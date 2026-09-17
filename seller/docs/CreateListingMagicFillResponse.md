# CreateListingMagicFillResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filled** | **map[string]interface{}** |  | 
**Count** | **float32** |  | 
**AiAvailable** | **bool** | False when the seller has no AI provider configured — &#x60;count: 0&#x60; then means \&quot;nothing new could be resolved\&quot; rather than \&quot;already complete\&quot;, a distinction the caller (web&#39;s useMagicFill toast) needs but &#x60;count&#x60; alone can&#39;t express. True whenever a provider IS configured, even if the AI call itself failed for some other reason (network, bad key) — that&#39;s a different problem than \&quot;connect a provider.\&quot; | 

## Methods

### NewCreateListingMagicFillResponse

`func NewCreateListingMagicFillResponse(filled map[string]interface{}, count float32, aiAvailable bool, ) *CreateListingMagicFillResponse`

NewCreateListingMagicFillResponse instantiates a new CreateListingMagicFillResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateListingMagicFillResponseWithDefaults

`func NewCreateListingMagicFillResponseWithDefaults() *CreateListingMagicFillResponse`

NewCreateListingMagicFillResponseWithDefaults instantiates a new CreateListingMagicFillResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilled

`func (o *CreateListingMagicFillResponse) GetFilled() map[string]interface{}`

GetFilled returns the Filled field if non-nil, zero value otherwise.

### GetFilledOk

`func (o *CreateListingMagicFillResponse) GetFilledOk() (*map[string]interface{}, bool)`

GetFilledOk returns a tuple with the Filled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilled

`func (o *CreateListingMagicFillResponse) SetFilled(v map[string]interface{})`

SetFilled sets Filled field to given value.


### GetCount

`func (o *CreateListingMagicFillResponse) GetCount() float32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *CreateListingMagicFillResponse) GetCountOk() (*float32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *CreateListingMagicFillResponse) SetCount(v float32)`

SetCount sets Count field to given value.


### GetAiAvailable

`func (o *CreateListingMagicFillResponse) GetAiAvailable() bool`

GetAiAvailable returns the AiAvailable field if non-nil, zero value otherwise.

### GetAiAvailableOk

`func (o *CreateListingMagicFillResponse) GetAiAvailableOk() (*bool, bool)`

GetAiAvailableOk returns a tuple with the AiAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiAvailable

`func (o *CreateListingMagicFillResponse) SetAiAvailable(v bool)`

SetAiAvailable sets AiAvailable field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


