# GetCustomerResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NormalizedHandle** | **string** |  | 
**Platforms** | **[]string** |  | 
**TotalOrders** | **float32** |  | 
**TotalSpent** | **float32** |  | 
**Orders** | [**[]GetCustomerResponseOrders**](GetCustomerResponseOrders.md) |  | 
**Contact** | [**GetCustomerResponseContact**](GetCustomerResponseContact.md) |  | 

## Methods

### NewGetCustomerResponse

`func NewGetCustomerResponse(normalizedHandle string, platforms []string, totalOrders float32, totalSpent float32, orders []GetCustomerResponseOrders, contact GetCustomerResponseContact, ) *GetCustomerResponse`

NewGetCustomerResponse instantiates a new GetCustomerResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerResponseWithDefaults

`func NewGetCustomerResponseWithDefaults() *GetCustomerResponse`

NewGetCustomerResponseWithDefaults instantiates a new GetCustomerResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNormalizedHandle

`func (o *GetCustomerResponse) GetNormalizedHandle() string`

GetNormalizedHandle returns the NormalizedHandle field if non-nil, zero value otherwise.

### GetNormalizedHandleOk

`func (o *GetCustomerResponse) GetNormalizedHandleOk() (*string, bool)`

GetNormalizedHandleOk returns a tuple with the NormalizedHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedHandle

`func (o *GetCustomerResponse) SetNormalizedHandle(v string)`

SetNormalizedHandle sets NormalizedHandle field to given value.


### GetPlatforms

`func (o *GetCustomerResponse) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *GetCustomerResponse) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *GetCustomerResponse) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.


### GetTotalOrders

`func (o *GetCustomerResponse) GetTotalOrders() float32`

GetTotalOrders returns the TotalOrders field if non-nil, zero value otherwise.

### GetTotalOrdersOk

`func (o *GetCustomerResponse) GetTotalOrdersOk() (*float32, bool)`

GetTotalOrdersOk returns a tuple with the TotalOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOrders

`func (o *GetCustomerResponse) SetTotalOrders(v float32)`

SetTotalOrders sets TotalOrders field to given value.


### GetTotalSpent

`func (o *GetCustomerResponse) GetTotalSpent() float32`

GetTotalSpent returns the TotalSpent field if non-nil, zero value otherwise.

### GetTotalSpentOk

`func (o *GetCustomerResponse) GetTotalSpentOk() (*float32, bool)`

GetTotalSpentOk returns a tuple with the TotalSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpent

`func (o *GetCustomerResponse) SetTotalSpent(v float32)`

SetTotalSpent sets TotalSpent field to given value.


### GetOrders

`func (o *GetCustomerResponse) GetOrders() []GetCustomerResponseOrders`

GetOrders returns the Orders field if non-nil, zero value otherwise.

### GetOrdersOk

`func (o *GetCustomerResponse) GetOrdersOk() (*[]GetCustomerResponseOrders, bool)`

GetOrdersOk returns a tuple with the Orders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrders

`func (o *GetCustomerResponse) SetOrders(v []GetCustomerResponseOrders)`

SetOrders sets Orders field to given value.


### GetContact

`func (o *GetCustomerResponse) GetContact() GetCustomerResponseContact`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *GetCustomerResponse) GetContactOk() (*GetCustomerResponseContact, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *GetCustomerResponse) SetContact(v GetCustomerResponseContact)`

SetContact sets Contact field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


