# \ProductionApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateProductionCandidateAccept**](ProductionApi.md#CreateProductionCandidateAccept) | **Post** /v1/production/candidates/{id}/accept | Accept a graduated design&#39;s production terms (edition size, cost, holder access).
[**CreateProductionCandidatePo**](ProductionApi.md#CreateProductionCandidatePo) | **Post** /v1/production/candidates/{id}/po | Ops: open a purchase order for a candidate. Admin PAT only.
[**CreateProductionCandidateVote**](ProductionApi.md#CreateProductionCandidateVote) | **Post** /v1/production/candidates/{id}/vote | Back (support: true) or un-back a design on the production ballot. One vote per person per design; voters must have bought on Crossly.
[**CreateProductionDesignBrief**](ProductionApi.md#CreateProductionDesignBrief) | **Post** /v1/production/designs/{cosmeticId}/brief | Set a design&#39;s production brief (materials, finish, box, reference images).
[**CreateProductionDesignConsent**](ProductionApi.md#CreateProductionDesignConsent) | **Post** /v1/production/designs/{cosmeticId}/consent | Opt a design in or out of physical production.
[**CreateProductionDesignSpec**](ProductionApi.md#CreateProductionDesignSpec) | **Post** /v1/production/designs/{cosmeticId}/spec | Set a design&#39;s production spec (category, materials, complexity, packaging).
[**CreateProductionEstimate**](ProductionApi.md#CreateProductionEstimate) | **Post** /v1/production/estimate | Estimate a run&#39;s landed cost from a spec + edition size (overshoots for safety).
[**CreateProductionRunPurchase**](ProductionApi.md#CreateProductionRunPurchase) | **Post** /v1/production/runs/{id}/purchase | Buy a limited-edition unit of a physical run (holder pricing applies).
[**CreateReceiptRelist**](ProductionApi.md#CreateReceiptRelist) | **Post** /v1/receipts/{id}/relist | Sell it again: a draft listing of something you bought on Crossly, carrying its purchase receipt so the next buyer inherits its provenance.
[**GetProductionBallot**](ProductionApi.md#GetProductionBallot) | **Get** /v1/production/ballot | The ballot: designs waiting for a physical run, most backed first, and whether you may vote.
[**GetProductionCandidate**](ProductionApi.md#GetProductionCandidate) | **Get** /v1/production/candidates | Your designs that have graduated toward a physical run.
[**GetProductionDesignPref**](ProductionApi.md#GetProductionDesignPref) | **Get** /v1/production/designs/{cosmeticId}/prefs | Read a design&#39;s production preferences (spec, brief, opt-out).
[**GetProductionDrop**](ProductionApi.md#GetProductionDrop) | **Get** /v1/production/drops | Limited-edition units you own, with fulfillment status.
[**GetProductionEarning**](ProductionApi.md#GetProductionEarning) | **Get** /v1/production/earnings | Your CBX earnings from physical production (royalties + balance + lines).
[**GetProductionLine**](ProductionApi.md#GetProductionLine) | **Get** /v1/production/lines | Your production lines dashboard (candidate + run + inventory).
[**GetProductionPo**](ProductionApi.md#GetProductionPo) | **Get** /v1/production/pos | Ops: list purchase orders (optionally for one candidate). Admin PAT only.
[**GetProductionPoReconcile**](ProductionApi.md#GetProductionPoReconcile) | **Get** /v1/production/pos/{poId}/reconcile | Ops: reconcile a PO&#39;s actuals against the estimate. Admin PAT only.
[**GetProductionRun**](ProductionApi.md#GetProductionRun) | **Get** /v1/production/runs | Browse open limited-edition physical runs (graduated designs).
[**GetShopListingProvenance**](ProductionApi.md#GetShopListingProvenance) | **Get** /v1/shop/listings/{slug}/provenance | Crossly provenance behind a shop listing: owners on Crossly and since when (never who or for how much). Null when not bought here.
[**UpdateProductionPo**](ProductionApi.md#UpdateProductionPo) | **Patch** /v1/production/pos/{poId} | Ops: update a PO (supplier, freight, actuals, QC, status). Admin PAT only.



## CreateProductionCandidateAccept

> CreateProductionCandidateAcceptResponse CreateProductionCandidateAccept(ctx, id).Execute()

Accept a graduated design's production terms (edition size, cost, holder access).

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
    resp, r, err := apiClient.ProductionApi.CreateProductionCandidateAccept(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.CreateProductionCandidateAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProductionCandidateAccept`: CreateProductionCandidateAcceptResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.CreateProductionCandidateAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductionCandidateAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateProductionCandidateAcceptResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateProductionCandidatePo

> UpdateProductionPoResponse CreateProductionCandidatePo(ctx, id).Execute()

Ops: open a purchase order for a candidate. Admin PAT only.

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
    resp, r, err := apiClient.ProductionApi.CreateProductionCandidatePo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.CreateProductionCandidatePo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProductionCandidatePo`: UpdateProductionPoResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.CreateProductionCandidatePo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductionCandidatePoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateProductionPoResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateProductionCandidateVote

> CreateProductionCandidateVoteResponse CreateProductionCandidateVote(ctx, id).Execute()

Back (support: true) or un-back a design on the production ballot. One vote per person per design; voters must have bought on Crossly.

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
    resp, r, err := apiClient.ProductionApi.CreateProductionCandidateVote(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.CreateProductionCandidateVote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProductionCandidateVote`: CreateProductionCandidateVoteResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.CreateProductionCandidateVote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductionCandidateVoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateProductionCandidateVoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateProductionDesignBrief

> CreateProductionDesignSpecResponse CreateProductionDesignBrief(ctx, cosmeticId).Execute()

Set a design's production brief (materials, finish, box, reference images).

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
    cosmeticId := "cosmeticId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductionApi.CreateProductionDesignBrief(context.Background(), cosmeticId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.CreateProductionDesignBrief``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProductionDesignBrief`: CreateProductionDesignSpecResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.CreateProductionDesignBrief`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cosmeticId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductionDesignBriefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateProductionDesignSpecResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateProductionDesignConsent

> CreateProductionDesignSpecResponse CreateProductionDesignConsent(ctx, cosmeticId).Execute()

Opt a design in or out of physical production.

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
    cosmeticId := "cosmeticId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductionApi.CreateProductionDesignConsent(context.Background(), cosmeticId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.CreateProductionDesignConsent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProductionDesignConsent`: CreateProductionDesignSpecResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.CreateProductionDesignConsent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cosmeticId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductionDesignConsentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateProductionDesignSpecResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateProductionDesignSpec

> CreateProductionDesignSpecResponse CreateProductionDesignSpec(ctx, cosmeticId).Execute()

Set a design's production spec (category, materials, complexity, packaging).

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
    cosmeticId := "cosmeticId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductionApi.CreateProductionDesignSpec(context.Background(), cosmeticId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.CreateProductionDesignSpec``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProductionDesignSpec`: CreateProductionDesignSpecResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.CreateProductionDesignSpec`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cosmeticId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductionDesignSpecRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateProductionDesignSpecResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateProductionEstimate

> CreateProductionEstimateResponse CreateProductionEstimate(ctx).Execute()

Estimate a run's landed cost from a spec + edition size (overshoots for safety).

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
    resp, r, err := apiClient.ProductionApi.CreateProductionEstimate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.CreateProductionEstimate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProductionEstimate`: CreateProductionEstimateResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.CreateProductionEstimate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductionEstimateRequest struct via the builder pattern


### Return type

**CreateProductionEstimateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateProductionRunPurchase

> CreateProductionRunPurchaseResponse CreateProductionRunPurchase(ctx, id).Execute()

Buy a limited-edition unit of a physical run (holder pricing applies).

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
    resp, r, err := apiClient.ProductionApi.CreateProductionRunPurchase(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.CreateProductionRunPurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateProductionRunPurchase`: CreateProductionRunPurchaseResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.CreateProductionRunPurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateProductionRunPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateProductionRunPurchaseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateReceiptRelist

> CreateReceiptRelistResponse CreateReceiptRelist(ctx, id).Execute()

Sell it again: a draft listing of something you bought on Crossly, carrying its purchase receipt so the next buyer inherits its provenance.

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
    resp, r, err := apiClient.ProductionApi.CreateReceiptRelist(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.CreateReceiptRelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReceiptRelist`: CreateReceiptRelistResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.CreateReceiptRelist`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReceiptRelistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateReceiptRelistResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetProductionBallot

> GetProductionBallotResponse GetProductionBallot(ctx).Execute()

The ballot: designs waiting for a physical run, most backed first, and whether you may vote.

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
    resp, r, err := apiClient.ProductionApi.GetProductionBallot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.GetProductionBallot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductionBallot`: GetProductionBallotResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.GetProductionBallot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductionBallotRequest struct via the builder pattern


### Return type

**GetProductionBallotResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetProductionCandidate

> GetProductionCandidateResponse GetProductionCandidate(ctx).Execute()

Your designs that have graduated toward a physical run.

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
    resp, r, err := apiClient.ProductionApi.GetProductionCandidate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.GetProductionCandidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductionCandidate`: GetProductionCandidateResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.GetProductionCandidate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductionCandidateRequest struct via the builder pattern


### Return type

**GetProductionCandidateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetProductionDesignPref

> GetProductionDesignPrefResponse GetProductionDesignPref(ctx, cosmeticId).Execute()

Read a design's production preferences (spec, brief, opt-out).

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
    cosmeticId := "cosmeticId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductionApi.GetProductionDesignPref(context.Background(), cosmeticId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.GetProductionDesignPref``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductionDesignPref`: GetProductionDesignPrefResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.GetProductionDesignPref`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cosmeticId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductionDesignPrefRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetProductionDesignPrefResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetProductionDrop

> GetProductionDropResponse GetProductionDrop(ctx).Execute()

Limited-edition units you own, with fulfillment status.

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
    resp, r, err := apiClient.ProductionApi.GetProductionDrop(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.GetProductionDrop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductionDrop`: GetProductionDropResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.GetProductionDrop`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductionDropRequest struct via the builder pattern


### Return type

**GetProductionDropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetProductionEarning

> GetProductionEarningResponse GetProductionEarning(ctx).Execute()

Your CBX earnings from physical production (royalties + balance + lines).

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
    resp, r, err := apiClient.ProductionApi.GetProductionEarning(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.GetProductionEarning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductionEarning`: GetProductionEarningResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.GetProductionEarning`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductionEarningRequest struct via the builder pattern


### Return type

**GetProductionEarningResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetProductionLine

> GetProductionLineResponse GetProductionLine(ctx).Execute()

Your production lines dashboard (candidate + run + inventory).

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
    resp, r, err := apiClient.ProductionApi.GetProductionLine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.GetProductionLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductionLine`: GetProductionLineResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.GetProductionLine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductionLineRequest struct via the builder pattern


### Return type

**GetProductionLineResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetProductionPo

> GetProductionPoResponse GetProductionPo(ctx).CandidateId(candidateId).Execute()

Ops: list purchase orders (optionally for one candidate). Admin PAT only.

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
    candidateId := "candidateId_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductionApi.GetProductionPo(context.Background()).CandidateId(candidateId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.GetProductionPo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductionPo`: GetProductionPoResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.GetProductionPo`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetProductionPoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **candidateId** | **string** |  | 

### Return type

**GetProductionPoResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetProductionPoReconcile

> GetProductionPoReconcileResponse GetProductionPoReconcile(ctx, poId).Execute()

Ops: reconcile a PO's actuals against the estimate. Admin PAT only.

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
    poId := "poId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductionApi.GetProductionPoReconcile(context.Background(), poId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.GetProductionPoReconcile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductionPoReconcile`: GetProductionPoReconcileResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.GetProductionPoReconcile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductionPoReconcileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetProductionPoReconcileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetProductionRun

> GetProductionRunResponse GetProductionRun(ctx).Execute()

Browse open limited-edition physical runs (graduated designs).

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
    resp, r, err := apiClient.ProductionApi.GetProductionRun(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.GetProductionRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetProductionRun`: GetProductionRunResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.GetProductionRun`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProductionRunRequest struct via the builder pattern


### Return type

**GetProductionRunResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetShopListingProvenance

> GetShopListingProvenanceResponse GetShopListingProvenance(ctx, slug).Execute()

Crossly provenance behind a shop listing: owners on Crossly and since when (never who or for how much). Null when not bought here.

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
    resp, r, err := apiClient.ProductionApi.GetShopListingProvenance(context.Background(), slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.GetShopListingProvenance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShopListingProvenance`: GetShopListingProvenanceResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.GetShopListingProvenance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopListingProvenanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetShopListingProvenanceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateProductionPo

> UpdateProductionPoResponse UpdateProductionPo(ctx, poId).Execute()

Ops: update a PO (supplier, freight, actuals, QC, status). Admin PAT only.

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
    poId := "poId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ProductionApi.UpdateProductionPo(context.Background(), poId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ProductionApi.UpdateProductionPo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProductionPo`: UpdateProductionPoResponse
    fmt.Fprintf(os.Stdout, "Response from `ProductionApi.UpdateProductionPo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**poId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProductionPoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateProductionPoResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

