# CreateCbxClaimSendResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | **string** |  | 
**TxSig** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateCbxClaimSendResponse

`func NewCreateCbxClaimSendResponse(status string, ) *CreateCbxClaimSendResponse`

NewCreateCbxClaimSendResponse instantiates a new CreateCbxClaimSendResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCbxClaimSendResponseWithDefaults

`func NewCreateCbxClaimSendResponseWithDefaults() *CreateCbxClaimSendResponse`

NewCreateCbxClaimSendResponseWithDefaults instantiates a new CreateCbxClaimSendResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CreateCbxClaimSendResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateCbxClaimSendResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateCbxClaimSendResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetTxSig

`func (o *CreateCbxClaimSendResponse) GetTxSig() string`

GetTxSig returns the TxSig field if non-nil, zero value otherwise.

### GetTxSigOk

`func (o *CreateCbxClaimSendResponse) GetTxSigOk() (*string, bool)`

GetTxSigOk returns a tuple with the TxSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxSig

`func (o *CreateCbxClaimSendResponse) SetTxSig(v string)`

SetTxSig sets TxSig field to given value.

### HasTxSig

`func (o *CreateCbxClaimSendResponse) HasTxSig() bool`

HasTxSig returns a boolean if a field has been set.

### SetTxSigNil

`func (o *CreateCbxClaimSendResponse) SetTxSigNil(b bool)`

 SetTxSigNil sets the value for TxSig to be an explicit nil

### UnsetTxSig
`func (o *CreateCbxClaimSendResponse) UnsetTxSig()`

UnsetTxSig ensures that no value is present for TxSig, not even an explicit nil
### GetReason

`func (o *CreateCbxClaimSendResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateCbxClaimSendResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateCbxClaimSendResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateCbxClaimSendResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *CreateCbxClaimSendResponse) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *CreateCbxClaimSendResponse) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


