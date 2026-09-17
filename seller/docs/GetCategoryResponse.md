# GetCategoryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Main** | [**GetCategoryResponseMain**](GetCategoryResponseMain.md) |  | 
**Sub** | **map[string]interface{}** |  | 

## Methods

### NewGetCategoryResponse

`func NewGetCategoryResponse(main GetCategoryResponseMain, sub map[string]interface{}, ) *GetCategoryResponse`

NewGetCategoryResponse instantiates a new GetCategoryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCategoryResponseWithDefaults

`func NewGetCategoryResponseWithDefaults() *GetCategoryResponse`

NewGetCategoryResponseWithDefaults instantiates a new GetCategoryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMain

`func (o *GetCategoryResponse) GetMain() GetCategoryResponseMain`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *GetCategoryResponse) GetMainOk() (*GetCategoryResponseMain, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *GetCategoryResponse) SetMain(v GetCategoryResponseMain)`

SetMain sets Main field to given value.


### GetSub

`func (o *GetCategoryResponse) GetSub() map[string]interface{}`

GetSub returns the Sub field if non-nil, zero value otherwise.

### GetSubOk

`func (o *GetCategoryResponse) GetSubOk() (*map[string]interface{}, bool)`

GetSubOk returns a tuple with the Sub field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSub

`func (o *GetCategoryResponse) SetSub(v map[string]interface{})`

SetSub sets Sub field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


