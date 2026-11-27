// Print every listing in your Crossly store, page by page, with retries,
// backoff and typed errors from crossly.NewClient.
//
//	go get github.com/crossly-dev/crossly-go/seller
//	CROSSLY_TOKEN=crossly_pat_... go run .
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	crossly "github.com/crossly-dev/crossly-go/seller"
)

func main() {
	ctx := context.Background()
	client := crossly.NewClient(os.Getenv("CROSSLY_TOKEN"))

	rows := crossly.Paginate(ctx, 100, func(ctx context.Context, page, limit int32) ([]crossly.ListListingsItem, *crossly.Pagination, error) {
		r, resp, err := client.ListingsApi.ListListings(ctx).Page(page).Limit(limit).Execute()
		if err != nil {
			if apiErr, ok := crossly.AsAPIError(err, resp); ok {
				return nil, nil, apiErr
			}
			return nil, nil, err
		}
		return r.Data, r.Pagination, nil
	})
	for item, err := range rows {
		if err != nil {
			log.Fatal(err) // crossly: 401 unauthorized: … (correlationId …)
		}
		fmt.Printf("%s  %-10s  %s\n", item.GetId(), item.GetStatus(), item.GetTitle())
	}
}
