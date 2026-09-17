# \AdsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAdOffsiteCampaign**](AdsApi.md#CreateAdOffsiteCampaign) | **Post** /v1/ads/offsite/campaigns | Launch an offsite campaign for an item.
[**CreateAdOffsiteResume**](AdsApi.md#CreateAdOffsiteResume) | **Post** /v1/ads/offsite/resume | Clear an auto-pause and resume offsite spend.
[**GetAdOffsite**](AdsApi.md#GetAdOffsite) | **Get** /v1/ads/offsite | Your offsite-ads opt-in and its terms.
[**GetAdOffsiteEligibility**](AdsApi.md#GetAdOffsiteEligibility) | **Get** /v1/ads/offsite/eligibility | Whether an item can run offsite, and why not.
[**GetAdOffsiteReport**](AdsApi.md#GetAdOffsiteReport) | **Get** /v1/ads/offsite/report | What your offsite budget bought — including the misses.
[**UpdateAdOffsite**](AdsApi.md#UpdateAdOffsite) | **Put** /v1/ads/offsite | Turn offsite ads on or off. Yours alone to set.



## CreateAdOffsiteCampaign

> CreateAdOffsiteCampaignResponse CreateAdOffsiteCampaign(ctx).Execute()

Launch an offsite campaign for an item.



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
    resp, r, err := apiClient.AdsApi.CreateAdOffsiteCampaign(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.CreateAdOffsiteCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAdOffsiteCampaign`: CreateAdOffsiteCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.CreateAdOffsiteCampaign`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdOffsiteCampaignRequest struct via the builder pattern


### Return type

[**CreateAdOffsiteCampaignResponse**](CreateAdOffsiteCampaignResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAdOffsiteResume

> CreateAdOffsiteResumeResponse CreateAdOffsiteResume(ctx).Execute()

Clear an auto-pause and resume offsite spend.



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
    resp, r, err := apiClient.AdsApi.CreateAdOffsiteResume(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.CreateAdOffsiteResume``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAdOffsiteResume`: CreateAdOffsiteResumeResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.CreateAdOffsiteResume`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdOffsiteResumeRequest struct via the builder pattern


### Return type

[**CreateAdOffsiteResumeResponse**](CreateAdOffsiteResumeResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdOffsite

> GetAdOffsiteResponse GetAdOffsite(ctx).Execute()

Your offsite-ads opt-in and its terms.



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
    resp, r, err := apiClient.AdsApi.GetAdOffsite(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetAdOffsite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdOffsite`: GetAdOffsiteResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetAdOffsite`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdOffsiteRequest struct via the builder pattern


### Return type

[**GetAdOffsiteResponse**](GetAdOffsiteResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdOffsiteEligibility

> GetAdOffsiteEligibilityResponse GetAdOffsiteEligibility(ctx).Execute()

Whether an item can run offsite, and why not.



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
    resp, r, err := apiClient.AdsApi.GetAdOffsiteEligibility(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetAdOffsiteEligibility``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdOffsiteEligibility`: GetAdOffsiteEligibilityResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetAdOffsiteEligibility`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdOffsiteEligibilityRequest struct via the builder pattern


### Return type

[**GetAdOffsiteEligibilityResponse**](GetAdOffsiteEligibilityResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAdOffsiteReport

> GetAdOffsiteReportResponse GetAdOffsiteReport(ctx).Execute()

What your offsite budget bought — including the misses.



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
    resp, r, err := apiClient.AdsApi.GetAdOffsiteReport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.GetAdOffsiteReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAdOffsiteReport`: GetAdOffsiteReportResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.GetAdOffsiteReport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAdOffsiteReportRequest struct via the builder pattern


### Return type

[**GetAdOffsiteReportResponse**](GetAdOffsiteReportResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAdOffsite

> UpdateAdOffsiteResponse UpdateAdOffsite(ctx).Execute()

Turn offsite ads on or off. Yours alone to set.



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
    resp, r, err := apiClient.AdsApi.UpdateAdOffsite(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdsApi.UpdateAdOffsite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAdOffsite`: UpdateAdOffsiteResponse
    fmt.Fprintf(os.Stdout, "Response from `AdsApi.UpdateAdOffsite`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAdOffsiteRequest struct via the builder pattern


### Return type

[**UpdateAdOffsiteResponse**](UpdateAdOffsiteResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

