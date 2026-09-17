# GetConnectionEmailResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Type** | **string** |  | 
**IsConnected** | **bool** |  | 
**Email** | **string** |  | 
**Host** | Pointer to **NullableString** |  | [optional] 
**Port** | Pointer to **NullableFloat32** |  | [optional] 
**LastSyncedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewGetConnectionEmailResponseData

`func NewGetConnectionEmailResponseData(id string, type_ string, isConnected bool, email string, ) *GetConnectionEmailResponseData`

NewGetConnectionEmailResponseData instantiates a new GetConnectionEmailResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionEmailResponseDataWithDefaults

`func NewGetConnectionEmailResponseDataWithDefaults() *GetConnectionEmailResponseData`

NewGetConnectionEmailResponseDataWithDefaults instantiates a new GetConnectionEmailResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetConnectionEmailResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetConnectionEmailResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetConnectionEmailResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetType

`func (o *GetConnectionEmailResponseData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetConnectionEmailResponseData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetConnectionEmailResponseData) SetType(v string)`

SetType sets Type field to given value.


### GetIsConnected

`func (o *GetConnectionEmailResponseData) GetIsConnected() bool`

GetIsConnected returns the IsConnected field if non-nil, zero value otherwise.

### GetIsConnectedOk

`func (o *GetConnectionEmailResponseData) GetIsConnectedOk() (*bool, bool)`

GetIsConnectedOk returns a tuple with the IsConnected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsConnected

`func (o *GetConnectionEmailResponseData) SetIsConnected(v bool)`

SetIsConnected sets IsConnected field to given value.


### GetEmail

`func (o *GetConnectionEmailResponseData) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetConnectionEmailResponseData) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetConnectionEmailResponseData) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetHost

`func (o *GetConnectionEmailResponseData) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetConnectionEmailResponseData) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetConnectionEmailResponseData) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *GetConnectionEmailResponseData) HasHost() bool`

HasHost returns a boolean if a field has been set.

### SetHostNil

`func (o *GetConnectionEmailResponseData) SetHostNil(b bool)`

 SetHostNil sets the value for Host to be an explicit nil

### UnsetHost
`func (o *GetConnectionEmailResponseData) UnsetHost()`

UnsetHost ensures that no value is present for Host, not even an explicit nil
### GetPort

`func (o *GetConnectionEmailResponseData) GetPort() float32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetConnectionEmailResponseData) GetPortOk() (*float32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetConnectionEmailResponseData) SetPort(v float32)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetConnectionEmailResponseData) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *GetConnectionEmailResponseData) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *GetConnectionEmailResponseData) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil
### GetLastSyncedAt

`func (o *GetConnectionEmailResponseData) GetLastSyncedAt() time.Time`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *GetConnectionEmailResponseData) GetLastSyncedAtOk() (*time.Time, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *GetConnectionEmailResponseData) SetLastSyncedAt(v time.Time)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *GetConnectionEmailResponseData) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### SetLastSyncedAtNil

`func (o *GetConnectionEmailResponseData) SetLastSyncedAtNil(b bool)`

 SetLastSyncedAtNil sets the value for LastSyncedAt to be an explicit nil

### UnsetLastSyncedAt
`func (o *GetConnectionEmailResponseData) UnsetLastSyncedAt()`

UnsetLastSyncedAt ensures that no value is present for LastSyncedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


