# GetConnectionHealthResponseAnchors

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Expected** | **[]string** | What we were looking for. Empty ⇒ this platform is unmonitored. | 
**Observations** | **[]map[string]interface{}** |  | 
**Present** | **[]string** | Name found carrying a non-empty value — the only honest \&quot;logged in\&quot;. | 
**Empty** | **[]string** | Name found, value is the empty string. The Whatnot class. | 
**Missing** | **[]string** | Name not in the jar at all. | 
**CookieCount** | **float32** |  | 
**ObservedCookieNames** | **[]string** | Cookie names actually in the jar, truncated. This is the payload that turns \&quot;anchors missing\&quot; into a diagnosis: if the jar holds 30 cookies and none are ours, a rename is the likely story; if it holds three device cookies, the browser is signed out. NAMES ONLY — never values. | 

## Methods

### NewGetConnectionHealthResponseAnchors

`func NewGetConnectionHealthResponseAnchors(expected []string, observations []map[string]interface{}, present []string, empty []string, missing []string, cookieCount float32, observedCookieNames []string, ) *GetConnectionHealthResponseAnchors`

NewGetConnectionHealthResponseAnchors instantiates a new GetConnectionHealthResponseAnchors object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetConnectionHealthResponseAnchorsWithDefaults

`func NewGetConnectionHealthResponseAnchorsWithDefaults() *GetConnectionHealthResponseAnchors`

NewGetConnectionHealthResponseAnchorsWithDefaults instantiates a new GetConnectionHealthResponseAnchors object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpected

`func (o *GetConnectionHealthResponseAnchors) GetExpected() []string`

GetExpected returns the Expected field if non-nil, zero value otherwise.

### GetExpectedOk

`func (o *GetConnectionHealthResponseAnchors) GetExpectedOk() (*[]string, bool)`

GetExpectedOk returns a tuple with the Expected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpected

`func (o *GetConnectionHealthResponseAnchors) SetExpected(v []string)`

SetExpected sets Expected field to given value.


### GetObservations

`func (o *GetConnectionHealthResponseAnchors) GetObservations() []map[string]interface{}`

GetObservations returns the Observations field if non-nil, zero value otherwise.

### GetObservationsOk

`func (o *GetConnectionHealthResponseAnchors) GetObservationsOk() (*[]map[string]interface{}, bool)`

GetObservationsOk returns a tuple with the Observations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservations

`func (o *GetConnectionHealthResponseAnchors) SetObservations(v []map[string]interface{})`

SetObservations sets Observations field to given value.


### GetPresent

`func (o *GetConnectionHealthResponseAnchors) GetPresent() []string`

GetPresent returns the Present field if non-nil, zero value otherwise.

### GetPresentOk

`func (o *GetConnectionHealthResponseAnchors) GetPresentOk() (*[]string, bool)`

GetPresentOk returns a tuple with the Present field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresent

`func (o *GetConnectionHealthResponseAnchors) SetPresent(v []string)`

SetPresent sets Present field to given value.


### GetEmpty

`func (o *GetConnectionHealthResponseAnchors) GetEmpty() []string`

GetEmpty returns the Empty field if non-nil, zero value otherwise.

### GetEmptyOk

`func (o *GetConnectionHealthResponseAnchors) GetEmptyOk() (*[]string, bool)`

GetEmptyOk returns a tuple with the Empty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmpty

`func (o *GetConnectionHealthResponseAnchors) SetEmpty(v []string)`

SetEmpty sets Empty field to given value.


### GetMissing

`func (o *GetConnectionHealthResponseAnchors) GetMissing() []string`

GetMissing returns the Missing field if non-nil, zero value otherwise.

### GetMissingOk

`func (o *GetConnectionHealthResponseAnchors) GetMissingOk() (*[]string, bool)`

GetMissingOk returns a tuple with the Missing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMissing

`func (o *GetConnectionHealthResponseAnchors) SetMissing(v []string)`

SetMissing sets Missing field to given value.


### GetCookieCount

`func (o *GetConnectionHealthResponseAnchors) GetCookieCount() float32`

GetCookieCount returns the CookieCount field if non-nil, zero value otherwise.

### GetCookieCountOk

`func (o *GetConnectionHealthResponseAnchors) GetCookieCountOk() (*float32, bool)`

GetCookieCountOk returns a tuple with the CookieCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCookieCount

`func (o *GetConnectionHealthResponseAnchors) SetCookieCount(v float32)`

SetCookieCount sets CookieCount field to given value.


### GetObservedCookieNames

`func (o *GetConnectionHealthResponseAnchors) GetObservedCookieNames() []string`

GetObservedCookieNames returns the ObservedCookieNames field if non-nil, zero value otherwise.

### GetObservedCookieNamesOk

`func (o *GetConnectionHealthResponseAnchors) GetObservedCookieNamesOk() (*[]string, bool)`

GetObservedCookieNamesOk returns a tuple with the ObservedCookieNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObservedCookieNames

`func (o *GetConnectionHealthResponseAnchors) SetObservedCookieNames(v []string)`

SetObservedCookieNames sets ObservedCookieNames field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


