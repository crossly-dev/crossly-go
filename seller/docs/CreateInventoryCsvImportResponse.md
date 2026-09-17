# CreateInventoryCsvImportResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Problems** | [**[]CreateInventoryCsvImportResponseProblems**](CreateInventoryCsvImportResponseProblems.md) |  | 
**ProblemCount** | **float32** |  | 
**MaxRows** | **float32** |  | 
**Created** | **float32** |  | 
**Updated** | **float32** |  | 
**Usable** | **float32** | Rows that mapped cleanly. &#x60;created + updated&#x60; when not a dry run. | 
**TotalRows** | **float32** |  | 
**ListingsCreated** | **float32** |  | 
**DryRun** | **bool** | True when nothing was written — a preview pass. | 

## Methods

### NewCreateInventoryCsvImportResponse

`func NewCreateInventoryCsvImportResponse(problems []CreateInventoryCsvImportResponseProblems, problemCount float32, maxRows float32, created float32, updated float32, usable float32, totalRows float32, listingsCreated float32, dryRun bool, ) *CreateInventoryCsvImportResponse`

NewCreateInventoryCsvImportResponse instantiates a new CreateInventoryCsvImportResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInventoryCsvImportResponseWithDefaults

`func NewCreateInventoryCsvImportResponseWithDefaults() *CreateInventoryCsvImportResponse`

NewCreateInventoryCsvImportResponseWithDefaults instantiates a new CreateInventoryCsvImportResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProblems

`func (o *CreateInventoryCsvImportResponse) GetProblems() []CreateInventoryCsvImportResponseProblems`

GetProblems returns the Problems field if non-nil, zero value otherwise.

### GetProblemsOk

`func (o *CreateInventoryCsvImportResponse) GetProblemsOk() (*[]CreateInventoryCsvImportResponseProblems, bool)`

GetProblemsOk returns a tuple with the Problems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblems

`func (o *CreateInventoryCsvImportResponse) SetProblems(v []CreateInventoryCsvImportResponseProblems)`

SetProblems sets Problems field to given value.


### GetProblemCount

`func (o *CreateInventoryCsvImportResponse) GetProblemCount() float32`

GetProblemCount returns the ProblemCount field if non-nil, zero value otherwise.

### GetProblemCountOk

`func (o *CreateInventoryCsvImportResponse) GetProblemCountOk() (*float32, bool)`

GetProblemCountOk returns a tuple with the ProblemCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProblemCount

`func (o *CreateInventoryCsvImportResponse) SetProblemCount(v float32)`

SetProblemCount sets ProblemCount field to given value.


### GetMaxRows

`func (o *CreateInventoryCsvImportResponse) GetMaxRows() float32`

GetMaxRows returns the MaxRows field if non-nil, zero value otherwise.

### GetMaxRowsOk

`func (o *CreateInventoryCsvImportResponse) GetMaxRowsOk() (*float32, bool)`

GetMaxRowsOk returns a tuple with the MaxRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxRows

`func (o *CreateInventoryCsvImportResponse) SetMaxRows(v float32)`

SetMaxRows sets MaxRows field to given value.


### GetCreated

`func (o *CreateInventoryCsvImportResponse) GetCreated() float32`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CreateInventoryCsvImportResponse) GetCreatedOk() (*float32, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CreateInventoryCsvImportResponse) SetCreated(v float32)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *CreateInventoryCsvImportResponse) GetUpdated() float32`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CreateInventoryCsvImportResponse) GetUpdatedOk() (*float32, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CreateInventoryCsvImportResponse) SetUpdated(v float32)`

SetUpdated sets Updated field to given value.


### GetUsable

`func (o *CreateInventoryCsvImportResponse) GetUsable() float32`

GetUsable returns the Usable field if non-nil, zero value otherwise.

### GetUsableOk

`func (o *CreateInventoryCsvImportResponse) GetUsableOk() (*float32, bool)`

GetUsableOk returns a tuple with the Usable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsable

`func (o *CreateInventoryCsvImportResponse) SetUsable(v float32)`

SetUsable sets Usable field to given value.


### GetTotalRows

`func (o *CreateInventoryCsvImportResponse) GetTotalRows() float32`

GetTotalRows returns the TotalRows field if non-nil, zero value otherwise.

### GetTotalRowsOk

`func (o *CreateInventoryCsvImportResponse) GetTotalRowsOk() (*float32, bool)`

GetTotalRowsOk returns a tuple with the TotalRows field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalRows

`func (o *CreateInventoryCsvImportResponse) SetTotalRows(v float32)`

SetTotalRows sets TotalRows field to given value.


### GetListingsCreated

`func (o *CreateInventoryCsvImportResponse) GetListingsCreated() float32`

GetListingsCreated returns the ListingsCreated field if non-nil, zero value otherwise.

### GetListingsCreatedOk

`func (o *CreateInventoryCsvImportResponse) GetListingsCreatedOk() (*float32, bool)`

GetListingsCreatedOk returns a tuple with the ListingsCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingsCreated

`func (o *CreateInventoryCsvImportResponse) SetListingsCreated(v float32)`

SetListingsCreated sets ListingsCreated field to given value.


### GetDryRun

`func (o *CreateInventoryCsvImportResponse) GetDryRun() bool`

GetDryRun returns the DryRun field if non-nil, zero value otherwise.

### GetDryRunOk

`func (o *CreateInventoryCsvImportResponse) GetDryRunOk() (*bool, bool)`

GetDryRunOk returns a tuple with the DryRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDryRun

`func (o *CreateInventoryCsvImportResponse) SetDryRun(v bool)`

SetDryRun sets DryRun field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


