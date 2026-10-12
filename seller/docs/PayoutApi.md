# \PayoutApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePayoutConnectOnboard**](PayoutApi.md#CreatePayoutConnectOnboard) | **Post** /v1/payout/connect/onboard | Start or resume payout-account onboarding; returns a Stripe-hosted URL.
[**GetPayoutConnect**](PayoutApi.md#GetPayoutConnect) | **Get** /v1/payout/connect | Your Stripe payout account: status, country, agreement and payout currency.
[**GetPayoutEstimate**](PayoutApi.md#GetPayoutEstimate) | **Get** /v1/payout/estimate | What one platform nets at a given price, after fees and shipping.
[**GetPayoutGrossForNet**](PayoutApi.md#GetPayoutGrossForNet) | **Get** /v1/payout/gross-for-net | The gross price needed to clear a target net on one platform.
[**ListPayoutCompare**](PayoutApi.md#ListPayoutCompare) | **Get** /v1/payout/compare | Rank platforms by what they net at a given price. Defaults to connected ones.



## CreatePayoutConnectOnboard

> CreatePayoutConnectOnboardResponse CreatePayoutConnectOnboard(ctx).CreatePayoutConnectOnboardRequest(createPayoutConnectOnboardRequest).Execute()

Start or resume payout-account onboarding; returns a Stripe-hosted URL.



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
    createPayoutConnectOnboardRequest := *openapiclient.NewCreatePayoutConnectOnboardRequest() // CreatePayoutConnectOnboardRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PayoutApi.CreatePayoutConnectOnboard(context.Background()).CreatePayoutConnectOnboardRequest(createPayoutConnectOnboardRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.CreatePayoutConnectOnboard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePayoutConnectOnboard`: CreatePayoutConnectOnboardResponse
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.CreatePayoutConnectOnboard`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePayoutConnectOnboardRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createPayoutConnectOnboardRequest** | **CreatePayoutConnectOnboardRequest** |  | 

### Return type

**CreatePayoutConnectOnboardResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetPayoutConnect

> GetPayoutConnectResponse GetPayoutConnect(ctx).Execute()

Your Stripe payout account: status, country, agreement and payout currency.



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
    resp, r, err := apiClient.PayoutApi.GetPayoutConnect(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.GetPayoutConnect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayoutConnect`: GetPayoutConnectResponse
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.GetPayoutConnect`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutConnectRequest struct via the builder pattern


### Return type

**GetPayoutConnectResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetPayoutEstimate

> GetPayoutEstimateResponse GetPayoutEstimate(ctx).Execute()

What one platform nets at a given price, after fees and shipping.

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
    resp, r, err := apiClient.PayoutApi.GetPayoutEstimate(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.GetPayoutEstimate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayoutEstimate`: GetPayoutEstimateResponse
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.GetPayoutEstimate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutEstimateRequest struct via the builder pattern


### Return type

**GetPayoutEstimateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetPayoutGrossForNet

> GetPayoutGrossForNetResponse GetPayoutGrossForNet(ctx).Execute()

The gross price needed to clear a target net on one platform.

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
    resp, r, err := apiClient.PayoutApi.GetPayoutGrossForNet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.GetPayoutGrossForNet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPayoutGrossForNet`: GetPayoutGrossForNetResponse
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.GetPayoutGrossForNet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPayoutGrossForNetRequest struct via the builder pattern


### Return type

**GetPayoutGrossForNetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListPayoutCompare

> ListPayoutCompareResponse ListPayoutCompare(ctx).Execute()

Rank platforms by what they net at a given price. Defaults to connected ones.

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
    resp, r, err := apiClient.PayoutApi.ListPayoutCompare(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PayoutApi.ListPayoutCompare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPayoutCompare`: ListPayoutCompareResponse
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.ListPayoutCompare`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPayoutCompareRequest struct via the builder pattern


### Return type

**ListPayoutCompareResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

