# GetCbxClaimResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClaimId** | **string** |  | 
**Status** | **string** |  | 
**RequestedBaseUnits** | **string** |  | 
**NetBaseUnits** | **string** |  | 
**TxSig** | Pointer to **NullableString** |  | [optional] 
**FailureReason** | Pointer to **NullableString** |  | [optional] 
**Attempts** | **float32** |  | 
**CreatedAt** | **time.Time** |  | 
**ConfirmedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewGetCbxClaimResponse

`func NewGetCbxClaimResponse(claimId string, status string, requestedBaseUnits string, netBaseUnits string, attempts float32, createdAt time.Time, ) *GetCbxClaimResponse`

NewGetCbxClaimResponse instantiates a new GetCbxClaimResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCbxClaimResponseWithDefaults

`func NewGetCbxClaimResponseWithDefaults() *GetCbxClaimResponse`

NewGetCbxClaimResponseWithDefaults instantiates a new GetCbxClaimResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClaimId

`func (o *GetCbxClaimResponse) GetClaimId() string`

GetClaimId returns the ClaimId field if non-nil, zero value otherwise.

### GetClaimIdOk

`func (o *GetCbxClaimResponse) GetClaimIdOk() (*string, bool)`

GetClaimIdOk returns a tuple with the ClaimId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClaimId

`func (o *GetCbxClaimResponse) SetClaimId(v string)`

SetClaimId sets ClaimId field to given value.


### GetStatus

`func (o *GetCbxClaimResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCbxClaimResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCbxClaimResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRequestedBaseUnits

`func (o *GetCbxClaimResponse) GetRequestedBaseUnits() string`

GetRequestedBaseUnits returns the RequestedBaseUnits field if non-nil, zero value otherwise.

### GetRequestedBaseUnitsOk

`func (o *GetCbxClaimResponse) GetRequestedBaseUnitsOk() (*string, bool)`

GetRequestedBaseUnitsOk returns a tuple with the RequestedBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBaseUnits

`func (o *GetCbxClaimResponse) SetRequestedBaseUnits(v string)`

SetRequestedBaseUnits sets RequestedBaseUnits field to given value.


### GetNetBaseUnits

`func (o *GetCbxClaimResponse) GetNetBaseUnits() string`

GetNetBaseUnits returns the NetBaseUnits field if non-nil, zero value otherwise.

### GetNetBaseUnitsOk

`func (o *GetCbxClaimResponse) GetNetBaseUnitsOk() (*string, bool)`

GetNetBaseUnitsOk returns a tuple with the NetBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetBaseUnits

`func (o *GetCbxClaimResponse) SetNetBaseUnits(v string)`

SetNetBaseUnits sets NetBaseUnits field to given value.


### GetTxSig

`func (o *GetCbxClaimResponse) GetTxSig() string`

GetTxSig returns the TxSig field if non-nil, zero value otherwise.

### GetTxSigOk

`func (o *GetCbxClaimResponse) GetTxSigOk() (*string, bool)`

GetTxSigOk returns a tuple with the TxSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxSig

`func (o *GetCbxClaimResponse) SetTxSig(v string)`

SetTxSig sets TxSig field to given value.

### HasTxSig

`func (o *GetCbxClaimResponse) HasTxSig() bool`

HasTxSig returns a boolean if a field has been set.

### SetTxSigNil

`func (o *GetCbxClaimResponse) SetTxSigNil(b bool)`

 SetTxSigNil sets the value for TxSig to be an explicit nil

### UnsetTxSig
`func (o *GetCbxClaimResponse) UnsetTxSig()`

UnsetTxSig ensures that no value is present for TxSig, not even an explicit nil
### GetFailureReason

`func (o *GetCbxClaimResponse) GetFailureReason() string`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *GetCbxClaimResponse) GetFailureReasonOk() (*string, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *GetCbxClaimResponse) SetFailureReason(v string)`

SetFailureReason sets FailureReason field to given value.

### HasFailureReason

`func (o *GetCbxClaimResponse) HasFailureReason() bool`

HasFailureReason returns a boolean if a field has been set.

### SetFailureReasonNil

`func (o *GetCbxClaimResponse) SetFailureReasonNil(b bool)`

 SetFailureReasonNil sets the value for FailureReason to be an explicit nil

### UnsetFailureReason
`func (o *GetCbxClaimResponse) UnsetFailureReason()`

UnsetFailureReason ensures that no value is present for FailureReason, not even an explicit nil
### GetAttempts

`func (o *GetCbxClaimResponse) GetAttempts() float32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *GetCbxClaimResponse) GetAttemptsOk() (*float32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *GetCbxClaimResponse) SetAttempts(v float32)`

SetAttempts sets Attempts field to given value.


### GetCreatedAt

`func (o *GetCbxClaimResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetCbxClaimResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetCbxClaimResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetConfirmedAt

`func (o *GetCbxClaimResponse) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *GetCbxClaimResponse) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *GetCbxClaimResponse) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *GetCbxClaimResponse) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### SetConfirmedAtNil

`func (o *GetCbxClaimResponse) SetConfirmedAtNil(b bool)`

 SetConfirmedAtNil sets the value for ConfirmedAt to be an explicit nil

### UnsetConfirmedAt
`func (o *GetCbxClaimResponse) UnsetConfirmedAt()`

UnsetConfirmedAt ensures that no value is present for ConfirmedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


