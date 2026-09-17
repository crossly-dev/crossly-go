# CreateBuyerCheckoutResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListingSlug** | **string** |  | 
**Quantity** | **float32** |  | 
**TotalCents** | **float32** |  | 
**Currency** | **string** |  | 
**PaymentStatus** | **string** |  | 
**PaymentIntentId** | Pointer to **NullableString** |  | [optional] 
**OrderId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateBuyerCheckoutResponse

`func NewCreateBuyerCheckoutResponse(listingSlug string, quantity float32, totalCents float32, currency string, paymentStatus string, ) *CreateBuyerCheckoutResponse`

NewCreateBuyerCheckoutResponse instantiates a new CreateBuyerCheckoutResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerCheckoutResponseWithDefaults

`func NewCreateBuyerCheckoutResponseWithDefaults() *CreateBuyerCheckoutResponse`

NewCreateBuyerCheckoutResponseWithDefaults instantiates a new CreateBuyerCheckoutResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListingSlug

`func (o *CreateBuyerCheckoutResponse) GetListingSlug() string`

GetListingSlug returns the ListingSlug field if non-nil, zero value otherwise.

### GetListingSlugOk

`func (o *CreateBuyerCheckoutResponse) GetListingSlugOk() (*string, bool)`

GetListingSlugOk returns a tuple with the ListingSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingSlug

`func (o *CreateBuyerCheckoutResponse) SetListingSlug(v string)`

SetListingSlug sets ListingSlug field to given value.


### GetQuantity

`func (o *CreateBuyerCheckoutResponse) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateBuyerCheckoutResponse) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateBuyerCheckoutResponse) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetTotalCents

`func (o *CreateBuyerCheckoutResponse) GetTotalCents() float32`

GetTotalCents returns the TotalCents field if non-nil, zero value otherwise.

### GetTotalCentsOk

`func (o *CreateBuyerCheckoutResponse) GetTotalCentsOk() (*float32, bool)`

GetTotalCentsOk returns a tuple with the TotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCents

`func (o *CreateBuyerCheckoutResponse) SetTotalCents(v float32)`

SetTotalCents sets TotalCents field to given value.


### GetCurrency

`func (o *CreateBuyerCheckoutResponse) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CreateBuyerCheckoutResponse) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CreateBuyerCheckoutResponse) SetCurrency(v string)`

SetCurrency sets Currency field to given value.


### GetPaymentStatus

`func (o *CreateBuyerCheckoutResponse) GetPaymentStatus() string`

GetPaymentStatus returns the PaymentStatus field if non-nil, zero value otherwise.

### GetPaymentStatusOk

`func (o *CreateBuyerCheckoutResponse) GetPaymentStatusOk() (*string, bool)`

GetPaymentStatusOk returns a tuple with the PaymentStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentStatus

`func (o *CreateBuyerCheckoutResponse) SetPaymentStatus(v string)`

SetPaymentStatus sets PaymentStatus field to given value.


### GetPaymentIntentId

`func (o *CreateBuyerCheckoutResponse) GetPaymentIntentId() string`

GetPaymentIntentId returns the PaymentIntentId field if non-nil, zero value otherwise.

### GetPaymentIntentIdOk

`func (o *CreateBuyerCheckoutResponse) GetPaymentIntentIdOk() (*string, bool)`

GetPaymentIntentIdOk returns a tuple with the PaymentIntentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentIntentId

`func (o *CreateBuyerCheckoutResponse) SetPaymentIntentId(v string)`

SetPaymentIntentId sets PaymentIntentId field to given value.

### HasPaymentIntentId

`func (o *CreateBuyerCheckoutResponse) HasPaymentIntentId() bool`

HasPaymentIntentId returns a boolean if a field has been set.

### SetPaymentIntentIdNil

`func (o *CreateBuyerCheckoutResponse) SetPaymentIntentIdNil(b bool)`

 SetPaymentIntentIdNil sets the value for PaymentIntentId to be an explicit nil

### UnsetPaymentIntentId
`func (o *CreateBuyerCheckoutResponse) UnsetPaymentIntentId()`

UnsetPaymentIntentId ensures that no value is present for PaymentIntentId, not even an explicit nil
### GetOrderId

`func (o *CreateBuyerCheckoutResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *CreateBuyerCheckoutResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *CreateBuyerCheckoutResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.

### HasOrderId

`func (o *CreateBuyerCheckoutResponse) HasOrderId() bool`

HasOrderId returns a boolean if a field has been set.

### SetOrderIdNil

`func (o *CreateBuyerCheckoutResponse) SetOrderIdNil(b bool)`

 SetOrderIdNil sets the value for OrderId to be an explicit nil

### UnsetOrderId
`func (o *CreateBuyerCheckoutResponse) UnsetOrderId()`

UnsetOrderId ensures that no value is present for OrderId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


