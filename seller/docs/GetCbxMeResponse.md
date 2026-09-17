# GetCbxMeResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MerchantId** | **string** |  | 
**Slug** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**Scopes** | **[]string** |  | 
**Terms** | [**GetCbxMeResponseTerms**](GetCbxMeResponseTerms.md) |  | 

## Methods

### NewGetCbxMeResponse

`func NewGetCbxMeResponse(merchantId string, scopes []string, terms GetCbxMeResponseTerms, ) *GetCbxMeResponse`

NewGetCbxMeResponse instantiates a new GetCbxMeResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCbxMeResponseWithDefaults

`func NewGetCbxMeResponseWithDefaults() *GetCbxMeResponse`

NewGetCbxMeResponseWithDefaults instantiates a new GetCbxMeResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMerchantId

`func (o *GetCbxMeResponse) GetMerchantId() string`

GetMerchantId returns the MerchantId field if non-nil, zero value otherwise.

### GetMerchantIdOk

`func (o *GetCbxMeResponse) GetMerchantIdOk() (*string, bool)`

GetMerchantIdOk returns a tuple with the MerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMerchantId

`func (o *GetCbxMeResponse) SetMerchantId(v string)`

SetMerchantId sets MerchantId field to given value.


### GetSlug

`func (o *GetCbxMeResponse) GetSlug() string`

GetSlug returns the Slug field if non-nil, zero value otherwise.

### GetSlugOk

`func (o *GetCbxMeResponse) GetSlugOk() (*string, bool)`

GetSlugOk returns a tuple with the Slug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlug

`func (o *GetCbxMeResponse) SetSlug(v string)`

SetSlug sets Slug field to given value.

### HasSlug

`func (o *GetCbxMeResponse) HasSlug() bool`

HasSlug returns a boolean if a field has been set.

### SetSlugNil

`func (o *GetCbxMeResponse) SetSlugNil(b bool)`

 SetSlugNil sets the value for Slug to be an explicit nil

### UnsetSlug
`func (o *GetCbxMeResponse) UnsetSlug()`

UnsetSlug ensures that no value is present for Slug, not even an explicit nil
### GetName

`func (o *GetCbxMeResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCbxMeResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCbxMeResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCbxMeResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GetCbxMeResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GetCbxMeResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetStatus

`func (o *GetCbxMeResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetCbxMeResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetCbxMeResponse) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetCbxMeResponse) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *GetCbxMeResponse) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *GetCbxMeResponse) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetScopes

`func (o *GetCbxMeResponse) GetScopes() []string`

GetScopes returns the Scopes field if non-nil, zero value otherwise.

### GetScopesOk

`func (o *GetCbxMeResponse) GetScopesOk() (*[]string, bool)`

GetScopesOk returns a tuple with the Scopes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScopes

`func (o *GetCbxMeResponse) SetScopes(v []string)`

SetScopes sets Scopes field to given value.


### GetTerms

`func (o *GetCbxMeResponse) GetTerms() GetCbxMeResponseTerms`

GetTerms returns the Terms field if non-nil, zero value otherwise.

### GetTermsOk

`func (o *GetCbxMeResponse) GetTermsOk() (*GetCbxMeResponseTerms, bool)`

GetTermsOk returns a tuple with the Terms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTerms

`func (o *GetCbxMeResponse) SetTerms(v GetCbxMeResponseTerms)`

SetTerms sets Terms field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


