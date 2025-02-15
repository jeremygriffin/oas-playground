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

// MTOShipmentType The type of shipment.   * `HHG` = Household goods move   * `HHG_INTO_NTS_DOMESTIC` = HHG into Non-temporary storage (NTS)   * `HHG_OUTOF_NTS_DOMESTIC` = HHG out of Non-temporary storage (NTS Release)   * `PPM` = Personally Procured Move also known as Do It Yourself (DITY) 
type MTOShipmentType string

// List of MTOShipmentType
const (
	BOAT_HAUL_AWAY MTOShipmentType = "BOAT_HAUL_AWAY"
	BOAT_TOW_AWAY MTOShipmentType = "BOAT_TOW_AWAY"
	HHG MTOShipmentType = "HHG"
	HHG_INTO_NTS_DOMESTIC MTOShipmentType = "HHG_INTO_NTS_DOMESTIC"
	HHG_OUTOF_NTS_DOMESTIC MTOShipmentType = "HHG_OUTOF_NTS_DOMESTIC"
	INTERNATIONAL_HHG MTOShipmentType = "INTERNATIONAL_HHG"
	INTERNATIONAL_UB MTOShipmentType = "INTERNATIONAL_UB"
	MOTORHOME MTOShipmentType = "MOTORHOME"
	PPM MTOShipmentType = "PPM"
)

// All allowed values of MTOShipmentType enum
var AllowedMTOShipmentTypeEnumValues = []MTOShipmentType{
	"BOAT_HAUL_AWAY",
	"BOAT_TOW_AWAY",
	"HHG",
	"HHG_INTO_NTS_DOMESTIC",
	"HHG_OUTOF_NTS_DOMESTIC",
	"INTERNATIONAL_HHG",
	"INTERNATIONAL_UB",
	"MOTORHOME",
	"PPM",
}

func (v *MTOShipmentType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MTOShipmentType(value)
	for _, existing := range AllowedMTOShipmentTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MTOShipmentType", value)
}

// NewMTOShipmentTypeFromValue returns a pointer to a valid MTOShipmentType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMTOShipmentTypeFromValue(v string) (*MTOShipmentType, error) {
	ev := MTOShipmentType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MTOShipmentType: valid values are %v", v, AllowedMTOShipmentTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MTOShipmentType) IsValid() bool {
	for _, existing := range AllowedMTOShipmentTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MTOShipmentType value
func (v MTOShipmentType) Ptr() *MTOShipmentType {
	return &v
}

type NullableMTOShipmentType struct {
	value *MTOShipmentType
	isSet bool
}

func (v NullableMTOShipmentType) Get() *MTOShipmentType {
	return v.value
}

func (v *NullableMTOShipmentType) Set(val *MTOShipmentType) {
	v.value = val
	v.isSet = true
}

func (v NullableMTOShipmentType) IsSet() bool {
	return v.isSet
}

func (v *NullableMTOShipmentType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMTOShipmentType(val *MTOShipmentType) *NullableMTOShipmentType {
	return &NullableMTOShipmentType{value: val, isSet: true}
}

func (v NullableMTOShipmentType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMTOShipmentType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

