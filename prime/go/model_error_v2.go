/*
MilMove Prime V2 API

The Prime V2 API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v2/`. 

API version: 0.0.1
Contact: milmove-developers@caci.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the ErrorV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorV2{}

// ErrorV2 struct for ErrorV2
type ErrorV2 struct {
	Title string `json:"title"`
	Detail string `json:"detail"`
	Instance *string `json:"instance,omitempty"`
}

type _ErrorV2 ErrorV2

// NewErrorV2 instantiates a new ErrorV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorV2(title string, detail string) *ErrorV2 {
	this := ErrorV2{}
	this.Title = title
	this.Detail = detail
	return &this
}

// NewErrorV2WithDefaults instantiates a new ErrorV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorV2WithDefaults() *ErrorV2 {
	this := ErrorV2{}
	return &this
}

// GetTitle returns the Title field value
func (o *ErrorV2) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ErrorV2) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ErrorV2) SetTitle(v string) {
	o.Title = v
}

// GetDetail returns the Detail field value
func (o *ErrorV2) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *ErrorV2) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *ErrorV2) SetDetail(v string) {
	o.Detail = v
}

// GetInstance returns the Instance field value if set, zero value otherwise.
func (o *ErrorV2) GetInstance() string {
	if o == nil || IsNil(o.Instance) {
		var ret string
		return ret
	}
	return *o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorV2) GetInstanceOk() (*string, bool) {
	if o == nil || IsNil(o.Instance) {
		return nil, false
	}
	return o.Instance, true
}

// HasInstance returns a boolean if a field has been set.
func (o *ErrorV2) HasInstance() bool {
	if o != nil && !IsNil(o.Instance) {
		return true
	}

	return false
}

// SetInstance gets a reference to the given string and assigns it to the Instance field.
func (o *ErrorV2) SetInstance(v string) {
	o.Instance = &v
}

func (o ErrorV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["detail"] = o.Detail
	if !IsNil(o.Instance) {
		toSerialize["instance"] = o.Instance
	}
	return toSerialize, nil
}

func (o *ErrorV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"detail",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varErrorV2 := _ErrorV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varErrorV2)

	if err != nil {
		return err
	}

	*o = ErrorV2(varErrorV2)

	return err
}

type NullableErrorV2 struct {
	value *ErrorV2
	isSet bool
}

func (v NullableErrorV2) Get() *ErrorV2 {
	return v.value
}

func (v *NullableErrorV2) Set(val *ErrorV2) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorV2) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorV2(val *ErrorV2) *NullableErrorV2 {
	return &NullableErrorV2{value: val, isSet: true}
}

func (v NullableErrorV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


