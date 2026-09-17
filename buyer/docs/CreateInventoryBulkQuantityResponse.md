# CreateInventoryBulkQuantityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Affected** | **float32** | Rows whose available stock actually changed. | 
**Skipped** | **float32** | Ids that did not move. Either they were already at that number, or they  aren&#39;t this seller&#39;s. The two are deliberately not distinguished: telling  a caller \&quot;that id isn&#39;t yours\&quot; confirms the id exists. | 
**BulkJobId** | Pointer to **NullableString** | Watchable job for the marketplace fan-out, when one was started. | [optional] 

## Methods

### NewCreateInventoryBulkQuantityResponse

`func NewCreateInventoryBulkQuantityResponse(affected float32, skipped float32, ) *CreateInventoryBulkQuantityResponse`

NewCreateInventoryBulkQuantityResponse instantiates a new CreateInventoryBulkQuantityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInventoryBulkQuantityResponseWithDefaults

`func NewCreateInventoryBulkQuantityResponseWithDefaults() *CreateInventoryBulkQuantityResponse`

NewCreateInventoryBulkQuantityResponseWithDefaults instantiates a new CreateInventoryBulkQuantityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAffected

`func (o *CreateInventoryBulkQuantityResponse) GetAffected() float32`

GetAffected returns the Affected field if non-nil, zero value otherwise.

### GetAffectedOk

`func (o *CreateInventoryBulkQuantityResponse) GetAffectedOk() (*float32, bool)`

GetAffectedOk returns a tuple with the Affected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAffected

`func (o *CreateInventoryBulkQuantityResponse) SetAffected(v float32)`

SetAffected sets Affected field to given value.


### GetSkipped

`func (o *CreateInventoryBulkQuantityResponse) GetSkipped() float32`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *CreateInventoryBulkQuantityResponse) GetSkippedOk() (*float32, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *CreateInventoryBulkQuantityResponse) SetSkipped(v float32)`

SetSkipped sets Skipped field to given value.


### GetBulkJobId

`func (o *CreateInventoryBulkQuantityResponse) GetBulkJobId() string`

GetBulkJobId returns the BulkJobId field if non-nil, zero value otherwise.

### GetBulkJobIdOk

`func (o *CreateInventoryBulkQuantityResponse) GetBulkJobIdOk() (*string, bool)`

GetBulkJobIdOk returns a tuple with the BulkJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkJobId

`func (o *CreateInventoryBulkQuantityResponse) SetBulkJobId(v string)`

SetBulkJobId sets BulkJobId field to given value.

### HasBulkJobId

`func (o *CreateInventoryBulkQuantityResponse) HasBulkJobId() bool`

HasBulkJobId returns a boolean if a field has been set.

### SetBulkJobIdNil

`func (o *CreateInventoryBulkQuantityResponse) SetBulkJobIdNil(b bool)`

 SetBulkJobIdNil sets the value for BulkJobId to be an explicit nil

### UnsetBulkJobId
`func (o *CreateInventoryBulkQuantityResponse) UnsetBulkJobId()`

UnsetBulkJobId ensures that no value is present for BulkJobId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


