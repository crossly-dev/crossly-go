# ListTaxonomySuggestItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExternalId** | **string** |  | 
**Display** | **string** |  | 
**Level** | **float32** |  | 
**Catalog** | [**ListTaxonomySuggestItemCatalog**](ListTaxonomySuggestItemCatalog.md) |  | 
**Breadcrumb** | **string** | \&quot;Kids &gt; Toys &gt; Building Sets &amp; Blocks\&quot; for the picker list. | 
**Displays** | [**ListTaxonomySuggestItemDisplays**](ListTaxonomySuggestItemDisplays.md) |  | 

## Methods

### NewListTaxonomySuggestItem

`func NewListTaxonomySuggestItem(externalId string, display string, level float32, catalog ListTaxonomySuggestItemCatalog, breadcrumb string, displays ListTaxonomySuggestItemDisplays, ) *ListTaxonomySuggestItem`

NewListTaxonomySuggestItem instantiates a new ListTaxonomySuggestItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTaxonomySuggestItemWithDefaults

`func NewListTaxonomySuggestItemWithDefaults() *ListTaxonomySuggestItem`

NewListTaxonomySuggestItemWithDefaults instantiates a new ListTaxonomySuggestItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExternalId

`func (o *ListTaxonomySuggestItem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListTaxonomySuggestItem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListTaxonomySuggestItem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.


### GetDisplay

`func (o *ListTaxonomySuggestItem) GetDisplay() string`

GetDisplay returns the Display field if non-nil, zero value otherwise.

### GetDisplayOk

`func (o *ListTaxonomySuggestItem) GetDisplayOk() (*string, bool)`

GetDisplayOk returns a tuple with the Display field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplay

`func (o *ListTaxonomySuggestItem) SetDisplay(v string)`

SetDisplay sets Display field to given value.


### GetLevel

`func (o *ListTaxonomySuggestItem) GetLevel() float32`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *ListTaxonomySuggestItem) GetLevelOk() (*float32, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *ListTaxonomySuggestItem) SetLevel(v float32)`

SetLevel sets Level field to given value.


### GetCatalog

`func (o *ListTaxonomySuggestItem) GetCatalog() ListTaxonomySuggestItemCatalog`

GetCatalog returns the Catalog field if non-nil, zero value otherwise.

### GetCatalogOk

`func (o *ListTaxonomySuggestItem) GetCatalogOk() (*ListTaxonomySuggestItemCatalog, bool)`

GetCatalogOk returns a tuple with the Catalog field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalog

`func (o *ListTaxonomySuggestItem) SetCatalog(v ListTaxonomySuggestItemCatalog)`

SetCatalog sets Catalog field to given value.


### GetBreadcrumb

`func (o *ListTaxonomySuggestItem) GetBreadcrumb() string`

GetBreadcrumb returns the Breadcrumb field if non-nil, zero value otherwise.

### GetBreadcrumbOk

`func (o *ListTaxonomySuggestItem) GetBreadcrumbOk() (*string, bool)`

GetBreadcrumbOk returns a tuple with the Breadcrumb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBreadcrumb

`func (o *ListTaxonomySuggestItem) SetBreadcrumb(v string)`

SetBreadcrumb sets Breadcrumb field to given value.


### GetDisplays

`func (o *ListTaxonomySuggestItem) GetDisplays() ListTaxonomySuggestItemDisplays`

GetDisplays returns the Displays field if non-nil, zero value otherwise.

### GetDisplaysOk

`func (o *ListTaxonomySuggestItem) GetDisplaysOk() (*ListTaxonomySuggestItemDisplays, bool)`

GetDisplaysOk returns a tuple with the Displays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplays

`func (o *ListTaxonomySuggestItem) SetDisplays(v ListTaxonomySuggestItemDisplays)`

SetDisplays sets Displays field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


