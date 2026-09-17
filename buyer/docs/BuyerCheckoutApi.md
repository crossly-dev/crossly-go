# \BuyerCheckoutApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBuyerCheckout**](BuyerCheckoutApi.md#CreateBuyerCheckout) | **Post** /v1/buyer/checkout | Buy a listing without being present.
[**GetBuyerCheckoutControl**](BuyerCheckoutApi.md#GetBuyerCheckoutControl) | **Get** /v1/buyer/checkout/controls | What this key is allowed to spend.
[**UpdateBuyerCheckoutControl**](BuyerCheckoutApi.md#UpdateBuyerCheckoutControl) | **Put** /v1/buyer/checkout/controls | Switch this key on for spending, and set its limits.



## CreateBuyerCheckout

> CreateBuyerCheckoutResponse CreateBuyerCheckout(ctx).Execute()

Buy a listing without being present.



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
    resp, r, err := apiClient.BuyerCheckoutApi.CreateBuyerCheckout(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCheckoutApi.CreateBuyerCheckout``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBuyerCheckout`: CreateBuyerCheckoutResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCheckoutApi.CreateBuyerCheckout`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBuyerCheckoutRequest struct via the builder pattern


### Return type

[**CreateBuyerCheckoutResponse**](CreateBuyerCheckoutResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBuyerCheckoutControl

> GetBuyerCheckoutControlResponse GetBuyerCheckoutControl(ctx).Execute()

What this key is allowed to spend.



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
    resp, r, err := apiClient.BuyerCheckoutApi.GetBuyerCheckoutControl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCheckoutApi.GetBuyerCheckoutControl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBuyerCheckoutControl`: GetBuyerCheckoutControlResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCheckoutApi.GetBuyerCheckoutControl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBuyerCheckoutControlRequest struct via the builder pattern


### Return type

[**GetBuyerCheckoutControlResponse**](GetBuyerCheckoutControlResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBuyerCheckoutControl

> UpdateBuyerCheckoutControlResponse UpdateBuyerCheckoutControl(ctx).Execute()

Switch this key on for spending, and set its limits.



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
    resp, r, err := apiClient.BuyerCheckoutApi.UpdateBuyerCheckoutControl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BuyerCheckoutApi.UpdateBuyerCheckoutControl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBuyerCheckoutControl`: UpdateBuyerCheckoutControlResponse
    fmt.Fprintf(os.Stdout, "Response from `BuyerCheckoutApi.UpdateBuyerCheckoutControl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBuyerCheckoutControlRequest struct via the builder pattern


### Return type

[**UpdateBuyerCheckoutControlResponse**](UpdateBuyerCheckoutControlResponse.md)

### Authorization

[BuyerOAuth](../README.md#BuyerOAuth)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

