# CreateMagicScanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RunId** | **string** |  | 
**EbayMatch** | Pointer to [**NullableCreateMagicScanResponseEbayMatch**](CreateMagicScanResponseEbayMatch.md) |  | [optional] 
**TopHits** | [**[]CreateMagicScanResponseTopHits**](CreateMagicScanResponseTopHits.md) | Unified top-10-globally list, ranked by CLIP visual similarity to  the seller&#39;s source photo. Each hit carries its origin platform. | 
**EbayHits** | [**[]CreateMagicScanResponseEbayHits**](CreateMagicScanResponseEbayHits.md) | Legacy compat — UI&#39;s existing render. ebayHits now &#x3D;&#x3D; visually-  validated eBay subset; otherMatches is re-grouped from topHits. | 
**OtherMatches** | **map[string]interface{}** |  | 
**ImageUrls** | **[]string** | Every photo the seller uploaded for this scan, primary first. | 
**VisionAspects** | **map[string]interface{}** | Vision-LLM aspects extracted across all photos. Populated only  when the seller has magic-list-vision-aspects enabled + a vision  provider configured. Empty otherwise. | 
**PossibleDuplicates** | [**[]CreateMagicScanResponsePossibleDuplicates**](CreateMagicScanResponsePossibleDuplicates.md) | The seller&#39;s OWN listings/inventory that this scan probably duplicates  (image + fuzzy-title self-dedup). Empty when nothing matched. Drives the  \&quot;you may already have this\&quot; prompt. | 
**Cached** | **bool** |  | 

## Methods

### NewCreateMagicScanResponse

`func NewCreateMagicScanResponse(runId string, topHits []CreateMagicScanResponseTopHits, ebayHits []CreateMagicScanResponseEbayHits, otherMatches map[string]interface{}, imageUrls []string, visionAspects map[string]interface{}, possibleDuplicates []CreateMagicScanResponsePossibleDuplicates, cached bool, ) *CreateMagicScanResponse`

NewCreateMagicScanResponse instantiates a new CreateMagicScanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMagicScanResponseWithDefaults

`func NewCreateMagicScanResponseWithDefaults() *CreateMagicScanResponse`

NewCreateMagicScanResponseWithDefaults instantiates a new CreateMagicScanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRunId

`func (o *CreateMagicScanResponse) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *CreateMagicScanResponse) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *CreateMagicScanResponse) SetRunId(v string)`

SetRunId sets RunId field to given value.


### GetEbayMatch

`func (o *CreateMagicScanResponse) GetEbayMatch() CreateMagicScanResponseEbayMatch`

GetEbayMatch returns the EbayMatch field if non-nil, zero value otherwise.

### GetEbayMatchOk

`func (o *CreateMagicScanResponse) GetEbayMatchOk() (*CreateMagicScanResponseEbayMatch, bool)`

GetEbayMatchOk returns a tuple with the EbayMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayMatch

`func (o *CreateMagicScanResponse) SetEbayMatch(v CreateMagicScanResponseEbayMatch)`

SetEbayMatch sets EbayMatch field to given value.

### HasEbayMatch

`func (o *CreateMagicScanResponse) HasEbayMatch() bool`

HasEbayMatch returns a boolean if a field has been set.

### SetEbayMatchNil

`func (o *CreateMagicScanResponse) SetEbayMatchNil(b bool)`

 SetEbayMatchNil sets the value for EbayMatch to be an explicit nil

### UnsetEbayMatch
`func (o *CreateMagicScanResponse) UnsetEbayMatch()`

UnsetEbayMatch ensures that no value is present for EbayMatch, not even an explicit nil
### GetTopHits

`func (o *CreateMagicScanResponse) GetTopHits() []CreateMagicScanResponseTopHits`

GetTopHits returns the TopHits field if non-nil, zero value otherwise.

### GetTopHitsOk

`func (o *CreateMagicScanResponse) GetTopHitsOk() (*[]CreateMagicScanResponseTopHits, bool)`

GetTopHitsOk returns a tuple with the TopHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopHits

`func (o *CreateMagicScanResponse) SetTopHits(v []CreateMagicScanResponseTopHits)`

SetTopHits sets TopHits field to given value.


### GetEbayHits

`func (o *CreateMagicScanResponse) GetEbayHits() []CreateMagicScanResponseEbayHits`

GetEbayHits returns the EbayHits field if non-nil, zero value otherwise.

### GetEbayHitsOk

`func (o *CreateMagicScanResponse) GetEbayHitsOk() (*[]CreateMagicScanResponseEbayHits, bool)`

GetEbayHitsOk returns a tuple with the EbayHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbayHits

`func (o *CreateMagicScanResponse) SetEbayHits(v []CreateMagicScanResponseEbayHits)`

SetEbayHits sets EbayHits field to given value.


### GetOtherMatches

`func (o *CreateMagicScanResponse) GetOtherMatches() map[string]interface{}`

GetOtherMatches returns the OtherMatches field if non-nil, zero value otherwise.

### GetOtherMatchesOk

`func (o *CreateMagicScanResponse) GetOtherMatchesOk() (*map[string]interface{}, bool)`

GetOtherMatchesOk returns a tuple with the OtherMatches field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOtherMatches

`func (o *CreateMagicScanResponse) SetOtherMatches(v map[string]interface{})`

SetOtherMatches sets OtherMatches field to given value.


### GetImageUrls

`func (o *CreateMagicScanResponse) GetImageUrls() []string`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *CreateMagicScanResponse) GetImageUrlsOk() (*[]string, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *CreateMagicScanResponse) SetImageUrls(v []string)`

SetImageUrls sets ImageUrls field to given value.


### GetVisionAspects

`func (o *CreateMagicScanResponse) GetVisionAspects() map[string]interface{}`

GetVisionAspects returns the VisionAspects field if non-nil, zero value otherwise.

### GetVisionAspectsOk

`func (o *CreateMagicScanResponse) GetVisionAspectsOk() (*map[string]interface{}, bool)`

GetVisionAspectsOk returns a tuple with the VisionAspects field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisionAspects

`func (o *CreateMagicScanResponse) SetVisionAspects(v map[string]interface{})`

SetVisionAspects sets VisionAspects field to given value.


### GetPossibleDuplicates

`func (o *CreateMagicScanResponse) GetPossibleDuplicates() []CreateMagicScanResponsePossibleDuplicates`

GetPossibleDuplicates returns the PossibleDuplicates field if non-nil, zero value otherwise.

### GetPossibleDuplicatesOk

`func (o *CreateMagicScanResponse) GetPossibleDuplicatesOk() (*[]CreateMagicScanResponsePossibleDuplicates, bool)`

GetPossibleDuplicatesOk returns a tuple with the PossibleDuplicates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPossibleDuplicates

`func (o *CreateMagicScanResponse) SetPossibleDuplicates(v []CreateMagicScanResponsePossibleDuplicates)`

SetPossibleDuplicates sets PossibleDuplicates field to given value.


### GetCached

`func (o *CreateMagicScanResponse) GetCached() bool`

GetCached returns the Cached field if non-nil, zero value otherwise.

### GetCachedOk

`func (o *CreateMagicScanResponse) GetCachedOk() (*bool, bool)`

GetCachedOk returns a tuple with the Cached field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCached

`func (o *CreateMagicScanResponse) SetCached(v bool)`

SetCached sets Cached field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


