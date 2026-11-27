# \ChannelInventoryApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteChannelInventory**](ChannelInventoryApi.md#DeleteChannelInventory) | **Delete** /v1/channel-inventory/{platform}/{slot} | Reset a channel&#39;s inventory controls to the default (no cap, no buffer, sync on, end on sell-out).
[**GetChannelInventory**](ChannelInventoryApi.md#GetChannelInventory) | **Get** /v1/channel-inventory/{platform}/{slot} | Get one channel&#39;s inventory controls.
[**ListChannelInventory**](ChannelInventoryApi.md#ListChannelInventory) | **Get** /v1/channel-inventory | List every sales channel with its inventory controls (cap, buffer, sync, out-of-stock action, exclusions).
[**UpdateChannelInventory**](ChannelInventoryApi.md#UpdateChannelInventory) | **Put** /v1/channel-inventory/{platform}/{slot} | Set a channel&#39;s quantity cap, buffer, sync toggle, out-of-stock action and exclusion lists (merge write).



## DeleteChannelInventory

> ListChannelInventoryItem DeleteChannelInventory(ctx, platform, slot).Execute()

Reset a channel's inventory controls to the default (no cap, no buffer, sync on, end on sell-out).

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
    slot := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelInventoryApi.DeleteChannelInventory(context.Background(), platform, slot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelInventoryApi.DeleteChannelInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteChannelInventory`: ListChannelInventoryItem
    fmt.Fprintf(os.Stdout, "Response from `ChannelInventoryApi.DeleteChannelInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 
**slot** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteChannelInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**ListChannelInventoryItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetChannelInventory

> ListChannelInventoryItem GetChannelInventory(ctx, platform, slot).Execute()

Get one channel's inventory controls.

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
    slot := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelInventoryApi.GetChannelInventory(context.Background(), platform, slot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelInventoryApi.GetChannelInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetChannelInventory`: ListChannelInventoryItem
    fmt.Fprintf(os.Stdout, "Response from `ChannelInventoryApi.GetChannelInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 
**slot** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetChannelInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**ListChannelInventoryItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListChannelInventory

> ListChannelInventoryResponse ListChannelInventory(ctx).Execute()

List every sales channel with its inventory controls (cap, buffer, sync, out-of-stock action, exclusions).

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
    resp, r, err := apiClient.ChannelInventoryApi.ListChannelInventory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelInventoryApi.ListChannelInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListChannelInventory`: ListChannelInventoryResponse
    fmt.Fprintf(os.Stdout, "Response from `ChannelInventoryApi.ListChannelInventory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListChannelInventoryRequest struct via the builder pattern


### Return type

**ListChannelInventoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateChannelInventory

> ListChannelInventoryItem UpdateChannelInventory(ctx, platform, slot).Execute()

Set a channel's quantity cap, buffer, sync toggle, out-of-stock action and exclusion lists (merge write).

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
    slot := int32(56) // int32 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ChannelInventoryApi.UpdateChannelInventory(context.Background(), platform, slot).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ChannelInventoryApi.UpdateChannelInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateChannelInventory`: ListChannelInventoryItem
    fmt.Fprintf(os.Stdout, "Response from `ChannelInventoryApi.UpdateChannelInventory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 
**slot** | **int32** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateChannelInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**ListChannelInventoryItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

