# GetBuyerScanSessionResponseCaptures

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**MatchMethod** | **string** |  | 
**Identifier** | Pointer to [**NullableGetBuyerScanSessionResponseIdentifier**](GetBuyerScanSessionResponseIdentifier.md) |  | [optional] 
**Verdict** | **string** |  | 
**DisplayLine** | Pointer to **NullableString** |  | [optional] 
**SavingCents** | Pointer to **NullableFloat32** |  | [optional] 
**CreatedAt** | **string** |  | 

## Methods

### NewGetBuyerScanSessionResponseCaptures

`func NewGetBuyerScanSessionResponseCaptures(id string, matchMethod string, verdict string, createdAt string, ) *GetBuyerScanSessionResponseCaptures`

NewGetBuyerScanSessionResponseCaptures instantiates a new GetBuyerScanSessionResponseCaptures object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBuyerScanSessionResponseCapturesWithDefaults

`func NewGetBuyerScanSessionResponseCapturesWithDefaults() *GetBuyerScanSessionResponseCaptures`

NewGetBuyerScanSessionResponseCapturesWithDefaults instantiates a new GetBuyerScanSessionResponseCaptures object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetBuyerScanSessionResponseCaptures) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetBuyerScanSessionResponseCaptures) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetBuyerScanSessionResponseCaptures) SetId(v string)`

SetId sets Id field to given value.


### GetMatchMethod

`func (o *GetBuyerScanSessionResponseCaptures) GetMatchMethod() string`

GetMatchMethod returns the MatchMethod field if non-nil, zero value otherwise.

### GetMatchMethodOk

`func (o *GetBuyerScanSessionResponseCaptures) GetMatchMethodOk() (*string, bool)`

GetMatchMethodOk returns a tuple with the MatchMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchMethod

`func (o *GetBuyerScanSessionResponseCaptures) SetMatchMethod(v string)`

SetMatchMethod sets MatchMethod field to given value.


### GetIdentifier

`func (o *GetBuyerScanSessionResponseCaptures) GetIdentifier() GetBuyerScanSessionResponseIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *GetBuyerScanSessionResponseCaptures) GetIdentifierOk() (*GetBuyerScanSessionResponseIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *GetBuyerScanSessionResponseCaptures) SetIdentifier(v GetBuyerScanSessionResponseIdentifier)`

SetIdentifier sets Identifier field to given value.

### HasIdentifier

`func (o *GetBuyerScanSessionResponseCaptures) HasIdentifier() bool`

HasIdentifier returns a boolean if a field has been set.

### SetIdentifierNil

`func (o *GetBuyerScanSessionResponseCaptures) SetIdentifierNil(b bool)`

 SetIdentifierNil sets the value for Identifier to be an explicit nil

### UnsetIdentifier
`func (o *GetBuyerScanSessionResponseCaptures) UnsetIdentifier()`

UnsetIdentifier ensures that no value is present for Identifier, not even an explicit nil
### GetVerdict

`func (o *GetBuyerScanSessionResponseCaptures) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *GetBuyerScanSessionResponseCaptures) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *GetBuyerScanSessionResponseCaptures) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.


### GetDisplayLine

`func (o *GetBuyerScanSessionResponseCaptures) GetDisplayLine() string`

GetDisplayLine returns the DisplayLine field if non-nil, zero value otherwise.

### GetDisplayLineOk

`func (o *GetBuyerScanSessionResponseCaptures) GetDisplayLineOk() (*string, bool)`

GetDisplayLineOk returns a tuple with the DisplayLine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayLine

`func (o *GetBuyerScanSessionResponseCaptures) SetDisplayLine(v string)`

SetDisplayLine sets DisplayLine field to given value.

### HasDisplayLine

`func (o *GetBuyerScanSessionResponseCaptures) HasDisplayLine() bool`

HasDisplayLine returns a boolean if a field has been set.

### SetDisplayLineNil

`func (o *GetBuyerScanSessionResponseCaptures) SetDisplayLineNil(b bool)`

 SetDisplayLineNil sets the value for DisplayLine to be an explicit nil

### UnsetDisplayLine
`func (o *GetBuyerScanSessionResponseCaptures) UnsetDisplayLine()`

UnsetDisplayLine ensures that no value is present for DisplayLine, not even an explicit nil
### GetSavingCents

`func (o *GetBuyerScanSessionResponseCaptures) GetSavingCents() float32`

GetSavingCents returns the SavingCents field if non-nil, zero value otherwise.

### GetSavingCentsOk

`func (o *GetBuyerScanSessionResponseCaptures) GetSavingCentsOk() (*float32, bool)`

GetSavingCentsOk returns a tuple with the SavingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingCents

`func (o *GetBuyerScanSessionResponseCaptures) SetSavingCents(v float32)`

SetSavingCents sets SavingCents field to given value.

### HasSavingCents

`func (o *GetBuyerScanSessionResponseCaptures) HasSavingCents() bool`

HasSavingCents returns a boolean if a field has been set.

### SetSavingCentsNil

`func (o *GetBuyerScanSessionResponseCaptures) SetSavingCentsNil(b bool)`

 SetSavingCentsNil sets the value for SavingCents to be an explicit nil

### UnsetSavingCents
`func (o *GetBuyerScanSessionResponseCaptures) UnsetSavingCents()`

UnsetSavingCents ensures that no value is present for SavingCents, not even an explicit nil
### GetCreatedAt

`func (o *GetBuyerScanSessionResponseCaptures) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetBuyerScanSessionResponseCaptures) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetBuyerScanSessionResponseCaptures) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


