# CreateListingCheckDuplicateResponseMatches

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Suggested** | **string** | What we&#39;d offer to do about this match. A suggestion for which button to  feature — never a decision. All three actions stay available. | 
**VariationGroupId** | Pointer to **NullableString** | The variation group to ADD to, when the match already belongs to one.  Null means there is no group yet and choosing &#x60;variation&#x60; creates one from  the match plus the new listing. Without this the UI has to guess, and  guessing wrong means either a second group beside the first or a silent  no-op. | [optional] 
**ListingId** | Pointer to **NullableString** | The seller&#39;s existing listing this scan probably duplicates (null if the  match landed only on an inventory item with no listing row). | [optional] 
**InventoryItemId** | Pointer to **NullableString** | The inventory item behind that listing, when linked. Drives the  \&quot;View inventory\&quot; button. | [optional] 
**Title** | **string** |  | 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**MatchType** | **string** |  | 
**Score** | **float32** | 0–1 confidence. Image matches report 1; title matches the similarity. | 

## Methods

### NewCreateListingCheckDuplicateResponseMatches

`func NewCreateListingCheckDuplicateResponseMatches(suggested string, title string, matchType string, score float32, ) *CreateListingCheckDuplicateResponseMatches`

NewCreateListingCheckDuplicateResponseMatches instantiates a new CreateListingCheckDuplicateResponseMatches object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateListingCheckDuplicateResponseMatchesWithDefaults

`func NewCreateListingCheckDuplicateResponseMatchesWithDefaults() *CreateListingCheckDuplicateResponseMatches`

NewCreateListingCheckDuplicateResponseMatchesWithDefaults instantiates a new CreateListingCheckDuplicateResponseMatches object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuggested

`func (o *CreateListingCheckDuplicateResponseMatches) GetSuggested() string`

GetSuggested returns the Suggested field if non-nil, zero value otherwise.

### GetSuggestedOk

`func (o *CreateListingCheckDuplicateResponseMatches) GetSuggestedOk() (*string, bool)`

GetSuggestedOk returns a tuple with the Suggested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuggested

`func (o *CreateListingCheckDuplicateResponseMatches) SetSuggested(v string)`

SetSuggested sets Suggested field to given value.


### GetVariationGroupId

`func (o *CreateListingCheckDuplicateResponseMatches) GetVariationGroupId() string`

GetVariationGroupId returns the VariationGroupId field if non-nil, zero value otherwise.

### GetVariationGroupIdOk

`func (o *CreateListingCheckDuplicateResponseMatches) GetVariationGroupIdOk() (*string, bool)`

GetVariationGroupIdOk returns a tuple with the VariationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationGroupId

`func (o *CreateListingCheckDuplicateResponseMatches) SetVariationGroupId(v string)`

SetVariationGroupId sets VariationGroupId field to given value.

### HasVariationGroupId

`func (o *CreateListingCheckDuplicateResponseMatches) HasVariationGroupId() bool`

HasVariationGroupId returns a boolean if a field has been set.

### SetVariationGroupIdNil

`func (o *CreateListingCheckDuplicateResponseMatches) SetVariationGroupIdNil(b bool)`

 SetVariationGroupIdNil sets the value for VariationGroupId to be an explicit nil

### UnsetVariationGroupId
`func (o *CreateListingCheckDuplicateResponseMatches) UnsetVariationGroupId()`

UnsetVariationGroupId ensures that no value is present for VariationGroupId, not even an explicit nil
### GetListingId

`func (o *CreateListingCheckDuplicateResponseMatches) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *CreateListingCheckDuplicateResponseMatches) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *CreateListingCheckDuplicateResponseMatches) SetListingId(v string)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *CreateListingCheckDuplicateResponseMatches) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### SetListingIdNil

`func (o *CreateListingCheckDuplicateResponseMatches) SetListingIdNil(b bool)`

 SetListingIdNil sets the value for ListingId to be an explicit nil

### UnsetListingId
`func (o *CreateListingCheckDuplicateResponseMatches) UnsetListingId()`

UnsetListingId ensures that no value is present for ListingId, not even an explicit nil
### GetInventoryItemId

`func (o *CreateListingCheckDuplicateResponseMatches) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *CreateListingCheckDuplicateResponseMatches) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *CreateListingCheckDuplicateResponseMatches) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *CreateListingCheckDuplicateResponseMatches) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### SetInventoryItemIdNil

`func (o *CreateListingCheckDuplicateResponseMatches) SetInventoryItemIdNil(b bool)`

 SetInventoryItemIdNil sets the value for InventoryItemId to be an explicit nil

### UnsetInventoryItemId
`func (o *CreateListingCheckDuplicateResponseMatches) UnsetInventoryItemId()`

UnsetInventoryItemId ensures that no value is present for InventoryItemId, not even an explicit nil
### GetTitle

`func (o *CreateListingCheckDuplicateResponseMatches) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateListingCheckDuplicateResponseMatches) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateListingCheckDuplicateResponseMatches) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetImageUrl

`func (o *CreateListingCheckDuplicateResponseMatches) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateListingCheckDuplicateResponseMatches) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateListingCheckDuplicateResponseMatches) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateListingCheckDuplicateResponseMatches) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *CreateListingCheckDuplicateResponseMatches) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *CreateListingCheckDuplicateResponseMatches) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetMatchType

`func (o *CreateListingCheckDuplicateResponseMatches) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *CreateListingCheckDuplicateResponseMatches) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *CreateListingCheckDuplicateResponseMatches) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.


### GetScore

`func (o *CreateListingCheckDuplicateResponseMatches) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CreateListingCheckDuplicateResponseMatches) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CreateListingCheckDuplicateResponseMatches) SetScore(v float32)`

SetScore sets Score field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


