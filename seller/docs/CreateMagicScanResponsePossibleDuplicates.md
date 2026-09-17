# CreateMagicScanResponsePossibleDuplicates

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListingId** | Pointer to **NullableString** | The seller&#39;s existing listing this scan probably duplicates (null if the  match landed only on an inventory item with no listing row). | [optional] 
**InventoryItemId** | Pointer to **NullableString** | The inventory item behind that listing, when linked. Drives the  \&quot;View inventory\&quot; button. | [optional] 
**Title** | **string** |  | 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**MatchType** | **string** |  | 
**Score** | **float32** | 0–1 confidence. Image matches report 1; title matches the similarity. | 
**VariationGroupId** | Pointer to **NullableString** | The variation group the matched listing already belongs to, if any.    This is what turns \&quot;you already have this\&quot; into something useful for a  seller scanning a size run. Scan the Medium, scan the Large, and the  second scan lands here — the honest answer is usually neither \&quot;it&#39;s the  same one\&quot; nor \&quot;it&#39;s different\&quot;, it&#39;s \&quot;it&#39;s another size of that\&quot;. Which  of the two offers to make depends entirely on this field:      null      → offer to CREATE a group from the match and the new listing    set       → offer to ADD the new listing to the group that exists    Without it the UI would have to guess, and guessing wrong means either a  second group beside the first or a silent no-op. | [optional] 

## Methods

### NewCreateMagicScanResponsePossibleDuplicates

`func NewCreateMagicScanResponsePossibleDuplicates(title string, matchType string, score float32, ) *CreateMagicScanResponsePossibleDuplicates`

NewCreateMagicScanResponsePossibleDuplicates instantiates a new CreateMagicScanResponsePossibleDuplicates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMagicScanResponsePossibleDuplicatesWithDefaults

`func NewCreateMagicScanResponsePossibleDuplicatesWithDefaults() *CreateMagicScanResponsePossibleDuplicates`

NewCreateMagicScanResponsePossibleDuplicatesWithDefaults instantiates a new CreateMagicScanResponsePossibleDuplicates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingId

`func (o *CreateMagicScanResponsePossibleDuplicates) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *CreateMagicScanResponsePossibleDuplicates) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *CreateMagicScanResponsePossibleDuplicates) SetListingId(v string)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *CreateMagicScanResponsePossibleDuplicates) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### SetListingIdNil

`func (o *CreateMagicScanResponsePossibleDuplicates) SetListingIdNil(b bool)`

 SetListingIdNil sets the value for ListingId to be an explicit nil

### UnsetListingId
`func (o *CreateMagicScanResponsePossibleDuplicates) UnsetListingId()`

UnsetListingId ensures that no value is present for ListingId, not even an explicit nil
### GetInventoryItemId

`func (o *CreateMagicScanResponsePossibleDuplicates) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *CreateMagicScanResponsePossibleDuplicates) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *CreateMagicScanResponsePossibleDuplicates) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *CreateMagicScanResponsePossibleDuplicates) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### SetInventoryItemIdNil

`func (o *CreateMagicScanResponsePossibleDuplicates) SetInventoryItemIdNil(b bool)`

 SetInventoryItemIdNil sets the value for InventoryItemId to be an explicit nil

### UnsetInventoryItemId
`func (o *CreateMagicScanResponsePossibleDuplicates) UnsetInventoryItemId()`

UnsetInventoryItemId ensures that no value is present for InventoryItemId, not even an explicit nil
### GetTitle

`func (o *CreateMagicScanResponsePossibleDuplicates) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateMagicScanResponsePossibleDuplicates) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateMagicScanResponsePossibleDuplicates) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetImageUrl

`func (o *CreateMagicScanResponsePossibleDuplicates) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *CreateMagicScanResponsePossibleDuplicates) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *CreateMagicScanResponsePossibleDuplicates) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *CreateMagicScanResponsePossibleDuplicates) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *CreateMagicScanResponsePossibleDuplicates) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *CreateMagicScanResponsePossibleDuplicates) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetMatchType

`func (o *CreateMagicScanResponsePossibleDuplicates) GetMatchType() string`

GetMatchType returns the MatchType field if non-nil, zero value otherwise.

### GetMatchTypeOk

`func (o *CreateMagicScanResponsePossibleDuplicates) GetMatchTypeOk() (*string, bool)`

GetMatchTypeOk returns a tuple with the MatchType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchType

`func (o *CreateMagicScanResponsePossibleDuplicates) SetMatchType(v string)`

SetMatchType sets MatchType field to given value.


### GetScore

`func (o *CreateMagicScanResponsePossibleDuplicates) GetScore() float32`

GetScore returns the Score field if non-nil, zero value otherwise.

### GetScoreOk

`func (o *CreateMagicScanResponsePossibleDuplicates) GetScoreOk() (*float32, bool)`

GetScoreOk returns a tuple with the Score field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScore

`func (o *CreateMagicScanResponsePossibleDuplicates) SetScore(v float32)`

SetScore sets Score field to given value.


### GetVariationGroupId

`func (o *CreateMagicScanResponsePossibleDuplicates) GetVariationGroupId() string`

GetVariationGroupId returns the VariationGroupId field if non-nil, zero value otherwise.

### GetVariationGroupIdOk

`func (o *CreateMagicScanResponsePossibleDuplicates) GetVariationGroupIdOk() (*string, bool)`

GetVariationGroupIdOk returns a tuple with the VariationGroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariationGroupId

`func (o *CreateMagicScanResponsePossibleDuplicates) SetVariationGroupId(v string)`

SetVariationGroupId sets VariationGroupId field to given value.

### HasVariationGroupId

`func (o *CreateMagicScanResponsePossibleDuplicates) HasVariationGroupId() bool`

HasVariationGroupId returns a boolean if a field has been set.

### SetVariationGroupIdNil

`func (o *CreateMagicScanResponsePossibleDuplicates) SetVariationGroupIdNil(b bool)`

 SetVariationGroupIdNil sets the value for VariationGroupId to be an explicit nil

### UnsetVariationGroupId
`func (o *CreateMagicScanResponsePossibleDuplicates) UnsetVariationGroupId()`

UnsetVariationGroupId ensures that no value is present for VariationGroupId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


