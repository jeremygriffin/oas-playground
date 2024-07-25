# PaymentRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IsFinal** | Pointer to **bool** |  | [optional] [default to false]
**MoveTaskOrderID** | Pointer to **string** |  | [optional] 
**RejectionReason** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**PaymentRequestStatusV3V3**](PaymentRequestStatusV3.md) |  | [optional] 
**PaymentRequestNumber** | Pointer to **string** |  | [optional] [readonly] 
**RecalculationOfPaymentRequestID** | Pointer to **NullableString** |  | [optional] [readonly] 
**ProofOfServiceDocs** | Pointer to [**[]ProofOfServiceDocV3V3**](ProofOfServiceDocV3V3.md) |  | [optional] 
**PaymentServiceItems** | Pointer to [**[]PaymentServiceItemV3V3**](PaymentServiceItemV3V3.md) |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPaymentRequestV3

`func NewPaymentRequestV3() *PaymentRequestV3`

NewPaymentRequestV3 instantiates a new PaymentRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestV3WithDefaults

`func NewPaymentRequestV3WithDefaults() *PaymentRequestV3`

NewPaymentRequestV3WithDefaults instantiates a new PaymentRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentRequestV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentRequestV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentRequestV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentRequestV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsFinal

`func (o *PaymentRequestV3) GetIsFinal() bool`

GetIsFinal returns the IsFinal field if non-nil, zero value otherwise.

### GetIsFinalOk

`func (o *PaymentRequestV3) GetIsFinalOk() (*bool, bool)`

GetIsFinalOk returns a tuple with the IsFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinal

`func (o *PaymentRequestV3) SetIsFinal(v bool)`

SetIsFinal sets IsFinal field to given value.

### HasIsFinal

`func (o *PaymentRequestV3) HasIsFinal() bool`

HasIsFinal returns a boolean if a field has been set.

### GetMoveTaskOrderID

`func (o *PaymentRequestV3) GetMoveTaskOrderID() string`

GetMoveTaskOrderID returns the MoveTaskOrderID field if non-nil, zero value otherwise.

### GetMoveTaskOrderIDOk

`func (o *PaymentRequestV3) GetMoveTaskOrderIDOk() (*string, bool)`

GetMoveTaskOrderIDOk returns a tuple with the MoveTaskOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTaskOrderID

`func (o *PaymentRequestV3) SetMoveTaskOrderID(v string)`

SetMoveTaskOrderID sets MoveTaskOrderID field to given value.

### HasMoveTaskOrderID

`func (o *PaymentRequestV3) HasMoveTaskOrderID() bool`

HasMoveTaskOrderID returns a boolean if a field has been set.

### GetRejectionReason

`func (o *PaymentRequestV3) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *PaymentRequestV3) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *PaymentRequestV3) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *PaymentRequestV3) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### SetRejectionReasonNil

`func (o *PaymentRequestV3) SetRejectionReasonNil(b bool)`

 SetRejectionReasonNil sets the value for RejectionReason to be an explicit nil

### UnsetRejectionReason
`func (o *PaymentRequestV3) UnsetRejectionReason()`

UnsetRejectionReason ensures that no value is present for RejectionReason, not even an explicit nil
### GetStatus

`func (o *PaymentRequestV3) GetStatus() PaymentRequestStatusV3V3`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentRequestV3) GetStatusOk() (*PaymentRequestStatusV3V3, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentRequestV3) SetStatus(v PaymentRequestStatusV3V3)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentRequestV3) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPaymentRequestNumber

`func (o *PaymentRequestV3) GetPaymentRequestNumber() string`

GetPaymentRequestNumber returns the PaymentRequestNumber field if non-nil, zero value otherwise.

### GetPaymentRequestNumberOk

`func (o *PaymentRequestV3) GetPaymentRequestNumberOk() (*string, bool)`

GetPaymentRequestNumberOk returns a tuple with the PaymentRequestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestNumber

`func (o *PaymentRequestV3) SetPaymentRequestNumber(v string)`

SetPaymentRequestNumber sets PaymentRequestNumber field to given value.

### HasPaymentRequestNumber

`func (o *PaymentRequestV3) HasPaymentRequestNumber() bool`

HasPaymentRequestNumber returns a boolean if a field has been set.

### GetRecalculationOfPaymentRequestID

`func (o *PaymentRequestV3) GetRecalculationOfPaymentRequestID() string`

GetRecalculationOfPaymentRequestID returns the RecalculationOfPaymentRequestID field if non-nil, zero value otherwise.

### GetRecalculationOfPaymentRequestIDOk

`func (o *PaymentRequestV3) GetRecalculationOfPaymentRequestIDOk() (*string, bool)`

GetRecalculationOfPaymentRequestIDOk returns a tuple with the RecalculationOfPaymentRequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecalculationOfPaymentRequestID

`func (o *PaymentRequestV3) SetRecalculationOfPaymentRequestID(v string)`

SetRecalculationOfPaymentRequestID sets RecalculationOfPaymentRequestID field to given value.

### HasRecalculationOfPaymentRequestID

`func (o *PaymentRequestV3) HasRecalculationOfPaymentRequestID() bool`

HasRecalculationOfPaymentRequestID returns a boolean if a field has been set.

### SetRecalculationOfPaymentRequestIDNil

`func (o *PaymentRequestV3) SetRecalculationOfPaymentRequestIDNil(b bool)`

 SetRecalculationOfPaymentRequestIDNil sets the value for RecalculationOfPaymentRequestID to be an explicit nil

### UnsetRecalculationOfPaymentRequestID
`func (o *PaymentRequestV3) UnsetRecalculationOfPaymentRequestID()`

UnsetRecalculationOfPaymentRequestID ensures that no value is present for RecalculationOfPaymentRequestID, not even an explicit nil
### GetProofOfServiceDocs

`func (o *PaymentRequestV3) GetProofOfServiceDocs() []ProofOfServiceDocV3V3`

GetProofOfServiceDocs returns the ProofOfServiceDocs field if non-nil, zero value otherwise.

### GetProofOfServiceDocsOk

`func (o *PaymentRequestV3) GetProofOfServiceDocsOk() (*[]ProofOfServiceDocV3V3, bool)`

GetProofOfServiceDocsOk returns a tuple with the ProofOfServiceDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofOfServiceDocs

`func (o *PaymentRequestV3) SetProofOfServiceDocs(v []ProofOfServiceDocV3V3)`

SetProofOfServiceDocs sets ProofOfServiceDocs field to given value.

### HasProofOfServiceDocs

`func (o *PaymentRequestV3) HasProofOfServiceDocs() bool`

HasProofOfServiceDocs returns a boolean if a field has been set.

### GetPaymentServiceItems

`func (o *PaymentRequestV3) GetPaymentServiceItems() []PaymentServiceItemV3V3`

GetPaymentServiceItems returns the PaymentServiceItems field if non-nil, zero value otherwise.

### GetPaymentServiceItemsOk

`func (o *PaymentRequestV3) GetPaymentServiceItemsOk() (*[]PaymentServiceItemV3V3, bool)`

GetPaymentServiceItemsOk returns a tuple with the PaymentServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceItems

`func (o *PaymentRequestV3) SetPaymentServiceItems(v []PaymentServiceItemV3V3)`

SetPaymentServiceItems sets PaymentServiceItems field to given value.

### HasPaymentServiceItems

`func (o *PaymentRequestV3) HasPaymentServiceItems() bool`

HasPaymentServiceItems returns a boolean if a field has been set.

### GetETag

`func (o *PaymentRequestV3) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *PaymentRequestV3) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *PaymentRequestV3) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *PaymentRequestV3) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


