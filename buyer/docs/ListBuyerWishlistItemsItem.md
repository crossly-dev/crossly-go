# ListBuyerWishlistItemsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**PlatformListingId** | **string** |  | 
**AddedAt** | **time.Time** |  | 
**Slug** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListBuyerWishlistItemsItem

`func NewListBuyerWishlistItemsItem(id string, platformListingId string, addedAt time.Time, ) *ListBuyerWishlistItemsItem`

NewListBuyerWishlistItemsItem instantiates a new ListBuyerWishlistItemsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBuyerWishlistItemsItemWithDefaults

`func NewListBuyerWishlistItemsItemWithDefaults() *ListBuyerWishlistItemsItem`

NewListBuyerWishlistItemsItemWithDefaults instantiates a new ListBuyerWishlistItemsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListBuyerWishlistItemsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBuyerWishlistItemsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBuyerWishlistItemsItem) SetId(v string)`

SetId sets Id field to given value.


### GetPlatformListingId

`func (o *ListBuyerWishlistItemsItem) GetPlatformListingId() string`

GetPlatformListingId returns the PlatformListingId field if non-nil, zero value otherwise.

### GetPlatformListingIdOk

`func (o *ListBuyerWishlistItemsItem) GetPlatformListingIdOk() (*string, bool)`

GetPlatformListingIdOk returns a tuple with the PlatformListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformListingId

`func (o *ListBuyerWishlistItemsItem) SetPlatformListingId(v string)`

SetPlatformListingId sets PlatformListingId field to given value.


### GetAddedAt

`func (o *ListBuyerWishlistItemsItem) GetAddedAt() time.Time`

GetAddedAt returns the AddedAt field if non-nil, zero value otherwise.

### GetAddedAtOk

`func (o *ListBuyerWishlistItemsItem) GetAddedAtOk() (*time.Time, bool)`

GetAddedAtOk returns a tuple with the AddedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddedAt

`func (o *ListBuyerWishlistItemsItem) SetAddedAt(v time.Time)`

SetAddedAt sets AddedAt field to given value.


### GetSlug

`func (o *ListBuyerWishlistItemsItem) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *ListBuyerWishlistItemsItem) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *ListBuyerWishlistItemsItem) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *ListBuyerWishlistItemsItem) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *ListBuyerWishlistItemsItem) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *ListBuyerWishlistItemsItem) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetTitle

`func (o *ListBuyerWishlistItemsItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListBuyerWishlistItemsItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListBuyerWishlistItemsItem) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *ListBuyerWishlistItemsItem) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *ListBuyerWishlistItemsItem) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *ListBuyerWishlistItemsItem) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


