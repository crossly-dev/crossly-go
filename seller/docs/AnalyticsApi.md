# \AnalyticsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAnalyticBookkeeping**](AnalyticsApi.md#GetAnalyticBookkeeping) | **Get** /v1/analytics/bookkeeping | Monthly P&amp;L + per-platform breakdown for a calendar year.
[**GetAnalyticByPlatform**](AnalyticsApi.md#GetAnalyticByPlatform) | **Get** /v1/analytics/by-platform | Sales + revenue grouped by platform for the last N days.
[**GetAnalyticDashboard**](AnalyticsApi.md#GetAnalyticDashboard) | **Get** /v1/analytics/dashboard | Composite dashboard: KPIs + breakdowns + recent activity.
[**GetAnalyticItem**](AnalyticsApi.md#GetAnalyticItem) | **Get** /v1/analytics/items | Per-item P&amp;L for sold inventory.
[**GetAnalyticSummary**](AnalyticsApi.md#GetAnalyticSummary) | **Get** /v1/analytics/summary | Headline KPIs for the last N days.
[**GetAnalyticTimesery**](AnalyticsApi.md#GetAnalyticTimesery) | **Get** /v1/analytics/timeseries | Daily sales + revenue series for the last N days.
[**GetAnalyticToday**](AnalyticsApi.md#GetAnalyticToday) | **Get** /v1/analytics/today | Today&#39;s checklist + 14-day activity streak.
[**ListInsightByPlatform**](AnalyticsApi.md#ListInsightByPlatform) | **Get** /v1/insights/by-platform | Platform velocity + margin insight (90-day window).



## GetAnalyticBookkeeping

> GetAnalyticBookkeepingResponse GetAnalyticBookkeeping(ctx).Execute()

Monthly P&L + per-platform breakdown for a calendar year.

### Example

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
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticBookkeeping(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticBookkeeping``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticBookkeeping`: GetAnalyticBookkeepingResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticBookkeeping`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticBookkeepingRequest struct via the builder pattern


### Return type

[**GetAnalyticBookkeepingResponse**](GetAnalyticBookkeepingResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticByPlatform

> GetAnalyticByPlatformResponse GetAnalyticByPlatform(ctx).Execute()

Sales + revenue grouped by platform for the last N days.

### Example

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
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticByPlatform(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticByPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticByPlatform`: GetAnalyticByPlatformResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticByPlatform`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticByPlatformRequest struct via the builder pattern


### Return type

[**GetAnalyticByPlatformResponse**](GetAnalyticByPlatformResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticDashboard

> GetAnalyticDashboardResponse GetAnalyticDashboard(ctx).Execute()

Composite dashboard: KPIs + breakdowns + recent activity.

### Example

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
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticDashboard(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticDashboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticDashboard`: GetAnalyticDashboardResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticDashboard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticDashboardRequest struct via the builder pattern


### Return type

[**GetAnalyticDashboardResponse**](GetAnalyticDashboardResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticItem

> GetAnalyticItemResponse GetAnalyticItem(ctx).Execute()

Per-item P&L for sold inventory.

### Example

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
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticItem(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticItem`: GetAnalyticItemResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticItem`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticItemRequest struct via the builder pattern


### Return type

[**GetAnalyticItemResponse**](GetAnalyticItemResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticSummary

> GetAnalyticSummaryResponse GetAnalyticSummary(ctx).Days(days).Execute()

Headline KPIs for the last N days.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    days := int32(56) // int32 |  (optional) (default to 30)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticSummary(context.Background()).Days(days).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticSummary`: GetAnalyticSummaryResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** |  | [default to 30]

### Return type

[**GetAnalyticSummaryResponse**](GetAnalyticSummaryResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticTimesery

> GetAnalyticTimeseryResponse GetAnalyticTimesery(ctx).Execute()

Daily sales + revenue series for the last N days.

### Example

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
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticTimesery(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticTimesery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticTimesery`: GetAnalyticTimeseryResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticTimesery`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticTimeseryRequest struct via the builder pattern


### Return type

[**GetAnalyticTimeseryResponse**](GetAnalyticTimeseryResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAnalyticToday

> GetAnalyticTodayResponse GetAnalyticToday(ctx).Execute()

Today's checklist + 14-day activity streak.

### Example

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
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticToday(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticToday``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticToday`: GetAnalyticTodayResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticToday`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticTodayRequest struct via the builder pattern


### Return type

[**GetAnalyticTodayResponse**](GetAnalyticTodayResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInsightByPlatform

> V1List ListInsightByPlatform(ctx).Execute()

Platform velocity + margin insight (90-day window).

### Example

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
    resp, r, err := apiClient.AnalyticsApi.ListInsightByPlatform(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.ListInsightByPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInsightByPlatform`: V1List
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.ListInsightByPlatform`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListInsightByPlatformRequest struct via the builder pattern


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

