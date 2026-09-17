# \ListingsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateListing**](ListingsApi.md#CreateListing) | **Post** /v1/listings | Create a listing and fan out crosspost jobs across platforms.
[**CreateListingBulkCheckStatus**](ListingsApi.md#CreateListingBulkCheckStatus) | **Post** /v1/listings/bulk-check-status | Check listing status on platforms
[**CreateListingBulkCrosspost**](ListingsApi.md#CreateListingBulkCrosspost) | **Post** /v1/listings/bulk-crosspost | Bulk crosspost (no delist phase)
[**CreateListingBulkDelete**](ListingsApi.md#CreateListingBulkDelete) | **Post** /v1/listings/bulk-delete | Bulk archive + delist
[**CreateListingBulkDelist**](ListingsApi.md#CreateListingBulkDelist) | **Post** /v1/listings/bulk-delist | Bulk delist from platforms
[**CreateListingBulkDelistPreview**](ListingsApi.md#CreateListingBulkDelistPreview) | **Post** /v1/listings/bulk-delist-preview | Preview which marketplaces a delist would touch
[**CreateListingBulkHardDelete**](ListingsApi.md#CreateListingBulkHardDelete) | **Post** /v1/listings/bulk-hard-delete | Permanently delete archived listings
[**CreateListingBulkRelist**](ListingsApi.md#CreateListingBulkRelist) | **Post** /v1/listings/bulk-relist | Bulk relist across platforms
[**CreateListingBulkUpdate**](ListingsApi.md#CreateListingBulkUpdate) | **Post** /v1/listings/bulk-update | Bulk update listing fields
[**CreateListingById**](ListingsApi.md#CreateListingById) | **Post** /v1/listings/by-ids | Fetch hydrated listings by ID
[**CreateListingCheckDuplicate**](ListingsApi.md#CreateListingCheckDuplicate) | **Post** /v1/listings/check-duplicates | Check whether the seller already owns something matching this title/photo, and what to do about it.
[**CreateListingCombine**](ListingsApi.md#CreateListingCombine) | **Post** /v1/listings/combine | Combine duplicate listings into one: sums their stock, delists and archives the rest.
[**CreateListingDiscrepancyResolve**](ListingsApi.md#CreateListingDiscrepancyResolve) | **Post** /v1/listings/{id}/discrepancies/{discrepancyId}/resolve | Resolve a detected marketplace-drift discrepancy: accept the platform value, push ours back, relist to apply it, or dismiss.
[**CreateListingImportByUrl**](ListingsApi.md#CreateListingImportByUrl) | **Post** /v1/listings/{id}/import-by-url | Attach a real platform listing to this listing by pasting its live URL.
[**CreateListingMagicFill**](ListingsApi.md#CreateListingMagicFill) | **Post** /v1/listings/{id}/magic-fill | Auto-fill empty fields on one platform tab from the master listing + AI/deterministic taxonomy resolution.
[**DeleteListing**](ListingsApi.md#DeleteListing) | **Delete** /v1/listings/{id} | Delist a listing (optionally narrowed to specific platforms via ?platforms&#x3D;).
[**GetListing**](ListingsApi.md#GetListing) | **Get** /v1/listings/{id} | Get one listing with its platform rows.
[**GetListingFacet**](ListingsApi.md#GetListingFacet) | **Get** /v1/listings/facets | Distinct brands + categories across listings + inventory.
[**GetListingSkuExist**](ListingsApi.md#GetListingSkuExist) | **Get** /v1/listings/sku-exists | Check whether a SKU is already used by one of this user&#39;s items.
[**ListListingDiscrepancies**](ListingsApi.md#ListListingDiscrepancies) | **Get** /v1/listings/{id}/discrepancies | List detected marketplace-drift discrepancies for a listing.
[**ListListingIds**](ListingsApi.md#ListListingIds) | **Get** /v1/listings/ids | Filter listings → return matching id list (no pagination).
[**ListListings**](ListingsApi.md#ListListings) | **Get** /v1/listings | List active platform listings.
[**UpdateListing**](ListingsApi.md#UpdateListing) | **Patch** /v1/listings/{id} | Edit a listing and fan out update jobs to existing platform listings.



## CreateListing

> CreateListingResponse CreateListing(ctx).Execute()

Create a listing and fan out crosspost jobs across platforms.

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
    resp, r, err := apiClient.ListingsApi.CreateListing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListing`: CreateListingResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingRequest struct via the builder pattern


### Return type

[**CreateListingResponse**](CreateListingResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListingBulkCheckStatus

> CreateListingBulkCheckStatusResponse CreateListingBulkCheckStatus(ctx).Execute()

Check listing status on platforms

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
    resp, r, err := apiClient.ListingsApi.CreateListingBulkCheckStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingBulkCheckStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingBulkCheckStatus`: CreateListingBulkCheckStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingBulkCheckStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingBulkCheckStatusRequest struct via the builder pattern


### Return type

[**CreateListingBulkCheckStatusResponse**](CreateListingBulkCheckStatusResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListingBulkCrosspost

> CreateListingBulkCrosspostResponse CreateListingBulkCrosspost(ctx).Execute()

Bulk crosspost (no delist phase)

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
    resp, r, err := apiClient.ListingsApi.CreateListingBulkCrosspost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingBulkCrosspost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingBulkCrosspost`: CreateListingBulkCrosspostResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingBulkCrosspost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingBulkCrosspostRequest struct via the builder pattern


### Return type

[**CreateListingBulkCrosspostResponse**](CreateListingBulkCrosspostResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListingBulkDelete

> CreateListingBulkDeleteResponse CreateListingBulkDelete(ctx).Execute()

Bulk archive + delist

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
    resp, r, err := apiClient.ListingsApi.CreateListingBulkDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingBulkDelete`: CreateListingBulkDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingBulkDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingBulkDeleteRequest struct via the builder pattern


### Return type

[**CreateListingBulkDeleteResponse**](CreateListingBulkDeleteResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListingBulkDelist

> CreateListingBulkDelistResponse CreateListingBulkDelist(ctx).Execute()

Bulk delist from platforms

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
    resp, r, err := apiClient.ListingsApi.CreateListingBulkDelist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingBulkDelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingBulkDelist`: CreateListingBulkDelistResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingBulkDelist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingBulkDelistRequest struct via the builder pattern


### Return type

[**CreateListingBulkDelistResponse**](CreateListingBulkDelistResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListingBulkDelistPreview

> CreateListingBulkDelistPreviewResponse CreateListingBulkDelistPreview(ctx).Execute()

Preview which marketplaces a delist would touch

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
    resp, r, err := apiClient.ListingsApi.CreateListingBulkDelistPreview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingBulkDelistPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingBulkDelistPreview`: CreateListingBulkDelistPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingBulkDelistPreview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingBulkDelistPreviewRequest struct via the builder pattern


### Return type

[**CreateListingBulkDelistPreviewResponse**](CreateListingBulkDelistPreviewResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListingBulkHardDelete

> CreateListingBulkHardDeleteResponse CreateListingBulkHardDelete(ctx).Execute()

Permanently delete archived listings

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
    resp, r, err := apiClient.ListingsApi.CreateListingBulkHardDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingBulkHardDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingBulkHardDelete`: CreateListingBulkHardDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingBulkHardDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingBulkHardDeleteRequest struct via the builder pattern


### Return type

[**CreateListingBulkHardDeleteResponse**](CreateListingBulkHardDeleteResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListingBulkRelist

> CreateListingBulkRelistResponse CreateListingBulkRelist(ctx).Execute()

Bulk relist across platforms

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
    resp, r, err := apiClient.ListingsApi.CreateListingBulkRelist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingBulkRelist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingBulkRelist`: CreateListingBulkRelistResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingBulkRelist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingBulkRelistRequest struct via the builder pattern


### Return type

[**CreateListingBulkRelistResponse**](CreateListingBulkRelistResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListingBulkUpdate

> CreateListingBulkUpdateResponse CreateListingBulkUpdate(ctx).Execute()

Bulk update listing fields

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
    resp, r, err := apiClient.ListingsApi.CreateListingBulkUpdate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingBulkUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingBulkUpdate`: CreateListingBulkUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingBulkUpdate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingBulkUpdateRequest struct via the builder pattern


### Return type

[**CreateListingBulkUpdateResponse**](CreateListingBulkUpdateResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListingById

> CreateListingByIdResponse CreateListingById(ctx).Execute()

Fetch hydrated listings by ID

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
    resp, r, err := apiClient.ListingsApi.CreateListingById(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingById`: CreateListingByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingById`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingByIdRequest struct via the builder pattern


### Return type

[**CreateListingByIdResponse**](CreateListingByIdResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListingCheckDuplicate

> CreateListingCheckDuplicateResponse CreateListingCheckDuplicate(ctx).Execute()

Check whether the seller already owns something matching this title/photo, and what to do about it.

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
    resp, r, err := apiClient.ListingsApi.CreateListingCheckDuplicate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingCheckDuplicate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingCheckDuplicate`: CreateListingCheckDuplicateResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingCheckDuplicate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingCheckDuplicateRequest struct via the builder pattern


### Return type

[**CreateListingCheckDuplicateResponse**](CreateListingCheckDuplicateResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListingCombine

> CreateListingCombineResponse CreateListingCombine(ctx).Execute()

Combine duplicate listings into one: sums their stock, delists and archives the rest.

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
    resp, r, err := apiClient.ListingsApi.CreateListingCombine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingCombine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingCombine`: CreateListingCombineResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingCombine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingCombineRequest struct via the builder pattern


### Return type

[**CreateListingCombineResponse**](CreateListingCombineResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListingDiscrepancyResolve

> CreateListingDiscrepancyResolveResponse CreateListingDiscrepancyResolve(ctx, id, discrepancyId).Execute()

Resolve a detected marketplace-drift discrepancy: accept the platform value, push ours back, relist to apply it, or dismiss.

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
    discrepancyId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsApi.CreateListingDiscrepancyResolve(context.Background(), id, discrepancyId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingDiscrepancyResolve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingDiscrepancyResolve`: CreateListingDiscrepancyResolveResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingDiscrepancyResolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**discrepancyId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingDiscrepancyResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**CreateListingDiscrepancyResolveResponse**](CreateListingDiscrepancyResolveResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListingImportByUrl

> CreateListingImportByUrlResponse CreateListingImportByUrl(ctx, id).Execute()

Attach a real platform listing to this listing by pasting its live URL.

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
    resp, r, err := apiClient.ListingsApi.CreateListingImportByUrl(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingImportByUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingImportByUrl`: CreateListingImportByUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingImportByUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingImportByUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateListingImportByUrlResponse**](CreateListingImportByUrlResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateListingMagicFill

> CreateListingMagicFillResponse CreateListingMagicFill(ctx, id).Execute()

Auto-fill empty fields on one platform tab from the master listing + AI/deterministic taxonomy resolution.

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
    resp, r, err := apiClient.ListingsApi.CreateListingMagicFill(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingMagicFill``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingMagicFill`: CreateListingMagicFillResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingMagicFill`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingMagicFillRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateListingMagicFillResponse**](CreateListingMagicFillResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteListing

> DeleteListingResponse DeleteListing(ctx, id).Platforms(platforms).Execute()

Delist a listing (optionally narrowed to specific platforms via ?platforms=).

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
    platforms := "platforms_example" // string | Comma-separated platform slugs to limit the delist fan-out. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsApi.DeleteListing(context.Background(), id).Platforms(platforms).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.DeleteListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteListing`: DeleteListingResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.DeleteListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **platforms** | **string** | Comma-separated platform slugs to limit the delist fan-out. | 

### Return type

[**DeleteListingResponse**](DeleteListingResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListing

> GetListingResponse GetListing(ctx, id).Execute()

Get one listing with its platform rows.

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
    resp, r, err := apiClient.ListingsApi.GetListing(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListing`: GetListingResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetListingResponse**](GetListingResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingFacet

> GetListingFacetResponse GetListingFacet(ctx).Execute()

Distinct brands + categories across listings + inventory.

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
    resp, r, err := apiClient.ListingsApi.GetListingFacet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetListingFacet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListingFacet`: GetListingFacetResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetListingFacet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingFacetRequest struct via the builder pattern


### Return type

[**GetListingFacetResponse**](GetListingFacetResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetListingSkuExist

> GetListingSkuExistResponse GetListingSkuExist(ctx).Sku(sku).Execute()

Check whether a SKU is already used by one of this user's items.

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
    sku := "sku_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsApi.GetListingSkuExist(context.Background()).Sku(sku).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetListingSkuExist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListingSkuExist`: GetListingSkuExistResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetListingSkuExist`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetListingSkuExistRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sku** | **string** |  | 

### Return type

[**GetListingSkuExistResponse**](GetListingSkuExistResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListListingDiscrepancies

> V1List ListListingDiscrepancies(ctx, id).Execute()

List detected marketplace-drift discrepancies for a listing.

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
    resp, r, err := apiClient.ListingsApi.ListListingDiscrepancies(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListListingDiscrepancies``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListListingDiscrepancies`: V1List
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListListingDiscrepancies`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListListingDiscrepanciesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


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


## ListListingIds

> V1List ListListingIds(ctx).Execute()

Filter listings → return matching id list (no pagination).

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
    resp, r, err := apiClient.ListingsApi.ListListingIds(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListListingIds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListListingIds`: V1List
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListListingIds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListListingIdsRequest struct via the builder pattern


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


## ListListings

> V1List ListListings(ctx).Page(page).Limit(limit).Platform(platform).Status(status).Execute()

List active platform listings.

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
    page := int32(56) // int32 |  (optional) (default to 1)
    limit := int32(56) // int32 |  (optional) (default to 25)
    platform := "platform_example" // string |  (optional)
    status := "status_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsApi.ListListings(context.Background()).Page(page).Limit(limit).Platform(platform).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListListings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListListings`: V1List
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListListings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListListingsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 25]
 **platform** | **string** |  | 
 **status** | **string** |  | 

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


## UpdateListing

> UpdateListingResponse UpdateListing(ctx, id).Execute()

Edit a listing and fan out update jobs to existing platform listings.

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
    resp, r, err := apiClient.ListingsApi.UpdateListing(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.UpdateListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateListing`: UpdateListingResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.UpdateListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UpdateListingResponse**](UpdateListingResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

