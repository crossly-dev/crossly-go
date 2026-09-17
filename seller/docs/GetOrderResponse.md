# GetOrderResponse

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

### NewGetOrderResponse

`func NewGetOrderResponse(id string, createdAt time.Time, updatedAt time.Time, userId string, quantity float32, status string, platform string, salesChannel string, oversoldBy float32, isDisputed bool, arrivalConditionPhotos []string, fulfillmentMethod string, ) *GetOrderResponse`

NewGetOrderResponse instantiates a new GetOrderResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOrderResponseWithDefaults

`func NewGetOrderResponseWithDefaults() *GetOrderResponse`

NewGetOrderResponseWithDefaults instantiates a new GetOrderResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOrderResponse) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOrderResponse) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOrderResponse) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *GetOrderResponse) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *GetOrderResponse) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *GetOrderResponse) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *GetOrderResponse) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *GetOrderResponse) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *GetOrderResponse) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *GetOrderResponse) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *GetOrderResponse) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *GetOrderResponse) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetQuantity

`func (o *GetOrderResponse) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *GetOrderResponse) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *GetOrderResponse) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetNotes

`func (o *GetOrderResponse) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *GetOrderResponse) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *GetOrderResponse) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *GetOrderResponse) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *GetOrderResponse) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *GetOrderResponse) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetStatus

`func (o *GetOrderResponse) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetOrderResponse) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetOrderResponse) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlatform

`func (o *GetOrderResponse) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetOrderResponse) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetOrderResponse) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetCancelledAt

`func (o *GetOrderResponse) GetCancelledAt() time.Time`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *GetOrderResponse) GetCancelledAtOk() (*time.Time, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *GetOrderResponse) SetCancelledAt(v time.Time)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *GetOrderResponse) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### SetCancelledAtNil

`func (o *GetOrderResponse) SetCancelledAtNil(b bool)`

 SetCancelledAtNil sets the value for CancelledAt to be an explicit nil

### UnsetCancelledAt
`func (o *GetOrderResponse) UnsetCancelledAt()`

UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
### GetListingId

`func (o *GetOrderResponse) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *GetOrderResponse) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *GetOrderResponse) SetListingId(v string)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *GetOrderResponse) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### SetListingIdNil

`func (o *GetOrderResponse) SetListingIdNil(b bool)`

 SetListingIdNil sets the value for ListingId to be an explicit nil

### UnsetListingId
`func (o *GetOrderResponse) UnsetListingId()`

UnsetListingId ensures that no value is present for ListingId, not even an explicit nil
### GetInventoryItemId

`func (o *GetOrderResponse) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *GetOrderResponse) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *GetOrderResponse) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *GetOrderResponse) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### SetInventoryItemIdNil

`func (o *GetOrderResponse) SetInventoryItemIdNil(b bool)`

 SetInventoryItemIdNil sets the value for InventoryItemId to be an explicit nil

### UnsetInventoryItemId
`func (o *GetOrderResponse) UnsetInventoryItemId()`

UnsetInventoryItemId ensures that no value is present for InventoryItemId, not even an explicit nil
### GetPlatformListingId

`func (o *GetOrderResponse) GetPlatformListingId() string`

GetPlatformListingId returns the PlatformListingId field if non-nil, zero value otherwise.

### GetPlatformListingIdOk

`func (o *GetOrderResponse) GetPlatformListingIdOk() (*string, bool)`

GetPlatformListingIdOk returns a tuple with the PlatformListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformListingId

`func (o *GetOrderResponse) SetPlatformListingId(v string)`

SetPlatformListingId sets PlatformListingId field to given value.

### HasPlatformListingId

`func (o *GetOrderResponse) HasPlatformListingId() bool`

HasPlatformListingId returns a boolean if a field has been set.

### SetPlatformListingIdNil

`func (o *GetOrderResponse) SetPlatformListingIdNil(b bool)`

 SetPlatformListingIdNil sets the value for PlatformListingId to be an explicit nil

### UnsetPlatformListingId
`func (o *GetOrderResponse) UnsetPlatformListingId()`

UnsetPlatformListingId ensures that no value is present for PlatformListingId, not even an explicit nil
### GetHandlingTimeDays

`func (o *GetOrderResponse) GetHandlingTimeDays() float32`

GetHandlingTimeDays returns the HandlingTimeDays field if non-nil, zero value otherwise.

### GetHandlingTimeDaysOk

`func (o *GetOrderResponse) GetHandlingTimeDaysOk() (*float32, bool)`

GetHandlingTimeDaysOk returns a tuple with the HandlingTimeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingTimeDays

`func (o *GetOrderResponse) SetHandlingTimeDays(v float32)`

SetHandlingTimeDays sets HandlingTimeDays field to given value.

### HasHandlingTimeDays

`func (o *GetOrderResponse) HasHandlingTimeDays() bool`

HasHandlingTimeDays returns a boolean if a field has been set.

### SetHandlingTimeDaysNil

`func (o *GetOrderResponse) SetHandlingTimeDaysNil(b bool)`

 SetHandlingTimeDaysNil sets the value for HandlingTimeDays to be an explicit nil

### UnsetHandlingTimeDays
`func (o *GetOrderResponse) UnsetHandlingTimeDays()`

UnsetHandlingTimeDays ensures that no value is present for HandlingTimeDays, not even an explicit nil
### GetPlatformOrderId

`func (o *GetOrderResponse) GetPlatformOrderId() string`

GetPlatformOrderId returns the PlatformOrderId field if non-nil, zero value otherwise.

### GetPlatformOrderIdOk

`func (o *GetOrderResponse) GetPlatformOrderIdOk() (*string, bool)`

GetPlatformOrderIdOk returns a tuple with the PlatformOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformOrderId

`func (o *GetOrderResponse) SetPlatformOrderId(v string)`

SetPlatformOrderId sets PlatformOrderId field to given value.

### HasPlatformOrderId

`func (o *GetOrderResponse) HasPlatformOrderId() bool`

HasPlatformOrderId returns a boolean if a field has been set.

### SetPlatformOrderIdNil

`func (o *GetOrderResponse) SetPlatformOrderIdNil(b bool)`

 SetPlatformOrderIdNil sets the value for PlatformOrderId to be an explicit nil

### UnsetPlatformOrderId
`func (o *GetOrderResponse) UnsetPlatformOrderId()`

UnsetPlatformOrderId ensures that no value is present for PlatformOrderId, not even an explicit nil
### GetBuyerUsername

`func (o *GetOrderResponse) GetBuyerUsername() string`

GetBuyerUsername returns the BuyerUsername field if non-nil, zero value otherwise.

### GetBuyerUsernameOk

`func (o *GetOrderResponse) GetBuyerUsernameOk() (*string, bool)`

GetBuyerUsernameOk returns a tuple with the BuyerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerUsername

`func (o *GetOrderResponse) SetBuyerUsername(v string)`

SetBuyerUsername sets BuyerUsername field to given value.

### HasBuyerUsername

`func (o *GetOrderResponse) HasBuyerUsername() bool`

HasBuyerUsername returns a boolean if a field has been set.

### SetBuyerUsernameNil

`func (o *GetOrderResponse) SetBuyerUsernameNil(b bool)`

 SetBuyerUsernameNil sets the value for BuyerUsername to be an explicit nil

### UnsetBuyerUsername
`func (o *GetOrderResponse) UnsetBuyerUsername()`

UnsetBuyerUsername ensures that no value is present for BuyerUsername, not even an explicit nil
### GetSalesChannel

`func (o *GetOrderResponse) GetSalesChannel() string`

GetSalesChannel returns the SalesChannel field if non-nil, zero value otherwise.

### GetSalesChannelOk

`func (o *GetOrderResponse) GetSalesChannelOk() (*string, bool)`

GetSalesChannelOk returns a tuple with the SalesChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesChannel

`func (o *GetOrderResponse) SetSalesChannel(v string)`

SetSalesChannel sets SalesChannel field to given value.


### GetChannelLocationId

`func (o *GetOrderResponse) GetChannelLocationId() string`

GetChannelLocationId returns the ChannelLocationId field if non-nil, zero value otherwise.

### GetChannelLocationIdOk

`func (o *GetOrderResponse) GetChannelLocationIdOk() (*string, bool)`

GetChannelLocationIdOk returns a tuple with the ChannelLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLocationId

`func (o *GetOrderResponse) SetChannelLocationId(v string)`

SetChannelLocationId sets ChannelLocationId field to given value.

### HasChannelLocationId

`func (o *GetOrderResponse) HasChannelLocationId() bool`

HasChannelLocationId returns a boolean if a field has been set.

### SetChannelLocationIdNil

`func (o *GetOrderResponse) SetChannelLocationIdNil(b bool)`

 SetChannelLocationIdNil sets the value for ChannelLocationId to be an explicit nil

### UnsetChannelLocationId
`func (o *GetOrderResponse) UnsetChannelLocationId()`

UnsetChannelLocationId ensures that no value is present for ChannelLocationId, not even an explicit nil
### GetPackagePresetId

`func (o *GetOrderResponse) GetPackagePresetId() string`

GetPackagePresetId returns the PackagePresetId field if non-nil, zero value otherwise.

### GetPackagePresetIdOk

`func (o *GetOrderResponse) GetPackagePresetIdOk() (*string, bool)`

GetPackagePresetIdOk returns a tuple with the PackagePresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagePresetId

`func (o *GetOrderResponse) SetPackagePresetId(v string)`

SetPackagePresetId sets PackagePresetId field to given value.

### HasPackagePresetId

`func (o *GetOrderResponse) HasPackagePresetId() bool`

HasPackagePresetId returns a boolean if a field has been set.

### SetPackagePresetIdNil

`func (o *GetOrderResponse) SetPackagePresetIdNil(b bool)`

 SetPackagePresetIdNil sets the value for PackagePresetId to be an explicit nil

### UnsetPackagePresetId
`func (o *GetOrderResponse) UnsetPackagePresetId()`

UnsetPackagePresetId ensures that no value is present for PackagePresetId, not even an explicit nil
### GetCarrier

`func (o *GetOrderResponse) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *GetOrderResponse) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *GetOrderResponse) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *GetOrderResponse) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### SetCarrierNil

`func (o *GetOrderResponse) SetCarrierNil(b bool)`

 SetCarrierNil sets the value for Carrier to be an explicit nil

### UnsetCarrier
`func (o *GetOrderResponse) UnsetCarrier()`

UnsetCarrier ensures that no value is present for Carrier, not even an explicit nil
### GetService

`func (o *GetOrderResponse) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *GetOrderResponse) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *GetOrderResponse) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *GetOrderResponse) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *GetOrderResponse) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *GetOrderResponse) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetTrackingNumber

`func (o *GetOrderResponse) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *GetOrderResponse) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *GetOrderResponse) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *GetOrderResponse) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *GetOrderResponse) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *GetOrderResponse) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
### GetEasypostShipmentId

`func (o *GetOrderResponse) GetEasypostShipmentId() string`

GetEasypostShipmentId returns the EasypostShipmentId field if non-nil, zero value otherwise.

### GetEasypostShipmentIdOk

`func (o *GetOrderResponse) GetEasypostShipmentIdOk() (*string, bool)`

GetEasypostShipmentIdOk returns a tuple with the EasypostShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasypostShipmentId

`func (o *GetOrderResponse) SetEasypostShipmentId(v string)`

SetEasypostShipmentId sets EasypostShipmentId field to given value.

### HasEasypostShipmentId

`func (o *GetOrderResponse) HasEasypostShipmentId() bool`

HasEasypostShipmentId returns a boolean if a field has been set.

### SetEasypostShipmentIdNil

`func (o *GetOrderResponse) SetEasypostShipmentIdNil(b bool)`

 SetEasypostShipmentIdNil sets the value for EasypostShipmentId to be an explicit nil

### UnsetEasypostShipmentId
`func (o *GetOrderResponse) UnsetEasypostShipmentId()`

UnsetEasypostShipmentId ensures that no value is present for EasypostShipmentId, not even an explicit nil
### GetEasypostTrackerId

`func (o *GetOrderResponse) GetEasypostTrackerId() string`

GetEasypostTrackerId returns the EasypostTrackerId field if non-nil, zero value otherwise.

### GetEasypostTrackerIdOk

`func (o *GetOrderResponse) GetEasypostTrackerIdOk() (*string, bool)`

GetEasypostTrackerIdOk returns a tuple with the EasypostTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasypostTrackerId

`func (o *GetOrderResponse) SetEasypostTrackerId(v string)`

SetEasypostTrackerId sets EasypostTrackerId field to given value.

### HasEasypostTrackerId

`func (o *GetOrderResponse) HasEasypostTrackerId() bool`

HasEasypostTrackerId returns a boolean if a field has been set.

### SetEasypostTrackerIdNil

`func (o *GetOrderResponse) SetEasypostTrackerIdNil(b bool)`

 SetEasypostTrackerIdNil sets the value for EasypostTrackerId to be an explicit nil

### UnsetEasypostTrackerId
`func (o *GetOrderResponse) UnsetEasypostTrackerId()`

UnsetEasypostTrackerId ensures that no value is present for EasypostTrackerId, not even an explicit nil
### GetEasypostRateId

`func (o *GetOrderResponse) GetEasypostRateId() string`

GetEasypostRateId returns the EasypostRateId field if non-nil, zero value otherwise.

### GetEasypostRateIdOk

`func (o *GetOrderResponse) GetEasypostRateIdOk() (*string, bool)`

GetEasypostRateIdOk returns a tuple with the EasypostRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasypostRateId

`func (o *GetOrderResponse) SetEasypostRateId(v string)`

SetEasypostRateId sets EasypostRateId field to given value.

### HasEasypostRateId

`func (o *GetOrderResponse) HasEasypostRateId() bool`

HasEasypostRateId returns a boolean if a field has been set.

### SetEasypostRateIdNil

`func (o *GetOrderResponse) SetEasypostRateIdNil(b bool)`

 SetEasypostRateIdNil sets the value for EasypostRateId to be an explicit nil

### UnsetEasypostRateId
`func (o *GetOrderResponse) UnsetEasypostRateId()`

UnsetEasypostRateId ensures that no value is present for EasypostRateId, not even an explicit nil
### GetShippingLabelUrl

`func (o *GetOrderResponse) GetShippingLabelUrl() string`

GetShippingLabelUrl returns the ShippingLabelUrl field if non-nil, zero value otherwise.

### GetShippingLabelUrlOk

`func (o *GetOrderResponse) GetShippingLabelUrlOk() (*string, bool)`

GetShippingLabelUrlOk returns a tuple with the ShippingLabelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelUrl

`func (o *GetOrderResponse) SetShippingLabelUrl(v string)`

SetShippingLabelUrl sets ShippingLabelUrl field to given value.

### HasShippingLabelUrl

`func (o *GetOrderResponse) HasShippingLabelUrl() bool`

HasShippingLabelUrl returns a boolean if a field has been set.

### SetShippingLabelUrlNil

`func (o *GetOrderResponse) SetShippingLabelUrlNil(b bool)`

 SetShippingLabelUrlNil sets the value for ShippingLabelUrl to be an explicit nil

### UnsetShippingLabelUrl
`func (o *GetOrderResponse) UnsetShippingLabelUrl()`

UnsetShippingLabelUrl ensures that no value is present for ShippingLabelUrl, not even an explicit nil
### GetLabelPurchasedAt

`func (o *GetOrderResponse) GetLabelPurchasedAt() time.Time`

GetLabelPurchasedAt returns the LabelPurchasedAt field if non-nil, zero value otherwise.

### GetLabelPurchasedAtOk

`func (o *GetOrderResponse) GetLabelPurchasedAtOk() (*time.Time, bool)`

GetLabelPurchasedAtOk returns a tuple with the LabelPurchasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPurchasedAt

`func (o *GetOrderResponse) SetLabelPurchasedAt(v time.Time)`

SetLabelPurchasedAt sets LabelPurchasedAt field to given value.

### HasLabelPurchasedAt

`func (o *GetOrderResponse) HasLabelPurchasedAt() bool`

HasLabelPurchasedAt returns a boolean if a field has been set.

### SetLabelPurchasedAtNil

`func (o *GetOrderResponse) SetLabelPurchasedAtNil(b bool)`

 SetLabelPurchasedAtNil sets the value for LabelPurchasedAt to be an explicit nil

### UnsetLabelPurchasedAt
`func (o *GetOrderResponse) UnsetLabelPurchasedAt()`

UnsetLabelPurchasedAt ensures that no value is present for LabelPurchasedAt, not even an explicit nil
### GetShippedAt

`func (o *GetOrderResponse) GetShippedAt() time.Time`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *GetOrderResponse) GetShippedAtOk() (*time.Time, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *GetOrderResponse) SetShippedAt(v time.Time)`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *GetOrderResponse) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### SetShippedAtNil

`func (o *GetOrderResponse) SetShippedAtNil(b bool)`

 SetShippedAtNil sets the value for ShippedAt to be an explicit nil

### UnsetShippedAt
`func (o *GetOrderResponse) UnsetShippedAt()`

UnsetShippedAt ensures that no value is present for ShippedAt, not even an explicit nil
### GetEstimatedDelivery

`func (o *GetOrderResponse) GetEstimatedDelivery() string`

GetEstimatedDelivery returns the EstimatedDelivery field if non-nil, zero value otherwise.

### GetEstimatedDeliveryOk

`func (o *GetOrderResponse) GetEstimatedDeliveryOk() (*string, bool)`

GetEstimatedDeliveryOk returns a tuple with the EstimatedDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDelivery

`func (o *GetOrderResponse) SetEstimatedDelivery(v string)`

SetEstimatedDelivery sets EstimatedDelivery field to given value.

### HasEstimatedDelivery

`func (o *GetOrderResponse) HasEstimatedDelivery() bool`

HasEstimatedDelivery returns a boolean if a field has been set.

### SetEstimatedDeliveryNil

`func (o *GetOrderResponse) SetEstimatedDeliveryNil(b bool)`

 SetEstimatedDeliveryNil sets the value for EstimatedDelivery to be an explicit nil

### UnsetEstimatedDelivery
`func (o *GetOrderResponse) UnsetEstimatedDelivery()`

UnsetEstimatedDelivery ensures that no value is present for EstimatedDelivery, not even an explicit nil
### GetDeliveredAt

`func (o *GetOrderResponse) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *GetOrderResponse) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *GetOrderResponse) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *GetOrderResponse) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### SetDeliveredAtNil

`func (o *GetOrderResponse) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *GetOrderResponse) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil
### GetCarrierStatus

`func (o *GetOrderResponse) GetCarrierStatus() string`

GetCarrierStatus returns the CarrierStatus field if non-nil, zero value otherwise.

### GetCarrierStatusOk

`func (o *GetOrderResponse) GetCarrierStatusOk() (*string, bool)`

GetCarrierStatusOk returns a tuple with the CarrierStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierStatus

`func (o *GetOrderResponse) SetCarrierStatus(v string)`

SetCarrierStatus sets CarrierStatus field to given value.

### HasCarrierStatus

`func (o *GetOrderResponse) HasCarrierStatus() bool`

HasCarrierStatus returns a boolean if a field has been set.

### SetCarrierStatusNil

`func (o *GetOrderResponse) SetCarrierStatusNil(b bool)`

 SetCarrierStatusNil sets the value for CarrierStatus to be an explicit nil

### UnsetCarrierStatus
`func (o *GetOrderResponse) UnsetCarrierStatus()`

UnsetCarrierStatus ensures that no value is present for CarrierStatus, not even an explicit nil
### GetCarrierStatusDetail

`func (o *GetOrderResponse) GetCarrierStatusDetail() string`

GetCarrierStatusDetail returns the CarrierStatusDetail field if non-nil, zero value otherwise.

### GetCarrierStatusDetailOk

`func (o *GetOrderResponse) GetCarrierStatusDetailOk() (*string, bool)`

GetCarrierStatusDetailOk returns a tuple with the CarrierStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierStatusDetail

`func (o *GetOrderResponse) SetCarrierStatusDetail(v string)`

SetCarrierStatusDetail sets CarrierStatusDetail field to given value.

### HasCarrierStatusDetail

`func (o *GetOrderResponse) HasCarrierStatusDetail() bool`

HasCarrierStatusDetail returns a boolean if a field has been set.

### SetCarrierStatusDetailNil

`func (o *GetOrderResponse) SetCarrierStatusDetailNil(b bool)`

 SetCarrierStatusDetailNil sets the value for CarrierStatusDetail to be an explicit nil

### UnsetCarrierStatusDetail
`func (o *GetOrderResponse) UnsetCarrierStatusDetail()`

UnsetCarrierStatusDetail ensures that no value is present for CarrierStatusDetail, not even an explicit nil
### GetTrackingHistory

`func (o *GetOrderResponse) GetTrackingHistory() []ListOrdersItemTrackingHistory`

GetTrackingHistory returns the TrackingHistory field if non-nil, zero value otherwise.

### GetTrackingHistoryOk

`func (o *GetOrderResponse) GetTrackingHistoryOk() (*[]ListOrdersItemTrackingHistory, bool)`

GetTrackingHistoryOk returns a tuple with the TrackingHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingHistory

`func (o *GetOrderResponse) SetTrackingHistory(v []ListOrdersItemTrackingHistory)`

SetTrackingHistory sets TrackingHistory field to given value.

### HasTrackingHistory

`func (o *GetOrderResponse) HasTrackingHistory() bool`

HasTrackingHistory returns a boolean if a field has been set.

### SetTrackingHistoryNil

`func (o *GetOrderResponse) SetTrackingHistoryNil(b bool)`

 SetTrackingHistoryNil sets the value for TrackingHistory to be an explicit nil

### UnsetTrackingHistory
`func (o *GetOrderResponse) UnsetTrackingHistory()`

UnsetTrackingHistory ensures that no value is present for TrackingHistory, not even an explicit nil
### GetDeliveryLocation

`func (o *GetOrderResponse) GetDeliveryLocation() string`

GetDeliveryLocation returns the DeliveryLocation field if non-nil, zero value otherwise.

### GetDeliveryLocationOk

`func (o *GetOrderResponse) GetDeliveryLocationOk() (*string, bool)`

GetDeliveryLocationOk returns a tuple with the DeliveryLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLocation

`func (o *GetOrderResponse) SetDeliveryLocation(v string)`

SetDeliveryLocation sets DeliveryLocation field to given value.

### HasDeliveryLocation

`func (o *GetOrderResponse) HasDeliveryLocation() bool`

HasDeliveryLocation returns a boolean if a field has been set.

### SetDeliveryLocationNil

`func (o *GetOrderResponse) SetDeliveryLocationNil(b bool)`

 SetDeliveryLocationNil sets the value for DeliveryLocation to be an explicit nil

### UnsetDeliveryLocation
`func (o *GetOrderResponse) UnsetDeliveryLocation()`

UnsetDeliveryLocation ensures that no value is present for DeliveryLocation, not even an explicit nil
### GetDeliverySignature

`func (o *GetOrderResponse) GetDeliverySignature() string`

GetDeliverySignature returns the DeliverySignature field if non-nil, zero value otherwise.

### GetDeliverySignatureOk

`func (o *GetOrderResponse) GetDeliverySignatureOk() (*string, bool)`

GetDeliverySignatureOk returns a tuple with the DeliverySignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverySignature

`func (o *GetOrderResponse) SetDeliverySignature(v string)`

SetDeliverySignature sets DeliverySignature field to given value.

### HasDeliverySignature

`func (o *GetOrderResponse) HasDeliverySignature() bool`

HasDeliverySignature returns a boolean if a field has been set.

### SetDeliverySignatureNil

`func (o *GetOrderResponse) SetDeliverySignatureNil(b bool)`

 SetDeliverySignatureNil sets the value for DeliverySignature to be an explicit nil

### UnsetDeliverySignature
`func (o *GetOrderResponse) UnsetDeliverySignature()`

UnsetDeliverySignature ensures that no value is present for DeliverySignature, not even an explicit nil
### GetTrackingSubmittedAt

`func (o *GetOrderResponse) GetTrackingSubmittedAt() time.Time`

GetTrackingSubmittedAt returns the TrackingSubmittedAt field if non-nil, zero value otherwise.

### GetTrackingSubmittedAtOk

`func (o *GetOrderResponse) GetTrackingSubmittedAtOk() (*time.Time, bool)`

GetTrackingSubmittedAtOk returns a tuple with the TrackingSubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingSubmittedAt

`func (o *GetOrderResponse) SetTrackingSubmittedAt(v time.Time)`

SetTrackingSubmittedAt sets TrackingSubmittedAt field to given value.

### HasTrackingSubmittedAt

`func (o *GetOrderResponse) HasTrackingSubmittedAt() bool`

HasTrackingSubmittedAt returns a boolean if a field has been set.

### SetTrackingSubmittedAtNil

`func (o *GetOrderResponse) SetTrackingSubmittedAtNil(b bool)`

 SetTrackingSubmittedAtNil sets the value for TrackingSubmittedAt to be an explicit nil

### UnsetTrackingSubmittedAt
`func (o *GetOrderResponse) UnsetTrackingSubmittedAt()`

UnsetTrackingSubmittedAt ensures that no value is present for TrackingSubmittedAt, not even an explicit nil
### GetTrackingSubmitStatus

`func (o *GetOrderResponse) GetTrackingSubmitStatus() string`

GetTrackingSubmitStatus returns the TrackingSubmitStatus field if non-nil, zero value otherwise.

### GetTrackingSubmitStatusOk

`func (o *GetOrderResponse) GetTrackingSubmitStatusOk() (*string, bool)`

GetTrackingSubmitStatusOk returns a tuple with the TrackingSubmitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingSubmitStatus

`func (o *GetOrderResponse) SetTrackingSubmitStatus(v string)`

SetTrackingSubmitStatus sets TrackingSubmitStatus field to given value.

### HasTrackingSubmitStatus

`func (o *GetOrderResponse) HasTrackingSubmitStatus() bool`

HasTrackingSubmitStatus returns a boolean if a field has been set.

### SetTrackingSubmitStatusNil

`func (o *GetOrderResponse) SetTrackingSubmitStatusNil(b bool)`

 SetTrackingSubmitStatusNil sets the value for TrackingSubmitStatus to be an explicit nil

### UnsetTrackingSubmitStatus
`func (o *GetOrderResponse) UnsetTrackingSubmitStatus()`

UnsetTrackingSubmitStatus ensures that no value is present for TrackingSubmitStatus, not even an explicit nil
### GetLabelCost

`func (o *GetOrderResponse) GetLabelCost() string`

GetLabelCost returns the LabelCost field if non-nil, zero value otherwise.

### GetLabelCostOk

`func (o *GetOrderResponse) GetLabelCostOk() (*string, bool)`

GetLabelCostOk returns a tuple with the LabelCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelCost

`func (o *GetOrderResponse) SetLabelCost(v string)`

SetLabelCost sets LabelCost field to given value.

### HasLabelCost

`func (o *GetOrderResponse) HasLabelCost() bool`

HasLabelCost returns a boolean if a field has been set.

### SetLabelCostNil

`func (o *GetOrderResponse) SetLabelCostNil(b bool)`

 SetLabelCostNil sets the value for LabelCost to be an explicit nil

### UnsetLabelCost
`func (o *GetOrderResponse) UnsetLabelCost()`

UnsetLabelCost ensures that no value is present for LabelCost, not even an explicit nil
### GetCostOfGoods

`func (o *GetOrderResponse) GetCostOfGoods() string`

GetCostOfGoods returns the CostOfGoods field if non-nil, zero value otherwise.

### GetCostOfGoodsOk

`func (o *GetOrderResponse) GetCostOfGoodsOk() (*string, bool)`

GetCostOfGoodsOk returns a tuple with the CostOfGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOfGoods

`func (o *GetOrderResponse) SetCostOfGoods(v string)`

SetCostOfGoods sets CostOfGoods field to given value.

### HasCostOfGoods

`func (o *GetOrderResponse) HasCostOfGoods() bool`

HasCostOfGoods returns a boolean if a field has been set.

### SetCostOfGoodsNil

`func (o *GetOrderResponse) SetCostOfGoodsNil(b bool)`

 SetCostOfGoodsNil sets the value for CostOfGoods to be an explicit nil

### UnsetCostOfGoods
`func (o *GetOrderResponse) UnsetCostOfGoods()`

UnsetCostOfGoods ensures that no value is present for CostOfGoods, not even an explicit nil
### GetRequestedCarrier

`func (o *GetOrderResponse) GetRequestedCarrier() string`

GetRequestedCarrier returns the RequestedCarrier field if non-nil, zero value otherwise.

### GetRequestedCarrierOk

`func (o *GetOrderResponse) GetRequestedCarrierOk() (*string, bool)`

GetRequestedCarrierOk returns a tuple with the RequestedCarrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCarrier

`func (o *GetOrderResponse) SetRequestedCarrier(v string)`

SetRequestedCarrier sets RequestedCarrier field to given value.

### HasRequestedCarrier

`func (o *GetOrderResponse) HasRequestedCarrier() bool`

HasRequestedCarrier returns a boolean if a field has been set.

### SetRequestedCarrierNil

`func (o *GetOrderResponse) SetRequestedCarrierNil(b bool)`

 SetRequestedCarrierNil sets the value for RequestedCarrier to be an explicit nil

### UnsetRequestedCarrier
`func (o *GetOrderResponse) UnsetRequestedCarrier()`

UnsetRequestedCarrier ensures that no value is present for RequestedCarrier, not even an explicit nil
### GetRequestedService

`func (o *GetOrderResponse) GetRequestedService() string`

GetRequestedService returns the RequestedService field if non-nil, zero value otherwise.

### GetRequestedServiceOk

`func (o *GetOrderResponse) GetRequestedServiceOk() (*string, bool)`

GetRequestedServiceOk returns a tuple with the RequestedService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedService

`func (o *GetOrderResponse) SetRequestedService(v string)`

SetRequestedService sets RequestedService field to given value.

### HasRequestedService

`func (o *GetOrderResponse) HasRequestedService() bool`

HasRequestedService returns a boolean if a field has been set.

### SetRequestedServiceNil

`func (o *GetOrderResponse) SetRequestedServiceNil(b bool)`

 SetRequestedServiceNil sets the value for RequestedService to be an explicit nil

### UnsetRequestedService
`func (o *GetOrderResponse) UnsetRequestedService()`

UnsetRequestedService ensures that no value is present for RequestedService, not even an explicit nil
### GetShipByAt

`func (o *GetOrderResponse) GetShipByAt() time.Time`

GetShipByAt returns the ShipByAt field if non-nil, zero value otherwise.

### GetShipByAtOk

`func (o *GetOrderResponse) GetShipByAtOk() (*time.Time, bool)`

GetShipByAtOk returns a tuple with the ShipByAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipByAt

`func (o *GetOrderResponse) SetShipByAt(v time.Time)`

SetShipByAt sets ShipByAt field to given value.

### HasShipByAt

`func (o *GetOrderResponse) HasShipByAt() bool`

HasShipByAt returns a boolean if a field has been set.

### SetShipByAtNil

`func (o *GetOrderResponse) SetShipByAtNil(b bool)`

 SetShipByAtNil sets the value for ShipByAt to be an explicit nil

### UnsetShipByAt
`func (o *GetOrderResponse) UnsetShipByAt()`

UnsetShipByAt ensures that no value is present for ShipByAt, not even an explicit nil
### GetShipByAlertedAt

`func (o *GetOrderResponse) GetShipByAlertedAt() time.Time`

GetShipByAlertedAt returns the ShipByAlertedAt field if non-nil, zero value otherwise.

### GetShipByAlertedAtOk

`func (o *GetOrderResponse) GetShipByAlertedAtOk() (*time.Time, bool)`

GetShipByAlertedAtOk returns a tuple with the ShipByAlertedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipByAlertedAt

`func (o *GetOrderResponse) SetShipByAlertedAt(v time.Time)`

SetShipByAlertedAt sets ShipByAlertedAt field to given value.

### HasShipByAlertedAt

`func (o *GetOrderResponse) HasShipByAlertedAt() bool`

HasShipByAlertedAt returns a boolean if a field has been set.

### SetShipByAlertedAtNil

`func (o *GetOrderResponse) SetShipByAlertedAtNil(b bool)`

 SetShipByAlertedAtNil sets the value for ShipByAlertedAt to be an explicit nil

### UnsetShipByAlertedAt
`func (o *GetOrderResponse) UnsetShipByAlertedAt()`

UnsetShipByAlertedAt ensures that no value is present for ShipByAlertedAt, not even an explicit nil
### GetOversoldBy

`func (o *GetOrderResponse) GetOversoldBy() float32`

GetOversoldBy returns the OversoldBy field if non-nil, zero value otherwise.

### GetOversoldByOk

`func (o *GetOrderResponse) GetOversoldByOk() (*float32, bool)`

GetOversoldByOk returns a tuple with the OversoldBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversoldBy

`func (o *GetOrderResponse) SetOversoldBy(v float32)`

SetOversoldBy sets OversoldBy field to given value.


### GetLastStatusCheckAt

`func (o *GetOrderResponse) GetLastStatusCheckAt() time.Time`

GetLastStatusCheckAt returns the LastStatusCheckAt field if non-nil, zero value otherwise.

### GetLastStatusCheckAtOk

`func (o *GetOrderResponse) GetLastStatusCheckAtOk() (*time.Time, bool)`

GetLastStatusCheckAtOk returns a tuple with the LastStatusCheckAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusCheckAt

`func (o *GetOrderResponse) SetLastStatusCheckAt(v time.Time)`

SetLastStatusCheckAt sets LastStatusCheckAt field to given value.

### HasLastStatusCheckAt

`func (o *GetOrderResponse) HasLastStatusCheckAt() bool`

HasLastStatusCheckAt returns a boolean if a field has been set.

### SetLastStatusCheckAtNil

`func (o *GetOrderResponse) SetLastStatusCheckAtNil(b bool)`

 SetLastStatusCheckAtNil sets the value for LastStatusCheckAt to be an explicit nil

### UnsetLastStatusCheckAt
`func (o *GetOrderResponse) UnsetLastStatusCheckAt()`

UnsetLastStatusCheckAt ensures that no value is present for LastStatusCheckAt, not even an explicit nil
### GetLastChatCheckAt

`func (o *GetOrderResponse) GetLastChatCheckAt() time.Time`

GetLastChatCheckAt returns the LastChatCheckAt field if non-nil, zero value otherwise.

### GetLastChatCheckAtOk

`func (o *GetOrderResponse) GetLastChatCheckAtOk() (*time.Time, bool)`

GetLastChatCheckAtOk returns a tuple with the LastChatCheckAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChatCheckAt

`func (o *GetOrderResponse) SetLastChatCheckAt(v time.Time)`

SetLastChatCheckAt sets LastChatCheckAt field to given value.

### HasLastChatCheckAt

`func (o *GetOrderResponse) HasLastChatCheckAt() bool`

HasLastChatCheckAt returns a boolean if a field has been set.

### SetLastChatCheckAtNil

`func (o *GetOrderResponse) SetLastChatCheckAtNil(b bool)`

 SetLastChatCheckAtNil sets the value for LastChatCheckAt to be an explicit nil

### UnsetLastChatCheckAt
`func (o *GetOrderResponse) UnsetLastChatCheckAt()`

UnsetLastChatCheckAt ensures that no value is present for LastChatCheckAt, not even an explicit nil
### GetIsDisputed

`func (o *GetOrderResponse) GetIsDisputed() bool`

GetIsDisputed returns the IsDisputed field if non-nil, zero value otherwise.

### GetIsDisputedOk

`func (o *GetOrderResponse) GetIsDisputedOk() (*bool, bool)`

GetIsDisputedOk returns a tuple with the IsDisputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisputed

`func (o *GetOrderResponse) SetIsDisputed(v bool)`

SetIsDisputed sets IsDisputed field to given value.


### GetDisputeReason

`func (o *GetOrderResponse) GetDisputeReason() string`

GetDisputeReason returns the DisputeReason field if non-nil, zero value otherwise.

### GetDisputeReasonOk

`func (o *GetOrderResponse) GetDisputeReasonOk() (*string, bool)`

GetDisputeReasonOk returns a tuple with the DisputeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeReason

`func (o *GetOrderResponse) SetDisputeReason(v string)`

SetDisputeReason sets DisputeReason field to given value.

### HasDisputeReason

`func (o *GetOrderResponse) HasDisputeReason() bool`

HasDisputeReason returns a boolean if a field has been set.

### SetDisputeReasonNil

`func (o *GetOrderResponse) SetDisputeReasonNil(b bool)`

 SetDisputeReasonNil sets the value for DisputeReason to be an explicit nil

### UnsetDisputeReason
`func (o *GetOrderResponse) UnsetDisputeReason()`

UnsetDisputeReason ensures that no value is present for DisputeReason, not even an explicit nil
### GetDisputePlatformCaseId

`func (o *GetOrderResponse) GetDisputePlatformCaseId() string`

GetDisputePlatformCaseId returns the DisputePlatformCaseId field if non-nil, zero value otherwise.

### GetDisputePlatformCaseIdOk

`func (o *GetOrderResponse) GetDisputePlatformCaseIdOk() (*string, bool)`

GetDisputePlatformCaseIdOk returns a tuple with the DisputePlatformCaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputePlatformCaseId

`func (o *GetOrderResponse) SetDisputePlatformCaseId(v string)`

SetDisputePlatformCaseId sets DisputePlatformCaseId field to given value.

### HasDisputePlatformCaseId

`func (o *GetOrderResponse) HasDisputePlatformCaseId() bool`

HasDisputePlatformCaseId returns a boolean if a field has been set.

### SetDisputePlatformCaseIdNil

`func (o *GetOrderResponse) SetDisputePlatformCaseIdNil(b bool)`

 SetDisputePlatformCaseIdNil sets the value for DisputePlatformCaseId to be an explicit nil

### UnsetDisputePlatformCaseId
`func (o *GetOrderResponse) UnsetDisputePlatformCaseId()`

UnsetDisputePlatformCaseId ensures that no value is present for DisputePlatformCaseId, not even an explicit nil
### GetDisputeResolvedAt

`func (o *GetOrderResponse) GetDisputeResolvedAt() time.Time`

GetDisputeResolvedAt returns the DisputeResolvedAt field if non-nil, zero value otherwise.

### GetDisputeResolvedAtOk

`func (o *GetOrderResponse) GetDisputeResolvedAtOk() (*time.Time, bool)`

GetDisputeResolvedAtOk returns a tuple with the DisputeResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeResolvedAt

`func (o *GetOrderResponse) SetDisputeResolvedAt(v time.Time)`

SetDisputeResolvedAt sets DisputeResolvedAt field to given value.

### HasDisputeResolvedAt

`func (o *GetOrderResponse) HasDisputeResolvedAt() bool`

HasDisputeResolvedAt returns a boolean if a field has been set.

### SetDisputeResolvedAtNil

`func (o *GetOrderResponse) SetDisputeResolvedAtNil(b bool)`

 SetDisputeResolvedAtNil sets the value for DisputeResolvedAt to be an explicit nil

### UnsetDisputeResolvedAt
`func (o *GetOrderResponse) UnsetDisputeResolvedAt()`

UnsetDisputeResolvedAt ensures that no value is present for DisputeResolvedAt, not even an explicit nil
### GetRefundAmount

`func (o *GetOrderResponse) GetRefundAmount() string`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *GetOrderResponse) GetRefundAmountOk() (*string, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *GetOrderResponse) SetRefundAmount(v string)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *GetOrderResponse) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### SetRefundAmountNil

`func (o *GetOrderResponse) SetRefundAmountNil(b bool)`

 SetRefundAmountNil sets the value for RefundAmount to be an explicit nil

### UnsetRefundAmount
`func (o *GetOrderResponse) UnsetRefundAmount()`

UnsetRefundAmount ensures that no value is present for RefundAmount, not even an explicit nil
### GetRefundReason

`func (o *GetOrderResponse) GetRefundReason() string`

GetRefundReason returns the RefundReason field if non-nil, zero value otherwise.

### GetRefundReasonOk

`func (o *GetOrderResponse) GetRefundReasonOk() (*string, bool)`

GetRefundReasonOk returns a tuple with the RefundReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundReason

`func (o *GetOrderResponse) SetRefundReason(v string)`

SetRefundReason sets RefundReason field to given value.

### HasRefundReason

`func (o *GetOrderResponse) HasRefundReason() bool`

HasRefundReason returns a boolean if a field has been set.

### SetRefundReasonNil

`func (o *GetOrderResponse) SetRefundReasonNil(b bool)`

 SetRefundReasonNil sets the value for RefundReason to be an explicit nil

### UnsetRefundReason
`func (o *GetOrderResponse) UnsetRefundReason()`

UnsetRefundReason ensures that no value is present for RefundReason, not even an explicit nil
### GetRefundPlatformId

`func (o *GetOrderResponse) GetRefundPlatformId() string`

GetRefundPlatformId returns the RefundPlatformId field if non-nil, zero value otherwise.

### GetRefundPlatformIdOk

`func (o *GetOrderResponse) GetRefundPlatformIdOk() (*string, bool)`

GetRefundPlatformIdOk returns a tuple with the RefundPlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundPlatformId

`func (o *GetOrderResponse) SetRefundPlatformId(v string)`

SetRefundPlatformId sets RefundPlatformId field to given value.

### HasRefundPlatformId

`func (o *GetOrderResponse) HasRefundPlatformId() bool`

HasRefundPlatformId returns a boolean if a field has been set.

### SetRefundPlatformIdNil

`func (o *GetOrderResponse) SetRefundPlatformIdNil(b bool)`

 SetRefundPlatformIdNil sets the value for RefundPlatformId to be an explicit nil

### UnsetRefundPlatformId
`func (o *GetOrderResponse) UnsetRefundPlatformId()`

UnsetRefundPlatformId ensures that no value is present for RefundPlatformId, not even an explicit nil
### GetRefundedAt

`func (o *GetOrderResponse) GetRefundedAt() time.Time`

GetRefundedAt returns the RefundedAt field if non-nil, zero value otherwise.

### GetRefundedAtOk

`func (o *GetOrderResponse) GetRefundedAtOk() (*time.Time, bool)`

GetRefundedAtOk returns a tuple with the RefundedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAt

`func (o *GetOrderResponse) SetRefundedAt(v time.Time)`

SetRefundedAt sets RefundedAt field to given value.

### HasRefundedAt

`func (o *GetOrderResponse) HasRefundedAt() bool`

HasRefundedAt returns a boolean if a field has been set.

### SetRefundedAtNil

`func (o *GetOrderResponse) SetRefundedAtNil(b bool)`

 SetRefundedAtNil sets the value for RefundedAt to be an explicit nil

### UnsetRefundedAt
`func (o *GetOrderResponse) UnsetRefundedAt()`

UnsetRefundedAt ensures that no value is present for RefundedAt, not even an explicit nil
### GetCancellationReason

`func (o *GetOrderResponse) GetCancellationReason() string`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *GetOrderResponse) GetCancellationReasonOk() (*string, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationReason

`func (o *GetOrderResponse) SetCancellationReason(v string)`

SetCancellationReason sets CancellationReason field to given value.

### HasCancellationReason

`func (o *GetOrderResponse) HasCancellationReason() bool`

HasCancellationReason returns a boolean if a field has been set.

### SetCancellationReasonNil

`func (o *GetOrderResponse) SetCancellationReasonNil(b bool)`

 SetCancellationReasonNil sets the value for CancellationReason to be an explicit nil

### UnsetCancellationReason
`func (o *GetOrderResponse) UnsetCancellationReason()`

UnsetCancellationReason ensures that no value is present for CancellationReason, not even an explicit nil
### GetArrivalConditionRequestedAt

`func (o *GetOrderResponse) GetArrivalConditionRequestedAt() time.Time`

GetArrivalConditionRequestedAt returns the ArrivalConditionRequestedAt field if non-nil, zero value otherwise.

### GetArrivalConditionRequestedAtOk

`func (o *GetOrderResponse) GetArrivalConditionRequestedAtOk() (*time.Time, bool)`

GetArrivalConditionRequestedAtOk returns a tuple with the ArrivalConditionRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalConditionRequestedAt

`func (o *GetOrderResponse) SetArrivalConditionRequestedAt(v time.Time)`

SetArrivalConditionRequestedAt sets ArrivalConditionRequestedAt field to given value.

### HasArrivalConditionRequestedAt

`func (o *GetOrderResponse) HasArrivalConditionRequestedAt() bool`

HasArrivalConditionRequestedAt returns a boolean if a field has been set.

### SetArrivalConditionRequestedAtNil

`func (o *GetOrderResponse) SetArrivalConditionRequestedAtNil(b bool)`

 SetArrivalConditionRequestedAtNil sets the value for ArrivalConditionRequestedAt to be an explicit nil

### UnsetArrivalConditionRequestedAt
`func (o *GetOrderResponse) UnsetArrivalConditionRequestedAt()`

UnsetArrivalConditionRequestedAt ensures that no value is present for ArrivalConditionRequestedAt, not even an explicit nil
### GetArrivalConditionSubmittedAt

`func (o *GetOrderResponse) GetArrivalConditionSubmittedAt() time.Time`

GetArrivalConditionSubmittedAt returns the ArrivalConditionSubmittedAt field if non-nil, zero value otherwise.

### GetArrivalConditionSubmittedAtOk

`func (o *GetOrderResponse) GetArrivalConditionSubmittedAtOk() (*time.Time, bool)`

GetArrivalConditionSubmittedAtOk returns a tuple with the ArrivalConditionSubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalConditionSubmittedAt

`func (o *GetOrderResponse) SetArrivalConditionSubmittedAt(v time.Time)`

SetArrivalConditionSubmittedAt sets ArrivalConditionSubmittedAt field to given value.

### HasArrivalConditionSubmittedAt

`func (o *GetOrderResponse) HasArrivalConditionSubmittedAt() bool`

HasArrivalConditionSubmittedAt returns a boolean if a field has been set.

### SetArrivalConditionSubmittedAtNil

`func (o *GetOrderResponse) SetArrivalConditionSubmittedAtNil(b bool)`

 SetArrivalConditionSubmittedAtNil sets the value for ArrivalConditionSubmittedAt to be an explicit nil

### UnsetArrivalConditionSubmittedAt
`func (o *GetOrderResponse) UnsetArrivalConditionSubmittedAt()`

UnsetArrivalConditionSubmittedAt ensures that no value is present for ArrivalConditionSubmittedAt, not even an explicit nil
### GetArrivalConditionDeclinedAt

`func (o *GetOrderResponse) GetArrivalConditionDeclinedAt() time.Time`

GetArrivalConditionDeclinedAt returns the ArrivalConditionDeclinedAt field if non-nil, zero value otherwise.

### GetArrivalConditionDeclinedAtOk

`func (o *GetOrderResponse) GetArrivalConditionDeclinedAtOk() (*time.Time, bool)`

GetArrivalConditionDeclinedAtOk returns a tuple with the ArrivalConditionDeclinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalConditionDeclinedAt

`func (o *GetOrderResponse) SetArrivalConditionDeclinedAt(v time.Time)`

SetArrivalConditionDeclinedAt sets ArrivalConditionDeclinedAt field to given value.

### HasArrivalConditionDeclinedAt

`func (o *GetOrderResponse) HasArrivalConditionDeclinedAt() bool`

HasArrivalConditionDeclinedAt returns a boolean if a field has been set.

### SetArrivalConditionDeclinedAtNil

`func (o *GetOrderResponse) SetArrivalConditionDeclinedAtNil(b bool)`

 SetArrivalConditionDeclinedAtNil sets the value for ArrivalConditionDeclinedAt to be an explicit nil

### UnsetArrivalConditionDeclinedAt
`func (o *GetOrderResponse) UnsetArrivalConditionDeclinedAt()`

UnsetArrivalConditionDeclinedAt ensures that no value is present for ArrivalConditionDeclinedAt, not even an explicit nil
### GetDeliveryPhotoUrl

`func (o *GetOrderResponse) GetDeliveryPhotoUrl() string`

GetDeliveryPhotoUrl returns the DeliveryPhotoUrl field if non-nil, zero value otherwise.

### GetDeliveryPhotoUrlOk

`func (o *GetOrderResponse) GetDeliveryPhotoUrlOk() (*string, bool)`

GetDeliveryPhotoUrlOk returns a tuple with the DeliveryPhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPhotoUrl

`func (o *GetOrderResponse) SetDeliveryPhotoUrl(v string)`

SetDeliveryPhotoUrl sets DeliveryPhotoUrl field to given value.

### HasDeliveryPhotoUrl

`func (o *GetOrderResponse) HasDeliveryPhotoUrl() bool`

HasDeliveryPhotoUrl returns a boolean if a field has been set.

### SetDeliveryPhotoUrlNil

`func (o *GetOrderResponse) SetDeliveryPhotoUrlNil(b bool)`

 SetDeliveryPhotoUrlNil sets the value for DeliveryPhotoUrl to be an explicit nil

### UnsetDeliveryPhotoUrl
`func (o *GetOrderResponse) UnsetDeliveryPhotoUrl()`

UnsetDeliveryPhotoUrl ensures that no value is present for DeliveryPhotoUrl, not even an explicit nil
### GetPurchaseOrderRef

`func (o *GetOrderResponse) GetPurchaseOrderRef() string`

GetPurchaseOrderRef returns the PurchaseOrderRef field if non-nil, zero value otherwise.

### GetPurchaseOrderRefOk

`func (o *GetOrderResponse) GetPurchaseOrderRefOk() (*string, bool)`

GetPurchaseOrderRefOk returns a tuple with the PurchaseOrderRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderRef

`func (o *GetOrderResponse) SetPurchaseOrderRef(v string)`

SetPurchaseOrderRef sets PurchaseOrderRef field to given value.

### HasPurchaseOrderRef

`func (o *GetOrderResponse) HasPurchaseOrderRef() bool`

HasPurchaseOrderRef returns a boolean if a field has been set.

### SetPurchaseOrderRefNil

`func (o *GetOrderResponse) SetPurchaseOrderRefNil(b bool)`

 SetPurchaseOrderRefNil sets the value for PurchaseOrderRef to be an explicit nil

### UnsetPurchaseOrderRef
`func (o *GetOrderResponse) UnsetPurchaseOrderRef()`

UnsetPurchaseOrderRef ensures that no value is present for PurchaseOrderRef, not even an explicit nil
### GetArrivalConditionPhotos

`func (o *GetOrderResponse) GetArrivalConditionPhotos() []string`

GetArrivalConditionPhotos returns the ArrivalConditionPhotos field if non-nil, zero value otherwise.

### GetArrivalConditionPhotosOk

`func (o *GetOrderResponse) GetArrivalConditionPhotosOk() (*[]string, bool)`

GetArrivalConditionPhotosOk returns a tuple with the ArrivalConditionPhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalConditionPhotos

`func (o *GetOrderResponse) SetArrivalConditionPhotos(v []string)`

SetArrivalConditionPhotos sets ArrivalConditionPhotos field to given value.


### GetBuyerEmail

`func (o *GetOrderResponse) GetBuyerEmail() string`

GetBuyerEmail returns the BuyerEmail field if non-nil, zero value otherwise.

### GetBuyerEmailOk

`func (o *GetOrderResponse) GetBuyerEmailOk() (*string, bool)`

GetBuyerEmailOk returns a tuple with the BuyerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerEmail

`func (o *GetOrderResponse) SetBuyerEmail(v string)`

SetBuyerEmail sets BuyerEmail field to given value.

### HasBuyerEmail

`func (o *GetOrderResponse) HasBuyerEmail() bool`

HasBuyerEmail returns a boolean if a field has been set.

### SetBuyerEmailNil

`func (o *GetOrderResponse) SetBuyerEmailNil(b bool)`

 SetBuyerEmailNil sets the value for BuyerEmail to be an explicit nil

### UnsetBuyerEmail
`func (o *GetOrderResponse) UnsetBuyerEmail()`

UnsetBuyerEmail ensures that no value is present for BuyerEmail, not even an explicit nil
### GetFulfillmentMethod

`func (o *GetOrderResponse) GetFulfillmentMethod() string`

GetFulfillmentMethod returns the FulfillmentMethod field if non-nil, zero value otherwise.

### GetFulfillmentMethodOk

`func (o *GetOrderResponse) GetFulfillmentMethodOk() (*string, bool)`

GetFulfillmentMethodOk returns a tuple with the FulfillmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentMethod

`func (o *GetOrderResponse) SetFulfillmentMethod(v string)`

SetFulfillmentMethod sets FulfillmentMethod field to given value.


### GetDeletedAt

`func (o *GetOrderResponse) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *GetOrderResponse) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *GetOrderResponse) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *GetOrderResponse) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *GetOrderResponse) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *GetOrderResponse) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


