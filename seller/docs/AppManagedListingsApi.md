# \AppManagedListingsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateListingCommandAck**](AppManagedListingsApi.md#CreateListingCommandAck) | **Post** /v1/listing-commands/{id}/ack | Acknowledge a command (you will act on it). Extends its deadline: 24 h for listing commands; tracking 2 h, cancel 4 h, refund 24 h, message 1 h.
[**CreateListingCommandComplete**](AppManagedListingsApi.md#CreateListingCommandComplete) | **Post** /v1/listing-commands/{id}/complete | Report a command done on the marketplace. relist: newExternalId when a new id was issued; refund_order: platformRefundId + amount (required); submit_tracking: trackingNumber/carrier as the marketplace holds them; message_buyer: platformMessageId.
[**CreateListingCommandFail**](AppManagedListingsApi.md#CreateListingCommandFail) | **Post** /v1/listing-commands/{id}/fail | Report a command failed, with a reason the seller will see.
[**CreateListingEvent**](AppManagedListingsApi.md#CreateListingEvent) | **Post** /v1/listings/{id}/events | Report a marketplace event (item.updated / ended / sold / relisted) for a listing your app manages.
[**CreateListingEvent2**](AppManagedListingsApi.md#CreateListingEvent2) | **Post** /v1/listings/events | Report up to 100 marketplace events keyed by platform + externalId. One bad event never fails the rest.
[**CreateListingLink**](AppManagedListingsApi.md#CreateListingLink) | **Post** /v1/listings/link | Link a live marketplace listing your app already manages into Crossly (no re-listing). Idempotent.
[**CreateListingLinkBulk**](AppManagedListingsApi.md#CreateListingLinkBulk) | **Post** /v1/listings/link/bulk | Link up to 100 live marketplace listings. Each item succeeds or fails on its own; re-sending is safe.
[**CreateListingPlatformReclone**](AppManagedListingsApi.md#CreateListingPlatformReclone) | **Post** /v1/listings/{id}/platforms/{platform}/reclone | End this marketplace item and relist it under Crossly. Requires consent: true.
[**CreateListingPlatformRecloneCancel**](AppManagedListingsApi.md#CreateListingPlatformRecloneCancel) | **Post** /v1/listings/{id}/platforms/{platform}/reclone/cancel | Cancel a relist still waiting for you to end the old item.
[**CreateListingPlatformRecloneSellerEnded**](AppManagedListingsApi.md#CreateListingPlatformRecloneSellerEnded) | **Post** /v1/listings/{id}/platforms/{platform}/reclone/seller-ended | You ended the old item on the marketplace yourself (Done).
[**CreateListingPlatformTakeOver**](AppManagedListingsApi.md#CreateListingPlatformTakeOver) | **Post** /v1/listings/{id}/platforms/{platform}/take-over | Take over an app-managed item: Crossly acts on the marketplace itself (needs your own connection).
[**CreateListingUnlink**](AppManagedListingsApi.md#CreateListingUnlink) | **Post** /v1/listings/{id}/unlink | Stop the app managing this listing. It stays live everywhere; the app&#39;s Crossly Shop fee ends for future sales.
[**CreateOrderEvent**](AppManagedListingsApi.md#CreateOrderEvent) | **Post** /v1/orders/events | Report up to 100 order events (shipped, cancelled, refunded, delivered) keyed by platform + platformOrderId. Each completes the matching open order command. One bad event never fails the rest.
[**CreateOrderTakeOver**](AppManagedListingsApi.md#CreateOrderTakeOver) | **Post** /v1/orders/{id}/take-over | Take over the app-managed item behind an order; its actions then go through your own connection.
[**GetListingManagement**](AppManagedListingsApi.md#GetListingManagement) | **Get** /v1/listings/{id}/management | Which app manages this listing (null when none). Its fee is charged only on Crossly Shop sales, never on eBay or other marketplaces.
[**GetListingPlatformReclone**](AppManagedListingsApi.md#GetListingPlatformReclone) | **Get** /v1/listings/{id}/platforms/{platform}/reclone | The latest relist-under-Crossly run for this listing on a platform (null when none).
[**ListListingCommands**](AppManagedListingsApi.md#ListListingCommands) | **Get** /v1/listing-commands | Your app&#39;s listing (delist/update/relist) and order (submit_tracking, remove_tracking, cancel_order, refund_order, message_buyer) commands for this seller (default: open ones), oldest first.



## CreateListingCommandAck

> ListListingCommandsItem CreateListingCommandAck(ctx, id).Execute()

Acknowledge a command (you will act on it). Extends its deadline: 24 h for listing commands; tracking 2 h, cancel 4 h, refund 24 h, message 1 h.

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
    resp, r, err := apiClient.AppManagedListingsApi.CreateListingCommandAck(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.CreateListingCommandAck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingCommandAck`: ListListingCommandsItem
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.CreateListingCommandAck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingCommandAckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListListingCommandsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingCommandComplete

> ListListingCommandsItem CreateListingCommandComplete(ctx, id).Execute()

Report a command done on the marketplace. relist: newExternalId when a new id was issued; refund_order: platformRefundId + amount (required); submit_tracking: trackingNumber/carrier as the marketplace holds them; message_buyer: platformMessageId.

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
    resp, r, err := apiClient.AppManagedListingsApi.CreateListingCommandComplete(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.CreateListingCommandComplete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingCommandComplete`: ListListingCommandsItem
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.CreateListingCommandComplete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingCommandCompleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListListingCommandsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingCommandFail

> ListListingCommandsItem CreateListingCommandFail(ctx, id).Execute()

Report a command failed, with a reason the seller will see.

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
    resp, r, err := apiClient.AppManagedListingsApi.CreateListingCommandFail(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.CreateListingCommandFail``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingCommandFail`: ListListingCommandsItem
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.CreateListingCommandFail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingCommandFailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListListingCommandsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingEvent

> CreateListingEventResponse CreateListingEvent(ctx, id).Execute()

Report a marketplace event (item.updated / ended / sold / relisted) for a listing your app manages.



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
    resp, r, err := apiClient.AppManagedListingsApi.CreateListingEvent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.CreateListingEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingEvent`: CreateListingEventResponse
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.CreateListingEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateListingEventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingEvent2

> CreateListingEvent2Response CreateListingEvent2(ctx).Execute()

Report up to 100 marketplace events keyed by platform + externalId. One bad event never fails the rest.

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
    resp, r, err := apiClient.AppManagedListingsApi.CreateListingEvent2(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.CreateListingEvent2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingEvent2`: CreateListingEvent2Response
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.CreateListingEvent2`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingEvent2Request struct via the builder pattern


### Return type

**CreateListingEvent2Response**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingLink

> CreateListingLinkResponse CreateListingLink(ctx).Execute()

Link a live marketplace listing your app already manages into Crossly (no re-listing). Idempotent.



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
    resp, r, err := apiClient.AppManagedListingsApi.CreateListingLink(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.CreateListingLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingLink`: CreateListingLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.CreateListingLink`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingLinkRequest struct via the builder pattern


### Return type

**CreateListingLinkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingLinkBulk

> CreateListingLinkBulkResponse CreateListingLinkBulk(ctx).Execute()

Link up to 100 live marketplace listings. Each item succeeds or fails on its own; re-sending is safe.

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
    resp, r, err := apiClient.AppManagedListingsApi.CreateListingLinkBulk(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.CreateListingLinkBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingLinkBulk`: CreateListingLinkBulkResponse
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.CreateListingLinkBulk`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingLinkBulkRequest struct via the builder pattern


### Return type

**CreateListingLinkBulkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingPlatformReclone

> CreateListingPlatformRecloneResponse CreateListingPlatformReclone(ctx, id, platform).Execute()

End this marketplace item and relist it under Crossly. Requires consent: true.



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
    platform := "platform_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppManagedListingsApi.CreateListingPlatformReclone(context.Background(), id, platform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.CreateListingPlatformReclone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingPlatformReclone`: CreateListingPlatformRecloneResponse
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.CreateListingPlatformReclone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingPlatformRecloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateListingPlatformRecloneResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingPlatformRecloneCancel

> CreateListingPlatformRecloneResponse CreateListingPlatformRecloneCancel(ctx, id, platform).Execute()

Cancel a relist still waiting for you to end the old item.

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
    platform := "platform_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppManagedListingsApi.CreateListingPlatformRecloneCancel(context.Background(), id, platform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.CreateListingPlatformRecloneCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingPlatformRecloneCancel`: CreateListingPlatformRecloneResponse
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.CreateListingPlatformRecloneCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingPlatformRecloneCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateListingPlatformRecloneResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingPlatformRecloneSellerEnded

> CreateListingPlatformRecloneResponse CreateListingPlatformRecloneSellerEnded(ctx, id, platform).Execute()

You ended the old item on the marketplace yourself (Done).

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
    platform := "platform_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppManagedListingsApi.CreateListingPlatformRecloneSellerEnded(context.Background(), id, platform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.CreateListingPlatformRecloneSellerEnded``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingPlatformRecloneSellerEnded`: CreateListingPlatformRecloneResponse
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.CreateListingPlatformRecloneSellerEnded`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingPlatformRecloneSellerEndedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateListingPlatformRecloneResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingPlatformTakeOver

> CreateListingPlatformTakeOverResponse CreateListingPlatformTakeOver(ctx, id, platform).Execute()

Take over an app-managed item: Crossly acts on the marketplace itself (needs your own connection).



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
    platform := "platform_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppManagedListingsApi.CreateListingPlatformTakeOver(context.Background(), id, platform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.CreateListingPlatformTakeOver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingPlatformTakeOver`: CreateListingPlatformTakeOverResponse
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.CreateListingPlatformTakeOver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingPlatformTakeOverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateListingPlatformTakeOverResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateListingUnlink

> CreateListingUnlinkResponse CreateListingUnlink(ctx, id).Execute()

Stop the app managing this listing. It stays live everywhere; the app's Crossly Shop fee ends for future sales.

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
    resp, r, err := apiClient.AppManagedListingsApi.CreateListingUnlink(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.CreateListingUnlink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateListingUnlink`: CreateListingUnlinkResponse
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.CreateListingUnlink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateListingUnlinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateListingUnlinkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderEvent

> CreateOrderEventResponse CreateOrderEvent(ctx).Execute()

Report up to 100 order events (shipped, cancelled, refunded, delivered) keyed by platform + platformOrderId. Each completes the matching open order command. One bad event never fails the rest.

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
    resp, r, err := apiClient.AppManagedListingsApi.CreateOrderEvent(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.CreateOrderEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderEvent`: CreateOrderEventResponse
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.CreateOrderEvent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderEventRequest struct via the builder pattern


### Return type

**CreateOrderEventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateOrderTakeOver

> CreateOrderTakeOverResponse CreateOrderTakeOver(ctx, id).Execute()

Take over the app-managed item behind an order; its actions then go through your own connection.

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
    resp, r, err := apiClient.AppManagedListingsApi.CreateOrderTakeOver(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.CreateOrderTakeOver``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateOrderTakeOver`: CreateOrderTakeOverResponse
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.CreateOrderTakeOver`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOrderTakeOverRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateOrderTakeOverResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetListingManagement

> GetListingManagementResponse GetListingManagement(ctx, id).Execute()

Which app manages this listing (null when none). Its fee is charged only on Crossly Shop sales, never on eBay or other marketplaces.

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
    resp, r, err := apiClient.AppManagedListingsApi.GetListingManagement(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.GetListingManagement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListingManagement`: GetListingManagementResponse
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.GetListingManagement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingManagementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetListingManagementResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetListingPlatformReclone

> GetListingPlatformRecloneResponse GetListingPlatformReclone(ctx, id, platform).Execute()

The latest relist-under-Crossly run for this listing on a platform (null when none).

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
    platform := "platform_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppManagedListingsApi.GetListingPlatformReclone(context.Background(), id, platform).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.GetListingPlatformReclone``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetListingPlatformReclone`: GetListingPlatformRecloneResponse
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.GetListingPlatformReclone`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**platform** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetListingPlatformRecloneRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetListingPlatformRecloneResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListListingCommands

> ListListingCommandsResponse ListListingCommands(ctx).State(state).Type_(type_).Limit(limit).Page(page).Execute()

Your app's listing (delist/update/relist) and order (submit_tracking, remove_tracking, cancel_order, refund_order, message_buyer) commands for this seller (default: open ones), oldest first.

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
    state := "state_example" // string | Comma-separated: pending, acked, completed, failed, expired, cancelled (optional)
    type_ := "type__example" // string | Comma-separated: delist, update, relist, submit_tracking, remove_tracking, cancel_order, refund_order, message_buyer (optional)
    limit := int32(56) // int32 |  (optional)
    page := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppManagedListingsApi.ListListingCommands(context.Background()).State(state).Type_(type_).Limit(limit).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppManagedListingsApi.ListListingCommands``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListListingCommands`: ListListingCommandsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppManagedListingsApi.ListListingCommands`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListListingCommandsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **state** | **string** | Comma-separated: pending, acked, completed, failed, expired, cancelled | 
 **type_** | **string** | Comma-separated: delist, update, relist, submit_tracking, remove_tracking, cancel_order, refund_order, message_buyer | 
 **limit** | **int32** |  | 
 **page** | **int32** |  | 

### Return type

**ListListingCommandsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

