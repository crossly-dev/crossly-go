# \BillingApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingUpgrade**](BillingApi.md#CreateBillingUpgrade) | **Post** /v1/billing/upgrade | Start an upgrade to a higher plan.
[**ListBillingAppAllowances**](BillingApi.md#ListBillingAppAllowances) | **Get** /v1/billing/app-allowances | Your listing allowances through OAuth apps (used of limit, per app, with its platform scope).
[**ListBillingPlans**](BillingApi.md#ListBillingPlans) | **Get** /v1/billing/plans | List the plans with their prices, limits and features.



## CreateBillingUpgrade

> CreateBillingUpgradeResponse CreateBillingUpgrade(ctx).CreateBillingUpgradeRequest(createBillingUpgradeRequest).Execute()

Start an upgrade to a higher plan.



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
    createBillingUpgradeRequest := *openapiclient.NewCreateBillingUpgradeRequest("Tier_example") // CreateBillingUpgradeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BillingApi.CreateBillingUpgrade(context.Background()).CreateBillingUpgradeRequest(createBillingUpgradeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.CreateBillingUpgrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBillingUpgrade`: CreateBillingUpgradeResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.CreateBillingUpgrade`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBillingUpgradeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createBillingUpgradeRequest** | **CreateBillingUpgradeRequest** |  | 

### Return type

**CreateBillingUpgradeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBillingAppAllowances

> ListBillingAppAllowancesResponse ListBillingAppAllowances(ctx).Execute()

Your listing allowances through OAuth apps (used of limit, per app, with its platform scope).



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
    resp, r, err := apiClient.BillingApi.ListBillingAppAllowances(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ListBillingAppAllowances``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBillingAppAllowances`: ListBillingAppAllowancesResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ListBillingAppAllowances`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBillingAppAllowancesRequest struct via the builder pattern


### Return type

**ListBillingAppAllowancesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBillingPlans

> ListBillingPlansResponse ListBillingPlans(ctx).Execute()

List the plans with their prices, limits and features.



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
    resp, r, err := apiClient.BillingApi.ListBillingPlans(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BillingApi.ListBillingPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBillingPlans`: ListBillingPlansResponse
    fmt.Fprintf(os.Stdout, "Response from `BillingApi.ListBillingPlans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBillingPlansRequest struct via the builder pattern


### Return type

**ListBillingPlansResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

