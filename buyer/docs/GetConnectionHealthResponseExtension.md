# GetConnectionHealthResponseExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Online** | **bool** | Heartbeat within the last 15 minutes. False is NOT proof of breakage — a closed browser looks the same — so it is always paired with the durable timestamp below. | 
**LastBrowserPushAt** | Pointer to **NullableString** | Newest browser push across every account. The durable answer to \&quot;when did the extension last do anything\&quot;, surviving a Redis flush. | [optional] 
**LastBrowserPushAgo** | **string** |  | 
**ReportedVersion** | Pointer to **NullableString** | Highest version any platform&#39;s browser report carried, or null when the extension has never told us (it does not send one today — see the service README notes in routes/extension-health.ts). | [optional] 

## Methods

### NewGetConnectionHealthResponseExtension

`func NewGetConnectionHealthResponseExtension(online bool, lastBrowserPushAgo string, ) *GetConnectionHealthResponseExtension`

NewGetConnectionHealthResponseExtension instantiates a new GetConnectionHealthResponseExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionHealthResponseExtensionWithDefaults

`func NewGetConnectionHealthResponseExtensionWithDefaults() *GetConnectionHealthResponseExtension`

NewGetConnectionHealthResponseExtensionWithDefaults instantiates a new GetConnectionHealthResponseExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOnline

`func (o *GetConnectionHealthResponseExtension) GetOnline() bool`

GetOnline returns the Online field if non-nil, zero value otherwise.

### GetOnlineOk

`func (o *GetConnectionHealthResponseExtension) GetOnlineOk() (*bool, bool)`

GetOnlineOk returns a tuple with the Online field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnline

`func (o *GetConnectionHealthResponseExtension) SetOnline(v bool)`

SetOnline sets Online field to given value.


### GetLastBrowserPushAt

`func (o *GetConnectionHealthResponseExtension) GetLastBrowserPushAt() string`

GetLastBrowserPushAt returns the LastBrowserPushAt field if non-nil, zero value otherwise.

### GetLastBrowserPushAtOk

`func (o *GetConnectionHealthResponseExtension) GetLastBrowserPushAtOk() (*string, bool)`

GetLastBrowserPushAtOk returns a tuple with the LastBrowserPushAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBrowserPushAt

`func (o *GetConnectionHealthResponseExtension) SetLastBrowserPushAt(v string)`

SetLastBrowserPushAt sets LastBrowserPushAt field to given value.

### HasLastBrowserPushAt

`func (o *GetConnectionHealthResponseExtension) HasLastBrowserPushAt() bool`

HasLastBrowserPushAt returns a boolean if a field has been set.

### SetLastBrowserPushAtNil

`func (o *GetConnectionHealthResponseExtension) SetLastBrowserPushAtNil(b bool)`

 SetLastBrowserPushAtNil sets the value for LastBrowserPushAt to be an explicit nil

### UnsetLastBrowserPushAt
`func (o *GetConnectionHealthResponseExtension) UnsetLastBrowserPushAt()`

UnsetLastBrowserPushAt ensures that no value is present for LastBrowserPushAt, not even an explicit nil
### GetLastBrowserPushAgo

`func (o *GetConnectionHealthResponseExtension) GetLastBrowserPushAgo() string`

GetLastBrowserPushAgo returns the LastBrowserPushAgo field if non-nil, zero value otherwise.

### GetLastBrowserPushAgoOk

`func (o *GetConnectionHealthResponseExtension) GetLastBrowserPushAgoOk() (*string, bool)`

GetLastBrowserPushAgoOk returns a tuple with the LastBrowserPushAgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBrowserPushAgo

`func (o *GetConnectionHealthResponseExtension) SetLastBrowserPushAgo(v string)`

SetLastBrowserPushAgo sets LastBrowserPushAgo field to given value.


### GetReportedVersion

`func (o *GetConnectionHealthResponseExtension) GetReportedVersion() string`

GetReportedVersion returns the ReportedVersion field if non-nil, zero value otherwise.

### GetReportedVersionOk

`func (o *GetConnectionHealthResponseExtension) GetReportedVersionOk() (*string, bool)`

GetReportedVersionOk returns a tuple with the ReportedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportedVersion

`func (o *GetConnectionHealthResponseExtension) SetReportedVersion(v string)`

SetReportedVersion sets ReportedVersion field to given value.

### HasReportedVersion

`func (o *GetConnectionHealthResponseExtension) HasReportedVersion() bool`

HasReportedVersion returns a boolean if a field has been set.

### SetReportedVersionNil

`func (o *GetConnectionHealthResponseExtension) SetReportedVersionNil(b bool)`

 SetReportedVersionNil sets the value for ReportedVersion to be an explicit nil

### UnsetReportedVersion
`func (o *GetConnectionHealthResponseExtension) UnsetReportedVersion()`

UnsetReportedVersion ensures that no value is present for ReportedVersion, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


