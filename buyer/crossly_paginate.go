// Hand-written client ergonomics, vendored into the generated package by
// codegen/integrations/extras/go/finish.mjs. Edit it in the Crossly monorepo
// (tools/sdk-repo/extras/go/), not here.

package crossly_buyer

import (
	"context"
	"iter"
)

// PageFunc fetches one page of a list endpoint: its rows and its pagination block.
type PageFunc[T any] func(ctx context.Context, page, limit int32) ([]T, *Pagination, error)

// Paginate walks every page of a list endpoint and yields each row, fetching the
// next page only when the loop reaches it. It stops at the last page (totalPages,
// total, or a short page) and on the first error, which it yields once.
//
//	rows := crossly.Paginate(ctx, 100, func(ctx context.Context, page, limit int32) ([]crossly.ListListingsItem, *crossly.Pagination, error) {
//		r, _, err := client.ListingsApi.ListListings(ctx).Page(page).Limit(limit).Execute()
//		if err != nil {
//			return nil, nil, err
//		}
//		return r.Data, r.Pagination, nil
//	})
//	for item, err := range rows {
//		if err != nil { return err }
//		fmt.Println(item.GetId())
//	}
func Paginate[T any](ctx context.Context, limit int32, fetch PageFunc[T]) iter.Seq2[T, error] {
	if limit <= 0 {
		limit = 100
	}
	return func(yield func(T, error) bool) {
		for page := int32(1); ; page++ {
			if err := ctx.Err(); err != nil {
				var zero T
				yield(zero, err)
				return
			}
			rows, p, err := fetch(ctx, page, limit)
			if err != nil {
				var zero T
				yield(zero, err)
				return
			}
			for _, row := range rows {
				if !yield(row, nil) {
					return
				}
			}
			if lastPage(len(rows), p, page, limit) {
				return
			}
		}
	}
}

func lastPage(n int, p *Pagination, page, limit int32) bool {
	if n == 0 || p == nil {
		// No pagination block: the endpoint returned everything at once.
		return true
	}
	if p.TotalPages != nil {
		return page >= *p.TotalPages
	}
	if p.Total != nil {
		return int64(page)*int64(limit) >= int64(*p.Total)
	}
	per := p.Limit
	if per <= 0 {
		per = limit
	}
	return int32(n) < per
}
