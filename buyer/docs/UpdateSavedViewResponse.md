# UpdateSavedViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Resource** | **string** |  | 
**Name** | **string** |  | 
**ViewMode** | Pointer to **NullableString** |  | [optional] 
**IsDefault** | **bool** |  | 
**SortOrder** | **float32** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewUpdateSavedViewResponse

`func NewUpdateSavedViewResponse(id string, userId string, resource string, name string, isDefault bool, sortOrder float32, createdAt time.Time, updatedAt time.Time, ) *UpdateSavedViewResponse`

NewUpdateSavedViewResponse instantiates a new UpdateSavedViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSavedViewResponseWithDefaults

`func NewUpdateSavedViewResponseWithDefaults() *UpdateSavedViewResponse`

NewUpdateSavedViewResponseWithDefaults instantiates a new UpdateSavedViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateSavedViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateSavedViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateSavedViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *UpdateSavedViewResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateSavedViewResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateSavedViewResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetResource

`func (o *UpdateSavedViewResponse) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *UpdateSavedViewResponse) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *UpdateSavedViewResponse) SetResource(v string)`

SetResource sets Resource field to given value.


### GetName

`func (o *UpdateSavedViewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSavedViewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSavedViewResponse) SetName(v string)`

SetName sets Name field to given value.


### GetViewMode

`func (o *UpdateSavedViewResponse) GetViewMode() string`

GetViewMode returns the ViewMode field if non-nil, zero value otherwise.

### GetViewModeOk

`func (o *UpdateSavedViewResponse) GetViewModeOk() (*string, bool)`

GetViewModeOk returns a tuple with the ViewMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMode

`func (o *UpdateSavedViewResponse) SetViewMode(v string)`

SetViewMode sets ViewMode field to given value.

### HasViewMode

`func (o *UpdateSavedViewResponse) HasViewMode() bool`

HasViewMode returns a boolean if a field has been set.

### SetViewModeNil

`func (o *UpdateSavedViewResponse) SetViewModeNil(b bool)`

 SetViewModeNil sets the value for ViewMode to be an explicit nil

### UnsetViewMode
`func (o *UpdateSavedViewResponse) UnsetViewMode()`

UnsetViewMode ensures that no value is present for ViewMode, not even an explicit nil
### GetIsDefault

`func (o *UpdateSavedViewResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UpdateSavedViewResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UpdateSavedViewResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetSortOrder

`func (o *UpdateSavedViewResponse) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *UpdateSavedViewResponse) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *UpdateSavedViewResponse) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.


### GetCreatedAt

`func (o *UpdateSavedViewResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateSavedViewResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateSavedViewResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UpdateSavedViewResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdateSavedViewResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdateSavedViewResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


