# GetAutomationRuleExportByIdResponse

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

### NewGetAutomationRuleExportByIdResponse

`func NewGetAutomationRuleExportByIdResponse(name string, action GetAutomationRuleExportResponseAction, trigger GetAutomationRuleExportResponseAction, schemaVersion float32, ) *GetAutomationRuleExportByIdResponse`

NewGetAutomationRuleExportByIdResponse instantiates a new GetAutomationRuleExportByIdResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAutomationRuleExportByIdResponseWithDefaults

`func NewGetAutomationRuleExportByIdResponseWithDefaults() *GetAutomationRuleExportByIdResponse`

NewGetAutomationRuleExportByIdResponseWithDefaults instantiates a new GetAutomationRuleExportByIdResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetAutomationRuleExportByIdResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetAutomationRuleExportByIdResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetAutomationRuleExportByIdResponse) SetName(v string)`

SetName sets Name field to given value.


### GetAction

`func (o *GetAutomationRuleExportByIdResponse) GetAction() GetAutomationRuleExportResponseAction`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *GetAutomationRuleExportByIdResponse) GetActionOk() (*GetAutomationRuleExportResponseAction, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *GetAutomationRuleExportByIdResponse) SetAction(v GetAutomationRuleExportResponseAction)`

SetAction sets Action field to given value.


### GetTrigger

`func (o *GetAutomationRuleExportByIdResponse) GetTrigger() GetAutomationRuleExportResponseAction`

GetTrigger returns the Trigger field if non-nil, zero value otherwise.

### GetTriggerOk

`func (o *GetAutomationRuleExportByIdResponse) GetTriggerOk() (*GetAutomationRuleExportResponseAction, bool)`

GetTriggerOk returns a tuple with the Trigger field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrigger

`func (o *GetAutomationRuleExportByIdResponse) SetTrigger(v GetAutomationRuleExportResponseAction)`

SetTrigger sets Trigger field to given value.


### GetSchemaVersion

`func (o *GetAutomationRuleExportByIdResponse) GetSchemaVersion() float32`

GetSchemaVersion returns the SchemaVersion field if non-nil, zero value otherwise.

### GetSchemaVersionOk

`func (o *GetAutomationRuleExportByIdResponse) GetSchemaVersionOk() (*float32, bool)`

GetSchemaVersionOk returns a tuple with the SchemaVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchemaVersion

`func (o *GetAutomationRuleExportByIdResponse) SetSchemaVersion(v float32)`

SetSchemaVersion sets SchemaVersion field to given value.


### GetCondition

`func (o *GetAutomationRuleExportByIdResponse) GetCondition() GetAutomationRuleExportResponseCondition`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GetAutomationRuleExportByIdResponse) GetConditionOk() (*GetAutomationRuleExportResponseCondition, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *GetAutomationRuleExportByIdResponse) SetCondition(v GetAutomationRuleExportResponseCondition)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *GetAutomationRuleExportByIdResponse) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *GetAutomationRuleExportByIdResponse) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *GetAutomationRuleExportByIdResponse) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetMetadata

`func (o *GetAutomationRuleExportByIdResponse) GetMetadata() GetAutomationRuleExportResponseMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *GetAutomationRuleExportByIdResponse) GetMetadataOk() (*GetAutomationRuleExportResponseMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *GetAutomationRuleExportByIdResponse) SetMetadata(v GetAutomationRuleExportResponseMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *GetAutomationRuleExportByIdResponse) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *GetAutomationRuleExportByIdResponse) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *GetAutomationRuleExportByIdResponse) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetDescription

`func (o *GetAutomationRuleExportByIdResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetAutomationRuleExportByIdResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetAutomationRuleExportByIdResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetAutomationRuleExportByIdResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetAutomationRuleExportByIdResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetAutomationRuleExportByIdResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPlatforms

`func (o *GetAutomationRuleExportByIdResponse) GetPlatforms() []string`

GetPlatforms returns the Platforms field if non-nil, zero value otherwise.

### GetPlatformsOk

`func (o *GetAutomationRuleExportByIdResponse) GetPlatformsOk() (*[]string, bool)`

GetPlatformsOk returns a tuple with the Platforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatforms

`func (o *GetAutomationRuleExportByIdResponse) SetPlatforms(v []string)`

SetPlatforms sets Platforms field to given value.

### HasPlatforms

`func (o *GetAutomationRuleExportByIdResponse) HasPlatforms() bool`

HasPlatforms returns a boolean if a field has been set.

### SetPlatformsNil

`func (o *GetAutomationRuleExportByIdResponse) SetPlatformsNil(b bool)`

 SetPlatformsNil sets the value for Platforms to be an explicit nil

### UnsetPlatforms
`func (o *GetAutomationRuleExportByIdResponse) UnsetPlatforms()`

UnsetPlatforms ensures that no value is present for Platforms, not even an explicit nil
### GetSchema

`func (o *GetAutomationRuleExportByIdResponse) GetSchema() string`

GetSchema returns the Schema field if non-nil, zero value otherwise.

### GetSchemaOk

`func (o *GetAutomationRuleExportByIdResponse) GetSchemaOk() (*string, bool)`

GetSchemaOk returns a tuple with the Schema field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchema

`func (o *GetAutomationRuleExportByIdResponse) SetSchema(v string)`

SetSchema sets Schema field to given value.

### HasSchema

`func (o *GetAutomationRuleExportByIdResponse) HasSchema() bool`

HasSchema returns a boolean if a field has been set.

### SetSchemaNil

`func (o *GetAutomationRuleExportByIdResponse) SetSchemaNil(b bool)`

 SetSchemaNil sets the value for Schema to be an explicit nil

### UnsetSchema
`func (o *GetAutomationRuleExportByIdResponse) UnsetSchema()`

UnsetSchema ensures that no value is present for Schema, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


