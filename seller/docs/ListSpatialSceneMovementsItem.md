# ListSpatialSceneMovementsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**At** | **string** |  | 
**Kind** | **string** |  | 
**Source** | **string** |  | 
**Quantity** | **float32** |  | 
**ItemId** | **string** |  | 
**UnitId** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**FromNodeId** | Pointer to **NullableString** |  | [optional] 
**ToNodeId** | Pointer to **NullableString** |  | [optional] 
**FromCode** | Pointer to **NullableString** |  | [optional] 
**ToCode** | Pointer to **NullableString** |  | [optional] 
**ActorId** | Pointer to **NullableString** | NULL when nothing human did it, OR when attribution is redacted for this caller. &#x60;actorsRedacted&#x60; on the payload distinguishes the two. | [optional] 

## Methods

### NewListSpatialSceneMovementsItem

`func NewListSpatialSceneMovementsItem(id string, at string, kind string, source string, quantity float32, itemId string, ) *ListSpatialSceneMovementsItem`

NewListSpatialSceneMovementsItem instantiates a new ListSpatialSceneMovementsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSpatialSceneMovementsItemWithDefaults

`func NewListSpatialSceneMovementsItemWithDefaults() *ListSpatialSceneMovementsItem`

NewListSpatialSceneMovementsItemWithDefaults instantiates a new ListSpatialSceneMovementsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSpatialSceneMovementsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSpatialSceneMovementsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSpatialSceneMovementsItem) SetId(v string)`

SetId sets Id field to given value.


### GetAt

`func (o *ListSpatialSceneMovementsItem) GetAt() string`

GetAt returns the At field if non-nil, zero value otherwise.

### GetAtOk

`func (o *ListSpatialSceneMovementsItem) GetAtOk() (*string, bool)`

GetAtOk returns a tuple with the At field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAt

`func (o *ListSpatialSceneMovementsItem) SetAt(v string)`

SetAt sets At field to given value.


### GetKind

`func (o *ListSpatialSceneMovementsItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListSpatialSceneMovementsItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListSpatialSceneMovementsItem) SetKind(v string)`

SetKind sets Kind field to given value.


### GetSource

`func (o *ListSpatialSceneMovementsItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListSpatialSceneMovementsItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListSpatialSceneMovementsItem) SetSource(v string)`

SetSource sets Source field to given value.


### GetQuantity

`func (o *ListSpatialSceneMovementsItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ListSpatialSceneMovementsItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ListSpatialSceneMovementsItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetItemId

`func (o *ListSpatialSceneMovementsItem) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *ListSpatialSceneMovementsItem) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *ListSpatialSceneMovementsItem) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetUnitId

`func (o *ListSpatialSceneMovementsItem) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *ListSpatialSceneMovementsItem) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *ListSpatialSceneMovementsItem) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *ListSpatialSceneMovementsItem) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *ListSpatialSceneMovementsItem) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *ListSpatialSceneMovementsItem) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetTitle

`func (o *ListSpatialSceneMovementsItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListSpatialSceneMovementsItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListSpatialSceneMovementsItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListSpatialSceneMovementsItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ListSpatialSceneMovementsItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ListSpatialSceneMovementsItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetFromNodeId

`func (o *ListSpatialSceneMovementsItem) GetFromNodeId() string`

GetFromNodeId returns the FromNodeId field if non-nil, zero value otherwise.

### GetFromNodeIdOk

`func (o *ListSpatialSceneMovementsItem) GetFromNodeIdOk() (*string, bool)`

GetFromNodeIdOk returns a tuple with the FromNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromNodeId

`func (o *ListSpatialSceneMovementsItem) SetFromNodeId(v string)`

SetFromNodeId sets FromNodeId field to given value.

### HasFromNodeId

`func (o *ListSpatialSceneMovementsItem) HasFromNodeId() bool`

HasFromNodeId returns a boolean if a field has been set.

### SetFromNodeIdNil

`func (o *ListSpatialSceneMovementsItem) SetFromNodeIdNil(b bool)`

 SetFromNodeIdNil sets the value for FromNodeId to be an explicit nil

### UnsetFromNodeId
`func (o *ListSpatialSceneMovementsItem) UnsetFromNodeId()`

UnsetFromNodeId ensures that no value is present for FromNodeId, not even an explicit nil
### GetToNodeId

`func (o *ListSpatialSceneMovementsItem) GetToNodeId() string`

GetToNodeId returns the ToNodeId field if non-nil, zero value otherwise.

### GetToNodeIdOk

`func (o *ListSpatialSceneMovementsItem) GetToNodeIdOk() (*string, bool)`

GetToNodeIdOk returns a tuple with the ToNodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToNodeId

`func (o *ListSpatialSceneMovementsItem) SetToNodeId(v string)`

SetToNodeId sets ToNodeId field to given value.

### HasToNodeId

`func (o *ListSpatialSceneMovementsItem) HasToNodeId() bool`

HasToNodeId returns a boolean if a field has been set.

### SetToNodeIdNil

`func (o *ListSpatialSceneMovementsItem) SetToNodeIdNil(b bool)`

 SetToNodeIdNil sets the value for ToNodeId to be an explicit nil

### UnsetToNodeId
`func (o *ListSpatialSceneMovementsItem) UnsetToNodeId()`

UnsetToNodeId ensures that no value is present for ToNodeId, not even an explicit nil
### GetFromCode

`func (o *ListSpatialSceneMovementsItem) GetFromCode() string`

GetFromCode returns the FromCode field if non-nil, zero value otherwise.

### GetFromCodeOk

`func (o *ListSpatialSceneMovementsItem) GetFromCodeOk() (*string, bool)`

GetFromCodeOk returns a tuple with the FromCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromCode

`func (o *ListSpatialSceneMovementsItem) SetFromCode(v string)`

SetFromCode sets FromCode field to given value.

### HasFromCode

`func (o *ListSpatialSceneMovementsItem) HasFromCode() bool`

HasFromCode returns a boolean if a field has been set.

### SetFromCodeNil

`func (o *ListSpatialSceneMovementsItem) SetFromCodeNil(b bool)`

 SetFromCodeNil sets the value for FromCode to be an explicit nil

### UnsetFromCode
`func (o *ListSpatialSceneMovementsItem) UnsetFromCode()`

UnsetFromCode ensures that no value is present for FromCode, not even an explicit nil
### GetToCode

`func (o *ListSpatialSceneMovementsItem) GetToCode() string`

GetToCode returns the ToCode field if non-nil, zero value otherwise.

### GetToCodeOk

`func (o *ListSpatialSceneMovementsItem) GetToCodeOk() (*string, bool)`

GetToCodeOk returns a tuple with the ToCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToCode

`func (o *ListSpatialSceneMovementsItem) SetToCode(v string)`

SetToCode sets ToCode field to given value.

### HasToCode

`func (o *ListSpatialSceneMovementsItem) HasToCode() bool`

HasToCode returns a boolean if a field has been set.

### SetToCodeNil

`func (o *ListSpatialSceneMovementsItem) SetToCodeNil(b bool)`

 SetToCodeNil sets the value for ToCode to be an explicit nil

### UnsetToCode
`func (o *ListSpatialSceneMovementsItem) UnsetToCode()`

UnsetToCode ensures that no value is present for ToCode, not even an explicit nil
### GetActorId

`func (o *ListSpatialSceneMovementsItem) GetActorId() string`

GetActorId returns the ActorId field if non-nil, zero value otherwise.

### GetActorIdOk

`func (o *ListSpatialSceneMovementsItem) GetActorIdOk() (*string, bool)`

GetActorIdOk returns a tuple with the ActorId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorId

`func (o *ListSpatialSceneMovementsItem) SetActorId(v string)`

SetActorId sets ActorId field to given value.

### HasActorId

`func (o *ListSpatialSceneMovementsItem) HasActorId() bool`

HasActorId returns a boolean if a field has been set.

### SetActorIdNil

`func (o *ListSpatialSceneMovementsItem) SetActorIdNil(b bool)`

 SetActorIdNil sets the value for ActorId to be an explicit nil

### UnsetActorId
`func (o *ListSpatialSceneMovementsItem) UnsetActorId()`

UnsetActorId ensures that no value is present for ActorId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


