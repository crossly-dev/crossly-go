# GetAnalyticByPlatformResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Days** | **float32** |  | 
**Breakdown** | [**[]GetAnalyticByPlatformResponseBreakdown**](GetAnalyticByPlatformResponseBreakdown.md) |  | 

## Methods

### NewGetAnalyticByPlatformResponse

`func NewGetAnalyticByPlatformResponse(days float32, breakdown []GetAnalyticByPlatformResponseBreakdown, ) *GetAnalyticByPlatformResponse`

NewGetAnalyticByPlatformResponse instantiates a new GetAnalyticByPlatformResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalyticByPlatformResponseWithDefaults

`func NewGetAnalyticByPlatformResponseWithDefaults() *GetAnalyticByPlatformResponse`

NewGetAnalyticByPlatformResponseWithDefaults instantiates a new GetAnalyticByPlatformResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDays

`func (o *GetAnalyticByPlatformResponse) GetDays() float32`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *GetAnalyticByPlatformResponse) GetDaysOk() (*float32, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *GetAnalyticByPlatformResponse) SetDays(v float32)`

SetDays sets Days field to given value.


### GetBreakdown

`func (o *GetAnalyticByPlatformResponse) GetBreakdown() []GetAnalyticByPlatformResponseBreakdown`

GetBreakdown returns the Breakdown field if non-nil, zero value otherwise.

### GetBreakdownOk

`func (o *GetAnalyticByPlatformResponse) GetBreakdownOk() (*[]GetAnalyticByPlatformResponseBreakdown, bool)`

GetBreakdownOk returns a tuple with the Breakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdown

`func (o *GetAnalyticByPlatformResponse) SetBreakdown(v []GetAnalyticByPlatformResponseBreakdown)`

SetBreakdown sets Breakdown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


