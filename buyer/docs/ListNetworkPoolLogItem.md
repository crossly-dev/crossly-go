# ListNetworkPoolLogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Platform** | **string** |  | 
**ActionType** | **string** |  | 
**TargetUserId** | **string** |  | 
**PerformedAt** | **time.Time** |  | 
**RuleId** | Pointer to **NullableString** |  | [optional] 
**EngagerUserId** | **string** |  | 
**TargetListingId** | Pointer to **NullableString** |  | [optional] 
**TargetPlatformListingId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListNetworkPoolLogItem

`func NewListNetworkPoolLogItem(id string, platform string, actionType string, targetUserId string, performedAt time.Time, engagerUserId string, ) *ListNetworkPoolLogItem`

NewListNetworkPoolLogItem instantiates a new ListNetworkPoolLogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkPoolLogItemWithDefaults

`func NewListNetworkPoolLogItemWithDefaults() *ListNetworkPoolLogItem`

NewListNetworkPoolLogItemWithDefaults instantiates a new ListNetworkPoolLogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListNetworkPoolLogItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListNetworkPoolLogItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListNetworkPoolLogItem) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *ListNetworkPoolLogItem) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListNetworkPoolLogItem) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListNetworkPoolLogItem) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetActionType

`func (o *ListNetworkPoolLogItem) GetActionType() string`

GetActionType returns the ActionType field if non-nil, zero value otherwise.

### GetActionTypeOk

`func (o *ListNetworkPoolLogItem) GetActionTypeOk() (*string, bool)`

GetActionTypeOk returns a tuple with the ActionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionType

`func (o *ListNetworkPoolLogItem) SetActionType(v string)`

SetActionType sets ActionType field to given value.


### GetTargetUserId

`func (o *ListNetworkPoolLogItem) GetTargetUserId() string`

GetTargetUserId returns the TargetUserId field if non-nil, zero value otherwise.

### GetTargetUserIdOk

`func (o *ListNetworkPoolLogItem) GetTargetUserIdOk() (*string, bool)`

GetTargetUserIdOk returns a tuple with the TargetUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetUserId

`func (o *ListNetworkPoolLogItem) SetTargetUserId(v string)`

SetTargetUserId sets TargetUserId field to given value.


### GetPerformedAt

`func (o *ListNetworkPoolLogItem) GetPerformedAt() time.Time`

GetPerformedAt returns the PerformedAt field if non-nil, zero value otherwise.

### GetPerformedAtOk

`func (o *ListNetworkPoolLogItem) GetPerformedAtOk() (*time.Time, bool)`

GetPerformedAtOk returns a tuple with the PerformedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerformedAt

`func (o *ListNetworkPoolLogItem) SetPerformedAt(v time.Time)`

SetPerformedAt sets PerformedAt field to given value.


### GetRuleId

`func (o *ListNetworkPoolLogItem) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ListNetworkPoolLogItem) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ListNetworkPoolLogItem) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.

### HasRuleId

`func (o *ListNetworkPoolLogItem) HasRuleId() bool`

HasRuleId returns a boolean if a field has been set.

### SetRuleIdNil

`func (o *ListNetworkPoolLogItem) SetRuleIdNil(b bool)`

 SetRuleIdNil sets the value for RuleId to be an explicit nil

### UnsetRuleId
`func (o *ListNetworkPoolLogItem) UnsetRuleId()`

UnsetRuleId ensures that no value is present for RuleId, not even an explicit nil
### GetEngagerUserId

`func (o *ListNetworkPoolLogItem) GetEngagerUserId() string`

GetEngagerUserId returns the EngagerUserId field if non-nil, zero value otherwise.

### GetEngagerUserIdOk

`func (o *ListNetworkPoolLogItem) GetEngagerUserIdOk() (*string, bool)`

GetEngagerUserIdOk returns a tuple with the EngagerUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagerUserId

`func (o *ListNetworkPoolLogItem) SetEngagerUserId(v string)`

SetEngagerUserId sets EngagerUserId field to given value.


### GetTargetListingId

`func (o *ListNetworkPoolLogItem) GetTargetListingId() string`

GetTargetListingId returns the TargetListingId field if non-nil, zero value otherwise.

### GetTargetListingIdOk

`func (o *ListNetworkPoolLogItem) GetTargetListingIdOk() (*string, bool)`

GetTargetListingIdOk returns a tuple with the TargetListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetListingId

`func (o *ListNetworkPoolLogItem) SetTargetListingId(v string)`

SetTargetListingId sets TargetListingId field to given value.

### HasTargetListingId

`func (o *ListNetworkPoolLogItem) HasTargetListingId() bool`

HasTargetListingId returns a boolean if a field has been set.

### SetTargetListingIdNil

`func (o *ListNetworkPoolLogItem) SetTargetListingIdNil(b bool)`

 SetTargetListingIdNil sets the value for TargetListingId to be an explicit nil

### UnsetTargetListingId
`func (o *ListNetworkPoolLogItem) UnsetTargetListingId()`

UnsetTargetListingId ensures that no value is present for TargetListingId, not even an explicit nil
### GetTargetPlatformListingId

`func (o *ListNetworkPoolLogItem) GetTargetPlatformListingId() string`

GetTargetPlatformListingId returns the TargetPlatformListingId field if non-nil, zero value otherwise.

### GetTargetPlatformListingIdOk

`func (o *ListNetworkPoolLogItem) GetTargetPlatformListingIdOk() (*string, bool)`

GetTargetPlatformListingIdOk returns a tuple with the TargetPlatformListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetPlatformListingId

`func (o *ListNetworkPoolLogItem) SetTargetPlatformListingId(v string)`

SetTargetPlatformListingId sets TargetPlatformListingId field to given value.

### HasTargetPlatformListingId

`func (o *ListNetworkPoolLogItem) HasTargetPlatformListingId() bool`

HasTargetPlatformListingId returns a boolean if a field has been set.

### SetTargetPlatformListingIdNil

`func (o *ListNetworkPoolLogItem) SetTargetPlatformListingIdNil(b bool)`

 SetTargetPlatformListingIdNil sets the value for TargetPlatformListingId to be an explicit nil

### UnsetTargetPlatformListingId
`func (o *ListNetworkPoolLogItem) UnsetTargetPlatformListingId()`

UnsetTargetPlatformListingId ensures that no value is present for TargetPlatformListingId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


