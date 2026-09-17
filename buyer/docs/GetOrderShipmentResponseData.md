# GetOrderShipmentResponseData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Source** | **string** |  | 
**OrderId** | **string** |  | 
**Carrier** | Pointer to **NullableString** |  | [optional] 
**Service** | Pointer to **NullableString** |  | [optional] 
**TrackingNumber** | Pointer to **NullableString** |  | [optional] 
**VoidedAt** | Pointer to **NullableTime** |  | [optional] 
**LabelCostCents** | Pointer to **NullableFloat32** |  | [optional] 
**LabelUrl** | Pointer to **NullableString** |  | [optional] 
**ShippoShipmentId** | Pointer to **NullableString** |  | [optional] 
**TrackingRegisteredAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewGetOrderShipmentResponseData

`func NewGetOrderShipmentResponseData(id string, createdAt time.Time, userId string, source string, orderId string, ) *GetOrderShipmentResponseData`

NewGetOrderShipmentResponseData instantiates a new GetOrderShipmentResponseData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderShipmentResponseDataWithDefaults

`func NewGetOrderShipmentResponseDataWithDefaults() *GetOrderShipmentResponseData`

NewGetOrderShipmentResponseDataWithDefaults instantiates a new GetOrderShipmentResponseData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrderShipmentResponseData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrderShipmentResponseData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrderShipmentResponseData) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetOrderShipmentResponseData) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOrderShipmentResponseData) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOrderShipmentResponseData) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUserId

`func (o *GetOrderShipmentResponseData) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetOrderShipmentResponseData) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetOrderShipmentResponseData) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetSource

`func (o *GetOrderShipmentResponseData) GetSource() string`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *GetOrderShipmentResponseData) GetSourceOk() (*string, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *GetOrderShipmentResponseData) SetSource(v string)`

SetSource sets Source field to given value.


### GetOrderId

`func (o *GetOrderShipmentResponseData) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetOrderShipmentResponseData) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetOrderShipmentResponseData) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetCarrier

`func (o *GetOrderShipmentResponseData) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *GetOrderShipmentResponseData) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *GetOrderShipmentResponseData) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *GetOrderShipmentResponseData) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### SetCarrierNil

`func (o *GetOrderShipmentResponseData) SetCarrierNil(b bool)`

 SetCarrierNil sets the value for Carrier to be an explicit nil

### UnsetCarrier
`func (o *GetOrderShipmentResponseData) UnsetCarrier()`

UnsetCarrier ensures that no value is present for Carrier, not even an explicit nil
### GetService

`func (o *GetOrderShipmentResponseData) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetOrderShipmentResponseData) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetOrderShipmentResponseData) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GetOrderShipmentResponseData) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *GetOrderShipmentResponseData) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *GetOrderShipmentResponseData) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetTrackingNumber

`func (o *GetOrderShipmentResponseData) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *GetOrderShipmentResponseData) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *GetOrderShipmentResponseData) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *GetOrderShipmentResponseData) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *GetOrderShipmentResponseData) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *GetOrderShipmentResponseData) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
### GetVoidedAt

`func (o *GetOrderShipmentResponseData) GetVoidedAt() time.Time`

GetVoidedAt returns the VoidedAt field if non-nil, zero value otherwise.

### GetVoidedAtOk

`func (o *GetOrderShipmentResponseData) GetVoidedAtOk() (*time.Time, bool)`

GetVoidedAtOk returns a tuple with the VoidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoidedAt

`func (o *GetOrderShipmentResponseData) SetVoidedAt(v time.Time)`

SetVoidedAt sets VoidedAt field to given value.

### HasVoidedAt

`func (o *GetOrderShipmentResponseData) HasVoidedAt() bool`

HasVoidedAt returns a boolean if a field has been set.

### SetVoidedAtNil

`func (o *GetOrderShipmentResponseData) SetVoidedAtNil(b bool)`

 SetVoidedAtNil sets the value for VoidedAt to be an explicit nil

### UnsetVoidedAt
`func (o *GetOrderShipmentResponseData) UnsetVoidedAt()`

UnsetVoidedAt ensures that no value is present for VoidedAt, not even an explicit nil
### GetLabelCostCents

`func (o *GetOrderShipmentResponseData) GetLabelCostCents() float32`

GetLabelCostCents returns the LabelCostCents field if non-nil, zero value otherwise.

### GetLabelCostCentsOk

`func (o *GetOrderShipmentResponseData) GetLabelCostCentsOk() (*float32, bool)`

GetLabelCostCentsOk returns a tuple with the LabelCostCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelCostCents

`func (o *GetOrderShipmentResponseData) SetLabelCostCents(v float32)`

SetLabelCostCents sets LabelCostCents field to given value.

### HasLabelCostCents

`func (o *GetOrderShipmentResponseData) HasLabelCostCents() bool`

HasLabelCostCents returns a boolean if a field has been set.

### SetLabelCostCentsNil

`func (o *GetOrderShipmentResponseData) SetLabelCostCentsNil(b bool)`

 SetLabelCostCentsNil sets the value for LabelCostCents to be an explicit nil

### UnsetLabelCostCents
`func (o *GetOrderShipmentResponseData) UnsetLabelCostCents()`

UnsetLabelCostCents ensures that no value is present for LabelCostCents, not even an explicit nil
### GetLabelUrl

`func (o *GetOrderShipmentResponseData) GetLabelUrl() string`

GetLabelUrl returns the LabelUrl field if non-nil, zero value otherwise.

### GetLabelUrlOk

`func (o *GetOrderShipmentResponseData) GetLabelUrlOk() (*string, bool)`

GetLabelUrlOk returns a tuple with the LabelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelUrl

`func (o *GetOrderShipmentResponseData) SetLabelUrl(v string)`

SetLabelUrl sets LabelUrl field to given value.

### HasLabelUrl

`func (o *GetOrderShipmentResponseData) HasLabelUrl() bool`

HasLabelUrl returns a boolean if a field has been set.

### SetLabelUrlNil

`func (o *GetOrderShipmentResponseData) SetLabelUrlNil(b bool)`

 SetLabelUrlNil sets the value for LabelUrl to be an explicit nil

### UnsetLabelUrl
`func (o *GetOrderShipmentResponseData) UnsetLabelUrl()`

UnsetLabelUrl ensures that no value is present for LabelUrl, not even an explicit nil
### GetShippoShipmentId

`func (o *GetOrderShipmentResponseData) GetShippoShipmentId() string`

GetShippoShipmentId returns the ShippoShipmentId field if non-nil, zero value otherwise.

### GetShippoShipmentIdOk

`func (o *GetOrderShipmentResponseData) GetShippoShipmentIdOk() (*string, bool)`

GetShippoShipmentIdOk returns a tuple with the ShippoShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippoShipmentId

`func (o *GetOrderShipmentResponseData) SetShippoShipmentId(v string)`

SetShippoShipmentId sets ShippoShipmentId field to given value.

### HasShippoShipmentId

`func (o *GetOrderShipmentResponseData) HasShippoShipmentId() bool`

HasShippoShipmentId returns a boolean if a field has been set.

### SetShippoShipmentIdNil

`func (o *GetOrderShipmentResponseData) SetShippoShipmentIdNil(b bool)`

 SetShippoShipmentIdNil sets the value for ShippoShipmentId to be an explicit nil

### UnsetShippoShipmentId
`func (o *GetOrderShipmentResponseData) UnsetShippoShipmentId()`

UnsetShippoShipmentId ensures that no value is present for ShippoShipmentId, not even an explicit nil
### GetTrackingRegisteredAt

`func (o *GetOrderShipmentResponseData) GetTrackingRegisteredAt() time.Time`

GetTrackingRegisteredAt returns the TrackingRegisteredAt field if non-nil, zero value otherwise.

### GetTrackingRegisteredAtOk

`func (o *GetOrderShipmentResponseData) GetTrackingRegisteredAtOk() (*time.Time, bool)`

GetTrackingRegisteredAtOk returns a tuple with the TrackingRegisteredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingRegisteredAt

`func (o *GetOrderShipmentResponseData) SetTrackingRegisteredAt(v time.Time)`

SetTrackingRegisteredAt sets TrackingRegisteredAt field to given value.

### HasTrackingRegisteredAt

`func (o *GetOrderShipmentResponseData) HasTrackingRegisteredAt() bool`

HasTrackingRegisteredAt returns a boolean if a field has been set.

### SetTrackingRegisteredAtNil

`func (o *GetOrderShipmentResponseData) SetTrackingRegisteredAtNil(b bool)`

 SetTrackingRegisteredAtNil sets the value for TrackingRegisteredAt to be an explicit nil

### UnsetTrackingRegisteredAt
`func (o *GetOrderShipmentResponseData) UnsetTrackingRegisteredAt()`

UnsetTrackingRegisteredAt ensures that no value is present for TrackingRegisteredAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


