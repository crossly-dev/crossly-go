# \NFTMarketApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateNftMarketExternalResale**](NFTMarketApi.md#CreateNftMarketExternalResale) | **Post** /v1/nft-market/external-resale | Ops: record a Crossly NFT resold on ANY marketplace — maps mint → design and funds the pot. Admin PAT only.
[**CreateNftMarketListing**](NFTMarketApi.md#CreateNftMarketListing) | **Post** /v1/nft-market/listings | List an NFT you own for resale on the Crossly marketplace.
[**CreateNftMarketListingCancel**](NFTMarketApi.md#CreateNftMarketListingCancel) | **Post** /v1/nft-market/listings/{id}/cancel | Cancel one of your open NFT listings.
[**CreateNftMarketListingConfirmSale**](NFTMarketApi.md#CreateNftMarketListingConfirmSale) | **Post** /v1/nft-market/listings/{id}/confirm-sale | Ops: record a finalized on-Crossly NFT sale (verifies the royalty landed, funds the pot). Admin PAT only.
[**GetNftMarketListing**](NFTMarketApi.md#GetNftMarketListing) | **Get** /v1/nft-market/listings | Browse open on-platform NFT listings (Crossly cosmetics for resale in CBX/SOL).



## CreateNftMarketExternalResale

> CreateNftMarketExternalResaleResponse CreateNftMarketExternalResale(ctx).Execute()

Ops: record a Crossly NFT resold on ANY marketplace — maps mint → design and funds the pot. Admin PAT only.

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
    resp, r, err := apiClient.NFTMarketApi.CreateNftMarketExternalResale(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTMarketApi.CreateNftMarketExternalResale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNftMarketExternalResale`: CreateNftMarketExternalResaleResponse
    fmt.Fprintf(os.Stdout, "Response from `NFTMarketApi.CreateNftMarketExternalResale`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNftMarketExternalResaleRequest struct via the builder pattern


### Return type

**CreateNftMarketExternalResaleResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateNftMarketListing

> CreateNftMarketListingResponse CreateNftMarketListing(ctx).Execute()

List an NFT you own for resale on the Crossly marketplace.

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
    resp, r, err := apiClient.NFTMarketApi.CreateNftMarketListing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTMarketApi.CreateNftMarketListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNftMarketListing`: CreateNftMarketListingResponse
    fmt.Fprintf(os.Stdout, "Response from `NFTMarketApi.CreateNftMarketListing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNftMarketListingRequest struct via the builder pattern


### Return type

**CreateNftMarketListingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateNftMarketListingCancel

> OkResponse CreateNftMarketListingCancel(ctx, id).Execute()

Cancel one of your open NFT listings.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFTMarketApi.CreateNftMarketListingCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTMarketApi.CreateNftMarketListingCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNftMarketListingCancel`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `NFTMarketApi.CreateNftMarketListingCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNftMarketListingCancelRequest struct via the builder pattern


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


## CreateNftMarketListingConfirmSale

> OkResponse CreateNftMarketListingConfirmSale(ctx, id).Execute()

Ops: record a finalized on-Crossly NFT sale (verifies the royalty landed, funds the pot). Admin PAT only.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.NFTMarketApi.CreateNftMarketListingConfirmSale(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTMarketApi.CreateNftMarketListingConfirmSale``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateNftMarketListingConfirmSale`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `NFTMarketApi.CreateNftMarketListingConfirmSale`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateNftMarketListingConfirmSaleRequest struct via the builder pattern


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


## GetNftMarketListing

> GetNftMarketListingResponse GetNftMarketListing(ctx).Execute()

Browse open on-platform NFT listings (Crossly cosmetics for resale in CBX/SOL).

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
    resp, r, err := apiClient.NFTMarketApi.GetNftMarketListing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NFTMarketApi.GetNftMarketListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNftMarketListing`: GetNftMarketListingResponse
    fmt.Fprintf(os.Stdout, "Response from `NFTMarketApi.GetNftMarketListing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNftMarketListingRequest struct via the builder pattern


### Return type

**GetNftMarketListingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

