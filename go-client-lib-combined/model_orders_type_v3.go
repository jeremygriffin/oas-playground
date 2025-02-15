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

// OrdersTypeV3 the model 'OrdersTypeV3'
type OrdersTypeV3 string

// List of OrdersType
const (
	PERMANENT_CHANGE_OF_STATION OrdersTypeV3 = "PERMANENT_CHANGE_OF_STATION"
	LOCAL_MOVE OrdersTypeV3 = "LOCAL_MOVE"
	RETIREMENT OrdersTypeV3 = "RETIREMENT"
	SEPARATION OrdersTypeV3 = "SEPARATION"
	WOUNDED_WARRIOR OrdersTypeV3 = "WOUNDED_WARRIOR"
	BLUEBARK OrdersTypeV3 = "BLUEBARK"
	SAFETY OrdersTypeV3 = "SAFETY"
)

// All allowed values of OrdersTypeV3 enum
var AllowedOrdersTypeV3EnumValues = []OrdersTypeV3{
	"PERMANENT_CHANGE_OF_STATION",
	"LOCAL_MOVE",
	"RETIREMENT",
	"SEPARATION",
	"WOUNDED_WARRIOR",
	"BLUEBARK",
	"SAFETY",
}

func (v *OrdersTypeV3) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OrdersTypeV3(value)
	for _, existing := range AllowedOrdersTypeV3EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OrdersTypeV3", value)
}

// NewOrdersTypeV3FromValue returns a pointer to a valid OrdersTypeV3
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOrdersTypeV3FromValue(v string) (*OrdersTypeV3, error) {
	ev := OrdersTypeV3(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OrdersTypeV3: valid values are %v", v, AllowedOrdersTypeV3EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OrdersTypeV3) IsValid() bool {
	for _, existing := range AllowedOrdersTypeV3EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OrdersType value
func (v OrdersTypeV3) Ptr() *OrdersTypeV3 {
	return &v
}

type NullableOrdersTypeV3 struct {
	value *OrdersTypeV3
	isSet bool
}

func (v NullableOrdersTypeV3) Get() *OrdersTypeV3 {
	return v.value
}

func (v *NullableOrdersTypeV3) Set(val *OrdersTypeV3) {
	v.value = val
	v.isSet = true
}

func (v NullableOrdersTypeV3) IsSet() bool {
	return v.isSet
}

func (v *NullableOrdersTypeV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOrdersTypeV3(val *OrdersTypeV3) *NullableOrdersTypeV3 {
	return &NullableOrdersTypeV3{value: val, isSet: true}
}

func (v NullableOrdersTypeV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOrdersTypeV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

