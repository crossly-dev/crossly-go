# \DiscountsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAutomaticDiscount**](DiscountsApi.md#CreateAutomaticDiscount) | **Post** /v1/automatic-discounts | Create an automatic discount.
[**CreateAutomaticDiscountPreview**](DiscountsApi.md#CreateAutomaticDiscountPreview) | **Post** /v1/automatic-discounts/preview | Preview a hypothetical cart of your listings against your automatic discounts (or an unsaved draft): what each line gets.
[**CreateDiscountCode**](DiscountsApi.md#CreateDiscountCode) | **Post** /v1/discount-codes | Create a discount code: percent, fixed, free shipping, buy X get Y, spend tiers or a set price, with requirements, item scope, limits (total, per buyer, per day, budget, regions), schedule and stacking.
[**CreateDiscountCodeGenerate**](DiscountsApi.md#CreateDiscountCodeGenerate) | **Post** /v1/discount-codes/{id}/generate | Generate up to 5,000 unique single-use codes with this code’s rules (their uses count toward its limits).
[**CreateDiscountCodeStatus**](DiscountsApi.md#CreateDiscountCodeStatus) | **Post** /v1/discount-codes/{id}/status | Pause, resume (active) or archive a discount code.
[**CreateDiscountHolder**](DiscountsApi.md#CreateDiscountHolder) | **Post** /v1/discounts/holders | Create a Collection holders deal in the store (X-Storefront-Id; default store when omitted).
[**DeleteAutomaticDiscount**](DiscountsApi.md#DeleteAutomaticDiscount) | **Delete** /v1/automatic-discounts/{id} | Delete an automatic discount.
[**DeleteDiscountHolder**](DiscountsApi.md#DeleteDiscountHolder) | **Delete** /v1/discounts/holders/{id} | Delete a Collection holders deal. Returns &#x60;{ ok }&#x60;.
[**GetDiscountCode**](DiscountsApi.md#GetDiscountCode) | **Get** /v1/discount-codes/{id} | One discount code with its rules and usage.
[**GetDiscountCodeShare**](DiscountsApi.md#GetDiscountCodeShare) | **Get** /v1/discount-codes/{id}/share | The share link that pre-applies the code at checkout (the store page with ?code&#x3D;) and a QR code of it as a PNG data URL.
[**GetDiscountHolder**](DiscountsApi.md#GetDiscountHolder) | **Get** /v1/discounts/holders/{id} | One Collection holders deal with its usage (uses, cents given).
[**ListAutomaticDiscounts**](DiscountsApi.md#ListAutomaticDiscounts) | **Get** /v1/automatic-discounts | Your automatic discounts (buy more, save more): quantity and mixed-item tiers, scope, channels and schedule.
[**ListDiscountCodeRedemptions**](DiscountsApi.md#ListDiscountCodeRedemptions) | **Get** /v1/discount-codes/{id}/redemptions | A code’s usage history: orders, masked buyers, region, channel and the amount given, newest first.
[**ListDiscountCodes**](DiscountsApi.md#ListDiscountCodes) | **Get** /v1/discount-codes | Your discount codes: kind, status (active, scheduled, paused, expired, exhausted, archived), every rule, and usage.
[**ListDiscountHolders**](DiscountsApi.md#ListDiscountHolders) | **Get** /v1/discounts/holders | The store’s Collection holders deals: a percent off for buyers holding any token of a collection (by chain + contract), with scope, channels, schedule, stacking and limits.
[**UpdateAutomaticDiscount**](DiscountsApi.md#UpdateAutomaticDiscount) | **Put** /v1/automatic-discounts/{id} | Replace an automatic discount.
[**UpdateDiscountCode**](DiscountsApi.md#UpdateDiscountCode) | **Patch** /v1/discount-codes/{id} | Update a discount code. The code itself can be renamed only until it is first used.
[**UpdateDiscountHolder**](DiscountsApi.md#UpdateDiscountHolder) | **Put** /v1/discounts/holders/{id} | Replace a Collection holders deal.



## CreateAutomaticDiscount

> CreateAutomaticDiscountResponse CreateAutomaticDiscount(ctx).Execute()

Create an automatic discount.

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
    resp, r, err := apiClient.DiscountsApi.CreateAutomaticDiscount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.CreateAutomaticDiscount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAutomaticDiscount`: CreateAutomaticDiscountResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.CreateAutomaticDiscount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutomaticDiscountRequest struct via the builder pattern


### Return type

**CreateAutomaticDiscountResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateAutomaticDiscountPreview

> CreateAutomaticDiscountPreviewResponse CreateAutomaticDiscountPreview(ctx).Execute()

Preview a hypothetical cart of your listings against your automatic discounts (or an unsaved draft): what each line gets.

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
    resp, r, err := apiClient.DiscountsApi.CreateAutomaticDiscountPreview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.CreateAutomaticDiscountPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAutomaticDiscountPreview`: CreateAutomaticDiscountPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.CreateAutomaticDiscountPreview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAutomaticDiscountPreviewRequest struct via the builder pattern


### Return type

**CreateAutomaticDiscountPreviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDiscountCode

> DiscountCodeResponse CreateDiscountCode(ctx).Execute()

Create a discount code: percent, fixed, free shipping, buy X get Y, spend tiers or a set price, with requirements, item scope, limits (total, per buyer, per day, budget, regions), schedule and stacking.

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
    resp, r, err := apiClient.DiscountsApi.CreateDiscountCode(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.CreateDiscountCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDiscountCode`: DiscountCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.CreateDiscountCode`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiscountCodeRequest struct via the builder pattern


### Return type

**DiscountCodeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDiscountCodeGenerate

> CreateDiscountCodeGenerateResponse CreateDiscountCodeGenerate(ctx, id).Execute()

Generate up to 5,000 unique single-use codes with this code’s rules (their uses count toward its limits).

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
    resp, r, err := apiClient.DiscountsApi.CreateDiscountCodeGenerate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.CreateDiscountCodeGenerate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDiscountCodeGenerate`: CreateDiscountCodeGenerateResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.CreateDiscountCodeGenerate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiscountCodeGenerateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateDiscountCodeGenerateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDiscountCodeStatus

> DiscountCodeResponse CreateDiscountCodeStatus(ctx, id).Execute()

Pause, resume (active) or archive a discount code.

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
    resp, r, err := apiClient.DiscountsApi.CreateDiscountCodeStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.CreateDiscountCodeStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDiscountCodeStatus`: DiscountCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.CreateDiscountCodeStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiscountCodeStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DiscountCodeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDiscountHolder

> CreateDiscountHolderResponse CreateDiscountHolder(ctx).Execute()

Create a Collection holders deal in the store (X-Storefront-Id; default store when omitted).

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
    resp, r, err := apiClient.DiscountsApi.CreateDiscountHolder(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.CreateDiscountHolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDiscountHolder`: CreateDiscountHolderResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.CreateDiscountHolder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDiscountHolderRequest struct via the builder pattern


### Return type

**CreateDiscountHolderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteAutomaticDiscount

> UpdateGroupBuyResponse DeleteAutomaticDiscount(ctx, id).Execute()

Delete an automatic discount.

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
    resp, r, err := apiClient.DiscountsApi.DeleteAutomaticDiscount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.DeleteAutomaticDiscount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAutomaticDiscount`: UpdateGroupBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.DeleteAutomaticDiscount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAutomaticDiscountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateGroupBuyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteDiscountHolder

> OkResponse DeleteDiscountHolder(ctx, id).Execute()

Delete a Collection holders deal. Returns `{ ok }`.

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
    resp, r, err := apiClient.DiscountsApi.DeleteDiscountHolder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.DeleteDiscountHolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDiscountHolder`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.DeleteDiscountHolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDiscountHolderRequest struct via the builder pattern


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


## GetDiscountCode

> GetDiscountCodeResponse GetDiscountCode(ctx, id).Execute()

One discount code with its rules and usage.

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
    resp, r, err := apiClient.DiscountsApi.GetDiscountCode(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.GetDiscountCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscountCode`: GetDiscountCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.GetDiscountCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscountCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetDiscountCodeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetDiscountCodeShare

> GetDiscountCodeShareResponse GetDiscountCodeShare(ctx, id).Execute()

The share link that pre-applies the code at checkout (the store page with ?code=) and a QR code of it as a PNG data URL.

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
    resp, r, err := apiClient.DiscountsApi.GetDiscountCodeShare(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.GetDiscountCodeShare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscountCodeShare`: GetDiscountCodeShareResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.GetDiscountCodeShare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscountCodeShareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetDiscountCodeShareResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetDiscountHolder

> GetDiscountHolderResponse GetDiscountHolder(ctx, id).Execute()

One Collection holders deal with its usage (uses, cents given).

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
    resp, r, err := apiClient.DiscountsApi.GetDiscountHolder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.GetDiscountHolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDiscountHolder`: GetDiscountHolderResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.GetDiscountHolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDiscountHolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetDiscountHolderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListAutomaticDiscounts

> ListAutomaticDiscountsResponse ListAutomaticDiscounts(ctx).Execute()

Your automatic discounts (buy more, save more): quantity and mixed-item tiers, scope, channels and schedule.

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
    resp, r, err := apiClient.DiscountsApi.ListAutomaticDiscounts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.ListAutomaticDiscounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAutomaticDiscounts`: ListAutomaticDiscountsResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.ListAutomaticDiscounts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAutomaticDiscountsRequest struct via the builder pattern


### Return type

**ListAutomaticDiscountsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListDiscountCodeRedemptions

> ListDiscountCodeRedemptionsResponse ListDiscountCodeRedemptions(ctx, id).Execute()

A code’s usage history: orders, masked buyers, region, channel and the amount given, newest first.

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
    resp, r, err := apiClient.DiscountsApi.ListDiscountCodeRedemptions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.ListDiscountCodeRedemptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDiscountCodeRedemptions`: ListDiscountCodeRedemptionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.ListDiscountCodeRedemptions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDiscountCodeRedemptionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListDiscountCodeRedemptionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListDiscountCodes

> ListDiscountCodesResponse ListDiscountCodes(ctx).Execute()

Your discount codes: kind, status (active, scheduled, paused, expired, exhausted, archived), every rule, and usage.

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
    resp, r, err := apiClient.DiscountsApi.ListDiscountCodes(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.ListDiscountCodes``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDiscountCodes`: ListDiscountCodesResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.ListDiscountCodes`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDiscountCodesRequest struct via the builder pattern


### Return type

**ListDiscountCodesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListDiscountHolders

> ListDiscountHoldersResponse ListDiscountHolders(ctx).Execute()

The store’s Collection holders deals: a percent off for buyers holding any token of a collection (by chain + contract), with scope, channels, schedule, stacking and limits.

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
    resp, r, err := apiClient.DiscountsApi.ListDiscountHolders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.ListDiscountHolders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDiscountHolders`: ListDiscountHoldersResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.ListDiscountHolders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDiscountHoldersRequest struct via the builder pattern


### Return type

**ListDiscountHoldersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateAutomaticDiscount

> CreateAutomaticDiscountResponse UpdateAutomaticDiscount(ctx, id).Execute()

Replace an automatic discount.

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
    resp, r, err := apiClient.DiscountsApi.UpdateAutomaticDiscount(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.UpdateAutomaticDiscount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAutomaticDiscount`: CreateAutomaticDiscountResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.UpdateAutomaticDiscount`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAutomaticDiscountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateAutomaticDiscountResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateDiscountCode

> DiscountCodeResponse UpdateDiscountCode(ctx, id).Execute()

Update a discount code. The code itself can be renamed only until it is first used.

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
    resp, r, err := apiClient.DiscountsApi.UpdateDiscountCode(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.UpdateDiscountCode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDiscountCode`: DiscountCodeResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.UpdateDiscountCode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDiscountCodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DiscountCodeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateDiscountHolder

> CreateDiscountHolderResponse UpdateDiscountHolder(ctx, id).Execute()

Replace a Collection holders deal.

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
    resp, r, err := apiClient.DiscountsApi.UpdateDiscountHolder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DiscountsApi.UpdateDiscountHolder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDiscountHolder`: CreateDiscountHolderResponse
    fmt.Fprintf(os.Stdout, "Response from `DiscountsApi.UpdateDiscountHolder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDiscountHolderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateDiscountHolderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

