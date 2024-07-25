# MoveTaskOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**MoveCode** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**OrderID** | Pointer to **string** |  | [optional] 
**Order** | Pointer to [**Order**](Order.md) |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**AvailableToPrimeAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**PrimeCounselingCompletedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**PaymentRequests** | [**[]PaymentRequest**](PaymentRequest.md) |  | 
**MtoServiceItems** | [**[]MTOServiceItem**](MTOServiceItem.md) |  | 
**MtoShipments** | [**[]MTOShipmentWithoutServiceItems**](MTOShipmentWithoutServiceItems.md) | A list of shipments without their associated service items. | 
**PpmType** | Pointer to **string** |  | [optional] 
**PpmEstimatedWeight** | Pointer to **int32** |  | [optional] 
**ExcessWeightQualifiedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ExcessWeightAcknowledgedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ExcessWeightUploadId** | Pointer to **NullableString** |  | [optional] [readonly] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewMoveTaskOrder

`func NewMoveTaskOrder(paymentRequests []PaymentRequest, mtoServiceItems []MTOServiceItem, mtoShipments []MTOShipmentWithoutServiceItems, ) *MoveTaskOrder`

NewMoveTaskOrder instantiates a new MoveTaskOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveTaskOrderWithDefaults

`func NewMoveTaskOrderWithDefaults() *MoveTaskOrder`

NewMoveTaskOrderWithDefaults instantiates a new MoveTaskOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MoveTaskOrder) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoveTaskOrder) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoveTaskOrder) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MoveTaskOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMoveCode

`func (o *MoveTaskOrder) GetMoveCode() string`

GetMoveCode returns the MoveCode field if non-nil, zero value otherwise.

### GetMoveCodeOk

`func (o *MoveTaskOrder) GetMoveCodeOk() (*string, bool)`

GetMoveCodeOk returns a tuple with the MoveCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveCode

`func (o *MoveTaskOrder) SetMoveCode(v string)`

SetMoveCode sets MoveCode field to given value.

### HasMoveCode

`func (o *MoveTaskOrder) HasMoveCode() bool`

HasMoveCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MoveTaskOrder) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MoveTaskOrder) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MoveTaskOrder) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MoveTaskOrder) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOrderID

`func (o *MoveTaskOrder) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *MoveTaskOrder) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *MoveTaskOrder) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.

### HasOrderID

`func (o *MoveTaskOrder) HasOrderID() bool`

HasOrderID returns a boolean if a field has been set.

### GetOrder

`func (o *MoveTaskOrder) GetOrder() Order`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MoveTaskOrder) GetOrderOk() (*Order, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MoveTaskOrder) SetOrder(v Order)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MoveTaskOrder) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReferenceId

`func (o *MoveTaskOrder) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *MoveTaskOrder) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *MoveTaskOrder) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *MoveTaskOrder) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetAvailableToPrimeAt

`func (o *MoveTaskOrder) GetAvailableToPrimeAt() time.Time`

GetAvailableToPrimeAt returns the AvailableToPrimeAt field if non-nil, zero value otherwise.

### GetAvailableToPrimeAtOk

`func (o *MoveTaskOrder) GetAvailableToPrimeAtOk() (*time.Time, bool)`

GetAvailableToPrimeAtOk returns a tuple with the AvailableToPrimeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableToPrimeAt

`func (o *MoveTaskOrder) SetAvailableToPrimeAt(v time.Time)`

SetAvailableToPrimeAt sets AvailableToPrimeAt field to given value.

### HasAvailableToPrimeAt

`func (o *MoveTaskOrder) HasAvailableToPrimeAt() bool`

HasAvailableToPrimeAt returns a boolean if a field has been set.

### SetAvailableToPrimeAtNil

`func (o *MoveTaskOrder) SetAvailableToPrimeAtNil(b bool)`

 SetAvailableToPrimeAtNil sets the value for AvailableToPrimeAt to be an explicit nil

### UnsetAvailableToPrimeAt
`func (o *MoveTaskOrder) UnsetAvailableToPrimeAt()`

UnsetAvailableToPrimeAt ensures that no value is present for AvailableToPrimeAt, not even an explicit nil
### GetUpdatedAt

`func (o *MoveTaskOrder) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MoveTaskOrder) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MoveTaskOrder) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MoveTaskOrder) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPrimeCounselingCompletedAt

`func (o *MoveTaskOrder) GetPrimeCounselingCompletedAt() time.Time`

GetPrimeCounselingCompletedAt returns the PrimeCounselingCompletedAt field if non-nil, zero value otherwise.

### GetPrimeCounselingCompletedAtOk

`func (o *MoveTaskOrder) GetPrimeCounselingCompletedAtOk() (*time.Time, bool)`

GetPrimeCounselingCompletedAtOk returns a tuple with the PrimeCounselingCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeCounselingCompletedAt

`func (o *MoveTaskOrder) SetPrimeCounselingCompletedAt(v time.Time)`

SetPrimeCounselingCompletedAt sets PrimeCounselingCompletedAt field to given value.

### HasPrimeCounselingCompletedAt

`func (o *MoveTaskOrder) HasPrimeCounselingCompletedAt() bool`

HasPrimeCounselingCompletedAt returns a boolean if a field has been set.

### SetPrimeCounselingCompletedAtNil

`func (o *MoveTaskOrder) SetPrimeCounselingCompletedAtNil(b bool)`

 SetPrimeCounselingCompletedAtNil sets the value for PrimeCounselingCompletedAt to be an explicit nil

### UnsetPrimeCounselingCompletedAt
`func (o *MoveTaskOrder) UnsetPrimeCounselingCompletedAt()`

UnsetPrimeCounselingCompletedAt ensures that no value is present for PrimeCounselingCompletedAt, not even an explicit nil
### GetPaymentRequests

`func (o *MoveTaskOrder) GetPaymentRequests() []PaymentRequest`

GetPaymentRequests returns the PaymentRequests field if non-nil, zero value otherwise.

### GetPaymentRequestsOk

`func (o *MoveTaskOrder) GetPaymentRequestsOk() (*[]PaymentRequest, bool)`

GetPaymentRequestsOk returns a tuple with the PaymentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequests

`func (o *MoveTaskOrder) SetPaymentRequests(v []PaymentRequest)`

SetPaymentRequests sets PaymentRequests field to given value.


### GetMtoServiceItems

`func (o *MoveTaskOrder) GetMtoServiceItems() []MTOServiceItem`

GetMtoServiceItems returns the MtoServiceItems field if non-nil, zero value otherwise.

### GetMtoServiceItemsOk

`func (o *MoveTaskOrder) GetMtoServiceItemsOk() (*[]MTOServiceItem, bool)`

GetMtoServiceItemsOk returns a tuple with the MtoServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItems

`func (o *MoveTaskOrder) SetMtoServiceItems(v []MTOServiceItem)`

SetMtoServiceItems sets MtoServiceItems field to given value.


### GetMtoShipments

`func (o *MoveTaskOrder) GetMtoShipments() []MTOShipmentWithoutServiceItems`

GetMtoShipments returns the MtoShipments field if non-nil, zero value otherwise.

### GetMtoShipmentsOk

`func (o *MoveTaskOrder) GetMtoShipmentsOk() (*[]MTOShipmentWithoutServiceItems, bool)`

GetMtoShipmentsOk returns a tuple with the MtoShipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoShipments

`func (o *MoveTaskOrder) SetMtoShipments(v []MTOShipmentWithoutServiceItems)`

SetMtoShipments sets MtoShipments field to given value.


### GetPpmType

`func (o *MoveTaskOrder) GetPpmType() string`

GetPpmType returns the PpmType field if non-nil, zero value otherwise.

### GetPpmTypeOk

`func (o *MoveTaskOrder) GetPpmTypeOk() (*string, bool)`

GetPpmTypeOk returns a tuple with the PpmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpmType

`func (o *MoveTaskOrder) SetPpmType(v string)`

SetPpmType sets PpmType field to given value.

### HasPpmType

`func (o *MoveTaskOrder) HasPpmType() bool`

HasPpmType returns a boolean if a field has been set.

### GetPpmEstimatedWeight

`func (o *MoveTaskOrder) GetPpmEstimatedWeight() int32`

GetPpmEstimatedWeight returns the PpmEstimatedWeight field if non-nil, zero value otherwise.

### GetPpmEstimatedWeightOk

`func (o *MoveTaskOrder) GetPpmEstimatedWeightOk() (*int32, bool)`

GetPpmEstimatedWeightOk returns a tuple with the PpmEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpmEstimatedWeight

`func (o *MoveTaskOrder) SetPpmEstimatedWeight(v int32)`

SetPpmEstimatedWeight sets PpmEstimatedWeight field to given value.

### HasPpmEstimatedWeight

`func (o *MoveTaskOrder) HasPpmEstimatedWeight() bool`

HasPpmEstimatedWeight returns a boolean if a field has been set.

### GetExcessWeightQualifiedAt

`func (o *MoveTaskOrder) GetExcessWeightQualifiedAt() time.Time`

GetExcessWeightQualifiedAt returns the ExcessWeightQualifiedAt field if non-nil, zero value otherwise.

### GetExcessWeightQualifiedAtOk

`func (o *MoveTaskOrder) GetExcessWeightQualifiedAtOk() (*time.Time, bool)`

GetExcessWeightQualifiedAtOk returns a tuple with the ExcessWeightQualifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcessWeightQualifiedAt

`func (o *MoveTaskOrder) SetExcessWeightQualifiedAt(v time.Time)`

SetExcessWeightQualifiedAt sets ExcessWeightQualifiedAt field to given value.

### HasExcessWeightQualifiedAt

`func (o *MoveTaskOrder) HasExcessWeightQualifiedAt() bool`

HasExcessWeightQualifiedAt returns a boolean if a field has been set.

### SetExcessWeightQualifiedAtNil

`func (o *MoveTaskOrder) SetExcessWeightQualifiedAtNil(b bool)`

 SetExcessWeightQualifiedAtNil sets the value for ExcessWeightQualifiedAt to be an explicit nil

### UnsetExcessWeightQualifiedAt
`func (o *MoveTaskOrder) UnsetExcessWeightQualifiedAt()`

UnsetExcessWeightQualifiedAt ensures that no value is present for ExcessWeightQualifiedAt, not even an explicit nil
### GetExcessWeightAcknowledgedAt

`func (o *MoveTaskOrder) GetExcessWeightAcknowledgedAt() time.Time`

GetExcessWeightAcknowledgedAt returns the ExcessWeightAcknowledgedAt field if non-nil, zero value otherwise.

### GetExcessWeightAcknowledgedAtOk

`func (o *MoveTaskOrder) GetExcessWeightAcknowledgedAtOk() (*time.Time, bool)`

GetExcessWeightAcknowledgedAtOk returns a tuple with the ExcessWeightAcknowledgedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcessWeightAcknowledgedAt

`func (o *MoveTaskOrder) SetExcessWeightAcknowledgedAt(v time.Time)`

SetExcessWeightAcknowledgedAt sets ExcessWeightAcknowledgedAt field to given value.

### HasExcessWeightAcknowledgedAt

`func (o *MoveTaskOrder) HasExcessWeightAcknowledgedAt() bool`

HasExcessWeightAcknowledgedAt returns a boolean if a field has been set.

### SetExcessWeightAcknowledgedAtNil

`func (o *MoveTaskOrder) SetExcessWeightAcknowledgedAtNil(b bool)`

 SetExcessWeightAcknowledgedAtNil sets the value for ExcessWeightAcknowledgedAt to be an explicit nil

### UnsetExcessWeightAcknowledgedAt
`func (o *MoveTaskOrder) UnsetExcessWeightAcknowledgedAt()`

UnsetExcessWeightAcknowledgedAt ensures that no value is present for ExcessWeightAcknowledgedAt, not even an explicit nil
### GetExcessWeightUploadId

`func (o *MoveTaskOrder) GetExcessWeightUploadId() string`

GetExcessWeightUploadId returns the ExcessWeightUploadId field if non-nil, zero value otherwise.

### GetExcessWeightUploadIdOk

`func (o *MoveTaskOrder) GetExcessWeightUploadIdOk() (*string, bool)`

GetExcessWeightUploadIdOk returns a tuple with the ExcessWeightUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcessWeightUploadId

`func (o *MoveTaskOrder) SetExcessWeightUploadId(v string)`

SetExcessWeightUploadId sets ExcessWeightUploadId field to given value.

### HasExcessWeightUploadId

`func (o *MoveTaskOrder) HasExcessWeightUploadId() bool`

HasExcessWeightUploadId returns a boolean if a field has been set.

### SetExcessWeightUploadIdNil

`func (o *MoveTaskOrder) SetExcessWeightUploadIdNil(b bool)`

 SetExcessWeightUploadIdNil sets the value for ExcessWeightUploadId to be an explicit nil

### UnsetExcessWeightUploadId
`func (o *MoveTaskOrder) UnsetExcessWeightUploadId()`

UnsetExcessWeightUploadId ensures that no value is present for ExcessWeightUploadId, not even an explicit nil
### GetETag

`func (o *MoveTaskOrder) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MoveTaskOrder) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MoveTaskOrder) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MoveTaskOrder) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


