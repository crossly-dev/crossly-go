# GetPayoutEstimateResponse

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

### NewGetPayoutEstimateResponse

`func NewGetPayoutEstimateResponse(platform string, grossCents float32, feeCents float32, feeSource string, shippingCents float32, shippingSource string, netCents float32, takeHomePct float32, assumptions []string, ) *GetPayoutEstimateResponse`

NewGetPayoutEstimateResponse instantiates a new GetPayoutEstimateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayoutEstimateResponseWithDefaults

`func NewGetPayoutEstimateResponseWithDefaults() *GetPayoutEstimateResponse`

NewGetPayoutEstimateResponseWithDefaults instantiates a new GetPayoutEstimateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *GetPayoutEstimateResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetPayoutEstimateResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetPayoutEstimateResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetGrossCents

`func (o *GetPayoutEstimateResponse) GetGrossCents() float32`

GetGrossCents returns the GrossCents field if non-nil, zero value otherwise.

### GetGrossCentsOk

`func (o *GetPayoutEstimateResponse) GetGrossCentsOk() (*float32, bool)`

GetGrossCentsOk returns a tuple with the GrossCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossCents

`func (o *GetPayoutEstimateResponse) SetGrossCents(v float32)`

SetGrossCents sets GrossCents field to given value.


### GetFeeCents

`func (o *GetPayoutEstimateResponse) GetFeeCents() float32`

GetFeeCents returns the FeeCents field if non-nil, zero value otherwise.

### GetFeeCentsOk

`func (o *GetPayoutEstimateResponse) GetFeeCentsOk() (*float32, bool)`

GetFeeCentsOk returns a tuple with the FeeCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeCents

`func (o *GetPayoutEstimateResponse) SetFeeCents(v float32)`

SetFeeCents sets FeeCents field to given value.


### GetFeeSource

`func (o *GetPayoutEstimateResponse) GetFeeSource() string`

GetFeeSource returns the FeeSource field if non-nil, zero value otherwise.

### GetFeeSourceOk

`func (o *GetPayoutEstimateResponse) GetFeeSourceOk() (*string, bool)`

GetFeeSourceOk returns a tuple with the FeeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeSource

`func (o *GetPayoutEstimateResponse) SetFeeSource(v string)`

SetFeeSource sets FeeSource field to given value.


### GetShippingCents

`func (o *GetPayoutEstimateResponse) GetShippingCents() float32`

GetShippingCents returns the ShippingCents field if non-nil, zero value otherwise.

### GetShippingCentsOk

`func (o *GetPayoutEstimateResponse) GetShippingCentsOk() (*float32, bool)`

GetShippingCentsOk returns a tuple with the ShippingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCents

`func (o *GetPayoutEstimateResponse) SetShippingCents(v float32)`

SetShippingCents sets ShippingCents field to given value.


### GetShippingSource

`func (o *GetPayoutEstimateResponse) GetShippingSource() string`

GetShippingSource returns the ShippingSource field if non-nil, zero value otherwise.

### GetShippingSourceOk

`func (o *GetPayoutEstimateResponse) GetShippingSourceOk() (*string, bool)`

GetShippingSourceOk returns a tuple with the ShippingSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingSource

`func (o *GetPayoutEstimateResponse) SetShippingSource(v string)`

SetShippingSource sets ShippingSource field to given value.


### GetNetCents

`func (o *GetPayoutEstimateResponse) GetNetCents() float32`

GetNetCents returns the NetCents field if non-nil, zero value otherwise.

### GetNetCentsOk

`func (o *GetPayoutEstimateResponse) GetNetCentsOk() (*float32, bool)`

GetNetCentsOk returns a tuple with the NetCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetCents

`func (o *GetPayoutEstimateResponse) SetNetCents(v float32)`

SetNetCents sets NetCents field to given value.


### GetTakeHomePct

`func (o *GetPayoutEstimateResponse) GetTakeHomePct() float32`

GetTakeHomePct returns the TakeHomePct field if non-nil, zero value otherwise.

### GetTakeHomePctOk

`func (o *GetPayoutEstimateResponse) GetTakeHomePctOk() (*float32, bool)`

GetTakeHomePctOk returns a tuple with the TakeHomePct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakeHomePct

`func (o *GetPayoutEstimateResponse) SetTakeHomePct(v float32)`

SetTakeHomePct sets TakeHomePct field to given value.


### GetAssumptions

`func (o *GetPayoutEstimateResponse) GetAssumptions() []string`

GetAssumptions returns the Assumptions field if non-nil, zero value otherwise.

### GetAssumptionsOk

`func (o *GetPayoutEstimateResponse) GetAssumptionsOk() (*[]string, bool)`

GetAssumptionsOk returns a tuple with the Assumptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssumptions

`func (o *GetPayoutEstimateResponse) SetAssumptions(v []string)`

SetAssumptions sets Assumptions field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


