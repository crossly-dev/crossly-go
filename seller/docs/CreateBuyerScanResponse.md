# CreateBuyerScanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verdict** | **string** |  | 
**Crossly** | Pointer to [**NullableGetBuyerAnywhereResponseCrossly**](GetBuyerAnywhereResponseCrossly.md) |  | [optional] 
**Offsite** | Pointer to [**NullableGetBuyerAnywhereResponseOffsite**](GetBuyerAnywhereResponseOffsite.md) |  | [optional] 
**Alternates** | [**[]GetBuyerAnywhereResponseAlternates**](GetBuyerAnywhereResponseAlternates.md) |  | 
**SavingCents** | Pointer to **NullableFloat32** |  | [optional] 
**ShippingUnknown** | **bool** |  | 
**MatchMethod** | **string** |  | 
**Comparable** | **bool** |  | 
**Confidence** | **float32** |  | 
**Identifier** | [**CreateBuyerScanResponseIdentifier**](CreateBuyerScanResponseIdentifier.md) |  | 
**VisualMatches** | **[]interface{}** |  | 

## Methods

### NewCreateBuyerScanResponse

`func NewCreateBuyerScanResponse(verdict string, alternates []GetBuyerAnywhereResponseAlternates, shippingUnknown bool, matchMethod string, comparable bool, confidence float32, identifier CreateBuyerScanResponseIdentifier, visualMatches []interface{}, ) *CreateBuyerScanResponse`

NewCreateBuyerScanResponse instantiates a new CreateBuyerScanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerScanResponseWithDefaults

`func NewCreateBuyerScanResponseWithDefaults() *CreateBuyerScanResponse`

NewCreateBuyerScanResponseWithDefaults instantiates a new CreateBuyerScanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerdict

`func (o *CreateBuyerScanResponse) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *CreateBuyerScanResponse) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *CreateBuyerScanResponse) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.


### GetCrossly

`func (o *CreateBuyerScanResponse) GetCrossly() GetBuyerAnywhereResponseCrossly`

GetCrossly returns the Crossly field if non-nil, zero value otherwise.

### GetCrosslyOk

`func (o *CreateBuyerScanResponse) GetCrosslyOk() (*GetBuyerAnywhereResponseCrossly, bool)`

GetCrosslyOk returns a tuple with the Crossly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossly

`func (o *CreateBuyerScanResponse) SetCrossly(v GetBuyerAnywhereResponseCrossly)`

SetCrossly sets Crossly field to given value.

### HasCrossly

`func (o *CreateBuyerScanResponse) HasCrossly() bool`

HasCrossly returns a boolean if a field has been set.

### SetCrosslyNil

`func (o *CreateBuyerScanResponse) SetCrosslyNil(b bool)`

 SetCrosslyNil sets the value for Crossly to be an explicit nil

### UnsetCrossly
`func (o *CreateBuyerScanResponse) UnsetCrossly()`

UnsetCrossly ensures that no value is present for Crossly, not even an explicit nil
### GetOffsite

`func (o *CreateBuyerScanResponse) GetOffsite() GetBuyerAnywhereResponseOffsite`

GetOffsite returns the Offsite field if non-nil, zero value otherwise.

### GetOffsiteOk

`func (o *CreateBuyerScanResponse) GetOffsiteOk() (*GetBuyerAnywhereResponseOffsite, bool)`

GetOffsiteOk returns a tuple with the Offsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsite

`func (o *CreateBuyerScanResponse) SetOffsite(v GetBuyerAnywhereResponseOffsite)`

SetOffsite sets Offsite field to given value.

### HasOffsite

`func (o *CreateBuyerScanResponse) HasOffsite() bool`

HasOffsite returns a boolean if a field has been set.

### SetOffsiteNil

`func (o *CreateBuyerScanResponse) SetOffsiteNil(b bool)`

 SetOffsiteNil sets the value for Offsite to be an explicit nil

### UnsetOffsite
`func (o *CreateBuyerScanResponse) UnsetOffsite()`

UnsetOffsite ensures that no value is present for Offsite, not even an explicit nil
### GetAlternates

`func (o *CreateBuyerScanResponse) GetAlternates() []GetBuyerAnywhereResponseAlternates`

GetAlternates returns the Alternates field if non-nil, zero value otherwise.

### GetAlternatesOk

`func (o *CreateBuyerScanResponse) GetAlternatesOk() (*[]GetBuyerAnywhereResponseAlternates, bool)`

GetAlternatesOk returns a tuple with the Alternates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternates

`func (o *CreateBuyerScanResponse) SetAlternates(v []GetBuyerAnywhereResponseAlternates)`

SetAlternates sets Alternates field to given value.


### GetSavingCents

`func (o *CreateBuyerScanResponse) GetSavingCents() float32`

GetSavingCents returns the SavingCents field if non-nil, zero value otherwise.

### GetSavingCentsOk

`func (o *CreateBuyerScanResponse) GetSavingCentsOk() (*float32, bool)`

GetSavingCentsOk returns a tuple with the SavingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingCents

`func (o *CreateBuyerScanResponse) SetSavingCents(v float32)`

SetSavingCents sets SavingCents field to given value.

### HasSavingCents

`func (o *CreateBuyerScanResponse) HasSavingCents() bool`

HasSavingCents returns a boolean if a field has been set.

### SetSavingCentsNil

`func (o *CreateBuyerScanResponse) SetSavingCentsNil(b bool)`

 SetSavingCentsNil sets the value for SavingCents to be an explicit nil

### UnsetSavingCents
`func (o *CreateBuyerScanResponse) UnsetSavingCents()`

UnsetSavingCents ensures that no value is present for SavingCents, not even an explicit nil
### GetShippingUnknown

`func (o *CreateBuyerScanResponse) GetShippingUnknown() bool`

GetShippingUnknown returns the ShippingUnknown field if non-nil, zero value otherwise.

### GetShippingUnknownOk

`func (o *CreateBuyerScanResponse) GetShippingUnknownOk() (*bool, bool)`

GetShippingUnknownOk returns a tuple with the ShippingUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingUnknown

`func (o *CreateBuyerScanResponse) SetShippingUnknown(v bool)`

SetShippingUnknown sets ShippingUnknown field to given value.


### GetMatchMethod

`func (o *CreateBuyerScanResponse) GetMatchMethod() string`

GetMatchMethod returns the MatchMethod field if non-nil, zero value otherwise.

### GetMatchMethodOk

`func (o *CreateBuyerScanResponse) GetMatchMethodOk() (*string, bool)`

GetMatchMethodOk returns a tuple with the MatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMethod

`func (o *CreateBuyerScanResponse) SetMatchMethod(v string)`

SetMatchMethod sets MatchMethod field to given value.


### GetComparable

`func (o *CreateBuyerScanResponse) GetComparable() bool`

GetComparable returns the Comparable field if non-nil, zero value otherwise.

### GetComparableOk

`func (o *CreateBuyerScanResponse) GetComparableOk() (*bool, bool)`

GetComparableOk returns a tuple with the Comparable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComparable

`func (o *CreateBuyerScanResponse) SetComparable(v bool)`

SetComparable sets Comparable field to given value.


### GetConfidence

`func (o *CreateBuyerScanResponse) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CreateBuyerScanResponse) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CreateBuyerScanResponse) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.


### GetIdentifier

`func (o *CreateBuyerScanResponse) GetIdentifier() CreateBuyerScanResponseIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateBuyerScanResponse) GetIdentifierOk() (*CreateBuyerScanResponseIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateBuyerScanResponse) SetIdentifier(v CreateBuyerScanResponseIdentifier)`

SetIdentifier sets Identifier field to given value.


### GetVisualMatches

`func (o *CreateBuyerScanResponse) GetVisualMatches() []interface{}`

GetVisualMatches returns the VisualMatches field if non-nil, zero value otherwise.

### GetVisualMatchesOk

`func (o *CreateBuyerScanResponse) GetVisualMatchesOk() (*[]interface{}, bool)`

GetVisualMatchesOk returns a tuple with the VisualMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualMatches

`func (o *CreateBuyerScanResponse) SetVisualMatches(v []interface{})`

SetVisualMatches sets VisualMatches field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


