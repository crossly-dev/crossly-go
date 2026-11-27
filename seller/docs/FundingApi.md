# \FundingApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFundingApply**](FundingApi.md#CreateFundingApply) | **Post** /v1/funding/apply | Apply for Crossly Funding. Decided at once on the checklist rules: approved opens your line (limit &#x3D; 25% of 90 days of settled Crossly payouts), declined lists why, needs_review goes to a person.
[**CreateFundingDraw**](FundingApi.md#CreateFundingDraw) | **Post** /v1/funding/draws | Draw on your line for a supplier order. Send the supplier, the invoice (an uploaded invoiceFileId from POST /v1/funding/invoice-files, a reference, or a URL) and the APR from the quote you accepted (acceptedAprBps); refused if your rate changed. The advance is held as in-platform Crossly balance, never cash. Idempotent on idempotencyKey.
[**CreateFundingDrawApprove**](FundingApi.md#CreateFundingDrawApprove) | **Post** /v1/funding/draws/{id}/approve | Approve a held draw. The two-person rule needs two distinct approvers; the requester never approves their own.
[**CreateFundingDrawQuote**](FundingApi.md#CreateFundingDrawQuote) | **Post** /v1/funding/draws/quote | The disclosure for a draw of amountCents: APR, the simple-interest fee for the full term, the total and the due date, plus your available headroom. Nothing is drawn.
[**CreateFundingDrawReject**](FundingApi.md#CreateFundingDrawReject) | **Post** /v1/funding/draws/{id}/reject | Reject a held draw. Nothing was drawn, so nothing is owed.
[**CreateFundingDrawUse**](FundingApi.md#CreateFundingDrawUse) | **Post** /v1/funding/draws/use | Request a draw for an approved use. Refused by the spend controls, held for approval (stage requested — no debt until disbursed), or disbursed at once as restricted funding CBX. A DEBT with interest: show the disclosure first.
[**CreateFundingInvoiceFile**](FundingApi.md#CreateFundingInvoiceFile) | **Post** /v1/funding/invoice-files | Sign an upload for a supplier invoice: a PDF or image (application/pdf, image/jpeg, image/png, image/webp, image/heic) of at most 15 MB. PUT the bytes to uploadUrl with the same Content-Type and exactly contentLength bytes, then draw with invoiceFileId &#x3D; file.id. Stored privately; one draw per file.
[**CreateFundingPayCrosslyFee**](FundingApi.md#CreateFundingPayCrosslyFee) | **Post** /v1/funding/pay/crossly-fees | Pay an open invoice of your Crossly plan with restricted funding CBX (no card is charged; the invoice is marked paid). Idempotent per invoice.
[**CreateFundingPayShippingLabel**](FundingApi.md#CreateFundingPayShippingLabel) | **Post** /v1/funding/pay/shipping-labels | Top up your shipping wallet with restricted funding CBX; labels you buy are then debited from it. Idempotent on idempotencyKey. Refused when the spend controls forbid shipping labels or it is over your approval threshold with no approved draw.
[**CreateFundingPayVaultFee**](FundingApi.md#CreateFundingPayVaultFee) | **Post** /v1/funding/pay/vault-fees | Ship a stored Vault unit home and pay its fee (shipping, storage owed, early withdrawal) with restricted funding CBX instead of your card.
[**CreateFundingPoolStake**](FundingApi.md#CreateFundingPoolStake) | **Post** /v1/funding/pool/stake | Stake earned CBX into the funding pool. Refused, with the reason, until the pool is live and staking is open.
[**CreateFundingPoolUnstake**](FundingApi.md#CreateFundingPoolUnstake) | **Post** /v1/funding/pool/unstake | Queue an unstake. Paid first-in first-out once your lockup has passed.
[**CreateFundingPoolUnstakeCancel**](FundingApi.md#CreateFundingPoolUnstakeCancel) | **Post** /v1/funding/pool/unstake/{id}/cancel | Cancel a queued unstake that has not been paid yet.
[**CreateFundingPoolWaitlist**](FundingApi.md#CreateFundingPoolWaitlist) | **Post** /v1/funding/pool/waitlist | Join the staker waitlist: told the moment staking into the funding pool opens.
[**CreateFundingRepay**](FundingApi.md#CreateFundingRepay) | **Post** /v1/funding/repay | Repay. source \&quot;balance\&quot; (default) takes it from your available Stripe payout balance and applies it at once. source \&quot;card\&quot; returns a PaymentIntent client secret to confirm with Stripe; source \&quot;checkout\&quot; returns checkoutUrl, a hosted Stripe Checkout page the seller pays on (no Stripe SDK needed). source \&quot;cbx\&quot; pays early from your EARNED CBX at the obligation rate (granted and funding CBX never repay a line) and applies at once. Card and checkout apply when the payment succeeds. Never more than you owe; idempotent on idempotencyKey.
[**DeleteFundingMember**](FundingApi.md#DeleteFundingMember) | **Delete** /v1/funding/members/{memberUserId} | Remove a team member’s funding limit: they can no longer draw.
[**GetFundingControl**](FundingApi.md#GetFundingControl) | **Get** /v1/funding/controls | Your funding spend controls: approved uses, supplier allow/block lists, the approval threshold, the two-person rule and each team member’s funding limit.
[**GetFundingDraw**](FundingApi.md#GetFundingDraw) | **Get** /v1/funding/draws/{id} | One draw: amount, use, supplier, stage and its timeline (requested → approved → disbursed → spent → repaid).
[**GetFundingDrawInvoice**](FundingApi.md#GetFundingDrawInvoice) | **Get** /v1/funding/draws/{id}/invoice | A 15-minute signed link to a draw&#39;s uploaded supplier invoice, with the file&#39;s name, type and size. 404 when the draw has no uploaded file.
[**GetFundingLine**](FundingApi.md#GetFundingLine) | **Get** /v1/funding/line | Your line: limit, available, outstanding (principal + interest), APR, interest accrued since the last posting, the payoff figure, next due date, anything overdue, the per-draw schedule and your auto-repay share.
[**GetFundingOffer**](FundingApi.md#GetFundingOffer) | **Get** /v1/funding/offer | Your pre-qualified Crossly Funding limit and the eligibility checklist: each rule, whether you pass it, and what to do if not. Works with or without a line.
[**GetFundingPool**](FundingApi.md#GetFundingPool) | **Get** /v1/funding/pool | The Crossly Funding pool: live or not and why (no token yet, no pool, no price, no capacity), capacity, utilisation, the reserve, and the uses and suppliers it allows.
[**GetFundingPoolStake**](FundingApi.md#GetFundingPoolStake) | **Get** /v1/funding/pool/stake | Your position in the funding pool, the trailing APY reading (never a promise), lockup, unstake queue, history and the risk disclosure.
[**GetFundingSpendable**](FundingApi.md#GetFundingSpendable) | **Get** /v1/funding/spendable | What your restricted funding CBX can pay right now, in cents, for each Crossly use (shipping labels, Crossly fees, ads, Vault fees). 0 &#x3D; none, or your spend controls refuse that use.
[**GetFundingStatement**](FundingApi.md#GetFundingStatement) | **Get** /v1/funding/statements | A monthly statement (?month&#x3D;YYYY-MM, default the latest): opening balance, every draw, interest posting and payment, and the closing balance; plus the months available.
[**ListFundingApprovals**](FundingApi.md#ListFundingApprovals) | **Get** /v1/funding/approvals | Draws your spend controls hold for approval, with the approvals so far and how many are needed.
[**ListFundingDraws**](FundingApi.md#ListFundingDraws) | **Get** /v1/funding/draws | Every draw on your line: what it paid for (a supplier order, or a Crossly checkout), its APR, estimated fee and due date.
[**ListFundingReceipts**](FundingApi.md#ListFundingReceipts) | **Get** /v1/funding/receipts | Your funding payments to Crossly, newest first: the use, what it paid for, cents, status (pending, paid, failed) and the draw it was stamped against.
[**ListFundingRepayments**](FundingApi.md#ListFundingRepayments) | **Get** /v1/funding/repayments | Manual repayments (card / bank, or from your payout balance), newest first.
[**UpdateFundingAutorepay**](FundingApi.md#UpdateFundingAutorepay) | **Put** /v1/funding/autorepay | Set the share of each Crossly payout applied to what you owe: a whole percent from 10 to 100 (default 20). Never more than you owe is taken.
[**UpdateFundingControl**](FundingApi.md#UpdateFundingControl) | **Put** /v1/funding/controls | Change your funding spend controls. Enforced at every spend: wholesale checkout, labels, ads, Crossly and Vault fees.
[**UpdateFundingMember**](FundingApi.md#UpdateFundingMember) | **Put** /v1/funding/members/{memberUserId} | Set a team member’s funding limit (per draw or per month).



## CreateFundingApply

> CreateFundingApplyResponse CreateFundingApply(ctx).Execute()

Apply for Crossly Funding. Decided at once on the checklist rules: approved opens your line (limit = 25% of 90 days of settled Crossly payouts), declined lists why, needs_review goes to a person.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingApply(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingApply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingApply`: CreateFundingApplyResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingApply`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingApplyRequest struct via the builder pattern


### Return type

**CreateFundingApplyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFundingDraw

> CreateFundingDrawResponse CreateFundingDraw(ctx).Execute()

Draw on your line for a supplier order. Send the supplier, the invoice (an uploaded invoiceFileId from POST /v1/funding/invoice-files, a reference, or a URL) and the APR from the quote you accepted (acceptedAprBps); refused if your rate changed. The advance is held as in-platform Crossly balance, never cash. Idempotent on idempotencyKey.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingDraw(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingDraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingDraw`: CreateFundingDrawResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingDraw`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingDrawRequest struct via the builder pattern


### Return type

**CreateFundingDrawResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFundingDrawApprove

> CreateFundingDrawRejectResponse CreateFundingDrawApprove(ctx, id).Execute()

Approve a held draw. The two-person rule needs two distinct approvers; the requester never approves their own.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingDrawApprove(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingDrawApprove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingDrawApprove`: CreateFundingDrawRejectResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingDrawApprove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingDrawApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateFundingDrawRejectResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFundingDrawQuote

> CreateFundingDrawQuoteResponse CreateFundingDrawQuote(ctx).Execute()

The disclosure for a draw of amountCents: APR, the simple-interest fee for the full term, the total and the due date, plus your available headroom. Nothing is drawn.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingDrawQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingDrawQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingDrawQuote`: CreateFundingDrawQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingDrawQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingDrawQuoteRequest struct via the builder pattern


### Return type

**CreateFundingDrawQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFundingDrawReject

> CreateFundingDrawRejectResponse CreateFundingDrawReject(ctx, id).Execute()

Reject a held draw. Nothing was drawn, so nothing is owed.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingDrawReject(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingDrawReject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingDrawReject`: CreateFundingDrawRejectResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingDrawReject`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingDrawRejectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateFundingDrawRejectResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFundingDrawUse

> CreateFundingDrawUseResponse CreateFundingDrawUse(ctx).Execute()

Request a draw for an approved use. Refused by the spend controls, held for approval (stage requested — no debt until disbursed), or disbursed at once as restricted funding CBX. A DEBT with interest: show the disclosure first.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingDrawUse(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingDrawUse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingDrawUse`: CreateFundingDrawUseResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingDrawUse`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingDrawUseRequest struct via the builder pattern


### Return type

**CreateFundingDrawUseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFundingInvoiceFile

> CreateFundingInvoiceFileResponse CreateFundingInvoiceFile(ctx).Execute()

Sign an upload for a supplier invoice: a PDF or image (application/pdf, image/jpeg, image/png, image/webp, image/heic) of at most 15 MB. PUT the bytes to uploadUrl with the same Content-Type and exactly contentLength bytes, then draw with invoiceFileId = file.id. Stored privately; one draw per file.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingInvoiceFile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingInvoiceFile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingInvoiceFile`: CreateFundingInvoiceFileResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingInvoiceFile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingInvoiceFileRequest struct via the builder pattern


### Return type

**CreateFundingInvoiceFileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFundingPayCrosslyFee

> CreateFundingPayCrosslyFeeResponse CreateFundingPayCrosslyFee(ctx).Execute()

Pay an open invoice of your Crossly plan with restricted funding CBX (no card is charged; the invoice is marked paid). Idempotent per invoice.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingPayCrosslyFee(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingPayCrosslyFee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingPayCrosslyFee`: CreateFundingPayCrosslyFeeResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingPayCrosslyFee`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingPayCrosslyFeeRequest struct via the builder pattern


### Return type

**CreateFundingPayCrosslyFeeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFundingPayShippingLabel

> CreateFundingPayShippingLabelResponse CreateFundingPayShippingLabel(ctx).Execute()

Top up your shipping wallet with restricted funding CBX; labels you buy are then debited from it. Idempotent on idempotencyKey. Refused when the spend controls forbid shipping labels or it is over your approval threshold with no approved draw.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingPayShippingLabel(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingPayShippingLabel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingPayShippingLabel`: CreateFundingPayShippingLabelResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingPayShippingLabel`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingPayShippingLabelRequest struct via the builder pattern


### Return type

**CreateFundingPayShippingLabelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFundingPayVaultFee

> CreateFundingPayVaultFeeResponse CreateFundingPayVaultFee(ctx).Execute()

Ship a stored Vault unit home and pay its fee (shipping, storage owed, early withdrawal) with restricted funding CBX instead of your card.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingPayVaultFee(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingPayVaultFee``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingPayVaultFee`: CreateFundingPayVaultFeeResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingPayVaultFee`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingPayVaultFeeRequest struct via the builder pattern


### Return type

**CreateFundingPayVaultFeeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFundingPoolStake

> CreateFundingPoolStakeResponse CreateFundingPoolStake(ctx).Execute()

Stake earned CBX into the funding pool. Refused, with the reason, until the pool is live and staking is open.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingPoolStake(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingPoolStake``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingPoolStake`: CreateFundingPoolStakeResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingPoolStake`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingPoolStakeRequest struct via the builder pattern


### Return type

**CreateFundingPoolStakeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFundingPoolUnstake

> FundingPoolResponse CreateFundingPoolUnstake(ctx).Execute()

Queue an unstake. Paid first-in first-out once your lockup has passed.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingPoolUnstake(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingPoolUnstake``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingPoolUnstake`: FundingPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingPoolUnstake`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingPoolUnstakeRequest struct via the builder pattern


### Return type

**FundingPoolResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFundingPoolUnstakeCancel

> FundingPoolResponse CreateFundingPoolUnstakeCancel(ctx, id).Execute()

Cancel a queued unstake that has not been paid yet.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingPoolUnstakeCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingPoolUnstakeCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingPoolUnstakeCancel`: FundingPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingPoolUnstakeCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingPoolUnstakeCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**FundingPoolResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFundingPoolWaitlist

> FundingPoolResponse CreateFundingPoolWaitlist(ctx).Execute()

Join the staker waitlist: told the moment staking into the funding pool opens.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingPoolWaitlist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingPoolWaitlist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingPoolWaitlist`: FundingPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingPoolWaitlist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingPoolWaitlistRequest struct via the builder pattern


### Return type

**FundingPoolResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFundingRepay

> CreateFundingRepayResponse CreateFundingRepay(ctx).Execute()

Repay. source \"balance\" (default) takes it from your available Stripe payout balance and applies it at once. source \"card\" returns a PaymentIntent client secret to confirm with Stripe; source \"checkout\" returns checkoutUrl, a hosted Stripe Checkout page the seller pays on (no Stripe SDK needed). source \"cbx\" pays early from your EARNED CBX at the obligation rate (granted and funding CBX never repay a line) and applies at once. Card and checkout apply when the payment succeeds. Never more than you owe; idempotent on idempotencyKey.

### Example

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
    resp, r, err := apiClient.FundingApi.CreateFundingRepay(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.CreateFundingRepay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFundingRepay`: CreateFundingRepayResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.CreateFundingRepay`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFundingRepayRequest struct via the builder pattern


### Return type

**CreateFundingRepayResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteFundingMember

> DeleteFundingMemberResponse DeleteFundingMember(ctx, memberUserId).Execute()

Remove a team member’s funding limit: they can no longer draw.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberUserId := "memberUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundingApi.DeleteFundingMember(context.Background(), memberUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.DeleteFundingMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFundingMember`: DeleteFundingMemberResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.DeleteFundingMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFundingMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteFundingMemberResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetFundingControl

> GetFundingControlResponse GetFundingControl(ctx).Execute()

Your funding spend controls: approved uses, supplier allow/block lists, the approval threshold, the two-person rule and each team member’s funding limit.

### Example

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
    resp, r, err := apiClient.FundingApi.GetFundingControl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.GetFundingControl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingControl`: GetFundingControlResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.GetFundingControl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingControlRequest struct via the builder pattern


### Return type

**GetFundingControlResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetFundingDraw

> GetFundingDrawResponse GetFundingDraw(ctx, id).Execute()

One draw: amount, use, supplier, stage and its timeline (requested → approved → disbursed → spent → repaid).

### Example

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
    resp, r, err := apiClient.FundingApi.GetFundingDraw(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.GetFundingDraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingDraw`: GetFundingDrawResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.GetFundingDraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingDrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetFundingDrawResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetFundingDrawInvoice

> GetFundingDrawInvoiceResponse GetFundingDrawInvoice(ctx, id).Execute()

A 15-minute signed link to a draw's uploaded supplier invoice, with the file's name, type and size. 404 when the draw has no uploaded file.

### Example

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
    resp, r, err := apiClient.FundingApi.GetFundingDrawInvoice(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.GetFundingDrawInvoice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingDrawInvoice`: GetFundingDrawInvoiceResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.GetFundingDrawInvoice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingDrawInvoiceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetFundingDrawInvoiceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetFundingLine

> GetFundingLineResponse GetFundingLine(ctx).Execute()

Your line: limit, available, outstanding (principal + interest), APR, interest accrued since the last posting, the payoff figure, next due date, anything overdue, the per-draw schedule and your auto-repay share.

### Example

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
    resp, r, err := apiClient.FundingApi.GetFundingLine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.GetFundingLine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingLine`: GetFundingLineResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.GetFundingLine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingLineRequest struct via the builder pattern


### Return type

**GetFundingLineResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetFundingOffer

> GetFundingOfferResponse GetFundingOffer(ctx).Execute()

Your pre-qualified Crossly Funding limit and the eligibility checklist: each rule, whether you pass it, and what to do if not. Works with or without a line.

### Example

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
    resp, r, err := apiClient.FundingApi.GetFundingOffer(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.GetFundingOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingOffer`: GetFundingOfferResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.GetFundingOffer`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingOfferRequest struct via the builder pattern


### Return type

**GetFundingOfferResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetFundingPool

> GetFundingPoolResponse GetFundingPool(ctx).Execute()

The Crossly Funding pool: live or not and why (no token yet, no pool, no price, no capacity), capacity, utilisation, the reserve, and the uses and suppliers it allows.

### Example

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
    resp, r, err := apiClient.FundingApi.GetFundingPool(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.GetFundingPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingPool`: GetFundingPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.GetFundingPool`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingPoolRequest struct via the builder pattern


### Return type

**GetFundingPoolResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetFundingPoolStake

> FundingPoolResponse GetFundingPoolStake(ctx).Execute()

Your position in the funding pool, the trailing APY reading (never a promise), lockup, unstake queue, history and the risk disclosure.

### Example

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
    resp, r, err := apiClient.FundingApi.GetFundingPoolStake(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.GetFundingPoolStake``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingPoolStake`: FundingPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.GetFundingPoolStake`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingPoolStakeRequest struct via the builder pattern


### Return type

**FundingPoolResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetFundingSpendable

> GetFundingSpendableResponse GetFundingSpendable(ctx).Execute()

What your restricted funding CBX can pay right now, in cents, for each Crossly use (shipping labels, Crossly fees, ads, Vault fees). 0 = none, or your spend controls refuse that use.

### Example

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
    resp, r, err := apiClient.FundingApi.GetFundingSpendable(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.GetFundingSpendable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingSpendable`: GetFundingSpendableResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.GetFundingSpendable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingSpendableRequest struct via the builder pattern


### Return type

**GetFundingSpendableResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetFundingStatement

> GetFundingStatementResponse GetFundingStatement(ctx).Execute()

A monthly statement (?month=YYYY-MM, default the latest): opening balance, every draw, interest posting and payment, and the closing balance; plus the months available.

### Example

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
    resp, r, err := apiClient.FundingApi.GetFundingStatement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.GetFundingStatement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFundingStatement`: GetFundingStatementResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.GetFundingStatement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetFundingStatementRequest struct via the builder pattern


### Return type

**GetFundingStatementResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListFundingApprovals

> ListFundingApprovalsResponse ListFundingApprovals(ctx).Execute()

Draws your spend controls hold for approval, with the approvals so far and how many are needed.

### Example

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
    resp, r, err := apiClient.FundingApi.ListFundingApprovals(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.ListFundingApprovals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFundingApprovals`: ListFundingApprovalsResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.ListFundingApprovals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFundingApprovalsRequest struct via the builder pattern


### Return type

**ListFundingApprovalsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListFundingDraws

> ListFundingDrawsResponse ListFundingDraws(ctx).Execute()

Every draw on your line: what it paid for (a supplier order, or a Crossly checkout), its APR, estimated fee and due date.

### Example

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
    resp, r, err := apiClient.FundingApi.ListFundingDraws(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.ListFundingDraws``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFundingDraws`: ListFundingDrawsResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.ListFundingDraws`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFundingDrawsRequest struct via the builder pattern


### Return type

**ListFundingDrawsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListFundingReceipts

> ListFundingReceiptsResponse ListFundingReceipts(ctx).Execute()

Your funding payments to Crossly, newest first: the use, what it paid for, cents, status (pending, paid, failed) and the draw it was stamped against.

### Example

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
    resp, r, err := apiClient.FundingApi.ListFundingReceipts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.ListFundingReceipts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFundingReceipts`: ListFundingReceiptsResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.ListFundingReceipts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFundingReceiptsRequest struct via the builder pattern


### Return type

**ListFundingReceiptsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListFundingRepayments

> ListFundingRepaymentsResponse ListFundingRepayments(ctx).Execute()

Manual repayments (card / bank, or from your payout balance), newest first.

### Example

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
    resp, r, err := apiClient.FundingApi.ListFundingRepayments(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.ListFundingRepayments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFundingRepayments`: ListFundingRepaymentsResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.ListFundingRepayments`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFundingRepaymentsRequest struct via the builder pattern


### Return type

**ListFundingRepaymentsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateFundingAutorepay

> UpdateFundingAutorepayResponse UpdateFundingAutorepay(ctx).Execute()

Set the share of each Crossly payout applied to what you owe: a whole percent from 10 to 100 (default 20). Never more than you owe is taken.

### Example

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
    resp, r, err := apiClient.FundingApi.UpdateFundingAutorepay(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.UpdateFundingAutorepay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFundingAutorepay`: UpdateFundingAutorepayResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.UpdateFundingAutorepay`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFundingAutorepayRequest struct via the builder pattern


### Return type

**UpdateFundingAutorepayResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateFundingControl

> UpdateFundingControlResponse UpdateFundingControl(ctx).Execute()

Change your funding spend controls. Enforced at every spend: wholesale checkout, labels, ads, Crossly and Vault fees.

### Example

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
    resp, r, err := apiClient.FundingApi.UpdateFundingControl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.UpdateFundingControl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFundingControl`: UpdateFundingControlResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.UpdateFundingControl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFundingControlRequest struct via the builder pattern


### Return type

**UpdateFundingControlResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateFundingMember

> UpdateFundingMemberResponse UpdateFundingMember(ctx, memberUserId).Execute()

Set a team member’s funding limit (per draw or per month).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    memberUserId := "memberUserId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.FundingApi.UpdateFundingMember(context.Background(), memberUserId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FundingApi.UpdateFundingMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFundingMember`: UpdateFundingMemberResponse
    fmt.Fprintf(os.Stdout, "Response from `FundingApi.UpdateFundingMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberUserId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFundingMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateFundingMemberResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

