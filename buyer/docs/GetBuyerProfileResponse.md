# GetBuyerProfileResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Email** | **string** |  | 
**DisplayName** | Pointer to **NullableString** |  | [optional] 
**ShippingAddress** | Pointer to **map[string]interface{}** |  | [optional] 
**BucksBalanceCents** | **float32** |  | 

## Methods

### NewGetBuyerProfileResponse

`func NewGetBuyerProfileResponse(id string, email string, bucksBalanceCents float32, ) *GetBuyerProfileResponse`

NewGetBuyerProfileResponse instantiates a new GetBuyerProfileResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBuyerProfileResponseWithDefaults

`func NewGetBuyerProfileResponseWithDefaults() *GetBuyerProfileResponse`

NewGetBuyerProfileResponseWithDefaults instantiates a new GetBuyerProfileResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetBuyerProfileResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBuyerProfileResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBuyerProfileResponse) SetId(v string)`

SetId sets Id field to given value.


### GetEmail

`func (o *GetBuyerProfileResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetBuyerProfileResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetBuyerProfileResponse) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetDisplayName

`func (o *GetBuyerProfileResponse) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *GetBuyerProfileResponse) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *GetBuyerProfileResponse) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *GetBuyerProfileResponse) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### SetDisplayNameNil

`func (o *GetBuyerProfileResponse) SetDisplayNameNil(b bool)`

 SetDisplayNameNil sets the value for DisplayName to be an explicit nil

### UnsetDisplayName
`func (o *GetBuyerProfileResponse) UnsetDisplayName()`

UnsetDisplayName ensures that no value is present for DisplayName, not even an explicit nil
### GetShippingAddress

`func (o *GetBuyerProfileResponse) GetShippingAddress() map[string]interface{}`

GetShippingAddress returns the ShippingAddress field if non-nil, zero value otherwise.

### GetShippingAddressOk

`func (o *GetBuyerProfileResponse) GetShippingAddressOk() (*map[string]interface{}, bool)`

GetShippingAddressOk returns a tuple with the ShippingAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingAddress

`func (o *GetBuyerProfileResponse) SetShippingAddress(v map[string]interface{})`

SetShippingAddress sets ShippingAddress field to given value.

### HasShippingAddress

`func (o *GetBuyerProfileResponse) HasShippingAddress() bool`

HasShippingAddress returns a boolean if a field has been set.

### SetShippingAddressNil

`func (o *GetBuyerProfileResponse) SetShippingAddressNil(b bool)`

 SetShippingAddressNil sets the value for ShippingAddress to be an explicit nil

### UnsetShippingAddress
`func (o *GetBuyerProfileResponse) UnsetShippingAddress()`

UnsetShippingAddress ensures that no value is present for ShippingAddress, not even an explicit nil
### GetBucksBalanceCents

`func (o *GetBuyerProfileResponse) GetBucksBalanceCents() float32`

GetBucksBalanceCents returns the BucksBalanceCents field if non-nil, zero value otherwise.

### GetBucksBalanceCentsOk

`func (o *GetBuyerProfileResponse) GetBucksBalanceCentsOk() (*float32, bool)`

GetBucksBalanceCentsOk returns a tuple with the BucksBalanceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucksBalanceCents

`func (o *GetBuyerProfileResponse) SetBucksBalanceCents(v float32)`

SetBucksBalanceCents sets BucksBalanceCents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


