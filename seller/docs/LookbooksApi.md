# \LookbooksApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateLookbook**](LookbooksApi.md#CreateLookbook) | **Post** /v1/lookbooks | Create a lookbook. Pass templateId (classic, editorial, hero, magazine, drop, sale, postcard) to start it from a template styled with your storefront brand kit.
[**CreateLookbookCampaign**](LookbooksApi.md#CreateLookbookCampaign) | **Post** /v1/lookbook-campaigns | Email a lookbook to customers now or at scheduledAt. Audience kind subscribers | all | repeat | segment {filters} | handles {handles}; only people with marketing consent who have not unsubscribed receive it.
[**CreateLookbookCampaignCancel**](LookbooksApi.md#CreateLookbookCampaignCancel) | **Post** /v1/lookbook-campaigns/{id}/cancel | Cancel a scheduled campaign or stop one that is sending (emails already sent stay sent).
[**CreateLookbookCampaignPreview**](LookbooksApi.md#CreateLookbookCampaignPreview) | **Post** /v1/lookbook-campaigns/preview | Preview a lookbook email campaign without sending: how many of the audience can receive it (opted in, not suppressed) and the rendered HTML/text.
[**CreateLookbookShareLink**](LookbooksApi.md#CreateLookbookShareLink) | **Post** /v1/lookbooks/{id}/share-links | Make a share link for someone without a Crossly login: role view, comment or edit; expires after ttlHours (1 to 2160, default 168). The url is shown once.
[**CreatePrintMailJob**](LookbooksApi.md#CreatePrintMailJob) | **Post** /v1/print-mail/jobs | Draft a printed lookbook mailing (beta): renders the print-ready PDF with bleed and estimates the cost from your per-piece price. Mails nothing until confirmed.
[**CreatePrintMailJobConfirm**](LookbooksApi.md#CreatePrintMailJobConfirm) | **Post** /v1/print-mail/jobs/{id}/confirm | Confirm a drafted print job: acknowledgeCharges must be true and pieces must equal the job&#39;s piece count. Your Lob or PostGrid account is billed.
[**GetLookbookCampaign**](LookbooksApi.md#GetLookbookCampaign) | **Get** /v1/lookbook-campaigns/{id} | One lookbook campaign and its counts.
[**GetLookbookDesign**](LookbooksApi.md#GetLookbookDesign) | **Get** /v1/lookbooks/{id}/design | A lookbook design document (pages of text, photo, listing, shape and sticker elements) with live listing title/price/photo, your brand kit and rendered page images.
[**ListEmailProviders**](LookbooksApi.md#ListEmailProviders) | **Get** /v1/email-providers | Your BYOK email providers (SMTP, Mailgun, SendGrid, SES, Resend, Postmark, Brevo, Mailjet, SparkPost, MailerSend, Mandrill, Plunk and more): label, from address, default, last test. Secrets are never returned.
[**ListLookbookCampaignRecipients**](LookbooksApi.md#ListLookbookCampaignRecipients) | **Get** /v1/lookbook-campaigns/{id}/recipients | Per-recipient results for a campaign: status (queued, sent, failed, skipped + reason), provider message id, opens, clicks, bounce, unsubscribe.
[**ListLookbookCampaigns**](LookbooksApi.md#ListLookbookCampaigns) | **Get** /v1/lookbook-campaigns | Your lookbook campaigns, newest first, with sent / failed / skipped / opened / clicked / bounced / unsubscribed counts.
[**ListLookbooks**](LookbooksApi.md#ListLookbooks) | **Get** /v1/lookbooks | Your lookbooks: id, slug, title, published, current design version and item count.
[**ListPrintMailJobs**](LookbooksApi.md#ListPrintMailJobs) | **Get** /v1/print-mail/jobs | Print-and-mail jobs (beta): provider, format, pieces, estimate, status and per-piece results.
[**UpdateLookbookDesign**](LookbooksApi.md#UpdateLookbookDesign) | **Put** /v1/lookbooks/{id}/design | Save a new design version. Images must already be on the Crossly CDN; listings must be your own Crossly listings. Pass baseVersion to get a 409 instead of overwriting a newer save.



## CreateLookbook

> CreateLookbookResponse CreateLookbook(ctx).Execute()

Create a lookbook. Pass templateId (classic, editorial, hero, magazine, drop, sale, postcard) to start it from a template styled with your storefront brand kit.

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
    resp, r, err := apiClient.LookbooksApi.CreateLookbook(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.CreateLookbook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLookbook`: CreateLookbookResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.CreateLookbook`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLookbookRequest struct via the builder pattern


### Return type

**CreateLookbookResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLookbookCampaign

> LookbookCampaignResponse CreateLookbookCampaign(ctx).Execute()

Email a lookbook to customers now or at scheduledAt. Audience kind subscribers | all | repeat | segment {filters} | handles {handles}; only people with marketing consent who have not unsubscribed receive it.

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
    resp, r, err := apiClient.LookbooksApi.CreateLookbookCampaign(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.CreateLookbookCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLookbookCampaign`: LookbookCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.CreateLookbookCampaign`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLookbookCampaignRequest struct via the builder pattern


### Return type

**LookbookCampaignResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLookbookCampaignCancel

> LookbookCampaignResponse CreateLookbookCampaignCancel(ctx, id).Execute()

Cancel a scheduled campaign or stop one that is sending (emails already sent stay sent).

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
    resp, r, err := apiClient.LookbooksApi.CreateLookbookCampaignCancel(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.CreateLookbookCampaignCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLookbookCampaignCancel`: LookbookCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.CreateLookbookCampaignCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLookbookCampaignCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**LookbookCampaignResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLookbookCampaignPreview

> CreateLookbookCampaignPreviewResponse CreateLookbookCampaignPreview(ctx).Execute()

Preview a lookbook email campaign without sending: how many of the audience can receive it (opted in, not suppressed) and the rendered HTML/text.

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
    resp, r, err := apiClient.LookbooksApi.CreateLookbookCampaignPreview(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.CreateLookbookCampaignPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLookbookCampaignPreview`: CreateLookbookCampaignPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.CreateLookbookCampaignPreview`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLookbookCampaignPreviewRequest struct via the builder pattern


### Return type

**CreateLookbookCampaignPreviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateLookbookShareLink

> CreateLookbookShareLinkResponse CreateLookbookShareLink(ctx, id).Execute()

Make a share link for someone without a Crossly login: role view, comment or edit; expires after ttlHours (1 to 2160, default 168). The url is shown once.

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
    resp, r, err := apiClient.LookbooksApi.CreateLookbookShareLink(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.CreateLookbookShareLink``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateLookbookShareLink`: CreateLookbookShareLinkResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.CreateLookbookShareLink`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLookbookShareLinkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateLookbookShareLinkResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreatePrintMailJob

> CreatePrintMailJobResponse CreatePrintMailJob(ctx).Execute()

Draft a printed lookbook mailing (beta): renders the print-ready PDF with bleed and estimates the cost from your per-piece price. Mails nothing until confirmed.

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
    resp, r, err := apiClient.LookbooksApi.CreatePrintMailJob(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.CreatePrintMailJob``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrintMailJob`: CreatePrintMailJobResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.CreatePrintMailJob`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrintMailJobRequest struct via the builder pattern


### Return type

**CreatePrintMailJobResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreatePrintMailJobConfirm

> CreatePrintMailJobConfirmResponse CreatePrintMailJobConfirm(ctx, id).Execute()

Confirm a drafted print job: acknowledgeCharges must be true and pieces must equal the job's piece count. Your Lob or PostGrid account is billed.

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
    resp, r, err := apiClient.LookbooksApi.CreatePrintMailJobConfirm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.CreatePrintMailJobConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePrintMailJobConfirm`: CreatePrintMailJobConfirmResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.CreatePrintMailJobConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePrintMailJobConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreatePrintMailJobConfirmResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetLookbookCampaign

> LookbookCampaignResponse GetLookbookCampaign(ctx, id).Execute()

One lookbook campaign and its counts.

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
    resp, r, err := apiClient.LookbooksApi.GetLookbookCampaign(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.GetLookbookCampaign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLookbookCampaign`: LookbookCampaignResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.GetLookbookCampaign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLookbookCampaignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**LookbookCampaignResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetLookbookDesign

> GetLookbookDesignResponse GetLookbookDesign(ctx, id).Execute()

A lookbook design document (pages of text, photo, listing, shape and sticker elements) with live listing title/price/photo, your brand kit and rendered page images.

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
    resp, r, err := apiClient.LookbooksApi.GetLookbookDesign(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.GetLookbookDesign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLookbookDesign`: GetLookbookDesignResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.GetLookbookDesign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLookbookDesignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetLookbookDesignResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListEmailProviders

> ListEmailProvidersResponse ListEmailProviders(ctx).Execute()

Your BYOK email providers (SMTP, Mailgun, SendGrid, SES, Resend, Postmark, Brevo, Mailjet, SparkPost, MailerSend, Mandrill, Plunk and more): label, from address, default, last test. Secrets are never returned.

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
    resp, r, err := apiClient.LookbooksApi.ListEmailProviders(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.ListEmailProviders``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailProviders`: ListEmailProvidersResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.ListEmailProviders`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListEmailProvidersRequest struct via the builder pattern


### Return type

**ListEmailProvidersResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListLookbookCampaignRecipients

> ListLookbookCampaignRecipientsResponse ListLookbookCampaignRecipients(ctx, id).Execute()

Per-recipient results for a campaign: status (queued, sent, failed, skipped + reason), provider message id, opens, clicks, bounce, unsubscribe.

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
    resp, r, err := apiClient.LookbooksApi.ListLookbookCampaignRecipients(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.ListLookbookCampaignRecipients``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLookbookCampaignRecipients`: ListLookbookCampaignRecipientsResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.ListLookbookCampaignRecipients`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListLookbookCampaignRecipientsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListLookbookCampaignRecipientsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListLookbookCampaigns

> ListLookbookCampaignsResponse ListLookbookCampaigns(ctx).Execute()

Your lookbook campaigns, newest first, with sent / failed / skipped / opened / clicked / bounced / unsubscribed counts.

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
    resp, r, err := apiClient.LookbooksApi.ListLookbookCampaigns(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.ListLookbookCampaigns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLookbookCampaigns`: ListLookbookCampaignsResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.ListLookbookCampaigns`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLookbookCampaignsRequest struct via the builder pattern


### Return type

**ListLookbookCampaignsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListLookbooks

> ListLookbooksResponse ListLookbooks(ctx).Execute()

Your lookbooks: id, slug, title, published, current design version and item count.

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
    resp, r, err := apiClient.LookbooksApi.ListLookbooks(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.ListLookbooks``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListLookbooks`: ListLookbooksResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.ListLookbooks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListLookbooksRequest struct via the builder pattern


### Return type

**ListLookbooksResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListPrintMailJobs

> ListPrintMailJobsResponse ListPrintMailJobs(ctx).Execute()

Print-and-mail jobs (beta): provider, format, pieces, estimate, status and per-piece results.

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
    resp, r, err := apiClient.LookbooksApi.ListPrintMailJobs(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.ListPrintMailJobs``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListPrintMailJobs`: ListPrintMailJobsResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.ListPrintMailJobs`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListPrintMailJobsRequest struct via the builder pattern


### Return type

**ListPrintMailJobsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateLookbookDesign

> UpdateLookbookDesignResponse UpdateLookbookDesign(ctx, id).Execute()

Save a new design version. Images must already be on the Crossly CDN; listings must be your own Crossly listings. Pass baseVersion to get a 409 instead of overwriting a newer save.

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
    resp, r, err := apiClient.LookbooksApi.UpdateLookbookDesign(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `LookbooksApi.UpdateLookbookDesign``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateLookbookDesign`: UpdateLookbookDesignResponse
    fmt.Fprintf(os.Stdout, "Response from `LookbooksApi.UpdateLookbookDesign`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLookbookDesignRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**UpdateLookbookDesignResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

