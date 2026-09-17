# CreateListingImportByUrlResponseListing

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PlatformListingId** | **string** |  | 
**Title** | **string** |  | 
**Description** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableFloat32** |  | [optional] 
**Images** | Pointer to **[]string** |  | [optional] 
**Brand** | Pointer to **NullableString** |  | [optional] 
**Condition** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableString** |  | [optional] 
**Color** | Pointer to **[]string** | Color(s) the platform&#39;s own response exposes (e.g. Poshmark&#39;s  &#x60;colors&#x60; array). Only set when actually present in the scrape. | [optional] 
**Sku** | Pointer to **NullableString** |  | [optional] 
**Url** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **[]string** | The platform&#39;s OWN category, as a top-down path (e.g. Poshmark&#39;s  [\&quot;Kids\&quot;, \&quot;Toys\&quot;, \&quot;Dolls &amp; Accessories\&quot;] from department → category →  feature). Only set when the platform&#39;s list/scrape response actually  carries this — never guessed. Mapped onto Crossly&#39;s category.main/  sub/sub2 (listings) or categoryMain/categorySub (inventory_items) at  create/enrich time. | [optional] 
**CategoryId** | Pointer to **NullableString** | The platform&#39;s raw numeric category id, same space as  overrides.&lt;platform&gt;.categoryId — distinct from &#x60;category&#x60;&#39;s  human-readable path. Only set by platforms whose id space is directly  comparable to what we publish (currently eBay&#39;s drift-check second  call — see diff-fields/second-call.ts&#39;s ebayRemoteFields). eBay&#39;s own  CategoryName wording/depth is a different vocabulary from Crossly&#39;s  master taxonomy breadcrumb and will essentially never string-match  it, so drift-detection compares ids instead of names. | [optional] 
**Tags** | Pointer to **[]string** | Tag-like strings the platform&#39;s own response exposes (e.g. Poshmark&#39;s  marketing \&quot;experience\&quot; tags). Distinct from a full search-tag  generator — just whatever real tag data the scrape already carries. | [optional] 
**Material** | Pointer to **NullableString** | Structured item aspects — the same shape the listing form&#39;s Item  Details section captures (and templates already persist), now sourced  from the platform&#39;s own data instead of only manual entry. Only set  when the platform&#39;s response genuinely carries a semantically-matching  field (eBay item specifics, Facebook attributes, Poshmark catalog,  Depop&#39;s detail-call ride-along) — never derived or guessed. | [optional] 
**Style** | Pointer to **NullableString** |  | [optional] 
**Pattern** | Pointer to **NullableString** |  | [optional] 
**Department** | Pointer to **NullableString** |  | [optional] 
**Gender** | Pointer to **NullableString** |  | [optional] 
**ItemType** | Pointer to **NullableString** |  | [optional] 
**SizeSystem** | Pointer to **NullableString** |  | [optional] 
**ItemSpecifics** | Pointer to **map[string]interface{}** | Category-specific facets from a per-item DETAIL call (e.g. Facebook&#39;s  Age Range/Character/Age Group), keyed the same way eBay item-specifics  are: aspect name -&gt; value array. Only set by platforms with a genuine  per-item attribute source — costs one extra call per listing, so  populated by a dedicated enrichment pass, not the main list mapper.  Maps onto listings.itemSpecifics; inventory_items has no equivalent  column. | [optional] 
**ListedAt** | Pointer to **NullableTime** | When the listing was first published on the platform. Used by  the Advanced filter&#39;s listedAfter / listedBefore knobs. Optional  because not every platform returns it on the listing endpoint. | [optional] 
**Quantity** | Pointer to **NullableFloat32** | Units this platform reports. Only meaningful for platforms  &#x60;PLATFORM_QUANTITY_SYNC&#x60; marks &#39;native&#39; — relist platforms show one item  and say 1 forever, so they leave this undefined rather than voting with  a number they cannot actually express. See &#x60;_quantity.ts&#x60;. | [optional] 
**AccountSlot** | Pointer to **NullableFloat32** | Which of the user&#39;s connected accounts on this platform this listing  was scraped from — stamped by fetchCookieListings as it loops each  connected account (see listActiveForPlatform in  user-platform-accounts/read.ts). Undefined for API-track platforms  (single connection, no multi-account concept) and for any cookie path  that hasn&#39;t been threaded yet; importOne treats undefined as slot 1,  matching the historical single-account default. | [optional] 
**WeightOz** | Pointer to **NullableFloat32** | Total item weight in ounces (already summed, not split lb+oz —  the write path converts to the DB&#39;s weightLb+weightOz split). | [optional] 
**DimensionLIn** | Pointer to **NullableFloat32** |  | [optional] 
**DimensionWIn** | Pointer to **NullableFloat32** |  | [optional] 
**DimensionHIn** | Pointer to **NullableFloat32** |  | [optional] 
**HandlingTimeDays** | Pointer to **NullableFloat32** | Max days the platform&#39;s own listing commits to ship within (eBay&#39;s  DispatchTimeMax, Etsy&#39;s processing_max). | [optional] 
**BestOfferAutoAcceptCents** | Pointer to **NullableFloat32** |  | [optional] 
**BestOfferAutoDeclineCents** | Pointer to **NullableFloat32** |  | [optional] 
**ItemLocation** | Pointer to **NullableString** | Free text as the platform itself expresses it (e.g. eBay&#39;s &#x60;Location&#x60;  is a single seller-typed string like \&quot;Austin, TX\&quot;, not a structured  address) — never parsed into city/state. | [optional] 
**ItemLocationZip** | Pointer to **NullableString** |  | [optional] 
**ItemLocationCountry** | Pointer to **NullableString** |  | [optional] 
**DeclaredShippingCost** | Pointer to **NullableFloat32** | What the platform&#39;s OWN listing declares shipping costs — reference  only, distinct from a realized post-sale shipping cost. | [optional] 
**ReturnPolicyText** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewCreateListingImportByUrlResponseListing

`func NewCreateListingImportByUrlResponseListing(platformListingId string, title string, ) *CreateListingImportByUrlResponseListing`

NewCreateListingImportByUrlResponseListing instantiates a new CreateListingImportByUrlResponseListing object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateListingImportByUrlResponseListingWithDefaults

`func NewCreateListingImportByUrlResponseListingWithDefaults() *CreateListingImportByUrlResponseListing`

NewCreateListingImportByUrlResponseListingWithDefaults instantiates a new CreateListingImportByUrlResponseListing object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlatformListingId

`func (o *CreateListingImportByUrlResponseListing) GetPlatformListingId() string`

GetPlatformListingId returns the PlatformListingId field if non-nil, zero value otherwise.

### GetPlatformListingIdOk

`func (o *CreateListingImportByUrlResponseListing) GetPlatformListingIdOk() (*string, bool)`

GetPlatformListingIdOk returns a tuple with the PlatformListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformListingId

`func (o *CreateListingImportByUrlResponseListing) SetPlatformListingId(v string)`

SetPlatformListingId sets PlatformListingId field to given value.


### GetTitle

`func (o *CreateListingImportByUrlResponseListing) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateListingImportByUrlResponseListing) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateListingImportByUrlResponseListing) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CreateListingImportByUrlResponseListing) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateListingImportByUrlResponseListing) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateListingImportByUrlResponseListing) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateListingImportByUrlResponseListing) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateListingImportByUrlResponseListing) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateListingImportByUrlResponseListing) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetPrice

`func (o *CreateListingImportByUrlResponseListing) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateListingImportByUrlResponseListing) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateListingImportByUrlResponseListing) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CreateListingImportByUrlResponseListing) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *CreateListingImportByUrlResponseListing) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *CreateListingImportByUrlResponseListing) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetImages

`func (o *CreateListingImportByUrlResponseListing) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CreateListingImportByUrlResponseListing) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CreateListingImportByUrlResponseListing) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *CreateListingImportByUrlResponseListing) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *CreateListingImportByUrlResponseListing) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *CreateListingImportByUrlResponseListing) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetBrand

`func (o *CreateListingImportByUrlResponseListing) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CreateListingImportByUrlResponseListing) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CreateListingImportByUrlResponseListing) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CreateListingImportByUrlResponseListing) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *CreateListingImportByUrlResponseListing) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *CreateListingImportByUrlResponseListing) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetCondition

`func (o *CreateListingImportByUrlResponseListing) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CreateListingImportByUrlResponseListing) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CreateListingImportByUrlResponseListing) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *CreateListingImportByUrlResponseListing) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *CreateListingImportByUrlResponseListing) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *CreateListingImportByUrlResponseListing) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetSize

`func (o *CreateListingImportByUrlResponseListing) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateListingImportByUrlResponseListing) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateListingImportByUrlResponseListing) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateListingImportByUrlResponseListing) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *CreateListingImportByUrlResponseListing) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *CreateListingImportByUrlResponseListing) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetColor

`func (o *CreateListingImportByUrlResponseListing) GetColor() []string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateListingImportByUrlResponseListing) GetColorOk() (*[]string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateListingImportByUrlResponseListing) SetColor(v []string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateListingImportByUrlResponseListing) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CreateListingImportByUrlResponseListing) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CreateListingImportByUrlResponseListing) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetSku

`func (o *CreateListingImportByUrlResponseListing) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CreateListingImportByUrlResponseListing) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CreateListingImportByUrlResponseListing) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CreateListingImportByUrlResponseListing) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *CreateListingImportByUrlResponseListing) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *CreateListingImportByUrlResponseListing) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetUrl

`func (o *CreateListingImportByUrlResponseListing) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *CreateListingImportByUrlResponseListing) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *CreateListingImportByUrlResponseListing) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *CreateListingImportByUrlResponseListing) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### SetUrlNil

`func (o *CreateListingImportByUrlResponseListing) SetUrlNil(b bool)`

 SetUrlNil sets the value for Url to be an explicit nil

### UnsetUrl
`func (o *CreateListingImportByUrlResponseListing) UnsetUrl()`

UnsetUrl ensures that no value is present for Url, not even an explicit nil
### GetCategory

`func (o *CreateListingImportByUrlResponseListing) GetCategory() []string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateListingImportByUrlResponseListing) GetCategoryOk() (*[]string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateListingImportByUrlResponseListing) SetCategory(v []string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateListingImportByUrlResponseListing) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *CreateListingImportByUrlResponseListing) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *CreateListingImportByUrlResponseListing) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCategoryId

`func (o *CreateListingImportByUrlResponseListing) GetCategoryId() string`

GetCategoryId returns the CategoryId field if non-nil, zero value otherwise.

### GetCategoryIdOk

`func (o *CreateListingImportByUrlResponseListing) GetCategoryIdOk() (*string, bool)`

GetCategoryIdOk returns a tuple with the CategoryId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategoryId

`func (o *CreateListingImportByUrlResponseListing) SetCategoryId(v string)`

SetCategoryId sets CategoryId field to given value.

### HasCategoryId

`func (o *CreateListingImportByUrlResponseListing) HasCategoryId() bool`

HasCategoryId returns a boolean if a field has been set.

### SetCategoryIdNil

`func (o *CreateListingImportByUrlResponseListing) SetCategoryIdNil(b bool)`

 SetCategoryIdNil sets the value for CategoryId to be an explicit nil

### UnsetCategoryId
`func (o *CreateListingImportByUrlResponseListing) UnsetCategoryId()`

UnsetCategoryId ensures that no value is present for CategoryId, not even an explicit nil
### GetTags

`func (o *CreateListingImportByUrlResponseListing) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateListingImportByUrlResponseListing) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateListingImportByUrlResponseListing) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateListingImportByUrlResponseListing) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateListingImportByUrlResponseListing) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateListingImportByUrlResponseListing) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetMaterial

`func (o *CreateListingImportByUrlResponseListing) GetMaterial() string`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *CreateListingImportByUrlResponseListing) GetMaterialOk() (*string, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *CreateListingImportByUrlResponseListing) SetMaterial(v string)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *CreateListingImportByUrlResponseListing) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### SetMaterialNil

`func (o *CreateListingImportByUrlResponseListing) SetMaterialNil(b bool)`

 SetMaterialNil sets the value for Material to be an explicit nil

### UnsetMaterial
`func (o *CreateListingImportByUrlResponseListing) UnsetMaterial()`

UnsetMaterial ensures that no value is present for Material, not even an explicit nil
### GetStyle

`func (o *CreateListingImportByUrlResponseListing) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *CreateListingImportByUrlResponseListing) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *CreateListingImportByUrlResponseListing) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *CreateListingImportByUrlResponseListing) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *CreateListingImportByUrlResponseListing) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *CreateListingImportByUrlResponseListing) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetPattern

`func (o *CreateListingImportByUrlResponseListing) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *CreateListingImportByUrlResponseListing) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *CreateListingImportByUrlResponseListing) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *CreateListingImportByUrlResponseListing) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPatternNil

`func (o *CreateListingImportByUrlResponseListing) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *CreateListingImportByUrlResponseListing) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil
### GetDepartment

`func (o *CreateListingImportByUrlResponseListing) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *CreateListingImportByUrlResponseListing) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *CreateListingImportByUrlResponseListing) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *CreateListingImportByUrlResponseListing) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *CreateListingImportByUrlResponseListing) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *CreateListingImportByUrlResponseListing) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetGender

`func (o *CreateListingImportByUrlResponseListing) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CreateListingImportByUrlResponseListing) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CreateListingImportByUrlResponseListing) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CreateListingImportByUrlResponseListing) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *CreateListingImportByUrlResponseListing) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *CreateListingImportByUrlResponseListing) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetItemType

`func (o *CreateListingImportByUrlResponseListing) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *CreateListingImportByUrlResponseListing) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *CreateListingImportByUrlResponseListing) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *CreateListingImportByUrlResponseListing) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *CreateListingImportByUrlResponseListing) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *CreateListingImportByUrlResponseListing) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetSizeSystem

`func (o *CreateListingImportByUrlResponseListing) GetSizeSystem() string`

GetSizeSystem returns the SizeSystem field if non-nil, zero value otherwise.

### GetSizeSystemOk

`func (o *CreateListingImportByUrlResponseListing) GetSizeSystemOk() (*string, bool)`

GetSizeSystemOk returns a tuple with the SizeSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSystem

`func (o *CreateListingImportByUrlResponseListing) SetSizeSystem(v string)`

SetSizeSystem sets SizeSystem field to given value.

### HasSizeSystem

`func (o *CreateListingImportByUrlResponseListing) HasSizeSystem() bool`

HasSizeSystem returns a boolean if a field has been set.

### SetSizeSystemNil

`func (o *CreateListingImportByUrlResponseListing) SetSizeSystemNil(b bool)`

 SetSizeSystemNil sets the value for SizeSystem to be an explicit nil

### UnsetSizeSystem
`func (o *CreateListingImportByUrlResponseListing) UnsetSizeSystem()`

UnsetSizeSystem ensures that no value is present for SizeSystem, not even an explicit nil
### GetItemSpecifics

`func (o *CreateListingImportByUrlResponseListing) GetItemSpecifics() map[string]interface{}`

GetItemSpecifics returns the ItemSpecifics field if non-nil, zero value otherwise.

### GetItemSpecificsOk

`func (o *CreateListingImportByUrlResponseListing) GetItemSpecificsOk() (*map[string]interface{}, bool)`

GetItemSpecificsOk returns a tuple with the ItemSpecifics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSpecifics

`func (o *CreateListingImportByUrlResponseListing) SetItemSpecifics(v map[string]interface{})`

SetItemSpecifics sets ItemSpecifics field to given value.

### HasItemSpecifics

`func (o *CreateListingImportByUrlResponseListing) HasItemSpecifics() bool`

HasItemSpecifics returns a boolean if a field has been set.

### SetItemSpecificsNil

`func (o *CreateListingImportByUrlResponseListing) SetItemSpecificsNil(b bool)`

 SetItemSpecificsNil sets the value for ItemSpecifics to be an explicit nil

### UnsetItemSpecifics
`func (o *CreateListingImportByUrlResponseListing) UnsetItemSpecifics()`

UnsetItemSpecifics ensures that no value is present for ItemSpecifics, not even an explicit nil
### GetListedAt

`func (o *CreateListingImportByUrlResponseListing) GetListedAt() time.Time`

GetListedAt returns the ListedAt field if non-nil, zero value otherwise.

### GetListedAtOk

`func (o *CreateListingImportByUrlResponseListing) GetListedAtOk() (*time.Time, bool)`

GetListedAtOk returns a tuple with the ListedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListedAt

`func (o *CreateListingImportByUrlResponseListing) SetListedAt(v time.Time)`

SetListedAt sets ListedAt field to given value.

### HasListedAt

`func (o *CreateListingImportByUrlResponseListing) HasListedAt() bool`

HasListedAt returns a boolean if a field has been set.

### SetListedAtNil

`func (o *CreateListingImportByUrlResponseListing) SetListedAtNil(b bool)`

 SetListedAtNil sets the value for ListedAt to be an explicit nil

### UnsetListedAt
`func (o *CreateListingImportByUrlResponseListing) UnsetListedAt()`

UnsetListedAt ensures that no value is present for ListedAt, not even an explicit nil
### GetQuantity

`func (o *CreateListingImportByUrlResponseListing) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateListingImportByUrlResponseListing) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateListingImportByUrlResponseListing) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CreateListingImportByUrlResponseListing) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *CreateListingImportByUrlResponseListing) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *CreateListingImportByUrlResponseListing) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetAccountSlot

`func (o *CreateListingImportByUrlResponseListing) GetAccountSlot() float32`

GetAccountSlot returns the AccountSlot field if non-nil, zero value otherwise.

### GetAccountSlotOk

`func (o *CreateListingImportByUrlResponseListing) GetAccountSlotOk() (*float32, bool)`

GetAccountSlotOk returns a tuple with the AccountSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountSlot

`func (o *CreateListingImportByUrlResponseListing) SetAccountSlot(v float32)`

SetAccountSlot sets AccountSlot field to given value.

### HasAccountSlot

`func (o *CreateListingImportByUrlResponseListing) HasAccountSlot() bool`

HasAccountSlot returns a boolean if a field has been set.

### SetAccountSlotNil

`func (o *CreateListingImportByUrlResponseListing) SetAccountSlotNil(b bool)`

 SetAccountSlotNil sets the value for AccountSlot to be an explicit nil

### UnsetAccountSlot
`func (o *CreateListingImportByUrlResponseListing) UnsetAccountSlot()`

UnsetAccountSlot ensures that no value is present for AccountSlot, not even an explicit nil
### GetWeightOz

`func (o *CreateListingImportByUrlResponseListing) GetWeightOz() float32`

GetWeightOz returns the WeightOz field if non-nil, zero value otherwise.

### GetWeightOzOk

`func (o *CreateListingImportByUrlResponseListing) GetWeightOzOk() (*float32, bool)`

GetWeightOzOk returns a tuple with the WeightOz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightOz

`func (o *CreateListingImportByUrlResponseListing) SetWeightOz(v float32)`

SetWeightOz sets WeightOz field to given value.

### HasWeightOz

`func (o *CreateListingImportByUrlResponseListing) HasWeightOz() bool`

HasWeightOz returns a boolean if a field has been set.

### SetWeightOzNil

`func (o *CreateListingImportByUrlResponseListing) SetWeightOzNil(b bool)`

 SetWeightOzNil sets the value for WeightOz to be an explicit nil

### UnsetWeightOz
`func (o *CreateListingImportByUrlResponseListing) UnsetWeightOz()`

UnsetWeightOz ensures that no value is present for WeightOz, not even an explicit nil
### GetDimensionLIn

`func (o *CreateListingImportByUrlResponseListing) GetDimensionLIn() float32`

GetDimensionLIn returns the DimensionLIn field if non-nil, zero value otherwise.

### GetDimensionLInOk

`func (o *CreateListingImportByUrlResponseListing) GetDimensionLInOk() (*float32, bool)`

GetDimensionLInOk returns a tuple with the DimensionLIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionLIn

`func (o *CreateListingImportByUrlResponseListing) SetDimensionLIn(v float32)`

SetDimensionLIn sets DimensionLIn field to given value.

### HasDimensionLIn

`func (o *CreateListingImportByUrlResponseListing) HasDimensionLIn() bool`

HasDimensionLIn returns a boolean if a field has been set.

### SetDimensionLInNil

`func (o *CreateListingImportByUrlResponseListing) SetDimensionLInNil(b bool)`

 SetDimensionLInNil sets the value for DimensionLIn to be an explicit nil

### UnsetDimensionLIn
`func (o *CreateListingImportByUrlResponseListing) UnsetDimensionLIn()`

UnsetDimensionLIn ensures that no value is present for DimensionLIn, not even an explicit nil
### GetDimensionWIn

`func (o *CreateListingImportByUrlResponseListing) GetDimensionWIn() float32`

GetDimensionWIn returns the DimensionWIn field if non-nil, zero value otherwise.

### GetDimensionWInOk

`func (o *CreateListingImportByUrlResponseListing) GetDimensionWInOk() (*float32, bool)`

GetDimensionWInOk returns a tuple with the DimensionWIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionWIn

`func (o *CreateListingImportByUrlResponseListing) SetDimensionWIn(v float32)`

SetDimensionWIn sets DimensionWIn field to given value.

### HasDimensionWIn

`func (o *CreateListingImportByUrlResponseListing) HasDimensionWIn() bool`

HasDimensionWIn returns a boolean if a field has been set.

### SetDimensionWInNil

`func (o *CreateListingImportByUrlResponseListing) SetDimensionWInNil(b bool)`

 SetDimensionWInNil sets the value for DimensionWIn to be an explicit nil

### UnsetDimensionWIn
`func (o *CreateListingImportByUrlResponseListing) UnsetDimensionWIn()`

UnsetDimensionWIn ensures that no value is present for DimensionWIn, not even an explicit nil
### GetDimensionHIn

`func (o *CreateListingImportByUrlResponseListing) GetDimensionHIn() float32`

GetDimensionHIn returns the DimensionHIn field if non-nil, zero value otherwise.

### GetDimensionHInOk

`func (o *CreateListingImportByUrlResponseListing) GetDimensionHInOk() (*float32, bool)`

GetDimensionHInOk returns a tuple with the DimensionHIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionHIn

`func (o *CreateListingImportByUrlResponseListing) SetDimensionHIn(v float32)`

SetDimensionHIn sets DimensionHIn field to given value.

### HasDimensionHIn

`func (o *CreateListingImportByUrlResponseListing) HasDimensionHIn() bool`

HasDimensionHIn returns a boolean if a field has been set.

### SetDimensionHInNil

`func (o *CreateListingImportByUrlResponseListing) SetDimensionHInNil(b bool)`

 SetDimensionHInNil sets the value for DimensionHIn to be an explicit nil

### UnsetDimensionHIn
`func (o *CreateListingImportByUrlResponseListing) UnsetDimensionHIn()`

UnsetDimensionHIn ensures that no value is present for DimensionHIn, not even an explicit nil
### GetHandlingTimeDays

`func (o *CreateListingImportByUrlResponseListing) GetHandlingTimeDays() float32`

GetHandlingTimeDays returns the HandlingTimeDays field if non-nil, zero value otherwise.

### GetHandlingTimeDaysOk

`func (o *CreateListingImportByUrlResponseListing) GetHandlingTimeDaysOk() (*float32, bool)`

GetHandlingTimeDaysOk returns a tuple with the HandlingTimeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingTimeDays

`func (o *CreateListingImportByUrlResponseListing) SetHandlingTimeDays(v float32)`

SetHandlingTimeDays sets HandlingTimeDays field to given value.

### HasHandlingTimeDays

`func (o *CreateListingImportByUrlResponseListing) HasHandlingTimeDays() bool`

HasHandlingTimeDays returns a boolean if a field has been set.

### SetHandlingTimeDaysNil

`func (o *CreateListingImportByUrlResponseListing) SetHandlingTimeDaysNil(b bool)`

 SetHandlingTimeDaysNil sets the value for HandlingTimeDays to be an explicit nil

### UnsetHandlingTimeDays
`func (o *CreateListingImportByUrlResponseListing) UnsetHandlingTimeDays()`

UnsetHandlingTimeDays ensures that no value is present for HandlingTimeDays, not even an explicit nil
### GetBestOfferAutoAcceptCents

`func (o *CreateListingImportByUrlResponseListing) GetBestOfferAutoAcceptCents() float32`

GetBestOfferAutoAcceptCents returns the BestOfferAutoAcceptCents field if non-nil, zero value otherwise.

### GetBestOfferAutoAcceptCentsOk

`func (o *CreateListingImportByUrlResponseListing) GetBestOfferAutoAcceptCentsOk() (*float32, bool)`

GetBestOfferAutoAcceptCentsOk returns a tuple with the BestOfferAutoAcceptCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestOfferAutoAcceptCents

`func (o *CreateListingImportByUrlResponseListing) SetBestOfferAutoAcceptCents(v float32)`

SetBestOfferAutoAcceptCents sets BestOfferAutoAcceptCents field to given value.

### HasBestOfferAutoAcceptCents

`func (o *CreateListingImportByUrlResponseListing) HasBestOfferAutoAcceptCents() bool`

HasBestOfferAutoAcceptCents returns a boolean if a field has been set.

### SetBestOfferAutoAcceptCentsNil

`func (o *CreateListingImportByUrlResponseListing) SetBestOfferAutoAcceptCentsNil(b bool)`

 SetBestOfferAutoAcceptCentsNil sets the value for BestOfferAutoAcceptCents to be an explicit nil

### UnsetBestOfferAutoAcceptCents
`func (o *CreateListingImportByUrlResponseListing) UnsetBestOfferAutoAcceptCents()`

UnsetBestOfferAutoAcceptCents ensures that no value is present for BestOfferAutoAcceptCents, not even an explicit nil
### GetBestOfferAutoDeclineCents

`func (o *CreateListingImportByUrlResponseListing) GetBestOfferAutoDeclineCents() float32`

GetBestOfferAutoDeclineCents returns the BestOfferAutoDeclineCents field if non-nil, zero value otherwise.

### GetBestOfferAutoDeclineCentsOk

`func (o *CreateListingImportByUrlResponseListing) GetBestOfferAutoDeclineCentsOk() (*float32, bool)`

GetBestOfferAutoDeclineCentsOk returns a tuple with the BestOfferAutoDeclineCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBestOfferAutoDeclineCents

`func (o *CreateListingImportByUrlResponseListing) SetBestOfferAutoDeclineCents(v float32)`

SetBestOfferAutoDeclineCents sets BestOfferAutoDeclineCents field to given value.

### HasBestOfferAutoDeclineCents

`func (o *CreateListingImportByUrlResponseListing) HasBestOfferAutoDeclineCents() bool`

HasBestOfferAutoDeclineCents returns a boolean if a field has been set.

### SetBestOfferAutoDeclineCentsNil

`func (o *CreateListingImportByUrlResponseListing) SetBestOfferAutoDeclineCentsNil(b bool)`

 SetBestOfferAutoDeclineCentsNil sets the value for BestOfferAutoDeclineCents to be an explicit nil

### UnsetBestOfferAutoDeclineCents
`func (o *CreateListingImportByUrlResponseListing) UnsetBestOfferAutoDeclineCents()`

UnsetBestOfferAutoDeclineCents ensures that no value is present for BestOfferAutoDeclineCents, not even an explicit nil
### GetItemLocation

`func (o *CreateListingImportByUrlResponseListing) GetItemLocation() string`

GetItemLocation returns the ItemLocation field if non-nil, zero value otherwise.

### GetItemLocationOk

`func (o *CreateListingImportByUrlResponseListing) GetItemLocationOk() (*string, bool)`

GetItemLocationOk returns a tuple with the ItemLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLocation

`func (o *CreateListingImportByUrlResponseListing) SetItemLocation(v string)`

SetItemLocation sets ItemLocation field to given value.

### HasItemLocation

`func (o *CreateListingImportByUrlResponseListing) HasItemLocation() bool`

HasItemLocation returns a boolean if a field has been set.

### SetItemLocationNil

`func (o *CreateListingImportByUrlResponseListing) SetItemLocationNil(b bool)`

 SetItemLocationNil sets the value for ItemLocation to be an explicit nil

### UnsetItemLocation
`func (o *CreateListingImportByUrlResponseListing) UnsetItemLocation()`

UnsetItemLocation ensures that no value is present for ItemLocation, not even an explicit nil
### GetItemLocationZip

`func (o *CreateListingImportByUrlResponseListing) GetItemLocationZip() string`

GetItemLocationZip returns the ItemLocationZip field if non-nil, zero value otherwise.

### GetItemLocationZipOk

`func (o *CreateListingImportByUrlResponseListing) GetItemLocationZipOk() (*string, bool)`

GetItemLocationZipOk returns a tuple with the ItemLocationZip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLocationZip

`func (o *CreateListingImportByUrlResponseListing) SetItemLocationZip(v string)`

SetItemLocationZip sets ItemLocationZip field to given value.

### HasItemLocationZip

`func (o *CreateListingImportByUrlResponseListing) HasItemLocationZip() bool`

HasItemLocationZip returns a boolean if a field has been set.

### SetItemLocationZipNil

`func (o *CreateListingImportByUrlResponseListing) SetItemLocationZipNil(b bool)`

 SetItemLocationZipNil sets the value for ItemLocationZip to be an explicit nil

### UnsetItemLocationZip
`func (o *CreateListingImportByUrlResponseListing) UnsetItemLocationZip()`

UnsetItemLocationZip ensures that no value is present for ItemLocationZip, not even an explicit nil
### GetItemLocationCountry

`func (o *CreateListingImportByUrlResponseListing) GetItemLocationCountry() string`

GetItemLocationCountry returns the ItemLocationCountry field if non-nil, zero value otherwise.

### GetItemLocationCountryOk

`func (o *CreateListingImportByUrlResponseListing) GetItemLocationCountryOk() (*string, bool)`

GetItemLocationCountryOk returns a tuple with the ItemLocationCountry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemLocationCountry

`func (o *CreateListingImportByUrlResponseListing) SetItemLocationCountry(v string)`

SetItemLocationCountry sets ItemLocationCountry field to given value.

### HasItemLocationCountry

`func (o *CreateListingImportByUrlResponseListing) HasItemLocationCountry() bool`

HasItemLocationCountry returns a boolean if a field has been set.

### SetItemLocationCountryNil

`func (o *CreateListingImportByUrlResponseListing) SetItemLocationCountryNil(b bool)`

 SetItemLocationCountryNil sets the value for ItemLocationCountry to be an explicit nil

### UnsetItemLocationCountry
`func (o *CreateListingImportByUrlResponseListing) UnsetItemLocationCountry()`

UnsetItemLocationCountry ensures that no value is present for ItemLocationCountry, not even an explicit nil
### GetDeclaredShippingCost

`func (o *CreateListingImportByUrlResponseListing) GetDeclaredShippingCost() float32`

GetDeclaredShippingCost returns the DeclaredShippingCost field if non-nil, zero value otherwise.

### GetDeclaredShippingCostOk

`func (o *CreateListingImportByUrlResponseListing) GetDeclaredShippingCostOk() (*float32, bool)`

GetDeclaredShippingCostOk returns a tuple with the DeclaredShippingCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeclaredShippingCost

`func (o *CreateListingImportByUrlResponseListing) SetDeclaredShippingCost(v float32)`

SetDeclaredShippingCost sets DeclaredShippingCost field to given value.

### HasDeclaredShippingCost

`func (o *CreateListingImportByUrlResponseListing) HasDeclaredShippingCost() bool`

HasDeclaredShippingCost returns a boolean if a field has been set.

### SetDeclaredShippingCostNil

`func (o *CreateListingImportByUrlResponseListing) SetDeclaredShippingCostNil(b bool)`

 SetDeclaredShippingCostNil sets the value for DeclaredShippingCost to be an explicit nil

### UnsetDeclaredShippingCost
`func (o *CreateListingImportByUrlResponseListing) UnsetDeclaredShippingCost()`

UnsetDeclaredShippingCost ensures that no value is present for DeclaredShippingCost, not even an explicit nil
### GetReturnPolicyText

`func (o *CreateListingImportByUrlResponseListing) GetReturnPolicyText() string`

GetReturnPolicyText returns the ReturnPolicyText field if non-nil, zero value otherwise.

### GetReturnPolicyTextOk

`func (o *CreateListingImportByUrlResponseListing) GetReturnPolicyTextOk() (*string, bool)`

GetReturnPolicyTextOk returns a tuple with the ReturnPolicyText field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnPolicyText

`func (o *CreateListingImportByUrlResponseListing) SetReturnPolicyText(v string)`

SetReturnPolicyText sets ReturnPolicyText field to given value.

### HasReturnPolicyText

`func (o *CreateListingImportByUrlResponseListing) HasReturnPolicyText() bool`

HasReturnPolicyText returns a boolean if a field has been set.

### SetReturnPolicyTextNil

`func (o *CreateListingImportByUrlResponseListing) SetReturnPolicyTextNil(b bool)`

 SetReturnPolicyTextNil sets the value for ReturnPolicyText to be an explicit nil

### UnsetReturnPolicyText
`func (o *CreateListingImportByUrlResponseListing) UnsetReturnPolicyText()`

UnsetReturnPolicyText ensures that no value is present for ReturnPolicyText, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


