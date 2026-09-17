# GetConnectionHealthResponseAccounts

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformName** | **string** |  | 
**AccountId** | Pointer to **NullableString** |  | [optional] 
**AccountSlot** | Pointer to **NullableFloat32** |  | [optional] 
**Label** | Pointer to **NullableString** |  | [optional] 
**PlatformUsername** | Pointer to **NullableString** |  | [optional] 
**Summary** | **string** | Plain-English \&quot;what is true\&quot; + \&quot;what to do\&quot;. Never empty. | 
**Action** | **string** |  | 
**Liveness** | [**GetConnectionHealthResponseLiveness**](GetConnectionHealthResponseLiveness.md) |  | 
**Browser** | Pointer to [**NullableGetConnectionHealthResponseBrowser**](GetConnectionHealthResponseBrowser.md) |  | [optional] 
**Platform** | **string** |  | 
**State** | **string** |  | 
**Severity** | **string** |  | 
**Audience** | **string** |  | 
**Anchors** | Pointer to [**NullableGetConnectionHealthResponseAnchors**](GetConnectionHealthResponseAnchors.md) |  | [optional] 
**Notes** | **[]string** | Secondary observations that do not change the verdict but change the debugging. Always safe to show; never the only thing shown. | 

## Methods

### NewGetConnectionHealthResponseAccounts

`func NewGetConnectionHealthResponseAccounts(platformName string, summary string, action string, liveness GetConnectionHealthResponseLiveness, platform string, state string, severity string, audience string, notes []string, ) *GetConnectionHealthResponseAccounts`

NewGetConnectionHealthResponseAccounts instantiates a new GetConnectionHealthResponseAccounts object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionHealthResponseAccountsWithDefaults

`func NewGetConnectionHealthResponseAccountsWithDefaults() *GetConnectionHealthResponseAccounts`

NewGetConnectionHealthResponseAccountsWithDefaults instantiates a new GetConnectionHealthResponseAccounts object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformName

`func (o *GetConnectionHealthResponseAccounts) GetPlatformName() string`

GetPlatformName returns the PlatformName field if non-nil, zero value otherwise.

### GetPlatformNameOk

`func (o *GetConnectionHealthResponseAccounts) GetPlatformNameOk() (*string, bool)`

GetPlatformNameOk returns a tuple with the PlatformName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformName

`func (o *GetConnectionHealthResponseAccounts) SetPlatformName(v string)`

SetPlatformName sets PlatformName field to given value.


### GetAccountId

`func (o *GetConnectionHealthResponseAccounts) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetConnectionHealthResponseAccounts) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetConnectionHealthResponseAccounts) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetConnectionHealthResponseAccounts) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *GetConnectionHealthResponseAccounts) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *GetConnectionHealthResponseAccounts) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetAccountSlot

`func (o *GetConnectionHealthResponseAccounts) GetAccountSlot() float32`

GetAccountSlot returns the AccountSlot field if non-nil, zero value otherwise.

### GetAccountSlotOk

`func (o *GetConnectionHealthResponseAccounts) GetAccountSlotOk() (*float32, bool)`

GetAccountSlotOk returns a tuple with the AccountSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSlot

`func (o *GetConnectionHealthResponseAccounts) SetAccountSlot(v float32)`

SetAccountSlot sets AccountSlot field to given value.

### HasAccountSlot

`func (o *GetConnectionHealthResponseAccounts) HasAccountSlot() bool`

HasAccountSlot returns a boolean if a field has been set.

### SetAccountSlotNil

`func (o *GetConnectionHealthResponseAccounts) SetAccountSlotNil(b bool)`

 SetAccountSlotNil sets the value for AccountSlot to be an explicit nil

### UnsetAccountSlot
`func (o *GetConnectionHealthResponseAccounts) UnsetAccountSlot()`

UnsetAccountSlot ensures that no value is present for AccountSlot, not even an explicit nil
### GetLabel

`func (o *GetConnectionHealthResponseAccounts) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetConnectionHealthResponseAccounts) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetConnectionHealthResponseAccounts) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *GetConnectionHealthResponseAccounts) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### SetLabelNil

`func (o *GetConnectionHealthResponseAccounts) SetLabelNil(b bool)`

 SetLabelNil sets the value for Label to be an explicit nil

### UnsetLabel
`func (o *GetConnectionHealthResponseAccounts) UnsetLabel()`

UnsetLabel ensures that no value is present for Label, not even an explicit nil
### GetPlatformUsername

`func (o *GetConnectionHealthResponseAccounts) GetPlatformUsername() string`

GetPlatformUsername returns the PlatformUsername field if non-nil, zero value otherwise.

### GetPlatformUsernameOk

`func (o *GetConnectionHealthResponseAccounts) GetPlatformUsernameOk() (*string, bool)`

GetPlatformUsernameOk returns a tuple with the PlatformUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformUsername

`func (o *GetConnectionHealthResponseAccounts) SetPlatformUsername(v string)`

SetPlatformUsername sets PlatformUsername field to given value.

### HasPlatformUsername

`func (o *GetConnectionHealthResponseAccounts) HasPlatformUsername() bool`

HasPlatformUsername returns a boolean if a field has been set.

### SetPlatformUsernameNil

`func (o *GetConnectionHealthResponseAccounts) SetPlatformUsernameNil(b bool)`

 SetPlatformUsernameNil sets the value for PlatformUsername to be an explicit nil

### UnsetPlatformUsername
`func (o *GetConnectionHealthResponseAccounts) UnsetPlatformUsername()`

UnsetPlatformUsername ensures that no value is present for PlatformUsername, not even an explicit nil
### GetSummary

`func (o *GetConnectionHealthResponseAccounts) GetSummary() string`

GetSummary returns the Summary field if non-nil, zero value otherwise.

### GetSummaryOk

`func (o *GetConnectionHealthResponseAccounts) GetSummaryOk() (*string, bool)`

GetSummaryOk returns a tuple with the Summary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSummary

`func (o *GetConnectionHealthResponseAccounts) SetSummary(v string)`

SetSummary sets Summary field to given value.


### GetAction

`func (o *GetConnectionHealthResponseAccounts) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetConnectionHealthResponseAccounts) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetConnectionHealthResponseAccounts) SetAction(v string)`

SetAction sets Action field to given value.


### GetLiveness

`func (o *GetConnectionHealthResponseAccounts) GetLiveness() GetConnectionHealthResponseLiveness`

GetLiveness returns the Liveness field if non-nil, zero value otherwise.

### GetLivenessOk

`func (o *GetConnectionHealthResponseAccounts) GetLivenessOk() (*GetConnectionHealthResponseLiveness, bool)`

GetLivenessOk returns a tuple with the Liveness field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLiveness

`func (o *GetConnectionHealthResponseAccounts) SetLiveness(v GetConnectionHealthResponseLiveness)`

SetLiveness sets Liveness field to given value.


### GetBrowser

`func (o *GetConnectionHealthResponseAccounts) GetBrowser() GetConnectionHealthResponseBrowser`

GetBrowser returns the Browser field if non-nil, zero value otherwise.

### GetBrowserOk

`func (o *GetConnectionHealthResponseAccounts) GetBrowserOk() (*GetConnectionHealthResponseBrowser, bool)`

GetBrowserOk returns a tuple with the Browser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrowser

`func (o *GetConnectionHealthResponseAccounts) SetBrowser(v GetConnectionHealthResponseBrowser)`

SetBrowser sets Browser field to given value.

### HasBrowser

`func (o *GetConnectionHealthResponseAccounts) HasBrowser() bool`

HasBrowser returns a boolean if a field has been set.

### SetBrowserNil

`func (o *GetConnectionHealthResponseAccounts) SetBrowserNil(b bool)`

 SetBrowserNil sets the value for Browser to be an explicit nil

### UnsetBrowser
`func (o *GetConnectionHealthResponseAccounts) UnsetBrowser()`

UnsetBrowser ensures that no value is present for Browser, not even an explicit nil
### GetPlatform

`func (o *GetConnectionHealthResponseAccounts) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetConnectionHealthResponseAccounts) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetConnectionHealthResponseAccounts) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetState

`func (o *GetConnectionHealthResponseAccounts) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *GetConnectionHealthResponseAccounts) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *GetConnectionHealthResponseAccounts) SetState(v string)`

SetState sets State field to given value.


### GetSeverity

`func (o *GetConnectionHealthResponseAccounts) GetSeverity() string`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetConnectionHealthResponseAccounts) GetSeverityOk() (*string, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetConnectionHealthResponseAccounts) SetSeverity(v string)`

SetSeverity sets Severity field to given value.


### GetAudience

`func (o *GetConnectionHealthResponseAccounts) GetAudience() string`

GetAudience returns the Audience field if non-nil, zero value otherwise.

### GetAudienceOk

`func (o *GetConnectionHealthResponseAccounts) GetAudienceOk() (*string, bool)`

GetAudienceOk returns a tuple with the Audience field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAudience

`func (o *GetConnectionHealthResponseAccounts) SetAudience(v string)`

SetAudience sets Audience field to given value.


### GetAnchors

`func (o *GetConnectionHealthResponseAccounts) GetAnchors() GetConnectionHealthResponseAnchors`

GetAnchors returns the Anchors field if non-nil, zero value otherwise.

### GetAnchorsOk

`func (o *GetConnectionHealthResponseAccounts) GetAnchorsOk() (*GetConnectionHealthResponseAnchors, bool)`

GetAnchorsOk returns a tuple with the Anchors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnchors

`func (o *GetConnectionHealthResponseAccounts) SetAnchors(v GetConnectionHealthResponseAnchors)`

SetAnchors sets Anchors field to given value.

### HasAnchors

`func (o *GetConnectionHealthResponseAccounts) HasAnchors() bool`

HasAnchors returns a boolean if a field has been set.

### SetAnchorsNil

`func (o *GetConnectionHealthResponseAccounts) SetAnchorsNil(b bool)`

 SetAnchorsNil sets the value for Anchors to be an explicit nil

### UnsetAnchors
`func (o *GetConnectionHealthResponseAccounts) UnsetAnchors()`

UnsetAnchors ensures that no value is present for Anchors, not even an explicit nil
### GetNotes

`func (o *GetConnectionHealthResponseAccounts) GetNotes() []string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetConnectionHealthResponseAccounts) GetNotesOk() (*[]string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetConnectionHealthResponseAccounts) SetNotes(v []string)`

SetNotes sets Notes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


