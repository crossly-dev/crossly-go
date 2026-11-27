// Tests for the hand-written ergonomics. Copied in by codegen/ci/check-go.sh for
// the check and removed afterwards: not shipped.

package crossly

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"
)

func TestCrosslyRetryHonoursRetryAfterAndKeepsIdempotencyKey(t *testing.T) {
	var calls int32
	keys := make(chan string, 4)
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		n := atomic.AddInt32(&calls, 1)
		keys <- r.Header.Get("Idempotency-Key")
		if got := r.Header.Get("User-Agent"); got != UserAgent {
			t.Errorf("User-Agent = %q, want %q", got, UserAgent)
		}
		if got := r.Header.Get("Authorization"); got != "Bearer crossly_pat_test" {
			t.Errorf("Authorization = %q", got)
		}
		body, _ := io.ReadAll(r.Body)
		if string(body) != `{"a":1}` {
			t.Errorf("attempt %d body = %q", n, body)
		}
		if n == 1 {
			w.Header().Set("Retry-After", "0")
			w.WriteHeader(http.StatusTooManyRequests)
			_, _ = w.Write([]byte(`{"error":{"code":"rate_limited","message":"slow down"},"correlationId":"c1"}`))
			return
		}
		if n == 2 {
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}))
	defer srv.Close()

	client := NewClient("crossly_pat_test", WithBaseURL(srv.URL))
	req, _ := http.NewRequest(http.MethodPost, srv.URL+"/v1/x", strings.NewReader(`{"a":1}`))
	req.Header.Set("User-Agent", UserAgent)
	req.Header.Set("Authorization", "Bearer crossly_pat_test")
	resp, err := client.GetConfig().HTTPClient.Do(req)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != http.StatusCreated || atomic.LoadInt32(&calls) != 3 {
		t.Fatalf("status %d after %d calls", resp.StatusCode, calls)
	}
	k1, k2, k3 := <-keys, <-keys, <-keys
	if k1 == "" || k1 != k2 || k2 != k3 {
		t.Fatalf("idempotency keys differ or are empty: %q %q %q", k1, k2, k3)
	}
}

func TestCrosslyRetryGivesUpOnLongRetryAfter(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.Header().Set("Retry-After", "3600")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer srv.Close()
	rt := &RetryTransport{MaxRetries: 3}
	resp, err := (&http.Client{Transport: rt}).Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if resp.StatusCode != 429 || calls != 1 {
		t.Fatalf("status %d after %d calls, want one 429", resp.StatusCode, calls)
	}
}

func TestCrosslyNoRetryOn4xx(t *testing.T) {
	var calls int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&calls, 1)
		w.WriteHeader(http.StatusNotFound)
	}))
	defer srv.Close()
	resp, err := (&http.Client{Transport: &RetryTransport{MaxRetries: 3}}).Get(srv.URL)
	if err != nil {
		t.Fatal(err)
	}
	resp.Body.Close()
	if calls != 1 {
		t.Fatalf("%d calls for a 404", calls)
	}
}

func TestCrosslyRetryAfterDate(t *testing.T) {
	h := http.Header{}
	h.Set("Retry-After", time.Now().Add(2*time.Second).UTC().Format(http.TimeFormat))
	d, ok := RetryAfter(h)
	if !ok || d <= 0 || d > 3*time.Second {
		t.Fatalf("RetryAfter(date) = %v, %v", d, ok)
	}
}

func TestCrosslyAsAPIError(t *testing.T) {
	err := &GenericOpenAPIError{
		body:  []byte(`{"error":{"code":"not_found","message":"No such listing","details":{"id":"x"}},"correlationId":"corr_1"}`),
		error: "404 Not Found",
	}
	resp := &http.Response{StatusCode: 404, Header: http.Header{}}
	apiErr, ok := AsAPIError(err, resp)
	if !ok || apiErr.Code != "not_found" || apiErr.CorrelationID != "corr_1" || apiErr.StatusCode != 404 {
		t.Fatalf("AsAPIError = %+v, %v", apiErr, ok)
	}
	if _, ok := AsAPIError(errors.New("dial tcp: refused"), nil); ok {
		t.Fatal("a transport error is not an APIError")
	}
}

func TestCrosslyPaginate(t *testing.T) {
	total := int32(5)
	pages := 0
	var got []int
	for v, err := range Paginate(context.Background(), 2, func(_ context.Context, page, limit int32) ([]int, *Pagination, error) {
		pages++
		all := []int{1, 2, 3, 4, 5}
		lo := int((page - 1) * limit)
		hi := lo + int(limit)
		if hi > len(all) {
			hi = len(all)
		}
		return all[lo:hi], &Pagination{Page: page, Limit: limit, Total: &total}, nil
	}) {
		if err != nil {
			t.Fatal(err)
		}
		got = append(got, v)
	}
	if len(got) != 5 || pages != 3 {
		t.Fatalf("got %v over %d pages", got, pages)
	}
	// Breaking early fetches no further pages.
	pages = 0
	for range Paginate(context.Background(), 2, func(_ context.Context, page, limit int32) ([]int, *Pagination, error) {
		pages++
		return []int{1, 2}, &Pagination{Page: page, Limit: limit}, nil
	}) {
		break
	}
	if pages != 1 {
		t.Fatalf("break fetched %d pages", pages)
	}
}
