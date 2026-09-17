# GetSpatialSceneResponseItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UnitId** | Pointer to **NullableString** |  | [optional] 
**Title** | **string** |  | 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**ThumbUrl** | Pointer to **NullableString** | A small copy of &#x60;imageUrl&#x60;, when the catalog has one.    The room binds this for everything except the few items you are standing  in front of. A 600x600 original costs 1.83 MiB of VRAM; a 150x150 thumb  costs 0.11 MiB, and at more than a couple of metres they are the same  handful of pixels on screen. Null when the catalog never made one, which  the renderer treats as \&quot;use the original\&quot; rather than as \&quot;draw nothing\&quot;. | [optional] 
**Size** | Pointer to [**NullableGetSpatialSceneResponseSize**](GetSpatialSceneResponseSize.md) |  | [optional] 
**CostCents** | Pointer to **NullableFloat32** | Cents the seller paid. Drives the &#x60;value&#x60; overlay and capital density. | [optional] 
**AgeDays** | Pointer to **NullableFloat32** | Days since first listed anywhere. Drives the &#x60;aging&#x60; overlay. | [optional] 
**Status** | **string** |  | 
**Location** | Pointer to **NullableString** | Free-text or structured location, when the unit has one. | [optional] 
**MarketTagged** | **bool** | True when this row resolved to a catalog product. | 

## Methods

### NewGetSpatialSceneResponseItems

`func NewGetSpatialSceneResponseItems(id string, title string, status string, marketTagged bool, ) *GetSpatialSceneResponseItems`

NewGetSpatialSceneResponseItems instantiates a new GetSpatialSceneResponseItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSpatialSceneResponseItemsWithDefaults

`func NewGetSpatialSceneResponseItemsWithDefaults() *GetSpatialSceneResponseItems`

NewGetSpatialSceneResponseItemsWithDefaults instantiates a new GetSpatialSceneResponseItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSpatialSceneResponseItems) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSpatialSceneResponseItems) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSpatialSceneResponseItems) SetId(v string)`

SetId sets Id field to given value.


### GetUnitId

`func (o *GetSpatialSceneResponseItems) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *GetSpatialSceneResponseItems) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *GetSpatialSceneResponseItems) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *GetSpatialSceneResponseItems) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *GetSpatialSceneResponseItems) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *GetSpatialSceneResponseItems) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetTitle

`func (o *GetSpatialSceneResponseItems) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetSpatialSceneResponseItems) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetSpatialSceneResponseItems) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetImageUrl

`func (o *GetSpatialSceneResponseItems) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GetSpatialSceneResponseItems) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GetSpatialSceneResponseItems) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GetSpatialSceneResponseItems) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *GetSpatialSceneResponseItems) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *GetSpatialSceneResponseItems) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetThumbUrl

`func (o *GetSpatialSceneResponseItems) GetThumbUrl() string`

GetThumbUrl returns the ThumbUrl field if non-nil, zero value otherwise.

### GetThumbUrlOk

`func (o *GetSpatialSceneResponseItems) GetThumbUrlOk() (*string, bool)`

GetThumbUrlOk returns a tuple with the ThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbUrl

`func (o *GetSpatialSceneResponseItems) SetThumbUrl(v string)`

SetThumbUrl sets ThumbUrl field to given value.

### HasThumbUrl

`func (o *GetSpatialSceneResponseItems) HasThumbUrl() bool`

HasThumbUrl returns a boolean if a field has been set.

### SetThumbUrlNil

`func (o *GetSpatialSceneResponseItems) SetThumbUrlNil(b bool)`

 SetThumbUrlNil sets the value for ThumbUrl to be an explicit nil

### UnsetThumbUrl
`func (o *GetSpatialSceneResponseItems) UnsetThumbUrl()`

UnsetThumbUrl ensures that no value is present for ThumbUrl, not even an explicit nil
### GetSize

`func (o *GetSpatialSceneResponseItems) GetSize() GetSpatialSceneResponseSize`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetSpatialSceneResponseItems) GetSizeOk() (*GetSpatialSceneResponseSize, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetSpatialSceneResponseItems) SetSize(v GetSpatialSceneResponseSize)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetSpatialSceneResponseItems) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *GetSpatialSceneResponseItems) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *GetSpatialSceneResponseItems) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetCostCents

`func (o *GetSpatialSceneResponseItems) GetCostCents() float32`

GetCostCents returns the CostCents field if non-nil, zero value otherwise.

### GetCostCentsOk

`func (o *GetSpatialSceneResponseItems) GetCostCentsOk() (*float32, bool)`

GetCostCentsOk returns a tuple with the CostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostCents

`func (o *GetSpatialSceneResponseItems) SetCostCents(v float32)`

SetCostCents sets CostCents field to given value.

### HasCostCents

`func (o *GetSpatialSceneResponseItems) HasCostCents() bool`

HasCostCents returns a boolean if a field has been set.

### SetCostCentsNil

`func (o *GetSpatialSceneResponseItems) SetCostCentsNil(b bool)`

 SetCostCentsNil sets the value for CostCents to be an explicit nil

### UnsetCostCents
`func (o *GetSpatialSceneResponseItems) UnsetCostCents()`

UnsetCostCents ensures that no value is present for CostCents, not even an explicit nil
### GetAgeDays

`func (o *GetSpatialSceneResponseItems) GetAgeDays() float32`

GetAgeDays returns the AgeDays field if non-nil, zero value otherwise.

### GetAgeDaysOk

`func (o *GetSpatialSceneResponseItems) GetAgeDaysOk() (*float32, bool)`

GetAgeDaysOk returns a tuple with the AgeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgeDays

`func (o *GetSpatialSceneResponseItems) SetAgeDays(v float32)`

SetAgeDays sets AgeDays field to given value.

### HasAgeDays

`func (o *GetSpatialSceneResponseItems) HasAgeDays() bool`

HasAgeDays returns a boolean if a field has been set.

### SetAgeDaysNil

`func (o *GetSpatialSceneResponseItems) SetAgeDaysNil(b bool)`

 SetAgeDaysNil sets the value for AgeDays to be an explicit nil

### UnsetAgeDays
`func (o *GetSpatialSceneResponseItems) UnsetAgeDays()`

UnsetAgeDays ensures that no value is present for AgeDays, not even an explicit nil
### GetStatus

`func (o *GetSpatialSceneResponseItems) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSpatialSceneResponseItems) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSpatialSceneResponseItems) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetLocation

`func (o *GetSpatialSceneResponseItems) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GetSpatialSceneResponseItems) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GetSpatialSceneResponseItems) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GetSpatialSceneResponseItems) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### SetLocationNil

`func (o *GetSpatialSceneResponseItems) SetLocationNil(b bool)`

 SetLocationNil sets the value for Location to be an explicit nil

### UnsetLocation
`func (o *GetSpatialSceneResponseItems) UnsetLocation()`

UnsetLocation ensures that no value is present for Location, not even an explicit nil
### GetMarketTagged

`func (o *GetSpatialSceneResponseItems) GetMarketTagged() bool`

GetMarketTagged returns the MarketTagged field if non-nil, zero value otherwise.

### GetMarketTaggedOk

`func (o *GetSpatialSceneResponseItems) GetMarketTaggedOk() (*bool, bool)`

GetMarketTaggedOk returns a tuple with the MarketTagged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketTagged

`func (o *GetSpatialSceneResponseItems) SetMarketTagged(v bool)`

SetMarketTagged sets MarketTagged field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


