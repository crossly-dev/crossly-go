# \OffersApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateOfferRespond**](OffersApi.md#CreateOfferRespond) | **Post** /v1/offers/{id}/respond | Accept, decline, or counter a buyer offer on a Crossly marketplace listing.
[**GetOffer**](OffersApi.md#GetOffer) | **Get** /v1/offers | List buyer offers on your Crossly marketplace listings, including bundles.



## CreateOfferRespond

> CreateOfferRespondResponse CreateOfferRespond(ctx, id).InlineObject2(inlineObject2).Execute()

Accept, decline, or counter a buyer offer on a Crossly marketplace listing.

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
    id := "id_example" // string | Offer UUID.
    inlineObject2 := *openapiclient.NewInlineObject2("Action_example") // InlineObject2 | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OffersApi.CreateOfferRespond(context.Background(), id).InlineObject2(inlineObject2).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OffersApi.CreateOfferRespond``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOfferRespond`: CreateOfferRespondResponse
    fmt.Fprintf(os.Stdout, "Response from `OffersApi.CreateOfferRespond`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Offer UUID. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOfferRespondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **inlineObject2** | [**InlineObject2**](InlineObject2.md) |  | 

### Return type

[**CreateOfferRespondResponse**](CreateOfferRespondResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOffer

> GetOfferResponse GetOffer(ctx).Status(status).Limit(limit).Execute()

List buyer offers on your Crossly marketplace listings, including bundles.

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
    status := "status_example" // string | Filter to one status. Omit for all. (optional)
    limit := int32(56) // int32 |  (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OffersApi.GetOffer(context.Background()).Status(status).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OffersApi.GetOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOffer`: GetOfferResponse
    fmt.Fprintf(os.Stdout, "Response from `OffersApi.GetOffer`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Filter to one status. Omit for all. | 
 **limit** | **int32** |  | [default to 50]

### Return type

[**GetOfferResponse**](GetOfferResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

