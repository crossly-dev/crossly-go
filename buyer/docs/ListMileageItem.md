# ListMileageItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Date** | **string** |  | 
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Miles** | **string** |  | 
**Purpose** | **string** |  | 
**StartLocation** | Pointer to **NullableString** |  | [optional] 
**EndLocation** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewListMileageItem

`func NewListMileageItem(date string, id string, createdAt time.Time, userId string, miles string, purpose string, ) *ListMileageItem`

NewListMileageItem instantiates a new ListMileageItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMileageItemWithDefaults

`func NewListMileageItemWithDefaults() *ListMileageItem`

NewListMileageItemWithDefaults instantiates a new ListMileageItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDate

`func (o *ListMileageItem) GetDate() string`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ListMileageItem) GetDateOk() (*string, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ListMileageItem) SetDate(v string)`

SetDate sets Date field to given value.


### GetId

`func (o *ListMileageItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListMileageItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListMileageItem) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListMileageItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListMileageItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListMileageItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *ListMileageItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListMileageItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListMileageItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetNotes

`func (o *ListMileageItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ListMileageItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ListMileageItem) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ListMileageItem) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ListMileageItem) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ListMileageItem) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetMiles

`func (o *ListMileageItem) GetMiles() string`

GetMiles returns the Miles field if non-nil, zero value otherwise.

### GetMilesOk

`func (o *ListMileageItem) GetMilesOk() (*string, bool)`

GetMilesOk returns a tuple with the Miles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMiles

`func (o *ListMileageItem) SetMiles(v string)`

SetMiles sets Miles field to given value.


### GetPurpose

`func (o *ListMileageItem) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *ListMileageItem) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *ListMileageItem) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.


### GetStartLocation

`func (o *ListMileageItem) GetStartLocation() string`

GetStartLocation returns the StartLocation field if non-nil, zero value otherwise.

### GetStartLocationOk

`func (o *ListMileageItem) GetStartLocationOk() (*string, bool)`

GetStartLocationOk returns a tuple with the StartLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartLocation

`func (o *ListMileageItem) SetStartLocation(v string)`

SetStartLocation sets StartLocation field to given value.

### HasStartLocation

`func (o *ListMileageItem) HasStartLocation() bool`

HasStartLocation returns a boolean if a field has been set.

### SetStartLocationNil

`func (o *ListMileageItem) SetStartLocationNil(b bool)`

 SetStartLocationNil sets the value for StartLocation to be an explicit nil

### UnsetStartLocation
`func (o *ListMileageItem) UnsetStartLocation()`

UnsetStartLocation ensures that no value is present for StartLocation, not even an explicit nil
### GetEndLocation

`func (o *ListMileageItem) GetEndLocation() string`

GetEndLocation returns the EndLocation field if non-nil, zero value otherwise.

### GetEndLocationOk

`func (o *ListMileageItem) GetEndLocationOk() (*string, bool)`

GetEndLocationOk returns a tuple with the EndLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndLocation

`func (o *ListMileageItem) SetEndLocation(v string)`

SetEndLocation sets EndLocation field to given value.

### HasEndLocation

`func (o *ListMileageItem) HasEndLocation() bool`

HasEndLocation returns a boolean if a field has been set.

### SetEndLocationNil

`func (o *ListMileageItem) SetEndLocationNil(b bool)`

 SetEndLocationNil sets the value for EndLocation to be an explicit nil

### UnsetEndLocation
`func (o *ListMileageItem) UnsetEndLocation()`

UnsetEndLocation ensures that no value is present for EndLocation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


