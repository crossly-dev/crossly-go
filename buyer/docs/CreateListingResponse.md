# CreateListingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Listing** | [**CreateListingResponseListing**](CreateListingResponseListing.md) |  | 
**Jobs** | [**[]CreateListingResponseJobs**](CreateListingResponseJobs.md) |  | 
**Skipped** | [**[]CreateListingResponseSkipped**](CreateListingResponseSkipped.md) |  | 
**BulkJobId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateListingResponse

`func NewCreateListingResponse(listing CreateListingResponseListing, jobs []CreateListingResponseJobs, skipped []CreateListingResponseSkipped, ) *CreateListingResponse`

NewCreateListingResponse instantiates a new CreateListingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateListingResponseWithDefaults

`func NewCreateListingResponseWithDefaults() *CreateListingResponse`

NewCreateListingResponseWithDefaults instantiates a new CreateListingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListing

`func (o *CreateListingResponse) GetListing() CreateListingResponseListing`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *CreateListingResponse) GetListingOk() (*CreateListingResponseListing, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *CreateListingResponse) SetListing(v CreateListingResponseListing)`

SetListing sets Listing field to given value.


### GetJobs

`func (o *CreateListingResponse) GetJobs() []CreateListingResponseJobs`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *CreateListingResponse) GetJobsOk() (*[]CreateListingResponseJobs, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *CreateListingResponse) SetJobs(v []CreateListingResponseJobs)`

SetJobs sets Jobs field to given value.


### GetSkipped

`func (o *CreateListingResponse) GetSkipped() []CreateListingResponseSkipped`

GetSkipped returns the Skipped field if non-nil, zero value otherwise.

### GetSkippedOk

`func (o *CreateListingResponse) GetSkippedOk() (*[]CreateListingResponseSkipped, bool)`

GetSkippedOk returns a tuple with the Skipped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkipped

`func (o *CreateListingResponse) SetSkipped(v []CreateListingResponseSkipped)`

SetSkipped sets Skipped field to given value.


### GetBulkJobId

`func (o *CreateListingResponse) GetBulkJobId() string`

GetBulkJobId returns the BulkJobId field if non-nil, zero value otherwise.

### GetBulkJobIdOk

`func (o *CreateListingResponse) GetBulkJobIdOk() (*string, bool)`

GetBulkJobIdOk returns a tuple with the BulkJobId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBulkJobId

`func (o *CreateListingResponse) SetBulkJobId(v string)`

SetBulkJobId sets BulkJobId field to given value.

### HasBulkJobId

`func (o *CreateListingResponse) HasBulkJobId() bool`

HasBulkJobId returns a boolean if a field has been set.

### SetBulkJobIdNil

`func (o *CreateListingResponse) SetBulkJobIdNil(b bool)`

 SetBulkJobIdNil sets the value for BulkJobId to be an explicit nil

### UnsetBulkJobId
`func (o *CreateListingResponse) UnsetBulkJobId()`

UnsetBulkJobId ensures that no value is present for BulkJobId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


