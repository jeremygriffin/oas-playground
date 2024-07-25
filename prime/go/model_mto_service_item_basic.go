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

// checks if the MTOServiceItemBasic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MTOServiceItemBasic{}

// MTOServiceItemBasic Describes a basic service item subtype of a MTOServiceItem.
type MTOServiceItemBasic struct {
	MTOServiceItem
	ReServiceCode ReServiceCode `json:"reServiceCode"`
}

type _MTOServiceItemBasic MTOServiceItemBasic

// NewMTOServiceItemBasic instantiates a new MTOServiceItemBasic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMTOServiceItemBasic(reServiceCode ReServiceCode, moveTaskOrderID string, modelType MTOServiceItemModelType) *MTOServiceItemBasic {
	this := MTOServiceItemBasic{}
	this.MoveTaskOrderID = moveTaskOrderID
	this.ModelType = modelType
	this.ReServiceCode = reServiceCode
	return &this
}

// NewMTOServiceItemBasicWithDefaults instantiates a new MTOServiceItemBasic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMTOServiceItemBasicWithDefaults() *MTOServiceItemBasic {
	this := MTOServiceItemBasic{}
	return &this
}

// GetReServiceCode returns the ReServiceCode field value
func (o *MTOServiceItemBasic) GetReServiceCode() ReServiceCode {
	if o == nil {
		var ret ReServiceCode
		return ret
	}

	return o.ReServiceCode
}

// GetReServiceCodeOk returns a tuple with the ReServiceCode field value
// and a boolean to check if the value has been set.
func (o *MTOServiceItemBasic) GetReServiceCodeOk() (*ReServiceCode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReServiceCode, true
}

// SetReServiceCode sets field value
func (o *MTOServiceItemBasic) SetReServiceCode(v ReServiceCode) {
	o.ReServiceCode = v
}

func (o MTOServiceItemBasic) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MTOServiceItemBasic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMTOServiceItem, errMTOServiceItem := json.Marshal(o.MTOServiceItem)
	if errMTOServiceItem != nil {
		return map[string]interface{}{}, errMTOServiceItem
	}
	errMTOServiceItem = json.Unmarshal([]byte(serializedMTOServiceItem), &toSerialize)
	if errMTOServiceItem != nil {
		return map[string]interface{}{}, errMTOServiceItem
	}
	toSerialize["reServiceCode"] = o.ReServiceCode
	return toSerialize, nil
}

func (o *MTOServiceItemBasic) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reServiceCode",
		"moveTaskOrderID",
		"modelType",
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

	varMTOServiceItemBasic := _MTOServiceItemBasic{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMTOServiceItemBasic)

	if err != nil {
		return err
	}

	*o = MTOServiceItemBasic(varMTOServiceItemBasic)

	return err
}

type NullableMTOServiceItemBasic struct {
	value *MTOServiceItemBasic
	isSet bool
}

func (v NullableMTOServiceItemBasic) Get() *MTOServiceItemBasic {
	return v.value
}

func (v *NullableMTOServiceItemBasic) Set(val *MTOServiceItemBasic) {
	v.value = val
	v.isSet = true
}

func (v NullableMTOServiceItemBasic) IsSet() bool {
	return v.isSet
}

func (v *NullableMTOServiceItemBasic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMTOServiceItemBasic(val *MTOServiceItemBasic) *NullableMTOServiceItemBasic {
	return &NullableMTOServiceItemBasic{value: val, isSet: true}
}

func (v NullableMTOServiceItemBasic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMTOServiceItemBasic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


