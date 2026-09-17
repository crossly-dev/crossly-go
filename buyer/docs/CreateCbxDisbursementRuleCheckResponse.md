# CreateCbxDisbursementRuleCheckResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outcome** | **string** |  | 
**Fired** | **bool** |  | 
**PoolBaseUnits** | **string** |  | 
**ThresholdBaseUnits** | **string** |  | 
**DistributeBaseUnits** | **string** |  | 
**CoverageBps** | Pointer to **NullableFloat32** |  | [optional] 
**CampaignId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateCbxDisbursementRuleCheckResponse

`func NewCreateCbxDisbursementRuleCheckResponse(outcome string, fired bool, poolBaseUnits string, thresholdBaseUnits string, distributeBaseUnits string, ) *CreateCbxDisbursementRuleCheckResponse`

NewCreateCbxDisbursementRuleCheckResponse instantiates a new CreateCbxDisbursementRuleCheckResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCbxDisbursementRuleCheckResponseWithDefaults

`func NewCreateCbxDisbursementRuleCheckResponseWithDefaults() *CreateCbxDisbursementRuleCheckResponse`

NewCreateCbxDisbursementRuleCheckResponseWithDefaults instantiates a new CreateCbxDisbursementRuleCheckResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutcome

`func (o *CreateCbxDisbursementRuleCheckResponse) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *CreateCbxDisbursementRuleCheckResponse) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *CreateCbxDisbursementRuleCheckResponse) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetFired

`func (o *CreateCbxDisbursementRuleCheckResponse) GetFired() bool`

GetFired returns the Fired field if non-nil, zero value otherwise.

### GetFiredOk

`func (o *CreateCbxDisbursementRuleCheckResponse) GetFiredOk() (*bool, bool)`

GetFiredOk returns a tuple with the Fired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFired

`func (o *CreateCbxDisbursementRuleCheckResponse) SetFired(v bool)`

SetFired sets Fired field to given value.


### GetPoolBaseUnits

`func (o *CreateCbxDisbursementRuleCheckResponse) GetPoolBaseUnits() string`

GetPoolBaseUnits returns the PoolBaseUnits field if non-nil, zero value otherwise.

### GetPoolBaseUnitsOk

`func (o *CreateCbxDisbursementRuleCheckResponse) GetPoolBaseUnitsOk() (*string, bool)`

GetPoolBaseUnitsOk returns a tuple with the PoolBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolBaseUnits

`func (o *CreateCbxDisbursementRuleCheckResponse) SetPoolBaseUnits(v string)`

SetPoolBaseUnits sets PoolBaseUnits field to given value.


### GetThresholdBaseUnits

`func (o *CreateCbxDisbursementRuleCheckResponse) GetThresholdBaseUnits() string`

GetThresholdBaseUnits returns the ThresholdBaseUnits field if non-nil, zero value otherwise.

### GetThresholdBaseUnitsOk

`func (o *CreateCbxDisbursementRuleCheckResponse) GetThresholdBaseUnitsOk() (*string, bool)`

GetThresholdBaseUnitsOk returns a tuple with the ThresholdBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdBaseUnits

`func (o *CreateCbxDisbursementRuleCheckResponse) SetThresholdBaseUnits(v string)`

SetThresholdBaseUnits sets ThresholdBaseUnits field to given value.


### GetDistributeBaseUnits

`func (o *CreateCbxDisbursementRuleCheckResponse) GetDistributeBaseUnits() string`

GetDistributeBaseUnits returns the DistributeBaseUnits field if non-nil, zero value otherwise.

### GetDistributeBaseUnitsOk

`func (o *CreateCbxDisbursementRuleCheckResponse) GetDistributeBaseUnitsOk() (*string, bool)`

GetDistributeBaseUnitsOk returns a tuple with the DistributeBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributeBaseUnits

`func (o *CreateCbxDisbursementRuleCheckResponse) SetDistributeBaseUnits(v string)`

SetDistributeBaseUnits sets DistributeBaseUnits field to given value.


### GetCoverageBps

`func (o *CreateCbxDisbursementRuleCheckResponse) GetCoverageBps() float32`

GetCoverageBps returns the CoverageBps field if non-nil, zero value otherwise.

### GetCoverageBpsOk

`func (o *CreateCbxDisbursementRuleCheckResponse) GetCoverageBpsOk() (*float32, bool)`

GetCoverageBpsOk returns a tuple with the CoverageBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageBps

`func (o *CreateCbxDisbursementRuleCheckResponse) SetCoverageBps(v float32)`

SetCoverageBps sets CoverageBps field to given value.

### HasCoverageBps

`func (o *CreateCbxDisbursementRuleCheckResponse) HasCoverageBps() bool`

HasCoverageBps returns a boolean if a field has been set.

### SetCoverageBpsNil

`func (o *CreateCbxDisbursementRuleCheckResponse) SetCoverageBpsNil(b bool)`

 SetCoverageBpsNil sets the value for CoverageBps to be an explicit nil

### UnsetCoverageBps
`func (o *CreateCbxDisbursementRuleCheckResponse) UnsetCoverageBps()`

UnsetCoverageBps ensures that no value is present for CoverageBps, not even an explicit nil
### GetCampaignId

`func (o *CreateCbxDisbursementRuleCheckResponse) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CreateCbxDisbursementRuleCheckResponse) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CreateCbxDisbursementRuleCheckResponse) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.

### HasCampaignId

`func (o *CreateCbxDisbursementRuleCheckResponse) HasCampaignId() bool`

HasCampaignId returns a boolean if a field has been set.

### SetCampaignIdNil

`func (o *CreateCbxDisbursementRuleCheckResponse) SetCampaignIdNil(b bool)`

 SetCampaignIdNil sets the value for CampaignId to be an explicit nil

### UnsetCampaignId
`func (o *CreateCbxDisbursementRuleCheckResponse) UnsetCampaignId()`

UnsetCampaignId ensures that no value is present for CampaignId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


