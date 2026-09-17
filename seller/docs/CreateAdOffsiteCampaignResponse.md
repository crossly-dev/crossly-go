# CreateAdOffsiteCampaignResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignId** | **string** |  | 
**Network** | **string** |  | 
**Status** | **string** |  | 
**CreativeHeadline** | Pointer to **NullableString** |  | [optional] 
**DailyBudgetCents** | **float32** |  | 

## Methods

### NewCreateAdOffsiteCampaignResponse

`func NewCreateAdOffsiteCampaignResponse(campaignId string, network string, status string, dailyBudgetCents float32, ) *CreateAdOffsiteCampaignResponse`

NewCreateAdOffsiteCampaignResponse instantiates a new CreateAdOffsiteCampaignResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAdOffsiteCampaignResponseWithDefaults

`func NewCreateAdOffsiteCampaignResponseWithDefaults() *CreateAdOffsiteCampaignResponse`

NewCreateAdOffsiteCampaignResponseWithDefaults instantiates a new CreateAdOffsiteCampaignResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignId

`func (o *CreateAdOffsiteCampaignResponse) GetCampaignId() string`

GetCampaignId returns the CampaignId field if non-nil, zero value otherwise.

### GetCampaignIdOk

`func (o *CreateAdOffsiteCampaignResponse) GetCampaignIdOk() (*string, bool)`

GetCampaignIdOk returns a tuple with the CampaignId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignId

`func (o *CreateAdOffsiteCampaignResponse) SetCampaignId(v string)`

SetCampaignId sets CampaignId field to given value.


### GetNetwork

`func (o *CreateAdOffsiteCampaignResponse) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *CreateAdOffsiteCampaignResponse) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *CreateAdOffsiteCampaignResponse) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetStatus

`func (o *CreateAdOffsiteCampaignResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateAdOffsiteCampaignResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateAdOffsiteCampaignResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCreativeHeadline

`func (o *CreateAdOffsiteCampaignResponse) GetCreativeHeadline() string`

GetCreativeHeadline returns the CreativeHeadline field if non-nil, zero value otherwise.

### GetCreativeHeadlineOk

`func (o *CreateAdOffsiteCampaignResponse) GetCreativeHeadlineOk() (*string, bool)`

GetCreativeHeadlineOk returns a tuple with the CreativeHeadline field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreativeHeadline

`func (o *CreateAdOffsiteCampaignResponse) SetCreativeHeadline(v string)`

SetCreativeHeadline sets CreativeHeadline field to given value.

### HasCreativeHeadline

`func (o *CreateAdOffsiteCampaignResponse) HasCreativeHeadline() bool`

HasCreativeHeadline returns a boolean if a field has been set.

### SetCreativeHeadlineNil

`func (o *CreateAdOffsiteCampaignResponse) SetCreativeHeadlineNil(b bool)`

 SetCreativeHeadlineNil sets the value for CreativeHeadline to be an explicit nil

### UnsetCreativeHeadline
`func (o *CreateAdOffsiteCampaignResponse) UnsetCreativeHeadline()`

UnsetCreativeHeadline ensures that no value is present for CreativeHeadline, not even an explicit nil
### GetDailyBudgetCents

`func (o *CreateAdOffsiteCampaignResponse) GetDailyBudgetCents() float32`

GetDailyBudgetCents returns the DailyBudgetCents field if non-nil, zero value otherwise.

### GetDailyBudgetCentsOk

`func (o *CreateAdOffsiteCampaignResponse) GetDailyBudgetCentsOk() (*float32, bool)`

GetDailyBudgetCentsOk returns a tuple with the DailyBudgetCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDailyBudgetCents

`func (o *CreateAdOffsiteCampaignResponse) SetDailyBudgetCents(v float32)`

SetDailyBudgetCents sets DailyBudgetCents field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


