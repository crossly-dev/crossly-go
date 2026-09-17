# CreateOrderRefundResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RefundId** | **string** |  | 
**Amount** | **float32** |  | 

## Methods

### NewCreateOrderRefundResponse

`func NewCreateOrderRefundResponse(refundId string, amount float32, ) *CreateOrderRefundResponse`

NewCreateOrderRefundResponse instantiates a new CreateOrderRefundResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderRefundResponseWithDefaults

`func NewCreateOrderRefundResponseWithDefaults() *CreateOrderRefundResponse`

NewCreateOrderRefundResponseWithDefaults instantiates a new CreateOrderRefundResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRefundId

`func (o *CreateOrderRefundResponse) GetRefundId() string`

GetRefundId returns the RefundId field if non-nil, zero value otherwise.

### GetRefundIdOk

`func (o *CreateOrderRefundResponse) GetRefundIdOk() (*string, bool)`

GetRefundIdOk returns a tuple with the RefundId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundId

`func (o *CreateOrderRefundResponse) SetRefundId(v string)`

SetRefundId sets RefundId field to given value.


### GetAmount

`func (o *CreateOrderRefundResponse) GetAmount() float32`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *CreateOrderRefundResponse) GetAmountOk() (*float32, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *CreateOrderRefundResponse) SetAmount(v float32)`

SetAmount sets Amount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


