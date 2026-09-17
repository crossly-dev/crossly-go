# ListBuyerScanSessionsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Device** | **string** |  | 
**Label** | Pointer to **NullableString** |  | [optional] 
**StartedAt** | **string** |  | 
**EndedAt** | Pointer to **NullableString** |  | [optional] 
**CaptureCount** | **float32** |  | 
**Live** | **bool** |  | 

## Methods

### NewListBuyerScanSessionsItem

`func NewListBuyerScanSessionsItem(id string, device string, startedAt string, captureCount float32, live bool, ) *ListBuyerScanSessionsItem`

NewListBuyerScanSessionsItem instantiates a new ListBuyerScanSessionsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBuyerScanSessionsItemWithDefaults

`func NewListBuyerScanSessionsItemWithDefaults() *ListBuyerScanSessionsItem`

NewListBuyerScanSessionsItemWithDefaults instantiates a new ListBuyerScanSessionsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListBuyerScanSessionsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBuyerScanSessionsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBuyerScanSessionsItem) SetId(v string)`

SetId sets Id field to given value.


### GetDevice

`func (o *ListBuyerScanSessionsItem) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *ListBuyerScanSessionsItem) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *ListBuyerScanSessionsItem) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetLabel

`func (o *ListBuyerScanSessionsItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ListBuyerScanSessionsItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ListBuyerScanSessionsItem) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ListBuyerScanSessionsItem) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *ListBuyerScanSessionsItem) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *ListBuyerScanSessionsItem) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetStartedAt

`func (o *ListBuyerScanSessionsItem) GetStartedAt() string`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ListBuyerScanSessionsItem) GetStartedAtOk() (*string, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ListBuyerScanSessionsItem) SetStartedAt(v string)`

SetStartedAt sets StartedAt field to given value.


### GetEndedAt

`func (o *ListBuyerScanSessionsItem) GetEndedAt() string`

GetEndedAt returns the EndedAt field if non-nil, zero value otherwise.

### GetEndedAtOk

`func (o *ListBuyerScanSessionsItem) GetEndedAtOk() (*string, bool)`

GetEndedAtOk returns a tuple with the EndedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndedAt

`func (o *ListBuyerScanSessionsItem) SetEndedAt(v string)`

SetEndedAt sets EndedAt field to given value.

### HasEndedAt

`func (o *ListBuyerScanSessionsItem) HasEndedAt() bool`

HasEndedAt returns a boolean if a field has been set.

### SetEndedAtNil

`func (o *ListBuyerScanSessionsItem) SetEndedAtNil(b bool)`

 SetEndedAtNil sets the value for EndedAt to be an explicit nil

### UnsetEndedAt
`func (o *ListBuyerScanSessionsItem) UnsetEndedAt()`

UnsetEndedAt ensures that no value is present for EndedAt, not even an explicit nil
### GetCaptureCount

`func (o *ListBuyerScanSessionsItem) GetCaptureCount() float32`

GetCaptureCount returns the CaptureCount field if non-nil, zero value otherwise.

### GetCaptureCountOk

`func (o *ListBuyerScanSessionsItem) GetCaptureCountOk() (*float32, bool)`

GetCaptureCountOk returns a tuple with the CaptureCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptureCount

`func (o *ListBuyerScanSessionsItem) SetCaptureCount(v float32)`

SetCaptureCount sets CaptureCount field to given value.


### GetLive

`func (o *ListBuyerScanSessionsItem) GetLive() bool`

GetLive returns the Live field if non-nil, zero value otherwise.

### GetLiveOk

`func (o *ListBuyerScanSessionsItem) GetLiveOk() (*bool, bool)`

GetLiveOk returns a tuple with the Live field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLive

`func (o *ListBuyerScanSessionsItem) SetLive(v bool)`

SetLive sets Live field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


