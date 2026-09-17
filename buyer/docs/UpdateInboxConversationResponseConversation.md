# UpdateInboxConversationResponseConversation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewUpdateInboxConversationResponseConversation

`func NewUpdateInboxConversationResponseConversation(id string, createdAt time.Time, userId string, status string, platform string, accountSlot float32, platformConversationId string, isRead bool, isHidden bool, hasOffer bool, ) *UpdateInboxConversationResponseConversation`

NewUpdateInboxConversationResponseConversation instantiates a new UpdateInboxConversationResponseConversation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInboxConversationResponseConversationWithDefaults

`func NewUpdateInboxConversationResponseConversationWithDefaults() *UpdateInboxConversationResponseConversation`

NewUpdateInboxConversationResponseConversationWithDefaults instantiates a new UpdateInboxConversationResponseConversation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateInboxConversationResponseConversation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateInboxConversationResponseConversation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateInboxConversationResponseConversation) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *UpdateInboxConversationResponseConversation) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateInboxConversationResponseConversation) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateInboxConversationResponseConversation) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *UpdateInboxConversationResponseConversation) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateInboxConversationResponseConversation) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateInboxConversationResponseConversation) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetStatus

`func (o *UpdateInboxConversationResponseConversation) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateInboxConversationResponseConversation) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateInboxConversationResponseConversation) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlatform

`func (o *UpdateInboxConversationResponseConversation) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdateInboxConversationResponseConversation) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdateInboxConversationResponseConversation) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetListingId

`func (o *UpdateInboxConversationResponseConversation) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *UpdateInboxConversationResponseConversation) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *UpdateInboxConversationResponseConversation) SetListingId(v string)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *UpdateInboxConversationResponseConversation) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### SetListingIdNil

`func (o *UpdateInboxConversationResponseConversation) SetListingIdNil(b bool)`

 SetListingIdNil sets the value for ListingId to be an explicit nil

### UnsetListingId
`func (o *UpdateInboxConversationResponseConversation) UnsetListingId()`

UnsetListingId ensures that no value is present for ListingId, not even an explicit nil
### GetAccountSlot

`func (o *UpdateInboxConversationResponseConversation) GetAccountSlot() float32`

GetAccountSlot returns the AccountSlot field if non-nil, zero value otherwise.

### GetAccountSlotOk

`func (o *UpdateInboxConversationResponseConversation) GetAccountSlotOk() (*float32, bool)`

GetAccountSlotOk returns a tuple with the AccountSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSlot

`func (o *UpdateInboxConversationResponseConversation) SetAccountSlot(v float32)`

SetAccountSlot sets AccountSlot field to given value.


### GetInventoryItemId

`func (o *UpdateInboxConversationResponseConversation) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *UpdateInboxConversationResponseConversation) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *UpdateInboxConversationResponseConversation) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *UpdateInboxConversationResponseConversation) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### SetInventoryItemIdNil

`func (o *UpdateInboxConversationResponseConversation) SetInventoryItemIdNil(b bool)`

 SetInventoryItemIdNil sets the value for InventoryItemId to be an explicit nil

### UnsetInventoryItemId
`func (o *UpdateInboxConversationResponseConversation) UnsetInventoryItemId()`

UnsetInventoryItemId ensures that no value is present for InventoryItemId, not even an explicit nil
### GetBuyerUsername

`func (o *UpdateInboxConversationResponseConversation) GetBuyerUsername() string`

GetBuyerUsername returns the BuyerUsername field if non-nil, zero value otherwise.

### GetBuyerUsernameOk

`func (o *UpdateInboxConversationResponseConversation) GetBuyerUsernameOk() (*string, bool)`

GetBuyerUsernameOk returns a tuple with the BuyerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerUsername

`func (o *UpdateInboxConversationResponseConversation) SetBuyerUsername(v string)`

SetBuyerUsername sets BuyerUsername field to given value.

### HasBuyerUsername

`func (o *UpdateInboxConversationResponseConversation) HasBuyerUsername() bool`

HasBuyerUsername returns a boolean if a field has been set.

### SetBuyerUsernameNil

`func (o *UpdateInboxConversationResponseConversation) SetBuyerUsernameNil(b bool)`

 SetBuyerUsernameNil sets the value for BuyerUsername to be an explicit nil

### UnsetBuyerUsername
`func (o *UpdateInboxConversationResponseConversation) UnsetBuyerUsername()`

UnsetBuyerUsername ensures that no value is present for BuyerUsername, not even an explicit nil
### GetSubject

`func (o *UpdateInboxConversationResponseConversation) GetSubject() string`

GetSubject returns the Subject field if non-nil, zero value otherwise.

### GetSubjectOk

`func (o *UpdateInboxConversationResponseConversation) GetSubjectOk() (*string, bool)`

GetSubjectOk returns a tuple with the Subject field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubject

`func (o *UpdateInboxConversationResponseConversation) SetSubject(v string)`

SetSubject sets Subject field to given value.

### HasSubject

`func (o *UpdateInboxConversationResponseConversation) HasSubject() bool`

HasSubject returns a boolean if a field has been set.

### SetSubjectNil

`func (o *UpdateInboxConversationResponseConversation) SetSubjectNil(b bool)`

 SetSubjectNil sets the value for Subject to be an explicit nil

### UnsetSubject
`func (o *UpdateInboxConversationResponseConversation) UnsetSubject()`

UnsetSubject ensures that no value is present for Subject, not even an explicit nil
### GetPlatformConversationId

`func (o *UpdateInboxConversationResponseConversation) GetPlatformConversationId() string`

GetPlatformConversationId returns the PlatformConversationId field if non-nil, zero value otherwise.

### GetPlatformConversationIdOk

`func (o *UpdateInboxConversationResponseConversation) GetPlatformConversationIdOk() (*string, bool)`

GetPlatformConversationIdOk returns a tuple with the PlatformConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformConversationId

`func (o *UpdateInboxConversationResponseConversation) SetPlatformConversationId(v string)`

SetPlatformConversationId sets PlatformConversationId field to given value.


### GetBuyerPlatformId

`func (o *UpdateInboxConversationResponseConversation) GetBuyerPlatformId() string`

GetBuyerPlatformId returns the BuyerPlatformId field if non-nil, zero value otherwise.

### GetBuyerPlatformIdOk

`func (o *UpdateInboxConversationResponseConversation) GetBuyerPlatformIdOk() (*string, bool)`

GetBuyerPlatformIdOk returns a tuple with the BuyerPlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerPlatformId

`func (o *UpdateInboxConversationResponseConversation) SetBuyerPlatformId(v string)`

SetBuyerPlatformId sets BuyerPlatformId field to given value.

### HasBuyerPlatformId

`func (o *UpdateInboxConversationResponseConversation) HasBuyerPlatformId() bool`

HasBuyerPlatformId returns a boolean if a field has been set.

### SetBuyerPlatformIdNil

`func (o *UpdateInboxConversationResponseConversation) SetBuyerPlatformIdNil(b bool)`

 SetBuyerPlatformIdNil sets the value for BuyerPlatformId to be an explicit nil

### UnsetBuyerPlatformId
`func (o *UpdateInboxConversationResponseConversation) UnsetBuyerPlatformId()`

UnsetBuyerPlatformId ensures that no value is present for BuyerPlatformId, not even an explicit nil
### GetIsRead

`func (o *UpdateInboxConversationResponseConversation) GetIsRead() bool`

GetIsRead returns the IsRead field if non-nil, zero value otherwise.

### GetIsReadOk

`func (o *UpdateInboxConversationResponseConversation) GetIsReadOk() (*bool, bool)`

GetIsReadOk returns a tuple with the IsRead field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsRead

`func (o *UpdateInboxConversationResponseConversation) SetIsRead(v bool)`

SetIsRead sets IsRead field to given value.


### GetIsHidden

`func (o *UpdateInboxConversationResponseConversation) GetIsHidden() bool`

GetIsHidden returns the IsHidden field if non-nil, zero value otherwise.

### GetIsHiddenOk

`func (o *UpdateInboxConversationResponseConversation) GetIsHiddenOk() (*bool, bool)`

GetIsHiddenOk returns a tuple with the IsHidden field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsHidden

`func (o *UpdateInboxConversationResponseConversation) SetIsHidden(v bool)`

SetIsHidden sets IsHidden field to given value.


### GetHasOffer

`func (o *UpdateInboxConversationResponseConversation) GetHasOffer() bool`

GetHasOffer returns the HasOffer field if non-nil, zero value otherwise.

### GetHasOfferOk

`func (o *UpdateInboxConversationResponseConversation) GetHasOfferOk() (*bool, bool)`

GetHasOfferOk returns a tuple with the HasOffer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasOffer

`func (o *UpdateInboxConversationResponseConversation) SetHasOffer(v bool)`

SetHasOffer sets HasOffer field to given value.


### GetOfferAmount

`func (o *UpdateInboxConversationResponseConversation) GetOfferAmount() string`

GetOfferAmount returns the OfferAmount field if non-nil, zero value otherwise.

### GetOfferAmountOk

`func (o *UpdateInboxConversationResponseConversation) GetOfferAmountOk() (*string, bool)`

GetOfferAmountOk returns a tuple with the OfferAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferAmount

`func (o *UpdateInboxConversationResponseConversation) SetOfferAmount(v string)`

SetOfferAmount sets OfferAmount field to given value.

### HasOfferAmount

`func (o *UpdateInboxConversationResponseConversation) HasOfferAmount() bool`

HasOfferAmount returns a boolean if a field has been set.

### SetOfferAmountNil

`func (o *UpdateInboxConversationResponseConversation) SetOfferAmountNil(b bool)`

 SetOfferAmountNil sets the value for OfferAmount to be an explicit nil

### UnsetOfferAmount
`func (o *UpdateInboxConversationResponseConversation) UnsetOfferAmount()`

UnsetOfferAmount ensures that no value is present for OfferAmount, not even an explicit nil
### GetOfferStatus

`func (o *UpdateInboxConversationResponseConversation) GetOfferStatus() string`

GetOfferStatus returns the OfferStatus field if non-nil, zero value otherwise.

### GetOfferStatusOk

`func (o *UpdateInboxConversationResponseConversation) GetOfferStatusOk() (*string, bool)`

GetOfferStatusOk returns a tuple with the OfferStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferStatus

`func (o *UpdateInboxConversationResponseConversation) SetOfferStatus(v string)`

SetOfferStatus sets OfferStatus field to given value.

### HasOfferStatus

`func (o *UpdateInboxConversationResponseConversation) HasOfferStatus() bool`

HasOfferStatus returns a boolean if a field has been set.

### SetOfferStatusNil

`func (o *UpdateInboxConversationResponseConversation) SetOfferStatusNil(b bool)`

 SetOfferStatusNil sets the value for OfferStatus to be an explicit nil

### UnsetOfferStatus
`func (o *UpdateInboxConversationResponseConversation) UnsetOfferStatus()`

UnsetOfferStatus ensures that no value is present for OfferStatus, not even an explicit nil
### GetPlatformConversationUrl

`func (o *UpdateInboxConversationResponseConversation) GetPlatformConversationUrl() string`

GetPlatformConversationUrl returns the PlatformConversationUrl field if non-nil, zero value otherwise.

### GetPlatformConversationUrlOk

`func (o *UpdateInboxConversationResponseConversation) GetPlatformConversationUrlOk() (*string, bool)`

GetPlatformConversationUrlOk returns a tuple with the PlatformConversationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformConversationUrl

`func (o *UpdateInboxConversationResponseConversation) SetPlatformConversationUrl(v string)`

SetPlatformConversationUrl sets PlatformConversationUrl field to given value.

### HasPlatformConversationUrl

`func (o *UpdateInboxConversationResponseConversation) HasPlatformConversationUrl() bool`

HasPlatformConversationUrl returns a boolean if a field has been set.

### SetPlatformConversationUrlNil

`func (o *UpdateInboxConversationResponseConversation) SetPlatformConversationUrlNil(b bool)`

 SetPlatformConversationUrlNil sets the value for PlatformConversationUrl to be an explicit nil

### UnsetPlatformConversationUrl
`func (o *UpdateInboxConversationResponseConversation) UnsetPlatformConversationUrl()`

UnsetPlatformConversationUrl ensures that no value is present for PlatformConversationUrl, not even an explicit nil
### GetLastMessageAt

`func (o *UpdateInboxConversationResponseConversation) GetLastMessageAt() time.Time`

GetLastMessageAt returns the LastMessageAt field if non-nil, zero value otherwise.

### GetLastMessageAtOk

`func (o *UpdateInboxConversationResponseConversation) GetLastMessageAtOk() (*time.Time, bool)`

GetLastMessageAtOk returns a tuple with the LastMessageAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastMessageAt

`func (o *UpdateInboxConversationResponseConversation) SetLastMessageAt(v time.Time)`

SetLastMessageAt sets LastMessageAt field to given value.

### HasLastMessageAt

`func (o *UpdateInboxConversationResponseConversation) HasLastMessageAt() bool`

HasLastMessageAt returns a boolean if a field has been set.

### SetLastMessageAtNil

`func (o *UpdateInboxConversationResponseConversation) SetLastMessageAtNil(b bool)`

 SetLastMessageAtNil sets the value for LastMessageAt to be an explicit nil

### UnsetLastMessageAt
`func (o *UpdateInboxConversationResponseConversation) UnsetLastMessageAt()`

UnsetLastMessageAt ensures that no value is present for LastMessageAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


