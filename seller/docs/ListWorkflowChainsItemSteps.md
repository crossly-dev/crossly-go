# ListWorkflowChainsItemSteps

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Position** | **float32** |  | 
**ChainId** | **string** |  | 
**StepType** | **string** |  | 
**ActionId** | Pointer to **NullableString** |  | [optional] 
**WaitMs** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewListWorkflowChainsItemSteps

`func NewListWorkflowChainsItemSteps(id string, position float32, chainId string, stepType string, ) *ListWorkflowChainsItemSteps`

NewListWorkflowChainsItemSteps instantiates a new ListWorkflowChainsItemSteps object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWorkflowChainsItemStepsWithDefaults

`func NewListWorkflowChainsItemStepsWithDefaults() *ListWorkflowChainsItemSteps`

NewListWorkflowChainsItemStepsWithDefaults instantiates a new ListWorkflowChainsItemSteps object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListWorkflowChainsItemSteps) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListWorkflowChainsItemSteps) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListWorkflowChainsItemSteps) SetId(v string)`

SetId sets Id field to given value.


### GetPosition

`func (o *ListWorkflowChainsItemSteps) GetPosition() float32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *ListWorkflowChainsItemSteps) GetPositionOk() (*float32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *ListWorkflowChainsItemSteps) SetPosition(v float32)`

SetPosition sets Position field to given value.


### GetChainId

`func (o *ListWorkflowChainsItemSteps) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ListWorkflowChainsItemSteps) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ListWorkflowChainsItemSteps) SetChainId(v string)`

SetChainId sets ChainId field to given value.


### GetStepType

`func (o *ListWorkflowChainsItemSteps) GetStepType() string`

GetStepType returns the StepType field if non-nil, zero value otherwise.

### GetStepTypeOk

`func (o *ListWorkflowChainsItemSteps) GetStepTypeOk() (*string, bool)`

GetStepTypeOk returns a tuple with the StepType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStepType

`func (o *ListWorkflowChainsItemSteps) SetStepType(v string)`

SetStepType sets StepType field to given value.


### GetActionId

`func (o *ListWorkflowChainsItemSteps) GetActionId() string`

GetActionId returns the ActionId field if non-nil, zero value otherwise.

### GetActionIdOk

`func (o *ListWorkflowChainsItemSteps) GetActionIdOk() (*string, bool)`

GetActionIdOk returns a tuple with the ActionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionId

`func (o *ListWorkflowChainsItemSteps) SetActionId(v string)`

SetActionId sets ActionId field to given value.

### HasActionId

`func (o *ListWorkflowChainsItemSteps) HasActionId() bool`

HasActionId returns a boolean if a field has been set.

### SetActionIdNil

`func (o *ListWorkflowChainsItemSteps) SetActionIdNil(b bool)`

 SetActionIdNil sets the value for ActionId to be an explicit nil

### UnsetActionId
`func (o *ListWorkflowChainsItemSteps) UnsetActionId()`

UnsetActionId ensures that no value is present for ActionId, not even an explicit nil
### GetWaitMs

`func (o *ListWorkflowChainsItemSteps) GetWaitMs() float32`

GetWaitMs returns the WaitMs field if non-nil, zero value otherwise.

### GetWaitMsOk

`func (o *ListWorkflowChainsItemSteps) GetWaitMsOk() (*float32, bool)`

GetWaitMsOk returns a tuple with the WaitMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitMs

`func (o *ListWorkflowChainsItemSteps) SetWaitMs(v float32)`

SetWaitMs sets WaitMs field to given value.

### HasWaitMs

`func (o *ListWorkflowChainsItemSteps) HasWaitMs() bool`

HasWaitMs returns a boolean if a field has been set.

### SetWaitMsNil

`func (o *ListWorkflowChainsItemSteps) SetWaitMsNil(b bool)`

 SetWaitMsNil sets the value for WaitMs to be an explicit nil

### UnsetWaitMs
`func (o *ListWorkflowChainsItemSteps) UnsetWaitMs()`

UnsetWaitMs ensures that no value is present for WaitMs, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


