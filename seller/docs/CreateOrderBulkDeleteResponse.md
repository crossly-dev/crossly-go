# CreateOrderBulkDeleteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | **float32** |  | 
**HardDelete** | **bool** |  | 

## Methods

### NewCreateOrderBulkDeleteResponse

`func NewCreateOrderBulkDeleteResponse(deleted float32, hardDelete bool, ) *CreateOrderBulkDeleteResponse`

NewCreateOrderBulkDeleteResponse instantiates a new CreateOrderBulkDeleteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderBulkDeleteResponseWithDefaults

`func NewCreateOrderBulkDeleteResponseWithDefaults() *CreateOrderBulkDeleteResponse`

NewCreateOrderBulkDeleteResponseWithDefaults instantiates a new CreateOrderBulkDeleteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *CreateOrderBulkDeleteResponse) GetDeleted() float32`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *CreateOrderBulkDeleteResponse) GetDeletedOk() (*float32, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *CreateOrderBulkDeleteResponse) SetDeleted(v float32)`

SetDeleted sets Deleted field to given value.


### GetHardDelete

`func (o *CreateOrderBulkDeleteResponse) GetHardDelete() bool`

GetHardDelete returns the HardDelete field if non-nil, zero value otherwise.

### GetHardDeleteOk

`func (o *CreateOrderBulkDeleteResponse) GetHardDeleteOk() (*bool, bool)`

GetHardDeleteOk returns a tuple with the HardDelete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHardDelete

`func (o *CreateOrderBulkDeleteResponse) SetHardDelete(v bool)`

SetHardDelete sets HardDelete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


