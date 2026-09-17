# CreateAutomationRuleRunNowResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | **bool** |  | 
**JobId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateAutomationRuleRunNowResponse

`func NewCreateAutomationRuleRunNowResponse(success bool, ) *CreateAutomationRuleRunNowResponse`

NewCreateAutomationRuleRunNowResponse instantiates a new CreateAutomationRuleRunNowResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutomationRuleRunNowResponseWithDefaults

`func NewCreateAutomationRuleRunNowResponseWithDefaults() *CreateAutomationRuleRunNowResponse`

NewCreateAutomationRuleRunNowResponseWithDefaults instantiates a new CreateAutomationRuleRunNowResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *CreateAutomationRuleRunNowResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateAutomationRuleRunNowResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateAutomationRuleRunNowResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.


### GetJobId

`func (o *CreateAutomationRuleRunNowResponse) GetJobId() string`

GetJobId returns the JobId field if non-nil, zero value otherwise.

### GetJobIdOk

`func (o *CreateAutomationRuleRunNowResponse) GetJobIdOk() (*string, bool)`

GetJobIdOk returns a tuple with the JobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobId

`func (o *CreateAutomationRuleRunNowResponse) SetJobId(v string)`

SetJobId sets JobId field to given value.

### HasJobId

`func (o *CreateAutomationRuleRunNowResponse) HasJobId() bool`

HasJobId returns a boolean if a field has been set.

### SetJobIdNil

`func (o *CreateAutomationRuleRunNowResponse) SetJobIdNil(b bool)`

 SetJobIdNil sets the value for JobId to be an explicit nil

### UnsetJobId
`func (o *CreateAutomationRuleRunNowResponse) UnsetJobId()`

UnsetJobId ensures that no value is present for JobId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


