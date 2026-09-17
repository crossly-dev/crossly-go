# ListReturnsItem

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

### NewListReturnsItem

`func NewListReturnsItem(id string, createdAt time.Time, userId string, status string, requestedAt time.Time, orderId string, identityCheck string, ) *ListReturnsItem`

NewListReturnsItem instantiates a new ListReturnsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListReturnsItemWithDefaults

`func NewListReturnsItemWithDefaults() *ListReturnsItem`

NewListReturnsItemWithDefaults instantiates a new ListReturnsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListReturnsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListReturnsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListReturnsItem) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListReturnsItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListReturnsItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListReturnsItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *ListReturnsItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListReturnsItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListReturnsItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetNotes

`func (o *ListReturnsItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ListReturnsItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ListReturnsItem) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ListReturnsItem) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ListReturnsItem) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ListReturnsItem) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetStatus

`func (o *ListReturnsItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListReturnsItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListReturnsItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *ListReturnsItem) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *ListReturnsItem) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *ListReturnsItem) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *ListReturnsItem) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *ListReturnsItem) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *ListReturnsItem) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetRequestedAt

`func (o *ListReturnsItem) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *ListReturnsItem) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *ListReturnsItem) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.


### GetOrderId

`func (o *ListReturnsItem) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ListReturnsItem) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ListReturnsItem) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetReceivedAt

`func (o *ListReturnsItem) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *ListReturnsItem) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *ListReturnsItem) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *ListReturnsItem) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### SetReceivedAtNil

`func (o *ListReturnsItem) SetReceivedAtNil(b bool)`

 SetReceivedAtNil sets the value for ReceivedAt to be an explicit nil

### UnsetReceivedAt
`func (o *ListReturnsItem) UnsetReceivedAt()`

UnsetReceivedAt ensures that no value is present for ReceivedAt, not even an explicit nil
### GetRefundAmount

`func (o *ListReturnsItem) GetRefundAmount() string`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *ListReturnsItem) GetRefundAmountOk() (*string, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *ListReturnsItem) SetRefundAmount(v string)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *ListReturnsItem) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### SetRefundAmountNil

`func (o *ListReturnsItem) SetRefundAmountNil(b bool)`

 SetRefundAmountNil sets the value for RefundAmount to be an explicit nil

### UnsetRefundAmount
`func (o *ListReturnsItem) UnsetRefundAmount()`

UnsetRefundAmount ensures that no value is present for RefundAmount, not even an explicit nil
### GetRestockedToInventoryItemId

`func (o *ListReturnsItem) GetRestockedToInventoryItemId() string`

GetRestockedToInventoryItemId returns the RestockedToInventoryItemId field if non-nil, zero value otherwise.

### GetRestockedToInventoryItemIdOk

`func (o *ListReturnsItem) GetRestockedToInventoryItemIdOk() (*string, bool)`

GetRestockedToInventoryItemIdOk returns a tuple with the RestockedToInventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestockedToInventoryItemId

`func (o *ListReturnsItem) SetRestockedToInventoryItemId(v string)`

SetRestockedToInventoryItemId sets RestockedToInventoryItemId field to given value.

### HasRestockedToInventoryItemId

`func (o *ListReturnsItem) HasRestockedToInventoryItemId() bool`

HasRestockedToInventoryItemId returns a boolean if a field has been set.

### SetRestockedToInventoryItemIdNil

`func (o *ListReturnsItem) SetRestockedToInventoryItemIdNil(b bool)`

 SetRestockedToInventoryItemIdNil sets the value for RestockedToInventoryItemId to be an explicit nil

### UnsetRestockedToInventoryItemId
`func (o *ListReturnsItem) UnsetRestockedToInventoryItemId()`

UnsetRestockedToInventoryItemId ensures that no value is present for RestockedToInventoryItemId, not even an explicit nil
### GetInspectedAt

`func (o *ListReturnsItem) GetInspectedAt() time.Time`

GetInspectedAt returns the InspectedAt field if non-nil, zero value otherwise.

### GetInspectedAtOk

`func (o *ListReturnsItem) GetInspectedAtOk() (*time.Time, bool)`

GetInspectedAtOk returns a tuple with the InspectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedAt

`func (o *ListReturnsItem) SetInspectedAt(v time.Time)`

SetInspectedAt sets InspectedAt field to given value.

### HasInspectedAt

`func (o *ListReturnsItem) HasInspectedAt() bool`

HasInspectedAt returns a boolean if a field has been set.

### SetInspectedAtNil

`func (o *ListReturnsItem) SetInspectedAtNil(b bool)`

 SetInspectedAtNil sets the value for InspectedAt to be an explicit nil

### UnsetInspectedAt
`func (o *ListReturnsItem) UnsetInspectedAt()`

UnsetInspectedAt ensures that no value is present for InspectedAt, not even an explicit nil
### GetInspectedUnitId

`func (o *ListReturnsItem) GetInspectedUnitId() string`

GetInspectedUnitId returns the InspectedUnitId field if non-nil, zero value otherwise.

### GetInspectedUnitIdOk

`func (o *ListReturnsItem) GetInspectedUnitIdOk() (*string, bool)`

GetInspectedUnitIdOk returns a tuple with the InspectedUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedUnitId

`func (o *ListReturnsItem) SetInspectedUnitId(v string)`

SetInspectedUnitId sets InspectedUnitId field to given value.

### HasInspectedUnitId

`func (o *ListReturnsItem) HasInspectedUnitId() bool`

HasInspectedUnitId returns a boolean if a field has been set.

### SetInspectedUnitIdNil

`func (o *ListReturnsItem) SetInspectedUnitIdNil(b bool)`

 SetInspectedUnitIdNil sets the value for InspectedUnitId to be an explicit nil

### UnsetInspectedUnitId
`func (o *ListReturnsItem) UnsetInspectedUnitId()`

UnsetInspectedUnitId ensures that no value is present for InspectedUnitId, not even an explicit nil
### GetInspectedIdentifier

`func (o *ListReturnsItem) GetInspectedIdentifier() string`

GetInspectedIdentifier returns the InspectedIdentifier field if non-nil, zero value otherwise.

### GetInspectedIdentifierOk

`func (o *ListReturnsItem) GetInspectedIdentifierOk() (*string, bool)`

GetInspectedIdentifierOk returns a tuple with the InspectedIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedIdentifier

`func (o *ListReturnsItem) SetInspectedIdentifier(v string)`

SetInspectedIdentifier sets InspectedIdentifier field to given value.

### HasInspectedIdentifier

`func (o *ListReturnsItem) HasInspectedIdentifier() bool`

HasInspectedIdentifier returns a boolean if a field has been set.

### SetInspectedIdentifierNil

`func (o *ListReturnsItem) SetInspectedIdentifierNil(b bool)`

 SetInspectedIdentifierNil sets the value for InspectedIdentifier to be an explicit nil

### UnsetInspectedIdentifier
`func (o *ListReturnsItem) UnsetInspectedIdentifier()`

UnsetInspectedIdentifier ensures that no value is present for InspectedIdentifier, not even an explicit nil
### GetIdentityCheck

`func (o *ListReturnsItem) GetIdentityCheck() string`

GetIdentityCheck returns the IdentityCheck field if non-nil, zero value otherwise.

### GetIdentityCheckOk

`func (o *ListReturnsItem) GetIdentityCheckOk() (*string, bool)`

GetIdentityCheckOk returns a tuple with the IdentityCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCheck

`func (o *ListReturnsItem) SetIdentityCheck(v string)`

SetIdentityCheck sets IdentityCheck field to given value.


### GetRestockedAt

`func (o *ListReturnsItem) GetRestockedAt() time.Time`

GetRestockedAt returns the RestockedAt field if non-nil, zero value otherwise.

### GetRestockedAtOk

`func (o *ListReturnsItem) GetRestockedAtOk() (*time.Time, bool)`

GetRestockedAtOk returns a tuple with the RestockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestockedAt

`func (o *ListReturnsItem) SetRestockedAt(v time.Time)`

SetRestockedAt sets RestockedAt field to given value.

### HasRestockedAt

`func (o *ListReturnsItem) HasRestockedAt() bool`

HasRestockedAt returns a boolean if a field has been set.

### SetRestockedAtNil

`func (o *ListReturnsItem) SetRestockedAtNil(b bool)`

 SetRestockedAtNil sets the value for RestockedAt to be an explicit nil

### UnsetRestockedAt
`func (o *ListReturnsItem) UnsetRestockedAt()`

UnsetRestockedAt ensures that no value is present for RestockedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


