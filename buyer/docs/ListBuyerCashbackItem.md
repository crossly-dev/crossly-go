# ListBuyerCashbackItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Status** | **string** |  | 
**StoreName** | **string** |  | 
**StoreHost** | **string** |  | 
**OrderCents** | Pointer to **NullableFloat32** |  | [optional] 
**CashbackCents** | Pointer to **NullableFloat32** | What WE will pay. Never the commission — that is our revenue and is not the shopper&#39;s business. | [optional] 
**ActivatedAt** | **time.Time** |  | 
**ConfirmedAt** | Pointer to **NullableTime** |  | [optional] 
**PaidAt** | Pointer to **NullableTime** |  | [optional] 
**ExpiredAt** | Pointer to **NullableTime** |  | [optional] 
**RejectedReason** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListBuyerCashbackItem

`func NewListBuyerCashbackItem(id string, status string, storeName string, storeHost string, activatedAt time.Time, ) *ListBuyerCashbackItem`

NewListBuyerCashbackItem instantiates a new ListBuyerCashbackItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBuyerCashbackItemWithDefaults

`func NewListBuyerCashbackItemWithDefaults() *ListBuyerCashbackItem`

NewListBuyerCashbackItemWithDefaults instantiates a new ListBuyerCashbackItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListBuyerCashbackItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBuyerCashbackItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBuyerCashbackItem) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *ListBuyerCashbackItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListBuyerCashbackItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListBuyerCashbackItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetStoreName

`func (o *ListBuyerCashbackItem) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *ListBuyerCashbackItem) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *ListBuyerCashbackItem) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.


### GetStoreHost

`func (o *ListBuyerCashbackItem) GetStoreHost() string`

GetStoreHost returns the StoreHost field if non-nil, zero value otherwise.

### GetStoreHostOk

`func (o *ListBuyerCashbackItem) GetStoreHostOk() (*string, bool)`

GetStoreHostOk returns a tuple with the StoreHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreHost

`func (o *ListBuyerCashbackItem) SetStoreHost(v string)`

SetStoreHost sets StoreHost field to given value.


### GetOrderCents

`func (o *ListBuyerCashbackItem) GetOrderCents() float32`

GetOrderCents returns the OrderCents field if non-nil, zero value otherwise.

### GetOrderCentsOk

`func (o *ListBuyerCashbackItem) GetOrderCentsOk() (*float32, bool)`

GetOrderCentsOk returns a tuple with the OrderCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderCents

`func (o *ListBuyerCashbackItem) SetOrderCents(v float32)`

SetOrderCents sets OrderCents field to given value.

### HasOrderCents

`func (o *ListBuyerCashbackItem) HasOrderCents() bool`

HasOrderCents returns a boolean if a field has been set.

### SetOrderCentsNil

`func (o *ListBuyerCashbackItem) SetOrderCentsNil(b bool)`

 SetOrderCentsNil sets the value for OrderCents to be an explicit nil

### UnsetOrderCents
`func (o *ListBuyerCashbackItem) UnsetOrderCents()`

UnsetOrderCents ensures that no value is present for OrderCents, not even an explicit nil
### GetCashbackCents

`func (o *ListBuyerCashbackItem) GetCashbackCents() float32`

GetCashbackCents returns the CashbackCents field if non-nil, zero value otherwise.

### GetCashbackCentsOk

`func (o *ListBuyerCashbackItem) GetCashbackCentsOk() (*float32, bool)`

GetCashbackCentsOk returns a tuple with the CashbackCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCashbackCents

`func (o *ListBuyerCashbackItem) SetCashbackCents(v float32)`

SetCashbackCents sets CashbackCents field to given value.

### HasCashbackCents

`func (o *ListBuyerCashbackItem) HasCashbackCents() bool`

HasCashbackCents returns a boolean if a field has been set.

### SetCashbackCentsNil

`func (o *ListBuyerCashbackItem) SetCashbackCentsNil(b bool)`

 SetCashbackCentsNil sets the value for CashbackCents to be an explicit nil

### UnsetCashbackCents
`func (o *ListBuyerCashbackItem) UnsetCashbackCents()`

UnsetCashbackCents ensures that no value is present for CashbackCents, not even an explicit nil
### GetActivatedAt

`func (o *ListBuyerCashbackItem) GetActivatedAt() time.Time`

GetActivatedAt returns the ActivatedAt field if non-nil, zero value otherwise.

### GetActivatedAtOk

`func (o *ListBuyerCashbackItem) GetActivatedAtOk() (*time.Time, bool)`

GetActivatedAtOk returns a tuple with the ActivatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivatedAt

`func (o *ListBuyerCashbackItem) SetActivatedAt(v time.Time)`

SetActivatedAt sets ActivatedAt field to given value.


### GetConfirmedAt

`func (o *ListBuyerCashbackItem) GetConfirmedAt() time.Time`

GetConfirmedAt returns the ConfirmedAt field if non-nil, zero value otherwise.

### GetConfirmedAtOk

`func (o *ListBuyerCashbackItem) GetConfirmedAtOk() (*time.Time, bool)`

GetConfirmedAtOk returns a tuple with the ConfirmedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmedAt

`func (o *ListBuyerCashbackItem) SetConfirmedAt(v time.Time)`

SetConfirmedAt sets ConfirmedAt field to given value.

### HasConfirmedAt

`func (o *ListBuyerCashbackItem) HasConfirmedAt() bool`

HasConfirmedAt returns a boolean if a field has been set.

### SetConfirmedAtNil

`func (o *ListBuyerCashbackItem) SetConfirmedAtNil(b bool)`

 SetConfirmedAtNil sets the value for ConfirmedAt to be an explicit nil

### UnsetConfirmedAt
`func (o *ListBuyerCashbackItem) UnsetConfirmedAt()`

UnsetConfirmedAt ensures that no value is present for ConfirmedAt, not even an explicit nil
### GetPaidAt

`func (o *ListBuyerCashbackItem) GetPaidAt() time.Time`

GetPaidAt returns the PaidAt field if non-nil, zero value otherwise.

### GetPaidAtOk

`func (o *ListBuyerCashbackItem) GetPaidAtOk() (*time.Time, bool)`

GetPaidAtOk returns a tuple with the PaidAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaidAt

`func (o *ListBuyerCashbackItem) SetPaidAt(v time.Time)`

SetPaidAt sets PaidAt field to given value.

### HasPaidAt

`func (o *ListBuyerCashbackItem) HasPaidAt() bool`

HasPaidAt returns a boolean if a field has been set.

### SetPaidAtNil

`func (o *ListBuyerCashbackItem) SetPaidAtNil(b bool)`

 SetPaidAtNil sets the value for PaidAt to be an explicit nil

### UnsetPaidAt
`func (o *ListBuyerCashbackItem) UnsetPaidAt()`

UnsetPaidAt ensures that no value is present for PaidAt, not even an explicit nil
### GetExpiredAt

`func (o *ListBuyerCashbackItem) GetExpiredAt() time.Time`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *ListBuyerCashbackItem) GetExpiredAtOk() (*time.Time, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *ListBuyerCashbackItem) SetExpiredAt(v time.Time)`

SetExpiredAt sets ExpiredAt field to given value.

### HasExpiredAt

`func (o *ListBuyerCashbackItem) HasExpiredAt() bool`

HasExpiredAt returns a boolean if a field has been set.

### SetExpiredAtNil

`func (o *ListBuyerCashbackItem) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *ListBuyerCashbackItem) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil
### GetRejectedReason

`func (o *ListBuyerCashbackItem) GetRejectedReason() string`

GetRejectedReason returns the RejectedReason field if non-nil, zero value otherwise.

### GetRejectedReasonOk

`func (o *ListBuyerCashbackItem) GetRejectedReasonOk() (*string, bool)`

GetRejectedReasonOk returns a tuple with the RejectedReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectedReason

`func (o *ListBuyerCashbackItem) SetRejectedReason(v string)`

SetRejectedReason sets RejectedReason field to given value.

### HasRejectedReason

`func (o *ListBuyerCashbackItem) HasRejectedReason() bool`

HasRejectedReason returns a boolean if a field has been set.

### SetRejectedReasonNil

`func (o *ListBuyerCashbackItem) SetRejectedReasonNil(b bool)`

 SetRejectedReasonNil sets the value for RejectedReason to be an explicit nil

### UnsetRejectedReason
`func (o *ListBuyerCashbackItem) UnsetRejectedReason()`

UnsetRejectedReason ensures that no value is present for RejectedReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


