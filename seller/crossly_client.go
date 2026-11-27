// Hand-written client ergonomics, vendored into the generated package by
// codegen/integrations/extras/go/finish.mjs. Not generated: edit it in the
// Crossly monorepo (tools/sdk-repo/extras/go/), not here.

package crossly

import (
	"crypto/rand"
	"encoding/hex"
	"io"
	"math"
	mrand "math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// UserAgent is sent on every request made through NewClient.
const UserAgent = "crossly-go/" + SDKVersion

// ClientOption tunes NewClient.
type ClientOption func(*clientSettings)

type clientSettings struct {
	baseURL    string
	timeout    time.Duration
	transport  http.RoundTripper
	retry      RetryTransport
	httpClient *http.Client
}

// WithBaseURL points the client at another server (default https://crossly.net/api).
func WithBaseURL(url string) ClientOption { return func(s *clientSettings) { s.baseURL = url } }

// WithTimeout bounds each attempt, including reading the response body (default 30s).
func WithTimeout(d time.Duration) ClientOption { return func(s *clientSettings) { s.timeout = d } }

// WithMaxRetries sets how many times a retryable failure is retried (default 3; 0 disables retries).
func WithMaxRetries(n int) ClientOption { return func(s *clientSettings) { s.retry.MaxRetries = n } }

// WithMaxRetryAfter caps how long a Retry-After header may make the client wait
// (default 60s). A longer Retry-After is returned to the caller instead of slept through.
func WithMaxRetryAfter(d time.Duration) ClientOption {
	return func(s *clientSettings) { s.retry.MaxRetryAfter = d }
}

// WithoutIdempotencyKeys stops NewClient adding an Idempotency-Key to POST and PATCH requests.
// Writes are then retried only on 429, which the API returns before doing anything.
func WithoutIdempotencyKeys() ClientOption {
	return func(s *clientSettings) { s.retry.NoIdempotencyKeys = true }
}

// WithTransport sets the RoundTripper under the retry layer (default http.DefaultTransport).
func WithTransport(rt http.RoundTripper) ClientOption { return func(s *clientSettings) { s.transport = rt } }

// NewClient returns the generated APIClient, authenticated with a Personal Access
// Token (crossly_pat_…) or an OAuth token, behind a RetryTransport:
//
//   - 429 and 5xx responses and network errors are retried with exponential
//     backoff and jitter, honouring Retry-After;
//   - POST and PATCH get an Idempotency-Key, the same on every retry, so a
//     retried write is applied once (the API replays the first response);
//   - every request carries User-Agent crossly-go/<version>.
//
// Every generated service is on the result: client.ListingsApi, client.OrdersApi, …
func NewClient(token string, opts ...ClientOption) *APIClient {
	s := clientSettings{timeout: 30 * time.Second, retry: RetryTransport{MaxRetries: 3}}
	for _, o := range opts {
		o(&s)
	}
	s.retry.Base = s.transport
	cfg := NewConfiguration()
	cfg.UserAgent = UserAgent
	if s.baseURL != "" {
		cfg.Servers = ServerConfigurations{{URL: strings.TrimRight(s.baseURL, "/")}}
	}
	if token != "" {
		cfg.AddDefaultHeader("Authorization", "Bearer "+token)
	}
	rt := s.retry
	cfg.HTTPClient = &http.Client{Timeout: s.timeout, Transport: &rt}
	return NewAPIClient(cfg)
}

// RetryTransport is the http.RoundTripper NewClient installs. Use it directly to
// add Crossly's retry policy to a client you configure yourself:
//
//	cfg := crossly.NewConfiguration()
//	cfg.HTTPClient = &http.Client{Transport: &crossly.RetryTransport{MaxRetries: 3}}
type RetryTransport struct {
	// Base performs each attempt. Nil means http.DefaultTransport.
	Base http.RoundTripper
	// MaxRetries is the number of retries after the first attempt.
	MaxRetries int
	// MaxRetryAfter caps a server-requested wait. Zero means 60s.
	MaxRetryAfter time.Duration
	// NoIdempotencyKeys disables the automatic Idempotency-Key on POST and PATCH.
	NoIdempotencyKeys bool
}

const (
	backoffBase = 500 * time.Millisecond
	backoffCap  = 8 * time.Second
)

// RoundTrip implements http.RoundTripper.
func (t *RetryTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	base := t.Base
	if base == nil {
		base = http.DefaultTransport
	}
	write := req.Method == http.MethodPost || req.Method == http.MethodPatch
	if write && !t.NoIdempotencyKeys && req.Header.Get("Idempotency-Key") == "" {
		req = req.Clone(req.Context())
		req.Header.Set("Idempotency-Key", NewIdempotencyKey())
	}
	if req.Header.Get("User-Agent") == "" {
		req = req.Clone(req.Context())
		req.Header.Set("User-Agent", UserAgent)
	}
	keyed := req.Header.Get("Idempotency-Key") != ""
	maxWait := t.MaxRetryAfter
	if maxWait == 0 {
		maxWait = 60 * time.Second
	}

	for attempt := 0; ; attempt++ {
		try := req
		if attempt > 0 {
			try = req.Clone(req.Context())
			if req.Body != nil && req.Body != http.NoBody {
				body, err := req.GetBody()
				if err != nil {
					return nil, err
				}
				try.Body = body
			}
		}
		resp, err := base.RoundTrip(try)

		// May this failure be retried at all?
		canRetry := attempt < t.MaxRetries && req.Context().Err() == nil &&
			(req.Body == nil || req.Body == http.NoBody || req.GetBody != nil)
		if canRetry {
			switch {
			case err != nil:
				// The request may have reached the server: a write is only safe to repeat with a key.
				canRetry = !write || keyed
			case resp.StatusCode == http.StatusTooManyRequests:
				// Rate limited before the handler ran: always safe.
			case resp.StatusCode >= 500 && resp.StatusCode != http.StatusNotImplemented:
				canRetry = !write || keyed
			default:
				canRetry = false
			}
		}
		if !canRetry {
			return resp, err
		}

		delay := backoff(attempt)
		if resp != nil {
			if wait, ok := RetryAfter(resp.Header); ok {
				if wait > maxWait {
					return resp, nil
				}
				delay = wait
			}
			_, _ = io.Copy(io.Discard, io.LimitReader(resp.Body, 1<<16))
			resp.Body.Close()
		}
		timer := time.NewTimer(delay)
		select {
		case <-req.Context().Done():
			timer.Stop()
			return nil, req.Context().Err()
		case <-timer.C:
		}
	}
}

// backoff is full-jitter exponential: a random wait in [0, min(cap, base*2^attempt)).
func backoff(attempt int) time.Duration {
	ceiling := float64(backoffBase) * math.Pow(2, float64(attempt))
	if ceiling > float64(backoffCap) {
		ceiling = float64(backoffCap)
	}
	return time.Duration(mrand.Int63n(int64(ceiling)) + 1)
}

// RetryAfter reads a Retry-After header: delta-seconds or an HTTP date (RFC 9110).
func RetryAfter(h http.Header) (time.Duration, bool) {
	v := strings.TrimSpace(h.Get("Retry-After"))
	if v == "" {
		return 0, false
	}
	if secs, err := strconv.ParseFloat(v, 64); err == nil && secs >= 0 {
		return time.Duration(secs * float64(time.Second)), true
	}
	if at, err := http.ParseTime(v); err == nil {
		d := time.Until(at)
		if d < 0 {
			d = 0
		}
		return d, true
	}
	return 0, false
}

// NewIdempotencyKey returns a random key for the Idempotency-Key header. The API
// replays the first response to a POST or PATCH that repeats a key within 24 hours.
// Set the header yourself when the same logical write may be sent from more than one process.
func NewIdempotencyKey() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		return strconv.FormatInt(time.Now().UnixNano(), 36)
	}
	return hex.EncodeToString(b[:])
}

