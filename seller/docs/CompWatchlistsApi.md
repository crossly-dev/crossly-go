# \CompWatchlistsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCompWatchlist**](CompWatchlistsApi.md#CreateCompWatchlist) | **Post** /v1/comp-watchlists | Create a sold-comp watchlist.
[**CreateCompWatchlistScrape**](CompWatchlistsApi.md#CreateCompWatchlistScrape) | **Post** /v1/comp-watchlists/{id}/scrape | Manually trigger a watchlist scrape.
[**DeleteCompWatchlist**](CompWatchlistsApi.md#DeleteCompWatchlist) | **Delete** /v1/comp-watchlists/{id} | Delete a sold-comp watchlist.
[**ListCompWatchlistRecent**](CompWatchlistsApi.md#ListCompWatchlistRecent) | **Get** /v1/comp-watchlists/{id}/recent | Recent external sold comps matching this watchlist.
[**ListCompWatchlists**](CompWatchlistsApi.md#ListCompWatchlists) | **Get** /v1/comp-watchlists | List the seller&#39;s sold-comp watchlists.



## CreateCompWatchlist

> ListCompWatchlistsItem CreateCompWatchlist(ctx).Execute()

Create a sold-comp watchlist.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompWatchlistsApi.CreateCompWatchlist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompWatchlistsApi.CreateCompWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCompWatchlist`: ListCompWatchlistsItem
    fmt.Fprintf(os.Stdout, "Response from `CompWatchlistsApi.CreateCompWatchlist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompWatchlistRequest struct via the builder pattern


### Return type

**ListCompWatchlistsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCompWatchlistScrape

> CreateCompWatchlistScrapeResponse CreateCompWatchlistScrape(ctx, id).Execute()

Manually trigger a watchlist scrape.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompWatchlistsApi.CreateCompWatchlistScrape(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompWatchlistsApi.CreateCompWatchlistScrape``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCompWatchlistScrape`: CreateCompWatchlistScrapeResponse
    fmt.Fprintf(os.Stdout, "Response from `CompWatchlistsApi.CreateCompWatchlistScrape`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCompWatchlistScrapeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateCompWatchlistScrapeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteCompWatchlist

> OkResponse DeleteCompWatchlist(ctx, id).Execute()

Delete a sold-comp watchlist.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompWatchlistsApi.DeleteCompWatchlist(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompWatchlistsApi.DeleteCompWatchlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCompWatchlist`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `CompWatchlistsApi.DeleteCompWatchlist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCompWatchlistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCompWatchlistRecent

> ListCompWatchlistRecentResponse ListCompWatchlistRecent(ctx, id).Execute()

Recent external sold comps matching this watchlist.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompWatchlistsApi.ListCompWatchlistRecent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompWatchlistsApi.ListCompWatchlistRecent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompWatchlistRecent`: ListCompWatchlistRecentResponse
    fmt.Fprintf(os.Stdout, "Response from `CompWatchlistsApi.ListCompWatchlistRecent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCompWatchlistRecentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListCompWatchlistRecentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCompWatchlists

> ListCompWatchlistsResponse ListCompWatchlists(ctx).Execute()

List the seller's sold-comp watchlists.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CompWatchlistsApi.ListCompWatchlists(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CompWatchlistsApi.ListCompWatchlists``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCompWatchlists`: ListCompWatchlistsResponse
    fmt.Fprintf(os.Stdout, "Response from `CompWatchlistsApi.ListCompWatchlists`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCompWatchlistsRequest struct via the builder pattern


### Return type

**ListCompWatchlistsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

