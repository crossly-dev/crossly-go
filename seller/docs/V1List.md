# V1List

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **[]interface{}** | The rows. | 
**Pagination** | Pointer to [**Pagination**](Pagination.md) |  | [optional] 
**Meta** | Pointer to **map[string]map[string]interface{}** | Endpoint-specific extras that are NOT pagination — e.g. &#x60;pendingCents&#x60;, &#x60;basis&#x60;. Omitted when empty. | [optional] 

## Methods

### NewV1List

`func NewV1List(data []interface{}, ) *V1List`

NewV1List instantiates a new V1List object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ListWithDefaults

`func NewV1ListWithDefaults() *V1List`

NewV1ListWithDefaults instantiates a new V1List object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *V1List) GetData() []interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *V1List) GetDataOk() (*[]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *V1List) SetData(v []interface{})`

SetData sets Data field to given value.


### GetPagination

`func (o *V1List) GetPagination() Pagination`

GetPagination returns the Pagination field if non-nil, zero value otherwise.

### GetPaginationOk

`func (o *V1List) GetPaginationOk() (*Pagination, bool)`

GetPaginationOk returns a tuple with the Pagination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPagination

`func (o *V1List) SetPagination(v Pagination)`

SetPagination sets Pagination field to given value.

### HasPagination

`func (o *V1List) HasPagination() bool`

HasPagination returns a boolean if a field has been set.

### GetMeta

`func (o *V1List) GetMeta() map[string]map[string]interface{}`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *V1List) GetMetaOk() (*map[string]map[string]interface{}, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *V1List) SetMeta(v map[string]map[string]interface{})`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *V1List) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


