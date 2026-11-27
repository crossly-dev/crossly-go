# \CreatorStudioApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCreatorCardSet**](CreatorStudioApi.md#CreateCreatorCardSet) | **Post** /v1/creator/cards/sets | Start a card set: a game id (3-24 of a-z, 0-9, -), a set code (MG1) and a name. Answers a starter draft that already passes conformance. Paid seller plans only.
[**CreateCreatorCardSetArt**](CreatorStudioApi.md#CreateCreatorCardSetArt) | **Post** /v1/creator/cards/sets/{id}/art/{slug} | One card&#39;s art: a PNG (base64, at most 1 MB and 2048 px), scanned and stored until the version is approved.
[**CreateCreatorCardSetCheck**](CreatorStudioApi.md#CreateCreatorCardSetCheck) | **Post** /v1/creator/cards/sets/{id}/check | The full check publish runs: draft issues, tcgStaticConformance and bot-vs-bot games of the set.
[**CreateCreatorCardSetPublish**](CreatorStudioApi.md#CreateCreatorCardSetPublish) | **Post** /v1/creator/cards/sets/{id}/publish | Publish: the set must pass its check; it is packaged (game.json, cards.json, the designer rules, art) and submitted as a new immutable version for review, through the same intake as a hand-made card game.
[**CreateCreatorCreationSubmit**](CreatorStudioApi.md#CreateCreatorCreationSubmit) | **Post** /v1/creator/creations/{format}/{id}/submit | Send one creation to review through its format&#39;s own submit (wearable, model, decor, interactive, card: its whole set). Games are published from world-kit.
[**CreateCreatorImport**](CreatorStudioApi.md#CreateCreatorImport) | **Post** /v1/creator/imports | Validate a filled-in template (manifest text, CSV text, every path in the upload): issues with file, line and column, and a preview. Creates nothing.
[**CreateCreatorImportCommit**](CreatorStudioApi.md#CreateCreatorImportCommit) | **Post** /v1/creator/imports/{id}/commit | Create the drafts of a validated template upload (drafts only; nothing is submitted to review).
[**CreateCreatorItemModel**](CreatorStudioApi.md#CreateCreatorItemModel) | **Post** /v1/creator/items/models | Upload a .glb (base64, 4 MB): validated, measured and stored (quarantined until approved). An over-budget model is refused with every failing number.
[**CreateCreatorItemModelCheck**](CreatorStudioApi.md#CreateCreatorItemModelCheck) | **Post** /v1/creator/items/models/check | The budget report for a .glb (base64), storing nothing: every limit (triangles, draw calls, texture edge per phone/desktop; size, nodes, meshes, images, animations; bones for wearables) with pass/fail, problems and LOD suggestions.
[**CreateCreatorItemModelSubmit**](CreatorStudioApi.md#CreateCreatorItemModelSubmit) | **Post** /v1/creator/items/models/{id}/submit | Send an uploaded model to art review.
[**CreateCreatorItemPreview**](CreatorStudioApi.md#CreateCreatorItemPreview) | **Post** /v1/creator/items/previews | Upload a preview PNG (base64, at most 1 MB and 2048 px): parsed, scanned and stored on the Crossly CDN; answers its https URL.
[**CreateCreatorItemWearable**](CreatorStudioApi.md#CreateCreatorItemWearable) | **Post** /v1/creator/items/wearables | Start a wearable from a template: slug, name, colourways (panel -&gt; #rrggbb), a COSM price and optional edition size. Text is filtered; answers the automated checks. Paid seller plans only.
[**CreateCreatorItemWearableSubmit**](CreatorStudioApi.md#CreateCreatorItemWearableSubmit) | **Post** /v1/creator/items/wearables/{id}/submit | Submit a wearable whose checks pass: each colourway becomes a COSM design (inactive) filed for review. Needs an https preview image (POST /v1/creator/items/previews).
[**CreateCreatorReviewHide**](CreatorStudioApi.md#CreateCreatorReviewHide) | **Post** /v1/creator/reviews/{reviewId}/hide | Hide a buyer review from your public page (it is kept, and its rating still counts).
[**CreateCreatorReviewReply**](CreatorStudioApi.md#CreateCreatorReviewReply) | **Post** /v1/creator/reviews/{reviewId}/reply | Reply to a buyer review of your creation (or change your reply).
[**CreateCreatorReviewReport**](CreatorStudioApi.md#CreateCreatorReviewReport) | **Post** /v1/creator/reviews/{reviewId}/report | Report a buyer review to Crossly (abuse, spam, not a real buyer).
[**CreateCreatorReviewUnhide**](CreatorStudioApi.md#CreateCreatorReviewUnhide) | **Post** /v1/creator/reviews/{reviewId}/unhide | Show a hidden buyer review on your page again.
[**DeleteCreatorCardSet**](CreatorStudioApi.md#DeleteCreatorCardSet) | **Delete** /v1/creator/cards/sets/{id} | Delete a set draft that was never published.
[**DeleteCreatorItemModel**](CreatorStudioApi.md#DeleteCreatorItemModel) | **Delete** /v1/creator/items/models/{id} | Delete a draft or rejected model.
[**DeleteCreatorItemWearable**](CreatorStudioApi.md#DeleteCreatorItemWearable) | **Delete** /v1/creator/items/wearables/{id} | Delete a draft wearable.
[**GetCreatorCardSet**](CreatorStudioApi.md#GetCreatorCardSet) | **Get** /v1/creator/cards/sets/{id} | One set draft with its check: draft issues by field and tcgStaticConformance problems.
[**GetCreatorCardSetProgress**](CreatorStudioApi.md#GetCreatorCardSetProgress) | **Get** /v1/creator/cards/sets/{id}/progress | A set&#39;s three steps: live on COSM, the mint threshold toward NFTs, and the print-run cost (trading_cards estimate) against production funding banked.
[**GetCreatorCollection**](CreatorStudioApi.md#GetCreatorCollection) | **Get** /v1/creator/collections/{collectionId} | One collection (a card set, a launch or a claimed contract): its six-stage road to the chain, the launch (phase timeline, proceeds, recent mints), sales by marketplace, royalty estimate, holders, its cards or creations with their earnings, holder perks set at mint, and which actions apply.
[**GetCreatorCollectionPerk**](CreatorStudioApi.md#GetCreatorCollectionPerk) | **Get** /v1/creator/collections/{collectionId}/perks | A collection&#39;s holder perks (percent off at your shop, early access, a World item, a loyalty tier), whether they can be edited, and why not (a card set not started yet, a contract claimed from elsewhere).
[**GetCreatorCreation**](CreatorStudioApi.md#GetCreatorCreation) | **Get** /v1/creator/creations/{format}/{id} | One creation: its status timeline and review history, the files it uses, where it sells, its sales, its buyer reviews and its mint status.
[**GetCreatorCreationStatus**](CreatorStudioApi.md#GetCreatorCreationStatus) | **Get** /v1/creator/creations/{format}/{id}/status | Where one creation stands: its status, the reviewer&#39;s last note and its status timeline.
[**GetCreatorEarningExport**](CreatorStudioApi.md#GetCreatorEarningExport) | **Get** /v1/creator/earnings/export | Download your creator earnings: format&#x3D;csv (every ledger row for the filter) or format&#x3D;pdf (a statement with per-rail totals for the same filter: rail, source, group, gameId, state; period defaults to this month).
[**GetCreatorEarningReconcile**](CreatorStudioApi.md#GetCreatorEarningReconcile) | **Get** /v1/creator/earnings/reconcile | Reconcile your creator earnings: every source recomputed directly from its own ledger, side by side with the earnings view. ok is true when every source matches.
[**GetCreatorImport**](CreatorStudioApi.md#GetCreatorImport) | **Get** /v1/creator/imports/{id} | A stored template validation.
[**GetCreatorItemModel**](CreatorStudioApi.md#GetCreatorItemModel) | **Get** /v1/creator/items/models/{id} | One of your models with its budget report.
[**GetCreatorItemTemplate**](CreatorStudioApi.md#GetCreatorItemTemplate) | **Get** /v1/creator/items/templates | Wearable templates: each garment the Crossly avatar wears, its recolourable panels, default colours and measured cost (triangles, draw calls, bones, bytes, phone LOD), plus the per-wearable budget.
[**GetCreatorItemWearable**](CreatorStudioApi.md#GetCreatorItemWearable) | **Get** /v1/creator/items/wearables/{id} | One of your wearables with its checks, review state and note.
[**GetCreatorMinting**](CreatorStudioApi.md#GetCreatorMinting) | **Get** /v1/creator/minting | Everything you take on chain as collection rows with a six-stage status (designed, eligible, drafted, deploying, live, ended): gettingReady and launched. Game items (source game_item: tradeable game products, NFT declarations for your games) carry href and gameId.
[**GetCreatorOverview**](CreatorStudioApi.md#GetCreatorOverview) | **Get** /v1/creator/overview | Creator Studio overview: earnings this month and last by rail, live creations, review queue, publish lock and next steps.
[**GetCreatorTemplate**](CreatorStudioApi.md#GetCreatorTemplate) | **Get** /v1/creator/templates/{format} | The template files for a format (wearable, model, decor, interactive, card_set): crossly-creation.json, creations.csv where it applies, and a README.
[**GetCreatorTemplateZip**](CreatorStudioApi.md#GetCreatorTemplateZip) | **Get** /v1/creator/templates/{format}/zip | A format&#39;s template as a .zip download (the same files as GET /templates/:format, plus its empty folders).
[**ListCreatorCardSets**](CreatorStudioApi.md#ListCreatorCardSets) | **Get** /v1/creator/cards/sets | Your card set drafts, each with its road from COSM to print: live on COSM, the mint threshold (collectors, packs opened, resales) and the print-run cost against what NFT sales have banked.
[**ListCreatorCreationReviews**](CreatorStudioApi.md#ListCreatorCreationReviews) | **Get** /v1/creator/creations/{format}/{id}/reviews | Buyer reviews of one creation, newest first (hidden ones included).
[**ListCreatorCreations**](CreatorStudioApi.md#ListCreatorCreations) | **Get** /v1/creator/creations | Every creation of your shop in one list: wearables, 3D objects, decor, interactive objects, each card of your card sets, Classic games and card games, with status, collection, rating, earnings and mint status.
[**ListCreatorEarningTrend**](CreatorStudioApi.md#ListCreatorEarningTrend) | **Get** /v1/creator/earnings/trend | Your creator earnings over time: net (available + paid) per day, week or month (UTC) and rail, for the same filter as the ledger. meta.totals holds the per-rail and per-source totals of that filter.
[**ListCreatorEarnings**](CreatorStudioApi.md#ListCreatorEarnings) | **Get** /v1/creator/earnings | Your creator earnings ledger across every rail (cash, COSM, CBX), newest first. Each row names the source ledger row it is. meta.totals holds per-rail and per-source totals for the same filter; meta.nextCursor pages.
[**ListCreatorItemModels**](CreatorStudioApi.md#ListCreatorItemModels) | **Get** /v1/creator/items/models | Your uploaded 3D models with their budget reports and review state.
[**ListCreatorItemWearables**](CreatorStudioApi.md#ListCreatorItemWearables) | **Get** /v1/creator/items/wearables | Your wearables: drafts, in review, live or with changes requested, each with its automated checks and per-colourway designs.
[**ListCreatorItems**](CreatorStudioApi.md#ListCreatorItems) | **Get** /v1/creator/items | Every creation of your shop — wearables, decor, uploaded 3D models, interactive objects, Classic games, card games (incl. unpublished designer sets), physical runs — with status, price and earnings.
[**ListCreatorReviews**](CreatorStudioApi.md#ListCreatorReviews) | **Get** /v1/creator/reviews | Every creator submission (items, interactive objects, Classic games, card games) and its review decision with the reviewer note.
[**UpdateCreatorCardSet**](CreatorStudioApi.md#UpdateCreatorCardSet) | **Put** /v1/creator/cards/sets/{id} | Save the whole draft: cards (name, type, rarity, rules text, flavour, power/draw/guard), COSM prices and supply, frame colours, print plan. Every word goes through the text filter. Art is merged: a card without &#x60;art&#x60; (or with art as read) keeps its stored art, &#x60;art: null&#x60; clears it; only the art upload sets art. Answers the check.
[**UpdateCreatorCollectionPerk**](CreatorStudioApi.md#UpdateCreatorCollectionPerk) | **Put** /v1/creator/collections/{collectionId}/perks | Replace a collection&#39;s holder perks: &#x60;{ perks: [...] }&#x60;, at most one of each type: percent_off { percent 1-50 }, early_access { hours 1-720 }, world_item { creationRef, name }, loyalty_tier { tierId, tierName } (one of your loyalty tiers), each with an optional note. 409 when the collection has no launch yet.
[**UpdateCreatorItemWearable**](CreatorStudioApi.md#UpdateCreatorItemWearable) | **Patch** /v1/creator/items/wearables/{id} | Edit a draft wearable (name, description, colourways, price, edition). Submitted wearables are frozen.



## CreateCreatorCardSet

> CreatorCardSetResponse CreateCreatorCardSet(ctx).Execute()

Start a card set: a game id (3-24 of a-z, 0-9, -), a set code (MG1) and a name. Answers a starter draft that already passes conformance. Paid seller plans only.

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
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorCardSet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorCardSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorCardSet`: CreatorCardSetResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorCardSet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorCardSetRequest struct via the builder pattern


### Return type

**CreatorCardSetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorCardSetArt

> CreatorCardSetResponse CreateCreatorCardSetArt(ctx, id, slug).Execute()

One card's art: a PNG (base64, at most 1 MB and 2048 px), scanned and stored until the version is approved.

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
    slug := "slug_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorCardSetArt(context.Background(), id, slug).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorCardSetArt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorCardSetArt`: CreatorCardSetResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorCardSetArt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**slug** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorCardSetArtRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreatorCardSetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorCardSetCheck

> CreatorCardSetResponse CreateCreatorCardSetCheck(ctx, id).Execute()

The full check publish runs: draft issues, tcgStaticConformance and bot-vs-bot games of the set.

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
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorCardSetCheck(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorCardSetCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorCardSetCheck`: CreatorCardSetResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorCardSetCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorCardSetCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatorCardSetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorCardSetPublish

> CreateCreatorCardSetPublishResponse CreateCreatorCardSetPublish(ctx, id).Execute()

Publish: the set must pass its check; it is packaged (game.json, cards.json, the designer rules, art) and submitted as a new immutable version for review, through the same intake as a hand-made card game.

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
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorCardSetPublish(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorCardSetPublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorCardSetPublish`: CreateCreatorCardSetPublishResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorCardSetPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorCardSetPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateCreatorCardSetPublishResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorCreationSubmit

> CreateCreatorCreationSubmitResponse CreateCreatorCreationSubmit(ctx, format, id).Execute()

Send one creation to review through its format's own submit (wearable, model, decor, interactive, card: its whole set). Games are published from world-kit.

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
    format := "format_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorCreationSubmit(context.Background(), format, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorCreationSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorCreationSubmit`: CreateCreatorCreationSubmitResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorCreationSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorCreationSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateCreatorCreationSubmitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorImport

> GetCreatorImportResponse CreateCreatorImport(ctx).Execute()

Validate a filled-in template (manifest text, CSV text, every path in the upload): issues with file, line and column, and a preview. Creates nothing.

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
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorImport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorImport`: GetCreatorImportResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorImport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorImportRequest struct via the builder pattern


### Return type

**GetCreatorImportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorImportCommit

> CreateCreatorImportCommitResponse CreateCreatorImportCommit(ctx, id).Execute()

Create the drafts of a validated template upload (drafts only; nothing is submitted to review).

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
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorImportCommit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorImportCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorImportCommit`: CreateCreatorImportCommitResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorImportCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorImportCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateCreatorImportCommitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorItemModel

> CreatorItemModelResponse CreateCreatorItemModel(ctx).Execute()

Upload a .glb (base64, 4 MB): validated, measured and stored (quarantined until approved). An over-budget model is refused with every failing number.

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
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorItemModel(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorItemModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorItemModel`: CreatorItemModelResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorItemModel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorItemModelRequest struct via the builder pattern


### Return type

**CreatorItemModelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorItemModelCheck

> CreateCreatorItemModelCheckResponse CreateCreatorItemModelCheck(ctx).Execute()

The budget report for a .glb (base64), storing nothing: every limit (triangles, draw calls, texture edge per phone/desktop; size, nodes, meshes, images, animations; bones for wearables) with pass/fail, problems and LOD suggestions.

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
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorItemModelCheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorItemModelCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorItemModelCheck`: CreateCreatorItemModelCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorItemModelCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorItemModelCheckRequest struct via the builder pattern


### Return type

**CreateCreatorItemModelCheckResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorItemModelSubmit

> CreatorItemModelResponse CreateCreatorItemModelSubmit(ctx, id).Execute()

Send an uploaded model to art review.

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
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorItemModelSubmit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorItemModelSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorItemModelSubmit`: CreatorItemModelResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorItemModelSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorItemModelSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatorItemModelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorItemPreview

> CreateCreatorItemPreviewResponse CreateCreatorItemPreview(ctx).Execute()

Upload a preview PNG (base64, at most 1 MB and 2048 px): parsed, scanned and stored on the Crossly CDN; answers its https URL.

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
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorItemPreview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorItemPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorItemPreview`: CreateCreatorItemPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorItemPreview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorItemPreviewRequest struct via the builder pattern


### Return type

**CreateCreatorItemPreviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorItemWearable

> CreatorItemWearableResponse CreateCreatorItemWearable(ctx).Execute()

Start a wearable from a template: slug, name, colourways (panel -> #rrggbb), a COSM price and optional edition size. Text is filtered; answers the automated checks. Paid seller plans only.

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
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorItemWearable(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorItemWearable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorItemWearable`: CreatorItemWearableResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorItemWearable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorItemWearableRequest struct via the builder pattern


### Return type

**CreatorItemWearableResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorItemWearableSubmit

> CreatorItemWearableResponse CreateCreatorItemWearableSubmit(ctx, id).Execute()

Submit a wearable whose checks pass: each colourway becomes a COSM design (inactive) filed for review. Needs an https preview image (POST /v1/creator/items/previews).

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
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorItemWearableSubmit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorItemWearableSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorItemWearableSubmit`: CreatorItemWearableResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorItemWearableSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorItemWearableSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatorItemWearableResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorReviewHide

> CreatorReviewResponse CreateCreatorReviewHide(ctx, reviewId).Execute()

Hide a buyer review from your public page (it is kept, and its rating still counts).

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
    reviewId := "reviewId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorReviewHide(context.Background(), reviewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorReviewHide``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorReviewHide`: CreatorReviewResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorReviewHide`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reviewId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorReviewHideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatorReviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorReviewReply

> CreatorReviewResponse CreateCreatorReviewReply(ctx, reviewId).Execute()

Reply to a buyer review of your creation (or change your reply).

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
    reviewId := "reviewId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorReviewReply(context.Background(), reviewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorReviewReply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorReviewReply`: CreatorReviewResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorReviewReply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reviewId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorReviewReplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatorReviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorReviewReport

> CreatorReviewResponse CreateCreatorReviewReport(ctx, reviewId).Execute()

Report a buyer review to Crossly (abuse, spam, not a real buyer).

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
    reviewId := "reviewId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorReviewReport(context.Background(), reviewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorReviewReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorReviewReport`: CreatorReviewResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorReviewReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reviewId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorReviewReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatorReviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorReviewUnhide

> CreatorReviewResponse CreateCreatorReviewUnhide(ctx, reviewId).Execute()

Show a hidden buyer review on your page again.

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
    reviewId := "reviewId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.CreateCreatorReviewUnhide(context.Background(), reviewId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.CreateCreatorReviewUnhide``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorReviewUnhide`: CreatorReviewResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.CreateCreatorReviewUnhide`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**reviewId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorReviewUnhideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatorReviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteCreatorCardSet

> DeleteCreatorCardSet(ctx, id).Execute()

Delete a set draft that was never published.

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
    resp, r, err := apiClient.CreatorStudioApi.DeleteCreatorCardSet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.DeleteCreatorCardSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCreatorCardSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteCreatorItemModel

> DeleteCreatorItemModel(ctx, id).Execute()

Delete a draft or rejected model.

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
    resp, r, err := apiClient.CreatorStudioApi.DeleteCreatorItemModel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.DeleteCreatorItemModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCreatorItemModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteCreatorItemWearable

> DeleteCreatorItemWearable(ctx, id).Execute()

Delete a draft wearable.

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
    resp, r, err := apiClient.CreatorStudioApi.DeleteCreatorItemWearable(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.DeleteCreatorItemWearable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCreatorItemWearableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorCardSet

> CreatorCardSetResponse GetCreatorCardSet(ctx, id).Execute()

One set draft with its check: draft issues by field and tcgStaticConformance problems.

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
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorCardSet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorCardSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorCardSet`: CreatorCardSetResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorCardSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorCardSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatorCardSetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorCardSetProgress

> GetCreatorCardSetProgressResponse GetCreatorCardSetProgress(ctx, id).Execute()

A set's three steps: live on COSM, the mint threshold toward NFTs, and the print-run cost (trading_cards estimate) against production funding banked.

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
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorCardSetProgress(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorCardSetProgress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorCardSetProgress`: GetCreatorCardSetProgressResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorCardSetProgress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorCardSetProgressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetCreatorCardSetProgressResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorCollection

> GetCreatorCollectionResponse GetCreatorCollection(ctx, collectionId).Execute()

One collection (a card set, a launch or a claimed contract): its six-stage road to the chain, the launch (phase timeline, proceeds, recent mints), sales by marketplace, royalty estimate, holders, its cards or creations with their earnings, holder perks set at mint, and which actions apply.

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
    collectionId := "collectionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorCollection(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorCollection`: GetCreatorCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetCreatorCollectionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorCollectionPerk

> GetCreatorCollectionPerkResponse GetCreatorCollectionPerk(ctx, collectionId).Execute()

A collection's holder perks (percent off at your shop, early access, a World item, a loyalty tier), whether they can be edited, and why not (a card set not started yet, a contract claimed from elsewhere).

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
    collectionId := "collectionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorCollectionPerk(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorCollectionPerk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorCollectionPerk`: GetCreatorCollectionPerkResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorCollectionPerk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorCollectionPerkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetCreatorCollectionPerkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorCreation

> GetCreatorCreationResponse GetCreatorCreation(ctx, format, id).Execute()

One creation: its status timeline and review history, the files it uses, where it sells, its sales, its buyer reviews and its mint status.

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
    format := "format_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorCreation(context.Background(), format, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorCreation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorCreation`: GetCreatorCreationResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorCreation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorCreationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetCreatorCreationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorCreationStatus

> GetCreatorCreationStatusResponse GetCreatorCreationStatus(ctx, format, id).Execute()

Where one creation stands: its status, the reviewer's last note and its status timeline.

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
    format := "format_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorCreationStatus(context.Background(), format, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorCreationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorCreationStatus`: GetCreatorCreationStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorCreationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorCreationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetCreatorCreationStatusResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorEarningExport

> string GetCreatorEarningExport(ctx).Format(format).Rail(rail).Source(source).State(state).Group(group).GameId(gameId).From(from).To(to).Execute()

Download your creator earnings: format=csv (every ledger row for the filter) or format=pdf (a statement with per-rail totals for the same filter: rail, source, group, gameId, state; period defaults to this month).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    format := "format_example" // string |  (optional)
    rail := "rail_example" // string |  (optional)
    source := "source_example" // string |  (optional)
    state := "state_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    gameId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    from := time.Now() // time.Time |  (optional)
    to := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorEarningExport(context.Background()).Format(format).Rail(rail).Source(source).State(state).Group(group).GameId(gameId).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorEarningExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorEarningExport`: string
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorEarningExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorEarningExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **format** | **string** |  | 
 **rail** | **string** |  | 
 **source** | **string** |  | 
 **state** | **string** |  | 
 **group** | **string** |  | 
 **gameId** | **string** |  | 
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 

### Return type

**string**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/csv, application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorEarningReconcile

> GetCreatorEarningReconcileResponse GetCreatorEarningReconcile(ctx).Execute()

Reconcile your creator earnings: every source recomputed directly from its own ledger, side by side with the earnings view. ok is true when every source matches.

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
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorEarningReconcile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorEarningReconcile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorEarningReconcile`: GetCreatorEarningReconcileResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorEarningReconcile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorEarningReconcileRequest struct via the builder pattern


### Return type

**GetCreatorEarningReconcileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorImport

> GetCreatorImportResponse GetCreatorImport(ctx, id).Execute()

A stored template validation.

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
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorImport(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorImport`: GetCreatorImportResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorImport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetCreatorImportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorItemModel

> CreatorItemModelResponse GetCreatorItemModel(ctx, id).Execute()

One of your models with its budget report.

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
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorItemModel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorItemModel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorItemModel`: CreatorItemModelResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorItemModel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorItemModelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatorItemModelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorItemTemplate

> GetCreatorItemTemplateResponse GetCreatorItemTemplate(ctx).Execute()

Wearable templates: each garment the Crossly avatar wears, its recolourable panels, default colours and measured cost (triangles, draw calls, bones, bytes, phone LOD), plus the per-wearable budget.

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
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorItemTemplate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorItemTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorItemTemplate`: GetCreatorItemTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorItemTemplate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorItemTemplateRequest struct via the builder pattern


### Return type

**GetCreatorItemTemplateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorItemWearable

> CreatorItemWearableResponse GetCreatorItemWearable(ctx, id).Execute()

One of your wearables with its checks, review state and note.

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
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorItemWearable(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorItemWearable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorItemWearable`: CreatorItemWearableResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorItemWearable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorItemWearableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatorItemWearableResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorMinting

> GetCreatorMintingResponse GetCreatorMinting(ctx).Execute()

Everything you take on chain as collection rows with a six-stage status (designed, eligible, drafted, deploying, live, ended): gettingReady and launched. Game items (source game_item: tradeable game products, NFT declarations for your games) carry href and gameId.

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
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorMinting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorMinting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorMinting`: GetCreatorMintingResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorMinting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorMintingRequest struct via the builder pattern


### Return type

**GetCreatorMintingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorOverview

> GetCreatorOverviewResponse GetCreatorOverview(ctx).Execute()

Creator Studio overview: earnings this month and last by rail, live creations, review queue, publish lock and next steps.

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
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorOverview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorOverview`: GetCreatorOverviewResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorOverview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorOverviewRequest struct via the builder pattern


### Return type

**GetCreatorOverviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorTemplate

> GetCreatorTemplateResponse GetCreatorTemplate(ctx, format).Execute()

The template files for a format (wearable, model, decor, interactive, card_set): crossly-creation.json, creations.csv where it applies, and a README.

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
    format := "format_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorTemplate(context.Background(), format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorTemplate`: GetCreatorTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetCreatorTemplateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorTemplateZip

> *os.File GetCreatorTemplateZip(ctx, format).Execute()

A format's template as a .zip download (the same files as GET /templates/:format, plus its empty folders).

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
    format := "format_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.GetCreatorTemplateZip(context.Background(), format).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.GetCreatorTemplateZip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorTemplateZip`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.GetCreatorTemplateZip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorTemplateZipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

***os.File**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCreatorCardSets

> ListCreatorCardSetsResponse ListCreatorCardSets(ctx).Execute()

Your card set drafts, each with its road from COSM to print: live on COSM, the mint threshold (collectors, packs opened, resales) and the print-run cost against what NFT sales have banked.

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
    resp, r, err := apiClient.CreatorStudioApi.ListCreatorCardSets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.ListCreatorCardSets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreatorCardSets`: ListCreatorCardSetsResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.ListCreatorCardSets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCreatorCardSetsRequest struct via the builder pattern


### Return type

**ListCreatorCardSetsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCreatorCreationReviews

> ListCreatorCreationReviewsResponse ListCreatorCreationReviews(ctx, format, id).Execute()

Buyer reviews of one creation, newest first (hidden ones included).

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
    format := "format_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.ListCreatorCreationReviews(context.Background(), format, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.ListCreatorCreationReviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreatorCreationReviews`: ListCreatorCreationReviewsResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.ListCreatorCreationReviews`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**format** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCreatorCreationReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**ListCreatorCreationReviewsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCreatorCreations

> ListCreatorCreationsResponse ListCreatorCreations(ctx).Execute()

Every creation of your shop in one list: wearables, 3D objects, decor, interactive objects, each card of your card sets, Classic games and card games, with status, collection, rating, earnings and mint status.

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
    resp, r, err := apiClient.CreatorStudioApi.ListCreatorCreations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.ListCreatorCreations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreatorCreations`: ListCreatorCreationsResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.ListCreatorCreations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCreatorCreationsRequest struct via the builder pattern


### Return type

**ListCreatorCreationsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCreatorEarningTrend

> ListCreatorEarningTrendResponse ListCreatorEarningTrend(ctx).Bucket(bucket).Rail(rail).Group(group).GameId(gameId).From(from).To(to).Execute()

Your creator earnings over time: net (available + paid) per day, week or month (UTC) and rail, for the same filter as the ledger. meta.totals holds the per-rail and per-source totals of that filter.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    bucket := "bucket_example" // string |  (optional)
    rail := "rail_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    gameId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    from := time.Now() // time.Time |  (optional)
    to := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.ListCreatorEarningTrend(context.Background()).Bucket(bucket).Rail(rail).Group(group).GameId(gameId).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.ListCreatorEarningTrend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreatorEarningTrend`: ListCreatorEarningTrendResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.ListCreatorEarningTrend`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCreatorEarningTrendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bucket** | **string** |  | 
 **rail** | **string** |  | 
 **group** | **string** |  | 
 **gameId** | **string** |  | 
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 

### Return type

**ListCreatorEarningTrendResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCreatorEarnings

> ListCreatorEarningsResponse ListCreatorEarnings(ctx).Rail(rail).Source(source).Group(group).GameId(gameId).State(state).From(from).To(to).Cursor(cursor).Limit(limit).Execute()

Your creator earnings ledger across every rail (cash, COSM, CBX), newest first. Each row names the source ledger row it is. meta.totals holds per-rail and per-source totals for the same filter; meta.nextCursor pages.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    rail := "rail_example" // string |  (optional)
    source := "source_example" // string |  (optional)
    group := "group_example" // string |  (optional)
    gameId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    state := "state_example" // string |  (optional)
    from := time.Now() // time.Time |  (optional)
    to := time.Now() // time.Time |  (optional)
    cursor := "cursor_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.ListCreatorEarnings(context.Background()).Rail(rail).Source(source).Group(group).GameId(gameId).State(state).From(from).To(to).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.ListCreatorEarnings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreatorEarnings`: ListCreatorEarningsResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.ListCreatorEarnings`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCreatorEarningsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **rail** | **string** |  | 
 **source** | **string** |  | 
 **group** | **string** |  | 
 **gameId** | **string** |  | 
 **state** | **string** |  | 
 **from** | **time.Time** |  | 
 **to** | **time.Time** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

**ListCreatorEarningsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCreatorItemModels

> ListCreatorItemModelsResponse ListCreatorItemModels(ctx).Execute()

Your uploaded 3D models with their budget reports and review state.

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
    resp, r, err := apiClient.CreatorStudioApi.ListCreatorItemModels(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.ListCreatorItemModels``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreatorItemModels`: ListCreatorItemModelsResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.ListCreatorItemModels`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCreatorItemModelsRequest struct via the builder pattern


### Return type

**ListCreatorItemModelsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCreatorItemWearables

> ListCreatorItemWearablesResponse ListCreatorItemWearables(ctx).Execute()

Your wearables: drafts, in review, live or with changes requested, each with its automated checks and per-colourway designs.

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
    resp, r, err := apiClient.CreatorStudioApi.ListCreatorItemWearables(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.ListCreatorItemWearables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreatorItemWearables`: ListCreatorItemWearablesResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.ListCreatorItemWearables`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCreatorItemWearablesRequest struct via the builder pattern


### Return type

**ListCreatorItemWearablesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCreatorItems

> ListCreatorItemsResponse ListCreatorItems(ctx).Kind(kind).Status(status).Rail(rail).Execute()

Every creation of your shop — wearables, decor, uploaded 3D models, interactive objects, Classic games, card games (incl. unpublished designer sets), physical runs — with status, price and earnings.

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
    kind := "kind_example" // string |  (optional)
    status := "status_example" // string |  (optional)
    rail := "rail_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.ListCreatorItems(context.Background()).Kind(kind).Status(status).Rail(rail).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.ListCreatorItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreatorItems`: ListCreatorItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.ListCreatorItems`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCreatorItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** |  | 
 **status** | **string** |  | 
 **rail** | **string** |  | 

### Return type

**ListCreatorItemsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCreatorReviews

> ListCreatorReviewsResponse ListCreatorReviews(ctx).Kind(kind).State(state).Limit(limit).Execute()

Every creator submission (items, interactive objects, Classic games, card games) and its review decision with the reviewer note.

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
    kind := "kind_example" // string |  (optional)
    state := "state_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.ListCreatorReviews(context.Background()).Kind(kind).State(state).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.ListCreatorReviews``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreatorReviews`: ListCreatorReviewsResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.ListCreatorReviews`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCreatorReviewsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** |  | 
 **state** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

**ListCreatorReviewsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateCreatorCardSet

> CreatorCardSetResponse UpdateCreatorCardSet(ctx, id).Execute()

Save the whole draft: cards (name, type, rarity, rules text, flavour, power/draw/guard), COSM prices and supply, frame colours, print plan. Every word goes through the text filter. Art is merged: a card without `art` (or with art as read) keeps its stored art, `art: null` clears it; only the art upload sets art. Answers the check.

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
    resp, r, err := apiClient.CreatorStudioApi.UpdateCreatorCardSet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.UpdateCreatorCardSet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCreatorCardSet`: CreatorCardSetResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.UpdateCreatorCardSet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCreatorCardSetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatorCardSetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateCreatorCollectionPerk

> GetCreatorCollectionPerkResponse UpdateCreatorCollectionPerk(ctx, collectionId).Execute()

Replace a collection's holder perks: `{ perks: [...] }`, at most one of each type: percent_off { percent 1-50 }, early_access { hours 1-720 }, world_item { creationRef, name }, loyalty_tier { tierId, tierName } (one of your loyalty tiers), each with an optional note. 409 when the collection has no launch yet.

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
    collectionId := "collectionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorStudioApi.UpdateCreatorCollectionPerk(context.Background(), collectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.UpdateCreatorCollectionPerk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCreatorCollectionPerk`: GetCreatorCollectionPerkResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.UpdateCreatorCollectionPerk`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**collectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCreatorCollectionPerkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetCreatorCollectionPerkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateCreatorItemWearable

> CreatorItemWearableResponse UpdateCreatorItemWearable(ctx, id).Execute()

Edit a draft wearable (name, description, colourways, price, edition). Submitted wearables are frozen.

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
    resp, r, err := apiClient.CreatorStudioApi.UpdateCreatorItemWearable(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorStudioApi.UpdateCreatorItemWearable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCreatorItemWearable`: CreatorItemWearableResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorStudioApi.UpdateCreatorItemWearable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCreatorItemWearableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatorItemWearableResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

