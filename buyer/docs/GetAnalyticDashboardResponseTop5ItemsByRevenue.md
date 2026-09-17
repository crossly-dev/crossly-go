# GetAnalyticDashboardResponseTop5ItemsByRevenue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SaleId** | **string** |  | 
**Platform** | **string** |  | 
**SalePrice** | Pointer to **NullableFloat32** |  | [optional] 
**NetProfit** | Pointer to **NullableFloat32** |  | [optional] 
**DetectedAt** | **time.Time** |  | 
**ItemTitle** | Pointer to **NullableString** |  | [optional] 
**ItemImage** | **string** |  | 

## Methods

### NewGetAnalyticDashboardResponseTop5ItemsByRevenue

`func NewGetAnalyticDashboardResponseTop5ItemsByRevenue(saleId string, platform string, detectedAt time.Time, itemImage string, ) *GetAnalyticDashboardResponseTop5ItemsByRevenue`

NewGetAnalyticDashboardResponseTop5ItemsByRevenue instantiates a new GetAnalyticDashboardResponseTop5ItemsByRevenue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAnalyticDashboardResponseTop5ItemsByRevenueWithDefaults

`func NewGetAnalyticDashboardResponseTop5ItemsByRevenueWithDefaults() *GetAnalyticDashboardResponseTop5ItemsByRevenue`

NewGetAnalyticDashboardResponseTop5ItemsByRevenueWithDefaults instantiates a new GetAnalyticDashboardResponseTop5ItemsByRevenue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSaleId

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) GetSaleId() string`

GetSaleId returns the SaleId field if non-nil, zero value otherwise.

### GetSaleIdOk

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) GetSaleIdOk() (*string, bool)`

GetSaleIdOk returns a tuple with the SaleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaleId

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) SetSaleId(v string)`

SetSaleId sets SaleId field to given value.


### GetPlatform

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetSalePrice

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) GetSalePrice() float32`

GetSalePrice returns the SalePrice field if non-nil, zero value otherwise.

### GetSalePriceOk

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) GetSalePriceOk() (*float32, bool)`

GetSalePriceOk returns a tuple with the SalePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalePrice

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) SetSalePrice(v float32)`

SetSalePrice sets SalePrice field to given value.

### HasSalePrice

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) HasSalePrice() bool`

HasSalePrice returns a boolean if a field has been set.

### SetSalePriceNil

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) SetSalePriceNil(b bool)`

 SetSalePriceNil sets the value for SalePrice to be an explicit nil

### UnsetSalePrice
`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) UnsetSalePrice()`

UnsetSalePrice ensures that no value is present for SalePrice, not even an explicit nil
### GetNetProfit

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) GetNetProfit() float32`

GetNetProfit returns the NetProfit field if non-nil, zero value otherwise.

### GetNetProfitOk

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) GetNetProfitOk() (*float32, bool)`

GetNetProfitOk returns a tuple with the NetProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetProfit

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) SetNetProfit(v float32)`

SetNetProfit sets NetProfit field to given value.

### HasNetProfit

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) HasNetProfit() bool`

HasNetProfit returns a boolean if a field has been set.

### SetNetProfitNil

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) SetNetProfitNil(b bool)`

 SetNetProfitNil sets the value for NetProfit to be an explicit nil

### UnsetNetProfit
`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) UnsetNetProfit()`

UnsetNetProfit ensures that no value is present for NetProfit, not even an explicit nil
### GetDetectedAt

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) GetDetectedAt() time.Time`

GetDetectedAt returns the DetectedAt field if non-nil, zero value otherwise.

### GetDetectedAtOk

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) GetDetectedAtOk() (*time.Time, bool)`

GetDetectedAtOk returns a tuple with the DetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedAt

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) SetDetectedAt(v time.Time)`

SetDetectedAt sets DetectedAt field to given value.


### GetItemTitle

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) GetItemTitle() string`

GetItemTitle returns the ItemTitle field if non-nil, zero value otherwise.

### GetItemTitleOk

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) GetItemTitleOk() (*string, bool)`

GetItemTitleOk returns a tuple with the ItemTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTitle

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) SetItemTitle(v string)`

SetItemTitle sets ItemTitle field to given value.

### HasItemTitle

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) HasItemTitle() bool`

HasItemTitle returns a boolean if a field has been set.

### SetItemTitleNil

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) SetItemTitleNil(b bool)`

 SetItemTitleNil sets the value for ItemTitle to be an explicit nil

### UnsetItemTitle
`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) UnsetItemTitle()`

UnsetItemTitle ensures that no value is present for ItemTitle, not even an explicit nil
### GetItemImage

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) GetItemImage() string`

GetItemImage returns the ItemImage field if non-nil, zero value otherwise.

### GetItemImageOk

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) GetItemImageOk() (*string, bool)`

GetItemImageOk returns a tuple with the ItemImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemImage

`func (o *GetAnalyticDashboardResponseTop5ItemsByRevenue) SetItemImage(v string)`

SetItemImage sets ItemImage field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


