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

// MTOAgentType The type for this agent. `RELEASING` means they have authority on pickup, `RECEIVING` means they can receive the shipment on delivery. 
type MTOAgentType string

// List of MTOAgentType
const (
	RELEASING_AGENT MTOAgentType = "RELEASING_AGENT"
	RECEIVING_AGENT MTOAgentType = "RECEIVING_AGENT"
)

// All allowed values of MTOAgentType enum
var AllowedMTOAgentTypeEnumValues = []MTOAgentType{
	"RELEASING_AGENT",
	"RECEIVING_AGENT",
}

func (v *MTOAgentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MTOAgentType(value)
	for _, existing := range AllowedMTOAgentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MTOAgentType", value)
}

// NewMTOAgentTypeFromValue returns a pointer to a valid MTOAgentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMTOAgentTypeFromValue(v string) (*MTOAgentType, error) {
	ev := MTOAgentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MTOAgentType: valid values are %v", v, AllowedMTOAgentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MTOAgentType) IsValid() bool {
	for _, existing := range AllowedMTOAgentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MTOAgentType value
func (v MTOAgentType) Ptr() *MTOAgentType {
	return &v
}

type NullableMTOAgentType struct {
	value *MTOAgentType
	isSet bool
}

func (v NullableMTOAgentType) Get() *MTOAgentType {
	return v.value
}

func (v *NullableMTOAgentType) Set(val *MTOAgentType) {
	v.value = val
	v.isSet = true
}

func (v NullableMTOAgentType) IsSet() bool {
	return v.isSet
}

func (v *NullableMTOAgentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMTOAgentType(val *MTOAgentType) *NullableMTOAgentType {
	return &NullableMTOAgentType{value: val, isSet: true}
}

func (v NullableMTOAgentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMTOAgentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

