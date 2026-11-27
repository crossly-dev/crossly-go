# \TCGCreatorsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTcgCreatorGame**](TCGCreatorsApi.md#CreateTcgCreatorGame) | **Post** /v1/tcg/creator/games | Submit a card-game package version for review: game.json, cards.json, rules/game.js, thumb.png, art/_*.png as { path, base64 }. Checked, conformance-run, art quarantined, then reviewed by a person.
[**CreateTcgCreatorGameCheck**](TCGCreatorsApi.md#CreateTcgCreatorGameCheck) | **Post** /v1/tcg/creator/games/check | Check a card-game package exactly as a submission is checked -- the static report and bot games played against your rules in the sandbox -- storing nothing. Files as { path, base64 }.
[**CreateTcgCreatorGamePackOpen**](TCGCreatorsApi.md#CreateTcgCreatorGamePackOpen) | **Post** /v1/tcg/creator/games/{id}/packs/{packId}/open | Open a booster you hold: the game&#39;s rules pick the pulls, the pack is used up, and one COSM copy per card joins your collection (logged with its seed).
[**CreateTcgCreatorGameStarter**](TCGCreatorsApi.md#CreateTcgCreatorGameStarter) | **Post** /v1/tcg/creator/games/{id}/starter | Claim a creator game&#39;s free starter deck by its key (once per account; account-bound, never tradeable).
[**CreateTcgCreatorGameVersionWithdraw**](TCGCreatorsApi.md#CreateTcgCreatorGameVersionWithdraw) | **Post** /v1/tcg/creator/games/{id}/versions/{v}/withdraw | Withdraw one of your versions that is still waiting for review.
[**GetTcgCreatorGame**](TCGCreatorsApi.md#GetTcgCreatorGame) | **Get** /v1/tcg/creator/games/{id} | One live creator card game: its set, every card (rules text, rarity, art URL, game data) and its COSM prices.
[**GetTcgCreatorGameCollection**](TCGCreatorsApi.md#GetTcgCreatorGameCollection) | **Get** /v1/tcg/creator/games/{id}/collection | Your playable copies of a creator game: card id -&gt; copies (free starter plus every COSM copy you hold).
[**GetTcgCreatorGamePopularity**](TCGCreatorsApi.md#GetTcgCreatorGamePopularity) | **Get** /v1/tcg/creator/games/{id}/popularity | The set&#39;s progress toward going on chain: distinct collectors, packs opened and resales against the bar; when it passes, every holder may mint their cards as NFTs.
[**GetTcgCreatorGameShelf**](TCGCreatorsApi.md#GetTcgCreatorGameShelf) | **Get** /v1/tcg/creator/games/{id}/shelf | The game&#39;s COSM shelf: the booster and every single, each with its COSM design id (buy via /api/v1/cosm/cosmetics/:id/buy), price, edition cap and copies issued.
[**ListTcgCreatorGamePacks**](TCGCreatorsApi.md#ListTcgCreatorGamePacks) | **Get** /v1/tcg/creator/games/{id}/packs | Your unopened boosters of a creator game.
[**ListTcgCreatorGames**](TCGCreatorsApi.md#ListTcgCreatorGames) | **Get** /v1/tcg/creator/games | Live creator card games: name, set, description, typical match length, frame colours, thumbnail.
[**ListTcgCreatorMine**](TCGCreatorsApi.md#ListTcgCreatorMine) | **Get** /v1/tcg/creator/mine | Your creator card games and every version with its review state and note.



## CreateTcgCreatorGame

> CreateTcgCreatorGameResponse CreateTcgCreatorGame(ctx).Execute()

Submit a card-game package version for review: game.json, cards.json, rules/game.js, thumb.png, art/_*.png as { path, base64 }. Checked, conformance-run, art quarantined, then reviewed by a person.

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
    resp, r, err := apiClient.TCGCreatorsApi.CreateTcgCreatorGame(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGCreatorsApi.CreateTcgCreatorGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgCreatorGame`: CreateTcgCreatorGameResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGCreatorsApi.CreateTcgCreatorGame`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgCreatorGameRequest struct via the builder pattern


### Return type

**CreateTcgCreatorGameResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgCreatorGameCheck

> CreateTcgCreatorGameCheckResponse CreateTcgCreatorGameCheck(ctx).Execute()

Check a card-game package exactly as a submission is checked -- the static report and bot games played against your rules in the sandbox -- storing nothing. Files as { path, base64 }.

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
    resp, r, err := apiClient.TCGCreatorsApi.CreateTcgCreatorGameCheck(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGCreatorsApi.CreateTcgCreatorGameCheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgCreatorGameCheck`: CreateTcgCreatorGameCheckResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGCreatorsApi.CreateTcgCreatorGameCheck`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgCreatorGameCheckRequest struct via the builder pattern


### Return type

**CreateTcgCreatorGameCheckResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgCreatorGamePackOpen

> CreateTcgCreatorGamePackOpenResponse CreateTcgCreatorGamePackOpen(ctx, id, packId).Execute()

Open a booster you hold: the game's rules pick the pulls, the pack is used up, and one COSM copy per card joins your collection (logged with its seed).

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
    packId := "packId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TCGCreatorsApi.CreateTcgCreatorGamePackOpen(context.Background(), id, packId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGCreatorsApi.CreateTcgCreatorGamePackOpen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgCreatorGamePackOpen`: CreateTcgCreatorGamePackOpenResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGCreatorsApi.CreateTcgCreatorGamePackOpen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**packId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgCreatorGamePackOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**CreateTcgCreatorGamePackOpenResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgCreatorGameStarter

> CreateTcgStarterResponse CreateTcgCreatorGameStarter(ctx, id).Execute()

Claim a creator game's free starter deck by its key (once per account; account-bound, never tradeable).

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
    resp, r, err := apiClient.TCGCreatorsApi.CreateTcgCreatorGameStarter(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGCreatorsApi.CreateTcgCreatorGameStarter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgCreatorGameStarter`: CreateTcgStarterResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGCreatorsApi.CreateTcgCreatorGameStarter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgCreatorGameStarterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateTcgStarterResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgCreatorGameVersionWithdraw

> CreateTcgCreatorGameVersionWithdraw(ctx, id, v).Execute()

Withdraw one of your versions that is still waiting for review.

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
    v := "v_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TCGCreatorsApi.CreateTcgCreatorGameVersionWithdraw(context.Background(), id, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGCreatorsApi.CreateTcgCreatorGameVersionWithdraw``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**v** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgCreatorGameVersionWithdrawRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTcgCreatorGame

> GetTcgCreatorGameResponse GetTcgCreatorGame(ctx, id).Execute()

One live creator card game: its set, every card (rules text, rarity, art URL, game data) and its COSM prices.

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
    resp, r, err := apiClient.TCGCreatorsApi.GetTcgCreatorGame(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGCreatorsApi.GetTcgCreatorGame``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTcgCreatorGame`: GetTcgCreatorGameResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGCreatorsApi.GetTcgCreatorGame`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcgCreatorGameRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetTcgCreatorGameResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTcgCreatorGameCollection

> GetTcgCreatorGameCollectionResponse GetTcgCreatorGameCollection(ctx, id).Execute()

Your playable copies of a creator game: card id -> copies (free starter plus every COSM copy you hold).

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
    resp, r, err := apiClient.TCGCreatorsApi.GetTcgCreatorGameCollection(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGCreatorsApi.GetTcgCreatorGameCollection``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTcgCreatorGameCollection`: GetTcgCreatorGameCollectionResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGCreatorsApi.GetTcgCreatorGameCollection`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcgCreatorGameCollectionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetTcgCreatorGameCollectionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTcgCreatorGamePopularity

> GetTcgCreatorGamePopularityResponse GetTcgCreatorGamePopularity(ctx, id).Execute()

The set's progress toward going on chain: distinct collectors, packs opened and resales against the bar; when it passes, every holder may mint their cards as NFTs.

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
    resp, r, err := apiClient.TCGCreatorsApi.GetTcgCreatorGamePopularity(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGCreatorsApi.GetTcgCreatorGamePopularity``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTcgCreatorGamePopularity`: GetTcgCreatorGamePopularityResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGCreatorsApi.GetTcgCreatorGamePopularity`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcgCreatorGamePopularityRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetTcgCreatorGamePopularityResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTcgCreatorGameShelf

> GetTcgCreatorGameShelfResponse GetTcgCreatorGameShelf(ctx, id).Execute()

The game's COSM shelf: the booster and every single, each with its COSM design id (buy via /api/v1/cosm/cosmetics/:id/buy), price, edition cap and copies issued.

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
    resp, r, err := apiClient.TCGCreatorsApi.GetTcgCreatorGameShelf(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGCreatorsApi.GetTcgCreatorGameShelf``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTcgCreatorGameShelf`: GetTcgCreatorGameShelfResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGCreatorsApi.GetTcgCreatorGameShelf`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcgCreatorGameShelfRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetTcgCreatorGameShelfResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListTcgCreatorGamePacks

> ListTcgCreatorGamePacksResponse ListTcgCreatorGamePacks(ctx, id).Execute()

Your unopened boosters of a creator game.

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
    resp, r, err := apiClient.TCGCreatorsApi.ListTcgCreatorGamePacks(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGCreatorsApi.ListTcgCreatorGamePacks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTcgCreatorGamePacks`: ListTcgCreatorGamePacksResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGCreatorsApi.ListTcgCreatorGamePacks`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTcgCreatorGamePacksRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListTcgCreatorGamePacksResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListTcgCreatorGames

> ListTcgCreatorGamesResponse ListTcgCreatorGames(ctx).Execute()

Live creator card games: name, set, description, typical match length, frame colours, thumbnail.

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
    resp, r, err := apiClient.TCGCreatorsApi.ListTcgCreatorGames(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGCreatorsApi.ListTcgCreatorGames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTcgCreatorGames`: ListTcgCreatorGamesResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGCreatorsApi.ListTcgCreatorGames`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTcgCreatorGamesRequest struct via the builder pattern


### Return type

**ListTcgCreatorGamesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListTcgCreatorMine

> ListTcgCreatorMineResponse ListTcgCreatorMine(ctx).Execute()

Your creator card games and every version with its review state and note.

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
    resp, r, err := apiClient.TCGCreatorsApi.ListTcgCreatorMine(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGCreatorsApi.ListTcgCreatorMine``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTcgCreatorMine`: ListTcgCreatorMineResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGCreatorsApi.ListTcgCreatorMine`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTcgCreatorMineRequest struct via the builder pattern


### Return type

**ListTcgCreatorMineResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

