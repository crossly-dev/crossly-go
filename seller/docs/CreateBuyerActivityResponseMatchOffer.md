# CreateBuyerActivityResponseMatchOffer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**PriceCents** | **float32** |  | 
**Condition** | Pointer to **NullableString** | Free-text on listings, an order-book grade (&#39;DS&#39;) on asks. Never null on asks; frequently null on listings, which is itself informative. | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** |  | 
**Slug** | Pointer to **NullableString** | The public slug, when this offer is a marketplace listing.  Present so Scout can buy it without parsing the URL it was given back. Null on an order-book ask, which is deliberate rather than an omission: an ask is a price in a book, not a thing with a checkout, and a Buy button on one would be a promise the market side cannot keep. | [optional] 
**Available** | Pointer to **NullableFloat32** | How many the seller has. Caps the quantity stepper honestly. | [optional] 

## Methods

### NewCreateBuyerActivityResponseMatchOffer

`func NewCreateBuyerActivityResponseMatchOffer(kind string, priceCents float32, url string, ) *CreateBuyerActivityResponseMatchOffer`

NewCreateBuyerActivityResponseMatchOffer instantiates a new CreateBuyerActivityResponseMatchOffer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerActivityResponseMatchOfferWithDefaults

`func NewCreateBuyerActivityResponseMatchOfferWithDefaults() *CreateBuyerActivityResponseMatchOffer`

NewCreateBuyerActivityResponseMatchOfferWithDefaults instantiates a new CreateBuyerActivityResponseMatchOffer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateBuyerActivityResponseMatchOffer) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateBuyerActivityResponseMatchOffer) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateBuyerActivityResponseMatchOffer) SetKind(v string)`

SetKind sets Kind field to given value.


### GetPriceCents

`func (o *CreateBuyerActivityResponseMatchOffer) GetPriceCents() float32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *CreateBuyerActivityResponseMatchOffer) GetPriceCentsOk() (*float32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *CreateBuyerActivityResponseMatchOffer) SetPriceCents(v float32)`

SetPriceCents sets PriceCents field to given value.


### GetCondition

`func (o *CreateBuyerActivityResponseMatchOffer) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CreateBuyerActivityResponseMatchOffer) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CreateBuyerActivityResponseMatchOffer) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *CreateBuyerActivityResponseMatchOffer) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *CreateBuyerActivityResponseMatchOffer) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *CreateBuyerActivityResponseMatchOffer) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetTitle

`func (o *CreateBuyerActivityResponseMatchOffer) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateBuyerActivityResponseMatchOffer) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateBuyerActivityResponseMatchOffer) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateBuyerActivityResponseMatchOffer) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CreateBuyerActivityResponseMatchOffer) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateBuyerActivityResponseMatchOffer) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetImageUrl

`func (o *CreateBuyerActivityResponseMatchOffer) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateBuyerActivityResponseMatchOffer) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateBuyerActivityResponseMatchOffer) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateBuyerActivityResponseMatchOffer) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *CreateBuyerActivityResponseMatchOffer) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *CreateBuyerActivityResponseMatchOffer) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetUrl

`func (o *CreateBuyerActivityResponseMatchOffer) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateBuyerActivityResponseMatchOffer) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateBuyerActivityResponseMatchOffer) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetSlug

`func (o *CreateBuyerActivityResponseMatchOffer) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CreateBuyerActivityResponseMatchOffer) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CreateBuyerActivityResponseMatchOffer) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *CreateBuyerActivityResponseMatchOffer) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *CreateBuyerActivityResponseMatchOffer) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *CreateBuyerActivityResponseMatchOffer) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetAvailable

`func (o *CreateBuyerActivityResponseMatchOffer) GetAvailable() float32`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *CreateBuyerActivityResponseMatchOffer) GetAvailableOk() (*float32, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *CreateBuyerActivityResponseMatchOffer) SetAvailable(v float32)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *CreateBuyerActivityResponseMatchOffer) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### SetAvailableNil

`func (o *CreateBuyerActivityResponseMatchOffer) SetAvailableNil(b bool)`

 SetAvailableNil sets the value for Available to be an explicit nil

### UnsetAvailable
`func (o *CreateBuyerActivityResponseMatchOffer) UnsetAvailable()`

UnsetAvailable ensures that no value is present for Available, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


