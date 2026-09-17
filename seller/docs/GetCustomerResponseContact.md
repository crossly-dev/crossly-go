# GetCustomerResponseContact

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Address** | Pointer to [**NullableGetCustomerResponseContactAddress**](GetCustomerResponseContactAddress.md) |  | [optional] 
**AddressPlatform** | Pointer to **NullableString** | Where the address came from, so the UI can say. | [optional] 

## Methods

### NewGetCustomerResponseContact

`func NewGetCustomerResponseContact() *GetCustomerResponseContact`

NewGetCustomerResponseContact instantiates a new GetCustomerResponseContact object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCustomerResponseContactWithDefaults

`func NewGetCustomerResponseContactWithDefaults() *GetCustomerResponseContact`

NewGetCustomerResponseContactWithDefaults instantiates a new GetCustomerResponseContact object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *GetCustomerResponseContact) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *GetCustomerResponseContact) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *GetCustomerResponseContact) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *GetCustomerResponseContact) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *GetCustomerResponseContact) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *GetCustomerResponseContact) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetName

`func (o *GetCustomerResponseContact) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCustomerResponseContact) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCustomerResponseContact) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCustomerResponseContact) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GetCustomerResponseContact) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GetCustomerResponseContact) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetPhone

`func (o *GetCustomerResponseContact) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *GetCustomerResponseContact) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *GetCustomerResponseContact) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *GetCustomerResponseContact) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *GetCustomerResponseContact) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *GetCustomerResponseContact) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetAddress

`func (o *GetCustomerResponseContact) GetAddress() GetCustomerResponseContactAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *GetCustomerResponseContact) GetAddressOk() (*GetCustomerResponseContactAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *GetCustomerResponseContact) SetAddress(v GetCustomerResponseContactAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *GetCustomerResponseContact) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### SetAddressNil

`func (o *GetCustomerResponseContact) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *GetCustomerResponseContact) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetAddressPlatform

`func (o *GetCustomerResponseContact) GetAddressPlatform() string`

GetAddressPlatform returns the AddressPlatform field if non-nil, zero value otherwise.

### GetAddressPlatformOk

`func (o *GetCustomerResponseContact) GetAddressPlatformOk() (*string, bool)`

GetAddressPlatformOk returns a tuple with the AddressPlatform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressPlatform

`func (o *GetCustomerResponseContact) SetAddressPlatform(v string)`

SetAddressPlatform sets AddressPlatform field to given value.

### HasAddressPlatform

`func (o *GetCustomerResponseContact) HasAddressPlatform() bool`

HasAddressPlatform returns a boolean if a field has been set.

### SetAddressPlatformNil

`func (o *GetCustomerResponseContact) SetAddressPlatformNil(b bool)`

 SetAddressPlatformNil sets the value for AddressPlatform to be an explicit nil

### UnsetAddressPlatform
`func (o *GetCustomerResponseContact) UnsetAddressPlatform()`

UnsetAddressPlatform ensures that no value is present for AddressPlatform, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


