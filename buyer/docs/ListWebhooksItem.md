# ListWebhooksItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Url** | **string** |  | 
**Events** | **[]string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Enabled** | **bool** |  | 
**LastDeliveryAt** | Pointer to **NullableTime** |  | [optional] 
**LastErrorAt** | Pointer to **NullableTime** |  | [optional] 
**LastError** | Pointer to **NullableString** |  | [optional] 
**ConsecutiveFailures** | **float32** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewListWebhooksItem

`func NewListWebhooksItem(id string, url string, events []string, enabled bool, consecutiveFailures float32, createdAt time.Time, ) *ListWebhooksItem`

NewListWebhooksItem instantiates a new ListWebhooksItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListWebhooksItemWithDefaults

`func NewListWebhooksItemWithDefaults() *ListWebhooksItem`

NewListWebhooksItemWithDefaults instantiates a new ListWebhooksItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListWebhooksItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListWebhooksItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListWebhooksItem) SetId(v string)`

SetId sets Id field to given value.


### GetUrl

`func (o *ListWebhooksItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ListWebhooksItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ListWebhooksItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetEvents

`func (o *ListWebhooksItem) GetEvents() []string`

GetEvents returns the Events field if non-nil, zero value otherwise.

### GetEventsOk

`func (o *ListWebhooksItem) GetEventsOk() (*[]string, bool)`

GetEventsOk returns a tuple with the Events field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvents

`func (o *ListWebhooksItem) SetEvents(v []string)`

SetEvents sets Events field to given value.


### GetDescription

`func (o *ListWebhooksItem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListWebhooksItem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListWebhooksItem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListWebhooksItem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListWebhooksItem) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListWebhooksItem) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetEnabled

`func (o *ListWebhooksItem) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListWebhooksItem) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListWebhooksItem) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLastDeliveryAt

`func (o *ListWebhooksItem) GetLastDeliveryAt() time.Time`

GetLastDeliveryAt returns the LastDeliveryAt field if non-nil, zero value otherwise.

### GetLastDeliveryAtOk

`func (o *ListWebhooksItem) GetLastDeliveryAtOk() (*time.Time, bool)`

GetLastDeliveryAtOk returns a tuple with the LastDeliveryAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeliveryAt

`func (o *ListWebhooksItem) SetLastDeliveryAt(v time.Time)`

SetLastDeliveryAt sets LastDeliveryAt field to given value.

### HasLastDeliveryAt

`func (o *ListWebhooksItem) HasLastDeliveryAt() bool`

HasLastDeliveryAt returns a boolean if a field has been set.

### SetLastDeliveryAtNil

`func (o *ListWebhooksItem) SetLastDeliveryAtNil(b bool)`

 SetLastDeliveryAtNil sets the value for LastDeliveryAt to be an explicit nil

### UnsetLastDeliveryAt
`func (o *ListWebhooksItem) UnsetLastDeliveryAt()`

UnsetLastDeliveryAt ensures that no value is present for LastDeliveryAt, not even an explicit nil
### GetLastErrorAt

`func (o *ListWebhooksItem) GetLastErrorAt() time.Time`

GetLastErrorAt returns the LastErrorAt field if non-nil, zero value otherwise.

### GetLastErrorAtOk

`func (o *ListWebhooksItem) GetLastErrorAtOk() (*time.Time, bool)`

GetLastErrorAtOk returns a tuple with the LastErrorAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastErrorAt

`func (o *ListWebhooksItem) SetLastErrorAt(v time.Time)`

SetLastErrorAt sets LastErrorAt field to given value.

### HasLastErrorAt

`func (o *ListWebhooksItem) HasLastErrorAt() bool`

HasLastErrorAt returns a boolean if a field has been set.

### SetLastErrorAtNil

`func (o *ListWebhooksItem) SetLastErrorAtNil(b bool)`

 SetLastErrorAtNil sets the value for LastErrorAt to be an explicit nil

### UnsetLastErrorAt
`func (o *ListWebhooksItem) UnsetLastErrorAt()`

UnsetLastErrorAt ensures that no value is present for LastErrorAt, not even an explicit nil
### GetLastError

`func (o *ListWebhooksItem) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ListWebhooksItem) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ListWebhooksItem) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ListWebhooksItem) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *ListWebhooksItem) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *ListWebhooksItem) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetConsecutiveFailures

`func (o *ListWebhooksItem) GetConsecutiveFailures() float32`

GetConsecutiveFailures returns the ConsecutiveFailures field if non-nil, zero value otherwise.

### GetConsecutiveFailuresOk

`func (o *ListWebhooksItem) GetConsecutiveFailuresOk() (*float32, bool)`

GetConsecutiveFailuresOk returns a tuple with the ConsecutiveFailures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsecutiveFailures

`func (o *ListWebhooksItem) SetConsecutiveFailures(v float32)`

SetConsecutiveFailures sets ConsecutiveFailures field to given value.


### GetCreatedAt

`func (o *ListWebhooksItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListWebhooksItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListWebhooksItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


