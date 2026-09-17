# GetOrderEvidenceResponseArrival

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**State** | **string** |  | 
**WindowHours** | **float32** |  | 
**RequestedAt** | Pointer to **NullableTime** |  | [optional] 
**ClosesAt** | Pointer to **NullableTime** | When the window shuts. Null when nothing was ever asked. | [optional] 
**Detail** | **string** |  | 
**Photos** | **[]string** |  | 

## Methods

### NewGetOrderEvidenceResponseArrival

`func NewGetOrderEvidenceResponseArrival(state string, windowHours float32, detail string, photos []string, ) *GetOrderEvidenceResponseArrival`

NewGetOrderEvidenceResponseArrival instantiates a new GetOrderEvidenceResponseArrival object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderEvidenceResponseArrivalWithDefaults

`func NewGetOrderEvidenceResponseArrivalWithDefaults() *GetOrderEvidenceResponseArrival`

NewGetOrderEvidenceResponseArrivalWithDefaults instantiates a new GetOrderEvidenceResponseArrival object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetState

`func (o *GetOrderEvidenceResponseArrival) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetOrderEvidenceResponseArrival) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetOrderEvidenceResponseArrival) SetState(v string)`

SetState sets State field to given value.


### GetWindowHours

`func (o *GetOrderEvidenceResponseArrival) GetWindowHours() float32`

GetWindowHours returns the WindowHours field if non-nil, zero value otherwise.

### GetWindowHoursOk

`func (o *GetOrderEvidenceResponseArrival) GetWindowHoursOk() (*float32, bool)`

GetWindowHoursOk returns a tuple with the WindowHours field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowHours

`func (o *GetOrderEvidenceResponseArrival) SetWindowHours(v float32)`

SetWindowHours sets WindowHours field to given value.


### GetRequestedAt

`func (o *GetOrderEvidenceResponseArrival) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *GetOrderEvidenceResponseArrival) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *GetOrderEvidenceResponseArrival) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *GetOrderEvidenceResponseArrival) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### SetRequestedAtNil

`func (o *GetOrderEvidenceResponseArrival) SetRequestedAtNil(b bool)`

 SetRequestedAtNil sets the value for RequestedAt to be an explicit nil

### UnsetRequestedAt
`func (o *GetOrderEvidenceResponseArrival) UnsetRequestedAt()`

UnsetRequestedAt ensures that no value is present for RequestedAt, not even an explicit nil
### GetClosesAt

`func (o *GetOrderEvidenceResponseArrival) GetClosesAt() time.Time`

GetClosesAt returns the ClosesAt field if non-nil, zero value otherwise.

### GetClosesAtOk

`func (o *GetOrderEvidenceResponseArrival) GetClosesAtOk() (*time.Time, bool)`

GetClosesAtOk returns a tuple with the ClosesAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosesAt

`func (o *GetOrderEvidenceResponseArrival) SetClosesAt(v time.Time)`

SetClosesAt sets ClosesAt field to given value.

### HasClosesAt

`func (o *GetOrderEvidenceResponseArrival) HasClosesAt() bool`

HasClosesAt returns a boolean if a field has been set.

### SetClosesAtNil

`func (o *GetOrderEvidenceResponseArrival) SetClosesAtNil(b bool)`

 SetClosesAtNil sets the value for ClosesAt to be an explicit nil

### UnsetClosesAt
`func (o *GetOrderEvidenceResponseArrival) UnsetClosesAt()`

UnsetClosesAt ensures that no value is present for ClosesAt, not even an explicit nil
### GetDetail

`func (o *GetOrderEvidenceResponseArrival) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *GetOrderEvidenceResponseArrival) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *GetOrderEvidenceResponseArrival) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetPhotos

`func (o *GetOrderEvidenceResponseArrival) GetPhotos() []string`

GetPhotos returns the Photos field if non-nil, zero value otherwise.

### GetPhotosOk

`func (o *GetOrderEvidenceResponseArrival) GetPhotosOk() (*[]string, bool)`

GetPhotosOk returns a tuple with the Photos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhotos

`func (o *GetOrderEvidenceResponseArrival) SetPhotos(v []string)`

SetPhotos sets Photos field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


