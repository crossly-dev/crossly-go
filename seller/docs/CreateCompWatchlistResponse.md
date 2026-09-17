# CreateCompWatchlistResponse

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

### NewCreateCompWatchlistResponse

`func NewCreateCompWatchlistResponse(id string, name string, createdAt time.Time, userId string, ) *CreateCompWatchlistResponse`

NewCreateCompWatchlistResponse instantiates a new CreateCompWatchlistResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCompWatchlistResponseWithDefaults

`func NewCreateCompWatchlistResponseWithDefaults() *CreateCompWatchlistResponse`

NewCreateCompWatchlistResponseWithDefaults instantiates a new CreateCompWatchlistResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateCompWatchlistResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateCompWatchlistResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateCompWatchlistResponse) SetId(v string)`

SetId sets Id field to given value.


### GetBrand

`func (o *CreateCompWatchlistResponse) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CreateCompWatchlistResponse) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CreateCompWatchlistResponse) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CreateCompWatchlistResponse) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *CreateCompWatchlistResponse) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *CreateCompWatchlistResponse) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetName

`func (o *CreateCompWatchlistResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateCompWatchlistResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateCompWatchlistResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *CreateCompWatchlistResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateCompWatchlistResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateCompWatchlistResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *CreateCompWatchlistResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateCompWatchlistResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateCompWatchlistResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetCondition

`func (o *CreateCompWatchlistResponse) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CreateCompWatchlistResponse) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CreateCompWatchlistResponse) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *CreateCompWatchlistResponse) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *CreateCompWatchlistResponse) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *CreateCompWatchlistResponse) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetCategoryMain

`func (o *CreateCompWatchlistResponse) GetCategoryMain() string`

GetCategoryMain returns the CategoryMain field if non-nil, zero value otherwise.

### GetCategoryMainOk

`func (o *CreateCompWatchlistResponse) GetCategoryMainOk() (*string, bool)`

GetCategoryMainOk returns a tuple with the CategoryMain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryMain

`func (o *CreateCompWatchlistResponse) SetCategoryMain(v string)`

SetCategoryMain sets CategoryMain field to given value.

### HasCategoryMain

`func (o *CreateCompWatchlistResponse) HasCategoryMain() bool`

HasCategoryMain returns a boolean if a field has been set.

### SetCategoryMainNil

`func (o *CreateCompWatchlistResponse) SetCategoryMainNil(b bool)`

 SetCategoryMainNil sets the value for CategoryMain to be an explicit nil

### UnsetCategoryMain
`func (o *CreateCompWatchlistResponse) UnsetCategoryMain()`

UnsetCategoryMain ensures that no value is present for CategoryMain, not even an explicit nil
### GetPlatform

`func (o *CreateCompWatchlistResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateCompWatchlistResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateCompWatchlistResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *CreateCompWatchlistResponse) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *CreateCompWatchlistResponse) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *CreateCompWatchlistResponse) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetMinPrice

`func (o *CreateCompWatchlistResponse) GetMinPrice() string`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *CreateCompWatchlistResponse) GetMinPriceOk() (*string, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *CreateCompWatchlistResponse) SetMinPrice(v string)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *CreateCompWatchlistResponse) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### SetMinPriceNil

`func (o *CreateCompWatchlistResponse) SetMinPriceNil(b bool)`

 SetMinPriceNil sets the value for MinPrice to be an explicit nil

### UnsetMinPrice
`func (o *CreateCompWatchlistResponse) UnsetMinPrice()`

UnsetMinPrice ensures that no value is present for MinPrice, not even an explicit nil
### GetMaxPrice

`func (o *CreateCompWatchlistResponse) GetMaxPrice() string`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *CreateCompWatchlistResponse) GetMaxPriceOk() (*string, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *CreateCompWatchlistResponse) SetMaxPrice(v string)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *CreateCompWatchlistResponse) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### SetMaxPriceNil

`func (o *CreateCompWatchlistResponse) SetMaxPriceNil(b bool)`

 SetMaxPriceNil sets the value for MaxPrice to be an explicit nil

### UnsetMaxPrice
`func (o *CreateCompWatchlistResponse) UnsetMaxPrice()`

UnsetMaxPrice ensures that no value is present for MaxPrice, not even an explicit nil
### GetLastMatchedAt

`func (o *CreateCompWatchlistResponse) GetLastMatchedAt() time.Time`

GetLastMatchedAt returns the LastMatchedAt field if non-nil, zero value otherwise.

### GetLastMatchedAtOk

`func (o *CreateCompWatchlistResponse) GetLastMatchedAtOk() (*time.Time, bool)`

GetLastMatchedAtOk returns a tuple with the LastMatchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMatchedAt

`func (o *CreateCompWatchlistResponse) SetLastMatchedAt(v time.Time)`

SetLastMatchedAt sets LastMatchedAt field to given value.

### HasLastMatchedAt

`func (o *CreateCompWatchlistResponse) HasLastMatchedAt() bool`

HasLastMatchedAt returns a boolean if a field has been set.

### SetLastMatchedAtNil

`func (o *CreateCompWatchlistResponse) SetLastMatchedAtNil(b bool)`

 SetLastMatchedAtNil sets the value for LastMatchedAt to be an explicit nil

### UnsetLastMatchedAt
`func (o *CreateCompWatchlistResponse) UnsetLastMatchedAt()`

UnsetLastMatchedAt ensures that no value is present for LastMatchedAt, not even an explicit nil
### GetLastScrapedAt

`func (o *CreateCompWatchlistResponse) GetLastScrapedAt() time.Time`

GetLastScrapedAt returns the LastScrapedAt field if non-nil, zero value otherwise.

### GetLastScrapedAtOk

`func (o *CreateCompWatchlistResponse) GetLastScrapedAtOk() (*time.Time, bool)`

GetLastScrapedAtOk returns a tuple with the LastScrapedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastScrapedAt

`func (o *CreateCompWatchlistResponse) SetLastScrapedAt(v time.Time)`

SetLastScrapedAt sets LastScrapedAt field to given value.

### HasLastScrapedAt

`func (o *CreateCompWatchlistResponse) HasLastScrapedAt() bool`

HasLastScrapedAt returns a boolean if a field has been set.

### SetLastScrapedAtNil

`func (o *CreateCompWatchlistResponse) SetLastScrapedAtNil(b bool)`

 SetLastScrapedAtNil sets the value for LastScrapedAt to be an explicit nil

### UnsetLastScrapedAt
`func (o *CreateCompWatchlistResponse) UnsetLastScrapedAt()`

UnsetLastScrapedAt ensures that no value is present for LastScrapedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


