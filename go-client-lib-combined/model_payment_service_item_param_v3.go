/*
MilMove Prime V3 API

The Prime V3 API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v3/`. 

API version: 0.0.1
Contact: milmove-developers@caci.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the PaymentServiceItemParamV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PaymentServiceItemParamV3{}

// PaymentServiceItemParamV3 struct for PaymentServiceItemParamV3
type PaymentServiceItemParamV3 struct {
	Id *string `json:"id,omitempty"`
	PaymentServiceItemID *string `json:"paymentServiceItemID,omitempty"`
	Key *ServiceItemParamNameV3V3 `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
	Type *ServiceItemParamTypeV3V3 `json:"type,omitempty"`
	Origin *ServiceItemParamOriginV3V3 `json:"origin,omitempty"`
	ETag *string `json:"eTag,omitempty"`
}

// NewPaymentServiceItemParamV3 instantiates a new PaymentServiceItemParamV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPaymentServiceItemParamV3() *PaymentServiceItemParamV3 {
	this := PaymentServiceItemParamV3{}
	return &this
}

// NewPaymentServiceItemParamV3WithDefaults instantiates a new PaymentServiceItemParamV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPaymentServiceItemParamV3WithDefaults() *PaymentServiceItemParamV3 {
	this := PaymentServiceItemParamV3{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PaymentServiceItemParamV3) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceItemParamV3) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PaymentServiceItemParamV3) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PaymentServiceItemParamV3) SetId(v string) {
	o.Id = &v
}

// GetPaymentServiceItemID returns the PaymentServiceItemID field value if set, zero value otherwise.
func (o *PaymentServiceItemParamV3) GetPaymentServiceItemID() string {
	if o == nil || IsNil(o.PaymentServiceItemID) {
		var ret string
		return ret
	}
	return *o.PaymentServiceItemID
}

// GetPaymentServiceItemIDOk returns a tuple with the PaymentServiceItemID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceItemParamV3) GetPaymentServiceItemIDOk() (*string, bool) {
	if o == nil || IsNil(o.PaymentServiceItemID) {
		return nil, false
	}
	return o.PaymentServiceItemID, true
}

// HasPaymentServiceItemID returns a boolean if a field has been set.
func (o *PaymentServiceItemParamV3) HasPaymentServiceItemID() bool {
	if o != nil && !IsNil(o.PaymentServiceItemID) {
		return true
	}

	return false
}

// SetPaymentServiceItemID gets a reference to the given string and assigns it to the PaymentServiceItemID field.
func (o *PaymentServiceItemParamV3) SetPaymentServiceItemID(v string) {
	o.PaymentServiceItemID = &v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *PaymentServiceItemParamV3) GetKey() ServiceItemParamNameV3V3 {
	if o == nil || IsNil(o.Key) {
		var ret ServiceItemParamNameV3V3
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceItemParamV3) GetKeyOk() (*ServiceItemParamNameV3V3, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *PaymentServiceItemParamV3) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given ServiceItemParamNameV3V3 and assigns it to the Key field.
func (o *PaymentServiceItemParamV3) SetKey(v ServiceItemParamNameV3V3) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *PaymentServiceItemParamV3) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceItemParamV3) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *PaymentServiceItemParamV3) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *PaymentServiceItemParamV3) SetValue(v string) {
	o.Value = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *PaymentServiceItemParamV3) GetType() ServiceItemParamTypeV3V3 {
	if o == nil || IsNil(o.Type) {
		var ret ServiceItemParamTypeV3V3
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceItemParamV3) GetTypeOk() (*ServiceItemParamTypeV3V3, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *PaymentServiceItemParamV3) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given ServiceItemParamTypeV3V3 and assigns it to the Type field.
func (o *PaymentServiceItemParamV3) SetType(v ServiceItemParamTypeV3V3) {
	o.Type = &v
}

// GetOrigin returns the Origin field value if set, zero value otherwise.
func (o *PaymentServiceItemParamV3) GetOrigin() ServiceItemParamOriginV3V3 {
	if o == nil || IsNil(o.Origin) {
		var ret ServiceItemParamOriginV3V3
		return ret
	}
	return *o.Origin
}

// GetOriginOk returns a tuple with the Origin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceItemParamV3) GetOriginOk() (*ServiceItemParamOriginV3V3, bool) {
	if o == nil || IsNil(o.Origin) {
		return nil, false
	}
	return o.Origin, true
}

// HasOrigin returns a boolean if a field has been set.
func (o *PaymentServiceItemParamV3) HasOrigin() bool {
	if o != nil && !IsNil(o.Origin) {
		return true
	}

	return false
}

// SetOrigin gets a reference to the given ServiceItemParamOriginV3V3 and assigns it to the Origin field.
func (o *PaymentServiceItemParamV3) SetOrigin(v ServiceItemParamOriginV3V3) {
	o.Origin = &v
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *PaymentServiceItemParamV3) GetETag() string {
	if o == nil || IsNil(o.ETag) {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PaymentServiceItemParamV3) GetETagOk() (*string, bool) {
	if o == nil || IsNil(o.ETag) {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *PaymentServiceItemParamV3) HasETag() bool {
	if o != nil && !IsNil(o.ETag) {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *PaymentServiceItemParamV3) SetETag(v string) {
	o.ETag = &v
}

func (o PaymentServiceItemParamV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PaymentServiceItemParamV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.PaymentServiceItemID) {
		toSerialize["paymentServiceItemID"] = o.PaymentServiceItemID
	}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Origin) {
		toSerialize["origin"] = o.Origin
	}
	if !IsNil(o.ETag) {
		toSerialize["eTag"] = o.ETag
	}
	return toSerialize, nil
}

type NullablePaymentServiceItemParamV3 struct {
	value *PaymentServiceItemParamV3
	isSet bool
}

func (v NullablePaymentServiceItemParamV3) Get() *PaymentServiceItemParamV3 {
	return v.value
}

func (v *NullablePaymentServiceItemParamV3) Set(val *PaymentServiceItemParamV3) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServiceItemParamV3) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServiceItemParamV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServiceItemParamV3(val *PaymentServiceItemParamV3) *NullablePaymentServiceItemParamV3 {
	return &NullablePaymentServiceItemParamV3{value: val, isSet: true}
}

func (v NullablePaymentServiceItemParamV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServiceItemParamV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


