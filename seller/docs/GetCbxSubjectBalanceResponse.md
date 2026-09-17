# GetCbxSubjectBalanceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SubjectId** | **string** |  | 
**AvailableBaseUnits** | **string** |  | 
**PendingCents** | **float32** |  | 

## Methods

### NewGetCbxSubjectBalanceResponse

`func NewGetCbxSubjectBalanceResponse(subjectId string, availableBaseUnits string, pendingCents float32, ) *GetCbxSubjectBalanceResponse`

NewGetCbxSubjectBalanceResponse instantiates a new GetCbxSubjectBalanceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCbxSubjectBalanceResponseWithDefaults

`func NewGetCbxSubjectBalanceResponseWithDefaults() *GetCbxSubjectBalanceResponse`

NewGetCbxSubjectBalanceResponseWithDefaults instantiates a new GetCbxSubjectBalanceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSubjectId

`func (o *GetCbxSubjectBalanceResponse) GetSubjectId() string`

GetSubjectId returns the SubjectId field if non-nil, zero value otherwise.

### GetSubjectIdOk

`func (o *GetCbxSubjectBalanceResponse) GetSubjectIdOk() (*string, bool)`

GetSubjectIdOk returns a tuple with the SubjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectId

`func (o *GetCbxSubjectBalanceResponse) SetSubjectId(v string)`

SetSubjectId sets SubjectId field to given value.


### GetAvailableBaseUnits

`func (o *GetCbxSubjectBalanceResponse) GetAvailableBaseUnits() string`

GetAvailableBaseUnits returns the AvailableBaseUnits field if non-nil, zero value otherwise.

### GetAvailableBaseUnitsOk

`func (o *GetCbxSubjectBalanceResponse) GetAvailableBaseUnitsOk() (*string, bool)`

GetAvailableBaseUnitsOk returns a tuple with the AvailableBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBaseUnits

`func (o *GetCbxSubjectBalanceResponse) SetAvailableBaseUnits(v string)`

SetAvailableBaseUnits sets AvailableBaseUnits field to given value.


### GetPendingCents

`func (o *GetCbxSubjectBalanceResponse) GetPendingCents() float32`

GetPendingCents returns the PendingCents field if non-nil, zero value otherwise.

### GetPendingCentsOk

`func (o *GetCbxSubjectBalanceResponse) GetPendingCentsOk() (*float32, bool)`

GetPendingCentsOk returns a tuple with the PendingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPendingCents

`func (o *GetCbxSubjectBalanceResponse) SetPendingCents(v float32)`

SetPendingCents sets PendingCents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


