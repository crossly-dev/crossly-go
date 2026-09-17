# CreateCbxSubjectSpendPlanResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Legs** | [**[]CreateCbxSubjectSpendPlanResponseLegs**](CreateCbxSubjectSpendPlanResponseLegs.md) |  | 
**ShortfallBaseUnits** | **string** |  | 
**FullyFunded** | **bool** |  | 

## Methods

### NewCreateCbxSubjectSpendPlanResponse

`func NewCreateCbxSubjectSpendPlanResponse(legs []CreateCbxSubjectSpendPlanResponseLegs, shortfallBaseUnits string, fullyFunded bool, ) *CreateCbxSubjectSpendPlanResponse`

NewCreateCbxSubjectSpendPlanResponse instantiates a new CreateCbxSubjectSpendPlanResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateCbxSubjectSpendPlanResponseWithDefaults

`func NewCreateCbxSubjectSpendPlanResponseWithDefaults() *CreateCbxSubjectSpendPlanResponse`

NewCreateCbxSubjectSpendPlanResponseWithDefaults instantiates a new CreateCbxSubjectSpendPlanResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLegs

`func (o *CreateCbxSubjectSpendPlanResponse) GetLegs() []CreateCbxSubjectSpendPlanResponseLegs`

GetLegs returns the Legs field if non-nil, zero value otherwise.

### GetLegsOk

`func (o *CreateCbxSubjectSpendPlanResponse) GetLegsOk() (*[]CreateCbxSubjectSpendPlanResponseLegs, bool)`

GetLegsOk returns a tuple with the Legs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLegs

`func (o *CreateCbxSubjectSpendPlanResponse) SetLegs(v []CreateCbxSubjectSpendPlanResponseLegs)`

SetLegs sets Legs field to given value.


### GetShortfallBaseUnits

`func (o *CreateCbxSubjectSpendPlanResponse) GetShortfallBaseUnits() string`

GetShortfallBaseUnits returns the ShortfallBaseUnits field if non-nil, zero value otherwise.

### GetShortfallBaseUnitsOk

`func (o *CreateCbxSubjectSpendPlanResponse) GetShortfallBaseUnitsOk() (*string, bool)`

GetShortfallBaseUnitsOk returns a tuple with the ShortfallBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortfallBaseUnits

`func (o *CreateCbxSubjectSpendPlanResponse) SetShortfallBaseUnits(v string)`

SetShortfallBaseUnits sets ShortfallBaseUnits field to given value.


### GetFullyFunded

`func (o *CreateCbxSubjectSpendPlanResponse) GetFullyFunded() bool`

GetFullyFunded returns the FullyFunded field if non-nil, zero value otherwise.

### GetFullyFundedOk

`func (o *CreateCbxSubjectSpendPlanResponse) GetFullyFundedOk() (*bool, bool)`

GetFullyFundedOk returns a tuple with the FullyFunded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullyFunded

`func (o *CreateCbxSubjectSpendPlanResponse) SetFullyFunded(v bool)`

SetFullyFunded sets FullyFunded field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


