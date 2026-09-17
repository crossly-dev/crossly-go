# UpdatePolicyPresetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Kind** | **string** | &#39;return&#39; | &#39;shipping&#39; | &#39;payment&#39; — validated at the route layer. | 
**Name** | **string** |  | 
**IsDefault** | **bool** | Auto-selected on new listings. At most one true per (user, kind). | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewUpdatePolicyPresetResponse

`func NewUpdatePolicyPresetResponse(id string, userId string, kind string, name string, isDefault bool, createdAt time.Time, updatedAt time.Time, ) *UpdatePolicyPresetResponse`

NewUpdatePolicyPresetResponse instantiates a new UpdatePolicyPresetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePolicyPresetResponseWithDefaults

`func NewUpdatePolicyPresetResponseWithDefaults() *UpdatePolicyPresetResponse`

NewUpdatePolicyPresetResponseWithDefaults instantiates a new UpdatePolicyPresetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdatePolicyPresetResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatePolicyPresetResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatePolicyPresetResponse) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *UpdatePolicyPresetResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdatePolicyPresetResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdatePolicyPresetResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetKind

`func (o *UpdatePolicyPresetResponse) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *UpdatePolicyPresetResponse) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *UpdatePolicyPresetResponse) SetKind(v string)`

SetKind sets Kind field to given value.


### GetName

`func (o *UpdatePolicyPresetResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePolicyPresetResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePolicyPresetResponse) SetName(v string)`

SetName sets Name field to given value.


### GetIsDefault

`func (o *UpdatePolicyPresetResponse) GetIsDefault() bool`

GetIsDefault returns the IsDefault field if non-nil, zero value otherwise.

### GetIsDefaultOk

`func (o *UpdatePolicyPresetResponse) GetIsDefaultOk() (*bool, bool)`

GetIsDefaultOk returns a tuple with the IsDefault field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDefault

`func (o *UpdatePolicyPresetResponse) SetIsDefault(v bool)`

SetIsDefault sets IsDefault field to given value.


### GetCreatedAt

`func (o *UpdatePolicyPresetResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdatePolicyPresetResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdatePolicyPresetResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UpdatePolicyPresetResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdatePolicyPresetResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdatePolicyPresetResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


