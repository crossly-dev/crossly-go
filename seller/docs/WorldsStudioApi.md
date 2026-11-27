# \WorldsStudioApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorldExperienceAssetMesh**](WorldsStudioApi.md#CreateWorldExperienceAssetMesh) | **Post** /v1/worlds/experiences/{experienceId}/assets/meshes | Import a .glb (base64) as a MeshPart: safety parse + budgets, stored on the Crossly CDN, a phone LOD over 5k triangles, collision Box or convex Hull. Returns the asset and the MeshPart node; with placeId it is inserted into that place through its co-editing session.
[**CreateWorldExperienceAssetMeshCheck**](WorldsStudioApi.md#CreateWorldExperienceAssetMeshCheck) | **Post** /v1/worlds/experiences/{experienceId}/assets/meshes/check | Check a .glb (base64) for import without storing it: the budget report per device, plain-English errors with a fix list, whether a phone LOD will be made, mesh bounds and hull points.
[**CreateWorldExperiencePlaceAssistant**](WorldsStudioApi.md#CreateWorldExperiencePlaceAssistant) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/assistant | Ask the Studio Assistant (mode explain | edit | build): returns an explanation and a proposed change set (scripts, inserts, edits) checked against your role and linted. Nothing is applied.
[**CreateWorldExperiencePlaceAssistantApply**](WorldsStudioApi.md#CreateWorldExperiencePlaceAssistantApply) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/assistant/apply | Apply accepted Assistant changes to the place as one co-editing batch (activity via assistant, revertible). Re-checked against the place as it is now.
[**CreateWorldExperiencePlaceStudioEdit**](WorldsStudioApi.md#CreateWorldExperiencePlaceStudioEdit) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/edits | Apply set / insert / move / delete steps by path as ONE batch through Team Create (live in every open Studio, one revertible activity entry). All-or-nothing: any bad step refuses the batch with every problem listed.
[**CreateWorldExperiencePlaceStudioInsertAsset**](WorldsStudioApi.md#CreateWorldExperiencePlaceStudioInsertAsset) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/insert-asset | Insert a Toolbox asset under a path (default Workspace; ServerScripts for scripts). Paid assets need the experience owner&#39;s licence.
[**CreateWorldExperiencePlaceStudioPlaytest**](WorldsStudioApi.md#CreateWorldExperiencePlaceStudioPlaytest) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/playtest | Cloud test: flush the session, freeze a private test version and start a fleet server for it. 201 ready (ticket) or 202 starting; outputCursor reads only this test&#39;s Output.
[**CreateWorldExperiencePlaceStudioPlaytestStop**](WorldsStudioApi.md#CreateWorldExperiencePlaceStudioPlaytestStop) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/playtest/stop | Stop this place&#39;s cloud-test servers now.
[**CreateWorldExperiencePlaceStudioPublish**](WorldsStudioApi.md#CreateWorldExperiencePlaceStudioPublish) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/publish | Prepare a publish from the live session: checks, details, rating answers, access. Never submits and never acknowledges the money review — the person does (or submit_worlds_publish).
[**CreateWorldExperiencePlaceStudioScriptCommit**](WorldsStudioApi.md#CreateWorldExperiencePlaceStudioScriptCommit) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/script/commit | Commit your draft with a three-way merge onto the script as it is now (collaborators&#39; edits kept). 409 merge_conflict carries merged text with conflict blocks and theirs.
[**DeleteWorldExperiencePlaceStudioScript**](WorldsStudioApi.md#DeleteWorldExperiencePlaceStudioScript) | **Delete** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/script | Discard your draft of a script.
[**GetWorldAssistantStatus**](WorldsStudioApi.md#GetWorldAssistantStatus) | **Get** /v1/worlds/assistant/status | Studio Assistant status: whether Crossly&#39;s model key is configured, whether you have your own AI key, and your plan allowance (asks left today, budget left this month).
[**GetWorldExperiencePlaceStudioFind**](WorldsStudioApi.md#GetWorldExperiencePlaceStudioFind) | **Get** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/find | Find instances by name (contains), exact class and/or tag.
[**GetWorldExperiencePlaceStudioInstance**](WorldsStudioApi.md#GetWorldExperiencePlaceStudioInstance) | **Get** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/instance | One instance by path (or session id): class, properties, attributes, tags, children. Script source is read with …/studio/script.
[**GetWorldExperiencePlaceStudioOutput**](WorldsStudioApi.md#GetWorldExperiencePlaceStudioOutput) | **Get** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/output | Server Output (print, warnings, errors) after a cursor: every line of cloud tests, errors of live servers (scope&#x3D;all). Kept 24 hours.
[**GetWorldExperiencePlaceStudioProblem**](WorldsStudioApi.md#GetWorldExperiencePlaceStudioProblem) | **Get** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/problems | The publish checks on the place as the live session holds it: script lint, budgets per device, settings, shop listings, Toolbox licences. publishable &#x3D; no errors.
[**GetWorldExperiencePlaceStudioScript**](WorldsStudioApi.md#GetWorldExperiencePlaceStudioScript) | **Get** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/script | A script&#39;s current source, your draft of it (if any, with whether someone committed since) and the lint of each.
[**GetWorldExperiencePlaceStudioSnapshot**](WorldsStudioApi.md#GetWorldExperiencePlaceStudioSnapshot) | **Get** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/snapshot | A plan (top) or front drawing of the parts under a path as SVG, with bounds, spawns and unanchored parts that will fall.
[**GetWorldExperiencePlaceStudioTree**](WorldsStudioApi.md#GetWorldExperiencePlaceStudioTree) | **Get** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/tree | The live instance tree (what every open Studio shows) by dotted path, depth levels down (default 3, max 12).
[**GetWorldStudioClass**](WorldsStudioApi.md#GetWorldStudioClass) | **Get** /v1/worlds/studio/classes | The engine&#39;s class table: insertable classes, services, and for ?name&#x3D; one class with every property, its type, default and the exact JSON a place (and a studio edit) takes.
[**ListWorldExperienceAssets**](WorldsStudioApi.md#ListWorldExperienceAssets) | **Get** /v1/worlds/experiences/{experienceId}/assets | Meshes imported into this experience, newest first: paths, phone LOD, bounds, hull size, budget report, moderation (pending until a reviewed publish uses them) and a URL you can load now.
[**UpdateWorldExperiencePlaceStudioScript**](WorldsStudioApi.md#UpdateWorldExperiencePlaceStudioScript) | **Put** /v1/worlds/experiences/{experienceId}/places/{placeId}/studio/script | Save your draft of a script (the place does not change until you commit). Returns lint and the diff from the current source.



## CreateWorldExperienceAssetMesh

> CreateWorldExperienceAssetMeshResponse CreateWorldExperienceAssetMesh(ctx, experienceId).Execute()

Import a .glb (base64) as a MeshPart: safety parse + budgets, stored on the Crossly CDN, a phone LOD over 5k triangles, collision Box or convex Hull. Returns the asset and the MeshPart node; with placeId it is inserted into that place through its co-editing session.

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
    resp, r, err := apiClient.WorldsStudioApi.CreateWorldExperienceAssetMesh(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.CreateWorldExperienceAssetMesh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceAssetMesh`: CreateWorldExperienceAssetMeshResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.CreateWorldExperienceAssetMesh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceAssetMeshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceAssetMeshResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceAssetMeshCheck

> CreateWorldExperienceAssetMeshCheckResponse CreateWorldExperienceAssetMeshCheck(ctx, experienceId).Execute()

Check a .glb (base64) for import without storing it: the budget report per device, plain-English errors with a fix list, whether a phone LOD will be made, mesh bounds and hull points.

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
    resp, r, err := apiClient.WorldsStudioApi.CreateWorldExperienceAssetMeshCheck(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.CreateWorldExperienceAssetMeshCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceAssetMeshCheck`: CreateWorldExperienceAssetMeshCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.CreateWorldExperienceAssetMeshCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceAssetMeshCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceAssetMeshCheckResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceAssistant

> CreateWorldExperiencePlaceAssistantResponse CreateWorldExperiencePlaceAssistant(ctx, experienceId, placeId).Execute()

Ask the Studio Assistant (mode explain | edit | build): returns an explanation and a proposed change set (scripts, inserts, edits) checked against your role and linted. Nothing is applied.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.CreateWorldExperiencePlaceAssistant(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.CreateWorldExperiencePlaceAssistant``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceAssistant`: CreateWorldExperiencePlaceAssistantResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.CreateWorldExperiencePlaceAssistant`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceAssistantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperiencePlaceAssistantResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceAssistantApply

> CreateWorldExperiencePlaceAssistantApplyResponse CreateWorldExperiencePlaceAssistantApply(ctx, experienceId, placeId).Execute()

Apply accepted Assistant changes to the place as one co-editing batch (activity via assistant, revertible). Re-checked against the place as it is now.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.CreateWorldExperiencePlaceAssistantApply(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.CreateWorldExperiencePlaceAssistantApply``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceAssistantApply`: CreateWorldExperiencePlaceAssistantApplyResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.CreateWorldExperiencePlaceAssistantApply`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceAssistantApplyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperiencePlaceAssistantApplyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceStudioEdit

> CreateWorldExperiencePlaceStudioEditResponse CreateWorldExperiencePlaceStudioEdit(ctx, experienceId, placeId).Execute()

Apply set / insert / move / delete steps by path as ONE batch through Team Create (live in every open Studio, one revertible activity entry). All-or-nothing: any bad step refuses the batch with every problem listed.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.CreateWorldExperiencePlaceStudioEdit(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.CreateWorldExperiencePlaceStudioEdit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceStudioEdit`: CreateWorldExperiencePlaceStudioEditResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.CreateWorldExperiencePlaceStudioEdit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceStudioEditRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperiencePlaceStudioEditResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceStudioInsertAsset

> CreateWorldExperiencePlaceStudioInsertAssetResponse CreateWorldExperiencePlaceStudioInsertAsset(ctx, experienceId, placeId).Execute()

Insert a Toolbox asset under a path (default Workspace; ServerScripts for scripts). Paid assets need the experience owner's licence.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.CreateWorldExperiencePlaceStudioInsertAsset(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.CreateWorldExperiencePlaceStudioInsertAsset``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceStudioInsertAsset`: CreateWorldExperiencePlaceStudioInsertAssetResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.CreateWorldExperiencePlaceStudioInsertAsset`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceStudioInsertAssetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperiencePlaceStudioInsertAssetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceStudioPlaytest

> CreateWorldExperiencePlaceStudioPlaytestResponse CreateWorldExperiencePlaceStudioPlaytest(ctx, experienceId, placeId).Execute()

Cloud test: flush the session, freeze a private test version and start a fleet server for it. 201 ready (ticket) or 202 starting; outputCursor reads only this test's Output.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.CreateWorldExperiencePlaceStudioPlaytest(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.CreateWorldExperiencePlaceStudioPlaytest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceStudioPlaytest`: CreateWorldExperiencePlaceStudioPlaytestResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.CreateWorldExperiencePlaceStudioPlaytest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceStudioPlaytestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperiencePlaceStudioPlaytestResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceStudioPlaytestStop

> CreateWorldExperiencePlaceStudioPlaytestStopResponse CreateWorldExperiencePlaceStudioPlaytestStop(ctx, experienceId, placeId).Execute()

Stop this place's cloud-test servers now.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.CreateWorldExperiencePlaceStudioPlaytestStop(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.CreateWorldExperiencePlaceStudioPlaytestStop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceStudioPlaytestStop`: CreateWorldExperiencePlaceStudioPlaytestStopResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.CreateWorldExperiencePlaceStudioPlaytestStop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceStudioPlaytestStopRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperiencePlaceStudioPlaytestStopResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceStudioPublish

> CreateWorldExperiencePlaceStudioPublishResponse CreateWorldExperiencePlaceStudioPublish(ctx, experienceId, placeId).Execute()

Prepare a publish from the live session: checks, details, rating answers, access. Never submits and never acknowledges the money review — the person does (or submit_worlds_publish).

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.CreateWorldExperiencePlaceStudioPublish(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.CreateWorldExperiencePlaceStudioPublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceStudioPublish`: CreateWorldExperiencePlaceStudioPublishResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.CreateWorldExperiencePlaceStudioPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceStudioPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperiencePlaceStudioPublishResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceStudioScriptCommit

> CreateWorldExperiencePlaceStudioScriptCommitResponse CreateWorldExperiencePlaceStudioScriptCommit(ctx, experienceId, placeId).Execute()

Commit your draft with a three-way merge onto the script as it is now (collaborators' edits kept). 409 merge_conflict carries merged text with conflict blocks and theirs.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.CreateWorldExperiencePlaceStudioScriptCommit(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.CreateWorldExperiencePlaceStudioScriptCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceStudioScriptCommit`: CreateWorldExperiencePlaceStudioScriptCommitResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.CreateWorldExperiencePlaceStudioScriptCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceStudioScriptCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperiencePlaceStudioScriptCommitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorldExperiencePlaceStudioScript

> DeleteWorldExperiencePlaceCollabScriptResponse DeleteWorldExperiencePlaceStudioScript(ctx, experienceId, placeId).Execute()

Discard your draft of a script.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.DeleteWorldExperiencePlaceStudioScript(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.DeleteWorldExperiencePlaceStudioScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldExperiencePlaceStudioScript`: DeleteWorldExperiencePlaceCollabScriptResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.DeleteWorldExperiencePlaceStudioScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldExperiencePlaceStudioScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**DeleteWorldExperiencePlaceCollabScriptResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldAssistantStatus

> GetWorldAssistantStatusResponse GetWorldAssistantStatus(ctx).Execute()

Studio Assistant status: whether Crossly's model key is configured, whether you have your own AI key, and your plan allowance (asks left today, budget left this month).

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
    resp, r, err := apiClient.WorldsStudioApi.GetWorldAssistantStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.GetWorldAssistantStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldAssistantStatus`: GetWorldAssistantStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.GetWorldAssistantStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldAssistantStatusRequest struct via the builder pattern


### Return type

**GetWorldAssistantStatusResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperiencePlaceStudioFind

> GetWorldExperiencePlaceStudioFindResponse GetWorldExperiencePlaceStudioFind(ctx, experienceId, placeId).Execute()

Find instances by name (contains), exact class and/or tag.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.GetWorldExperiencePlaceStudioFind(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.GetWorldExperiencePlaceStudioFind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperiencePlaceStudioFind`: GetWorldExperiencePlaceStudioFindResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.GetWorldExperiencePlaceStudioFind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperiencePlaceStudioFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperiencePlaceStudioFindResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperiencePlaceStudioInstance

> GetWorldExperiencePlaceStudioInstanceResponse GetWorldExperiencePlaceStudioInstance(ctx, experienceId, placeId).Execute()

One instance by path (or session id): class, properties, attributes, tags, children. Script source is read with …/studio/script.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.GetWorldExperiencePlaceStudioInstance(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.GetWorldExperiencePlaceStudioInstance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperiencePlaceStudioInstance`: GetWorldExperiencePlaceStudioInstanceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.GetWorldExperiencePlaceStudioInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperiencePlaceStudioInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperiencePlaceStudioInstanceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperiencePlaceStudioOutput

> GetWorldExperiencePlaceStudioOutputResponse GetWorldExperiencePlaceStudioOutput(ctx, experienceId, placeId).Execute()

Server Output (print, warnings, errors) after a cursor: every line of cloud tests, errors of live servers (scope=all). Kept 24 hours.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.GetWorldExperiencePlaceStudioOutput(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.GetWorldExperiencePlaceStudioOutput``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperiencePlaceStudioOutput`: GetWorldExperiencePlaceStudioOutputResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.GetWorldExperiencePlaceStudioOutput`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperiencePlaceStudioOutputRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperiencePlaceStudioOutputResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperiencePlaceStudioProblem

> GetWorldExperiencePlaceStudioProblemResponse GetWorldExperiencePlaceStudioProblem(ctx, experienceId, placeId).Execute()

The publish checks on the place as the live session holds it: script lint, budgets per device, settings, shop listings, Toolbox licences. publishable = no errors.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.GetWorldExperiencePlaceStudioProblem(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.GetWorldExperiencePlaceStudioProblem``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperiencePlaceStudioProblem`: GetWorldExperiencePlaceStudioProblemResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.GetWorldExperiencePlaceStudioProblem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperiencePlaceStudioProblemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperiencePlaceStudioProblemResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperiencePlaceStudioScript

> GetWorldExperiencePlaceStudioScriptResponse GetWorldExperiencePlaceStudioScript(ctx, experienceId, placeId).Execute()

A script's current source, your draft of it (if any, with whether someone committed since) and the lint of each.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.GetWorldExperiencePlaceStudioScript(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.GetWorldExperiencePlaceStudioScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperiencePlaceStudioScript`: GetWorldExperiencePlaceStudioScriptResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.GetWorldExperiencePlaceStudioScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperiencePlaceStudioScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperiencePlaceStudioScriptResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperiencePlaceStudioSnapshot

> GetWorldExperiencePlaceStudioSnapshotResponse GetWorldExperiencePlaceStudioSnapshot(ctx, experienceId, placeId).Execute()

A plan (top) or front drawing of the parts under a path as SVG, with bounds, spawns and unanchored parts that will fall.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.GetWorldExperiencePlaceStudioSnapshot(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.GetWorldExperiencePlaceStudioSnapshot``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperiencePlaceStudioSnapshot`: GetWorldExperiencePlaceStudioSnapshotResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.GetWorldExperiencePlaceStudioSnapshot`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperiencePlaceStudioSnapshotRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperiencePlaceStudioSnapshotResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperiencePlaceStudioTree

> GetWorldExperiencePlaceStudioTreeResponse GetWorldExperiencePlaceStudioTree(ctx, experienceId, placeId).Execute()

The live instance tree (what every open Studio shows) by dotted path, depth levels down (default 3, max 12).

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.GetWorldExperiencePlaceStudioTree(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.GetWorldExperiencePlaceStudioTree``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperiencePlaceStudioTree`: GetWorldExperiencePlaceStudioTreeResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.GetWorldExperiencePlaceStudioTree`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperiencePlaceStudioTreeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperiencePlaceStudioTreeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldStudioClass

> GetWorldStudioClassResponse GetWorldStudioClass(ctx).Execute()

The engine's class table: insertable classes, services, and for ?name= one class with every property, its type, default and the exact JSON a place (and a studio edit) takes.

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
    resp, r, err := apiClient.WorldsStudioApi.GetWorldStudioClass(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.GetWorldStudioClass``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldStudioClass`: GetWorldStudioClassResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.GetWorldStudioClass`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldStudioClassRequest struct via the builder pattern


### Return type

**GetWorldStudioClassResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceAssets

> ListWorldExperienceAssetsResponse ListWorldExperienceAssets(ctx, experienceId).Execute()

Meshes imported into this experience, newest first: paths, phone LOD, bounds, hull size, budget report, moderation (pending until a reviewed publish uses them) and a URL you can load now.

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
    resp, r, err := apiClient.WorldsStudioApi.ListWorldExperienceAssets(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.ListWorldExperienceAssets``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceAssets`: ListWorldExperienceAssetsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.ListWorldExperienceAssets`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceAssetsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceAssetsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperiencePlaceStudioScript

> UpdateWorldExperiencePlaceStudioScriptResponse UpdateWorldExperiencePlaceStudioScript(ctx, experienceId, placeId).Execute()

Save your draft of a script (the place does not change until you commit). Returns lint and the diff from the current source.

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
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsStudioApi.UpdateWorldExperiencePlaceStudioScript(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsStudioApi.UpdateWorldExperiencePlaceStudioScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperiencePlaceStudioScript`: UpdateWorldExperiencePlaceStudioScriptResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsStudioApi.UpdateWorldExperiencePlaceStudioScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperiencePlaceStudioScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**UpdateWorldExperiencePlaceStudioScriptResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

