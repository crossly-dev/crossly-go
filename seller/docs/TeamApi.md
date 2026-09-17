# \TeamApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTeamAccept**](TeamApi.md#CreateTeamAccept) | **Post** /v1/team/accept | Accept a pending team invitation by raw token.
[**CreateTeamInvite**](TeamApi.md#CreateTeamInvite) | **Post** /v1/team/invite | Mint a team invitation; returns the one-time accept URL.
[**CreateTeamLeave**](TeamApi.md#CreateTeamLeave) | **Post** /v1/team/leave | Leave every team this user is currently a member of.
[**CreateTeamRevoke**](TeamApi.md#CreateTeamRevoke) | **Post** /v1/team/revoke | Revoke a pending invite OR an active team member.
[**GetTeam**](TeamApi.md#GetTeam) | **Get** /v1/team | List pending team invitations + active members.
[**UpdateTeam**](TeamApi.md#UpdateTeam) | **Patch** /v1/team/{memberId} | Update a team member&#39;s scopes (owner only).



## CreateTeamAccept

> CreateTeamAcceptResponse CreateTeamAccept(ctx).Execute()

Accept a pending team invitation by raw token.

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
    resp, r, err := apiClient.TeamApi.CreateTeamAccept(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.CreateTeamAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTeamAccept`: CreateTeamAcceptResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamApi.CreateTeamAccept`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamAcceptRequest struct via the builder pattern


### Return type

[**CreateTeamAcceptResponse**](CreateTeamAcceptResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTeamInvite

> CreateTeamInviteResponse CreateTeamInvite(ctx).Execute()

Mint a team invitation; returns the one-time accept URL.

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
    resp, r, err := apiClient.TeamApi.CreateTeamInvite(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.CreateTeamInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTeamInvite`: CreateTeamInviteResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamApi.CreateTeamInvite`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamInviteRequest struct via the builder pattern


### Return type

[**CreateTeamInviteResponse**](CreateTeamInviteResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTeamLeave

> CreateTeamLeaveResponse CreateTeamLeave(ctx).Execute()

Leave every team this user is currently a member of.

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
    resp, r, err := apiClient.TeamApi.CreateTeamLeave(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.CreateTeamLeave``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTeamLeave`: CreateTeamLeaveResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamApi.CreateTeamLeave`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamLeaveRequest struct via the builder pattern


### Return type

[**CreateTeamLeaveResponse**](CreateTeamLeaveResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTeamRevoke

> CreateTeamRevokeResponse CreateTeamRevoke(ctx).Execute()

Revoke a pending invite OR an active team member.

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
    resp, r, err := apiClient.TeamApi.CreateTeamRevoke(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.CreateTeamRevoke``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTeamRevoke`: CreateTeamRevokeResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamApi.CreateTeamRevoke`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTeamRevokeRequest struct via the builder pattern


### Return type

[**CreateTeamRevokeResponse**](CreateTeamRevokeResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTeam

> GetTeamResponse GetTeam(ctx).Execute()

List pending team invitations + active members.

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
    resp, r, err := apiClient.TeamApi.GetTeam(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.GetTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTeam`: GetTeamResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamApi.GetTeam`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTeamRequest struct via the builder pattern


### Return type

[**GetTeamResponse**](GetTeamResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateTeam

> UpdateTeamResponse UpdateTeam(ctx, memberId).Execute()

Update a team member's scopes (owner only).

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
    memberId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TeamApi.UpdateTeam(context.Background(), memberId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TeamApi.UpdateTeam``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateTeam`: UpdateTeamResponse
    fmt.Fprintf(os.Stdout, "Response from `TeamApi.UpdateTeam`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**memberId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTeamRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UpdateTeamResponse**](UpdateTeamResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

