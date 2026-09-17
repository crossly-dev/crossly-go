# ListConnectionsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Platform** | **string** |  | 
**IsActive** | **bool** |  | 
**ShopUrl** | Pointer to **NullableString** |  | [optional] 
**LastSeenActive** | Pointer to **NullableTime** |  | [optional] 
**TokenExpiresAt** | Pointer to **NullableTime** |  | [optional] 
**CreatedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewListConnectionsItem

`func NewListConnectionsItem(id string, platform string, isActive bool, ) *ListConnectionsItem`

NewListConnectionsItem instantiates a new ListConnectionsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectionsItemWithDefaults

`func NewListConnectionsItemWithDefaults() *ListConnectionsItem`

NewListConnectionsItemWithDefaults instantiates a new ListConnectionsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListConnectionsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListConnectionsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListConnectionsItem) SetId(v string)`

SetId sets Id field to given value.


### GetPlatform

`func (o *ListConnectionsItem) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListConnectionsItem) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListConnectionsItem) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetIsActive

`func (o *ListConnectionsItem) GetIsActive() bool`

GetIsActive returns the IsActive field if non-nil, zero value otherwise.

### GetIsActiveOk

`func (o *ListConnectionsItem) GetIsActiveOk() (*bool, bool)`

GetIsActiveOk returns a tuple with the IsActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsActive

`func (o *ListConnectionsItem) SetIsActive(v bool)`

SetIsActive sets IsActive field to given value.


### GetShopUrl

`func (o *ListConnectionsItem) GetShopUrl() string`

GetShopUrl returns the ShopUrl field if non-nil, zero value otherwise.

### GetShopUrlOk

`func (o *ListConnectionsItem) GetShopUrlOk() (*string, bool)`

GetShopUrlOk returns a tuple with the ShopUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShopUrl

`func (o *ListConnectionsItem) SetShopUrl(v string)`

SetShopUrl sets ShopUrl field to given value.

### HasShopUrl

`func (o *ListConnectionsItem) HasShopUrl() bool`

HasShopUrl returns a boolean if a field has been set.

### SetShopUrlNil

`func (o *ListConnectionsItem) SetShopUrlNil(b bool)`

 SetShopUrlNil sets the value for ShopUrl to be an explicit nil

### UnsetShopUrl
`func (o *ListConnectionsItem) UnsetShopUrl()`

UnsetShopUrl ensures that no value is present for ShopUrl, not even an explicit nil
### GetLastSeenActive

`func (o *ListConnectionsItem) GetLastSeenActive() time.Time`

GetLastSeenActive returns the LastSeenActive field if non-nil, zero value otherwise.

### GetLastSeenActiveOk

`func (o *ListConnectionsItem) GetLastSeenActiveOk() (*time.Time, bool)`

GetLastSeenActiveOk returns a tuple with the LastSeenActive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSeenActive

`func (o *ListConnectionsItem) SetLastSeenActive(v time.Time)`

SetLastSeenActive sets LastSeenActive field to given value.

### HasLastSeenActive

`func (o *ListConnectionsItem) HasLastSeenActive() bool`

HasLastSeenActive returns a boolean if a field has been set.

### SetLastSeenActiveNil

`func (o *ListConnectionsItem) SetLastSeenActiveNil(b bool)`

 SetLastSeenActiveNil sets the value for LastSeenActive to be an explicit nil

### UnsetLastSeenActive
`func (o *ListConnectionsItem) UnsetLastSeenActive()`

UnsetLastSeenActive ensures that no value is present for LastSeenActive, not even an explicit nil
### GetTokenExpiresAt

`func (o *ListConnectionsItem) GetTokenExpiresAt() time.Time`

GetTokenExpiresAt returns the TokenExpiresAt field if non-nil, zero value otherwise.

### GetTokenExpiresAtOk

`func (o *ListConnectionsItem) GetTokenExpiresAtOk() (*time.Time, bool)`

GetTokenExpiresAtOk returns a tuple with the TokenExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenExpiresAt

`func (o *ListConnectionsItem) SetTokenExpiresAt(v time.Time)`

SetTokenExpiresAt sets TokenExpiresAt field to given value.

### HasTokenExpiresAt

`func (o *ListConnectionsItem) HasTokenExpiresAt() bool`

HasTokenExpiresAt returns a boolean if a field has been set.

### SetTokenExpiresAtNil

`func (o *ListConnectionsItem) SetTokenExpiresAtNil(b bool)`

 SetTokenExpiresAtNil sets the value for TokenExpiresAt to be an explicit nil

### UnsetTokenExpiresAt
`func (o *ListConnectionsItem) UnsetTokenExpiresAt()`

UnsetTokenExpiresAt ensures that no value is present for TokenExpiresAt, not even an explicit nil
### GetCreatedAt

`func (o *ListConnectionsItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListConnectionsItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListConnectionsItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListConnectionsItem) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### SetCreatedAtNil

`func (o *ListConnectionsItem) SetCreatedAtNil(b bool)`

 SetCreatedAtNil sets the value for CreatedAt to be an explicit nil

### UnsetCreatedAt
`func (o *ListConnectionsItem) UnsetCreatedAt()`

UnsetCreatedAt ensures that no value is present for CreatedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


