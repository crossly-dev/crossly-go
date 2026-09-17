# CreateMagicScanSynthesizeResponsePayloadGrading

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GraderSlug** | **string** |  | 
**Grade** | **string** | As the grader writes it: &#39;10&#39;, &#39;9.8&#39;, &#39;MS-65&#39;. | 
**Qualifier** | Pointer to **NullableString** |  | [optional] 
**SealGrade** | Pointer to **NullableString** | Second axis on a &#x60;dual&#x60; scale — WATA&#39;s seal grade. | [optional] 
**CertNumber** | Pointer to **NullableString** |  | [optional] 
**GradeKey** | **string** | Derived. Never assign by hand — call &#x60;gradeKey()&#x60;. | 
**VerifiedAt** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateMagicScanSynthesizeResponsePayloadGrading

`func NewCreateMagicScanSynthesizeResponsePayloadGrading(graderSlug string, grade string, gradeKey string, ) *CreateMagicScanSynthesizeResponsePayloadGrading`

NewCreateMagicScanSynthesizeResponsePayloadGrading instantiates a new CreateMagicScanSynthesizeResponsePayloadGrading object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMagicScanSynthesizeResponsePayloadGradingWithDefaults

`func NewCreateMagicScanSynthesizeResponsePayloadGradingWithDefaults() *CreateMagicScanSynthesizeResponsePayloadGrading`

NewCreateMagicScanSynthesizeResponsePayloadGradingWithDefaults instantiates a new CreateMagicScanSynthesizeResponsePayloadGrading object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGraderSlug

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) GetGraderSlug() string`

GetGraderSlug returns the GraderSlug field if non-nil, zero value otherwise.

### GetGraderSlugOk

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) GetGraderSlugOk() (*string, bool)`

GetGraderSlugOk returns a tuple with the GraderSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGraderSlug

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) SetGraderSlug(v string)`

SetGraderSlug sets GraderSlug field to given value.


### GetGrade

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) GetGrade() string`

GetGrade returns the Grade field if non-nil, zero value otherwise.

### GetGradeOk

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) GetGradeOk() (*string, bool)`

GetGradeOk returns a tuple with the Grade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrade

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) SetGrade(v string)`

SetGrade sets Grade field to given value.


### GetQualifier

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) GetQualifier() string`

GetQualifier returns the Qualifier field if non-nil, zero value otherwise.

### GetQualifierOk

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) GetQualifierOk() (*string, bool)`

GetQualifierOk returns a tuple with the Qualifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualifier

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) SetQualifier(v string)`

SetQualifier sets Qualifier field to given value.

### HasQualifier

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) HasQualifier() bool`

HasQualifier returns a boolean if a field has been set.

### SetQualifierNil

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) SetQualifierNil(b bool)`

 SetQualifierNil sets the value for Qualifier to be an explicit nil

### UnsetQualifier
`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) UnsetQualifier()`

UnsetQualifier ensures that no value is present for Qualifier, not even an explicit nil
### GetSealGrade

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) GetSealGrade() string`

GetSealGrade returns the SealGrade field if non-nil, zero value otherwise.

### GetSealGradeOk

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) GetSealGradeOk() (*string, bool)`

GetSealGradeOk returns a tuple with the SealGrade field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSealGrade

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) SetSealGrade(v string)`

SetSealGrade sets SealGrade field to given value.

### HasSealGrade

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) HasSealGrade() bool`

HasSealGrade returns a boolean if a field has been set.

### SetSealGradeNil

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) SetSealGradeNil(b bool)`

 SetSealGradeNil sets the value for SealGrade to be an explicit nil

### UnsetSealGrade
`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) UnsetSealGrade()`

UnsetSealGrade ensures that no value is present for SealGrade, not even an explicit nil
### GetCertNumber

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) GetCertNumber() string`

GetCertNumber returns the CertNumber field if non-nil, zero value otherwise.

### GetCertNumberOk

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) GetCertNumberOk() (*string, bool)`

GetCertNumberOk returns a tuple with the CertNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertNumber

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) SetCertNumber(v string)`

SetCertNumber sets CertNumber field to given value.

### HasCertNumber

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) HasCertNumber() bool`

HasCertNumber returns a boolean if a field has been set.

### SetCertNumberNil

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) SetCertNumberNil(b bool)`

 SetCertNumberNil sets the value for CertNumber to be an explicit nil

### UnsetCertNumber
`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) UnsetCertNumber()`

UnsetCertNumber ensures that no value is present for CertNumber, not even an explicit nil
### GetGradeKey

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) GetGradeKey() string`

GetGradeKey returns the GradeKey field if non-nil, zero value otherwise.

### GetGradeKeyOk

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) GetGradeKeyOk() (*string, bool)`

GetGradeKeyOk returns a tuple with the GradeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeKey

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) SetGradeKey(v string)`

SetGradeKey sets GradeKey field to given value.


### GetVerifiedAt

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) GetVerifiedAt() string`

GetVerifiedAt returns the VerifiedAt field if non-nil, zero value otherwise.

### GetVerifiedAtOk

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) GetVerifiedAtOk() (*string, bool)`

GetVerifiedAtOk returns a tuple with the VerifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerifiedAt

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) SetVerifiedAt(v string)`

SetVerifiedAt sets VerifiedAt field to given value.

### HasVerifiedAt

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) HasVerifiedAt() bool`

HasVerifiedAt returns a boolean if a field has been set.

### SetVerifiedAtNil

`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) SetVerifiedAtNil(b bool)`

 SetVerifiedAtNil sets the value for VerifiedAt to be an explicit nil

### UnsetVerifiedAt
`func (o *CreateMagicScanSynthesizeResponsePayloadGrading) UnsetVerifiedAt()`

UnsetVerifiedAt ensures that no value is present for VerifiedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


