// Runs the shared vectors against the Go port.
package main

import (
	"encoding/json"
	"fmt"
	"os"

	webhooks "github.com/crossly-dev/crossly-go/webhooks"
)

type vector struct {
	Name   string `json:"name"`
	Body   string `json:"body"`
	Header string `json:"header"`
	Secret string `json:"secret"`
	Now    int64  `json:"now"`
	Expect string `json:"expect"`
}

func main() {
	raw, err := os.ReadFile("../../test-vectors.json")
	if err != nil {
		fmt.Println("cannot read vectors:", err)
		os.Exit(1)
	}
	var file struct {
		ToleranceSeconds int      `json:"toleranceSeconds"`
		Vectors          []vector `json:"vectors"`
	}
	if err := json.Unmarshal(raw, &file); err != nil {
		fmt.Println("bad vectors:", err)
		os.Exit(1)
	}

	ok, bad := 0, 0
	for _, v := range file.Vectors {
		got := "ok"
		_, err := webhooks.Verify([]byte(v.Body), v.Header, v.Secret,
			&webhooks.Options{ToleranceSeconds: file.ToleranceSeconds, Now: v.Now})
		if err != nil {
			if e, isWebhookErr := err.(*webhooks.Error); isWebhookErr {
				got = e.Reason
			} else {
				got = "threw"
			}
		}
		if got == v.Expect {
			ok++
		} else {
			bad++
			fmt.Printf("  MISMATCH %s: expected %s, got %s\n", v.Name, v.Expect, got)
		}
	}
	fmt.Printf("  go: %d/%d vectors agree\n", ok, len(file.Vectors))
	if bad > 0 {
		os.Exit(1)
	}
}
