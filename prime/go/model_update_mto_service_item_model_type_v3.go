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

// UpdateMTOServiceItemModelTypeV3 Using this list, choose the correct modelType in the dropdown, corresponding to the service item type.   * DDDSIT - UpdateMTOServiceItemSIT   * DOPSIT - UpdateMTOServiceItemSIT   * DOASIT - UpdateMTOServiceItemSIT   * DOFSIT - UpdateMTOServiceItemSIT   * DDSHUT - UpdateMTOServiceItemShuttle   * DOSHUT - UpdateMTOServiceItemShuttle  The documentation will then update with the supported fields. 
type UpdateMTOServiceItemModelTypeV3 string

// List of UpdateMTOServiceItemModelType
const (
	UPDATE_MTO_SERVICE_ITEM_SIT UpdateMTOServiceItemModelTypeV3 = "UpdateMTOServiceItemSIT"
	UPDATE_MTO_SERVICE_ITEM_SHUTTLE UpdateMTOServiceItemModelTypeV3 = "UpdateMTOServiceItemShuttle"
)

// All allowed values of UpdateMTOServiceItemModelTypeV3 enum
var AllowedUpdateMTOServiceItemModelTypeV3EnumValues = []UpdateMTOServiceItemModelTypeV3{
	"UpdateMTOServiceItemSIT",
	"UpdateMTOServiceItemShuttle",
}

func (v *UpdateMTOServiceItemModelTypeV3) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UpdateMTOServiceItemModelTypeV3(value)
	for _, existing := range AllowedUpdateMTOServiceItemModelTypeV3EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UpdateMTOServiceItemModelTypeV3", value)
}

// NewUpdateMTOServiceItemModelTypeV3FromValue returns a pointer to a valid UpdateMTOServiceItemModelTypeV3
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUpdateMTOServiceItemModelTypeV3FromValue(v string) (*UpdateMTOServiceItemModelTypeV3, error) {
	ev := UpdateMTOServiceItemModelTypeV3(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UpdateMTOServiceItemModelTypeV3: valid values are %v", v, AllowedUpdateMTOServiceItemModelTypeV3EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UpdateMTOServiceItemModelTypeV3) IsValid() bool {
	for _, existing := range AllowedUpdateMTOServiceItemModelTypeV3EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpdateMTOServiceItemModelType value
func (v UpdateMTOServiceItemModelTypeV3) Ptr() *UpdateMTOServiceItemModelTypeV3 {
	return &v
}

type NullableUpdateMTOServiceItemModelTypeV3 struct {
	value *UpdateMTOServiceItemModelTypeV3
	isSet bool
}

func (v NullableUpdateMTOServiceItemModelTypeV3) Get() *UpdateMTOServiceItemModelTypeV3 {
	return v.value
}

func (v *NullableUpdateMTOServiceItemModelTypeV3) Set(val *UpdateMTOServiceItemModelTypeV3) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMTOServiceItemModelTypeV3) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMTOServiceItemModelTypeV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMTOServiceItemModelTypeV3(val *UpdateMTOServiceItemModelTypeV3) *NullableUpdateMTOServiceItemModelTypeV3 {
	return &NullableUpdateMTOServiceItemModelTypeV3{value: val, isSet: true}
}

func (v NullableUpdateMTOServiceItemModelTypeV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMTOServiceItemModelTypeV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

