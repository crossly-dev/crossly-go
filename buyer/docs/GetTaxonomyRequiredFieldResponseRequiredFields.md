# GetTaxonomyRequiredFieldResponseRequiredFields

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** |  | 
**Label** | **string** |  | 
**Master** | **bool** |  | 
**Required** | **bool** |  | 
**DataType** | **string** |  | 
**HasEnumValues** | **bool** |  | 
**EnumValues** | Pointer to **[]string** |  | [optional] 
**Hint** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetTaxonomyRequiredFieldResponseRequiredFields

`func NewGetTaxonomyRequiredFieldResponseRequiredFields(key string, label string, master bool, required bool, dataType string, hasEnumValues bool, ) *GetTaxonomyRequiredFieldResponseRequiredFields`

NewGetTaxonomyRequiredFieldResponseRequiredFields instantiates a new GetTaxonomyRequiredFieldResponseRequiredFields object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaxonomyRequiredFieldResponseRequiredFieldsWithDefaults

`func NewGetTaxonomyRequiredFieldResponseRequiredFieldsWithDefaults() *GetTaxonomyRequiredFieldResponseRequiredFields`

NewGetTaxonomyRequiredFieldResponseRequiredFieldsWithDefaults instantiates a new GetTaxonomyRequiredFieldResponseRequiredFields object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMaster

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetMaster() bool`

GetMaster returns the Master field if non-nil, zero value otherwise.

### GetMasterOk

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetMasterOk() (*bool, bool)`

GetMasterOk returns a tuple with the Master field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaster

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) SetMaster(v bool)`

SetMaster sets Master field to given value.


### GetRequired

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetDataType

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetHasEnumValues

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetHasEnumValues() bool`

GetHasEnumValues returns the HasEnumValues field if non-nil, zero value otherwise.

### GetHasEnumValuesOk

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetHasEnumValuesOk() (*bool, bool)`

GetHasEnumValuesOk returns a tuple with the HasEnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEnumValues

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) SetHasEnumValues(v bool)`

SetHasEnumValues sets HasEnumValues field to given value.


### GetEnumValues

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetEnumValues() []string`

GetEnumValues returns the EnumValues field if non-nil, zero value otherwise.

### GetEnumValuesOk

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetEnumValuesOk() (*[]string, bool)`

GetEnumValuesOk returns a tuple with the EnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValues

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) SetEnumValues(v []string)`

SetEnumValues sets EnumValues field to given value.

### HasEnumValues

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) HasEnumValues() bool`

HasEnumValues returns a boolean if a field has been set.

### SetEnumValuesNil

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) SetEnumValuesNil(b bool)`

 SetEnumValuesNil sets the value for EnumValues to be an explicit nil

### UnsetEnumValues
`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) UnsetEnumValues()`

UnsetEnumValues ensures that no value is present for EnumValues, not even an explicit nil
### GetHint

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetHint() string`

GetHint returns the Hint field if non-nil, zero value otherwise.

### GetHintOk

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) GetHintOk() (*string, bool)`

GetHintOk returns a tuple with the Hint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHint

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) SetHint(v string)`

SetHint sets Hint field to given value.

### HasHint

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) HasHint() bool`

HasHint returns a boolean if a field has been set.

### SetHintNil

`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) SetHintNil(b bool)`

 SetHintNil sets the value for Hint to be an explicit nil

### UnsetHint
`func (o *GetTaxonomyRequiredFieldResponseRequiredFields) UnsetHint()`

UnsetHint ensures that no value is present for Hint, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


