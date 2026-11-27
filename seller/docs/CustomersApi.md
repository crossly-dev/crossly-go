# \CustomersApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCustomerBulkDelete**](CustomersApi.md#CreateCustomerBulkDelete) | **Post** /v1/customers/bulk-delete | Bulk blocklist customer handles.
[**CreateCustomerBulkExport**](CustomersApi.md#CreateCustomerBulkExport) | **Post** /v1/customers/bulk-export | Bulk export aggregated customers as CSV.
[**CreateCustomerRefreshProfile**](CustomersApi.md#CreateCustomerRefreshProfile) | **Post** /v1/customers/{handle}/refresh-profile | Queue a fresh read of a customer&#39;s buyer profiles on every platform they ordered on.
[**CreateLoyaltyAirdropCancel**](CustomersApi.md#CreateLoyaltyAirdropCancel) | **Post** /v1/loyalty/airdrops/{id}/cancel | Cancel a queued or failed airdrop (one already sent is the member’s).
[**CreateLoyaltyAirdropCheck**](CustomersApi.md#CreateLoyaltyAirdropCheck) | **Post** /v1/loyalty/airdrops/{id}/check | Read a submitted airdrop’s receipt again now.
[**CreateLoyaltyAirdropPrepare**](CustomersApi.md#CreateLoyaltyAirdropPrepare) | **Post** /v1/loyalty/airdrops/{id}/prepare | The mint transaction for your launch owner wallet to send: chain, from, to (your launch contract) and calldata, to the member’s verified wallet. Refused airdrop_not_sendable with a reason when it cannot be built.
[**CreateLoyaltyAirdropSubmit**](CustomersApi.md#CreateLoyaltyAirdropSubmit) | **Post** /v1/loyalty/airdrops/{id}/submit | Record the hash your wallet broadcast; Crossly reads the receipt (now, then on a schedule) and marks it sent or failed.
[**CreateLoyaltyProgramInvite**](CustomersApi.md#CreateLoyaltyProgramInvite) | **Post** /v1/loyalty/program/invite | Invite Crossly buyers to your loyalty program by user id (up to 500); returns how many were newly invited.
[**CreateLoyaltyProgramMemberOverride**](CustomersApi.md#CreateLoyaltyProgramMemberOverride) | **Post** /v1/loyalty/program/members/{id}/override | Pin a member to a tier (tierId null clears the pin), with a reason.
[**CreateLoyaltyTier**](CustomersApi.md#CreateLoyaltyTier) | **Post** /v1/loyalty/tiers | Add a tier: name, look, threshold and typed perks (% off, free shipping, early access, members-only listings, wholesale access, token gates, airdrops, Bucks, World items, event tickets, …).
[**DeleteLoyaltyTier**](CustomersApi.md#DeleteLoyaltyTier) | **Delete** /v1/loyalty/tiers/{id} | Delete a tier; its members are re-evaluated.
[**GetCustomer**](CustomersApi.md#GetCustomer) | **Get** /v1/customers/{handle} | Get one customer: recent 50 orders, buyer profiles per platform, and stats over all of their orders.
[**GetCustomerNote**](CustomersApi.md#GetCustomerNote) | **Get** /v1/customers/{handle}/notes | Your private notes on a customer (empty when none). Works before their first order.
[**GetLoyaltyAirdrop**](CustomersApi.md#GetLoyaltyAirdrop) | **Get** /v1/loyalty/airdrops/{id} | One loyalty NFT airdrop: member, collection, recipient wallet, transaction hash and status.
[**GetLoyaltyClub**](CustomersApi.md#GetLoyaltyClub) | **Get** /v1/loyalty/club | Your loyalty club: tiers and how many members are in each.
[**GetLoyaltyProgram**](CustomersApi.md#GetLoyaltyProgram) | **Get** /v1/loyalty/program | Your loyalty program: qualification (spend, orders, points), window, joining, downgrade and grace, and its tiers with typed perks.
[**GetLoyaltyProgramMember**](CustomersApi.md#GetLoyaltyProgramMember) | **Get** /v1/loyalty/program/members/{id} | One member: standing, tier, progress, perks held, history and any override.
[**ListCustomers**](CustomersApi.md#ListCustomers) | **Get** /v1/customers | List aggregated customers (group-by lower(buyer_username)), each with its best avatar, display name and star rating.
[**ListLoyaltyAirdrops**](CustomersApi.md#ListLoyaltyAirdrops) | **Get** /v1/loyalty/airdrops | NFT airdrops your loyalty tiers earned members: queued to send, submitted (waiting on the receipt), sent, failed (with why) or cancelled.
[**ListLoyaltyMembers**](CustomersApi.md#ListLoyaltyMembers) | **Get** /v1/loyalty/members | Your loyalty club’s members, by spend with you.
[**ListLoyaltyProgramInviteCandidates**](CustomersApi.md#ListLoyaltyProgramInviteCandidates) | **Get** /v1/loyalty/program/invite-candidates | Crossly buyers who bought from you and are not loyalty members yet (q searches by name): who you can invite.
[**ListLoyaltyProgramMembers**](CustomersApi.md#ListLoyaltyProgramMembers) | **Get** /v1/loyalty/program/members | Your loyalty members with tier, standing, grace and overrides; search, tier, status and sort filters.
[**UpdateCustomerNote**](CustomersApi.md#UpdateCustomerNote) | **Put** /v1/customers/{handle}/notes | Save your private notes on a customer (an empty string clears them).
[**UpdateLoyaltyClub**](CustomersApi.md#UpdateLoyaltyClub) | **Put** /v1/loyalty/club | Create or update your loyalty club: name, on/off, and up to 6 spend tiers each with a discount (≤ 30%) and perks.
[**UpdateLoyaltyProgram**](CustomersApi.md#UpdateLoyaltyProgram) | **Put** /v1/loyalty/program | Create or update your loyalty program settings.
[**UpdateLoyaltyTier**](CustomersApi.md#UpdateLoyaltyTier) | **Put** /v1/loyalty/tiers/{id} | Replace a tier. Members holding it get the edited perks.



## CreateCustomerBulkDelete

> CreateCustomerBulkDeleteResponse CreateCustomerBulkDelete(ctx).Execute()

Bulk blocklist customer handles.

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
    resp, r, err := apiClient.CustomersApi.CreateCustomerBulkDelete(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateCustomerBulkDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerBulkDelete`: CreateCustomerBulkDeleteResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateCustomerBulkDelete`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerBulkDeleteRequest struct via the builder pattern


### Return type

**CreateCustomerBulkDeleteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCustomerBulkExport

> string CreateCustomerBulkExport(ctx).Execute()

Bulk export aggregated customers as CSV.

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
    resp, r, err := apiClient.CustomersApi.CreateCustomerBulkExport(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateCustomerBulkExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerBulkExport`: string
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateCustomerBulkExport`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerBulkExportRequest struct via the builder pattern


### Return type

**string**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateCustomerRefreshProfile

> CreateCustomerRefreshProfileResponse CreateCustomerRefreshProfile(ctx, handle).Execute()

Queue a fresh read of a customer's buyer profiles on every platform they ordered on.

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
    handle := "handle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateCustomerRefreshProfile(context.Background(), handle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateCustomerRefreshProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCustomerRefreshProfile`: CreateCustomerRefreshProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateCustomerRefreshProfile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCustomerRefreshProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateCustomerRefreshProfileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLoyaltyAirdropCancel

> LoyaltyAirdropResponse CreateLoyaltyAirdropCancel(ctx, id).Execute()

Cancel a queued or failed airdrop (one already sent is the member’s).

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateLoyaltyAirdropCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateLoyaltyAirdropCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoyaltyAirdropCancel`: LoyaltyAirdropResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateLoyaltyAirdropCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoyaltyAirdropCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**LoyaltyAirdropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLoyaltyAirdropCheck

> LoyaltyAirdropResponse CreateLoyaltyAirdropCheck(ctx, id).Execute()

Read a submitted airdrop’s receipt again now.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateLoyaltyAirdropCheck(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateLoyaltyAirdropCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoyaltyAirdropCheck`: LoyaltyAirdropResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateLoyaltyAirdropCheck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoyaltyAirdropCheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**LoyaltyAirdropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLoyaltyAirdropPrepare

> CreateLoyaltyAirdropPrepareResponse CreateLoyaltyAirdropPrepare(ctx, id).Execute()

The mint transaction for your launch owner wallet to send: chain, from, to (your launch contract) and calldata, to the member’s verified wallet. Refused airdrop_not_sendable with a reason when it cannot be built.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateLoyaltyAirdropPrepare(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateLoyaltyAirdropPrepare``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoyaltyAirdropPrepare`: CreateLoyaltyAirdropPrepareResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateLoyaltyAirdropPrepare`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoyaltyAirdropPrepareRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateLoyaltyAirdropPrepareResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLoyaltyAirdropSubmit

> LoyaltyAirdropResponse CreateLoyaltyAirdropSubmit(ctx, id).Execute()

Record the hash your wallet broadcast; Crossly reads the receipt (now, then on a schedule) and marks it sent or failed.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateLoyaltyAirdropSubmit(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateLoyaltyAirdropSubmit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoyaltyAirdropSubmit`: LoyaltyAirdropResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateLoyaltyAirdropSubmit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoyaltyAirdropSubmitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**LoyaltyAirdropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLoyaltyProgramInvite

> CreateLoyaltyProgramInviteResponse CreateLoyaltyProgramInvite(ctx).Execute()

Invite Crossly buyers to your loyalty program by user id (up to 500); returns how many were newly invited.

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
    resp, r, err := apiClient.CustomersApi.CreateLoyaltyProgramInvite(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateLoyaltyProgramInvite``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoyaltyProgramInvite`: CreateLoyaltyProgramInviteResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateLoyaltyProgramInvite`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoyaltyProgramInviteRequest struct via the builder pattern


### Return type

**CreateLoyaltyProgramInviteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLoyaltyProgramMemberOverride

> GetLoyaltyProgramMemberResponse CreateLoyaltyProgramMemberOverride(ctx, id).Execute()

Pin a member to a tier (tierId null clears the pin), with a reason.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.CreateLoyaltyProgramMemberOverride(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateLoyaltyProgramMemberOverride``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoyaltyProgramMemberOverride`: GetLoyaltyProgramMemberResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateLoyaltyProgramMemberOverride`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoyaltyProgramMemberOverrideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetLoyaltyProgramMemberResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLoyaltyTier

> CreateLoyaltyTierResponse CreateLoyaltyTier(ctx).Execute()

Add a tier: name, look, threshold and typed perks (% off, free shipping, early access, members-only listings, wholesale access, token gates, airdrops, Bucks, World items, event tickets, …).

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
    resp, r, err := apiClient.CustomersApi.CreateLoyaltyTier(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.CreateLoyaltyTier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLoyaltyTier`: CreateLoyaltyTierResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.CreateLoyaltyTier`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLoyaltyTierRequest struct via the builder pattern


### Return type

**CreateLoyaltyTierResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteLoyaltyTier

> UpdateGroupBuyResponse DeleteLoyaltyTier(ctx, id).Execute()

Delete a tier; its members are re-evaluated.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.DeleteLoyaltyTier(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.DeleteLoyaltyTier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteLoyaltyTier`: UpdateGroupBuyResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.DeleteLoyaltyTier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLoyaltyTierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateGroupBuyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCustomer

> GetCustomerResponse GetCustomer(ctx, handle).Execute()

Get one customer: recent 50 orders, buyer profiles per platform, and stats over all of their orders.

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
    handle := "handle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.GetCustomer(context.Background(), handle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetCustomer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomer`: GetCustomerResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetCustomer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetCustomerResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetCustomerNote

> GetCustomerNoteResponse GetCustomerNote(ctx, handle).Execute()

Your private notes on a customer (empty when none). Works before their first order.

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
    handle := "handle_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.GetCustomerNote(context.Background(), handle).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetCustomerNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCustomerNote`: GetCustomerNoteResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetCustomerNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCustomerNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetCustomerNoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetLoyaltyAirdrop

> LoyaltyAirdropResponse GetLoyaltyAirdrop(ctx, id).Execute()

One loyalty NFT airdrop: member, collection, recipient wallet, transaction hash and status.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.GetLoyaltyAirdrop(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetLoyaltyAirdrop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoyaltyAirdrop`: LoyaltyAirdropResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetLoyaltyAirdrop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyAirdropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**LoyaltyAirdropResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetLoyaltyClub

> GetLoyaltyClubResponse GetLoyaltyClub(ctx).Execute()

Your loyalty club: tiers and how many members are in each.

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
    resp, r, err := apiClient.CustomersApi.GetLoyaltyClub(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetLoyaltyClub``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoyaltyClub`: GetLoyaltyClubResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetLoyaltyClub`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyClubRequest struct via the builder pattern


### Return type

**GetLoyaltyClubResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetLoyaltyProgram

> GetLoyaltyProgramResponse GetLoyaltyProgram(ctx).Execute()

Your loyalty program: qualification (spend, orders, points), window, joining, downgrade and grace, and its tiers with typed perks.

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
    resp, r, err := apiClient.CustomersApi.GetLoyaltyProgram(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetLoyaltyProgram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoyaltyProgram`: GetLoyaltyProgramResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetLoyaltyProgram`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramRequest struct via the builder pattern


### Return type

**GetLoyaltyProgramResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetLoyaltyProgramMember

> GetLoyaltyProgramMemberResponse GetLoyaltyProgramMember(ctx, id).Execute()

One member: standing, tier, progress, perks held, history and any override.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.GetLoyaltyProgramMember(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.GetLoyaltyProgramMember``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLoyaltyProgramMember`: GetLoyaltyProgramMemberResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.GetLoyaltyProgramMember`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLoyaltyProgramMemberRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetLoyaltyProgramMemberResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListCustomers

> ListCustomersResponse ListCustomers(ctx).Execute()

List aggregated customers (group-by lower(buyer_username)), each with its best avatar, display name and star rating.

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
    resp, r, err := apiClient.CustomersApi.ListCustomers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.ListCustomers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCustomers`: ListCustomersResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.ListCustomers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListCustomersRequest struct via the builder pattern


### Return type

**ListCustomersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListLoyaltyAirdrops

> ListLoyaltyAirdropsResponse ListLoyaltyAirdrops(ctx).Execute()

NFT airdrops your loyalty tiers earned members: queued to send, submitted (waiting on the receipt), sent, failed (with why) or cancelled.

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
    resp, r, err := apiClient.CustomersApi.ListLoyaltyAirdrops(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.ListLoyaltyAirdrops``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoyaltyAirdrops`: ListLoyaltyAirdropsResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.ListLoyaltyAirdrops`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLoyaltyAirdropsRequest struct via the builder pattern


### Return type

**ListLoyaltyAirdropsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListLoyaltyMembers

> ListLoyaltyMembersResponse ListLoyaltyMembers(ctx).Execute()

Your loyalty club’s members, by spend with you.

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
    resp, r, err := apiClient.CustomersApi.ListLoyaltyMembers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.ListLoyaltyMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoyaltyMembers`: ListLoyaltyMembersResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.ListLoyaltyMembers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLoyaltyMembersRequest struct via the builder pattern


### Return type

**ListLoyaltyMembersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListLoyaltyProgramInviteCandidates

> ListLoyaltyProgramInviteCandidatesResponse ListLoyaltyProgramInviteCandidates(ctx).Execute()

Crossly buyers who bought from you and are not loyalty members yet (q searches by name): who you can invite.

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
    resp, r, err := apiClient.CustomersApi.ListLoyaltyProgramInviteCandidates(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.ListLoyaltyProgramInviteCandidates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoyaltyProgramInviteCandidates`: ListLoyaltyProgramInviteCandidatesResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.ListLoyaltyProgramInviteCandidates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLoyaltyProgramInviteCandidatesRequest struct via the builder pattern


### Return type

**ListLoyaltyProgramInviteCandidatesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListLoyaltyProgramMembers

> ListLoyaltyProgramMembersResponse ListLoyaltyProgramMembers(ctx).Execute()

Your loyalty members with tier, standing, grace and overrides; search, tier, status and sort filters.

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
    resp, r, err := apiClient.CustomersApi.ListLoyaltyProgramMembers(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.ListLoyaltyProgramMembers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLoyaltyProgramMembers`: ListLoyaltyProgramMembersResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.ListLoyaltyProgramMembers`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLoyaltyProgramMembersRequest struct via the builder pattern


### Return type

**ListLoyaltyProgramMembersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateCustomerNote

> GetCustomerNoteResponse UpdateCustomerNote(ctx, handle).UpdateCustomerNoteRequest(updateCustomerNoteRequest).Execute()

Save your private notes on a customer (an empty string clears them).

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
    handle := "handle_example" // string | 
    updateCustomerNoteRequest := *openapiclient.NewUpdateCustomerNoteRequest("Notes_example") // UpdateCustomerNoteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.UpdateCustomerNote(context.Background(), handle).UpdateCustomerNoteRequest(updateCustomerNoteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.UpdateCustomerNote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCustomerNote`: GetCustomerNoteResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.UpdateCustomerNote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**handle** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCustomerNoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateCustomerNoteRequest** | **UpdateCustomerNoteRequest** |  | 

### Return type

**GetCustomerNoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateLoyaltyClub

> GetLoyaltyClubResponse UpdateLoyaltyClub(ctx).Execute()

Create or update your loyalty club: name, on/off, and up to 6 spend tiers each with a discount (≤ 30%) and perks.

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
    resp, r, err := apiClient.CustomersApi.UpdateLoyaltyClub(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.UpdateLoyaltyClub``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLoyaltyClub`: GetLoyaltyClubResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.UpdateLoyaltyClub`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoyaltyClubRequest struct via the builder pattern


### Return type

**GetLoyaltyClubResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateLoyaltyProgram

> UpdateLoyaltyProgramResponse UpdateLoyaltyProgram(ctx).Execute()

Create or update your loyalty program settings.

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
    resp, r, err := apiClient.CustomersApi.UpdateLoyaltyProgram(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.UpdateLoyaltyProgram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLoyaltyProgram`: UpdateLoyaltyProgramResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.UpdateLoyaltyProgram`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoyaltyProgramRequest struct via the builder pattern


### Return type

**UpdateLoyaltyProgramResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateLoyaltyTier

> CreateLoyaltyTierResponse UpdateLoyaltyTier(ctx, id).Execute()

Replace a tier. Members holding it get the edited perks.

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
    id := "id_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CustomersApi.UpdateLoyaltyTier(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CustomersApi.UpdateLoyaltyTier``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLoyaltyTier`: CreateLoyaltyTierResponse
    fmt.Fprintf(os.Stdout, "Response from `CustomersApi.UpdateLoyaltyTier`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLoyaltyTierRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateLoyaltyTierResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

