# ListListingDiscrepanciesItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Source** | **string** |  | 
**Platform** | **string** |  | 
**ListingId** | **string** |  | 
**AccountSlot** | **float32** |  | 
**LastError** | Pointer to **NullableString** |  | [optional] 
**Field** | **string** |  | 
**OurValue** | Pointer to **NullableString** |  | [optional] 
**RemoteValue** | Pointer to **NullableString** |  | [optional] 
**FirstDetectedAt** | **time.Time** |  | 
**LastDetectedAt** | **time.Time** |  | 
**DismissedAt** | Pointer to **NullableTime** |  | [optional] 
**DismissedRemoteValue** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListListingDiscrepanciesItem

`func NewListListingDiscrepanciesItem(id string, createdAt time.Time, updatedAt time.Time, userId string, source string, platform string, listingId string, accountSlot float32, field string, firstDetectedAt time.Time, lastDetectedAt time.Time, ) *ListListingDiscrepanciesItem`

NewListListingDiscrepanciesItem instantiates a new ListListingDiscrepanciesItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListListingDiscrepanciesItemWithDefaults

`func NewListListingDiscrepanciesItemWithDefaults() *ListListingDiscrepanciesItem`

NewListListingDiscrepanciesItemWithDefaults instantiates a new ListListingDiscrepanciesItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListListingDiscrepanciesItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListListingDiscrepanciesItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListListingDiscrepanciesItem) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListListingDiscrepanciesItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListListingDiscrepanciesItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListListingDiscrepanciesItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListListingDiscrepanciesItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListListingDiscrepanciesItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListListingDiscrepanciesItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *ListListingDiscrepanciesItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListListingDiscrepanciesItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListListingDiscrepanciesItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetSource

`func (o *ListListingDiscrepanciesItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListListingDiscrepanciesItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListListingDiscrepanciesItem) SetSource(v string)`

SetSource sets Source field to given value.


### GetPlatform

`func (o *ListListingDiscrepanciesItem) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListListingDiscrepanciesItem) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListListingDiscrepanciesItem) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetListingId

`func (o *ListListingDiscrepanciesItem) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *ListListingDiscrepanciesItem) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *ListListingDiscrepanciesItem) SetListingId(v string)`

SetListingId sets ListingId field to given value.


### GetAccountSlot

`func (o *ListListingDiscrepanciesItem) GetAccountSlot() float32`

GetAccountSlot returns the AccountSlot field if non-nil, zero value otherwise.

### GetAccountSlotOk

`func (o *ListListingDiscrepanciesItem) GetAccountSlotOk() (*float32, bool)`

GetAccountSlotOk returns a tuple with the AccountSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSlot

`func (o *ListListingDiscrepanciesItem) SetAccountSlot(v float32)`

SetAccountSlot sets AccountSlot field to given value.


### GetLastError

`func (o *ListListingDiscrepanciesItem) GetLastError() string`

GetLastError returns the LastError field if non-nil, zero value otherwise.

### GetLastErrorOk

`func (o *ListListingDiscrepanciesItem) GetLastErrorOk() (*string, bool)`

GetLastErrorOk returns a tuple with the LastError field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastError

`func (o *ListListingDiscrepanciesItem) SetLastError(v string)`

SetLastError sets LastError field to given value.

### HasLastError

`func (o *ListListingDiscrepanciesItem) HasLastError() bool`

HasLastError returns a boolean if a field has been set.

### SetLastErrorNil

`func (o *ListListingDiscrepanciesItem) SetLastErrorNil(b bool)`

 SetLastErrorNil sets the value for LastError to be an explicit nil

### UnsetLastError
`func (o *ListListingDiscrepanciesItem) UnsetLastError()`

UnsetLastError ensures that no value is present for LastError, not even an explicit nil
### GetField

`func (o *ListListingDiscrepanciesItem) GetField() string`

GetField returns the Field field if non-nil, zero value otherwise.

### GetFieldOk

`func (o *ListListingDiscrepanciesItem) GetFieldOk() (*string, bool)`

GetFieldOk returns a tuple with the Field field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetField

`func (o *ListListingDiscrepanciesItem) SetField(v string)`

SetField sets Field field to given value.


### GetOurValue

`func (o *ListListingDiscrepanciesItem) GetOurValue() string`

GetOurValue returns the OurValue field if non-nil, zero value otherwise.

### GetOurValueOk

`func (o *ListListingDiscrepanciesItem) GetOurValueOk() (*string, bool)`

GetOurValueOk returns a tuple with the OurValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOurValue

`func (o *ListListingDiscrepanciesItem) SetOurValue(v string)`

SetOurValue sets OurValue field to given value.

### HasOurValue

`func (o *ListListingDiscrepanciesItem) HasOurValue() bool`

HasOurValue returns a boolean if a field has been set.

### SetOurValueNil

`func (o *ListListingDiscrepanciesItem) SetOurValueNil(b bool)`

 SetOurValueNil sets the value for OurValue to be an explicit nil

### UnsetOurValue
`func (o *ListListingDiscrepanciesItem) UnsetOurValue()`

UnsetOurValue ensures that no value is present for OurValue, not even an explicit nil
### GetRemoteValue

`func (o *ListListingDiscrepanciesItem) GetRemoteValue() string`

GetRemoteValue returns the RemoteValue field if non-nil, zero value otherwise.

### GetRemoteValueOk

`func (o *ListListingDiscrepanciesItem) GetRemoteValueOk() (*string, bool)`

GetRemoteValueOk returns a tuple with the RemoteValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteValue

`func (o *ListListingDiscrepanciesItem) SetRemoteValue(v string)`

SetRemoteValue sets RemoteValue field to given value.

### HasRemoteValue

`func (o *ListListingDiscrepanciesItem) HasRemoteValue() bool`

HasRemoteValue returns a boolean if a field has been set.

### SetRemoteValueNil

`func (o *ListListingDiscrepanciesItem) SetRemoteValueNil(b bool)`

 SetRemoteValueNil sets the value for RemoteValue to be an explicit nil

### UnsetRemoteValue
`func (o *ListListingDiscrepanciesItem) UnsetRemoteValue()`

UnsetRemoteValue ensures that no value is present for RemoteValue, not even an explicit nil
### GetFirstDetectedAt

`func (o *ListListingDiscrepanciesItem) GetFirstDetectedAt() time.Time`

GetFirstDetectedAt returns the FirstDetectedAt field if non-nil, zero value otherwise.

### GetFirstDetectedAtOk

`func (o *ListListingDiscrepanciesItem) GetFirstDetectedAtOk() (*time.Time, bool)`

GetFirstDetectedAtOk returns a tuple with the FirstDetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstDetectedAt

`func (o *ListListingDiscrepanciesItem) SetFirstDetectedAt(v time.Time)`

SetFirstDetectedAt sets FirstDetectedAt field to given value.


### GetLastDetectedAt

`func (o *ListListingDiscrepanciesItem) GetLastDetectedAt() time.Time`

GetLastDetectedAt returns the LastDetectedAt field if non-nil, zero value otherwise.

### GetLastDetectedAtOk

`func (o *ListListingDiscrepanciesItem) GetLastDetectedAtOk() (*time.Time, bool)`

GetLastDetectedAtOk returns a tuple with the LastDetectedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDetectedAt

`func (o *ListListingDiscrepanciesItem) SetLastDetectedAt(v time.Time)`

SetLastDetectedAt sets LastDetectedAt field to given value.


### GetDismissedAt

`func (o *ListListingDiscrepanciesItem) GetDismissedAt() time.Time`

GetDismissedAt returns the DismissedAt field if non-nil, zero value otherwise.

### GetDismissedAtOk

`func (o *ListListingDiscrepanciesItem) GetDismissedAtOk() (*time.Time, bool)`

GetDismissedAtOk returns a tuple with the DismissedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedAt

`func (o *ListListingDiscrepanciesItem) SetDismissedAt(v time.Time)`

SetDismissedAt sets DismissedAt field to given value.

### HasDismissedAt

`func (o *ListListingDiscrepanciesItem) HasDismissedAt() bool`

HasDismissedAt returns a boolean if a field has been set.

### SetDismissedAtNil

`func (o *ListListingDiscrepanciesItem) SetDismissedAtNil(b bool)`

 SetDismissedAtNil sets the value for DismissedAt to be an explicit nil

### UnsetDismissedAt
`func (o *ListListingDiscrepanciesItem) UnsetDismissedAt()`

UnsetDismissedAt ensures that no value is present for DismissedAt, not even an explicit nil
### GetDismissedRemoteValue

`func (o *ListListingDiscrepanciesItem) GetDismissedRemoteValue() string`

GetDismissedRemoteValue returns the DismissedRemoteValue field if non-nil, zero value otherwise.

### GetDismissedRemoteValueOk

`func (o *ListListingDiscrepanciesItem) GetDismissedRemoteValueOk() (*string, bool)`

GetDismissedRemoteValueOk returns a tuple with the DismissedRemoteValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDismissedRemoteValue

`func (o *ListListingDiscrepanciesItem) SetDismissedRemoteValue(v string)`

SetDismissedRemoteValue sets DismissedRemoteValue field to given value.

### HasDismissedRemoteValue

`func (o *ListListingDiscrepanciesItem) HasDismissedRemoteValue() bool`

HasDismissedRemoteValue returns a boolean if a field has been set.

### SetDismissedRemoteValueNil

`func (o *ListListingDiscrepanciesItem) SetDismissedRemoteValueNil(b bool)`

 SetDismissedRemoteValueNil sets the value for DismissedRemoteValue to be an explicit nil

### UnsetDismissedRemoteValue
`func (o *ListListingDiscrepanciesItem) UnsetDismissedRemoteValue()`

UnsetDismissedRemoteValue ensures that no value is present for DismissedRemoteValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


