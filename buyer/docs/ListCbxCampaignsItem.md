# ListCbxCampaignsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **string** |  | 
**Name** | **string** |  | 
**Status** | **string** |  | 
**BudgetBaseUnits** | **string** |  | 
**WindowStart** | **time.Time** |  | 
**WindowEnd** | **time.Time** |  | 
**PreviewRecipientCount** | Pointer to **NullableFloat32** |  | [optional] 
**ExecutedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewListCbxCampaignsItem

`func NewListCbxCampaignsItem(campaignId string, name string, status string, budgetBaseUnits string, windowStart time.Time, windowEnd time.Time, ) *ListCbxCampaignsItem`

NewListCbxCampaignsItem instantiates a new ListCbxCampaignsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCbxCampaignsItemWithDefaults

`func NewListCbxCampaignsItemWithDefaults() *ListCbxCampaignsItem`

NewListCbxCampaignsItemWithDefaults instantiates a new ListCbxCampaignsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *ListCbxCampaignsItem) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *ListCbxCampaignsItem) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *ListCbxCampaignsItem) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### GetName

`func (o *ListCbxCampaignsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCbxCampaignsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCbxCampaignsItem) SetName(v string)`

SetName sets Name field to given value.


### GetStatus

`func (o *ListCbxCampaignsItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListCbxCampaignsItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListCbxCampaignsItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetBudgetBaseUnits

`func (o *ListCbxCampaignsItem) GetBudgetBaseUnits() string`

GetBudgetBaseUnits returns the BudgetBaseUnits field if non-nil, zero value otherwise.

### GetBudgetBaseUnitsOk

`func (o *ListCbxCampaignsItem) GetBudgetBaseUnitsOk() (*string, bool)`

GetBudgetBaseUnitsOk returns a tuple with the BudgetBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetBaseUnits

`func (o *ListCbxCampaignsItem) SetBudgetBaseUnits(v string)`

SetBudgetBaseUnits sets BudgetBaseUnits field to given value.


### GetWindowStart

`func (o *ListCbxCampaignsItem) GetWindowStart() time.Time`

GetWindowStart returns the WindowStart field if non-nil, zero value otherwise.

### GetWindowStartOk

`func (o *ListCbxCampaignsItem) GetWindowStartOk() (*time.Time, bool)`

GetWindowStartOk returns a tuple with the WindowStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowStart

`func (o *ListCbxCampaignsItem) SetWindowStart(v time.Time)`

SetWindowStart sets WindowStart field to given value.


### GetWindowEnd

`func (o *ListCbxCampaignsItem) GetWindowEnd() time.Time`

GetWindowEnd returns the WindowEnd field if non-nil, zero value otherwise.

### GetWindowEndOk

`func (o *ListCbxCampaignsItem) GetWindowEndOk() (*time.Time, bool)`

GetWindowEndOk returns a tuple with the WindowEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowEnd

`func (o *ListCbxCampaignsItem) SetWindowEnd(v time.Time)`

SetWindowEnd sets WindowEnd field to given value.


### GetPreviewRecipientCount

`func (o *ListCbxCampaignsItem) GetPreviewRecipientCount() float32`

GetPreviewRecipientCount returns the PreviewRecipientCount field if non-nil, zero value otherwise.

### GetPreviewRecipientCountOk

`func (o *ListCbxCampaignsItem) GetPreviewRecipientCountOk() (*float32, bool)`

GetPreviewRecipientCountOk returns a tuple with the PreviewRecipientCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreviewRecipientCount

`func (o *ListCbxCampaignsItem) SetPreviewRecipientCount(v float32)`

SetPreviewRecipientCount sets PreviewRecipientCount field to given value.

### HasPreviewRecipientCount

`func (o *ListCbxCampaignsItem) HasPreviewRecipientCount() bool`

HasPreviewRecipientCount returns a boolean if a field has been set.

### SetPreviewRecipientCountNil

`func (o *ListCbxCampaignsItem) SetPreviewRecipientCountNil(b bool)`

 SetPreviewRecipientCountNil sets the value for PreviewRecipientCount to be an explicit nil

### UnsetPreviewRecipientCount
`func (o *ListCbxCampaignsItem) UnsetPreviewRecipientCount()`

UnsetPreviewRecipientCount ensures that no value is present for PreviewRecipientCount, not even an explicit nil
### GetExecutedAt

`func (o *ListCbxCampaignsItem) GetExecutedAt() time.Time`

GetExecutedAt returns the ExecutedAt field if non-nil, zero value otherwise.

### GetExecutedAtOk

`func (o *ListCbxCampaignsItem) GetExecutedAtOk() (*time.Time, bool)`

GetExecutedAtOk returns a tuple with the ExecutedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutedAt

`func (o *ListCbxCampaignsItem) SetExecutedAt(v time.Time)`

SetExecutedAt sets ExecutedAt field to given value.

### HasExecutedAt

`func (o *ListCbxCampaignsItem) HasExecutedAt() bool`

HasExecutedAt returns a boolean if a field has been set.

### SetExecutedAtNil

`func (o *ListCbxCampaignsItem) SetExecutedAtNil(b bool)`

 SetExecutedAtNil sets the value for ExecutedAt to be an explicit nil

### UnsetExecutedAt
`func (o *ListCbxCampaignsItem) UnsetExecutedAt()`

UnsetExecutedAt ensures that no value is present for ExecutedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


