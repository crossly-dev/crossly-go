# CreateBuyerActivityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ObservationId** | Pointer to **NullableString** |  | [optional] 
**Match** | Pointer to [**NullableCreateBuyerActivityResponseMatch**](CreateBuyerActivityResponseMatch.md) |  | [optional] 

## Methods

### NewCreateBuyerActivityResponse

`func NewCreateBuyerActivityResponse() *CreateBuyerActivityResponse`

NewCreateBuyerActivityResponse instantiates a new CreateBuyerActivityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerActivityResponseWithDefaults

`func NewCreateBuyerActivityResponseWithDefaults() *CreateBuyerActivityResponse`

NewCreateBuyerActivityResponseWithDefaults instantiates a new CreateBuyerActivityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetObservationId

`func (o *CreateBuyerActivityResponse) GetObservationId() string`

GetObservationId returns the ObservationId field if non-nil, zero value otherwise.

### GetObservationIdOk

`func (o *CreateBuyerActivityResponse) GetObservationIdOk() (*string, bool)`

GetObservationIdOk returns a tuple with the ObservationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservationId

`func (o *CreateBuyerActivityResponse) SetObservationId(v string)`

SetObservationId sets ObservationId field to given value.

### HasObservationId

`func (o *CreateBuyerActivityResponse) HasObservationId() bool`

HasObservationId returns a boolean if a field has been set.

### SetObservationIdNil

`func (o *CreateBuyerActivityResponse) SetObservationIdNil(b bool)`

 SetObservationIdNil sets the value for ObservationId to be an explicit nil

### UnsetObservationId
`func (o *CreateBuyerActivityResponse) UnsetObservationId()`

UnsetObservationId ensures that no value is present for ObservationId, not even an explicit nil
### GetMatch

`func (o *CreateBuyerActivityResponse) GetMatch() CreateBuyerActivityResponseMatch`

GetMatch returns the Match field if non-nil, zero value otherwise.

### GetMatchOk

`func (o *CreateBuyerActivityResponse) GetMatchOk() (*CreateBuyerActivityResponseMatch, bool)`

GetMatchOk returns a tuple with the Match field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatch

`func (o *CreateBuyerActivityResponse) SetMatch(v CreateBuyerActivityResponseMatch)`

SetMatch sets Match field to given value.

### HasMatch

`func (o *CreateBuyerActivityResponse) HasMatch() bool`

HasMatch returns a boolean if a field has been set.

### SetMatchNil

`func (o *CreateBuyerActivityResponse) SetMatchNil(b bool)`

 SetMatchNil sets the value for Match to be an explicit nil

### UnsetMatch
`func (o *CreateBuyerActivityResponse) UnsetMatch()`

UnsetMatch ensures that no value is present for Match, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


