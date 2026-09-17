# GetAutomationRuleResponseRule

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

### NewGetAutomationRuleResponseRule

`func NewGetAutomationRuleResponseRule(id string, name string, createdAt time.Time, updatedAt time.Time, userId string, isActive bool, triggerType string, actionType string, runCount float32, ) *GetAutomationRuleResponseRule`

NewGetAutomationRuleResponseRule instantiates a new GetAutomationRuleResponseRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAutomationRuleResponseRuleWithDefaults

`func NewGetAutomationRuleResponseRuleWithDefaults() *GetAutomationRuleResponseRule`

NewGetAutomationRuleResponseRuleWithDefaults instantiates a new GetAutomationRuleResponseRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAutomationRuleResponseRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAutomationRuleResponseRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAutomationRuleResponseRule) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *GetAutomationRuleResponseRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAutomationRuleResponseRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAutomationRuleResponseRule) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *GetAutomationRuleResponseRule) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetAutomationRuleResponseRule) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetAutomationRuleResponseRule) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetAutomationRuleResponseRule) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetAutomationRuleResponseRule) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetAutomationRuleResponseRule) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *GetAutomationRuleResponseRule) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetAutomationRuleResponseRule) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetAutomationRuleResponseRule) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetIsActive

`func (o *GetAutomationRuleResponseRule) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *GetAutomationRuleResponseRule) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *GetAutomationRuleResponseRule) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetPlatforms

`func (o *GetAutomationRuleResponseRule) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *GetAutomationRuleResponseRule) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *GetAutomationRuleResponseRule) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *GetAutomationRuleResponseRule) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### SetPlatformsNil

`func (o *GetAutomationRuleResponseRule) SetPlatformsNil(b bool)`

 SetPlatformsNil sets the value for Platforms to be an explicit nil

### UnsetPlatforms
`func (o *GetAutomationRuleResponseRule) UnsetPlatforms()`

UnsetPlatforms ensures that no value is present for Platforms, not even an explicit nil
### GetTriggerType

`func (o *GetAutomationRuleResponseRule) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *GetAutomationRuleResponseRule) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *GetAutomationRuleResponseRule) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.


### GetConditionType

`func (o *GetAutomationRuleResponseRule) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *GetAutomationRuleResponseRule) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *GetAutomationRuleResponseRule) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.

### HasConditionType

`func (o *GetAutomationRuleResponseRule) HasConditionType() bool`

HasConditionType returns a boolean if a field has been set.

### SetConditionTypeNil

`func (o *GetAutomationRuleResponseRule) SetConditionTypeNil(b bool)`

 SetConditionTypeNil sets the value for ConditionType to be an explicit nil

### UnsetConditionType
`func (o *GetAutomationRuleResponseRule) UnsetConditionType()`

UnsetConditionType ensures that no value is present for ConditionType, not even an explicit nil
### GetActionType

`func (o *GetAutomationRuleResponseRule) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *GetAutomationRuleResponseRule) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *GetAutomationRuleResponseRule) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetLastRunAt

`func (o *GetAutomationRuleResponseRule) GetLastRunAt() time.Time`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *GetAutomationRuleResponseRule) GetLastRunAtOk() (*time.Time, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *GetAutomationRuleResponseRule) SetLastRunAt(v time.Time)`

SetLastRunAt sets LastRunAt field to given value.

### HasLastRunAt

`func (o *GetAutomationRuleResponseRule) HasLastRunAt() bool`

HasLastRunAt returns a boolean if a field has been set.

### SetLastRunAtNil

`func (o *GetAutomationRuleResponseRule) SetLastRunAtNil(b bool)`

 SetLastRunAtNil sets the value for LastRunAt to be an explicit nil

### UnsetLastRunAt
`func (o *GetAutomationRuleResponseRule) UnsetLastRunAt()`

UnsetLastRunAt ensures that no value is present for LastRunAt, not even an explicit nil
### GetNextRunAt

`func (o *GetAutomationRuleResponseRule) GetNextRunAt() time.Time`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *GetAutomationRuleResponseRule) GetNextRunAtOk() (*time.Time, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *GetAutomationRuleResponseRule) SetNextRunAt(v time.Time)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *GetAutomationRuleResponseRule) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### SetNextRunAtNil

`func (o *GetAutomationRuleResponseRule) SetNextRunAtNil(b bool)`

 SetNextRunAtNil sets the value for NextRunAt to be an explicit nil

### UnsetNextRunAt
`func (o *GetAutomationRuleResponseRule) UnsetNextRunAt()`

UnsetNextRunAt ensures that no value is present for NextRunAt, not even an explicit nil
### GetRunCount

`func (o *GetAutomationRuleResponseRule) GetRunCount() float32`

GetRunCount returns the RunCount field if non-nil, zero value otherwise.

### GetRunCountOk

`func (o *GetAutomationRuleResponseRule) GetRunCountOk() (*float32, bool)`

GetRunCountOk returns a tuple with the RunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCount

`func (o *GetAutomationRuleResponseRule) SetRunCount(v float32)`

SetRunCount sets RunCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


