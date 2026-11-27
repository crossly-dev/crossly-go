# \TaxonomyApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCategoryById**](TaxonomyApi.md#GetCategoryById) | **Get** /v1/categories/{id} | One canonical category and the category it maps to on each marketplace.
[**GetTaxonomyCategory**](TaxonomyApi.md#GetTaxonomyCategory) | **Get** /v1/taxonomy/{platform}/categories | Categories for a platform. Default is top-level; pass &#x60;?parent&#x3D;&lt;categoryId&gt;&#x60; to drill down one level (supported on cookie platforms whose recipe returns flat parent_id-linked rows).
[**GetTaxonomyCategoryAspect**](TaxonomyApi.md#GetTaxonomyCategoryAspect) | **Get** /v1/taxonomy/{platform}/categories/{id}/aspects | Item-specific aspects (eBay) / properties (Etsy) / hard-coded enums (cookie platforms) for a category.
[**GetTaxonomyCategoryChildren**](TaxonomyApi.md#GetTaxonomyCategoryChildren) | **Get** /v1/taxonomy/{platform}/categories/{id}/children | Direct children of a category node.
[**GetTaxonomyRequiredField**](TaxonomyApi.md#GetTaxonomyRequiredField) | **Get** /v1/taxonomy/{platform}/required-fields | Normalized field schema the seller needs to fill before crossposting to this platform. Combines master fields (title/description/price/condition) with platform-specific overrides.
[**ListCategoryAspects**](TaxonomyApi.md#ListCategoryAspects) | **Get** /v1/categories/{id}/aspects | Item specifics for a canonical category: required/optional, known values, per-marketplace names.
[**ListCategorySearch**](TaxonomyApi.md#ListCategorySearch) | **Get** /v1/categories/search | Search or browse Crossly&#39;s canonical categories (the ids listings take as categoryId).
[**ListTaxonomyEbayCharities**](TaxonomyApi.md#ListTaxonomyEbayCharities) | **Get** /v1/taxonomy/ebay/charities | Search eBay for Charity organisations (the charityId an eBay listing donates to).
[**ListTaxonomySuggest**](TaxonomyApi.md#ListTaxonomySuggest) | **Get** /v1/taxonomy/{platform}/suggest | Reverse lookup — suggest categories matching a search phrase: eBay live; Poshmark, Mercari, Depop, Grailed, Vinted, Whatnot, Vestiaire, OfferUp, Facebook and Curtsy over their harvested trees. Etsy answers an empty list (browse its tree instead).



## GetCategoryById

> GetCategoryByIdResponse GetCategoryById(ctx, id).Execute()

One canonical category and the category it maps to on each marketplace.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.GetCategoryById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.GetCategoryById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategoryById`: GetCategoryByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.GetCategoryById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetCategoryByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


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

**GetTaxonomyCategoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**GetTaxonomyCategoryAspectResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**GetTaxonomyCategoryChildrenResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**GetTaxonomyRequiredFieldResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCategoryAspects

> ListCategoryAspectsResponse ListCategoryAspects(ctx, id).ValuesLimit(valuesLimit).Execute()

Item specifics for a canonical category: required/optional, known values, per-marketplace names.



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
    id := "id_example" // string | 
    valuesLimit := int32(56) // int32 |  (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.ListCategoryAspects(context.Background(), id).ValuesLimit(valuesLimit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.ListCategoryAspects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCategoryAspects`: ListCategoryAspectsResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.ListCategoryAspects`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCategoryAspectsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **valuesLimit** | **int32** |  | [default to 50]

### Return type

**ListCategoryAspectsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCategorySearch

> ListCategorySearchResponse ListCategorySearch(ctx).Q(q).Path(path).Limit(limit).Execute()

Search or browse Crossly's canonical categories (the ids listings take as categoryId).



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
    q := "q_example" // string |  (optional)
    path := "path_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.ListCategorySearch(context.Background()).Q(q).Path(path).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.ListCategorySearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCategorySearch`: ListCategorySearchResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.ListCategorySearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCategorySearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **path** | **string** |  | 
 **limit** | **int32** |  | [default to 25]

### Return type

**ListCategorySearchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListTaxonomyEbayCharities

> ListTaxonomyEbayCharitiesResponse ListTaxonomyEbayCharities(ctx).Q(q).Limit(limit).Marketplace(marketplace).Execute()

Search eBay for Charity organisations (the charityId an eBay listing donates to).



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
    q := "q_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 20)
    marketplace := "marketplace_example" // string |  (optional) (default to "EBAY_US")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxonomyApi.ListTaxonomyEbayCharities(context.Background()).Q(q).Limit(limit).Marketplace(marketplace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxonomyApi.ListTaxonomyEbayCharities``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaxonomyEbayCharities`: ListTaxonomyEbayCharitiesResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxonomyApi.ListTaxonomyEbayCharities`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTaxonomyEbayCharitiesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** |  | 
 **limit** | **int32** |  | [default to 20]
 **marketplace** | **string** |  | [default to &quot;EBAY_US&quot;]

### Return type

**ListTaxonomyEbayCharitiesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListTaxonomySuggest

> ListTaxonomySuggestResponse ListTaxonomySuggest(ctx, platform).Q(q).Execute()

Reverse lookup — suggest categories matching a search phrase: eBay live; Poshmark, Mercari, Depop, Grailed, Vinted, Whatnot, Vestiaire, OfferUp, Facebook and Curtsy over their harvested trees. Etsy answers an empty list (browse its tree instead).

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
    // response from `ListTaxonomySuggest`: ListTaxonomySuggestResponse
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

**ListTaxonomySuggestResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

