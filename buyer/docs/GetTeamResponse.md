# GetTeamResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invitations** | [**[]GetTeamResponseInvitations**](GetTeamResponseInvitations.md) |  | 
**Members** | [**[]GetTeamResponseMembers**](GetTeamResponseMembers.md) |  | 

## Methods

### NewGetTeamResponse

`func NewGetTeamResponse(invitations []GetTeamResponseInvitations, members []GetTeamResponseMembers, ) *GetTeamResponse`

NewGetTeamResponse instantiates a new GetTeamResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTeamResponseWithDefaults

`func NewGetTeamResponseWithDefaults() *GetTeamResponse`

NewGetTeamResponseWithDefaults instantiates a new GetTeamResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvitations

`func (o *GetTeamResponse) GetInvitations() []GetTeamResponseInvitations`

GetInvitations returns the Invitations field if non-nil, zero value otherwise.

### GetInvitationsOk

`func (o *GetTeamResponse) GetInvitationsOk() (*[]GetTeamResponseInvitations, bool)`

GetInvitationsOk returns a tuple with the Invitations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvitations

`func (o *GetTeamResponse) SetInvitations(v []GetTeamResponseInvitations)`

SetInvitations sets Invitations field to given value.


### GetMembers

`func (o *GetTeamResponse) GetMembers() []GetTeamResponseMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *GetTeamResponse) GetMembersOk() (*[]GetTeamResponseMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *GetTeamResponse) SetMembers(v []GetTeamResponseMembers)`

SetMembers sets Members field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


