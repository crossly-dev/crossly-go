# ListCustomersItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NormalizedHandle** | **string** |  | 
**Platforms** | **[]string** |  | 
**TotalOrders** | **float32** |  | 
**TotalSpent** | **float32** |  | 

## Methods

### NewListCustomersItem

`func NewListCustomersItem(normalizedHandle string, platforms []string, totalOrders float32, totalSpent float32, ) *ListCustomersItem`

NewListCustomersItem instantiates a new ListCustomersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCustomersItemWithDefaults

`func NewListCustomersItemWithDefaults() *ListCustomersItem`

NewListCustomersItemWithDefaults instantiates a new ListCustomersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNormalizedHandle

`func (o *ListCustomersItem) GetNormalizedHandle() string`

GetNormalizedHandle returns the NormalizedHandle field if non-nil, zero value otherwise.

### GetNormalizedHandleOk

`func (o *ListCustomersItem) GetNormalizedHandleOk() (*string, bool)`

GetNormalizedHandleOk returns a tuple with the NormalizedHandle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormalizedHandle

`func (o *ListCustomersItem) SetNormalizedHandle(v string)`

SetNormalizedHandle sets NormalizedHandle field to given value.


### GetPlatforms

`func (o *ListCustomersItem) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *ListCustomersItem) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *ListCustomersItem) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.


### GetTotalOrders

`func (o *ListCustomersItem) GetTotalOrders() float32`

GetTotalOrders returns the TotalOrders field if non-nil, zero value otherwise.

### GetTotalOrdersOk

`func (o *ListCustomersItem) GetTotalOrdersOk() (*float32, bool)`

GetTotalOrdersOk returns a tuple with the TotalOrders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalOrders

`func (o *ListCustomersItem) SetTotalOrders(v float32)`

SetTotalOrders sets TotalOrders field to given value.


### GetTotalSpent

`func (o *ListCustomersItem) GetTotalSpent() float32`

GetTotalSpent returns the TotalSpent field if non-nil, zero value otherwise.

### GetTotalSpentOk

`func (o *ListCustomersItem) GetTotalSpentOk() (*float32, bool)`

GetTotalSpentOk returns a tuple with the TotalSpent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalSpent

`func (o *ListCustomersItem) SetTotalSpent(v float32)`

SetTotalSpent sets TotalSpent field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


