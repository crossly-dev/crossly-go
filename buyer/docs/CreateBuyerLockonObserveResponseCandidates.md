# CreateBuyerLockonObserveResponseCandidates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdentifierNs** | **string** | Where it came from, so a confirmation can be validated. | 
**IdentifierValue** | **string** |  | 
**Title** | **string** |  | 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**PriceCents** | Pointer to **NullableFloat32** |  | [optional] 
**StoreName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateBuyerLockonObserveResponseCandidates

`func NewCreateBuyerLockonObserveResponseCandidates(identifierNs string, identifierValue string, title string, ) *CreateBuyerLockonObserveResponseCandidates`

NewCreateBuyerLockonObserveResponseCandidates instantiates a new CreateBuyerLockonObserveResponseCandidates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerLockonObserveResponseCandidatesWithDefaults

`func NewCreateBuyerLockonObserveResponseCandidatesWithDefaults() *CreateBuyerLockonObserveResponseCandidates`

NewCreateBuyerLockonObserveResponseCandidatesWithDefaults instantiates a new CreateBuyerLockonObserveResponseCandidates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifierNs

`func (o *CreateBuyerLockonObserveResponseCandidates) GetIdentifierNs() string`

GetIdentifierNs returns the IdentifierNs field if non-nil, zero value otherwise.

### GetIdentifierNsOk

`func (o *CreateBuyerLockonObserveResponseCandidates) GetIdentifierNsOk() (*string, bool)`

GetIdentifierNsOk returns a tuple with the IdentifierNs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierNs

`func (o *CreateBuyerLockonObserveResponseCandidates) SetIdentifierNs(v string)`

SetIdentifierNs sets IdentifierNs field to given value.


### GetIdentifierValue

`func (o *CreateBuyerLockonObserveResponseCandidates) GetIdentifierValue() string`

GetIdentifierValue returns the IdentifierValue field if non-nil, zero value otherwise.

### GetIdentifierValueOk

`func (o *CreateBuyerLockonObserveResponseCandidates) GetIdentifierValueOk() (*string, bool)`

GetIdentifierValueOk returns a tuple with the IdentifierValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierValue

`func (o *CreateBuyerLockonObserveResponseCandidates) SetIdentifierValue(v string)`

SetIdentifierValue sets IdentifierValue field to given value.


### GetTitle

`func (o *CreateBuyerLockonObserveResponseCandidates) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateBuyerLockonObserveResponseCandidates) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateBuyerLockonObserveResponseCandidates) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetImageUrl

`func (o *CreateBuyerLockonObserveResponseCandidates) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateBuyerLockonObserveResponseCandidates) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateBuyerLockonObserveResponseCandidates) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateBuyerLockonObserveResponseCandidates) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *CreateBuyerLockonObserveResponseCandidates) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *CreateBuyerLockonObserveResponseCandidates) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetPriceCents

`func (o *CreateBuyerLockonObserveResponseCandidates) GetPriceCents() float32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *CreateBuyerLockonObserveResponseCandidates) GetPriceCentsOk() (*float32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *CreateBuyerLockonObserveResponseCandidates) SetPriceCents(v float32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *CreateBuyerLockonObserveResponseCandidates) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### SetPriceCentsNil

`func (o *CreateBuyerLockonObserveResponseCandidates) SetPriceCentsNil(b bool)`

 SetPriceCentsNil sets the value for PriceCents to be an explicit nil

### UnsetPriceCents
`func (o *CreateBuyerLockonObserveResponseCandidates) UnsetPriceCents()`

UnsetPriceCents ensures that no value is present for PriceCents, not even an explicit nil
### GetStoreName

`func (o *CreateBuyerLockonObserveResponseCandidates) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *CreateBuyerLockonObserveResponseCandidates) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *CreateBuyerLockonObserveResponseCandidates) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.

### HasStoreName

`func (o *CreateBuyerLockonObserveResponseCandidates) HasStoreName() bool`

HasStoreName returns a boolean if a field has been set.

### SetStoreNameNil

`func (o *CreateBuyerLockonObserveResponseCandidates) SetStoreNameNil(b bool)`

 SetStoreNameNil sets the value for StoreName to be an explicit nil

### UnsetStoreName
`func (o *CreateBuyerLockonObserveResponseCandidates) UnsetStoreName()`

UnsetStoreName ensures that no value is present for StoreName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


