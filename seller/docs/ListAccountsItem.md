# ListAccountsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Platform** | **string** |  | 
**AccountSlot** | **float32** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**PlatformUsername** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**LastSyncedAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewListAccountsItem

`func NewListAccountsItem(id string, platform string, accountSlot float32, status string, createdAt time.Time, ) *ListAccountsItem`

NewListAccountsItem instantiates a new ListAccountsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAccountsItemWithDefaults

`func NewListAccountsItemWithDefaults() *ListAccountsItem`

NewListAccountsItemWithDefaults instantiates a new ListAccountsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListAccountsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAccountsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAccountsItem) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *ListAccountsItem) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListAccountsItem) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListAccountsItem) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetAccountSlot

`func (o *ListAccountsItem) GetAccountSlot() float32`

GetAccountSlot returns the AccountSlot field if non-nil, zero value otherwise.

### GetAccountSlotOk

`func (o *ListAccountsItem) GetAccountSlotOk() (*float32, bool)`

GetAccountSlotOk returns a tuple with the AccountSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSlot

`func (o *ListAccountsItem) SetAccountSlot(v float32)`

SetAccountSlot sets AccountSlot field to given value.


### GetLabel

`func (o *ListAccountsItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ListAccountsItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ListAccountsItem) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ListAccountsItem) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *ListAccountsItem) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *ListAccountsItem) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetPlatformUsername

`func (o *ListAccountsItem) GetPlatformUsername() string`

GetPlatformUsername returns the PlatformUsername field if non-nil, zero value otherwise.

### GetPlatformUsernameOk

`func (o *ListAccountsItem) GetPlatformUsernameOk() (*string, bool)`

GetPlatformUsernameOk returns a tuple with the PlatformUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUsername

`func (o *ListAccountsItem) SetPlatformUsername(v string)`

SetPlatformUsername sets PlatformUsername field to given value.

### HasPlatformUsername

`func (o *ListAccountsItem) HasPlatformUsername() bool`

HasPlatformUsername returns a boolean if a field has been set.

### SetPlatformUsernameNil

`func (o *ListAccountsItem) SetPlatformUsernameNil(b bool)`

 SetPlatformUsernameNil sets the value for PlatformUsername to be an explicit nil

### UnsetPlatformUsername
`func (o *ListAccountsItem) UnsetPlatformUsername()`

UnsetPlatformUsername ensures that no value is present for PlatformUsername, not even an explicit nil
### GetStatus

`func (o *ListAccountsItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListAccountsItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListAccountsItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLastSyncedAt

`func (o *ListAccountsItem) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *ListAccountsItem) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *ListAccountsItem) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *ListAccountsItem) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### SetLastSyncedAtNil

`func (o *ListAccountsItem) SetLastSyncedAtNil(b bool)`

 SetLastSyncedAtNil sets the value for LastSyncedAt to be an explicit nil

### UnsetLastSyncedAt
`func (o *ListAccountsItem) UnsetLastSyncedAt()`

UnsetLastSyncedAt ensures that no value is present for LastSyncedAt, not even an explicit nil
### GetCreatedAt

`func (o *ListAccountsItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListAccountsItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListAccountsItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


