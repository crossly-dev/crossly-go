# \WorkflowsApi

All URIs are relative to *https://crossly.net/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkflow**](WorkflowsApi.md#CreateWorkflow) | **Post** /v1/workflows | Create a workflow: from blank ({ name, graph? }), from a template ({ from: \&quot;template\&quot;, templateId, setup?, activate? }) or from an export document ({ from: \&quot;import\&quot;, document }). New workflows start paused; activate one with POST /v1/workflows/{id}/activate. Returns the workflow and its validation issues.
[**CreateWorkflowActivate**](WorkflowsApi.md#CreateWorkflowActivate) | **Post** /v1/workflows/{id}/activate | Activate a workflow: its triggers start firing. A graph with blocking issues is refused (422, details.issues).
[**CreateWorkflowApprovalDecide**](WorkflowsApi.md#CreateWorkflowApprovalDecide) | **Post** /v1/workflow-approvals/{id}/decide | Approve or reject a pending flow.approval request: the waiting run resumes down the approved or rejected port. 409 when it was already decided (another device, Telegram, or the timeout).
[**CreateWorkflowChain**](WorkflowsApi.md#CreateWorkflowChain) | **Post** /v1/workflow-chains | Create a multi-step workflow chain.
[**CreateWorkflowChainRunNow**](WorkflowsApi.md#CreateWorkflowChainRunNow) | **Post** /v1/workflow-chains/{id}/run-now | Enqueue an ad-hoc run of a workflow chain.
[**CreateWorkflowChainToggle**](WorkflowsApi.md#CreateWorkflowChainToggle) | **Post** /v1/workflow-chains/{id}/toggle | Flip a workflow chain between active and inactive.
[**CreateWorkflowDeactivate**](WorkflowsApi.md#CreateWorkflowDeactivate) | **Post** /v1/workflows/{id}/deactivate | Pause a workflow: its triggers stop firing. Runs already started finish.
[**CreateWorkflowDuplicate**](WorkflowsApi.md#CreateWorkflowDuplicate) | **Post** /v1/workflows/{id}/duplicate | Copy a workflow (graph, description, tags) into a new paused workflow.
[**CreateWorkflowExpressionPreview**](WorkflowsApi.md#CreateWorkflowExpressionPreview) | **Post** /v1/workflows/expressions/preview | Resolve one param value (\&quot;{{ $json.priceCents * 0.9 }}\&quot;) with the engine&#39;s evaluator against an item, the trigger item, earlier nodes&#39; items or a stored run. A bad expression is an answer (ok&#x3D;false with the message), not an HTTP error. Changes nothing.
[**CreateWorkflowImport**](WorkflowsApi.md#CreateWorkflowImport) | **Post** /v1/workflows/import | Import a crossly.workflow export document (or { document, name }) as a new, paused workflow. Integration and workflow references are cleared (they belong to the exporting account) and listed in stripped.
[**CreateWorkflowRun**](WorkflowsApi.md#CreateWorkflowRun) | **Post** /v1/workflows/{id}/run | Run a workflow (active or paused) live now, on items you pass or on selected listings / orders / inventory items (one kind per call, up to 1,000). The graph must validate (422 with issues otherwise).
[**CreateWorkflowRunCancel**](WorkflowsApi.md#CreateWorkflowRunCancel) | **Post** /v1/workflow-runs/{runId}/cancel | Cancel a queued, running or waiting run. 409 when it already finished.
[**CreateWorkflowRunOn**](WorkflowsApi.md#CreateWorkflowRunOn) | **Post** /v1/workflows/run-on | \&quot;Run workflow on…\&quot;: the same as POST /v1/workflows/{id}/run with the workflow named in the body.
[**CreateWorkflowRunRetry**](WorkflowsApi.md#CreateWorkflowRunRetry) | **Post** /v1/workflow-runs/{runId}/retry | Retry a failed run from the node that failed: a new run (retryOfRunId) seeded from the failed run&#39;s checkpoint, so nodes that succeeded are not repeated. 409 unless the run failed.
[**CreateWorkflowSaleEnd**](WorkflowsApi.md#CreateWorkflowSaleEnd) | **Post** /v1/workflows/sales/{saleEventId}/end | End a sale now. A sale that has not started is cancelled; a live one closes and every discounted price is put back now (reverted &#x3D; how many listings were restored). 404 for a sale that is not yours.
[**CreateWorkflowTemplateUse**](WorkflowsApi.md#CreateWorkflowTemplateUse) | **Post** /v1/workflows/templates/{id}/use | Create a workflow from a template, filling its setup fields; activate&#x3D;true activates it when it validates.
[**CreateWorkflowTest**](WorkflowsApi.md#CreateWorkflowTest) | **Post** /v1/workflows/{id}/test | Test a workflow (the saved graph, or a draft graph you pass). mode&#x3D;test runs it for real, using pinned data where a node has it; mode&#x3D;dry is a safe test: write, destructive and money nodes report what they would do and change nothing. Trigger items: the ones you pass, else pinned, else the last real ones, else the trigger sample.
[**CreateWorkflowValidate**](WorkflowsApi.md#CreateWorkflowValidate) | **Post** /v1/workflows/validate | Validate a graph without saving it: unknown node types, missing required params, bad ports, cycles, expression errors. blocking&#x3D;true means activation would be refused.
[**CreateWorkflowVersionRestore**](WorkflowsApi.md#CreateWorkflowVersionRestore) | **Post** /v1/workflows/{id}/versions/{v}/restore | Restore an earlier version&#39;s graph as a new version. Pass version (the current one you read) to refuse the restore when someone saved meanwhile (409). Body optional: { version?: integer }.
[**CreateWorkflowWebhookRotate**](WorkflowsApi.md#CreateWorkflowWebhookRotate) | **Post** /v1/workflows/{id}/webhooks/{nodeId}/rotate | Issue a new inbound URL for an on.webhook trigger node (the old one stops working). The URL, and the signing secret when withSecret&#x3D;true, are returned ONCE: store them now.
[**DeleteWorkflow**](WorkflowsApi.md#DeleteWorkflow) | **Delete** /v1/workflows/{id} | Delete (archive) a workflow: its triggers stop, open runs are cancelled, run history stays readable.
[**DeleteWorkflowChain**](WorkflowsApi.md#DeleteWorkflowChain) | **Delete** /v1/workflow-chains/{id} | Delete a workflow chain (cascades steps + runs).
[**GetWorkflow**](WorkflowsApi.md#GetWorkflow) | **Get** /v1/workflows/{id} | Get one workflow with its full graph and current version.
[**GetWorkflowCatalog**](WorkflowsApi.md#GetWorkflowCatalog) | **Get** /v1/workflows/catalog | Everything needed to build a workflow: every node type (params, ports, output fields, risk, platforms capability), which platforms each capability covers and which of those you have connected, the expression helpers and bindings ({{ $json.listingId }}…) and the condition operators. Read this before composing a graph rather than guessing node types or param keys.
[**GetWorkflowChain**](WorkflowsApi.md#GetWorkflowChain) | **Get** /v1/workflow-chains/{id} | Get one workflow chain with its steps.
[**GetWorkflowExport**](WorkflowsApi.md#GetWorkflowExport) | **Get** /v1/workflows/{id}/export | Export a workflow as a portable document (format crossly.workflow) that POST /v1/workflows/import accepts. Credentials and account handles are stripped; pinned test data only with pinnedData&#x3D;true.
[**GetWorkflowRun**](WorkflowsApi.md#GetWorkflowRun) | **Get** /v1/workflow-runs/{runId} | One run: its summary, the graph version it ran, and every node execution (status, item counts per port, timing, errors). Item data per node: GET …/nodes/{nodeId}.
[**GetWorkflowRunNode**](WorkflowsApi.md#GetWorkflowRunNode) | **Get** /v1/workflow-runs/{runId}/nodes/{nodeId} | Every execution of one node in a run with its input and output items (each port keeps the first 50 items, 16 KB per item; truncated&#x3D;true when cut). Data older than the retention window, or not kept by settings.saveSuccessData&#x3D;false, is pruned: that execution has input/output null and pruned&#x3D;true, and when every execution is pruned the response is { data: null, meta: { pruned: true } } (never an empty list).
[**ListWorkflowApprovals**](WorkflowsApi.md#ListWorkflowApprovals) | **Get** /v1/workflow-approvals | Approval requests from flow.approval nodes, newest first (pending only unless status says otherwise). Decide one with POST /v1/workflow-approvals/{id}/decide.
[**ListWorkflowChains**](WorkflowsApi.md#ListWorkflowChains) | **Get** /v1/workflow-chains | List workflow chains with their step graph.
[**ListWorkflowRuns**](WorkflowsApi.md#ListWorkflowRuns) | **Get** /v1/workflows/{id}/runs | Run history of one workflow, newest first, with per-node status and item counts.
[**ListWorkflowRuns2**](WorkflowsApi.md#ListWorkflowRuns2) | **Get** /v1/workflow-runs | Runs of every workflow, newest first. Filter by workflow, status, mode and time.
[**ListWorkflowSaleActive**](WorkflowsApi.md#ListWorkflowSaleActive) | **Get** /v1/workflows/sales/active | Every live or upcoming sale (a sales event started by a sale.start node) with the workflow that runs it: percent off, window, and how many listings it applies to.
[**ListWorkflowTemplates**](WorkflowsApi.md#ListWorkflowTemplates) | **Get** /v1/workflows/templates | Ready-made workflows (weekend sale, stale-listing markdowns, offer responder, ship-by alerts…), each a complete graph plus the setup fields its \&quot;use\&quot; call asks for.
[**ListWorkflowVersions**](WorkflowsApi.md#ListWorkflowVersions) | **Get** /v1/workflows/{id}/versions | The saved versions of a workflow, newest first (restore one with POST …/versions/{v}/restore).
[**ListWorkflows**](WorkflowsApi.md#ListWorkflows) | **Get** /v1/workflows | List workflows, newest change first. A workflow is a graph of nodes: one or more triggers, then actions, lookups and flow nodes (if, switch, wait, approval…). Each row carries its trigger and node types and a 7-day run sparkline, not the graph (GET /v1/workflows/{id} has it).
[**UpdateWorkflow**](WorkflowsApi.md#UpdateWorkflow) | **Put** /v1/workflows/{id} | Save a workflow. Pass the version you edited: a save over a newer version is 409 version_conflict (re-read and merge). Saving an active workflow validates it like activation does (422 with issues).
[**UpdateWorkflowChain**](WorkflowsApi.md#UpdateWorkflowChain) | **Put** /v1/workflow-chains/{id} | Replace a workflow chain wholesale.



## CreateWorkflow

> CreateWorkflowResponse CreateWorkflow(ctx).CreateWorkflowRequest(createWorkflowRequest).Execute()

Create a workflow: from blank ({ name, graph? }), from a template ({ from: \"template\", templateId, setup?, activate? }) or from an export document ({ from: \"import\", document }). New workflows start paused; activate one with POST /v1/workflows/{id}/activate. Returns the workflow and its validation issues.

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
    createWorkflowRequest := *openapiclient.NewCreateWorkflowRequest() // CreateWorkflowRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflow(context.Background()).CreateWorkflowRequest(createWorkflowRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflow`: CreateWorkflowResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflow`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkflowRequest** | **CreateWorkflowRequest** |  | 

### Return type

**CreateWorkflowResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowActivate

> WorkflowResponse CreateWorkflowActivate(ctx, id).Execute()

Activate a workflow: its triggers start firing. A graph with blocking issues is refused (422, details.issues).

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowActivate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowActivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowActivate`: WorkflowResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowActivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowActivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorkflowResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowApprovalDecide

> CreateWorkflowApprovalDecideResponse CreateWorkflowApprovalDecide(ctx, id).CreateWorkflowApprovalDecideRequest(createWorkflowApprovalDecideRequest).Execute()

Approve or reject a pending flow.approval request: the waiting run resumes down the approved or rejected port. 409 when it was already decided (another device, Telegram, or the timeout).

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Approval id from GET /v1/workflow-approvals.
    createWorkflowApprovalDecideRequest := *openapiclient.NewCreateWorkflowApprovalDecideRequest("Decision_example") // CreateWorkflowApprovalDecideRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowApprovalDecide(context.Background(), id).CreateWorkflowApprovalDecideRequest(createWorkflowApprovalDecideRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowApprovalDecide``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowApprovalDecide`: CreateWorkflowApprovalDecideResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowApprovalDecide`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Approval id from GET /v1/workflow-approvals. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowApprovalDecideRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createWorkflowApprovalDecideRequest** | **CreateWorkflowApprovalDecideRequest** |  | 

### Return type

**CreateWorkflowApprovalDecideResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowChain

> IdResponse CreateWorkflowChain(ctx).Execute()

Create a multi-step workflow chain.



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
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowChain(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowChain`: IdResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowChain`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowChainRequest struct via the builder pattern


### Return type

**IdResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowChainRunNow

> CreateWorkflowChainRunNowResponse CreateWorkflowChainRunNow(ctx, id).Execute()

Enqueue an ad-hoc run of a workflow chain.



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
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowChainRunNow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowChainRunNow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowChainRunNow`: CreateWorkflowChainRunNowResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowChainRunNow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowChainRunNowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorkflowChainRunNowResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowChainToggle

> CreateWorkflowChainToggleResponse CreateWorkflowChainToggle(ctx, id).Execute()

Flip a workflow chain between active and inactive.



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
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowChainToggle(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowChainToggle``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowChainToggle`: CreateWorkflowChainToggleResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowChainToggle`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowChainToggleRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorkflowChainToggleResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowDeactivate

> WorkflowResponse CreateWorkflowDeactivate(ctx, id).Execute()

Pause a workflow: its triggers stop firing. Runs already started finish.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowDeactivate(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowDeactivate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowDeactivate`: WorkflowResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowDeactivate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowDeactivateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**WorkflowResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowDuplicate

> CreateWorkflowDuplicateResponse CreateWorkflowDuplicate(ctx, id).CreateWorkflowDuplicateRequest(createWorkflowDuplicateRequest).Execute()

Copy a workflow (graph, description, tags) into a new paused workflow.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow id.
    createWorkflowDuplicateRequest := *openapiclient.NewCreateWorkflowDuplicateRequest() // CreateWorkflowDuplicateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowDuplicate(context.Background(), id).CreateWorkflowDuplicateRequest(createWorkflowDuplicateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowDuplicate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowDuplicate`: CreateWorkflowDuplicateResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowDuplicate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowDuplicateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createWorkflowDuplicateRequest** | **CreateWorkflowDuplicateRequest** |  | 

### Return type

**CreateWorkflowDuplicateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowExpressionPreview

> CreateWorkflowExpressionPreviewResponse CreateWorkflowExpressionPreview(ctx).CreateWorkflowExpressionPreviewRequest(createWorkflowExpressionPreviewRequest).Execute()

Resolve one param value (\"{{ $json.priceCents * 0.9 }}\") with the engine's evaluator against an item, the trigger item, earlier nodes' items or a stored run. A bad expression is an answer (ok=false with the message), not an HTTP error. Changes nothing.

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
    createWorkflowExpressionPreviewRequest := *openapiclient.NewCreateWorkflowExpressionPreviewRequest() // CreateWorkflowExpressionPreviewRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowExpressionPreview(context.Background()).CreateWorkflowExpressionPreviewRequest(createWorkflowExpressionPreviewRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowExpressionPreview``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowExpressionPreview`: CreateWorkflowExpressionPreviewResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowExpressionPreview`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowExpressionPreviewRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkflowExpressionPreviewRequest** | **CreateWorkflowExpressionPreviewRequest** |  | 

### Return type

**CreateWorkflowExpressionPreviewResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowImport

> CreateWorkflowImportResponse CreateWorkflowImport(ctx).RequestBody(requestBody).Execute()

Import a crossly.workflow export document (or { document, name }) as a new, paused workflow. Integration and workflow references are cleared (they belong to the exporting account) and listed in stripped.

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
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowImport(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowImport`: CreateWorkflowImportResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

**CreateWorkflowImportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowRun

> CreateWorkflowRunResponse CreateWorkflowRun(ctx, id).CreateWorkflowRunRequest(createWorkflowRunRequest).Execute()

Run a workflow (active or paused) live now, on items you pass or on selected listings / orders / inventory items (one kind per call, up to 1,000). The graph must validate (422 with issues otherwise).

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow id.
    createWorkflowRunRequest := *openapiclient.NewCreateWorkflowRunRequest() // CreateWorkflowRunRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowRun(context.Background(), id).CreateWorkflowRunRequest(createWorkflowRunRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowRun`: CreateWorkflowRunResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createWorkflowRunRequest** | **CreateWorkflowRunRequest** |  | 

### Return type

**CreateWorkflowRunResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowRunCancel

> CreateWorkflowRunCancelResponse CreateWorkflowRunCancel(ctx, runId).Execute()

Cancel a queued, running or waiting run. 409 when it already finished.

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
    runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow run id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowRunCancel(context.Background(), runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowRunCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowRunCancel`: CreateWorkflowRunCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowRunCancel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Workflow run id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowRunCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorkflowRunCancelResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowRunOn

> CreateWorkflowRunOnResponse CreateWorkflowRunOn(ctx).CreateWorkflowRunOnRequest(createWorkflowRunOnRequest).Execute()

\"Run workflow on…\": the same as POST /v1/workflows/{id}/run with the workflow named in the body.

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
    createWorkflowRunOnRequest := *openapiclient.NewCreateWorkflowRunOnRequest("WorkflowId_example") // CreateWorkflowRunOnRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowRunOn(context.Background()).CreateWorkflowRunOnRequest(createWorkflowRunOnRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowRunOn``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowRunOn`: CreateWorkflowRunOnResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowRunOn`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowRunOnRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkflowRunOnRequest** | **CreateWorkflowRunOnRequest** |  | 

### Return type

**CreateWorkflowRunOnResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowRunRetry

> CreateWorkflowRunRetryResponse CreateWorkflowRunRetry(ctx, runId).Execute()

Retry a failed run from the node that failed: a new run (retryOfRunId) seeded from the failed run's checkpoint, so nodes that succeeded are not repeated. 409 unless the run failed.

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
    runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow run id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowRunRetry(context.Background(), runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowRunRetry``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowRunRetry`: CreateWorkflowRunRetryResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowRunRetry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Workflow run id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowRunRetryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorkflowRunRetryResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowSaleEnd

> CreateWorkflowSaleEndResponse CreateWorkflowSaleEnd(ctx, saleEventId).Execute()

End a sale now. A sale that has not started is cancelled; a live one closes and every discounted price is put back now (reverted = how many listings were restored). 404 for a sale that is not yours.

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
    saleEventId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Sale id from GET /v1/workflows/sales/active.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowSaleEnd(context.Background(), saleEventId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowSaleEnd``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowSaleEnd`: CreateWorkflowSaleEndResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowSaleEnd`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**saleEventId** | **string** | Sale id from GET /v1/workflows/sales/active. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowSaleEndRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**CreateWorkflowSaleEndResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowTemplateUse

> CreateWorkflowResponse CreateWorkflowTemplateUse(ctx, id).CreateWorkflowTemplateUseRequest(createWorkflowTemplateUseRequest).Execute()

Create a workflow from a template, filling its setup fields; activate=true activates it when it validates.

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
    id := "id_example" // string | Template id from GET /v1/workflows/templates.
    createWorkflowTemplateUseRequest := *openapiclient.NewCreateWorkflowTemplateUseRequest() // CreateWorkflowTemplateUseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowTemplateUse(context.Background(), id).CreateWorkflowTemplateUseRequest(createWorkflowTemplateUseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowTemplateUse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowTemplateUse`: CreateWorkflowResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowTemplateUse`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Template id from GET /v1/workflows/templates. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowTemplateUseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createWorkflowTemplateUseRequest** | **CreateWorkflowTemplateUseRequest** |  | 

### Return type

**CreateWorkflowResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowTest

> CreateWorkflowTestResponse CreateWorkflowTest(ctx, id).CreateWorkflowTestRequest(createWorkflowTestRequest).Execute()

Test a workflow (the saved graph, or a draft graph you pass). mode=test runs it for real, using pinned data where a node has it; mode=dry is a safe test: write, destructive and money nodes report what they would do and change nothing. Trigger items: the ones you pass, else pinned, else the last real ones, else the trigger sample.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow id.
    createWorkflowTestRequest := *openapiclient.NewCreateWorkflowTestRequest() // CreateWorkflowTestRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowTest(context.Background(), id).CreateWorkflowTestRequest(createWorkflowTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowTest``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowTest`: CreateWorkflowTestResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowTest`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowTestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createWorkflowTestRequest** | **CreateWorkflowTestRequest** |  | 

### Return type

**CreateWorkflowTestResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowValidate

> CreateWorkflowValidateResponse CreateWorkflowValidate(ctx).CreateWorkflowValidateRequest(createWorkflowValidateRequest).Execute()

Validate a graph without saving it: unknown node types, missing required params, bad ports, cycles, expression errors. blocking=true means activation would be refused.

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
    createWorkflowValidateRequest := *openapiclient.NewCreateWorkflowValidateRequest(map[string]interface{}{"key": interface{}(123)}) // CreateWorkflowValidateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowValidate(context.Background()).CreateWorkflowValidateRequest(createWorkflowValidateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowValidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowValidate`: CreateWorkflowValidateResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowValidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowValidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createWorkflowValidateRequest** | **CreateWorkflowValidateRequest** |  | 

### Return type

**CreateWorkflowValidateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowVersionRestore

> WorkflowResponse CreateWorkflowVersionRestore(ctx, id, v).Execute()

Restore an earlier version's graph as a new version. Pass version (the current one you read) to refuse the restore when someone saved meanwhile (409). Body optional: { version?: integer }.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow id.
    v := int32(56) // int32 | The version number to restore.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowVersionRestore(context.Background(), id, v).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowVersionRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowVersionRestore`: WorkflowResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowVersionRestore`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id. | 
**v** | **int32** | The version number to restore. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowVersionRestoreRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**WorkflowResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## CreateWorkflowWebhookRotate

> CreateWorkflowWebhookRotateResponse CreateWorkflowWebhookRotate(ctx, id, nodeId).CreateWorkflowWebhookRotateRequest(createWorkflowWebhookRotateRequest).Execute()

Issue a new inbound URL for an on.webhook trigger node (the old one stops working). The URL, and the signing secret when withSecret=true, are returned ONCE: store them now.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow id.
    nodeId := "nodeId_example" // string | The on.webhook trigger node id.
    createWorkflowWebhookRotateRequest := *openapiclient.NewCreateWorkflowWebhookRotateRequest() // CreateWorkflowWebhookRotateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.CreateWorkflowWebhookRotate(context.Background(), id, nodeId).CreateWorkflowWebhookRotateRequest(createWorkflowWebhookRotateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.CreateWorkflowWebhookRotate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkflowWebhookRotate`: CreateWorkflowWebhookRotateResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.CreateWorkflowWebhookRotate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id. | 
**nodeId** | **string** | The on.webhook trigger node id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkflowWebhookRotateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **createWorkflowWebhookRotateRequest** | **CreateWorkflowWebhookRotateRequest** |  | 

### Return type

**CreateWorkflowWebhookRotateResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorkflow

> DeleteWorkflowResponse DeleteWorkflow(ctx, id).Execute()

Delete (archive) a workflow: its triggers stop, open runs are cancelled, run history stays readable.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.DeleteWorkflow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.DeleteWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorkflow`: DeleteWorkflowResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.DeleteWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**DeleteWorkflowResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## DeleteWorkflowChain

> OkResponse DeleteWorkflowChain(ctx, id).Execute()

Delete a workflow chain (cascades steps + runs).



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
    resp, r, err := apiClient.WorkflowsApi.DeleteWorkflowChain(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.DeleteWorkflowChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorkflowChain`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.DeleteWorkflowChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkflowChainRequest struct via the builder pattern


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


## GetWorkflow

> GetWorkflowResponse GetWorkflow(ctx, id).Execute()

Get one workflow with its full graph and current version.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.GetWorkflow(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.GetWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflow`: GetWorkflowResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.GetWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorkflowResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorkflowCatalog

> GetWorkflowCatalogResponse GetWorkflowCatalog(ctx).Execute()

Everything needed to build a workflow: every node type (params, ports, output fields, risk, platforms capability), which platforms each capability covers and which of those you have connected, the expression helpers and bindings ({{ $json.listingId }}…) and the condition operators. Read this before composing a graph rather than guessing node types or param keys.

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
    resp, r, err := apiClient.WorkflowsApi.GetWorkflowCatalog(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.GetWorkflowCatalog``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowCatalog`: GetWorkflowCatalogResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.GetWorkflowCatalog`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowCatalogRequest struct via the builder pattern


### Return type

**GetWorkflowCatalogResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorkflowChain

> GetWorkflowChainResponse GetWorkflowChain(ctx, id).Execute()

Get one workflow chain with its steps.



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
    resp, r, err := apiClient.WorkflowsApi.GetWorkflowChain(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.GetWorkflowChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowChain`: GetWorkflowChainResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.GetWorkflowChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowChainRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorkflowChainResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorkflowExport

> GetWorkflowExportResponse GetWorkflowExport(ctx, id).PinnedData(pinnedData).Execute()

Export a workflow as a portable document (format crossly.workflow) that POST /v1/workflows/import accepts. Credentials and account handles are stripped; pinned test data only with pinnedData=true.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow id.
    pinnedData := "pinnedData_example" // string | Keep the nodes' pinned test data. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.GetWorkflowExport(context.Background(), id).PinnedData(pinnedData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.GetWorkflowExport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowExport`: GetWorkflowExportResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.GetWorkflowExport`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowExportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pinnedData** | **string** | Keep the nodes&#39; pinned test data. | 

### Return type

**GetWorkflowExportResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorkflowRun

> GetWorkflowRunResponse GetWorkflowRun(ctx, runId).Execute()

One run: its summary, the graph version it ran, and every node execution (status, item counts per port, timing, errors). Item data per node: GET …/nodes/{nodeId}.

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
    runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow run id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.GetWorkflowRun(context.Background(), runId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.GetWorkflowRun``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowRun`: GetWorkflowRunResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.GetWorkflowRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Workflow run id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**GetWorkflowRunResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## GetWorkflowRunNode

> GetWorkflowRunNodeResponse GetWorkflowRunNode(ctx, runId, nodeId).Execute()

Every execution of one node in a run with its input and output items (each port keeps the first 50 items, 16 KB per item; truncated=true when cut). Data older than the retention window, or not kept by settings.saveSuccessData=false, is pruned: that execution has input/output null and pruned=true, and when every execution is pruned the response is { data: null, meta: { pruned: true } } (never an empty list).

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
    runId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow run id.
    nodeId := "nodeId_example" // string | Node id within the run's graph.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.GetWorkflowRunNode(context.Background(), runId, nodeId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.GetWorkflowRunNode``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkflowRunNode`: GetWorkflowRunNodeResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.GetWorkflowRunNode`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**runId** | **string** | Workflow run id. | 
**nodeId** | **string** | Node id within the run&#39;s graph. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkflowRunNodeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

**GetWorkflowRunNodeResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorkflowApprovals

> ListWorkflowApprovalsResponse ListWorkflowApprovals(ctx).Status(status).WorkflowId(workflowId).Page(page).Limit(limit).Execute()

Approval requests from flow.approval nodes, newest first (pending only unless status says otherwise). Decide one with POST /v1/workflow-approvals/{id}/decide.

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
    status := "status_example" // string | Comma-separated: pending, approved, rejected, expired. Default pending. (optional)
    workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only approvals of this workflow. (optional)
    page := int32(56) // int32 | 1-indexed page. (optional) (default to 1)
    limit := int32(56) // int32 | Rows per page (1-100). (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.ListWorkflowApprovals(context.Background()).Status(status).WorkflowId(workflowId).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.ListWorkflowApprovals``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowApprovals`: ListWorkflowApprovalsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.ListWorkflowApprovals`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowApprovalsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **status** | **string** | Comma-separated: pending, approved, rejected, expired. Default pending. | 
 **workflowId** | **string** | Only approvals of this workflow. | 
 **page** | **int32** | 1-indexed page. | [default to 1]
 **limit** | **int32** | Rows per page (1-100). | [default to 25]

### Return type

**ListWorkflowApprovalsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorkflowChains

> ListWorkflowChainsResponse ListWorkflowChains(ctx).Execute()

List workflow chains with their step graph.



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
    resp, r, err := apiClient.WorkflowsApi.ListWorkflowChains(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.ListWorkflowChains``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowChains`: ListWorkflowChainsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.ListWorkflowChains`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowChainsRequest struct via the builder pattern


### Return type

**ListWorkflowChainsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorkflowRuns

> ListWorkflowRunsResponse ListWorkflowRuns(ctx, id).Status(status).Mode(mode).Since(since).Until(until).Page(page).Limit(limit).Execute()

Run history of one workflow, newest first, with per-node status and item counts.

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow id.
    status := "status_example" // string | Comma-separated run statuses: queued, running, waiting, succeeded, failed, cancelled. (optional)
    mode := "mode_example" // string | live (triggers), test (editor test run) or dry (safe test: no writes). (optional)
    since := time.Now() // time.Time | Runs queued at or after this time (ISO 8601). (optional)
    until := time.Now() // time.Time | Runs queued at or before this time (ISO 8601). (optional)
    page := int32(56) // int32 | 1-indexed page. (optional) (default to 1)
    limit := int32(56) // int32 | Rows per page (1-100). (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.ListWorkflowRuns(context.Background(), id).Status(status).Mode(mode).Since(since).Until(until).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.ListWorkflowRuns``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowRuns`: ListWorkflowRunsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.ListWorkflowRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **status** | **string** | Comma-separated run statuses: queued, running, waiting, succeeded, failed, cancelled. | 
 **mode** | **string** | live (triggers), test (editor test run) or dry (safe test: no writes). | 
 **since** | **time.Time** | Runs queued at or after this time (ISO 8601). | 
 **until** | **time.Time** | Runs queued at or before this time (ISO 8601). | 
 **page** | **int32** | 1-indexed page. | [default to 1]
 **limit** | **int32** | Rows per page (1-100). | [default to 25]

### Return type

**ListWorkflowRunsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorkflowRuns2

> ListWorkflowRunsResponse ListWorkflowRuns2(ctx).WorkflowId(workflowId).Status(status).Mode(mode).Since(since).Until(until).Page(page).Limit(limit).Execute()

Runs of every workflow, newest first. Filter by workflow, status, mode and time.

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
    workflowId := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Only runs of this workflow. (optional)
    status := "status_example" // string | Comma-separated run statuses: queued, running, waiting, succeeded, failed, cancelled. (optional)
    mode := "mode_example" // string | live (triggers), test (editor test run) or dry (safe test: no writes). (optional)
    since := time.Now() // time.Time | Runs queued at or after this time (ISO 8601). (optional)
    until := time.Now() // time.Time | Runs queued at or before this time (ISO 8601). (optional)
    page := int32(56) // int32 | 1-indexed page. (optional) (default to 1)
    limit := int32(56) // int32 | Rows per page (1-100). (optional) (default to 25)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.ListWorkflowRuns2(context.Background()).WorkflowId(workflowId).Status(status).Mode(mode).Since(since).Until(until).Page(page).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.ListWorkflowRuns2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowRuns2`: ListWorkflowRunsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.ListWorkflowRuns2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowRuns2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **workflowId** | **string** | Only runs of this workflow. | 
 **status** | **string** | Comma-separated run statuses: queued, running, waiting, succeeded, failed, cancelled. | 
 **mode** | **string** | live (triggers), test (editor test run) or dry (safe test: no writes). | 
 **since** | **time.Time** | Runs queued at or after this time (ISO 8601). | 
 **until** | **time.Time** | Runs queued at or before this time (ISO 8601). | 
 **page** | **int32** | 1-indexed page. | [default to 1]
 **limit** | **int32** | Rows per page (1-100). | [default to 25]

### Return type

**ListWorkflowRunsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorkflowSaleActive

> ListWorkflowSaleActiveResponse ListWorkflowSaleActive(ctx).Execute()

Every live or upcoming sale (a sales event started by a sale.start node) with the workflow that runs it: percent off, window, and how many listings it applies to.

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
    resp, r, err := apiClient.WorkflowsApi.ListWorkflowSaleActive(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.ListWorkflowSaleActive``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowSaleActive`: ListWorkflowSaleActiveResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.ListWorkflowSaleActive`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowSaleActiveRequest struct via the builder pattern


### Return type

**ListWorkflowSaleActiveResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorkflowTemplates

> ListWorkflowTemplatesResponse ListWorkflowTemplates(ctx).Category(category).Q(q).Execute()

Ready-made workflows (weekend sale, stale-listing markdowns, offer responder, ship-by alerts…), each a complete graph plus the setup fields its \"use\" call asks for.

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
    category := "category_example" // string | Exact template category, e.g. \"Sales & markdowns\". (optional)
    q := "q_example" // string | Matches name, description, category or tags. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.ListWorkflowTemplates(context.Background()).Category(category).Q(q).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.ListWorkflowTemplates``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowTemplates`: ListWorkflowTemplatesResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.ListWorkflowTemplates`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowTemplatesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **category** | **string** | Exact template category, e.g. \&quot;Sales &amp; markdowns\&quot;. | 
 **q** | **string** | Matches name, description, category or tags. | 

### Return type

**ListWorkflowTemplatesResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorkflowVersions

> ListWorkflowVersionsResponse ListWorkflowVersions(ctx, id).Execute()

The saved versions of a workflow, newest first (restore one with POST …/versions/{v}/restore).

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow id.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.ListWorkflowVersions(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.ListWorkflowVersions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflowVersions`: ListWorkflowVersionsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.ListWorkflowVersions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowVersionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**ListWorkflowVersionsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## ListWorkflows

> ListWorkflowsResponse ListWorkflows(ctx).Active(active).Tag(tag).Trigger(trigger).Q(q).LegacyRef(legacyRef).Execute()

List workflows, newest change first. A workflow is a graph of nodes: one or more triggers, then actions, lookups and flow nodes (if, switch, wait, approval…). Each row carries its trigger and node types and a 7-day run sparkline, not the graph (GET /v1/workflows/{id} has it).

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
    active := "active_example" // string | Only active (true) or paused (false) workflows. (optional)
    tag := "tag_example" // string | Only workflows carrying this tag. (optional)
    trigger := "trigger_example" // string | Only workflows started by this trigger node type, e.g. on.listing.sold. (optional)
    q := "q_example" // string | Name contains (case-insensitive). (optional)
    legacyRef := "legacyRef_example" // string | The workflow an old automation became: rule:<id>, chain:<id> or sale_event:<id>. (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.ListWorkflows(context.Background()).Active(active).Tag(tag).Trigger(trigger).Q(q).LegacyRef(legacyRef).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.ListWorkflows``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkflows`: ListWorkflowsResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.ListWorkflows`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkflowsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **string** | Only active (true) or paused (false) workflows. | 
 **tag** | **string** | Only workflows carrying this tag. | 
 **trigger** | **string** | Only workflows started by this trigger node type, e.g. on.listing.sold. | 
 **q** | **string** | Name contains (case-insensitive). | 
 **legacyRef** | **string** | The workflow an old automation became: rule:&lt;id&gt;, chain:&lt;id&gt; or sale_event:&lt;id&gt;. | 

### Return type

**ListWorkflowsResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorkflow

> WorkflowResponse UpdateWorkflow(ctx, id).UpdateWorkflowRequest(updateWorkflowRequest).Execute()

Save a workflow. Pass the version you edited: a save over a newer version is 409 version_conflict (re-read and merge). Saving an active workflow validates it like activation does (422 with issues).

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
    id := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | Workflow id.
    updateWorkflowRequest := *openapiclient.NewUpdateWorkflowRequest(int32(123)) // UpdateWorkflowRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkflowsApi.UpdateWorkflow(context.Background(), id).UpdateWorkflowRequest(updateWorkflowRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.UpdateWorkflow``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflow`: WorkflowResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.UpdateWorkflow`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Workflow id. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateWorkflowRequest** | **UpdateWorkflowRequest** |  | 

### Return type

**WorkflowResponse**

### Authorization

[PersonalAccessToken](../README.md#PersonalAccessToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to README]](../README.md)


## UpdateWorkflowChain

> OkResponse UpdateWorkflowChain(ctx, id).Execute()

Replace a workflow chain wholesale.



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
    resp, r, err := apiClient.WorkflowsApi.UpdateWorkflowChain(context.Background(), id).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkflowsApi.UpdateWorkflowChain``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateWorkflowChain`: OkResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkflowsApi.UpdateWorkflowChain`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWorkflowChainRequest struct via the builder pattern


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

