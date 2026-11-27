# \AnalyticsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAnalyticExperiment**](AnalyticsApi.md#CreateAnalyticExperiment) | **Post** /v1/analytics/experiments | Start a listing experiment. split: 2-4 live platform listings observed over the same window; sequential: one platform listing, before vs after changeAt (optionally applying a new price).
[**CreateAnalyticExperimentDeclare**](AnalyticsApi.md#CreateAnalyticExperimentDeclare) | **Post** /v1/analytics/experiments/{id}/declare | Record the outcome: a winning arm, no difference, or inconclusive. Concludes the experiment.
[**CreateAnalyticExperimentStop**](AnalyticsApi.md#CreateAnalyticExperimentStop) | **Post** /v1/analytics/experiments/{id}/stop | Stop an experiment: no more data is counted after now.
[**DeleteAnalyticExperiment**](AnalyticsApi.md#DeleteAnalyticExperiment) | **Delete** /v1/analytics/experiments/{id} | Delete an experiment. The listings are untouched.
[**GetAnalyticBookkeeping**](AnalyticsApi.md#GetAnalyticBookkeeping) | **Get** /v1/analytics/bookkeeping | Monthly P&amp;L + per-platform breakdown for a calendar year.
[**GetAnalyticByPlatform**](AnalyticsApi.md#GetAnalyticByPlatform) | **Get** /v1/analytics/by-platform | Sales, net sales, fees and net profit by platform for the last N days.
[**GetAnalyticCompare**](AnalyticsApi.md#GetAnalyticCompare) | **Get** /v1/analytics/compare | Compare 2-4 of your listings: columns of price, quality, channels, views, favorites, offers and outcome, engagement curves, the fields that differ, and observational tests of favorite rate, view-to-offer rate and views per day.
[**GetAnalyticDashboard**](AnalyticsApi.md#GetAnalyticDashboard) | **Get** /v1/analytics/dashboard | Composite dashboard: KPIs + breakdowns + recent activity.
[**GetAnalyticExperiment**](AnalyticsApi.md#GetAnalyticExperiment) | **Get** /v1/analytics/experiments | Your listing A/B experiments, newest first, with status, days running and the current verdict.
[**GetAnalyticExperimentById**](AnalyticsApi.md#GetAnalyticExperimentById) | **Get** /v1/analytics/experiments/{id} | One experiment: arms with their windows, exposure and rates, the primary-metric test per arm against A, guardrail tests, the daily series and what to do next.
[**GetAnalyticInventory**](AnalyticsApi.md#GetAnalyticInventory) | **Get** /v1/analytics/inventory | Unsold stock: aging buckets with cost, listed value and market value, the 20 listings most in need of attention, shelf value over time, cohort sell-through at the chosen maturity, listed vs sold velocity and data health.
[**GetAnalyticItem**](AnalyticsApi.md#GetAnalyticItem) | **Get** /v1/analytics/items | Per-item P&amp;L for sold inventory.
[**GetAnalyticItemByListingId**](AnalyticsApi.md#GetAnalyticItemByListingId) | **Get** /v1/analytics/items/{listingId} | One listing in depth: channels with views, favorites and offers, the merged timeline, engagement series, price history, comps, money waterfall, days on market, offers, messages, experiments and suggestions.
[**GetAnalyticMarketplace**](AnalyticsApi.md#GetAnalyticMarketplace) | **Get** /v1/analytics/marketplaces | Every marketplace side by side: net sales and profit, units, ASP, take rate, margin, shipping subsidy, days to sell, sell-through, cancel and return rates, offers and engagement; with vs&#x3D;2-3 platforms, a head-to-head over listings live on all of them.
[**GetAnalyticMarketplaceByPlatform**](AnalyticsApi.md#GetAnalyticMarketplaceByPlatform) | **Get** /v1/analytics/marketplaces/{platform} | One marketplace in depth: its row, the money series, fee breakdown, days-to-sell and price-realization histograms, categories, top items, at-risk listings, a day-by-hour sales heatmap and offer stats.
[**GetAnalyticOverview**](AnalyticsApi.md#GetAnalyticOverview) | **Get** /v1/analytics/overview | KPIs with deltas (net sales, net profit, units, ASP, margin, sell-through, median days to sell, active listings, take rate), the trend series, by-marketplace split, top categories, brands and items, sell-through cohort, data health and insights.
[**GetAnalyticPerformance**](AnalyticsApi.md#GetAnalyticPerformance) | **Get** /v1/analytics/performance | Per-listing performance: channels, views, favorites, offers and conversations in the period, units sold, net sales and profit, margin, days to sell and a 0-100 health score, paged and sortable.
[**GetAnalyticProgram**](AnalyticsApi.md#GetAnalyticProgram) | **Get** /v1/analytics/programs/{program} | One program&#39;s report over a date range (default the last 30 days, at most 366): KPIs, a daily trend and a breakdown table. Money in integer cents.
[**GetAnalyticSegment**](AnalyticsApi.md#GetAnalyticSegment) | **Get** /v1/analytics/segments | Niches along one dimension: listed and active counts, units sold, cohort sell-through, median days to sell, net sales and profit, margin, ROI, ASP and discount from ask; rows under minN are flagged lowSample.
[**GetAnalyticSegmentByKey**](AnalyticsApi.md#GetAnalyticSegmentByKey) | **Get** /v1/analytics/segments/{dim}/{key} | One niche in depth: its row, series, by-marketplace split, days-to-sell histogram, price bands, top items and the 10 oldest unsold listings in it.
[**GetAnalyticSummary**](AnalyticsApi.md#GetAnalyticSummary) | **Get** /v1/analytics/summary | Headline KPIs for the last N days (the same numbers as the analytics overview).
[**GetAnalyticTimesery**](AnalyticsApi.md#GetAnalyticTimesery) | **Get** /v1/analytics/timeseries | Sales and net sales over the last N days, zero-filled, one row per calendar day.
[**GetAnalyticToday**](AnalyticsApi.md#GetAnalyticToday) | **Get** /v1/analytics/today | Today&#39;s checklist + 14-day activity streak.
[**ListInsightByPlatform**](AnalyticsApi.md#ListInsightByPlatform) | **Get** /v1/insights/by-platform | Platform velocity + margin insight (90-day window): units, net sales, median days to sell, ASP and net after the real marketplace fees.
[**UpdateAnalyticExperiment**](AnalyticsApi.md#UpdateAnalyticExperiment) | **Patch** /v1/analytics/experiments/{id} | Rename an experiment or change its hypothesis, notes or planned end.



## CreateAnalyticExperiment

> AnalyticExperimentResponse CreateAnalyticExperiment(ctx).Execute()

Start a listing experiment. split: 2-4 live platform listings observed over the same window; sequential: one platform listing, before vs after changeAt (optionally applying a new price).

### Example

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
    resp, r, err := apiClient.AnalyticsApi.CreateAnalyticExperiment(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.CreateAnalyticExperiment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAnalyticExperiment`: AnalyticExperimentResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.CreateAnalyticExperiment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnalyticExperimentRequest struct via the builder pattern


### Return type

**AnalyticExperimentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateAnalyticExperimentDeclare

> AnalyticExperimentResponse CreateAnalyticExperimentDeclare(ctx, id).Execute()

Record the outcome: a winning arm, no difference, or inconclusive. Concludes the experiment.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The experiment id (a uuid; any other id is a 404).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.CreateAnalyticExperimentDeclare(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.CreateAnalyticExperimentDeclare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAnalyticExperimentDeclare`: AnalyticExperimentResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.CreateAnalyticExperimentDeclare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The experiment id (a uuid; any other id is a 404). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnalyticExperimentDeclareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**AnalyticExperimentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateAnalyticExperimentStop

> AnalyticExperimentResponse CreateAnalyticExperimentStop(ctx, id).Execute()

Stop an experiment: no more data is counted after now.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The experiment id (a uuid; any other id is a 404).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.CreateAnalyticExperimentStop(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.CreateAnalyticExperimentStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAnalyticExperimentStop`: AnalyticExperimentResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.CreateAnalyticExperimentStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The experiment id (a uuid; any other id is a 404). | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAnalyticExperimentStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**AnalyticExperimentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteAnalyticExperiment

> DeleteStoryResponse DeleteAnalyticExperiment(ctx, id).Execute()

Delete an experiment. The listings are untouched.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The experiment id (a uuid; any other id is a 404).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.DeleteAnalyticExperiment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.DeleteAnalyticExperiment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAnalyticExperiment`: DeleteStoryResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.DeleteAnalyticExperiment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The experiment id (a uuid; any other id is a 404). | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAnalyticExperimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteStoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


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

**GetAnalyticBookkeepingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticByPlatform

> GetAnalyticByPlatformResponse GetAnalyticByPlatform(ctx).Days(days).Tz(tz).Execute()

Sales, net sales, fees and net profit by platform for the last N days.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    days := int32(56) // int32 | Calendar days back, today included. Larger values are clamped to 365. (optional) (default to 30)
    tz := "tz_example" // string | IANA time zone for day boundaries. Default UTC. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticByPlatform(context.Background()).Days(days).Tz(tz).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticByPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticByPlatform`: GetAnalyticByPlatformResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticByPlatform`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticByPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** | Calendar days back, today included. Larger values are clamped to 365. | [default to 30]
 **tz** | **string** | IANA time zone for day boundaries. Default UTC. | 

### Return type

**GetAnalyticByPlatformResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticCompare

> GetAnalyticCompareResponse GetAnalyticCompare(ctx).Ids(ids).Align(align).Platform(platform).Tz(tz).Execute()

Compare 2-4 of your listings: columns of price, quality, channels, views, favorites, offers and outcome, engagement curves, the fields that differ, and observational tests of favorite rate, view-to-offer rate and views per day.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    ids := "ids_example" // string | Comma-separated listing ids, 2 to 4. The first is arm A.
    align := "align_example" // string | age = x is days since first listed (default); calendar = x is the date. (optional)
    platform := "platform_example" // string | Restrict engagement to one channel (platform id). (optional)
    tz := "tz_example" // string | IANA time zone. Default UTC. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticCompare(context.Background()).Ids(ids).Align(align).Platform(platform).Tz(tz).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticCompare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticCompare`: GetAnalyticCompareResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticCompare`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticCompareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ids** | **string** | Comma-separated listing ids, 2 to 4. The first is arm A. | 
 **align** | **string** | age &#x3D; x is days since first listed (default); calendar &#x3D; x is the date. | 
 **platform** | **string** | Restrict engagement to one channel (platform id). | 
 **tz** | **string** | IANA time zone. Default UTC. | 

### Return type

**GetAnalyticCompareResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**GetAnalyticDashboardResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticExperiment

> GetAnalyticExperimentResponse GetAnalyticExperiment(ctx).Status(status).Page(page).Size(size).Execute()

Your listing A/B experiments, newest first, with status, days running and the current verdict.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    status := "status_example" // string | Default all. (optional)
    page := int32(56) // int32 | 1-indexed page. Default 1. (optional)
    size := int32(56) // int32 | Rows per page. Default 20. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticExperiment(context.Background()).Status(status).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticExperiment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticExperiment`: GetAnalyticExperimentResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticExperiment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticExperimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Default all. | 
 **page** | **int32** | 1-indexed page. Default 1. | 
 **size** | **int32** | Rows per page. Default 20. | 

### Return type

**GetAnalyticExperimentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticExperimentById

> GetAnalyticExperimentByIdResponse GetAnalyticExperimentById(ctx, id).Execute()

One experiment: arms with their windows, exposure and rates, the primary-metric test per arm against A, guardrail tests, the daily series and what to do next.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The experiment id (a uuid).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticExperimentById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticExperimentById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticExperimentById`: GetAnalyticExperimentByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticExperimentById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The experiment id (a uuid). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticExperimentByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetAnalyticExperimentByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticInventory

> GetAnalyticInventoryResponse GetAnalyticInventory(ctx).Range_(range_).From(from).To(to).Compare(compare).CmpFrom(cmpFrom).CmpTo(cmpTo).Labels(labels).Platforms(platforms).Category(category).Brand(brand).Tz(tz).Maturity(maturity).Execute()

Unsold stock: aging buckets with cost, listed value and market value, the 20 listings most in need of attention, shelf value over time, cohort sell-through at the chosen maturity, listed vs sold velocity and data health.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    range_ := "range__example" // string | Reporting window. Default 30d. custom needs from and to. (optional)
    from := "from_example" // string | YYYY-MM-DD, inclusive, in tz. Only with range=custom. (optional)
    to := "to_example" // string | YYYY-MM-DD, inclusive, in tz. Only with range=custom. (optional)
    compare := "compare_example" // string | Comparison window for deltas: prev (same length just before), yoy, none, custom (cmpFrom/cmpTo). (optional)
    cmpFrom := "cmpFrom_example" // string | YYYY-MM-DD. Only with compare=custom. (optional)
    cmpTo := "cmpTo_example" // string | YYYY-MM-DD. Only with compare=custom. (optional)
    labels := "labels_example" // string | Comma-separated inventory labels (any of them). (optional)
    platforms := "platforms_example" // string | Comma-separated platform ids. (optional)
    category := "category_example" // string | Canonical main-category key. (optional)
    brand := "brand_example" // string | Brand name. (optional)
    tz := "tz_example" // string | IANA time zone for day boundaries. Default UTC. An invalid name is a 400. (optional)
    maturity := int32(56) // int32 | Days an item has to sell within to count as sold in the cohort. Default 30. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticInventory(context.Background()).Range_(range_).From(from).To(to).Compare(compare).CmpFrom(cmpFrom).CmpTo(cmpTo).Labels(labels).Platforms(platforms).Category(category).Brand(brand).Tz(tz).Maturity(maturity).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticInventory`: GetAnalyticInventoryResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticInventory`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticInventoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Reporting window. Default 30d. custom needs from and to. | 
 **from** | **string** | YYYY-MM-DD, inclusive, in tz. Only with range&#x3D;custom. | 
 **to** | **string** | YYYY-MM-DD, inclusive, in tz. Only with range&#x3D;custom. | 
 **compare** | **string** | Comparison window for deltas: prev (same length just before), yoy, none, custom (cmpFrom/cmpTo). | 
 **cmpFrom** | **string** | YYYY-MM-DD. Only with compare&#x3D;custom. | 
 **cmpTo** | **string** | YYYY-MM-DD. Only with compare&#x3D;custom. | 
 **labels** | **string** | Comma-separated inventory labels (any of them). | 
 **platforms** | **string** | Comma-separated platform ids. | 
 **category** | **string** | Canonical main-category key. | 
 **brand** | **string** | Brand name. | 
 **tz** | **string** | IANA time zone for day boundaries. Default UTC. An invalid name is a 400. | 
 **maturity** | **int32** | Days an item has to sell within to count as sold in the cohort. Default 30. | 

### Return type

**GetAnalyticInventoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**GetAnalyticItemResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticItemByListingId

> GetAnalyticItemByListingIdResponse GetAnalyticItemByListingId(ctx, listingId).By(by).Tz(tz).Execute()

One listing in depth: channels with views, favorites and offers, the merged timeline, engagement series, price history, comps, money waterfall, days on market, offers, messages, experiments and suggestions.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    listingId := "listingId_example" // string | A listings.id you own (or an inventory item id with by=inventory).
    by := "by_example" // string | inventory = the id is an inventory item id; the answer is { redirectListingId }. Default listing. (optional)
    tz := "tz_example" // string | IANA time zone for day boundaries. Default UTC. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticItemByListingId(context.Background(), listingId).By(by).Tz(tz).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticItemByListingId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticItemByListingId`: GetAnalyticItemByListingIdResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticItemByListingId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**listingId** | **string** | A listings.id you own (or an inventory item id with by&#x3D;inventory). | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticItemByListingIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **by** | **string** | inventory &#x3D; the id is an inventory item id; the answer is { redirectListingId }. Default listing. | 
 **tz** | **string** | IANA time zone for day boundaries. Default UTC. | 

### Return type

**GetAnalyticItemByListingIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticMarketplace

> GetAnalyticMarketplaceResponse GetAnalyticMarketplace(ctx).Range_(range_).From(from).To(to).Compare(compare).CmpFrom(cmpFrom).CmpTo(cmpTo).Labels(labels).Platforms(platforms).Category(category).Brand(brand).Tz(tz).Vs(vs).Sort(sort).Dir(dir).Execute()

Every marketplace side by side: net sales and profit, units, ASP, take rate, margin, shipping subsidy, days to sell, sell-through, cancel and return rates, offers and engagement; with vs=2-3 platforms, a head-to-head over listings live on all of them.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    range_ := "range__example" // string | Reporting window. Default 30d. custom needs from and to. (optional)
    from := "from_example" // string | YYYY-MM-DD, inclusive, in tz. Only with range=custom. (optional)
    to := "to_example" // string | YYYY-MM-DD, inclusive, in tz. Only with range=custom. (optional)
    compare := "compare_example" // string | Comparison window for deltas: prev (same length just before), yoy, none, custom (cmpFrom/cmpTo). (optional)
    cmpFrom := "cmpFrom_example" // string | YYYY-MM-DD. Only with compare=custom. (optional)
    cmpTo := "cmpTo_example" // string | YYYY-MM-DD. Only with compare=custom. (optional)
    labels := "labels_example" // string | Comma-separated inventory labels (any of them). (optional)
    platforms := "platforms_example" // string | Comma-separated platform ids. (optional)
    category := "category_example" // string | Canonical main-category key. (optional)
    brand := "brand_example" // string | Brand name. (optional)
    tz := "tz_example" // string | IANA time zone for day boundaries. Default UTC. An invalid name is a 400. (optional)
    vs := "vs_example" // string | Comma-separated 2-3 platform ids for the head-to-head. (optional)
    sort := "sort_example" // string | Default netSales. (optional)
    dir := "dir_example" // string | Default desc. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticMarketplace(context.Background()).Range_(range_).From(from).To(to).Compare(compare).CmpFrom(cmpFrom).CmpTo(cmpTo).Labels(labels).Platforms(platforms).Category(category).Brand(brand).Tz(tz).Vs(vs).Sort(sort).Dir(dir).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticMarketplace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticMarketplace`: GetAnalyticMarketplaceResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticMarketplace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticMarketplaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Reporting window. Default 30d. custom needs from and to. | 
 **from** | **string** | YYYY-MM-DD, inclusive, in tz. Only with range&#x3D;custom. | 
 **to** | **string** | YYYY-MM-DD, inclusive, in tz. Only with range&#x3D;custom. | 
 **compare** | **string** | Comparison window for deltas: prev (same length just before), yoy, none, custom (cmpFrom/cmpTo). | 
 **cmpFrom** | **string** | YYYY-MM-DD. Only with compare&#x3D;custom. | 
 **cmpTo** | **string** | YYYY-MM-DD. Only with compare&#x3D;custom. | 
 **labels** | **string** | Comma-separated inventory labels (any of them). | 
 **platforms** | **string** | Comma-separated platform ids. | 
 **category** | **string** | Canonical main-category key. | 
 **brand** | **string** | Brand name. | 
 **tz** | **string** | IANA time zone for day boundaries. Default UTC. An invalid name is a 400. | 
 **vs** | **string** | Comma-separated 2-3 platform ids for the head-to-head. | 
 **sort** | **string** | Default netSales. | 
 **dir** | **string** | Default desc. | 

### Return type

**GetAnalyticMarketplaceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticMarketplaceByPlatform

> GetAnalyticMarketplaceByPlatformResponse GetAnalyticMarketplaceByPlatform(ctx, platform).Range_(range_).From(from).To(to).Compare(compare).CmpFrom(cmpFrom).CmpTo(cmpTo).Labels(labels).Platforms(platforms).Category(category).Brand(brand).Tz(tz).Execute()

One marketplace in depth: its row, the money series, fee breakdown, days-to-sell and price-realization histograms, categories, top items, at-risk listings, a day-by-hour sales heatmap and offer stats.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    platform := "platform_example" // string | A platform id, e.g. ebay, poshmark.
    range_ := "range__example" // string | Reporting window. Default 30d. custom needs from and to. (optional)
    from := "from_example" // string | YYYY-MM-DD, inclusive, in tz. Only with range=custom. (optional)
    to := "to_example" // string | YYYY-MM-DD, inclusive, in tz. Only with range=custom. (optional)
    compare := "compare_example" // string | Comparison window for deltas: prev (same length just before), yoy, none, custom (cmpFrom/cmpTo). (optional)
    cmpFrom := "cmpFrom_example" // string | YYYY-MM-DD. Only with compare=custom. (optional)
    cmpTo := "cmpTo_example" // string | YYYY-MM-DD. Only with compare=custom. (optional)
    labels := "labels_example" // string | Comma-separated inventory labels (any of them). (optional)
    platforms := "platforms_example" // string | Comma-separated platform ids. (optional)
    category := "category_example" // string | Canonical main-category key. (optional)
    brand := "brand_example" // string | Brand name. (optional)
    tz := "tz_example" // string | IANA time zone for day boundaries. Default UTC. An invalid name is a 400. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticMarketplaceByPlatform(context.Background(), platform).Range_(range_).From(from).To(to).Compare(compare).CmpFrom(cmpFrom).CmpTo(cmpTo).Labels(labels).Platforms(platforms).Category(category).Brand(brand).Tz(tz).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticMarketplaceByPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticMarketplaceByPlatform`: GetAnalyticMarketplaceByPlatformResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticMarketplaceByPlatform`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** | A platform id, e.g. ebay, poshmark. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticMarketplaceByPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **range_** | **string** | Reporting window. Default 30d. custom needs from and to. | 
 **from** | **string** | YYYY-MM-DD, inclusive, in tz. Only with range&#x3D;custom. | 
 **to** | **string** | YYYY-MM-DD, inclusive, in tz. Only with range&#x3D;custom. | 
 **compare** | **string** | Comparison window for deltas: prev (same length just before), yoy, none, custom (cmpFrom/cmpTo). | 
 **cmpFrom** | **string** | YYYY-MM-DD. Only with compare&#x3D;custom. | 
 **cmpTo** | **string** | YYYY-MM-DD. Only with compare&#x3D;custom. | 
 **labels** | **string** | Comma-separated inventory labels (any of them). | 
 **platforms** | **string** | Comma-separated platform ids. | 
 **category** | **string** | Canonical main-category key. | 
 **brand** | **string** | Brand name. | 
 **tz** | **string** | IANA time zone for day boundaries. Default UTC. An invalid name is a 400. | 

### Return type

**GetAnalyticMarketplaceByPlatformResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticOverview

> GetAnalyticOverviewResponse GetAnalyticOverview(ctx).Range_(range_).From(from).To(to).Compare(compare).CmpFrom(cmpFrom).CmpTo(cmpTo).Labels(labels).Platforms(platforms).Category(category).Brand(brand).Tz(tz).Stores(stores).Execute()

KPIs with deltas (net sales, net profit, units, ASP, margin, sell-through, median days to sell, active listings, take rate), the trend series, by-marketplace split, top categories, brands and items, sell-through cohort, data health and insights.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    range_ := "range__example" // string | Reporting window. Default 30d. custom needs from and to. (optional)
    from := "from_example" // string | YYYY-MM-DD, inclusive, in tz. Only with range=custom. (optional)
    to := "to_example" // string | YYYY-MM-DD, inclusive, in tz. Only with range=custom. (optional)
    compare := "compare_example" // string | Comparison window for deltas: prev (same length just before), yoy, none, custom (cmpFrom/cmpTo). (optional)
    cmpFrom := "cmpFrom_example" // string | YYYY-MM-DD. Only with compare=custom. (optional)
    cmpTo := "cmpTo_example" // string | YYYY-MM-DD. Only with compare=custom. (optional)
    labels := "labels_example" // string | Comma-separated inventory labels (any of them). (optional)
    platforms := "platforms_example" // string | Comma-separated platform ids. (optional)
    category := "category_example" // string | Canonical main-category key. (optional)
    brand := "brand_example" // string | Brand name. (optional)
    tz := "tz_example" // string | IANA time zone for day boundaries. Default UTC. An invalid name is a 400. (optional)
    stores := "stores_example" // string | Comma-separated store ids. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticOverview(context.Background()).Range_(range_).From(from).To(to).Compare(compare).CmpFrom(cmpFrom).CmpTo(cmpTo).Labels(labels).Platforms(platforms).Category(category).Brand(brand).Tz(tz).Stores(stores).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticOverview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticOverview`: GetAnalyticOverviewResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticOverview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticOverviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Reporting window. Default 30d. custom needs from and to. | 
 **from** | **string** | YYYY-MM-DD, inclusive, in tz. Only with range&#x3D;custom. | 
 **to** | **string** | YYYY-MM-DD, inclusive, in tz. Only with range&#x3D;custom. | 
 **compare** | **string** | Comparison window for deltas: prev (same length just before), yoy, none, custom (cmpFrom/cmpTo). | 
 **cmpFrom** | **string** | YYYY-MM-DD. Only with compare&#x3D;custom. | 
 **cmpTo** | **string** | YYYY-MM-DD. Only with compare&#x3D;custom. | 
 **labels** | **string** | Comma-separated inventory labels (any of them). | 
 **platforms** | **string** | Comma-separated platform ids. | 
 **category** | **string** | Canonical main-category key. | 
 **brand** | **string** | Brand name. | 
 **tz** | **string** | IANA time zone for day boundaries. Default UTC. An invalid name is a 400. | 
 **stores** | **string** | Comma-separated store ids. | 

### Return type

**GetAnalyticOverviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticPerformance

> GetAnalyticPerformanceResponse GetAnalyticPerformance(ctx).Range_(range_).From(from).To(to).Compare(compare).CmpFrom(cmpFrom).CmpTo(cmpTo).Labels(labels).Platforms(platforms).Category(category).Brand(brand).Tz(tz).Status(status).Q(q).Sort(sort).Dir(dir).Page(page).Size(size).MinDays(minDays).Execute()

Per-listing performance: channels, views, favorites, offers and conversations in the period, units sold, net sales and profit, margin, days to sell and a 0-100 health score, paged and sortable.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    range_ := "range__example" // string | Reporting window. Default 30d. custom needs from and to. (optional)
    from := "from_example" // string | YYYY-MM-DD, inclusive, in tz. Only with range=custom. (optional)
    to := "to_example" // string | YYYY-MM-DD, inclusive, in tz. Only with range=custom. (optional)
    compare := "compare_example" // string | Comparison window for deltas: prev (same length just before), yoy, none, custom (cmpFrom/cmpTo). (optional)
    cmpFrom := "cmpFrom_example" // string | YYYY-MM-DD. Only with compare=custom. (optional)
    cmpTo := "cmpTo_example" // string | YYYY-MM-DD. Only with compare=custom. (optional)
    labels := "labels_example" // string | Comma-separated inventory labels (any of them). (optional)
    platforms := "platforms_example" // string | Comma-separated platform ids. (optional)
    category := "category_example" // string | Canonical main-category key. (optional)
    brand := "brand_example" // string | Brand name. (optional)
    tz := "tz_example" // string | IANA time zone for day boundaries. Default UTC. An invalid name is a 400. (optional)
    status := "status_example" // string | unsold = listed at least minDays with no completed sale. (optional)
    q := "q_example" // string | Title, SKU or brand contains. (optional)
    sort := "sort_example" // string | Default lastActivity. (optional)
    dir := "dir_example" // string | Default desc. (optional)
    page := int32(56) // int32 | 1-indexed page. Default 1. (optional)
    size := int32(56) // int32 | Rows per page. Default 50. (optional)
    minDays := int32(56) // int32 | Minimum days listed (for status=unsold). Default 0. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticPerformance(context.Background()).Range_(range_).From(from).To(to).Compare(compare).CmpFrom(cmpFrom).CmpTo(cmpTo).Labels(labels).Platforms(platforms).Category(category).Brand(brand).Tz(tz).Status(status).Q(q).Sort(sort).Dir(dir).Page(page).Size(size).MinDays(minDays).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticPerformance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticPerformance`: GetAnalyticPerformanceResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticPerformance`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticPerformanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Reporting window. Default 30d. custom needs from and to. | 
 **from** | **string** | YYYY-MM-DD, inclusive, in tz. Only with range&#x3D;custom. | 
 **to** | **string** | YYYY-MM-DD, inclusive, in tz. Only with range&#x3D;custom. | 
 **compare** | **string** | Comparison window for deltas: prev (same length just before), yoy, none, custom (cmpFrom/cmpTo). | 
 **cmpFrom** | **string** | YYYY-MM-DD. Only with compare&#x3D;custom. | 
 **cmpTo** | **string** | YYYY-MM-DD. Only with compare&#x3D;custom. | 
 **labels** | **string** | Comma-separated inventory labels (any of them). | 
 **platforms** | **string** | Comma-separated platform ids. | 
 **category** | **string** | Canonical main-category key. | 
 **brand** | **string** | Brand name. | 
 **tz** | **string** | IANA time zone for day boundaries. Default UTC. An invalid name is a 400. | 
 **status** | **string** | unsold &#x3D; listed at least minDays with no completed sale. | 
 **q** | **string** | Title, SKU or brand contains. | 
 **sort** | **string** | Default lastActivity. | 
 **dir** | **string** | Default desc. | 
 **page** | **int32** | 1-indexed page. Default 1. | 
 **size** | **int32** | Rows per page. Default 50. | 
 **minDays** | **int32** | Minimum days listed (for status&#x3D;unsold). Default 0. | 

### Return type

**GetAnalyticPerformanceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticProgram

> GetAnalyticProgramResponse GetAnalyticProgram(ctx, program).From(from).To(to).By(by).Execute()

One program's report over a date range (default the last 30 days, at most 366): KPIs, a daily trend and a breakdown table. Money in integer cents.

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
    program := "program_example" // string | 
    from := time.Now() // string | First UTC day (YYYY-MM-DD), inclusive. (optional)
    to := time.Now() // string | Last UTC day (YYYY-MM-DD), inclusive. Default today. (optional)
    by := "by_example" // string | Breakdown dimension, when the program offers one (worlds: showroom, property, event). (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticProgram(context.Background(), program).From(from).To(to).By(by).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticProgram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticProgram`: GetAnalyticProgramResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticProgram`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**program** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticProgramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **string** | First UTC day (YYYY-MM-DD), inclusive. | 
 **to** | **string** | Last UTC day (YYYY-MM-DD), inclusive. Default today. | 
 **by** | **string** | Breakdown dimension, when the program offers one (worlds: showroom, property, event). | 

### Return type

**GetAnalyticProgramResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticSegment

> GetAnalyticSegmentResponse GetAnalyticSegment(ctx).Range_(range_).From(from).To(to).Compare(compare).CmpFrom(cmpFrom).CmpTo(cmpTo).Labels(labels).Platforms(platforms).Category(category).Brand(brand).Tz(tz).Dim(dim).Sort(sort).Dir(dir).MinN(minN).Q(q).Execute()

Niches along one dimension: listed and active counts, units sold, cohort sell-through, median days to sell, net sales and profit, margin, ROI, ASP and discount from ask; rows under minN are flagged lowSample.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    range_ := "range__example" // string | Reporting window. Default 30d. custom needs from and to. (optional)
    from := "from_example" // string | YYYY-MM-DD, inclusive, in tz. Only with range=custom. (optional)
    to := "to_example" // string | YYYY-MM-DD, inclusive, in tz. Only with range=custom. (optional)
    compare := "compare_example" // string | Comparison window for deltas: prev (same length just before), yoy, none, custom (cmpFrom/cmpTo). (optional)
    cmpFrom := "cmpFrom_example" // string | YYYY-MM-DD. Only with compare=custom. (optional)
    cmpTo := "cmpTo_example" // string | YYYY-MM-DD. Only with compare=custom. (optional)
    labels := "labels_example" // string | Comma-separated inventory labels (any of them). (optional)
    platforms := "platforms_example" // string | Comma-separated platform ids. (optional)
    category := "category_example" // string | Canonical main-category key. (optional)
    brand := "brand_example" // string | Brand name. (optional)
    tz := "tz_example" // string | IANA time zone for day boundaries. Default UTC. An invalid name is a 400. (optional)
    dim := "dim_example" // string | category = main category; subcategory = main > sub; brand; condition; size; source = where you acquired it; department; priceBand = ask in <$10, $10-25, $25-50, $50-100, $100-250, $250+; ageBand = days live at sale or now (<7, 7-30, 30-90, 90+). (optional)
    sort := "sort_example" // string | Default netSales. (optional)
    dir := "dir_example" // string | Default desc. (optional)
    minN := int32(56) // int32 | Minimum sample for a row to count as a winner or loser. Default 3. (optional)
    q := "q_example" // string | Segment label contains. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticSegment(context.Background()).Range_(range_).From(from).To(to).Compare(compare).CmpFrom(cmpFrom).CmpTo(cmpTo).Labels(labels).Platforms(platforms).Category(category).Brand(brand).Tz(tz).Dim(dim).Sort(sort).Dir(dir).MinN(minN).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticSegment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticSegment`: GetAnalyticSegmentResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticSegment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticSegmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **range_** | **string** | Reporting window. Default 30d. custom needs from and to. | 
 **from** | **string** | YYYY-MM-DD, inclusive, in tz. Only with range&#x3D;custom. | 
 **to** | **string** | YYYY-MM-DD, inclusive, in tz. Only with range&#x3D;custom. | 
 **compare** | **string** | Comparison window for deltas: prev (same length just before), yoy, none, custom (cmpFrom/cmpTo). | 
 **cmpFrom** | **string** | YYYY-MM-DD. Only with compare&#x3D;custom. | 
 **cmpTo** | **string** | YYYY-MM-DD. Only with compare&#x3D;custom. | 
 **labels** | **string** | Comma-separated inventory labels (any of them). | 
 **platforms** | **string** | Comma-separated platform ids. | 
 **category** | **string** | Canonical main-category key. | 
 **brand** | **string** | Brand name. | 
 **tz** | **string** | IANA time zone for day boundaries. Default UTC. An invalid name is a 400. | 
 **dim** | **string** | category &#x3D; main category; subcategory &#x3D; main &gt; sub; brand; condition; size; source &#x3D; where you acquired it; department; priceBand &#x3D; ask in &lt;$10, $10-25, $25-50, $50-100, $100-250, $250+; ageBand &#x3D; days live at sale or now (&lt;7, 7-30, 30-90, 90+). | 
 **sort** | **string** | Default netSales. | 
 **dir** | **string** | Default desc. | 
 **minN** | **int32** | Minimum sample for a row to count as a winner or loser. Default 3. | 
 **q** | **string** | Segment label contains. | 

### Return type

**GetAnalyticSegmentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticSegmentByKey

> GetAnalyticSegmentByKeyResponse GetAnalyticSegmentByKey(ctx, dim, key).Range_(range_).From(from).To(to).Compare(compare).CmpFrom(cmpFrom).CmpTo(cmpTo).Labels(labels).Platforms(platforms).Category(category).Brand(brand).Tz(tz).Execute()

One niche in depth: its row, series, by-marketplace split, days-to-sell histogram, price bands, top items and the 10 oldest unsold listings in it.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    dim := "dim_example" // string | One of category, subcategory, brand, condition, size, source, department, priceBand, ageBand. category = main category; subcategory = main > sub; brand; condition; size; source = where you acquired it; department; priceBand = ask in <$10, $10-25, $25-50, $50-100, $100-250, $250+; ageBand = days live at sale or now (<7, 7-30, 30-90, 90+).
    key := "key_example" // string | The segment key from a segments row (URL-encoded). A key with no data in scope is a 404.
    range_ := "range__example" // string | Reporting window. Default 30d. custom needs from and to. (optional)
    from := "from_example" // string | YYYY-MM-DD, inclusive, in tz. Only with range=custom. (optional)
    to := "to_example" // string | YYYY-MM-DD, inclusive, in tz. Only with range=custom. (optional)
    compare := "compare_example" // string | Comparison window for deltas: prev (same length just before), yoy, none, custom (cmpFrom/cmpTo). (optional)
    cmpFrom := "cmpFrom_example" // string | YYYY-MM-DD. Only with compare=custom. (optional)
    cmpTo := "cmpTo_example" // string | YYYY-MM-DD. Only with compare=custom. (optional)
    labels := "labels_example" // string | Comma-separated inventory labels (any of them). (optional)
    platforms := "platforms_example" // string | Comma-separated platform ids. (optional)
    category := "category_example" // string | Canonical main-category key. (optional)
    brand := "brand_example" // string | Brand name. (optional)
    tz := "tz_example" // string | IANA time zone for day boundaries. Default UTC. An invalid name is a 400. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticSegmentByKey(context.Background(), dim, key).Range_(range_).From(from).To(to).Compare(compare).CmpFrom(cmpFrom).CmpTo(cmpTo).Labels(labels).Platforms(platforms).Category(category).Brand(brand).Tz(tz).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticSegmentByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticSegmentByKey`: GetAnalyticSegmentByKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticSegmentByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**dim** | **string** | One of category, subcategory, brand, condition, size, source, department, priceBand, ageBand. category &#x3D; main category; subcategory &#x3D; main &gt; sub; brand; condition; size; source &#x3D; where you acquired it; department; priceBand &#x3D; ask in &lt;$10, $10-25, $25-50, $50-100, $100-250, $250+; ageBand &#x3D; days live at sale or now (&lt;7, 7-30, 30-90, 90+). | 
**key** | **string** | The segment key from a segments row (URL-encoded). A key with no data in scope is a 404. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticSegmentByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **range_** | **string** | Reporting window. Default 30d. custom needs from and to. | 
 **from** | **string** | YYYY-MM-DD, inclusive, in tz. Only with range&#x3D;custom. | 
 **to** | **string** | YYYY-MM-DD, inclusive, in tz. Only with range&#x3D;custom. | 
 **compare** | **string** | Comparison window for deltas: prev (same length just before), yoy, none, custom (cmpFrom/cmpTo). | 
 **cmpFrom** | **string** | YYYY-MM-DD. Only with compare&#x3D;custom. | 
 **cmpTo** | **string** | YYYY-MM-DD. Only with compare&#x3D;custom. | 
 **labels** | **string** | Comma-separated inventory labels (any of them). | 
 **platforms** | **string** | Comma-separated platform ids. | 
 **category** | **string** | Canonical main-category key. | 
 **brand** | **string** | Brand name. | 
 **tz** | **string** | IANA time zone for day boundaries. Default UTC. An invalid name is a 400. | 

### Return type

**GetAnalyticSegmentByKeyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticSummary

> GetAnalyticSummaryResponse GetAnalyticSummary(ctx).Days(days).Tz(tz).Execute()

Headline KPIs for the last N days (the same numbers as the analytics overview).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    days := int32(56) // int32 | Calendar days back, today included. Larger values are clamped to 365. (optional) (default to 30)
    tz := "tz_example" // string | IANA time zone for day boundaries. Default UTC. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticSummary(context.Background()).Days(days).Tz(tz).Execute()
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
 **days** | **int32** | Calendar days back, today included. Larger values are clamped to 365. | [default to 30]
 **tz** | **string** | IANA time zone for day boundaries. Default UTC. | 

### Return type

**GetAnalyticSummaryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAnalyticTimesery

> GetAnalyticTimeseryResponse GetAnalyticTimesery(ctx).Days(days).Tz(tz).Execute()

Sales and net sales over the last N days, zero-filled, one row per calendar day.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    days := int32(56) // int32 | Calendar days back, today included. Larger values are clamped to 365. (optional) (default to 30)
    tz := "tz_example" // string | IANA time zone for day boundaries. Default UTC. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.GetAnalyticTimesery(context.Background()).Days(days).Tz(tz).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.GetAnalyticTimesery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAnalyticTimesery`: GetAnalyticTimeseryResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.GetAnalyticTimesery`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAnalyticTimeseryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** | Calendar days back, today included. Larger values are clamped to 365. | [default to 30]
 **tz** | **string** | IANA time zone for day boundaries. Default UTC. | 

### Return type

**GetAnalyticTimeseryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**GetAnalyticTodayResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListInsightByPlatform

> ListInsightByPlatformResponse ListInsightByPlatform(ctx).Tz(tz).Execute()

Platform velocity + margin insight (90-day window): units, net sales, median days to sell, ASP and net after the real marketplace fees.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tz := "tz_example" // string | IANA time zone for day boundaries. Default UTC. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.ListInsightByPlatform(context.Background()).Tz(tz).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.ListInsightByPlatform``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInsightByPlatform`: ListInsightByPlatformResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.ListInsightByPlatform`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInsightByPlatformRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **tz** | **string** | IANA time zone for day boundaries. Default UTC. | 

### Return type

**ListInsightByPlatformResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateAnalyticExperiment

> AnalyticExperimentResponse UpdateAnalyticExperiment(ctx, id).Execute()

Rename an experiment or change its hypothesis, notes or planned end.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    id := "id_example" // string | The experiment id (a uuid; any other id is a 404).

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AnalyticsApi.UpdateAnalyticExperiment(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AnalyticsApi.UpdateAnalyticExperiment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAnalyticExperiment`: AnalyticExperimentResponse
    fmt.Fprintf(os.Stdout, "Response from `AnalyticsApi.UpdateAnalyticExperiment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The experiment id (a uuid; any other id is a 404). | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAnalyticExperimentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**AnalyticExperimentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

