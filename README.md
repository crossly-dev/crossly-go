# Crossly Go SDK

[![ci](https://github.com/crossly-dev/crossly-go/actions/workflows/ci.yml/badge.svg)](https://github.com/crossly-dev/crossly-go/actions/workflows/ci.yml) [![release](https://github.com/crossly-dev/crossly-go/actions/workflows/release.yml/badge.svg)](https://github.com/crossly-dev/crossly-go/actions/workflows/release.yml) [![license](https://img.shields.io/badge/license-MIT-blue)](LICENSE)

Official Go clients for the [Crossly](https://crossly.net) API: list once and sell across every marketplace, read
orders, manage inventory, and receive signed webhooks. Generated from the public OpenAPI specs and released
automatically when they change, with signed build provenance.

## Install

| client | for | install |
|---|---|---|
| seller | your store: listings, crossposting, orders, analytics | `go get github.com/crossly-dev/crossly-go/seller` |
| buyer | shopping: cart, orders, wishlists, offers, cashback | `go get github.com/crossly-dev/crossly-go/buyer` |

## Two clients, two credentials

| client | credential |
|---|---|
| seller | a Personal Access Token (`crossly_pat_…`, Settings → Developer) or a seller OAuth token |
| buyer | a buyer OAuth token (`crossly_oat_…`) with `buyer:*` scopes |

They are not interchangeable: a seller token is refused by the buyer API and the reverse, even for the same
Crossly account. A 401 from a call that looks right usually means the other client's token.

## First call

```go
import (
	"context"
	"fmt"
	"log"
	"os"

	crossly "github.com/crossly-dev/crossly-go/seller"
)

// Retries 429/5xx with backoff (honouring Retry-After), adds an Idempotency-Key to POST/PATCH,
// sends User-Agent crossly-go/<version>; every generated service hangs off the client.
client := crossly.NewClient(os.Getenv("CROSSLY_TOKEN")) // crossly_pat_…
page, resp, err := client.ListingsApi.ListListings(context.Background()).Page(1).Limit(25).Execute()
if err != nil {
	if apiErr, ok := crossly.AsAPIError(err, resp); ok {
		log.Fatalf("%s (%s, correlationId %s)", apiErr.Message, apiErr.Code, apiErr.CorrelationID)
	}
	log.Fatal(err)
}
for _, item := range page.Data {
	fmt.Println(item.GetId(), item.GetTitle(), item.GetStatus())
}
```

List calls return `{ data, pagination }` and page with `page` and `limit`. Money fields named `…_cents` are
integer cents. A failed call returns an error; `crossly.AsAPIError(err, resp)` gives an `*APIError` with `StatusCode`, `Code`, `Message`, `Details` and the `CorrelationID` support can trace.

## Webhooks

Crossly signs every webhook with `Crossly-Signature: t=<unix>,v1=<hex HMAC-SHA256 of "t.body">`. The verifier
ships in the seller and buyer packages and checks the signature in constant time, rejects a timestamp more than
five minutes old, and returns the parsed event:

```go
import "github.com/crossly-dev/crossly-go/seller/webhooks"

// net/http (and chi, which takes the same func(http.Handler) http.Handler):
mux.Handle("POST /webhooks/crossly", webhooks.Handler(secret, nil, func(w http.ResponseWriter, r *http.Request, ev *webhooks.Event) {
	log.Println(ev.Type, ev.ID)
	w.WriteHeader(http.StatusNoContent)
}))

// or by hand, over the exact bytes received:
event, err := webhooks.Verify(rawBody, r.Header.Get("Crossly-Signature"), secret, nil)
```

Verify the raw bytes, not a re-serialised body. `webhooks/` holds the same verifier on its own, with the shared
test vectors every language's port is checked against on each build.

## Client helpers

Reference: [pkg.go.dev/github.com/crossly-dev/crossly-go/seller](https://pkg.go.dev/github.com/crossly-dev/crossly-go/seller) ·
[buyer](https://pkg.go.dev/github.com/crossly-dev/crossly-go/buyer). Each client is its own module, tagged `seller/vX.Y.Z` and
`buyer/vX.Y.Z`; Go 1.23 or newer.

| helper | what it does |
|---|---|
| `crossly.NewClient(token, opts...)` | the generated `*APIClient`, authenticated, behind `RetryTransport`; options `WithTimeout` (default 30 s), `WithMaxRetries` (3), `WithMaxRetryAfter` (60 s), `WithBaseURL`, `WithTransport`, `WithoutIdempotencyKeys` |
| `crossly.RetryTransport` | the `http.RoundTripper` on its own: retries 429, 5xx and network errors with full-jitter exponential backoff, waits exactly `Retry-After` when sent, and gives a POST/PATCH one `Idempotency-Key` for all its attempts (the API replays the first response for 24 h) |
| `crossly.Paginate(ctx, limit, fetch)` | an `iter.Seq2[T, error]` over every row of a list endpoint, fetching pages as the loop reaches them |
| `crossly.AsAPIError(err, resp)` | the error envelope as `*APIError{StatusCode, Code, Message, Details, CorrelationID, RetryAfter}` |

```go
rows := crossly.Paginate(ctx, 100, func(ctx context.Context, page, limit int32) ([]crossly.ListListingsItem, *crossly.Pagination, error) {
	r, _, err := client.ListingsApi.ListListings(ctx).Page(page).Limit(limit).Execute()
	if err != nil {
		return nil, nil, err
	}
	return r.Data, r.Pagination, nil
})
for item, err := range rows {
	if err != nil {
		return err
	}
	fmt.Println(item.GetId())
}
```

The same helpers are in the buyer module (`crossly_buyer.NewClient`, …) for a buyer OAuth token.

## Webhook middleware

`github.com/crossly-dev/crossly-go/seller/webhooks` has no dependencies beyond the standard library:
`webhooks.Middleware(secret, opts)` is a `func(http.Handler) http.Handler` (net/http, chi, gorilla, …) that
answers an unverified delivery with a 400 and puts the event on the request context
(`webhooks.EventFromContext`); `webhooks.Handler` wraps one endpoint; `webhooks.VerifyRequest` is the
framework-free building block. Gin and Echo adapters are separate modules, so the SDK never pulls in a
framework you do not use:

```go
// chi
r.With(webhooks.Middleware(secret, nil)).Post("/webhooks/crossly", func(w http.ResponseWriter, r *http.Request) {
	ev, _ := webhooks.EventFromContext(r.Context())
	…
})

// Gin: go get github.com/crossly-dev/crossly-go/seller/webhooks/ginx
r.POST("/webhooks/crossly", ginx.Middleware(secret, nil), func(c *gin.Context) {
	ev := ginx.Event(c)
	…
})

// Echo: go get github.com/crossly-dev/crossly-go/seller/webhooks/echox
e.POST("/webhooks/crossly", func(c echo.Context) error {
	ev := echox.Event(c)
	…
}, echox.Middleware(secret, nil))
```

Deliveries larger than 1 MiB are refused with a 413. [examples/](examples/) has a runnable lister and receiver.

## Where things are

| client | API classes | models | API reference pages |
|---|---|---|---|
| seller | `seller/` `api_*.go` (package root) | `seller/` `model_<area>.go` (package root) | `seller/docs/` |
| buyer | `buyer/` `api_*.go` (package root) | `buyer/` `model_<area>.go` (package root) | `buyer/docs/` |

One class per API group (`OrdersApi`, `ListingsApi`, …), one model per type, the
models grouped one folder per product area (`vault`, `world`, `order`, …). A type
several endpoints share is ONE model, named after where it first appears
(`GetOrderResponse`, `ListOrdersItem`) or, for small shared shapes such as `{ ok }`,
after its fields (`OkResponse`). Each client README lists every method.

## Versions and changelog

Semantic versioning. While the version is 0.x, a release that removes an operation or a model bumps the minor
version; anything else bumps the patch. Every [CHANGELOG](CHANGELOG.md) entry lists the operations and models
that were added, changed or removed in the specs it was generated from.

## Verifying a release

Every release artifact is built on GitHub-hosted runners by [release.yml](.github/workflows/release.yml) and
carries a signed SLSA provenance attestation:

```sh
gh attestation verify <downloaded-file> --repo crossly-dev/crossly-go
```

## Contributing

`seller/`, `buyer/` and `webhooks/` are generated, so a pull request that edits them would be overwritten by the
next release. Open an issue instead: the fix goes into the API spec or the generator, and every language picks it
up. [CONTRIBUTING.md](CONTRIBUTING.md) explains how to regenerate locally.

## More

- API reference and guides: https://crossly.net/docs
- Other languages: https://github.com/crossly-dev
- TypeScript (Node and browsers): [`@crossly/sdk`](https://github.com/crossly-dev/crossly-sdk); command line: [`@crossly/cli`](https://github.com/crossly-dev/crossly-cli); AI agents: [`@crossly/mcp`](https://github.com/crossly-dev/crossly-mcp)

## License

[MIT](LICENSE)
