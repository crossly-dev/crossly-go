# GetAdOffsiteReportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SpendCents** | **float32** |  | 
**MediaCostCents** | **float32** |  | 
**ManagementFeeCents** | **float32** |  | 
**Impressions** | **float32** |  | 
**Clicks** | **float32** |  | 
**Conversions** | **float32** |  | 
**AttributedRevenueCents** | **float32** |  | 
**SpilloverConversions** | **float32** |  | 
**SpilloverRevenueCents** | **float32** |  | 
**ReturnBps** | **float32** |  | 
**Enabled** | **bool** |  | 
**AutoPausedAt** | Pointer to **NullableString** |  | [optional] 
**AutoPausedReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetAdOffsiteReportResponse

`func NewGetAdOffsiteReportResponse(spendCents float32, mediaCostCents float32, managementFeeCents float32, impressions float32, clicks float32, conversions float32, attributedRevenueCents float32, spilloverConversions float32, spilloverRevenueCents float32, returnBps float32, enabled bool, ) *GetAdOffsiteReportResponse`

NewGetAdOffsiteReportResponse instantiates a new GetAdOffsiteReportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAdOffsiteReportResponseWithDefaults

`func NewGetAdOffsiteReportResponseWithDefaults() *GetAdOffsiteReportResponse`

NewGetAdOffsiteReportResponseWithDefaults instantiates a new GetAdOffsiteReportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSpendCents

`func (o *GetAdOffsiteReportResponse) GetSpendCents() float32`

GetSpendCents returns the SpendCents field if non-nil, zero value otherwise.

### GetSpendCentsOk

`func (o *GetAdOffsiteReportResponse) GetSpendCentsOk() (*float32, bool)`

GetSpendCentsOk returns a tuple with the SpendCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpendCents

`func (o *GetAdOffsiteReportResponse) SetSpendCents(v float32)`

SetSpendCents sets SpendCents field to given value.


### GetMediaCostCents

`func (o *GetAdOffsiteReportResponse) GetMediaCostCents() float32`

GetMediaCostCents returns the MediaCostCents field if non-nil, zero value otherwise.

### GetMediaCostCentsOk

`func (o *GetAdOffsiteReportResponse) GetMediaCostCentsOk() (*float32, bool)`

GetMediaCostCentsOk returns a tuple with the MediaCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMediaCostCents

`func (o *GetAdOffsiteReportResponse) SetMediaCostCents(v float32)`

SetMediaCostCents sets MediaCostCents field to given value.


### GetManagementFeeCents

`func (o *GetAdOffsiteReportResponse) GetManagementFeeCents() float32`

GetManagementFeeCents returns the ManagementFeeCents field if non-nil, zero value otherwise.

### GetManagementFeeCentsOk

`func (o *GetAdOffsiteReportResponse) GetManagementFeeCentsOk() (*float32, bool)`

GetManagementFeeCentsOk returns a tuple with the ManagementFeeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagementFeeCents

`func (o *GetAdOffsiteReportResponse) SetManagementFeeCents(v float32)`

SetManagementFeeCents sets ManagementFeeCents field to given value.


### GetImpressions

`func (o *GetAdOffsiteReportResponse) GetImpressions() float32`

GetImpressions returns the Impressions field if non-nil, zero value otherwise.

### GetImpressionsOk

`func (o *GetAdOffsiteReportResponse) GetImpressionsOk() (*float32, bool)`

GetImpressionsOk returns a tuple with the Impressions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImpressions

`func (o *GetAdOffsiteReportResponse) SetImpressions(v float32)`

SetImpressions sets Impressions field to given value.


### GetClicks

`func (o *GetAdOffsiteReportResponse) GetClicks() float32`

GetClicks returns the Clicks field if non-nil, zero value otherwise.

### GetClicksOk

`func (o *GetAdOffsiteReportResponse) GetClicksOk() (*float32, bool)`

GetClicksOk returns a tuple with the Clicks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClicks

`func (o *GetAdOffsiteReportResponse) SetClicks(v float32)`

SetClicks sets Clicks field to given value.


### GetConversions

`func (o *GetAdOffsiteReportResponse) GetConversions() float32`

GetConversions returns the Conversions field if non-nil, zero value otherwise.

### GetConversionsOk

`func (o *GetAdOffsiteReportResponse) GetConversionsOk() (*float32, bool)`

GetConversionsOk returns a tuple with the Conversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversions

`func (o *GetAdOffsiteReportResponse) SetConversions(v float32)`

SetConversions sets Conversions field to given value.


### GetAttributedRevenueCents

`func (o *GetAdOffsiteReportResponse) GetAttributedRevenueCents() float32`

GetAttributedRevenueCents returns the AttributedRevenueCents field if non-nil, zero value otherwise.

### GetAttributedRevenueCentsOk

`func (o *GetAdOffsiteReportResponse) GetAttributedRevenueCentsOk() (*float32, bool)`

GetAttributedRevenueCentsOk returns a tuple with the AttributedRevenueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributedRevenueCents

`func (o *GetAdOffsiteReportResponse) SetAttributedRevenueCents(v float32)`

SetAttributedRevenueCents sets AttributedRevenueCents field to given value.


### GetSpilloverConversions

`func (o *GetAdOffsiteReportResponse) GetSpilloverConversions() float32`

GetSpilloverConversions returns the SpilloverConversions field if non-nil, zero value otherwise.

### GetSpilloverConversionsOk

`func (o *GetAdOffsiteReportResponse) GetSpilloverConversionsOk() (*float32, bool)`

GetSpilloverConversionsOk returns a tuple with the SpilloverConversions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpilloverConversions

`func (o *GetAdOffsiteReportResponse) SetSpilloverConversions(v float32)`

SetSpilloverConversions sets SpilloverConversions field to given value.


### GetSpilloverRevenueCents

`func (o *GetAdOffsiteReportResponse) GetSpilloverRevenueCents() float32`

GetSpilloverRevenueCents returns the SpilloverRevenueCents field if non-nil, zero value otherwise.

### GetSpilloverRevenueCentsOk

`func (o *GetAdOffsiteReportResponse) GetSpilloverRevenueCentsOk() (*float32, bool)`

GetSpilloverRevenueCentsOk returns a tuple with the SpilloverRevenueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpilloverRevenueCents

`func (o *GetAdOffsiteReportResponse) SetSpilloverRevenueCents(v float32)`

SetSpilloverRevenueCents sets SpilloverRevenueCents field to given value.


### GetReturnBps

`func (o *GetAdOffsiteReportResponse) GetReturnBps() float32`

GetReturnBps returns the ReturnBps field if non-nil, zero value otherwise.

### GetReturnBpsOk

`func (o *GetAdOffsiteReportResponse) GetReturnBpsOk() (*float32, bool)`

GetReturnBpsOk returns a tuple with the ReturnBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnBps

`func (o *GetAdOffsiteReportResponse) SetReturnBps(v float32)`

SetReturnBps sets ReturnBps field to given value.


### GetEnabled

`func (o *GetAdOffsiteReportResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetAdOffsiteReportResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetAdOffsiteReportResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetAutoPausedAt

`func (o *GetAdOffsiteReportResponse) GetAutoPausedAt() string`

GetAutoPausedAt returns the AutoPausedAt field if non-nil, zero value otherwise.

### GetAutoPausedAtOk

`func (o *GetAdOffsiteReportResponse) GetAutoPausedAtOk() (*string, bool)`

GetAutoPausedAtOk returns a tuple with the AutoPausedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPausedAt

`func (o *GetAdOffsiteReportResponse) SetAutoPausedAt(v string)`

SetAutoPausedAt sets AutoPausedAt field to given value.

### HasAutoPausedAt

`func (o *GetAdOffsiteReportResponse) HasAutoPausedAt() bool`

HasAutoPausedAt returns a boolean if a field has been set.

### SetAutoPausedAtNil

`func (o *GetAdOffsiteReportResponse) SetAutoPausedAtNil(b bool)`

 SetAutoPausedAtNil sets the value for AutoPausedAt to be an explicit nil

### UnsetAutoPausedAt
`func (o *GetAdOffsiteReportResponse) UnsetAutoPausedAt()`

UnsetAutoPausedAt ensures that no value is present for AutoPausedAt, not even an explicit nil
### GetAutoPausedReason

`func (o *GetAdOffsiteReportResponse) GetAutoPausedReason() string`

GetAutoPausedReason returns the AutoPausedReason field if non-nil, zero value otherwise.

### GetAutoPausedReasonOk

`func (o *GetAdOffsiteReportResponse) GetAutoPausedReasonOk() (*string, bool)`

GetAutoPausedReasonOk returns a tuple with the AutoPausedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoPausedReason

`func (o *GetAdOffsiteReportResponse) SetAutoPausedReason(v string)`

SetAutoPausedReason sets AutoPausedReason field to given value.

### HasAutoPausedReason

`func (o *GetAdOffsiteReportResponse) HasAutoPausedReason() bool`

HasAutoPausedReason returns a boolean if a field has been set.

### SetAutoPausedReasonNil

`func (o *GetAdOffsiteReportResponse) SetAutoPausedReasonNil(b bool)`

 SetAutoPausedReasonNil sets the value for AutoPausedReason to be an explicit nil

### UnsetAutoPausedReason
`func (o *GetAdOffsiteReportResponse) UnsetAutoPausedReason()`

UnsetAutoPausedReason ensures that no value is present for AutoPausedReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


