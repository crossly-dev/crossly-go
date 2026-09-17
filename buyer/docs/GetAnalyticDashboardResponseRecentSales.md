# GetAnalyticDashboardResponseRecentSales

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SaleId** | **string** |  | 
**Platform** | **string** |  | 
**PlatformOrderId** | Pointer to **NullableString** |  | [optional] 
**SalePrice** | Pointer to **NullableFloat32** |  | [optional] 
**PlatformFee** | Pointer to **NullableFloat32** |  | [optional] 
**ShippingCost** | Pointer to **NullableFloat32** |  | [optional] 
**CostOfGoods** | Pointer to **NullableFloat32** |  | [optional] 
**NetProfit** | Pointer to **NullableFloat32** |  | [optional] 
**DetectedBy** | **string** |  | 
**DetectedAt** | **time.Time** |  | 
**ItemTitle** | Pointer to **NullableString** |  | [optional] 
**ItemImage** | **string** |  | 

## Methods

### NewGetAnalyticDashboardResponseRecentSales

`func NewGetAnalyticDashboardResponseRecentSales(saleId string, platform string, detectedBy string, detectedAt time.Time, itemImage string, ) *GetAnalyticDashboardResponseRecentSales`

NewGetAnalyticDashboardResponseRecentSales instantiates a new GetAnalyticDashboardResponseRecentSales object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalyticDashboardResponseRecentSalesWithDefaults

`func NewGetAnalyticDashboardResponseRecentSalesWithDefaults() *GetAnalyticDashboardResponseRecentSales`

NewGetAnalyticDashboardResponseRecentSalesWithDefaults instantiates a new GetAnalyticDashboardResponseRecentSales object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaleId

`func (o *GetAnalyticDashboardResponseRecentSales) GetSaleId() string`

GetSaleId returns the SaleId field if non-nil, zero value otherwise.

### GetSaleIdOk

`func (o *GetAnalyticDashboardResponseRecentSales) GetSaleIdOk() (*string, bool)`

GetSaleIdOk returns a tuple with the SaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleId

`func (o *GetAnalyticDashboardResponseRecentSales) SetSaleId(v string)`

SetSaleId sets SaleId field to given value.


### GetPlatform

`func (o *GetAnalyticDashboardResponseRecentSales) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetAnalyticDashboardResponseRecentSales) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetAnalyticDashboardResponseRecentSales) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPlatformOrderId

`func (o *GetAnalyticDashboardResponseRecentSales) GetPlatformOrderId() string`

GetPlatformOrderId returns the PlatformOrderId field if non-nil, zero value otherwise.

### GetPlatformOrderIdOk

`func (o *GetAnalyticDashboardResponseRecentSales) GetPlatformOrderIdOk() (*string, bool)`

GetPlatformOrderIdOk returns a tuple with the PlatformOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformOrderId

`func (o *GetAnalyticDashboardResponseRecentSales) SetPlatformOrderId(v string)`

SetPlatformOrderId sets PlatformOrderId field to given value.

### HasPlatformOrderId

`func (o *GetAnalyticDashboardResponseRecentSales) HasPlatformOrderId() bool`

HasPlatformOrderId returns a boolean if a field has been set.

### SetPlatformOrderIdNil

`func (o *GetAnalyticDashboardResponseRecentSales) SetPlatformOrderIdNil(b bool)`

 SetPlatformOrderIdNil sets the value for PlatformOrderId to be an explicit nil

### UnsetPlatformOrderId
`func (o *GetAnalyticDashboardResponseRecentSales) UnsetPlatformOrderId()`

UnsetPlatformOrderId ensures that no value is present for PlatformOrderId, not even an explicit nil
### GetSalePrice

`func (o *GetAnalyticDashboardResponseRecentSales) GetSalePrice() float32`

GetSalePrice returns the SalePrice field if non-nil, zero value otherwise.

### GetSalePriceOk

`func (o *GetAnalyticDashboardResponseRecentSales) GetSalePriceOk() (*float32, bool)`

GetSalePriceOk returns a tuple with the SalePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalePrice

`func (o *GetAnalyticDashboardResponseRecentSales) SetSalePrice(v float32)`

SetSalePrice sets SalePrice field to given value.

### HasSalePrice

`func (o *GetAnalyticDashboardResponseRecentSales) HasSalePrice() bool`

HasSalePrice returns a boolean if a field has been set.

### SetSalePriceNil

`func (o *GetAnalyticDashboardResponseRecentSales) SetSalePriceNil(b bool)`

 SetSalePriceNil sets the value for SalePrice to be an explicit nil

### UnsetSalePrice
`func (o *GetAnalyticDashboardResponseRecentSales) UnsetSalePrice()`

UnsetSalePrice ensures that no value is present for SalePrice, not even an explicit nil
### GetPlatformFee

`func (o *GetAnalyticDashboardResponseRecentSales) GetPlatformFee() float32`

GetPlatformFee returns the PlatformFee field if non-nil, zero value otherwise.

### GetPlatformFeeOk

`func (o *GetAnalyticDashboardResponseRecentSales) GetPlatformFeeOk() (*float32, bool)`

GetPlatformFeeOk returns a tuple with the PlatformFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformFee

`func (o *GetAnalyticDashboardResponseRecentSales) SetPlatformFee(v float32)`

SetPlatformFee sets PlatformFee field to given value.

### HasPlatformFee

`func (o *GetAnalyticDashboardResponseRecentSales) HasPlatformFee() bool`

HasPlatformFee returns a boolean if a field has been set.

### SetPlatformFeeNil

`func (o *GetAnalyticDashboardResponseRecentSales) SetPlatformFeeNil(b bool)`

 SetPlatformFeeNil sets the value for PlatformFee to be an explicit nil

### UnsetPlatformFee
`func (o *GetAnalyticDashboardResponseRecentSales) UnsetPlatformFee()`

UnsetPlatformFee ensures that no value is present for PlatformFee, not even an explicit nil
### GetShippingCost

`func (o *GetAnalyticDashboardResponseRecentSales) GetShippingCost() float32`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *GetAnalyticDashboardResponseRecentSales) GetShippingCostOk() (*float32, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *GetAnalyticDashboardResponseRecentSales) SetShippingCost(v float32)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *GetAnalyticDashboardResponseRecentSales) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### SetShippingCostNil

`func (o *GetAnalyticDashboardResponseRecentSales) SetShippingCostNil(b bool)`

 SetShippingCostNil sets the value for ShippingCost to be an explicit nil

### UnsetShippingCost
`func (o *GetAnalyticDashboardResponseRecentSales) UnsetShippingCost()`

UnsetShippingCost ensures that no value is present for ShippingCost, not even an explicit nil
### GetCostOfGoods

`func (o *GetAnalyticDashboardResponseRecentSales) GetCostOfGoods() float32`

GetCostOfGoods returns the CostOfGoods field if non-nil, zero value otherwise.

### GetCostOfGoodsOk

`func (o *GetAnalyticDashboardResponseRecentSales) GetCostOfGoodsOk() (*float32, bool)`

GetCostOfGoodsOk returns a tuple with the CostOfGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOfGoods

`func (o *GetAnalyticDashboardResponseRecentSales) SetCostOfGoods(v float32)`

SetCostOfGoods sets CostOfGoods field to given value.

### HasCostOfGoods

`func (o *GetAnalyticDashboardResponseRecentSales) HasCostOfGoods() bool`

HasCostOfGoods returns a boolean if a field has been set.

### SetCostOfGoodsNil

`func (o *GetAnalyticDashboardResponseRecentSales) SetCostOfGoodsNil(b bool)`

 SetCostOfGoodsNil sets the value for CostOfGoods to be an explicit nil

### UnsetCostOfGoods
`func (o *GetAnalyticDashboardResponseRecentSales) UnsetCostOfGoods()`

UnsetCostOfGoods ensures that no value is present for CostOfGoods, not even an explicit nil
### GetNetProfit

`func (o *GetAnalyticDashboardResponseRecentSales) GetNetProfit() float32`

GetNetProfit returns the NetProfit field if non-nil, zero value otherwise.

### GetNetProfitOk

`func (o *GetAnalyticDashboardResponseRecentSales) GetNetProfitOk() (*float32, bool)`

GetNetProfitOk returns a tuple with the NetProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetProfit

`func (o *GetAnalyticDashboardResponseRecentSales) SetNetProfit(v float32)`

SetNetProfit sets NetProfit field to given value.

### HasNetProfit

`func (o *GetAnalyticDashboardResponseRecentSales) HasNetProfit() bool`

HasNetProfit returns a boolean if a field has been set.

### SetNetProfitNil

`func (o *GetAnalyticDashboardResponseRecentSales) SetNetProfitNil(b bool)`

 SetNetProfitNil sets the value for NetProfit to be an explicit nil

### UnsetNetProfit
`func (o *GetAnalyticDashboardResponseRecentSales) UnsetNetProfit()`

UnsetNetProfit ensures that no value is present for NetProfit, not even an explicit nil
### GetDetectedBy

`func (o *GetAnalyticDashboardResponseRecentSales) GetDetectedBy() string`

GetDetectedBy returns the DetectedBy field if non-nil, zero value otherwise.

### GetDetectedByOk

`func (o *GetAnalyticDashboardResponseRecentSales) GetDetectedByOk() (*string, bool)`

GetDetectedByOk returns a tuple with the DetectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedBy

`func (o *GetAnalyticDashboardResponseRecentSales) SetDetectedBy(v string)`

SetDetectedBy sets DetectedBy field to given value.


### GetDetectedAt

`func (o *GetAnalyticDashboardResponseRecentSales) GetDetectedAt() time.Time`

GetDetectedAt returns the DetectedAt field if non-nil, zero value otherwise.

### GetDetectedAtOk

`func (o *GetAnalyticDashboardResponseRecentSales) GetDetectedAtOk() (*time.Time, bool)`

GetDetectedAtOk returns a tuple with the DetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedAt

`func (o *GetAnalyticDashboardResponseRecentSales) SetDetectedAt(v time.Time)`

SetDetectedAt sets DetectedAt field to given value.


### GetItemTitle

`func (o *GetAnalyticDashboardResponseRecentSales) GetItemTitle() string`

GetItemTitle returns the ItemTitle field if non-nil, zero value otherwise.

### GetItemTitleOk

`func (o *GetAnalyticDashboardResponseRecentSales) GetItemTitleOk() (*string, bool)`

GetItemTitleOk returns a tuple with the ItemTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTitle

`func (o *GetAnalyticDashboardResponseRecentSales) SetItemTitle(v string)`

SetItemTitle sets ItemTitle field to given value.

### HasItemTitle

`func (o *GetAnalyticDashboardResponseRecentSales) HasItemTitle() bool`

HasItemTitle returns a boolean if a field has been set.

### SetItemTitleNil

`func (o *GetAnalyticDashboardResponseRecentSales) SetItemTitleNil(b bool)`

 SetItemTitleNil sets the value for ItemTitle to be an explicit nil

### UnsetItemTitle
`func (o *GetAnalyticDashboardResponseRecentSales) UnsetItemTitle()`

UnsetItemTitle ensures that no value is present for ItemTitle, not even an explicit nil
### GetItemImage

`func (o *GetAnalyticDashboardResponseRecentSales) GetItemImage() string`

GetItemImage returns the ItemImage field if non-nil, zero value otherwise.

### GetItemImageOk

`func (o *GetAnalyticDashboardResponseRecentSales) GetItemImageOk() (*string, bool)`

GetItemImageOk returns a tuple with the ItemImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemImage

`func (o *GetAnalyticDashboardResponseRecentSales) SetItemImage(v string)`

SetItemImage sets ItemImage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


