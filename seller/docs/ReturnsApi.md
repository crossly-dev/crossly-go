# \ReturnsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePolicyReturnReceived**](ReturnsApi.md#CreatePolicyReturnReceived) | **Post** /v1/policy-returns/{id}/received | Mark a return received: refunds the buyer, less the restocking or try fee you keep.
[**CreateReturn**](ReturnsApi.md#CreateReturn) | **Post** /v1/returns | Open a return record on an order (rejects if another open return exists).
[**CreateReturnAccept**](ReturnsApi.md#CreateReturnAccept) | **Post** /v1/returns/{id}/accept | Accept a Crossly-marketplace return request: a prepaid return label to your return address (the buyer has 14 days to ship), a return record linked to it, and optionally a refund now (refund: { amount? } in dollars; needs orders:refund). Without one, the rest is refunded when the return is delivered. Safe to repeat. Crossposted orders: not_supported_on_platform.
[**CreateReturnDecline**](ReturnsApi.md#CreateReturnDecline) | **Post** /v1/returns/{id}/decline | Decline a Crossly-marketplace return request with a reason (required, 8-2000 chars), posted to the buyer as your response. It does not close the request: the buyer can reply or ask Crossly to decide. Not after the label is issued.
[**GetReturn**](ReturnsApi.md#GetReturn) | **Get** /v1/returns/{id} | Get one return record.
[**ListPolicyReturns**](ReturnsApi.md#ListPolicyReturns) | **Get** /v1/policy-returns | Returns Crossly buyers opened under your return policy or try at home; filter by orderId.
[**ListReturns**](ReturnsApi.md#ListReturns) | **Get** /v1/returns | List returns. Default: your return records (the physical-return workflow), status&#x3D;open|closed|&lt;exact&gt;. source&#x3D;request: the OPEN return requests buyers made on the Crossly marketplace, newest first; accept or decline each with POST /v1/returns/{id}/accept | decline.
[**UpdateReturn**](ReturnsApi.md#UpdateReturn) | **Patch** /v1/returns/{id} | Transition return status (received/inspected/restocked) and bump inventory on restock.



## CreatePolicyReturnReceived

> CreatePolicyReturnReceivedResponse CreatePolicyReturnReceived(ctx, id).Execute()

Mark a return received: refunds the buyer, less the restocking or try fee you keep.

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
    resp, r, err := apiClient.ReturnsApi.CreatePolicyReturnReceived(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.CreatePolicyReturnReceived``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePolicyReturnReceived`: CreatePolicyReturnReceivedResponse
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.CreatePolicyReturnReceived`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePolicyReturnReceivedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatePolicyReturnReceivedResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateReturn

> GetReturnResponse CreateReturn(ctx).Execute()

Open a return record on an order (rejects if another open return exists).

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
    resp, r, err := apiClient.ReturnsApi.CreateReturn(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.CreateReturn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReturn`: GetReturnResponse
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.CreateReturn`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReturnRequest struct via the builder pattern


### Return type

**GetReturnResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateReturnAccept

> CreateReturnAcceptResponse CreateReturnAccept(ctx, id).Execute()

Accept a Crossly-marketplace return request: a prepaid return label to your return address (the buyer has 14 days to ship), a return record linked to it, and optionally a refund now (refund: { amount? } in dollars; needs orders:refund). Without one, the rest is refunded when the return is delivered. Safe to repeat. Crossposted orders: not_supported_on_platform.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnsApi.CreateReturnAccept(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.CreateReturnAccept``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReturnAccept`: CreateReturnAcceptResponse
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.CreateReturnAccept`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReturnAcceptRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateReturnAcceptResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateReturnDecline

> ListReturnsItem CreateReturnDecline(ctx, id).CreateReturnDeclineRequest(createReturnDeclineRequest).Execute()

Decline a Crossly-marketplace return request with a reason (required, 8-2000 chars), posted to the buyer as your response. It does not close the request: the buyer can reply or ask Crossly to decide. Not after the label is issued.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 
    createReturnDeclineRequest := *openapiclient.NewCreateReturnDeclineRequest("Reason_example") // CreateReturnDeclineRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnsApi.CreateReturnDecline(context.Background(), id).CreateReturnDeclineRequest(createReturnDeclineRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.CreateReturnDecline``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateReturnDecline`: ListReturnsItem
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.CreateReturnDecline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateReturnDeclineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createReturnDeclineRequest** | **CreateReturnDeclineRequest** |  | 

### Return type

**ListReturnsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetReturn

> GetReturnResponse GetReturn(ctx, id).Execute()

Get one return record.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnsApi.GetReturn(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.GetReturn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetReturn`: GetReturnResponse
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.GetReturn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetReturnResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListPolicyReturns

> ListPolicyReturnsResponse ListPolicyReturns(ctx).Execute()

Returns Crossly buyers opened under your return policy or try at home; filter by orderId.

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
    resp, r, err := apiClient.ReturnsApi.ListPolicyReturns(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.ListPolicyReturns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPolicyReturns`: ListPolicyReturnsResponse
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.ListPolicyReturns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPolicyReturnsRequest struct via the builder pattern


### Return type

**ListPolicyReturnsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListReturns

> ListReturnsResponse ListReturns(ctx).Execute()

List returns. Default: your return records (the physical-return workflow), status=open|closed|<exact>. source=request: the OPEN return requests buyers made on the Crossly marketplace, newest first; accept or decline each with POST /v1/returns/{id}/accept | decline.

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
    resp, r, err := apiClient.ReturnsApi.ListReturns(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.ListReturns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListReturns`: ListReturnsResponse
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.ListReturns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListReturnsRequest struct via the builder pattern


### Return type

**ListReturnsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateReturn

> GetReturnResponse UpdateReturn(ctx, id).Execute()

Transition return status (received/inspected/restocked) and bump inventory on restock.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ReturnsApi.UpdateReturn(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ReturnsApi.UpdateReturn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateReturn`: GetReturnResponse
    fmt.Fprintf(os.Stdout, "Response from `ReturnsApi.UpdateReturn`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReturnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetReturnResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

