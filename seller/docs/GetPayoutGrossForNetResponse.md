# GetPayoutGrossForNetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrossCents** | **float32** |  | 
**Estimate** | [**GetPayoutGrossForNetResponseEstimate**](GetPayoutGrossForNetResponseEstimate.md) |  | 

## Methods

### NewGetPayoutGrossForNetResponse

`func NewGetPayoutGrossForNetResponse(grossCents float32, estimate GetPayoutGrossForNetResponseEstimate, ) *GetPayoutGrossForNetResponse`

NewGetPayoutGrossForNetResponse instantiates a new GetPayoutGrossForNetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPayoutGrossForNetResponseWithDefaults

`func NewGetPayoutGrossForNetResponseWithDefaults() *GetPayoutGrossForNetResponse`

NewGetPayoutGrossForNetResponseWithDefaults instantiates a new GetPayoutGrossForNetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrossCents

`func (o *GetPayoutGrossForNetResponse) GetGrossCents() float32`

GetGrossCents returns the GrossCents field if non-nil, zero value otherwise.

### GetGrossCentsOk

`func (o *GetPayoutGrossForNetResponse) GetGrossCentsOk() (*float32, bool)`

GetGrossCentsOk returns a tuple with the GrossCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossCents

`func (o *GetPayoutGrossForNetResponse) SetGrossCents(v float32)`

SetGrossCents sets GrossCents field to given value.


### GetEstimate

`func (o *GetPayoutGrossForNetResponse) GetEstimate() GetPayoutGrossForNetResponseEstimate`

GetEstimate returns the Estimate field if non-nil, zero value otherwise.

### GetEstimateOk

`func (o *GetPayoutGrossForNetResponse) GetEstimateOk() (*GetPayoutGrossForNetResponseEstimate, bool)`

GetEstimateOk returns a tuple with the Estimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimate

`func (o *GetPayoutGrossForNetResponse) SetEstimate(v GetPayoutGrossForNetResponseEstimate)`

SetEstimate sets Estimate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


