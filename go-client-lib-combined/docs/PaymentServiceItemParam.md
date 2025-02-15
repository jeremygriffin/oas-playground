# PaymentServiceItemParam

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**PaymentServiceItemID** | Pointer to **string** |  | [optional] 
**Key** | Pointer to [**ServiceItemParamName**](ServiceItemParamName.md) |  | [optional] 
**Value** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**ServiceItemParamType**](ServiceItemParamType.md) |  | [optional] 
**Origin** | Pointer to [**ServiceItemParamOrigin**](ServiceItemParamOrigin.md) |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewPaymentServiceItemParam

`func NewPaymentServiceItemParam() *PaymentServiceItemParam`

NewPaymentServiceItemParam instantiates a new PaymentServiceItemParam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentServiceItemParamWithDefaults

`func NewPaymentServiceItemParamWithDefaults() *PaymentServiceItemParam`

NewPaymentServiceItemParamWithDefaults instantiates a new PaymentServiceItemParam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PaymentServiceItemParam) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PaymentServiceItemParam) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PaymentServiceItemParam) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PaymentServiceItemParam) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPaymentServiceItemID

`func (o *PaymentServiceItemParam) GetPaymentServiceItemID() string`

GetPaymentServiceItemID returns the PaymentServiceItemID field if non-nil, zero value otherwise.

### GetPaymentServiceItemIDOk

`func (o *PaymentServiceItemParam) GetPaymentServiceItemIDOk() (*string, bool)`

GetPaymentServiceItemIDOk returns a tuple with the PaymentServiceItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentServiceItemID

`func (o *PaymentServiceItemParam) SetPaymentServiceItemID(v string)`

SetPaymentServiceItemID sets PaymentServiceItemID field to given value.

### HasPaymentServiceItemID

`func (o *PaymentServiceItemParam) HasPaymentServiceItemID() bool`

HasPaymentServiceItemID returns a boolean if a field has been set.

### GetKey

`func (o *PaymentServiceItemParam) GetKey() ServiceItemParamName`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *PaymentServiceItemParam) GetKeyOk() (*ServiceItemParamName, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *PaymentServiceItemParam) SetKey(v ServiceItemParamName)`

SetKey sets Key field to given value.

### HasKey

`func (o *PaymentServiceItemParam) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetValue

`func (o *PaymentServiceItemParam) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PaymentServiceItemParam) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PaymentServiceItemParam) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *PaymentServiceItemParam) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetType

`func (o *PaymentServiceItemParam) GetType() ServiceItemParamType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PaymentServiceItemParam) GetTypeOk() (*ServiceItemParamType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PaymentServiceItemParam) SetType(v ServiceItemParamType)`

SetType sets Type field to given value.

### HasType

`func (o *PaymentServiceItemParam) HasType() bool`

HasType returns a boolean if a field has been set.

### GetOrigin

`func (o *PaymentServiceItemParam) GetOrigin() ServiceItemParamOrigin`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *PaymentServiceItemParam) GetOriginOk() (*ServiceItemParamOrigin, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *PaymentServiceItemParam) SetOrigin(v ServiceItemParamOrigin)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *PaymentServiceItemParam) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.

### GetETag

`func (o *PaymentServiceItemParam) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *PaymentServiceItemParam) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *PaymentServiceItemParam) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *PaymentServiceItemParam) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


