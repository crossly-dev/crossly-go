# \AIApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAiCategorize**](AIApi.md#CreateAiCategorize) | **Post** /v1/ai/categorize | Taxonomy guess from a single image URL.
[**CreateAiCategorizeFromImage**](AIApi.md#CreateAiCategorizeFromImage) | **Post** /v1/ai/categorize-from-image | Taxonomy guess from a single base64 image.
[**CreateAiEnhanceDescription**](AIApi.md#CreateAiEnhanceDescription) | **Post** /v1/ai/enhance-description | SEO-rewrite a listing description.
[**CreateAiEnhanceListing**](AIApi.md#CreateAiEnhanceListing) | **Post** /v1/ai/enhance-listing | Rewrite title + description + tags in one call.
[**CreateAiEnhanceTitle**](AIApi.md#CreateAiEnhanceTitle) | **Post** /v1/ai/enhance-title | SEO-rewrite a listing title.
[**CreateAiExtractReceipt**](AIApi.md#CreateAiExtractReceipt) | **Post** /v1/ai/extract-receipt | Structured data extraction from a receipt photo.
[**CreateAiGenerateListing**](AIApi.md#CreateAiGenerateListing) | **Post** /v1/ai/generate-listing | Generate full listing fields from up to 4 image URLs.
[**CreateAiHelp**](AIApi.md#CreateAiHelp) | **Post** /v1/ai/help | In-app help Q&amp;A grounded in supplied docs.
[**CreateAiMagicListing**](AIApi.md#CreateAiMagicListing) | **Post** /v1/ai/magic-listing | Generate full listing fields from base64 photos.
[**CreateAiTestKey**](AIApi.md#CreateAiTestKey) | **Post** /v1/ai/test-key | Live-ping a candidate BYO-key.
[**DeleteAiKey**](AIApi.md#DeleteAiKey) | **Delete** /v1/ai/key | Remove the BYO-key for a provider.
[**GetAiProvider**](AIApi.md#GetAiProvider) | **Get** /v1/ai/providers | Static catalog of supported AI providers.
[**GetAiStatus**](AIApi.md#GetAiStatus) | **Get** /v1/ai/status | BYO-key state for the calling user.
[**UpdateAiKey**](AIApi.md#UpdateAiKey) | **Put** /v1/ai/key | Save an encrypted BYO-key for an AI provider.



## CreateAiCategorize

> CreateAiCategorizeResponse CreateAiCategorize(ctx).Execute()

Taxonomy guess from a single image URL.

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
    resp, r, err := apiClient.AIApi.CreateAiCategorize(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIApi.CreateAiCategorize``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiCategorize`: CreateAiCategorizeResponse
    fmt.Fprintf(os.Stdout, "Response from `AIApi.CreateAiCategorize`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiCategorizeRequest struct via the builder pattern


### Return type

[**CreateAiCategorizeResponse**](CreateAiCategorizeResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiCategorizeFromImage

> CreateAiCategorizeFromImageResponse CreateAiCategorizeFromImage(ctx).Execute()

Taxonomy guess from a single base64 image.

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
    resp, r, err := apiClient.AIApi.CreateAiCategorizeFromImage(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIApi.CreateAiCategorizeFromImage``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiCategorizeFromImage`: CreateAiCategorizeFromImageResponse
    fmt.Fprintf(os.Stdout, "Response from `AIApi.CreateAiCategorizeFromImage`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiCategorizeFromImageRequest struct via the builder pattern


### Return type

[**CreateAiCategorizeFromImageResponse**](CreateAiCategorizeFromImageResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiEnhanceDescription

> CreateAiEnhanceDescriptionResponse CreateAiEnhanceDescription(ctx).Execute()

SEO-rewrite a listing description.

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
    resp, r, err := apiClient.AIApi.CreateAiEnhanceDescription(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIApi.CreateAiEnhanceDescription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiEnhanceDescription`: CreateAiEnhanceDescriptionResponse
    fmt.Fprintf(os.Stdout, "Response from `AIApi.CreateAiEnhanceDescription`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiEnhanceDescriptionRequest struct via the builder pattern


### Return type

[**CreateAiEnhanceDescriptionResponse**](CreateAiEnhanceDescriptionResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiEnhanceListing

> CreateAiEnhanceListingResponse CreateAiEnhanceListing(ctx).Execute()

Rewrite title + description + tags in one call.

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
    resp, r, err := apiClient.AIApi.CreateAiEnhanceListing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIApi.CreateAiEnhanceListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiEnhanceListing`: CreateAiEnhanceListingResponse
    fmt.Fprintf(os.Stdout, "Response from `AIApi.CreateAiEnhanceListing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiEnhanceListingRequest struct via the builder pattern


### Return type

[**CreateAiEnhanceListingResponse**](CreateAiEnhanceListingResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiEnhanceTitle

> CreateAiEnhanceTitleResponse CreateAiEnhanceTitle(ctx).Execute()

SEO-rewrite a listing title.

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
    resp, r, err := apiClient.AIApi.CreateAiEnhanceTitle(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIApi.CreateAiEnhanceTitle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiEnhanceTitle`: CreateAiEnhanceTitleResponse
    fmt.Fprintf(os.Stdout, "Response from `AIApi.CreateAiEnhanceTitle`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiEnhanceTitleRequest struct via the builder pattern


### Return type

[**CreateAiEnhanceTitleResponse**](CreateAiEnhanceTitleResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiExtractReceipt

> CreateAiExtractReceiptResponse CreateAiExtractReceipt(ctx).Execute()

Structured data extraction from a receipt photo.

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
    resp, r, err := apiClient.AIApi.CreateAiExtractReceipt(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIApi.CreateAiExtractReceipt``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiExtractReceipt`: CreateAiExtractReceiptResponse
    fmt.Fprintf(os.Stdout, "Response from `AIApi.CreateAiExtractReceipt`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiExtractReceiptRequest struct via the builder pattern


### Return type

[**CreateAiExtractReceiptResponse**](CreateAiExtractReceiptResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiGenerateListing

> CreateAiGenerateListingResponse CreateAiGenerateListing(ctx).Execute()

Generate full listing fields from up to 4 image URLs.

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
    resp, r, err := apiClient.AIApi.CreateAiGenerateListing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIApi.CreateAiGenerateListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiGenerateListing`: CreateAiGenerateListingResponse
    fmt.Fprintf(os.Stdout, "Response from `AIApi.CreateAiGenerateListing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiGenerateListingRequest struct via the builder pattern


### Return type

[**CreateAiGenerateListingResponse**](CreateAiGenerateListingResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiHelp

> CreateAiHelpResponse CreateAiHelp(ctx).Execute()

In-app help Q&A grounded in supplied docs.

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
    resp, r, err := apiClient.AIApi.CreateAiHelp(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIApi.CreateAiHelp``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiHelp`: CreateAiHelpResponse
    fmt.Fprintf(os.Stdout, "Response from `AIApi.CreateAiHelp`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiHelpRequest struct via the builder pattern


### Return type

[**CreateAiHelpResponse**](CreateAiHelpResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiMagicListing

> CreateAiMagicListingResponse CreateAiMagicListing(ctx).Execute()

Generate full listing fields from base64 photos.

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
    resp, r, err := apiClient.AIApi.CreateAiMagicListing(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIApi.CreateAiMagicListing``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiMagicListing`: CreateAiMagicListingResponse
    fmt.Fprintf(os.Stdout, "Response from `AIApi.CreateAiMagicListing`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiMagicListingRequest struct via the builder pattern


### Return type

[**CreateAiMagicListingResponse**](CreateAiMagicListingResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateAiTestKey

> CreateAiTestKeyResponse CreateAiTestKey(ctx).Execute()

Live-ping a candidate BYO-key.

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
    resp, r, err := apiClient.AIApi.CreateAiTestKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIApi.CreateAiTestKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAiTestKey`: CreateAiTestKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AIApi.CreateAiTestKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAiTestKeyRequest struct via the builder pattern


### Return type

[**CreateAiTestKeyResponse**](CreateAiTestKeyResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteAiKey

> DeleteAiKeyResponse DeleteAiKey(ctx).Execute()

Remove the BYO-key for a provider.

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
    resp, r, err := apiClient.AIApi.DeleteAiKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIApi.DeleteAiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteAiKey`: DeleteAiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AIApi.DeleteAiKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAiKeyRequest struct via the builder pattern


### Return type

[**DeleteAiKeyResponse**](DeleteAiKeyResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiProvider

> GetAiProviderResponse GetAiProvider(ctx).Execute()

Static catalog of supported AI providers.

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
    resp, r, err := apiClient.AIApi.GetAiProvider(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIApi.GetAiProvider``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiProvider`: GetAiProviderResponse
    fmt.Fprintf(os.Stdout, "Response from `AIApi.GetAiProvider`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiProviderRequest struct via the builder pattern


### Return type

[**GetAiProviderResponse**](GetAiProviderResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAiStatus

> GetAiStatusResponse GetAiStatus(ctx).Execute()

BYO-key state for the calling user.

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
    resp, r, err := apiClient.AIApi.GetAiStatus(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIApi.GetAiStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAiStatus`: GetAiStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `AIApi.GetAiStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetAiStatusRequest struct via the builder pattern


### Return type

[**GetAiStatusResponse**](GetAiStatusResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAiKey

> UpdateAiKeyResponse UpdateAiKey(ctx).Execute()

Save an encrypted BYO-key for an AI provider.

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
    resp, r, err := apiClient.AIApi.UpdateAiKey(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AIApi.UpdateAiKey``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAiKey`: UpdateAiKeyResponse
    fmt.Fprintf(os.Stdout, "Response from `AIApi.UpdateAiKey`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAiKeyRequest struct via the builder pattern


### Return type

[**UpdateAiKeyResponse**](UpdateAiKeyResponse.md)

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

