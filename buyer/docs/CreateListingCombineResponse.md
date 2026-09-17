# CreateListingCombineResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeepListingId** | **string** |  | 
**MergedListingIds** | **[]string** | Listings archived into the keeper. | 
**NewQuantityAvailable** | **float32** | Stock on the keeper&#39;s item after summing. | 
**UnitsMoved** | **float32** | Units moved off the merged items. | 
**DelistJobsDispatched** | **float32** | Delist jobs dispatched for the merged listings&#39; live platforms. | 

## Methods

### NewCreateListingCombineResponse

`func NewCreateListingCombineResponse(keepListingId string, mergedListingIds []string, newQuantityAvailable float32, unitsMoved float32, delistJobsDispatched float32, ) *CreateListingCombineResponse`

NewCreateListingCombineResponse instantiates a new CreateListingCombineResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateListingCombineResponseWithDefaults

`func NewCreateListingCombineResponseWithDefaults() *CreateListingCombineResponse`

NewCreateListingCombineResponseWithDefaults instantiates a new CreateListingCombineResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeepListingId

`func (o *CreateListingCombineResponse) GetKeepListingId() string`

GetKeepListingId returns the KeepListingId field if non-nil, zero value otherwise.

### GetKeepListingIdOk

`func (o *CreateListingCombineResponse) GetKeepListingIdOk() (*string, bool)`

GetKeepListingIdOk returns a tuple with the KeepListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeepListingId

`func (o *CreateListingCombineResponse) SetKeepListingId(v string)`

SetKeepListingId sets KeepListingId field to given value.


### GetMergedListingIds

`func (o *CreateListingCombineResponse) GetMergedListingIds() []string`

GetMergedListingIds returns the MergedListingIds field if non-nil, zero value otherwise.

### GetMergedListingIdsOk

`func (o *CreateListingCombineResponse) GetMergedListingIdsOk() (*[]string, bool)`

GetMergedListingIdsOk returns a tuple with the MergedListingIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergedListingIds

`func (o *CreateListingCombineResponse) SetMergedListingIds(v []string)`

SetMergedListingIds sets MergedListingIds field to given value.


### GetNewQuantityAvailable

`func (o *CreateListingCombineResponse) GetNewQuantityAvailable() float32`

GetNewQuantityAvailable returns the NewQuantityAvailable field if non-nil, zero value otherwise.

### GetNewQuantityAvailableOk

`func (o *CreateListingCombineResponse) GetNewQuantityAvailableOk() (*float32, bool)`

GetNewQuantityAvailableOk returns a tuple with the NewQuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewQuantityAvailable

`func (o *CreateListingCombineResponse) SetNewQuantityAvailable(v float32)`

SetNewQuantityAvailable sets NewQuantityAvailable field to given value.


### GetUnitsMoved

`func (o *CreateListingCombineResponse) GetUnitsMoved() float32`

GetUnitsMoved returns the UnitsMoved field if non-nil, zero value otherwise.

### GetUnitsMovedOk

`func (o *CreateListingCombineResponse) GetUnitsMovedOk() (*float32, bool)`

GetUnitsMovedOk returns a tuple with the UnitsMoved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsMoved

`func (o *CreateListingCombineResponse) SetUnitsMoved(v float32)`

SetUnitsMoved sets UnitsMoved field to given value.


### GetDelistJobsDispatched

`func (o *CreateListingCombineResponse) GetDelistJobsDispatched() float32`

GetDelistJobsDispatched returns the DelistJobsDispatched field if non-nil, zero value otherwise.

### GetDelistJobsDispatchedOk

`func (o *CreateListingCombineResponse) GetDelistJobsDispatchedOk() (*float32, bool)`

GetDelistJobsDispatchedOk returns a tuple with the DelistJobsDispatched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelistJobsDispatched

`func (o *CreateListingCombineResponse) SetDelistJobsDispatched(v float32)`

SetDelistJobsDispatched sets DelistJobsDispatched field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


