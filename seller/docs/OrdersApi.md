# \OrdersApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLiveShipmentLabel**](OrdersApi.md#CreateLiveShipmentLabel) | **Post** /v1/live-shipments/{id}/label | Buy ONE label for a show box, rated for all its orders packed together; every order takes its tracking and the buyer’s extra postage is refunded.
[**CreateLiveShipmentShip**](OrdersApi.md#CreateLiveShipmentShip) | **Post** /v1/live-shipments/{id}/ship | Ship a show box with your own label: every order is marked shipped with this tracking and the buyer’s extra postage is refunded.
[**CreateOrderBulkDelete**](OrdersApi.md#CreateOrderBulkDelete) | **Post** /v1/orders/bulk-delete | Bulk soft- or hard-delete orders by id.
[**CreateOrderBulkExport**](OrdersApi.md#CreateOrderBulkExport) | **Post** /v1/orders/bulk-export | Bulk export selected orders as CSV.
[**CreateOrderBulkMarkDisputed**](OrdersApi.md#CreateOrderBulkMarkDisputed) | **Post** /v1/orders/bulk-mark-disputed | Bulk mark orders as disputed with a shared reason.
[**CreateOrderBulkMarkShipped**](OrdersApi.md#CreateOrderBulkMarkShipped) | **Post** /v1/orders/bulk-mark-shipped | Bulk flip status&#x3D;shipped, clear the reserved-stock marker, and notify each platform.
[**CreateOrderBulkPackingSlip**](OrdersApi.md#CreateOrderBulkPackingSlip) | **Post** /v1/orders/bulk-packing-slips | Bulk packing slips PDF.
[**CreateOrderCancel**](OrdersApi.md#CreateOrderCancel) | **Post** /v1/orders/{id}/cancel | Cancel an order before it ships. Crossly-marketplace orders refund the buyer (also needs orders:refund) and restock; crossposted orders go to the marketplace, or fail not_supported_on_platform (recordOnly: true records a cancel made on the marketplace). Fires order.cancelled.
[**CreateOrderCorrectTracking**](OrdersApi.md#CreateOrderCorrectTracking) | **Post** /v1/orders/{id}/correct-tracking | Correct a wrong tracking number the marketplace already holds.
[**CreateOrderDispute**](OrdersApi.md#CreateOrderDispute) | **Post** /v1/orders/{id}/dispute | Flag an order as disputed.
[**CreateOrderImport**](OrdersApi.md#CreateOrderImport) | **Post** /v1/orders/import | Manually (re-)pull one or more connected platforms&#39; order history for a day-window.
[**CreateOrderLabel**](OrdersApi.md#CreateOrderLabel) | **Post** /v1/orders/{id}/label | Purchase a Crossly Shipping label. An order that already has a label answers 409 label_exists unless allowSecondLabel is true.
[**CreateOrderLabelVoid**](OrdersApi.md#CreateOrderLabelVoid) | **Post** /v1/orders/{id}/label/void | Void a Crossly Shipping label and refund it to the shipping wallet.
[**CreateOrderLinkListing**](OrdersApi.md#CreateOrderLinkListing) | **Post** /v1/orders/{id}/link-listing | Link an unmatched sale to one of your listings (marks it sold, delists it elsewhere).
[**CreateOrderMessage**](OrdersApi.md#CreateOrderMessage) | **Post** /v1/orders/{id}/message | Message the order&#39;s buyer on the marketplace.
[**CreateOrderOrderChat**](OrdersApi.md#CreateOrderOrderChat) | **Post** /v1/orders/{id}/order-chat | Send a message to the order&#39;s buyer chat on the marketplace. Mirrored into the Crossly inbox.
[**CreateOrderPickup**](OrdersApi.md#CreateOrderPickup) | **Post** /v1/orders/pickups | Book a carrier pickup for labelled orders (one carrier per pickup; window at least an hour, up to 14 days ahead).
[**CreateOrderPickupLabel**](OrdersApi.md#CreateOrderPickupLabel) | **Post** /v1/orders/{id}/pickup-label | Buy the carrier label for a Crossly order the buyer collects from a pickup point (debits your shipping wallet). 402 when the wallet is short. An order that already has a label answers 409 label_exists unless allowSecondLabel is true.
[**CreateOrderPlatformLabelPackage**](OrdersApi.md#CreateOrderPlatformLabelPackage) | **Post** /v1/orders/{id}/platform-label/package | Save the package (weight oz, box inches) the marketplace prices the label from.
[**CreateOrderPlatformLabelPackageQuote**](OrdersApi.md#CreateOrderPlatformLabelPackageQuote) | **Post** /v1/orders/{id}/platform-label/package/quote | Quote what a change to the marketplace shipment&#39;s package costs (cents).
[**CreateOrderPlatformLabelPull**](OrdersApi.md#CreateOrderPlatformLabelPull) | **Post** /v1/orders/{id}/platform-label/pull | Pull the marketplace&#39;s prepaid label into Crossly (every platform with a label descriptor).
[**CreateOrderPlatformLabelPurchase**](OrdersApi.md#CreateOrderPlatformLabelPurchase) | **Post** /v1/orders/{id}/platform-label/purchase | Buy a replacement / extra marketplace label, or have the marketplace generate it, and record it.
[**CreateOrderPlatformLabelQuote**](OrdersApi.md#CreateOrderPlatformLabelQuote) | **Post** /v1/orders/{id}/platform-label/quote | Quote what buying a marketplace label costs, before you are charged.
[**CreateOrderPullPlatformLabel**](OrdersApi.md#CreateOrderPullPlatformLabel) | **Post** /v1/orders/{id}/pull-platform-label | Pull a pre-paid label from cookie platforms (Poshmark/Mercari).
[**CreateOrderRate**](OrdersApi.md#CreateOrderRate) | **Post** /v1/orders/{id}/rates | Crossly Shipping rate quote for an order.
[**CreateOrderRefund**](OrdersApi.md#CreateOrderRefund) | **Post** /v1/orders/{id}/refund | DEPRECATED: use POST /v1/orders/{id}/refunds. Full or partial refund (amount in dollars). Same idempotency and cap as /refunds: the same request twice refunds once (Idempotency-Key, else the body), never past what the buyer paid.
[**CreateOrderRefundById**](OrdersApi.md#CreateOrderRefundById) | **Post** /v1/orders/{id}/refunds | Refund an order, in full or in part. &#x60;amount&#x60; is dollars; omit it for everything left (or the line items&#39; value). Crossly-marketplace orders refund the buyer&#39;s card; crossposted orders refund through the marketplace, or fail not_supported_on_platform. Never refunds past the order (refund_exceeds_remaining). The same request twice refunds once: send an Idempotency-Key (without one, the body is the key). A partial refund keeps the order&#39;s status. What the platform takes (full/partial, its own reason keys, a comment) is GET /v1/orders/{id}/refund-options; a request outside it fails partial_refund_unsupported, invalid_refund_reason or refund_exceeds_item_amount. Fires order.refunded.
[**CreateOrderRefundResolve**](OrdersApi.md#CreateOrderRefundResolve) | **Post** /v1/orders/{id}/refunds/{refundId}/resolve | Settle an &#x60;unverified&#x60; refund (one Crossly could not confirm with the marketplace after its request was interrupted) once you have checked the marketplace: succeeded (with the amount actually refunded, dollars) or failed (nothing moved; the same Idempotency-Key can then try again). Crossly-marketplace refunds are confirmed by support instead.
[**CreateOrderScanForm**](OrdersApi.md#CreateOrderScanForm) | **Post** /v1/orders/scan-forms | Create SCAN forms (one per carrier account and ship-from address) for labels the carrier has not scanned yet.
[**CreateOrderTracking**](OrdersApi.md#CreateOrderTracking) | **Post** /v1/orders/{id}/tracking | Submit tracking number + carrier and notify the platform.
[**CreateOrderWhatnotLabel**](OrdersApi.md#CreateOrderWhatnotLabel) | **Post** /v1/orders/{id}/whatnot/label | Generate the Whatnot shipping label and record it on the order.
[**CreateOrderWhatnotPackage**](OrdersApi.md#CreateOrderWhatnotPackage) | **Post** /v1/orders/{id}/whatnot/package | Save a Whatnot shipment package (weight oz, box inches).
[**CreateOrderWhatnotPackageQuote**](OrdersApi.md#CreateOrderWhatnotPackageQuote) | **Post** /v1/orders/{id}/whatnot/package/quote | Quote what a Whatnot package change costs (cents).
[**CreateRentalApprove**](OrdersApi.md#CreateRentalApprove) | **Post** /v1/rentals/{id}/approve | Approve a requested rental (offers with requireApproval): the renter’s held card is charged.
[**CreateRentalCancel**](OrdersApi.md#CreateRentalCancel) | **Post** /v1/rentals/{id}/cancel | Cancel a booked rental that has not shipped: the renter is refunded in full.
[**CreateRentalClaim**](OrdersApi.md#CreateRentalClaim) | **Post** /v1/rentals/{id}/claim | Claim part or all of the deposit (cents, a note, photo URLs): within 3 days of the return, or 14 days after the end if it never came back.
[**CreateRentalDecline**](OrdersApi.md#CreateRentalDecline) | **Post** /v1/rentals/{id}/decline | Decline a requested rental: the hold on the renter’s card is released.
[**CreateRentalDelivered**](OrdersApi.md#CreateRentalDelivered) | **Post** /v1/rentals/{id}/delivered | Mark a rental delivered to the renter.
[**CreateRentalExtensionDecide**](OrdersApi.md#CreateRentalExtensionDecide) | **Post** /v1/rentals/{id}/extensions/{extensionId}/decide | Approve or decline a renter’s request for more days.
[**CreateRentalInspect**](OrdersApi.md#CreateRentalInspect) | **Post** /v1/rentals/{id}/inspect | Record the return inspection: the checklist, photos, and optionally release the deposit.
[**CreateRentalLabel**](OrdersApi.md#CreateRentalLabel) | **Post** /v1/rentals/{id}/labels | Buy both labels for a booked rental: outbound to the renter, and the return label they use.
[**CreateRentalNotReturned**](OrdersApi.md#CreateRentalNotReturned) | **Post** /v1/rentals/{id}/not-returned | Report a rental not returned (14 days past due): the deposit goes to you.
[**CreateRentalReleaseDeposit**](OrdersApi.md#CreateRentalReleaseDeposit) | **Post** /v1/rentals/{id}/release-deposit | Release the deposit now, without a claim.
[**CreateRentalReturned**](OrdersApi.md#CreateRentalReturned) | **Post** /v1/rentals/{id}/returned | Mark a rental back with you. You then have 3 days to claim damage before the deposit goes back.
[**CreateRentalShipped**](OrdersApi.md#CreateRentalShipped) | **Post** /v1/rentals/{id}/shipped | Mark a booked rental shipped, with at least 2 photos of its condition as it leaves (URLs you uploaded) — the evidence any later claim is judged against.
[**GetOrder**](OrdersApi.md#GetOrder) | **Get** /v1/orders/{id} | Get one order, with its sale price, fees, SKU and line items.
[**GetOrderCancelEligibility**](OrdersApi.md#GetOrderCancelEligibility) | **Get** /v1/orders/{id}/cancel-eligibility | Whether this order can be cancelled, and which reason codes the marketplace accepts.
[**GetOrderCount**](OrdersApi.md#GetOrderCount) | **Get** /v1/orders/counts | Tab badge counts by status.
[**GetOrderDeliveryPhoto**](OrdersApi.md#GetOrderDeliveryPhoto) | **Get** /v1/orders/{id}/delivery-photo | The courier&#39;s delivery photo for this order, pulled automatically, with where it came from.
[**GetOrderEvidence**](OrdersApi.md#GetOrderEvidence) | **Get** /v1/orders/{id}/evidence | Everything recorded about how this order was packed and how it arrived.
[**GetOrderLabelPrint**](OrdersApi.md#GetOrderLabelPrint) | **Get** /v1/orders/{id}/label/print | The order&#39;s shipping label as a PDF sized for printing (cropped, upright, fitted).
[**GetOrderOrderChat**](OrdersApi.md#GetOrderOrderChat) | **Get** /v1/orders/{id}/order-chat | The order&#39;s buyer chat thread on the marketplace (separate from listing comments): { orderId, platform, messages: [{ id, sender (buyer|seller|platform), body, sentAt, attachments? }], canSend }, oldest message first.
[**GetOrderPackingSlip**](OrdersApi.md#GetOrderPackingSlip) | **Get** /v1/orders/{id}/packing-slip | Single-order packing slip PDF.
[**GetOrderPlatformLabelOption**](OrdersApi.md#GetOrderPlatformLabelOption) | **Get** /v1/orders/{id}/platform-label/options | What buying a marketplace label offers for this order (weight tiers, mailbox state).
[**GetOrderPlatformLabelPackingSlip**](OrdersApi.md#GetOrderPlatformLabelPackingSlip) | **Get** /v1/orders/{id}/platform-label/packing-slip | A fresh link to the marketplace&#39;s own packing slip PDF.
[**GetOrderPlatformLabelState**](OrdersApi.md#GetOrderPlatformLabelState) | **Get** /v1/orders/{id}/platform-label/state | The marketplace&#39;s shipment for an order (package, label, tracking). Records a label found on it.
[**GetOrderPlatformLabelUrl**](OrdersApi.md#GetOrderPlatformLabelUrl) | **Get** /v1/orders/{id}/platform-label/url | A fresh, short-lived link to the marketplace&#39;s label sheet. Never stored.
[**GetOrderProofOfDelivery**](OrdersApi.md#GetOrderProofOfDelivery) | **Get** /v1/orders/{id}/proof-of-delivery | Proof-of-delivery evidence from the carrier&#39;s scan record (PDF, or JSON).
[**GetOrderRefundOption**](OrdersApi.md#GetOrderRefundOption) | **Get** /v1/orders/{id}/refund-options | What a refund on this order may be, from the platform&#39;s declared capability: whether Crossly can refund it at all, full and/or partial, what a partial is measured against (total &#x3D; what the buyer paid, item &#x3D; the item price only), whether it may precede shipment, the platform&#39;s own reason keys (when listed, a refund&#39;s reason must be one of them), whether a comment is sent, whether the refund IS cancelling the sale (cancelsSale: Mercari, full only), whether the window closes when the order completes (untilCompleted; windowClosed &#x3D; it has) and a one-line note on what the platform does not allow.
[**GetOrderSalesTax**](OrdersApi.md#GetOrderSalesTax) | **Get** /v1/orders/{id}/sales-tax | Who collected sales tax on one Crossly marketplace order, where, and the seller notice.
[**GetOrderScanForm**](OrdersApi.md#GetOrderScanForm) | **Get** /v1/orders/scan-forms/{id} | One SCAN form; a queued form is re-read from the carrier and its PDF appears once built.
[**GetOrderShipment**](OrdersApi.md#GetOrderShipment) | **Get** /v1/orders/{id}/shipments | List the parcels an order shipped in.
[**GetOrderWhatnotPackingSlip**](OrdersApi.md#GetOrderWhatnotPackingSlip) | **Get** /v1/orders/{id}/whatnot/packing-slip | A fresh link to Whatnot&#39;s packing slip PDF.
[**GetOrderWhatnotShipment**](OrdersApi.md#GetOrderWhatnotShipment) | **Get** /v1/orders/{id}/whatnot/shipment | Whatnot shipment for an order (package, label, tracking). Records a label found on it.
[**GetRental**](OrdersApi.md#GetRental) | **Get** /v1/rentals/{id} | One rental: its timeline, both sides’ condition photos, payments, extensions and what can be done next.
[**GetRentalCalendar**](OrdersApi.md#GetRentalCalendar) | **Get** /v1/rentals/calendar | Every rentable item’s bookings (per unit), turnaround buffers and blackouts between two dates (≤ 93 days).
[**ListLayaways**](OrdersApi.md#ListLayaways) | **Get** /v1/layaways | Your items on layaway (active and paid off). Do not ship an active one: its order moves to needs_fulfillment when paid in full.
[**ListLiveShipments**](OrdersApi.md#ListLiveShipments) | **Get** /v1/live-shipments | Your show boxes: every buyer’s orders from one live show, to ship together.
[**ListOrderPickupReady**](OrdersApi.md#ListOrderPickupReady) | **Get** /v1/orders/pickup-ready | Orders with a purchased Crossly Shipping label a courier can collect or a SCAN form can cover, grouped by carrier.
[**ListOrderPickups**](OrdersApi.md#ListOrderPickups) | **Get** /v1/orders/pickups | Booked carrier pickups, newest first.
[**ListOrderRefunds**](OrdersApi.md#ListOrderRefunds) | **Get** /v1/orders/{id}/refunds | Every refund requested on this order through the API or a cancel, newest first (amounts in dollars).
[**ListOrderScanForms**](OrdersApi.md#ListOrderScanForms) | **Get** /v1/orders/scan-forms | SCAN forms, newest first.
[**ListOrderUnits**](OrdersApi.md#ListOrderUnits) | **Get** /v1/orders/{id}/units | List the identified units that shipped on an order.
[**ListOrders**](OrdersApi.md#ListOrders) | **Get** /v1/orders | List orders, each with its sale price, fees, SKU and line items.
[**ListRentalHistory**](OrdersApi.md#ListRentalHistory) | **Get** /v1/rentals/history | Active rentals and history, filtered by view, status, dates, item or a search.
[**ListRentals**](OrdersApi.md#ListRentals) | **Get** /v1/rentals | Rentals of your items, with the ship-to and the outbound label.
[**UpdateOrder**](OrdersApi.md#UpdateOrder) | **Patch** /v1/orders/{id} | Update order (status, notes, tracking).



## CreateLiveShipmentLabel

> CreateLiveShipmentShipResponse CreateLiveShipmentLabel(ctx, id).Execute()

Buy ONE label for a show box, rated for all its orders packed together; every order takes its tracking and the buyer’s extra postage is refunded.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateLiveShipmentLabel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateLiveShipmentLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLiveShipmentLabel`: CreateLiveShipmentShipResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateLiveShipmentLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLiveShipmentLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateLiveShipmentShipResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLiveShipmentShip

> CreateLiveShipmentShipResponse CreateLiveShipmentShip(ctx, id).Execute()

Ship a show box with your own label: every order is marked shipped with this tracking and the buyer’s extra postage is refunded.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateLiveShipmentShip(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateLiveShipmentShip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLiveShipmentShip`: CreateLiveShipmentShipResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateLiveShipmentShip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLiveShipmentShipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateLiveShipmentShipResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderBulkDelete

> CreateOrderBulkDeleteResponse CreateOrderBulkDelete(ctx).Execute()

Bulk soft- or hard-delete orders by id.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderBulkDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderBulkDelete`: CreateOrderBulkDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderBulkDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderBulkDeleteRequest struct via the builder pattern


### Return type

**CreateOrderBulkDeleteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderBulkExport

> string CreateOrderBulkExport(ctx).Execute()

Bulk export selected orders as CSV.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderBulkExport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderBulkExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderBulkExport`: string
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderBulkExport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderBulkExportRequest struct via the builder pattern


### Return type

**string**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderBulkMarkDisputed

> UpdatedResponse CreateOrderBulkMarkDisputed(ctx).Execute()

Bulk mark orders as disputed with a shared reason.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderBulkMarkDisputed(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderBulkMarkDisputed``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderBulkMarkDisputed`: UpdatedResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderBulkMarkDisputed`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderBulkMarkDisputedRequest struct via the builder pattern


### Return type

**UpdatedResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderBulkMarkShipped

> CreateOrderBulkMarkShippedResponse CreateOrderBulkMarkShipped(ctx).Execute()

Bulk flip status=shipped, clear the reserved-stock marker, and notify each platform.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderBulkMarkShipped(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderBulkMarkShipped``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderBulkMarkShipped`: CreateOrderBulkMarkShippedResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderBulkMarkShipped`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderBulkMarkShippedRequest struct via the builder pattern


### Return type

**CreateOrderBulkMarkShippedResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderBulkPackingSlip

> *os.File CreateOrderBulkPackingSlip(ctx).Execute()

Bulk packing slips PDF.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderBulkPackingSlip(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderBulkPackingSlip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderBulkPackingSlip`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderBulkPackingSlip`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderBulkPackingSlipRequest struct via the builder pattern


### Return type

***os.File**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderCancel

> CreateOrderCancelResponse CreateOrderCancel(ctx, id).Execute()

Cancel an order before it ships. Crossly-marketplace orders refund the buyer (also needs orders:refund) and restock; crossposted orders go to the marketplace, or fail not_supported_on_platform (recordOnly: true records a cancel made on the marketplace). Fires order.cancelled.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderCancel`: CreateOrderCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderCancelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderCorrectTracking

> CreateOrderCorrectTrackingResponse CreateOrderCorrectTracking(ctx, id).Execute()

Correct a wrong tracking number the marketplace already holds.



### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderCorrectTracking(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderCorrectTracking``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderCorrectTracking`: CreateOrderCorrectTrackingResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderCorrectTracking`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderCorrectTrackingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderCorrectTrackingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderDispute

> CreateOrderDisputeResponse CreateOrderDispute(ctx, id).Execute()

Flag an order as disputed.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderDispute(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderDispute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderDispute`: CreateOrderDisputeResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderDispute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderDisputeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderDisputeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderImport

> CreateOrderImportResponse CreateOrderImport(ctx).Execute()

Manually (re-)pull one or more connected platforms' order history for a day-window.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderImport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderImport`: CreateOrderImportResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderImport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderImportRequest struct via the builder pattern


### Return type

**CreateOrderImportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderLabel

> CreateOrderLabelResponse CreateOrderLabel(ctx, id).Execute()

Purchase a Crossly Shipping label. An order that already has a label answers 409 label_exists unless allowSecondLabel is true.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderLabel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderLabel`: CreateOrderLabelResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderLabelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderLabelVoid

> CreateOrderLabelVoidResponse CreateOrderLabelVoid(ctx, id).Execute()

Void a Crossly Shipping label and refund it to the shipping wallet.



### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderLabelVoid(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderLabelVoid``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderLabelVoid`: CreateOrderLabelVoidResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderLabelVoid`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderLabelVoidRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderLabelVoidResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderLinkListing

> CreateOrderLinkListingResponse CreateOrderLinkListing(ctx, id).Execute()

Link an unmatched sale to one of your listings (marks it sold, delists it elsewhere).



### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderLinkListing(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderLinkListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderLinkListing`: CreateOrderLinkListingResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderLinkListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderLinkListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderLinkListingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderMessage

> CreateOrderMessageResponse CreateOrderMessage(ctx, id).Execute()

Message the order's buyer on the marketplace.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderMessage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderMessage`: CreateOrderMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderMessageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderOrderChat

> OkResponse CreateOrderOrderChat(ctx, id).CreateOrderOrderChatRequest(createOrderOrderChatRequest).Execute()

Send a message to the order's buyer chat on the marketplace. Mirrored into the Crossly inbox.

### Example

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
    createOrderOrderChatRequest := *openapiclient.NewCreateOrderOrderChatRequest("Body_example") // CreateOrderOrderChatRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.CreateOrderOrderChat(context.Background(), id).CreateOrderOrderChatRequest(createOrderOrderChatRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderOrderChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderOrderChat`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderOrderChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderOrderChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createOrderOrderChatRequest** | **CreateOrderOrderChatRequest** |  | 

### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderPickup

> CreateOrderPickupResponse CreateOrderPickup(ctx).Execute()

Book a carrier pickup for labelled orders (one carrier per pickup; window at least an hour, up to 14 days ahead).

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderPickup(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderPickup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderPickup`: CreateOrderPickupResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderPickup`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderPickupRequest struct via the builder pattern


### Return type

**CreateOrderPickupResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderPickupLabel

> CreateOrderPickupLabelResponse CreateOrderPickupLabel(ctx, id).Execute()

Buy the carrier label for a Crossly order the buyer collects from a pickup point (debits your shipping wallet). 402 when the wallet is short. An order that already has a label answers 409 label_exists unless allowSecondLabel is true.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderPickupLabel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderPickupLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderPickupLabel`: CreateOrderPickupLabelResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderPickupLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderPickupLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderPickupLabelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderPlatformLabelPackage

> GetOrderWhatnotShipmentResponse CreateOrderPlatformLabelPackage(ctx, id).Execute()

Save the package (weight oz, box inches) the marketplace prices the label from.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderPlatformLabelPackage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderPlatformLabelPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderPlatformLabelPackage`: GetOrderWhatnotShipmentResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderPlatformLabelPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderPlatformLabelPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderWhatnotShipmentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderPlatformLabelPackageQuote

> CreateOrderWhatnotPackageQuoteResponse CreateOrderPlatformLabelPackageQuote(ctx, id).Execute()

Quote what a change to the marketplace shipment's package costs (cents).

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderPlatformLabelPackageQuote(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderPlatformLabelPackageQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderPlatformLabelPackageQuote`: CreateOrderWhatnotPackageQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderPlatformLabelPackageQuote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderPlatformLabelPackageQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderWhatnotPackageQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderPlatformLabelPull

> CreateOrderPlatformLabelPullResponse CreateOrderPlatformLabelPull(ctx, id).Execute()

Pull the marketplace's prepaid label into Crossly (every platform with a label descriptor).

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderPlatformLabelPull(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderPlatformLabelPull``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderPlatformLabelPull`: CreateOrderPlatformLabelPullResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderPlatformLabelPull`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderPlatformLabelPullRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderPlatformLabelPullResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderPlatformLabelPurchase

> CreateOrderPlatformLabelPurchaseResponse CreateOrderPlatformLabelPurchase(ctx, id).Execute()

Buy a replacement / extra marketplace label, or have the marketplace generate it, and record it.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderPlatformLabelPurchase(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderPlatformLabelPurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderPlatformLabelPurchase`: CreateOrderPlatformLabelPurchaseResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderPlatformLabelPurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderPlatformLabelPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderPlatformLabelPurchaseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderPlatformLabelQuote

> CreateOrderPlatformLabelQuoteResponse CreateOrderPlatformLabelQuote(ctx, id).Execute()

Quote what buying a marketplace label costs, before you are charged.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderPlatformLabelQuote(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderPlatformLabelQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderPlatformLabelQuote`: CreateOrderPlatformLabelQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderPlatformLabelQuote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderPlatformLabelQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderPlatformLabelQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderPullPlatformLabel

> CreateOrderPullPlatformLabelResponse CreateOrderPullPlatformLabel(ctx, id).Execute()

Pull a pre-paid label from cookie platforms (Poshmark/Mercari).

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderPullPlatformLabel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderPullPlatformLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderPullPlatformLabel`: CreateOrderPullPlatformLabelResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderPullPlatformLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderPullPlatformLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderPullPlatformLabelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderRate

> CreateOrderRateResponse CreateOrderRate(ctx, id).Execute()

Crossly Shipping rate quote for an order.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderRate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderRate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderRate`: CreateOrderRateResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderRate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderRateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderRefund

> CreateOrderRefundResponse CreateOrderRefund(ctx, id).Execute()

DEPRECATED: use POST /v1/orders/{id}/refunds. Full or partial refund (amount in dollars). Same idempotency and cap as /refunds: the same request twice refunds once (Idempotency-Key, else the body), never past what the buyer paid.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderRefund(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderRefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderRefund`: CreateOrderRefundResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRefundRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderRefundResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderRefundById

> CreateOrderRefundByIdResponse CreateOrderRefundById(ctx, id).Execute()

Refund an order, in full or in part. `amount` is dollars; omit it for everything left (or the line items' value). Crossly-marketplace orders refund the buyer's card; crossposted orders refund through the marketplace, or fail not_supported_on_platform. Never refunds past the order (refund_exceeds_remaining). The same request twice refunds once: send an Idempotency-Key (without one, the body is the key). A partial refund keeps the order's status. What the platform takes (full/partial, its own reason keys, a comment) is GET /v1/orders/{id}/refund-options; a request outside it fails partial_refund_unsupported, invalid_refund_reason or refund_exceeds_item_amount. Fires order.refunded.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderRefundById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderRefundById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderRefundById`: CreateOrderRefundByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderRefundById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRefundByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderRefundByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderRefundResolve

> ListOrderRefundsItem CreateOrderRefundResolve(ctx, id, refundId).Execute()

Settle an `unverified` refund (one Crossly could not confirm with the marketplace after its request was interrupted) once you have checked the marketplace: succeeded (with the amount actually refunded, dollars) or failed (nothing moved; the same Idempotency-Key can then try again). Crossly-marketplace refunds are confirmed by support instead.

### Example

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
    refundId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.CreateOrderRefundResolve(context.Background(), id, refundId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderRefundResolve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderRefundResolve`: ListOrderRefundsItem
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderRefundResolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**refundId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderRefundResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**ListOrderRefundsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderScanForm

> CreateOrderScanFormResponse CreateOrderScanForm(ctx).Execute()

Create SCAN forms (one per carrier account and ship-from address) for labels the carrier has not scanned yet.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderScanForm(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderScanForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderScanForm`: CreateOrderScanFormResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderScanForm`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderScanFormRequest struct via the builder pattern


### Return type

**CreateOrderScanFormResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderTracking

> CreateOrderTrackingResponse CreateOrderTracking(ctx, id).Execute()

Submit tracking number + carrier and notify the platform.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderTracking(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderTracking``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderTracking`: CreateOrderTrackingResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderTracking`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderTrackingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderTrackingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderWhatnotLabel

> CreateOrderWhatnotLabelResponse CreateOrderWhatnotLabel(ctx, id).Execute()

Generate the Whatnot shipping label and record it on the order.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderWhatnotLabel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderWhatnotLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderWhatnotLabel`: CreateOrderWhatnotLabelResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderWhatnotLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderWhatnotLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderWhatnotLabelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderWhatnotPackage

> GetOrderWhatnotShipmentResponse CreateOrderWhatnotPackage(ctx, id).Execute()

Save a Whatnot shipment package (weight oz, box inches).

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderWhatnotPackage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderWhatnotPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderWhatnotPackage`: GetOrderWhatnotShipmentResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderWhatnotPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderWhatnotPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderWhatnotShipmentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderWhatnotPackageQuote

> CreateOrderWhatnotPackageQuoteResponse CreateOrderWhatnotPackageQuote(ctx, id).Execute()

Quote what a Whatnot package change costs (cents).

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateOrderWhatnotPackageQuote(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateOrderWhatnotPackageQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderWhatnotPackageQuote`: CreateOrderWhatnotPackageQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateOrderWhatnotPackageQuote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderWhatnotPackageQuoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderWhatnotPackageQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateRentalApprove

> RentalResponse CreateRentalApprove(ctx, id).Execute()

Approve a requested rental (offers with requireApproval): the renter’s held card is charged.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateRentalApprove(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateRentalApprove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRentalApprove`: RentalResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateRentalApprove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRentalApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**RentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateRentalCancel

> RentalResponse CreateRentalCancel(ctx, id).Execute()

Cancel a booked rental that has not shipped: the renter is refunded in full.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateRentalCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateRentalCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRentalCancel`: RentalResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateRentalCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRentalCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**RentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateRentalClaim

> RentalResponse CreateRentalClaim(ctx, id).Execute()

Claim part or all of the deposit (cents, a note, photo URLs): within 3 days of the return, or 14 days after the end if it never came back.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateRentalClaim(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateRentalClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRentalClaim`: RentalResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateRentalClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRentalClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**RentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateRentalDecline

> RentalResponse CreateRentalDecline(ctx, id).Execute()

Decline a requested rental: the hold on the renter’s card is released.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateRentalDecline(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateRentalDecline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRentalDecline`: RentalResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateRentalDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRentalDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**RentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateRentalDelivered

> RentalResponse CreateRentalDelivered(ctx, id).Execute()

Mark a rental delivered to the renter.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateRentalDelivered(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateRentalDelivered``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRentalDelivered`: RentalResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateRentalDelivered`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRentalDeliveredRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**RentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateRentalExtensionDecide

> CreateRentalExtensionDecideResponse CreateRentalExtensionDecide(ctx, id, extensionId).Execute()

Approve or decline a renter’s request for more days.

### Example

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
    extensionId := "extensionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.CreateRentalExtensionDecide(context.Background(), id, extensionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateRentalExtensionDecide``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRentalExtensionDecide`: CreateRentalExtensionDecideResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateRentalExtensionDecide`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**extensionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRentalExtensionDecideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateRentalExtensionDecideResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateRentalInspect

> RentalResponse CreateRentalInspect(ctx, id).Execute()

Record the return inspection: the checklist, photos, and optionally release the deposit.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateRentalInspect(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateRentalInspect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRentalInspect`: RentalResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateRentalInspect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRentalInspectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**RentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateRentalLabel

> RentalResponse CreateRentalLabel(ctx, id).Execute()

Buy both labels for a booked rental: outbound to the renter, and the return label they use.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateRentalLabel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateRentalLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRentalLabel`: RentalResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateRentalLabel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRentalLabelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**RentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateRentalNotReturned

> RentalResponse CreateRentalNotReturned(ctx, id).Execute()

Report a rental not returned (14 days past due): the deposit goes to you.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateRentalNotReturned(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateRentalNotReturned``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRentalNotReturned`: RentalResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateRentalNotReturned`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRentalNotReturnedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**RentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateRentalReleaseDeposit

> RentalResponse CreateRentalReleaseDeposit(ctx, id).Execute()

Release the deposit now, without a claim.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateRentalReleaseDeposit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateRentalReleaseDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRentalReleaseDeposit`: RentalResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateRentalReleaseDeposit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRentalReleaseDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**RentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateRentalReturned

> RentalResponse CreateRentalReturned(ctx, id).Execute()

Mark a rental back with you. You then have 3 days to claim damage before the deposit goes back.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateRentalReturned(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateRentalReturned``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRentalReturned`: RentalResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateRentalReturned`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRentalReturnedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**RentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateRentalShipped

> RentalResponse CreateRentalShipped(ctx, id).Execute()

Mark a booked rental shipped, with at least 2 photos of its condition as it leaves (URLs you uploaded) — the evidence any later claim is judged against.

### Example

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
    resp, r, err := apiClient.OrdersApi.CreateRentalShipped(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.CreateRentalShipped``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateRentalShipped`: RentalResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.CreateRentalShipped`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRentalShippedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**RentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrder

> GetOrderResponse GetOrder(ctx, id).Execute()

Get one order, with its sale price, fees, SKU and line items.



### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrder`: GetOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderCancelEligibility

> GetOrderCancelEligibilityResponse GetOrderCancelEligibility(ctx, id).Execute()

Whether this order can be cancelled, and which reason codes the marketplace accepts.

### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderCancelEligibility(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderCancelEligibility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderCancelEligibility`: GetOrderCancelEligibilityResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderCancelEligibility`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderCancelEligibilityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderCancelEligibilityResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderCount

> map[string]float32 GetOrderCount(ctx).Execute()

Tab badge counts by status.

### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderCount`: map[string]float32
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderCountRequest struct via the builder pattern


### Return type

**map[string]float32**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderDeliveryPhoto

> GetOrderDeliveryPhotoResponse GetOrderDeliveryPhoto(ctx, id).Execute()

The courier's delivery photo for this order, pulled automatically, with where it came from.



### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderDeliveryPhoto(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderDeliveryPhoto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderDeliveryPhoto`: GetOrderDeliveryPhotoResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderDeliveryPhoto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderDeliveryPhotoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderDeliveryPhotoResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderEvidence

> GetOrderEvidenceResponse GetOrderEvidence(ctx, id).Execute()

Everything recorded about how this order was packed and how it arrived.



### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderEvidence(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderEvidence``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderEvidence`: GetOrderEvidenceResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderEvidence`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderEvidenceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderEvidenceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderLabelPrint

> *os.File GetOrderLabelPrint(ctx, id).Size(size).Flip(flip).ParcelId(parcelId).Execute()

The order's shipping label as a PDF sized for printing (cropped, upright, fitted).

### Example

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
    size := "size_example" // string | Paper. Default 4x6. (optional)
    flip := "flip_example" // string | Turn the label a half turn. (optional)
    parcelId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | A specific parcel's label; default the order's primary. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.GetOrderLabelPrint(context.Background(), id).Size(size).Flip(flip).ParcelId(parcelId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderLabelPrint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderLabelPrint`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderLabelPrint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderLabelPrintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **string** | Paper. Default 4x6. | 
 **flip** | **string** | Turn the label a half turn. | 
 **parcelId** | **string** | A specific parcel&#39;s label; default the order&#39;s primary. | 

### Return type

***os.File**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderOrderChat

> GetOrderOrderChatResponse GetOrderOrderChat(ctx, id).Execute()

The order's buyer chat thread on the marketplace (separate from listing comments): { orderId, platform, messages: [{ id, sender (buyer|seller|platform), body, sentAt, attachments? }], canSend }, oldest message first.

### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderOrderChat(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderOrderChat``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderOrderChat`: GetOrderOrderChatResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderOrderChat`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderOrderChatRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderOrderChatResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderPackingSlip

> *os.File GetOrderPackingSlip(ctx, id).Execute()

Single-order packing slip PDF.

### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderPackingSlip(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderPackingSlip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderPackingSlip`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderPackingSlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderPackingSlipRequest struct via the builder pattern


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


## GetOrderPlatformLabelOption

> GetOrderPlatformLabelOptionResponse GetOrderPlatformLabelOption(ctx, id).Execute()

What buying a marketplace label offers for this order (weight tiers, mailbox state).

### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderPlatformLabelOption(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderPlatformLabelOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderPlatformLabelOption`: GetOrderPlatformLabelOptionResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderPlatformLabelOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderPlatformLabelOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderPlatformLabelOptionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderPlatformLabelPackingSlip

> UrlExpiresInSecondsResponse GetOrderPlatformLabelPackingSlip(ctx, id).Execute()

A fresh link to the marketplace's own packing slip PDF.

### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderPlatformLabelPackingSlip(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderPlatformLabelPackingSlip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderPlatformLabelPackingSlip`: UrlExpiresInSecondsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderPlatformLabelPackingSlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderPlatformLabelPackingSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UrlExpiresInSecondsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderPlatformLabelState

> GetOrderWhatnotShipmentResponse GetOrderPlatformLabelState(ctx, id).Execute()

The marketplace's shipment for an order (package, label, tracking). Records a label found on it.

### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderPlatformLabelState(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderPlatformLabelState``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderPlatformLabelState`: GetOrderWhatnotShipmentResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderPlatformLabelState`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderPlatformLabelStateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderWhatnotShipmentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderPlatformLabelUrl

> UrlExpiresInSecondsResponse GetOrderPlatformLabelUrl(ctx, id).Execute()

A fresh, short-lived link to the marketplace's label sheet. Never stored.

### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderPlatformLabelUrl(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderPlatformLabelUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderPlatformLabelUrl`: UrlExpiresInSecondsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderPlatformLabelUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderPlatformLabelUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UrlExpiresInSecondsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderProofOfDelivery

> GetOrderProofOfDeliveryResponse GetOrderProofOfDelivery(ctx, id).Execute()

Proof-of-delivery evidence from the carrier's scan record (PDF, or JSON).

### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderProofOfDelivery(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderProofOfDelivery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderProofOfDelivery`: GetOrderProofOfDeliveryResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderProofOfDelivery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderProofOfDeliveryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderProofOfDeliveryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderRefundOption

> GetOrderRefundOptionResponse GetOrderRefundOption(ctx, id).Execute()

What a refund on this order may be, from the platform's declared capability: whether Crossly can refund it at all, full and/or partial, what a partial is measured against (total = what the buyer paid, item = the item price only), whether it may precede shipment, the platform's own reason keys (when listed, a refund's reason must be one of them), whether a comment is sent, whether the refund IS cancelling the sale (cancelsSale: Mercari, full only), whether the window closes when the order completes (untilCompleted; windowClosed = it has) and a one-line note on what the platform does not allow.

### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderRefundOption(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderRefundOption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderRefundOption`: GetOrderRefundOptionResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderRefundOption`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderRefundOptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderRefundOptionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderSalesTax

> GetOrderSalesTaxResponse GetOrderSalesTax(ctx, id).Execute()

Who collected sales tax on one Crossly marketplace order, where, and the seller notice.

### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderSalesTax(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderSalesTax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderSalesTax`: GetOrderSalesTaxResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderSalesTax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderSalesTaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderSalesTaxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderScanForm

> GetOrderScanFormResponse GetOrderScanForm(ctx, id).Execute()

One SCAN form; a queued form is re-read from the carrier and its PDF appears once built.

### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderScanForm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderScanForm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderScanForm`: GetOrderScanFormResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderScanForm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderScanFormRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderScanFormResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderShipment

> GetOrderShipmentResponse GetOrderShipment(ctx, id).Execute()

List the parcels an order shipped in.



### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderShipment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderShipment`: GetOrderShipmentResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderShipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderShipmentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderWhatnotPackingSlip

> UrlExpiresInSecondsResponse GetOrderWhatnotPackingSlip(ctx, id).Execute()

A fresh link to Whatnot's packing slip PDF.

### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderWhatnotPackingSlip(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderWhatnotPackingSlip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderWhatnotPackingSlip`: UrlExpiresInSecondsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderWhatnotPackingSlip`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderWhatnotPackingSlipRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UrlExpiresInSecondsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderWhatnotShipment

> GetOrderWhatnotShipmentResponse GetOrderWhatnotShipment(ctx, id).Execute()

Whatnot shipment for an order (package, label, tracking). Records a label found on it.

### Example

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
    resp, r, err := apiClient.OrdersApi.GetOrderWhatnotShipment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetOrderWhatnotShipment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderWhatnotShipment`: GetOrderWhatnotShipmentResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetOrderWhatnotShipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderWhatnotShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderWhatnotShipmentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetRental

> GetRentalResponse GetRental(ctx, id).Execute()

One rental: its timeline, both sides’ condition photos, payments, extensions and what can be done next.

### Example

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
    resp, r, err := apiClient.OrdersApi.GetRental(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetRental``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRental`: GetRentalResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetRental`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRentalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetRentalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetRentalCalendar

> GetRentalCalendarResponse GetRentalCalendar(ctx).Execute()

Every rentable item’s bookings (per unit), turnaround buffers and blackouts between two dates (≤ 93 days).

### Example

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
    resp, r, err := apiClient.OrdersApi.GetRentalCalendar(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.GetRentalCalendar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetRentalCalendar`: GetRentalCalendarResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.GetRentalCalendar`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetRentalCalendarRequest struct via the builder pattern


### Return type

**GetRentalCalendarResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListLayaways

> ListLayawaysResponse ListLayaways(ctx).Execute()

Your items on layaway (active and paid off). Do not ship an active one: its order moves to needs_fulfillment when paid in full.

### Example

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
    resp, r, err := apiClient.OrdersApi.ListLayaways(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ListLayaways``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLayaways`: ListLayawaysResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ListLayaways`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLayawaysRequest struct via the builder pattern


### Return type

**ListLayawaysResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListLiveShipments

> ListLiveShipmentsResponse ListLiveShipments(ctx).Execute()

Your show boxes: every buyer’s orders from one live show, to ship together.

### Example

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
    resp, r, err := apiClient.OrdersApi.ListLiveShipments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ListLiveShipments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLiveShipments`: ListLiveShipmentsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ListLiveShipments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLiveShipmentsRequest struct via the builder pattern


### Return type

**ListLiveShipmentsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListOrderPickupReady

> ListOrderPickupReadyResponse ListOrderPickupReady(ctx).Execute()

Orders with a purchased Crossly Shipping label a courier can collect or a SCAN form can cover, grouped by carrier.

### Example

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
    resp, r, err := apiClient.OrdersApi.ListOrderPickupReady(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ListOrderPickupReady``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrderPickupReady`: ListOrderPickupReadyResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ListOrderPickupReady`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrderPickupReadyRequest struct via the builder pattern


### Return type

**ListOrderPickupReadyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListOrderPickups

> ListOrderPickupsResponse ListOrderPickups(ctx).Execute()

Booked carrier pickups, newest first.

### Example

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
    resp, r, err := apiClient.OrdersApi.ListOrderPickups(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ListOrderPickups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrderPickups`: ListOrderPickupsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ListOrderPickups`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrderPickupsRequest struct via the builder pattern


### Return type

**ListOrderPickupsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListOrderRefunds

> ListOrderRefundsResponse ListOrderRefunds(ctx, id).Execute()

Every refund requested on this order through the API or a cancel, newest first (amounts in dollars).

### Example

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
    resp, r, err := apiClient.OrdersApi.ListOrderRefunds(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ListOrderRefunds``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrderRefunds`: ListOrderRefundsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ListOrderRefunds`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrderRefundsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListOrderRefundsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListOrderScanForms

> ListOrderScanFormsResponse ListOrderScanForms(ctx).Execute()

SCAN forms, newest first.

### Example

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
    resp, r, err := apiClient.OrdersApi.ListOrderScanForms(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ListOrderScanForms``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrderScanForms`: ListOrderScanFormsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ListOrderScanForms`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListOrderScanFormsRequest struct via the builder pattern


### Return type

**ListOrderScanFormsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListOrderUnits

> ListOrderUnitsResponse ListOrderUnits(ctx, id).Execute()

List the identified units that shipped on an order.



### Example

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
    resp, r, err := apiClient.OrdersApi.ListOrderUnits(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ListOrderUnits``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrderUnits`: ListOrderUnitsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ListOrderUnits`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListOrderUnitsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListOrderUnitsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListOrders

> ListOrdersResponse ListOrders(ctx).Page(page).Limit(limit).Status(status).Platform(platform).Channel(channel).Details(details).Unlinked(unlinked).UpdatedSince(updatedSince).Execute()

List orders, each with its sale price, fees, SKU and line items.



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
    status := "status_example" // string |  (optional)
    platform := "platform_example" // string |  (optional)
    channel := "channel_example" // string | Only orders sold through this business channel (e.g. wholesale). (optional)
    details := "details_example" // string | pending = only orders whose sale details (price, buyer) have not loaded yet. (optional)
    unlinked := "unlinked_example" // string | 1 or true = only orders whose sale is not linked to any listing (an orphan sale). (optional)
    updatedSince := time.Now() // time.Time | Only orders created or changed at or after this time (ISO 8601): poll with the time of your last call to catch new orders and status/tracking changes. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.OrdersApi.ListOrders(context.Background()).Page(page).Limit(limit).Status(status).Platform(platform).Channel(channel).Details(details).Unlinked(unlinked).UpdatedSince(updatedSince).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ListOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListOrders`: ListOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ListOrders`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListOrdersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 25]
 **status** | **string** |  | 
 **platform** | **string** |  | 
 **channel** | **string** | Only orders sold through this business channel (e.g. wholesale). | 
 **details** | **string** | pending &#x3D; only orders whose sale details (price, buyer) have not loaded yet. | 
 **unlinked** | **string** | 1 or true &#x3D; only orders whose sale is not linked to any listing (an orphan sale). | 
 **updatedSince** | **time.Time** | Only orders created or changed at or after this time (ISO 8601): poll with the time of your last call to catch new orders and status/tracking changes. | 

### Return type

**ListOrdersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListRentalHistory

> ListRentalHistoryResponse ListRentalHistory(ctx).Execute()

Active rentals and history, filtered by view, status, dates, item or a search.

### Example

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
    resp, r, err := apiClient.OrdersApi.ListRentalHistory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ListRentalHistory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRentalHistory`: ListRentalHistoryResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ListRentalHistory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRentalHistoryRequest struct via the builder pattern


### Return type

**ListRentalHistoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListRentals

> ListRentalsResponse ListRentals(ctx).Execute()

Rentals of your items, with the ship-to and the outbound label.

### Example

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
    resp, r, err := apiClient.OrdersApi.ListRentals(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.ListRentals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListRentals`: ListRentalsResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.ListRentals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListRentalsRequest struct via the builder pattern


### Return type

**ListRentalsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateOrder

> UpdateOrderResponse UpdateOrder(ctx, id).Execute()

Update order (status, notes, tracking).

### Example

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
    resp, r, err := apiClient.OrdersApi.UpdateOrder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `OrdersApi.UpdateOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateOrder`: UpdateOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `OrdersApi.UpdateOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateOrderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

