# \AccountApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccountCancelDeletion**](AccountApi.md#CreateAccountCancelDeletion) | **Post** /v1/account/cancel-deletion | Cancel a pending account deletion.
[**CreateAccountIdentityVerificationStart**](AccountApi.md#CreateAccountIdentityVerificationStart) | **Post** /v1/account/identity-verification/start | Start the identity check (photo ID + selfie). Returns a Stripe-hosted URL the person completes.
[**CreateAccountLogoutAll**](AccountApi.md#CreateAccountLogoutAll) | **Post** /v1/account/logout-all | Revoke every browser auth session for this user.
[**CreateAccountReconfirmDeletion**](AccountApi.md#CreateAccountReconfirmDeletion) | **Post** /v1/account/reconfirm-deletion | Confirm again a deletion paused because the account was in use.
[**CreateAccountRequestDeletion**](AccountApi.md#CreateAccountRequestDeletion) | **Post** /v1/account/request-deletion | Schedule account deletion after a grace period. Requires confirmEmail.
[**DeleteAuthSession**](AccountApi.md#DeleteAuthSession) | **Delete** /v1/auth/sessions | Revoke all active browser sessions.
[**DeleteAuthSessionBySessionId**](AccountApi.md#DeleteAuthSessionBySessionId) | **Delete** /v1/auth/sessions/{sessionId} | Revoke a single browser session by id.
[**DeleteConnectedApp**](AccountApi.md#DeleteConnectedApp) | **Delete** /v1/connected-apps/{grantId} | Disconnect a third-party app. Its tokens stop working immediately.
[**GetAccountClosureChecklist**](AccountApi.md#GetAccountClosureChecklist) | **Get** /v1/account/closure-checklist | What closing the account will and will not do, before you request it.
[**GetAccountDeletionStatus**](AccountApi.md#GetAccountDeletionStatus) | **Get** /v1/account/deletion-status | Get the currently-pending deletion request, if any.
[**GetAccountIdentityVerification**](AccountApi.md#GetAccountIdentityVerification) | **Get** /v1/account/identity-verification | Your identity verification: one status for the whole account, read by every feature that needs ID.
[**GetMe**](AccountApi.md#GetMe) | **Get** /v1/me | Identity check — authenticated user + PAT scopes + account state.
[**GetMembership**](AccountApi.md#GetMembership) | **Get** /v1/membership | The caller&#39;s Crossly One membership status and benefits.
[**ListAuthSessions**](AccountApi.md#ListAuthSessions) | **Get** /v1/auth/sessions | List active browser auth sessions.
[**ListConnectedApps**](AccountApi.md#ListConnectedApps) | **Get** /v1/connected-apps | List third-party OAuth apps with access to this account.



## CreateAccountCancelDeletion

> CreateAccountCancelDeletionResponse CreateAccountCancelDeletion(ctx).Execute()

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
    // response from `CreateAccountCancelDeletion`: CreateAccountCancelDeletionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CreateAccountCancelDeletion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountCancelDeletionRequest struct via the builder pattern


### Return type

**CreateAccountCancelDeletionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateAccountIdentityVerificationStart

> CreateWorldMatureVerifyResponse CreateAccountIdentityVerificationStart(ctx).CreateAccountIdentityVerificationStartRequest(createAccountIdentityVerificationStartRequest).Execute()

Start the identity check (photo ID + selfie). Returns a Stripe-hosted URL the person completes.



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
    createAccountIdentityVerificationStartRequest := *openapiclient.NewCreateAccountIdentityVerificationStartRequest() // CreateAccountIdentityVerificationStartRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AccountApi.CreateAccountIdentityVerificationStart(context.Background()).CreateAccountIdentityVerificationStartRequest(createAccountIdentityVerificationStartRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CreateAccountIdentityVerificationStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountIdentityVerificationStart`: CreateWorldMatureVerifyResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CreateAccountIdentityVerificationStart`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountIdentityVerificationStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createAccountIdentityVerificationStartRequest** | **CreateAccountIdentityVerificationStartRequest** |  | 

### Return type

**CreateWorldMatureVerifyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
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


## CreateAccountReconfirmDeletion

> CreateAccountReconfirmDeletionResponse CreateAccountReconfirmDeletion(ctx).Execute()

Confirm again a deletion paused because the account was in use.

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
    resp, r, err := apiClient.AccountApi.CreateAccountReconfirmDeletion(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.CreateAccountReconfirmDeletion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAccountReconfirmDeletion`: CreateAccountReconfirmDeletionResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.CreateAccountReconfirmDeletion`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccountReconfirmDeletionRequest struct via the builder pattern


### Return type

**CreateAccountReconfirmDeletionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateAccountRequestDeletion

> CreateAccountRequestDeletionResponse CreateAccountRequestDeletion(ctx).Execute()

Schedule account deletion after a grace period. Requires confirmEmail.

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
    sessionId := "sessionId_example" // string | 

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


## GetAccountClosureChecklist

> GetAccountClosureChecklistResponse GetAccountClosureChecklist(ctx).Execute()

What closing the account will and will not do, before you request it.

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
    resp, r, err := apiClient.AccountApi.GetAccountClosureChecklist(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetAccountClosureChecklist``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountClosureChecklist`: GetAccountClosureChecklistResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetAccountClosureChecklist`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountClosureChecklistRequest struct via the builder pattern


### Return type

**GetAccountClosureChecklistResponse**

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


## GetAccountIdentityVerification

> GetAccountIdentityVerificationResponse GetAccountIdentityVerification(ctx).Execute()

Your identity verification: one status for the whole account, read by every feature that needs ID.



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
    resp, r, err := apiClient.AccountApi.GetAccountIdentityVerification(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountApi.GetAccountIdentityVerification``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAccountIdentityVerification`: GetAccountIdentityVerificationResponse
    fmt.Fprintf(os.Stdout, "Response from `AccountApi.GetAccountIdentityVerification`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccountIdentityVerificationRequest struct via the builder pattern


### Return type

**GetAccountIdentityVerificationResponse**

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

