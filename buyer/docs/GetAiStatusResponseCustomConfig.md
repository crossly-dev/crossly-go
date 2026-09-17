# GetAiStatusResponseCustomConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BaseUrl** | **string** |  | 
**TextModel** | **string** |  | 
**VisionModel** | Pointer to **NullableString** |  | [optional] 
**SupportsVision** | **bool** |  | 
**SupportsJsonMode** | **bool** |  | 

## Methods

### NewGetAiStatusResponseCustomConfig

`func NewGetAiStatusResponseCustomConfig(baseUrl string, textModel string, supportsVision bool, supportsJsonMode bool, ) *GetAiStatusResponseCustomConfig`

NewGetAiStatusResponseCustomConfig instantiates a new GetAiStatusResponseCustomConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAiStatusResponseCustomConfigWithDefaults

`func NewGetAiStatusResponseCustomConfigWithDefaults() *GetAiStatusResponseCustomConfig`

NewGetAiStatusResponseCustomConfigWithDefaults instantiates a new GetAiStatusResponseCustomConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBaseUrl

`func (o *GetAiStatusResponseCustomConfig) GetBaseUrl() string`

GetBaseUrl returns the BaseUrl field if non-nil, zero value otherwise.

### GetBaseUrlOk

`func (o *GetAiStatusResponseCustomConfig) GetBaseUrlOk() (*string, bool)`

GetBaseUrlOk returns a tuple with the BaseUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUrl

`func (o *GetAiStatusResponseCustomConfig) SetBaseUrl(v string)`

SetBaseUrl sets BaseUrl field to given value.


### GetTextModel

`func (o *GetAiStatusResponseCustomConfig) GetTextModel() string`

GetTextModel returns the TextModel field if non-nil, zero value otherwise.

### GetTextModelOk

`func (o *GetAiStatusResponseCustomConfig) GetTextModelOk() (*string, bool)`

GetTextModelOk returns a tuple with the TextModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextModel

`func (o *GetAiStatusResponseCustomConfig) SetTextModel(v string)`

SetTextModel sets TextModel field to given value.


### GetVisionModel

`func (o *GetAiStatusResponseCustomConfig) GetVisionModel() string`

GetVisionModel returns the VisionModel field if non-nil, zero value otherwise.

### GetVisionModelOk

`func (o *GetAiStatusResponseCustomConfig) GetVisionModelOk() (*string, bool)`

GetVisionModelOk returns a tuple with the VisionModel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisionModel

`func (o *GetAiStatusResponseCustomConfig) SetVisionModel(v string)`

SetVisionModel sets VisionModel field to given value.

### HasVisionModel

`func (o *GetAiStatusResponseCustomConfig) HasVisionModel() bool`

HasVisionModel returns a boolean if a field has been set.

### SetVisionModelNil

`func (o *GetAiStatusResponseCustomConfig) SetVisionModelNil(b bool)`

 SetVisionModelNil sets the value for VisionModel to be an explicit nil

### UnsetVisionModel
`func (o *GetAiStatusResponseCustomConfig) UnsetVisionModel()`

UnsetVisionModel ensures that no value is present for VisionModel, not even an explicit nil
### GetSupportsVision

`func (o *GetAiStatusResponseCustomConfig) GetSupportsVision() bool`

GetSupportsVision returns the SupportsVision field if non-nil, zero value otherwise.

### GetSupportsVisionOk

`func (o *GetAiStatusResponseCustomConfig) GetSupportsVisionOk() (*bool, bool)`

GetSupportsVisionOk returns a tuple with the SupportsVision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsVision

`func (o *GetAiStatusResponseCustomConfig) SetSupportsVision(v bool)`

SetSupportsVision sets SupportsVision field to given value.


### GetSupportsJsonMode

`func (o *GetAiStatusResponseCustomConfig) GetSupportsJsonMode() bool`

GetSupportsJsonMode returns the SupportsJsonMode field if non-nil, zero value otherwise.

### GetSupportsJsonModeOk

`func (o *GetAiStatusResponseCustomConfig) GetSupportsJsonModeOk() (*bool, bool)`

GetSupportsJsonModeOk returns a tuple with the SupportsJsonMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsJsonMode

`func (o *GetAiStatusResponseCustomConfig) SetSupportsJsonMode(v bool)`

SetSupportsJsonMode sets SupportsJsonMode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


