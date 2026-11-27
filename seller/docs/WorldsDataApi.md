# \WorldsDataApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorldExperienceDatastoreEntryIncrement**](WorldsDataApi.md#CreateWorldExperienceDatastoreEntryIncrement) | **Post** /v1/worlds/experiences/{experienceId}/datastores/{store}/entries/{key}/increment | Atomically add &#x60;by&#x60; to a numeric key (absent counts as 0).
[**CreateWorldExperienceMemoryQueueItem**](WorldsDataApi.md#CreateWorldExperienceMemoryQueueItem) | **Post** /v1/worlds/experiences/{experienceId}/memory/queues/{queue}/items | Add an item to a MemoryStore queue (ttlSeconds ≤ 45 days, priority 0-1000).
[**CreateWorldExperienceMemoryQueueRead**](WorldsDataApi.md#CreateWorldExperienceMemoryQueueRead) | **Post** /v1/worlds/experiences/{experienceId}/memory/queues/{queue}/read | Read and hide up to &#x60;count&#x60; queue items for invisibilitySeconds; remove them with the readId.
[**CreateWorldExperienceMessaging**](WorldsDataApi.md#CreateWorldExperienceMessaging) | **Post** /v1/worlds/experiences/{experienceId}/messaging/{topic} | Publish a message (≤ 1 KB) to every live server of the experience subscribed to the topic.
[**CreateWorldExperienceOrderedEntryIncrement**](WorldsDataApi.md#CreateWorldExperienceOrderedEntryIncrement) | **Post** /v1/worlds/experiences/{experienceId}/ordered/{store}/entries/{key}/increment | Atomically add &#x60;by&#x60; to an OrderedStore entry.
[**DeleteWorldExperienceDatastoreEntry**](WorldsDataApi.md#DeleteWorldExperienceDatastoreEntry) | **Delete** /v1/worlds/experiences/{experienceId}/datastores/{store}/entries/{key} | Remove a key (a tombstone version; history stays readable for 30 days). Honours If-Match.
[**DeleteWorldExperienceMemoryMap**](WorldsDataApi.md#DeleteWorldExperienceMemoryMap) | **Delete** /v1/worlds/experiences/{experienceId}/memory/maps/{map}/{key} | Remove a hash map item.
[**DeleteWorldExperienceMemoryQueueRead**](WorldsDataApi.md#DeleteWorldExperienceMemoryQueueRead) | **Delete** /v1/worlds/experiences/{experienceId}/memory/queues/{queue}/reads/{readId} | Remove the items of one read (409 if the read lapsed).
[**DeleteWorldExperienceMemorySorted**](WorldsDataApi.md#DeleteWorldExperienceMemorySorted) | **Delete** /v1/worlds/experiences/{experienceId}/memory/sorted/{map}/{key} | Remove a sorted map item.
[**DeleteWorldExperienceOrderedEntry**](WorldsDataApi.md#DeleteWorldExperienceOrderedEntry) | **Delete** /v1/worlds/experiences/{experienceId}/ordered/{store}/entries/{key} | Remove an OrderedStore entry.
[**GetWorldExperienceDataUsage**](WorldsDataApi.md#GetWorldExperienceDataUsage) | **Get** /v1/worlds/experiences/{experienceId}/data/usage | Live quotas and usage: CCU, request budgets per minute, DataStore storage, MemoryStore memory.
[**GetWorldExperienceDatastoreEntry**](WorldsDataApi.md#GetWorldExperienceDatastoreEntry) | **Get** /v1/worlds/experiences/{experienceId}/datastores/{store}/entries/{key} | The current value of a key: version (ETag \&quot;v&lt;n&gt;\&quot;), userIds and attributes.
[**GetWorldExperienceDatastoreEntryVersion**](WorldsDataApi.md#GetWorldExperienceDatastoreEntryVersion) | **Get** /v1/worlds/experiences/{experienceId}/datastores/{store}/entries/{key}/versions/{version} | One version of a key, including a removed one.
[**GetWorldExperienceMemoryMap**](WorldsDataApi.md#GetWorldExperienceMemoryMap) | **Get** /v1/worlds/experiences/{experienceId}/memory/maps/{map} | Items of a MemoryStore hash map (cursor-paged, unordered).
[**GetWorldExperienceMemoryMapByKey**](WorldsDataApi.md#GetWorldExperienceMemoryMapByKey) | **Get** /v1/worlds/experiences/{experienceId}/memory/maps/{map}/{key} | One MemoryStore hash map item with its version and expiry.
[**GetWorldExperienceMemoryQueue**](WorldsDataApi.md#GetWorldExperienceMemoryQueue) | **Get** /v1/worlds/experiences/{experienceId}/memory/queues/{queue} | Queue length (visible, in flight) and the head, without hiding anything.
[**GetWorldExperienceMemorySorted**](WorldsDataApi.md#GetWorldExperienceMemorySorted) | **Get** /v1/worlds/experiences/{experienceId}/memory/sorted/{map} | A MemoryStore sorted map in order (sortKey, then key); ?order&#x3D;&amp;cursor&#x3D;&amp;limit&#x3D;.
[**GetWorldExperienceMemorySortedByKey**](WorldsDataApi.md#GetWorldExperienceMemorySortedByKey) | **Get** /v1/worlds/experiences/{experienceId}/memory/sorted/{map}/{key} | One sorted map item.
[**GetWorldExperienceOrdered**](WorldsDataApi.md#GetWorldExperienceOrdered) | **Get** /v1/worlds/experiences/{experienceId}/ordered/{store} | An OrderedStore page ranked by value; ?order&#x3D;desc|asc&amp;min&#x3D;&amp;max&#x3D;&amp;cursor&#x3D;&amp;limit&#x3D;.
[**GetWorldExperienceOrderedEntry**](WorldsDataApi.md#GetWorldExperienceOrderedEntry) | **Get** /v1/worlds/experiences/{experienceId}/ordered/{store}/entries/{key} | One OrderedStore entry.
[**ListWorldExperienceDatastoreEntries**](WorldsDataApi.md#ListWorldExperienceDatastoreEntries) | **Get** /v1/worlds/experiences/{experienceId}/datastores/{store}/entries | Keys of a DataStore in key order; ?prefix&#x3D;&amp;scope&#x3D;&amp;cursor&#x3D;&amp;limit&#x3D;.
[**ListWorldExperienceDatastoreEntryVersions**](WorldsDataApi.md#ListWorldExperienceDatastoreEntryVersions) | **Get** /v1/worlds/experiences/{experienceId}/datastores/{store}/entries/{key}/versions | Versions of a key, newest first (30-day history; the current version is kept).
[**ListWorldExperienceDatastores**](WorldsDataApi.md#ListWorldExperienceDatastores) | **Get** /v1/worlds/experiences/{experienceId}/datastores | The DataStores of an experience, with key counts and bytes.
[**UpdateWorldExperienceDatastoreEntry**](WorldsDataApi.md#UpdateWorldExperienceDatastoreEntry) | **Put** /v1/worlds/experiences/{experienceId}/datastores/{store}/entries/{key} | Write a new version. If-Match \&quot;v&lt;n&gt;\&quot; (or ifVersion) writes only at that version; If-None-Match: * only if absent. 412 on a mismatch.
[**UpdateWorldExperienceMemoryMap**](WorldsDataApi.md#UpdateWorldExperienceMemoryMap) | **Put** /v1/worlds/experiences/{experienceId}/memory/maps/{map}/{key} | Set a hash map item (ttlSeconds required; ifVersion for compare-and-set).
[**UpdateWorldExperienceMemorySorted**](WorldsDataApi.md#UpdateWorldExperienceMemorySorted) | **Put** /v1/worlds/experiences/{experienceId}/memory/sorted/{map}/{key} | Set a sorted map item (ttlSeconds, optional sortKey and ifVersion).
[**UpdateWorldExperienceOrderedEntry**](WorldsDataApi.md#UpdateWorldExperienceOrderedEntry) | **Put** /v1/worlds/experiences/{experienceId}/ordered/{store}/entries/{key} | Set an OrderedStore entry to a whole number.



## CreateWorldExperienceDatastoreEntryIncrement

> CreateWorldExperienceDatastoreEntryIncrementResponse CreateWorldExperienceDatastoreEntryIncrement(ctx, experienceId, store, key).Execute()

Atomically add `by` to a numeric key (absent counts as 0).

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
    store := "store_example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.CreateWorldExperienceDatastoreEntryIncrement(context.Background(), experienceId, store, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.CreateWorldExperienceDatastoreEntryIncrement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceDatastoreEntryIncrement`: CreateWorldExperienceDatastoreEntryIncrementResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.CreateWorldExperienceDatastoreEntryIncrement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**store** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceDatastoreEntryIncrementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**CreateWorldExperienceDatastoreEntryIncrementResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceMemoryQueueItem

> CreateWorldPartyInviteResponse CreateWorldExperienceMemoryQueueItem(ctx, experienceId, queue).Execute()

Add an item to a MemoryStore queue (ttlSeconds ≤ 45 days, priority 0-1000).

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
    queue := "queue_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.CreateWorldExperienceMemoryQueueItem(context.Background(), experienceId, queue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.CreateWorldExperienceMemoryQueueItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceMemoryQueueItem`: CreateWorldPartyInviteResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.CreateWorldExperienceMemoryQueueItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**queue** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceMemoryQueueItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldPartyInviteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceMemoryQueueRead

> CreateWorldExperienceMemoryQueueReadResponse CreateWorldExperienceMemoryQueueRead(ctx, experienceId, queue).Execute()

Read and hide up to `count` queue items for invisibilitySeconds; remove them with the readId.

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
    queue := "queue_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.CreateWorldExperienceMemoryQueueRead(context.Background(), experienceId, queue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.CreateWorldExperienceMemoryQueueRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceMemoryQueueRead`: CreateWorldExperienceMemoryQueueReadResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.CreateWorldExperienceMemoryQueueRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**queue** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceMemoryQueueReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperienceMemoryQueueReadResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceMessaging

> CreateWorldExperienceMessagingResponse CreateWorldExperienceMessaging(ctx, experienceId, topic).Execute()

Publish a message (≤ 1 KB) to every live server of the experience subscribed to the topic.

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
    topic := "topic_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.CreateWorldExperienceMessaging(context.Background(), experienceId, topic).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.CreateWorldExperienceMessaging``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceMessaging`: CreateWorldExperienceMessagingResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.CreateWorldExperienceMessaging`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**topic** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceMessagingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperienceMessagingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceOrderedEntryIncrement

> WorldExperienceOrderedEntryResponse CreateWorldExperienceOrderedEntryIncrement(ctx, experienceId, store, key).Execute()

Atomically add `by` to an OrderedStore entry.

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
    store := "store_example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.CreateWorldExperienceOrderedEntryIncrement(context.Background(), experienceId, store, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.CreateWorldExperienceOrderedEntryIncrement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceOrderedEntryIncrement`: WorldExperienceOrderedEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.CreateWorldExperienceOrderedEntryIncrement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**store** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceOrderedEntryIncrementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**WorldExperienceOrderedEntryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorldExperienceDatastoreEntry

> DeleteWorldExperienceDatastoreEntryResponse DeleteWorldExperienceDatastoreEntry(ctx, experienceId, store, key).Execute()

Remove a key (a tombstone version; history stays readable for 30 days). Honours If-Match.

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
    store := "store_example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.DeleteWorldExperienceDatastoreEntry(context.Background(), experienceId, store, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.DeleteWorldExperienceDatastoreEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldExperienceDatastoreEntry`: DeleteWorldExperienceDatastoreEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.DeleteWorldExperienceDatastoreEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**store** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldExperienceDatastoreEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**DeleteWorldExperienceDatastoreEntryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorldExperienceMemoryMap

> DeleteSpatialSceneGrantResponse DeleteWorldExperienceMemoryMap(ctx, experienceId, map_, key).Execute()

Remove a hash map item.

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
    map_ := "map__example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.DeleteWorldExperienceMemoryMap(context.Background(), experienceId, map_, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.DeleteWorldExperienceMemoryMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldExperienceMemoryMap`: DeleteSpatialSceneGrantResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.DeleteWorldExperienceMemoryMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**map_** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldExperienceMemoryMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**DeleteSpatialSceneGrantResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorldExperienceMemoryQueueRead

> DeleteSpatialSceneNodeResponse DeleteWorldExperienceMemoryQueueRead(ctx, experienceId, queue, readId).Execute()

Remove the items of one read (409 if the read lapsed).

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
    queue := "queue_example" // string | 
    readId := "readId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.DeleteWorldExperienceMemoryQueueRead(context.Background(), experienceId, queue, readId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.DeleteWorldExperienceMemoryQueueRead``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldExperienceMemoryQueueRead`: DeleteSpatialSceneNodeResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.DeleteWorldExperienceMemoryQueueRead`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**queue** | **string** |  | 
**readId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldExperienceMemoryQueueReadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**DeleteSpatialSceneNodeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorldExperienceMemorySorted

> DeleteSpatialSceneGrantResponse DeleteWorldExperienceMemorySorted(ctx, experienceId, map_, key).Execute()

Remove a sorted map item.

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
    map_ := "map__example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.DeleteWorldExperienceMemorySorted(context.Background(), experienceId, map_, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.DeleteWorldExperienceMemorySorted``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldExperienceMemorySorted`: DeleteSpatialSceneGrantResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.DeleteWorldExperienceMemorySorted`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**map_** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldExperienceMemorySortedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**DeleteSpatialSceneGrantResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorldExperienceOrderedEntry

> DeleteWorldExperienceOrderedEntryResponse DeleteWorldExperienceOrderedEntry(ctx, experienceId, store, key).Execute()

Remove an OrderedStore entry.

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
    store := "store_example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.DeleteWorldExperienceOrderedEntry(context.Background(), experienceId, store, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.DeleteWorldExperienceOrderedEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldExperienceOrderedEntry`: DeleteWorldExperienceOrderedEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.DeleteWorldExperienceOrderedEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**store** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldExperienceOrderedEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**DeleteWorldExperienceOrderedEntryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceDataUsage

> GetWorldExperienceDataUsageResponse GetWorldExperienceDataUsage(ctx, experienceId).Execute()

Live quotas and usage: CCU, request budgets per minute, DataStore storage, MemoryStore memory.

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
    resp, r, err := apiClient.WorldsDataApi.GetWorldExperienceDataUsage(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.GetWorldExperienceDataUsage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceDataUsage`: GetWorldExperienceDataUsageResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.GetWorldExperienceDataUsage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceDataUsageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceDataUsageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceDatastoreEntry

> GetWorldExperienceDatastoreEntryResponse GetWorldExperienceDatastoreEntry(ctx, experienceId, store, key).Execute()

The current value of a key: version (ETag \"v<n>\"), userIds and attributes.

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
    store := "store_example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.GetWorldExperienceDatastoreEntry(context.Background(), experienceId, store, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.GetWorldExperienceDatastoreEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceDatastoreEntry`: GetWorldExperienceDatastoreEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.GetWorldExperienceDatastoreEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**store** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceDatastoreEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**GetWorldExperienceDatastoreEntryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceDatastoreEntryVersion

> GetWorldExperienceDatastoreEntryResponse GetWorldExperienceDatastoreEntryVersion(ctx, experienceId, store, key, version).Execute()

One version of a key, including a removed one.

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
    store := "store_example" // string | 
    key := "key_example" // string | 
    version := "version_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.GetWorldExperienceDatastoreEntryVersion(context.Background(), experienceId, store, key, version).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.GetWorldExperienceDatastoreEntryVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceDatastoreEntryVersion`: GetWorldExperienceDatastoreEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.GetWorldExperienceDatastoreEntryVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**store** | **string** |  | 
**key** | **string** |  | 
**version** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceDatastoreEntryVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

**GetWorldExperienceDatastoreEntryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceMemoryMap

> GetWorldExperienceMemoryMapResponse GetWorldExperienceMemoryMap(ctx, experienceId, map_).Execute()

Items of a MemoryStore hash map (cursor-paged, unordered).

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
    map_ := "map__example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.GetWorldExperienceMemoryMap(context.Background(), experienceId, map_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.GetWorldExperienceMemoryMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceMemoryMap`: GetWorldExperienceMemoryMapResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.GetWorldExperienceMemoryMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**map_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceMemoryMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperienceMemoryMapResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceMemoryMapByKey

> GetWorldExperienceMemoryMapByKeyResponse GetWorldExperienceMemoryMapByKey(ctx, experienceId, map_, key).Execute()

One MemoryStore hash map item with its version and expiry.

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
    map_ := "map__example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.GetWorldExperienceMemoryMapByKey(context.Background(), experienceId, map_, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.GetWorldExperienceMemoryMapByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceMemoryMapByKey`: GetWorldExperienceMemoryMapByKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.GetWorldExperienceMemoryMapByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**map_** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceMemoryMapByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**GetWorldExperienceMemoryMapByKeyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceMemoryQueue

> GetWorldExperienceMemoryQueueResponse GetWorldExperienceMemoryQueue(ctx, experienceId, queue).Execute()

Queue length (visible, in flight) and the head, without hiding anything.

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
    queue := "queue_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.GetWorldExperienceMemoryQueue(context.Background(), experienceId, queue).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.GetWorldExperienceMemoryQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceMemoryQueue`: GetWorldExperienceMemoryQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.GetWorldExperienceMemoryQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**queue** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceMemoryQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperienceMemoryQueueResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceMemorySorted

> GetWorldExperienceMemoryMapResponse GetWorldExperienceMemorySorted(ctx, experienceId, map_).Execute()

A MemoryStore sorted map in order (sortKey, then key); ?order=&cursor=&limit=.

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
    map_ := "map__example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.GetWorldExperienceMemorySorted(context.Background(), experienceId, map_).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.GetWorldExperienceMemorySorted``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceMemorySorted`: GetWorldExperienceMemoryMapResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.GetWorldExperienceMemorySorted`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**map_** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceMemorySortedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperienceMemoryMapResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceMemorySortedByKey

> GetWorldExperienceMemoryMapByKeyResponse GetWorldExperienceMemorySortedByKey(ctx, experienceId, map_, key).Execute()

One sorted map item.

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
    map_ := "map__example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.GetWorldExperienceMemorySortedByKey(context.Background(), experienceId, map_, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.GetWorldExperienceMemorySortedByKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceMemorySortedByKey`: GetWorldExperienceMemoryMapByKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.GetWorldExperienceMemorySortedByKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**map_** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceMemorySortedByKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**GetWorldExperienceMemoryMapByKeyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceOrdered

> GetWorldExperienceOrderedResponse GetWorldExperienceOrdered(ctx, experienceId, store).Execute()

An OrderedStore page ranked by value; ?order=desc|asc&min=&max=&cursor=&limit=.

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
    store := "store_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.GetWorldExperienceOrdered(context.Background(), experienceId, store).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.GetWorldExperienceOrdered``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceOrdered`: GetWorldExperienceOrderedResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.GetWorldExperienceOrdered`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**store** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceOrderedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperienceOrderedResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceOrderedEntry

> WorldExperienceOrderedEntryResponse GetWorldExperienceOrderedEntry(ctx, experienceId, store, key).Execute()

One OrderedStore entry.

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
    store := "store_example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.GetWorldExperienceOrderedEntry(context.Background(), experienceId, store, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.GetWorldExperienceOrderedEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceOrderedEntry`: WorldExperienceOrderedEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.GetWorldExperienceOrderedEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**store** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceOrderedEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**WorldExperienceOrderedEntryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceDatastoreEntries

> ListWorldExperienceDatastoreEntriesResponse ListWorldExperienceDatastoreEntries(ctx, experienceId, store).Execute()

Keys of a DataStore in key order; ?prefix=&scope=&cursor=&limit=.

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
    store := "store_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.ListWorldExperienceDatastoreEntries(context.Background(), experienceId, store).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.ListWorldExperienceDatastoreEntries``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceDatastoreEntries`: ListWorldExperienceDatastoreEntriesResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.ListWorldExperienceDatastoreEntries`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**store** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceDatastoreEntriesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**ListWorldExperienceDatastoreEntriesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceDatastoreEntryVersions

> ListWorldExperienceDatastoreEntryVersionsResponse ListWorldExperienceDatastoreEntryVersions(ctx, experienceId, store, key).Execute()

Versions of a key, newest first (30-day history; the current version is kept).

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
    store := "store_example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.ListWorldExperienceDatastoreEntryVersions(context.Background(), experienceId, store, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.ListWorldExperienceDatastoreEntryVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceDatastoreEntryVersions`: ListWorldExperienceDatastoreEntryVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.ListWorldExperienceDatastoreEntryVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**store** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceDatastoreEntryVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**ListWorldExperienceDatastoreEntryVersionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceDatastores

> ListWorldExperienceDatastoresResponse ListWorldExperienceDatastores(ctx, experienceId).Execute()

The DataStores of an experience, with key counts and bytes.

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
    resp, r, err := apiClient.WorldsDataApi.ListWorldExperienceDatastores(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.ListWorldExperienceDatastores``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceDatastores`: ListWorldExperienceDatastoresResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.ListWorldExperienceDatastores`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceDatastoresRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceDatastoresResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceDatastoreEntry

> UpdateWorldExperienceDatastoreEntryResponse UpdateWorldExperienceDatastoreEntry(ctx, experienceId, store, key).Execute()

Write a new version. If-Match \"v<n>\" (or ifVersion) writes only at that version; If-None-Match: * only if absent. 412 on a mismatch.

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
    store := "store_example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.UpdateWorldExperienceDatastoreEntry(context.Background(), experienceId, store, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.UpdateWorldExperienceDatastoreEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceDatastoreEntry`: UpdateWorldExperienceDatastoreEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.UpdateWorldExperienceDatastoreEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**store** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceDatastoreEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**UpdateWorldExperienceDatastoreEntryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceMemoryMap

> UpdateWorldExperienceMemoryMapResponse UpdateWorldExperienceMemoryMap(ctx, experienceId, map_, key).Execute()

Set a hash map item (ttlSeconds required; ifVersion for compare-and-set).

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
    map_ := "map__example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.UpdateWorldExperienceMemoryMap(context.Background(), experienceId, map_, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.UpdateWorldExperienceMemoryMap``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceMemoryMap`: UpdateWorldExperienceMemoryMapResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.UpdateWorldExperienceMemoryMap`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**map_** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceMemoryMapRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**UpdateWorldExperienceMemoryMapResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceMemorySorted

> UpdateWorldExperienceMemoryMapResponse UpdateWorldExperienceMemorySorted(ctx, experienceId, map_, key).Execute()

Set a sorted map item (ttlSeconds, optional sortKey and ifVersion).

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
    map_ := "map__example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.UpdateWorldExperienceMemorySorted(context.Background(), experienceId, map_, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.UpdateWorldExperienceMemorySorted``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceMemorySorted`: UpdateWorldExperienceMemoryMapResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.UpdateWorldExperienceMemorySorted`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**map_** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceMemorySortedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**UpdateWorldExperienceMemoryMapResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceOrderedEntry

> WorldExperienceOrderedEntryResponse UpdateWorldExperienceOrderedEntry(ctx, experienceId, store, key).Execute()

Set an OrderedStore entry to a whole number.

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
    store := "store_example" // string | 
    key := "key_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsDataApi.UpdateWorldExperienceOrderedEntry(context.Background(), experienceId, store, key).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsDataApi.UpdateWorldExperienceOrderedEntry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceOrderedEntry`: WorldExperienceOrderedEntryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsDataApi.UpdateWorldExperienceOrderedEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**store** | **string** |  | 
**key** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceOrderedEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**WorldExperienceOrderedEntryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

