# MTOServiceItemV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the service item. | [optional] [readonly] 
**MoveTaskOrderID** | **string** | The ID of the move for this service item. | 
**MtoShipmentID** | Pointer to **string** | The ID of the shipment this service is for, if any. Optional. | [optional] 
**ReServiceName** | Pointer to **string** | The full descriptive name of the service. | [optional] [readonly] 
**Status** | Pointer to [**MTOServiceItemStatusV2V2**](MTOServiceItemStatusV2.md) |  | [optional] 
**RejectionReason** | Pointer to **NullableString** | The reason why this service item was rejected by the TOO. | [optional] [readonly] 
**ModelType** | [**MTOServiceItemModelTypeV2V2**](MTOServiceItemModelTypeV2.md) |  | 
**ServiceRequestDocuments** | Pointer to [**[]ServiceRequestDocumentV2V2**](ServiceRequestDocumentV2V2.md) |  | [optional] 
**ETag** | Pointer to **string** | A hash unique to this service item that should be used as the \&quot;If-Match\&quot; header for any updates. | [optional] [readonly] 

## Methods

### NewMTOServiceItemV2

`func NewMTOServiceItemV2(moveTaskOrderID string, modelType MTOServiceItemModelTypeV2V2, ) *MTOServiceItemV2`

NewMTOServiceItemV2 instantiates a new MTOServiceItemV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTOServiceItemV2WithDefaults

`func NewMTOServiceItemV2WithDefaults() *MTOServiceItemV2`

NewMTOServiceItemV2WithDefaults instantiates a new MTOServiceItemV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MTOServiceItemV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MTOServiceItemV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MTOServiceItemV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MTOServiceItemV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMoveTaskOrderID

`func (o *MTOServiceItemV2) GetMoveTaskOrderID() string`

GetMoveTaskOrderID returns the MoveTaskOrderID field if non-nil, zero value otherwise.

### GetMoveTaskOrderIDOk

`func (o *MTOServiceItemV2) GetMoveTaskOrderIDOk() (*string, bool)`

GetMoveTaskOrderIDOk returns a tuple with the MoveTaskOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTaskOrderID

`func (o *MTOServiceItemV2) SetMoveTaskOrderID(v string)`

SetMoveTaskOrderID sets MoveTaskOrderID field to given value.


### GetMtoShipmentID

`func (o *MTOServiceItemV2) GetMtoShipmentID() string`

GetMtoShipmentID returns the MtoShipmentID field if non-nil, zero value otherwise.

### GetMtoShipmentIDOk

`func (o *MTOServiceItemV2) GetMtoShipmentIDOk() (*string, bool)`

GetMtoShipmentIDOk returns a tuple with the MtoShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoShipmentID

`func (o *MTOServiceItemV2) SetMtoShipmentID(v string)`

SetMtoShipmentID sets MtoShipmentID field to given value.

### HasMtoShipmentID

`func (o *MTOServiceItemV2) HasMtoShipmentID() bool`

HasMtoShipmentID returns a boolean if a field has been set.

### GetReServiceName

`func (o *MTOServiceItemV2) GetReServiceName() string`

GetReServiceName returns the ReServiceName field if non-nil, zero value otherwise.

### GetReServiceNameOk

`func (o *MTOServiceItemV2) GetReServiceNameOk() (*string, bool)`

GetReServiceNameOk returns a tuple with the ReServiceName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReServiceName

`func (o *MTOServiceItemV2) SetReServiceName(v string)`

SetReServiceName sets ReServiceName field to given value.

### HasReServiceName

`func (o *MTOServiceItemV2) HasReServiceName() bool`

HasReServiceName returns a boolean if a field has been set.

### GetStatus

`func (o *MTOServiceItemV2) GetStatus() MTOServiceItemStatusV2V2`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MTOServiceItemV2) GetStatusOk() (*MTOServiceItemStatusV2V2, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MTOServiceItemV2) SetStatus(v MTOServiceItemStatusV2V2)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MTOServiceItemV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRejectionReason

`func (o *MTOServiceItemV2) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *MTOServiceItemV2) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *MTOServiceItemV2) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *MTOServiceItemV2) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### SetRejectionReasonNil

`func (o *MTOServiceItemV2) SetRejectionReasonNil(b bool)`

 SetRejectionReasonNil sets the value for RejectionReason to be an explicit nil

### UnsetRejectionReason
`func (o *MTOServiceItemV2) UnsetRejectionReason()`

UnsetRejectionReason ensures that no value is present for RejectionReason, not even an explicit nil
### GetModelType

`func (o *MTOServiceItemV2) GetModelType() MTOServiceItemModelTypeV2V2`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *MTOServiceItemV2) GetModelTypeOk() (*MTOServiceItemModelTypeV2V2, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *MTOServiceItemV2) SetModelType(v MTOServiceItemModelTypeV2V2)`

SetModelType sets ModelType field to given value.


### GetServiceRequestDocuments

`func (o *MTOServiceItemV2) GetServiceRequestDocuments() []ServiceRequestDocumentV2V2`

GetServiceRequestDocuments returns the ServiceRequestDocuments field if non-nil, zero value otherwise.

### GetServiceRequestDocumentsOk

`func (o *MTOServiceItemV2) GetServiceRequestDocumentsOk() (*[]ServiceRequestDocumentV2V2, bool)`

GetServiceRequestDocumentsOk returns a tuple with the ServiceRequestDocuments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceRequestDocuments

`func (o *MTOServiceItemV2) SetServiceRequestDocuments(v []ServiceRequestDocumentV2V2)`

SetServiceRequestDocuments sets ServiceRequestDocuments field to given value.

### HasServiceRequestDocuments

`func (o *MTOServiceItemV2) HasServiceRequestDocuments() bool`

HasServiceRequestDocuments returns a boolean if a field has been set.

### GetETag

`func (o *MTOServiceItemV2) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MTOServiceItemV2) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MTOServiceItemV2) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MTOServiceItemV2) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


