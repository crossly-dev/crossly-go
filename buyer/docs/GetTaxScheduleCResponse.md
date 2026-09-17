# GetTaxScheduleCResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Year** | **float32** |  | 
**BusinessName** | Pointer to **NullableString** |  | [optional] 
**Ein** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to **NullableString** |  | [optional] 
**GrossReceipts** | **float32** |  | 
**ReturnsAndAllowances** | **float32** |  | 
**NetSales** | **float32** |  | 
**Expenses** | [**GetTaxScheduleCResponseExpenses**](GetTaxScheduleCResponseExpenses.md) |  | 
**NetProfit** | **float32** |  | 
**Counts** | [**GetTaxScheduleCResponseCounts**](GetTaxScheduleCResponseCounts.md) |  | 

## Methods

### NewGetTaxScheduleCResponse

`func NewGetTaxScheduleCResponse(year float32, grossReceipts float32, returnsAndAllowances float32, netSales float32, expenses GetTaxScheduleCResponseExpenses, netProfit float32, counts GetTaxScheduleCResponseCounts, ) *GetTaxScheduleCResponse`

NewGetTaxScheduleCResponse instantiates a new GetTaxScheduleCResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaxScheduleCResponseWithDefaults

`func NewGetTaxScheduleCResponseWithDefaults() *GetTaxScheduleCResponse`

NewGetTaxScheduleCResponseWithDefaults instantiates a new GetTaxScheduleCResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetYear

`func (o *GetTaxScheduleCResponse) GetYear() float32`

GetYear returns the Year field if non-nil, zero value otherwise.

### GetYearOk

`func (o *GetTaxScheduleCResponse) GetYearOk() (*float32, bool)`

GetYearOk returns a tuple with the Year field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetYear

`func (o *GetTaxScheduleCResponse) SetYear(v float32)`

SetYear sets Year field to given value.


### GetBusinessName

`func (o *GetTaxScheduleCResponse) GetBusinessName() string`

GetBusinessName returns the BusinessName field if non-nil, zero value otherwise.

### GetBusinessNameOk

`func (o *GetTaxScheduleCResponse) GetBusinessNameOk() (*string, bool)`

GetBusinessNameOk returns a tuple with the BusinessName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessName

`func (o *GetTaxScheduleCResponse) SetBusinessName(v string)`

SetBusinessName sets BusinessName field to given value.

### HasBusinessName

`func (o *GetTaxScheduleCResponse) HasBusinessName() bool`

HasBusinessName returns a boolean if a field has been set.

### SetBusinessNameNil

`func (o *GetTaxScheduleCResponse) SetBusinessNameNil(b bool)`

 SetBusinessNameNil sets the value for BusinessName to be an explicit nil

### UnsetBusinessName
`func (o *GetTaxScheduleCResponse) UnsetBusinessName()`

UnsetBusinessName ensures that no value is present for BusinessName, not even an explicit nil
### GetEin

`func (o *GetTaxScheduleCResponse) GetEin() string`

GetEin returns the Ein field if non-nil, zero value otherwise.

### GetEinOk

`func (o *GetTaxScheduleCResponse) GetEinOk() (*string, bool)`

GetEinOk returns a tuple with the Ein field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEin

`func (o *GetTaxScheduleCResponse) SetEin(v string)`

SetEin sets Ein field to given value.

### HasEin

`func (o *GetTaxScheduleCResponse) HasEin() bool`

HasEin returns a boolean if a field has been set.

### SetEinNil

`func (o *GetTaxScheduleCResponse) SetEinNil(b bool)`

 SetEinNil sets the value for Ein to be an explicit nil

### UnsetEin
`func (o *GetTaxScheduleCResponse) UnsetEin()`

UnsetEin ensures that no value is present for Ein, not even an explicit nil
### GetAddress

`func (o *GetTaxScheduleCResponse) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetTaxScheduleCResponse) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetTaxScheduleCResponse) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetTaxScheduleCResponse) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *GetTaxScheduleCResponse) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *GetTaxScheduleCResponse) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetGrossReceipts

`func (o *GetTaxScheduleCResponse) GetGrossReceipts() float32`

GetGrossReceipts returns the GrossReceipts field if non-nil, zero value otherwise.

### GetGrossReceiptsOk

`func (o *GetTaxScheduleCResponse) GetGrossReceiptsOk() (*float32, bool)`

GetGrossReceiptsOk returns a tuple with the GrossReceipts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrossReceipts

`func (o *GetTaxScheduleCResponse) SetGrossReceipts(v float32)`

SetGrossReceipts sets GrossReceipts field to given value.


### GetReturnsAndAllowances

`func (o *GetTaxScheduleCResponse) GetReturnsAndAllowances() float32`

GetReturnsAndAllowances returns the ReturnsAndAllowances field if non-nil, zero value otherwise.

### GetReturnsAndAllowancesOk

`func (o *GetTaxScheduleCResponse) GetReturnsAndAllowancesOk() (*float32, bool)`

GetReturnsAndAllowancesOk returns a tuple with the ReturnsAndAllowances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnsAndAllowances

`func (o *GetTaxScheduleCResponse) SetReturnsAndAllowances(v float32)`

SetReturnsAndAllowances sets ReturnsAndAllowances field to given value.


### GetNetSales

`func (o *GetTaxScheduleCResponse) GetNetSales() float32`

GetNetSales returns the NetSales field if non-nil, zero value otherwise.

### GetNetSalesOk

`func (o *GetTaxScheduleCResponse) GetNetSalesOk() (*float32, bool)`

GetNetSalesOk returns a tuple with the NetSales field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetSales

`func (o *GetTaxScheduleCResponse) SetNetSales(v float32)`

SetNetSales sets NetSales field to given value.


### GetExpenses

`func (o *GetTaxScheduleCResponse) GetExpenses() GetTaxScheduleCResponseExpenses`

GetExpenses returns the Expenses field if non-nil, zero value otherwise.

### GetExpensesOk

`func (o *GetTaxScheduleCResponse) GetExpensesOk() (*GetTaxScheduleCResponseExpenses, bool)`

GetExpensesOk returns a tuple with the Expenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpenses

`func (o *GetTaxScheduleCResponse) SetExpenses(v GetTaxScheduleCResponseExpenses)`

SetExpenses sets Expenses field to given value.


### GetNetProfit

`func (o *GetTaxScheduleCResponse) GetNetProfit() float32`

GetNetProfit returns the NetProfit field if non-nil, zero value otherwise.

### GetNetProfitOk

`func (o *GetTaxScheduleCResponse) GetNetProfitOk() (*float32, bool)`

GetNetProfitOk returns a tuple with the NetProfit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetProfit

`func (o *GetTaxScheduleCResponse) SetNetProfit(v float32)`

SetNetProfit sets NetProfit field to given value.


### GetCounts

`func (o *GetTaxScheduleCResponse) GetCounts() GetTaxScheduleCResponseCounts`

GetCounts returns the Counts field if non-nil, zero value otherwise.

### GetCountsOk

`func (o *GetTaxScheduleCResponse) GetCountsOk() (*GetTaxScheduleCResponseCounts, bool)`

GetCountsOk returns a tuple with the Counts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounts

`func (o *GetTaxScheduleCResponse) SetCounts(v GetTaxScheduleCResponseCounts)`

SetCounts sets Counts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


