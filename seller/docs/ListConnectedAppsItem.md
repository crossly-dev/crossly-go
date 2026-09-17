# ListConnectedAppsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantId** | **string** |  | 
**Scopes** | **[]string** |  | 
**ConnectedAt** | **time.Time** |  | 
**LastUsedAt** | Pointer to **NullableTime** |  | [optional] 
**AppId** | **string** |  | 
**AppName** | **string** |  | 
**AppDescription** | Pointer to **NullableString** |  | [optional] 
**HomepageUrl** | Pointer to **NullableString** |  | [optional] 
**DeveloperName** | Pointer to **NullableString** |  | [optional] 
**DeveloperEmail** | **string** |  | 

## Methods

### NewListConnectedAppsItem

`func NewListConnectedAppsItem(grantId string, scopes []string, connectedAt time.Time, appId string, appName string, developerEmail string, ) *ListConnectedAppsItem`

NewListConnectedAppsItem instantiates a new ListConnectedAppsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListConnectedAppsItemWithDefaults

`func NewListConnectedAppsItemWithDefaults() *ListConnectedAppsItem`

NewListConnectedAppsItemWithDefaults instantiates a new ListConnectedAppsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantId

`func (o *ListConnectedAppsItem) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *ListConnectedAppsItem) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *ListConnectedAppsItem) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.


### GetScopes

`func (o *ListConnectedAppsItem) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *ListConnectedAppsItem) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *ListConnectedAppsItem) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetConnectedAt

`func (o *ListConnectedAppsItem) GetConnectedAt() time.Time`

GetConnectedAt returns the ConnectedAt field if non-nil, zero value otherwise.

### GetConnectedAtOk

`func (o *ListConnectedAppsItem) GetConnectedAtOk() (*time.Time, bool)`

GetConnectedAtOk returns a tuple with the ConnectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectedAt

`func (o *ListConnectedAppsItem) SetConnectedAt(v time.Time)`

SetConnectedAt sets ConnectedAt field to given value.


### GetLastUsedAt

`func (o *ListConnectedAppsItem) GetLastUsedAt() time.Time`

GetLastUsedAt returns the LastUsedAt field if non-nil, zero value otherwise.

### GetLastUsedAtOk

`func (o *ListConnectedAppsItem) GetLastUsedAtOk() (*time.Time, bool)`

GetLastUsedAtOk returns a tuple with the LastUsedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUsedAt

`func (o *ListConnectedAppsItem) SetLastUsedAt(v time.Time)`

SetLastUsedAt sets LastUsedAt field to given value.

### HasLastUsedAt

`func (o *ListConnectedAppsItem) HasLastUsedAt() bool`

HasLastUsedAt returns a boolean if a field has been set.

### SetLastUsedAtNil

`func (o *ListConnectedAppsItem) SetLastUsedAtNil(b bool)`

 SetLastUsedAtNil sets the value for LastUsedAt to be an explicit nil

### UnsetLastUsedAt
`func (o *ListConnectedAppsItem) UnsetLastUsedAt()`

UnsetLastUsedAt ensures that no value is present for LastUsedAt, not even an explicit nil
### GetAppId

`func (o *ListConnectedAppsItem) GetAppId() string`

GetAppId returns the AppId field if non-nil, zero value otherwise.

### GetAppIdOk

`func (o *ListConnectedAppsItem) GetAppIdOk() (*string, bool)`

GetAppIdOk returns a tuple with the AppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppId

`func (o *ListConnectedAppsItem) SetAppId(v string)`

SetAppId sets AppId field to given value.


### GetAppName

`func (o *ListConnectedAppsItem) GetAppName() string`

GetAppName returns the AppName field if non-nil, zero value otherwise.

### GetAppNameOk

`func (o *ListConnectedAppsItem) GetAppNameOk() (*string, bool)`

GetAppNameOk returns a tuple with the AppName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppName

`func (o *ListConnectedAppsItem) SetAppName(v string)`

SetAppName sets AppName field to given value.


### GetAppDescription

`func (o *ListConnectedAppsItem) GetAppDescription() string`

GetAppDescription returns the AppDescription field if non-nil, zero value otherwise.

### GetAppDescriptionOk

`func (o *ListConnectedAppsItem) GetAppDescriptionOk() (*string, bool)`

GetAppDescriptionOk returns a tuple with the AppDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDescription

`func (o *ListConnectedAppsItem) SetAppDescription(v string)`

SetAppDescription sets AppDescription field to given value.

### HasAppDescription

`func (o *ListConnectedAppsItem) HasAppDescription() bool`

HasAppDescription returns a boolean if a field has been set.

### SetAppDescriptionNil

`func (o *ListConnectedAppsItem) SetAppDescriptionNil(b bool)`

 SetAppDescriptionNil sets the value for AppDescription to be an explicit nil

### UnsetAppDescription
`func (o *ListConnectedAppsItem) UnsetAppDescription()`

UnsetAppDescription ensures that no value is present for AppDescription, not even an explicit nil
### GetHomepageUrl

`func (o *ListConnectedAppsItem) GetHomepageUrl() string`

GetHomepageUrl returns the HomepageUrl field if non-nil, zero value otherwise.

### GetHomepageUrlOk

`func (o *ListConnectedAppsItem) GetHomepageUrlOk() (*string, bool)`

GetHomepageUrlOk returns a tuple with the HomepageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHomepageUrl

`func (o *ListConnectedAppsItem) SetHomepageUrl(v string)`

SetHomepageUrl sets HomepageUrl field to given value.

### HasHomepageUrl

`func (o *ListConnectedAppsItem) HasHomepageUrl() bool`

HasHomepageUrl returns a boolean if a field has been set.

### SetHomepageUrlNil

`func (o *ListConnectedAppsItem) SetHomepageUrlNil(b bool)`

 SetHomepageUrlNil sets the value for HomepageUrl to be an explicit nil

### UnsetHomepageUrl
`func (o *ListConnectedAppsItem) UnsetHomepageUrl()`

UnsetHomepageUrl ensures that no value is present for HomepageUrl, not even an explicit nil
### GetDeveloperName

`func (o *ListConnectedAppsItem) GetDeveloperName() string`

GetDeveloperName returns the DeveloperName field if non-nil, zero value otherwise.

### GetDeveloperNameOk

`func (o *ListConnectedAppsItem) GetDeveloperNameOk() (*string, bool)`

GetDeveloperNameOk returns a tuple with the DeveloperName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperName

`func (o *ListConnectedAppsItem) SetDeveloperName(v string)`

SetDeveloperName sets DeveloperName field to given value.

### HasDeveloperName

`func (o *ListConnectedAppsItem) HasDeveloperName() bool`

HasDeveloperName returns a boolean if a field has been set.

### SetDeveloperNameNil

`func (o *ListConnectedAppsItem) SetDeveloperNameNil(b bool)`

 SetDeveloperNameNil sets the value for DeveloperName to be an explicit nil

### UnsetDeveloperName
`func (o *ListConnectedAppsItem) UnsetDeveloperName()`

UnsetDeveloperName ensures that no value is present for DeveloperName, not even an explicit nil
### GetDeveloperEmail

`func (o *ListConnectedAppsItem) GetDeveloperEmail() string`

GetDeveloperEmail returns the DeveloperEmail field if non-nil, zero value otherwise.

### GetDeveloperEmailOk

`func (o *ListConnectedAppsItem) GetDeveloperEmailOk() (*string, bool)`

GetDeveloperEmailOk returns a tuple with the DeveloperEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeveloperEmail

`func (o *ListConnectedAppsItem) SetDeveloperEmail(v string)`

SetDeveloperEmail sets DeveloperEmail field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


