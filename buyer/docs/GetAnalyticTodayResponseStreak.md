# GetAnalyticTodayResponseStreak

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Current** | **float32** |  | 
**Longest** | **float32** |  | 
**LastActiveDate** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetAnalyticTodayResponseStreak

`func NewGetAnalyticTodayResponseStreak(current float32, longest float32, ) *GetAnalyticTodayResponseStreak`

NewGetAnalyticTodayResponseStreak instantiates a new GetAnalyticTodayResponseStreak object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalyticTodayResponseStreakWithDefaults

`func NewGetAnalyticTodayResponseStreakWithDefaults() *GetAnalyticTodayResponseStreak`

NewGetAnalyticTodayResponseStreakWithDefaults instantiates a new GetAnalyticTodayResponseStreak object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrent

`func (o *GetAnalyticTodayResponseStreak) GetCurrent() float32`

GetCurrent returns the Current field if non-nil, zero value otherwise.

### GetCurrentOk

`func (o *GetAnalyticTodayResponseStreak) GetCurrentOk() (*float32, bool)`

GetCurrentOk returns a tuple with the Current field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrent

`func (o *GetAnalyticTodayResponseStreak) SetCurrent(v float32)`

SetCurrent sets Current field to given value.


### GetLongest

`func (o *GetAnalyticTodayResponseStreak) GetLongest() float32`

GetLongest returns the Longest field if non-nil, zero value otherwise.

### GetLongestOk

`func (o *GetAnalyticTodayResponseStreak) GetLongestOk() (*float32, bool)`

GetLongestOk returns a tuple with the Longest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLongest

`func (o *GetAnalyticTodayResponseStreak) SetLongest(v float32)`

SetLongest sets Longest field to given value.


### GetLastActiveDate

`func (o *GetAnalyticTodayResponseStreak) GetLastActiveDate() string`

GetLastActiveDate returns the LastActiveDate field if non-nil, zero value otherwise.

### GetLastActiveDateOk

`func (o *GetAnalyticTodayResponseStreak) GetLastActiveDateOk() (*string, bool)`

GetLastActiveDateOk returns a tuple with the LastActiveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActiveDate

`func (o *GetAnalyticTodayResponseStreak) SetLastActiveDate(v string)`

SetLastActiveDate sets LastActiveDate field to given value.

### HasLastActiveDate

`func (o *GetAnalyticTodayResponseStreak) HasLastActiveDate() bool`

HasLastActiveDate returns a boolean if a field has been set.

### SetLastActiveDateNil

`func (o *GetAnalyticTodayResponseStreak) SetLastActiveDateNil(b bool)`

 SetLastActiveDateNil sets the value for LastActiveDate to be an explicit nil

### UnsetLastActiveDate
`func (o *GetAnalyticTodayResponseStreak) UnsetLastActiveDate()`

UnsetLastActiveDate ensures that no value is present for LastActiveDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


