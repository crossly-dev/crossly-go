# GetPayoutGrossForNetResponseEstimate

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

### NewGetPayoutGrossForNetResponseEstimate

`func NewGetPayoutGrossForNetResponseEstimate(platform string, grossCents float32, feeCents float32, feeSource string, shippingCents float32, shippingSource string, netCents float32, takeHomePct float32, assumptions []string, ) *GetPayoutGrossForNetResponseEstimate`

NewGetPayoutGrossForNetResponseEstimate instantiates a new GetPayoutGrossForNetResponseEstimate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayoutGrossForNetResponseEstimateWithDefaults

`func NewGetPayoutGrossForNetResponseEstimateWithDefaults() *GetPayoutGrossForNetResponseEstimate`

NewGetPayoutGrossForNetResponseEstimateWithDefaults instantiates a new GetPayoutGrossForNetResponseEstimate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *GetPayoutGrossForNetResponseEstimate) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetPayoutGrossForNetResponseEstimate) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetPayoutGrossForNetResponseEstimate) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetGrossCents

`func (o *GetPayoutGrossForNetResponseEstimate) GetGrossCents() float32`

GetGrossCents returns the GrossCents field if non-nil, zero value otherwise.

### GetGrossCentsOk

`func (o *GetPayoutGrossForNetResponseEstimate) GetGrossCentsOk() (*float32, bool)`

GetGrossCentsOk returns a tuple with the GrossCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossCents

`func (o *GetPayoutGrossForNetResponseEstimate) SetGrossCents(v float32)`

SetGrossCents sets GrossCents field to given value.


### GetFeeCents

`func (o *GetPayoutGrossForNetResponseEstimate) GetFeeCents() float32`

GetFeeCents returns the FeeCents field if non-nil, zero value otherwise.

### GetFeeCentsOk

`func (o *GetPayoutGrossForNetResponseEstimate) GetFeeCentsOk() (*float32, bool)`

GetFeeCentsOk returns a tuple with the FeeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCents

`func (o *GetPayoutGrossForNetResponseEstimate) SetFeeCents(v float32)`

SetFeeCents sets FeeCents field to given value.


### GetFeeSource

`func (o *GetPayoutGrossForNetResponseEstimate) GetFeeSource() string`

GetFeeSource returns the FeeSource field if non-nil, zero value otherwise.

### GetFeeSourceOk

`func (o *GetPayoutGrossForNetResponseEstimate) GetFeeSourceOk() (*string, bool)`

GetFeeSourceOk returns a tuple with the FeeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeSource

`func (o *GetPayoutGrossForNetResponseEstimate) SetFeeSource(v string)`

SetFeeSource sets FeeSource field to given value.


### GetShippingCents

`func (o *GetPayoutGrossForNetResponseEstimate) GetShippingCents() float32`

GetShippingCents returns the ShippingCents field if non-nil, zero value otherwise.

### GetShippingCentsOk

`func (o *GetPayoutGrossForNetResponseEstimate) GetShippingCentsOk() (*float32, bool)`

GetShippingCentsOk returns a tuple with the ShippingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCents

`func (o *GetPayoutGrossForNetResponseEstimate) SetShippingCents(v float32)`

SetShippingCents sets ShippingCents field to given value.


### GetShippingSource

`func (o *GetPayoutGrossForNetResponseEstimate) GetShippingSource() string`

GetShippingSource returns the ShippingSource field if non-nil, zero value otherwise.

### GetShippingSourceOk

`func (o *GetPayoutGrossForNetResponseEstimate) GetShippingSourceOk() (*string, bool)`

GetShippingSourceOk returns a tuple with the ShippingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingSource

`func (o *GetPayoutGrossForNetResponseEstimate) SetShippingSource(v string)`

SetShippingSource sets ShippingSource field to given value.


### GetNetCents

`func (o *GetPayoutGrossForNetResponseEstimate) GetNetCents() float32`

GetNetCents returns the NetCents field if non-nil, zero value otherwise.

### GetNetCentsOk

`func (o *GetPayoutGrossForNetResponseEstimate) GetNetCentsOk() (*float32, bool)`

GetNetCentsOk returns a tuple with the NetCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCents

`func (o *GetPayoutGrossForNetResponseEstimate) SetNetCents(v float32)`

SetNetCents sets NetCents field to given value.


### GetTakeHomePct

`func (o *GetPayoutGrossForNetResponseEstimate) GetTakeHomePct() float32`

GetTakeHomePct returns the TakeHomePct field if non-nil, zero value otherwise.

### GetTakeHomePctOk

`func (o *GetPayoutGrossForNetResponseEstimate) GetTakeHomePctOk() (*float32, bool)`

GetTakeHomePctOk returns a tuple with the TakeHomePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeHomePct

`func (o *GetPayoutGrossForNetResponseEstimate) SetTakeHomePct(v float32)`

SetTakeHomePct sets TakeHomePct field to given value.


### GetAssumptions

`func (o *GetPayoutGrossForNetResponseEstimate) GetAssumptions() []string`

GetAssumptions returns the Assumptions field if non-nil, zero value otherwise.

### GetAssumptionsOk

`func (o *GetPayoutGrossForNetResponseEstimate) GetAssumptionsOk() (*[]string, bool)`

GetAssumptionsOk returns a tuple with the Assumptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumptions

`func (o *GetPayoutGrossForNetResponseEstimate) SetAssumptions(v []string)`

SetAssumptions sets Assumptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


