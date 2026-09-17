# CreateOrderCancelResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** |  | 
**Platform** | **string** |  | 
**CancelledInCrossly** | **bool** |  | 
**PlatformCancel** | **string** |  | 

## Methods

### NewCreateOrderCancelResponse

`func NewCreateOrderCancelResponse(orderId string, platform string, cancelledInCrossly bool, platformCancel string, ) *CreateOrderCancelResponse`

NewCreateOrderCancelResponse instantiates a new CreateOrderCancelResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderCancelResponseWithDefaults

`func NewCreateOrderCancelResponseWithDefaults() *CreateOrderCancelResponse`

NewCreateOrderCancelResponseWithDefaults instantiates a new CreateOrderCancelResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CreateOrderCancelResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CreateOrderCancelResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CreateOrderCancelResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetPlatform

`func (o *CreateOrderCancelResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateOrderCancelResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateOrderCancelResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetCancelledInCrossly

`func (o *CreateOrderCancelResponse) GetCancelledInCrossly() bool`

GetCancelledInCrossly returns the CancelledInCrossly field if non-nil, zero value otherwise.

### GetCancelledInCrosslyOk

`func (o *CreateOrderCancelResponse) GetCancelledInCrosslyOk() (*bool, bool)`

GetCancelledInCrosslyOk returns a tuple with the CancelledInCrossly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledInCrossly

`func (o *CreateOrderCancelResponse) SetCancelledInCrossly(v bool)`

SetCancelledInCrossly sets CancelledInCrossly field to given value.


### GetPlatformCancel

`func (o *CreateOrderCancelResponse) GetPlatformCancel() string`

GetPlatformCancel returns the PlatformCancel field if non-nil, zero value otherwise.

### GetPlatformCancelOk

`func (o *CreateOrderCancelResponse) GetPlatformCancelOk() (*string, bool)`

GetPlatformCancelOk returns a tuple with the PlatformCancel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformCancel

`func (o *CreateOrderCancelResponse) SetPlatformCancel(v string)`

SetPlatformCancel sets PlatformCancel field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


