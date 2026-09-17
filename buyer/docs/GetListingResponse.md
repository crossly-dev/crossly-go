# GetListingResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformListings** | [**[]GetListingResponsePlatformListings**](GetListingResponsePlatformListings.md) |  | 
**Id** | **string** |  | 
**Brand** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Sku** | Pointer to **NullableString** |  | [optional] 
**Quantity** | **float32** |  | 
**QuantityAvailable** | **float32** |  | 
**Condition** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableString** |  | [optional] 
**Material** | Pointer to **NullableString** |  | [optional] 
**Style** | Pointer to **NullableString** |  | [optional] 
**Pattern** | Pointer to **NullableString** |  | [optional] 
**Department** | Pointer to **NullableString** |  | [optional] 
**Gender** | Pointer to **NullableString** |  | [optional] 
**ItemType** | Pointer to **NullableString** |  | [optional] 
**SizeSystem** | Pointer to **NullableString** |  | [optional] 
**Color** | **[]string** |  | 
**Images** | Pointer to **[]string** |  | [optional] 
**VideoUrl** | Pointer to **NullableString** |  | [optional] 
**Tags** | **[]string** |  | 
**PriceFloorCents** | Pointer to **NullableFloat32** |  | [optional] 
**FloorIsNet** | **bool** |  | 
**WeightLb** | Pointer to **NullableString** |  | [optional] 
**WeightOz** | Pointer to **NullableString** |  | [optional] 
**DimensionLIn** | Pointer to **NullableString** |  | [optional] 
**DimensionWIn** | Pointer to **NullableString** |  | [optional] 
**DimensionHIn** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**Source** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**DescriptionHtml** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableString** |  | [optional] 
**GradeKey** | Pointer to **NullableString** |  | [optional] 
**Grading** | Pointer to [**NullableListListingsItemGrading**](ListListingsItemGrading.md) |  | [optional] 
**PublishAt** | Pointer to **NullableTime** |  | [optional] 
**ScheduledPlatforms** | Pointer to **[]string** |  | [optional] 
**ParentListingId** | Pointer to **NullableString** |  | [optional] 
**IsBundle** | **bool** |  | 
**AutomationAssignedRuleIds** | **[]string** |  | 
**AutomationBlockedRuleIds** | **[]string** |  | 
**AutomationAssignedChainIds** | **[]string** |  | 
**AutomationBlockedChainIds** | **[]string** |  | 
**HsCode** | Pointer to **NullableString** |  | [optional] 
**CountryOfOrigin** | Pointer to **NullableString** |  | [optional] 
**DelistedAt** | Pointer to **NullableTime** |  | [optional] 
**SoldAt** | Pointer to **NullableTime** |  | [optional] 
**DuplicateOfListingId** | Pointer to **NullableString** |  | [optional] 
**ClientDraftId** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetListingResponse

`func NewGetListingResponse(platformListings []GetListingResponsePlatformListings, id string, createdAt time.Time, updatedAt time.Time, userId string, quantity float32, quantityAvailable float32, color []string, tags []string, floorIsNet bool, status string, source string, isBundle bool, automationAssignedRuleIds []string, automationBlockedRuleIds []string, automationAssignedChainIds []string, automationBlockedChainIds []string, ) *GetListingResponse`

NewGetListingResponse instantiates a new GetListingResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetListingResponseWithDefaults

`func NewGetListingResponseWithDefaults() *GetListingResponse`

NewGetListingResponseWithDefaults instantiates a new GetListingResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformListings

`func (o *GetListingResponse) GetPlatformListings() []GetListingResponsePlatformListings`

GetPlatformListings returns the PlatformListings field if non-nil, zero value otherwise.

### GetPlatformListingsOk

`func (o *GetListingResponse) GetPlatformListingsOk() (*[]GetListingResponsePlatformListings, bool)`

GetPlatformListingsOk returns a tuple with the PlatformListings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformListings

`func (o *GetListingResponse) SetPlatformListings(v []GetListingResponsePlatformListings)`

SetPlatformListings sets PlatformListings field to given value.


### GetId

`func (o *GetListingResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetListingResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetListingResponse) SetId(v string)`

SetId sets Id field to given value.


### GetBrand

`func (o *GetListingResponse) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *GetListingResponse) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *GetListingResponse) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *GetListingResponse) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *GetListingResponse) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *GetListingResponse) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetName

`func (o *GetListingResponse) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetListingResponse) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetListingResponse) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetListingResponse) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *GetListingResponse) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *GetListingResponse) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetCreatedAt

`func (o *GetListingResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetListingResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetListingResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetListingResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetListingResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetListingResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *GetListingResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetListingResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetListingResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetSku

`func (o *GetListingResponse) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *GetListingResponse) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *GetListingResponse) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *GetListingResponse) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *GetListingResponse) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *GetListingResponse) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetQuantity

`func (o *GetListingResponse) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GetListingResponse) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GetListingResponse) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetQuantityAvailable

`func (o *GetListingResponse) GetQuantityAvailable() float32`

GetQuantityAvailable returns the QuantityAvailable field if non-nil, zero value otherwise.

### GetQuantityAvailableOk

`func (o *GetListingResponse) GetQuantityAvailableOk() (*float32, bool)`

GetQuantityAvailableOk returns a tuple with the QuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAvailable

`func (o *GetListingResponse) SetQuantityAvailable(v float32)`

SetQuantityAvailable sets QuantityAvailable field to given value.


### GetCondition

`func (o *GetListingResponse) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *GetListingResponse) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *GetListingResponse) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *GetListingResponse) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *GetListingResponse) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *GetListingResponse) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetSize

`func (o *GetListingResponse) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetListingResponse) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetListingResponse) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetListingResponse) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *GetListingResponse) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *GetListingResponse) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetMaterial

`func (o *GetListingResponse) GetMaterial() string`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *GetListingResponse) GetMaterialOk() (*string, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *GetListingResponse) SetMaterial(v string)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *GetListingResponse) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### SetMaterialNil

`func (o *GetListingResponse) SetMaterialNil(b bool)`

 SetMaterialNil sets the value for Material to be an explicit nil

### UnsetMaterial
`func (o *GetListingResponse) UnsetMaterial()`

UnsetMaterial ensures that no value is present for Material, not even an explicit nil
### GetStyle

`func (o *GetListingResponse) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *GetListingResponse) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *GetListingResponse) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *GetListingResponse) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *GetListingResponse) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *GetListingResponse) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetPattern

`func (o *GetListingResponse) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *GetListingResponse) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *GetListingResponse) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *GetListingResponse) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPatternNil

`func (o *GetListingResponse) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *GetListingResponse) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil
### GetDepartment

`func (o *GetListingResponse) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *GetListingResponse) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *GetListingResponse) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *GetListingResponse) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *GetListingResponse) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *GetListingResponse) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetGender

`func (o *GetListingResponse) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *GetListingResponse) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *GetListingResponse) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *GetListingResponse) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *GetListingResponse) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *GetListingResponse) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetItemType

`func (o *GetListingResponse) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *GetListingResponse) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *GetListingResponse) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *GetListingResponse) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *GetListingResponse) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *GetListingResponse) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetSizeSystem

`func (o *GetListingResponse) GetSizeSystem() string`

GetSizeSystem returns the SizeSystem field if non-nil, zero value otherwise.

### GetSizeSystemOk

`func (o *GetListingResponse) GetSizeSystemOk() (*string, bool)`

GetSizeSystemOk returns a tuple with the SizeSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSystem

`func (o *GetListingResponse) SetSizeSystem(v string)`

SetSizeSystem sets SizeSystem field to given value.

### HasSizeSystem

`func (o *GetListingResponse) HasSizeSystem() bool`

HasSizeSystem returns a boolean if a field has been set.

### SetSizeSystemNil

`func (o *GetListingResponse) SetSizeSystemNil(b bool)`

 SetSizeSystemNil sets the value for SizeSystem to be an explicit nil

### UnsetSizeSystem
`func (o *GetListingResponse) UnsetSizeSystem()`

UnsetSizeSystem ensures that no value is present for SizeSystem, not even an explicit nil
### GetColor

`func (o *GetListingResponse) GetColor() []string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *GetListingResponse) GetColorOk() (*[]string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *GetListingResponse) SetColor(v []string)`

SetColor sets Color field to given value.


### GetImages

`func (o *GetListingResponse) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *GetListingResponse) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *GetListingResponse) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *GetListingResponse) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *GetListingResponse) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *GetListingResponse) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetVideoUrl

`func (o *GetListingResponse) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *GetListingResponse) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *GetListingResponse) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *GetListingResponse) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### SetVideoUrlNil

`func (o *GetListingResponse) SetVideoUrlNil(b bool)`

 SetVideoUrlNil sets the value for VideoUrl to be an explicit nil

### UnsetVideoUrl
`func (o *GetListingResponse) UnsetVideoUrl()`

UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
### GetTags

`func (o *GetListingResponse) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GetListingResponse) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GetListingResponse) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetPriceFloorCents

`func (o *GetListingResponse) GetPriceFloorCents() float32`

GetPriceFloorCents returns the PriceFloorCents field if non-nil, zero value otherwise.

### GetPriceFloorCentsOk

`func (o *GetListingResponse) GetPriceFloorCentsOk() (*float32, bool)`

GetPriceFloorCentsOk returns a tuple with the PriceFloorCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFloorCents

`func (o *GetListingResponse) SetPriceFloorCents(v float32)`

SetPriceFloorCents sets PriceFloorCents field to given value.

### HasPriceFloorCents

`func (o *GetListingResponse) HasPriceFloorCents() bool`

HasPriceFloorCents returns a boolean if a field has been set.

### SetPriceFloorCentsNil

`func (o *GetListingResponse) SetPriceFloorCentsNil(b bool)`

 SetPriceFloorCentsNil sets the value for PriceFloorCents to be an explicit nil

### UnsetPriceFloorCents
`func (o *GetListingResponse) UnsetPriceFloorCents()`

UnsetPriceFloorCents ensures that no value is present for PriceFloorCents, not even an explicit nil
### GetFloorIsNet

`func (o *GetListingResponse) GetFloorIsNet() bool`

GetFloorIsNet returns the FloorIsNet field if non-nil, zero value otherwise.

### GetFloorIsNetOk

`func (o *GetListingResponse) GetFloorIsNetOk() (*bool, bool)`

GetFloorIsNetOk returns a tuple with the FloorIsNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorIsNet

`func (o *GetListingResponse) SetFloorIsNet(v bool)`

SetFloorIsNet sets FloorIsNet field to given value.


### GetWeightLb

`func (o *GetListingResponse) GetWeightLb() string`

GetWeightLb returns the WeightLb field if non-nil, zero value otherwise.

### GetWeightLbOk

`func (o *GetListingResponse) GetWeightLbOk() (*string, bool)`

GetWeightLbOk returns a tuple with the WeightLb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightLb

`func (o *GetListingResponse) SetWeightLb(v string)`

SetWeightLb sets WeightLb field to given value.

### HasWeightLb

`func (o *GetListingResponse) HasWeightLb() bool`

HasWeightLb returns a boolean if a field has been set.

### SetWeightLbNil

`func (o *GetListingResponse) SetWeightLbNil(b bool)`

 SetWeightLbNil sets the value for WeightLb to be an explicit nil

### UnsetWeightLb
`func (o *GetListingResponse) UnsetWeightLb()`

UnsetWeightLb ensures that no value is present for WeightLb, not even an explicit nil
### GetWeightOz

`func (o *GetListingResponse) GetWeightOz() string`

GetWeightOz returns the WeightOz field if non-nil, zero value otherwise.

### GetWeightOzOk

`func (o *GetListingResponse) GetWeightOzOk() (*string, bool)`

GetWeightOzOk returns a tuple with the WeightOz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightOz

`func (o *GetListingResponse) SetWeightOz(v string)`

SetWeightOz sets WeightOz field to given value.

### HasWeightOz

`func (o *GetListingResponse) HasWeightOz() bool`

HasWeightOz returns a boolean if a field has been set.

### SetWeightOzNil

`func (o *GetListingResponse) SetWeightOzNil(b bool)`

 SetWeightOzNil sets the value for WeightOz to be an explicit nil

### UnsetWeightOz
`func (o *GetListingResponse) UnsetWeightOz()`

UnsetWeightOz ensures that no value is present for WeightOz, not even an explicit nil
### GetDimensionLIn

`func (o *GetListingResponse) GetDimensionLIn() string`

GetDimensionLIn returns the DimensionLIn field if non-nil, zero value otherwise.

### GetDimensionLInOk

`func (o *GetListingResponse) GetDimensionLInOk() (*string, bool)`

GetDimensionLInOk returns a tuple with the DimensionLIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionLIn

`func (o *GetListingResponse) SetDimensionLIn(v string)`

SetDimensionLIn sets DimensionLIn field to given value.

### HasDimensionLIn

`func (o *GetListingResponse) HasDimensionLIn() bool`

HasDimensionLIn returns a boolean if a field has been set.

### SetDimensionLInNil

`func (o *GetListingResponse) SetDimensionLInNil(b bool)`

 SetDimensionLInNil sets the value for DimensionLIn to be an explicit nil

### UnsetDimensionLIn
`func (o *GetListingResponse) UnsetDimensionLIn()`

UnsetDimensionLIn ensures that no value is present for DimensionLIn, not even an explicit nil
### GetDimensionWIn

`func (o *GetListingResponse) GetDimensionWIn() string`

GetDimensionWIn returns the DimensionWIn field if non-nil, zero value otherwise.

### GetDimensionWInOk

`func (o *GetListingResponse) GetDimensionWInOk() (*string, bool)`

GetDimensionWInOk returns a tuple with the DimensionWIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionWIn

`func (o *GetListingResponse) SetDimensionWIn(v string)`

SetDimensionWIn sets DimensionWIn field to given value.

### HasDimensionWIn

`func (o *GetListingResponse) HasDimensionWIn() bool`

HasDimensionWIn returns a boolean if a field has been set.

### SetDimensionWInNil

`func (o *GetListingResponse) SetDimensionWInNil(b bool)`

 SetDimensionWInNil sets the value for DimensionWIn to be an explicit nil

### UnsetDimensionWIn
`func (o *GetListingResponse) UnsetDimensionWIn()`

UnsetDimensionWIn ensures that no value is present for DimensionWIn, not even an explicit nil
### GetDimensionHIn

`func (o *GetListingResponse) GetDimensionHIn() string`

GetDimensionHIn returns the DimensionHIn field if non-nil, zero value otherwise.

### GetDimensionHInOk

`func (o *GetListingResponse) GetDimensionHInOk() (*string, bool)`

GetDimensionHInOk returns a tuple with the DimensionHIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionHIn

`func (o *GetListingResponse) SetDimensionHIn(v string)`

SetDimensionHIn sets DimensionHIn field to given value.

### HasDimensionHIn

`func (o *GetListingResponse) HasDimensionHIn() bool`

HasDimensionHIn returns a boolean if a field has been set.

### SetDimensionHInNil

`func (o *GetListingResponse) SetDimensionHInNil(b bool)`

 SetDimensionHInNil sets the value for DimensionHIn to be an explicit nil

### UnsetDimensionHIn
`func (o *GetListingResponse) UnsetDimensionHIn()`

UnsetDimensionHIn ensures that no value is present for DimensionHIn, not even an explicit nil
### GetStatus

`func (o *GetListingResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetListingResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetListingResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSource

`func (o *GetListingResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetListingResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetListingResponse) SetSource(v string)`

SetSource sets Source field to given value.


### GetDescription

`func (o *GetListingResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetListingResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetListingResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetListingResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetListingResponse) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetListingResponse) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetTitle

`func (o *GetListingResponse) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *GetListingResponse) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *GetListingResponse) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *GetListingResponse) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *GetListingResponse) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *GetListingResponse) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescriptionHtml

`func (o *GetListingResponse) GetDescriptionHtml() string`

GetDescriptionHtml returns the DescriptionHtml field if non-nil, zero value otherwise.

### GetDescriptionHtmlOk

`func (o *GetListingResponse) GetDescriptionHtmlOk() (*string, bool)`

GetDescriptionHtmlOk returns a tuple with the DescriptionHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionHtml

`func (o *GetListingResponse) SetDescriptionHtml(v string)`

SetDescriptionHtml sets DescriptionHtml field to given value.

### HasDescriptionHtml

`func (o *GetListingResponse) HasDescriptionHtml() bool`

HasDescriptionHtml returns a boolean if a field has been set.

### SetDescriptionHtmlNil

`func (o *GetListingResponse) SetDescriptionHtmlNil(b bool)`

 SetDescriptionHtmlNil sets the value for DescriptionHtml to be an explicit nil

### UnsetDescriptionHtml
`func (o *GetListingResponse) UnsetDescriptionHtml()`

UnsetDescriptionHtml ensures that no value is present for DescriptionHtml, not even an explicit nil
### GetPrice

`func (o *GetListingResponse) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetListingResponse) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetListingResponse) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetListingResponse) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *GetListingResponse) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *GetListingResponse) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetGradeKey

`func (o *GetListingResponse) GetGradeKey() string`

GetGradeKey returns the GradeKey field if non-nil, zero value otherwise.

### GetGradeKeyOk

`func (o *GetListingResponse) GetGradeKeyOk() (*string, bool)`

GetGradeKeyOk returns a tuple with the GradeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeKey

`func (o *GetListingResponse) SetGradeKey(v string)`

SetGradeKey sets GradeKey field to given value.

### HasGradeKey

`func (o *GetListingResponse) HasGradeKey() bool`

HasGradeKey returns a boolean if a field has been set.

### SetGradeKeyNil

`func (o *GetListingResponse) SetGradeKeyNil(b bool)`

 SetGradeKeyNil sets the value for GradeKey to be an explicit nil

### UnsetGradeKey
`func (o *GetListingResponse) UnsetGradeKey()`

UnsetGradeKey ensures that no value is present for GradeKey, not even an explicit nil
### GetGrading

`func (o *GetListingResponse) GetGrading() ListListingsItemGrading`

GetGrading returns the Grading field if non-nil, zero value otherwise.

### GetGradingOk

`func (o *GetListingResponse) GetGradingOk() (*ListListingsItemGrading, bool)`

GetGradingOk returns a tuple with the Grading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrading

`func (o *GetListingResponse) SetGrading(v ListListingsItemGrading)`

SetGrading sets Grading field to given value.

### HasGrading

`func (o *GetListingResponse) HasGrading() bool`

HasGrading returns a boolean if a field has been set.

### SetGradingNil

`func (o *GetListingResponse) SetGradingNil(b bool)`

 SetGradingNil sets the value for Grading to be an explicit nil

### UnsetGrading
`func (o *GetListingResponse) UnsetGrading()`

UnsetGrading ensures that no value is present for Grading, not even an explicit nil
### GetPublishAt

`func (o *GetListingResponse) GetPublishAt() time.Time`

GetPublishAt returns the PublishAt field if non-nil, zero value otherwise.

### GetPublishAtOk

`func (o *GetListingResponse) GetPublishAtOk() (*time.Time, bool)`

GetPublishAtOk returns a tuple with the PublishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishAt

`func (o *GetListingResponse) SetPublishAt(v time.Time)`

SetPublishAt sets PublishAt field to given value.

### HasPublishAt

`func (o *GetListingResponse) HasPublishAt() bool`

HasPublishAt returns a boolean if a field has been set.

### SetPublishAtNil

`func (o *GetListingResponse) SetPublishAtNil(b bool)`

 SetPublishAtNil sets the value for PublishAt to be an explicit nil

### UnsetPublishAt
`func (o *GetListingResponse) UnsetPublishAt()`

UnsetPublishAt ensures that no value is present for PublishAt, not even an explicit nil
### GetScheduledPlatforms

`func (o *GetListingResponse) GetScheduledPlatforms() []string`

GetScheduledPlatforms returns the ScheduledPlatforms field if non-nil, zero value otherwise.

### GetScheduledPlatformsOk

`func (o *GetListingResponse) GetScheduledPlatformsOk() (*[]string, bool)`

GetScheduledPlatformsOk returns a tuple with the ScheduledPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPlatforms

`func (o *GetListingResponse) SetScheduledPlatforms(v []string)`

SetScheduledPlatforms sets ScheduledPlatforms field to given value.

### HasScheduledPlatforms

`func (o *GetListingResponse) HasScheduledPlatforms() bool`

HasScheduledPlatforms returns a boolean if a field has been set.

### SetScheduledPlatformsNil

`func (o *GetListingResponse) SetScheduledPlatformsNil(b bool)`

 SetScheduledPlatformsNil sets the value for ScheduledPlatforms to be an explicit nil

### UnsetScheduledPlatforms
`func (o *GetListingResponse) UnsetScheduledPlatforms()`

UnsetScheduledPlatforms ensures that no value is present for ScheduledPlatforms, not even an explicit nil
### GetParentListingId

`func (o *GetListingResponse) GetParentListingId() string`

GetParentListingId returns the ParentListingId field if non-nil, zero value otherwise.

### GetParentListingIdOk

`func (o *GetListingResponse) GetParentListingIdOk() (*string, bool)`

GetParentListingIdOk returns a tuple with the ParentListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentListingId

`func (o *GetListingResponse) SetParentListingId(v string)`

SetParentListingId sets ParentListingId field to given value.

### HasParentListingId

`func (o *GetListingResponse) HasParentListingId() bool`

HasParentListingId returns a boolean if a field has been set.

### SetParentListingIdNil

`func (o *GetListingResponse) SetParentListingIdNil(b bool)`

 SetParentListingIdNil sets the value for ParentListingId to be an explicit nil

### UnsetParentListingId
`func (o *GetListingResponse) UnsetParentListingId()`

UnsetParentListingId ensures that no value is present for ParentListingId, not even an explicit nil
### GetIsBundle

`func (o *GetListingResponse) GetIsBundle() bool`

GetIsBundle returns the IsBundle field if non-nil, zero value otherwise.

### GetIsBundleOk

`func (o *GetListingResponse) GetIsBundleOk() (*bool, bool)`

GetIsBundleOk returns a tuple with the IsBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBundle

`func (o *GetListingResponse) SetIsBundle(v bool)`

SetIsBundle sets IsBundle field to given value.


### GetAutomationAssignedRuleIds

`func (o *GetListingResponse) GetAutomationAssignedRuleIds() []string`

GetAutomationAssignedRuleIds returns the AutomationAssignedRuleIds field if non-nil, zero value otherwise.

### GetAutomationAssignedRuleIdsOk

`func (o *GetListingResponse) GetAutomationAssignedRuleIdsOk() (*[]string, bool)`

GetAutomationAssignedRuleIdsOk returns a tuple with the AutomationAssignedRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationAssignedRuleIds

`func (o *GetListingResponse) SetAutomationAssignedRuleIds(v []string)`

SetAutomationAssignedRuleIds sets AutomationAssignedRuleIds field to given value.


### GetAutomationBlockedRuleIds

`func (o *GetListingResponse) GetAutomationBlockedRuleIds() []string`

GetAutomationBlockedRuleIds returns the AutomationBlockedRuleIds field if non-nil, zero value otherwise.

### GetAutomationBlockedRuleIdsOk

`func (o *GetListingResponse) GetAutomationBlockedRuleIdsOk() (*[]string, bool)`

GetAutomationBlockedRuleIdsOk returns a tuple with the AutomationBlockedRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationBlockedRuleIds

`func (o *GetListingResponse) SetAutomationBlockedRuleIds(v []string)`

SetAutomationBlockedRuleIds sets AutomationBlockedRuleIds field to given value.


### GetAutomationAssignedChainIds

`func (o *GetListingResponse) GetAutomationAssignedChainIds() []string`

GetAutomationAssignedChainIds returns the AutomationAssignedChainIds field if non-nil, zero value otherwise.

### GetAutomationAssignedChainIdsOk

`func (o *GetListingResponse) GetAutomationAssignedChainIdsOk() (*[]string, bool)`

GetAutomationAssignedChainIdsOk returns a tuple with the AutomationAssignedChainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationAssignedChainIds

`func (o *GetListingResponse) SetAutomationAssignedChainIds(v []string)`

SetAutomationAssignedChainIds sets AutomationAssignedChainIds field to given value.


### GetAutomationBlockedChainIds

`func (o *GetListingResponse) GetAutomationBlockedChainIds() []string`

GetAutomationBlockedChainIds returns the AutomationBlockedChainIds field if non-nil, zero value otherwise.

### GetAutomationBlockedChainIdsOk

`func (o *GetListingResponse) GetAutomationBlockedChainIdsOk() (*[]string, bool)`

GetAutomationBlockedChainIdsOk returns a tuple with the AutomationBlockedChainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationBlockedChainIds

`func (o *GetListingResponse) SetAutomationBlockedChainIds(v []string)`

SetAutomationBlockedChainIds sets AutomationBlockedChainIds field to given value.


### GetHsCode

`func (o *GetListingResponse) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *GetListingResponse) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *GetListingResponse) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *GetListingResponse) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### SetHsCodeNil

`func (o *GetListingResponse) SetHsCodeNil(b bool)`

 SetHsCodeNil sets the value for HsCode to be an explicit nil

### UnsetHsCode
`func (o *GetListingResponse) UnsetHsCode()`

UnsetHsCode ensures that no value is present for HsCode, not even an explicit nil
### GetCountryOfOrigin

`func (o *GetListingResponse) GetCountryOfOrigin() string`

GetCountryOfOrigin returns the CountryOfOrigin field if non-nil, zero value otherwise.

### GetCountryOfOriginOk

`func (o *GetListingResponse) GetCountryOfOriginOk() (*string, bool)`

GetCountryOfOriginOk returns a tuple with the CountryOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfOrigin

`func (o *GetListingResponse) SetCountryOfOrigin(v string)`

SetCountryOfOrigin sets CountryOfOrigin field to given value.

### HasCountryOfOrigin

`func (o *GetListingResponse) HasCountryOfOrigin() bool`

HasCountryOfOrigin returns a boolean if a field has been set.

### SetCountryOfOriginNil

`func (o *GetListingResponse) SetCountryOfOriginNil(b bool)`

 SetCountryOfOriginNil sets the value for CountryOfOrigin to be an explicit nil

### UnsetCountryOfOrigin
`func (o *GetListingResponse) UnsetCountryOfOrigin()`

UnsetCountryOfOrigin ensures that no value is present for CountryOfOrigin, not even an explicit nil
### GetDelistedAt

`func (o *GetListingResponse) GetDelistedAt() time.Time`

GetDelistedAt returns the DelistedAt field if non-nil, zero value otherwise.

### GetDelistedAtOk

`func (o *GetListingResponse) GetDelistedAtOk() (*time.Time, bool)`

GetDelistedAtOk returns a tuple with the DelistedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelistedAt

`func (o *GetListingResponse) SetDelistedAt(v time.Time)`

SetDelistedAt sets DelistedAt field to given value.

### HasDelistedAt

`func (o *GetListingResponse) HasDelistedAt() bool`

HasDelistedAt returns a boolean if a field has been set.

### SetDelistedAtNil

`func (o *GetListingResponse) SetDelistedAtNil(b bool)`

 SetDelistedAtNil sets the value for DelistedAt to be an explicit nil

### UnsetDelistedAt
`func (o *GetListingResponse) UnsetDelistedAt()`

UnsetDelistedAt ensures that no value is present for DelistedAt, not even an explicit nil
### GetSoldAt

`func (o *GetListingResponse) GetSoldAt() time.Time`

GetSoldAt returns the SoldAt field if non-nil, zero value otherwise.

### GetSoldAtOk

`func (o *GetListingResponse) GetSoldAtOk() (*time.Time, bool)`

GetSoldAtOk returns a tuple with the SoldAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldAt

`func (o *GetListingResponse) SetSoldAt(v time.Time)`

SetSoldAt sets SoldAt field to given value.

### HasSoldAt

`func (o *GetListingResponse) HasSoldAt() bool`

HasSoldAt returns a boolean if a field has been set.

### SetSoldAtNil

`func (o *GetListingResponse) SetSoldAtNil(b bool)`

 SetSoldAtNil sets the value for SoldAt to be an explicit nil

### UnsetSoldAt
`func (o *GetListingResponse) UnsetSoldAt()`

UnsetSoldAt ensures that no value is present for SoldAt, not even an explicit nil
### GetDuplicateOfListingId

`func (o *GetListingResponse) GetDuplicateOfListingId() string`

GetDuplicateOfListingId returns the DuplicateOfListingId field if non-nil, zero value otherwise.

### GetDuplicateOfListingIdOk

`func (o *GetListingResponse) GetDuplicateOfListingIdOk() (*string, bool)`

GetDuplicateOfListingIdOk returns a tuple with the DuplicateOfListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateOfListingId

`func (o *GetListingResponse) SetDuplicateOfListingId(v string)`

SetDuplicateOfListingId sets DuplicateOfListingId field to given value.

### HasDuplicateOfListingId

`func (o *GetListingResponse) HasDuplicateOfListingId() bool`

HasDuplicateOfListingId returns a boolean if a field has been set.

### SetDuplicateOfListingIdNil

`func (o *GetListingResponse) SetDuplicateOfListingIdNil(b bool)`

 SetDuplicateOfListingIdNil sets the value for DuplicateOfListingId to be an explicit nil

### UnsetDuplicateOfListingId
`func (o *GetListingResponse) UnsetDuplicateOfListingId()`

UnsetDuplicateOfListingId ensures that no value is present for DuplicateOfListingId, not even an explicit nil
### GetClientDraftId

`func (o *GetListingResponse) GetClientDraftId() string`

GetClientDraftId returns the ClientDraftId field if non-nil, zero value otherwise.

### GetClientDraftIdOk

`func (o *GetListingResponse) GetClientDraftIdOk() (*string, bool)`

GetClientDraftIdOk returns a tuple with the ClientDraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDraftId

`func (o *GetListingResponse) SetClientDraftId(v string)`

SetClientDraftId sets ClientDraftId field to given value.

### HasClientDraftId

`func (o *GetListingResponse) HasClientDraftId() bool`

HasClientDraftId returns a boolean if a field has been set.

### SetClientDraftIdNil

`func (o *GetListingResponse) SetClientDraftIdNil(b bool)`

 SetClientDraftIdNil sets the value for ClientDraftId to be an explicit nil

### UnsetClientDraftId
`func (o *GetListingResponse) UnsetClientDraftId()`

UnsetClientDraftId ensures that no value is present for ClientDraftId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


