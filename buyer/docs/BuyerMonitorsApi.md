# \BuyerMonitorsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBuyerMonitor**](BuyerMonitorsApi.md#CreateBuyerMonitor) | **Post** /v1/buyer/monitors | Watch a search, and be told when it matches.
[**DeleteBuyerMonitor**](BuyerMonitorsApi.md#DeleteBuyerMonitor) | **Delete** /v1/buyer/monitors/{id} | Delete a monitor.
[**ListBuyerMonitorMatches**](BuyerMonitorsApi.md#ListBuyerMonitorMatches) | **Get** /v1/buyer/monitors/{id}/matches | What this monitor has matched.
[**ListBuyerMonitors**](BuyerMonitorsApi.md#ListBuyerMonitors) | **Get** /v1/buyer/monitors | Your monitors.
[**UpdateBuyerMonitor**](BuyerMonitorsApi.md#UpdateBuyerMonitor) | **Patch** /v1/buyer/monitors/{id} | Pause, resume or rename a monitor.



## CreateBuyerMonitor

> ListBuyerMonitorsItem CreateBuyerMonitor(ctx).Execute()

Watch a search, and be told when it matches.



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
    resp, r, err := apiClient.BuyerMonitorsApi.CreateBuyerMonitor(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerMonitorsApi.CreateBuyerMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerMonitor`: ListBuyerMonitorsItem
    fmt.Fprintf(os.Stdout, "Response from `BuyerMonitorsApi.CreateBuyerMonitor`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerMonitorRequest struct via the builder pattern


### Return type

**ListBuyerMonitorsItem**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBuyerMonitor

> DeleteBuyerMonitor(ctx, id).Execute()

Delete a monitor.

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
    resp, r, err := apiClient.BuyerMonitorsApi.DeleteBuyerMonitor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerMonitorsApi.DeleteBuyerMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBuyerMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerMonitorMatches

> ListBuyerMonitorMatchesResponse ListBuyerMonitorMatches(ctx, id).Execute()

What this monitor has matched.



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
    resp, r, err := apiClient.BuyerMonitorsApi.ListBuyerMonitorMatches(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerMonitorsApi.ListBuyerMonitorMatches``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerMonitorMatches`: ListBuyerMonitorMatchesResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerMonitorsApi.ListBuyerMonitorMatches`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerMonitorMatchesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListBuyerMonitorMatchesResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerMonitors

> ListBuyerMonitorsResponse ListBuyerMonitors(ctx).Execute()

Your monitors.

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
    resp, r, err := apiClient.BuyerMonitorsApi.ListBuyerMonitors(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerMonitorsApi.ListBuyerMonitors``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerMonitors`: ListBuyerMonitorsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerMonitorsApi.ListBuyerMonitors`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerMonitorsRequest struct via the builder pattern


### Return type

**ListBuyerMonitorsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBuyerMonitor

> ListBuyerMonitorsItem UpdateBuyerMonitor(ctx, id).Execute()

Pause, resume or rename a monitor.

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
    resp, r, err := apiClient.BuyerMonitorsApi.UpdateBuyerMonitor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerMonitorsApi.UpdateBuyerMonitor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBuyerMonitor`: ListBuyerMonitorsItem
    fmt.Fprintf(os.Stdout, "Response from `BuyerMonitorsApi.UpdateBuyerMonitor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerMonitorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListBuyerMonitorsItem**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

