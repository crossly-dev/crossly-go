# CreateInventoryCsvImportResponseProblems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Line** | **float32** | 1-based, and counting the header — so it matches what the seller sees  in their spreadsheet. Off-by-one here makes every error unfindable. | 
**Message** | **string** |  | 

## Methods

### NewCreateInventoryCsvImportResponseProblems

`func NewCreateInventoryCsvImportResponseProblems(line float32, message string, ) *CreateInventoryCsvImportResponseProblems`

NewCreateInventoryCsvImportResponseProblems instantiates a new CreateInventoryCsvImportResponseProblems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInventoryCsvImportResponseProblemsWithDefaults

`func NewCreateInventoryCsvImportResponseProblemsWithDefaults() *CreateInventoryCsvImportResponseProblems`

NewCreateInventoryCsvImportResponseProblemsWithDefaults instantiates a new CreateInventoryCsvImportResponseProblems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLine

`func (o *CreateInventoryCsvImportResponseProblems) GetLine() float32`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *CreateInventoryCsvImportResponseProblems) GetLineOk() (*float32, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *CreateInventoryCsvImportResponseProblems) SetLine(v float32)`

SetLine sets Line field to given value.


### GetMessage

`func (o *CreateInventoryCsvImportResponseProblems) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *CreateInventoryCsvImportResponseProblems) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *CreateInventoryCsvImportResponseProblems) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


