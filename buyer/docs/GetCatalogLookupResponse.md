# GetCatalogLookupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Identifier** | [**GetCatalogLookupResponseIdentifier**](GetCatalogLookupResponseIdentifier.md) |  | 
**Offers** | [**[]GetCatalogLookupResponseOffers**](GetCatalogLookupResponseOffers.md) |  | 

## Methods

### NewGetCatalogLookupResponse

`func NewGetCatalogLookupResponse(identifier GetCatalogLookupResponseIdentifier, offers []GetCatalogLookupResponseOffers, ) *GetCatalogLookupResponse`

NewGetCatalogLookupResponse instantiates a new GetCatalogLookupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCatalogLookupResponseWithDefaults

`func NewGetCatalogLookupResponseWithDefaults() *GetCatalogLookupResponse`

NewGetCatalogLookupResponseWithDefaults instantiates a new GetCatalogLookupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdentifier

`func (o *GetCatalogLookupResponse) GetIdentifier() GetCatalogLookupResponseIdentifier`

GetIdentifier returns the Identifier field if non-nil, zero value otherwise.

### GetIdentifierOk

`func (o *GetCatalogLookupResponse) GetIdentifierOk() (*GetCatalogLookupResponseIdentifier, bool)`

GetIdentifierOk returns a tuple with the Identifier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentifier

`func (o *GetCatalogLookupResponse) SetIdentifier(v GetCatalogLookupResponseIdentifier)`

SetIdentifier sets Identifier field to given value.


### GetOffers

`func (o *GetCatalogLookupResponse) GetOffers() []GetCatalogLookupResponseOffers`

GetOffers returns the Offers field if non-nil, zero value otherwise.

### GetOffersOk

`func (o *GetCatalogLookupResponse) GetOffersOk() (*[]GetCatalogLookupResponseOffers, bool)`

GetOffersOk returns a tuple with the Offers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffers

`func (o *GetCatalogLookupResponse) SetOffers(v []GetCatalogLookupResponseOffers)`

SetOffers sets Offers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


