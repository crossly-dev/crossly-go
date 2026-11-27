module github.com/crossly-dev/crossly-go/seller/webhooks/echox

go 1.25.0

require (
	github.com/crossly-dev/crossly-go/seller v0.2.0
	github.com/labstack/echo/v4 v4.16.0
)

require (
	github.com/labstack/gommon v0.5.0 // indirect
	github.com/mattn/go-colorable v0.1.15 // indirect
	github.com/mattn/go-isatty v0.0.22 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	golang.org/x/crypto v0.53.0 // indirect
	golang.org/x/net v0.56.0 // indirect
	golang.org/x/sys v0.46.0 // indirect
	golang.org/x/text v0.40.0 // indirect
)

// Builds inside this repository against the seller module next to it. Ignored by
// consumers (replace only applies in the main module), who get seller v0.2.0.
replace github.com/crossly-dev/crossly-go/seller => ../..
