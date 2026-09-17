# GetOrderCancelEligibilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Eligible** | **bool** |  | 
**EligibleCancelReason** | **[]interface{}** |  | 
**FailureReason** | **[]interface{}** |  | 
**Source** | **string** |  | 

## Methods

### NewGetOrderCancelEligibilityResponse

`func NewGetOrderCancelEligibilityResponse(eligible bool, eligibleCancelReason []interface{}, failureReason []interface{}, source string, ) *GetOrderCancelEligibilityResponse`

NewGetOrderCancelEligibilityResponse instantiates a new GetOrderCancelEligibilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderCancelEligibilityResponseWithDefaults

`func NewGetOrderCancelEligibilityResponseWithDefaults() *GetOrderCancelEligibilityResponse`

NewGetOrderCancelEligibilityResponseWithDefaults instantiates a new GetOrderCancelEligibilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEligible

`func (o *GetOrderCancelEligibilityResponse) GetEligible() bool`

GetEligible returns the Eligible field if non-nil, zero value otherwise.

### GetEligibleOk

`func (o *GetOrderCancelEligibilityResponse) GetEligibleOk() (*bool, bool)`

GetEligibleOk returns a tuple with the Eligible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligible

`func (o *GetOrderCancelEligibilityResponse) SetEligible(v bool)`

SetEligible sets Eligible field to given value.


### GetEligibleCancelReason

`func (o *GetOrderCancelEligibilityResponse) GetEligibleCancelReason() []interface{}`

GetEligibleCancelReason returns the EligibleCancelReason field if non-nil, zero value otherwise.

### GetEligibleCancelReasonOk

`func (o *GetOrderCancelEligibilityResponse) GetEligibleCancelReasonOk() (*[]interface{}, bool)`

GetEligibleCancelReasonOk returns a tuple with the EligibleCancelReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEligibleCancelReason

`func (o *GetOrderCancelEligibilityResponse) SetEligibleCancelReason(v []interface{})`

SetEligibleCancelReason sets EligibleCancelReason field to given value.


### GetFailureReason

`func (o *GetOrderCancelEligibilityResponse) GetFailureReason() []interface{}`

GetFailureReason returns the FailureReason field if non-nil, zero value otherwise.

### GetFailureReasonOk

`func (o *GetOrderCancelEligibilityResponse) GetFailureReasonOk() (*[]interface{}, bool)`

GetFailureReasonOk returns a tuple with the FailureReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFailureReason

`func (o *GetOrderCancelEligibilityResponse) SetFailureReason(v []interface{})`

SetFailureReason sets FailureReason field to given value.


### GetSource

`func (o *GetOrderCancelEligibilityResponse) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetOrderCancelEligibilityResponse) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetOrderCancelEligibilityResponse) SetSource(v string)`

SetSource sets Source field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


