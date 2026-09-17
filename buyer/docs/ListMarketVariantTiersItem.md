# ListMarketVariantTiersItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Label** | **string** |  | 
**Condition** | **string** |  | 
**GradeKey** | Pointer to **NullableString** | null &#x3D; ungraded, which is the tier every pre-grading order sits in. | [optional] 
**LowestAskCents** | Pointer to **NullableFloat32** |  | [optional] 
**HighestBidCents** | Pointer to **NullableFloat32** |  | [optional] 
**OpenAsks** | **float32** |  | 
**OpenBids** | **float32** |  | 
**LastSaleCents** | Pointer to **NullableFloat32** |  | [optional] 
**LastSaleAt** | Pointer to **NullableTime** |  | [optional] 
**TradesCount** | **float32** |  | 

## Methods

### NewListMarketVariantTiersItem

`func NewListMarketVariantTiersItem(label string, condition string, openAsks float32, openBids float32, tradesCount float32, ) *ListMarketVariantTiersItem`

NewListMarketVariantTiersItem instantiates a new ListMarketVariantTiersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMarketVariantTiersItemWithDefaults

`func NewListMarketVariantTiersItemWithDefaults() *ListMarketVariantTiersItem`

NewListMarketVariantTiersItemWithDefaults instantiates a new ListMarketVariantTiersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLabel

`func (o *ListMarketVariantTiersItem) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ListMarketVariantTiersItem) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ListMarketVariantTiersItem) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetCondition

`func (o *ListMarketVariantTiersItem) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ListMarketVariantTiersItem) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ListMarketVariantTiersItem) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetGradeKey

`func (o *ListMarketVariantTiersItem) GetGradeKey() string`

GetGradeKey returns the GradeKey field if non-nil, zero value otherwise.

### GetGradeKeyOk

`func (o *ListMarketVariantTiersItem) GetGradeKeyOk() (*string, bool)`

GetGradeKeyOk returns a tuple with the GradeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeKey

`func (o *ListMarketVariantTiersItem) SetGradeKey(v string)`

SetGradeKey sets GradeKey field to given value.

### HasGradeKey

`func (o *ListMarketVariantTiersItem) HasGradeKey() bool`

HasGradeKey returns a boolean if a field has been set.

### SetGradeKeyNil

`func (o *ListMarketVariantTiersItem) SetGradeKeyNil(b bool)`

 SetGradeKeyNil sets the value for GradeKey to be an explicit nil

### UnsetGradeKey
`func (o *ListMarketVariantTiersItem) UnsetGradeKey()`

UnsetGradeKey ensures that no value is present for GradeKey, not even an explicit nil
### GetLowestAskCents

`func (o *ListMarketVariantTiersItem) GetLowestAskCents() float32`

GetLowestAskCents returns the LowestAskCents field if non-nil, zero value otherwise.

### GetLowestAskCentsOk

`func (o *ListMarketVariantTiersItem) GetLowestAskCentsOk() (*float32, bool)`

GetLowestAskCentsOk returns a tuple with the LowestAskCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestAskCents

`func (o *ListMarketVariantTiersItem) SetLowestAskCents(v float32)`

SetLowestAskCents sets LowestAskCents field to given value.

### HasLowestAskCents

`func (o *ListMarketVariantTiersItem) HasLowestAskCents() bool`

HasLowestAskCents returns a boolean if a field has been set.

### SetLowestAskCentsNil

`func (o *ListMarketVariantTiersItem) SetLowestAskCentsNil(b bool)`

 SetLowestAskCentsNil sets the value for LowestAskCents to be an explicit nil

### UnsetLowestAskCents
`func (o *ListMarketVariantTiersItem) UnsetLowestAskCents()`

UnsetLowestAskCents ensures that no value is present for LowestAskCents, not even an explicit nil
### GetHighestBidCents

`func (o *ListMarketVariantTiersItem) GetHighestBidCents() float32`

GetHighestBidCents returns the HighestBidCents field if non-nil, zero value otherwise.

### GetHighestBidCentsOk

`func (o *ListMarketVariantTiersItem) GetHighestBidCentsOk() (*float32, bool)`

GetHighestBidCentsOk returns a tuple with the HighestBidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestBidCents

`func (o *ListMarketVariantTiersItem) SetHighestBidCents(v float32)`

SetHighestBidCents sets HighestBidCents field to given value.

### HasHighestBidCents

`func (o *ListMarketVariantTiersItem) HasHighestBidCents() bool`

HasHighestBidCents returns a boolean if a field has been set.

### SetHighestBidCentsNil

`func (o *ListMarketVariantTiersItem) SetHighestBidCentsNil(b bool)`

 SetHighestBidCentsNil sets the value for HighestBidCents to be an explicit nil

### UnsetHighestBidCents
`func (o *ListMarketVariantTiersItem) UnsetHighestBidCents()`

UnsetHighestBidCents ensures that no value is present for HighestBidCents, not even an explicit nil
### GetOpenAsks

`func (o *ListMarketVariantTiersItem) GetOpenAsks() float32`

GetOpenAsks returns the OpenAsks field if non-nil, zero value otherwise.

### GetOpenAsksOk

`func (o *ListMarketVariantTiersItem) GetOpenAsksOk() (*float32, bool)`

GetOpenAsksOk returns a tuple with the OpenAsks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenAsks

`func (o *ListMarketVariantTiersItem) SetOpenAsks(v float32)`

SetOpenAsks sets OpenAsks field to given value.


### GetOpenBids

`func (o *ListMarketVariantTiersItem) GetOpenBids() float32`

GetOpenBids returns the OpenBids field if non-nil, zero value otherwise.

### GetOpenBidsOk

`func (o *ListMarketVariantTiersItem) GetOpenBidsOk() (*float32, bool)`

GetOpenBidsOk returns a tuple with the OpenBids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenBids

`func (o *ListMarketVariantTiersItem) SetOpenBids(v float32)`

SetOpenBids sets OpenBids field to given value.


### GetLastSaleCents

`func (o *ListMarketVariantTiersItem) GetLastSaleCents() float32`

GetLastSaleCents returns the LastSaleCents field if non-nil, zero value otherwise.

### GetLastSaleCentsOk

`func (o *ListMarketVariantTiersItem) GetLastSaleCentsOk() (*float32, bool)`

GetLastSaleCentsOk returns a tuple with the LastSaleCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSaleCents

`func (o *ListMarketVariantTiersItem) SetLastSaleCents(v float32)`

SetLastSaleCents sets LastSaleCents field to given value.

### HasLastSaleCents

`func (o *ListMarketVariantTiersItem) HasLastSaleCents() bool`

HasLastSaleCents returns a boolean if a field has been set.

### SetLastSaleCentsNil

`func (o *ListMarketVariantTiersItem) SetLastSaleCentsNil(b bool)`

 SetLastSaleCentsNil sets the value for LastSaleCents to be an explicit nil

### UnsetLastSaleCents
`func (o *ListMarketVariantTiersItem) UnsetLastSaleCents()`

UnsetLastSaleCents ensures that no value is present for LastSaleCents, not even an explicit nil
### GetLastSaleAt

`func (o *ListMarketVariantTiersItem) GetLastSaleAt() time.Time`

GetLastSaleAt returns the LastSaleAt field if non-nil, zero value otherwise.

### GetLastSaleAtOk

`func (o *ListMarketVariantTiersItem) GetLastSaleAtOk() (*time.Time, bool)`

GetLastSaleAtOk returns a tuple with the LastSaleAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSaleAt

`func (o *ListMarketVariantTiersItem) SetLastSaleAt(v time.Time)`

SetLastSaleAt sets LastSaleAt field to given value.

### HasLastSaleAt

`func (o *ListMarketVariantTiersItem) HasLastSaleAt() bool`

HasLastSaleAt returns a boolean if a field has been set.

### SetLastSaleAtNil

`func (o *ListMarketVariantTiersItem) SetLastSaleAtNil(b bool)`

 SetLastSaleAtNil sets the value for LastSaleAt to be an explicit nil

### UnsetLastSaleAt
`func (o *ListMarketVariantTiersItem) UnsetLastSaleAt()`

UnsetLastSaleAt ensures that no value is present for LastSaleAt, not even an explicit nil
### GetTradesCount

`func (o *ListMarketVariantTiersItem) GetTradesCount() float32`

GetTradesCount returns the TradesCount field if non-nil, zero value otherwise.

### GetTradesCountOk

`func (o *ListMarketVariantTiersItem) GetTradesCountOk() (*float32, bool)`

GetTradesCountOk returns a tuple with the TradesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradesCount

`func (o *ListMarketVariantTiersItem) SetTradesCount(v float32)`

SetTradesCount sets TradesCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


