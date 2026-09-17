# GetInventoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformListings** | [**[]GetInventoryResponsePlatformListings**](GetInventoryResponsePlatformListings.md) |  | 
**Category** | Pointer to [**NullableListInventoryItemCategory**](ListInventoryItemCategory.md) |  | [optional] 

## Methods

### NewGetInventoryResponse

`func NewGetInventoryResponse(platformListings []GetInventoryResponsePlatformListings, ) *GetInventoryResponse`

NewGetInventoryResponse instantiates a new GetInventoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInventoryResponseWithDefaults

`func NewGetInventoryResponseWithDefaults() *GetInventoryResponse`

NewGetInventoryResponseWithDefaults instantiates a new GetInventoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformListings

`func (o *GetInventoryResponse) GetPlatformListings() []GetInventoryResponsePlatformListings`

GetPlatformListings returns the PlatformListings field if non-nil, zero value otherwise.

### GetPlatformListingsOk

`func (o *GetInventoryResponse) GetPlatformListingsOk() (*[]GetInventoryResponsePlatformListings, bool)`

GetPlatformListingsOk returns a tuple with the PlatformListings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformListings

`func (o *GetInventoryResponse) SetPlatformListings(v []GetInventoryResponsePlatformListings)`

SetPlatformListings sets PlatformListings field to given value.


### GetCategory

`func (o *GetInventoryResponse) GetCategory() ListInventoryItemCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetInventoryResponse) GetCategoryOk() (*ListInventoryItemCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetInventoryResponse) SetCategory(v ListInventoryItemCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetInventoryResponse) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetInventoryResponse) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetInventoryResponse) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


