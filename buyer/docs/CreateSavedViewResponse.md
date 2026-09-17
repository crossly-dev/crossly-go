# CreateSavedViewResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**IsDefault** | **bool** |  | 
**SortOrder** | **float32** |  | 
**Resource** | **string** |  | 
**ViewMode** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateSavedViewResponse

`func NewCreateSavedViewResponse(id string, name string, createdAt time.Time, updatedAt time.Time, userId string, isDefault bool, sortOrder float32, resource string, ) *CreateSavedViewResponse`

NewCreateSavedViewResponse instantiates a new CreateSavedViewResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSavedViewResponseWithDefaults

`func NewCreateSavedViewResponseWithDefaults() *CreateSavedViewResponse`

NewCreateSavedViewResponseWithDefaults instantiates a new CreateSavedViewResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CreateSavedViewResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateSavedViewResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateSavedViewResponse) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *CreateSavedViewResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateSavedViewResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateSavedViewResponse) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *CreateSavedViewResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateSavedViewResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateSavedViewResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreateSavedViewResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateSavedViewResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateSavedViewResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *CreateSavedViewResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateSavedViewResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateSavedViewResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetIsDefault

`func (o *CreateSavedViewResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *CreateSavedViewResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *CreateSavedViewResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetSortOrder

`func (o *CreateSavedViewResponse) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *CreateSavedViewResponse) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *CreateSavedViewResponse) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.


### GetResource

`func (o *CreateSavedViewResponse) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *CreateSavedViewResponse) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *CreateSavedViewResponse) SetResource(v string)`

SetResource sets Resource field to given value.


### GetViewMode

`func (o *CreateSavedViewResponse) GetViewMode() string`

GetViewMode returns the ViewMode field if non-nil, zero value otherwise.

### GetViewModeOk

`func (o *CreateSavedViewResponse) GetViewModeOk() (*string, bool)`

GetViewModeOk returns a tuple with the ViewMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMode

`func (o *CreateSavedViewResponse) SetViewMode(v string)`

SetViewMode sets ViewMode field to given value.

### HasViewMode

`func (o *CreateSavedViewResponse) HasViewMode() bool`

HasViewMode returns a boolean if a field has been set.

### SetViewModeNil

`func (o *CreateSavedViewResponse) SetViewModeNil(b bool)`

 SetViewModeNil sets the value for ViewMode to be an explicit nil

### UnsetViewMode
`func (o *CreateSavedViewResponse) UnsetViewMode()`

UnsetViewMode ensures that no value is present for ViewMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


