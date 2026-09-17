# GetAnalyticTimeseryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **float32** |  | 
**Series** | [**[]GetAnalyticTimeseryResponseSeries**](GetAnalyticTimeseryResponseSeries.md) |  | 

## Methods

### NewGetAnalyticTimeseryResponse

`func NewGetAnalyticTimeseryResponse(days float32, series []GetAnalyticTimeseryResponseSeries, ) *GetAnalyticTimeseryResponse`

NewGetAnalyticTimeseryResponse instantiates a new GetAnalyticTimeseryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalyticTimeseryResponseWithDefaults

`func NewGetAnalyticTimeseryResponseWithDefaults() *GetAnalyticTimeseryResponse`

NewGetAnalyticTimeseryResponseWithDefaults instantiates a new GetAnalyticTimeseryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *GetAnalyticTimeseryResponse) GetDays() float32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *GetAnalyticTimeseryResponse) GetDaysOk() (*float32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *GetAnalyticTimeseryResponse) SetDays(v float32)`

SetDays sets Days field to given value.


### GetSeries

`func (o *GetAnalyticTimeseryResponse) GetSeries() []GetAnalyticTimeseryResponseSeries`

GetSeries returns the Series field if non-nil, zero value otherwise.

### GetSeriesOk

`func (o *GetAnalyticTimeseryResponse) GetSeriesOk() (*[]GetAnalyticTimeseryResponseSeries, bool)`

GetSeriesOk returns a tuple with the Series field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeries

`func (o *GetAnalyticTimeseryResponse) SetSeries(v []GetAnalyticTimeseryResponseSeries)`

SetSeries sets Series field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


