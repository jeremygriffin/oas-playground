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
	"bytes"
	"fmt"
)

// checks if the ValidationErrorV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ValidationErrorV3{}

// ValidationErrorV3 struct for ValidationErrorV3
type ValidationErrorV3 struct {
	Title string `json:"title"`
	Detail string `json:"detail"`
	Instance string `json:"instance"`
	InvalidFields map[string][]string `json:"invalidFields"`
}

type _ValidationErrorV3 ValidationErrorV3

// NewValidationErrorV3 instantiates a new ValidationErrorV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewValidationErrorV3(title string, detail string, instance string, invalidFields map[string][]string) *ValidationErrorV3 {
	this := ValidationErrorV3{}
	this.Title = title
	this.Detail = detail
	this.Instance = instance
	this.InvalidFields = invalidFields
	return &this
}

// NewValidationErrorV3WithDefaults instantiates a new ValidationErrorV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewValidationErrorV3WithDefaults() *ValidationErrorV3 {
	this := ValidationErrorV3{}
	return &this
}

// GetTitle returns the Title field value
func (o *ValidationErrorV3) GetTitle() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Title
}

// GetTitleOk returns a tuple with the Title field value
// and a boolean to check if the value has been set.
func (o *ValidationErrorV3) GetTitleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Title, true
}

// SetTitle sets field value
func (o *ValidationErrorV3) SetTitle(v string) {
	o.Title = v
}

// GetDetail returns the Detail field value
func (o *ValidationErrorV3) GetDetail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *ValidationErrorV3) GetDetailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Detail, true
}

// SetDetail sets field value
func (o *ValidationErrorV3) SetDetail(v string) {
	o.Detail = v
}

// GetInstance returns the Instance field value
func (o *ValidationErrorV3) GetInstance() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Instance
}

// GetInstanceOk returns a tuple with the Instance field value
// and a boolean to check if the value has been set.
func (o *ValidationErrorV3) GetInstanceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Instance, true
}

// SetInstance sets field value
func (o *ValidationErrorV3) SetInstance(v string) {
	o.Instance = v
}

// GetInvalidFields returns the InvalidFields field value
func (o *ValidationErrorV3) GetInvalidFields() map[string][]string {
	if o == nil {
		var ret map[string][]string
		return ret
	}

	return o.InvalidFields
}

// GetInvalidFieldsOk returns a tuple with the InvalidFields field value
// and a boolean to check if the value has been set.
func (o *ValidationErrorV3) GetInvalidFieldsOk() (*map[string][]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InvalidFields, true
}

// SetInvalidFields sets field value
func (o *ValidationErrorV3) SetInvalidFields(v map[string][]string) {
	o.InvalidFields = v
}

func (o ValidationErrorV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ValidationErrorV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["title"] = o.Title
	toSerialize["detail"] = o.Detail
	toSerialize["instance"] = o.Instance
	toSerialize["invalidFields"] = o.InvalidFields
	return toSerialize, nil
}

func (o *ValidationErrorV3) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"title",
		"detail",
		"instance",
		"invalidFields",
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

	varValidationErrorV3 := _ValidationErrorV3{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varValidationErrorV3)

	if err != nil {
		return err
	}

	*o = ValidationErrorV3(varValidationErrorV3)

	return err
}

type NullableValidationErrorV3 struct {
	value *ValidationErrorV3
	isSet bool
}

func (v NullableValidationErrorV3) Get() *ValidationErrorV3 {
	return v.value
}

func (v *NullableValidationErrorV3) Set(val *ValidationErrorV3) {
	v.value = val
	v.isSet = true
}

func (v NullableValidationErrorV3) IsSet() bool {
	return v.isSet
}

func (v *NullableValidationErrorV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValidationErrorV3(val *ValidationErrorV3) *NullableValidationErrorV3 {
	return &NullableValidationErrorV3{value: val, isSet: true}
}

func (v NullableValidationErrorV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValidationErrorV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


