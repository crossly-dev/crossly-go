# CreateListingImportByUrlResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Outcome** | **string** |  | 
**Listing** | Pointer to [**NullableCreateListingImportByUrlResponseListing**](CreateListingImportByUrlResponseListing.md) |  | [optional] 

## Methods

### NewCreateListingImportByUrlResponse

`func NewCreateListingImportByUrlResponse(outcome string, ) *CreateListingImportByUrlResponse`

NewCreateListingImportByUrlResponse instantiates a new CreateListingImportByUrlResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateListingImportByUrlResponseWithDefaults

`func NewCreateListingImportByUrlResponseWithDefaults() *CreateListingImportByUrlResponse`

NewCreateListingImportByUrlResponseWithDefaults instantiates a new CreateListingImportByUrlResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOutcome

`func (o *CreateListingImportByUrlResponse) GetOutcome() string`

GetOutcome returns the Outcome field if non-nil, zero value otherwise.

### GetOutcomeOk

`func (o *CreateListingImportByUrlResponse) GetOutcomeOk() (*string, bool)`

GetOutcomeOk returns a tuple with the Outcome field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcome

`func (o *CreateListingImportByUrlResponse) SetOutcome(v string)`

SetOutcome sets Outcome field to given value.


### GetListing

`func (o *CreateListingImportByUrlResponse) GetListing() CreateListingImportByUrlResponseListing`

GetListing returns the Listing field if non-nil, zero value otherwise.

### GetListingOk

`func (o *CreateListingImportByUrlResponse) GetListingOk() (*CreateListingImportByUrlResponseListing, bool)`

GetListingOk returns a tuple with the Listing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListing

`func (o *CreateListingImportByUrlResponse) SetListing(v CreateListingImportByUrlResponseListing)`

SetListing sets Listing field to given value.

### HasListing

`func (o *CreateListingImportByUrlResponse) HasListing() bool`

HasListing returns a boolean if a field has been set.

### SetListingNil

`func (o *CreateListingImportByUrlResponse) SetListingNil(b bool)`

 SetListingNil sets the value for Listing to be an explicit nil

### UnsetListing
`func (o *CreateListingImportByUrlResponse) UnsetListing()`

UnsetListing ensures that no value is present for Listing, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


