# \SourcingApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateSourcingReceipt**](SourcingApi.md#CreateSourcingReceipt) | **Post** /v1/sourcing/receipts | Append a parsed receipt to the sourcing ledger.
[**GetSourcingReceipt**](SourcingApi.md#GetSourcingReceipt) | **Get** /v1/sourcing/receipts | List parsed sourcing receipts in this user&#39;s ledger.
[**ListSourcingDemand**](SourcingApi.md#ListSourcingDemand) | **Get** /v1/sourcing/demand | Items buyers looked for on other sites that Crossly did not have.
[**ListSourcingDemandMine**](SourcingApi.md#ListSourcingDemandMine) | **Get** /v1/sourcing/demand/mine | Unmet buyer demand for items you hold or have sold before.



## CreateSourcingReceipt

> CreateSourcingReceiptResponse CreateSourcingReceipt(ctx).Execute()

Append a parsed receipt to the sourcing ledger.

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
    resp, r, err := apiClient.SourcingApi.CreateSourcingReceipt(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcingApi.CreateSourcingReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSourcingReceipt`: CreateSourcingReceiptResponse
    fmt.Fprintf(os.Stdout, "Response from `SourcingApi.CreateSourcingReceipt`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSourcingReceiptRequest struct via the builder pattern


### Return type

[**CreateSourcingReceiptResponse**](CreateSourcingReceiptResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetSourcingReceipt

> GetSourcingReceiptResponse GetSourcingReceipt(ctx).Execute()

List parsed sourcing receipts in this user's ledger.

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
    resp, r, err := apiClient.SourcingApi.GetSourcingReceipt(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcingApi.GetSourcingReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetSourcingReceipt`: GetSourcingReceiptResponse
    fmt.Fprintf(os.Stdout, "Response from `SourcingApi.GetSourcingReceipt`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSourcingReceiptRequest struct via the builder pattern


### Return type

[**GetSourcingReceiptResponse**](GetSourcingReceiptResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourcingDemand

> V1List ListSourcingDemand(ctx).Days(days).MinLooks(minLooks).Limit(limit).Execute()

Items buyers looked for on other sites that Crossly did not have.



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
    days := int32(56) // int32 |  (optional) (default to 30)
    minLooks := int32(56) // int32 |  (optional) (default to 3)
    limit := int32(56) // int32 |  (optional) (default to 50)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcingApi.ListSourcingDemand(context.Background()).Days(days).MinLooks(minLooks).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcingApi.ListSourcingDemand``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourcingDemand`: V1List
    fmt.Fprintf(os.Stdout, "Response from `SourcingApi.ListSourcingDemand`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSourcingDemandRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** |  | [default to 30]
 **minLooks** | **int32** |  | [default to 3]
 **limit** | **int32** |  | [default to 50]

### Return type

[**V1List**](V1List.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSourcingDemandMine

> V1List ListSourcingDemandMine(ctx).Days(days).MinLookers(minLookers).Limit(limit).Execute()

Unmet buyer demand for items you hold or have sold before.

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
    days := int32(56) // int32 |  (optional) (default to 60)
    minLookers := int32(56) // int32 |  (optional) (default to 2)
    limit := int32(56) // int32 |  (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SourcingApi.ListSourcingDemandMine(context.Background()).Days(days).MinLookers(minLookers).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SourcingApi.ListSourcingDemandMine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSourcingDemandMine`: V1List
    fmt.Fprintf(os.Stdout, "Response from `SourcingApi.ListSourcingDemandMine`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSourcingDemandMineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **days** | **int32** |  | [default to 60]
 **minLookers** | **int32** |  | [default to 2]
 **limit** | **int32** |  | [default to 25]

### Return type

[**V1List**](V1List.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

