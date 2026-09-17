# GetOrderEvidenceResponseGrade

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Grade** | **string** |  | 
**Points** | **float32** |  | 
**Signals** | [**[]GetOrderEvidenceResponseGradeSignals**](GetOrderEvidenceResponseGradeSignals.md) |  | 
**NextStep** | Pointer to **NullableString** | The single most valuable thing not yet done, phrased as an action.  One, not a checklist. A seller given six things to fix does none of them; a seller given the one worth the most does that one. | [optional] 

## Methods

### NewGetOrderEvidenceResponseGrade

`func NewGetOrderEvidenceResponseGrade(grade string, points float32, signals []GetOrderEvidenceResponseGradeSignals, ) *GetOrderEvidenceResponseGrade`

NewGetOrderEvidenceResponseGrade instantiates a new GetOrderEvidenceResponseGrade object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderEvidenceResponseGradeWithDefaults

`func NewGetOrderEvidenceResponseGradeWithDefaults() *GetOrderEvidenceResponseGrade`

NewGetOrderEvidenceResponseGradeWithDefaults instantiates a new GetOrderEvidenceResponseGrade object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrade

`func (o *GetOrderEvidenceResponseGrade) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *GetOrderEvidenceResponseGrade) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *GetOrderEvidenceResponseGrade) SetGrade(v string)`

SetGrade sets Grade field to given value.


### GetPoints

`func (o *GetOrderEvidenceResponseGrade) GetPoints() float32`

GetPoints returns the Points field if non-nil, zero value otherwise.

### GetPointsOk

`func (o *GetOrderEvidenceResponseGrade) GetPointsOk() (*float32, bool)`

GetPointsOk returns a tuple with the Points field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPoints

`func (o *GetOrderEvidenceResponseGrade) SetPoints(v float32)`

SetPoints sets Points field to given value.


### GetSignals

`func (o *GetOrderEvidenceResponseGrade) GetSignals() []GetOrderEvidenceResponseGradeSignals`

GetSignals returns the Signals field if non-nil, zero value otherwise.

### GetSignalsOk

`func (o *GetOrderEvidenceResponseGrade) GetSignalsOk() (*[]GetOrderEvidenceResponseGradeSignals, bool)`

GetSignalsOk returns a tuple with the Signals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignals

`func (o *GetOrderEvidenceResponseGrade) SetSignals(v []GetOrderEvidenceResponseGradeSignals)`

SetSignals sets Signals field to given value.


### GetNextStep

`func (o *GetOrderEvidenceResponseGrade) GetNextStep() string`

GetNextStep returns the NextStep field if non-nil, zero value otherwise.

### GetNextStepOk

`func (o *GetOrderEvidenceResponseGrade) GetNextStepOk() (*string, bool)`

GetNextStepOk returns a tuple with the NextStep field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextStep

`func (o *GetOrderEvidenceResponseGrade) SetNextStep(v string)`

SetNextStep sets NextStep field to given value.

### HasNextStep

`func (o *GetOrderEvidenceResponseGrade) HasNextStep() bool`

HasNextStep returns a boolean if a field has been set.

### SetNextStepNil

`func (o *GetOrderEvidenceResponseGrade) SetNextStepNil(b bool)`

 SetNextStepNil sets the value for NextStep to be an explicit nil

### UnsetNextStep
`func (o *GetOrderEvidenceResponseGrade) UnsetNextStep()`

UnsetNextStep ensures that no value is present for NextStep, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


