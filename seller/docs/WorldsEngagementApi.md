# \WorldsEngagementApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorldExperienceExperienceEvent**](WorldsEngagementApi.md#CreateWorldExperienceExperienceEvent) | **Post** /v1/worlds/experiences/{experienceId}/experience-events | Schedule an Experience Event shown on the game&#39;s card: title, description, start/end (5 min–14 days, ≤ 180 days ahead), thumbnail, place. With liveEventId it follows that live-ops event&#39;s schedule.
[**CreateWorldExperienceExperienceEventCancel**](WorldsEngagementApi.md#CreateWorldExperienceExperienceEventCancel) | **Post** /v1/worlds/experiences/{experienceId}/experience-events/{eventId}/cancel | Cancel an event (it disappears from the game&#39;s card; no reminder goes out).
[**CreateWorldExperienceExperienceEventRsvp**](WorldsEngagementApi.md#CreateWorldExperienceExperienceEventRsvp) | **Post** /v1/worlds/experiences/{experienceId}/experience-events/{eventId}/rsvp | RSVP (going: true) or take it back. remind: true also turns the experience&#39;s notifications on, so the 15-minute reminder can reach you.
[**CreateWorldExperienceNotification**](WorldsEngagementApi.md#CreateWorldExperienceNotification) | **Post** /v1/worlds/experiences/{experienceId}/notifications | Send an approved template to one player (sub) or up to 100 (subs). Only players who turned notifications on, 13+, not banned, at most 2 a day per player and outside their quiet hours receive it; every other send is logged as suppressed with its reason.
[**CreateWorldExperienceNotificationTemplateRetire**](WorldsEngagementApi.md#CreateWorldExperienceNotificationTemplateRetire) | **Post** /v1/worlds/experiences/{experienceId}/notification-templates/{key}/retire | Retire a template: sends by its key stop.
[**GetWorldExperienceNotificationOptIn**](WorldsEngagementApi.md#GetWorldExperienceNotificationOptIn) | **Get** /v1/worlds/experiences/{experienceId}/notifications/opt-in | Whether you turned this experience&#39;s notifications on.
[**GetWorldExperienceNotificationReport**](WorldsEngagementApi.md#GetWorldExperienceNotificationReport) | **Get** /v1/worlds/experiences/{experienceId}/notifications/report | Sends over the last N days (default 7, max 90): delivered, suppressed by reason, per template, and how many players turned notifications on.
[**ListWorldExperienceEventUpcoming**](WorldsEngagementApi.md#ListWorldExperienceEventUpcoming) | **Get** /v1/worlds/experience-events/upcoming | Events across public experiences in the next 7 days (and live now), soonest first — the Carnival Plaza&#39;s \&quot;Coming up\&quot; board.
[**ListWorldExperienceExperienceEventUpcoming**](WorldsEngagementApi.md#ListWorldExperienceExperienceEventUpcoming) | **Get** /v1/worlds/experiences/{experienceId}/experience-events/upcoming | An experience&#39;s upcoming and live events as players see them, with whether you are going.
[**ListWorldExperienceExperienceEvents**](WorldsEngagementApi.md#ListWorldExperienceExperienceEvents) | **Get** /v1/worlds/experiences/{experienceId}/experience-events | Your scheduled Experience Events (upcoming, live and the last 30 days) with RSVP counts.
[**ListWorldExperienceNotificationTemplates**](WorldsEngagementApi.md#ListWorldExperienceNotificationTemplates) | **Get** /v1/worlds/experiences/{experienceId}/notification-templates | Notification templates: key, title and body with {placeholders}, review status (approved ones can be sent).
[**ListWorldMeExperienceEvents**](WorldsEngagementApi.md#ListWorldMeExperienceEvents) | **Get** /v1/worlds/me/experience-events | Events you are going to that have not ended.
[**ListWorldMeNotificationOptIns**](WorldsEngagementApi.md#ListWorldMeNotificationOptIns) | **Get** /v1/worlds/me/notification-opt-ins | Every experience whose notifications you turned on.
[**UpdateWorldExperienceExperienceEvent**](WorldsEngagementApi.md#UpdateWorldExperienceExperienceEvent) | **Patch** /v1/worlds/experiences/{experienceId}/experience-events/{eventId} | Change an event. A started event keeps its start; a new start time earns RSVPs a new reminder.
[**UpdateWorldExperienceNotificationOptIn**](WorldsEngagementApi.md#UpdateWorldExperienceNotificationOptIn) | **Put** /v1/worlds/experiences/{experienceId}/notifications/opt-in | Turn this experience&#39;s notifications on or off for you. timeZone (IANA) keeps them out of your night.
[**UpdateWorldExperienceNotificationTemplate**](WorldsEngagementApi.md#UpdateWorldExperienceNotificationTemplate) | **Put** /v1/worlds/experiences/{experienceId}/notification-templates/{key} | Create or reword a template. Clean words with no link are approved at once; anything the text filter flags, or any link, waits for Crossly review (24 h). Rewording re-reviews.



## CreateWorldExperienceExperienceEvent

> WorldExperienceExperienceEventResponse CreateWorldExperienceExperienceEvent(ctx, experienceId).Execute()

Schedule an Experience Event shown on the game's card: title, description, start/end (5 min–14 days, ≤ 180 days ahead), thumbnail, place. With liveEventId it follows that live-ops event's schedule.

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
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsEngagementApi.CreateWorldExperienceExperienceEvent(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.CreateWorldExperienceExperienceEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceExperienceEvent`: WorldExperienceExperienceEventResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.CreateWorldExperienceExperienceEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceExperienceEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceExperienceEventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceExperienceEventCancel

> WorldExperienceExperienceEventResponse CreateWorldExperienceExperienceEventCancel(ctx, experienceId, eventId).Execute()

Cancel an event (it disappears from the game's card; no reminder goes out).

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
    experienceId := "experienceId_example" // string | 
    eventId := "eventId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsEngagementApi.CreateWorldExperienceExperienceEventCancel(context.Background(), experienceId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.CreateWorldExperienceExperienceEventCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceExperienceEventCancel`: WorldExperienceExperienceEventResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.CreateWorldExperienceExperienceEventCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceExperienceEventCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceExperienceEventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceExperienceEventRsvp

> WorldExperienceExperienceEventResponse CreateWorldExperienceExperienceEventRsvp(ctx, experienceId, eventId).Execute()

RSVP (going: true) or take it back. remind: true also turns the experience's notifications on, so the 15-minute reminder can reach you.

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
    experienceId := "experienceId_example" // string | 
    eventId := "eventId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsEngagementApi.CreateWorldExperienceExperienceEventRsvp(context.Background(), experienceId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.CreateWorldExperienceExperienceEventRsvp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceExperienceEventRsvp`: WorldExperienceExperienceEventResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.CreateWorldExperienceExperienceEventRsvp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceExperienceEventRsvpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceExperienceEventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceNotification

> CreateWorldExperienceNotificationResponse CreateWorldExperienceNotification(ctx, experienceId).Execute()

Send an approved template to one player (sub) or up to 100 (subs). Only players who turned notifications on, 13+, not banned, at most 2 a day per player and outside their quiet hours receive it; every other send is logged as suppressed with its reason.

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
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsEngagementApi.CreateWorldExperienceNotification(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.CreateWorldExperienceNotification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceNotification`: CreateWorldExperienceNotificationResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.CreateWorldExperienceNotification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceNotificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceNotificationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceNotificationTemplateRetire

> UpdateWorldExperienceNotificationTemplateResponse CreateWorldExperienceNotificationTemplateRetire(ctx, experienceId, key).Execute()

Retire a template: sends by its key stop.

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
    experienceId := "experienceId_example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsEngagementApi.CreateWorldExperienceNotificationTemplateRetire(context.Background(), experienceId, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.CreateWorldExperienceNotificationTemplateRetire``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceNotificationTemplateRetire`: UpdateWorldExperienceNotificationTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.CreateWorldExperienceNotificationTemplateRetire`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceNotificationTemplateRetireRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**UpdateWorldExperienceNotificationTemplateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceNotificationOptIn

> GetWorldExperienceNotificationOptInResponse GetWorldExperienceNotificationOptIn(ctx, experienceId).Execute()

Whether you turned this experience's notifications on.

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
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsEngagementApi.GetWorldExperienceNotificationOptIn(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.GetWorldExperienceNotificationOptIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceNotificationOptIn`: GetWorldExperienceNotificationOptInResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.GetWorldExperienceNotificationOptIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceNotificationOptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceNotificationOptInResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceNotificationReport

> GetWorldExperienceNotificationReportResponse GetWorldExperienceNotificationReport(ctx, experienceId).Execute()

Sends over the last N days (default 7, max 90): delivered, suppressed by reason, per template, and how many players turned notifications on.

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
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsEngagementApi.GetWorldExperienceNotificationReport(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.GetWorldExperienceNotificationReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceNotificationReport`: GetWorldExperienceNotificationReportResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.GetWorldExperienceNotificationReport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceNotificationReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceNotificationReportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceEventUpcoming

> ListWorldExperienceEventUpcomingResponse ListWorldExperienceEventUpcoming(ctx).Execute()

Events across public experiences in the next 7 days (and live now), soonest first — the Carnival Plaza's \"Coming up\" board.

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
    resp, r, err := apiClient.WorldsEngagementApi.ListWorldExperienceEventUpcoming(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.ListWorldExperienceEventUpcoming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceEventUpcoming`: ListWorldExperienceEventUpcomingResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.ListWorldExperienceEventUpcoming`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceEventUpcomingRequest struct via the builder pattern


### Return type

**ListWorldExperienceEventUpcomingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceExperienceEventUpcoming

> ListWorldExperienceExperienceEventsResponse ListWorldExperienceExperienceEventUpcoming(ctx, experienceId).Execute()

An experience's upcoming and live events as players see them, with whether you are going.

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
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsEngagementApi.ListWorldExperienceExperienceEventUpcoming(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.ListWorldExperienceExperienceEventUpcoming``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceExperienceEventUpcoming`: ListWorldExperienceExperienceEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.ListWorldExperienceExperienceEventUpcoming`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceExperienceEventUpcomingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceExperienceEventsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceExperienceEvents

> ListWorldExperienceExperienceEventsResponse ListWorldExperienceExperienceEvents(ctx, experienceId).Execute()

Your scheduled Experience Events (upcoming, live and the last 30 days) with RSVP counts.

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
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsEngagementApi.ListWorldExperienceExperienceEvents(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.ListWorldExperienceExperienceEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceExperienceEvents`: ListWorldExperienceExperienceEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.ListWorldExperienceExperienceEvents`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceExperienceEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceExperienceEventsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceNotificationTemplates

> ListWorldExperienceNotificationTemplatesResponse ListWorldExperienceNotificationTemplates(ctx, experienceId).Execute()

Notification templates: key, title and body with {placeholders}, review status (approved ones can be sent).

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
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsEngagementApi.ListWorldExperienceNotificationTemplates(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.ListWorldExperienceNotificationTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceNotificationTemplates`: ListWorldExperienceNotificationTemplatesResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.ListWorldExperienceNotificationTemplates`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceNotificationTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceNotificationTemplatesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldMeExperienceEvents

> ListWorldMeExperienceEventsResponse ListWorldMeExperienceEvents(ctx).Execute()

Events you are going to that have not ended.

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
    resp, r, err := apiClient.WorldsEngagementApi.ListWorldMeExperienceEvents(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.ListWorldMeExperienceEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldMeExperienceEvents`: ListWorldMeExperienceEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.ListWorldMeExperienceEvents`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldMeExperienceEventsRequest struct via the builder pattern


### Return type

**ListWorldMeExperienceEventsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldMeNotificationOptIns

> ListWorldMeNotificationOptInsResponse ListWorldMeNotificationOptIns(ctx).Execute()

Every experience whose notifications you turned on.

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
    resp, r, err := apiClient.WorldsEngagementApi.ListWorldMeNotificationOptIns(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.ListWorldMeNotificationOptIns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldMeNotificationOptIns`: ListWorldMeNotificationOptInsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.ListWorldMeNotificationOptIns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldMeNotificationOptInsRequest struct via the builder pattern


### Return type

**ListWorldMeNotificationOptInsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceExperienceEvent

> WorldExperienceExperienceEventResponse UpdateWorldExperienceExperienceEvent(ctx, experienceId, eventId).Execute()

Change an event. A started event keeps its start; a new start time earns RSVPs a new reminder.

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
    experienceId := "experienceId_example" // string | 
    eventId := "eventId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsEngagementApi.UpdateWorldExperienceExperienceEvent(context.Background(), experienceId, eventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.UpdateWorldExperienceExperienceEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceExperienceEvent`: WorldExperienceExperienceEventResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.UpdateWorldExperienceExperienceEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**eventId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceExperienceEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceExperienceEventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceNotificationOptIn

> GetWorldExperienceNotificationOptInResponse UpdateWorldExperienceNotificationOptIn(ctx, experienceId).Execute()

Turn this experience's notifications on or off for you. timeZone (IANA) keeps them out of your night.

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
    experienceId := "experienceId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsEngagementApi.UpdateWorldExperienceNotificationOptIn(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.UpdateWorldExperienceNotificationOptIn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceNotificationOptIn`: GetWorldExperienceNotificationOptInResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.UpdateWorldExperienceNotificationOptIn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceNotificationOptInRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceNotificationOptInResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceNotificationTemplate

> UpdateWorldExperienceNotificationTemplateResponse UpdateWorldExperienceNotificationTemplate(ctx, experienceId, key).Execute()

Create or reword a template. Clean words with no link are approved at once; anything the text filter flags, or any link, waits for Crossly review (24 h). Rewording re-reviews.

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
    experienceId := "experienceId_example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsEngagementApi.UpdateWorldExperienceNotificationTemplate(context.Background(), experienceId, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsEngagementApi.UpdateWorldExperienceNotificationTemplate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceNotificationTemplate`: UpdateWorldExperienceNotificationTemplateResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsEngagementApi.UpdateWorldExperienceNotificationTemplate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceNotificationTemplateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**UpdateWorldExperienceNotificationTemplateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

