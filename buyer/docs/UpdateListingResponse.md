# UpdateListingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Jobs** | [**[]CreateListingResponseJobs**](CreateListingResponseJobs.md) |  | 

## Methods

### NewUpdateListingResponse

`func NewUpdateListingResponse(jobs []CreateListingResponseJobs, ) *UpdateListingResponse`

NewUpdateListingResponse instantiates a new UpdateListingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateListingResponseWithDefaults

`func NewUpdateListingResponseWithDefaults() *UpdateListingResponse`

NewUpdateListingResponseWithDefaults instantiates a new UpdateListingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetJobs

`func (o *UpdateListingResponse) GetJobs() []CreateListingResponseJobs`

GetJobs returns the Jobs field if non-nil, zero value otherwise.

### GetJobsOk

`func (o *UpdateListingResponse) GetJobsOk() (*[]CreateListingResponseJobs, bool)`

GetJobsOk returns a tuple with the Jobs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJobs

`func (o *UpdateListingResponse) SetJobs(v []CreateListingResponseJobs)`

SetJobs sets Jobs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


