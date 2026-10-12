# \BuyerCatalogApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBuyerIdentify**](BuyerCatalogApi.md#CreateBuyerIdentify) | **Post** /v1/buyer/identify | Identify a held object and return a HUD-ready answer.
[**CreateBuyerLockon**](BuyerCatalogApi.md#CreateBuyerLockon) | **Post** /v1/buyer/lockons | Lock on to an object the buyer is holding.
[**CreateBuyerLockonConfirm**](BuyerCatalogApi.md#CreateBuyerLockonConfirm) | **Post** /v1/buyer/lockons/{id}/confirm | The buyer picked one of the candidates.
[**CreateBuyerLockonObserve**](BuyerCatalogApi.md#CreateBuyerLockonObserve) | **Post** /v1/buyer/lockons/{id}/observe | Add what this frame revealed, and get the current best answer.
[**CreateBuyerScan**](BuyerCatalogApi.md#CreateBuyerScan) | **Post** /v1/buyer/scan | Identify a physical item and find the cheapest place to buy it.
[**CreateBuyerScanSession**](BuyerCatalogApi.md#CreateBuyerScanSession) | **Post** /v1/buyer/scan/sessions | Open a Live Shop session.
[**CreateBuyerScanSessionEnd**](BuyerCatalogApi.md#CreateBuyerScanSessionEnd) | **Post** /v1/buyer/scan/sessions/{id}/end | Close a Live Shop session.
[**GetBuyerAnywhere**](BuyerCatalogApi.md#GetBuyerAnywhere) | **Get** /v1/buyer/anywhere | Cheapest source for an item — Crossly first, then other retailers.
[**GetBuyerCatalogFacet**](BuyerCatalogApi.md#GetBuyerCatalogFacet) | **Get** /v1/buyer/catalog/facets | Brands, categories and conditions that currently have stock.
[**GetBuyerCatalogFx**](BuyerCatalogApi.md#GetBuyerCatalogFx) | **Get** /v1/buyer/catalog/fx | Display-only exchange rates (units per 1 USD).
[**GetBuyerCatalogListing**](BuyerCatalogApi.md#GetBuyerCatalogListing) | **Get** /v1/buyer/catalog/listings/{slug} | One listing, in full.
[**GetBuyerCatalogListingAvailability**](BuyerCatalogApi.md#GetBuyerCatalogListingAvailability) | **Get** /v1/buyer/catalog/listings/{slug}/availability | Is it still buyable, and at what price.
[**GetBuyerCatalogListingShipping**](BuyerCatalogApi.md#GetBuyerCatalogListingShipping) | **Get** /v1/buyer/catalog/listings/{slug}/shipping | Shipping services and prices to a destination.
[**GetBuyerListingLandedCost**](BuyerCatalogApi.md#GetBuyerListingLandedCost) | **Get** /v1/buyer/listings/{slug}/landed-cost | Import duties, taxes and fees to a country (landed cost).
[**GetBuyerScanSession**](BuyerCatalogApi.md#GetBuyerScanSession) | **Get** /v1/buyer/scan/sessions/{id} | One trip and everything it found.
[**ListBuyerCatalogBrowse**](BuyerCatalogApi.md#ListBuyerCatalogBrowse) | **Get** /v1/buyer/catalog/browse | Browse the way the shop does — buying format, returns, size, color, sold items, Best Match.
[**ListBuyerCatalogListingSoldHistory**](BuyerCatalogApi.md#ListBuyerCatalogListingSoldHistory) | **Get** /v1/buyer/catalog/listings/{slug}/sold-history | What THIS listing has sold for, newest first (last 50), with the median.
[**ListBuyerCatalogSearch**](BuyerCatalogApi.md#ListBuyerCatalogSearch) | **Get** /v1/buyer/catalog/search | Search the Crossly catalogue.
[**ListBuyerCatalogSellers**](BuyerCatalogApi.md#ListBuyerCatalogSellers) | **Get** /v1/buyer/catalog/sellers | Find a seller by shop handle or shop name.
[**ListBuyerScanSessions**](BuyerCatalogApi.md#ListBuyerScanSessions) | **Get** /v1/buyer/scan/sessions | Your scanning trips, newest first.



## CreateBuyerIdentify

> CreateBuyerIdentifyResponse CreateBuyerIdentify(ctx).Execute()

Identify a held object and return a HUD-ready answer.



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
    resp, r, err := apiClient.BuyerCatalogApi.CreateBuyerIdentify(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.CreateBuyerIdentify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerIdentify`: CreateBuyerIdentifyResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.CreateBuyerIdentify`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerIdentifyRequest struct via the builder pattern


### Return type

**CreateBuyerIdentifyResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerLockon

> CreateBuyerLockonResponse CreateBuyerLockon(ctx).Execute()

Lock on to an object the buyer is holding.



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
    resp, r, err := apiClient.BuyerCatalogApi.CreateBuyerLockon(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.CreateBuyerLockon``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerLockon`: CreateBuyerLockonResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.CreateBuyerLockon`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerLockonRequest struct via the builder pattern


### Return type

**CreateBuyerLockonResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerLockonConfirm

> CreateBuyerLockonConfirmResponse CreateBuyerLockonConfirm(ctx, id).Execute()

The buyer picked one of the candidates.



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
    resp, r, err := apiClient.BuyerCatalogApi.CreateBuyerLockonConfirm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.CreateBuyerLockonConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerLockonConfirm`: CreateBuyerLockonConfirmResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.CreateBuyerLockonConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerLockonConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerLockonConfirmResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerLockonObserve

> CreateBuyerLockonConfirmResponse CreateBuyerLockonObserve(ctx, id).Execute()

Add what this frame revealed, and get the current best answer.



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
    resp, r, err := apiClient.BuyerCatalogApi.CreateBuyerLockonObserve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.CreateBuyerLockonObserve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerLockonObserve`: CreateBuyerLockonConfirmResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.CreateBuyerLockonObserve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerLockonObserveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerLockonConfirmResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerScan

> CreateBuyerScanResponse CreateBuyerScan(ctx).Execute()

Identify a physical item and find the cheapest place to buy it.



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
    resp, r, err := apiClient.BuyerCatalogApi.CreateBuyerScan(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.CreateBuyerScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerScan`: CreateBuyerScanResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.CreateBuyerScan`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerScanRequest struct via the builder pattern


### Return type

**CreateBuyerScanResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerScanSession

> CreateBuyerScanSessionResponse CreateBuyerScanSession(ctx).Execute()

Open a Live Shop session.



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
    resp, r, err := apiClient.BuyerCatalogApi.CreateBuyerScanSession(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.CreateBuyerScanSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerScanSession`: CreateBuyerScanSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.CreateBuyerScanSession`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerScanSessionRequest struct via the builder pattern


### Return type

**CreateBuyerScanSessionResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBuyerScanSessionEnd

> CreateBuyerScanSessionEndResponse CreateBuyerScanSessionEnd(ctx, id).Execute()

Close a Live Shop session.

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
    resp, r, err := apiClient.BuyerCatalogApi.CreateBuyerScanSessionEnd(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.CreateBuyerScanSessionEnd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerScanSessionEnd`: CreateBuyerScanSessionEndResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.CreateBuyerScanSessionEnd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerScanSessionEndRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBuyerScanSessionEndResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerAnywhere

> GetBuyerAnywhereResponse GetBuyerAnywhere(ctx).Execute()

Cheapest source for an item — Crossly first, then other retailers.



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
    resp, r, err := apiClient.BuyerCatalogApi.GetBuyerAnywhere(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.GetBuyerAnywhere``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerAnywhere`: GetBuyerAnywhereResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.GetBuyerAnywhere`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerAnywhereRequest struct via the builder pattern


### Return type

**GetBuyerAnywhereResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerCatalogFacet

> GetBuyerCatalogFacetResponse GetBuyerCatalogFacet(ctx).Execute()

Brands, categories and conditions that currently have stock.



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
    resp, r, err := apiClient.BuyerCatalogApi.GetBuyerCatalogFacet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.GetBuyerCatalogFacet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerCatalogFacet`: GetBuyerCatalogFacetResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.GetBuyerCatalogFacet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerCatalogFacetRequest struct via the builder pattern


### Return type

**GetBuyerCatalogFacetResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerCatalogFx

> GetBuyerCatalogFxResponse GetBuyerCatalogFx(ctx).Execute()

Display-only exchange rates (units per 1 USD).



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
    resp, r, err := apiClient.BuyerCatalogApi.GetBuyerCatalogFx(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.GetBuyerCatalogFx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerCatalogFx`: GetBuyerCatalogFxResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.GetBuyerCatalogFx`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerCatalogFxRequest struct via the builder pattern


### Return type

**GetBuyerCatalogFxResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerCatalogListing

> GetBuyerCatalogListingResponse GetBuyerCatalogListing(ctx, slug).Execute()

One listing, in full.

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
    slug := "slug_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerCatalogApi.GetBuyerCatalogListing(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.GetBuyerCatalogListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerCatalogListing`: GetBuyerCatalogListingResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.GetBuyerCatalogListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerCatalogListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerCatalogListingResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerCatalogListingAvailability

> GetBuyerCatalogListingAvailabilityResponse GetBuyerCatalogListingAvailability(ctx, slug).Execute()

Is it still buyable, and at what price.



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
    slug := "slug_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerCatalogApi.GetBuyerCatalogListingAvailability(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.GetBuyerCatalogListingAvailability``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerCatalogListingAvailability`: GetBuyerCatalogListingAvailabilityResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.GetBuyerCatalogListingAvailability`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerCatalogListingAvailabilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerCatalogListingAvailabilityResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerCatalogListingShipping

> GetBuyerCatalogListingShippingResponse GetBuyerCatalogListingShipping(ctx, slug).Zip(zip).Country(country).Quantity(quantity).Execute()

Shipping services and prices to a destination.



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
    slug := "slug_example" // string | 
    zip := "zip_example" // string |  (optional)
    country := "country_example" // string |  (optional)
    quantity := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerCatalogApi.GetBuyerCatalogListingShipping(context.Background(), slug).Zip(zip).Country(country).Quantity(quantity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.GetBuyerCatalogListingShipping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerCatalogListingShipping`: GetBuyerCatalogListingShippingResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.GetBuyerCatalogListingShipping`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerCatalogListingShippingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **zip** | **string** |  | 
 **country** | **string** |  | 
 **quantity** | **int32** |  | 

### Return type

**GetBuyerCatalogListingShippingResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerListingLandedCost

> GetBuyerListingLandedCostResponse GetBuyerListingLandedCost(ctx, slug).Country(country).PostalCode(postalCode).Quantity(quantity).Execute()

Import duties, taxes and fees to a country (landed cost).



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
    country := "country_example" // string | 
    slug := "slug_example" // string | 
    postalCode := "postalCode_example" // string |  (optional)
    quantity := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerCatalogApi.GetBuyerListingLandedCost(context.Background(), slug).Country(country).PostalCode(postalCode).Quantity(quantity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.GetBuyerListingLandedCost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerListingLandedCost`: GetBuyerListingLandedCostResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.GetBuyerListingLandedCost`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerListingLandedCostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **country** | **string** |  | 

 **postalCode** | **string** |  | 
 **quantity** | **int32** |  | 

### Return type

**GetBuyerListingLandedCostResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBuyerScanSession

> GetBuyerScanSessionResponse GetBuyerScanSession(ctx, id).Execute()

One trip and everything it found.



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
    resp, r, err := apiClient.BuyerCatalogApi.GetBuyerScanSession(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.GetBuyerScanSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerScanSession`: GetBuyerScanSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.GetBuyerScanSession`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerScanSessionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBuyerScanSessionResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerCatalogBrowse

> ListBuyerCatalogBrowseResponse ListBuyerCatalogBrowse(ctx).Q(q).Format(format).Returns(returns).Sold(sold).Size(size).Color(color).ShipsTo(shipsTo).Sort(sort).MinPriceCents(minPriceCents).MaxPriceCents(maxPriceCents).Page(page).Limit(limit).Execute()

Browse the way the shop does — buying format, returns, size, color, sold items, Best Match.



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
    q := "q_example" // string | Free text; every word must appear in the title or brand. (optional)
    format := "format_example" // string | auction, buy_it_now, or offers (accepts offers). (optional)
    returns := "returns_example" // string | true = only listings that accept returns. (optional)
    sold := "sold_example" // string | true = recent sales instead of live listings. (optional)
    size := "size_example" // string | Exact size, case-insensitive (M does not match XL). (optional)
    color := "color_example" // string | Exact colour, case-insensitive. (optional)
    shipsTo := "shipsTo_example" // string | ISO 3166-1 alpha-2 destination (e.g. GB): only listings whose seller ships there, decided as checkout decides it (the listing's shipping policy and its own restricted countries). An embargoed destination matches nothing. Omitted, or `any`, means no destination filter — the storefront grid defaults to the buyer's country, this endpoint does not. (optional)
    sort := "sort_example" // string | Default best_match with q, newest without. (optional)
    minPriceCents := int32(56) // int32 | Lowest price, in cents. (optional)
    maxPriceCents := int32(56) // int32 | Highest price, in cents. (optional)
    page := int32(56) // int32 | 1-indexed page. (optional) (default to 1)
    limit := int32(56) // int32 | Rows per page. (optional) (default to 24)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerCatalogApi.ListBuyerCatalogBrowse(context.Background()).Q(q).Format(format).Returns(returns).Sold(sold).Size(size).Color(color).ShipsTo(shipsTo).Sort(sort).MinPriceCents(minPriceCents).MaxPriceCents(maxPriceCents).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.ListBuyerCatalogBrowse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerCatalogBrowse`: ListBuyerCatalogBrowseResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.ListBuyerCatalogBrowse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerCatalogBrowseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | Free text; every word must appear in the title or brand. | 
 **format** | **string** | auction, buy_it_now, or offers (accepts offers). | 
 **returns** | **string** | true &#x3D; only listings that accept returns. | 
 **sold** | **string** | true &#x3D; recent sales instead of live listings. | 
 **size** | **string** | Exact size, case-insensitive (M does not match XL). | 
 **color** | **string** | Exact colour, case-insensitive. | 
 **shipsTo** | **string** | ISO 3166-1 alpha-2 destination (e.g. GB): only listings whose seller ships there, decided as checkout decides it (the listing&#39;s shipping policy and its own restricted countries). An embargoed destination matches nothing. Omitted, or &#x60;any&#x60;, means no destination filter — the storefront grid defaults to the buyer&#39;s country, this endpoint does not. | 
 **sort** | **string** | Default best_match with q, newest without. | 
 **minPriceCents** | **int32** | Lowest price, in cents. | 
 **maxPriceCents** | **int32** | Highest price, in cents. | 
 **page** | **int32** | 1-indexed page. | [default to 1]
 **limit** | **int32** | Rows per page. | [default to 24]

### Return type

**ListBuyerCatalogBrowseResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerCatalogListingSoldHistory

> ListBuyerCatalogListingSoldHistoryResponse ListBuyerCatalogListingSoldHistory(ctx, slug).Execute()

What THIS listing has sold for, newest first (last 50), with the median.



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
    slug := "slug_example" // string | The listing slug.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerCatalogApi.ListBuyerCatalogListingSoldHistory(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.ListBuyerCatalogListingSoldHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerCatalogListingSoldHistory`: ListBuyerCatalogListingSoldHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.ListBuyerCatalogListingSoldHistory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** | The listing slug. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerCatalogListingSoldHistoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListBuyerCatalogListingSoldHistoryResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerCatalogSearch

> ListBuyerCatalogSearchResponse ListBuyerCatalogSearch(ctx).Execute()

Search the Crossly catalogue.



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
    resp, r, err := apiClient.BuyerCatalogApi.ListBuyerCatalogSearch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.ListBuyerCatalogSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerCatalogSearch`: ListBuyerCatalogSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.ListBuyerCatalogSearch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerCatalogSearchRequest struct via the builder pattern


### Return type

**ListBuyerCatalogSearchResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerCatalogSellers

> ListBuyerCatalogSellersResponse ListBuyerCatalogSellers(ctx).Q(q).Limit(limit).Execute()

Find a seller by shop handle or shop name.



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
    q := "q_example" // string | At least two characters of a handle or shop name.
    limit := int32(56) // int32 | How many shops to return. (optional) (default to 12)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BuyerCatalogApi.ListBuyerCatalogSellers(context.Background()).Q(q).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.ListBuyerCatalogSellers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerCatalogSellers`: ListBuyerCatalogSellersResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.ListBuyerCatalogSellers`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerCatalogSellersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | At least two characters of a handle or shop name. | 
 **limit** | **int32** | How many shops to return. | [default to 12]

### Return type

**ListBuyerCatalogSellersResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBuyerScanSessions

> ListBuyerScanSessionsResponse ListBuyerScanSessions(ctx).Execute()

Your scanning trips, newest first.

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
    resp, r, err := apiClient.BuyerCatalogApi.ListBuyerScanSessions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCatalogApi.ListBuyerScanSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBuyerScanSessions`: ListBuyerScanSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.ListBuyerScanSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerScanSessionsRequest struct via the builder pattern


### Return type

**ListBuyerScanSessionsResponse**

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

