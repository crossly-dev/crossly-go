# \MagicApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMagicAutoDecisionUndo**](MagicApi.md#CreateMagicAutoDecisionUndo) | **Post** /v1/magic/auto/decisions/{decisionId}/undo | Undo an automated Magic List decision, deleting its unpublished draft.
[**CreateMagicAutoItem**](MagicApi.md#CreateMagicAutoItem) | **Post** /v1/magic/auto/items | Scan one item and match it automatically (drafts only, never publishes).
[**CreateMagicAutoRun**](MagicApi.md#CreateMagicAutoRun) | **Post** /v1/magic/auto/runs/{runId} | Score a Magic List run automatically; a confident match in auto mode becomes a draft.
[**CreateMagicAutoRunSwipe**](MagicApi.md#CreateMagicAutoRunSwipe) | **Post** /v1/magic/auto/runs/{runId}/swipes | Record a seller swipe session (kept and rejected match keys) for a Magic List run.
[**CreateMagicScan**](MagicApi.md#CreateMagicScan) | **Post** /v1/magic/scan | Run a Magic List image scan.
[**CreateMagicScanSynthesize**](MagicApi.md#CreateMagicScanSynthesize) | **Post** /v1/magic/scan/{runId}/synthesize | Synthesize a draft from confirmed matches.
[**GetMagicAutoSetting**](MagicApi.md#GetMagicAutoSetting) | **Get** /v1/magic/auto/settings | The seller’s Magic List automated-matching mode.
[**GetMagicDraft**](MagicApi.md#GetMagicDraft) | **Get** /v1/magic/drafts/{draftId} | Get a synthesized Magic List draft.
[**ListMagicAutoRunDecisions**](MagicApi.md#ListMagicAutoRunDecisions) | **Get** /v1/magic/auto/runs/{runId}/decisions | Every match decision round (automated and seller) for a Magic List run.
[**ListMagicRecent**](MagicApi.md#ListMagicRecent) | **Get** /v1/magic/recent | Recent Magic List scans for this seller.
[**UpdateMagicAutoSetting**](MagicApi.md#UpdateMagicAutoSetting) | **Put** /v1/magic/auto/settings | Set the Magic List automated-matching mode (off | auto | review_all).



## CreateMagicAutoDecisionUndo

> CreateMagicAutoDecisionUndoResponse CreateMagicAutoDecisionUndo(ctx, decisionId).Execute()

Undo an automated Magic List decision, deleting its unpublished draft.

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
    decisionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MagicApi.CreateMagicAutoDecisionUndo(context.Background(), decisionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MagicApi.CreateMagicAutoDecisionUndo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMagicAutoDecisionUndo`: CreateMagicAutoDecisionUndoResponse
    fmt.Fprintf(os.Stdout, "Response from `MagicApi.CreateMagicAutoDecisionUndo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decisionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMagicAutoDecisionUndoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateMagicAutoDecisionUndoResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateMagicAutoItem

> CreateMagicAutoItemResponse CreateMagicAutoItem(ctx).Execute()

Scan one item and match it automatically (drafts only, never publishes).

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
    resp, r, err := apiClient.MagicApi.CreateMagicAutoItem(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MagicApi.CreateMagicAutoItem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMagicAutoItem`: CreateMagicAutoItemResponse
    fmt.Fprintf(os.Stdout, "Response from `MagicApi.CreateMagicAutoItem`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMagicAutoItemRequest struct via the builder pattern


### Return type

**CreateMagicAutoItemResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateMagicAutoRun

> CreateMagicAutoRunResponse CreateMagicAutoRun(ctx, runId).Execute()

Score a Magic List run automatically; a confident match in auto mode becomes a draft.

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
    resp, r, err := apiClient.MagicApi.CreateMagicAutoRun(context.Background(), runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MagicApi.CreateMagicAutoRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMagicAutoRun`: CreateMagicAutoRunResponse
    fmt.Fprintf(os.Stdout, "Response from `MagicApi.CreateMagicAutoRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMagicAutoRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateMagicAutoRunResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateMagicAutoRunSwipe

> CreateMagicAutoRunSwipeResponse CreateMagicAutoRunSwipe(ctx, runId).Execute()

Record a seller swipe session (kept and rejected match keys) for a Magic List run.

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
    resp, r, err := apiClient.MagicApi.CreateMagicAutoRunSwipe(context.Background(), runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MagicApi.CreateMagicAutoRunSwipe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMagicAutoRunSwipe`: CreateMagicAutoRunSwipeResponse
    fmt.Fprintf(os.Stdout, "Response from `MagicApi.CreateMagicAutoRunSwipe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMagicAutoRunSwipeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateMagicAutoRunSwipeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


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

**CreateMagicScanResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**CreateMagicScanSynthesizeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetMagicAutoSetting

> GetMagicAutoSettingResponse GetMagicAutoSetting(ctx).Execute()

The seller’s Magic List automated-matching mode.



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
    resp, r, err := apiClient.MagicApi.GetMagicAutoSetting(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MagicApi.GetMagicAutoSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMagicAutoSetting`: GetMagicAutoSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `MagicApi.GetMagicAutoSetting`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMagicAutoSettingRequest struct via the builder pattern


### Return type

**GetMagicAutoSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
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

**GetMagicDraftResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListMagicAutoRunDecisions

> ListMagicAutoRunDecisionsResponse ListMagicAutoRunDecisions(ctx, runId).Execute()

Every match decision round (automated and seller) for a Magic List run.

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
    resp, r, err := apiClient.MagicApi.ListMagicAutoRunDecisions(context.Background(), runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MagicApi.ListMagicAutoRunDecisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMagicAutoRunDecisions`: ListMagicAutoRunDecisionsResponse
    fmt.Fprintf(os.Stdout, "Response from `MagicApi.ListMagicAutoRunDecisions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListMagicAutoRunDecisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListMagicAutoRunDecisionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListMagicRecent

> ListMagicRecentResponse ListMagicRecent(ctx).Execute()

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
    // response from `ListMagicRecent`: ListMagicRecentResponse
    fmt.Fprintf(os.Stdout, "Response from `MagicApi.ListMagicRecent`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMagicRecentRequest struct via the builder pattern


### Return type

**ListMagicRecentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateMagicAutoSetting

> GetMagicAutoSettingResponse UpdateMagicAutoSetting(ctx).UpdateMagicAutoSettingRequest(updateMagicAutoSettingRequest).Execute()

Set the Magic List automated-matching mode (off | auto | review_all).

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
    updateMagicAutoSettingRequest := *openapiclient.NewUpdateMagicAutoSettingRequest("Mode_example") // UpdateMagicAutoSettingRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MagicApi.UpdateMagicAutoSetting(context.Background()).UpdateMagicAutoSettingRequest(updateMagicAutoSettingRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MagicApi.UpdateMagicAutoSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateMagicAutoSetting`: GetMagicAutoSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `MagicApi.UpdateMagicAutoSetting`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMagicAutoSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateMagicAutoSettingRequest** | **UpdateMagicAutoSettingRequest** |  | 

### Return type

**GetMagicAutoSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

