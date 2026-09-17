# \CBXApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCbxAccrual**](CBXApi.md#CreateCbxAccrual) | **Post** /v1/cbx/accruals | Record cashback a user earned, in cents.
[**CreateCbxAccrualPurchase**](CBXApi.md#CreateCbxAccrualPurchase) | **Post** /v1/cbx/accruals/purchase | Accrue cashback for an order at the resolved rate.
[**CreateCbxAccrualReverse**](CBXApi.md#CreateCbxAccrualReverse) | **Post** /v1/cbx/accruals/{accrualId}/reverse | Claw back a pending accrual — a refund, a cancellation, fraud.
[**CreateCbxAdCreditPurchase**](CBXApi.md#CreateCbxAdCreditPurchase) | **Post** /v1/cbx/ad-credit/purchase | Claim ad credit against a CBX transfer you sent.
[**CreateCbxAdCreditQuote**](CBXApi.md#CreateCbxAdCreditQuote) | **Post** /v1/cbx/ad-credit/quote | What a given number of tokens buys in ad credit.
[**CreateCbxAdCreditSpend**](CBXApi.md#CreateCbxAdCreditSpend) | **Post** /v1/cbx/ad-credit/spend | Consume credit for a billing period.
[**CreateCbxBoost**](CBXApi.md#CreateCbxBoost) | **Post** /v1/cbx/boosts | Fund elevated cashback on matching items.
[**CreateCbxBoostPause**](CBXApi.md#CreateCbxBoostPause) | **Post** /v1/cbx/boosts/{boostId}/pause | Stop a boost from matching further orders.
[**CreateCbxCampaign**](CBXApi.md#CreateCbxCampaign) | **Post** /v1/cbx/campaigns | Create a campaign in draft.
[**CreateCbxCampaignApprove**](CBXApi.md#CreateCbxCampaignApprove) | **Post** /v1/cbx/campaigns/{campaignId}/approve | Approve the previewed recipient list.
[**CreateCbxCampaignExecute**](CBXApi.md#CreateCbxCampaignExecute) | **Post** /v1/cbx/campaigns/{campaignId}/execute | Pay an approved campaign.
[**CreateCbxCampaignPreview**](CBXApi.md#CreateCbxCampaignPreview) | **Post** /v1/cbx/campaigns/{campaignId}/preview | Compute the recipient list without paying it.
[**CreateCbxClaim**](CBXApi.md#CreateCbxClaim) | **Post** /v1/cbx/claims | Reserve a claim. Debits the balance and queues the transfer.
[**CreateCbxClaimQuote**](CBXApi.md#CreateCbxClaimQuote) | **Post** /v1/cbx/claims/quote | What a claim would cost, without committing to it.
[**CreateCbxClaimSend**](CBXApi.md#CreateCbxClaimSend) | **Post** /v1/cbx/claims/{claimId}/send | Send a reserved claim on chain.
[**CreateCbxCreditDraw**](CBXApi.md#CreateCbxCreditDraw) | **Post** /v1/cbx/credit/draw | Draw against a line, receiving grant credit.
[**CreateCbxCreditFreeze**](CBXApi.md#CreateCbxCreditFreeze) | **Post** /v1/cbx/credit/freeze | Stop new draws. Leaves the drawn balance on its terms.
[**CreateCbxCreditRefresh**](CBXApi.md#CreateCbxCreditRefresh) | **Post** /v1/cbx/credit/refresh | Recompute a limit from trading history and stake.
[**CreateCbxCreditRepay**](CBXApi.md#CreateCbxCreditRepay) | **Post** /v1/cbx/credit/repay | Apply a repayment to a line.
[**CreateCbxDisbursementRule**](CBXApi.md#CreateCbxDisbursementRule) | **Post** /v1/cbx/disbursement-rules | Fire a distribution when the events pool crosses a threshold.
[**CreateCbxDisbursementRuleActive**](CBXApi.md#CreateCbxDisbursementRuleActive) | **Post** /v1/cbx/disbursement-rules/{ruleId}/active | Enable or disable a rule.
[**CreateCbxDisbursementRuleCheck**](CBXApi.md#CreateCbxDisbursementRuleCheck) | **Post** /v1/cbx/disbursement-rules/{ruleId}/check | Evaluate a rule now. Fires it if every gate passes.
[**CreateCbxEarnTier**](CBXApi.md#CreateCbxEarnTier) | **Post** /v1/cbx/earn-tiers | Define an earn term.
[**CreateCbxRateQuote**](CBXApi.md#CreateCbxRateQuote) | **Post** /v1/cbx/rates/quote | What would this order earn, and why.
[**CreateCbxRedemption**](CBXApi.md#CreateCbxRedemption) | **Post** /v1/cbx/redemptions | Pay for a service in CBX.
[**CreateCbxRedemptionQuote**](CBXApi.md#CreateCbxRedemptionQuote) | **Post** /v1/cbx/redemptions/quote | What a service costs in tokens right now.
[**CreateCbxRevenueSweep**](CBXApi.md#CreateCbxRevenueSweep) | **Post** /v1/cbx/revenue/sweep | Move accrued revenue from the reserve to your revenue wallet.
[**CreateCbxSpend**](CBXApi.md#CreateCbxSpend) | **Post** /v1/cbx/spends | Redeem a user&#39;s CBX against an order.
[**CreateCbxSpendReverse**](CBXApi.md#CreateCbxSpendReverse) | **Post** /v1/cbx/spends/{externalId}/reverse | Refund a spend — give the tokens back and claw the skim back.
[**CreateCbxStakeTier**](CBXApi.md#CreateCbxStakeTier) | **Post** /v1/cbx/stake-tiers | Define a staking tier.
[**CreateCbxSubject**](CBXApi.md#CreateCbxSubject) | **Post** /v1/cbx/subjects | Map one of your user ids to a CBX subject.
[**CreateCbxSubjectGrant**](CBXApi.md#CreateCbxSubjectGrant) | **Post** /v1/cbx/subjects/{subjectId}/grants | Issue grant credit — in-platform, non-withdrawable.
[**CreateCbxSubjectSpendPlan**](CBXApi.md#CreateCbxSubjectSpendPlan) | **Post** /v1/cbx/subjects/{subjectId}/spend-plan | Which balances would pay for a spend, and in what order.
[**CreateCbxSubjectStake**](CBXApi.md#CreateCbxSubjectStake) | **Post** /v1/cbx/subjects/{subjectId}/stake | Lock a subject&#39;s tokens for a tier.
[**CreateCbxSubjectStakeUnstake**](CBXApi.md#CreateCbxSubjectStakeUnstake) | **Post** /v1/cbx/subjects/{subjectId}/stake/unstake | Start the cooldown. Tokens unlock when it elapses.
[**CreateCbxWalletChallenge**](CBXApi.md#CreateCbxWalletChallenge) | **Post** /v1/cbx/wallets/challenge | Start wallet verification. Returns a message for the user to sign.
[**CreateCbxWalletPaymentConfirm**](CBXApi.md#CreateCbxWalletPaymentConfirm) | **Post** /v1/cbx/wallet-payments/confirm | Present the signature. Returns a ship / do-not-ship decision.
[**CreateCbxWalletPaymentQuote**](CBXApi.md#CreateCbxWalletPaymentQuote) | **Post** /v1/cbx/wallet-payments/quote | Build a transfer for the buyer to sign themselves.
[**CreateCbxWalletPaymentResolve**](CBXApi.md#CreateCbxWalletPaymentResolve) | **Post** /v1/cbx/wallet-payments/{paymentId}/resolve | A human decides on a held payment.
[**CreateCbxWalletVerify**](CBXApi.md#CreateCbxWalletVerify) | **Post** /v1/cbx/wallets/verify | Complete wallet verification with the user&#39;s signature.
[**GetCbxAdCredit**](CBXApi.md#GetCbxAdCredit) | **Get** /v1/cbx/ad-credit | Unspent advertising credit, in cents.
[**GetCbxClaim**](CBXApi.md#GetCbxClaim) | **Get** /v1/cbx/claims/{claimId} | A claim&#39;s current state.
[**GetCbxCredit**](CBXApi.md#GetCbxCredit) | **Get** /v1/cbx/credit | A seller&#39;s wholesale credit line.
[**GetCbxMe**](CBXApi.md#GetCbxMe) | **Get** /v1/cbx/me | Identity check — which merchant this key belongs to, and its terms.
[**GetCbxPool**](CBXApi.md#GetCbxPool) | **Get** /v1/cbx/pool | Your events-pool balance.
[**GetCbxRevenue**](CBXApi.md#GetCbxRevenue) | **Get** /v1/cbx/revenue | Operator revenue accrued and not yet withdrawn.
[**GetCbxSubjectBalance**](CBXApi.md#GetCbxSubjectBalance) | **Get** /v1/cbx/subjects/{subjectId}/balance | What a subject holds: pending cents and available CBX.
[**GetCbxSubjectBalanceBySubjectId**](CBXApi.md#GetCbxSubjectBalanceBySubjectId) | **Get** /v1/cbx/subjects/{subjectId}/balances | All three balances a subject holds.
[**GetCbxSubjectSpent**](CBXApi.md#GetCbxSubjectSpent) | **Get** /v1/cbx/subjects/{subjectId}/spent | Total CBX a subject has spent in your marketplace.
[**GetCbxSubjectStake**](CBXApi.md#GetCbxSubjectStake) | **Get** /v1/cbx/subjects/{subjectId}/stake | A subject&#39;s staking state and spendable balance.
[**GetCbxSubjectWallet**](CBXApi.md#GetCbxSubjectWallet) | **Get** /v1/cbx/subjects/{subjectId}/wallet | The verified payout address for a subject, if any.
[**GetCbxTreasury**](CBXApi.md#GetCbxTreasury) | **Get** /v1/cbx/treasury | Your most recent reserve reconciliation.
[**ListCbxAdCreditLedger**](CBXApi.md#ListCbxAdCreditLedger) | **Get** /v1/cbx/ad-credit/ledger | Ad-credit movements, newest first.
[**ListCbxBoosts**](CBXApi.md#ListCbxBoosts) | **Get** /v1/cbx/boosts | Your funded cashback boosts, newest first.
[**ListCbxCampaignPayouts**](CBXApi.md#ListCbxCampaignPayouts) | **Get** /v1/cbx/campaigns/{campaignId}/payouts | What a campaign actually paid, with the weight behind each amount.
[**ListCbxCampaigns**](CBXApi.md#ListCbxCampaigns) | **Get** /v1/cbx/campaigns | Your campaigns, newest first.
[**ListCbxDisbursementProgress**](CBXApi.md#ListCbxDisbursementProgress) | **Get** /v1/cbx/disbursement-progress | How close each rule is to firing — the public counter.
[**ListCbxDisbursementRules**](CBXApi.md#ListCbxDisbursementRules) | **Get** /v1/cbx/disbursement-rules | Threshold rules that fire community distributions.
[**ListCbxEarnTiers**](CBXApi.md#ListCbxEarnTiers) | **Get** /v1/cbx/earn-tiers | Earn terms on offer — longer maturation, higher rate.
[**ListCbxRedemptionServices**](CBXApi.md#ListCbxRedemptionServices) | **Get** /v1/cbx/redemptions/services | Services payable in CBX, and the discount each carries.
[**ListCbxStakeTiers**](CBXApi.md#ListCbxStakeTiers) | **Get** /v1/cbx/stake-tiers | Staking tiers — what locking tokens buys.
[**ListCbxSubjectGrants**](CBXApi.md#ListCbxSubjectGrants) | **Get** /v1/cbx/subjects/{subjectId}/grants | Live grants, soonest-expiring first.
[**ListCbxSubjectLedger**](CBXApi.md#ListCbxSubjectLedger) | **Get** /v1/cbx/subjects/{subjectId}/ledger | A subject&#39;s CBX ledger, newest first.
[**ListCbxWalletPaymentReview**](CBXApi.md#ListCbxWalletPaymentReview) | **Get** /v1/cbx/wallet-payments/review | Payments held for a human — the ops queue.



## CreateCbxAccrual

> CreateCbxAccrualResponse CreateCbxAccrual(ctx).Execute()

Record cashback a user earned, in cents.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxAccrual(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxAccrual``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxAccrual`: CreateCbxAccrualResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxAccrual`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxAccrualRequest struct via the builder pattern


### Return type

[**CreateCbxAccrualResponse**](CreateCbxAccrualResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxAccrualPurchase

> CreateCbxAccrualPurchaseResponse CreateCbxAccrualPurchase(ctx).Execute()

Accrue cashback for an order at the resolved rate.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxAccrualPurchase(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxAccrualPurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxAccrualPurchase`: CreateCbxAccrualPurchaseResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxAccrualPurchase`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxAccrualPurchaseRequest struct via the builder pattern


### Return type

[**CreateCbxAccrualPurchaseResponse**](CreateCbxAccrualPurchaseResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxAccrualReverse

> CreateCbxAccrualReverseResponse CreateCbxAccrualReverse(ctx, accrualId).Execute()

Claw back a pending accrual — a refund, a cancellation, fraud.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accrualId := "accrualId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.CreateCbxAccrualReverse(context.Background(), accrualId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxAccrualReverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxAccrualReverse`: CreateCbxAccrualReverseResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxAccrualReverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accrualId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxAccrualReverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCbxAccrualReverseResponse**](CreateCbxAccrualReverseResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxAdCreditPurchase

> CreateCbxAdCreditPurchaseResponse CreateCbxAdCreditPurchase(ctx).Execute()

Claim ad credit against a CBX transfer you sent.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxAdCreditPurchase(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxAdCreditPurchase``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxAdCreditPurchase`: CreateCbxAdCreditPurchaseResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxAdCreditPurchase`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxAdCreditPurchaseRequest struct via the builder pattern


### Return type

[**CreateCbxAdCreditPurchaseResponse**](CreateCbxAdCreditPurchaseResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxAdCreditQuote

> CreateCbxAdCreditQuoteResponse CreateCbxAdCreditQuote(ctx).Execute()

What a given number of tokens buys in ad credit.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxAdCreditQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxAdCreditQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxAdCreditQuote`: CreateCbxAdCreditQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxAdCreditQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxAdCreditQuoteRequest struct via the builder pattern


### Return type

[**CreateCbxAdCreditQuoteResponse**](CreateCbxAdCreditQuoteResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxAdCreditSpend

> CreateCbxAdCreditSpendResponse CreateCbxAdCreditSpend(ctx).Execute()

Consume credit for a billing period.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxAdCreditSpend(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxAdCreditSpend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxAdCreditSpend`: CreateCbxAdCreditSpendResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxAdCreditSpend`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxAdCreditSpendRequest struct via the builder pattern


### Return type

[**CreateCbxAdCreditSpendResponse**](CreateCbxAdCreditSpendResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxBoost

> CreateCbxBoostResponse CreateCbxBoost(ctx).Execute()

Fund elevated cashback on matching items.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxBoost(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxBoost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxBoost`: CreateCbxBoostResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxBoost`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxBoostRequest struct via the builder pattern


### Return type

[**CreateCbxBoostResponse**](CreateCbxBoostResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxBoostPause

> CreateCbxBoostPauseResponse CreateCbxBoostPause(ctx, boostId).Execute()

Stop a boost from matching further orders.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    boostId := "boostId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.CreateCbxBoostPause(context.Background(), boostId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxBoostPause``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxBoostPause`: CreateCbxBoostPauseResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxBoostPause`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**boostId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxBoostPauseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCbxBoostPauseResponse**](CreateCbxBoostPauseResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxCampaign

> CreateCbxCampaignResponse CreateCbxCampaign(ctx).Execute()

Create a campaign in draft.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxCampaign(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxCampaign`: CreateCbxCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxCampaign`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxCampaignRequest struct via the builder pattern


### Return type

[**CreateCbxCampaignResponse**](CreateCbxCampaignResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxCampaignApprove

> CreateCbxCampaignApproveResponse CreateCbxCampaignApprove(ctx, campaignId).Execute()

Approve the previewed recipient list.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    campaignId := "campaignId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.CreateCbxCampaignApprove(context.Background(), campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxCampaignApprove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxCampaignApprove`: CreateCbxCampaignApproveResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxCampaignApprove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxCampaignApproveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCbxCampaignApproveResponse**](CreateCbxCampaignApproveResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxCampaignExecute

> CreateCbxCampaignExecuteResponse CreateCbxCampaignExecute(ctx, campaignId).Execute()

Pay an approved campaign.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    campaignId := "campaignId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.CreateCbxCampaignExecute(context.Background(), campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxCampaignExecute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxCampaignExecute`: CreateCbxCampaignExecuteResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxCampaignExecute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxCampaignExecuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCbxCampaignExecuteResponse**](CreateCbxCampaignExecuteResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxCampaignPreview

> CreateCbxCampaignPreviewResponse CreateCbxCampaignPreview(ctx, campaignId).Execute()

Compute the recipient list without paying it.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    campaignId := "campaignId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.CreateCbxCampaignPreview(context.Background(), campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxCampaignPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxCampaignPreview`: CreateCbxCampaignPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxCampaignPreview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxCampaignPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCbxCampaignPreviewResponse**](CreateCbxCampaignPreviewResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxClaim

> CreateCbxClaimResponse CreateCbxClaim(ctx).Execute()

Reserve a claim. Debits the balance and queues the transfer.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxClaim(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxClaim`: CreateCbxClaimResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxClaim`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxClaimRequest struct via the builder pattern


### Return type

[**CreateCbxClaimResponse**](CreateCbxClaimResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxClaimQuote

> CreateCbxClaimQuoteResponse CreateCbxClaimQuote(ctx).Execute()

What a claim would cost, without committing to it.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxClaimQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxClaimQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxClaimQuote`: CreateCbxClaimQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxClaimQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxClaimQuoteRequest struct via the builder pattern


### Return type

[**CreateCbxClaimQuoteResponse**](CreateCbxClaimQuoteResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxClaimSend

> CreateCbxClaimSendResponse CreateCbxClaimSend(ctx, claimId).Execute()

Send a reserved claim on chain.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    claimId := "claimId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.CreateCbxClaimSend(context.Background(), claimId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxClaimSend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxClaimSend`: CreateCbxClaimSendResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxClaimSend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**claimId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxClaimSendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCbxClaimSendResponse**](CreateCbxClaimSendResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxCreditDraw

> CreateCbxCreditDrawResponse CreateCbxCreditDraw(ctx).Execute()

Draw against a line, receiving grant credit.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxCreditDraw(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxCreditDraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxCreditDraw`: CreateCbxCreditDrawResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxCreditDraw`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxCreditDrawRequest struct via the builder pattern


### Return type

[**CreateCbxCreditDrawResponse**](CreateCbxCreditDrawResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxCreditFreeze

> CreateCbxCreditFreezeResponse CreateCbxCreditFreeze(ctx).Execute()

Stop new draws. Leaves the drawn balance on its terms.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxCreditFreeze(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxCreditFreeze``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxCreditFreeze`: CreateCbxCreditFreezeResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxCreditFreeze`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxCreditFreezeRequest struct via the builder pattern


### Return type

[**CreateCbxCreditFreezeResponse**](CreateCbxCreditFreezeResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxCreditRefresh

> CreateCbxCreditRefreshResponse CreateCbxCreditRefresh(ctx).Execute()

Recompute a limit from trading history and stake.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxCreditRefresh(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxCreditRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxCreditRefresh`: CreateCbxCreditRefreshResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxCreditRefresh`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxCreditRefreshRequest struct via the builder pattern


### Return type

[**CreateCbxCreditRefreshResponse**](CreateCbxCreditRefreshResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxCreditRepay

> CreateCbxCreditRepayResponse CreateCbxCreditRepay(ctx).Execute()

Apply a repayment to a line.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxCreditRepay(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxCreditRepay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxCreditRepay`: CreateCbxCreditRepayResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxCreditRepay`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxCreditRepayRequest struct via the builder pattern


### Return type

[**CreateCbxCreditRepayResponse**](CreateCbxCreditRepayResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxDisbursementRule

> CreateCbxDisbursementRuleResponse CreateCbxDisbursementRule(ctx).Execute()

Fire a distribution when the events pool crosses a threshold.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxDisbursementRule(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxDisbursementRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxDisbursementRule`: CreateCbxDisbursementRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxDisbursementRule`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxDisbursementRuleRequest struct via the builder pattern


### Return type

[**CreateCbxDisbursementRuleResponse**](CreateCbxDisbursementRuleResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxDisbursementRuleActive

> CreateCbxDisbursementRuleActiveResponse CreateCbxDisbursementRuleActive(ctx, ruleId).Execute()

Enable or disable a rule.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.CreateCbxDisbursementRuleActive(context.Background(), ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxDisbursementRuleActive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxDisbursementRuleActive`: CreateCbxDisbursementRuleActiveResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxDisbursementRuleActive`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxDisbursementRuleActiveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCbxDisbursementRuleActiveResponse**](CreateCbxDisbursementRuleActiveResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxDisbursementRuleCheck

> CreateCbxDisbursementRuleCheckResponse CreateCbxDisbursementRuleCheck(ctx, ruleId).Execute()

Evaluate a rule now. Fires it if every gate passes.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ruleId := "ruleId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.CreateCbxDisbursementRuleCheck(context.Background(), ruleId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxDisbursementRuleCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxDisbursementRuleCheck`: CreateCbxDisbursementRuleCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxDisbursementRuleCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ruleId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxDisbursementRuleCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCbxDisbursementRuleCheckResponse**](CreateCbxDisbursementRuleCheckResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxEarnTier

> CreateCbxEarnTierResponse CreateCbxEarnTier(ctx).Execute()

Define an earn term.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxEarnTier(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxEarnTier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxEarnTier`: CreateCbxEarnTierResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxEarnTier`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxEarnTierRequest struct via the builder pattern


### Return type

[**CreateCbxEarnTierResponse**](CreateCbxEarnTierResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxRateQuote

> CreateCbxRateQuoteResponse CreateCbxRateQuote(ctx).Execute()

What would this order earn, and why.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxRateQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxRateQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxRateQuote`: CreateCbxRateQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxRateQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxRateQuoteRequest struct via the builder pattern


### Return type

[**CreateCbxRateQuoteResponse**](CreateCbxRateQuoteResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxRedemption

> CreateCbxRedemptionResponse CreateCbxRedemption(ctx).Execute()

Pay for a service in CBX.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxRedemption(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxRedemption``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxRedemption`: CreateCbxRedemptionResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxRedemption`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxRedemptionRequest struct via the builder pattern


### Return type

[**CreateCbxRedemptionResponse**](CreateCbxRedemptionResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxRedemptionQuote

> CreateCbxRedemptionQuoteResponse CreateCbxRedemptionQuote(ctx).Execute()

What a service costs in tokens right now.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxRedemptionQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxRedemptionQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxRedemptionQuote`: CreateCbxRedemptionQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxRedemptionQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxRedemptionQuoteRequest struct via the builder pattern


### Return type

[**CreateCbxRedemptionQuoteResponse**](CreateCbxRedemptionQuoteResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxRevenueSweep

> CreateCbxRevenueSweepResponse CreateCbxRevenueSweep(ctx).Execute()

Move accrued revenue from the reserve to your revenue wallet.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxRevenueSweep(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxRevenueSweep``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxRevenueSweep`: CreateCbxRevenueSweepResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxRevenueSweep`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxRevenueSweepRequest struct via the builder pattern


### Return type

[**CreateCbxRevenueSweepResponse**](CreateCbxRevenueSweepResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxSpend

> CreateCbxSpendResponse CreateCbxSpend(ctx).Execute()

Redeem a user's CBX against an order.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxSpend(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxSpend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxSpend`: CreateCbxSpendResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxSpend`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxSpendRequest struct via the builder pattern


### Return type

[**CreateCbxSpendResponse**](CreateCbxSpendResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxSpendReverse

> CreateCbxSpendReverseResponse CreateCbxSpendReverse(ctx, externalId).Execute()

Refund a spend — give the tokens back and claw the skim back.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    externalId := "externalId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.CreateCbxSpendReverse(context.Background(), externalId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxSpendReverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxSpendReverse`: CreateCbxSpendReverseResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxSpendReverse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxSpendReverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCbxSpendReverseResponse**](CreateCbxSpendReverseResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxStakeTier

> CreateCbxStakeTierResponse CreateCbxStakeTier(ctx).Execute()

Define a staking tier.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxStakeTier(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxStakeTier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxStakeTier`: CreateCbxStakeTierResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxStakeTier`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxStakeTierRequest struct via the builder pattern


### Return type

[**CreateCbxStakeTierResponse**](CreateCbxStakeTierResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxSubject

> CreateCbxSubjectResponse CreateCbxSubject(ctx).Execute()

Map one of your user ids to a CBX subject.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxSubject(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxSubject``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxSubject`: CreateCbxSubjectResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxSubject`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxSubjectRequest struct via the builder pattern


### Return type

[**CreateCbxSubjectResponse**](CreateCbxSubjectResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxSubjectGrant

> CreateCbxSubjectGrantResponse CreateCbxSubjectGrant(ctx, subjectId).Execute()

Issue grant credit — in-platform, non-withdrawable.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subjectId := "subjectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.CreateCbxSubjectGrant(context.Background(), subjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxSubjectGrant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxSubjectGrant`: CreateCbxSubjectGrantResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxSubjectGrant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxSubjectGrantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCbxSubjectGrantResponse**](CreateCbxSubjectGrantResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxSubjectSpendPlan

> CreateCbxSubjectSpendPlanResponse CreateCbxSubjectSpendPlan(ctx, subjectId).Execute()

Which balances would pay for a spend, and in what order.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subjectId := "subjectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.CreateCbxSubjectSpendPlan(context.Background(), subjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxSubjectSpendPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxSubjectSpendPlan`: CreateCbxSubjectSpendPlanResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxSubjectSpendPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxSubjectSpendPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCbxSubjectSpendPlanResponse**](CreateCbxSubjectSpendPlanResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxSubjectStake

> CreateCbxSubjectStakeResponse CreateCbxSubjectStake(ctx, subjectId).Execute()

Lock a subject's tokens for a tier.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subjectId := "subjectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.CreateCbxSubjectStake(context.Background(), subjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxSubjectStake``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxSubjectStake`: CreateCbxSubjectStakeResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxSubjectStake`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxSubjectStakeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCbxSubjectStakeResponse**](CreateCbxSubjectStakeResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxSubjectStakeUnstake

> CreateCbxSubjectStakeUnstakeResponse CreateCbxSubjectStakeUnstake(ctx, subjectId).Execute()

Start the cooldown. Tokens unlock when it elapses.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subjectId := "subjectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.CreateCbxSubjectStakeUnstake(context.Background(), subjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxSubjectStakeUnstake``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxSubjectStakeUnstake`: CreateCbxSubjectStakeUnstakeResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxSubjectStakeUnstake`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxSubjectStakeUnstakeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCbxSubjectStakeUnstakeResponse**](CreateCbxSubjectStakeUnstakeResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxWalletChallenge

> CreateCbxWalletChallengeResponse CreateCbxWalletChallenge(ctx).Execute()

Start wallet verification. Returns a message for the user to sign.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxWalletChallenge(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxWalletChallenge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxWalletChallenge`: CreateCbxWalletChallengeResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxWalletChallenge`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxWalletChallengeRequest struct via the builder pattern


### Return type

[**CreateCbxWalletChallengeResponse**](CreateCbxWalletChallengeResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxWalletPaymentConfirm

> CreateCbxWalletPaymentConfirmResponse CreateCbxWalletPaymentConfirm(ctx).Execute()

Present the signature. Returns a ship / do-not-ship decision.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxWalletPaymentConfirm(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxWalletPaymentConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxWalletPaymentConfirm`: CreateCbxWalletPaymentConfirmResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxWalletPaymentConfirm`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxWalletPaymentConfirmRequest struct via the builder pattern


### Return type

[**CreateCbxWalletPaymentConfirmResponse**](CreateCbxWalletPaymentConfirmResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxWalletPaymentQuote

> CreateCbxWalletPaymentQuoteResponse CreateCbxWalletPaymentQuote(ctx).Execute()

Build a transfer for the buyer to sign themselves.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxWalletPaymentQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxWalletPaymentQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxWalletPaymentQuote`: CreateCbxWalletPaymentQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxWalletPaymentQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxWalletPaymentQuoteRequest struct via the builder pattern


### Return type

[**CreateCbxWalletPaymentQuoteResponse**](CreateCbxWalletPaymentQuoteResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxWalletPaymentResolve

> CreateCbxWalletPaymentResolveResponse CreateCbxWalletPaymentResolve(ctx, paymentId).Execute()

A human decides on a held payment.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentId := "paymentId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.CreateCbxWalletPaymentResolve(context.Background(), paymentId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxWalletPaymentResolve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxWalletPaymentResolve`: CreateCbxWalletPaymentResolveResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxWalletPaymentResolve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxWalletPaymentResolveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateCbxWalletPaymentResolveResponse**](CreateCbxWalletPaymentResolveResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCbxWalletVerify

> CreateCbxWalletVerifyResponse CreateCbxWalletVerify(ctx).Execute()

Complete wallet verification with the user's signature.



### Example

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
    resp, r, err := apiClient.CBXApi.CreateCbxWalletVerify(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.CreateCbxWalletVerify``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCbxWalletVerify`: CreateCbxWalletVerifyResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.CreateCbxWalletVerify`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCbxWalletVerifyRequest struct via the builder pattern


### Return type

[**CreateCbxWalletVerifyResponse**](CreateCbxWalletVerifyResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCbxAdCredit

> GetCbxAdCreditResponse GetCbxAdCredit(ctx).Execute()

Unspent advertising credit, in cents.



### Example

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
    resp, r, err := apiClient.CBXApi.GetCbxAdCredit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.GetCbxAdCredit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCbxAdCredit`: GetCbxAdCreditResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.GetCbxAdCredit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCbxAdCreditRequest struct via the builder pattern


### Return type

[**GetCbxAdCreditResponse**](GetCbxAdCreditResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCbxClaim

> GetCbxClaimResponse GetCbxClaim(ctx, claimId).Execute()

A claim's current state.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    claimId := "claimId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.GetCbxClaim(context.Background(), claimId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.GetCbxClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCbxClaim`: GetCbxClaimResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.GetCbxClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**claimId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCbxClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCbxClaimResponse**](GetCbxClaimResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCbxCredit

> GetCbxCreditResponse GetCbxCredit(ctx).Execute()

A seller's wholesale credit line.



### Example

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
    resp, r, err := apiClient.CBXApi.GetCbxCredit(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.GetCbxCredit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCbxCredit`: GetCbxCreditResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.GetCbxCredit`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCbxCreditRequest struct via the builder pattern


### Return type

[**GetCbxCreditResponse**](GetCbxCreditResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCbxMe

> GetCbxMeResponse GetCbxMe(ctx).Execute()

Identity check — which merchant this key belongs to, and its terms.



### Example

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
    resp, r, err := apiClient.CBXApi.GetCbxMe(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.GetCbxMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCbxMe`: GetCbxMeResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.GetCbxMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCbxMeRequest struct via the builder pattern


### Return type

[**GetCbxMeResponse**](GetCbxMeResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCbxPool

> GetCbxPoolResponse GetCbxPool(ctx).Execute()

Your events-pool balance.



### Example

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
    resp, r, err := apiClient.CBXApi.GetCbxPool(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.GetCbxPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCbxPool`: GetCbxPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.GetCbxPool`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCbxPoolRequest struct via the builder pattern


### Return type

[**GetCbxPoolResponse**](GetCbxPoolResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCbxRevenue

> GetCbxRevenueResponse GetCbxRevenue(ctx).Execute()

Operator revenue accrued and not yet withdrawn.



### Example

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
    resp, r, err := apiClient.CBXApi.GetCbxRevenue(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.GetCbxRevenue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCbxRevenue`: GetCbxRevenueResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.GetCbxRevenue`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCbxRevenueRequest struct via the builder pattern


### Return type

[**GetCbxRevenueResponse**](GetCbxRevenueResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCbxSubjectBalance

> GetCbxSubjectBalanceResponse GetCbxSubjectBalance(ctx, subjectId).Execute()

What a subject holds: pending cents and available CBX.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subjectId := "subjectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.GetCbxSubjectBalance(context.Background(), subjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.GetCbxSubjectBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCbxSubjectBalance`: GetCbxSubjectBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.GetCbxSubjectBalance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCbxSubjectBalanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCbxSubjectBalanceResponse**](GetCbxSubjectBalanceResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCbxSubjectBalanceBySubjectId

> GetCbxSubjectBalanceBySubjectIdResponse GetCbxSubjectBalanceBySubjectId(ctx, subjectId).Execute()

All three balances a subject holds.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subjectId := "subjectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.GetCbxSubjectBalanceBySubjectId(context.Background(), subjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.GetCbxSubjectBalanceBySubjectId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCbxSubjectBalanceBySubjectId`: GetCbxSubjectBalanceBySubjectIdResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.GetCbxSubjectBalanceBySubjectId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCbxSubjectBalanceBySubjectIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCbxSubjectBalanceBySubjectIdResponse**](GetCbxSubjectBalanceBySubjectIdResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCbxSubjectSpent

> GetCbxSubjectSpentResponse GetCbxSubjectSpent(ctx, subjectId).Execute()

Total CBX a subject has spent in your marketplace.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subjectId := "subjectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.GetCbxSubjectSpent(context.Background(), subjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.GetCbxSubjectSpent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCbxSubjectSpent`: GetCbxSubjectSpentResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.GetCbxSubjectSpent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCbxSubjectSpentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCbxSubjectSpentResponse**](GetCbxSubjectSpentResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCbxSubjectStake

> GetCbxSubjectStakeResponse GetCbxSubjectStake(ctx, subjectId).Execute()

A subject's staking state and spendable balance.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subjectId := "subjectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.GetCbxSubjectStake(context.Background(), subjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.GetCbxSubjectStake``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCbxSubjectStake`: GetCbxSubjectStakeResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.GetCbxSubjectStake`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCbxSubjectStakeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCbxSubjectStakeResponse**](GetCbxSubjectStakeResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCbxSubjectWallet

> GetCbxSubjectWalletResponse GetCbxSubjectWallet(ctx, subjectId).Execute()

The verified payout address for a subject, if any.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subjectId := "subjectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.GetCbxSubjectWallet(context.Background(), subjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.GetCbxSubjectWallet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCbxSubjectWallet`: GetCbxSubjectWalletResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.GetCbxSubjectWallet`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCbxSubjectWalletRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetCbxSubjectWalletResponse**](GetCbxSubjectWalletResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCbxTreasury

> GetCbxTreasuryResponse GetCbxTreasury(ctx).Execute()

Your most recent reserve reconciliation.



### Example

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
    resp, r, err := apiClient.CBXApi.GetCbxTreasury(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.GetCbxTreasury``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCbxTreasury`: GetCbxTreasuryResponse
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.GetCbxTreasury`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCbxTreasuryRequest struct via the builder pattern


### Return type

[**GetCbxTreasuryResponse**](GetCbxTreasuryResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCbxAdCreditLedger

> V1List ListCbxAdCreditLedger(ctx).Execute()

Ad-credit movements, newest first.



### Example

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
    resp, r, err := apiClient.CBXApi.ListCbxAdCreditLedger(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.ListCbxAdCreditLedger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCbxAdCreditLedger`: V1List
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.ListCbxAdCreditLedger`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCbxAdCreditLedgerRequest struct via the builder pattern


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


## ListCbxBoosts

> V1List ListCbxBoosts(ctx).Execute()

Your funded cashback boosts, newest first.



### Example

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
    resp, r, err := apiClient.CBXApi.ListCbxBoosts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.ListCbxBoosts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCbxBoosts`: V1List
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.ListCbxBoosts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCbxBoostsRequest struct via the builder pattern


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


## ListCbxCampaignPayouts

> V1List ListCbxCampaignPayouts(ctx, campaignId).Execute()

What a campaign actually paid, with the weight behind each amount.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    campaignId := "campaignId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.ListCbxCampaignPayouts(context.Background(), campaignId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.ListCbxCampaignPayouts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCbxCampaignPayouts`: V1List
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.ListCbxCampaignPayouts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCbxCampaignPayoutsRequest struct via the builder pattern


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


## ListCbxCampaigns

> V1List ListCbxCampaigns(ctx).Execute()

Your campaigns, newest first.

### Example

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
    resp, r, err := apiClient.CBXApi.ListCbxCampaigns(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.ListCbxCampaigns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCbxCampaigns`: V1List
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.ListCbxCampaigns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCbxCampaignsRequest struct via the builder pattern


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


## ListCbxDisbursementProgress

> V1List ListCbxDisbursementProgress(ctx).Execute()

How close each rule is to firing — the public counter.



### Example

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
    resp, r, err := apiClient.CBXApi.ListCbxDisbursementProgress(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.ListCbxDisbursementProgress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCbxDisbursementProgress`: V1List
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.ListCbxDisbursementProgress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCbxDisbursementProgressRequest struct via the builder pattern


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


## ListCbxDisbursementRules

> V1List ListCbxDisbursementRules(ctx).Execute()

Threshold rules that fire community distributions.

### Example

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
    resp, r, err := apiClient.CBXApi.ListCbxDisbursementRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.ListCbxDisbursementRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCbxDisbursementRules`: V1List
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.ListCbxDisbursementRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCbxDisbursementRulesRequest struct via the builder pattern


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


## ListCbxEarnTiers

> V1List ListCbxEarnTiers(ctx).Execute()

Earn terms on offer — longer maturation, higher rate.



### Example

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
    resp, r, err := apiClient.CBXApi.ListCbxEarnTiers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.ListCbxEarnTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCbxEarnTiers`: V1List
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.ListCbxEarnTiers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCbxEarnTiersRequest struct via the builder pattern


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


## ListCbxRedemptionServices

> V1List ListCbxRedemptionServices(ctx).Execute()

Services payable in CBX, and the discount each carries.



### Example

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
    resp, r, err := apiClient.CBXApi.ListCbxRedemptionServices(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.ListCbxRedemptionServices``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCbxRedemptionServices`: V1List
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.ListCbxRedemptionServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCbxRedemptionServicesRequest struct via the builder pattern


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


## ListCbxStakeTiers

> V1List ListCbxStakeTiers(ctx).Execute()

Staking tiers — what locking tokens buys.



### Example

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
    resp, r, err := apiClient.CBXApi.ListCbxStakeTiers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.ListCbxStakeTiers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCbxStakeTiers`: V1List
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.ListCbxStakeTiers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCbxStakeTiersRequest struct via the builder pattern


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


## ListCbxSubjectGrants

> V1List ListCbxSubjectGrants(ctx, subjectId).Execute()

Live grants, soonest-expiring first.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subjectId := "subjectId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.ListCbxSubjectGrants(context.Background(), subjectId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.ListCbxSubjectGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCbxSubjectGrants`: V1List
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.ListCbxSubjectGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCbxSubjectGrantsRequest struct via the builder pattern


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


## ListCbxSubjectLedger

> V1List ListCbxSubjectLedger(ctx, subjectId).Limit(limit).Execute()

A subject's CBX ledger, newest first.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    subjectId := "subjectId_example" // string | 
    limit := int32(56) // int32 |  (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CBXApi.ListCbxSubjectLedger(context.Background(), subjectId).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.ListCbxSubjectLedger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCbxSubjectLedger`: V1List
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.ListCbxSubjectLedger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subjectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCbxSubjectLedgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **int32** |  | [default to 50]

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


## ListCbxWalletPaymentReview

> V1List ListCbxWalletPaymentReview(ctx).Execute()

Payments held for a human — the ops queue.



### Example

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
    resp, r, err := apiClient.CBXApi.ListCbxWalletPaymentReview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CBXApi.ListCbxWalletPaymentReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCbxWalletPaymentReview`: V1List
    fmt.Fprintf(os.Stdout, "Response from `CBXApi.ListCbxWalletPaymentReview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCbxWalletPaymentReviewRequest struct via the builder pattern


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

