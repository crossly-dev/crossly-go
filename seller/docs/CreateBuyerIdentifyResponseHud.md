# CreateBuyerIdentifyResponseHud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headline** | **string** |  | 
**Subline** | Pointer to **NullableString** |  | [optional] 
**Facts** | [**[]CreateBuyerIdentifyResponseHudFacts**](CreateBuyerIdentifyResponseHudFacts.md) |  | 
**Action** | [**CreateBuyerIdentifyResponseHudAction**](CreateBuyerIdentifyResponseHudAction.md) |  | 
**Tone** | **string** | &#x60;win&#x60; earned a saving, &#x60;info&#x60; found something, &#x60;none&#x60; found nothing. | 
**ImageUrl** | Pointer to **NullableString** | Image for the lens card, when there is a match worth showing. | [optional] 

## Methods

### NewCreateBuyerIdentifyResponseHud

`func NewCreateBuyerIdentifyResponseHud(headline string, facts []CreateBuyerIdentifyResponseHudFacts, action CreateBuyerIdentifyResponseHudAction, tone string, ) *CreateBuyerIdentifyResponseHud`

NewCreateBuyerIdentifyResponseHud instantiates a new CreateBuyerIdentifyResponseHud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerIdentifyResponseHudWithDefaults

`func NewCreateBuyerIdentifyResponseHudWithDefaults() *CreateBuyerIdentifyResponseHud`

NewCreateBuyerIdentifyResponseHudWithDefaults instantiates a new CreateBuyerIdentifyResponseHud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeadline

`func (o *CreateBuyerIdentifyResponseHud) GetHeadline() string`

GetHeadline returns the Headline field if non-nil, zero value otherwise.

### GetHeadlineOk

`func (o *CreateBuyerIdentifyResponseHud) GetHeadlineOk() (*string, bool)`

GetHeadlineOk returns a tuple with the Headline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeadline

`func (o *CreateBuyerIdentifyResponseHud) SetHeadline(v string)`

SetHeadline sets Headline field to given value.


### GetSubline

`func (o *CreateBuyerIdentifyResponseHud) GetSubline() string`

GetSubline returns the Subline field if non-nil, zero value otherwise.

### GetSublineOk

`func (o *CreateBuyerIdentifyResponseHud) GetSublineOk() (*string, bool)`

GetSublineOk returns a tuple with the Subline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubline

`func (o *CreateBuyerIdentifyResponseHud) SetSubline(v string)`

SetSubline sets Subline field to given value.

### HasSubline

`func (o *CreateBuyerIdentifyResponseHud) HasSubline() bool`

HasSubline returns a boolean if a field has been set.

### SetSublineNil

`func (o *CreateBuyerIdentifyResponseHud) SetSublineNil(b bool)`

 SetSublineNil sets the value for Subline to be an explicit nil

### UnsetSubline
`func (o *CreateBuyerIdentifyResponseHud) UnsetSubline()`

UnsetSubline ensures that no value is present for Subline, not even an explicit nil
### GetFacts

`func (o *CreateBuyerIdentifyResponseHud) GetFacts() []CreateBuyerIdentifyResponseHudFacts`

GetFacts returns the Facts field if non-nil, zero value otherwise.

### GetFactsOk

`func (o *CreateBuyerIdentifyResponseHud) GetFactsOk() (*[]CreateBuyerIdentifyResponseHudFacts, bool)`

GetFactsOk returns a tuple with the Facts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacts

`func (o *CreateBuyerIdentifyResponseHud) SetFacts(v []CreateBuyerIdentifyResponseHudFacts)`

SetFacts sets Facts field to given value.


### GetAction

`func (o *CreateBuyerIdentifyResponseHud) GetAction() CreateBuyerIdentifyResponseHudAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CreateBuyerIdentifyResponseHud) GetActionOk() (*CreateBuyerIdentifyResponseHudAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CreateBuyerIdentifyResponseHud) SetAction(v CreateBuyerIdentifyResponseHudAction)`

SetAction sets Action field to given value.


### GetTone

`func (o *CreateBuyerIdentifyResponseHud) GetTone() string`

GetTone returns the Tone field if non-nil, zero value otherwise.

### GetToneOk

`func (o *CreateBuyerIdentifyResponseHud) GetToneOk() (*string, bool)`

GetToneOk returns a tuple with the Tone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTone

`func (o *CreateBuyerIdentifyResponseHud) SetTone(v string)`

SetTone sets Tone field to given value.


### GetImageUrl

`func (o *CreateBuyerIdentifyResponseHud) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateBuyerIdentifyResponseHud) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateBuyerIdentifyResponseHud) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateBuyerIdentifyResponseHud) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *CreateBuyerIdentifyResponseHud) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *CreateBuyerIdentifyResponseHud) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


