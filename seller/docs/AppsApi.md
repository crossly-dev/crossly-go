# \AppsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetApp**](AppsApi.md#GetApp) | **Get** /v1/apps/{clientId} | Get one listed app by client id or profile slug. 404 if it is not in the directory.
[**GetAppEarning**](AppsApi.md#GetAppEarning) | **Get** /v1/apps/{id}/earnings | Your OAuth app&#39;s revenue-share earnings: summary + ledger. Owner only.
[**GetAppListingAllowance**](AppsApi.md#GetAppListingAllowance) | **Get** /v1/apps/{id}/listing-allowance | Your OAuth app&#39;s listing allowance (read-only): whether staff let sellers list above their plan through it, its limitMode (separate pool or shared with the plan), limitPeriod (total or monthly), the per-seller limit and the platform scope. Owner only.
[**ListApps**](AppsApi.md#ListApps) | **Get** /v1/apps | Browse the public Crossly app directories: seller apps, buyer apps or both (search, category, sort, featured, paging).



## GetApp

> ListAppsItem GetApp(ctx, clientId).Execute()

Get one listed app by client id or profile slug. 404 if it is not in the directory.

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
    clientId := "clientId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.GetApp(context.Background(), clientId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.GetApp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApp`: ListAppsItem
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.GetApp`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**clientId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListAppsItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAppEarning

> GetAppEarningResponse GetAppEarning(ctx, id).Execute()

Your OAuth app's revenue-share earnings: summary + ledger. Owner only.

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
    resp, r, err := apiClient.AppsApi.GetAppEarning(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.GetAppEarning``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppEarning`: GetAppEarningResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.GetAppEarning`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppEarningRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetAppEarningResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetAppListingAllowance

> GetAppListingAllowanceResponse GetAppListingAllowance(ctx, id).Execute()

Your OAuth app's listing allowance (read-only): whether staff let sellers list above their plan through it, its limitMode (separate pool or shared with the plan), limitPeriod (total or monthly), the per-seller limit and the platform scope. Owner only.

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
    resp, r, err := apiClient.AppsApi.GetAppListingAllowance(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.GetAppListingAllowance``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAppListingAllowance`: GetAppListingAllowanceResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.GetAppListingAllowance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAppListingAllowanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetAppListingAllowanceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListApps

> ListAppsResponse ListApps(ctx).Directory(directory).Sort(sort).Featured(featured).Category(category).Q(q).Limit(limit).Cursor(cursor).Execute()

Browse the public Crossly app directories: seller apps, buyer apps or both (search, category, sort, featured, paging).

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
    directory := "directory_example" // string |  (optional)
    sort := "sort_example" // string |  (optional)
    featured := "featured_example" // string |  (optional)
    category := "category_example" // string |  (optional)
    q := "q_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional)
    cursor := "cursor_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AppsApi.ListApps(context.Background()).Directory(directory).Sort(sort).Featured(featured).Category(category).Q(q).Limit(limit).Cursor(cursor).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AppsApi.ListApps``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListApps`: ListAppsResponse
    fmt.Fprintf(os.Stdout, "Response from `AppsApi.ListApps`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAppsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **directory** | **string** |  | 
 **sort** | **string** |  | 
 **featured** | **string** |  | 
 **category** | **string** |  | 
 **q** | **string** |  | 
 **limit** | **int32** |  | 
 **cursor** | **string** |  | 

### Return type

**ListAppsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

