# \InboxApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInboxAiSuggest**](InboxApi.md#CreateInboxAiSuggest) | **Post** /v1/inbox/ai-suggest | AI reply suggestion for a conversation.
[**CreateInboxCannedResponse**](InboxApi.md#CreateInboxCannedResponse) | **Post** /v1/inbox/canned-responses | Create a canned response (added to the end of the list).
[**CreateInboxConversationBulk**](InboxApi.md#CreateInboxConversationBulk) | **Post** /v1/inbox/conversations/bulk | Bulk mark read / mark unread / soft-delete conversations.
[**CreateInboxConversationBulkAiRespond**](InboxApi.md#CreateInboxConversationBulkAiRespond) | **Post** /v1/inbox/conversations/bulk-ai-respond | AI reply suggestion for multiple conversations — draft or send.
[**CreateInboxConversationOfferAction**](InboxApi.md#CreateInboxConversationOfferAction) | **Post** /v1/inbox/conversations/{id}/offer-action | Accept / counter / decline an active offer on a conversation.
[**CreateInboxMessageTriage**](InboxApi.md#CreateInboxMessageTriage) | **Post** /v1/inbox/messages/{id}/triage | Manually re-triage a buyer message.
[**CreateInboxOffer**](InboxApi.md#CreateInboxOffer) | **Post** /v1/inbox/{id}/offer | Accept, counter, or decline an offer on a conversation.
[**CreateInboxReply**](InboxApi.md#CreateInboxReply) | **Post** /v1/inbox/{id}/reply | Send a reply to a conversation thread.
[**CreateInboxSendOffer**](InboxApi.md#CreateInboxSendOffer) | **Post** /v1/inbox/{id}/send-offer | Make the buyer an offer inside a conversation (seller → buyer). Only where the conversation&#39;s canSendOffer is true (Grailed today). &#x60;amount&#x60; is in the listing&#39;s currency, above 0 and below the listing price (whole dollars on Grailed). The offer is in the thread at once as a message with offerAmount + offerStatus, deliveryStatus &#39;sending&#39; until the platform confirms.
[**CreateVideoRequestAnswer**](InboxApi.md#CreateVideoRequestAnswer) | **Post** /v1/video-requests/{id}/answer | Answer a video request with a clip you uploaded (POST /v1/video-requests/upload-url, then PUT the file to uploadUrl). The buyer is notified and the clip shows on the listing.
[**CreateVideoRequestDecline**](InboxApi.md#CreateVideoRequestDecline) | **Post** /v1/video-requests/{id}/decline | Decline a video request, with an optional note to the buyer.
[**CreateVideoRequestUploadUrl**](InboxApi.md#CreateVideoRequestUploadUrl) | **Post** /v1/video-requests/upload-url | A presigned upload for a video answer: PUT the file to uploadUrl with the same Content-Type and exact Content-Length, then answer with publicUrl.
[**DeleteInboxCannedResponse**](InboxApi.md#DeleteInboxCannedResponse) | **Delete** /v1/inbox/canned-responses/{id} | Delete a canned response.
[**GetInbox**](InboxApi.md#GetInbox) | **Get** /v1/inbox/{id} | Get one conversation with its messages and the item it is about (itemTitle, itemThumbnail, itemPrice, itemCurrency, itemUrl on the platform, listingId in Crossly, crosslyListingSlug, canSendOffer). An offer made in the thread is a message with offerAmount + offerStatus.
[**GetInboxCannedResponse**](InboxApi.md#GetInboxCannedResponse) | **Get** /v1/inbox/canned-responses | List canned responses in the seller&#39;s order.
[**GetInboxConversationMessage**](InboxApi.md#GetInboxConversationMessage) | **Get** /v1/inbox/conversations/{id}/messages | Paginated messages for a conversation.
[**GetInboxConversationUnreadCount**](InboxApi.md#GetInboxConversationUnreadCount) | **Get** /v1/inbox/conversations/unread-count | Unread conversation count (threads only; /v1/inbox/counts adds offers).
[**GetInboxCount**](InboxApi.md#GetInboxCount) | **Get** /v1/inbox/counts | Inbox badges: unread and needs-reply, split into conversations and offers.
[**ListInbox**](InboxApi.md#ListInbox) | **Get** /v1/inbox | List conversations. &#x60;filter&#x60; narrows to the rows GET /v1/inbox/counts counts.
[**ListVideoRequests**](InboxApi.md#ListVideoRequests) | **Get** /v1/video-requests | Buyers asking for a video of your items. status: open (default), answered, declined or all.
[**UpdateInboxCannedResponse**](InboxApi.md#UpdateInboxCannedResponse) | **Put** /v1/inbox/canned-responses/{id} | Update a canned response.
[**UpdateInboxCannedResponseOrder**](InboxApi.md#UpdateInboxCannedResponseOrder) | **Put** /v1/inbox/canned-responses/order | Reorder canned responses.
[**UpdateInboxConversation**](InboxApi.md#UpdateInboxConversation) | **Patch** /v1/inbox/conversations/{id} | Mark read / change status / close conversation.



## CreateInboxAiSuggest

> CreateInboxAiSuggestResponse CreateInboxAiSuggest(ctx).Execute()

AI reply suggestion for a conversation.

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
    resp, r, err := apiClient.InboxApi.CreateInboxAiSuggest(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.CreateInboxAiSuggest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInboxAiSuggest`: CreateInboxAiSuggestResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.CreateInboxAiSuggest`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInboxAiSuggestRequest struct via the builder pattern


### Return type

**CreateInboxAiSuggestResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateInboxCannedResponse

> CreateInboxCannedResponseResponse CreateInboxCannedResponse(ctx).Execute()

Create a canned response (added to the end of the list).

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
    resp, r, err := apiClient.InboxApi.CreateInboxCannedResponse(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.CreateInboxCannedResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInboxCannedResponse`: CreateInboxCannedResponseResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.CreateInboxCannedResponse`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInboxCannedResponseRequest struct via the builder pattern


### Return type

**CreateInboxCannedResponseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateInboxConversationBulk

> CreateInboxConversationBulkResponse CreateInboxConversationBulk(ctx).Execute()

Bulk mark read / mark unread / soft-delete conversations.

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
    resp, r, err := apiClient.InboxApi.CreateInboxConversationBulk(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.CreateInboxConversationBulk``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInboxConversationBulk`: CreateInboxConversationBulkResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.CreateInboxConversationBulk`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInboxConversationBulkRequest struct via the builder pattern


### Return type

**CreateInboxConversationBulkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateInboxConversationBulkAiRespond

> CreateInboxConversationBulkAiRespondResponse CreateInboxConversationBulkAiRespond(ctx).Execute()

AI reply suggestion for multiple conversations — draft or send.

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
    resp, r, err := apiClient.InboxApi.CreateInboxConversationBulkAiRespond(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.CreateInboxConversationBulkAiRespond``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInboxConversationBulkAiRespond`: CreateInboxConversationBulkAiRespondResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.CreateInboxConversationBulkAiRespond`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInboxConversationBulkAiRespondRequest struct via the builder pattern


### Return type

**CreateInboxConversationBulkAiRespondResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateInboxConversationOfferAction

> UpdateInboxConversationResponse CreateInboxConversationOfferAction(ctx, id).Execute()

Accept / counter / decline an active offer on a conversation.

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
    resp, r, err := apiClient.InboxApi.CreateInboxConversationOfferAction(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.CreateInboxConversationOfferAction``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInboxConversationOfferAction`: UpdateInboxConversationResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.CreateInboxConversationOfferAction`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInboxConversationOfferActionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateInboxConversationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateInboxMessageTriage

> CreateInboxMessageTriageResponse CreateInboxMessageTriage(ctx, id).Execute()

Manually re-triage a buyer message.

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
    resp, r, err := apiClient.InboxApi.CreateInboxMessageTriage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.CreateInboxMessageTriage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInboxMessageTriage`: CreateInboxMessageTriageResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.CreateInboxMessageTriage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInboxMessageTriageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateInboxMessageTriageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateInboxOffer

> CreateInboxOfferResponse CreateInboxOffer(ctx, id).Execute()

Accept, counter, or decline an offer on a conversation.

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
    resp, r, err := apiClient.InboxApi.CreateInboxOffer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.CreateInboxOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInboxOffer`: CreateInboxOfferResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.CreateInboxOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInboxOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateInboxOfferResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateInboxReply

> CreateInboxReplyResponse CreateInboxReply(ctx, id).Execute()

Send a reply to a conversation thread.

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
    resp, r, err := apiClient.InboxApi.CreateInboxReply(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.CreateInboxReply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInboxReply`: CreateInboxReplyResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.CreateInboxReply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInboxReplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateInboxReplyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateInboxSendOffer

> CreateInboxSendOfferResponse CreateInboxSendOffer(ctx, id).CreateInboxSendOfferRequest(createInboxSendOfferRequest).Execute()

Make the buyer an offer inside a conversation (seller → buyer). Only where the conversation's canSendOffer is true (Grailed today). `amount` is in the listing's currency, above 0 and below the listing price (whole dollars on Grailed). The offer is in the thread at once as a message with offerAmount + offerStatus, deliveryStatus 'sending' until the platform confirms.

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
    createInboxSendOfferRequest := *openapiclient.NewCreateInboxSendOfferRequest(float32(123)) // CreateInboxSendOfferRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InboxApi.CreateInboxSendOffer(context.Background(), id).CreateInboxSendOfferRequest(createInboxSendOfferRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.CreateInboxSendOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInboxSendOffer`: CreateInboxSendOfferResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.CreateInboxSendOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInboxSendOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInboxSendOfferRequest** | **CreateInboxSendOfferRequest** |  | 

### Return type

**CreateInboxSendOfferResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVideoRequestAnswer

> CreateVideoRequestAnswerResponse CreateVideoRequestAnswer(ctx, id).Execute()

Answer a video request with a clip you uploaded (POST /v1/video-requests/upload-url, then PUT the file to uploadUrl). The buyer is notified and the clip shows on the listing.

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
    resp, r, err := apiClient.InboxApi.CreateVideoRequestAnswer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.CreateVideoRequestAnswer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVideoRequestAnswer`: CreateVideoRequestAnswerResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.CreateVideoRequestAnswer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVideoRequestAnswerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVideoRequestAnswerResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVideoRequestDecline

> CreateVideoRequestAnswerResponse CreateVideoRequestDecline(ctx, id).Execute()

Decline a video request, with an optional note to the buyer.

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
    resp, r, err := apiClient.InboxApi.CreateVideoRequestDecline(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.CreateVideoRequestDecline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVideoRequestDecline`: CreateVideoRequestAnswerResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.CreateVideoRequestDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVideoRequestDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateVideoRequestAnswerResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateVideoRequestUploadUrl

> CreateVideoRequestUploadUrlResponse CreateVideoRequestUploadUrl(ctx).Execute()

A presigned upload for a video answer: PUT the file to uploadUrl with the same Content-Type and exact Content-Length, then answer with publicUrl.

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
    resp, r, err := apiClient.InboxApi.CreateVideoRequestUploadUrl(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.CreateVideoRequestUploadUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateVideoRequestUploadUrl`: CreateVideoRequestUploadUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.CreateVideoRequestUploadUrl`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateVideoRequestUploadUrlRequest struct via the builder pattern


### Return type

**CreateVideoRequestUploadUrlResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteInboxCannedResponse

> DeleteInboxCannedResponse(ctx, id).Execute()

Delete a canned response.

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
    resp, r, err := apiClient.InboxApi.DeleteInboxCannedResponse(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.DeleteInboxCannedResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInboxCannedResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetInbox

> GetInboxResponse GetInbox(ctx, id).Execute()

Get one conversation with its messages and the item it is about (itemTitle, itemThumbnail, itemPrice, itemCurrency, itemUrl on the platform, listingId in Crossly, crosslyListingSlug, canSendOffer). An offer made in the thread is a message with offerAmount + offerStatus.

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
    resp, r, err := apiClient.InboxApi.GetInbox(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.GetInbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInbox`: GetInboxResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.GetInbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetInboxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetInboxCannedResponse

> GetInboxCannedResponseResponse GetInboxCannedResponse(ctx).Execute()

List canned responses in the seller's order.

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
    resp, r, err := apiClient.InboxApi.GetInboxCannedResponse(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.GetInboxCannedResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInboxCannedResponse`: GetInboxCannedResponseResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.GetInboxCannedResponse`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxCannedResponseRequest struct via the builder pattern


### Return type

**GetInboxCannedResponseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetInboxConversationMessage

> GetInboxConversationMessageResponse GetInboxConversationMessage(ctx, id).Execute()

Paginated messages for a conversation.

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
    resp, r, err := apiClient.InboxApi.GetInboxConversationMessage(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.GetInboxConversationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInboxConversationMessage`: GetInboxConversationMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.GetInboxConversationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxConversationMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetInboxConversationMessageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetInboxConversationUnreadCount

> CreateWeb3LaunchArtResponse GetInboxConversationUnreadCount(ctx).Execute()

Unread conversation count (threads only; /v1/inbox/counts adds offers).

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
    resp, r, err := apiClient.InboxApi.GetInboxConversationUnreadCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.GetInboxConversationUnreadCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInboxConversationUnreadCount`: CreateWeb3LaunchArtResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.GetInboxConversationUnreadCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxConversationUnreadCountRequest struct via the builder pattern


### Return type

**CreateWeb3LaunchArtResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetInboxCount

> GetInboxCountResponse GetInboxCount(ctx).Execute()

Inbox badges: unread and needs-reply, split into conversations and offers.

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
    resp, r, err := apiClient.InboxApi.GetInboxCount(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.GetInboxCount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInboxCount`: GetInboxCountResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.GetInboxCount`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetInboxCountRequest struct via the builder pattern


### Return type

**GetInboxCountResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListInbox

> ListInboxResponse ListInbox(ctx).Page(page).Limit(limit).Filter(filter).Execute()

List conversations. `filter` narrows to the rows GET /v1/inbox/counts counts.

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
    page := int32(56) // int32 |  (optional) (default to 1)
    limit := int32(56) // int32 |  (optional) (default to 25)
    filter := "filter_example" // string | unread | needs_reply (services/inbox/attention.ts). Omit for all. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InboxApi.ListInbox(context.Background()).Page(page).Limit(limit).Filter(filter).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.ListInbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInbox`: ListInboxResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.ListInbox`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListInboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 25]
 **filter** | **string** | unread | needs_reply (services/inbox/attention.ts). Omit for all. | 

### Return type

**ListInboxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListVideoRequests

> ListVideoRequestsResponse ListVideoRequests(ctx).Execute()

Buyers asking for a video of your items. status: open (default), answered, declined or all.

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
    resp, r, err := apiClient.InboxApi.ListVideoRequests(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.ListVideoRequests``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListVideoRequests`: ListVideoRequestsResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.ListVideoRequests`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListVideoRequestsRequest struct via the builder pattern


### Return type

**ListVideoRequestsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateInboxCannedResponse

> CreateInboxCannedResponseResponse UpdateInboxCannedResponse(ctx, id).Execute()

Update a canned response.

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
    resp, r, err := apiClient.InboxApi.UpdateInboxCannedResponse(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.UpdateInboxCannedResponse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInboxCannedResponse`: CreateInboxCannedResponseResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.UpdateInboxCannedResponse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInboxCannedResponseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateInboxCannedResponseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateInboxCannedResponseOrder

> GetInboxCannedResponseResponse UpdateInboxCannedResponseOrder(ctx).UpdateInboxCannedResponseOrderRequest(updateInboxCannedResponseOrderRequest).Execute()

Reorder canned responses.



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
    updateInboxCannedResponseOrderRequest := *openapiclient.NewUpdateInboxCannedResponseOrderRequest([]string{"Ids_example"}) // UpdateInboxCannedResponseOrderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.InboxApi.UpdateInboxCannedResponseOrder(context.Background()).UpdateInboxCannedResponseOrderRequest(updateInboxCannedResponseOrderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.UpdateInboxCannedResponseOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInboxCannedResponseOrder`: GetInboxCannedResponseResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.UpdateInboxCannedResponseOrder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInboxCannedResponseOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateInboxCannedResponseOrderRequest** | **UpdateInboxCannedResponseOrderRequest** |  | 

### Return type

**GetInboxCannedResponseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateInboxConversation

> UpdateInboxConversationResponse UpdateInboxConversation(ctx, id).Execute()

Mark read / change status / close conversation.

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
    resp, r, err := apiClient.InboxApi.UpdateInboxConversation(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `InboxApi.UpdateInboxConversation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInboxConversation`: UpdateInboxConversationResponse
    fmt.Fprintf(os.Stdout, "Response from `InboxApi.UpdateInboxConversation`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInboxConversationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateInboxConversationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

