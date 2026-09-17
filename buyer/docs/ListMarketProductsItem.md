# ListMarketProductsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Brand** | Pointer to **NullableString** |  | [optional] 
**Model** | Pointer to **NullableString** |  | [optional] 
**CategorySlug** | **string** |  | 
**StyleCode** | Pointer to **NullableString** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**LowestAskCentsFrom** | Pointer to **NullableFloat32** |  | [optional] 
**HighestBidCents** | Pointer to **NullableFloat32** |  | [optional] 
**LastSaleCents** | Pointer to **NullableFloat32** |  | [optional] 
**TradesCount** | **float32** |  | 
**GradedGradeKeys** | **[]string** |  | 

## Methods

### NewListMarketProductsItem

`func NewListMarketProductsItem(id string, name string, categorySlug string, tradesCount float32, gradedGradeKeys []string, ) *ListMarketProductsItem`

NewListMarketProductsItem instantiates a new ListMarketProductsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMarketProductsItemWithDefaults

`func NewListMarketProductsItemWithDefaults() *ListMarketProductsItem`

NewListMarketProductsItemWithDefaults instantiates a new ListMarketProductsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListMarketProductsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListMarketProductsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListMarketProductsItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListMarketProductsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListMarketProductsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListMarketProductsItem) SetName(v string)`

SetName sets Name field to given value.


### GetBrand

`func (o *ListMarketProductsItem) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ListMarketProductsItem) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ListMarketProductsItem) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ListMarketProductsItem) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *ListMarketProductsItem) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *ListMarketProductsItem) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetModel

`func (o *ListMarketProductsItem) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *ListMarketProductsItem) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *ListMarketProductsItem) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *ListMarketProductsItem) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *ListMarketProductsItem) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *ListMarketProductsItem) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetCategorySlug

`func (o *ListMarketProductsItem) GetCategorySlug() string`

GetCategorySlug returns the CategorySlug field if non-nil, zero value otherwise.

### GetCategorySlugOk

`func (o *ListMarketProductsItem) GetCategorySlugOk() (*string, bool)`

GetCategorySlugOk returns a tuple with the CategorySlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySlug

`func (o *ListMarketProductsItem) SetCategorySlug(v string)`

SetCategorySlug sets CategorySlug field to given value.


### GetStyleCode

`func (o *ListMarketProductsItem) GetStyleCode() string`

GetStyleCode returns the StyleCode field if non-nil, zero value otherwise.

### GetStyleCodeOk

`func (o *ListMarketProductsItem) GetStyleCodeOk() (*string, bool)`

GetStyleCodeOk returns a tuple with the StyleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleCode

`func (o *ListMarketProductsItem) SetStyleCode(v string)`

SetStyleCode sets StyleCode field to given value.

### HasStyleCode

`func (o *ListMarketProductsItem) HasStyleCode() bool`

HasStyleCode returns a boolean if a field has been set.

### SetStyleCodeNil

`func (o *ListMarketProductsItem) SetStyleCodeNil(b bool)`

 SetStyleCodeNil sets the value for StyleCode to be an explicit nil

### UnsetStyleCode
`func (o *ListMarketProductsItem) UnsetStyleCode()`

UnsetStyleCode ensures that no value is present for StyleCode, not even an explicit nil
### GetImageUrl

`func (o *ListMarketProductsItem) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ListMarketProductsItem) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ListMarketProductsItem) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *ListMarketProductsItem) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *ListMarketProductsItem) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *ListMarketProductsItem) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetLowestAskCentsFrom

`func (o *ListMarketProductsItem) GetLowestAskCentsFrom() float32`

GetLowestAskCentsFrom returns the LowestAskCentsFrom field if non-nil, zero value otherwise.

### GetLowestAskCentsFromOk

`func (o *ListMarketProductsItem) GetLowestAskCentsFromOk() (*float32, bool)`

GetLowestAskCentsFromOk returns a tuple with the LowestAskCentsFrom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestAskCentsFrom

`func (o *ListMarketProductsItem) SetLowestAskCentsFrom(v float32)`

SetLowestAskCentsFrom sets LowestAskCentsFrom field to given value.

### HasLowestAskCentsFrom

`func (o *ListMarketProductsItem) HasLowestAskCentsFrom() bool`

HasLowestAskCentsFrom returns a boolean if a field has been set.

### SetLowestAskCentsFromNil

`func (o *ListMarketProductsItem) SetLowestAskCentsFromNil(b bool)`

 SetLowestAskCentsFromNil sets the value for LowestAskCentsFrom to be an explicit nil

### UnsetLowestAskCentsFrom
`func (o *ListMarketProductsItem) UnsetLowestAskCentsFrom()`

UnsetLowestAskCentsFrom ensures that no value is present for LowestAskCentsFrom, not even an explicit nil
### GetHighestBidCents

`func (o *ListMarketProductsItem) GetHighestBidCents() float32`

GetHighestBidCents returns the HighestBidCents field if non-nil, zero value otherwise.

### GetHighestBidCentsOk

`func (o *ListMarketProductsItem) GetHighestBidCentsOk() (*float32, bool)`

GetHighestBidCentsOk returns a tuple with the HighestBidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestBidCents

`func (o *ListMarketProductsItem) SetHighestBidCents(v float32)`

SetHighestBidCents sets HighestBidCents field to given value.

### HasHighestBidCents

`func (o *ListMarketProductsItem) HasHighestBidCents() bool`

HasHighestBidCents returns a boolean if a field has been set.

### SetHighestBidCentsNil

`func (o *ListMarketProductsItem) SetHighestBidCentsNil(b bool)`

 SetHighestBidCentsNil sets the value for HighestBidCents to be an explicit nil

### UnsetHighestBidCents
`func (o *ListMarketProductsItem) UnsetHighestBidCents()`

UnsetHighestBidCents ensures that no value is present for HighestBidCents, not even an explicit nil
### GetLastSaleCents

`func (o *ListMarketProductsItem) GetLastSaleCents() float32`

GetLastSaleCents returns the LastSaleCents field if non-nil, zero value otherwise.

### GetLastSaleCentsOk

`func (o *ListMarketProductsItem) GetLastSaleCentsOk() (*float32, bool)`

GetLastSaleCentsOk returns a tuple with the LastSaleCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSaleCents

`func (o *ListMarketProductsItem) SetLastSaleCents(v float32)`

SetLastSaleCents sets LastSaleCents field to given value.

### HasLastSaleCents

`func (o *ListMarketProductsItem) HasLastSaleCents() bool`

HasLastSaleCents returns a boolean if a field has been set.

### SetLastSaleCentsNil

`func (o *ListMarketProductsItem) SetLastSaleCentsNil(b bool)`

 SetLastSaleCentsNil sets the value for LastSaleCents to be an explicit nil

### UnsetLastSaleCents
`func (o *ListMarketProductsItem) UnsetLastSaleCents()`

UnsetLastSaleCents ensures that no value is present for LastSaleCents, not even an explicit nil
### GetTradesCount

`func (o *ListMarketProductsItem) GetTradesCount() float32`

GetTradesCount returns the TradesCount field if non-nil, zero value otherwise.

### GetTradesCountOk

`func (o *ListMarketProductsItem) GetTradesCountOk() (*float32, bool)`

GetTradesCountOk returns a tuple with the TradesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradesCount

`func (o *ListMarketProductsItem) SetTradesCount(v float32)`

SetTradesCount sets TradesCount field to given value.


### GetGradedGradeKeys

`func (o *ListMarketProductsItem) GetGradedGradeKeys() []string`

GetGradedGradeKeys returns the GradedGradeKeys field if non-nil, zero value otherwise.

### GetGradedGradeKeysOk

`func (o *ListMarketProductsItem) GetGradedGradeKeysOk() (*[]string, bool)`

GetGradedGradeKeysOk returns a tuple with the GradedGradeKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradedGradeKeys

`func (o *ListMarketProductsItem) SetGradedGradeKeys(v []string)`

SetGradedGradeKeys sets GradedGradeKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


