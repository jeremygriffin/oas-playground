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
	"bytes"
	"fmt"
)

// checks if the Amendments type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Amendments{}

// Amendments Metadata outlining number of amendments for given order. 
type Amendments struct {
	// The total count of amendments.
	Total int32 `json:"total"`
	// The total count of amendments available since specified time.
	AvailableSince int32 `json:"availableSince"`
}

type _Amendments Amendments

// NewAmendments instantiates a new Amendments object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAmendments(total int32, availableSince int32) *Amendments {
	this := Amendments{}
	this.Total = total
	this.AvailableSince = availableSince
	return &this
}

// NewAmendmentsWithDefaults instantiates a new Amendments object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAmendmentsWithDefaults() *Amendments {
	this := Amendments{}
	return &this
}

// GetTotal returns the Total field value
func (o *Amendments) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *Amendments) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *Amendments) SetTotal(v int32) {
	o.Total = v
}

// GetAvailableSince returns the AvailableSince field value
func (o *Amendments) GetAvailableSince() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.AvailableSince
}

// GetAvailableSinceOk returns a tuple with the AvailableSince field value
// and a boolean to check if the value has been set.
func (o *Amendments) GetAvailableSinceOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvailableSince, true
}

// SetAvailableSince sets field value
func (o *Amendments) SetAvailableSince(v int32) {
	o.AvailableSince = v
}

func (o Amendments) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Amendments) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["total"] = o.Total
	toSerialize["availableSince"] = o.AvailableSince
	return toSerialize, nil
}

func (o *Amendments) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total",
		"availableSince",
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

	varAmendments := _Amendments{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAmendments)

	if err != nil {
		return err
	}

	*o = Amendments(varAmendments)

	return err
}

type NullableAmendments struct {
	value *Amendments
	isSet bool
}

func (v NullableAmendments) Get() *Amendments {
	return v.value
}

func (v *NullableAmendments) Set(val *Amendments) {
	v.value = val
	v.isSet = true
}

func (v NullableAmendments) IsSet() bool {
	return v.isSet
}

func (v *NullableAmendments) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmendments(val *Amendments) *NullableAmendments {
	return &NullableAmendments{value: val, isSet: true}
}

func (v NullableAmendments) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmendments) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


