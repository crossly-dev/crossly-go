# GetOrderEvidenceResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grade** | [**GetOrderEvidenceResponseGrade**](GetOrderEvidenceResponseGrade.md) |  | 
**Captures** | [**[]GetOrderEvidenceResponseCaptures**](GetOrderEvidenceResponseCaptures.md) |  | 
**Seal** | [**GetOrderEvidenceResponseSeal**](GetOrderEvidenceResponseSeal.md) |  | 
**Arrival** | [**GetOrderEvidenceResponseArrival**](GetOrderEvidenceResponseArrival.md) |  | 
**Units** | [**[]GetOrderEvidenceResponseUnits**](GetOrderEvidenceResponseUnits.md) |  | 

## Methods

### NewGetOrderEvidenceResponse

`func NewGetOrderEvidenceResponse(grade GetOrderEvidenceResponseGrade, captures []GetOrderEvidenceResponseCaptures, seal GetOrderEvidenceResponseSeal, arrival GetOrderEvidenceResponseArrival, units []GetOrderEvidenceResponseUnits, ) *GetOrderEvidenceResponse`

NewGetOrderEvidenceResponse instantiates a new GetOrderEvidenceResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderEvidenceResponseWithDefaults

`func NewGetOrderEvidenceResponseWithDefaults() *GetOrderEvidenceResponse`

NewGetOrderEvidenceResponseWithDefaults instantiates a new GetOrderEvidenceResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrade

`func (o *GetOrderEvidenceResponse) GetGrade() GetOrderEvidenceResponseGrade`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *GetOrderEvidenceResponse) GetGradeOk() (*GetOrderEvidenceResponseGrade, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *GetOrderEvidenceResponse) SetGrade(v GetOrderEvidenceResponseGrade)`

SetGrade sets Grade field to given value.


### GetCaptures

`func (o *GetOrderEvidenceResponse) GetCaptures() []GetOrderEvidenceResponseCaptures`

GetCaptures returns the Captures field if non-nil, zero value otherwise.

### GetCapturesOk

`func (o *GetOrderEvidenceResponse) GetCapturesOk() (*[]GetOrderEvidenceResponseCaptures, bool)`

GetCapturesOk returns a tuple with the Captures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptures

`func (o *GetOrderEvidenceResponse) SetCaptures(v []GetOrderEvidenceResponseCaptures)`

SetCaptures sets Captures field to given value.


### GetSeal

`func (o *GetOrderEvidenceResponse) GetSeal() GetOrderEvidenceResponseSeal`

GetSeal returns the Seal field if non-nil, zero value otherwise.

### GetSealOk

`func (o *GetOrderEvidenceResponse) GetSealOk() (*GetOrderEvidenceResponseSeal, bool)`

GetSealOk returns a tuple with the Seal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeal

`func (o *GetOrderEvidenceResponse) SetSeal(v GetOrderEvidenceResponseSeal)`

SetSeal sets Seal field to given value.


### GetArrival

`func (o *GetOrderEvidenceResponse) GetArrival() GetOrderEvidenceResponseArrival`

GetArrival returns the Arrival field if non-nil, zero value otherwise.

### GetArrivalOk

`func (o *GetOrderEvidenceResponse) GetArrivalOk() (*GetOrderEvidenceResponseArrival, bool)`

GetArrivalOk returns a tuple with the Arrival field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrival

`func (o *GetOrderEvidenceResponse) SetArrival(v GetOrderEvidenceResponseArrival)`

SetArrival sets Arrival field to given value.


### GetUnits

`func (o *GetOrderEvidenceResponse) GetUnits() []GetOrderEvidenceResponseUnits`

GetUnits returns the Units field if non-nil, zero value otherwise.

### GetUnitsOk

`func (o *GetOrderEvidenceResponse) GetUnitsOk() (*[]GetOrderEvidenceResponseUnits, bool)`

GetUnitsOk returns a tuple with the Units field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnits

`func (o *GetOrderEvidenceResponse) SetUnits(v []GetOrderEvidenceResponseUnits)`

SetUnits sets Units field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


