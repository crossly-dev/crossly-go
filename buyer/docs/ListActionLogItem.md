# ListActionLogItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**Source** | **string** |  | 
**Platform** | Pointer to **NullableString** |  | [optional] 
**Action** | **string** |  | 
**LatencyMs** | Pointer to **NullableFloat32** |  | [optional] 
**ErrorClass** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**UserAgent** | Pointer to **NullableString** |  | [optional] 
**OauthAppId** | Pointer to **NullableString** |  | [optional] 
**Category** | **string** |  | 
**ActorUserId** | Pointer to **NullableString** |  | [optional] 
**FinishedAt** | Pointer to **NullableTime** |  | [optional] 
**CorrelationId** | **string** |  | 
**Track** | Pointer to **NullableString** |  | [optional] 
**TargetType** | Pointer to **NullableString** |  | [optional] 
**TargetId** | Pointer to **NullableString** |  | [optional] 
**HttpStatus** | Pointer to **NullableFloat32** |  | [optional] 
**ActorEmail** | Pointer to **NullableString** | Resolved from actorUserId so the UI can say \&quot;Jane relisted this\&quot; rather than printing a UUID. Null for worker/system actions, which genuinely had no human actor. | [optional] 
**ActorDisplayName** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListActionLogItem

`func NewListActionLogItem(id string, createdAt time.Time, status string, source string, action string, category string, correlationId string, ) *ListActionLogItem`

NewListActionLogItem instantiates a new ListActionLogItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListActionLogItemWithDefaults

`func NewListActionLogItemWithDefaults() *ListActionLogItem`

NewListActionLogItemWithDefaults instantiates a new ListActionLogItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListActionLogItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListActionLogItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListActionLogItem) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListActionLogItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListActionLogItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListActionLogItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *ListActionLogItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListActionLogItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListActionLogItem) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ListActionLogItem) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ListActionLogItem) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ListActionLogItem) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetStatus

`func (o *ListActionLogItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListActionLogItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListActionLogItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSource

`func (o *ListActionLogItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListActionLogItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListActionLogItem) SetSource(v string)`

SetSource sets Source field to given value.


### GetPlatform

`func (o *ListActionLogItem) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListActionLogItem) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListActionLogItem) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListActionLogItem) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *ListActionLogItem) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *ListActionLogItem) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetAction

`func (o *ListActionLogItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ListActionLogItem) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ListActionLogItem) SetAction(v string)`

SetAction sets Action field to given value.


### GetLatencyMs

`func (o *ListActionLogItem) GetLatencyMs() float32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *ListActionLogItem) GetLatencyMsOk() (*float32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *ListActionLogItem) SetLatencyMs(v float32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *ListActionLogItem) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### SetLatencyMsNil

`func (o *ListActionLogItem) SetLatencyMsNil(b bool)`

 SetLatencyMsNil sets the value for LatencyMs to be an explicit nil

### UnsetLatencyMs
`func (o *ListActionLogItem) UnsetLatencyMs()`

UnsetLatencyMs ensures that no value is present for LatencyMs, not even an explicit nil
### GetErrorClass

`func (o *ListActionLogItem) GetErrorClass() string`

GetErrorClass returns the ErrorClass field if non-nil, zero value otherwise.

### GetErrorClassOk

`func (o *ListActionLogItem) GetErrorClassOk() (*string, bool)`

GetErrorClassOk returns a tuple with the ErrorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorClass

`func (o *ListActionLogItem) SetErrorClass(v string)`

SetErrorClass sets ErrorClass field to given value.

### HasErrorClass

`func (o *ListActionLogItem) HasErrorClass() bool`

HasErrorClass returns a boolean if a field has been set.

### SetErrorClassNil

`func (o *ListActionLogItem) SetErrorClassNil(b bool)`

 SetErrorClassNil sets the value for ErrorClass to be an explicit nil

### UnsetErrorClass
`func (o *ListActionLogItem) UnsetErrorClass()`

UnsetErrorClass ensures that no value is present for ErrorClass, not even an explicit nil
### GetErrorMessage

`func (o *ListActionLogItem) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListActionLogItem) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListActionLogItem) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListActionLogItem) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ListActionLogItem) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ListActionLogItem) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetIpAddress

`func (o *ListActionLogItem) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ListActionLogItem) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ListActionLogItem) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ListActionLogItem) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *ListActionLogItem) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *ListActionLogItem) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetUserAgent

`func (o *ListActionLogItem) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ListActionLogItem) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ListActionLogItem) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ListActionLogItem) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *ListActionLogItem) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *ListActionLogItem) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
### GetOauthAppId

`func (o *ListActionLogItem) GetOauthAppId() string`

GetOauthAppId returns the OauthAppId field if non-nil, zero value otherwise.

### GetOauthAppIdOk

`func (o *ListActionLogItem) GetOauthAppIdOk() (*string, bool)`

GetOauthAppIdOk returns a tuple with the OauthAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthAppId

`func (o *ListActionLogItem) SetOauthAppId(v string)`

SetOauthAppId sets OauthAppId field to given value.

### HasOauthAppId

`func (o *ListActionLogItem) HasOauthAppId() bool`

HasOauthAppId returns a boolean if a field has been set.

### SetOauthAppIdNil

`func (o *ListActionLogItem) SetOauthAppIdNil(b bool)`

 SetOauthAppIdNil sets the value for OauthAppId to be an explicit nil

### UnsetOauthAppId
`func (o *ListActionLogItem) UnsetOauthAppId()`

UnsetOauthAppId ensures that no value is present for OauthAppId, not even an explicit nil
### GetCategory

`func (o *ListActionLogItem) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListActionLogItem) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListActionLogItem) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetActorUserId

`func (o *ListActionLogItem) GetActorUserId() string`

GetActorUserId returns the ActorUserId field if non-nil, zero value otherwise.

### GetActorUserIdOk

`func (o *ListActionLogItem) GetActorUserIdOk() (*string, bool)`

GetActorUserIdOk returns a tuple with the ActorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorUserId

`func (o *ListActionLogItem) SetActorUserId(v string)`

SetActorUserId sets ActorUserId field to given value.

### HasActorUserId

`func (o *ListActionLogItem) HasActorUserId() bool`

HasActorUserId returns a boolean if a field has been set.

### SetActorUserIdNil

`func (o *ListActionLogItem) SetActorUserIdNil(b bool)`

 SetActorUserIdNil sets the value for ActorUserId to be an explicit nil

### UnsetActorUserId
`func (o *ListActionLogItem) UnsetActorUserId()`

UnsetActorUserId ensures that no value is present for ActorUserId, not even an explicit nil
### GetFinishedAt

`func (o *ListActionLogItem) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ListActionLogItem) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ListActionLogItem) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ListActionLogItem) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *ListActionLogItem) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *ListActionLogItem) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetCorrelationId

`func (o *ListActionLogItem) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *ListActionLogItem) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *ListActionLogItem) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.


### GetTrack

`func (o *ListActionLogItem) GetTrack() string`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *ListActionLogItem) GetTrackOk() (*string, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *ListActionLogItem) SetTrack(v string)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *ListActionLogItem) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### SetTrackNil

`func (o *ListActionLogItem) SetTrackNil(b bool)`

 SetTrackNil sets the value for Track to be an explicit nil

### UnsetTrack
`func (o *ListActionLogItem) UnsetTrack()`

UnsetTrack ensures that no value is present for Track, not even an explicit nil
### GetTargetType

`func (o *ListActionLogItem) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *ListActionLogItem) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *ListActionLogItem) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *ListActionLogItem) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *ListActionLogItem) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *ListActionLogItem) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTargetId

`func (o *ListActionLogItem) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *ListActionLogItem) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *ListActionLogItem) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *ListActionLogItem) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *ListActionLogItem) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *ListActionLogItem) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetHttpStatus

`func (o *ListActionLogItem) GetHttpStatus() float32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *ListActionLogItem) GetHttpStatusOk() (*float32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *ListActionLogItem) SetHttpStatus(v float32)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *ListActionLogItem) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### SetHttpStatusNil

`func (o *ListActionLogItem) SetHttpStatusNil(b bool)`

 SetHttpStatusNil sets the value for HttpStatus to be an explicit nil

### UnsetHttpStatus
`func (o *ListActionLogItem) UnsetHttpStatus()`

UnsetHttpStatus ensures that no value is present for HttpStatus, not even an explicit nil
### GetActorEmail

`func (o *ListActionLogItem) GetActorEmail() string`

GetActorEmail returns the ActorEmail field if non-nil, zero value otherwise.

### GetActorEmailOk

`func (o *ListActionLogItem) GetActorEmailOk() (*string, bool)`

GetActorEmailOk returns a tuple with the ActorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorEmail

`func (o *ListActionLogItem) SetActorEmail(v string)`

SetActorEmail sets ActorEmail field to given value.

### HasActorEmail

`func (o *ListActionLogItem) HasActorEmail() bool`

HasActorEmail returns a boolean if a field has been set.

### SetActorEmailNil

`func (o *ListActionLogItem) SetActorEmailNil(b bool)`

 SetActorEmailNil sets the value for ActorEmail to be an explicit nil

### UnsetActorEmail
`func (o *ListActionLogItem) UnsetActorEmail()`

UnsetActorEmail ensures that no value is present for ActorEmail, not even an explicit nil
### GetActorDisplayName

`func (o *ListActionLogItem) GetActorDisplayName() string`

GetActorDisplayName returns the ActorDisplayName field if non-nil, zero value otherwise.

### GetActorDisplayNameOk

`func (o *ListActionLogItem) GetActorDisplayNameOk() (*string, bool)`

GetActorDisplayNameOk returns a tuple with the ActorDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorDisplayName

`func (o *ListActionLogItem) SetActorDisplayName(v string)`

SetActorDisplayName sets ActorDisplayName field to given value.

### HasActorDisplayName

`func (o *ListActionLogItem) HasActorDisplayName() bool`

HasActorDisplayName returns a boolean if a field has been set.

### SetActorDisplayNameNil

`func (o *ListActionLogItem) SetActorDisplayNameNil(b bool)`

 SetActorDisplayNameNil sets the value for ActorDisplayName to be an explicit nil

### UnsetActorDisplayName
`func (o *ListActionLogItem) UnsetActorDisplayName()`

UnsetActorDisplayName ensures that no value is present for ActorDisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


