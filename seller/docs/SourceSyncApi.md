# \SourceSyncApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSourceSyncSync**](SourceSyncApi.md#CreateSourceSyncSync) | **Post** /v1/source-sync/{platform}/sync | Reconcile a store now: re-read its whole catalogue and apply every change (runs in the background; 202).
[**GetSourceSyncLocation**](SourceSyncApi.md#GetSourceSyncLocation) | **Get** /v1/source-sync/{platform}/locations | POS only (lightspeed_r, lightspeed_x, square): the shops/outlets/locations on the account and which one Crossly syncs stock and sales with.
[**ListSourceSync**](SourceSyncApi.md#ListSourceSync) | **Get** /v1/source-sync | List source-capable stores with their continuous-sync settings and status.
[**ListSourceSyncEvents**](SourceSyncApi.md#ListSourceSyncEvents) | **Get** /v1/source-sync/events | Recent source-sync activity: each webhook delivery and reconcile change, with its outcome.
[**UpdateSourceSync**](SourceSyncApi.md#UpdateSourceSync) | **Patch** /v1/source-sync/{platform} | Update a store&#39;s source-sync settings: on/off, source-owned fields, deletion handling, auto-export channels.
[**UpdateSourceSyncLocation**](SourceSyncApi.md#UpdateSourceSyncLocation) | **Put** /v1/source-sync/{platform}/location | POS only: choose the shop (R-Series) / outlet (X-Series) / Square location whose stock and sales Crossly syncs. Must be one GET .../locations lists.



## CreateSourceSyncSync

> CreateSourceSyncSyncResponse CreateSourceSyncSync(ctx, platform).Execute()

Reconcile a store now: re-read its whole catalogue and apply every change (runs in the background; 202).

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
    platform := "platform_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceSyncApi.CreateSourceSyncSync(context.Background(), platform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceSyncApi.CreateSourceSyncSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSourceSyncSync`: CreateSourceSyncSyncResponse
    fmt.Fprintf(os.Stdout, "Response from `SourceSyncApi.CreateSourceSyncSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourceSyncSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateSourceSyncSyncResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetSourceSyncLocation

> GetSourceSyncLocationResponse GetSourceSyncLocation(ctx, platform).Execute()

POS only (lightspeed_r, lightspeed_x, square): the shops/outlets/locations on the account and which one Crossly syncs stock and sales with.

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
    platform := "platform_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceSyncApi.GetSourceSyncLocation(context.Background(), platform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceSyncApi.GetSourceSyncLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourceSyncLocation`: GetSourceSyncLocationResponse
    fmt.Fprintf(os.Stdout, "Response from `SourceSyncApi.GetSourceSyncLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourceSyncLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetSourceSyncLocationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListSourceSync

> ListSourceSyncResponse ListSourceSync(ctx).Execute()

List source-capable stores with their continuous-sync settings and status.

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
    resp, r, err := apiClient.SourceSyncApi.ListSourceSync(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceSyncApi.ListSourceSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourceSync`: ListSourceSyncResponse
    fmt.Fprintf(os.Stdout, "Response from `SourceSyncApi.ListSourceSync`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSourceSyncRequest struct via the builder pattern


### Return type

**ListSourceSyncResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListSourceSyncEvents

> ListSourceSyncEventsResponse ListSourceSyncEvents(ctx).Platform(platform).Limit(limit).Execute()

Recent source-sync activity: each webhook delivery and reconcile change, with its outcome.

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
    platform := "platform_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceSyncApi.ListSourceSyncEvents(context.Background()).Platform(platform).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceSyncApi.ListSourceSyncEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourceSyncEvents`: ListSourceSyncEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `SourceSyncApi.ListSourceSyncEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSourceSyncEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platform** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

**ListSourceSyncEventsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateSourceSync

> ListSourceSyncItem UpdateSourceSync(ctx, platform).Execute()

Update a store's source-sync settings: on/off, source-owned fields, deletion handling, auto-export channels.

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
    platform := "platform_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceSyncApi.UpdateSourceSync(context.Background(), platform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceSyncApi.UpdateSourceSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSourceSync`: ListSourceSyncItem
    fmt.Fprintf(os.Stdout, "Response from `SourceSyncApi.UpdateSourceSync`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListSourceSyncItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateSourceSyncLocation

> GetSourceSyncLocationResponse UpdateSourceSyncLocation(ctx, platform).UpdateSourceSyncLocationRequest(updateSourceSyncLocationRequest).Execute()

POS only: choose the shop (R-Series) / outlet (X-Series) / Square location whose stock and sales Crossly syncs. Must be one GET .../locations lists.

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
    platform := "platform_example" // string | 
    updateSourceSyncLocationRequest := *openapiclient.NewUpdateSourceSyncLocationRequest("LocationId_example") // UpdateSourceSyncLocationRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourceSyncApi.UpdateSourceSyncLocation(context.Background(), platform).UpdateSourceSyncLocationRequest(updateSourceSyncLocationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourceSyncApi.UpdateSourceSyncLocation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSourceSyncLocation`: GetSourceSyncLocationResponse
    fmt.Fprintf(os.Stdout, "Response from `SourceSyncApi.UpdateSourceSyncLocation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSourceSyncLocationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateSourceSyncLocationRequest** | **UpdateSourceSyncLocationRequest** |  | 

### Return type

**GetSourceSyncLocationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

