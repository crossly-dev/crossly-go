# ListRestockPromptsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ListingId** | **string** |  | 
**InventoryItemId** | Pointer to **NullableString** |  | [optional] 
**Platforms** | **[]string** |  | 
**NewQuantity** | **float32** |  | 
**CreatedAt** | **time.Time** |  | 
**ListingTitle** | Pointer to **NullableString** |  | [optional] 
**ListingImages** | Pointer to **[]string** |  | [optional] 
**ListingPrice** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListRestockPromptsItem

`func NewListRestockPromptsItem(id string, listingId string, platforms []string, newQuantity float32, createdAt time.Time, ) *ListRestockPromptsItem`

NewListRestockPromptsItem instantiates a new ListRestockPromptsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListRestockPromptsItemWithDefaults

`func NewListRestockPromptsItemWithDefaults() *ListRestockPromptsItem`

NewListRestockPromptsItemWithDefaults instantiates a new ListRestockPromptsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListRestockPromptsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListRestockPromptsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListRestockPromptsItem) SetId(v string)`

SetId sets Id field to given value.


### GetListingId

`func (o *ListRestockPromptsItem) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *ListRestockPromptsItem) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *ListRestockPromptsItem) SetListingId(v string)`

SetListingId sets ListingId field to given value.


### GetInventoryItemId

`func (o *ListRestockPromptsItem) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *ListRestockPromptsItem) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *ListRestockPromptsItem) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *ListRestockPromptsItem) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### SetInventoryItemIdNil

`func (o *ListRestockPromptsItem) SetInventoryItemIdNil(b bool)`

 SetInventoryItemIdNil sets the value for InventoryItemId to be an explicit nil

### UnsetInventoryItemId
`func (o *ListRestockPromptsItem) UnsetInventoryItemId()`

UnsetInventoryItemId ensures that no value is present for InventoryItemId, not even an explicit nil
### GetPlatforms

`func (o *ListRestockPromptsItem) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *ListRestockPromptsItem) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *ListRestockPromptsItem) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.


### GetNewQuantity

`func (o *ListRestockPromptsItem) GetNewQuantity() float32`

GetNewQuantity returns the NewQuantity field if non-nil, zero value otherwise.

### GetNewQuantityOk

`func (o *ListRestockPromptsItem) GetNewQuantityOk() (*float32, bool)`

GetNewQuantityOk returns a tuple with the NewQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewQuantity

`func (o *ListRestockPromptsItem) SetNewQuantity(v float32)`

SetNewQuantity sets NewQuantity field to given value.


### GetCreatedAt

`func (o *ListRestockPromptsItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListRestockPromptsItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListRestockPromptsItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetListingTitle

`func (o *ListRestockPromptsItem) GetListingTitle() string`

GetListingTitle returns the ListingTitle field if non-nil, zero value otherwise.

### GetListingTitleOk

`func (o *ListRestockPromptsItem) GetListingTitleOk() (*string, bool)`

GetListingTitleOk returns a tuple with the ListingTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingTitle

`func (o *ListRestockPromptsItem) SetListingTitle(v string)`

SetListingTitle sets ListingTitle field to given value.

### HasListingTitle

`func (o *ListRestockPromptsItem) HasListingTitle() bool`

HasListingTitle returns a boolean if a field has been set.

### SetListingTitleNil

`func (o *ListRestockPromptsItem) SetListingTitleNil(b bool)`

 SetListingTitleNil sets the value for ListingTitle to be an explicit nil

### UnsetListingTitle
`func (o *ListRestockPromptsItem) UnsetListingTitle()`

UnsetListingTitle ensures that no value is present for ListingTitle, not even an explicit nil
### GetListingImages

`func (o *ListRestockPromptsItem) GetListingImages() []string`

GetListingImages returns the ListingImages field if non-nil, zero value otherwise.

### GetListingImagesOk

`func (o *ListRestockPromptsItem) GetListingImagesOk() (*[]string, bool)`

GetListingImagesOk returns a tuple with the ListingImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingImages

`func (o *ListRestockPromptsItem) SetListingImages(v []string)`

SetListingImages sets ListingImages field to given value.

### HasListingImages

`func (o *ListRestockPromptsItem) HasListingImages() bool`

HasListingImages returns a boolean if a field has been set.

### SetListingImagesNil

`func (o *ListRestockPromptsItem) SetListingImagesNil(b bool)`

 SetListingImagesNil sets the value for ListingImages to be an explicit nil

### UnsetListingImages
`func (o *ListRestockPromptsItem) UnsetListingImages()`

UnsetListingImages ensures that no value is present for ListingImages, not even an explicit nil
### GetListingPrice

`func (o *ListRestockPromptsItem) GetListingPrice() string`

GetListingPrice returns the ListingPrice field if non-nil, zero value otherwise.

### GetListingPriceOk

`func (o *ListRestockPromptsItem) GetListingPriceOk() (*string, bool)`

GetListingPriceOk returns a tuple with the ListingPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingPrice

`func (o *ListRestockPromptsItem) SetListingPrice(v string)`

SetListingPrice sets ListingPrice field to given value.

### HasListingPrice

`func (o *ListRestockPromptsItem) HasListingPrice() bool`

HasListingPrice returns a boolean if a field has been set.

### SetListingPriceNil

`func (o *ListRestockPromptsItem) SetListingPriceNil(b bool)`

 SetListingPriceNil sets the value for ListingPrice to be an explicit nil

### UnsetListingPrice
`func (o *ListRestockPromptsItem) UnsetListingPrice()`

UnsetListingPrice ensures that no value is present for ListingPrice, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


