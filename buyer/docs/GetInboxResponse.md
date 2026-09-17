# GetInboxResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Messages** | [**[]GetInboxResponseMessages**](GetInboxResponseMessages.md) |  | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Status** | **string** |  | 
**Platform** | **string** |  | 
**ListingId** | Pointer to **NullableString** |  | [optional] 
**AccountSlot** | **float32** |  | 
**InventoryItemId** | Pointer to **NullableString** |  | [optional] 
**BuyerUsername** | Pointer to **NullableString** |  | [optional] 
**Subject** | Pointer to **NullableString** |  | [optional] 
**PlatformConversationId** | **string** |  | 
**BuyerPlatformId** | Pointer to **NullableString** |  | [optional] 
**IsRead** | **bool** |  | 
**IsHidden** | **bool** |  | 
**HasOffer** | **bool** |  | 
**OfferAmount** | Pointer to **NullableString** |  | [optional] 
**OfferStatus** | Pointer to **NullableString** |  | [optional] 
**PlatformConversationUrl** | Pointer to **NullableString** |  | [optional] 
**LastMessageAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewGetInboxResponse

`func NewGetInboxResponse(messages []GetInboxResponseMessages, id string, createdAt time.Time, userId string, status string, platform string, accountSlot float32, platformConversationId string, isRead bool, isHidden bool, hasOffer bool, ) *GetInboxResponse`

NewGetInboxResponse instantiates a new GetInboxResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInboxResponseWithDefaults

`func NewGetInboxResponseWithDefaults() *GetInboxResponse`

NewGetInboxResponseWithDefaults instantiates a new GetInboxResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessages

`func (o *GetInboxResponse) GetMessages() []GetInboxResponseMessages`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *GetInboxResponse) GetMessagesOk() (*[]GetInboxResponseMessages, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *GetInboxResponse) SetMessages(v []GetInboxResponseMessages)`

SetMessages sets Messages field to given value.


### GetId

`func (o *GetInboxResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInboxResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInboxResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetInboxResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetInboxResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetInboxResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *GetInboxResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetInboxResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetInboxResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetStatus

`func (o *GetInboxResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetInboxResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetInboxResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlatform

`func (o *GetInboxResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetInboxResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetInboxResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetListingId

`func (o *GetInboxResponse) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *GetInboxResponse) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *GetInboxResponse) SetListingId(v string)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *GetInboxResponse) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### SetListingIdNil

`func (o *GetInboxResponse) SetListingIdNil(b bool)`

 SetListingIdNil sets the value for ListingId to be an explicit nil

### UnsetListingId
`func (o *GetInboxResponse) UnsetListingId()`

UnsetListingId ensures that no value is present for ListingId, not even an explicit nil
### GetAccountSlot

`func (o *GetInboxResponse) GetAccountSlot() float32`

GetAccountSlot returns the AccountSlot field if non-nil, zero value otherwise.

### GetAccountSlotOk

`func (o *GetInboxResponse) GetAccountSlotOk() (*float32, bool)`

GetAccountSlotOk returns a tuple with the AccountSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSlot

`func (o *GetInboxResponse) SetAccountSlot(v float32)`

SetAccountSlot sets AccountSlot field to given value.


### GetInventoryItemId

`func (o *GetInboxResponse) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *GetInboxResponse) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *GetInboxResponse) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *GetInboxResponse) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### SetInventoryItemIdNil

`func (o *GetInboxResponse) SetInventoryItemIdNil(b bool)`

 SetInventoryItemIdNil sets the value for InventoryItemId to be an explicit nil

### UnsetInventoryItemId
`func (o *GetInboxResponse) UnsetInventoryItemId()`

UnsetInventoryItemId ensures that no value is present for InventoryItemId, not even an explicit nil
### GetBuyerUsername

`func (o *GetInboxResponse) GetBuyerUsername() string`

GetBuyerUsername returns the BuyerUsername field if non-nil, zero value otherwise.

### GetBuyerUsernameOk

`func (o *GetInboxResponse) GetBuyerUsernameOk() (*string, bool)`

GetBuyerUsernameOk returns a tuple with the BuyerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerUsername

`func (o *GetInboxResponse) SetBuyerUsername(v string)`

SetBuyerUsername sets BuyerUsername field to given value.

### HasBuyerUsername

`func (o *GetInboxResponse) HasBuyerUsername() bool`

HasBuyerUsername returns a boolean if a field has been set.

### SetBuyerUsernameNil

`func (o *GetInboxResponse) SetBuyerUsernameNil(b bool)`

 SetBuyerUsernameNil sets the value for BuyerUsername to be an explicit nil

### UnsetBuyerUsername
`func (o *GetInboxResponse) UnsetBuyerUsername()`

UnsetBuyerUsername ensures that no value is present for BuyerUsername, not even an explicit nil
### GetSubject

`func (o *GetInboxResponse) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *GetInboxResponse) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *GetInboxResponse) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *GetInboxResponse) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *GetInboxResponse) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *GetInboxResponse) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetPlatformConversationId

`func (o *GetInboxResponse) GetPlatformConversationId() string`

GetPlatformConversationId returns the PlatformConversationId field if non-nil, zero value otherwise.

### GetPlatformConversationIdOk

`func (o *GetInboxResponse) GetPlatformConversationIdOk() (*string, bool)`

GetPlatformConversationIdOk returns a tuple with the PlatformConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformConversationId

`func (o *GetInboxResponse) SetPlatformConversationId(v string)`

SetPlatformConversationId sets PlatformConversationId field to given value.


### GetBuyerPlatformId

`func (o *GetInboxResponse) GetBuyerPlatformId() string`

GetBuyerPlatformId returns the BuyerPlatformId field if non-nil, zero value otherwise.

### GetBuyerPlatformIdOk

`func (o *GetInboxResponse) GetBuyerPlatformIdOk() (*string, bool)`

GetBuyerPlatformIdOk returns a tuple with the BuyerPlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerPlatformId

`func (o *GetInboxResponse) SetBuyerPlatformId(v string)`

SetBuyerPlatformId sets BuyerPlatformId field to given value.

### HasBuyerPlatformId

`func (o *GetInboxResponse) HasBuyerPlatformId() bool`

HasBuyerPlatformId returns a boolean if a field has been set.

### SetBuyerPlatformIdNil

`func (o *GetInboxResponse) SetBuyerPlatformIdNil(b bool)`

 SetBuyerPlatformIdNil sets the value for BuyerPlatformId to be an explicit nil

### UnsetBuyerPlatformId
`func (o *GetInboxResponse) UnsetBuyerPlatformId()`

UnsetBuyerPlatformId ensures that no value is present for BuyerPlatformId, not even an explicit nil
### GetIsRead

`func (o *GetInboxResponse) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *GetInboxResponse) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *GetInboxResponse) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.


### GetIsHidden

`func (o *GetInboxResponse) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *GetInboxResponse) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *GetInboxResponse) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.


### GetHasOffer

`func (o *GetInboxResponse) GetHasOffer() bool`

GetHasOffer returns the HasOffer field if non-nil, zero value otherwise.

### GetHasOfferOk

`func (o *GetInboxResponse) GetHasOfferOk() (*bool, bool)`

GetHasOfferOk returns a tuple with the HasOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOffer

`func (o *GetInboxResponse) SetHasOffer(v bool)`

SetHasOffer sets HasOffer field to given value.


### GetOfferAmount

`func (o *GetInboxResponse) GetOfferAmount() string`

GetOfferAmount returns the OfferAmount field if non-nil, zero value otherwise.

### GetOfferAmountOk

`func (o *GetInboxResponse) GetOfferAmountOk() (*string, bool)`

GetOfferAmountOk returns a tuple with the OfferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferAmount

`func (o *GetInboxResponse) SetOfferAmount(v string)`

SetOfferAmount sets OfferAmount field to given value.

### HasOfferAmount

`func (o *GetInboxResponse) HasOfferAmount() bool`

HasOfferAmount returns a boolean if a field has been set.

### SetOfferAmountNil

`func (o *GetInboxResponse) SetOfferAmountNil(b bool)`

 SetOfferAmountNil sets the value for OfferAmount to be an explicit nil

### UnsetOfferAmount
`func (o *GetInboxResponse) UnsetOfferAmount()`

UnsetOfferAmount ensures that no value is present for OfferAmount, not even an explicit nil
### GetOfferStatus

`func (o *GetInboxResponse) GetOfferStatus() string`

GetOfferStatus returns the OfferStatus field if non-nil, zero value otherwise.

### GetOfferStatusOk

`func (o *GetInboxResponse) GetOfferStatusOk() (*string, bool)`

GetOfferStatusOk returns a tuple with the OfferStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferStatus

`func (o *GetInboxResponse) SetOfferStatus(v string)`

SetOfferStatus sets OfferStatus field to given value.

### HasOfferStatus

`func (o *GetInboxResponse) HasOfferStatus() bool`

HasOfferStatus returns a boolean if a field has been set.

### SetOfferStatusNil

`func (o *GetInboxResponse) SetOfferStatusNil(b bool)`

 SetOfferStatusNil sets the value for OfferStatus to be an explicit nil

### UnsetOfferStatus
`func (o *GetInboxResponse) UnsetOfferStatus()`

UnsetOfferStatus ensures that no value is present for OfferStatus, not even an explicit nil
### GetPlatformConversationUrl

`func (o *GetInboxResponse) GetPlatformConversationUrl() string`

GetPlatformConversationUrl returns the PlatformConversationUrl field if non-nil, zero value otherwise.

### GetPlatformConversationUrlOk

`func (o *GetInboxResponse) GetPlatformConversationUrlOk() (*string, bool)`

GetPlatformConversationUrlOk returns a tuple with the PlatformConversationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformConversationUrl

`func (o *GetInboxResponse) SetPlatformConversationUrl(v string)`

SetPlatformConversationUrl sets PlatformConversationUrl field to given value.

### HasPlatformConversationUrl

`func (o *GetInboxResponse) HasPlatformConversationUrl() bool`

HasPlatformConversationUrl returns a boolean if a field has been set.

### SetPlatformConversationUrlNil

`func (o *GetInboxResponse) SetPlatformConversationUrlNil(b bool)`

 SetPlatformConversationUrlNil sets the value for PlatformConversationUrl to be an explicit nil

### UnsetPlatformConversationUrl
`func (o *GetInboxResponse) UnsetPlatformConversationUrl()`

UnsetPlatformConversationUrl ensures that no value is present for PlatformConversationUrl, not even an explicit nil
### GetLastMessageAt

`func (o *GetInboxResponse) GetLastMessageAt() time.Time`

GetLastMessageAt returns the LastMessageAt field if non-nil, zero value otherwise.

### GetLastMessageAtOk

`func (o *GetInboxResponse) GetLastMessageAtOk() (*time.Time, bool)`

GetLastMessageAtOk returns a tuple with the LastMessageAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageAt

`func (o *GetInboxResponse) SetLastMessageAt(v time.Time)`

SetLastMessageAt sets LastMessageAt field to given value.

### HasLastMessageAt

`func (o *GetInboxResponse) HasLastMessageAt() bool`

HasLastMessageAt returns a boolean if a field has been set.

### SetLastMessageAtNil

`func (o *GetInboxResponse) SetLastMessageAtNil(b bool)`

 SetLastMessageAtNil sets the value for LastMessageAt to be an explicit nil

### UnsetLastMessageAt
`func (o *GetInboxResponse) UnsetLastMessageAt()`

UnsetLastMessageAt ensures that no value is present for LastMessageAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


