# CreateMagicScanResponseTopHits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** |  | 
**Title** | **string** |  | 
**PriceCents** | **float32** |  | 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**ListingUrl** | Pointer to **NullableString** |  | [optional] 
**Origin** | **string** | Origin marker for the UI badge. | 
**VisualSim** | Pointer to **NullableFloat32** | CLIP cosine [0,1]; populated after the visual-rank pass. | [optional] 
**State** | **string** | &#39;active&#39; &#x3D; currently for sale; &#39;sold&#39; &#x3D; historical comp. UI  renders distinct badges so the seller can see both at a glance. | 

## Methods

### NewCreateMagicScanResponseTopHits

`func NewCreateMagicScanResponseTopHits(platform string, title string, priceCents float32, origin string, state string, ) *CreateMagicScanResponseTopHits`

NewCreateMagicScanResponseTopHits instantiates a new CreateMagicScanResponseTopHits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMagicScanResponseTopHitsWithDefaults

`func NewCreateMagicScanResponseTopHitsWithDefaults() *CreateMagicScanResponseTopHits`

NewCreateMagicScanResponseTopHitsWithDefaults instantiates a new CreateMagicScanResponseTopHits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *CreateMagicScanResponseTopHits) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateMagicScanResponseTopHits) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateMagicScanResponseTopHits) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetTitle

`func (o *CreateMagicScanResponseTopHits) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateMagicScanResponseTopHits) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateMagicScanResponseTopHits) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPriceCents

`func (o *CreateMagicScanResponseTopHits) GetPriceCents() float32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *CreateMagicScanResponseTopHits) GetPriceCentsOk() (*float32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *CreateMagicScanResponseTopHits) SetPriceCents(v float32)`

SetPriceCents sets PriceCents field to given value.


### GetImageUrl

`func (o *CreateMagicScanResponseTopHits) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateMagicScanResponseTopHits) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateMagicScanResponseTopHits) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateMagicScanResponseTopHits) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *CreateMagicScanResponseTopHits) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *CreateMagicScanResponseTopHits) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetListingUrl

`func (o *CreateMagicScanResponseTopHits) GetListingUrl() string`

GetListingUrl returns the ListingUrl field if non-nil, zero value otherwise.

### GetListingUrlOk

`func (o *CreateMagicScanResponseTopHits) GetListingUrlOk() (*string, bool)`

GetListingUrlOk returns a tuple with the ListingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingUrl

`func (o *CreateMagicScanResponseTopHits) SetListingUrl(v string)`

SetListingUrl sets ListingUrl field to given value.

### HasListingUrl

`func (o *CreateMagicScanResponseTopHits) HasListingUrl() bool`

HasListingUrl returns a boolean if a field has been set.

### SetListingUrlNil

`func (o *CreateMagicScanResponseTopHits) SetListingUrlNil(b bool)`

 SetListingUrlNil sets the value for ListingUrl to be an explicit nil

### UnsetListingUrl
`func (o *CreateMagicScanResponseTopHits) UnsetListingUrl()`

UnsetListingUrl ensures that no value is present for ListingUrl, not even an explicit nil
### GetOrigin

`func (o *CreateMagicScanResponseTopHits) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *CreateMagicScanResponseTopHits) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *CreateMagicScanResponseTopHits) SetOrigin(v string)`

SetOrigin sets Origin field to given value.


### GetVisualSim

`func (o *CreateMagicScanResponseTopHits) GetVisualSim() float32`

GetVisualSim returns the VisualSim field if non-nil, zero value otherwise.

### GetVisualSimOk

`func (o *CreateMagicScanResponseTopHits) GetVisualSimOk() (*float32, bool)`

GetVisualSimOk returns a tuple with the VisualSim field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisualSim

`func (o *CreateMagicScanResponseTopHits) SetVisualSim(v float32)`

SetVisualSim sets VisualSim field to given value.

### HasVisualSim

`func (o *CreateMagicScanResponseTopHits) HasVisualSim() bool`

HasVisualSim returns a boolean if a field has been set.

### SetVisualSimNil

`func (o *CreateMagicScanResponseTopHits) SetVisualSimNil(b bool)`

 SetVisualSimNil sets the value for VisualSim to be an explicit nil

### UnsetVisualSim
`func (o *CreateMagicScanResponseTopHits) UnsetVisualSim()`

UnsetVisualSim ensures that no value is present for VisualSim, not even an explicit nil
### GetState

`func (o *CreateMagicScanResponseTopHits) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *CreateMagicScanResponseTopHits) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *CreateMagicScanResponseTopHits) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


