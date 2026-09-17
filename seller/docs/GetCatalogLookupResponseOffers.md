# GetCatalogLookupResponseOffers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**PriceCents** | **float32** |  | 
**Currency** | **string** |  | 
**Condition** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** |  | 
**ListingSlug** | Pointer to **NullableString** | Null on an order-book ask, which is a price rather than a purchasable listing. | [optional] 
**Available** | Pointer to **NullableFloat32** |  | [optional] 
**Catalog** | Pointer to [**NullableGetCatalogLookupResponseCatalog**](GetCatalogLookupResponseCatalog.md) |  | [optional] 

## Methods

### NewGetCatalogLookupResponseOffers

`func NewGetCatalogLookupResponseOffers(kind string, priceCents float32, currency string, url string, ) *GetCatalogLookupResponseOffers`

NewGetCatalogLookupResponseOffers instantiates a new GetCatalogLookupResponseOffers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogLookupResponseOffersWithDefaults

`func NewGetCatalogLookupResponseOffersWithDefaults() *GetCatalogLookupResponseOffers`

NewGetCatalogLookupResponseOffersWithDefaults instantiates a new GetCatalogLookupResponseOffers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *GetCatalogLookupResponseOffers) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetCatalogLookupResponseOffers) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetCatalogLookupResponseOffers) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPriceCents

`func (o *GetCatalogLookupResponseOffers) GetPriceCents() float32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *GetCatalogLookupResponseOffers) GetPriceCentsOk() (*float32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *GetCatalogLookupResponseOffers) SetPriceCents(v float32)`

SetPriceCents sets PriceCents field to given value.


### GetCurrency

`func (o *GetCatalogLookupResponseOffers) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetCatalogLookupResponseOffers) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetCatalogLookupResponseOffers) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCondition

`func (o *GetCatalogLookupResponseOffers) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GetCatalogLookupResponseOffers) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *GetCatalogLookupResponseOffers) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *GetCatalogLookupResponseOffers) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *GetCatalogLookupResponseOffers) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *GetCatalogLookupResponseOffers) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetTitle

`func (o *GetCatalogLookupResponseOffers) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetCatalogLookupResponseOffers) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetCatalogLookupResponseOffers) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetCatalogLookupResponseOffers) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *GetCatalogLookupResponseOffers) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *GetCatalogLookupResponseOffers) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetImageUrl

`func (o *GetCatalogLookupResponseOffers) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GetCatalogLookupResponseOffers) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GetCatalogLookupResponseOffers) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GetCatalogLookupResponseOffers) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *GetCatalogLookupResponseOffers) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *GetCatalogLookupResponseOffers) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetUrl

`func (o *GetCatalogLookupResponseOffers) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetCatalogLookupResponseOffers) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetCatalogLookupResponseOffers) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetListingSlug

`func (o *GetCatalogLookupResponseOffers) GetListingSlug() string`

GetListingSlug returns the ListingSlug field if non-nil, zero value otherwise.

### GetListingSlugOk

`func (o *GetCatalogLookupResponseOffers) GetListingSlugOk() (*string, bool)`

GetListingSlugOk returns a tuple with the ListingSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingSlug

`func (o *GetCatalogLookupResponseOffers) SetListingSlug(v string)`

SetListingSlug sets ListingSlug field to given value.

### HasListingSlug

`func (o *GetCatalogLookupResponseOffers) HasListingSlug() bool`

HasListingSlug returns a boolean if a field has been set.

### SetListingSlugNil

`func (o *GetCatalogLookupResponseOffers) SetListingSlugNil(b bool)`

 SetListingSlugNil sets the value for ListingSlug to be an explicit nil

### UnsetListingSlug
`func (o *GetCatalogLookupResponseOffers) UnsetListingSlug()`

UnsetListingSlug ensures that no value is present for ListingSlug, not even an explicit nil
### GetAvailable

`func (o *GetCatalogLookupResponseOffers) GetAvailable() float32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *GetCatalogLookupResponseOffers) GetAvailableOk() (*float32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *GetCatalogLookupResponseOffers) SetAvailable(v float32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *GetCatalogLookupResponseOffers) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *GetCatalogLookupResponseOffers) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *GetCatalogLookupResponseOffers) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil
### GetCatalog

`func (o *GetCatalogLookupResponseOffers) GetCatalog() GetCatalogLookupResponseCatalog`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *GetCatalogLookupResponseOffers) GetCatalogOk() (*GetCatalogLookupResponseCatalog, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *GetCatalogLookupResponseOffers) SetCatalog(v GetCatalogLookupResponseCatalog)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *GetCatalogLookupResponseOffers) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *GetCatalogLookupResponseOffers) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *GetCatalogLookupResponseOffers) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


