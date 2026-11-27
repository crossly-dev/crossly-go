# \WorldsGalleryApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorldExperienceGallery**](WorldsGalleryApi.md#CreateWorldExperienceGallery) | **Post** /v1/worlds/experiences/{experienceId}/gallery | Add an uploaded screenshot or video (your own upload&#39;s publicUrl; a video may carry a poster image). Images are scanned at once; every item waits for a person and shows on the game&#39;s page once approved.
[**CreateWorldExperienceGalleryUploadUrl**](WorldsGalleryApi.md#CreateWorldExperienceGalleryUploadUrl) | **Post** /v1/worlds/experiences/{experienceId}/gallery/upload-url | A presigned upload for a gallery screenshot (JPEG/PNG/WebP/GIF, ≤ 10 MB) or video (MP4/QuickTime/WebM): PUT the file to uploadUrl with the same Content-Type and exact Content-Length, then add publicUrl to the gallery.
[**DeleteWorldExperienceGallery**](WorldsGalleryApi.md#DeleteWorldExperienceGallery) | **Delete** /v1/worlds/experiences/{experienceId}/gallery/{mediaId} | Take an item off the game&#39;s page now.
[**GetWorldExperienceGallery**](WorldsGalleryApi.md#GetWorldExperienceGallery) | **Get** /v1/worlds/experiences/{experienceId}/gallery | The game&#39;s gallery: every screenshot and video with its review state (pending / approved / rejected with the reviewer&#39;s note), position and caption, plus the limits.
[**UpdateWorldExperienceGallery**](WorldsGalleryApi.md#UpdateWorldExperienceGallery) | **Patch** /v1/worlds/experiences/{experienceId}/gallery/{mediaId} | Change an item&#39;s caption (filtered; an approved item with a new caption goes back to review).
[**UpdateWorldExperienceGalleryOrder**](WorldsGalleryApi.md#UpdateWorldExperienceGalleryOrder) | **Put** /v1/worlds/experiences/{experienceId}/gallery/order | Set the order on the game&#39;s page: the ids you list first, the rest after in their current order.
[**UpdateWorldExperienceGalleryRole**](WorldsGalleryApi.md#UpdateWorldExperienceGalleryRole) | **Put** /v1/worlds/experiences/{experienceId}/gallery/{mediaId}/role | Make an item the game&#39;s icon or thumbnail (an image), its trailer (a video), or put it back in the gallery. The previous holder of that role goes back to the gallery.
[**UpdateWorldExperienceGalleryTrailerUrl**](WorldsGalleryApi.md#UpdateWorldExperienceGalleryTrailerUrl) | **Put** /v1/worlds/experiences/{experienceId}/gallery/trailer-url | Set the game&#39;s trailer to a YouTube or Vimeo video link (https; youtube.com, youtu.be, vimeo.com, player.vimeo.com only; any other host is 400 trailer_host_not_allowed). It waits for a person like any gallery item; the current trailer goes back to the gallery (an uploaded clip) or off the page (a link).



## CreateWorldExperienceGallery

> WorldExperienceGalleryResponse CreateWorldExperienceGallery(ctx, experienceId).Execute()

Add an uploaded screenshot or video (your own upload's publicUrl; a video may carry a poster image). Images are scanned at once; every item waits for a person and shows on the game's page once approved.

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
    resp, r, err := apiClient.WorldsGalleryApi.CreateWorldExperienceGallery(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsGalleryApi.CreateWorldExperienceGallery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceGallery`: WorldExperienceGalleryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsGalleryApi.CreateWorldExperienceGallery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceGalleryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceGalleryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceGalleryUploadUrl

> CreateWorldExperienceGalleryUploadUrlResponse CreateWorldExperienceGalleryUploadUrl(ctx, experienceId).Execute()

A presigned upload for a gallery screenshot (JPEG/PNG/WebP/GIF, ≤ 10 MB) or video (MP4/QuickTime/WebM): PUT the file to uploadUrl with the same Content-Type and exact Content-Length, then add publicUrl to the gallery.

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
    resp, r, err := apiClient.WorldsGalleryApi.CreateWorldExperienceGalleryUploadUrl(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsGalleryApi.CreateWorldExperienceGalleryUploadUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceGalleryUploadUrl`: CreateWorldExperienceGalleryUploadUrlResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsGalleryApi.CreateWorldExperienceGalleryUploadUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceGalleryUploadUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceGalleryUploadUrlResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorldExperienceGallery

> DeleteSpatialSceneGrantResponse DeleteWorldExperienceGallery(ctx, experienceId, mediaId).Execute()

Take an item off the game's page now.

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
    mediaId := "mediaId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsGalleryApi.DeleteWorldExperienceGallery(context.Background(), experienceId, mediaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsGalleryApi.DeleteWorldExperienceGallery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldExperienceGallery`: DeleteSpatialSceneGrantResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsGalleryApi.DeleteWorldExperienceGallery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**mediaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldExperienceGalleryRequest struct via the builder pattern


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


## GetWorldExperienceGallery

> GetWorldExperienceGalleryResponse GetWorldExperienceGallery(ctx, experienceId).Execute()

The game's gallery: every screenshot and video with its review state (pending / approved / rejected with the reviewer's note), position and caption, plus the limits.

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
    resp, r, err := apiClient.WorldsGalleryApi.GetWorldExperienceGallery(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsGalleryApi.GetWorldExperienceGallery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceGallery`: GetWorldExperienceGalleryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsGalleryApi.GetWorldExperienceGallery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceGalleryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceGalleryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceGallery

> WorldExperienceGalleryResponse UpdateWorldExperienceGallery(ctx, experienceId, mediaId).Execute()

Change an item's caption (filtered; an approved item with a new caption goes back to review).

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
    mediaId := "mediaId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsGalleryApi.UpdateWorldExperienceGallery(context.Background(), experienceId, mediaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsGalleryApi.UpdateWorldExperienceGallery``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceGallery`: WorldExperienceGalleryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsGalleryApi.UpdateWorldExperienceGallery`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**mediaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceGalleryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorldExperienceGalleryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceGalleryOrder

> UpdateWorldExperienceGalleryRoleResponse UpdateWorldExperienceGalleryOrder(ctx, experienceId).Execute()

Set the order on the game's page: the ids you list first, the rest after in their current order.

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
    resp, r, err := apiClient.WorldsGalleryApi.UpdateWorldExperienceGalleryOrder(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsGalleryApi.UpdateWorldExperienceGalleryOrder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceGalleryOrder`: UpdateWorldExperienceGalleryRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsGalleryApi.UpdateWorldExperienceGalleryOrder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceGalleryOrderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateWorldExperienceGalleryRoleResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceGalleryRole

> UpdateWorldExperienceGalleryRoleResponse UpdateWorldExperienceGalleryRole(ctx, experienceId, mediaId).Execute()

Make an item the game's icon or thumbnail (an image), its trailer (a video), or put it back in the gallery. The previous holder of that role goes back to the gallery.

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
    mediaId := "mediaId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsGalleryApi.UpdateWorldExperienceGalleryRole(context.Background(), experienceId, mediaId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsGalleryApi.UpdateWorldExperienceGalleryRole``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceGalleryRole`: UpdateWorldExperienceGalleryRoleResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsGalleryApi.UpdateWorldExperienceGalleryRole`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**mediaId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceGalleryRoleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**UpdateWorldExperienceGalleryRoleResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceGalleryTrailerUrl

> WorldExperienceGalleryResponse UpdateWorldExperienceGalleryTrailerUrl(ctx, experienceId).Execute()

Set the game's trailer to a YouTube or Vimeo video link (https; youtube.com, youtu.be, vimeo.com, player.vimeo.com only; any other host is 400 trailer_host_not_allowed). It waits for a person like any gallery item; the current trailer goes back to the gallery (an uploaded clip) or off the page (a link).

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
    resp, r, err := apiClient.WorldsGalleryApi.UpdateWorldExperienceGalleryTrailerUrl(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsGalleryApi.UpdateWorldExperienceGalleryTrailerUrl``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceGalleryTrailerUrl`: WorldExperienceGalleryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsGalleryApi.UpdateWorldExperienceGalleryTrailerUrl`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceGalleryTrailerUrlRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorldExperienceGalleryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

