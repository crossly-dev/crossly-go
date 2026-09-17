# CreateCbxRevenueSweepResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Swept** | **bool** |  | 
**BaseUnits** | **string** |  | 
**TxSig** | Pointer to **NullableString** |  | [optional] 
**Reason** | Pointer to **NullableString** |  | [optional] 
**Breakdown** | Pointer to [**NullableCreateCbxRevenueSweepResponseBreakdown**](CreateCbxRevenueSweepResponseBreakdown.md) |  | [optional] 

## Methods

### NewCreateCbxRevenueSweepResponse

`func NewCreateCbxRevenueSweepResponse(swept bool, baseUnits string, ) *CreateCbxRevenueSweepResponse`

NewCreateCbxRevenueSweepResponse instantiates a new CreateCbxRevenueSweepResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCbxRevenueSweepResponseWithDefaults

`func NewCreateCbxRevenueSweepResponseWithDefaults() *CreateCbxRevenueSweepResponse`

NewCreateCbxRevenueSweepResponseWithDefaults instantiates a new CreateCbxRevenueSweepResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSwept

`func (o *CreateCbxRevenueSweepResponse) GetSwept() bool`

GetSwept returns the Swept field if non-nil, zero value otherwise.

### GetSweptOk

`func (o *CreateCbxRevenueSweepResponse) GetSweptOk() (*bool, bool)`

GetSweptOk returns a tuple with the Swept field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwept

`func (o *CreateCbxRevenueSweepResponse) SetSwept(v bool)`

SetSwept sets Swept field to given value.


### GetBaseUnits

`func (o *CreateCbxRevenueSweepResponse) GetBaseUnits() string`

GetBaseUnits returns the BaseUnits field if non-nil, zero value otherwise.

### GetBaseUnitsOk

`func (o *CreateCbxRevenueSweepResponse) GetBaseUnitsOk() (*string, bool)`

GetBaseUnitsOk returns a tuple with the BaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUnits

`func (o *CreateCbxRevenueSweepResponse) SetBaseUnits(v string)`

SetBaseUnits sets BaseUnits field to given value.


### GetTxSig

`func (o *CreateCbxRevenueSweepResponse) GetTxSig() string`

GetTxSig returns the TxSig field if non-nil, zero value otherwise.

### GetTxSigOk

`func (o *CreateCbxRevenueSweepResponse) GetTxSigOk() (*string, bool)`

GetTxSigOk returns a tuple with the TxSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxSig

`func (o *CreateCbxRevenueSweepResponse) SetTxSig(v string)`

SetTxSig sets TxSig field to given value.

### HasTxSig

`func (o *CreateCbxRevenueSweepResponse) HasTxSig() bool`

HasTxSig returns a boolean if a field has been set.

### SetTxSigNil

`func (o *CreateCbxRevenueSweepResponse) SetTxSigNil(b bool)`

 SetTxSigNil sets the value for TxSig to be an explicit nil

### UnsetTxSig
`func (o *CreateCbxRevenueSweepResponse) UnsetTxSig()`

UnsetTxSig ensures that no value is present for TxSig, not even an explicit nil
### GetReason

`func (o *CreateCbxRevenueSweepResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateCbxRevenueSweepResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateCbxRevenueSweepResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateCbxRevenueSweepResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *CreateCbxRevenueSweepResponse) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *CreateCbxRevenueSweepResponse) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetBreakdown

`func (o *CreateCbxRevenueSweepResponse) GetBreakdown() CreateCbxRevenueSweepResponseBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *CreateCbxRevenueSweepResponse) GetBreakdownOk() (*CreateCbxRevenueSweepResponseBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *CreateCbxRevenueSweepResponse) SetBreakdown(v CreateCbxRevenueSweepResponseBreakdown)`

SetBreakdown sets Breakdown field to given value.

### HasBreakdown

`func (o *CreateCbxRevenueSweepResponse) HasBreakdown() bool`

HasBreakdown returns a boolean if a field has been set.

### SetBreakdownNil

`func (o *CreateCbxRevenueSweepResponse) SetBreakdownNil(b bool)`

 SetBreakdownNil sets the value for Breakdown to be an explicit nil

### UnsetBreakdown
`func (o *CreateCbxRevenueSweepResponse) UnsetBreakdown()`

UnsetBreakdown ensures that no value is present for Breakdown, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


