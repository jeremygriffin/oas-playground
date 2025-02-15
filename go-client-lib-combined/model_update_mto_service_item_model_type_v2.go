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
	"fmt"
)

// UpdateMTOServiceItemModelTypeV2 Using this list, choose the correct modelType in the dropdown, corresponding to the service item type.   * DDDSIT - UpdateMTOServiceItemSIT   * DOPSIT - UpdateMTOServiceItemSIT   * DOASIT - UpdateMTOServiceItemSIT   * DOFSIT - UpdateMTOServiceItemSIT   * DDSHUT - UpdateMTOServiceItemShuttle   * DOSHUT - UpdateMTOServiceItemShuttle  The documentation will then update with the supported fields. 
type UpdateMTOServiceItemModelTypeV2 string

// List of UpdateMTOServiceItemModelType
const (
	UPDATE_MTO_SERVICE_ITEM_SIT UpdateMTOServiceItemModelTypeV2 = "UpdateMTOServiceItemSIT"
	UPDATE_MTO_SERVICE_ITEM_SHUTTLE UpdateMTOServiceItemModelTypeV2 = "UpdateMTOServiceItemShuttle"
)

// All allowed values of UpdateMTOServiceItemModelTypeV2 enum
var AllowedUpdateMTOServiceItemModelTypeV2EnumValues = []UpdateMTOServiceItemModelTypeV2{
	"UpdateMTOServiceItemSIT",
	"UpdateMTOServiceItemShuttle",
}

func (v *UpdateMTOServiceItemModelTypeV2) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UpdateMTOServiceItemModelTypeV2(value)
	for _, existing := range AllowedUpdateMTOServiceItemModelTypeV2EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UpdateMTOServiceItemModelTypeV2", value)
}

// NewUpdateMTOServiceItemModelTypeV2FromValue returns a pointer to a valid UpdateMTOServiceItemModelTypeV2
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUpdateMTOServiceItemModelTypeV2FromValue(v string) (*UpdateMTOServiceItemModelTypeV2, error) {
	ev := UpdateMTOServiceItemModelTypeV2(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UpdateMTOServiceItemModelTypeV2: valid values are %v", v, AllowedUpdateMTOServiceItemModelTypeV2EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UpdateMTOServiceItemModelTypeV2) IsValid() bool {
	for _, existing := range AllowedUpdateMTOServiceItemModelTypeV2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateMTOServiceItemModelType value
func (v UpdateMTOServiceItemModelTypeV2) Ptr() *UpdateMTOServiceItemModelTypeV2 {
	return &v
}

type NullableUpdateMTOServiceItemModelTypeV2 struct {
	value *UpdateMTOServiceItemModelTypeV2
	isSet bool
}

func (v NullableUpdateMTOServiceItemModelTypeV2) Get() *UpdateMTOServiceItemModelTypeV2 {
	return v.value
}

func (v *NullableUpdateMTOServiceItemModelTypeV2) Set(val *UpdateMTOServiceItemModelTypeV2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMTOServiceItemModelTypeV2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMTOServiceItemModelTypeV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMTOServiceItemModelTypeV2(val *UpdateMTOServiceItemModelTypeV2) *NullableUpdateMTOServiceItemModelTypeV2 {
	return &NullableUpdateMTOServiceItemModelTypeV2{value: val, isSet: true}
}

func (v NullableUpdateMTOServiceItemModelTypeV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMTOServiceItemModelTypeV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

