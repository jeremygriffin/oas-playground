# MTOServiceItemV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the service item. | [optional] [readonly] 
**MoveTaskOrderID** | **string** | The ID of the move for this service item. | 
**MtoShipmentID** | Pointer to **string** | The ID of the shipment this service is for, if any. Optional. | [optional] 
**ReServiceName** | Pointer to **string** | The full descriptive name of the service. | [optional] [readonly] 
**Status** | Pointer to [**MTOServiceItemStatusV3V3**](MTOServiceItemStatusV3.md) |  | [optional] 
**RejectionReason** | Pointer to **NullableString** | The reason why this service item was rejected by the TOO. | [optional] [readonly] 
**ModelType** | [**MTOServiceItemModelTypeV3V3**](MTOServiceItemModelTypeV3.md) |  | 
**ServiceRequestDocuments** | Pointer to [**[]ServiceRequestDocumentV3V3**](ServiceRequestDocumentV3V3.md) |  | [optional] 
**ETag** | Pointer to **string** | A hash unique to this service item that should be used as the \&quot;If-Match\&quot; header for any updates. | [optional] [readonly] 

## Methods

### NewMTOServiceItemV3

`func NewMTOServiceItemV3(moveTaskOrderID string, modelType MTOServiceItemModelTypeV3V3, ) *MTOServiceItemV3`

NewMTOServiceItemV3 instantiates a new MTOServiceItemV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTOServiceItemV3WithDefaults

`func NewMTOServiceItemV3WithDefaults() *MTOServiceItemV3`

NewMTOServiceItemV3WithDefaults instantiates a new MTOServiceItemV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MTOServiceItemV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MTOServiceItemV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MTOServiceItemV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MTOServiceItemV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMoveTaskOrderID

`func (o *MTOServiceItemV3) GetMoveTaskOrderID() string`

GetMoveTaskOrderID returns the MoveTaskOrderID field if non-nil, zero value otherwise.

### GetMoveTaskOrderIDOk

`func (o *MTOServiceItemV3) GetMoveTaskOrderIDOk() (*string, bool)`

GetMoveTaskOrderIDOk returns a tuple with the MoveTaskOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTaskOrderID

`func (o *MTOServiceItemV3) SetMoveTaskOrderID(v string)`

SetMoveTaskOrderID sets MoveTaskOrderID field to given value.


### GetMtoShipmentID

`func (o *MTOServiceItemV3) GetMtoShipmentID() string`

GetMtoShipmentID returns the MtoShipmentID field if non-nil, zero value otherwise.

### GetMtoShipmentIDOk

`func (o *MTOServiceItemV3) GetMtoShipmentIDOk() (*string, bool)`

GetMtoShipmentIDOk returns a tuple with the MtoShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoShipmentID

`func (o *MTOServiceItemV3) SetMtoShipmentID(v string)`

SetMtoShipmentID sets MtoShipmentID field to given value.

### HasMtoShipmentID

`func (o *MTOServiceItemV3) HasMtoShipmentID() bool`

HasMtoShipmentID returns a boolean if a field has been set.

### GetReServiceName

`func (o *MTOServiceItemV3) GetReServiceName() string`

GetReServiceName returns the ReServiceName field if non-nil, zero value otherwise.

### GetReServiceNameOk

`func (o *MTOServiceItemV3) GetReServiceNameOk() (*string, bool)`

GetReServiceNameOk returns a tuple with the ReServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReServiceName

`func (o *MTOServiceItemV3) SetReServiceName(v string)`

SetReServiceName sets ReServiceName field to given value.

### HasReServiceName

`func (o *MTOServiceItemV3) HasReServiceName() bool`

HasReServiceName returns a boolean if a field has been set.

### GetStatus

`func (o *MTOServiceItemV3) GetStatus() MTOServiceItemStatusV3V3`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MTOServiceItemV3) GetStatusOk() (*MTOServiceItemStatusV3V3, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MTOServiceItemV3) SetStatus(v MTOServiceItemStatusV3V3)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MTOServiceItemV3) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRejectionReason

`func (o *MTOServiceItemV3) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *MTOServiceItemV3) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *MTOServiceItemV3) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *MTOServiceItemV3) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### SetRejectionReasonNil

`func (o *MTOServiceItemV3) SetRejectionReasonNil(b bool)`

 SetRejectionReasonNil sets the value for RejectionReason to be an explicit nil

### UnsetRejectionReason
`func (o *MTOServiceItemV3) UnsetRejectionReason()`

UnsetRejectionReason ensures that no value is present for RejectionReason, not even an explicit nil
### GetModelType

`func (o *MTOServiceItemV3) GetModelType() MTOServiceItemModelTypeV3V3`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *MTOServiceItemV3) GetModelTypeOk() (*MTOServiceItemModelTypeV3V3, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *MTOServiceItemV3) SetModelType(v MTOServiceItemModelTypeV3V3)`

SetModelType sets ModelType field to given value.


### GetServiceRequestDocuments

`func (o *MTOServiceItemV3) GetServiceRequestDocuments() []ServiceRequestDocumentV3V3`

GetServiceRequestDocuments returns the ServiceRequestDocuments field if non-nil, zero value otherwise.

### GetServiceRequestDocumentsOk

`func (o *MTOServiceItemV3) GetServiceRequestDocumentsOk() (*[]ServiceRequestDocumentV3V3, bool)`

GetServiceRequestDocumentsOk returns a tuple with the ServiceRequestDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRequestDocuments

`func (o *MTOServiceItemV3) SetServiceRequestDocuments(v []ServiceRequestDocumentV3V3)`

SetServiceRequestDocuments sets ServiceRequestDocuments field to given value.

### HasServiceRequestDocuments

`func (o *MTOServiceItemV3) HasServiceRequestDocuments() bool`

HasServiceRequestDocuments returns a boolean if a field has been set.

### GetETag

`func (o *MTOServiceItemV3) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MTOServiceItemV3) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MTOServiceItemV3) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MTOServiceItemV3) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


