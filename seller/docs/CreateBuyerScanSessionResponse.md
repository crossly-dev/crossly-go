# CreateBuyerScanSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Device** | **string** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**StartedAt** | **string** |  | 
**CaptureCount** | **float32** |  | 

## Methods

### NewCreateBuyerScanSessionResponse

`func NewCreateBuyerScanSessionResponse(id string, device string, startedAt string, captureCount float32, ) *CreateBuyerScanSessionResponse`

NewCreateBuyerScanSessionResponse instantiates a new CreateBuyerScanSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerScanSessionResponseWithDefaults

`func NewCreateBuyerScanSessionResponseWithDefaults() *CreateBuyerScanSessionResponse`

NewCreateBuyerScanSessionResponseWithDefaults instantiates a new CreateBuyerScanSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateBuyerScanSessionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateBuyerScanSessionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateBuyerScanSessionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetDevice

`func (o *CreateBuyerScanSessionResponse) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CreateBuyerScanSessionResponse) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CreateBuyerScanSessionResponse) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetLabel

`func (o *CreateBuyerScanSessionResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateBuyerScanSessionResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateBuyerScanSessionResponse) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateBuyerScanSessionResponse) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *CreateBuyerScanSessionResponse) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *CreateBuyerScanSessionResponse) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetStartedAt

`func (o *CreateBuyerScanSessionResponse) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CreateBuyerScanSessionResponse) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CreateBuyerScanSessionResponse) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### GetCaptureCount

`func (o *CreateBuyerScanSessionResponse) GetCaptureCount() float32`

GetCaptureCount returns the CaptureCount field if non-nil, zero value otherwise.

### GetCaptureCountOk

`func (o *CreateBuyerScanSessionResponse) GetCaptureCountOk() (*float32, bool)`

GetCaptureCountOk returns a tuple with the CaptureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureCount

`func (o *CreateBuyerScanSessionResponse) SetCaptureCount(v float32)`

SetCaptureCount sets CaptureCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


