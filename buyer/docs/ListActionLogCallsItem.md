# ListActionLogCallsItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UserId** | Pointer to **NullableString** |  | [optional] 
**Platform** | **string** |  | 
**Action** | **string** |  | 
**Step** | Pointer to **NullableFloat32** |  | [optional] 
**Url** | **string** |  | 
**Method** | **string** |  | 
**StatusCode** | Pointer to **NullableFloat32** |  | [optional] 
**LatencyMs** | Pointer to **NullableFloat32** |  | [optional] 
**ErrorClass** | Pointer to **NullableString** |  | [optional] 
**ErrorMessage** | Pointer to **NullableString** |  | [optional] 
**RequestBodyExcerpt** | Pointer to **NullableString** |  | [optional] 
**ResponseBodyExcerpt** | Pointer to **NullableString** |  | [optional] 
**RecipeId** | Pointer to **NullableString** |  | [optional] 
**DocIdOrHash** | Pointer to **NullableString** |  | [optional] 
**ProxyId** | Pointer to **NullableString** |  | [optional] 
**CorrelationId** | **string** |  | 
**ActionEventId** | Pointer to **NullableString** |  | [optional] 
**Ok** | **bool** |  | 

## Methods

### NewListActionLogCallsItem

`func NewListActionLogCallsItem(id string, createdAt time.Time, platform string, action string, url string, method string, correlationId string, ok bool, ) *ListActionLogCallsItem`

NewListActionLogCallsItem instantiates a new ListActionLogCallsItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListActionLogCallsItemWithDefaults

`func NewListActionLogCallsItemWithDefaults() *ListActionLogCallsItem`

NewListActionLogCallsItemWithDefaults instantiates a new ListActionLogCallsItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListActionLogCallsItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListActionLogCallsItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListActionLogCallsItem) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListActionLogCallsItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListActionLogCallsItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListActionLogCallsItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *ListActionLogCallsItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListActionLogCallsItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListActionLogCallsItem) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ListActionLogCallsItem) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### SetUserIdNil

`func (o *ListActionLogCallsItem) SetUserIdNil(b bool)`

 SetUserIdNil sets the value for UserId to be an explicit nil

### UnsetUserId
`func (o *ListActionLogCallsItem) UnsetUserId()`

UnsetUserId ensures that no value is present for UserId, not even an explicit nil
### GetPlatform

`func (o *ListActionLogCallsItem) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListActionLogCallsItem) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListActionLogCallsItem) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetAction

`func (o *ListActionLogCallsItem) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *ListActionLogCallsItem) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *ListActionLogCallsItem) SetAction(v string)`

SetAction sets Action field to given value.


### GetStep

`func (o *ListActionLogCallsItem) GetStep() float32`

GetStep returns the Step field if non-nil, zero value otherwise.

### GetStepOk

`func (o *ListActionLogCallsItem) GetStepOk() (*float32, bool)`

GetStepOk returns a tuple with the Step field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStep

`func (o *ListActionLogCallsItem) SetStep(v float32)`

SetStep sets Step field to given value.

### HasStep

`func (o *ListActionLogCallsItem) HasStep() bool`

HasStep returns a boolean if a field has been set.

### SetStepNil

`func (o *ListActionLogCallsItem) SetStepNil(b bool)`

 SetStepNil sets the value for Step to be an explicit nil

### UnsetStep
`func (o *ListActionLogCallsItem) UnsetStep()`

UnsetStep ensures that no value is present for Step, not even an explicit nil
### GetUrl

`func (o *ListActionLogCallsItem) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ListActionLogCallsItem) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ListActionLogCallsItem) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetMethod

`func (o *ListActionLogCallsItem) GetMethod() string`

GetMethod returns the Method field if non-nil, zero value otherwise.

### GetMethodOk

`func (o *ListActionLogCallsItem) GetMethodOk() (*string, bool)`

GetMethodOk returns a tuple with the Method field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethod

`func (o *ListActionLogCallsItem) SetMethod(v string)`

SetMethod sets Method field to given value.


### GetStatusCode

`func (o *ListActionLogCallsItem) GetStatusCode() float32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ListActionLogCallsItem) GetStatusCodeOk() (*float32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ListActionLogCallsItem) SetStatusCode(v float32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ListActionLogCallsItem) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### SetStatusCodeNil

`func (o *ListActionLogCallsItem) SetStatusCodeNil(b bool)`

 SetStatusCodeNil sets the value for StatusCode to be an explicit nil

### UnsetStatusCode
`func (o *ListActionLogCallsItem) UnsetStatusCode()`

UnsetStatusCode ensures that no value is present for StatusCode, not even an explicit nil
### GetLatencyMs

`func (o *ListActionLogCallsItem) GetLatencyMs() float32`

GetLatencyMs returns the LatencyMs field if non-nil, zero value otherwise.

### GetLatencyMsOk

`func (o *ListActionLogCallsItem) GetLatencyMsOk() (*float32, bool)`

GetLatencyMsOk returns a tuple with the LatencyMs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatencyMs

`func (o *ListActionLogCallsItem) SetLatencyMs(v float32)`

SetLatencyMs sets LatencyMs field to given value.

### HasLatencyMs

`func (o *ListActionLogCallsItem) HasLatencyMs() bool`

HasLatencyMs returns a boolean if a field has been set.

### SetLatencyMsNil

`func (o *ListActionLogCallsItem) SetLatencyMsNil(b bool)`

 SetLatencyMsNil sets the value for LatencyMs to be an explicit nil

### UnsetLatencyMs
`func (o *ListActionLogCallsItem) UnsetLatencyMs()`

UnsetLatencyMs ensures that no value is present for LatencyMs, not even an explicit nil
### GetErrorClass

`func (o *ListActionLogCallsItem) GetErrorClass() string`

GetErrorClass returns the ErrorClass field if non-nil, zero value otherwise.

### GetErrorClassOk

`func (o *ListActionLogCallsItem) GetErrorClassOk() (*string, bool)`

GetErrorClassOk returns a tuple with the ErrorClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorClass

`func (o *ListActionLogCallsItem) SetErrorClass(v string)`

SetErrorClass sets ErrorClass field to given value.

### HasErrorClass

`func (o *ListActionLogCallsItem) HasErrorClass() bool`

HasErrorClass returns a boolean if a field has been set.

### SetErrorClassNil

`func (o *ListActionLogCallsItem) SetErrorClassNil(b bool)`

 SetErrorClassNil sets the value for ErrorClass to be an explicit nil

### UnsetErrorClass
`func (o *ListActionLogCallsItem) UnsetErrorClass()`

UnsetErrorClass ensures that no value is present for ErrorClass, not even an explicit nil
### GetErrorMessage

`func (o *ListActionLogCallsItem) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *ListActionLogCallsItem) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *ListActionLogCallsItem) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *ListActionLogCallsItem) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### SetErrorMessageNil

`func (o *ListActionLogCallsItem) SetErrorMessageNil(b bool)`

 SetErrorMessageNil sets the value for ErrorMessage to be an explicit nil

### UnsetErrorMessage
`func (o *ListActionLogCallsItem) UnsetErrorMessage()`

UnsetErrorMessage ensures that no value is present for ErrorMessage, not even an explicit nil
### GetRequestBodyExcerpt

`func (o *ListActionLogCallsItem) GetRequestBodyExcerpt() string`

GetRequestBodyExcerpt returns the RequestBodyExcerpt field if non-nil, zero value otherwise.

### GetRequestBodyExcerptOk

`func (o *ListActionLogCallsItem) GetRequestBodyExcerptOk() (*string, bool)`

GetRequestBodyExcerptOk returns a tuple with the RequestBodyExcerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBodyExcerpt

`func (o *ListActionLogCallsItem) SetRequestBodyExcerpt(v string)`

SetRequestBodyExcerpt sets RequestBodyExcerpt field to given value.

### HasRequestBodyExcerpt

`func (o *ListActionLogCallsItem) HasRequestBodyExcerpt() bool`

HasRequestBodyExcerpt returns a boolean if a field has been set.

### SetRequestBodyExcerptNil

`func (o *ListActionLogCallsItem) SetRequestBodyExcerptNil(b bool)`

 SetRequestBodyExcerptNil sets the value for RequestBodyExcerpt to be an explicit nil

### UnsetRequestBodyExcerpt
`func (o *ListActionLogCallsItem) UnsetRequestBodyExcerpt()`

UnsetRequestBodyExcerpt ensures that no value is present for RequestBodyExcerpt, not even an explicit nil
### GetResponseBodyExcerpt

`func (o *ListActionLogCallsItem) GetResponseBodyExcerpt() string`

GetResponseBodyExcerpt returns the ResponseBodyExcerpt field if non-nil, zero value otherwise.

### GetResponseBodyExcerptOk

`func (o *ListActionLogCallsItem) GetResponseBodyExcerptOk() (*string, bool)`

GetResponseBodyExcerptOk returns a tuple with the ResponseBodyExcerpt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponseBodyExcerpt

`func (o *ListActionLogCallsItem) SetResponseBodyExcerpt(v string)`

SetResponseBodyExcerpt sets ResponseBodyExcerpt field to given value.

### HasResponseBodyExcerpt

`func (o *ListActionLogCallsItem) HasResponseBodyExcerpt() bool`

HasResponseBodyExcerpt returns a boolean if a field has been set.

### SetResponseBodyExcerptNil

`func (o *ListActionLogCallsItem) SetResponseBodyExcerptNil(b bool)`

 SetResponseBodyExcerptNil sets the value for ResponseBodyExcerpt to be an explicit nil

### UnsetResponseBodyExcerpt
`func (o *ListActionLogCallsItem) UnsetResponseBodyExcerpt()`

UnsetResponseBodyExcerpt ensures that no value is present for ResponseBodyExcerpt, not even an explicit nil
### GetRecipeId

`func (o *ListActionLogCallsItem) GetRecipeId() string`

GetRecipeId returns the RecipeId field if non-nil, zero value otherwise.

### GetRecipeIdOk

`func (o *ListActionLogCallsItem) GetRecipeIdOk() (*string, bool)`

GetRecipeIdOk returns a tuple with the RecipeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipeId

`func (o *ListActionLogCallsItem) SetRecipeId(v string)`

SetRecipeId sets RecipeId field to given value.

### HasRecipeId

`func (o *ListActionLogCallsItem) HasRecipeId() bool`

HasRecipeId returns a boolean if a field has been set.

### SetRecipeIdNil

`func (o *ListActionLogCallsItem) SetRecipeIdNil(b bool)`

 SetRecipeIdNil sets the value for RecipeId to be an explicit nil

### UnsetRecipeId
`func (o *ListActionLogCallsItem) UnsetRecipeId()`

UnsetRecipeId ensures that no value is present for RecipeId, not even an explicit nil
### GetDocIdOrHash

`func (o *ListActionLogCallsItem) GetDocIdOrHash() string`

GetDocIdOrHash returns the DocIdOrHash field if non-nil, zero value otherwise.

### GetDocIdOrHashOk

`func (o *ListActionLogCallsItem) GetDocIdOrHashOk() (*string, bool)`

GetDocIdOrHashOk returns a tuple with the DocIdOrHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocIdOrHash

`func (o *ListActionLogCallsItem) SetDocIdOrHash(v string)`

SetDocIdOrHash sets DocIdOrHash field to given value.

### HasDocIdOrHash

`func (o *ListActionLogCallsItem) HasDocIdOrHash() bool`

HasDocIdOrHash returns a boolean if a field has been set.

### SetDocIdOrHashNil

`func (o *ListActionLogCallsItem) SetDocIdOrHashNil(b bool)`

 SetDocIdOrHashNil sets the value for DocIdOrHash to be an explicit nil

### UnsetDocIdOrHash
`func (o *ListActionLogCallsItem) UnsetDocIdOrHash()`

UnsetDocIdOrHash ensures that no value is present for DocIdOrHash, not even an explicit nil
### GetProxyId

`func (o *ListActionLogCallsItem) GetProxyId() string`

GetProxyId returns the ProxyId field if non-nil, zero value otherwise.

### GetProxyIdOk

`func (o *ListActionLogCallsItem) GetProxyIdOk() (*string, bool)`

GetProxyIdOk returns a tuple with the ProxyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyId

`func (o *ListActionLogCallsItem) SetProxyId(v string)`

SetProxyId sets ProxyId field to given value.

### HasProxyId

`func (o *ListActionLogCallsItem) HasProxyId() bool`

HasProxyId returns a boolean if a field has been set.

### SetProxyIdNil

`func (o *ListActionLogCallsItem) SetProxyIdNil(b bool)`

 SetProxyIdNil sets the value for ProxyId to be an explicit nil

### UnsetProxyId
`func (o *ListActionLogCallsItem) UnsetProxyId()`

UnsetProxyId ensures that no value is present for ProxyId, not even an explicit nil
### GetCorrelationId

`func (o *ListActionLogCallsItem) GetCorrelationId() string`

GetCorrelationId returns the CorrelationId field if non-nil, zero value otherwise.

### GetCorrelationIdOk

`func (o *ListActionLogCallsItem) GetCorrelationIdOk() (*string, bool)`

GetCorrelationIdOk returns a tuple with the CorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCorrelationId

`func (o *ListActionLogCallsItem) SetCorrelationId(v string)`

SetCorrelationId sets CorrelationId field to given value.


### GetActionEventId

`func (o *ListActionLogCallsItem) GetActionEventId() string`

GetActionEventId returns the ActionEventId field if non-nil, zero value otherwise.

### GetActionEventIdOk

`func (o *ListActionLogCallsItem) GetActionEventIdOk() (*string, bool)`

GetActionEventIdOk returns a tuple with the ActionEventId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActionEventId

`func (o *ListActionLogCallsItem) SetActionEventId(v string)`

SetActionEventId sets ActionEventId field to given value.

### HasActionEventId

`func (o *ListActionLogCallsItem) HasActionEventId() bool`

HasActionEventId returns a boolean if a field has been set.

### SetActionEventIdNil

`func (o *ListActionLogCallsItem) SetActionEventIdNil(b bool)`

 SetActionEventIdNil sets the value for ActionEventId to be an explicit nil

### UnsetActionEventId
`func (o *ListActionLogCallsItem) UnsetActionEventId()`

UnsetActionEventId ensures that no value is present for ActionEventId, not even an explicit nil
### GetOk

`func (o *ListActionLogCallsItem) GetOk() bool`

GetOk returns the Ok field if non-nil, zero value otherwise.

### GetOkOk

`func (o *ListActionLogCallsItem) GetOkOk() (*bool, bool)`

GetOkOk returns a tuple with the Ok field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOk

`func (o *ListActionLogCallsItem) SetOk(v bool)`

SetOk sets Ok field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


