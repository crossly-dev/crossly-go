# \FieldRulesApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateFieldRule**](FieldRulesApi.md#CreateFieldRule) | **Post** /v1/field-rules | Create a per-channel field rule.
[**CreateFieldRulePreview**](FieldRulesApi.md#CreateFieldRulePreview) | **Post** /v1/field-rules/preview | Preview field-rule output for a listing on one channel.
[**CreateFieldRuleReorder**](FieldRulesApi.md#CreateFieldRuleReorder) | **Post** /v1/field-rules/reorder | Set the run order of several field rules.
[**DeleteFieldRule**](FieldRulesApi.md#DeleteFieldRule) | **Delete** /v1/field-rules/{id} | Delete a field rule.
[**GetFieldRule**](FieldRulesApi.md#GetFieldRule) | **Get** /v1/field-rules/{id} | Get one field rule.
[**ListFieldRules**](FieldRulesApi.md#ListFieldRules) | **Get** /v1/field-rules | List the seller&#39;s per-channel field rules.
[**UpdateFieldRule**](FieldRulesApi.md#UpdateFieldRule) | **Patch** /v1/field-rules/{id} | Update a field rule.



## CreateFieldRule

> ListFieldRulesItem CreateFieldRule(ctx).Execute()

Create a per-channel field rule.

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
    resp, r, err := apiClient.FieldRulesApi.CreateFieldRule(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldRulesApi.CreateFieldRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFieldRule`: ListFieldRulesItem
    fmt.Fprintf(os.Stdout, "Response from `FieldRulesApi.CreateFieldRule`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFieldRuleRequest struct via the builder pattern


### Return type

**ListFieldRulesItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFieldRulePreview

> CreateFieldRulePreviewResponse CreateFieldRulePreview(ctx).Execute()

Preview field-rule output for a listing on one channel.

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
    resp, r, err := apiClient.FieldRulesApi.CreateFieldRulePreview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldRulesApi.CreateFieldRulePreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFieldRulePreview`: CreateFieldRulePreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `FieldRulesApi.CreateFieldRulePreview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFieldRulePreviewRequest struct via the builder pattern


### Return type

**CreateFieldRulePreviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateFieldRuleReorder

> OkResponse CreateFieldRuleReorder(ctx).Execute()

Set the run order of several field rules.

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
    resp, r, err := apiClient.FieldRulesApi.CreateFieldRuleReorder(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldRulesApi.CreateFieldRuleReorder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateFieldRuleReorder`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `FieldRulesApi.CreateFieldRuleReorder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFieldRuleReorderRequest struct via the builder pattern


### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteFieldRule

> OkResponse DeleteFieldRule(ctx, id).Execute()

Delete a field rule.

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
    resp, r, err := apiClient.FieldRulesApi.DeleteFieldRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldRulesApi.DeleteFieldRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteFieldRule`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `FieldRulesApi.DeleteFieldRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFieldRuleRequest struct via the builder pattern


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


## GetFieldRule

> ListFieldRulesItem GetFieldRule(ctx, id).Execute()

Get one field rule.

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
    resp, r, err := apiClient.FieldRulesApi.GetFieldRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldRulesApi.GetFieldRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFieldRule`: ListFieldRulesItem
    fmt.Fprintf(os.Stdout, "Response from `FieldRulesApi.GetFieldRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFieldRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListFieldRulesItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListFieldRules

> ListFieldRulesResponse ListFieldRules(ctx).Platform(platform).Execute()

List the seller's per-channel field rules.

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
    resp, r, err := apiClient.FieldRulesApi.ListFieldRules(context.Background()).Platform(platform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldRulesApi.ListFieldRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFieldRules`: ListFieldRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `FieldRulesApi.ListFieldRules`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFieldRulesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platform** | **string** |  | 

### Return type

**ListFieldRulesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateFieldRule

> ListFieldRulesItem UpdateFieldRule(ctx, id).Execute()

Update a field rule.

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
    resp, r, err := apiClient.FieldRulesApi.UpdateFieldRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `FieldRulesApi.UpdateFieldRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateFieldRule`: ListFieldRulesItem
    fmt.Fprintf(os.Stdout, "Response from `FieldRulesApi.UpdateFieldRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFieldRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListFieldRulesItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

