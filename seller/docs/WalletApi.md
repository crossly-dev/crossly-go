# \WalletApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteWalletDefault**](WalletApi.md#DeleteWalletDefault) | **Delete** /v1/wallet/defaults | Clear a default (the purpose falls back to your only wallet, or none) or one override (&#x60;refKind&#x60; + &#x60;refKey&#x60;).
[**ListWalletDefaults**](WalletApi.md#ListWalletDefaults) | **Get** /v1/wallet/defaults | Your default wallet per purpose (royalties, deploys, drops, game rewards, staking, CBX withdrawals), with per-collection and per-game overrides and what each resolves to today.
[**UpdateWalletDefault**](WalletApi.md#UpdateWalletDefault) | **Put** /v1/wallet/defaults | Set the default wallet for a purpose, or an override for one collection (&#x60;ref: { kind: \&quot;collection\&quot;, chain, contract }&#x60;) or game (&#x60;ref: { kind: \&quot;game\&quot;, experienceId }&#x60;). The wallet must be yours, verified, and on a network the purpose accepts.



## DeleteWalletDefault

> WalletResponse DeleteWalletDefault(ctx).Execute()

Clear a default (the purpose falls back to your only wallet, or none) or one override (`refKind` + `refKey`).

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
    resp, r, err := apiClient.WalletApi.DeleteWalletDefault(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.DeleteWalletDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWalletDefault`: WalletResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.DeleteWalletDefault`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWalletDefaultRequest struct via the builder pattern


### Return type

**WalletResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWalletDefaults

> WalletResponse ListWalletDefaults(ctx).Execute()

Your default wallet per purpose (royalties, deploys, drops, game rewards, staking, CBX withdrawals), with per-collection and per-game overrides and what each resolves to today.

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
    resp, r, err := apiClient.WalletApi.ListWalletDefaults(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.ListWalletDefaults``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWalletDefaults`: WalletResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.ListWalletDefaults`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWalletDefaultsRequest struct via the builder pattern


### Return type

**WalletResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWalletDefault

> WalletResponse UpdateWalletDefault(ctx).Execute()

Set the default wallet for a purpose, or an override for one collection (`ref: { kind: \"collection\", chain, contract }`) or game (`ref: { kind: \"game\", experienceId }`). The wallet must be yours, verified, and on a network the purpose accepts.

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
    resp, r, err := apiClient.WalletApi.UpdateWalletDefault(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WalletApi.UpdateWalletDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWalletDefault`: WalletResponse
    fmt.Fprintf(os.Stdout, "Response from `WalletApi.UpdateWalletDefault`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWalletDefaultRequest struct via the builder pattern


### Return type

**WalletResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

