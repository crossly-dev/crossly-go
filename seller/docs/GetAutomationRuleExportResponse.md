# GetAutomationRuleExportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schema** | **string** |  | 
**SchemaVersion** | **float32** |  | 
**ExportedAt** | **string** |  | 
**Recipes** | [**[]GetAutomationRuleExportResponseRecipes**](GetAutomationRuleExportResponseRecipes.md) |  | 

## Methods

### NewGetAutomationRuleExportResponse

`func NewGetAutomationRuleExportResponse(schema string, schemaVersion float32, exportedAt string, recipes []GetAutomationRuleExportResponseRecipes, ) *GetAutomationRuleExportResponse`

NewGetAutomationRuleExportResponse instantiates a new GetAutomationRuleExportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAutomationRuleExportResponseWithDefaults

`func NewGetAutomationRuleExportResponseWithDefaults() *GetAutomationRuleExportResponse`

NewGetAutomationRuleExportResponseWithDefaults instantiates a new GetAutomationRuleExportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchema

`func (o *GetAutomationRuleExportResponse) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *GetAutomationRuleExportResponse) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *GetAutomationRuleExportResponse) SetSchema(v string)`

SetSchema sets Schema field to given value.


### GetSchemaVersion

`func (o *GetAutomationRuleExportResponse) GetSchemaVersion() float32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetAutomationRuleExportResponse) GetSchemaVersionOk() (*float32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetAutomationRuleExportResponse) SetSchemaVersion(v float32)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetExportedAt

`func (o *GetAutomationRuleExportResponse) GetExportedAt() string`

GetExportedAt returns the ExportedAt field if non-nil, zero value otherwise.

### GetExportedAtOk

`func (o *GetAutomationRuleExportResponse) GetExportedAtOk() (*string, bool)`

GetExportedAtOk returns a tuple with the ExportedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportedAt

`func (o *GetAutomationRuleExportResponse) SetExportedAt(v string)`

SetExportedAt sets ExportedAt field to given value.


### GetRecipes

`func (o *GetAutomationRuleExportResponse) GetRecipes() []GetAutomationRuleExportResponseRecipes`

GetRecipes returns the Recipes field if non-nil, zero value otherwise.

### GetRecipesOk

`func (o *GetAutomationRuleExportResponse) GetRecipesOk() (*[]GetAutomationRuleExportResponseRecipes, bool)`

GetRecipesOk returns a tuple with the Recipes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipes

`func (o *GetAutomationRuleExportResponse) SetRecipes(v []GetAutomationRuleExportResponseRecipes)`

SetRecipes sets Recipes field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


