# \WorldsMoneyApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorldEarningPayoutAccountOnboard**](WorldsMoneyApi.md#CreateWorldEarningPayoutAccountOnboard) | **Post** /v1/worlds/earnings/payout-account/onboard | Start or resume Stripe onboarding for your creator payout account. Returns a one-time Stripe URL. Body: { country? } (fixed once the account exists).
[**CreateWorldExperienceDropTable**](WorldsMoneyApi.md#CreateWorldExperienceDropTable) | **Post** /v1/worlds/experiences/{experienceId}/drop-tables | Declare a drop table for a paid random item: cosmetic or in-experience outcomes only. Crossly pre-simulates and hashes it.
[**CreateWorldExperienceProduct**](WorldsMoneyApi.md#CreateWorldExperienceProduct) | **Post** /v1/worlds/experiences/{experienceId}/products | Create a product: price 50-100,000 cents (a private server may be free), optional per-country prices. Paid random items need an approved drop table and are never tradeable.
[**CreateWorldExperienceProductArchive**](WorldsMoneyApi.md#CreateWorldExperienceProductArchive) | **Post** /v1/worlds/experiences/{experienceId}/products/{productId}/archive | Take a product off sale for good. Owned passes and running subscriptions keep working.
[**CreateWorldExperiencePurchase**](WorldsMoneyApi.md#CreateWorldExperiencePurchase) | **Post** /v1/worlds/experiences/{experienceId}/purchases | Open a purchase sheet for a product (paid access, a pass…). Returns checkoutPath: the player pays only in the Crossly checkout.
[**CreateWorldExperienceReceiptAck**](WorldsMoneyApi.md#CreateWorldExperienceReceiptAck) | **Post** /v1/worlds/experiences/{experienceId}/receipts/{orderId}/ack | Acknowledge a grant your back-end made off-server. Idempotent: a second ack reports duplicate.
[**CreateWorldMeReceiptRefund**](WorldsMoneyApi.md#CreateWorldMeReceiptRefund) | **Post** /v1/worlds/me/receipts/{receiptId}/refund | Refund a developer product bought by mistake: within the self-refund window and before the experience granted it.
[**CreateWorldMeSubscriptionCancel**](WorldsMoneyApi.md#CreateWorldMeSubscriptionCancel) | **Post** /v1/worlds/me/subscriptions/{id}/cancel | Cancel: no more charges; the benefit lasts to the end of the paid month.
[**CreateWorldMeSubscriptionPrivateServerCode**](WorldsMoneyApi.md#CreateWorldMeSubscriptionPrivateServerCode) | **Post** /v1/worlds/me/subscriptions/{id}/private-server-code | A fresh access code for your private server (the previous one stops working). Join with reservedCode.
[**CreateWorldPurchaseCancel**](WorldsMoneyApi.md#CreateWorldPurchaseCancel) | **Post** /v1/worlds/purchases/{intentId}/cancel | Cancel an unpaid purchase sheet.
[**GetWorldEarning**](WorldsMoneyApi.md#GetWorldEarning) | **Get** /v1/worlds/earnings | Your Crossly Carnival earnings: pending, available, held, paid; next payout; what cashing out still needs.
[**GetWorldEarningPayoutAccount**](WorldsMoneyApi.md#GetWorldEarningPayoutAccount) | **Get** /v1/worlds/earnings/payout-account | Your creator payout account (your own, separate from any store payout account): not_started, pending, restricted, active or payouts_unavailable.
[**GetWorldEconomy**](WorldsMoneyApi.md#GetWorldEconomy) | **Get** /v1/worlds/economy | The published payout formula: splits, thresholds, scheduled changes (90-day notice) and guarantees.
[**GetWorldExperienceRevenue**](WorldsMoneyApi.md#GetWorldExperienceRevenue) | **Get** /v1/worlds/experiences/{experienceId}/revenue | Gross, tax, refunds, chargebacks, net and your share per product, plus pending / available / paid.
[**GetWorldExperienceStore**](WorldsMoneyApi.md#GetWorldExperienceStore) | **Get** /v1/worlds/experiences/{experienceId}/store | An experience&#39;s store as you see it: your regional price, odds for paid random items, what you own.
[**GetWorldPurchase**](WorldsMoneyApi.md#GetWorldPurchase) | **Get** /v1/worlds/purchases/{intentId} | A purchase sheet: item, price, tax, total, odds, status.
[**ListWorldEarningAffiliate**](WorldsMoneyApi.md#ListWorldEarningAffiliate) | **Get** /v1/worlds/earnings/affiliate | Real-goods orders attributed to your games: the fee, your 20% share, and whether it is posted.
[**ListWorldEarningLedger**](WorldsMoneyApi.md#ListWorldEarningLedger) | **Get** /v1/worlds/earnings/ledger | Your creator ledger, newest first: every earning and reversal. Cursor-paged.
[**ListWorldEarningPayouts**](WorldsMoneyApi.md#ListWorldEarningPayouts) | **Get** /v1/worlds/earnings/payouts | Your monthly Crossly Carnival payouts.
[**ListWorldExperienceDropTables**](WorldsMoneyApi.md#ListWorldExperienceDropTables) | **Get** /v1/worlds/experiences/{experienceId}/drop-tables | Declared drop tables with their odds, simulation (10^6 draws) and hash.
[**ListWorldExperienceProducts**](WorldsMoneyApi.md#ListWorldExperienceProducts) | **Get** /v1/worlds/experiences/{experienceId}/products | What an experience sells: passes, developer products, subscriptions, private servers, paid access (?includeArchived&#x3D;true).
[**ListWorldExperienceReceipts**](WorldsMoneyApi.md#ListWorldExperienceReceipts) | **Get** /v1/worlds/experiences/{experienceId}/receipts | Every purchase: orderId, sub, product, price, net, state (paid/delivered/granted/refunded/revoked). Cursor-paged; ?since&#x3D;.
[**ListWorldMeSubscriptions**](WorldsMoneyApi.md#ListWorldMeSubscriptions) | **Get** /v1/worlds/me/subscriptions | Your experience subscriptions and private servers.
[**UpdateWorldExperienceProduct**](WorldsMoneyApi.md#UpdateWorldExperienceProduct) | **Patch** /v1/worlds/experiences/{experienceId}/products/{productId} | Edit a product: name, description, icon, price, regional prices, hourly limit, on/off sale. Kind and place never change.



## CreateWorldEarningPayoutAccountOnboard

> CreateWorldEarningPayoutAccountOnboardResponse CreateWorldEarningPayoutAccountOnboard(ctx).Execute()

Start or resume Stripe onboarding for your creator payout account. Returns a one-time Stripe URL. Body: { country? } (fixed once the account exists).

### Example

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
    resp, r, err := apiClient.WorldsMoneyApi.CreateWorldEarningPayoutAccountOnboard(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.CreateWorldEarningPayoutAccountOnboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldEarningPayoutAccountOnboard`: CreateWorldEarningPayoutAccountOnboardResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.CreateWorldEarningPayoutAccountOnboard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldEarningPayoutAccountOnboardRequest struct via the builder pattern


### Return type

**CreateWorldEarningPayoutAccountOnboardResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceDropTable

> CreateWorldExperienceDropTableResponse CreateWorldExperienceDropTable(ctx, experienceId).Execute()

Declare a drop table for a paid random item: cosmetic or in-experience outcomes only. Crossly pre-simulates and hashes it.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsMoneyApi.CreateWorldExperienceDropTable(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.CreateWorldExperienceDropTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceDropTable`: CreateWorldExperienceDropTableResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.CreateWorldExperienceDropTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceDropTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceDropTableResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceProduct

> WorldExperienceProductResponse CreateWorldExperienceProduct(ctx, experienceId).Execute()

Create a product: price 50-100,000 cents (a private server may be free), optional per-country prices. Paid random items need an approved drop table and are never tradeable.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsMoneyApi.CreateWorldExperienceProduct(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.CreateWorldExperienceProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceProduct`: WorldExperienceProductResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.CreateWorldExperienceProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceProductResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceProductArchive

> WorldExperienceProductResponse CreateWorldExperienceProductArchive(ctx, experienceId, productId).Execute()

Take a product off sale for good. Owned passes and running subscriptions keep working.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    productId := "productId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsMoneyApi.CreateWorldExperienceProductArchive(context.Background(), experienceId, productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.CreateWorldExperienceProductArchive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceProductArchive`: WorldExperienceProductResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.CreateWorldExperienceProductArchive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceProductArchiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceProductResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePurchase

> CreateWorldExperiencePurchaseResponse CreateWorldExperiencePurchase(ctx, experienceId).Execute()

Open a purchase sheet for a product (paid access, a pass…). Returns checkoutPath: the player pays only in the Crossly checkout.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsMoneyApi.CreateWorldExperiencePurchase(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.CreateWorldExperiencePurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePurchase`: CreateWorldExperiencePurchaseResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.CreateWorldExperiencePurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperiencePurchaseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceReceiptAck

> CreateWorldExperienceReceiptAckResponse CreateWorldExperienceReceiptAck(ctx, experienceId, orderId).Execute()

Acknowledge a grant your back-end made off-server. Idempotent: a second ack reports duplicate.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    orderId := "orderId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsMoneyApi.CreateWorldExperienceReceiptAck(context.Background(), experienceId, orderId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.CreateWorldExperienceReceiptAck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceReceiptAck`: CreateWorldExperienceReceiptAckResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.CreateWorldExperienceReceiptAck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**orderId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceReceiptAckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperienceReceiptAckResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldMeReceiptRefund

> CreateWorldMeReceiptRefundResponse CreateWorldMeReceiptRefund(ctx, receiptId).Execute()

Refund a developer product bought by mistake: within the self-refund window and before the experience granted it.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    receiptId := "receiptId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsMoneyApi.CreateWorldMeReceiptRefund(context.Background(), receiptId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.CreateWorldMeReceiptRefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldMeReceiptRefund`: CreateWorldMeReceiptRefundResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.CreateWorldMeReceiptRefund`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**receiptId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldMeReceiptRefundRequest struct via the builder pattern


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


## CreateWorldMeSubscriptionCancel

> CreateWorldMeSubscriptionCancelResponse CreateWorldMeSubscriptionCancel(ctx, id).Execute()

Cancel: no more charges; the benefit lasts to the end of the paid month.

### Example

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
    resp, r, err := apiClient.WorldsMoneyApi.CreateWorldMeSubscriptionCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.CreateWorldMeSubscriptionCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldMeSubscriptionCancel`: CreateWorldMeSubscriptionCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.CreateWorldMeSubscriptionCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldMeSubscriptionCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldMeSubscriptionCancelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldMeSubscriptionPrivateServerCode

> CreateWorldMeSubscriptionPrivateServerCodeResponse CreateWorldMeSubscriptionPrivateServerCode(ctx, id).Execute()

A fresh access code for your private server (the previous one stops working). Join with reservedCode.

### Example

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
    resp, r, err := apiClient.WorldsMoneyApi.CreateWorldMeSubscriptionPrivateServerCode(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.CreateWorldMeSubscriptionPrivateServerCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldMeSubscriptionPrivateServerCode`: CreateWorldMeSubscriptionPrivateServerCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.CreateWorldMeSubscriptionPrivateServerCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldMeSubscriptionPrivateServerCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldMeSubscriptionPrivateServerCodeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldPurchaseCancel

> CreateWorldSpaceJoinResponse CreateWorldPurchaseCancel(ctx, intentId).Execute()

Cancel an unpaid purchase sheet.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    intentId := "intentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsMoneyApi.CreateWorldPurchaseCancel(context.Background(), intentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.CreateWorldPurchaseCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldPurchaseCancel`: CreateWorldSpaceJoinResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.CreateWorldPurchaseCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**intentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldPurchaseCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldSpaceJoinResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldEarning

> GetWorldEarningResponse GetWorldEarning(ctx).Execute()

Your Crossly Carnival earnings: pending, available, held, paid; next payout; what cashing out still needs.

### Example

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
    resp, r, err := apiClient.WorldsMoneyApi.GetWorldEarning(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.GetWorldEarning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldEarning`: GetWorldEarningResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.GetWorldEarning`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldEarningRequest struct via the builder pattern


### Return type

**GetWorldEarningResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldEarningPayoutAccount

> GetWorldEarningPayoutAccountResponse GetWorldEarningPayoutAccount(ctx).Execute()

Your creator payout account (your own, separate from any store payout account): not_started, pending, restricted, active or payouts_unavailable.

### Example

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
    resp, r, err := apiClient.WorldsMoneyApi.GetWorldEarningPayoutAccount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.GetWorldEarningPayoutAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldEarningPayoutAccount`: GetWorldEarningPayoutAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.GetWorldEarningPayoutAccount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldEarningPayoutAccountRequest struct via the builder pattern


### Return type

**GetWorldEarningPayoutAccountResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldEconomy

> GetWorldEconomyResponse GetWorldEconomy(ctx).Execute()

The published payout formula: splits, thresholds, scheduled changes (90-day notice) and guarantees.

### Example

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
    resp, r, err := apiClient.WorldsMoneyApi.GetWorldEconomy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.GetWorldEconomy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldEconomy`: GetWorldEconomyResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.GetWorldEconomy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldEconomyRequest struct via the builder pattern


### Return type

**GetWorldEconomyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceRevenue

> GetWorldExperienceRevenueResponse GetWorldExperienceRevenue(ctx, experienceId).Execute()

Gross, tax, refunds, chargebacks, net and your share per product, plus pending / available / paid.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsMoneyApi.GetWorldExperienceRevenue(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.GetWorldExperienceRevenue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceRevenue`: GetWorldExperienceRevenueResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.GetWorldExperienceRevenue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceRevenueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceRevenueResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceStore

> GetWorldExperienceStoreResponse GetWorldExperienceStore(ctx, experienceId).Execute()

An experience's store as you see it: your regional price, odds for paid random items, what you own.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsMoneyApi.GetWorldExperienceStore(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.GetWorldExperienceStore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceStore`: GetWorldExperienceStoreResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.GetWorldExperienceStore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceStoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceStoreResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldPurchase

> GetWorldPurchaseResponse GetWorldPurchase(ctx, intentId).Execute()

A purchase sheet: item, price, tax, total, odds, status.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    intentId := "intentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsMoneyApi.GetWorldPurchase(context.Background(), intentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.GetWorldPurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldPurchase`: GetWorldPurchaseResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.GetWorldPurchase`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**intentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldPurchaseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldPurchaseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldEarningAffiliate

> ListWorldEarningAffiliateResponse ListWorldEarningAffiliate(ctx).Execute()

Real-goods orders attributed to your games: the fee, your 20% share, and whether it is posted.

### Example

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
    resp, r, err := apiClient.WorldsMoneyApi.ListWorldEarningAffiliate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.ListWorldEarningAffiliate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldEarningAffiliate`: ListWorldEarningAffiliateResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.ListWorldEarningAffiliate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldEarningAffiliateRequest struct via the builder pattern


### Return type

**ListWorldEarningAffiliateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldEarningLedger

> ListWorldEarningLedgerResponse ListWorldEarningLedger(ctx).Execute()

Your creator ledger, newest first: every earning and reversal. Cursor-paged.

### Example

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
    resp, r, err := apiClient.WorldsMoneyApi.ListWorldEarningLedger(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.ListWorldEarningLedger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldEarningLedger`: ListWorldEarningLedgerResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.ListWorldEarningLedger`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldEarningLedgerRequest struct via the builder pattern


### Return type

**ListWorldEarningLedgerResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldEarningPayouts

> ListWorldEarningPayoutsResponse ListWorldEarningPayouts(ctx).Execute()

Your monthly Crossly Carnival payouts.

### Example

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
    resp, r, err := apiClient.WorldsMoneyApi.ListWorldEarningPayouts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.ListWorldEarningPayouts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldEarningPayouts`: ListWorldEarningPayoutsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.ListWorldEarningPayouts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldEarningPayoutsRequest struct via the builder pattern


### Return type

**ListWorldEarningPayoutsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceDropTables

> ListWorldExperienceDropTablesResponse ListWorldExperienceDropTables(ctx, experienceId).Execute()

Declared drop tables with their odds, simulation (10^6 draws) and hash.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsMoneyApi.ListWorldExperienceDropTables(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.ListWorldExperienceDropTables``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceDropTables`: ListWorldExperienceDropTablesResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.ListWorldExperienceDropTables`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceDropTablesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceDropTablesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceProducts

> ListWorldExperienceProductsResponse ListWorldExperienceProducts(ctx, experienceId).Execute()

What an experience sells: passes, developer products, subscriptions, private servers, paid access (?includeArchived=true).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsMoneyApi.ListWorldExperienceProducts(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.ListWorldExperienceProducts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceProducts`: ListWorldExperienceProductsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.ListWorldExperienceProducts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceProductsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceProductsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceReceipts

> ListWorldExperienceReceiptsResponse ListWorldExperienceReceipts(ctx, experienceId).Execute()

Every purchase: orderId, sub, product, price, net, state (paid/delivered/granted/refunded/revoked). Cursor-paged; ?since=.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsMoneyApi.ListWorldExperienceReceipts(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.ListWorldExperienceReceipts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceReceipts`: ListWorldExperienceReceiptsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.ListWorldExperienceReceipts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceReceiptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceReceiptsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldMeSubscriptions

> ListWorldMeSubscriptionsResponse ListWorldMeSubscriptions(ctx).Execute()

Your experience subscriptions and private servers.

### Example

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
    resp, r, err := apiClient.WorldsMoneyApi.ListWorldMeSubscriptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.ListWorldMeSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldMeSubscriptions`: ListWorldMeSubscriptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.ListWorldMeSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldMeSubscriptionsRequest struct via the builder pattern


### Return type

**ListWorldMeSubscriptionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceProduct

> WorldExperienceProductResponse UpdateWorldExperienceProduct(ctx, experienceId, productId).Execute()

Edit a product: name, description, icon, price, regional prices, hourly limit, on/off sale. Kind and place never change.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    experienceId := "experienceId_example" // string | 
    productId := "productId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsMoneyApi.UpdateWorldExperienceProduct(context.Background(), experienceId, productId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsMoneyApi.UpdateWorldExperienceProduct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceProduct`: WorldExperienceProductResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsMoneyApi.UpdateWorldExperienceProduct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**productId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceProductRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceProductResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

