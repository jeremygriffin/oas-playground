# MoveTaskOrderV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**MoveCode** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**OrderID** | Pointer to **string** |  | [optional] 
**Order** | Pointer to [**OrderV2V2**](OrderV2.md) |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**AvailableToPrimeAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**PrimeCounselingCompletedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**PaymentRequests** | [**[]PaymentRequestV2V2**](PaymentRequestV2V2.md) |  | 
**MtoServiceItems** | [**[]MTOServiceItemV2V2**](MTOServiceItemV2V2.md) |  | 
**MtoShipments** | [**[]MTOShipmentWithoutServiceItemsV2V2**](MTOShipmentWithoutServiceItemsV2V2.md) | A list of shipments without their associated service items. | 
**PpmType** | Pointer to **string** |  | [optional] 
**PpmEstimatedWeight** | Pointer to **int32** |  | [optional] 
**ExcessWeightQualifiedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ExcessWeightAcknowledgedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ExcessWeightUploadId** | Pointer to **NullableString** |  | [optional] [readonly] 
**ContractNumber** | Pointer to **string** |  | [optional] [readonly] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewMoveTaskOrderV2

`func NewMoveTaskOrderV2(paymentRequests []PaymentRequestV2V2, mtoServiceItems []MTOServiceItemV2V2, mtoShipments []MTOShipmentWithoutServiceItemsV2V2, ) *MoveTaskOrderV2`

NewMoveTaskOrderV2 instantiates a new MoveTaskOrderV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveTaskOrderV2WithDefaults

`func NewMoveTaskOrderV2WithDefaults() *MoveTaskOrderV2`

NewMoveTaskOrderV2WithDefaults instantiates a new MoveTaskOrderV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MoveTaskOrderV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoveTaskOrderV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoveTaskOrderV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MoveTaskOrderV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMoveCode

`func (o *MoveTaskOrderV2) GetMoveCode() string`

GetMoveCode returns the MoveCode field if non-nil, zero value otherwise.

### GetMoveCodeOk

`func (o *MoveTaskOrderV2) GetMoveCodeOk() (*string, bool)`

GetMoveCodeOk returns a tuple with the MoveCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveCode

`func (o *MoveTaskOrderV2) SetMoveCode(v string)`

SetMoveCode sets MoveCode field to given value.

### HasMoveCode

`func (o *MoveTaskOrderV2) HasMoveCode() bool`

HasMoveCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MoveTaskOrderV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MoveTaskOrderV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MoveTaskOrderV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MoveTaskOrderV2) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOrderID

`func (o *MoveTaskOrderV2) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *MoveTaskOrderV2) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *MoveTaskOrderV2) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.

### HasOrderID

`func (o *MoveTaskOrderV2) HasOrderID() bool`

HasOrderID returns a boolean if a field has been set.

### GetOrder

`func (o *MoveTaskOrderV2) GetOrder() OrderV2V2`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MoveTaskOrderV2) GetOrderOk() (*OrderV2V2, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MoveTaskOrderV2) SetOrder(v OrderV2V2)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MoveTaskOrderV2) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReferenceId

`func (o *MoveTaskOrderV2) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *MoveTaskOrderV2) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *MoveTaskOrderV2) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *MoveTaskOrderV2) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetAvailableToPrimeAt

`func (o *MoveTaskOrderV2) GetAvailableToPrimeAt() time.Time`

GetAvailableToPrimeAt returns the AvailableToPrimeAt field if non-nil, zero value otherwise.

### GetAvailableToPrimeAtOk

`func (o *MoveTaskOrderV2) GetAvailableToPrimeAtOk() (*time.Time, bool)`

GetAvailableToPrimeAtOk returns a tuple with the AvailableToPrimeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableToPrimeAt

`func (o *MoveTaskOrderV2) SetAvailableToPrimeAt(v time.Time)`

SetAvailableToPrimeAt sets AvailableToPrimeAt field to given value.

### HasAvailableToPrimeAt

`func (o *MoveTaskOrderV2) HasAvailableToPrimeAt() bool`

HasAvailableToPrimeAt returns a boolean if a field has been set.

### SetAvailableToPrimeAtNil

`func (o *MoveTaskOrderV2) SetAvailableToPrimeAtNil(b bool)`

 SetAvailableToPrimeAtNil sets the value for AvailableToPrimeAt to be an explicit nil

### UnsetAvailableToPrimeAt
`func (o *MoveTaskOrderV2) UnsetAvailableToPrimeAt()`

UnsetAvailableToPrimeAt ensures that no value is present for AvailableToPrimeAt, not even an explicit nil
### GetUpdatedAt

`func (o *MoveTaskOrderV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MoveTaskOrderV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MoveTaskOrderV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MoveTaskOrderV2) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPrimeCounselingCompletedAt

`func (o *MoveTaskOrderV2) GetPrimeCounselingCompletedAt() time.Time`

GetPrimeCounselingCompletedAt returns the PrimeCounselingCompletedAt field if non-nil, zero value otherwise.

### GetPrimeCounselingCompletedAtOk

`func (o *MoveTaskOrderV2) GetPrimeCounselingCompletedAtOk() (*time.Time, bool)`

GetPrimeCounselingCompletedAtOk returns a tuple with the PrimeCounselingCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeCounselingCompletedAt

`func (o *MoveTaskOrderV2) SetPrimeCounselingCompletedAt(v time.Time)`

SetPrimeCounselingCompletedAt sets PrimeCounselingCompletedAt field to given value.

### HasPrimeCounselingCompletedAt

`func (o *MoveTaskOrderV2) HasPrimeCounselingCompletedAt() bool`

HasPrimeCounselingCompletedAt returns a boolean if a field has been set.

### SetPrimeCounselingCompletedAtNil

`func (o *MoveTaskOrderV2) SetPrimeCounselingCompletedAtNil(b bool)`

 SetPrimeCounselingCompletedAtNil sets the value for PrimeCounselingCompletedAt to be an explicit nil

### UnsetPrimeCounselingCompletedAt
`func (o *MoveTaskOrderV2) UnsetPrimeCounselingCompletedAt()`

UnsetPrimeCounselingCompletedAt ensures that no value is present for PrimeCounselingCompletedAt, not even an explicit nil
### GetPaymentRequests

`func (o *MoveTaskOrderV2) GetPaymentRequests() []PaymentRequestV2V2`

GetPaymentRequests returns the PaymentRequests field if non-nil, zero value otherwise.

### GetPaymentRequestsOk

`func (o *MoveTaskOrderV2) GetPaymentRequestsOk() (*[]PaymentRequestV2V2, bool)`

GetPaymentRequestsOk returns a tuple with the PaymentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequests

`func (o *MoveTaskOrderV2) SetPaymentRequests(v []PaymentRequestV2V2)`

SetPaymentRequests sets PaymentRequests field to given value.


### GetMtoServiceItems

`func (o *MoveTaskOrderV2) GetMtoServiceItems() []MTOServiceItemV2V2`

GetMtoServiceItems returns the MtoServiceItems field if non-nil, zero value otherwise.

### GetMtoServiceItemsOk

`func (o *MoveTaskOrderV2) GetMtoServiceItemsOk() (*[]MTOServiceItemV2V2, bool)`

GetMtoServiceItemsOk returns a tuple with the MtoServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItems

`func (o *MoveTaskOrderV2) SetMtoServiceItems(v []MTOServiceItemV2V2)`

SetMtoServiceItems sets MtoServiceItems field to given value.


### GetMtoShipments

`func (o *MoveTaskOrderV2) GetMtoShipments() []MTOShipmentWithoutServiceItemsV2V2`

GetMtoShipments returns the MtoShipments field if non-nil, zero value otherwise.

### GetMtoShipmentsOk

`func (o *MoveTaskOrderV2) GetMtoShipmentsOk() (*[]MTOShipmentWithoutServiceItemsV2V2, bool)`

GetMtoShipmentsOk returns a tuple with the MtoShipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoShipments

`func (o *MoveTaskOrderV2) SetMtoShipments(v []MTOShipmentWithoutServiceItemsV2V2)`

SetMtoShipments sets MtoShipments field to given value.


### GetPpmType

`func (o *MoveTaskOrderV2) GetPpmType() string`

GetPpmType returns the PpmType field if non-nil, zero value otherwise.

### GetPpmTypeOk

`func (o *MoveTaskOrderV2) GetPpmTypeOk() (*string, bool)`

GetPpmTypeOk returns a tuple with the PpmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpmType

`func (o *MoveTaskOrderV2) SetPpmType(v string)`

SetPpmType sets PpmType field to given value.

### HasPpmType

`func (o *MoveTaskOrderV2) HasPpmType() bool`

HasPpmType returns a boolean if a field has been set.

### GetPpmEstimatedWeight

`func (o *MoveTaskOrderV2) GetPpmEstimatedWeight() int32`

GetPpmEstimatedWeight returns the PpmEstimatedWeight field if non-nil, zero value otherwise.

### GetPpmEstimatedWeightOk

`func (o *MoveTaskOrderV2) GetPpmEstimatedWeightOk() (*int32, bool)`

GetPpmEstimatedWeightOk returns a tuple with the PpmEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpmEstimatedWeight

`func (o *MoveTaskOrderV2) SetPpmEstimatedWeight(v int32)`

SetPpmEstimatedWeight sets PpmEstimatedWeight field to given value.

### HasPpmEstimatedWeight

`func (o *MoveTaskOrderV2) HasPpmEstimatedWeight() bool`

HasPpmEstimatedWeight returns a boolean if a field has been set.

### GetExcessWeightQualifiedAt

`func (o *MoveTaskOrderV2) GetExcessWeightQualifiedAt() time.Time`

GetExcessWeightQualifiedAt returns the ExcessWeightQualifiedAt field if non-nil, zero value otherwise.

### GetExcessWeightQualifiedAtOk

`func (o *MoveTaskOrderV2) GetExcessWeightQualifiedAtOk() (*time.Time, bool)`

GetExcessWeightQualifiedAtOk returns a tuple with the ExcessWeightQualifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcessWeightQualifiedAt

`func (o *MoveTaskOrderV2) SetExcessWeightQualifiedAt(v time.Time)`

SetExcessWeightQualifiedAt sets ExcessWeightQualifiedAt field to given value.

### HasExcessWeightQualifiedAt

`func (o *MoveTaskOrderV2) HasExcessWeightQualifiedAt() bool`

HasExcessWeightQualifiedAt returns a boolean if a field has been set.

### SetExcessWeightQualifiedAtNil

`func (o *MoveTaskOrderV2) SetExcessWeightQualifiedAtNil(b bool)`

 SetExcessWeightQualifiedAtNil sets the value for ExcessWeightQualifiedAt to be an explicit nil

### UnsetExcessWeightQualifiedAt
`func (o *MoveTaskOrderV2) UnsetExcessWeightQualifiedAt()`

UnsetExcessWeightQualifiedAt ensures that no value is present for ExcessWeightQualifiedAt, not even an explicit nil
### GetExcessWeightAcknowledgedAt

`func (o *MoveTaskOrderV2) GetExcessWeightAcknowledgedAt() time.Time`

GetExcessWeightAcknowledgedAt returns the ExcessWeightAcknowledgedAt field if non-nil, zero value otherwise.

### GetExcessWeightAcknowledgedAtOk

`func (o *MoveTaskOrderV2) GetExcessWeightAcknowledgedAtOk() (*time.Time, bool)`

GetExcessWeightAcknowledgedAtOk returns a tuple with the ExcessWeightAcknowledgedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcessWeightAcknowledgedAt

`func (o *MoveTaskOrderV2) SetExcessWeightAcknowledgedAt(v time.Time)`

SetExcessWeightAcknowledgedAt sets ExcessWeightAcknowledgedAt field to given value.

### HasExcessWeightAcknowledgedAt

`func (o *MoveTaskOrderV2) HasExcessWeightAcknowledgedAt() bool`

HasExcessWeightAcknowledgedAt returns a boolean if a field has been set.

### SetExcessWeightAcknowledgedAtNil

`func (o *MoveTaskOrderV2) SetExcessWeightAcknowledgedAtNil(b bool)`

 SetExcessWeightAcknowledgedAtNil sets the value for ExcessWeightAcknowledgedAt to be an explicit nil

### UnsetExcessWeightAcknowledgedAt
`func (o *MoveTaskOrderV2) UnsetExcessWeightAcknowledgedAt()`

UnsetExcessWeightAcknowledgedAt ensures that no value is present for ExcessWeightAcknowledgedAt, not even an explicit nil
### GetExcessWeightUploadId

`func (o *MoveTaskOrderV2) GetExcessWeightUploadId() string`

GetExcessWeightUploadId returns the ExcessWeightUploadId field if non-nil, zero value otherwise.

### GetExcessWeightUploadIdOk

`func (o *MoveTaskOrderV2) GetExcessWeightUploadIdOk() (*string, bool)`

GetExcessWeightUploadIdOk returns a tuple with the ExcessWeightUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcessWeightUploadId

`func (o *MoveTaskOrderV2) SetExcessWeightUploadId(v string)`

SetExcessWeightUploadId sets ExcessWeightUploadId field to given value.

### HasExcessWeightUploadId

`func (o *MoveTaskOrderV2) HasExcessWeightUploadId() bool`

HasExcessWeightUploadId returns a boolean if a field has been set.

### SetExcessWeightUploadIdNil

`func (o *MoveTaskOrderV2) SetExcessWeightUploadIdNil(b bool)`

 SetExcessWeightUploadIdNil sets the value for ExcessWeightUploadId to be an explicit nil

### UnsetExcessWeightUploadId
`func (o *MoveTaskOrderV2) UnsetExcessWeightUploadId()`

UnsetExcessWeightUploadId ensures that no value is present for ExcessWeightUploadId, not even an explicit nil
### GetContractNumber

`func (o *MoveTaskOrderV2) GetContractNumber() string`

GetContractNumber returns the ContractNumber field if non-nil, zero value otherwise.

### GetContractNumberOk

`func (o *MoveTaskOrderV2) GetContractNumberOk() (*string, bool)`

GetContractNumberOk returns a tuple with the ContractNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractNumber

`func (o *MoveTaskOrderV2) SetContractNumber(v string)`

SetContractNumber sets ContractNumber field to given value.

### HasContractNumber

`func (o *MoveTaskOrderV2) HasContractNumber() bool`

HasContractNumber returns a boolean if a field has been set.

### GetETag

`func (o *MoveTaskOrderV2) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MoveTaskOrderV2) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MoveTaskOrderV2) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MoveTaskOrderV2) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


