# CreateOrderMessageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** |  | 
**Sent** | **bool** |  | 
**PlatformConversationId** | **string** |  | 

## Methods

### NewCreateOrderMessageResponse

`func NewCreateOrderMessageResponse(orderId string, sent bool, platformConversationId string, ) *CreateOrderMessageResponse`

NewCreateOrderMessageResponse instantiates a new CreateOrderMessageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderMessageResponseWithDefaults

`func NewCreateOrderMessageResponseWithDefaults() *CreateOrderMessageResponse`

NewCreateOrderMessageResponseWithDefaults instantiates a new CreateOrderMessageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *CreateOrderMessageResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CreateOrderMessageResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CreateOrderMessageResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetSent

`func (o *CreateOrderMessageResponse) GetSent() bool`

GetSent returns the Sent field if non-nil, zero value otherwise.

### GetSentOk

`func (o *CreateOrderMessageResponse) GetSentOk() (*bool, bool)`

GetSentOk returns a tuple with the Sent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSent

`func (o *CreateOrderMessageResponse) SetSent(v bool)`

SetSent sets Sent field to given value.


### GetPlatformConversationId

`func (o *CreateOrderMessageResponse) GetPlatformConversationId() string`

GetPlatformConversationId returns the PlatformConversationId field if non-nil, zero value otherwise.

### GetPlatformConversationIdOk

`func (o *CreateOrderMessageResponse) GetPlatformConversationIdOk() (*string, bool)`

GetPlatformConversationIdOk returns a tuple with the PlatformConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformConversationId

`func (o *CreateOrderMessageResponse) SetPlatformConversationId(v string)`

SetPlatformConversationId sets PlatformConversationId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


