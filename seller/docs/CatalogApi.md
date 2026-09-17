# \CatalogApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCatalogLookup**](CatalogApi.md#GetCatalogLookup) | **Get** /v1/catalog/lookup | Live Crossly offers for a product identifier (barcode, style code, LEGO set…).



## GetCatalogLookup

> GetCatalogLookupResponse GetCatalogLookup(ctx).Namespace(namespace).Value(value).Execute()

Live Crossly offers for a product identifier (barcode, style code, LEGO set…).



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
    namespace := "namespace_example" // string | 
    value := "value_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CatalogApi.GetCatalogLookup(context.Background()).Namespace(namespace).Value(value).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CatalogApi.GetCatalogLookup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCatalogLookup`: GetCatalogLookupResponse
    fmt.Fprintf(os.Stdout, "Response from `CatalogApi.GetCatalogLookup`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCatalogLookupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **namespace** | **string** |  | 
 **value** | **string** |  | 

### Return type

[**GetCatalogLookupResponse**](GetCatalogLookupResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

