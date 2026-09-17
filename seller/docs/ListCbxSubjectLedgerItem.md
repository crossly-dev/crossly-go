# ListCbxSubjectLedgerItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**DeltaBaseUnits** | **string** |  | 
**Kind** | **string** |  | 
**Memo** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 

## Methods

### NewListCbxSubjectLedgerItem

`func NewListCbxSubjectLedgerItem(id string, deltaBaseUnits string, kind string, createdAt time.Time, ) *ListCbxSubjectLedgerItem`

NewListCbxSubjectLedgerItem instantiates a new ListCbxSubjectLedgerItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCbxSubjectLedgerItemWithDefaults

`func NewListCbxSubjectLedgerItemWithDefaults() *ListCbxSubjectLedgerItem`

NewListCbxSubjectLedgerItemWithDefaults instantiates a new ListCbxSubjectLedgerItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCbxSubjectLedgerItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCbxSubjectLedgerItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCbxSubjectLedgerItem) SetId(v string)`

SetId sets Id field to given value.


### GetDeltaBaseUnits

`func (o *ListCbxSubjectLedgerItem) GetDeltaBaseUnits() string`

GetDeltaBaseUnits returns the DeltaBaseUnits field if non-nil, zero value otherwise.

### GetDeltaBaseUnitsOk

`func (o *ListCbxSubjectLedgerItem) GetDeltaBaseUnitsOk() (*string, bool)`

GetDeltaBaseUnitsOk returns a tuple with the DeltaBaseUnits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeltaBaseUnits

`func (o *ListCbxSubjectLedgerItem) SetDeltaBaseUnits(v string)`

SetDeltaBaseUnits sets DeltaBaseUnits field to given value.


### GetKind

`func (o *ListCbxSubjectLedgerItem) GetKind() string`

GetKind returns the Kind field if non-nil, zero value otherwise.

### GetKindOk

`func (o *ListCbxSubjectLedgerItem) GetKindOk() (*string, bool)`

GetKindOk returns a tuple with the Kind field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKind

`func (o *ListCbxSubjectLedgerItem) SetKind(v string)`

SetKind sets Kind field to given value.


### GetMemo

`func (o *ListCbxSubjectLedgerItem) GetMemo() string`

GetMemo returns the Memo field if non-nil, zero value otherwise.

### GetMemoOk

`func (o *ListCbxSubjectLedgerItem) GetMemoOk() (*string, bool)`

GetMemoOk returns a tuple with the Memo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemo

`func (o *ListCbxSubjectLedgerItem) SetMemo(v string)`

SetMemo sets Memo field to given value.

### HasMemo

`func (o *ListCbxSubjectLedgerItem) HasMemo() bool`

HasMemo returns a boolean if a field has been set.

### SetMemoNil

`func (o *ListCbxSubjectLedgerItem) SetMemoNil(b bool)`

 SetMemoNil sets the value for Memo to be an explicit nil

### UnsetMemo
`func (o *ListCbxSubjectLedgerItem) UnsetMemo()`

UnsetMemo ensures that no value is present for Memo, not even an explicit nil
### GetCreatedAt

`func (o *ListCbxSubjectLedgerItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListCbxSubjectLedgerItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListCbxSubjectLedgerItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


