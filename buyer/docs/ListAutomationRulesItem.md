# ListAutomationRulesItem

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

### NewListAutomationRulesItem

`func NewListAutomationRulesItem(id string, name string, createdAt time.Time, updatedAt time.Time, userId string, isActive bool, triggerType string, actionType string, runCount float32, ) *ListAutomationRulesItem`

NewListAutomationRulesItem instantiates a new ListAutomationRulesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAutomationRulesItemWithDefaults

`func NewListAutomationRulesItemWithDefaults() *ListAutomationRulesItem`

NewListAutomationRulesItemWithDefaults instantiates a new ListAutomationRulesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListAutomationRulesItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAutomationRulesItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAutomationRulesItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListAutomationRulesItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListAutomationRulesItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListAutomationRulesItem) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *ListAutomationRulesItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListAutomationRulesItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListAutomationRulesItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListAutomationRulesItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListAutomationRulesItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListAutomationRulesItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *ListAutomationRulesItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListAutomationRulesItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListAutomationRulesItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetIsActive

`func (o *ListAutomationRulesItem) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ListAutomationRulesItem) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ListAutomationRulesItem) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetPlatforms

`func (o *ListAutomationRulesItem) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *ListAutomationRulesItem) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *ListAutomationRulesItem) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *ListAutomationRulesItem) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### SetPlatformsNil

`func (o *ListAutomationRulesItem) SetPlatformsNil(b bool)`

 SetPlatformsNil sets the value for Platforms to be an explicit nil

### UnsetPlatforms
`func (o *ListAutomationRulesItem) UnsetPlatforms()`

UnsetPlatforms ensures that no value is present for Platforms, not even an explicit nil
### GetTriggerType

`func (o *ListAutomationRulesItem) GetTriggerType() string`

GetTriggerType returns the TriggerType field if non-nil, zero value otherwise.

### GetTriggerTypeOk

`func (o *ListAutomationRulesItem) GetTriggerTypeOk() (*string, bool)`

GetTriggerTypeOk returns a tuple with the TriggerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerType

`func (o *ListAutomationRulesItem) SetTriggerType(v string)`

SetTriggerType sets TriggerType field to given value.


### GetConditionType

`func (o *ListAutomationRulesItem) GetConditionType() string`

GetConditionType returns the ConditionType field if non-nil, zero value otherwise.

### GetConditionTypeOk

`func (o *ListAutomationRulesItem) GetConditionTypeOk() (*string, bool)`

GetConditionTypeOk returns a tuple with the ConditionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionType

`func (o *ListAutomationRulesItem) SetConditionType(v string)`

SetConditionType sets ConditionType field to given value.

### HasConditionType

`func (o *ListAutomationRulesItem) HasConditionType() bool`

HasConditionType returns a boolean if a field has been set.

### SetConditionTypeNil

`func (o *ListAutomationRulesItem) SetConditionTypeNil(b bool)`

 SetConditionTypeNil sets the value for ConditionType to be an explicit nil

### UnsetConditionType
`func (o *ListAutomationRulesItem) UnsetConditionType()`

UnsetConditionType ensures that no value is present for ConditionType, not even an explicit nil
### GetActionType

`func (o *ListAutomationRulesItem) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *ListAutomationRulesItem) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *ListAutomationRulesItem) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetLastRunAt

`func (o *ListAutomationRulesItem) GetLastRunAt() time.Time`

GetLastRunAt returns the LastRunAt field if non-nil, zero value otherwise.

### GetLastRunAtOk

`func (o *ListAutomationRulesItem) GetLastRunAtOk() (*time.Time, bool)`

GetLastRunAtOk returns a tuple with the LastRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastRunAt

`func (o *ListAutomationRulesItem) SetLastRunAt(v time.Time)`

SetLastRunAt sets LastRunAt field to given value.

### HasLastRunAt

`func (o *ListAutomationRulesItem) HasLastRunAt() bool`

HasLastRunAt returns a boolean if a field has been set.

### SetLastRunAtNil

`func (o *ListAutomationRulesItem) SetLastRunAtNil(b bool)`

 SetLastRunAtNil sets the value for LastRunAt to be an explicit nil

### UnsetLastRunAt
`func (o *ListAutomationRulesItem) UnsetLastRunAt()`

UnsetLastRunAt ensures that no value is present for LastRunAt, not even an explicit nil
### GetNextRunAt

`func (o *ListAutomationRulesItem) GetNextRunAt() time.Time`

GetNextRunAt returns the NextRunAt field if non-nil, zero value otherwise.

### GetNextRunAtOk

`func (o *ListAutomationRulesItem) GetNextRunAtOk() (*time.Time, bool)`

GetNextRunAtOk returns a tuple with the NextRunAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunAt

`func (o *ListAutomationRulesItem) SetNextRunAt(v time.Time)`

SetNextRunAt sets NextRunAt field to given value.

### HasNextRunAt

`func (o *ListAutomationRulesItem) HasNextRunAt() bool`

HasNextRunAt returns a boolean if a field has been set.

### SetNextRunAtNil

`func (o *ListAutomationRulesItem) SetNextRunAtNil(b bool)`

 SetNextRunAtNil sets the value for NextRunAt to be an explicit nil

### UnsetNextRunAt
`func (o *ListAutomationRulesItem) UnsetNextRunAt()`

UnsetNextRunAt ensures that no value is present for NextRunAt, not even an explicit nil
### GetRunCount

`func (o *ListAutomationRulesItem) GetRunCount() float32`

GetRunCount returns the RunCount field if non-nil, zero value otherwise.

### GetRunCountOk

`func (o *ListAutomationRulesItem) GetRunCountOk() (*float32, bool)`

GetRunCountOk returns a tuple with the RunCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunCount

`func (o *ListAutomationRulesItem) SetRunCount(v float32)`

SetRunCount sets RunCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


