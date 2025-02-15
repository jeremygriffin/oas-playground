/*
MilMove Prime API

The Prime API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v1/`. 

API version: 0.0.1
Contact: milmove-developers@caci.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the ServiceItemParamsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceItemParamsInner{}

// ServiceItemParamsInner struct for ServiceItemParamsInner
type ServiceItemParamsInner struct {
	Key *string `json:"key,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewServiceItemParamsInner instantiates a new ServiceItemParamsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceItemParamsInner() *ServiceItemParamsInner {
	this := ServiceItemParamsInner{}
	return &this
}

// NewServiceItemParamsInnerWithDefaults instantiates a new ServiceItemParamsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceItemParamsInnerWithDefaults() *ServiceItemParamsInner {
	this := ServiceItemParamsInner{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *ServiceItemParamsInner) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceItemParamsInner) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *ServiceItemParamsInner) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *ServiceItemParamsInner) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *ServiceItemParamsInner) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceItemParamsInner) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *ServiceItemParamsInner) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *ServiceItemParamsInner) SetValue(v string) {
	o.Value = &v
}

func (o ServiceItemParamsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceItemParamsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableServiceItemParamsInner struct {
	value *ServiceItemParamsInner
	isSet bool
}

func (v NullableServiceItemParamsInner) Get() *ServiceItemParamsInner {
	return v.value
}

func (v *NullableServiceItemParamsInner) Set(val *ServiceItemParamsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceItemParamsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceItemParamsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceItemParamsInner(val *ServiceItemParamsInner) *NullableServiceItemParamsInner {
	return &NullableServiceItemParamsInner{value: val, isSet: true}
}

func (v NullableServiceItemParamsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceItemParamsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


