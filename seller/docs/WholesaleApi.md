# \WholesaleApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWholesaleBuyerInvite**](WholesaleApi.md#CreateWholesaleBuyerInvite) | **Post** /v1/wholesale/buyers/invite | Invite a buyer by email to one or more of your catalogs. Binds to their account now if they have one, or when they sign up.
[**CreateWholesaleBuyerRemove**](WholesaleApi.md#CreateWholesaleBuyerRemove) | **Post** /v1/wholesale/buyers/remove | Revoke every catalog grant a buyer (or invited email) has with you. Their terms and invoices are kept.
[**CreateWholesaleCatalog**](WholesaleApi.md#CreateWholesaleCatalog) | **Post** /v1/wholesale/catalogs | Create a catalog. visibility public | public_prices_private (prices for approved buyers) | private_discoverable (listed, opened on approval) | link (anyone with the URL) | invite (named buyers only; \&quot;private\&quot; is accepted as invite); autoApprove never | verified | all; the cascade terms (minOrderQuantity, casePack, leadTimeDays, verifiedOnly, priceVisibility, maxPerOrder, backorder, restockEta, preset ids) default every item; listed advertises a public catalog in the wholesale directory (never its prices).
[**CreateWholesaleCatalogItem**](WholesaleApi.md#CreateWholesaleCatalogItem) | **Post** /v1/wholesale/catalogs/{id}/items | Add your Crossly listings to a catalog (up to 500), each with an optional wholesalePriceCents, minOrderQuantity and casePack. Items already in the catalog have their terms updated.
[**CreateWholesaleDrop**](WholesaleApi.md#CreateWholesaleDrop) | **Post** /v1/wholesale/drops | Create a draft trade drop: name, dropAt (ISO; when it opens), audience trade|catalog (+ audienceCatalogId), kind release|prebook. A pre-book needs depositBps (100-9900) and shipBy (YYYY-MM-DD), and takes balanceNetDays (0-90) for the balance invoice. listingIds are your listing ids, each on Crossly. A pre-book makes its listings pre-orders, so buyers pay the deposit at checkout. Schedule it with PATCH status \&quot;scheduled\&quot;.
[**CreateWholesaleDropRelease**](WholesaleApi.md#CreateWholesaleDropRelease) | **Post** /v1/wholesale/drops/{id}/release | Release a live pre-book: every reserved balance becomes an invoice on account to its buyer (due in the drop balanceNetDays), emailed, and the drop ends. Safe to retry: only balances not yet invoiced are invoiced. Returns { invoiced, balanceCents, invoiceIds }.
[**CreateWholesaleFreightRequestAnswer**](WholesaleApi.md#CreateWholesaleFreightRequestAnswer) | **Post** /v1/wholesale/freight-requests/{id}/answer | Answer an open freight request with a figure in CENTS (and an optional note).
[**CreateWholesaleGroup**](WholesaleApi.md#CreateWholesaleGroup) | **Post** /v1/wholesale/groups | Create a customer group with a discount in basis points (up to 9000). A buyer in several groups gets the best rate, never a stacked one.
[**CreateWholesaleGroupMember**](WholesaleApi.md#CreateWholesaleGroupMember) | **Post** /v1/wholesale/groups/{id}/members | Add a buyer to a group by account id or by the email of an existing Crossly account.
[**CreateWholesaleInvoiceCreditNote**](WholesaleApi.md#CreateWholesaleInvoiceCreditNote) | **Post** /v1/wholesale/invoices/{id}/credit-notes | Issue a credit note on an on-account invoice (short shipment, damaged, pricing error, returned, goodwill, other): at most the balance due. Numbered &lt;invoice&gt;-CN&lt;n&gt;; settling the rest closes the invoice paid. Not a refund.
[**CreateWholesaleInvoiceLateFee**](WholesaleApi.md#CreateWholesaleInvoiceLateFee) | **Post** /v1/wholesale/invoices/{id}/late-fee | Add your late fee (invoice settings lateFeeBps) to an invoice by hand, once. The hourly sweep adds it automatically when an invoice goes overdue.
[**CreateWholesaleInvoiceMarkPaid**](WholesaleApi.md#CreateWholesaleInvoiceMarkPaid) | **Post** /v1/wholesale/invoices/{id}/mark-paid | Record a payment the buyer made to you directly: method ach | check | wire | card | other, a reference, and amountCents (omit for the whole balance). When the invoice is settled, Crossly&#39;s fee and the tax it remits are charged to your billing card.
[**CreateWholesaleInvoiceSend**](WholesaleApi.md#CreateWholesaleInvoiceSend) | **Post** /v1/wholesale/invoices/{id}/send | Email the invoice to the buyer again (an overdue one goes as an overdue notice), with the link to pay it.
[**CreateWholesaleInvoiceVoid**](WholesaleApi.md#CreateWholesaleInvoiceVoid) | **Post** /v1/wholesale/invoices/{id}/void | Void an invoice nobody has paid anything on (e.g. the order was cancelled). It stops counting against the buyer&#39;s credit limit.
[**CreateWholesaleInvoiceWriteOff**](WholesaleApi.md#CreateWholesaleInvoiceWriteOff) | **Post** /v1/wholesale/invoices/{id}/write-off | Write off what is left on an invoice (bad debt): it closes as written_off and stops counting against the buyer&#39;s credit.
[**CreateWholesaleProfileShipsFromValidate**](WholesaleApi.md#CreateWholesaleProfileShipsFromValidate) | **Post** /v1/wholesale/profile/ships-from/validate | Check a ships-from origin ({ postalCode, country, city?, region? }) with the shipping validator without saving it. 400 invalid_ships_from with the issues; 422 for an embargoed origin.
[**CreateWholesaleQuoteAnswer**](WholesaleApi.md#CreateWholesaleQuoteAnswer) | **Post** /v1/wholesale/quotes/{id}/answer | Answer (or revise) a quote: a unit price in CENTS per item you quote, an expiresAt (ISO, at most 90 days out) and an optional note. Items left out are not quoted. The buyer is notified and can accept until expiry.
[**CreateWholesaleQuoteDecline**](WholesaleApi.md#CreateWholesaleQuoteDecline) | **Post** /v1/wholesale/quotes/{id}/decline | Decline a quote request, with an optional reason the buyer sees.
[**CreateWholesaleRequestDecide**](WholesaleApi.md#CreateWholesaleRequestDecide) | **Post** /v1/wholesale/requests/{id}/decide | Approve (grants access to that catalog; optional groupId adds the buyer to a customer group) or decline (optional reason) an access request.
[**CreateWholesaleTierBulk**](WholesaleApi.md#CreateWholesaleTierBulk) | **Post** /v1/wholesale/tiers/bulk | One ladder for many listings: breaks as [{ minQuantity, discountBps }] off each listing&#39;s own price. Returns how many were updated and which were skipped and why.
[**DeleteWholesaleBuyerTerm**](WholesaleApi.md#DeleteWholesaleBuyerTerm) | **Delete** /v1/wholesale/buyers/{buyerUserId}/terms | Switch a buyer&#39;s account terms off. Invoices already issued are unaffected.
[**DeleteWholesaleCatalog**](WholesaleApi.md#DeleteWholesaleCatalog) | **Delete** /v1/wholesale/catalogs/{id} | Delete a catalog with its items, grants and access requests. Orders and invoices are untouched.
[**DeleteWholesaleCatalogItem**](WholesaleApi.md#DeleteWholesaleCatalogItem) | **Delete** /v1/wholesale/catalogs/{id}/items/{platformListingId} | Take an item out of a catalog. The listing itself is untouched.
[**DeleteWholesaleGroup**](WholesaleApi.md#DeleteWholesaleGroup) | **Delete** /v1/wholesale/groups/{id} | Delete a customer group. Its members keep any other group or account rate.
[**DeleteWholesaleGroupMember**](WholesaleApi.md#DeleteWholesaleGroupMember) | **Delete** /v1/wholesale/groups/{id}/members/{buyerUserId} | Take a buyer out of a group.
[**GetWholesaleCatalogLineSheet**](WholesaleApi.md#GetWholesaleCatalogLineSheet) | **Get** /v1/wholesale/catalogs/{id}/line-sheet | The catalog&#39;s line sheet: your company details and terms, then every item with retail and trade price, minimum, case pack and quantity breaks.
[**GetWholesaleCatalogShare**](WholesaleApi.md#GetWholesaleCatalogShare) | **Get** /v1/wholesale/catalogs/{id}/share | The catalog&#39;s buyer link (null until you have a shop username), to share or print as a QR code.
[**GetWholesaleCategory**](WholesaleApi.md#GetWholesaleCategory) | **Get** /v1/wholesale/categories | The wholesale category taxonomy (id, label, group). A catalog or trade profile takes up to 12 ids, or \&quot;other:&lt;words&gt;\&quot;.
[**GetWholesaleDrop**](WholesaleApi.md#GetWholesaleDrop) | **Get** /v1/wholesale/drops/{id} | One trade drop with its listings and, for a pre-book, every reservation: buyer, item, quantity, deposit, balance, status and the balance invoice once released.
[**GetWholesaleInvoice**](WholesaleApi.md#GetWholesaleInvoice) | **Get** /v1/wholesale/invoices/{id} | One invoice with its lines, PO reference, payments and whether Crossly dues on it are collected.
[**GetWholesaleOverview**](WholesaleApi.md#GetWholesaleOverview) | **Get** /v1/wholesale/overview | Your wholesale store at a glance: the setup checklist, 30-day wholesale revenue, open orders, outstanding invoices, active buyers and pending access requests.
[**GetWholesaleProfile**](WholesaleApi.md#GetWholesaleProfile) | **Get** /v1/wholesale/profile | A store&#39;s trade profile (X-Storefront-Id; default store when omitted): company, tagline, logo and banner, contact, ships-from origin, lead time, order minimums (opening, first order, reorder), the default MOQ / case pack / max per order / backorder / price visibility, return, payment and shipping preset ids, invoicing and its settings, buyer requirements and the directory defaults.
[**GetWholesaleQuote**](WholesaleApi.md#GetWholesaleQuote) | **Get** /v1/wholesale/quotes/{id} | One quote request.
[**GetWholesaleReserve**](WholesaleApi.md#GetWholesaleReserve) | **Get** /v1/wholesale/reserve/{inventoryItemId} | How many units of an inventory item are held back for wholesale, and what retail channels advertise.
[**GetWholesaleTermPreview**](WholesaleApi.md#GetWholesaleTermPreview) | **Get** /v1/wholesale/terms/preview | The trade terms that apply to a catalog item (?catalogId&amp;platformListingId), a catalog (?catalogId) or the trade profile alone, each with its source (item, catalog, profile, default), plus the return, payment and shipping policies they resolve to.
[**ListWholesaleBuyers**](WholesaleApi.md#ListWholesaleBuyers) | **Get** /v1/wholesale/buyers | Your trade buyers: catalogs they can open, customer groups, account terms, outstanding balance, wholesale spend, business verification and resale-certificate status (read-only — Crossly reviews those).
[**ListWholesaleCatalogItems**](WholesaleApi.md#ListWholesaleCatalogItems) | **Get** /v1/wholesale/catalogs/{id}/items | A catalog&#39;s items in merchandising order, each with its retail price, trade price, minimum order quantity, case pack and quantity breaks.
[**ListWholesaleCatalogs**](WholesaleApi.md#ListWholesaleCatalogs) | **Get** /v1/wholesale/catalogs | Your wholesale catalogs with item, buyer-grant and pending-request counts.
[**ListWholesaleDrops**](WholesaleApi.md#ListWholesaleDrops) | **Get** /v1/wholesale/drops | Your trade drops and pre-books: audience (trade &#x3D; your approved trade buyers, catalog &#x3D; whoever can open audienceCatalogId), kind (release | prebook), deposit, ship date, status, and for a pre-book its reservation totals (buyers, units, deposits taken, balances still to invoice).
[**ListWholesaleFreightRequests**](WholesaleApi.md#ListWholesaleFreightRequests) | **Get** /v1/wholesale/freight-requests | Freight (LTL) quote requests from buyers whose order is over your pallet threshold.
[**ListWholesaleGroups**](WholesaleApi.md#ListWholesaleGroups) | **Get** /v1/wholesale/groups | Your customer groups with their discount and members.
[**ListWholesaleInvoices**](WholesaleApi.md#ListWholesaleInvoices) | **Get** /v1/wholesale/invoices | On-account invoices, newest first, with balance due. ?status&#x3D;unpaid for open, partial and overdue together.
[**ListWholesaleOrders**](WholesaleApi.md#ListWholesaleOrders) | **Get** /v1/wholesale/orders | Wholesale orders: Crossly orders priced through your trade terms (catalog price, quantity break or account rate) or placed on account. Fulfil them like any order.
[**ListWholesaleQuotes**](WholesaleApi.md#ListWholesaleQuotes) | **Get** /v1/wholesale/quotes | Quote requests (RFQ) from trade buyers: items with quantity and target price, your quoted prices and expiry. ?status&#x3D;open lists the ones waiting on you or the buyer.
[**ListWholesaleRequests**](WholesaleApi.md#ListWholesaleRequests) | **Get** /v1/wholesale/requests | Requests from buyers to open your private or trade-only catalogs, with each buyer&#39;s business-verification and tax-exempt status. ?status&#x3D;all includes decided ones.
[**ListWholesaleSetupSteps**](WholesaleApi.md#ListWholesaleSetupSteps) | **Get** /v1/wholesale/setup-steps | The wholesale setup checklist for a store: profile, ships-from, payment, returns, a catalog with items, visibility chosen, live. Each step is done because the thing exists.
[**UpdateWholesaleBuyerTerm**](WholesaleApi.md#UpdateWholesaleBuyerTerm) | **Put** /v1/wholesale/buyers/{buyerUserId}/terms | Set a buyer&#39;s account terms: netDays (0-120) for paying on account, creditLimitCents (the most they may owe you at once; null &#x3D; no limit) and discountPercentBps (their negotiated rate).
[**UpdateWholesaleCatalog**](WholesaleApi.md#UpdateWholesaleCatalog) | **Patch** /v1/wholesale/catalogs/{id} | Edit a catalog: name, slug, description, cover image, visibility, opening-order minimum, trade gate, directory listing, active.
[**UpdateWholesaleCatalogItem**](WholesaleApi.md#UpdateWholesaleCatalogItem) | **Patch** /v1/wholesale/catalogs/{id}/items | Change many catalog items at once: items[] sets each row (trade price, MOQ, case pack, lead time, max per order, backorder, restock date, verified-only, price visibility, preset ids); or platformListingIds + set applies the same values to each, set.pctOffRetailBps pricing each from its retail price.
[**UpdateWholesaleCatalogItemByPlatformListingId**](WholesaleApi.md#UpdateWholesaleCatalogItemByPlatformListingId) | **Patch** /v1/wholesale/catalogs/{id}/items/{platformListingId} | Change one item&#39;s trade price and its own trade terms in this catalog (MOQ, case pack, lead time, max per order, backorder, restock date, verified-only, price visibility, preset ids). null clears &#x3D; inherit from the catalog.
[**UpdateWholesaleCatalogItemOrder**](WholesaleApi.md#UpdateWholesaleCatalogItemOrder) | **Put** /v1/wholesale/catalogs/{id}/items/order | Set the merchandising order: platform listing ids, first to last. Items left out keep their relative order after these.
[**UpdateWholesaleDrop**](WholesaleApi.md#UpdateWholesaleDrop) | **Patch** /v1/wholesale/drops/{id} | Edit a trade drop (any create field but kind), schedule it (status \&quot;scheduled\&quot;, with dropAt in the future) or cancel it. A live drop can only be cancelled or have endsAt / perCustomerLimit changed.
[**UpdateWholesaleGroup**](WholesaleApi.md#UpdateWholesaleGroup) | **Patch** /v1/wholesale/groups/{id} | Rename a group, change its discount, or pause it (active&#x3D;false).
[**UpdateWholesaleListingTier**](WholesaleApi.md#UpdateWholesaleListingTier) | **Put** /v1/wholesale/listings/{platformListingId}/tiers | Replace one listing&#39;s quantity breaks: [{ minQuantity ≥ 2, unitPriceCents }], each cheaper than the last and than the list price. An empty list removes them. Applies to every buyer at checkout.
[**UpdateWholesaleProfile**](WholesaleApi.md#UpdateWholesaleProfile) | **Put** /v1/wholesale/profile | Save a store&#39;s trade profile. Omitted fields are unchanged; null clears. shipsFrom is checked by the shipping validator (400 invalid_ships_from). The profile is the bottom of the trade-terms cascade: catalogs and items that leave a term empty inherit it — saving never rewrites catalogs. Preset ids must be your own presets of that kind.
[**UpdateWholesaleReserve**](WholesaleApi.md#UpdateWholesaleReserve) | **Put** /v1/wholesale/reserve/{inventoryItemId} | Reserve N units of an inventory item for wholesale: retail channels advertise the rest and come down when only the reserve is left. 0 clears it.



## CreateWholesaleBuyerInvite

> CreateWholesaleBuyerInviteResponse CreateWholesaleBuyerInvite(ctx).Execute()

Invite a buyer by email to one or more of your catalogs. Binds to their account now if they have one, or when they sign up.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleBuyerInvite(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleBuyerInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleBuyerInvite`: CreateWholesaleBuyerInviteResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleBuyerInvite`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleBuyerInviteRequest struct via the builder pattern


### Return type

**CreateWholesaleBuyerInviteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleBuyerRemove

> CreateWholesaleBuyerRemoveResponse CreateWholesaleBuyerRemove(ctx).Execute()

Revoke every catalog grant a buyer (or invited email) has with you. Their terms and invoices are kept.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleBuyerRemove(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleBuyerRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleBuyerRemove`: CreateWholesaleBuyerRemoveResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleBuyerRemove`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleBuyerRemoveRequest struct via the builder pattern


### Return type

**CreateWholesaleBuyerRemoveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleCatalog

> CreateWholesaleCatalogResponse CreateWholesaleCatalog(ctx).Execute()

Create a catalog. visibility public | public_prices_private (prices for approved buyers) | private_discoverable (listed, opened on approval) | link (anyone with the URL) | invite (named buyers only; \"private\" is accepted as invite); autoApprove never | verified | all; the cascade terms (minOrderQuantity, casePack, leadTimeDays, verifiedOnly, priceVisibility, maxPerOrder, backorder, restockEta, preset ids) default every item; listed advertises a public catalog in the wholesale directory (never its prices).

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleCatalog(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleCatalog`: CreateWholesaleCatalogResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleCatalogRequest struct via the builder pattern


### Return type

**CreateWholesaleCatalogResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleCatalogItem

> CreateWholesaleCatalogItemResponse CreateWholesaleCatalogItem(ctx, id).Execute()

Add your Crossly listings to a catalog (up to 500), each with an optional wholesalePriceCents, minOrderQuantity and casePack. Items already in the catalog have their terms updated.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleCatalogItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleCatalogItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleCatalogItem`: CreateWholesaleCatalogItemResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleCatalogItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleCatalogItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWholesaleCatalogItemResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleDrop

> CreateWholesaleDropResponse CreateWholesaleDrop(ctx).Execute()

Create a draft trade drop: name, dropAt (ISO; when it opens), audience trade|catalog (+ audienceCatalogId), kind release|prebook. A pre-book needs depositBps (100-9900) and shipBy (YYYY-MM-DD), and takes balanceNetDays (0-90) for the balance invoice. listingIds are your listing ids, each on Crossly. A pre-book makes its listings pre-orders, so buyers pay the deposit at checkout. Schedule it with PATCH status \"scheduled\".

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleDrop(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleDrop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleDrop`: CreateWholesaleDropResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleDrop`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleDropRequest struct via the builder pattern


### Return type

**CreateWholesaleDropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleDropRelease

> CreateWholesaleDropReleaseResponse CreateWholesaleDropRelease(ctx, id).Execute()

Release a live pre-book: every reserved balance becomes an invoice on account to its buyer (due in the drop balanceNetDays), emailed, and the drop ends. Safe to retry: only balances not yet invoiced are invoiced. Returns { invoiced, balanceCents, invoiceIds }.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleDropRelease(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleDropRelease``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleDropRelease`: CreateWholesaleDropReleaseResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleDropRelease`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleDropReleaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWholesaleDropReleaseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleFreightRequestAnswer

> CreateWholesaleFreightRequestAnswerResponse CreateWholesaleFreightRequestAnswer(ctx, id).Execute()

Answer an open freight request with a figure in CENTS (and an optional note).

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleFreightRequestAnswer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleFreightRequestAnswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleFreightRequestAnswer`: CreateWholesaleFreightRequestAnswerResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleFreightRequestAnswer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleFreightRequestAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWholesaleFreightRequestAnswerResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleGroup

> CreateWholesaleGroupResponse CreateWholesaleGroup(ctx).Execute()

Create a customer group with a discount in basis points (up to 9000). A buyer in several groups gets the best rate, never a stacked one.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleGroup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleGroup`: CreateWholesaleGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleGroup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleGroupRequest struct via the builder pattern


### Return type

**CreateWholesaleGroupResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleGroupMember

> CreateWholesaleGroupMemberResponse CreateWholesaleGroupMember(ctx, id).Execute()

Add a buyer to a group by account id or by the email of an existing Crossly account.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleGroupMember(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleGroupMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleGroupMember`: CreateWholesaleGroupMemberResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleGroupMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWholesaleGroupMemberResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleInvoiceCreditNote

> CreateWholesaleInvoiceCreditNoteResponse CreateWholesaleInvoiceCreditNote(ctx, id).Execute()

Issue a credit note on an on-account invoice (short shipment, damaged, pricing error, returned, goodwill, other): at most the balance due. Numbered <invoice>-CN<n>; settling the rest closes the invoice paid. Not a refund.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleInvoiceCreditNote(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleInvoiceCreditNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleInvoiceCreditNote`: CreateWholesaleInvoiceCreditNoteResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleInvoiceCreditNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleInvoiceCreditNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWholesaleInvoiceCreditNoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleInvoiceLateFee

> WholesaleInvoiceResponse CreateWholesaleInvoiceLateFee(ctx, id).Execute()

Add your late fee (invoice settings lateFeeBps) to an invoice by hand, once. The hourly sweep adds it automatically when an invoice goes overdue.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleInvoiceLateFee(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleInvoiceLateFee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleInvoiceLateFee`: WholesaleInvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleInvoiceLateFee`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleInvoiceLateFeeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WholesaleInvoiceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleInvoiceMarkPaid

> WholesaleInvoiceResponse CreateWholesaleInvoiceMarkPaid(ctx, id).Execute()

Record a payment the buyer made to you directly: method ach | check | wire | card | other, a reference, and amountCents (omit for the whole balance). When the invoice is settled, Crossly's fee and the tax it remits are charged to your billing card.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleInvoiceMarkPaid(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleInvoiceMarkPaid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleInvoiceMarkPaid`: WholesaleInvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleInvoiceMarkPaid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleInvoiceMarkPaidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WholesaleInvoiceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleInvoiceSend

> CreateWholesaleInvoiceSendResponse CreateWholesaleInvoiceSend(ctx, id).Execute()

Email the invoice to the buyer again (an overdue one goes as an overdue notice), with the link to pay it.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleInvoiceSend(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleInvoiceSend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleInvoiceSend`: CreateWholesaleInvoiceSendResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleInvoiceSend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleInvoiceSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWholesaleInvoiceSendResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleInvoiceVoid

> WholesaleInvoiceResponse CreateWholesaleInvoiceVoid(ctx, id).Execute()

Void an invoice nobody has paid anything on (e.g. the order was cancelled). It stops counting against the buyer's credit limit.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleInvoiceVoid(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleInvoiceVoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleInvoiceVoid`: WholesaleInvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleInvoiceVoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleInvoiceVoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WholesaleInvoiceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleInvoiceWriteOff

> WholesaleInvoiceResponse CreateWholesaleInvoiceWriteOff(ctx, id).Execute()

Write off what is left on an invoice (bad debt): it closes as written_off and stops counting against the buyer's credit.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleInvoiceWriteOff(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleInvoiceWriteOff``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleInvoiceWriteOff`: WholesaleInvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleInvoiceWriteOff`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleInvoiceWriteOffRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WholesaleInvoiceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleProfileShipsFromValidate

> CreateWholesaleProfileShipsFromValidateResponse CreateWholesaleProfileShipsFromValidate(ctx).Execute()

Check a ships-from origin ({ postalCode, country, city?, region? }) with the shipping validator without saving it. 400 invalid_ships_from with the issues; 422 for an embargoed origin.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleProfileShipsFromValidate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleProfileShipsFromValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleProfileShipsFromValidate`: CreateWholesaleProfileShipsFromValidateResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleProfileShipsFromValidate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleProfileShipsFromValidateRequest struct via the builder pattern


### Return type

**CreateWholesaleProfileShipsFromValidateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleQuoteAnswer

> WholesaleQuoteResponse CreateWholesaleQuoteAnswer(ctx, id).Execute()

Answer (or revise) a quote: a unit price in CENTS per item you quote, an expiresAt (ISO, at most 90 days out) and an optional note. Items left out are not quoted. The buyer is notified and can accept until expiry.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleQuoteAnswer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleQuoteAnswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleQuoteAnswer`: WholesaleQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleQuoteAnswer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleQuoteAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WholesaleQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleQuoteDecline

> WholesaleQuoteResponse CreateWholesaleQuoteDecline(ctx, id).Execute()

Decline a quote request, with an optional reason the buyer sees.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleQuoteDecline(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleQuoteDecline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleQuoteDecline`: WholesaleQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleQuoteDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleQuoteDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WholesaleQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleRequestDecide

> CreateWholesaleRequestDecideResponse CreateWholesaleRequestDecide(ctx, id).Execute()

Approve (grants access to that catalog; optional groupId adds the buyer to a customer group) or decline (optional reason) an access request.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleRequestDecide(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleRequestDecide``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleRequestDecide`: CreateWholesaleRequestDecideResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleRequestDecide`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleRequestDecideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWholesaleRequestDecideResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWholesaleTierBulk

> CreateWholesaleTierBulkResponse CreateWholesaleTierBulk(ctx).Execute()

One ladder for many listings: breaks as [{ minQuantity, discountBps }] off each listing's own price. Returns how many were updated and which were skipped and why.

### Example

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
    resp, r, err := apiClient.WholesaleApi.CreateWholesaleTierBulk(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.CreateWholesaleTierBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWholesaleTierBulk`: CreateWholesaleTierBulkResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.CreateWholesaleTierBulk`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWholesaleTierBulkRequest struct via the builder pattern


### Return type

**CreateWholesaleTierBulkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWholesaleBuyerTerm

> UpdateGroupBuyResponse DeleteWholesaleBuyerTerm(ctx, buyerUserId).Execute()

Switch a buyer's account terms off. Invoices already issued are unaffected.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    buyerUserId := "buyerUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WholesaleApi.DeleteWholesaleBuyerTerm(context.Background(), buyerUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.DeleteWholesaleBuyerTerm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWholesaleBuyerTerm`: UpdateGroupBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.DeleteWholesaleBuyerTerm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyerUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWholesaleBuyerTermRequest struct via the builder pattern


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


## DeleteWholesaleCatalog

> UpdateGroupBuyResponse DeleteWholesaleCatalog(ctx, id).Execute()

Delete a catalog with its items, grants and access requests. Orders and invoices are untouched.

### Example

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
    resp, r, err := apiClient.WholesaleApi.DeleteWholesaleCatalog(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.DeleteWholesaleCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWholesaleCatalog`: UpdateGroupBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.DeleteWholesaleCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWholesaleCatalogRequest struct via the builder pattern


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


## DeleteWholesaleCatalogItem

> UpdateGroupBuyResponse DeleteWholesaleCatalogItem(ctx, id, platformListingId).Execute()

Take an item out of a catalog. The listing itself is untouched.

### Example

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
    platformListingId := "platformListingId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WholesaleApi.DeleteWholesaleCatalogItem(context.Background(), id, platformListingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.DeleteWholesaleCatalogItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWholesaleCatalogItem`: UpdateGroupBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.DeleteWholesaleCatalogItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**platformListingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWholesaleCatalogItemRequest struct via the builder pattern


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


## DeleteWholesaleGroup

> UpdateGroupBuyResponse DeleteWholesaleGroup(ctx, id).Execute()

Delete a customer group. Its members keep any other group or account rate.

### Example

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
    resp, r, err := apiClient.WholesaleApi.DeleteWholesaleGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.DeleteWholesaleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWholesaleGroup`: UpdateGroupBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.DeleteWholesaleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWholesaleGroupRequest struct via the builder pattern


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


## DeleteWholesaleGroupMember

> UpdateGroupBuyResponse DeleteWholesaleGroupMember(ctx, id, buyerUserId).Execute()

Take a buyer out of a group.

### Example

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
    buyerUserId := "buyerUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WholesaleApi.DeleteWholesaleGroupMember(context.Background(), id, buyerUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.DeleteWholesaleGroupMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWholesaleGroupMember`: UpdateGroupBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.DeleteWholesaleGroupMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**buyerUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWholesaleGroupMemberRequest struct via the builder pattern


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


## GetWholesaleCatalogLineSheet

> GetWholesaleCatalogLineSheetResponse GetWholesaleCatalogLineSheet(ctx, id).Execute()

The catalog's line sheet: your company details and terms, then every item with retail and trade price, minimum, case pack and quantity breaks.

### Example

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
    resp, r, err := apiClient.WholesaleApi.GetWholesaleCatalogLineSheet(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.GetWholesaleCatalogLineSheet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWholesaleCatalogLineSheet`: GetWholesaleCatalogLineSheetResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.GetWholesaleCatalogLineSheet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWholesaleCatalogLineSheetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWholesaleCatalogLineSheetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWholesaleCatalogShare

> GetWholesaleCatalogShareResponse GetWholesaleCatalogShare(ctx, id).Execute()

The catalog's buyer link (null until you have a shop username), to share or print as a QR code.

### Example

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
    resp, r, err := apiClient.WholesaleApi.GetWholesaleCatalogShare(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.GetWholesaleCatalogShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWholesaleCatalogShare`: GetWholesaleCatalogShareResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.GetWholesaleCatalogShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWholesaleCatalogShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWholesaleCatalogShareResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWholesaleCategory

> GetWholesaleCategoryResponse GetWholesaleCategory(ctx).Execute()

The wholesale category taxonomy (id, label, group). A catalog or trade profile takes up to 12 ids, or \"other:<words>\".

### Example

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
    resp, r, err := apiClient.WholesaleApi.GetWholesaleCategory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.GetWholesaleCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWholesaleCategory`: GetWholesaleCategoryResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.GetWholesaleCategory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWholesaleCategoryRequest struct via the builder pattern


### Return type

**GetWholesaleCategoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWholesaleDrop

> GetWholesaleDropResponse GetWholesaleDrop(ctx, id).Execute()

One trade drop with its listings and, for a pre-book, every reservation: buyer, item, quantity, deposit, balance, status and the balance invoice once released.

### Example

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
    resp, r, err := apiClient.WholesaleApi.GetWholesaleDrop(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.GetWholesaleDrop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWholesaleDrop`: GetWholesaleDropResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.GetWholesaleDrop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWholesaleDropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWholesaleDropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWholesaleInvoice

> WholesaleInvoiceResponse GetWholesaleInvoice(ctx, id).Execute()

One invoice with its lines, PO reference, payments and whether Crossly dues on it are collected.

### Example

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
    resp, r, err := apiClient.WholesaleApi.GetWholesaleInvoice(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.GetWholesaleInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWholesaleInvoice`: WholesaleInvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.GetWholesaleInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWholesaleInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WholesaleInvoiceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWholesaleOverview

> GetWholesaleOverviewResponse GetWholesaleOverview(ctx).Execute()

Your wholesale store at a glance: the setup checklist, 30-day wholesale revenue, open orders, outstanding invoices, active buyers and pending access requests.

### Example

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
    resp, r, err := apiClient.WholesaleApi.GetWholesaleOverview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.GetWholesaleOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWholesaleOverview`: GetWholesaleOverviewResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.GetWholesaleOverview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWholesaleOverviewRequest struct via the builder pattern


### Return type

**GetWholesaleOverviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWholesaleProfile

> GetWholesaleProfileResponse GetWholesaleProfile(ctx).Execute()

A store's trade profile (X-Storefront-Id; default store when omitted): company, tagline, logo and banner, contact, ships-from origin, lead time, order minimums (opening, first order, reorder), the default MOQ / case pack / max per order / backorder / price visibility, return, payment and shipping preset ids, invoicing and its settings, buyer requirements and the directory defaults.

### Example

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
    resp, r, err := apiClient.WholesaleApi.GetWholesaleProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.GetWholesaleProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWholesaleProfile`: GetWholesaleProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.GetWholesaleProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWholesaleProfileRequest struct via the builder pattern


### Return type

**GetWholesaleProfileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWholesaleQuote

> WholesaleQuoteResponse GetWholesaleQuote(ctx, id).Execute()

One quote request.

### Example

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
    resp, r, err := apiClient.WholesaleApi.GetWholesaleQuote(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.GetWholesaleQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWholesaleQuote`: WholesaleQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.GetWholesaleQuote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWholesaleQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WholesaleQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWholesaleReserve

> GetWholesaleReserveResponse GetWholesaleReserve(ctx, inventoryItemId).Execute()

How many units of an inventory item are held back for wholesale, and what retail channels advertise.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inventoryItemId := "inventoryItemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WholesaleApi.GetWholesaleReserve(context.Background(), inventoryItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.GetWholesaleReserve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWholesaleReserve`: GetWholesaleReserveResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.GetWholesaleReserve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWholesaleReserveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWholesaleReserveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWholesaleTermPreview

> GetWholesaleTermPreviewResponse GetWholesaleTermPreview(ctx).Execute()

The trade terms that apply to a catalog item (?catalogId&platformListingId), a catalog (?catalogId) or the trade profile alone, each with its source (item, catalog, profile, default), plus the return, payment and shipping policies they resolve to.

### Example

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
    resp, r, err := apiClient.WholesaleApi.GetWholesaleTermPreview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.GetWholesaleTermPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWholesaleTermPreview`: GetWholesaleTermPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.GetWholesaleTermPreview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWholesaleTermPreviewRequest struct via the builder pattern


### Return type

**GetWholesaleTermPreviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWholesaleBuyers

> ListWholesaleBuyersResponse ListWholesaleBuyers(ctx).Execute()

Your trade buyers: catalogs they can open, customer groups, account terms, outstanding balance, wholesale spend, business verification and resale-certificate status (read-only — Crossly reviews those).

### Example

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
    resp, r, err := apiClient.WholesaleApi.ListWholesaleBuyers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.ListWholesaleBuyers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWholesaleBuyers`: ListWholesaleBuyersResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.ListWholesaleBuyers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWholesaleBuyersRequest struct via the builder pattern


### Return type

**ListWholesaleBuyersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWholesaleCatalogItems

> ListWholesaleCatalogItemsResponse ListWholesaleCatalogItems(ctx, id).Execute()

A catalog's items in merchandising order, each with its retail price, trade price, minimum order quantity, case pack and quantity breaks.

### Example

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
    resp, r, err := apiClient.WholesaleApi.ListWholesaleCatalogItems(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.ListWholesaleCatalogItems``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWholesaleCatalogItems`: ListWholesaleCatalogItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.ListWholesaleCatalogItems`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWholesaleCatalogItemsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWholesaleCatalogItemsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWholesaleCatalogs

> ListWholesaleCatalogsResponse ListWholesaleCatalogs(ctx).Execute()

Your wholesale catalogs with item, buyer-grant and pending-request counts.

### Example

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
    resp, r, err := apiClient.WholesaleApi.ListWholesaleCatalogs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.ListWholesaleCatalogs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWholesaleCatalogs`: ListWholesaleCatalogsResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.ListWholesaleCatalogs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWholesaleCatalogsRequest struct via the builder pattern


### Return type

**ListWholesaleCatalogsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWholesaleDrops

> ListWholesaleDropsResponse ListWholesaleDrops(ctx).Execute()

Your trade drops and pre-books: audience (trade = your approved trade buyers, catalog = whoever can open audienceCatalogId), kind (release | prebook), deposit, ship date, status, and for a pre-book its reservation totals (buyers, units, deposits taken, balances still to invoice).

### Example

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
    resp, r, err := apiClient.WholesaleApi.ListWholesaleDrops(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.ListWholesaleDrops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWholesaleDrops`: ListWholesaleDropsResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.ListWholesaleDrops`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWholesaleDropsRequest struct via the builder pattern


### Return type

**ListWholesaleDropsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWholesaleFreightRequests

> ListWholesaleFreightRequestsResponse ListWholesaleFreightRequests(ctx).Execute()

Freight (LTL) quote requests from buyers whose order is over your pallet threshold.

### Example

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
    resp, r, err := apiClient.WholesaleApi.ListWholesaleFreightRequests(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.ListWholesaleFreightRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWholesaleFreightRequests`: ListWholesaleFreightRequestsResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.ListWholesaleFreightRequests`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWholesaleFreightRequestsRequest struct via the builder pattern


### Return type

**ListWholesaleFreightRequestsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWholesaleGroups

> ListWholesaleGroupsResponse ListWholesaleGroups(ctx).Execute()

Your customer groups with their discount and members.

### Example

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
    resp, r, err := apiClient.WholesaleApi.ListWholesaleGroups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.ListWholesaleGroups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWholesaleGroups`: ListWholesaleGroupsResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.ListWholesaleGroups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWholesaleGroupsRequest struct via the builder pattern


### Return type

**ListWholesaleGroupsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWholesaleInvoices

> ListWholesaleInvoicesResponse ListWholesaleInvoices(ctx).Execute()

On-account invoices, newest first, with balance due. ?status=unpaid for open, partial and overdue together.

### Example

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
    resp, r, err := apiClient.WholesaleApi.ListWholesaleInvoices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.ListWholesaleInvoices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWholesaleInvoices`: ListWholesaleInvoicesResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.ListWholesaleInvoices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWholesaleInvoicesRequest struct via the builder pattern


### Return type

**ListWholesaleInvoicesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWholesaleOrders

> ListWholesaleOrdersResponse ListWholesaleOrders(ctx).Execute()

Wholesale orders: Crossly orders priced through your trade terms (catalog price, quantity break or account rate) or placed on account. Fulfil them like any order.

### Example

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
    resp, r, err := apiClient.WholesaleApi.ListWholesaleOrders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.ListWholesaleOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWholesaleOrders`: ListWholesaleOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.ListWholesaleOrders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWholesaleOrdersRequest struct via the builder pattern


### Return type

**ListWholesaleOrdersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWholesaleQuotes

> ListWholesaleQuotesResponse ListWholesaleQuotes(ctx).Execute()

Quote requests (RFQ) from trade buyers: items with quantity and target price, your quoted prices and expiry. ?status=open lists the ones waiting on you or the buyer.

### Example

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
    resp, r, err := apiClient.WholesaleApi.ListWholesaleQuotes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.ListWholesaleQuotes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWholesaleQuotes`: ListWholesaleQuotesResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.ListWholesaleQuotes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWholesaleQuotesRequest struct via the builder pattern


### Return type

**ListWholesaleQuotesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWholesaleRequests

> ListWholesaleRequestsResponse ListWholesaleRequests(ctx).Execute()

Requests from buyers to open your private or trade-only catalogs, with each buyer's business-verification and tax-exempt status. ?status=all includes decided ones.

### Example

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
    resp, r, err := apiClient.WholesaleApi.ListWholesaleRequests(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.ListWholesaleRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWholesaleRequests`: ListWholesaleRequestsResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.ListWholesaleRequests`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWholesaleRequestsRequest struct via the builder pattern


### Return type

**ListWholesaleRequestsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWholesaleSetupSteps

> ListWholesaleSetupStepsResponse ListWholesaleSetupSteps(ctx).Execute()

The wholesale setup checklist for a store: profile, ships-from, payment, returns, a catalog with items, visibility chosen, live. Each step is done because the thing exists.

### Example

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
    resp, r, err := apiClient.WholesaleApi.ListWholesaleSetupSteps(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.ListWholesaleSetupSteps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWholesaleSetupSteps`: ListWholesaleSetupStepsResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.ListWholesaleSetupSteps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWholesaleSetupStepsRequest struct via the builder pattern


### Return type

**ListWholesaleSetupStepsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWholesaleBuyerTerm

> UpdateWholesaleBuyerTermResponse UpdateWholesaleBuyerTerm(ctx, buyerUserId).Execute()

Set a buyer's account terms: netDays (0-120) for paying on account, creditLimitCents (the most they may owe you at once; null = no limit) and discountPercentBps (their negotiated rate).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    buyerUserId := "buyerUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WholesaleApi.UpdateWholesaleBuyerTerm(context.Background(), buyerUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.UpdateWholesaleBuyerTerm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWholesaleBuyerTerm`: UpdateWholesaleBuyerTermResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.UpdateWholesaleBuyerTerm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**buyerUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWholesaleBuyerTermRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateWholesaleBuyerTermResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWholesaleCatalog

> CreateWholesaleCatalogResponse UpdateWholesaleCatalog(ctx, id).Execute()

Edit a catalog: name, slug, description, cover image, visibility, opening-order minimum, trade gate, directory listing, active.

### Example

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
    resp, r, err := apiClient.WholesaleApi.UpdateWholesaleCatalog(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.UpdateWholesaleCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWholesaleCatalog`: CreateWholesaleCatalogResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.UpdateWholesaleCatalog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWholesaleCatalogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWholesaleCatalogResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWholesaleCatalogItem

> UpdateWholesaleCatalogItemResponse UpdateWholesaleCatalogItem(ctx, id).Execute()

Change many catalog items at once: items[] sets each row (trade price, MOQ, case pack, lead time, max per order, backorder, restock date, verified-only, price visibility, preset ids); or platformListingIds + set applies the same values to each, set.pctOffRetailBps pricing each from its retail price.

### Example

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
    resp, r, err := apiClient.WholesaleApi.UpdateWholesaleCatalogItem(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.UpdateWholesaleCatalogItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWholesaleCatalogItem`: UpdateWholesaleCatalogItemResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.UpdateWholesaleCatalogItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWholesaleCatalogItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateWholesaleCatalogItemResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWholesaleCatalogItemByPlatformListingId

> UpdateGroupBuyResponse UpdateWholesaleCatalogItemByPlatformListingId(ctx, id, platformListingId).Execute()

Change one item's trade price and its own trade terms in this catalog (MOQ, case pack, lead time, max per order, backorder, restock date, verified-only, price visibility, preset ids). null clears = inherit from the catalog.

### Example

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
    platformListingId := "platformListingId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WholesaleApi.UpdateWholesaleCatalogItemByPlatformListingId(context.Background(), id, platformListingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.UpdateWholesaleCatalogItemByPlatformListingId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWholesaleCatalogItemByPlatformListingId`: UpdateGroupBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.UpdateWholesaleCatalogItemByPlatformListingId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**platformListingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWholesaleCatalogItemByPlatformListingIdRequest struct via the builder pattern


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


## UpdateWholesaleCatalogItemOrder

> ListWholesaleCatalogItemsResponse UpdateWholesaleCatalogItemOrder(ctx, id).Execute()

Set the merchandising order: platform listing ids, first to last. Items left out keep their relative order after these.

### Example

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
    resp, r, err := apiClient.WholesaleApi.UpdateWholesaleCatalogItemOrder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.UpdateWholesaleCatalogItemOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWholesaleCatalogItemOrder`: ListWholesaleCatalogItemsResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.UpdateWholesaleCatalogItemOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWholesaleCatalogItemOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWholesaleCatalogItemsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWholesaleDrop

> CreateWholesaleDropResponse UpdateWholesaleDrop(ctx, id).Execute()

Edit a trade drop (any create field but kind), schedule it (status \"scheduled\", with dropAt in the future) or cancel it. A live drop can only be cancelled or have endsAt / perCustomerLimit changed.

### Example

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
    resp, r, err := apiClient.WholesaleApi.UpdateWholesaleDrop(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.UpdateWholesaleDrop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWholesaleDrop`: CreateWholesaleDropResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.UpdateWholesaleDrop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWholesaleDropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWholesaleDropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWholesaleGroup

> CreateWholesaleGroupResponse UpdateWholesaleGroup(ctx, id).Execute()

Rename a group, change its discount, or pause it (active=false).

### Example

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
    resp, r, err := apiClient.WholesaleApi.UpdateWholesaleGroup(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.UpdateWholesaleGroup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWholesaleGroup`: CreateWholesaleGroupResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.UpdateWholesaleGroup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWholesaleGroupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWholesaleGroupResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWholesaleListingTier

> UpdateWholesaleListingTierResponse UpdateWholesaleListingTier(ctx, platformListingId).Execute()

Replace one listing's quantity breaks: [{ minQuantity ≥ 2, unitPriceCents }], each cheaper than the last and than the list price. An empty list removes them. Applies to every buyer at checkout.

### Example

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
    resp, r, err := apiClient.WholesaleApi.UpdateWholesaleListingTier(context.Background(), platformListingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.UpdateWholesaleListingTier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWholesaleListingTier`: UpdateWholesaleListingTierResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.UpdateWholesaleListingTier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platformListingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWholesaleListingTierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateWholesaleListingTierResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWholesaleProfile

> UpdateWholesaleProfileResponse UpdateWholesaleProfile(ctx).Execute()

Save a store's trade profile. Omitted fields are unchanged; null clears. shipsFrom is checked by the shipping validator (400 invalid_ships_from). The profile is the bottom of the trade-terms cascade: catalogs and items that leave a term empty inherit it — saving never rewrites catalogs. Preset ids must be your own presets of that kind.

### Example

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
    resp, r, err := apiClient.WholesaleApi.UpdateWholesaleProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.UpdateWholesaleProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWholesaleProfile`: UpdateWholesaleProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.UpdateWholesaleProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWholesaleProfileRequest struct via the builder pattern


### Return type

**UpdateWholesaleProfileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWholesaleReserve

> GetWholesaleReserveResponse UpdateWholesaleReserve(ctx, inventoryItemId).Execute()

Reserve N units of an inventory item for wholesale: retail channels advertise the rest and come down when only the reserve is left. 0 clears it.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    inventoryItemId := "inventoryItemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WholesaleApi.UpdateWholesaleReserve(context.Background(), inventoryItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WholesaleApi.UpdateWholesaleReserve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWholesaleReserve`: GetWholesaleReserveResponse
    fmt.Fprintf(os.Stdout, "Response from `WholesaleApi.UpdateWholesaleReserve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWholesaleReserveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWholesaleReserveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

