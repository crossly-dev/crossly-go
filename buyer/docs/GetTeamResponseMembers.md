# GetTeamResponseMembers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MemberUserId** | **string** |  | 
**MemberEmail** | Pointer to **NullableString** |  | [optional] 
**MemberDisplayName** | Pointer to **NullableString** |  | [optional] 
**Role** | **string** |  | 
**Scopes** | **[]string** |  | 
**RevokedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewGetTeamResponseMembers

`func NewGetTeamResponseMembers(id string, memberUserId string, role string, scopes []string, createdAt time.Time, ) *GetTeamResponseMembers`

NewGetTeamResponseMembers instantiates a new GetTeamResponseMembers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTeamResponseMembersWithDefaults

`func NewGetTeamResponseMembersWithDefaults() *GetTeamResponseMembers`

NewGetTeamResponseMembersWithDefaults instantiates a new GetTeamResponseMembers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTeamResponseMembers) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTeamResponseMembers) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTeamResponseMembers) SetId(v string)`

SetId sets Id field to given value.


### GetMemberUserId

`func (o *GetTeamResponseMembers) GetMemberUserId() string`

GetMemberUserId returns the MemberUserId field if non-nil, zero value otherwise.

### GetMemberUserIdOk

`func (o *GetTeamResponseMembers) GetMemberUserIdOk() (*string, bool)`

GetMemberUserIdOk returns a tuple with the MemberUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberUserId

`func (o *GetTeamResponseMembers) SetMemberUserId(v string)`

SetMemberUserId sets MemberUserId field to given value.


### GetMemberEmail

`func (o *GetTeamResponseMembers) GetMemberEmail() string`

GetMemberEmail returns the MemberEmail field if non-nil, zero value otherwise.

### GetMemberEmailOk

`func (o *GetTeamResponseMembers) GetMemberEmailOk() (*string, bool)`

GetMemberEmailOk returns a tuple with the MemberEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberEmail

`func (o *GetTeamResponseMembers) SetMemberEmail(v string)`

SetMemberEmail sets MemberEmail field to given value.

### HasMemberEmail

`func (o *GetTeamResponseMembers) HasMemberEmail() bool`

HasMemberEmail returns a boolean if a field has been set.

### SetMemberEmailNil

`func (o *GetTeamResponseMembers) SetMemberEmailNil(b bool)`

 SetMemberEmailNil sets the value for MemberEmail to be an explicit nil

### UnsetMemberEmail
`func (o *GetTeamResponseMembers) UnsetMemberEmail()`

UnsetMemberEmail ensures that no value is present for MemberEmail, not even an explicit nil
### GetMemberDisplayName

`func (o *GetTeamResponseMembers) GetMemberDisplayName() string`

GetMemberDisplayName returns the MemberDisplayName field if non-nil, zero value otherwise.

### GetMemberDisplayNameOk

`func (o *GetTeamResponseMembers) GetMemberDisplayNameOk() (*string, bool)`

GetMemberDisplayNameOk returns a tuple with the MemberDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberDisplayName

`func (o *GetTeamResponseMembers) SetMemberDisplayName(v string)`

SetMemberDisplayName sets MemberDisplayName field to given value.

### HasMemberDisplayName

`func (o *GetTeamResponseMembers) HasMemberDisplayName() bool`

HasMemberDisplayName returns a boolean if a field has been set.

### SetMemberDisplayNameNil

`func (o *GetTeamResponseMembers) SetMemberDisplayNameNil(b bool)`

 SetMemberDisplayNameNil sets the value for MemberDisplayName to be an explicit nil

### UnsetMemberDisplayName
`func (o *GetTeamResponseMembers) UnsetMemberDisplayName()`

UnsetMemberDisplayName ensures that no value is present for MemberDisplayName, not even an explicit nil
### GetRole

`func (o *GetTeamResponseMembers) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetTeamResponseMembers) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetTeamResponseMembers) SetRole(v string)`

SetRole sets Role field to given value.


### GetScopes

`func (o *GetTeamResponseMembers) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GetTeamResponseMembers) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GetTeamResponseMembers) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetRevokedAt

`func (o *GetTeamResponseMembers) GetRevokedAt() time.Time`

GetRevokedAt returns the RevokedAt field if non-nil, zero value otherwise.

### GetRevokedAtOk

`func (o *GetTeamResponseMembers) GetRevokedAtOk() (*time.Time, bool)`

GetRevokedAtOk returns a tuple with the RevokedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevokedAt

`func (o *GetTeamResponseMembers) SetRevokedAt(v time.Time)`

SetRevokedAt sets RevokedAt field to given value.

### HasRevokedAt

`func (o *GetTeamResponseMembers) HasRevokedAt() bool`

HasRevokedAt returns a boolean if a field has been set.

### SetRevokedAtNil

`func (o *GetTeamResponseMembers) SetRevokedAtNil(b bool)`

 SetRevokedAtNil sets the value for RevokedAt to be an explicit nil

### UnsetRevokedAt
`func (o *GetTeamResponseMembers) UnsetRevokedAt()`

UnsetRevokedAt ensures that no value is present for RevokedAt, not even an explicit nil
### GetCreatedAt

`func (o *GetTeamResponseMembers) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetTeamResponseMembers) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetTeamResponseMembers) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


