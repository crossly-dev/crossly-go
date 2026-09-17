# UpdateBuyerMonitorResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** |  | 
**PausedReason** | Pointer to **NullableString** |  | [optional] 
**MatchCount** | **float32** |  | 
**LastCheckedAt** | Pointer to **NullableString** |  | [optional] 
**LastMatchAt** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **string** |  | 
**WebhookSecret** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**Name** | **string** |  | 
**Kind** | **string** |  | 
**Delivery** | **string** |  | 
**WebhookUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateBuyerMonitorResponse

`func NewUpdateBuyerMonitorResponse(active bool, matchCount float32, createdAt string, id string, name string, kind string, delivery string, ) *UpdateBuyerMonitorResponse`

NewUpdateBuyerMonitorResponse instantiates a new UpdateBuyerMonitorResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateBuyerMonitorResponseWithDefaults

`func NewUpdateBuyerMonitorResponseWithDefaults() *UpdateBuyerMonitorResponse`

NewUpdateBuyerMonitorResponseWithDefaults instantiates a new UpdateBuyerMonitorResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *UpdateBuyerMonitorResponse) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateBuyerMonitorResponse) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateBuyerMonitorResponse) SetActive(v bool)`

SetActive sets Active field to given value.


### GetPausedReason

`func (o *UpdateBuyerMonitorResponse) GetPausedReason() string`

GetPausedReason returns the PausedReason field if non-nil, zero value otherwise.

### GetPausedReasonOk

`func (o *UpdateBuyerMonitorResponse) GetPausedReasonOk() (*string, bool)`

GetPausedReasonOk returns a tuple with the PausedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPausedReason

`func (o *UpdateBuyerMonitorResponse) SetPausedReason(v string)`

SetPausedReason sets PausedReason field to given value.

### HasPausedReason

`func (o *UpdateBuyerMonitorResponse) HasPausedReason() bool`

HasPausedReason returns a boolean if a field has been set.

### SetPausedReasonNil

`func (o *UpdateBuyerMonitorResponse) SetPausedReasonNil(b bool)`

 SetPausedReasonNil sets the value for PausedReason to be an explicit nil

### UnsetPausedReason
`func (o *UpdateBuyerMonitorResponse) UnsetPausedReason()`

UnsetPausedReason ensures that no value is present for PausedReason, not even an explicit nil
### GetMatchCount

`func (o *UpdateBuyerMonitorResponse) GetMatchCount() float32`

GetMatchCount returns the MatchCount field if non-nil, zero value otherwise.

### GetMatchCountOk

`func (o *UpdateBuyerMonitorResponse) GetMatchCountOk() (*float32, bool)`

GetMatchCountOk returns a tuple with the MatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCount

`func (o *UpdateBuyerMonitorResponse) SetMatchCount(v float32)`

SetMatchCount sets MatchCount field to given value.


### GetLastCheckedAt

`func (o *UpdateBuyerMonitorResponse) GetLastCheckedAt() string`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *UpdateBuyerMonitorResponse) GetLastCheckedAtOk() (*string, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *UpdateBuyerMonitorResponse) SetLastCheckedAt(v string)`

SetLastCheckedAt sets LastCheckedAt field to given value.

### HasLastCheckedAt

`func (o *UpdateBuyerMonitorResponse) HasLastCheckedAt() bool`

HasLastCheckedAt returns a boolean if a field has been set.

### SetLastCheckedAtNil

`func (o *UpdateBuyerMonitorResponse) SetLastCheckedAtNil(b bool)`

 SetLastCheckedAtNil sets the value for LastCheckedAt to be an explicit nil

### UnsetLastCheckedAt
`func (o *UpdateBuyerMonitorResponse) UnsetLastCheckedAt()`

UnsetLastCheckedAt ensures that no value is present for LastCheckedAt, not even an explicit nil
### GetLastMatchAt

`func (o *UpdateBuyerMonitorResponse) GetLastMatchAt() string`

GetLastMatchAt returns the LastMatchAt field if non-nil, zero value otherwise.

### GetLastMatchAtOk

`func (o *UpdateBuyerMonitorResponse) GetLastMatchAtOk() (*string, bool)`

GetLastMatchAtOk returns a tuple with the LastMatchAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMatchAt

`func (o *UpdateBuyerMonitorResponse) SetLastMatchAt(v string)`

SetLastMatchAt sets LastMatchAt field to given value.

### HasLastMatchAt

`func (o *UpdateBuyerMonitorResponse) HasLastMatchAt() bool`

HasLastMatchAt returns a boolean if a field has been set.

### SetLastMatchAtNil

`func (o *UpdateBuyerMonitorResponse) SetLastMatchAtNil(b bool)`

 SetLastMatchAtNil sets the value for LastMatchAt to be an explicit nil

### UnsetLastMatchAt
`func (o *UpdateBuyerMonitorResponse) UnsetLastMatchAt()`

UnsetLastMatchAt ensures that no value is present for LastMatchAt, not even an explicit nil
### GetCreatedAt

`func (o *UpdateBuyerMonitorResponse) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateBuyerMonitorResponse) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateBuyerMonitorResponse) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.


### GetWebhookSecret

`func (o *UpdateBuyerMonitorResponse) GetWebhookSecret() string`

GetWebhookSecret returns the WebhookSecret field if non-nil, zero value otherwise.

### GetWebhookSecretOk

`func (o *UpdateBuyerMonitorResponse) GetWebhookSecretOk() (*string, bool)`

GetWebhookSecretOk returns a tuple with the WebhookSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookSecret

`func (o *UpdateBuyerMonitorResponse) SetWebhookSecret(v string)`

SetWebhookSecret sets WebhookSecret field to given value.

### HasWebhookSecret

`func (o *UpdateBuyerMonitorResponse) HasWebhookSecret() bool`

HasWebhookSecret returns a boolean if a field has been set.

### SetWebhookSecretNil

`func (o *UpdateBuyerMonitorResponse) SetWebhookSecretNil(b bool)`

 SetWebhookSecretNil sets the value for WebhookSecret to be an explicit nil

### UnsetWebhookSecret
`func (o *UpdateBuyerMonitorResponse) UnsetWebhookSecret()`

UnsetWebhookSecret ensures that no value is present for WebhookSecret, not even an explicit nil
### GetId

`func (o *UpdateBuyerMonitorResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateBuyerMonitorResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateBuyerMonitorResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *UpdateBuyerMonitorResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateBuyerMonitorResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateBuyerMonitorResponse) SetName(v string)`

SetName sets Name field to given value.


### GetKind

`func (o *UpdateBuyerMonitorResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UpdateBuyerMonitorResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UpdateBuyerMonitorResponse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetDelivery

`func (o *UpdateBuyerMonitorResponse) GetDelivery() string`

GetDelivery returns the Delivery field if non-nil, zero value otherwise.

### GetDeliveryOk

`func (o *UpdateBuyerMonitorResponse) GetDeliveryOk() (*string, bool)`

GetDeliveryOk returns a tuple with the Delivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelivery

`func (o *UpdateBuyerMonitorResponse) SetDelivery(v string)`

SetDelivery sets Delivery field to given value.


### GetWebhookUrl

`func (o *UpdateBuyerMonitorResponse) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *UpdateBuyerMonitorResponse) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *UpdateBuyerMonitorResponse) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *UpdateBuyerMonitorResponse) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### SetWebhookUrlNil

`func (o *UpdateBuyerMonitorResponse) SetWebhookUrlNil(b bool)`

 SetWebhookUrlNil sets the value for WebhookUrl to be an explicit nil

### UnsetWebhookUrl
`func (o *UpdateBuyerMonitorResponse) UnsetWebhookUrl()`

UnsetWebhookUrl ensures that no value is present for WebhookUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


