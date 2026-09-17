# \ReferenceApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetBrand**](ReferenceApi.md#GetBrand) | **Get** /v1/brands | Search the Crossly brand index. Returns up to 50 matches.
[**GetCategory**](ReferenceApi.md#GetCategory) | **Get** /v1/categories | List Crossly&#39;s canonical category tree.
[**GetDepartment**](ReferenceApi.md#GetDepartment) | **Get** /v1/departments | Search the eBay-sourced \&quot;Department\&quot; item-specific values.
[**GetGender**](ReferenceApi.md#GetGender) | **Get** /v1/genders | Search the eBay-sourced \&quot;Gender\&quot; item-specific values.
[**GetPattern**](ReferenceApi.md#GetPattern) | **Get** /v1/patterns | Search the eBay-sourced \&quot;Pattern\&quot; item-specific values.
[**GetSizeSystem**](ReferenceApi.md#GetSizeSystem) | **Get** /v1/size-systems | Search the Poshmark + Vestiaire size-system union (US/UK/EU/AU/FR/KR).
[**GetStyle**](ReferenceApi.md#GetStyle) | **Get** /v1/styles | Search the eBay-sourced \&quot;Style\&quot; item-specific values.
[**GetType**](ReferenceApi.md#GetType) | **Get** /v1/types | Search the eBay-sourced \&quot;Type\&quot; item-specific values.



## GetBrand

> GetBrandResponse GetBrand(ctx).Execute()

Search the Crossly brand index. Returns up to 50 matches.

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
    resp, r, err := apiClient.ReferenceApi.GetBrand(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceApi.GetBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrand`: GetBrandResponse
    fmt.Fprintf(os.Stdout, "Response from `ReferenceApi.GetBrand`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandRequest struct via the builder pattern


### Return type

[**GetBrandResponse**](GetBrandResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCategory

> GetCategoryResponse GetCategory(ctx).Execute()

List Crossly's canonical category tree.

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
    resp, r, err := apiClient.ReferenceApi.GetCategory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceApi.GetCategory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCategory`: GetCategoryResponse
    fmt.Fprintf(os.Stdout, "Response from `ReferenceApi.GetCategory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCategoryRequest struct via the builder pattern


### Return type

[**GetCategoryResponse**](GetCategoryResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDepartment

> GetDepartmentResponse GetDepartment(ctx).Execute()

Search the eBay-sourced \"Department\" item-specific values.

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
    resp, r, err := apiClient.ReferenceApi.GetDepartment(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceApi.GetDepartment``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDepartment`: GetDepartmentResponse
    fmt.Fprintf(os.Stdout, "Response from `ReferenceApi.GetDepartment`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDepartmentRequest struct via the builder pattern


### Return type

[**GetDepartmentResponse**](GetDepartmentResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGender

> GetGenderResponse GetGender(ctx).Execute()

Search the eBay-sourced \"Gender\" item-specific values.

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
    resp, r, err := apiClient.ReferenceApi.GetGender(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceApi.GetGender``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGender`: GetGenderResponse
    fmt.Fprintf(os.Stdout, "Response from `ReferenceApi.GetGender`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGenderRequest struct via the builder pattern


### Return type

[**GetGenderResponse**](GetGenderResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPattern

> GetPatternResponse GetPattern(ctx).Execute()

Search the eBay-sourced \"Pattern\" item-specific values.

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
    resp, r, err := apiClient.ReferenceApi.GetPattern(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceApi.GetPattern``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPattern`: GetPatternResponse
    fmt.Fprintf(os.Stdout, "Response from `ReferenceApi.GetPattern`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPatternRequest struct via the builder pattern


### Return type

[**GetPatternResponse**](GetPatternResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSizeSystem

> GetSizeSystemResponse GetSizeSystem(ctx).Execute()

Search the Poshmark + Vestiaire size-system union (US/UK/EU/AU/FR/KR).

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
    resp, r, err := apiClient.ReferenceApi.GetSizeSystem(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceApi.GetSizeSystem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSizeSystem`: GetSizeSystemResponse
    fmt.Fprintf(os.Stdout, "Response from `ReferenceApi.GetSizeSystem`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSizeSystemRequest struct via the builder pattern


### Return type

[**GetSizeSystemResponse**](GetSizeSystemResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStyle

> GetStyleResponse GetStyle(ctx).Execute()

Search the eBay-sourced \"Style\" item-specific values.

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
    resp, r, err := apiClient.ReferenceApi.GetStyle(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceApi.GetStyle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStyle`: GetStyleResponse
    fmt.Fprintf(os.Stdout, "Response from `ReferenceApi.GetStyle`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStyleRequest struct via the builder pattern


### Return type

[**GetStyleResponse**](GetStyleResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetType

> GetTypeResponse GetType(ctx).Execute()

Search the eBay-sourced \"Type\" item-specific values.

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
    resp, r, err := apiClient.ReferenceApi.GetType(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReferenceApi.GetType``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetType`: GetTypeResponse
    fmt.Fprintf(os.Stdout, "Response from `ReferenceApi.GetType`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTypeRequest struct via the builder pattern


### Return type

[**GetTypeResponse**](GetTypeResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

