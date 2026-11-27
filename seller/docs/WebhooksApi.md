# \WebhooksApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWebhook**](WebhooksApi.md#CreateWebhook) | **Post** /v1/webhooks | Register a webhook endpoint (returns the signing secret once).
[**CreateWebhookDeliveryRedeliver**](WebhooksApi.md#CreateWebhookDeliveryRedeliver) | **Post** /v1/webhooks/{id}/deliveries/{deliveryId}/redeliver | Redeliver one past delivery: a new delivery (new Crossly-Delivery-Id) of the same event id and payload.
[**CreateWebhookTest**](WebhooksApi.md#CreateWebhookTest) | **Post** /v1/webhooks/{id}/test | Fire a test delivery to one webhook: test.ping, or a sample of any event (body {\&quot;event\&quot;:\&quot;order.created\&quot;}, payload marked test: true).
[**DeleteWebhook**](WebhooksApi.md#DeleteWebhook) | **Delete** /v1/webhooks/{id} | Delete a webhook endpoint.
[**GetWebhookStream**](WebhooksApi.md#GetWebhookStream) | **Get** /v1/webhooks/stream | Stream this account&#39;s webhook events as they happen (SSE).
[**ListWebhookDeliveries**](WebhooksApi.md#ListWebhookDeliveries) | **Get** /v1/webhooks/{id}/deliveries | List deliveries to one webhook (30 days, newest first): event, status, attempts, last HTTP status and error, payload.
[**ListWebhooks**](WebhooksApi.md#ListWebhooks) | **Get** /v1/webhooks | List your registered webhook endpoints.



## CreateWebhook

> CreateWebhookResponse CreateWebhook(ctx).Execute()

Register a webhook endpoint (returns the signing secret once).

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
    resp, r, err := apiClient.WebhooksApi.CreateWebhook(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.CreateWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhook`: CreateWebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.CreateWebhook`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookRequest struct via the builder pattern


### Return type

**CreateWebhookResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWebhookDeliveryRedeliver

> CreateWebhookDeliveryRedeliverResponse CreateWebhookDeliveryRedeliver(ctx, id, deliveryId).Execute()

Redeliver one past delivery: a new delivery (new Crossly-Delivery-Id) of the same event id and payload.

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
    deliveryId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.CreateWebhookDeliveryRedeliver(context.Background(), id, deliveryId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.CreateWebhookDeliveryRedeliver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhookDeliveryRedeliver`: CreateWebhookDeliveryRedeliverResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.CreateWebhookDeliveryRedeliver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**deliveryId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookDeliveryRedeliverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWebhookDeliveryRedeliverResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWebhookTest

> CreateWebhookTestResponse CreateWebhookTest(ctx, id).Execute()

Fire a test delivery to one webhook: test.ping, or a sample of any event (body {\"event\":\"order.created\"}, payload marked test: true).

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
    resp, r, err := apiClient.WebhooksApi.CreateWebhookTest(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.CreateWebhookTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWebhookTest`: CreateWebhookTestResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.CreateWebhookTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWebhookTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWebhookTestResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWebhook

> OkResponse DeleteWebhook(ctx, id).Execute()

Delete a webhook endpoint.

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
    resp, r, err := apiClient.WebhooksApi.DeleteWebhook(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.DeleteWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWebhook`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.DeleteWebhook`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWebhookRequest struct via the builder pattern


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


## GetWebhookStream

> string GetWebhookStream(ctx).Execute()

Stream this account's webhook events as they happen (SSE).



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
    resp, r, err := apiClient.WebhooksApi.GetWebhookStream(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.GetWebhookStream``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWebhookStream`: string
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.GetWebhookStream`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWebhookStreamRequest struct via the builder pattern


### Return type

**string**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/event-stream, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWebhookDeliveries

> ListWebhookDeliveriesResponse ListWebhookDeliveries(ctx, id).Status(status).Limit(limit).Execute()

List deliveries to one webhook (30 days, newest first): event, status, attempts, last HTTP status and error, payload.

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
    status := "status_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WebhooksApi.ListWebhookDeliveries(context.Background(), id).Status(status).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.ListWebhookDeliveries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhookDeliveries`: ListWebhookDeliveriesResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.ListWebhookDeliveries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhookDeliveriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** |  | 
 **limit** | **int32** |  | [default to 25]

### Return type

**ListWebhookDeliveriesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWebhooks

> ListWebhooksResponse ListWebhooks(ctx).Execute()

List your registered webhook endpoints.

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
    resp, r, err := apiClient.WebhooksApi.ListWebhooks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WebhooksApi.ListWebhooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWebhooks`: ListWebhooksResponse
    fmt.Fprintf(os.Stdout, "Response from `WebhooksApi.ListWebhooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWebhooksRequest struct via the builder pattern


### Return type

**ListWebhooksResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

