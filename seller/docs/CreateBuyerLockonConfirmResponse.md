# CreateBuyerLockonConfirmResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LockonId** | **string** |  | 
**Status** | **string** |  | 
**Identifier** | Pointer to [**NullableCreateBuyerIdentifyResponseIdentifier**](CreateBuyerIdentifyResponseIdentifier.md) |  | [optional] 
**Candidates** | [**[]CreateBuyerLockonObserveResponseCandidates**](CreateBuyerLockonObserveResponseCandidates.md) | Present when we could not settle it alone. Show them; a pinch on one is the cheapest, strongest disambiguation available. | 
**ObservationCount** | **float32** |  | 
**VisionCalls** | **float32** |  | 
**VisionQuotaExhausted** | **bool** |  | 
**Verdict** | **string** |  | 
**Crossly** | Pointer to [**NullableGetBuyerAnywhereResponseCrossly**](GetBuyerAnywhereResponseCrossly.md) |  | [optional] 
**Offsite** | Pointer to [**NullableGetBuyerAnywhereResponseOffsite**](GetBuyerAnywhereResponseOffsite.md) |  | [optional] 
**Alternates** | [**[]GetBuyerAnywhereResponseAlternates**](GetBuyerAnywhereResponseAlternates.md) |  | 
**SavingCents** | Pointer to **NullableFloat32** |  | [optional] 
**ShippingUnknown** | **bool** |  | 
**Hud** | [**CreateBuyerIdentifyResponseHud**](CreateBuyerIdentifyResponseHud.md) |  | 

## Methods

### NewCreateBuyerLockonConfirmResponse

`func NewCreateBuyerLockonConfirmResponse(lockonId string, status string, candidates []CreateBuyerLockonObserveResponseCandidates, observationCount float32, visionCalls float32, visionQuotaExhausted bool, verdict string, alternates []GetBuyerAnywhereResponseAlternates, shippingUnknown bool, hud CreateBuyerIdentifyResponseHud, ) *CreateBuyerLockonConfirmResponse`

NewCreateBuyerLockonConfirmResponse instantiates a new CreateBuyerLockonConfirmResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerLockonConfirmResponseWithDefaults

`func NewCreateBuyerLockonConfirmResponseWithDefaults() *CreateBuyerLockonConfirmResponse`

NewCreateBuyerLockonConfirmResponseWithDefaults instantiates a new CreateBuyerLockonConfirmResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLockonId

`func (o *CreateBuyerLockonConfirmResponse) GetLockonId() string`

GetLockonId returns the LockonId field if non-nil, zero value otherwise.

### GetLockonIdOk

`func (o *CreateBuyerLockonConfirmResponse) GetLockonIdOk() (*string, bool)`

GetLockonIdOk returns a tuple with the LockonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockonId

`func (o *CreateBuyerLockonConfirmResponse) SetLockonId(v string)`

SetLockonId sets LockonId field to given value.


### GetStatus

`func (o *CreateBuyerLockonConfirmResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateBuyerLockonConfirmResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateBuyerLockonConfirmResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIdentifier

`func (o *CreateBuyerLockonConfirmResponse) GetIdentifier() CreateBuyerIdentifyResponseIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateBuyerLockonConfirmResponse) GetIdentifierOk() (*CreateBuyerIdentifyResponseIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateBuyerLockonConfirmResponse) SetIdentifier(v CreateBuyerIdentifyResponseIdentifier)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CreateBuyerLockonConfirmResponse) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *CreateBuyerLockonConfirmResponse) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *CreateBuyerLockonConfirmResponse) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetCandidates

`func (o *CreateBuyerLockonConfirmResponse) GetCandidates() []CreateBuyerLockonObserveResponseCandidates`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *CreateBuyerLockonConfirmResponse) GetCandidatesOk() (*[]CreateBuyerLockonObserveResponseCandidates, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *CreateBuyerLockonConfirmResponse) SetCandidates(v []CreateBuyerLockonObserveResponseCandidates)`

SetCandidates sets Candidates field to given value.


### GetObservationCount

`func (o *CreateBuyerLockonConfirmResponse) GetObservationCount() float32`

GetObservationCount returns the ObservationCount field if non-nil, zero value otherwise.

### GetObservationCountOk

`func (o *CreateBuyerLockonConfirmResponse) GetObservationCountOk() (*float32, bool)`

GetObservationCountOk returns a tuple with the ObservationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationCount

`func (o *CreateBuyerLockonConfirmResponse) SetObservationCount(v float32)`

SetObservationCount sets ObservationCount field to given value.


### GetVisionCalls

`func (o *CreateBuyerLockonConfirmResponse) GetVisionCalls() float32`

GetVisionCalls returns the VisionCalls field if non-nil, zero value otherwise.

### GetVisionCallsOk

`func (o *CreateBuyerLockonConfirmResponse) GetVisionCallsOk() (*float32, bool)`

GetVisionCallsOk returns a tuple with the VisionCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisionCalls

`func (o *CreateBuyerLockonConfirmResponse) SetVisionCalls(v float32)`

SetVisionCalls sets VisionCalls field to given value.


### GetVisionQuotaExhausted

`func (o *CreateBuyerLockonConfirmResponse) GetVisionQuotaExhausted() bool`

GetVisionQuotaExhausted returns the VisionQuotaExhausted field if non-nil, zero value otherwise.

### GetVisionQuotaExhaustedOk

`func (o *CreateBuyerLockonConfirmResponse) GetVisionQuotaExhaustedOk() (*bool, bool)`

GetVisionQuotaExhaustedOk returns a tuple with the VisionQuotaExhausted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisionQuotaExhausted

`func (o *CreateBuyerLockonConfirmResponse) SetVisionQuotaExhausted(v bool)`

SetVisionQuotaExhausted sets VisionQuotaExhausted field to given value.


### GetVerdict

`func (o *CreateBuyerLockonConfirmResponse) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *CreateBuyerLockonConfirmResponse) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *CreateBuyerLockonConfirmResponse) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.


### GetCrossly

`func (o *CreateBuyerLockonConfirmResponse) GetCrossly() GetBuyerAnywhereResponseCrossly`

GetCrossly returns the Crossly field if non-nil, zero value otherwise.

### GetCrosslyOk

`func (o *CreateBuyerLockonConfirmResponse) GetCrosslyOk() (*GetBuyerAnywhereResponseCrossly, bool)`

GetCrosslyOk returns a tuple with the Crossly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossly

`func (o *CreateBuyerLockonConfirmResponse) SetCrossly(v GetBuyerAnywhereResponseCrossly)`

SetCrossly sets Crossly field to given value.

### HasCrossly

`func (o *CreateBuyerLockonConfirmResponse) HasCrossly() bool`

HasCrossly returns a boolean if a field has been set.

### SetCrosslyNil

`func (o *CreateBuyerLockonConfirmResponse) SetCrosslyNil(b bool)`

 SetCrosslyNil sets the value for Crossly to be an explicit nil

### UnsetCrossly
`func (o *CreateBuyerLockonConfirmResponse) UnsetCrossly()`

UnsetCrossly ensures that no value is present for Crossly, not even an explicit nil
### GetOffsite

`func (o *CreateBuyerLockonConfirmResponse) GetOffsite() GetBuyerAnywhereResponseOffsite`

GetOffsite returns the Offsite field if non-nil, zero value otherwise.

### GetOffsiteOk

`func (o *CreateBuyerLockonConfirmResponse) GetOffsiteOk() (*GetBuyerAnywhereResponseOffsite, bool)`

GetOffsiteOk returns a tuple with the Offsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsite

`func (o *CreateBuyerLockonConfirmResponse) SetOffsite(v GetBuyerAnywhereResponseOffsite)`

SetOffsite sets Offsite field to given value.

### HasOffsite

`func (o *CreateBuyerLockonConfirmResponse) HasOffsite() bool`

HasOffsite returns a boolean if a field has been set.

### SetOffsiteNil

`func (o *CreateBuyerLockonConfirmResponse) SetOffsiteNil(b bool)`

 SetOffsiteNil sets the value for Offsite to be an explicit nil

### UnsetOffsite
`func (o *CreateBuyerLockonConfirmResponse) UnsetOffsite()`

UnsetOffsite ensures that no value is present for Offsite, not even an explicit nil
### GetAlternates

`func (o *CreateBuyerLockonConfirmResponse) GetAlternates() []GetBuyerAnywhereResponseAlternates`

GetAlternates returns the Alternates field if non-nil, zero value otherwise.

### GetAlternatesOk

`func (o *CreateBuyerLockonConfirmResponse) GetAlternatesOk() (*[]GetBuyerAnywhereResponseAlternates, bool)`

GetAlternatesOk returns a tuple with the Alternates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternates

`func (o *CreateBuyerLockonConfirmResponse) SetAlternates(v []GetBuyerAnywhereResponseAlternates)`

SetAlternates sets Alternates field to given value.


### GetSavingCents

`func (o *CreateBuyerLockonConfirmResponse) GetSavingCents() float32`

GetSavingCents returns the SavingCents field if non-nil, zero value otherwise.

### GetSavingCentsOk

`func (o *CreateBuyerLockonConfirmResponse) GetSavingCentsOk() (*float32, bool)`

GetSavingCentsOk returns a tuple with the SavingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingCents

`func (o *CreateBuyerLockonConfirmResponse) SetSavingCents(v float32)`

SetSavingCents sets SavingCents field to given value.

### HasSavingCents

`func (o *CreateBuyerLockonConfirmResponse) HasSavingCents() bool`

HasSavingCents returns a boolean if a field has been set.

### SetSavingCentsNil

`func (o *CreateBuyerLockonConfirmResponse) SetSavingCentsNil(b bool)`

 SetSavingCentsNil sets the value for SavingCents to be an explicit nil

### UnsetSavingCents
`func (o *CreateBuyerLockonConfirmResponse) UnsetSavingCents()`

UnsetSavingCents ensures that no value is present for SavingCents, not even an explicit nil
### GetShippingUnknown

`func (o *CreateBuyerLockonConfirmResponse) GetShippingUnknown() bool`

GetShippingUnknown returns the ShippingUnknown field if non-nil, zero value otherwise.

### GetShippingUnknownOk

`func (o *CreateBuyerLockonConfirmResponse) GetShippingUnknownOk() (*bool, bool)`

GetShippingUnknownOk returns a tuple with the ShippingUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingUnknown

`func (o *CreateBuyerLockonConfirmResponse) SetShippingUnknown(v bool)`

SetShippingUnknown sets ShippingUnknown field to given value.


### GetHud

`func (o *CreateBuyerLockonConfirmResponse) GetHud() CreateBuyerIdentifyResponseHud`

GetHud returns the Hud field if non-nil, zero value otherwise.

### GetHudOk

`func (o *CreateBuyerLockonConfirmResponse) GetHudOk() (*CreateBuyerIdentifyResponseHud, bool)`

GetHudOk returns a tuple with the Hud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHud

`func (o *CreateBuyerLockonConfirmResponse) SetHud(v CreateBuyerIdentifyResponseHud)`

SetHud sets Hud field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


