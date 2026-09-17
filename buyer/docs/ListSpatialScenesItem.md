# ListSpatialScenesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | **string** |  | 
**Name** | **string** |  | 
**CategorySlug** | Pointer to **NullableString** |  | [optional] 
**Visibility** | **string** |  | 
**PublicSlug** | Pointer to **NullableString** |  | [optional] 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewListSpatialScenesItem

`func NewListSpatialScenesItem(id string, kind string, name string, visibility string, updatedAt time.Time, ) *ListSpatialScenesItem`

NewListSpatialScenesItem instantiates a new ListSpatialScenesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSpatialScenesItemWithDefaults

`func NewListSpatialScenesItemWithDefaults() *ListSpatialScenesItem`

NewListSpatialScenesItemWithDefaults instantiates a new ListSpatialScenesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSpatialScenesItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSpatialScenesItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSpatialScenesItem) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *ListSpatialScenesItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListSpatialScenesItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListSpatialScenesItem) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *ListSpatialScenesItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSpatialScenesItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSpatialScenesItem) SetName(v string)`

SetName sets Name field to given value.


### GetCategorySlug

`func (o *ListSpatialScenesItem) GetCategorySlug() string`

GetCategorySlug returns the CategorySlug field if non-nil, zero value otherwise.

### GetCategorySlugOk

`func (o *ListSpatialScenesItem) GetCategorySlugOk() (*string, bool)`

GetCategorySlugOk returns a tuple with the CategorySlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySlug

`func (o *ListSpatialScenesItem) SetCategorySlug(v string)`

SetCategorySlug sets CategorySlug field to given value.

### HasCategorySlug

`func (o *ListSpatialScenesItem) HasCategorySlug() bool`

HasCategorySlug returns a boolean if a field has been set.

### SetCategorySlugNil

`func (o *ListSpatialScenesItem) SetCategorySlugNil(b bool)`

 SetCategorySlugNil sets the value for CategorySlug to be an explicit nil

### UnsetCategorySlug
`func (o *ListSpatialScenesItem) UnsetCategorySlug()`

UnsetCategorySlug ensures that no value is present for CategorySlug, not even an explicit nil
### GetVisibility

`func (o *ListSpatialScenesItem) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListSpatialScenesItem) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListSpatialScenesItem) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.


### GetPublicSlug

`func (o *ListSpatialScenesItem) GetPublicSlug() string`

GetPublicSlug returns the PublicSlug field if non-nil, zero value otherwise.

### GetPublicSlugOk

`func (o *ListSpatialScenesItem) GetPublicSlugOk() (*string, bool)`

GetPublicSlugOk returns a tuple with the PublicSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicSlug

`func (o *ListSpatialScenesItem) SetPublicSlug(v string)`

SetPublicSlug sets PublicSlug field to given value.

### HasPublicSlug

`func (o *ListSpatialScenesItem) HasPublicSlug() bool`

HasPublicSlug returns a boolean if a field has been set.

### SetPublicSlugNil

`func (o *ListSpatialScenesItem) SetPublicSlugNil(b bool)`

 SetPublicSlugNil sets the value for PublicSlug to be an explicit nil

### UnsetPublicSlug
`func (o *ListSpatialScenesItem) UnsetPublicSlug()`

UnsetPublicSlug ensures that no value is present for PublicSlug, not even an explicit nil
### GetUpdatedAt

`func (o *ListSpatialScenesItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListSpatialScenesItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListSpatialScenesItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


