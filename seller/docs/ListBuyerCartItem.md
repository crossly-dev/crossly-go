# ListBuyerCartItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | **bool** |  | 
**Id** | **string** |  | 
**PlatformListingId** | **string** |  | 
**Quantity** | **float32** |  | 
**UnitPriceCents** | **float32** |  | 
**Slug** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**ListingStatus** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListBuyerCartItem

`func NewListBuyerCartItem(available bool, id string, platformListingId string, quantity float32, unitPriceCents float32, ) *ListBuyerCartItem`

NewListBuyerCartItem instantiates a new ListBuyerCartItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBuyerCartItemWithDefaults

`func NewListBuyerCartItemWithDefaults() *ListBuyerCartItem`

NewListBuyerCartItemWithDefaults instantiates a new ListBuyerCartItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *ListBuyerCartItem) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *ListBuyerCartItem) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *ListBuyerCartItem) SetAvailable(v bool)`

SetAvailable sets Available field to given value.


### GetId

`func (o *ListBuyerCartItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBuyerCartItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBuyerCartItem) SetId(v string)`

SetId sets Id field to given value.


### GetPlatformListingId

`func (o *ListBuyerCartItem) GetPlatformListingId() string`

GetPlatformListingId returns the PlatformListingId field if non-nil, zero value otherwise.

### GetPlatformListingIdOk

`func (o *ListBuyerCartItem) GetPlatformListingIdOk() (*string, bool)`

GetPlatformListingIdOk returns a tuple with the PlatformListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformListingId

`func (o *ListBuyerCartItem) SetPlatformListingId(v string)`

SetPlatformListingId sets PlatformListingId field to given value.


### GetQuantity

`func (o *ListBuyerCartItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ListBuyerCartItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ListBuyerCartItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetUnitPriceCents

`func (o *ListBuyerCartItem) GetUnitPriceCents() float32`

GetUnitPriceCents returns the UnitPriceCents field if non-nil, zero value otherwise.

### GetUnitPriceCentsOk

`func (o *ListBuyerCartItem) GetUnitPriceCentsOk() (*float32, bool)`

GetUnitPriceCentsOk returns a tuple with the UnitPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitPriceCents

`func (o *ListBuyerCartItem) SetUnitPriceCents(v float32)`

SetUnitPriceCents sets UnitPriceCents field to given value.


### GetSlug

`func (o *ListBuyerCartItem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ListBuyerCartItem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ListBuyerCartItem) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ListBuyerCartItem) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *ListBuyerCartItem) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *ListBuyerCartItem) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetTitle

`func (o *ListBuyerCartItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListBuyerCartItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListBuyerCartItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListBuyerCartItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ListBuyerCartItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ListBuyerCartItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetListingStatus

`func (o *ListBuyerCartItem) GetListingStatus() string`

GetListingStatus returns the ListingStatus field if non-nil, zero value otherwise.

### GetListingStatusOk

`func (o *ListBuyerCartItem) GetListingStatusOk() (*string, bool)`

GetListingStatusOk returns a tuple with the ListingStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingStatus

`func (o *ListBuyerCartItem) SetListingStatus(v string)`

SetListingStatus sets ListingStatus field to given value.

### HasListingStatus

`func (o *ListBuyerCartItem) HasListingStatus() bool`

HasListingStatus returns a boolean if a field has been set.

### SetListingStatusNil

`func (o *ListBuyerCartItem) SetListingStatusNil(b bool)`

 SetListingStatusNil sets the value for ListingStatus to be an explicit nil

### UnsetListingStatus
`func (o *ListBuyerCartItem) UnsetListingStatus()`

UnsetListingStatus ensures that no value is present for ListingStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


