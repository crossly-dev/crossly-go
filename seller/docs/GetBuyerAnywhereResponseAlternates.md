# GetBuyerAnywhereResponseAlternates

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

### NewGetBuyerAnywhereResponseAlternates

`func NewGetBuyerAnywhereResponseAlternates(storeId string, host string, storeName string, title string, priceCents float32, currency string, url string, buyerCashbackCents float32, deliveredCents float32, shippingUnknown bool, ) *GetBuyerAnywhereResponseAlternates`

NewGetBuyerAnywhereResponseAlternates instantiates a new GetBuyerAnywhereResponseAlternates object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBuyerAnywhereResponseAlternatesWithDefaults

`func NewGetBuyerAnywhereResponseAlternatesWithDefaults() *GetBuyerAnywhereResponseAlternates`

NewGetBuyerAnywhereResponseAlternatesWithDefaults instantiates a new GetBuyerAnywhereResponseAlternates object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStoreId

`func (o *GetBuyerAnywhereResponseAlternates) GetStoreId() string`

GetStoreId returns the StoreId field if non-nil, zero value otherwise.

### GetStoreIdOk

`func (o *GetBuyerAnywhereResponseAlternates) GetStoreIdOk() (*string, bool)`

GetStoreIdOk returns a tuple with the StoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreId

`func (o *GetBuyerAnywhereResponseAlternates) SetStoreId(v string)`

SetStoreId sets StoreId field to given value.


### GetHost

`func (o *GetBuyerAnywhereResponseAlternates) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *GetBuyerAnywhereResponseAlternates) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *GetBuyerAnywhereResponseAlternates) SetHost(v string)`

SetHost sets Host field to given value.


### GetStoreName

`func (o *GetBuyerAnywhereResponseAlternates) GetStoreName() string`

GetStoreName returns the StoreName field if non-nil, zero value otherwise.

### GetStoreNameOk

`func (o *GetBuyerAnywhereResponseAlternates) GetStoreNameOk() (*string, bool)`

GetStoreNameOk returns a tuple with the StoreName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreName

`func (o *GetBuyerAnywhereResponseAlternates) SetStoreName(v string)`

SetStoreName sets StoreName field to given value.


### GetTitle

`func (o *GetBuyerAnywhereResponseAlternates) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetBuyerAnywhereResponseAlternates) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetBuyerAnywhereResponseAlternates) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetPriceCents

`func (o *GetBuyerAnywhereResponseAlternates) GetPriceCents() float32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *GetBuyerAnywhereResponseAlternates) GetPriceCentsOk() (*float32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *GetBuyerAnywhereResponseAlternates) SetPriceCents(v float32)`

SetPriceCents sets PriceCents field to given value.


### GetShippingCents

`func (o *GetBuyerAnywhereResponseAlternates) GetShippingCents() float32`

GetShippingCents returns the ShippingCents field if non-nil, zero value otherwise.

### GetShippingCentsOk

`func (o *GetBuyerAnywhereResponseAlternates) GetShippingCentsOk() (*float32, bool)`

GetShippingCentsOk returns a tuple with the ShippingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCents

`func (o *GetBuyerAnywhereResponseAlternates) SetShippingCents(v float32)`

SetShippingCents sets ShippingCents field to given value.

### HasShippingCents

`func (o *GetBuyerAnywhereResponseAlternates) HasShippingCents() bool`

HasShippingCents returns a boolean if a field has been set.

### SetShippingCentsNil

`func (o *GetBuyerAnywhereResponseAlternates) SetShippingCentsNil(b bool)`

 SetShippingCentsNil sets the value for ShippingCents to be an explicit nil

### UnsetShippingCents
`func (o *GetBuyerAnywhereResponseAlternates) UnsetShippingCents()`

UnsetShippingCents ensures that no value is present for ShippingCents, not even an explicit nil
### GetCurrency

`func (o *GetBuyerAnywhereResponseAlternates) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetBuyerAnywhereResponseAlternates) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetBuyerAnywhereResponseAlternates) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetCondition

`func (o *GetBuyerAnywhereResponseAlternates) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GetBuyerAnywhereResponseAlternates) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *GetBuyerAnywhereResponseAlternates) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *GetBuyerAnywhereResponseAlternates) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *GetBuyerAnywhereResponseAlternates) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *GetBuyerAnywhereResponseAlternates) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetUrl

`func (o *GetBuyerAnywhereResponseAlternates) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetBuyerAnywhereResponseAlternates) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetBuyerAnywhereResponseAlternates) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetImageUrl

`func (o *GetBuyerAnywhereResponseAlternates) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GetBuyerAnywhereResponseAlternates) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GetBuyerAnywhereResponseAlternates) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GetBuyerAnywhereResponseAlternates) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *GetBuyerAnywhereResponseAlternates) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *GetBuyerAnywhereResponseAlternates) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetBuyerCashbackCents

`func (o *GetBuyerAnywhereResponseAlternates) GetBuyerCashbackCents() float32`

GetBuyerCashbackCents returns the BuyerCashbackCents field if non-nil, zero value otherwise.

### GetBuyerCashbackCentsOk

`func (o *GetBuyerAnywhereResponseAlternates) GetBuyerCashbackCentsOk() (*float32, bool)`

GetBuyerCashbackCentsOk returns a tuple with the BuyerCashbackCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerCashbackCents

`func (o *GetBuyerAnywhereResponseAlternates) SetBuyerCashbackCents(v float32)`

SetBuyerCashbackCents sets BuyerCashbackCents field to given value.


### GetDeliveredCents

`func (o *GetBuyerAnywhereResponseAlternates) GetDeliveredCents() float32`

GetDeliveredCents returns the DeliveredCents field if non-nil, zero value otherwise.

### GetDeliveredCentsOk

`func (o *GetBuyerAnywhereResponseAlternates) GetDeliveredCentsOk() (*float32, bool)`

GetDeliveredCentsOk returns a tuple with the DeliveredCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredCents

`func (o *GetBuyerAnywhereResponseAlternates) SetDeliveredCents(v float32)`

SetDeliveredCents sets DeliveredCents field to given value.


### GetShippingUnknown

`func (o *GetBuyerAnywhereResponseAlternates) GetShippingUnknown() bool`

GetShippingUnknown returns the ShippingUnknown field if non-nil, zero value otherwise.

### GetShippingUnknownOk

`func (o *GetBuyerAnywhereResponseAlternates) GetShippingUnknownOk() (*bool, bool)`

GetShippingUnknownOk returns a tuple with the ShippingUnknown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingUnknown

`func (o *GetBuyerAnywhereResponseAlternates) SetShippingUnknown(v bool)`

SetShippingUnknown sets ShippingUnknown field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


