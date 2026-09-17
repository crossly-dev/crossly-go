# CreateCbxWalletPaymentQuoteResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TransactionBase64** | **string** |  | 
**BaseUnits** | **string** |  | 
**ValueCents** | **float32** |  | 
**CentsPerToken** | **float32** |  | 
**SourceTokenAccount** | **string** |  | 
**DestinationTokenAccount** | **string** |  | 
**DestinationOwner** | **string** |  | 
**RecentBlockhash** | **string** |  | 
**LastValidBlockHeight** | **float32** |  | 

## Methods

### NewCreateCbxWalletPaymentQuoteResponse

`func NewCreateCbxWalletPaymentQuoteResponse(transactionBase64 string, baseUnits string, valueCents float32, centsPerToken float32, sourceTokenAccount string, destinationTokenAccount string, destinationOwner string, recentBlockhash string, lastValidBlockHeight float32, ) *CreateCbxWalletPaymentQuoteResponse`

NewCreateCbxWalletPaymentQuoteResponse instantiates a new CreateCbxWalletPaymentQuoteResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCbxWalletPaymentQuoteResponseWithDefaults

`func NewCreateCbxWalletPaymentQuoteResponseWithDefaults() *CreateCbxWalletPaymentQuoteResponse`

NewCreateCbxWalletPaymentQuoteResponseWithDefaults instantiates a new CreateCbxWalletPaymentQuoteResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTransactionBase64

`func (o *CreateCbxWalletPaymentQuoteResponse) GetTransactionBase64() string`

GetTransactionBase64 returns the TransactionBase64 field if non-nil, zero value otherwise.

### GetTransactionBase64Ok

`func (o *CreateCbxWalletPaymentQuoteResponse) GetTransactionBase64Ok() (*string, bool)`

GetTransactionBase64Ok returns a tuple with the TransactionBase64 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransactionBase64

`func (o *CreateCbxWalletPaymentQuoteResponse) SetTransactionBase64(v string)`

SetTransactionBase64 sets TransactionBase64 field to given value.


### GetBaseUnits

`func (o *CreateCbxWalletPaymentQuoteResponse) GetBaseUnits() string`

GetBaseUnits returns the BaseUnits field if non-nil, zero value otherwise.

### GetBaseUnitsOk

`func (o *CreateCbxWalletPaymentQuoteResponse) GetBaseUnitsOk() (*string, bool)`

GetBaseUnitsOk returns a tuple with the BaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBaseUnits

`func (o *CreateCbxWalletPaymentQuoteResponse) SetBaseUnits(v string)`

SetBaseUnits sets BaseUnits field to given value.


### GetValueCents

`func (o *CreateCbxWalletPaymentQuoteResponse) GetValueCents() float32`

GetValueCents returns the ValueCents field if non-nil, zero value otherwise.

### GetValueCentsOk

`func (o *CreateCbxWalletPaymentQuoteResponse) GetValueCentsOk() (*float32, bool)`

GetValueCentsOk returns a tuple with the ValueCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValueCents

`func (o *CreateCbxWalletPaymentQuoteResponse) SetValueCents(v float32)`

SetValueCents sets ValueCents field to given value.


### GetCentsPerToken

`func (o *CreateCbxWalletPaymentQuoteResponse) GetCentsPerToken() float32`

GetCentsPerToken returns the CentsPerToken field if non-nil, zero value otherwise.

### GetCentsPerTokenOk

`func (o *CreateCbxWalletPaymentQuoteResponse) GetCentsPerTokenOk() (*float32, bool)`

GetCentsPerTokenOk returns a tuple with the CentsPerToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCentsPerToken

`func (o *CreateCbxWalletPaymentQuoteResponse) SetCentsPerToken(v float32)`

SetCentsPerToken sets CentsPerToken field to given value.


### GetSourceTokenAccount

`func (o *CreateCbxWalletPaymentQuoteResponse) GetSourceTokenAccount() string`

GetSourceTokenAccount returns the SourceTokenAccount field if non-nil, zero value otherwise.

### GetSourceTokenAccountOk

`func (o *CreateCbxWalletPaymentQuoteResponse) GetSourceTokenAccountOk() (*string, bool)`

GetSourceTokenAccountOk returns a tuple with the SourceTokenAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceTokenAccount

`func (o *CreateCbxWalletPaymentQuoteResponse) SetSourceTokenAccount(v string)`

SetSourceTokenAccount sets SourceTokenAccount field to given value.


### GetDestinationTokenAccount

`func (o *CreateCbxWalletPaymentQuoteResponse) GetDestinationTokenAccount() string`

GetDestinationTokenAccount returns the DestinationTokenAccount field if non-nil, zero value otherwise.

### GetDestinationTokenAccountOk

`func (o *CreateCbxWalletPaymentQuoteResponse) GetDestinationTokenAccountOk() (*string, bool)`

GetDestinationTokenAccountOk returns a tuple with the DestinationTokenAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationTokenAccount

`func (o *CreateCbxWalletPaymentQuoteResponse) SetDestinationTokenAccount(v string)`

SetDestinationTokenAccount sets DestinationTokenAccount field to given value.


### GetDestinationOwner

`func (o *CreateCbxWalletPaymentQuoteResponse) GetDestinationOwner() string`

GetDestinationOwner returns the DestinationOwner field if non-nil, zero value otherwise.

### GetDestinationOwnerOk

`func (o *CreateCbxWalletPaymentQuoteResponse) GetDestinationOwnerOk() (*string, bool)`

GetDestinationOwnerOk returns a tuple with the DestinationOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationOwner

`func (o *CreateCbxWalletPaymentQuoteResponse) SetDestinationOwner(v string)`

SetDestinationOwner sets DestinationOwner field to given value.


### GetRecentBlockhash

`func (o *CreateCbxWalletPaymentQuoteResponse) GetRecentBlockhash() string`

GetRecentBlockhash returns the RecentBlockhash field if non-nil, zero value otherwise.

### GetRecentBlockhashOk

`func (o *CreateCbxWalletPaymentQuoteResponse) GetRecentBlockhashOk() (*string, bool)`

GetRecentBlockhashOk returns a tuple with the RecentBlockhash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecentBlockhash

`func (o *CreateCbxWalletPaymentQuoteResponse) SetRecentBlockhash(v string)`

SetRecentBlockhash sets RecentBlockhash field to given value.


### GetLastValidBlockHeight

`func (o *CreateCbxWalletPaymentQuoteResponse) GetLastValidBlockHeight() float32`

GetLastValidBlockHeight returns the LastValidBlockHeight field if non-nil, zero value otherwise.

### GetLastValidBlockHeightOk

`func (o *CreateCbxWalletPaymentQuoteResponse) GetLastValidBlockHeightOk() (*float32, bool)`

GetLastValidBlockHeightOk returns a tuple with the LastValidBlockHeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastValidBlockHeight

`func (o *CreateCbxWalletPaymentQuoteResponse) SetLastValidBlockHeight(v float32)`

SetLastValidBlockHeight sets LastValidBlockHeight field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


