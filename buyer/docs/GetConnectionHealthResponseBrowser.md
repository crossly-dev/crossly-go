# GetConnectionHealthResponseBrowser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Observations** | **[]map[string]interface{}** |  | 
**CookieCount** | **float32** |  | 
**ExtensionVersion** | Pointer to **NullableString** |  | [optional] 
**ObservedAt** | **string** |  | 
**ObservedAgo** | **string** |  | 

## Methods

### NewGetConnectionHealthResponseBrowser

`func NewGetConnectionHealthResponseBrowser(observations []map[string]interface{}, cookieCount float32, observedAt string, observedAgo string, ) *GetConnectionHealthResponseBrowser`

NewGetConnectionHealthResponseBrowser instantiates a new GetConnectionHealthResponseBrowser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionHealthResponseBrowserWithDefaults

`func NewGetConnectionHealthResponseBrowserWithDefaults() *GetConnectionHealthResponseBrowser`

NewGetConnectionHealthResponseBrowserWithDefaults instantiates a new GetConnectionHealthResponseBrowser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObservations

`func (o *GetConnectionHealthResponseBrowser) GetObservations() []map[string]interface{}`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *GetConnectionHealthResponseBrowser) GetObservationsOk() (*[]map[string]interface{}, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *GetConnectionHealthResponseBrowser) SetObservations(v []map[string]interface{})`

SetObservations sets Observations field to given value.


### GetCookieCount

`func (o *GetConnectionHealthResponseBrowser) GetCookieCount() float32`

GetCookieCount returns the CookieCount field if non-nil, zero value otherwise.

### GetCookieCountOk

`func (o *GetConnectionHealthResponseBrowser) GetCookieCountOk() (*float32, bool)`

GetCookieCountOk returns a tuple with the CookieCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieCount

`func (o *GetConnectionHealthResponseBrowser) SetCookieCount(v float32)`

SetCookieCount sets CookieCount field to given value.


### GetExtensionVersion

`func (o *GetConnectionHealthResponseBrowser) GetExtensionVersion() string`

GetExtensionVersion returns the ExtensionVersion field if non-nil, zero value otherwise.

### GetExtensionVersionOk

`func (o *GetConnectionHealthResponseBrowser) GetExtensionVersionOk() (*string, bool)`

GetExtensionVersionOk returns a tuple with the ExtensionVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensionVersion

`func (o *GetConnectionHealthResponseBrowser) SetExtensionVersion(v string)`

SetExtensionVersion sets ExtensionVersion field to given value.

### HasExtensionVersion

`func (o *GetConnectionHealthResponseBrowser) HasExtensionVersion() bool`

HasExtensionVersion returns a boolean if a field has been set.

### SetExtensionVersionNil

`func (o *GetConnectionHealthResponseBrowser) SetExtensionVersionNil(b bool)`

 SetExtensionVersionNil sets the value for ExtensionVersion to be an explicit nil

### UnsetExtensionVersion
`func (o *GetConnectionHealthResponseBrowser) UnsetExtensionVersion()`

UnsetExtensionVersion ensures that no value is present for ExtensionVersion, not even an explicit nil
### GetObservedAt

`func (o *GetConnectionHealthResponseBrowser) GetObservedAt() string`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *GetConnectionHealthResponseBrowser) GetObservedAtOk() (*string, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *GetConnectionHealthResponseBrowser) SetObservedAt(v string)`

SetObservedAt sets ObservedAt field to given value.


### GetObservedAgo

`func (o *GetConnectionHealthResponseBrowser) GetObservedAgo() string`

GetObservedAgo returns the ObservedAgo field if non-nil, zero value otherwise.

### GetObservedAgoOk

`func (o *GetConnectionHealthResponseBrowser) GetObservedAgoOk() (*string, bool)`

GetObservedAgoOk returns a tuple with the ObservedAgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAgo

`func (o *GetConnectionHealthResponseBrowser) SetObservedAgo(v string)`

SetObservedAgo sets ObservedAgo field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


