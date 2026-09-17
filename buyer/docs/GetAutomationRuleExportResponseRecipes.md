# GetAutomationRuleExportResponseRecipes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Action** | [**GetAutomationRuleExportResponseAction**](GetAutomationRuleExportResponseAction.md) |  | 
**Trigger** | [**GetAutomationRuleExportResponseAction**](GetAutomationRuleExportResponseAction.md) |  | 
**SchemaVersion** | **float32** |  | 
**Condition** | Pointer to [**NullableGetAutomationRuleExportResponseCondition**](GetAutomationRuleExportResponseCondition.md) |  | [optional] 
**Metadata** | Pointer to [**NullableGetAutomationRuleExportResponseMetadata**](GetAutomationRuleExportResponseMetadata.md) |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Platforms** | Pointer to **[]string** |  | [optional] 
**Schema** | Pointer to **NullableString** | Optional reference to the public schema URL — purely cosmetic. | [optional] 

## Methods

### NewGetAutomationRuleExportResponseRecipes

`func NewGetAutomationRuleExportResponseRecipes(name string, action GetAutomationRuleExportResponseAction, trigger GetAutomationRuleExportResponseAction, schemaVersion float32, ) *GetAutomationRuleExportResponseRecipes`

NewGetAutomationRuleExportResponseRecipes instantiates a new GetAutomationRuleExportResponseRecipes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAutomationRuleExportResponseRecipesWithDefaults

`func NewGetAutomationRuleExportResponseRecipesWithDefaults() *GetAutomationRuleExportResponseRecipes`

NewGetAutomationRuleExportResponseRecipesWithDefaults instantiates a new GetAutomationRuleExportResponseRecipes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetAutomationRuleExportResponseRecipes) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAutomationRuleExportResponseRecipes) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAutomationRuleExportResponseRecipes) SetName(v string)`

SetName sets Name field to given value.


### GetAction

`func (o *GetAutomationRuleExportResponseRecipes) GetAction() GetAutomationRuleExportResponseAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetAutomationRuleExportResponseRecipes) GetActionOk() (*GetAutomationRuleExportResponseAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetAutomationRuleExportResponseRecipes) SetAction(v GetAutomationRuleExportResponseAction)`

SetAction sets Action field to given value.


### GetTrigger

`func (o *GetAutomationRuleExportResponseRecipes) GetTrigger() GetAutomationRuleExportResponseAction`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *GetAutomationRuleExportResponseRecipes) GetTriggerOk() (*GetAutomationRuleExportResponseAction, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *GetAutomationRuleExportResponseRecipes) SetTrigger(v GetAutomationRuleExportResponseAction)`

SetTrigger sets Trigger field to given value.


### GetSchemaVersion

`func (o *GetAutomationRuleExportResponseRecipes) GetSchemaVersion() float32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetAutomationRuleExportResponseRecipes) GetSchemaVersionOk() (*float32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetAutomationRuleExportResponseRecipes) SetSchemaVersion(v float32)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetCondition

`func (o *GetAutomationRuleExportResponseRecipes) GetCondition() GetAutomationRuleExportResponseCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GetAutomationRuleExportResponseRecipes) GetConditionOk() (*GetAutomationRuleExportResponseCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *GetAutomationRuleExportResponseRecipes) SetCondition(v GetAutomationRuleExportResponseCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *GetAutomationRuleExportResponseRecipes) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *GetAutomationRuleExportResponseRecipes) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *GetAutomationRuleExportResponseRecipes) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetMetadata

`func (o *GetAutomationRuleExportResponseRecipes) GetMetadata() GetAutomationRuleExportResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetAutomationRuleExportResponseRecipes) GetMetadataOk() (*GetAutomationRuleExportResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetAutomationRuleExportResponseRecipes) SetMetadata(v GetAutomationRuleExportResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetAutomationRuleExportResponseRecipes) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GetAutomationRuleExportResponseRecipes) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GetAutomationRuleExportResponseRecipes) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDescription

`func (o *GetAutomationRuleExportResponseRecipes) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAutomationRuleExportResponseRecipes) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAutomationRuleExportResponseRecipes) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAutomationRuleExportResponseRecipes) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetAutomationRuleExportResponseRecipes) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetAutomationRuleExportResponseRecipes) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPlatforms

`func (o *GetAutomationRuleExportResponseRecipes) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *GetAutomationRuleExportResponseRecipes) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *GetAutomationRuleExportResponseRecipes) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *GetAutomationRuleExportResponseRecipes) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### SetPlatformsNil

`func (o *GetAutomationRuleExportResponseRecipes) SetPlatformsNil(b bool)`

 SetPlatformsNil sets the value for Platforms to be an explicit nil

### UnsetPlatforms
`func (o *GetAutomationRuleExportResponseRecipes) UnsetPlatforms()`

UnsetPlatforms ensures that no value is present for Platforms, not even an explicit nil
### GetSchema

`func (o *GetAutomationRuleExportResponseRecipes) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *GetAutomationRuleExportResponseRecipes) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *GetAutomationRuleExportResponseRecipes) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *GetAutomationRuleExportResponseRecipes) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *GetAutomationRuleExportResponseRecipes) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *GetAutomationRuleExportResponseRecipes) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


