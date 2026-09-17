# CreateCbxWalletPaymentConfirmResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PaymentId** | **string** |  | 
**Duplicate** | **bool** |  | 
**ReleaseDecision** | **string** |  | 
**ReleaseReason** | **string** |  | 
**BaseUnits** | **string** |  | 
**ValueCents** | **float32** |  | 
**CentsPerToken** | **float32** |  | 
**PayerAddress** | **string** |  | 
**RiskLevel** | Pointer to **NullableString** |  | [optional] 
**RiskExposures** | **[]string** |  | 
**ScreeningProvider** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateCbxWalletPaymentConfirmResponse

`func NewCreateCbxWalletPaymentConfirmResponse(paymentId string, duplicate bool, releaseDecision string, releaseReason string, baseUnits string, valueCents float32, centsPerToken float32, payerAddress string, riskExposures []string, ) *CreateCbxWalletPaymentConfirmResponse`

NewCreateCbxWalletPaymentConfirmResponse instantiates a new CreateCbxWalletPaymentConfirmResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCbxWalletPaymentConfirmResponseWithDefaults

`func NewCreateCbxWalletPaymentConfirmResponseWithDefaults() *CreateCbxWalletPaymentConfirmResponse`

NewCreateCbxWalletPaymentConfirmResponseWithDefaults instantiates a new CreateCbxWalletPaymentConfirmResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPaymentId

`func (o *CreateCbxWalletPaymentConfirmResponse) GetPaymentId() string`

GetPaymentId returns the PaymentId field if non-nil, zero value otherwise.

### GetPaymentIdOk

`func (o *CreateCbxWalletPaymentConfirmResponse) GetPaymentIdOk() (*string, bool)`

GetPaymentIdOk returns a tuple with the PaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentId

`func (o *CreateCbxWalletPaymentConfirmResponse) SetPaymentId(v string)`

SetPaymentId sets PaymentId field to given value.


### GetDuplicate

`func (o *CreateCbxWalletPaymentConfirmResponse) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *CreateCbxWalletPaymentConfirmResponse) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *CreateCbxWalletPaymentConfirmResponse) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.


### GetReleaseDecision

`func (o *CreateCbxWalletPaymentConfirmResponse) GetReleaseDecision() string`

GetReleaseDecision returns the ReleaseDecision field if non-nil, zero value otherwise.

### GetReleaseDecisionOk

`func (o *CreateCbxWalletPaymentConfirmResponse) GetReleaseDecisionOk() (*string, bool)`

GetReleaseDecisionOk returns a tuple with the ReleaseDecision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDecision

`func (o *CreateCbxWalletPaymentConfirmResponse) SetReleaseDecision(v string)`

SetReleaseDecision sets ReleaseDecision field to given value.


### GetReleaseReason

`func (o *CreateCbxWalletPaymentConfirmResponse) GetReleaseReason() string`

GetReleaseReason returns the ReleaseReason field if non-nil, zero value otherwise.

### GetReleaseReasonOk

`func (o *CreateCbxWalletPaymentConfirmResponse) GetReleaseReasonOk() (*string, bool)`

GetReleaseReasonOk returns a tuple with the ReleaseReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseReason

`func (o *CreateCbxWalletPaymentConfirmResponse) SetReleaseReason(v string)`

SetReleaseReason sets ReleaseReason field to given value.


### GetBaseUnits

`func (o *CreateCbxWalletPaymentConfirmResponse) GetBaseUnits() string`

GetBaseUnits returns the BaseUnits field if non-nil, zero value otherwise.

### GetBaseUnitsOk

`func (o *CreateCbxWalletPaymentConfirmResponse) GetBaseUnitsOk() (*string, bool)`

GetBaseUnitsOk returns a tuple with the BaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUnits

`func (o *CreateCbxWalletPaymentConfirmResponse) SetBaseUnits(v string)`

SetBaseUnits sets BaseUnits field to given value.


### GetValueCents

`func (o *CreateCbxWalletPaymentConfirmResponse) GetValueCents() float32`

GetValueCents returns the ValueCents field if non-nil, zero value otherwise.

### GetValueCentsOk

`func (o *CreateCbxWalletPaymentConfirmResponse) GetValueCentsOk() (*float32, bool)`

GetValueCentsOk returns a tuple with the ValueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCents

`func (o *CreateCbxWalletPaymentConfirmResponse) SetValueCents(v float32)`

SetValueCents sets ValueCents field to given value.


### GetCentsPerToken

`func (o *CreateCbxWalletPaymentConfirmResponse) GetCentsPerToken() float32`

GetCentsPerToken returns the CentsPerToken field if non-nil, zero value otherwise.

### GetCentsPerTokenOk

`func (o *CreateCbxWalletPaymentConfirmResponse) GetCentsPerTokenOk() (*float32, bool)`

GetCentsPerTokenOk returns a tuple with the CentsPerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentsPerToken

`func (o *CreateCbxWalletPaymentConfirmResponse) SetCentsPerToken(v float32)`

SetCentsPerToken sets CentsPerToken field to given value.


### GetPayerAddress

`func (o *CreateCbxWalletPaymentConfirmResponse) GetPayerAddress() string`

GetPayerAddress returns the PayerAddress field if non-nil, zero value otherwise.

### GetPayerAddressOk

`func (o *CreateCbxWalletPaymentConfirmResponse) GetPayerAddressOk() (*string, bool)`

GetPayerAddressOk returns a tuple with the PayerAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayerAddress

`func (o *CreateCbxWalletPaymentConfirmResponse) SetPayerAddress(v string)`

SetPayerAddress sets PayerAddress field to given value.


### GetRiskLevel

`func (o *CreateCbxWalletPaymentConfirmResponse) GetRiskLevel() string`

GetRiskLevel returns the RiskLevel field if non-nil, zero value otherwise.

### GetRiskLevelOk

`func (o *CreateCbxWalletPaymentConfirmResponse) GetRiskLevelOk() (*string, bool)`

GetRiskLevelOk returns a tuple with the RiskLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskLevel

`func (o *CreateCbxWalletPaymentConfirmResponse) SetRiskLevel(v string)`

SetRiskLevel sets RiskLevel field to given value.

### HasRiskLevel

`func (o *CreateCbxWalletPaymentConfirmResponse) HasRiskLevel() bool`

HasRiskLevel returns a boolean if a field has been set.

### SetRiskLevelNil

`func (o *CreateCbxWalletPaymentConfirmResponse) SetRiskLevelNil(b bool)`

 SetRiskLevelNil sets the value for RiskLevel to be an explicit nil

### UnsetRiskLevel
`func (o *CreateCbxWalletPaymentConfirmResponse) UnsetRiskLevel()`

UnsetRiskLevel ensures that no value is present for RiskLevel, not even an explicit nil
### GetRiskExposures

`func (o *CreateCbxWalletPaymentConfirmResponse) GetRiskExposures() []string`

GetRiskExposures returns the RiskExposures field if non-nil, zero value otherwise.

### GetRiskExposuresOk

`func (o *CreateCbxWalletPaymentConfirmResponse) GetRiskExposuresOk() (*[]string, bool)`

GetRiskExposuresOk returns a tuple with the RiskExposures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRiskExposures

`func (o *CreateCbxWalletPaymentConfirmResponse) SetRiskExposures(v []string)`

SetRiskExposures sets RiskExposures field to given value.


### GetScreeningProvider

`func (o *CreateCbxWalletPaymentConfirmResponse) GetScreeningProvider() string`

GetScreeningProvider returns the ScreeningProvider field if non-nil, zero value otherwise.

### GetScreeningProviderOk

`func (o *CreateCbxWalletPaymentConfirmResponse) GetScreeningProviderOk() (*string, bool)`

GetScreeningProviderOk returns a tuple with the ScreeningProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScreeningProvider

`func (o *CreateCbxWalletPaymentConfirmResponse) SetScreeningProvider(v string)`

SetScreeningProvider sets ScreeningProvider field to given value.

### HasScreeningProvider

`func (o *CreateCbxWalletPaymentConfirmResponse) HasScreeningProvider() bool`

HasScreeningProvider returns a boolean if a field has been set.

### SetScreeningProviderNil

`func (o *CreateCbxWalletPaymentConfirmResponse) SetScreeningProviderNil(b bool)`

 SetScreeningProviderNil sets the value for ScreeningProvider to be an explicit nil

### UnsetScreeningProvider
`func (o *CreateCbxWalletPaymentConfirmResponse) UnsetScreeningProvider()`

UnsetScreeningProvider ensures that no value is present for ScreeningProvider, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


