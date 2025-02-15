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

// PaymentServiceItemStatusV2 the model 'PaymentServiceItemStatusV2'
type PaymentServiceItemStatusV2 string

// List of PaymentServiceItemStatus
const (
	REQUESTED PaymentServiceItemStatusV2 = "REQUESTED"
	APPROVED PaymentServiceItemStatusV2 = "APPROVED"
	DENIED PaymentServiceItemStatusV2 = "DENIED"
	SENT_TO_GEX PaymentServiceItemStatusV2 = "SENT_TO_GEX"
	PAID PaymentServiceItemStatusV2 = "PAID"
	EDI_ERROR PaymentServiceItemStatusV2 = "EDI_ERROR"
)

// All allowed values of PaymentServiceItemStatusV2 enum
var AllowedPaymentServiceItemStatusV2EnumValues = []PaymentServiceItemStatusV2{
	"REQUESTED",
	"APPROVED",
	"DENIED",
	"SENT_TO_GEX",
	"PAID",
	"EDI_ERROR",
}

func (v *PaymentServiceItemStatusV2) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentServiceItemStatusV2(value)
	for _, existing := range AllowedPaymentServiceItemStatusV2EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentServiceItemStatusV2", value)
}

// NewPaymentServiceItemStatusV2FromValue returns a pointer to a valid PaymentServiceItemStatusV2
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentServiceItemStatusV2FromValue(v string) (*PaymentServiceItemStatusV2, error) {
	ev := PaymentServiceItemStatusV2(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentServiceItemStatusV2: valid values are %v", v, AllowedPaymentServiceItemStatusV2EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentServiceItemStatusV2) IsValid() bool {
	for _, existing := range AllowedPaymentServiceItemStatusV2EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentServiceItemStatus value
func (v PaymentServiceItemStatusV2) Ptr() *PaymentServiceItemStatusV2 {
	return &v
}

type NullablePaymentServiceItemStatusV2 struct {
	value *PaymentServiceItemStatusV2
	isSet bool
}

func (v NullablePaymentServiceItemStatusV2) Get() *PaymentServiceItemStatusV2 {
	return v.value
}

func (v *NullablePaymentServiceItemStatusV2) Set(val *PaymentServiceItemStatusV2) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentServiceItemStatusV2) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentServiceItemStatusV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentServiceItemStatusV2(val *PaymentServiceItemStatusV2) *NullablePaymentServiceItemStatusV2 {
	return &NullablePaymentServiceItemStatusV2{value: val, isSet: true}
}

func (v NullablePaymentServiceItemStatusV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentServiceItemStatusV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

