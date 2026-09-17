# GetMeTemplateResponse

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

### NewGetMeTemplateResponse

`func NewGetMeTemplateResponse(id string, name string, createdAt time.Time, updatedAt time.Time, userId string, isDefault bool, scope string, sortOrder float32, ) *GetMeTemplateResponse`

NewGetMeTemplateResponse instantiates a new GetMeTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMeTemplateResponseWithDefaults

`func NewGetMeTemplateResponseWithDefaults() *GetMeTemplateResponse`

NewGetMeTemplateResponseWithDefaults instantiates a new GetMeTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMeTemplateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMeTemplateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMeTemplateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetBrand

`func (o *GetMeTemplateResponse) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *GetMeTemplateResponse) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *GetMeTemplateResponse) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *GetMeTemplateResponse) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *GetMeTemplateResponse) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *GetMeTemplateResponse) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetName

`func (o *GetMeTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMeTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMeTemplateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *GetMeTemplateResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetMeTemplateResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetMeTemplateResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetMeTemplateResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetMeTemplateResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetMeTemplateResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *GetMeTemplateResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetMeTemplateResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetMeTemplateResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCondition

`func (o *GetMeTemplateResponse) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GetMeTemplateResponse) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *GetMeTemplateResponse) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *GetMeTemplateResponse) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *GetMeTemplateResponse) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *GetMeTemplateResponse) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetSize

`func (o *GetMeTemplateResponse) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetMeTemplateResponse) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetMeTemplateResponse) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetMeTemplateResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *GetMeTemplateResponse) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *GetMeTemplateResponse) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetMaterial

`func (o *GetMeTemplateResponse) GetMaterial() string`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *GetMeTemplateResponse) GetMaterialOk() (*string, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *GetMeTemplateResponse) SetMaterial(v string)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *GetMeTemplateResponse) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### SetMaterialNil

`func (o *GetMeTemplateResponse) SetMaterialNil(b bool)`

 SetMaterialNil sets the value for Material to be an explicit nil

### UnsetMaterial
`func (o *GetMeTemplateResponse) UnsetMaterial()`

UnsetMaterial ensures that no value is present for Material, not even an explicit nil
### GetStyle

`func (o *GetMeTemplateResponse) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *GetMeTemplateResponse) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *GetMeTemplateResponse) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *GetMeTemplateResponse) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *GetMeTemplateResponse) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *GetMeTemplateResponse) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetPattern

`func (o *GetMeTemplateResponse) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *GetMeTemplateResponse) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *GetMeTemplateResponse) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *GetMeTemplateResponse) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPatternNil

`func (o *GetMeTemplateResponse) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *GetMeTemplateResponse) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil
### GetDepartment

`func (o *GetMeTemplateResponse) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *GetMeTemplateResponse) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *GetMeTemplateResponse) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *GetMeTemplateResponse) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *GetMeTemplateResponse) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *GetMeTemplateResponse) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetGender

`func (o *GetMeTemplateResponse) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *GetMeTemplateResponse) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *GetMeTemplateResponse) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *GetMeTemplateResponse) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *GetMeTemplateResponse) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *GetMeTemplateResponse) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetItemType

`func (o *GetMeTemplateResponse) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *GetMeTemplateResponse) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *GetMeTemplateResponse) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *GetMeTemplateResponse) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *GetMeTemplateResponse) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *GetMeTemplateResponse) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetSizeSystem

`func (o *GetMeTemplateResponse) GetSizeSystem() string`

GetSizeSystem returns the SizeSystem field if non-nil, zero value otherwise.

### GetSizeSystemOk

`func (o *GetMeTemplateResponse) GetSizeSystemOk() (*string, bool)`

GetSizeSystemOk returns a tuple with the SizeSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSystem

`func (o *GetMeTemplateResponse) SetSizeSystem(v string)`

SetSizeSystem sets SizeSystem field to given value.

### HasSizeSystem

`func (o *GetMeTemplateResponse) HasSizeSystem() bool`

HasSizeSystem returns a boolean if a field has been set.

### SetSizeSystemNil

`func (o *GetMeTemplateResponse) SetSizeSystemNil(b bool)`

 SetSizeSystemNil sets the value for SizeSystem to be an explicit nil

### UnsetSizeSystem
`func (o *GetMeTemplateResponse) UnsetSizeSystem()`

UnsetSizeSystem ensures that no value is present for SizeSystem, not even an explicit nil
### GetColor

`func (o *GetMeTemplateResponse) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *GetMeTemplateResponse) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *GetMeTemplateResponse) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *GetMeTemplateResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *GetMeTemplateResponse) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *GetMeTemplateResponse) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetTags

`func (o *GetMeTemplateResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetMeTemplateResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetMeTemplateResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GetMeTemplateResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *GetMeTemplateResponse) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *GetMeTemplateResponse) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetNotes

`func (o *GetMeTemplateResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetMeTemplateResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetMeTemplateResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetMeTemplateResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *GetMeTemplateResponse) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *GetMeTemplateResponse) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetWeightOz

`func (o *GetMeTemplateResponse) GetWeightOz() float32`

GetWeightOz returns the WeightOz field if non-nil, zero value otherwise.

### GetWeightOzOk

`func (o *GetMeTemplateResponse) GetWeightOzOk() (*float32, bool)`

GetWeightOzOk returns a tuple with the WeightOz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightOz

`func (o *GetMeTemplateResponse) SetWeightOz(v float32)`

SetWeightOz sets WeightOz field to given value.

### HasWeightOz

`func (o *GetMeTemplateResponse) HasWeightOz() bool`

HasWeightOz returns a boolean if a field has been set.

### SetWeightOzNil

`func (o *GetMeTemplateResponse) SetWeightOzNil(b bool)`

 SetWeightOzNil sets the value for WeightOz to be an explicit nil

### UnsetWeightOz
`func (o *GetMeTemplateResponse) UnsetWeightOz()`

UnsetWeightOz ensures that no value is present for WeightOz, not even an explicit nil
### GetIsDefault

`func (o *GetMeTemplateResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *GetMeTemplateResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *GetMeTemplateResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetDescription

`func (o *GetMeTemplateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetMeTemplateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetMeTemplateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetMeTemplateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetMeTemplateResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetMeTemplateResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTitle

`func (o *GetMeTemplateResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetMeTemplateResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetMeTemplateResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetMeTemplateResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *GetMeTemplateResponse) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *GetMeTemplateResponse) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetScope

`func (o *GetMeTemplateResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *GetMeTemplateResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *GetMeTemplateResponse) SetScope(v string)`

SetScope sets Scope field to given value.


### GetShareToken

`func (o *GetMeTemplateResponse) GetShareToken() string`

GetShareToken returns the ShareToken field if non-nil, zero value otherwise.

### GetShareTokenOk

`func (o *GetMeTemplateResponse) GetShareTokenOk() (*string, bool)`

GetShareTokenOk returns a tuple with the ShareToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareToken

`func (o *GetMeTemplateResponse) SetShareToken(v string)`

SetShareToken sets ShareToken field to given value.

### HasShareToken

`func (o *GetMeTemplateResponse) HasShareToken() bool`

HasShareToken returns a boolean if a field has been set.

### SetShareTokenNil

`func (o *GetMeTemplateResponse) SetShareTokenNil(b bool)`

 SetShareTokenNil sets the value for ShareToken to be an explicit nil

### UnsetShareToken
`func (o *GetMeTemplateResponse) UnsetShareToken()`

UnsetShareToken ensures that no value is present for ShareToken, not even an explicit nil
### GetDescriptionVariants

`func (o *GetMeTemplateResponse) GetDescriptionVariants() []string`

GetDescriptionVariants returns the DescriptionVariants field if non-nil, zero value otherwise.

### GetDescriptionVariantsOk

`func (o *GetMeTemplateResponse) GetDescriptionVariantsOk() (*[]string, bool)`

GetDescriptionVariantsOk returns a tuple with the DescriptionVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionVariants

`func (o *GetMeTemplateResponse) SetDescriptionVariants(v []string)`

SetDescriptionVariants sets DescriptionVariants field to given value.

### HasDescriptionVariants

`func (o *GetMeTemplateResponse) HasDescriptionVariants() bool`

HasDescriptionVariants returns a boolean if a field has been set.

### SetDescriptionVariantsNil

`func (o *GetMeTemplateResponse) SetDescriptionVariantsNil(b bool)`

 SetDescriptionVariantsNil sets the value for DescriptionVariants to be an explicit nil

### UnsetDescriptionVariants
`func (o *GetMeTemplateResponse) UnsetDescriptionVariants()`

UnsetDescriptionVariants ensures that no value is present for DescriptionVariants, not even an explicit nil
### GetTitleVariants

`func (o *GetMeTemplateResponse) GetTitleVariants() []string`

GetTitleVariants returns the TitleVariants field if non-nil, zero value otherwise.

### GetTitleVariantsOk

`func (o *GetMeTemplateResponse) GetTitleVariantsOk() (*[]string, bool)`

GetTitleVariantsOk returns a tuple with the TitleVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleVariants

`func (o *GetMeTemplateResponse) SetTitleVariants(v []string)`

SetTitleVariants sets TitleVariants field to given value.

### HasTitleVariants

`func (o *GetMeTemplateResponse) HasTitleVariants() bool`

HasTitleVariants returns a boolean if a field has been set.

### SetTitleVariantsNil

`func (o *GetMeTemplateResponse) SetTitleVariantsNil(b bool)`

 SetTitleVariantsNil sets the value for TitleVariants to be an explicit nil

### UnsetTitleVariants
`func (o *GetMeTemplateResponse) UnsetTitleVariants()`

UnsetTitleVariants ensures that no value is present for TitleVariants, not even an explicit nil
### GetDefaultForCategory

`func (o *GetMeTemplateResponse) GetDefaultForCategory() string`

GetDefaultForCategory returns the DefaultForCategory field if non-nil, zero value otherwise.

### GetDefaultForCategoryOk

`func (o *GetMeTemplateResponse) GetDefaultForCategoryOk() (*string, bool)`

GetDefaultForCategoryOk returns a tuple with the DefaultForCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForCategory

`func (o *GetMeTemplateResponse) SetDefaultForCategory(v string)`

SetDefaultForCategory sets DefaultForCategory field to given value.

### HasDefaultForCategory

`func (o *GetMeTemplateResponse) HasDefaultForCategory() bool`

HasDefaultForCategory returns a boolean if a field has been set.

### SetDefaultForCategoryNil

`func (o *GetMeTemplateResponse) SetDefaultForCategoryNil(b bool)`

 SetDefaultForCategoryNil sets the value for DefaultForCategory to be an explicit nil

### UnsetDefaultForCategory
`func (o *GetMeTemplateResponse) UnsetDefaultForCategory()`

UnsetDefaultForCategory ensures that no value is present for DefaultForCategory, not even an explicit nil
### GetSortOrder

`func (o *GetMeTemplateResponse) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *GetMeTemplateResponse) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *GetMeTemplateResponse) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


