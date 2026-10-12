# \ScanApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListScanCaptures**](ScanApi.md#ListScanCaptures) | **Get** /v1/scan/captures | Your capture sessions and what they produced.
[**ListScanModels**](ScanApi.md#ListScanModels) | **Get** /v1/scan/models | Which items already have a 3D model.



## ListScanCaptures

> ListScanCapturesResponse ListScanCaptures(ctx).Status(status).Limit(limit).Execute()

Your capture sessions and what they produced.



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
    status := "status_example" // string | Filter to one status. (optional)
    limit := int32(56) // int32 |  (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScanApi.ListScanCaptures(context.Background()).Status(status).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanApi.ListScanCaptures``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScanCaptures`: ListScanCapturesResponse
    fmt.Fprintf(os.Stdout, "Response from `ScanApi.ListScanCaptures`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListScanCapturesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter to one status. | 
 **limit** | **int32** |  | [default to 25]

### Return type

**ListScanCapturesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListScanModels

> ListScanModelsResponse ListScanModels(ctx).VariantIds(variantIds).InventoryItemIds(inventoryItemIds).Execute()

Which items already have a 3D model.



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
    variantIds := "variantIds_example" // string | Comma-separated catalogue variant ids, max 200. (optional)
    inventoryItemIds := "inventoryItemIds_example" // string | Comma-separated inventory item ids, max 200. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ScanApi.ListScanModels(context.Background()).VariantIds(variantIds).InventoryItemIds(inventoryItemIds).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ScanApi.ListScanModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListScanModels`: ListScanModelsResponse
    fmt.Fprintf(os.Stdout, "Response from `ScanApi.ListScanModels`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListScanModelsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **variantIds** | **string** | Comma-separated catalogue variant ids, max 200. | 
 **inventoryItemIds** | **string** | Comma-separated inventory item ids, max 200. | 

### Return type

**ListScanModelsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

