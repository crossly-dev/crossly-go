# CreateMileageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Miles** | **float32** |  | 
**Date** | **string** |  | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Purpose** | **string** |  | 
**StartLocation** | Pointer to **NullableString** |  | [optional] 
**EndLocation** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateMileageResponse

`func NewCreateMileageResponse(miles float32, date string, id string, createdAt time.Time, userId string, purpose string, ) *CreateMileageResponse`

NewCreateMileageResponse instantiates a new CreateMileageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMileageResponseWithDefaults

`func NewCreateMileageResponseWithDefaults() *CreateMileageResponse`

NewCreateMileageResponseWithDefaults instantiates a new CreateMileageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMiles

`func (o *CreateMileageResponse) GetMiles() float32`

GetMiles returns the Miles field if non-nil, zero value otherwise.

### GetMilesOk

`func (o *CreateMileageResponse) GetMilesOk() (*float32, bool)`

GetMilesOk returns a tuple with the Miles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiles

`func (o *CreateMileageResponse) SetMiles(v float32)`

SetMiles sets Miles field to given value.


### GetDate

`func (o *CreateMileageResponse) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *CreateMileageResponse) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *CreateMileageResponse) SetDate(v string)`

SetDate sets Date field to given value.


### GetId

`func (o *CreateMileageResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateMileageResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateMileageResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *CreateMileageResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateMileageResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateMileageResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *CreateMileageResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateMileageResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateMileageResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetNotes

`func (o *CreateMileageResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateMileageResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateMileageResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateMileageResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *CreateMileageResponse) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *CreateMileageResponse) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetPurpose

`func (o *CreateMileageResponse) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *CreateMileageResponse) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *CreateMileageResponse) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetStartLocation

`func (o *CreateMileageResponse) GetStartLocation() string`

GetStartLocation returns the StartLocation field if non-nil, zero value otherwise.

### GetStartLocationOk

`func (o *CreateMileageResponse) GetStartLocationOk() (*string, bool)`

GetStartLocationOk returns a tuple with the StartLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLocation

`func (o *CreateMileageResponse) SetStartLocation(v string)`

SetStartLocation sets StartLocation field to given value.

### HasStartLocation

`func (o *CreateMileageResponse) HasStartLocation() bool`

HasStartLocation returns a boolean if a field has been set.

### SetStartLocationNil

`func (o *CreateMileageResponse) SetStartLocationNil(b bool)`

 SetStartLocationNil sets the value for StartLocation to be an explicit nil

### UnsetStartLocation
`func (o *CreateMileageResponse) UnsetStartLocation()`

UnsetStartLocation ensures that no value is present for StartLocation, not even an explicit nil
### GetEndLocation

`func (o *CreateMileageResponse) GetEndLocation() string`

GetEndLocation returns the EndLocation field if non-nil, zero value otherwise.

### GetEndLocationOk

`func (o *CreateMileageResponse) GetEndLocationOk() (*string, bool)`

GetEndLocationOk returns a tuple with the EndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLocation

`func (o *CreateMileageResponse) SetEndLocation(v string)`

SetEndLocation sets EndLocation field to given value.

### HasEndLocation

`func (o *CreateMileageResponse) HasEndLocation() bool`

HasEndLocation returns a boolean if a field has been set.

### SetEndLocationNil

`func (o *CreateMileageResponse) SetEndLocationNil(b bool)`

 SetEndLocationNil sets the value for EndLocation to be an explicit nil

### UnsetEndLocation
`func (o *CreateMileageResponse) UnsetEndLocation()`

UnsetEndLocation ensures that no value is present for EndLocation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


