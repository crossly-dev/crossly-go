# \ActivityApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateActionLogExport**](ActivityApi.md#CreateActionLogExport) | **Post** /v1/action-log/exports | Queue an export of the Activity Log: every event (csv, ndjson, pdf) or a summary report (pdf, ndjson). Poll GET /api/v1/action-log/exports/{exportId}, then download it. Exports never contain IP addresses.
[**CreateActionLogSearchTranslate**](ActivityApi.md#CreateActionLogSearchTranslate) | **Post** /v1/action-log/search/translate | Turn a question in words (\&quot;failed delists on Mercari last week\&quot;) into an Activity Log search: pass the returned &#x60;q&#x60; to GET /api/v1/action-log (or the histogram and csv/ndjson exports). Nothing is run. A 404 with details.code &#x60;not_enabled&#x60; means natural-language search is off for this account.
[**CreateDetectedEventEventConfirm**](ActivityApi.md#CreateDetectedEventEventConfirm) | **Post** /v1/detected-events/events/{id}/confirm | Answer an open ask (canConfirm true): optionally which item it is (itemId, one of the match candidates) and which actions to run instead of the ask&#39;s preselected ones. Returns the event. 409 ask_not_open when the event is not waiting.
[**CreateDetectedEventEventUndo**](ActivityApi.md#CreateDetectedEventEventUndo) | **Post** /v1/detected-events/events/{id}/undo | Undo what Crossly did for an event (canUndo true): every executed action&#39;s stored inverse runs, newest first (restore stock, relist, unlink…). Returns the event. 409 nothing_to_undo when nothing is reversible.
[**CreateDetectedEventRule**](ActivityApi.md#CreateDetectedEventRule) | **Post** /v1/detected-events/rules | Create a detection rule: which detected events it applies to (filter + optional condition over any event axis) and what Crossly does then (and else). Without order it is placed most-specific-first, above the default set. Returns the stored rule with its id.
[**CreateDetectedEventRulePlace**](ActivityApi.md#CreateDetectedEventRulePlace) | **Post** /v1/detected-events/rules/place | Save one rule and the whole evaluation order in one transaction: rule.id is an existing id (replace) or any new client id (create); ids is every rule id top-down, naming the rule by that id. Returns the stored rule and the list.
[**CreateDetectedEventRuleReorder**](ActivityApi.md#CreateDetectedEventRuleReorder) | **Post** /v1/detected-events/rules/reorder | Set the evaluation order: ids top-down, defaults included (ids from the rules list). Rules are re-numbered 10, 20, 30…; a default moved above one of your rules is stored as your edit. Returns the list.
[**CreateDetectedEventRuleResetDefault**](ActivityApi.md#CreateDetectedEventRuleResetDefault) | **Post** /v1/detected-events/rules/reset-defaults | Drop your edits of the shipped default rules (keys, or all when omitted). Returns the list.
[**CreateDetectedEventSimulate**](ActivityApi.md#CreateDetectedEventSimulate) | **Post** /v1/detected-events/simulate | Run one event through the rules without acting: the winning rule, rules it shadowed, the trace, the guards that held actions back and the plan per target with how it would run (direct / relay / ask / skip). event needs only kind; every axis left out takes its default. rules &#x3D; an unsaved draft list.
[**CreateDetectedEventSimulateImpact**](ActivityApi.md#CreateDetectedEventSimulateImpact) | **Post** /v1/detected-events/simulate/impact | What a rule edit would change: your detected events of the last days (default 30, at most 500) re-resolved against the draft rules and against the stored ones. Per event: winning rule and actions before and after, and whether it changed. Nothing runs.
[**CreateEmailEventConfirm**](ActivityApi.md#CreateEmailEventConfirm) | **Post** /v1/email-events/{id}/confirm | Confirm an email-detected event: re-run it as certain (a possible sale records the sale and delists the item elsewhere). When several listings could be meant, pass platformListingDbId to choose one.
[**CreateEmailEventDismiss**](ActivityApi.md#CreateEmailEventDismiss) | **Post** /v1/email-events/{id}/dismiss | Dismiss an email-detected event (\&quot;not a sale\&quot;): closes the card without acting.
[**CreateEventPolicyDecisionConfirm**](ActivityApi.md#CreateEventPolicyDecisionConfirm) | **Post** /v1/event-policies/decisions/{id}/confirm | Confirm a review decision from a webhook, order poller or app relay (canConfirm true): the event runs again with the seller confirmed, so a review rule runs the strongest allowed non-review action for that cell. Returns the new decision. 409 when the decision was already confirmed or is not confirmable (email review decisions are confirmed on their review card).
[**DeleteDetectedEventRule**](ActivityApi.md#DeleteDetectedEventRule) | **Delete** /v1/detected-events/rules/{id} | Delete a detection rule. Deleting your edit of a default (its id, or default:&lt;key&gt;) brings the shipped default back; deleting a translated event-policy rule deletes that policy rule too.
[**DeleteEventPolicyRule**](ActivityApi.md#DeleteEventPolicyRule) | **Delete** /v1/event-policies/rules/{id} | Delete a rule: its cell falls back to the next scope (platform → global → the built-in default).
[**GetActionLog**](ActivityApi.md#GetActionLog) | **Get** /v1/action-log/{id} | One Activity Log event by id, with &#x60;view&#x60; including its steps and, for bulk jobs, its children.
[**GetActionLogExport**](ActivityApi.md#GetActionLogExport) | **Get** /v1/action-log/exports/{exportId} | The state of one Activity Log export.
[**GetActionLogExportDownload**](ActivityApi.md#GetActionLogExportDownload) | **Get** /v1/action-log/exports/{exportId}/download | Download a finished Activity Log export (for 24 hours after it completes).
[**GetActionLogFacet**](ActivityApi.md#GetActionLogFacet) | **Get** /v1/action-log/facets | Filter options of the caller&#39;s Activity Log (last 90 days): areas with their actions and counts, the team members who acted, platforms and outcomes.
[**GetDetectedEventEvent**](ActivityApi.md#GetDetectedEventEvent) | **Get** /v1/detected-events/events/{id} | One detected event with every report of it (first + corroborations, with correlation and contradiction), the guards that held actions back, the plan and each executed action&#39;s outcome and reversibility.
[**GetEventPolicy**](ActivityApi.md#GetEventPolicy) | **Get** /v1/event-policies | What Crossly does with each detected marketplace event (sale, listing ended, offer, cancellation, label, shipped, delivered, refund, return, payout, message), per who owns the item (crossly, imported, app_managed, other_tool, untracked). Each matrix cell lists the allowed actions and the effective one, with the scope it comes from (app &gt; platform &gt; global &gt; default). Pass platform (or appId) to see the matrix as that platform&#39;s (or app&#39;s) items see it. Also returns the rules, platforms and connected apps.
[**ListActionLog**](ActivityApi.md#ListActionLog) | **Get** /v1/action-log | List Activity Log events, newest first. Each event carries &#x60;view&#x60; (title, sentence, fields, failure advice, who did it) and typed &#x60;detail&#x60;. Page with &#x60;cursor&#x60;; &#x60;pagination.total&#x60; counts matches.
[**ListActionLogCalls**](ActivityApi.md#ListActionLogCalls) | **Get** /v1/action-log/{id}/calls | The outbound platform HTTP calls under an event (oldest first) — platform, operation, method, host, status, latency, error class/message, proxy + recipe id. Answers \&quot;what was sent / what went wrong\&quot;. The URL path, request/response bodies and persisted-query hash are never returned.
[**ListActionLogHistogram**](ActivityApi.md#ListActionLogHistogram) | **Get** /v1/action-log/histogram | Events per UTC day for the list filters (default: the last 30 days; at most 366).
[**ListDetectedEventEvents**](ActivityApi.md#ListDetectedEventEvents) | **Get** /v1/detected-events/events | Every event Crossly detected (sales, order changes, listing changes, offers, messages, account notices), newest first: all its axes, the rule that decided it, its corroborating reports, the plan per target and how each action went, with canConfirm (an open ask) and canUndo. Cursor-paged; filter by kinds, platforms statuses and accounts (comma-separated) and since (ISO time).
[**ListDetectedEventRules**](ActivityApi.md#ListDetectedEventRules) | **Get** /v1/detected-events/rules | The seller&#39;s detection rules, top-down in the order they are evaluated (first match wins unless a rule says continue): the seller&#39;s own rules, translations of their event-policy rules (source legacy), then the shipped default set (source default, ids default:&lt;key&gt;; a stored edit of a default replaces it). Each rule has a filter (event kinds, platforms, accounts, apps, tags, items), an optional condition over any event axis, the actions it runs (then / else) and its guard loosenings.
[**ListEmailEvents**](ActivityApi.md#ListEmailEvents) | **Get** /v1/email-events | List email-detected events waiting for the seller (status&#x3D;needs_review): a possible sale, ended listing, cancellation or other event an email reported that was not certain enough to act on automatically.
[**ListEventPolicyDecisions**](ActivityApi.md#ListEventPolicyDecisions) | **Get** /v1/event-policies/decisions | The decision log, newest first: for each detected event, what the rule or default asked for, what was done (after capability fallback and the email confidence ceiling), the rule that decided it, and the outcome (done / skipped / failed / review). Cursor-paged; filter by event, ownership, platform or outcome (outcome&#x3D;review: still waiting for the seller). Kept 90 days.
[**UpdateDetectedEventRule**](ActivityApi.md#UpdateDetectedEventRule) | **Put** /v1/detected-events/rules/{id} | Replace a detection rule. An id default:&lt;key&gt; stores your edit of that shipped default (delete it, or reset-defaults, to go back to the shipped one).
[**UpdateEventPolicyRule**](ActivityApi.md#UpdateEventPolicyRule) | **Put** /v1/event-policies/rules | Set what Crossly does with one kind of detected event for one kind of item: scope global (scopeValue null), platform (scopeValue &#x3D; platform id) or app (scopeValue &#x3D; OAuth app id, app_managed items only). Replaces the rule for the same cell and scope. The action must be one the matrix allows for that cell; import actions exist only for untracked items.



## CreateActionLogExport

> GetActionLogExportResponse CreateActionLogExport(ctx).CreateActionLogExportRequest(createActionLogExportRequest).Execute()

Queue an export of the Activity Log: every event (csv, ndjson, pdf) or a summary report (pdf, ndjson). Poll GET /api/v1/action-log/exports/{exportId}, then download it. Exports never contain IP addresses.

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
    createActionLogExportRequest := *openapiclient.NewCreateActionLogExportRequest("Format_example") // CreateActionLogExportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.CreateActionLogExport(context.Background()).CreateActionLogExportRequest(createActionLogExportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.CreateActionLogExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateActionLogExport`: GetActionLogExportResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.CreateActionLogExport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateActionLogExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createActionLogExportRequest** | **CreateActionLogExportRequest** |  | 

### Return type

**GetActionLogExportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateActionLogSearchTranslate

> CreateActionLogSearchTranslateResponse CreateActionLogSearchTranslate(ctx).CreateActionLogSearchTranslateRequest(createActionLogSearchTranslateRequest).Execute()

Turn a question in words (\"failed delists on Mercari last week\") into an Activity Log search: pass the returned `q` to GET /api/v1/action-log (or the histogram and csv/ndjson exports). Nothing is run. A 404 with details.code `not_enabled` means natural-language search is off for this account.

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
    createActionLogSearchTranslateRequest := *openapiclient.NewCreateActionLogSearchTranslateRequest("Text_example") // CreateActionLogSearchTranslateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.CreateActionLogSearchTranslate(context.Background()).CreateActionLogSearchTranslateRequest(createActionLogSearchTranslateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.CreateActionLogSearchTranslate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateActionLogSearchTranslate`: CreateActionLogSearchTranslateResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.CreateActionLogSearchTranslate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateActionLogSearchTranslateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createActionLogSearchTranslateRequest** | **CreateActionLogSearchTranslateRequest** |  | 

### Return type

**CreateActionLogSearchTranslateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDetectedEventEventConfirm

> CreateDetectedEventEventUndoResponse CreateDetectedEventEventConfirm(ctx, id).CreateDetectedEventEventConfirmRequest(createDetectedEventEventConfirmRequest).Execute()

Answer an open ask (canConfirm true): optionally which item it is (itemId, one of the match candidates) and which actions to run instead of the ask's preselected ones. Returns the event. 409 ask_not_open when the event is not waiting.

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
    createDetectedEventEventConfirmRequest := *openapiclient.NewCreateDetectedEventEventConfirmRequest() // CreateDetectedEventEventConfirmRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.CreateDetectedEventEventConfirm(context.Background(), id).CreateDetectedEventEventConfirmRequest(createDetectedEventEventConfirmRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.CreateDetectedEventEventConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDetectedEventEventConfirm`: CreateDetectedEventEventUndoResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.CreateDetectedEventEventConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDetectedEventEventConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDetectedEventEventConfirmRequest** | **CreateDetectedEventEventConfirmRequest** |  | 

### Return type

**CreateDetectedEventEventUndoResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDetectedEventEventUndo

> CreateDetectedEventEventUndoResponse CreateDetectedEventEventUndo(ctx, id).Execute()

Undo what Crossly did for an event (canUndo true): every executed action's stored inverse runs, newest first (restore stock, relist, unlink…). Returns the event. 409 nothing_to_undo when nothing is reversible.

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
    resp, r, err := apiClient.ActivityApi.CreateDetectedEventEventUndo(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.CreateDetectedEventEventUndo``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDetectedEventEventUndo`: CreateDetectedEventEventUndoResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.CreateDetectedEventEventUndo`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDetectedEventEventUndoRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateDetectedEventEventUndoResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDetectedEventRule

> CreateDetectedEventRuleResponse CreateDetectedEventRule(ctx).CreateDetectedEventRuleRequest(createDetectedEventRuleRequest).Execute()

Create a detection rule: which detected events it applies to (filter + optional condition over any event axis) and what Crossly does then (and else). Without order it is placed most-specific-first, above the default set. Returns the stored rule with its id.

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
    createDetectedEventRuleRequest := *openapiclient.NewCreateDetectedEventRuleRequest("Name_example", map[string]interface{}{"key": interface{}(123)}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // CreateDetectedEventRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.CreateDetectedEventRule(context.Background()).CreateDetectedEventRuleRequest(createDetectedEventRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.CreateDetectedEventRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDetectedEventRule`: CreateDetectedEventRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.CreateDetectedEventRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDetectedEventRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDetectedEventRuleRequest** | **CreateDetectedEventRuleRequest** |  | 

### Return type

**CreateDetectedEventRuleResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDetectedEventRulePlace

> CreateDetectedEventRulePlaceResponse CreateDetectedEventRulePlace(ctx).CreateDetectedEventRulePlaceRequest(createDetectedEventRulePlaceRequest).Execute()

Save one rule and the whole evaluation order in one transaction: rule.id is an existing id (replace) or any new client id (create); ids is every rule id top-down, naming the rule by that id. Returns the stored rule and the list.

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
    createDetectedEventRulePlaceRequest := *openapiclient.NewCreateDetectedEventRulePlaceRequest(*openapiclient.NewCreateDetectedEventRuleRequest("Name_example", map[string]interface{}{"key": interface{}(123)}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}), []string{"Ids_example"}) // CreateDetectedEventRulePlaceRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.CreateDetectedEventRulePlace(context.Background()).CreateDetectedEventRulePlaceRequest(createDetectedEventRulePlaceRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.CreateDetectedEventRulePlace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDetectedEventRulePlace`: CreateDetectedEventRulePlaceResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.CreateDetectedEventRulePlace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDetectedEventRulePlaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDetectedEventRulePlaceRequest** | **CreateDetectedEventRulePlaceRequest** |  | 

### Return type

**CreateDetectedEventRulePlaceResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDetectedEventRuleReorder

> CreateDetectedEventRuleReorderResponse CreateDetectedEventRuleReorder(ctx).CreateDetectedEventRuleReorderRequest(createDetectedEventRuleReorderRequest).Execute()

Set the evaluation order: ids top-down, defaults included (ids from the rules list). Rules are re-numbered 10, 20, 30…; a default moved above one of your rules is stored as your edit. Returns the list.

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
    createDetectedEventRuleReorderRequest := *openapiclient.NewCreateDetectedEventRuleReorderRequest([]string{"Ids_example"}) // CreateDetectedEventRuleReorderRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.CreateDetectedEventRuleReorder(context.Background()).CreateDetectedEventRuleReorderRequest(createDetectedEventRuleReorderRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.CreateDetectedEventRuleReorder``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDetectedEventRuleReorder`: CreateDetectedEventRuleReorderResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.CreateDetectedEventRuleReorder`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDetectedEventRuleReorderRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDetectedEventRuleReorderRequest** | **CreateDetectedEventRuleReorderRequest** |  | 

### Return type

**CreateDetectedEventRuleReorderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDetectedEventRuleResetDefault

> CreateDetectedEventRuleReorderResponse CreateDetectedEventRuleResetDefault(ctx).CreateDetectedEventRuleResetDefaultRequest(createDetectedEventRuleResetDefaultRequest).Execute()

Drop your edits of the shipped default rules (keys, or all when omitted). Returns the list.

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
    createDetectedEventRuleResetDefaultRequest := *openapiclient.NewCreateDetectedEventRuleResetDefaultRequest() // CreateDetectedEventRuleResetDefaultRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.CreateDetectedEventRuleResetDefault(context.Background()).CreateDetectedEventRuleResetDefaultRequest(createDetectedEventRuleResetDefaultRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.CreateDetectedEventRuleResetDefault``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDetectedEventRuleResetDefault`: CreateDetectedEventRuleReorderResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.CreateDetectedEventRuleResetDefault`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDetectedEventRuleResetDefaultRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDetectedEventRuleResetDefaultRequest** | **CreateDetectedEventRuleResetDefaultRequest** |  | 

### Return type

**CreateDetectedEventRuleReorderResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDetectedEventSimulate

> CreateDetectedEventSimulateResponse CreateDetectedEventSimulate(ctx).CreateDetectedEventSimulateRequest(createDetectedEventSimulateRequest).Execute()

Run one event through the rules without acting: the winning rule, rules it shadowed, the trace, the guards that held actions back and the plan per target with how it would run (direct / relay / ask / skip). event needs only kind; every axis left out takes its default. rules = an unsaved draft list.

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
    createDetectedEventSimulateRequest := *openapiclient.NewCreateDetectedEventSimulateRequest(map[string]interface{}{"key": interface{}(123)}) // CreateDetectedEventSimulateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.CreateDetectedEventSimulate(context.Background()).CreateDetectedEventSimulateRequest(createDetectedEventSimulateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.CreateDetectedEventSimulate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDetectedEventSimulate`: CreateDetectedEventSimulateResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.CreateDetectedEventSimulate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDetectedEventSimulateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDetectedEventSimulateRequest** | **CreateDetectedEventSimulateRequest** |  | 

### Return type

**CreateDetectedEventSimulateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateDetectedEventSimulateImpact

> CreateDetectedEventSimulateImpactResponse CreateDetectedEventSimulateImpact(ctx).CreateDetectedEventSimulateImpactRequest(createDetectedEventSimulateImpactRequest).Execute()

What a rule edit would change: your detected events of the last days (default 30, at most 500) re-resolved against the draft rules and against the stored ones. Per event: winning rule and actions before and after, and whether it changed. Nothing runs.

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
    createDetectedEventSimulateImpactRequest := *openapiclient.NewCreateDetectedEventSimulateImpactRequest([]map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // CreateDetectedEventSimulateImpactRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.CreateDetectedEventSimulateImpact(context.Background()).CreateDetectedEventSimulateImpactRequest(createDetectedEventSimulateImpactRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.CreateDetectedEventSimulateImpact``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDetectedEventSimulateImpact`: CreateDetectedEventSimulateImpactResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.CreateDetectedEventSimulateImpact`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDetectedEventSimulateImpactRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createDetectedEventSimulateImpactRequest** | **CreateDetectedEventSimulateImpactRequest** |  | 

### Return type

**CreateDetectedEventSimulateImpactResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateEmailEventConfirm

> CreateEmailEventConfirmResponse CreateEmailEventConfirm(ctx, id).CreateEmailEventConfirmRequest(createEmailEventConfirmRequest).Execute()

Confirm an email-detected event: re-run it as certain (a possible sale records the sale and delists the item elsewhere). When several listings could be meant, pass platformListingDbId to choose one.

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
    createEmailEventConfirmRequest := *openapiclient.NewCreateEmailEventConfirmRequest() // CreateEmailEventConfirmRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.CreateEmailEventConfirm(context.Background(), id).CreateEmailEventConfirmRequest(createEmailEventConfirmRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.CreateEmailEventConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailEventConfirm`: CreateEmailEventConfirmResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.CreateEmailEventConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailEventConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createEmailEventConfirmRequest** | **CreateEmailEventConfirmRequest** |  | 

### Return type

**CreateEmailEventConfirmResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateEmailEventDismiss

> IdStatusResponse CreateEmailEventDismiss(ctx, id).Execute()

Dismiss an email-detected event (\"not a sale\"): closes the card without acting.

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
    resp, r, err := apiClient.ActivityApi.CreateEmailEventDismiss(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.CreateEmailEventDismiss``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEmailEventDismiss`: IdStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.CreateEmailEventDismiss`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEmailEventDismissRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**IdStatusResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateEventPolicyDecisionConfirm

> CreateEventPolicyDecisionConfirmResponse CreateEventPolicyDecisionConfirm(ctx, id).Execute()

Confirm a review decision from a webhook, order poller or app relay (canConfirm true): the event runs again with the seller confirmed, so a review rule runs the strongest allowed non-review action for that cell. Returns the new decision. 409 when the decision was already confirmed or is not confirmable (email review decisions are confirmed on their review card).

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
    resp, r, err := apiClient.ActivityApi.CreateEventPolicyDecisionConfirm(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.CreateEventPolicyDecisionConfirm``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateEventPolicyDecisionConfirm`: CreateEventPolicyDecisionConfirmResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.CreateEventPolicyDecisionConfirm`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEventPolicyDecisionConfirmRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateEventPolicyDecisionConfirmResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteDetectedEventRule

> DeleteStoryResponse DeleteDetectedEventRule(ctx, id).Execute()

Delete a detection rule. Deleting your edit of a default (its id, or default:<key>) brings the shipped default back; deleting a translated event-policy rule deletes that policy rule too.

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
    resp, r, err := apiClient.ActivityApi.DeleteDetectedEventRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.DeleteDetectedEventRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteDetectedEventRule`: DeleteStoryResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.DeleteDetectedEventRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDetectedEventRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteStoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteEventPolicyRule

> DeleteStoryResponse DeleteEventPolicyRule(ctx, id).Execute()

Delete a rule: its cell falls back to the next scope (platform → global → the built-in default).

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
    resp, r, err := apiClient.ActivityApi.DeleteEventPolicyRule(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.DeleteEventPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteEventPolicyRule`: DeleteStoryResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.DeleteEventPolicyRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEventPolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteStoryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetActionLog

> ListActionLogItem GetActionLog(ctx, id).Execute()

One Activity Log event by id, with `view` including its steps and, for bulk jobs, its children.

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
    resp, r, err := apiClient.ActivityApi.GetActionLog(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GetActionLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActionLog`: ListActionLogItem
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GetActionLog`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActionLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListActionLogItem**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetActionLogExport

> GetActionLogExportResponse GetActionLogExport(ctx, exportId).Execute()

The state of one Activity Log export.

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
    exportId := "exportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.GetActionLogExport(context.Background(), exportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GetActionLogExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActionLogExport`: GetActionLogExportResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GetActionLogExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActionLogExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetActionLogExportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetActionLogExportDownload

> GetActionLogExportDownloadResponse GetActionLogExportDownload(ctx, exportId).Execute()

Download a finished Activity Log export (for 24 hours after it completes).

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
    exportId := "exportId_example" // string | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.GetActionLogExportDownload(context.Background(), exportId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GetActionLogExportDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActionLogExportDownload`: GetActionLogExportDownloadResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GetActionLogExportDownload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetActionLogExportDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetActionLogExportDownloadResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetActionLogFacet

> GetActionLogFacetResponse GetActionLogFacet(ctx).Execute()

Filter options of the caller's Activity Log (last 90 days): areas with their actions and counts, the team members who acted, platforms and outcomes.

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
    resp, r, err := apiClient.ActivityApi.GetActionLogFacet(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GetActionLogFacet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetActionLogFacet`: GetActionLogFacetResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GetActionLogFacet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetActionLogFacetRequest struct via the builder pattern


### Return type

**GetActionLogFacetResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetDetectedEventEvent

> GetDetectedEventEventResponse GetDetectedEventEvent(ctx, id).Execute()

One detected event with every report of it (first + corroborations, with correlation and contradiction), the guards that held actions back, the plan and each executed action's outcome and reversibility.

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
    resp, r, err := apiClient.ActivityApi.GetDetectedEventEvent(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GetDetectedEventEvent``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDetectedEventEvent`: GetDetectedEventEventResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GetDetectedEventEvent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDetectedEventEventRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetDetectedEventEventResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetEventPolicy

> GetEventPolicyResponse GetEventPolicy(ctx).Platform(platform).AppId(appId).Execute()

What Crossly does with each detected marketplace event (sale, listing ended, offer, cancellation, label, shipped, delivered, refund, return, payout, message), per who owns the item (crossly, imported, app_managed, other_tool, untracked). Each matrix cell lists the allowed actions and the effective one, with the scope it comes from (app > platform > global > default). Pass platform (or appId) to see the matrix as that platform's (or app's) items see it. Also returns the rules, platforms and connected apps.

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
    platform := "platform_example" // string |  (optional)
    appId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.GetEventPolicy(context.Background()).Platform(platform).AppId(appId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.GetEventPolicy``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetEventPolicy`: GetEventPolicyResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.GetEventPolicy`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetEventPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platform** | **string** |  | 
 **appId** | **string** |  | 

### Return type

**GetEventPolicyResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListActionLog

> ListActionLogResponse ListActionLog(ctx).Platform(platform).Action(action).Category(category).Area(area).Status(status).Source(source).TargetType(targetType).TargetId(targetId).ActorUserId(actorUserId).Q(q).Since(since).Until(until).At(at).Cursor(cursor).Limit(limit).Offset(offset).Execute()

List Activity Log events, newest first. Each event carries `view` (title, sentence, fields, failure advice, who did it) and typed `detail`. Page with `cursor`; `pagination.total` counts matches.

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    platform := "platform_example" // string | Platform id (`ebay`, `poshmark`, ...). (optional)
    action := "action_example" // string | A canonical action key (`listing.delist`, see the public action catalog) or a legacy action word. (optional)
    category := "category_example" // string |  (optional)
    area := "area_example" // string | A seller area: every action of that area (the grouped filter of the dashboard). (optional)
    status := "status_example" // string |  (optional)
    source := "source_example" // string |  (optional)
    targetType := "targetType_example" // string | With targetId: events about that object (any of their targets). (optional)
    targetId := "targetId_example" // string |  (optional)
    actorUserId := "actorUserId_example" // string | Only what one team member did. (optional)
    q := "q_example" // string | A search in the activity query language (`action:listing.* -outcome:failure \"free text\"`), ANDed with the other filters. POST /api/v1/action-log/search/translate turns a question in words into one. (optional)
    since := time.Now() // time.Time | Inclusive lower bound on when the action happened. (optional)
    until := time.Now() // time.Time | Exclusive upper bound. (optional)
    at := time.Now() // time.Time | Jump to a time: the page starts at the newest event at or before it. (optional)
    cursor := "cursor_example" // string | `pagination.nextCursor` of the previous page. (optional)
    limit := int32(56) // int32 |  (optional) (default to 50)
    offset := int32(56) // int32 | Deprecated: use cursor. offset + limit may reach at most 1000 events. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ListActionLog(context.Background()).Platform(platform).Action(action).Category(category).Area(area).Status(status).Source(source).TargetType(targetType).TargetId(targetId).ActorUserId(actorUserId).Q(q).Since(since).Until(until).At(at).Cursor(cursor).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ListActionLog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActionLog`: ListActionLogResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ListActionLog`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListActionLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platform** | **string** | Platform id (&#x60;ebay&#x60;, &#x60;poshmark&#x60;, ...). | 
 **action** | **string** | A canonical action key (&#x60;listing.delist&#x60;, see the public action catalog) or a legacy action word. | 
 **category** | **string** |  | 
 **area** | **string** | A seller area: every action of that area (the grouped filter of the dashboard). | 
 **status** | **string** |  | 
 **source** | **string** |  | 
 **targetType** | **string** | With targetId: events about that object (any of their targets). | 
 **targetId** | **string** |  | 
 **actorUserId** | **string** | Only what one team member did. | 
 **q** | **string** | A search in the activity query language (&#x60;action:listing.* -outcome:failure \&quot;free text\&quot;&#x60;), ANDed with the other filters. POST /api/v1/action-log/search/translate turns a question in words into one. | 
 **since** | **time.Time** | Inclusive lower bound on when the action happened. | 
 **until** | **time.Time** | Exclusive upper bound. | 
 **at** | **time.Time** | Jump to a time: the page starts at the newest event at or before it. | 
 **cursor** | **string** | &#x60;pagination.nextCursor&#x60; of the previous page. | 
 **limit** | **int32** |  | [default to 50]
 **offset** | **int32** | Deprecated: use cursor. offset + limit may reach at most 1000 events. | 

### Return type

**ListActionLogResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListActionLogCalls

> ListActionLogCallsResponse ListActionLogCalls(ctx, id).Execute()

The outbound platform HTTP calls under an event (oldest first) — platform, operation, method, host, status, latency, error class/message, proxy + recipe id. Answers \"what was sent / what went wrong\". The URL path, request/response bodies and persisted-query hash are never returned.

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
    resp, r, err := apiClient.ActivityApi.ListActionLogCalls(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ListActionLogCalls``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActionLogCalls`: ListActionLogCallsResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ListActionLogCalls`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListActionLogCallsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListActionLogCallsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListActionLogHistogram

> ListActionLogHistogramResponse ListActionLogHistogram(ctx).Platform(platform).Action(action).Category(category).Area(area).Status(status).Source(source).TargetType(targetType).TargetId(targetId).ActorUserId(actorUserId).Q(q).Since(since).Until(until).Execute()

Events per UTC day for the list filters (default: the last 30 days; at most 366).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    platform := "platform_example" // string | Platform id (`ebay`, `poshmark`, ...). (optional)
    action := "action_example" // string | A canonical action key (`listing.delist`, see the public action catalog) or a legacy action word. (optional)
    category := "category_example" // string |  (optional)
    area := "area_example" // string | A seller area: every action of that area (the grouped filter of the dashboard). (optional)
    status := "status_example" // string |  (optional)
    source := "source_example" // string |  (optional)
    targetType := "targetType_example" // string | With targetId: events about that object (any of their targets). (optional)
    targetId := "targetId_example" // string |  (optional)
    actorUserId := "actorUserId_example" // string | Only what one team member did. (optional)
    q := "q_example" // string | A search in the activity query language (`action:listing.* -outcome:failure \"free text\"`), ANDed with the other filters. POST /api/v1/action-log/search/translate turns a question in words into one. (optional)
    since := time.Now() // time.Time | Inclusive lower bound on when the action happened. (optional)
    until := time.Now() // time.Time | Exclusive upper bound. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ListActionLogHistogram(context.Background()).Platform(platform).Action(action).Category(category).Area(area).Status(status).Source(source).TargetType(targetType).TargetId(targetId).ActorUserId(actorUserId).Q(q).Since(since).Until(until).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ListActionLogHistogram``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListActionLogHistogram`: ListActionLogHistogramResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ListActionLogHistogram`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListActionLogHistogramRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **platform** | **string** | Platform id (&#x60;ebay&#x60;, &#x60;poshmark&#x60;, ...). | 
 **action** | **string** | A canonical action key (&#x60;listing.delist&#x60;, see the public action catalog) or a legacy action word. | 
 **category** | **string** |  | 
 **area** | **string** | A seller area: every action of that area (the grouped filter of the dashboard). | 
 **status** | **string** |  | 
 **source** | **string** |  | 
 **targetType** | **string** | With targetId: events about that object (any of their targets). | 
 **targetId** | **string** |  | 
 **actorUserId** | **string** | Only what one team member did. | 
 **q** | **string** | A search in the activity query language (&#x60;action:listing.* -outcome:failure \&quot;free text\&quot;&#x60;), ANDed with the other filters. POST /api/v1/action-log/search/translate turns a question in words into one. | 
 **since** | **time.Time** | Inclusive lower bound on when the action happened. | 
 **until** | **time.Time** | Exclusive upper bound. | 

### Return type

**ListActionLogHistogramResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListDetectedEventEvents

> ListDetectedEventEventsResponse ListDetectedEventEvents(ctx).Kinds(kinds).Platforms(platforms).Accounts(accounts).Statuses(statuses).Since(since).Cursor(cursor).Limit(limit).Execute()

Every event Crossly detected (sales, order changes, listing changes, offers, messages, account notices), newest first: all its axes, the rule that decided it, its corroborating reports, the plan per target and how each action went, with canConfirm (an open ask) and canUndo. Cursor-paged; filter by kinds, platforms statuses and accounts (comma-separated) and since (ISO time).

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    kinds := "kinds_example" // string |  (optional)
    platforms := "platforms_example" // string |  (optional)
    accounts := "accounts_example" // string | Comma-separated platform account ids (scope.account). (optional)
    statuses := "statuses_example" // string | Comma-separated: acted, held, asked, confirmed, ignored, undone, failed, received, corrected (optional)
    since := time.Now() // time.Time |  (optional)
    cursor := "cursor_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ListDetectedEventEvents(context.Background()).Kinds(kinds).Platforms(platforms).Accounts(accounts).Statuses(statuses).Since(since).Cursor(cursor).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ListDetectedEventEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDetectedEventEvents`: ListDetectedEventEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ListDetectedEventEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDetectedEventEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **kinds** | **string** |  | 
 **platforms** | **string** |  | 
 **accounts** | **string** | Comma-separated platform account ids (scope.account). | 
 **statuses** | **string** | Comma-separated: acted, held, asked, confirmed, ignored, undone, failed, received, corrected | 
 **since** | **time.Time** |  | 
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 

### Return type

**ListDetectedEventEventsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListDetectedEventRules

> ListDetectedEventRulesResponse ListDetectedEventRules(ctx).Execute()

The seller's detection rules, top-down in the order they are evaluated (first match wins unless a rule says continue): the seller's own rules, translations of their event-policy rules (source legacy), then the shipped default set (source default, ids default:<key>; a stored edit of a default replaces it). Each rule has a filter (event kinds, platforms, accounts, apps, tags, items), an optional condition over any event axis, the actions it runs (then / else) and its guard loosenings.

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
    resp, r, err := apiClient.ActivityApi.ListDetectedEventRules(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ListDetectedEventRules``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDetectedEventRules`: ListDetectedEventRulesResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ListDetectedEventRules`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListDetectedEventRulesRequest struct via the builder pattern


### Return type

**ListDetectedEventRulesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListEmailEvents

> ListEmailEventsResponse ListEmailEvents(ctx).Status(status).Execute()

List email-detected events waiting for the seller (status=needs_review): a possible sale, ended listing, cancellation or other event an email reported that was not certain enough to act on automatically.

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
    status := "status_example" // string |  (optional) (default to "needs_review")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ListEmailEvents(context.Background()).Status(status).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ListEmailEvents``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEmailEvents`: ListEmailEventsResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ListEmailEvents`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEmailEventsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** |  | [default to &quot;needs_review&quot;]

### Return type

**ListEmailEventsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListEventPolicyDecisions

> ListEventPolicyDecisionsResponse ListEventPolicyDecisions(ctx).Cursor(cursor).Limit(limit).Event(event).Ownership(ownership).Platform(platform).Outcome(outcome).Execute()

The decision log, newest first: for each detected event, what the rule or default asked for, what was done (after capability fallback and the email confidence ceiling), the rule that decided it, and the outcome (done / skipped / failed / review). Cursor-paged; filter by event, ownership, platform or outcome (outcome=review: still waiting for the seller). Kept 90 days.

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
    cursor := "cursor_example" // string |  (optional)
    limit := int32(56) // int32 |  (optional)
    event := "event_example" // string |  (optional)
    ownership := "ownership_example" // string |  (optional)
    platform := "platform_example" // string |  (optional)
    outcome := "outcome_example" // string |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.ListEventPolicyDecisions(context.Background()).Cursor(cursor).Limit(limit).Event(event).Ownership(ownership).Platform(platform).Outcome(outcome).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.ListEventPolicyDecisions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListEventPolicyDecisions`: ListEventPolicyDecisionsResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.ListEventPolicyDecisions`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListEventPolicyDecisionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cursor** | **string** |  | 
 **limit** | **int32** |  | 
 **event** | **string** |  | 
 **ownership** | **string** |  | 
 **platform** | **string** |  | 
 **outcome** | **string** |  | 

### Return type

**ListEventPolicyDecisionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateDetectedEventRule

> CreateDetectedEventRuleResponse UpdateDetectedEventRule(ctx, id).CreateDetectedEventRuleRequest(createDetectedEventRuleRequest).Execute()

Replace a detection rule. An id default:<key> stores your edit of that shipped default (delete it, or reset-defaults, to go back to the shipped one).

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
    createDetectedEventRuleRequest := *openapiclient.NewCreateDetectedEventRuleRequest("Name_example", map[string]interface{}{"key": interface{}(123)}, []map[string]interface{}{map[string]interface{}{"key": interface{}(123)}}) // CreateDetectedEventRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.UpdateDetectedEventRule(context.Background(), id).CreateDetectedEventRuleRequest(createDetectedEventRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.UpdateDetectedEventRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDetectedEventRule`: CreateDetectedEventRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.UpdateDetectedEventRule`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDetectedEventRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createDetectedEventRuleRequest** | **CreateDetectedEventRuleRequest** |  | 

### Return type

**CreateDetectedEventRuleResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateEventPolicyRule

> UpdateEventPolicyRuleResponse UpdateEventPolicyRule(ctx).UpdateEventPolicyRuleRequest(updateEventPolicyRuleRequest).Execute()

Set what Crossly does with one kind of detected event for one kind of item: scope global (scopeValue null), platform (scopeValue = platform id) or app (scopeValue = OAuth app id, app_managed items only). Replaces the rule for the same cell and scope. The action must be one the matrix allows for that cell; import actions exist only for untracked items.

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
    updateEventPolicyRuleRequest := *openapiclient.NewUpdateEventPolicyRuleRequest("Scope_example", "Event_example", "Ownership_example", "Action_example") // UpdateEventPolicyRuleRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.ActivityApi.UpdateEventPolicyRule(context.Background()).UpdateEventPolicyRuleRequest(updateEventPolicyRuleRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ActivityApi.UpdateEventPolicyRule``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateEventPolicyRule`: UpdateEventPolicyRuleResponse
    fmt.Fprintf(os.Stdout, "Response from `ActivityApi.UpdateEventPolicyRule`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEventPolicyRuleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateEventPolicyRuleRequest** | **UpdateEventPolicyRuleRequest** |  | 

### Return type

**UpdateEventPolicyRuleResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)

