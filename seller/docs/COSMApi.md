# \COSMApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCosmAttestationAnchor**](COSMApi.md#CreateCosmAttestationAnchor) | **Post** /v1/cosm/attestations/{id}/anchor | Record a signature for a commitment anchored BY HAND.
[**CreateCosmAuction**](COSMApi.md#CreateCosmAuction) | **Post** /v1/cosm/auction | List an owned cosmetic for COSM.
[**CreateCosmAuctionBuy**](COSMApi.md#CreateCosmAuctionBuy) | **Post** /v1/cosm/auction/{id}/buy | Buy an open listing.
[**CreateCosmCosmeticBuy**](COSMApi.md#CreateCosmCosmeticBuy) | **Post** /v1/cosm/cosmetics/{id}/buy | Buy an earned-tier cosmetic. The COSM is DESTROYED.
[**CreateCosmCosmeticBuyCbx**](COSMApi.md#CreateCosmCosmeticBuyCbx) | **Post** /v1/cosm/cosmetics/{id}/buy-cbx | Buy a dollar-priced cosmetic, paying in CBX.
[**CreateCosmCosmeticClaim**](COSMApi.md#CreateCosmCosmeticClaim) | **Post** /v1/cosm/cosmetics/{id}/claim | Claim a stake-exclusive cosmetic once tenure qualifies.
[**CreateCosmFusion**](COSMApi.md#CreateCosmFusion) | **Post** /v1/cosm/fusion/{recipeId} | Fuse two cosmetics plus COSM into a rarer one.
[**CreateCosmHandle**](COSMApi.md#CreateCosmHandle) | **Post** /v1/cosm/handles | Claim a vanity handle, burning COSM.
[**CreateCosmWithdrawal**](COSMApi.md#CreateCosmWithdrawal) | **Post** /v1/cosm/withdrawals | Request a withdrawal to your own proven wallet. REFUSES today.
[**DeleteCosmHandle**](COSMApi.md#DeleteCosmHandle) | **Delete** /v1/cosm/handles | Release a handle. NOTHING IS REFUNDED.
[**DeleteCosmPeg**](COSMApi.md#DeleteCosmPeg) | **Delete** /v1/cosm/pegs/{listingId} | Stop repricing. The current ask stays where it is.
[**GetCosmAttestationProof**](COSMApi.md#GetCosmAttestationProof) | **Get** /v1/cosm/attestations/{id}/proof | A Merkle inclusion proof for one subject, against a published root.
[**GetCosmAuctionQuote**](COSMApi.md#GetCosmAuctionQuote) | **Get** /v1/cosm/auction/quote | What an auction listing would cost in fees.
[**GetCosmBalance**](COSMApi.md#GetCosmBalance) | **Get** /v1/cosm/balance | A subject&#39;s COSM balance and the earn rate in force.
[**GetCosmEmployeeFeeQuote**](COSMApi.md#GetCosmEmployeeFeeQuote) | **Get** /v1/cosm/employee/fee-quote | What one shop-employee conversation will cost, in COSM.
[**GetCosmHandle**](COSMApi.md#GetCosmHandle) | **Get** /v1/cosm/handles/{handle} | Is this handle available, and what does it cost?
[**GetCosmStakeTenure**](COSMApi.md#GetCosmStakeTenure) | **Get** /v1/cosm/stake-tenure | How long a subject has held a stake continuously, in days.
[**GetCosmSupply**](COSMApi.md#GetCosmSupply) | **Get** /v1/cosm/supply | Total supply — minted, burned, circulating.
[**ListCosmAttestations**](COSMApi.md#ListCosmAttestations) | **Get** /v1/cosm/attestations | Published Merkle commitments. What you can check instead of trusting us.
[**ListCosmAuction**](COSMApi.md#ListCosmAuction) | **Get** /v1/cosm/auction | Open auction listings, priced in COSM.
[**ListCosmCosmetics**](COSMApi.md#ListCosmCosmetics) | **Get** /v1/cosm/cosmetics | The cosmetics catalogue.
[**ListCosmGoals**](COSMApi.md#ListCosmGoals) | **Get** /v1/cosm/goals | Community burn goal progress.
[**ListCosmInventory**](COSMApi.md#ListCosmInventory) | **Get** /v1/cosm/inventory | Cosmetics a subject holds.
[**ListCosmPegs**](COSMApi.md#ListCosmPegs) | **Get** /v1/cosm/pegs | CBX-pegged listings and why each last moved.
[**ListCosmSupplyEpochs**](COSMApi.md#ListCosmSupplyEpochs) | **Get** /v1/cosm/supply/epochs | Per-epoch emission against the backstop.
[**ListCosmSupplySchedule**](COSMApi.md#ListCosmSupplySchedule) | **Get** /v1/cosm/supply/schedule | The published emission taper. Fixed years ahead.
[**ListCosmUnlocks**](COSMApi.md#ListCosmUnlocks) | **Get** /v1/cosm/unlocks | Which spatial cosmetic values this subject may apply.
[**ListCosmWithdrawals**](COSMApi.md#ListCosmWithdrawals) | **Get** /v1/cosm/withdrawals | Withdrawal history, and whether withdrawals are open at all.
[**UpdateCosmCosmeticUnlock**](COSMApi.md#UpdateCosmCosmeticUnlock) | **Put** /v1/cosm/cosmetics/{id}/unlock | Set (or clear) what a cosmetic unlocks spatially.
[**UpdateCosmPeg**](COSMApi.md#UpdateCosmPeg) | **Put** /v1/cosm/pegs | Set or update a CBX price peg. OFF unless you turn it on.



## CreateCosmAttestationAnchor

> CreateCosmAttestationAnchorResponse CreateCosmAttestationAnchor(ctx, id).Execute()

Record a signature for a commitment anchored BY HAND.



### Example

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
    resp, r, err := apiClient.COSMApi.CreateCosmAttestationAnchor(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.CreateCosmAttestationAnchor``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCosmAttestationAnchor`: CreateCosmAttestationAnchorResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.CreateCosmAttestationAnchor`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCosmAttestationAnchorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateCosmAttestationAnchorResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCosmAuction

> CreateCosmAuctionResponse CreateCosmAuction(ctx).Execute()

List an owned cosmetic for COSM.



### Example

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
    resp, r, err := apiClient.COSMApi.CreateCosmAuction(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.CreateCosmAuction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCosmAuction`: CreateCosmAuctionResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.CreateCosmAuction`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCosmAuctionRequest struct via the builder pattern


### Return type

**CreateCosmAuctionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCosmAuctionBuy

> CreateCosmFusionResponse CreateCosmAuctionBuy(ctx, id).Execute()

Buy an open listing.



### Example

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
    resp, r, err := apiClient.COSMApi.CreateCosmAuctionBuy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.CreateCosmAuctionBuy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCosmAuctionBuy`: CreateCosmFusionResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.CreateCosmAuctionBuy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCosmAuctionBuyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateCosmFusionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCosmCosmeticBuy

> CreateCosmCosmeticBuyResponse CreateCosmCosmeticBuy(ctx, id).Execute()

Buy an earned-tier cosmetic. The COSM is DESTROYED.



### Example

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
    resp, r, err := apiClient.COSMApi.CreateCosmCosmeticBuy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.CreateCosmCosmeticBuy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCosmCosmeticBuy`: CreateCosmCosmeticBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.CreateCosmCosmeticBuy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCosmCosmeticBuyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateCosmCosmeticBuyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCosmCosmeticBuyCbx

> CreateCosmCosmeticBuyCbxResponse CreateCosmCosmeticBuyCbx(ctx, id).Execute()

Buy a dollar-priced cosmetic, paying in CBX.



### Example

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
    resp, r, err := apiClient.COSMApi.CreateCosmCosmeticBuyCbx(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.CreateCosmCosmeticBuyCbx``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCosmCosmeticBuyCbx`: CreateCosmCosmeticBuyCbxResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.CreateCosmCosmeticBuyCbx`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCosmCosmeticBuyCbxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateCosmCosmeticBuyCbxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCosmCosmeticClaim

> CreateCosmCosmeticClaimResponse CreateCosmCosmeticClaim(ctx, id).Execute()

Claim a stake-exclusive cosmetic once tenure qualifies.



### Example

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
    resp, r, err := apiClient.COSMApi.CreateCosmCosmeticClaim(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.CreateCosmCosmeticClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCosmCosmeticClaim`: CreateCosmCosmeticClaimResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.CreateCosmCosmeticClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCosmCosmeticClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateCosmCosmeticClaimResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCosmFusion

> CreateCosmFusionResponse CreateCosmFusion(ctx, recipeId).Execute()

Fuse two cosmetics plus COSM into a rarer one.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    recipeId := "recipeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.COSMApi.CreateCosmFusion(context.Background(), recipeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.CreateCosmFusion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCosmFusion`: CreateCosmFusionResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.CreateCosmFusion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**recipeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCosmFusionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateCosmFusionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCosmHandle

> CreateCosmHandleResponse CreateCosmHandle(ctx).Execute()

Claim a vanity handle, burning COSM.



### Example

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
    resp, r, err := apiClient.COSMApi.CreateCosmHandle(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.CreateCosmHandle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCosmHandle`: CreateCosmHandleResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.CreateCosmHandle`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCosmHandleRequest struct via the builder pattern


### Return type

**CreateCosmHandleResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCosmWithdrawal

> CreateCosmWithdrawalResponse CreateCosmWithdrawal(ctx).Execute()

Request a withdrawal to your own proven wallet. REFUSES today.



### Example

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
    resp, r, err := apiClient.COSMApi.CreateCosmWithdrawal(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.CreateCosmWithdrawal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCosmWithdrawal`: CreateCosmWithdrawalResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.CreateCosmWithdrawal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCosmWithdrawalRequest struct via the builder pattern


### Return type

**CreateCosmWithdrawalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteCosmHandle

> DeleteCosmHandleResponse DeleteCosmHandle(ctx).Execute()

Release a handle. NOTHING IS REFUNDED.



### Example

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
    resp, r, err := apiClient.COSMApi.DeleteCosmHandle(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.DeleteCosmHandle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCosmHandle`: DeleteCosmHandleResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.DeleteCosmHandle`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCosmHandleRequest struct via the builder pattern


### Return type

**DeleteCosmHandleResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteCosmPeg

> DeleteCosmPegResponse DeleteCosmPeg(ctx, listingId).Execute()

Stop repricing. The current ask stays where it is.



### Example

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
    resp, r, err := apiClient.COSMApi.DeleteCosmPeg(context.Background(), listingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.DeleteCosmPeg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCosmPeg`: DeleteCosmPegResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.DeleteCosmPeg`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCosmPegRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteCosmPegResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCosmAttestationProof

> GetCosmAttestationProofResponse GetCosmAttestationProof(ctx, id).Execute()

A Merkle inclusion proof for one subject, against a published root.



### Example

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
    resp, r, err := apiClient.COSMApi.GetCosmAttestationProof(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.GetCosmAttestationProof``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCosmAttestationProof`: GetCosmAttestationProofResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.GetCosmAttestationProof`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCosmAttestationProofRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetCosmAttestationProofResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCosmAuctionQuote

> GetCosmAuctionQuoteResponse GetCosmAuctionQuote(ctx).Execute()

What an auction listing would cost in fees.



### Example

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
    resp, r, err := apiClient.COSMApi.GetCosmAuctionQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.GetCosmAuctionQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCosmAuctionQuote`: GetCosmAuctionQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.GetCosmAuctionQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCosmAuctionQuoteRequest struct via the builder pattern


### Return type

**GetCosmAuctionQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCosmBalance

> GetCosmBalanceResponse GetCosmBalance(ctx).Execute()

A subject's COSM balance and the earn rate in force.



### Example

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
    resp, r, err := apiClient.COSMApi.GetCosmBalance(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.GetCosmBalance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCosmBalance`: GetCosmBalanceResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.GetCosmBalance`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCosmBalanceRequest struct via the builder pattern


### Return type

**GetCosmBalanceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCosmEmployeeFeeQuote

> GetCosmEmployeeFeeQuoteResponse GetCosmEmployeeFeeQuote(ctx).Execute()

What one shop-employee conversation will cost, in COSM.



### Example

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
    resp, r, err := apiClient.COSMApi.GetCosmEmployeeFeeQuote(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.GetCosmEmployeeFeeQuote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCosmEmployeeFeeQuote`: GetCosmEmployeeFeeQuoteResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.GetCosmEmployeeFeeQuote`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCosmEmployeeFeeQuoteRequest struct via the builder pattern


### Return type

**GetCosmEmployeeFeeQuoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCosmHandle

> GetCosmHandleResponse GetCosmHandle(ctx, handle).Execute()

Is this handle available, and what does it cost?



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    handle := "handle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.COSMApi.GetCosmHandle(context.Background(), handle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.GetCosmHandle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCosmHandle`: GetCosmHandleResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.GetCosmHandle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCosmHandleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetCosmHandleResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCosmStakeTenure

> GetCosmStakeTenureResponse GetCosmStakeTenure(ctx).Execute()

How long a subject has held a stake continuously, in days.



### Example

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
    resp, r, err := apiClient.COSMApi.GetCosmStakeTenure(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.GetCosmStakeTenure``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCosmStakeTenure`: GetCosmStakeTenureResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.GetCosmStakeTenure`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCosmStakeTenureRequest struct via the builder pattern


### Return type

**GetCosmStakeTenureResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCosmSupply

> GetCosmSupplyResponse GetCosmSupply(ctx).Execute()

Total supply — minted, burned, circulating.



### Example

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
    resp, r, err := apiClient.COSMApi.GetCosmSupply(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.GetCosmSupply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCosmSupply`: GetCosmSupplyResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.GetCosmSupply`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCosmSupplyRequest struct via the builder pattern


### Return type

**GetCosmSupplyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCosmAttestations

> ListCosmAttestationsResponse ListCosmAttestations(ctx).Execute()

Published Merkle commitments. What you can check instead of trusting us.



### Example

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
    resp, r, err := apiClient.COSMApi.ListCosmAttestations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.ListCosmAttestations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCosmAttestations`: ListCosmAttestationsResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.ListCosmAttestations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCosmAttestationsRequest struct via the builder pattern


### Return type

**ListCosmAttestationsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCosmAuction

> ListCosmAuctionResponse ListCosmAuction(ctx).Execute()

Open auction listings, priced in COSM.



### Example

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
    resp, r, err := apiClient.COSMApi.ListCosmAuction(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.ListCosmAuction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCosmAuction`: ListCosmAuctionResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.ListCosmAuction`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCosmAuctionRequest struct via the builder pattern


### Return type

**ListCosmAuctionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCosmCosmetics

> ListCosmCosmeticsResponse ListCosmCosmetics(ctx).Execute()

The cosmetics catalogue.



### Example

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
    resp, r, err := apiClient.COSMApi.ListCosmCosmetics(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.ListCosmCosmetics``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCosmCosmetics`: ListCosmCosmeticsResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.ListCosmCosmetics`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCosmCosmeticsRequest struct via the builder pattern


### Return type

**ListCosmCosmeticsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCosmGoals

> ListCosmGoalsResponse ListCosmGoals(ctx).Execute()

Community burn goal progress.



### Example

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
    resp, r, err := apiClient.COSMApi.ListCosmGoals(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.ListCosmGoals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCosmGoals`: ListCosmGoalsResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.ListCosmGoals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCosmGoalsRequest struct via the builder pattern


### Return type

**ListCosmGoalsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCosmInventory

> ListCosmInventoryResponse ListCosmInventory(ctx).Execute()

Cosmetics a subject holds.



### Example

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
    resp, r, err := apiClient.COSMApi.ListCosmInventory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.ListCosmInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCosmInventory`: ListCosmInventoryResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.ListCosmInventory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCosmInventoryRequest struct via the builder pattern


### Return type

**ListCosmInventoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCosmPegs

> ListCosmPegsResponse ListCosmPegs(ctx).Execute()

CBX-pegged listings and why each last moved.



### Example

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
    resp, r, err := apiClient.COSMApi.ListCosmPegs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.ListCosmPegs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCosmPegs`: ListCosmPegsResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.ListCosmPegs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCosmPegsRequest struct via the builder pattern


### Return type

**ListCosmPegsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCosmSupplyEpochs

> ListCosmSupplyEpochsResponse ListCosmSupplyEpochs(ctx).Execute()

Per-epoch emission against the backstop.



### Example

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
    resp, r, err := apiClient.COSMApi.ListCosmSupplyEpochs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.ListCosmSupplyEpochs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCosmSupplyEpochs`: ListCosmSupplyEpochsResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.ListCosmSupplyEpochs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCosmSupplyEpochsRequest struct via the builder pattern


### Return type

**ListCosmSupplyEpochsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCosmSupplySchedule

> ListCosmSupplyScheduleResponse ListCosmSupplySchedule(ctx).Execute()

The published emission taper. Fixed years ahead.



### Example

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
    resp, r, err := apiClient.COSMApi.ListCosmSupplySchedule(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.ListCosmSupplySchedule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCosmSupplySchedule`: ListCosmSupplyScheduleResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.ListCosmSupplySchedule`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCosmSupplyScheduleRequest struct via the builder pattern


### Return type

**ListCosmSupplyScheduleResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCosmUnlocks

> ListCosmUnlocksResponse ListCosmUnlocks(ctx).Execute()

Which spatial cosmetic values this subject may apply.



### Example

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
    resp, r, err := apiClient.COSMApi.ListCosmUnlocks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.ListCosmUnlocks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCosmUnlocks`: ListCosmUnlocksResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.ListCosmUnlocks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCosmUnlocksRequest struct via the builder pattern


### Return type

**ListCosmUnlocksResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCosmWithdrawals

> ListCosmWithdrawalsResponse ListCosmWithdrawals(ctx).Execute()

Withdrawal history, and whether withdrawals are open at all.



### Example

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
    resp, r, err := apiClient.COSMApi.ListCosmWithdrawals(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.ListCosmWithdrawals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCosmWithdrawals`: ListCosmWithdrawalsResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.ListCosmWithdrawals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCosmWithdrawalsRequest struct via the builder pattern


### Return type

**ListCosmWithdrawalsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateCosmCosmeticUnlock

> UpdateCosmCosmeticUnlockResponse UpdateCosmCosmeticUnlock(ctx, id).Execute()

Set (or clear) what a cosmetic unlocks spatially.



### Example

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
    resp, r, err := apiClient.COSMApi.UpdateCosmCosmeticUnlock(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.UpdateCosmCosmeticUnlock``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCosmCosmeticUnlock`: UpdateCosmCosmeticUnlockResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.UpdateCosmCosmeticUnlock`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCosmCosmeticUnlockRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateCosmCosmeticUnlockResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateCosmPeg

> UpdateCosmPegResponse UpdateCosmPeg(ctx).Execute()

Set or update a CBX price peg. OFF unless you turn it on.



### Example

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
    resp, r, err := apiClient.COSMApi.UpdateCosmPeg(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `COSMApi.UpdateCosmPeg``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCosmPeg`: UpdateCosmPegResponse
    fmt.Fprintf(os.Stdout, "Response from `COSMApi.UpdateCosmPeg`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCosmPegRequest struct via the builder pattern


### Return type

**UpdateCosmPegResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

