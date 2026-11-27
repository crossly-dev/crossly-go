# \CreatorApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCreatorExperienceKey**](CreatorApi.md#CreateCreatorExperienceKey) | **Post** /v1/creator/experience-keys | Issue an experience API key for a back-end or CI: name, creator scopes (worlds:*, items:write — only ones your token holds), an IP allowlist (CIDRs; 0.0.0.0/0 + ::/0 for anywhere), optional experienceIds and expiresInDays. The token is in the response ONCE.
[**CreateCreatorExperienceKeyRotate**](CreatorApi.md#CreateCreatorExperienceKeyRotate) | **Post** /v1/creator/experience-keys/{id}/rotate | Rotate a key: a replacement with the same settings (its token shown ONCE); the old key keeps working for overlapHours (default 24, 0 &#x3D; revoke now, max 168).
[**DeleteCreatorExperienceKey**](CreatorApi.md#DeleteCreatorExperienceKey) | **Delete** /v1/creator/experience-keys/{id} | Revoke an experience key now. Never needs a plan: a credential must always be revocable.
[**GetCreatorEntitlement**](CreatorApi.md#GetCreatorEntitlement) | **Get** /v1/creator/entitlements | What you can do in Creator Studio: your plan, whether you can author, grandfathering, and every limit with used and remaining. Read by web, apps, SDK and MCP alike.
[**ListCreatorExperienceKeys**](CreatorApi.md#ListCreatorExperienceKeys) | **Get** /v1/creator/experience-keys | Your experience API keys, newest first (never the secret). ?includeRevoked&#x3D;true for the history; ?experienceId&#x3D; for the keys that reach one game (a game&#39;s own keys are narrowed to it alone). Cursor-paged: pass pagination.nextCursor back as ?cursor&#x3D;.



## CreateCreatorExperienceKey

> CreateCreatorExperienceKeyResponse CreateCreatorExperienceKey(ctx).Execute()

Issue an experience API key for a back-end or CI: name, creator scopes (worlds:*, items:write — only ones your token holds), an IP allowlist (CIDRs; 0.0.0.0/0 + ::/0 for anywhere), optional experienceIds and expiresInDays. The token is in the response ONCE.

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
    resp, r, err := apiClient.CreatorApi.CreateCreatorExperienceKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorApi.CreateCreatorExperienceKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorExperienceKey`: CreateCreatorExperienceKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorApi.CreateCreatorExperienceKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorExperienceKeyRequest struct via the builder pattern


### Return type

**CreateCreatorExperienceKeyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCreatorExperienceKeyRotate

> CreateCreatorExperienceKeyRotateResponse CreateCreatorExperienceKeyRotate(ctx, id).Execute()

Rotate a key: a replacement with the same settings (its token shown ONCE); the old key keeps working for overlapHours (default 24, 0 = revoke now, max 168).

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
    resp, r, err := apiClient.CreatorApi.CreateCreatorExperienceKeyRotate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorApi.CreateCreatorExperienceKeyRotate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCreatorExperienceKeyRotate`: CreateCreatorExperienceKeyRotateResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorApi.CreateCreatorExperienceKeyRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCreatorExperienceKeyRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateCreatorExperienceKeyRotateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteCreatorExperienceKey

> DeleteCreatorExperienceKeyResponse DeleteCreatorExperienceKey(ctx, id).Execute()

Revoke an experience key now. Never needs a plan: a credential must always be revocable.

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
    resp, r, err := apiClient.CreatorApi.DeleteCreatorExperienceKey(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorApi.DeleteCreatorExperienceKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCreatorExperienceKey`: DeleteCreatorExperienceKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorApi.DeleteCreatorExperienceKey`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCreatorExperienceKeyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteCreatorExperienceKeyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCreatorEntitlement

> GetCreatorEntitlementResponse GetCreatorEntitlement(ctx).Execute()

What you can do in Creator Studio: your plan, whether you can author, grandfathering, and every limit with used and remaining. Read by web, apps, SDK and MCP alike.

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
    resp, r, err := apiClient.CreatorApi.GetCreatorEntitlement(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorApi.GetCreatorEntitlement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCreatorEntitlement`: GetCreatorEntitlementResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorApi.GetCreatorEntitlement`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCreatorEntitlementRequest struct via the builder pattern


### Return type

**GetCreatorEntitlementResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCreatorExperienceKeys

> ListCreatorExperienceKeysResponse ListCreatorExperienceKeys(ctx).Execute()

Your experience API keys, newest first (never the secret). ?includeRevoked=true for the history; ?experienceId= for the keys that reach one game (a game's own keys are narrowed to it alone). Cursor-paged: pass pagination.nextCursor back as ?cursor=.

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
    resp, r, err := apiClient.CreatorApi.ListCreatorExperienceKeys(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CreatorApi.ListCreatorExperienceKeys``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCreatorExperienceKeys`: ListCreatorExperienceKeysResponse
    fmt.Fprintf(os.Stdout, "Response from `CreatorApi.ListCreatorExperienceKeys`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCreatorExperienceKeysRequest struct via the builder pattern


### Return type

**ListCreatorExperienceKeysResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

