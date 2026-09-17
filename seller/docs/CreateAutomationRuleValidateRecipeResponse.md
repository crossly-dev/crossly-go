# CreateAutomationRuleValidateRecipeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Valid** | **bool** |  | 
**Error** | **string** |  | 
**Details** | [**CreateAutomationRuleValidateRecipeResponseDetails**](CreateAutomationRuleValidateRecipeResponseDetails.md) |  | 

## Methods

### NewCreateAutomationRuleValidateRecipeResponse

`func NewCreateAutomationRuleValidateRecipeResponse(valid bool, error_ string, details CreateAutomationRuleValidateRecipeResponseDetails, ) *CreateAutomationRuleValidateRecipeResponse`

NewCreateAutomationRuleValidateRecipeResponse instantiates a new CreateAutomationRuleValidateRecipeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutomationRuleValidateRecipeResponseWithDefaults

`func NewCreateAutomationRuleValidateRecipeResponseWithDefaults() *CreateAutomationRuleValidateRecipeResponse`

NewCreateAutomationRuleValidateRecipeResponseWithDefaults instantiates a new CreateAutomationRuleValidateRecipeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetValid

`func (o *CreateAutomationRuleValidateRecipeResponse) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *CreateAutomationRuleValidateRecipeResponse) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *CreateAutomationRuleValidateRecipeResponse) SetValid(v bool)`

SetValid sets Valid field to given value.


### GetError

`func (o *CreateAutomationRuleValidateRecipeResponse) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CreateAutomationRuleValidateRecipeResponse) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CreateAutomationRuleValidateRecipeResponse) SetError(v string)`

SetError sets Error field to given value.


### GetDetails

`func (o *CreateAutomationRuleValidateRecipeResponse) GetDetails() CreateAutomationRuleValidateRecipeResponseDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CreateAutomationRuleValidateRecipeResponse) GetDetailsOk() (*CreateAutomationRuleValidateRecipeResponseDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CreateAutomationRuleValidateRecipeResponse) SetDetails(v CreateAutomationRuleValidateRecipeResponseDetails)`

SetDetails sets Details field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


