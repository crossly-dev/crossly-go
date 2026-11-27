# \TaxApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMileage**](TaxApi.md#CreateMileage) | **Post** /v1/mileage | Create a mileage entry.
[**DeleteMileage**](TaxApi.md#DeleteMileage) | **Delete** /v1/mileage/{id} | Delete a mileage entry.
[**GetBookTaxProfile**](TaxApi.md#GetBookTaxProfile) | **Get** /v1/books/tax-profile | Your tax profile: time zone for books boundaries, filing status, state, set-aside and rate assumptions, accounting method, and business name, EIN and address.
[**GetMileageSummary**](TaxApi.md#GetMileageSummary) | **Get** /v1/mileage/summary | Annual mileage totals + IRS deduction. Each trip is priced at the IRS rate on its own date (2026 has two rates, listed in &#x60;rates&#x60;); &#x60;ratePerMile&#x60; is the January rate.
[**GetOrderSalesTax**](TaxApi.md#GetOrderSalesTax) | **Get** /v1/orders/{id}/sales-tax | Who collected sales tax on one Crossly marketplace order, where, and the seller notice.
[**GetTaxSalesTax**](TaxApi.md#GetTaxSalesTax) | **Get** /v1/tax/sales-tax | Crossly marketplace sales Crossly did not collect sales tax on, by jurisdiction.
[**GetTaxScheduleC**](TaxApi.md#GetTaxScheduleC) | **Get** /v1/tax/schedule-c | Schedule C JSON for a tax year (default the current year), from the same P&amp;L as GET /v1/books: line 1 gross receipts before refunds, line 2 refunds, COGS in Part III, completed sales only. GET /v1/books has the full P&amp;L by period.
[**ListMileage**](TaxApi.md#ListMileage) | **Get** /v1/mileage | List mileage entries, newest trip first, paged (limit up to 500). Narrow by year and/or an inclusive from/to date range.
[**ListTaxCrosslyCollects**](TaxApi.md#ListTaxCrosslyCollects) | **Get** /v1/tax/crossly-collects | Where Crossly collects sales tax today (its active Stripe Tax registrations).
[**ListTaxSalesTaxNotCollected**](TaxApi.md#ListTaxSalesTaxNotCollected) | **Get** /v1/tax/sales-tax/not-collected | The individual Crossly sales Crossly did not collect sales tax on, newest first.
[**UpdateBookTaxProfile**](TaxApi.md#UpdateBookTaxProfile) | **Put** /v1/books/tax-profile | Update your tax profile: time zone, filing status, state, set-aside and rate assumptions (basis points), accounting method, business name, EIN and address. Fields left out are unchanged.
[**UpdateMileage**](TaxApi.md#UpdateMileage) | **Patch** /v1/mileage/{id} | Update a mileage entry.



## CreateMileage

> CreateMileageResponse CreateMileage(ctx).Execute()

Create a mileage entry.

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
    resp, r, err := apiClient.TaxApi.CreateMileage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxApi.CreateMileage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMileage`: CreateMileageResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxApi.CreateMileage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMileageRequest struct via the builder pattern


### Return type

**CreateMileageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteMileage

> OkResponse DeleteMileage(ctx, id).Execute()

Delete a mileage entry.

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
    resp, r, err := apiClient.TaxApi.DeleteMileage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxApi.DeleteMileage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMileage`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxApi.DeleteMileage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMileageRequest struct via the builder pattern


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


## GetBookTaxProfile

> GetBookTaxProfileResponse GetBookTaxProfile(ctx).Execute()

Your tax profile: time zone for books boundaries, filing status, state, set-aside and rate assumptions, accounting method, and business name, EIN and address.

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
    resp, r, err := apiClient.TaxApi.GetBookTaxProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxApi.GetBookTaxProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookTaxProfile`: GetBookTaxProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxApi.GetBookTaxProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBookTaxProfileRequest struct via the builder pattern


### Return type

**GetBookTaxProfileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetMileageSummary

> GetMileageSummaryResponse GetMileageSummary(ctx).Year(year).Execute()

Annual mileage totals + IRS deduction. Each trip is priced at the IRS rate on its own date (2026 has two rates, listed in `rates`); `ratePerMile` is the January rate.

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
    year := int32(56) // int32 | Calendar tax year. Default the current year. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxApi.GetMileageSummary(context.Background()).Year(year).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxApi.GetMileageSummary``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMileageSummary`: GetMileageSummaryResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxApi.GetMileageSummary`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMileageSummaryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **int32** | Calendar tax year. Default the current year. | 

### Return type

**GetMileageSummaryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetOrderSalesTax

> GetOrderSalesTaxResponse GetOrderSalesTax(ctx, id).Execute()

Who collected sales tax on one Crossly marketplace order, where, and the seller notice.

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
    resp, r, err := apiClient.TaxApi.GetOrderSalesTax(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxApi.GetOrderSalesTax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetOrderSalesTax`: GetOrderSalesTaxResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxApi.GetOrderSalesTax`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOrderSalesTaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetOrderSalesTaxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTaxSalesTax

> GetTaxSalesTaxResponse GetTaxSalesTax(ctx).From(from).To(to).Execute()

Crossly marketplace sales Crossly did not collect sales tax on, by jurisdiction.

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
    from := "from_example" // string | YYYY-MM-DD, inclusive. Default 1 January this year. (optional)
    to := "to_example" // string | YYYY-MM-DD, inclusive. Default now. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxApi.GetTaxSalesTax(context.Background()).From(from).To(to).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxApi.GetTaxSalesTax``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaxSalesTax`: GetTaxSalesTaxResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxApi.GetTaxSalesTax`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxSalesTaxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** | YYYY-MM-DD, inclusive. Default 1 January this year. | 
 **to** | **string** | YYYY-MM-DD, inclusive. Default now. | 

### Return type

**GetTaxSalesTaxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTaxScheduleC

> GetTaxScheduleCResponse GetTaxScheduleC(ctx).Execute()

Schedule C JSON for a tax year (default the current year), from the same P&L as GET /v1/books: line 1 gross receipts before refunds, line 2 refunds, COGS in Part III, completed sales only. GET /v1/books has the full P&L by period.

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
    resp, r, err := apiClient.TaxApi.GetTaxScheduleC(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxApi.GetTaxScheduleC``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTaxScheduleC`: GetTaxScheduleCResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxApi.GetTaxScheduleC`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTaxScheduleCRequest struct via the builder pattern


### Return type

**GetTaxScheduleCResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListMileage

> ListMileageResponse ListMileage(ctx).Year(year).From(from).To(to).Page(page).Limit(limit).Execute()

List mileage entries, newest trip first, paged (limit up to 500). Narrow by year and/or an inclusive from/to date range.

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
    year := int32(56) // int32 | Only trips dated in this calendar year. (optional)
    from := "from_example" // string | YYYY-MM-DD, inclusive. Only trips on or after this date. Combines with year. (optional)
    to := "to_example" // string | YYYY-MM-DD, inclusive. Only trips on or before this date; from <= to. (optional)
    page := int32(56) // int32 | 1-indexed page. Default 1. (optional)
    limit := int32(56) // int32 | Rows per page, default 100, clamped to 500. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxApi.ListMileage(context.Background()).Year(year).From(from).To(to).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxApi.ListMileage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMileage`: ListMileageResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxApi.ListMileage`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMileageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **int32** | Only trips dated in this calendar year. | 
 **from** | **string** | YYYY-MM-DD, inclusive. Only trips on or after this date. Combines with year. | 
 **to** | **string** | YYYY-MM-DD, inclusive. Only trips on or before this date; from &lt;&#x3D; to. | 
 **page** | **int32** | 1-indexed page. Default 1. | 
 **limit** | **int32** | Rows per page, default 100, clamped to 500. | 

### Return type

**ListMileageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListTaxCrosslyCollects

> ListTaxCrosslyCollectsResponse ListTaxCrosslyCollects(ctx).Execute()

Where Crossly collects sales tax today (its active Stripe Tax registrations).

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
    resp, r, err := apiClient.TaxApi.ListTaxCrosslyCollects(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxApi.ListTaxCrosslyCollects``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaxCrosslyCollects`: ListTaxCrosslyCollectsResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxApi.ListTaxCrosslyCollects`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTaxCrosslyCollectsRequest struct via the builder pattern


### Return type

**ListTaxCrosslyCollectsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListTaxSalesTaxNotCollected

> ListTaxSalesTaxNotCollectedResponse ListTaxSalesTaxNotCollected(ctx).From(from).To(to).Page(page).Limit(limit).Execute()

The individual Crossly sales Crossly did not collect sales tax on, newest first.

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
    from := "from_example" // string |  (optional)
    to := "to_example" // string |  (optional)
    page := int32(56) // int32 |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TaxApi.ListTaxSalesTaxNotCollected(context.Background()).From(from).To(to).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxApi.ListTaxSalesTaxNotCollected``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTaxSalesTaxNotCollected`: ListTaxSalesTaxNotCollectedResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxApi.ListTaxSalesTaxNotCollected`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListTaxSalesTaxNotCollectedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **string** |  | 
 **to** | **string** |  | 
 **page** | **int32** |  | 
 **limit** | **int32** |  | 

### Return type

**ListTaxSalesTaxNotCollectedResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBookTaxProfile

> GetBookTaxProfileResponse UpdateBookTaxProfile(ctx).Execute()

Update your tax profile: time zone, filing status, state, set-aside and rate assumptions (basis points), accounting method, business name, EIN and address. Fields left out are unchanged.

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
    resp, r, err := apiClient.TaxApi.UpdateBookTaxProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxApi.UpdateBookTaxProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBookTaxProfile`: GetBookTaxProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxApi.UpdateBookTaxProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBookTaxProfileRequest struct via the builder pattern


### Return type

**GetBookTaxProfileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateMileage

> CreateMileageResponse UpdateMileage(ctx, id).Execute()

Update a mileage entry.

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
    resp, r, err := apiClient.TaxApi.UpdateMileage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TaxApi.UpdateMileage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMileage`: CreateMileageResponse
    fmt.Fprintf(os.Stdout, "Response from `TaxApi.UpdateMileage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMileageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateMileageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

