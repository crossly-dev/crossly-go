# \BillingApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBillingUpgrade**](BillingApi.md#CreateBillingUpgrade) | **Post** /v1/billing/upgrade | Start an upgrade to a higher plan.



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

