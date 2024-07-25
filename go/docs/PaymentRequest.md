# PaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**IsFinal** | Pointer to **bool** |  | [optional] [default to false]
**MoveTaskOrderID** | Pointer to **string** |  | [optional] 
**RejectionReason** | Pointer to **NullableString** |  | [optional] 
**Status** | Pointer to [**PaymentRequestStatus**](PaymentRequestStatus.md) |  | [optional] 
**PaymentRequestNumber** | Pointer to **string** |  | [optional] [readonly] 
**RecalculationOfPaymentRequestID** | Pointer to **NullableString** |  | [optional] [readonly] 
**ProofOfServiceDocs** | Pointer to [**[]ProofOfServiceDoc**](ProofOfServiceDoc.md) |  | [optional] 
**PaymentServiceItems** | Pointer to [**[]PaymentServiceItem**](PaymentServiceItem.md) |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPaymentRequest

`func NewPaymentRequest() *PaymentRequest`

NewPaymentRequest instantiates a new PaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentRequestWithDefaults

`func NewPaymentRequestWithDefaults() *PaymentRequest`

NewPaymentRequestWithDefaults instantiates a new PaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIsFinal

`func (o *PaymentRequest) GetIsFinal() bool`

GetIsFinal returns the IsFinal field if non-nil, zero value otherwise.

### GetIsFinalOk

`func (o *PaymentRequest) GetIsFinalOk() (*bool, bool)`

GetIsFinalOk returns a tuple with the IsFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinal

`func (o *PaymentRequest) SetIsFinal(v bool)`

SetIsFinal sets IsFinal field to given value.

### HasIsFinal

`func (o *PaymentRequest) HasIsFinal() bool`

HasIsFinal returns a boolean if a field has been set.

### GetMoveTaskOrderID

`func (o *PaymentRequest) GetMoveTaskOrderID() string`

GetMoveTaskOrderID returns the MoveTaskOrderID field if non-nil, zero value otherwise.

### GetMoveTaskOrderIDOk

`func (o *PaymentRequest) GetMoveTaskOrderIDOk() (*string, bool)`

GetMoveTaskOrderIDOk returns a tuple with the MoveTaskOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTaskOrderID

`func (o *PaymentRequest) SetMoveTaskOrderID(v string)`

SetMoveTaskOrderID sets MoveTaskOrderID field to given value.

### HasMoveTaskOrderID

`func (o *PaymentRequest) HasMoveTaskOrderID() bool`

HasMoveTaskOrderID returns a boolean if a field has been set.

### GetRejectionReason

`func (o *PaymentRequest) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *PaymentRequest) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *PaymentRequest) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *PaymentRequest) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### SetRejectionReasonNil

`func (o *PaymentRequest) SetRejectionReasonNil(b bool)`

 SetRejectionReasonNil sets the value for RejectionReason to be an explicit nil

### UnsetRejectionReason
`func (o *PaymentRequest) UnsetRejectionReason()`

UnsetRejectionReason ensures that no value is present for RejectionReason, not even an explicit nil
### GetStatus

`func (o *PaymentRequest) GetStatus() PaymentRequestStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentRequest) GetStatusOk() (*PaymentRequestStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentRequest) SetStatus(v PaymentRequestStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPaymentRequestNumber

`func (o *PaymentRequest) GetPaymentRequestNumber() string`

GetPaymentRequestNumber returns the PaymentRequestNumber field if non-nil, zero value otherwise.

### GetPaymentRequestNumberOk

`func (o *PaymentRequest) GetPaymentRequestNumberOk() (*string, bool)`

GetPaymentRequestNumberOk returns a tuple with the PaymentRequestNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestNumber

`func (o *PaymentRequest) SetPaymentRequestNumber(v string)`

SetPaymentRequestNumber sets PaymentRequestNumber field to given value.

### HasPaymentRequestNumber

`func (o *PaymentRequest) HasPaymentRequestNumber() bool`

HasPaymentRequestNumber returns a boolean if a field has been set.

### GetRecalculationOfPaymentRequestID

`func (o *PaymentRequest) GetRecalculationOfPaymentRequestID() string`

GetRecalculationOfPaymentRequestID returns the RecalculationOfPaymentRequestID field if non-nil, zero value otherwise.

### GetRecalculationOfPaymentRequestIDOk

`func (o *PaymentRequest) GetRecalculationOfPaymentRequestIDOk() (*string, bool)`

GetRecalculationOfPaymentRequestIDOk returns a tuple with the RecalculationOfPaymentRequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecalculationOfPaymentRequestID

`func (o *PaymentRequest) SetRecalculationOfPaymentRequestID(v string)`

SetRecalculationOfPaymentRequestID sets RecalculationOfPaymentRequestID field to given value.

### HasRecalculationOfPaymentRequestID

`func (o *PaymentRequest) HasRecalculationOfPaymentRequestID() bool`

HasRecalculationOfPaymentRequestID returns a boolean if a field has been set.

### SetRecalculationOfPaymentRequestIDNil

`func (o *PaymentRequest) SetRecalculationOfPaymentRequestIDNil(b bool)`

 SetRecalculationOfPaymentRequestIDNil sets the value for RecalculationOfPaymentRequestID to be an explicit nil

### UnsetRecalculationOfPaymentRequestID
`func (o *PaymentRequest) UnsetRecalculationOfPaymentRequestID()`

UnsetRecalculationOfPaymentRequestID ensures that no value is present for RecalculationOfPaymentRequestID, not even an explicit nil
### GetProofOfServiceDocs

`func (o *PaymentRequest) GetProofOfServiceDocs() []ProofOfServiceDoc`

GetProofOfServiceDocs returns the ProofOfServiceDocs field if non-nil, zero value otherwise.

### GetProofOfServiceDocsOk

`func (o *PaymentRequest) GetProofOfServiceDocsOk() (*[]ProofOfServiceDoc, bool)`

GetProofOfServiceDocsOk returns a tuple with the ProofOfServiceDocs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProofOfServiceDocs

`func (o *PaymentRequest) SetProofOfServiceDocs(v []ProofOfServiceDoc)`

SetProofOfServiceDocs sets ProofOfServiceDocs field to given value.

### HasProofOfServiceDocs

`func (o *PaymentRequest) HasProofOfServiceDocs() bool`

HasProofOfServiceDocs returns a boolean if a field has been set.

### GetPaymentServiceItems

`func (o *PaymentRequest) GetPaymentServiceItems() []PaymentServiceItem`

GetPaymentServiceItems returns the PaymentServiceItems field if non-nil, zero value otherwise.

### GetPaymentServiceItemsOk

`func (o *PaymentRequest) GetPaymentServiceItemsOk() (*[]PaymentServiceItem, bool)`

GetPaymentServiceItemsOk returns a tuple with the PaymentServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceItems

`func (o *PaymentRequest) SetPaymentServiceItems(v []PaymentServiceItem)`

SetPaymentServiceItems sets PaymentServiceItems field to given value.

### HasPaymentServiceItems

`func (o *PaymentRequest) HasPaymentServiceItems() bool`

HasPaymentServiceItems returns a boolean if a field has been set.

### GetETag

`func (o *PaymentRequest) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *PaymentRequest) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *PaymentRequest) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *PaymentRequest) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


