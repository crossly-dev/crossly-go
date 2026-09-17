# ListCompWatchlistRecentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Platform** | **string** |  | 
**Title** | **string** |  | 
**Brand** | Pointer to **NullableString** |  | [optional] 
**SoldPrice** | **string** |  | 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**ListingUrl** | Pointer to **NullableString** |  | [optional] 
**ScrapedAt** | **string** |  | 
**SoldAtApprox** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListCompWatchlistRecentItem

`func NewListCompWatchlistRecentItem(id string, platform string, title string, soldPrice string, scrapedAt string, ) *ListCompWatchlistRecentItem`

NewListCompWatchlistRecentItem instantiates a new ListCompWatchlistRecentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCompWatchlistRecentItemWithDefaults

`func NewListCompWatchlistRecentItemWithDefaults() *ListCompWatchlistRecentItem`

NewListCompWatchlistRecentItemWithDefaults instantiates a new ListCompWatchlistRecentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCompWatchlistRecentItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCompWatchlistRecentItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCompWatchlistRecentItem) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *ListCompWatchlistRecentItem) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListCompWatchlistRecentItem) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListCompWatchlistRecentItem) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetTitle

`func (o *ListCompWatchlistRecentItem) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ListCompWatchlistRecentItem) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ListCompWatchlistRecentItem) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetBrand

`func (o *ListCompWatchlistRecentItem) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ListCompWatchlistRecentItem) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ListCompWatchlistRecentItem) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ListCompWatchlistRecentItem) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *ListCompWatchlistRecentItem) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *ListCompWatchlistRecentItem) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetSoldPrice

`func (o *ListCompWatchlistRecentItem) GetSoldPrice() string`

GetSoldPrice returns the SoldPrice field if non-nil, zero value otherwise.

### GetSoldPriceOk

`func (o *ListCompWatchlistRecentItem) GetSoldPriceOk() (*string, bool)`

GetSoldPriceOk returns a tuple with the SoldPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldPrice

`func (o *ListCompWatchlistRecentItem) SetSoldPrice(v string)`

SetSoldPrice sets SoldPrice field to given value.


### GetImageUrl

`func (o *ListCompWatchlistRecentItem) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ListCompWatchlistRecentItem) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ListCompWatchlistRecentItem) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ListCompWatchlistRecentItem) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *ListCompWatchlistRecentItem) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *ListCompWatchlistRecentItem) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetListingUrl

`func (o *ListCompWatchlistRecentItem) GetListingUrl() string`

GetListingUrl returns the ListingUrl field if non-nil, zero value otherwise.

### GetListingUrlOk

`func (o *ListCompWatchlistRecentItem) GetListingUrlOk() (*string, bool)`

GetListingUrlOk returns a tuple with the ListingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingUrl

`func (o *ListCompWatchlistRecentItem) SetListingUrl(v string)`

SetListingUrl sets ListingUrl field to given value.

### HasListingUrl

`func (o *ListCompWatchlistRecentItem) HasListingUrl() bool`

HasListingUrl returns a boolean if a field has been set.

### SetListingUrlNil

`func (o *ListCompWatchlistRecentItem) SetListingUrlNil(b bool)`

 SetListingUrlNil sets the value for ListingUrl to be an explicit nil

### UnsetListingUrl
`func (o *ListCompWatchlistRecentItem) UnsetListingUrl()`

UnsetListingUrl ensures that no value is present for ListingUrl, not even an explicit nil
### GetScrapedAt

`func (o *ListCompWatchlistRecentItem) GetScrapedAt() string`

GetScrapedAt returns the ScrapedAt field if non-nil, zero value otherwise.

### GetScrapedAtOk

`func (o *ListCompWatchlistRecentItem) GetScrapedAtOk() (*string, bool)`

GetScrapedAtOk returns a tuple with the ScrapedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScrapedAt

`func (o *ListCompWatchlistRecentItem) SetScrapedAt(v string)`

SetScrapedAt sets ScrapedAt field to given value.


### GetSoldAtApprox

`func (o *ListCompWatchlistRecentItem) GetSoldAtApprox() string`

GetSoldAtApprox returns the SoldAtApprox field if non-nil, zero value otherwise.

### GetSoldAtApproxOk

`func (o *ListCompWatchlistRecentItem) GetSoldAtApproxOk() (*string, bool)`

GetSoldAtApproxOk returns a tuple with the SoldAtApprox field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldAtApprox

`func (o *ListCompWatchlistRecentItem) SetSoldAtApprox(v string)`

SetSoldAtApprox sets SoldAtApprox field to given value.

### HasSoldAtApprox

`func (o *ListCompWatchlistRecentItem) HasSoldAtApprox() bool`

HasSoldAtApprox returns a boolean if a field has been set.

### SetSoldAtApproxNil

`func (o *ListCompWatchlistRecentItem) SetSoldAtApproxNil(b bool)`

 SetSoldAtApproxNil sets the value for SoldAtApprox to be an explicit nil

### UnsetSoldAtApprox
`func (o *ListCompWatchlistRecentItem) UnsetSoldAtApprox()`

UnsetSoldAtApprox ensures that no value is present for SoldAtApprox, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


