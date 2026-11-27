# \AdsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdOffsiteCampaign**](AdsApi.md#CreateAdOffsiteCampaign) | **Post** /v1/ads/offsite/campaigns | Launch an offsite campaign for an item.
[**CreateAdOffsiteResume**](AdsApi.md#CreateAdOffsiteResume) | **Post** /v1/ads/offsite/resume | Clear an auto-pause and resume offsite spend.
[**CreateListingAdBulkAdd**](AdsApi.md#CreateListingAdBulkAdd) | **Post** /v1/listing-ads/bulk-add | Add listings to a campaign by master listing id: each listing&#39;s live row on &#x60;platform&#x60; joins campaignId (a Crossly campaign uuid, or the marketplace&#39;s campaign id). A listing not live there is its own failed result.
[**CreateListingAdCampaign**](AdsApi.md#CreateListingAdCampaign) | **Post** /v1/listing-ads/{platform}/campaigns | Create a campaign on the marketplace: program (promoted_listing_rate &#x3D; eBay General, a % of the sale; cpc_campaign &#x3D; eBay Priority, a daily budget), name, ratePct or dailyBudgetCents, startsAt/endsAt, and the program fields. Idempotency-Key supported.
[**CreateListingAdCampaignListing**](AdsApi.md#CreateListingAdCampaignListing) | **Post** /v1/listing-ads/{platform}/campaigns/{campaignId}/listings | Promote listings in a marketplace campaign. platformListingIds are Crossly platform-listing ids on that marketplace; ratePct sets the ad rate (eBay General). One result per listing.
[**CreateListingAdCampaignListingRemove**](AdsApi.md#CreateListingAdCampaignListingRemove) | **Post** /v1/listing-ads/{platform}/campaigns/{campaignId}/listings/remove | Stop promoting listings in a marketplace campaign. One result per listing.
[**CreateSponsoredAdBillingTopupLink**](AdsApi.md#CreateSponsoredAdBillingTopupLink) | **Post** /v1/sponsored-ads/billing/topup-link | A Stripe Checkout link that adds amountCents ($5–$10,000) to the ad balance when the seller pays it. Nothing is charged until they do. Use an Idempotency-Key to get the same link back on a retry.
[**CreateSponsoredAdBoostRemove**](AdsApi.md#CreateSponsoredAdBoostRemove) | **Post** /v1/sponsored-ads/boosts/remove | Stop boosting listings in bulk.
[**CreateSponsoredAdCampaign**](AdsApi.md#CreateSponsoredAdCampaign) | **Post** /v1/sponsored-ads/campaigns | Create a Sponsored Listings campaign: name, dailyBudgetCents (≥100), optional totalBudgetCents, defaultBidCents (5–10000; the max bid under bidStrategy \&quot;auto\&quot;), placements (search, browse, product_page, home_rail, storefront, worlds_surface), goal (sales | visits), optional targetCategories / targetRegions (ISO countries) / targetDevices (mobile, desktop, app), optional startsAt/endsAt, listings [{platformListingId, bidCents?, boostBps? (100–2000, pay-on-sale Boost; null clears)}], keywords [{keyword, matchType broad|phrase|exact, negative?, bidCents?}], worldTargets [{kind showroom|convention, id, bidCents?, status?}] (your showroom or a convention you host, promoted in the Crossly World directory; at least one listing or world target), status active|draft. A catalog boost: targetType \&quot;wholesale_catalog\&quot; + targetCatalogId (your catalog), no listings, keywords or world targets; it ranks Sponsored in the wholesale directory. bidStrategy defaults to your account default. Clicks are charged from your ad funds (ad balance and/or CBX, ad credit first); with none it waits as out_of_budget.
[**CreateSponsoredAdCampaignDuplicate**](AdsApi.md#CreateSponsoredAdCampaignDuplicate) | **Post** /v1/sponsored-ads/campaigns/{id}/duplicate | Copy a campaign (settings, listings, keywords) as a new draft.
[**CreateSponsoredAdCampaignStatus**](AdsApi.md#CreateSponsoredAdCampaignStatus) | **Post** /v1/sponsored-ads/campaigns/{id}/status | Set a campaign active (launch/resume), paused, or ended (final). Resuming with no ad funds to pay for clicks leaves it out_of_budget until you top up.
[**GetAdOffsite**](AdsApi.md#GetAdOffsite) | **Get** /v1/ads/offsite | Your offsite-ads opt-in and its terms.
[**GetAdOffsiteEligibility**](AdsApi.md#GetAdOffsiteEligibility) | **Get** /v1/ads/offsite/eligibility | Whether an item can run offsite, and why not.
[**GetAdOffsiteReport**](AdsApi.md#GetAdOffsiteReport) | **Get** /v1/ads/offsite/report | What your offsite budget bought — including the misses.
[**GetListingAdListing**](AdsApi.md#GetListingAdListing) | **Get** /v1/listing-ads/listings/{listingId} | Every promotion of one listing: its Crossly campaigns and Boost, and each marketplace ad with status, campaign, rate and last sync.
[**GetSponsoredAdBilling**](AdsApi.md#GetSponsoredAdBilling) | **Get** /v1/sponsored-ads/billing | What pays for your ads: headroom across the billing rails (ad balance and/or CBX), unbilled click spend, settlement charges, daily spend, Boost fees and CBX ad-credit grants; funds is the full funds summary (same as /v1/sponsored-ads/funds). refresh&#x3D;true re-reads the balance.
[**GetSponsoredAdBillingStatement**](AdsApi.md#GetSponsoredAdBillingStatement) | **Get** /v1/sponsored-ads/billing/statements/{month} | One month&#39;s ads statement: opening and closing ad balance, top-ups, click spend by campaign, how clicks were paid, credits, refunds and Boost fees.
[**GetSponsoredAdBillingStatementPdf**](AdsApi.md#GetSponsoredAdBillingStatementPdf) | **Get** /v1/sponsored-ads/billing/statements/{month}/pdf | One month&#39;s ads statement as a PDF.
[**GetSponsoredAdBillingTopupReceipt**](AdsApi.md#GetSponsoredAdBillingTopupReceipt) | **Get** /v1/sponsored-ads/billing/topups/{id}/receipt | One ad balance top-up&#39;s receipt as a PDF (each top-up in the billing history carries its id and receipt number).
[**GetSponsoredAdCampaign**](AdsApi.md#GetSponsoredAdCampaign) | **Get** /v1/sponsored-ads/campaigns/{id} | One campaign: settings, its listings (bid, suggested bid, live, per-listing stats) and keywords. A catalog boost (targetType wholesale_catalog) has &#x60;catalogs&#x60;: per boosted catalog impressions, clicks, spend, attributed trade orders and revenue, CTR and ROAS.
[**GetSponsoredAdFund**](AdsApi.md#GetSponsoredAdFund) | **Get** /v1/sponsored-ads/funds | What can pay for clicks now: the billing rails in effect (ad balance and/or CBX), the ad balance and ad credit, CBX headroom, unbilled clicks, available cents, auto top-up, the saved payment method and the refundable balance. refresh&#x3D;true re-reads CBX.
[**GetSponsoredAdReport**](AdsApi.md#GetSponsoredAdReport) | **Get** /v1/sponsored-ads/report | Sponsored report for a range: totals (spend, impressions, clicks, CTR, average CPC, attributed orders and revenue, ROAS), a zero-filled daily series, per-campaign rows (targetType listings | wholesale_catalog) and &#x60;catalogs&#x60;, one row per boosted catalog. Listing orders are Crossly purchases of the advertised listing within 7 days of a click by the same buyer; catalog orders are trade orders with that seller within 7 days of a catalog click by the same buyer. Catalog impressions count once per viewer per day.
[**GetSponsoredAdSetting**](AdsApi.md#GetSponsoredAdSetting) | **Get** /v1/sponsored-ads/settings | Ads quick settings: adsPaused, accountDailyCapCents and spentTodayCents, defaultBidStrategy, autoTopup.
[**ListListingAdCampaigns**](AdsApi.md#ListListingAdCampaigns) | **Get** /v1/listing-ads/{platform}/campaigns | Your ad campaigns on one marketplace (cached up to 15 minutes; refresh&#x3D;true asks the marketplace now).
[**ListListingAdPlatforms**](AdsApi.md#ListListingAdPlatforms) | **Get** /v1/listing-ads/platforms | The ads registry: each marketplace that sells promotion, its programs (kind, scope, campaign and listing fields), which verbs are wired / need a capture / are unsupported, and whether you are connected there.
[**ListSponsoredAdBillingHistory**](AdsApi.md#ListSponsoredAdBillingHistory) | **Get** /v1/sponsored-ads/billing/history | Billing history newest first: top-ups (status, receipt number), click charges (how each was paid: ad balance, credit, CBX, written off), credits, refunds and Boost fees. kind filters; before (ISO time) pages; limit ≤ 100.
[**ListSponsoredAdBillingStatements**](AdsApi.md#ListSponsoredAdBillingStatements) | **Get** /v1/sponsored-ads/billing/statements | Months (YYYY-MM) with any ad billing activity, newest first.
[**ListSponsoredAdBoosts**](AdsApi.md#ListSponsoredAdBoosts) | **Get** /v1/sponsored-ads/boosts | Boosted (pay-on-sale) listings: the rate, sales and fees taken so far. A boost lifts organic ranking by up to 10% and costs its rate only when the listing sells on Crossly.
[**ListSponsoredAdCampaigns**](AdsApi.md#ListSponsoredAdCampaigns) | **Get** /v1/sponsored-ads/campaigns | Your Sponsored Listings campaigns with status, pacing (spent today, daily cap reached), budgets, bids, placements and stats for the range (from/to, YYYY-MM-DD; default last 30 days).
[**ListSponsoredAdListings**](AdsApi.md#ListSponsoredAdListings) | **Get** /v1/sponsored-ads/listings | Per-listing ad performance across campaigns for a range, plus suggestions: live listings with views but few sales that are not sponsored yet, with a suggested bid.
[**UpdateAdOffsite**](AdsApi.md#UpdateAdOffsite) | **Put** /v1/ads/offsite | Turn offsite ads on or off. Yours alone to set.
[**UpdateListingAdCampaign**](AdsApi.md#UpdateListingAdCampaign) | **Patch** /v1/listing-ads/{platform}/campaigns/{campaignId} | Change a marketplace campaign: name, dailyBudgetCents, endsAt (null clears), status paused | running | ended.
[**UpdateListingAdListingRate**](AdsApi.md#UpdateListingAdListingRate) | **Put** /v1/listing-ads/{platform}/listings/{platformListingId}/rate | Change one promoted listing&#39;s ad rate (percent of the sale) in the campaign it is in.
[**UpdateSponsoredAdBillingAutoTopup**](AdsApi.md#UpdateSponsoredAdBillingAutoTopup) | **Put** /v1/sponsored-ads/billing/auto-topup | Auto top-up: enabled, thresholdCents (top up when the ad balance drops below it), amountCents (how much). Needs a saved payment method (the first top-up saves it).
[**UpdateSponsoredAdBoost**](AdsApi.md#UpdateSponsoredAdBoost) | **Put** /v1/sponsored-ads/boosts | Boost listings in bulk (platformListingIds, upliftBps 100–2000 &#x3D; 1–20%), or change their rate. Your own live Crossly listings only.
[**UpdateSponsoredAdCampaign**](AdsApi.md#UpdateSponsoredAdCampaign) | **Patch** /v1/sponsored-ads/campaigns/{id} | Edit a campaign. Any of the create fields; listings and keywords, when given, replace the current set. Raising the total budget resumes a campaign stopped by it.
[**UpdateSponsoredAdSetting**](AdsApi.md#UpdateSponsoredAdSetting) | **Patch** /v1/sponsored-ads/settings | Change ads quick settings: adsPaused (pause / resume every ad), accountDailyCapCents (null &#x3D; no cap), defaultBidStrategy (auto | manual) for new campaigns, autoTopup {enabled, thresholdCents?, amountCents?}, railPreference (cash_first | cbx_first, both rails live only).



## CreateAdOffsiteCampaign

> CreateAdOffsiteCampaignResponse CreateAdOffsiteCampaign(ctx).Execute()

Launch an offsite campaign for an item.



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
    resp, r, err := apiClient.AdsApi.CreateAdOffsiteCampaign(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.CreateAdOffsiteCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAdOffsiteCampaign`: CreateAdOffsiteCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.CreateAdOffsiteCampaign`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdOffsiteCampaignRequest struct via the builder pattern


### Return type

**CreateAdOffsiteCampaignResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateAdOffsiteResume

> CreateAdOffsiteResumeResponse CreateAdOffsiteResume(ctx).Execute()

Clear an auto-pause and resume offsite spend.



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
    resp, r, err := apiClient.AdsApi.CreateAdOffsiteResume(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.CreateAdOffsiteResume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAdOffsiteResume`: CreateAdOffsiteResumeResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.CreateAdOffsiteResume`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdOffsiteResumeRequest struct via the builder pattern


### Return type

**CreateAdOffsiteResumeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingAdBulkAdd

> CreateListingAdBulkAddResponse CreateListingAdBulkAdd(ctx).Execute()

Add listings to a campaign by master listing id: each listing's live row on `platform` joins campaignId (a Crossly campaign uuid, or the marketplace's campaign id). A listing not live there is its own failed result.

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
    resp, r, err := apiClient.AdsApi.CreateListingAdBulkAdd(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.CreateListingAdBulkAdd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingAdBulkAdd`: CreateListingAdBulkAddResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.CreateListingAdBulkAdd`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingAdBulkAddRequest struct via the builder pattern


### Return type

**CreateListingAdBulkAddResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingAdCampaign

> CreateListingAdCampaignResponse CreateListingAdCampaign(ctx, platform).Execute()

Create a campaign on the marketplace: program (promoted_listing_rate = eBay General, a % of the sale; cpc_campaign = eBay Priority, a daily budget), name, ratePct or dailyBudgetCents, startsAt/endsAt, and the program fields. Idempotency-Key supported.

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
    platform := "platform_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsApi.CreateListingAdCampaign(context.Background(), platform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.CreateListingAdCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingAdCampaign`: CreateListingAdCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.CreateListingAdCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingAdCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateListingAdCampaignResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingAdCampaignListing

> CreateListingAdCampaignListingResponse CreateListingAdCampaignListing(ctx, platform, campaignId).Execute()

Promote listings in a marketplace campaign. platformListingIds are Crossly platform-listing ids on that marketplace; ratePct sets the ad rate (eBay General). One result per listing.

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
    platform := "platform_example" // string | 
    campaignId := "campaignId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsApi.CreateListingAdCampaignListing(context.Background(), platform, campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.CreateListingAdCampaignListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingAdCampaignListing`: CreateListingAdCampaignListingResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.CreateListingAdCampaignListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingAdCampaignListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateListingAdCampaignListingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingAdCampaignListingRemove

> CreateListingAdCampaignListingResponse CreateListingAdCampaignListingRemove(ctx, platform, campaignId).Execute()

Stop promoting listings in a marketplace campaign. One result per listing.

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
    platform := "platform_example" // string | 
    campaignId := "campaignId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsApi.CreateListingAdCampaignListingRemove(context.Background(), platform, campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.CreateListingAdCampaignListingRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingAdCampaignListingRemove`: CreateListingAdCampaignListingResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.CreateListingAdCampaignListingRemove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingAdCampaignListingRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateListingAdCampaignListingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateSponsoredAdBillingTopupLink

> CreateSponsoredAdBillingTopupLinkResponse CreateSponsoredAdBillingTopupLink(ctx).Execute()

A Stripe Checkout link that adds amountCents ($5–$10,000) to the ad balance when the seller pays it. Nothing is charged until they do. Use an Idempotency-Key to get the same link back on a retry.

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
    resp, r, err := apiClient.AdsApi.CreateSponsoredAdBillingTopupLink(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.CreateSponsoredAdBillingTopupLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSponsoredAdBillingTopupLink`: CreateSponsoredAdBillingTopupLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.CreateSponsoredAdBillingTopupLink`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsoredAdBillingTopupLinkRequest struct via the builder pattern


### Return type

**CreateSponsoredAdBillingTopupLinkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateSponsoredAdBoostRemove

> CreateSponsoredAdBoostRemoveResponse CreateSponsoredAdBoostRemove(ctx).Execute()

Stop boosting listings in bulk.

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
    resp, r, err := apiClient.AdsApi.CreateSponsoredAdBoostRemove(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.CreateSponsoredAdBoostRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSponsoredAdBoostRemove`: CreateSponsoredAdBoostRemoveResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.CreateSponsoredAdBoostRemove`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsoredAdBoostRemoveRequest struct via the builder pattern


### Return type

**CreateSponsoredAdBoostRemoveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateSponsoredAdCampaign

> SponsoredAdCampaignResponse CreateSponsoredAdCampaign(ctx).Execute()

Create a Sponsored Listings campaign: name, dailyBudgetCents (≥100), optional totalBudgetCents, defaultBidCents (5–10000; the max bid under bidStrategy \"auto\"), placements (search, browse, product_page, home_rail, storefront, worlds_surface), goal (sales | visits), optional targetCategories / targetRegions (ISO countries) / targetDevices (mobile, desktop, app), optional startsAt/endsAt, listings [{platformListingId, bidCents?, boostBps? (100–2000, pay-on-sale Boost; null clears)}], keywords [{keyword, matchType broad|phrase|exact, negative?, bidCents?}], worldTargets [{kind showroom|convention, id, bidCents?, status?}] (your showroom or a convention you host, promoted in the Crossly World directory; at least one listing or world target), status active|draft. A catalog boost: targetType \"wholesale_catalog\" + targetCatalogId (your catalog), no listings, keywords or world targets; it ranks Sponsored in the wholesale directory. bidStrategy defaults to your account default. Clicks are charged from your ad funds (ad balance and/or CBX, ad credit first); with none it waits as out_of_budget.

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
    resp, r, err := apiClient.AdsApi.CreateSponsoredAdCampaign(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.CreateSponsoredAdCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSponsoredAdCampaign`: SponsoredAdCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.CreateSponsoredAdCampaign`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsoredAdCampaignRequest struct via the builder pattern


### Return type

**SponsoredAdCampaignResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateSponsoredAdCampaignDuplicate

> SponsoredAdCampaignResponse CreateSponsoredAdCampaignDuplicate(ctx, id).Execute()

Copy a campaign (settings, listings, keywords) as a new draft.

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
    resp, r, err := apiClient.AdsApi.CreateSponsoredAdCampaignDuplicate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.CreateSponsoredAdCampaignDuplicate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSponsoredAdCampaignDuplicate`: SponsoredAdCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.CreateSponsoredAdCampaignDuplicate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsoredAdCampaignDuplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**SponsoredAdCampaignResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateSponsoredAdCampaignStatus

> SponsoredAdCampaignResponse CreateSponsoredAdCampaignStatus(ctx, id).Execute()

Set a campaign active (launch/resume), paused, or ended (final). Resuming with no ad funds to pay for clicks leaves it out_of_budget until you top up.

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
    resp, r, err := apiClient.AdsApi.CreateSponsoredAdCampaignStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.CreateSponsoredAdCampaignStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSponsoredAdCampaignStatus`: SponsoredAdCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.CreateSponsoredAdCampaignStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSponsoredAdCampaignStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**SponsoredAdCampaignResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAdOffsite

> GetAdOffsiteResponse GetAdOffsite(ctx).Execute()

Your offsite-ads opt-in and its terms.



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
    resp, r, err := apiClient.AdsApi.GetAdOffsite(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetAdOffsite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdOffsite`: GetAdOffsiteResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetAdOffsite`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdOffsiteRequest struct via the builder pattern


### Return type

**GetAdOffsiteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAdOffsiteEligibility

> GetAdOffsiteEligibilityResponse GetAdOffsiteEligibility(ctx).Execute()

Whether an item can run offsite, and why not.



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
    resp, r, err := apiClient.AdsApi.GetAdOffsiteEligibility(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetAdOffsiteEligibility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdOffsiteEligibility`: GetAdOffsiteEligibilityResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetAdOffsiteEligibility`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdOffsiteEligibilityRequest struct via the builder pattern


### Return type

**GetAdOffsiteEligibilityResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAdOffsiteReport

> GetAdOffsiteReportResponse GetAdOffsiteReport(ctx).Execute()

What your offsite budget bought — including the misses.



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
    resp, r, err := apiClient.AdsApi.GetAdOffsiteReport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetAdOffsiteReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdOffsiteReport`: GetAdOffsiteReportResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetAdOffsiteReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdOffsiteReportRequest struct via the builder pattern


### Return type

**GetAdOffsiteReportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetListingAdListing

> GetListingAdListingResponse GetListingAdListing(ctx, listingId).Execute()

Every promotion of one listing: its Crossly campaigns and Boost, and each marketplace ad with status, campaign, rate and last sync.

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
    listingId := "listingId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsApi.GetListingAdListing(context.Background(), listingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetListingAdListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListingAdListing`: GetListingAdListingResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetListingAdListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingAdListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetListingAdListingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetSponsoredAdBilling

> GetSponsoredAdBillingResponse GetSponsoredAdBilling(ctx).Execute()

What pays for your ads: headroom across the billing rails (ad balance and/or CBX), unbilled click spend, settlement charges, daily spend, Boost fees and CBX ad-credit grants; funds is the full funds summary (same as /v1/sponsored-ads/funds). refresh=true re-reads the balance.

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
    resp, r, err := apiClient.AdsApi.GetSponsoredAdBilling(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetSponsoredAdBilling``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSponsoredAdBilling`: GetSponsoredAdBillingResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetSponsoredAdBilling`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsoredAdBillingRequest struct via the builder pattern


### Return type

**GetSponsoredAdBillingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetSponsoredAdBillingStatement

> GetSponsoredAdBillingStatementResponse GetSponsoredAdBillingStatement(ctx, month).Execute()

One month's ads statement: opening and closing ad balance, top-ups, click spend by campaign, how clicks were paid, credits, refunds and Boost fees.

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
    month := "month_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsApi.GetSponsoredAdBillingStatement(context.Background(), month).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetSponsoredAdBillingStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSponsoredAdBillingStatement`: GetSponsoredAdBillingStatementResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetSponsoredAdBillingStatement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**month** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsoredAdBillingStatementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetSponsoredAdBillingStatementResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetSponsoredAdBillingStatementPdf

> *os.File GetSponsoredAdBillingStatementPdf(ctx, month).Execute()

One month's ads statement as a PDF.

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
    month := "month_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsApi.GetSponsoredAdBillingStatementPdf(context.Background(), month).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetSponsoredAdBillingStatementPdf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSponsoredAdBillingStatementPdf`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetSponsoredAdBillingStatementPdf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**month** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsoredAdBillingStatementPdfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

***os.File**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetSponsoredAdBillingTopupReceipt

> *os.File GetSponsoredAdBillingTopupReceipt(ctx, id).Execute()

One ad balance top-up's receipt as a PDF (each top-up in the billing history carries its id and receipt number).

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
    resp, r, err := apiClient.AdsApi.GetSponsoredAdBillingTopupReceipt(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetSponsoredAdBillingTopupReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSponsoredAdBillingTopupReceipt`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetSponsoredAdBillingTopupReceipt`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsoredAdBillingTopupReceiptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

***os.File**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetSponsoredAdCampaign

> SponsoredAdCampaignResponse GetSponsoredAdCampaign(ctx, id).Execute()

One campaign: settings, its listings (bid, suggested bid, live, per-listing stats) and keywords. A catalog boost (targetType wholesale_catalog) has `catalogs`: per boosted catalog impressions, clicks, spend, attributed trade orders and revenue, CTR and ROAS.

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
    resp, r, err := apiClient.AdsApi.GetSponsoredAdCampaign(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetSponsoredAdCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSponsoredAdCampaign`: SponsoredAdCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetSponsoredAdCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsoredAdCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**SponsoredAdCampaignResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetSponsoredAdFund

> GetSponsoredAdFundResponse GetSponsoredAdFund(ctx).Execute()

What can pay for clicks now: the billing rails in effect (ad balance and/or CBX), the ad balance and ad credit, CBX headroom, unbilled clicks, available cents, auto top-up, the saved payment method and the refundable balance. refresh=true re-reads CBX.

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
    resp, r, err := apiClient.AdsApi.GetSponsoredAdFund(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetSponsoredAdFund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSponsoredAdFund`: GetSponsoredAdFundResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetSponsoredAdFund`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsoredAdFundRequest struct via the builder pattern


### Return type

**GetSponsoredAdFundResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetSponsoredAdReport

> GetSponsoredAdReportResponse GetSponsoredAdReport(ctx).Execute()

Sponsored report for a range: totals (spend, impressions, clicks, CTR, average CPC, attributed orders and revenue, ROAS), a zero-filled daily series, per-campaign rows (targetType listings | wholesale_catalog) and `catalogs`, one row per boosted catalog. Listing orders are Crossly purchases of the advertised listing within 7 days of a click by the same buyer; catalog orders are trade orders with that seller within 7 days of a catalog click by the same buyer. Catalog impressions count once per viewer per day.

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
    resp, r, err := apiClient.AdsApi.GetSponsoredAdReport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetSponsoredAdReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSponsoredAdReport`: GetSponsoredAdReportResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetSponsoredAdReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsoredAdReportRequest struct via the builder pattern


### Return type

**GetSponsoredAdReportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetSponsoredAdSetting

> SponsoredAdResponse GetSponsoredAdSetting(ctx).Execute()

Ads quick settings: adsPaused, accountDailyCapCents and spentTodayCents, defaultBidStrategy, autoTopup.

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
    resp, r, err := apiClient.AdsApi.GetSponsoredAdSetting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetSponsoredAdSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSponsoredAdSetting`: SponsoredAdResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetSponsoredAdSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSponsoredAdSettingRequest struct via the builder pattern


### Return type

**SponsoredAdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListListingAdCampaigns

> ListListingAdCampaignsResponse ListListingAdCampaigns(ctx, platform).Execute()

Your ad campaigns on one marketplace (cached up to 15 minutes; refresh=true asks the marketplace now).

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
    platform := "platform_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsApi.ListListingAdCampaigns(context.Background(), platform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.ListListingAdCampaigns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListListingAdCampaigns`: ListListingAdCampaignsResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.ListListingAdCampaigns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListListingAdCampaignsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListListingAdCampaignsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListListingAdPlatforms

> ListListingAdPlatformsResponse ListListingAdPlatforms(ctx).Execute()

The ads registry: each marketplace that sells promotion, its programs (kind, scope, campaign and listing fields), which verbs are wired / need a capture / are unsupported, and whether you are connected there.

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
    resp, r, err := apiClient.AdsApi.ListListingAdPlatforms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.ListListingAdPlatforms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListListingAdPlatforms`: ListListingAdPlatformsResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.ListListingAdPlatforms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListListingAdPlatformsRequest struct via the builder pattern


### Return type

**ListListingAdPlatformsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListSponsoredAdBillingHistory

> ListSponsoredAdBillingHistoryResponse ListSponsoredAdBillingHistory(ctx).Execute()

Billing history newest first: top-ups (status, receipt number), click charges (how each was paid: ad balance, credit, CBX, written off), credits, refunds and Boost fees. kind filters; before (ISO time) pages; limit ≤ 100.

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
    resp, r, err := apiClient.AdsApi.ListSponsoredAdBillingHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.ListSponsoredAdBillingHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSponsoredAdBillingHistory`: ListSponsoredAdBillingHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.ListSponsoredAdBillingHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSponsoredAdBillingHistoryRequest struct via the builder pattern


### Return type

**ListSponsoredAdBillingHistoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListSponsoredAdBillingStatements

> ListSponsoredAdBillingStatementsResponse ListSponsoredAdBillingStatements(ctx).Execute()

Months (YYYY-MM) with any ad billing activity, newest first.

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
    resp, r, err := apiClient.AdsApi.ListSponsoredAdBillingStatements(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.ListSponsoredAdBillingStatements``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSponsoredAdBillingStatements`: ListSponsoredAdBillingStatementsResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.ListSponsoredAdBillingStatements`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSponsoredAdBillingStatementsRequest struct via the builder pattern


### Return type

**ListSponsoredAdBillingStatementsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListSponsoredAdBoosts

> ListSponsoredAdBoostsResponse ListSponsoredAdBoosts(ctx).Execute()

Boosted (pay-on-sale) listings: the rate, sales and fees taken so far. A boost lifts organic ranking by up to 10% and costs its rate only when the listing sells on Crossly.

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
    resp, r, err := apiClient.AdsApi.ListSponsoredAdBoosts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.ListSponsoredAdBoosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSponsoredAdBoosts`: ListSponsoredAdBoostsResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.ListSponsoredAdBoosts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSponsoredAdBoostsRequest struct via the builder pattern


### Return type

**ListSponsoredAdBoostsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListSponsoredAdCampaigns

> ListSponsoredAdCampaignsResponse ListSponsoredAdCampaigns(ctx).Execute()

Your Sponsored Listings campaigns with status, pacing (spent today, daily cap reached), budgets, bids, placements and stats for the range (from/to, YYYY-MM-DD; default last 30 days).

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
    resp, r, err := apiClient.AdsApi.ListSponsoredAdCampaigns(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.ListSponsoredAdCampaigns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSponsoredAdCampaigns`: ListSponsoredAdCampaignsResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.ListSponsoredAdCampaigns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSponsoredAdCampaignsRequest struct via the builder pattern


### Return type

**ListSponsoredAdCampaignsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListSponsoredAdListings

> ListSponsoredAdListingsResponse ListSponsoredAdListings(ctx).Execute()

Per-listing ad performance across campaigns for a range, plus suggestions: live listings with views but few sales that are not sponsored yet, with a suggested bid.

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
    resp, r, err := apiClient.AdsApi.ListSponsoredAdListings(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.ListSponsoredAdListings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSponsoredAdListings`: ListSponsoredAdListingsResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.ListSponsoredAdListings`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSponsoredAdListingsRequest struct via the builder pattern


### Return type

**ListSponsoredAdListingsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateAdOffsite

> UpdateAdOffsiteResponse UpdateAdOffsite(ctx).Execute()

Turn offsite ads on or off. Yours alone to set.



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
    resp, r, err := apiClient.AdsApi.UpdateAdOffsite(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.UpdateAdOffsite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAdOffsite`: UpdateAdOffsiteResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.UpdateAdOffsite`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdOffsiteRequest struct via the builder pattern


### Return type

**UpdateAdOffsiteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateListingAdCampaign

> CreateListingAdCampaignResponse UpdateListingAdCampaign(ctx, platform, campaignId).Execute()

Change a marketplace campaign: name, dailyBudgetCents, endsAt (null clears), status paused | running | ended.

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
    platform := "platform_example" // string | 
    campaignId := "campaignId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsApi.UpdateListingAdCampaign(context.Background(), platform, campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.UpdateListingAdCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateListingAdCampaign`: CreateListingAdCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.UpdateListingAdCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListingAdCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateListingAdCampaignResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateListingAdListingRate

> UpdateListingAdListingRateResponse UpdateListingAdListingRate(ctx, platform, platformListingId).Execute()

Change one promoted listing's ad rate (percent of the sale) in the campaign it is in.

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
    platform := "platform_example" // string | 
    platformListingId := "platformListingId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdsApi.UpdateListingAdListingRate(context.Background(), platform, platformListingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.UpdateListingAdListingRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateListingAdListingRate`: UpdateListingAdListingRateResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.UpdateListingAdListingRate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 
**platformListingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListingAdListingRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**UpdateListingAdListingRateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateSponsoredAdBillingAutoTopup

> SponsoredAdResponse UpdateSponsoredAdBillingAutoTopup(ctx).Execute()

Auto top-up: enabled, thresholdCents (top up when the ad balance drops below it), amountCents (how much). Needs a saved payment method (the first top-up saves it).

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
    resp, r, err := apiClient.AdsApi.UpdateSponsoredAdBillingAutoTopup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.UpdateSponsoredAdBillingAutoTopup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSponsoredAdBillingAutoTopup`: SponsoredAdResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.UpdateSponsoredAdBillingAutoTopup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSponsoredAdBillingAutoTopupRequest struct via the builder pattern


### Return type

**SponsoredAdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateSponsoredAdBoost

> UpdateSponsoredAdBoostResponse UpdateSponsoredAdBoost(ctx).Execute()

Boost listings in bulk (platformListingIds, upliftBps 100–2000 = 1–20%), or change their rate. Your own live Crossly listings only.

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
    resp, r, err := apiClient.AdsApi.UpdateSponsoredAdBoost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.UpdateSponsoredAdBoost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSponsoredAdBoost`: UpdateSponsoredAdBoostResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.UpdateSponsoredAdBoost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSponsoredAdBoostRequest struct via the builder pattern


### Return type

**UpdateSponsoredAdBoostResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateSponsoredAdCampaign

> SponsoredAdCampaignResponse UpdateSponsoredAdCampaign(ctx, id).Execute()

Edit a campaign. Any of the create fields; listings and keywords, when given, replace the current set. Raising the total budget resumes a campaign stopped by it.

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
    resp, r, err := apiClient.AdsApi.UpdateSponsoredAdCampaign(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.UpdateSponsoredAdCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSponsoredAdCampaign`: SponsoredAdCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.UpdateSponsoredAdCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSponsoredAdCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**SponsoredAdCampaignResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateSponsoredAdSetting

> SponsoredAdResponse UpdateSponsoredAdSetting(ctx).Execute()

Change ads quick settings: adsPaused (pause / resume every ad), accountDailyCapCents (null = no cap), defaultBidStrategy (auto | manual) for new campaigns, autoTopup {enabled, thresholdCents?, amountCents?}, railPreference (cash_first | cbx_first, both rails live only).

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
    resp, r, err := apiClient.AdsApi.UpdateSponsoredAdSetting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.UpdateSponsoredAdSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSponsoredAdSetting`: SponsoredAdResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.UpdateSponsoredAdSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSponsoredAdSettingRequest struct via the builder pattern


### Return type

**SponsoredAdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

