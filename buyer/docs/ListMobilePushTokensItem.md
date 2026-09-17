# ListMobilePushTokensItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** |  | 
**Platform** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**LastUsedAt** | **time.Time** |  | 

## Methods

### NewListMobilePushTokensItem

`func NewListMobilePushTokensItem(token string, platform string, createdAt time.Time, lastUsedAt time.Time, ) *ListMobilePushTokensItem`

NewListMobilePushTokensItem instantiates a new ListMobilePushTokensItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMobilePushTokensItemWithDefaults

`func NewListMobilePushTokensItemWithDefaults() *ListMobilePushTokensItem`

NewListMobilePushTokensItemWithDefaults instantiates a new ListMobilePushTokensItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ListMobilePushTokensItem) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListMobilePushTokensItem) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListMobilePushTokensItem) SetToken(v string)`

SetToken sets Token field to given value.


### GetPlatform

`func (o *ListMobilePushTokensItem) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListMobilePushTokensItem) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListMobilePushTokensItem) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetCreatedAt

`func (o *ListMobilePushTokensItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListMobilePushTokensItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListMobilePushTokensItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetLastUsedAt

`func (o *ListMobilePushTokensItem) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ListMobilePushTokensItem) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ListMobilePushTokensItem) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


