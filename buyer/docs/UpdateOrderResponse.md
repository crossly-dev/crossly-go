# UpdateOrderResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**CreatedAt** | **time.Time** |  | 
**UpdatedAt** | **time.Time** |  | 
**UserId** | **string** |  | 
**Quantity** | **float32** |  | 
**Notes** | Pointer to **NullableString** |  | [optional] 
**Status** | **string** |  | 
**Platform** | **string** |  | 
**CancelledAt** | Pointer to **NullableTime** |  | [optional] 
**ListingId** | Pointer to **NullableString** |  | [optional] 
**InventoryItemId** | Pointer to **NullableString** |  | [optional] 
**PlatformListingId** | Pointer to **NullableString** |  | [optional] 
**HandlingTimeDays** | Pointer to **NullableFloat32** |  | [optional] 
**PlatformOrderId** | Pointer to **NullableString** |  | [optional] 
**BuyerUsername** | Pointer to **NullableString** |  | [optional] 
**SalesChannel** | **string** |  | 
**ChannelLocationId** | Pointer to **NullableString** |  | [optional] 
**PackagePresetId** | Pointer to **NullableString** |  | [optional] 
**Carrier** | Pointer to **NullableString** |  | [optional] 
**Service** | Pointer to **NullableString** |  | [optional] 
**TrackingNumber** | Pointer to **NullableString** |  | [optional] 
**EasypostShipmentId** | Pointer to **NullableString** |  | [optional] 
**EasypostTrackerId** | Pointer to **NullableString** |  | [optional] 
**EasypostRateId** | Pointer to **NullableString** |  | [optional] 
**ShippingLabelUrl** | Pointer to **NullableString** |  | [optional] 
**LabelPurchasedAt** | Pointer to **NullableTime** |  | [optional] 
**ShippedAt** | Pointer to **NullableTime** |  | [optional] 
**EstimatedDelivery** | Pointer to **NullableString** |  | [optional] 
**DeliveredAt** | Pointer to **NullableTime** |  | [optional] 
**CarrierStatus** | Pointer to **NullableString** |  | [optional] 
**CarrierStatusDetail** | Pointer to **NullableString** |  | [optional] 
**TrackingHistory** | Pointer to [**[]ListOrdersItemTrackingHistory**](ListOrdersItemTrackingHistory.md) |  | [optional] 
**DeliveryLocation** | Pointer to **NullableString** |  | [optional] 
**DeliverySignature** | Pointer to **NullableString** |  | [optional] 
**TrackingSubmittedAt** | Pointer to **NullableTime** |  | [optional] 
**TrackingSubmitStatus** | Pointer to **NullableString** |  | [optional] 
**LabelCost** | Pointer to **NullableString** |  | [optional] 
**CostOfGoods** | Pointer to **NullableString** |  | [optional] 
**RequestedCarrier** | Pointer to **NullableString** |  | [optional] 
**RequestedService** | Pointer to **NullableString** |  | [optional] 
**ShipByAt** | Pointer to **NullableTime** |  | [optional] 
**ShipByAlertedAt** | Pointer to **NullableTime** |  | [optional] 
**OversoldBy** | **float32** |  | 
**LastStatusCheckAt** | Pointer to **NullableTime** |  | [optional] 
**LastChatCheckAt** | Pointer to **NullableTime** |  | [optional] 
**IsDisputed** | **bool** |  | 
**DisputeReason** | Pointer to **NullableString** |  | [optional] 
**DisputePlatformCaseId** | Pointer to **NullableString** |  | [optional] 
**DisputeResolvedAt** | Pointer to **NullableTime** |  | [optional] 
**RefundAmount** | Pointer to **NullableString** |  | [optional] 
**RefundReason** | Pointer to **NullableString** |  | [optional] 
**RefundPlatformId** | Pointer to **NullableString** |  | [optional] 
**RefundedAt** | Pointer to **NullableTime** |  | [optional] 
**CancellationReason** | Pointer to **NullableString** |  | [optional] 
**ArrivalConditionRequestedAt** | Pointer to **NullableTime** |  | [optional] 
**ArrivalConditionSubmittedAt** | Pointer to **NullableTime** |  | [optional] 
**ArrivalConditionDeclinedAt** | Pointer to **NullableTime** |  | [optional] 
**DeliveryPhotoUrl** | Pointer to **NullableString** |  | [optional] 
**PurchaseOrderRef** | Pointer to **NullableString** |  | [optional] 
**ArrivalConditionPhotos** | **[]string** |  | 
**BuyerEmail** | Pointer to **NullableString** |  | [optional] 
**FulfillmentMethod** | **string** |  | 
**DeletedAt** | Pointer to **NullableTime** |  | [optional] 

## Methods

### NewUpdateOrderResponse

`func NewUpdateOrderResponse(id string, createdAt time.Time, updatedAt time.Time, userId string, quantity float32, status string, platform string, salesChannel string, oversoldBy float32, isDisputed bool, arrivalConditionPhotos []string, fulfillmentMethod string, ) *UpdateOrderResponse`

NewUpdateOrderResponse instantiates a new UpdateOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOrderResponseWithDefaults

`func NewUpdateOrderResponseWithDefaults() *UpdateOrderResponse`

NewUpdateOrderResponseWithDefaults instantiates a new UpdateOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateOrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOrderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *UpdateOrderResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UpdateOrderResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UpdateOrderResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *UpdateOrderResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UpdateOrderResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UpdateOrderResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *UpdateOrderResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UpdateOrderResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UpdateOrderResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetQuantity

`func (o *UpdateOrderResponse) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *UpdateOrderResponse) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *UpdateOrderResponse) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetNotes

`func (o *UpdateOrderResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *UpdateOrderResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *UpdateOrderResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *UpdateOrderResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *UpdateOrderResponse) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *UpdateOrderResponse) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetStatus

`func (o *UpdateOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlatform

`func (o *UpdateOrderResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdateOrderResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdateOrderResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetCancelledAt

`func (o *UpdateOrderResponse) GetCancelledAt() time.Time`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *UpdateOrderResponse) GetCancelledAtOk() (*time.Time, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *UpdateOrderResponse) SetCancelledAt(v time.Time)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *UpdateOrderResponse) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### SetCancelledAtNil

`func (o *UpdateOrderResponse) SetCancelledAtNil(b bool)`

 SetCancelledAtNil sets the value for CancelledAt to be an explicit nil

### UnsetCancelledAt
`func (o *UpdateOrderResponse) UnsetCancelledAt()`

UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
### GetListingId

`func (o *UpdateOrderResponse) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *UpdateOrderResponse) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *UpdateOrderResponse) SetListingId(v string)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *UpdateOrderResponse) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### SetListingIdNil

`func (o *UpdateOrderResponse) SetListingIdNil(b bool)`

 SetListingIdNil sets the value for ListingId to be an explicit nil

### UnsetListingId
`func (o *UpdateOrderResponse) UnsetListingId()`

UnsetListingId ensures that no value is present for ListingId, not even an explicit nil
### GetInventoryItemId

`func (o *UpdateOrderResponse) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *UpdateOrderResponse) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *UpdateOrderResponse) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *UpdateOrderResponse) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### SetInventoryItemIdNil

`func (o *UpdateOrderResponse) SetInventoryItemIdNil(b bool)`

 SetInventoryItemIdNil sets the value for InventoryItemId to be an explicit nil

### UnsetInventoryItemId
`func (o *UpdateOrderResponse) UnsetInventoryItemId()`

UnsetInventoryItemId ensures that no value is present for InventoryItemId, not even an explicit nil
### GetPlatformListingId

`func (o *UpdateOrderResponse) GetPlatformListingId() string`

GetPlatformListingId returns the PlatformListingId field if non-nil, zero value otherwise.

### GetPlatformListingIdOk

`func (o *UpdateOrderResponse) GetPlatformListingIdOk() (*string, bool)`

GetPlatformListingIdOk returns a tuple with the PlatformListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformListingId

`func (o *UpdateOrderResponse) SetPlatformListingId(v string)`

SetPlatformListingId sets PlatformListingId field to given value.

### HasPlatformListingId

`func (o *UpdateOrderResponse) HasPlatformListingId() bool`

HasPlatformListingId returns a boolean if a field has been set.

### SetPlatformListingIdNil

`func (o *UpdateOrderResponse) SetPlatformListingIdNil(b bool)`

 SetPlatformListingIdNil sets the value for PlatformListingId to be an explicit nil

### UnsetPlatformListingId
`func (o *UpdateOrderResponse) UnsetPlatformListingId()`

UnsetPlatformListingId ensures that no value is present for PlatformListingId, not even an explicit nil
### GetHandlingTimeDays

`func (o *UpdateOrderResponse) GetHandlingTimeDays() float32`

GetHandlingTimeDays returns the HandlingTimeDays field if non-nil, zero value otherwise.

### GetHandlingTimeDaysOk

`func (o *UpdateOrderResponse) GetHandlingTimeDaysOk() (*float32, bool)`

GetHandlingTimeDaysOk returns a tuple with the HandlingTimeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingTimeDays

`func (o *UpdateOrderResponse) SetHandlingTimeDays(v float32)`

SetHandlingTimeDays sets HandlingTimeDays field to given value.

### HasHandlingTimeDays

`func (o *UpdateOrderResponse) HasHandlingTimeDays() bool`

HasHandlingTimeDays returns a boolean if a field has been set.

### SetHandlingTimeDaysNil

`func (o *UpdateOrderResponse) SetHandlingTimeDaysNil(b bool)`

 SetHandlingTimeDaysNil sets the value for HandlingTimeDays to be an explicit nil

### UnsetHandlingTimeDays
`func (o *UpdateOrderResponse) UnsetHandlingTimeDays()`

UnsetHandlingTimeDays ensures that no value is present for HandlingTimeDays, not even an explicit nil
### GetPlatformOrderId

`func (o *UpdateOrderResponse) GetPlatformOrderId() string`

GetPlatformOrderId returns the PlatformOrderId field if non-nil, zero value otherwise.

### GetPlatformOrderIdOk

`func (o *UpdateOrderResponse) GetPlatformOrderIdOk() (*string, bool)`

GetPlatformOrderIdOk returns a tuple with the PlatformOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformOrderId

`func (o *UpdateOrderResponse) SetPlatformOrderId(v string)`

SetPlatformOrderId sets PlatformOrderId field to given value.

### HasPlatformOrderId

`func (o *UpdateOrderResponse) HasPlatformOrderId() bool`

HasPlatformOrderId returns a boolean if a field has been set.

### SetPlatformOrderIdNil

`func (o *UpdateOrderResponse) SetPlatformOrderIdNil(b bool)`

 SetPlatformOrderIdNil sets the value for PlatformOrderId to be an explicit nil

### UnsetPlatformOrderId
`func (o *UpdateOrderResponse) UnsetPlatformOrderId()`

UnsetPlatformOrderId ensures that no value is present for PlatformOrderId, not even an explicit nil
### GetBuyerUsername

`func (o *UpdateOrderResponse) GetBuyerUsername() string`

GetBuyerUsername returns the BuyerUsername field if non-nil, zero value otherwise.

### GetBuyerUsernameOk

`func (o *UpdateOrderResponse) GetBuyerUsernameOk() (*string, bool)`

GetBuyerUsernameOk returns a tuple with the BuyerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerUsername

`func (o *UpdateOrderResponse) SetBuyerUsername(v string)`

SetBuyerUsername sets BuyerUsername field to given value.

### HasBuyerUsername

`func (o *UpdateOrderResponse) HasBuyerUsername() bool`

HasBuyerUsername returns a boolean if a field has been set.

### SetBuyerUsernameNil

`func (o *UpdateOrderResponse) SetBuyerUsernameNil(b bool)`

 SetBuyerUsernameNil sets the value for BuyerUsername to be an explicit nil

### UnsetBuyerUsername
`func (o *UpdateOrderResponse) UnsetBuyerUsername()`

UnsetBuyerUsername ensures that no value is present for BuyerUsername, not even an explicit nil
### GetSalesChannel

`func (o *UpdateOrderResponse) GetSalesChannel() string`

GetSalesChannel returns the SalesChannel field if non-nil, zero value otherwise.

### GetSalesChannelOk

`func (o *UpdateOrderResponse) GetSalesChannelOk() (*string, bool)`

GetSalesChannelOk returns a tuple with the SalesChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesChannel

`func (o *UpdateOrderResponse) SetSalesChannel(v string)`

SetSalesChannel sets SalesChannel field to given value.


### GetChannelLocationId

`func (o *UpdateOrderResponse) GetChannelLocationId() string`

GetChannelLocationId returns the ChannelLocationId field if non-nil, zero value otherwise.

### GetChannelLocationIdOk

`func (o *UpdateOrderResponse) GetChannelLocationIdOk() (*string, bool)`

GetChannelLocationIdOk returns a tuple with the ChannelLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLocationId

`func (o *UpdateOrderResponse) SetChannelLocationId(v string)`

SetChannelLocationId sets ChannelLocationId field to given value.

### HasChannelLocationId

`func (o *UpdateOrderResponse) HasChannelLocationId() bool`

HasChannelLocationId returns a boolean if a field has been set.

### SetChannelLocationIdNil

`func (o *UpdateOrderResponse) SetChannelLocationIdNil(b bool)`

 SetChannelLocationIdNil sets the value for ChannelLocationId to be an explicit nil

### UnsetChannelLocationId
`func (o *UpdateOrderResponse) UnsetChannelLocationId()`

UnsetChannelLocationId ensures that no value is present for ChannelLocationId, not even an explicit nil
### GetPackagePresetId

`func (o *UpdateOrderResponse) GetPackagePresetId() string`

GetPackagePresetId returns the PackagePresetId field if non-nil, zero value otherwise.

### GetPackagePresetIdOk

`func (o *UpdateOrderResponse) GetPackagePresetIdOk() (*string, bool)`

GetPackagePresetIdOk returns a tuple with the PackagePresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagePresetId

`func (o *UpdateOrderResponse) SetPackagePresetId(v string)`

SetPackagePresetId sets PackagePresetId field to given value.

### HasPackagePresetId

`func (o *UpdateOrderResponse) HasPackagePresetId() bool`

HasPackagePresetId returns a boolean if a field has been set.

### SetPackagePresetIdNil

`func (o *UpdateOrderResponse) SetPackagePresetIdNil(b bool)`

 SetPackagePresetIdNil sets the value for PackagePresetId to be an explicit nil

### UnsetPackagePresetId
`func (o *UpdateOrderResponse) UnsetPackagePresetId()`

UnsetPackagePresetId ensures that no value is present for PackagePresetId, not even an explicit nil
### GetCarrier

`func (o *UpdateOrderResponse) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *UpdateOrderResponse) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *UpdateOrderResponse) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *UpdateOrderResponse) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### SetCarrierNil

`func (o *UpdateOrderResponse) SetCarrierNil(b bool)`

 SetCarrierNil sets the value for Carrier to be an explicit nil

### UnsetCarrier
`func (o *UpdateOrderResponse) UnsetCarrier()`

UnsetCarrier ensures that no value is present for Carrier, not even an explicit nil
### GetService

`func (o *UpdateOrderResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *UpdateOrderResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *UpdateOrderResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *UpdateOrderResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *UpdateOrderResponse) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *UpdateOrderResponse) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetTrackingNumber

`func (o *UpdateOrderResponse) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *UpdateOrderResponse) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *UpdateOrderResponse) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *UpdateOrderResponse) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *UpdateOrderResponse) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *UpdateOrderResponse) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
### GetEasypostShipmentId

`func (o *UpdateOrderResponse) GetEasypostShipmentId() string`

GetEasypostShipmentId returns the EasypostShipmentId field if non-nil, zero value otherwise.

### GetEasypostShipmentIdOk

`func (o *UpdateOrderResponse) GetEasypostShipmentIdOk() (*string, bool)`

GetEasypostShipmentIdOk returns a tuple with the EasypostShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasypostShipmentId

`func (o *UpdateOrderResponse) SetEasypostShipmentId(v string)`

SetEasypostShipmentId sets EasypostShipmentId field to given value.

### HasEasypostShipmentId

`func (o *UpdateOrderResponse) HasEasypostShipmentId() bool`

HasEasypostShipmentId returns a boolean if a field has been set.

### SetEasypostShipmentIdNil

`func (o *UpdateOrderResponse) SetEasypostShipmentIdNil(b bool)`

 SetEasypostShipmentIdNil sets the value for EasypostShipmentId to be an explicit nil

### UnsetEasypostShipmentId
`func (o *UpdateOrderResponse) UnsetEasypostShipmentId()`

UnsetEasypostShipmentId ensures that no value is present for EasypostShipmentId, not even an explicit nil
### GetEasypostTrackerId

`func (o *UpdateOrderResponse) GetEasypostTrackerId() string`

GetEasypostTrackerId returns the EasypostTrackerId field if non-nil, zero value otherwise.

### GetEasypostTrackerIdOk

`func (o *UpdateOrderResponse) GetEasypostTrackerIdOk() (*string, bool)`

GetEasypostTrackerIdOk returns a tuple with the EasypostTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasypostTrackerId

`func (o *UpdateOrderResponse) SetEasypostTrackerId(v string)`

SetEasypostTrackerId sets EasypostTrackerId field to given value.

### HasEasypostTrackerId

`func (o *UpdateOrderResponse) HasEasypostTrackerId() bool`

HasEasypostTrackerId returns a boolean if a field has been set.

### SetEasypostTrackerIdNil

`func (o *UpdateOrderResponse) SetEasypostTrackerIdNil(b bool)`

 SetEasypostTrackerIdNil sets the value for EasypostTrackerId to be an explicit nil

### UnsetEasypostTrackerId
`func (o *UpdateOrderResponse) UnsetEasypostTrackerId()`

UnsetEasypostTrackerId ensures that no value is present for EasypostTrackerId, not even an explicit nil
### GetEasypostRateId

`func (o *UpdateOrderResponse) GetEasypostRateId() string`

GetEasypostRateId returns the EasypostRateId field if non-nil, zero value otherwise.

### GetEasypostRateIdOk

`func (o *UpdateOrderResponse) GetEasypostRateIdOk() (*string, bool)`

GetEasypostRateIdOk returns a tuple with the EasypostRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasypostRateId

`func (o *UpdateOrderResponse) SetEasypostRateId(v string)`

SetEasypostRateId sets EasypostRateId field to given value.

### HasEasypostRateId

`func (o *UpdateOrderResponse) HasEasypostRateId() bool`

HasEasypostRateId returns a boolean if a field has been set.

### SetEasypostRateIdNil

`func (o *UpdateOrderResponse) SetEasypostRateIdNil(b bool)`

 SetEasypostRateIdNil sets the value for EasypostRateId to be an explicit nil

### UnsetEasypostRateId
`func (o *UpdateOrderResponse) UnsetEasypostRateId()`

UnsetEasypostRateId ensures that no value is present for EasypostRateId, not even an explicit nil
### GetShippingLabelUrl

`func (o *UpdateOrderResponse) GetShippingLabelUrl() string`

GetShippingLabelUrl returns the ShippingLabelUrl field if non-nil, zero value otherwise.

### GetShippingLabelUrlOk

`func (o *UpdateOrderResponse) GetShippingLabelUrlOk() (*string, bool)`

GetShippingLabelUrlOk returns a tuple with the ShippingLabelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelUrl

`func (o *UpdateOrderResponse) SetShippingLabelUrl(v string)`

SetShippingLabelUrl sets ShippingLabelUrl field to given value.

### HasShippingLabelUrl

`func (o *UpdateOrderResponse) HasShippingLabelUrl() bool`

HasShippingLabelUrl returns a boolean if a field has been set.

### SetShippingLabelUrlNil

`func (o *UpdateOrderResponse) SetShippingLabelUrlNil(b bool)`

 SetShippingLabelUrlNil sets the value for ShippingLabelUrl to be an explicit nil

### UnsetShippingLabelUrl
`func (o *UpdateOrderResponse) UnsetShippingLabelUrl()`

UnsetShippingLabelUrl ensures that no value is present for ShippingLabelUrl, not even an explicit nil
### GetLabelPurchasedAt

`func (o *UpdateOrderResponse) GetLabelPurchasedAt() time.Time`

GetLabelPurchasedAt returns the LabelPurchasedAt field if non-nil, zero value otherwise.

### GetLabelPurchasedAtOk

`func (o *UpdateOrderResponse) GetLabelPurchasedAtOk() (*time.Time, bool)`

GetLabelPurchasedAtOk returns a tuple with the LabelPurchasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPurchasedAt

`func (o *UpdateOrderResponse) SetLabelPurchasedAt(v time.Time)`

SetLabelPurchasedAt sets LabelPurchasedAt field to given value.

### HasLabelPurchasedAt

`func (o *UpdateOrderResponse) HasLabelPurchasedAt() bool`

HasLabelPurchasedAt returns a boolean if a field has been set.

### SetLabelPurchasedAtNil

`func (o *UpdateOrderResponse) SetLabelPurchasedAtNil(b bool)`

 SetLabelPurchasedAtNil sets the value for LabelPurchasedAt to be an explicit nil

### UnsetLabelPurchasedAt
`func (o *UpdateOrderResponse) UnsetLabelPurchasedAt()`

UnsetLabelPurchasedAt ensures that no value is present for LabelPurchasedAt, not even an explicit nil
### GetShippedAt

`func (o *UpdateOrderResponse) GetShippedAt() time.Time`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *UpdateOrderResponse) GetShippedAtOk() (*time.Time, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *UpdateOrderResponse) SetShippedAt(v time.Time)`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *UpdateOrderResponse) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### SetShippedAtNil

`func (o *UpdateOrderResponse) SetShippedAtNil(b bool)`

 SetShippedAtNil sets the value for ShippedAt to be an explicit nil

### UnsetShippedAt
`func (o *UpdateOrderResponse) UnsetShippedAt()`

UnsetShippedAt ensures that no value is present for ShippedAt, not even an explicit nil
### GetEstimatedDelivery

`func (o *UpdateOrderResponse) GetEstimatedDelivery() string`

GetEstimatedDelivery returns the EstimatedDelivery field if non-nil, zero value otherwise.

### GetEstimatedDeliveryOk

`func (o *UpdateOrderResponse) GetEstimatedDeliveryOk() (*string, bool)`

GetEstimatedDeliveryOk returns a tuple with the EstimatedDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDelivery

`func (o *UpdateOrderResponse) SetEstimatedDelivery(v string)`

SetEstimatedDelivery sets EstimatedDelivery field to given value.

### HasEstimatedDelivery

`func (o *UpdateOrderResponse) HasEstimatedDelivery() bool`

HasEstimatedDelivery returns a boolean if a field has been set.

### SetEstimatedDeliveryNil

`func (o *UpdateOrderResponse) SetEstimatedDeliveryNil(b bool)`

 SetEstimatedDeliveryNil sets the value for EstimatedDelivery to be an explicit nil

### UnsetEstimatedDelivery
`func (o *UpdateOrderResponse) UnsetEstimatedDelivery()`

UnsetEstimatedDelivery ensures that no value is present for EstimatedDelivery, not even an explicit nil
### GetDeliveredAt

`func (o *UpdateOrderResponse) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *UpdateOrderResponse) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *UpdateOrderResponse) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *UpdateOrderResponse) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### SetDeliveredAtNil

`func (o *UpdateOrderResponse) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *UpdateOrderResponse) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil
### GetCarrierStatus

`func (o *UpdateOrderResponse) GetCarrierStatus() string`

GetCarrierStatus returns the CarrierStatus field if non-nil, zero value otherwise.

### GetCarrierStatusOk

`func (o *UpdateOrderResponse) GetCarrierStatusOk() (*string, bool)`

GetCarrierStatusOk returns a tuple with the CarrierStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierStatus

`func (o *UpdateOrderResponse) SetCarrierStatus(v string)`

SetCarrierStatus sets CarrierStatus field to given value.

### HasCarrierStatus

`func (o *UpdateOrderResponse) HasCarrierStatus() bool`

HasCarrierStatus returns a boolean if a field has been set.

### SetCarrierStatusNil

`func (o *UpdateOrderResponse) SetCarrierStatusNil(b bool)`

 SetCarrierStatusNil sets the value for CarrierStatus to be an explicit nil

### UnsetCarrierStatus
`func (o *UpdateOrderResponse) UnsetCarrierStatus()`

UnsetCarrierStatus ensures that no value is present for CarrierStatus, not even an explicit nil
### GetCarrierStatusDetail

`func (o *UpdateOrderResponse) GetCarrierStatusDetail() string`

GetCarrierStatusDetail returns the CarrierStatusDetail field if non-nil, zero value otherwise.

### GetCarrierStatusDetailOk

`func (o *UpdateOrderResponse) GetCarrierStatusDetailOk() (*string, bool)`

GetCarrierStatusDetailOk returns a tuple with the CarrierStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierStatusDetail

`func (o *UpdateOrderResponse) SetCarrierStatusDetail(v string)`

SetCarrierStatusDetail sets CarrierStatusDetail field to given value.

### HasCarrierStatusDetail

`func (o *UpdateOrderResponse) HasCarrierStatusDetail() bool`

HasCarrierStatusDetail returns a boolean if a field has been set.

### SetCarrierStatusDetailNil

`func (o *UpdateOrderResponse) SetCarrierStatusDetailNil(b bool)`

 SetCarrierStatusDetailNil sets the value for CarrierStatusDetail to be an explicit nil

### UnsetCarrierStatusDetail
`func (o *UpdateOrderResponse) UnsetCarrierStatusDetail()`

UnsetCarrierStatusDetail ensures that no value is present for CarrierStatusDetail, not even an explicit nil
### GetTrackingHistory

`func (o *UpdateOrderResponse) GetTrackingHistory() []ListOrdersItemTrackingHistory`

GetTrackingHistory returns the TrackingHistory field if non-nil, zero value otherwise.

### GetTrackingHistoryOk

`func (o *UpdateOrderResponse) GetTrackingHistoryOk() (*[]ListOrdersItemTrackingHistory, bool)`

GetTrackingHistoryOk returns a tuple with the TrackingHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingHistory

`func (o *UpdateOrderResponse) SetTrackingHistory(v []ListOrdersItemTrackingHistory)`

SetTrackingHistory sets TrackingHistory field to given value.

### HasTrackingHistory

`func (o *UpdateOrderResponse) HasTrackingHistory() bool`

HasTrackingHistory returns a boolean if a field has been set.

### SetTrackingHistoryNil

`func (o *UpdateOrderResponse) SetTrackingHistoryNil(b bool)`

 SetTrackingHistoryNil sets the value for TrackingHistory to be an explicit nil

### UnsetTrackingHistory
`func (o *UpdateOrderResponse) UnsetTrackingHistory()`

UnsetTrackingHistory ensures that no value is present for TrackingHistory, not even an explicit nil
### GetDeliveryLocation

`func (o *UpdateOrderResponse) GetDeliveryLocation() string`

GetDeliveryLocation returns the DeliveryLocation field if non-nil, zero value otherwise.

### GetDeliveryLocationOk

`func (o *UpdateOrderResponse) GetDeliveryLocationOk() (*string, bool)`

GetDeliveryLocationOk returns a tuple with the DeliveryLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLocation

`func (o *UpdateOrderResponse) SetDeliveryLocation(v string)`

SetDeliveryLocation sets DeliveryLocation field to given value.

### HasDeliveryLocation

`func (o *UpdateOrderResponse) HasDeliveryLocation() bool`

HasDeliveryLocation returns a boolean if a field has been set.

### SetDeliveryLocationNil

`func (o *UpdateOrderResponse) SetDeliveryLocationNil(b bool)`

 SetDeliveryLocationNil sets the value for DeliveryLocation to be an explicit nil

### UnsetDeliveryLocation
`func (o *UpdateOrderResponse) UnsetDeliveryLocation()`

UnsetDeliveryLocation ensures that no value is present for DeliveryLocation, not even an explicit nil
### GetDeliverySignature

`func (o *UpdateOrderResponse) GetDeliverySignature() string`

GetDeliverySignature returns the DeliverySignature field if non-nil, zero value otherwise.

### GetDeliverySignatureOk

`func (o *UpdateOrderResponse) GetDeliverySignatureOk() (*string, bool)`

GetDeliverySignatureOk returns a tuple with the DeliverySignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverySignature

`func (o *UpdateOrderResponse) SetDeliverySignature(v string)`

SetDeliverySignature sets DeliverySignature field to given value.

### HasDeliverySignature

`func (o *UpdateOrderResponse) HasDeliverySignature() bool`

HasDeliverySignature returns a boolean if a field has been set.

### SetDeliverySignatureNil

`func (o *UpdateOrderResponse) SetDeliverySignatureNil(b bool)`

 SetDeliverySignatureNil sets the value for DeliverySignature to be an explicit nil

### UnsetDeliverySignature
`func (o *UpdateOrderResponse) UnsetDeliverySignature()`

UnsetDeliverySignature ensures that no value is present for DeliverySignature, not even an explicit nil
### GetTrackingSubmittedAt

`func (o *UpdateOrderResponse) GetTrackingSubmittedAt() time.Time`

GetTrackingSubmittedAt returns the TrackingSubmittedAt field if non-nil, zero value otherwise.

### GetTrackingSubmittedAtOk

`func (o *UpdateOrderResponse) GetTrackingSubmittedAtOk() (*time.Time, bool)`

GetTrackingSubmittedAtOk returns a tuple with the TrackingSubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingSubmittedAt

`func (o *UpdateOrderResponse) SetTrackingSubmittedAt(v time.Time)`

SetTrackingSubmittedAt sets TrackingSubmittedAt field to given value.

### HasTrackingSubmittedAt

`func (o *UpdateOrderResponse) HasTrackingSubmittedAt() bool`

HasTrackingSubmittedAt returns a boolean if a field has been set.

### SetTrackingSubmittedAtNil

`func (o *UpdateOrderResponse) SetTrackingSubmittedAtNil(b bool)`

 SetTrackingSubmittedAtNil sets the value for TrackingSubmittedAt to be an explicit nil

### UnsetTrackingSubmittedAt
`func (o *UpdateOrderResponse) UnsetTrackingSubmittedAt()`

UnsetTrackingSubmittedAt ensures that no value is present for TrackingSubmittedAt, not even an explicit nil
### GetTrackingSubmitStatus

`func (o *UpdateOrderResponse) GetTrackingSubmitStatus() string`

GetTrackingSubmitStatus returns the TrackingSubmitStatus field if non-nil, zero value otherwise.

### GetTrackingSubmitStatusOk

`func (o *UpdateOrderResponse) GetTrackingSubmitStatusOk() (*string, bool)`

GetTrackingSubmitStatusOk returns a tuple with the TrackingSubmitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingSubmitStatus

`func (o *UpdateOrderResponse) SetTrackingSubmitStatus(v string)`

SetTrackingSubmitStatus sets TrackingSubmitStatus field to given value.

### HasTrackingSubmitStatus

`func (o *UpdateOrderResponse) HasTrackingSubmitStatus() bool`

HasTrackingSubmitStatus returns a boolean if a field has been set.

### SetTrackingSubmitStatusNil

`func (o *UpdateOrderResponse) SetTrackingSubmitStatusNil(b bool)`

 SetTrackingSubmitStatusNil sets the value for TrackingSubmitStatus to be an explicit nil

### UnsetTrackingSubmitStatus
`func (o *UpdateOrderResponse) UnsetTrackingSubmitStatus()`

UnsetTrackingSubmitStatus ensures that no value is present for TrackingSubmitStatus, not even an explicit nil
### GetLabelCost

`func (o *UpdateOrderResponse) GetLabelCost() string`

GetLabelCost returns the LabelCost field if non-nil, zero value otherwise.

### GetLabelCostOk

`func (o *UpdateOrderResponse) GetLabelCostOk() (*string, bool)`

GetLabelCostOk returns a tuple with the LabelCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelCost

`func (o *UpdateOrderResponse) SetLabelCost(v string)`

SetLabelCost sets LabelCost field to given value.

### HasLabelCost

`func (o *UpdateOrderResponse) HasLabelCost() bool`

HasLabelCost returns a boolean if a field has been set.

### SetLabelCostNil

`func (o *UpdateOrderResponse) SetLabelCostNil(b bool)`

 SetLabelCostNil sets the value for LabelCost to be an explicit nil

### UnsetLabelCost
`func (o *UpdateOrderResponse) UnsetLabelCost()`

UnsetLabelCost ensures that no value is present for LabelCost, not even an explicit nil
### GetCostOfGoods

`func (o *UpdateOrderResponse) GetCostOfGoods() string`

GetCostOfGoods returns the CostOfGoods field if non-nil, zero value otherwise.

### GetCostOfGoodsOk

`func (o *UpdateOrderResponse) GetCostOfGoodsOk() (*string, bool)`

GetCostOfGoodsOk returns a tuple with the CostOfGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOfGoods

`func (o *UpdateOrderResponse) SetCostOfGoods(v string)`

SetCostOfGoods sets CostOfGoods field to given value.

### HasCostOfGoods

`func (o *UpdateOrderResponse) HasCostOfGoods() bool`

HasCostOfGoods returns a boolean if a field has been set.

### SetCostOfGoodsNil

`func (o *UpdateOrderResponse) SetCostOfGoodsNil(b bool)`

 SetCostOfGoodsNil sets the value for CostOfGoods to be an explicit nil

### UnsetCostOfGoods
`func (o *UpdateOrderResponse) UnsetCostOfGoods()`

UnsetCostOfGoods ensures that no value is present for CostOfGoods, not even an explicit nil
### GetRequestedCarrier

`func (o *UpdateOrderResponse) GetRequestedCarrier() string`

GetRequestedCarrier returns the RequestedCarrier field if non-nil, zero value otherwise.

### GetRequestedCarrierOk

`func (o *UpdateOrderResponse) GetRequestedCarrierOk() (*string, bool)`

GetRequestedCarrierOk returns a tuple with the RequestedCarrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCarrier

`func (o *UpdateOrderResponse) SetRequestedCarrier(v string)`

SetRequestedCarrier sets RequestedCarrier field to given value.

### HasRequestedCarrier

`func (o *UpdateOrderResponse) HasRequestedCarrier() bool`

HasRequestedCarrier returns a boolean if a field has been set.

### SetRequestedCarrierNil

`func (o *UpdateOrderResponse) SetRequestedCarrierNil(b bool)`

 SetRequestedCarrierNil sets the value for RequestedCarrier to be an explicit nil

### UnsetRequestedCarrier
`func (o *UpdateOrderResponse) UnsetRequestedCarrier()`

UnsetRequestedCarrier ensures that no value is present for RequestedCarrier, not even an explicit nil
### GetRequestedService

`func (o *UpdateOrderResponse) GetRequestedService() string`

GetRequestedService returns the RequestedService field if non-nil, zero value otherwise.

### GetRequestedServiceOk

`func (o *UpdateOrderResponse) GetRequestedServiceOk() (*string, bool)`

GetRequestedServiceOk returns a tuple with the RequestedService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedService

`func (o *UpdateOrderResponse) SetRequestedService(v string)`

SetRequestedService sets RequestedService field to given value.

### HasRequestedService

`func (o *UpdateOrderResponse) HasRequestedService() bool`

HasRequestedService returns a boolean if a field has been set.

### SetRequestedServiceNil

`func (o *UpdateOrderResponse) SetRequestedServiceNil(b bool)`

 SetRequestedServiceNil sets the value for RequestedService to be an explicit nil

### UnsetRequestedService
`func (o *UpdateOrderResponse) UnsetRequestedService()`

UnsetRequestedService ensures that no value is present for RequestedService, not even an explicit nil
### GetShipByAt

`func (o *UpdateOrderResponse) GetShipByAt() time.Time`

GetShipByAt returns the ShipByAt field if non-nil, zero value otherwise.

### GetShipByAtOk

`func (o *UpdateOrderResponse) GetShipByAtOk() (*time.Time, bool)`

GetShipByAtOk returns a tuple with the ShipByAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipByAt

`func (o *UpdateOrderResponse) SetShipByAt(v time.Time)`

SetShipByAt sets ShipByAt field to given value.

### HasShipByAt

`func (o *UpdateOrderResponse) HasShipByAt() bool`

HasShipByAt returns a boolean if a field has been set.

### SetShipByAtNil

`func (o *UpdateOrderResponse) SetShipByAtNil(b bool)`

 SetShipByAtNil sets the value for ShipByAt to be an explicit nil

### UnsetShipByAt
`func (o *UpdateOrderResponse) UnsetShipByAt()`

UnsetShipByAt ensures that no value is present for ShipByAt, not even an explicit nil
### GetShipByAlertedAt

`func (o *UpdateOrderResponse) GetShipByAlertedAt() time.Time`

GetShipByAlertedAt returns the ShipByAlertedAt field if non-nil, zero value otherwise.

### GetShipByAlertedAtOk

`func (o *UpdateOrderResponse) GetShipByAlertedAtOk() (*time.Time, bool)`

GetShipByAlertedAtOk returns a tuple with the ShipByAlertedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipByAlertedAt

`func (o *UpdateOrderResponse) SetShipByAlertedAt(v time.Time)`

SetShipByAlertedAt sets ShipByAlertedAt field to given value.

### HasShipByAlertedAt

`func (o *UpdateOrderResponse) HasShipByAlertedAt() bool`

HasShipByAlertedAt returns a boolean if a field has been set.

### SetShipByAlertedAtNil

`func (o *UpdateOrderResponse) SetShipByAlertedAtNil(b bool)`

 SetShipByAlertedAtNil sets the value for ShipByAlertedAt to be an explicit nil

### UnsetShipByAlertedAt
`func (o *UpdateOrderResponse) UnsetShipByAlertedAt()`

UnsetShipByAlertedAt ensures that no value is present for ShipByAlertedAt, not even an explicit nil
### GetOversoldBy

`func (o *UpdateOrderResponse) GetOversoldBy() float32`

GetOversoldBy returns the OversoldBy field if non-nil, zero value otherwise.

### GetOversoldByOk

`func (o *UpdateOrderResponse) GetOversoldByOk() (*float32, bool)`

GetOversoldByOk returns a tuple with the OversoldBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversoldBy

`func (o *UpdateOrderResponse) SetOversoldBy(v float32)`

SetOversoldBy sets OversoldBy field to given value.


### GetLastStatusCheckAt

`func (o *UpdateOrderResponse) GetLastStatusCheckAt() time.Time`

GetLastStatusCheckAt returns the LastStatusCheckAt field if non-nil, zero value otherwise.

### GetLastStatusCheckAtOk

`func (o *UpdateOrderResponse) GetLastStatusCheckAtOk() (*time.Time, bool)`

GetLastStatusCheckAtOk returns a tuple with the LastStatusCheckAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusCheckAt

`func (o *UpdateOrderResponse) SetLastStatusCheckAt(v time.Time)`

SetLastStatusCheckAt sets LastStatusCheckAt field to given value.

### HasLastStatusCheckAt

`func (o *UpdateOrderResponse) HasLastStatusCheckAt() bool`

HasLastStatusCheckAt returns a boolean if a field has been set.

### SetLastStatusCheckAtNil

`func (o *UpdateOrderResponse) SetLastStatusCheckAtNil(b bool)`

 SetLastStatusCheckAtNil sets the value for LastStatusCheckAt to be an explicit nil

### UnsetLastStatusCheckAt
`func (o *UpdateOrderResponse) UnsetLastStatusCheckAt()`

UnsetLastStatusCheckAt ensures that no value is present for LastStatusCheckAt, not even an explicit nil
### GetLastChatCheckAt

`func (o *UpdateOrderResponse) GetLastChatCheckAt() time.Time`

GetLastChatCheckAt returns the LastChatCheckAt field if non-nil, zero value otherwise.

### GetLastChatCheckAtOk

`func (o *UpdateOrderResponse) GetLastChatCheckAtOk() (*time.Time, bool)`

GetLastChatCheckAtOk returns a tuple with the LastChatCheckAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChatCheckAt

`func (o *UpdateOrderResponse) SetLastChatCheckAt(v time.Time)`

SetLastChatCheckAt sets LastChatCheckAt field to given value.

### HasLastChatCheckAt

`func (o *UpdateOrderResponse) HasLastChatCheckAt() bool`

HasLastChatCheckAt returns a boolean if a field has been set.

### SetLastChatCheckAtNil

`func (o *UpdateOrderResponse) SetLastChatCheckAtNil(b bool)`

 SetLastChatCheckAtNil sets the value for LastChatCheckAt to be an explicit nil

### UnsetLastChatCheckAt
`func (o *UpdateOrderResponse) UnsetLastChatCheckAt()`

UnsetLastChatCheckAt ensures that no value is present for LastChatCheckAt, not even an explicit nil
### GetIsDisputed

`func (o *UpdateOrderResponse) GetIsDisputed() bool`

GetIsDisputed returns the IsDisputed field if non-nil, zero value otherwise.

### GetIsDisputedOk

`func (o *UpdateOrderResponse) GetIsDisputedOk() (*bool, bool)`

GetIsDisputedOk returns a tuple with the IsDisputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisputed

`func (o *UpdateOrderResponse) SetIsDisputed(v bool)`

SetIsDisputed sets IsDisputed field to given value.


### GetDisputeReason

`func (o *UpdateOrderResponse) GetDisputeReason() string`

GetDisputeReason returns the DisputeReason field if non-nil, zero value otherwise.

### GetDisputeReasonOk

`func (o *UpdateOrderResponse) GetDisputeReasonOk() (*string, bool)`

GetDisputeReasonOk returns a tuple with the DisputeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeReason

`func (o *UpdateOrderResponse) SetDisputeReason(v string)`

SetDisputeReason sets DisputeReason field to given value.

### HasDisputeReason

`func (o *UpdateOrderResponse) HasDisputeReason() bool`

HasDisputeReason returns a boolean if a field has been set.

### SetDisputeReasonNil

`func (o *UpdateOrderResponse) SetDisputeReasonNil(b bool)`

 SetDisputeReasonNil sets the value for DisputeReason to be an explicit nil

### UnsetDisputeReason
`func (o *UpdateOrderResponse) UnsetDisputeReason()`

UnsetDisputeReason ensures that no value is present for DisputeReason, not even an explicit nil
### GetDisputePlatformCaseId

`func (o *UpdateOrderResponse) GetDisputePlatformCaseId() string`

GetDisputePlatformCaseId returns the DisputePlatformCaseId field if non-nil, zero value otherwise.

### GetDisputePlatformCaseIdOk

`func (o *UpdateOrderResponse) GetDisputePlatformCaseIdOk() (*string, bool)`

GetDisputePlatformCaseIdOk returns a tuple with the DisputePlatformCaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputePlatformCaseId

`func (o *UpdateOrderResponse) SetDisputePlatformCaseId(v string)`

SetDisputePlatformCaseId sets DisputePlatformCaseId field to given value.

### HasDisputePlatformCaseId

`func (o *UpdateOrderResponse) HasDisputePlatformCaseId() bool`

HasDisputePlatformCaseId returns a boolean if a field has been set.

### SetDisputePlatformCaseIdNil

`func (o *UpdateOrderResponse) SetDisputePlatformCaseIdNil(b bool)`

 SetDisputePlatformCaseIdNil sets the value for DisputePlatformCaseId to be an explicit nil

### UnsetDisputePlatformCaseId
`func (o *UpdateOrderResponse) UnsetDisputePlatformCaseId()`

UnsetDisputePlatformCaseId ensures that no value is present for DisputePlatformCaseId, not even an explicit nil
### GetDisputeResolvedAt

`func (o *UpdateOrderResponse) GetDisputeResolvedAt() time.Time`

GetDisputeResolvedAt returns the DisputeResolvedAt field if non-nil, zero value otherwise.

### GetDisputeResolvedAtOk

`func (o *UpdateOrderResponse) GetDisputeResolvedAtOk() (*time.Time, bool)`

GetDisputeResolvedAtOk returns a tuple with the DisputeResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeResolvedAt

`func (o *UpdateOrderResponse) SetDisputeResolvedAt(v time.Time)`

SetDisputeResolvedAt sets DisputeResolvedAt field to given value.

### HasDisputeResolvedAt

`func (o *UpdateOrderResponse) HasDisputeResolvedAt() bool`

HasDisputeResolvedAt returns a boolean if a field has been set.

### SetDisputeResolvedAtNil

`func (o *UpdateOrderResponse) SetDisputeResolvedAtNil(b bool)`

 SetDisputeResolvedAtNil sets the value for DisputeResolvedAt to be an explicit nil

### UnsetDisputeResolvedAt
`func (o *UpdateOrderResponse) UnsetDisputeResolvedAt()`

UnsetDisputeResolvedAt ensures that no value is present for DisputeResolvedAt, not even an explicit nil
### GetRefundAmount

`func (o *UpdateOrderResponse) GetRefundAmount() string`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *UpdateOrderResponse) GetRefundAmountOk() (*string, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *UpdateOrderResponse) SetRefundAmount(v string)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *UpdateOrderResponse) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### SetRefundAmountNil

`func (o *UpdateOrderResponse) SetRefundAmountNil(b bool)`

 SetRefundAmountNil sets the value for RefundAmount to be an explicit nil

### UnsetRefundAmount
`func (o *UpdateOrderResponse) UnsetRefundAmount()`

UnsetRefundAmount ensures that no value is present for RefundAmount, not even an explicit nil
### GetRefundReason

`func (o *UpdateOrderResponse) GetRefundReason() string`

GetRefundReason returns the RefundReason field if non-nil, zero value otherwise.

### GetRefundReasonOk

`func (o *UpdateOrderResponse) GetRefundReasonOk() (*string, bool)`

GetRefundReasonOk returns a tuple with the RefundReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundReason

`func (o *UpdateOrderResponse) SetRefundReason(v string)`

SetRefundReason sets RefundReason field to given value.

### HasRefundReason

`func (o *UpdateOrderResponse) HasRefundReason() bool`

HasRefundReason returns a boolean if a field has been set.

### SetRefundReasonNil

`func (o *UpdateOrderResponse) SetRefundReasonNil(b bool)`

 SetRefundReasonNil sets the value for RefundReason to be an explicit nil

### UnsetRefundReason
`func (o *UpdateOrderResponse) UnsetRefundReason()`

UnsetRefundReason ensures that no value is present for RefundReason, not even an explicit nil
### GetRefundPlatformId

`func (o *UpdateOrderResponse) GetRefundPlatformId() string`

GetRefundPlatformId returns the RefundPlatformId field if non-nil, zero value otherwise.

### GetRefundPlatformIdOk

`func (o *UpdateOrderResponse) GetRefundPlatformIdOk() (*string, bool)`

GetRefundPlatformIdOk returns a tuple with the RefundPlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundPlatformId

`func (o *UpdateOrderResponse) SetRefundPlatformId(v string)`

SetRefundPlatformId sets RefundPlatformId field to given value.

### HasRefundPlatformId

`func (o *UpdateOrderResponse) HasRefundPlatformId() bool`

HasRefundPlatformId returns a boolean if a field has been set.

### SetRefundPlatformIdNil

`func (o *UpdateOrderResponse) SetRefundPlatformIdNil(b bool)`

 SetRefundPlatformIdNil sets the value for RefundPlatformId to be an explicit nil

### UnsetRefundPlatformId
`func (o *UpdateOrderResponse) UnsetRefundPlatformId()`

UnsetRefundPlatformId ensures that no value is present for RefundPlatformId, not even an explicit nil
### GetRefundedAt

`func (o *UpdateOrderResponse) GetRefundedAt() time.Time`

GetRefundedAt returns the RefundedAt field if non-nil, zero value otherwise.

### GetRefundedAtOk

`func (o *UpdateOrderResponse) GetRefundedAtOk() (*time.Time, bool)`

GetRefundedAtOk returns a tuple with the RefundedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAt

`func (o *UpdateOrderResponse) SetRefundedAt(v time.Time)`

SetRefundedAt sets RefundedAt field to given value.

### HasRefundedAt

`func (o *UpdateOrderResponse) HasRefundedAt() bool`

HasRefundedAt returns a boolean if a field has been set.

### SetRefundedAtNil

`func (o *UpdateOrderResponse) SetRefundedAtNil(b bool)`

 SetRefundedAtNil sets the value for RefundedAt to be an explicit nil

### UnsetRefundedAt
`func (o *UpdateOrderResponse) UnsetRefundedAt()`

UnsetRefundedAt ensures that no value is present for RefundedAt, not even an explicit nil
### GetCancellationReason

`func (o *UpdateOrderResponse) GetCancellationReason() string`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *UpdateOrderResponse) GetCancellationReasonOk() (*string, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationReason

`func (o *UpdateOrderResponse) SetCancellationReason(v string)`

SetCancellationReason sets CancellationReason field to given value.

### HasCancellationReason

`func (o *UpdateOrderResponse) HasCancellationReason() bool`

HasCancellationReason returns a boolean if a field has been set.

### SetCancellationReasonNil

`func (o *UpdateOrderResponse) SetCancellationReasonNil(b bool)`

 SetCancellationReasonNil sets the value for CancellationReason to be an explicit nil

### UnsetCancellationReason
`func (o *UpdateOrderResponse) UnsetCancellationReason()`

UnsetCancellationReason ensures that no value is present for CancellationReason, not even an explicit nil
### GetArrivalConditionRequestedAt

`func (o *UpdateOrderResponse) GetArrivalConditionRequestedAt() time.Time`

GetArrivalConditionRequestedAt returns the ArrivalConditionRequestedAt field if non-nil, zero value otherwise.

### GetArrivalConditionRequestedAtOk

`func (o *UpdateOrderResponse) GetArrivalConditionRequestedAtOk() (*time.Time, bool)`

GetArrivalConditionRequestedAtOk returns a tuple with the ArrivalConditionRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalConditionRequestedAt

`func (o *UpdateOrderResponse) SetArrivalConditionRequestedAt(v time.Time)`

SetArrivalConditionRequestedAt sets ArrivalConditionRequestedAt field to given value.

### HasArrivalConditionRequestedAt

`func (o *UpdateOrderResponse) HasArrivalConditionRequestedAt() bool`

HasArrivalConditionRequestedAt returns a boolean if a field has been set.

### SetArrivalConditionRequestedAtNil

`func (o *UpdateOrderResponse) SetArrivalConditionRequestedAtNil(b bool)`

 SetArrivalConditionRequestedAtNil sets the value for ArrivalConditionRequestedAt to be an explicit nil

### UnsetArrivalConditionRequestedAt
`func (o *UpdateOrderResponse) UnsetArrivalConditionRequestedAt()`

UnsetArrivalConditionRequestedAt ensures that no value is present for ArrivalConditionRequestedAt, not even an explicit nil
### GetArrivalConditionSubmittedAt

`func (o *UpdateOrderResponse) GetArrivalConditionSubmittedAt() time.Time`

GetArrivalConditionSubmittedAt returns the ArrivalConditionSubmittedAt field if non-nil, zero value otherwise.

### GetArrivalConditionSubmittedAtOk

`func (o *UpdateOrderResponse) GetArrivalConditionSubmittedAtOk() (*time.Time, bool)`

GetArrivalConditionSubmittedAtOk returns a tuple with the ArrivalConditionSubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalConditionSubmittedAt

`func (o *UpdateOrderResponse) SetArrivalConditionSubmittedAt(v time.Time)`

SetArrivalConditionSubmittedAt sets ArrivalConditionSubmittedAt field to given value.

### HasArrivalConditionSubmittedAt

`func (o *UpdateOrderResponse) HasArrivalConditionSubmittedAt() bool`

HasArrivalConditionSubmittedAt returns a boolean if a field has been set.

### SetArrivalConditionSubmittedAtNil

`func (o *UpdateOrderResponse) SetArrivalConditionSubmittedAtNil(b bool)`

 SetArrivalConditionSubmittedAtNil sets the value for ArrivalConditionSubmittedAt to be an explicit nil

### UnsetArrivalConditionSubmittedAt
`func (o *UpdateOrderResponse) UnsetArrivalConditionSubmittedAt()`

UnsetArrivalConditionSubmittedAt ensures that no value is present for ArrivalConditionSubmittedAt, not even an explicit nil
### GetArrivalConditionDeclinedAt

`func (o *UpdateOrderResponse) GetArrivalConditionDeclinedAt() time.Time`

GetArrivalConditionDeclinedAt returns the ArrivalConditionDeclinedAt field if non-nil, zero value otherwise.

### GetArrivalConditionDeclinedAtOk

`func (o *UpdateOrderResponse) GetArrivalConditionDeclinedAtOk() (*time.Time, bool)`

GetArrivalConditionDeclinedAtOk returns a tuple with the ArrivalConditionDeclinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalConditionDeclinedAt

`func (o *UpdateOrderResponse) SetArrivalConditionDeclinedAt(v time.Time)`

SetArrivalConditionDeclinedAt sets ArrivalConditionDeclinedAt field to given value.

### HasArrivalConditionDeclinedAt

`func (o *UpdateOrderResponse) HasArrivalConditionDeclinedAt() bool`

HasArrivalConditionDeclinedAt returns a boolean if a field has been set.

### SetArrivalConditionDeclinedAtNil

`func (o *UpdateOrderResponse) SetArrivalConditionDeclinedAtNil(b bool)`

 SetArrivalConditionDeclinedAtNil sets the value for ArrivalConditionDeclinedAt to be an explicit nil

### UnsetArrivalConditionDeclinedAt
`func (o *UpdateOrderResponse) UnsetArrivalConditionDeclinedAt()`

UnsetArrivalConditionDeclinedAt ensures that no value is present for ArrivalConditionDeclinedAt, not even an explicit nil
### GetDeliveryPhotoUrl

`func (o *UpdateOrderResponse) GetDeliveryPhotoUrl() string`

GetDeliveryPhotoUrl returns the DeliveryPhotoUrl field if non-nil, zero value otherwise.

### GetDeliveryPhotoUrlOk

`func (o *UpdateOrderResponse) GetDeliveryPhotoUrlOk() (*string, bool)`

GetDeliveryPhotoUrlOk returns a tuple with the DeliveryPhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPhotoUrl

`func (o *UpdateOrderResponse) SetDeliveryPhotoUrl(v string)`

SetDeliveryPhotoUrl sets DeliveryPhotoUrl field to given value.

### HasDeliveryPhotoUrl

`func (o *UpdateOrderResponse) HasDeliveryPhotoUrl() bool`

HasDeliveryPhotoUrl returns a boolean if a field has been set.

### SetDeliveryPhotoUrlNil

`func (o *UpdateOrderResponse) SetDeliveryPhotoUrlNil(b bool)`

 SetDeliveryPhotoUrlNil sets the value for DeliveryPhotoUrl to be an explicit nil

### UnsetDeliveryPhotoUrl
`func (o *UpdateOrderResponse) UnsetDeliveryPhotoUrl()`

UnsetDeliveryPhotoUrl ensures that no value is present for DeliveryPhotoUrl, not even an explicit nil
### GetPurchaseOrderRef

`func (o *UpdateOrderResponse) GetPurchaseOrderRef() string`

GetPurchaseOrderRef returns the PurchaseOrderRef field if non-nil, zero value otherwise.

### GetPurchaseOrderRefOk

`func (o *UpdateOrderResponse) GetPurchaseOrderRefOk() (*string, bool)`

GetPurchaseOrderRefOk returns a tuple with the PurchaseOrderRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderRef

`func (o *UpdateOrderResponse) SetPurchaseOrderRef(v string)`

SetPurchaseOrderRef sets PurchaseOrderRef field to given value.

### HasPurchaseOrderRef

`func (o *UpdateOrderResponse) HasPurchaseOrderRef() bool`

HasPurchaseOrderRef returns a boolean if a field has been set.

### SetPurchaseOrderRefNil

`func (o *UpdateOrderResponse) SetPurchaseOrderRefNil(b bool)`

 SetPurchaseOrderRefNil sets the value for PurchaseOrderRef to be an explicit nil

### UnsetPurchaseOrderRef
`func (o *UpdateOrderResponse) UnsetPurchaseOrderRef()`

UnsetPurchaseOrderRef ensures that no value is present for PurchaseOrderRef, not even an explicit nil
### GetArrivalConditionPhotos

`func (o *UpdateOrderResponse) GetArrivalConditionPhotos() []string`

GetArrivalConditionPhotos returns the ArrivalConditionPhotos field if non-nil, zero value otherwise.

### GetArrivalConditionPhotosOk

`func (o *UpdateOrderResponse) GetArrivalConditionPhotosOk() (*[]string, bool)`

GetArrivalConditionPhotosOk returns a tuple with the ArrivalConditionPhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalConditionPhotos

`func (o *UpdateOrderResponse) SetArrivalConditionPhotos(v []string)`

SetArrivalConditionPhotos sets ArrivalConditionPhotos field to given value.


### GetBuyerEmail

`func (o *UpdateOrderResponse) GetBuyerEmail() string`

GetBuyerEmail returns the BuyerEmail field if non-nil, zero value otherwise.

### GetBuyerEmailOk

`func (o *UpdateOrderResponse) GetBuyerEmailOk() (*string, bool)`

GetBuyerEmailOk returns a tuple with the BuyerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerEmail

`func (o *UpdateOrderResponse) SetBuyerEmail(v string)`

SetBuyerEmail sets BuyerEmail field to given value.

### HasBuyerEmail

`func (o *UpdateOrderResponse) HasBuyerEmail() bool`

HasBuyerEmail returns a boolean if a field has been set.

### SetBuyerEmailNil

`func (o *UpdateOrderResponse) SetBuyerEmailNil(b bool)`

 SetBuyerEmailNil sets the value for BuyerEmail to be an explicit nil

### UnsetBuyerEmail
`func (o *UpdateOrderResponse) UnsetBuyerEmail()`

UnsetBuyerEmail ensures that no value is present for BuyerEmail, not even an explicit nil
### GetFulfillmentMethod

`func (o *UpdateOrderResponse) GetFulfillmentMethod() string`

GetFulfillmentMethod returns the FulfillmentMethod field if non-nil, zero value otherwise.

### GetFulfillmentMethodOk

`func (o *UpdateOrderResponse) GetFulfillmentMethodOk() (*string, bool)`

GetFulfillmentMethodOk returns a tuple with the FulfillmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentMethod

`func (o *UpdateOrderResponse) SetFulfillmentMethod(v string)`

SetFulfillmentMethod sets FulfillmentMethod field to given value.


### GetDeletedAt

`func (o *UpdateOrderResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *UpdateOrderResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *UpdateOrderResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *UpdateOrderResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *UpdateOrderResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *UpdateOrderResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


