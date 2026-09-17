# \PayoutApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetPayoutEstimate**](PayoutApi.md#GetPayoutEstimate) | **Get** /v1/payout/estimate | What one platform nets at a given price, after fees and shipping.
[**GetPayoutGrossForNet**](PayoutApi.md#GetPayoutGrossForNet) | **Get** /v1/payout/gross-for-net | The gross price needed to clear a target net on one platform.
[**ListPayoutCompare**](PayoutApi.md#ListPayoutCompare) | **Get** /v1/payout/compare | Rank platforms by what they net at a given price. Defaults to connected ones.



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

[**GetPayoutEstimateResponse**](GetPayoutEstimateResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
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

[**GetPayoutGrossForNetResponse**](GetPayoutGrossForNetResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPayoutCompare

> V1List ListPayoutCompare(ctx).Execute()

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
    // response from `ListPayoutCompare`: V1List
    fmt.Fprintf(os.Stdout, "Response from `PayoutApi.ListPayoutCompare`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPayoutCompareRequest struct via the builder pattern


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

