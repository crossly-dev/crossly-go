# ListSourcingDemandMineItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentifierValue** | **string** |  | 
**Lookers** | **float32** | How many distinct shoppers looked, in the window. | 
**Misses** | **float32** | How many of those looks Crossly could not answer at all. | 
**MedianRetailCents** | Pointer to **NullableFloat32** | What the retailers were charging, median of what Scout saw. | [optional] 
**Relation** | **string** | &#39;in_stock&#39; — it is in their inventory. &#39;sold_before&#39; — they have sold one. | 
**InventoryItemId** | Pointer to **NullableString** | Their own row, for the link. | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**LastSoldCents** | Pointer to **NullableFloat32** | What they got for it last time, when they have sold one. | [optional] 
**LastSoldAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewListSourcingDemandMineItem

`func NewListSourcingDemandMineItem(identifierValue string, lookers float32, misses float32, relation string, ) *ListSourcingDemandMineItem`

NewListSourcingDemandMineItem instantiates a new ListSourcingDemandMineItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSourcingDemandMineItemWithDefaults

`func NewListSourcingDemandMineItemWithDefaults() *ListSourcingDemandMineItem`

NewListSourcingDemandMineItemWithDefaults instantiates a new ListSourcingDemandMineItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifierValue

`func (o *ListSourcingDemandMineItem) GetIdentifierValue() string`

GetIdentifierValue returns the IdentifierValue field if non-nil, zero value otherwise.

### GetIdentifierValueOk

`func (o *ListSourcingDemandMineItem) GetIdentifierValueOk() (*string, bool)`

GetIdentifierValueOk returns a tuple with the IdentifierValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierValue

`func (o *ListSourcingDemandMineItem) SetIdentifierValue(v string)`

SetIdentifierValue sets IdentifierValue field to given value.


### GetLookers

`func (o *ListSourcingDemandMineItem) GetLookers() float32`

GetLookers returns the Lookers field if non-nil, zero value otherwise.

### GetLookersOk

`func (o *ListSourcingDemandMineItem) GetLookersOk() (*float32, bool)`

GetLookersOk returns a tuple with the Lookers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLookers

`func (o *ListSourcingDemandMineItem) SetLookers(v float32)`

SetLookers sets Lookers field to given value.


### GetMisses

`func (o *ListSourcingDemandMineItem) GetMisses() float32`

GetMisses returns the Misses field if non-nil, zero value otherwise.

### GetMissesOk

`func (o *ListSourcingDemandMineItem) GetMissesOk() (*float32, bool)`

GetMissesOk returns a tuple with the Misses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMisses

`func (o *ListSourcingDemandMineItem) SetMisses(v float32)`

SetMisses sets Misses field to given value.


### GetMedianRetailCents

`func (o *ListSourcingDemandMineItem) GetMedianRetailCents() float32`

GetMedianRetailCents returns the MedianRetailCents field if non-nil, zero value otherwise.

### GetMedianRetailCentsOk

`func (o *ListSourcingDemandMineItem) GetMedianRetailCentsOk() (*float32, bool)`

GetMedianRetailCentsOk returns a tuple with the MedianRetailCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianRetailCents

`func (o *ListSourcingDemandMineItem) SetMedianRetailCents(v float32)`

SetMedianRetailCents sets MedianRetailCents field to given value.

### HasMedianRetailCents

`func (o *ListSourcingDemandMineItem) HasMedianRetailCents() bool`

HasMedianRetailCents returns a boolean if a field has been set.

### SetMedianRetailCentsNil

`func (o *ListSourcingDemandMineItem) SetMedianRetailCentsNil(b bool)`

 SetMedianRetailCentsNil sets the value for MedianRetailCents to be an explicit nil

### UnsetMedianRetailCents
`func (o *ListSourcingDemandMineItem) UnsetMedianRetailCents()`

UnsetMedianRetailCents ensures that no value is present for MedianRetailCents, not even an explicit nil
### GetRelation

`func (o *ListSourcingDemandMineItem) GetRelation() string`

GetRelation returns the Relation field if non-nil, zero value otherwise.

### GetRelationOk

`func (o *ListSourcingDemandMineItem) GetRelationOk() (*string, bool)`

GetRelationOk returns a tuple with the Relation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelation

`func (o *ListSourcingDemandMineItem) SetRelation(v string)`

SetRelation sets Relation field to given value.


### GetInventoryItemId

`func (o *ListSourcingDemandMineItem) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *ListSourcingDemandMineItem) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *ListSourcingDemandMineItem) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *ListSourcingDemandMineItem) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### SetInventoryItemIdNil

`func (o *ListSourcingDemandMineItem) SetInventoryItemIdNil(b bool)`

 SetInventoryItemIdNil sets the value for InventoryItemId to be an explicit nil

### UnsetInventoryItemId
`func (o *ListSourcingDemandMineItem) UnsetInventoryItemId()`

UnsetInventoryItemId ensures that no value is present for InventoryItemId, not even an explicit nil
### GetTitle

`func (o *ListSourcingDemandMineItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListSourcingDemandMineItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListSourcingDemandMineItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListSourcingDemandMineItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ListSourcingDemandMineItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ListSourcingDemandMineItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetLastSoldCents

`func (o *ListSourcingDemandMineItem) GetLastSoldCents() float32`

GetLastSoldCents returns the LastSoldCents field if non-nil, zero value otherwise.

### GetLastSoldCentsOk

`func (o *ListSourcingDemandMineItem) GetLastSoldCentsOk() (*float32, bool)`

GetLastSoldCentsOk returns a tuple with the LastSoldCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSoldCents

`func (o *ListSourcingDemandMineItem) SetLastSoldCents(v float32)`

SetLastSoldCents sets LastSoldCents field to given value.

### HasLastSoldCents

`func (o *ListSourcingDemandMineItem) HasLastSoldCents() bool`

HasLastSoldCents returns a boolean if a field has been set.

### SetLastSoldCentsNil

`func (o *ListSourcingDemandMineItem) SetLastSoldCentsNil(b bool)`

 SetLastSoldCentsNil sets the value for LastSoldCents to be an explicit nil

### UnsetLastSoldCents
`func (o *ListSourcingDemandMineItem) UnsetLastSoldCents()`

UnsetLastSoldCents ensures that no value is present for LastSoldCents, not even an explicit nil
### GetLastSoldAt

`func (o *ListSourcingDemandMineItem) GetLastSoldAt() time.Time`

GetLastSoldAt returns the LastSoldAt field if non-nil, zero value otherwise.

### GetLastSoldAtOk

`func (o *ListSourcingDemandMineItem) GetLastSoldAtOk() (*time.Time, bool)`

GetLastSoldAtOk returns a tuple with the LastSoldAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSoldAt

`func (o *ListSourcingDemandMineItem) SetLastSoldAt(v time.Time)`

SetLastSoldAt sets LastSoldAt field to given value.

### HasLastSoldAt

`func (o *ListSourcingDemandMineItem) HasLastSoldAt() bool`

HasLastSoldAt returns a boolean if a field has been set.

### SetLastSoldAtNil

`func (o *ListSourcingDemandMineItem) SetLastSoldAtNil(b bool)`

 SetLastSoldAtNil sets the value for LastSoldAt to be an explicit nil

### UnsetLastSoldAt
`func (o *ListSourcingDemandMineItem) UnsetLastSoldAt()`

UnsetLastSoldAt ensures that no value is present for LastSoldAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


