# CreateMagicScanResponseEbayHits

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | **string** |  | 
**LegacyItemId** | Pointer to **NullableString** |  | [optional] 
**Title** | **string** |  | 
**Brand** | Pointer to **NullableString** |  | [optional] 
**PriceCents** | Pointer to **NullableFloat32** | Normalized cents. eBay returns string + currency on &#x60;price.value&#x60;. | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 
**Condition** | Pointer to **NullableString** |  | [optional] 
**CategoryId** | Pointer to **NullableString** | Top-level category eBay assigned to the match (id + path). | [optional] 
**CategoryPath** | Pointer to **NullableString** |  | [optional] 
**ItemUrl** | **string** |  | 
**ThumbnailUrl** | Pointer to **NullableString** |  | [optional] 
**Aspects** | Pointer to **map[string]interface{}** | Loosely-typed aspect bag — Brand, Color, Material, etc. when eBay inlines them. Always inspected defensively by the synthesizer. | [optional] 

## Methods

### NewCreateMagicScanResponseEbayHits

`func NewCreateMagicScanResponseEbayHits(itemId string, title string, itemUrl string, ) *CreateMagicScanResponseEbayHits`

NewCreateMagicScanResponseEbayHits instantiates a new CreateMagicScanResponseEbayHits object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMagicScanResponseEbayHitsWithDefaults

`func NewCreateMagicScanResponseEbayHitsWithDefaults() *CreateMagicScanResponseEbayHits`

NewCreateMagicScanResponseEbayHitsWithDefaults instantiates a new CreateMagicScanResponseEbayHits object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *CreateMagicScanResponseEbayHits) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *CreateMagicScanResponseEbayHits) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *CreateMagicScanResponseEbayHits) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetLegacyItemId

`func (o *CreateMagicScanResponseEbayHits) GetLegacyItemId() string`

GetLegacyItemId returns the LegacyItemId field if non-nil, zero value otherwise.

### GetLegacyItemIdOk

`func (o *CreateMagicScanResponseEbayHits) GetLegacyItemIdOk() (*string, bool)`

GetLegacyItemIdOk returns a tuple with the LegacyItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyItemId

`func (o *CreateMagicScanResponseEbayHits) SetLegacyItemId(v string)`

SetLegacyItemId sets LegacyItemId field to given value.

### HasLegacyItemId

`func (o *CreateMagicScanResponseEbayHits) HasLegacyItemId() bool`

HasLegacyItemId returns a boolean if a field has been set.

### SetLegacyItemIdNil

`func (o *CreateMagicScanResponseEbayHits) SetLegacyItemIdNil(b bool)`

 SetLegacyItemIdNil sets the value for LegacyItemId to be an explicit nil

### UnsetLegacyItemId
`func (o *CreateMagicScanResponseEbayHits) UnsetLegacyItemId()`

UnsetLegacyItemId ensures that no value is present for LegacyItemId, not even an explicit nil
### GetTitle

`func (o *CreateMagicScanResponseEbayHits) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateMagicScanResponseEbayHits) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateMagicScanResponseEbayHits) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBrand

`func (o *CreateMagicScanResponseEbayHits) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CreateMagicScanResponseEbayHits) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CreateMagicScanResponseEbayHits) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CreateMagicScanResponseEbayHits) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *CreateMagicScanResponseEbayHits) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *CreateMagicScanResponseEbayHits) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetPriceCents

`func (o *CreateMagicScanResponseEbayHits) GetPriceCents() float32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *CreateMagicScanResponseEbayHits) GetPriceCentsOk() (*float32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *CreateMagicScanResponseEbayHits) SetPriceCents(v float32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *CreateMagicScanResponseEbayHits) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### SetPriceCentsNil

`func (o *CreateMagicScanResponseEbayHits) SetPriceCentsNil(b bool)`

 SetPriceCentsNil sets the value for PriceCents to be an explicit nil

### UnsetPriceCents
`func (o *CreateMagicScanResponseEbayHits) UnsetPriceCents()`

UnsetPriceCents ensures that no value is present for PriceCents, not even an explicit nil
### GetCurrency

`func (o *CreateMagicScanResponseEbayHits) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateMagicScanResponseEbayHits) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateMagicScanResponseEbayHits) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateMagicScanResponseEbayHits) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *CreateMagicScanResponseEbayHits) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CreateMagicScanResponseEbayHits) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCondition

`func (o *CreateMagicScanResponseEbayHits) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CreateMagicScanResponseEbayHits) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CreateMagicScanResponseEbayHits) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *CreateMagicScanResponseEbayHits) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *CreateMagicScanResponseEbayHits) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *CreateMagicScanResponseEbayHits) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetCategoryId

`func (o *CreateMagicScanResponseEbayHits) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CreateMagicScanResponseEbayHits) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CreateMagicScanResponseEbayHits) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *CreateMagicScanResponseEbayHits) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### SetCategoryIdNil

`func (o *CreateMagicScanResponseEbayHits) SetCategoryIdNil(b bool)`

 SetCategoryIdNil sets the value for CategoryId to be an explicit nil

### UnsetCategoryId
`func (o *CreateMagicScanResponseEbayHits) UnsetCategoryId()`

UnsetCategoryId ensures that no value is present for CategoryId, not even an explicit nil
### GetCategoryPath

`func (o *CreateMagicScanResponseEbayHits) GetCategoryPath() string`

GetCategoryPath returns the CategoryPath field if non-nil, zero value otherwise.

### GetCategoryPathOk

`func (o *CreateMagicScanResponseEbayHits) GetCategoryPathOk() (*string, bool)`

GetCategoryPathOk returns a tuple with the CategoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryPath

`func (o *CreateMagicScanResponseEbayHits) SetCategoryPath(v string)`

SetCategoryPath sets CategoryPath field to given value.

### HasCategoryPath

`func (o *CreateMagicScanResponseEbayHits) HasCategoryPath() bool`

HasCategoryPath returns a boolean if a field has been set.

### SetCategoryPathNil

`func (o *CreateMagicScanResponseEbayHits) SetCategoryPathNil(b bool)`

 SetCategoryPathNil sets the value for CategoryPath to be an explicit nil

### UnsetCategoryPath
`func (o *CreateMagicScanResponseEbayHits) UnsetCategoryPath()`

UnsetCategoryPath ensures that no value is present for CategoryPath, not even an explicit nil
### GetItemUrl

`func (o *CreateMagicScanResponseEbayHits) GetItemUrl() string`

GetItemUrl returns the ItemUrl field if non-nil, zero value otherwise.

### GetItemUrlOk

`func (o *CreateMagicScanResponseEbayHits) GetItemUrlOk() (*string, bool)`

GetItemUrlOk returns a tuple with the ItemUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemUrl

`func (o *CreateMagicScanResponseEbayHits) SetItemUrl(v string)`

SetItemUrl sets ItemUrl field to given value.


### GetThumbnailUrl

`func (o *CreateMagicScanResponseEbayHits) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *CreateMagicScanResponseEbayHits) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *CreateMagicScanResponseEbayHits) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *CreateMagicScanResponseEbayHits) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *CreateMagicScanResponseEbayHits) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *CreateMagicScanResponseEbayHits) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetAspects

`func (o *CreateMagicScanResponseEbayHits) GetAspects() map[string]interface{}`

GetAspects returns the Aspects field if non-nil, zero value otherwise.

### GetAspectsOk

`func (o *CreateMagicScanResponseEbayHits) GetAspectsOk() (*map[string]interface{}, bool)`

GetAspectsOk returns a tuple with the Aspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspects

`func (o *CreateMagicScanResponseEbayHits) SetAspects(v map[string]interface{})`

SetAspects sets Aspects field to given value.

### HasAspects

`func (o *CreateMagicScanResponseEbayHits) HasAspects() bool`

HasAspects returns a boolean if a field has been set.

### SetAspectsNil

`func (o *CreateMagicScanResponseEbayHits) SetAspectsNil(b bool)`

 SetAspectsNil sets the value for Aspects to be an explicit nil

### UnsetAspects
`func (o *CreateMagicScanResponseEbayHits) UnsetAspects()`

UnsetAspects ensures that no value is present for Aspects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


