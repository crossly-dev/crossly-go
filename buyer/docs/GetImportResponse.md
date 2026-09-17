# GetImportResponse

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

### NewGetImportResponse

`func NewGetImportResponse(id string, createdAt time.Time, userId string, status string, platform string, importMode string, includeInactive bool, mergeIntoExistingListing bool, mergeIntoExistingInventory bool, totalCount float32, importedCount float32, skippedCount float32, errorCount float32, ) *GetImportResponse`

NewGetImportResponse instantiates a new GetImportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImportResponseWithDefaults

`func NewGetImportResponseWithDefaults() *GetImportResponse`

NewGetImportResponseWithDefaults instantiates a new GetImportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetImportResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetImportResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetImportResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetImportResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetImportResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetImportResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *GetImportResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetImportResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetImportResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetStatus

`func (o *GetImportResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetImportResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetImportResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlatform

`func (o *GetImportResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetImportResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetImportResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetLastError

`func (o *GetImportResponse) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *GetImportResponse) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *GetImportResponse) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *GetImportResponse) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *GetImportResponse) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *GetImportResponse) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetStartedAt

`func (o *GetImportResponse) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetImportResponse) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetImportResponse) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GetImportResponse) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *GetImportResponse) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *GetImportResponse) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetFilters

`func (o *GetImportResponse) GetFilters() ListImportsItemFilters`

GetFilters returns the Filters field if non-nil, zero value otherwise.

### GetFiltersOk

`func (o *GetImportResponse) GetFiltersOk() (*ListImportsItemFilters, bool)`

GetFiltersOk returns a tuple with the Filters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilters

`func (o *GetImportResponse) SetFilters(v ListImportsItemFilters)`

SetFilters sets Filters field to given value.

### HasFilters

`func (o *GetImportResponse) HasFilters() bool`

HasFilters returns a boolean if a field has been set.

### SetFiltersNil

`func (o *GetImportResponse) SetFiltersNil(b bool)`

 SetFiltersNil sets the value for Filters to be an explicit nil

### UnsetFilters
`func (o *GetImportResponse) UnsetFilters()`

UnsetFilters ensures that no value is present for Filters, not even an explicit nil
### GetFinishedAt

`func (o *GetImportResponse) GetFinishedAt() time.Time`

GetFinishedAt returns the FinishedAt field if non-nil, zero value otherwise.

### GetFinishedAtOk

`func (o *GetImportResponse) GetFinishedAtOk() (*time.Time, bool)`

GetFinishedAtOk returns a tuple with the FinishedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFinishedAt

`func (o *GetImportResponse) SetFinishedAt(v time.Time)`

SetFinishedAt sets FinishedAt field to given value.

### HasFinishedAt

`func (o *GetImportResponse) HasFinishedAt() bool`

HasFinishedAt returns a boolean if a field has been set.

### SetFinishedAtNil

`func (o *GetImportResponse) SetFinishedAtNil(b bool)`

 SetFinishedAtNil sets the value for FinishedAt to be an explicit nil

### UnsetFinishedAt
`func (o *GetImportResponse) UnsetFinishedAt()`

UnsetFinishedAt ensures that no value is present for FinishedAt, not even an explicit nil
### GetBatchId

`func (o *GetImportResponse) GetBatchId() string`

GetBatchId returns the BatchId field if non-nil, zero value otherwise.

### GetBatchIdOk

`func (o *GetImportResponse) GetBatchIdOk() (*string, bool)`

GetBatchIdOk returns a tuple with the BatchId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBatchId

`func (o *GetImportResponse) SetBatchId(v string)`

SetBatchId sets BatchId field to given value.

### HasBatchId

`func (o *GetImportResponse) HasBatchId() bool`

HasBatchId returns a boolean if a field has been set.

### SetBatchIdNil

`func (o *GetImportResponse) SetBatchIdNil(b bool)`

 SetBatchIdNil sets the value for BatchId to be an explicit nil

### UnsetBatchId
`func (o *GetImportResponse) UnsetBatchId()`

UnsetBatchId ensures that no value is present for BatchId, not even an explicit nil
### GetImportMode

`func (o *GetImportResponse) GetImportMode() string`

GetImportMode returns the ImportMode field if non-nil, zero value otherwise.

### GetImportModeOk

`func (o *GetImportResponse) GetImportModeOk() (*string, bool)`

GetImportModeOk returns a tuple with the ImportMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportMode

`func (o *GetImportResponse) SetImportMode(v string)`

SetImportMode sets ImportMode field to given value.


### GetIncludeInactive

`func (o *GetImportResponse) GetIncludeInactive() bool`

GetIncludeInactive returns the IncludeInactive field if non-nil, zero value otherwise.

### GetIncludeInactiveOk

`func (o *GetImportResponse) GetIncludeInactiveOk() (*bool, bool)`

GetIncludeInactiveOk returns a tuple with the IncludeInactive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeInactive

`func (o *GetImportResponse) SetIncludeInactive(v bool)`

SetIncludeInactive sets IncludeInactive field to given value.


### GetMergeIntoExistingListing

`func (o *GetImportResponse) GetMergeIntoExistingListing() bool`

GetMergeIntoExistingListing returns the MergeIntoExistingListing field if non-nil, zero value otherwise.

### GetMergeIntoExistingListingOk

`func (o *GetImportResponse) GetMergeIntoExistingListingOk() (*bool, bool)`

GetMergeIntoExistingListingOk returns a tuple with the MergeIntoExistingListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeIntoExistingListing

`func (o *GetImportResponse) SetMergeIntoExistingListing(v bool)`

SetMergeIntoExistingListing sets MergeIntoExistingListing field to given value.


### GetMergeIntoExistingInventory

`func (o *GetImportResponse) GetMergeIntoExistingInventory() bool`

GetMergeIntoExistingInventory returns the MergeIntoExistingInventory field if non-nil, zero value otherwise.

### GetMergeIntoExistingInventoryOk

`func (o *GetImportResponse) GetMergeIntoExistingInventoryOk() (*bool, bool)`

GetMergeIntoExistingInventoryOk returns a tuple with the MergeIntoExistingInventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMergeIntoExistingInventory

`func (o *GetImportResponse) SetMergeIntoExistingInventory(v bool)`

SetMergeIntoExistingInventory sets MergeIntoExistingInventory field to given value.


### GetTotalCount

`func (o *GetImportResponse) GetTotalCount() float32`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *GetImportResponse) GetTotalCountOk() (*float32, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *GetImportResponse) SetTotalCount(v float32)`

SetTotalCount sets TotalCount field to given value.


### GetImportedCount

`func (o *GetImportResponse) GetImportedCount() float32`

GetImportedCount returns the ImportedCount field if non-nil, zero value otherwise.

### GetImportedCountOk

`func (o *GetImportResponse) GetImportedCountOk() (*float32, bool)`

GetImportedCountOk returns a tuple with the ImportedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedCount

`func (o *GetImportResponse) SetImportedCount(v float32)`

SetImportedCount sets ImportedCount field to given value.


### GetSkippedCount

`func (o *GetImportResponse) GetSkippedCount() float32`

GetSkippedCount returns the SkippedCount field if non-nil, zero value otherwise.

### GetSkippedCountOk

`func (o *GetImportResponse) GetSkippedCountOk() (*float32, bool)`

GetSkippedCountOk returns a tuple with the SkippedCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkippedCount

`func (o *GetImportResponse) SetSkippedCount(v float32)`

SetSkippedCount sets SkippedCount field to given value.


### GetErrorCount

`func (o *GetImportResponse) GetErrorCount() float32`

GetErrorCount returns the ErrorCount field if non-nil, zero value otherwise.

### GetErrorCountOk

`func (o *GetImportResponse) GetErrorCountOk() (*float32, bool)`

GetErrorCountOk returns a tuple with the ErrorCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCount

`func (o *GetImportResponse) SetErrorCount(v float32)`

SetErrorCount sets ErrorCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


