# \CreatorRewardsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCreatorRewardFormula**](CreatorRewardsApi.md#GetCreatorRewardFormula) | **Get** /v1/creator-rewards/formula | The Creator Rewards formula: the pool&#39;s share of Crossly&#39;s Crossly Carnival take in force, the counting rules, and any scheduled change (announced at least 90 days ahead).
[**GetCreatorRewardMe**](CreatorRewardsApi.md#GetCreatorRewardMe) | **Get** /v1/creator-rewards/me/{month} | One month of your Creator Rewards (YYYY-MM).
[**GetPublicCreatorRewardReconciliation**](CreatorRewardsApi.md#GetPublicCreatorRewardReconciliation) | **Get** /v1/public/creator-rewards/reconciliation | Public monthly reconciliation of Creator Rewards (no auth): per month the pool funded from Crossly&#39;s take, carried in, distributed, paid (from the ledger), carried over, the formula version and weights, eligible experiences and every check; plus the formula in force and scheduled changes (90-day notice).
[**ListCreatorRewardMe**](CreatorRewardsApi.md#ListCreatorRewardMe) | **Get** /v1/creator-rewards/me | Your Creator Rewards: each month, each experience you own or are paid for (counted minutes and players, exclusions, its amount, your part).
[**ListCreatorRewardMonths**](CreatorRewardsApi.md#ListCreatorRewardMonths) | **Get** /v1/creator-rewards/months | Every calculated month: Crossly&#39;s take, the funded share, the pool, what was distributed and carried, counted minutes and exclusions by reason.



## GetCreatorRewardFormula

> GetCreatorRewardFormulaResponse GetCreatorRewardFormula(ctx).Execute()

The Creator Rewards formula: the pool's share of Crossly's Crossly Carnival take in force, the counting rules, and any scheduled change (announced at least 90 days ahead).

### Example

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
    resp, r, err := apiClient.CreatorRewardsApi.GetCreatorRewardFormula(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorRewardsApi.GetCreatorRewardFormula``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorRewardFormula`: GetCreatorRewardFormulaResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorRewardsApi.GetCreatorRewardFormula`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorRewardFormulaRequest struct via the builder pattern


### Return type

**GetCreatorRewardFormulaResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorRewardMe

> GetCreatorRewardMeResponse GetCreatorRewardMe(ctx, month).Execute()

One month of your Creator Rewards (YYYY-MM).

### Example

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
    resp, r, err := apiClient.CreatorRewardsApi.GetCreatorRewardMe(context.Background(), month).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorRewardsApi.GetCreatorRewardMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorRewardMe`: GetCreatorRewardMeResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorRewardsApi.GetCreatorRewardMe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**month** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorRewardMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetCreatorRewardMeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetPublicCreatorRewardReconciliation

> GetPublicCreatorRewardReconciliationResponse GetPublicCreatorRewardReconciliation(ctx).Limit(limit).Execute()

Public monthly reconciliation of Creator Rewards (no auth): per month the pool funded from Crossly's take, carried in, distributed, paid (from the ledger), carried over, the formula version and weights, eligible experiences and every check; plus the formula in force and scheduled changes (90-day notice).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    limit := int32(56) // int32 | Months to return, newest first (default 24). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CreatorRewardsApi.GetPublicCreatorRewardReconciliation(context.Background()).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorRewardsApi.GetPublicCreatorRewardReconciliation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPublicCreatorRewardReconciliation`: GetPublicCreatorRewardReconciliationResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorRewardsApi.GetPublicCreatorRewardReconciliation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetPublicCreatorRewardReconciliationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Months to return, newest first (default 24). | 

### Return type

**GetPublicCreatorRewardReconciliationResponse**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCreatorRewardMe

> ListCreatorRewardMeResponse ListCreatorRewardMe(ctx).Execute()

Your Creator Rewards: each month, each experience you own or are paid for (counted minutes and players, exclusions, its amount, your part).

### Example

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
    resp, r, err := apiClient.CreatorRewardsApi.ListCreatorRewardMe(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorRewardsApi.ListCreatorRewardMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreatorRewardMe`: ListCreatorRewardMeResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorRewardsApi.ListCreatorRewardMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCreatorRewardMeRequest struct via the builder pattern


### Return type

**ListCreatorRewardMeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCreatorRewardMonths

> ListCreatorRewardMonthsResponse ListCreatorRewardMonths(ctx).Execute()

Every calculated month: Crossly's take, the funded share, the pool, what was distributed and carried, counted minutes and exclusions by reason.

### Example

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
    resp, r, err := apiClient.CreatorRewardsApi.ListCreatorRewardMonths(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorRewardsApi.ListCreatorRewardMonths``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreatorRewardMonths`: ListCreatorRewardMonthsResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorRewardsApi.ListCreatorRewardMonths`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCreatorRewardMonthsRequest struct via the builder pattern


### Return type

**ListCreatorRewardMonthsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

