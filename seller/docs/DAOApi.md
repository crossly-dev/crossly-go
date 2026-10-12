# \DAOApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDaoProposal**](DAOApi.md#CreateDaoProposal) | **Post** /v1/dao/proposals | Raise a proposal. Starts as a draft.
[**CreateDaoProposalFinalise**](DAOApi.md#CreateDaoProposalFinalise) | **Post** /v1/dao/proposals/{id}/finalise | Close a proposal whose window has passed and record the outcome.
[**CreateDaoProposalOpen**](DAOApi.md#CreateDaoProposalOpen) | **Post** /v1/dao/proposals/{id}/open | Put your draft to a vote. Snapshots voting power.
[**CreateDaoProposalVeto**](DAOApi.md#CreateDaoProposalVeto) | **Post** /v1/dao/proposals/{id}/veto | Operator block. There is no matching power to PASS.
[**CreateDaoProposalVote**](DAOApi.md#CreateDaoProposalVote) | **Post** /v1/dao/proposals/{id}/vote | Cast or change a vote while the window is open.
[**DeleteDaoProposal**](DAOApi.md#DeleteDaoProposal) | **Delete** /v1/dao/proposals/{id} | Withdraw your own DRAFT.
[**GetDaoProposalTally**](DAOApi.md#GetDaoProposalTally) | **Get** /v1/dao/proposals/{id}/tally | The live count.
[**ListDaoProposals**](DAOApi.md#ListDaoProposals) | **Get** /v1/dao/proposals | The proposal queue.



## CreateDaoProposal

> CreateDaoProposalResponse CreateDaoProposal(ctx).Execute()

Raise a proposal. Starts as a draft.



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
    resp, r, err := apiClient.DAOApi.CreateDaoProposal(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAOApi.CreateDaoProposal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDaoProposal`: CreateDaoProposalResponse
    fmt.Fprintf(os.Stdout, "Response from `DAOApi.CreateDaoProposal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDaoProposalRequest struct via the builder pattern


### Return type

**CreateDaoProposalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDaoProposalFinalise

> CreateDaoProposalFinaliseResponse CreateDaoProposalFinalise(ctx, id).Execute()

Close a proposal whose window has passed and record the outcome.



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
    resp, r, err := apiClient.DAOApi.CreateDaoProposalFinalise(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAOApi.CreateDaoProposalFinalise``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDaoProposalFinalise`: CreateDaoProposalFinaliseResponse
    fmt.Fprintf(os.Stdout, "Response from `DAOApi.CreateDaoProposalFinalise`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDaoProposalFinaliseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateDaoProposalFinaliseResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDaoProposalOpen

> CreateDaoProposalOpenResponse CreateDaoProposalOpen(ctx, id).Execute()

Put your draft to a vote. Snapshots voting power.



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
    resp, r, err := apiClient.DAOApi.CreateDaoProposalOpen(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAOApi.CreateDaoProposalOpen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDaoProposalOpen`: CreateDaoProposalOpenResponse
    fmt.Fprintf(os.Stdout, "Response from `DAOApi.CreateDaoProposalOpen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDaoProposalOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateDaoProposalOpenResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDaoProposalVeto

> CreateDaoProposalVetoResponse CreateDaoProposalVeto(ctx, id).Execute()

Operator block. There is no matching power to PASS.



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
    resp, r, err := apiClient.DAOApi.CreateDaoProposalVeto(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAOApi.CreateDaoProposalVeto``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDaoProposalVeto`: CreateDaoProposalVetoResponse
    fmt.Fprintf(os.Stdout, "Response from `DAOApi.CreateDaoProposalVeto`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDaoProposalVetoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateDaoProposalVetoResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDaoProposalVote

> CreateDaoProposalVoteResponse CreateDaoProposalVote(ctx, id).Execute()

Cast or change a vote while the window is open.



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
    resp, r, err := apiClient.DAOApi.CreateDaoProposalVote(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAOApi.CreateDaoProposalVote``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDaoProposalVote`: CreateDaoProposalVoteResponse
    fmt.Fprintf(os.Stdout, "Response from `DAOApi.CreateDaoProposalVote`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDaoProposalVoteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateDaoProposalVoteResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteDaoProposal

> DeleteDaoProposalResponse DeleteDaoProposal(ctx, id).Execute()

Withdraw your own DRAFT.



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
    resp, r, err := apiClient.DAOApi.DeleteDaoProposal(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAOApi.DeleteDaoProposal``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDaoProposal`: DeleteDaoProposalResponse
    fmt.Fprintf(os.Stdout, "Response from `DAOApi.DeleteDaoProposal`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDaoProposalRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteDaoProposalResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetDaoProposalTally

> GetDaoProposalTallyResponse GetDaoProposalTally(ctx, id).Execute()

The live count.



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
    resp, r, err := apiClient.DAOApi.GetDaoProposalTally(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAOApi.GetDaoProposalTally``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDaoProposalTally`: GetDaoProposalTallyResponse
    fmt.Fprintf(os.Stdout, "Response from `DAOApi.GetDaoProposalTally`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDaoProposalTallyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetDaoProposalTallyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListDaoProposals

> ListDaoProposalsResponse ListDaoProposals(ctx).Execute()

The proposal queue.



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
    resp, r, err := apiClient.DAOApi.ListDaoProposals(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DAOApi.ListDaoProposals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDaoProposals`: ListDaoProposalsResponse
    fmt.Fprintf(os.Stdout, "Response from `DAOApi.ListDaoProposals`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDaoProposalsRequest struct via the builder pattern


### Return type

**ListDaoProposalsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

