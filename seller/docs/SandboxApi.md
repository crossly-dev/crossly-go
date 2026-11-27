# \SandboxApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSandboxReset**](SandboxApi.md#CreateSandboxReset) | **Post** /v1/sandbox/reset | Restore the sandbox seller to its seed. Your grant and webhook subscriptions are kept.
[**CreateSandboxSimulateInventory**](SandboxApi.md#CreateSandboxSimulateInventory) | **Post** /v1/sandbox/simulate/inventory | Change stock (set, increase or decrease); the new quantity is pushed to every channel.
[**CreateSandboxSimulateListingSoldOut**](SandboxApi.md#CreateSandboxSimulateListingSoldOut) | **Post** /v1/sandbox/simulate/listing.sold_out | A listing sells out on a marketplace with no order (stock moves, other channels delist).
[**CreateSandboxSimulateOrder**](SandboxApi.md#CreateSandboxSimulateOrder) | **Post** /v1/sandbox/simulate/order | Simulate a sale through the real sale pipeline (order.created, sale.detected, stock and cross-channel updates).
[**CreateSandboxSimulateOrderAddress**](SandboxApi.md#CreateSandboxSimulateOrderAddress) | **Post** /v1/sandbox/simulate/order.address | The buyer address of an order simulated with address \&quot;delayed\&quot; arrives now (order.address_available).
[**CreateSandboxSimulateOrderCancelled**](SandboxApi.md#CreateSandboxSimulateOrderCancelled) | **Post** /v1/sandbox/simulate/order.cancelled | The buyer or the marketplace cancels an order (order.cancelled).
[**CreateSandboxSimulateOrderDelivered**](SandboxApi.md#CreateSandboxSimulateOrderDelivered) | **Post** /v1/sandbox/simulate/order.delivered | The carrier delivers a shipped order (order.delivered).
[**CreateSandboxSimulateRefund**](SandboxApi.md#CreateSandboxSimulateRefund) | **Post** /v1/sandbox/simulate/refund | Refund an order in full or in part (refund.issued).
[**GetSandbox**](SandboxApi.md#GetSandbox) | **Get** /v1/sandbox | The sandbox seller this token acts on, its simulated channels and the failure modes.
[**ListSandboxEvents**](SandboxApi.md#ListSandboxEvents) | **Get** /v1/sandbox/events | The simulated marketplaces&#39; call log: verb, channel, outcome, latency.



## CreateSandboxReset

> CreateSandboxResetResponse CreateSandboxReset(ctx).Execute()

Restore the sandbox seller to its seed. Your grant and webhook subscriptions are kept.

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
    resp, r, err := apiClient.SandboxApi.CreateSandboxReset(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxApi.CreateSandboxReset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSandboxReset`: CreateSandboxResetResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxApi.CreateSandboxReset`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSandboxResetRequest struct via the builder pattern


### Return type

**CreateSandboxResetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateSandboxSimulateInventory

> CreateInboxConversationBulkResponse CreateSandboxSimulateInventory(ctx).Execute()

Change stock (set, increase or decrease); the new quantity is pushed to every channel.

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
    resp, r, err := apiClient.SandboxApi.CreateSandboxSimulateInventory(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxApi.CreateSandboxSimulateInventory``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSandboxSimulateInventory`: CreateInboxConversationBulkResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxApi.CreateSandboxSimulateInventory`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSandboxSimulateInventoryRequest struct via the builder pattern


### Return type

**CreateInboxConversationBulkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateSandboxSimulateListingSoldOut

> CreateSandboxSimulateListingSoldOutResponse CreateSandboxSimulateListingSoldOut(ctx).Execute()

A listing sells out on a marketplace with no order (stock moves, other channels delist).

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
    resp, r, err := apiClient.SandboxApi.CreateSandboxSimulateListingSoldOut(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxApi.CreateSandboxSimulateListingSoldOut``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSandboxSimulateListingSoldOut`: CreateSandboxSimulateListingSoldOutResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxApi.CreateSandboxSimulateListingSoldOut`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSandboxSimulateListingSoldOutRequest struct via the builder pattern


### Return type

**CreateSandboxSimulateListingSoldOutResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateSandboxSimulateOrder

> CreateSandboxSimulateOrderResponse CreateSandboxSimulateOrder(ctx).Execute()

Simulate a sale through the real sale pipeline (order.created, sale.detected, stock and cross-channel updates).

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
    resp, r, err := apiClient.SandboxApi.CreateSandboxSimulateOrder(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxApi.CreateSandboxSimulateOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSandboxSimulateOrder`: CreateSandboxSimulateOrderResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxApi.CreateSandboxSimulateOrder`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSandboxSimulateOrderRequest struct via the builder pattern


### Return type

**CreateSandboxSimulateOrderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateSandboxSimulateOrderAddress

> CreateSandboxSimulateOrderAddressResponse CreateSandboxSimulateOrderAddress(ctx).Execute()

The buyer address of an order simulated with address \"delayed\" arrives now (order.address_available).

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
    resp, r, err := apiClient.SandboxApi.CreateSandboxSimulateOrderAddress(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxApi.CreateSandboxSimulateOrderAddress``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSandboxSimulateOrderAddress`: CreateSandboxSimulateOrderAddressResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxApi.CreateSandboxSimulateOrderAddress`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSandboxSimulateOrderAddressRequest struct via the builder pattern


### Return type

**CreateSandboxSimulateOrderAddressResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateSandboxSimulateOrderCancelled

> CreateSandboxSimulateOrderCancelledResponse CreateSandboxSimulateOrderCancelled(ctx).Execute()

The buyer or the marketplace cancels an order (order.cancelled).

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
    resp, r, err := apiClient.SandboxApi.CreateSandboxSimulateOrderCancelled(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxApi.CreateSandboxSimulateOrderCancelled``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSandboxSimulateOrderCancelled`: CreateSandboxSimulateOrderCancelledResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxApi.CreateSandboxSimulateOrderCancelled`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSandboxSimulateOrderCancelledRequest struct via the builder pattern


### Return type

**CreateSandboxSimulateOrderCancelledResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateSandboxSimulateOrderDelivered

> CreateSandboxSimulateOrderDeliveredResponse CreateSandboxSimulateOrderDelivered(ctx).Execute()

The carrier delivers a shipped order (order.delivered).

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
    resp, r, err := apiClient.SandboxApi.CreateSandboxSimulateOrderDelivered(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxApi.CreateSandboxSimulateOrderDelivered``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSandboxSimulateOrderDelivered`: CreateSandboxSimulateOrderDeliveredResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxApi.CreateSandboxSimulateOrderDelivered`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSandboxSimulateOrderDeliveredRequest struct via the builder pattern


### Return type

**CreateSandboxSimulateOrderDeliveredResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateSandboxSimulateRefund

> CreateSandboxSimulateRefundResponse CreateSandboxSimulateRefund(ctx).Execute()

Refund an order in full or in part (refund.issued).

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
    resp, r, err := apiClient.SandboxApi.CreateSandboxSimulateRefund(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxApi.CreateSandboxSimulateRefund``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSandboxSimulateRefund`: CreateSandboxSimulateRefundResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxApi.CreateSandboxSimulateRefund`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSandboxSimulateRefundRequest struct via the builder pattern


### Return type

**CreateSandboxSimulateRefundResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetSandbox

> GetSandboxResponse GetSandbox(ctx).Execute()

The sandbox seller this token acts on, its simulated channels and the failure modes.

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
    resp, r, err := apiClient.SandboxApi.GetSandbox(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxApi.GetSandbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSandbox`: GetSandboxResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxApi.GetSandbox`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSandboxRequest struct via the builder pattern


### Return type

**GetSandboxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListSandboxEvents

> ListSandboxEventsResponse ListSandboxEvents(ctx).Execute()

The simulated marketplaces' call log: verb, channel, outcome, latency.

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
    resp, r, err := apiClient.SandboxApi.ListSandboxEvents(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SandboxApi.ListSandboxEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSandboxEvents`: ListSandboxEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `SandboxApi.ListSandboxEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListSandboxEventsRequest struct via the builder pattern


### Return type

**ListSandboxEventsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

