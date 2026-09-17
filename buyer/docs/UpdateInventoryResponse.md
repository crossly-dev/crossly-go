# UpdateInventoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to [**NullableListInventoryItemCategory**](ListInventoryItemCategory.md) |  | [optional] 

## Methods

### NewUpdateInventoryResponse

`func NewUpdateInventoryResponse() *UpdateInventoryResponse`

NewUpdateInventoryResponse instantiates a new UpdateInventoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInventoryResponseWithDefaults

`func NewUpdateInventoryResponseWithDefaults() *UpdateInventoryResponse`

NewUpdateInventoryResponseWithDefaults instantiates a new UpdateInventoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *UpdateInventoryResponse) GetCategory() ListInventoryItemCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *UpdateInventoryResponse) GetCategoryOk() (*ListInventoryItemCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *UpdateInventoryResponse) SetCategory(v ListInventoryItemCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *UpdateInventoryResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *UpdateInventoryResponse) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *UpdateInventoryResponse) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


