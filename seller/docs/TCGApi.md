# \TCGApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTcgChallenge**](TCGApi.md#CreateTcgChallenge) | **Post** /v1/tcg/{game}/challenges | Open a 1v1 challenge with one of your decks; another player joins it.
[**CreateTcgDeck**](TCGApi.md#CreateTcgDeck) | **Post** /v1/tcg/{game}/decks | Create (no id) or update a deck: its lead (Crew Boss, Mastermind or Keeper) and 40 cards you own, checked against the deck rules of its game.
[**CreateTcgMatchAct**](TCGApi.md#CreateTcgMatchAct) | **Post** /v1/tcg/matches/{id}/act | Make one move: { action } as listed in the match&#39;s legal moves. The engine checks it.
[**CreateTcgMatchJoin**](TCGApi.md#CreateTcgMatchJoin) | **Post** /v1/tcg/matches/{id}/join | Join an open 1v1 challenge with one of your decks.
[**CreateTcgPackOpen**](TCGApi.md#CreateTcgPackOpen) | **Post** /v1/tcg/packs/{id}/open | Open one of your earned packs: its eight cards join your collection.
[**CreateTcgPractice**](TCGApi.md#CreateTcgPractice) | **Post** /v1/tcg/{game}/practice | Start a practice match against the house bot with one of your decks.
[**CreateTcgQueue**](TCGApi.md#CreateTcgQueue) | **Post** /v1/tcg/{game}/queue | Join the ranked queue with an owned, legal deck (free). Returns the match at once when an opponent fits.
[**CreateTcgStarter**](TCGApi.md#CreateTcgStarter) | **Post** /v1/tcg/{game}/starter | Claim the free starter deck for a game, choosing its faction (once per account; account-bound, untradeable).
[**CreateTcgTable**](TCGApi.md#CreateTcgTable) | **Post** /v1/tcg/{game}/table | Sit at a table with your free starter deck (claimed now if you never did: starter names it, else one is picked for you): joins the oldest open 1v1 challenge, or opens one; already waiting, your own seat back.
[**CreateTcgTournament**](TCGApi.md#CreateTcgTournament) | **Post** /v1/tcg/{game}/tournaments | Host a Swiss tournament: name, optional top cut (4 or 8), round length, player cap and start time.
[**CreateTcgTournamentDrop**](TCGApi.md#CreateTcgTournamentDrop) | **Post** /v1/tcg/tournaments/{id}/drop | Leave a tournament before it starts, or drop from it (your results stand).
[**CreateTcgTournamentRegister**](TCGApi.md#CreateTcgTournamentRegister) | **Post** /v1/tcg/tournaments/{id}/register | Enter a tournament with an owned, legal deck before it starts.
[**CreateTcgTournamentStart**](TCGApi.md#CreateTcgTournamentStart) | **Post** /v1/tcg/tournaments/{id}/start | As the host, start the tournament: Swiss rounds are sized to the field and round 1 is dealt.
[**DeleteTcgDeck**](TCGApi.md#DeleteTcgDeck) | **Delete** /v1/tcg/decks/{id} | Delete one of your decks.
[**DeleteTcgQueue**](TCGApi.md#DeleteTcgQueue) | **Delete** /v1/tcg/{game}/queue | Leave the ranked queue.
[**GetTcgCreatorGamePrint**](TCGApi.md#GetTcgCreatorGamePrint) | **Get** /v1/tcg/creator/games/{id}/print | Print file links for your own creator card game (owner only): the set&#39;s print file, home sheets and a proof, for its newest version.
[**GetTcgCreatorGamePrintByKind**](TCGApi.md#GetTcgCreatorGamePrintByKind) | **Get** /v1/tcg/creator/games/{id}/print/{kind} | A PDF of your own creator card game (owner only): kind set | sheet | proof; ?version&#x3D;, ?paper&#x3D;letter|a4, ?card&#x3D; for a proof. Binary (application/pdf).
[**GetTcgLadder**](TCGApi.md#GetTcgLadder) | **Get** /v1/tcg/{game}/ladder | The season&#39;s ranked leaderboard for a game: place, rating, tier and record (named only where the player chose to be).
[**GetTcgLadderMe**](TCGApi.md#GetTcgLadderMe) | **Get** /v1/tcg/{game}/ladder/me | Your ranked rating, tier, place and record this season, and whether you are queued.
[**GetTcgMatch**](TCGApi.md#GetTcgMatch) | **Get** /v1/tcg/{game}/matches | Your recent matches, and open challenges you could join.
[**GetTcgMatchById**](TCGApi.md#GetTcgMatchById) | **Get** /v1/tcg/matches/{id} | Your view of a match (no hidden information) and, on your turn, every legal move.
[**GetTcgMe**](TCGApi.md#GetTcgMe) | **Get** /v1/tcg/{game}/me | Your starter faction, collection (card id -&gt; copies), decks and earned packs.
[**GetTcgPrint**](TCGApi.md#GetTcgPrint) | **Get** /v1/tcg/{game}/print | Print files for one of Crossly&#39;s own card games: a card proof (template), 3 x 3 Letter/A4 home sheets with crop marks, and the set&#39;s print file (63 x 88 mm, 3 mm bleed).
[**GetTcgTournament**](TCGApi.md#GetTcgTournament) | **Get** /v1/tcg/tournaments/{id} | A tournament&#39;s standings (points, opponents&#39; and game win %), this round&#39;s pairings and your seat.
[**ListTcgCards**](TCGApi.md#ListTcgCards) | **Get** /v1/tcg/{game}/cards | The game&#39;s card pool: every card with its rules text, stats, rarity and collector number.
[**ListTcgGames**](TCGApi.md#ListTcgGames) | **Get** /v1/tcg/games | Crossly&#39;s own card games: id, name, seats and card count.
[**ListTcgTournaments**](TCGApi.md#ListTcgTournaments) | **Get** /v1/tcg/{game}/tournaments | A game&#39;s recent and upcoming Swiss tournaments (free to enter; prizes are earned packs).



## CreateTcgChallenge

> TcgResponse CreateTcgChallenge(ctx, game).Execute()

Open a 1v1 challenge with one of your decks; another player joins it.

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
    resp, r, err := apiClient.TCGApi.CreateTcgChallenge(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.CreateTcgChallenge``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgChallenge`: TcgResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.CreateTcgChallenge`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgChallengeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**TcgResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgDeck

> CreateTcgDeckResponse CreateTcgDeck(ctx, game).Execute()

Create (no id) or update a deck: its lead (Crew Boss, Mastermind or Keeper) and 40 cards you own, checked against the deck rules of its game.

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
    resp, r, err := apiClient.TCGApi.CreateTcgDeck(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.CreateTcgDeck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgDeck`: CreateTcgDeckResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.CreateTcgDeck`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgDeckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateTcgDeckResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgMatchAct

> TcgResponse CreateTcgMatchAct(ctx, id).Execute()

Make one move: { action } as listed in the match's legal moves. The engine checks it.

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
    resp, r, err := apiClient.TCGApi.CreateTcgMatchAct(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.CreateTcgMatchAct``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgMatchAct`: TcgResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.CreateTcgMatchAct`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgMatchActRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**TcgResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgMatchJoin

> TcgResponse CreateTcgMatchJoin(ctx, id).Execute()

Join an open 1v1 challenge with one of your decks.

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
    resp, r, err := apiClient.TCGApi.CreateTcgMatchJoin(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.CreateTcgMatchJoin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgMatchJoin`: TcgResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.CreateTcgMatchJoin`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgMatchJoinRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**TcgResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgPackOpen

> CreateTcgPackOpenResponse CreateTcgPackOpen(ctx, id).Execute()

Open one of your earned packs: its eight cards join your collection.

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
    resp, r, err := apiClient.TCGApi.CreateTcgPackOpen(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.CreateTcgPackOpen``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgPackOpen`: CreateTcgPackOpenResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.CreateTcgPackOpen`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgPackOpenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateTcgPackOpenResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgPractice

> TcgResponse CreateTcgPractice(ctx, game).Execute()

Start a practice match against the house bot with one of your decks.

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
    resp, r, err := apiClient.TCGApi.CreateTcgPractice(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.CreateTcgPractice``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgPractice`: TcgResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.CreateTcgPractice`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgPracticeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**TcgResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgQueue

> CreateTcgQueueResponse CreateTcgQueue(ctx, game).Execute()

Join the ranked queue with an owned, legal deck (free). Returns the match at once when an opponent fits.

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
    resp, r, err := apiClient.TCGApi.CreateTcgQueue(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.CreateTcgQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgQueue`: CreateTcgQueueResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.CreateTcgQueue`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgQueueRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateTcgQueueResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgStarter

> CreateTcgStarterResponse CreateTcgStarter(ctx, game).Execute()

Claim the free starter deck for a game, choosing its faction (once per account; account-bound, untradeable).

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
    resp, r, err := apiClient.TCGApi.CreateTcgStarter(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.CreateTcgStarter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgStarter`: CreateTcgStarterResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.CreateTcgStarter`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgStarterRequest struct via the builder pattern


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


## CreateTcgTable

> TcgResponse CreateTcgTable(ctx, game).Execute()

Sit at a table with your free starter deck (claimed now if you never did: starter names it, else one is picked for you): joins the oldest open 1v1 challenge, or opens one; already waiting, your own seat back.

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
    resp, r, err := apiClient.TCGApi.CreateTcgTable(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.CreateTcgTable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgTable`: TcgResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.CreateTcgTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**TcgResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgTournament

> CreateTcgTournamentResponse CreateTcgTournament(ctx, game).Execute()

Host a Swiss tournament: name, optional top cut (4 or 8), round length, player cap and start time.

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
    resp, r, err := apiClient.TCGApi.CreateTcgTournament(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.CreateTcgTournament``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgTournament`: CreateTcgTournamentResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.CreateTcgTournament`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgTournamentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateTcgTournamentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgTournamentDrop

> TcgTournamentResponse CreateTcgTournamentDrop(ctx, id).Execute()

Leave a tournament before it starts, or drop from it (your results stand).

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
    resp, r, err := apiClient.TCGApi.CreateTcgTournamentDrop(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.CreateTcgTournamentDrop``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgTournamentDrop`: TcgTournamentResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.CreateTcgTournamentDrop`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgTournamentDropRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**TcgTournamentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgTournamentRegister

> TcgTournamentResponse CreateTcgTournamentRegister(ctx, id).Execute()

Enter a tournament with an owned, legal deck before it starts.

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
    resp, r, err := apiClient.TCGApi.CreateTcgTournamentRegister(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.CreateTcgTournamentRegister``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgTournamentRegister`: TcgTournamentResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.CreateTcgTournamentRegister`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgTournamentRegisterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**TcgTournamentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTcgTournamentStart

> TcgTournamentResponse CreateTcgTournamentStart(ctx, id).Execute()

As the host, start the tournament: Swiss rounds are sized to the field and round 1 is dealt.

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
    resp, r, err := apiClient.TCGApi.CreateTcgTournamentStart(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.CreateTcgTournamentStart``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTcgTournamentStart`: TcgTournamentResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.CreateTcgTournamentStart`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTcgTournamentStartRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**TcgTournamentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteTcgDeck

> DeleteTcgDeck(ctx, id).Execute()

Delete one of your decks.

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
    resp, r, err := apiClient.TCGApi.DeleteTcgDeck(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.DeleteTcgDeck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTcgDeckRequest struct via the builder pattern


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


## DeleteTcgQueue

> DeleteTcgQueue(ctx, game).Execute()

Leave the ranked queue.

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
    resp, r, err := apiClient.TCGApi.DeleteTcgQueue(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.DeleteTcgQueue``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTcgQueueRequest struct via the builder pattern


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


## GetTcgCreatorGamePrint

> GetTcgCreatorGamePrintResponse GetTcgCreatorGamePrint(ctx, id).Execute()

Print file links for your own creator card game (owner only): the set's print file, home sheets and a proof, for its newest version.

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
    resp, r, err := apiClient.TCGApi.GetTcgCreatorGamePrint(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.GetTcgCreatorGamePrint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTcgCreatorGamePrint`: GetTcgCreatorGamePrintResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.GetTcgCreatorGamePrint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcgCreatorGamePrintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetTcgCreatorGamePrintResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTcgCreatorGamePrintByKind

> *os.File GetTcgCreatorGamePrintByKind(ctx, id, kind).Execute()

A PDF of your own creator card game (owner only): kind set | sheet | proof; ?version=, ?paper=letter|a4, ?card= for a proof. Binary (application/pdf).

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
    kind := "kind_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TCGApi.GetTcgCreatorGamePrintByKind(context.Background(), id, kind).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.GetTcgCreatorGamePrintByKind``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTcgCreatorGamePrintByKind`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.GetTcgCreatorGamePrintByKind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**kind** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcgCreatorGamePrintByKindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

***os.File**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTcgLadder

> GetTcgLadderResponse GetTcgLadder(ctx, game).Execute()

The season's ranked leaderboard for a game: place, rating, tier and record (named only where the player chose to be).

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
    resp, r, err := apiClient.TCGApi.GetTcgLadder(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.GetTcgLadder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTcgLadder`: GetTcgLadderResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.GetTcgLadder`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcgLadderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetTcgLadderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTcgLadderMe

> GetTcgLadderMeResponse GetTcgLadderMe(ctx, game).Execute()

Your ranked rating, tier, place and record this season, and whether you are queued.

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
    resp, r, err := apiClient.TCGApi.GetTcgLadderMe(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.GetTcgLadderMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTcgLadderMe`: GetTcgLadderMeResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.GetTcgLadderMe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcgLadderMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetTcgLadderMeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTcgMatch

> GetTcgMatchResponse GetTcgMatch(ctx, game).Execute()

Your recent matches, and open challenges you could join.

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
    resp, r, err := apiClient.TCGApi.GetTcgMatch(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.GetTcgMatch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTcgMatch`: GetTcgMatchResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.GetTcgMatch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcgMatchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetTcgMatchResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTcgMatchById

> TcgResponse GetTcgMatchById(ctx, id).Execute()

Your view of a match (no hidden information) and, on your turn, every legal move.

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
    resp, r, err := apiClient.TCGApi.GetTcgMatchById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.GetTcgMatchById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTcgMatchById`: TcgResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.GetTcgMatchById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcgMatchByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**TcgResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTcgMe

> GetTcgMeResponse GetTcgMe(ctx, game).Execute()

Your starter faction, collection (card id -> copies), decks and earned packs.

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
    resp, r, err := apiClient.TCGApi.GetTcgMe(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.GetTcgMe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTcgMe`: GetTcgMeResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.GetTcgMe`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcgMeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetTcgMeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTcgPrint

> GetTcgPrintResponse GetTcgPrint(ctx, game).Execute()

Print files for one of Crossly's own card games: a card proof (template), 3 x 3 Letter/A4 home sheets with crop marks, and the set's print file (63 x 88 mm, 3 mm bleed).

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
    resp, r, err := apiClient.TCGApi.GetTcgPrint(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.GetTcgPrint``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTcgPrint`: GetTcgPrintResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.GetTcgPrint`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcgPrintRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetTcgPrintResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTcgTournament

> TcgTournamentResponse GetTcgTournament(ctx, id).Execute()

A tournament's standings (points, opponents' and game win %), this round's pairings and your seat.

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
    resp, r, err := apiClient.TCGApi.GetTcgTournament(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.GetTcgTournament``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTcgTournament`: TcgTournamentResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.GetTcgTournament`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTcgTournamentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**TcgTournamentResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListTcgCards

> ListTcgCardsResponse ListTcgCards(ctx, game).Execute()

The game's card pool: every card with its rules text, stats, rarity and collector number.

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
    resp, r, err := apiClient.TCGApi.ListTcgCards(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.ListTcgCards``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTcgCards`: ListTcgCardsResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.ListTcgCards`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTcgCardsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListTcgCardsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListTcgGames

> ListTcgGamesResponse ListTcgGames(ctx).Execute()

Crossly's own card games: id, name, seats and card count.

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
    resp, r, err := apiClient.TCGApi.ListTcgGames(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.ListTcgGames``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTcgGames`: ListTcgGamesResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.ListTcgGames`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListTcgGamesRequest struct via the builder pattern


### Return type

**ListTcgGamesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListTcgTournaments

> ListTcgTournamentsResponse ListTcgTournaments(ctx, game).Execute()

A game's recent and upcoming Swiss tournaments (free to enter; prizes are earned packs).

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
    resp, r, err := apiClient.TCGApi.ListTcgTournaments(context.Background(), game).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TCGApi.ListTcgTournaments``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListTcgTournaments`: ListTcgTournamentsResponse
    fmt.Fprintf(os.Stdout, "Response from `TCGApi.ListTcgTournaments`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**game** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListTcgTournamentsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListTcgTournamentsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

