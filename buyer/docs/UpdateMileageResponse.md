# UpdateMileageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Miles** | **float32** |  | 
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Date** | **string** |  | 
**Purpose** | **string** |  | 
**StartLocation** | Pointer to **NullableString** |  | [optional] 
**EndLocation** | Pointer to **NullableString** |  | [optional] 
**Notes** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewUpdateMileageResponse

`func NewUpdateMileageResponse(miles float32, id string, userId string, date string, purpose string, createdAt time.Time, ) *UpdateMileageResponse`

NewUpdateMileageResponse instantiates a new UpdateMileageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMileageResponseWithDefaults

`func NewUpdateMileageResponseWithDefaults() *UpdateMileageResponse`

NewUpdateMileageResponseWithDefaults instantiates a new UpdateMileageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMiles

`func (o *UpdateMileageResponse) GetMiles() float32`

GetMiles returns the Miles field if non-nil, zero value otherwise.

### GetMilesOk

`func (o *UpdateMileageResponse) GetMilesOk() (*float32, bool)`

GetMilesOk returns a tuple with the Miles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiles

`func (o *UpdateMileageResponse) SetMiles(v float32)`

SetMiles sets Miles field to given value.


### GetId

`func (o *UpdateMileageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateMileageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateMileageResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *UpdateMileageResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateMileageResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateMileageResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetDate

`func (o *UpdateMileageResponse) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *UpdateMileageResponse) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *UpdateMileageResponse) SetDate(v string)`

SetDate sets Date field to given value.


### GetPurpose

`func (o *UpdateMileageResponse) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *UpdateMileageResponse) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *UpdateMileageResponse) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetStartLocation

`func (o *UpdateMileageResponse) GetStartLocation() string`

GetStartLocation returns the StartLocation field if non-nil, zero value otherwise.

### GetStartLocationOk

`func (o *UpdateMileageResponse) GetStartLocationOk() (*string, bool)`

GetStartLocationOk returns a tuple with the StartLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLocation

`func (o *UpdateMileageResponse) SetStartLocation(v string)`

SetStartLocation sets StartLocation field to given value.

### HasStartLocation

`func (o *UpdateMileageResponse) HasStartLocation() bool`

HasStartLocation returns a boolean if a field has been set.

### SetStartLocationNil

`func (o *UpdateMileageResponse) SetStartLocationNil(b bool)`

 SetStartLocationNil sets the value for StartLocation to be an explicit nil

### UnsetStartLocation
`func (o *UpdateMileageResponse) UnsetStartLocation()`

UnsetStartLocation ensures that no value is present for StartLocation, not even an explicit nil
### GetEndLocation

`func (o *UpdateMileageResponse) GetEndLocation() string`

GetEndLocation returns the EndLocation field if non-nil, zero value otherwise.

### GetEndLocationOk

`func (o *UpdateMileageResponse) GetEndLocationOk() (*string, bool)`

GetEndLocationOk returns a tuple with the EndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLocation

`func (o *UpdateMileageResponse) SetEndLocation(v string)`

SetEndLocation sets EndLocation field to given value.

### HasEndLocation

`func (o *UpdateMileageResponse) HasEndLocation() bool`

HasEndLocation returns a boolean if a field has been set.

### SetEndLocationNil

`func (o *UpdateMileageResponse) SetEndLocationNil(b bool)`

 SetEndLocationNil sets the value for EndLocation to be an explicit nil

### UnsetEndLocation
`func (o *UpdateMileageResponse) UnsetEndLocation()`

UnsetEndLocation ensures that no value is present for EndLocation, not even an explicit nil
### GetNotes

`func (o *UpdateMileageResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateMileageResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateMileageResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateMileageResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *UpdateMileageResponse) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *UpdateMileageResponse) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetCreatedAt

`func (o *UpdateMileageResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateMileageResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateMileageResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


