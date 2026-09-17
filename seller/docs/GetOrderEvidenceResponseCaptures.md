# GetOrderEvidenceResponseCaptures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Kind** | **string** |  | 
**Status** | **string** |  | 
**Verdict** | Pointer to **NullableString** |  | [optional] 
**VideoUrl** | Pointer to **NullableString** |  | [optional] 
**DurationMs** | Pointer to **NullableFloat32** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewGetOrderEvidenceResponseCaptures

`func NewGetOrderEvidenceResponseCaptures(id string, kind string, status string, createdAt time.Time, ) *GetOrderEvidenceResponseCaptures`

NewGetOrderEvidenceResponseCaptures instantiates a new GetOrderEvidenceResponseCaptures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderEvidenceResponseCapturesWithDefaults

`func NewGetOrderEvidenceResponseCapturesWithDefaults() *GetOrderEvidenceResponseCaptures`

NewGetOrderEvidenceResponseCapturesWithDefaults instantiates a new GetOrderEvidenceResponseCaptures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrderEvidenceResponseCaptures) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrderEvidenceResponseCaptures) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrderEvidenceResponseCaptures) SetId(v string)`

SetId sets Id field to given value.


### GetKind

`func (o *GetOrderEvidenceResponseCaptures) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *GetOrderEvidenceResponseCaptures) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *GetOrderEvidenceResponseCaptures) SetKind(v string)`

SetKind sets Kind field to given value.


### GetStatus

`func (o *GetOrderEvidenceResponseCaptures) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrderEvidenceResponseCaptures) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrderEvidenceResponseCaptures) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetVerdict

`func (o *GetOrderEvidenceResponseCaptures) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *GetOrderEvidenceResponseCaptures) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *GetOrderEvidenceResponseCaptures) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.

### HasVerdict

`func (o *GetOrderEvidenceResponseCaptures) HasVerdict() bool`

HasVerdict returns a boolean if a field has been set.

### SetVerdictNil

`func (o *GetOrderEvidenceResponseCaptures) SetVerdictNil(b bool)`

 SetVerdictNil sets the value for Verdict to be an explicit nil

### UnsetVerdict
`func (o *GetOrderEvidenceResponseCaptures) UnsetVerdict()`

UnsetVerdict ensures that no value is present for Verdict, not even an explicit nil
### GetVideoUrl

`func (o *GetOrderEvidenceResponseCaptures) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *GetOrderEvidenceResponseCaptures) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *GetOrderEvidenceResponseCaptures) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *GetOrderEvidenceResponseCaptures) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### SetVideoUrlNil

`func (o *GetOrderEvidenceResponseCaptures) SetVideoUrlNil(b bool)`

 SetVideoUrlNil sets the value for VideoUrl to be an explicit nil

### UnsetVideoUrl
`func (o *GetOrderEvidenceResponseCaptures) UnsetVideoUrl()`

UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
### GetDurationMs

`func (o *GetOrderEvidenceResponseCaptures) GetDurationMs() float32`

GetDurationMs returns the DurationMs field if non-nil, zero value otherwise.

### GetDurationMsOk

`func (o *GetOrderEvidenceResponseCaptures) GetDurationMsOk() (*float32, bool)`

GetDurationMsOk returns a tuple with the DurationMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationMs

`func (o *GetOrderEvidenceResponseCaptures) SetDurationMs(v float32)`

SetDurationMs sets DurationMs field to given value.

### HasDurationMs

`func (o *GetOrderEvidenceResponseCaptures) HasDurationMs() bool`

HasDurationMs returns a boolean if a field has been set.

### SetDurationMsNil

`func (o *GetOrderEvidenceResponseCaptures) SetDurationMsNil(b bool)`

 SetDurationMsNil sets the value for DurationMs to be an explicit nil

### UnsetDurationMs
`func (o *GetOrderEvidenceResponseCaptures) UnsetDurationMs()`

UnsetDurationMs ensures that no value is present for DurationMs, not even an explicit nil
### GetCreatedAt

`func (o *GetOrderEvidenceResponseCaptures) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOrderEvidenceResponseCaptures) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOrderEvidenceResponseCaptures) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


