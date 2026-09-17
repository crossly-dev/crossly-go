# CreateInventoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**NullableListInventoryItemCategory**](ListInventoryItemCategory.md) |  | [optional] 

## Methods

### NewCreateInventoryResponse

`func NewCreateInventoryResponse() *CreateInventoryResponse`

NewCreateInventoryResponse instantiates a new CreateInventoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInventoryResponseWithDefaults

`func NewCreateInventoryResponseWithDefaults() *CreateInventoryResponse`

NewCreateInventoryResponseWithDefaults instantiates a new CreateInventoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *CreateInventoryResponse) GetCategory() ListInventoryItemCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateInventoryResponse) GetCategoryOk() (*ListInventoryItemCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateInventoryResponse) SetCategory(v ListInventoryItemCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateInventoryResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *CreateInventoryResponse) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *CreateInventoryResponse) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


