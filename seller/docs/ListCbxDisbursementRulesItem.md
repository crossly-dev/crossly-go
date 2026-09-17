# ListCbxDisbursementRulesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**ThresholdBaseUnits** | **string** |  | 
**DistributeBps** | **float32** |  | 
**CheckCadenceHours** | **float32** |  | 
**RequireCoverageBps** | **float32** |  | 
**Metric** | **string** |  | 
**LookbackDays** | **float32** |  | 
**MaxRecipients** | **float32** |  | 
**IsActive** | **bool** |  | 
**LastFiredAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListCbxDisbursementRulesItem

`func NewListCbxDisbursementRulesItem(id string, name string, thresholdBaseUnits string, distributeBps float32, checkCadenceHours float32, requireCoverageBps float32, metric string, lookbackDays float32, maxRecipients float32, isActive bool, ) *ListCbxDisbursementRulesItem`

NewListCbxDisbursementRulesItem instantiates a new ListCbxDisbursementRulesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCbxDisbursementRulesItemWithDefaults

`func NewListCbxDisbursementRulesItemWithDefaults() *ListCbxDisbursementRulesItem`

NewListCbxDisbursementRulesItemWithDefaults instantiates a new ListCbxDisbursementRulesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCbxDisbursementRulesItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCbxDisbursementRulesItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCbxDisbursementRulesItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListCbxDisbursementRulesItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCbxDisbursementRulesItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCbxDisbursementRulesItem) SetName(v string)`

SetName sets Name field to given value.


### GetThresholdBaseUnits

`func (o *ListCbxDisbursementRulesItem) GetThresholdBaseUnits() string`

GetThresholdBaseUnits returns the ThresholdBaseUnits field if non-nil, zero value otherwise.

### GetThresholdBaseUnitsOk

`func (o *ListCbxDisbursementRulesItem) GetThresholdBaseUnitsOk() (*string, bool)`

GetThresholdBaseUnitsOk returns a tuple with the ThresholdBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThresholdBaseUnits

`func (o *ListCbxDisbursementRulesItem) SetThresholdBaseUnits(v string)`

SetThresholdBaseUnits sets ThresholdBaseUnits field to given value.


### GetDistributeBps

`func (o *ListCbxDisbursementRulesItem) GetDistributeBps() float32`

GetDistributeBps returns the DistributeBps field if non-nil, zero value otherwise.

### GetDistributeBpsOk

`func (o *ListCbxDisbursementRulesItem) GetDistributeBpsOk() (*float32, bool)`

GetDistributeBpsOk returns a tuple with the DistributeBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistributeBps

`func (o *ListCbxDisbursementRulesItem) SetDistributeBps(v float32)`

SetDistributeBps sets DistributeBps field to given value.


### GetCheckCadenceHours

`func (o *ListCbxDisbursementRulesItem) GetCheckCadenceHours() float32`

GetCheckCadenceHours returns the CheckCadenceHours field if non-nil, zero value otherwise.

### GetCheckCadenceHoursOk

`func (o *ListCbxDisbursementRulesItem) GetCheckCadenceHoursOk() (*float32, bool)`

GetCheckCadenceHoursOk returns a tuple with the CheckCadenceHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckCadenceHours

`func (o *ListCbxDisbursementRulesItem) SetCheckCadenceHours(v float32)`

SetCheckCadenceHours sets CheckCadenceHours field to given value.


### GetRequireCoverageBps

`func (o *ListCbxDisbursementRulesItem) GetRequireCoverageBps() float32`

GetRequireCoverageBps returns the RequireCoverageBps field if non-nil, zero value otherwise.

### GetRequireCoverageBpsOk

`func (o *ListCbxDisbursementRulesItem) GetRequireCoverageBpsOk() (*float32, bool)`

GetRequireCoverageBpsOk returns a tuple with the RequireCoverageBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequireCoverageBps

`func (o *ListCbxDisbursementRulesItem) SetRequireCoverageBps(v float32)`

SetRequireCoverageBps sets RequireCoverageBps field to given value.


### GetMetric

`func (o *ListCbxDisbursementRulesItem) GetMetric() string`

GetMetric returns the Metric field if non-nil, zero value otherwise.

### GetMetricOk

`func (o *ListCbxDisbursementRulesItem) GetMetricOk() (*string, bool)`

GetMetricOk returns a tuple with the Metric field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetric

`func (o *ListCbxDisbursementRulesItem) SetMetric(v string)`

SetMetric sets Metric field to given value.


### GetLookbackDays

`func (o *ListCbxDisbursementRulesItem) GetLookbackDays() float32`

GetLookbackDays returns the LookbackDays field if non-nil, zero value otherwise.

### GetLookbackDaysOk

`func (o *ListCbxDisbursementRulesItem) GetLookbackDaysOk() (*float32, bool)`

GetLookbackDaysOk returns a tuple with the LookbackDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookbackDays

`func (o *ListCbxDisbursementRulesItem) SetLookbackDays(v float32)`

SetLookbackDays sets LookbackDays field to given value.


### GetMaxRecipients

`func (o *ListCbxDisbursementRulesItem) GetMaxRecipients() float32`

GetMaxRecipients returns the MaxRecipients field if non-nil, zero value otherwise.

### GetMaxRecipientsOk

`func (o *ListCbxDisbursementRulesItem) GetMaxRecipientsOk() (*float32, bool)`

GetMaxRecipientsOk returns a tuple with the MaxRecipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRecipients

`func (o *ListCbxDisbursementRulesItem) SetMaxRecipients(v float32)`

SetMaxRecipients sets MaxRecipients field to given value.


### GetIsActive

`func (o *ListCbxDisbursementRulesItem) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ListCbxDisbursementRulesItem) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ListCbxDisbursementRulesItem) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetLastFiredAt

`func (o *ListCbxDisbursementRulesItem) GetLastFiredAt() string`

GetLastFiredAt returns the LastFiredAt field if non-nil, zero value otherwise.

### GetLastFiredAtOk

`func (o *ListCbxDisbursementRulesItem) GetLastFiredAtOk() (*string, bool)`

GetLastFiredAtOk returns a tuple with the LastFiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastFiredAt

`func (o *ListCbxDisbursementRulesItem) SetLastFiredAt(v string)`

SetLastFiredAt sets LastFiredAt field to given value.

### HasLastFiredAt

`func (o *ListCbxDisbursementRulesItem) HasLastFiredAt() bool`

HasLastFiredAt returns a boolean if a field has been set.

### SetLastFiredAtNil

`func (o *ListCbxDisbursementRulesItem) SetLastFiredAtNil(b bool)`

 SetLastFiredAtNil sets the value for LastFiredAt to be an explicit nil

### UnsetLastFiredAt
`func (o *ListCbxDisbursementRulesItem) UnsetLastFiredAt()`

UnsetLastFiredAt ensures that no value is present for LastFiredAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


