# GetMarketProductResponseVariants

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Label** | **string** |  | 
**SortKey** | **float32** |  | 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**ExternalSource** | Pointer to **NullableString** |  | [optional] 
**ExternalSyncedAt** | Pointer to **NullableTime** |  | [optional] 
**LastSaleCents** | Pointer to **NullableFloat32** |  | [optional] 
**TradesCount** | **float32** |  | 
**SkuId** | **string** |  | 
**LastSaleAt** | Pointer to **NullableTime** |  | [optional] 
**BestBidCents** | Pointer to **NullableFloat32** |  | [optional] 
**BestAskCents** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewGetMarketProductResponseVariants

`func NewGetMarketProductResponseVariants(id string, label string, sortKey float32, tradesCount float32, skuId string, ) *GetMarketProductResponseVariants`

NewGetMarketProductResponseVariants instantiates a new GetMarketProductResponseVariants object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketProductResponseVariantsWithDefaults

`func NewGetMarketProductResponseVariantsWithDefaults() *GetMarketProductResponseVariants`

NewGetMarketProductResponseVariantsWithDefaults instantiates a new GetMarketProductResponseVariants object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMarketProductResponseVariants) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMarketProductResponseVariants) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMarketProductResponseVariants) SetId(v string)`

SetId sets Id field to given value.


### GetLabel

`func (o *GetMarketProductResponseVariants) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetMarketProductResponseVariants) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetMarketProductResponseVariants) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetSortKey

`func (o *GetMarketProductResponseVariants) GetSortKey() float32`

GetSortKey returns the SortKey field if non-nil, zero value otherwise.

### GetSortKeyOk

`func (o *GetMarketProductResponseVariants) GetSortKeyOk() (*float32, bool)`

GetSortKeyOk returns a tuple with the SortKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortKey

`func (o *GetMarketProductResponseVariants) SetSortKey(v float32)`

SetSortKey sets SortKey field to given value.


### GetExternalId

`func (o *GetMarketProductResponseVariants) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetMarketProductResponseVariants) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetMarketProductResponseVariants) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetMarketProductResponseVariants) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetMarketProductResponseVariants) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetMarketProductResponseVariants) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetExternalSource

`func (o *GetMarketProductResponseVariants) GetExternalSource() string`

GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.

### GetExternalSourceOk

`func (o *GetMarketProductResponseVariants) GetExternalSourceOk() (*string, bool)`

GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSource

`func (o *GetMarketProductResponseVariants) SetExternalSource(v string)`

SetExternalSource sets ExternalSource field to given value.

### HasExternalSource

`func (o *GetMarketProductResponseVariants) HasExternalSource() bool`

HasExternalSource returns a boolean if a field has been set.

### SetExternalSourceNil

`func (o *GetMarketProductResponseVariants) SetExternalSourceNil(b bool)`

 SetExternalSourceNil sets the value for ExternalSource to be an explicit nil

### UnsetExternalSource
`func (o *GetMarketProductResponseVariants) UnsetExternalSource()`

UnsetExternalSource ensures that no value is present for ExternalSource, not even an explicit nil
### GetExternalSyncedAt

`func (o *GetMarketProductResponseVariants) GetExternalSyncedAt() time.Time`

GetExternalSyncedAt returns the ExternalSyncedAt field if non-nil, zero value otherwise.

### GetExternalSyncedAtOk

`func (o *GetMarketProductResponseVariants) GetExternalSyncedAtOk() (*time.Time, bool)`

GetExternalSyncedAtOk returns a tuple with the ExternalSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSyncedAt

`func (o *GetMarketProductResponseVariants) SetExternalSyncedAt(v time.Time)`

SetExternalSyncedAt sets ExternalSyncedAt field to given value.

### HasExternalSyncedAt

`func (o *GetMarketProductResponseVariants) HasExternalSyncedAt() bool`

HasExternalSyncedAt returns a boolean if a field has been set.

### SetExternalSyncedAtNil

`func (o *GetMarketProductResponseVariants) SetExternalSyncedAtNil(b bool)`

 SetExternalSyncedAtNil sets the value for ExternalSyncedAt to be an explicit nil

### UnsetExternalSyncedAt
`func (o *GetMarketProductResponseVariants) UnsetExternalSyncedAt()`

UnsetExternalSyncedAt ensures that no value is present for ExternalSyncedAt, not even an explicit nil
### GetLastSaleCents

`func (o *GetMarketProductResponseVariants) GetLastSaleCents() float32`

GetLastSaleCents returns the LastSaleCents field if non-nil, zero value otherwise.

### GetLastSaleCentsOk

`func (o *GetMarketProductResponseVariants) GetLastSaleCentsOk() (*float32, bool)`

GetLastSaleCentsOk returns a tuple with the LastSaleCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSaleCents

`func (o *GetMarketProductResponseVariants) SetLastSaleCents(v float32)`

SetLastSaleCents sets LastSaleCents field to given value.

### HasLastSaleCents

`func (o *GetMarketProductResponseVariants) HasLastSaleCents() bool`

HasLastSaleCents returns a boolean if a field has been set.

### SetLastSaleCentsNil

`func (o *GetMarketProductResponseVariants) SetLastSaleCentsNil(b bool)`

 SetLastSaleCentsNil sets the value for LastSaleCents to be an explicit nil

### UnsetLastSaleCents
`func (o *GetMarketProductResponseVariants) UnsetLastSaleCents()`

UnsetLastSaleCents ensures that no value is present for LastSaleCents, not even an explicit nil
### GetTradesCount

`func (o *GetMarketProductResponseVariants) GetTradesCount() float32`

GetTradesCount returns the TradesCount field if non-nil, zero value otherwise.

### GetTradesCountOk

`func (o *GetMarketProductResponseVariants) GetTradesCountOk() (*float32, bool)`

GetTradesCountOk returns a tuple with the TradesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradesCount

`func (o *GetMarketProductResponseVariants) SetTradesCount(v float32)`

SetTradesCount sets TradesCount field to given value.


### GetSkuId

`func (o *GetMarketProductResponseVariants) GetSkuId() string`

GetSkuId returns the SkuId field if non-nil, zero value otherwise.

### GetSkuIdOk

`func (o *GetMarketProductResponseVariants) GetSkuIdOk() (*string, bool)`

GetSkuIdOk returns a tuple with the SkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkuId

`func (o *GetMarketProductResponseVariants) SetSkuId(v string)`

SetSkuId sets SkuId field to given value.


### GetLastSaleAt

`func (o *GetMarketProductResponseVariants) GetLastSaleAt() time.Time`

GetLastSaleAt returns the LastSaleAt field if non-nil, zero value otherwise.

### GetLastSaleAtOk

`func (o *GetMarketProductResponseVariants) GetLastSaleAtOk() (*time.Time, bool)`

GetLastSaleAtOk returns a tuple with the LastSaleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSaleAt

`func (o *GetMarketProductResponseVariants) SetLastSaleAt(v time.Time)`

SetLastSaleAt sets LastSaleAt field to given value.

### HasLastSaleAt

`func (o *GetMarketProductResponseVariants) HasLastSaleAt() bool`

HasLastSaleAt returns a boolean if a field has been set.

### SetLastSaleAtNil

`func (o *GetMarketProductResponseVariants) SetLastSaleAtNil(b bool)`

 SetLastSaleAtNil sets the value for LastSaleAt to be an explicit nil

### UnsetLastSaleAt
`func (o *GetMarketProductResponseVariants) UnsetLastSaleAt()`

UnsetLastSaleAt ensures that no value is present for LastSaleAt, not even an explicit nil
### GetBestBidCents

`func (o *GetMarketProductResponseVariants) GetBestBidCents() float32`

GetBestBidCents returns the BestBidCents field if non-nil, zero value otherwise.

### GetBestBidCentsOk

`func (o *GetMarketProductResponseVariants) GetBestBidCentsOk() (*float32, bool)`

GetBestBidCentsOk returns a tuple with the BestBidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestBidCents

`func (o *GetMarketProductResponseVariants) SetBestBidCents(v float32)`

SetBestBidCents sets BestBidCents field to given value.

### HasBestBidCents

`func (o *GetMarketProductResponseVariants) HasBestBidCents() bool`

HasBestBidCents returns a boolean if a field has been set.

### SetBestBidCentsNil

`func (o *GetMarketProductResponseVariants) SetBestBidCentsNil(b bool)`

 SetBestBidCentsNil sets the value for BestBidCents to be an explicit nil

### UnsetBestBidCents
`func (o *GetMarketProductResponseVariants) UnsetBestBidCents()`

UnsetBestBidCents ensures that no value is present for BestBidCents, not even an explicit nil
### GetBestAskCents

`func (o *GetMarketProductResponseVariants) GetBestAskCents() float32`

GetBestAskCents returns the BestAskCents field if non-nil, zero value otherwise.

### GetBestAskCentsOk

`func (o *GetMarketProductResponseVariants) GetBestAskCentsOk() (*float32, bool)`

GetBestAskCentsOk returns a tuple with the BestAskCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestAskCents

`func (o *GetMarketProductResponseVariants) SetBestAskCents(v float32)`

SetBestAskCents sets BestAskCents field to given value.

### HasBestAskCents

`func (o *GetMarketProductResponseVariants) HasBestAskCents() bool`

HasBestAskCents returns a boolean if a field has been set.

### SetBestAskCentsNil

`func (o *GetMarketProductResponseVariants) SetBestAskCentsNil(b bool)`

 SetBestAskCentsNil sets the value for BestAskCents to be an explicit nil

### UnsetBestAskCents
`func (o *GetMarketProductResponseVariants) UnsetBestAskCents()`

UnsetBestAskCents ensures that no value is present for BestAskCents, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


