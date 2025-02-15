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
	"fmt"
)

// MTOServiceItemModelTypeV3 Describes all model sub-types for a MTOServiceItem model.  Using this list, choose the correct modelType in the dropdown, corresponding to the service item type.   * DOFSIT, DOASIT - MTOServiceItemOriginSIT   * DDFSIT, DDASIT - MTOServiceItemDestSIT   * DOSHUT, DDSHUT - MTOServiceItemShuttle   * DCRT, DUCRT - MTOServiceItemDomesticCrating  The documentation will then update with the supported fields. 
type MTOServiceItemModelTypeV3 string

// List of MTOServiceItemModelType
const (
	MTO_SERVICE_ITEM_BASIC MTOServiceItemModelTypeV3 = "MTOServiceItemBasic"
	MTO_SERVICE_ITEM_ORIGIN_SIT MTOServiceItemModelTypeV3 = "MTOServiceItemOriginSIT"
	MTO_SERVICE_ITEM_DEST_SIT MTOServiceItemModelTypeV3 = "MTOServiceItemDestSIT"
	MTO_SERVICE_ITEM_SHUTTLE MTOServiceItemModelTypeV3 = "MTOServiceItemShuttle"
	MTO_SERVICE_ITEM_DOMESTIC_CRATING MTOServiceItemModelTypeV3 = "MTOServiceItemDomesticCrating"
)

// All allowed values of MTOServiceItemModelTypeV3 enum
var AllowedMTOServiceItemModelTypeV3EnumValues = []MTOServiceItemModelTypeV3{
	"MTOServiceItemBasic",
	"MTOServiceItemOriginSIT",
	"MTOServiceItemDestSIT",
	"MTOServiceItemShuttle",
	"MTOServiceItemDomesticCrating",
}

func (v *MTOServiceItemModelTypeV3) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MTOServiceItemModelTypeV3(value)
	for _, existing := range AllowedMTOServiceItemModelTypeV3EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MTOServiceItemModelTypeV3", value)
}

// NewMTOServiceItemModelTypeV3FromValue returns a pointer to a valid MTOServiceItemModelTypeV3
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMTOServiceItemModelTypeV3FromValue(v string) (*MTOServiceItemModelTypeV3, error) {
	ev := MTOServiceItemModelTypeV3(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MTOServiceItemModelTypeV3: valid values are %v", v, AllowedMTOServiceItemModelTypeV3EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MTOServiceItemModelTypeV3) IsValid() bool {
	for _, existing := range AllowedMTOServiceItemModelTypeV3EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MTOServiceItemModelType value
func (v MTOServiceItemModelTypeV3) Ptr() *MTOServiceItemModelTypeV3 {
	return &v
}

type NullableMTOServiceItemModelTypeV3 struct {
	value *MTOServiceItemModelTypeV3
	isSet bool
}

func (v NullableMTOServiceItemModelTypeV3) Get() *MTOServiceItemModelTypeV3 {
	return v.value
}

func (v *NullableMTOServiceItemModelTypeV3) Set(val *MTOServiceItemModelTypeV3) {
	v.value = val
	v.isSet = true
}

func (v NullableMTOServiceItemModelTypeV3) IsSet() bool {
	return v.isSet
}

func (v *NullableMTOServiceItemModelTypeV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMTOServiceItemModelTypeV3(val *MTOServiceItemModelTypeV3) *NullableMTOServiceItemModelTypeV3 {
	return &NullableMTOServiceItemModelTypeV3{value: val, isSet: true}
}

func (v NullableMTOServiceItemModelTypeV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMTOServiceItemModelTypeV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

