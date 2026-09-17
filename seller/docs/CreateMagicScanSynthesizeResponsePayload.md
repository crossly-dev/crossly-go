# CreateMagicScanSynthesizeResponsePayload

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Description** | **string** |  | 
**Brand** | Pointer to **NullableString** |  | [optional] 
**Grading** | Pointer to [**NullableCreateMagicScanSynthesizeResponsePayloadGrading**](CreateMagicScanSynthesizeResponsePayloadGrading.md) |  | [optional] 
**Color** | Pointer to **NullableString** |  | [optional] 
**Material** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableString** |  | [optional] 
**SizeSystem** | Pointer to **NullableString** | Poshmark/Vestiaire size system — \&quot;US\&quot;, \&quot;EU\&quot;, \&quot;UK\&quot;, etc. | [optional] 
**WeightOz** | Pointer to **NullableFloat32** |  | [optional] 
**Dimensions** | Pointer to [**NullableCreateMagicScanSynthesizeResponsePayloadDimensions**](CreateMagicScanSynthesizeResponsePayloadDimensions.md) |  | [optional] 
**PriceCents** | Pointer to **NullableFloat32** | Echoed user-selected price (cents). Set on every synthesize call  so the master form&#39;s defaultPrice input gets populated regardless  of which path (AI or heuristic) generated the rest. | [optional] 
**SellerNote** | Pointer to **NullableString** | The seller&#39;s own words about this item — the pre-scan hint, or the  post-scan note that supersedes it (see synthesize/seller-note.ts).  Carried on the payload so the TAXONOMY resolvers see it too: category  and facet picking happen in their own LLM calls, which never saw the  note even after the main synthesis prompt did. A seller correcting the  variant was still getting the category of the wrong one. | [optional] 
**PriceLowCents** | Pointer to **NullableFloat32** | Comp price band (cents), outlier-trimmed — low/median/high across the  matched comps. The UI surfaces the range so the seller prices  strategically instead of trusting one number; priceCents defaults to  the median when the seller hasn&#39;t picked. | [optional] 
**PriceMedianCents** | Pointer to **NullableFloat32** |  | [optional] 
**PriceHighCents** | Pointer to **NullableFloat32** |  | [optional] 
**Quantity** | Pointer to **NullableFloat32** | Echoed seller-supplied listable-unit count from Magic List. This is  the listing&#39;s ADVERTISED stock (how many of this listing to sell).  Hydrates the form&#39;s &#x60;quantity&#x60; field. Undefined &#x3D; leave form default (1). | [optional] 
**InventoryQuantity** | Pointer to **NullableFloat32** | Physical units the seller actually has in stock — becomes the auto-  created inventory item&#39;s quantity/quantityAvailable. Distinct from  &#x60;quantity&#x60; (advertised stock) and &#x60;unitsPerListing&#x60; (composition).  Undefined &#x3D; fall back to &#x60;quantity&#x60;. | [optional] 
**UnitsPerListing** | Pointer to **NullableFloat32** | How many physical units are bundled inside ONE listing (composition;  e.g. a pack of 4). Becomes listing_inventory_items.quantity so the  multi-channel fulfillable &#x3D; floor(inventoryQuantity / unitsPerListing).  Undefined &#x3D; 1. | [optional] 
**CostBasisCents** | Pointer to **NullableFloat32** | What the seller paid, per physical unit, in cents — the auto-created  inventory item&#39;s &#x60;costBasisCents&#x60;. Nothing wrote that column from any  create path, so every Magic List item had a NULL cost basis and every  sale of one reported its full sale price as profit. Undefined &#x3D; unknown,  which stays NULL (0 would be a claim that the item was free). | [optional] 
**ChosenDimensions** | Pointer to [**NullableCreateMagicScanSynthesizeResponsePayloadChosenDimensions**](CreateMagicScanSynthesizeResponsePayloadChosenDimensions.md) |  | [optional] 
**ImageUrls** | Pointer to **[]string** | Every photo the seller uploaded for the source scan, primary  first. The NewListingPage hydrator drops these straight into the  form&#39;s images state so the seller doesn&#39;t have to re-upload. | [optional] 
**Category** | Pointer to [**NullableCreateMagicScanSynthesizeResponsePayloadCategory**](CreateMagicScanSynthesizeResponsePayloadCategory.md) |  | [optional] 
**Department** | Pointer to **NullableString** | eBay item-specifics: Department/Gender/Style/Pattern/Type. These  drive the eBay-required aspects on the form&#39;s Item Details  section. Mostly only eBay matches have them. | [optional] 
**Gender** | Pointer to **NullableString** |  | [optional] 
**Style** | Pointer to **NullableString** |  | [optional] 
**Pattern** | Pointer to **NullableString** |  | [optional] 
**ItemType** | Pointer to **NullableString** |  | [optional] 
**Upc** | Pointer to **NullableString** | Universal product code, when eBay&#39;s enriched detail surfaced it.  Hydrates the master form&#39;s UPC/GTIN field. | [optional] 
**HandlingTimeDays** | Pointer to **NullableString** |  | [optional] 
**ReturnWindowDays** | Pointer to **NullableString** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ItemSpecifics** | Pointer to **NullableString** | Master-level item specifics (JSON Record&lt;string,string[]&gt;). Item  specifics are now master-owned + fanned to every aspect platform, so the  form&#39;s shared Item-specifics editor prefills from this. Mirrored from the  richest resolved per-platform specifics blob (eBay after full-aspect  enrichment). Undefined &#x3D; no specifics. | [optional] 
**PerPlatformOverrides** | **map[string]interface{}** | Per-platform required-fields map — slotted directly into  platformOverrides when the form hydrates. | 
**Confidence** | **float32** | AI&#39;s confidence in the synthesis (0-1). Surfaced to the seller  so they can decide whether to skim or trust + click. | 
**Notes** | Pointer to **NullableString** | AI&#39;s free-text rationale for the merge — what it pulled from  where. Helps the seller spot a bad merge before listing. | [optional] 
**SectionApplicability** | Pointer to [**NullableCreateMagicScanSynthesizeResponsePayloadSectionApplicability**](CreateMagicScanSynthesizeResponsePayloadSectionApplicability.md) |  | [optional] 

## Methods

### NewCreateMagicScanSynthesizeResponsePayload

`func NewCreateMagicScanSynthesizeResponsePayload(title string, description string, perPlatformOverrides map[string]interface{}, confidence float32, ) *CreateMagicScanSynthesizeResponsePayload`

NewCreateMagicScanSynthesizeResponsePayload instantiates a new CreateMagicScanSynthesizeResponsePayload object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMagicScanSynthesizeResponsePayloadWithDefaults

`func NewCreateMagicScanSynthesizeResponsePayloadWithDefaults() *CreateMagicScanSynthesizeResponsePayload`

NewCreateMagicScanSynthesizeResponsePayloadWithDefaults instantiates a new CreateMagicScanSynthesizeResponsePayload object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *CreateMagicScanSynthesizeResponsePayload) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateMagicScanSynthesizeResponsePayload) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDescription

`func (o *CreateMagicScanSynthesizeResponsePayload) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateMagicScanSynthesizeResponsePayload) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetBrand

`func (o *CreateMagicScanSynthesizeResponsePayload) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CreateMagicScanSynthesizeResponsePayload) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CreateMagicScanSynthesizeResponsePayload) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetGrading

`func (o *CreateMagicScanSynthesizeResponsePayload) GetGrading() CreateMagicScanSynthesizeResponsePayloadGrading`

GetGrading returns the Grading field if non-nil, zero value otherwise.

### GetGradingOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetGradingOk() (*CreateMagicScanSynthesizeResponsePayloadGrading, bool)`

GetGradingOk returns a tuple with the Grading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrading

`func (o *CreateMagicScanSynthesizeResponsePayload) SetGrading(v CreateMagicScanSynthesizeResponsePayloadGrading)`

SetGrading sets Grading field to given value.

### HasGrading

`func (o *CreateMagicScanSynthesizeResponsePayload) HasGrading() bool`

HasGrading returns a boolean if a field has been set.

### SetGradingNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetGradingNil(b bool)`

 SetGradingNil sets the value for Grading to be an explicit nil

### UnsetGrading
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetGrading()`

UnsetGrading ensures that no value is present for Grading, not even an explicit nil
### GetColor

`func (o *CreateMagicScanSynthesizeResponsePayload) GetColor() string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetColorOk() (*string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateMagicScanSynthesizeResponsePayload) SetColor(v string)`

SetColor sets Color field to given value.

### HasColor

`func (o *CreateMagicScanSynthesizeResponsePayload) HasColor() bool`

HasColor returns a boolean if a field has been set.

### SetColorNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetColorNil(b bool)`

 SetColorNil sets the value for Color to be an explicit nil

### UnsetColor
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetColor()`

UnsetColor ensures that no value is present for Color, not even an explicit nil
### GetMaterial

`func (o *CreateMagicScanSynthesizeResponsePayload) GetMaterial() string`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetMaterialOk() (*string, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *CreateMagicScanSynthesizeResponsePayload) SetMaterial(v string)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *CreateMagicScanSynthesizeResponsePayload) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### SetMaterialNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetMaterialNil(b bool)`

 SetMaterialNil sets the value for Material to be an explicit nil

### UnsetMaterial
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetMaterial()`

UnsetMaterial ensures that no value is present for Material, not even an explicit nil
### GetSize

`func (o *CreateMagicScanSynthesizeResponsePayload) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateMagicScanSynthesizeResponsePayload) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateMagicScanSynthesizeResponsePayload) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetSizeSystem

`func (o *CreateMagicScanSynthesizeResponsePayload) GetSizeSystem() string`

GetSizeSystem returns the SizeSystem field if non-nil, zero value otherwise.

### GetSizeSystemOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetSizeSystemOk() (*string, bool)`

GetSizeSystemOk returns a tuple with the SizeSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSystem

`func (o *CreateMagicScanSynthesizeResponsePayload) SetSizeSystem(v string)`

SetSizeSystem sets SizeSystem field to given value.

### HasSizeSystem

`func (o *CreateMagicScanSynthesizeResponsePayload) HasSizeSystem() bool`

HasSizeSystem returns a boolean if a field has been set.

### SetSizeSystemNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetSizeSystemNil(b bool)`

 SetSizeSystemNil sets the value for SizeSystem to be an explicit nil

### UnsetSizeSystem
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetSizeSystem()`

UnsetSizeSystem ensures that no value is present for SizeSystem, not even an explicit nil
### GetWeightOz

`func (o *CreateMagicScanSynthesizeResponsePayload) GetWeightOz() float32`

GetWeightOz returns the WeightOz field if non-nil, zero value otherwise.

### GetWeightOzOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetWeightOzOk() (*float32, bool)`

GetWeightOzOk returns a tuple with the WeightOz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightOz

`func (o *CreateMagicScanSynthesizeResponsePayload) SetWeightOz(v float32)`

SetWeightOz sets WeightOz field to given value.

### HasWeightOz

`func (o *CreateMagicScanSynthesizeResponsePayload) HasWeightOz() bool`

HasWeightOz returns a boolean if a field has been set.

### SetWeightOzNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetWeightOzNil(b bool)`

 SetWeightOzNil sets the value for WeightOz to be an explicit nil

### UnsetWeightOz
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetWeightOz()`

UnsetWeightOz ensures that no value is present for WeightOz, not even an explicit nil
### GetDimensions

`func (o *CreateMagicScanSynthesizeResponsePayload) GetDimensions() CreateMagicScanSynthesizeResponsePayloadDimensions`

GetDimensions returns the Dimensions field if non-nil, zero value otherwise.

### GetDimensionsOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetDimensionsOk() (*CreateMagicScanSynthesizeResponsePayloadDimensions, bool)`

GetDimensionsOk returns a tuple with the Dimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensions

`func (o *CreateMagicScanSynthesizeResponsePayload) SetDimensions(v CreateMagicScanSynthesizeResponsePayloadDimensions)`

SetDimensions sets Dimensions field to given value.

### HasDimensions

`func (o *CreateMagicScanSynthesizeResponsePayload) HasDimensions() bool`

HasDimensions returns a boolean if a field has been set.

### SetDimensionsNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetDimensionsNil(b bool)`

 SetDimensionsNil sets the value for Dimensions to be an explicit nil

### UnsetDimensions
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetDimensions()`

UnsetDimensions ensures that no value is present for Dimensions, not even an explicit nil
### GetPriceCents

`func (o *CreateMagicScanSynthesizeResponsePayload) GetPriceCents() float32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetPriceCentsOk() (*float32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *CreateMagicScanSynthesizeResponsePayload) SetPriceCents(v float32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *CreateMagicScanSynthesizeResponsePayload) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### SetPriceCentsNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetPriceCentsNil(b bool)`

 SetPriceCentsNil sets the value for PriceCents to be an explicit nil

### UnsetPriceCents
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetPriceCents()`

UnsetPriceCents ensures that no value is present for PriceCents, not even an explicit nil
### GetSellerNote

`func (o *CreateMagicScanSynthesizeResponsePayload) GetSellerNote() string`

GetSellerNote returns the SellerNote field if non-nil, zero value otherwise.

### GetSellerNoteOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetSellerNoteOk() (*string, bool)`

GetSellerNoteOk returns a tuple with the SellerNote field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSellerNote

`func (o *CreateMagicScanSynthesizeResponsePayload) SetSellerNote(v string)`

SetSellerNote sets SellerNote field to given value.

### HasSellerNote

`func (o *CreateMagicScanSynthesizeResponsePayload) HasSellerNote() bool`

HasSellerNote returns a boolean if a field has been set.

### SetSellerNoteNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetSellerNoteNil(b bool)`

 SetSellerNoteNil sets the value for SellerNote to be an explicit nil

### UnsetSellerNote
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetSellerNote()`

UnsetSellerNote ensures that no value is present for SellerNote, not even an explicit nil
### GetPriceLowCents

`func (o *CreateMagicScanSynthesizeResponsePayload) GetPriceLowCents() float32`

GetPriceLowCents returns the PriceLowCents field if non-nil, zero value otherwise.

### GetPriceLowCentsOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetPriceLowCentsOk() (*float32, bool)`

GetPriceLowCentsOk returns a tuple with the PriceLowCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceLowCents

`func (o *CreateMagicScanSynthesizeResponsePayload) SetPriceLowCents(v float32)`

SetPriceLowCents sets PriceLowCents field to given value.

### HasPriceLowCents

`func (o *CreateMagicScanSynthesizeResponsePayload) HasPriceLowCents() bool`

HasPriceLowCents returns a boolean if a field has been set.

### SetPriceLowCentsNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetPriceLowCentsNil(b bool)`

 SetPriceLowCentsNil sets the value for PriceLowCents to be an explicit nil

### UnsetPriceLowCents
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetPriceLowCents()`

UnsetPriceLowCents ensures that no value is present for PriceLowCents, not even an explicit nil
### GetPriceMedianCents

`func (o *CreateMagicScanSynthesizeResponsePayload) GetPriceMedianCents() float32`

GetPriceMedianCents returns the PriceMedianCents field if non-nil, zero value otherwise.

### GetPriceMedianCentsOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetPriceMedianCentsOk() (*float32, bool)`

GetPriceMedianCentsOk returns a tuple with the PriceMedianCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceMedianCents

`func (o *CreateMagicScanSynthesizeResponsePayload) SetPriceMedianCents(v float32)`

SetPriceMedianCents sets PriceMedianCents field to given value.

### HasPriceMedianCents

`func (o *CreateMagicScanSynthesizeResponsePayload) HasPriceMedianCents() bool`

HasPriceMedianCents returns a boolean if a field has been set.

### SetPriceMedianCentsNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetPriceMedianCentsNil(b bool)`

 SetPriceMedianCentsNil sets the value for PriceMedianCents to be an explicit nil

### UnsetPriceMedianCents
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetPriceMedianCents()`

UnsetPriceMedianCents ensures that no value is present for PriceMedianCents, not even an explicit nil
### GetPriceHighCents

`func (o *CreateMagicScanSynthesizeResponsePayload) GetPriceHighCents() float32`

GetPriceHighCents returns the PriceHighCents field if non-nil, zero value otherwise.

### GetPriceHighCentsOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetPriceHighCentsOk() (*float32, bool)`

GetPriceHighCentsOk returns a tuple with the PriceHighCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceHighCents

`func (o *CreateMagicScanSynthesizeResponsePayload) SetPriceHighCents(v float32)`

SetPriceHighCents sets PriceHighCents field to given value.

### HasPriceHighCents

`func (o *CreateMagicScanSynthesizeResponsePayload) HasPriceHighCents() bool`

HasPriceHighCents returns a boolean if a field has been set.

### SetPriceHighCentsNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetPriceHighCentsNil(b bool)`

 SetPriceHighCentsNil sets the value for PriceHighCents to be an explicit nil

### UnsetPriceHighCents
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetPriceHighCents()`

UnsetPriceHighCents ensures that no value is present for PriceHighCents, not even an explicit nil
### GetQuantity

`func (o *CreateMagicScanSynthesizeResponsePayload) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateMagicScanSynthesizeResponsePayload) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CreateMagicScanSynthesizeResponsePayload) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### SetQuantityNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetQuantityNil(b bool)`

 SetQuantityNil sets the value for Quantity to be an explicit nil

### UnsetQuantity
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetQuantity()`

UnsetQuantity ensures that no value is present for Quantity, not even an explicit nil
### GetInventoryQuantity

`func (o *CreateMagicScanSynthesizeResponsePayload) GetInventoryQuantity() float32`

GetInventoryQuantity returns the InventoryQuantity field if non-nil, zero value otherwise.

### GetInventoryQuantityOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetInventoryQuantityOk() (*float32, bool)`

GetInventoryQuantityOk returns a tuple with the InventoryQuantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryQuantity

`func (o *CreateMagicScanSynthesizeResponsePayload) SetInventoryQuantity(v float32)`

SetInventoryQuantity sets InventoryQuantity field to given value.

### HasInventoryQuantity

`func (o *CreateMagicScanSynthesizeResponsePayload) HasInventoryQuantity() bool`

HasInventoryQuantity returns a boolean if a field has been set.

### SetInventoryQuantityNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetInventoryQuantityNil(b bool)`

 SetInventoryQuantityNil sets the value for InventoryQuantity to be an explicit nil

### UnsetInventoryQuantity
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetInventoryQuantity()`

UnsetInventoryQuantity ensures that no value is present for InventoryQuantity, not even an explicit nil
### GetUnitsPerListing

`func (o *CreateMagicScanSynthesizeResponsePayload) GetUnitsPerListing() float32`

GetUnitsPerListing returns the UnitsPerListing field if non-nil, zero value otherwise.

### GetUnitsPerListingOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetUnitsPerListingOk() (*float32, bool)`

GetUnitsPerListingOk returns a tuple with the UnitsPerListing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnitsPerListing

`func (o *CreateMagicScanSynthesizeResponsePayload) SetUnitsPerListing(v float32)`

SetUnitsPerListing sets UnitsPerListing field to given value.

### HasUnitsPerListing

`func (o *CreateMagicScanSynthesizeResponsePayload) HasUnitsPerListing() bool`

HasUnitsPerListing returns a boolean if a field has been set.

### SetUnitsPerListingNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetUnitsPerListingNil(b bool)`

 SetUnitsPerListingNil sets the value for UnitsPerListing to be an explicit nil

### UnsetUnitsPerListing
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetUnitsPerListing()`

UnsetUnitsPerListing ensures that no value is present for UnitsPerListing, not even an explicit nil
### GetCostBasisCents

`func (o *CreateMagicScanSynthesizeResponsePayload) GetCostBasisCents() float32`

GetCostBasisCents returns the CostBasisCents field if non-nil, zero value otherwise.

### GetCostBasisCentsOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetCostBasisCentsOk() (*float32, bool)`

GetCostBasisCentsOk returns a tuple with the CostBasisCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostBasisCents

`func (o *CreateMagicScanSynthesizeResponsePayload) SetCostBasisCents(v float32)`

SetCostBasisCents sets CostBasisCents field to given value.

### HasCostBasisCents

`func (o *CreateMagicScanSynthesizeResponsePayload) HasCostBasisCents() bool`

HasCostBasisCents returns a boolean if a field has been set.

### SetCostBasisCentsNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetCostBasisCentsNil(b bool)`

 SetCostBasisCentsNil sets the value for CostBasisCents to be an explicit nil

### UnsetCostBasisCents
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetCostBasisCents()`

UnsetCostBasisCents ensures that no value is present for CostBasisCents, not even an explicit nil
### GetChosenDimensions

`func (o *CreateMagicScanSynthesizeResponsePayload) GetChosenDimensions() CreateMagicScanSynthesizeResponsePayloadChosenDimensions`

GetChosenDimensions returns the ChosenDimensions field if non-nil, zero value otherwise.

### GetChosenDimensionsOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetChosenDimensionsOk() (*CreateMagicScanSynthesizeResponsePayloadChosenDimensions, bool)`

GetChosenDimensionsOk returns a tuple with the ChosenDimensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChosenDimensions

`func (o *CreateMagicScanSynthesizeResponsePayload) SetChosenDimensions(v CreateMagicScanSynthesizeResponsePayloadChosenDimensions)`

SetChosenDimensions sets ChosenDimensions field to given value.

### HasChosenDimensions

`func (o *CreateMagicScanSynthesizeResponsePayload) HasChosenDimensions() bool`

HasChosenDimensions returns a boolean if a field has been set.

### SetChosenDimensionsNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetChosenDimensionsNil(b bool)`

 SetChosenDimensionsNil sets the value for ChosenDimensions to be an explicit nil

### UnsetChosenDimensions
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetChosenDimensions()`

UnsetChosenDimensions ensures that no value is present for ChosenDimensions, not even an explicit nil
### GetImageUrls

`func (o *CreateMagicScanSynthesizeResponsePayload) GetImageUrls() []string`

GetImageUrls returns the ImageUrls field if non-nil, zero value otherwise.

### GetImageUrlsOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetImageUrlsOk() (*[]string, bool)`

GetImageUrlsOk returns a tuple with the ImageUrls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrls

`func (o *CreateMagicScanSynthesizeResponsePayload) SetImageUrls(v []string)`

SetImageUrls sets ImageUrls field to given value.

### HasImageUrls

`func (o *CreateMagicScanSynthesizeResponsePayload) HasImageUrls() bool`

HasImageUrls returns a boolean if a field has been set.

### SetImageUrlsNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetImageUrlsNil(b bool)`

 SetImageUrlsNil sets the value for ImageUrls to be an explicit nil

### UnsetImageUrls
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetImageUrls()`

UnsetImageUrls ensures that no value is present for ImageUrls, not even an explicit nil
### GetCategory

`func (o *CreateMagicScanSynthesizeResponsePayload) GetCategory() CreateMagicScanSynthesizeResponsePayloadCategory`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetCategoryOk() (*CreateMagicScanSynthesizeResponsePayloadCategory, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *CreateMagicScanSynthesizeResponsePayload) SetCategory(v CreateMagicScanSynthesizeResponsePayloadCategory)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *CreateMagicScanSynthesizeResponsePayload) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetDepartment

`func (o *CreateMagicScanSynthesizeResponsePayload) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *CreateMagicScanSynthesizeResponsePayload) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *CreateMagicScanSynthesizeResponsePayload) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetGender

`func (o *CreateMagicScanSynthesizeResponsePayload) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CreateMagicScanSynthesizeResponsePayload) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CreateMagicScanSynthesizeResponsePayload) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetStyle

`func (o *CreateMagicScanSynthesizeResponsePayload) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *CreateMagicScanSynthesizeResponsePayload) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *CreateMagicScanSynthesizeResponsePayload) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetPattern

`func (o *CreateMagicScanSynthesizeResponsePayload) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *CreateMagicScanSynthesizeResponsePayload) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *CreateMagicScanSynthesizeResponsePayload) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPatternNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil
### GetItemType

`func (o *CreateMagicScanSynthesizeResponsePayload) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *CreateMagicScanSynthesizeResponsePayload) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *CreateMagicScanSynthesizeResponsePayload) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetUpc

`func (o *CreateMagicScanSynthesizeResponsePayload) GetUpc() string`

GetUpc returns the Upc field if non-nil, zero value otherwise.

### GetUpcOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetUpcOk() (*string, bool)`

GetUpcOk returns a tuple with the Upc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpc

`func (o *CreateMagicScanSynthesizeResponsePayload) SetUpc(v string)`

SetUpc sets Upc field to given value.

### HasUpc

`func (o *CreateMagicScanSynthesizeResponsePayload) HasUpc() bool`

HasUpc returns a boolean if a field has been set.

### SetUpcNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetUpcNil(b bool)`

 SetUpcNil sets the value for Upc to be an explicit nil

### UnsetUpc
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetUpc()`

UnsetUpc ensures that no value is present for Upc, not even an explicit nil
### GetHandlingTimeDays

`func (o *CreateMagicScanSynthesizeResponsePayload) GetHandlingTimeDays() string`

GetHandlingTimeDays returns the HandlingTimeDays field if non-nil, zero value otherwise.

### GetHandlingTimeDaysOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetHandlingTimeDaysOk() (*string, bool)`

GetHandlingTimeDaysOk returns a tuple with the HandlingTimeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingTimeDays

`func (o *CreateMagicScanSynthesizeResponsePayload) SetHandlingTimeDays(v string)`

SetHandlingTimeDays sets HandlingTimeDays field to given value.

### HasHandlingTimeDays

`func (o *CreateMagicScanSynthesizeResponsePayload) HasHandlingTimeDays() bool`

HasHandlingTimeDays returns a boolean if a field has been set.

### SetHandlingTimeDaysNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetHandlingTimeDaysNil(b bool)`

 SetHandlingTimeDaysNil sets the value for HandlingTimeDays to be an explicit nil

### UnsetHandlingTimeDays
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetHandlingTimeDays()`

UnsetHandlingTimeDays ensures that no value is present for HandlingTimeDays, not even an explicit nil
### GetReturnWindowDays

`func (o *CreateMagicScanSynthesizeResponsePayload) GetReturnWindowDays() string`

GetReturnWindowDays returns the ReturnWindowDays field if non-nil, zero value otherwise.

### GetReturnWindowDaysOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetReturnWindowDaysOk() (*string, bool)`

GetReturnWindowDaysOk returns a tuple with the ReturnWindowDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReturnWindowDays

`func (o *CreateMagicScanSynthesizeResponsePayload) SetReturnWindowDays(v string)`

SetReturnWindowDays sets ReturnWindowDays field to given value.

### HasReturnWindowDays

`func (o *CreateMagicScanSynthesizeResponsePayload) HasReturnWindowDays() bool`

HasReturnWindowDays returns a boolean if a field has been set.

### SetReturnWindowDaysNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetReturnWindowDaysNil(b bool)`

 SetReturnWindowDaysNil sets the value for ReturnWindowDays to be an explicit nil

### UnsetReturnWindowDays
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetReturnWindowDays()`

UnsetReturnWindowDays ensures that no value is present for ReturnWindowDays, not even an explicit nil
### GetTags

`func (o *CreateMagicScanSynthesizeResponsePayload) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateMagicScanSynthesizeResponsePayload) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CreateMagicScanSynthesizeResponsePayload) HasTags() bool`

HasTags returns a boolean if a field has been set.

### SetTagsNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetTagsNil(b bool)`

 SetTagsNil sets the value for Tags to be an explicit nil

### UnsetTags
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetTags()`

UnsetTags ensures that no value is present for Tags, not even an explicit nil
### GetItemSpecifics

`func (o *CreateMagicScanSynthesizeResponsePayload) GetItemSpecifics() string`

GetItemSpecifics returns the ItemSpecifics field if non-nil, zero value otherwise.

### GetItemSpecificsOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetItemSpecificsOk() (*string, bool)`

GetItemSpecificsOk returns a tuple with the ItemSpecifics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemSpecifics

`func (o *CreateMagicScanSynthesizeResponsePayload) SetItemSpecifics(v string)`

SetItemSpecifics sets ItemSpecifics field to given value.

### HasItemSpecifics

`func (o *CreateMagicScanSynthesizeResponsePayload) HasItemSpecifics() bool`

HasItemSpecifics returns a boolean if a field has been set.

### SetItemSpecificsNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetItemSpecificsNil(b bool)`

 SetItemSpecificsNil sets the value for ItemSpecifics to be an explicit nil

### UnsetItemSpecifics
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetItemSpecifics()`

UnsetItemSpecifics ensures that no value is present for ItemSpecifics, not even an explicit nil
### GetPerPlatformOverrides

`func (o *CreateMagicScanSynthesizeResponsePayload) GetPerPlatformOverrides() map[string]interface{}`

GetPerPlatformOverrides returns the PerPlatformOverrides field if non-nil, zero value otherwise.

### GetPerPlatformOverridesOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetPerPlatformOverridesOk() (*map[string]interface{}, bool)`

GetPerPlatformOverridesOk returns a tuple with the PerPlatformOverrides field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPerPlatformOverrides

`func (o *CreateMagicScanSynthesizeResponsePayload) SetPerPlatformOverrides(v map[string]interface{})`

SetPerPlatformOverrides sets PerPlatformOverrides field to given value.


### GetConfidence

`func (o *CreateMagicScanSynthesizeResponsePayload) GetConfidence() float32`

GetConfidence returns the Confidence field if non-nil, zero value otherwise.

### GetConfidenceOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetConfidenceOk() (*float32, bool)`

GetConfidenceOk returns a tuple with the Confidence field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfidence

`func (o *CreateMagicScanSynthesizeResponsePayload) SetConfidence(v float32)`

SetConfidence sets Confidence field to given value.


### GetNotes

`func (o *CreateMagicScanSynthesizeResponsePayload) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *CreateMagicScanSynthesizeResponsePayload) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *CreateMagicScanSynthesizeResponsePayload) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetSectionApplicability

`func (o *CreateMagicScanSynthesizeResponsePayload) GetSectionApplicability() CreateMagicScanSynthesizeResponsePayloadSectionApplicability`

GetSectionApplicability returns the SectionApplicability field if non-nil, zero value otherwise.

### GetSectionApplicabilityOk

`func (o *CreateMagicScanSynthesizeResponsePayload) GetSectionApplicabilityOk() (*CreateMagicScanSynthesizeResponsePayloadSectionApplicability, bool)`

GetSectionApplicabilityOk returns a tuple with the SectionApplicability field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSectionApplicability

`func (o *CreateMagicScanSynthesizeResponsePayload) SetSectionApplicability(v CreateMagicScanSynthesizeResponsePayloadSectionApplicability)`

SetSectionApplicability sets SectionApplicability field to given value.

### HasSectionApplicability

`func (o *CreateMagicScanSynthesizeResponsePayload) HasSectionApplicability() bool`

HasSectionApplicability returns a boolean if a field has been set.

### SetSectionApplicabilityNil

`func (o *CreateMagicScanSynthesizeResponsePayload) SetSectionApplicabilityNil(b bool)`

 SetSectionApplicabilityNil sets the value for SectionApplicability to be an explicit nil

### UnsetSectionApplicability
`func (o *CreateMagicScanSynthesizeResponsePayload) UnsetSectionApplicability()`

UnsetSectionApplicability ensures that no value is present for SectionApplicability, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


