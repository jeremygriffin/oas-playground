# PaymentServiceItemV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**PaymentRequestID** | Pointer to **string** |  | [optional] 
**MtoServiceItemID** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**PaymentServiceItemStatusV2V2**](PaymentServiceItemStatusV2.md) |  | [optional] 
**PriceCents** | Pointer to **NullableInt32** |  | [optional] 
**RejectionReason** | Pointer to **NullableString** |  | [optional] 
**ReferenceID** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**PaymentServiceItemParams** | Pointer to [**[]PaymentServiceItemParamV2V2**](PaymentServiceItemParamV2V2.md) |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPaymentServiceItemV2

`func NewPaymentServiceItemV2() *PaymentServiceItemV2`

NewPaymentServiceItemV2 instantiates a new PaymentServiceItemV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceItemV2WithDefaults

`func NewPaymentServiceItemV2WithDefaults() *PaymentServiceItemV2`

NewPaymentServiceItemV2WithDefaults instantiates a new PaymentServiceItemV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentServiceItemV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentServiceItemV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentServiceItemV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentServiceItemV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPaymentRequestID

`func (o *PaymentServiceItemV2) GetPaymentRequestID() string`

GetPaymentRequestID returns the PaymentRequestID field if non-nil, zero value otherwise.

### GetPaymentRequestIDOk

`func (o *PaymentServiceItemV2) GetPaymentRequestIDOk() (*string, bool)`

GetPaymentRequestIDOk returns a tuple with the PaymentRequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestID

`func (o *PaymentServiceItemV2) SetPaymentRequestID(v string)`

SetPaymentRequestID sets PaymentRequestID field to given value.

### HasPaymentRequestID

`func (o *PaymentServiceItemV2) HasPaymentRequestID() bool`

HasPaymentRequestID returns a boolean if a field has been set.

### GetMtoServiceItemID

`func (o *PaymentServiceItemV2) GetMtoServiceItemID() string`

GetMtoServiceItemID returns the MtoServiceItemID field if non-nil, zero value otherwise.

### GetMtoServiceItemIDOk

`func (o *PaymentServiceItemV2) GetMtoServiceItemIDOk() (*string, bool)`

GetMtoServiceItemIDOk returns a tuple with the MtoServiceItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItemID

`func (o *PaymentServiceItemV2) SetMtoServiceItemID(v string)`

SetMtoServiceItemID sets MtoServiceItemID field to given value.

### HasMtoServiceItemID

`func (o *PaymentServiceItemV2) HasMtoServiceItemID() bool`

HasMtoServiceItemID returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentServiceItemV2) GetStatus() PaymentServiceItemStatusV2V2`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentServiceItemV2) GetStatusOk() (*PaymentServiceItemStatusV2V2, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentServiceItemV2) SetStatus(v PaymentServiceItemStatusV2V2)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentServiceItemV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriceCents

`func (o *PaymentServiceItemV2) GetPriceCents() int32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *PaymentServiceItemV2) GetPriceCentsOk() (*int32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *PaymentServiceItemV2) SetPriceCents(v int32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *PaymentServiceItemV2) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### SetPriceCentsNil

`func (o *PaymentServiceItemV2) SetPriceCentsNil(b bool)`

 SetPriceCentsNil sets the value for PriceCents to be an explicit nil

### UnsetPriceCents
`func (o *PaymentServiceItemV2) UnsetPriceCents()`

UnsetPriceCents ensures that no value is present for PriceCents, not even an explicit nil
### GetRejectionReason

`func (o *PaymentServiceItemV2) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *PaymentServiceItemV2) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *PaymentServiceItemV2) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *PaymentServiceItemV2) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### SetRejectionReasonNil

`func (o *PaymentServiceItemV2) SetRejectionReasonNil(b bool)`

 SetRejectionReasonNil sets the value for RejectionReason to be an explicit nil

### UnsetRejectionReason
`func (o *PaymentServiceItemV2) UnsetRejectionReason()`

UnsetRejectionReason ensures that no value is present for RejectionReason, not even an explicit nil
### GetReferenceID

`func (o *PaymentServiceItemV2) GetReferenceID() map[string]interface{}`

GetReferenceID returns the ReferenceID field if non-nil, zero value otherwise.

### GetReferenceIDOk

`func (o *PaymentServiceItemV2) GetReferenceIDOk() (*map[string]interface{}, bool)`

GetReferenceIDOk returns a tuple with the ReferenceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceID

`func (o *PaymentServiceItemV2) SetReferenceID(v map[string]interface{})`

SetReferenceID sets ReferenceID field to given value.

### HasReferenceID

`func (o *PaymentServiceItemV2) HasReferenceID() bool`

HasReferenceID returns a boolean if a field has been set.

### GetPaymentServiceItemParams

`func (o *PaymentServiceItemV2) GetPaymentServiceItemParams() []PaymentServiceItemParamV2V2`

GetPaymentServiceItemParams returns the PaymentServiceItemParams field if non-nil, zero value otherwise.

### GetPaymentServiceItemParamsOk

`func (o *PaymentServiceItemV2) GetPaymentServiceItemParamsOk() (*[]PaymentServiceItemParamV2V2, bool)`

GetPaymentServiceItemParamsOk returns a tuple with the PaymentServiceItemParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceItemParams

`func (o *PaymentServiceItemV2) SetPaymentServiceItemParams(v []PaymentServiceItemParamV2V2)`

SetPaymentServiceItemParams sets PaymentServiceItemParams field to given value.

### HasPaymentServiceItemParams

`func (o *PaymentServiceItemV2) HasPaymentServiceItemParams() bool`

HasPaymentServiceItemParams returns a boolean if a field has been set.

### GetETag

`func (o *PaymentServiceItemV2) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *PaymentServiceItemV2) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *PaymentServiceItemV2) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *PaymentServiceItemV2) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


