# GetOrderProofOfDeliveryResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OrderId** | **string** |  | 
**Platform** | **string** |  | 
**PlatformOrderId** | Pointer to **NullableString** |  | [optional] 
**ItemTitle** | Pointer to **NullableString** |  | [optional] 
**BuyerUsername** | Pointer to **NullableString** |  | [optional] 
**ShipToPostalCode** | Pointer to **NullableString** | The ZIP we shipped to, for comparison against the delivery scan. | [optional] 
**ShipToCityState** | Pointer to **NullableString** |  | [optional] 
**Carrier** | Pointer to **NullableString** |  | [optional] 
**TrackingNumber** | Pointer to **NullableString** |  | [optional] 
**TrackingUrl** | Pointer to **NullableString** |  | [optional] 
**ShippedAt** | Pointer to **NullableString** |  | [optional] 
**DeliveredAt** | Pointer to **NullableString** |  | [optional] 
**DeliveryLocation** | Pointer to **NullableString** |  | [optional] 
**Signature** | Pointer to **NullableString** | Null means the carrier captured none — NOT that delivery is unproven. | [optional] 
**Scans** | [**[]GetOrderProofOfDeliveryResponseScans**](GetOrderProofOfDeliveryResponseScans.md) |  | 
**Gaps** | **[]string** | Why this document is weak, stated plainly so the seller isn&#39;t surprised  by the marketplace&#39;s response. | 

## Methods

### NewGetOrderProofOfDeliveryResponse

`func NewGetOrderProofOfDeliveryResponse(orderId string, platform string, scans []GetOrderProofOfDeliveryResponseScans, gaps []string, ) *GetOrderProofOfDeliveryResponse`

NewGetOrderProofOfDeliveryResponse instantiates a new GetOrderProofOfDeliveryResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderProofOfDeliveryResponseWithDefaults

`func NewGetOrderProofOfDeliveryResponseWithDefaults() *GetOrderProofOfDeliveryResponse`

NewGetOrderProofOfDeliveryResponseWithDefaults instantiates a new GetOrderProofOfDeliveryResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrderId

`func (o *GetOrderProofOfDeliveryResponse) GetOrderId() string`

GetOrderId returns the OrderId field if non-nil, zero value otherwise.

### GetOrderIdOk

`func (o *GetOrderProofOfDeliveryResponse) GetOrderIdOk() (*string, bool)`

GetOrderIdOk returns a tuple with the OrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderId

`func (o *GetOrderProofOfDeliveryResponse) SetOrderId(v string)`

SetOrderId sets OrderId field to given value.


### GetPlatform

`func (o *GetOrderProofOfDeliveryResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetOrderProofOfDeliveryResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetOrderProofOfDeliveryResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetPlatformOrderId

`func (o *GetOrderProofOfDeliveryResponse) GetPlatformOrderId() string`

GetPlatformOrderId returns the PlatformOrderId field if non-nil, zero value otherwise.

### GetPlatformOrderIdOk

`func (o *GetOrderProofOfDeliveryResponse) GetPlatformOrderIdOk() (*string, bool)`

GetPlatformOrderIdOk returns a tuple with the PlatformOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformOrderId

`func (o *GetOrderProofOfDeliveryResponse) SetPlatformOrderId(v string)`

SetPlatformOrderId sets PlatformOrderId field to given value.

### HasPlatformOrderId

`func (o *GetOrderProofOfDeliveryResponse) HasPlatformOrderId() bool`

HasPlatformOrderId returns a boolean if a field has been set.

### SetPlatformOrderIdNil

`func (o *GetOrderProofOfDeliveryResponse) SetPlatformOrderIdNil(b bool)`

 SetPlatformOrderIdNil sets the value for PlatformOrderId to be an explicit nil

### UnsetPlatformOrderId
`func (o *GetOrderProofOfDeliveryResponse) UnsetPlatformOrderId()`

UnsetPlatformOrderId ensures that no value is present for PlatformOrderId, not even an explicit nil
### GetItemTitle

`func (o *GetOrderProofOfDeliveryResponse) GetItemTitle() string`

GetItemTitle returns the ItemTitle field if non-nil, zero value otherwise.

### GetItemTitleOk

`func (o *GetOrderProofOfDeliveryResponse) GetItemTitleOk() (*string, bool)`

GetItemTitleOk returns a tuple with the ItemTitle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItemTitle

`func (o *GetOrderProofOfDeliveryResponse) SetItemTitle(v string)`

SetItemTitle sets ItemTitle field to given value.

### HasItemTitle

`func (o *GetOrderProofOfDeliveryResponse) HasItemTitle() bool`

HasItemTitle returns a boolean if a field has been set.

### SetItemTitleNil

`func (o *GetOrderProofOfDeliveryResponse) SetItemTitleNil(b bool)`

 SetItemTitleNil sets the value for ItemTitle to be an explicit nil

### UnsetItemTitle
`func (o *GetOrderProofOfDeliveryResponse) UnsetItemTitle()`

UnsetItemTitle ensures that no value is present for ItemTitle, not even an explicit nil
### GetBuyerUsername

`func (o *GetOrderProofOfDeliveryResponse) GetBuyerUsername() string`

GetBuyerUsername returns the BuyerUsername field if non-nil, zero value otherwise.

### GetBuyerUsernameOk

`func (o *GetOrderProofOfDeliveryResponse) GetBuyerUsernameOk() (*string, bool)`

GetBuyerUsernameOk returns a tuple with the BuyerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerUsername

`func (o *GetOrderProofOfDeliveryResponse) SetBuyerUsername(v string)`

SetBuyerUsername sets BuyerUsername field to given value.

### HasBuyerUsername

`func (o *GetOrderProofOfDeliveryResponse) HasBuyerUsername() bool`

HasBuyerUsername returns a boolean if a field has been set.

### SetBuyerUsernameNil

`func (o *GetOrderProofOfDeliveryResponse) SetBuyerUsernameNil(b bool)`

 SetBuyerUsernameNil sets the value for BuyerUsername to be an explicit nil

### UnsetBuyerUsername
`func (o *GetOrderProofOfDeliveryResponse) UnsetBuyerUsername()`

UnsetBuyerUsername ensures that no value is present for BuyerUsername, not even an explicit nil
### GetShipToPostalCode

`func (o *GetOrderProofOfDeliveryResponse) GetShipToPostalCode() string`

GetShipToPostalCode returns the ShipToPostalCode field if non-nil, zero value otherwise.

### GetShipToPostalCodeOk

`func (o *GetOrderProofOfDeliveryResponse) GetShipToPostalCodeOk() (*string, bool)`

GetShipToPostalCodeOk returns a tuple with the ShipToPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToPostalCode

`func (o *GetOrderProofOfDeliveryResponse) SetShipToPostalCode(v string)`

SetShipToPostalCode sets ShipToPostalCode field to given value.

### HasShipToPostalCode

`func (o *GetOrderProofOfDeliveryResponse) HasShipToPostalCode() bool`

HasShipToPostalCode returns a boolean if a field has been set.

### SetShipToPostalCodeNil

`func (o *GetOrderProofOfDeliveryResponse) SetShipToPostalCodeNil(b bool)`

 SetShipToPostalCodeNil sets the value for ShipToPostalCode to be an explicit nil

### UnsetShipToPostalCode
`func (o *GetOrderProofOfDeliveryResponse) UnsetShipToPostalCode()`

UnsetShipToPostalCode ensures that no value is present for ShipToPostalCode, not even an explicit nil
### GetShipToCityState

`func (o *GetOrderProofOfDeliveryResponse) GetShipToCityState() string`

GetShipToCityState returns the ShipToCityState field if non-nil, zero value otherwise.

### GetShipToCityStateOk

`func (o *GetOrderProofOfDeliveryResponse) GetShipToCityStateOk() (*string, bool)`

GetShipToCityStateOk returns a tuple with the ShipToCityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipToCityState

`func (o *GetOrderProofOfDeliveryResponse) SetShipToCityState(v string)`

SetShipToCityState sets ShipToCityState field to given value.

### HasShipToCityState

`func (o *GetOrderProofOfDeliveryResponse) HasShipToCityState() bool`

HasShipToCityState returns a boolean if a field has been set.

### SetShipToCityStateNil

`func (o *GetOrderProofOfDeliveryResponse) SetShipToCityStateNil(b bool)`

 SetShipToCityStateNil sets the value for ShipToCityState to be an explicit nil

### UnsetShipToCityState
`func (o *GetOrderProofOfDeliveryResponse) UnsetShipToCityState()`

UnsetShipToCityState ensures that no value is present for ShipToCityState, not even an explicit nil
### GetCarrier

`func (o *GetOrderProofOfDeliveryResponse) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *GetOrderProofOfDeliveryResponse) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *GetOrderProofOfDeliveryResponse) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *GetOrderProofOfDeliveryResponse) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### SetCarrierNil

`func (o *GetOrderProofOfDeliveryResponse) SetCarrierNil(b bool)`

 SetCarrierNil sets the value for Carrier to be an explicit nil

### UnsetCarrier
`func (o *GetOrderProofOfDeliveryResponse) UnsetCarrier()`

UnsetCarrier ensures that no value is present for Carrier, not even an explicit nil
### GetTrackingNumber

`func (o *GetOrderProofOfDeliveryResponse) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *GetOrderProofOfDeliveryResponse) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *GetOrderProofOfDeliveryResponse) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *GetOrderProofOfDeliveryResponse) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *GetOrderProofOfDeliveryResponse) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *GetOrderProofOfDeliveryResponse) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
### GetTrackingUrl

`func (o *GetOrderProofOfDeliveryResponse) GetTrackingUrl() string`

GetTrackingUrl returns the TrackingUrl field if non-nil, zero value otherwise.

### GetTrackingUrlOk

`func (o *GetOrderProofOfDeliveryResponse) GetTrackingUrlOk() (*string, bool)`

GetTrackingUrlOk returns a tuple with the TrackingUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingUrl

`func (o *GetOrderProofOfDeliveryResponse) SetTrackingUrl(v string)`

SetTrackingUrl sets TrackingUrl field to given value.

### HasTrackingUrl

`func (o *GetOrderProofOfDeliveryResponse) HasTrackingUrl() bool`

HasTrackingUrl returns a boolean if a field has been set.

### SetTrackingUrlNil

`func (o *GetOrderProofOfDeliveryResponse) SetTrackingUrlNil(b bool)`

 SetTrackingUrlNil sets the value for TrackingUrl to be an explicit nil

### UnsetTrackingUrl
`func (o *GetOrderProofOfDeliveryResponse) UnsetTrackingUrl()`

UnsetTrackingUrl ensures that no value is present for TrackingUrl, not even an explicit nil
### GetShippedAt

`func (o *GetOrderProofOfDeliveryResponse) GetShippedAt() string`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *GetOrderProofOfDeliveryResponse) GetShippedAtOk() (*string, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *GetOrderProofOfDeliveryResponse) SetShippedAt(v string)`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *GetOrderProofOfDeliveryResponse) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### SetShippedAtNil

`func (o *GetOrderProofOfDeliveryResponse) SetShippedAtNil(b bool)`

 SetShippedAtNil sets the value for ShippedAt to be an explicit nil

### UnsetShippedAt
`func (o *GetOrderProofOfDeliveryResponse) UnsetShippedAt()`

UnsetShippedAt ensures that no value is present for ShippedAt, not even an explicit nil
### GetDeliveredAt

`func (o *GetOrderProofOfDeliveryResponse) GetDeliveredAt() string`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *GetOrderProofOfDeliveryResponse) GetDeliveredAtOk() (*string, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *GetOrderProofOfDeliveryResponse) SetDeliveredAt(v string)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *GetOrderProofOfDeliveryResponse) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### SetDeliveredAtNil

`func (o *GetOrderProofOfDeliveryResponse) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *GetOrderProofOfDeliveryResponse) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil
### GetDeliveryLocation

`func (o *GetOrderProofOfDeliveryResponse) GetDeliveryLocation() string`

GetDeliveryLocation returns the DeliveryLocation field if non-nil, zero value otherwise.

### GetDeliveryLocationOk

`func (o *GetOrderProofOfDeliveryResponse) GetDeliveryLocationOk() (*string, bool)`

GetDeliveryLocationOk returns a tuple with the DeliveryLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLocation

`func (o *GetOrderProofOfDeliveryResponse) SetDeliveryLocation(v string)`

SetDeliveryLocation sets DeliveryLocation field to given value.

### HasDeliveryLocation

`func (o *GetOrderProofOfDeliveryResponse) HasDeliveryLocation() bool`

HasDeliveryLocation returns a boolean if a field has been set.

### SetDeliveryLocationNil

`func (o *GetOrderProofOfDeliveryResponse) SetDeliveryLocationNil(b bool)`

 SetDeliveryLocationNil sets the value for DeliveryLocation to be an explicit nil

### UnsetDeliveryLocation
`func (o *GetOrderProofOfDeliveryResponse) UnsetDeliveryLocation()`

UnsetDeliveryLocation ensures that no value is present for DeliveryLocation, not even an explicit nil
### GetSignature

`func (o *GetOrderProofOfDeliveryResponse) GetSignature() string`

GetSignature returns the Signature field if non-nil, zero value otherwise.

### GetSignatureOk

`func (o *GetOrderProofOfDeliveryResponse) GetSignatureOk() (*string, bool)`

GetSignatureOk returns a tuple with the Signature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSignature

`func (o *GetOrderProofOfDeliveryResponse) SetSignature(v string)`

SetSignature sets Signature field to given value.

### HasSignature

`func (o *GetOrderProofOfDeliveryResponse) HasSignature() bool`

HasSignature returns a boolean if a field has been set.

### SetSignatureNil

`func (o *GetOrderProofOfDeliveryResponse) SetSignatureNil(b bool)`

 SetSignatureNil sets the value for Signature to be an explicit nil

### UnsetSignature
`func (o *GetOrderProofOfDeliveryResponse) UnsetSignature()`

UnsetSignature ensures that no value is present for Signature, not even an explicit nil
### GetScans

`func (o *GetOrderProofOfDeliveryResponse) GetScans() []GetOrderProofOfDeliveryResponseScans`

GetScans returns the Scans field if non-nil, zero value otherwise.

### GetScansOk

`func (o *GetOrderProofOfDeliveryResponse) GetScansOk() (*[]GetOrderProofOfDeliveryResponseScans, bool)`

GetScansOk returns a tuple with the Scans field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScans

`func (o *GetOrderProofOfDeliveryResponse) SetScans(v []GetOrderProofOfDeliveryResponseScans)`

SetScans sets Scans field to given value.


### GetGaps

`func (o *GetOrderProofOfDeliveryResponse) GetGaps() []string`

GetGaps returns the Gaps field if non-nil, zero value otherwise.

### GetGapsOk

`func (o *GetOrderProofOfDeliveryResponse) GetGapsOk() (*[]string, bool)`

GetGapsOk returns a tuple with the Gaps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGaps

`func (o *GetOrderProofOfDeliveryResponse) SetGaps(v []string)`

SetGaps sets Gaps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


