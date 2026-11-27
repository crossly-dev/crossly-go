# \SearchApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSearch**](SearchApi.md#GetSearch) | **Get** /v1/search | Search everything in your business at once: inventory, listings, orders, sales, customers, conversations, offers, shipments and returns. Exact identifiers (SKU, barcode, order number, tracking number, buyer email, marketplace listing id or a pasted listing URL) rank first, then prefixes, then fuzzy title / brand / buyer matches. Operators: sku:, order:, buyer:, tracking:, in:&lt;type&gt;, status:, platform:. Groups your token has no read scope for are left out.



## GetSearch

> GetSearchResponse GetSearch(ctx).Q(q).Types(types).Limit(limit).Execute()

Search everything in your business at once: inventory, listings, orders, sales, customers, conversations, offers, shipments and returns. Exact identifiers (SKU, barcode, order number, tracking number, buyer email, marketplace listing id or a pasted listing URL) rank first, then prefixes, then fuzzy title / brand / buyer matches. Operators: sku:, order:, buyer:, tracking:, in:<type>, status:, platform:. Groups your token has no read scope for are left out.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    q := "q_example" // string | What to search for (operators allowed). (optional)
    types := "types_example" // string | Comma-separated groups to search: inventory, listings, orders, sales, customers, conversations, offers, shipments, returns. (optional)
    limit := int32(56) // int32 | Hits per group. (optional) (default to 5)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SearchApi.GetSearch(context.Background()).Q(q).Types(types).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SearchApi.GetSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSearch`: GetSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `SearchApi.GetSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | What to search for (operators allowed). | 
 **types** | **string** | Comma-separated groups to search: inventory, listings, orders, sales, customers, conversations, offers, shipments, returns. | 
 **limit** | **int32** | Hits per group. | [default to 5]

### Return type

**GetSearchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

