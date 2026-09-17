# ListCbxDisbursementProgressItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleId** | **string** |  | 
**Name** | **string** |  | 
**PoolBaseUnits** | **string** |  | 
**ThresholdBaseUnits** | **string** |  | 
**ProgressBps** | **float32** |  | 
**NextEligibleAt** | Pointer to **NullableString** |  | [optional] 
**LastFiredAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListCbxDisbursementProgressItem

`func NewListCbxDisbursementProgressItem(ruleId string, name string, poolBaseUnits string, thresholdBaseUnits string, progressBps float32, ) *ListCbxDisbursementProgressItem`

NewListCbxDisbursementProgressItem instantiates a new ListCbxDisbursementProgressItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCbxDisbursementProgressItemWithDefaults

`func NewListCbxDisbursementProgressItemWithDefaults() *ListCbxDisbursementProgressItem`

NewListCbxDisbursementProgressItemWithDefaults instantiates a new ListCbxDisbursementProgressItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleId

`func (o *ListCbxDisbursementProgressItem) GetRuleId() string`

GetRuleId returns the RuleId field if non-nil, zero value otherwise.

### GetRuleIdOk

`func (o *ListCbxDisbursementProgressItem) GetRuleIdOk() (*string, bool)`

GetRuleIdOk returns a tuple with the RuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleId

`func (o *ListCbxDisbursementProgressItem) SetRuleId(v string)`

SetRuleId sets RuleId field to given value.


### GetName

`func (o *ListCbxDisbursementProgressItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCbxDisbursementProgressItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCbxDisbursementProgressItem) SetName(v string)`

SetName sets Name field to given value.


### GetPoolBaseUnits

`func (o *ListCbxDisbursementProgressItem) GetPoolBaseUnits() string`

GetPoolBaseUnits returns the PoolBaseUnits field if non-nil, zero value otherwise.

### GetPoolBaseUnitsOk

`func (o *ListCbxDisbursementProgressItem) GetPoolBaseUnitsOk() (*string, bool)`

GetPoolBaseUnitsOk returns a tuple with the PoolBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoolBaseUnits

`func (o *ListCbxDisbursementProgressItem) SetPoolBaseUnits(v string)`

SetPoolBaseUnits sets PoolBaseUnits field to given value.


### GetThresholdBaseUnits

`func (o *ListCbxDisbursementProgressItem) GetThresholdBaseUnits() string`

GetThresholdBaseUnits returns the ThresholdBaseUnits field if non-nil, zero value otherwise.

### GetThresholdBaseUnitsOk

`func (o *ListCbxDisbursementProgressItem) GetThresholdBaseUnitsOk() (*string, bool)`

GetThresholdBaseUnitsOk returns a tuple with the ThresholdBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdBaseUnits

`func (o *ListCbxDisbursementProgressItem) SetThresholdBaseUnits(v string)`

SetThresholdBaseUnits sets ThresholdBaseUnits field to given value.


### GetProgressBps

`func (o *ListCbxDisbursementProgressItem) GetProgressBps() float32`

GetProgressBps returns the ProgressBps field if non-nil, zero value otherwise.

### GetProgressBpsOk

`func (o *ListCbxDisbursementProgressItem) GetProgressBpsOk() (*float32, bool)`

GetProgressBpsOk returns a tuple with the ProgressBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgressBps

`func (o *ListCbxDisbursementProgressItem) SetProgressBps(v float32)`

SetProgressBps sets ProgressBps field to given value.


### GetNextEligibleAt

`func (o *ListCbxDisbursementProgressItem) GetNextEligibleAt() string`

GetNextEligibleAt returns the NextEligibleAt field if non-nil, zero value otherwise.

### GetNextEligibleAtOk

`func (o *ListCbxDisbursementProgressItem) GetNextEligibleAtOk() (*string, bool)`

GetNextEligibleAtOk returns a tuple with the NextEligibleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextEligibleAt

`func (o *ListCbxDisbursementProgressItem) SetNextEligibleAt(v string)`

SetNextEligibleAt sets NextEligibleAt field to given value.

### HasNextEligibleAt

`func (o *ListCbxDisbursementProgressItem) HasNextEligibleAt() bool`

HasNextEligibleAt returns a boolean if a field has been set.

### SetNextEligibleAtNil

`func (o *ListCbxDisbursementProgressItem) SetNextEligibleAtNil(b bool)`

 SetNextEligibleAtNil sets the value for NextEligibleAt to be an explicit nil

### UnsetNextEligibleAt
`func (o *ListCbxDisbursementProgressItem) UnsetNextEligibleAt()`

UnsetNextEligibleAt ensures that no value is present for NextEligibleAt, not even an explicit nil
### GetLastFiredAt

`func (o *ListCbxDisbursementProgressItem) GetLastFiredAt() string`

GetLastFiredAt returns the LastFiredAt field if non-nil, zero value otherwise.

### GetLastFiredAtOk

`func (o *ListCbxDisbursementProgressItem) GetLastFiredAtOk() (*string, bool)`

GetLastFiredAtOk returns a tuple with the LastFiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFiredAt

`func (o *ListCbxDisbursementProgressItem) SetLastFiredAt(v string)`

SetLastFiredAt sets LastFiredAt field to given value.

### HasLastFiredAt

`func (o *ListCbxDisbursementProgressItem) HasLastFiredAt() bool`

HasLastFiredAt returns a boolean if a field has been set.

### SetLastFiredAtNil

`func (o *ListCbxDisbursementProgressItem) SetLastFiredAtNil(b bool)`

 SetLastFiredAtNil sets the value for LastFiredAt to be an explicit nil

### UnsetLastFiredAt
`func (o *ListCbxDisbursementProgressItem) UnsetLastFiredAt()`

UnsetLastFiredAt ensures that no value is present for LastFiredAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


