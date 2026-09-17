# ListBuyerMonitorMatchesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListingSlug** | **string** |  | 
**TriggerPriceCents** | **float32** |  | 
**DeliveredAt** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **string** |  | 

## Methods

### NewListBuyerMonitorMatchesItem

`func NewListBuyerMonitorMatchesItem(listingSlug string, triggerPriceCents float32, createdAt string, ) *ListBuyerMonitorMatchesItem`

NewListBuyerMonitorMatchesItem instantiates a new ListBuyerMonitorMatchesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBuyerMonitorMatchesItemWithDefaults

`func NewListBuyerMonitorMatchesItemWithDefaults() *ListBuyerMonitorMatchesItem`

NewListBuyerMonitorMatchesItemWithDefaults instantiates a new ListBuyerMonitorMatchesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingSlug

`func (o *ListBuyerMonitorMatchesItem) GetListingSlug() string`

GetListingSlug returns the ListingSlug field if non-nil, zero value otherwise.

### GetListingSlugOk

`func (o *ListBuyerMonitorMatchesItem) GetListingSlugOk() (*string, bool)`

GetListingSlugOk returns a tuple with the ListingSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingSlug

`func (o *ListBuyerMonitorMatchesItem) SetListingSlug(v string)`

SetListingSlug sets ListingSlug field to given value.


### GetTriggerPriceCents

`func (o *ListBuyerMonitorMatchesItem) GetTriggerPriceCents() float32`

GetTriggerPriceCents returns the TriggerPriceCents field if non-nil, zero value otherwise.

### GetTriggerPriceCentsOk

`func (o *ListBuyerMonitorMatchesItem) GetTriggerPriceCentsOk() (*float32, bool)`

GetTriggerPriceCentsOk returns a tuple with the TriggerPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggerPriceCents

`func (o *ListBuyerMonitorMatchesItem) SetTriggerPriceCents(v float32)`

SetTriggerPriceCents sets TriggerPriceCents field to given value.


### GetDeliveredAt

`func (o *ListBuyerMonitorMatchesItem) GetDeliveredAt() string`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *ListBuyerMonitorMatchesItem) GetDeliveredAtOk() (*string, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *ListBuyerMonitorMatchesItem) SetDeliveredAt(v string)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *ListBuyerMonitorMatchesItem) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### SetDeliveredAtNil

`func (o *ListBuyerMonitorMatchesItem) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *ListBuyerMonitorMatchesItem) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil
### GetCreatedAt

`func (o *ListBuyerMonitorMatchesItem) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListBuyerMonitorMatchesItem) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListBuyerMonitorMatchesItem) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


