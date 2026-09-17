# GetMarketVariantBookResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Condition** | **string** |  | 
**GradeKey** | Pointer to **NullableString** |  | [optional] 
**Bids** | [**[]GetMarketVariantBookResponseBids**](GetMarketVariantBookResponseBids.md) |  | 
**Asks** | [**[]GetMarketVariantBookResponseBids**](GetMarketVariantBookResponseBids.md) |  | 

## Methods

### NewGetMarketVariantBookResponse

`func NewGetMarketVariantBookResponse(condition string, bids []GetMarketVariantBookResponseBids, asks []GetMarketVariantBookResponseBids, ) *GetMarketVariantBookResponse`

NewGetMarketVariantBookResponse instantiates a new GetMarketVariantBookResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketVariantBookResponseWithDefaults

`func NewGetMarketVariantBookResponseWithDefaults() *GetMarketVariantBookResponse`

NewGetMarketVariantBookResponseWithDefaults instantiates a new GetMarketVariantBookResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCondition

`func (o *GetMarketVariantBookResponse) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GetMarketVariantBookResponse) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *GetMarketVariantBookResponse) SetCondition(v string)`

SetCondition sets Condition field to given value.


### GetGradeKey

`func (o *GetMarketVariantBookResponse) GetGradeKey() string`

GetGradeKey returns the GradeKey field if non-nil, zero value otherwise.

### GetGradeKeyOk

`func (o *GetMarketVariantBookResponse) GetGradeKeyOk() (*string, bool)`

GetGradeKeyOk returns a tuple with the GradeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeKey

`func (o *GetMarketVariantBookResponse) SetGradeKey(v string)`

SetGradeKey sets GradeKey field to given value.

### HasGradeKey

`func (o *GetMarketVariantBookResponse) HasGradeKey() bool`

HasGradeKey returns a boolean if a field has been set.

### SetGradeKeyNil

`func (o *GetMarketVariantBookResponse) SetGradeKeyNil(b bool)`

 SetGradeKeyNil sets the value for GradeKey to be an explicit nil

### UnsetGradeKey
`func (o *GetMarketVariantBookResponse) UnsetGradeKey()`

UnsetGradeKey ensures that no value is present for GradeKey, not even an explicit nil
### GetBids

`func (o *GetMarketVariantBookResponse) GetBids() []GetMarketVariantBookResponseBids`

GetBids returns the Bids field if non-nil, zero value otherwise.

### GetBidsOk

`func (o *GetMarketVariantBookResponse) GetBidsOk() (*[]GetMarketVariantBookResponseBids, bool)`

GetBidsOk returns a tuple with the Bids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBids

`func (o *GetMarketVariantBookResponse) SetBids(v []GetMarketVariantBookResponseBids)`

SetBids sets Bids field to given value.


### GetAsks

`func (o *GetMarketVariantBookResponse) GetAsks() []GetMarketVariantBookResponseBids`

GetAsks returns the Asks field if non-nil, zero value otherwise.

### GetAsksOk

`func (o *GetMarketVariantBookResponse) GetAsksOk() (*[]GetMarketVariantBookResponseBids, bool)`

GetAsksOk returns a tuple with the Asks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAsks

`func (o *GetMarketVariantBookResponse) SetAsks(v []GetMarketVariantBookResponseBids)`

SetAsks sets Asks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


