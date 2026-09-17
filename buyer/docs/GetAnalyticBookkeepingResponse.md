# GetAnalyticBookkeepingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Year** | **float32** |  | 
**MonthlyPL** | [**[]GetAnalyticBookkeepingResponseMonthlyPL**](GetAnalyticBookkeepingResponseMonthlyPL.md) |  | 
**AnnualTotals** | [**GetAnalyticBookkeepingResponseAnnualTotals**](GetAnalyticBookkeepingResponseAnnualTotals.md) |  | 
**PlatformBreakdown** | [**[]GetAnalyticBookkeepingResponsePlatformBreakdown**](GetAnalyticBookkeepingResponsePlatformBreakdown.md) |  | 

## Methods

### NewGetAnalyticBookkeepingResponse

`func NewGetAnalyticBookkeepingResponse(year float32, monthlyPL []GetAnalyticBookkeepingResponseMonthlyPL, annualTotals GetAnalyticBookkeepingResponseAnnualTotals, platformBreakdown []GetAnalyticBookkeepingResponsePlatformBreakdown, ) *GetAnalyticBookkeepingResponse`

NewGetAnalyticBookkeepingResponse instantiates a new GetAnalyticBookkeepingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalyticBookkeepingResponseWithDefaults

`func NewGetAnalyticBookkeepingResponseWithDefaults() *GetAnalyticBookkeepingResponse`

NewGetAnalyticBookkeepingResponseWithDefaults instantiates a new GetAnalyticBookkeepingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYear

`func (o *GetAnalyticBookkeepingResponse) GetYear() float32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *GetAnalyticBookkeepingResponse) GetYearOk() (*float32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *GetAnalyticBookkeepingResponse) SetYear(v float32)`

SetYear sets Year field to given value.


### GetMonthlyPL

`func (o *GetAnalyticBookkeepingResponse) GetMonthlyPL() []GetAnalyticBookkeepingResponseMonthlyPL`

GetMonthlyPL returns the MonthlyPL field if non-nil, zero value otherwise.

### GetMonthlyPLOk

`func (o *GetAnalyticBookkeepingResponse) GetMonthlyPLOk() (*[]GetAnalyticBookkeepingResponseMonthlyPL, bool)`

GetMonthlyPLOk returns a tuple with the MonthlyPL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthlyPL

`func (o *GetAnalyticBookkeepingResponse) SetMonthlyPL(v []GetAnalyticBookkeepingResponseMonthlyPL)`

SetMonthlyPL sets MonthlyPL field to given value.


### GetAnnualTotals

`func (o *GetAnalyticBookkeepingResponse) GetAnnualTotals() GetAnalyticBookkeepingResponseAnnualTotals`

GetAnnualTotals returns the AnnualTotals field if non-nil, zero value otherwise.

### GetAnnualTotalsOk

`func (o *GetAnalyticBookkeepingResponse) GetAnnualTotalsOk() (*GetAnalyticBookkeepingResponseAnnualTotals, bool)`

GetAnnualTotalsOk returns a tuple with the AnnualTotals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnualTotals

`func (o *GetAnalyticBookkeepingResponse) SetAnnualTotals(v GetAnalyticBookkeepingResponseAnnualTotals)`

SetAnnualTotals sets AnnualTotals field to given value.


### GetPlatformBreakdown

`func (o *GetAnalyticBookkeepingResponse) GetPlatformBreakdown() []GetAnalyticBookkeepingResponsePlatformBreakdown`

GetPlatformBreakdown returns the PlatformBreakdown field if non-nil, zero value otherwise.

### GetPlatformBreakdownOk

`func (o *GetAnalyticBookkeepingResponse) GetPlatformBreakdownOk() (*[]GetAnalyticBookkeepingResponsePlatformBreakdown, bool)`

GetPlatformBreakdownOk returns a tuple with the PlatformBreakdown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformBreakdown

`func (o *GetAnalyticBookkeepingResponse) SetPlatformBreakdown(v []GetAnalyticBookkeepingResponsePlatformBreakdown)`

SetPlatformBreakdown sets PlatformBreakdown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


