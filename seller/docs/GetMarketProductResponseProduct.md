# GetMarketProductResponseProduct

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Brand** | Pointer to **NullableString** |  | [optional] 
**Name** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**Status** | **string** |  | 
**RetailCents** | Pointer to **NullableFloat32** |  | [optional] 
**CategorySlug** | **string** |  | 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**ImageUrl** | Pointer to **NullableString** |  | [optional] 
**Model** | Pointer to **NullableString** |  | [optional] 
**Colorway** | Pointer to **NullableString** |  | [optional] 
**StyleCode** | Pointer to **NullableString** |  | [optional] 
**ReleaseDate** | Pointer to **NullableString** |  | [optional] 
**SubmittedByUserId** | Pointer to **NullableString** |  | [optional] 
**ApprovedByUserId** | Pointer to **NullableString** |  | [optional] 
**ApprovedAt** | Pointer to **NullableTime** |  | [optional] 
**RejectionReason** | Pointer to **NullableString** |  | [optional] 
**ProposedVariants** | [**[]GetMarketProductResponseProductProposedVariants**](GetMarketProductResponseProductProposedVariants.md) |  | 
**Verification** | **map[string]interface{}** |  | 
**AiVerdict** | Pointer to **map[string]interface{}** |  | [optional] 
**SubmissionSource** | Pointer to **NullableString** |  | [optional] 
**DuplicateOfSkuId** | Pointer to **NullableString** |  | [optional] 
**PreferredProviderSlug** | Pointer to **NullableString** |  | [optional] 
**AllowPreBid** | **bool** |  | 
**PreBidOpensAt** | Pointer to **NullableTime** |  | [optional] 
**Featured** | **bool** |  | 
**ExternalSource** | Pointer to **NullableString** |  | [optional] 
**ExternalSyncedAt** | Pointer to **NullableTime** |  | [optional] 
**GradedGradeKeys** | **[]string** |  | 
**LowestAskCents** | Pointer to **NullableFloat32** |  | [optional] 
**HighestBidCents** | Pointer to **NullableFloat32** |  | [optional] 
**LastSaleCents** | Pointer to **NullableFloat32** |  | [optional] 
**TradesCount** | **float32** |  | 

## Methods

### NewGetMarketProductResponseProduct

`func NewGetMarketProductResponseProduct(id string, name string, createdAt time.Time, status string, categorySlug string, proposedVariants []GetMarketProductResponseProductProposedVariants, verification map[string]interface{}, allowPreBid bool, featured bool, gradedGradeKeys []string, tradesCount float32, ) *GetMarketProductResponseProduct`

NewGetMarketProductResponseProduct instantiates a new GetMarketProductResponseProduct object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetMarketProductResponseProductWithDefaults

`func NewGetMarketProductResponseProductWithDefaults() *GetMarketProductResponseProduct`

NewGetMarketProductResponseProductWithDefaults instantiates a new GetMarketProductResponseProduct object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetMarketProductResponseProduct) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetMarketProductResponseProduct) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetMarketProductResponseProduct) SetId(v string)`

SetId sets Id field to given value.


### GetBrand

`func (o *GetMarketProductResponseProduct) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *GetMarketProductResponseProduct) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *GetMarketProductResponseProduct) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *GetMarketProductResponseProduct) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *GetMarketProductResponseProduct) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *GetMarketProductResponseProduct) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetName

`func (o *GetMarketProductResponseProduct) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetMarketProductResponseProduct) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetMarketProductResponseProduct) SetName(v string)`

SetName sets Name field to given value.


### GetCreatedAt

`func (o *GetMarketProductResponseProduct) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetMarketProductResponseProduct) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetMarketProductResponseProduct) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetStatus

`func (o *GetMarketProductResponseProduct) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetMarketProductResponseProduct) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetMarketProductResponseProduct) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetRetailCents

`func (o *GetMarketProductResponseProduct) GetRetailCents() float32`

GetRetailCents returns the RetailCents field if non-nil, zero value otherwise.

### GetRetailCentsOk

`func (o *GetMarketProductResponseProduct) GetRetailCentsOk() (*float32, bool)`

GetRetailCentsOk returns a tuple with the RetailCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetailCents

`func (o *GetMarketProductResponseProduct) SetRetailCents(v float32)`

SetRetailCents sets RetailCents field to given value.

### HasRetailCents

`func (o *GetMarketProductResponseProduct) HasRetailCents() bool`

HasRetailCents returns a boolean if a field has been set.

### SetRetailCentsNil

`func (o *GetMarketProductResponseProduct) SetRetailCentsNil(b bool)`

 SetRetailCentsNil sets the value for RetailCents to be an explicit nil

### UnsetRetailCents
`func (o *GetMarketProductResponseProduct) UnsetRetailCents()`

UnsetRetailCents ensures that no value is present for RetailCents, not even an explicit nil
### GetCategorySlug

`func (o *GetMarketProductResponseProduct) GetCategorySlug() string`

GetCategorySlug returns the CategorySlug field if non-nil, zero value otherwise.

### GetCategorySlugOk

`func (o *GetMarketProductResponseProduct) GetCategorySlugOk() (*string, bool)`

GetCategorySlugOk returns a tuple with the CategorySlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategorySlug

`func (o *GetMarketProductResponseProduct) SetCategorySlug(v string)`

SetCategorySlug sets CategorySlug field to given value.


### GetExternalId

`func (o *GetMarketProductResponseProduct) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetMarketProductResponseProduct) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetMarketProductResponseProduct) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetMarketProductResponseProduct) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetMarketProductResponseProduct) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetMarketProductResponseProduct) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetImageUrl

`func (o *GetMarketProductResponseProduct) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *GetMarketProductResponseProduct) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *GetMarketProductResponseProduct) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.

### HasImageUrl

`func (o *GetMarketProductResponseProduct) HasImageUrl() bool`

HasImageUrl returns a boolean if a field has been set.

### SetImageUrlNil

`func (o *GetMarketProductResponseProduct) SetImageUrlNil(b bool)`

 SetImageUrlNil sets the value for ImageUrl to be an explicit nil

### UnsetImageUrl
`func (o *GetMarketProductResponseProduct) UnsetImageUrl()`

UnsetImageUrl ensures that no value is present for ImageUrl, not even an explicit nil
### GetModel

`func (o *GetMarketProductResponseProduct) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetMarketProductResponseProduct) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetMarketProductResponseProduct) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *GetMarketProductResponseProduct) HasModel() bool`

HasModel returns a boolean if a field has been set.

### SetModelNil

`func (o *GetMarketProductResponseProduct) SetModelNil(b bool)`

 SetModelNil sets the value for Model to be an explicit nil

### UnsetModel
`func (o *GetMarketProductResponseProduct) UnsetModel()`

UnsetModel ensures that no value is present for Model, not even an explicit nil
### GetColorway

`func (o *GetMarketProductResponseProduct) GetColorway() string`

GetColorway returns the Colorway field if non-nil, zero value otherwise.

### GetColorwayOk

`func (o *GetMarketProductResponseProduct) GetColorwayOk() (*string, bool)`

GetColorwayOk returns a tuple with the Colorway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColorway

`func (o *GetMarketProductResponseProduct) SetColorway(v string)`

SetColorway sets Colorway field to given value.

### HasColorway

`func (o *GetMarketProductResponseProduct) HasColorway() bool`

HasColorway returns a boolean if a field has been set.

### SetColorwayNil

`func (o *GetMarketProductResponseProduct) SetColorwayNil(b bool)`

 SetColorwayNil sets the value for Colorway to be an explicit nil

### UnsetColorway
`func (o *GetMarketProductResponseProduct) UnsetColorway()`

UnsetColorway ensures that no value is present for Colorway, not even an explicit nil
### GetStyleCode

`func (o *GetMarketProductResponseProduct) GetStyleCode() string`

GetStyleCode returns the StyleCode field if non-nil, zero value otherwise.

### GetStyleCodeOk

`func (o *GetMarketProductResponseProduct) GetStyleCodeOk() (*string, bool)`

GetStyleCodeOk returns a tuple with the StyleCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyleCode

`func (o *GetMarketProductResponseProduct) SetStyleCode(v string)`

SetStyleCode sets StyleCode field to given value.

### HasStyleCode

`func (o *GetMarketProductResponseProduct) HasStyleCode() bool`

HasStyleCode returns a boolean if a field has been set.

### SetStyleCodeNil

`func (o *GetMarketProductResponseProduct) SetStyleCodeNil(b bool)`

 SetStyleCodeNil sets the value for StyleCode to be an explicit nil

### UnsetStyleCode
`func (o *GetMarketProductResponseProduct) UnsetStyleCode()`

UnsetStyleCode ensures that no value is present for StyleCode, not even an explicit nil
### GetReleaseDate

`func (o *GetMarketProductResponseProduct) GetReleaseDate() string`

GetReleaseDate returns the ReleaseDate field if non-nil, zero value otherwise.

### GetReleaseDateOk

`func (o *GetMarketProductResponseProduct) GetReleaseDateOk() (*string, bool)`

GetReleaseDateOk returns a tuple with the ReleaseDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleaseDate

`func (o *GetMarketProductResponseProduct) SetReleaseDate(v string)`

SetReleaseDate sets ReleaseDate field to given value.

### HasReleaseDate

`func (o *GetMarketProductResponseProduct) HasReleaseDate() bool`

HasReleaseDate returns a boolean if a field has been set.

### SetReleaseDateNil

`func (o *GetMarketProductResponseProduct) SetReleaseDateNil(b bool)`

 SetReleaseDateNil sets the value for ReleaseDate to be an explicit nil

### UnsetReleaseDate
`func (o *GetMarketProductResponseProduct) UnsetReleaseDate()`

UnsetReleaseDate ensures that no value is present for ReleaseDate, not even an explicit nil
### GetSubmittedByUserId

`func (o *GetMarketProductResponseProduct) GetSubmittedByUserId() string`

GetSubmittedByUserId returns the SubmittedByUserId field if non-nil, zero value otherwise.

### GetSubmittedByUserIdOk

`func (o *GetMarketProductResponseProduct) GetSubmittedByUserIdOk() (*string, bool)`

GetSubmittedByUserIdOk returns a tuple with the SubmittedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedByUserId

`func (o *GetMarketProductResponseProduct) SetSubmittedByUserId(v string)`

SetSubmittedByUserId sets SubmittedByUserId field to given value.

### HasSubmittedByUserId

`func (o *GetMarketProductResponseProduct) HasSubmittedByUserId() bool`

HasSubmittedByUserId returns a boolean if a field has been set.

### SetSubmittedByUserIdNil

`func (o *GetMarketProductResponseProduct) SetSubmittedByUserIdNil(b bool)`

 SetSubmittedByUserIdNil sets the value for SubmittedByUserId to be an explicit nil

### UnsetSubmittedByUserId
`func (o *GetMarketProductResponseProduct) UnsetSubmittedByUserId()`

UnsetSubmittedByUserId ensures that no value is present for SubmittedByUserId, not even an explicit nil
### GetApprovedByUserId

`func (o *GetMarketProductResponseProduct) GetApprovedByUserId() string`

GetApprovedByUserId returns the ApprovedByUserId field if non-nil, zero value otherwise.

### GetApprovedByUserIdOk

`func (o *GetMarketProductResponseProduct) GetApprovedByUserIdOk() (*string, bool)`

GetApprovedByUserIdOk returns a tuple with the ApprovedByUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedByUserId

`func (o *GetMarketProductResponseProduct) SetApprovedByUserId(v string)`

SetApprovedByUserId sets ApprovedByUserId field to given value.

### HasApprovedByUserId

`func (o *GetMarketProductResponseProduct) HasApprovedByUserId() bool`

HasApprovedByUserId returns a boolean if a field has been set.

### SetApprovedByUserIdNil

`func (o *GetMarketProductResponseProduct) SetApprovedByUserIdNil(b bool)`

 SetApprovedByUserIdNil sets the value for ApprovedByUserId to be an explicit nil

### UnsetApprovedByUserId
`func (o *GetMarketProductResponseProduct) UnsetApprovedByUserId()`

UnsetApprovedByUserId ensures that no value is present for ApprovedByUserId, not even an explicit nil
### GetApprovedAt

`func (o *GetMarketProductResponseProduct) GetApprovedAt() time.Time`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *GetMarketProductResponseProduct) GetApprovedAtOk() (*time.Time, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *GetMarketProductResponseProduct) SetApprovedAt(v time.Time)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *GetMarketProductResponseProduct) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### SetApprovedAtNil

`func (o *GetMarketProductResponseProduct) SetApprovedAtNil(b bool)`

 SetApprovedAtNil sets the value for ApprovedAt to be an explicit nil

### UnsetApprovedAt
`func (o *GetMarketProductResponseProduct) UnsetApprovedAt()`

UnsetApprovedAt ensures that no value is present for ApprovedAt, not even an explicit nil
### GetRejectionReason

`func (o *GetMarketProductResponseProduct) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *GetMarketProductResponseProduct) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *GetMarketProductResponseProduct) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *GetMarketProductResponseProduct) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### SetRejectionReasonNil

`func (o *GetMarketProductResponseProduct) SetRejectionReasonNil(b bool)`

 SetRejectionReasonNil sets the value for RejectionReason to be an explicit nil

### UnsetRejectionReason
`func (o *GetMarketProductResponseProduct) UnsetRejectionReason()`

UnsetRejectionReason ensures that no value is present for RejectionReason, not even an explicit nil
### GetProposedVariants

`func (o *GetMarketProductResponseProduct) GetProposedVariants() []GetMarketProductResponseProductProposedVariants`

GetProposedVariants returns the ProposedVariants field if non-nil, zero value otherwise.

### GetProposedVariantsOk

`func (o *GetMarketProductResponseProduct) GetProposedVariantsOk() (*[]GetMarketProductResponseProductProposedVariants, bool)`

GetProposedVariantsOk returns a tuple with the ProposedVariants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProposedVariants

`func (o *GetMarketProductResponseProduct) SetProposedVariants(v []GetMarketProductResponseProductProposedVariants)`

SetProposedVariants sets ProposedVariants field to given value.


### GetVerification

`func (o *GetMarketProductResponseProduct) GetVerification() map[string]interface{}`

GetVerification returns the Verification field if non-nil, zero value otherwise.

### GetVerificationOk

`func (o *GetMarketProductResponseProduct) GetVerificationOk() (*map[string]interface{}, bool)`

GetVerificationOk returns a tuple with the Verification field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerification

`func (o *GetMarketProductResponseProduct) SetVerification(v map[string]interface{})`

SetVerification sets Verification field to given value.


### GetAiVerdict

`func (o *GetMarketProductResponseProduct) GetAiVerdict() map[string]interface{}`

GetAiVerdict returns the AiVerdict field if non-nil, zero value otherwise.

### GetAiVerdictOk

`func (o *GetMarketProductResponseProduct) GetAiVerdictOk() (*map[string]interface{}, bool)`

GetAiVerdictOk returns a tuple with the AiVerdict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAiVerdict

`func (o *GetMarketProductResponseProduct) SetAiVerdict(v map[string]interface{})`

SetAiVerdict sets AiVerdict field to given value.

### HasAiVerdict

`func (o *GetMarketProductResponseProduct) HasAiVerdict() bool`

HasAiVerdict returns a boolean if a field has been set.

### SetAiVerdictNil

`func (o *GetMarketProductResponseProduct) SetAiVerdictNil(b bool)`

 SetAiVerdictNil sets the value for AiVerdict to be an explicit nil

### UnsetAiVerdict
`func (o *GetMarketProductResponseProduct) UnsetAiVerdict()`

UnsetAiVerdict ensures that no value is present for AiVerdict, not even an explicit nil
### GetSubmissionSource

`func (o *GetMarketProductResponseProduct) GetSubmissionSource() string`

GetSubmissionSource returns the SubmissionSource field if non-nil, zero value otherwise.

### GetSubmissionSourceOk

`func (o *GetMarketProductResponseProduct) GetSubmissionSourceOk() (*string, bool)`

GetSubmissionSourceOk returns a tuple with the SubmissionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmissionSource

`func (o *GetMarketProductResponseProduct) SetSubmissionSource(v string)`

SetSubmissionSource sets SubmissionSource field to given value.

### HasSubmissionSource

`func (o *GetMarketProductResponseProduct) HasSubmissionSource() bool`

HasSubmissionSource returns a boolean if a field has been set.

### SetSubmissionSourceNil

`func (o *GetMarketProductResponseProduct) SetSubmissionSourceNil(b bool)`

 SetSubmissionSourceNil sets the value for SubmissionSource to be an explicit nil

### UnsetSubmissionSource
`func (o *GetMarketProductResponseProduct) UnsetSubmissionSource()`

UnsetSubmissionSource ensures that no value is present for SubmissionSource, not even an explicit nil
### GetDuplicateOfSkuId

`func (o *GetMarketProductResponseProduct) GetDuplicateOfSkuId() string`

GetDuplicateOfSkuId returns the DuplicateOfSkuId field if non-nil, zero value otherwise.

### GetDuplicateOfSkuIdOk

`func (o *GetMarketProductResponseProduct) GetDuplicateOfSkuIdOk() (*string, bool)`

GetDuplicateOfSkuIdOk returns a tuple with the DuplicateOfSkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateOfSkuId

`func (o *GetMarketProductResponseProduct) SetDuplicateOfSkuId(v string)`

SetDuplicateOfSkuId sets DuplicateOfSkuId field to given value.

### HasDuplicateOfSkuId

`func (o *GetMarketProductResponseProduct) HasDuplicateOfSkuId() bool`

HasDuplicateOfSkuId returns a boolean if a field has been set.

### SetDuplicateOfSkuIdNil

`func (o *GetMarketProductResponseProduct) SetDuplicateOfSkuIdNil(b bool)`

 SetDuplicateOfSkuIdNil sets the value for DuplicateOfSkuId to be an explicit nil

### UnsetDuplicateOfSkuId
`func (o *GetMarketProductResponseProduct) UnsetDuplicateOfSkuId()`

UnsetDuplicateOfSkuId ensures that no value is present for DuplicateOfSkuId, not even an explicit nil
### GetPreferredProviderSlug

`func (o *GetMarketProductResponseProduct) GetPreferredProviderSlug() string`

GetPreferredProviderSlug returns the PreferredProviderSlug field if non-nil, zero value otherwise.

### GetPreferredProviderSlugOk

`func (o *GetMarketProductResponseProduct) GetPreferredProviderSlugOk() (*string, bool)`

GetPreferredProviderSlugOk returns a tuple with the PreferredProviderSlug field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreferredProviderSlug

`func (o *GetMarketProductResponseProduct) SetPreferredProviderSlug(v string)`

SetPreferredProviderSlug sets PreferredProviderSlug field to given value.

### HasPreferredProviderSlug

`func (o *GetMarketProductResponseProduct) HasPreferredProviderSlug() bool`

HasPreferredProviderSlug returns a boolean if a field has been set.

### SetPreferredProviderSlugNil

`func (o *GetMarketProductResponseProduct) SetPreferredProviderSlugNil(b bool)`

 SetPreferredProviderSlugNil sets the value for PreferredProviderSlug to be an explicit nil

### UnsetPreferredProviderSlug
`func (o *GetMarketProductResponseProduct) UnsetPreferredProviderSlug()`

UnsetPreferredProviderSlug ensures that no value is present for PreferredProviderSlug, not even an explicit nil
### GetAllowPreBid

`func (o *GetMarketProductResponseProduct) GetAllowPreBid() bool`

GetAllowPreBid returns the AllowPreBid field if non-nil, zero value otherwise.

### GetAllowPreBidOk

`func (o *GetMarketProductResponseProduct) GetAllowPreBidOk() (*bool, bool)`

GetAllowPreBidOk returns a tuple with the AllowPreBid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllowPreBid

`func (o *GetMarketProductResponseProduct) SetAllowPreBid(v bool)`

SetAllowPreBid sets AllowPreBid field to given value.


### GetPreBidOpensAt

`func (o *GetMarketProductResponseProduct) GetPreBidOpensAt() time.Time`

GetPreBidOpensAt returns the PreBidOpensAt field if non-nil, zero value otherwise.

### GetPreBidOpensAtOk

`func (o *GetMarketProductResponseProduct) GetPreBidOpensAtOk() (*time.Time, bool)`

GetPreBidOpensAtOk returns a tuple with the PreBidOpensAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreBidOpensAt

`func (o *GetMarketProductResponseProduct) SetPreBidOpensAt(v time.Time)`

SetPreBidOpensAt sets PreBidOpensAt field to given value.

### HasPreBidOpensAt

`func (o *GetMarketProductResponseProduct) HasPreBidOpensAt() bool`

HasPreBidOpensAt returns a boolean if a field has been set.

### SetPreBidOpensAtNil

`func (o *GetMarketProductResponseProduct) SetPreBidOpensAtNil(b bool)`

 SetPreBidOpensAtNil sets the value for PreBidOpensAt to be an explicit nil

### UnsetPreBidOpensAt
`func (o *GetMarketProductResponseProduct) UnsetPreBidOpensAt()`

UnsetPreBidOpensAt ensures that no value is present for PreBidOpensAt, not even an explicit nil
### GetFeatured

`func (o *GetMarketProductResponseProduct) GetFeatured() bool`

GetFeatured returns the Featured field if non-nil, zero value otherwise.

### GetFeaturedOk

`func (o *GetMarketProductResponseProduct) GetFeaturedOk() (*bool, bool)`

GetFeaturedOk returns a tuple with the Featured field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFeatured

`func (o *GetMarketProductResponseProduct) SetFeatured(v bool)`

SetFeatured sets Featured field to given value.


### GetExternalSource

`func (o *GetMarketProductResponseProduct) GetExternalSource() string`

GetExternalSource returns the ExternalSource field if non-nil, zero value otherwise.

### GetExternalSourceOk

`func (o *GetMarketProductResponseProduct) GetExternalSourceOk() (*string, bool)`

GetExternalSourceOk returns a tuple with the ExternalSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSource

`func (o *GetMarketProductResponseProduct) SetExternalSource(v string)`

SetExternalSource sets ExternalSource field to given value.

### HasExternalSource

`func (o *GetMarketProductResponseProduct) HasExternalSource() bool`

HasExternalSource returns a boolean if a field has been set.

### SetExternalSourceNil

`func (o *GetMarketProductResponseProduct) SetExternalSourceNil(b bool)`

 SetExternalSourceNil sets the value for ExternalSource to be an explicit nil

### UnsetExternalSource
`func (o *GetMarketProductResponseProduct) UnsetExternalSource()`

UnsetExternalSource ensures that no value is present for ExternalSource, not even an explicit nil
### GetExternalSyncedAt

`func (o *GetMarketProductResponseProduct) GetExternalSyncedAt() time.Time`

GetExternalSyncedAt returns the ExternalSyncedAt field if non-nil, zero value otherwise.

### GetExternalSyncedAtOk

`func (o *GetMarketProductResponseProduct) GetExternalSyncedAtOk() (*time.Time, bool)`

GetExternalSyncedAtOk returns a tuple with the ExternalSyncedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalSyncedAt

`func (o *GetMarketProductResponseProduct) SetExternalSyncedAt(v time.Time)`

SetExternalSyncedAt sets ExternalSyncedAt field to given value.

### HasExternalSyncedAt

`func (o *GetMarketProductResponseProduct) HasExternalSyncedAt() bool`

HasExternalSyncedAt returns a boolean if a field has been set.

### SetExternalSyncedAtNil

`func (o *GetMarketProductResponseProduct) SetExternalSyncedAtNil(b bool)`

 SetExternalSyncedAtNil sets the value for ExternalSyncedAt to be an explicit nil

### UnsetExternalSyncedAt
`func (o *GetMarketProductResponseProduct) UnsetExternalSyncedAt()`

UnsetExternalSyncedAt ensures that no value is present for ExternalSyncedAt, not even an explicit nil
### GetGradedGradeKeys

`func (o *GetMarketProductResponseProduct) GetGradedGradeKeys() []string`

GetGradedGradeKeys returns the GradedGradeKeys field if non-nil, zero value otherwise.

### GetGradedGradeKeysOk

`func (o *GetMarketProductResponseProduct) GetGradedGradeKeysOk() (*[]string, bool)`

GetGradedGradeKeysOk returns a tuple with the GradedGradeKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradedGradeKeys

`func (o *GetMarketProductResponseProduct) SetGradedGradeKeys(v []string)`

SetGradedGradeKeys sets GradedGradeKeys field to given value.


### GetLowestAskCents

`func (o *GetMarketProductResponseProduct) GetLowestAskCents() float32`

GetLowestAskCents returns the LowestAskCents field if non-nil, zero value otherwise.

### GetLowestAskCentsOk

`func (o *GetMarketProductResponseProduct) GetLowestAskCentsOk() (*float32, bool)`

GetLowestAskCentsOk returns a tuple with the LowestAskCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLowestAskCents

`func (o *GetMarketProductResponseProduct) SetLowestAskCents(v float32)`

SetLowestAskCents sets LowestAskCents field to given value.

### HasLowestAskCents

`func (o *GetMarketProductResponseProduct) HasLowestAskCents() bool`

HasLowestAskCents returns a boolean if a field has been set.

### SetLowestAskCentsNil

`func (o *GetMarketProductResponseProduct) SetLowestAskCentsNil(b bool)`

 SetLowestAskCentsNil sets the value for LowestAskCents to be an explicit nil

### UnsetLowestAskCents
`func (o *GetMarketProductResponseProduct) UnsetLowestAskCents()`

UnsetLowestAskCents ensures that no value is present for LowestAskCents, not even an explicit nil
### GetHighestBidCents

`func (o *GetMarketProductResponseProduct) GetHighestBidCents() float32`

GetHighestBidCents returns the HighestBidCents field if non-nil, zero value otherwise.

### GetHighestBidCentsOk

`func (o *GetMarketProductResponseProduct) GetHighestBidCentsOk() (*float32, bool)`

GetHighestBidCentsOk returns a tuple with the HighestBidCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHighestBidCents

`func (o *GetMarketProductResponseProduct) SetHighestBidCents(v float32)`

SetHighestBidCents sets HighestBidCents field to given value.

### HasHighestBidCents

`func (o *GetMarketProductResponseProduct) HasHighestBidCents() bool`

HasHighestBidCents returns a boolean if a field has been set.

### SetHighestBidCentsNil

`func (o *GetMarketProductResponseProduct) SetHighestBidCentsNil(b bool)`

 SetHighestBidCentsNil sets the value for HighestBidCents to be an explicit nil

### UnsetHighestBidCents
`func (o *GetMarketProductResponseProduct) UnsetHighestBidCents()`

UnsetHighestBidCents ensures that no value is present for HighestBidCents, not even an explicit nil
### GetLastSaleCents

`func (o *GetMarketProductResponseProduct) GetLastSaleCents() float32`

GetLastSaleCents returns the LastSaleCents field if non-nil, zero value otherwise.

### GetLastSaleCentsOk

`func (o *GetMarketProductResponseProduct) GetLastSaleCentsOk() (*float32, bool)`

GetLastSaleCentsOk returns a tuple with the LastSaleCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSaleCents

`func (o *GetMarketProductResponseProduct) SetLastSaleCents(v float32)`

SetLastSaleCents sets LastSaleCents field to given value.

### HasLastSaleCents

`func (o *GetMarketProductResponseProduct) HasLastSaleCents() bool`

HasLastSaleCents returns a boolean if a field has been set.

### SetLastSaleCentsNil

`func (o *GetMarketProductResponseProduct) SetLastSaleCentsNil(b bool)`

 SetLastSaleCentsNil sets the value for LastSaleCents to be an explicit nil

### UnsetLastSaleCents
`func (o *GetMarketProductResponseProduct) UnsetLastSaleCents()`

UnsetLastSaleCents ensures that no value is present for LastSaleCents, not even an explicit nil
### GetTradesCount

`func (o *GetMarketProductResponseProduct) GetTradesCount() float32`

GetTradesCount returns the TradesCount field if non-nil, zero value otherwise.

### GetTradesCountOk

`func (o *GetMarketProductResponseProduct) GetTradesCountOk() (*float32, bool)`

GetTradesCountOk returns a tuple with the TradesCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTradesCount

`func (o *GetMarketProductResponseProduct) SetTradesCount(v float32)`

SetTradesCount sets TradesCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


