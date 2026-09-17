# GetAdOffsiteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** |  | 
**MaxOffsiteShareBps** | Pointer to **NullableFloat32** |  | [optional] 
**MinReturnBps** | Pointer to **NullableFloat32** |  | [optional] 
**ReturnWindowDays** | Pointer to **NullableFloat32** |  | [optional] 
**AllowedNetworks** | Pointer to **[]string** |  | [optional] 
**AutoPausedAt** | Pointer to **NullableString** |  | [optional] 
**AutoPausedReason** | Pointer to **NullableString** |  | [optional] 
**OptedInAt** | Pointer to **NullableString** |  | [optional] 
**Terms** | [**GetAdOffsiteResponseTerms**](GetAdOffsiteResponseTerms.md) |  | 
**NetworkWired** | **bool** |  | 

## Methods

### NewGetAdOffsiteResponse

`func NewGetAdOffsiteResponse(enabled bool, terms GetAdOffsiteResponseTerms, networkWired bool, ) *GetAdOffsiteResponse`

NewGetAdOffsiteResponse instantiates a new GetAdOffsiteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdOffsiteResponseWithDefaults

`func NewGetAdOffsiteResponseWithDefaults() *GetAdOffsiteResponse`

NewGetAdOffsiteResponseWithDefaults instantiates a new GetAdOffsiteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *GetAdOffsiteResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetAdOffsiteResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetAdOffsiteResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetMaxOffsiteShareBps

`func (o *GetAdOffsiteResponse) GetMaxOffsiteShareBps() float32`

GetMaxOffsiteShareBps returns the MaxOffsiteShareBps field if non-nil, zero value otherwise.

### GetMaxOffsiteShareBpsOk

`func (o *GetAdOffsiteResponse) GetMaxOffsiteShareBpsOk() (*float32, bool)`

GetMaxOffsiteShareBpsOk returns a tuple with the MaxOffsiteShareBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxOffsiteShareBps

`func (o *GetAdOffsiteResponse) SetMaxOffsiteShareBps(v float32)`

SetMaxOffsiteShareBps sets MaxOffsiteShareBps field to given value.

### HasMaxOffsiteShareBps

`func (o *GetAdOffsiteResponse) HasMaxOffsiteShareBps() bool`

HasMaxOffsiteShareBps returns a boolean if a field has been set.

### SetMaxOffsiteShareBpsNil

`func (o *GetAdOffsiteResponse) SetMaxOffsiteShareBpsNil(b bool)`

 SetMaxOffsiteShareBpsNil sets the value for MaxOffsiteShareBps to be an explicit nil

### UnsetMaxOffsiteShareBps
`func (o *GetAdOffsiteResponse) UnsetMaxOffsiteShareBps()`

UnsetMaxOffsiteShareBps ensures that no value is present for MaxOffsiteShareBps, not even an explicit nil
### GetMinReturnBps

`func (o *GetAdOffsiteResponse) GetMinReturnBps() float32`

GetMinReturnBps returns the MinReturnBps field if non-nil, zero value otherwise.

### GetMinReturnBpsOk

`func (o *GetAdOffsiteResponse) GetMinReturnBpsOk() (*float32, bool)`

GetMinReturnBpsOk returns a tuple with the MinReturnBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinReturnBps

`func (o *GetAdOffsiteResponse) SetMinReturnBps(v float32)`

SetMinReturnBps sets MinReturnBps field to given value.

### HasMinReturnBps

`func (o *GetAdOffsiteResponse) HasMinReturnBps() bool`

HasMinReturnBps returns a boolean if a field has been set.

### SetMinReturnBpsNil

`func (o *GetAdOffsiteResponse) SetMinReturnBpsNil(b bool)`

 SetMinReturnBpsNil sets the value for MinReturnBps to be an explicit nil

### UnsetMinReturnBps
`func (o *GetAdOffsiteResponse) UnsetMinReturnBps()`

UnsetMinReturnBps ensures that no value is present for MinReturnBps, not even an explicit nil
### GetReturnWindowDays

`func (o *GetAdOffsiteResponse) GetReturnWindowDays() float32`

GetReturnWindowDays returns the ReturnWindowDays field if non-nil, zero value otherwise.

### GetReturnWindowDaysOk

`func (o *GetAdOffsiteResponse) GetReturnWindowDaysOk() (*float32, bool)`

GetReturnWindowDaysOk returns a tuple with the ReturnWindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnWindowDays

`func (o *GetAdOffsiteResponse) SetReturnWindowDays(v float32)`

SetReturnWindowDays sets ReturnWindowDays field to given value.

### HasReturnWindowDays

`func (o *GetAdOffsiteResponse) HasReturnWindowDays() bool`

HasReturnWindowDays returns a boolean if a field has been set.

### SetReturnWindowDaysNil

`func (o *GetAdOffsiteResponse) SetReturnWindowDaysNil(b bool)`

 SetReturnWindowDaysNil sets the value for ReturnWindowDays to be an explicit nil

### UnsetReturnWindowDays
`func (o *GetAdOffsiteResponse) UnsetReturnWindowDays()`

UnsetReturnWindowDays ensures that no value is present for ReturnWindowDays, not even an explicit nil
### GetAllowedNetworks

`func (o *GetAdOffsiteResponse) GetAllowedNetworks() []string`

GetAllowedNetworks returns the AllowedNetworks field if non-nil, zero value otherwise.

### GetAllowedNetworksOk

`func (o *GetAdOffsiteResponse) GetAllowedNetworksOk() (*[]string, bool)`

GetAllowedNetworksOk returns a tuple with the AllowedNetworks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowedNetworks

`func (o *GetAdOffsiteResponse) SetAllowedNetworks(v []string)`

SetAllowedNetworks sets AllowedNetworks field to given value.

### HasAllowedNetworks

`func (o *GetAdOffsiteResponse) HasAllowedNetworks() bool`

HasAllowedNetworks returns a boolean if a field has been set.

### SetAllowedNetworksNil

`func (o *GetAdOffsiteResponse) SetAllowedNetworksNil(b bool)`

 SetAllowedNetworksNil sets the value for AllowedNetworks to be an explicit nil

### UnsetAllowedNetworks
`func (o *GetAdOffsiteResponse) UnsetAllowedNetworks()`

UnsetAllowedNetworks ensures that no value is present for AllowedNetworks, not even an explicit nil
### GetAutoPausedAt

`func (o *GetAdOffsiteResponse) GetAutoPausedAt() string`

GetAutoPausedAt returns the AutoPausedAt field if non-nil, zero value otherwise.

### GetAutoPausedAtOk

`func (o *GetAdOffsiteResponse) GetAutoPausedAtOk() (*string, bool)`

GetAutoPausedAtOk returns a tuple with the AutoPausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPausedAt

`func (o *GetAdOffsiteResponse) SetAutoPausedAt(v string)`

SetAutoPausedAt sets AutoPausedAt field to given value.

### HasAutoPausedAt

`func (o *GetAdOffsiteResponse) HasAutoPausedAt() bool`

HasAutoPausedAt returns a boolean if a field has been set.

### SetAutoPausedAtNil

`func (o *GetAdOffsiteResponse) SetAutoPausedAtNil(b bool)`

 SetAutoPausedAtNil sets the value for AutoPausedAt to be an explicit nil

### UnsetAutoPausedAt
`func (o *GetAdOffsiteResponse) UnsetAutoPausedAt()`

UnsetAutoPausedAt ensures that no value is present for AutoPausedAt, not even an explicit nil
### GetAutoPausedReason

`func (o *GetAdOffsiteResponse) GetAutoPausedReason() string`

GetAutoPausedReason returns the AutoPausedReason field if non-nil, zero value otherwise.

### GetAutoPausedReasonOk

`func (o *GetAdOffsiteResponse) GetAutoPausedReasonOk() (*string, bool)`

GetAutoPausedReasonOk returns a tuple with the AutoPausedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPausedReason

`func (o *GetAdOffsiteResponse) SetAutoPausedReason(v string)`

SetAutoPausedReason sets AutoPausedReason field to given value.

### HasAutoPausedReason

`func (o *GetAdOffsiteResponse) HasAutoPausedReason() bool`

HasAutoPausedReason returns a boolean if a field has been set.

### SetAutoPausedReasonNil

`func (o *GetAdOffsiteResponse) SetAutoPausedReasonNil(b bool)`

 SetAutoPausedReasonNil sets the value for AutoPausedReason to be an explicit nil

### UnsetAutoPausedReason
`func (o *GetAdOffsiteResponse) UnsetAutoPausedReason()`

UnsetAutoPausedReason ensures that no value is present for AutoPausedReason, not even an explicit nil
### GetOptedInAt

`func (o *GetAdOffsiteResponse) GetOptedInAt() string`

GetOptedInAt returns the OptedInAt field if non-nil, zero value otherwise.

### GetOptedInAtOk

`func (o *GetAdOffsiteResponse) GetOptedInAtOk() (*string, bool)`

GetOptedInAtOk returns a tuple with the OptedInAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptedInAt

`func (o *GetAdOffsiteResponse) SetOptedInAt(v string)`

SetOptedInAt sets OptedInAt field to given value.

### HasOptedInAt

`func (o *GetAdOffsiteResponse) HasOptedInAt() bool`

HasOptedInAt returns a boolean if a field has been set.

### SetOptedInAtNil

`func (o *GetAdOffsiteResponse) SetOptedInAtNil(b bool)`

 SetOptedInAtNil sets the value for OptedInAt to be an explicit nil

### UnsetOptedInAt
`func (o *GetAdOffsiteResponse) UnsetOptedInAt()`

UnsetOptedInAt ensures that no value is present for OptedInAt, not even an explicit nil
### GetTerms

`func (o *GetAdOffsiteResponse) GetTerms() GetAdOffsiteResponseTerms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *GetAdOffsiteResponse) GetTermsOk() (*GetAdOffsiteResponseTerms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *GetAdOffsiteResponse) SetTerms(v GetAdOffsiteResponseTerms)`

SetTerms sets Terms field to given value.


### GetNetworkWired

`func (o *GetAdOffsiteResponse) GetNetworkWired() bool`

GetNetworkWired returns the NetworkWired field if non-nil, zero value otherwise.

### GetNetworkWiredOk

`func (o *GetAdOffsiteResponse) GetNetworkWiredOk() (*bool, bool)`

GetNetworkWiredOk returns a tuple with the NetworkWired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkWired

`func (o *GetAdOffsiteResponse) SetNetworkWired(v bool)`

SetNetworkWired sets NetworkWired field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


