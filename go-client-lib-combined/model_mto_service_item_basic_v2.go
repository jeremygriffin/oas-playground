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
	"bytes"
	"fmt"
)

// checks if the MTOServiceItemBasicV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MTOServiceItemBasicV2{}

// MTOServiceItemBasicV2 Describes a basic service item subtype of a MTOServiceItem.
type MTOServiceItemBasicV2 struct {
	MTOServiceItemV2
	ReServiceCode ReServiceCodeV2V2 `json:"reServiceCode"`
}

type _MTOServiceItemBasicV2 MTOServiceItemBasicV2

// NewMTOServiceItemBasicV2 instantiates a new MTOServiceItemBasicV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMTOServiceItemBasicV2(reServiceCode ReServiceCodeV2V2, moveTaskOrderID string, modelType MTOServiceItemModelTypeV2V2) *MTOServiceItemBasicV2 {
	this := MTOServiceItemBasicV2{}
	this.MoveTaskOrderID = moveTaskOrderID
	this.ModelType = modelType
	this.ReServiceCode = reServiceCode
	return &this
}

// NewMTOServiceItemBasicV2WithDefaults instantiates a new MTOServiceItemBasicV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMTOServiceItemBasicV2WithDefaults() *MTOServiceItemBasicV2 {
	this := MTOServiceItemBasicV2{}
	return &this
}

// GetReServiceCode returns the ReServiceCode field value
func (o *MTOServiceItemBasicV2) GetReServiceCode() ReServiceCodeV2V2 {
	if o == nil {
		var ret ReServiceCodeV2V2
		return ret
	}

	return o.ReServiceCode
}

// GetReServiceCodeOk returns a tuple with the ReServiceCode field value
// and a boolean to check if the value has been set.
func (o *MTOServiceItemBasicV2) GetReServiceCodeOk() (*ReServiceCodeV2V2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReServiceCode, true
}

// SetReServiceCode sets field value
func (o *MTOServiceItemBasicV2) SetReServiceCode(v ReServiceCodeV2V2) {
	o.ReServiceCode = v
}

func (o MTOServiceItemBasicV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MTOServiceItemBasicV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedMTOServiceItemV2, errMTOServiceItemV2 := json.Marshal(o.MTOServiceItemV2)
	if errMTOServiceItemV2 != nil {
		return map[string]interface{}{}, errMTOServiceItemV2
	}
	errMTOServiceItemV2 = json.Unmarshal([]byte(serializedMTOServiceItemV2), &toSerialize)
	if errMTOServiceItemV2 != nil {
		return map[string]interface{}{}, errMTOServiceItemV2
	}
	toSerialize["reServiceCode"] = o.ReServiceCode
	return toSerialize, nil
}

func (o *MTOServiceItemBasicV2) UnmarshalJSON(data []byte) (err error) {
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

	varMTOServiceItemBasicV2 := _MTOServiceItemBasicV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMTOServiceItemBasicV2)

	if err != nil {
		return err
	}

	*o = MTOServiceItemBasicV2(varMTOServiceItemBasicV2)

	return err
}

type NullableMTOServiceItemBasicV2 struct {
	value *MTOServiceItemBasicV2
	isSet bool
}

func (v NullableMTOServiceItemBasicV2) Get() *MTOServiceItemBasicV2 {
	return v.value
}

func (v *NullableMTOServiceItemBasicV2) Set(val *MTOServiceItemBasicV2) {
	v.value = val
	v.isSet = true
}

func (v NullableMTOServiceItemBasicV2) IsSet() bool {
	return v.isSet
}

func (v *NullableMTOServiceItemBasicV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMTOServiceItemBasicV2(val *MTOServiceItemBasicV2) *NullableMTOServiceItemBasicV2 {
	return &NullableMTOServiceItemBasicV2{value: val, isSet: true}
}

func (v NullableMTOServiceItemBasicV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMTOServiceItemBasicV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


