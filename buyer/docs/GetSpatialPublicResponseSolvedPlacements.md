# GetSpatialPublicResponseSolvedPlacements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ItemId** | **string** |  | 
**UnitId** | Pointer to **NullableString** |  | [optional] 
**ContainerKey** | **string** |  | 
**ContainerIndex** | **float32** |  | 
**SlotIndex** | **float32** |  | 
**Position** | **map[string]interface{}** | Local position inside the container, metres. | 
**GroupPath** | **[]string** |  | 
**Pinned** | **bool** |  | 

## Methods

### NewGetSpatialPublicResponseSolvedPlacements

`func NewGetSpatialPublicResponseSolvedPlacements(itemId string, containerKey string, containerIndex float32, slotIndex float32, position map[string]interface{}, groupPath []string, pinned bool, ) *GetSpatialPublicResponseSolvedPlacements`

NewGetSpatialPublicResponseSolvedPlacements instantiates a new GetSpatialPublicResponseSolvedPlacements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSpatialPublicResponseSolvedPlacementsWithDefaults

`func NewGetSpatialPublicResponseSolvedPlacementsWithDefaults() *GetSpatialPublicResponseSolvedPlacements`

NewGetSpatialPublicResponseSolvedPlacementsWithDefaults instantiates a new GetSpatialPublicResponseSolvedPlacements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItemId

`func (o *GetSpatialPublicResponseSolvedPlacements) GetItemId() string`

GetItemId returns the ItemId field if non-nil, zero value otherwise.

### GetItemIdOk

`func (o *GetSpatialPublicResponseSolvedPlacements) GetItemIdOk() (*string, bool)`

GetItemIdOk returns a tuple with the ItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemId

`func (o *GetSpatialPublicResponseSolvedPlacements) SetItemId(v string)`

SetItemId sets ItemId field to given value.


### GetUnitId

`func (o *GetSpatialPublicResponseSolvedPlacements) GetUnitId() string`

GetUnitId returns the UnitId field if non-nil, zero value otherwise.

### GetUnitIdOk

`func (o *GetSpatialPublicResponseSolvedPlacements) GetUnitIdOk() (*string, bool)`

GetUnitIdOk returns a tuple with the UnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitId

`func (o *GetSpatialPublicResponseSolvedPlacements) SetUnitId(v string)`

SetUnitId sets UnitId field to given value.

### HasUnitId

`func (o *GetSpatialPublicResponseSolvedPlacements) HasUnitId() bool`

HasUnitId returns a boolean if a field has been set.

### SetUnitIdNil

`func (o *GetSpatialPublicResponseSolvedPlacements) SetUnitIdNil(b bool)`

 SetUnitIdNil sets the value for UnitId to be an explicit nil

### UnsetUnitId
`func (o *GetSpatialPublicResponseSolvedPlacements) UnsetUnitId()`

UnsetUnitId ensures that no value is present for UnitId, not even an explicit nil
### GetContainerKey

`func (o *GetSpatialPublicResponseSolvedPlacements) GetContainerKey() string`

GetContainerKey returns the ContainerKey field if non-nil, zero value otherwise.

### GetContainerKeyOk

`func (o *GetSpatialPublicResponseSolvedPlacements) GetContainerKeyOk() (*string, bool)`

GetContainerKeyOk returns a tuple with the ContainerKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerKey

`func (o *GetSpatialPublicResponseSolvedPlacements) SetContainerKey(v string)`

SetContainerKey sets ContainerKey field to given value.


### GetContainerIndex

`func (o *GetSpatialPublicResponseSolvedPlacements) GetContainerIndex() float32`

GetContainerIndex returns the ContainerIndex field if non-nil, zero value otherwise.

### GetContainerIndexOk

`func (o *GetSpatialPublicResponseSolvedPlacements) GetContainerIndexOk() (*float32, bool)`

GetContainerIndexOk returns a tuple with the ContainerIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerIndex

`func (o *GetSpatialPublicResponseSolvedPlacements) SetContainerIndex(v float32)`

SetContainerIndex sets ContainerIndex field to given value.


### GetSlotIndex

`func (o *GetSpatialPublicResponseSolvedPlacements) GetSlotIndex() float32`

GetSlotIndex returns the SlotIndex field if non-nil, zero value otherwise.

### GetSlotIndexOk

`func (o *GetSpatialPublicResponseSolvedPlacements) GetSlotIndexOk() (*float32, bool)`

GetSlotIndexOk returns a tuple with the SlotIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlotIndex

`func (o *GetSpatialPublicResponseSolvedPlacements) SetSlotIndex(v float32)`

SetSlotIndex sets SlotIndex field to given value.


### GetPosition

`func (o *GetSpatialPublicResponseSolvedPlacements) GetPosition() map[string]interface{}`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *GetSpatialPublicResponseSolvedPlacements) GetPositionOk() (*map[string]interface{}, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *GetSpatialPublicResponseSolvedPlacements) SetPosition(v map[string]interface{})`

SetPosition sets Position field to given value.


### GetGroupPath

`func (o *GetSpatialPublicResponseSolvedPlacements) GetGroupPath() []string`

GetGroupPath returns the GroupPath field if non-nil, zero value otherwise.

### GetGroupPathOk

`func (o *GetSpatialPublicResponseSolvedPlacements) GetGroupPathOk() (*[]string, bool)`

GetGroupPathOk returns a tuple with the GroupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPath

`func (o *GetSpatialPublicResponseSolvedPlacements) SetGroupPath(v []string)`

SetGroupPath sets GroupPath field to given value.


### GetPinned

`func (o *GetSpatialPublicResponseSolvedPlacements) GetPinned() bool`

GetPinned returns the Pinned field if non-nil, zero value otherwise.

### GetPinnedOk

`func (o *GetSpatialPublicResponseSolvedPlacements) GetPinnedOk() (*bool, bool)`

GetPinnedOk returns a tuple with the Pinned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPinned

`func (o *GetSpatialPublicResponseSolvedPlacements) SetPinned(v bool)`

SetPinned sets Pinned field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


