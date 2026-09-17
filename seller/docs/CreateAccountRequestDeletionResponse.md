# CreateAccountRequestDeletionResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ScheduledFor** | **time.Time** |  | 
**GraceDays** | **float32** |  | 

## Methods

### NewCreateAccountRequestDeletionResponse

`func NewCreateAccountRequestDeletionResponse(id string, scheduledFor time.Time, graceDays float32, ) *CreateAccountRequestDeletionResponse`

NewCreateAccountRequestDeletionResponse instantiates a new CreateAccountRequestDeletionResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAccountRequestDeletionResponseWithDefaults

`func NewCreateAccountRequestDeletionResponseWithDefaults() *CreateAccountRequestDeletionResponse`

NewCreateAccountRequestDeletionResponseWithDefaults instantiates a new CreateAccountRequestDeletionResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateAccountRequestDeletionResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateAccountRequestDeletionResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateAccountRequestDeletionResponse) SetId(v string)`

SetId sets Id field to given value.


### GetScheduledFor

`func (o *CreateAccountRequestDeletionResponse) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *CreateAccountRequestDeletionResponse) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *CreateAccountRequestDeletionResponse) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.


### GetGraceDays

`func (o *CreateAccountRequestDeletionResponse) GetGraceDays() float32`

GetGraceDays returns the GraceDays field if non-nil, zero value otherwise.

### GetGraceDaysOk

`func (o *CreateAccountRequestDeletionResponse) GetGraceDaysOk() (*float32, bool)`

GetGraceDaysOk returns a tuple with the GraceDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraceDays

`func (o *CreateAccountRequestDeletionResponse) SetGraceDays(v float32)`

SetGraceDays sets GraceDays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


