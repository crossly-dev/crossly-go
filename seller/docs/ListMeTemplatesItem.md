# ListMeTemplatesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Brand** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Condition** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableString** |  | [optional] 
**Material** | Pointer to **NullableString** |  | [optional] 
**Style** | Pointer to **NullableString** |  | [optional] 
**Pattern** | Pointer to **NullableString** |  | [optional] 
**Department** | Pointer to **NullableString** |  | [optional] 
**Gender** | Pointer to **NullableString** |  | [optional] 
**ItemType** | Pointer to **NullableString** |  | [optional] 
**SizeSystem** | Pointer to **NullableString** |  | [optional] 
**Color** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**WeightOz** | Pointer to **NullableFloat32** |  | [optional] 
**IsDefault** | **bool** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Scope** | **string** |  | 
**ShareToken** | Pointer to **NullableString** |  | [optional] 
**DescriptionVariants** | Pointer to **[]string** |  | [optional] 
**TitleVariants** | Pointer to **[]string** |  | [optional] 
**DefaultForCategory** | Pointer to **NullableString** |  | [optional] 
**SortOrder** | **float32** |  | 

## Methods

### NewListMeTemplatesItem

`func NewListMeTemplatesItem(id string, name string, createdAt time.Time, updatedAt time.Time, userId string, isDefault bool, scope string, sortOrder float32, ) *ListMeTemplatesItem`

NewListMeTemplatesItem instantiates a new ListMeTemplatesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMeTemplatesItemWithDefaults

`func NewListMeTemplatesItemWithDefaults() *ListMeTemplatesItem`

NewListMeTemplatesItemWithDefaults instantiates a new ListMeTemplatesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListMeTemplatesItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListMeTemplatesItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListMeTemplatesItem) SetId(v string)`

SetId sets Id field to given value.


### GetBrand

`func (o *ListMeTemplatesItem) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ListMeTemplatesItem) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ListMeTemplatesItem) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ListMeTemplatesItem) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *ListMeTemplatesItem) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *ListMeTemplatesItem) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetName

`func (o *ListMeTemplatesItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListMeTemplatesItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListMeTemplatesItem) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *ListMeTemplatesItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListMeTemplatesItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListMeTemplatesItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListMeTemplatesItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListMeTemplatesItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListMeTemplatesItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *ListMeTemplatesItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListMeTemplatesItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListMeTemplatesItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCondition

`func (o *ListMeTemplatesItem) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ListMeTemplatesItem) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ListMeTemplatesItem) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ListMeTemplatesItem) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *ListMeTemplatesItem) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *ListMeTemplatesItem) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetSize

`func (o *ListMeTemplatesItem) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListMeTemplatesItem) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListMeTemplatesItem) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ListMeTemplatesItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *ListMeTemplatesItem) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *ListMeTemplatesItem) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetMaterial

`func (o *ListMeTemplatesItem) GetMaterial() string`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *ListMeTemplatesItem) GetMaterialOk() (*string, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *ListMeTemplatesItem) SetMaterial(v string)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *ListMeTemplatesItem) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### SetMaterialNil

`func (o *ListMeTemplatesItem) SetMaterialNil(b bool)`

 SetMaterialNil sets the value for Material to be an explicit nil

### UnsetMaterial
`func (o *ListMeTemplatesItem) UnsetMaterial()`

UnsetMaterial ensures that no value is present for Material, not even an explicit nil
### GetStyle

`func (o *ListMeTemplatesItem) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ListMeTemplatesItem) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ListMeTemplatesItem) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ListMeTemplatesItem) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *ListMeTemplatesItem) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *ListMeTemplatesItem) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetPattern

`func (o *ListMeTemplatesItem) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ListMeTemplatesItem) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ListMeTemplatesItem) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *ListMeTemplatesItem) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPatternNil

`func (o *ListMeTemplatesItem) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *ListMeTemplatesItem) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil
### GetDepartment

`func (o *ListMeTemplatesItem) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ListMeTemplatesItem) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ListMeTemplatesItem) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ListMeTemplatesItem) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *ListMeTemplatesItem) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *ListMeTemplatesItem) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetGender

`func (o *ListMeTemplatesItem) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ListMeTemplatesItem) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ListMeTemplatesItem) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ListMeTemplatesItem) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ListMeTemplatesItem) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ListMeTemplatesItem) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetItemType

`func (o *ListMeTemplatesItem) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ListMeTemplatesItem) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ListMeTemplatesItem) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *ListMeTemplatesItem) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *ListMeTemplatesItem) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *ListMeTemplatesItem) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetSizeSystem

`func (o *ListMeTemplatesItem) GetSizeSystem() string`

GetSizeSystem returns the SizeSystem field if non-nil, zero value otherwise.

### GetSizeSystemOk

`func (o *ListMeTemplatesItem) GetSizeSystemOk() (*string, bool)`

GetSizeSystemOk returns a tuple with the SizeSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSystem

`func (o *ListMeTemplatesItem) SetSizeSystem(v string)`

SetSizeSystem sets SizeSystem field to given value.

### HasSizeSystem

`func (o *ListMeTemplatesItem) HasSizeSystem() bool`

HasSizeSystem returns a boolean if a field has been set.

### SetSizeSystemNil

`func (o *ListMeTemplatesItem) SetSizeSystemNil(b bool)`

 SetSizeSystemNil sets the value for SizeSystem to be an explicit nil

### UnsetSizeSystem
`func (o *ListMeTemplatesItem) UnsetSizeSystem()`

UnsetSizeSystem ensures that no value is present for SizeSystem, not even an explicit nil
### GetColor

`func (o *ListMeTemplatesItem) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ListMeTemplatesItem) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ListMeTemplatesItem) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *ListMeTemplatesItem) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *ListMeTemplatesItem) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *ListMeTemplatesItem) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetTags

`func (o *ListMeTemplatesItem) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListMeTemplatesItem) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListMeTemplatesItem) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ListMeTemplatesItem) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *ListMeTemplatesItem) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *ListMeTemplatesItem) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetNotes

`func (o *ListMeTemplatesItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ListMeTemplatesItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ListMeTemplatesItem) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ListMeTemplatesItem) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ListMeTemplatesItem) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ListMeTemplatesItem) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetWeightOz

`func (o *ListMeTemplatesItem) GetWeightOz() float32`

GetWeightOz returns the WeightOz field if non-nil, zero value otherwise.

### GetWeightOzOk

`func (o *ListMeTemplatesItem) GetWeightOzOk() (*float32, bool)`

GetWeightOzOk returns a tuple with the WeightOz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightOz

`func (o *ListMeTemplatesItem) SetWeightOz(v float32)`

SetWeightOz sets WeightOz field to given value.

### HasWeightOz

`func (o *ListMeTemplatesItem) HasWeightOz() bool`

HasWeightOz returns a boolean if a field has been set.

### SetWeightOzNil

`func (o *ListMeTemplatesItem) SetWeightOzNil(b bool)`

 SetWeightOzNil sets the value for WeightOz to be an explicit nil

### UnsetWeightOz
`func (o *ListMeTemplatesItem) UnsetWeightOz()`

UnsetWeightOz ensures that no value is present for WeightOz, not even an explicit nil
### GetIsDefault

`func (o *ListMeTemplatesItem) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ListMeTemplatesItem) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ListMeTemplatesItem) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetDescription

`func (o *ListMeTemplatesItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListMeTemplatesItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListMeTemplatesItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListMeTemplatesItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListMeTemplatesItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListMeTemplatesItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTitle

`func (o *ListMeTemplatesItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListMeTemplatesItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListMeTemplatesItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListMeTemplatesItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ListMeTemplatesItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ListMeTemplatesItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetScope

`func (o *ListMeTemplatesItem) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ListMeTemplatesItem) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ListMeTemplatesItem) SetScope(v string)`

SetScope sets Scope field to given value.


### GetShareToken

`func (o *ListMeTemplatesItem) GetShareToken() string`

GetShareToken returns the ShareToken field if non-nil, zero value otherwise.

### GetShareTokenOk

`func (o *ListMeTemplatesItem) GetShareTokenOk() (*string, bool)`

GetShareTokenOk returns a tuple with the ShareToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareToken

`func (o *ListMeTemplatesItem) SetShareToken(v string)`

SetShareToken sets ShareToken field to given value.

### HasShareToken

`func (o *ListMeTemplatesItem) HasShareToken() bool`

HasShareToken returns a boolean if a field has been set.

### SetShareTokenNil

`func (o *ListMeTemplatesItem) SetShareTokenNil(b bool)`

 SetShareTokenNil sets the value for ShareToken to be an explicit nil

### UnsetShareToken
`func (o *ListMeTemplatesItem) UnsetShareToken()`

UnsetShareToken ensures that no value is present for ShareToken, not even an explicit nil
### GetDescriptionVariants

`func (o *ListMeTemplatesItem) GetDescriptionVariants() []string`

GetDescriptionVariants returns the DescriptionVariants field if non-nil, zero value otherwise.

### GetDescriptionVariantsOk

`func (o *ListMeTemplatesItem) GetDescriptionVariantsOk() (*[]string, bool)`

GetDescriptionVariantsOk returns a tuple with the DescriptionVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionVariants

`func (o *ListMeTemplatesItem) SetDescriptionVariants(v []string)`

SetDescriptionVariants sets DescriptionVariants field to given value.

### HasDescriptionVariants

`func (o *ListMeTemplatesItem) HasDescriptionVariants() bool`

HasDescriptionVariants returns a boolean if a field has been set.

### SetDescriptionVariantsNil

`func (o *ListMeTemplatesItem) SetDescriptionVariantsNil(b bool)`

 SetDescriptionVariantsNil sets the value for DescriptionVariants to be an explicit nil

### UnsetDescriptionVariants
`func (o *ListMeTemplatesItem) UnsetDescriptionVariants()`

UnsetDescriptionVariants ensures that no value is present for DescriptionVariants, not even an explicit nil
### GetTitleVariants

`func (o *ListMeTemplatesItem) GetTitleVariants() []string`

GetTitleVariants returns the TitleVariants field if non-nil, zero value otherwise.

### GetTitleVariantsOk

`func (o *ListMeTemplatesItem) GetTitleVariantsOk() (*[]string, bool)`

GetTitleVariantsOk returns a tuple with the TitleVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleVariants

`func (o *ListMeTemplatesItem) SetTitleVariants(v []string)`

SetTitleVariants sets TitleVariants field to given value.

### HasTitleVariants

`func (o *ListMeTemplatesItem) HasTitleVariants() bool`

HasTitleVariants returns a boolean if a field has been set.

### SetTitleVariantsNil

`func (o *ListMeTemplatesItem) SetTitleVariantsNil(b bool)`

 SetTitleVariantsNil sets the value for TitleVariants to be an explicit nil

### UnsetTitleVariants
`func (o *ListMeTemplatesItem) UnsetTitleVariants()`

UnsetTitleVariants ensures that no value is present for TitleVariants, not even an explicit nil
### GetDefaultForCategory

`func (o *ListMeTemplatesItem) GetDefaultForCategory() string`

GetDefaultForCategory returns the DefaultForCategory field if non-nil, zero value otherwise.

### GetDefaultForCategoryOk

`func (o *ListMeTemplatesItem) GetDefaultForCategoryOk() (*string, bool)`

GetDefaultForCategoryOk returns a tuple with the DefaultForCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForCategory

`func (o *ListMeTemplatesItem) SetDefaultForCategory(v string)`

SetDefaultForCategory sets DefaultForCategory field to given value.

### HasDefaultForCategory

`func (o *ListMeTemplatesItem) HasDefaultForCategory() bool`

HasDefaultForCategory returns a boolean if a field has been set.

### SetDefaultForCategoryNil

`func (o *ListMeTemplatesItem) SetDefaultForCategoryNil(b bool)`

 SetDefaultForCategoryNil sets the value for DefaultForCategory to be an explicit nil

### UnsetDefaultForCategory
`func (o *ListMeTemplatesItem) UnsetDefaultForCategory()`

UnsetDefaultForCategory ensures that no value is present for DefaultForCategory, not even an explicit nil
### GetSortOrder

`func (o *ListMeTemplatesItem) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ListMeTemplatesItem) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ListMeTemplatesItem) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


