# \BrandingApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBrandingDrop**](BrandingApi.md#CreateBrandingDrop) | **Post** /v1/branding/drops | Create a drop as a draft. &#x60;config&#x60; is the full drop config (what, where: shop|world|live|storefront|featured, format, access, limits, pricing, schedule, promotion, fulfilment). Returns rule problems to fix before publishing.
[**CreateBrandingDropControl**](BrandingApi.md#CreateBrandingDropControl) | **Post** /v1/branding/drops/{id}/controls/{control} | Live controls: pause, resume, extend (body endsAt), release-wave, end, draw (raffle).
[**CreateBrandingDropPublish**](BrandingApi.md#CreateBrandingDropPublish) | **Post** /v1/branding/drops/{id}/publish | Publish: scheduled, or pending_review for a brand partner without self-publish. Brand partners need a verified business on Stripe (403 business_verification_required).
[**CreateBrandingEventCancel**](BrandingApi.md#CreateBrandingEventCancel) | **Post** /v1/branding/events/{id}/cancel | Cancel a published event with the reason attendees see.
[**CreateBrandingEventCheckIn**](BrandingApi.md#CreateBrandingEventCheckIn) | **Post** /v1/branding/events/{id}/check-in | Check a ticket in at the door by its QR token.
[**CreateBrandingEventMessage**](BrandingApi.md#CreateBrandingEventMessage) | **Post** /v1/branding/events/{id}/messages | Message the event&#39;s attendees: all, going, or ticketed.
[**DeleteBrandingDrop**](BrandingApi.md#DeleteBrandingDrop) | **Delete** /v1/branding/drops/{id} | Delete a draft, or cancel a scheduled drop. A live drop is ended with the end control.
[**GetBrandingBrand**](BrandingApi.md#GetBrandingBrand) | **Get** /v1/branding/brand | Your brand profile (approved brand partners).
[**GetBrandingBusinessVerification**](BrandingApi.md#GetBrandingBusinessVerification) | **Get** /v1/branding/business-verification | Whether your Stripe account is a verified business (company, EIN/tax id verified, payouts on), with reasons and the page to fix it.
[**GetBrandingDrop**](BrandingApi.md#GetBrandingDrop) | **Get** /v1/branding/drops/{id} | One of your drops.
[**GetBrandingEvent**](BrandingApi.md#GetBrandingEvent) | **Get** /v1/branding/events | Your events: upcoming grouped by date (today, this_week, later), or past.
[**GetBrandingSetting**](BrandingApi.md#GetBrandingSetting) | **Get** /v1/branding/settings | The store’s Branding defaults new drops start from: timezone, notifyFollowers, botQueue, botCaptcha, cartHoldMinutes, leftovers.
[**ListBrandingDrops**](BrandingApi.md#ListBrandingDrops) | **Get** /v1/branding/drops | Your drops, newest first, with config, waves, counts (entries, holding, sold, buyers) and phase. ?scope&#x3D;upcoming|live|past|drafts|all.
[**UpdateBrandingBrand**](BrandingApi.md#UpdateBrandingBrand) | **Patch** /v1/branding/brand | Edit your brand profile: logoUrl, bannerUrl, story, links.
[**UpdateBrandingDrop**](BrandingApi.md#UpdateBrandingDrop) | **Patch** /v1/branding/drops/{id} | Edit a drop that is not live yet. A published drop refuses rule problems (400 drop_config_problems).
[**UpdateBrandingEventSetting**](BrandingApi.md#UpdateBrandingEventSetting) | **Put** /v1/branding/events/{id}/settings | Save an event&#39;s full settings (where, access, reminders, check-in, recording, linked drops) and ticket tiers.
[**UpdateBrandingSetting**](BrandingApi.md#UpdateBrandingSetting) | **Patch** /v1/branding/settings | Change the store’s Branding defaults (any subset of the fields).



## CreateBrandingDrop

> CreateBrandingDropResponse CreateBrandingDrop(ctx).Execute()

Create a drop as a draft. `config` is the full drop config (what, where: shop|world|live|storefront|featured, format, access, limits, pricing, schedule, promotion, fulfilment). Returns rule problems to fix before publishing.

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
    resp, r, err := apiClient.BrandingApi.CreateBrandingDrop(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.CreateBrandingDrop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrandingDrop`: CreateBrandingDropResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.CreateBrandingDrop`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandingDropRequest struct via the builder pattern


### Return type

**CreateBrandingDropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBrandingDropControl

> GetBrandingDropResponse CreateBrandingDropControl(ctx, id, control).Execute()

Live controls: pause, resume, extend (body endsAt), release-wave, end, draw (raffle).

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
    control := "control_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.BrandingApi.CreateBrandingDropControl(context.Background(), id, control).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.CreateBrandingDropControl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrandingDropControl`: GetBrandingDropResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.CreateBrandingDropControl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**control** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandingDropControlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetBrandingDropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBrandingDropPublish

> CreateBrandingDropPublishResponse CreateBrandingDropPublish(ctx, id).Execute()

Publish: scheduled, or pending_review for a brand partner without self-publish. Brand partners need a verified business on Stripe (403 business_verification_required).

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
    resp, r, err := apiClient.BrandingApi.CreateBrandingDropPublish(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.CreateBrandingDropPublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrandingDropPublish`: CreateBrandingDropPublishResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.CreateBrandingDropPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandingDropPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBrandingDropPublishResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBrandingEventCancel

> CreateBrandingEventCancelResponse CreateBrandingEventCancel(ctx, id).Execute()

Cancel a published event with the reason attendees see.

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
    resp, r, err := apiClient.BrandingApi.CreateBrandingEventCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.CreateBrandingEventCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrandingEventCancel`: CreateBrandingEventCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.CreateBrandingEventCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandingEventCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBrandingEventCancelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBrandingEventCheckIn

> CreateBrandingEventCheckInResponse CreateBrandingEventCheckIn(ctx, id).Execute()

Check a ticket in at the door by its QR token.

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
    resp, r, err := apiClient.BrandingApi.CreateBrandingEventCheckIn(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.CreateBrandingEventCheckIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrandingEventCheckIn`: CreateBrandingEventCheckInResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.CreateBrandingEventCheckIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandingEventCheckInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBrandingEventCheckInResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateBrandingEventMessage

> CreateBrandingEventMessageResponse CreateBrandingEventMessage(ctx, id).Execute()

Message the event's attendees: all, going, or ticketed.

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
    resp, r, err := apiClient.BrandingApi.CreateBrandingEventMessage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.CreateBrandingEventMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBrandingEventMessage`: CreateBrandingEventMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.CreateBrandingEventMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateBrandingEventMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBrandingEventMessageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteBrandingDrop

> DeleteBrandingDropResponse DeleteBrandingDrop(ctx, id).Execute()

Delete a draft, or cancel a scheduled drop. A live drop is ended with the end control.

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
    resp, r, err := apiClient.BrandingApi.DeleteBrandingDrop(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.DeleteBrandingDrop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBrandingDrop`: DeleteBrandingDropResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.DeleteBrandingDrop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBrandingDropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteBrandingDropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBrandingBrand

> GetBrandingBrandResponse GetBrandingBrand(ctx).Execute()

Your brand profile (approved brand partners).

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
    resp, r, err := apiClient.BrandingApi.GetBrandingBrand(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.GetBrandingBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandingBrand`: GetBrandingBrandResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.GetBrandingBrand`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingBrandRequest struct via the builder pattern


### Return type

**GetBrandingBrandResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBrandingBusinessVerification

> GetBrandingBusinessVerificationResponse GetBrandingBusinessVerification(ctx).Execute()

Whether your Stripe account is a verified business (company, EIN/tax id verified, payouts on), with reasons and the page to fix it.

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
    resp, r, err := apiClient.BrandingApi.GetBrandingBusinessVerification(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.GetBrandingBusinessVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandingBusinessVerification`: GetBrandingBusinessVerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.GetBrandingBusinessVerification`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingBusinessVerificationRequest struct via the builder pattern


### Return type

**GetBrandingBusinessVerificationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBrandingDrop

> GetBrandingDropResponse GetBrandingDrop(ctx, id).Execute()

One of your drops.

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
    resp, r, err := apiClient.BrandingApi.GetBrandingDrop(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.GetBrandingDrop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandingDrop`: GetBrandingDropResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.GetBrandingDrop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingDropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetBrandingDropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBrandingEvent

> GetBrandingEventResponse GetBrandingEvent(ctx).Execute()

Your events: upcoming grouped by date (today, this_week, later), or past.

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
    resp, r, err := apiClient.BrandingApi.GetBrandingEvent(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.GetBrandingEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandingEvent`: GetBrandingEventResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.GetBrandingEvent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingEventRequest struct via the builder pattern


### Return type

**GetBrandingEventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetBrandingSetting

> GetBrandingSettingResponse GetBrandingSetting(ctx).Execute()

The store’s Branding defaults new drops start from: timezone, notifyFollowers, botQueue, botCaptcha, cartHoldMinutes, leftovers.

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
    resp, r, err := apiClient.BrandingApi.GetBrandingSetting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.GetBrandingSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBrandingSetting`: GetBrandingSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.GetBrandingSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBrandingSettingRequest struct via the builder pattern


### Return type

**GetBrandingSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListBrandingDrops

> ListBrandingDropsResponse ListBrandingDrops(ctx).Execute()

Your drops, newest first, with config, waves, counts (entries, holding, sold, buyers) and phase. ?scope=upcoming|live|past|drafts|all.

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
    resp, r, err := apiClient.BrandingApi.ListBrandingDrops(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.ListBrandingDrops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBrandingDrops`: ListBrandingDropsResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.ListBrandingDrops`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListBrandingDropsRequest struct via the builder pattern


### Return type

**ListBrandingDropsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBrandingBrand

> GetBrandingBrandResponse UpdateBrandingBrand(ctx).Execute()

Edit your brand profile: logoUrl, bannerUrl, story, links.

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
    resp, r, err := apiClient.BrandingApi.UpdateBrandingBrand(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.UpdateBrandingBrand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBrandingBrand`: GetBrandingBrandResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.UpdateBrandingBrand`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrandingBrandRequest struct via the builder pattern


### Return type

**GetBrandingBrandResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBrandingDrop

> CreateBrandingDropResponse UpdateBrandingDrop(ctx, id).Execute()

Edit a drop that is not live yet. A published drop refuses rule problems (400 drop_config_problems).

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
    resp, r, err := apiClient.BrandingApi.UpdateBrandingDrop(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.UpdateBrandingDrop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBrandingDrop`: CreateBrandingDropResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.UpdateBrandingDrop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrandingDropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateBrandingDropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBrandingEventSetting

> UpdateBrandingEventSettingResponse UpdateBrandingEventSetting(ctx, id).Execute()

Save an event's full settings (where, access, reminders, check-in, recording, linked drops) and ticket tiers.

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
    resp, r, err := apiClient.BrandingApi.UpdateBrandingEventSetting(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.UpdateBrandingEventSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBrandingEventSetting`: UpdateBrandingEventSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.UpdateBrandingEventSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrandingEventSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateBrandingEventSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateBrandingSetting

> GetBrandingSettingResponse UpdateBrandingSetting(ctx).Execute()

Change the store’s Branding defaults (any subset of the fields).

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
    resp, r, err := apiClient.BrandingApi.UpdateBrandingSetting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BrandingApi.UpdateBrandingSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBrandingSetting`: GetBrandingSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `BrandingApi.UpdateBrandingSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBrandingSettingRequest struct via the builder pattern


### Return type

**GetBrandingSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

