# GetReturnResponse

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

### NewGetReturnResponse

`func NewGetReturnResponse(id string, createdAt time.Time, userId string, status string, requestedAt time.Time, orderId string, identityCheck string, ) *GetReturnResponse`

NewGetReturnResponse instantiates a new GetReturnResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetReturnResponseWithDefaults

`func NewGetReturnResponseWithDefaults() *GetReturnResponse`

NewGetReturnResponseWithDefaults instantiates a new GetReturnResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetReturnResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetReturnResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetReturnResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetReturnResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetReturnResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetReturnResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *GetReturnResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetReturnResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetReturnResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetNotes

`func (o *GetReturnResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetReturnResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetReturnResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetReturnResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *GetReturnResponse) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *GetReturnResponse) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetStatus

`func (o *GetReturnResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetReturnResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetReturnResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetReason

`func (o *GetReturnResponse) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *GetReturnResponse) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *GetReturnResponse) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *GetReturnResponse) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *GetReturnResponse) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *GetReturnResponse) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetRequestedAt

`func (o *GetReturnResponse) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *GetReturnResponse) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *GetReturnResponse) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.


### GetOrderId

`func (o *GetReturnResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetReturnResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetReturnResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetReceivedAt

`func (o *GetReturnResponse) GetReceivedAt() time.Time`

GetReceivedAt returns the ReceivedAt field if non-nil, zero value otherwise.

### GetReceivedAtOk

`func (o *GetReturnResponse) GetReceivedAtOk() (*time.Time, bool)`

GetReceivedAtOk returns a tuple with the ReceivedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceivedAt

`func (o *GetReturnResponse) SetReceivedAt(v time.Time)`

SetReceivedAt sets ReceivedAt field to given value.

### HasReceivedAt

`func (o *GetReturnResponse) HasReceivedAt() bool`

HasReceivedAt returns a boolean if a field has been set.

### SetReceivedAtNil

`func (o *GetReturnResponse) SetReceivedAtNil(b bool)`

 SetReceivedAtNil sets the value for ReceivedAt to be an explicit nil

### UnsetReceivedAt
`func (o *GetReturnResponse) UnsetReceivedAt()`

UnsetReceivedAt ensures that no value is present for ReceivedAt, not even an explicit nil
### GetRefundAmount

`func (o *GetReturnResponse) GetRefundAmount() string`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *GetReturnResponse) GetRefundAmountOk() (*string, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *GetReturnResponse) SetRefundAmount(v string)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *GetReturnResponse) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### SetRefundAmountNil

`func (o *GetReturnResponse) SetRefundAmountNil(b bool)`

 SetRefundAmountNil sets the value for RefundAmount to be an explicit nil

### UnsetRefundAmount
`func (o *GetReturnResponse) UnsetRefundAmount()`

UnsetRefundAmount ensures that no value is present for RefundAmount, not even an explicit nil
### GetRestockedToInventoryItemId

`func (o *GetReturnResponse) GetRestockedToInventoryItemId() string`

GetRestockedToInventoryItemId returns the RestockedToInventoryItemId field if non-nil, zero value otherwise.

### GetRestockedToInventoryItemIdOk

`func (o *GetReturnResponse) GetRestockedToInventoryItemIdOk() (*string, bool)`

GetRestockedToInventoryItemIdOk returns a tuple with the RestockedToInventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestockedToInventoryItemId

`func (o *GetReturnResponse) SetRestockedToInventoryItemId(v string)`

SetRestockedToInventoryItemId sets RestockedToInventoryItemId field to given value.

### HasRestockedToInventoryItemId

`func (o *GetReturnResponse) HasRestockedToInventoryItemId() bool`

HasRestockedToInventoryItemId returns a boolean if a field has been set.

### SetRestockedToInventoryItemIdNil

`func (o *GetReturnResponse) SetRestockedToInventoryItemIdNil(b bool)`

 SetRestockedToInventoryItemIdNil sets the value for RestockedToInventoryItemId to be an explicit nil

### UnsetRestockedToInventoryItemId
`func (o *GetReturnResponse) UnsetRestockedToInventoryItemId()`

UnsetRestockedToInventoryItemId ensures that no value is present for RestockedToInventoryItemId, not even an explicit nil
### GetInspectedAt

`func (o *GetReturnResponse) GetInspectedAt() time.Time`

GetInspectedAt returns the InspectedAt field if non-nil, zero value otherwise.

### GetInspectedAtOk

`func (o *GetReturnResponse) GetInspectedAtOk() (*time.Time, bool)`

GetInspectedAtOk returns a tuple with the InspectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedAt

`func (o *GetReturnResponse) SetInspectedAt(v time.Time)`

SetInspectedAt sets InspectedAt field to given value.

### HasInspectedAt

`func (o *GetReturnResponse) HasInspectedAt() bool`

HasInspectedAt returns a boolean if a field has been set.

### SetInspectedAtNil

`func (o *GetReturnResponse) SetInspectedAtNil(b bool)`

 SetInspectedAtNil sets the value for InspectedAt to be an explicit nil

### UnsetInspectedAt
`func (o *GetReturnResponse) UnsetInspectedAt()`

UnsetInspectedAt ensures that no value is present for InspectedAt, not even an explicit nil
### GetInspectedUnitId

`func (o *GetReturnResponse) GetInspectedUnitId() string`

GetInspectedUnitId returns the InspectedUnitId field if non-nil, zero value otherwise.

### GetInspectedUnitIdOk

`func (o *GetReturnResponse) GetInspectedUnitIdOk() (*string, bool)`

GetInspectedUnitIdOk returns a tuple with the InspectedUnitId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedUnitId

`func (o *GetReturnResponse) SetInspectedUnitId(v string)`

SetInspectedUnitId sets InspectedUnitId field to given value.

### HasInspectedUnitId

`func (o *GetReturnResponse) HasInspectedUnitId() bool`

HasInspectedUnitId returns a boolean if a field has been set.

### SetInspectedUnitIdNil

`func (o *GetReturnResponse) SetInspectedUnitIdNil(b bool)`

 SetInspectedUnitIdNil sets the value for InspectedUnitId to be an explicit nil

### UnsetInspectedUnitId
`func (o *GetReturnResponse) UnsetInspectedUnitId()`

UnsetInspectedUnitId ensures that no value is present for InspectedUnitId, not even an explicit nil
### GetInspectedIdentifier

`func (o *GetReturnResponse) GetInspectedIdentifier() string`

GetInspectedIdentifier returns the InspectedIdentifier field if non-nil, zero value otherwise.

### GetInspectedIdentifierOk

`func (o *GetReturnResponse) GetInspectedIdentifierOk() (*string, bool)`

GetInspectedIdentifierOk returns a tuple with the InspectedIdentifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInspectedIdentifier

`func (o *GetReturnResponse) SetInspectedIdentifier(v string)`

SetInspectedIdentifier sets InspectedIdentifier field to given value.

### HasInspectedIdentifier

`func (o *GetReturnResponse) HasInspectedIdentifier() bool`

HasInspectedIdentifier returns a boolean if a field has been set.

### SetInspectedIdentifierNil

`func (o *GetReturnResponse) SetInspectedIdentifierNil(b bool)`

 SetInspectedIdentifierNil sets the value for InspectedIdentifier to be an explicit nil

### UnsetInspectedIdentifier
`func (o *GetReturnResponse) UnsetInspectedIdentifier()`

UnsetInspectedIdentifier ensures that no value is present for InspectedIdentifier, not even an explicit nil
### GetIdentityCheck

`func (o *GetReturnResponse) GetIdentityCheck() string`

GetIdentityCheck returns the IdentityCheck field if non-nil, zero value otherwise.

### GetIdentityCheckOk

`func (o *GetReturnResponse) GetIdentityCheckOk() (*string, bool)`

GetIdentityCheckOk returns a tuple with the IdentityCheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityCheck

`func (o *GetReturnResponse) SetIdentityCheck(v string)`

SetIdentityCheck sets IdentityCheck field to given value.


### GetRestockedAt

`func (o *GetReturnResponse) GetRestockedAt() time.Time`

GetRestockedAt returns the RestockedAt field if non-nil, zero value otherwise.

### GetRestockedAtOk

`func (o *GetReturnResponse) GetRestockedAtOk() (*time.Time, bool)`

GetRestockedAtOk returns a tuple with the RestockedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestockedAt

`func (o *GetReturnResponse) SetRestockedAt(v time.Time)`

SetRestockedAt sets RestockedAt field to given value.

### HasRestockedAt

`func (o *GetReturnResponse) HasRestockedAt() bool`

HasRestockedAt returns a boolean if a field has been set.

### SetRestockedAtNil

`func (o *GetReturnResponse) SetRestockedAtNil(b bool)`

 SetRestockedAtNil sets the value for RestockedAt to be an explicit nil

### UnsetRestockedAt
`func (o *GetReturnResponse) UnsetRestockedAt()`

UnsetRestockedAt ensures that no value is present for RestockedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


