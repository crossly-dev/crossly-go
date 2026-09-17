# ListInventoryUnitsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Status** | **string** |  | 
**OrderId** | Pointer to **NullableString** |  | [optional] 
**SoldAt** | Pointer to **NullableTime** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**Identifiers** | [**[]ListInventoryUnitsItemIdentifiers**](ListInventoryUnitsItemIdentifiers.md) |  | 

## Methods

### NewListInventoryUnitsItem

`func NewListInventoryUnitsItem(id string, status string, createdAt time.Time, identifiers []ListInventoryUnitsItemIdentifiers, ) *ListInventoryUnitsItem`

NewListInventoryUnitsItem instantiates a new ListInventoryUnitsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInventoryUnitsItemWithDefaults

`func NewListInventoryUnitsItemWithDefaults() *ListInventoryUnitsItem`

NewListInventoryUnitsItemWithDefaults instantiates a new ListInventoryUnitsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInventoryUnitsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInventoryUnitsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInventoryUnitsItem) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *ListInventoryUnitsItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListInventoryUnitsItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListInventoryUnitsItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetOrderId

`func (o *ListInventoryUnitsItem) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ListInventoryUnitsItem) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ListInventoryUnitsItem) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ListInventoryUnitsItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *ListInventoryUnitsItem) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *ListInventoryUnitsItem) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetSoldAt

`func (o *ListInventoryUnitsItem) GetSoldAt() time.Time`

GetSoldAt returns the SoldAt field if non-nil, zero value otherwise.

### GetSoldAtOk

`func (o *ListInventoryUnitsItem) GetSoldAtOk() (*time.Time, bool)`

GetSoldAtOk returns a tuple with the SoldAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldAt

`func (o *ListInventoryUnitsItem) SetSoldAt(v time.Time)`

SetSoldAt sets SoldAt field to given value.

### HasSoldAt

`func (o *ListInventoryUnitsItem) HasSoldAt() bool`

HasSoldAt returns a boolean if a field has been set.

### SetSoldAtNil

`func (o *ListInventoryUnitsItem) SetSoldAtNil(b bool)`

 SetSoldAtNil sets the value for SoldAt to be an explicit nil

### UnsetSoldAt
`func (o *ListInventoryUnitsItem) UnsetSoldAt()`

UnsetSoldAt ensures that no value is present for SoldAt, not even an explicit nil
### GetNotes

`func (o *ListInventoryUnitsItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ListInventoryUnitsItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ListInventoryUnitsItem) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ListInventoryUnitsItem) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ListInventoryUnitsItem) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ListInventoryUnitsItem) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetCreatedAt

`func (o *ListInventoryUnitsItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListInventoryUnitsItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListInventoryUnitsItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetIdentifiers

`func (o *ListInventoryUnitsItem) GetIdentifiers() []ListInventoryUnitsItemIdentifiers`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *ListInventoryUnitsItem) GetIdentifiersOk() (*[]ListInventoryUnitsItemIdentifiers, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *ListInventoryUnitsItem) SetIdentifiers(v []ListInventoryUnitsItemIdentifiers)`

SetIdentifiers sets Identifiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


