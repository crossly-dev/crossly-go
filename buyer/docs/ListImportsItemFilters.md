# ListImportsItemFilters

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ListedAfter** | Pointer to **NullableString** |  | [optional] 
**ListedBefore** | Pointer to **NullableString** |  | [optional] 
**MinPrice** | Pointer to **NullableFloat32** |  | [optional] 
**MaxPrice** | Pointer to **NullableFloat32** |  | [optional] 
**ConditionIncludes** | Pointer to **[]string** |  | [optional] 
**MaxImport** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewListImportsItemFilters

`func NewListImportsItemFilters() *ListImportsItemFilters`

NewListImportsItemFilters instantiates a new ListImportsItemFilters object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImportsItemFiltersWithDefaults

`func NewListImportsItemFiltersWithDefaults() *ListImportsItemFilters`

NewListImportsItemFiltersWithDefaults instantiates a new ListImportsItemFilters object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetListedAfter

`func (o *ListImportsItemFilters) GetListedAfter() string`

GetListedAfter returns the ListedAfter field if non-nil, zero value otherwise.

### GetListedAfterOk

`func (o *ListImportsItemFilters) GetListedAfterOk() (*string, bool)`

GetListedAfterOk returns a tuple with the ListedAfter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedAfter

`func (o *ListImportsItemFilters) SetListedAfter(v string)`

SetListedAfter sets ListedAfter field to given value.

### HasListedAfter

`func (o *ListImportsItemFilters) HasListedAfter() bool`

HasListedAfter returns a boolean if a field has been set.

### SetListedAfterNil

`func (o *ListImportsItemFilters) SetListedAfterNil(b bool)`

 SetListedAfterNil sets the value for ListedAfter to be an explicit nil

### UnsetListedAfter
`func (o *ListImportsItemFilters) UnsetListedAfter()`

UnsetListedAfter ensures that no value is present for ListedAfter, not even an explicit nil
### GetListedBefore

`func (o *ListImportsItemFilters) GetListedBefore() string`

GetListedBefore returns the ListedBefore field if non-nil, zero value otherwise.

### GetListedBeforeOk

`func (o *ListImportsItemFilters) GetListedBeforeOk() (*string, bool)`

GetListedBeforeOk returns a tuple with the ListedBefore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedBefore

`func (o *ListImportsItemFilters) SetListedBefore(v string)`

SetListedBefore sets ListedBefore field to given value.

### HasListedBefore

`func (o *ListImportsItemFilters) HasListedBefore() bool`

HasListedBefore returns a boolean if a field has been set.

### SetListedBeforeNil

`func (o *ListImportsItemFilters) SetListedBeforeNil(b bool)`

 SetListedBeforeNil sets the value for ListedBefore to be an explicit nil

### UnsetListedBefore
`func (o *ListImportsItemFilters) UnsetListedBefore()`

UnsetListedBefore ensures that no value is present for ListedBefore, not even an explicit nil
### GetMinPrice

`func (o *ListImportsItemFilters) GetMinPrice() float32`

GetMinPrice returns the MinPrice field if non-nil, zero value otherwise.

### GetMinPriceOk

`func (o *ListImportsItemFilters) GetMinPriceOk() (*float32, bool)`

GetMinPriceOk returns a tuple with the MinPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPrice

`func (o *ListImportsItemFilters) SetMinPrice(v float32)`

SetMinPrice sets MinPrice field to given value.

### HasMinPrice

`func (o *ListImportsItemFilters) HasMinPrice() bool`

HasMinPrice returns a boolean if a field has been set.

### SetMinPriceNil

`func (o *ListImportsItemFilters) SetMinPriceNil(b bool)`

 SetMinPriceNil sets the value for MinPrice to be an explicit nil

### UnsetMinPrice
`func (o *ListImportsItemFilters) UnsetMinPrice()`

UnsetMinPrice ensures that no value is present for MinPrice, not even an explicit nil
### GetMaxPrice

`func (o *ListImportsItemFilters) GetMaxPrice() float32`

GetMaxPrice returns the MaxPrice field if non-nil, zero value otherwise.

### GetMaxPriceOk

`func (o *ListImportsItemFilters) GetMaxPriceOk() (*float32, bool)`

GetMaxPriceOk returns a tuple with the MaxPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPrice

`func (o *ListImportsItemFilters) SetMaxPrice(v float32)`

SetMaxPrice sets MaxPrice field to given value.

### HasMaxPrice

`func (o *ListImportsItemFilters) HasMaxPrice() bool`

HasMaxPrice returns a boolean if a field has been set.

### SetMaxPriceNil

`func (o *ListImportsItemFilters) SetMaxPriceNil(b bool)`

 SetMaxPriceNil sets the value for MaxPrice to be an explicit nil

### UnsetMaxPrice
`func (o *ListImportsItemFilters) UnsetMaxPrice()`

UnsetMaxPrice ensures that no value is present for MaxPrice, not even an explicit nil
### GetConditionIncludes

`func (o *ListImportsItemFilters) GetConditionIncludes() []string`

GetConditionIncludes returns the ConditionIncludes field if non-nil, zero value otherwise.

### GetConditionIncludesOk

`func (o *ListImportsItemFilters) GetConditionIncludesOk() (*[]string, bool)`

GetConditionIncludesOk returns a tuple with the ConditionIncludes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionIncludes

`func (o *ListImportsItemFilters) SetConditionIncludes(v []string)`

SetConditionIncludes sets ConditionIncludes field to given value.

### HasConditionIncludes

`func (o *ListImportsItemFilters) HasConditionIncludes() bool`

HasConditionIncludes returns a boolean if a field has been set.

### SetConditionIncludesNil

`func (o *ListImportsItemFilters) SetConditionIncludesNil(b bool)`

 SetConditionIncludesNil sets the value for ConditionIncludes to be an explicit nil

### UnsetConditionIncludes
`func (o *ListImportsItemFilters) UnsetConditionIncludes()`

UnsetConditionIncludes ensures that no value is present for ConditionIncludes, not even an explicit nil
### GetMaxImport

`func (o *ListImportsItemFilters) GetMaxImport() float32`

GetMaxImport returns the MaxImport field if non-nil, zero value otherwise.

### GetMaxImportOk

`func (o *ListImportsItemFilters) GetMaxImportOk() (*float32, bool)`

GetMaxImportOk returns a tuple with the MaxImport field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxImport

`func (o *ListImportsItemFilters) SetMaxImport(v float32)`

SetMaxImport sets MaxImport field to given value.

### HasMaxImport

`func (o *ListImportsItemFilters) HasMaxImport() bool`

HasMaxImport returns a boolean if a field has been set.

### SetMaxImportNil

`func (o *ListImportsItemFilters) SetMaxImportNil(b bool)`

 SetMaxImportNil sets the value for MaxImport to be an explicit nil

### UnsetMaxImport
`func (o *ListImportsItemFilters) UnsetMaxImport()`

UnsetMaxImport ensures that no value is present for MaxImport, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


