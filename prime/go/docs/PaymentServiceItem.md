# PaymentServiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**PaymentRequestID** | Pointer to **string** |  | [optional] 
**MtoServiceItemID** | Pointer to **string** |  | [optional] 
**Status** | Pointer to [**PaymentServiceItemStatus**](PaymentServiceItemStatus.md) |  | [optional] 
**PriceCents** | Pointer to **NullableInt32** |  | [optional] 
**RejectionReason** | Pointer to **NullableString** |  | [optional] 
**ReferenceID** | Pointer to **map[string]interface{}** |  | [optional] [readonly] 
**PaymentServiceItemParams** | Pointer to [**[]PaymentServiceItemParam**](PaymentServiceItemParam.md) |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPaymentServiceItem

`func NewPaymentServiceItem() *PaymentServiceItem`

NewPaymentServiceItem instantiates a new PaymentServiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceItemWithDefaults

`func NewPaymentServiceItemWithDefaults() *PaymentServiceItem`

NewPaymentServiceItemWithDefaults instantiates a new PaymentServiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentServiceItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentServiceItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentServiceItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentServiceItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPaymentRequestID

`func (o *PaymentServiceItem) GetPaymentRequestID() string`

GetPaymentRequestID returns the PaymentRequestID field if non-nil, zero value otherwise.

### GetPaymentRequestIDOk

`func (o *PaymentServiceItem) GetPaymentRequestIDOk() (*string, bool)`

GetPaymentRequestIDOk returns a tuple with the PaymentRequestID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentRequestID

`func (o *PaymentServiceItem) SetPaymentRequestID(v string)`

SetPaymentRequestID sets PaymentRequestID field to given value.

### HasPaymentRequestID

`func (o *PaymentServiceItem) HasPaymentRequestID() bool`

HasPaymentRequestID returns a boolean if a field has been set.

### GetMtoServiceItemID

`func (o *PaymentServiceItem) GetMtoServiceItemID() string`

GetMtoServiceItemID returns the MtoServiceItemID field if non-nil, zero value otherwise.

### GetMtoServiceItemIDOk

`func (o *PaymentServiceItem) GetMtoServiceItemIDOk() (*string, bool)`

GetMtoServiceItemIDOk returns a tuple with the MtoServiceItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItemID

`func (o *PaymentServiceItem) SetMtoServiceItemID(v string)`

SetMtoServiceItemID sets MtoServiceItemID field to given value.

### HasMtoServiceItemID

`func (o *PaymentServiceItem) HasMtoServiceItemID() bool`

HasMtoServiceItemID returns a boolean if a field has been set.

### GetStatus

`func (o *PaymentServiceItem) GetStatus() PaymentServiceItemStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PaymentServiceItem) GetStatusOk() (*PaymentServiceItemStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PaymentServiceItem) SetStatus(v PaymentServiceItemStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *PaymentServiceItem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPriceCents

`func (o *PaymentServiceItem) GetPriceCents() int32`

GetPriceCents returns the PriceCents field if non-nil, zero value otherwise.

### GetPriceCentsOk

`func (o *PaymentServiceItem) GetPriceCentsOk() (*int32, bool)`

GetPriceCentsOk returns a tuple with the PriceCents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceCents

`func (o *PaymentServiceItem) SetPriceCents(v int32)`

SetPriceCents sets PriceCents field to given value.

### HasPriceCents

`func (o *PaymentServiceItem) HasPriceCents() bool`

HasPriceCents returns a boolean if a field has been set.

### SetPriceCentsNil

`func (o *PaymentServiceItem) SetPriceCentsNil(b bool)`

 SetPriceCentsNil sets the value for PriceCents to be an explicit nil

### UnsetPriceCents
`func (o *PaymentServiceItem) UnsetPriceCents()`

UnsetPriceCents ensures that no value is present for PriceCents, not even an explicit nil
### GetRejectionReason

`func (o *PaymentServiceItem) GetRejectionReason() string`

GetRejectionReason returns the RejectionReason field if non-nil, zero value otherwise.

### GetRejectionReasonOk

`func (o *PaymentServiceItem) GetRejectionReasonOk() (*string, bool)`

GetRejectionReasonOk returns a tuple with the RejectionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRejectionReason

`func (o *PaymentServiceItem) SetRejectionReason(v string)`

SetRejectionReason sets RejectionReason field to given value.

### HasRejectionReason

`func (o *PaymentServiceItem) HasRejectionReason() bool`

HasRejectionReason returns a boolean if a field has been set.

### SetRejectionReasonNil

`func (o *PaymentServiceItem) SetRejectionReasonNil(b bool)`

 SetRejectionReasonNil sets the value for RejectionReason to be an explicit nil

### UnsetRejectionReason
`func (o *PaymentServiceItem) UnsetRejectionReason()`

UnsetRejectionReason ensures that no value is present for RejectionReason, not even an explicit nil
### GetReferenceID

`func (o *PaymentServiceItem) GetReferenceID() map[string]interface{}`

GetReferenceID returns the ReferenceID field if non-nil, zero value otherwise.

### GetReferenceIDOk

`func (o *PaymentServiceItem) GetReferenceIDOk() (*map[string]interface{}, bool)`

GetReferenceIDOk returns a tuple with the ReferenceID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceID

`func (o *PaymentServiceItem) SetReferenceID(v map[string]interface{})`

SetReferenceID sets ReferenceID field to given value.

### HasReferenceID

`func (o *PaymentServiceItem) HasReferenceID() bool`

HasReferenceID returns a boolean if a field has been set.

### GetPaymentServiceItemParams

`func (o *PaymentServiceItem) GetPaymentServiceItemParams() []PaymentServiceItemParam`

GetPaymentServiceItemParams returns the PaymentServiceItemParams field if non-nil, zero value otherwise.

### GetPaymentServiceItemParamsOk

`func (o *PaymentServiceItem) GetPaymentServiceItemParamsOk() (*[]PaymentServiceItemParam, bool)`

GetPaymentServiceItemParamsOk returns a tuple with the PaymentServiceItemParams field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceItemParams

`func (o *PaymentServiceItem) SetPaymentServiceItemParams(v []PaymentServiceItemParam)`

SetPaymentServiceItemParams sets PaymentServiceItemParams field to given value.

### HasPaymentServiceItemParams

`func (o *PaymentServiceItem) HasPaymentServiceItemParams() bool`

HasPaymentServiceItemParams returns a boolean if a field has been set.

### GetETag

`func (o *PaymentServiceItem) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *PaymentServiceItem) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *PaymentServiceItem) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *PaymentServiceItem) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


