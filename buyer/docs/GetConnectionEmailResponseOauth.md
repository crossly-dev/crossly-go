# GetConnectionEmailResponseOauth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**IsConnected** | **bool** |  | 
**Email** | **string** |  | 
**LastSyncedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewGetConnectionEmailResponseOauth

`func NewGetConnectionEmailResponseOauth(id string, type_ string, isConnected bool, email string, ) *GetConnectionEmailResponseOauth`

NewGetConnectionEmailResponseOauth instantiates a new GetConnectionEmailResponseOauth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionEmailResponseOauthWithDefaults

`func NewGetConnectionEmailResponseOauthWithDefaults() *GetConnectionEmailResponseOauth`

NewGetConnectionEmailResponseOauthWithDefaults instantiates a new GetConnectionEmailResponseOauth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetConnectionEmailResponseOauth) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetConnectionEmailResponseOauth) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetConnectionEmailResponseOauth) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *GetConnectionEmailResponseOauth) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetConnectionEmailResponseOauth) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetConnectionEmailResponseOauth) SetType(v string)`

SetType sets Type field to given value.


### GetIsConnected

`func (o *GetConnectionEmailResponseOauth) GetIsConnected() bool`

GetIsConnected returns the IsConnected field if non-nil, zero value otherwise.

### GetIsConnectedOk

`func (o *GetConnectionEmailResponseOauth) GetIsConnectedOk() (*bool, bool)`

GetIsConnectedOk returns a tuple with the IsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnected

`func (o *GetConnectionEmailResponseOauth) SetIsConnected(v bool)`

SetIsConnected sets IsConnected field to given value.


### GetEmail

`func (o *GetConnectionEmailResponseOauth) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetConnectionEmailResponseOauth) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetConnectionEmailResponseOauth) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLastSyncedAt

`func (o *GetConnectionEmailResponseOauth) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *GetConnectionEmailResponseOauth) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *GetConnectionEmailResponseOauth) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *GetConnectionEmailResponseOauth) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### SetLastSyncedAtNil

`func (o *GetConnectionEmailResponseOauth) SetLastSyncedAtNil(b bool)`

 SetLastSyncedAtNil sets the value for LastSyncedAt to be an explicit nil

### UnsetLastSyncedAt
`func (o *GetConnectionEmailResponseOauth) UnsetLastSyncedAt()`

UnsetLastSyncedAt ensures that no value is present for LastSyncedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


