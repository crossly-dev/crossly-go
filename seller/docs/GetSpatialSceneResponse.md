# GetSpatialSceneResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Breakdowns** | [**[]GetSpatialSceneResponseBreakdowns**](GetSpatialSceneResponseBreakdowns.md) | Facets that would actually split this collection, most-dividing first.  Drives the breakdown picker — see availableBreakdowns. | 
**Scene** | [**GetSpatialSceneResponseScene**](GetSpatialSceneResponseScene.md) |  | 
**Profile** | [**GetSpatialPublicResponseProfile**](GetSpatialPublicResponseProfile.md) |  | 
**Solved** | [**GetSpatialPublicResponseSolved**](GetSpatialPublicResponseSolved.md) |  | 
**Items** | [**[]GetSpatialSceneResponseItems**](GetSpatialSceneResponseItems.md) |  | 
**Stats** | [**GetSpatialSceneResponseStats**](GetSpatialSceneResponseStats.md) |  | 

## Methods

### NewGetSpatialSceneResponse

`func NewGetSpatialSceneResponse(breakdowns []GetSpatialSceneResponseBreakdowns, scene GetSpatialSceneResponseScene, profile GetSpatialPublicResponseProfile, solved GetSpatialPublicResponseSolved, items []GetSpatialSceneResponseItems, stats GetSpatialSceneResponseStats, ) *GetSpatialSceneResponse`

NewGetSpatialSceneResponse instantiates a new GetSpatialSceneResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSpatialSceneResponseWithDefaults

`func NewGetSpatialSceneResponseWithDefaults() *GetSpatialSceneResponse`

NewGetSpatialSceneResponseWithDefaults instantiates a new GetSpatialSceneResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBreakdowns

`func (o *GetSpatialSceneResponse) GetBreakdowns() []GetSpatialSceneResponseBreakdowns`

GetBreakdowns returns the Breakdowns field if non-nil, zero value otherwise.

### GetBreakdownsOk

`func (o *GetSpatialSceneResponse) GetBreakdownsOk() (*[]GetSpatialSceneResponseBreakdowns, bool)`

GetBreakdownsOk returns a tuple with the Breakdowns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreakdowns

`func (o *GetSpatialSceneResponse) SetBreakdowns(v []GetSpatialSceneResponseBreakdowns)`

SetBreakdowns sets Breakdowns field to given value.


### GetScene

`func (o *GetSpatialSceneResponse) GetScene() GetSpatialSceneResponseScene`

GetScene returns the Scene field if non-nil, zero value otherwise.

### GetSceneOk

`func (o *GetSpatialSceneResponse) GetSceneOk() (*GetSpatialSceneResponseScene, bool)`

GetSceneOk returns a tuple with the Scene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScene

`func (o *GetSpatialSceneResponse) SetScene(v GetSpatialSceneResponseScene)`

SetScene sets Scene field to given value.


### GetProfile

`func (o *GetSpatialSceneResponse) GetProfile() GetSpatialPublicResponseProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetSpatialSceneResponse) GetProfileOk() (*GetSpatialPublicResponseProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetSpatialSceneResponse) SetProfile(v GetSpatialPublicResponseProfile)`

SetProfile sets Profile field to given value.


### GetSolved

`func (o *GetSpatialSceneResponse) GetSolved() GetSpatialPublicResponseSolved`

GetSolved returns the Solved field if non-nil, zero value otherwise.

### GetSolvedOk

`func (o *GetSpatialSceneResponse) GetSolvedOk() (*GetSpatialPublicResponseSolved, bool)`

GetSolvedOk returns a tuple with the Solved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolved

`func (o *GetSpatialSceneResponse) SetSolved(v GetSpatialPublicResponseSolved)`

SetSolved sets Solved field to given value.


### GetItems

`func (o *GetSpatialSceneResponse) GetItems() []GetSpatialSceneResponseItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetSpatialSceneResponse) GetItemsOk() (*[]GetSpatialSceneResponseItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetSpatialSceneResponse) SetItems(v []GetSpatialSceneResponseItems)`

SetItems sets Items field to given value.


### GetStats

`func (o *GetSpatialSceneResponse) GetStats() GetSpatialSceneResponseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetSpatialSceneResponse) GetStatsOk() (*GetSpatialSceneResponseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetSpatialSceneResponse) SetStats(v GetSpatialSceneResponseStats)`

SetStats sets Stats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


