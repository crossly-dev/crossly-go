# GetCbxSubjectStakeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BalanceBaseUnits** | **string** |  | 
**StakedBaseUnits** | **string** |  | 
**AvailableBaseUnits** | **string** |  | 
**TierSlug** | Pointer to **NullableString** |  | [optional] 
**FeeDiscountBps** | **float32** |  | 
**EarnBoostBps** | **float32** |  | 
**Status** | Pointer to **NullableString** |  | [optional] 
**UnlocksAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetCbxSubjectStakeResponse

`func NewGetCbxSubjectStakeResponse(balanceBaseUnits string, stakedBaseUnits string, availableBaseUnits string, feeDiscountBps float32, earnBoostBps float32, ) *GetCbxSubjectStakeResponse`

NewGetCbxSubjectStakeResponse instantiates a new GetCbxSubjectStakeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCbxSubjectStakeResponseWithDefaults

`func NewGetCbxSubjectStakeResponseWithDefaults() *GetCbxSubjectStakeResponse`

NewGetCbxSubjectStakeResponseWithDefaults instantiates a new GetCbxSubjectStakeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBalanceBaseUnits

`func (o *GetCbxSubjectStakeResponse) GetBalanceBaseUnits() string`

GetBalanceBaseUnits returns the BalanceBaseUnits field if non-nil, zero value otherwise.

### GetBalanceBaseUnitsOk

`func (o *GetCbxSubjectStakeResponse) GetBalanceBaseUnitsOk() (*string, bool)`

GetBalanceBaseUnitsOk returns a tuple with the BalanceBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalanceBaseUnits

`func (o *GetCbxSubjectStakeResponse) SetBalanceBaseUnits(v string)`

SetBalanceBaseUnits sets BalanceBaseUnits field to given value.


### GetStakedBaseUnits

`func (o *GetCbxSubjectStakeResponse) GetStakedBaseUnits() string`

GetStakedBaseUnits returns the StakedBaseUnits field if non-nil, zero value otherwise.

### GetStakedBaseUnitsOk

`func (o *GetCbxSubjectStakeResponse) GetStakedBaseUnitsOk() (*string, bool)`

GetStakedBaseUnitsOk returns a tuple with the StakedBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakedBaseUnits

`func (o *GetCbxSubjectStakeResponse) SetStakedBaseUnits(v string)`

SetStakedBaseUnits sets StakedBaseUnits field to given value.


### GetAvailableBaseUnits

`func (o *GetCbxSubjectStakeResponse) GetAvailableBaseUnits() string`

GetAvailableBaseUnits returns the AvailableBaseUnits field if non-nil, zero value otherwise.

### GetAvailableBaseUnitsOk

`func (o *GetCbxSubjectStakeResponse) GetAvailableBaseUnitsOk() (*string, bool)`

GetAvailableBaseUnitsOk returns a tuple with the AvailableBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableBaseUnits

`func (o *GetCbxSubjectStakeResponse) SetAvailableBaseUnits(v string)`

SetAvailableBaseUnits sets AvailableBaseUnits field to given value.


### GetTierSlug

`func (o *GetCbxSubjectStakeResponse) GetTierSlug() string`

GetTierSlug returns the TierSlug field if non-nil, zero value otherwise.

### GetTierSlugOk

`func (o *GetCbxSubjectStakeResponse) GetTierSlugOk() (*string, bool)`

GetTierSlugOk returns a tuple with the TierSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierSlug

`func (o *GetCbxSubjectStakeResponse) SetTierSlug(v string)`

SetTierSlug sets TierSlug field to given value.

### HasTierSlug

`func (o *GetCbxSubjectStakeResponse) HasTierSlug() bool`

HasTierSlug returns a boolean if a field has been set.

### SetTierSlugNil

`func (o *GetCbxSubjectStakeResponse) SetTierSlugNil(b bool)`

 SetTierSlugNil sets the value for TierSlug to be an explicit nil

### UnsetTierSlug
`func (o *GetCbxSubjectStakeResponse) UnsetTierSlug()`

UnsetTierSlug ensures that no value is present for TierSlug, not even an explicit nil
### GetFeeDiscountBps

`func (o *GetCbxSubjectStakeResponse) GetFeeDiscountBps() float32`

GetFeeDiscountBps returns the FeeDiscountBps field if non-nil, zero value otherwise.

### GetFeeDiscountBpsOk

`func (o *GetCbxSubjectStakeResponse) GetFeeDiscountBpsOk() (*float32, bool)`

GetFeeDiscountBpsOk returns a tuple with the FeeDiscountBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeDiscountBps

`func (o *GetCbxSubjectStakeResponse) SetFeeDiscountBps(v float32)`

SetFeeDiscountBps sets FeeDiscountBps field to given value.


### GetEarnBoostBps

`func (o *GetCbxSubjectStakeResponse) GetEarnBoostBps() float32`

GetEarnBoostBps returns the EarnBoostBps field if non-nil, zero value otherwise.

### GetEarnBoostBpsOk

`func (o *GetCbxSubjectStakeResponse) GetEarnBoostBpsOk() (*float32, bool)`

GetEarnBoostBpsOk returns a tuple with the EarnBoostBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEarnBoostBps

`func (o *GetCbxSubjectStakeResponse) SetEarnBoostBps(v float32)`

SetEarnBoostBps sets EarnBoostBps field to given value.


### GetStatus

`func (o *GetCbxSubjectStakeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCbxSubjectStakeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCbxSubjectStakeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCbxSubjectStakeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetCbxSubjectStakeResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetCbxSubjectStakeResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetUnlocksAt

`func (o *GetCbxSubjectStakeResponse) GetUnlocksAt() string`

GetUnlocksAt returns the UnlocksAt field if non-nil, zero value otherwise.

### GetUnlocksAtOk

`func (o *GetCbxSubjectStakeResponse) GetUnlocksAtOk() (*string, bool)`

GetUnlocksAtOk returns a tuple with the UnlocksAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnlocksAt

`func (o *GetCbxSubjectStakeResponse) SetUnlocksAt(v string)`

SetUnlocksAt sets UnlocksAt field to given value.

### HasUnlocksAt

`func (o *GetCbxSubjectStakeResponse) HasUnlocksAt() bool`

HasUnlocksAt returns a boolean if a field has been set.

### SetUnlocksAtNil

`func (o *GetCbxSubjectStakeResponse) SetUnlocksAtNil(b bool)`

 SetUnlocksAtNil sets the value for UnlocksAt to be an explicit nil

### UnsetUnlocksAt
`func (o *GetCbxSubjectStakeResponse) UnsetUnlocksAt()`

UnsetUnlocksAt ensures that no value is present for UnlocksAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


