# GetSpatialPublicResponseProfile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CategorySlug** | **string** | Matches &#x60;crossly_market_categories.slug&#x60;, or &#39;*&#39; for the fallback. | 
**Label** | **string** | Room title, shown in the switcher. | 
**Tagline** | **string** | One line of why this space is shaped the way it is. Surfaced in the UI. | 
**Presentation** | **string** |  | 
**ItemModelVariants** | Pointer to **[]string** | Other meshes items in this category may be drawn as, chosen PER ITEM.    One model per category is right for a card room, where every object is the  same object. It is wrong for a wardrobe: a rail holds tees and jeans and  jackets, and drawing all of them as a tee would be a worse lie than the  flat quad it replaced, because a wrong SHAPE reads as information.    The renderer picks from &#x60;[itemModel, ...itemModelVariants]&#x60; using the  item&#39;s title — the same heuristic &#x60;silhouetteFor&#x60; already uses to choose a  garment outline, and for the same reason: the title is the only signal  present on every item, and a wrong guess costs a slightly odd shape rather  than the wrong item. &#x60;itemModel&#x60; is the fallback when nothing matches.    Deliberately NOT fuzzy-matched against a catalog — this picks a SHAPE, not  an identity. See docs/IDENTIFIER-FIRST.md for where that line sits. | [optional] 
**ItemSize** | [**GetSpatialPublicResponseProfileItemSize**](GetSpatialPublicResponseProfileItemSize.md) |  | 
**Containers** | [**[]GetSpatialPublicResponseProfileContainers**](GetSpatialPublicResponseProfileContainers.md) | The container ladder, OUTERMOST FIRST. A binder holds pages, a page holds  cards. The solver walks this to decide what to create next when the  current container fills up. | 
**DefaultGroupBy** | **[]string** | Default grouping, in order. Each level becomes a divider or a container. | 
**DefaultSortBy** | **[]string** | Default ordering inside a group. | 
**GradedVariant** | Pointer to **NullableString** | A SEPARATE profile for graded/sealed copies of the same category.    This is not a flourish. Nobody puts a slabbed card in a binder — it does  not fit and it would be vandalism. Graded cards go on a wall, raw cards go  in pockets, and a space that ignores that is immediately wrong to the only  people who would use it. Same for CGC comics. | [optional] 

## Methods

### NewGetSpatialPublicResponseProfile

`func NewGetSpatialPublicResponseProfile(categorySlug string, label string, tagline string, presentation string, itemSize GetSpatialPublicResponseProfileItemSize, containers []GetSpatialPublicResponseProfileContainers, defaultGroupBy []string, defaultSortBy []string, ) *GetSpatialPublicResponseProfile`

NewGetSpatialPublicResponseProfile instantiates a new GetSpatialPublicResponseProfile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSpatialPublicResponseProfileWithDefaults

`func NewGetSpatialPublicResponseProfileWithDefaults() *GetSpatialPublicResponseProfile`

NewGetSpatialPublicResponseProfileWithDefaults instantiates a new GetSpatialPublicResponseProfile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategorySlug

`func (o *GetSpatialPublicResponseProfile) GetCategorySlug() string`

GetCategorySlug returns the CategorySlug field if non-nil, zero value otherwise.

### GetCategorySlugOk

`func (o *GetSpatialPublicResponseProfile) GetCategorySlugOk() (*string, bool)`

GetCategorySlugOk returns a tuple with the CategorySlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySlug

`func (o *GetSpatialPublicResponseProfile) SetCategorySlug(v string)`

SetCategorySlug sets CategorySlug field to given value.


### GetLabel

`func (o *GetSpatialPublicResponseProfile) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetSpatialPublicResponseProfile) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetSpatialPublicResponseProfile) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetTagline

`func (o *GetSpatialPublicResponseProfile) GetTagline() string`

GetTagline returns the Tagline field if non-nil, zero value otherwise.

### GetTaglineOk

`func (o *GetSpatialPublicResponseProfile) GetTaglineOk() (*string, bool)`

GetTaglineOk returns a tuple with the Tagline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTagline

`func (o *GetSpatialPublicResponseProfile) SetTagline(v string)`

SetTagline sets Tagline field to given value.


### GetPresentation

`func (o *GetSpatialPublicResponseProfile) GetPresentation() string`

GetPresentation returns the Presentation field if non-nil, zero value otherwise.

### GetPresentationOk

`func (o *GetSpatialPublicResponseProfile) GetPresentationOk() (*string, bool)`

GetPresentationOk returns a tuple with the Presentation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresentation

`func (o *GetSpatialPublicResponseProfile) SetPresentation(v string)`

SetPresentation sets Presentation field to given value.


### GetItemModelVariants

`func (o *GetSpatialPublicResponseProfile) GetItemModelVariants() []string`

GetItemModelVariants returns the ItemModelVariants field if non-nil, zero value otherwise.

### GetItemModelVariantsOk

`func (o *GetSpatialPublicResponseProfile) GetItemModelVariantsOk() (*[]string, bool)`

GetItemModelVariantsOk returns a tuple with the ItemModelVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemModelVariants

`func (o *GetSpatialPublicResponseProfile) SetItemModelVariants(v []string)`

SetItemModelVariants sets ItemModelVariants field to given value.

### HasItemModelVariants

`func (o *GetSpatialPublicResponseProfile) HasItemModelVariants() bool`

HasItemModelVariants returns a boolean if a field has been set.

### SetItemModelVariantsNil

`func (o *GetSpatialPublicResponseProfile) SetItemModelVariantsNil(b bool)`

 SetItemModelVariantsNil sets the value for ItemModelVariants to be an explicit nil

### UnsetItemModelVariants
`func (o *GetSpatialPublicResponseProfile) UnsetItemModelVariants()`

UnsetItemModelVariants ensures that no value is present for ItemModelVariants, not even an explicit nil
### GetItemSize

`func (o *GetSpatialPublicResponseProfile) GetItemSize() GetSpatialPublicResponseProfileItemSize`

GetItemSize returns the ItemSize field if non-nil, zero value otherwise.

### GetItemSizeOk

`func (o *GetSpatialPublicResponseProfile) GetItemSizeOk() (*GetSpatialPublicResponseProfileItemSize, bool)`

GetItemSizeOk returns a tuple with the ItemSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSize

`func (o *GetSpatialPublicResponseProfile) SetItemSize(v GetSpatialPublicResponseProfileItemSize)`

SetItemSize sets ItemSize field to given value.


### GetContainers

`func (o *GetSpatialPublicResponseProfile) GetContainers() []GetSpatialPublicResponseProfileContainers`

GetContainers returns the Containers field if non-nil, zero value otherwise.

### GetContainersOk

`func (o *GetSpatialPublicResponseProfile) GetContainersOk() (*[]GetSpatialPublicResponseProfileContainers, bool)`

GetContainersOk returns a tuple with the Containers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainers

`func (o *GetSpatialPublicResponseProfile) SetContainers(v []GetSpatialPublicResponseProfileContainers)`

SetContainers sets Containers field to given value.


### GetDefaultGroupBy

`func (o *GetSpatialPublicResponseProfile) GetDefaultGroupBy() []string`

GetDefaultGroupBy returns the DefaultGroupBy field if non-nil, zero value otherwise.

### GetDefaultGroupByOk

`func (o *GetSpatialPublicResponseProfile) GetDefaultGroupByOk() (*[]string, bool)`

GetDefaultGroupByOk returns a tuple with the DefaultGroupBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupBy

`func (o *GetSpatialPublicResponseProfile) SetDefaultGroupBy(v []string)`

SetDefaultGroupBy sets DefaultGroupBy field to given value.


### GetDefaultSortBy

`func (o *GetSpatialPublicResponseProfile) GetDefaultSortBy() []string`

GetDefaultSortBy returns the DefaultSortBy field if non-nil, zero value otherwise.

### GetDefaultSortByOk

`func (o *GetSpatialPublicResponseProfile) GetDefaultSortByOk() (*[]string, bool)`

GetDefaultSortByOk returns a tuple with the DefaultSortBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSortBy

`func (o *GetSpatialPublicResponseProfile) SetDefaultSortBy(v []string)`

SetDefaultSortBy sets DefaultSortBy field to given value.


### GetGradedVariant

`func (o *GetSpatialPublicResponseProfile) GetGradedVariant() string`

GetGradedVariant returns the GradedVariant field if non-nil, zero value otherwise.

### GetGradedVariantOk

`func (o *GetSpatialPublicResponseProfile) GetGradedVariantOk() (*string, bool)`

GetGradedVariantOk returns a tuple with the GradedVariant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradedVariant

`func (o *GetSpatialPublicResponseProfile) SetGradedVariant(v string)`

SetGradedVariant sets GradedVariant field to given value.

### HasGradedVariant

`func (o *GetSpatialPublicResponseProfile) HasGradedVariant() bool`

HasGradedVariant returns a boolean if a field has been set.

### SetGradedVariantNil

`func (o *GetSpatialPublicResponseProfile) SetGradedVariantNil(b bool)`

 SetGradedVariantNil sets the value for GradedVariant to be an explicit nil

### UnsetGradedVariant
`func (o *GetSpatialPublicResponseProfile) UnsetGradedVariant()`

UnsetGradedVariant ensures that no value is present for GradedVariant, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


