# ListSpatialPublicOffersItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | **string** |  | 
**ListingSlug** | **string** | platform_listings.platform_listing_id — what POST /cart/items takes. | 
**PriceCents** | **float32** |  | 
**Currency** | **string** |  | 
**Available** | Pointer to **NullableFloat32** | Remaining stock, when the listing declares one. Null &#x3D; unknown. | [optional] 
**Condition** | Pointer to **NullableString** | What a buyer is entitled to know before they add it. | [optional] 
**GradeKey** | Pointer to **NullableString** | Canonical grade key when the item is slabbed, e.g. &#x60;psa-10&#x60;. | [optional] 

## Methods

### NewListSpatialPublicOffersItem

`func NewListSpatialPublicOffersItem(itemId string, listingSlug string, priceCents float32, currency string, ) *ListSpatialPublicOffersItem`

NewListSpatialPublicOffersItem instantiates a new ListSpatialPublicOffersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSpatialPublicOffersItemWithDefaults

`func NewListSpatialPublicOffersItemWithDefaults() *ListSpatialPublicOffersItem`

NewListSpatialPublicOffersItemWithDefaults instantiates a new ListSpatialPublicOffersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *ListSpatialPublicOffersItem) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ListSpatialPublicOffersItem) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ListSpatialPublicOffersItem) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetListingSlug

`func (o *ListSpatialPublicOffersItem) GetListingSlug() string`

GetListingSlug returns the ListingSlug field if non-nil, zero value otherwise.

### GetListingSlugOk

`func (o *ListSpatialPublicOffersItem) GetListingSlugOk() (*string, bool)`

GetListingSlugOk returns a tuple with the ListingSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingSlug

`func (o *ListSpatialPublicOffersItem) SetListingSlug(v string)`

SetListingSlug sets ListingSlug field to given value.


### GetPriceCents

`func (o *ListSpatialPublicOffersItem) GetPriceCents() float32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *ListSpatialPublicOffersItem) GetPriceCentsOk() (*float32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *ListSpatialPublicOffersItem) SetPriceCents(v float32)`

SetPriceCents sets PriceCents field to given value.


### GetCurrency

`func (o *ListSpatialPublicOffersItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListSpatialPublicOffersItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListSpatialPublicOffersItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetAvailable

`func (o *ListSpatialPublicOffersItem) GetAvailable() float32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ListSpatialPublicOffersItem) GetAvailableOk() (*float32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ListSpatialPublicOffersItem) SetAvailable(v float32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *ListSpatialPublicOffersItem) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *ListSpatialPublicOffersItem) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *ListSpatialPublicOffersItem) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetCondition

`func (o *ListSpatialPublicOffersItem) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ListSpatialPublicOffersItem) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ListSpatialPublicOffersItem) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ListSpatialPublicOffersItem) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *ListSpatialPublicOffersItem) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *ListSpatialPublicOffersItem) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetGradeKey

`func (o *ListSpatialPublicOffersItem) GetGradeKey() string`

GetGradeKey returns the GradeKey field if non-nil, zero value otherwise.

### GetGradeKeyOk

`func (o *ListSpatialPublicOffersItem) GetGradeKeyOk() (*string, bool)`

GetGradeKeyOk returns a tuple with the GradeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeKey

`func (o *ListSpatialPublicOffersItem) SetGradeKey(v string)`

SetGradeKey sets GradeKey field to given value.

### HasGradeKey

`func (o *ListSpatialPublicOffersItem) HasGradeKey() bool`

HasGradeKey returns a boolean if a field has been set.

### SetGradeKeyNil

`func (o *ListSpatialPublicOffersItem) SetGradeKeyNil(b bool)`

 SetGradeKeyNil sets the value for GradeKey to be an explicit nil

### UnsetGradeKey
`func (o *ListSpatialPublicOffersItem) UnsetGradeKey()`

UnsetGradeKey ensures that no value is present for GradeKey, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


