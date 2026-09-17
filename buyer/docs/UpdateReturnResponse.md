# UpdateReturnResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**Reason** | Pointer to **NullableString** |  | [optional] 
**RequestedAt** | **time.Time** |  | 
**OrderId** | **string** |  | 
**ReceivedAt** | Pointer to **NullableTime** |  | [optional] 
**RefundAmount** | Pointer to **NullableString** |  | [optional] 
**RestockedToInventoryItemId** | Pointer to **NullableString** |  | [optional] 
**InspectedAt** | Pointer to **NullableTime** |  | [optional] 
**InspectedUnitId** | Pointer to **NullableString** |  | [optional] 
**InspectedIdentifier** | Pointer to **NullableString** |  | [optional] 
**IdentityCheck** | **string** |  | 
**RestockedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewUpdateReturnResponse

`func NewUpdateReturnResponse(id string, createdAt time.Time, userId string, status string, requestedAt time.Time, orderId string, identityCheck string, ) *UpdateReturnResponse`

NewUpdateReturnResponse instantiates a new UpdateReturnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateReturnResponseWithDefaults

`func NewUpdateReturnResponseWithDefaults() *UpdateReturnResponse`

NewUpdateReturnResponseWithDefaults instantiates a new UpdateReturnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateReturnResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateReturnResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateReturnResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *UpdateReturnResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateReturnResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateReturnResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *UpdateReturnResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateReturnResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateReturnResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetNotes

`func (o *UpdateReturnResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateReturnResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateReturnResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateReturnResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *UpdateReturnResponse) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *UpdateReturnResponse) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetStatus

`func (o *UpdateReturnResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateReturnResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateReturnResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *UpdateReturnResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *UpdateReturnResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *UpdateReturnResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *UpdateReturnResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *UpdateReturnResponse) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *UpdateReturnResponse) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetRequestedAt

`func (o *UpdateReturnResponse) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *UpdateReturnResponse) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *UpdateReturnResponse) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.


### GetOrderId

`func (o *UpdateReturnResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *UpdateReturnResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *UpdateReturnResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetReceivedAt

`func (o *UpdateReturnResponse) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *UpdateReturnResponse) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *UpdateReturnResponse) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *UpdateReturnResponse) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### SetReceivedAtNil

`func (o *UpdateReturnResponse) SetReceivedAtNil(b bool)`

 SetReceivedAtNil sets the value for ReceivedAt to be an explicit nil

### UnsetReceivedAt
`func (o *UpdateReturnResponse) UnsetReceivedAt()`

UnsetReceivedAt ensures that no value is present for ReceivedAt, not even an explicit nil
### GetRefundAmount

`func (o *UpdateReturnResponse) GetRefundAmount() string`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *UpdateReturnResponse) GetRefundAmountOk() (*string, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *UpdateReturnResponse) SetRefundAmount(v string)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *UpdateReturnResponse) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### SetRefundAmountNil

`func (o *UpdateReturnResponse) SetRefundAmountNil(b bool)`

 SetRefundAmountNil sets the value for RefundAmount to be an explicit nil

### UnsetRefundAmount
`func (o *UpdateReturnResponse) UnsetRefundAmount()`

UnsetRefundAmount ensures that no value is present for RefundAmount, not even an explicit nil
### GetRestockedToInventoryItemId

`func (o *UpdateReturnResponse) GetRestockedToInventoryItemId() string`

GetRestockedToInventoryItemId returns the RestockedToInventoryItemId field if non-nil, zero value otherwise.

### GetRestockedToInventoryItemIdOk

`func (o *UpdateReturnResponse) GetRestockedToInventoryItemIdOk() (*string, bool)`

GetRestockedToInventoryItemIdOk returns a tuple with the RestockedToInventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestockedToInventoryItemId

`func (o *UpdateReturnResponse) SetRestockedToInventoryItemId(v string)`

SetRestockedToInventoryItemId sets RestockedToInventoryItemId field to given value.

### HasRestockedToInventoryItemId

`func (o *UpdateReturnResponse) HasRestockedToInventoryItemId() bool`

HasRestockedToInventoryItemId returns a boolean if a field has been set.

### SetRestockedToInventoryItemIdNil

`func (o *UpdateReturnResponse) SetRestockedToInventoryItemIdNil(b bool)`

 SetRestockedToInventoryItemIdNil sets the value for RestockedToInventoryItemId to be an explicit nil

### UnsetRestockedToInventoryItemId
`func (o *UpdateReturnResponse) UnsetRestockedToInventoryItemId()`

UnsetRestockedToInventoryItemId ensures that no value is present for RestockedToInventoryItemId, not even an explicit nil
### GetInspectedAt

`func (o *UpdateReturnResponse) GetInspectedAt() time.Time`

GetInspectedAt returns the InspectedAt field if non-nil, zero value otherwise.

### GetInspectedAtOk

`func (o *UpdateReturnResponse) GetInspectedAtOk() (*time.Time, bool)`

GetInspectedAtOk returns a tuple with the InspectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedAt

`func (o *UpdateReturnResponse) SetInspectedAt(v time.Time)`

SetInspectedAt sets InspectedAt field to given value.

### HasInspectedAt

`func (o *UpdateReturnResponse) HasInspectedAt() bool`

HasInspectedAt returns a boolean if a field has been set.

### SetInspectedAtNil

`func (o *UpdateReturnResponse) SetInspectedAtNil(b bool)`

 SetInspectedAtNil sets the value for InspectedAt to be an explicit nil

### UnsetInspectedAt
`func (o *UpdateReturnResponse) UnsetInspectedAt()`

UnsetInspectedAt ensures that no value is present for InspectedAt, not even an explicit nil
### GetInspectedUnitId

`func (o *UpdateReturnResponse) GetInspectedUnitId() string`

GetInspectedUnitId returns the InspectedUnitId field if non-nil, zero value otherwise.

### GetInspectedUnitIdOk

`func (o *UpdateReturnResponse) GetInspectedUnitIdOk() (*string, bool)`

GetInspectedUnitIdOk returns a tuple with the InspectedUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedUnitId

`func (o *UpdateReturnResponse) SetInspectedUnitId(v string)`

SetInspectedUnitId sets InspectedUnitId field to given value.

### HasInspectedUnitId

`func (o *UpdateReturnResponse) HasInspectedUnitId() bool`

HasInspectedUnitId returns a boolean if a field has been set.

### SetInspectedUnitIdNil

`func (o *UpdateReturnResponse) SetInspectedUnitIdNil(b bool)`

 SetInspectedUnitIdNil sets the value for InspectedUnitId to be an explicit nil

### UnsetInspectedUnitId
`func (o *UpdateReturnResponse) UnsetInspectedUnitId()`

UnsetInspectedUnitId ensures that no value is present for InspectedUnitId, not even an explicit nil
### GetInspectedIdentifier

`func (o *UpdateReturnResponse) GetInspectedIdentifier() string`

GetInspectedIdentifier returns the InspectedIdentifier field if non-nil, zero value otherwise.

### GetInspectedIdentifierOk

`func (o *UpdateReturnResponse) GetInspectedIdentifierOk() (*string, bool)`

GetInspectedIdentifierOk returns a tuple with the InspectedIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedIdentifier

`func (o *UpdateReturnResponse) SetInspectedIdentifier(v string)`

SetInspectedIdentifier sets InspectedIdentifier field to given value.

### HasInspectedIdentifier

`func (o *UpdateReturnResponse) HasInspectedIdentifier() bool`

HasInspectedIdentifier returns a boolean if a field has been set.

### SetInspectedIdentifierNil

`func (o *UpdateReturnResponse) SetInspectedIdentifierNil(b bool)`

 SetInspectedIdentifierNil sets the value for InspectedIdentifier to be an explicit nil

### UnsetInspectedIdentifier
`func (o *UpdateReturnResponse) UnsetInspectedIdentifier()`

UnsetInspectedIdentifier ensures that no value is present for InspectedIdentifier, not even an explicit nil
### GetIdentityCheck

`func (o *UpdateReturnResponse) GetIdentityCheck() string`

GetIdentityCheck returns the IdentityCheck field if non-nil, zero value otherwise.

### GetIdentityCheckOk

`func (o *UpdateReturnResponse) GetIdentityCheckOk() (*string, bool)`

GetIdentityCheckOk returns a tuple with the IdentityCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCheck

`func (o *UpdateReturnResponse) SetIdentityCheck(v string)`

SetIdentityCheck sets IdentityCheck field to given value.


### GetRestockedAt

`func (o *UpdateReturnResponse) GetRestockedAt() time.Time`

GetRestockedAt returns the RestockedAt field if non-nil, zero value otherwise.

### GetRestockedAtOk

`func (o *UpdateReturnResponse) GetRestockedAtOk() (*time.Time, bool)`

GetRestockedAtOk returns a tuple with the RestockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestockedAt

`func (o *UpdateReturnResponse) SetRestockedAt(v time.Time)`

SetRestockedAt sets RestockedAt field to given value.

### HasRestockedAt

`func (o *UpdateReturnResponse) HasRestockedAt() bool`

HasRestockedAt returns a boolean if a field has been set.

### SetRestockedAtNil

`func (o *UpdateReturnResponse) SetRestockedAtNil(b bool)`

 SetRestockedAtNil sets the value for RestockedAt to be an explicit nil

### UnsetRestockedAt
`func (o *UpdateReturnResponse) UnsetRestockedAt()`

UnsetRestockedAt ensures that no value is present for RestockedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


