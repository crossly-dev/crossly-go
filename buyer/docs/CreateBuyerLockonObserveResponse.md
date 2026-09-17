# CreateBuyerLockonObserveResponse

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

### NewCreateBuyerLockonObserveResponse

`func NewCreateBuyerLockonObserveResponse(lockonId string, status string, candidates []CreateBuyerLockonObserveResponseCandidates, observationCount float32, visionCalls float32, visionQuotaExhausted bool, verdict string, alternates []GetBuyerAnywhereResponseAlternates, shippingUnknown bool, hud CreateBuyerIdentifyResponseHud, ) *CreateBuyerLockonObserveResponse`

NewCreateBuyerLockonObserveResponse instantiates a new CreateBuyerLockonObserveResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerLockonObserveResponseWithDefaults

`func NewCreateBuyerLockonObserveResponseWithDefaults() *CreateBuyerLockonObserveResponse`

NewCreateBuyerLockonObserveResponseWithDefaults instantiates a new CreateBuyerLockonObserveResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLockonId

`func (o *CreateBuyerLockonObserveResponse) GetLockonId() string`

GetLockonId returns the LockonId field if non-nil, zero value otherwise.

### GetLockonIdOk

`func (o *CreateBuyerLockonObserveResponse) GetLockonIdOk() (*string, bool)`

GetLockonIdOk returns a tuple with the LockonId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLockonId

`func (o *CreateBuyerLockonObserveResponse) SetLockonId(v string)`

SetLockonId sets LockonId field to given value.


### GetStatus

`func (o *CreateBuyerLockonObserveResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateBuyerLockonObserveResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateBuyerLockonObserveResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetIdentifier

`func (o *CreateBuyerLockonObserveResponse) GetIdentifier() CreateBuyerIdentifyResponseIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *CreateBuyerLockonObserveResponse) GetIdentifierOk() (*CreateBuyerIdentifyResponseIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *CreateBuyerLockonObserveResponse) SetIdentifier(v CreateBuyerIdentifyResponseIdentifier)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *CreateBuyerLockonObserveResponse) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *CreateBuyerLockonObserveResponse) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *CreateBuyerLockonObserveResponse) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetCandidates

`func (o *CreateBuyerLockonObserveResponse) GetCandidates() []CreateBuyerLockonObserveResponseCandidates`

GetCandidates returns the Candidates field if non-nil, zero value otherwise.

### GetCandidatesOk

`func (o *CreateBuyerLockonObserveResponse) GetCandidatesOk() (*[]CreateBuyerLockonObserveResponseCandidates, bool)`

GetCandidatesOk returns a tuple with the Candidates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCandidates

`func (o *CreateBuyerLockonObserveResponse) SetCandidates(v []CreateBuyerLockonObserveResponseCandidates)`

SetCandidates sets Candidates field to given value.


### GetObservationCount

`func (o *CreateBuyerLockonObserveResponse) GetObservationCount() float32`

GetObservationCount returns the ObservationCount field if non-nil, zero value otherwise.

### GetObservationCountOk

`func (o *CreateBuyerLockonObserveResponse) GetObservationCountOk() (*float32, bool)`

GetObservationCountOk returns a tuple with the ObservationCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationCount

`func (o *CreateBuyerLockonObserveResponse) SetObservationCount(v float32)`

SetObservationCount sets ObservationCount field to given value.


### GetVisionCalls

`func (o *CreateBuyerLockonObserveResponse) GetVisionCalls() float32`

GetVisionCalls returns the VisionCalls field if non-nil, zero value otherwise.

### GetVisionCallsOk

`func (o *CreateBuyerLockonObserveResponse) GetVisionCallsOk() (*float32, bool)`

GetVisionCallsOk returns a tuple with the VisionCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisionCalls

`func (o *CreateBuyerLockonObserveResponse) SetVisionCalls(v float32)`

SetVisionCalls sets VisionCalls field to given value.


### GetVisionQuotaExhausted

`func (o *CreateBuyerLockonObserveResponse) GetVisionQuotaExhausted() bool`

GetVisionQuotaExhausted returns the VisionQuotaExhausted field if non-nil, zero value otherwise.

### GetVisionQuotaExhaustedOk

`func (o *CreateBuyerLockonObserveResponse) GetVisionQuotaExhaustedOk() (*bool, bool)`

GetVisionQuotaExhaustedOk returns a tuple with the VisionQuotaExhausted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisionQuotaExhausted

`func (o *CreateBuyerLockonObserveResponse) SetVisionQuotaExhausted(v bool)`

SetVisionQuotaExhausted sets VisionQuotaExhausted field to given value.


### GetVerdict

`func (o *CreateBuyerLockonObserveResponse) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *CreateBuyerLockonObserveResponse) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *CreateBuyerLockonObserveResponse) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.


### GetCrossly

`func (o *CreateBuyerLockonObserveResponse) GetCrossly() GetBuyerAnywhereResponseCrossly`

GetCrossly returns the Crossly field if non-nil, zero value otherwise.

### GetCrosslyOk

`func (o *CreateBuyerLockonObserveResponse) GetCrosslyOk() (*GetBuyerAnywhereResponseCrossly, bool)`

GetCrosslyOk returns a tuple with the Crossly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossly

`func (o *CreateBuyerLockonObserveResponse) SetCrossly(v GetBuyerAnywhereResponseCrossly)`

SetCrossly sets Crossly field to given value.

### HasCrossly

`func (o *CreateBuyerLockonObserveResponse) HasCrossly() bool`

HasCrossly returns a boolean if a field has been set.

### SetCrosslyNil

`func (o *CreateBuyerLockonObserveResponse) SetCrosslyNil(b bool)`

 SetCrosslyNil sets the value for Crossly to be an explicit nil

### UnsetCrossly
`func (o *CreateBuyerLockonObserveResponse) UnsetCrossly()`

UnsetCrossly ensures that no value is present for Crossly, not even an explicit nil
### GetOffsite

`func (o *CreateBuyerLockonObserveResponse) GetOffsite() GetBuyerAnywhereResponseOffsite`

GetOffsite returns the Offsite field if non-nil, zero value otherwise.

### GetOffsiteOk

`func (o *CreateBuyerLockonObserveResponse) GetOffsiteOk() (*GetBuyerAnywhereResponseOffsite, bool)`

GetOffsiteOk returns a tuple with the Offsite field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffsite

`func (o *CreateBuyerLockonObserveResponse) SetOffsite(v GetBuyerAnywhereResponseOffsite)`

SetOffsite sets Offsite field to given value.

### HasOffsite

`func (o *CreateBuyerLockonObserveResponse) HasOffsite() bool`

HasOffsite returns a boolean if a field has been set.

### SetOffsiteNil

`func (o *CreateBuyerLockonObserveResponse) SetOffsiteNil(b bool)`

 SetOffsiteNil sets the value for Offsite to be an explicit nil

### UnsetOffsite
`func (o *CreateBuyerLockonObserveResponse) UnsetOffsite()`

UnsetOffsite ensures that no value is present for Offsite, not even an explicit nil
### GetAlternates

`func (o *CreateBuyerLockonObserveResponse) GetAlternates() []GetBuyerAnywhereResponseAlternates`

GetAlternates returns the Alternates field if non-nil, zero value otherwise.

### GetAlternatesOk

`func (o *CreateBuyerLockonObserveResponse) GetAlternatesOk() (*[]GetBuyerAnywhereResponseAlternates, bool)`

GetAlternatesOk returns a tuple with the Alternates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternates

`func (o *CreateBuyerLockonObserveResponse) SetAlternates(v []GetBuyerAnywhereResponseAlternates)`

SetAlternates sets Alternates field to given value.


### GetSavingCents

`func (o *CreateBuyerLockonObserveResponse) GetSavingCents() float32`

GetSavingCents returns the SavingCents field if non-nil, zero value otherwise.

### GetSavingCentsOk

`func (o *CreateBuyerLockonObserveResponse) GetSavingCentsOk() (*float32, bool)`

GetSavingCentsOk returns a tuple with the SavingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingCents

`func (o *CreateBuyerLockonObserveResponse) SetSavingCents(v float32)`

SetSavingCents sets SavingCents field to given value.

### HasSavingCents

`func (o *CreateBuyerLockonObserveResponse) HasSavingCents() bool`

HasSavingCents returns a boolean if a field has been set.

### SetSavingCentsNil

`func (o *CreateBuyerLockonObserveResponse) SetSavingCentsNil(b bool)`

 SetSavingCentsNil sets the value for SavingCents to be an explicit nil

### UnsetSavingCents
`func (o *CreateBuyerLockonObserveResponse) UnsetSavingCents()`

UnsetSavingCents ensures that no value is present for SavingCents, not even an explicit nil
### GetShippingUnknown

`func (o *CreateBuyerLockonObserveResponse) GetShippingUnknown() bool`

GetShippingUnknown returns the ShippingUnknown field if non-nil, zero value otherwise.

### GetShippingUnknownOk

`func (o *CreateBuyerLockonObserveResponse) GetShippingUnknownOk() (*bool, bool)`

GetShippingUnknownOk returns a tuple with the ShippingUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingUnknown

`func (o *CreateBuyerLockonObserveResponse) SetShippingUnknown(v bool)`

SetShippingUnknown sets ShippingUnknown field to given value.


### GetHud

`func (o *CreateBuyerLockonObserveResponse) GetHud() CreateBuyerIdentifyResponseHud`

GetHud returns the Hud field if non-nil, zero value otherwise.

### GetHudOk

`func (o *CreateBuyerLockonObserveResponse) GetHudOk() (*CreateBuyerIdentifyResponseHud, bool)`

GetHudOk returns a tuple with the Hud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHud

`func (o *CreateBuyerLockonObserveResponse) SetHud(v CreateBuyerIdentifyResponseHud)`

SetHud sets Hud field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


