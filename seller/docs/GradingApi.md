# \GradingApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGradingRunReserve**](GradingApi.md#CreateGradingRunReserve) | **Post** /v1/grading/runs/{id}/reserve | Hold a slot in a Grading Run. Free now; charged on your saved card only if the run locks. Full-access token and Idempotency-Key required.
[**CreateGradingSubmissionCancel**](GradingApi.md#CreateGradingSubmissionCancel) | **Post** /v1/grading/submissions/{id}/cancel | Pull a reserved card out of its run. Only while it is reserved; nothing was charged.
[**CreateGradingSubmissionShipped**](GradingApi.md#CreateGradingSubmissionShipped) | **Post** /v1/grading/submissions/{id}/shipped | Say you mailed a charged card to the hub, with optional tracking.
[**CreateGradingVerifyCert**](GradingApi.md#CreateGradingVerifyCert) | **Post** /v1/grading/verify-cert | Look a cert number up with its grader. found&#x3D;false when the grader has no record.
[**CreateVaultMarketGradingOrder**](GradingApi.md#CreateVaultMarketGradingOrder) | **Post** /v1/vault-markets/grading/orders | Start a grading submission: cards (vault units or inventory items) to one grader at one service. Priced from the price book; vault units are held. Nothing is charged until you pay.
[**CreateVaultMarketGradingOrderCancel**](GradingApi.md#CreateVaultMarketGradingOrderCancel) | **Post** /v1/vault-markets/grading/orders/{id}/cancel | Cancel a draft submission; its vault units go back on their shelf. A paid submission cannot be cancelled here.
[**CreateVaultMarketGradingOrderPay**](GradingApi.md#CreateVaultMarketGradingOrderPay) | **Post** /v1/vault-markets/grading/orders/{id}/pay | Pay for a draft submission. savedCard: true charges your saved card (paid, or requires_action with a client secret for 3-D Secure). Full-access token and Idempotency-Key required.
[**CreateVaultMarketGradingOrderPayConfirm**](GradingApi.md#CreateVaultMarketGradingOrderPayConfirm) | **Post** /v1/vault-markets/grading/orders/{id}/pay/confirm | After a 3-D Secure step or a card form: confirm the PaymentIntent with Stripe and mark the submission paid.
[**GetGradingPriceBook**](GradingApi.md#GetGradingPriceBook) | **Get** /v1/grading/price-book | The grading price book: every grader, its services (per-card price, minimum cards, max declared value, turnaround) and add-ons, plus Crossly handling. The same data the dashboard prices from.
[**GetGradingWorth**](GradingApi.md#GetGradingWorth) | **Get** /v1/grading/worth/{inventoryItemId} | Worth grading? Raw value vs what each grade is worth at each grader, against the cost. Never predicts the grade.
[**GetVaultMarketGradingOrder**](GradingApi.md#GetVaultMarketGradingOrder) | **Get** /v1/vault-markets/grading/orders/{id} | One grading submission: its timeline, every card with its grade, cert number and photos, the costs and the tracking.
[**ListGradingRateCard**](GradingApi.md#ListGradingRateCard) | **Get** /v1/grading/rate-card | Every grader and tier Crossly can send to, with the per-card price in cents (null &#x3D; paused at the grader).
[**ListGradingRuns**](GradingApi.md#ListGradingRuns) | **Get** /v1/grading/runs | Grading Runs taking reservations now, soonest cutoff first. meta.open is false while runs are switched off.
[**ListGradingSubmissions**](GradingApi.md#ListGradingSubmissions) | **Get** /v1/grading/submissions | Your cards out for grading (runs and Direct), newest first. active&#x3D;1 keeps only the ones still in flight.
[**ListVaultMarketGradingCandidates**](GradingApi.md#ListVaultMarketGradingCandidates) | **Get** /v1/vault-markets/grading/candidates | What you can send to grading: raw units already in your vault and inventory items, each with a declared value to start from and the reason when one cannot go.
[**ListVaultMarketGradingOrders**](GradingApi.md#ListVaultMarketGradingOrders) | **Get** /v1/vault-markets/grading/orders | Your grading submissions, newest first. Filter by status, active, grader, tier and dates; q searches the number, card names and cert numbers.



## CreateGradingRunReserve

> IdResponse CreateGradingRunReserve(ctx, id).Execute()

Hold a slot in a Grading Run. Free now; charged on your saved card only if the run locks. Full-access token and Idempotency-Key required.

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
    resp, r, err := apiClient.GradingApi.CreateGradingRunReserve(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.CreateGradingRunReserve``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGradingRunReserve`: IdResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.CreateGradingRunReserve`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGradingRunReserveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**IdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateGradingSubmissionCancel

> OkResponse CreateGradingSubmissionCancel(ctx, id).Execute()

Pull a reserved card out of its run. Only while it is reserved; nothing was charged.

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
    resp, r, err := apiClient.GradingApi.CreateGradingSubmissionCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.CreateGradingSubmissionCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGradingSubmissionCancel`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.CreateGradingSubmissionCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGradingSubmissionCancelRequest struct via the builder pattern


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


## CreateGradingSubmissionShipped

> OkResponse CreateGradingSubmissionShipped(ctx, id).Execute()

Say you mailed a charged card to the hub, with optional tracking.

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
    resp, r, err := apiClient.GradingApi.CreateGradingSubmissionShipped(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.CreateGradingSubmissionShipped``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGradingSubmissionShipped`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.CreateGradingSubmissionShipped`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGradingSubmissionShippedRequest struct via the builder pattern


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


## CreateGradingVerifyCert

> CreateGradingVerifyCertResponse CreateGradingVerifyCert(ctx).Execute()

Look a cert number up with its grader. found=false when the grader has no record.

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
    resp, r, err := apiClient.GradingApi.CreateGradingVerifyCert(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.CreateGradingVerifyCert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGradingVerifyCert`: CreateGradingVerifyCertResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.CreateGradingVerifyCert`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGradingVerifyCertRequest struct via the builder pattern


### Return type

**CreateGradingVerifyCertResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGradingOrder

> CreateVaultMarketGradingOrderResponse CreateVaultMarketGradingOrder(ctx).Execute()

Start a grading submission: cards (vault units or inventory items) to one grader at one service. Priced from the price book; vault units are held. Nothing is charged until you pay.

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
    resp, r, err := apiClient.GradingApi.CreateVaultMarketGradingOrder(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.CreateVaultMarketGradingOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGradingOrder`: CreateVaultMarketGradingOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.CreateVaultMarketGradingOrder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGradingOrderRequest struct via the builder pattern


### Return type

**CreateVaultMarketGradingOrderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGradingOrderCancel

> StatusResponse CreateVaultMarketGradingOrderCancel(ctx, id).Execute()

Cancel a draft submission; its vault units go back on their shelf. A paid submission cannot be cancelled here.

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
    resp, r, err := apiClient.GradingApi.CreateVaultMarketGradingOrderCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.CreateVaultMarketGradingOrderCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGradingOrderCancel`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.CreateVaultMarketGradingOrderCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGradingOrderCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StatusResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGradingOrderPay

> CreateVaultMarketGradingOrderPayResponse CreateVaultMarketGradingOrderPay(ctx, id).Execute()

Pay for a draft submission. savedCard: true charges your saved card (paid, or requires_action with a client secret for 3-D Secure). Full-access token and Idempotency-Key required.

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
    resp, r, err := apiClient.GradingApi.CreateVaultMarketGradingOrderPay(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.CreateVaultMarketGradingOrderPay``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGradingOrderPay`: CreateVaultMarketGradingOrderPayResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.CreateVaultMarketGradingOrderPay`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGradingOrderPayRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVaultMarketGradingOrderPayResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVaultMarketGradingOrderPayConfirm

> StatusResponse CreateVaultMarketGradingOrderPayConfirm(ctx, id).Execute()

After a 3-D Secure step or a card form: confirm the PaymentIntent with Stripe and mark the submission paid.

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
    resp, r, err := apiClient.GradingApi.CreateVaultMarketGradingOrderPayConfirm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.CreateVaultMarketGradingOrderPayConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVaultMarketGradingOrderPayConfirm`: StatusResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.CreateVaultMarketGradingOrderPayConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVaultMarketGradingOrderPayConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StatusResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetGradingPriceBook

> GetGradingPriceBookResponse GetGradingPriceBook(ctx).Execute()

The grading price book: every grader, its services (per-card price, minimum cards, max declared value, turnaround) and add-ons, plus Crossly handling. The same data the dashboard prices from.

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
    resp, r, err := apiClient.GradingApi.GetGradingPriceBook(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.GetGradingPriceBook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGradingPriceBook`: GetGradingPriceBookResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.GetGradingPriceBook`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetGradingPriceBookRequest struct via the builder pattern


### Return type

**GetGradingPriceBookResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetGradingWorth

> GetGradingWorthResponse GetGradingWorth(ctx, inventoryItemId).Execute()

Worth grading? Raw value vs what each grade is worth at each grader, against the cost. Never predicts the grade.

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
    inventoryItemId := "inventoryItemId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GradingApi.GetGradingWorth(context.Background(), inventoryItemId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.GetGradingWorth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGradingWorth`: GetGradingWorthResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.GetGradingWorth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**inventoryItemId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGradingWorthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetGradingWorthResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetVaultMarketGradingOrder

> GetVaultMarketGradingOrderResponse GetVaultMarketGradingOrder(ctx, id).Execute()

One grading submission: its timeline, every card with its grade, cert number and photos, the costs and the tracking.

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
    resp, r, err := apiClient.GradingApi.GetVaultMarketGradingOrder(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.GetVaultMarketGradingOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetVaultMarketGradingOrder`: GetVaultMarketGradingOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.GetVaultMarketGradingOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetVaultMarketGradingOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetVaultMarketGradingOrderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListGradingRateCard

> ListGradingRateCardResponse ListGradingRateCard(ctx).Execute()

Every grader and tier Crossly can send to, with the per-card price in cents (null = paused at the grader).

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
    resp, r, err := apiClient.GradingApi.ListGradingRateCard(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.ListGradingRateCard``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGradingRateCard`: ListGradingRateCardResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.ListGradingRateCard`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGradingRateCardRequest struct via the builder pattern


### Return type

**ListGradingRateCardResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListGradingRuns

> ListGradingRunsResponse ListGradingRuns(ctx).Execute()

Grading Runs taking reservations now, soonest cutoff first. meta.open is false while runs are switched off.

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
    resp, r, err := apiClient.GradingApi.ListGradingRuns(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.ListGradingRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGradingRuns`: ListGradingRunsResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.ListGradingRuns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListGradingRunsRequest struct via the builder pattern


### Return type

**ListGradingRunsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListGradingSubmissions

> ListGradingSubmissionsResponse ListGradingSubmissions(ctx).Active(active).Execute()

Your cards out for grading (runs and Direct), newest first. active=1 keeps only the ones still in flight.

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
    active := "active_example" // string | Only cards that have not finished. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.GradingApi.ListGradingSubmissions(context.Background()).Active(active).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.ListGradingSubmissions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGradingSubmissions`: ListGradingSubmissionsResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.ListGradingSubmissions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGradingSubmissionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **string** | Only cards that have not finished. | 

### Return type

**ListGradingSubmissionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketGradingCandidates

> ListVaultMarketGradingCandidatesResponse ListVaultMarketGradingCandidates(ctx).Execute()

What you can send to grading: raw units already in your vault and inventory items, each with a declared value to start from and the reason when one cannot go.

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
    resp, r, err := apiClient.GradingApi.ListVaultMarketGradingCandidates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.ListVaultMarketGradingCandidates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketGradingCandidates`: ListVaultMarketGradingCandidatesResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.ListVaultMarketGradingCandidates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketGradingCandidatesRequest struct via the builder pattern


### Return type

**ListVaultMarketGradingCandidatesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVaultMarketGradingOrders

> ListVaultMarketGradingOrdersResponse ListVaultMarketGradingOrders(ctx).Execute()

Your grading submissions, newest first. Filter by status, active, grader, tier and dates; q searches the number, card names and cert numbers.

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
    resp, r, err := apiClient.GradingApi.ListVaultMarketGradingOrders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `GradingApi.ListVaultMarketGradingOrders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVaultMarketGradingOrders`: ListVaultMarketGradingOrdersResponse
    fmt.Fprintf(os.Stdout, "Response from `GradingApi.ListVaultMarketGradingOrders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVaultMarketGradingOrdersRequest struct via the builder pattern


### Return type

**ListVaultMarketGradingOrdersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

