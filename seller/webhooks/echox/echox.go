// Package echox is Crossly webhook verification for Echo (v4). It is its own Go
// module so the core SDK (github.com/crossly-dev/crossly-go/seller) never pulls in Echo.
//
//	go get github.com/crossly-dev/crossly-go/seller/webhooks/echox
//
//	e.POST("/webhooks/crossly", func(c echo.Context) error {
//		ev := echox.Event(c)
//		log.Println(ev.Type, ev.ID)
//		return c.NoContent(http.StatusNoContent)
//	}, echox.Middleware(os.Getenv("CROSSLY_WEBHOOK_SECRET"), nil))
//
// Hand-written; edit it in the Crossly monorepo (tools/sdk-repo/extras/go/echox/).
package echox

import (
	"github.com/crossly-dev/crossly-go/seller/webhooks"
	"github.com/labstack/echo/v4"
)

// ContextKey is the echo.Context key the verified *webhooks.Event is stored under.
const ContextKey = "crossly.webhook.event"

// Middleware verifies the Crossly-Signature over the raw body. A request that does
// not verify is answered with a 400 JSON error and never reaches the handler; a
// verified one runs with the event on the context (Event). The body stays readable.
func Middleware(secret string, opts *webhooks.Options) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			req := c.Request()
			ev, _, err := webhooks.VerifyRequest(req, secret, opts)
			if err != nil {
				return c.JSON(webhooks.StatusFor(err), webhooks.ErrorBody(err))
			}
			c.Set(ContextKey, ev)
			c.SetRequest(req.WithContext(webhooks.WithEvent(req.Context(), ev)))
			return next(c)
		}
	}
}

// Event returns the event Middleware verified, or nil outside it.
func Event(c echo.Context) *webhooks.Event {
	ev, _ := c.Get(ContextKey).(*webhooks.Event)
	return ev
}
