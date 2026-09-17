# GetTeamResponseInvitations

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Pending** | **bool** |  | 
**Id** | **string** |  | 
**InviteeEmail** | **string** |  | 
**Role** | **string** |  | 
**Scopes** | **[]string** |  | 
**ExpiresAt** | **time.Time** |  | 
**AcceptedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewGetTeamResponseInvitations

`func NewGetTeamResponseInvitations(pending bool, id string, inviteeEmail string, role string, scopes []string, expiresAt time.Time, createdAt time.Time, ) *GetTeamResponseInvitations`

NewGetTeamResponseInvitations instantiates a new GetTeamResponseInvitations object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTeamResponseInvitationsWithDefaults

`func NewGetTeamResponseInvitationsWithDefaults() *GetTeamResponseInvitations`

NewGetTeamResponseInvitationsWithDefaults instantiates a new GetTeamResponseInvitations object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPending

`func (o *GetTeamResponseInvitations) GetPending() bool`

GetPending returns the Pending field if non-nil, zero value otherwise.

### GetPendingOk

`func (o *GetTeamResponseInvitations) GetPendingOk() (*bool, bool)`

GetPendingOk returns a tuple with the Pending field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPending

`func (o *GetTeamResponseInvitations) SetPending(v bool)`

SetPending sets Pending field to given value.


### GetId

`func (o *GetTeamResponseInvitations) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTeamResponseInvitations) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTeamResponseInvitations) SetId(v string)`

SetId sets Id field to given value.


### GetInviteeEmail

`func (o *GetTeamResponseInvitations) GetInviteeEmail() string`

GetInviteeEmail returns the InviteeEmail field if non-nil, zero value otherwise.

### GetInviteeEmailOk

`func (o *GetTeamResponseInvitations) GetInviteeEmailOk() (*string, bool)`

GetInviteeEmailOk returns a tuple with the InviteeEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInviteeEmail

`func (o *GetTeamResponseInvitations) SetInviteeEmail(v string)`

SetInviteeEmail sets InviteeEmail field to given value.


### GetRole

`func (o *GetTeamResponseInvitations) GetRole() string`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *GetTeamResponseInvitations) GetRoleOk() (*string, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *GetTeamResponseInvitations) SetRole(v string)`

SetRole sets Role field to given value.


### GetScopes

`func (o *GetTeamResponseInvitations) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GetTeamResponseInvitations) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GetTeamResponseInvitations) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetExpiresAt

`func (o *GetTeamResponseInvitations) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetTeamResponseInvitations) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetTeamResponseInvitations) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetAcceptedAt

`func (o *GetTeamResponseInvitations) GetAcceptedAt() time.Time`

GetAcceptedAt returns the AcceptedAt field if non-nil, zero value otherwise.

### GetAcceptedAtOk

`func (o *GetTeamResponseInvitations) GetAcceptedAtOk() (*time.Time, bool)`

GetAcceptedAtOk returns a tuple with the AcceptedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcceptedAt

`func (o *GetTeamResponseInvitations) SetAcceptedAt(v time.Time)`

SetAcceptedAt sets AcceptedAt field to given value.

### HasAcceptedAt

`func (o *GetTeamResponseInvitations) HasAcceptedAt() bool`

HasAcceptedAt returns a boolean if a field has been set.

### SetAcceptedAtNil

`func (o *GetTeamResponseInvitations) SetAcceptedAtNil(b bool)`

 SetAcceptedAtNil sets the value for AcceptedAt to be an explicit nil

### UnsetAcceptedAt
`func (o *GetTeamResponseInvitations) UnsetAcceptedAt()`

UnsetAcceptedAt ensures that no value is present for AcceptedAt, not even an explicit nil
### GetCreatedAt

`func (o *GetTeamResponseInvitations) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetTeamResponseInvitations) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetTeamResponseInvitations) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


