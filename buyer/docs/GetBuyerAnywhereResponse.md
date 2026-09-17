# GetBuyerAnywhereResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verdict** | **string** |  | 
**SavingCents** | Pointer to **NullableFloat32** |  | [optional] 
**ShippingUnknown** | **bool** |  | 
**Crossly** | Pointer to [**NullableGetBuyerAnywhereResponseCrossly**](GetBuyerAnywhereResponseCrossly.md) |  | [optional] 
**Offsite** | Pointer to [**NullableGetBuyerAnywhereResponseOffsite**](GetBuyerAnywhereResponseOffsite.md) |  | [optional] 
**Alternates** | [**[]GetBuyerAnywhereResponseAlternates**](GetBuyerAnywhereResponseAlternates.md) |  | 

## Methods

### NewGetBuyerAnywhereResponse

`func NewGetBuyerAnywhereResponse(verdict string, shippingUnknown bool, alternates []GetBuyerAnywhereResponseAlternates, ) *GetBuyerAnywhereResponse`

NewGetBuyerAnywhereResponse instantiates a new GetBuyerAnywhereResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBuyerAnywhereResponseWithDefaults

`func NewGetBuyerAnywhereResponseWithDefaults() *GetBuyerAnywhereResponse`

NewGetBuyerAnywhereResponseWithDefaults instantiates a new GetBuyerAnywhereResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerdict

`func (o *GetBuyerAnywhereResponse) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *GetBuyerAnywhereResponse) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *GetBuyerAnywhereResponse) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.


### GetSavingCents

`func (o *GetBuyerAnywhereResponse) GetSavingCents() float32`

GetSavingCents returns the SavingCents field if non-nil, zero value otherwise.

### GetSavingCentsOk

`func (o *GetBuyerAnywhereResponse) GetSavingCentsOk() (*float32, bool)`

GetSavingCentsOk returns a tuple with the SavingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingCents

`func (o *GetBuyerAnywhereResponse) SetSavingCents(v float32)`

SetSavingCents sets SavingCents field to given value.

### HasSavingCents

`func (o *GetBuyerAnywhereResponse) HasSavingCents() bool`

HasSavingCents returns a boolean if a field has been set.

### SetSavingCentsNil

`func (o *GetBuyerAnywhereResponse) SetSavingCentsNil(b bool)`

 SetSavingCentsNil sets the value for SavingCents to be an explicit nil

### UnsetSavingCents
`func (o *GetBuyerAnywhereResponse) UnsetSavingCents()`

UnsetSavingCents ensures that no value is present for SavingCents, not even an explicit nil
### GetShippingUnknown

`func (o *GetBuyerAnywhereResponse) GetShippingUnknown() bool`

GetShippingUnknown returns the ShippingUnknown field if non-nil, zero value otherwise.

### GetShippingUnknownOk

`func (o *GetBuyerAnywhereResponse) GetShippingUnknownOk() (*bool, bool)`

GetShippingUnknownOk returns a tuple with the ShippingUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingUnknown

`func (o *GetBuyerAnywhereResponse) SetShippingUnknown(v bool)`

SetShippingUnknown sets ShippingUnknown field to given value.


### GetCrossly

`func (o *GetBuyerAnywhereResponse) GetCrossly() GetBuyerAnywhereResponseCrossly`

GetCrossly returns the Crossly field if non-nil, zero value otherwise.

### GetCrosslyOk

`func (o *GetBuyerAnywhereResponse) GetCrosslyOk() (*GetBuyerAnywhereResponseCrossly, bool)`

GetCrosslyOk returns a tuple with the Crossly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossly

`func (o *GetBuyerAnywhereResponse) SetCrossly(v GetBuyerAnywhereResponseCrossly)`

SetCrossly sets Crossly field to given value.

### HasCrossly

`func (o *GetBuyerAnywhereResponse) HasCrossly() bool`

HasCrossly returns a boolean if a field has been set.

### SetCrosslyNil

`func (o *GetBuyerAnywhereResponse) SetCrosslyNil(b bool)`

 SetCrosslyNil sets the value for Crossly to be an explicit nil

### UnsetCrossly
`func (o *GetBuyerAnywhereResponse) UnsetCrossly()`

UnsetCrossly ensures that no value is present for Crossly, not even an explicit nil
### GetOffsite

`func (o *GetBuyerAnywhereResponse) GetOffsite() GetBuyerAnywhereResponseOffsite`

GetOffsite returns the Offsite field if non-nil, zero value otherwise.

### GetOffsiteOk

`func (o *GetBuyerAnywhereResponse) GetOffsiteOk() (*GetBuyerAnywhereResponseOffsite, bool)`

GetOffsiteOk returns a tuple with the Offsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsite

`func (o *GetBuyerAnywhereResponse) SetOffsite(v GetBuyerAnywhereResponseOffsite)`

SetOffsite sets Offsite field to given value.

### HasOffsite

`func (o *GetBuyerAnywhereResponse) HasOffsite() bool`

HasOffsite returns a boolean if a field has been set.

### SetOffsiteNil

`func (o *GetBuyerAnywhereResponse) SetOffsiteNil(b bool)`

 SetOffsiteNil sets the value for Offsite to be an explicit nil

### UnsetOffsite
`func (o *GetBuyerAnywhereResponse) UnsetOffsite()`

UnsetOffsite ensures that no value is present for Offsite, not even an explicit nil
### GetAlternates

`func (o *GetBuyerAnywhereResponse) GetAlternates() []GetBuyerAnywhereResponseAlternates`

GetAlternates returns the Alternates field if non-nil, zero value otherwise.

### GetAlternatesOk

`func (o *GetBuyerAnywhereResponse) GetAlternatesOk() (*[]GetBuyerAnywhereResponseAlternates, bool)`

GetAlternatesOk returns a tuple with the Alternates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternates

`func (o *GetBuyerAnywhereResponse) SetAlternates(v []GetBuyerAnywhereResponseAlternates)`

SetAlternates sets Alternates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


