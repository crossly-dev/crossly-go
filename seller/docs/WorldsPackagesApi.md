# \WorldsPackagesApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorldExperiencePlacePackage**](WorldsPackagesApi.md#CreateWorldExperiencePlacePackage) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/packages | Insert a copy of a package (default: the newest version, auto-update on).
[**CreateWorldExperiencePlacePackageUpdate**](WorldsPackagesApi.md#CreateWorldExperiencePlacePackageUpdate) | **Post** /v1/worlds/experiences/{experienceId}/places/{placeId}/packages/{nodeId}/update | Update a copy to a version (default newest). A copy edited in the place needs force: true (its edits are replaced).
[**CreateWorldPackage**](WorldsPackagesApi.md#CreateWorldPackage) | **Post** /v1/worlds/packages | Create a package (v1) from one node (place JSON) or from a node in a place — which then becomes copy v1, auto-updating.
[**CreateWorldPackageToolbox**](WorldsPackagesApi.md#CreateWorldPackageToolbox) | **Post** /v1/worlds/packages/{packageId}/toolbox | List the latest version in the Toolbox as an asset of kind package (Toolbox review and pricing rules apply).
[**CreateWorldPackageVersion**](WorldsPackagesApi.md#CreateWorldPackageVersion) | **Post** /v1/worlds/packages/{packageId}/versions | Publish a new version; every auto-updating, unmodified copy in your places is updated through its co-editing session. Returns what happened per place.
[**DeleteWorldPackage**](WorldsPackagesApi.md#DeleteWorldPackage) | **Delete** /v1/worlds/packages/{packageId} | Archive a package. Copies in places keep working; they just stop receiving updates.
[**GetWorldPackage**](WorldsPackagesApi.md#GetWorldPackage) | **Get** /v1/worlds/packages/{packageId} | A package with its versions.
[**GetWorldPackageVersion**](WorldsPackagesApi.md#GetWorldPackageVersion) | **Get** /v1/worlds/packages/{packageId}/versions/{number} | One version, with its payload (one root node in place JSON).
[**ListWorldExperiencePlacePackages**](WorldsPackagesApi.md#ListWorldExperiencePlacePackages) | **Get** /v1/worlds/experiences/{experienceId}/places/{placeId}/packages | Package copies in a place: version, newest version, modified, auto-update.
[**ListWorldPackages**](WorldsPackagesApi.md#ListWorldPackages) | **Get** /v1/worlds/packages | Your package library: reusable models and modules, newest first, with how many places hold copies.
[**UpdateWorldExperiencePlacePackage**](WorldsPackagesApi.md#UpdateWorldExperiencePlacePackage) | **Patch** /v1/worlds/experiences/{experienceId}/places/{placeId}/packages/{nodeId} | Turn auto-update on or off for one copy.
[**UpdateWorldPackage**](WorldsPackagesApi.md#UpdateWorldPackage) | **Patch** /v1/worlds/packages/{packageId} | Rename a package or change its description.



## CreateWorldExperiencePlacePackage

> CreateWorldExperiencePlacePackageResponse CreateWorldExperiencePlacePackage(ctx, experienceId, placeId).Execute()

Insert a copy of a package (default: the newest version, auto-update on).

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
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsPackagesApi.CreateWorldExperiencePlacePackage(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsPackagesApi.CreateWorldExperiencePlacePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlacePackage`: CreateWorldExperiencePlacePackageResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsPackagesApi.CreateWorldExperiencePlacePackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlacePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperiencePlacePackageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldExperiencePlacePackageUpdate

> CreateWorldExperiencePlacePackageUpdateResponse CreateWorldExperiencePlacePackageUpdate(ctx, experienceId, placeId, nodeId).Execute()

Update a copy to a version (default newest). A copy edited in the place needs force: true (its edits are replaced).

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
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 
    nodeId := "nodeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsPackagesApi.CreateWorldExperiencePlacePackageUpdate(context.Background(), experienceId, placeId, nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsPackagesApi.CreateWorldExperiencePlacePackageUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperiencePlacePackageUpdate`: CreateWorldExperiencePlacePackageUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsPackagesApi.CreateWorldExperiencePlacePackageUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperiencePlacePackageUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**CreateWorldExperiencePlacePackageUpdateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldPackage

> GetWorldPackageResponse CreateWorldPackage(ctx).Execute()

Create a package (v1) from one node (place JSON) or from a node in a place — which then becomes copy v1, auto-updating.

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
    resp, r, err := apiClient.WorldsPackagesApi.CreateWorldPackage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsPackagesApi.CreateWorldPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldPackage`: GetWorldPackageResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsPackagesApi.CreateWorldPackage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldPackageRequest struct via the builder pattern


### Return type

**GetWorldPackageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldPackageToolbox

> CreateWorldPackageToolboxResponse CreateWorldPackageToolbox(ctx, packageId).Execute()

List the latest version in the Toolbox as an asset of kind package (Toolbox review and pricing rules apply).

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
    packageId := "packageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsPackagesApi.CreateWorldPackageToolbox(context.Background(), packageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsPackagesApi.CreateWorldPackageToolbox``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldPackageToolbox`: CreateWorldPackageToolboxResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsPackagesApi.CreateWorldPackageToolbox`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldPackageToolboxRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldPackageToolboxResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorldPackageVersion

> CreateWorldPackageVersionResponse CreateWorldPackageVersion(ctx, packageId).Execute()

Publish a new version; every auto-updating, unmodified copy in your places is updated through its co-editing session. Returns what happened per place.

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
    packageId := "packageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsPackagesApi.CreateWorldPackageVersion(context.Background(), packageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsPackagesApi.CreateWorldPackageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldPackageVersion`: CreateWorldPackageVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsPackagesApi.CreateWorldPackageVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldPackageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldPackageVersionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorldPackage

> DeleteWorldPackageResponse DeleteWorldPackage(ctx, packageId).Execute()

Archive a package. Copies in places keep working; they just stop receiving updates.

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
    packageId := "packageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsPackagesApi.DeleteWorldPackage(context.Background(), packageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsPackagesApi.DeleteWorldPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorldPackage`: DeleteWorldPackageResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsPackagesApi.DeleteWorldPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorldPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteWorldPackageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldPackage

> GetWorldPackageResponse GetWorldPackage(ctx, packageId).Execute()

A package with its versions.

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
    packageId := "packageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsPackagesApi.GetWorldPackage(context.Background(), packageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsPackagesApi.GetWorldPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldPackage`: GetWorldPackageResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsPackagesApi.GetWorldPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldPackageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldPackageVersion

> GetWorldPackageVersionResponse GetWorldPackageVersion(ctx, packageId, number).Execute()

One version, with its payload (one root node in place JSON).

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
    packageId := "packageId_example" // string | 
    number := "number_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsPackagesApi.GetWorldPackageVersion(context.Background(), packageId, number).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsPackagesApi.GetWorldPackageVersion``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldPackageVersion`: GetWorldPackageVersionResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsPackagesApi.GetWorldPackageVersion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | **string** |  | 
**number** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldPackageVersionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorldPackageVersionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldExperiencePlacePackages

> ListWorldExperiencePlacePackagesResponse ListWorldExperiencePlacePackages(ctx, experienceId, placeId).Execute()

Package copies in a place: version, newest version, modified, auto-update.

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
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsPackagesApi.ListWorldExperiencePlacePackages(context.Background(), experienceId, placeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsPackagesApi.ListWorldExperiencePlacePackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldExperiencePlacePackages`: ListWorldExperiencePlacePackagesResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsPackagesApi.ListWorldExperiencePlacePackages`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldExperiencePlacePackagesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**ListWorldExperiencePlacePackagesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorldPackages

> ListWorldPackagesResponse ListWorldPackages(ctx).Execute()

Your package library: reusable models and modules, newest first, with how many places hold copies.

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
    resp, r, err := apiClient.WorldsPackagesApi.ListWorldPackages(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsPackagesApi.ListWorldPackages``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorldPackages`: ListWorldPackagesResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsPackagesApi.ListWorldPackages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorldPackagesRequest struct via the builder pattern


### Return type

**ListWorldPackagesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperiencePlacePackage

> UpdateWorldExperiencePlacePackageResponse UpdateWorldExperiencePlacePackage(ctx, experienceId, placeId, nodeId).Execute()

Turn auto-update on or off for one copy.

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
    experienceId := "experienceId_example" // string | 
    placeId := "placeId_example" // string | 
    nodeId := "nodeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsPackagesApi.UpdateWorldExperiencePlacePackage(context.Background(), experienceId, placeId, nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsPackagesApi.UpdateWorldExperiencePlacePackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperiencePlacePackage`: UpdateWorldExperiencePlacePackageResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsPackagesApi.UpdateWorldExperiencePlacePackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**placeId** | **string** |  | 
**nodeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperiencePlacePackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

**UpdateWorldExperiencePlacePackageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldPackage

> DeleteWorldPackageResponse UpdateWorldPackage(ctx, packageId).Execute()

Rename a package or change its description.

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
    packageId := "packageId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsPackagesApi.UpdateWorldPackage(context.Background(), packageId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsPackagesApi.UpdateWorldPackage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldPackage`: DeleteWorldPackageResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsPackagesApi.UpdateWorldPackage`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**packageId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldPackageRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteWorldPackageResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

