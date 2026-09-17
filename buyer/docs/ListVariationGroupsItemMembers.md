# ListVariationGroupsItemMembers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListingId** | **string** |  | 
**AxisValue** | Pointer to **NullableString** |  | [optional] 
**SortKey** | **float32** |  | 
**Title** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableString** |  | [optional] 
**Images** | Pointer to **[]string** |  | [optional] 
**Status** | **string** |  | 
**Sku** | Pointer to **NullableString** |  | [optional] 
**QuantityAvailable** | **float32** |  | 

## Methods

### NewListVariationGroupsItemMembers

`func NewListVariationGroupsItemMembers(listingId string, sortKey float32, status string, quantityAvailable float32, ) *ListVariationGroupsItemMembers`

NewListVariationGroupsItemMembers instantiates a new ListVariationGroupsItemMembers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVariationGroupsItemMembersWithDefaults

`func NewListVariationGroupsItemMembersWithDefaults() *ListVariationGroupsItemMembers`

NewListVariationGroupsItemMembersWithDefaults instantiates a new ListVariationGroupsItemMembers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingId

`func (o *ListVariationGroupsItemMembers) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *ListVariationGroupsItemMembers) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *ListVariationGroupsItemMembers) SetListingId(v string)`

SetListingId sets ListingId field to given value.


### GetAxisValue

`func (o *ListVariationGroupsItemMembers) GetAxisValue() string`

GetAxisValue returns the AxisValue field if non-nil, zero value otherwise.

### GetAxisValueOk

`func (o *ListVariationGroupsItemMembers) GetAxisValueOk() (*string, bool)`

GetAxisValueOk returns a tuple with the AxisValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAxisValue

`func (o *ListVariationGroupsItemMembers) SetAxisValue(v string)`

SetAxisValue sets AxisValue field to given value.

### HasAxisValue

`func (o *ListVariationGroupsItemMembers) HasAxisValue() bool`

HasAxisValue returns a boolean if a field has been set.

### SetAxisValueNil

`func (o *ListVariationGroupsItemMembers) SetAxisValueNil(b bool)`

 SetAxisValueNil sets the value for AxisValue to be an explicit nil

### UnsetAxisValue
`func (o *ListVariationGroupsItemMembers) UnsetAxisValue()`

UnsetAxisValue ensures that no value is present for AxisValue, not even an explicit nil
### GetSortKey

`func (o *ListVariationGroupsItemMembers) GetSortKey() float32`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *ListVariationGroupsItemMembers) GetSortKeyOk() (*float32, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *ListVariationGroupsItemMembers) SetSortKey(v float32)`

SetSortKey sets SortKey field to given value.


### GetTitle

`func (o *ListVariationGroupsItemMembers) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListVariationGroupsItemMembers) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListVariationGroupsItemMembers) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListVariationGroupsItemMembers) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ListVariationGroupsItemMembers) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ListVariationGroupsItemMembers) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetPrice

`func (o *ListVariationGroupsItemMembers) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListVariationGroupsItemMembers) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListVariationGroupsItemMembers) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListVariationGroupsItemMembers) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *ListVariationGroupsItemMembers) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *ListVariationGroupsItemMembers) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetImages

`func (o *ListVariationGroupsItemMembers) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ListVariationGroupsItemMembers) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ListVariationGroupsItemMembers) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *ListVariationGroupsItemMembers) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *ListVariationGroupsItemMembers) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *ListVariationGroupsItemMembers) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetStatus

`func (o *ListVariationGroupsItemMembers) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListVariationGroupsItemMembers) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListVariationGroupsItemMembers) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSku

`func (o *ListVariationGroupsItemMembers) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ListVariationGroupsItemMembers) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ListVariationGroupsItemMembers) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ListVariationGroupsItemMembers) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *ListVariationGroupsItemMembers) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *ListVariationGroupsItemMembers) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetQuantityAvailable

`func (o *ListVariationGroupsItemMembers) GetQuantityAvailable() float32`

GetQuantityAvailable returns the QuantityAvailable field if non-nil, zero value otherwise.

### GetQuantityAvailableOk

`func (o *ListVariationGroupsItemMembers) GetQuantityAvailableOk() (*float32, bool)`

GetQuantityAvailableOk returns a tuple with the QuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAvailable

`func (o *ListVariationGroupsItemMembers) SetQuantityAvailable(v float32)`

SetQuantityAvailable sets QuantityAvailable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


