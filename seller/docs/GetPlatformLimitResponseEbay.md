# GetPlatformLimitResponseEbay

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** |  | 
**Used** | **float32** |  | 
**Limit** | **float32** |  | 
**Remaining** | **float32** |  | 
**TierConfigured** | **bool** | False if the user hasn&#39;t picked a tier (we default to 250 but flag it so the UI can prompt). | 
**RespectQuota** | **bool** |  | 
**PeriodStart** | **string** |  | 
**PerOverageFeeUsd** | **float32** | Approximate cost if &#x60;used&#x60; overflows &#x60;limit&#x60; — informational. | 
**SellingCap** | Pointer to [**NullableGetPlatformLimitResponseEbaySellingCap**](GetPlatformLimitResponseEbaySellingCap.md) |  | [optional] 

## Methods

### NewGetPlatformLimitResponseEbay

`func NewGetPlatformLimitResponseEbay(platform string, used float32, limit float32, remaining float32, tierConfigured bool, respectQuota bool, periodStart string, perOverageFeeUsd float32, ) *GetPlatformLimitResponseEbay`

NewGetPlatformLimitResponseEbay instantiates a new GetPlatformLimitResponseEbay object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlatformLimitResponseEbayWithDefaults

`func NewGetPlatformLimitResponseEbayWithDefaults() *GetPlatformLimitResponseEbay`

NewGetPlatformLimitResponseEbayWithDefaults instantiates a new GetPlatformLimitResponseEbay object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *GetPlatformLimitResponseEbay) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetPlatformLimitResponseEbay) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetPlatformLimitResponseEbay) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetUsed

`func (o *GetPlatformLimitResponseEbay) GetUsed() float32`

GetUsed returns the Used field if non-nil, zero value otherwise.

### GetUsedOk

`func (o *GetPlatformLimitResponseEbay) GetUsedOk() (*float32, bool)`

GetUsedOk returns a tuple with the Used field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsed

`func (o *GetPlatformLimitResponseEbay) SetUsed(v float32)`

SetUsed sets Used field to given value.


### GetLimit

`func (o *GetPlatformLimitResponseEbay) GetLimit() float32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *GetPlatformLimitResponseEbay) GetLimitOk() (*float32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *GetPlatformLimitResponseEbay) SetLimit(v float32)`

SetLimit sets Limit field to given value.


### GetRemaining

`func (o *GetPlatformLimitResponseEbay) GetRemaining() float32`

GetRemaining returns the Remaining field if non-nil, zero value otherwise.

### GetRemainingOk

`func (o *GetPlatformLimitResponseEbay) GetRemainingOk() (*float32, bool)`

GetRemainingOk returns a tuple with the Remaining field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemaining

`func (o *GetPlatformLimitResponseEbay) SetRemaining(v float32)`

SetRemaining sets Remaining field to given value.


### GetTierConfigured

`func (o *GetPlatformLimitResponseEbay) GetTierConfigured() bool`

GetTierConfigured returns the TierConfigured field if non-nil, zero value otherwise.

### GetTierConfiguredOk

`func (o *GetPlatformLimitResponseEbay) GetTierConfiguredOk() (*bool, bool)`

GetTierConfiguredOk returns a tuple with the TierConfigured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierConfigured

`func (o *GetPlatformLimitResponseEbay) SetTierConfigured(v bool)`

SetTierConfigured sets TierConfigured field to given value.


### GetRespectQuota

`func (o *GetPlatformLimitResponseEbay) GetRespectQuota() bool`

GetRespectQuota returns the RespectQuota field if non-nil, zero value otherwise.

### GetRespectQuotaOk

`func (o *GetPlatformLimitResponseEbay) GetRespectQuotaOk() (*bool, bool)`

GetRespectQuotaOk returns a tuple with the RespectQuota field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRespectQuota

`func (o *GetPlatformLimitResponseEbay) SetRespectQuota(v bool)`

SetRespectQuota sets RespectQuota field to given value.


### GetPeriodStart

`func (o *GetPlatformLimitResponseEbay) GetPeriodStart() string`

GetPeriodStart returns the PeriodStart field if non-nil, zero value otherwise.

### GetPeriodStartOk

`func (o *GetPlatformLimitResponseEbay) GetPeriodStartOk() (*string, bool)`

GetPeriodStartOk returns a tuple with the PeriodStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriodStart

`func (o *GetPlatformLimitResponseEbay) SetPeriodStart(v string)`

SetPeriodStart sets PeriodStart field to given value.


### GetPerOverageFeeUsd

`func (o *GetPlatformLimitResponseEbay) GetPerOverageFeeUsd() float32`

GetPerOverageFeeUsd returns the PerOverageFeeUsd field if non-nil, zero value otherwise.

### GetPerOverageFeeUsdOk

`func (o *GetPlatformLimitResponseEbay) GetPerOverageFeeUsdOk() (*float32, bool)`

GetPerOverageFeeUsdOk returns a tuple with the PerOverageFeeUsd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerOverageFeeUsd

`func (o *GetPlatformLimitResponseEbay) SetPerOverageFeeUsd(v float32)`

SetPerOverageFeeUsd sets PerOverageFeeUsd field to given value.


### GetSellingCap

`func (o *GetPlatformLimitResponseEbay) GetSellingCap() GetPlatformLimitResponseEbaySellingCap`

GetSellingCap returns the SellingCap field if non-nil, zero value otherwise.

### GetSellingCapOk

`func (o *GetPlatformLimitResponseEbay) GetSellingCapOk() (*GetPlatformLimitResponseEbaySellingCap, bool)`

GetSellingCapOk returns a tuple with the SellingCap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellingCap

`func (o *GetPlatformLimitResponseEbay) SetSellingCap(v GetPlatformLimitResponseEbaySellingCap)`

SetSellingCap sets SellingCap field to given value.

### HasSellingCap

`func (o *GetPlatformLimitResponseEbay) HasSellingCap() bool`

HasSellingCap returns a boolean if a field has been set.

### SetSellingCapNil

`func (o *GetPlatformLimitResponseEbay) SetSellingCapNil(b bool)`

 SetSellingCapNil sets the value for SellingCap to be an explicit nil

### UnsetSellingCap
`func (o *GetPlatformLimitResponseEbay) UnsetSellingCap()`

UnsetSellingCap ensures that no value is present for SellingCap, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


