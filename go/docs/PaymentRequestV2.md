# PaymentRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IsFinal** | Pointer to **bool** |  | [optional] [default to false]
**MoveTaskOrderID** | Pointer to **string** |  | [optional] 
**RejectionReason** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**PaymentRequestStatusV2V2**](PaymentRequestStatusV2.md) |  | [optional] 
**PaymentRequestNumber** | Pointer to **string** |  | [optional] [readonly] 
**RecalculationOfPaymentRequestID** | Pointer to **NullableString** |  | [optional] [readonly] 
**ProofOfServiceDocs** | Pointer to [**[]ProofOfServiceDocV2V2**](ProofOfServiceDocV2V2.md) |  | [optional] 
**PaymentServiceItems** | Pointer to [**[]PaymentServiceItemV2V2**](PaymentServiceItemV2V2.md) |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPaymentRequestV2

`func NewPaymentRequestV2() *PaymentRequestV2`

NewPaymentRequestV2 instantiates a new PaymentRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestV2WithDefaults

`func NewPaymentRequestV2WithDefaults() *PaymentRequestV2`

NewPaymentRequestV2WithDefaults instantiates a new PaymentRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentRequestV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentRequestV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentRequestV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentRequestV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsFinal

`func (o *PaymentRequestV2) GetIsFinal() bool`

GetIsFinal returns the IsFinal field if non-nil, zero value otherwise.

### GetIsFinalOk

`func (o *PaymentRequestV2) GetIsFinalOk() (*bool, bool)`

GetIsFinalOk returns a tuple with the IsFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinal

`func (o *PaymentRequestV2) SetIsFinal(v bool)`

SetIsFinal sets IsFinal field to given value.

### HasIsFinal

`func (o *PaymentRequestV2) HasIsFinal() bool`

HasIsFinal returns a boolean if a field has been set.

### GetMoveTaskOrderID

`func (o *PaymentRequestV2) GetMoveTaskOrderID() string`

GetMoveTaskOrderID returns the MoveTaskOrderID field if non-nil, zero value otherwise.

### GetMoveTaskOrderIDOk

`func (o *PaymentRequestV2) GetMoveTaskOrderIDOk() (*string, bool)`

GetMoveTaskOrderIDOk returns a tuple with the MoveTaskOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTaskOrderID

`func (o *PaymentRequestV2) SetMoveTaskOrderID(v string)`

SetMoveTaskOrderID sets MoveTaskOrderID field to given value.

### HasMoveTaskOrderID

`func (o *PaymentRequestV2) HasMoveTaskOrderID() bool`

HasMoveTaskOrderID returns a boolean if a field has been set.

### GetRejectionReason

`func (o *PaymentRequestV2) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *PaymentRequestV2) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *PaymentRequestV2) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *PaymentRequestV2) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### SetRejectionReasonNil

`func (o *PaymentRequestV2) SetRejectionReasonNil(b bool)`

 SetRejectionReasonNil sets the value for RejectionReason to be an explicit nil

### UnsetRejectionReason
`func (o *PaymentRequestV2) UnsetRejectionReason()`

UnsetRejectionReason ensures that no value is present for RejectionReason, not even an explicit nil
### GetStatus

`func (o *PaymentRequestV2) GetStatus() PaymentRequestStatusV2V2`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentRequestV2) GetStatusOk() (*PaymentRequestStatusV2V2, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentRequestV2) SetStatus(v PaymentRequestStatusV2V2)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentRequestV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPaymentRequestNumber

`func (o *PaymentRequestV2) GetPaymentRequestNumber() string`

GetPaymentRequestNumber returns the PaymentRequestNumber field if non-nil, zero value otherwise.

### GetPaymentRequestNumberOk

`func (o *PaymentRequestV2) GetPaymentRequestNumberOk() (*string, bool)`

GetPaymentRequestNumberOk returns a tuple with the PaymentRequestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestNumber

`func (o *PaymentRequestV2) SetPaymentRequestNumber(v string)`

SetPaymentRequestNumber sets PaymentRequestNumber field to given value.

### HasPaymentRequestNumber

`func (o *PaymentRequestV2) HasPaymentRequestNumber() bool`

HasPaymentRequestNumber returns a boolean if a field has been set.

### GetRecalculationOfPaymentRequestID

`func (o *PaymentRequestV2) GetRecalculationOfPaymentRequestID() string`

GetRecalculationOfPaymentRequestID returns the RecalculationOfPaymentRequestID field if non-nil, zero value otherwise.

### GetRecalculationOfPaymentRequestIDOk

`func (o *PaymentRequestV2) GetRecalculationOfPaymentRequestIDOk() (*string, bool)`

GetRecalculationOfPaymentRequestIDOk returns a tuple with the RecalculationOfPaymentRequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecalculationOfPaymentRequestID

`func (o *PaymentRequestV2) SetRecalculationOfPaymentRequestID(v string)`

SetRecalculationOfPaymentRequestID sets RecalculationOfPaymentRequestID field to given value.

### HasRecalculationOfPaymentRequestID

`func (o *PaymentRequestV2) HasRecalculationOfPaymentRequestID() bool`

HasRecalculationOfPaymentRequestID returns a boolean if a field has been set.

### SetRecalculationOfPaymentRequestIDNil

`func (o *PaymentRequestV2) SetRecalculationOfPaymentRequestIDNil(b bool)`

 SetRecalculationOfPaymentRequestIDNil sets the value for RecalculationOfPaymentRequestID to be an explicit nil

### UnsetRecalculationOfPaymentRequestID
`func (o *PaymentRequestV2) UnsetRecalculationOfPaymentRequestID()`

UnsetRecalculationOfPaymentRequestID ensures that no value is present for RecalculationOfPaymentRequestID, not even an explicit nil
### GetProofOfServiceDocs

`func (o *PaymentRequestV2) GetProofOfServiceDocs() []ProofOfServiceDocV2V2`

GetProofOfServiceDocs returns the ProofOfServiceDocs field if non-nil, zero value otherwise.

### GetProofOfServiceDocsOk

`func (o *PaymentRequestV2) GetProofOfServiceDocsOk() (*[]ProofOfServiceDocV2V2, bool)`

GetProofOfServiceDocsOk returns a tuple with the ProofOfServiceDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofOfServiceDocs

`func (o *PaymentRequestV2) SetProofOfServiceDocs(v []ProofOfServiceDocV2V2)`

SetProofOfServiceDocs sets ProofOfServiceDocs field to given value.

### HasProofOfServiceDocs

`func (o *PaymentRequestV2) HasProofOfServiceDocs() bool`

HasProofOfServiceDocs returns a boolean if a field has been set.

### GetPaymentServiceItems

`func (o *PaymentRequestV2) GetPaymentServiceItems() []PaymentServiceItemV2V2`

GetPaymentServiceItems returns the PaymentServiceItems field if non-nil, zero value otherwise.

### GetPaymentServiceItemsOk

`func (o *PaymentRequestV2) GetPaymentServiceItemsOk() (*[]PaymentServiceItemV2V2, bool)`

GetPaymentServiceItemsOk returns a tuple with the PaymentServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceItems

`func (o *PaymentRequestV2) SetPaymentServiceItems(v []PaymentServiceItemV2V2)`

SetPaymentServiceItems sets PaymentServiceItems field to given value.

### HasPaymentServiceItems

`func (o *PaymentRequestV2) HasPaymentServiceItems() bool`

HasPaymentServiceItems returns a boolean if a field has been set.

### GetETag

`func (o *PaymentRequestV2) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *PaymentRequestV2) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *PaymentRequestV2) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *PaymentRequestV2) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


