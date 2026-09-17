# GetBuyerCatalogListingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**PriceCents** | **float32** |  | 
**CompareAtCents** | Pointer to **NullableFloat32** |  | [optional] 
**Currency** | **string** |  | 
**Status** | **string** |  | 
**Available** | **bool** |  | 
**QuantityAvailable** | **float32** |  | 
**Condition** | Pointer to **NullableString** |  | [optional] 
**Brand** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableString** |  | [optional] 
**Color** | Pointer to **[]string** |  | [optional] 
**CategoryMain** | Pointer to **NullableString** |  | [optional] 
**CategorySub** | Pointer to **NullableString** |  | [optional] 
**Images** | **[]string** |  | 
**SellerUsername** | Pointer to **NullableString** |  | [optional] 
**SellerDisplayName** | Pointer to **NullableString** |  | [optional] 
**ListedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetBuyerCatalogListingResponse

`func NewGetBuyerCatalogListingResponse(slug string, title string, priceCents float32, currency string, status string, available bool, quantityAvailable float32, images []string, ) *GetBuyerCatalogListingResponse`

NewGetBuyerCatalogListingResponse instantiates a new GetBuyerCatalogListingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBuyerCatalogListingResponseWithDefaults

`func NewGetBuyerCatalogListingResponseWithDefaults() *GetBuyerCatalogListingResponse`

NewGetBuyerCatalogListingResponseWithDefaults instantiates a new GetBuyerCatalogListingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *GetBuyerCatalogListingResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *GetBuyerCatalogListingResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *GetBuyerCatalogListingResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTitle

`func (o *GetBuyerCatalogListingResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetBuyerCatalogListingResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetBuyerCatalogListingResponse) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *GetBuyerCatalogListingResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetBuyerCatalogListingResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetBuyerCatalogListingResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetBuyerCatalogListingResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetBuyerCatalogListingResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetBuyerCatalogListingResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPriceCents

`func (o *GetBuyerCatalogListingResponse) GetPriceCents() float32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *GetBuyerCatalogListingResponse) GetPriceCentsOk() (*float32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *GetBuyerCatalogListingResponse) SetPriceCents(v float32)`

SetPriceCents sets PriceCents field to given value.


### GetCompareAtCents

`func (o *GetBuyerCatalogListingResponse) GetCompareAtCents() float32`

GetCompareAtCents returns the CompareAtCents field if non-nil, zero value otherwise.

### GetCompareAtCentsOk

`func (o *GetBuyerCatalogListingResponse) GetCompareAtCentsOk() (*float32, bool)`

GetCompareAtCentsOk returns a tuple with the CompareAtCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompareAtCents

`func (o *GetBuyerCatalogListingResponse) SetCompareAtCents(v float32)`

SetCompareAtCents sets CompareAtCents field to given value.

### HasCompareAtCents

`func (o *GetBuyerCatalogListingResponse) HasCompareAtCents() bool`

HasCompareAtCents returns a boolean if a field has been set.

### SetCompareAtCentsNil

`func (o *GetBuyerCatalogListingResponse) SetCompareAtCentsNil(b bool)`

 SetCompareAtCentsNil sets the value for CompareAtCents to be an explicit nil

### UnsetCompareAtCents
`func (o *GetBuyerCatalogListingResponse) UnsetCompareAtCents()`

UnsetCompareAtCents ensures that no value is present for CompareAtCents, not even an explicit nil
### GetCurrency

`func (o *GetBuyerCatalogListingResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetBuyerCatalogListingResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetBuyerCatalogListingResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetStatus

`func (o *GetBuyerCatalogListingResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetBuyerCatalogListingResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetBuyerCatalogListingResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAvailable

`func (o *GetBuyerCatalogListingResponse) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *GetBuyerCatalogListingResponse) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *GetBuyerCatalogListingResponse) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetQuantityAvailable

`func (o *GetBuyerCatalogListingResponse) GetQuantityAvailable() float32`

GetQuantityAvailable returns the QuantityAvailable field if non-nil, zero value otherwise.

### GetQuantityAvailableOk

`func (o *GetBuyerCatalogListingResponse) GetQuantityAvailableOk() (*float32, bool)`

GetQuantityAvailableOk returns a tuple with the QuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAvailable

`func (o *GetBuyerCatalogListingResponse) SetQuantityAvailable(v float32)`

SetQuantityAvailable sets QuantityAvailable field to given value.


### GetCondition

`func (o *GetBuyerCatalogListingResponse) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GetBuyerCatalogListingResponse) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *GetBuyerCatalogListingResponse) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *GetBuyerCatalogListingResponse) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *GetBuyerCatalogListingResponse) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *GetBuyerCatalogListingResponse) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetBrand

`func (o *GetBuyerCatalogListingResponse) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *GetBuyerCatalogListingResponse) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *GetBuyerCatalogListingResponse) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *GetBuyerCatalogListingResponse) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *GetBuyerCatalogListingResponse) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *GetBuyerCatalogListingResponse) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetSize

`func (o *GetBuyerCatalogListingResponse) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetBuyerCatalogListingResponse) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetBuyerCatalogListingResponse) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetBuyerCatalogListingResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *GetBuyerCatalogListingResponse) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *GetBuyerCatalogListingResponse) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetColor

`func (o *GetBuyerCatalogListingResponse) GetColor() []string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *GetBuyerCatalogListingResponse) GetColorOk() (*[]string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *GetBuyerCatalogListingResponse) SetColor(v []string)`

SetColor sets Color field to given value.

### HasColor

`func (o *GetBuyerCatalogListingResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *GetBuyerCatalogListingResponse) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *GetBuyerCatalogListingResponse) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetCategoryMain

`func (o *GetBuyerCatalogListingResponse) GetCategoryMain() string`

GetCategoryMain returns the CategoryMain field if non-nil, zero value otherwise.

### GetCategoryMainOk

`func (o *GetBuyerCatalogListingResponse) GetCategoryMainOk() (*string, bool)`

GetCategoryMainOk returns a tuple with the CategoryMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryMain

`func (o *GetBuyerCatalogListingResponse) SetCategoryMain(v string)`

SetCategoryMain sets CategoryMain field to given value.

### HasCategoryMain

`func (o *GetBuyerCatalogListingResponse) HasCategoryMain() bool`

HasCategoryMain returns a boolean if a field has been set.

### SetCategoryMainNil

`func (o *GetBuyerCatalogListingResponse) SetCategoryMainNil(b bool)`

 SetCategoryMainNil sets the value for CategoryMain to be an explicit nil

### UnsetCategoryMain
`func (o *GetBuyerCatalogListingResponse) UnsetCategoryMain()`

UnsetCategoryMain ensures that no value is present for CategoryMain, not even an explicit nil
### GetCategorySub

`func (o *GetBuyerCatalogListingResponse) GetCategorySub() string`

GetCategorySub returns the CategorySub field if non-nil, zero value otherwise.

### GetCategorySubOk

`func (o *GetBuyerCatalogListingResponse) GetCategorySubOk() (*string, bool)`

GetCategorySubOk returns a tuple with the CategorySub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySub

`func (o *GetBuyerCatalogListingResponse) SetCategorySub(v string)`

SetCategorySub sets CategorySub field to given value.

### HasCategorySub

`func (o *GetBuyerCatalogListingResponse) HasCategorySub() bool`

HasCategorySub returns a boolean if a field has been set.

### SetCategorySubNil

`func (o *GetBuyerCatalogListingResponse) SetCategorySubNil(b bool)`

 SetCategorySubNil sets the value for CategorySub to be an explicit nil

### UnsetCategorySub
`func (o *GetBuyerCatalogListingResponse) UnsetCategorySub()`

UnsetCategorySub ensures that no value is present for CategorySub, not even an explicit nil
### GetImages

`func (o *GetBuyerCatalogListingResponse) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *GetBuyerCatalogListingResponse) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *GetBuyerCatalogListingResponse) SetImages(v []string)`

SetImages sets Images field to given value.


### GetSellerUsername

`func (o *GetBuyerCatalogListingResponse) GetSellerUsername() string`

GetSellerUsername returns the SellerUsername field if non-nil, zero value otherwise.

### GetSellerUsernameOk

`func (o *GetBuyerCatalogListingResponse) GetSellerUsernameOk() (*string, bool)`

GetSellerUsernameOk returns a tuple with the SellerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerUsername

`func (o *GetBuyerCatalogListingResponse) SetSellerUsername(v string)`

SetSellerUsername sets SellerUsername field to given value.

### HasSellerUsername

`func (o *GetBuyerCatalogListingResponse) HasSellerUsername() bool`

HasSellerUsername returns a boolean if a field has been set.

### SetSellerUsernameNil

`func (o *GetBuyerCatalogListingResponse) SetSellerUsernameNil(b bool)`

 SetSellerUsernameNil sets the value for SellerUsername to be an explicit nil

### UnsetSellerUsername
`func (o *GetBuyerCatalogListingResponse) UnsetSellerUsername()`

UnsetSellerUsername ensures that no value is present for SellerUsername, not even an explicit nil
### GetSellerDisplayName

`func (o *GetBuyerCatalogListingResponse) GetSellerDisplayName() string`

GetSellerDisplayName returns the SellerDisplayName field if non-nil, zero value otherwise.

### GetSellerDisplayNameOk

`func (o *GetBuyerCatalogListingResponse) GetSellerDisplayNameOk() (*string, bool)`

GetSellerDisplayNameOk returns a tuple with the SellerDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerDisplayName

`func (o *GetBuyerCatalogListingResponse) SetSellerDisplayName(v string)`

SetSellerDisplayName sets SellerDisplayName field to given value.

### HasSellerDisplayName

`func (o *GetBuyerCatalogListingResponse) HasSellerDisplayName() bool`

HasSellerDisplayName returns a boolean if a field has been set.

### SetSellerDisplayNameNil

`func (o *GetBuyerCatalogListingResponse) SetSellerDisplayNameNil(b bool)`

 SetSellerDisplayNameNil sets the value for SellerDisplayName to be an explicit nil

### UnsetSellerDisplayName
`func (o *GetBuyerCatalogListingResponse) UnsetSellerDisplayName()`

UnsetSellerDisplayName ensures that no value is present for SellerDisplayName, not even an explicit nil
### GetListedAt

`func (o *GetBuyerCatalogListingResponse) GetListedAt() string`

GetListedAt returns the ListedAt field if non-nil, zero value otherwise.

### GetListedAtOk

`func (o *GetBuyerCatalogListingResponse) GetListedAtOk() (*string, bool)`

GetListedAtOk returns a tuple with the ListedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedAt

`func (o *GetBuyerCatalogListingResponse) SetListedAt(v string)`

SetListedAt sets ListedAt field to given value.

### HasListedAt

`func (o *GetBuyerCatalogListingResponse) HasListedAt() bool`

HasListedAt returns a boolean if a field has been set.

### SetListedAtNil

`func (o *GetBuyerCatalogListingResponse) SetListedAtNil(b bool)`

 SetListedAtNil sets the value for ListedAt to be an explicit nil

### UnsetListedAt
`func (o *GetBuyerCatalogListingResponse) UnsetListedAt()`

UnsetListedAt ensures that no value is present for ListedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


