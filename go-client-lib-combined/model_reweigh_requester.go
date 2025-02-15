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
	"fmt"
)

// ReweighRequester the model 'ReweighRequester'
type ReweighRequester string

// List of ReweighRequester
const (
	CUSTOMER ReweighRequester = "CUSTOMER"
	PRIME ReweighRequester = "PRIME"
	SYSTEM ReweighRequester = "SYSTEM"
	TOO ReweighRequester = "TOO"
)

// All allowed values of ReweighRequester enum
var AllowedReweighRequesterEnumValues = []ReweighRequester{
	"CUSTOMER",
	"PRIME",
	"SYSTEM",
	"TOO",
}

func (v *ReweighRequester) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReweighRequester(value)
	for _, existing := range AllowedReweighRequesterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReweighRequester", value)
}

// NewReweighRequesterFromValue returns a pointer to a valid ReweighRequester
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReweighRequesterFromValue(v string) (*ReweighRequester, error) {
	ev := ReweighRequester(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReweighRequester: valid values are %v", v, AllowedReweighRequesterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReweighRequester) IsValid() bool {
	for _, existing := range AllowedReweighRequesterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReweighRequester value
func (v ReweighRequester) Ptr() *ReweighRequester {
	return &v
}

type NullableReweighRequester struct {
	value *ReweighRequester
	isSet bool
}

func (v NullableReweighRequester) Get() *ReweighRequester {
	return v.value
}

func (v *NullableReweighRequester) Set(val *ReweighRequester) {
	v.value = val
	v.isSet = true
}

func (v NullableReweighRequester) IsSet() bool {
	return v.isSet
}

func (v *NullableReweighRequester) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReweighRequester(val *ReweighRequester) *NullableReweighRequester {
	return &NullableReweighRequester{value: val, isSet: true}
}

func (v NullableReweighRequester) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReweighRequester) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

