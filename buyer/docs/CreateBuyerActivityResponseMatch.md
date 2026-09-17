# CreateBuyerActivityResponseMatch

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Verdict** | **string** |  | 
**Offer** | [**CreateBuyerActivityResponseMatchOffer**](CreateBuyerActivityResponseMatchOffer.md) |  | 
**SavingCents** | Pointer to **NullableFloat32** | Positive when cheaper. Null when no page price was supplied. | [optional] 
**ConditionComparable** | **bool** | TRUE means the offer&#39;s condition and the page&#39;s are comparable. False means we matched the item but not its state — a used Crossly copy against a retailer&#39;s new one — and the UI must say so rather than claim a saving. | 
**ShippingKnown** | **bool** | Always false today. Crossly shipping is computed at checkout from the buyer&#39;s address, which Scout does not have and should not send. Present so the surface that renders \&quot;before shipping\&quot; is reading a fact rather than hard-coding an assumption that stops being true when we add it. | 
**Catalog** | Pointer to [**NullableGetCatalogLookupResponseCatalog**](GetCatalogLookupResponseCatalog.md) |  | [optional] 
**Alternates** | [**[]CreateBuyerActivityResponseMatchOffer**](CreateBuyerActivityResponseMatchOffer.md) | Other buyable offers for the SAME item, cheapest first, best excluded.  Deliberately not \&quot;you might also like\&quot;. We have no behavioural data to build that from, and inventing it would put unrelated items under a badge whose entire value is that it only ever appears when we have the thing the buyer is actually looking at.  What these ARE is the same product from other sellers, in other conditions, at other prices — which is the choice a buyer on a product page genuinely wants, and the one a single \&quot;cheapest\&quot; result hides. | 

## Methods

### NewCreateBuyerActivityResponseMatch

`func NewCreateBuyerActivityResponseMatch(verdict string, offer CreateBuyerActivityResponseMatchOffer, conditionComparable bool, shippingKnown bool, alternates []CreateBuyerActivityResponseMatchOffer, ) *CreateBuyerActivityResponseMatch`

NewCreateBuyerActivityResponseMatch instantiates a new CreateBuyerActivityResponseMatch object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateBuyerActivityResponseMatchWithDefaults

`func NewCreateBuyerActivityResponseMatchWithDefaults() *CreateBuyerActivityResponseMatch`

NewCreateBuyerActivityResponseMatchWithDefaults instantiates a new CreateBuyerActivityResponseMatch object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVerdict

`func (o *CreateBuyerActivityResponseMatch) GetVerdict() string`

GetVerdict returns the Verdict field if non-nil, zero value otherwise.

### GetVerdictOk

`func (o *CreateBuyerActivityResponseMatch) GetVerdictOk() (*string, bool)`

GetVerdictOk returns a tuple with the Verdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerdict

`func (o *CreateBuyerActivityResponseMatch) SetVerdict(v string)`

SetVerdict sets Verdict field to given value.


### GetOffer

`func (o *CreateBuyerActivityResponseMatch) GetOffer() CreateBuyerActivityResponseMatchOffer`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *CreateBuyerActivityResponseMatch) GetOfferOk() (*CreateBuyerActivityResponseMatchOffer, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *CreateBuyerActivityResponseMatch) SetOffer(v CreateBuyerActivityResponseMatchOffer)`

SetOffer sets Offer field to given value.


### GetSavingCents

`func (o *CreateBuyerActivityResponseMatch) GetSavingCents() float32`

GetSavingCents returns the SavingCents field if non-nil, zero value otherwise.

### GetSavingCentsOk

`func (o *CreateBuyerActivityResponseMatch) GetSavingCentsOk() (*float32, bool)`

GetSavingCentsOk returns a tuple with the SavingCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavingCents

`func (o *CreateBuyerActivityResponseMatch) SetSavingCents(v float32)`

SetSavingCents sets SavingCents field to given value.

### HasSavingCents

`func (o *CreateBuyerActivityResponseMatch) HasSavingCents() bool`

HasSavingCents returns a boolean if a field has been set.

### SetSavingCentsNil

`func (o *CreateBuyerActivityResponseMatch) SetSavingCentsNil(b bool)`

 SetSavingCentsNil sets the value for SavingCents to be an explicit nil

### UnsetSavingCents
`func (o *CreateBuyerActivityResponseMatch) UnsetSavingCents()`

UnsetSavingCents ensures that no value is present for SavingCents, not even an explicit nil
### GetConditionComparable

`func (o *CreateBuyerActivityResponseMatch) GetConditionComparable() bool`

GetConditionComparable returns the ConditionComparable field if non-nil, zero value otherwise.

### GetConditionComparableOk

`func (o *CreateBuyerActivityResponseMatch) GetConditionComparableOk() (*bool, bool)`

GetConditionComparableOk returns a tuple with the ConditionComparable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConditionComparable

`func (o *CreateBuyerActivityResponseMatch) SetConditionComparable(v bool)`

SetConditionComparable sets ConditionComparable field to given value.


### GetShippingKnown

`func (o *CreateBuyerActivityResponseMatch) GetShippingKnown() bool`

GetShippingKnown returns the ShippingKnown field if non-nil, zero value otherwise.

### GetShippingKnownOk

`func (o *CreateBuyerActivityResponseMatch) GetShippingKnownOk() (*bool, bool)`

GetShippingKnownOk returns a tuple with the ShippingKnown field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingKnown

`func (o *CreateBuyerActivityResponseMatch) SetShippingKnown(v bool)`

SetShippingKnown sets ShippingKnown field to given value.


### GetCatalog

`func (o *CreateBuyerActivityResponseMatch) GetCatalog() GetCatalogLookupResponseCatalog`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *CreateBuyerActivityResponseMatch) GetCatalogOk() (*GetCatalogLookupResponseCatalog, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *CreateBuyerActivityResponseMatch) SetCatalog(v GetCatalogLookupResponseCatalog)`

SetCatalog sets Catalog field to given value.

### HasCatalog

`func (o *CreateBuyerActivityResponseMatch) HasCatalog() bool`

HasCatalog returns a boolean if a field has been set.

### SetCatalogNil

`func (o *CreateBuyerActivityResponseMatch) SetCatalogNil(b bool)`

 SetCatalogNil sets the value for Catalog to be an explicit nil

### UnsetCatalog
`func (o *CreateBuyerActivityResponseMatch) UnsetCatalog()`

UnsetCatalog ensures that no value is present for Catalog, not even an explicit nil
### GetAlternates

`func (o *CreateBuyerActivityResponseMatch) GetAlternates() []CreateBuyerActivityResponseMatchOffer`

GetAlternates returns the Alternates field if non-nil, zero value otherwise.

### GetAlternatesOk

`func (o *CreateBuyerActivityResponseMatch) GetAlternatesOk() (*[]CreateBuyerActivityResponseMatchOffer, bool)`

GetAlternatesOk returns a tuple with the Alternates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlternates

`func (o *CreateBuyerActivityResponseMatch) SetAlternates(v []CreateBuyerActivityResponseMatchOffer)`

SetAlternates sets Alternates field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


