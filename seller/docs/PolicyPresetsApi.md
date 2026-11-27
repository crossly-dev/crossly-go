# \PolicyPresetsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGeoLocationCheck**](PolicyPresetsApi.md#CreateGeoLocationCheck) | **Post** /v1/geo/location-check | Check whether a shipping policy ships to an address.
[**CreatePolicyPreset**](PolicyPresetsApi.md#CreatePolicyPreset) | **Post** /v1/policy-presets | Create a return / shipping / payment policy preset.
[**DeletePolicyPreset**](PolicyPresetsApi.md#DeletePolicyPreset) | **Delete** /v1/policy-presets/{id} | Delete a policy preset. Listings that used it switch to the default preset of its kind.
[**GetPolicy**](PolicyPresetsApi.md#GetPolicy) | **Get** /v1/policies | The seller&#39;s shipping / return / payment policy presets by kind, with each kind&#39;s default.
[**ListPolicyPresets**](PolicyPresetsApi.md#ListPolicyPresets) | **Get** /v1/policy-presets | List the seller&#39;s return / shipping / payment policy presets. meta.migrationNotices lists defaults the policy migration demoted (kind still has no default).
[**UpdatePolicyPreset**](PolicyPresetsApi.md#UpdatePolicyPreset) | **Patch** /v1/policy-presets/{id} | Update a policy preset.



## CreateGeoLocationCheck

> CreateGeoLocationCheckResponse CreateGeoLocationCheck(ctx).Execute()

Check whether a shipping policy ships to an address.



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
    resp, r, err := apiClient.PolicyPresetsApi.CreateGeoLocationCheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyPresetsApi.CreateGeoLocationCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGeoLocationCheck`: CreateGeoLocationCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyPresetsApi.CreateGeoLocationCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateGeoLocationCheckRequest struct via the builder pattern


### Return type

**CreateGeoLocationCheckResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreatePolicyPreset

> CreatePolicyPresetResponse CreatePolicyPreset(ctx).Execute()

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
    // response from `CreatePolicyPreset`: CreatePolicyPresetResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyPresetsApi.CreatePolicyPreset`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyPresetRequest struct via the builder pattern


### Return type

**CreatePolicyPresetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeletePolicyPreset

> DeletePolicyPresetResponse DeletePolicyPreset(ctx, id).Execute()

Delete a policy preset. Listings that used it switch to the default preset of its kind.

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
    // response from `DeletePolicyPreset`: DeletePolicyPresetResponse
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

**DeletePolicyPresetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetPolicy

> GetPolicyResponse GetPolicy(ctx).Execute()

The seller's shipping / return / payment policy presets by kind, with each kind's default.



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
    resp, r, err := apiClient.PolicyPresetsApi.GetPolicy(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PolicyPresetsApi.GetPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetPolicy`: GetPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `PolicyPresetsApi.GetPolicy`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPolicyRequest struct via the builder pattern


### Return type

**GetPolicyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListPolicyPresets

> ListPolicyPresetsResponse ListPolicyPresets(ctx).Kind(kind).Execute()

List the seller's return / shipping / payment policy presets. meta.migrationNotices lists defaults the policy migration demoted (kind still has no default).

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

> CreatePolicyPresetResponse UpdatePolicyPreset(ctx, id).Execute()

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
    // response from `UpdatePolicyPreset`: CreatePolicyPresetResponse
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

**CreatePolicyPresetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

