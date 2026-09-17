# UpdateMeTemplateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Scope** | **string** |  | 
**Name** | **string** |  | 
**Notes** | Pointer to **NullableString** | Optional short blurb the seller can attach to remember what it&#39;s for. | [optional] 
**Description** | Pointer to **NullableString** | Primary description body. For scope&#x3D;&#39;description&#39; this is the  snippet body; for scope&#x3D;&#39;listing&#39; this is the default description  the seller wants pre-filled. | [optional] 
**DescriptionVariants** | Pointer to **[]string** | A/B variants for description. Populated for scope&#x3D;&#39;listing&#39;;  typically null for scope&#x3D;&#39;description&#39; (a snippet is one string). | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**TitleVariants** | Pointer to **[]string** |  | [optional] 
**Brand** | Pointer to **NullableString** |  | [optional] 
**Condition** | Pointer to **NullableString** | Master condition enum — new/like_new/good/fair/poor. | [optional] 
**Color** | Pointer to **NullableString** |  | [optional] 
**Material** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableString** |  | [optional] 
**SizeSystem** | Pointer to **NullableString** |  | [optional] 
**WeightOz** | Pointer to **NullableFloat32** |  | [optional] 
**Department** | Pointer to **NullableString** |  | [optional] 
**Gender** | Pointer to **NullableString** |  | [optional] 
**Style** | Pointer to **NullableString** |  | [optional] 
**Pattern** | Pointer to **NullableString** |  | [optional] 
**ItemType** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**DefaultForCategory** | Pointer to **NullableString** | When set, form&#39;s category picker prompts \&quot;Use your default for  this category\&quot; on match. | [optional] 
**IsDefault** | **bool** |  | 
**ShareToken** | Pointer to **NullableString** | URL-safe random token. Populated by POST /me/templates/:id/share;  the /public/templates/:token route surfaces a read-only view any  visitor can browse + import. | [optional] 
**SortOrder** | **float32** | Snippet ordering — kept for scope&#x3D;&#39;description&#39; back-compat with  the description_templates.sort_order behavior. | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewUpdateMeTemplateResponse

`func NewUpdateMeTemplateResponse(id string, userId string, scope string, name string, isDefault bool, sortOrder float32, createdAt time.Time, updatedAt time.Time, ) *UpdateMeTemplateResponse`

NewUpdateMeTemplateResponse instantiates a new UpdateMeTemplateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMeTemplateResponseWithDefaults

`func NewUpdateMeTemplateResponseWithDefaults() *UpdateMeTemplateResponse`

NewUpdateMeTemplateResponseWithDefaults instantiates a new UpdateMeTemplateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateMeTemplateResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateMeTemplateResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateMeTemplateResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *UpdateMeTemplateResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateMeTemplateResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateMeTemplateResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetScope

`func (o *UpdateMeTemplateResponse) GetScope() string`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *UpdateMeTemplateResponse) GetScopeOk() (*string, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *UpdateMeTemplateResponse) SetScope(v string)`

SetScope sets Scope field to given value.


### GetName

`func (o *UpdateMeTemplateResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateMeTemplateResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateMeTemplateResponse) SetName(v string)`

SetName sets Name field to given value.


### GetNotes

`func (o *UpdateMeTemplateResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateMeTemplateResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateMeTemplateResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateMeTemplateResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *UpdateMeTemplateResponse) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *UpdateMeTemplateResponse) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetDescription

`func (o *UpdateMeTemplateResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateMeTemplateResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateMeTemplateResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateMeTemplateResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateMeTemplateResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateMeTemplateResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDescriptionVariants

`func (o *UpdateMeTemplateResponse) GetDescriptionVariants() []string`

GetDescriptionVariants returns the DescriptionVariants field if non-nil, zero value otherwise.

### GetDescriptionVariantsOk

`func (o *UpdateMeTemplateResponse) GetDescriptionVariantsOk() (*[]string, bool)`

GetDescriptionVariantsOk returns a tuple with the DescriptionVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionVariants

`func (o *UpdateMeTemplateResponse) SetDescriptionVariants(v []string)`

SetDescriptionVariants sets DescriptionVariants field to given value.

### HasDescriptionVariants

`func (o *UpdateMeTemplateResponse) HasDescriptionVariants() bool`

HasDescriptionVariants returns a boolean if a field has been set.

### SetDescriptionVariantsNil

`func (o *UpdateMeTemplateResponse) SetDescriptionVariantsNil(b bool)`

 SetDescriptionVariantsNil sets the value for DescriptionVariants to be an explicit nil

### UnsetDescriptionVariants
`func (o *UpdateMeTemplateResponse) UnsetDescriptionVariants()`

UnsetDescriptionVariants ensures that no value is present for DescriptionVariants, not even an explicit nil
### GetTitle

`func (o *UpdateMeTemplateResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *UpdateMeTemplateResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *UpdateMeTemplateResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *UpdateMeTemplateResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *UpdateMeTemplateResponse) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *UpdateMeTemplateResponse) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetTitleVariants

`func (o *UpdateMeTemplateResponse) GetTitleVariants() []string`

GetTitleVariants returns the TitleVariants field if non-nil, zero value otherwise.

### GetTitleVariantsOk

`func (o *UpdateMeTemplateResponse) GetTitleVariantsOk() (*[]string, bool)`

GetTitleVariantsOk returns a tuple with the TitleVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitleVariants

`func (o *UpdateMeTemplateResponse) SetTitleVariants(v []string)`

SetTitleVariants sets TitleVariants field to given value.

### HasTitleVariants

`func (o *UpdateMeTemplateResponse) HasTitleVariants() bool`

HasTitleVariants returns a boolean if a field has been set.

### SetTitleVariantsNil

`func (o *UpdateMeTemplateResponse) SetTitleVariantsNil(b bool)`

 SetTitleVariantsNil sets the value for TitleVariants to be an explicit nil

### UnsetTitleVariants
`func (o *UpdateMeTemplateResponse) UnsetTitleVariants()`

UnsetTitleVariants ensures that no value is present for TitleVariants, not even an explicit nil
### GetBrand

`func (o *UpdateMeTemplateResponse) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *UpdateMeTemplateResponse) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *UpdateMeTemplateResponse) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *UpdateMeTemplateResponse) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *UpdateMeTemplateResponse) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *UpdateMeTemplateResponse) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetCondition

`func (o *UpdateMeTemplateResponse) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *UpdateMeTemplateResponse) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *UpdateMeTemplateResponse) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *UpdateMeTemplateResponse) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *UpdateMeTemplateResponse) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *UpdateMeTemplateResponse) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetColor

`func (o *UpdateMeTemplateResponse) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *UpdateMeTemplateResponse) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *UpdateMeTemplateResponse) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *UpdateMeTemplateResponse) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *UpdateMeTemplateResponse) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *UpdateMeTemplateResponse) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetMaterial

`func (o *UpdateMeTemplateResponse) GetMaterial() string`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *UpdateMeTemplateResponse) GetMaterialOk() (*string, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *UpdateMeTemplateResponse) SetMaterial(v string)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *UpdateMeTemplateResponse) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### SetMaterialNil

`func (o *UpdateMeTemplateResponse) SetMaterialNil(b bool)`

 SetMaterialNil sets the value for Material to be an explicit nil

### UnsetMaterial
`func (o *UpdateMeTemplateResponse) UnsetMaterial()`

UnsetMaterial ensures that no value is present for Material, not even an explicit nil
### GetSize

`func (o *UpdateMeTemplateResponse) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *UpdateMeTemplateResponse) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *UpdateMeTemplateResponse) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *UpdateMeTemplateResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *UpdateMeTemplateResponse) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *UpdateMeTemplateResponse) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetSizeSystem

`func (o *UpdateMeTemplateResponse) GetSizeSystem() string`

GetSizeSystem returns the SizeSystem field if non-nil, zero value otherwise.

### GetSizeSystemOk

`func (o *UpdateMeTemplateResponse) GetSizeSystemOk() (*string, bool)`

GetSizeSystemOk returns a tuple with the SizeSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSystem

`func (o *UpdateMeTemplateResponse) SetSizeSystem(v string)`

SetSizeSystem sets SizeSystem field to given value.

### HasSizeSystem

`func (o *UpdateMeTemplateResponse) HasSizeSystem() bool`

HasSizeSystem returns a boolean if a field has been set.

### SetSizeSystemNil

`func (o *UpdateMeTemplateResponse) SetSizeSystemNil(b bool)`

 SetSizeSystemNil sets the value for SizeSystem to be an explicit nil

### UnsetSizeSystem
`func (o *UpdateMeTemplateResponse) UnsetSizeSystem()`

UnsetSizeSystem ensures that no value is present for SizeSystem, not even an explicit nil
### GetWeightOz

`func (o *UpdateMeTemplateResponse) GetWeightOz() float32`

GetWeightOz returns the WeightOz field if non-nil, zero value otherwise.

### GetWeightOzOk

`func (o *UpdateMeTemplateResponse) GetWeightOzOk() (*float32, bool)`

GetWeightOzOk returns a tuple with the WeightOz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightOz

`func (o *UpdateMeTemplateResponse) SetWeightOz(v float32)`

SetWeightOz sets WeightOz field to given value.

### HasWeightOz

`func (o *UpdateMeTemplateResponse) HasWeightOz() bool`

HasWeightOz returns a boolean if a field has been set.

### SetWeightOzNil

`func (o *UpdateMeTemplateResponse) SetWeightOzNil(b bool)`

 SetWeightOzNil sets the value for WeightOz to be an explicit nil

### UnsetWeightOz
`func (o *UpdateMeTemplateResponse) UnsetWeightOz()`

UnsetWeightOz ensures that no value is present for WeightOz, not even an explicit nil
### GetDepartment

`func (o *UpdateMeTemplateResponse) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *UpdateMeTemplateResponse) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *UpdateMeTemplateResponse) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *UpdateMeTemplateResponse) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *UpdateMeTemplateResponse) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *UpdateMeTemplateResponse) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetGender

`func (o *UpdateMeTemplateResponse) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *UpdateMeTemplateResponse) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *UpdateMeTemplateResponse) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *UpdateMeTemplateResponse) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *UpdateMeTemplateResponse) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *UpdateMeTemplateResponse) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetStyle

`func (o *UpdateMeTemplateResponse) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *UpdateMeTemplateResponse) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *UpdateMeTemplateResponse) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *UpdateMeTemplateResponse) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *UpdateMeTemplateResponse) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *UpdateMeTemplateResponse) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetPattern

`func (o *UpdateMeTemplateResponse) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *UpdateMeTemplateResponse) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *UpdateMeTemplateResponse) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *UpdateMeTemplateResponse) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPatternNil

`func (o *UpdateMeTemplateResponse) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *UpdateMeTemplateResponse) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil
### GetItemType

`func (o *UpdateMeTemplateResponse) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *UpdateMeTemplateResponse) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *UpdateMeTemplateResponse) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *UpdateMeTemplateResponse) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *UpdateMeTemplateResponse) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *UpdateMeTemplateResponse) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetTags

`func (o *UpdateMeTemplateResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *UpdateMeTemplateResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *UpdateMeTemplateResponse) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *UpdateMeTemplateResponse) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *UpdateMeTemplateResponse) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *UpdateMeTemplateResponse) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetDefaultForCategory

`func (o *UpdateMeTemplateResponse) GetDefaultForCategory() string`

GetDefaultForCategory returns the DefaultForCategory field if non-nil, zero value otherwise.

### GetDefaultForCategoryOk

`func (o *UpdateMeTemplateResponse) GetDefaultForCategoryOk() (*string, bool)`

GetDefaultForCategoryOk returns a tuple with the DefaultForCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForCategory

`func (o *UpdateMeTemplateResponse) SetDefaultForCategory(v string)`

SetDefaultForCategory sets DefaultForCategory field to given value.

### HasDefaultForCategory

`func (o *UpdateMeTemplateResponse) HasDefaultForCategory() bool`

HasDefaultForCategory returns a boolean if a field has been set.

### SetDefaultForCategoryNil

`func (o *UpdateMeTemplateResponse) SetDefaultForCategoryNil(b bool)`

 SetDefaultForCategoryNil sets the value for DefaultForCategory to be an explicit nil

### UnsetDefaultForCategory
`func (o *UpdateMeTemplateResponse) UnsetDefaultForCategory()`

UnsetDefaultForCategory ensures that no value is present for DefaultForCategory, not even an explicit nil
### GetIsDefault

`func (o *UpdateMeTemplateResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UpdateMeTemplateResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UpdateMeTemplateResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetShareToken

`func (o *UpdateMeTemplateResponse) GetShareToken() string`

GetShareToken returns the ShareToken field if non-nil, zero value otherwise.

### GetShareTokenOk

`func (o *UpdateMeTemplateResponse) GetShareTokenOk() (*string, bool)`

GetShareTokenOk returns a tuple with the ShareToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShareToken

`func (o *UpdateMeTemplateResponse) SetShareToken(v string)`

SetShareToken sets ShareToken field to given value.

### HasShareToken

`func (o *UpdateMeTemplateResponse) HasShareToken() bool`

HasShareToken returns a boolean if a field has been set.

### SetShareTokenNil

`func (o *UpdateMeTemplateResponse) SetShareTokenNil(b bool)`

 SetShareTokenNil sets the value for ShareToken to be an explicit nil

### UnsetShareToken
`func (o *UpdateMeTemplateResponse) UnsetShareToken()`

UnsetShareToken ensures that no value is present for ShareToken, not even an explicit nil
### GetSortOrder

`func (o *UpdateMeTemplateResponse) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *UpdateMeTemplateResponse) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *UpdateMeTemplateResponse) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.


### GetCreatedAt

`func (o *UpdateMeTemplateResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateMeTemplateResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateMeTemplateResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UpdateMeTemplateResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdateMeTemplateResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdateMeTemplateResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


