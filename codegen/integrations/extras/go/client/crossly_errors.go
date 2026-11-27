// Hand-written client ergonomics, vendored into the generated package by
// codegen/integrations/extras/go/finish.mjs. Edit it in the Crossly monorepo
// (tools/sdk-repo/extras/go/), not here.

package crossly

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// APIError is a failed call, decoded from the API's error envelope
//
//	{ "error": { "code", "message", "details" }, "correlationId" }
//
// Code is the stable machine key to branch on (not_found, rate_limited,
// idempotency_key_reused, …); CorrelationID is what Crossly support traces.
type APIError struct {
	StatusCode    int
	Code          string
	Message       string
	Details       json.RawMessage
	CorrelationID string
	// RetryAfter is the server's Retry-After, when it sent one (zero otherwise).
	RetryAfter time.Duration
	// Body is the raw response body.
	Body []byte
}

func (e *APIError) Error() string {
	msg := e.Message
	if msg == "" {
		msg = http.StatusText(e.StatusCode)
	}
	s := fmt.Sprintf("crossly: %d %s: %s", e.StatusCode, e.Code, msg)
	if e.CorrelationID != "" {
		s += " (correlationId " + e.CorrelationID + ")"
	}
	return s
}

// AsAPIError turns what a generated call returned into an *APIError. Pass both
// the error and the *http.Response:
//
//	page, resp, err := client.ListingsApi.ListListings(ctx).Execute()
//	if apiErr, ok := crossly.AsAPIError(err, resp); ok && apiErr.Code == "rate_limited" { … }
//
// It returns false for a nil error and for failures that never reached the API
// (DNS, TLS, a cancelled context): inspect err itself for those.
func AsAPIError(err error, resp *http.Response) (*APIError, bool) {
	if err == nil {
		return nil, false
	}
	var already *APIError
	if errors.As(err, &already) {
		return already, true
	}
	var body []byte
	var gp *GenericOpenAPIError
	var gv GenericOpenAPIError
	switch {
	case errors.As(err, &gp) && gp != nil:
		body = gp.Body()
	case errors.As(err, &gv):
		body = gv.Body()
	default:
		return nil, false
	}
	out := &APIError{Body: body}
	if resp != nil {
		out.StatusCode = resp.StatusCode
		out.CorrelationID = resp.Header.Get("X-Correlation-Id")
		out.RetryAfter, _ = RetryAfter(resp.Header)
	} else if code, convErr := strconv.Atoi(strings.SplitN(err.Error(), " ", 2)[0]); convErr == nil {
		out.StatusCode = code
	}
	var env struct {
		Error *struct {
			Code    string          `json:"code"`
			Message string          `json:"message"`
			Details json.RawMessage `json:"details"`
		} `json:"error"`
		CorrelationID string `json:"correlationId"`
	}
	if json.Unmarshal(body, &env) == nil {
		if env.Error != nil {
			out.Code, out.Message, out.Details = env.Error.Code, env.Error.Message, env.Error.Details
		}
		if env.CorrelationID != "" {
			out.CorrelationID = env.CorrelationID
		}
	}
	if out.Message == "" {
		out.Message = err.Error()
	}
	return out, true
}
