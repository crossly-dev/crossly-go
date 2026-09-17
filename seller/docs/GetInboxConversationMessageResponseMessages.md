# GetInboxConversationMessageResponseMessages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ConversationId** | **string** |  | 
**PlatformMessageId** | Pointer to **NullableString** |  | [optional] 
**Sender** | **string** |  | 
**Body** | Pointer to **NullableString** |  | [optional] 
**Attachments** | Pointer to **[]string** |  | [optional] 
**SentAt** | **time.Time** |  | 
**SyncedAt** | Pointer to **NullableTime** |  | [optional] 
**AiIntent** | Pointer to **NullableString** |  | [optional] 
**AiUrgency** | Pointer to **NullableString** |  | [optional] 
**AiScamScore** | Pointer to **NullableFloat32** |  | [optional] 
**AiSuggestedReply** | Pointer to **NullableString** |  | [optional] 
**AiTriagedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewGetInboxConversationMessageResponseMessages

`func NewGetInboxConversationMessageResponseMessages(id string, conversationId string, sender string, sentAt time.Time, ) *GetInboxConversationMessageResponseMessages`

NewGetInboxConversationMessageResponseMessages instantiates a new GetInboxConversationMessageResponseMessages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxConversationMessageResponseMessagesWithDefaults

`func NewGetInboxConversationMessageResponseMessagesWithDefaults() *GetInboxConversationMessageResponseMessages`

NewGetInboxConversationMessageResponseMessagesWithDefaults instantiates a new GetInboxConversationMessageResponseMessages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInboxConversationMessageResponseMessages) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInboxConversationMessageResponseMessages) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInboxConversationMessageResponseMessages) SetId(v string)`

SetId sets Id field to given value.


### GetConversationId

`func (o *GetInboxConversationMessageResponseMessages) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *GetInboxConversationMessageResponseMessages) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *GetInboxConversationMessageResponseMessages) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetPlatformMessageId

`func (o *GetInboxConversationMessageResponseMessages) GetPlatformMessageId() string`

GetPlatformMessageId returns the PlatformMessageId field if non-nil, zero value otherwise.

### GetPlatformMessageIdOk

`func (o *GetInboxConversationMessageResponseMessages) GetPlatformMessageIdOk() (*string, bool)`

GetPlatformMessageIdOk returns a tuple with the PlatformMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformMessageId

`func (o *GetInboxConversationMessageResponseMessages) SetPlatformMessageId(v string)`

SetPlatformMessageId sets PlatformMessageId field to given value.

### HasPlatformMessageId

`func (o *GetInboxConversationMessageResponseMessages) HasPlatformMessageId() bool`

HasPlatformMessageId returns a boolean if a field has been set.

### SetPlatformMessageIdNil

`func (o *GetInboxConversationMessageResponseMessages) SetPlatformMessageIdNil(b bool)`

 SetPlatformMessageIdNil sets the value for PlatformMessageId to be an explicit nil

### UnsetPlatformMessageId
`func (o *GetInboxConversationMessageResponseMessages) UnsetPlatformMessageId()`

UnsetPlatformMessageId ensures that no value is present for PlatformMessageId, not even an explicit nil
### GetSender

`func (o *GetInboxConversationMessageResponseMessages) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *GetInboxConversationMessageResponseMessages) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *GetInboxConversationMessageResponseMessages) SetSender(v string)`

SetSender sets Sender field to given value.


### GetBody

`func (o *GetInboxConversationMessageResponseMessages) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GetInboxConversationMessageResponseMessages) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GetInboxConversationMessageResponseMessages) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GetInboxConversationMessageResponseMessages) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *GetInboxConversationMessageResponseMessages) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *GetInboxConversationMessageResponseMessages) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetAttachments

`func (o *GetInboxConversationMessageResponseMessages) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetInboxConversationMessageResponseMessages) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetInboxConversationMessageResponseMessages) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetInboxConversationMessageResponseMessages) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *GetInboxConversationMessageResponseMessages) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *GetInboxConversationMessageResponseMessages) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetSentAt

`func (o *GetInboxConversationMessageResponseMessages) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *GetInboxConversationMessageResponseMessages) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *GetInboxConversationMessageResponseMessages) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.


### GetSyncedAt

`func (o *GetInboxConversationMessageResponseMessages) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *GetInboxConversationMessageResponseMessages) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *GetInboxConversationMessageResponseMessages) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *GetInboxConversationMessageResponseMessages) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.

### SetSyncedAtNil

`func (o *GetInboxConversationMessageResponseMessages) SetSyncedAtNil(b bool)`

 SetSyncedAtNil sets the value for SyncedAt to be an explicit nil

### UnsetSyncedAt
`func (o *GetInboxConversationMessageResponseMessages) UnsetSyncedAt()`

UnsetSyncedAt ensures that no value is present for SyncedAt, not even an explicit nil
### GetAiIntent

`func (o *GetInboxConversationMessageResponseMessages) GetAiIntent() string`

GetAiIntent returns the AiIntent field if non-nil, zero value otherwise.

### GetAiIntentOk

`func (o *GetInboxConversationMessageResponseMessages) GetAiIntentOk() (*string, bool)`

GetAiIntentOk returns a tuple with the AiIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiIntent

`func (o *GetInboxConversationMessageResponseMessages) SetAiIntent(v string)`

SetAiIntent sets AiIntent field to given value.

### HasAiIntent

`func (o *GetInboxConversationMessageResponseMessages) HasAiIntent() bool`

HasAiIntent returns a boolean if a field has been set.

### SetAiIntentNil

`func (o *GetInboxConversationMessageResponseMessages) SetAiIntentNil(b bool)`

 SetAiIntentNil sets the value for AiIntent to be an explicit nil

### UnsetAiIntent
`func (o *GetInboxConversationMessageResponseMessages) UnsetAiIntent()`

UnsetAiIntent ensures that no value is present for AiIntent, not even an explicit nil
### GetAiUrgency

`func (o *GetInboxConversationMessageResponseMessages) GetAiUrgency() string`

GetAiUrgency returns the AiUrgency field if non-nil, zero value otherwise.

### GetAiUrgencyOk

`func (o *GetInboxConversationMessageResponseMessages) GetAiUrgencyOk() (*string, bool)`

GetAiUrgencyOk returns a tuple with the AiUrgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiUrgency

`func (o *GetInboxConversationMessageResponseMessages) SetAiUrgency(v string)`

SetAiUrgency sets AiUrgency field to given value.

### HasAiUrgency

`func (o *GetInboxConversationMessageResponseMessages) HasAiUrgency() bool`

HasAiUrgency returns a boolean if a field has been set.

### SetAiUrgencyNil

`func (o *GetInboxConversationMessageResponseMessages) SetAiUrgencyNil(b bool)`

 SetAiUrgencyNil sets the value for AiUrgency to be an explicit nil

### UnsetAiUrgency
`func (o *GetInboxConversationMessageResponseMessages) UnsetAiUrgency()`

UnsetAiUrgency ensures that no value is present for AiUrgency, not even an explicit nil
### GetAiScamScore

`func (o *GetInboxConversationMessageResponseMessages) GetAiScamScore() float32`

GetAiScamScore returns the AiScamScore field if non-nil, zero value otherwise.

### GetAiScamScoreOk

`func (o *GetInboxConversationMessageResponseMessages) GetAiScamScoreOk() (*float32, bool)`

GetAiScamScoreOk returns a tuple with the AiScamScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiScamScore

`func (o *GetInboxConversationMessageResponseMessages) SetAiScamScore(v float32)`

SetAiScamScore sets AiScamScore field to given value.

### HasAiScamScore

`func (o *GetInboxConversationMessageResponseMessages) HasAiScamScore() bool`

HasAiScamScore returns a boolean if a field has been set.

### SetAiScamScoreNil

`func (o *GetInboxConversationMessageResponseMessages) SetAiScamScoreNil(b bool)`

 SetAiScamScoreNil sets the value for AiScamScore to be an explicit nil

### UnsetAiScamScore
`func (o *GetInboxConversationMessageResponseMessages) UnsetAiScamScore()`

UnsetAiScamScore ensures that no value is present for AiScamScore, not even an explicit nil
### GetAiSuggestedReply

`func (o *GetInboxConversationMessageResponseMessages) GetAiSuggestedReply() string`

GetAiSuggestedReply returns the AiSuggestedReply field if non-nil, zero value otherwise.

### GetAiSuggestedReplyOk

`func (o *GetInboxConversationMessageResponseMessages) GetAiSuggestedReplyOk() (*string, bool)`

GetAiSuggestedReplyOk returns a tuple with the AiSuggestedReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSuggestedReply

`func (o *GetInboxConversationMessageResponseMessages) SetAiSuggestedReply(v string)`

SetAiSuggestedReply sets AiSuggestedReply field to given value.

### HasAiSuggestedReply

`func (o *GetInboxConversationMessageResponseMessages) HasAiSuggestedReply() bool`

HasAiSuggestedReply returns a boolean if a field has been set.

### SetAiSuggestedReplyNil

`func (o *GetInboxConversationMessageResponseMessages) SetAiSuggestedReplyNil(b bool)`

 SetAiSuggestedReplyNil sets the value for AiSuggestedReply to be an explicit nil

### UnsetAiSuggestedReply
`func (o *GetInboxConversationMessageResponseMessages) UnsetAiSuggestedReply()`

UnsetAiSuggestedReply ensures that no value is present for AiSuggestedReply, not even an explicit nil
### GetAiTriagedAt

`func (o *GetInboxConversationMessageResponseMessages) GetAiTriagedAt() time.Time`

GetAiTriagedAt returns the AiTriagedAt field if non-nil, zero value otherwise.

### GetAiTriagedAtOk

`func (o *GetInboxConversationMessageResponseMessages) GetAiTriagedAtOk() (*time.Time, bool)`

GetAiTriagedAtOk returns a tuple with the AiTriagedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiTriagedAt

`func (o *GetInboxConversationMessageResponseMessages) SetAiTriagedAt(v time.Time)`

SetAiTriagedAt sets AiTriagedAt field to given value.

### HasAiTriagedAt

`func (o *GetInboxConversationMessageResponseMessages) HasAiTriagedAt() bool`

HasAiTriagedAt returns a boolean if a field has been set.

### SetAiTriagedAtNil

`func (o *GetInboxConversationMessageResponseMessages) SetAiTriagedAtNil(b bool)`

 SetAiTriagedAtNil sets the value for AiTriagedAt to be an explicit nil

### UnsetAiTriagedAt
`func (o *GetInboxConversationMessageResponseMessages) UnsetAiTriagedAt()`

UnsetAiTriagedAt ensures that no value is present for AiTriagedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


