# MTOServiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the service item. | [optional] [readonly] 
**MoveTaskOrderID** | **string** | The ID of the move for this service item. | 
**MtoShipmentID** | Pointer to **string** | The ID of the shipment this service is for, if any. Optional. | [optional] 
**ReServiceName** | Pointer to **string** | The full descriptive name of the service. | [optional] [readonly] 
**Status** | Pointer to [**MTOServiceItemStatus**](MTOServiceItemStatus.md) |  | [optional] 
**RejectionReason** | Pointer to **NullableString** | The reason why this service item was rejected by the TOO. | [optional] [readonly] 
**ModelType** | [**MTOServiceItemModelType**](MTOServiceItemModelType.md) |  | 
**ServiceRequestDocuments** | Pointer to [**[]ServiceRequestDocument**](ServiceRequestDocument.md) |  | [optional] 
**ETag** | Pointer to **string** | A hash unique to this service item that should be used as the \&quot;If-Match\&quot; header for any updates. | [optional] [readonly] 

## Methods

### NewMTOServiceItem

`func NewMTOServiceItem(moveTaskOrderID string, modelType MTOServiceItemModelType, ) *MTOServiceItem`

NewMTOServiceItem instantiates a new MTOServiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTOServiceItemWithDefaults

`func NewMTOServiceItemWithDefaults() *MTOServiceItem`

NewMTOServiceItemWithDefaults instantiates a new MTOServiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MTOServiceItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MTOServiceItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MTOServiceItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MTOServiceItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMoveTaskOrderID

`func (o *MTOServiceItem) GetMoveTaskOrderID() string`

GetMoveTaskOrderID returns the MoveTaskOrderID field if non-nil, zero value otherwise.

### GetMoveTaskOrderIDOk

`func (o *MTOServiceItem) GetMoveTaskOrderIDOk() (*string, bool)`

GetMoveTaskOrderIDOk returns a tuple with the MoveTaskOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTaskOrderID

`func (o *MTOServiceItem) SetMoveTaskOrderID(v string)`

SetMoveTaskOrderID sets MoveTaskOrderID field to given value.


### GetMtoShipmentID

`func (o *MTOServiceItem) GetMtoShipmentID() string`

GetMtoShipmentID returns the MtoShipmentID field if non-nil, zero value otherwise.

### GetMtoShipmentIDOk

`func (o *MTOServiceItem) GetMtoShipmentIDOk() (*string, bool)`

GetMtoShipmentIDOk returns a tuple with the MtoShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoShipmentID

`func (o *MTOServiceItem) SetMtoShipmentID(v string)`

SetMtoShipmentID sets MtoShipmentID field to given value.

### HasMtoShipmentID

`func (o *MTOServiceItem) HasMtoShipmentID() bool`

HasMtoShipmentID returns a boolean if a field has been set.

### GetReServiceName

`func (o *MTOServiceItem) GetReServiceName() string`

GetReServiceName returns the ReServiceName field if non-nil, zero value otherwise.

### GetReServiceNameOk

`func (o *MTOServiceItem) GetReServiceNameOk() (*string, bool)`

GetReServiceNameOk returns a tuple with the ReServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReServiceName

`func (o *MTOServiceItem) SetReServiceName(v string)`

SetReServiceName sets ReServiceName field to given value.

### HasReServiceName

`func (o *MTOServiceItem) HasReServiceName() bool`

HasReServiceName returns a boolean if a field has been set.

### GetStatus

`func (o *MTOServiceItem) GetStatus() MTOServiceItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MTOServiceItem) GetStatusOk() (*MTOServiceItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MTOServiceItem) SetStatus(v MTOServiceItemStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MTOServiceItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRejectionReason

`func (o *MTOServiceItem) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *MTOServiceItem) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *MTOServiceItem) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *MTOServiceItem) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### SetRejectionReasonNil

`func (o *MTOServiceItem) SetRejectionReasonNil(b bool)`

 SetRejectionReasonNil sets the value for RejectionReason to be an explicit nil

### UnsetRejectionReason
`func (o *MTOServiceItem) UnsetRejectionReason()`

UnsetRejectionReason ensures that no value is present for RejectionReason, not even an explicit nil
### GetModelType

`func (o *MTOServiceItem) GetModelType() MTOServiceItemModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *MTOServiceItem) GetModelTypeOk() (*MTOServiceItemModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *MTOServiceItem) SetModelType(v MTOServiceItemModelType)`

SetModelType sets ModelType field to given value.


### GetServiceRequestDocuments

`func (o *MTOServiceItem) GetServiceRequestDocuments() []ServiceRequestDocument`

GetServiceRequestDocuments returns the ServiceRequestDocuments field if non-nil, zero value otherwise.

### GetServiceRequestDocumentsOk

`func (o *MTOServiceItem) GetServiceRequestDocumentsOk() (*[]ServiceRequestDocument, bool)`

GetServiceRequestDocumentsOk returns a tuple with the ServiceRequestDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRequestDocuments

`func (o *MTOServiceItem) SetServiceRequestDocuments(v []ServiceRequestDocument)`

SetServiceRequestDocuments sets ServiceRequestDocuments field to given value.

### HasServiceRequestDocuments

`func (o *MTOServiceItem) HasServiceRequestDocuments() bool`

HasServiceRequestDocuments returns a boolean if a field has been set.

### GetETag

`func (o *MTOServiceItem) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MTOServiceItem) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MTOServiceItem) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MTOServiceItem) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


