# \CrosslyWalletApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTcgNftReconcile**](CrosslyWalletApi.md#CreateTcgNftReconcile) | **Post** /v1/tcg/nfts/reconcile | Re-read your wallets on chain so your card counts follow what you hold (cards you moved away leave; cards you received join).
[**CreateWalletCrossly**](CrosslyWalletApi.md#CreateWalletCrossly) | **Post** /v1/wallet/crossly | Create your Crossly wallet now (it is otherwise created for you in the background). Idempotent.
[**GetWalletCrossly**](CrosslyWalletApi.md#GetWalletCrossly) | **Get** /v1/wallet/crossly | Your Crossly wallet (self-custody Solana): status, address, whether it is configured, and your card NFT counts by status.
[**ListTcgNfts**](CrosslyWalletApi.md#ListTcgNfts) | **Get** /v1/tcg/{game}/nfts | Your card NFTs for a Crossly TCG game: each copy, its mint status, asset id and whether it is soulbound (starter).



## CreateTcgNftReconcile

> CreateTcgNftReconcileResponse CreateTcgNftReconcile(ctx).Execute()

Re-read your wallets on chain so your card counts follow what you hold (cards you moved away leave; cards you received join).

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
    resp, r, err := apiClient.CrosslyWalletApi.CreateTcgNftReconcile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CrosslyWalletApi.CreateTcgNftReconcile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgNftReconcile`: CreateTcgNftReconcileResponse
    fmt.Fprintf(os.Stdout, "Response from `CrosslyWalletApi.CreateTcgNftReconcile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgNftReconcileRequest struct via the builder pattern


### Return type

**CreateTcgNftReconcileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWalletCrossly

> GetWalletCrosslyResponse CreateWalletCrossly(ctx).Execute()

Create your Crossly wallet now (it is otherwise created for you in the background). Idempotent.

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
    resp, r, err := apiClient.CrosslyWalletApi.CreateWalletCrossly(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CrosslyWalletApi.CreateWalletCrossly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWalletCrossly`: GetWalletCrosslyResponse
    fmt.Fprintf(os.Stdout, "Response from `CrosslyWalletApi.CreateWalletCrossly`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWalletCrosslyRequest struct via the builder pattern


### Return type

**GetWalletCrosslyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWalletCrossly

> GetWalletCrosslyResponse GetWalletCrossly(ctx).Execute()

Your Crossly wallet (self-custody Solana): status, address, whether it is configured, and your card NFT counts by status.

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
    resp, r, err := apiClient.CrosslyWalletApi.GetWalletCrossly(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CrosslyWalletApi.GetWalletCrossly``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWalletCrossly`: GetWalletCrosslyResponse
    fmt.Fprintf(os.Stdout, "Response from `CrosslyWalletApi.GetWalletCrossly`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWalletCrosslyRequest struct via the builder pattern


### Return type

**GetWalletCrosslyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListTcgNfts

> ListTcgNftsResponse ListTcgNfts(ctx, game).Execute()

Your card NFTs for a Crossly TCG game: each copy, its mint status, asset id and whether it is soulbound (starter).

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
    game := "game_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.CrosslyWalletApi.ListTcgNfts(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CrosslyWalletApi.ListTcgNfts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTcgNfts`: ListTcgNftsResponse
    fmt.Fprintf(os.Stdout, "Response from `CrosslyWalletApi.ListTcgNfts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTcgNftsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListTcgNftsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

