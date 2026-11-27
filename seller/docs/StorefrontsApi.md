# \StorefrontsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateStorefront**](StorefrontsApi.md#CreateStorefront) | **Post** /v1/storefronts | Create a store: name (1-60) and optional slug (3-40 lowercase letters, numbers, hyphens; derived from the name when omitted). Your first store becomes the default.
[**CreateStorefrontListingMove**](StorefrontsApi.md#CreateStorefrontListingMove) | **Post** /v1/storefronts/{id}/listings/{platformListingId}/move | Move a Crossly listing (by its platform_listing id) into this store.
[**CreateStorefrontOnboardingSkip**](StorefrontsApi.md#CreateStorefrontOnboardingSkip) | **Post** /v1/storefronts/onboarding/skip | Skip (or dismiss) the setup wizard: ensures your default store exists, with a generated name you can rename any time. Idempotent.
[**CreateStorefrontOnboardingSuggest**](StorefrontsApi.md#CreateStorefrontOnboardingSuggest) | **Post** /v1/storefronts/onboarding/suggest | Suggest another generated store name with a free matching slug.
[**CreateStorefrontPolicyFromTerm**](StorefrontsApi.md#CreateStorefrontPolicyFromTerm) | **Post** /v1/storefronts/{id}/policies/from-terms | Create a policy preset from this store&#39;s current terms (body: { kind: &#39;shipping&#39; | &#39;return&#39; | &#39;payment&#39; }) and assign it to the store. Returns { preset, policies }.
[**CreateStorefrontPolicyUseShared**](StorefrontsApi.md#CreateStorefrontPolicyUseShared) | **Post** /v1/storefronts/{id}/policies/use-shared | Make a store use your shared policies: a non-default store drops its own return/shipping/payment policy copies and shows your default policies from then on. No change on the default store. Returns the store settings (policySource per kind, usesOwnPolicy).
[**CreateStorefrontPublish**](StorefrontsApi.md#CreateStorefrontPublish) | **Post** /v1/storefronts/{id}/publish | Flip a store live. Fails with storefront_requirements_missing (and the missing list) when ship-from, shipping policy, return policy or payouts are not set.
[**CreateStorefrontSectionById**](StorefrontsApi.md#CreateStorefrontSectionById) | **Post** /v1/storefronts/{id}/sections | Add a section to a store (kind hero | listings | collections | about | policies | social | announcement | faq | creations).
[**DeleteStorefront**](StorefrontsApi.md#DeleteStorefront) | **Delete** /v1/storefronts/{id} | Archive a store. If it was the default, your oldest remaining store becomes the default.
[**DeleteStorefrontSectionBySectionId**](StorefrontsApi.md#DeleteStorefrontSectionBySectionId) | **Delete** /v1/storefronts/{id}/sections/{sectionId} | Remove a section from a store. Listings on a removed shelf stay live.
[**DeleteStorefrontTeamAccess**](StorefrontsApi.md#DeleteStorefrontTeamAccess) | **Delete** /v1/storefronts/team-access | Revoke a team member&#39;s access to one store (or their all-stores grant).
[**GetHubSetup**](StorefrontsApi.md#GetHubSetup) | **Get** /v1/hub/setup/{section} | A Hub section&#39;s setup checklist: each step, whether it is done and the page that completes it. Derived from what exists, never stored.
[**GetStorefrontById**](StorefrontsApi.md#GetStorefrontById) | **Get** /v1/storefronts/{id} | One store: name, slug, status (draft|live|paused|archived), isDefault, nameAuto, onboardedAt.
[**GetStorefrontOnboarding**](StorefrontsApi.md#GetStorefrontOnboarding) | **Get** /v1/storefronts/onboarding | Crossly Shop setup state: whether you have a store, the generated name/slug a skip would use, each optional step and whether it is done, the listing-gate requirements, payouts readiness and the public URL once live.
[**GetStorefrontPolicy**](StorefrontsApi.md#GetStorefrontPolicy) | **Get** /v1/storefronts/{id}/policies | A store&#39;s policies: per kind (shipping, return, payment) the preset the store picked (presetId, null &#x3D; your default) and what the store resolves to (resolvedPresetId, resolvedPresetName, level: store | seller_default | legacy | default, body).
[**GetStorefrontRequirement**](StorefrontsApi.md#GetStorefrontRequirement) | **Get** /v1/storefronts/{id}/requirements | What this store still needs before a NEW Crossly listing can publish into it: ship_from, shipping_policy, return_policy, payouts — with where to fix each.
[**GetStorefrontSetting**](StorefrontsApi.md#GetStorefrontSetting) | **Get** /v1/storefronts/{id}/settings | A store&#39;s settings: design (banner, logo, theme, SEO, social, published), ship-from, return address, policies (storePolicies: the store&#39;s preset picks + what each kind resolves to), custom-order flag, featured listings, bundle discount, custom domain. DEPRECATED aliases, resolved through the policy cascade: offersFreeShipping, freeShippingThresholdCents, offersLocalPickup, localCourierEnabled, acceptsLayaway, sellerProtections, requestArrivalCondition, arrivalConditionWindowHours — use /v1/storefronts/:id/policies and the preset fields.
[**GetStorefrontSlugAvailable**](StorefrontsApi.md#GetStorefrontSlugAvailable) | **Get** /v1/storefronts/slug-available | Is this store slug free? Checks shape, reserved words, current slugs and retired ones.
[**ListStorefrontSectionsById**](StorefrontsApi.md#ListStorefrontSectionsById) | **Get** /v1/storefronts/{id}/sections | A store&#39;s sections in display order.
[**ListStorefrontTeamAccess**](StorefrontsApi.md#ListStorefrontTeamAccess) | **Get** /v1/storefronts/team-access | Every active per-store grant you have made to team members (storefrontId null &#x3D; all stores).
[**ListStorefronts**](StorefrontsApi.md#ListStorefronts) | **Get** /v1/storefronts | Your stores (default first). includeArchived&#x3D;true adds archived ones.
[**UpdateStorefrontById**](StorefrontsApi.md#UpdateStorefrontById) | **Patch** /v1/storefronts/{id} | Rename a store. Never changes the slug or your @handle.
[**UpdateStorefrontDefault**](StorefrontsApi.md#UpdateStorefrontDefault) | **Put** /v1/storefronts/{id}/default | Make this store your default (where a Crossly publish with no storefrontId lands).
[**UpdateStorefrontPolicy**](StorefrontsApi.md#UpdateStorefrontPolicy) | **Put** /v1/storefronts/{id}/policies | Set a store&#39;s policies: shippingPresetId, returnPresetId, paymentPresetId (a policy preset of that kind you own, or null for your default). Omitted &#x3D; unchanged. Listings in the store with no preset of their own follow these.
[**UpdateStorefrontSectionBySectionId**](StorefrontsApi.md#UpdateStorefrontSectionBySectionId) | **Patch** /v1/storefronts/{id}/sections/{sectionId} | Rename, show/hide, refill (listings shelves) or reconfigure one section of a store.
[**UpdateStorefrontSectionReorderById**](StorefrontsApi.md#UpdateStorefrontSectionReorderById) | **Put** /v1/storefronts/{id}/sections/reorder | Set a store&#39;s section order in one atomic call: every section id, top to bottom.
[**UpdateStorefrontSetting**](StorefrontsApi.md#UpdateStorefrontSetting) | **Put** /v1/storefronts/{id}/settings | Update a store&#39;s settings. Omitted fields are unchanged, null clears. Policies replace wholesale. shippingPresetId / returnPresetId / paymentPresetId pick the store&#39;s presets. The checkout-term fields (offersFreeShipping … arrivalConditionWindowHours) are DEPRECATED aliases: they are written into the store&#39;s preset.
[**UpdateStorefrontSlug**](StorefrontsApi.md#UpdateStorefrontSlug) | **Put** /v1/storefronts/{id}/slug | Change a store&#39;s slug. The old slug keeps redirecting to the store.
[**UpdateStorefrontStatus**](StorefrontsApi.md#UpdateStorefrontStatus) | **Put** /v1/storefronts/{id}/status | Set a store&#39;s status: draft | live | paused | archived. A live store is public once it has a live listing.
[**UpdateStorefrontTeamAccess**](StorefrontsApi.md#UpdateStorefrontTeamAccess) | **Put** /v1/storefronts/team-access | Grant (or replace) a team member&#39;s access to one store, or to every store (storefrontId null), with scopes.



## CreateStorefront

> StorefrontResponse CreateStorefront(ctx).Execute()

Create a store: name (1-60) and optional slug (3-40 lowercase letters, numbers, hyphens; derived from the name when omitted). Your first store becomes the default.

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
    resp, r, err := apiClient.StorefrontsApi.CreateStorefront(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.CreateStorefront``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefront`: StorefrontResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.CreateStorefront`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontRequest struct via the builder pattern


### Return type

**StorefrontResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontListingMove

> OkResponse CreateStorefrontListingMove(ctx, id, platformListingId).Execute()

Move a Crossly listing (by its platform_listing id) into this store.

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
    platformListingId := "platformListingId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorefrontsApi.CreateStorefrontListingMove(context.Background(), id, platformListingId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.CreateStorefrontListingMove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontListingMove`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.CreateStorefrontListingMove`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**platformListingId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontListingMoveRequest struct via the builder pattern


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


## CreateStorefrontOnboardingSkip

> StorefrontResponse CreateStorefrontOnboardingSkip(ctx).Execute()

Skip (or dismiss) the setup wizard: ensures your default store exists, with a generated name you can rename any time. Idempotent.

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
    resp, r, err := apiClient.StorefrontsApi.CreateStorefrontOnboardingSkip(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.CreateStorefrontOnboardingSkip``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontOnboardingSkip`: StorefrontResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.CreateStorefrontOnboardingSkip`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontOnboardingSkipRequest struct via the builder pattern


### Return type

**StorefrontResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontOnboardingSuggest

> CreateStorefrontOnboardingSuggestResponse CreateStorefrontOnboardingSuggest(ctx).Execute()

Suggest another generated store name with a free matching slug.

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
    resp, r, err := apiClient.StorefrontsApi.CreateStorefrontOnboardingSuggest(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.CreateStorefrontOnboardingSuggest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontOnboardingSuggest`: CreateStorefrontOnboardingSuggestResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.CreateStorefrontOnboardingSuggest`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontOnboardingSuggestRequest struct via the builder pattern


### Return type

**CreateStorefrontOnboardingSuggestResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontPolicyFromTerm

> CreateStorefrontPolicyFromTermResponse CreateStorefrontPolicyFromTerm(ctx, id).Execute()

Create a policy preset from this store's current terms (body: { kind: 'shipping' | 'return' | 'payment' }) and assign it to the store. Returns { preset, policies }.

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
    resp, r, err := apiClient.StorefrontsApi.CreateStorefrontPolicyFromTerm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.CreateStorefrontPolicyFromTerm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontPolicyFromTerm`: CreateStorefrontPolicyFromTermResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.CreateStorefrontPolicyFromTerm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontPolicyFromTermRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateStorefrontPolicyFromTermResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontPolicyUseShared

> UpdateStorefrontSettingResponse CreateStorefrontPolicyUseShared(ctx, id).Execute()

Make a store use your shared policies: a non-default store drops its own return/shipping/payment policy copies and shows your default policies from then on. No change on the default store. Returns the store settings (policySource per kind, usesOwnPolicy).

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
    resp, r, err := apiClient.StorefrontsApi.CreateStorefrontPolicyUseShared(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.CreateStorefrontPolicyUseShared``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontPolicyUseShared`: UpdateStorefrontSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.CreateStorefrontPolicyUseShared`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontPolicyUseSharedRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateStorefrontSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontPublish

> StorefrontResponse CreateStorefrontPublish(ctx, id).Execute()

Flip a store live. Fails with storefront_requirements_missing (and the missing list) when ship-from, shipping policy, return policy or payouts are not set.

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
    resp, r, err := apiClient.StorefrontsApi.CreateStorefrontPublish(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.CreateStorefrontPublish``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontPublish`: StorefrontResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.CreateStorefrontPublish`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontPublishRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateStorefrontSectionById

> StorefrontSectionResponse CreateStorefrontSectionById(ctx, id).Execute()

Add a section to a store (kind hero | listings | collections | about | policies | social | announcement | faq | creations).

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
    resp, r, err := apiClient.StorefrontsApi.CreateStorefrontSectionById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.CreateStorefrontSectionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateStorefrontSectionById`: StorefrontSectionResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.CreateStorefrontSectionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStorefrontSectionByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontSectionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteStorefront

> StorefrontResponse DeleteStorefront(ctx, id).Execute()

Archive a store. If it was the default, your oldest remaining store becomes the default.

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
    resp, r, err := apiClient.StorefrontsApi.DeleteStorefront(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.DeleteStorefront``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStorefront`: StorefrontResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.DeleteStorefront`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorefrontRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteStorefrontSectionBySectionId

> OkResponse DeleteStorefrontSectionBySectionId(ctx, id, sectionId).Execute()

Remove a section from a store. Listings on a removed shelf stay live.

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
    sectionId := "sectionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorefrontsApi.DeleteStorefrontSectionBySectionId(context.Background(), id, sectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.DeleteStorefrontSectionBySectionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStorefrontSectionBySectionId`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.DeleteStorefrontSectionBySectionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**sectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorefrontSectionBySectionIdRequest struct via the builder pattern


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


## DeleteStorefrontTeamAccess

> OkResponse DeleteStorefrontTeamAccess(ctx).Execute()

Revoke a team member's access to one store (or their all-stores grant).

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
    resp, r, err := apiClient.StorefrontsApi.DeleteStorefrontTeamAccess(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.DeleteStorefrontTeamAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteStorefrontTeamAccess`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.DeleteStorefrontTeamAccess`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStorefrontTeamAccessRequest struct via the builder pattern


### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetHubSetup

> GetHubSetupResponse GetHubSetup(ctx, section).Execute()

A Hub section's setup checklist: each step, whether it is done and the page that completes it. Derived from what exists, never stored.

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
    section := "section_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorefrontsApi.GetHubSetup(context.Background(), section).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.GetHubSetup``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetHubSetup`: GetHubSetupResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.GetHubSetup`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**section** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetHubSetupRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetHubSetupResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetStorefrontById

> StorefrontResponse GetStorefrontById(ctx, id).Execute()

One store: name, slug, status (draft|live|paused|archived), isDefault, nameAuto, onboardedAt.

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
    resp, r, err := apiClient.StorefrontsApi.GetStorefrontById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.GetStorefrontById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorefrontById`: StorefrontResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.GetStorefrontById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorefrontByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetStorefrontOnboarding

> GetStorefrontOnboardingResponse GetStorefrontOnboarding(ctx).Execute()

Crossly Shop setup state: whether you have a store, the generated name/slug a skip would use, each optional step and whether it is done, the listing-gate requirements, payouts readiness and the public URL once live.

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
    resp, r, err := apiClient.StorefrontsApi.GetStorefrontOnboarding(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.GetStorefrontOnboarding``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorefrontOnboarding`: GetStorefrontOnboardingResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.GetStorefrontOnboarding`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorefrontOnboardingRequest struct via the builder pattern


### Return type

**GetStorefrontOnboardingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetStorefrontPolicy

> GetStorefrontPolicyResponse GetStorefrontPolicy(ctx, id).Execute()

A store's policies: per kind (shipping, return, payment) the preset the store picked (presetId, null = your default) and what the store resolves to (resolvedPresetId, resolvedPresetName, level: store | seller_default | legacy | default, body).

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
    resp, r, err := apiClient.StorefrontsApi.GetStorefrontPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.GetStorefrontPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorefrontPolicy`: GetStorefrontPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.GetStorefrontPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorefrontPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetStorefrontPolicyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetStorefrontRequirement

> GetStorefrontRequirementResponse GetStorefrontRequirement(ctx, id).Execute()

What this store still needs before a NEW Crossly listing can publish into it: ship_from, shipping_policy, return_policy, payouts — with where to fix each.

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
    resp, r, err := apiClient.StorefrontsApi.GetStorefrontRequirement(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.GetStorefrontRequirement``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorefrontRequirement`: GetStorefrontRequirementResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.GetStorefrontRequirement`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorefrontRequirementRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetStorefrontRequirementResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetStorefrontSetting

> GetStorefrontSettingResponse GetStorefrontSetting(ctx, id).Execute()

A store's settings: design (banner, logo, theme, SEO, social, published), ship-from, return address, policies (storePolicies: the store's preset picks + what each kind resolves to), custom-order flag, featured listings, bundle discount, custom domain. DEPRECATED aliases, resolved through the policy cascade: offersFreeShipping, freeShippingThresholdCents, offersLocalPickup, localCourierEnabled, acceptsLayaway, sellerProtections, requestArrivalCondition, arrivalConditionWindowHours — use /v1/storefronts/:id/policies and the preset fields.

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
    resp, r, err := apiClient.StorefrontsApi.GetStorefrontSetting(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.GetStorefrontSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorefrontSetting`: GetStorefrontSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.GetStorefrontSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorefrontSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetStorefrontSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetStorefrontSlugAvailable

> GetStorefrontSlugAvailableResponse GetStorefrontSlugAvailable(ctx).Execute()

Is this store slug free? Checks shape, reserved words, current slugs and retired ones.

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
    resp, r, err := apiClient.StorefrontsApi.GetStorefrontSlugAvailable(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.GetStorefrontSlugAvailable``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStorefrontSlugAvailable`: GetStorefrontSlugAvailableResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.GetStorefrontSlugAvailable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStorefrontSlugAvailableRequest struct via the builder pattern


### Return type

**GetStorefrontSlugAvailableResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListStorefrontSectionsById

> StorefrontResponse4 ListStorefrontSectionsById(ctx, id).Execute()

A store's sections in display order.

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
    resp, r, err := apiClient.StorefrontsApi.ListStorefrontSectionsById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.ListStorefrontSectionsById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStorefrontSectionsById`: StorefrontResponse4
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.ListStorefrontSectionsById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListStorefrontSectionsByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontResponse4**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListStorefrontTeamAccess

> ListStorefrontTeamAccessResponse ListStorefrontTeamAccess(ctx).Execute()

Every active per-store grant you have made to team members (storefrontId null = all stores).

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
    resp, r, err := apiClient.StorefrontsApi.ListStorefrontTeamAccess(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.ListStorefrontTeamAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStorefrontTeamAccess`: ListStorefrontTeamAccessResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.ListStorefrontTeamAccess`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStorefrontTeamAccessRequest struct via the builder pattern


### Return type

**ListStorefrontTeamAccessResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListStorefronts

> ListStorefrontsResponse ListStorefronts(ctx).Execute()

Your stores (default first). includeArchived=true adds archived ones.

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
    resp, r, err := apiClient.StorefrontsApi.ListStorefronts(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.ListStorefronts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListStorefronts`: ListStorefrontsResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.ListStorefronts`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListStorefrontsRequest struct via the builder pattern


### Return type

**ListStorefrontsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontById

> StorefrontResponse UpdateStorefrontById(ctx, id).Execute()

Rename a store. Never changes the slug or your @handle.

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
    resp, r, err := apiClient.StorefrontsApi.UpdateStorefrontById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.UpdateStorefrontById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontById`: StorefrontResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.UpdateStorefrontById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontDefault

> StorefrontResponse UpdateStorefrontDefault(ctx, id).Execute()

Make this store your default (where a Crossly publish with no storefrontId lands).

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
    resp, r, err := apiClient.StorefrontsApi.UpdateStorefrontDefault(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.UpdateStorefrontDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontDefault`: StorefrontResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.UpdateStorefrontDefault`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontPolicy

> GetStorefrontPolicyResponse UpdateStorefrontPolicy(ctx, id).Execute()

Set a store's policies: shippingPresetId, returnPresetId, paymentPresetId (a policy preset of that kind you own, or null for your default). Omitted = unchanged. Listings in the store with no preset of their own follow these.

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
    resp, r, err := apiClient.StorefrontsApi.UpdateStorefrontPolicy(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.UpdateStorefrontPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontPolicy`: GetStorefrontPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.UpdateStorefrontPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetStorefrontPolicyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontSectionBySectionId

> StorefrontSectionResponse UpdateStorefrontSectionBySectionId(ctx, id, sectionId).Execute()

Rename, show/hide, refill (listings shelves) or reconfigure one section of a store.

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
    sectionId := "sectionId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.StorefrontsApi.UpdateStorefrontSectionBySectionId(context.Background(), id, sectionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.UpdateStorefrontSectionBySectionId``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontSectionBySectionId`: StorefrontSectionResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.UpdateStorefrontSectionBySectionId`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 
**sectionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontSectionBySectionIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**StorefrontSectionResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontSectionReorderById

> StorefrontResponse4 UpdateStorefrontSectionReorderById(ctx, id).Execute()

Set a store's section order in one atomic call: every section id, top to bottom.

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
    resp, r, err := apiClient.StorefrontsApi.UpdateStorefrontSectionReorderById(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.UpdateStorefrontSectionReorderById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontSectionReorderById`: StorefrontResponse4
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.UpdateStorefrontSectionReorderById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontSectionReorderByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontResponse4**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontSetting

> UpdateStorefrontSettingResponse UpdateStorefrontSetting(ctx, id).Execute()

Update a store's settings. Omitted fields are unchanged, null clears. Policies replace wholesale. shippingPresetId / returnPresetId / paymentPresetId pick the store's presets. The checkout-term fields (offersFreeShipping … arrivalConditionWindowHours) are DEPRECATED aliases: they are written into the store's preset.

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
    resp, r, err := apiClient.StorefrontsApi.UpdateStorefrontSetting(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.UpdateStorefrontSetting``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontSetting`: UpdateStorefrontSettingResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.UpdateStorefrontSetting`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontSettingRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateStorefrontSettingResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontSlug

> StorefrontResponse UpdateStorefrontSlug(ctx, id).Execute()

Change a store's slug. The old slug keeps redirecting to the store.

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
    resp, r, err := apiClient.StorefrontsApi.UpdateStorefrontSlug(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.UpdateStorefrontSlug``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontSlug`: StorefrontResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.UpdateStorefrontSlug`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontSlugRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontStatus

> StorefrontResponse UpdateStorefrontStatus(ctx, id).Execute()

Set a store's status: draft | live | paused | archived. A live store is public once it has a live listing.

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
    resp, r, err := apiClient.StorefrontsApi.UpdateStorefrontStatus(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.UpdateStorefrontStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontStatus`: StorefrontResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.UpdateStorefrontStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**StorefrontResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateStorefrontTeamAccess

> OkResponse UpdateStorefrontTeamAccess(ctx).Execute()

Grant (or replace) a team member's access to one store, or to every store (storefrontId null), with scopes.

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
    resp, r, err := apiClient.StorefrontsApi.UpdateStorefrontTeamAccess(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `StorefrontsApi.UpdateStorefrontTeamAccess``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateStorefrontTeamAccess`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `StorefrontsApi.UpdateStorefrontTeamAccess`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStorefrontTeamAccessRequest struct via the builder pattern


### Return type

**OkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

