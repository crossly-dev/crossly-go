# CreateCbxAccrualPurchaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccrualId** | Pointer to **NullableString** |  | [optional] 
**Duplicate** | **bool** |  | 
**Cents** | **float32** |  | 
**RateBps** | Pointer to **NullableFloat32** |  | [optional] 
**MaturesAt** | Pointer to **NullableString** |  | [optional] 
**BoostBudgetExhausted** | **bool** |  | 

## Methods

### NewCreateCbxAccrualPurchaseResponse

`func NewCreateCbxAccrualPurchaseResponse(duplicate bool, cents float32, boostBudgetExhausted bool, ) *CreateCbxAccrualPurchaseResponse`

NewCreateCbxAccrualPurchaseResponse instantiates a new CreateCbxAccrualPurchaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCbxAccrualPurchaseResponseWithDefaults

`func NewCreateCbxAccrualPurchaseResponseWithDefaults() *CreateCbxAccrualPurchaseResponse`

NewCreateCbxAccrualPurchaseResponseWithDefaults instantiates a new CreateCbxAccrualPurchaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccrualId

`func (o *CreateCbxAccrualPurchaseResponse) GetAccrualId() string`

GetAccrualId returns the AccrualId field if non-nil, zero value otherwise.

### GetAccrualIdOk

`func (o *CreateCbxAccrualPurchaseResponse) GetAccrualIdOk() (*string, bool)`

GetAccrualIdOk returns a tuple with the AccrualId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccrualId

`func (o *CreateCbxAccrualPurchaseResponse) SetAccrualId(v string)`

SetAccrualId sets AccrualId field to given value.

### HasAccrualId

`func (o *CreateCbxAccrualPurchaseResponse) HasAccrualId() bool`

HasAccrualId returns a boolean if a field has been set.

### SetAccrualIdNil

`func (o *CreateCbxAccrualPurchaseResponse) SetAccrualIdNil(b bool)`

 SetAccrualIdNil sets the value for AccrualId to be an explicit nil

### UnsetAccrualId
`func (o *CreateCbxAccrualPurchaseResponse) UnsetAccrualId()`

UnsetAccrualId ensures that no value is present for AccrualId, not even an explicit nil
### GetDuplicate

`func (o *CreateCbxAccrualPurchaseResponse) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *CreateCbxAccrualPurchaseResponse) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *CreateCbxAccrualPurchaseResponse) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.


### GetCents

`func (o *CreateCbxAccrualPurchaseResponse) GetCents() float32`

GetCents returns the Cents field if non-nil, zero value otherwise.

### GetCentsOk

`func (o *CreateCbxAccrualPurchaseResponse) GetCentsOk() (*float32, bool)`

GetCentsOk returns a tuple with the Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCents

`func (o *CreateCbxAccrualPurchaseResponse) SetCents(v float32)`

SetCents sets Cents field to given value.


### GetRateBps

`func (o *CreateCbxAccrualPurchaseResponse) GetRateBps() float32`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *CreateCbxAccrualPurchaseResponse) GetRateBpsOk() (*float32, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *CreateCbxAccrualPurchaseResponse) SetRateBps(v float32)`

SetRateBps sets RateBps field to given value.

### HasRateBps

`func (o *CreateCbxAccrualPurchaseResponse) HasRateBps() bool`

HasRateBps returns a boolean if a field has been set.

### SetRateBpsNil

`func (o *CreateCbxAccrualPurchaseResponse) SetRateBpsNil(b bool)`

 SetRateBpsNil sets the value for RateBps to be an explicit nil

### UnsetRateBps
`func (o *CreateCbxAccrualPurchaseResponse) UnsetRateBps()`

UnsetRateBps ensures that no value is present for RateBps, not even an explicit nil
### GetMaturesAt

`func (o *CreateCbxAccrualPurchaseResponse) GetMaturesAt() string`

GetMaturesAt returns the MaturesAt field if non-nil, zero value otherwise.

### GetMaturesAtOk

`func (o *CreateCbxAccrualPurchaseResponse) GetMaturesAtOk() (*string, bool)`

GetMaturesAtOk returns a tuple with the MaturesAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturesAt

`func (o *CreateCbxAccrualPurchaseResponse) SetMaturesAt(v string)`

SetMaturesAt sets MaturesAt field to given value.

### HasMaturesAt

`func (o *CreateCbxAccrualPurchaseResponse) HasMaturesAt() bool`

HasMaturesAt returns a boolean if a field has been set.

### SetMaturesAtNil

`func (o *CreateCbxAccrualPurchaseResponse) SetMaturesAtNil(b bool)`

 SetMaturesAtNil sets the value for MaturesAt to be an explicit nil

### UnsetMaturesAt
`func (o *CreateCbxAccrualPurchaseResponse) UnsetMaturesAt()`

UnsetMaturesAt ensures that no value is present for MaturesAt, not even an explicit nil
### GetBoostBudgetExhausted

`func (o *CreateCbxAccrualPurchaseResponse) GetBoostBudgetExhausted() bool`

GetBoostBudgetExhausted returns the BoostBudgetExhausted field if non-nil, zero value otherwise.

### GetBoostBudgetExhaustedOk

`func (o *CreateCbxAccrualPurchaseResponse) GetBoostBudgetExhaustedOk() (*bool, bool)`

GetBoostBudgetExhaustedOk returns a tuple with the BoostBudgetExhausted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostBudgetExhausted

`func (o *CreateCbxAccrualPurchaseResponse) SetBoostBudgetExhausted(v bool)`

SetBoostBudgetExhausted sets BoostBudgetExhausted field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


