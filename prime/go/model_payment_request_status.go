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

// PaymentRequestStatus the model 'PaymentRequestStatus'
type PaymentRequestStatus string

// List of PaymentRequestStatus
const (
	PENDING PaymentRequestStatus = "PENDING"
	REVIEWED PaymentRequestStatus = "REVIEWED"
	REVIEWED_AND_ALL_SERVICE_ITEMS_REJECTED PaymentRequestStatus = "REVIEWED_AND_ALL_SERVICE_ITEMS_REJECTED"
	SENT_TO_GEX PaymentRequestStatus = "SENT_TO_GEX"
	RECEIVED_BY_GEX PaymentRequestStatus = "RECEIVED_BY_GEX"
	PAID PaymentRequestStatus = "PAID"
	EDI_ERROR PaymentRequestStatus = "EDI_ERROR"
	DEPRECATED PaymentRequestStatus = "DEPRECATED"
)

// All allowed values of PaymentRequestStatus enum
var AllowedPaymentRequestStatusEnumValues = []PaymentRequestStatus{
	"PENDING",
	"REVIEWED",
	"REVIEWED_AND_ALL_SERVICE_ITEMS_REJECTED",
	"SENT_TO_GEX",
	"RECEIVED_BY_GEX",
	"PAID",
	"EDI_ERROR",
	"DEPRECATED",
}

func (v *PaymentRequestStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PaymentRequestStatus(value)
	for _, existing := range AllowedPaymentRequestStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PaymentRequestStatus", value)
}

// NewPaymentRequestStatusFromValue returns a pointer to a valid PaymentRequestStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPaymentRequestStatusFromValue(v string) (*PaymentRequestStatus, error) {
	ev := PaymentRequestStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PaymentRequestStatus: valid values are %v", v, AllowedPaymentRequestStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PaymentRequestStatus) IsValid() bool {
	for _, existing := range AllowedPaymentRequestStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PaymentRequestStatus value
func (v PaymentRequestStatus) Ptr() *PaymentRequestStatus {
	return &v
}

type NullablePaymentRequestStatus struct {
	value *PaymentRequestStatus
	isSet bool
}

func (v NullablePaymentRequestStatus) Get() *PaymentRequestStatus {
	return v.value
}

func (v *NullablePaymentRequestStatus) Set(val *PaymentRequestStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePaymentRequestStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePaymentRequestStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePaymentRequestStatus(val *PaymentRequestStatus) *NullablePaymentRequestStatus {
	return &NullablePaymentRequestStatus{value: val, isSet: true}
}

func (v NullablePaymentRequestStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePaymentRequestStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

