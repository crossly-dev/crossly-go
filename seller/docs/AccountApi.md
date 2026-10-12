# \AccountApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountCancelDeletion**](AccountApi.md#CreateAccountCancelDeletion) | **Post** /v1/account/cancel-deletion | Cancel a pending account deletion.
[**CreateAccountLogoutAll**](AccountApi.md#CreateAccountLogoutAll) | **Post** /v1/account/logout-all | Revoke every browser auth session for this user.
[**CreateAccountRequestDeletion**](AccountApi.md#CreateAccountRequestDeletion) | **Post** /v1/account/request-deletion | Schedule account deletion after a grace period.
[**DeleteAuthSession**](AccountApi.md#DeleteAuthSession) | **Delete** /v1/auth/sessions | Revoke all active browser sessions.
[**DeleteAuthSessionBySessionId**](AccountApi.md#DeleteAuthSessionBySessionId) | **Delete** /v1/auth/sessions/{sessionId} | Revoke a single browser session by id.
[**DeleteConnectedApp**](AccountApi.md#DeleteConnectedApp) | **Delete** /v1/connected-apps/{grantId} | Disconnect a third-party app. Its tokens stop working immediately.
[**GetAccountDeletionStatus**](AccountApi.md#GetAccountDeletionStatus) | **Get** /v1/account/deletion-status | Get the currently-pending deletion request, if any.
[**GetMe**](AccountApi.md#GetMe) | **Get** /v1/me | Identity check — authenticated user + PAT scopes + account state.
[**GetMembership**](AccountApi.md#GetMembership) | **Get** /v1/membership | The caller&#39;s Crossly One membership status and benefits.
[**ListAuthSessions**](AccountApi.md#ListAuthSessions) | **Get** /v1/auth/sessions | List active browser auth sessions.
[**ListConnectedApps**](AccountApi.md#ListConnectedApps) | **Get** /v1/connected-apps | List third-party OAuth apps with access to this account.



## CreateAccountCancelDeletion

> DeleteDropResponse CreateAccountCancelDeletion(ctx).Execute()

Cancel a pending account deletion.

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
    resp, r, err := apiClient.AccountApi.CreateAccountCancelDeletion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CreateAccountCancelDeletion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountCancelDeletion`: DeleteDropResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CreateAccountCancelDeletion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountCancelDeletionRequest struct via the builder pattern


### Return type

**DeleteDropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateAccountLogoutAll

> CreateAccountLogoutAllResponse CreateAccountLogoutAll(ctx).Execute()

Revoke every browser auth session for this user.

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
    resp, r, err := apiClient.AccountApi.CreateAccountLogoutAll(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CreateAccountLogoutAll``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountLogoutAll`: CreateAccountLogoutAllResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CreateAccountLogoutAll`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountLogoutAllRequest struct via the builder pattern


### Return type

**CreateAccountLogoutAllResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateAccountRequestDeletion

> CreateAccountRequestDeletionResponse CreateAccountRequestDeletion(ctx).Execute()

Schedule account deletion after a grace period.

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
    resp, r, err := apiClient.AccountApi.CreateAccountRequestDeletion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CreateAccountRequestDeletion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountRequestDeletion`: CreateAccountRequestDeletionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CreateAccountRequestDeletion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountRequestDeletionRequest struct via the builder pattern


### Return type

**CreateAccountRequestDeletionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteAuthSession

> DeleteAuthSessionResponse DeleteAuthSession(ctx).Execute()

Revoke all active browser sessions.

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
    resp, r, err := apiClient.AccountApi.DeleteAuthSession(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.DeleteAuthSession``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAuthSession`: DeleteAuthSessionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.DeleteAuthSession`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthSessionRequest struct via the builder pattern


### Return type

**DeleteAuthSessionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteAuthSessionBySessionId

> SuccessResponse DeleteAuthSessionBySessionId(ctx, sessionId).Execute()

Revoke a single browser session by id.

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
    sessionId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.DeleteAuthSessionBySessionId(context.Background(), sessionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.DeleteAuthSessionBySessionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAuthSessionBySessionId`: SuccessResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.DeleteAuthSessionBySessionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sessionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAuthSessionBySessionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**SuccessResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteConnectedApp

> OkResponse DeleteConnectedApp(ctx, grantId).Execute()

Disconnect a third-party app. Its tokens stop working immediately.

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
    grantId := "grantId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.DeleteConnectedApp(context.Background(), grantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.DeleteConnectedApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteConnectedApp`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.DeleteConnectedApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**grantId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteConnectedAppRequest struct via the builder pattern


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


## GetAccountDeletionStatus

> GetAccountDeletionStatusResponse GetAccountDeletionStatus(ctx).Execute()

Get the currently-pending deletion request, if any.

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
    resp, r, err := apiClient.AccountApi.GetAccountDeletionStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetAccountDeletionStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountDeletionStatus`: GetAccountDeletionStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetAccountDeletionStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountDeletionStatusRequest struct via the builder pattern


### Return type

**GetAccountDeletionStatusResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetMe

> GetMeResponse GetMe(ctx).Execute()

Identity check — authenticated user + PAT scopes + account state.

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
    resp, r, err := apiClient.AccountApi.GetMe(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMe`: GetMeResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetMe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMeRequest struct via the builder pattern


### Return type

**GetMeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetMembership

> GetMembershipResponse GetMembership(ctx).Execute()

The caller's Crossly One membership status and benefits.



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
    resp, r, err := apiClient.AccountApi.GetMembership(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetMembership``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMembership`: GetMembershipResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetMembership`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetMembershipRequest struct via the builder pattern


### Return type

**GetMembershipResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListAuthSessions

> ListAuthSessionsResponse ListAuthSessions(ctx).Execute()

List active browser auth sessions.

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
    resp, r, err := apiClient.AccountApi.ListAuthSessions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.ListAuthSessions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAuthSessions`: ListAuthSessionsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.ListAuthSessions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAuthSessionsRequest struct via the builder pattern


### Return type

**ListAuthSessionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListConnectedApps

> ListConnectedAppsResponse ListConnectedApps(ctx).Execute()

List third-party OAuth apps with access to this account.

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
    resp, r, err := apiClient.AccountApi.ListConnectedApps(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.ListConnectedApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectedApps`: ListConnectedAppsResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.ListConnectedApps`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListConnectedAppsRequest struct via the builder pattern


### Return type

**ListConnectedAppsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

