# ListCbxBoostsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**TargetKind** | **string** |  | 
**TargetValue** | Pointer to **NullableString** |  | [optional] 
**BoostedRateBps** | **float32** |  | 
**BudgetCents** | **float32** |  | 
**SpentCents** | **float32** |  | 
**RemainingCents** | **float32** |  | 
**StartsAt** | **string** |  | 
**EndsAt** | **string** |  | 
**Status** | **string** |  | 

## Methods

### NewListCbxBoostsItem

`func NewListCbxBoostsItem(id string, name string, targetKind string, boostedRateBps float32, budgetCents float32, spentCents float32, remainingCents float32, startsAt string, endsAt string, status string, ) *ListCbxBoostsItem`

NewListCbxBoostsItem instantiates a new ListCbxBoostsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCbxBoostsItemWithDefaults

`func NewListCbxBoostsItemWithDefaults() *ListCbxBoostsItem`

NewListCbxBoostsItemWithDefaults instantiates a new ListCbxBoostsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCbxBoostsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCbxBoostsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCbxBoostsItem) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *ListCbxBoostsItem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCbxBoostsItem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCbxBoostsItem) SetName(v string)`

SetName sets Name field to given value.


### GetTargetKind

`func (o *ListCbxBoostsItem) GetTargetKind() string`

GetTargetKind returns the TargetKind field if non-nil, zero value otherwise.

### GetTargetKindOk

`func (o *ListCbxBoostsItem) GetTargetKindOk() (*string, bool)`

GetTargetKindOk returns a tuple with the TargetKind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetKind

`func (o *ListCbxBoostsItem) SetTargetKind(v string)`

SetTargetKind sets TargetKind field to given value.


### GetTargetValue

`func (o *ListCbxBoostsItem) GetTargetValue() string`

GetTargetValue returns the TargetValue field if non-nil, zero value otherwise.

### GetTargetValueOk

`func (o *ListCbxBoostsItem) GetTargetValueOk() (*string, bool)`

GetTargetValueOk returns a tuple with the TargetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetValue

`func (o *ListCbxBoostsItem) SetTargetValue(v string)`

SetTargetValue sets TargetValue field to given value.

### HasTargetValue

`func (o *ListCbxBoostsItem) HasTargetValue() bool`

HasTargetValue returns a boolean if a field has been set.

### SetTargetValueNil

`func (o *ListCbxBoostsItem) SetTargetValueNil(b bool)`

 SetTargetValueNil sets the value for TargetValue to be an explicit nil

### UnsetTargetValue
`func (o *ListCbxBoostsItem) UnsetTargetValue()`

UnsetTargetValue ensures that no value is present for TargetValue, not even an explicit nil
### GetBoostedRateBps

`func (o *ListCbxBoostsItem) GetBoostedRateBps() float32`

GetBoostedRateBps returns the BoostedRateBps field if non-nil, zero value otherwise.

### GetBoostedRateBpsOk

`func (o *ListCbxBoostsItem) GetBoostedRateBpsOk() (*float32, bool)`

GetBoostedRateBpsOk returns a tuple with the BoostedRateBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBoostedRateBps

`func (o *ListCbxBoostsItem) SetBoostedRateBps(v float32)`

SetBoostedRateBps sets BoostedRateBps field to given value.


### GetBudgetCents

`func (o *ListCbxBoostsItem) GetBudgetCents() float32`

GetBudgetCents returns the BudgetCents field if non-nil, zero value otherwise.

### GetBudgetCentsOk

`func (o *ListCbxBoostsItem) GetBudgetCentsOk() (*float32, bool)`

GetBudgetCentsOk returns a tuple with the BudgetCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBudgetCents

`func (o *ListCbxBoostsItem) SetBudgetCents(v float32)`

SetBudgetCents sets BudgetCents field to given value.


### GetSpentCents

`func (o *ListCbxBoostsItem) GetSpentCents() float32`

GetSpentCents returns the SpentCents field if non-nil, zero value otherwise.

### GetSpentCentsOk

`func (o *ListCbxBoostsItem) GetSpentCentsOk() (*float32, bool)`

GetSpentCentsOk returns a tuple with the SpentCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpentCents

`func (o *ListCbxBoostsItem) SetSpentCents(v float32)`

SetSpentCents sets SpentCents field to given value.


### GetRemainingCents

`func (o *ListCbxBoostsItem) GetRemainingCents() float32`

GetRemainingCents returns the RemainingCents field if non-nil, zero value otherwise.

### GetRemainingCentsOk

`func (o *ListCbxBoostsItem) GetRemainingCentsOk() (*float32, bool)`

GetRemainingCentsOk returns a tuple with the RemainingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemainingCents

`func (o *ListCbxBoostsItem) SetRemainingCents(v float32)`

SetRemainingCents sets RemainingCents field to given value.


### GetStartsAt

`func (o *ListCbxBoostsItem) GetStartsAt() string`

GetStartsAt returns the StartsAt field if non-nil, zero value otherwise.

### GetStartsAtOk

`func (o *ListCbxBoostsItem) GetStartsAtOk() (*string, bool)`

GetStartsAtOk returns a tuple with the StartsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartsAt

`func (o *ListCbxBoostsItem) SetStartsAt(v string)`

SetStartsAt sets StartsAt field to given value.


### GetEndsAt

`func (o *ListCbxBoostsItem) GetEndsAt() string`

GetEndsAt returns the EndsAt field if non-nil, zero value otherwise.

### GetEndsAtOk

`func (o *ListCbxBoostsItem) GetEndsAtOk() (*string, bool)`

GetEndsAtOk returns a tuple with the EndsAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndsAt

`func (o *ListCbxBoostsItem) SetEndsAt(v string)`

SetEndsAt sets EndsAt field to given value.


### GetStatus

`func (o *ListCbxBoostsItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListCbxBoostsItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListCbxBoostsItem) SetStatus(v string)`

SetStatus sets Status field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


