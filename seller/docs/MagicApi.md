# \MagicApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMagicScan**](MagicApi.md#CreateMagicScan) | **Post** /v1/magic/scan | Run a Magic List image scan.
[**CreateMagicScanSynthesize**](MagicApi.md#CreateMagicScanSynthesize) | **Post** /v1/magic/scan/{runId}/synthesize | Synthesize a draft from confirmed matches.
[**GetMagicDraft**](MagicApi.md#GetMagicDraft) | **Get** /v1/magic/drafts/{draftId} | Get a synthesized Magic List draft.
[**ListMagicRecent**](MagicApi.md#ListMagicRecent) | **Get** /v1/magic/recent | Recent Magic List scans for this seller.



## CreateMagicScan

> CreateMagicScanResponse CreateMagicScan(ctx).Execute()

Run a Magic List image scan.

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
    resp, r, err := apiClient.MagicApi.CreateMagicScan(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MagicApi.CreateMagicScan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMagicScan`: CreateMagicScanResponse
    fmt.Fprintf(os.Stdout, "Response from `MagicApi.CreateMagicScan`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMagicScanRequest struct via the builder pattern


### Return type

[**CreateMagicScanResponse**](CreateMagicScanResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMagicScanSynthesize

> CreateMagicScanSynthesizeResponse CreateMagicScanSynthesize(ctx, runId).Execute()

Synthesize a draft from confirmed matches.

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
    runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MagicApi.CreateMagicScanSynthesize(context.Background(), runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MagicApi.CreateMagicScanSynthesize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMagicScanSynthesize`: CreateMagicScanSynthesizeResponse
    fmt.Fprintf(os.Stdout, "Response from `MagicApi.CreateMagicScanSynthesize`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMagicScanSynthesizeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**CreateMagicScanSynthesizeResponse**](CreateMagicScanSynthesizeResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMagicDraft

> GetMagicDraftResponse GetMagicDraft(ctx, draftId).Execute()

Get a synthesized Magic List draft.

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
    draftId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MagicApi.GetMagicDraft(context.Background(), draftId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MagicApi.GetMagicDraft``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMagicDraft`: GetMagicDraftResponse
    fmt.Fprintf(os.Stdout, "Response from `MagicApi.GetMagicDraft`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**draftId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetMagicDraftRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetMagicDraftResponse**](GetMagicDraftResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMagicRecent

> V1List ListMagicRecent(ctx).Execute()

Recent Magic List scans for this seller.

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
    resp, r, err := apiClient.MagicApi.ListMagicRecent(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MagicApi.ListMagicRecent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMagicRecent`: V1List
    fmt.Fprintf(os.Stdout, "Response from `MagicApi.ListMagicRecent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMagicRecentRequest struct via the builder pattern


### Return type

[**V1List**](V1List.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

