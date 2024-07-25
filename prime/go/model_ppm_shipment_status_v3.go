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

// PPMShipmentStatusV3 Status of the PPM Shipment:   * **DRAFT**: The customer has created the PPM shipment but has not yet submitted their move for counseling.   * **SUBMITTED**: The shipment belongs to a move that has been submitted by the customer or has been created by a Service Counselor or Prime Contractor for a submitted move.   * **WAITING_ON_CUSTOMER**: The PPM shipment has been approved and the customer may now provide their actual move closeout information and documentation required to get paid.   * **NEEDS_ADVANCE_APPROVAL**: The shipment was counseled by the Prime Contractor and approved but an advance was requested so will need further financial approval from the government.   * **NEEDS_CLOSEOUT**: The customer has provided their closeout weight tickets, receipts, and expenses and certified it for the Service Counselor to approve, exclude or reject.   * **CLOSEOUT_COMPLETE**: The Service Counselor has reviewed all of the customer's PPM closeout documentation and authorizes the customer can download and submit their finalized SSW packet. 
type PPMShipmentStatusV3 string

// List of PPMShipmentStatus
const (
	DRAFT PPMShipmentStatusV3 = "DRAFT"
	SUBMITTED PPMShipmentStatusV3 = "SUBMITTED"
	WAITING_ON_CUSTOMER PPMShipmentStatusV3 = "WAITING_ON_CUSTOMER"
	NEEDS_ADVANCE_APPROVAL PPMShipmentStatusV3 = "NEEDS_ADVANCE_APPROVAL"
	NEEDS_CLOSEOUT PPMShipmentStatusV3 = "NEEDS_CLOSEOUT"
	CLOSEOUT_COMPLETE PPMShipmentStatusV3 = "CLOSEOUT_COMPLETE"
)

// All allowed values of PPMShipmentStatusV3 enum
var AllowedPPMShipmentStatusV3EnumValues = []PPMShipmentStatusV3{
	"DRAFT",
	"SUBMITTED",
	"WAITING_ON_CUSTOMER",
	"NEEDS_ADVANCE_APPROVAL",
	"NEEDS_CLOSEOUT",
	"CLOSEOUT_COMPLETE",
}

func (v *PPMShipmentStatusV3) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PPMShipmentStatusV3(value)
	for _, existing := range AllowedPPMShipmentStatusV3EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PPMShipmentStatusV3", value)
}

// NewPPMShipmentStatusV3FromValue returns a pointer to a valid PPMShipmentStatusV3
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPPMShipmentStatusV3FromValue(v string) (*PPMShipmentStatusV3, error) {
	ev := PPMShipmentStatusV3(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PPMShipmentStatusV3: valid values are %v", v, AllowedPPMShipmentStatusV3EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PPMShipmentStatusV3) IsValid() bool {
	for _, existing := range AllowedPPMShipmentStatusV3EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PPMShipmentStatus value
func (v PPMShipmentStatusV3) Ptr() *PPMShipmentStatusV3 {
	return &v
}

type NullablePPMShipmentStatusV3 struct {
	value *PPMShipmentStatusV3
	isSet bool
}

func (v NullablePPMShipmentStatusV3) Get() *PPMShipmentStatusV3 {
	return v.value
}

func (v *NullablePPMShipmentStatusV3) Set(val *PPMShipmentStatusV3) {
	v.value = val
	v.isSet = true
}

func (v NullablePPMShipmentStatusV3) IsSet() bool {
	return v.isSet
}

func (v *NullablePPMShipmentStatusV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePPMShipmentStatusV3(val *PPMShipmentStatusV3) *NullablePPMShipmentStatusV3 {
	return &NullablePPMShipmentStatusV3{value: val, isSet: true}
}

func (v NullablePPMShipmentStatusV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePPMShipmentStatusV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

