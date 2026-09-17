# CreateCbxWalletChallengeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** |  | 
**Message** | **string** | The exact text to present for signing. | 
**Nonce** | **string** |  | 
**ExpiresInMs** | **float32** |  | 

## Methods

### NewCreateCbxWalletChallengeResponse

`func NewCreateCbxWalletChallengeResponse(walletId string, message string, nonce string, expiresInMs float32, ) *CreateCbxWalletChallengeResponse`

NewCreateCbxWalletChallengeResponse instantiates a new CreateCbxWalletChallengeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCbxWalletChallengeResponseWithDefaults

`func NewCreateCbxWalletChallengeResponseWithDefaults() *CreateCbxWalletChallengeResponse`

NewCreateCbxWalletChallengeResponseWithDefaults instantiates a new CreateCbxWalletChallengeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *CreateCbxWalletChallengeResponse) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *CreateCbxWalletChallengeResponse) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *CreateCbxWalletChallengeResponse) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetMessage

`func (o *CreateCbxWalletChallengeResponse) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateCbxWalletChallengeResponse) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateCbxWalletChallengeResponse) SetMessage(v string)`

SetMessage sets Message field to given value.


### GetNonce

`func (o *CreateCbxWalletChallengeResponse) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *CreateCbxWalletChallengeResponse) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *CreateCbxWalletChallengeResponse) SetNonce(v string)`

SetNonce sets Nonce field to given value.


### GetExpiresInMs

`func (o *CreateCbxWalletChallengeResponse) GetExpiresInMs() float32`

GetExpiresInMs returns the ExpiresInMs field if non-nil, zero value otherwise.

### GetExpiresInMsOk

`func (o *CreateCbxWalletChallengeResponse) GetExpiresInMsOk() (*float32, bool)`

GetExpiresInMsOk returns a tuple with the ExpiresInMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresInMs

`func (o *CreateCbxWalletChallengeResponse) SetExpiresInMs(v float32)`

SetExpiresInMs sets ExpiresInMs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


