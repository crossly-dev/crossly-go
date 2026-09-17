# ListInventoryActivityItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Platform** | Pointer to **NullableString** |  | [optional] 
**Type** | **string** |  | 
**OrderId** | Pointer to **NullableString** |  | [optional] 
**InventoryItemId** | Pointer to **NullableString** |  | [optional] 
**Delta** | Pointer to **NullableFloat32** |  | [optional] 
**QuantityBefore** | Pointer to **NullableFloat32** |  | [optional] 
**QuantityAfter** | Pointer to **NullableFloat32** |  | [optional] 
**SaleId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListInventoryActivityItem

`func NewListInventoryActivityItem(id string, createdAt time.Time, userId string, type_ string, ) *ListInventoryActivityItem`

NewListInventoryActivityItem instantiates a new ListInventoryActivityItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInventoryActivityItemWithDefaults

`func NewListInventoryActivityItemWithDefaults() *ListInventoryActivityItem`

NewListInventoryActivityItemWithDefaults instantiates a new ListInventoryActivityItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInventoryActivityItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInventoryActivityItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInventoryActivityItem) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListInventoryActivityItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListInventoryActivityItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListInventoryActivityItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *ListInventoryActivityItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListInventoryActivityItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListInventoryActivityItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetPlatform

`func (o *ListInventoryActivityItem) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListInventoryActivityItem) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListInventoryActivityItem) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListInventoryActivityItem) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *ListInventoryActivityItem) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *ListInventoryActivityItem) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetType

`func (o *ListInventoryActivityItem) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListInventoryActivityItem) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListInventoryActivityItem) SetType(v string)`

SetType sets Type field to given value.


### GetOrderId

`func (o *ListInventoryActivityItem) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ListInventoryActivityItem) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ListInventoryActivityItem) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ListInventoryActivityItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *ListInventoryActivityItem) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *ListInventoryActivityItem) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetInventoryItemId

`func (o *ListInventoryActivityItem) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *ListInventoryActivityItem) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *ListInventoryActivityItem) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *ListInventoryActivityItem) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### SetInventoryItemIdNil

`func (o *ListInventoryActivityItem) SetInventoryItemIdNil(b bool)`

 SetInventoryItemIdNil sets the value for InventoryItemId to be an explicit nil

### UnsetInventoryItemId
`func (o *ListInventoryActivityItem) UnsetInventoryItemId()`

UnsetInventoryItemId ensures that no value is present for InventoryItemId, not even an explicit nil
### GetDelta

`func (o *ListInventoryActivityItem) GetDelta() float32`

GetDelta returns the Delta field if non-nil, zero value otherwise.

### GetDeltaOk

`func (o *ListInventoryActivityItem) GetDeltaOk() (*float32, bool)`

GetDeltaOk returns a tuple with the Delta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelta

`func (o *ListInventoryActivityItem) SetDelta(v float32)`

SetDelta sets Delta field to given value.

### HasDelta

`func (o *ListInventoryActivityItem) HasDelta() bool`

HasDelta returns a boolean if a field has been set.

### SetDeltaNil

`func (o *ListInventoryActivityItem) SetDeltaNil(b bool)`

 SetDeltaNil sets the value for Delta to be an explicit nil

### UnsetDelta
`func (o *ListInventoryActivityItem) UnsetDelta()`

UnsetDelta ensures that no value is present for Delta, not even an explicit nil
### GetQuantityBefore

`func (o *ListInventoryActivityItem) GetQuantityBefore() float32`

GetQuantityBefore returns the QuantityBefore field if non-nil, zero value otherwise.

### GetQuantityBeforeOk

`func (o *ListInventoryActivityItem) GetQuantityBeforeOk() (*float32, bool)`

GetQuantityBeforeOk returns a tuple with the QuantityBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityBefore

`func (o *ListInventoryActivityItem) SetQuantityBefore(v float32)`

SetQuantityBefore sets QuantityBefore field to given value.

### HasQuantityBefore

`func (o *ListInventoryActivityItem) HasQuantityBefore() bool`

HasQuantityBefore returns a boolean if a field has been set.

### SetQuantityBeforeNil

`func (o *ListInventoryActivityItem) SetQuantityBeforeNil(b bool)`

 SetQuantityBeforeNil sets the value for QuantityBefore to be an explicit nil

### UnsetQuantityBefore
`func (o *ListInventoryActivityItem) UnsetQuantityBefore()`

UnsetQuantityBefore ensures that no value is present for QuantityBefore, not even an explicit nil
### GetQuantityAfter

`func (o *ListInventoryActivityItem) GetQuantityAfter() float32`

GetQuantityAfter returns the QuantityAfter field if non-nil, zero value otherwise.

### GetQuantityAfterOk

`func (o *ListInventoryActivityItem) GetQuantityAfterOk() (*float32, bool)`

GetQuantityAfterOk returns a tuple with the QuantityAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAfter

`func (o *ListInventoryActivityItem) SetQuantityAfter(v float32)`

SetQuantityAfter sets QuantityAfter field to given value.

### HasQuantityAfter

`func (o *ListInventoryActivityItem) HasQuantityAfter() bool`

HasQuantityAfter returns a boolean if a field has been set.

### SetQuantityAfterNil

`func (o *ListInventoryActivityItem) SetQuantityAfterNil(b bool)`

 SetQuantityAfterNil sets the value for QuantityAfter to be an explicit nil

### UnsetQuantityAfter
`func (o *ListInventoryActivityItem) UnsetQuantityAfter()`

UnsetQuantityAfter ensures that no value is present for QuantityAfter, not even an explicit nil
### GetSaleId

`func (o *ListInventoryActivityItem) GetSaleId() string`

GetSaleId returns the SaleId field if non-nil, zero value otherwise.

### GetSaleIdOk

`func (o *ListInventoryActivityItem) GetSaleIdOk() (*string, bool)`

GetSaleIdOk returns a tuple with the SaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleId

`func (o *ListInventoryActivityItem) SetSaleId(v string)`

SetSaleId sets SaleId field to given value.

### HasSaleId

`func (o *ListInventoryActivityItem) HasSaleId() bool`

HasSaleId returns a boolean if a field has been set.

### SetSaleIdNil

`func (o *ListInventoryActivityItem) SetSaleIdNil(b bool)`

 SetSaleIdNil sets the value for SaleId to be an explicit nil

### UnsetSaleId
`func (o *ListInventoryActivityItem) UnsetSaleId()`

UnsetSaleId ensures that no value is present for SaleId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


