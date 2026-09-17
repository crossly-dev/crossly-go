# GetInboxResponseMessages

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**VideoUrl** | Pointer to **NullableString** |  | [optional] 
**SyncedAt** | Pointer to **NullableTime** |  | [optional] 
**SentAt** | **time.Time** |  | 
**Body** | Pointer to **NullableString** |  | [optional] 
**ConversationId** | **string** |  | 
**PlatformMessageId** | Pointer to **NullableString** |  | [optional] 
**Sender** | **string** |  | 
**Attachments** | Pointer to **[]string** |  | [optional] 
**AiIntent** | Pointer to **NullableString** |  | [optional] 
**AiUrgency** | Pointer to **NullableString** |  | [optional] 
**AiScamScore** | Pointer to **NullableFloat32** |  | [optional] 
**AiSuggestedReply** | Pointer to **NullableString** |  | [optional] 
**AiTriagedAt** | Pointer to **NullableTime** |  | [optional] 
**StickerId** | Pointer to **NullableString** |  | [optional] 
**GifUrl** | Pointer to **NullableString** |  | [optional] 
**VideoThumbUrl** | Pointer to **NullableString** |  | [optional] 
**VoiceUrl** | Pointer to **NullableString** |  | [optional] 
**VoiceDurationSeconds** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewGetInboxResponseMessages

`func NewGetInboxResponseMessages(id string, sentAt time.Time, conversationId string, sender string, ) *GetInboxResponseMessages`

NewGetInboxResponseMessages instantiates a new GetInboxResponseMessages object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxResponseMessagesWithDefaults

`func NewGetInboxResponseMessagesWithDefaults() *GetInboxResponseMessages`

NewGetInboxResponseMessagesWithDefaults instantiates a new GetInboxResponseMessages object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInboxResponseMessages) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInboxResponseMessages) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInboxResponseMessages) SetId(v string)`

SetId sets Id field to given value.


### GetVideoUrl

`func (o *GetInboxResponseMessages) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *GetInboxResponseMessages) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *GetInboxResponseMessages) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *GetInboxResponseMessages) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### SetVideoUrlNil

`func (o *GetInboxResponseMessages) SetVideoUrlNil(b bool)`

 SetVideoUrlNil sets the value for VideoUrl to be an explicit nil

### UnsetVideoUrl
`func (o *GetInboxResponseMessages) UnsetVideoUrl()`

UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
### GetSyncedAt

`func (o *GetInboxResponseMessages) GetSyncedAt() time.Time`

GetSyncedAt returns the SyncedAt field if non-nil, zero value otherwise.

### GetSyncedAtOk

`func (o *GetInboxResponseMessages) GetSyncedAtOk() (*time.Time, bool)`

GetSyncedAtOk returns a tuple with the SyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncedAt

`func (o *GetInboxResponseMessages) SetSyncedAt(v time.Time)`

SetSyncedAt sets SyncedAt field to given value.

### HasSyncedAt

`func (o *GetInboxResponseMessages) HasSyncedAt() bool`

HasSyncedAt returns a boolean if a field has been set.

### SetSyncedAtNil

`func (o *GetInboxResponseMessages) SetSyncedAtNil(b bool)`

 SetSyncedAtNil sets the value for SyncedAt to be an explicit nil

### UnsetSyncedAt
`func (o *GetInboxResponseMessages) UnsetSyncedAt()`

UnsetSyncedAt ensures that no value is present for SyncedAt, not even an explicit nil
### GetSentAt

`func (o *GetInboxResponseMessages) GetSentAt() time.Time`

GetSentAt returns the SentAt field if non-nil, zero value otherwise.

### GetSentAtOk

`func (o *GetInboxResponseMessages) GetSentAtOk() (*time.Time, bool)`

GetSentAtOk returns a tuple with the SentAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSentAt

`func (o *GetInboxResponseMessages) SetSentAt(v time.Time)`

SetSentAt sets SentAt field to given value.


### GetBody

`func (o *GetInboxResponseMessages) GetBody() string`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *GetInboxResponseMessages) GetBodyOk() (*string, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *GetInboxResponseMessages) SetBody(v string)`

SetBody sets Body field to given value.

### HasBody

`func (o *GetInboxResponseMessages) HasBody() bool`

HasBody returns a boolean if a field has been set.

### SetBodyNil

`func (o *GetInboxResponseMessages) SetBodyNil(b bool)`

 SetBodyNil sets the value for Body to be an explicit nil

### UnsetBody
`func (o *GetInboxResponseMessages) UnsetBody()`

UnsetBody ensures that no value is present for Body, not even an explicit nil
### GetConversationId

`func (o *GetInboxResponseMessages) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *GetInboxResponseMessages) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *GetInboxResponseMessages) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.


### GetPlatformMessageId

`func (o *GetInboxResponseMessages) GetPlatformMessageId() string`

GetPlatformMessageId returns the PlatformMessageId field if non-nil, zero value otherwise.

### GetPlatformMessageIdOk

`func (o *GetInboxResponseMessages) GetPlatformMessageIdOk() (*string, bool)`

GetPlatformMessageIdOk returns a tuple with the PlatformMessageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformMessageId

`func (o *GetInboxResponseMessages) SetPlatformMessageId(v string)`

SetPlatformMessageId sets PlatformMessageId field to given value.

### HasPlatformMessageId

`func (o *GetInboxResponseMessages) HasPlatformMessageId() bool`

HasPlatformMessageId returns a boolean if a field has been set.

### SetPlatformMessageIdNil

`func (o *GetInboxResponseMessages) SetPlatformMessageIdNil(b bool)`

 SetPlatformMessageIdNil sets the value for PlatformMessageId to be an explicit nil

### UnsetPlatformMessageId
`func (o *GetInboxResponseMessages) UnsetPlatformMessageId()`

UnsetPlatformMessageId ensures that no value is present for PlatformMessageId, not even an explicit nil
### GetSender

`func (o *GetInboxResponseMessages) GetSender() string`

GetSender returns the Sender field if non-nil, zero value otherwise.

### GetSenderOk

`func (o *GetInboxResponseMessages) GetSenderOk() (*string, bool)`

GetSenderOk returns a tuple with the Sender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSender

`func (o *GetInboxResponseMessages) SetSender(v string)`

SetSender sets Sender field to given value.


### GetAttachments

`func (o *GetInboxResponseMessages) GetAttachments() []string`

GetAttachments returns the Attachments field if non-nil, zero value otherwise.

### GetAttachmentsOk

`func (o *GetInboxResponseMessages) GetAttachmentsOk() (*[]string, bool)`

GetAttachmentsOk returns a tuple with the Attachments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachments

`func (o *GetInboxResponseMessages) SetAttachments(v []string)`

SetAttachments sets Attachments field to given value.

### HasAttachments

`func (o *GetInboxResponseMessages) HasAttachments() bool`

HasAttachments returns a boolean if a field has been set.

### SetAttachmentsNil

`func (o *GetInboxResponseMessages) SetAttachmentsNil(b bool)`

 SetAttachmentsNil sets the value for Attachments to be an explicit nil

### UnsetAttachments
`func (o *GetInboxResponseMessages) UnsetAttachments()`

UnsetAttachments ensures that no value is present for Attachments, not even an explicit nil
### GetAiIntent

`func (o *GetInboxResponseMessages) GetAiIntent() string`

GetAiIntent returns the AiIntent field if non-nil, zero value otherwise.

### GetAiIntentOk

`func (o *GetInboxResponseMessages) GetAiIntentOk() (*string, bool)`

GetAiIntentOk returns a tuple with the AiIntent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiIntent

`func (o *GetInboxResponseMessages) SetAiIntent(v string)`

SetAiIntent sets AiIntent field to given value.

### HasAiIntent

`func (o *GetInboxResponseMessages) HasAiIntent() bool`

HasAiIntent returns a boolean if a field has been set.

### SetAiIntentNil

`func (o *GetInboxResponseMessages) SetAiIntentNil(b bool)`

 SetAiIntentNil sets the value for AiIntent to be an explicit nil

### UnsetAiIntent
`func (o *GetInboxResponseMessages) UnsetAiIntent()`

UnsetAiIntent ensures that no value is present for AiIntent, not even an explicit nil
### GetAiUrgency

`func (o *GetInboxResponseMessages) GetAiUrgency() string`

GetAiUrgency returns the AiUrgency field if non-nil, zero value otherwise.

### GetAiUrgencyOk

`func (o *GetInboxResponseMessages) GetAiUrgencyOk() (*string, bool)`

GetAiUrgencyOk returns a tuple with the AiUrgency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiUrgency

`func (o *GetInboxResponseMessages) SetAiUrgency(v string)`

SetAiUrgency sets AiUrgency field to given value.

### HasAiUrgency

`func (o *GetInboxResponseMessages) HasAiUrgency() bool`

HasAiUrgency returns a boolean if a field has been set.

### SetAiUrgencyNil

`func (o *GetInboxResponseMessages) SetAiUrgencyNil(b bool)`

 SetAiUrgencyNil sets the value for AiUrgency to be an explicit nil

### UnsetAiUrgency
`func (o *GetInboxResponseMessages) UnsetAiUrgency()`

UnsetAiUrgency ensures that no value is present for AiUrgency, not even an explicit nil
### GetAiScamScore

`func (o *GetInboxResponseMessages) GetAiScamScore() float32`

GetAiScamScore returns the AiScamScore field if non-nil, zero value otherwise.

### GetAiScamScoreOk

`func (o *GetInboxResponseMessages) GetAiScamScoreOk() (*float32, bool)`

GetAiScamScoreOk returns a tuple with the AiScamScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiScamScore

`func (o *GetInboxResponseMessages) SetAiScamScore(v float32)`

SetAiScamScore sets AiScamScore field to given value.

### HasAiScamScore

`func (o *GetInboxResponseMessages) HasAiScamScore() bool`

HasAiScamScore returns a boolean if a field has been set.

### SetAiScamScoreNil

`func (o *GetInboxResponseMessages) SetAiScamScoreNil(b bool)`

 SetAiScamScoreNil sets the value for AiScamScore to be an explicit nil

### UnsetAiScamScore
`func (o *GetInboxResponseMessages) UnsetAiScamScore()`

UnsetAiScamScore ensures that no value is present for AiScamScore, not even an explicit nil
### GetAiSuggestedReply

`func (o *GetInboxResponseMessages) GetAiSuggestedReply() string`

GetAiSuggestedReply returns the AiSuggestedReply field if non-nil, zero value otherwise.

### GetAiSuggestedReplyOk

`func (o *GetInboxResponseMessages) GetAiSuggestedReplyOk() (*string, bool)`

GetAiSuggestedReplyOk returns a tuple with the AiSuggestedReply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiSuggestedReply

`func (o *GetInboxResponseMessages) SetAiSuggestedReply(v string)`

SetAiSuggestedReply sets AiSuggestedReply field to given value.

### HasAiSuggestedReply

`func (o *GetInboxResponseMessages) HasAiSuggestedReply() bool`

HasAiSuggestedReply returns a boolean if a field has been set.

### SetAiSuggestedReplyNil

`func (o *GetInboxResponseMessages) SetAiSuggestedReplyNil(b bool)`

 SetAiSuggestedReplyNil sets the value for AiSuggestedReply to be an explicit nil

### UnsetAiSuggestedReply
`func (o *GetInboxResponseMessages) UnsetAiSuggestedReply()`

UnsetAiSuggestedReply ensures that no value is present for AiSuggestedReply, not even an explicit nil
### GetAiTriagedAt

`func (o *GetInboxResponseMessages) GetAiTriagedAt() time.Time`

GetAiTriagedAt returns the AiTriagedAt field if non-nil, zero value otherwise.

### GetAiTriagedAtOk

`func (o *GetInboxResponseMessages) GetAiTriagedAtOk() (*time.Time, bool)`

GetAiTriagedAtOk returns a tuple with the AiTriagedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiTriagedAt

`func (o *GetInboxResponseMessages) SetAiTriagedAt(v time.Time)`

SetAiTriagedAt sets AiTriagedAt field to given value.

### HasAiTriagedAt

`func (o *GetInboxResponseMessages) HasAiTriagedAt() bool`

HasAiTriagedAt returns a boolean if a field has been set.

### SetAiTriagedAtNil

`func (o *GetInboxResponseMessages) SetAiTriagedAtNil(b bool)`

 SetAiTriagedAtNil sets the value for AiTriagedAt to be an explicit nil

### UnsetAiTriagedAt
`func (o *GetInboxResponseMessages) UnsetAiTriagedAt()`

UnsetAiTriagedAt ensures that no value is present for AiTriagedAt, not even an explicit nil
### GetStickerId

`func (o *GetInboxResponseMessages) GetStickerId() string`

GetStickerId returns the StickerId field if non-nil, zero value otherwise.

### GetStickerIdOk

`func (o *GetInboxResponseMessages) GetStickerIdOk() (*string, bool)`

GetStickerIdOk returns a tuple with the StickerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStickerId

`func (o *GetInboxResponseMessages) SetStickerId(v string)`

SetStickerId sets StickerId field to given value.

### HasStickerId

`func (o *GetInboxResponseMessages) HasStickerId() bool`

HasStickerId returns a boolean if a field has been set.

### SetStickerIdNil

`func (o *GetInboxResponseMessages) SetStickerIdNil(b bool)`

 SetStickerIdNil sets the value for StickerId to be an explicit nil

### UnsetStickerId
`func (o *GetInboxResponseMessages) UnsetStickerId()`

UnsetStickerId ensures that no value is present for StickerId, not even an explicit nil
### GetGifUrl

`func (o *GetInboxResponseMessages) GetGifUrl() string`

GetGifUrl returns the GifUrl field if non-nil, zero value otherwise.

### GetGifUrlOk

`func (o *GetInboxResponseMessages) GetGifUrlOk() (*string, bool)`

GetGifUrlOk returns a tuple with the GifUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGifUrl

`func (o *GetInboxResponseMessages) SetGifUrl(v string)`

SetGifUrl sets GifUrl field to given value.

### HasGifUrl

`func (o *GetInboxResponseMessages) HasGifUrl() bool`

HasGifUrl returns a boolean if a field has been set.

### SetGifUrlNil

`func (o *GetInboxResponseMessages) SetGifUrlNil(b bool)`

 SetGifUrlNil sets the value for GifUrl to be an explicit nil

### UnsetGifUrl
`func (o *GetInboxResponseMessages) UnsetGifUrl()`

UnsetGifUrl ensures that no value is present for GifUrl, not even an explicit nil
### GetVideoThumbUrl

`func (o *GetInboxResponseMessages) GetVideoThumbUrl() string`

GetVideoThumbUrl returns the VideoThumbUrl field if non-nil, zero value otherwise.

### GetVideoThumbUrlOk

`func (o *GetInboxResponseMessages) GetVideoThumbUrlOk() (*string, bool)`

GetVideoThumbUrlOk returns a tuple with the VideoThumbUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoThumbUrl

`func (o *GetInboxResponseMessages) SetVideoThumbUrl(v string)`

SetVideoThumbUrl sets VideoThumbUrl field to given value.

### HasVideoThumbUrl

`func (o *GetInboxResponseMessages) HasVideoThumbUrl() bool`

HasVideoThumbUrl returns a boolean if a field has been set.

### SetVideoThumbUrlNil

`func (o *GetInboxResponseMessages) SetVideoThumbUrlNil(b bool)`

 SetVideoThumbUrlNil sets the value for VideoThumbUrl to be an explicit nil

### UnsetVideoThumbUrl
`func (o *GetInboxResponseMessages) UnsetVideoThumbUrl()`

UnsetVideoThumbUrl ensures that no value is present for VideoThumbUrl, not even an explicit nil
### GetVoiceUrl

`func (o *GetInboxResponseMessages) GetVoiceUrl() string`

GetVoiceUrl returns the VoiceUrl field if non-nil, zero value otherwise.

### GetVoiceUrlOk

`func (o *GetInboxResponseMessages) GetVoiceUrlOk() (*string, bool)`

GetVoiceUrlOk returns a tuple with the VoiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceUrl

`func (o *GetInboxResponseMessages) SetVoiceUrl(v string)`

SetVoiceUrl sets VoiceUrl field to given value.

### HasVoiceUrl

`func (o *GetInboxResponseMessages) HasVoiceUrl() bool`

HasVoiceUrl returns a boolean if a field has been set.

### SetVoiceUrlNil

`func (o *GetInboxResponseMessages) SetVoiceUrlNil(b bool)`

 SetVoiceUrlNil sets the value for VoiceUrl to be an explicit nil

### UnsetVoiceUrl
`func (o *GetInboxResponseMessages) UnsetVoiceUrl()`

UnsetVoiceUrl ensures that no value is present for VoiceUrl, not even an explicit nil
### GetVoiceDurationSeconds

`func (o *GetInboxResponseMessages) GetVoiceDurationSeconds() float32`

GetVoiceDurationSeconds returns the VoiceDurationSeconds field if non-nil, zero value otherwise.

### GetVoiceDurationSecondsOk

`func (o *GetInboxResponseMessages) GetVoiceDurationSecondsOk() (*float32, bool)`

GetVoiceDurationSecondsOk returns a tuple with the VoiceDurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoiceDurationSeconds

`func (o *GetInboxResponseMessages) SetVoiceDurationSeconds(v float32)`

SetVoiceDurationSeconds sets VoiceDurationSeconds field to given value.

### HasVoiceDurationSeconds

`func (o *GetInboxResponseMessages) HasVoiceDurationSeconds() bool`

HasVoiceDurationSeconds returns a boolean if a field has been set.

### SetVoiceDurationSecondsNil

`func (o *GetInboxResponseMessages) SetVoiceDurationSecondsNil(b bool)`

 SetVoiceDurationSecondsNil sets the value for VoiceDurationSeconds to be an explicit nil

### UnsetVoiceDurationSeconds
`func (o *GetInboxResponseMessages) UnsetVoiceDurationSeconds()`

UnsetVoiceDurationSeconds ensures that no value is present for VoiceDurationSeconds, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


