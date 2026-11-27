# \DeveloperAppsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeveloperAppCertificationSubmit**](DeveloperAppsApi.md#CreateDeveloperAppCertificationSubmit) | **Post** /v1/developer/apps/{appId}/certification/submit | Send the application to Crossly review. Needs a passing conformance run on this client version and report-menu evidence.
[**CreateDeveloperAppCertificationWithdraw**](DeveloperAppsApi.md#CreateDeveloperAppCertificationWithdraw) | **Post** /v1/developer/apps/{appId}/certification/withdraw | Take the application out of review to edit it.
[**CreateDeveloperAppClientVersion**](DeveloperAppsApi.md#CreateDeveloperAppClientVersion) | **Post** /v1/developer/apps/{appId}/client-versions | Register a new release of your certified client and start its conformance run. A pass on the same major version carries the certification automatically; a new major version goes to Crossly review; a failure keeps your previous certified version.
[**CreateDeveloperAppClientVersionRun**](DeveloperAppsApi.md#CreateDeveloperAppClientVersionRun) | **Post** /v1/developer/apps/{appId}/client-versions/{versionId}/run | Start a new conformance run for a release whose run failed, expired or was aborted.
[**CreateDeveloperAppConformanceRun**](DeveloperAppsApi.md#CreateDeveloperAppConformanceRun) | **Post** /v1/developer/apps/{appId}/conformance-runs | Start a conformance run: a test experience only your app, signed in as you, can join. Join it with POST /v1/worlds/experiences/{experienceId}/join and follow the instruction.
[**CreateDeveloperAppConformanceRunAbort**](DeveloperAppsApi.md#CreateDeveloperAppConformanceRunAbort) | **Post** /v1/developer/apps/{appId}/conformance-runs/{runId}/abort | Abort an open run.
[**CreateDeveloperAppVerificationCheckDomain**](DeveloperAppsApi.md#CreateDeveloperAppVerificationCheckDomain) | **Post** /v1/developer/apps/{appId}/verification/check-domain | Look up the DNS TXT record crossly-app-verification&#x3D;&lt;token&gt; on your domain (or _crossly.&lt;domain&gt;) now.
[**CreateDeveloperAppVerificationSubmit**](DeveloperAppsApi.md#CreateDeveloperAppVerificationSubmit) | **Post** /v1/developer/apps/{appId}/verification/submit | Send the verification to Crossly review (each scope is decided separately).
[**CreateDeveloperAppVerificationWithdraw**](DeveloperAppsApi.md#CreateDeveloperAppVerificationWithdraw) | **Post** /v1/developer/apps/{appId}/verification/withdraw | Take the verification out of review to edit it.
[**GetDeveloperAppCertification**](DeveloperAppsApi.md#GetDeveloperAppCertification) | **Get** /v1/developer/apps/{appId}/certification | The Crossly Carnival client certification application (null until you create it): status, client, evidence, the passing run and what still blocks submission.
[**GetDeveloperAppClientVersion**](DeveloperAppsApi.md#GetDeveloperAppClientVersion) | **Get** /v1/developer/apps/{appId}/client-versions/{versionId} | One client release and its re-certification.
[**GetDeveloperAppConformanceRun**](DeveloperAppsApi.md#GetDeveloperAppConformanceRun) | **Get** /v1/developer/apps/{appId}/conformance-runs/{runId} | A run&#39;s step, instruction and every check with its verdict.
[**GetDeveloperAppHealth**](DeveloperAppsApi.md#GetDeveloperAppHealth) | **Get** /v1/developer/apps/{appId}/health | Your app&#39;s API health: write-request counts by status class (24 h, 7 d) across every seller using it, top error codes with advice, a retry-storm flag, the last 50 failed requests (message and correlation id only on requests made with your own account), and from Atlas: how many people authorized the app and with which scopes (counts, never who) and the token endpoint&#39;s OAuth errors by code (e.g. invalid_grant) with the latest one.
[**GetDeveloperAppProfile**](DeveloperAppsApi.md#GetDeveloperAppProfile) | **Get** /v1/developer/apps/{appId}/profile | Your app&#39;s public directory profile: every editable field, the listing status, what still needs review, and &#x60;preview&#x60; — the profile exactly as /seller-apps or /buyer-apps will show it.
[**GetDeveloperAppReview**](DeveloperAppsApi.md#GetDeveloperAppReview) | **Get** /v1/developer/apps/{appId}/review | Your app&#39;s Crossly Carnival client certification, verification and review history in one read.
[**GetDeveloperAppVerification**](DeveloperAppsApi.md#GetDeveloperAppVerification) | **Get** /v1/developer/apps/{appId}/verification | The app verification for worlds:publish / worlds:monetize (null until created): domain proof, policies, justifications and per-scope decisions.
[**ListDeveloperAppClientVersions**](DeveloperAppsApi.md#ListDeveloperAppClientVersions) | **Get** /v1/developer/apps/{appId}/client-versions | Releases of your certified client, newest first: the re-certification status (awaiting_run, running, carried, ops_review, approved, rejected, failed, superseded), the rule that decided it and its checks.
[**ListDeveloperAppConformanceRuns**](DeveloperAppsApi.md#ListDeveloperAppConformanceRuns) | **Get** /v1/developer/apps/{appId}/conformance-runs | Your app&#39;s conformance runs, newest first.
[**UpdateDeveloperAppCertification**](DeveloperAppsApi.md#UpdateDeveloperAppCertification) | **Put** /v1/developer/apps/{appId}/certification | Create or edit the application: client name, platform, version, build and source URLs, description and evidence (report menu, chat, Policy UI). Changing the client clears the passing run.
[**UpdateDeveloperAppListing**](DeveloperAppsApi.md#UpdateDeveloperAppListing) | **Put** /v1/developer/apps/{appId}/listing | Turn \&quot;Listed publicly\&quot; on or off. On: an app Crossly approved whose name, scopes and audience are unchanged is listed at once; anything else goes to review (pending_review). Off: unlisted immediately.
[**UpdateDeveloperAppProfile**](DeveloperAppsApi.md#UpdateDeveloperAppProfile) | **Patch** /v1/developer/apps/{appId}/profile | Edit your app&#39;s directory profile: name, slug, tagline, description, longDescription (markdown), category, audience (auto|seller|buyer|both), logoUrl, bannerUrl (4:1, 1600×400), screenshots (≤6), demoVideoUrl, homepageUrl, installUrl, support/privacy/terms links, docsLinks [{label,url}] and pricingNote. Absent keys are kept, null clears. Images not on Crossly&#39;s CDN are copied onto it. Renaming a listed app or changing its audience sends it back to review.
[**UpdateDeveloperAppVerification**](DeveloperAppsApi.md#UpdateDeveloperAppVerification) | **Put** /v1/developer/apps/{appId}/verification | Create or edit the verification: domain, requested scopes with a justification each, privacy policy, terms, support email, demo.



## CreateDeveloperAppCertificationSubmit

> DeveloperAppCertificationResponse CreateDeveloperAppCertificationSubmit(ctx, appId).Execute()

Send the application to Crossly review. Needs a passing conformance run on this client version and report-menu evidence.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.CreateDeveloperAppCertificationSubmit(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.CreateDeveloperAppCertificationSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeveloperAppCertificationSubmit`: DeveloperAppCertificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.CreateDeveloperAppCertificationSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeveloperAppCertificationSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeveloperAppCertificationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDeveloperAppCertificationWithdraw

> DeveloperAppCertificationResponse CreateDeveloperAppCertificationWithdraw(ctx, appId).Execute()

Take the application out of review to edit it.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.CreateDeveloperAppCertificationWithdraw(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.CreateDeveloperAppCertificationWithdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeveloperAppCertificationWithdraw`: DeveloperAppCertificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.CreateDeveloperAppCertificationWithdraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeveloperAppCertificationWithdrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeveloperAppCertificationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDeveloperAppClientVersion

> CreateDeveloperAppClientVersionResponse CreateDeveloperAppClientVersion(ctx, appId).Execute()

Register a new release of your certified client and start its conformance run. A pass on the same major version carries the certification automatically; a new major version goes to Crossly review; a failure keeps your previous certified version.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.CreateDeveloperAppClientVersion(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.CreateDeveloperAppClientVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeveloperAppClientVersion`: CreateDeveloperAppClientVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.CreateDeveloperAppClientVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeveloperAppClientVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateDeveloperAppClientVersionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDeveloperAppClientVersionRun

> CreateDeveloperAppClientVersionResponse CreateDeveloperAppClientVersionRun(ctx, appId, versionId).Execute()

Start a new conformance run for a release whose run failed, expired or was aborted.

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
    appId := "appId_example" // string | 
    versionId := "versionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.CreateDeveloperAppClientVersionRun(context.Background(), appId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.CreateDeveloperAppClientVersionRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeveloperAppClientVersionRun`: CreateDeveloperAppClientVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.CreateDeveloperAppClientVersionRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeveloperAppClientVersionRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateDeveloperAppClientVersionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDeveloperAppConformanceRun

> GetDeveloperAppConformanceRunResponse CreateDeveloperAppConformanceRun(ctx, appId).Execute()

Start a conformance run: a test experience only your app, signed in as you, can join. Join it with POST /v1/worlds/experiences/{experienceId}/join and follow the instruction.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.CreateDeveloperAppConformanceRun(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.CreateDeveloperAppConformanceRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeveloperAppConformanceRun`: GetDeveloperAppConformanceRunResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.CreateDeveloperAppConformanceRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeveloperAppConformanceRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetDeveloperAppConformanceRunResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDeveloperAppConformanceRunAbort

> CreateDeveloperAppConformanceRunAbortResponse CreateDeveloperAppConformanceRunAbort(ctx, appId, runId).Execute()

Abort an open run.

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
    appId := "appId_example" // string | 
    runId := "runId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.CreateDeveloperAppConformanceRunAbort(context.Background(), appId, runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.CreateDeveloperAppConformanceRunAbort``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeveloperAppConformanceRunAbort`: CreateDeveloperAppConformanceRunAbortResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.CreateDeveloperAppConformanceRunAbort`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeveloperAppConformanceRunAbortRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateDeveloperAppConformanceRunAbortResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDeveloperAppVerificationCheckDomain

> DeveloperAppVerificationResponse CreateDeveloperAppVerificationCheckDomain(ctx, appId).Execute()

Look up the DNS TXT record crossly-app-verification=<token> on your domain (or _crossly.<domain>) now.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.CreateDeveloperAppVerificationCheckDomain(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.CreateDeveloperAppVerificationCheckDomain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeveloperAppVerificationCheckDomain`: DeveloperAppVerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.CreateDeveloperAppVerificationCheckDomain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeveloperAppVerificationCheckDomainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeveloperAppVerificationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDeveloperAppVerificationSubmit

> DeveloperAppVerificationResponse CreateDeveloperAppVerificationSubmit(ctx, appId).Execute()

Send the verification to Crossly review (each scope is decided separately).

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.CreateDeveloperAppVerificationSubmit(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.CreateDeveloperAppVerificationSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeveloperAppVerificationSubmit`: DeveloperAppVerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.CreateDeveloperAppVerificationSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeveloperAppVerificationSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeveloperAppVerificationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDeveloperAppVerificationWithdraw

> DeveloperAppVerificationResponse CreateDeveloperAppVerificationWithdraw(ctx, appId).Execute()

Take the verification out of review to edit it.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.CreateDeveloperAppVerificationWithdraw(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.CreateDeveloperAppVerificationWithdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeveloperAppVerificationWithdraw`: DeveloperAppVerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.CreateDeveloperAppVerificationWithdraw`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeveloperAppVerificationWithdrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeveloperAppVerificationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetDeveloperAppCertification

> GetDeveloperAppCertificationResponse GetDeveloperAppCertification(ctx, appId).Execute()

The Crossly Carnival client certification application (null until you create it): status, client, evidence, the passing run and what still blocks submission.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.GetDeveloperAppCertification(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.GetDeveloperAppCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeveloperAppCertification`: GetDeveloperAppCertificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.GetDeveloperAppCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeveloperAppCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetDeveloperAppCertificationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetDeveloperAppClientVersion

> GetDeveloperAppClientVersionResponse GetDeveloperAppClientVersion(ctx, appId, versionId).Execute()

One client release and its re-certification.

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
    appId := "appId_example" // string | 
    versionId := "versionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.GetDeveloperAppClientVersion(context.Background(), appId, versionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.GetDeveloperAppClientVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeveloperAppClientVersion`: GetDeveloperAppClientVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.GetDeveloperAppClientVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**versionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeveloperAppClientVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetDeveloperAppClientVersionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetDeveloperAppConformanceRun

> GetDeveloperAppConformanceRunResponse GetDeveloperAppConformanceRun(ctx, appId, runId).Execute()

A run's step, instruction and every check with its verdict.

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
    appId := "appId_example" // string | 
    runId := "runId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.GetDeveloperAppConformanceRun(context.Background(), appId, runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.GetDeveloperAppConformanceRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeveloperAppConformanceRun`: GetDeveloperAppConformanceRunResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.GetDeveloperAppConformanceRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 
**runId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeveloperAppConformanceRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetDeveloperAppConformanceRunResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetDeveloperAppHealth

> GetDeveloperAppHealthResponse GetDeveloperAppHealth(ctx, appId).Execute()

Your app's API health: write-request counts by status class (24 h, 7 d) across every seller using it, top error codes with advice, a retry-storm flag, the last 50 failed requests (message and correlation id only on requests made with your own account), and from Atlas: how many people authorized the app and with which scopes (counts, never who) and the token endpoint's OAuth errors by code (e.g. invalid_grant) with the latest one.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.GetDeveloperAppHealth(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.GetDeveloperAppHealth``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeveloperAppHealth`: GetDeveloperAppHealthResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.GetDeveloperAppHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeveloperAppHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetDeveloperAppHealthResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetDeveloperAppProfile

> GetDeveloperAppProfileResponse GetDeveloperAppProfile(ctx, appId).Execute()

Your app's public directory profile: every editable field, the listing status, what still needs review, and `preview` — the profile exactly as /seller-apps or /buyer-apps will show it.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.GetDeveloperAppProfile(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.GetDeveloperAppProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeveloperAppProfile`: GetDeveloperAppProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.GetDeveloperAppProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeveloperAppProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetDeveloperAppProfileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetDeveloperAppReview

> GetDeveloperAppReviewResponse GetDeveloperAppReview(ctx, appId).Execute()

Your app's Crossly Carnival client certification, verification and review history in one read.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.GetDeveloperAppReview(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.GetDeveloperAppReview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeveloperAppReview`: GetDeveloperAppReviewResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.GetDeveloperAppReview`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeveloperAppReviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetDeveloperAppReviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetDeveloperAppVerification

> GetDeveloperAppVerificationResponse GetDeveloperAppVerification(ctx, appId).Execute()

The app verification for worlds:publish / worlds:monetize (null until created): domain proof, policies, justifications and per-scope decisions.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.GetDeveloperAppVerification(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.GetDeveloperAppVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeveloperAppVerification`: GetDeveloperAppVerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.GetDeveloperAppVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeveloperAppVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetDeveloperAppVerificationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListDeveloperAppClientVersions

> ListDeveloperAppClientVersionsResponse ListDeveloperAppClientVersions(ctx, appId).Execute()

Releases of your certified client, newest first: the re-certification status (awaiting_run, running, carried, ops_review, approved, rejected, failed, superseded), the rule that decided it and its checks.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.ListDeveloperAppClientVersions(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.ListDeveloperAppClientVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeveloperAppClientVersions`: ListDeveloperAppClientVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.ListDeveloperAppClientVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeveloperAppClientVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListDeveloperAppClientVersionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListDeveloperAppConformanceRuns

> ListDeveloperAppConformanceRunsResponse ListDeveloperAppConformanceRuns(ctx, appId).Execute()

Your app's conformance runs, newest first.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.ListDeveloperAppConformanceRuns(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.ListDeveloperAppConformanceRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeveloperAppConformanceRuns`: ListDeveloperAppConformanceRunsResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.ListDeveloperAppConformanceRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeveloperAppConformanceRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListDeveloperAppConformanceRunsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateDeveloperAppCertification

> DeveloperAppCertificationResponse UpdateDeveloperAppCertification(ctx, appId).Execute()

Create or edit the application: client name, platform, version, build and source URLs, description and evidence (report menu, chat, Policy UI). Changing the client clears the passing run.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.UpdateDeveloperAppCertification(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.UpdateDeveloperAppCertification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeveloperAppCertification`: DeveloperAppCertificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.UpdateDeveloperAppCertification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeveloperAppCertificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeveloperAppCertificationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateDeveloperAppListing

> UpdateDeveloperAppListingResponse UpdateDeveloperAppListing(ctx, appId).Execute()

Turn \"Listed publicly\" on or off. On: an app Crossly approved whose name, scopes and audience are unchanged is listed at once; anything else goes to review (pending_review). Off: unlisted immediately.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.UpdateDeveloperAppListing(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.UpdateDeveloperAppListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeveloperAppListing`: UpdateDeveloperAppListingResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.UpdateDeveloperAppListing`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeveloperAppListingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateDeveloperAppListingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateDeveloperAppProfile

> UpdateDeveloperAppProfileResponse UpdateDeveloperAppProfile(ctx, appId).Execute()

Edit your app's directory profile: name, slug, tagline, description, longDescription (markdown), category, audience (auto|seller|buyer|both), logoUrl, bannerUrl (4:1, 1600×400), screenshots (≤6), demoVideoUrl, homepageUrl, installUrl, support/privacy/terms links, docsLinks [{label,url}] and pricingNote. Absent keys are kept, null clears. Images not on Crossly's CDN are copied onto it. Renaming a listed app or changing its audience sends it back to review.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.UpdateDeveloperAppProfile(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.UpdateDeveloperAppProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeveloperAppProfile`: UpdateDeveloperAppProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.UpdateDeveloperAppProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeveloperAppProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateDeveloperAppProfileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateDeveloperAppVerification

> DeveloperAppVerificationResponse UpdateDeveloperAppVerification(ctx, appId).Execute()

Create or edit the verification: domain, requested scopes with a justification each, privacy policy, terms, support email, demo.

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
    appId := "appId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.DeveloperAppsApi.UpdateDeveloperAppVerification(context.Background(), appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeveloperAppsApi.UpdateDeveloperAppVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeveloperAppVerification`: DeveloperAppVerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `DeveloperAppsApi.UpdateDeveloperAppVerification`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeveloperAppVerificationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeveloperAppVerificationResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

