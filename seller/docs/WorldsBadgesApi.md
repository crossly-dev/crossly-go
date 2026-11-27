# \WorldsBadgesApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorldExperienceBadge**](WorldsBadgesApi.md#CreateWorldExperienceBadge) | **Post** /v1/worlds/experiences/{experienceId}/badges | Add a badge (at most 100 per game). name 1-50 characters, description up to 1000, iconSha256 &#x3D; the sha256 a png upload answered (POST /v1/world/interactive/assets, kind png); a new icon is shown once a person approved it.
[**GetWorldBadgeMine**](WorldsBadgesApi.md#GetWorldBadgeMine) | **Get** /v1/worlds/badges/mine | Every game badge you hold, newest first, and whether your person page shows them.
[**GetWorldExperienceBadge**](WorldsBadgesApi.md#GetWorldExperienceBadge) | **Get** /v1/worlds/experiences/{experienceId}/badges | A game&#39;s badges, oldest first: name, description, icon and its review state, on/off, how many players hold each; plus the player count (the rarity denominator) and the limits. Scripts award them with BadgeService.AwardBadge(player, badgeId).
[**GetWorldGameBadge**](WorldsBadgesApi.md#GetWorldGameBadge) | **Get** /v1/worlds/games/{game}/badges | A Crossly Carnival game&#39;s badges as players see them: name, description, approved icon, how many hold it, rarity (% of the game&#39;s players) and, for you, when you earned it. :game is a Worlds experience id or key.
[**GetWorldPlayerBadge**](WorldsBadgesApi.md#GetWorldPlayerBadge) | **Get** /v1/worlds/players/{username}/badges | The game badges a person shows on their person page (only when they switched it on; Mature and non-public games never appear).
[**UpdateWorldBadgeProfile**](WorldsBadgesApi.md#UpdateWorldBadgeProfile) | **Put** /v1/worlds/badges/profile | Show (or hide) the game badges you hold on your person page. Off by default.
[**UpdateWorldExperienceBadge**](WorldsBadgesApi.md#UpdateWorldExperienceBadge) | **Patch** /v1/worlds/experiences/{experienceId}/badges/{badgeId} | Edit a badge: name, description, icon (null removes it), enabled. Badges are never deleted (players keep what they earned); turn one off to hide it and stop AwardBadge.



## CreateWorldExperienceBadge

> CreateWorldExperienceBadgeResponse CreateWorldExperienceBadge(ctx, experienceId).Execute()

Add a badge (at most 100 per game). name 1-50 characters, description up to 1000, iconSha256 = the sha256 a png upload answered (POST /v1/world/interactive/assets, kind png); a new icon is shown once a person approved it.

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsBadgesApi.CreateWorldExperienceBadge(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsBadgesApi.CreateWorldExperienceBadge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorldExperienceBadge`: CreateWorldExperienceBadgeResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsBadgesApi.CreateWorldExperienceBadge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorldExperienceBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorldExperienceBadgeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldBadgeMine

> GetWorldBadgeMineResponse GetWorldBadgeMine(ctx).Execute()

Every game badge you hold, newest first, and whether your person page shows them.

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
    resp, r, err := apiClient.WorldsBadgesApi.GetWorldBadgeMine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsBadgesApi.GetWorldBadgeMine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldBadgeMine`: GetWorldBadgeMineResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsBadgesApi.GetWorldBadgeMine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldBadgeMineRequest struct via the builder pattern


### Return type

**GetWorldBadgeMineResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldExperienceBadge

> GetWorldExperienceBadgeResponse GetWorldExperienceBadge(ctx, experienceId).Execute()

A game's badges, oldest first: name, description, icon and its review state, on/off, how many players hold each; plus the player count (the rarity denominator) and the limits. Scripts award them with BadgeService.AwardBadge(player, badgeId).

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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsBadgesApi.GetWorldExperienceBadge(context.Background(), experienceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsBadgesApi.GetWorldExperienceBadge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldExperienceBadge`: GetWorldExperienceBadgeResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsBadgesApi.GetWorldExperienceBadge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldExperienceBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldExperienceBadgeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldGameBadge

> GetWorldGameBadgeResponse GetWorldGameBadge(ctx, game).Execute()

A Crossly Carnival game's badges as players see them: name, description, approved icon, how many hold it, rarity (% of the game's players) and, for you, when you earned it. :game is a Worlds experience id or key.

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
    resp, r, err := apiClient.WorldsBadgesApi.GetWorldGameBadge(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsBadgesApi.GetWorldGameBadge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldGameBadge`: GetWorldGameBadgeResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsBadgesApi.GetWorldGameBadge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldGameBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldGameBadgeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorldPlayerBadge

> GetWorldPlayerBadgeResponse GetWorldPlayerBadge(ctx, username).Execute()

The game badges a person shows on their person page (only when they switched it on; Mature and non-public games never appear).

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
    username := "username_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsBadgesApi.GetWorldPlayerBadge(context.Background(), username).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsBadgesApi.GetWorldPlayerBadge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorldPlayerBadge`: GetWorldPlayerBadgeResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsBadgesApi.GetWorldPlayerBadge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**username** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorldPlayerBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorldPlayerBadgeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldBadgeProfile

> UpdateWorldBadgeProfileResponse UpdateWorldBadgeProfile(ctx).Execute()

Show (or hide) the game badges you hold on your person page. Off by default.

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
    resp, r, err := apiClient.WorldsBadgesApi.UpdateWorldBadgeProfile(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsBadgesApi.UpdateWorldBadgeProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldBadgeProfile`: UpdateWorldBadgeProfileResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsBadgesApi.UpdateWorldBadgeProfile`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldBadgeProfileRequest struct via the builder pattern


### Return type

**UpdateWorldBadgeProfileResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorldExperienceBadge

> CreateWorldExperienceBadgeResponse UpdateWorldExperienceBadge(ctx, experienceId, badgeId).Execute()

Edit a badge: name, description, icon (null removes it), enabled. Badges are never deleted (players keep what they earned); turn one off to hide it and stop AwardBadge.

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
    badgeId := "badgeId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorldsBadgesApi.UpdateWorldExperienceBadge(context.Background(), experienceId, badgeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorldsBadgesApi.UpdateWorldExperienceBadge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorldExperienceBadge`: CreateWorldExperienceBadgeResponse
    fmt.Fprintf(os.Stdout, "Response from `WorldsBadgesApi.UpdateWorldExperienceBadge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**experienceId** | **string** |  | 
**badgeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorldExperienceBadgeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateWorldExperienceBadgeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

