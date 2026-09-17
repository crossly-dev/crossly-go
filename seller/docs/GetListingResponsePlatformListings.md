# GetListingResponsePlatformListings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Quantity** | **float32** |  | 
**Status** | **string** |  | 
**Platform** | **string** |  | 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**ListingId** | **string** |  | 
**AccountSlot** | **float32** |  | 
**DelistedAt** | Pointer to **NullableTime** |  | [optional] 
**SoldAt** | Pointer to **NullableTime** |  | [optional] 
**InventoryItemId** | Pointer to **NullableString** |  | [optional] 
**PlatformListingId** | **string** |  | 
**PlatformListingUrl** | Pointer to **NullableString** |  | [optional] 
**StockLocation** | **string** |  | 
**ListedAt** | Pointer to **NullableTime** |  | [optional] 
**SharedAt** | Pointer to **NullableTime** |  | [optional] 
**DelistReason** | Pointer to **NullableString** |  | [optional] 
**LastCheckedAt** | Pointer to **NullableTime** |  | [optional] 
**Views** | **float32** |  | 
**Favorites** | **float32** |  | 
**CommentsCount** | **float32** |  | 
**LastStatsSyncedAt** | Pointer to **NullableTime** |  | [optional] 
**Hashtags** | Pointer to **[]string** |  | [optional] 
**IsDraft** | **bool** |  | 
**ScheduledPublishAt** | Pointer to **NullableTime** |  | [optional] 
**IsDutchAuction** | **bool** |  | 
**DutchStartCents** | Pointer to **NullableFloat32** |  | [optional] 
**DutchFloorCents** | Pointer to **NullableFloat32** |  | [optional] 
**DutchEndsAt** | Pointer to **NullableTime** |  | [optional] 
**IsPreOrder** | **bool** |  | 
**ExpectedShipBy** | Pointer to **NullableString** |  | [optional] 
**PreorderDepositPercentBps** | Pointer to **NullableFloat32** |  | [optional] 
**BoostedAt** | Pointer to **NullableTime** |  | [optional] 
**ImageCaption** | Pointer to **NullableString** |  | [optional] 
**DurationDays** | Pointer to **NullableFloat32** |  | [optional] 
**EndsAt** | Pointer to **NullableTime** |  | [optional] 
**BestOfferEnabled** | **bool** |  | 
**BestOfferAutoAcceptCents** | Pointer to **NullableFloat32** |  | [optional] 
**BestOfferAutoDeclineCents** | Pointer to **NullableFloat32** |  | [optional] 
**AestheticTags** | Pointer to **[]string** |  | [optional] 
**AllowsLocalPickup** | **bool** |  | 
**PickupRadiusMiles** | Pointer to **NullableFloat32** |  | [optional] 
**PickupLat** | Pointer to **NullableString** |  | [optional] 
**PickupLng** | Pointer to **NullableString** |  | [optional] 
**RefurbishedTier** | Pointer to **NullableString** |  | [optional] 
**CharitySlug** | Pointer to **NullableString** |  | [optional] 
**CharityPercentBps** | Pointer to **NullableFloat32** |  | [optional] 
**CrosslyReturnPolicyId** | Pointer to **NullableString** |  | [optional] 
**CrosslyShippingPolicyId** | Pointer to **NullableString** |  | [optional] 
**CrosslyPaymentPolicyId** | Pointer to **NullableString** |  | [optional] 
**RetailPriceCents** | Pointer to **NullableFloat32** |  | [optional] 
**RestrictedCountries** | Pointer to **[]string** |  | [optional] 
**AuctionExtendSeconds** | Pointer to **NullableFloat32** |  | [optional] 
**AuthenticationProviderSlug** | Pointer to **NullableString** |  | [optional] 
**AuthenticationRequired** | **bool** |  | 
**HasVariants** | **bool** |  | 
**CharityId** | Pointer to **NullableString** |  | [optional] 
**MinOrderQuantity** | **float32** |  | 
**WholesaleUnitPriceCents** | Pointer to **NullableFloat32** |  | [optional] 
**WholesaleMinQuantity** | Pointer to **NullableFloat32** |  | [optional] 
**AuctionReserveCents** | Pointer to **NullableFloat32** |  | [optional] 
**IsGtc** | **bool** |  | 
**HandlingTimeDays** | **float32** |  | 
**ItemLocation** | Pointer to **NullableString** |  | [optional] 
**ItemLocationZip** | Pointer to **NullableString** |  | [optional] 
**ItemLocationCountry** | Pointer to **NullableString** |  | [optional] 
**DeclaredShippingCost** | Pointer to **NullableString** |  | [optional] 
**ReturnPolicyText** | Pointer to **NullableString** |  | [optional] 
**LastDriftCheckedAt** | Pointer to **NullableTime** |  | [optional] 
**VerificationFailed** | **bool** |  | 
**PublishedOverrides** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetListingResponsePlatformListings

`func NewGetListingResponsePlatformListings(id string, userId string, quantity float32, status string, platform string, listingId string, accountSlot float32, platformListingId string, stockLocation string, views float32, favorites float32, commentsCount float32, isDraft bool, isDutchAuction bool, isPreOrder bool, bestOfferEnabled bool, allowsLocalPickup bool, authenticationRequired bool, hasVariants bool, minOrderQuantity float32, isGtc bool, handlingTimeDays float32, verificationFailed bool, ) *GetListingResponsePlatformListings`

NewGetListingResponsePlatformListings instantiates a new GetListingResponsePlatformListings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListingResponsePlatformListingsWithDefaults

`func NewGetListingResponsePlatformListingsWithDefaults() *GetListingResponsePlatformListings`

NewGetListingResponsePlatformListingsWithDefaults instantiates a new GetListingResponsePlatformListings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetListingResponsePlatformListings) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetListingResponsePlatformListings) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetListingResponsePlatformListings) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *GetListingResponsePlatformListings) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetListingResponsePlatformListings) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetListingResponsePlatformListings) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetQuantity

`func (o *GetListingResponsePlatformListings) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GetListingResponsePlatformListings) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GetListingResponsePlatformListings) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetStatus

`func (o *GetListingResponsePlatformListings) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetListingResponsePlatformListings) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetListingResponsePlatformListings) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlatform

`func (o *GetListingResponsePlatformListings) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetListingResponsePlatformListings) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetListingResponsePlatformListings) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetErrorMessage

`func (o *GetListingResponsePlatformListings) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *GetListingResponsePlatformListings) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *GetListingResponsePlatformListings) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *GetListingResponsePlatformListings) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *GetListingResponsePlatformListings) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *GetListingResponsePlatformListings) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetListingId

`func (o *GetListingResponsePlatformListings) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *GetListingResponsePlatformListings) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *GetListingResponsePlatformListings) SetListingId(v string)`

SetListingId sets ListingId field to given value.


### GetAccountSlot

`func (o *GetListingResponsePlatformListings) GetAccountSlot() float32`

GetAccountSlot returns the AccountSlot field if non-nil, zero value otherwise.

### GetAccountSlotOk

`func (o *GetListingResponsePlatformListings) GetAccountSlotOk() (*float32, bool)`

GetAccountSlotOk returns a tuple with the AccountSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSlot

`func (o *GetListingResponsePlatformListings) SetAccountSlot(v float32)`

SetAccountSlot sets AccountSlot field to given value.


### GetDelistedAt

`func (o *GetListingResponsePlatformListings) GetDelistedAt() time.Time`

GetDelistedAt returns the DelistedAt field if non-nil, zero value otherwise.

### GetDelistedAtOk

`func (o *GetListingResponsePlatformListings) GetDelistedAtOk() (*time.Time, bool)`

GetDelistedAtOk returns a tuple with the DelistedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelistedAt

`func (o *GetListingResponsePlatformListings) SetDelistedAt(v time.Time)`

SetDelistedAt sets DelistedAt field to given value.

### HasDelistedAt

`func (o *GetListingResponsePlatformListings) HasDelistedAt() bool`

HasDelistedAt returns a boolean if a field has been set.

### SetDelistedAtNil

`func (o *GetListingResponsePlatformListings) SetDelistedAtNil(b bool)`

 SetDelistedAtNil sets the value for DelistedAt to be an explicit nil

### UnsetDelistedAt
`func (o *GetListingResponsePlatformListings) UnsetDelistedAt()`

UnsetDelistedAt ensures that no value is present for DelistedAt, not even an explicit nil
### GetSoldAt

`func (o *GetListingResponsePlatformListings) GetSoldAt() time.Time`

GetSoldAt returns the SoldAt field if non-nil, zero value otherwise.

### GetSoldAtOk

`func (o *GetListingResponsePlatformListings) GetSoldAtOk() (*time.Time, bool)`

GetSoldAtOk returns a tuple with the SoldAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldAt

`func (o *GetListingResponsePlatformListings) SetSoldAt(v time.Time)`

SetSoldAt sets SoldAt field to given value.

### HasSoldAt

`func (o *GetListingResponsePlatformListings) HasSoldAt() bool`

HasSoldAt returns a boolean if a field has been set.

### SetSoldAtNil

`func (o *GetListingResponsePlatformListings) SetSoldAtNil(b bool)`

 SetSoldAtNil sets the value for SoldAt to be an explicit nil

### UnsetSoldAt
`func (o *GetListingResponsePlatformListings) UnsetSoldAt()`

UnsetSoldAt ensures that no value is present for SoldAt, not even an explicit nil
### GetInventoryItemId

`func (o *GetListingResponsePlatformListings) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *GetListingResponsePlatformListings) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *GetListingResponsePlatformListings) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *GetListingResponsePlatformListings) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### SetInventoryItemIdNil

`func (o *GetListingResponsePlatformListings) SetInventoryItemIdNil(b bool)`

 SetInventoryItemIdNil sets the value for InventoryItemId to be an explicit nil

### UnsetInventoryItemId
`func (o *GetListingResponsePlatformListings) UnsetInventoryItemId()`

UnsetInventoryItemId ensures that no value is present for InventoryItemId, not even an explicit nil
### GetPlatformListingId

`func (o *GetListingResponsePlatformListings) GetPlatformListingId() string`

GetPlatformListingId returns the PlatformListingId field if non-nil, zero value otherwise.

### GetPlatformListingIdOk

`func (o *GetListingResponsePlatformListings) GetPlatformListingIdOk() (*string, bool)`

GetPlatformListingIdOk returns a tuple with the PlatformListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformListingId

`func (o *GetListingResponsePlatformListings) SetPlatformListingId(v string)`

SetPlatformListingId sets PlatformListingId field to given value.


### GetPlatformListingUrl

`func (o *GetListingResponsePlatformListings) GetPlatformListingUrl() string`

GetPlatformListingUrl returns the PlatformListingUrl field if non-nil, zero value otherwise.

### GetPlatformListingUrlOk

`func (o *GetListingResponsePlatformListings) GetPlatformListingUrlOk() (*string, bool)`

GetPlatformListingUrlOk returns a tuple with the PlatformListingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformListingUrl

`func (o *GetListingResponsePlatformListings) SetPlatformListingUrl(v string)`

SetPlatformListingUrl sets PlatformListingUrl field to given value.

### HasPlatformListingUrl

`func (o *GetListingResponsePlatformListings) HasPlatformListingUrl() bool`

HasPlatformListingUrl returns a boolean if a field has been set.

### SetPlatformListingUrlNil

`func (o *GetListingResponsePlatformListings) SetPlatformListingUrlNil(b bool)`

 SetPlatformListingUrlNil sets the value for PlatformListingUrl to be an explicit nil

### UnsetPlatformListingUrl
`func (o *GetListingResponsePlatformListings) UnsetPlatformListingUrl()`

UnsetPlatformListingUrl ensures that no value is present for PlatformListingUrl, not even an explicit nil
### GetStockLocation

`func (o *GetListingResponsePlatformListings) GetStockLocation() string`

GetStockLocation returns the StockLocation field if non-nil, zero value otherwise.

### GetStockLocationOk

`func (o *GetListingResponsePlatformListings) GetStockLocationOk() (*string, bool)`

GetStockLocationOk returns a tuple with the StockLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStockLocation

`func (o *GetListingResponsePlatformListings) SetStockLocation(v string)`

SetStockLocation sets StockLocation field to given value.


### GetListedAt

`func (o *GetListingResponsePlatformListings) GetListedAt() time.Time`

GetListedAt returns the ListedAt field if non-nil, zero value otherwise.

### GetListedAtOk

`func (o *GetListingResponsePlatformListings) GetListedAtOk() (*time.Time, bool)`

GetListedAtOk returns a tuple with the ListedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedAt

`func (o *GetListingResponsePlatformListings) SetListedAt(v time.Time)`

SetListedAt sets ListedAt field to given value.

### HasListedAt

`func (o *GetListingResponsePlatformListings) HasListedAt() bool`

HasListedAt returns a boolean if a field has been set.

### SetListedAtNil

`func (o *GetListingResponsePlatformListings) SetListedAtNil(b bool)`

 SetListedAtNil sets the value for ListedAt to be an explicit nil

### UnsetListedAt
`func (o *GetListingResponsePlatformListings) UnsetListedAt()`

UnsetListedAt ensures that no value is present for ListedAt, not even an explicit nil
### GetSharedAt

`func (o *GetListingResponsePlatformListings) GetSharedAt() time.Time`

GetSharedAt returns the SharedAt field if non-nil, zero value otherwise.

### GetSharedAtOk

`func (o *GetListingResponsePlatformListings) GetSharedAtOk() (*time.Time, bool)`

GetSharedAtOk returns a tuple with the SharedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAt

`func (o *GetListingResponsePlatformListings) SetSharedAt(v time.Time)`

SetSharedAt sets SharedAt field to given value.

### HasSharedAt

`func (o *GetListingResponsePlatformListings) HasSharedAt() bool`

HasSharedAt returns a boolean if a field has been set.

### SetSharedAtNil

`func (o *GetListingResponsePlatformListings) SetSharedAtNil(b bool)`

 SetSharedAtNil sets the value for SharedAt to be an explicit nil

### UnsetSharedAt
`func (o *GetListingResponsePlatformListings) UnsetSharedAt()`

UnsetSharedAt ensures that no value is present for SharedAt, not even an explicit nil
### GetDelistReason

`func (o *GetListingResponsePlatformListings) GetDelistReason() string`

GetDelistReason returns the DelistReason field if non-nil, zero value otherwise.

### GetDelistReasonOk

`func (o *GetListingResponsePlatformListings) GetDelistReasonOk() (*string, bool)`

GetDelistReasonOk returns a tuple with the DelistReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelistReason

`func (o *GetListingResponsePlatformListings) SetDelistReason(v string)`

SetDelistReason sets DelistReason field to given value.

### HasDelistReason

`func (o *GetListingResponsePlatformListings) HasDelistReason() bool`

HasDelistReason returns a boolean if a field has been set.

### SetDelistReasonNil

`func (o *GetListingResponsePlatformListings) SetDelistReasonNil(b bool)`

 SetDelistReasonNil sets the value for DelistReason to be an explicit nil

### UnsetDelistReason
`func (o *GetListingResponsePlatformListings) UnsetDelistReason()`

UnsetDelistReason ensures that no value is present for DelistReason, not even an explicit nil
### GetLastCheckedAt

`func (o *GetListingResponsePlatformListings) GetLastCheckedAt() time.Time`

GetLastCheckedAt returns the LastCheckedAt field if non-nil, zero value otherwise.

### GetLastCheckedAtOk

`func (o *GetListingResponsePlatformListings) GetLastCheckedAtOk() (*time.Time, bool)`

GetLastCheckedAtOk returns a tuple with the LastCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCheckedAt

`func (o *GetListingResponsePlatformListings) SetLastCheckedAt(v time.Time)`

SetLastCheckedAt sets LastCheckedAt field to given value.

### HasLastCheckedAt

`func (o *GetListingResponsePlatformListings) HasLastCheckedAt() bool`

HasLastCheckedAt returns a boolean if a field has been set.

### SetLastCheckedAtNil

`func (o *GetListingResponsePlatformListings) SetLastCheckedAtNil(b bool)`

 SetLastCheckedAtNil sets the value for LastCheckedAt to be an explicit nil

### UnsetLastCheckedAt
`func (o *GetListingResponsePlatformListings) UnsetLastCheckedAt()`

UnsetLastCheckedAt ensures that no value is present for LastCheckedAt, not even an explicit nil
### GetViews

`func (o *GetListingResponsePlatformListings) GetViews() float32`

GetViews returns the Views field if non-nil, zero value otherwise.

### GetViewsOk

`func (o *GetListingResponsePlatformListings) GetViewsOk() (*float32, bool)`

GetViewsOk returns a tuple with the Views field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViews

`func (o *GetListingResponsePlatformListings) SetViews(v float32)`

SetViews sets Views field to given value.


### GetFavorites

`func (o *GetListingResponsePlatformListings) GetFavorites() float32`

GetFavorites returns the Favorites field if non-nil, zero value otherwise.

### GetFavoritesOk

`func (o *GetListingResponsePlatformListings) GetFavoritesOk() (*float32, bool)`

GetFavoritesOk returns a tuple with the Favorites field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFavorites

`func (o *GetListingResponsePlatformListings) SetFavorites(v float32)`

SetFavorites sets Favorites field to given value.


### GetCommentsCount

`func (o *GetListingResponsePlatformListings) GetCommentsCount() float32`

GetCommentsCount returns the CommentsCount field if non-nil, zero value otherwise.

### GetCommentsCountOk

`func (o *GetListingResponsePlatformListings) GetCommentsCountOk() (*float32, bool)`

GetCommentsCountOk returns a tuple with the CommentsCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommentsCount

`func (o *GetListingResponsePlatformListings) SetCommentsCount(v float32)`

SetCommentsCount sets CommentsCount field to given value.


### GetLastStatsSyncedAt

`func (o *GetListingResponsePlatformListings) GetLastStatsSyncedAt() time.Time`

GetLastStatsSyncedAt returns the LastStatsSyncedAt field if non-nil, zero value otherwise.

### GetLastStatsSyncedAtOk

`func (o *GetListingResponsePlatformListings) GetLastStatsSyncedAtOk() (*time.Time, bool)`

GetLastStatsSyncedAtOk returns a tuple with the LastStatsSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatsSyncedAt

`func (o *GetListingResponsePlatformListings) SetLastStatsSyncedAt(v time.Time)`

SetLastStatsSyncedAt sets LastStatsSyncedAt field to given value.

### HasLastStatsSyncedAt

`func (o *GetListingResponsePlatformListings) HasLastStatsSyncedAt() bool`

HasLastStatsSyncedAt returns a boolean if a field has been set.

### SetLastStatsSyncedAtNil

`func (o *GetListingResponsePlatformListings) SetLastStatsSyncedAtNil(b bool)`

 SetLastStatsSyncedAtNil sets the value for LastStatsSyncedAt to be an explicit nil

### UnsetLastStatsSyncedAt
`func (o *GetListingResponsePlatformListings) UnsetLastStatsSyncedAt()`

UnsetLastStatsSyncedAt ensures that no value is present for LastStatsSyncedAt, not even an explicit nil
### GetHashtags

`func (o *GetListingResponsePlatformListings) GetHashtags() []string`

GetHashtags returns the Hashtags field if non-nil, zero value otherwise.

### GetHashtagsOk

`func (o *GetListingResponsePlatformListings) GetHashtagsOk() (*[]string, bool)`

GetHashtagsOk returns a tuple with the Hashtags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHashtags

`func (o *GetListingResponsePlatformListings) SetHashtags(v []string)`

SetHashtags sets Hashtags field to given value.

### HasHashtags

`func (o *GetListingResponsePlatformListings) HasHashtags() bool`

HasHashtags returns a boolean if a field has been set.

### SetHashtagsNil

`func (o *GetListingResponsePlatformListings) SetHashtagsNil(b bool)`

 SetHashtagsNil sets the value for Hashtags to be an explicit nil

### UnsetHashtags
`func (o *GetListingResponsePlatformListings) UnsetHashtags()`

UnsetHashtags ensures that no value is present for Hashtags, not even an explicit nil
### GetIsDraft

`func (o *GetListingResponsePlatformListings) GetIsDraft() bool`

GetIsDraft returns the IsDraft field if non-nil, zero value otherwise.

### GetIsDraftOk

`func (o *GetListingResponsePlatformListings) GetIsDraftOk() (*bool, bool)`

GetIsDraftOk returns a tuple with the IsDraft field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDraft

`func (o *GetListingResponsePlatformListings) SetIsDraft(v bool)`

SetIsDraft sets IsDraft field to given value.


### GetScheduledPublishAt

`func (o *GetListingResponsePlatformListings) GetScheduledPublishAt() time.Time`

GetScheduledPublishAt returns the ScheduledPublishAt field if non-nil, zero value otherwise.

### GetScheduledPublishAtOk

`func (o *GetListingResponsePlatformListings) GetScheduledPublishAtOk() (*time.Time, bool)`

GetScheduledPublishAtOk returns a tuple with the ScheduledPublishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPublishAt

`func (o *GetListingResponsePlatformListings) SetScheduledPublishAt(v time.Time)`

SetScheduledPublishAt sets ScheduledPublishAt field to given value.

### HasScheduledPublishAt

`func (o *GetListingResponsePlatformListings) HasScheduledPublishAt() bool`

HasScheduledPublishAt returns a boolean if a field has been set.

### SetScheduledPublishAtNil

`func (o *GetListingResponsePlatformListings) SetScheduledPublishAtNil(b bool)`

 SetScheduledPublishAtNil sets the value for ScheduledPublishAt to be an explicit nil

### UnsetScheduledPublishAt
`func (o *GetListingResponsePlatformListings) UnsetScheduledPublishAt()`

UnsetScheduledPublishAt ensures that no value is present for ScheduledPublishAt, not even an explicit nil
### GetIsDutchAuction

`func (o *GetListingResponsePlatformListings) GetIsDutchAuction() bool`

GetIsDutchAuction returns the IsDutchAuction field if non-nil, zero value otherwise.

### GetIsDutchAuctionOk

`func (o *GetListingResponsePlatformListings) GetIsDutchAuctionOk() (*bool, bool)`

GetIsDutchAuctionOk returns a tuple with the IsDutchAuction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDutchAuction

`func (o *GetListingResponsePlatformListings) SetIsDutchAuction(v bool)`

SetIsDutchAuction sets IsDutchAuction field to given value.


### GetDutchStartCents

`func (o *GetListingResponsePlatformListings) GetDutchStartCents() float32`

GetDutchStartCents returns the DutchStartCents field if non-nil, zero value otherwise.

### GetDutchStartCentsOk

`func (o *GetListingResponsePlatformListings) GetDutchStartCentsOk() (*float32, bool)`

GetDutchStartCentsOk returns a tuple with the DutchStartCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutchStartCents

`func (o *GetListingResponsePlatformListings) SetDutchStartCents(v float32)`

SetDutchStartCents sets DutchStartCents field to given value.

### HasDutchStartCents

`func (o *GetListingResponsePlatformListings) HasDutchStartCents() bool`

HasDutchStartCents returns a boolean if a field has been set.

### SetDutchStartCentsNil

`func (o *GetListingResponsePlatformListings) SetDutchStartCentsNil(b bool)`

 SetDutchStartCentsNil sets the value for DutchStartCents to be an explicit nil

### UnsetDutchStartCents
`func (o *GetListingResponsePlatformListings) UnsetDutchStartCents()`

UnsetDutchStartCents ensures that no value is present for DutchStartCents, not even an explicit nil
### GetDutchFloorCents

`func (o *GetListingResponsePlatformListings) GetDutchFloorCents() float32`

GetDutchFloorCents returns the DutchFloorCents field if non-nil, zero value otherwise.

### GetDutchFloorCentsOk

`func (o *GetListingResponsePlatformListings) GetDutchFloorCentsOk() (*float32, bool)`

GetDutchFloorCentsOk returns a tuple with the DutchFloorCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutchFloorCents

`func (o *GetListingResponsePlatformListings) SetDutchFloorCents(v float32)`

SetDutchFloorCents sets DutchFloorCents field to given value.

### HasDutchFloorCents

`func (o *GetListingResponsePlatformListings) HasDutchFloorCents() bool`

HasDutchFloorCents returns a boolean if a field has been set.

### SetDutchFloorCentsNil

`func (o *GetListingResponsePlatformListings) SetDutchFloorCentsNil(b bool)`

 SetDutchFloorCentsNil sets the value for DutchFloorCents to be an explicit nil

### UnsetDutchFloorCents
`func (o *GetListingResponsePlatformListings) UnsetDutchFloorCents()`

UnsetDutchFloorCents ensures that no value is present for DutchFloorCents, not even an explicit nil
### GetDutchEndsAt

`func (o *GetListingResponsePlatformListings) GetDutchEndsAt() time.Time`

GetDutchEndsAt returns the DutchEndsAt field if non-nil, zero value otherwise.

### GetDutchEndsAtOk

`func (o *GetListingResponsePlatformListings) GetDutchEndsAtOk() (*time.Time, bool)`

GetDutchEndsAtOk returns a tuple with the DutchEndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDutchEndsAt

`func (o *GetListingResponsePlatformListings) SetDutchEndsAt(v time.Time)`

SetDutchEndsAt sets DutchEndsAt field to given value.

### HasDutchEndsAt

`func (o *GetListingResponsePlatformListings) HasDutchEndsAt() bool`

HasDutchEndsAt returns a boolean if a field has been set.

### SetDutchEndsAtNil

`func (o *GetListingResponsePlatformListings) SetDutchEndsAtNil(b bool)`

 SetDutchEndsAtNil sets the value for DutchEndsAt to be an explicit nil

### UnsetDutchEndsAt
`func (o *GetListingResponsePlatformListings) UnsetDutchEndsAt()`

UnsetDutchEndsAt ensures that no value is present for DutchEndsAt, not even an explicit nil
### GetIsPreOrder

`func (o *GetListingResponsePlatformListings) GetIsPreOrder() bool`

GetIsPreOrder returns the IsPreOrder field if non-nil, zero value otherwise.

### GetIsPreOrderOk

`func (o *GetListingResponsePlatformListings) GetIsPreOrderOk() (*bool, bool)`

GetIsPreOrderOk returns a tuple with the IsPreOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPreOrder

`func (o *GetListingResponsePlatformListings) SetIsPreOrder(v bool)`

SetIsPreOrder sets IsPreOrder field to given value.


### GetExpectedShipBy

`func (o *GetListingResponsePlatformListings) GetExpectedShipBy() string`

GetExpectedShipBy returns the ExpectedShipBy field if non-nil, zero value otherwise.

### GetExpectedShipByOk

`func (o *GetListingResponsePlatformListings) GetExpectedShipByOk() (*string, bool)`

GetExpectedShipByOk returns a tuple with the ExpectedShipBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedShipBy

`func (o *GetListingResponsePlatformListings) SetExpectedShipBy(v string)`

SetExpectedShipBy sets ExpectedShipBy field to given value.

### HasExpectedShipBy

`func (o *GetListingResponsePlatformListings) HasExpectedShipBy() bool`

HasExpectedShipBy returns a boolean if a field has been set.

### SetExpectedShipByNil

`func (o *GetListingResponsePlatformListings) SetExpectedShipByNil(b bool)`

 SetExpectedShipByNil sets the value for ExpectedShipBy to be an explicit nil

### UnsetExpectedShipBy
`func (o *GetListingResponsePlatformListings) UnsetExpectedShipBy()`

UnsetExpectedShipBy ensures that no value is present for ExpectedShipBy, not even an explicit nil
### GetPreorderDepositPercentBps

`func (o *GetListingResponsePlatformListings) GetPreorderDepositPercentBps() float32`

GetPreorderDepositPercentBps returns the PreorderDepositPercentBps field if non-nil, zero value otherwise.

### GetPreorderDepositPercentBpsOk

`func (o *GetListingResponsePlatformListings) GetPreorderDepositPercentBpsOk() (*float32, bool)`

GetPreorderDepositPercentBpsOk returns a tuple with the PreorderDepositPercentBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreorderDepositPercentBps

`func (o *GetListingResponsePlatformListings) SetPreorderDepositPercentBps(v float32)`

SetPreorderDepositPercentBps sets PreorderDepositPercentBps field to given value.

### HasPreorderDepositPercentBps

`func (o *GetListingResponsePlatformListings) HasPreorderDepositPercentBps() bool`

HasPreorderDepositPercentBps returns a boolean if a field has been set.

### SetPreorderDepositPercentBpsNil

`func (o *GetListingResponsePlatformListings) SetPreorderDepositPercentBpsNil(b bool)`

 SetPreorderDepositPercentBpsNil sets the value for PreorderDepositPercentBps to be an explicit nil

### UnsetPreorderDepositPercentBps
`func (o *GetListingResponsePlatformListings) UnsetPreorderDepositPercentBps()`

UnsetPreorderDepositPercentBps ensures that no value is present for PreorderDepositPercentBps, not even an explicit nil
### GetBoostedAt

`func (o *GetListingResponsePlatformListings) GetBoostedAt() time.Time`

GetBoostedAt returns the BoostedAt field if non-nil, zero value otherwise.

### GetBoostedAtOk

`func (o *GetListingResponsePlatformListings) GetBoostedAtOk() (*time.Time, bool)`

GetBoostedAtOk returns a tuple with the BoostedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostedAt

`func (o *GetListingResponsePlatformListings) SetBoostedAt(v time.Time)`

SetBoostedAt sets BoostedAt field to given value.

### HasBoostedAt

`func (o *GetListingResponsePlatformListings) HasBoostedAt() bool`

HasBoostedAt returns a boolean if a field has been set.

### SetBoostedAtNil

`func (o *GetListingResponsePlatformListings) SetBoostedAtNil(b bool)`

 SetBoostedAtNil sets the value for BoostedAt to be an explicit nil

### UnsetBoostedAt
`func (o *GetListingResponsePlatformListings) UnsetBoostedAt()`

UnsetBoostedAt ensures that no value is present for BoostedAt, not even an explicit nil
### GetImageCaption

`func (o *GetListingResponsePlatformListings) GetImageCaption() string`

GetImageCaption returns the ImageCaption field if non-nil, zero value otherwise.

### GetImageCaptionOk

`func (o *GetListingResponsePlatformListings) GetImageCaptionOk() (*string, bool)`

GetImageCaptionOk returns a tuple with the ImageCaption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageCaption

`func (o *GetListingResponsePlatformListings) SetImageCaption(v string)`

SetImageCaption sets ImageCaption field to given value.

### HasImageCaption

`func (o *GetListingResponsePlatformListings) HasImageCaption() bool`

HasImageCaption returns a boolean if a field has been set.

### SetImageCaptionNil

`func (o *GetListingResponsePlatformListings) SetImageCaptionNil(b bool)`

 SetImageCaptionNil sets the value for ImageCaption to be an explicit nil

### UnsetImageCaption
`func (o *GetListingResponsePlatformListings) UnsetImageCaption()`

UnsetImageCaption ensures that no value is present for ImageCaption, not even an explicit nil
### GetDurationDays

`func (o *GetListingResponsePlatformListings) GetDurationDays() float32`

GetDurationDays returns the DurationDays field if non-nil, zero value otherwise.

### GetDurationDaysOk

`func (o *GetListingResponsePlatformListings) GetDurationDaysOk() (*float32, bool)`

GetDurationDaysOk returns a tuple with the DurationDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationDays

`func (o *GetListingResponsePlatformListings) SetDurationDays(v float32)`

SetDurationDays sets DurationDays field to given value.

### HasDurationDays

`func (o *GetListingResponsePlatformListings) HasDurationDays() bool`

HasDurationDays returns a boolean if a field has been set.

### SetDurationDaysNil

`func (o *GetListingResponsePlatformListings) SetDurationDaysNil(b bool)`

 SetDurationDaysNil sets the value for DurationDays to be an explicit nil

### UnsetDurationDays
`func (o *GetListingResponsePlatformListings) UnsetDurationDays()`

UnsetDurationDays ensures that no value is present for DurationDays, not even an explicit nil
### GetEndsAt

`func (o *GetListingResponsePlatformListings) GetEndsAt() time.Time`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *GetListingResponsePlatformListings) GetEndsAtOk() (*time.Time, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *GetListingResponsePlatformListings) SetEndsAt(v time.Time)`

SetEndsAt sets EndsAt field to given value.

### HasEndsAt

`func (o *GetListingResponsePlatformListings) HasEndsAt() bool`

HasEndsAt returns a boolean if a field has been set.

### SetEndsAtNil

`func (o *GetListingResponsePlatformListings) SetEndsAtNil(b bool)`

 SetEndsAtNil sets the value for EndsAt to be an explicit nil

### UnsetEndsAt
`func (o *GetListingResponsePlatformListings) UnsetEndsAt()`

UnsetEndsAt ensures that no value is present for EndsAt, not even an explicit nil
### GetBestOfferEnabled

`func (o *GetListingResponsePlatformListings) GetBestOfferEnabled() bool`

GetBestOfferEnabled returns the BestOfferEnabled field if non-nil, zero value otherwise.

### GetBestOfferEnabledOk

`func (o *GetListingResponsePlatformListings) GetBestOfferEnabledOk() (*bool, bool)`

GetBestOfferEnabledOk returns a tuple with the BestOfferEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestOfferEnabled

`func (o *GetListingResponsePlatformListings) SetBestOfferEnabled(v bool)`

SetBestOfferEnabled sets BestOfferEnabled field to given value.


### GetBestOfferAutoAcceptCents

`func (o *GetListingResponsePlatformListings) GetBestOfferAutoAcceptCents() float32`

GetBestOfferAutoAcceptCents returns the BestOfferAutoAcceptCents field if non-nil, zero value otherwise.

### GetBestOfferAutoAcceptCentsOk

`func (o *GetListingResponsePlatformListings) GetBestOfferAutoAcceptCentsOk() (*float32, bool)`

GetBestOfferAutoAcceptCentsOk returns a tuple with the BestOfferAutoAcceptCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestOfferAutoAcceptCents

`func (o *GetListingResponsePlatformListings) SetBestOfferAutoAcceptCents(v float32)`

SetBestOfferAutoAcceptCents sets BestOfferAutoAcceptCents field to given value.

### HasBestOfferAutoAcceptCents

`func (o *GetListingResponsePlatformListings) HasBestOfferAutoAcceptCents() bool`

HasBestOfferAutoAcceptCents returns a boolean if a field has been set.

### SetBestOfferAutoAcceptCentsNil

`func (o *GetListingResponsePlatformListings) SetBestOfferAutoAcceptCentsNil(b bool)`

 SetBestOfferAutoAcceptCentsNil sets the value for BestOfferAutoAcceptCents to be an explicit nil

### UnsetBestOfferAutoAcceptCents
`func (o *GetListingResponsePlatformListings) UnsetBestOfferAutoAcceptCents()`

UnsetBestOfferAutoAcceptCents ensures that no value is present for BestOfferAutoAcceptCents, not even an explicit nil
### GetBestOfferAutoDeclineCents

`func (o *GetListingResponsePlatformListings) GetBestOfferAutoDeclineCents() float32`

GetBestOfferAutoDeclineCents returns the BestOfferAutoDeclineCents field if non-nil, zero value otherwise.

### GetBestOfferAutoDeclineCentsOk

`func (o *GetListingResponsePlatformListings) GetBestOfferAutoDeclineCentsOk() (*float32, bool)`

GetBestOfferAutoDeclineCentsOk returns a tuple with the BestOfferAutoDeclineCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestOfferAutoDeclineCents

`func (o *GetListingResponsePlatformListings) SetBestOfferAutoDeclineCents(v float32)`

SetBestOfferAutoDeclineCents sets BestOfferAutoDeclineCents field to given value.

### HasBestOfferAutoDeclineCents

`func (o *GetListingResponsePlatformListings) HasBestOfferAutoDeclineCents() bool`

HasBestOfferAutoDeclineCents returns a boolean if a field has been set.

### SetBestOfferAutoDeclineCentsNil

`func (o *GetListingResponsePlatformListings) SetBestOfferAutoDeclineCentsNil(b bool)`

 SetBestOfferAutoDeclineCentsNil sets the value for BestOfferAutoDeclineCents to be an explicit nil

### UnsetBestOfferAutoDeclineCents
`func (o *GetListingResponsePlatformListings) UnsetBestOfferAutoDeclineCents()`

UnsetBestOfferAutoDeclineCents ensures that no value is present for BestOfferAutoDeclineCents, not even an explicit nil
### GetAestheticTags

`func (o *GetListingResponsePlatformListings) GetAestheticTags() []string`

GetAestheticTags returns the AestheticTags field if non-nil, zero value otherwise.

### GetAestheticTagsOk

`func (o *GetListingResponsePlatformListings) GetAestheticTagsOk() (*[]string, bool)`

GetAestheticTagsOk returns a tuple with the AestheticTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAestheticTags

`func (o *GetListingResponsePlatformListings) SetAestheticTags(v []string)`

SetAestheticTags sets AestheticTags field to given value.

### HasAestheticTags

`func (o *GetListingResponsePlatformListings) HasAestheticTags() bool`

HasAestheticTags returns a boolean if a field has been set.

### SetAestheticTagsNil

`func (o *GetListingResponsePlatformListings) SetAestheticTagsNil(b bool)`

 SetAestheticTagsNil sets the value for AestheticTags to be an explicit nil

### UnsetAestheticTags
`func (o *GetListingResponsePlatformListings) UnsetAestheticTags()`

UnsetAestheticTags ensures that no value is present for AestheticTags, not even an explicit nil
### GetAllowsLocalPickup

`func (o *GetListingResponsePlatformListings) GetAllowsLocalPickup() bool`

GetAllowsLocalPickup returns the AllowsLocalPickup field if non-nil, zero value otherwise.

### GetAllowsLocalPickupOk

`func (o *GetListingResponsePlatformListings) GetAllowsLocalPickupOk() (*bool, bool)`

GetAllowsLocalPickupOk returns a tuple with the AllowsLocalPickup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowsLocalPickup

`func (o *GetListingResponsePlatformListings) SetAllowsLocalPickup(v bool)`

SetAllowsLocalPickup sets AllowsLocalPickup field to given value.


### GetPickupRadiusMiles

`func (o *GetListingResponsePlatformListings) GetPickupRadiusMiles() float32`

GetPickupRadiusMiles returns the PickupRadiusMiles field if non-nil, zero value otherwise.

### GetPickupRadiusMilesOk

`func (o *GetListingResponsePlatformListings) GetPickupRadiusMilesOk() (*float32, bool)`

GetPickupRadiusMilesOk returns a tuple with the PickupRadiusMiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupRadiusMiles

`func (o *GetListingResponsePlatformListings) SetPickupRadiusMiles(v float32)`

SetPickupRadiusMiles sets PickupRadiusMiles field to given value.

### HasPickupRadiusMiles

`func (o *GetListingResponsePlatformListings) HasPickupRadiusMiles() bool`

HasPickupRadiusMiles returns a boolean if a field has been set.

### SetPickupRadiusMilesNil

`func (o *GetListingResponsePlatformListings) SetPickupRadiusMilesNil(b bool)`

 SetPickupRadiusMilesNil sets the value for PickupRadiusMiles to be an explicit nil

### UnsetPickupRadiusMiles
`func (o *GetListingResponsePlatformListings) UnsetPickupRadiusMiles()`

UnsetPickupRadiusMiles ensures that no value is present for PickupRadiusMiles, not even an explicit nil
### GetPickupLat

`func (o *GetListingResponsePlatformListings) GetPickupLat() string`

GetPickupLat returns the PickupLat field if non-nil, zero value otherwise.

### GetPickupLatOk

`func (o *GetListingResponsePlatformListings) GetPickupLatOk() (*string, bool)`

GetPickupLatOk returns a tuple with the PickupLat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupLat

`func (o *GetListingResponsePlatformListings) SetPickupLat(v string)`

SetPickupLat sets PickupLat field to given value.

### HasPickupLat

`func (o *GetListingResponsePlatformListings) HasPickupLat() bool`

HasPickupLat returns a boolean if a field has been set.

### SetPickupLatNil

`func (o *GetListingResponsePlatformListings) SetPickupLatNil(b bool)`

 SetPickupLatNil sets the value for PickupLat to be an explicit nil

### UnsetPickupLat
`func (o *GetListingResponsePlatformListings) UnsetPickupLat()`

UnsetPickupLat ensures that no value is present for PickupLat, not even an explicit nil
### GetPickupLng

`func (o *GetListingResponsePlatformListings) GetPickupLng() string`

GetPickupLng returns the PickupLng field if non-nil, zero value otherwise.

### GetPickupLngOk

`func (o *GetListingResponsePlatformListings) GetPickupLngOk() (*string, bool)`

GetPickupLngOk returns a tuple with the PickupLng field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupLng

`func (o *GetListingResponsePlatformListings) SetPickupLng(v string)`

SetPickupLng sets PickupLng field to given value.

### HasPickupLng

`func (o *GetListingResponsePlatformListings) HasPickupLng() bool`

HasPickupLng returns a boolean if a field has been set.

### SetPickupLngNil

`func (o *GetListingResponsePlatformListings) SetPickupLngNil(b bool)`

 SetPickupLngNil sets the value for PickupLng to be an explicit nil

### UnsetPickupLng
`func (o *GetListingResponsePlatformListings) UnsetPickupLng()`

UnsetPickupLng ensures that no value is present for PickupLng, not even an explicit nil
### GetRefurbishedTier

`func (o *GetListingResponsePlatformListings) GetRefurbishedTier() string`

GetRefurbishedTier returns the RefurbishedTier field if non-nil, zero value otherwise.

### GetRefurbishedTierOk

`func (o *GetListingResponsePlatformListings) GetRefurbishedTierOk() (*string, bool)`

GetRefurbishedTierOk returns a tuple with the RefurbishedTier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefurbishedTier

`func (o *GetListingResponsePlatformListings) SetRefurbishedTier(v string)`

SetRefurbishedTier sets RefurbishedTier field to given value.

### HasRefurbishedTier

`func (o *GetListingResponsePlatformListings) HasRefurbishedTier() bool`

HasRefurbishedTier returns a boolean if a field has been set.

### SetRefurbishedTierNil

`func (o *GetListingResponsePlatformListings) SetRefurbishedTierNil(b bool)`

 SetRefurbishedTierNil sets the value for RefurbishedTier to be an explicit nil

### UnsetRefurbishedTier
`func (o *GetListingResponsePlatformListings) UnsetRefurbishedTier()`

UnsetRefurbishedTier ensures that no value is present for RefurbishedTier, not even an explicit nil
### GetCharitySlug

`func (o *GetListingResponsePlatformListings) GetCharitySlug() string`

GetCharitySlug returns the CharitySlug field if non-nil, zero value otherwise.

### GetCharitySlugOk

`func (o *GetListingResponsePlatformListings) GetCharitySlugOk() (*string, bool)`

GetCharitySlugOk returns a tuple with the CharitySlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharitySlug

`func (o *GetListingResponsePlatformListings) SetCharitySlug(v string)`

SetCharitySlug sets CharitySlug field to given value.

### HasCharitySlug

`func (o *GetListingResponsePlatformListings) HasCharitySlug() bool`

HasCharitySlug returns a boolean if a field has been set.

### SetCharitySlugNil

`func (o *GetListingResponsePlatformListings) SetCharitySlugNil(b bool)`

 SetCharitySlugNil sets the value for CharitySlug to be an explicit nil

### UnsetCharitySlug
`func (o *GetListingResponsePlatformListings) UnsetCharitySlug()`

UnsetCharitySlug ensures that no value is present for CharitySlug, not even an explicit nil
### GetCharityPercentBps

`func (o *GetListingResponsePlatformListings) GetCharityPercentBps() float32`

GetCharityPercentBps returns the CharityPercentBps field if non-nil, zero value otherwise.

### GetCharityPercentBpsOk

`func (o *GetListingResponsePlatformListings) GetCharityPercentBpsOk() (*float32, bool)`

GetCharityPercentBpsOk returns a tuple with the CharityPercentBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharityPercentBps

`func (o *GetListingResponsePlatformListings) SetCharityPercentBps(v float32)`

SetCharityPercentBps sets CharityPercentBps field to given value.

### HasCharityPercentBps

`func (o *GetListingResponsePlatformListings) HasCharityPercentBps() bool`

HasCharityPercentBps returns a boolean if a field has been set.

### SetCharityPercentBpsNil

`func (o *GetListingResponsePlatformListings) SetCharityPercentBpsNil(b bool)`

 SetCharityPercentBpsNil sets the value for CharityPercentBps to be an explicit nil

### UnsetCharityPercentBps
`func (o *GetListingResponsePlatformListings) UnsetCharityPercentBps()`

UnsetCharityPercentBps ensures that no value is present for CharityPercentBps, not even an explicit nil
### GetCrosslyReturnPolicyId

`func (o *GetListingResponsePlatformListings) GetCrosslyReturnPolicyId() string`

GetCrosslyReturnPolicyId returns the CrosslyReturnPolicyId field if non-nil, zero value otherwise.

### GetCrosslyReturnPolicyIdOk

`func (o *GetListingResponsePlatformListings) GetCrosslyReturnPolicyIdOk() (*string, bool)`

GetCrosslyReturnPolicyIdOk returns a tuple with the CrosslyReturnPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrosslyReturnPolicyId

`func (o *GetListingResponsePlatformListings) SetCrosslyReturnPolicyId(v string)`

SetCrosslyReturnPolicyId sets CrosslyReturnPolicyId field to given value.

### HasCrosslyReturnPolicyId

`func (o *GetListingResponsePlatformListings) HasCrosslyReturnPolicyId() bool`

HasCrosslyReturnPolicyId returns a boolean if a field has been set.

### SetCrosslyReturnPolicyIdNil

`func (o *GetListingResponsePlatformListings) SetCrosslyReturnPolicyIdNil(b bool)`

 SetCrosslyReturnPolicyIdNil sets the value for CrosslyReturnPolicyId to be an explicit nil

### UnsetCrosslyReturnPolicyId
`func (o *GetListingResponsePlatformListings) UnsetCrosslyReturnPolicyId()`

UnsetCrosslyReturnPolicyId ensures that no value is present for CrosslyReturnPolicyId, not even an explicit nil
### GetCrosslyShippingPolicyId

`func (o *GetListingResponsePlatformListings) GetCrosslyShippingPolicyId() string`

GetCrosslyShippingPolicyId returns the CrosslyShippingPolicyId field if non-nil, zero value otherwise.

### GetCrosslyShippingPolicyIdOk

`func (o *GetListingResponsePlatformListings) GetCrosslyShippingPolicyIdOk() (*string, bool)`

GetCrosslyShippingPolicyIdOk returns a tuple with the CrosslyShippingPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrosslyShippingPolicyId

`func (o *GetListingResponsePlatformListings) SetCrosslyShippingPolicyId(v string)`

SetCrosslyShippingPolicyId sets CrosslyShippingPolicyId field to given value.

### HasCrosslyShippingPolicyId

`func (o *GetListingResponsePlatformListings) HasCrosslyShippingPolicyId() bool`

HasCrosslyShippingPolicyId returns a boolean if a field has been set.

### SetCrosslyShippingPolicyIdNil

`func (o *GetListingResponsePlatformListings) SetCrosslyShippingPolicyIdNil(b bool)`

 SetCrosslyShippingPolicyIdNil sets the value for CrosslyShippingPolicyId to be an explicit nil

### UnsetCrosslyShippingPolicyId
`func (o *GetListingResponsePlatformListings) UnsetCrosslyShippingPolicyId()`

UnsetCrosslyShippingPolicyId ensures that no value is present for CrosslyShippingPolicyId, not even an explicit nil
### GetCrosslyPaymentPolicyId

`func (o *GetListingResponsePlatformListings) GetCrosslyPaymentPolicyId() string`

GetCrosslyPaymentPolicyId returns the CrosslyPaymentPolicyId field if non-nil, zero value otherwise.

### GetCrosslyPaymentPolicyIdOk

`func (o *GetListingResponsePlatformListings) GetCrosslyPaymentPolicyIdOk() (*string, bool)`

GetCrosslyPaymentPolicyIdOk returns a tuple with the CrosslyPaymentPolicyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrosslyPaymentPolicyId

`func (o *GetListingResponsePlatformListings) SetCrosslyPaymentPolicyId(v string)`

SetCrosslyPaymentPolicyId sets CrosslyPaymentPolicyId field to given value.

### HasCrosslyPaymentPolicyId

`func (o *GetListingResponsePlatformListings) HasCrosslyPaymentPolicyId() bool`

HasCrosslyPaymentPolicyId returns a boolean if a field has been set.

### SetCrosslyPaymentPolicyIdNil

`func (o *GetListingResponsePlatformListings) SetCrosslyPaymentPolicyIdNil(b bool)`

 SetCrosslyPaymentPolicyIdNil sets the value for CrosslyPaymentPolicyId to be an explicit nil

### UnsetCrosslyPaymentPolicyId
`func (o *GetListingResponsePlatformListings) UnsetCrosslyPaymentPolicyId()`

UnsetCrosslyPaymentPolicyId ensures that no value is present for CrosslyPaymentPolicyId, not even an explicit nil
### GetRetailPriceCents

`func (o *GetListingResponsePlatformListings) GetRetailPriceCents() float32`

GetRetailPriceCents returns the RetailPriceCents field if non-nil, zero value otherwise.

### GetRetailPriceCentsOk

`func (o *GetListingResponsePlatformListings) GetRetailPriceCentsOk() (*float32, bool)`

GetRetailPriceCentsOk returns a tuple with the RetailPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailPriceCents

`func (o *GetListingResponsePlatformListings) SetRetailPriceCents(v float32)`

SetRetailPriceCents sets RetailPriceCents field to given value.

### HasRetailPriceCents

`func (o *GetListingResponsePlatformListings) HasRetailPriceCents() bool`

HasRetailPriceCents returns a boolean if a field has been set.

### SetRetailPriceCentsNil

`func (o *GetListingResponsePlatformListings) SetRetailPriceCentsNil(b bool)`

 SetRetailPriceCentsNil sets the value for RetailPriceCents to be an explicit nil

### UnsetRetailPriceCents
`func (o *GetListingResponsePlatformListings) UnsetRetailPriceCents()`

UnsetRetailPriceCents ensures that no value is present for RetailPriceCents, not even an explicit nil
### GetRestrictedCountries

`func (o *GetListingResponsePlatformListings) GetRestrictedCountries() []string`

GetRestrictedCountries returns the RestrictedCountries field if non-nil, zero value otherwise.

### GetRestrictedCountriesOk

`func (o *GetListingResponsePlatformListings) GetRestrictedCountriesOk() (*[]string, bool)`

GetRestrictedCountriesOk returns a tuple with the RestrictedCountries field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestrictedCountries

`func (o *GetListingResponsePlatformListings) SetRestrictedCountries(v []string)`

SetRestrictedCountries sets RestrictedCountries field to given value.

### HasRestrictedCountries

`func (o *GetListingResponsePlatformListings) HasRestrictedCountries() bool`

HasRestrictedCountries returns a boolean if a field has been set.

### SetRestrictedCountriesNil

`func (o *GetListingResponsePlatformListings) SetRestrictedCountriesNil(b bool)`

 SetRestrictedCountriesNil sets the value for RestrictedCountries to be an explicit nil

### UnsetRestrictedCountries
`func (o *GetListingResponsePlatformListings) UnsetRestrictedCountries()`

UnsetRestrictedCountries ensures that no value is present for RestrictedCountries, not even an explicit nil
### GetAuctionExtendSeconds

`func (o *GetListingResponsePlatformListings) GetAuctionExtendSeconds() float32`

GetAuctionExtendSeconds returns the AuctionExtendSeconds field if non-nil, zero value otherwise.

### GetAuctionExtendSecondsOk

`func (o *GetListingResponsePlatformListings) GetAuctionExtendSecondsOk() (*float32, bool)`

GetAuctionExtendSecondsOk returns a tuple with the AuctionExtendSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionExtendSeconds

`func (o *GetListingResponsePlatformListings) SetAuctionExtendSeconds(v float32)`

SetAuctionExtendSeconds sets AuctionExtendSeconds field to given value.

### HasAuctionExtendSeconds

`func (o *GetListingResponsePlatformListings) HasAuctionExtendSeconds() bool`

HasAuctionExtendSeconds returns a boolean if a field has been set.

### SetAuctionExtendSecondsNil

`func (o *GetListingResponsePlatformListings) SetAuctionExtendSecondsNil(b bool)`

 SetAuctionExtendSecondsNil sets the value for AuctionExtendSeconds to be an explicit nil

### UnsetAuctionExtendSeconds
`func (o *GetListingResponsePlatformListings) UnsetAuctionExtendSeconds()`

UnsetAuctionExtendSeconds ensures that no value is present for AuctionExtendSeconds, not even an explicit nil
### GetAuthenticationProviderSlug

`func (o *GetListingResponsePlatformListings) GetAuthenticationProviderSlug() string`

GetAuthenticationProviderSlug returns the AuthenticationProviderSlug field if non-nil, zero value otherwise.

### GetAuthenticationProviderSlugOk

`func (o *GetListingResponsePlatformListings) GetAuthenticationProviderSlugOk() (*string, bool)`

GetAuthenticationProviderSlugOk returns a tuple with the AuthenticationProviderSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationProviderSlug

`func (o *GetListingResponsePlatformListings) SetAuthenticationProviderSlug(v string)`

SetAuthenticationProviderSlug sets AuthenticationProviderSlug field to given value.

### HasAuthenticationProviderSlug

`func (o *GetListingResponsePlatformListings) HasAuthenticationProviderSlug() bool`

HasAuthenticationProviderSlug returns a boolean if a field has been set.

### SetAuthenticationProviderSlugNil

`func (o *GetListingResponsePlatformListings) SetAuthenticationProviderSlugNil(b bool)`

 SetAuthenticationProviderSlugNil sets the value for AuthenticationProviderSlug to be an explicit nil

### UnsetAuthenticationProviderSlug
`func (o *GetListingResponsePlatformListings) UnsetAuthenticationProviderSlug()`

UnsetAuthenticationProviderSlug ensures that no value is present for AuthenticationProviderSlug, not even an explicit nil
### GetAuthenticationRequired

`func (o *GetListingResponsePlatformListings) GetAuthenticationRequired() bool`

GetAuthenticationRequired returns the AuthenticationRequired field if non-nil, zero value otherwise.

### GetAuthenticationRequiredOk

`func (o *GetListingResponsePlatformListings) GetAuthenticationRequiredOk() (*bool, bool)`

GetAuthenticationRequiredOk returns a tuple with the AuthenticationRequired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthenticationRequired

`func (o *GetListingResponsePlatformListings) SetAuthenticationRequired(v bool)`

SetAuthenticationRequired sets AuthenticationRequired field to given value.


### GetHasVariants

`func (o *GetListingResponsePlatformListings) GetHasVariants() bool`

GetHasVariants returns the HasVariants field if non-nil, zero value otherwise.

### GetHasVariantsOk

`func (o *GetListingResponsePlatformListings) GetHasVariantsOk() (*bool, bool)`

GetHasVariantsOk returns a tuple with the HasVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasVariants

`func (o *GetListingResponsePlatformListings) SetHasVariants(v bool)`

SetHasVariants sets HasVariants field to given value.


### GetCharityId

`func (o *GetListingResponsePlatformListings) GetCharityId() string`

GetCharityId returns the CharityId field if non-nil, zero value otherwise.

### GetCharityIdOk

`func (o *GetListingResponsePlatformListings) GetCharityIdOk() (*string, bool)`

GetCharityIdOk returns a tuple with the CharityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCharityId

`func (o *GetListingResponsePlatformListings) SetCharityId(v string)`

SetCharityId sets CharityId field to given value.

### HasCharityId

`func (o *GetListingResponsePlatformListings) HasCharityId() bool`

HasCharityId returns a boolean if a field has been set.

### SetCharityIdNil

`func (o *GetListingResponsePlatformListings) SetCharityIdNil(b bool)`

 SetCharityIdNil sets the value for CharityId to be an explicit nil

### UnsetCharityId
`func (o *GetListingResponsePlatformListings) UnsetCharityId()`

UnsetCharityId ensures that no value is present for CharityId, not even an explicit nil
### GetMinOrderQuantity

`func (o *GetListingResponsePlatformListings) GetMinOrderQuantity() float32`

GetMinOrderQuantity returns the MinOrderQuantity field if non-nil, zero value otherwise.

### GetMinOrderQuantityOk

`func (o *GetListingResponsePlatformListings) GetMinOrderQuantityOk() (*float32, bool)`

GetMinOrderQuantityOk returns a tuple with the MinOrderQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinOrderQuantity

`func (o *GetListingResponsePlatformListings) SetMinOrderQuantity(v float32)`

SetMinOrderQuantity sets MinOrderQuantity field to given value.


### GetWholesaleUnitPriceCents

`func (o *GetListingResponsePlatformListings) GetWholesaleUnitPriceCents() float32`

GetWholesaleUnitPriceCents returns the WholesaleUnitPriceCents field if non-nil, zero value otherwise.

### GetWholesaleUnitPriceCentsOk

`func (o *GetListingResponsePlatformListings) GetWholesaleUnitPriceCentsOk() (*float32, bool)`

GetWholesaleUnitPriceCentsOk returns a tuple with the WholesaleUnitPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWholesaleUnitPriceCents

`func (o *GetListingResponsePlatformListings) SetWholesaleUnitPriceCents(v float32)`

SetWholesaleUnitPriceCents sets WholesaleUnitPriceCents field to given value.

### HasWholesaleUnitPriceCents

`func (o *GetListingResponsePlatformListings) HasWholesaleUnitPriceCents() bool`

HasWholesaleUnitPriceCents returns a boolean if a field has been set.

### SetWholesaleUnitPriceCentsNil

`func (o *GetListingResponsePlatformListings) SetWholesaleUnitPriceCentsNil(b bool)`

 SetWholesaleUnitPriceCentsNil sets the value for WholesaleUnitPriceCents to be an explicit nil

### UnsetWholesaleUnitPriceCents
`func (o *GetListingResponsePlatformListings) UnsetWholesaleUnitPriceCents()`

UnsetWholesaleUnitPriceCents ensures that no value is present for WholesaleUnitPriceCents, not even an explicit nil
### GetWholesaleMinQuantity

`func (o *GetListingResponsePlatformListings) GetWholesaleMinQuantity() float32`

GetWholesaleMinQuantity returns the WholesaleMinQuantity field if non-nil, zero value otherwise.

### GetWholesaleMinQuantityOk

`func (o *GetListingResponsePlatformListings) GetWholesaleMinQuantityOk() (*float32, bool)`

GetWholesaleMinQuantityOk returns a tuple with the WholesaleMinQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWholesaleMinQuantity

`func (o *GetListingResponsePlatformListings) SetWholesaleMinQuantity(v float32)`

SetWholesaleMinQuantity sets WholesaleMinQuantity field to given value.

### HasWholesaleMinQuantity

`func (o *GetListingResponsePlatformListings) HasWholesaleMinQuantity() bool`

HasWholesaleMinQuantity returns a boolean if a field has been set.

### SetWholesaleMinQuantityNil

`func (o *GetListingResponsePlatformListings) SetWholesaleMinQuantityNil(b bool)`

 SetWholesaleMinQuantityNil sets the value for WholesaleMinQuantity to be an explicit nil

### UnsetWholesaleMinQuantity
`func (o *GetListingResponsePlatformListings) UnsetWholesaleMinQuantity()`

UnsetWholesaleMinQuantity ensures that no value is present for WholesaleMinQuantity, not even an explicit nil
### GetAuctionReserveCents

`func (o *GetListingResponsePlatformListings) GetAuctionReserveCents() float32`

GetAuctionReserveCents returns the AuctionReserveCents field if non-nil, zero value otherwise.

### GetAuctionReserveCentsOk

`func (o *GetListingResponsePlatformListings) GetAuctionReserveCentsOk() (*float32, bool)`

GetAuctionReserveCentsOk returns a tuple with the AuctionReserveCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuctionReserveCents

`func (o *GetListingResponsePlatformListings) SetAuctionReserveCents(v float32)`

SetAuctionReserveCents sets AuctionReserveCents field to given value.

### HasAuctionReserveCents

`func (o *GetListingResponsePlatformListings) HasAuctionReserveCents() bool`

HasAuctionReserveCents returns a boolean if a field has been set.

### SetAuctionReserveCentsNil

`func (o *GetListingResponsePlatformListings) SetAuctionReserveCentsNil(b bool)`

 SetAuctionReserveCentsNil sets the value for AuctionReserveCents to be an explicit nil

### UnsetAuctionReserveCents
`func (o *GetListingResponsePlatformListings) UnsetAuctionReserveCents()`

UnsetAuctionReserveCents ensures that no value is present for AuctionReserveCents, not even an explicit nil
### GetIsGtc

`func (o *GetListingResponsePlatformListings) GetIsGtc() bool`

GetIsGtc returns the IsGtc field if non-nil, zero value otherwise.

### GetIsGtcOk

`func (o *GetListingResponsePlatformListings) GetIsGtcOk() (*bool, bool)`

GetIsGtcOk returns a tuple with the IsGtc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsGtc

`func (o *GetListingResponsePlatformListings) SetIsGtc(v bool)`

SetIsGtc sets IsGtc field to given value.


### GetHandlingTimeDays

`func (o *GetListingResponsePlatformListings) GetHandlingTimeDays() float32`

GetHandlingTimeDays returns the HandlingTimeDays field if non-nil, zero value otherwise.

### GetHandlingTimeDaysOk

`func (o *GetListingResponsePlatformListings) GetHandlingTimeDaysOk() (*float32, bool)`

GetHandlingTimeDaysOk returns a tuple with the HandlingTimeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingTimeDays

`func (o *GetListingResponsePlatformListings) SetHandlingTimeDays(v float32)`

SetHandlingTimeDays sets HandlingTimeDays field to given value.


### GetItemLocation

`func (o *GetListingResponsePlatformListings) GetItemLocation() string`

GetItemLocation returns the ItemLocation field if non-nil, zero value otherwise.

### GetItemLocationOk

`func (o *GetListingResponsePlatformListings) GetItemLocationOk() (*string, bool)`

GetItemLocationOk returns a tuple with the ItemLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLocation

`func (o *GetListingResponsePlatformListings) SetItemLocation(v string)`

SetItemLocation sets ItemLocation field to given value.

### HasItemLocation

`func (o *GetListingResponsePlatformListings) HasItemLocation() bool`

HasItemLocation returns a boolean if a field has been set.

### SetItemLocationNil

`func (o *GetListingResponsePlatformListings) SetItemLocationNil(b bool)`

 SetItemLocationNil sets the value for ItemLocation to be an explicit nil

### UnsetItemLocation
`func (o *GetListingResponsePlatformListings) UnsetItemLocation()`

UnsetItemLocation ensures that no value is present for ItemLocation, not even an explicit nil
### GetItemLocationZip

`func (o *GetListingResponsePlatformListings) GetItemLocationZip() string`

GetItemLocationZip returns the ItemLocationZip field if non-nil, zero value otherwise.

### GetItemLocationZipOk

`func (o *GetListingResponsePlatformListings) GetItemLocationZipOk() (*string, bool)`

GetItemLocationZipOk returns a tuple with the ItemLocationZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLocationZip

`func (o *GetListingResponsePlatformListings) SetItemLocationZip(v string)`

SetItemLocationZip sets ItemLocationZip field to given value.

### HasItemLocationZip

`func (o *GetListingResponsePlatformListings) HasItemLocationZip() bool`

HasItemLocationZip returns a boolean if a field has been set.

### SetItemLocationZipNil

`func (o *GetListingResponsePlatformListings) SetItemLocationZipNil(b bool)`

 SetItemLocationZipNil sets the value for ItemLocationZip to be an explicit nil

### UnsetItemLocationZip
`func (o *GetListingResponsePlatformListings) UnsetItemLocationZip()`

UnsetItemLocationZip ensures that no value is present for ItemLocationZip, not even an explicit nil
### GetItemLocationCountry

`func (o *GetListingResponsePlatformListings) GetItemLocationCountry() string`

GetItemLocationCountry returns the ItemLocationCountry field if non-nil, zero value otherwise.

### GetItemLocationCountryOk

`func (o *GetListingResponsePlatformListings) GetItemLocationCountryOk() (*string, bool)`

GetItemLocationCountryOk returns a tuple with the ItemLocationCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLocationCountry

`func (o *GetListingResponsePlatformListings) SetItemLocationCountry(v string)`

SetItemLocationCountry sets ItemLocationCountry field to given value.

### HasItemLocationCountry

`func (o *GetListingResponsePlatformListings) HasItemLocationCountry() bool`

HasItemLocationCountry returns a boolean if a field has been set.

### SetItemLocationCountryNil

`func (o *GetListingResponsePlatformListings) SetItemLocationCountryNil(b bool)`

 SetItemLocationCountryNil sets the value for ItemLocationCountry to be an explicit nil

### UnsetItemLocationCountry
`func (o *GetListingResponsePlatformListings) UnsetItemLocationCountry()`

UnsetItemLocationCountry ensures that no value is present for ItemLocationCountry, not even an explicit nil
### GetDeclaredShippingCost

`func (o *GetListingResponsePlatformListings) GetDeclaredShippingCost() string`

GetDeclaredShippingCost returns the DeclaredShippingCost field if non-nil, zero value otherwise.

### GetDeclaredShippingCostOk

`func (o *GetListingResponsePlatformListings) GetDeclaredShippingCostOk() (*string, bool)`

GetDeclaredShippingCostOk returns a tuple with the DeclaredShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredShippingCost

`func (o *GetListingResponsePlatformListings) SetDeclaredShippingCost(v string)`

SetDeclaredShippingCost sets DeclaredShippingCost field to given value.

### HasDeclaredShippingCost

`func (o *GetListingResponsePlatformListings) HasDeclaredShippingCost() bool`

HasDeclaredShippingCost returns a boolean if a field has been set.

### SetDeclaredShippingCostNil

`func (o *GetListingResponsePlatformListings) SetDeclaredShippingCostNil(b bool)`

 SetDeclaredShippingCostNil sets the value for DeclaredShippingCost to be an explicit nil

### UnsetDeclaredShippingCost
`func (o *GetListingResponsePlatformListings) UnsetDeclaredShippingCost()`

UnsetDeclaredShippingCost ensures that no value is present for DeclaredShippingCost, not even an explicit nil
### GetReturnPolicyText

`func (o *GetListingResponsePlatformListings) GetReturnPolicyText() string`

GetReturnPolicyText returns the ReturnPolicyText field if non-nil, zero value otherwise.

### GetReturnPolicyTextOk

`func (o *GetListingResponsePlatformListings) GetReturnPolicyTextOk() (*string, bool)`

GetReturnPolicyTextOk returns a tuple with the ReturnPolicyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicyText

`func (o *GetListingResponsePlatformListings) SetReturnPolicyText(v string)`

SetReturnPolicyText sets ReturnPolicyText field to given value.

### HasReturnPolicyText

`func (o *GetListingResponsePlatformListings) HasReturnPolicyText() bool`

HasReturnPolicyText returns a boolean if a field has been set.

### SetReturnPolicyTextNil

`func (o *GetListingResponsePlatformListings) SetReturnPolicyTextNil(b bool)`

 SetReturnPolicyTextNil sets the value for ReturnPolicyText to be an explicit nil

### UnsetReturnPolicyText
`func (o *GetListingResponsePlatformListings) UnsetReturnPolicyText()`

UnsetReturnPolicyText ensures that no value is present for ReturnPolicyText, not even an explicit nil
### GetLastDriftCheckedAt

`func (o *GetListingResponsePlatformListings) GetLastDriftCheckedAt() time.Time`

GetLastDriftCheckedAt returns the LastDriftCheckedAt field if non-nil, zero value otherwise.

### GetLastDriftCheckedAtOk

`func (o *GetListingResponsePlatformListings) GetLastDriftCheckedAtOk() (*time.Time, bool)`

GetLastDriftCheckedAtOk returns a tuple with the LastDriftCheckedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDriftCheckedAt

`func (o *GetListingResponsePlatformListings) SetLastDriftCheckedAt(v time.Time)`

SetLastDriftCheckedAt sets LastDriftCheckedAt field to given value.

### HasLastDriftCheckedAt

`func (o *GetListingResponsePlatformListings) HasLastDriftCheckedAt() bool`

HasLastDriftCheckedAt returns a boolean if a field has been set.

### SetLastDriftCheckedAtNil

`func (o *GetListingResponsePlatformListings) SetLastDriftCheckedAtNil(b bool)`

 SetLastDriftCheckedAtNil sets the value for LastDriftCheckedAt to be an explicit nil

### UnsetLastDriftCheckedAt
`func (o *GetListingResponsePlatformListings) UnsetLastDriftCheckedAt()`

UnsetLastDriftCheckedAt ensures that no value is present for LastDriftCheckedAt, not even an explicit nil
### GetVerificationFailed

`func (o *GetListingResponsePlatformListings) GetVerificationFailed() bool`

GetVerificationFailed returns the VerificationFailed field if non-nil, zero value otherwise.

### GetVerificationFailedOk

`func (o *GetListingResponsePlatformListings) GetVerificationFailedOk() (*bool, bool)`

GetVerificationFailedOk returns a tuple with the VerificationFailed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationFailed

`func (o *GetListingResponsePlatformListings) SetVerificationFailed(v bool)`

SetVerificationFailed sets VerificationFailed field to given value.


### GetPublishedOverrides

`func (o *GetListingResponsePlatformListings) GetPublishedOverrides() map[string]interface{}`

GetPublishedOverrides returns the PublishedOverrides field if non-nil, zero value otherwise.

### GetPublishedOverridesOk

`func (o *GetListingResponsePlatformListings) GetPublishedOverridesOk() (*map[string]interface{}, bool)`

GetPublishedOverridesOk returns a tuple with the PublishedOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishedOverrides

`func (o *GetListingResponsePlatformListings) SetPublishedOverrides(v map[string]interface{})`

SetPublishedOverrides sets PublishedOverrides field to given value.

### HasPublishedOverrides

`func (o *GetListingResponsePlatformListings) HasPublishedOverrides() bool`

HasPublishedOverrides returns a boolean if a field has been set.

### SetPublishedOverridesNil

`func (o *GetListingResponsePlatformListings) SetPublishedOverridesNil(b bool)`

 SetPublishedOverridesNil sets the value for PublishedOverrides to be an explicit nil

### UnsetPublishedOverrides
`func (o *GetListingResponsePlatformListings) UnsetPublishedOverrides()`

UnsetPublishedOverrides ensures that no value is present for PublishedOverrides, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


