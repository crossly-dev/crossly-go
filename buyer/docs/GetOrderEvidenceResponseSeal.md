# GetOrderEvidenceResponseSeal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Assessment** | [**GetOrderEvidenceResponseSealAssessment**](GetOrderEvidenceResponseSealAssessment.md) |  | 
**DispatchSearchedFrames** | Pointer to **NullableFloat32** | How many packing frames were searched to find the dispatch frame shot closest to the arrival angle. Null when the dispatch reading was simply the frame the recorder ended on.  Surfaced rather than kept internal: \&quot;the best of eighteen frames matched\&quot; is a weaker claim than \&quot;the frame we took matched\&quot;, and a reviewer has to be able to tell them apart. See frame-match.ts. | [optional] 
**HasDispatch** | **bool** |  | 
**HasArrival** | **bool** |  | 
**HasCourierPhoto** | **bool** | A courier photo exists, even if the seal was not legible in it. | 

## Methods

### NewGetOrderEvidenceResponseSeal

`func NewGetOrderEvidenceResponseSeal(assessment GetOrderEvidenceResponseSealAssessment, hasDispatch bool, hasArrival bool, hasCourierPhoto bool, ) *GetOrderEvidenceResponseSeal`

NewGetOrderEvidenceResponseSeal instantiates a new GetOrderEvidenceResponseSeal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderEvidenceResponseSealWithDefaults

`func NewGetOrderEvidenceResponseSealWithDefaults() *GetOrderEvidenceResponseSeal`

NewGetOrderEvidenceResponseSealWithDefaults instantiates a new GetOrderEvidenceResponseSeal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssessment

`func (o *GetOrderEvidenceResponseSeal) GetAssessment() GetOrderEvidenceResponseSealAssessment`

GetAssessment returns the Assessment field if non-nil, zero value otherwise.

### GetAssessmentOk

`func (o *GetOrderEvidenceResponseSeal) GetAssessmentOk() (*GetOrderEvidenceResponseSealAssessment, bool)`

GetAssessmentOk returns a tuple with the Assessment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssessment

`func (o *GetOrderEvidenceResponseSeal) SetAssessment(v GetOrderEvidenceResponseSealAssessment)`

SetAssessment sets Assessment field to given value.


### GetDispatchSearchedFrames

`func (o *GetOrderEvidenceResponseSeal) GetDispatchSearchedFrames() float32`

GetDispatchSearchedFrames returns the DispatchSearchedFrames field if non-nil, zero value otherwise.

### GetDispatchSearchedFramesOk

`func (o *GetOrderEvidenceResponseSeal) GetDispatchSearchedFramesOk() (*float32, bool)`

GetDispatchSearchedFramesOk returns a tuple with the DispatchSearchedFrames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDispatchSearchedFrames

`func (o *GetOrderEvidenceResponseSeal) SetDispatchSearchedFrames(v float32)`

SetDispatchSearchedFrames sets DispatchSearchedFrames field to given value.

### HasDispatchSearchedFrames

`func (o *GetOrderEvidenceResponseSeal) HasDispatchSearchedFrames() bool`

HasDispatchSearchedFrames returns a boolean if a field has been set.

### SetDispatchSearchedFramesNil

`func (o *GetOrderEvidenceResponseSeal) SetDispatchSearchedFramesNil(b bool)`

 SetDispatchSearchedFramesNil sets the value for DispatchSearchedFrames to be an explicit nil

### UnsetDispatchSearchedFrames
`func (o *GetOrderEvidenceResponseSeal) UnsetDispatchSearchedFrames()`

UnsetDispatchSearchedFrames ensures that no value is present for DispatchSearchedFrames, not even an explicit nil
### GetHasDispatch

`func (o *GetOrderEvidenceResponseSeal) GetHasDispatch() bool`

GetHasDispatch returns the HasDispatch field if non-nil, zero value otherwise.

### GetHasDispatchOk

`func (o *GetOrderEvidenceResponseSeal) GetHasDispatchOk() (*bool, bool)`

GetHasDispatchOk returns a tuple with the HasDispatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDispatch

`func (o *GetOrderEvidenceResponseSeal) SetHasDispatch(v bool)`

SetHasDispatch sets HasDispatch field to given value.


### GetHasArrival

`func (o *GetOrderEvidenceResponseSeal) GetHasArrival() bool`

GetHasArrival returns the HasArrival field if non-nil, zero value otherwise.

### GetHasArrivalOk

`func (o *GetOrderEvidenceResponseSeal) GetHasArrivalOk() (*bool, bool)`

GetHasArrivalOk returns a tuple with the HasArrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasArrival

`func (o *GetOrderEvidenceResponseSeal) SetHasArrival(v bool)`

SetHasArrival sets HasArrival field to given value.


### GetHasCourierPhoto

`func (o *GetOrderEvidenceResponseSeal) GetHasCourierPhoto() bool`

GetHasCourierPhoto returns the HasCourierPhoto field if non-nil, zero value otherwise.

### GetHasCourierPhotoOk

`func (o *GetOrderEvidenceResponseSeal) GetHasCourierPhotoOk() (*bool, bool)`

GetHasCourierPhotoOk returns a tuple with the HasCourierPhoto field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasCourierPhoto

`func (o *GetOrderEvidenceResponseSeal) SetHasCourierPhoto(v bool)`

SetHasCourierPhoto sets HasCourierPhoto field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


