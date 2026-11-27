// net/http middleware around Verify. Hand-written, vendored next to the
// verifier by codegen/integrations/extras/go/finish.mjs; edit it in the Crossly
// monorepo (tools/sdk-repo/extras/go/), not here.

package webhooks

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// SignatureHeader is the header Crossly signs every delivery with.
const SignatureHeader = "Crossly-Signature"

// MaxBodyBytes caps how much of a delivery is read (1 MiB); a larger body is a 413.
const MaxBodyBytes = 1 << 20

type ctxKey struct{}

// WithEvent returns a copy of ctx carrying a verified event.
func WithEvent(ctx context.Context, ev *Event) context.Context { return context.WithValue(ctx, ctxKey{}, ev) }

// EventFromContext returns the event Middleware verified for this request.
func EventFromContext(ctx context.Context) (*Event, bool) {
	ev, ok := ctx.Value(ctxKey{}).(*Event)
	return ev, ok && ev != nil
}

// ErrBodyTooLarge is returned by VerifyRequest for a body over MaxBodyBytes.
var ErrBodyTooLarge = errors.New("webhook body exceeds MaxBodyBytes")

// VerifyRequest reads r.Body (the exact bytes, never re-serialised), verifies the
// Crossly-Signature header and returns the event and the raw body. r.Body is
// replaced with a reader over the same bytes, so a later handler can read it again.
func VerifyRequest(r *http.Request, secret string, opts *Options) (*Event, []byte, error) {
	raw, err := io.ReadAll(io.LimitReader(r.Body, MaxBodyBytes+1))
	_ = r.Body.Close()
	if err != nil {
		return nil, nil, err
	}
	if len(raw) > MaxBodyBytes {
		return nil, nil, ErrBodyTooLarge
	}
	r.Body = io.NopCloser(bytes.NewReader(raw))
	ev, err := Verify(raw, r.Header.Get(SignatureHeader), secret, opts)
	if err != nil {
		return nil, raw, err
	}
	return ev, raw, nil
}

// StatusFor is the HTTP status a receiver should answer a failed verification with.
func StatusFor(err error) int {
	if errors.Is(err, ErrBodyTooLarge) {
		return http.StatusRequestEntityTooLarge
	}
	var e *Error
	if errors.As(err, &e) {
		if e.Reason == ReasonMissingSecret {
			return http.StatusInternalServerError
		}
		return http.StatusBadRequest
	}
	return http.StatusBadRequest
}

// ErrorBody is the JSON a receiver answers a failed verification with:
// {"error":{"code":"<reason>","message":"…"}}.
func ErrorBody(err error) map[string]any {
	code := "invalid_request"
	var e *Error
	if errors.As(err, &e) {
		code = e.Reason
	} else if errors.Is(err, ErrBodyTooLarge) {
		code = "body_too_large"
	}
	return map[string]any{"error": map[string]string{"code": code, "message": err.Error()}}
}

// WriteError answers a failed verification.
func WriteError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(StatusFor(err))
	_ = json.NewEncoder(w).Encode(ErrorBody(err))
}

// Middleware verifies every request before next runs and puts the event on the
// request context (EventFromContext). An unverified request gets a 400 and never
// reaches next. It is a plain func(http.Handler) http.Handler, so it also plugs
// straight into chi (r.Use / r.With) and anything else built on net/http.
func Middleware(secret string, opts *Options) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ev, _, err := VerifyRequest(r, secret, opts)
			if err != nil {
				WriteError(w, err)
				return
			}
			next.ServeHTTP(w, r.WithContext(WithEvent(r.Context(), ev)))
		})
	}
}

// Handler is Middleware for a single endpoint: fn runs only for verified events.
// Answer 2xx quickly; Crossly retries anything else.
func Handler(secret string, opts *Options, fn func(w http.ResponseWriter, r *http.Request, ev *Event)) http.Handler {
	return Middleware(secret, opts)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ev, _ := EventFromContext(r.Context())
		fn(w, r, ev)
	}))
}
