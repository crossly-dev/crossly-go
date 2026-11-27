# \ConversationsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateConversationMessage**](ConversationsApi.md#CreateConversationMessage) | **Post** /v1/conversations/{id}/messages | Reply to the buyer, on the marketplace the conversation is on (Crossly marketplace: straight into the thread). The message is in the thread at once with deliveryStatus &#39;sending&#39;; it settles to &#39;sent&#39; or &#39;failed&#39;. reply_not_supported when the thread cannot take a reply (a no-reply sender, a platform with no reply wire).
[**ListConversationMessages**](ConversationsApi.md#ListConversationMessages) | **Get** /v1/conversations/{id}/messages | A conversation&#39;s messages, oldest first. sender is buyer | seller; a reply you sent shows deliveryStatus sending → sent | failed. since (ISO 8601): only messages sent at or after it.
[**ListConversations**](ConversationsApi.md#ListConversations) | **Get** /v1/conversations | List buyer conversations, newest activity first. orderId: the conversation(s) about that order (same platform, its listing, its buyer). listingId: about that listing. updatedSince (ISO 8601): only threads with activity at or after it, for polling. search: contains-match on the buyer username or the subject (case-insensitive). Pair with the message.received webhook.



## CreateConversationMessage

> CreateConversationMessageResponse CreateConversationMessage(ctx, id).CreateConversationMessageRequest(createConversationMessageRequest).Execute()

Reply to the buyer, on the marketplace the conversation is on (Crossly marketplace: straight into the thread). The message is in the thread at once with deliveryStatus 'sending'; it settles to 'sent' or 'failed'. reply_not_supported when the thread cannot take a reply (a no-reply sender, a platform with no reply wire).

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
    createConversationMessageRequest := *openapiclient.NewCreateConversationMessageRequest("Body_example") // CreateConversationMessageRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConversationsApi.CreateConversationMessage(context.Background(), id).CreateConversationMessageRequest(createConversationMessageRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConversationsApi.CreateConversationMessage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateConversationMessage`: CreateConversationMessageResponse
    fmt.Fprintf(os.Stdout, "Response from `ConversationsApi.CreateConversationMessage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateConversationMessageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createConversationMessageRequest** | **CreateConversationMessageRequest** |  | 

### Return type

**CreateConversationMessageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListConversationMessages

> ListConversationMessagesResponse ListConversationMessages(ctx, id).Page(page).Limit(limit).Since(since).Execute()

A conversation's messages, oldest first. sender is buyer | seller; a reply you sent shows deliveryStatus sending → sent | failed. since (ISO 8601): only messages sent at or after it.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    page := int32(56) // int32 |  (optional) (default to 1)
    limit := int32(56) // int32 |  (optional) (default to 50)
    since := time.Now() // time.Time |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConversationsApi.ListConversationMessages(context.Background(), id).Page(page).Limit(limit).Since(since).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConversationsApi.ListConversationMessages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConversationMessages`: ListConversationMessagesResponse
    fmt.Fprintf(os.Stdout, "Response from `ConversationsApi.ListConversationMessages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListConversationMessagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 50]
 **since** | **time.Time** |  | 

### Return type

**ListConversationMessagesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListConversations

> ListConversationsResponse ListConversations(ctx).Page(page).Limit(limit).OrderId(orderId).ListingId(listingId).Platform(platform).UpdatedSince(updatedSince).Search(search).Execute()

List buyer conversations, newest activity first. orderId: the conversation(s) about that order (same platform, its listing, its buyer). listingId: about that listing. updatedSince (ISO 8601): only threads with activity at or after it, for polling. search: contains-match on the buyer username or the subject (case-insensitive). Pair with the message.received webhook.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    page := int32(56) // int32 |  (optional) (default to 1)
    limit := int32(56) // int32 |  (optional) (default to 25)
    orderId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    listingId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)
    platform := "platform_example" // string |  (optional)
    updatedSince := time.Now() // time.Time |  (optional)
    search := "search_example" // string | Contains-match (case-insensitive) on the buyer username or the thread subject. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ConversationsApi.ListConversations(context.Background()).Page(page).Limit(limit).OrderId(orderId).ListingId(listingId).Platform(platform).UpdatedSince(updatedSince).Search(search).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ConversationsApi.ListConversations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConversations`: ListConversationsResponse
    fmt.Fprintf(os.Stdout, "Response from `ConversationsApi.ListConversations`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConversationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **page** | **int32** |  | [default to 1]
 **limit** | **int32** |  | [default to 25]
 **orderId** | **string** |  | 
 **listingId** | **string** |  | 
 **platform** | **string** |  | 
 **updatedSince** | **time.Time** |  | 
 **search** | **string** | Contains-match (case-insensitive) on the buyer username or the thread subject. | 

### Return type

**ListConversationsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

