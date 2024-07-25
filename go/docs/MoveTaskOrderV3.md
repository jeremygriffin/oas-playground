# MoveTaskOrderV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**MoveCode** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**OrderID** | Pointer to **string** |  | [optional] 
**Order** | Pointer to [**OrderV3V3**](OrderV3.md) |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**AvailableToPrimeAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**PrimeCounselingCompletedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**PaymentRequests** | [**[]PaymentRequestV3V3**](PaymentRequestV3V3.md) |  | 
**MtoServiceItems** | [**[]MTOServiceItemV3V3**](MTOServiceItemV3V3.md) |  | 
**MtoShipments** | [**[]MTOShipmentWithoutServiceItemsV3V3**](MTOShipmentWithoutServiceItemsV3V3.md) | A list of shipments without their associated service items. | 
**PpmType** | Pointer to **string** |  | [optional] 
**PpmEstimatedWeight** | Pointer to **int32** |  | [optional] 
**ExcessWeightQualifiedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ExcessWeightAcknowledgedAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**ExcessWeightUploadId** | Pointer to **NullableString** |  | [optional] [readonly] 
**ContractNumber** | Pointer to **string** |  | [optional] [readonly] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewMoveTaskOrderV3

`func NewMoveTaskOrderV3(paymentRequests []PaymentRequestV3V3, mtoServiceItems []MTOServiceItemV3V3, mtoShipments []MTOShipmentWithoutServiceItemsV3V3, ) *MoveTaskOrderV3`

NewMoveTaskOrderV3 instantiates a new MoveTaskOrderV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMoveTaskOrderV3WithDefaults

`func NewMoveTaskOrderV3WithDefaults() *MoveTaskOrderV3`

NewMoveTaskOrderV3WithDefaults instantiates a new MoveTaskOrderV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MoveTaskOrderV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MoveTaskOrderV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MoveTaskOrderV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MoveTaskOrderV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMoveCode

`func (o *MoveTaskOrderV3) GetMoveCode() string`

GetMoveCode returns the MoveCode field if non-nil, zero value otherwise.

### GetMoveCodeOk

`func (o *MoveTaskOrderV3) GetMoveCodeOk() (*string, bool)`

GetMoveCodeOk returns a tuple with the MoveCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveCode

`func (o *MoveTaskOrderV3) SetMoveCode(v string)`

SetMoveCode sets MoveCode field to given value.

### HasMoveCode

`func (o *MoveTaskOrderV3) HasMoveCode() bool`

HasMoveCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MoveTaskOrderV3) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MoveTaskOrderV3) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MoveTaskOrderV3) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MoveTaskOrderV3) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOrderID

`func (o *MoveTaskOrderV3) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *MoveTaskOrderV3) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *MoveTaskOrderV3) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.

### HasOrderID

`func (o *MoveTaskOrderV3) HasOrderID() bool`

HasOrderID returns a boolean if a field has been set.

### GetOrder

`func (o *MoveTaskOrderV3) GetOrder() OrderV3V3`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *MoveTaskOrderV3) GetOrderOk() (*OrderV3V3, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *MoveTaskOrderV3) SetOrder(v OrderV3V3)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *MoveTaskOrderV3) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetReferenceId

`func (o *MoveTaskOrderV3) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *MoveTaskOrderV3) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *MoveTaskOrderV3) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *MoveTaskOrderV3) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetAvailableToPrimeAt

`func (o *MoveTaskOrderV3) GetAvailableToPrimeAt() time.Time`

GetAvailableToPrimeAt returns the AvailableToPrimeAt field if non-nil, zero value otherwise.

### GetAvailableToPrimeAtOk

`func (o *MoveTaskOrderV3) GetAvailableToPrimeAtOk() (*time.Time, bool)`

GetAvailableToPrimeAtOk returns a tuple with the AvailableToPrimeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableToPrimeAt

`func (o *MoveTaskOrderV3) SetAvailableToPrimeAt(v time.Time)`

SetAvailableToPrimeAt sets AvailableToPrimeAt field to given value.

### HasAvailableToPrimeAt

`func (o *MoveTaskOrderV3) HasAvailableToPrimeAt() bool`

HasAvailableToPrimeAt returns a boolean if a field has been set.

### SetAvailableToPrimeAtNil

`func (o *MoveTaskOrderV3) SetAvailableToPrimeAtNil(b bool)`

 SetAvailableToPrimeAtNil sets the value for AvailableToPrimeAt to be an explicit nil

### UnsetAvailableToPrimeAt
`func (o *MoveTaskOrderV3) UnsetAvailableToPrimeAt()`

UnsetAvailableToPrimeAt ensures that no value is present for AvailableToPrimeAt, not even an explicit nil
### GetUpdatedAt

`func (o *MoveTaskOrderV3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MoveTaskOrderV3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MoveTaskOrderV3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MoveTaskOrderV3) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPrimeCounselingCompletedAt

`func (o *MoveTaskOrderV3) GetPrimeCounselingCompletedAt() time.Time`

GetPrimeCounselingCompletedAt returns the PrimeCounselingCompletedAt field if non-nil, zero value otherwise.

### GetPrimeCounselingCompletedAtOk

`func (o *MoveTaskOrderV3) GetPrimeCounselingCompletedAtOk() (*time.Time, bool)`

GetPrimeCounselingCompletedAtOk returns a tuple with the PrimeCounselingCompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeCounselingCompletedAt

`func (o *MoveTaskOrderV3) SetPrimeCounselingCompletedAt(v time.Time)`

SetPrimeCounselingCompletedAt sets PrimeCounselingCompletedAt field to given value.

### HasPrimeCounselingCompletedAt

`func (o *MoveTaskOrderV3) HasPrimeCounselingCompletedAt() bool`

HasPrimeCounselingCompletedAt returns a boolean if a field has been set.

### SetPrimeCounselingCompletedAtNil

`func (o *MoveTaskOrderV3) SetPrimeCounselingCompletedAtNil(b bool)`

 SetPrimeCounselingCompletedAtNil sets the value for PrimeCounselingCompletedAt to be an explicit nil

### UnsetPrimeCounselingCompletedAt
`func (o *MoveTaskOrderV3) UnsetPrimeCounselingCompletedAt()`

UnsetPrimeCounselingCompletedAt ensures that no value is present for PrimeCounselingCompletedAt, not even an explicit nil
### GetPaymentRequests

`func (o *MoveTaskOrderV3) GetPaymentRequests() []PaymentRequestV3V3`

GetPaymentRequests returns the PaymentRequests field if non-nil, zero value otherwise.

### GetPaymentRequestsOk

`func (o *MoveTaskOrderV3) GetPaymentRequestsOk() (*[]PaymentRequestV3V3, bool)`

GetPaymentRequestsOk returns a tuple with the PaymentRequests field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequests

`func (o *MoveTaskOrderV3) SetPaymentRequests(v []PaymentRequestV3V3)`

SetPaymentRequests sets PaymentRequests field to given value.


### GetMtoServiceItems

`func (o *MoveTaskOrderV3) GetMtoServiceItems() []MTOServiceItemV3V3`

GetMtoServiceItems returns the MtoServiceItems field if non-nil, zero value otherwise.

### GetMtoServiceItemsOk

`func (o *MoveTaskOrderV3) GetMtoServiceItemsOk() (*[]MTOServiceItemV3V3, bool)`

GetMtoServiceItemsOk returns a tuple with the MtoServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItems

`func (o *MoveTaskOrderV3) SetMtoServiceItems(v []MTOServiceItemV3V3)`

SetMtoServiceItems sets MtoServiceItems field to given value.


### GetMtoShipments

`func (o *MoveTaskOrderV3) GetMtoShipments() []MTOShipmentWithoutServiceItemsV3V3`

GetMtoShipments returns the MtoShipments field if non-nil, zero value otherwise.

### GetMtoShipmentsOk

`func (o *MoveTaskOrderV3) GetMtoShipmentsOk() (*[]MTOShipmentWithoutServiceItemsV3V3, bool)`

GetMtoShipmentsOk returns a tuple with the MtoShipments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoShipments

`func (o *MoveTaskOrderV3) SetMtoShipments(v []MTOShipmentWithoutServiceItemsV3V3)`

SetMtoShipments sets MtoShipments field to given value.


### GetPpmType

`func (o *MoveTaskOrderV3) GetPpmType() string`

GetPpmType returns the PpmType field if non-nil, zero value otherwise.

### GetPpmTypeOk

`func (o *MoveTaskOrderV3) GetPpmTypeOk() (*string, bool)`

GetPpmTypeOk returns a tuple with the PpmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpmType

`func (o *MoveTaskOrderV3) SetPpmType(v string)`

SetPpmType sets PpmType field to given value.

### HasPpmType

`func (o *MoveTaskOrderV3) HasPpmType() bool`

HasPpmType returns a boolean if a field has been set.

### GetPpmEstimatedWeight

`func (o *MoveTaskOrderV3) GetPpmEstimatedWeight() int32`

GetPpmEstimatedWeight returns the PpmEstimatedWeight field if non-nil, zero value otherwise.

### GetPpmEstimatedWeightOk

`func (o *MoveTaskOrderV3) GetPpmEstimatedWeightOk() (*int32, bool)`

GetPpmEstimatedWeightOk returns a tuple with the PpmEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpmEstimatedWeight

`func (o *MoveTaskOrderV3) SetPpmEstimatedWeight(v int32)`

SetPpmEstimatedWeight sets PpmEstimatedWeight field to given value.

### HasPpmEstimatedWeight

`func (o *MoveTaskOrderV3) HasPpmEstimatedWeight() bool`

HasPpmEstimatedWeight returns a boolean if a field has been set.

### GetExcessWeightQualifiedAt

`func (o *MoveTaskOrderV3) GetExcessWeightQualifiedAt() time.Time`

GetExcessWeightQualifiedAt returns the ExcessWeightQualifiedAt field if non-nil, zero value otherwise.

### GetExcessWeightQualifiedAtOk

`func (o *MoveTaskOrderV3) GetExcessWeightQualifiedAtOk() (*time.Time, bool)`

GetExcessWeightQualifiedAtOk returns a tuple with the ExcessWeightQualifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcessWeightQualifiedAt

`func (o *MoveTaskOrderV3) SetExcessWeightQualifiedAt(v time.Time)`

SetExcessWeightQualifiedAt sets ExcessWeightQualifiedAt field to given value.

### HasExcessWeightQualifiedAt

`func (o *MoveTaskOrderV3) HasExcessWeightQualifiedAt() bool`

HasExcessWeightQualifiedAt returns a boolean if a field has been set.

### SetExcessWeightQualifiedAtNil

`func (o *MoveTaskOrderV3) SetExcessWeightQualifiedAtNil(b bool)`

 SetExcessWeightQualifiedAtNil sets the value for ExcessWeightQualifiedAt to be an explicit nil

### UnsetExcessWeightQualifiedAt
`func (o *MoveTaskOrderV3) UnsetExcessWeightQualifiedAt()`

UnsetExcessWeightQualifiedAt ensures that no value is present for ExcessWeightQualifiedAt, not even an explicit nil
### GetExcessWeightAcknowledgedAt

`func (o *MoveTaskOrderV3) GetExcessWeightAcknowledgedAt() time.Time`

GetExcessWeightAcknowledgedAt returns the ExcessWeightAcknowledgedAt field if non-nil, zero value otherwise.

### GetExcessWeightAcknowledgedAtOk

`func (o *MoveTaskOrderV3) GetExcessWeightAcknowledgedAtOk() (*time.Time, bool)`

GetExcessWeightAcknowledgedAtOk returns a tuple with the ExcessWeightAcknowledgedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcessWeightAcknowledgedAt

`func (o *MoveTaskOrderV3) SetExcessWeightAcknowledgedAt(v time.Time)`

SetExcessWeightAcknowledgedAt sets ExcessWeightAcknowledgedAt field to given value.

### HasExcessWeightAcknowledgedAt

`func (o *MoveTaskOrderV3) HasExcessWeightAcknowledgedAt() bool`

HasExcessWeightAcknowledgedAt returns a boolean if a field has been set.

### SetExcessWeightAcknowledgedAtNil

`func (o *MoveTaskOrderV3) SetExcessWeightAcknowledgedAtNil(b bool)`

 SetExcessWeightAcknowledgedAtNil sets the value for ExcessWeightAcknowledgedAt to be an explicit nil

### UnsetExcessWeightAcknowledgedAt
`func (o *MoveTaskOrderV3) UnsetExcessWeightAcknowledgedAt()`

UnsetExcessWeightAcknowledgedAt ensures that no value is present for ExcessWeightAcknowledgedAt, not even an explicit nil
### GetExcessWeightUploadId

`func (o *MoveTaskOrderV3) GetExcessWeightUploadId() string`

GetExcessWeightUploadId returns the ExcessWeightUploadId field if non-nil, zero value otherwise.

### GetExcessWeightUploadIdOk

`func (o *MoveTaskOrderV3) GetExcessWeightUploadIdOk() (*string, bool)`

GetExcessWeightUploadIdOk returns a tuple with the ExcessWeightUploadId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcessWeightUploadId

`func (o *MoveTaskOrderV3) SetExcessWeightUploadId(v string)`

SetExcessWeightUploadId sets ExcessWeightUploadId field to given value.

### HasExcessWeightUploadId

`func (o *MoveTaskOrderV3) HasExcessWeightUploadId() bool`

HasExcessWeightUploadId returns a boolean if a field has been set.

### SetExcessWeightUploadIdNil

`func (o *MoveTaskOrderV3) SetExcessWeightUploadIdNil(b bool)`

 SetExcessWeightUploadIdNil sets the value for ExcessWeightUploadId to be an explicit nil

### UnsetExcessWeightUploadId
`func (o *MoveTaskOrderV3) UnsetExcessWeightUploadId()`

UnsetExcessWeightUploadId ensures that no value is present for ExcessWeightUploadId, not even an explicit nil
### GetContractNumber

`func (o *MoveTaskOrderV3) GetContractNumber() string`

GetContractNumber returns the ContractNumber field if non-nil, zero value otherwise.

### GetContractNumberOk

`func (o *MoveTaskOrderV3) GetContractNumberOk() (*string, bool)`

GetContractNumberOk returns a tuple with the ContractNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractNumber

`func (o *MoveTaskOrderV3) SetContractNumber(v string)`

SetContractNumber sets ContractNumber field to given value.

### HasContractNumber

`func (o *MoveTaskOrderV3) HasContractNumber() bool`

HasContractNumber returns a boolean if a field has been set.

### GetETag

`func (o *MoveTaskOrderV3) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MoveTaskOrderV3) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MoveTaskOrderV3) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MoveTaskOrderV3) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


