# ListOrdersItem

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

### NewListOrdersItem

`func NewListOrdersItem(id string, createdAt time.Time, updatedAt time.Time, userId string, quantity float32, status string, platform string, salesChannel string, oversoldBy float32, isDisputed bool, arrivalConditionPhotos []string, fulfillmentMethod string, ) *ListOrdersItem`

NewListOrdersItem instantiates a new ListOrdersItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOrdersItemWithDefaults

`func NewListOrdersItemWithDefaults() *ListOrdersItem`

NewListOrdersItemWithDefaults instantiates a new ListOrdersItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListOrdersItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListOrdersItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListOrdersItem) SetId(v string)`

SetId sets Id field to given value.


### GetCreatedAt

`func (o *ListOrdersItem) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListOrdersItem) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListOrdersItem) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *ListOrdersItem) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListOrdersItem) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListOrdersItem) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.


### GetUserId

`func (o *ListOrdersItem) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ListOrdersItem) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ListOrdersItem) SetUserId(v string)`

SetUserId sets UserId field to given value.


### GetQuantity

`func (o *ListOrdersItem) GetQuantity() float32`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *ListOrdersItem) GetQuantityOk() (*float32, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *ListOrdersItem) SetQuantity(v float32)`

SetQuantity sets Quantity field to given value.


### GetNotes

`func (o *ListOrdersItem) GetNotes() string`

GetNotes returns the Notes field if non-nil, zero value otherwise.

### GetNotesOk

`func (o *ListOrdersItem) GetNotesOk() (*string, bool)`

GetNotesOk returns a tuple with the Notes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotes

`func (o *ListOrdersItem) SetNotes(v string)`

SetNotes sets Notes field to given value.

### HasNotes

`func (o *ListOrdersItem) HasNotes() bool`

HasNotes returns a boolean if a field has been set.

### SetNotesNil

`func (o *ListOrdersItem) SetNotesNil(b bool)`

 SetNotesNil sets the value for Notes to be an explicit nil

### UnsetNotes
`func (o *ListOrdersItem) UnsetNotes()`

UnsetNotes ensures that no value is present for Notes, not even an explicit nil
### GetStatus

`func (o *ListOrdersItem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListOrdersItem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListOrdersItem) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetPlatform

`func (o *ListOrdersItem) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListOrdersItem) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListOrdersItem) SetPlatform(v string)`

SetPlatform sets Platform field to given value.


### GetCancelledAt

`func (o *ListOrdersItem) GetCancelledAt() time.Time`

GetCancelledAt returns the CancelledAt field if non-nil, zero value otherwise.

### GetCancelledAtOk

`func (o *ListOrdersItem) GetCancelledAtOk() (*time.Time, bool)`

GetCancelledAtOk returns a tuple with the CancelledAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancelledAt

`func (o *ListOrdersItem) SetCancelledAt(v time.Time)`

SetCancelledAt sets CancelledAt field to given value.

### HasCancelledAt

`func (o *ListOrdersItem) HasCancelledAt() bool`

HasCancelledAt returns a boolean if a field has been set.

### SetCancelledAtNil

`func (o *ListOrdersItem) SetCancelledAtNil(b bool)`

 SetCancelledAtNil sets the value for CancelledAt to be an explicit nil

### UnsetCancelledAt
`func (o *ListOrdersItem) UnsetCancelledAt()`

UnsetCancelledAt ensures that no value is present for CancelledAt, not even an explicit nil
### GetListingId

`func (o *ListOrdersItem) GetListingId() string`

GetListingId returns the ListingId field if non-nil, zero value otherwise.

### GetListingIdOk

`func (o *ListOrdersItem) GetListingIdOk() (*string, bool)`

GetListingIdOk returns a tuple with the ListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetListingId

`func (o *ListOrdersItem) SetListingId(v string)`

SetListingId sets ListingId field to given value.

### HasListingId

`func (o *ListOrdersItem) HasListingId() bool`

HasListingId returns a boolean if a field has been set.

### SetListingIdNil

`func (o *ListOrdersItem) SetListingIdNil(b bool)`

 SetListingIdNil sets the value for ListingId to be an explicit nil

### UnsetListingId
`func (o *ListOrdersItem) UnsetListingId()`

UnsetListingId ensures that no value is present for ListingId, not even an explicit nil
### GetInventoryItemId

`func (o *ListOrdersItem) GetInventoryItemId() string`

GetInventoryItemId returns the InventoryItemId field if non-nil, zero value otherwise.

### GetInventoryItemIdOk

`func (o *ListOrdersItem) GetInventoryItemIdOk() (*string, bool)`

GetInventoryItemIdOk returns a tuple with the InventoryItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventoryItemId

`func (o *ListOrdersItem) SetInventoryItemId(v string)`

SetInventoryItemId sets InventoryItemId field to given value.

### HasInventoryItemId

`func (o *ListOrdersItem) HasInventoryItemId() bool`

HasInventoryItemId returns a boolean if a field has been set.

### SetInventoryItemIdNil

`func (o *ListOrdersItem) SetInventoryItemIdNil(b bool)`

 SetInventoryItemIdNil sets the value for InventoryItemId to be an explicit nil

### UnsetInventoryItemId
`func (o *ListOrdersItem) UnsetInventoryItemId()`

UnsetInventoryItemId ensures that no value is present for InventoryItemId, not even an explicit nil
### GetPlatformListingId

`func (o *ListOrdersItem) GetPlatformListingId() string`

GetPlatformListingId returns the PlatformListingId field if non-nil, zero value otherwise.

### GetPlatformListingIdOk

`func (o *ListOrdersItem) GetPlatformListingIdOk() (*string, bool)`

GetPlatformListingIdOk returns a tuple with the PlatformListingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformListingId

`func (o *ListOrdersItem) SetPlatformListingId(v string)`

SetPlatformListingId sets PlatformListingId field to given value.

### HasPlatformListingId

`func (o *ListOrdersItem) HasPlatformListingId() bool`

HasPlatformListingId returns a boolean if a field has been set.

### SetPlatformListingIdNil

`func (o *ListOrdersItem) SetPlatformListingIdNil(b bool)`

 SetPlatformListingIdNil sets the value for PlatformListingId to be an explicit nil

### UnsetPlatformListingId
`func (o *ListOrdersItem) UnsetPlatformListingId()`

UnsetPlatformListingId ensures that no value is present for PlatformListingId, not even an explicit nil
### GetHandlingTimeDays

`func (o *ListOrdersItem) GetHandlingTimeDays() float32`

GetHandlingTimeDays returns the HandlingTimeDays field if non-nil, zero value otherwise.

### GetHandlingTimeDaysOk

`func (o *ListOrdersItem) GetHandlingTimeDaysOk() (*float32, bool)`

GetHandlingTimeDaysOk returns a tuple with the HandlingTimeDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHandlingTimeDays

`func (o *ListOrdersItem) SetHandlingTimeDays(v float32)`

SetHandlingTimeDays sets HandlingTimeDays field to given value.

### HasHandlingTimeDays

`func (o *ListOrdersItem) HasHandlingTimeDays() bool`

HasHandlingTimeDays returns a boolean if a field has been set.

### SetHandlingTimeDaysNil

`func (o *ListOrdersItem) SetHandlingTimeDaysNil(b bool)`

 SetHandlingTimeDaysNil sets the value for HandlingTimeDays to be an explicit nil

### UnsetHandlingTimeDays
`func (o *ListOrdersItem) UnsetHandlingTimeDays()`

UnsetHandlingTimeDays ensures that no value is present for HandlingTimeDays, not even an explicit nil
### GetPlatformOrderId

`func (o *ListOrdersItem) GetPlatformOrderId() string`

GetPlatformOrderId returns the PlatformOrderId field if non-nil, zero value otherwise.

### GetPlatformOrderIdOk

`func (o *ListOrdersItem) GetPlatformOrderIdOk() (*string, bool)`

GetPlatformOrderIdOk returns a tuple with the PlatformOrderId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatformOrderId

`func (o *ListOrdersItem) SetPlatformOrderId(v string)`

SetPlatformOrderId sets PlatformOrderId field to given value.

### HasPlatformOrderId

`func (o *ListOrdersItem) HasPlatformOrderId() bool`

HasPlatformOrderId returns a boolean if a field has been set.

### SetPlatformOrderIdNil

`func (o *ListOrdersItem) SetPlatformOrderIdNil(b bool)`

 SetPlatformOrderIdNil sets the value for PlatformOrderId to be an explicit nil

### UnsetPlatformOrderId
`func (o *ListOrdersItem) UnsetPlatformOrderId()`

UnsetPlatformOrderId ensures that no value is present for PlatformOrderId, not even an explicit nil
### GetBuyerUsername

`func (o *ListOrdersItem) GetBuyerUsername() string`

GetBuyerUsername returns the BuyerUsername field if non-nil, zero value otherwise.

### GetBuyerUsernameOk

`func (o *ListOrdersItem) GetBuyerUsernameOk() (*string, bool)`

GetBuyerUsernameOk returns a tuple with the BuyerUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerUsername

`func (o *ListOrdersItem) SetBuyerUsername(v string)`

SetBuyerUsername sets BuyerUsername field to given value.

### HasBuyerUsername

`func (o *ListOrdersItem) HasBuyerUsername() bool`

HasBuyerUsername returns a boolean if a field has been set.

### SetBuyerUsernameNil

`func (o *ListOrdersItem) SetBuyerUsernameNil(b bool)`

 SetBuyerUsernameNil sets the value for BuyerUsername to be an explicit nil

### UnsetBuyerUsername
`func (o *ListOrdersItem) UnsetBuyerUsername()`

UnsetBuyerUsername ensures that no value is present for BuyerUsername, not even an explicit nil
### GetSalesChannel

`func (o *ListOrdersItem) GetSalesChannel() string`

GetSalesChannel returns the SalesChannel field if non-nil, zero value otherwise.

### GetSalesChannelOk

`func (o *ListOrdersItem) GetSalesChannelOk() (*string, bool)`

GetSalesChannelOk returns a tuple with the SalesChannel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSalesChannel

`func (o *ListOrdersItem) SetSalesChannel(v string)`

SetSalesChannel sets SalesChannel field to given value.


### GetChannelLocationId

`func (o *ListOrdersItem) GetChannelLocationId() string`

GetChannelLocationId returns the ChannelLocationId field if non-nil, zero value otherwise.

### GetChannelLocationIdOk

`func (o *ListOrdersItem) GetChannelLocationIdOk() (*string, bool)`

GetChannelLocationIdOk returns a tuple with the ChannelLocationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChannelLocationId

`func (o *ListOrdersItem) SetChannelLocationId(v string)`

SetChannelLocationId sets ChannelLocationId field to given value.

### HasChannelLocationId

`func (o *ListOrdersItem) HasChannelLocationId() bool`

HasChannelLocationId returns a boolean if a field has been set.

### SetChannelLocationIdNil

`func (o *ListOrdersItem) SetChannelLocationIdNil(b bool)`

 SetChannelLocationIdNil sets the value for ChannelLocationId to be an explicit nil

### UnsetChannelLocationId
`func (o *ListOrdersItem) UnsetChannelLocationId()`

UnsetChannelLocationId ensures that no value is present for ChannelLocationId, not even an explicit nil
### GetPackagePresetId

`func (o *ListOrdersItem) GetPackagePresetId() string`

GetPackagePresetId returns the PackagePresetId field if non-nil, zero value otherwise.

### GetPackagePresetIdOk

`func (o *ListOrdersItem) GetPackagePresetIdOk() (*string, bool)`

GetPackagePresetIdOk returns a tuple with the PackagePresetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackagePresetId

`func (o *ListOrdersItem) SetPackagePresetId(v string)`

SetPackagePresetId sets PackagePresetId field to given value.

### HasPackagePresetId

`func (o *ListOrdersItem) HasPackagePresetId() bool`

HasPackagePresetId returns a boolean if a field has been set.

### SetPackagePresetIdNil

`func (o *ListOrdersItem) SetPackagePresetIdNil(b bool)`

 SetPackagePresetIdNil sets the value for PackagePresetId to be an explicit nil

### UnsetPackagePresetId
`func (o *ListOrdersItem) UnsetPackagePresetId()`

UnsetPackagePresetId ensures that no value is present for PackagePresetId, not even an explicit nil
### GetCarrier

`func (o *ListOrdersItem) GetCarrier() string`

GetCarrier returns the Carrier field if non-nil, zero value otherwise.

### GetCarrierOk

`func (o *ListOrdersItem) GetCarrierOk() (*string, bool)`

GetCarrierOk returns a tuple with the Carrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrier

`func (o *ListOrdersItem) SetCarrier(v string)`

SetCarrier sets Carrier field to given value.

### HasCarrier

`func (o *ListOrdersItem) HasCarrier() bool`

HasCarrier returns a boolean if a field has been set.

### SetCarrierNil

`func (o *ListOrdersItem) SetCarrierNil(b bool)`

 SetCarrierNil sets the value for Carrier to be an explicit nil

### UnsetCarrier
`func (o *ListOrdersItem) UnsetCarrier()`

UnsetCarrier ensures that no value is present for Carrier, not even an explicit nil
### GetService

`func (o *ListOrdersItem) GetService() string`

GetService returns the Service field if non-nil, zero value otherwise.

### GetServiceOk

`func (o *ListOrdersItem) GetServiceOk() (*string, bool)`

GetServiceOk returns a tuple with the Service field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetService

`func (o *ListOrdersItem) SetService(v string)`

SetService sets Service field to given value.

### HasService

`func (o *ListOrdersItem) HasService() bool`

HasService returns a boolean if a field has been set.

### SetServiceNil

`func (o *ListOrdersItem) SetServiceNil(b bool)`

 SetServiceNil sets the value for Service to be an explicit nil

### UnsetService
`func (o *ListOrdersItem) UnsetService()`

UnsetService ensures that no value is present for Service, not even an explicit nil
### GetTrackingNumber

`func (o *ListOrdersItem) GetTrackingNumber() string`

GetTrackingNumber returns the TrackingNumber field if non-nil, zero value otherwise.

### GetTrackingNumberOk

`func (o *ListOrdersItem) GetTrackingNumberOk() (*string, bool)`

GetTrackingNumberOk returns a tuple with the TrackingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingNumber

`func (o *ListOrdersItem) SetTrackingNumber(v string)`

SetTrackingNumber sets TrackingNumber field to given value.

### HasTrackingNumber

`func (o *ListOrdersItem) HasTrackingNumber() bool`

HasTrackingNumber returns a boolean if a field has been set.

### SetTrackingNumberNil

`func (o *ListOrdersItem) SetTrackingNumberNil(b bool)`

 SetTrackingNumberNil sets the value for TrackingNumber to be an explicit nil

### UnsetTrackingNumber
`func (o *ListOrdersItem) UnsetTrackingNumber()`

UnsetTrackingNumber ensures that no value is present for TrackingNumber, not even an explicit nil
### GetEasypostShipmentId

`func (o *ListOrdersItem) GetEasypostShipmentId() string`

GetEasypostShipmentId returns the EasypostShipmentId field if non-nil, zero value otherwise.

### GetEasypostShipmentIdOk

`func (o *ListOrdersItem) GetEasypostShipmentIdOk() (*string, bool)`

GetEasypostShipmentIdOk returns a tuple with the EasypostShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasypostShipmentId

`func (o *ListOrdersItem) SetEasypostShipmentId(v string)`

SetEasypostShipmentId sets EasypostShipmentId field to given value.

### HasEasypostShipmentId

`func (o *ListOrdersItem) HasEasypostShipmentId() bool`

HasEasypostShipmentId returns a boolean if a field has been set.

### SetEasypostShipmentIdNil

`func (o *ListOrdersItem) SetEasypostShipmentIdNil(b bool)`

 SetEasypostShipmentIdNil sets the value for EasypostShipmentId to be an explicit nil

### UnsetEasypostShipmentId
`func (o *ListOrdersItem) UnsetEasypostShipmentId()`

UnsetEasypostShipmentId ensures that no value is present for EasypostShipmentId, not even an explicit nil
### GetEasypostTrackerId

`func (o *ListOrdersItem) GetEasypostTrackerId() string`

GetEasypostTrackerId returns the EasypostTrackerId field if non-nil, zero value otherwise.

### GetEasypostTrackerIdOk

`func (o *ListOrdersItem) GetEasypostTrackerIdOk() (*string, bool)`

GetEasypostTrackerIdOk returns a tuple with the EasypostTrackerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasypostTrackerId

`func (o *ListOrdersItem) SetEasypostTrackerId(v string)`

SetEasypostTrackerId sets EasypostTrackerId field to given value.

### HasEasypostTrackerId

`func (o *ListOrdersItem) HasEasypostTrackerId() bool`

HasEasypostTrackerId returns a boolean if a field has been set.

### SetEasypostTrackerIdNil

`func (o *ListOrdersItem) SetEasypostTrackerIdNil(b bool)`

 SetEasypostTrackerIdNil sets the value for EasypostTrackerId to be an explicit nil

### UnsetEasypostTrackerId
`func (o *ListOrdersItem) UnsetEasypostTrackerId()`

UnsetEasypostTrackerId ensures that no value is present for EasypostTrackerId, not even an explicit nil
### GetEasypostRateId

`func (o *ListOrdersItem) GetEasypostRateId() string`

GetEasypostRateId returns the EasypostRateId field if non-nil, zero value otherwise.

### GetEasypostRateIdOk

`func (o *ListOrdersItem) GetEasypostRateIdOk() (*string, bool)`

GetEasypostRateIdOk returns a tuple with the EasypostRateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEasypostRateId

`func (o *ListOrdersItem) SetEasypostRateId(v string)`

SetEasypostRateId sets EasypostRateId field to given value.

### HasEasypostRateId

`func (o *ListOrdersItem) HasEasypostRateId() bool`

HasEasypostRateId returns a boolean if a field has been set.

### SetEasypostRateIdNil

`func (o *ListOrdersItem) SetEasypostRateIdNil(b bool)`

 SetEasypostRateIdNil sets the value for EasypostRateId to be an explicit nil

### UnsetEasypostRateId
`func (o *ListOrdersItem) UnsetEasypostRateId()`

UnsetEasypostRateId ensures that no value is present for EasypostRateId, not even an explicit nil
### GetShippingLabelUrl

`func (o *ListOrdersItem) GetShippingLabelUrl() string`

GetShippingLabelUrl returns the ShippingLabelUrl field if non-nil, zero value otherwise.

### GetShippingLabelUrlOk

`func (o *ListOrdersItem) GetShippingLabelUrlOk() (*string, bool)`

GetShippingLabelUrlOk returns a tuple with the ShippingLabelUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippingLabelUrl

`func (o *ListOrdersItem) SetShippingLabelUrl(v string)`

SetShippingLabelUrl sets ShippingLabelUrl field to given value.

### HasShippingLabelUrl

`func (o *ListOrdersItem) HasShippingLabelUrl() bool`

HasShippingLabelUrl returns a boolean if a field has been set.

### SetShippingLabelUrlNil

`func (o *ListOrdersItem) SetShippingLabelUrlNil(b bool)`

 SetShippingLabelUrlNil sets the value for ShippingLabelUrl to be an explicit nil

### UnsetShippingLabelUrl
`func (o *ListOrdersItem) UnsetShippingLabelUrl()`

UnsetShippingLabelUrl ensures that no value is present for ShippingLabelUrl, not even an explicit nil
### GetLabelPurchasedAt

`func (o *ListOrdersItem) GetLabelPurchasedAt() time.Time`

GetLabelPurchasedAt returns the LabelPurchasedAt field if non-nil, zero value otherwise.

### GetLabelPurchasedAtOk

`func (o *ListOrdersItem) GetLabelPurchasedAtOk() (*time.Time, bool)`

GetLabelPurchasedAtOk returns a tuple with the LabelPurchasedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelPurchasedAt

`func (o *ListOrdersItem) SetLabelPurchasedAt(v time.Time)`

SetLabelPurchasedAt sets LabelPurchasedAt field to given value.

### HasLabelPurchasedAt

`func (o *ListOrdersItem) HasLabelPurchasedAt() bool`

HasLabelPurchasedAt returns a boolean if a field has been set.

### SetLabelPurchasedAtNil

`func (o *ListOrdersItem) SetLabelPurchasedAtNil(b bool)`

 SetLabelPurchasedAtNil sets the value for LabelPurchasedAt to be an explicit nil

### UnsetLabelPurchasedAt
`func (o *ListOrdersItem) UnsetLabelPurchasedAt()`

UnsetLabelPurchasedAt ensures that no value is present for LabelPurchasedAt, not even an explicit nil
### GetShippedAt

`func (o *ListOrdersItem) GetShippedAt() time.Time`

GetShippedAt returns the ShippedAt field if non-nil, zero value otherwise.

### GetShippedAtOk

`func (o *ListOrdersItem) GetShippedAtOk() (*time.Time, bool)`

GetShippedAtOk returns a tuple with the ShippedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShippedAt

`func (o *ListOrdersItem) SetShippedAt(v time.Time)`

SetShippedAt sets ShippedAt field to given value.

### HasShippedAt

`func (o *ListOrdersItem) HasShippedAt() bool`

HasShippedAt returns a boolean if a field has been set.

### SetShippedAtNil

`func (o *ListOrdersItem) SetShippedAtNil(b bool)`

 SetShippedAtNil sets the value for ShippedAt to be an explicit nil

### UnsetShippedAt
`func (o *ListOrdersItem) UnsetShippedAt()`

UnsetShippedAt ensures that no value is present for ShippedAt, not even an explicit nil
### GetEstimatedDelivery

`func (o *ListOrdersItem) GetEstimatedDelivery() string`

GetEstimatedDelivery returns the EstimatedDelivery field if non-nil, zero value otherwise.

### GetEstimatedDeliveryOk

`func (o *ListOrdersItem) GetEstimatedDeliveryOk() (*string, bool)`

GetEstimatedDeliveryOk returns a tuple with the EstimatedDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDelivery

`func (o *ListOrdersItem) SetEstimatedDelivery(v string)`

SetEstimatedDelivery sets EstimatedDelivery field to given value.

### HasEstimatedDelivery

`func (o *ListOrdersItem) HasEstimatedDelivery() bool`

HasEstimatedDelivery returns a boolean if a field has been set.

### SetEstimatedDeliveryNil

`func (o *ListOrdersItem) SetEstimatedDeliveryNil(b bool)`

 SetEstimatedDeliveryNil sets the value for EstimatedDelivery to be an explicit nil

### UnsetEstimatedDelivery
`func (o *ListOrdersItem) UnsetEstimatedDelivery()`

UnsetEstimatedDelivery ensures that no value is present for EstimatedDelivery, not even an explicit nil
### GetDeliveredAt

`func (o *ListOrdersItem) GetDeliveredAt() time.Time`

GetDeliveredAt returns the DeliveredAt field if non-nil, zero value otherwise.

### GetDeliveredAtOk

`func (o *ListOrdersItem) GetDeliveredAtOk() (*time.Time, bool)`

GetDeliveredAtOk returns a tuple with the DeliveredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveredAt

`func (o *ListOrdersItem) SetDeliveredAt(v time.Time)`

SetDeliveredAt sets DeliveredAt field to given value.

### HasDeliveredAt

`func (o *ListOrdersItem) HasDeliveredAt() bool`

HasDeliveredAt returns a boolean if a field has been set.

### SetDeliveredAtNil

`func (o *ListOrdersItem) SetDeliveredAtNil(b bool)`

 SetDeliveredAtNil sets the value for DeliveredAt to be an explicit nil

### UnsetDeliveredAt
`func (o *ListOrdersItem) UnsetDeliveredAt()`

UnsetDeliveredAt ensures that no value is present for DeliveredAt, not even an explicit nil
### GetCarrierStatus

`func (o *ListOrdersItem) GetCarrierStatus() string`

GetCarrierStatus returns the CarrierStatus field if non-nil, zero value otherwise.

### GetCarrierStatusOk

`func (o *ListOrdersItem) GetCarrierStatusOk() (*string, bool)`

GetCarrierStatusOk returns a tuple with the CarrierStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierStatus

`func (o *ListOrdersItem) SetCarrierStatus(v string)`

SetCarrierStatus sets CarrierStatus field to given value.

### HasCarrierStatus

`func (o *ListOrdersItem) HasCarrierStatus() bool`

HasCarrierStatus returns a boolean if a field has been set.

### SetCarrierStatusNil

`func (o *ListOrdersItem) SetCarrierStatusNil(b bool)`

 SetCarrierStatusNil sets the value for CarrierStatus to be an explicit nil

### UnsetCarrierStatus
`func (o *ListOrdersItem) UnsetCarrierStatus()`

UnsetCarrierStatus ensures that no value is present for CarrierStatus, not even an explicit nil
### GetCarrierStatusDetail

`func (o *ListOrdersItem) GetCarrierStatusDetail() string`

GetCarrierStatusDetail returns the CarrierStatusDetail field if non-nil, zero value otherwise.

### GetCarrierStatusDetailOk

`func (o *ListOrdersItem) GetCarrierStatusDetailOk() (*string, bool)`

GetCarrierStatusDetailOk returns a tuple with the CarrierStatusDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCarrierStatusDetail

`func (o *ListOrdersItem) SetCarrierStatusDetail(v string)`

SetCarrierStatusDetail sets CarrierStatusDetail field to given value.

### HasCarrierStatusDetail

`func (o *ListOrdersItem) HasCarrierStatusDetail() bool`

HasCarrierStatusDetail returns a boolean if a field has been set.

### SetCarrierStatusDetailNil

`func (o *ListOrdersItem) SetCarrierStatusDetailNil(b bool)`

 SetCarrierStatusDetailNil sets the value for CarrierStatusDetail to be an explicit nil

### UnsetCarrierStatusDetail
`func (o *ListOrdersItem) UnsetCarrierStatusDetail()`

UnsetCarrierStatusDetail ensures that no value is present for CarrierStatusDetail, not even an explicit nil
### GetTrackingHistory

`func (o *ListOrdersItem) GetTrackingHistory() []ListOrdersItemTrackingHistory`

GetTrackingHistory returns the TrackingHistory field if non-nil, zero value otherwise.

### GetTrackingHistoryOk

`func (o *ListOrdersItem) GetTrackingHistoryOk() (*[]ListOrdersItemTrackingHistory, bool)`

GetTrackingHistoryOk returns a tuple with the TrackingHistory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingHistory

`func (o *ListOrdersItem) SetTrackingHistory(v []ListOrdersItemTrackingHistory)`

SetTrackingHistory sets TrackingHistory field to given value.

### HasTrackingHistory

`func (o *ListOrdersItem) HasTrackingHistory() bool`

HasTrackingHistory returns a boolean if a field has been set.

### SetTrackingHistoryNil

`func (o *ListOrdersItem) SetTrackingHistoryNil(b bool)`

 SetTrackingHistoryNil sets the value for TrackingHistory to be an explicit nil

### UnsetTrackingHistory
`func (o *ListOrdersItem) UnsetTrackingHistory()`

UnsetTrackingHistory ensures that no value is present for TrackingHistory, not even an explicit nil
### GetDeliveryLocation

`func (o *ListOrdersItem) GetDeliveryLocation() string`

GetDeliveryLocation returns the DeliveryLocation field if non-nil, zero value otherwise.

### GetDeliveryLocationOk

`func (o *ListOrdersItem) GetDeliveryLocationOk() (*string, bool)`

GetDeliveryLocationOk returns a tuple with the DeliveryLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryLocation

`func (o *ListOrdersItem) SetDeliveryLocation(v string)`

SetDeliveryLocation sets DeliveryLocation field to given value.

### HasDeliveryLocation

`func (o *ListOrdersItem) HasDeliveryLocation() bool`

HasDeliveryLocation returns a boolean if a field has been set.

### SetDeliveryLocationNil

`func (o *ListOrdersItem) SetDeliveryLocationNil(b bool)`

 SetDeliveryLocationNil sets the value for DeliveryLocation to be an explicit nil

### UnsetDeliveryLocation
`func (o *ListOrdersItem) UnsetDeliveryLocation()`

UnsetDeliveryLocation ensures that no value is present for DeliveryLocation, not even an explicit nil
### GetDeliverySignature

`func (o *ListOrdersItem) GetDeliverySignature() string`

GetDeliverySignature returns the DeliverySignature field if non-nil, zero value otherwise.

### GetDeliverySignatureOk

`func (o *ListOrdersItem) GetDeliverySignatureOk() (*string, bool)`

GetDeliverySignatureOk returns a tuple with the DeliverySignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliverySignature

`func (o *ListOrdersItem) SetDeliverySignature(v string)`

SetDeliverySignature sets DeliverySignature field to given value.

### HasDeliverySignature

`func (o *ListOrdersItem) HasDeliverySignature() bool`

HasDeliverySignature returns a boolean if a field has been set.

### SetDeliverySignatureNil

`func (o *ListOrdersItem) SetDeliverySignatureNil(b bool)`

 SetDeliverySignatureNil sets the value for DeliverySignature to be an explicit nil

### UnsetDeliverySignature
`func (o *ListOrdersItem) UnsetDeliverySignature()`

UnsetDeliverySignature ensures that no value is present for DeliverySignature, not even an explicit nil
### GetTrackingSubmittedAt

`func (o *ListOrdersItem) GetTrackingSubmittedAt() time.Time`

GetTrackingSubmittedAt returns the TrackingSubmittedAt field if non-nil, zero value otherwise.

### GetTrackingSubmittedAtOk

`func (o *ListOrdersItem) GetTrackingSubmittedAtOk() (*time.Time, bool)`

GetTrackingSubmittedAtOk returns a tuple with the TrackingSubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingSubmittedAt

`func (o *ListOrdersItem) SetTrackingSubmittedAt(v time.Time)`

SetTrackingSubmittedAt sets TrackingSubmittedAt field to given value.

### HasTrackingSubmittedAt

`func (o *ListOrdersItem) HasTrackingSubmittedAt() bool`

HasTrackingSubmittedAt returns a boolean if a field has been set.

### SetTrackingSubmittedAtNil

`func (o *ListOrdersItem) SetTrackingSubmittedAtNil(b bool)`

 SetTrackingSubmittedAtNil sets the value for TrackingSubmittedAt to be an explicit nil

### UnsetTrackingSubmittedAt
`func (o *ListOrdersItem) UnsetTrackingSubmittedAt()`

UnsetTrackingSubmittedAt ensures that no value is present for TrackingSubmittedAt, not even an explicit nil
### GetTrackingSubmitStatus

`func (o *ListOrdersItem) GetTrackingSubmitStatus() string`

GetTrackingSubmitStatus returns the TrackingSubmitStatus field if non-nil, zero value otherwise.

### GetTrackingSubmitStatusOk

`func (o *ListOrdersItem) GetTrackingSubmitStatusOk() (*string, bool)`

GetTrackingSubmitStatusOk returns a tuple with the TrackingSubmitStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrackingSubmitStatus

`func (o *ListOrdersItem) SetTrackingSubmitStatus(v string)`

SetTrackingSubmitStatus sets TrackingSubmitStatus field to given value.

### HasTrackingSubmitStatus

`func (o *ListOrdersItem) HasTrackingSubmitStatus() bool`

HasTrackingSubmitStatus returns a boolean if a field has been set.

### SetTrackingSubmitStatusNil

`func (o *ListOrdersItem) SetTrackingSubmitStatusNil(b bool)`

 SetTrackingSubmitStatusNil sets the value for TrackingSubmitStatus to be an explicit nil

### UnsetTrackingSubmitStatus
`func (o *ListOrdersItem) UnsetTrackingSubmitStatus()`

UnsetTrackingSubmitStatus ensures that no value is present for TrackingSubmitStatus, not even an explicit nil
### GetLabelCost

`func (o *ListOrdersItem) GetLabelCost() string`

GetLabelCost returns the LabelCost field if non-nil, zero value otherwise.

### GetLabelCostOk

`func (o *ListOrdersItem) GetLabelCostOk() (*string, bool)`

GetLabelCostOk returns a tuple with the LabelCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabelCost

`func (o *ListOrdersItem) SetLabelCost(v string)`

SetLabelCost sets LabelCost field to given value.

### HasLabelCost

`func (o *ListOrdersItem) HasLabelCost() bool`

HasLabelCost returns a boolean if a field has been set.

### SetLabelCostNil

`func (o *ListOrdersItem) SetLabelCostNil(b bool)`

 SetLabelCostNil sets the value for LabelCost to be an explicit nil

### UnsetLabelCost
`func (o *ListOrdersItem) UnsetLabelCost()`

UnsetLabelCost ensures that no value is present for LabelCost, not even an explicit nil
### GetCostOfGoods

`func (o *ListOrdersItem) GetCostOfGoods() string`

GetCostOfGoods returns the CostOfGoods field if non-nil, zero value otherwise.

### GetCostOfGoodsOk

`func (o *ListOrdersItem) GetCostOfGoodsOk() (*string, bool)`

GetCostOfGoodsOk returns a tuple with the CostOfGoods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCostOfGoods

`func (o *ListOrdersItem) SetCostOfGoods(v string)`

SetCostOfGoods sets CostOfGoods field to given value.

### HasCostOfGoods

`func (o *ListOrdersItem) HasCostOfGoods() bool`

HasCostOfGoods returns a boolean if a field has been set.

### SetCostOfGoodsNil

`func (o *ListOrdersItem) SetCostOfGoodsNil(b bool)`

 SetCostOfGoodsNil sets the value for CostOfGoods to be an explicit nil

### UnsetCostOfGoods
`func (o *ListOrdersItem) UnsetCostOfGoods()`

UnsetCostOfGoods ensures that no value is present for CostOfGoods, not even an explicit nil
### GetRequestedCarrier

`func (o *ListOrdersItem) GetRequestedCarrier() string`

GetRequestedCarrier returns the RequestedCarrier field if non-nil, zero value otherwise.

### GetRequestedCarrierOk

`func (o *ListOrdersItem) GetRequestedCarrierOk() (*string, bool)`

GetRequestedCarrierOk returns a tuple with the RequestedCarrier field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedCarrier

`func (o *ListOrdersItem) SetRequestedCarrier(v string)`

SetRequestedCarrier sets RequestedCarrier field to given value.

### HasRequestedCarrier

`func (o *ListOrdersItem) HasRequestedCarrier() bool`

HasRequestedCarrier returns a boolean if a field has been set.

### SetRequestedCarrierNil

`func (o *ListOrdersItem) SetRequestedCarrierNil(b bool)`

 SetRequestedCarrierNil sets the value for RequestedCarrier to be an explicit nil

### UnsetRequestedCarrier
`func (o *ListOrdersItem) UnsetRequestedCarrier()`

UnsetRequestedCarrier ensures that no value is present for RequestedCarrier, not even an explicit nil
### GetRequestedService

`func (o *ListOrdersItem) GetRequestedService() string`

GetRequestedService returns the RequestedService field if non-nil, zero value otherwise.

### GetRequestedServiceOk

`func (o *ListOrdersItem) GetRequestedServiceOk() (*string, bool)`

GetRequestedServiceOk returns a tuple with the RequestedService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedService

`func (o *ListOrdersItem) SetRequestedService(v string)`

SetRequestedService sets RequestedService field to given value.

### HasRequestedService

`func (o *ListOrdersItem) HasRequestedService() bool`

HasRequestedService returns a boolean if a field has been set.

### SetRequestedServiceNil

`func (o *ListOrdersItem) SetRequestedServiceNil(b bool)`

 SetRequestedServiceNil sets the value for RequestedService to be an explicit nil

### UnsetRequestedService
`func (o *ListOrdersItem) UnsetRequestedService()`

UnsetRequestedService ensures that no value is present for RequestedService, not even an explicit nil
### GetShipByAt

`func (o *ListOrdersItem) GetShipByAt() time.Time`

GetShipByAt returns the ShipByAt field if non-nil, zero value otherwise.

### GetShipByAtOk

`func (o *ListOrdersItem) GetShipByAtOk() (*time.Time, bool)`

GetShipByAtOk returns a tuple with the ShipByAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipByAt

`func (o *ListOrdersItem) SetShipByAt(v time.Time)`

SetShipByAt sets ShipByAt field to given value.

### HasShipByAt

`func (o *ListOrdersItem) HasShipByAt() bool`

HasShipByAt returns a boolean if a field has been set.

### SetShipByAtNil

`func (o *ListOrdersItem) SetShipByAtNil(b bool)`

 SetShipByAtNil sets the value for ShipByAt to be an explicit nil

### UnsetShipByAt
`func (o *ListOrdersItem) UnsetShipByAt()`

UnsetShipByAt ensures that no value is present for ShipByAt, not even an explicit nil
### GetShipByAlertedAt

`func (o *ListOrdersItem) GetShipByAlertedAt() time.Time`

GetShipByAlertedAt returns the ShipByAlertedAt field if non-nil, zero value otherwise.

### GetShipByAlertedAtOk

`func (o *ListOrdersItem) GetShipByAlertedAtOk() (*time.Time, bool)`

GetShipByAlertedAtOk returns a tuple with the ShipByAlertedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipByAlertedAt

`func (o *ListOrdersItem) SetShipByAlertedAt(v time.Time)`

SetShipByAlertedAt sets ShipByAlertedAt field to given value.

### HasShipByAlertedAt

`func (o *ListOrdersItem) HasShipByAlertedAt() bool`

HasShipByAlertedAt returns a boolean if a field has been set.

### SetShipByAlertedAtNil

`func (o *ListOrdersItem) SetShipByAlertedAtNil(b bool)`

 SetShipByAlertedAtNil sets the value for ShipByAlertedAt to be an explicit nil

### UnsetShipByAlertedAt
`func (o *ListOrdersItem) UnsetShipByAlertedAt()`

UnsetShipByAlertedAt ensures that no value is present for ShipByAlertedAt, not even an explicit nil
### GetOversoldBy

`func (o *ListOrdersItem) GetOversoldBy() float32`

GetOversoldBy returns the OversoldBy field if non-nil, zero value otherwise.

### GetOversoldByOk

`func (o *ListOrdersItem) GetOversoldByOk() (*float32, bool)`

GetOversoldByOk returns a tuple with the OversoldBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOversoldBy

`func (o *ListOrdersItem) SetOversoldBy(v float32)`

SetOversoldBy sets OversoldBy field to given value.


### GetLastStatusCheckAt

`func (o *ListOrdersItem) GetLastStatusCheckAt() time.Time`

GetLastStatusCheckAt returns the LastStatusCheckAt field if non-nil, zero value otherwise.

### GetLastStatusCheckAtOk

`func (o *ListOrdersItem) GetLastStatusCheckAtOk() (*time.Time, bool)`

GetLastStatusCheckAtOk returns a tuple with the LastStatusCheckAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastStatusCheckAt

`func (o *ListOrdersItem) SetLastStatusCheckAt(v time.Time)`

SetLastStatusCheckAt sets LastStatusCheckAt field to given value.

### HasLastStatusCheckAt

`func (o *ListOrdersItem) HasLastStatusCheckAt() bool`

HasLastStatusCheckAt returns a boolean if a field has been set.

### SetLastStatusCheckAtNil

`func (o *ListOrdersItem) SetLastStatusCheckAtNil(b bool)`

 SetLastStatusCheckAtNil sets the value for LastStatusCheckAt to be an explicit nil

### UnsetLastStatusCheckAt
`func (o *ListOrdersItem) UnsetLastStatusCheckAt()`

UnsetLastStatusCheckAt ensures that no value is present for LastStatusCheckAt, not even an explicit nil
### GetLastChatCheckAt

`func (o *ListOrdersItem) GetLastChatCheckAt() time.Time`

GetLastChatCheckAt returns the LastChatCheckAt field if non-nil, zero value otherwise.

### GetLastChatCheckAtOk

`func (o *ListOrdersItem) GetLastChatCheckAtOk() (*time.Time, bool)`

GetLastChatCheckAtOk returns a tuple with the LastChatCheckAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastChatCheckAt

`func (o *ListOrdersItem) SetLastChatCheckAt(v time.Time)`

SetLastChatCheckAt sets LastChatCheckAt field to given value.

### HasLastChatCheckAt

`func (o *ListOrdersItem) HasLastChatCheckAt() bool`

HasLastChatCheckAt returns a boolean if a field has been set.

### SetLastChatCheckAtNil

`func (o *ListOrdersItem) SetLastChatCheckAtNil(b bool)`

 SetLastChatCheckAtNil sets the value for LastChatCheckAt to be an explicit nil

### UnsetLastChatCheckAt
`func (o *ListOrdersItem) UnsetLastChatCheckAt()`

UnsetLastChatCheckAt ensures that no value is present for LastChatCheckAt, not even an explicit nil
### GetIsDisputed

`func (o *ListOrdersItem) GetIsDisputed() bool`

GetIsDisputed returns the IsDisputed field if non-nil, zero value otherwise.

### GetIsDisputedOk

`func (o *ListOrdersItem) GetIsDisputedOk() (*bool, bool)`

GetIsDisputedOk returns a tuple with the IsDisputed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDisputed

`func (o *ListOrdersItem) SetIsDisputed(v bool)`

SetIsDisputed sets IsDisputed field to given value.


### GetDisputeReason

`func (o *ListOrdersItem) GetDisputeReason() string`

GetDisputeReason returns the DisputeReason field if non-nil, zero value otherwise.

### GetDisputeReasonOk

`func (o *ListOrdersItem) GetDisputeReasonOk() (*string, bool)`

GetDisputeReasonOk returns a tuple with the DisputeReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeReason

`func (o *ListOrdersItem) SetDisputeReason(v string)`

SetDisputeReason sets DisputeReason field to given value.

### HasDisputeReason

`func (o *ListOrdersItem) HasDisputeReason() bool`

HasDisputeReason returns a boolean if a field has been set.

### SetDisputeReasonNil

`func (o *ListOrdersItem) SetDisputeReasonNil(b bool)`

 SetDisputeReasonNil sets the value for DisputeReason to be an explicit nil

### UnsetDisputeReason
`func (o *ListOrdersItem) UnsetDisputeReason()`

UnsetDisputeReason ensures that no value is present for DisputeReason, not even an explicit nil
### GetDisputePlatformCaseId

`func (o *ListOrdersItem) GetDisputePlatformCaseId() string`

GetDisputePlatformCaseId returns the DisputePlatformCaseId field if non-nil, zero value otherwise.

### GetDisputePlatformCaseIdOk

`func (o *ListOrdersItem) GetDisputePlatformCaseIdOk() (*string, bool)`

GetDisputePlatformCaseIdOk returns a tuple with the DisputePlatformCaseId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputePlatformCaseId

`func (o *ListOrdersItem) SetDisputePlatformCaseId(v string)`

SetDisputePlatformCaseId sets DisputePlatformCaseId field to given value.

### HasDisputePlatformCaseId

`func (o *ListOrdersItem) HasDisputePlatformCaseId() bool`

HasDisputePlatformCaseId returns a boolean if a field has been set.

### SetDisputePlatformCaseIdNil

`func (o *ListOrdersItem) SetDisputePlatformCaseIdNil(b bool)`

 SetDisputePlatformCaseIdNil sets the value for DisputePlatformCaseId to be an explicit nil

### UnsetDisputePlatformCaseId
`func (o *ListOrdersItem) UnsetDisputePlatformCaseId()`

UnsetDisputePlatformCaseId ensures that no value is present for DisputePlatformCaseId, not even an explicit nil
### GetDisputeResolvedAt

`func (o *ListOrdersItem) GetDisputeResolvedAt() time.Time`

GetDisputeResolvedAt returns the DisputeResolvedAt field if non-nil, zero value otherwise.

### GetDisputeResolvedAtOk

`func (o *ListOrdersItem) GetDisputeResolvedAtOk() (*time.Time, bool)`

GetDisputeResolvedAtOk returns a tuple with the DisputeResolvedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisputeResolvedAt

`func (o *ListOrdersItem) SetDisputeResolvedAt(v time.Time)`

SetDisputeResolvedAt sets DisputeResolvedAt field to given value.

### HasDisputeResolvedAt

`func (o *ListOrdersItem) HasDisputeResolvedAt() bool`

HasDisputeResolvedAt returns a boolean if a field has been set.

### SetDisputeResolvedAtNil

`func (o *ListOrdersItem) SetDisputeResolvedAtNil(b bool)`

 SetDisputeResolvedAtNil sets the value for DisputeResolvedAt to be an explicit nil

### UnsetDisputeResolvedAt
`func (o *ListOrdersItem) UnsetDisputeResolvedAt()`

UnsetDisputeResolvedAt ensures that no value is present for DisputeResolvedAt, not even an explicit nil
### GetRefundAmount

`func (o *ListOrdersItem) GetRefundAmount() string`

GetRefundAmount returns the RefundAmount field if non-nil, zero value otherwise.

### GetRefundAmountOk

`func (o *ListOrdersItem) GetRefundAmountOk() (*string, bool)`

GetRefundAmountOk returns a tuple with the RefundAmount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundAmount

`func (o *ListOrdersItem) SetRefundAmount(v string)`

SetRefundAmount sets RefundAmount field to given value.

### HasRefundAmount

`func (o *ListOrdersItem) HasRefundAmount() bool`

HasRefundAmount returns a boolean if a field has been set.

### SetRefundAmountNil

`func (o *ListOrdersItem) SetRefundAmountNil(b bool)`

 SetRefundAmountNil sets the value for RefundAmount to be an explicit nil

### UnsetRefundAmount
`func (o *ListOrdersItem) UnsetRefundAmount()`

UnsetRefundAmount ensures that no value is present for RefundAmount, not even an explicit nil
### GetRefundReason

`func (o *ListOrdersItem) GetRefundReason() string`

GetRefundReason returns the RefundReason field if non-nil, zero value otherwise.

### GetRefundReasonOk

`func (o *ListOrdersItem) GetRefundReasonOk() (*string, bool)`

GetRefundReasonOk returns a tuple with the RefundReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundReason

`func (o *ListOrdersItem) SetRefundReason(v string)`

SetRefundReason sets RefundReason field to given value.

### HasRefundReason

`func (o *ListOrdersItem) HasRefundReason() bool`

HasRefundReason returns a boolean if a field has been set.

### SetRefundReasonNil

`func (o *ListOrdersItem) SetRefundReasonNil(b bool)`

 SetRefundReasonNil sets the value for RefundReason to be an explicit nil

### UnsetRefundReason
`func (o *ListOrdersItem) UnsetRefundReason()`

UnsetRefundReason ensures that no value is present for RefundReason, not even an explicit nil
### GetRefundPlatformId

`func (o *ListOrdersItem) GetRefundPlatformId() string`

GetRefundPlatformId returns the RefundPlatformId field if non-nil, zero value otherwise.

### GetRefundPlatformIdOk

`func (o *ListOrdersItem) GetRefundPlatformIdOk() (*string, bool)`

GetRefundPlatformIdOk returns a tuple with the RefundPlatformId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundPlatformId

`func (o *ListOrdersItem) SetRefundPlatformId(v string)`

SetRefundPlatformId sets RefundPlatformId field to given value.

### HasRefundPlatformId

`func (o *ListOrdersItem) HasRefundPlatformId() bool`

HasRefundPlatformId returns a boolean if a field has been set.

### SetRefundPlatformIdNil

`func (o *ListOrdersItem) SetRefundPlatformIdNil(b bool)`

 SetRefundPlatformIdNil sets the value for RefundPlatformId to be an explicit nil

### UnsetRefundPlatformId
`func (o *ListOrdersItem) UnsetRefundPlatformId()`

UnsetRefundPlatformId ensures that no value is present for RefundPlatformId, not even an explicit nil
### GetRefundedAt

`func (o *ListOrdersItem) GetRefundedAt() time.Time`

GetRefundedAt returns the RefundedAt field if non-nil, zero value otherwise.

### GetRefundedAtOk

`func (o *ListOrdersItem) GetRefundedAtOk() (*time.Time, bool)`

GetRefundedAtOk returns a tuple with the RefundedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefundedAt

`func (o *ListOrdersItem) SetRefundedAt(v time.Time)`

SetRefundedAt sets RefundedAt field to given value.

### HasRefundedAt

`func (o *ListOrdersItem) HasRefundedAt() bool`

HasRefundedAt returns a boolean if a field has been set.

### SetRefundedAtNil

`func (o *ListOrdersItem) SetRefundedAtNil(b bool)`

 SetRefundedAtNil sets the value for RefundedAt to be an explicit nil

### UnsetRefundedAt
`func (o *ListOrdersItem) UnsetRefundedAt()`

UnsetRefundedAt ensures that no value is present for RefundedAt, not even an explicit nil
### GetCancellationReason

`func (o *ListOrdersItem) GetCancellationReason() string`

GetCancellationReason returns the CancellationReason field if non-nil, zero value otherwise.

### GetCancellationReasonOk

`func (o *ListOrdersItem) GetCancellationReasonOk() (*string, bool)`

GetCancellationReasonOk returns a tuple with the CancellationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCancellationReason

`func (o *ListOrdersItem) SetCancellationReason(v string)`

SetCancellationReason sets CancellationReason field to given value.

### HasCancellationReason

`func (o *ListOrdersItem) HasCancellationReason() bool`

HasCancellationReason returns a boolean if a field has been set.

### SetCancellationReasonNil

`func (o *ListOrdersItem) SetCancellationReasonNil(b bool)`

 SetCancellationReasonNil sets the value for CancellationReason to be an explicit nil

### UnsetCancellationReason
`func (o *ListOrdersItem) UnsetCancellationReason()`

UnsetCancellationReason ensures that no value is present for CancellationReason, not even an explicit nil
### GetArrivalConditionRequestedAt

`func (o *ListOrdersItem) GetArrivalConditionRequestedAt() time.Time`

GetArrivalConditionRequestedAt returns the ArrivalConditionRequestedAt field if non-nil, zero value otherwise.

### GetArrivalConditionRequestedAtOk

`func (o *ListOrdersItem) GetArrivalConditionRequestedAtOk() (*time.Time, bool)`

GetArrivalConditionRequestedAtOk returns a tuple with the ArrivalConditionRequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalConditionRequestedAt

`func (o *ListOrdersItem) SetArrivalConditionRequestedAt(v time.Time)`

SetArrivalConditionRequestedAt sets ArrivalConditionRequestedAt field to given value.

### HasArrivalConditionRequestedAt

`func (o *ListOrdersItem) HasArrivalConditionRequestedAt() bool`

HasArrivalConditionRequestedAt returns a boolean if a field has been set.

### SetArrivalConditionRequestedAtNil

`func (o *ListOrdersItem) SetArrivalConditionRequestedAtNil(b bool)`

 SetArrivalConditionRequestedAtNil sets the value for ArrivalConditionRequestedAt to be an explicit nil

### UnsetArrivalConditionRequestedAt
`func (o *ListOrdersItem) UnsetArrivalConditionRequestedAt()`

UnsetArrivalConditionRequestedAt ensures that no value is present for ArrivalConditionRequestedAt, not even an explicit nil
### GetArrivalConditionSubmittedAt

`func (o *ListOrdersItem) GetArrivalConditionSubmittedAt() time.Time`

GetArrivalConditionSubmittedAt returns the ArrivalConditionSubmittedAt field if non-nil, zero value otherwise.

### GetArrivalConditionSubmittedAtOk

`func (o *ListOrdersItem) GetArrivalConditionSubmittedAtOk() (*time.Time, bool)`

GetArrivalConditionSubmittedAtOk returns a tuple with the ArrivalConditionSubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalConditionSubmittedAt

`func (o *ListOrdersItem) SetArrivalConditionSubmittedAt(v time.Time)`

SetArrivalConditionSubmittedAt sets ArrivalConditionSubmittedAt field to given value.

### HasArrivalConditionSubmittedAt

`func (o *ListOrdersItem) HasArrivalConditionSubmittedAt() bool`

HasArrivalConditionSubmittedAt returns a boolean if a field has been set.

### SetArrivalConditionSubmittedAtNil

`func (o *ListOrdersItem) SetArrivalConditionSubmittedAtNil(b bool)`

 SetArrivalConditionSubmittedAtNil sets the value for ArrivalConditionSubmittedAt to be an explicit nil

### UnsetArrivalConditionSubmittedAt
`func (o *ListOrdersItem) UnsetArrivalConditionSubmittedAt()`

UnsetArrivalConditionSubmittedAt ensures that no value is present for ArrivalConditionSubmittedAt, not even an explicit nil
### GetArrivalConditionDeclinedAt

`func (o *ListOrdersItem) GetArrivalConditionDeclinedAt() time.Time`

GetArrivalConditionDeclinedAt returns the ArrivalConditionDeclinedAt field if non-nil, zero value otherwise.

### GetArrivalConditionDeclinedAtOk

`func (o *ListOrdersItem) GetArrivalConditionDeclinedAtOk() (*time.Time, bool)`

GetArrivalConditionDeclinedAtOk returns a tuple with the ArrivalConditionDeclinedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalConditionDeclinedAt

`func (o *ListOrdersItem) SetArrivalConditionDeclinedAt(v time.Time)`

SetArrivalConditionDeclinedAt sets ArrivalConditionDeclinedAt field to given value.

### HasArrivalConditionDeclinedAt

`func (o *ListOrdersItem) HasArrivalConditionDeclinedAt() bool`

HasArrivalConditionDeclinedAt returns a boolean if a field has been set.

### SetArrivalConditionDeclinedAtNil

`func (o *ListOrdersItem) SetArrivalConditionDeclinedAtNil(b bool)`

 SetArrivalConditionDeclinedAtNil sets the value for ArrivalConditionDeclinedAt to be an explicit nil

### UnsetArrivalConditionDeclinedAt
`func (o *ListOrdersItem) UnsetArrivalConditionDeclinedAt()`

UnsetArrivalConditionDeclinedAt ensures that no value is present for ArrivalConditionDeclinedAt, not even an explicit nil
### GetDeliveryPhotoUrl

`func (o *ListOrdersItem) GetDeliveryPhotoUrl() string`

GetDeliveryPhotoUrl returns the DeliveryPhotoUrl field if non-nil, zero value otherwise.

### GetDeliveryPhotoUrlOk

`func (o *ListOrdersItem) GetDeliveryPhotoUrlOk() (*string, bool)`

GetDeliveryPhotoUrlOk returns a tuple with the DeliveryPhotoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryPhotoUrl

`func (o *ListOrdersItem) SetDeliveryPhotoUrl(v string)`

SetDeliveryPhotoUrl sets DeliveryPhotoUrl field to given value.

### HasDeliveryPhotoUrl

`func (o *ListOrdersItem) HasDeliveryPhotoUrl() bool`

HasDeliveryPhotoUrl returns a boolean if a field has been set.

### SetDeliveryPhotoUrlNil

`func (o *ListOrdersItem) SetDeliveryPhotoUrlNil(b bool)`

 SetDeliveryPhotoUrlNil sets the value for DeliveryPhotoUrl to be an explicit nil

### UnsetDeliveryPhotoUrl
`func (o *ListOrdersItem) UnsetDeliveryPhotoUrl()`

UnsetDeliveryPhotoUrl ensures that no value is present for DeliveryPhotoUrl, not even an explicit nil
### GetPurchaseOrderRef

`func (o *ListOrdersItem) GetPurchaseOrderRef() string`

GetPurchaseOrderRef returns the PurchaseOrderRef field if non-nil, zero value otherwise.

### GetPurchaseOrderRefOk

`func (o *ListOrdersItem) GetPurchaseOrderRefOk() (*string, bool)`

GetPurchaseOrderRefOk returns a tuple with the PurchaseOrderRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurchaseOrderRef

`func (o *ListOrdersItem) SetPurchaseOrderRef(v string)`

SetPurchaseOrderRef sets PurchaseOrderRef field to given value.

### HasPurchaseOrderRef

`func (o *ListOrdersItem) HasPurchaseOrderRef() bool`

HasPurchaseOrderRef returns a boolean if a field has been set.

### SetPurchaseOrderRefNil

`func (o *ListOrdersItem) SetPurchaseOrderRefNil(b bool)`

 SetPurchaseOrderRefNil sets the value for PurchaseOrderRef to be an explicit nil

### UnsetPurchaseOrderRef
`func (o *ListOrdersItem) UnsetPurchaseOrderRef()`

UnsetPurchaseOrderRef ensures that no value is present for PurchaseOrderRef, not even an explicit nil
### GetArrivalConditionPhotos

`func (o *ListOrdersItem) GetArrivalConditionPhotos() []string`

GetArrivalConditionPhotos returns the ArrivalConditionPhotos field if non-nil, zero value otherwise.

### GetArrivalConditionPhotosOk

`func (o *ListOrdersItem) GetArrivalConditionPhotosOk() (*[]string, bool)`

GetArrivalConditionPhotosOk returns a tuple with the ArrivalConditionPhotos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArrivalConditionPhotos

`func (o *ListOrdersItem) SetArrivalConditionPhotos(v []string)`

SetArrivalConditionPhotos sets ArrivalConditionPhotos field to given value.


### GetBuyerEmail

`func (o *ListOrdersItem) GetBuyerEmail() string`

GetBuyerEmail returns the BuyerEmail field if non-nil, zero value otherwise.

### GetBuyerEmailOk

`func (o *ListOrdersItem) GetBuyerEmailOk() (*string, bool)`

GetBuyerEmailOk returns a tuple with the BuyerEmail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuyerEmail

`func (o *ListOrdersItem) SetBuyerEmail(v string)`

SetBuyerEmail sets BuyerEmail field to given value.

### HasBuyerEmail

`func (o *ListOrdersItem) HasBuyerEmail() bool`

HasBuyerEmail returns a boolean if a field has been set.

### SetBuyerEmailNil

`func (o *ListOrdersItem) SetBuyerEmailNil(b bool)`

 SetBuyerEmailNil sets the value for BuyerEmail to be an explicit nil

### UnsetBuyerEmail
`func (o *ListOrdersItem) UnsetBuyerEmail()`

UnsetBuyerEmail ensures that no value is present for BuyerEmail, not even an explicit nil
### GetFulfillmentMethod

`func (o *ListOrdersItem) GetFulfillmentMethod() string`

GetFulfillmentMethod returns the FulfillmentMethod field if non-nil, zero value otherwise.

### GetFulfillmentMethodOk

`func (o *ListOrdersItem) GetFulfillmentMethodOk() (*string, bool)`

GetFulfillmentMethodOk returns a tuple with the FulfillmentMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFulfillmentMethod

`func (o *ListOrdersItem) SetFulfillmentMethod(v string)`

SetFulfillmentMethod sets FulfillmentMethod field to given value.


### GetDeletedAt

`func (o *ListOrdersItem) GetDeletedAt() time.Time`

GetDeletedAt returns the DeletedAt field if non-nil, zero value otherwise.

### GetDeletedAtOk

`func (o *ListOrdersItem) GetDeletedAtOk() (*time.Time, bool)`

GetDeletedAtOk returns a tuple with the DeletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletedAt

`func (o *ListOrdersItem) SetDeletedAt(v time.Time)`

SetDeletedAt sets DeletedAt field to given value.

### HasDeletedAt

`func (o *ListOrdersItem) HasDeletedAt() bool`

HasDeletedAt returns a boolean if a field has been set.

### SetDeletedAtNil

`func (o *ListOrdersItem) SetDeletedAtNil(b bool)`

 SetDeletedAtNil sets the value for DeletedAt to be an explicit nil

### UnsetDeletedAt
`func (o *ListOrdersItem) UnsetDeletedAt()`

UnsetDeletedAt ensures that no value is present for DeletedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


