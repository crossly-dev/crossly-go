# \TradeBlockApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTradeBlockEvolutionClaim**](TradeBlockApi.md#CreateTradeBlockEvolutionClaim) | **Post** /v1/trade-block/evolutions/{id}/claim | Burn COSM and mint the evolved (or set-master) NFT for one of your unlocks.
[**CreateTradeBlockEvolutionRecipe**](TradeBlockApi.md#CreateTradeBlockEvolutionRecipe) | **Post** /v1/trade-block/evolution-recipes | Author a Trade Evolution recipe (a graded item or pair that, when swapped, unlocks a cosmetic you designed). Born unreviewed until ops approves.
[**CreateTradeBlockOfferRespond**](TradeBlockApi.md#CreateTradeBlockOfferRespond) | **Post** /v1/trade-block/offers/{id}/respond | Accept, decline, counter, or withdraw a Trade Block offer.
[**CreateTradeBlockPostClose**](TradeBlockApi.md#CreateTradeBlockPostClose) | **Post** /v1/trade-block/posts/{id}/close | Close one of your posts so it stops taking offers.
[**CreateTradeBlockPostOffer**](TradeBlockApi.md#CreateTradeBlockPostOffer) | **Post** /v1/trade-block/posts/{id}/offers | Respond to a post with a concrete swap: items on each side (0..N), optional cash.
[**CreateTradeBlockSetRecipe**](TradeBlockApi.md#CreateTradeBlockSetRecipe) | **Post** /v1/trade-block/set-recipes | Author a set-completion recipe (a set of items that, once assembled through trading, unlocks a set-master cosmetic you designed). Born unreviewed until ops approves.
[**GetTradeBlockEvolution**](TradeBlockApi.md#GetTradeBlockEvolution) | **Get** /v1/trade-block/evolutions | List your Trade Evolution unlocks — completed swaps that unlocked an evolved or set-master NFT you can mint by burning COSM.
[**GetTradeBlockPost**](TradeBlockApi.md#GetTradeBlockPost) | **Get** /v1/trade-block/posts | Browse open Trade Block posts (wants + offers) on the Crossly Market.
[**GetTradeBlockPostById**](TradeBlockApi.md#GetTradeBlockPostById) | **Get** /v1/trade-block/posts/{id} | Get one Trade Block post with its items and open-offer count.
[**GetTradeBlockTrade**](TradeBlockApi.md#GetTradeBlockTrade) | **Get** /v1/trade-block/trades | List the trades you are a party to, newest first.



## CreateTradeBlockEvolutionClaim

> CreateTradeBlockEvolutionClaimResponse CreateTradeBlockEvolutionClaim(ctx, id).Execute()

Burn COSM and mint the evolved (or set-master) NFT for one of your unlocks.

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
    resp, r, err := apiClient.TradeBlockApi.CreateTradeBlockEvolutionClaim(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeBlockApi.CreateTradeBlockEvolutionClaim``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTradeBlockEvolutionClaim`: CreateTradeBlockEvolutionClaimResponse
    fmt.Fprintf(os.Stdout, "Response from `TradeBlockApi.CreateTradeBlockEvolutionClaim`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeBlockEvolutionClaimRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateTradeBlockEvolutionClaimResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTradeBlockEvolutionRecipe

> CreateTradeBlockEvolutionRecipeResponse CreateTradeBlockEvolutionRecipe(ctx).Execute()

Author a Trade Evolution recipe (a graded item or pair that, when swapped, unlocks a cosmetic you designed). Born unreviewed until ops approves.

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
    resp, r, err := apiClient.TradeBlockApi.CreateTradeBlockEvolutionRecipe(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeBlockApi.CreateTradeBlockEvolutionRecipe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTradeBlockEvolutionRecipe`: CreateTradeBlockEvolutionRecipeResponse
    fmt.Fprintf(os.Stdout, "Response from `TradeBlockApi.CreateTradeBlockEvolutionRecipe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeBlockEvolutionRecipeRequest struct via the builder pattern


### Return type

**CreateTradeBlockEvolutionRecipeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTradeBlockOfferRespond

> CreateTradeBlockOfferRespondResponse CreateTradeBlockOfferRespond(ctx, id).Execute()

Accept, decline, counter, or withdraw a Trade Block offer.

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
    resp, r, err := apiClient.TradeBlockApi.CreateTradeBlockOfferRespond(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeBlockApi.CreateTradeBlockOfferRespond``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTradeBlockOfferRespond`: CreateTradeBlockOfferRespondResponse
    fmt.Fprintf(os.Stdout, "Response from `TradeBlockApi.CreateTradeBlockOfferRespond`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeBlockOfferRespondRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateTradeBlockOfferRespondResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTradeBlockPostClose

> OkResponse CreateTradeBlockPostClose(ctx, id).Execute()

Close one of your posts so it stops taking offers.

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
    resp, r, err := apiClient.TradeBlockApi.CreateTradeBlockPostClose(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeBlockApi.CreateTradeBlockPostClose``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTradeBlockPostClose`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `TradeBlockApi.CreateTradeBlockPostClose`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeBlockPostCloseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTradeBlockPostOffer

> IdResponse CreateTradeBlockPostOffer(ctx, id).Execute()

Respond to a post with a concrete swap: items on each side (0..N), optional cash.

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
    resp, r, err := apiClient.TradeBlockApi.CreateTradeBlockPostOffer(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeBlockApi.CreateTradeBlockPostOffer``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTradeBlockPostOffer`: IdResponse
    fmt.Fprintf(os.Stdout, "Response from `TradeBlockApi.CreateTradeBlockPostOffer`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeBlockPostOfferRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**IdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateTradeBlockSetRecipe

> CreateTradeBlockSetRecipeResponse CreateTradeBlockSetRecipe(ctx).Execute()

Author a set-completion recipe (a set of items that, once assembled through trading, unlocks a set-master cosmetic you designed). Born unreviewed until ops approves.

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
    resp, r, err := apiClient.TradeBlockApi.CreateTradeBlockSetRecipe(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeBlockApi.CreateTradeBlockSetRecipe``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTradeBlockSetRecipe`: CreateTradeBlockSetRecipeResponse
    fmt.Fprintf(os.Stdout, "Response from `TradeBlockApi.CreateTradeBlockSetRecipe`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTradeBlockSetRecipeRequest struct via the builder pattern


### Return type

**CreateTradeBlockSetRecipeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTradeBlockEvolution

> GetTradeBlockEvolutionResponse GetTradeBlockEvolution(ctx).Execute()

List your Trade Evolution unlocks — completed swaps that unlocked an evolved or set-master NFT you can mint by burning COSM.

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
    resp, r, err := apiClient.TradeBlockApi.GetTradeBlockEvolution(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeBlockApi.GetTradeBlockEvolution``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTradeBlockEvolution`: GetTradeBlockEvolutionResponse
    fmt.Fprintf(os.Stdout, "Response from `TradeBlockApi.GetTradeBlockEvolution`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeBlockEvolutionRequest struct via the builder pattern


### Return type

**GetTradeBlockEvolutionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTradeBlockPost

> GetTradeBlockPostResponse GetTradeBlockPost(ctx).Kind(kind).CategorySlug(categorySlug).Brand(brand).Model(model).Q(q).Limit(limit).Offset(offset).Execute()

Browse open Trade Block posts (wants + offers) on the Crossly Market.

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
    kind := "kind_example" // string | Filter to wants or offers. (optional)
    categorySlug := "categorySlug_example" // string | Filter to a market category. (optional)
    brand := "brand_example" // string | Filter to posts with a product of this brand (case-insensitive). (optional)
    model := "model_example" // string | Filter to posts with a product of this catalogue model (case-insensitive). (optional)
    q := "q_example" // string | Search title + want text. (optional)
    limit := int32(56) // int32 |  (optional) (default to 40)
    offset := int32(56) // int32 |  (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.TradeBlockApi.GetTradeBlockPost(context.Background()).Kind(kind).CategorySlug(categorySlug).Brand(brand).Model(model).Q(q).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeBlockApi.GetTradeBlockPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTradeBlockPost`: GetTradeBlockPostResponse
    fmt.Fprintf(os.Stdout, "Response from `TradeBlockApi.GetTradeBlockPost`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeBlockPostRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kind** | **string** | Filter to wants or offers. | 
 **categorySlug** | **string** | Filter to a market category. | 
 **brand** | **string** | Filter to posts with a product of this brand (case-insensitive). | 
 **model** | **string** | Filter to posts with a product of this catalogue model (case-insensitive). | 
 **q** | **string** | Search title + want text. | 
 **limit** | **int32** |  | [default to 40]
 **offset** | **int32** |  | [default to 0]

### Return type

**GetTradeBlockPostResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTradeBlockPostById

> GetTradeBlockPostByIdResponse GetTradeBlockPostById(ctx, id).Execute()

Get one Trade Block post with its items and open-offer count.

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
    resp, r, err := apiClient.TradeBlockApi.GetTradeBlockPostById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeBlockApi.GetTradeBlockPostById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTradeBlockPostById`: GetTradeBlockPostByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `TradeBlockApi.GetTradeBlockPostById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeBlockPostByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetTradeBlockPostByIdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetTradeBlockTrade

> GetTradeBlockTradeResponse GetTradeBlockTrade(ctx).Execute()

List the trades you are a party to, newest first.

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
    resp, r, err := apiClient.TradeBlockApi.GetTradeBlockTrade(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TradeBlockApi.GetTradeBlockTrade``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTradeBlockTrade`: GetTradeBlockTradeResponse
    fmt.Fprintf(os.Stdout, "Response from `TradeBlockApi.GetTradeBlockTrade`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetTradeBlockTradeRequest struct via the builder pattern


### Return type

**GetTradeBlockTradeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

