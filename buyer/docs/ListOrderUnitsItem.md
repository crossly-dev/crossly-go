# ListOrderUnitsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UnitId** | **string** |  | 
**Identifiers** | [**[]ListOrderUnitsItemIdentifiers**](ListOrderUnitsItemIdentifiers.md) |  | 

## Methods

### NewListOrderUnitsItem

`func NewListOrderUnitsItem(unitId string, identifiers []ListOrderUnitsItemIdentifiers, ) *ListOrderUnitsItem`

NewListOrderUnitsItem instantiates a new ListOrderUnitsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrderUnitsItemWithDefaults

`func NewListOrderUnitsItemWithDefaults() *ListOrderUnitsItem`

NewListOrderUnitsItemWithDefaults instantiates a new ListOrderUnitsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUnitId

`func (o *ListOrderUnitsItem) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *ListOrderUnitsItem) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *ListOrderUnitsItem) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.


### GetIdentifiers

`func (o *ListOrderUnitsItem) GetIdentifiers() []ListOrderUnitsItemIdentifiers`

GetIdentifiers returns the Identifiers field if non-nil, zero value otherwise.

### GetIdentifiersOk

`func (o *ListOrderUnitsItem) GetIdentifiersOk() (*[]ListOrderUnitsItemIdentifiers, bool)`

GetIdentifiersOk returns a tuple with the Identifiers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifiers

`func (o *ListOrderUnitsItem) SetIdentifiers(v []ListOrderUnitsItemIdentifiers)`

SetIdentifiers sets Identifiers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


