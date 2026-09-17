# CreateBuyerCartQuoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Kind** | **string** |  | 
**TaxCents** | **float32** |  | 
**ShippingCents** | **float32** |  | 
**TotalCents** | **float32** |  | 
**Currency** | **string** |  | 
**ItemsTotalCents** | **float32** |  | 
**PickupCartItemIds** | **[]string** | Lines being collected in person, so a summary can name what ships free. | 
**TaxComplete** | **bool** | False means there is no saved delivery address, so &#x60;taxCents&#x60; is a floor rather than a final figure — not that tax is zero. | 

## Methods

### NewCreateBuyerCartQuoteResponse

`func NewCreateBuyerCartQuoteResponse(kind string, taxCents float32, shippingCents float32, totalCents float32, currency string, itemsTotalCents float32, pickupCartItemIds []string, taxComplete bool, ) *CreateBuyerCartQuoteResponse`

NewCreateBuyerCartQuoteResponse instantiates a new CreateBuyerCartQuoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerCartQuoteResponseWithDefaults

`func NewCreateBuyerCartQuoteResponseWithDefaults() *CreateBuyerCartQuoteResponse`

NewCreateBuyerCartQuoteResponseWithDefaults instantiates a new CreateBuyerCartQuoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKind

`func (o *CreateBuyerCartQuoteResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *CreateBuyerCartQuoteResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *CreateBuyerCartQuoteResponse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetTaxCents

`func (o *CreateBuyerCartQuoteResponse) GetTaxCents() float32`

GetTaxCents returns the TaxCents field if non-nil, zero value otherwise.

### GetTaxCentsOk

`func (o *CreateBuyerCartQuoteResponse) GetTaxCentsOk() (*float32, bool)`

GetTaxCentsOk returns a tuple with the TaxCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxCents

`func (o *CreateBuyerCartQuoteResponse) SetTaxCents(v float32)`

SetTaxCents sets TaxCents field to given value.


### GetShippingCents

`func (o *CreateBuyerCartQuoteResponse) GetShippingCents() float32`

GetShippingCents returns the ShippingCents field if non-nil, zero value otherwise.

### GetShippingCentsOk

`func (o *CreateBuyerCartQuoteResponse) GetShippingCentsOk() (*float32, bool)`

GetShippingCentsOk returns a tuple with the ShippingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingCents

`func (o *CreateBuyerCartQuoteResponse) SetShippingCents(v float32)`

SetShippingCents sets ShippingCents field to given value.


### GetTotalCents

`func (o *CreateBuyerCartQuoteResponse) GetTotalCents() float32`

GetTotalCents returns the TotalCents field if non-nil, zero value otherwise.

### GetTotalCentsOk

`func (o *CreateBuyerCartQuoteResponse) GetTotalCentsOk() (*float32, bool)`

GetTotalCentsOk returns a tuple with the TotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCents

`func (o *CreateBuyerCartQuoteResponse) SetTotalCents(v float32)`

SetTotalCents sets TotalCents field to given value.


### GetCurrency

`func (o *CreateBuyerCartQuoteResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateBuyerCartQuoteResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateBuyerCartQuoteResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetItemsTotalCents

`func (o *CreateBuyerCartQuoteResponse) GetItemsTotalCents() float32`

GetItemsTotalCents returns the ItemsTotalCents field if non-nil, zero value otherwise.

### GetItemsTotalCentsOk

`func (o *CreateBuyerCartQuoteResponse) GetItemsTotalCentsOk() (*float32, bool)`

GetItemsTotalCentsOk returns a tuple with the ItemsTotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemsTotalCents

`func (o *CreateBuyerCartQuoteResponse) SetItemsTotalCents(v float32)`

SetItemsTotalCents sets ItemsTotalCents field to given value.


### GetPickupCartItemIds

`func (o *CreateBuyerCartQuoteResponse) GetPickupCartItemIds() []string`

GetPickupCartItemIds returns the PickupCartItemIds field if non-nil, zero value otherwise.

### GetPickupCartItemIdsOk

`func (o *CreateBuyerCartQuoteResponse) GetPickupCartItemIdsOk() (*[]string, bool)`

GetPickupCartItemIdsOk returns a tuple with the PickupCartItemIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupCartItemIds

`func (o *CreateBuyerCartQuoteResponse) SetPickupCartItemIds(v []string)`

SetPickupCartItemIds sets PickupCartItemIds field to given value.


### GetTaxComplete

`func (o *CreateBuyerCartQuoteResponse) GetTaxComplete() bool`

GetTaxComplete returns the TaxComplete field if non-nil, zero value otherwise.

### GetTaxCompleteOk

`func (o *CreateBuyerCartQuoteResponse) GetTaxCompleteOk() (*bool, bool)`

GetTaxCompleteOk returns a tuple with the TaxComplete field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxComplete

`func (o *CreateBuyerCartQuoteResponse) SetTaxComplete(v bool)`

SetTaxComplete sets TaxComplete field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


