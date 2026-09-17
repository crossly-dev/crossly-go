# ListBuyerCatalogSearchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** |  | 
**Title** | **string** |  | 
**PriceCents** | **float32** |  | 
**CompareAtCents** | Pointer to **NullableFloat32** | MSRP above the ask, or null. Never fabricated from a stale value. | [optional] 
**Currency** | **string** |  | 
**Condition** | Pointer to **NullableString** |  | [optional] 
**Brand** | Pointer to **NullableString** |  | [optional] 
**CategoryMain** | Pointer to **NullableString** |  | [optional] 
**CategorySub** | Pointer to **NullableString** |  | [optional] 
**Thumbnail** | Pointer to **NullableString** |  | [optional] 
**Images** | **[]string** |  | 
**SellerUsername** | Pointer to **NullableString** |  | [optional] 
**SellerDisplayName** | Pointer to **NullableString** |  | [optional] 
**QuantityAvailable** | **float32** | Units a buyer can actually take right now. Reserved units are excluded. | 
**ListedAt** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** | Canonical buyer-facing URL, so a client never has to build one. | 

## Methods

### NewListBuyerCatalogSearchItem

`func NewListBuyerCatalogSearchItem(slug string, title string, priceCents float32, currency string, images []string, quantityAvailable float32, url string, ) *ListBuyerCatalogSearchItem`

NewListBuyerCatalogSearchItem instantiates a new ListBuyerCatalogSearchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBuyerCatalogSearchItemWithDefaults

`func NewListBuyerCatalogSearchItemWithDefaults() *ListBuyerCatalogSearchItem`

NewListBuyerCatalogSearchItemWithDefaults instantiates a new ListBuyerCatalogSearchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *ListBuyerCatalogSearchItem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ListBuyerCatalogSearchItem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ListBuyerCatalogSearchItem) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTitle

`func (o *ListBuyerCatalogSearchItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListBuyerCatalogSearchItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListBuyerCatalogSearchItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPriceCents

`func (o *ListBuyerCatalogSearchItem) GetPriceCents() float32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *ListBuyerCatalogSearchItem) GetPriceCentsOk() (*float32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *ListBuyerCatalogSearchItem) SetPriceCents(v float32)`

SetPriceCents sets PriceCents field to given value.


### GetCompareAtCents

`func (o *ListBuyerCatalogSearchItem) GetCompareAtCents() float32`

GetCompareAtCents returns the CompareAtCents field if non-nil, zero value otherwise.

### GetCompareAtCentsOk

`func (o *ListBuyerCatalogSearchItem) GetCompareAtCentsOk() (*float32, bool)`

GetCompareAtCentsOk returns a tuple with the CompareAtCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtCents

`func (o *ListBuyerCatalogSearchItem) SetCompareAtCents(v float32)`

SetCompareAtCents sets CompareAtCents field to given value.

### HasCompareAtCents

`func (o *ListBuyerCatalogSearchItem) HasCompareAtCents() bool`

HasCompareAtCents returns a boolean if a field has been set.

### SetCompareAtCentsNil

`func (o *ListBuyerCatalogSearchItem) SetCompareAtCentsNil(b bool)`

 SetCompareAtCentsNil sets the value for CompareAtCents to be an explicit nil

### UnsetCompareAtCents
`func (o *ListBuyerCatalogSearchItem) UnsetCompareAtCents()`

UnsetCompareAtCents ensures that no value is present for CompareAtCents, not even an explicit nil
### GetCurrency

`func (o *ListBuyerCatalogSearchItem) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListBuyerCatalogSearchItem) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListBuyerCatalogSearchItem) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCondition

`func (o *ListBuyerCatalogSearchItem) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ListBuyerCatalogSearchItem) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ListBuyerCatalogSearchItem) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ListBuyerCatalogSearchItem) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *ListBuyerCatalogSearchItem) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *ListBuyerCatalogSearchItem) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetBrand

`func (o *ListBuyerCatalogSearchItem) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ListBuyerCatalogSearchItem) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ListBuyerCatalogSearchItem) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ListBuyerCatalogSearchItem) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *ListBuyerCatalogSearchItem) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *ListBuyerCatalogSearchItem) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetCategoryMain

`func (o *ListBuyerCatalogSearchItem) GetCategoryMain() string`

GetCategoryMain returns the CategoryMain field if non-nil, zero value otherwise.

### GetCategoryMainOk

`func (o *ListBuyerCatalogSearchItem) GetCategoryMainOk() (*string, bool)`

GetCategoryMainOk returns a tuple with the CategoryMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryMain

`func (o *ListBuyerCatalogSearchItem) SetCategoryMain(v string)`

SetCategoryMain sets CategoryMain field to given value.

### HasCategoryMain

`func (o *ListBuyerCatalogSearchItem) HasCategoryMain() bool`

HasCategoryMain returns a boolean if a field has been set.

### SetCategoryMainNil

`func (o *ListBuyerCatalogSearchItem) SetCategoryMainNil(b bool)`

 SetCategoryMainNil sets the value for CategoryMain to be an explicit nil

### UnsetCategoryMain
`func (o *ListBuyerCatalogSearchItem) UnsetCategoryMain()`

UnsetCategoryMain ensures that no value is present for CategoryMain, not even an explicit nil
### GetCategorySub

`func (o *ListBuyerCatalogSearchItem) GetCategorySub() string`

GetCategorySub returns the CategorySub field if non-nil, zero value otherwise.

### GetCategorySubOk

`func (o *ListBuyerCatalogSearchItem) GetCategorySubOk() (*string, bool)`

GetCategorySubOk returns a tuple with the CategorySub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySub

`func (o *ListBuyerCatalogSearchItem) SetCategorySub(v string)`

SetCategorySub sets CategorySub field to given value.

### HasCategorySub

`func (o *ListBuyerCatalogSearchItem) HasCategorySub() bool`

HasCategorySub returns a boolean if a field has been set.

### SetCategorySubNil

`func (o *ListBuyerCatalogSearchItem) SetCategorySubNil(b bool)`

 SetCategorySubNil sets the value for CategorySub to be an explicit nil

### UnsetCategorySub
`func (o *ListBuyerCatalogSearchItem) UnsetCategorySub()`

UnsetCategorySub ensures that no value is present for CategorySub, not even an explicit nil
### GetThumbnail

`func (o *ListBuyerCatalogSearchItem) GetThumbnail() string`

GetThumbnail returns the Thumbnail field if non-nil, zero value otherwise.

### GetThumbnailOk

`func (o *ListBuyerCatalogSearchItem) GetThumbnailOk() (*string, bool)`

GetThumbnailOk returns a tuple with the Thumbnail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThumbnail

`func (o *ListBuyerCatalogSearchItem) SetThumbnail(v string)`

SetThumbnail sets Thumbnail field to given value.

### HasThumbnail

`func (o *ListBuyerCatalogSearchItem) HasThumbnail() bool`

HasThumbnail returns a boolean if a field has been set.

### SetThumbnailNil

`func (o *ListBuyerCatalogSearchItem) SetThumbnailNil(b bool)`

 SetThumbnailNil sets the value for Thumbnail to be an explicit nil

### UnsetThumbnail
`func (o *ListBuyerCatalogSearchItem) UnsetThumbnail()`

UnsetThumbnail ensures that no value is present for Thumbnail, not even an explicit nil
### GetImages

`func (o *ListBuyerCatalogSearchItem) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ListBuyerCatalogSearchItem) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ListBuyerCatalogSearchItem) SetImages(v []string)`

SetImages sets Images field to given value.


### GetSellerUsername

`func (o *ListBuyerCatalogSearchItem) GetSellerUsername() string`

GetSellerUsername returns the SellerUsername field if non-nil, zero value otherwise.

### GetSellerUsernameOk

`func (o *ListBuyerCatalogSearchItem) GetSellerUsernameOk() (*string, bool)`

GetSellerUsernameOk returns a tuple with the SellerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerUsername

`func (o *ListBuyerCatalogSearchItem) SetSellerUsername(v string)`

SetSellerUsername sets SellerUsername field to given value.

### HasSellerUsername

`func (o *ListBuyerCatalogSearchItem) HasSellerUsername() bool`

HasSellerUsername returns a boolean if a field has been set.

### SetSellerUsernameNil

`func (o *ListBuyerCatalogSearchItem) SetSellerUsernameNil(b bool)`

 SetSellerUsernameNil sets the value for SellerUsername to be an explicit nil

### UnsetSellerUsername
`func (o *ListBuyerCatalogSearchItem) UnsetSellerUsername()`

UnsetSellerUsername ensures that no value is present for SellerUsername, not even an explicit nil
### GetSellerDisplayName

`func (o *ListBuyerCatalogSearchItem) GetSellerDisplayName() string`

GetSellerDisplayName returns the SellerDisplayName field if non-nil, zero value otherwise.

### GetSellerDisplayNameOk

`func (o *ListBuyerCatalogSearchItem) GetSellerDisplayNameOk() (*string, bool)`

GetSellerDisplayNameOk returns a tuple with the SellerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerDisplayName

`func (o *ListBuyerCatalogSearchItem) SetSellerDisplayName(v string)`

SetSellerDisplayName sets SellerDisplayName field to given value.

### HasSellerDisplayName

`func (o *ListBuyerCatalogSearchItem) HasSellerDisplayName() bool`

HasSellerDisplayName returns a boolean if a field has been set.

### SetSellerDisplayNameNil

`func (o *ListBuyerCatalogSearchItem) SetSellerDisplayNameNil(b bool)`

 SetSellerDisplayNameNil sets the value for SellerDisplayName to be an explicit nil

### UnsetSellerDisplayName
`func (o *ListBuyerCatalogSearchItem) UnsetSellerDisplayName()`

UnsetSellerDisplayName ensures that no value is present for SellerDisplayName, not even an explicit nil
### GetQuantityAvailable

`func (o *ListBuyerCatalogSearchItem) GetQuantityAvailable() float32`

GetQuantityAvailable returns the QuantityAvailable field if non-nil, zero value otherwise.

### GetQuantityAvailableOk

`func (o *ListBuyerCatalogSearchItem) GetQuantityAvailableOk() (*float32, bool)`

GetQuantityAvailableOk returns a tuple with the QuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAvailable

`func (o *ListBuyerCatalogSearchItem) SetQuantityAvailable(v float32)`

SetQuantityAvailable sets QuantityAvailable field to given value.


### GetListedAt

`func (o *ListBuyerCatalogSearchItem) GetListedAt() string`

GetListedAt returns the ListedAt field if non-nil, zero value otherwise.

### GetListedAtOk

`func (o *ListBuyerCatalogSearchItem) GetListedAtOk() (*string, bool)`

GetListedAtOk returns a tuple with the ListedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedAt

`func (o *ListBuyerCatalogSearchItem) SetListedAt(v string)`

SetListedAt sets ListedAt field to given value.

### HasListedAt

`func (o *ListBuyerCatalogSearchItem) HasListedAt() bool`

HasListedAt returns a boolean if a field has been set.

### SetListedAtNil

`func (o *ListBuyerCatalogSearchItem) SetListedAtNil(b bool)`

 SetListedAtNil sets the value for ListedAt to be an explicit nil

### UnsetListedAt
`func (o *ListBuyerCatalogSearchItem) UnsetListedAt()`

UnsetListedAt ensures that no value is present for ListedAt, not even an explicit nil
### GetUrl

`func (o *ListBuyerCatalogSearchItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ListBuyerCatalogSearchItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ListBuyerCatalogSearchItem) SetUrl(v string)`

SetUrl sets Url field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


