# CreateAutomationRuleResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**IsActive** | **bool** |  | 
**Platforms** | Pointer to **[]string** |  | [optional] 
**TriggerType** | **string** |  | 
**ConditionType** | Pointer to **NullableString** |  | [optional] 
**ActionType** | **string** |  | 
**LastRunAt** | Pointer to **NullableTime** |  | [optional] 
**NextRunAt** | Pointer to **NullableTime** |  | [optional] 
**RunCount** | **float32** |  | 

## Methods

### NewCreateAutomationRuleResponse

`func NewCreateAutomationRuleResponse(id string, name string, createdAt time.Time, updatedAt time.Time, userId string, isActive bool, triggerType string, actionType string, runCount float32, ) *CreateAutomationRuleResponse`

NewCreateAutomationRuleResponse instantiates a new CreateAutomationRuleResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutomationRuleResponseWithDefaults

`func NewCreateAutomationRuleResponseWithDefaults() *CreateAutomationRuleResponse`

NewCreateAutomationRuleResponseWithDefaults instantiates a new CreateAutomationRuleResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateAutomationRuleResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAutomationRuleResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAutomationRuleResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreateAutomationRuleResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateAutomationRuleResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateAutomationRuleResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *CreateAutomationRuleResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateAutomationRuleResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateAutomationRuleResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreateAutomationRuleResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateAutomationRuleResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateAutomationRuleResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *CreateAutomationRuleResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateAutomationRuleResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateAutomationRuleResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetIsActive

`func (o *CreateAutomationRuleResponse) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *CreateAutomationRuleResponse) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *CreateAutomationRuleResponse) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetPlatforms

`func (o *CreateAutomationRuleResponse) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *CreateAutomationRuleResponse) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *CreateAutomationRuleResponse) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *CreateAutomationRuleResponse) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### SetPlatformsNil

`func (o *CreateAutomationRuleResponse) SetPlatformsNil(b bool)`

 SetPlatformsNil sets the value for Platforms to be an explicit nil

### UnsetPlatforms
`func (o *CreateAutomationRuleResponse) UnsetPlatforms()`

UnsetPlatforms ensures that no value is present for Platforms, not even an explicit nil
### GetTriggerType

`func (o *CreateAutomationRuleResponse) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *CreateAutomationRuleResponse) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *CreateAutomationRuleResponse) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.


### GetConditionType

`func (o *CreateAutomationRuleResponse) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *CreateAutomationRuleResponse) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *CreateAutomationRuleResponse) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.

### HasConditionType

`func (o *CreateAutomationRuleResponse) HasConditionType() bool`

HasConditionType returns a boolean if a field has been set.

### SetConditionTypeNil

`func (o *CreateAutomationRuleResponse) SetConditionTypeNil(b bool)`

 SetConditionTypeNil sets the value for ConditionType to be an explicit nil

### UnsetConditionType
`func (o *CreateAutomationRuleResponse) UnsetConditionType()`

UnsetConditionType ensures that no value is present for ConditionType, not even an explicit nil
### GetActionType

`func (o *CreateAutomationRuleResponse) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *CreateAutomationRuleResponse) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *CreateAutomationRuleResponse) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetLastRunAt

`func (o *CreateAutomationRuleResponse) GetLastRunAt() time.Time`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *CreateAutomationRuleResponse) GetLastRunAtOk() (*time.Time, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *CreateAutomationRuleResponse) SetLastRunAt(v time.Time)`

SetLastRunAt sets LastRunAt field to given value.

### HasLastRunAt

`func (o *CreateAutomationRuleResponse) HasLastRunAt() bool`

HasLastRunAt returns a boolean if a field has been set.

### SetLastRunAtNil

`func (o *CreateAutomationRuleResponse) SetLastRunAtNil(b bool)`

 SetLastRunAtNil sets the value for LastRunAt to be an explicit nil

### UnsetLastRunAt
`func (o *CreateAutomationRuleResponse) UnsetLastRunAt()`

UnsetLastRunAt ensures that no value is present for LastRunAt, not even an explicit nil
### GetNextRunAt

`func (o *CreateAutomationRuleResponse) GetNextRunAt() time.Time`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *CreateAutomationRuleResponse) GetNextRunAtOk() (*time.Time, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *CreateAutomationRuleResponse) SetNextRunAt(v time.Time)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *CreateAutomationRuleResponse) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### SetNextRunAtNil

`func (o *CreateAutomationRuleResponse) SetNextRunAtNil(b bool)`

 SetNextRunAtNil sets the value for NextRunAt to be an explicit nil

### UnsetNextRunAt
`func (o *CreateAutomationRuleResponse) UnsetNextRunAt()`

UnsetNextRunAt ensures that no value is present for NextRunAt, not even an explicit nil
### GetRunCount

`func (o *CreateAutomationRuleResponse) GetRunCount() float32`

GetRunCount returns the RunCount field if non-nil, zero value otherwise.

### GetRunCountOk

`func (o *CreateAutomationRuleResponse) GetRunCountOk() (*float32, bool)`

GetRunCountOk returns a tuple with the RunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCount

`func (o *CreateAutomationRuleResponse) SetRunCount(v float32)`

SetRunCount sets RunCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


