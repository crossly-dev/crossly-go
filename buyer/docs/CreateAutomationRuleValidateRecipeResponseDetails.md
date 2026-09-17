# CreateAutomationRuleValidateRecipeResponseDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FormErrors** | **[]string** |  | 
**FieldErrors** | [**CreateAutomationRuleValidateRecipeResponseDetailsFieldErrors**](CreateAutomationRuleValidateRecipeResponseDetailsFieldErrors.md) |  | 

## Methods

### NewCreateAutomationRuleValidateRecipeResponseDetails

`func NewCreateAutomationRuleValidateRecipeResponseDetails(formErrors []string, fieldErrors CreateAutomationRuleValidateRecipeResponseDetailsFieldErrors, ) *CreateAutomationRuleValidateRecipeResponseDetails`

NewCreateAutomationRuleValidateRecipeResponseDetails instantiates a new CreateAutomationRuleValidateRecipeResponseDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateAutomationRuleValidateRecipeResponseDetailsWithDefaults

`func NewCreateAutomationRuleValidateRecipeResponseDetailsWithDefaults() *CreateAutomationRuleValidateRecipeResponseDetails`

NewCreateAutomationRuleValidateRecipeResponseDetailsWithDefaults instantiates a new CreateAutomationRuleValidateRecipeResponseDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFormErrors

`func (o *CreateAutomationRuleValidateRecipeResponseDetails) GetFormErrors() []string`

GetFormErrors returns the FormErrors field if non-nil, zero value otherwise.

### GetFormErrorsOk

`func (o *CreateAutomationRuleValidateRecipeResponseDetails) GetFormErrorsOk() (*[]string, bool)`

GetFormErrorsOk returns a tuple with the FormErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormErrors

`func (o *CreateAutomationRuleValidateRecipeResponseDetails) SetFormErrors(v []string)`

SetFormErrors sets FormErrors field to given value.


### GetFieldErrors

`func (o *CreateAutomationRuleValidateRecipeResponseDetails) GetFieldErrors() CreateAutomationRuleValidateRecipeResponseDetailsFieldErrors`

GetFieldErrors returns the FieldErrors field if non-nil, zero value otherwise.

### GetFieldErrorsOk

`func (o *CreateAutomationRuleValidateRecipeResponseDetails) GetFieldErrorsOk() (*CreateAutomationRuleValidateRecipeResponseDetailsFieldErrors, bool)`

GetFieldErrorsOk returns a tuple with the FieldErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldErrors

`func (o *CreateAutomationRuleValidateRecipeResponseDetails) SetFieldErrors(v CreateAutomationRuleValidateRecipeResponseDetailsFieldErrors)`

SetFieldErrors sets FieldErrors field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


