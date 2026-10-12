# \PolicyPresetsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyPreset**](PolicyPresetsApi.md#CreatePolicyPreset) | **Post** /v1/policy-presets | Create a return / shipping / payment policy preset.
[**DeletePolicyPreset**](PolicyPresetsApi.md#DeletePolicyPreset) | **Delete** /v1/policy-presets/{id} | Delete a policy preset.
[**ListPolicyPresets**](PolicyPresetsApi.md#ListPolicyPresets) | **Get** /v1/policy-presets | List the seller&#39;s return / shipping / payment policy presets.
[**UpdatePolicyPreset**](PolicyPresetsApi.md#UpdatePolicyPreset) | **Patch** /v1/policy-presets/{id} | Update a policy preset.



## CreatePolicyPreset

> ListPolicyPresetsItem CreatePolicyPreset(ctx).Execute()

Create a return / shipping / payment policy preset.

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
    resp, r, err := apiClient.PolicyPresetsApi.CreatePolicyPreset(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyPresetsApi.CreatePolicyPreset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyPreset`: ListPolicyPresetsItem
    fmt.Fprintf(os.Stdout, "Response from `PolicyPresetsApi.CreatePolicyPreset`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyPresetRequest struct via the builder pattern


### Return type

**ListPolicyPresetsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeletePolicyPreset

> OkResponse DeletePolicyPreset(ctx, id).Execute()

Delete a policy preset.

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
    resp, r, err := apiClient.PolicyPresetsApi.DeletePolicyPreset(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyPresetsApi.DeletePolicyPreset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeletePolicyPreset`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyPresetsApi.DeletePolicyPreset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePolicyPresetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListPolicyPresets

> ListPolicyPresetsResponse ListPolicyPresets(ctx).Kind(kind).Execute()

List the seller's return / shipping / payment policy presets.

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
    kind := "kind_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PolicyPresetsApi.ListPolicyPresets(context.Background()).Kind(kind).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyPresetsApi.ListPolicyPresets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicyPresets`: ListPolicyPresetsResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyPresetsApi.ListPolicyPresets`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyPresetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** |  | 

### Return type

**ListPolicyPresetsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdatePolicyPreset

> ListPolicyPresetsItem UpdatePolicyPreset(ctx, id).Execute()

Update a policy preset.

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
    resp, r, err := apiClient.PolicyPresetsApi.UpdatePolicyPreset(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyPresetsApi.UpdatePolicyPreset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePolicyPreset`: ListPolicyPresetsItem
    fmt.Fprintf(os.Stdout, "Response from `PolicyPresetsApi.UpdatePolicyPreset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePolicyPresetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListPolicyPresetsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

