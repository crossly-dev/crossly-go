# GetBuyerPreferenceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Retailers** | [**[]GetBuyerPreferenceResponseRetailers**](GetBuyerPreferenceResponseRetailers.md) | Retailers they shop, most-seen first. | 
**Brands** | [**[]GetBuyerPreferenceResponseBrands**](GetBuyerPreferenceResponseBrands.md) | What they look at, by catalog brand where we could resolve one. | 
**PriceBand** | Pointer to [**NullableGetBuyerPreferenceResponsePriceBand**](GetBuyerPreferenceResponsePriceBand.md) |  | [optional] 
**MatchRate** | Pointer to **NullableFloat32** | How much of what they want we can actually supply. | [optional] 
**Observations** | **float32** |  | 

## Methods

### NewGetBuyerPreferenceResponse

`func NewGetBuyerPreferenceResponse(retailers []GetBuyerPreferenceResponseRetailers, brands []GetBuyerPreferenceResponseBrands, observations float32, ) *GetBuyerPreferenceResponse`

NewGetBuyerPreferenceResponse instantiates a new GetBuyerPreferenceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBuyerPreferenceResponseWithDefaults

`func NewGetBuyerPreferenceResponseWithDefaults() *GetBuyerPreferenceResponse`

NewGetBuyerPreferenceResponseWithDefaults instantiates a new GetBuyerPreferenceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRetailers

`func (o *GetBuyerPreferenceResponse) GetRetailers() []GetBuyerPreferenceResponseRetailers`

GetRetailers returns the Retailers field if non-nil, zero value otherwise.

### GetRetailersOk

`func (o *GetBuyerPreferenceResponse) GetRetailersOk() (*[]GetBuyerPreferenceResponseRetailers, bool)`

GetRetailersOk returns a tuple with the Retailers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailers

`func (o *GetBuyerPreferenceResponse) SetRetailers(v []GetBuyerPreferenceResponseRetailers)`

SetRetailers sets Retailers field to given value.


### GetBrands

`func (o *GetBuyerPreferenceResponse) GetBrands() []GetBuyerPreferenceResponseBrands`

GetBrands returns the Brands field if non-nil, zero value otherwise.

### GetBrandsOk

`func (o *GetBuyerPreferenceResponse) GetBrandsOk() (*[]GetBuyerPreferenceResponseBrands, bool)`

GetBrandsOk returns a tuple with the Brands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrands

`func (o *GetBuyerPreferenceResponse) SetBrands(v []GetBuyerPreferenceResponseBrands)`

SetBrands sets Brands field to given value.


### GetPriceBand

`func (o *GetBuyerPreferenceResponse) GetPriceBand() GetBuyerPreferenceResponsePriceBand`

GetPriceBand returns the PriceBand field if non-nil, zero value otherwise.

### GetPriceBandOk

`func (o *GetBuyerPreferenceResponse) GetPriceBandOk() (*GetBuyerPreferenceResponsePriceBand, bool)`

GetPriceBandOk returns a tuple with the PriceBand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceBand

`func (o *GetBuyerPreferenceResponse) SetPriceBand(v GetBuyerPreferenceResponsePriceBand)`

SetPriceBand sets PriceBand field to given value.

### HasPriceBand

`func (o *GetBuyerPreferenceResponse) HasPriceBand() bool`

HasPriceBand returns a boolean if a field has been set.

### SetPriceBandNil

`func (o *GetBuyerPreferenceResponse) SetPriceBandNil(b bool)`

 SetPriceBandNil sets the value for PriceBand to be an explicit nil

### UnsetPriceBand
`func (o *GetBuyerPreferenceResponse) UnsetPriceBand()`

UnsetPriceBand ensures that no value is present for PriceBand, not even an explicit nil
### GetMatchRate

`func (o *GetBuyerPreferenceResponse) GetMatchRate() float32`

GetMatchRate returns the MatchRate field if non-nil, zero value otherwise.

### GetMatchRateOk

`func (o *GetBuyerPreferenceResponse) GetMatchRateOk() (*float32, bool)`

GetMatchRateOk returns a tuple with the MatchRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchRate

`func (o *GetBuyerPreferenceResponse) SetMatchRate(v float32)`

SetMatchRate sets MatchRate field to given value.

### HasMatchRate

`func (o *GetBuyerPreferenceResponse) HasMatchRate() bool`

HasMatchRate returns a boolean if a field has been set.

### SetMatchRateNil

`func (o *GetBuyerPreferenceResponse) SetMatchRateNil(b bool)`

 SetMatchRateNil sets the value for MatchRate to be an explicit nil

### UnsetMatchRate
`func (o *GetBuyerPreferenceResponse) UnsetMatchRate()`

UnsetMatchRate ensures that no value is present for MatchRate, not even an explicit nil
### GetObservations

`func (o *GetBuyerPreferenceResponse) GetObservations() float32`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *GetBuyerPreferenceResponse) GetObservationsOk() (*float32, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *GetBuyerPreferenceResponse) SetObservations(v float32)`

SetObservations sets Observations field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


