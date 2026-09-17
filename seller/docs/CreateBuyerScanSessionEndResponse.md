# CreateBuyerScanSessionEndResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Device** | **string** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**StartedAt** | **string** |  | 
**EndedAt** | Pointer to **NullableString** |  | [optional] 
**CaptureCount** | **float32** |  | 

## Methods

### NewCreateBuyerScanSessionEndResponse

`func NewCreateBuyerScanSessionEndResponse(id string, device string, startedAt string, captureCount float32, ) *CreateBuyerScanSessionEndResponse`

NewCreateBuyerScanSessionEndResponse instantiates a new CreateBuyerScanSessionEndResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerScanSessionEndResponseWithDefaults

`func NewCreateBuyerScanSessionEndResponseWithDefaults() *CreateBuyerScanSessionEndResponse`

NewCreateBuyerScanSessionEndResponseWithDefaults instantiates a new CreateBuyerScanSessionEndResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateBuyerScanSessionEndResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateBuyerScanSessionEndResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateBuyerScanSessionEndResponse) SetId(v string)`

SetId sets Id field to given value.


### GetDevice

`func (o *CreateBuyerScanSessionEndResponse) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CreateBuyerScanSessionEndResponse) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CreateBuyerScanSessionEndResponse) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetLabel

`func (o *CreateBuyerScanSessionEndResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *CreateBuyerScanSessionEndResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *CreateBuyerScanSessionEndResponse) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *CreateBuyerScanSessionEndResponse) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *CreateBuyerScanSessionEndResponse) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *CreateBuyerScanSessionEndResponse) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetStartedAt

`func (o *CreateBuyerScanSessionEndResponse) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CreateBuyerScanSessionEndResponse) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CreateBuyerScanSessionEndResponse) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### GetEndedAt

`func (o *CreateBuyerScanSessionEndResponse) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *CreateBuyerScanSessionEndResponse) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *CreateBuyerScanSessionEndResponse) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *CreateBuyerScanSessionEndResponse) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *CreateBuyerScanSessionEndResponse) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *CreateBuyerScanSessionEndResponse) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetCaptureCount

`func (o *CreateBuyerScanSessionEndResponse) GetCaptureCount() float32`

GetCaptureCount returns the CaptureCount field if non-nil, zero value otherwise.

### GetCaptureCountOk

`func (o *CreateBuyerScanSessionEndResponse) GetCaptureCountOk() (*float32, bool)`

GetCaptureCountOk returns a tuple with the CaptureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureCount

`func (o *CreateBuyerScanSessionEndResponse) SetCaptureCount(v float32)`

SetCaptureCount sets CaptureCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


