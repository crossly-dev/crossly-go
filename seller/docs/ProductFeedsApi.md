# \ProductFeedsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProductFeed**](ProductFeedsApi.md#CreateProductFeed) | **Post** /v1/product-feeds | Create a hosted product feed for a channel. Returns its secret URL.
[**CreateProductFeedRotate**](ProductFeedsApi.md#CreateProductFeedRotate) | **Post** /v1/product-feeds/{id}/rotate | Rotate the feed token. The old URL stops answering immediately; paste the new one into the channel.
[**DeleteProductFeed**](ProductFeedsApi.md#DeleteProductFeed) | **Delete** /v1/product-feeds/{id} | Delete a feed. Its URL stops answering immediately.
[**GetProductFeedReport**](ProductFeedsApi.md#GetProductFeedReport) | **Get** /v1/product-feeds/{id}/report | Validate a feed: per-item errors (held back) and warnings (sent, limited reach), plus sample rendered items.
[**ListProductFeeds**](ProductFeedsApi.md#ListProductFeeds) | **Get** /v1/product-feeds | List your hosted product feeds (Google Merchant Center, Meta catalogue, TikTok catalog) with their URLs and last-fetch stats.
[**UpdateProductFeed**](ProductFeedsApi.md#UpdateProductFeed) | **Patch** /v1/product-feeds/{id} | Edit a feed: name, default format, enabled, listing filters, output settings.



## CreateProductFeed

> ListProductFeedsItem CreateProductFeed(ctx).Execute()

Create a hosted product feed for a channel. Returns its secret URL.

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
    resp, r, err := apiClient.ProductFeedsApi.CreateProductFeed(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductFeedsApi.CreateProductFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProductFeed`: ListProductFeedsItem
    fmt.Fprintf(os.Stdout, "Response from `ProductFeedsApi.CreateProductFeed`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductFeedRequest struct via the builder pattern


### Return type

**ListProductFeedsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateProductFeedRotate

> ListProductFeedsItem CreateProductFeedRotate(ctx, id).Execute()

Rotate the feed token. The old URL stops answering immediately; paste the new one into the channel.

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
    resp, r, err := apiClient.ProductFeedsApi.CreateProductFeedRotate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductFeedsApi.CreateProductFeedRotate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProductFeedRotate`: ListProductFeedsItem
    fmt.Fprintf(os.Stdout, "Response from `ProductFeedsApi.CreateProductFeedRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductFeedRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListProductFeedsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteProductFeed

> OkResponse DeleteProductFeed(ctx, id).Execute()

Delete a feed. Its URL stops answering immediately.

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
    resp, r, err := apiClient.ProductFeedsApi.DeleteProductFeed(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductFeedsApi.DeleteProductFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteProductFeed`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductFeedsApi.DeleteProductFeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteProductFeedRequest struct via the builder pattern


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


## GetProductFeedReport

> GetProductFeedReportResponse GetProductFeedReport(ctx, id).Execute()

Validate a feed: per-item errors (held back) and warnings (sent, limited reach), plus sample rendered items.

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
    resp, r, err := apiClient.ProductFeedsApi.GetProductFeedReport(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductFeedsApi.GetProductFeedReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductFeedReport`: GetProductFeedReportResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductFeedsApi.GetProductFeedReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductFeedReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetProductFeedReportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListProductFeeds

> ListProductFeedsResponse ListProductFeeds(ctx).Execute()

List your hosted product feeds (Google Merchant Center, Meta catalogue, TikTok catalog) with their URLs and last-fetch stats.

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
    resp, r, err := apiClient.ProductFeedsApi.ListProductFeeds(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductFeedsApi.ListProductFeeds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListProductFeeds`: ListProductFeedsResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductFeedsApi.ListProductFeeds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListProductFeedsRequest struct via the builder pattern


### Return type

**ListProductFeedsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateProductFeed

> ListProductFeedsItem UpdateProductFeed(ctx, id).Execute()

Edit a feed: name, default format, enabled, listing filters, output settings.

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
    resp, r, err := apiClient.ProductFeedsApi.UpdateProductFeed(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductFeedsApi.UpdateProductFeed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProductFeed`: ListProductFeedsItem
    fmt.Fprintf(os.Stdout, "Response from `ProductFeedsApi.UpdateProductFeed`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductFeedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListProductFeedsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

