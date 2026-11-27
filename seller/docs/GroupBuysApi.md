# \GroupBuysApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupBuy**](GroupBuysApi.md#CreateGroupBuy) | **Post** /v1/group-buys | Run a group buy on a listing with stock: price tiers by committed quantity, a goal and minimum to proceed, a per-buyer max, a deadline, authorize-now or pay-now-refund-the-difference, a goal-miss policy (cancel, proceed, extend once), visibility, placement and share copy. The original body (tiers by buyer count, minBuyers, days) is still accepted.
[**CreateGroupBuyCancel**](GroupBuysApi.md#CreateGroupBuyCancel) | **Post** /v1/group-buys/{id}/cancel | Cancel a group buy: every hold is released and every charge refunded.
[**CreateGroupBuyClose**](GroupBuysApi.md#CreateGroupBuyClose) | **Post** /v1/group-buys/{id}/close | Close a group buy now: every buyer pays the price reached (the minimum and goal-miss policy still apply).
[**CreateGroupBuyExtend**](GroupBuysApi.md#CreateGroupBuyExtend) | **Post** /v1/group-buys/{id}/extend | Move the deadline later by N days (authorize-now pools run at most 6 days, pay-now pools 30).
[**CreateGroupBuyMessage**](GroupBuysApi.md#CreateGroupBuyMessage) | **Post** /v1/group-buys/{id}/message | Send a notification to every buyer committed to the group buy.
[**GetGroupBuy**](GroupBuysApi.md#GetGroupBuy) | **Get** /v1/group-buys/{id} | One group buy: progress towards each tier, committed buyers (masked), time left, revenue at the current tier and the activity feed.
[**GetGroupBuyParticipantsCsv**](GroupBuysApi.md#GetGroupBuyParticipantsCsv) | **Get** /v1/group-buys/{id}/participants.csv | Every participant of a group buy as CSV: quantity, status, amounts and ship-to.
[**ListGroupBuys**](GroupBuysApi.md#ListGroupBuys) | **Get** /v1/group-buys | Your group buys, open ones first: tiers, buyers joined, the price reached and when it ends.
[**UpdateGroupBuy**](GroupBuysApi.md#UpdateGroupBuy) | **Patch** /v1/group-buys/{id} | Edit a group buy. Copy, visibility, placement and the per-buyer max change any time; tiers, goal, minimum, payment mode and goal-miss policy only before anyone commits.



## CreateGroupBuy

> CreateBreakResponse CreateGroupBuy(ctx).Execute()

Run a group buy on a listing with stock: price tiers by committed quantity, a goal and minimum to proceed, a per-buyer max, a deadline, authorize-now or pay-now-refund-the-difference, a goal-miss policy (cancel, proceed, extend once), visibility, placement and share copy. The original body (tiers by buyer count, minBuyers, days) is still accepted.

### Example

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
    resp, r, err := apiClient.GroupBuysApi.CreateGroupBuy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupBuysApi.CreateGroupBuy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroupBuy`: CreateBreakResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupBuysApi.CreateGroupBuy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupBuyRequest struct via the builder pattern


### Return type

**CreateBreakResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateGroupBuyCancel

> CreateGroupBuyCancelResponse CreateGroupBuyCancel(ctx, id).Execute()

Cancel a group buy: every hold is released and every charge refunded.

### Example

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
    resp, r, err := apiClient.GroupBuysApi.CreateGroupBuyCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupBuysApi.CreateGroupBuyCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroupBuyCancel`: CreateGroupBuyCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupBuysApi.CreateGroupBuyCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupBuyCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateGroupBuyCancelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateGroupBuyClose

> CreateGroupBuyCloseResponse CreateGroupBuyClose(ctx, id).Execute()

Close a group buy now: every buyer pays the price reached (the minimum and goal-miss policy still apply).

### Example

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
    resp, r, err := apiClient.GroupBuysApi.CreateGroupBuyClose(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupBuysApi.CreateGroupBuyClose``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroupBuyClose`: CreateGroupBuyCloseResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupBuysApi.CreateGroupBuyClose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupBuyCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateGroupBuyCloseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateGroupBuyExtend

> CreateGroupBuyExtendResponse CreateGroupBuyExtend(ctx, id).Execute()

Move the deadline later by N days (authorize-now pools run at most 6 days, pay-now pools 30).

### Example

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
    resp, r, err := apiClient.GroupBuysApi.CreateGroupBuyExtend(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupBuysApi.CreateGroupBuyExtend``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroupBuyExtend`: CreateGroupBuyExtendResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupBuysApi.CreateGroupBuyExtend`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupBuyExtendRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateGroupBuyExtendResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateGroupBuyMessage

> CreateGroupBuyMessageResponse CreateGroupBuyMessage(ctx, id).Execute()

Send a notification to every buyer committed to the group buy.

### Example

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
    resp, r, err := apiClient.GroupBuysApi.CreateGroupBuyMessage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupBuysApi.CreateGroupBuyMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroupBuyMessage`: CreateGroupBuyMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupBuysApi.CreateGroupBuyMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupBuyMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateGroupBuyMessageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetGroupBuy

> GetGroupBuyResponse GetGroupBuy(ctx, id).Execute()

One group buy: progress towards each tier, committed buyers (masked), time left, revenue at the current tier and the activity feed.

### Example

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
    resp, r, err := apiClient.GroupBuysApi.GetGroupBuy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupBuysApi.GetGroupBuy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupBuy`: GetGroupBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupBuysApi.GetGroupBuy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupBuyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetGroupBuyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetGroupBuyParticipantsCsv

> string GetGroupBuyParticipantsCsv(ctx, id).Execute()

Every participant of a group buy as CSV: quantity, status, amounts and ship-to.

### Example

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
    resp, r, err := apiClient.GroupBuysApi.GetGroupBuyParticipantsCsv(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupBuysApi.GetGroupBuyParticipantsCsv``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupBuyParticipantsCsv`: string
    fmt.Fprintf(os.Stdout, "Response from `GroupBuysApi.GetGroupBuyParticipantsCsv`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupBuyParticipantsCsvRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListGroupBuys

> ListGroupBuysResponse ListGroupBuys(ctx).Execute()

Your group buys, open ones first: tiers, buyers joined, the price reached and when it ends.

### Example

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
    resp, r, err := apiClient.GroupBuysApi.ListGroupBuys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupBuysApi.ListGroupBuys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupBuys`: ListGroupBuysResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupBuysApi.ListGroupBuys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGroupBuysRequest struct via the builder pattern


### Return type

**ListGroupBuysResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateGroupBuy

> UpdateGroupBuyResponse UpdateGroupBuy(ctx, id).Execute()

Edit a group buy. Copy, visibility, placement and the per-buyer max change any time; tiers, goal, minimum, payment mode and goal-miss policy only before anyone commits.

### Example

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
    resp, r, err := apiClient.GroupBuysApi.UpdateGroupBuy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GroupBuysApi.UpdateGroupBuy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroupBuy`: UpdateGroupBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `GroupBuysApi.UpdateGroupBuy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupBuyRequest struct via the builder pattern


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

