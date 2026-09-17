# GetAiStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HasKey** | **bool** |  | 
**AiEnabled** | **bool** |  | 
**ActiveProvider** | **string** |  | 
**Providers** | **map[string]interface{}** |  | 
**Capabilities** | [**GetAiStatusResponseCapabilities**](GetAiStatusResponseCapabilities.md) |  | 
**CustomConfig** | Pointer to [**NullableGetAiStatusResponseCustomConfig**](GetAiStatusResponseCustomConfig.md) |  | [optional] 

## Methods

### NewGetAiStatusResponse

`func NewGetAiStatusResponse(hasKey bool, aiEnabled bool, activeProvider string, providers map[string]interface{}, capabilities GetAiStatusResponseCapabilities, ) *GetAiStatusResponse`

NewGetAiStatusResponse instantiates a new GetAiStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAiStatusResponseWithDefaults

`func NewGetAiStatusResponseWithDefaults() *GetAiStatusResponse`

NewGetAiStatusResponseWithDefaults instantiates a new GetAiStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHasKey

`func (o *GetAiStatusResponse) GetHasKey() bool`

GetHasKey returns the HasKey field if non-nil, zero value otherwise.

### GetHasKeyOk

`func (o *GetAiStatusResponse) GetHasKeyOk() (*bool, bool)`

GetHasKeyOk returns a tuple with the HasKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasKey

`func (o *GetAiStatusResponse) SetHasKey(v bool)`

SetHasKey sets HasKey field to given value.


### GetAiEnabled

`func (o *GetAiStatusResponse) GetAiEnabled() bool`

GetAiEnabled returns the AiEnabled field if non-nil, zero value otherwise.

### GetAiEnabledOk

`func (o *GetAiStatusResponse) GetAiEnabledOk() (*bool, bool)`

GetAiEnabledOk returns a tuple with the AiEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiEnabled

`func (o *GetAiStatusResponse) SetAiEnabled(v bool)`

SetAiEnabled sets AiEnabled field to given value.


### GetActiveProvider

`func (o *GetAiStatusResponse) GetActiveProvider() string`

GetActiveProvider returns the ActiveProvider field if non-nil, zero value otherwise.

### GetActiveProviderOk

`func (o *GetAiStatusResponse) GetActiveProviderOk() (*string, bool)`

GetActiveProviderOk returns a tuple with the ActiveProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveProvider

`func (o *GetAiStatusResponse) SetActiveProvider(v string)`

SetActiveProvider sets ActiveProvider field to given value.


### GetProviders

`func (o *GetAiStatusResponse) GetProviders() map[string]interface{}`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *GetAiStatusResponse) GetProvidersOk() (*map[string]interface{}, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *GetAiStatusResponse) SetProviders(v map[string]interface{})`

SetProviders sets Providers field to given value.


### GetCapabilities

`func (o *GetAiStatusResponse) GetCapabilities() GetAiStatusResponseCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *GetAiStatusResponse) GetCapabilitiesOk() (*GetAiStatusResponseCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *GetAiStatusResponse) SetCapabilities(v GetAiStatusResponseCapabilities)`

SetCapabilities sets Capabilities field to given value.


### GetCustomConfig

`func (o *GetAiStatusResponse) GetCustomConfig() GetAiStatusResponseCustomConfig`

GetCustomConfig returns the CustomConfig field if non-nil, zero value otherwise.

### GetCustomConfigOk

`func (o *GetAiStatusResponse) GetCustomConfigOk() (*GetAiStatusResponseCustomConfig, bool)`

GetCustomConfigOk returns a tuple with the CustomConfig field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomConfig

`func (o *GetAiStatusResponse) SetCustomConfig(v GetAiStatusResponseCustomConfig)`

SetCustomConfig sets CustomConfig field to given value.

### HasCustomConfig

`func (o *GetAiStatusResponse) HasCustomConfig() bool`

HasCustomConfig returns a boolean if a field has been set.

### SetCustomConfigNil

`func (o *GetAiStatusResponse) SetCustomConfigNil(b bool)`

 SetCustomConfigNil sets the value for CustomConfig to be an explicit nil

### UnsetCustomConfig
`func (o *GetAiStatusResponse) UnsetCustomConfig()`

UnsetCustomConfig ensures that no value is present for CustomConfig, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


