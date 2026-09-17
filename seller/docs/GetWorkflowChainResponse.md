# GetWorkflowChainResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Chain** | [**GetWorkflowChainResponseChain**](GetWorkflowChainResponseChain.md) |  | 
**Steps** | [**[]ListWorkflowChainsItemSteps**](ListWorkflowChainsItemSteps.md) |  | 

## Methods

### NewGetWorkflowChainResponse

`func NewGetWorkflowChainResponse(chain GetWorkflowChainResponseChain, steps []ListWorkflowChainsItemSteps, ) *GetWorkflowChainResponse`

NewGetWorkflowChainResponse instantiates a new GetWorkflowChainResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetWorkflowChainResponseWithDefaults

`func NewGetWorkflowChainResponseWithDefaults() *GetWorkflowChainResponse`

NewGetWorkflowChainResponseWithDefaults instantiates a new GetWorkflowChainResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChain

`func (o *GetWorkflowChainResponse) GetChain() GetWorkflowChainResponseChain`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *GetWorkflowChainResponse) GetChainOk() (*GetWorkflowChainResponseChain, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *GetWorkflowChainResponse) SetChain(v GetWorkflowChainResponseChain)`

SetChain sets Chain field to given value.


### GetSteps

`func (o *GetWorkflowChainResponse) GetSteps() []ListWorkflowChainsItemSteps`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *GetWorkflowChainResponse) GetStepsOk() (*[]ListWorkflowChainsItemSteps, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *GetWorkflowChainResponse) SetSteps(v []ListWorkflowChainsItemSteps)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


