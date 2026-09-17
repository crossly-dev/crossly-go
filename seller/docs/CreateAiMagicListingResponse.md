# CreateAiMagicListingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to [**NullableCreateAiGenerateListingResponseDescription**](CreateAiGenerateListingResponseDescription.md) |  | [optional] 
**SuggestedPrice** | Pointer to [**NullableCreateAiGenerateListingResponseSuggestedPrice**](CreateAiGenerateListingResponseSuggestedPrice.md) |  | [optional] 
**PriceRange** | Pointer to [**NullableCreateAiGenerateListingResponsePriceRange**](CreateAiGenerateListingResponsePriceRange.md) |  | [optional] 
**Category** | Pointer to [**NullableCreateAiGenerateListingResponseDescription**](CreateAiGenerateListingResponseDescription.md) |  | [optional] 
**SubCategory** | Pointer to [**NullableCreateAiGenerateListingResponseDescription**](CreateAiGenerateListingResponseDescription.md) |  | [optional] 
**Brand** | Pointer to [**NullableCreateAiGenerateListingResponseDescription**](CreateAiGenerateListingResponseDescription.md) |  | [optional] 
**Condition** | Pointer to [**NullableCreateAiGenerateListingResponseDescription**](CreateAiGenerateListingResponseDescription.md) |  | [optional] 
**Colors** | Pointer to [**NullableCreateAiGenerateListingResponseColors**](CreateAiGenerateListingResponseColors.md) |  | [optional] 
**Size** | Pointer to [**NullableCreateAiGenerateListingResponseDescription**](CreateAiGenerateListingResponseDescription.md) |  | [optional] 
**Tags** | Pointer to [**NullableCreateAiGenerateListingResponseColors**](CreateAiGenerateListingResponseColors.md) |  | [optional] 
**SuggestedTitle** | Pointer to [**NullableCreateAiGenerateListingResponseDescription**](CreateAiGenerateListingResponseDescription.md) |  | [optional] 
**Material** | Pointer to [**NullableCreateAiGenerateListingResponseDescription**](CreateAiGenerateListingResponseDescription.md) |  | [optional] 
**Style** | Pointer to [**NullableCreateAiGenerateListingResponseDescription**](CreateAiGenerateListingResponseDescription.md) |  | [optional] 
**Pattern** | Pointer to [**NullableCreateAiGenerateListingResponseDescription**](CreateAiGenerateListingResponseDescription.md) |  | [optional] 
**Department** | Pointer to [**NullableCreateAiGenerateListingResponseDescription**](CreateAiGenerateListingResponseDescription.md) |  | [optional] 
**Gender** | Pointer to [**NullableCreateAiGenerateListingResponseDescription**](CreateAiGenerateListingResponseDescription.md) |  | [optional] 
**ItemType** | Pointer to [**NullableCreateAiGenerateListingResponseDescription**](CreateAiGenerateListingResponseDescription.md) |  | [optional] 
**SizeSystem** | Pointer to [**NullableCreateAiGenerateListingResponseDescription**](CreateAiGenerateListingResponseDescription.md) |  | [optional] 

## Methods

### NewCreateAiMagicListingResponse

`func NewCreateAiMagicListingResponse() *CreateAiMagicListingResponse`

NewCreateAiMagicListingResponse instantiates a new CreateAiMagicListingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAiMagicListingResponseWithDefaults

`func NewCreateAiMagicListingResponseWithDefaults() *CreateAiMagicListingResponse`

NewCreateAiMagicListingResponseWithDefaults instantiates a new CreateAiMagicListingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CreateAiMagicListingResponse) GetDescription() CreateAiGenerateListingResponseDescription`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateAiMagicListingResponse) GetDescriptionOk() (*CreateAiGenerateListingResponseDescription, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateAiMagicListingResponse) SetDescription(v CreateAiGenerateListingResponseDescription)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateAiMagicListingResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateAiMagicListingResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateAiMagicListingResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetSuggestedPrice

`func (o *CreateAiMagicListingResponse) GetSuggestedPrice() CreateAiGenerateListingResponseSuggestedPrice`

GetSuggestedPrice returns the SuggestedPrice field if non-nil, zero value otherwise.

### GetSuggestedPriceOk

`func (o *CreateAiMagicListingResponse) GetSuggestedPriceOk() (*CreateAiGenerateListingResponseSuggestedPrice, bool)`

GetSuggestedPriceOk returns a tuple with the SuggestedPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedPrice

`func (o *CreateAiMagicListingResponse) SetSuggestedPrice(v CreateAiGenerateListingResponseSuggestedPrice)`

SetSuggestedPrice sets SuggestedPrice field to given value.

### HasSuggestedPrice

`func (o *CreateAiMagicListingResponse) HasSuggestedPrice() bool`

HasSuggestedPrice returns a boolean if a field has been set.

### SetSuggestedPriceNil

`func (o *CreateAiMagicListingResponse) SetSuggestedPriceNil(b bool)`

 SetSuggestedPriceNil sets the value for SuggestedPrice to be an explicit nil

### UnsetSuggestedPrice
`func (o *CreateAiMagicListingResponse) UnsetSuggestedPrice()`

UnsetSuggestedPrice ensures that no value is present for SuggestedPrice, not even an explicit nil
### GetPriceRange

`func (o *CreateAiMagicListingResponse) GetPriceRange() CreateAiGenerateListingResponsePriceRange`

GetPriceRange returns the PriceRange field if non-nil, zero value otherwise.

### GetPriceRangeOk

`func (o *CreateAiMagicListingResponse) GetPriceRangeOk() (*CreateAiGenerateListingResponsePriceRange, bool)`

GetPriceRangeOk returns a tuple with the PriceRange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceRange

`func (o *CreateAiMagicListingResponse) SetPriceRange(v CreateAiGenerateListingResponsePriceRange)`

SetPriceRange sets PriceRange field to given value.

### HasPriceRange

`func (o *CreateAiMagicListingResponse) HasPriceRange() bool`

HasPriceRange returns a boolean if a field has been set.

### SetPriceRangeNil

`func (o *CreateAiMagicListingResponse) SetPriceRangeNil(b bool)`

 SetPriceRangeNil sets the value for PriceRange to be an explicit nil

### UnsetPriceRange
`func (o *CreateAiMagicListingResponse) UnsetPriceRange()`

UnsetPriceRange ensures that no value is present for PriceRange, not even an explicit nil
### GetCategory

`func (o *CreateAiMagicListingResponse) GetCategory() CreateAiGenerateListingResponseDescription`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateAiMagicListingResponse) GetCategoryOk() (*CreateAiGenerateListingResponseDescription, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateAiMagicListingResponse) SetCategory(v CreateAiGenerateListingResponseDescription)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateAiMagicListingResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *CreateAiMagicListingResponse) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *CreateAiMagicListingResponse) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetSubCategory

`func (o *CreateAiMagicListingResponse) GetSubCategory() CreateAiGenerateListingResponseDescription`

GetSubCategory returns the SubCategory field if non-nil, zero value otherwise.

### GetSubCategoryOk

`func (o *CreateAiMagicListingResponse) GetSubCategoryOk() (*CreateAiGenerateListingResponseDescription, bool)`

GetSubCategoryOk returns a tuple with the SubCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubCategory

`func (o *CreateAiMagicListingResponse) SetSubCategory(v CreateAiGenerateListingResponseDescription)`

SetSubCategory sets SubCategory field to given value.

### HasSubCategory

`func (o *CreateAiMagicListingResponse) HasSubCategory() bool`

HasSubCategory returns a boolean if a field has been set.

### SetSubCategoryNil

`func (o *CreateAiMagicListingResponse) SetSubCategoryNil(b bool)`

 SetSubCategoryNil sets the value for SubCategory to be an explicit nil

### UnsetSubCategory
`func (o *CreateAiMagicListingResponse) UnsetSubCategory()`

UnsetSubCategory ensures that no value is present for SubCategory, not even an explicit nil
### GetBrand

`func (o *CreateAiMagicListingResponse) GetBrand() CreateAiGenerateListingResponseDescription`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CreateAiMagicListingResponse) GetBrandOk() (*CreateAiGenerateListingResponseDescription, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CreateAiMagicListingResponse) SetBrand(v CreateAiGenerateListingResponseDescription)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CreateAiMagicListingResponse) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *CreateAiMagicListingResponse) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *CreateAiMagicListingResponse) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetCondition

`func (o *CreateAiMagicListingResponse) GetCondition() CreateAiGenerateListingResponseDescription`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CreateAiMagicListingResponse) GetConditionOk() (*CreateAiGenerateListingResponseDescription, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CreateAiMagicListingResponse) SetCondition(v CreateAiGenerateListingResponseDescription)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *CreateAiMagicListingResponse) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *CreateAiMagicListingResponse) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *CreateAiMagicListingResponse) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetColors

`func (o *CreateAiMagicListingResponse) GetColors() CreateAiGenerateListingResponseColors`

GetColors returns the Colors field if non-nil, zero value otherwise.

### GetColorsOk

`func (o *CreateAiMagicListingResponse) GetColorsOk() (*CreateAiGenerateListingResponseColors, bool)`

GetColorsOk returns a tuple with the Colors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColors

`func (o *CreateAiMagicListingResponse) SetColors(v CreateAiGenerateListingResponseColors)`

SetColors sets Colors field to given value.

### HasColors

`func (o *CreateAiMagicListingResponse) HasColors() bool`

HasColors returns a boolean if a field has been set.

### SetColorsNil

`func (o *CreateAiMagicListingResponse) SetColorsNil(b bool)`

 SetColorsNil sets the value for Colors to be an explicit nil

### UnsetColors
`func (o *CreateAiMagicListingResponse) UnsetColors()`

UnsetColors ensures that no value is present for Colors, not even an explicit nil
### GetSize

`func (o *CreateAiMagicListingResponse) GetSize() CreateAiGenerateListingResponseDescription`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateAiMagicListingResponse) GetSizeOk() (*CreateAiGenerateListingResponseDescription, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateAiMagicListingResponse) SetSize(v CreateAiGenerateListingResponseDescription)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateAiMagicListingResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *CreateAiMagicListingResponse) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *CreateAiMagicListingResponse) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetTags

`func (o *CreateAiMagicListingResponse) GetTags() CreateAiGenerateListingResponseColors`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateAiMagicListingResponse) GetTagsOk() (*CreateAiGenerateListingResponseColors, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateAiMagicListingResponse) SetTags(v CreateAiGenerateListingResponseColors)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateAiMagicListingResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateAiMagicListingResponse) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateAiMagicListingResponse) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetSuggestedTitle

`func (o *CreateAiMagicListingResponse) GetSuggestedTitle() CreateAiGenerateListingResponseDescription`

GetSuggestedTitle returns the SuggestedTitle field if non-nil, zero value otherwise.

### GetSuggestedTitleOk

`func (o *CreateAiMagicListingResponse) GetSuggestedTitleOk() (*CreateAiGenerateListingResponseDescription, bool)`

GetSuggestedTitleOk returns a tuple with the SuggestedTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggestedTitle

`func (o *CreateAiMagicListingResponse) SetSuggestedTitle(v CreateAiGenerateListingResponseDescription)`

SetSuggestedTitle sets SuggestedTitle field to given value.

### HasSuggestedTitle

`func (o *CreateAiMagicListingResponse) HasSuggestedTitle() bool`

HasSuggestedTitle returns a boolean if a field has been set.

### SetSuggestedTitleNil

`func (o *CreateAiMagicListingResponse) SetSuggestedTitleNil(b bool)`

 SetSuggestedTitleNil sets the value for SuggestedTitle to be an explicit nil

### UnsetSuggestedTitle
`func (o *CreateAiMagicListingResponse) UnsetSuggestedTitle()`

UnsetSuggestedTitle ensures that no value is present for SuggestedTitle, not even an explicit nil
### GetMaterial

`func (o *CreateAiMagicListingResponse) GetMaterial() CreateAiGenerateListingResponseDescription`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *CreateAiMagicListingResponse) GetMaterialOk() (*CreateAiGenerateListingResponseDescription, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *CreateAiMagicListingResponse) SetMaterial(v CreateAiGenerateListingResponseDescription)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *CreateAiMagicListingResponse) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### SetMaterialNil

`func (o *CreateAiMagicListingResponse) SetMaterialNil(b bool)`

 SetMaterialNil sets the value for Material to be an explicit nil

### UnsetMaterial
`func (o *CreateAiMagicListingResponse) UnsetMaterial()`

UnsetMaterial ensures that no value is present for Material, not even an explicit nil
### GetStyle

`func (o *CreateAiMagicListingResponse) GetStyle() CreateAiGenerateListingResponseDescription`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *CreateAiMagicListingResponse) GetStyleOk() (*CreateAiGenerateListingResponseDescription, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *CreateAiMagicListingResponse) SetStyle(v CreateAiGenerateListingResponseDescription)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *CreateAiMagicListingResponse) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *CreateAiMagicListingResponse) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *CreateAiMagicListingResponse) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetPattern

`func (o *CreateAiMagicListingResponse) GetPattern() CreateAiGenerateListingResponseDescription`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *CreateAiMagicListingResponse) GetPatternOk() (*CreateAiGenerateListingResponseDescription, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *CreateAiMagicListingResponse) SetPattern(v CreateAiGenerateListingResponseDescription)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *CreateAiMagicListingResponse) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPatternNil

`func (o *CreateAiMagicListingResponse) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *CreateAiMagicListingResponse) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil
### GetDepartment

`func (o *CreateAiMagicListingResponse) GetDepartment() CreateAiGenerateListingResponseDescription`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *CreateAiMagicListingResponse) GetDepartmentOk() (*CreateAiGenerateListingResponseDescription, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *CreateAiMagicListingResponse) SetDepartment(v CreateAiGenerateListingResponseDescription)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *CreateAiMagicListingResponse) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *CreateAiMagicListingResponse) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *CreateAiMagicListingResponse) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetGender

`func (o *CreateAiMagicListingResponse) GetGender() CreateAiGenerateListingResponseDescription`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CreateAiMagicListingResponse) GetGenderOk() (*CreateAiGenerateListingResponseDescription, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CreateAiMagicListingResponse) SetGender(v CreateAiGenerateListingResponseDescription)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CreateAiMagicListingResponse) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *CreateAiMagicListingResponse) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *CreateAiMagicListingResponse) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetItemType

`func (o *CreateAiMagicListingResponse) GetItemType() CreateAiGenerateListingResponseDescription`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *CreateAiMagicListingResponse) GetItemTypeOk() (*CreateAiGenerateListingResponseDescription, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *CreateAiMagicListingResponse) SetItemType(v CreateAiGenerateListingResponseDescription)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *CreateAiMagicListingResponse) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *CreateAiMagicListingResponse) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *CreateAiMagicListingResponse) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetSizeSystem

`func (o *CreateAiMagicListingResponse) GetSizeSystem() CreateAiGenerateListingResponseDescription`

GetSizeSystem returns the SizeSystem field if non-nil, zero value otherwise.

### GetSizeSystemOk

`func (o *CreateAiMagicListingResponse) GetSizeSystemOk() (*CreateAiGenerateListingResponseDescription, bool)`

GetSizeSystemOk returns a tuple with the SizeSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSystem

`func (o *CreateAiMagicListingResponse) SetSizeSystem(v CreateAiGenerateListingResponseDescription)`

SetSizeSystem sets SizeSystem field to given value.

### HasSizeSystem

`func (o *CreateAiMagicListingResponse) HasSizeSystem() bool`

HasSizeSystem returns a boolean if a field has been set.

### SetSizeSystemNil

`func (o *CreateAiMagicListingResponse) SetSizeSystemNil(b bool)`

 SetSizeSystemNil sets the value for SizeSystem to be an explicit nil

### UnsetSizeSystem
`func (o *CreateAiMagicListingResponse) UnsetSizeSystem()`

UnsetSizeSystem ensures that no value is present for SizeSystem, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


