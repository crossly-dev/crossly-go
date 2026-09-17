# CreateAccountResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Platform** | **string** |  | 
**AccountSlot** | **float32** |  | 
**ProxyAssigned** | **bool** |  | 

## Methods

### NewCreateAccountResponse

`func NewCreateAccountResponse(id string, platform string, accountSlot float32, proxyAssigned bool, ) *CreateAccountResponse`

NewCreateAccountResponse instantiates a new CreateAccountResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountResponseWithDefaults

`func NewCreateAccountResponseWithDefaults() *CreateAccountResponse`

NewCreateAccountResponseWithDefaults instantiates a new CreateAccountResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateAccountResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAccountResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAccountResponse) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *CreateAccountResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateAccountResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateAccountResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetAccountSlot

`func (o *CreateAccountResponse) GetAccountSlot() float32`

GetAccountSlot returns the AccountSlot field if non-nil, zero value otherwise.

### GetAccountSlotOk

`func (o *CreateAccountResponse) GetAccountSlotOk() (*float32, bool)`

GetAccountSlotOk returns a tuple with the AccountSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSlot

`func (o *CreateAccountResponse) SetAccountSlot(v float32)`

SetAccountSlot sets AccountSlot field to given value.


### GetProxyAssigned

`func (o *CreateAccountResponse) GetProxyAssigned() bool`

GetProxyAssigned returns the ProxyAssigned field if non-nil, zero value otherwise.

### GetProxyAssignedOk

`func (o *CreateAccountResponse) GetProxyAssignedOk() (*bool, bool)`

GetProxyAssignedOk returns a tuple with the ProxyAssigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyAssigned

`func (o *CreateAccountResponse) SetProxyAssigned(v bool)`

SetProxyAssigned sets ProxyAssigned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


