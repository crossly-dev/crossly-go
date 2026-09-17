# GetTaxonomyCategoryResponseCategories

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Label** | **string** |  | 
**Leaf** | **bool** |  | 
**ParentId** | Pointer to **NullableString** |  | [optional] 
**Path** | Pointer to **NullableString** | Root-first ancestor path including the leaf, when the caller supplied  it. Suggestions carry it; the tree walk does not. | [optional] 

## Methods

### NewGetTaxonomyCategoryResponseCategories

`func NewGetTaxonomyCategoryResponseCategories(id string, label string, leaf bool, ) *GetTaxonomyCategoryResponseCategories`

NewGetTaxonomyCategoryResponseCategories instantiates a new GetTaxonomyCategoryResponseCategories object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaxonomyCategoryResponseCategoriesWithDefaults

`func NewGetTaxonomyCategoryResponseCategoriesWithDefaults() *GetTaxonomyCategoryResponseCategories`

NewGetTaxonomyCategoryResponseCategoriesWithDefaults instantiates a new GetTaxonomyCategoryResponseCategories object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetTaxonomyCategoryResponseCategories) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetTaxonomyCategoryResponseCategories) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetTaxonomyCategoryResponseCategories) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *GetTaxonomyCategoryResponseCategories) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetTaxonomyCategoryResponseCategories) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetTaxonomyCategoryResponseCategories) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetLeaf

`func (o *GetTaxonomyCategoryResponseCategories) GetLeaf() bool`

GetLeaf returns the Leaf field if non-nil, zero value otherwise.

### GetLeafOk

`func (o *GetTaxonomyCategoryResponseCategories) GetLeafOk() (*bool, bool)`

GetLeafOk returns a tuple with the Leaf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaf

`func (o *GetTaxonomyCategoryResponseCategories) SetLeaf(v bool)`

SetLeaf sets Leaf field to given value.


### GetParentId

`func (o *GetTaxonomyCategoryResponseCategories) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GetTaxonomyCategoryResponseCategories) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GetTaxonomyCategoryResponseCategories) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GetTaxonomyCategoryResponseCategories) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### SetParentIdNil

`func (o *GetTaxonomyCategoryResponseCategories) SetParentIdNil(b bool)`

 SetParentIdNil sets the value for ParentId to be an explicit nil

### UnsetParentId
`func (o *GetTaxonomyCategoryResponseCategories) UnsetParentId()`

UnsetParentId ensures that no value is present for ParentId, not even an explicit nil
### GetPath

`func (o *GetTaxonomyCategoryResponseCategories) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GetTaxonomyCategoryResponseCategories) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GetTaxonomyCategoryResponseCategories) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *GetTaxonomyCategoryResponseCategories) HasPath() bool`

HasPath returns a boolean if a field has been set.

### SetPathNil

`func (o *GetTaxonomyCategoryResponseCategories) SetPathNil(b bool)`

 SetPathNil sets the value for Path to be an explicit nil

### UnsetPath
`func (o *GetTaxonomyCategoryResponseCategories) UnsetPath()`

UnsetPath ensures that no value is present for Path, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


