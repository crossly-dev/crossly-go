# ListSalesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Platform** | **string** |  | 
**CancelledAt** | Pointer to **NullableTime** |  | [optional] 
**OrderId** | Pointer to **NullableString** |  | [optional] 
**ListingId** | Pointer to **NullableString** |  | [optional] 
**InventoryItemId** | Pointer to **NullableString** |  | [optional] 
**PlatformListingId** | Pointer to **NullableString** |  | [optional] 
**PlatformOrderId** | Pointer to **NullableString** |  | [optional] 
**CostOfGoods** | Pointer to **NullableString** |  | [optional] 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 
**SalePrice** | Pointer to **NullableString** |  | [optional] 
**PlatformFee** | Pointer to **NullableString** |  | [optional] 
**ShippingCost** | Pointer to **NullableString** |  | [optional] 
**NetProfit** | Pointer to **NullableString** |  | [optional] 
**TaxCollected** | Pointer to **NullableString** |  | [optional] 
**ShippingCharged** | Pointer to **NullableString** |  | [optional] 
**FeeSource** | Pointer to **NullableString** |  | [optional] 
**DetectedBy** | **string** |  | 
**DetectedAt** | **time.Time** |  | 
**Documents** | Pointer to [**[]ListSalesItemDocuments**](ListSalesItemDocuments.md) |  | [optional] 
**RefundedAmount** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListSalesItem

`func NewListSalesItem(id string, userId string, platform string, detectedBy string, detectedAt time.Time, ) *ListSalesItem`

NewListSalesItem instantiates a new ListSalesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSalesItemWithDefaults

`func NewListSalesItemWithDefaults() *ListSalesItem`

NewListSalesItemWithDefaults instantiates a new ListSalesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSalesItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSalesItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSalesItem) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *ListSalesItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListSalesItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListSalesItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetPlatform

`func (o *ListSalesItem) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListSalesItem) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListSalesItem) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetCancelledAt

`func (o *ListSalesItem) GetCancelledAt() time.Time`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *ListSalesItem) GetCancelledAtOk() (*time.Time, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *ListSalesItem) SetCancelledAt(v time.Time)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *ListSalesItem) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### SetCancelledAtNil

`func (o *ListSalesItem) SetCancelledAtNil(b bool)`

 SetCancelledAtNil sets the value for CancelledAt to be an explicit nil

### UnsetCancelledAt
`func (o *ListSalesItem) UnsetCancelledAt()`

UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
### GetOrderId

`func (o *ListSalesItem) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *ListSalesItem) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *ListSalesItem) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *ListSalesItem) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *ListSalesItem) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *ListSalesItem) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil
### GetListingId

`func (o *ListSalesItem) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *ListSalesItem) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *ListSalesItem) SetListingId(v string)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *ListSalesItem) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### SetListingIdNil

`func (o *ListSalesItem) SetListingIdNil(b bool)`

 SetListingIdNil sets the value for ListingId to be an explicit nil

### UnsetListingId
`func (o *ListSalesItem) UnsetListingId()`

UnsetListingId ensures that no value is present for ListingId, not even an explicit nil
### GetInventoryItemId

`func (o *ListSalesItem) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *ListSalesItem) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *ListSalesItem) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *ListSalesItem) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### SetInventoryItemIdNil

`func (o *ListSalesItem) SetInventoryItemIdNil(b bool)`

 SetInventoryItemIdNil sets the value for InventoryItemId to be an explicit nil

### UnsetInventoryItemId
`func (o *ListSalesItem) UnsetInventoryItemId()`

UnsetInventoryItemId ensures that no value is present for InventoryItemId, not even an explicit nil
### GetPlatformListingId

`func (o *ListSalesItem) GetPlatformListingId() string`

GetPlatformListingId returns the PlatformListingId field if non-nil, zero value otherwise.

### GetPlatformListingIdOk

`func (o *ListSalesItem) GetPlatformListingIdOk() (*string, bool)`

GetPlatformListingIdOk returns a tuple with the PlatformListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformListingId

`func (o *ListSalesItem) SetPlatformListingId(v string)`

SetPlatformListingId sets PlatformListingId field to given value.

### HasPlatformListingId

`func (o *ListSalesItem) HasPlatformListingId() bool`

HasPlatformListingId returns a boolean if a field has been set.

### SetPlatformListingIdNil

`func (o *ListSalesItem) SetPlatformListingIdNil(b bool)`

 SetPlatformListingIdNil sets the value for PlatformListingId to be an explicit nil

### UnsetPlatformListingId
`func (o *ListSalesItem) UnsetPlatformListingId()`

UnsetPlatformListingId ensures that no value is present for PlatformListingId, not even an explicit nil
### GetPlatformOrderId

`func (o *ListSalesItem) GetPlatformOrderId() string`

GetPlatformOrderId returns the PlatformOrderId field if non-nil, zero value otherwise.

### GetPlatformOrderIdOk

`func (o *ListSalesItem) GetPlatformOrderIdOk() (*string, bool)`

GetPlatformOrderIdOk returns a tuple with the PlatformOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformOrderId

`func (o *ListSalesItem) SetPlatformOrderId(v string)`

SetPlatformOrderId sets PlatformOrderId field to given value.

### HasPlatformOrderId

`func (o *ListSalesItem) HasPlatformOrderId() bool`

HasPlatformOrderId returns a boolean if a field has been set.

### SetPlatformOrderIdNil

`func (o *ListSalesItem) SetPlatformOrderIdNil(b bool)`

 SetPlatformOrderIdNil sets the value for PlatformOrderId to be an explicit nil

### UnsetPlatformOrderId
`func (o *ListSalesItem) UnsetPlatformOrderId()`

UnsetPlatformOrderId ensures that no value is present for PlatformOrderId, not even an explicit nil
### GetCostOfGoods

`func (o *ListSalesItem) GetCostOfGoods() string`

GetCostOfGoods returns the CostOfGoods field if non-nil, zero value otherwise.

### GetCostOfGoodsOk

`func (o *ListSalesItem) GetCostOfGoodsOk() (*string, bool)`

GetCostOfGoodsOk returns a tuple with the CostOfGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOfGoods

`func (o *ListSalesItem) SetCostOfGoods(v string)`

SetCostOfGoods sets CostOfGoods field to given value.

### HasCostOfGoods

`func (o *ListSalesItem) HasCostOfGoods() bool`

HasCostOfGoods returns a boolean if a field has been set.

### SetCostOfGoodsNil

`func (o *ListSalesItem) SetCostOfGoodsNil(b bool)`

 SetCostOfGoodsNil sets the value for CostOfGoods to be an explicit nil

### UnsetCostOfGoods
`func (o *ListSalesItem) UnsetCostOfGoods()`

UnsetCostOfGoods ensures that no value is present for CostOfGoods, not even an explicit nil
### GetDeletedAt

`func (o *ListSalesItem) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ListSalesItem) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ListSalesItem) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ListSalesItem) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *ListSalesItem) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ListSalesItem) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil
### GetSalePrice

`func (o *ListSalesItem) GetSalePrice() string`

GetSalePrice returns the SalePrice field if non-nil, zero value otherwise.

### GetSalePriceOk

`func (o *ListSalesItem) GetSalePriceOk() (*string, bool)`

GetSalePriceOk returns a tuple with the SalePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalePrice

`func (o *ListSalesItem) SetSalePrice(v string)`

SetSalePrice sets SalePrice field to given value.

### HasSalePrice

`func (o *ListSalesItem) HasSalePrice() bool`

HasSalePrice returns a boolean if a field has been set.

### SetSalePriceNil

`func (o *ListSalesItem) SetSalePriceNil(b bool)`

 SetSalePriceNil sets the value for SalePrice to be an explicit nil

### UnsetSalePrice
`func (o *ListSalesItem) UnsetSalePrice()`

UnsetSalePrice ensures that no value is present for SalePrice, not even an explicit nil
### GetPlatformFee

`func (o *ListSalesItem) GetPlatformFee() string`

GetPlatformFee returns the PlatformFee field if non-nil, zero value otherwise.

### GetPlatformFeeOk

`func (o *ListSalesItem) GetPlatformFeeOk() (*string, bool)`

GetPlatformFeeOk returns a tuple with the PlatformFee field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformFee

`func (o *ListSalesItem) SetPlatformFee(v string)`

SetPlatformFee sets PlatformFee field to given value.

### HasPlatformFee

`func (o *ListSalesItem) HasPlatformFee() bool`

HasPlatformFee returns a boolean if a field has been set.

### SetPlatformFeeNil

`func (o *ListSalesItem) SetPlatformFeeNil(b bool)`

 SetPlatformFeeNil sets the value for PlatformFee to be an explicit nil

### UnsetPlatformFee
`func (o *ListSalesItem) UnsetPlatformFee()`

UnsetPlatformFee ensures that no value is present for PlatformFee, not even an explicit nil
### GetShippingCost

`func (o *ListSalesItem) GetShippingCost() string`

GetShippingCost returns the ShippingCost field if non-nil, zero value otherwise.

### GetShippingCostOk

`func (o *ListSalesItem) GetShippingCostOk() (*string, bool)`

GetShippingCostOk returns a tuple with the ShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCost

`func (o *ListSalesItem) SetShippingCost(v string)`

SetShippingCost sets ShippingCost field to given value.

### HasShippingCost

`func (o *ListSalesItem) HasShippingCost() bool`

HasShippingCost returns a boolean if a field has been set.

### SetShippingCostNil

`func (o *ListSalesItem) SetShippingCostNil(b bool)`

 SetShippingCostNil sets the value for ShippingCost to be an explicit nil

### UnsetShippingCost
`func (o *ListSalesItem) UnsetShippingCost()`

UnsetShippingCost ensures that no value is present for ShippingCost, not even an explicit nil
### GetNetProfit

`func (o *ListSalesItem) GetNetProfit() string`

GetNetProfit returns the NetProfit field if non-nil, zero value otherwise.

### GetNetProfitOk

`func (o *ListSalesItem) GetNetProfitOk() (*string, bool)`

GetNetProfitOk returns a tuple with the NetProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetProfit

`func (o *ListSalesItem) SetNetProfit(v string)`

SetNetProfit sets NetProfit field to given value.

### HasNetProfit

`func (o *ListSalesItem) HasNetProfit() bool`

HasNetProfit returns a boolean if a field has been set.

### SetNetProfitNil

`func (o *ListSalesItem) SetNetProfitNil(b bool)`

 SetNetProfitNil sets the value for NetProfit to be an explicit nil

### UnsetNetProfit
`func (o *ListSalesItem) UnsetNetProfit()`

UnsetNetProfit ensures that no value is present for NetProfit, not even an explicit nil
### GetTaxCollected

`func (o *ListSalesItem) GetTaxCollected() string`

GetTaxCollected returns the TaxCollected field if non-nil, zero value otherwise.

### GetTaxCollectedOk

`func (o *ListSalesItem) GetTaxCollectedOk() (*string, bool)`

GetTaxCollectedOk returns a tuple with the TaxCollected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCollected

`func (o *ListSalesItem) SetTaxCollected(v string)`

SetTaxCollected sets TaxCollected field to given value.

### HasTaxCollected

`func (o *ListSalesItem) HasTaxCollected() bool`

HasTaxCollected returns a boolean if a field has been set.

### SetTaxCollectedNil

`func (o *ListSalesItem) SetTaxCollectedNil(b bool)`

 SetTaxCollectedNil sets the value for TaxCollected to be an explicit nil

### UnsetTaxCollected
`func (o *ListSalesItem) UnsetTaxCollected()`

UnsetTaxCollected ensures that no value is present for TaxCollected, not even an explicit nil
### GetShippingCharged

`func (o *ListSalesItem) GetShippingCharged() string`

GetShippingCharged returns the ShippingCharged field if non-nil, zero value otherwise.

### GetShippingChargedOk

`func (o *ListSalesItem) GetShippingChargedOk() (*string, bool)`

GetShippingChargedOk returns a tuple with the ShippingCharged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCharged

`func (o *ListSalesItem) SetShippingCharged(v string)`

SetShippingCharged sets ShippingCharged field to given value.

### HasShippingCharged

`func (o *ListSalesItem) HasShippingCharged() bool`

HasShippingCharged returns a boolean if a field has been set.

### SetShippingChargedNil

`func (o *ListSalesItem) SetShippingChargedNil(b bool)`

 SetShippingChargedNil sets the value for ShippingCharged to be an explicit nil

### UnsetShippingCharged
`func (o *ListSalesItem) UnsetShippingCharged()`

UnsetShippingCharged ensures that no value is present for ShippingCharged, not even an explicit nil
### GetFeeSource

`func (o *ListSalesItem) GetFeeSource() string`

GetFeeSource returns the FeeSource field if non-nil, zero value otherwise.

### GetFeeSourceOk

`func (o *ListSalesItem) GetFeeSourceOk() (*string, bool)`

GetFeeSourceOk returns a tuple with the FeeSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeeSource

`func (o *ListSalesItem) SetFeeSource(v string)`

SetFeeSource sets FeeSource field to given value.

### HasFeeSource

`func (o *ListSalesItem) HasFeeSource() bool`

HasFeeSource returns a boolean if a field has been set.

### SetFeeSourceNil

`func (o *ListSalesItem) SetFeeSourceNil(b bool)`

 SetFeeSourceNil sets the value for FeeSource to be an explicit nil

### UnsetFeeSource
`func (o *ListSalesItem) UnsetFeeSource()`

UnsetFeeSource ensures that no value is present for FeeSource, not even an explicit nil
### GetDetectedBy

`func (o *ListSalesItem) GetDetectedBy() string`

GetDetectedBy returns the DetectedBy field if non-nil, zero value otherwise.

### GetDetectedByOk

`func (o *ListSalesItem) GetDetectedByOk() (*string, bool)`

GetDetectedByOk returns a tuple with the DetectedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedBy

`func (o *ListSalesItem) SetDetectedBy(v string)`

SetDetectedBy sets DetectedBy field to given value.


### GetDetectedAt

`func (o *ListSalesItem) GetDetectedAt() time.Time`

GetDetectedAt returns the DetectedAt field if non-nil, zero value otherwise.

### GetDetectedAtOk

`func (o *ListSalesItem) GetDetectedAtOk() (*time.Time, bool)`

GetDetectedAtOk returns a tuple with the DetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetectedAt

`func (o *ListSalesItem) SetDetectedAt(v time.Time)`

SetDetectedAt sets DetectedAt field to given value.


### GetDocuments

`func (o *ListSalesItem) GetDocuments() []ListSalesItemDocuments`

GetDocuments returns the Documents field if non-nil, zero value otherwise.

### GetDocumentsOk

`func (o *ListSalesItem) GetDocumentsOk() (*[]ListSalesItemDocuments, bool)`

GetDocumentsOk returns a tuple with the Documents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocuments

`func (o *ListSalesItem) SetDocuments(v []ListSalesItemDocuments)`

SetDocuments sets Documents field to given value.

### HasDocuments

`func (o *ListSalesItem) HasDocuments() bool`

HasDocuments returns a boolean if a field has been set.

### SetDocumentsNil

`func (o *ListSalesItem) SetDocumentsNil(b bool)`

 SetDocumentsNil sets the value for Documents to be an explicit nil

### UnsetDocuments
`func (o *ListSalesItem) UnsetDocuments()`

UnsetDocuments ensures that no value is present for Documents, not even an explicit nil
### GetRefundedAmount

`func (o *ListSalesItem) GetRefundedAmount() string`

GetRefundedAmount returns the RefundedAmount field if non-nil, zero value otherwise.

### GetRefundedAmountOk

`func (o *ListSalesItem) GetRefundedAmountOk() (*string, bool)`

GetRefundedAmountOk returns a tuple with the RefundedAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAmount

`func (o *ListSalesItem) SetRefundedAmount(v string)`

SetRefundedAmount sets RefundedAmount field to given value.

### HasRefundedAmount

`func (o *ListSalesItem) HasRefundedAmount() bool`

HasRefundedAmount returns a boolean if a field has been set.

### SetRefundedAmountNil

`func (o *ListSalesItem) SetRefundedAmountNil(b bool)`

 SetRefundedAmountNil sets the value for RefundedAmount to be an explicit nil

### UnsetRefundedAmount
`func (o *ListSalesItem) UnsetRefundedAmount()`

UnsetRefundedAmount ensures that no value is present for RefundedAmount, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


