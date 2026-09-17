# GetActionLogResponse

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

### NewGetActionLogResponse

`func NewGetActionLogResponse(id string, createdAt time.Time, status string, source string, action string, category string, correlationId string, ) *GetActionLogResponse`

NewGetActionLogResponse instantiates a new GetActionLogResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetActionLogResponseWithDefaults

`func NewGetActionLogResponseWithDefaults() *GetActionLogResponse`

NewGetActionLogResponseWithDefaults instantiates a new GetActionLogResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetActionLogResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetActionLogResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetActionLogResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetActionLogResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetActionLogResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetActionLogResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *GetActionLogResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetActionLogResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetActionLogResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *GetActionLogResponse) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *GetActionLogResponse) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *GetActionLogResponse) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetStatus

`func (o *GetActionLogResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetActionLogResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetActionLogResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSource

`func (o *GetActionLogResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetActionLogResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetActionLogResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetPlatform

`func (o *GetActionLogResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetActionLogResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetActionLogResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetActionLogResponse) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *GetActionLogResponse) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *GetActionLogResponse) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetAction

`func (o *GetActionLogResponse) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetActionLogResponse) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetActionLogResponse) SetAction(v string)`

SetAction sets Action field to given value.


### GetLatencyMs

`func (o *GetActionLogResponse) GetLatencyMs() float32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *GetActionLogResponse) GetLatencyMsOk() (*float32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *GetActionLogResponse) SetLatencyMs(v float32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *GetActionLogResponse) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### SetLatencyMsNil

`func (o *GetActionLogResponse) SetLatencyMsNil(b bool)`

 SetLatencyMsNil sets the value for LatencyMs to be an explicit nil

### UnsetLatencyMs
`func (o *GetActionLogResponse) UnsetLatencyMs()`

UnsetLatencyMs ensures that no value is present for LatencyMs, not even an explicit nil
### GetErrorClass

`func (o *GetActionLogResponse) GetErrorClass() string`

GetErrorClass returns the ErrorClass field if non-nil, zero value otherwise.

### GetErrorClassOk

`func (o *GetActionLogResponse) GetErrorClassOk() (*string, bool)`

GetErrorClassOk returns a tuple with the ErrorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorClass

`func (o *GetActionLogResponse) SetErrorClass(v string)`

SetErrorClass sets ErrorClass field to given value.

### HasErrorClass

`func (o *GetActionLogResponse) HasErrorClass() bool`

HasErrorClass returns a boolean if a field has been set.

### SetErrorClassNil

`func (o *GetActionLogResponse) SetErrorClassNil(b bool)`

 SetErrorClassNil sets the value for ErrorClass to be an explicit nil

### UnsetErrorClass
`func (o *GetActionLogResponse) UnsetErrorClass()`

UnsetErrorClass ensures that no value is present for ErrorClass, not even an explicit nil
### GetErrorMessage

`func (o *GetActionLogResponse) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetActionLogResponse) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetActionLogResponse) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetActionLogResponse) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *GetActionLogResponse) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *GetActionLogResponse) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetIpAddress

`func (o *GetActionLogResponse) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *GetActionLogResponse) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *GetActionLogResponse) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *GetActionLogResponse) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *GetActionLogResponse) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *GetActionLogResponse) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetUserAgent

`func (o *GetActionLogResponse) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *GetActionLogResponse) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *GetActionLogResponse) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *GetActionLogResponse) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *GetActionLogResponse) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *GetActionLogResponse) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
### GetOauthAppId

`func (o *GetActionLogResponse) GetOauthAppId() string`

GetOauthAppId returns the OauthAppId field if non-nil, zero value otherwise.

### GetOauthAppIdOk

`func (o *GetActionLogResponse) GetOauthAppIdOk() (*string, bool)`

GetOauthAppIdOk returns a tuple with the OauthAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthAppId

`func (o *GetActionLogResponse) SetOauthAppId(v string)`

SetOauthAppId sets OauthAppId field to given value.

### HasOauthAppId

`func (o *GetActionLogResponse) HasOauthAppId() bool`

HasOauthAppId returns a boolean if a field has been set.

### SetOauthAppIdNil

`func (o *GetActionLogResponse) SetOauthAppIdNil(b bool)`

 SetOauthAppIdNil sets the value for OauthAppId to be an explicit nil

### UnsetOauthAppId
`func (o *GetActionLogResponse) UnsetOauthAppId()`

UnsetOauthAppId ensures that no value is present for OauthAppId, not even an explicit nil
### GetCategory

`func (o *GetActionLogResponse) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetActionLogResponse) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetActionLogResponse) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetActorUserId

`func (o *GetActionLogResponse) GetActorUserId() string`

GetActorUserId returns the ActorUserId field if non-nil, zero value otherwise.

### GetActorUserIdOk

`func (o *GetActionLogResponse) GetActorUserIdOk() (*string, bool)`

GetActorUserIdOk returns a tuple with the ActorUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorUserId

`func (o *GetActionLogResponse) SetActorUserId(v string)`

SetActorUserId sets ActorUserId field to given value.

### HasActorUserId

`func (o *GetActionLogResponse) HasActorUserId() bool`

HasActorUserId returns a boolean if a field has been set.

### SetActorUserIdNil

`func (o *GetActionLogResponse) SetActorUserIdNil(b bool)`

 SetActorUserIdNil sets the value for ActorUserId to be an explicit nil

### UnsetActorUserId
`func (o *GetActionLogResponse) UnsetActorUserId()`

UnsetActorUserId ensures that no value is present for ActorUserId, not even an explicit nil
### GetFinishedAt

`func (o *GetActionLogResponse) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *GetActionLogResponse) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *GetActionLogResponse) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *GetActionLogResponse) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *GetActionLogResponse) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *GetActionLogResponse) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetCorrelationId

`func (o *GetActionLogResponse) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *GetActionLogResponse) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *GetActionLogResponse) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.


### GetTrack

`func (o *GetActionLogResponse) GetTrack() string`

GetTrack returns the Track field if non-nil, zero value otherwise.

### GetTrackOk

`func (o *GetActionLogResponse) GetTrackOk() (*string, bool)`

GetTrackOk returns a tuple with the Track field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrack

`func (o *GetActionLogResponse) SetTrack(v string)`

SetTrack sets Track field to given value.

### HasTrack

`func (o *GetActionLogResponse) HasTrack() bool`

HasTrack returns a boolean if a field has been set.

### SetTrackNil

`func (o *GetActionLogResponse) SetTrackNil(b bool)`

 SetTrackNil sets the value for Track to be an explicit nil

### UnsetTrack
`func (o *GetActionLogResponse) UnsetTrack()`

UnsetTrack ensures that no value is present for Track, not even an explicit nil
### GetTargetType

`func (o *GetActionLogResponse) GetTargetType() string`

GetTargetType returns the TargetType field if non-nil, zero value otherwise.

### GetTargetTypeOk

`func (o *GetActionLogResponse) GetTargetTypeOk() (*string, bool)`

GetTargetTypeOk returns a tuple with the TargetType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetType

`func (o *GetActionLogResponse) SetTargetType(v string)`

SetTargetType sets TargetType field to given value.

### HasTargetType

`func (o *GetActionLogResponse) HasTargetType() bool`

HasTargetType returns a boolean if a field has been set.

### SetTargetTypeNil

`func (o *GetActionLogResponse) SetTargetTypeNil(b bool)`

 SetTargetTypeNil sets the value for TargetType to be an explicit nil

### UnsetTargetType
`func (o *GetActionLogResponse) UnsetTargetType()`

UnsetTargetType ensures that no value is present for TargetType, not even an explicit nil
### GetTargetId

`func (o *GetActionLogResponse) GetTargetId() string`

GetTargetId returns the TargetId field if non-nil, zero value otherwise.

### GetTargetIdOk

`func (o *GetActionLogResponse) GetTargetIdOk() (*string, bool)`

GetTargetIdOk returns a tuple with the TargetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetId

`func (o *GetActionLogResponse) SetTargetId(v string)`

SetTargetId sets TargetId field to given value.

### HasTargetId

`func (o *GetActionLogResponse) HasTargetId() bool`

HasTargetId returns a boolean if a field has been set.

### SetTargetIdNil

`func (o *GetActionLogResponse) SetTargetIdNil(b bool)`

 SetTargetIdNil sets the value for TargetId to be an explicit nil

### UnsetTargetId
`func (o *GetActionLogResponse) UnsetTargetId()`

UnsetTargetId ensures that no value is present for TargetId, not even an explicit nil
### GetHttpStatus

`func (o *GetActionLogResponse) GetHttpStatus() float32`

GetHttpStatus returns the HttpStatus field if non-nil, zero value otherwise.

### GetHttpStatusOk

`func (o *GetActionLogResponse) GetHttpStatusOk() (*float32, bool)`

GetHttpStatusOk returns a tuple with the HttpStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpStatus

`func (o *GetActionLogResponse) SetHttpStatus(v float32)`

SetHttpStatus sets HttpStatus field to given value.

### HasHttpStatus

`func (o *GetActionLogResponse) HasHttpStatus() bool`

HasHttpStatus returns a boolean if a field has been set.

### SetHttpStatusNil

`func (o *GetActionLogResponse) SetHttpStatusNil(b bool)`

 SetHttpStatusNil sets the value for HttpStatus to be an explicit nil

### UnsetHttpStatus
`func (o *GetActionLogResponse) UnsetHttpStatus()`

UnsetHttpStatus ensures that no value is present for HttpStatus, not even an explicit nil
### GetActorEmail

`func (o *GetActionLogResponse) GetActorEmail() string`

GetActorEmail returns the ActorEmail field if non-nil, zero value otherwise.

### GetActorEmailOk

`func (o *GetActionLogResponse) GetActorEmailOk() (*string, bool)`

GetActorEmailOk returns a tuple with the ActorEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorEmail

`func (o *GetActionLogResponse) SetActorEmail(v string)`

SetActorEmail sets ActorEmail field to given value.

### HasActorEmail

`func (o *GetActionLogResponse) HasActorEmail() bool`

HasActorEmail returns a boolean if a field has been set.

### SetActorEmailNil

`func (o *GetActionLogResponse) SetActorEmailNil(b bool)`

 SetActorEmailNil sets the value for ActorEmail to be an explicit nil

### UnsetActorEmail
`func (o *GetActionLogResponse) UnsetActorEmail()`

UnsetActorEmail ensures that no value is present for ActorEmail, not even an explicit nil
### GetActorDisplayName

`func (o *GetActionLogResponse) GetActorDisplayName() string`

GetActorDisplayName returns the ActorDisplayName field if non-nil, zero value otherwise.

### GetActorDisplayNameOk

`func (o *GetActionLogResponse) GetActorDisplayNameOk() (*string, bool)`

GetActorDisplayNameOk returns a tuple with the ActorDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActorDisplayName

`func (o *GetActionLogResponse) SetActorDisplayName(v string)`

SetActorDisplayName sets ActorDisplayName field to given value.

### HasActorDisplayName

`func (o *GetActionLogResponse) HasActorDisplayName() bool`

HasActorDisplayName returns a boolean if a field has been set.

### SetActorDisplayNameNil

`func (o *GetActionLogResponse) SetActorDisplayNameNil(b bool)`

 SetActorDisplayNameNil sets the value for ActorDisplayName to be an explicit nil

### UnsetActorDisplayName
`func (o *GetActionLogResponse) UnsetActorDisplayName()`

UnsetActorDisplayName ensures that no value is present for ActorDisplayName, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


