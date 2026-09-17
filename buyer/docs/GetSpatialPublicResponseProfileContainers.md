# GetSpatialPublicResponseProfileContainers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Key** | **string** | Stable key. Persisted on every placement — never renamed. | 
**Label** | **string** | What the seller calls it. \&quot;Page 3\&quot;, \&quot;Case\&quot;, \&quot;Long box\&quot;. | 
**Model** | **string** | Which hand-authored mesh renders this container. A small closed set — the  art budget lives here and nowhere else. | 
**Size** | **map[string]interface{}** | Outer bounds, for packing containers into a room. | 
**Slots** | **map[string]interface{}** |  | 

## Methods

### NewGetSpatialPublicResponseProfileContainers

`func NewGetSpatialPublicResponseProfileContainers(key string, label string, model string, size map[string]interface{}, slots map[string]interface{}, ) *GetSpatialPublicResponseProfileContainers`

NewGetSpatialPublicResponseProfileContainers instantiates a new GetSpatialPublicResponseProfileContainers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSpatialPublicResponseProfileContainersWithDefaults

`func NewGetSpatialPublicResponseProfileContainersWithDefaults() *GetSpatialPublicResponseProfileContainers`

NewGetSpatialPublicResponseProfileContainersWithDefaults instantiates a new GetSpatialPublicResponseProfileContainers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKey

`func (o *GetSpatialPublicResponseProfileContainers) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *GetSpatialPublicResponseProfileContainers) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *GetSpatialPublicResponseProfileContainers) SetKey(v string)`

SetKey sets Key field to given value.


### GetLabel

`func (o *GetSpatialPublicResponseProfileContainers) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *GetSpatialPublicResponseProfileContainers) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *GetSpatialPublicResponseProfileContainers) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetModel

`func (o *GetSpatialPublicResponseProfileContainers) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *GetSpatialPublicResponseProfileContainers) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *GetSpatialPublicResponseProfileContainers) SetModel(v string)`

SetModel sets Model field to given value.


### GetSize

`func (o *GetSpatialPublicResponseProfileContainers) GetSize() map[string]interface{}`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetSpatialPublicResponseProfileContainers) GetSizeOk() (*map[string]interface{}, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetSpatialPublicResponseProfileContainers) SetSize(v map[string]interface{})`

SetSize sets Size field to given value.


### GetSlots

`func (o *GetSpatialPublicResponseProfileContainers) GetSlots() map[string]interface{}`

GetSlots returns the Slots field if non-nil, zero value otherwise.

### GetSlotsOk

`func (o *GetSpatialPublicResponseProfileContainers) GetSlotsOk() (*map[string]interface{}, bool)`

GetSlotsOk returns a tuple with the Slots field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSlots

`func (o *GetSpatialPublicResponseProfileContainers) SetSlots(v map[string]interface{})`

SetSlots sets Slots field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


