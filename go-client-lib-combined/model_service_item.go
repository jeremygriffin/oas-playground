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

// checks if the ServiceItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ServiceItem{}

// ServiceItem struct for ServiceItem
type ServiceItem struct {
	Id *string `json:"id,omitempty"`
	// This should be populated for the following service items:   * DOASIT(Domestic origin Additional day SIT)   * DDASIT(Domestic destination Additional day SIT)  Both take in the following param keys:   * `SITPaymentRequestStart`   * `SITPaymentRequestEnd`  The value of each is a date string in the format \"YYYY-MM-DD\" (e.g. \"2023-01-15\") 
	Params []ServiceItemParamsInner `json:"params,omitempty"`
	ETag *string `json:"eTag,omitempty"`
}

// NewServiceItem instantiates a new ServiceItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceItem() *ServiceItem {
	this := ServiceItem{}
	return &this
}

// NewServiceItemWithDefaults instantiates a new ServiceItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceItemWithDefaults() *ServiceItem {
	this := ServiceItem{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ServiceItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ServiceItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ServiceItem) SetId(v string) {
	o.Id = &v
}

// GetParams returns the Params field value if set, zero value otherwise.
func (o *ServiceItem) GetParams() []ServiceItemParamsInner {
	if o == nil || IsNil(o.Params) {
		var ret []ServiceItemParamsInner
		return ret
	}
	return o.Params
}

// GetParamsOk returns a tuple with the Params field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceItem) GetParamsOk() ([]ServiceItemParamsInner, bool) {
	if o == nil || IsNil(o.Params) {
		return nil, false
	}
	return o.Params, true
}

// HasParams returns a boolean if a field has been set.
func (o *ServiceItem) HasParams() bool {
	if o != nil && !IsNil(o.Params) {
		return true
	}

	return false
}

// SetParams gets a reference to the given []ServiceItemParamsInner and assigns it to the Params field.
func (o *ServiceItem) SetParams(v []ServiceItemParamsInner) {
	o.Params = v
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *ServiceItem) GetETag() string {
	if o == nil || IsNil(o.ETag) {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ServiceItem) GetETagOk() (*string, bool) {
	if o == nil || IsNil(o.ETag) {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *ServiceItem) HasETag() bool {
	if o != nil && !IsNil(o.ETag) {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *ServiceItem) SetETag(v string) {
	o.ETag = &v
}

func (o ServiceItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Params) {
		toSerialize["params"] = o.Params
	}
	if !IsNil(o.ETag) {
		toSerialize["eTag"] = o.ETag
	}
	return toSerialize, nil
}

type NullableServiceItem struct {
	value *ServiceItem
	isSet bool
}

func (v NullableServiceItem) Get() *ServiceItem {
	return v.value
}

func (v *NullableServiceItem) Set(val *ServiceItem) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceItem) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceItem(val *ServiceItem) *NullableServiceItem {
	return &NullableServiceItem{value: val, isSet: true}
}

func (v NullableServiceItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


