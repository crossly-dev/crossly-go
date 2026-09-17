# CreateInboxMessageTriageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AiIntent** | Pointer to **NullableString** |  | [optional] 
**AiUrgency** | Pointer to **NullableString** |  | [optional] 
**AiScamScore** | Pointer to **NullableFloat32** |  | [optional] 
**AiSuggestedReply** | Pointer to **NullableString** |  | [optional] 
**AiTriagedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewCreateInboxMessageTriageResponse

`func NewCreateInboxMessageTriageResponse() *CreateInboxMessageTriageResponse`

NewCreateInboxMessageTriageResponse instantiates a new CreateInboxMessageTriageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInboxMessageTriageResponseWithDefaults

`func NewCreateInboxMessageTriageResponseWithDefaults() *CreateInboxMessageTriageResponse`

NewCreateInboxMessageTriageResponseWithDefaults instantiates a new CreateInboxMessageTriageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAiIntent

`func (o *CreateInboxMessageTriageResponse) GetAiIntent() string`

GetAiIntent returns the AiIntent field if non-nil, zero value otherwise.

### GetAiIntentOk

`func (o *CreateInboxMessageTriageResponse) GetAiIntentOk() (*string, bool)`

GetAiIntentOk returns a tuple with the AiIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiIntent

`func (o *CreateInboxMessageTriageResponse) SetAiIntent(v string)`

SetAiIntent sets AiIntent field to given value.

### HasAiIntent

`func (o *CreateInboxMessageTriageResponse) HasAiIntent() bool`

HasAiIntent returns a boolean if a field has been set.

### SetAiIntentNil

`func (o *CreateInboxMessageTriageResponse) SetAiIntentNil(b bool)`

 SetAiIntentNil sets the value for AiIntent to be an explicit nil

### UnsetAiIntent
`func (o *CreateInboxMessageTriageResponse) UnsetAiIntent()`

UnsetAiIntent ensures that no value is present for AiIntent, not even an explicit nil
### GetAiUrgency

`func (o *CreateInboxMessageTriageResponse) GetAiUrgency() string`

GetAiUrgency returns the AiUrgency field if non-nil, zero value otherwise.

### GetAiUrgencyOk

`func (o *CreateInboxMessageTriageResponse) GetAiUrgencyOk() (*string, bool)`

GetAiUrgencyOk returns a tuple with the AiUrgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiUrgency

`func (o *CreateInboxMessageTriageResponse) SetAiUrgency(v string)`

SetAiUrgency sets AiUrgency field to given value.

### HasAiUrgency

`func (o *CreateInboxMessageTriageResponse) HasAiUrgency() bool`

HasAiUrgency returns a boolean if a field has been set.

### SetAiUrgencyNil

`func (o *CreateInboxMessageTriageResponse) SetAiUrgencyNil(b bool)`

 SetAiUrgencyNil sets the value for AiUrgency to be an explicit nil

### UnsetAiUrgency
`func (o *CreateInboxMessageTriageResponse) UnsetAiUrgency()`

UnsetAiUrgency ensures that no value is present for AiUrgency, not even an explicit nil
### GetAiScamScore

`func (o *CreateInboxMessageTriageResponse) GetAiScamScore() float32`

GetAiScamScore returns the AiScamScore field if non-nil, zero value otherwise.

### GetAiScamScoreOk

`func (o *CreateInboxMessageTriageResponse) GetAiScamScoreOk() (*float32, bool)`

GetAiScamScoreOk returns a tuple with the AiScamScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiScamScore

`func (o *CreateInboxMessageTriageResponse) SetAiScamScore(v float32)`

SetAiScamScore sets AiScamScore field to given value.

### HasAiScamScore

`func (o *CreateInboxMessageTriageResponse) HasAiScamScore() bool`

HasAiScamScore returns a boolean if a field has been set.

### SetAiScamScoreNil

`func (o *CreateInboxMessageTriageResponse) SetAiScamScoreNil(b bool)`

 SetAiScamScoreNil sets the value for AiScamScore to be an explicit nil

### UnsetAiScamScore
`func (o *CreateInboxMessageTriageResponse) UnsetAiScamScore()`

UnsetAiScamScore ensures that no value is present for AiScamScore, not even an explicit nil
### GetAiSuggestedReply

`func (o *CreateInboxMessageTriageResponse) GetAiSuggestedReply() string`

GetAiSuggestedReply returns the AiSuggestedReply field if non-nil, zero value otherwise.

### GetAiSuggestedReplyOk

`func (o *CreateInboxMessageTriageResponse) GetAiSuggestedReplyOk() (*string, bool)`

GetAiSuggestedReplyOk returns a tuple with the AiSuggestedReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSuggestedReply

`func (o *CreateInboxMessageTriageResponse) SetAiSuggestedReply(v string)`

SetAiSuggestedReply sets AiSuggestedReply field to given value.

### HasAiSuggestedReply

`func (o *CreateInboxMessageTriageResponse) HasAiSuggestedReply() bool`

HasAiSuggestedReply returns a boolean if a field has been set.

### SetAiSuggestedReplyNil

`func (o *CreateInboxMessageTriageResponse) SetAiSuggestedReplyNil(b bool)`

 SetAiSuggestedReplyNil sets the value for AiSuggestedReply to be an explicit nil

### UnsetAiSuggestedReply
`func (o *CreateInboxMessageTriageResponse) UnsetAiSuggestedReply()`

UnsetAiSuggestedReply ensures that no value is present for AiSuggestedReply, not even an explicit nil
### GetAiTriagedAt

`func (o *CreateInboxMessageTriageResponse) GetAiTriagedAt() time.Time`

GetAiTriagedAt returns the AiTriagedAt field if non-nil, zero value otherwise.

### GetAiTriagedAtOk

`func (o *CreateInboxMessageTriageResponse) GetAiTriagedAtOk() (*time.Time, bool)`

GetAiTriagedAtOk returns a tuple with the AiTriagedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiTriagedAt

`func (o *CreateInboxMessageTriageResponse) SetAiTriagedAt(v time.Time)`

SetAiTriagedAt sets AiTriagedAt field to given value.

### HasAiTriagedAt

`func (o *CreateInboxMessageTriageResponse) HasAiTriagedAt() bool`

HasAiTriagedAt returns a boolean if a field has been set.

### SetAiTriagedAtNil

`func (o *CreateInboxMessageTriageResponse) SetAiTriagedAtNil(b bool)`

 SetAiTriagedAtNil sets the value for AiTriagedAt to be an explicit nil

### UnsetAiTriagedAt
`func (o *CreateInboxMessageTriageResponse) UnsetAiTriagedAt()`

UnsetAiTriagedAt ensures that no value is present for AiTriagedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


