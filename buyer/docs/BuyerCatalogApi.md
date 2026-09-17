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
[**GetBuyerCatalogListing**](BuyerCatalogApi.md#GetBuyerCatalogListing) | **Get** /v1/buyer/catalog/listings/{slug} | One listing, in full.
[**GetBuyerCatalogListingAvailability**](BuyerCatalogApi.md#GetBuyerCatalogListingAvailability) | **Get** /v1/buyer/catalog/listings/{slug}/availability | Is it still buyable, and at what price.
[**GetBuyerScanSession**](BuyerCatalogApi.md#GetBuyerScanSession) | **Get** /v1/buyer/scan/sessions/{id} | One trip and everything it found.
[**ListBuyerCatalogSearch**](BuyerCatalogApi.md#ListBuyerCatalogSearch) | **Get** /v1/buyer/catalog/search | Search the Crossly catalogue.
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

[**CreateBuyerIdentifyResponse**](CreateBuyerIdentifyResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
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

[**CreateBuyerLockonResponse**](CreateBuyerLockonResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
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

[**CreateBuyerLockonConfirmResponse**](CreateBuyerLockonConfirmResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateBuyerLockonObserve

> CreateBuyerLockonObserveResponse CreateBuyerLockonObserve(ctx, id).Execute()

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
    // response from `CreateBuyerLockonObserve`: CreateBuyerLockonObserveResponse
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

[**CreateBuyerLockonObserveResponse**](CreateBuyerLockonObserveResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
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

[**CreateBuyerScanResponse**](CreateBuyerScanResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
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

[**CreateBuyerScanSessionResponse**](CreateBuyerScanSessionResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
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

[**CreateBuyerScanSessionEndResponse**](CreateBuyerScanSessionEndResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
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

[**GetBuyerAnywhereResponse**](GetBuyerAnywhereResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
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

[**GetBuyerCatalogFacetResponse**](GetBuyerCatalogFacetResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
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

[**GetBuyerCatalogListingResponse**](GetBuyerCatalogListingResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
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

[**GetBuyerCatalogListingAvailabilityResponse**](GetBuyerCatalogListingAvailabilityResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
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

[**GetBuyerScanSessionResponse**](GetBuyerScanSessionResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuyerCatalogSearch

> V1List ListBuyerCatalogSearch(ctx).Execute()

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
    // response from `ListBuyerCatalogSearch`: V1List
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.ListBuyerCatalogSearch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerCatalogSearchRequest struct via the builder pattern


### Return type

[**V1List**](V1List.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBuyerScanSessions

> V1List ListBuyerScanSessions(ctx).Execute()

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
    // response from `ListBuyerScanSessions`: V1List
    fmt.Fprintf(os.Stdout, "Response from `BuyerCatalogApi.ListBuyerScanSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBuyerScanSessionsRequest struct via the builder pattern


### Return type

[**V1List**](V1List.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

