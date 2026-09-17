# \NetworkApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNetworkPool**](NetworkApi.md#CreateNetworkPool) | **Post** /v1/network/pool | Join the Crossly Network reciprocal engagement pool.
[**DeleteNetworkPool**](NetworkApi.md#DeleteNetworkPool) | **Delete** /v1/network/pool | Leave the Crossly Network pool.
[**GetNetworkPool**](NetworkApi.md#GetNetworkPool) | **Get** /v1/network/pool | The seller&#39;s Crossly Network pool membership row.
[**GetNetworkPoolSize**](NetworkApi.md#GetNetworkPoolSize) | **Get** /v1/network/pool/size | Total members in the Crossly Network pool.
[**ListNetworkPoolLog**](NetworkApi.md#ListNetworkPoolLog) | **Get** /v1/network/pool/log | Recent engagement history — both sent and received.
[**UpdateNetworkPool**](NetworkApi.md#UpdateNetworkPool) | **Patch** /v1/network/pool | Update per-action toggles + platforms on pool membership.



## CreateNetworkPool

> CreateNetworkPoolResponse CreateNetworkPool(ctx).Execute()

Join the Crossly Network reciprocal engagement pool.

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
    resp, r, err := apiClient.NetworkApi.CreateNetworkPool(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.CreateNetworkPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNetworkPool`: CreateNetworkPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.CreateNetworkPool`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNetworkPoolRequest struct via the builder pattern


### Return type

[**CreateNetworkPoolResponse**](CreateNetworkPoolResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNetworkPool

> DeleteNetworkPoolResponse DeleteNetworkPool(ctx).Execute()

Leave the Crossly Network pool.

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
    resp, r, err := apiClient.NetworkApi.DeleteNetworkPool(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.DeleteNetworkPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteNetworkPool`: DeleteNetworkPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.DeleteNetworkPool`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNetworkPoolRequest struct via the builder pattern


### Return type

[**DeleteNetworkPoolResponse**](DeleteNetworkPoolResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPool

> GetNetworkPoolResponse GetNetworkPool(ctx).Execute()

The seller's Crossly Network pool membership row.

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
    resp, r, err := apiClient.NetworkApi.GetNetworkPool(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.GetNetworkPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPool`: GetNetworkPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.GetNetworkPool`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoolRequest struct via the builder pattern


### Return type

[**GetNetworkPoolResponse**](GetNetworkPoolResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNetworkPoolSize

> GetNetworkPoolSizeResponse GetNetworkPoolSize(ctx).Execute()

Total members in the Crossly Network pool.

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
    resp, r, err := apiClient.NetworkApi.GetNetworkPoolSize(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.GetNetworkPoolSize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNetworkPoolSize`: GetNetworkPoolSizeResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.GetNetworkPoolSize`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworkPoolSizeRequest struct via the builder pattern


### Return type

[**GetNetworkPoolSizeResponse**](GetNetworkPoolSizeResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNetworkPoolLog

> V1List ListNetworkPoolLog(ctx).Limit(limit).Execute()

Recent engagement history — both sent and received.

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
    limit := int32(56) // int32 |  (optional) (default to 100)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NetworkApi.ListNetworkPoolLog(context.Background()).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.ListNetworkPoolLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNetworkPoolLog`: V1List
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.ListNetworkPoolLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListNetworkPoolLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** |  | [default to 100]

### Return type

[**V1List**](V1List.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNetworkPool

> UpdateNetworkPoolResponse UpdateNetworkPool(ctx).Execute()

Update per-action toggles + platforms on pool membership.

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
    resp, r, err := apiClient.NetworkApi.UpdateNetworkPool(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NetworkApi.UpdateNetworkPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNetworkPool`: UpdateNetworkPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `NetworkApi.UpdateNetworkPool`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNetworkPoolRequest struct via the builder pattern


### Return type

[**UpdateNetworkPoolResponse**](UpdateNetworkPoolResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

