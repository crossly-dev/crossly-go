# ListMagicRecentItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**ImageUrl** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**TopMatchTitle** | **string** |  | 
**MatchCount** | **float32** |  | 

## Methods

### NewListMagicRecentItem

`func NewListMagicRecentItem(id string, imageUrl string, createdAt time.Time, topMatchTitle string, matchCount float32, ) *ListMagicRecentItem`

NewListMagicRecentItem instantiates a new ListMagicRecentItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMagicRecentItemWithDefaults

`func NewListMagicRecentItemWithDefaults() *ListMagicRecentItem`

NewListMagicRecentItemWithDefaults instantiates a new ListMagicRecentItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListMagicRecentItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListMagicRecentItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListMagicRecentItem) SetId(v string)`

SetId sets Id field to given value.


### GetImageUrl

`func (o *ListMagicRecentItem) GetImageUrl() string`

GetImageUrl returns the ImageUrl field if non-nil, zero value otherwise.

### GetImageUrlOk

`func (o *ListMagicRecentItem) GetImageUrlOk() (*string, bool)`

GetImageUrlOk returns a tuple with the ImageUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageUrl

`func (o *ListMagicRecentItem) SetImageUrl(v string)`

SetImageUrl sets ImageUrl field to given value.


### GetCreatedAt

`func (o *ListMagicRecentItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListMagicRecentItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListMagicRecentItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetTopMatchTitle

`func (o *ListMagicRecentItem) GetTopMatchTitle() string`

GetTopMatchTitle returns the TopMatchTitle field if non-nil, zero value otherwise.

### GetTopMatchTitleOk

`func (o *ListMagicRecentItem) GetTopMatchTitleOk() (*string, bool)`

GetTopMatchTitleOk returns a tuple with the TopMatchTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopMatchTitle

`func (o *ListMagicRecentItem) SetTopMatchTitle(v string)`

SetTopMatchTitle sets TopMatchTitle field to given value.


### GetMatchCount

`func (o *ListMagicRecentItem) GetMatchCount() float32`

GetMatchCount returns the MatchCount field if non-nil, zero value otherwise.

### GetMatchCountOk

`func (o *ListMagicRecentItem) GetMatchCountOk() (*float32, bool)`

GetMatchCountOk returns a tuple with the MatchCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchCount

`func (o *ListMagicRecentItem) SetMatchCount(v float32)`

SetMatchCount sets MatchCount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


