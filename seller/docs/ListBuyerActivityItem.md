# ListBuyerActivityItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Namespace** | **string** |  | 
**IdentifierValue** | **string** |  | 
**RetailHost** | Pointer to **NullableString** |  | [optional] 
**PagePriceCents** | Pointer to **NullableFloat32** |  | [optional] 
**Matched** | **bool** |  | 
**BestPriceCents** | Pointer to **NullableFloat32** |  | [optional] 
**ObservedAt** | **time.Time** |  | 

## Methods

### NewListBuyerActivityItem

`func NewListBuyerActivityItem(id string, namespace string, identifierValue string, matched bool, observedAt time.Time, ) *ListBuyerActivityItem`

NewListBuyerActivityItem instantiates a new ListBuyerActivityItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBuyerActivityItemWithDefaults

`func NewListBuyerActivityItemWithDefaults() *ListBuyerActivityItem`

NewListBuyerActivityItemWithDefaults instantiates a new ListBuyerActivityItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListBuyerActivityItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListBuyerActivityItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListBuyerActivityItem) SetId(v string)`

SetId sets Id field to given value.


### GetNamespace

`func (o *ListBuyerActivityItem) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *ListBuyerActivityItem) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *ListBuyerActivityItem) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetIdentifierValue

`func (o *ListBuyerActivityItem) GetIdentifierValue() string`

GetIdentifierValue returns the IdentifierValue field if non-nil, zero value otherwise.

### GetIdentifierValueOk

`func (o *ListBuyerActivityItem) GetIdentifierValueOk() (*string, bool)`

GetIdentifierValueOk returns a tuple with the IdentifierValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifierValue

`func (o *ListBuyerActivityItem) SetIdentifierValue(v string)`

SetIdentifierValue sets IdentifierValue field to given value.


### GetRetailHost

`func (o *ListBuyerActivityItem) GetRetailHost() string`

GetRetailHost returns the RetailHost field if non-nil, zero value otherwise.

### GetRetailHostOk

`func (o *ListBuyerActivityItem) GetRetailHostOk() (*string, bool)`

GetRetailHostOk returns a tuple with the RetailHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailHost

`func (o *ListBuyerActivityItem) SetRetailHost(v string)`

SetRetailHost sets RetailHost field to given value.

### HasRetailHost

`func (o *ListBuyerActivityItem) HasRetailHost() bool`

HasRetailHost returns a boolean if a field has been set.

### SetRetailHostNil

`func (o *ListBuyerActivityItem) SetRetailHostNil(b bool)`

 SetRetailHostNil sets the value for RetailHost to be an explicit nil

### UnsetRetailHost
`func (o *ListBuyerActivityItem) UnsetRetailHost()`

UnsetRetailHost ensures that no value is present for RetailHost, not even an explicit nil
### GetPagePriceCents

`func (o *ListBuyerActivityItem) GetPagePriceCents() float32`

GetPagePriceCents returns the PagePriceCents field if non-nil, zero value otherwise.

### GetPagePriceCentsOk

`func (o *ListBuyerActivityItem) GetPagePriceCentsOk() (*float32, bool)`

GetPagePriceCentsOk returns a tuple with the PagePriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagePriceCents

`func (o *ListBuyerActivityItem) SetPagePriceCents(v float32)`

SetPagePriceCents sets PagePriceCents field to given value.

### HasPagePriceCents

`func (o *ListBuyerActivityItem) HasPagePriceCents() bool`

HasPagePriceCents returns a boolean if a field has been set.

### SetPagePriceCentsNil

`func (o *ListBuyerActivityItem) SetPagePriceCentsNil(b bool)`

 SetPagePriceCentsNil sets the value for PagePriceCents to be an explicit nil

### UnsetPagePriceCents
`func (o *ListBuyerActivityItem) UnsetPagePriceCents()`

UnsetPagePriceCents ensures that no value is present for PagePriceCents, not even an explicit nil
### GetMatched

`func (o *ListBuyerActivityItem) GetMatched() bool`

GetMatched returns the Matched field if non-nil, zero value otherwise.

### GetMatchedOk

`func (o *ListBuyerActivityItem) GetMatchedOk() (*bool, bool)`

GetMatchedOk returns a tuple with the Matched field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatched

`func (o *ListBuyerActivityItem) SetMatched(v bool)`

SetMatched sets Matched field to given value.


### GetBestPriceCents

`func (o *ListBuyerActivityItem) GetBestPriceCents() float32`

GetBestPriceCents returns the BestPriceCents field if non-nil, zero value otherwise.

### GetBestPriceCentsOk

`func (o *ListBuyerActivityItem) GetBestPriceCentsOk() (*float32, bool)`

GetBestPriceCentsOk returns a tuple with the BestPriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestPriceCents

`func (o *ListBuyerActivityItem) SetBestPriceCents(v float32)`

SetBestPriceCents sets BestPriceCents field to given value.

### HasBestPriceCents

`func (o *ListBuyerActivityItem) HasBestPriceCents() bool`

HasBestPriceCents returns a boolean if a field has been set.

### SetBestPriceCentsNil

`func (o *ListBuyerActivityItem) SetBestPriceCentsNil(b bool)`

 SetBestPriceCentsNil sets the value for BestPriceCents to be an explicit nil

### UnsetBestPriceCents
`func (o *ListBuyerActivityItem) UnsetBestPriceCents()`

UnsetBestPriceCents ensures that no value is present for BestPriceCents, not even an explicit nil
### GetObservedAt

`func (o *ListBuyerActivityItem) GetObservedAt() time.Time`

GetObservedAt returns the ObservedAt field if non-nil, zero value otherwise.

### GetObservedAtOk

`func (o *ListBuyerActivityItem) GetObservedAtOk() (*time.Time, bool)`

GetObservedAtOk returns a tuple with the ObservedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedAt

`func (o *ListBuyerActivityItem) SetObservedAt(v time.Time)`

SetObservedAt sets ObservedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


