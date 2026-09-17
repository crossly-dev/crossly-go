# CreateReturnResponse

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

### NewCreateReturnResponse

`func NewCreateReturnResponse(id string, createdAt time.Time, userId string, status string, requestedAt time.Time, orderId string, identityCheck string, ) *CreateReturnResponse`

NewCreateReturnResponse instantiates a new CreateReturnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateReturnResponseWithDefaults

`func NewCreateReturnResponseWithDefaults() *CreateReturnResponse`

NewCreateReturnResponseWithDefaults instantiates a new CreateReturnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateReturnResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateReturnResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateReturnResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CreateReturnResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateReturnResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateReturnResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *CreateReturnResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateReturnResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateReturnResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetNotes

`func (o *CreateReturnResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateReturnResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateReturnResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateReturnResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *CreateReturnResponse) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *CreateReturnResponse) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetStatus

`func (o *CreateReturnResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateReturnResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateReturnResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *CreateReturnResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *CreateReturnResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *CreateReturnResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *CreateReturnResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *CreateReturnResponse) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *CreateReturnResponse) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetRequestedAt

`func (o *CreateReturnResponse) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *CreateReturnResponse) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *CreateReturnResponse) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.


### GetOrderId

`func (o *CreateReturnResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CreateReturnResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CreateReturnResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetReceivedAt

`func (o *CreateReturnResponse) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *CreateReturnResponse) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *CreateReturnResponse) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *CreateReturnResponse) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### SetReceivedAtNil

`func (o *CreateReturnResponse) SetReceivedAtNil(b bool)`

 SetReceivedAtNil sets the value for ReceivedAt to be an explicit nil

### UnsetReceivedAt
`func (o *CreateReturnResponse) UnsetReceivedAt()`

UnsetReceivedAt ensures that no value is present for ReceivedAt, not even an explicit nil
### GetRefundAmount

`func (o *CreateReturnResponse) GetRefundAmount() string`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *CreateReturnResponse) GetRefundAmountOk() (*string, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *CreateReturnResponse) SetRefundAmount(v string)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *CreateReturnResponse) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### SetRefundAmountNil

`func (o *CreateReturnResponse) SetRefundAmountNil(b bool)`

 SetRefundAmountNil sets the value for RefundAmount to be an explicit nil

### UnsetRefundAmount
`func (o *CreateReturnResponse) UnsetRefundAmount()`

UnsetRefundAmount ensures that no value is present for RefundAmount, not even an explicit nil
### GetRestockedToInventoryItemId

`func (o *CreateReturnResponse) GetRestockedToInventoryItemId() string`

GetRestockedToInventoryItemId returns the RestockedToInventoryItemId field if non-nil, zero value otherwise.

### GetRestockedToInventoryItemIdOk

`func (o *CreateReturnResponse) GetRestockedToInventoryItemIdOk() (*string, bool)`

GetRestockedToInventoryItemIdOk returns a tuple with the RestockedToInventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestockedToInventoryItemId

`func (o *CreateReturnResponse) SetRestockedToInventoryItemId(v string)`

SetRestockedToInventoryItemId sets RestockedToInventoryItemId field to given value.

### HasRestockedToInventoryItemId

`func (o *CreateReturnResponse) HasRestockedToInventoryItemId() bool`

HasRestockedToInventoryItemId returns a boolean if a field has been set.

### SetRestockedToInventoryItemIdNil

`func (o *CreateReturnResponse) SetRestockedToInventoryItemIdNil(b bool)`

 SetRestockedToInventoryItemIdNil sets the value for RestockedToInventoryItemId to be an explicit nil

### UnsetRestockedToInventoryItemId
`func (o *CreateReturnResponse) UnsetRestockedToInventoryItemId()`

UnsetRestockedToInventoryItemId ensures that no value is present for RestockedToInventoryItemId, not even an explicit nil
### GetInspectedAt

`func (o *CreateReturnResponse) GetInspectedAt() time.Time`

GetInspectedAt returns the InspectedAt field if non-nil, zero value otherwise.

### GetInspectedAtOk

`func (o *CreateReturnResponse) GetInspectedAtOk() (*time.Time, bool)`

GetInspectedAtOk returns a tuple with the InspectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedAt

`func (o *CreateReturnResponse) SetInspectedAt(v time.Time)`

SetInspectedAt sets InspectedAt field to given value.

### HasInspectedAt

`func (o *CreateReturnResponse) HasInspectedAt() bool`

HasInspectedAt returns a boolean if a field has been set.

### SetInspectedAtNil

`func (o *CreateReturnResponse) SetInspectedAtNil(b bool)`

 SetInspectedAtNil sets the value for InspectedAt to be an explicit nil

### UnsetInspectedAt
`func (o *CreateReturnResponse) UnsetInspectedAt()`

UnsetInspectedAt ensures that no value is present for InspectedAt, not even an explicit nil
### GetInspectedUnitId

`func (o *CreateReturnResponse) GetInspectedUnitId() string`

GetInspectedUnitId returns the InspectedUnitId field if non-nil, zero value otherwise.

### GetInspectedUnitIdOk

`func (o *CreateReturnResponse) GetInspectedUnitIdOk() (*string, bool)`

GetInspectedUnitIdOk returns a tuple with the InspectedUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedUnitId

`func (o *CreateReturnResponse) SetInspectedUnitId(v string)`

SetInspectedUnitId sets InspectedUnitId field to given value.

### HasInspectedUnitId

`func (o *CreateReturnResponse) HasInspectedUnitId() bool`

HasInspectedUnitId returns a boolean if a field has been set.

### SetInspectedUnitIdNil

`func (o *CreateReturnResponse) SetInspectedUnitIdNil(b bool)`

 SetInspectedUnitIdNil sets the value for InspectedUnitId to be an explicit nil

### UnsetInspectedUnitId
`func (o *CreateReturnResponse) UnsetInspectedUnitId()`

UnsetInspectedUnitId ensures that no value is present for InspectedUnitId, not even an explicit nil
### GetInspectedIdentifier

`func (o *CreateReturnResponse) GetInspectedIdentifier() string`

GetInspectedIdentifier returns the InspectedIdentifier field if non-nil, zero value otherwise.

### GetInspectedIdentifierOk

`func (o *CreateReturnResponse) GetInspectedIdentifierOk() (*string, bool)`

GetInspectedIdentifierOk returns a tuple with the InspectedIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedIdentifier

`func (o *CreateReturnResponse) SetInspectedIdentifier(v string)`

SetInspectedIdentifier sets InspectedIdentifier field to given value.

### HasInspectedIdentifier

`func (o *CreateReturnResponse) HasInspectedIdentifier() bool`

HasInspectedIdentifier returns a boolean if a field has been set.

### SetInspectedIdentifierNil

`func (o *CreateReturnResponse) SetInspectedIdentifierNil(b bool)`

 SetInspectedIdentifierNil sets the value for InspectedIdentifier to be an explicit nil

### UnsetInspectedIdentifier
`func (o *CreateReturnResponse) UnsetInspectedIdentifier()`

UnsetInspectedIdentifier ensures that no value is present for InspectedIdentifier, not even an explicit nil
### GetIdentityCheck

`func (o *CreateReturnResponse) GetIdentityCheck() string`

GetIdentityCheck returns the IdentityCheck field if non-nil, zero value otherwise.

### GetIdentityCheckOk

`func (o *CreateReturnResponse) GetIdentityCheckOk() (*string, bool)`

GetIdentityCheckOk returns a tuple with the IdentityCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCheck

`func (o *CreateReturnResponse) SetIdentityCheck(v string)`

SetIdentityCheck sets IdentityCheck field to given value.


### GetRestockedAt

`func (o *CreateReturnResponse) GetRestockedAt() time.Time`

GetRestockedAt returns the RestockedAt field if non-nil, zero value otherwise.

### GetRestockedAtOk

`func (o *CreateReturnResponse) GetRestockedAtOk() (*time.Time, bool)`

GetRestockedAtOk returns a tuple with the RestockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestockedAt

`func (o *CreateReturnResponse) SetRestockedAt(v time.Time)`

SetRestockedAt sets RestockedAt field to given value.

### HasRestockedAt

`func (o *CreateReturnResponse) HasRestockedAt() bool`

HasRestockedAt returns a boolean if a field has been set.

### SetRestockedAtNil

`func (o *CreateReturnResponse) SetRestockedAtNil(b bool)`

 SetRestockedAtNil sets the value for RestockedAt to be an explicit nil

### UnsetRestockedAt
`func (o *CreateReturnResponse) UnsetRestockedAt()`

UnsetRestockedAt ensures that no value is present for RestockedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


