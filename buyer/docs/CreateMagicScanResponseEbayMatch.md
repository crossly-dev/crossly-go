# CreateMagicScanResponseEbayMatch

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

### NewCreateMagicScanResponseEbayMatch

`func NewCreateMagicScanResponseEbayMatch(itemId string, title string, itemUrl string, ) *CreateMagicScanResponseEbayMatch`

NewCreateMagicScanResponseEbayMatch instantiates a new CreateMagicScanResponseEbayMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMagicScanResponseEbayMatchWithDefaults

`func NewCreateMagicScanResponseEbayMatchWithDefaults() *CreateMagicScanResponseEbayMatch`

NewCreateMagicScanResponseEbayMatchWithDefaults instantiates a new CreateMagicScanResponseEbayMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *CreateMagicScanResponseEbayMatch) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *CreateMagicScanResponseEbayMatch) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *CreateMagicScanResponseEbayMatch) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetLegacyItemId

`func (o *CreateMagicScanResponseEbayMatch) GetLegacyItemId() string`

GetLegacyItemId returns the LegacyItemId field if non-nil, zero value otherwise.

### GetLegacyItemIdOk

`func (o *CreateMagicScanResponseEbayMatch) GetLegacyItemIdOk() (*string, bool)`

GetLegacyItemIdOk returns a tuple with the LegacyItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegacyItemId

`func (o *CreateMagicScanResponseEbayMatch) SetLegacyItemId(v string)`

SetLegacyItemId sets LegacyItemId field to given value.

### HasLegacyItemId

`func (o *CreateMagicScanResponseEbayMatch) HasLegacyItemId() bool`

HasLegacyItemId returns a boolean if a field has been set.

### SetLegacyItemIdNil

`func (o *CreateMagicScanResponseEbayMatch) SetLegacyItemIdNil(b bool)`

 SetLegacyItemIdNil sets the value for LegacyItemId to be an explicit nil

### UnsetLegacyItemId
`func (o *CreateMagicScanResponseEbayMatch) UnsetLegacyItemId()`

UnsetLegacyItemId ensures that no value is present for LegacyItemId, not even an explicit nil
### GetTitle

`func (o *CreateMagicScanResponseEbayMatch) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateMagicScanResponseEbayMatch) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateMagicScanResponseEbayMatch) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBrand

`func (o *CreateMagicScanResponseEbayMatch) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CreateMagicScanResponseEbayMatch) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CreateMagicScanResponseEbayMatch) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CreateMagicScanResponseEbayMatch) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *CreateMagicScanResponseEbayMatch) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *CreateMagicScanResponseEbayMatch) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetPriceCents

`func (o *CreateMagicScanResponseEbayMatch) GetPriceCents() float32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *CreateMagicScanResponseEbayMatch) GetPriceCentsOk() (*float32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *CreateMagicScanResponseEbayMatch) SetPriceCents(v float32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *CreateMagicScanResponseEbayMatch) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### SetPriceCentsNil

`func (o *CreateMagicScanResponseEbayMatch) SetPriceCentsNil(b bool)`

 SetPriceCentsNil sets the value for PriceCents to be an explicit nil

### UnsetPriceCents
`func (o *CreateMagicScanResponseEbayMatch) UnsetPriceCents()`

UnsetPriceCents ensures that no value is present for PriceCents, not even an explicit nil
### GetCurrency

`func (o *CreateMagicScanResponseEbayMatch) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateMagicScanResponseEbayMatch) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateMagicScanResponseEbayMatch) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CreateMagicScanResponseEbayMatch) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *CreateMagicScanResponseEbayMatch) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *CreateMagicScanResponseEbayMatch) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetCondition

`func (o *CreateMagicScanResponseEbayMatch) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CreateMagicScanResponseEbayMatch) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CreateMagicScanResponseEbayMatch) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *CreateMagicScanResponseEbayMatch) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *CreateMagicScanResponseEbayMatch) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *CreateMagicScanResponseEbayMatch) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetCategoryId

`func (o *CreateMagicScanResponseEbayMatch) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CreateMagicScanResponseEbayMatch) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CreateMagicScanResponseEbayMatch) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *CreateMagicScanResponseEbayMatch) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### SetCategoryIdNil

`func (o *CreateMagicScanResponseEbayMatch) SetCategoryIdNil(b bool)`

 SetCategoryIdNil sets the value for CategoryId to be an explicit nil

### UnsetCategoryId
`func (o *CreateMagicScanResponseEbayMatch) UnsetCategoryId()`

UnsetCategoryId ensures that no value is present for CategoryId, not even an explicit nil
### GetCategoryPath

`func (o *CreateMagicScanResponseEbayMatch) GetCategoryPath() string`

GetCategoryPath returns the CategoryPath field if non-nil, zero value otherwise.

### GetCategoryPathOk

`func (o *CreateMagicScanResponseEbayMatch) GetCategoryPathOk() (*string, bool)`

GetCategoryPathOk returns a tuple with the CategoryPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryPath

`func (o *CreateMagicScanResponseEbayMatch) SetCategoryPath(v string)`

SetCategoryPath sets CategoryPath field to given value.

### HasCategoryPath

`func (o *CreateMagicScanResponseEbayMatch) HasCategoryPath() bool`

HasCategoryPath returns a boolean if a field has been set.

### SetCategoryPathNil

`func (o *CreateMagicScanResponseEbayMatch) SetCategoryPathNil(b bool)`

 SetCategoryPathNil sets the value for CategoryPath to be an explicit nil

### UnsetCategoryPath
`func (o *CreateMagicScanResponseEbayMatch) UnsetCategoryPath()`

UnsetCategoryPath ensures that no value is present for CategoryPath, not even an explicit nil
### GetItemUrl

`func (o *CreateMagicScanResponseEbayMatch) GetItemUrl() string`

GetItemUrl returns the ItemUrl field if non-nil, zero value otherwise.

### GetItemUrlOk

`func (o *CreateMagicScanResponseEbayMatch) GetItemUrlOk() (*string, bool)`

GetItemUrlOk returns a tuple with the ItemUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemUrl

`func (o *CreateMagicScanResponseEbayMatch) SetItemUrl(v string)`

SetItemUrl sets ItemUrl field to given value.


### GetThumbnailUrl

`func (o *CreateMagicScanResponseEbayMatch) GetThumbnailUrl() string`

GetThumbnailUrl returns the ThumbnailUrl field if non-nil, zero value otherwise.

### GetThumbnailUrlOk

`func (o *CreateMagicScanResponseEbayMatch) GetThumbnailUrlOk() (*string, bool)`

GetThumbnailUrlOk returns a tuple with the ThumbnailUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnailUrl

`func (o *CreateMagicScanResponseEbayMatch) SetThumbnailUrl(v string)`

SetThumbnailUrl sets ThumbnailUrl field to given value.

### HasThumbnailUrl

`func (o *CreateMagicScanResponseEbayMatch) HasThumbnailUrl() bool`

HasThumbnailUrl returns a boolean if a field has been set.

### SetThumbnailUrlNil

`func (o *CreateMagicScanResponseEbayMatch) SetThumbnailUrlNil(b bool)`

 SetThumbnailUrlNil sets the value for ThumbnailUrl to be an explicit nil

### UnsetThumbnailUrl
`func (o *CreateMagicScanResponseEbayMatch) UnsetThumbnailUrl()`

UnsetThumbnailUrl ensures that no value is present for ThumbnailUrl, not even an explicit nil
### GetAspects

`func (o *CreateMagicScanResponseEbayMatch) GetAspects() map[string]interface{}`

GetAspects returns the Aspects field if non-nil, zero value otherwise.

### GetAspectsOk

`func (o *CreateMagicScanResponseEbayMatch) GetAspectsOk() (*map[string]interface{}, bool)`

GetAspectsOk returns a tuple with the Aspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAspects

`func (o *CreateMagicScanResponseEbayMatch) SetAspects(v map[string]interface{})`

SetAspects sets Aspects field to given value.

### HasAspects

`func (o *CreateMagicScanResponseEbayMatch) HasAspects() bool`

HasAspects returns a boolean if a field has been set.

### SetAspectsNil

`func (o *CreateMagicScanResponseEbayMatch) SetAspectsNil(b bool)`

 SetAspectsNil sets the value for Aspects to be an explicit nil

### UnsetAspects
`func (o *CreateMagicScanResponseEbayMatch) UnsetAspects()`

UnsetAspects ensures that no value is present for Aspects, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


