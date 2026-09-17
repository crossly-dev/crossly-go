# ListAutomationRunsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Status** | **string** |  | 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**DurationMs** | Pointer to **NullableFloat32** |  | [optional] 
**ActionType** | **string** |  | 
**RuleId** | Pointer to **NullableString** |  | [optional] 
**ChainId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListAutomationRunsItem

`func NewListAutomationRunsItem(id string, createdAt time.Time, userId string, status string, actionType string, ) *ListAutomationRunsItem`

NewListAutomationRunsItem instantiates a new ListAutomationRunsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAutomationRunsItemWithDefaults

`func NewListAutomationRunsItemWithDefaults() *ListAutomationRunsItem`

NewListAutomationRunsItemWithDefaults instantiates a new ListAutomationRunsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListAutomationRunsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAutomationRunsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAutomationRunsItem) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListAutomationRunsItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListAutomationRunsItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListAutomationRunsItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *ListAutomationRunsItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListAutomationRunsItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListAutomationRunsItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetStatus

`func (o *ListAutomationRunsItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListAutomationRunsItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListAutomationRunsItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetErrorMessage

`func (o *ListAutomationRunsItem) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListAutomationRunsItem) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListAutomationRunsItem) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListAutomationRunsItem) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ListAutomationRunsItem) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ListAutomationRunsItem) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetDurationMs

`func (o *ListAutomationRunsItem) GetDurationMs() float32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *ListAutomationRunsItem) GetDurationMsOk() (*float32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *ListAutomationRunsItem) SetDurationMs(v float32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *ListAutomationRunsItem) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### SetDurationMsNil

`func (o *ListAutomationRunsItem) SetDurationMsNil(b bool)`

 SetDurationMsNil sets the value for DurationMs to be an explicit nil

### UnsetDurationMs
`func (o *ListAutomationRunsItem) UnsetDurationMs()`

UnsetDurationMs ensures that no value is present for DurationMs, not even an explicit nil
### GetActionType

`func (o *ListAutomationRunsItem) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *ListAutomationRunsItem) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *ListAutomationRunsItem) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetRuleId

`func (o *ListAutomationRunsItem) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ListAutomationRunsItem) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ListAutomationRunsItem) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *ListAutomationRunsItem) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### SetRuleIdNil

`func (o *ListAutomationRunsItem) SetRuleIdNil(b bool)`

 SetRuleIdNil sets the value for RuleId to be an explicit nil

### UnsetRuleId
`func (o *ListAutomationRunsItem) UnsetRuleId()`

UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
### GetChainId

`func (o *ListAutomationRunsItem) GetChainId() string`

GetChainId returns the ChainId field if non-nil, zero value otherwise.

### GetChainIdOk

`func (o *ListAutomationRunsItem) GetChainIdOk() (*string, bool)`

GetChainIdOk returns a tuple with the ChainId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainId

`func (o *ListAutomationRunsItem) SetChainId(v string)`

SetChainId sets ChainId field to given value.

### HasChainId

`func (o *ListAutomationRunsItem) HasChainId() bool`

HasChainId returns a boolean if a field has been set.

### SetChainIdNil

`func (o *ListAutomationRunsItem) SetChainIdNil(b bool)`

 SetChainIdNil sets the value for ChainId to be an explicit nil

### UnsetChainId
`func (o *ListAutomationRunsItem) UnsetChainId()`

UnsetChainId ensures that no value is present for ChainId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


