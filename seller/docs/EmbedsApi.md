# \EmbedsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEmbedKey**](EmbedsApi.md#CreateEmbedKey) | **Post** /v1/embeds/keys | Mint a publishable key for a site.
[**DeleteEmbedKey**](EmbedsApi.md#DeleteEmbedKey) | **Delete** /v1/embeds/keys/{id} | Revoke a publishable key.
[**ListEmbedKeys**](EmbedsApi.md#ListEmbedKeys) | **Get** /v1/embeds/keys | Your publishable keys.



## CreateEmbedKey

> CreateEmbedKeyResponse CreateEmbedKey(ctx).CreateEmbedKeyRequest(createEmbedKeyRequest).Execute()

Mint a publishable key for a site.

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
    createEmbedKeyRequest := *openapiclient.NewCreateEmbedKeyRequest("Name_example") // CreateEmbedKeyRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.EmbedsApi.CreateEmbedKey(context.Background()).CreateEmbedKeyRequest(createEmbedKeyRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmbedsApi.CreateEmbedKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmbedKey`: CreateEmbedKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `EmbedsApi.CreateEmbedKey`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmbedKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createEmbedKeyRequest** | **CreateEmbedKeyRequest** |  | 

### Return type

**CreateEmbedKeyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteEmbedKey

> DeleteEmbedKey(ctx, id).Execute()

Revoke a publishable key.



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
    resp, r, err := apiClient.EmbedsApi.DeleteEmbedKey(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmbedsApi.DeleteEmbedKey``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteEmbedKeyRequest struct via the builder pattern


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


## ListEmbedKeys

> ListEmbedKeysResponse ListEmbedKeys(ctx).Execute()

Your publishable keys.



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
    resp, r, err := apiClient.EmbedsApi.ListEmbedKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EmbedsApi.ListEmbedKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmbedKeys`: ListEmbedKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `EmbedsApi.ListEmbedKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEmbedKeysRequest struct via the builder pattern


### Return type

**ListEmbedKeysResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

