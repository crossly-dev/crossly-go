# GetConnectionHealthResponseLiveness

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastBrowserPushAt** | Pointer to **NullableString** | The clean \&quot;the browser pushed cookies\&quot; signal. | [optional] 
**LastBrowserPushAgo** | **string** |  | 
**LastSyncedAt** | Pointer to **NullableString** | Also stamped by the executor on any successful server-side call, so it is NOT evidence the extension is alive. Exposed for debugging only. | [optional] 
**LastUsedAt** | Pointer to **NullableString** | Last server-exec attempt, success or fail. | [optional] 
**LastUsedAgo** | **string** |  | 
**AuthFailureStreak** | **float32** |  | 

## Methods

### NewGetConnectionHealthResponseLiveness

`func NewGetConnectionHealthResponseLiveness(lastBrowserPushAgo string, lastUsedAgo string, authFailureStreak float32, ) *GetConnectionHealthResponseLiveness`

NewGetConnectionHealthResponseLiveness instantiates a new GetConnectionHealthResponseLiveness object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionHealthResponseLivenessWithDefaults

`func NewGetConnectionHealthResponseLivenessWithDefaults() *GetConnectionHealthResponseLiveness`

NewGetConnectionHealthResponseLivenessWithDefaults instantiates a new GetConnectionHealthResponseLiveness object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastBrowserPushAt

`func (o *GetConnectionHealthResponseLiveness) GetLastBrowserPushAt() string`

GetLastBrowserPushAt returns the LastBrowserPushAt field if non-nil, zero value otherwise.

### GetLastBrowserPushAtOk

`func (o *GetConnectionHealthResponseLiveness) GetLastBrowserPushAtOk() (*string, bool)`

GetLastBrowserPushAtOk returns a tuple with the LastBrowserPushAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBrowserPushAt

`func (o *GetConnectionHealthResponseLiveness) SetLastBrowserPushAt(v string)`

SetLastBrowserPushAt sets LastBrowserPushAt field to given value.

### HasLastBrowserPushAt

`func (o *GetConnectionHealthResponseLiveness) HasLastBrowserPushAt() bool`

HasLastBrowserPushAt returns a boolean if a field has been set.

### SetLastBrowserPushAtNil

`func (o *GetConnectionHealthResponseLiveness) SetLastBrowserPushAtNil(b bool)`

 SetLastBrowserPushAtNil sets the value for LastBrowserPushAt to be an explicit nil

### UnsetLastBrowserPushAt
`func (o *GetConnectionHealthResponseLiveness) UnsetLastBrowserPushAt()`

UnsetLastBrowserPushAt ensures that no value is present for LastBrowserPushAt, not even an explicit nil
### GetLastBrowserPushAgo

`func (o *GetConnectionHealthResponseLiveness) GetLastBrowserPushAgo() string`

GetLastBrowserPushAgo returns the LastBrowserPushAgo field if non-nil, zero value otherwise.

### GetLastBrowserPushAgoOk

`func (o *GetConnectionHealthResponseLiveness) GetLastBrowserPushAgoOk() (*string, bool)`

GetLastBrowserPushAgoOk returns a tuple with the LastBrowserPushAgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBrowserPushAgo

`func (o *GetConnectionHealthResponseLiveness) SetLastBrowserPushAgo(v string)`

SetLastBrowserPushAgo sets LastBrowserPushAgo field to given value.


### GetLastSyncedAt

`func (o *GetConnectionHealthResponseLiveness) GetLastSyncedAt() string`

GetLastSyncedAt returns the LastSyncedAt field if non-nil, zero value otherwise.

### GetLastSyncedAtOk

`func (o *GetConnectionHealthResponseLiveness) GetLastSyncedAtOk() (*string, bool)`

GetLastSyncedAtOk returns a tuple with the LastSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncedAt

`func (o *GetConnectionHealthResponseLiveness) SetLastSyncedAt(v string)`

SetLastSyncedAt sets LastSyncedAt field to given value.

### HasLastSyncedAt

`func (o *GetConnectionHealthResponseLiveness) HasLastSyncedAt() bool`

HasLastSyncedAt returns a boolean if a field has been set.

### SetLastSyncedAtNil

`func (o *GetConnectionHealthResponseLiveness) SetLastSyncedAtNil(b bool)`

 SetLastSyncedAtNil sets the value for LastSyncedAt to be an explicit nil

### UnsetLastSyncedAt
`func (o *GetConnectionHealthResponseLiveness) UnsetLastSyncedAt()`

UnsetLastSyncedAt ensures that no value is present for LastSyncedAt, not even an explicit nil
### GetLastUsedAt

`func (o *GetConnectionHealthResponseLiveness) GetLastUsedAt() string`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *GetConnectionHealthResponseLiveness) GetLastUsedAtOk() (*string, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *GetConnectionHealthResponseLiveness) SetLastUsedAt(v string)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *GetConnectionHealthResponseLiveness) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### SetLastUsedAtNil

`func (o *GetConnectionHealthResponseLiveness) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *GetConnectionHealthResponseLiveness) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetLastUsedAgo

`func (o *GetConnectionHealthResponseLiveness) GetLastUsedAgo() string`

GetLastUsedAgo returns the LastUsedAgo field if non-nil, zero value otherwise.

### GetLastUsedAgoOk

`func (o *GetConnectionHealthResponseLiveness) GetLastUsedAgoOk() (*string, bool)`

GetLastUsedAgoOk returns a tuple with the LastUsedAgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAgo

`func (o *GetConnectionHealthResponseLiveness) SetLastUsedAgo(v string)`

SetLastUsedAgo sets LastUsedAgo field to given value.


### GetAuthFailureStreak

`func (o *GetConnectionHealthResponseLiveness) GetAuthFailureStreak() float32`

GetAuthFailureStreak returns the AuthFailureStreak field if non-nil, zero value otherwise.

### GetAuthFailureStreakOk

`func (o *GetConnectionHealthResponseLiveness) GetAuthFailureStreakOk() (*float32, bool)`

GetAuthFailureStreakOk returns a tuple with the AuthFailureStreak field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthFailureStreak

`func (o *GetConnectionHealthResponseLiveness) SetAuthFailureStreak(v float32)`

SetAuthFailureStreak sets AuthFailureStreak field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


