# GetBuyerScanSessionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Device** | **string** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**StartedAt** | **string** |  | 
**EndedAt** | Pointer to **NullableString** |  | [optional] 
**Live** | **bool** |  | 
**CaptureCount** | **float32** |  | 
**SavedCents** | **float32** | Sum of measured savings. Unmeasured captures contribute 0, not null. | 
**Captures** | [**[]GetBuyerScanSessionResponseCaptures**](GetBuyerScanSessionResponseCaptures.md) |  | 

## Methods

### NewGetBuyerScanSessionResponse

`func NewGetBuyerScanSessionResponse(id string, device string, startedAt string, live bool, captureCount float32, savedCents float32, captures []GetBuyerScanSessionResponseCaptures, ) *GetBuyerScanSessionResponse`

NewGetBuyerScanSessionResponse instantiates a new GetBuyerScanSessionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBuyerScanSessionResponseWithDefaults

`func NewGetBuyerScanSessionResponseWithDefaults() *GetBuyerScanSessionResponse`

NewGetBuyerScanSessionResponseWithDefaults instantiates a new GetBuyerScanSessionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetBuyerScanSessionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBuyerScanSessionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBuyerScanSessionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetDevice

`func (o *GetBuyerScanSessionResponse) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *GetBuyerScanSessionResponse) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *GetBuyerScanSessionResponse) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetLabel

`func (o *GetBuyerScanSessionResponse) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetBuyerScanSessionResponse) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetBuyerScanSessionResponse) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetBuyerScanSessionResponse) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *GetBuyerScanSessionResponse) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *GetBuyerScanSessionResponse) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetStartedAt

`func (o *GetBuyerScanSessionResponse) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetBuyerScanSessionResponse) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetBuyerScanSessionResponse) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### GetEndedAt

`func (o *GetBuyerScanSessionResponse) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *GetBuyerScanSessionResponse) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *GetBuyerScanSessionResponse) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *GetBuyerScanSessionResponse) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *GetBuyerScanSessionResponse) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *GetBuyerScanSessionResponse) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetLive

`func (o *GetBuyerScanSessionResponse) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *GetBuyerScanSessionResponse) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *GetBuyerScanSessionResponse) SetLive(v bool)`

SetLive sets Live field to given value.


### GetCaptureCount

`func (o *GetBuyerScanSessionResponse) GetCaptureCount() float32`

GetCaptureCount returns the CaptureCount field if non-nil, zero value otherwise.

### GetCaptureCountOk

`func (o *GetBuyerScanSessionResponse) GetCaptureCountOk() (*float32, bool)`

GetCaptureCountOk returns a tuple with the CaptureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureCount

`func (o *GetBuyerScanSessionResponse) SetCaptureCount(v float32)`

SetCaptureCount sets CaptureCount field to given value.


### GetSavedCents

`func (o *GetBuyerScanSessionResponse) GetSavedCents() float32`

GetSavedCents returns the SavedCents field if non-nil, zero value otherwise.

### GetSavedCentsOk

`func (o *GetBuyerScanSessionResponse) GetSavedCentsOk() (*float32, bool)`

GetSavedCentsOk returns a tuple with the SavedCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavedCents

`func (o *GetBuyerScanSessionResponse) SetSavedCents(v float32)`

SetSavedCents sets SavedCents field to given value.


### GetCaptures

`func (o *GetBuyerScanSessionResponse) GetCaptures() []GetBuyerScanSessionResponseCaptures`

GetCaptures returns the Captures field if non-nil, zero value otherwise.

### GetCapturesOk

`func (o *GetBuyerScanSessionResponse) GetCapturesOk() (*[]GetBuyerScanSessionResponseCaptures, bool)`

GetCapturesOk returns a tuple with the Captures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptures

`func (o *GetBuyerScanSessionResponse) SetCaptures(v []GetBuyerScanSessionResponseCaptures)`

SetCaptures sets Captures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


