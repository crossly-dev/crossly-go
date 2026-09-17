# GetCbxTreasuryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TakenAt** | **time.Time** |  | 
**CoverageBps** | **float32** |  | 
**Ok** | **bool** |  | 
**ReserveBaseUnits** | **string** |  | 
**OutstandingBaseUnits** | **string** |  | 
**InFlightClaimBaseUnits** | **string** |  | 
**Notes** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetCbxTreasuryResponse

`func NewGetCbxTreasuryResponse(takenAt time.Time, coverageBps float32, ok bool, reserveBaseUnits string, outstandingBaseUnits string, inFlightClaimBaseUnits string, ) *GetCbxTreasuryResponse`

NewGetCbxTreasuryResponse instantiates a new GetCbxTreasuryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCbxTreasuryResponseWithDefaults

`func NewGetCbxTreasuryResponseWithDefaults() *GetCbxTreasuryResponse`

NewGetCbxTreasuryResponseWithDefaults instantiates a new GetCbxTreasuryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTakenAt

`func (o *GetCbxTreasuryResponse) GetTakenAt() time.Time`

GetTakenAt returns the TakenAt field if non-nil, zero value otherwise.

### GetTakenAtOk

`func (o *GetCbxTreasuryResponse) GetTakenAtOk() (*time.Time, bool)`

GetTakenAtOk returns a tuple with the TakenAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTakenAt

`func (o *GetCbxTreasuryResponse) SetTakenAt(v time.Time)`

SetTakenAt sets TakenAt field to given value.


### GetCoverageBps

`func (o *GetCbxTreasuryResponse) GetCoverageBps() float32`

GetCoverageBps returns the CoverageBps field if non-nil, zero value otherwise.

### GetCoverageBpsOk

`func (o *GetCbxTreasuryResponse) GetCoverageBpsOk() (*float32, bool)`

GetCoverageBpsOk returns a tuple with the CoverageBps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoverageBps

`func (o *GetCbxTreasuryResponse) SetCoverageBps(v float32)`

SetCoverageBps sets CoverageBps field to given value.


### GetOk

`func (o *GetCbxTreasuryResponse) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *GetCbxTreasuryResponse) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *GetCbxTreasuryResponse) SetOk(v bool)`

SetOk sets Ok field to given value.


### GetReserveBaseUnits

`func (o *GetCbxTreasuryResponse) GetReserveBaseUnits() string`

GetReserveBaseUnits returns the ReserveBaseUnits field if non-nil, zero value otherwise.

### GetReserveBaseUnitsOk

`func (o *GetCbxTreasuryResponse) GetReserveBaseUnitsOk() (*string, bool)`

GetReserveBaseUnitsOk returns a tuple with the ReserveBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReserveBaseUnits

`func (o *GetCbxTreasuryResponse) SetReserveBaseUnits(v string)`

SetReserveBaseUnits sets ReserveBaseUnits field to given value.


### GetOutstandingBaseUnits

`func (o *GetCbxTreasuryResponse) GetOutstandingBaseUnits() string`

GetOutstandingBaseUnits returns the OutstandingBaseUnits field if non-nil, zero value otherwise.

### GetOutstandingBaseUnitsOk

`func (o *GetCbxTreasuryResponse) GetOutstandingBaseUnitsOk() (*string, bool)`

GetOutstandingBaseUnitsOk returns a tuple with the OutstandingBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutstandingBaseUnits

`func (o *GetCbxTreasuryResponse) SetOutstandingBaseUnits(v string)`

SetOutstandingBaseUnits sets OutstandingBaseUnits field to given value.


### GetInFlightClaimBaseUnits

`func (o *GetCbxTreasuryResponse) GetInFlightClaimBaseUnits() string`

GetInFlightClaimBaseUnits returns the InFlightClaimBaseUnits field if non-nil, zero value otherwise.

### GetInFlightClaimBaseUnitsOk

`func (o *GetCbxTreasuryResponse) GetInFlightClaimBaseUnitsOk() (*string, bool)`

GetInFlightClaimBaseUnitsOk returns a tuple with the InFlightClaimBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInFlightClaimBaseUnits

`func (o *GetCbxTreasuryResponse) SetInFlightClaimBaseUnits(v string)`

SetInFlightClaimBaseUnits sets InFlightClaimBaseUnits field to given value.


### GetNotes

`func (o *GetCbxTreasuryResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetCbxTreasuryResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetCbxTreasuryResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetCbxTreasuryResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *GetCbxTreasuryResponse) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *GetCbxTreasuryResponse) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


