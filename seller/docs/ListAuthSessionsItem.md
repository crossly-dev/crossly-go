# ListAuthSessionsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsCurrent** | **bool** |  | 
**Id** | **string** |  | 
**UserAgent** | Pointer to **NullableString** |  | [optional] 
**IpAddress** | Pointer to **NullableString** |  | [optional] 
**IpCountry** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**LastUsedAt** | **time.Time** |  | 
**ExpiresAt** | **time.Time** |  | 

## Methods

### NewListAuthSessionsItem

`func NewListAuthSessionsItem(isCurrent bool, id string, createdAt time.Time, lastUsedAt time.Time, expiresAt time.Time, ) *ListAuthSessionsItem`

NewListAuthSessionsItem instantiates a new ListAuthSessionsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListAuthSessionsItemWithDefaults

`func NewListAuthSessionsItemWithDefaults() *ListAuthSessionsItem`

NewListAuthSessionsItemWithDefaults instantiates a new ListAuthSessionsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsCurrent

`func (o *ListAuthSessionsItem) GetIsCurrent() bool`

GetIsCurrent returns the IsCurrent field if non-nil, zero value otherwise.

### GetIsCurrentOk

`func (o *ListAuthSessionsItem) GetIsCurrentOk() (*bool, bool)`

GetIsCurrentOk returns a tuple with the IsCurrent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCurrent

`func (o *ListAuthSessionsItem) SetIsCurrent(v bool)`

SetIsCurrent sets IsCurrent field to given value.


### GetId

`func (o *ListAuthSessionsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListAuthSessionsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListAuthSessionsItem) SetId(v string)`

SetId sets Id field to given value.


### GetUserAgent

`func (o *ListAuthSessionsItem) GetUserAgent() string`

GetUserAgent returns the UserAgent field if non-nil, zero value otherwise.

### GetUserAgentOk

`func (o *ListAuthSessionsItem) GetUserAgentOk() (*string, bool)`

GetUserAgentOk returns a tuple with the UserAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserAgent

`func (o *ListAuthSessionsItem) SetUserAgent(v string)`

SetUserAgent sets UserAgent field to given value.

### HasUserAgent

`func (o *ListAuthSessionsItem) HasUserAgent() bool`

HasUserAgent returns a boolean if a field has been set.

### SetUserAgentNil

`func (o *ListAuthSessionsItem) SetUserAgentNil(b bool)`

 SetUserAgentNil sets the value for UserAgent to be an explicit nil

### UnsetUserAgent
`func (o *ListAuthSessionsItem) UnsetUserAgent()`

UnsetUserAgent ensures that no value is present for UserAgent, not even an explicit nil
### GetIpAddress

`func (o *ListAuthSessionsItem) GetIpAddress() string`

GetIpAddress returns the IpAddress field if non-nil, zero value otherwise.

### GetIpAddressOk

`func (o *ListAuthSessionsItem) GetIpAddressOk() (*string, bool)`

GetIpAddressOk returns a tuple with the IpAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpAddress

`func (o *ListAuthSessionsItem) SetIpAddress(v string)`

SetIpAddress sets IpAddress field to given value.

### HasIpAddress

`func (o *ListAuthSessionsItem) HasIpAddress() bool`

HasIpAddress returns a boolean if a field has been set.

### SetIpAddressNil

`func (o *ListAuthSessionsItem) SetIpAddressNil(b bool)`

 SetIpAddressNil sets the value for IpAddress to be an explicit nil

### UnsetIpAddress
`func (o *ListAuthSessionsItem) UnsetIpAddress()`

UnsetIpAddress ensures that no value is present for IpAddress, not even an explicit nil
### GetIpCountry

`func (o *ListAuthSessionsItem) GetIpCountry() string`

GetIpCountry returns the IpCountry field if non-nil, zero value otherwise.

### GetIpCountryOk

`func (o *ListAuthSessionsItem) GetIpCountryOk() (*string, bool)`

GetIpCountryOk returns a tuple with the IpCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpCountry

`func (o *ListAuthSessionsItem) SetIpCountry(v string)`

SetIpCountry sets IpCountry field to given value.

### HasIpCountry

`func (o *ListAuthSessionsItem) HasIpCountry() bool`

HasIpCountry returns a boolean if a field has been set.

### SetIpCountryNil

`func (o *ListAuthSessionsItem) SetIpCountryNil(b bool)`

 SetIpCountryNil sets the value for IpCountry to be an explicit nil

### UnsetIpCountry
`func (o *ListAuthSessionsItem) UnsetIpCountry()`

UnsetIpCountry ensures that no value is present for IpCountry, not even an explicit nil
### GetCreatedAt

`func (o *ListAuthSessionsItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListAuthSessionsItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListAuthSessionsItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastUsedAt

`func (o *ListAuthSessionsItem) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ListAuthSessionsItem) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ListAuthSessionsItem) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.


### GetExpiresAt

`func (o *ListAuthSessionsItem) GetExpiresAt() time.Time`

GetExpiresAt returns the ExpiresAt field if non-nil, zero value otherwise.

### GetExpiresAtOk

`func (o *ListAuthSessionsItem) GetExpiresAtOk() (*time.Time, bool)`

GetExpiresAtOk returns a tuple with the ExpiresAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiresAt

`func (o *ListAuthSessionsItem) SetExpiresAt(v time.Time)`

SetExpiresAt sets ExpiresAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


