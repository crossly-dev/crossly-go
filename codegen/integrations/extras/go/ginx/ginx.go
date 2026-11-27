// Package ginx is Crossly webhook verification for Gin. It is its own Go module so
// the core SDK (github.com/crossly-dev/crossly-go/seller) never pulls in Gin.
//
//	go get github.com/crossly-dev/crossly-go/seller/webhooks/ginx
//
//	r.POST("/webhooks/crossly", ginx.Middleware(os.Getenv("CROSSLY_WEBHOOK_SECRET"), nil), func(c *gin.Context) {
//		ev := ginx.Event(c)
//		log.Println(ev.Type, ev.ID)
//		c.Status(http.StatusNoContent)
//	})
//
// Hand-written; edit it in the Crossly monorepo (tools/sdk-repo/extras/go/ginx/).
package ginx

import (
	"github.com/crossly-dev/crossly-go/seller/webhooks"
	"github.com/gin-gonic/gin"
)

// ContextKey is the gin.Context key the verified *webhooks.Event is stored under.
const ContextKey = "crossly.webhook.event"

// Middleware verifies the Crossly-Signature over the raw body. A request that does
// not verify is aborted with a 400 JSON error; a verified one continues with the
// event on the context (Event, or webhooks.EventFromContext(c.Request.Context())).
// The body stays readable by the next handler.
func Middleware(secret string, opts *webhooks.Options) gin.HandlerFunc {
	return func(c *gin.Context) {
		ev, _, err := webhooks.VerifyRequest(c.Request, secret, opts)
		if err != nil {
			c.AbortWithStatusJSON(webhooks.StatusFor(err), webhooks.ErrorBody(err))
			return
		}
		c.Set(ContextKey, ev)
		c.Request = c.Request.WithContext(webhooks.WithEvent(c.Request.Context(), ev))
		c.Next()
	}
}

// Event returns the event Middleware verified, or nil outside it.
func Event(c *gin.Context) *webhooks.Event {
	if v, ok := c.Get(ContextKey); ok {
		if ev, ok := v.(*webhooks.Event); ok {
			return ev
		}
	}
	return nil
}
