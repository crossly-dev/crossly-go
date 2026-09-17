# ListNotificationIntegrationsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | **map[string]interface{}** |  | 
**Id** | **string** |  | 
**Provider** | **string** |  | 
**Name** | **string** |  | 
**Enabled** | **bool** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewListNotificationIntegrationsItem

`func NewListNotificationIntegrationsItem(config map[string]interface{}, id string, provider string, name string, enabled bool, createdAt time.Time, ) *ListNotificationIntegrationsItem`

NewListNotificationIntegrationsItem instantiates a new ListNotificationIntegrationsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNotificationIntegrationsItemWithDefaults

`func NewListNotificationIntegrationsItemWithDefaults() *ListNotificationIntegrationsItem`

NewListNotificationIntegrationsItemWithDefaults instantiates a new ListNotificationIntegrationsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ListNotificationIntegrationsItem) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListNotificationIntegrationsItem) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListNotificationIntegrationsItem) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.


### GetId

`func (o *ListNotificationIntegrationsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListNotificationIntegrationsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListNotificationIntegrationsItem) SetId(v string)`

SetId sets Id field to given value.


### GetProvider

`func (o *ListNotificationIntegrationsItem) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ListNotificationIntegrationsItem) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ListNotificationIntegrationsItem) SetProvider(v string)`

SetProvider sets Provider field to given value.


### GetName

`func (o *ListNotificationIntegrationsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListNotificationIntegrationsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListNotificationIntegrationsItem) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *ListNotificationIntegrationsItem) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListNotificationIntegrationsItem) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListNotificationIntegrationsItem) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetCreatedAt

`func (o *ListNotificationIntegrationsItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListNotificationIntegrationsItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListNotificationIntegrationsItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


