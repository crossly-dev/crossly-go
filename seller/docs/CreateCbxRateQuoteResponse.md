# CreateCbxRateQuoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RateBps** | **float32** |  | 
**Cents** | **float32** |  | 
**MaturationDays** | **float32** |  | 
**Source** | **string** |  | 
**TierSlug** | Pointer to **NullableString** |  | [optional] 
**BoostName** | Pointer to **NullableString** |  | [optional] 
**StakeBoostBps** | **float32** |  | 
**Clamped** | **bool** |  | 

## Methods

### NewCreateCbxRateQuoteResponse

`func NewCreateCbxRateQuoteResponse(rateBps float32, cents float32, maturationDays float32, source string, stakeBoostBps float32, clamped bool, ) *CreateCbxRateQuoteResponse`

NewCreateCbxRateQuoteResponse instantiates a new CreateCbxRateQuoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCbxRateQuoteResponseWithDefaults

`func NewCreateCbxRateQuoteResponseWithDefaults() *CreateCbxRateQuoteResponse`

NewCreateCbxRateQuoteResponseWithDefaults instantiates a new CreateCbxRateQuoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRateBps

`func (o *CreateCbxRateQuoteResponse) GetRateBps() float32`

GetRateBps returns the RateBps field if non-nil, zero value otherwise.

### GetRateBpsOk

`func (o *CreateCbxRateQuoteResponse) GetRateBpsOk() (*float32, bool)`

GetRateBpsOk returns a tuple with the RateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateBps

`func (o *CreateCbxRateQuoteResponse) SetRateBps(v float32)`

SetRateBps sets RateBps field to given value.


### GetCents

`func (o *CreateCbxRateQuoteResponse) GetCents() float32`

GetCents returns the Cents field if non-nil, zero value otherwise.

### GetCentsOk

`func (o *CreateCbxRateQuoteResponse) GetCentsOk() (*float32, bool)`

GetCentsOk returns a tuple with the Cents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCents

`func (o *CreateCbxRateQuoteResponse) SetCents(v float32)`

SetCents sets Cents field to given value.


### GetMaturationDays

`func (o *CreateCbxRateQuoteResponse) GetMaturationDays() float32`

GetMaturationDays returns the MaturationDays field if non-nil, zero value otherwise.

### GetMaturationDaysOk

`func (o *CreateCbxRateQuoteResponse) GetMaturationDaysOk() (*float32, bool)`

GetMaturationDaysOk returns a tuple with the MaturationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaturationDays

`func (o *CreateCbxRateQuoteResponse) SetMaturationDays(v float32)`

SetMaturationDays sets MaturationDays field to given value.


### GetSource

`func (o *CreateCbxRateQuoteResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateCbxRateQuoteResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateCbxRateQuoteResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetTierSlug

`func (o *CreateCbxRateQuoteResponse) GetTierSlug() string`

GetTierSlug returns the TierSlug field if non-nil, zero value otherwise.

### GetTierSlugOk

`func (o *CreateCbxRateQuoteResponse) GetTierSlugOk() (*string, bool)`

GetTierSlugOk returns a tuple with the TierSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierSlug

`func (o *CreateCbxRateQuoteResponse) SetTierSlug(v string)`

SetTierSlug sets TierSlug field to given value.

### HasTierSlug

`func (o *CreateCbxRateQuoteResponse) HasTierSlug() bool`

HasTierSlug returns a boolean if a field has been set.

### SetTierSlugNil

`func (o *CreateCbxRateQuoteResponse) SetTierSlugNil(b bool)`

 SetTierSlugNil sets the value for TierSlug to be an explicit nil

### UnsetTierSlug
`func (o *CreateCbxRateQuoteResponse) UnsetTierSlug()`

UnsetTierSlug ensures that no value is present for TierSlug, not even an explicit nil
### GetBoostName

`func (o *CreateCbxRateQuoteResponse) GetBoostName() string`

GetBoostName returns the BoostName field if non-nil, zero value otherwise.

### GetBoostNameOk

`func (o *CreateCbxRateQuoteResponse) GetBoostNameOk() (*string, bool)`

GetBoostNameOk returns a tuple with the BoostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostName

`func (o *CreateCbxRateQuoteResponse) SetBoostName(v string)`

SetBoostName sets BoostName field to given value.

### HasBoostName

`func (o *CreateCbxRateQuoteResponse) HasBoostName() bool`

HasBoostName returns a boolean if a field has been set.

### SetBoostNameNil

`func (o *CreateCbxRateQuoteResponse) SetBoostNameNil(b bool)`

 SetBoostNameNil sets the value for BoostName to be an explicit nil

### UnsetBoostName
`func (o *CreateCbxRateQuoteResponse) UnsetBoostName()`

UnsetBoostName ensures that no value is present for BoostName, not even an explicit nil
### GetStakeBoostBps

`func (o *CreateCbxRateQuoteResponse) GetStakeBoostBps() float32`

GetStakeBoostBps returns the StakeBoostBps field if non-nil, zero value otherwise.

### GetStakeBoostBpsOk

`func (o *CreateCbxRateQuoteResponse) GetStakeBoostBpsOk() (*float32, bool)`

GetStakeBoostBpsOk returns a tuple with the StakeBoostBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStakeBoostBps

`func (o *CreateCbxRateQuoteResponse) SetStakeBoostBps(v float32)`

SetStakeBoostBps sets StakeBoostBps field to given value.


### GetClamped

`func (o *CreateCbxRateQuoteResponse) GetClamped() bool`

GetClamped returns the Clamped field if non-nil, zero value otherwise.

### GetClampedOk

`func (o *CreateCbxRateQuoteResponse) GetClampedOk() (*bool, bool)`

GetClampedOk returns a tuple with the Clamped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClamped

`func (o *CreateCbxRateQuoteResponse) SetClamped(v bool)`

SetClamped sets Clamped field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


