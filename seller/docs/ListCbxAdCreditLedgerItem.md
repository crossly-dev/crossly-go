# ListCbxAdCreditLedgerItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DeltaCents** | **float32** |  | 
**Kind** | **string** |  | 
**BaseUnitsPaid** | Pointer to **NullableString** |  | [optional] 
**CentsPerToken** | Pointer to **NullableFloat32** |  | [optional] 
**TxSig** | Pointer to **NullableString** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Memo** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **string** |  | 

## Methods

### NewListCbxAdCreditLedgerItem

`func NewListCbxAdCreditLedgerItem(id string, deltaCents float32, kind string, createdAt string, ) *ListCbxAdCreditLedgerItem`

NewListCbxAdCreditLedgerItem instantiates a new ListCbxAdCreditLedgerItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCbxAdCreditLedgerItemWithDefaults

`func NewListCbxAdCreditLedgerItemWithDefaults() *ListCbxAdCreditLedgerItem`

NewListCbxAdCreditLedgerItemWithDefaults instantiates a new ListCbxAdCreditLedgerItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCbxAdCreditLedgerItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCbxAdCreditLedgerItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCbxAdCreditLedgerItem) SetId(v string)`

SetId sets Id field to given value.


### GetDeltaCents

`func (o *ListCbxAdCreditLedgerItem) GetDeltaCents() float32`

GetDeltaCents returns the DeltaCents field if non-nil, zero value otherwise.

### GetDeltaCentsOk

`func (o *ListCbxAdCreditLedgerItem) GetDeltaCentsOk() (*float32, bool)`

GetDeltaCentsOk returns a tuple with the DeltaCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaCents

`func (o *ListCbxAdCreditLedgerItem) SetDeltaCents(v float32)`

SetDeltaCents sets DeltaCents field to given value.


### GetKind

`func (o *ListCbxAdCreditLedgerItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListCbxAdCreditLedgerItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListCbxAdCreditLedgerItem) SetKind(v string)`

SetKind sets Kind field to given value.


### GetBaseUnitsPaid

`func (o *ListCbxAdCreditLedgerItem) GetBaseUnitsPaid() string`

GetBaseUnitsPaid returns the BaseUnitsPaid field if non-nil, zero value otherwise.

### GetBaseUnitsPaidOk

`func (o *ListCbxAdCreditLedgerItem) GetBaseUnitsPaidOk() (*string, bool)`

GetBaseUnitsPaidOk returns a tuple with the BaseUnitsPaid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUnitsPaid

`func (o *ListCbxAdCreditLedgerItem) SetBaseUnitsPaid(v string)`

SetBaseUnitsPaid sets BaseUnitsPaid field to given value.

### HasBaseUnitsPaid

`func (o *ListCbxAdCreditLedgerItem) HasBaseUnitsPaid() bool`

HasBaseUnitsPaid returns a boolean if a field has been set.

### SetBaseUnitsPaidNil

`func (o *ListCbxAdCreditLedgerItem) SetBaseUnitsPaidNil(b bool)`

 SetBaseUnitsPaidNil sets the value for BaseUnitsPaid to be an explicit nil

### UnsetBaseUnitsPaid
`func (o *ListCbxAdCreditLedgerItem) UnsetBaseUnitsPaid()`

UnsetBaseUnitsPaid ensures that no value is present for BaseUnitsPaid, not even an explicit nil
### GetCentsPerToken

`func (o *ListCbxAdCreditLedgerItem) GetCentsPerToken() float32`

GetCentsPerToken returns the CentsPerToken field if non-nil, zero value otherwise.

### GetCentsPerTokenOk

`func (o *ListCbxAdCreditLedgerItem) GetCentsPerTokenOk() (*float32, bool)`

GetCentsPerTokenOk returns a tuple with the CentsPerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentsPerToken

`func (o *ListCbxAdCreditLedgerItem) SetCentsPerToken(v float32)`

SetCentsPerToken sets CentsPerToken field to given value.

### HasCentsPerToken

`func (o *ListCbxAdCreditLedgerItem) HasCentsPerToken() bool`

HasCentsPerToken returns a boolean if a field has been set.

### SetCentsPerTokenNil

`func (o *ListCbxAdCreditLedgerItem) SetCentsPerTokenNil(b bool)`

 SetCentsPerTokenNil sets the value for CentsPerToken to be an explicit nil

### UnsetCentsPerToken
`func (o *ListCbxAdCreditLedgerItem) UnsetCentsPerToken()`

UnsetCentsPerToken ensures that no value is present for CentsPerToken, not even an explicit nil
### GetTxSig

`func (o *ListCbxAdCreditLedgerItem) GetTxSig() string`

GetTxSig returns the TxSig field if non-nil, zero value otherwise.

### GetTxSigOk

`func (o *ListCbxAdCreditLedgerItem) GetTxSigOk() (*string, bool)`

GetTxSigOk returns a tuple with the TxSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxSig

`func (o *ListCbxAdCreditLedgerItem) SetTxSig(v string)`

SetTxSig sets TxSig field to given value.

### HasTxSig

`func (o *ListCbxAdCreditLedgerItem) HasTxSig() bool`

HasTxSig returns a boolean if a field has been set.

### SetTxSigNil

`func (o *ListCbxAdCreditLedgerItem) SetTxSigNil(b bool)`

 SetTxSigNil sets the value for TxSig to be an explicit nil

### UnsetTxSig
`func (o *ListCbxAdCreditLedgerItem) UnsetTxSig()`

UnsetTxSig ensures that no value is present for TxSig, not even an explicit nil
### GetExternalId

`func (o *ListCbxAdCreditLedgerItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListCbxAdCreditLedgerItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListCbxAdCreditLedgerItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListCbxAdCreditLedgerItem) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *ListCbxAdCreditLedgerItem) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *ListCbxAdCreditLedgerItem) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetMemo

`func (o *ListCbxAdCreditLedgerItem) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ListCbxAdCreditLedgerItem) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ListCbxAdCreditLedgerItem) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *ListCbxAdCreditLedgerItem) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *ListCbxAdCreditLedgerItem) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *ListCbxAdCreditLedgerItem) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetCreatedAt

`func (o *ListCbxAdCreditLedgerItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListCbxAdCreditLedgerItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListCbxAdCreditLedgerItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


