# GetAccountDeletionStatusResponsePending

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**RequestedAt** | **time.Time** |  | 
**ScheduledFor** | **time.Time** |  | 

## Methods

### NewGetAccountDeletionStatusResponsePending

`func NewGetAccountDeletionStatusResponsePending(id string, requestedAt time.Time, scheduledFor time.Time, ) *GetAccountDeletionStatusResponsePending`

NewGetAccountDeletionStatusResponsePending instantiates a new GetAccountDeletionStatusResponsePending object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAccountDeletionStatusResponsePendingWithDefaults

`func NewGetAccountDeletionStatusResponsePendingWithDefaults() *GetAccountDeletionStatusResponsePending`

NewGetAccountDeletionStatusResponsePendingWithDefaults instantiates a new GetAccountDeletionStatusResponsePending object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetAccountDeletionStatusResponsePending) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetAccountDeletionStatusResponsePending) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetAccountDeletionStatusResponsePending) SetId(v string)`

SetId sets Id field to given value.


### GetRequestedAt

`func (o *GetAccountDeletionStatusResponsePending) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *GetAccountDeletionStatusResponsePending) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *GetAccountDeletionStatusResponsePending) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.


### GetScheduledFor

`func (o *GetAccountDeletionStatusResponsePending) GetScheduledFor() time.Time`

GetScheduledFor returns the ScheduledFor field if non-nil, zero value otherwise.

### GetScheduledForOk

`func (o *GetAccountDeletionStatusResponsePending) GetScheduledForOk() (*time.Time, bool)`

GetScheduledForOk returns a tuple with the ScheduledFor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledFor

`func (o *GetAccountDeletionStatusResponsePending) SetScheduledFor(v time.Time)`

SetScheduledFor sets ScheduledFor field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


