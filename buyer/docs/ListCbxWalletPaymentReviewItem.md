# ListCbxWalletPaymentReviewItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ExternalId** | **string** |  | 
**PayerAddress** | **string** |  | 
**BaseUnits** | **string** |  | 
**ValueCents** | **float32** |  | 
**RiskLevel** | Pointer to **NullableString** |  | [optional] 
**RiskExposures** | **[]string** |  | 
**ScreeningProvider** | Pointer to **NullableString** |  | [optional] 
**ReleaseReason** | Pointer to **NullableString** |  | [optional] 
**TxSig** | **string** |  | 
**ConfirmedAt** | **string** |  | 

## Methods

### NewListCbxWalletPaymentReviewItem

`func NewListCbxWalletPaymentReviewItem(id string, externalId string, payerAddress string, baseUnits string, valueCents float32, riskExposures []string, txSig string, confirmedAt string, ) *ListCbxWalletPaymentReviewItem`

NewListCbxWalletPaymentReviewItem instantiates a new ListCbxWalletPaymentReviewItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCbxWalletPaymentReviewItemWithDefaults

`func NewListCbxWalletPaymentReviewItemWithDefaults() *ListCbxWalletPaymentReviewItem`

NewListCbxWalletPaymentReviewItemWithDefaults instantiates a new ListCbxWalletPaymentReviewItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCbxWalletPaymentReviewItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCbxWalletPaymentReviewItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCbxWalletPaymentReviewItem) SetId(v string)`

SetId sets Id field to given value.


### GetExternalId

`func (o *ListCbxWalletPaymentReviewItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListCbxWalletPaymentReviewItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListCbxWalletPaymentReviewItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetPayerAddress

`func (o *ListCbxWalletPaymentReviewItem) GetPayerAddress() string`

GetPayerAddress returns the PayerAddress field if non-nil, zero value otherwise.

### GetPayerAddressOk

`func (o *ListCbxWalletPaymentReviewItem) GetPayerAddressOk() (*string, bool)`

GetPayerAddressOk returns a tuple with the PayerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerAddress

`func (o *ListCbxWalletPaymentReviewItem) SetPayerAddress(v string)`

SetPayerAddress sets PayerAddress field to given value.


### GetBaseUnits

`func (o *ListCbxWalletPaymentReviewItem) GetBaseUnits() string`

GetBaseUnits returns the BaseUnits field if non-nil, zero value otherwise.

### GetBaseUnitsOk

`func (o *ListCbxWalletPaymentReviewItem) GetBaseUnitsOk() (*string, bool)`

GetBaseUnitsOk returns a tuple with the BaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUnits

`func (o *ListCbxWalletPaymentReviewItem) SetBaseUnits(v string)`

SetBaseUnits sets BaseUnits field to given value.


### GetValueCents

`func (o *ListCbxWalletPaymentReviewItem) GetValueCents() float32`

GetValueCents returns the ValueCents field if non-nil, zero value otherwise.

### GetValueCentsOk

`func (o *ListCbxWalletPaymentReviewItem) GetValueCentsOk() (*float32, bool)`

GetValueCentsOk returns a tuple with the ValueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCents

`func (o *ListCbxWalletPaymentReviewItem) SetValueCents(v float32)`

SetValueCents sets ValueCents field to given value.


### GetRiskLevel

`func (o *ListCbxWalletPaymentReviewItem) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *ListCbxWalletPaymentReviewItem) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *ListCbxWalletPaymentReviewItem) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *ListCbxWalletPaymentReviewItem) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.

### SetRiskLevelNil

`func (o *ListCbxWalletPaymentReviewItem) SetRiskLevelNil(b bool)`

 SetRiskLevelNil sets the value for RiskLevel to be an explicit nil

### UnsetRiskLevel
`func (o *ListCbxWalletPaymentReviewItem) UnsetRiskLevel()`

UnsetRiskLevel ensures that no value is present for RiskLevel, not even an explicit nil
### GetRiskExposures

`func (o *ListCbxWalletPaymentReviewItem) GetRiskExposures() []string`

GetRiskExposures returns the RiskExposures field if non-nil, zero value otherwise.

### GetRiskExposuresOk

`func (o *ListCbxWalletPaymentReviewItem) GetRiskExposuresOk() (*[]string, bool)`

GetRiskExposuresOk returns a tuple with the RiskExposures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskExposures

`func (o *ListCbxWalletPaymentReviewItem) SetRiskExposures(v []string)`

SetRiskExposures sets RiskExposures field to given value.


### GetScreeningProvider

`func (o *ListCbxWalletPaymentReviewItem) GetScreeningProvider() string`

GetScreeningProvider returns the ScreeningProvider field if non-nil, zero value otherwise.

### GetScreeningProviderOk

`func (o *ListCbxWalletPaymentReviewItem) GetScreeningProviderOk() (*string, bool)`

GetScreeningProviderOk returns a tuple with the ScreeningProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreeningProvider

`func (o *ListCbxWalletPaymentReviewItem) SetScreeningProvider(v string)`

SetScreeningProvider sets ScreeningProvider field to given value.

### HasScreeningProvider

`func (o *ListCbxWalletPaymentReviewItem) HasScreeningProvider() bool`

HasScreeningProvider returns a boolean if a field has been set.

### SetScreeningProviderNil

`func (o *ListCbxWalletPaymentReviewItem) SetScreeningProviderNil(b bool)`

 SetScreeningProviderNil sets the value for ScreeningProvider to be an explicit nil

### UnsetScreeningProvider
`func (o *ListCbxWalletPaymentReviewItem) UnsetScreeningProvider()`

UnsetScreeningProvider ensures that no value is present for ScreeningProvider, not even an explicit nil
### GetReleaseReason

`func (o *ListCbxWalletPaymentReviewItem) GetReleaseReason() string`

GetReleaseReason returns the ReleaseReason field if non-nil, zero value otherwise.

### GetReleaseReasonOk

`func (o *ListCbxWalletPaymentReviewItem) GetReleaseReasonOk() (*string, bool)`

GetReleaseReasonOk returns a tuple with the ReleaseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseReason

`func (o *ListCbxWalletPaymentReviewItem) SetReleaseReason(v string)`

SetReleaseReason sets ReleaseReason field to given value.

### HasReleaseReason

`func (o *ListCbxWalletPaymentReviewItem) HasReleaseReason() bool`

HasReleaseReason returns a boolean if a field has been set.

### SetReleaseReasonNil

`func (o *ListCbxWalletPaymentReviewItem) SetReleaseReasonNil(b bool)`

 SetReleaseReasonNil sets the value for ReleaseReason to be an explicit nil

### UnsetReleaseReason
`func (o *ListCbxWalletPaymentReviewItem) UnsetReleaseReason()`

UnsetReleaseReason ensures that no value is present for ReleaseReason, not even an explicit nil
### GetTxSig

`func (o *ListCbxWalletPaymentReviewItem) GetTxSig() string`

GetTxSig returns the TxSig field if non-nil, zero value otherwise.

### GetTxSigOk

`func (o *ListCbxWalletPaymentReviewItem) GetTxSigOk() (*string, bool)`

GetTxSigOk returns a tuple with the TxSig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTxSig

`func (o *ListCbxWalletPaymentReviewItem) SetTxSig(v string)`

SetTxSig sets TxSig field to given value.


### GetConfirmedAt

`func (o *ListCbxWalletPaymentReviewItem) GetConfirmedAt() string`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *ListCbxWalletPaymentReviewItem) GetConfirmedAtOk() (*string, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *ListCbxWalletPaymentReviewItem) SetConfirmedAt(v string)`

SetConfirmedAt sets ConfirmedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


