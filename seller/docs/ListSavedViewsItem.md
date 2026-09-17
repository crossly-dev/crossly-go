# ListSavedViewsItem

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

### NewListSavedViewsItem

`func NewListSavedViewsItem(id string, name string, createdAt time.Time, updatedAt time.Time, userId string, isDefault bool, sortOrder float32, resource string, ) *ListSavedViewsItem`

NewListSavedViewsItem instantiates a new ListSavedViewsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSavedViewsItemWithDefaults

`func NewListSavedViewsItemWithDefaults() *ListSavedViewsItem`

NewListSavedViewsItemWithDefaults instantiates a new ListSavedViewsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSavedViewsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSavedViewsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSavedViewsItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListSavedViewsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSavedViewsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSavedViewsItem) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *ListSavedViewsItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListSavedViewsItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListSavedViewsItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListSavedViewsItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListSavedViewsItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListSavedViewsItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *ListSavedViewsItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListSavedViewsItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListSavedViewsItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetIsDefault

`func (o *ListSavedViewsItem) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *ListSavedViewsItem) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *ListSavedViewsItem) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetSortOrder

`func (o *ListSavedViewsItem) GetSortOrder() float32`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ListSavedViewsItem) GetSortOrderOk() (*float32, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ListSavedViewsItem) SetSortOrder(v float32)`

SetSortOrder sets SortOrder field to given value.


### GetResource

`func (o *ListSavedViewsItem) GetResource() string`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *ListSavedViewsItem) GetResourceOk() (*string, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *ListSavedViewsItem) SetResource(v string)`

SetResource sets Resource field to given value.


### GetViewMode

`func (o *ListSavedViewsItem) GetViewMode() string`

GetViewMode returns the ViewMode field if non-nil, zero value otherwise.

### GetViewModeOk

`func (o *ListSavedViewsItem) GetViewModeOk() (*string, bool)`

GetViewModeOk returns a tuple with the ViewMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetViewMode

`func (o *ListSavedViewsItem) SetViewMode(v string)`

SetViewMode sets ViewMode field to given value.

### HasViewMode

`func (o *ListSavedViewsItem) HasViewMode() bool`

HasViewMode returns a boolean if a field has been set.

### SetViewModeNil

`func (o *ListSavedViewsItem) SetViewModeNil(b bool)`

 SetViewModeNil sets the value for ViewMode to be an explicit nil

### UnsetViewMode
`func (o *ListSavedViewsItem) UnsetViewMode()`

UnsetViewMode ensures that no value is present for ViewMode, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


