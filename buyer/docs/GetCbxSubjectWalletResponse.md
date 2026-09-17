# GetCbxSubjectWalletResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** |  | 
**Address** | **string** |  | 
**Chain** | **string** |  | 
**VerifiedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewGetCbxSubjectWalletResponse

`func NewGetCbxSubjectWalletResponse(walletId string, address string, chain string, ) *GetCbxSubjectWalletResponse`

NewGetCbxSubjectWalletResponse instantiates a new GetCbxSubjectWalletResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCbxSubjectWalletResponseWithDefaults

`func NewGetCbxSubjectWalletResponseWithDefaults() *GetCbxSubjectWalletResponse`

NewGetCbxSubjectWalletResponseWithDefaults instantiates a new GetCbxSubjectWalletResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *GetCbxSubjectWalletResponse) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *GetCbxSubjectWalletResponse) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *GetCbxSubjectWalletResponse) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetAddress

`func (o *GetCbxSubjectWalletResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetCbxSubjectWalletResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetCbxSubjectWalletResponse) SetAddress(v string)`

SetAddress sets Address field to given value.


### GetChain

`func (o *GetCbxSubjectWalletResponse) GetChain() string`

GetChain returns the Chain field if non-nil, zero value otherwise.

### GetChainOk

`func (o *GetCbxSubjectWalletResponse) GetChainOk() (*string, bool)`

GetChainOk returns a tuple with the Chain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChain

`func (o *GetCbxSubjectWalletResponse) SetChain(v string)`

SetChain sets Chain field to given value.


### GetVerifiedAt

`func (o *GetCbxSubjectWalletResponse) GetVerifiedAt() time.Time`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *GetCbxSubjectWalletResponse) GetVerifiedAtOk() (*time.Time, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *GetCbxSubjectWalletResponse) SetVerifiedAt(v time.Time)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *GetCbxSubjectWalletResponse) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### SetVerifiedAtNil

`func (o *GetCbxSubjectWalletResponse) SetVerifiedAtNil(b bool)`

 SetVerifiedAtNil sets the value for VerifiedAt to be an explicit nil

### UnsetVerifiedAt
`func (o *GetCbxSubjectWalletResponse) UnsetVerifiedAt()`

UnsetVerifiedAt ensures that no value is present for VerifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


