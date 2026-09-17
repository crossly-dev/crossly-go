# ListSpatialPublicItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**PublicSlug** | **string** |  | 
**CategorySlug** | **string** |  | 
**ItemCount** | **float32** | Unsold stock that lands in this room. The same predicate the room uses. | 
**ForSaleCount** | **float32** | How many of those a visitor could buy right now. | 
**PriceFromCents** | Pointer to **NullableFloat32** | The cheapest and dearest thing for sale, in cents.  A BAND, deliberately, and never a quote: &#x60;listPublicSceneOffers&#x60; is the only authority on what a given object costs. Null when nothing is for sale — zero would read as free. | [optional] 
**PriceToCents** | Pointer to **NullableFloat32** |  | [optional] 
**PreviewImages** | **[]string** | Up to PREVIEW_IMAGES item images. Catalog art first, seller photo else. | 
**UpdatedAt** | Pointer to **NullableString** | Last time the room itself changed. ISO, or null if the row has no date. | [optional] 

## Methods

### NewListSpatialPublicItem

`func NewListSpatialPublicItem(name string, publicSlug string, categorySlug string, itemCount float32, forSaleCount float32, previewImages []string, ) *ListSpatialPublicItem`

NewListSpatialPublicItem instantiates a new ListSpatialPublicItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSpatialPublicItemWithDefaults

`func NewListSpatialPublicItemWithDefaults() *ListSpatialPublicItem`

NewListSpatialPublicItemWithDefaults instantiates a new ListSpatialPublicItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ListSpatialPublicItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSpatialPublicItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSpatialPublicItem) SetName(v string)`

SetName sets Name field to given value.


### GetPublicSlug

`func (o *ListSpatialPublicItem) GetPublicSlug() string`

GetPublicSlug returns the PublicSlug field if non-nil, zero value otherwise.

### GetPublicSlugOk

`func (o *ListSpatialPublicItem) GetPublicSlugOk() (*string, bool)`

GetPublicSlugOk returns a tuple with the PublicSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSlug

`func (o *ListSpatialPublicItem) SetPublicSlug(v string)`

SetPublicSlug sets PublicSlug field to given value.


### GetCategorySlug

`func (o *ListSpatialPublicItem) GetCategorySlug() string`

GetCategorySlug returns the CategorySlug field if non-nil, zero value otherwise.

### GetCategorySlugOk

`func (o *ListSpatialPublicItem) GetCategorySlugOk() (*string, bool)`

GetCategorySlugOk returns a tuple with the CategorySlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySlug

`func (o *ListSpatialPublicItem) SetCategorySlug(v string)`

SetCategorySlug sets CategorySlug field to given value.


### GetItemCount

`func (o *ListSpatialPublicItem) GetItemCount() float32`

GetItemCount returns the ItemCount field if non-nil, zero value otherwise.

### GetItemCountOk

`func (o *ListSpatialPublicItem) GetItemCountOk() (*float32, bool)`

GetItemCountOk returns a tuple with the ItemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemCount

`func (o *ListSpatialPublicItem) SetItemCount(v float32)`

SetItemCount sets ItemCount field to given value.


### GetForSaleCount

`func (o *ListSpatialPublicItem) GetForSaleCount() float32`

GetForSaleCount returns the ForSaleCount field if non-nil, zero value otherwise.

### GetForSaleCountOk

`func (o *ListSpatialPublicItem) GetForSaleCountOk() (*float32, bool)`

GetForSaleCountOk returns a tuple with the ForSaleCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForSaleCount

`func (o *ListSpatialPublicItem) SetForSaleCount(v float32)`

SetForSaleCount sets ForSaleCount field to given value.


### GetPriceFromCents

`func (o *ListSpatialPublicItem) GetPriceFromCents() float32`

GetPriceFromCents returns the PriceFromCents field if non-nil, zero value otherwise.

### GetPriceFromCentsOk

`func (o *ListSpatialPublicItem) GetPriceFromCentsOk() (*float32, bool)`

GetPriceFromCentsOk returns a tuple with the PriceFromCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFromCents

`func (o *ListSpatialPublicItem) SetPriceFromCents(v float32)`

SetPriceFromCents sets PriceFromCents field to given value.

### HasPriceFromCents

`func (o *ListSpatialPublicItem) HasPriceFromCents() bool`

HasPriceFromCents returns a boolean if a field has been set.

### SetPriceFromCentsNil

`func (o *ListSpatialPublicItem) SetPriceFromCentsNil(b bool)`

 SetPriceFromCentsNil sets the value for PriceFromCents to be an explicit nil

### UnsetPriceFromCents
`func (o *ListSpatialPublicItem) UnsetPriceFromCents()`

UnsetPriceFromCents ensures that no value is present for PriceFromCents, not even an explicit nil
### GetPriceToCents

`func (o *ListSpatialPublicItem) GetPriceToCents() float32`

GetPriceToCents returns the PriceToCents field if non-nil, zero value otherwise.

### GetPriceToCentsOk

`func (o *ListSpatialPublicItem) GetPriceToCentsOk() (*float32, bool)`

GetPriceToCentsOk returns a tuple with the PriceToCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceToCents

`func (o *ListSpatialPublicItem) SetPriceToCents(v float32)`

SetPriceToCents sets PriceToCents field to given value.

### HasPriceToCents

`func (o *ListSpatialPublicItem) HasPriceToCents() bool`

HasPriceToCents returns a boolean if a field has been set.

### SetPriceToCentsNil

`func (o *ListSpatialPublicItem) SetPriceToCentsNil(b bool)`

 SetPriceToCentsNil sets the value for PriceToCents to be an explicit nil

### UnsetPriceToCents
`func (o *ListSpatialPublicItem) UnsetPriceToCents()`

UnsetPriceToCents ensures that no value is present for PriceToCents, not even an explicit nil
### GetPreviewImages

`func (o *ListSpatialPublicItem) GetPreviewImages() []string`

GetPreviewImages returns the PreviewImages field if non-nil, zero value otherwise.

### GetPreviewImagesOk

`func (o *ListSpatialPublicItem) GetPreviewImagesOk() (*[]string, bool)`

GetPreviewImagesOk returns a tuple with the PreviewImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewImages

`func (o *ListSpatialPublicItem) SetPreviewImages(v []string)`

SetPreviewImages sets PreviewImages field to given value.


### GetUpdatedAt

`func (o *ListSpatialPublicItem) GetUpdatedAt() string`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListSpatialPublicItem) GetUpdatedAtOk() (*string, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListSpatialPublicItem) SetUpdatedAt(v string)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListSpatialPublicItem) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### SetUpdatedAtNil

`func (o *ListSpatialPublicItem) SetUpdatedAtNil(b bool)`

 SetUpdatedAtNil sets the value for UpdatedAt to be an explicit nil

### UnsetUpdatedAt
`func (o *ListSpatialPublicItem) UnsetUpdatedAt()`

UnsetUpdatedAt ensures that no value is present for UpdatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


