# GetOfferResponseOffers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Status** | **string** |  | 
**AmountCents** | **float32** |  | 
**ListedTotalCents** | Pointer to **NullableFloat32** | Asking total when the offer was made — what the buyer responded to. | [optional] 
**IsBundle** | **bool** |  | 
**Items** | **[]map[string]interface{}** |  | 
**Message** | Pointer to **NullableString** |  | [optional] 
**ParentOfferId** | Pointer to **NullableString** |  | [optional] 
**ExpiresAt** | **time.Time** |  | 
**DecidedAt** | Pointer to **NullableTime** |  | [optional] 
**ConsumedAt** | Pointer to **NullableTime** | Set once an accepted offer has actually been paid for. | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewGetOfferResponseOffers

`func NewGetOfferResponseOffers(id string, status string, amountCents float32, isBundle bool, items []map[string]interface{}, expiresAt time.Time, createdAt time.Time, ) *GetOfferResponseOffers`

NewGetOfferResponseOffers instantiates a new GetOfferResponseOffers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOfferResponseOffersWithDefaults

`func NewGetOfferResponseOffersWithDefaults() *GetOfferResponseOffers`

NewGetOfferResponseOffersWithDefaults instantiates a new GetOfferResponseOffers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOfferResponseOffers) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOfferResponseOffers) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOfferResponseOffers) SetId(v string)`

SetId sets Id field to given value.


### GetStatus

`func (o *GetOfferResponseOffers) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOfferResponseOffers) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOfferResponseOffers) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAmountCents

`func (o *GetOfferResponseOffers) GetAmountCents() float32`

GetAmountCents returns the AmountCents field if non-nil, zero value otherwise.

### GetAmountCentsOk

`func (o *GetOfferResponseOffers) GetAmountCentsOk() (*float32, bool)`

GetAmountCentsOk returns a tuple with the AmountCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmountCents

`func (o *GetOfferResponseOffers) SetAmountCents(v float32)`

SetAmountCents sets AmountCents field to given value.


### GetListedTotalCents

`func (o *GetOfferResponseOffers) GetListedTotalCents() float32`

GetListedTotalCents returns the ListedTotalCents field if non-nil, zero value otherwise.

### GetListedTotalCentsOk

`func (o *GetOfferResponseOffers) GetListedTotalCentsOk() (*float32, bool)`

GetListedTotalCentsOk returns a tuple with the ListedTotalCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedTotalCents

`func (o *GetOfferResponseOffers) SetListedTotalCents(v float32)`

SetListedTotalCents sets ListedTotalCents field to given value.

### HasListedTotalCents

`func (o *GetOfferResponseOffers) HasListedTotalCents() bool`

HasListedTotalCents returns a boolean if a field has been set.

### SetListedTotalCentsNil

`func (o *GetOfferResponseOffers) SetListedTotalCentsNil(b bool)`

 SetListedTotalCentsNil sets the value for ListedTotalCents to be an explicit nil

### UnsetListedTotalCents
`func (o *GetOfferResponseOffers) UnsetListedTotalCents()`

UnsetListedTotalCents ensures that no value is present for ListedTotalCents, not even an explicit nil
### GetIsBundle

`func (o *GetOfferResponseOffers) GetIsBundle() bool`

GetIsBundle returns the IsBundle field if non-nil, zero value otherwise.

### GetIsBundleOk

`func (o *GetOfferResponseOffers) GetIsBundleOk() (*bool, bool)`

GetIsBundleOk returns a tuple with the IsBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBundle

`func (o *GetOfferResponseOffers) SetIsBundle(v bool)`

SetIsBundle sets IsBundle field to given value.


### GetItems

`func (o *GetOfferResponseOffers) GetItems() []map[string]interface{}`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetOfferResponseOffers) GetItemsOk() (*[]map[string]interface{}, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetOfferResponseOffers) SetItems(v []map[string]interface{})`

SetItems sets Items field to given value.


### GetMessage

`func (o *GetOfferResponseOffers) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *GetOfferResponseOffers) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *GetOfferResponseOffers) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *GetOfferResponseOffers) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### SetMessageNil

`func (o *GetOfferResponseOffers) SetMessageNil(b bool)`

 SetMessageNil sets the value for Message to be an explicit nil

### UnsetMessage
`func (o *GetOfferResponseOffers) UnsetMessage()`

UnsetMessage ensures that no value is present for Message, not even an explicit nil
### GetParentOfferId

`func (o *GetOfferResponseOffers) GetParentOfferId() string`

GetParentOfferId returns the ParentOfferId field if non-nil, zero value otherwise.

### GetParentOfferIdOk

`func (o *GetOfferResponseOffers) GetParentOfferIdOk() (*string, bool)`

GetParentOfferIdOk returns a tuple with the ParentOfferId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentOfferId

`func (o *GetOfferResponseOffers) SetParentOfferId(v string)`

SetParentOfferId sets ParentOfferId field to given value.

### HasParentOfferId

`func (o *GetOfferResponseOffers) HasParentOfferId() bool`

HasParentOfferId returns a boolean if a field has been set.

### SetParentOfferIdNil

`func (o *GetOfferResponseOffers) SetParentOfferIdNil(b bool)`

 SetParentOfferIdNil sets the value for ParentOfferId to be an explicit nil

### UnsetParentOfferId
`func (o *GetOfferResponseOffers) UnsetParentOfferId()`

UnsetParentOfferId ensures that no value is present for ParentOfferId, not even an explicit nil
### GetExpiresAt

`func (o *GetOfferResponseOffers) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *GetOfferResponseOffers) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *GetOfferResponseOffers) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.


### GetDecidedAt

`func (o *GetOfferResponseOffers) GetDecidedAt() time.Time`

GetDecidedAt returns the DecidedAt field if non-nil, zero value otherwise.

### GetDecidedAtOk

`func (o *GetOfferResponseOffers) GetDecidedAtOk() (*time.Time, bool)`

GetDecidedAtOk returns a tuple with the DecidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecidedAt

`func (o *GetOfferResponseOffers) SetDecidedAt(v time.Time)`

SetDecidedAt sets DecidedAt field to given value.

### HasDecidedAt

`func (o *GetOfferResponseOffers) HasDecidedAt() bool`

HasDecidedAt returns a boolean if a field has been set.

### SetDecidedAtNil

`func (o *GetOfferResponseOffers) SetDecidedAtNil(b bool)`

 SetDecidedAtNil sets the value for DecidedAt to be an explicit nil

### UnsetDecidedAt
`func (o *GetOfferResponseOffers) UnsetDecidedAt()`

UnsetDecidedAt ensures that no value is present for DecidedAt, not even an explicit nil
### GetConsumedAt

`func (o *GetOfferResponseOffers) GetConsumedAt() time.Time`

GetConsumedAt returns the ConsumedAt field if non-nil, zero value otherwise.

### GetConsumedAtOk

`func (o *GetOfferResponseOffers) GetConsumedAtOk() (*time.Time, bool)`

GetConsumedAtOk returns a tuple with the ConsumedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumedAt

`func (o *GetOfferResponseOffers) SetConsumedAt(v time.Time)`

SetConsumedAt sets ConsumedAt field to given value.

### HasConsumedAt

`func (o *GetOfferResponseOffers) HasConsumedAt() bool`

HasConsumedAt returns a boolean if a field has been set.

### SetConsumedAtNil

`func (o *GetOfferResponseOffers) SetConsumedAtNil(b bool)`

 SetConsumedAtNil sets the value for ConsumedAt to be an explicit nil

### UnsetConsumedAt
`func (o *GetOfferResponseOffers) UnsetConsumedAt()`

UnsetConsumedAt ensures that no value is present for ConsumedAt, not even an explicit nil
### GetCreatedAt

`func (o *GetOfferResponseOffers) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOfferResponseOffers) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOfferResponseOffers) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


