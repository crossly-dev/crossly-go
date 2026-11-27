# \WorldsTeamApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorldCollaborationAccept**](WorldsTeamApi.md#CreateWorldCollaborationAccept) | **Post** /v1/worlds/collaborations/{experienceId}/accept | Accept an invitation.
[**CreateWorldCollaborationDecline**](WorldsTeamApi.md#CreateWorldCollaborationDecline) | **Post** /v1/worlds/collaborations/{experienceId}/decline | Decline an invitation.
[**CreateWorldExperienceActivityRevert**](WorldsTeamApi.md#CreateWorldExperienceActivityRevert) | **Post** /v1/worlds/experiences/{experienceId}/activity/{activityId}/revert | Revert an instance change: its undo is applied as a new edit by you; anything changed since by others is left alone.
[**CreateWorldExperienceCollaborator**](WorldsTeamApi.md#CreateWorldExperienceCollaborator) | **Post** /v1/worlds/experiences/{experienceId}/collaborators | Invite a Crossly handle with a role. Owner or admin; only the owner makes admins. Counts against the owner&#39;s plan cap.
[**CreateWorldExperiencePlaceCollabOp**](WorldsTeamApi.md#CreateWorldExperiencePlaceCollabOp) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/collab/ops | Apply a batch of ops (insert, delete, move, prop, attr, name, tags, world) in server order. Each op is checked against your role; a retried batch id is answered, not re-applied.
[**CreateWorldExperiencePlaceCollabSave**](WorldsTeamApi.md#CreateWorldExperiencePlaceCollabSave) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/collab/save | Write the session into the place draft now (publishing and cloud tests do this themselves).
[**CreateWorldExperiencePlaceCollabScriptCommit**](WorldsTeamApi.md#CreateWorldExperiencePlaceCollabScriptCommit) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/collab/scripts/{nodeId}/commit | Commit your draft with a three-way merge. 409 merge_conflict carries the merged text with conflict blocks and the current source.
[**DeleteWorldExperienceCollaborator**](WorldsTeamApi.md#DeleteWorldExperienceCollaborator) | **Delete** /v1/worlds/experiences/{experienceId}/collaborators/{userId} | Remove a collaborator, or leave (your own user id). Their revenue share returns to the owner from next month.
[**DeleteWorldExperiencePlaceCollabScript**](WorldsTeamApi.md#DeleteWorldExperiencePlaceCollabScript) | **Delete** /v1/worlds/experiences/{experienceId}/places/{placeId}/collab/scripts/{nodeId} | Discard your draft of a script.
[**GetWorldExperiencePlaceCollab**](WorldsTeamApi.md#GetWorldExperiencePlaceCollab) | **Get** /v1/worlds/experiences/{experienceId}/places/{placeId}/collab | The live co-editing document of a place (node ids every editor shares), its op sequence and your role.
[**GetWorldExperiencePlaceCollabOp**](WorldsTeamApi.md#GetWorldExperiencePlaceCollabOp) | **Get** /v1/worlds/experiences/{experienceId}/places/{placeId}/collab/ops | Ops after a sequence number, grouped by batch. 409 resync when they are no longer kept: reload the session.
[**GetWorldExperienceSplit**](WorldsTeamApi.md#GetWorldExperienceSplit) | **Get** /v1/worlds/experiences/{experienceId}/splits | Revenue split of the creator share: current, scheduled, history.
[**GetWorldExperienceTeam**](WorldsTeamApi.md#GetWorldExperienceTeam) | **Get** /v1/worlds/experiences/{experienceId}/team | The team: owner, collaborators with roles (owner, admin, editor, scripter, builder, tester, analyst), your role and permissions, the plan cap.
[**ListWorldCollaborations**](WorldsTeamApi.md#ListWorldCollaborations) | **Get** /v1/worlds/collaborations | Experiences you are invited to or work on, with your role.
[**ListWorldExperienceActivity**](WorldsTeamApi.md#ListWorldExperienceActivity) | **Get** /v1/worlds/experiences/{experienceId}/activity | Activity history, newest first: who changed what, when, from where (studio, api, mcp, cli); instance changes are revertible.
[**ListWorldExperiencePlaceCollabScripts**](WorldsTeamApi.md#ListWorldExperiencePlaceCollabScripts) | **Get** /v1/worlds/experiences/{experienceId}/places/{placeId}/collab/scripts | Your script drafts in this place, with whether someone committed since you started.
[**UpdateWorldExperienceCollaborator**](WorldsTeamApi.md#UpdateWorldExperienceCollaborator) | **Patch** /v1/worlds/experiences/{experienceId}/collaborators/{userId} | Change a collaborator&#39;s role.
[**UpdateWorldExperiencePlaceCollabScript**](WorldsTeamApi.md#UpdateWorldExperiencePlaceCollabScript) | **Put** /v1/worlds/experiences/{experienceId}/places/{placeId}/collab/scripts/{nodeId} | Save your own draft of a script (needs edit scripts).
[**UpdateWorldExperienceSplit**](WorldsTeamApi.md#UpdateWorldExperienceSplit) | **Put** /v1/worlds/experiences/{experienceId}/splits | Owner only: split the creator share (70% of net) in basis points summing to 10000. Takes effect the 1st of next month; every party is notified. The 70% total never changes.



## CreateWorldCollaborationAccept

> CreateWorldCollaborationAcceptResponse CreateWorldCollaborationAccept(ctx, experienceId).Execute()

Accept an invitation.

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
    resp, r, err := apiClient.WorldsTeamApi.CreateWorldCollaborationAccept(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.CreateWorldCollaborationAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldCollaborationAccept`: CreateWorldCollaborationAcceptResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.CreateWorldCollaborationAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldCollaborationAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldCollaborationAcceptResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldCollaborationDecline

> CreateWorldCollaborationDeclineResponse CreateWorldCollaborationDecline(ctx, experienceId).Execute()

Decline an invitation.

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
    resp, r, err := apiClient.WorldsTeamApi.CreateWorldCollaborationDecline(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.CreateWorldCollaborationDecline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldCollaborationDecline`: CreateWorldCollaborationDeclineResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.CreateWorldCollaborationDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldCollaborationDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldCollaborationDeclineResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceActivityRevert

> CreateWorldExperienceActivityRevertResponse CreateWorldExperienceActivityRevert(ctx, experienceId, activityId).Execute()

Revert an instance change: its undo is applied as a new edit by you; anything changed since by others is left alone.

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
    activityId := "activityId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsTeamApi.CreateWorldExperienceActivityRevert(context.Background(), experienceId, activityId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.CreateWorldExperienceActivityRevert``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceActivityRevert`: CreateWorldExperienceActivityRevertResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.CreateWorldExperienceActivityRevert`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**activityId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceActivityRevertRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperienceActivityRevertResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperienceCollaborator

> CreateWorldExperienceCollaboratorResponse CreateWorldExperienceCollaborator(ctx, experienceId).Execute()

Invite a Crossly handle with a role. Owner or admin; only the owner makes admins. Counts against the owner's plan cap.

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
    resp, r, err := apiClient.WorldsTeamApi.CreateWorldExperienceCollaborator(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.CreateWorldExperienceCollaborator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceCollaborator`: CreateWorldExperienceCollaboratorResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.CreateWorldExperienceCollaborator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceCollaboratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceCollaboratorResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceCollabOp

> CreateWorldExperiencePlaceCollabOpResponse CreateWorldExperiencePlaceCollabOp(ctx, experienceId, placeId).Execute()

Apply a batch of ops (insert, delete, move, prop, attr, name, tags, world) in server order. Each op is checked against your role; a retried batch id is answered, not re-applied.

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
    resp, r, err := apiClient.WorldsTeamApi.CreateWorldExperiencePlaceCollabOp(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.CreateWorldExperiencePlaceCollabOp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceCollabOp`: CreateWorldExperiencePlaceCollabOpResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.CreateWorldExperiencePlaceCollabOp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceCollabOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperiencePlaceCollabOpResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceCollabSave

> CreateWorldExperiencePlaceCollabSaveResponse CreateWorldExperiencePlaceCollabSave(ctx, experienceId, placeId).Execute()

Write the session into the place draft now (publishing and cloud tests do this themselves).

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
    resp, r, err := apiClient.WorldsTeamApi.CreateWorldExperiencePlaceCollabSave(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.CreateWorldExperiencePlaceCollabSave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceCollabSave`: CreateWorldExperiencePlaceCollabSaveResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.CreateWorldExperiencePlaceCollabSave`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceCollabSaveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperiencePlaceCollabSaveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlaceCollabScriptCommit

> CreateWorldExperiencePlaceCollabScriptCommitResponse CreateWorldExperiencePlaceCollabScriptCommit(ctx, experienceId, placeId, nodeId).Execute()

Commit your draft with a three-way merge. 409 merge_conflict carries the merged text with conflict blocks and the current source.

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
    nodeId := "nodeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsTeamApi.CreateWorldExperiencePlaceCollabScriptCommit(context.Background(), experienceId, placeId, nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.CreateWorldExperiencePlaceCollabScriptCommit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlaceCollabScriptCommit`: CreateWorldExperiencePlaceCollabScriptCommitResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.CreateWorldExperiencePlaceCollabScriptCommit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlaceCollabScriptCommitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**CreateWorldExperiencePlaceCollabScriptCommitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorldExperienceCollaborator

> DeleteSpatialSceneGrantResponse DeleteWorldExperienceCollaborator(ctx, experienceId, userId).Execute()

Remove a collaborator, or leave (your own user id). Their revenue share returns to the owner from next month.

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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsTeamApi.DeleteWorldExperienceCollaborator(context.Background(), experienceId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.DeleteWorldExperienceCollaborator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldExperienceCollaborator`: DeleteSpatialSceneGrantResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.DeleteWorldExperienceCollaborator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldExperienceCollaboratorRequest struct via the builder pattern


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


## DeleteWorldExperiencePlaceCollabScript

> DeleteWorldExperiencePlaceCollabScriptResponse DeleteWorldExperiencePlaceCollabScript(ctx, experienceId, placeId, nodeId).Execute()

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
    nodeId := "nodeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsTeamApi.DeleteWorldExperiencePlaceCollabScript(context.Background(), experienceId, placeId, nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.DeleteWorldExperiencePlaceCollabScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldExperiencePlaceCollabScript`: DeleteWorldExperiencePlaceCollabScriptResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.DeleteWorldExperiencePlaceCollabScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldExperiencePlaceCollabScriptRequest struct via the builder pattern


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


## GetWorldExperiencePlaceCollab

> GetWorldExperiencePlaceCollabResponse GetWorldExperiencePlaceCollab(ctx, experienceId, placeId).Execute()

The live co-editing document of a place (node ids every editor shares), its op sequence and your role.

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
    resp, r, err := apiClient.WorldsTeamApi.GetWorldExperiencePlaceCollab(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.GetWorldExperiencePlaceCollab``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperiencePlaceCollab`: GetWorldExperiencePlaceCollabResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.GetWorldExperiencePlaceCollab`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperiencePlaceCollabRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperiencePlaceCollabResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperiencePlaceCollabOp

> GetWorldExperiencePlaceCollabOpResponse GetWorldExperiencePlaceCollabOp(ctx, experienceId, placeId).Execute()

Ops after a sequence number, grouped by batch. 409 resync when they are no longer kept: reload the session.

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
    resp, r, err := apiClient.WorldsTeamApi.GetWorldExperiencePlaceCollabOp(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.GetWorldExperiencePlaceCollabOp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperiencePlaceCollabOp`: GetWorldExperiencePlaceCollabOpResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.GetWorldExperiencePlaceCollabOp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperiencePlaceCollabOpRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldExperiencePlaceCollabOpResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceSplit

> GetWorldExperienceSplitResponse GetWorldExperienceSplit(ctx, experienceId).Execute()

Revenue split of the creator share: current, scheduled, history.

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
    resp, r, err := apiClient.WorldsTeamApi.GetWorldExperienceSplit(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.GetWorldExperienceSplit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceSplit`: GetWorldExperienceSplitResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.GetWorldExperienceSplit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceSplitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceSplitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceTeam

> GetWorldExperienceTeamResponse GetWorldExperienceTeam(ctx, experienceId).Execute()

The team: owner, collaborators with roles (owner, admin, editor, scripter, builder, tester, analyst), your role and permissions, the plan cap.

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
    resp, r, err := apiClient.WorldsTeamApi.GetWorldExperienceTeam(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.GetWorldExperienceTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceTeam`: GetWorldExperienceTeamResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.GetWorldExperienceTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceTeamResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldCollaborations

> ListWorldCollaborationsResponse ListWorldCollaborations(ctx).Execute()

Experiences you are invited to or work on, with your role.

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
    resp, r, err := apiClient.WorldsTeamApi.ListWorldCollaborations(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.ListWorldCollaborations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldCollaborations`: ListWorldCollaborationsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.ListWorldCollaborations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldCollaborationsRequest struct via the builder pattern


### Return type

**ListWorldCollaborationsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperienceActivity

> ListWorldExperienceActivityResponse ListWorldExperienceActivity(ctx, experienceId).Execute()

Activity history, newest first: who changed what, when, from where (studio, api, mcp, cli); instance changes are revertible.

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
    resp, r, err := apiClient.WorldsTeamApi.ListWorldExperienceActivity(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.ListWorldExperienceActivity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperienceActivity`: ListWorldExperienceActivityResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.ListWorldExperienceActivity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperienceActivityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorldExperienceActivityResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperiencePlaceCollabScripts

> ListWorldExperiencePlaceCollabScriptsResponse ListWorldExperiencePlaceCollabScripts(ctx, experienceId, placeId).Execute()

Your script drafts in this place, with whether someone committed since you started.

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
    resp, r, err := apiClient.WorldsTeamApi.ListWorldExperiencePlaceCollabScripts(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.ListWorldExperiencePlaceCollabScripts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperiencePlaceCollabScripts`: ListWorldExperiencePlaceCollabScriptsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.ListWorldExperiencePlaceCollabScripts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperiencePlaceCollabScriptsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**ListWorldExperiencePlaceCollabScriptsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceCollaborator

> CreateWorldExperienceCollaboratorResponse UpdateWorldExperienceCollaborator(ctx, experienceId, userId).Execute()

Change a collaborator's role.

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
    userId := "userId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsTeamApi.UpdateWorldExperienceCollaborator(context.Background(), experienceId, userId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.UpdateWorldExperienceCollaborator``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceCollaborator`: CreateWorldExperienceCollaboratorResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.UpdateWorldExperienceCollaborator`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**userId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceCollaboratorRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperienceCollaboratorResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperiencePlaceCollabScript

> UpdateWorldExperiencePlaceCollabScriptResponse UpdateWorldExperiencePlaceCollabScript(ctx, experienceId, placeId, nodeId).Execute()

Save your own draft of a script (needs edit scripts).

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
    nodeId := "nodeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsTeamApi.UpdateWorldExperiencePlaceCollabScript(context.Background(), experienceId, placeId, nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.UpdateWorldExperiencePlaceCollabScript``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperiencePlaceCollabScript`: UpdateWorldExperiencePlaceCollabScriptResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.UpdateWorldExperiencePlaceCollabScript`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperiencePlaceCollabScriptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**UpdateWorldExperiencePlaceCollabScriptResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceSplit

> GetWorldExperienceSplitResponse UpdateWorldExperienceSplit(ctx, experienceId).Execute()

Owner only: split the creator share (70% of net) in basis points summing to 10000. Takes effect the 1st of next month; every party is notified. The 70% total never changes.

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
    resp, r, err := apiClient.WorldsTeamApi.UpdateWorldExperienceSplit(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsTeamApi.UpdateWorldExperienceSplit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceSplit`: GetWorldExperienceSplitResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsTeamApi.UpdateWorldExperienceSplit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceSplitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceSplitResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

