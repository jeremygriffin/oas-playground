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

// checks if the UpdateMTOServiceItemShuttleV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMTOServiceItemShuttleV2{}

// UpdateMTOServiceItemShuttleV2 Subtype used to provide the estimated weight and actual weight for shuttle. This is not creating a new service item but rather updating an existing service item. 
type UpdateMTOServiceItemShuttleV2 struct {
	UpdateMTOServiceItemV2
	// Provided by the movers, based on weight tickets. Relevant for shuttling (DDSHUT & DOSHUT) service items.
	ActualWeight NullableInt32 `json:"actualWeight,omitempty"`
	// An estimate of how much weight from a shipment will be included in a shuttling (DDSHUT & DOSHUT) service item.
	EstimatedWeight NullableInt32 `json:"estimatedWeight,omitempty"`
	// Service code allowed for this model type.
	ReServiceCode *string `json:"reServiceCode,omitempty"`
}

type _UpdateMTOServiceItemShuttleV2 UpdateMTOServiceItemShuttleV2

// NewUpdateMTOServiceItemShuttleV2 instantiates a new UpdateMTOServiceItemShuttleV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMTOServiceItemShuttleV2(modelType UpdateMTOServiceItemModelTypeV2V2) *UpdateMTOServiceItemShuttleV2 {
	this := UpdateMTOServiceItemShuttleV2{}
	this.ModelType = modelType
	return &this
}

// NewUpdateMTOServiceItemShuttleV2WithDefaults instantiates a new UpdateMTOServiceItemShuttleV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMTOServiceItemShuttleV2WithDefaults() *UpdateMTOServiceItemShuttleV2 {
	this := UpdateMTOServiceItemShuttleV2{}
	return &this
}

// GetActualWeight returns the ActualWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMTOServiceItemShuttleV2) GetActualWeight() int32 {
	if o == nil || IsNil(o.ActualWeight.Get()) {
		var ret int32
		return ret
	}
	return *o.ActualWeight.Get()
}

// GetActualWeightOk returns a tuple with the ActualWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMTOServiceItemShuttleV2) GetActualWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.ActualWeight.Get(), o.ActualWeight.IsSet()
}

// HasActualWeight returns a boolean if a field has been set.
func (o *UpdateMTOServiceItemShuttleV2) HasActualWeight() bool {
	if o != nil && o.ActualWeight.IsSet() {
		return true
	}

	return false
}

// SetActualWeight gets a reference to the given NullableInt32 and assigns it to the ActualWeight field.
func (o *UpdateMTOServiceItemShuttleV2) SetActualWeight(v int32) {
	o.ActualWeight.Set(&v)
}
// SetActualWeightNil sets the value for ActualWeight to be an explicit nil
func (o *UpdateMTOServiceItemShuttleV2) SetActualWeightNil() {
	o.ActualWeight.Set(nil)
}

// UnsetActualWeight ensures that no value is present for ActualWeight, not even an explicit nil
func (o *UpdateMTOServiceItemShuttleV2) UnsetActualWeight() {
	o.ActualWeight.Unset()
}

// GetEstimatedWeight returns the EstimatedWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMTOServiceItemShuttleV2) GetEstimatedWeight() int32 {
	if o == nil || IsNil(o.EstimatedWeight.Get()) {
		var ret int32
		return ret
	}
	return *o.EstimatedWeight.Get()
}

// GetEstimatedWeightOk returns a tuple with the EstimatedWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMTOServiceItemShuttleV2) GetEstimatedWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.EstimatedWeight.Get(), o.EstimatedWeight.IsSet()
}

// HasEstimatedWeight returns a boolean if a field has been set.
func (o *UpdateMTOServiceItemShuttleV2) HasEstimatedWeight() bool {
	if o != nil && o.EstimatedWeight.IsSet() {
		return true
	}

	return false
}

// SetEstimatedWeight gets a reference to the given NullableInt32 and assigns it to the EstimatedWeight field.
func (o *UpdateMTOServiceItemShuttleV2) SetEstimatedWeight(v int32) {
	o.EstimatedWeight.Set(&v)
}
// SetEstimatedWeightNil sets the value for EstimatedWeight to be an explicit nil
func (o *UpdateMTOServiceItemShuttleV2) SetEstimatedWeightNil() {
	o.EstimatedWeight.Set(nil)
}

// UnsetEstimatedWeight ensures that no value is present for EstimatedWeight, not even an explicit nil
func (o *UpdateMTOServiceItemShuttleV2) UnsetEstimatedWeight() {
	o.EstimatedWeight.Unset()
}

// GetReServiceCode returns the ReServiceCode field value if set, zero value otherwise.
func (o *UpdateMTOServiceItemShuttleV2) GetReServiceCode() string {
	if o == nil || IsNil(o.ReServiceCode) {
		var ret string
		return ret
	}
	return *o.ReServiceCode
}

// GetReServiceCodeOk returns a tuple with the ReServiceCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMTOServiceItemShuttleV2) GetReServiceCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ReServiceCode) {
		return nil, false
	}
	return o.ReServiceCode, true
}

// HasReServiceCode returns a boolean if a field has been set.
func (o *UpdateMTOServiceItemShuttleV2) HasReServiceCode() bool {
	if o != nil && !IsNil(o.ReServiceCode) {
		return true
	}

	return false
}

// SetReServiceCode gets a reference to the given string and assigns it to the ReServiceCode field.
func (o *UpdateMTOServiceItemShuttleV2) SetReServiceCode(v string) {
	o.ReServiceCode = &v
}

func (o UpdateMTOServiceItemShuttleV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMTOServiceItemShuttleV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedUpdateMTOServiceItemV2, errUpdateMTOServiceItemV2 := json.Marshal(o.UpdateMTOServiceItemV2)
	if errUpdateMTOServiceItemV2 != nil {
		return map[string]interface{}{}, errUpdateMTOServiceItemV2
	}
	errUpdateMTOServiceItemV2 = json.Unmarshal([]byte(serializedUpdateMTOServiceItemV2), &toSerialize)
	if errUpdateMTOServiceItemV2 != nil {
		return map[string]interface{}{}, errUpdateMTOServiceItemV2
	}
	if o.ActualWeight.IsSet() {
		toSerialize["actualWeight"] = o.ActualWeight.Get()
	}
	if o.EstimatedWeight.IsSet() {
		toSerialize["estimatedWeight"] = o.EstimatedWeight.Get()
	}
	if !IsNil(o.ReServiceCode) {
		toSerialize["reServiceCode"] = o.ReServiceCode
	}
	return toSerialize, nil
}

func (o *UpdateMTOServiceItemShuttleV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varUpdateMTOServiceItemShuttleV2 := _UpdateMTOServiceItemShuttleV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUpdateMTOServiceItemShuttleV2)

	if err != nil {
		return err
	}

	*o = UpdateMTOServiceItemShuttleV2(varUpdateMTOServiceItemShuttleV2)

	return err
}

type NullableUpdateMTOServiceItemShuttleV2 struct {
	value *UpdateMTOServiceItemShuttleV2
	isSet bool
}

func (v NullableUpdateMTOServiceItemShuttleV2) Get() *UpdateMTOServiceItemShuttleV2 {
	return v.value
}

func (v *NullableUpdateMTOServiceItemShuttleV2) Set(val *UpdateMTOServiceItemShuttleV2) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMTOServiceItemShuttleV2) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMTOServiceItemShuttleV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMTOServiceItemShuttleV2(val *UpdateMTOServiceItemShuttleV2) *NullableUpdateMTOServiceItemShuttleV2 {
	return &NullableUpdateMTOServiceItemShuttleV2{value: val, isSet: true}
}

func (v NullableUpdateMTOServiceItemShuttleV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMTOServiceItemShuttleV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


