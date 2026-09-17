# CreateCbxSubjectGrantResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GrantId** | Pointer to **NullableString** |  | [optional] 
**Duplicate** | **bool** |  | 

## Methods

### NewCreateCbxSubjectGrantResponse

`func NewCreateCbxSubjectGrantResponse(duplicate bool, ) *CreateCbxSubjectGrantResponse`

NewCreateCbxSubjectGrantResponse instantiates a new CreateCbxSubjectGrantResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCbxSubjectGrantResponseWithDefaults

`func NewCreateCbxSubjectGrantResponseWithDefaults() *CreateCbxSubjectGrantResponse`

NewCreateCbxSubjectGrantResponseWithDefaults instantiates a new CreateCbxSubjectGrantResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGrantId

`func (o *CreateCbxSubjectGrantResponse) GetGrantId() string`

GetGrantId returns the GrantId field if non-nil, zero value otherwise.

### GetGrantIdOk

`func (o *CreateCbxSubjectGrantResponse) GetGrantIdOk() (*string, bool)`

GetGrantIdOk returns a tuple with the GrantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrantId

`func (o *CreateCbxSubjectGrantResponse) SetGrantId(v string)`

SetGrantId sets GrantId field to given value.

### HasGrantId

`func (o *CreateCbxSubjectGrantResponse) HasGrantId() bool`

HasGrantId returns a boolean if a field has been set.

### SetGrantIdNil

`func (o *CreateCbxSubjectGrantResponse) SetGrantIdNil(b bool)`

 SetGrantIdNil sets the value for GrantId to be an explicit nil

### UnsetGrantId
`func (o *CreateCbxSubjectGrantResponse) UnsetGrantId()`

UnsetGrantId ensures that no value is present for GrantId, not even an explicit nil
### GetDuplicate

`func (o *CreateCbxSubjectGrantResponse) GetDuplicate() bool`

GetDuplicate returns the Duplicate field if non-nil, zero value otherwise.

### GetDuplicateOk

`func (o *CreateCbxSubjectGrantResponse) GetDuplicateOk() (*bool, bool)`

GetDuplicateOk returns a tuple with the Duplicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicate

`func (o *CreateCbxSubjectGrantResponse) SetDuplicate(v bool)`

SetDuplicate sets Duplicate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


