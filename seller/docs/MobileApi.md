# \MobileApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMobilePushTest**](MobileApi.md#CreateMobilePushTest) | **Post** /v1/mobile/push-test | Fire a no-op test push to this user&#39;s devices.
[**CreateMobilePushToken**](MobileApi.md#CreateMobilePushToken) | **Post** /v1/mobile/push-token | Register an Expo push token for this user.
[**DeleteMobilePushToken**](MobileApi.md#DeleteMobilePushToken) | **Delete** /v1/mobile/push-tokens | Clear ALL registered push tokens for this user.
[**ListMobilePushTokens**](MobileApi.md#ListMobilePushTokens) | **Get** /v1/mobile/push-tokens | List registered Expo push tokens (masked).



## CreateMobilePushTest

> CreateMobilePushTestResponse CreateMobilePushTest(ctx).Execute()

Fire a no-op test push to this user's devices.

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
    resp, r, err := apiClient.MobileApi.CreateMobilePushTest(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileApi.CreateMobilePushTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMobilePushTest`: CreateMobilePushTestResponse
    fmt.Fprintf(os.Stdout, "Response from `MobileApi.CreateMobilePushTest`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMobilePushTestRequest struct via the builder pattern


### Return type

[**CreateMobilePushTestResponse**](CreateMobilePushTestResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMobilePushToken

> CreateMobilePushTokenResponse CreateMobilePushToken(ctx).Execute()

Register an Expo push token for this user.

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
    resp, r, err := apiClient.MobileApi.CreateMobilePushToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileApi.CreateMobilePushToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateMobilePushToken`: CreateMobilePushTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `MobileApi.CreateMobilePushToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMobilePushTokenRequest struct via the builder pattern


### Return type

[**CreateMobilePushTokenResponse**](CreateMobilePushTokenResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMobilePushToken

> DeleteMobilePushTokenResponse DeleteMobilePushToken(ctx).Execute()

Clear ALL registered push tokens for this user.

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
    resp, r, err := apiClient.MobileApi.DeleteMobilePushToken(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileApi.DeleteMobilePushToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteMobilePushToken`: DeleteMobilePushTokenResponse
    fmt.Fprintf(os.Stdout, "Response from `MobileApi.DeleteMobilePushToken`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMobilePushTokenRequest struct via the builder pattern


### Return type

[**DeleteMobilePushTokenResponse**](DeleteMobilePushTokenResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMobilePushTokens

> V1List ListMobilePushTokens(ctx).Execute()

List registered Expo push tokens (masked).

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
    resp, r, err := apiClient.MobileApi.ListMobilePushTokens(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MobileApi.ListMobilePushTokens``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMobilePushTokens`: V1List
    fmt.Fprintf(os.Stdout, "Response from `MobileApi.ListMobilePushTokens`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListMobilePushTokensRequest struct via the builder pattern


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

