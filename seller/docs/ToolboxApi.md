# \ToolboxApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateToolboxAsset**](ToolboxApi.md#CreateToolboxAsset) | **Post** /v1/toolbox/assets | Publish an asset (place nodes). Scanned (lint, text, credits, uploads) then reviewed. Free by default; paid needs Pro and a verified identity.
[**CreateToolboxAssetPurchase**](ToolboxApi.md#CreateToolboxAssetPurchase) | **Post** /v1/toolbox/assets/{assetId}/purchase | Open a Crossly checkout for a paid asset. Nothing is charged here: a person pays at checkoutUrl. The author earns 80% of net; the licence is use in your own experiences.
[**CreateToolboxAssetRetire**](ToolboxApi.md#CreateToolboxAssetRetire) | **Post** /v1/toolbox/assets/{assetId}/retire | Take an asset off the store for good. Licences already sold keep working.
[**CreateToolboxAssetUse**](ToolboxApi.md#CreateToolboxAssetUse) | **Post** /v1/toolbox/assets/{assetId}/use | Get the asset&#39;s nodes stamped with its id for a place you can edit; insert:true also inserts them into the place&#39;s live session. Paid assets need the experience owner&#39;s licence.
[**CreateToolboxAssetVersion**](ToolboxApi.md#CreateToolboxAssetVersion) | **Post** /v1/toolbox/assets/{assetId}/versions | A new version (reviewed); the live one keeps serving until it is approved.
[**CreateToolboxPurchaseRefund**](ToolboxApi.md#CreateToolboxPurchaseRefund) | **Post** /v1/toolbox/purchases/{purchaseId}/refund | Refund a licence within 14 days while no published version of yours uses the asset. Ends the licence; the author&#39;s share is reversed.
[**DeleteToolboxAssetRatingMine**](ToolboxApi.md#DeleteToolboxAssetRatingMine) | **Delete** /v1/toolbox/assets/{assetId}/ratings/mine | Remove your rating of an asset.
[**GetToolboxAsset**](ToolboxApi.md#GetToolboxAsset) | **Get** /v1/toolbox/assets/{assetId} | One asset: versions, credits (assets it includes), licence, contents, your purchase.
[**GetToolboxAssetRating**](ToolboxApi.md#GetToolboxAssetRating) | **Get** /v1/toolbox/assets/{assetId}/ratings | Ratings of an asset: average stars, breakdown, recent reviews, yours, and whether you may rate it.
[**GetToolboxPurchase**](ToolboxApi.md#GetToolboxPurchase) | **Get** /v1/toolbox/purchases/{purchaseId} | A purchase: status, price, tax, refundable until.
[**ListToolboxAssets**](ToolboxApi.md#ListToolboxAssets) | **Get** /v1/toolbox/assets | Search the Toolbox: models, scripts, sounds, images, packages. Free and paid; &#x60;owned&#x60; says whether you may insert it now.
[**ListToolboxLibrary**](ToolboxApi.md#ListToolboxLibrary) | **Get** /v1/toolbox/library | Assets you may insert: licences you bought, your own, free ones you used.
[**ListToolboxMine**](ToolboxApi.md#ListToolboxMine) | **Get** /v1/toolbox/mine | Your published assets: review state, reviewer note, what each earned you.
[**ListToolboxSales**](ToolboxApi.md#ListToolboxSales) | **Get** /v1/toolbox/sales | Your Toolbox sales: gross, net, your 80% share after refunds. Never who bought.
[**UpdateToolboxAsset**](ToolboxApi.md#UpdateToolboxAsset) | **Patch** /v1/toolbox/assets/{assetId} | Edit name, description, tags, thumbnail or price.
[**UpdateToolboxAssetRatingMine**](ToolboxApi.md#UpdateToolboxAssetRatingMine) | **Put** /v1/toolbox/assets/{assetId}/ratings/mine | Rate an asset you bought or installed: 1–5 stars and an optional review (1000 characters). Not your own.



## CreateToolboxAsset

> ToolboxAssetResponse CreateToolboxAsset(ctx).Execute()

Publish an asset (place nodes). Scanned (lint, text, credits, uploads) then reviewed. Free by default; paid needs Pro and a verified identity.

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
    resp, r, err := apiClient.ToolboxApi.CreateToolboxAsset(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.CreateToolboxAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToolboxAsset`: ToolboxAssetResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.CreateToolboxAsset`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateToolboxAssetRequest struct via the builder pattern


### Return type

**ToolboxAssetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateToolboxAssetPurchase

> GetToolboxPurchaseResponse CreateToolboxAssetPurchase(ctx, assetId).Execute()

Open a Crossly checkout for a paid asset. Nothing is charged here: a person pays at checkoutUrl. The author earns 80% of net; the licence is use in your own experiences.

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
    assetId := "assetId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ToolboxApi.CreateToolboxAssetPurchase(context.Background(), assetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.CreateToolboxAssetPurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToolboxAssetPurchase`: GetToolboxPurchaseResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.CreateToolboxAssetPurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateToolboxAssetPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetToolboxPurchaseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateToolboxAssetRetire

> ToolboxAssetResponse CreateToolboxAssetRetire(ctx, assetId).Execute()

Take an asset off the store for good. Licences already sold keep working.

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
    assetId := "assetId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ToolboxApi.CreateToolboxAssetRetire(context.Background(), assetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.CreateToolboxAssetRetire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToolboxAssetRetire`: ToolboxAssetResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.CreateToolboxAssetRetire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateToolboxAssetRetireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ToolboxAssetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateToolboxAssetUse

> CreateToolboxAssetUseResponse CreateToolboxAssetUse(ctx, assetId).Execute()

Get the asset's nodes stamped with its id for a place you can edit; insert:true also inserts them into the place's live session. Paid assets need the experience owner's licence.

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
    assetId := "assetId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ToolboxApi.CreateToolboxAssetUse(context.Background(), assetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.CreateToolboxAssetUse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToolboxAssetUse`: CreateToolboxAssetUseResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.CreateToolboxAssetUse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateToolboxAssetUseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateToolboxAssetUseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateToolboxAssetVersion

> ToolboxAssetResponse CreateToolboxAssetVersion(ctx, assetId).Execute()

A new version (reviewed); the live one keeps serving until it is approved.

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
    assetId := "assetId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ToolboxApi.CreateToolboxAssetVersion(context.Background(), assetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.CreateToolboxAssetVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToolboxAssetVersion`: ToolboxAssetResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.CreateToolboxAssetVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateToolboxAssetVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ToolboxAssetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateToolboxPurchaseRefund

> CreateWorldMeReceiptRefundResponse CreateToolboxPurchaseRefund(ctx, purchaseId).Execute()

Refund a licence within 14 days while no published version of yours uses the asset. Ends the licence; the author's share is reversed.

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
    purchaseId := "purchaseId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ToolboxApi.CreateToolboxPurchaseRefund(context.Background(), purchaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.CreateToolboxPurchaseRefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateToolboxPurchaseRefund`: CreateWorldMeReceiptRefundResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.CreateToolboxPurchaseRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**purchaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateToolboxPurchaseRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldMeReceiptRefundResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteToolboxAssetRatingMine

> DeleteSpatialSceneGrantResponse DeleteToolboxAssetRatingMine(ctx, assetId).Execute()

Remove your rating of an asset.

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
    assetId := "assetId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ToolboxApi.DeleteToolboxAssetRatingMine(context.Background(), assetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.DeleteToolboxAssetRatingMine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteToolboxAssetRatingMine`: DeleteSpatialSceneGrantResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.DeleteToolboxAssetRatingMine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteToolboxAssetRatingMineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteSpatialSceneGrantResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetToolboxAsset

> GetToolboxAssetResponse GetToolboxAsset(ctx, assetId).Execute()

One asset: versions, credits (assets it includes), licence, contents, your purchase.

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
    assetId := "assetId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ToolboxApi.GetToolboxAsset(context.Background(), assetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.GetToolboxAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetToolboxAsset`: GetToolboxAssetResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.GetToolboxAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolboxAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetToolboxAssetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetToolboxAssetRating

> GetToolboxAssetRatingResponse GetToolboxAssetRating(ctx, assetId).Execute()

Ratings of an asset: average stars, breakdown, recent reviews, yours, and whether you may rate it.

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
    assetId := "assetId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ToolboxApi.GetToolboxAssetRating(context.Background(), assetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.GetToolboxAssetRating``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetToolboxAssetRating`: GetToolboxAssetRatingResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.GetToolboxAssetRating`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolboxAssetRatingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetToolboxAssetRatingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetToolboxPurchase

> GetToolboxPurchaseResponse GetToolboxPurchase(ctx, purchaseId).Execute()

A purchase: status, price, tax, refundable until.

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
    purchaseId := "purchaseId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ToolboxApi.GetToolboxPurchase(context.Background(), purchaseId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.GetToolboxPurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetToolboxPurchase`: GetToolboxPurchaseResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.GetToolboxPurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**purchaseId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetToolboxPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetToolboxPurchaseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListToolboxAssets

> ListToolboxAssetsResponse ListToolboxAssets(ctx).Execute()

Search the Toolbox: models, scripts, sounds, images, packages. Free and paid; `owned` says whether you may insert it now.

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
    resp, r, err := apiClient.ToolboxApi.ListToolboxAssets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.ListToolboxAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListToolboxAssets`: ListToolboxAssetsResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.ListToolboxAssets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListToolboxAssetsRequest struct via the builder pattern


### Return type

**ListToolboxAssetsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListToolboxLibrary

> ListToolboxLibraryResponse ListToolboxLibrary(ctx).Execute()

Assets you may insert: licences you bought, your own, free ones you used.

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
    resp, r, err := apiClient.ToolboxApi.ListToolboxLibrary(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.ListToolboxLibrary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListToolboxLibrary`: ListToolboxLibraryResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.ListToolboxLibrary`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListToolboxLibraryRequest struct via the builder pattern


### Return type

**ListToolboxLibraryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListToolboxMine

> ListToolboxMineResponse ListToolboxMine(ctx).Execute()

Your published assets: review state, reviewer note, what each earned you.

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
    resp, r, err := apiClient.ToolboxApi.ListToolboxMine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.ListToolboxMine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListToolboxMine`: ListToolboxMineResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.ListToolboxMine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListToolboxMineRequest struct via the builder pattern


### Return type

**ListToolboxMineResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListToolboxSales

> ListToolboxSalesResponse ListToolboxSales(ctx).Execute()

Your Toolbox sales: gross, net, your 80% share after refunds. Never who bought.

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
    resp, r, err := apiClient.ToolboxApi.ListToolboxSales(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.ListToolboxSales``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListToolboxSales`: ListToolboxSalesResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.ListToolboxSales`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListToolboxSalesRequest struct via the builder pattern


### Return type

**ListToolboxSalesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateToolboxAsset

> ToolboxAssetResponse UpdateToolboxAsset(ctx, assetId).Execute()

Edit name, description, tags, thumbnail or price.

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
    assetId := "assetId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ToolboxApi.UpdateToolboxAsset(context.Background(), assetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.UpdateToolboxAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateToolboxAsset`: ToolboxAssetResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.UpdateToolboxAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateToolboxAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ToolboxAssetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateToolboxAssetRatingMine

> UpdateToolboxAssetRatingMineResponse UpdateToolboxAssetRatingMine(ctx, assetId).Execute()

Rate an asset you bought or installed: 1–5 stars and an optional review (1000 characters). Not your own.

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
    assetId := "assetId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ToolboxApi.UpdateToolboxAssetRatingMine(context.Background(), assetId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ToolboxApi.UpdateToolboxAssetRatingMine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateToolboxAssetRatingMine`: UpdateToolboxAssetRatingMineResponse
    fmt.Fprintf(os.Stdout, "Response from `ToolboxApi.UpdateToolboxAssetRatingMine`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateToolboxAssetRatingMineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateToolboxAssetRatingMineResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

