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
	"bytes"
	"fmt"
)

// checks if the UpdateShipmentDestinationAddressV3 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateShipmentDestinationAddressV3{}

// UpdateShipmentDestinationAddressV3 UpdateShipmentDestinationAddress contains the fields required for the prime to request an update for the destination address on an MTO Shipment.
type UpdateShipmentDestinationAddressV3 struct {
	NewAddress AddressV3V3 `json:"newAddress"`
	// This is the remark the Prime has entered, which would be the reason there is an address change.
	ContractorRemarks string `json:"contractorRemarks"`
}

type _UpdateShipmentDestinationAddressV3 UpdateShipmentDestinationAddressV3

// NewUpdateShipmentDestinationAddressV3 instantiates a new UpdateShipmentDestinationAddressV3 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateShipmentDestinationAddressV3(newAddress AddressV3V3, contractorRemarks string) *UpdateShipmentDestinationAddressV3 {
	this := UpdateShipmentDestinationAddressV3{}
	this.NewAddress = newAddress
	this.ContractorRemarks = contractorRemarks
	return &this
}

// NewUpdateShipmentDestinationAddressV3WithDefaults instantiates a new UpdateShipmentDestinationAddressV3 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateShipmentDestinationAddressV3WithDefaults() *UpdateShipmentDestinationAddressV3 {
	this := UpdateShipmentDestinationAddressV3{}
	return &this
}

// GetNewAddress returns the NewAddress field value
func (o *UpdateShipmentDestinationAddressV3) GetNewAddress() AddressV3V3 {
	if o == nil {
		var ret AddressV3V3
		return ret
	}

	return o.NewAddress
}

// GetNewAddressOk returns a tuple with the NewAddress field value
// and a boolean to check if the value has been set.
func (o *UpdateShipmentDestinationAddressV3) GetNewAddressOk() (*AddressV3V3, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewAddress, true
}

// SetNewAddress sets field value
func (o *UpdateShipmentDestinationAddressV3) SetNewAddress(v AddressV3V3) {
	o.NewAddress = v
}

// GetContractorRemarks returns the ContractorRemarks field value
func (o *UpdateShipmentDestinationAddressV3) GetContractorRemarks() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractorRemarks
}

// GetContractorRemarksOk returns a tuple with the ContractorRemarks field value
// and a boolean to check if the value has been set.
func (o *UpdateShipmentDestinationAddressV3) GetContractorRemarksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractorRemarks, true
}

// SetContractorRemarks sets field value
func (o *UpdateShipmentDestinationAddressV3) SetContractorRemarks(v string) {
	o.ContractorRemarks = v
}

func (o UpdateShipmentDestinationAddressV3) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateShipmentDestinationAddressV3) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["newAddress"] = o.NewAddress
	toSerialize["contractorRemarks"] = o.ContractorRemarks
	return toSerialize, nil
}

func (o *UpdateShipmentDestinationAddressV3) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"newAddress",
		"contractorRemarks",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateShipmentDestinationAddressV3 := _UpdateShipmentDestinationAddressV3{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateShipmentDestinationAddressV3)

	if err != nil {
		return err
	}

	*o = UpdateShipmentDestinationAddressV3(varUpdateShipmentDestinationAddressV3)

	return err
}

type NullableUpdateShipmentDestinationAddressV3 struct {
	value *UpdateShipmentDestinationAddressV3
	isSet bool
}

func (v NullableUpdateShipmentDestinationAddressV3) Get() *UpdateShipmentDestinationAddressV3 {
	return v.value
}

func (v *NullableUpdateShipmentDestinationAddressV3) Set(val *UpdateShipmentDestinationAddressV3) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateShipmentDestinationAddressV3) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateShipmentDestinationAddressV3) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateShipmentDestinationAddressV3(val *UpdateShipmentDestinationAddressV3) *NullableUpdateShipmentDestinationAddressV3 {
	return &NullableUpdateShipmentDestinationAddressV3{value: val, isSet: true}
}

func (v NullableUpdateShipmentDestinationAddressV3) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateShipmentDestinationAddressV3) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


