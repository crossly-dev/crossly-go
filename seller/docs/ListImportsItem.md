# ListImportsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Status** | **string** |  | 
**Platform** | **string** |  | 
**LastError** | Pointer to **NullableString** |  | [optional] 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**Filters** | Pointer to [**NullableListImportsItemFilters**](ListImportsItemFilters.md) |  | [optional] 
**FinishedAt** | Pointer to **NullableTime** |  | [optional] 
**BatchId** | Pointer to **NullableString** |  | [optional] 
**ImportMode** | **string** |  | 
**IncludeInactive** | **bool** |  | 
**MergeIntoExistingListing** | **bool** |  | 
**MergeIntoExistingInventory** | **bool** |  | 
**TotalCount** | **float32** |  | 
**ImportedCount** | **float32** |  | 
**SkippedCount** | **float32** |  | 
**ErrorCount** | **float32** |  | 

## Methods

### NewListImportsItem

`func NewListImportsItem(id string, createdAt time.Time, userId string, status string, platform string, importMode string, includeInactive bool, mergeIntoExistingListing bool, mergeIntoExistingInventory bool, totalCount float32, importedCount float32, skippedCount float32, errorCount float32, ) *ListImportsItem`

NewListImportsItem instantiates a new ListImportsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImportsItemWithDefaults

`func NewListImportsItemWithDefaults() *ListImportsItem`

NewListImportsItemWithDefaults instantiates a new ListImportsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListImportsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListImportsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListImportsItem) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListImportsItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListImportsItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListImportsItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *ListImportsItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListImportsItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListImportsItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetStatus

`func (o *ListImportsItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListImportsItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListImportsItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlatform

`func (o *ListImportsItem) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListImportsItem) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListImportsItem) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetLastError

`func (o *ListImportsItem) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ListImportsItem) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ListImportsItem) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ListImportsItem) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *ListImportsItem) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *ListImportsItem) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetStartedAt

`func (o *ListImportsItem) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ListImportsItem) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ListImportsItem) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ListImportsItem) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *ListImportsItem) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *ListImportsItem) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetFilters

`func (o *ListImportsItem) GetFilters() ListImportsItemFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *ListImportsItem) GetFiltersOk() (*ListImportsItemFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *ListImportsItem) SetFilters(v ListImportsItemFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *ListImportsItem) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *ListImportsItem) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *ListImportsItem) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetFinishedAt

`func (o *ListImportsItem) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *ListImportsItem) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *ListImportsItem) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *ListImportsItem) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *ListImportsItem) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *ListImportsItem) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetBatchId

`func (o *ListImportsItem) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *ListImportsItem) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *ListImportsItem) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *ListImportsItem) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### SetBatchIdNil

`func (o *ListImportsItem) SetBatchIdNil(b bool)`

 SetBatchIdNil sets the value for BatchId to be an explicit nil

### UnsetBatchId
`func (o *ListImportsItem) UnsetBatchId()`

UnsetBatchId ensures that no value is present for BatchId, not even an explicit nil
### GetImportMode

`func (o *ListImportsItem) GetImportMode() string`

GetImportMode returns the ImportMode field if non-nil, zero value otherwise.

### GetImportModeOk

`func (o *ListImportsItem) GetImportModeOk() (*string, bool)`

GetImportModeOk returns a tuple with the ImportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMode

`func (o *ListImportsItem) SetImportMode(v string)`

SetImportMode sets ImportMode field to given value.


### GetIncludeInactive

`func (o *ListImportsItem) GetIncludeInactive() bool`

GetIncludeInactive returns the IncludeInactive field if non-nil, zero value otherwise.

### GetIncludeInactiveOk

`func (o *ListImportsItem) GetIncludeInactiveOk() (*bool, bool)`

GetIncludeInactiveOk returns a tuple with the IncludeInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInactive

`func (o *ListImportsItem) SetIncludeInactive(v bool)`

SetIncludeInactive sets IncludeInactive field to given value.


### GetMergeIntoExistingListing

`func (o *ListImportsItem) GetMergeIntoExistingListing() bool`

GetMergeIntoExistingListing returns the MergeIntoExistingListing field if non-nil, zero value otherwise.

### GetMergeIntoExistingListingOk

`func (o *ListImportsItem) GetMergeIntoExistingListingOk() (*bool, bool)`

GetMergeIntoExistingListingOk returns a tuple with the MergeIntoExistingListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeIntoExistingListing

`func (o *ListImportsItem) SetMergeIntoExistingListing(v bool)`

SetMergeIntoExistingListing sets MergeIntoExistingListing field to given value.


### GetMergeIntoExistingInventory

`func (o *ListImportsItem) GetMergeIntoExistingInventory() bool`

GetMergeIntoExistingInventory returns the MergeIntoExistingInventory field if non-nil, zero value otherwise.

### GetMergeIntoExistingInventoryOk

`func (o *ListImportsItem) GetMergeIntoExistingInventoryOk() (*bool, bool)`

GetMergeIntoExistingInventoryOk returns a tuple with the MergeIntoExistingInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeIntoExistingInventory

`func (o *ListImportsItem) SetMergeIntoExistingInventory(v bool)`

SetMergeIntoExistingInventory sets MergeIntoExistingInventory field to given value.


### GetTotalCount

`func (o *ListImportsItem) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ListImportsItem) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ListImportsItem) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.


### GetImportedCount

`func (o *ListImportsItem) GetImportedCount() float32`

GetImportedCount returns the ImportedCount field if non-nil, zero value otherwise.

### GetImportedCountOk

`func (o *ListImportsItem) GetImportedCountOk() (*float32, bool)`

GetImportedCountOk returns a tuple with the ImportedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedCount

`func (o *ListImportsItem) SetImportedCount(v float32)`

SetImportedCount sets ImportedCount field to given value.


### GetSkippedCount

`func (o *ListImportsItem) GetSkippedCount() float32`

GetSkippedCount returns the SkippedCount field if non-nil, zero value otherwise.

### GetSkippedCountOk

`func (o *ListImportsItem) GetSkippedCountOk() (*float32, bool)`

GetSkippedCountOk returns a tuple with the SkippedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedCount

`func (o *ListImportsItem) SetSkippedCount(v float32)`

SetSkippedCount sets SkippedCount field to given value.


### GetErrorCount

`func (o *ListImportsItem) GetErrorCount() float32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *ListImportsItem) GetErrorCountOk() (*float32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *ListImportsItem) SetErrorCount(v float32)`

SetErrorCount sets ErrorCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


