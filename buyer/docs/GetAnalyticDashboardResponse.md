# GetAnalyticDashboardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Revenue** | [**GetAnalyticDashboardResponseRevenue**](GetAnalyticDashboardResponseRevenue.md) |  | 
**Profit** | [**GetAnalyticDashboardResponseRevenue**](GetAnalyticDashboardResponseRevenue.md) |  | 
**SoldItemsCount** | [**GetAnalyticDashboardResponseRevenue**](GetAnalyticDashboardResponseRevenue.md) |  | 
**CancelledCount** | [**GetAnalyticDashboardResponseRevenue**](GetAnalyticDashboardResponseRevenue.md) |  | 
**CancelRate** | [**GetAnalyticDashboardResponseCancelRate**](GetAnalyticDashboardResponseCancelRate.md) |  | 
**RefundedAmount** | [**GetAnalyticDashboardResponseRevenue**](GetAnalyticDashboardResponseRevenue.md) |  | 
**RefundedCount** | [**GetAnalyticDashboardResponseRevenue**](GetAnalyticDashboardResponseRevenue.md) |  | 
**ListedItemsCount** | [**GetAnalyticDashboardResponseRevenue**](GetAnalyticDashboardResponseRevenue.md) |  | 
**AvgSalePrice** | [**GetAnalyticDashboardResponseRevenue**](GetAnalyticDashboardResponseRevenue.md) |  | 
**SellThroughRate** | Pointer to **NullableFloat32** |  | [optional] 
**SellThrough** | [**GetAnalyticDashboardResponseSellThrough**](GetAnalyticDashboardResponseSellThrough.md) |  | 
**AvgDaysToSell** | Pointer to **NullableFloat32** |  | [optional] 
**SalesByPlatform** | [**[]GetAnalyticDashboardResponseSalesByPlatform**](GetAnalyticDashboardResponseSalesByPlatform.md) |  | 
**AvgSalePriceByPlatform** | [**[]GetAnalyticDashboardResponseAvgSalePriceByPlatform**](GetAnalyticDashboardResponseAvgSalePriceByPlatform.md) |  | 
**RevenueByDay** | [**[]GetAnalyticDashboardResponseRevenueByDay**](GetAnalyticDashboardResponseRevenueByDay.md) |  | 
**TopCategories** | [**[]GetAnalyticDashboardResponseTopCategories**](GetAnalyticDashboardResponseTopCategories.md) |  | 
**TopSubcategories** | [**[]GetAnalyticDashboardResponseTopCategories**](GetAnalyticDashboardResponseTopCategories.md) |  | 
**TopBrands** | [**[]GetAnalyticDashboardResponseTopCategories**](GetAnalyticDashboardResponseTopCategories.md) |  | 
**Top5ItemsByRevenue** | [**[]GetAnalyticDashboardResponseTop5ItemsByRevenue**](GetAnalyticDashboardResponseTop5ItemsByRevenue.md) |  | 
**RecentSales** | [**[]GetAnalyticDashboardResponseRecentSales**](GetAnalyticDashboardResponseRecentSales.md) |  | 
**AvailableLabels** | **[]string** |  | 

## Methods

### NewGetAnalyticDashboardResponse

`func NewGetAnalyticDashboardResponse(revenue GetAnalyticDashboardResponseRevenue, profit GetAnalyticDashboardResponseRevenue, soldItemsCount GetAnalyticDashboardResponseRevenue, cancelledCount GetAnalyticDashboardResponseRevenue, cancelRate GetAnalyticDashboardResponseCancelRate, refundedAmount GetAnalyticDashboardResponseRevenue, refundedCount GetAnalyticDashboardResponseRevenue, listedItemsCount GetAnalyticDashboardResponseRevenue, avgSalePrice GetAnalyticDashboardResponseRevenue, sellThrough GetAnalyticDashboardResponseSellThrough, salesByPlatform []GetAnalyticDashboardResponseSalesByPlatform, avgSalePriceByPlatform []GetAnalyticDashboardResponseAvgSalePriceByPlatform, revenueByDay []GetAnalyticDashboardResponseRevenueByDay, topCategories []GetAnalyticDashboardResponseTopCategories, topSubcategories []GetAnalyticDashboardResponseTopCategories, topBrands []GetAnalyticDashboardResponseTopCategories, top5ItemsByRevenue []GetAnalyticDashboardResponseTop5ItemsByRevenue, recentSales []GetAnalyticDashboardResponseRecentSales, availableLabels []string, ) *GetAnalyticDashboardResponse`

NewGetAnalyticDashboardResponse instantiates a new GetAnalyticDashboardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalyticDashboardResponseWithDefaults

`func NewGetAnalyticDashboardResponseWithDefaults() *GetAnalyticDashboardResponse`

NewGetAnalyticDashboardResponseWithDefaults instantiates a new GetAnalyticDashboardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRevenue

`func (o *GetAnalyticDashboardResponse) GetRevenue() GetAnalyticDashboardResponseRevenue`

GetRevenue returns the Revenue field if non-nil, zero value otherwise.

### GetRevenueOk

`func (o *GetAnalyticDashboardResponse) GetRevenueOk() (*GetAnalyticDashboardResponseRevenue, bool)`

GetRevenueOk returns a tuple with the Revenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenue

`func (o *GetAnalyticDashboardResponse) SetRevenue(v GetAnalyticDashboardResponseRevenue)`

SetRevenue sets Revenue field to given value.


### GetProfit

`func (o *GetAnalyticDashboardResponse) GetProfit() GetAnalyticDashboardResponseRevenue`

GetProfit returns the Profit field if non-nil, zero value otherwise.

### GetProfitOk

`func (o *GetAnalyticDashboardResponse) GetProfitOk() (*GetAnalyticDashboardResponseRevenue, bool)`

GetProfitOk returns a tuple with the Profit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfit

`func (o *GetAnalyticDashboardResponse) SetProfit(v GetAnalyticDashboardResponseRevenue)`

SetProfit sets Profit field to given value.


### GetSoldItemsCount

`func (o *GetAnalyticDashboardResponse) GetSoldItemsCount() GetAnalyticDashboardResponseRevenue`

GetSoldItemsCount returns the SoldItemsCount field if non-nil, zero value otherwise.

### GetSoldItemsCountOk

`func (o *GetAnalyticDashboardResponse) GetSoldItemsCountOk() (*GetAnalyticDashboardResponseRevenue, bool)`

GetSoldItemsCountOk returns a tuple with the SoldItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldItemsCount

`func (o *GetAnalyticDashboardResponse) SetSoldItemsCount(v GetAnalyticDashboardResponseRevenue)`

SetSoldItemsCount sets SoldItemsCount field to given value.


### GetCancelledCount

`func (o *GetAnalyticDashboardResponse) GetCancelledCount() GetAnalyticDashboardResponseRevenue`

GetCancelledCount returns the CancelledCount field if non-nil, zero value otherwise.

### GetCancelledCountOk

`func (o *GetAnalyticDashboardResponse) GetCancelledCountOk() (*GetAnalyticDashboardResponseRevenue, bool)`

GetCancelledCountOk returns a tuple with the CancelledCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledCount

`func (o *GetAnalyticDashboardResponse) SetCancelledCount(v GetAnalyticDashboardResponseRevenue)`

SetCancelledCount sets CancelledCount field to given value.


### GetCancelRate

`func (o *GetAnalyticDashboardResponse) GetCancelRate() GetAnalyticDashboardResponseCancelRate`

GetCancelRate returns the CancelRate field if non-nil, zero value otherwise.

### GetCancelRateOk

`func (o *GetAnalyticDashboardResponse) GetCancelRateOk() (*GetAnalyticDashboardResponseCancelRate, bool)`

GetCancelRateOk returns a tuple with the CancelRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelRate

`func (o *GetAnalyticDashboardResponse) SetCancelRate(v GetAnalyticDashboardResponseCancelRate)`

SetCancelRate sets CancelRate field to given value.


### GetRefundedAmount

`func (o *GetAnalyticDashboardResponse) GetRefundedAmount() GetAnalyticDashboardResponseRevenue`

GetRefundedAmount returns the RefundedAmount field if non-nil, zero value otherwise.

### GetRefundedAmountOk

`func (o *GetAnalyticDashboardResponse) GetRefundedAmountOk() (*GetAnalyticDashboardResponseRevenue, bool)`

GetRefundedAmountOk returns a tuple with the RefundedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmount

`func (o *GetAnalyticDashboardResponse) SetRefundedAmount(v GetAnalyticDashboardResponseRevenue)`

SetRefundedAmount sets RefundedAmount field to given value.


### GetRefundedCount

`func (o *GetAnalyticDashboardResponse) GetRefundedCount() GetAnalyticDashboardResponseRevenue`

GetRefundedCount returns the RefundedCount field if non-nil, zero value otherwise.

### GetRefundedCountOk

`func (o *GetAnalyticDashboardResponse) GetRefundedCountOk() (*GetAnalyticDashboardResponseRevenue, bool)`

GetRefundedCountOk returns a tuple with the RefundedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedCount

`func (o *GetAnalyticDashboardResponse) SetRefundedCount(v GetAnalyticDashboardResponseRevenue)`

SetRefundedCount sets RefundedCount field to given value.


### GetListedItemsCount

`func (o *GetAnalyticDashboardResponse) GetListedItemsCount() GetAnalyticDashboardResponseRevenue`

GetListedItemsCount returns the ListedItemsCount field if non-nil, zero value otherwise.

### GetListedItemsCountOk

`func (o *GetAnalyticDashboardResponse) GetListedItemsCountOk() (*GetAnalyticDashboardResponseRevenue, bool)`

GetListedItemsCountOk returns a tuple with the ListedItemsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedItemsCount

`func (o *GetAnalyticDashboardResponse) SetListedItemsCount(v GetAnalyticDashboardResponseRevenue)`

SetListedItemsCount sets ListedItemsCount field to given value.


### GetAvgSalePrice

`func (o *GetAnalyticDashboardResponse) GetAvgSalePrice() GetAnalyticDashboardResponseRevenue`

GetAvgSalePrice returns the AvgSalePrice field if non-nil, zero value otherwise.

### GetAvgSalePriceOk

`func (o *GetAnalyticDashboardResponse) GetAvgSalePriceOk() (*GetAnalyticDashboardResponseRevenue, bool)`

GetAvgSalePriceOk returns a tuple with the AvgSalePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSalePrice

`func (o *GetAnalyticDashboardResponse) SetAvgSalePrice(v GetAnalyticDashboardResponseRevenue)`

SetAvgSalePrice sets AvgSalePrice field to given value.


### GetSellThroughRate

`func (o *GetAnalyticDashboardResponse) GetSellThroughRate() float32`

GetSellThroughRate returns the SellThroughRate field if non-nil, zero value otherwise.

### GetSellThroughRateOk

`func (o *GetAnalyticDashboardResponse) GetSellThroughRateOk() (*float32, bool)`

GetSellThroughRateOk returns a tuple with the SellThroughRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellThroughRate

`func (o *GetAnalyticDashboardResponse) SetSellThroughRate(v float32)`

SetSellThroughRate sets SellThroughRate field to given value.

### HasSellThroughRate

`func (o *GetAnalyticDashboardResponse) HasSellThroughRate() bool`

HasSellThroughRate returns a boolean if a field has been set.

### SetSellThroughRateNil

`func (o *GetAnalyticDashboardResponse) SetSellThroughRateNil(b bool)`

 SetSellThroughRateNil sets the value for SellThroughRate to be an explicit nil

### UnsetSellThroughRate
`func (o *GetAnalyticDashboardResponse) UnsetSellThroughRate()`

UnsetSellThroughRate ensures that no value is present for SellThroughRate, not even an explicit nil
### GetSellThrough

`func (o *GetAnalyticDashboardResponse) GetSellThrough() GetAnalyticDashboardResponseSellThrough`

GetSellThrough returns the SellThrough field if non-nil, zero value otherwise.

### GetSellThroughOk

`func (o *GetAnalyticDashboardResponse) GetSellThroughOk() (*GetAnalyticDashboardResponseSellThrough, bool)`

GetSellThroughOk returns a tuple with the SellThrough field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellThrough

`func (o *GetAnalyticDashboardResponse) SetSellThrough(v GetAnalyticDashboardResponseSellThrough)`

SetSellThrough sets SellThrough field to given value.


### GetAvgDaysToSell

`func (o *GetAnalyticDashboardResponse) GetAvgDaysToSell() float32`

GetAvgDaysToSell returns the AvgDaysToSell field if non-nil, zero value otherwise.

### GetAvgDaysToSellOk

`func (o *GetAnalyticDashboardResponse) GetAvgDaysToSellOk() (*float32, bool)`

GetAvgDaysToSellOk returns a tuple with the AvgDaysToSell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgDaysToSell

`func (o *GetAnalyticDashboardResponse) SetAvgDaysToSell(v float32)`

SetAvgDaysToSell sets AvgDaysToSell field to given value.

### HasAvgDaysToSell

`func (o *GetAnalyticDashboardResponse) HasAvgDaysToSell() bool`

HasAvgDaysToSell returns a boolean if a field has been set.

### SetAvgDaysToSellNil

`func (o *GetAnalyticDashboardResponse) SetAvgDaysToSellNil(b bool)`

 SetAvgDaysToSellNil sets the value for AvgDaysToSell to be an explicit nil

### UnsetAvgDaysToSell
`func (o *GetAnalyticDashboardResponse) UnsetAvgDaysToSell()`

UnsetAvgDaysToSell ensures that no value is present for AvgDaysToSell, not even an explicit nil
### GetSalesByPlatform

`func (o *GetAnalyticDashboardResponse) GetSalesByPlatform() []GetAnalyticDashboardResponseSalesByPlatform`

GetSalesByPlatform returns the SalesByPlatform field if non-nil, zero value otherwise.

### GetSalesByPlatformOk

`func (o *GetAnalyticDashboardResponse) GetSalesByPlatformOk() (*[]GetAnalyticDashboardResponseSalesByPlatform, bool)`

GetSalesByPlatformOk returns a tuple with the SalesByPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesByPlatform

`func (o *GetAnalyticDashboardResponse) SetSalesByPlatform(v []GetAnalyticDashboardResponseSalesByPlatform)`

SetSalesByPlatform sets SalesByPlatform field to given value.


### GetAvgSalePriceByPlatform

`func (o *GetAnalyticDashboardResponse) GetAvgSalePriceByPlatform() []GetAnalyticDashboardResponseAvgSalePriceByPlatform`

GetAvgSalePriceByPlatform returns the AvgSalePriceByPlatform field if non-nil, zero value otherwise.

### GetAvgSalePriceByPlatformOk

`func (o *GetAnalyticDashboardResponse) GetAvgSalePriceByPlatformOk() (*[]GetAnalyticDashboardResponseAvgSalePriceByPlatform, bool)`

GetAvgSalePriceByPlatformOk returns a tuple with the AvgSalePriceByPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvgSalePriceByPlatform

`func (o *GetAnalyticDashboardResponse) SetAvgSalePriceByPlatform(v []GetAnalyticDashboardResponseAvgSalePriceByPlatform)`

SetAvgSalePriceByPlatform sets AvgSalePriceByPlatform field to given value.


### GetRevenueByDay

`func (o *GetAnalyticDashboardResponse) GetRevenueByDay() []GetAnalyticDashboardResponseRevenueByDay`

GetRevenueByDay returns the RevenueByDay field if non-nil, zero value otherwise.

### GetRevenueByDayOk

`func (o *GetAnalyticDashboardResponse) GetRevenueByDayOk() (*[]GetAnalyticDashboardResponseRevenueByDay, bool)`

GetRevenueByDayOk returns a tuple with the RevenueByDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevenueByDay

`func (o *GetAnalyticDashboardResponse) SetRevenueByDay(v []GetAnalyticDashboardResponseRevenueByDay)`

SetRevenueByDay sets RevenueByDay field to given value.


### GetTopCategories

`func (o *GetAnalyticDashboardResponse) GetTopCategories() []GetAnalyticDashboardResponseTopCategories`

GetTopCategories returns the TopCategories field if non-nil, zero value otherwise.

### GetTopCategoriesOk

`func (o *GetAnalyticDashboardResponse) GetTopCategoriesOk() (*[]GetAnalyticDashboardResponseTopCategories, bool)`

GetTopCategoriesOk returns a tuple with the TopCategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopCategories

`func (o *GetAnalyticDashboardResponse) SetTopCategories(v []GetAnalyticDashboardResponseTopCategories)`

SetTopCategories sets TopCategories field to given value.


### GetTopSubcategories

`func (o *GetAnalyticDashboardResponse) GetTopSubcategories() []GetAnalyticDashboardResponseTopCategories`

GetTopSubcategories returns the TopSubcategories field if non-nil, zero value otherwise.

### GetTopSubcategoriesOk

`func (o *GetAnalyticDashboardResponse) GetTopSubcategoriesOk() (*[]GetAnalyticDashboardResponseTopCategories, bool)`

GetTopSubcategoriesOk returns a tuple with the TopSubcategories field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopSubcategories

`func (o *GetAnalyticDashboardResponse) SetTopSubcategories(v []GetAnalyticDashboardResponseTopCategories)`

SetTopSubcategories sets TopSubcategories field to given value.


### GetTopBrands

`func (o *GetAnalyticDashboardResponse) GetTopBrands() []GetAnalyticDashboardResponseTopCategories`

GetTopBrands returns the TopBrands field if non-nil, zero value otherwise.

### GetTopBrandsOk

`func (o *GetAnalyticDashboardResponse) GetTopBrandsOk() (*[]GetAnalyticDashboardResponseTopCategories, bool)`

GetTopBrandsOk returns a tuple with the TopBrands field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopBrands

`func (o *GetAnalyticDashboardResponse) SetTopBrands(v []GetAnalyticDashboardResponseTopCategories)`

SetTopBrands sets TopBrands field to given value.


### GetTop5ItemsByRevenue

`func (o *GetAnalyticDashboardResponse) GetTop5ItemsByRevenue() []GetAnalyticDashboardResponseTop5ItemsByRevenue`

GetTop5ItemsByRevenue returns the Top5ItemsByRevenue field if non-nil, zero value otherwise.

### GetTop5ItemsByRevenueOk

`func (o *GetAnalyticDashboardResponse) GetTop5ItemsByRevenueOk() (*[]GetAnalyticDashboardResponseTop5ItemsByRevenue, bool)`

GetTop5ItemsByRevenueOk returns a tuple with the Top5ItemsByRevenue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTop5ItemsByRevenue

`func (o *GetAnalyticDashboardResponse) SetTop5ItemsByRevenue(v []GetAnalyticDashboardResponseTop5ItemsByRevenue)`

SetTop5ItemsByRevenue sets Top5ItemsByRevenue field to given value.


### GetRecentSales

`func (o *GetAnalyticDashboardResponse) GetRecentSales() []GetAnalyticDashboardResponseRecentSales`

GetRecentSales returns the RecentSales field if non-nil, zero value otherwise.

### GetRecentSalesOk

`func (o *GetAnalyticDashboardResponse) GetRecentSalesOk() (*[]GetAnalyticDashboardResponseRecentSales, bool)`

GetRecentSalesOk returns a tuple with the RecentSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentSales

`func (o *GetAnalyticDashboardResponse) SetRecentSales(v []GetAnalyticDashboardResponseRecentSales)`

SetRecentSales sets RecentSales field to given value.


### GetAvailableLabels

`func (o *GetAnalyticDashboardResponse) GetAvailableLabels() []string`

GetAvailableLabels returns the AvailableLabels field if non-nil, zero value otherwise.

### GetAvailableLabelsOk

`func (o *GetAnalyticDashboardResponse) GetAvailableLabelsOk() (*[]string, bool)`

GetAvailableLabelsOk returns a tuple with the AvailableLabels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableLabels

`func (o *GetAnalyticDashboardResponse) SetAvailableLabels(v []string)`

SetAvailableLabels sets AvailableLabels field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


