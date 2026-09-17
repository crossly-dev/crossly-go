# GetBuyerAnywhereResponseOffsite

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StoreId** | **string** |  | 
**Host** | **string** | The retailer&#39;s hostname, e.g. &#x60;rei.com&#x60;. | 
**StoreName** | **string** |  | 
**Title** | **string** |  | 
**PriceCents** | **float32** |  | 
**ShippingCents** | Pointer to **NullableFloat32** | Null &#x3D; UNKNOWN, never free. | [optional] 
**Currency** | **string** |  | 
**Condition** | Pointer to **NullableString** |  | [optional] 
**Url** | **string** |  | 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**BuyerCashbackCents** | **float32** | What the buyer gets back, in cents, if they buy through us.  Shown because a cashback figure the buyer cannot see is a figure they have no reason to believe. Derived from the store&#39;s rate, never stored per offer — rates change and a copied one goes stale silently. | 
**DeliveredCents** | **float32** | Item + shipping when known; item alone otherwise. See &#x60;shippingUnknown&#x60;. | 
**ShippingUnknown** | **bool** |  | 

## Methods

### NewGetBuyerAnywhereResponseOffsite

`func NewGetBuyerAnywhereResponseOffsite(storeId string, host string, storeName string, title string, priceCents float32, currency string, url string, buyerCashbackCents float32, deliveredCents float32, shippingUnknown bool, ) *GetBuyerAnywhereResponseOffsite`

NewGetBuyerAnywhereResponseOffsite instantiates a new GetBuyerAnywhereResponseOffsite object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBuyerAnywhereResponseOffsiteWithDefaults

`func NewGetBuyerAnywhereResponseOffsiteWithDefaults() *GetBuyerAnywhereResponseOffsite`

NewGetBuyerAnywhereResponseOffsiteWithDefaults instantiates a new GetBuyerAnywhereResponseOffsite object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreId

`func (o *GetBuyerAnywhereResponseOffsite) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *GetBuyerAnywhereResponseOffsite) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *GetBuyerAnywhereResponseOffsite) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.


### GetHost

`func (o *GetBuyerAnywhereResponseOffsite) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetBuyerAnywhereResponseOffsite) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetBuyerAnywhereResponseOffsite) SetHost(v string)`

SetHost sets Host field to given value.


### GetStoreName

`func (o *GetBuyerAnywhereResponseOffsite) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *GetBuyerAnywhereResponseOffsite) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *GetBuyerAnywhereResponseOffsite) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.


### GetTitle

`func (o *GetBuyerAnywhereResponseOffsite) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetBuyerAnywhereResponseOffsite) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetBuyerAnywhereResponseOffsite) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPriceCents

`func (o *GetBuyerAnywhereResponseOffsite) GetPriceCents() float32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *GetBuyerAnywhereResponseOffsite) GetPriceCentsOk() (*float32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *GetBuyerAnywhereResponseOffsite) SetPriceCents(v float32)`

SetPriceCents sets PriceCents field to given value.


### GetShippingCents

`func (o *GetBuyerAnywhereResponseOffsite) GetShippingCents() float32`

GetShippingCents returns the ShippingCents field if non-nil, zero value otherwise.

### GetShippingCentsOk

`func (o *GetBuyerAnywhereResponseOffsite) GetShippingCentsOk() (*float32, bool)`

GetShippingCentsOk returns a tuple with the ShippingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCents

`func (o *GetBuyerAnywhereResponseOffsite) SetShippingCents(v float32)`

SetShippingCents sets ShippingCents field to given value.

### HasShippingCents

`func (o *GetBuyerAnywhereResponseOffsite) HasShippingCents() bool`

HasShippingCents returns a boolean if a field has been set.

### SetShippingCentsNil

`func (o *GetBuyerAnywhereResponseOffsite) SetShippingCentsNil(b bool)`

 SetShippingCentsNil sets the value for ShippingCents to be an explicit nil

### UnsetShippingCents
`func (o *GetBuyerAnywhereResponseOffsite) UnsetShippingCents()`

UnsetShippingCents ensures that no value is present for ShippingCents, not even an explicit nil
### GetCurrency

`func (o *GetBuyerAnywhereResponseOffsite) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetBuyerAnywhereResponseOffsite) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetBuyerAnywhereResponseOffsite) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCondition

`func (o *GetBuyerAnywhereResponseOffsite) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GetBuyerAnywhereResponseOffsite) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *GetBuyerAnywhereResponseOffsite) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *GetBuyerAnywhereResponseOffsite) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *GetBuyerAnywhereResponseOffsite) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *GetBuyerAnywhereResponseOffsite) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetUrl

`func (o *GetBuyerAnywhereResponseOffsite) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetBuyerAnywhereResponseOffsite) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetBuyerAnywhereResponseOffsite) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetImageUrl

`func (o *GetBuyerAnywhereResponseOffsite) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GetBuyerAnywhereResponseOffsite) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GetBuyerAnywhereResponseOffsite) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GetBuyerAnywhereResponseOffsite) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *GetBuyerAnywhereResponseOffsite) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *GetBuyerAnywhereResponseOffsite) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetBuyerCashbackCents

`func (o *GetBuyerAnywhereResponseOffsite) GetBuyerCashbackCents() float32`

GetBuyerCashbackCents returns the BuyerCashbackCents field if non-nil, zero value otherwise.

### GetBuyerCashbackCentsOk

`func (o *GetBuyerAnywhereResponseOffsite) GetBuyerCashbackCentsOk() (*float32, bool)`

GetBuyerCashbackCentsOk returns a tuple with the BuyerCashbackCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerCashbackCents

`func (o *GetBuyerAnywhereResponseOffsite) SetBuyerCashbackCents(v float32)`

SetBuyerCashbackCents sets BuyerCashbackCents field to given value.


### GetDeliveredCents

`func (o *GetBuyerAnywhereResponseOffsite) GetDeliveredCents() float32`

GetDeliveredCents returns the DeliveredCents field if non-nil, zero value otherwise.

### GetDeliveredCentsOk

`func (o *GetBuyerAnywhereResponseOffsite) GetDeliveredCentsOk() (*float32, bool)`

GetDeliveredCentsOk returns a tuple with the DeliveredCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredCents

`func (o *GetBuyerAnywhereResponseOffsite) SetDeliveredCents(v float32)`

SetDeliveredCents sets DeliveredCents field to given value.


### GetShippingUnknown

`func (o *GetBuyerAnywhereResponseOffsite) GetShippingUnknown() bool`

GetShippingUnknown returns the ShippingUnknown field if non-nil, zero value otherwise.

### GetShippingUnknownOk

`func (o *GetBuyerAnywhereResponseOffsite) GetShippingUnknownOk() (*bool, bool)`

GetShippingUnknownOk returns a tuple with the ShippingUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingUnknown

`func (o *GetBuyerAnywhereResponseOffsite) SetShippingUnknown(v bool)`

SetShippingUnknown sets ShippingUnknown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


