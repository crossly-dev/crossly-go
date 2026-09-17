# ListBuyerOrdersItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Status** | **string** |  | 
**BuyerPaidCents** | **float32** |  | 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewListBuyerOrdersItem

`func NewListBuyerOrdersItem(id string, status string, buyerPaidCents float32, createdAt time.Time, ) *ListBuyerOrdersItem`

NewListBuyerOrdersItem instantiates a new ListBuyerOrdersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBuyerOrdersItemWithDefaults

`func NewListBuyerOrdersItemWithDefaults() *ListBuyerOrdersItem`

NewListBuyerOrdersItemWithDefaults instantiates a new ListBuyerOrdersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListBuyerOrdersItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBuyerOrdersItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBuyerOrdersItem) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *ListBuyerOrdersItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListBuyerOrdersItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListBuyerOrdersItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBuyerPaidCents

`func (o *ListBuyerOrdersItem) GetBuyerPaidCents() float32`

GetBuyerPaidCents returns the BuyerPaidCents field if non-nil, zero value otherwise.

### GetBuyerPaidCentsOk

`func (o *ListBuyerOrdersItem) GetBuyerPaidCentsOk() (*float32, bool)`

GetBuyerPaidCentsOk returns a tuple with the BuyerPaidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerPaidCents

`func (o *ListBuyerOrdersItem) SetBuyerPaidCents(v float32)`

SetBuyerPaidCents sets BuyerPaidCents field to given value.


### GetCreatedAt

`func (o *ListBuyerOrdersItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListBuyerOrdersItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListBuyerOrdersItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


