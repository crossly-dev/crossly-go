# ListInventoryItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Brand** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Sku** | Pointer to **NullableString** |  | [optional] 
**Quantity** | **float32** |  | 
**QuantityAvailable** | **float32** |  | 
**QuantityReserved** | **float32** |  | 
**QuantitySold** | **float32** |  | 
**OriginalPrice** | Pointer to **NullableString** |  | [optional] 
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
**DefaultTitle** | **string** |  | 
**Images** | **[]string** |  | 
**VideoUrl** | Pointer to **NullableString** |  | [optional] 
**Tags** | **[]string** |  | 
**Labels** | **[]string** |  | 
**Notes** | Pointer to **NullableString** |  | [optional] 
**CustomLabel** | Pointer to **NullableString** |  | [optional] 
**WarehouseLocation** | Pointer to **NullableString** |  | [optional] 
**StorageBin** | Pointer to **NullableString** |  | [optional] 
**StorageNote** | Pointer to **NullableString** |  | [optional] 
**MarketSkuId** | Pointer to **NullableString** |  | [optional] 
**MarketVariantId** | Pointer to **NullableString** |  | [optional] 
**FirstListedAt** | Pointer to **NullableTime** |  | [optional] 
**IncludeTaxInLanded** | **bool** |  | 
**CostBasisCents** | Pointer to **NullableFloat32** |  | [optional] 
**PriceFloorCents** | Pointer to **NullableFloat32** |  | [optional] 
**FloorIsNet** | **bool** |  | 
**AcquiredAt** | Pointer to **NullableTime** |  | [optional] 
**AcquiredSource** | Pointer to **NullableString** |  | [optional] 
**WeightLb** | Pointer to **NullableString** |  | [optional] 
**WeightOz** | Pointer to **NullableString** |  | [optional] 
**DimensionLIn** | Pointer to **NullableString** |  | [optional] 
**DimensionWIn** | Pointer to **NullableString** |  | [optional] 
**DimensionHIn** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**Source** | **string** |  | 
**Category** | Pointer to [**NullableListInventoryItemCategory**](ListInventoryItemCategory.md) |  | [optional] 

## Methods

### NewListInventoryItem

`func NewListInventoryItem(id string, createdAt time.Time, updatedAt time.Time, userId string, quantity float32, quantityAvailable float32, quantityReserved float32, quantitySold float32, color []string, defaultTitle string, images []string, tags []string, labels []string, includeTaxInLanded bool, floorIsNet bool, status string, source string, ) *ListInventoryItem`

NewListInventoryItem instantiates a new ListInventoryItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInventoryItemWithDefaults

`func NewListInventoryItemWithDefaults() *ListInventoryItem`

NewListInventoryItemWithDefaults instantiates a new ListInventoryItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListInventoryItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListInventoryItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListInventoryItem) SetId(v string)`

SetId sets Id field to given value.


### GetBrand

`func (o *ListInventoryItem) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *ListInventoryItem) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *ListInventoryItem) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *ListInventoryItem) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *ListInventoryItem) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *ListInventoryItem) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetCreatedAt

`func (o *ListInventoryItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListInventoryItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListInventoryItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListInventoryItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListInventoryItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListInventoryItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *ListInventoryItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListInventoryItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListInventoryItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetSku

`func (o *ListInventoryItem) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *ListInventoryItem) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *ListInventoryItem) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *ListInventoryItem) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *ListInventoryItem) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *ListInventoryItem) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetQuantity

`func (o *ListInventoryItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ListInventoryItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ListInventoryItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetQuantityAvailable

`func (o *ListInventoryItem) GetQuantityAvailable() float32`

GetQuantityAvailable returns the QuantityAvailable field if non-nil, zero value otherwise.

### GetQuantityAvailableOk

`func (o *ListInventoryItem) GetQuantityAvailableOk() (*float32, bool)`

GetQuantityAvailableOk returns a tuple with the QuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAvailable

`func (o *ListInventoryItem) SetQuantityAvailable(v float32)`

SetQuantityAvailable sets QuantityAvailable field to given value.


### GetQuantityReserved

`func (o *ListInventoryItem) GetQuantityReserved() float32`

GetQuantityReserved returns the QuantityReserved field if non-nil, zero value otherwise.

### GetQuantityReservedOk

`func (o *ListInventoryItem) GetQuantityReservedOk() (*float32, bool)`

GetQuantityReservedOk returns a tuple with the QuantityReserved field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityReserved

`func (o *ListInventoryItem) SetQuantityReserved(v float32)`

SetQuantityReserved sets QuantityReserved field to given value.


### GetQuantitySold

`func (o *ListInventoryItem) GetQuantitySold() float32`

GetQuantitySold returns the QuantitySold field if non-nil, zero value otherwise.

### GetQuantitySoldOk

`func (o *ListInventoryItem) GetQuantitySoldOk() (*float32, bool)`

GetQuantitySoldOk returns a tuple with the QuantitySold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantitySold

`func (o *ListInventoryItem) SetQuantitySold(v float32)`

SetQuantitySold sets QuantitySold field to given value.


### GetOriginalPrice

`func (o *ListInventoryItem) GetOriginalPrice() string`

GetOriginalPrice returns the OriginalPrice field if non-nil, zero value otherwise.

### GetOriginalPriceOk

`func (o *ListInventoryItem) GetOriginalPriceOk() (*string, bool)`

GetOriginalPriceOk returns a tuple with the OriginalPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalPrice

`func (o *ListInventoryItem) SetOriginalPrice(v string)`

SetOriginalPrice sets OriginalPrice field to given value.

### HasOriginalPrice

`func (o *ListInventoryItem) HasOriginalPrice() bool`

HasOriginalPrice returns a boolean if a field has been set.

### SetOriginalPriceNil

`func (o *ListInventoryItem) SetOriginalPriceNil(b bool)`

 SetOriginalPriceNil sets the value for OriginalPrice to be an explicit nil

### UnsetOriginalPrice
`func (o *ListInventoryItem) UnsetOriginalPrice()`

UnsetOriginalPrice ensures that no value is present for OriginalPrice, not even an explicit nil
### GetCondition

`func (o *ListInventoryItem) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *ListInventoryItem) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *ListInventoryItem) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *ListInventoryItem) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *ListInventoryItem) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *ListInventoryItem) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetSize

`func (o *ListInventoryItem) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ListInventoryItem) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ListInventoryItem) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *ListInventoryItem) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *ListInventoryItem) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *ListInventoryItem) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetMaterial

`func (o *ListInventoryItem) GetMaterial() string`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *ListInventoryItem) GetMaterialOk() (*string, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *ListInventoryItem) SetMaterial(v string)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *ListInventoryItem) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### SetMaterialNil

`func (o *ListInventoryItem) SetMaterialNil(b bool)`

 SetMaterialNil sets the value for Material to be an explicit nil

### UnsetMaterial
`func (o *ListInventoryItem) UnsetMaterial()`

UnsetMaterial ensures that no value is present for Material, not even an explicit nil
### GetStyle

`func (o *ListInventoryItem) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *ListInventoryItem) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *ListInventoryItem) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *ListInventoryItem) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *ListInventoryItem) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *ListInventoryItem) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetPattern

`func (o *ListInventoryItem) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *ListInventoryItem) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *ListInventoryItem) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *ListInventoryItem) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPatternNil

`func (o *ListInventoryItem) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *ListInventoryItem) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil
### GetDepartment

`func (o *ListInventoryItem) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *ListInventoryItem) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *ListInventoryItem) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *ListInventoryItem) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *ListInventoryItem) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *ListInventoryItem) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetGender

`func (o *ListInventoryItem) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *ListInventoryItem) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *ListInventoryItem) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *ListInventoryItem) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *ListInventoryItem) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *ListInventoryItem) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetItemType

`func (o *ListInventoryItem) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *ListInventoryItem) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *ListInventoryItem) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *ListInventoryItem) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *ListInventoryItem) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *ListInventoryItem) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetSizeSystem

`func (o *ListInventoryItem) GetSizeSystem() string`

GetSizeSystem returns the SizeSystem field if non-nil, zero value otherwise.

### GetSizeSystemOk

`func (o *ListInventoryItem) GetSizeSystemOk() (*string, bool)`

GetSizeSystemOk returns a tuple with the SizeSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSystem

`func (o *ListInventoryItem) SetSizeSystem(v string)`

SetSizeSystem sets SizeSystem field to given value.

### HasSizeSystem

`func (o *ListInventoryItem) HasSizeSystem() bool`

HasSizeSystem returns a boolean if a field has been set.

### SetSizeSystemNil

`func (o *ListInventoryItem) SetSizeSystemNil(b bool)`

 SetSizeSystemNil sets the value for SizeSystem to be an explicit nil

### UnsetSizeSystem
`func (o *ListInventoryItem) UnsetSizeSystem()`

UnsetSizeSystem ensures that no value is present for SizeSystem, not even an explicit nil
### GetColor

`func (o *ListInventoryItem) GetColor() []string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *ListInventoryItem) GetColorOk() (*[]string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *ListInventoryItem) SetColor(v []string)`

SetColor sets Color field to given value.


### GetDefaultTitle

`func (o *ListInventoryItem) GetDefaultTitle() string`

GetDefaultTitle returns the DefaultTitle field if non-nil, zero value otherwise.

### GetDefaultTitleOk

`func (o *ListInventoryItem) GetDefaultTitleOk() (*string, bool)`

GetDefaultTitleOk returns a tuple with the DefaultTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultTitle

`func (o *ListInventoryItem) SetDefaultTitle(v string)`

SetDefaultTitle sets DefaultTitle field to given value.


### GetImages

`func (o *ListInventoryItem) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *ListInventoryItem) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *ListInventoryItem) SetImages(v []string)`

SetImages sets Images field to given value.


### GetVideoUrl

`func (o *ListInventoryItem) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *ListInventoryItem) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *ListInventoryItem) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *ListInventoryItem) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### SetVideoUrlNil

`func (o *ListInventoryItem) SetVideoUrlNil(b bool)`

 SetVideoUrlNil sets the value for VideoUrl to be an explicit nil

### UnsetVideoUrl
`func (o *ListInventoryItem) UnsetVideoUrl()`

UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
### GetTags

`func (o *ListInventoryItem) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ListInventoryItem) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ListInventoryItem) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetLabels

`func (o *ListInventoryItem) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListInventoryItem) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListInventoryItem) SetLabels(v []string)`

SetLabels sets Labels field to given value.


### GetNotes

`func (o *ListInventoryItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ListInventoryItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ListInventoryItem) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ListInventoryItem) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ListInventoryItem) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ListInventoryItem) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetCustomLabel

`func (o *ListInventoryItem) GetCustomLabel() string`

GetCustomLabel returns the CustomLabel field if non-nil, zero value otherwise.

### GetCustomLabelOk

`func (o *ListInventoryItem) GetCustomLabelOk() (*string, bool)`

GetCustomLabelOk returns a tuple with the CustomLabel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomLabel

`func (o *ListInventoryItem) SetCustomLabel(v string)`

SetCustomLabel sets CustomLabel field to given value.

### HasCustomLabel

`func (o *ListInventoryItem) HasCustomLabel() bool`

HasCustomLabel returns a boolean if a field has been set.

### SetCustomLabelNil

`func (o *ListInventoryItem) SetCustomLabelNil(b bool)`

 SetCustomLabelNil sets the value for CustomLabel to be an explicit nil

### UnsetCustomLabel
`func (o *ListInventoryItem) UnsetCustomLabel()`

UnsetCustomLabel ensures that no value is present for CustomLabel, not even an explicit nil
### GetWarehouseLocation

`func (o *ListInventoryItem) GetWarehouseLocation() string`

GetWarehouseLocation returns the WarehouseLocation field if non-nil, zero value otherwise.

### GetWarehouseLocationOk

`func (o *ListInventoryItem) GetWarehouseLocationOk() (*string, bool)`

GetWarehouseLocationOk returns a tuple with the WarehouseLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWarehouseLocation

`func (o *ListInventoryItem) SetWarehouseLocation(v string)`

SetWarehouseLocation sets WarehouseLocation field to given value.

### HasWarehouseLocation

`func (o *ListInventoryItem) HasWarehouseLocation() bool`

HasWarehouseLocation returns a boolean if a field has been set.

### SetWarehouseLocationNil

`func (o *ListInventoryItem) SetWarehouseLocationNil(b bool)`

 SetWarehouseLocationNil sets the value for WarehouseLocation to be an explicit nil

### UnsetWarehouseLocation
`func (o *ListInventoryItem) UnsetWarehouseLocation()`

UnsetWarehouseLocation ensures that no value is present for WarehouseLocation, not even an explicit nil
### GetStorageBin

`func (o *ListInventoryItem) GetStorageBin() string`

GetStorageBin returns the StorageBin field if non-nil, zero value otherwise.

### GetStorageBinOk

`func (o *ListInventoryItem) GetStorageBinOk() (*string, bool)`

GetStorageBinOk returns a tuple with the StorageBin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBin

`func (o *ListInventoryItem) SetStorageBin(v string)`

SetStorageBin sets StorageBin field to given value.

### HasStorageBin

`func (o *ListInventoryItem) HasStorageBin() bool`

HasStorageBin returns a boolean if a field has been set.

### SetStorageBinNil

`func (o *ListInventoryItem) SetStorageBinNil(b bool)`

 SetStorageBinNil sets the value for StorageBin to be an explicit nil

### UnsetStorageBin
`func (o *ListInventoryItem) UnsetStorageBin()`

UnsetStorageBin ensures that no value is present for StorageBin, not even an explicit nil
### GetStorageNote

`func (o *ListInventoryItem) GetStorageNote() string`

GetStorageNote returns the StorageNote field if non-nil, zero value otherwise.

### GetStorageNoteOk

`func (o *ListInventoryItem) GetStorageNoteOk() (*string, bool)`

GetStorageNoteOk returns a tuple with the StorageNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageNote

`func (o *ListInventoryItem) SetStorageNote(v string)`

SetStorageNote sets StorageNote field to given value.

### HasStorageNote

`func (o *ListInventoryItem) HasStorageNote() bool`

HasStorageNote returns a boolean if a field has been set.

### SetStorageNoteNil

`func (o *ListInventoryItem) SetStorageNoteNil(b bool)`

 SetStorageNoteNil sets the value for StorageNote to be an explicit nil

### UnsetStorageNote
`func (o *ListInventoryItem) UnsetStorageNote()`

UnsetStorageNote ensures that no value is present for StorageNote, not even an explicit nil
### GetMarketSkuId

`func (o *ListInventoryItem) GetMarketSkuId() string`

GetMarketSkuId returns the MarketSkuId field if non-nil, zero value otherwise.

### GetMarketSkuIdOk

`func (o *ListInventoryItem) GetMarketSkuIdOk() (*string, bool)`

GetMarketSkuIdOk returns a tuple with the MarketSkuId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketSkuId

`func (o *ListInventoryItem) SetMarketSkuId(v string)`

SetMarketSkuId sets MarketSkuId field to given value.

### HasMarketSkuId

`func (o *ListInventoryItem) HasMarketSkuId() bool`

HasMarketSkuId returns a boolean if a field has been set.

### SetMarketSkuIdNil

`func (o *ListInventoryItem) SetMarketSkuIdNil(b bool)`

 SetMarketSkuIdNil sets the value for MarketSkuId to be an explicit nil

### UnsetMarketSkuId
`func (o *ListInventoryItem) UnsetMarketSkuId()`

UnsetMarketSkuId ensures that no value is present for MarketSkuId, not even an explicit nil
### GetMarketVariantId

`func (o *ListInventoryItem) GetMarketVariantId() string`

GetMarketVariantId returns the MarketVariantId field if non-nil, zero value otherwise.

### GetMarketVariantIdOk

`func (o *ListInventoryItem) GetMarketVariantIdOk() (*string, bool)`

GetMarketVariantIdOk returns a tuple with the MarketVariantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarketVariantId

`func (o *ListInventoryItem) SetMarketVariantId(v string)`

SetMarketVariantId sets MarketVariantId field to given value.

### HasMarketVariantId

`func (o *ListInventoryItem) HasMarketVariantId() bool`

HasMarketVariantId returns a boolean if a field has been set.

### SetMarketVariantIdNil

`func (o *ListInventoryItem) SetMarketVariantIdNil(b bool)`

 SetMarketVariantIdNil sets the value for MarketVariantId to be an explicit nil

### UnsetMarketVariantId
`func (o *ListInventoryItem) UnsetMarketVariantId()`

UnsetMarketVariantId ensures that no value is present for MarketVariantId, not even an explicit nil
### GetFirstListedAt

`func (o *ListInventoryItem) GetFirstListedAt() time.Time`

GetFirstListedAt returns the FirstListedAt field if non-nil, zero value otherwise.

### GetFirstListedAtOk

`func (o *ListInventoryItem) GetFirstListedAtOk() (*time.Time, bool)`

GetFirstListedAtOk returns a tuple with the FirstListedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstListedAt

`func (o *ListInventoryItem) SetFirstListedAt(v time.Time)`

SetFirstListedAt sets FirstListedAt field to given value.

### HasFirstListedAt

`func (o *ListInventoryItem) HasFirstListedAt() bool`

HasFirstListedAt returns a boolean if a field has been set.

### SetFirstListedAtNil

`func (o *ListInventoryItem) SetFirstListedAtNil(b bool)`

 SetFirstListedAtNil sets the value for FirstListedAt to be an explicit nil

### UnsetFirstListedAt
`func (o *ListInventoryItem) UnsetFirstListedAt()`

UnsetFirstListedAt ensures that no value is present for FirstListedAt, not even an explicit nil
### GetIncludeTaxInLanded

`func (o *ListInventoryItem) GetIncludeTaxInLanded() bool`

GetIncludeTaxInLanded returns the IncludeTaxInLanded field if non-nil, zero value otherwise.

### GetIncludeTaxInLandedOk

`func (o *ListInventoryItem) GetIncludeTaxInLandedOk() (*bool, bool)`

GetIncludeTaxInLandedOk returns a tuple with the IncludeTaxInLanded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeTaxInLanded

`func (o *ListInventoryItem) SetIncludeTaxInLanded(v bool)`

SetIncludeTaxInLanded sets IncludeTaxInLanded field to given value.


### GetCostBasisCents

`func (o *ListInventoryItem) GetCostBasisCents() float32`

GetCostBasisCents returns the CostBasisCents field if non-nil, zero value otherwise.

### GetCostBasisCentsOk

`func (o *ListInventoryItem) GetCostBasisCentsOk() (*float32, bool)`

GetCostBasisCentsOk returns a tuple with the CostBasisCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBasisCents

`func (o *ListInventoryItem) SetCostBasisCents(v float32)`

SetCostBasisCents sets CostBasisCents field to given value.

### HasCostBasisCents

`func (o *ListInventoryItem) HasCostBasisCents() bool`

HasCostBasisCents returns a boolean if a field has been set.

### SetCostBasisCentsNil

`func (o *ListInventoryItem) SetCostBasisCentsNil(b bool)`

 SetCostBasisCentsNil sets the value for CostBasisCents to be an explicit nil

### UnsetCostBasisCents
`func (o *ListInventoryItem) UnsetCostBasisCents()`

UnsetCostBasisCents ensures that no value is present for CostBasisCents, not even an explicit nil
### GetPriceFloorCents

`func (o *ListInventoryItem) GetPriceFloorCents() float32`

GetPriceFloorCents returns the PriceFloorCents field if non-nil, zero value otherwise.

### GetPriceFloorCentsOk

`func (o *ListInventoryItem) GetPriceFloorCentsOk() (*float32, bool)`

GetPriceFloorCentsOk returns a tuple with the PriceFloorCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFloorCents

`func (o *ListInventoryItem) SetPriceFloorCents(v float32)`

SetPriceFloorCents sets PriceFloorCents field to given value.

### HasPriceFloorCents

`func (o *ListInventoryItem) HasPriceFloorCents() bool`

HasPriceFloorCents returns a boolean if a field has been set.

### SetPriceFloorCentsNil

`func (o *ListInventoryItem) SetPriceFloorCentsNil(b bool)`

 SetPriceFloorCentsNil sets the value for PriceFloorCents to be an explicit nil

### UnsetPriceFloorCents
`func (o *ListInventoryItem) UnsetPriceFloorCents()`

UnsetPriceFloorCents ensures that no value is present for PriceFloorCents, not even an explicit nil
### GetFloorIsNet

`func (o *ListInventoryItem) GetFloorIsNet() bool`

GetFloorIsNet returns the FloorIsNet field if non-nil, zero value otherwise.

### GetFloorIsNetOk

`func (o *ListInventoryItem) GetFloorIsNetOk() (*bool, bool)`

GetFloorIsNetOk returns a tuple with the FloorIsNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorIsNet

`func (o *ListInventoryItem) SetFloorIsNet(v bool)`

SetFloorIsNet sets FloorIsNet field to given value.


### GetAcquiredAt

`func (o *ListInventoryItem) GetAcquiredAt() time.Time`

GetAcquiredAt returns the AcquiredAt field if non-nil, zero value otherwise.

### GetAcquiredAtOk

`func (o *ListInventoryItem) GetAcquiredAtOk() (*time.Time, bool)`

GetAcquiredAtOk returns a tuple with the AcquiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquiredAt

`func (o *ListInventoryItem) SetAcquiredAt(v time.Time)`

SetAcquiredAt sets AcquiredAt field to given value.

### HasAcquiredAt

`func (o *ListInventoryItem) HasAcquiredAt() bool`

HasAcquiredAt returns a boolean if a field has been set.

### SetAcquiredAtNil

`func (o *ListInventoryItem) SetAcquiredAtNil(b bool)`

 SetAcquiredAtNil sets the value for AcquiredAt to be an explicit nil

### UnsetAcquiredAt
`func (o *ListInventoryItem) UnsetAcquiredAt()`

UnsetAcquiredAt ensures that no value is present for AcquiredAt, not even an explicit nil
### GetAcquiredSource

`func (o *ListInventoryItem) GetAcquiredSource() string`

GetAcquiredSource returns the AcquiredSource field if non-nil, zero value otherwise.

### GetAcquiredSourceOk

`func (o *ListInventoryItem) GetAcquiredSourceOk() (*string, bool)`

GetAcquiredSourceOk returns a tuple with the AcquiredSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAcquiredSource

`func (o *ListInventoryItem) SetAcquiredSource(v string)`

SetAcquiredSource sets AcquiredSource field to given value.

### HasAcquiredSource

`func (o *ListInventoryItem) HasAcquiredSource() bool`

HasAcquiredSource returns a boolean if a field has been set.

### SetAcquiredSourceNil

`func (o *ListInventoryItem) SetAcquiredSourceNil(b bool)`

 SetAcquiredSourceNil sets the value for AcquiredSource to be an explicit nil

### UnsetAcquiredSource
`func (o *ListInventoryItem) UnsetAcquiredSource()`

UnsetAcquiredSource ensures that no value is present for AcquiredSource, not even an explicit nil
### GetWeightLb

`func (o *ListInventoryItem) GetWeightLb() string`

GetWeightLb returns the WeightLb field if non-nil, zero value otherwise.

### GetWeightLbOk

`func (o *ListInventoryItem) GetWeightLbOk() (*string, bool)`

GetWeightLbOk returns a tuple with the WeightLb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightLb

`func (o *ListInventoryItem) SetWeightLb(v string)`

SetWeightLb sets WeightLb field to given value.

### HasWeightLb

`func (o *ListInventoryItem) HasWeightLb() bool`

HasWeightLb returns a boolean if a field has been set.

### SetWeightLbNil

`func (o *ListInventoryItem) SetWeightLbNil(b bool)`

 SetWeightLbNil sets the value for WeightLb to be an explicit nil

### UnsetWeightLb
`func (o *ListInventoryItem) UnsetWeightLb()`

UnsetWeightLb ensures that no value is present for WeightLb, not even an explicit nil
### GetWeightOz

`func (o *ListInventoryItem) GetWeightOz() string`

GetWeightOz returns the WeightOz field if non-nil, zero value otherwise.

### GetWeightOzOk

`func (o *ListInventoryItem) GetWeightOzOk() (*string, bool)`

GetWeightOzOk returns a tuple with the WeightOz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightOz

`func (o *ListInventoryItem) SetWeightOz(v string)`

SetWeightOz sets WeightOz field to given value.

### HasWeightOz

`func (o *ListInventoryItem) HasWeightOz() bool`

HasWeightOz returns a boolean if a field has been set.

### SetWeightOzNil

`func (o *ListInventoryItem) SetWeightOzNil(b bool)`

 SetWeightOzNil sets the value for WeightOz to be an explicit nil

### UnsetWeightOz
`func (o *ListInventoryItem) UnsetWeightOz()`

UnsetWeightOz ensures that no value is present for WeightOz, not even an explicit nil
### GetDimensionLIn

`func (o *ListInventoryItem) GetDimensionLIn() string`

GetDimensionLIn returns the DimensionLIn field if non-nil, zero value otherwise.

### GetDimensionLInOk

`func (o *ListInventoryItem) GetDimensionLInOk() (*string, bool)`

GetDimensionLInOk returns a tuple with the DimensionLIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionLIn

`func (o *ListInventoryItem) SetDimensionLIn(v string)`

SetDimensionLIn sets DimensionLIn field to given value.

### HasDimensionLIn

`func (o *ListInventoryItem) HasDimensionLIn() bool`

HasDimensionLIn returns a boolean if a field has been set.

### SetDimensionLInNil

`func (o *ListInventoryItem) SetDimensionLInNil(b bool)`

 SetDimensionLInNil sets the value for DimensionLIn to be an explicit nil

### UnsetDimensionLIn
`func (o *ListInventoryItem) UnsetDimensionLIn()`

UnsetDimensionLIn ensures that no value is present for DimensionLIn, not even an explicit nil
### GetDimensionWIn

`func (o *ListInventoryItem) GetDimensionWIn() string`

GetDimensionWIn returns the DimensionWIn field if non-nil, zero value otherwise.

### GetDimensionWInOk

`func (o *ListInventoryItem) GetDimensionWInOk() (*string, bool)`

GetDimensionWInOk returns a tuple with the DimensionWIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionWIn

`func (o *ListInventoryItem) SetDimensionWIn(v string)`

SetDimensionWIn sets DimensionWIn field to given value.

### HasDimensionWIn

`func (o *ListInventoryItem) HasDimensionWIn() bool`

HasDimensionWIn returns a boolean if a field has been set.

### SetDimensionWInNil

`func (o *ListInventoryItem) SetDimensionWInNil(b bool)`

 SetDimensionWInNil sets the value for DimensionWIn to be an explicit nil

### UnsetDimensionWIn
`func (o *ListInventoryItem) UnsetDimensionWIn()`

UnsetDimensionWIn ensures that no value is present for DimensionWIn, not even an explicit nil
### GetDimensionHIn

`func (o *ListInventoryItem) GetDimensionHIn() string`

GetDimensionHIn returns the DimensionHIn field if non-nil, zero value otherwise.

### GetDimensionHInOk

`func (o *ListInventoryItem) GetDimensionHInOk() (*string, bool)`

GetDimensionHInOk returns a tuple with the DimensionHIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionHIn

`func (o *ListInventoryItem) SetDimensionHIn(v string)`

SetDimensionHIn sets DimensionHIn field to given value.

### HasDimensionHIn

`func (o *ListInventoryItem) HasDimensionHIn() bool`

HasDimensionHIn returns a boolean if a field has been set.

### SetDimensionHInNil

`func (o *ListInventoryItem) SetDimensionHInNil(b bool)`

 SetDimensionHInNil sets the value for DimensionHIn to be an explicit nil

### UnsetDimensionHIn
`func (o *ListInventoryItem) UnsetDimensionHIn()`

UnsetDimensionHIn ensures that no value is present for DimensionHIn, not even an explicit nil
### GetStatus

`func (o *ListInventoryItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListInventoryItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListInventoryItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetSource

`func (o *ListInventoryItem) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *ListInventoryItem) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *ListInventoryItem) SetSource(v string)`

SetSource sets Source field to given value.


### GetCategory

`func (o *ListInventoryItem) GetCategory() ListInventoryItemCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListInventoryItem) GetCategoryOk() (*ListInventoryItemCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListInventoryItem) SetCategory(v ListInventoryItemCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListInventoryItem) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ListInventoryItem) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ListInventoryItem) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


