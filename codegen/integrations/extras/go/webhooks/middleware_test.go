// Copied in by codegen/ci/check-go.sh for the check and removed afterwards: not shipped.

package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"io"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"
)

func sign(secret, body string, t int64) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(strconv.FormatInt(t, 10) + "." + body))
	return "t=" + strconv.FormatInt(t, 10) + ",v1=" + hex.EncodeToString(mac.Sum(nil))
}

func TestMiddleware(t *testing.T) {
	const secret = "whsec_test"
	body := `{"id":"evt_1","type":"listing.sold","created":"2026-01-01T00:00:00Z","data":{"x":1}}`
	var seen *Event
	var reread string
	h := Middleware(secret, nil)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		seen, _ = EventFromContext(r.Context())
		b, _ := io.ReadAll(r.Body)
		reread = string(b)
		w.WriteHeader(http.StatusNoContent)
	}))

	good := httptest.NewRequest(http.MethodPost, "/hooks", strings.NewReader(body))
	good.Header.Set(SignatureHeader, sign(secret, body, time.Now().Unix()))
	rec := httptest.NewRecorder()
	h.ServeHTTP(rec, good)
	if rec.Code != http.StatusNoContent || seen == nil || seen.Type != "listing.sold" || reread != body {
		t.Fatalf("verified delivery: code %d, event %+v, body %q", rec.Code, seen, reread)
	}

	seen = nil
	bad := httptest.NewRequest(http.MethodPost, "/hooks", strings.NewReader(body+" "))
	bad.Header.Set(SignatureHeader, sign(secret, body, time.Now().Unix()))
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, bad)
	if rec.Code != http.StatusBadRequest || seen != nil || !strings.Contains(rec.Body.String(), ReasonBadSignature) {
		t.Fatalf("tampered delivery: code %d, body %s", rec.Code, rec.Body.String())
	}

	old := httptest.NewRequest(http.MethodPost, "/hooks", strings.NewReader(body))
	old.Header.Set(SignatureHeader, sign(secret, body, time.Now().Unix()-3600))
	rec = httptest.NewRecorder()
	h.ServeHTTP(rec, old)
	if rec.Code != http.StatusBadRequest || !strings.Contains(rec.Body.String(), ReasonOutOfTolerance) {
		t.Fatalf("replayed delivery: code %d, body %s", rec.Code, rec.Body.String())
	}
}
