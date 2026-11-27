# \StorefrontApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStorefrontHighlight**](StorefrontApi.md#CreateStorefrontHighlight) | **Post** /v1/storefront/highlights | Create a highlight: title (up to 24 characters), optional coverUrl, and optional storyIds (your live or archived stories) to save into it. Up to 50 per store.
[**CreateStorefrontHighlightItem**](StorefrontApi.md#CreateStorefrontHighlightItem) | **Post** /v1/storefront/highlights/{id}/items | Save stories into a highlight (appended in the order given). Each is copied, so it stays after the story expires. A story already in the highlight is skipped.
[**CreateStorefrontProfileSuggestShelve**](StorefrontApi.md#CreateStorefrontProfileSuggestShelve) | **Post** /v1/storefront/profile/suggest-shelves | Profile shelves proposed from your top tags and categories (not saved).
[**CreateStorefrontSection**](StorefrontApi.md#CreateStorefrontSection) | **Post** /v1/storefront/sections | Add a section at the bottom: kind hero | listings | collections | about | policies | social | announcement | faq (default listings), with its config.
[**CreateStorefrontSitePreview**](StorefrontApi.md#CreateStorefrontSitePreview) | **Post** /v1/storefront/site/preview | Resolve an unsaved design page like the live site would: every section with its listings and extras (preview: true).
[**CreateStorefrontSitePublish**](StorefrontApi.md#CreateStorefrontSitePublish) | **Post** /v1/storefront/site/publish | Publish the draft: it becomes the live site and a publish version is kept (the last 30 versions are kept).
[**CreateStorefrontSiteResolve**](StorefrontApi.md#CreateStorefrontSiteResolve) | **Post** /v1/storefront/site/resolve | The live listings a ListingSource (manual, tags, category, smart rule or all) selects from your store, and how many match.
[**CreateStorefrontSiteStarter**](StorefrontApi.md#CreateStorefrontSiteStarter) | **Post** /v1/storefront/site/starter | Replace the draft with a starter site (vintage, sneakers, streetwear, collectibles, boutique, home, minimal, drop, general). The current draft is kept as a version first.
[**CreateStorefrontSiteSuggest**](StorefrontApi.md#CreateStorefrontSiteSuggest) | **Post** /v1/storefront/site/suggest | Auto-build: website sections proposed from your top tags and categories (not saved).
[**CreateStorefrontSiteVersionRestore**](StorefrontApi.md#CreateStorefrontSiteVersionRestore) | **Post** /v1/storefront/site/versions/{id}/restore | Copy a saved version into the draft (the current draft is kept as a version first). Publish to make it live.
[**CreateStorefrontTradeSitePublish**](StorefrontApi.md#CreateStorefrontTradeSitePublish) | **Post** /v1/storefront/trade-site/publish | Publish the trade site draft (a version is kept).
[**CreateStorefrontTradeSiteResolve**](StorefrontApi.md#CreateStorefrontTradeSiteResolve) | **Post** /v1/storefront/trade-site/resolve | The listings a ListingSource selects on your trade site: only items in your active wholesale catalogs (narrowed further by source.catalogIds), and how many match.
[**CreateStorefrontTradeSiteStarter**](StorefrontApi.md#CreateStorefrontTradeSiteStarter) | **Post** /v1/storefront/trade-site/starter | Replace the trade site draft with a trade starter (trade_classic, trade_line_sheet, trade_brand_house, trade_apply_first). The current draft is kept as a version first.
[**CreateStorefrontTradeSiteSuggest**](StorefrontApi.md#CreateStorefrontTradeSiteSuggest) | **Post** /v1/storefront/trade-site/suggest | Trade Auto-build: trade site sections proposed from your wholesale catalogs (a spotlight, a rail per catalog, terms, apply banner). Not saved.
[**DeleteStorefrontHighlight**](StorefrontApi.md#DeleteStorefrontHighlight) | **Delete** /v1/storefront/highlights/{id} | Delete a highlight and its saved stories. The original stories are unaffected.
[**DeleteStorefrontHighlightItem**](StorefrontApi.md#DeleteStorefrontHighlightItem) | **Delete** /v1/storefront/highlights/{id}/items/{itemId} | Remove one saved story from a highlight.
[**DeleteStorefrontSection**](StorefrontApi.md#DeleteStorefrontSection) | **Delete** /v1/storefront/sections/{id} | Remove a section. Listings on a removed shelf stay live.
[**GetStorefront**](StorefrontApi.md#GetStorefront) | **Get** /v1/storefront | Your storefront design: banner, logo, tagline, bio, theme, SEO title and description, social links, and whether it is published.
[**GetStorefrontMerch**](StorefrontApi.md#GetStorefrontMerch) | **Get** /v1/storefront/merch | Your store&#39;s live listings by tag, category, condition and brand, with the price range: what product sources can select.
[**GetStorefrontProfile**](StorefrontApi.md#GetStorefrontProfile) | **Get** /v1/storefront/profile | Your shop profile (the Crossly Shop page): banner, logo, tagline, bio, accent, announcement, shelves, blocks and All items options.
[**GetStorefrontShipping**](StorefrontApi.md#GetStorefrontShipping) | **Get** /v1/storefront/shipping | Your Crossly shipping policy (domestic and international services, which countries you ship to) and your default ship-from address.
[**GetStorefrontSite**](StorefrontApi.md#GetStorefrontSite) | **Get** /v1/storefront/site | Your store website: the draft design, what is live, when each changed, whether the site is switched on, and its custom domain.
[**GetStorefrontTradeSite**](StorefrontApi.md#GetStorefrontTradeSite) | **Get** /v1/storefront/trade-site | Your store&#39;s trade (wholesale) site at /site/:slug/trade: the draft, what is live, whether it is on and its gating (public | prices_private | approved_only).
[**ListStorefrontHighlights**](StorefrontApi.md#ListStorefrontHighlights) | **Get** /v1/storefront/highlights | Your store’s story highlights in display order, each with its saved stories (media, overlays, product tie-in, when posted). Empty highlights included.
[**ListStorefrontSections**](StorefrontApi.md#ListStorefrontSections) | **Get** /v1/storefront/sections | Your storefront sections in display order: kind, title, visible, config, and listing ids for listings shelves.
[**ListStorefrontSiteVersions**](StorefrontApi.md#ListStorefrontSiteVersions) | **Get** /v1/storefront/site/versions | Saved versions of the website, newest first (publish, restore and template points; up to 30).
[**ListStorefrontStoryArchive**](StorefrontApi.md#ListStorefrontStoryArchive) | **Get** /v1/storefront/stories/archive | Your stories, live and expired, newest first: what can be saved to a highlight. Paginated: pass before &#x3D; meta.nextCursor of the previous page.
[**UpdateStorefront**](StorefrontApi.md#UpdateStorefront) | **Put** /v1/storefront | Update your storefront design. Omitted fields are unchanged, null clears. published&#x3D;false keeps the design as a draft and buyers see the standard storefront.
[**UpdateStorefrontHighlight**](StorefrontApi.md#UpdateStorefrontHighlight) | **Patch** /v1/storefront/highlights/{id} | Rename a highlight or change its cover (coverUrl null &#x3D; use the first story).
[**UpdateStorefrontHighlightItemOrder**](StorefrontApi.md#UpdateStorefrontHighlightItemOrder) | **Put** /v1/storefront/highlights/{id}/items/order | Set the order of a highlight’s saved stories: every item id, first to last.
[**UpdateStorefrontHighlightOrder**](StorefrontApi.md#UpdateStorefrontHighlightOrder) | **Put** /v1/storefront/highlights/order | Set the highlight order in one call: every highlight id, first to last.
[**UpdateStorefrontProfile**](StorefrontApi.md#UpdateStorefrontProfile) | **Put** /v1/storefront/profile | Replace your shop profile (a complete ShopProfile). Saved straight away.
[**UpdateStorefrontSection**](StorefrontApi.md#UpdateStorefrontSection) | **Patch** /v1/storefront/sections/{id} | Rename, show/hide, refill (listings shelves) or reconfigure a section.
[**UpdateStorefrontSectionReorder**](StorefrontApi.md#UpdateStorefrontSectionReorder) | **Put** /v1/storefront/sections/reorder | Set the section order in one atomic call: every section id, top to bottom.
[**UpdateStorefrontShipping**](StorefrontApi.md#UpdateStorefrontShipping) | **Put** /v1/storefront/shipping | Set your shipping policy (international.enabled, shipsTo worldwide or ISO codes, excludes) and/or your default ship-from address (any country; 422 for a sanctioned origin).
[**UpdateStorefrontSiteDraft**](StorefrontApi.md#UpdateStorefrontSiteDraft) | **Put** /v1/storefront/site/draft | Save the website draft (a complete SiteDesign). Validated strictly (a bad field is a 400 naming its path) and sanitized; buyers see nothing until you publish.
[**UpdateStorefrontSiteEnabled**](StorefrontApi.md#UpdateStorefrontSiteEnabled) | **Put** /v1/storefront/site/enabled | Switch the website on or off. On requires a publish first; off makes /site/:slug 404 and the custom domain show the shop profile.
[**UpdateStorefrontTradeSiteDraft**](StorefrontApi.md#UpdateStorefrontTradeSiteDraft) | **Put** /v1/storefront/trade-site/draft | Save the trade site draft (a complete SiteDesign; trade sections included). Validated strictly and sanitized.
[**UpdateStorefrontTradeSiteEnabled**](StorefrontApi.md#UpdateStorefrontTradeSiteEnabled) | **Put** /v1/storefront/trade-site/enabled | Switch the trade site on (needs a publish first) or off.
[**UpdateStorefrontTradeSiteGating**](StorefrontApi.md#UpdateStorefrontTradeSiteGating) | **Put** /v1/storefront/trade-site/gating | Who sees the trade site: public (everyone, with prices), prices_private (everyone, prices for approved buyers) or approved_only (approved buyers; others get the apply page).



## CreateStorefrontHighlight

> CreateStorefrontHighlightResponse CreateStorefrontHighlight(ctx).Execute()

Create a highlight: title (up to 24 characters), optional coverUrl, and optional storyIds (your live or archived stories) to save into it. Up to 50 per store.

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
    resp, r, err := apiClient.StorefrontApi.CreateStorefrontHighlight(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.CreateStorefrontHighlight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontHighlight`: CreateStorefrontHighlightResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.CreateStorefrontHighlight`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontHighlightRequest struct via the builder pattern


### Return type

**CreateStorefrontHighlightResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontHighlightItem

> StorefrontHighlightResponse CreateStorefrontHighlightItem(ctx, id).Execute()

Save stories into a highlight (appended in the order given). Each is copied, so it stays after the story expires. A story already in the highlight is skipped.

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
    resp, r, err := apiClient.StorefrontApi.CreateStorefrontHighlightItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.CreateStorefrontHighlightItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontHighlightItem`: StorefrontHighlightResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.CreateStorefrontHighlightItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontHighlightItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontHighlightResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontProfileSuggestShelve

> CreateStorefrontProfileSuggestShelveResponse CreateStorefrontProfileSuggestShelve(ctx).Execute()

Profile shelves proposed from your top tags and categories (not saved).

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
    resp, r, err := apiClient.StorefrontApi.CreateStorefrontProfileSuggestShelve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.CreateStorefrontProfileSuggestShelve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontProfileSuggestShelve`: CreateStorefrontProfileSuggestShelveResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.CreateStorefrontProfileSuggestShelve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontProfileSuggestShelveRequest struct via the builder pattern


### Return type

**CreateStorefrontProfileSuggestShelveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontSection

> StorefrontSectionResponse CreateStorefrontSection(ctx).Execute()

Add a section at the bottom: kind hero | listings | collections | about | policies | social | announcement | faq (default listings), with its config.

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
    resp, r, err := apiClient.StorefrontApi.CreateStorefrontSection(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.CreateStorefrontSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontSection`: StorefrontSectionResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.CreateStorefrontSection`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontSectionRequest struct via the builder pattern


### Return type

**StorefrontSectionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontSitePreview

> CreateStorefrontSitePreviewResponse CreateStorefrontSitePreview(ctx).Execute()

Resolve an unsaved design page like the live site would: every section with its listings and extras (preview: true).

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
    resp, r, err := apiClient.StorefrontApi.CreateStorefrontSitePreview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.CreateStorefrontSitePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontSitePreview`: CreateStorefrontSitePreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.CreateStorefrontSitePreview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontSitePreviewRequest struct via the builder pattern


### Return type

**CreateStorefrontSitePreviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontSitePublish

> StorefrontResponse2 CreateStorefrontSitePublish(ctx).Execute()

Publish the draft: it becomes the live site and a publish version is kept (the last 30 versions are kept).

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
    resp, r, err := apiClient.StorefrontApi.CreateStorefrontSitePublish(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.CreateStorefrontSitePublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontSitePublish`: StorefrontResponse2
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.CreateStorefrontSitePublish`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontSitePublishRequest struct via the builder pattern


### Return type

**StorefrontResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontSiteResolve

> CreateStorefrontSiteResolveResponse CreateStorefrontSiteResolve(ctx).Execute()

The live listings a ListingSource (manual, tags, category, smart rule or all) selects from your store, and how many match.

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
    resp, r, err := apiClient.StorefrontApi.CreateStorefrontSiteResolve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.CreateStorefrontSiteResolve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontSiteResolve`: CreateStorefrontSiteResolveResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.CreateStorefrontSiteResolve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontSiteResolveRequest struct via the builder pattern


### Return type

**CreateStorefrontSiteResolveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontSiteStarter

> StorefrontResponse2 CreateStorefrontSiteStarter(ctx).Execute()

Replace the draft with a starter site (vintage, sneakers, streetwear, collectibles, boutique, home, minimal, drop, general). The current draft is kept as a version first.

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
    resp, r, err := apiClient.StorefrontApi.CreateStorefrontSiteStarter(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.CreateStorefrontSiteStarter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontSiteStarter`: StorefrontResponse2
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.CreateStorefrontSiteStarter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontSiteStarterRequest struct via the builder pattern


### Return type

**StorefrontResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontSiteSuggest

> CreateStorefrontSiteSuggestResponse CreateStorefrontSiteSuggest(ctx).Execute()

Auto-build: website sections proposed from your top tags and categories (not saved).

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
    resp, r, err := apiClient.StorefrontApi.CreateStorefrontSiteSuggest(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.CreateStorefrontSiteSuggest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontSiteSuggest`: CreateStorefrontSiteSuggestResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.CreateStorefrontSiteSuggest`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontSiteSuggestRequest struct via the builder pattern


### Return type

**CreateStorefrontSiteSuggestResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontSiteVersionRestore

> StorefrontResponse3 CreateStorefrontSiteVersionRestore(ctx, id).Execute()

Copy a saved version into the draft (the current draft is kept as a version first). Publish to make it live.

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
    resp, r, err := apiClient.StorefrontApi.CreateStorefrontSiteVersionRestore(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.CreateStorefrontSiteVersionRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontSiteVersionRestore`: StorefrontResponse3
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.CreateStorefrontSiteVersionRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontSiteVersionRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontResponse3**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontTradeSitePublish

> StorefrontResponse2 CreateStorefrontTradeSitePublish(ctx).Execute()

Publish the trade site draft (a version is kept).

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
    resp, r, err := apiClient.StorefrontApi.CreateStorefrontTradeSitePublish(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.CreateStorefrontTradeSitePublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontTradeSitePublish`: StorefrontResponse2
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.CreateStorefrontTradeSitePublish`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontTradeSitePublishRequest struct via the builder pattern


### Return type

**StorefrontResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontTradeSiteResolve

> CreateStorefrontSiteResolveResponse CreateStorefrontTradeSiteResolve(ctx).Execute()

The listings a ListingSource selects on your trade site: only items in your active wholesale catalogs (narrowed further by source.catalogIds), and how many match.

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
    resp, r, err := apiClient.StorefrontApi.CreateStorefrontTradeSiteResolve(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.CreateStorefrontTradeSiteResolve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontTradeSiteResolve`: CreateStorefrontSiteResolveResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.CreateStorefrontTradeSiteResolve`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontTradeSiteResolveRequest struct via the builder pattern


### Return type

**CreateStorefrontSiteResolveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontTradeSiteStarter

> StorefrontResponse2 CreateStorefrontTradeSiteStarter(ctx).Execute()

Replace the trade site draft with a trade starter (trade_classic, trade_line_sheet, trade_brand_house, trade_apply_first). The current draft is kept as a version first.

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
    resp, r, err := apiClient.StorefrontApi.CreateStorefrontTradeSiteStarter(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.CreateStorefrontTradeSiteStarter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontTradeSiteStarter`: StorefrontResponse2
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.CreateStorefrontTradeSiteStarter`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontTradeSiteStarterRequest struct via the builder pattern


### Return type

**StorefrontResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontTradeSiteSuggest

> CreateStorefrontSiteSuggestResponse CreateStorefrontTradeSiteSuggest(ctx).Execute()

Trade Auto-build: trade site sections proposed from your wholesale catalogs (a spotlight, a rail per catalog, terms, apply banner). Not saved.

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
    resp, r, err := apiClient.StorefrontApi.CreateStorefrontTradeSiteSuggest(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.CreateStorefrontTradeSiteSuggest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontTradeSiteSuggest`: CreateStorefrontSiteSuggestResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.CreateStorefrontTradeSiteSuggest`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontTradeSiteSuggestRequest struct via the builder pattern


### Return type

**CreateStorefrontSiteSuggestResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteStorefrontHighlight

> DeleteStoryResponse DeleteStorefrontHighlight(ctx, id).Execute()

Delete a highlight and its saved stories. The original stories are unaffected.

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
    resp, r, err := apiClient.StorefrontApi.DeleteStorefrontHighlight(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.DeleteStorefrontHighlight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStorefrontHighlight`: DeleteStoryResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.DeleteStorefrontHighlight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorefrontHighlightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteStoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteStorefrontHighlightItem

> StorefrontHighlightResponse DeleteStorefrontHighlightItem(ctx, id, itemId).Execute()

Remove one saved story from a highlight.

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
    itemId := "itemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorefrontApi.DeleteStorefrontHighlightItem(context.Background(), id, itemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.DeleteStorefrontHighlightItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStorefrontHighlightItem`: StorefrontHighlightResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.DeleteStorefrontHighlightItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**itemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorefrontHighlightItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**StorefrontHighlightResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteStorefrontSection

> UpdateGroupBuyResponse DeleteStorefrontSection(ctx, id).Execute()

Remove a section. Listings on a removed shelf stay live.

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
    resp, r, err := apiClient.StorefrontApi.DeleteStorefrontSection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.DeleteStorefrontSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStorefrontSection`: UpdateGroupBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.DeleteStorefrontSection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorefrontSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateGroupBuyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetStorefront

> GetStorefrontResponse GetStorefront(ctx).Execute()

Your storefront design: banner, logo, tagline, bio, theme, SEO title and description, social links, and whether it is published.

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
    resp, r, err := apiClient.StorefrontApi.GetStorefront(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.GetStorefront``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorefront`: GetStorefrontResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.GetStorefront`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorefrontRequest struct via the builder pattern


### Return type

**GetStorefrontResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetStorefrontMerch

> GetStorefrontMerchResponse GetStorefrontMerch(ctx).Execute()

Your store's live listings by tag, category, condition and brand, with the price range: what product sources can select.

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
    resp, r, err := apiClient.StorefrontApi.GetStorefrontMerch(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.GetStorefrontMerch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorefrontMerch`: GetStorefrontMerchResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.GetStorefrontMerch`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorefrontMerchRequest struct via the builder pattern


### Return type

**GetStorefrontMerchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetStorefrontProfile

> GetStorefrontProfileResponse GetStorefrontProfile(ctx).Execute()

Your shop profile (the Crossly Shop page): banner, logo, tagline, bio, accent, announcement, shelves, blocks and All items options.

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
    resp, r, err := apiClient.StorefrontApi.GetStorefrontProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.GetStorefrontProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorefrontProfile`: GetStorefrontProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.GetStorefrontProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorefrontProfileRequest struct via the builder pattern


### Return type

**GetStorefrontProfileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetStorefrontShipping

> GetStorefrontShippingResponse GetStorefrontShipping(ctx).Execute()

Your Crossly shipping policy (domestic and international services, which countries you ship to) and your default ship-from address.

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
    resp, r, err := apiClient.StorefrontApi.GetStorefrontShipping(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.GetStorefrontShipping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorefrontShipping`: GetStorefrontShippingResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.GetStorefrontShipping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorefrontShippingRequest struct via the builder pattern


### Return type

**GetStorefrontShippingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetStorefrontSite

> StorefrontResponse3 GetStorefrontSite(ctx).Execute()

Your store website: the draft design, what is live, when each changed, whether the site is switched on, and its custom domain.

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
    resp, r, err := apiClient.StorefrontApi.GetStorefrontSite(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.GetStorefrontSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorefrontSite`: StorefrontResponse3
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.GetStorefrontSite`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorefrontSiteRequest struct via the builder pattern


### Return type

**StorefrontResponse3**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetStorefrontTradeSite

> StorefrontResponse3 GetStorefrontTradeSite(ctx).Execute()

Your store's trade (wholesale) site at /site/:slug/trade: the draft, what is live, whether it is on and its gating (public | prices_private | approved_only).

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
    resp, r, err := apiClient.StorefrontApi.GetStorefrontTradeSite(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.GetStorefrontTradeSite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorefrontTradeSite`: StorefrontResponse3
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.GetStorefrontTradeSite`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorefrontTradeSiteRequest struct via the builder pattern


### Return type

**StorefrontResponse3**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListStorefrontHighlights

> ListStorefrontHighlightsResponse ListStorefrontHighlights(ctx).Execute()

Your store’s story highlights in display order, each with its saved stories (media, overlays, product tie-in, when posted). Empty highlights included.

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
    resp, r, err := apiClient.StorefrontApi.ListStorefrontHighlights(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.ListStorefrontHighlights``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStorefrontHighlights`: ListStorefrontHighlightsResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.ListStorefrontHighlights`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStorefrontHighlightsRequest struct via the builder pattern


### Return type

**ListStorefrontHighlightsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListStorefrontSections

> StorefrontResponse4 ListStorefrontSections(ctx).Execute()

Your storefront sections in display order: kind, title, visible, config, and listing ids for listings shelves.

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
    resp, r, err := apiClient.StorefrontApi.ListStorefrontSections(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.ListStorefrontSections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStorefrontSections`: StorefrontResponse4
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.ListStorefrontSections`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStorefrontSectionsRequest struct via the builder pattern


### Return type

**StorefrontResponse4**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListStorefrontSiteVersions

> ListStorefrontSiteVersionsResponse ListStorefrontSiteVersions(ctx).Execute()

Saved versions of the website, newest first (publish, restore and template points; up to 30).

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
    resp, r, err := apiClient.StorefrontApi.ListStorefrontSiteVersions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.ListStorefrontSiteVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStorefrontSiteVersions`: ListStorefrontSiteVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.ListStorefrontSiteVersions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStorefrontSiteVersionsRequest struct via the builder pattern


### Return type

**ListStorefrontSiteVersionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListStorefrontStoryArchive

> ListStorefrontStoryArchiveResponse ListStorefrontStoryArchive(ctx).Execute()

Your stories, live and expired, newest first: what can be saved to a highlight. Paginated: pass before = meta.nextCursor of the previous page.

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
    resp, r, err := apiClient.StorefrontApi.ListStorefrontStoryArchive(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.ListStorefrontStoryArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStorefrontStoryArchive`: ListStorefrontStoryArchiveResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.ListStorefrontStoryArchive`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStorefrontStoryArchiveRequest struct via the builder pattern


### Return type

**ListStorefrontStoryArchiveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefront

> GetStorefrontResponse UpdateStorefront(ctx).Execute()

Update your storefront design. Omitted fields are unchanged, null clears. published=false keeps the design as a draft and buyers see the standard storefront.

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
    resp, r, err := apiClient.StorefrontApi.UpdateStorefront(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.UpdateStorefront``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefront`: GetStorefrontResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.UpdateStorefront`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontRequest struct via the builder pattern


### Return type

**GetStorefrontResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontHighlight

> StorefrontHighlightResponse UpdateStorefrontHighlight(ctx, id).Execute()

Rename a highlight or change its cover (coverUrl null = use the first story).

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
    resp, r, err := apiClient.StorefrontApi.UpdateStorefrontHighlight(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.UpdateStorefrontHighlight``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontHighlight`: StorefrontHighlightResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.UpdateStorefrontHighlight`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontHighlightRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontHighlightResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontHighlightItemOrder

> StorefrontHighlightResponse UpdateStorefrontHighlightItemOrder(ctx, id).Execute()

Set the order of a highlight’s saved stories: every item id, first to last.

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
    resp, r, err := apiClient.StorefrontApi.UpdateStorefrontHighlightItemOrder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.UpdateStorefrontHighlightItemOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontHighlightItemOrder`: StorefrontHighlightResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.UpdateStorefrontHighlightItemOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontHighlightItemOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontHighlightResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontHighlightOrder

> ListStorefrontHighlightsResponse UpdateStorefrontHighlightOrder(ctx).Execute()

Set the highlight order in one call: every highlight id, first to last.

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
    resp, r, err := apiClient.StorefrontApi.UpdateStorefrontHighlightOrder(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.UpdateStorefrontHighlightOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontHighlightOrder`: ListStorefrontHighlightsResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.UpdateStorefrontHighlightOrder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontHighlightOrderRequest struct via the builder pattern


### Return type

**ListStorefrontHighlightsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontProfile

> UpdateStorefrontProfileResponse UpdateStorefrontProfile(ctx).Execute()

Replace your shop profile (a complete ShopProfile). Saved straight away.

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
    resp, r, err := apiClient.StorefrontApi.UpdateStorefrontProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.UpdateStorefrontProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontProfile`: UpdateStorefrontProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.UpdateStorefrontProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontProfileRequest struct via the builder pattern


### Return type

**UpdateStorefrontProfileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontSection

> StorefrontSectionResponse UpdateStorefrontSection(ctx, id).Execute()

Rename, show/hide, refill (listings shelves) or reconfigure a section.

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
    resp, r, err := apiClient.StorefrontApi.UpdateStorefrontSection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.UpdateStorefrontSection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontSection`: StorefrontSectionResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.UpdateStorefrontSection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontSectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontSectionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontSectionReorder

> StorefrontResponse4 UpdateStorefrontSectionReorder(ctx).Execute()

Set the section order in one atomic call: every section id, top to bottom.

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
    resp, r, err := apiClient.StorefrontApi.UpdateStorefrontSectionReorder(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.UpdateStorefrontSectionReorder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontSectionReorder`: StorefrontResponse4
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.UpdateStorefrontSectionReorder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontSectionReorderRequest struct via the builder pattern


### Return type

**StorefrontResponse4**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontShipping

> GetStorefrontShippingResponse UpdateStorefrontShipping(ctx).Execute()

Set your shipping policy (international.enabled, shipsTo worldwide or ISO codes, excludes) and/or your default ship-from address (any country; 422 for a sanctioned origin).

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
    resp, r, err := apiClient.StorefrontApi.UpdateStorefrontShipping(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.UpdateStorefrontShipping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontShipping`: GetStorefrontShippingResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.UpdateStorefrontShipping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontShippingRequest struct via the builder pattern


### Return type

**GetStorefrontShippingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontSiteDraft

> StorefrontResponse2 UpdateStorefrontSiteDraft(ctx).Execute()

Save the website draft (a complete SiteDesign). Validated strictly (a bad field is a 400 naming its path) and sanitized; buyers see nothing until you publish.

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
    resp, r, err := apiClient.StorefrontApi.UpdateStorefrontSiteDraft(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.UpdateStorefrontSiteDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontSiteDraft`: StorefrontResponse2
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.UpdateStorefrontSiteDraft`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontSiteDraftRequest struct via the builder pattern


### Return type

**StorefrontResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontSiteEnabled

> StorefrontResponse3 UpdateStorefrontSiteEnabled(ctx).Execute()

Switch the website on or off. On requires a publish first; off makes /site/:slug 404 and the custom domain show the shop profile.

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
    resp, r, err := apiClient.StorefrontApi.UpdateStorefrontSiteEnabled(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.UpdateStorefrontSiteEnabled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontSiteEnabled`: StorefrontResponse3
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.UpdateStorefrontSiteEnabled`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontSiteEnabledRequest struct via the builder pattern


### Return type

**StorefrontResponse3**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontTradeSiteDraft

> StorefrontResponse2 UpdateStorefrontTradeSiteDraft(ctx).Execute()

Save the trade site draft (a complete SiteDesign; trade sections included). Validated strictly and sanitized.

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
    resp, r, err := apiClient.StorefrontApi.UpdateStorefrontTradeSiteDraft(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.UpdateStorefrontTradeSiteDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontTradeSiteDraft`: StorefrontResponse2
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.UpdateStorefrontTradeSiteDraft`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontTradeSiteDraftRequest struct via the builder pattern


### Return type

**StorefrontResponse2**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontTradeSiteEnabled

> StorefrontResponse3 UpdateStorefrontTradeSiteEnabled(ctx).Execute()

Switch the trade site on (needs a publish first) or off.

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
    resp, r, err := apiClient.StorefrontApi.UpdateStorefrontTradeSiteEnabled(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.UpdateStorefrontTradeSiteEnabled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontTradeSiteEnabled`: StorefrontResponse3
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.UpdateStorefrontTradeSiteEnabled`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontTradeSiteEnabledRequest struct via the builder pattern


### Return type

**StorefrontResponse3**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontTradeSiteGating

> StorefrontResponse3 UpdateStorefrontTradeSiteGating(ctx).Execute()

Who sees the trade site: public (everyone, with prices), prices_private (everyone, prices for approved buyers) or approved_only (approved buyers; others get the apply page).

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
    resp, r, err := apiClient.StorefrontApi.UpdateStorefrontTradeSiteGating(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontApi.UpdateStorefrontTradeSiteGating``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontTradeSiteGating`: StorefrontResponse3
    fmt.Fprintf(os.Stdout, "Response from `StorefrontApi.UpdateStorefrontTradeSiteGating`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontTradeSiteGatingRequest struct via the builder pattern


### Return type

**StorefrontResponse3**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

