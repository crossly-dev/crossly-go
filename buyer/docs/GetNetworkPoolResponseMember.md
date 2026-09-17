# GetNetworkPoolResponseMember

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Platforms** | **string** |  | 
**JoinedAt** | **time.Time** |  | 
**ShareEnabled** | **bool** |  | 
**FollowEnabled** | **bool** |  | 
**LikeEnabled** | **bool** |  | 
**LastEngagementAt** | Pointer to **NullableTime** |  | [optional] 
**EngagementsPerformed** | **float32** |  | 
**EngagementsReceived** | **float32** |  | 
**FlaggedUntil** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewGetNetworkPoolResponseMember

`func NewGetNetworkPoolResponseMember(id string, userId string, platforms string, joinedAt time.Time, shareEnabled bool, followEnabled bool, likeEnabled bool, engagementsPerformed float32, engagementsReceived float32, ) *GetNetworkPoolResponseMember`

NewGetNetworkPoolResponseMember instantiates a new GetNetworkPoolResponseMember object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkPoolResponseMemberWithDefaults

`func NewGetNetworkPoolResponseMemberWithDefaults() *GetNetworkPoolResponseMember`

NewGetNetworkPoolResponseMemberWithDefaults instantiates a new GetNetworkPoolResponseMember object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkPoolResponseMember) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkPoolResponseMember) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkPoolResponseMember) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *GetNetworkPoolResponseMember) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetNetworkPoolResponseMember) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetNetworkPoolResponseMember) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetPlatforms

`func (o *GetNetworkPoolResponseMember) GetPlatforms() string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *GetNetworkPoolResponseMember) GetPlatformsOk() (*string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *GetNetworkPoolResponseMember) SetPlatforms(v string)`

SetPlatforms sets Platforms field to given value.


### GetJoinedAt

`func (o *GetNetworkPoolResponseMember) GetJoinedAt() time.Time`

GetJoinedAt returns the JoinedAt field if non-nil, zero value otherwise.

### GetJoinedAtOk

`func (o *GetNetworkPoolResponseMember) GetJoinedAtOk() (*time.Time, bool)`

GetJoinedAtOk returns a tuple with the JoinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinedAt

`func (o *GetNetworkPoolResponseMember) SetJoinedAt(v time.Time)`

SetJoinedAt sets JoinedAt field to given value.


### GetShareEnabled

`func (o *GetNetworkPoolResponseMember) GetShareEnabled() bool`

GetShareEnabled returns the ShareEnabled field if non-nil, zero value otherwise.

### GetShareEnabledOk

`func (o *GetNetworkPoolResponseMember) GetShareEnabledOk() (*bool, bool)`

GetShareEnabledOk returns a tuple with the ShareEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareEnabled

`func (o *GetNetworkPoolResponseMember) SetShareEnabled(v bool)`

SetShareEnabled sets ShareEnabled field to given value.


### GetFollowEnabled

`func (o *GetNetworkPoolResponseMember) GetFollowEnabled() bool`

GetFollowEnabled returns the FollowEnabled field if non-nil, zero value otherwise.

### GetFollowEnabledOk

`func (o *GetNetworkPoolResponseMember) GetFollowEnabledOk() (*bool, bool)`

GetFollowEnabledOk returns a tuple with the FollowEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFollowEnabled

`func (o *GetNetworkPoolResponseMember) SetFollowEnabled(v bool)`

SetFollowEnabled sets FollowEnabled field to given value.


### GetLikeEnabled

`func (o *GetNetworkPoolResponseMember) GetLikeEnabled() bool`

GetLikeEnabled returns the LikeEnabled field if non-nil, zero value otherwise.

### GetLikeEnabledOk

`func (o *GetNetworkPoolResponseMember) GetLikeEnabledOk() (*bool, bool)`

GetLikeEnabledOk returns a tuple with the LikeEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLikeEnabled

`func (o *GetNetworkPoolResponseMember) SetLikeEnabled(v bool)`

SetLikeEnabled sets LikeEnabled field to given value.


### GetLastEngagementAt

`func (o *GetNetworkPoolResponseMember) GetLastEngagementAt() time.Time`

GetLastEngagementAt returns the LastEngagementAt field if non-nil, zero value otherwise.

### GetLastEngagementAtOk

`func (o *GetNetworkPoolResponseMember) GetLastEngagementAtOk() (*time.Time, bool)`

GetLastEngagementAtOk returns a tuple with the LastEngagementAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastEngagementAt

`func (o *GetNetworkPoolResponseMember) SetLastEngagementAt(v time.Time)`

SetLastEngagementAt sets LastEngagementAt field to given value.

### HasLastEngagementAt

`func (o *GetNetworkPoolResponseMember) HasLastEngagementAt() bool`

HasLastEngagementAt returns a boolean if a field has been set.

### SetLastEngagementAtNil

`func (o *GetNetworkPoolResponseMember) SetLastEngagementAtNil(b bool)`

 SetLastEngagementAtNil sets the value for LastEngagementAt to be an explicit nil

### UnsetLastEngagementAt
`func (o *GetNetworkPoolResponseMember) UnsetLastEngagementAt()`

UnsetLastEngagementAt ensures that no value is present for LastEngagementAt, not even an explicit nil
### GetEngagementsPerformed

`func (o *GetNetworkPoolResponseMember) GetEngagementsPerformed() float32`

GetEngagementsPerformed returns the EngagementsPerformed field if non-nil, zero value otherwise.

### GetEngagementsPerformedOk

`func (o *GetNetworkPoolResponseMember) GetEngagementsPerformedOk() (*float32, bool)`

GetEngagementsPerformedOk returns a tuple with the EngagementsPerformed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementsPerformed

`func (o *GetNetworkPoolResponseMember) SetEngagementsPerformed(v float32)`

SetEngagementsPerformed sets EngagementsPerformed field to given value.


### GetEngagementsReceived

`func (o *GetNetworkPoolResponseMember) GetEngagementsReceived() float32`

GetEngagementsReceived returns the EngagementsReceived field if non-nil, zero value otherwise.

### GetEngagementsReceivedOk

`func (o *GetNetworkPoolResponseMember) GetEngagementsReceivedOk() (*float32, bool)`

GetEngagementsReceivedOk returns a tuple with the EngagementsReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEngagementsReceived

`func (o *GetNetworkPoolResponseMember) SetEngagementsReceived(v float32)`

SetEngagementsReceived sets EngagementsReceived field to given value.


### GetFlaggedUntil

`func (o *GetNetworkPoolResponseMember) GetFlaggedUntil() time.Time`

GetFlaggedUntil returns the FlaggedUntil field if non-nil, zero value otherwise.

### GetFlaggedUntilOk

`func (o *GetNetworkPoolResponseMember) GetFlaggedUntilOk() (*time.Time, bool)`

GetFlaggedUntilOk returns a tuple with the FlaggedUntil field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlaggedUntil

`func (o *GetNetworkPoolResponseMember) SetFlaggedUntil(v time.Time)`

SetFlaggedUntil sets FlaggedUntil field to given value.

### HasFlaggedUntil

`func (o *GetNetworkPoolResponseMember) HasFlaggedUntil() bool`

HasFlaggedUntil returns a boolean if a field has been set.

### SetFlaggedUntilNil

`func (o *GetNetworkPoolResponseMember) SetFlaggedUntilNil(b bool)`

 SetFlaggedUntilNil sets the value for FlaggedUntil to be an explicit nil

### UnsetFlaggedUntil
`func (o *GetNetworkPoolResponseMember) UnsetFlaggedUntil()`

UnsetFlaggedUntil ensures that no value is present for FlaggedUntil, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


