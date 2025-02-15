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

// checks if the ProofOfServiceDocV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProofOfServiceDocV3{}

// ProofOfServiceDocV3 struct for ProofOfServiceDocV3
type ProofOfServiceDocV3 struct {
	Uploads []UploadWithOmissionsV3V3 `json:"uploads,omitempty"`
}

// NewProofOfServiceDocV3 instantiates a new ProofOfServiceDocV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProofOfServiceDocV3() *ProofOfServiceDocV3 {
	this := ProofOfServiceDocV3{}
	return &this
}

// NewProofOfServiceDocV3WithDefaults instantiates a new ProofOfServiceDocV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProofOfServiceDocV3WithDefaults() *ProofOfServiceDocV3 {
	this := ProofOfServiceDocV3{}
	return &this
}

// GetUploads returns the Uploads field value if set, zero value otherwise.
func (o *ProofOfServiceDocV3) GetUploads() []UploadWithOmissionsV3V3 {
	if o == nil || IsNil(o.Uploads) {
		var ret []UploadWithOmissionsV3V3
		return ret
	}
	return o.Uploads
}

// GetUploadsOk returns a tuple with the Uploads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProofOfServiceDocV3) GetUploadsOk() ([]UploadWithOmissionsV3V3, bool) {
	if o == nil || IsNil(o.Uploads) {
		return nil, false
	}
	return o.Uploads, true
}

// HasUploads returns a boolean if a field has been set.
func (o *ProofOfServiceDocV3) HasUploads() bool {
	if o != nil && !IsNil(o.Uploads) {
		return true
	}

	return false
}

// SetUploads gets a reference to the given []UploadWithOmissionsV3V3 and assigns it to the Uploads field.
func (o *ProofOfServiceDocV3) SetUploads(v []UploadWithOmissionsV3V3) {
	o.Uploads = v
}

func (o ProofOfServiceDocV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProofOfServiceDocV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uploads) {
		toSerialize["uploads"] = o.Uploads
	}
	return toSerialize, nil
}

type NullableProofOfServiceDocV3 struct {
	value *ProofOfServiceDocV3
	isSet bool
}

func (v NullableProofOfServiceDocV3) Get() *ProofOfServiceDocV3 {
	return v.value
}

func (v *NullableProofOfServiceDocV3) Set(val *ProofOfServiceDocV3) {
	v.value = val
	v.isSet = true
}

func (v NullableProofOfServiceDocV3) IsSet() bool {
	return v.isSet
}

func (v *NullableProofOfServiceDocV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProofOfServiceDocV3(val *ProofOfServiceDocV3) *NullableProofOfServiceDocV3 {
	return &NullableProofOfServiceDocV3{value: val, isSet: true}
}

func (v NullableProofOfServiceDocV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProofOfServiceDocV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


