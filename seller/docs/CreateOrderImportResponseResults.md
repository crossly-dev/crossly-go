# CreateOrderImportResponseResults

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Platform** | **string** |  | 
**Ok** | **bool** |  | 
**AccountsStarted** | Pointer to **NullableFloat32** | Cookie track: number of account slots kicked off. API track: always 1. | [optional] 

## Methods

### NewCreateOrderImportResponseResults

`func NewCreateOrderImportResponseResults(platform string, ok bool, ) *CreateOrderImportResponseResults`

NewCreateOrderImportResponseResults instantiates a new CreateOrderImportResponseResults object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateOrderImportResponseResultsWithDefaults

`func NewCreateOrderImportResponseResultsWithDefaults() *CreateOrderImportResponseResults`

NewCreateOrderImportResponseResultsWithDefaults instantiates a new CreateOrderImportResponseResults object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatform

`func (o *CreateOrderImportResponseResults) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *CreateOrderImportResponseResults) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *CreateOrderImportResponseResults) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetOk

`func (o *CreateOrderImportResponseResults) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *CreateOrderImportResponseResults) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *CreateOrderImportResponseResults) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetAccountsStarted

`func (o *CreateOrderImportResponseResults) GetAccountsStarted() float32`

GetAccountsStarted returns the AccountsStarted field if non-nil, zero value otherwise.

### GetAccountsStartedOk

`func (o *CreateOrderImportResponseResults) GetAccountsStartedOk() (*float32, bool)`

GetAccountsStartedOk returns a tuple with the AccountsStarted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountsStarted

`func (o *CreateOrderImportResponseResults) SetAccountsStarted(v float32)`

SetAccountsStarted sets AccountsStarted field to given value.

### HasAccountsStarted

`func (o *CreateOrderImportResponseResults) HasAccountsStarted() bool`

HasAccountsStarted returns a boolean if a field has been set.

### SetAccountsStartedNil

`func (o *CreateOrderImportResponseResults) SetAccountsStartedNil(b bool)`

 SetAccountsStartedNil sets the value for AccountsStarted to be an explicit nil

### UnsetAccountsStarted
`func (o *CreateOrderImportResponseResults) UnsetAccountsStarted()`

UnsetAccountsStarted ensures that no value is present for AccountsStarted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


