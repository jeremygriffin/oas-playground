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
)

// checks if the ProofOfServiceDocV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProofOfServiceDocV2{}

// ProofOfServiceDocV2 struct for ProofOfServiceDocV2
type ProofOfServiceDocV2 struct {
	Uploads []UploadWithOmissionsV2V2 `json:"uploads,omitempty"`
}

// NewProofOfServiceDocV2 instantiates a new ProofOfServiceDocV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProofOfServiceDocV2() *ProofOfServiceDocV2 {
	this := ProofOfServiceDocV2{}
	return &this
}

// NewProofOfServiceDocV2WithDefaults instantiates a new ProofOfServiceDocV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProofOfServiceDocV2WithDefaults() *ProofOfServiceDocV2 {
	this := ProofOfServiceDocV2{}
	return &this
}

// GetUploads returns the Uploads field value if set, zero value otherwise.
func (o *ProofOfServiceDocV2) GetUploads() []UploadWithOmissionsV2V2 {
	if o == nil || IsNil(o.Uploads) {
		var ret []UploadWithOmissionsV2V2
		return ret
	}
	return o.Uploads
}

// GetUploadsOk returns a tuple with the Uploads field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProofOfServiceDocV2) GetUploadsOk() ([]UploadWithOmissionsV2V2, bool) {
	if o == nil || IsNil(o.Uploads) {
		return nil, false
	}
	return o.Uploads, true
}

// HasUploads returns a boolean if a field has been set.
func (o *ProofOfServiceDocV2) HasUploads() bool {
	if o != nil && !IsNil(o.Uploads) {
		return true
	}

	return false
}

// SetUploads gets a reference to the given []UploadWithOmissionsV2V2 and assigns it to the Uploads field.
func (o *ProofOfServiceDocV2) SetUploads(v []UploadWithOmissionsV2V2) {
	o.Uploads = v
}

func (o ProofOfServiceDocV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProofOfServiceDocV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uploads) {
		toSerialize["uploads"] = o.Uploads
	}
	return toSerialize, nil
}

type NullableProofOfServiceDocV2 struct {
	value *ProofOfServiceDocV2
	isSet bool
}

func (v NullableProofOfServiceDocV2) Get() *ProofOfServiceDocV2 {
	return v.value
}

func (v *NullableProofOfServiceDocV2) Set(val *ProofOfServiceDocV2) {
	v.value = val
	v.isSet = true
}

func (v NullableProofOfServiceDocV2) IsSet() bool {
	return v.isSet
}

func (v *NullableProofOfServiceDocV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProofOfServiceDocV2(val *ProofOfServiceDocV2) *NullableProofOfServiceDocV2 {
	return &NullableProofOfServiceDocV2{value: val, isSet: true}
}

func (v NullableProofOfServiceDocV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProofOfServiceDocV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


