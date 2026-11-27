# \BulkBatchesApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBulkBatch**](BulkBatchesApi.md#CreateBulkBatch) | **Post** /v1/bulk-batches | Open an empty capture batch for a preset.
[**CreateBulkBatchIdentify**](BulkBatchesApi.md#CreateBulkBatchIdentify) | **Post** /v1/bulk-batches/{id}/identify | Identify items: match to Crossly Market, Magic List, or auto (by preset). Runs in the background.
[**CreateBulkBatchItem**](BulkBatchesApi.md#CreateBulkBatchItem) | **Post** /v1/bulk-batches/{id}/items | Append one item (every photo of it, primary first).
[**CreateBulkBatchItemBulk**](BulkBatchesApi.md#CreateBulkBatchItemBulk) | **Post** /v1/bulk-batches/{id}/items/bulk | Set price/quantity/condition/status on many items.
[**CreateBulkBatchPublish**](BulkBatchesApi.md#CreateBulkBatchPublish) | **Post** /v1/bulk-batches/{id}/publish | Publish reviewed items: drafts only (default) or live on destinations; adds quantity to an existing matching listing.
[**DeleteBulkBatch**](BulkBatchesApi.md#DeleteBulkBatch) | **Delete** /v1/bulk-batches/{id} | Discard a batch (listings it produced stay).
[**DeleteBulkBatchItem**](BulkBatchesApi.md#DeleteBulkBatchItem) | **Delete** /v1/bulk-batches/items/{id} | Remove one unpublished item.
[**GetBulkBatch**](BulkBatchesApi.md#GetBulkBatch) | **Get** /v1/bulk-batches/{id} | Get one bulk batch with its items (capture order) and matched catalog rows.
[**ListBulkBatches**](BulkBatchesApi.md#ListBulkBatches) | **Get** /v1/bulk-batches | List recent bulk list batches, newest first.
[**UpdateBulkBatch**](BulkBatchesApi.md#UpdateBulkBatch) | **Patch** /v1/bulk-batches/{id} | Change batch settings, or set stage \&quot;identify\&quot; when capture is done.
[**UpdateBulkBatchItem**](BulkBatchesApi.md#UpdateBulkBatchItem) | **Patch** /v1/bulk-batches/items/{id} | Correct one item (match, title, price, photos…).



## CreateBulkBatch

> ListBulkBatchesItem CreateBulkBatch(ctx).Execute()

Open an empty capture batch for a preset.

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
    resp, r, err := apiClient.BulkBatchesApi.CreateBulkBatch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkBatchesApi.CreateBulkBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBulkBatch`: ListBulkBatchesItem
    fmt.Fprintf(os.Stdout, "Response from `BulkBatchesApi.CreateBulkBatch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulkBatchRequest struct via the builder pattern


### Return type

**ListBulkBatchesItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBulkBatchIdentify

> CreateBulkBatchIdentifyResponse CreateBulkBatchIdentify(ctx, id).Execute()

Identify items: match to Crossly Market, Magic List, or auto (by preset). Runs in the background.

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
    resp, r, err := apiClient.BulkBatchesApi.CreateBulkBatchIdentify(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkBatchesApi.CreateBulkBatchIdentify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBulkBatchIdentify`: CreateBulkBatchIdentifyResponse
    fmt.Fprintf(os.Stdout, "Response from `BulkBatchesApi.CreateBulkBatchIdentify`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulkBatchIdentifyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBulkBatchIdentifyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBulkBatchItem

> CreateBulkBatchItemResponse CreateBulkBatchItem(ctx, id).Execute()

Append one item (every photo of it, primary first).

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
    resp, r, err := apiClient.BulkBatchesApi.CreateBulkBatchItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkBatchesApi.CreateBulkBatchItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBulkBatchItem`: CreateBulkBatchItemResponse
    fmt.Fprintf(os.Stdout, "Response from `BulkBatchesApi.CreateBulkBatchItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulkBatchItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBulkBatchItemResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBulkBatchItemBulk

> UpdatedResponse CreateBulkBatchItemBulk(ctx, id).Execute()

Set price/quantity/condition/status on many items.

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
    resp, r, err := apiClient.BulkBatchesApi.CreateBulkBatchItemBulk(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkBatchesApi.CreateBulkBatchItemBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBulkBatchItemBulk`: UpdatedResponse
    fmt.Fprintf(os.Stdout, "Response from `BulkBatchesApi.CreateBulkBatchItemBulk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulkBatchItemBulkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdatedResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBulkBatchPublish

> CreateBulkBatchPublishResponse CreateBulkBatchPublish(ctx, id).Execute()

Publish reviewed items: drafts only (default) or live on destinations; adds quantity to an existing matching listing.

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
    resp, r, err := apiClient.BulkBatchesApi.CreateBulkBatchPublish(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkBatchesApi.CreateBulkBatchPublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBulkBatchPublish`: CreateBulkBatchPublishResponse
    fmt.Fprintf(os.Stdout, "Response from `BulkBatchesApi.CreateBulkBatchPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBulkBatchPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBulkBatchPublishResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBulkBatch

> OkResponse DeleteBulkBatch(ctx, id).Execute()

Discard a batch (listings it produced stay).

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
    resp, r, err := apiClient.BulkBatchesApi.DeleteBulkBatch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkBatchesApi.DeleteBulkBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBulkBatch`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `BulkBatchesApi.DeleteBulkBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBulkBatchRequest struct via the builder pattern


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


## DeleteBulkBatchItem

> OkResponse DeleteBulkBatchItem(ctx, id).Execute()

Remove one unpublished item.

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
    resp, r, err := apiClient.BulkBatchesApi.DeleteBulkBatchItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkBatchesApi.DeleteBulkBatchItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBulkBatchItem`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `BulkBatchesApi.DeleteBulkBatchItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBulkBatchItemRequest struct via the builder pattern


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


## GetBulkBatch

> GetBulkBatchResponse GetBulkBatch(ctx, id).Execute()

Get one bulk batch with its items (capture order) and matched catalog rows.

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
    resp, r, err := apiClient.BulkBatchesApi.GetBulkBatch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkBatchesApi.GetBulkBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBulkBatch`: GetBulkBatchResponse
    fmt.Fprintf(os.Stdout, "Response from `BulkBatchesApi.GetBulkBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBulkBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBulkBatchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBulkBatches

> ListBulkBatchesResponse ListBulkBatches(ctx).Limit(limit).Execute()

List recent bulk list batches, newest first.

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
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BulkBatchesApi.ListBulkBatches(context.Background()).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkBatchesApi.ListBulkBatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBulkBatches`: ListBulkBatchesResponse
    fmt.Fprintf(os.Stdout, "Response from `BulkBatchesApi.ListBulkBatches`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBulkBatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | 

### Return type

**ListBulkBatchesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBulkBatch

> ListBulkBatchesItem UpdateBulkBatch(ctx, id).Execute()

Change batch settings, or set stage \"identify\" when capture is done.

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
    resp, r, err := apiClient.BulkBatchesApi.UpdateBulkBatch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkBatchesApi.UpdateBulkBatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBulkBatch`: ListBulkBatchesItem
    fmt.Fprintf(os.Stdout, "Response from `BulkBatchesApi.UpdateBulkBatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBulkBatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListBulkBatchesItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBulkBatchItem

> CreateBulkBatchItemResponse UpdateBulkBatchItem(ctx, id).Execute()

Correct one item (match, title, price, photos…).

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
    resp, r, err := apiClient.BulkBatchesApi.UpdateBulkBatchItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BulkBatchesApi.UpdateBulkBatchItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBulkBatchItem`: CreateBulkBatchItemResponse
    fmt.Fprintf(os.Stdout, "Response from `BulkBatchesApi.UpdateBulkBatchItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBulkBatchItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBulkBatchItemResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

