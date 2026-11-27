# \WorldsHttpApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorldExperienceHttpDomain**](WorldsHttpApi.md#CreateWorldExperienceHttpDomain) | **Post** /v1/worlds/experiences/{experienceId}/http/domains | Declare a domain server scripts may call (example.com or *.example.com). It is reviewed by a person before it works.
[**DeleteWorldExperienceHttpDomain**](WorldsHttpApi.md#DeleteWorldExperienceHttpDomain) | **Delete** /v1/worlds/experiences/{experienceId}/http/domains/{domainId} | Remove a domain: scripts lose access at once.
[**DeleteWorldExperienceHttpSecret**](WorldsHttpApi.md#DeleteWorldExperienceHttpSecret) | **Delete** /v1/worlds/experiences/{experienceId}/http/secrets/{name} | Delete a secret.
[**GetWorldExperienceHttp**](WorldsHttpApi.md#GetWorldExperienceHttp) | **Get** /v1/worlds/experiences/{experienceId}/http | Http settings: declared domains with their review state, secret names (never values), limits and rules.
[**UpdateWorldExperienceHttpSecret**](WorldsHttpApi.md#UpdateWorldExperienceHttpSecret) | **Put** /v1/worlds/experiences/{experienceId}/http/secrets/{name} | Set a secret (UPPER_SNAKE name). Scripts reference it as Http.secret(\&quot;NAME\&quot;); the value is never returned.



## CreateWorldExperienceHttpDomain

> CreateWorldExperienceHttpDomainResponse CreateWorldExperienceHttpDomain(ctx, experienceId).Execute()

Declare a domain server scripts may call (example.com or *.example.com). It is reviewed by a person before it works.

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
    resp, r, err := apiClient.WorldsHttpApi.CreateWorldExperienceHttpDomain(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsHttpApi.CreateWorldExperienceHttpDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceHttpDomain`: CreateWorldExperienceHttpDomainResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsHttpApi.CreateWorldExperienceHttpDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceHttpDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceHttpDomainResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorldExperienceHttpDomain

> DeleteSpatialSceneGrantResponse DeleteWorldExperienceHttpDomain(ctx, experienceId, domainId).Execute()

Remove a domain: scripts lose access at once.

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
    domainId := "domainId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsHttpApi.DeleteWorldExperienceHttpDomain(context.Background(), experienceId, domainId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsHttpApi.DeleteWorldExperienceHttpDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldExperienceHttpDomain`: DeleteSpatialSceneGrantResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsHttpApi.DeleteWorldExperienceHttpDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**domainId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldExperienceHttpDomainRequest struct via the builder pattern


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


## DeleteWorldExperienceHttpSecret

> DeleteSpatialSceneGrantResponse DeleteWorldExperienceHttpSecret(ctx, experienceId, name).Execute()

Delete a secret.

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsHttpApi.DeleteWorldExperienceHttpSecret(context.Background(), experienceId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsHttpApi.DeleteWorldExperienceHttpSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldExperienceHttpSecret`: DeleteSpatialSceneGrantResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsHttpApi.DeleteWorldExperienceHttpSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldExperienceHttpSecretRequest struct via the builder pattern


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


## GetWorldExperienceHttp

> GetWorldExperienceHttpResponse GetWorldExperienceHttp(ctx, experienceId).Execute()

Http settings: declared domains with their review state, secret names (never values), limits and rules.

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
    resp, r, err := apiClient.WorldsHttpApi.GetWorldExperienceHttp(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsHttpApi.GetWorldExperienceHttp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceHttp`: GetWorldExperienceHttpResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsHttpApi.GetWorldExperienceHttp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceHttpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceHttpResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceHttpSecret

> UpdateWorldExperienceHttpSecretResponse UpdateWorldExperienceHttpSecret(ctx, experienceId, name).Execute()

Set a secret (UPPER_SNAKE name). Scripts reference it as Http.secret(\"NAME\"); the value is never returned.

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
    name := "name_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsHttpApi.UpdateWorldExperienceHttpSecret(context.Background(), experienceId, name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsHttpApi.UpdateWorldExperienceHttpSecret``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceHttpSecret`: UpdateWorldExperienceHttpSecretResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsHttpApi.UpdateWorldExperienceHttpSecret`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**name** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceHttpSecretRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**UpdateWorldExperienceHttpSecretResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

