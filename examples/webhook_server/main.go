// A minimal Crossly webhook receiver on net/http. The signature is checked over
// the exact bytes received, before anything parses them.
//
//	go get github.com/crossly-dev/crossly-go/seller
//	CROSSLY_WEBHOOK_SECRET=... go run .
//	# then point a webhook endpoint (Settings → Webhooks) at http://<host>:8080/webhooks/crossly
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/crossly-dev/crossly-go/seller/webhooks"
)

func main() {
	secret := os.Getenv("CROSSLY_WEBHOOK_SECRET")
	mux := http.NewServeMux()
	mux.Handle("POST /webhooks/crossly", webhooks.Handler(secret, nil, func(w http.ResponseWriter, r *http.Request, ev *webhooks.Event) {
		log.Printf("%s %s: %s", ev.ID, ev.Type, ev.Data)
		// Answer fast; do slow work after replying, or Crossly retries the delivery.
		w.WriteHeader(http.StatusNoContent)
	}))
	log.Fatal(http.ListenAndServe(":8080", mux))
}
