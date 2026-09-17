# ListVariationGroupsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Members** | [**[]ListVariationGroupsItemMembers**](ListVariationGroupsItemMembers.md) |  | 
**Rollup** | [**ListVariationGroupsItemRollup**](ListVariationGroupsItemRollup.md) |  | 
**Id** | **string** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListVariationGroupsItem

`func NewListVariationGroupsItem(members []ListVariationGroupsItemMembers, rollup ListVariationGroupsItemRollup, id string, ) *ListVariationGroupsItem`

NewListVariationGroupsItem instantiates a new ListVariationGroupsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVariationGroupsItemWithDefaults

`func NewListVariationGroupsItemWithDefaults() *ListVariationGroupsItem`

NewListVariationGroupsItemWithDefaults instantiates a new ListVariationGroupsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMembers

`func (o *ListVariationGroupsItem) GetMembers() []ListVariationGroupsItemMembers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *ListVariationGroupsItem) GetMembersOk() (*[]ListVariationGroupsItemMembers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *ListVariationGroupsItem) SetMembers(v []ListVariationGroupsItemMembers)`

SetMembers sets Members field to given value.


### GetRollup

`func (o *ListVariationGroupsItem) GetRollup() ListVariationGroupsItemRollup`

GetRollup returns the Rollup field if non-nil, zero value otherwise.

### GetRollupOk

`func (o *ListVariationGroupsItem) GetRollupOk() (*ListVariationGroupsItemRollup, bool)`

GetRollupOk returns a tuple with the Rollup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRollup

`func (o *ListVariationGroupsItem) SetRollup(v ListVariationGroupsItemRollup)`

SetRollup sets Rollup field to given value.


### GetId

`func (o *ListVariationGroupsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListVariationGroupsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListVariationGroupsItem) SetId(v string)`

SetId sets Id field to given value.


### GetTitle

`func (o *ListVariationGroupsItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListVariationGroupsItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListVariationGroupsItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListVariationGroupsItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ListVariationGroupsItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ListVariationGroupsItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


