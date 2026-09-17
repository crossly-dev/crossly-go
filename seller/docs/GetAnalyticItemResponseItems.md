# GetAnalyticItemResponseItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SaleId** | **string** |  | 
**Platform** | **string** |  | 
**PlatformOrderId** | Pointer to **NullableString** |  | [optional] 
**ItemTitle** | Pointer to **NullableString** |  | [optional] 
**ItemBrand** | Pointer to **NullableString** |  | [optional] 
**ItemSku** | Pointer to **NullableString** |  | [optional] 
**SalePrice** | Pointer to **NullableFloat32** |  | [optional] 
**PlatformFee** | Pointer to **NullableFloat32** |  | [optional] 
**ShippingCost** | Pointer to **NullableFloat32** |  | [optional] 
**CostOfGoods** | Pointer to **NullableFloat32** |  | [optional] 
**NetProfit** | Pointer to **NullableFloat32** |  | [optional] 
**MarginPct** | **float32** |  | 
**DaysToSell** | Pointer to **NullableFloat32** |  | [optional] 
**DetectedAt** | **time.Time** |  | 

## Methods

### NewGetAnalyticItemResponseItems

`func NewGetAnalyticItemResponseItems(saleId string, platform string, marginPct float32, detectedAt time.Time, ) *GetAnalyticItemResponseItems`

NewGetAnalyticItemResponseItems instantiates a new GetAnalyticItemResponseItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalyticItemResponseItemsWithDefaults

`func NewGetAnalyticItemResponseItemsWithDefaults() *GetAnalyticItemResponseItems`

NewGetAnalyticItemResponseItemsWithDefaults instantiates a new GetAnalyticItemResponseItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaleId

`func (o *GetAnalyticItemResponseItems) GetSaleId() string`

GetSaleId returns the SaleId field if non-nil, zero value otherwise.

### GetSaleIdOk

`func (o *GetAnalyticItemResponseItems) GetSaleIdOk() (*string, bool)`

GetSaleIdOk returns a tuple with the SaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleId

`func (o *GetAnalyticItemResponseItems) SetSaleId(v string)`

SetSaleId sets SaleId field to given value.


### GetPlatform

`func (o *GetAnalyticItemResponseItems) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetAnalyticItemResponseItems) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetAnalyticItemResponseItems) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPlatformOrderId

`func (o *GetAnalyticItemResponseItems) GetPlatformOrderId() string`

GetPlatformOrderId returns the PlatformOrderId field if non-nil, zero value otherwise.

### GetPlatformOrderIdOk

`func (o *GetAnalyticItemResponseItems) GetPlatformOrderIdOk() (*string, bool)`

GetPlatformOrderIdOk returns a tuple with the PlatformOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformOrderId

`func (o *GetAnalyticItemResponseItems) SetPlatformOrderId(v string)`

SetPlatformOrderId sets PlatformOrderId field to given value.

### HasPlatformOrderId

`func (o *GetAnalyticItemResponseItems) HasPlatformOrderId() bool`

HasPlatformOrderId returns a boolean if a field has been set.

### SetPlatformOrderIdNil

`func (o *GetAnalyticItemResponseItems) SetPlatformOrderIdNil(b bool)`

 SetPlatformOrderIdNil sets the value for PlatformOrderId to be an explicit nil

### UnsetPlatformOrderId
`func (o *GetAnalyticItemResponseItems) UnsetPlatformOrderId()`

UnsetPlatformOrderId ensures that no value is present for PlatformOrderId, not even an explicit nil
### GetItemTitle

`func (o *GetAnalyticItemResponseItems) GetItemTitle() string`

GetItemTitle returns the ItemTitle field if non-nil, zero value otherwise.

### GetItemTitleOk

`func (o *GetAnalyticItemResponseItems) GetItemTitleOk() (*string, bool)`

GetItemTitleOk returns a tuple with the ItemTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTitle

`func (o *GetAnalyticItemResponseItems) SetItemTitle(v string)`

SetItemTitle sets ItemTitle field to given value.

### HasItemTitle

`func (o *GetAnalyticItemResponseItems) HasItemTitle() bool`

HasItemTitle returns a boolean if a field has been set.

### SetItemTitleNil

`func (o *GetAnalyticItemResponseItems) SetItemTitleNil(b bool)`

 SetItemTitleNil sets the value for ItemTitle to be an explicit nil

### UnsetItemTitle
`func (o *GetAnalyticItemResponseItems) UnsetItemTitle()`

UnsetItemTitle ensures that no value is present for ItemTitle, not even an explicit nil
### GetItemBrand

`func (o *GetAnalyticItemResponseItems) GetItemBrand() string`

GetItemBrand returns the ItemBrand field if non-nil, zero value otherwise.

### GetItemBrandOk

`func (o *GetAnalyticItemResponseItems) GetItemBrandOk() (*string, bool)`

GetItemBrandOk returns a tuple with the ItemBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemBrand

`func (o *GetAnalyticItemResponseItems) SetItemBrand(v string)`

SetItemBrand sets ItemBrand field to given value.

### HasItemBrand

`func (o *GetAnalyticItemResponseItems) HasItemBrand() bool`

HasItemBrand returns a boolean if a field has been set.

### SetItemBrandNil

`func (o *GetAnalyticItemResponseItems) SetItemBrandNil(b bool)`

 SetItemBrandNil sets the value for ItemBrand to be an explicit nil

### UnsetItemBrand
`func (o *GetAnalyticItemResponseItems) UnsetItemBrand()`

UnsetItemBrand ensures that no value is present for ItemBrand, not even an explicit nil
### GetItemSku

`func (o *GetAnalyticItemResponseItems) GetItemSku() string`

GetItemSku returns the ItemSku field if non-nil, zero value otherwise.

### GetItemSkuOk

`func (o *GetAnalyticItemResponseItems) GetItemSkuOk() (*string, bool)`

GetItemSkuOk returns a tuple with the ItemSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSku

`func (o *GetAnalyticItemResponseItems) SetItemSku(v string)`

SetItemSku sets ItemSku field to given value.

### HasItemSku

`func (o *GetAnalyticItemResponseItems) HasItemSku() bool`

HasItemSku returns a boolean if a field has been set.

### SetItemSkuNil

`func (o *GetAnalyticItemResponseItems) SetItemSkuNil(b bool)`

 SetItemSkuNil sets the value for ItemSku to be an explicit nil

### UnsetItemSku
`func (o *GetAnalyticItemResponseItems) UnsetItemSku()`

UnsetItemSku ensures that no value is present for ItemSku, not even an explicit nil
### GetSalePrice

`func (o *GetAnalyticItemResponseItems) GetSalePrice() float32`

GetSalePrice returns the SalePrice field if non-nil, zero value otherwise.

### GetSalePriceOk

`func (o *GetAnalyticItemResponseItems) GetSalePriceOk() (*float32, bool)`

GetSalePriceOk returns a tuple with the SalePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalePrice

`func (o *GetAnalyticItemResponseItems) SetSalePrice(v float32)`

SetSalePrice sets SalePrice field to given value.

### HasSalePrice

`func (o *GetAnalyticItemResponseItems) HasSalePrice() bool`

HasSalePrice returns a boolean if a field has been set.

### SetSalePriceNil

`func (o *GetAnalyticItemResponseItems) SetSalePriceNil(b bool)`

 SetSalePriceNil sets the value for SalePrice to be an explicit nil

### UnsetSalePrice
`func (o *GetAnalyticItemResponseItems) UnsetSalePrice()`

UnsetSalePrice ensures that no value is present for SalePrice, not even an explicit nil
### GetPlatformFee

`func (o *GetAnalyticItemResponseItems) GetPlatformFee() float32`

GetPlatformFee returns the PlatformFee field if non-nil, zero value otherwise.

### GetPlatformFeeOk

`func (o *GetAnalyticItemResponseItems) GetPlatformFeeOk() (*float32, bool)`

GetPlatformFeeOk returns a tuple with the PlatformFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformFee

`func (o *GetAnalyticItemResponseItems) SetPlatformFee(v float32)`

SetPlatformFee sets PlatformFee field to given value.

### HasPlatformFee

`func (o *GetAnalyticItemResponseItems) HasPlatformFee() bool`

HasPlatformFee returns a boolean if a field has been set.

### SetPlatformFeeNil

`func (o *GetAnalyticItemResponseItems) SetPlatformFeeNil(b bool)`

 SetPlatformFeeNil sets the value for PlatformFee to be an explicit nil

### UnsetPlatformFee
`func (o *GetAnalyticItemResponseItems) UnsetPlatformFee()`

UnsetPlatformFee ensures that no value is present for PlatformFee, not even an explicit nil
### GetShippingCost

`func (o *GetAnalyticItemResponseItems) GetShippingCost() float32`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *GetAnalyticItemResponseItems) GetShippingCostOk() (*float32, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *GetAnalyticItemResponseItems) SetShippingCost(v float32)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *GetAnalyticItemResponseItems) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### SetShippingCostNil

`func (o *GetAnalyticItemResponseItems) SetShippingCostNil(b bool)`

 SetShippingCostNil sets the value for ShippingCost to be an explicit nil

### UnsetShippingCost
`func (o *GetAnalyticItemResponseItems) UnsetShippingCost()`

UnsetShippingCost ensures that no value is present for ShippingCost, not even an explicit nil
### GetCostOfGoods

`func (o *GetAnalyticItemResponseItems) GetCostOfGoods() float32`

GetCostOfGoods returns the CostOfGoods field if non-nil, zero value otherwise.

### GetCostOfGoodsOk

`func (o *GetAnalyticItemResponseItems) GetCostOfGoodsOk() (*float32, bool)`

GetCostOfGoodsOk returns a tuple with the CostOfGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOfGoods

`func (o *GetAnalyticItemResponseItems) SetCostOfGoods(v float32)`

SetCostOfGoods sets CostOfGoods field to given value.

### HasCostOfGoods

`func (o *GetAnalyticItemResponseItems) HasCostOfGoods() bool`

HasCostOfGoods returns a boolean if a field has been set.

### SetCostOfGoodsNil

`func (o *GetAnalyticItemResponseItems) SetCostOfGoodsNil(b bool)`

 SetCostOfGoodsNil sets the value for CostOfGoods to be an explicit nil

### UnsetCostOfGoods
`func (o *GetAnalyticItemResponseItems) UnsetCostOfGoods()`

UnsetCostOfGoods ensures that no value is present for CostOfGoods, not even an explicit nil
### GetNetProfit

`func (o *GetAnalyticItemResponseItems) GetNetProfit() float32`

GetNetProfit returns the NetProfit field if non-nil, zero value otherwise.

### GetNetProfitOk

`func (o *GetAnalyticItemResponseItems) GetNetProfitOk() (*float32, bool)`

GetNetProfitOk returns a tuple with the NetProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetProfit

`func (o *GetAnalyticItemResponseItems) SetNetProfit(v float32)`

SetNetProfit sets NetProfit field to given value.

### HasNetProfit

`func (o *GetAnalyticItemResponseItems) HasNetProfit() bool`

HasNetProfit returns a boolean if a field has been set.

### SetNetProfitNil

`func (o *GetAnalyticItemResponseItems) SetNetProfitNil(b bool)`

 SetNetProfitNil sets the value for NetProfit to be an explicit nil

### UnsetNetProfit
`func (o *GetAnalyticItemResponseItems) UnsetNetProfit()`

UnsetNetProfit ensures that no value is present for NetProfit, not even an explicit nil
### GetMarginPct

`func (o *GetAnalyticItemResponseItems) GetMarginPct() float32`

GetMarginPct returns the MarginPct field if non-nil, zero value otherwise.

### GetMarginPctOk

`func (o *GetAnalyticItemResponseItems) GetMarginPctOk() (*float32, bool)`

GetMarginPctOk returns a tuple with the MarginPct field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarginPct

`func (o *GetAnalyticItemResponseItems) SetMarginPct(v float32)`

SetMarginPct sets MarginPct field to given value.


### GetDaysToSell

`func (o *GetAnalyticItemResponseItems) GetDaysToSell() float32`

GetDaysToSell returns the DaysToSell field if non-nil, zero value otherwise.

### GetDaysToSellOk

`func (o *GetAnalyticItemResponseItems) GetDaysToSellOk() (*float32, bool)`

GetDaysToSellOk returns a tuple with the DaysToSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaysToSell

`func (o *GetAnalyticItemResponseItems) SetDaysToSell(v float32)`

SetDaysToSell sets DaysToSell field to given value.

### HasDaysToSell

`func (o *GetAnalyticItemResponseItems) HasDaysToSell() bool`

HasDaysToSell returns a boolean if a field has been set.

### SetDaysToSellNil

`func (o *GetAnalyticItemResponseItems) SetDaysToSellNil(b bool)`

 SetDaysToSellNil sets the value for DaysToSell to be an explicit nil

### UnsetDaysToSell
`func (o *GetAnalyticItemResponseItems) UnsetDaysToSell()`

UnsetDaysToSell ensures that no value is present for DaysToSell, not even an explicit nil
### GetDetectedAt

`func (o *GetAnalyticItemResponseItems) GetDetectedAt() time.Time`

GetDetectedAt returns the DetectedAt field if non-nil, zero value otherwise.

### GetDetectedAtOk

`func (o *GetAnalyticItemResponseItems) GetDetectedAtOk() (*time.Time, bool)`

GetDetectedAtOk returns a tuple with the DetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedAt

`func (o *GetAnalyticItemResponseItems) SetDetectedAt(v time.Time)`

SetDetectedAt sets DetectedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


