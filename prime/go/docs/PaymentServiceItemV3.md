# PaymentServiceItemV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**PaymentRequestID** | Pointer to **string** |  | [optional] 
**MtoServiceItemID** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**PaymentServiceItemStatusV3V3**](PaymentServiceItemStatusV3.md) |  | [optional] 
**PriceCents** | Pointer to **NullableInt32** |  | [optional] 
**RejectionReason** | Pointer to **NullableString** |  | [optional] 
**ReferenceID** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**PaymentServiceItemParams** | Pointer to [**[]PaymentServiceItemParamV3V3**](PaymentServiceItemParamV3V3.md) |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPaymentServiceItemV3

`func NewPaymentServiceItemV3() *PaymentServiceItemV3`

NewPaymentServiceItemV3 instantiates a new PaymentServiceItemV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceItemV3WithDefaults

`func NewPaymentServiceItemV3WithDefaults() *PaymentServiceItemV3`

NewPaymentServiceItemV3WithDefaults instantiates a new PaymentServiceItemV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentServiceItemV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentServiceItemV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentServiceItemV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentServiceItemV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPaymentRequestID

`func (o *PaymentServiceItemV3) GetPaymentRequestID() string`

GetPaymentRequestID returns the PaymentRequestID field if non-nil, zero value otherwise.

### GetPaymentRequestIDOk

`func (o *PaymentServiceItemV3) GetPaymentRequestIDOk() (*string, bool)`

GetPaymentRequestIDOk returns a tuple with the PaymentRequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestID

`func (o *PaymentServiceItemV3) SetPaymentRequestID(v string)`

SetPaymentRequestID sets PaymentRequestID field to given value.

### HasPaymentRequestID

`func (o *PaymentServiceItemV3) HasPaymentRequestID() bool`

HasPaymentRequestID returns a boolean if a field has been set.

### GetMtoServiceItemID

`func (o *PaymentServiceItemV3) GetMtoServiceItemID() string`

GetMtoServiceItemID returns the MtoServiceItemID field if non-nil, zero value otherwise.

### GetMtoServiceItemIDOk

`func (o *PaymentServiceItemV3) GetMtoServiceItemIDOk() (*string, bool)`

GetMtoServiceItemIDOk returns a tuple with the MtoServiceItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItemID

`func (o *PaymentServiceItemV3) SetMtoServiceItemID(v string)`

SetMtoServiceItemID sets MtoServiceItemID field to given value.

### HasMtoServiceItemID

`func (o *PaymentServiceItemV3) HasMtoServiceItemID() bool`

HasMtoServiceItemID returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentServiceItemV3) GetStatus() PaymentServiceItemStatusV3V3`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentServiceItemV3) GetStatusOk() (*PaymentServiceItemStatusV3V3, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentServiceItemV3) SetStatus(v PaymentServiceItemStatusV3V3)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentServiceItemV3) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriceCents

`func (o *PaymentServiceItemV3) GetPriceCents() int32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *PaymentServiceItemV3) GetPriceCentsOk() (*int32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *PaymentServiceItemV3) SetPriceCents(v int32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *PaymentServiceItemV3) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### SetPriceCentsNil

`func (o *PaymentServiceItemV3) SetPriceCentsNil(b bool)`

 SetPriceCentsNil sets the value for PriceCents to be an explicit nil

### UnsetPriceCents
`func (o *PaymentServiceItemV3) UnsetPriceCents()`

UnsetPriceCents ensures that no value is present for PriceCents, not even an explicit nil
### GetRejectionReason

`func (o *PaymentServiceItemV3) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *PaymentServiceItemV3) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *PaymentServiceItemV3) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *PaymentServiceItemV3) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### SetRejectionReasonNil

`func (o *PaymentServiceItemV3) SetRejectionReasonNil(b bool)`

 SetRejectionReasonNil sets the value for RejectionReason to be an explicit nil

### UnsetRejectionReason
`func (o *PaymentServiceItemV3) UnsetRejectionReason()`

UnsetRejectionReason ensures that no value is present for RejectionReason, not even an explicit nil
### GetReferenceID

`func (o *PaymentServiceItemV3) GetReferenceID() map[string]interface{}`

GetReferenceID returns the ReferenceID field if non-nil, zero value otherwise.

### GetReferenceIDOk

`func (o *PaymentServiceItemV3) GetReferenceIDOk() (*map[string]interface{}, bool)`

GetReferenceIDOk returns a tuple with the ReferenceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceID

`func (o *PaymentServiceItemV3) SetReferenceID(v map[string]interface{})`

SetReferenceID sets ReferenceID field to given value.

### HasReferenceID

`func (o *PaymentServiceItemV3) HasReferenceID() bool`

HasReferenceID returns a boolean if a field has been set.

### GetPaymentServiceItemParams

`func (o *PaymentServiceItemV3) GetPaymentServiceItemParams() []PaymentServiceItemParamV3V3`

GetPaymentServiceItemParams returns the PaymentServiceItemParams field if non-nil, zero value otherwise.

### GetPaymentServiceItemParamsOk

`func (o *PaymentServiceItemV3) GetPaymentServiceItemParamsOk() (*[]PaymentServiceItemParamV3V3, bool)`

GetPaymentServiceItemParamsOk returns a tuple with the PaymentServiceItemParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceItemParams

`func (o *PaymentServiceItemV3) SetPaymentServiceItemParams(v []PaymentServiceItemParamV3V3)`

SetPaymentServiceItemParams sets PaymentServiceItemParams field to given value.

### HasPaymentServiceItemParams

`func (o *PaymentServiceItemV3) HasPaymentServiceItemParams() bool`

HasPaymentServiceItemParams returns a boolean if a field has been set.

### GetETag

`func (o *PaymentServiceItemV3) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *PaymentServiceItemV3) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *PaymentServiceItemV3) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *PaymentServiceItemV3) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


