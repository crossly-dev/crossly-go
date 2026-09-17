// Package webhooks verifies Crossly webhook deliveries.
//
//	Crossly-Signature: t=<unix seconds>,v1=<hex HMAC-SHA256>
//
// signed over "<t>.<rawBody>" with the endpoint's signing secret.
//
// Three ways to get this wrong, all silent:
//
//  1. Verifying a re-serialised body. Unmarshal then Marshal does not
//     round-trip byte for byte — Go sorts map keys and reformats numbers — so
//     genuine payloads fail and the usual fix is to stop verifying. Pass the
//     bytes you read off the wire.
//  2. Comparing with ==. A string compare returns early on the first differing
//     byte; hmac.Equal does not.
//  3. Ignoring the timestamp. Without it a captured request replays forever.
//     The timestamp is INSIDE the signed message, so it cannot be edited to
//     look fresh.
//
// Standard library only.
package webhooks

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// DefaultToleranceSeconds is how far the timestamp may be from now.
//
// Five minutes rather than five seconds: a retried delivery can sit in a queue
// and other people's clocks are not ours. Tighter rejects genuine traffic.
const DefaultToleranceSeconds = 300

// Reason codes. A caller should distinguish these — a forged request and a
// clock problem want different responses and different alerts.
const (
	ReasonMalformedHeader = "malformed_header"
	ReasonBadSignature    = "bad_signature"
	ReasonOutOfTolerance  = "timestamp_out_of_tolerance"
	ReasonMissingSecret   = "missing_secret"
)

// Error is returned when a webhook does not verify.
type Error struct {
	Reason  string
	Message string
}

func (e *Error) Error() string { return e.Message }

// Event is the envelope Crossly sends.
type Event struct {
	ID      string          `json:"id"`
	Type    string          `json:"type"`
	Created string          `json:"created"`
	Data    json.RawMessage `json:"data"`
}

// Options tunes verification. The zero value means "defaults".
type Options struct {
	// ToleranceSeconds defaults to DefaultToleranceSeconds when zero.
	ToleranceSeconds int
	// Now overrides the clock, for tests. Zero means time.Now().
	Now int64
}

// parseSignatureHeader pulls t and v1 out of the header.
//
// Field-wise rather than one regex, so a future v2= alongside v1= does not
// break existing verifiers — which is the entire reason the scheme is
// versioned.
func parseSignatureHeader(header string) (int64, string, bool) {
	var (
		t     int64
		v1    string
		haveT bool
	)

	for _, part := range strings.Split(header, ",") {
		key, value, found := strings.Cut(part, "=")
		if !found {
			continue
		}
		key = strings.TrimSpace(key)
		value = strings.TrimSpace(value)

		switch key {
		case "t":
			parsed, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				return 0, "", false
			}
			t, haveT = parsed, true
		case "v1":
			v1 = value
		}
	}

	if !haveT || v1 == "" {
		return 0, "", false
	}
	return t, v1, true
}

// Verify checks a webhook and returns the parsed event.
//
// rawBody must be the EXACT bytes received — io.ReadAll(r.Body), not a
// re-marshalled struct. See the package comment.
//
// It returns an error rather than a bool so a caller who ignores the result
// does not silently accept forged events; `_ = Verify(...)` is at least
// visible in review, where `if ok` that was never written is not.
func Verify(rawBody []byte, signatureHeader, secret string, opts *Options) (*Event, error) {
	if secret == "" {
		return nil, &Error{ReasonMissingSecret, "A webhook signing secret is required."}
	}
	if signatureHeader == "" {
		return nil, &Error{ReasonMalformedHeader, "No Crossly-Signature header on the request."}
	}

	timestamp, provided, ok := parseSignatureHeader(signatureHeader)
	if !ok {
		preview := signatureHeader
		if len(preview) > 60 {
			preview = preview[:60]
		}
		return nil, &Error{
			ReasonMalformedHeader,
			fmt.Sprintf(`Could not parse Crossly-Signature: expected "t=<unix>,v1=<hex>", got %q.`, preview),
		}
	}

	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(strconv.FormatInt(timestamp, 10) + "." + string(rawBody)))
	expected := hex.EncodeToString(mac.Sum(nil))

	// hmac.Equal, not ==. It is constant time and safe on a length mismatch,
	// which a truncated signature produces.
	if !hmac.Equal([]byte(expected), []byte(provided)) {
		return nil, &Error{
			ReasonBadSignature,
			"Signature did not match. If genuine payloads are failing, you are almost " +
				"certainly verifying a re-serialised body — pass the bytes you read off the wire.",
		}
	}

	// Freshness AFTER the signature, so an attacker learns nothing about
	// timestamps without already holding a valid signature.
	tolerance := opts.tolerance()
	drift := opts.now() - timestamp
	if drift < 0 {
		drift = -drift
	}
	if drift > int64(tolerance) {
		return nil, &Error{
			ReasonOutOfTolerance,
			fmt.Sprintf(
				"Timestamp is %ds away from now (tolerance %ds). This is a replay guard — "+
					"if it fires on live traffic, check your server clock.", drift, tolerance),
		}
	}

	var event Event
	if err := json.Unmarshal(rawBody, &event); err != nil {
		return nil, &Error{ReasonMalformedHeader, "Body verified but is not valid JSON: " + err.Error()}
	}
	return &event, nil
}

func (o *Options) tolerance() int {
	if o == nil || o.ToleranceSeconds == 0 {
		return DefaultToleranceSeconds
	}
	return o.ToleranceSeconds
}

func (o *Options) now() int64 {
	if o == nil || o.Now == 0 {
		return time.Now().Unix()
	}
	return o.Now
}
