# ListSourcingDemandItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Namespace** | **string** |  | 
**IdentifierValue** | **string** |  | 
**Looks** | **float32** | How many times anybody looked, in the window. | 
**Misses** | **float32** | How many of those we had nothing for. | 
**Retailers** | **float32** | Distinct retailers it was seen on — breadth, not just volume. | 
**MedianPageCents** | Pointer to **NullableFloat32** | What the retailers were charging, median of what we saw. | [optional] 
**LastSeenAt** | **time.Time** |  | 

## Methods

### NewListSourcingDemandItem

`func NewListSourcingDemandItem(namespace string, identifierValue string, looks float32, misses float32, retailers float32, lastSeenAt time.Time, ) *ListSourcingDemandItem`

NewListSourcingDemandItem instantiates a new ListSourcingDemandItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSourcingDemandItemWithDefaults

`func NewListSourcingDemandItemWithDefaults() *ListSourcingDemandItem`

NewListSourcingDemandItemWithDefaults instantiates a new ListSourcingDemandItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNamespace

`func (o *ListSourcingDemandItem) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListSourcingDemandItem) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListSourcingDemandItem) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetIdentifierValue

`func (o *ListSourcingDemandItem) GetIdentifierValue() string`

GetIdentifierValue returns the IdentifierValue field if non-nil, zero value otherwise.

### GetIdentifierValueOk

`func (o *ListSourcingDemandItem) GetIdentifierValueOk() (*string, bool)`

GetIdentifierValueOk returns a tuple with the IdentifierValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierValue

`func (o *ListSourcingDemandItem) SetIdentifierValue(v string)`

SetIdentifierValue sets IdentifierValue field to given value.


### GetLooks

`func (o *ListSourcingDemandItem) GetLooks() float32`

GetLooks returns the Looks field if non-nil, zero value otherwise.

### GetLooksOk

`func (o *ListSourcingDemandItem) GetLooksOk() (*float32, bool)`

GetLooksOk returns a tuple with the Looks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLooks

`func (o *ListSourcingDemandItem) SetLooks(v float32)`

SetLooks sets Looks field to given value.


### GetMisses

`func (o *ListSourcingDemandItem) GetMisses() float32`

GetMisses returns the Misses field if non-nil, zero value otherwise.

### GetMissesOk

`func (o *ListSourcingDemandItem) GetMissesOk() (*float32, bool)`

GetMissesOk returns a tuple with the Misses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMisses

`func (o *ListSourcingDemandItem) SetMisses(v float32)`

SetMisses sets Misses field to given value.


### GetRetailers

`func (o *ListSourcingDemandItem) GetRetailers() float32`

GetRetailers returns the Retailers field if non-nil, zero value otherwise.

### GetRetailersOk

`func (o *ListSourcingDemandItem) GetRetailersOk() (*float32, bool)`

GetRetailersOk returns a tuple with the Retailers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailers

`func (o *ListSourcingDemandItem) SetRetailers(v float32)`

SetRetailers sets Retailers field to given value.


### GetMedianPageCents

`func (o *ListSourcingDemandItem) GetMedianPageCents() float32`

GetMedianPageCents returns the MedianPageCents field if non-nil, zero value otherwise.

### GetMedianPageCentsOk

`func (o *ListSourcingDemandItem) GetMedianPageCentsOk() (*float32, bool)`

GetMedianPageCentsOk returns a tuple with the MedianPageCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMedianPageCents

`func (o *ListSourcingDemandItem) SetMedianPageCents(v float32)`

SetMedianPageCents sets MedianPageCents field to given value.

### HasMedianPageCents

`func (o *ListSourcingDemandItem) HasMedianPageCents() bool`

HasMedianPageCents returns a boolean if a field has been set.

### SetMedianPageCentsNil

`func (o *ListSourcingDemandItem) SetMedianPageCentsNil(b bool)`

 SetMedianPageCentsNil sets the value for MedianPageCents to be an explicit nil

### UnsetMedianPageCents
`func (o *ListSourcingDemandItem) UnsetMedianPageCents()`

UnsetMedianPageCents ensures that no value is present for MedianPageCents, not even an explicit nil
### GetLastSeenAt

`func (o *ListSourcingDemandItem) GetLastSeenAt() time.Time`

GetLastSeenAt returns the LastSeenAt field if non-nil, zero value otherwise.

### GetLastSeenAtOk

`func (o *ListSourcingDemandItem) GetLastSeenAtOk() (*time.Time, bool)`

GetLastSeenAtOk returns a tuple with the LastSeenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenAt

`func (o *ListSourcingDemandItem) SetLastSeenAt(v time.Time)`

SetLastSeenAt sets LastSeenAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


