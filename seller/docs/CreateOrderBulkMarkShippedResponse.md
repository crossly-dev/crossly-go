# CreateOrderBulkMarkShippedResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Updated** | **float32** |  | 
**Skipped** | **float32** | Cancelled/refunded orders in the request that were left untouched — reported so a caller can&#39;t mistake a partial success for \&quot;all done\&quot; and miss that a closed order was in the batch. | 

## Methods

### NewCreateOrderBulkMarkShippedResponse

`func NewCreateOrderBulkMarkShippedResponse(updated float32, skipped float32, ) *CreateOrderBulkMarkShippedResponse`

NewCreateOrderBulkMarkShippedResponse instantiates a new CreateOrderBulkMarkShippedResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderBulkMarkShippedResponseWithDefaults

`func NewCreateOrderBulkMarkShippedResponseWithDefaults() *CreateOrderBulkMarkShippedResponse`

NewCreateOrderBulkMarkShippedResponseWithDefaults instantiates a new CreateOrderBulkMarkShippedResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUpdated

`func (o *CreateOrderBulkMarkShippedResponse) GetUpdated() float32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CreateOrderBulkMarkShippedResponse) GetUpdatedOk() (*float32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CreateOrderBulkMarkShippedResponse) SetUpdated(v float32)`

SetUpdated sets Updated field to given value.


### GetSkipped

`func (o *CreateOrderBulkMarkShippedResponse) GetSkipped() float32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *CreateOrderBulkMarkShippedResponse) GetSkippedOk() (*float32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *CreateOrderBulkMarkShippedResponse) SetSkipped(v float32)`

SetSkipped sets Skipped field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


