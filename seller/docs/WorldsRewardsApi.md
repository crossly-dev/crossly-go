# \WorldsRewardsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorldExperienceRewardClose**](WorldsRewardsApi.md#CreateWorldExperienceRewardClose) | **Post** /v1/worlds/experiences/{experienceId}/rewards/close | Close the pool: the unspent balance returns to your Crossly Carnival earnings. Free rewards keep working.
[**CreateWorldExperienceRewardDefinition**](WorldsRewardsApi.md#CreateWorldExperienceRewardDefinition) | **Post** /v1/worlds/experiences/{experienceId}/rewards/definitions | Add a reward scripts can grant: pin, emote, decor, item (your own COSM design), cosm or cbx (with an amount).
[**CreateWorldExperienceRewardDeposit**](WorldsRewardsApi.md#CreateWorldExperienceRewardDeposit) | **Post** /v1/worlds/experiences/{experienceId}/rewards/deposits | Move available Crossly Carnival earnings into the pool ($1-$10,000). Owner only; idempotent on idempotencyKey.
[**CreateWorldExperienceRewardTopup**](WorldsRewardsApi.md#CreateWorldExperienceRewardTopup) | **Post** /v1/worlds/experiences/{experienceId}/rewards/topups | Open a card top-up ($5-$10,000 into the pool; the 2.9% + 30¢ card fee is added on top). Returns checkoutUrl, where the owner pays with Crossly checkout; the pool is credited once when the payment lands. Owner only.
[**CreateWorldRewardTopupCancel**](WorldsRewardsApi.md#CreateWorldRewardTopupCancel) | **Post** /v1/worlds/rewards/topups/{topupId}/cancel | Cancel an unpaid top-up sheet.
[**GetWorldExperienceReward**](WorldsRewardsApi.md#GetWorldExperienceReward) | **Get** /v1/worlds/experiences/{experienceId}/rewards | The funded reward pool: balance, reserved (owed) value, caps, today&#39;s grants, the reward catalogue and the published rates.
[**GetWorldRewardTopup**](WorldsRewardsApi.md#GetWorldRewardTopup) | **Get** /v1/worlds/rewards/topups/{topupId} | One card top-up: its status (open, paying, paid, cancelled, expired, refused), amounts and the pool it funds.
[**ListWorldExperienceRewardGrants**](WorldsRewardsApi.md#ListWorldExperienceRewardGrants) | **Get** /v1/worlds/experiences/{experienceId}/rewards/grants | Every grant, newest first: reward, player sub, amount, cost, status. Cursor-paged; ?sub&#x3D; for one player.
[**ListWorldExperienceRewardLedger**](WorldsRewardsApi.md#ListWorldExperienceRewardLedger) | **Get** /v1/worlds/experiences/{experienceId}/rewards/ledger | The pool&#39;s own ledger: deposits, grant draw-downs, refunds. Cursor-paged.
[**ListWorldExperienceRewardTopups**](WorldsRewardsApi.md#ListWorldExperienceRewardTopups) | **Get** /v1/worlds/experiences/{experienceId}/rewards/topups | Card top-ups of the reward pool, newest first: amount into the pool, card fee, status, what went back to the card.
[**UpdateWorldExperienceRewardDefinition**](WorldsRewardsApi.md#UpdateWorldExperienceRewardDefinition) | **Patch** /v1/worlds/experiences/{experienceId}/rewards/definitions/{rewardId} | Edit a reward: name, description, image, amount, per-player daily cap, on/off. Key and kind never change.
[**UpdateWorldExperienceRewardPool**](WorldsRewardsApi.md#UpdateWorldExperienceRewardPool) | **Patch** /v1/worlds/experiences/{experienceId}/rewards/pool | Set the pool&#39;s daily caps: total value per day, value and grant count per player per day.



## CreateWorldExperienceRewardClose

> CreateWorldExperienceRewardCloseResponse CreateWorldExperienceRewardClose(ctx, experienceId).Execute()

Close the pool: the unspent balance returns to your Crossly Carnival earnings. Free rewards keep working.

### Example

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
    resp, r, err := apiClient.WorldsRewardsApi.CreateWorldExperienceRewardClose(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsRewardsApi.CreateWorldExperienceRewardClose``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceRewardClose`: CreateWorldExperienceRewardCloseResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsRewardsApi.CreateWorldExperienceRewardClose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceRewardCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceRewardCloseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceRewardDefinition

> CreateWorldExperienceRewardDefinitionResponse CreateWorldExperienceRewardDefinition(ctx, experienceId).Execute()

Add a reward scripts can grant: pin, emote, decor, item (your own COSM design), cosm or cbx (with an amount).

### Example

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
    resp, r, err := apiClient.WorldsRewardsApi.CreateWorldExperienceRewardDefinition(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsRewardsApi.CreateWorldExperienceRewardDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceRewardDefinition`: CreateWorldExperienceRewardDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsRewardsApi.CreateWorldExperienceRewardDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceRewardDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceRewardDefinitionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceRewardDeposit

> CreateWorldExperienceRewardDepositResponse CreateWorldExperienceRewardDeposit(ctx, experienceId).Execute()

Move available Crossly Carnival earnings into the pool ($1-$10,000). Owner only; idempotent on idempotencyKey.

### Example

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
    resp, r, err := apiClient.WorldsRewardsApi.CreateWorldExperienceRewardDeposit(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsRewardsApi.CreateWorldExperienceRewardDeposit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceRewardDeposit`: CreateWorldExperienceRewardDepositResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsRewardsApi.CreateWorldExperienceRewardDeposit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceRewardDepositRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceRewardDepositResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceRewardTopup

> CreateWorldExperienceRewardTopupResponse CreateWorldExperienceRewardTopup(ctx, experienceId).Execute()

Open a card top-up ($5-$10,000 into the pool; the 2.9% + 30¢ card fee is added on top). Returns checkoutUrl, where the owner pays with Crossly checkout; the pool is credited once when the payment lands. Owner only.

### Example

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
    resp, r, err := apiClient.WorldsRewardsApi.CreateWorldExperienceRewardTopup(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsRewardsApi.CreateWorldExperienceRewardTopup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceRewardTopup`: CreateWorldExperienceRewardTopupResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsRewardsApi.CreateWorldExperienceRewardTopup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceRewardTopupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceRewardTopupResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldRewardTopupCancel

> CreateWorldSpaceJoinResponse CreateWorldRewardTopupCancel(ctx, topupId).Execute()

Cancel an unpaid top-up sheet.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    topupId := "topupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsRewardsApi.CreateWorldRewardTopupCancel(context.Background(), topupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsRewardsApi.CreateWorldRewardTopupCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldRewardTopupCancel`: CreateWorldSpaceJoinResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsRewardsApi.CreateWorldRewardTopupCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldRewardTopupCancelRequest struct via the builder pattern


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


## GetWorldExperienceReward

> GetWorldExperienceRewardResponse GetWorldExperienceReward(ctx, experienceId).Execute()

The funded reward pool: balance, reserved (owed) value, caps, today's grants, the reward catalogue and the published rates.

### Example

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
    resp, r, err := apiClient.WorldsRewardsApi.GetWorldExperienceReward(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsRewardsApi.GetWorldExperienceReward``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceReward`: GetWorldExperienceRewardResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsRewardsApi.GetWorldExperienceReward`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceRewardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceRewardResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldRewardTopup

> GetWorldRewardTopupResponse GetWorldRewardTopup(ctx, topupId).Execute()

One card top-up: its status (open, paying, paid, cancelled, expired, refused), amounts and the pool it funds.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    topupId := "topupId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsRewardsApi.GetWorldRewardTopup(context.Background(), topupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsRewardsApi.GetWorldRewardTopup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldRewardTopup`: GetWorldRewardTopupResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsRewardsApi.GetWorldRewardTopup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**topupId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldRewardTopupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldRewardTopupResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceRewardGrants

> ListWorldExperienceRewardGrantsResponse ListWorldExperienceRewardGrants(ctx, experienceId).Execute()

Every grant, newest first: reward, player sub, amount, cost, status. Cursor-paged; ?sub= for one player.

### Example

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
    resp, r, err := apiClient.WorldsRewardsApi.ListWorldExperienceRewardGrants(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsRewardsApi.ListWorldExperienceRewardGrants``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceRewardGrants`: ListWorldExperienceRewardGrantsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsRewardsApi.ListWorldExperienceRewardGrants`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceRewardGrantsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceRewardGrantsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceRewardLedger

> ListWorldExperienceRewardLedgerResponse ListWorldExperienceRewardLedger(ctx, experienceId).Execute()

The pool's own ledger: deposits, grant draw-downs, refunds. Cursor-paged.

### Example

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
    resp, r, err := apiClient.WorldsRewardsApi.ListWorldExperienceRewardLedger(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsRewardsApi.ListWorldExperienceRewardLedger``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceRewardLedger`: ListWorldExperienceRewardLedgerResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsRewardsApi.ListWorldExperienceRewardLedger`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceRewardLedgerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceRewardLedgerResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceRewardTopups

> ListWorldExperienceRewardTopupsResponse ListWorldExperienceRewardTopups(ctx, experienceId).Execute()

Card top-ups of the reward pool, newest first: amount into the pool, card fee, status, what went back to the card.

### Example

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
    resp, r, err := apiClient.WorldsRewardsApi.ListWorldExperienceRewardTopups(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsRewardsApi.ListWorldExperienceRewardTopups``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceRewardTopups`: ListWorldExperienceRewardTopupsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsRewardsApi.ListWorldExperienceRewardTopups`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceRewardTopupsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceRewardTopupsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceRewardDefinition

> CreateWorldExperienceRewardDefinitionResponse UpdateWorldExperienceRewardDefinition(ctx, experienceId, rewardId).Execute()

Edit a reward: name, description, image, amount, per-player daily cap, on/off. Key and kind never change.

### Example

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
    rewardId := "rewardId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsRewardsApi.UpdateWorldExperienceRewardDefinition(context.Background(), experienceId, rewardId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsRewardsApi.UpdateWorldExperienceRewardDefinition``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceRewardDefinition`: CreateWorldExperienceRewardDefinitionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsRewardsApi.UpdateWorldExperienceRewardDefinition`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**rewardId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceRewardDefinitionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperienceRewardDefinitionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceRewardPool

> UpdateWorldExperienceRewardPoolResponse UpdateWorldExperienceRewardPool(ctx, experienceId).Execute()

Set the pool's daily caps: total value per day, value and grant count per player per day.

### Example

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
    resp, r, err := apiClient.WorldsRewardsApi.UpdateWorldExperienceRewardPool(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsRewardsApi.UpdateWorldExperienceRewardPool``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceRewardPool`: UpdateWorldExperienceRewardPoolResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsRewardsApi.UpdateWorldExperienceRewardPool`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceRewardPoolRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateWorldExperienceRewardPoolResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

