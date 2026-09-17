# CreateListingByIdResponseItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EffectiveTitle** | **string** |  | 
**EffectivePrice** | Pointer to **NullableString** |  | [optional] 
**EffectiveDescription** | Pointer to **NullableString** |  | [optional] 
**EffectiveImages** | **[]string** |  | 
**EffectiveBrand** | Pointer to **NullableString** |  | [optional] 
**EffectiveCondition** | Pointer to **NullableString** |  | [optional] 
**EffectiveSize** | Pointer to **NullableString** |  | [optional] 
**EffectiveSku** | Pointer to **NullableString** |  | [optional] 
**EffectiveColor** | **[]string** |  | 
**EffectiveTags** | **[]string** |  | 
**PlatformListings** | **[]map[string]interface{}** |  | 
**InventoryItemId** | Pointer to **NullableString** |  | [optional] 
**Id** | **string** |  | 
**UserId** | **string** |  | 
**Name** | Pointer to **NullableString** |  | [optional] 
**Title** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DescriptionHtml** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableString** |  | [optional] 
**Images** | Pointer to **[]string** |  | [optional] 
**VideoUrl** | Pointer to **NullableString** | Optional single product video (R2/CDN URL). Shown on the Crossly buyer page. | [optional] 
**Status** | **string** |  | 
**Condition** | Pointer to **NullableString** |  | [optional] 
**GradeKey** | Pointer to **NullableString** | Third-party grading, when the item is slabbed. Migration 0277.    Separate from &#x60;condition&#x60; on purpose and never derived from it: a grade  is a claim about what a GRADING COMPANY certified, and inferring \&quot;PSA 10\&quot;  from a coarse condition would be a false authenticity claim. It is also  never filled from our own AI estimate (&#x60;bulk_market_items.grade&#x60;), which  carries an explicit \&quot;not a professional grade\&quot; disclaimer.    &#x60;gradeKey&#x60; is the canonical form from &#x60;gradeKey()&#x60; in  shared/constants/graders.ts; &#x60;grading&#x60; holds the full GradingInfo  including the cert number and whether a cert lookup verified it. | [optional] 
**Grading** | Pointer to [**NullableListListingsItemGrading**](ListListingsItemGrading.md) |  | [optional] 
**Brand** | Pointer to **NullableString** |  | [optional] 
**Size** | Pointer to **NullableString** |  | [optional] 
**Material** | Pointer to **NullableString** | Migration 0179 — see the matching fields on inventory_items above. | [optional] 
**Style** | Pointer to **NullableString** |  | [optional] 
**Pattern** | Pointer to **NullableString** |  | [optional] 
**Department** | Pointer to **NullableString** |  | [optional] 
**Gender** | Pointer to **NullableString** |  | [optional] 
**ItemType** | Pointer to **NullableString** |  | [optional] 
**SizeSystem** | Pointer to **NullableString** |  | [optional] 
**Color** | **[]string** |  | 
**Tags** | **[]string** |  | 
**Sku** | Pointer to **NullableString** |  | [optional] 
**Quantity** | **float32** |  | 
**QuantityAvailable** | **float32** |  | 
**WeightLb** | Pointer to **NullableString** |  | [optional] 
**WeightOz** | Pointer to **NullableString** |  | [optional] 
**DimensionLIn** | Pointer to **NullableString** |  | [optional] 
**DimensionWIn** | Pointer to **NullableString** |  | [optional] 
**DimensionHIn** | Pointer to **NullableString** |  | [optional] 
**PublishAt** | Pointer to **NullableTime** | Scheduled go-live time. When set on a draft, the listing-scheduler  worker waits until this passes then dispatches the crosspost to  scheduledPlatforms and flips status from &#39;draft&#39; to &#39;active&#39;. | [optional] 
**ScheduledPlatforms** | Pointer to **[]string** | Which platforms to publish to when publishAt fires. JSON array of  platform ids. Null/empty &#x3D; scheduler skips (listing won&#39;t auto-  publish, even after publishAt — gives the seller an escape hatch). | [optional] 
**ParentListingId** | Pointer to **NullableString** | Parent listing when this row is a CHILD in a listing chain. Null &#x3D;  standalone. What being a child means depends on the parent&#39;s  &#x60;groupKind&#x60; — see it. | [optional] 
**IsBundle** | **bool** |  | 
**AutomationAssignedRuleIds** | **[]string** | Per-listing automation overrides. See migration 0098.     automationAssignedRuleIds  — force-include for these rules   automationBlockedRuleIds   — exempt from these rules   automationAssignedChainIds — force-include for these workflow chains   automationBlockedChainIds  — exempt from these workflow chains | 
**AutomationBlockedRuleIds** | **[]string** |  | 
**AutomationAssignedChainIds** | **[]string** |  | 
**AutomationBlockedChainIds** | **[]string** |  | 
**HsCode** | Pointer to **NullableString** | Harmonised System customs code — international shipping declarations. | [optional] 
**CountryOfOrigin** | Pointer to **NullableString** | Customs country of origin. Distinct from the seller&#39;s location. | [optional] 
**PriceFloorCents** | Pointer to **NullableFloat32** | Never let a repricing rule go below this. On the ITEM because it is a  fact about the thing owned, not about any one rule — \&quot;this jacket never  goes below $45\&quot; should apply to every rule, and before this it was  expressible only as one rule per jacket. Listings inherit when null. | [optional] 
**FloorIsNet** | **bool** | When true the floor is a TAKE-HOME target, converted to a per-platform  gross at reprice time. A gross floor is four different promises across  four platforms; this is the one number a seller actually cares about. | 
**DelistedAt** | Pointer to **NullableTime** |  | [optional] 
**SoldAt** | Pointer to **NullableTime** |  | [optional] 
**Source** | **string** | Mirrors inventory_items.source. &#39;manual&#39; for every seller-created  listing; external-stub.ts sets &#39;external_sale&#39; on the synthetic  listing it fabricates for a sale detected on a platform id Crossly  never listed — those rows have no real photos/description of their  own (everything is lifted from the platform&#39;s sale payload) and are  otherwise indistinguishable from a real listing in the UI. | 
**DuplicateOfListingId** | Pointer to **NullableString** | Set when import&#39;s bin-packing (see _import-one.ts) created THIS  listing to hold a same-platform straggler it couldn&#39;t fit onto an  existing candidate listing for the same physical item — points at  the primary/first candidate. Purely informational: this listing  is a real, independently listable/delistable row, not a shadow.  Null for every ordinarily-created listing. | [optional] 
**ClientDraftId** | Pointer to **NullableString** | UUID minted on a seller&#39;s machine for a draft written offline.    The idempotency key for desktop sync. The failure it guards is a POST  that succeeds server-side whose reply is lost — the client cannot tell  that from a failure, retries, and one item becomes two live listings  against one piece of stock. Unique per user (partial index, migration  0268); null for every listing that did not come from an offline draft. | [optional] 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 

## Methods

### NewCreateListingByIdResponseItems

`func NewCreateListingByIdResponseItems(effectiveTitle string, effectiveImages []string, effectiveColor []string, effectiveTags []string, platformListings []map[string]interface{}, id string, userId string, status string, color []string, tags []string, quantity float32, quantityAvailable float32, isBundle bool, automationAssignedRuleIds []string, automationBlockedRuleIds []string, automationAssignedChainIds []string, automationBlockedChainIds []string, floorIsNet bool, source string, createdAt time.Time, updatedAt time.Time, ) *CreateListingByIdResponseItems`

NewCreateListingByIdResponseItems instantiates a new CreateListingByIdResponseItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateListingByIdResponseItemsWithDefaults

`func NewCreateListingByIdResponseItemsWithDefaults() *CreateListingByIdResponseItems`

NewCreateListingByIdResponseItemsWithDefaults instantiates a new CreateListingByIdResponseItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEffectiveTitle

`func (o *CreateListingByIdResponseItems) GetEffectiveTitle() string`

GetEffectiveTitle returns the EffectiveTitle field if non-nil, zero value otherwise.

### GetEffectiveTitleOk

`func (o *CreateListingByIdResponseItems) GetEffectiveTitleOk() (*string, bool)`

GetEffectiveTitleOk returns a tuple with the EffectiveTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTitle

`func (o *CreateListingByIdResponseItems) SetEffectiveTitle(v string)`

SetEffectiveTitle sets EffectiveTitle field to given value.


### GetEffectivePrice

`func (o *CreateListingByIdResponseItems) GetEffectivePrice() string`

GetEffectivePrice returns the EffectivePrice field if non-nil, zero value otherwise.

### GetEffectivePriceOk

`func (o *CreateListingByIdResponseItems) GetEffectivePriceOk() (*string, bool)`

GetEffectivePriceOk returns a tuple with the EffectivePrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectivePrice

`func (o *CreateListingByIdResponseItems) SetEffectivePrice(v string)`

SetEffectivePrice sets EffectivePrice field to given value.

### HasEffectivePrice

`func (o *CreateListingByIdResponseItems) HasEffectivePrice() bool`

HasEffectivePrice returns a boolean if a field has been set.

### SetEffectivePriceNil

`func (o *CreateListingByIdResponseItems) SetEffectivePriceNil(b bool)`

 SetEffectivePriceNil sets the value for EffectivePrice to be an explicit nil

### UnsetEffectivePrice
`func (o *CreateListingByIdResponseItems) UnsetEffectivePrice()`

UnsetEffectivePrice ensures that no value is present for EffectivePrice, not even an explicit nil
### GetEffectiveDescription

`func (o *CreateListingByIdResponseItems) GetEffectiveDescription() string`

GetEffectiveDescription returns the EffectiveDescription field if non-nil, zero value otherwise.

### GetEffectiveDescriptionOk

`func (o *CreateListingByIdResponseItems) GetEffectiveDescriptionOk() (*string, bool)`

GetEffectiveDescriptionOk returns a tuple with the EffectiveDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveDescription

`func (o *CreateListingByIdResponseItems) SetEffectiveDescription(v string)`

SetEffectiveDescription sets EffectiveDescription field to given value.

### HasEffectiveDescription

`func (o *CreateListingByIdResponseItems) HasEffectiveDescription() bool`

HasEffectiveDescription returns a boolean if a field has been set.

### SetEffectiveDescriptionNil

`func (o *CreateListingByIdResponseItems) SetEffectiveDescriptionNil(b bool)`

 SetEffectiveDescriptionNil sets the value for EffectiveDescription to be an explicit nil

### UnsetEffectiveDescription
`func (o *CreateListingByIdResponseItems) UnsetEffectiveDescription()`

UnsetEffectiveDescription ensures that no value is present for EffectiveDescription, not even an explicit nil
### GetEffectiveImages

`func (o *CreateListingByIdResponseItems) GetEffectiveImages() []string`

GetEffectiveImages returns the EffectiveImages field if non-nil, zero value otherwise.

### GetEffectiveImagesOk

`func (o *CreateListingByIdResponseItems) GetEffectiveImagesOk() (*[]string, bool)`

GetEffectiveImagesOk returns a tuple with the EffectiveImages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveImages

`func (o *CreateListingByIdResponseItems) SetEffectiveImages(v []string)`

SetEffectiveImages sets EffectiveImages field to given value.


### GetEffectiveBrand

`func (o *CreateListingByIdResponseItems) GetEffectiveBrand() string`

GetEffectiveBrand returns the EffectiveBrand field if non-nil, zero value otherwise.

### GetEffectiveBrandOk

`func (o *CreateListingByIdResponseItems) GetEffectiveBrandOk() (*string, bool)`

GetEffectiveBrandOk returns a tuple with the EffectiveBrand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveBrand

`func (o *CreateListingByIdResponseItems) SetEffectiveBrand(v string)`

SetEffectiveBrand sets EffectiveBrand field to given value.

### HasEffectiveBrand

`func (o *CreateListingByIdResponseItems) HasEffectiveBrand() bool`

HasEffectiveBrand returns a boolean if a field has been set.

### SetEffectiveBrandNil

`func (o *CreateListingByIdResponseItems) SetEffectiveBrandNil(b bool)`

 SetEffectiveBrandNil sets the value for EffectiveBrand to be an explicit nil

### UnsetEffectiveBrand
`func (o *CreateListingByIdResponseItems) UnsetEffectiveBrand()`

UnsetEffectiveBrand ensures that no value is present for EffectiveBrand, not even an explicit nil
### GetEffectiveCondition

`func (o *CreateListingByIdResponseItems) GetEffectiveCondition() string`

GetEffectiveCondition returns the EffectiveCondition field if non-nil, zero value otherwise.

### GetEffectiveConditionOk

`func (o *CreateListingByIdResponseItems) GetEffectiveConditionOk() (*string, bool)`

GetEffectiveConditionOk returns a tuple with the EffectiveCondition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveCondition

`func (o *CreateListingByIdResponseItems) SetEffectiveCondition(v string)`

SetEffectiveCondition sets EffectiveCondition field to given value.

### HasEffectiveCondition

`func (o *CreateListingByIdResponseItems) HasEffectiveCondition() bool`

HasEffectiveCondition returns a boolean if a field has been set.

### SetEffectiveConditionNil

`func (o *CreateListingByIdResponseItems) SetEffectiveConditionNil(b bool)`

 SetEffectiveConditionNil sets the value for EffectiveCondition to be an explicit nil

### UnsetEffectiveCondition
`func (o *CreateListingByIdResponseItems) UnsetEffectiveCondition()`

UnsetEffectiveCondition ensures that no value is present for EffectiveCondition, not even an explicit nil
### GetEffectiveSize

`func (o *CreateListingByIdResponseItems) GetEffectiveSize() string`

GetEffectiveSize returns the EffectiveSize field if non-nil, zero value otherwise.

### GetEffectiveSizeOk

`func (o *CreateListingByIdResponseItems) GetEffectiveSizeOk() (*string, bool)`

GetEffectiveSizeOk returns a tuple with the EffectiveSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveSize

`func (o *CreateListingByIdResponseItems) SetEffectiveSize(v string)`

SetEffectiveSize sets EffectiveSize field to given value.

### HasEffectiveSize

`func (o *CreateListingByIdResponseItems) HasEffectiveSize() bool`

HasEffectiveSize returns a boolean if a field has been set.

### SetEffectiveSizeNil

`func (o *CreateListingByIdResponseItems) SetEffectiveSizeNil(b bool)`

 SetEffectiveSizeNil sets the value for EffectiveSize to be an explicit nil

### UnsetEffectiveSize
`func (o *CreateListingByIdResponseItems) UnsetEffectiveSize()`

UnsetEffectiveSize ensures that no value is present for EffectiveSize, not even an explicit nil
### GetEffectiveSku

`func (o *CreateListingByIdResponseItems) GetEffectiveSku() string`

GetEffectiveSku returns the EffectiveSku field if non-nil, zero value otherwise.

### GetEffectiveSkuOk

`func (o *CreateListingByIdResponseItems) GetEffectiveSkuOk() (*string, bool)`

GetEffectiveSkuOk returns a tuple with the EffectiveSku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveSku

`func (o *CreateListingByIdResponseItems) SetEffectiveSku(v string)`

SetEffectiveSku sets EffectiveSku field to given value.

### HasEffectiveSku

`func (o *CreateListingByIdResponseItems) HasEffectiveSku() bool`

HasEffectiveSku returns a boolean if a field has been set.

### SetEffectiveSkuNil

`func (o *CreateListingByIdResponseItems) SetEffectiveSkuNil(b bool)`

 SetEffectiveSkuNil sets the value for EffectiveSku to be an explicit nil

### UnsetEffectiveSku
`func (o *CreateListingByIdResponseItems) UnsetEffectiveSku()`

UnsetEffectiveSku ensures that no value is present for EffectiveSku, not even an explicit nil
### GetEffectiveColor

`func (o *CreateListingByIdResponseItems) GetEffectiveColor() []string`

GetEffectiveColor returns the EffectiveColor field if non-nil, zero value otherwise.

### GetEffectiveColorOk

`func (o *CreateListingByIdResponseItems) GetEffectiveColorOk() (*[]string, bool)`

GetEffectiveColorOk returns a tuple with the EffectiveColor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveColor

`func (o *CreateListingByIdResponseItems) SetEffectiveColor(v []string)`

SetEffectiveColor sets EffectiveColor field to given value.


### GetEffectiveTags

`func (o *CreateListingByIdResponseItems) GetEffectiveTags() []string`

GetEffectiveTags returns the EffectiveTags field if non-nil, zero value otherwise.

### GetEffectiveTagsOk

`func (o *CreateListingByIdResponseItems) GetEffectiveTagsOk() (*[]string, bool)`

GetEffectiveTagsOk returns a tuple with the EffectiveTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEffectiveTags

`func (o *CreateListingByIdResponseItems) SetEffectiveTags(v []string)`

SetEffectiveTags sets EffectiveTags field to given value.


### GetPlatformListings

`func (o *CreateListingByIdResponseItems) GetPlatformListings() []map[string]interface{}`

GetPlatformListings returns the PlatformListings field if non-nil, zero value otherwise.

### GetPlatformListingsOk

`func (o *CreateListingByIdResponseItems) GetPlatformListingsOk() (*[]map[string]interface{}, bool)`

GetPlatformListingsOk returns a tuple with the PlatformListings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformListings

`func (o *CreateListingByIdResponseItems) SetPlatformListings(v []map[string]interface{})`

SetPlatformListings sets PlatformListings field to given value.


### GetInventoryItemId

`func (o *CreateListingByIdResponseItems) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *CreateListingByIdResponseItems) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *CreateListingByIdResponseItems) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *CreateListingByIdResponseItems) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### SetInventoryItemIdNil

`func (o *CreateListingByIdResponseItems) SetInventoryItemIdNil(b bool)`

 SetInventoryItemIdNil sets the value for InventoryItemId to be an explicit nil

### UnsetInventoryItemId
`func (o *CreateListingByIdResponseItems) UnsetInventoryItemId()`

UnsetInventoryItemId ensures that no value is present for InventoryItemId, not even an explicit nil
### GetId

`func (o *CreateListingByIdResponseItems) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CreateListingByIdResponseItems) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CreateListingByIdResponseItems) SetId(v string)`

SetId sets Id field to given value.


### GetUserId

`func (o *CreateListingByIdResponseItems) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *CreateListingByIdResponseItems) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *CreateListingByIdResponseItems) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetName

`func (o *CreateListingByIdResponseItems) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateListingByIdResponseItems) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateListingByIdResponseItems) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CreateListingByIdResponseItems) HasName() bool`

HasName returns a boolean if a field has been set.

### SetNameNil

`func (o *CreateListingByIdResponseItems) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *CreateListingByIdResponseItems) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetTitle

`func (o *CreateListingByIdResponseItems) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *CreateListingByIdResponseItems) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *CreateListingByIdResponseItems) SetTitle(v string)`

SetTitle sets Title field to given value.

### HasTitle

`func (o *CreateListingByIdResponseItems) HasTitle() bool`

HasTitle returns a boolean if a field has been set.

### SetTitleNil

`func (o *CreateListingByIdResponseItems) SetTitleNil(b bool)`

 SetTitleNil sets the value for Title to be an explicit nil

### UnsetTitle
`func (o *CreateListingByIdResponseItems) UnsetTitle()`

UnsetTitle ensures that no value is present for Title, not even an explicit nil
### GetDescription

`func (o *CreateListingByIdResponseItems) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateListingByIdResponseItems) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateListingByIdResponseItems) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateListingByIdResponseItems) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *CreateListingByIdResponseItems) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *CreateListingByIdResponseItems) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDescriptionHtml

`func (o *CreateListingByIdResponseItems) GetDescriptionHtml() string`

GetDescriptionHtml returns the DescriptionHtml field if non-nil, zero value otherwise.

### GetDescriptionHtmlOk

`func (o *CreateListingByIdResponseItems) GetDescriptionHtmlOk() (*string, bool)`

GetDescriptionHtmlOk returns a tuple with the DescriptionHtml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptionHtml

`func (o *CreateListingByIdResponseItems) SetDescriptionHtml(v string)`

SetDescriptionHtml sets DescriptionHtml field to given value.

### HasDescriptionHtml

`func (o *CreateListingByIdResponseItems) HasDescriptionHtml() bool`

HasDescriptionHtml returns a boolean if a field has been set.

### SetDescriptionHtmlNil

`func (o *CreateListingByIdResponseItems) SetDescriptionHtmlNil(b bool)`

 SetDescriptionHtmlNil sets the value for DescriptionHtml to be an explicit nil

### UnsetDescriptionHtml
`func (o *CreateListingByIdResponseItems) UnsetDescriptionHtml()`

UnsetDescriptionHtml ensures that no value is present for DescriptionHtml, not even an explicit nil
### GetPrice

`func (o *CreateListingByIdResponseItems) GetPrice() string`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CreateListingByIdResponseItems) GetPriceOk() (*string, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CreateListingByIdResponseItems) SetPrice(v string)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CreateListingByIdResponseItems) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *CreateListingByIdResponseItems) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *CreateListingByIdResponseItems) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetImages

`func (o *CreateListingByIdResponseItems) GetImages() []string`

GetImages returns the Images field if non-nil, zero value otherwise.

### GetImagesOk

`func (o *CreateListingByIdResponseItems) GetImagesOk() (*[]string, bool)`

GetImagesOk returns a tuple with the Images field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImages

`func (o *CreateListingByIdResponseItems) SetImages(v []string)`

SetImages sets Images field to given value.

### HasImages

`func (o *CreateListingByIdResponseItems) HasImages() bool`

HasImages returns a boolean if a field has been set.

### SetImagesNil

`func (o *CreateListingByIdResponseItems) SetImagesNil(b bool)`

 SetImagesNil sets the value for Images to be an explicit nil

### UnsetImages
`func (o *CreateListingByIdResponseItems) UnsetImages()`

UnsetImages ensures that no value is present for Images, not even an explicit nil
### GetVideoUrl

`func (o *CreateListingByIdResponseItems) GetVideoUrl() string`

GetVideoUrl returns the VideoUrl field if non-nil, zero value otherwise.

### GetVideoUrlOk

`func (o *CreateListingByIdResponseItems) GetVideoUrlOk() (*string, bool)`

GetVideoUrlOk returns a tuple with the VideoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVideoUrl

`func (o *CreateListingByIdResponseItems) SetVideoUrl(v string)`

SetVideoUrl sets VideoUrl field to given value.

### HasVideoUrl

`func (o *CreateListingByIdResponseItems) HasVideoUrl() bool`

HasVideoUrl returns a boolean if a field has been set.

### SetVideoUrlNil

`func (o *CreateListingByIdResponseItems) SetVideoUrlNil(b bool)`

 SetVideoUrlNil sets the value for VideoUrl to be an explicit nil

### UnsetVideoUrl
`func (o *CreateListingByIdResponseItems) UnsetVideoUrl()`

UnsetVideoUrl ensures that no value is present for VideoUrl, not even an explicit nil
### GetStatus

`func (o *CreateListingByIdResponseItems) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CreateListingByIdResponseItems) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CreateListingByIdResponseItems) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetCondition

`func (o *CreateListingByIdResponseItems) GetCondition() string`

GetCondition returns the Condition field if non-nil, zero value otherwise.

### GetConditionOk

`func (o *CreateListingByIdResponseItems) GetConditionOk() (*string, bool)`

GetConditionOk returns a tuple with the Condition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCondition

`func (o *CreateListingByIdResponseItems) SetCondition(v string)`

SetCondition sets Condition field to given value.

### HasCondition

`func (o *CreateListingByIdResponseItems) HasCondition() bool`

HasCondition returns a boolean if a field has been set.

### SetConditionNil

`func (o *CreateListingByIdResponseItems) SetConditionNil(b bool)`

 SetConditionNil sets the value for Condition to be an explicit nil

### UnsetCondition
`func (o *CreateListingByIdResponseItems) UnsetCondition()`

UnsetCondition ensures that no value is present for Condition, not even an explicit nil
### GetGradeKey

`func (o *CreateListingByIdResponseItems) GetGradeKey() string`

GetGradeKey returns the GradeKey field if non-nil, zero value otherwise.

### GetGradeKeyOk

`func (o *CreateListingByIdResponseItems) GetGradeKeyOk() (*string, bool)`

GetGradeKeyOk returns a tuple with the GradeKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGradeKey

`func (o *CreateListingByIdResponseItems) SetGradeKey(v string)`

SetGradeKey sets GradeKey field to given value.

### HasGradeKey

`func (o *CreateListingByIdResponseItems) HasGradeKey() bool`

HasGradeKey returns a boolean if a field has been set.

### SetGradeKeyNil

`func (o *CreateListingByIdResponseItems) SetGradeKeyNil(b bool)`

 SetGradeKeyNil sets the value for GradeKey to be an explicit nil

### UnsetGradeKey
`func (o *CreateListingByIdResponseItems) UnsetGradeKey()`

UnsetGradeKey ensures that no value is present for GradeKey, not even an explicit nil
### GetGrading

`func (o *CreateListingByIdResponseItems) GetGrading() ListListingsItemGrading`

GetGrading returns the Grading field if non-nil, zero value otherwise.

### GetGradingOk

`func (o *CreateListingByIdResponseItems) GetGradingOk() (*ListListingsItemGrading, bool)`

GetGradingOk returns a tuple with the Grading field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGrading

`func (o *CreateListingByIdResponseItems) SetGrading(v ListListingsItemGrading)`

SetGrading sets Grading field to given value.

### HasGrading

`func (o *CreateListingByIdResponseItems) HasGrading() bool`

HasGrading returns a boolean if a field has been set.

### SetGradingNil

`func (o *CreateListingByIdResponseItems) SetGradingNil(b bool)`

 SetGradingNil sets the value for Grading to be an explicit nil

### UnsetGrading
`func (o *CreateListingByIdResponseItems) UnsetGrading()`

UnsetGrading ensures that no value is present for Grading, not even an explicit nil
### GetBrand

`func (o *CreateListingByIdResponseItems) GetBrand() string`

GetBrand returns the Brand field if non-nil, zero value otherwise.

### GetBrandOk

`func (o *CreateListingByIdResponseItems) GetBrandOk() (*string, bool)`

GetBrandOk returns a tuple with the Brand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrand

`func (o *CreateListingByIdResponseItems) SetBrand(v string)`

SetBrand sets Brand field to given value.

### HasBrand

`func (o *CreateListingByIdResponseItems) HasBrand() bool`

HasBrand returns a boolean if a field has been set.

### SetBrandNil

`func (o *CreateListingByIdResponseItems) SetBrandNil(b bool)`

 SetBrandNil sets the value for Brand to be an explicit nil

### UnsetBrand
`func (o *CreateListingByIdResponseItems) UnsetBrand()`

UnsetBrand ensures that no value is present for Brand, not even an explicit nil
### GetSize

`func (o *CreateListingByIdResponseItems) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *CreateListingByIdResponseItems) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *CreateListingByIdResponseItems) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *CreateListingByIdResponseItems) HasSize() bool`

HasSize returns a boolean if a field has been set.

### SetSizeNil

`func (o *CreateListingByIdResponseItems) SetSizeNil(b bool)`

 SetSizeNil sets the value for Size to be an explicit nil

### UnsetSize
`func (o *CreateListingByIdResponseItems) UnsetSize()`

UnsetSize ensures that no value is present for Size, not even an explicit nil
### GetMaterial

`func (o *CreateListingByIdResponseItems) GetMaterial() string`

GetMaterial returns the Material field if non-nil, zero value otherwise.

### GetMaterialOk

`func (o *CreateListingByIdResponseItems) GetMaterialOk() (*string, bool)`

GetMaterialOk returns a tuple with the Material field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaterial

`func (o *CreateListingByIdResponseItems) SetMaterial(v string)`

SetMaterial sets Material field to given value.

### HasMaterial

`func (o *CreateListingByIdResponseItems) HasMaterial() bool`

HasMaterial returns a boolean if a field has been set.

### SetMaterialNil

`func (o *CreateListingByIdResponseItems) SetMaterialNil(b bool)`

 SetMaterialNil sets the value for Material to be an explicit nil

### UnsetMaterial
`func (o *CreateListingByIdResponseItems) UnsetMaterial()`

UnsetMaterial ensures that no value is present for Material, not even an explicit nil
### GetStyle

`func (o *CreateListingByIdResponseItems) GetStyle() string`

GetStyle returns the Style field if non-nil, zero value otherwise.

### GetStyleOk

`func (o *CreateListingByIdResponseItems) GetStyleOk() (*string, bool)`

GetStyleOk returns a tuple with the Style field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStyle

`func (o *CreateListingByIdResponseItems) SetStyle(v string)`

SetStyle sets Style field to given value.

### HasStyle

`func (o *CreateListingByIdResponseItems) HasStyle() bool`

HasStyle returns a boolean if a field has been set.

### SetStyleNil

`func (o *CreateListingByIdResponseItems) SetStyleNil(b bool)`

 SetStyleNil sets the value for Style to be an explicit nil

### UnsetStyle
`func (o *CreateListingByIdResponseItems) UnsetStyle()`

UnsetStyle ensures that no value is present for Style, not even an explicit nil
### GetPattern

`func (o *CreateListingByIdResponseItems) GetPattern() string`

GetPattern returns the Pattern field if non-nil, zero value otherwise.

### GetPatternOk

`func (o *CreateListingByIdResponseItems) GetPatternOk() (*string, bool)`

GetPatternOk returns a tuple with the Pattern field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPattern

`func (o *CreateListingByIdResponseItems) SetPattern(v string)`

SetPattern sets Pattern field to given value.

### HasPattern

`func (o *CreateListingByIdResponseItems) HasPattern() bool`

HasPattern returns a boolean if a field has been set.

### SetPatternNil

`func (o *CreateListingByIdResponseItems) SetPatternNil(b bool)`

 SetPatternNil sets the value for Pattern to be an explicit nil

### UnsetPattern
`func (o *CreateListingByIdResponseItems) UnsetPattern()`

UnsetPattern ensures that no value is present for Pattern, not even an explicit nil
### GetDepartment

`func (o *CreateListingByIdResponseItems) GetDepartment() string`

GetDepartment returns the Department field if non-nil, zero value otherwise.

### GetDepartmentOk

`func (o *CreateListingByIdResponseItems) GetDepartmentOk() (*string, bool)`

GetDepartmentOk returns a tuple with the Department field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDepartment

`func (o *CreateListingByIdResponseItems) SetDepartment(v string)`

SetDepartment sets Department field to given value.

### HasDepartment

`func (o *CreateListingByIdResponseItems) HasDepartment() bool`

HasDepartment returns a boolean if a field has been set.

### SetDepartmentNil

`func (o *CreateListingByIdResponseItems) SetDepartmentNil(b bool)`

 SetDepartmentNil sets the value for Department to be an explicit nil

### UnsetDepartment
`func (o *CreateListingByIdResponseItems) UnsetDepartment()`

UnsetDepartment ensures that no value is present for Department, not even an explicit nil
### GetGender

`func (o *CreateListingByIdResponseItems) GetGender() string`

GetGender returns the Gender field if non-nil, zero value otherwise.

### GetGenderOk

`func (o *CreateListingByIdResponseItems) GetGenderOk() (*string, bool)`

GetGenderOk returns a tuple with the Gender field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGender

`func (o *CreateListingByIdResponseItems) SetGender(v string)`

SetGender sets Gender field to given value.

### HasGender

`func (o *CreateListingByIdResponseItems) HasGender() bool`

HasGender returns a boolean if a field has been set.

### SetGenderNil

`func (o *CreateListingByIdResponseItems) SetGenderNil(b bool)`

 SetGenderNil sets the value for Gender to be an explicit nil

### UnsetGender
`func (o *CreateListingByIdResponseItems) UnsetGender()`

UnsetGender ensures that no value is present for Gender, not even an explicit nil
### GetItemType

`func (o *CreateListingByIdResponseItems) GetItemType() string`

GetItemType returns the ItemType field if non-nil, zero value otherwise.

### GetItemTypeOk

`func (o *CreateListingByIdResponseItems) GetItemTypeOk() (*string, bool)`

GetItemTypeOk returns a tuple with the ItemType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemType

`func (o *CreateListingByIdResponseItems) SetItemType(v string)`

SetItemType sets ItemType field to given value.

### HasItemType

`func (o *CreateListingByIdResponseItems) HasItemType() bool`

HasItemType returns a boolean if a field has been set.

### SetItemTypeNil

`func (o *CreateListingByIdResponseItems) SetItemTypeNil(b bool)`

 SetItemTypeNil sets the value for ItemType to be an explicit nil

### UnsetItemType
`func (o *CreateListingByIdResponseItems) UnsetItemType()`

UnsetItemType ensures that no value is present for ItemType, not even an explicit nil
### GetSizeSystem

`func (o *CreateListingByIdResponseItems) GetSizeSystem() string`

GetSizeSystem returns the SizeSystem field if non-nil, zero value otherwise.

### GetSizeSystemOk

`func (o *CreateListingByIdResponseItems) GetSizeSystemOk() (*string, bool)`

GetSizeSystemOk returns a tuple with the SizeSystem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeSystem

`func (o *CreateListingByIdResponseItems) SetSizeSystem(v string)`

SetSizeSystem sets SizeSystem field to given value.

### HasSizeSystem

`func (o *CreateListingByIdResponseItems) HasSizeSystem() bool`

HasSizeSystem returns a boolean if a field has been set.

### SetSizeSystemNil

`func (o *CreateListingByIdResponseItems) SetSizeSystemNil(b bool)`

 SetSizeSystemNil sets the value for SizeSystem to be an explicit nil

### UnsetSizeSystem
`func (o *CreateListingByIdResponseItems) UnsetSizeSystem()`

UnsetSizeSystem ensures that no value is present for SizeSystem, not even an explicit nil
### GetColor

`func (o *CreateListingByIdResponseItems) GetColor() []string`

GetColor returns the Color field if non-nil, zero value otherwise.

### GetColorOk

`func (o *CreateListingByIdResponseItems) GetColorOk() (*[]string, bool)`

GetColorOk returns a tuple with the Color field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColor

`func (o *CreateListingByIdResponseItems) SetColor(v []string)`

SetColor sets Color field to given value.


### GetTags

`func (o *CreateListingByIdResponseItems) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CreateListingByIdResponseItems) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CreateListingByIdResponseItems) SetTags(v []string)`

SetTags sets Tags field to given value.


### GetSku

`func (o *CreateListingByIdResponseItems) GetSku() string`

GetSku returns the Sku field if non-nil, zero value otherwise.

### GetSkuOk

`func (o *CreateListingByIdResponseItems) GetSkuOk() (*string, bool)`

GetSkuOk returns a tuple with the Sku field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSku

`func (o *CreateListingByIdResponseItems) SetSku(v string)`

SetSku sets Sku field to given value.

### HasSku

`func (o *CreateListingByIdResponseItems) HasSku() bool`

HasSku returns a boolean if a field has been set.

### SetSkuNil

`func (o *CreateListingByIdResponseItems) SetSkuNil(b bool)`

 SetSkuNil sets the value for Sku to be an explicit nil

### UnsetSku
`func (o *CreateListingByIdResponseItems) UnsetSku()`

UnsetSku ensures that no value is present for Sku, not even an explicit nil
### GetQuantity

`func (o *CreateListingByIdResponseItems) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CreateListingByIdResponseItems) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CreateListingByIdResponseItems) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetQuantityAvailable

`func (o *CreateListingByIdResponseItems) GetQuantityAvailable() float32`

GetQuantityAvailable returns the QuantityAvailable field if non-nil, zero value otherwise.

### GetQuantityAvailableOk

`func (o *CreateListingByIdResponseItems) GetQuantityAvailableOk() (*float32, bool)`

GetQuantityAvailableOk returns a tuple with the QuantityAvailable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantityAvailable

`func (o *CreateListingByIdResponseItems) SetQuantityAvailable(v float32)`

SetQuantityAvailable sets QuantityAvailable field to given value.


### GetWeightLb

`func (o *CreateListingByIdResponseItems) GetWeightLb() string`

GetWeightLb returns the WeightLb field if non-nil, zero value otherwise.

### GetWeightLbOk

`func (o *CreateListingByIdResponseItems) GetWeightLbOk() (*string, bool)`

GetWeightLbOk returns a tuple with the WeightLb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightLb

`func (o *CreateListingByIdResponseItems) SetWeightLb(v string)`

SetWeightLb sets WeightLb field to given value.

### HasWeightLb

`func (o *CreateListingByIdResponseItems) HasWeightLb() bool`

HasWeightLb returns a boolean if a field has been set.

### SetWeightLbNil

`func (o *CreateListingByIdResponseItems) SetWeightLbNil(b bool)`

 SetWeightLbNil sets the value for WeightLb to be an explicit nil

### UnsetWeightLb
`func (o *CreateListingByIdResponseItems) UnsetWeightLb()`

UnsetWeightLb ensures that no value is present for WeightLb, not even an explicit nil
### GetWeightOz

`func (o *CreateListingByIdResponseItems) GetWeightOz() string`

GetWeightOz returns the WeightOz field if non-nil, zero value otherwise.

### GetWeightOzOk

`func (o *CreateListingByIdResponseItems) GetWeightOzOk() (*string, bool)`

GetWeightOzOk returns a tuple with the WeightOz field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeightOz

`func (o *CreateListingByIdResponseItems) SetWeightOz(v string)`

SetWeightOz sets WeightOz field to given value.

### HasWeightOz

`func (o *CreateListingByIdResponseItems) HasWeightOz() bool`

HasWeightOz returns a boolean if a field has been set.

### SetWeightOzNil

`func (o *CreateListingByIdResponseItems) SetWeightOzNil(b bool)`

 SetWeightOzNil sets the value for WeightOz to be an explicit nil

### UnsetWeightOz
`func (o *CreateListingByIdResponseItems) UnsetWeightOz()`

UnsetWeightOz ensures that no value is present for WeightOz, not even an explicit nil
### GetDimensionLIn

`func (o *CreateListingByIdResponseItems) GetDimensionLIn() string`

GetDimensionLIn returns the DimensionLIn field if non-nil, zero value otherwise.

### GetDimensionLInOk

`func (o *CreateListingByIdResponseItems) GetDimensionLInOk() (*string, bool)`

GetDimensionLInOk returns a tuple with the DimensionLIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionLIn

`func (o *CreateListingByIdResponseItems) SetDimensionLIn(v string)`

SetDimensionLIn sets DimensionLIn field to given value.

### HasDimensionLIn

`func (o *CreateListingByIdResponseItems) HasDimensionLIn() bool`

HasDimensionLIn returns a boolean if a field has been set.

### SetDimensionLInNil

`func (o *CreateListingByIdResponseItems) SetDimensionLInNil(b bool)`

 SetDimensionLInNil sets the value for DimensionLIn to be an explicit nil

### UnsetDimensionLIn
`func (o *CreateListingByIdResponseItems) UnsetDimensionLIn()`

UnsetDimensionLIn ensures that no value is present for DimensionLIn, not even an explicit nil
### GetDimensionWIn

`func (o *CreateListingByIdResponseItems) GetDimensionWIn() string`

GetDimensionWIn returns the DimensionWIn field if non-nil, zero value otherwise.

### GetDimensionWInOk

`func (o *CreateListingByIdResponseItems) GetDimensionWInOk() (*string, bool)`

GetDimensionWInOk returns a tuple with the DimensionWIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionWIn

`func (o *CreateListingByIdResponseItems) SetDimensionWIn(v string)`

SetDimensionWIn sets DimensionWIn field to given value.

### HasDimensionWIn

`func (o *CreateListingByIdResponseItems) HasDimensionWIn() bool`

HasDimensionWIn returns a boolean if a field has been set.

### SetDimensionWInNil

`func (o *CreateListingByIdResponseItems) SetDimensionWInNil(b bool)`

 SetDimensionWInNil sets the value for DimensionWIn to be an explicit nil

### UnsetDimensionWIn
`func (o *CreateListingByIdResponseItems) UnsetDimensionWIn()`

UnsetDimensionWIn ensures that no value is present for DimensionWIn, not even an explicit nil
### GetDimensionHIn

`func (o *CreateListingByIdResponseItems) GetDimensionHIn() string`

GetDimensionHIn returns the DimensionHIn field if non-nil, zero value otherwise.

### GetDimensionHInOk

`func (o *CreateListingByIdResponseItems) GetDimensionHInOk() (*string, bool)`

GetDimensionHInOk returns a tuple with the DimensionHIn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDimensionHIn

`func (o *CreateListingByIdResponseItems) SetDimensionHIn(v string)`

SetDimensionHIn sets DimensionHIn field to given value.

### HasDimensionHIn

`func (o *CreateListingByIdResponseItems) HasDimensionHIn() bool`

HasDimensionHIn returns a boolean if a field has been set.

### SetDimensionHInNil

`func (o *CreateListingByIdResponseItems) SetDimensionHInNil(b bool)`

 SetDimensionHInNil sets the value for DimensionHIn to be an explicit nil

### UnsetDimensionHIn
`func (o *CreateListingByIdResponseItems) UnsetDimensionHIn()`

UnsetDimensionHIn ensures that no value is present for DimensionHIn, not even an explicit nil
### GetPublishAt

`func (o *CreateListingByIdResponseItems) GetPublishAt() time.Time`

GetPublishAt returns the PublishAt field if non-nil, zero value otherwise.

### GetPublishAtOk

`func (o *CreateListingByIdResponseItems) GetPublishAtOk() (*time.Time, bool)`

GetPublishAtOk returns a tuple with the PublishAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublishAt

`func (o *CreateListingByIdResponseItems) SetPublishAt(v time.Time)`

SetPublishAt sets PublishAt field to given value.

### HasPublishAt

`func (o *CreateListingByIdResponseItems) HasPublishAt() bool`

HasPublishAt returns a boolean if a field has been set.

### SetPublishAtNil

`func (o *CreateListingByIdResponseItems) SetPublishAtNil(b bool)`

 SetPublishAtNil sets the value for PublishAt to be an explicit nil

### UnsetPublishAt
`func (o *CreateListingByIdResponseItems) UnsetPublishAt()`

UnsetPublishAt ensures that no value is present for PublishAt, not even an explicit nil
### GetScheduledPlatforms

`func (o *CreateListingByIdResponseItems) GetScheduledPlatforms() []string`

GetScheduledPlatforms returns the ScheduledPlatforms field if non-nil, zero value otherwise.

### GetScheduledPlatformsOk

`func (o *CreateListingByIdResponseItems) GetScheduledPlatformsOk() (*[]string, bool)`

GetScheduledPlatformsOk returns a tuple with the ScheduledPlatforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPlatforms

`func (o *CreateListingByIdResponseItems) SetScheduledPlatforms(v []string)`

SetScheduledPlatforms sets ScheduledPlatforms field to given value.

### HasScheduledPlatforms

`func (o *CreateListingByIdResponseItems) HasScheduledPlatforms() bool`

HasScheduledPlatforms returns a boolean if a field has been set.

### SetScheduledPlatformsNil

`func (o *CreateListingByIdResponseItems) SetScheduledPlatformsNil(b bool)`

 SetScheduledPlatformsNil sets the value for ScheduledPlatforms to be an explicit nil

### UnsetScheduledPlatforms
`func (o *CreateListingByIdResponseItems) UnsetScheduledPlatforms()`

UnsetScheduledPlatforms ensures that no value is present for ScheduledPlatforms, not even an explicit nil
### GetParentListingId

`func (o *CreateListingByIdResponseItems) GetParentListingId() string`

GetParentListingId returns the ParentListingId field if non-nil, zero value otherwise.

### GetParentListingIdOk

`func (o *CreateListingByIdResponseItems) GetParentListingIdOk() (*string, bool)`

GetParentListingIdOk returns a tuple with the ParentListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentListingId

`func (o *CreateListingByIdResponseItems) SetParentListingId(v string)`

SetParentListingId sets ParentListingId field to given value.

### HasParentListingId

`func (o *CreateListingByIdResponseItems) HasParentListingId() bool`

HasParentListingId returns a boolean if a field has been set.

### SetParentListingIdNil

`func (o *CreateListingByIdResponseItems) SetParentListingIdNil(b bool)`

 SetParentListingIdNil sets the value for ParentListingId to be an explicit nil

### UnsetParentListingId
`func (o *CreateListingByIdResponseItems) UnsetParentListingId()`

UnsetParentListingId ensures that no value is present for ParentListingId, not even an explicit nil
### GetIsBundle

`func (o *CreateListingByIdResponseItems) GetIsBundle() bool`

GetIsBundle returns the IsBundle field if non-nil, zero value otherwise.

### GetIsBundleOk

`func (o *CreateListingByIdResponseItems) GetIsBundleOk() (*bool, bool)`

GetIsBundleOk returns a tuple with the IsBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsBundle

`func (o *CreateListingByIdResponseItems) SetIsBundle(v bool)`

SetIsBundle sets IsBundle field to given value.


### GetAutomationAssignedRuleIds

`func (o *CreateListingByIdResponseItems) GetAutomationAssignedRuleIds() []string`

GetAutomationAssignedRuleIds returns the AutomationAssignedRuleIds field if non-nil, zero value otherwise.

### GetAutomationAssignedRuleIdsOk

`func (o *CreateListingByIdResponseItems) GetAutomationAssignedRuleIdsOk() (*[]string, bool)`

GetAutomationAssignedRuleIdsOk returns a tuple with the AutomationAssignedRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationAssignedRuleIds

`func (o *CreateListingByIdResponseItems) SetAutomationAssignedRuleIds(v []string)`

SetAutomationAssignedRuleIds sets AutomationAssignedRuleIds field to given value.


### GetAutomationBlockedRuleIds

`func (o *CreateListingByIdResponseItems) GetAutomationBlockedRuleIds() []string`

GetAutomationBlockedRuleIds returns the AutomationBlockedRuleIds field if non-nil, zero value otherwise.

### GetAutomationBlockedRuleIdsOk

`func (o *CreateListingByIdResponseItems) GetAutomationBlockedRuleIdsOk() (*[]string, bool)`

GetAutomationBlockedRuleIdsOk returns a tuple with the AutomationBlockedRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationBlockedRuleIds

`func (o *CreateListingByIdResponseItems) SetAutomationBlockedRuleIds(v []string)`

SetAutomationBlockedRuleIds sets AutomationBlockedRuleIds field to given value.


### GetAutomationAssignedChainIds

`func (o *CreateListingByIdResponseItems) GetAutomationAssignedChainIds() []string`

GetAutomationAssignedChainIds returns the AutomationAssignedChainIds field if non-nil, zero value otherwise.

### GetAutomationAssignedChainIdsOk

`func (o *CreateListingByIdResponseItems) GetAutomationAssignedChainIdsOk() (*[]string, bool)`

GetAutomationAssignedChainIdsOk returns a tuple with the AutomationAssignedChainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationAssignedChainIds

`func (o *CreateListingByIdResponseItems) SetAutomationAssignedChainIds(v []string)`

SetAutomationAssignedChainIds sets AutomationAssignedChainIds field to given value.


### GetAutomationBlockedChainIds

`func (o *CreateListingByIdResponseItems) GetAutomationBlockedChainIds() []string`

GetAutomationBlockedChainIds returns the AutomationBlockedChainIds field if non-nil, zero value otherwise.

### GetAutomationBlockedChainIdsOk

`func (o *CreateListingByIdResponseItems) GetAutomationBlockedChainIdsOk() (*[]string, bool)`

GetAutomationBlockedChainIdsOk returns a tuple with the AutomationBlockedChainIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomationBlockedChainIds

`func (o *CreateListingByIdResponseItems) SetAutomationBlockedChainIds(v []string)`

SetAutomationBlockedChainIds sets AutomationBlockedChainIds field to given value.


### GetHsCode

`func (o *CreateListingByIdResponseItems) GetHsCode() string`

GetHsCode returns the HsCode field if non-nil, zero value otherwise.

### GetHsCodeOk

`func (o *CreateListingByIdResponseItems) GetHsCodeOk() (*string, bool)`

GetHsCodeOk returns a tuple with the HsCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHsCode

`func (o *CreateListingByIdResponseItems) SetHsCode(v string)`

SetHsCode sets HsCode field to given value.

### HasHsCode

`func (o *CreateListingByIdResponseItems) HasHsCode() bool`

HasHsCode returns a boolean if a field has been set.

### SetHsCodeNil

`func (o *CreateListingByIdResponseItems) SetHsCodeNil(b bool)`

 SetHsCodeNil sets the value for HsCode to be an explicit nil

### UnsetHsCode
`func (o *CreateListingByIdResponseItems) UnsetHsCode()`

UnsetHsCode ensures that no value is present for HsCode, not even an explicit nil
### GetCountryOfOrigin

`func (o *CreateListingByIdResponseItems) GetCountryOfOrigin() string`

GetCountryOfOrigin returns the CountryOfOrigin field if non-nil, zero value otherwise.

### GetCountryOfOriginOk

`func (o *CreateListingByIdResponseItems) GetCountryOfOriginOk() (*string, bool)`

GetCountryOfOriginOk returns a tuple with the CountryOfOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountryOfOrigin

`func (o *CreateListingByIdResponseItems) SetCountryOfOrigin(v string)`

SetCountryOfOrigin sets CountryOfOrigin field to given value.

### HasCountryOfOrigin

`func (o *CreateListingByIdResponseItems) HasCountryOfOrigin() bool`

HasCountryOfOrigin returns a boolean if a field has been set.

### SetCountryOfOriginNil

`func (o *CreateListingByIdResponseItems) SetCountryOfOriginNil(b bool)`

 SetCountryOfOriginNil sets the value for CountryOfOrigin to be an explicit nil

### UnsetCountryOfOrigin
`func (o *CreateListingByIdResponseItems) UnsetCountryOfOrigin()`

UnsetCountryOfOrigin ensures that no value is present for CountryOfOrigin, not even an explicit nil
### GetPriceFloorCents

`func (o *CreateListingByIdResponseItems) GetPriceFloorCents() float32`

GetPriceFloorCents returns the PriceFloorCents field if non-nil, zero value otherwise.

### GetPriceFloorCentsOk

`func (o *CreateListingByIdResponseItems) GetPriceFloorCentsOk() (*float32, bool)`

GetPriceFloorCentsOk returns a tuple with the PriceFloorCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceFloorCents

`func (o *CreateListingByIdResponseItems) SetPriceFloorCents(v float32)`

SetPriceFloorCents sets PriceFloorCents field to given value.

### HasPriceFloorCents

`func (o *CreateListingByIdResponseItems) HasPriceFloorCents() bool`

HasPriceFloorCents returns a boolean if a field has been set.

### SetPriceFloorCentsNil

`func (o *CreateListingByIdResponseItems) SetPriceFloorCentsNil(b bool)`

 SetPriceFloorCentsNil sets the value for PriceFloorCents to be an explicit nil

### UnsetPriceFloorCents
`func (o *CreateListingByIdResponseItems) UnsetPriceFloorCents()`

UnsetPriceFloorCents ensures that no value is present for PriceFloorCents, not even an explicit nil
### GetFloorIsNet

`func (o *CreateListingByIdResponseItems) GetFloorIsNet() bool`

GetFloorIsNet returns the FloorIsNet field if non-nil, zero value otherwise.

### GetFloorIsNetOk

`func (o *CreateListingByIdResponseItems) GetFloorIsNetOk() (*bool, bool)`

GetFloorIsNetOk returns a tuple with the FloorIsNet field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFloorIsNet

`func (o *CreateListingByIdResponseItems) SetFloorIsNet(v bool)`

SetFloorIsNet sets FloorIsNet field to given value.


### GetDelistedAt

`func (o *CreateListingByIdResponseItems) GetDelistedAt() time.Time`

GetDelistedAt returns the DelistedAt field if non-nil, zero value otherwise.

### GetDelistedAtOk

`func (o *CreateListingByIdResponseItems) GetDelistedAtOk() (*time.Time, bool)`

GetDelistedAtOk returns a tuple with the DelistedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDelistedAt

`func (o *CreateListingByIdResponseItems) SetDelistedAt(v time.Time)`

SetDelistedAt sets DelistedAt field to given value.

### HasDelistedAt

`func (o *CreateListingByIdResponseItems) HasDelistedAt() bool`

HasDelistedAt returns a boolean if a field has been set.

### SetDelistedAtNil

`func (o *CreateListingByIdResponseItems) SetDelistedAtNil(b bool)`

 SetDelistedAtNil sets the value for DelistedAt to be an explicit nil

### UnsetDelistedAt
`func (o *CreateListingByIdResponseItems) UnsetDelistedAt()`

UnsetDelistedAt ensures that no value is present for DelistedAt, not even an explicit nil
### GetSoldAt

`func (o *CreateListingByIdResponseItems) GetSoldAt() time.Time`

GetSoldAt returns the SoldAt field if non-nil, zero value otherwise.

### GetSoldAtOk

`func (o *CreateListingByIdResponseItems) GetSoldAtOk() (*time.Time, bool)`

GetSoldAtOk returns a tuple with the SoldAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoldAt

`func (o *CreateListingByIdResponseItems) SetSoldAt(v time.Time)`

SetSoldAt sets SoldAt field to given value.

### HasSoldAt

`func (o *CreateListingByIdResponseItems) HasSoldAt() bool`

HasSoldAt returns a boolean if a field has been set.

### SetSoldAtNil

`func (o *CreateListingByIdResponseItems) SetSoldAtNil(b bool)`

 SetSoldAtNil sets the value for SoldAt to be an explicit nil

### UnsetSoldAt
`func (o *CreateListingByIdResponseItems) UnsetSoldAt()`

UnsetSoldAt ensures that no value is present for SoldAt, not even an explicit nil
### GetSource

`func (o *CreateListingByIdResponseItems) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *CreateListingByIdResponseItems) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *CreateListingByIdResponseItems) SetSource(v string)`

SetSource sets Source field to given value.


### GetDuplicateOfListingId

`func (o *CreateListingByIdResponseItems) GetDuplicateOfListingId() string`

GetDuplicateOfListingId returns the DuplicateOfListingId field if non-nil, zero value otherwise.

### GetDuplicateOfListingIdOk

`func (o *CreateListingByIdResponseItems) GetDuplicateOfListingIdOk() (*string, bool)`

GetDuplicateOfListingIdOk returns a tuple with the DuplicateOfListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuplicateOfListingId

`func (o *CreateListingByIdResponseItems) SetDuplicateOfListingId(v string)`

SetDuplicateOfListingId sets DuplicateOfListingId field to given value.

### HasDuplicateOfListingId

`func (o *CreateListingByIdResponseItems) HasDuplicateOfListingId() bool`

HasDuplicateOfListingId returns a boolean if a field has been set.

### SetDuplicateOfListingIdNil

`func (o *CreateListingByIdResponseItems) SetDuplicateOfListingIdNil(b bool)`

 SetDuplicateOfListingIdNil sets the value for DuplicateOfListingId to be an explicit nil

### UnsetDuplicateOfListingId
`func (o *CreateListingByIdResponseItems) UnsetDuplicateOfListingId()`

UnsetDuplicateOfListingId ensures that no value is present for DuplicateOfListingId, not even an explicit nil
### GetClientDraftId

`func (o *CreateListingByIdResponseItems) GetClientDraftId() string`

GetClientDraftId returns the ClientDraftId field if non-nil, zero value otherwise.

### GetClientDraftIdOk

`func (o *CreateListingByIdResponseItems) GetClientDraftIdOk() (*string, bool)`

GetClientDraftIdOk returns a tuple with the ClientDraftId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientDraftId

`func (o *CreateListingByIdResponseItems) SetClientDraftId(v string)`

SetClientDraftId sets ClientDraftId field to given value.

### HasClientDraftId

`func (o *CreateListingByIdResponseItems) HasClientDraftId() bool`

HasClientDraftId returns a boolean if a field has been set.

### SetClientDraftIdNil

`func (o *CreateListingByIdResponseItems) SetClientDraftIdNil(b bool)`

 SetClientDraftIdNil sets the value for ClientDraftId to be an explicit nil

### UnsetClientDraftId
`func (o *CreateListingByIdResponseItems) UnsetClientDraftId()`

UnsetClientDraftId ensures that no value is present for ClientDraftId, not even an explicit nil
### GetCreatedAt

`func (o *CreateListingByIdResponseItems) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *CreateListingByIdResponseItems) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *CreateListingByIdResponseItems) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *CreateListingByIdResponseItems) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *CreateListingByIdResponseItems) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *CreateListingByIdResponseItems) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


