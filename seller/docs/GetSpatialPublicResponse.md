# GetSpatialPublicResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Scene** | [**GetSpatialPublicResponseScene**](GetSpatialPublicResponseScene.md) |  | 
**Profile** | [**GetSpatialPublicResponseProfile**](GetSpatialPublicResponseProfile.md) |  | 
**Solved** | [**GetSpatialPublicResponseSolved**](GetSpatialPublicResponseSolved.md) |  | 
**Items** | [**[]GetSpatialPublicResponseItems**](GetSpatialPublicResponseItems.md) | Enough to draw an object and label it. No cost, no location, no status. | 
**Stats** | [**GetSpatialPublicResponseStats**](GetSpatialPublicResponseStats.md) |  | 

## Methods

### NewGetSpatialPublicResponse

`func NewGetSpatialPublicResponse(scene GetSpatialPublicResponseScene, profile GetSpatialPublicResponseProfile, solved GetSpatialPublicResponseSolved, items []GetSpatialPublicResponseItems, stats GetSpatialPublicResponseStats, ) *GetSpatialPublicResponse`

NewGetSpatialPublicResponse instantiates a new GetSpatialPublicResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSpatialPublicResponseWithDefaults

`func NewGetSpatialPublicResponseWithDefaults() *GetSpatialPublicResponse`

NewGetSpatialPublicResponseWithDefaults instantiates a new GetSpatialPublicResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScene

`func (o *GetSpatialPublicResponse) GetScene() GetSpatialPublicResponseScene`

GetScene returns the Scene field if non-nil, zero value otherwise.

### GetSceneOk

`func (o *GetSpatialPublicResponse) GetSceneOk() (*GetSpatialPublicResponseScene, bool)`

GetSceneOk returns a tuple with the Scene field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScene

`func (o *GetSpatialPublicResponse) SetScene(v GetSpatialPublicResponseScene)`

SetScene sets Scene field to given value.


### GetProfile

`func (o *GetSpatialPublicResponse) GetProfile() GetSpatialPublicResponseProfile`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *GetSpatialPublicResponse) GetProfileOk() (*GetSpatialPublicResponseProfile, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *GetSpatialPublicResponse) SetProfile(v GetSpatialPublicResponseProfile)`

SetProfile sets Profile field to given value.


### GetSolved

`func (o *GetSpatialPublicResponse) GetSolved() GetSpatialPublicResponseSolved`

GetSolved returns the Solved field if non-nil, zero value otherwise.

### GetSolvedOk

`func (o *GetSpatialPublicResponse) GetSolvedOk() (*GetSpatialPublicResponseSolved, bool)`

GetSolvedOk returns a tuple with the Solved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSolved

`func (o *GetSpatialPublicResponse) SetSolved(v GetSpatialPublicResponseSolved)`

SetSolved sets Solved field to given value.


### GetItems

`func (o *GetSpatialPublicResponse) GetItems() []GetSpatialPublicResponseItems`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetSpatialPublicResponse) GetItemsOk() (*[]GetSpatialPublicResponseItems, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetSpatialPublicResponse) SetItems(v []GetSpatialPublicResponseItems)`

SetItems sets Items field to given value.


### GetStats

`func (o *GetSpatialPublicResponse) GetStats() GetSpatialPublicResponseStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetSpatialPublicResponse) GetStatsOk() (*GetSpatialPublicResponseStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetSpatialPublicResponse) SetStats(v GetSpatialPublicResponseStats)`

SetStats sets Stats field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


