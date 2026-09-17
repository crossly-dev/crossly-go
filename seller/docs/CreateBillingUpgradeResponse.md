# CreateBillingUpgradeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckoutUrl** | Pointer to **NullableString** |  | [optional] 
**Tier** | **string** |  | 
**Interval** | **string** |  | 
**Charged** | **bool** |  | 

## Methods

### NewCreateBillingUpgradeResponse

`func NewCreateBillingUpgradeResponse(tier string, interval string, charged bool, ) *CreateBillingUpgradeResponse`

NewCreateBillingUpgradeResponse instantiates a new CreateBillingUpgradeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBillingUpgradeResponseWithDefaults

`func NewCreateBillingUpgradeResponseWithDefaults() *CreateBillingUpgradeResponse`

NewCreateBillingUpgradeResponseWithDefaults instantiates a new CreateBillingUpgradeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckoutUrl

`func (o *CreateBillingUpgradeResponse) GetCheckoutUrl() string`

GetCheckoutUrl returns the CheckoutUrl field if non-nil, zero value otherwise.

### GetCheckoutUrlOk

`func (o *CreateBillingUpgradeResponse) GetCheckoutUrlOk() (*string, bool)`

GetCheckoutUrlOk returns a tuple with the CheckoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckoutUrl

`func (o *CreateBillingUpgradeResponse) SetCheckoutUrl(v string)`

SetCheckoutUrl sets CheckoutUrl field to given value.

### HasCheckoutUrl

`func (o *CreateBillingUpgradeResponse) HasCheckoutUrl() bool`

HasCheckoutUrl returns a boolean if a field has been set.

### SetCheckoutUrlNil

`func (o *CreateBillingUpgradeResponse) SetCheckoutUrlNil(b bool)`

 SetCheckoutUrlNil sets the value for CheckoutUrl to be an explicit nil

### UnsetCheckoutUrl
`func (o *CreateBillingUpgradeResponse) UnsetCheckoutUrl()`

UnsetCheckoutUrl ensures that no value is present for CheckoutUrl, not even an explicit nil
### GetTier

`func (o *CreateBillingUpgradeResponse) GetTier() string`

GetTier returns the Tier field if non-nil, zero value otherwise.

### GetTierOk

`func (o *CreateBillingUpgradeResponse) GetTierOk() (*string, bool)`

GetTierOk returns a tuple with the Tier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier

`func (o *CreateBillingUpgradeResponse) SetTier(v string)`

SetTier sets Tier field to given value.


### GetInterval

`func (o *CreateBillingUpgradeResponse) GetInterval() string`

GetInterval returns the Interval field if non-nil, zero value otherwise.

### GetIntervalOk

`func (o *CreateBillingUpgradeResponse) GetIntervalOk() (*string, bool)`

GetIntervalOk returns a tuple with the Interval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterval

`func (o *CreateBillingUpgradeResponse) SetInterval(v string)`

SetInterval sets Interval field to given value.


### GetCharged

`func (o *CreateBillingUpgradeResponse) GetCharged() bool`

GetCharged returns the Charged field if non-nil, zero value otherwise.

### GetChargedOk

`func (o *CreateBillingUpgradeResponse) GetChargedOk() (*bool, bool)`

GetChargedOk returns a tuple with the Charged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharged

`func (o *CreateBillingUpgradeResponse) SetCharged(v bool)`

SetCharged sets Charged field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


