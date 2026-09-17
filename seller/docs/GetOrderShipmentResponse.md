# GetOrderShipmentResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetOrderShipmentResponseData**](GetOrderShipmentResponseData.md) |  | 
**TotalLabelCostCents** | **float32** |  | 

## Methods

### NewGetOrderShipmentResponse

`func NewGetOrderShipmentResponse(data []GetOrderShipmentResponseData, totalLabelCostCents float32, ) *GetOrderShipmentResponse`

NewGetOrderShipmentResponse instantiates a new GetOrderShipmentResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderShipmentResponseWithDefaults

`func NewGetOrderShipmentResponseWithDefaults() *GetOrderShipmentResponse`

NewGetOrderShipmentResponseWithDefaults instantiates a new GetOrderShipmentResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetOrderShipmentResponse) GetData() []GetOrderShipmentResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetOrderShipmentResponse) GetDataOk() (*[]GetOrderShipmentResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetOrderShipmentResponse) SetData(v []GetOrderShipmentResponseData)`

SetData sets Data field to given value.


### GetTotalLabelCostCents

`func (o *GetOrderShipmentResponse) GetTotalLabelCostCents() float32`

GetTotalLabelCostCents returns the TotalLabelCostCents field if non-nil, zero value otherwise.

### GetTotalLabelCostCentsOk

`func (o *GetOrderShipmentResponse) GetTotalLabelCostCentsOk() (*float32, bool)`

GetTotalLabelCostCentsOk returns a tuple with the TotalLabelCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalLabelCostCents

`func (o *GetOrderShipmentResponse) SetTotalLabelCostCents(v float32)`

SetTotalLabelCostCents sets TotalLabelCostCents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


