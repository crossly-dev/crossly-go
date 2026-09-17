# GetAnalyticTodayResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checklist** | [**GetAnalyticTodayResponseChecklist**](GetAnalyticTodayResponseChecklist.md) |  | 
**Streak** | [**GetAnalyticTodayResponseStreak**](GetAnalyticTodayResponseStreak.md) |  | 
**ActivitySparkline** | **[]float32** |  | 

## Methods

### NewGetAnalyticTodayResponse

`func NewGetAnalyticTodayResponse(checklist GetAnalyticTodayResponseChecklist, streak GetAnalyticTodayResponseStreak, activitySparkline []float32, ) *GetAnalyticTodayResponse`

NewGetAnalyticTodayResponse instantiates a new GetAnalyticTodayResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalyticTodayResponseWithDefaults

`func NewGetAnalyticTodayResponseWithDefaults() *GetAnalyticTodayResponse`

NewGetAnalyticTodayResponseWithDefaults instantiates a new GetAnalyticTodayResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecklist

`func (o *GetAnalyticTodayResponse) GetChecklist() GetAnalyticTodayResponseChecklist`

GetChecklist returns the Checklist field if non-nil, zero value otherwise.

### GetChecklistOk

`func (o *GetAnalyticTodayResponse) GetChecklistOk() (*GetAnalyticTodayResponseChecklist, bool)`

GetChecklistOk returns a tuple with the Checklist field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecklist

`func (o *GetAnalyticTodayResponse) SetChecklist(v GetAnalyticTodayResponseChecklist)`

SetChecklist sets Checklist field to given value.


### GetStreak

`func (o *GetAnalyticTodayResponse) GetStreak() GetAnalyticTodayResponseStreak`

GetStreak returns the Streak field if non-nil, zero value otherwise.

### GetStreakOk

`func (o *GetAnalyticTodayResponse) GetStreakOk() (*GetAnalyticTodayResponseStreak, bool)`

GetStreakOk returns a tuple with the Streak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreak

`func (o *GetAnalyticTodayResponse) SetStreak(v GetAnalyticTodayResponseStreak)`

SetStreak sets Streak field to given value.


### GetActivitySparkline

`func (o *GetAnalyticTodayResponse) GetActivitySparkline() []float32`

GetActivitySparkline returns the ActivitySparkline field if non-nil, zero value otherwise.

### GetActivitySparklineOk

`func (o *GetAnalyticTodayResponse) GetActivitySparklineOk() (*[]float32, bool)`

GetActivitySparklineOk returns a tuple with the ActivitySparkline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivitySparkline

`func (o *GetAnalyticTodayResponse) SetActivitySparkline(v []float32)`

SetActivitySparkline sets ActivitySparkline field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


