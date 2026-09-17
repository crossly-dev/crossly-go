# GetPlatformLimitResponseEbaySellingCap

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmountLimit** | Pointer to **NullableFloat32** | Monthly $$ ceiling — null when seller has no $$ cap set. | [optional] 
**AmountUsed** | Pointer to **NullableFloat32** |  | [optional] 
**QuantityLimit** | Pointer to **NullableFloat32** | Monthly item-count ceiling — null when seller has no qty cap. | [optional] 
**QuantityUsed** | Pointer to **NullableFloat32** |  | [optional] 
**Currency** | **string** |  | 
**FetchedAt** | **string** | ISO timestamp this was fetched (used for cache freshness display). | 

## Methods

### NewGetPlatformLimitResponseEbaySellingCap

`func NewGetPlatformLimitResponseEbaySellingCap(currency string, fetchedAt string, ) *GetPlatformLimitResponseEbaySellingCap`

NewGetPlatformLimitResponseEbaySellingCap instantiates a new GetPlatformLimitResponseEbaySellingCap object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlatformLimitResponseEbaySellingCapWithDefaults

`func NewGetPlatformLimitResponseEbaySellingCapWithDefaults() *GetPlatformLimitResponseEbaySellingCap`

NewGetPlatformLimitResponseEbaySellingCapWithDefaults instantiates a new GetPlatformLimitResponseEbaySellingCap object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmountLimit

`func (o *GetPlatformLimitResponseEbaySellingCap) GetAmountLimit() float32`

GetAmountLimit returns the AmountLimit field if non-nil, zero value otherwise.

### GetAmountLimitOk

`func (o *GetPlatformLimitResponseEbaySellingCap) GetAmountLimitOk() (*float32, bool)`

GetAmountLimitOk returns a tuple with the AmountLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountLimit

`func (o *GetPlatformLimitResponseEbaySellingCap) SetAmountLimit(v float32)`

SetAmountLimit sets AmountLimit field to given value.

### HasAmountLimit

`func (o *GetPlatformLimitResponseEbaySellingCap) HasAmountLimit() bool`

HasAmountLimit returns a boolean if a field has been set.

### SetAmountLimitNil

`func (o *GetPlatformLimitResponseEbaySellingCap) SetAmountLimitNil(b bool)`

 SetAmountLimitNil sets the value for AmountLimit to be an explicit nil

### UnsetAmountLimit
`func (o *GetPlatformLimitResponseEbaySellingCap) UnsetAmountLimit()`

UnsetAmountLimit ensures that no value is present for AmountLimit, not even an explicit nil
### GetAmountUsed

`func (o *GetPlatformLimitResponseEbaySellingCap) GetAmountUsed() float32`

GetAmountUsed returns the AmountUsed field if non-nil, zero value otherwise.

### GetAmountUsedOk

`func (o *GetPlatformLimitResponseEbaySellingCap) GetAmountUsedOk() (*float32, bool)`

GetAmountUsedOk returns a tuple with the AmountUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountUsed

`func (o *GetPlatformLimitResponseEbaySellingCap) SetAmountUsed(v float32)`

SetAmountUsed sets AmountUsed field to given value.

### HasAmountUsed

`func (o *GetPlatformLimitResponseEbaySellingCap) HasAmountUsed() bool`

HasAmountUsed returns a boolean if a field has been set.

### SetAmountUsedNil

`func (o *GetPlatformLimitResponseEbaySellingCap) SetAmountUsedNil(b bool)`

 SetAmountUsedNil sets the value for AmountUsed to be an explicit nil

### UnsetAmountUsed
`func (o *GetPlatformLimitResponseEbaySellingCap) UnsetAmountUsed()`

UnsetAmountUsed ensures that no value is present for AmountUsed, not even an explicit nil
### GetQuantityLimit

`func (o *GetPlatformLimitResponseEbaySellingCap) GetQuantityLimit() float32`

GetQuantityLimit returns the QuantityLimit field if non-nil, zero value otherwise.

### GetQuantityLimitOk

`func (o *GetPlatformLimitResponseEbaySellingCap) GetQuantityLimitOk() (*float32, bool)`

GetQuantityLimitOk returns a tuple with the QuantityLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityLimit

`func (o *GetPlatformLimitResponseEbaySellingCap) SetQuantityLimit(v float32)`

SetQuantityLimit sets QuantityLimit field to given value.

### HasQuantityLimit

`func (o *GetPlatformLimitResponseEbaySellingCap) HasQuantityLimit() bool`

HasQuantityLimit returns a boolean if a field has been set.

### SetQuantityLimitNil

`func (o *GetPlatformLimitResponseEbaySellingCap) SetQuantityLimitNil(b bool)`

 SetQuantityLimitNil sets the value for QuantityLimit to be an explicit nil

### UnsetQuantityLimit
`func (o *GetPlatformLimitResponseEbaySellingCap) UnsetQuantityLimit()`

UnsetQuantityLimit ensures that no value is present for QuantityLimit, not even an explicit nil
### GetQuantityUsed

`func (o *GetPlatformLimitResponseEbaySellingCap) GetQuantityUsed() float32`

GetQuantityUsed returns the QuantityUsed field if non-nil, zero value otherwise.

### GetQuantityUsedOk

`func (o *GetPlatformLimitResponseEbaySellingCap) GetQuantityUsedOk() (*float32, bool)`

GetQuantityUsedOk returns a tuple with the QuantityUsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityUsed

`func (o *GetPlatformLimitResponseEbaySellingCap) SetQuantityUsed(v float32)`

SetQuantityUsed sets QuantityUsed field to given value.

### HasQuantityUsed

`func (o *GetPlatformLimitResponseEbaySellingCap) HasQuantityUsed() bool`

HasQuantityUsed returns a boolean if a field has been set.

### SetQuantityUsedNil

`func (o *GetPlatformLimitResponseEbaySellingCap) SetQuantityUsedNil(b bool)`

 SetQuantityUsedNil sets the value for QuantityUsed to be an explicit nil

### UnsetQuantityUsed
`func (o *GetPlatformLimitResponseEbaySellingCap) UnsetQuantityUsed()`

UnsetQuantityUsed ensures that no value is present for QuantityUsed, not even an explicit nil
### GetCurrency

`func (o *GetPlatformLimitResponseEbaySellingCap) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetPlatformLimitResponseEbaySellingCap) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetPlatformLimitResponseEbaySellingCap) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetFetchedAt

`func (o *GetPlatformLimitResponseEbaySellingCap) GetFetchedAt() string`

GetFetchedAt returns the FetchedAt field if non-nil, zero value otherwise.

### GetFetchedAtOk

`func (o *GetPlatformLimitResponseEbaySellingCap) GetFetchedAtOk() (*string, bool)`

GetFetchedAtOk returns a tuple with the FetchedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchedAt

`func (o *GetPlatformLimitResponseEbaySellingCap) SetFetchedAt(v string)`

SetFetchedAt sets FetchedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


