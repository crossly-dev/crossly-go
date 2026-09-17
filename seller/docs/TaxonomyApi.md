# \TaxonomyApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTaxonomyCategory**](TaxonomyApi.md#GetTaxonomyCategory) | **Get** /v1/taxonomy/{platform}/categories | Categories for a platform. Default is top-level; pass &#x60;?parent&#x3D;&lt;categoryId&gt;&#x60; to drill down one level (supported on cookie platforms whose recipe returns flat parent_id-linked rows).
[**GetTaxonomyCategoryAspect**](TaxonomyApi.md#GetTaxonomyCategoryAspect) | **Get** /v1/taxonomy/{platform}/categories/{id}/aspects | Item-specific aspects (eBay) / properties (Etsy) / hard-coded enums (cookie platforms) for a category.
[**GetTaxonomyCategoryChildren**](TaxonomyApi.md#GetTaxonomyCategoryChildren) | **Get** /v1/taxonomy/{platform}/categories/{id}/children | Direct children of a category node.
[**GetTaxonomyRequiredField**](TaxonomyApi.md#GetTaxonomyRequiredField) | **Get** /v1/taxonomy/{platform}/required-fields | Normalized field schema the seller needs to fill before crossposting to this platform. Combines master fields (title/description/price/condition) with platform-specific overrides.
[**ListTaxonomySuggest**](TaxonomyApi.md#ListTaxonomySuggest) | **Get** /v1/taxonomy/{platform}/suggest | Reverse lookup — suggest categories matching a search phrase. eBay-only today.



## GetTaxonomyCategory

> GetTaxonomyCategoryResponse GetTaxonomyCategory(ctx, platform).Parent(parent).Execute()

Categories for a platform. Default is top-level; pass `?parent=<categoryId>` to drill down one level (supported on cookie platforms whose recipe returns flat parent_id-linked rows).

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
    platform := "platform_example" // string | 
    parent := "parent_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.GetTaxonomyCategory(context.Background(), platform).Parent(parent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.GetTaxonomyCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaxonomyCategory`: GetTaxonomyCategoryResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.GetTaxonomyCategory`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxonomyCategoryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **parent** | **string** |  | 

### Return type

[**GetTaxonomyCategoryResponse**](GetTaxonomyCategoryResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxonomyCategoryAspect

> GetTaxonomyCategoryAspectResponse GetTaxonomyCategoryAspect(ctx, platform, id).Execute()

Item-specific aspects (eBay) / properties (Etsy) / hard-coded enums (cookie platforms) for a category.

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
    platform := "platform_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.GetTaxonomyCategoryAspect(context.Background(), platform, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.GetTaxonomyCategoryAspect``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaxonomyCategoryAspect`: GetTaxonomyCategoryAspectResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.GetTaxonomyCategoryAspect`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxonomyCategoryAspectRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetTaxonomyCategoryAspectResponse**](GetTaxonomyCategoryAspectResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxonomyCategoryChildren

> GetTaxonomyCategoryChildrenResponse GetTaxonomyCategoryChildren(ctx, platform, id).Execute()

Direct children of a category node.

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
    platform := "platform_example" // string | 
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.GetTaxonomyCategoryChildren(context.Background(), platform, id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.GetTaxonomyCategoryChildren``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaxonomyCategoryChildren`: GetTaxonomyCategoryChildrenResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.GetTaxonomyCategoryChildren`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxonomyCategoryChildrenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**GetTaxonomyCategoryChildrenResponse**](GetTaxonomyCategoryChildrenResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTaxonomyRequiredField

> GetTaxonomyRequiredFieldResponse GetTaxonomyRequiredField(ctx, platform).CategoryId(categoryId).Execute()

Normalized field schema the seller needs to fill before crossposting to this platform. Combines master fields (title/description/price/condition) with platform-specific overrides.

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
    platform := "platform_example" // string | 
    categoryId := "categoryId_example" // string | Optional — used to inline aspects when present. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.GetTaxonomyRequiredField(context.Background(), platform).CategoryId(categoryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.GetTaxonomyRequiredField``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaxonomyRequiredField`: GetTaxonomyRequiredFieldResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.GetTaxonomyRequiredField`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxonomyRequiredFieldRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **categoryId** | **string** | Optional — used to inline aspects when present. | 

### Return type

[**GetTaxonomyRequiredFieldResponse**](GetTaxonomyRequiredFieldResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListTaxonomySuggest

> V1List ListTaxonomySuggest(ctx, platform).Q(q).Execute()

Reverse lookup — suggest categories matching a search phrase. eBay-only today.

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
    platform := "platform_example" // string | 
    q := "q_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.ListTaxonomySuggest(context.Background(), platform).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.ListTaxonomySuggest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaxonomySuggest`: V1List
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.ListTaxonomySuggest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTaxonomySuggestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **q** | **string** |  | 

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

