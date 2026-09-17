# GetTaxonomyCategoryAspectResponseAspects

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Required** | **bool** |  | 
**DataType** | **string** | &#39;STRING&#39; | &#39;NUMBER&#39; | &#39;DATE&#39; — eBay&#39;s dataType per aspect. | 
**HasEnumValues** | **bool** | True when the aspect is selection-only (no free-text). | 
**EnumValues** | Pointer to **[]string** |  | [optional] 
**MaxLength** | Pointer to **NullableFloat32** | Helpful for client validation — max length when supplied. | [optional] 
**Cardinality** | **string** | Cardinality — SINGLE_VALUE / MULTIPLE_VALUES. | 

## Methods

### NewGetTaxonomyCategoryAspectResponseAspects

`func NewGetTaxonomyCategoryAspectResponseAspects(name string, required bool, dataType string, hasEnumValues bool, cardinality string, ) *GetTaxonomyCategoryAspectResponseAspects`

NewGetTaxonomyCategoryAspectResponseAspects instantiates a new GetTaxonomyCategoryAspectResponseAspects object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTaxonomyCategoryAspectResponseAspectsWithDefaults

`func NewGetTaxonomyCategoryAspectResponseAspectsWithDefaults() *GetTaxonomyCategoryAspectResponseAspects`

NewGetTaxonomyCategoryAspectResponseAspectsWithDefaults instantiates a new GetTaxonomyCategoryAspectResponseAspects object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *GetTaxonomyCategoryAspectResponseAspects) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetTaxonomyCategoryAspectResponseAspects) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetTaxonomyCategoryAspectResponseAspects) SetName(v string)`

SetName sets Name field to given value.


### GetRequired

`func (o *GetTaxonomyCategoryAspectResponseAspects) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *GetTaxonomyCategoryAspectResponseAspects) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *GetTaxonomyCategoryAspectResponseAspects) SetRequired(v bool)`

SetRequired sets Required field to given value.


### GetDataType

`func (o *GetTaxonomyCategoryAspectResponseAspects) GetDataType() string`

GetDataType returns the DataType field if non-nil, zero value otherwise.

### GetDataTypeOk

`func (o *GetTaxonomyCategoryAspectResponseAspects) GetDataTypeOk() (*string, bool)`

GetDataTypeOk returns a tuple with the DataType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataType

`func (o *GetTaxonomyCategoryAspectResponseAspects) SetDataType(v string)`

SetDataType sets DataType field to given value.


### GetHasEnumValues

`func (o *GetTaxonomyCategoryAspectResponseAspects) GetHasEnumValues() bool`

GetHasEnumValues returns the HasEnumValues field if non-nil, zero value otherwise.

### GetHasEnumValuesOk

`func (o *GetTaxonomyCategoryAspectResponseAspects) GetHasEnumValuesOk() (*bool, bool)`

GetHasEnumValuesOk returns a tuple with the HasEnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasEnumValues

`func (o *GetTaxonomyCategoryAspectResponseAspects) SetHasEnumValues(v bool)`

SetHasEnumValues sets HasEnumValues field to given value.


### GetEnumValues

`func (o *GetTaxonomyCategoryAspectResponseAspects) GetEnumValues() []string`

GetEnumValues returns the EnumValues field if non-nil, zero value otherwise.

### GetEnumValuesOk

`func (o *GetTaxonomyCategoryAspectResponseAspects) GetEnumValuesOk() (*[]string, bool)`

GetEnumValuesOk returns a tuple with the EnumValues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnumValues

`func (o *GetTaxonomyCategoryAspectResponseAspects) SetEnumValues(v []string)`

SetEnumValues sets EnumValues field to given value.

### HasEnumValues

`func (o *GetTaxonomyCategoryAspectResponseAspects) HasEnumValues() bool`

HasEnumValues returns a boolean if a field has been set.

### SetEnumValuesNil

`func (o *GetTaxonomyCategoryAspectResponseAspects) SetEnumValuesNil(b bool)`

 SetEnumValuesNil sets the value for EnumValues to be an explicit nil

### UnsetEnumValues
`func (o *GetTaxonomyCategoryAspectResponseAspects) UnsetEnumValues()`

UnsetEnumValues ensures that no value is present for EnumValues, not even an explicit nil
### GetMaxLength

`func (o *GetTaxonomyCategoryAspectResponseAspects) GetMaxLength() float32`

GetMaxLength returns the MaxLength field if non-nil, zero value otherwise.

### GetMaxLengthOk

`func (o *GetTaxonomyCategoryAspectResponseAspects) GetMaxLengthOk() (*float32, bool)`

GetMaxLengthOk returns a tuple with the MaxLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLength

`func (o *GetTaxonomyCategoryAspectResponseAspects) SetMaxLength(v float32)`

SetMaxLength sets MaxLength field to given value.

### HasMaxLength

`func (o *GetTaxonomyCategoryAspectResponseAspects) HasMaxLength() bool`

HasMaxLength returns a boolean if a field has been set.

### SetMaxLengthNil

`func (o *GetTaxonomyCategoryAspectResponseAspects) SetMaxLengthNil(b bool)`

 SetMaxLengthNil sets the value for MaxLength to be an explicit nil

### UnsetMaxLength
`func (o *GetTaxonomyCategoryAspectResponseAspects) UnsetMaxLength()`

UnsetMaxLength ensures that no value is present for MaxLength, not even an explicit nil
### GetCardinality

`func (o *GetTaxonomyCategoryAspectResponseAspects) GetCardinality() string`

GetCardinality returns the Cardinality field if non-nil, zero value otherwise.

### GetCardinalityOk

`func (o *GetTaxonomyCategoryAspectResponseAspects) GetCardinalityOk() (*string, bool)`

GetCardinalityOk returns a tuple with the Cardinality field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCardinality

`func (o *GetTaxonomyCategoryAspectResponseAspects) SetCardinality(v string)`

SetCardinality sets Cardinality field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


