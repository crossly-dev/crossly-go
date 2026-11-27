# \BooksApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBookExpense**](BooksApi.md#CreateBookExpense) | **Post** /v1/books/expenses | Log a business expense: date, amountCents, a Schedule C category, vendor, description, receipt link, optionally recurring monthly or yearly.
[**DeleteBookExpense**](BooksApi.md#DeleteBookExpense) | **Delete** /v1/books/expenses/{id} | Delete a logged expense. It leaves the P&amp;L and the exports.
[**GetBook**](BooksApi.md#GetBook) | **Get** /v1/books | Profit and loss for a year and period on the Schedule C lines: income, COGS, gross profit, expenses (marketplace fees, postage, mileage, your categories), net profit, monthly table, per-marketplace table with 1099-K expectations, Part III COGS, estimated tax and sales tax collected by marketplaces. Money in cents.
[**GetBook1099k**](BooksApi.md#GetBook1099k) | **Get** /v1/books/1099k | Per marketplace for a year: the 1099-K gross you recorded beside the gross Crossly computed, the difference, and whether a form is expected (federal: over $20,000 and over 200 transactions).
[**GetBookExpense**](BooksApi.md#GetBookExpense) | **Get** /v1/books/expenses | Your logged business expenses in a year and period (recurring ones expanded), with their Schedule C category. Amounts in cents.
[**GetBookTaxProfile**](BooksApi.md#GetBookTaxProfile) | **Get** /v1/books/tax-profile | Your tax profile: time zone for books boundaries, filing status, state, set-aside and rate assumptions, accounting method, and business name, EIN and address.
[**GetBookTransaction**](BooksApi.md#GetBookTransaction) | **Get** /v1/books/transactions | The sale rows behind any P&amp;L line, newest first: gross, refunded, fee and fee source, shipping, COGS, tax collected and net, in cents.
[**UpdateBook1099k**](BooksApi.md#UpdateBook1099k) | **Put** /v1/books/1099k | Record what a marketplace reported on its 1099-K for a year (gross in cents, transactions, notes). Answers the year&#39;s reconciliation, as GET /v1/books/1099k does.
[**UpdateBookExpense**](BooksApi.md#UpdateBookExpense) | **Patch** /v1/books/expenses/{id} | Change a logged expense.
[**UpdateBookTaxProfile**](BooksApi.md#UpdateBookTaxProfile) | **Put** /v1/books/tax-profile | Update your tax profile: time zone, filing status, state, set-aside and rate assumptions (basis points), accounting method, business name, EIN and address. Fields left out are unchanged.



## CreateBookExpense

> CreateBookExpenseResponse CreateBookExpense(ctx).Execute()

Log a business expense: date, amountCents, a Schedule C category, vendor, description, receipt link, optionally recurring monthly or yearly.

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
    resp, r, err := apiClient.BooksApi.CreateBookExpense(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BooksApi.CreateBookExpense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBookExpense`: CreateBookExpenseResponse
    fmt.Fprintf(os.Stdout, "Response from `BooksApi.CreateBookExpense`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBookExpenseRequest struct via the builder pattern


### Return type

**CreateBookExpenseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBookExpense

> DeleteStoryResponse DeleteBookExpense(ctx, id).Execute()

Delete a logged expense. It leaves the P&L and the exports.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The expense id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BooksApi.DeleteBookExpense(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BooksApi.DeleteBookExpense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBookExpense`: DeleteStoryResponse
    fmt.Fprintf(os.Stdout, "Response from `BooksApi.DeleteBookExpense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The expense id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBookExpenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteStoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBook

> GetBookResponse GetBook(ctx).Year(year).Period(period).Tz(tz).Execute()

Profit and loss for a year and period on the Schedule C lines: income, COGS, gross profit, expenses (marketplace fees, postage, mileage, your categories), net profit, monthly table, per-marketplace table with 1099-K expectations, Part III COGS, estimated tax and sales tax collected by marketplaces. Money in cents.

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
    year := int32(56) // int32 | Tax year. Default the current year. (optional)
    period := "period_example" // string | fy = full year (default), ytd, q1-q4, m01-m12. Boundaries are in your tax-profile time zone. (optional)
    tz := "tz_example" // string | IANA zone; default your tax profile zone. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BooksApi.GetBook(context.Background()).Year(year).Period(period).Tz(tz).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BooksApi.GetBook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBook`: GetBookResponse
    fmt.Fprintf(os.Stdout, "Response from `BooksApi.GetBook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **int32** | Tax year. Default the current year. | 
 **period** | **string** | fy &#x3D; full year (default), ytd, q1-q4, m01-m12. Boundaries are in your tax-profile time zone. | 
 **tz** | **string** | IANA zone; default your tax profile zone. | 

### Return type

**GetBookResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBook1099k

> GetBook1099kResponse GetBook1099k(ctx).Year(year).Execute()

Per marketplace for a year: the 1099-K gross you recorded beside the gross Crossly computed, the difference, and whether a form is expected (federal: over $20,000 and over 200 transactions).

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
    year := int32(56) // int32 | Tax year. Default the current year. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BooksApi.GetBook1099k(context.Background()).Year(year).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BooksApi.GetBook1099k``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBook1099k`: GetBook1099kResponse
    fmt.Fprintf(os.Stdout, "Response from `BooksApi.GetBook1099k`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBook1099kRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **int32** | Tax year. Default the current year. | 

### Return type

**GetBook1099kResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBookExpense

> GetBookExpenseResponse GetBookExpense(ctx).Year(year).Period(period).Category(category).Page(page).Size(size).Execute()

Your logged business expenses in a year and period (recurring ones expanded), with their Schedule C category. Amounts in cents.

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
    year := int32(56) // int32 | Tax year. Default the current year. (optional)
    period := "period_example" // string | fy = full year (default), ytd, q1-q4, m01-m12. Boundaries are in your tax-profile time zone. (optional)
    category := "category_example" // string | An expense category id, e.g. supplies, postage, software_subscriptions. (optional)
    page := int32(56) // int32 | 1-indexed page. Default 1. (optional)
    size := int32(56) // int32 | Rows per page, at most 100. Default 25. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BooksApi.GetBookExpense(context.Background()).Year(year).Period(period).Category(category).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BooksApi.GetBookExpense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookExpense`: GetBookExpenseResponse
    fmt.Fprintf(os.Stdout, "Response from `BooksApi.GetBookExpense`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBookExpenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **int32** | Tax year. Default the current year. | 
 **period** | **string** | fy &#x3D; full year (default), ytd, q1-q4, m01-m12. Boundaries are in your tax-profile time zone. | 
 **category** | **string** | An expense category id, e.g. supplies, postage, software_subscriptions. | 
 **page** | **int32** | 1-indexed page. Default 1. | 
 **size** | **int32** | Rows per page, at most 100. Default 25. | 

### Return type

**GetBookExpenseResponse**

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
    resp, r, err := apiClient.BooksApi.GetBookTaxProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BooksApi.GetBookTaxProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookTaxProfile`: GetBookTaxProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `BooksApi.GetBookTaxProfile`: %v\n", resp)
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


## GetBookTransaction

> GetBookTransactionResponse GetBookTransaction(ctx).Year(year).Period(period).Tz(tz).Platform(platform).Status(status).Page(page).Size(size).Execute()

The sale rows behind any P&L line, newest first: gross, refunded, fee and fee source, shipping, COGS, tax collected and net, in cents.

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
    year := int32(56) // int32 | Tax year. Default the current year. (optional)
    period := "period_example" // string | fy = full year (default), ytd, q1-q4, m01-m12. Boundaries are in your tax-profile time zone. (optional)
    tz := "tz_example" // string | IANA zone; default your tax profile zone. (optional)
    platform := "platform_example" // string | One platform id. (optional)
    status := "status_example" // string | completed = counted in income; pending = details not in yet (no money); default all. (optional)
    page := int32(56) // int32 | 1-indexed page. Default 1. (optional)
    size := int32(56) // int32 | Rows per page, at most 100. Default 50. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BooksApi.GetBookTransaction(context.Background()).Year(year).Period(period).Tz(tz).Platform(platform).Status(status).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BooksApi.GetBookTransaction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBookTransaction`: GetBookTransactionResponse
    fmt.Fprintf(os.Stdout, "Response from `BooksApi.GetBookTransaction`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBookTransactionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **year** | **int32** | Tax year. Default the current year. | 
 **period** | **string** | fy &#x3D; full year (default), ytd, q1-q4, m01-m12. Boundaries are in your tax-profile time zone. | 
 **tz** | **string** | IANA zone; default your tax profile zone. | 
 **platform** | **string** | One platform id. | 
 **status** | **string** | completed &#x3D; counted in income; pending &#x3D; details not in yet (no money); default all. | 
 **page** | **int32** | 1-indexed page. Default 1. | 
 **size** | **int32** | Rows per page, at most 100. Default 50. | 

### Return type

**GetBookTransactionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBook1099k

> GetBook1099kResponse UpdateBook1099k(ctx).Execute()

Record what a marketplace reported on its 1099-K for a year (gross in cents, transactions, notes). Answers the year's reconciliation, as GET /v1/books/1099k does.

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
    resp, r, err := apiClient.BooksApi.UpdateBook1099k(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BooksApi.UpdateBook1099k``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBook1099k`: GetBook1099kResponse
    fmt.Fprintf(os.Stdout, "Response from `BooksApi.UpdateBook1099k`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBook1099kRequest struct via the builder pattern


### Return type

**GetBook1099kResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBookExpense

> CreateBookExpenseResponse UpdateBookExpense(ctx, id).Execute()

Change a logged expense.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | The expense id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BooksApi.UpdateBookExpense(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BooksApi.UpdateBookExpense``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBookExpense`: CreateBookExpenseResponse
    fmt.Fprintf(os.Stdout, "Response from `BooksApi.UpdateBookExpense`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | The expense id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBookExpenseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBookExpenseResponse**

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
    resp, r, err := apiClient.BooksApi.UpdateBookTaxProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BooksApi.UpdateBookTaxProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBookTaxProfile`: GetBookTaxProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `BooksApi.UpdateBookTaxProfile`: %v\n", resp)
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

