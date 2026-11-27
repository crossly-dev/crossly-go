# \ListingsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAiDetectionDispute**](ListingsApi.md#CreateAiDetectionDispute) | **Post** /v1/ai-detections/{id}/dispute | Dispute an AI-image detection on your photo (sends it to Crossly review).
[**CreateEvent**](ListingsApi.md#CreateEvent) | **Post** /v1/events | Create an event. A draft unless publish:true. Linked listings, the live show, the drop and the discount code must all be yours.
[**CreateEventPublish**](ListingsApi.md#CreateEventPublish) | **Post** /v1/events/{id}/publish | Publish an event: its page goes live, and with notifyFollowers your followers are told once.
[**CreateListing**](ListingsApi.md#CreateListing) | **Post** /v1/listings | Create a listing and fan out crosspost jobs across platforms.
[**CreateListingBulk**](ListingsApi.md#CreateListingBulk) | **Post** /v1/listings/bulk | Create up to 50 listings in one call, with a result per item.
[**CreateListingBulkCheckStatus**](ListingsApi.md#CreateListingBulkCheckStatus) | **Post** /v1/listings/bulk-check-status | Check listing status on platforms
[**CreateListingBulkCrosspost**](ListingsApi.md#CreateListingBulkCrosspost) | **Post** /v1/listings/bulk-crosspost | Bulk crosspost (no delist phase)
[**CreateListingBulkDelete**](ListingsApi.md#CreateListingBulkDelete) | **Post** /v1/listings/bulk-delete | Bulk archive + delist
[**CreateListingBulkDelist**](ListingsApi.md#CreateListingBulkDelist) | **Post** /v1/listings/bulk-delist | Bulk delist from platforms
[**CreateListingBulkDelistPreview**](ListingsApi.md#CreateListingBulkDelistPreview) | **Post** /v1/listings/bulk-delist-preview | Preview which marketplaces a delist would touch
[**CreateListingBulkHardDelete**](ListingsApi.md#CreateListingBulkHardDelete) | **Post** /v1/listings/bulk-hard-delete | Permanently delete listings (live ones need allowLive)
[**CreateListingBulkRelist**](ListingsApi.md#CreateListingBulkRelist) | **Post** /v1/listings/bulk-relist | Bulk relist across platforms
[**CreateListingBulkUpdate**](ListingsApi.md#CreateListingBulkUpdate) | **Post** /v1/listings/bulk-update | Bulk update listing fields
[**CreateListingById**](ListingsApi.md#CreateListingById) | **Post** /v1/listings/by-ids | Fetch hydrated listings by ID
[**CreateListingCheckDuplicate**](ListingsApi.md#CreateListingCheckDuplicate) | **Post** /v1/listings/check-duplicates | Check whether the seller already owns something matching this title/photo, and what to do about it.
[**CreateListingCombine**](ListingsApi.md#CreateListingCombine) | **Post** /v1/listings/combine | Combine duplicate listings into one: sums their stock, delists and archives the rest.
[**CreateListingDiscrepancyResolve**](ListingsApi.md#CreateListingDiscrepancyResolve) | **Post** /v1/listings/{id}/discrepancies/{discrepancyId}/resolve | Resolve a detected marketplace-drift discrepancy: accept the platform value, push ours back, relist to apply it, or dismiss.
[**CreateListingImportByUrl**](ListingsApi.md#CreateListingImportByUrl) | **Post** /v1/listings/{id}/import-by-url | Attach a real platform listing to this listing by pasting its live URL.
[**CreateListingMagicFill**](ListingsApi.md#CreateListingMagicFill) | **Post** /v1/listings/{id}/magic-fill | Auto-fill empty fields on one platform tab from the master listing + AI/deterministic taxonomy resolution.
[**CreateListingMerge**](ListingsApi.md#CreateListingMerge) | **Post** /v1/listings/merge | Merge the same item listed on different marketplaces into one listing; all stay live.
[**CreateListingTagBulk**](ListingsApi.md#CreateListingTagBulk) | **Post** /v1/listings/tags/bulk | Add and/or remove tags on many listings (yours only; max 30 tags a listing, 40 chars a tag). &#x60;dryRun&#x60; reports what would change.
[**CreateListingTagRename**](ListingsApi.md#CreateListingTagRename) | **Post** /v1/listings/tags/rename | Rename a tag on every listing; merges into &#x60;to&#x60; when it already exists.
[**CreateLiveStream**](ListingsApi.md#CreateLiveStream) | **Post** /v1/live-streams | Schedule a live show: title, start time, optional cover, category and the listings to feature (becomes the lineup, in order).
[**CreateLiveStreamConnection**](ListingsApi.md#CreateLiveStreamConnection) | **Post** /v1/live-streams/{id}/connection | The show’s video connection, minted on first call and identical after: RTMP(S) server + stream key, SRT URL, WHIP URL for browser publishing (null where the ingest has none) and the HLS playback URL. configured&#x3D;false (with a reason) when the deployment has no streaming provider.
[**CreateLiveStreamConnectionRotate**](ListingsApi.md#CreateLiveStreamConnectionRotate) | **Post** /v1/live-streams/{id}/connection/rotate | Replace the show’s stream key (e.g. after it was shown on screen). Refused while the show is live.
[**CreateLiveStreamEnd**](ListingsApi.md#CreateLiveStreamEnd) | **Post** /v1/live-streams/{id}/end | End the show. Idempotent.
[**CreateLiveStreamPinLaunch**](ListingsApi.md#CreateLiveStreamPinLaunch) | **Post** /v1/live-streams/{id}/pin-launch | Pin one of your live NFT launches on your live show; viewers mint it from the stream. Returns &#x60;{ ok }&#x60;.
[**CreateLiveStreamStart**](ListingsApi.md#CreateLiveStreamStart) | **Post** /v1/live-streams/{id}/start | Go live: the show’s status becomes live and followers are notified once. Idempotent while live; refused once ended.
[**CreatePassport**](ListingsApi.md#CreatePassport) | **Post** /v1/passports | Register a product passport for a Crossly listing or an inventory item you sell, optionally bound to an NFC tag UID (hex). Active at once. Returns the URL to write to the NFC tag or print as a QR; ownership then follows the item’s sales. For several items with photos and labels, use the drafts flow (POST /passports/drafts).
[**CreatePassportDraft**](ListingsApi.md#CreatePassportDraft) | **Post** /v1/passports/drafts | Start a passport draft for up to 100 inventory items (or Crossly listings): one draft passport per item with its code reserved. Items whose every unit already has a passport come back in skipped.
[**CreatePassportDraftActivate**](ListingsApi.md#CreatePassportDraftActivate) | **Post** /v1/passports/drafts/{id}/activate | Activate every passport of a draft: their pages open for anyone who scans the tag or QR. Closes the draft’s capture link.
[**CreatePassportDraftCaptureLink**](ListingsApi.md#CreatePassportDraftCaptureLink) | **Post** /v1/passports/drafts/{id}/capture-link | A link (and its QR as SVG) any phone opens with no sign-in to add photos, serials and NFC tags to this draft. Open 30 minutes after its last use, 4 hours at most; minting a new one closes the old.
[**CreatePassportPrint**](ListingsApi.md#CreatePassportPrint) | **Post** /v1/passports/print | Print passport labels: sent to your printer agent (the Crossly desktop app includes one) when it is online and prints the template, else a signed PDF link with the reason. Templates: sticker-2in, sticker-1in, label-4x6, hang-tag, sheet-letter, sheet-a4, transfer-card.
[**CreatePassportRevoke**](ListingsApi.md#CreatePassportRevoke) | **Post** /v1/passports/{id}/revoke | Revoke a passport (a lost or misprinted tag). Refused once the item has sold — the passport is then part of its history.
[**CreatePassportTag**](ListingsApi.md#CreatePassportTag) | **Post** /v1/passports/{id}/tag | Bind (or re-bind) a passport to the NFC tag you wrote its nfcUrl to: the tag UID (hex), and whether you locked the tag read-only. Refused once the item has sold, for a locked tag, or a tag already on another item.
[**CreateRentalBulkRentable**](ListingsApi.md#CreateRentalBulkRentable) | **Post** /v1/rentals/bulk-rentable | Make up to 200 Crossly listings rentable with the same options (a day rate, or a share of each price).
[**CreateShopLimitRequest**](ListingsApi.md#CreateShopLimitRequest) | **Post** /v1/shop-limits/requests | Ask for a higher Crossly Shop selling limit.
[**CreateStory**](ListingsApi.md#CreateStory) | **Post** /v1/stories | Share a story: 1–10 slides played in order. Each slide is a photo or video plus optional overlays (positions are fractions of the 9:16 frame). Listing stickers must be your own Crossly listings and also tag the slide; @mentions must be real shops. audience: everyone | followers; expiresInHours: 6, 12, 24 or 48.
[**CreateVariationGroup**](ListingsApi.md#CreateVariationGroup) | **Post** /v1/variation-groups | Group listings you already have: a variation (one product, an axis such as Size) or a purchase-rule group (\&quot;buy any 3\&quot;).
[**CreateVariationGroupGenerate**](ListingsApi.md#CreateVariationGroupGenerate) | **Post** /v1/variation-groups/generate | Turn one listing into a variation group: each new value becomes its own listing (a copy, with its own price, quantity and SKU) and all of them one group.
[**CreateVariationGroupGenerateMatrix**](ListingsApi.md#CreateVariationGroupGenerateMatrix) | **Post** /v1/variation-groups/generate-matrix | Turn one listing into a variation group over up to three axes (Size × Color): each other cell is a listing you already have (listingId) or a new copy of this one with its own price, quantity and SKU.
[**CreateVariationGroupPublish**](ListingsApi.md#CreateVariationGroupPublish) | **Post** /v1/variation-groups/{id}/publish | Publish a variation group per platform: one listing with a variation picker where the platform supports it (eBay, Shopify, Etsy), one per option elsewhere. Each platform reports its own outcome.
[**CreateVideoStudioFinalize**](ListingsApi.md#CreateVideoStudioFinalize) | **Post** /v1/video-studio/finalize | Check an uploaded listing video (exists, MP4, size) and, with listingId, set it as that listing&#39;s video.
[**CreateVideoStudioUploadUrl**](ListingsApi.md#CreateVideoStudioUploadUrl) | **Post** /v1/video-studio/upload-url | Presign an R2 PUT (video/mp4, 15 min) for a listing video; finalize it after the upload.
[**DeleteEvent**](ListingsApi.md#DeleteEvent) | **Delete** /v1/events/{id} | Delete a draft, or cancel a published event (its page then says it was called off).
[**DeleteListing**](ListingsApi.md#DeleteListing) | **Delete** /v1/listings/{id} | Delist a listing (optionally narrowed to specific platforms via ?platforms&#x3D;).
[**DeleteListingTag**](ListingsApi.md#DeleteListingTag) | **Delete** /v1/listings/tags/{tag} | Remove a tag (any case) from every listing.
[**DeleteLiveStreamPinLaunch**](ListingsApi.md#DeleteLiveStreamPinLaunch) | **Delete** /v1/live-streams/{id}/pin-launch | Unpin the launch from your live show. Returns &#x60;{ ok }&#x60;.
[**DeleteRentalOffer**](ListingsApi.md#DeleteRentalOffer) | **Delete** /v1/rentals/offers/{platformListingId} | Stop offering a listing for rent (existing bookings stand).
[**DeleteStory**](ListingsApi.md#DeleteStory) | **Delete** /v1/stories/{id} | Take one story slide down before it expires.
[**DeleteVariationGroup**](ListingsApi.md#DeleteVariationGroup) | **Delete** /v1/variation-groups/{id} | Dissolve a group. Its options stay as ordinary listings.
[**GetEvent**](ListingsApi.md#GetEvent) | **Get** /v1/events/{id} | One of your events.
[**GetEventAnalytic**](ListingsApi.md#GetEventAnalytic) | **Get** /v1/events/{id}/analytics | Views, RSVPs, and what the linked listings sold on Crossly during the event (orders, units, revenue in cents).
[**GetListing**](ListingsApi.md#GetListing) | **Get** /v1/listings/{id} | Get one listing with its platform rows.
[**GetListingFacet**](ListingsApi.md#GetListingFacet) | **Get** /v1/listings/facets | Distinct brands + categories across listings + inventory.
[**GetListingSkuExist**](ListingsApi.md#GetListingSkuExist) | **Get** /v1/listings/sku-exists | Check whether a SKU is already used by one of this user&#39;s items.
[**GetLiveStreamRecap**](ListingsApi.md#GetLiveStreamRecap) | **Get** /v1/live-streams/{id}/recap | The show’s numbers: duration, peak viewers, items sold and gross, tips, chat messages, auctions run, lineup size and the replay URL.
[**GetPassport**](ListingsApi.md#GetPassport) | **Get** /v1/passports/{id} | One passport: its item and every listing of it, the event history (created, printed, tag bound, transferred, claimed…), scans per day for 30 days, custody, grading.
[**GetPassportDraft**](ListingsApi.md#GetPassportDraft) | **Get** /v1/passports/drafts/{id} | A passport draft and its passports (photos, serial, tag and label state of each).
[**GetPassportPrint**](ListingsApi.md#GetPassportPrint) | **Get** /v1/passports/print/{jobId} | Where a passport print job sent to your agent is: queued, sent, done, failed, expired or duplicate.
[**GetPricingSetting**](ListingsApi.md#GetPricingSetting) | **Get** /v1/pricing/settings | Selling options: own-price match (Crossly sells your items at your own lower price for the same item on another marketplace; dollar-priced copies only) and layaway (buyers may put items on layaway). Both on by default.
[**GetPricingSoldComp**](ListingsApi.md#GetPricingSoldComp) | **Get** /v1/pricing/sold-comps | What one catalog product (market_sku_id, in one condition) actually sold for across Crossly in the last 180 days: median, p25/p75, p10/p90 (null under 10 sales), sale count, last sold, and days from listing to sale (median, p25, p75). Pass marketSkuId (+ optional condition) or one of your listingIds. comps is null until there are 3+ sales from 2+ sellers; holdCode says why there is nothing (no_market_sku, mixed_inventory, below_privacy_floor) — the same reasons a sold_median reprice rule holds.
[**GetRentalSetting**](ListingsApi.md#GetRentalSetting) | **Get** /v1/rentals/settings | The store’s rental defaults (every item inherits them) and its payout timing.
[**GetShopLimit**](ListingsApi.md#GetShopLimit) | **Get** /v1/shop-limits | Your Crossly Shop selling limit: live listings and listed value, used of limit, per store, and the next step.
[**ListEventAttendees**](ListingsApi.md#ListEventAttendees) | **Get** /v1/events/{id}/attendees | Who RSVPd: display name, avatar, going or interested, and when. No contact details.
[**ListEvents**](ListingsApi.md#ListEvents) | **Get** /v1/events | Your events (launches, sales, live shows, drops, meetups, pop-ups, workshops, parties): schedule, status, phase, and going/interested counts.
[**ListListingAiDetections**](ListingsApi.md#ListListingAiDetections) | **Get** /v1/listings/{id}/ai-detections | Crossly&#39;s AI-image detections for a listing&#39;s photos.
[**ListListingDiscrepancies**](ListingsApi.md#ListListingDiscrepancies) | **Get** /v1/listings/{id}/discrepancies | List detected marketplace-drift discrepancies for a listing.
[**ListListingIds**](ListingsApi.md#ListListingIds) | **Get** /v1/listings/ids | Filter listings → return matching id list (no pagination).
[**ListListingTagSuggest**](ListingsApi.md#ListListingTagSuggest) | **Get** /v1/listings/tags/suggest | Deterministic tag suggestions for one listing (brand, category, size, condition, colors, style tags, era words in the title); &#x60;known&#x60; &#x3D; you already use the tag.
[**ListListingTags**](ListingsApi.md#ListListingTags) | **Get** /v1/listings/tags | Your listing tags with how many listings carry each, most used first (case-insensitive; &#x60;q&#x60; filters).
[**ListListings**](ListingsApi.md#ListListings) | **Get** /v1/listings | List active platform listings.
[**ListLiveStreams**](ListingsApi.md#ListLiveStreams) | **Get** /v1/live-streams | Your live shows, newest scheduled first: status, cover, category, connection method and peak viewers.
[**ListPassportEvents**](ListingsApi.md#ListPassportEvents) | **Get** /v1/passports/{id}/events | A passport’s event history, newest first (never a buyer or a price).
[**ListPassportPrintTargets**](ListingsApi.md#ListPassportPrintTargets) | **Get** /v1/passports/print-targets | Your printer agents (online, printers, default) and the passport label templates each can print.
[**ListPassports**](ListingsApi.md#ListPassports) | **Get** /v1/passports | Product passports you registered (optionally for one Crossly listing): the code and URL each tag carries.
[**ListRentalOffers**](ListingsApi.md#ListRentalOffers) | **Get** /v1/rentals/offers | What you offer for rent.
[**ListStories**](ListingsApi.md#ListStories) | **Get** /v1/stories | Your active story slides (newest run first, each run in play order), with their overlays: text, listing stickers, @mentions, links, emoji and drawing.
[**UpdateEvent**](ListingsApi.md#UpdateEvent) | **Patch** /v1/events/{id} | Edit an event. Omitted fields are unchanged; null clears.
[**UpdateHotspot**](ListingsApi.md#UpdateHotspot) | **Put** /v1/hotspots | Tag your listings at points on a story or lookbook photo (shop the look). Replaces that photo’s tags; at most 12.
[**UpdateListing**](ListingsApi.md#UpdateListing) | **Patch** /v1/listings/{id} | Edit a listing and fan out update jobs to existing platform listings.
[**UpdateListingCustom**](ListingsApi.md#UpdateListingCustom) | **Put** /v1/listings/{id}/customs | Set a listing’s customs: HS code (6–10 digits), country of origin, and countries this item must not ship to on Crossly. Omitted fields are left alone; null (or [] for restrictedCountries) clears.
[**UpdateListingPlatformListing**](ListingsApi.md#UpdateListingPlatformListing) | **Patch** /v1/listings/{id}/platform-listings/{plId} | Set one marketplace copy&#39;s price (e.g. the ebay.com copy of a listing also live on ebay.co.uk), in that copy&#39;s own currency.
[**UpdateLiveStream**](ListingsApi.md#UpdateLiveStream) | **Patch** /v1/live-streams/{id} | Edit a show that has not ended: title, description, cover, category, start time (not once live) or connection method (browser | phone | encoder).
[**UpdateLiveStreamLineup**](ListingsApi.md#UpdateLiveStreamLineup) | **Put** /v1/live-streams/{id}/lineup | Set a live show’s lineup: your Crossly listings, in the order you will sell them. Buyers see it before they join; a sale of a lineup item while the show is live is recorded at its replay moment.
[**UpdatePassport**](ListingsApi.md#UpdatePassport) | **Patch** /v1/passports/{id} | Change what a passport says (title, condition, maker, year, materials, notes) or add photo URLs (photos append: a passport keeps the ones it was registered with).
[**UpdatePassportDraft**](ListingsApi.md#UpdatePassportDraft) | **Patch** /v1/passports/drafts/{id} | Move a draft to a step, pick its label template, add items or remove draft passports.
[**UpdatePricingSetting**](ListingsApi.md#UpdatePricingSetting) | **Put** /v1/pricing/settings | Turn own-price match and/or layaway on or off for every one of your stores. One store: PUT /v1/storefronts/{id}/settings.
[**UpdateRentalOffer**](ListingsApi.md#UpdateRentalOffer) | **Put** /v1/rentals/offers/{platformListingId} | Offer a Crossly listing for rent: day rate, optional week rate, deposit, min/max days, turnaround buffer, rent-only (off sale while offered), seller approval, unavailable date ranges (blackoutRanges, inclusive YYYY-MM-DD) and a return policy. Keeps the listing’s own rental terms in step. Needs a ship-from address.
[**UpdateRentalSetting**](ListingsApi.md#UpdateRentalSetting) | **Put** /v1/rentals/settings | Set the store’s rental defaults and payout timing (on_start needs 3 rentals that came back).
[**UpdateVariationGroup**](ListingsApi.md#UpdateVariationGroup) | **Patch** /v1/variation-groups/{id} | Retitle a group, rename its axis, or change a group&#39;s purchase rule (written into every member&#39;s description).



## CreateAiDetectionDispute

> CreateAiDetectionDisputeResponse CreateAiDetectionDispute(ctx, id).Execute()

Dispute an AI-image detection on your photo (sends it to Crossly review).

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
    resp, r, err := apiClient.ListingsApi.CreateAiDetectionDispute(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateAiDetectionDispute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiDetectionDispute`: CreateAiDetectionDisputeResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateAiDetectionDispute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiDetectionDisputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateAiDetectionDisputeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateEvent

> EventResponse CreateEvent(ctx).Execute()

Create an event. A draft unless publish:true. Linked listings, the live show, the drop and the discount code must all be yours.

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
    resp, r, err := apiClient.ListingsApi.CreateEvent(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEvent`: EventResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateEvent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventRequest struct via the builder pattern


### Return type

**EventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateEventPublish

> EventResponse CreateEventPublish(ctx, id).Execute()

Publish an event: its page goes live, and with notifyFollowers your followers are told once.

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
    resp, r, err := apiClient.ListingsApi.CreateEventPublish(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateEventPublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEventPublish`: EventResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateEventPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**EventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


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

**CreateListingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingBulk

> CreateListingBulkResponse CreateListingBulk(ctx).Execute()

Create up to 50 listings in one call, with a result per item.



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
    resp, r, err := apiClient.ListingsApi.CreateListingBulk(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingBulk`: CreateListingBulkResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingBulk`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingBulkRequest struct via the builder pattern


### Return type

**CreateListingBulkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingBulkCheckStatus

> CreateListingBulkDeleteResponse CreateListingBulkCheckStatus(ctx).Execute()

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
    // response from `CreateListingBulkCheckStatus`: CreateListingBulkDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingBulkCheckStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingBulkCheckStatusRequest struct via the builder pattern


### Return type

**CreateListingBulkDeleteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingBulkCrosspost

> CreateListingBulkDeleteResponse CreateListingBulkCrosspost(ctx).Execute()

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
    // response from `CreateListingBulkCrosspost`: CreateListingBulkDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingBulkCrosspost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingBulkCrosspostRequest struct via the builder pattern


### Return type

**CreateListingBulkDeleteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**CreateListingBulkDeleteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingBulkDelist

> CreateListingBulkDeleteResponse CreateListingBulkDelist(ctx).Execute()

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
    // response from `CreateListingBulkDelist`: CreateListingBulkDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingBulkDelist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingBulkDelistRequest struct via the builder pattern


### Return type

**CreateListingBulkDeleteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**CreateListingBulkDelistPreviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingBulkHardDelete

> CreateListingBulkHardDeleteResponse CreateListingBulkHardDelete(ctx).Execute()

Permanently delete listings (live ones need allowLive)

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

**CreateListingBulkHardDeleteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingBulkRelist

> CreateListingBulkDeleteResponse CreateListingBulkRelist(ctx).Execute()

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
    // response from `CreateListingBulkRelist`: CreateListingBulkDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingBulkRelist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingBulkRelistRequest struct via the builder pattern


### Return type

**CreateListingBulkDeleteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**CreateListingBulkUpdateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**CreateListingByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**CreateListingCheckDuplicateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**CreateListingCombineResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**CreateListingDiscrepancyResolveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**CreateListingImportByUrlResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**CreateListingMagicFillResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingMerge

> CreateListingMergeResponse CreateListingMerge(ctx).Execute()

Merge the same item listed on different marketplaces into one listing; all stay live.

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
    resp, r, err := apiClient.ListingsApi.CreateListingMerge(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingMerge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingMerge`: CreateListingMergeResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingMerge`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingMergeRequest struct via the builder pattern


### Return type

**CreateListingMergeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingTagBulk

> CreateListingTagBulkResponse CreateListingTagBulk(ctx).Execute()

Add and/or remove tags on many listings (yours only; max 30 tags a listing, 40 chars a tag). `dryRun` reports what would change.

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
    resp, r, err := apiClient.ListingsApi.CreateListingTagBulk(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingTagBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingTagBulk`: CreateListingTagBulkResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingTagBulk`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingTagBulkRequest struct via the builder pattern


### Return type

**CreateListingTagBulkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingTagRename

> CreateListingTagRenameResponse CreateListingTagRename(ctx).Execute()

Rename a tag on every listing; merges into `to` when it already exists.

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
    resp, r, err := apiClient.ListingsApi.CreateListingTagRename(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateListingTagRename``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingTagRename`: CreateListingTagRenameResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateListingTagRename`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingTagRenameRequest struct via the builder pattern


### Return type

**CreateListingTagRenameResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLiveStream

> LiveStreamResponse CreateLiveStream(ctx).Execute()

Schedule a live show: title, start time, optional cover, category and the listings to feature (becomes the lineup, in order).

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
    resp, r, err := apiClient.ListingsApi.CreateLiveStream(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateLiveStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLiveStream`: LiveStreamResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateLiveStream`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLiveStreamRequest struct via the builder pattern


### Return type

**LiveStreamResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLiveStreamConnection

> CreateLiveStreamConnectionResponse CreateLiveStreamConnection(ctx, id).Execute()

The show’s video connection, minted on first call and identical after: RTMP(S) server + stream key, SRT URL, WHIP URL for browser publishing (null where the ingest has none) and the HLS playback URL. configured=false (with a reason) when the deployment has no streaming provider.

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
    resp, r, err := apiClient.ListingsApi.CreateLiveStreamConnection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateLiveStreamConnection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLiveStreamConnection`: CreateLiveStreamConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateLiveStreamConnection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLiveStreamConnectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateLiveStreamConnectionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLiveStreamConnectionRotate

> CreateLiveStreamConnectionResponse CreateLiveStreamConnectionRotate(ctx, id).Execute()

Replace the show’s stream key (e.g. after it was shown on screen). Refused while the show is live.

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
    resp, r, err := apiClient.ListingsApi.CreateLiveStreamConnectionRotate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateLiveStreamConnectionRotate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLiveStreamConnectionRotate`: CreateLiveStreamConnectionResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateLiveStreamConnectionRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLiveStreamConnectionRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateLiveStreamConnectionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLiveStreamEnd

> LiveStreamResponse CreateLiveStreamEnd(ctx, id).Execute()

End the show. Idempotent.

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
    resp, r, err := apiClient.ListingsApi.CreateLiveStreamEnd(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateLiveStreamEnd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLiveStreamEnd`: LiveStreamResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateLiveStreamEnd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLiveStreamEndRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**LiveStreamResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLiveStreamPinLaunch

> OkResponse CreateLiveStreamPinLaunch(ctx, id).Execute()

Pin one of your live NFT launches on your live show; viewers mint it from the stream. Returns `{ ok }`.

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
    resp, r, err := apiClient.ListingsApi.CreateLiveStreamPinLaunch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateLiveStreamPinLaunch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLiveStreamPinLaunch`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateLiveStreamPinLaunch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLiveStreamPinLaunchRequest struct via the builder pattern


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


## CreateLiveStreamStart

> LiveStreamResponse CreateLiveStreamStart(ctx, id).Execute()

Go live: the show’s status becomes live and followers are notified once. Idempotent while live; refused once ended.

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
    resp, r, err := apiClient.ListingsApi.CreateLiveStreamStart(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateLiveStreamStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLiveStreamStart`: LiveStreamResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateLiveStreamStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLiveStreamStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**LiveStreamResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreatePassport

> PassportResponse CreatePassport(ctx).Execute()

Register a product passport for a Crossly listing or an inventory item you sell, optionally bound to an NFC tag UID (hex). Active at once. Returns the URL to write to the NFC tag or print as a QR; ownership then follows the item’s sales. For several items with photos and labels, use the drafts flow (POST /passports/drafts).

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
    resp, r, err := apiClient.ListingsApi.CreatePassport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreatePassport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePassport`: PassportResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreatePassport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePassportRequest struct via the builder pattern


### Return type

**PassportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreatePassportDraft

> PassportDraftResponse CreatePassportDraft(ctx).Execute()

Start a passport draft for up to 100 inventory items (or Crossly listings): one draft passport per item with its code reserved. Items whose every unit already has a passport come back in skipped.

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
    resp, r, err := apiClient.ListingsApi.CreatePassportDraft(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreatePassportDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePassportDraft`: PassportDraftResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreatePassportDraft`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePassportDraftRequest struct via the builder pattern


### Return type

**PassportDraftResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreatePassportDraftActivate

> PassportDraftResponse CreatePassportDraftActivate(ctx, id).Execute()

Activate every passport of a draft: their pages open for anyone who scans the tag or QR. Closes the draft’s capture link.

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
    resp, r, err := apiClient.ListingsApi.CreatePassportDraftActivate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreatePassportDraftActivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePassportDraftActivate`: PassportDraftResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreatePassportDraftActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePassportDraftActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**PassportDraftResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreatePassportDraftCaptureLink

> CreatePassportDraftCaptureLinkResponse CreatePassportDraftCaptureLink(ctx, id).Execute()

A link (and its QR as SVG) any phone opens with no sign-in to add photos, serials and NFC tags to this draft. Open 30 minutes after its last use, 4 hours at most; minting a new one closes the old.

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
    resp, r, err := apiClient.ListingsApi.CreatePassportDraftCaptureLink(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreatePassportDraftCaptureLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePassportDraftCaptureLink`: CreatePassportDraftCaptureLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreatePassportDraftCaptureLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePassportDraftCaptureLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatePassportDraftCaptureLinkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreatePassportPrint

> CreatePassportPrintResponse CreatePassportPrint(ctx).Execute()

Print passport labels: sent to your printer agent (the Crossly desktop app includes one) when it is online and prints the template, else a signed PDF link with the reason. Templates: sticker-2in, sticker-1in, label-4x6, hang-tag, sheet-letter, sheet-a4, transfer-card.

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
    resp, r, err := apiClient.ListingsApi.CreatePassportPrint(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreatePassportPrint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePassportPrint`: CreatePassportPrintResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreatePassportPrint`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePassportPrintRequest struct via the builder pattern


### Return type

**CreatePassportPrintResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreatePassportRevoke

> PassportResponse CreatePassportRevoke(ctx, id).Execute()

Revoke a passport (a lost or misprinted tag). Refused once the item has sold — the passport is then part of its history.

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
    resp, r, err := apiClient.ListingsApi.CreatePassportRevoke(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreatePassportRevoke``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePassportRevoke`: PassportResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreatePassportRevoke`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePassportRevokeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**PassportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreatePassportTag

> PassportResponse CreatePassportTag(ctx, id).Execute()

Bind (or re-bind) a passport to the NFC tag you wrote its nfcUrl to: the tag UID (hex), and whether you locked the tag read-only. Refused once the item has sold, for a locked tag, or a tag already on another item.

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
    resp, r, err := apiClient.ListingsApi.CreatePassportTag(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreatePassportTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePassportTag`: PassportResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreatePassportTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePassportTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**PassportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateRentalBulkRentable

> CreateRentalBulkRentableResponse CreateRentalBulkRentable(ctx).Execute()

Make up to 200 Crossly listings rentable with the same options (a day rate, or a share of each price).

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
    resp, r, err := apiClient.ListingsApi.CreateRentalBulkRentable(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateRentalBulkRentable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRentalBulkRentable`: CreateRentalBulkRentableResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateRentalBulkRentable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRentalBulkRentableRequest struct via the builder pattern


### Return type

**CreateRentalBulkRentableResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateShopLimitRequest

> CreateShopLimitRequestResponse CreateShopLimitRequest(ctx).CreateShopLimitRequestRequest(createShopLimitRequestRequest).Execute()

Ask for a higher Crossly Shop selling limit.



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
    createShopLimitRequestRequest := *openapiclient.NewCreateShopLimitRequestRequest("Reason_example") // CreateShopLimitRequestRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsApi.CreateShopLimitRequest(context.Background()).CreateShopLimitRequestRequest(createShopLimitRequestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateShopLimitRequest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateShopLimitRequest`: CreateShopLimitRequestResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateShopLimitRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateShopLimitRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createShopLimitRequestRequest** | **CreateShopLimitRequestRequest** |  | 

### Return type

**CreateShopLimitRequestResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStory

> CreateStoryResponse CreateStory(ctx).Execute()

Share a story: 1–10 slides played in order. Each slide is a photo or video plus optional overlays (positions are fractions of the 9:16 frame). Listing stickers must be your own Crossly listings and also tag the slide; @mentions must be real shops. audience: everyone | followers; expiresInHours: 6, 12, 24 or 48.

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
    resp, r, err := apiClient.ListingsApi.CreateStory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateStory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStory`: CreateStoryResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateStory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStoryRequest struct via the builder pattern


### Return type

**CreateStoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVariationGroup

> CreateVariationGroupResponse CreateVariationGroup(ctx).Execute()

Group listings you already have: a variation (one product, an axis such as Size) or a purchase-rule group (\"buy any 3\").

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
    resp, r, err := apiClient.ListingsApi.CreateVariationGroup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateVariationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVariationGroup`: CreateVariationGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateVariationGroup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVariationGroupRequest struct via the builder pattern


### Return type

**CreateVariationGroupResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVariationGroupGenerate

> CreateVariationGroupGenerateResponse CreateVariationGroupGenerate(ctx).Execute()

Turn one listing into a variation group: each new value becomes its own listing (a copy, with its own price, quantity and SKU) and all of them one group.

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
    resp, r, err := apiClient.ListingsApi.CreateVariationGroupGenerate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateVariationGroupGenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVariationGroupGenerate`: CreateVariationGroupGenerateResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateVariationGroupGenerate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVariationGroupGenerateRequest struct via the builder pattern


### Return type

**CreateVariationGroupGenerateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVariationGroupGenerateMatrix

> CreateVariationGroupGenerateMatrixResponse CreateVariationGroupGenerateMatrix(ctx).Execute()

Turn one listing into a variation group over up to three axes (Size × Color): each other cell is a listing you already have (listingId) or a new copy of this one with its own price, quantity and SKU.

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
    resp, r, err := apiClient.ListingsApi.CreateVariationGroupGenerateMatrix(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateVariationGroupGenerateMatrix``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVariationGroupGenerateMatrix`: CreateVariationGroupGenerateMatrixResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateVariationGroupGenerateMatrix`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVariationGroupGenerateMatrixRequest struct via the builder pattern


### Return type

**CreateVariationGroupGenerateMatrixResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVariationGroupPublish

> CreateVariationGroupPublishResponse CreateVariationGroupPublish(ctx, id).Execute()

Publish a variation group per platform: one listing with a variation picker where the platform supports it (eBay, Shopify, Etsy), one per option elsewhere. Each platform reports its own outcome.

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
    resp, r, err := apiClient.ListingsApi.CreateVariationGroupPublish(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateVariationGroupPublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVariationGroupPublish`: CreateVariationGroupPublishResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateVariationGroupPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVariationGroupPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVariationGroupPublishResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVideoStudioFinalize

> CreateVideoStudioFinalizeResponse CreateVideoStudioFinalize(ctx).Execute()

Check an uploaded listing video (exists, MP4, size) and, with listingId, set it as that listing's video.

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
    resp, r, err := apiClient.ListingsApi.CreateVideoStudioFinalize(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateVideoStudioFinalize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVideoStudioFinalize`: CreateVideoStudioFinalizeResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateVideoStudioFinalize`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVideoStudioFinalizeRequest struct via the builder pattern


### Return type

**CreateVideoStudioFinalizeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVideoStudioUploadUrl

> CreateVideoStudioUploadUrlResponse CreateVideoStudioUploadUrl(ctx).Execute()

Presign an R2 PUT (video/mp4, 15 min) for a listing video; finalize it after the upload.

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
    resp, r, err := apiClient.ListingsApi.CreateVideoStudioUploadUrl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.CreateVideoStudioUploadUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVideoStudioUploadUrl`: CreateVideoStudioUploadUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.CreateVideoStudioUploadUrl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVideoStudioUploadUrlRequest struct via the builder pattern


### Return type

**CreateVideoStudioUploadUrlResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteEvent

> DeleteEventResponse DeleteEvent(ctx, id).Execute()

Delete a draft, or cancel a published event (its page then says it was called off).

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
    resp, r, err := apiClient.ListingsApi.DeleteEvent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.DeleteEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEvent`: DeleteEventResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.DeleteEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteEventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteListing

> DeleteListingResponse DeleteListing(ctx, id).Platforms(platforms).PlatformListingIds(platformListingIds).Execute()

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
    platformListingIds := "platformListingIds_example" // string | Comma-separated platform listing ids — end only those copies (e.g. one eBay site's copy). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsApi.DeleteListing(context.Background(), id).Platforms(platforms).PlatformListingIds(platformListingIds).Execute()
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
 **platformListingIds** | **string** | Comma-separated platform listing ids — end only those copies (e.g. one eBay site&#39;s copy). | 

### Return type

**DeleteListingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteListingTag

> UpdatedResponse DeleteListingTag(ctx, tag).Execute()

Remove a tag (any case) from every listing.

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
    tag := "tag_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsApi.DeleteListingTag(context.Background(), tag).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.DeleteListingTag``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteListingTag`: UpdatedResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.DeleteListingTag`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tag** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteListingTagRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdatedResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteLiveStreamPinLaunch

> OkResponse DeleteLiveStreamPinLaunch(ctx, id).Execute()

Unpin the launch from your live show. Returns `{ ok }`.

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
    resp, r, err := apiClient.ListingsApi.DeleteLiveStreamPinLaunch(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.DeleteLiveStreamPinLaunch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLiveStreamPinLaunch`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.DeleteLiveStreamPinLaunch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLiveStreamPinLaunchRequest struct via the builder pattern


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


## DeleteRentalOffer

> UpdateGroupBuyResponse DeleteRentalOffer(ctx, platformListingId).Execute()

Stop offering a listing for rent (existing bookings stand).

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
    platformListingId := "platformListingId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsApi.DeleteRentalOffer(context.Background(), platformListingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.DeleteRentalOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteRentalOffer`: UpdateGroupBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.DeleteRentalOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformListingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRentalOfferRequest struct via the builder pattern


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


## DeleteStory

> DeleteStoryResponse DeleteStory(ctx, id).Execute()

Take one story slide down before it expires.

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
    resp, r, err := apiClient.ListingsApi.DeleteStory(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.DeleteStory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStory`: DeleteStoryResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.DeleteStory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStoryRequest struct via the builder pattern


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


## DeleteVariationGroup

> UpdateGroupBuyResponse DeleteVariationGroup(ctx, id).Execute()

Dissolve a group. Its options stay as ordinary listings.

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
    resp, r, err := apiClient.ListingsApi.DeleteVariationGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.DeleteVariationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteVariationGroup`: UpdateGroupBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.DeleteVariationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteVariationGroupRequest struct via the builder pattern


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


## GetEvent

> EventResponse GetEvent(ctx, id).Execute()

One of your events.

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
    resp, r, err := apiClient.ListingsApi.GetEvent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEvent`: EventResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**EventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetEventAnalytic

> GetEventAnalyticResponse GetEventAnalytic(ctx, id).Execute()

Views, RSVPs, and what the linked listings sold on Crossly during the event (orders, units, revenue in cents).

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
    resp, r, err := apiClient.ListingsApi.GetEventAnalytic(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetEventAnalytic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventAnalytic`: GetEventAnalyticResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetEventAnalytic`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEventAnalyticRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetEventAnalyticResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**GetListingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**GetListingFacetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**GetListingSkuExistResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetLiveStreamRecap

> GetLiveStreamRecapResponse GetLiveStreamRecap(ctx, id).Execute()

The show’s numbers: duration, peak viewers, items sold and gross, tips, chat messages, auctions run, lineup size and the replay URL.

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
    resp, r, err := apiClient.ListingsApi.GetLiveStreamRecap(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetLiveStreamRecap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLiveStreamRecap`: GetLiveStreamRecapResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetLiveStreamRecap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLiveStreamRecapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetLiveStreamRecapResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetPassport

> GetPassportResponse GetPassport(ctx, id).Execute()

One passport: its item and every listing of it, the event history (created, printed, tag bound, transferred, claimed…), scans per day for 30 days, custody, grading.

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
    resp, r, err := apiClient.ListingsApi.GetPassport(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetPassport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPassport`: GetPassportResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetPassport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPassportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetPassportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetPassportDraft

> PassportDraftResponse GetPassportDraft(ctx, id).Execute()

A passport draft and its passports (photos, serial, tag and label state of each).

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
    resp, r, err := apiClient.ListingsApi.GetPassportDraft(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetPassportDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPassportDraft`: PassportDraftResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetPassportDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPassportDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**PassportDraftResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetPassportPrint

> GetPassportPrintResponse GetPassportPrint(ctx, jobId).Execute()

Where a passport print job sent to your agent is: queued, sent, done, failed, expired or duplicate.

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
    jobId := "jobId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsApi.GetPassportPrint(context.Background(), jobId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetPassportPrint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPassportPrint`: GetPassportPrintResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetPassportPrint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**jobId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPassportPrintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetPassportPrintResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetPricingSetting

> GetPricingSettingResponse GetPricingSetting(ctx).Execute()

Selling options: own-price match (Crossly sells your items at your own lower price for the same item on another marketplace; dollar-priced copies only) and layaway (buyers may put items on layaway). Both on by default.

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
    resp, r, err := apiClient.ListingsApi.GetPricingSetting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetPricingSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPricingSetting`: GetPricingSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetPricingSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingSettingRequest struct via the builder pattern


### Return type

**GetPricingSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetPricingSoldComp

> GetPricingSoldCompResponse GetPricingSoldComp(ctx).Execute()

What one catalog product (market_sku_id, in one condition) actually sold for across Crossly in the last 180 days: median, p25/p75, p10/p90 (null under 10 sales), sale count, last sold, and days from listing to sale (median, p25, p75). Pass marketSkuId (+ optional condition) or one of your listingIds. comps is null until there are 3+ sales from 2+ sellers; holdCode says why there is nothing (no_market_sku, mixed_inventory, below_privacy_floor) — the same reasons a sold_median reprice rule holds.

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
    resp, r, err := apiClient.ListingsApi.GetPricingSoldComp(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetPricingSoldComp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPricingSoldComp`: GetPricingSoldCompResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetPricingSoldComp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPricingSoldCompRequest struct via the builder pattern


### Return type

**GetPricingSoldCompResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetRentalSetting

> GetRentalSettingResponse GetRentalSetting(ctx).Execute()

The store’s rental defaults (every item inherits them) and its payout timing.

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
    resp, r, err := apiClient.ListingsApi.GetRentalSetting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetRentalSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRentalSetting`: GetRentalSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetRentalSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRentalSettingRequest struct via the builder pattern


### Return type

**GetRentalSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetShopLimit

> GetShopLimitResponse GetShopLimit(ctx).Execute()

Your Crossly Shop selling limit: live listings and listed value, used of limit, per store, and the next step.



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
    resp, r, err := apiClient.ListingsApi.GetShopLimit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.GetShopLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetShopLimit`: GetShopLimitResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.GetShopLimit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetShopLimitRequest struct via the builder pattern


### Return type

**GetShopLimitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListEventAttendees

> ListEventAttendeesResponse ListEventAttendees(ctx, id).Execute()

Who RSVPd: display name, avatar, going or interested, and when. No contact details.

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
    resp, r, err := apiClient.ListingsApi.ListEventAttendees(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListEventAttendees``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEventAttendees`: ListEventAttendeesResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListEventAttendees`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEventAttendeesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListEventAttendeesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListEvents

> ListEventsResponse ListEvents(ctx).Execute()

Your events (launches, sales, live shows, drops, meetups, pop-ups, workshops, parties): schedule, status, phase, and going/interested counts.

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
    resp, r, err := apiClient.ListingsApi.ListEvents(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEvents`: ListEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEventsRequest struct via the builder pattern


### Return type

**ListEventsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListListingAiDetections

> ListListingAiDetectionsResponse ListListingAiDetections(ctx, id).Execute()

Crossly's AI-image detections for a listing's photos.

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
    resp, r, err := apiClient.ListingsApi.ListListingAiDetections(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListListingAiDetections``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListListingAiDetections`: ListListingAiDetectionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListListingAiDetections`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListListingAiDetectionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListListingAiDetectionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListListingDiscrepancies

> ListListingDiscrepanciesResponse ListListingDiscrepancies(ctx, id).Execute()

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
    // response from `ListListingDiscrepancies`: ListListingDiscrepanciesResponse
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

**ListListingDiscrepanciesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListListingIds

> ListListingIdsResponse ListListingIds(ctx).Execute()

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
    // response from `ListListingIds`: ListListingIdsResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListListingIds`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListListingIdsRequest struct via the builder pattern


### Return type

**ListListingIdsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListListingTagSuggest

> ListListingTagSuggestResponse ListListingTagSuggest(ctx).ListingId(listingId).Execute()

Deterministic tag suggestions for one listing (brand, category, size, condition, colors, style tags, era words in the title); `known` = you already use the tag.

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
    listingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsApi.ListListingTagSuggest(context.Background()).ListingId(listingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListListingTagSuggest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListListingTagSuggest`: ListListingTagSuggestResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListListingTagSuggest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListListingTagSuggestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **listingId** | **string** |  | 

### Return type

**ListListingTagSuggestResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListListingTags

> ListListingTagsResponse ListListingTags(ctx).Q(q).Limit(limit).Execute()

Your listing tags with how many listings carry each, most used first (case-insensitive; `q` filters).

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
    q := "q_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsApi.ListListingTags(context.Background()).Q(q).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListListingTags``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListListingTags`: ListListingTagsResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListListingTags`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListListingTagsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

**ListListingTagsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListListings

> ListListingsResponse ListListings(ctx).Page(page).Limit(limit).Platform(platform).Status(status).Sku(sku).UpdatedSince(updatedSince).Kind(kind).Tag(tag).StillLive(stillLive).Execute()

List active platform listings.



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
    page := int32(56) // int32 |  (optional) (default to 1)
    limit := int32(56) // int32 |  (optional) (default to 25)
    platform := "platform_example" // string |  (optional)
    status := "status_example" // string |  (optional)
    sku := "sku_example" // string | Exact SKU match: look up the row a supplier product maps to. (optional)
    updatedSince := time.Now() // time.Time | Only rows created or changed at or after this time (ISO 8601), oldest change first. Poll with the updatedAt of the last row you saw. (optional)
    kind := "kind_example" // string | Physical goods or Web3 assets (NFT listings). Omit for both. (optional)
    tag := TODO // AnyOfstringarray | Only listings carrying ANY of these tags (case-insensitive): repeat `tag` or comma-separate. (optional)
    stillLive := true // bool | Only listings still active on at least one marketplace. With status=sold: sold items still buyable somewhere (a delist failed, is pending or never ran). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsApi.ListListings(context.Background()).Page(page).Limit(limit).Platform(platform).Status(status).Sku(sku).UpdatedSince(updatedSince).Kind(kind).Tag(tag).StillLive(stillLive).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListListings``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListListings`: ListListingsResponse
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
 **sku** | **string** | Exact SKU match: look up the row a supplier product maps to. | 
 **updatedSince** | **time.Time** | Only rows created or changed at or after this time (ISO 8601), oldest change first. Poll with the updatedAt of the last row you saw. | 
 **kind** | **string** | Physical goods or Web3 assets (NFT listings). Omit for both. | 
 **tag** | **AnyOfstringarray** | Only listings carrying ANY of these tags (case-insensitive): repeat &#x60;tag&#x60; or comma-separate. | 
 **stillLive** | **bool** | Only listings still active on at least one marketplace. With status&#x3D;sold: sold items still buyable somewhere (a delist failed, is pending or never ran). | 

### Return type

**ListListingsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListLiveStreams

> ListLiveStreamsResponse ListLiveStreams(ctx).Execute()

Your live shows, newest scheduled first: status, cover, category, connection method and peak viewers.

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
    resp, r, err := apiClient.ListingsApi.ListLiveStreams(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListLiveStreams``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLiveStreams`: ListLiveStreamsResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListLiveStreams`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLiveStreamsRequest struct via the builder pattern


### Return type

**ListLiveStreamsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListPassportEvents

> ListPassportEventsResponse ListPassportEvents(ctx, id).Execute()

A passport’s event history, newest first (never a buyer or a price).

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
    resp, r, err := apiClient.ListingsApi.ListPassportEvents(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListPassportEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPassportEvents`: ListPassportEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListPassportEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPassportEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListPassportEventsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListPassportPrintTargets

> ListPassportPrintTargetsResponse ListPassportPrintTargets(ctx).Execute()

Your printer agents (online, printers, default) and the passport label templates each can print.

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
    resp, r, err := apiClient.ListingsApi.ListPassportPrintTargets(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListPassportPrintTargets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPassportPrintTargets`: ListPassportPrintTargetsResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListPassportPrintTargets`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPassportPrintTargetsRequest struct via the builder pattern


### Return type

**ListPassportPrintTargetsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListPassports

> ListPassportsResponse ListPassports(ctx).Execute()

Product passports you registered (optionally for one Crossly listing): the code and URL each tag carries.

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
    resp, r, err := apiClient.ListingsApi.ListPassports(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListPassports``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPassports`: ListPassportsResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListPassports`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPassportsRequest struct via the builder pattern


### Return type

**ListPassportsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListRentalOffers

> ListRentalOffersResponse ListRentalOffers(ctx).Execute()

What you offer for rent.

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
    resp, r, err := apiClient.ListingsApi.ListRentalOffers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListRentalOffers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRentalOffers`: ListRentalOffersResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListRentalOffers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRentalOffersRequest struct via the builder pattern


### Return type

**ListRentalOffersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListStories

> ListStoriesResponse ListStories(ctx).Execute()

Your active story slides (newest run first, each run in play order), with their overlays: text, listing stickers, @mentions, links, emoji and drawing.

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
    resp, r, err := apiClient.ListingsApi.ListStories(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.ListStories``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStories`: ListStoriesResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.ListStories`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStoriesRequest struct via the builder pattern


### Return type

**ListStoriesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateEvent

> EventResponse UpdateEvent(ctx, id).Execute()

Edit an event. Omitted fields are unchanged; null clears.

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
    resp, r, err := apiClient.ListingsApi.UpdateEvent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.UpdateEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEvent`: EventResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.UpdateEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**EventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateHotspot

> UpdateHotspotResponse UpdateHotspot(ctx).Execute()

Tag your listings at points on a story or lookbook photo (shop the look). Replaces that photo’s tags; at most 12.

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
    resp, r, err := apiClient.ListingsApi.UpdateHotspot(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.UpdateHotspot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateHotspot`: UpdateHotspotResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.UpdateHotspot`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateHotspotRequest struct via the builder pattern


### Return type

**UpdateHotspotResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**UpdateListingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateListingCustom

> UpdateListingCustomResponse UpdateListingCustom(ctx, id).Execute()

Set a listing’s customs: HS code (6–10 digits), country of origin, and countries this item must not ship to on Crossly. Omitted fields are left alone; null (or [] for restrictedCountries) clears.

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
    resp, r, err := apiClient.ListingsApi.UpdateListingCustom(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.UpdateListingCustom``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateListingCustom`: UpdateListingCustomResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.UpdateListingCustom`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListingCustomRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateListingCustomResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateListingPlatformListing

> UpdateListingPlatformListingResponse UpdateListingPlatformListing(ctx, id, plId).Execute()

Set one marketplace copy's price (e.g. the ebay.com copy of a listing also live on ebay.co.uk), in that copy's own currency.

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
    plId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsApi.UpdateListingPlatformListing(context.Background(), id, plId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.UpdateListingPlatformListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateListingPlatformListing`: UpdateListingPlatformListingResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.UpdateListingPlatformListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**plId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateListingPlatformListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**UpdateListingPlatformListingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateLiveStream

> LiveStreamResponse UpdateLiveStream(ctx, id).Execute()

Edit a show that has not ended: title, description, cover, category, start time (not once live) or connection method (browser | phone | encoder).

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
    resp, r, err := apiClient.ListingsApi.UpdateLiveStream(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.UpdateLiveStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLiveStream`: LiveStreamResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.UpdateLiveStream`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLiveStreamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**LiveStreamResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateLiveStreamLineup

> UpdateLiveStreamLineupResponse UpdateLiveStreamLineup(ctx, id).Execute()

Set a live show’s lineup: your Crossly listings, in the order you will sell them. Buyers see it before they join; a sale of a lineup item while the show is live is recorded at its replay moment.

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
    resp, r, err := apiClient.ListingsApi.UpdateLiveStreamLineup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.UpdateLiveStreamLineup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLiveStreamLineup`: UpdateLiveStreamLineupResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.UpdateLiveStreamLineup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLiveStreamLineupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateLiveStreamLineupResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdatePassport

> UpdatePassportResponse UpdatePassport(ctx, id).Execute()

Change what a passport says (title, condition, maker, year, materials, notes) or add photo URLs (photos append: a passport keeps the ones it was registered with).

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
    resp, r, err := apiClient.ListingsApi.UpdatePassport(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.UpdatePassport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePassport`: UpdatePassportResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.UpdatePassport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePassportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdatePassportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdatePassportDraft

> PassportDraftResponse UpdatePassportDraft(ctx, id).Execute()

Move a draft to a step, pick its label template, add items or remove draft passports.

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
    resp, r, err := apiClient.ListingsApi.UpdatePassportDraft(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.UpdatePassportDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePassportDraft`: PassportDraftResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.UpdatePassportDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePassportDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**PassportDraftResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdatePricingSetting

> GetPricingSettingResponse UpdatePricingSetting(ctx).Execute()

Turn own-price match and/or layaway on or off for every one of your stores. One store: PUT /v1/storefronts/{id}/settings.

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
    resp, r, err := apiClient.ListingsApi.UpdatePricingSetting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.UpdatePricingSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePricingSetting`: GetPricingSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.UpdatePricingSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePricingSettingRequest struct via the builder pattern


### Return type

**GetPricingSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateRentalOffer

> UpdateRentalOfferResponse UpdateRentalOffer(ctx, platformListingId).Execute()

Offer a Crossly listing for rent: day rate, optional week rate, deposit, min/max days, turnaround buffer, rent-only (off sale while offered), seller approval, unavailable date ranges (blackoutRanges, inclusive YYYY-MM-DD) and a return policy. Keeps the listing’s own rental terms in step. Needs a ship-from address.

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
    platformListingId := "platformListingId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ListingsApi.UpdateRentalOffer(context.Background(), platformListingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.UpdateRentalOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRentalOffer`: UpdateRentalOfferResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.UpdateRentalOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformListingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRentalOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateRentalOfferResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateRentalSetting

> GetRentalSettingResponse UpdateRentalSetting(ctx).Execute()

Set the store’s rental defaults and payout timing (on_start needs 3 rentals that came back).

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
    resp, r, err := apiClient.ListingsApi.UpdateRentalSetting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.UpdateRentalSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateRentalSetting`: GetRentalSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.UpdateRentalSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRentalSettingRequest struct via the builder pattern


### Return type

**GetRentalSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateVariationGroup

> CreateVariationGroupResponse UpdateVariationGroup(ctx, id).Execute()

Retitle a group, rename its axis, or change a group's purchase rule (written into every member's description).

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
    resp, r, err := apiClient.ListingsApi.UpdateVariationGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ListingsApi.UpdateVariationGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateVariationGroup`: CreateVariationGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `ListingsApi.UpdateVariationGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateVariationGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVariationGroupResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

