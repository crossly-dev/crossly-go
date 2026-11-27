# \WorldsAdsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetWorldExperienceAd**](WorldsAdsApi.md#GetWorldExperienceAd) | **Get** /v1/worlds/experiences/{experienceId}/ads | The ads for a game&#39;s in-world ad surfaces (parts tagged CrosslyAdSurface), for the token&#39;s owner: Sponsored Listings labelled \&quot;Ad\&quot;, served only to signed-in adults where the game&#39;s creator turned ads on, never to its team. A click is POST /api/public/shop/ads/click { token }.
[**GetWorldExperienceAdReport**](WorldsAdsApi.md#GetWorldExperienceAdReport) | **Get** /v1/worlds/experiences/{experienceId}/ads/report | In-game ads over the last N days (default 30, max 90): impressions, clicks, charged clicks, advertiser spend and your earned share, per day.
[**GetWorldExperienceAdSetting**](WorldsAdsApi.md#GetWorldExperienceAdSetting) | **Get** /v1/worlds/experiences/{experienceId}/ad-settings | Whether this game shows in-game ads (opt-in, default off) and the creator&#39;s share of Crossly&#39;s ad revenue in force.
[**UpdateWorldExperienceAdSetting**](WorldsAdsApi.md#UpdateWorldExperienceAdSetting) | **Put** /v1/worlds/experiences/{experienceId}/ad-settings | Turn in-game ads on or off for this game. Ads fill the game&#39;s ad surfaces (parts tagged CrosslyAdSurface); nothing shows until it has one.



## GetWorldExperienceAd

> GetWorldExperienceAdResponse GetWorldExperienceAd(ctx, experienceId).Execute()

The ads for a game's in-world ad surfaces (parts tagged CrosslyAdSurface), for the token's owner: Sponsored Listings labelled \"Ad\", served only to signed-in adults where the game's creator turned ads on, never to its team. A click is POST /api/public/shop/ads/click { token }.

### Example

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
    resp, r, err := apiClient.WorldsAdsApi.GetWorldExperienceAd(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsAdsApi.GetWorldExperienceAd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceAd`: GetWorldExperienceAdResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsAdsApi.GetWorldExperienceAd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceAdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceAdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceAdReport

> GetWorldExperienceAdReportResponse GetWorldExperienceAdReport(ctx, experienceId).Execute()

In-game ads over the last N days (default 30, max 90): impressions, clicks, charged clicks, advertiser spend and your earned share, per day.

### Example

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
    resp, r, err := apiClient.WorldsAdsApi.GetWorldExperienceAdReport(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsAdsApi.GetWorldExperienceAdReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceAdReport`: GetWorldExperienceAdReportResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsAdsApi.GetWorldExperienceAdReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceAdReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceAdReportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceAdSetting

> GetWorldExperienceAdSettingResponse GetWorldExperienceAdSetting(ctx, experienceId).Execute()

Whether this game shows in-game ads (opt-in, default off) and the creator's share of Crossly's ad revenue in force.

### Example

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
    resp, r, err := apiClient.WorldsAdsApi.GetWorldExperienceAdSetting(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsAdsApi.GetWorldExperienceAdSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceAdSetting`: GetWorldExperienceAdSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsAdsApi.GetWorldExperienceAdSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceAdSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceAdSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceAdSetting

> GetWorldExperienceAdSettingResponse UpdateWorldExperienceAdSetting(ctx, experienceId).Execute()

Turn in-game ads on or off for this game. Ads fill the game's ad surfaces (parts tagged CrosslyAdSurface); nothing shows until it has one.

### Example

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
    resp, r, err := apiClient.WorldsAdsApi.UpdateWorldExperienceAdSetting(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsAdsApi.UpdateWorldExperienceAdSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceAdSetting`: GetWorldExperienceAdSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsAdsApi.UpdateWorldExperienceAdSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceAdSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceAdSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

