# GetConnectionEmailResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GetConnectionEmailResponseData**](GetConnectionEmailResponseData.md) |  | 
**Imap** | [**[]GetConnectionEmailResponseData**](GetConnectionEmailResponseData.md) |  | 
**Oauth** | [**[]GetConnectionEmailResponseOauth**](GetConnectionEmailResponseOauth.md) |  | 

## Methods

### NewGetConnectionEmailResponse

`func NewGetConnectionEmailResponse(data []GetConnectionEmailResponseData, imap []GetConnectionEmailResponseData, oauth []GetConnectionEmailResponseOauth, ) *GetConnectionEmailResponse`

NewGetConnectionEmailResponse instantiates a new GetConnectionEmailResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionEmailResponseWithDefaults

`func NewGetConnectionEmailResponseWithDefaults() *GetConnectionEmailResponse`

NewGetConnectionEmailResponseWithDefaults instantiates a new GetConnectionEmailResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GetConnectionEmailResponse) GetData() []GetConnectionEmailResponseData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GetConnectionEmailResponse) GetDataOk() (*[]GetConnectionEmailResponseData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GetConnectionEmailResponse) SetData(v []GetConnectionEmailResponseData)`

SetData sets Data field to given value.


### GetImap

`func (o *GetConnectionEmailResponse) GetImap() []GetConnectionEmailResponseData`

GetImap returns the Imap field if non-nil, zero value otherwise.

### GetImapOk

`func (o *GetConnectionEmailResponse) GetImapOk() (*[]GetConnectionEmailResponseData, bool)`

GetImapOk returns a tuple with the Imap field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImap

`func (o *GetConnectionEmailResponse) SetImap(v []GetConnectionEmailResponseData)`

SetImap sets Imap field to given value.


### GetOauth

`func (o *GetConnectionEmailResponse) GetOauth() []GetConnectionEmailResponseOauth`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *GetConnectionEmailResponse) GetOauthOk() (*[]GetConnectionEmailResponseOauth, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *GetConnectionEmailResponse) SetOauth(v []GetConnectionEmailResponseOauth)`

SetOauth sets Oauth field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


