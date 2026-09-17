# CreateBuyerIdentifyResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tier** | **string** |  | 
**Confidence** | **float32** |  | 
**Identifier** | Pointer to [**NullableCreateBuyerIdentifyResponseIdentifier**](CreateBuyerIdentifyResponseIdentifier.md) |  | [optional] 
**VisionLabel** | Pointer to **NullableString** |  | [optional] 
**VisionQuotaExhausted** | **bool** | Surfaced rather than hidden: \&quot;we could not look harder\&quot; and \&quot;we looked and found nothing\&quot; are different answers, and a client that cannot tell them apart shows the wrong message on both. | 
**VisualMatches** | [**[]CreateBuyerIdentifyResponseVisualMatches**](CreateBuyerIdentifyResponseVisualMatches.md) |  | 
**Verdict** | **string** |  | 
**Crossly** | Pointer to [**NullableGetBuyerAnywhereResponseCrossly**](GetBuyerAnywhereResponseCrossly.md) |  | [optional] 
**Offsite** | Pointer to [**NullableGetBuyerAnywhereResponseOffsite**](GetBuyerAnywhereResponseOffsite.md) |  | [optional] 
**Alternates** | [**[]GetBuyerAnywhereResponseAlternates**](GetBuyerAnywhereResponseAlternates.md) |  | 
**SavingCents** | Pointer to **NullableFloat32** |  | [optional] 
**ShippingUnknown** | **bool** |  | 
**Hud** | [**CreateBuyerIdentifyResponseHud**](CreateBuyerIdentifyResponseHud.md) |  | 

## Methods

### NewCreateBuyerIdentifyResponse

`func NewCreateBuyerIdentifyResponse(tier string, confidence float32, visionQuotaExhausted bool, visualMatches []CreateBuyerIdentifyResponseVisualMatches, verdict string, alternates []GetBuyerAnywhereResponseAlternates, shippingUnknown bool, hud CreateBuyerIdentifyResponseHud, ) *CreateBuyerIdentifyResponse`

NewCreateBuyerIdentifyResponse instantiates a new CreateBuyerIdentifyResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerIdentifyResponseWithDefaults

`func NewCreateBuyerIdentifyResponseWithDefaults() *CreateBuyerIdentifyResponse`

NewCreateBuyerIdentifyResponseWithDefaults instantiates a new CreateBuyerIdentifyResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTier

`func (o *CreateBuyerIdentifyResponse) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CreateBuyerIdentifyResponse) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CreateBuyerIdentifyResponse) SetTier(v string)`

SetTier sets Tier field to given value.


### GetConfidence

`func (o *CreateBuyerIdentifyResponse) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CreateBuyerIdentifyResponse) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CreateBuyerIdentifyResponse) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.


### GetIdentifier

`func (o *CreateBuyerIdentifyResponse) GetIdentifier() CreateBuyerIdentifyResponseIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateBuyerIdentifyResponse) GetIdentifierOk() (*CreateBuyerIdentifyResponseIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateBuyerIdentifyResponse) SetIdentifier(v CreateBuyerIdentifyResponseIdentifier)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CreateBuyerIdentifyResponse) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *CreateBuyerIdentifyResponse) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *CreateBuyerIdentifyResponse) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetVisionLabel

`func (o *CreateBuyerIdentifyResponse) GetVisionLabel() string`

GetVisionLabel returns the VisionLabel field if non-nil, zero value otherwise.

### GetVisionLabelOk

`func (o *CreateBuyerIdentifyResponse) GetVisionLabelOk() (*string, bool)`

GetVisionLabelOk returns a tuple with the VisionLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisionLabel

`func (o *CreateBuyerIdentifyResponse) SetVisionLabel(v string)`

SetVisionLabel sets VisionLabel field to given value.

### HasVisionLabel

`func (o *CreateBuyerIdentifyResponse) HasVisionLabel() bool`

HasVisionLabel returns a boolean if a field has been set.

### SetVisionLabelNil

`func (o *CreateBuyerIdentifyResponse) SetVisionLabelNil(b bool)`

 SetVisionLabelNil sets the value for VisionLabel to be an explicit nil

### UnsetVisionLabel
`func (o *CreateBuyerIdentifyResponse) UnsetVisionLabel()`

UnsetVisionLabel ensures that no value is present for VisionLabel, not even an explicit nil
### GetVisionQuotaExhausted

`func (o *CreateBuyerIdentifyResponse) GetVisionQuotaExhausted() bool`

GetVisionQuotaExhausted returns the VisionQuotaExhausted field if non-nil, zero value otherwise.

### GetVisionQuotaExhaustedOk

`func (o *CreateBuyerIdentifyResponse) GetVisionQuotaExhaustedOk() (*bool, bool)`

GetVisionQuotaExhaustedOk returns a tuple with the VisionQuotaExhausted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisionQuotaExhausted

`func (o *CreateBuyerIdentifyResponse) SetVisionQuotaExhausted(v bool)`

SetVisionQuotaExhausted sets VisionQuotaExhausted field to given value.


### GetVisualMatches

`func (o *CreateBuyerIdentifyResponse) GetVisualMatches() []CreateBuyerIdentifyResponseVisualMatches`

GetVisualMatches returns the VisualMatches field if non-nil, zero value otherwise.

### GetVisualMatchesOk

`func (o *CreateBuyerIdentifyResponse) GetVisualMatchesOk() (*[]CreateBuyerIdentifyResponseVisualMatches, bool)`

GetVisualMatchesOk returns a tuple with the VisualMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualMatches

`func (o *CreateBuyerIdentifyResponse) SetVisualMatches(v []CreateBuyerIdentifyResponseVisualMatches)`

SetVisualMatches sets VisualMatches field to given value.


### GetVerdict

`func (o *CreateBuyerIdentifyResponse) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *CreateBuyerIdentifyResponse) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *CreateBuyerIdentifyResponse) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.


### GetCrossly

`func (o *CreateBuyerIdentifyResponse) GetCrossly() GetBuyerAnywhereResponseCrossly`

GetCrossly returns the Crossly field if non-nil, zero value otherwise.

### GetCrosslyOk

`func (o *CreateBuyerIdentifyResponse) GetCrosslyOk() (*GetBuyerAnywhereResponseCrossly, bool)`

GetCrosslyOk returns a tuple with the Crossly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossly

`func (o *CreateBuyerIdentifyResponse) SetCrossly(v GetBuyerAnywhereResponseCrossly)`

SetCrossly sets Crossly field to given value.

### HasCrossly

`func (o *CreateBuyerIdentifyResponse) HasCrossly() bool`

HasCrossly returns a boolean if a field has been set.

### SetCrosslyNil

`func (o *CreateBuyerIdentifyResponse) SetCrosslyNil(b bool)`

 SetCrosslyNil sets the value for Crossly to be an explicit nil

### UnsetCrossly
`func (o *CreateBuyerIdentifyResponse) UnsetCrossly()`

UnsetCrossly ensures that no value is present for Crossly, not even an explicit nil
### GetOffsite

`func (o *CreateBuyerIdentifyResponse) GetOffsite() GetBuyerAnywhereResponseOffsite`

GetOffsite returns the Offsite field if non-nil, zero value otherwise.

### GetOffsiteOk

`func (o *CreateBuyerIdentifyResponse) GetOffsiteOk() (*GetBuyerAnywhereResponseOffsite, bool)`

GetOffsiteOk returns a tuple with the Offsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsite

`func (o *CreateBuyerIdentifyResponse) SetOffsite(v GetBuyerAnywhereResponseOffsite)`

SetOffsite sets Offsite field to given value.

### HasOffsite

`func (o *CreateBuyerIdentifyResponse) HasOffsite() bool`

HasOffsite returns a boolean if a field has been set.

### SetOffsiteNil

`func (o *CreateBuyerIdentifyResponse) SetOffsiteNil(b bool)`

 SetOffsiteNil sets the value for Offsite to be an explicit nil

### UnsetOffsite
`func (o *CreateBuyerIdentifyResponse) UnsetOffsite()`

UnsetOffsite ensures that no value is present for Offsite, not even an explicit nil
### GetAlternates

`func (o *CreateBuyerIdentifyResponse) GetAlternates() []GetBuyerAnywhereResponseAlternates`

GetAlternates returns the Alternates field if non-nil, zero value otherwise.

### GetAlternatesOk

`func (o *CreateBuyerIdentifyResponse) GetAlternatesOk() (*[]GetBuyerAnywhereResponseAlternates, bool)`

GetAlternatesOk returns a tuple with the Alternates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternates

`func (o *CreateBuyerIdentifyResponse) SetAlternates(v []GetBuyerAnywhereResponseAlternates)`

SetAlternates sets Alternates field to given value.


### GetSavingCents

`func (o *CreateBuyerIdentifyResponse) GetSavingCents() float32`

GetSavingCents returns the SavingCents field if non-nil, zero value otherwise.

### GetSavingCentsOk

`func (o *CreateBuyerIdentifyResponse) GetSavingCentsOk() (*float32, bool)`

GetSavingCentsOk returns a tuple with the SavingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingCents

`func (o *CreateBuyerIdentifyResponse) SetSavingCents(v float32)`

SetSavingCents sets SavingCents field to given value.

### HasSavingCents

`func (o *CreateBuyerIdentifyResponse) HasSavingCents() bool`

HasSavingCents returns a boolean if a field has been set.

### SetSavingCentsNil

`func (o *CreateBuyerIdentifyResponse) SetSavingCentsNil(b bool)`

 SetSavingCentsNil sets the value for SavingCents to be an explicit nil

### UnsetSavingCents
`func (o *CreateBuyerIdentifyResponse) UnsetSavingCents()`

UnsetSavingCents ensures that no value is present for SavingCents, not even an explicit nil
### GetShippingUnknown

`func (o *CreateBuyerIdentifyResponse) GetShippingUnknown() bool`

GetShippingUnknown returns the ShippingUnknown field if non-nil, zero value otherwise.

### GetShippingUnknownOk

`func (o *CreateBuyerIdentifyResponse) GetShippingUnknownOk() (*bool, bool)`

GetShippingUnknownOk returns a tuple with the ShippingUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingUnknown

`func (o *CreateBuyerIdentifyResponse) SetShippingUnknown(v bool)`

SetShippingUnknown sets ShippingUnknown field to given value.


### GetHud

`func (o *CreateBuyerIdentifyResponse) GetHud() CreateBuyerIdentifyResponseHud`

GetHud returns the Hud field if non-nil, zero value otherwise.

### GetHudOk

`func (o *CreateBuyerIdentifyResponse) GetHudOk() (*CreateBuyerIdentifyResponseHud, bool)`

GetHudOk returns a tuple with the Hud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHud

`func (o *CreateBuyerIdentifyResponse) SetHud(v CreateBuyerIdentifyResponseHud)`

SetHud sets Hud field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


