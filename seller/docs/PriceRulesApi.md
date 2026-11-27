# \PriceRulesApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePriceRulePreview**](PriceRulesApi.md#CreatePriceRulePreview) | **Post** /v1/price-rules/preview | Compute the price each marketplace would list a listing (or a price) at.
[**DeletePriceRule**](PriceRulesApi.md#DeletePriceRule) | **Delete** /v1/price-rules/{id} | Delete a price rule.
[**ListPriceRules**](PriceRulesApi.md#ListPriceRules) | **Get** /v1/price-rules | List the seller&#39;s per-channel price rules.
[**UpdatePriceRule**](PriceRulesApi.md#UpdatePriceRule) | **Put** /v1/price-rules | Create or replace the price rule for a platform (and optional account / category).
[**UpdatePriceRuleById**](PriceRulesApi.md#UpdatePriceRuleById) | **Patch** /v1/price-rules/{id} | Update a price rule.



## CreatePriceRulePreview

> CreatePriceRulePreviewResponse CreatePriceRulePreview(ctx).Execute()

Compute the price each marketplace would list a listing (or a price) at.

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
    resp, r, err := apiClient.PriceRulesApi.CreatePriceRulePreview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceRulesApi.CreatePriceRulePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePriceRulePreview`: CreatePriceRulePreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `PriceRulesApi.CreatePriceRulePreview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePriceRulePreviewRequest struct via the builder pattern


### Return type

**CreatePriceRulePreviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeletePriceRule

> OkResponse DeletePriceRule(ctx, id).Execute()

Delete a price rule.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceRulesApi.DeletePriceRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceRulesApi.DeletePriceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePriceRule`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `PriceRulesApi.DeletePriceRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePriceRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListPriceRules

> ListPriceRulesResponse ListPriceRules(ctx).Platform(platform).Execute()

List the seller's per-channel price rules.

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
    platform := "platform_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceRulesApi.ListPriceRules(context.Background()).Platform(platform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceRulesApi.ListPriceRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPriceRules`: ListPriceRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `PriceRulesApi.ListPriceRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPriceRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platform** | **string** |  | 

### Return type

**ListPriceRulesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdatePriceRule

> ListPriceRulesItem UpdatePriceRule(ctx).Execute()

Create or replace the price rule for a platform (and optional account / category).

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
    resp, r, err := apiClient.PriceRulesApi.UpdatePriceRule(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceRulesApi.UpdatePriceRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePriceRule`: ListPriceRulesItem
    fmt.Fprintf(os.Stdout, "Response from `PriceRulesApi.UpdatePriceRule`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePriceRuleRequest struct via the builder pattern


### Return type

**ListPriceRulesItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdatePriceRuleById

> ListPriceRulesItem UpdatePriceRuleById(ctx, id).Execute()

Update a price rule.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PriceRulesApi.UpdatePriceRuleById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PriceRulesApi.UpdatePriceRuleById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePriceRuleById`: ListPriceRulesItem
    fmt.Fprintf(os.Stdout, "Response from `PriceRulesApi.UpdatePriceRuleById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePriceRuleByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListPriceRulesItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

