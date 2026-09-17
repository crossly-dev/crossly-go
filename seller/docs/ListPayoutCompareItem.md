# ListPayoutCompareItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** |  | 
**GrossCents** | **float32** |  | 
**FeeCents** | **float32** |  | 
**FeeSource** | **string** |  | 
**ShippingCents** | **float32** |  | 
**ShippingSource** | **string** |  | 
**NetCents** | **float32** |  | 
**TakeHomePct** | **float32** | Net as a percentage of gross, for comparing across price points. | 
**Assumptions** | **[]string** | What we assumed, in the seller&#39;s words. Never empty when we guessed. | 

## Methods

### NewListPayoutCompareItem

`func NewListPayoutCompareItem(platform string, grossCents float32, feeCents float32, feeSource string, shippingCents float32, shippingSource string, netCents float32, takeHomePct float32, assumptions []string, ) *ListPayoutCompareItem`

NewListPayoutCompareItem instantiates a new ListPayoutCompareItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPayoutCompareItemWithDefaults

`func NewListPayoutCompareItemWithDefaults() *ListPayoutCompareItem`

NewListPayoutCompareItemWithDefaults instantiates a new ListPayoutCompareItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *ListPayoutCompareItem) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListPayoutCompareItem) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListPayoutCompareItem) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetGrossCents

`func (o *ListPayoutCompareItem) GetGrossCents() float32`

GetGrossCents returns the GrossCents field if non-nil, zero value otherwise.

### GetGrossCentsOk

`func (o *ListPayoutCompareItem) GetGrossCentsOk() (*float32, bool)`

GetGrossCentsOk returns a tuple with the GrossCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossCents

`func (o *ListPayoutCompareItem) SetGrossCents(v float32)`

SetGrossCents sets GrossCents field to given value.


### GetFeeCents

`func (o *ListPayoutCompareItem) GetFeeCents() float32`

GetFeeCents returns the FeeCents field if non-nil, zero value otherwise.

### GetFeeCentsOk

`func (o *ListPayoutCompareItem) GetFeeCentsOk() (*float32, bool)`

GetFeeCentsOk returns a tuple with the FeeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCents

`func (o *ListPayoutCompareItem) SetFeeCents(v float32)`

SetFeeCents sets FeeCents field to given value.


### GetFeeSource

`func (o *ListPayoutCompareItem) GetFeeSource() string`

GetFeeSource returns the FeeSource field if non-nil, zero value otherwise.

### GetFeeSourceOk

`func (o *ListPayoutCompareItem) GetFeeSourceOk() (*string, bool)`

GetFeeSourceOk returns a tuple with the FeeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeSource

`func (o *ListPayoutCompareItem) SetFeeSource(v string)`

SetFeeSource sets FeeSource field to given value.


### GetShippingCents

`func (o *ListPayoutCompareItem) GetShippingCents() float32`

GetShippingCents returns the ShippingCents field if non-nil, zero value otherwise.

### GetShippingCentsOk

`func (o *ListPayoutCompareItem) GetShippingCentsOk() (*float32, bool)`

GetShippingCentsOk returns a tuple with the ShippingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCents

`func (o *ListPayoutCompareItem) SetShippingCents(v float32)`

SetShippingCents sets ShippingCents field to given value.


### GetShippingSource

`func (o *ListPayoutCompareItem) GetShippingSource() string`

GetShippingSource returns the ShippingSource field if non-nil, zero value otherwise.

### GetShippingSourceOk

`func (o *ListPayoutCompareItem) GetShippingSourceOk() (*string, bool)`

GetShippingSourceOk returns a tuple with the ShippingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingSource

`func (o *ListPayoutCompareItem) SetShippingSource(v string)`

SetShippingSource sets ShippingSource field to given value.


### GetNetCents

`func (o *ListPayoutCompareItem) GetNetCents() float32`

GetNetCents returns the NetCents field if non-nil, zero value otherwise.

### GetNetCentsOk

`func (o *ListPayoutCompareItem) GetNetCentsOk() (*float32, bool)`

GetNetCentsOk returns a tuple with the NetCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCents

`func (o *ListPayoutCompareItem) SetNetCents(v float32)`

SetNetCents sets NetCents field to given value.


### GetTakeHomePct

`func (o *ListPayoutCompareItem) GetTakeHomePct() float32`

GetTakeHomePct returns the TakeHomePct field if non-nil, zero value otherwise.

### GetTakeHomePctOk

`func (o *ListPayoutCompareItem) GetTakeHomePctOk() (*float32, bool)`

GetTakeHomePctOk returns a tuple with the TakeHomePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeHomePct

`func (o *ListPayoutCompareItem) SetTakeHomePct(v float32)`

SetTakeHomePct sets TakeHomePct field to given value.


### GetAssumptions

`func (o *ListPayoutCompareItem) GetAssumptions() []string`

GetAssumptions returns the Assumptions field if non-nil, zero value otherwise.

### GetAssumptionsOk

`func (o *ListPayoutCompareItem) GetAssumptionsOk() (*[]string, bool)`

GetAssumptionsOk returns a tuple with the Assumptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumptions

`func (o *ListPayoutCompareItem) SetAssumptions(v []string)`

SetAssumptions sets Assumptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


