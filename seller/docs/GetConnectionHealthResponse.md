# GetConnectionHealthResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Extension** | [**GetConnectionHealthResponseExtension**](GetConnectionHealthResponseExtension.md) |  | 
**Accounts** | [**[]GetConnectionHealthResponseAccounts**](GetConnectionHealthResponseAccounts.md) |  | 
**GeneratedAt** | **string** |  | 

## Methods

### NewGetConnectionHealthResponse

`func NewGetConnectionHealthResponse(extension GetConnectionHealthResponseExtension, accounts []GetConnectionHealthResponseAccounts, generatedAt string, ) *GetConnectionHealthResponse`

NewGetConnectionHealthResponse instantiates a new GetConnectionHealthResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionHealthResponseWithDefaults

`func NewGetConnectionHealthResponseWithDefaults() *GetConnectionHealthResponse`

NewGetConnectionHealthResponseWithDefaults instantiates a new GetConnectionHealthResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExtension

`func (o *GetConnectionHealthResponse) GetExtension() GetConnectionHealthResponseExtension`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *GetConnectionHealthResponse) GetExtensionOk() (*GetConnectionHealthResponseExtension, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *GetConnectionHealthResponse) SetExtension(v GetConnectionHealthResponseExtension)`

SetExtension sets Extension field to given value.


### GetAccounts

`func (o *GetConnectionHealthResponse) GetAccounts() []GetConnectionHealthResponseAccounts`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetConnectionHealthResponse) GetAccountsOk() (*[]GetConnectionHealthResponseAccounts, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetConnectionHealthResponse) SetAccounts(v []GetConnectionHealthResponseAccounts)`

SetAccounts sets Accounts field to given value.


### GetGeneratedAt

`func (o *GetConnectionHealthResponse) GetGeneratedAt() string`

GetGeneratedAt returns the GeneratedAt field if non-nil, zero value otherwise.

### GetGeneratedAtOk

`func (o *GetConnectionHealthResponse) GetGeneratedAtOk() (*string, bool)`

GetGeneratedAtOk returns a tuple with the GeneratedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeneratedAt

`func (o *GetConnectionHealthResponse) SetGeneratedAt(v string)`

SetGeneratedAt sets GeneratedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


