# ListCompWatchlistsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Brand** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Condition** | Pointer to **NullableString** |  | [optional] 
**CategoryMain** | Pointer to **NullableString** |  | [optional] 
**Platform** | Pointer to **NullableString** |  | [optional] 
**MinPrice** | Pointer to **NullableString** |  | [optional] 
**MaxPrice** | Pointer to **NullableString** |  | [optional] 
**LastMatchedAt** | Pointer to **NullableTime** |  | [optional] 
**LastScrapedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewListCompWatchlistsItem

`func NewListCompWatchlistsItem(id string, name string, createdAt time.Time, userId string, ) *ListCompWatchlistsItem`

NewListCompWatchlistsItem instantiates a new ListCompWatchlistsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCompWatchlistsItemWithDefaults

`func NewListCompWatchlistsItemWithDefaults() *ListCompWatchlistsItem`

NewListCompWatchlistsItemWithDefaults instantiates a new ListCompWatchlistsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCompWatchlistsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCompWatchlistsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCompWatchlistsItem) SetId(v string)`

SetId sets Id field to given value.


### GetBrand

`func (o *ListCompWatchlistsItem) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ListCompWatchlistsItem) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ListCompWatchlistsItem) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ListCompWatchlistsItem) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *ListCompWatchlistsItem) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *ListCompWatchlistsItem) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetName

`func (o *ListCompWatchlistsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCompWatchlistsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCompWatchlistsItem) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *ListCompWatchlistsItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListCompWatchlistsItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListCompWatchlistsItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *ListCompWatchlistsItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListCompWatchlistsItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListCompWatchlistsItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCondition

`func (o *ListCompWatchlistsItem) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ListCompWatchlistsItem) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ListCompWatchlistsItem) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ListCompWatchlistsItem) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *ListCompWatchlistsItem) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *ListCompWatchlistsItem) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetCategoryMain

`func (o *ListCompWatchlistsItem) GetCategoryMain() string`

GetCategoryMain returns the CategoryMain field if non-nil, zero value otherwise.

### GetCategoryMainOk

`func (o *ListCompWatchlistsItem) GetCategoryMainOk() (*string, bool)`

GetCategoryMainOk returns a tuple with the CategoryMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryMain

`func (o *ListCompWatchlistsItem) SetCategoryMain(v string)`

SetCategoryMain sets CategoryMain field to given value.

### HasCategoryMain

`func (o *ListCompWatchlistsItem) HasCategoryMain() bool`

HasCategoryMain returns a boolean if a field has been set.

### SetCategoryMainNil

`func (o *ListCompWatchlistsItem) SetCategoryMainNil(b bool)`

 SetCategoryMainNil sets the value for CategoryMain to be an explicit nil

### UnsetCategoryMain
`func (o *ListCompWatchlistsItem) UnsetCategoryMain()`

UnsetCategoryMain ensures that no value is present for CategoryMain, not even an explicit nil
### GetPlatform

`func (o *ListCompWatchlistsItem) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListCompWatchlistsItem) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListCompWatchlistsItem) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListCompWatchlistsItem) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *ListCompWatchlistsItem) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *ListCompWatchlistsItem) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetMinPrice

`func (o *ListCompWatchlistsItem) GetMinPrice() string`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *ListCompWatchlistsItem) GetMinPriceOk() (*string, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *ListCompWatchlistsItem) SetMinPrice(v string)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *ListCompWatchlistsItem) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### SetMinPriceNil

`func (o *ListCompWatchlistsItem) SetMinPriceNil(b bool)`

 SetMinPriceNil sets the value for MinPrice to be an explicit nil

### UnsetMinPrice
`func (o *ListCompWatchlistsItem) UnsetMinPrice()`

UnsetMinPrice ensures that no value is present for MinPrice, not even an explicit nil
### GetMaxPrice

`func (o *ListCompWatchlistsItem) GetMaxPrice() string`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *ListCompWatchlistsItem) GetMaxPriceOk() (*string, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *ListCompWatchlistsItem) SetMaxPrice(v string)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *ListCompWatchlistsItem) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### SetMaxPriceNil

`func (o *ListCompWatchlistsItem) SetMaxPriceNil(b bool)`

 SetMaxPriceNil sets the value for MaxPrice to be an explicit nil

### UnsetMaxPrice
`func (o *ListCompWatchlistsItem) UnsetMaxPrice()`

UnsetMaxPrice ensures that no value is present for MaxPrice, not even an explicit nil
### GetLastMatchedAt

`func (o *ListCompWatchlistsItem) GetLastMatchedAt() time.Time`

GetLastMatchedAt returns the LastMatchedAt field if non-nil, zero value otherwise.

### GetLastMatchedAtOk

`func (o *ListCompWatchlistsItem) GetLastMatchedAtOk() (*time.Time, bool)`

GetLastMatchedAtOk returns a tuple with the LastMatchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMatchedAt

`func (o *ListCompWatchlistsItem) SetLastMatchedAt(v time.Time)`

SetLastMatchedAt sets LastMatchedAt field to given value.

### HasLastMatchedAt

`func (o *ListCompWatchlistsItem) HasLastMatchedAt() bool`

HasLastMatchedAt returns a boolean if a field has been set.

### SetLastMatchedAtNil

`func (o *ListCompWatchlistsItem) SetLastMatchedAtNil(b bool)`

 SetLastMatchedAtNil sets the value for LastMatchedAt to be an explicit nil

### UnsetLastMatchedAt
`func (o *ListCompWatchlistsItem) UnsetLastMatchedAt()`

UnsetLastMatchedAt ensures that no value is present for LastMatchedAt, not even an explicit nil
### GetLastScrapedAt

`func (o *ListCompWatchlistsItem) GetLastScrapedAt() time.Time`

GetLastScrapedAt returns the LastScrapedAt field if non-nil, zero value otherwise.

### GetLastScrapedAtOk

`func (o *ListCompWatchlistsItem) GetLastScrapedAtOk() (*time.Time, bool)`

GetLastScrapedAtOk returns a tuple with the LastScrapedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScrapedAt

`func (o *ListCompWatchlistsItem) SetLastScrapedAt(v time.Time)`

SetLastScrapedAt sets LastScrapedAt field to given value.

### HasLastScrapedAt

`func (o *ListCompWatchlistsItem) HasLastScrapedAt() bool`

HasLastScrapedAt returns a boolean if a field has been set.

### SetLastScrapedAtNil

`func (o *ListCompWatchlistsItem) SetLastScrapedAtNil(b bool)`

 SetLastScrapedAtNil sets the value for LastScrapedAt to be an explicit nil

### UnsetLastScrapedAt
`func (o *ListCompWatchlistsItem) UnsetLastScrapedAt()`

UnsetLastScrapedAt ensures that no value is present for LastScrapedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


