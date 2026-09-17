# CreateBuyerIdentifyResponseVisualMatches

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Slug** | **string** |  | 
**Title** | **string** |  | 
**PriceCents** | **float32** |  | 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**Score** | **float32** | Cosine similarity. Higher is closer. | 

## Methods

### NewCreateBuyerIdentifyResponseVisualMatches

`func NewCreateBuyerIdentifyResponseVisualMatches(slug string, title string, priceCents float32, score float32, ) *CreateBuyerIdentifyResponseVisualMatches`

NewCreateBuyerIdentifyResponseVisualMatches instantiates a new CreateBuyerIdentifyResponseVisualMatches object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerIdentifyResponseVisualMatchesWithDefaults

`func NewCreateBuyerIdentifyResponseVisualMatchesWithDefaults() *CreateBuyerIdentifyResponseVisualMatches`

NewCreateBuyerIdentifyResponseVisualMatchesWithDefaults instantiates a new CreateBuyerIdentifyResponseVisualMatches object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSlug

`func (o *CreateBuyerIdentifyResponseVisualMatches) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *CreateBuyerIdentifyResponseVisualMatches) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *CreateBuyerIdentifyResponseVisualMatches) SetSlug(v string)`

SetSlug sets Slug field to given value.


### GetTitle

`func (o *CreateBuyerIdentifyResponseVisualMatches) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateBuyerIdentifyResponseVisualMatches) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateBuyerIdentifyResponseVisualMatches) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPriceCents

`func (o *CreateBuyerIdentifyResponseVisualMatches) GetPriceCents() float32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *CreateBuyerIdentifyResponseVisualMatches) GetPriceCentsOk() (*float32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *CreateBuyerIdentifyResponseVisualMatches) SetPriceCents(v float32)`

SetPriceCents sets PriceCents field to given value.


### GetImageUrl

`func (o *CreateBuyerIdentifyResponseVisualMatches) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateBuyerIdentifyResponseVisualMatches) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateBuyerIdentifyResponseVisualMatches) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateBuyerIdentifyResponseVisualMatches) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *CreateBuyerIdentifyResponseVisualMatches) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *CreateBuyerIdentifyResponseVisualMatches) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetScore

`func (o *CreateBuyerIdentifyResponseVisualMatches) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CreateBuyerIdentifyResponseVisualMatches) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CreateBuyerIdentifyResponseVisualMatches) SetScore(v float32)`

SetScore sets Score field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


