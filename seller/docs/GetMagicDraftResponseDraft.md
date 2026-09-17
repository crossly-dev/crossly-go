# GetMagicDraftResponseDraft

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**ListingId** | Pointer to **NullableString** |  | [optional] 
**RunId** | **string** |  | 
**SelectedMatchKeys** | Pointer to **[]string** |  | [optional] 
**ChosenPriceCents** | Pointer to **NullableFloat32** |  | [optional] 
**ChosenQuantity** | Pointer to **NullableFloat32** |  | [optional] 
**ChosenCondition** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetMagicDraftResponseDraft

`func NewGetMagicDraftResponseDraft(id string, createdAt time.Time, userId string, runId string, ) *GetMagicDraftResponseDraft`

NewGetMagicDraftResponseDraft instantiates a new GetMagicDraftResponseDraft object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMagicDraftResponseDraftWithDefaults

`func NewGetMagicDraftResponseDraftWithDefaults() *GetMagicDraftResponseDraft`

NewGetMagicDraftResponseDraftWithDefaults instantiates a new GetMagicDraftResponseDraft object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMagicDraftResponseDraft) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMagicDraftResponseDraft) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMagicDraftResponseDraft) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetMagicDraftResponseDraft) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetMagicDraftResponseDraft) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetMagicDraftResponseDraft) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *GetMagicDraftResponseDraft) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetMagicDraftResponseDraft) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetMagicDraftResponseDraft) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetListingId

`func (o *GetMagicDraftResponseDraft) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *GetMagicDraftResponseDraft) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *GetMagicDraftResponseDraft) SetListingId(v string)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *GetMagicDraftResponseDraft) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### SetListingIdNil

`func (o *GetMagicDraftResponseDraft) SetListingIdNil(b bool)`

 SetListingIdNil sets the value for ListingId to be an explicit nil

### UnsetListingId
`func (o *GetMagicDraftResponseDraft) UnsetListingId()`

UnsetListingId ensures that no value is present for ListingId, not even an explicit nil
### GetRunId

`func (o *GetMagicDraftResponseDraft) GetRunId() string`

GetRunId returns the RunId field if non-nil, zero value otherwise.

### GetRunIdOk

`func (o *GetMagicDraftResponseDraft) GetRunIdOk() (*string, bool)`

GetRunIdOk returns a tuple with the RunId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRunId

`func (o *GetMagicDraftResponseDraft) SetRunId(v string)`

SetRunId sets RunId field to given value.


### GetSelectedMatchKeys

`func (o *GetMagicDraftResponseDraft) GetSelectedMatchKeys() []string`

GetSelectedMatchKeys returns the SelectedMatchKeys field if non-nil, zero value otherwise.

### GetSelectedMatchKeysOk

`func (o *GetMagicDraftResponseDraft) GetSelectedMatchKeysOk() (*[]string, bool)`

GetSelectedMatchKeysOk returns a tuple with the SelectedMatchKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectedMatchKeys

`func (o *GetMagicDraftResponseDraft) SetSelectedMatchKeys(v []string)`

SetSelectedMatchKeys sets SelectedMatchKeys field to given value.

### HasSelectedMatchKeys

`func (o *GetMagicDraftResponseDraft) HasSelectedMatchKeys() bool`

HasSelectedMatchKeys returns a boolean if a field has been set.

### SetSelectedMatchKeysNil

`func (o *GetMagicDraftResponseDraft) SetSelectedMatchKeysNil(b bool)`

 SetSelectedMatchKeysNil sets the value for SelectedMatchKeys to be an explicit nil

### UnsetSelectedMatchKeys
`func (o *GetMagicDraftResponseDraft) UnsetSelectedMatchKeys()`

UnsetSelectedMatchKeys ensures that no value is present for SelectedMatchKeys, not even an explicit nil
### GetChosenPriceCents

`func (o *GetMagicDraftResponseDraft) GetChosenPriceCents() float32`

GetChosenPriceCents returns the ChosenPriceCents field if non-nil, zero value otherwise.

### GetChosenPriceCentsOk

`func (o *GetMagicDraftResponseDraft) GetChosenPriceCentsOk() (*float32, bool)`

GetChosenPriceCentsOk returns a tuple with the ChosenPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChosenPriceCents

`func (o *GetMagicDraftResponseDraft) SetChosenPriceCents(v float32)`

SetChosenPriceCents sets ChosenPriceCents field to given value.

### HasChosenPriceCents

`func (o *GetMagicDraftResponseDraft) HasChosenPriceCents() bool`

HasChosenPriceCents returns a boolean if a field has been set.

### SetChosenPriceCentsNil

`func (o *GetMagicDraftResponseDraft) SetChosenPriceCentsNil(b bool)`

 SetChosenPriceCentsNil sets the value for ChosenPriceCents to be an explicit nil

### UnsetChosenPriceCents
`func (o *GetMagicDraftResponseDraft) UnsetChosenPriceCents()`

UnsetChosenPriceCents ensures that no value is present for ChosenPriceCents, not even an explicit nil
### GetChosenQuantity

`func (o *GetMagicDraftResponseDraft) GetChosenQuantity() float32`

GetChosenQuantity returns the ChosenQuantity field if non-nil, zero value otherwise.

### GetChosenQuantityOk

`func (o *GetMagicDraftResponseDraft) GetChosenQuantityOk() (*float32, bool)`

GetChosenQuantityOk returns a tuple with the ChosenQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChosenQuantity

`func (o *GetMagicDraftResponseDraft) SetChosenQuantity(v float32)`

SetChosenQuantity sets ChosenQuantity field to given value.

### HasChosenQuantity

`func (o *GetMagicDraftResponseDraft) HasChosenQuantity() bool`

HasChosenQuantity returns a boolean if a field has been set.

### SetChosenQuantityNil

`func (o *GetMagicDraftResponseDraft) SetChosenQuantityNil(b bool)`

 SetChosenQuantityNil sets the value for ChosenQuantity to be an explicit nil

### UnsetChosenQuantity
`func (o *GetMagicDraftResponseDraft) UnsetChosenQuantity()`

UnsetChosenQuantity ensures that no value is present for ChosenQuantity, not even an explicit nil
### GetChosenCondition

`func (o *GetMagicDraftResponseDraft) GetChosenCondition() string`

GetChosenCondition returns the ChosenCondition field if non-nil, zero value otherwise.

### GetChosenConditionOk

`func (o *GetMagicDraftResponseDraft) GetChosenConditionOk() (*string, bool)`

GetChosenConditionOk returns a tuple with the ChosenCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChosenCondition

`func (o *GetMagicDraftResponseDraft) SetChosenCondition(v string)`

SetChosenCondition sets ChosenCondition field to given value.

### HasChosenCondition

`func (o *GetMagicDraftResponseDraft) HasChosenCondition() bool`

HasChosenCondition returns a boolean if a field has been set.

### SetChosenConditionNil

`func (o *GetMagicDraftResponseDraft) SetChosenConditionNil(b bool)`

 SetChosenConditionNil sets the value for ChosenCondition to be an explicit nil

### UnsetChosenCondition
`func (o *GetMagicDraftResponseDraft) UnsetChosenCondition()`

UnsetChosenCondition ensures that no value is present for ChosenCondition, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


