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

// ReServiceCodeV3 This is the full list of service items that can be found on a shipment. Not all service items may be requested by the Prime, but may be returned in a response.  Documentation of all the service items will be provided. 
type ReServiceCodeV3 string

// List of ReServiceCode
const (
	CS ReServiceCodeV3 = "CS"
	DBHF ReServiceCodeV3 = "DBHF"
	DBTF ReServiceCodeV3 = "DBTF"
	DCRT ReServiceCodeV3 = "DCRT"
	DDASIT ReServiceCodeV3 = "DDASIT"
	DDDSIT ReServiceCodeV3 = "DDDSIT"
	DDFSIT ReServiceCodeV3 = "DDFSIT"
	DDP ReServiceCodeV3 = "DDP"
	DDSHUT ReServiceCodeV3 = "DDSHUT"
	DLH ReServiceCodeV3 = "DLH"
	DMHF ReServiceCodeV3 = "DMHF"
	DNPK ReServiceCodeV3 = "DNPK"
	DOASIT ReServiceCodeV3 = "DOASIT"
	DOFSIT ReServiceCodeV3 = "DOFSIT"
	DOP ReServiceCodeV3 = "DOP"
	DOPSIT ReServiceCodeV3 = "DOPSIT"
	DOSHUT ReServiceCodeV3 = "DOSHUT"
	DPK ReServiceCodeV3 = "DPK"
	DSH ReServiceCodeV3 = "DSH"
	DUCRT ReServiceCodeV3 = "DUCRT"
	DUPK ReServiceCodeV3 = "DUPK"
	FSC ReServiceCodeV3 = "FSC"
	IBHF ReServiceCodeV3 = "IBHF"
	IBTF ReServiceCodeV3 = "IBTF"
	ICOLH ReServiceCodeV3 = "ICOLH"
	ICOUB ReServiceCodeV3 = "ICOUB"
	ICRT ReServiceCodeV3 = "ICRT"
	ICRTSA ReServiceCodeV3 = "ICRTSA"
	IDASIT ReServiceCodeV3 = "IDASIT"
	IDDSIT ReServiceCodeV3 = "IDDSIT"
	IDFSIT ReServiceCodeV3 = "IDFSIT"
	IDSHUT ReServiceCodeV3 = "IDSHUT"
	IHPK ReServiceCodeV3 = "IHPK"
	IHUPK ReServiceCodeV3 = "IHUPK"
	INPK ReServiceCodeV3 = "INPK"
	IOASIT ReServiceCodeV3 = "IOASIT"
	IOCLH ReServiceCodeV3 = "IOCLH"
	IOCUB ReServiceCodeV3 = "IOCUB"
	IOFSIT ReServiceCodeV3 = "IOFSIT"
	IOOLH ReServiceCodeV3 = "IOOLH"
	IOOUB ReServiceCodeV3 = "IOOUB"
	IOPSIT ReServiceCodeV3 = "IOPSIT"
	IOSHUT ReServiceCodeV3 = "IOSHUT"
	IUBPK ReServiceCodeV3 = "IUBPK"
	IUBUPK ReServiceCodeV3 = "IUBUPK"
	IUCRT ReServiceCodeV3 = "IUCRT"
	MS ReServiceCodeV3 = "MS"
	NSTH ReServiceCodeV3 = "NSTH"
	NSTUB ReServiceCodeV3 = "NSTUB"
)

// All allowed values of ReServiceCodeV3 enum
var AllowedReServiceCodeV3EnumValues = []ReServiceCodeV3{
	"CS",
	"DBHF",
	"DBTF",
	"DCRT",
	"DDASIT",
	"DDDSIT",
	"DDFSIT",
	"DDP",
	"DDSHUT",
	"DLH",
	"DMHF",
	"DNPK",
	"DOASIT",
	"DOFSIT",
	"DOP",
	"DOPSIT",
	"DOSHUT",
	"DPK",
	"DSH",
	"DUCRT",
	"DUPK",
	"FSC",
	"IBHF",
	"IBTF",
	"ICOLH",
	"ICOUB",
	"ICRT",
	"ICRTSA",
	"IDASIT",
	"IDDSIT",
	"IDFSIT",
	"IDSHUT",
	"IHPK",
	"IHUPK",
	"INPK",
	"IOASIT",
	"IOCLH",
	"IOCUB",
	"IOFSIT",
	"IOOLH",
	"IOOUB",
	"IOPSIT",
	"IOSHUT",
	"IUBPK",
	"IUBUPK",
	"IUCRT",
	"MS",
	"NSTH",
	"NSTUB",
}

func (v *ReServiceCodeV3) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReServiceCodeV3(value)
	for _, existing := range AllowedReServiceCodeV3EnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReServiceCodeV3", value)
}

// NewReServiceCodeV3FromValue returns a pointer to a valid ReServiceCodeV3
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReServiceCodeV3FromValue(v string) (*ReServiceCodeV3, error) {
	ev := ReServiceCodeV3(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReServiceCodeV3: valid values are %v", v, AllowedReServiceCodeV3EnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReServiceCodeV3) IsValid() bool {
	for _, existing := range AllowedReServiceCodeV3EnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReServiceCode value
func (v ReServiceCodeV3) Ptr() *ReServiceCodeV3 {
	return &v
}

type NullableReServiceCodeV3 struct {
	value *ReServiceCodeV3
	isSet bool
}

func (v NullableReServiceCodeV3) Get() *ReServiceCodeV3 {
	return v.value
}

func (v *NullableReServiceCodeV3) Set(val *ReServiceCodeV3) {
	v.value = val
	v.isSet = true
}

func (v NullableReServiceCodeV3) IsSet() bool {
	return v.isSet
}

func (v *NullableReServiceCodeV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReServiceCodeV3(val *ReServiceCodeV3) *NullableReServiceCodeV3 {
	return &NullableReServiceCodeV3{value: val, isSet: true}
}

func (v NullableReServiceCodeV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReServiceCodeV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

