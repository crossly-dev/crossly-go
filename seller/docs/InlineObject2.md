# InlineObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** |  | 
**CounterCents** | Pointer to **int32** | Required when action is \&quot;counter\&quot;. Your asking price for the set. | [optional] 

## Methods

### NewInlineObject2

`func NewInlineObject2(action string, ) *InlineObject2`

NewInlineObject2 instantiates a new InlineObject2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInlineObject2WithDefaults

`func NewInlineObject2WithDefaults() *InlineObject2`

NewInlineObject2WithDefaults instantiates a new InlineObject2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *InlineObject2) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *InlineObject2) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *InlineObject2) SetAction(v string)`

SetAction sets Action field to given value.


### GetCounterCents

`func (o *InlineObject2) GetCounterCents() int32`

GetCounterCents returns the CounterCents field if non-nil, zero value otherwise.

### GetCounterCentsOk

`func (o *InlineObject2) GetCounterCentsOk() (*int32, bool)`

GetCounterCentsOk returns a tuple with the CounterCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounterCents

`func (o *InlineObject2) SetCounterCents(v int32)`

SetCounterCents sets CounterCents field to given value.

### HasCounterCents

`func (o *InlineObject2) HasCounterCents() bool`

HasCounterCents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


