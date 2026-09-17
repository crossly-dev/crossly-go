# GetMarketProductResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Product** | [**GetMarketProductResponseProduct**](GetMarketProductResponseProduct.md) |  | 
**Variants** | [**[]GetMarketProductResponseVariants**](GetMarketProductResponseVariants.md) |  | 

## Methods

### NewGetMarketProductResponse

`func NewGetMarketProductResponse(product GetMarketProductResponseProduct, variants []GetMarketProductResponseVariants, ) *GetMarketProductResponse`

NewGetMarketProductResponse instantiates a new GetMarketProductResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketProductResponseWithDefaults

`func NewGetMarketProductResponseWithDefaults() *GetMarketProductResponse`

NewGetMarketProductResponseWithDefaults instantiates a new GetMarketProductResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProduct

`func (o *GetMarketProductResponse) GetProduct() GetMarketProductResponseProduct`

GetProduct returns the Product field if non-nil, zero value otherwise.

### GetProductOk

`func (o *GetMarketProductResponse) GetProductOk() (*GetMarketProductResponseProduct, bool)`

GetProductOk returns a tuple with the Product field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProduct

`func (o *GetMarketProductResponse) SetProduct(v GetMarketProductResponseProduct)`

SetProduct sets Product field to given value.


### GetVariants

`func (o *GetMarketProductResponse) GetVariants() []GetMarketProductResponseVariants`

GetVariants returns the Variants field if non-nil, zero value otherwise.

### GetVariantsOk

`func (o *GetMarketProductResponse) GetVariantsOk() (*[]GetMarketProductResponseVariants, bool)`

GetVariantsOk returns a tuple with the Variants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariants

`func (o *GetMarketProductResponse) SetVariants(v []GetMarketProductResponseVariants)`

SetVariants sets Variants field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


