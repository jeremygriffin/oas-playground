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

// checks if the MTOServiceItemDimension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MTOServiceItemDimension{}

// MTOServiceItemDimension The dimensions for either the item or the crate associated with a crating service item.
type MTOServiceItemDimension struct {
	Id *string `json:"id,omitempty"`
	// Length in thousandth inches. 1000 thou = 1 inch.
	Length int32 `json:"length"`
	// Width in thousandth inches. 1000 thou = 1 inch.
	Width int32 `json:"width"`
	// Height in thousandth inches. 1000 thou = 1 inch.
	Height int32 `json:"height"`
}

type _MTOServiceItemDimension MTOServiceItemDimension

// NewMTOServiceItemDimension instantiates a new MTOServiceItemDimension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMTOServiceItemDimension(length int32, width int32, height int32) *MTOServiceItemDimension {
	this := MTOServiceItemDimension{}
	this.Length = length
	this.Width = width
	this.Height = height
	return &this
}

// NewMTOServiceItemDimensionWithDefaults instantiates a new MTOServiceItemDimension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMTOServiceItemDimensionWithDefaults() *MTOServiceItemDimension {
	this := MTOServiceItemDimension{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MTOServiceItemDimension) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOServiceItemDimension) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MTOServiceItemDimension) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MTOServiceItemDimension) SetId(v string) {
	o.Id = &v
}

// GetLength returns the Length field value
func (o *MTOServiceItemDimension) GetLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Length
}

// GetLengthOk returns a tuple with the Length field value
// and a boolean to check if the value has been set.
func (o *MTOServiceItemDimension) GetLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Length, true
}

// SetLength sets field value
func (o *MTOServiceItemDimension) SetLength(v int32) {
	o.Length = v
}

// GetWidth returns the Width field value
func (o *MTOServiceItemDimension) GetWidth() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Width
}

// GetWidthOk returns a tuple with the Width field value
// and a boolean to check if the value has been set.
func (o *MTOServiceItemDimension) GetWidthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Width, true
}

// SetWidth sets field value
func (o *MTOServiceItemDimension) SetWidth(v int32) {
	o.Width = v
}

// GetHeight returns the Height field value
func (o *MTOServiceItemDimension) GetHeight() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Height
}

// GetHeightOk returns a tuple with the Height field value
// and a boolean to check if the value has been set.
func (o *MTOServiceItemDimension) GetHeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Height, true
}

// SetHeight sets field value
func (o *MTOServiceItemDimension) SetHeight(v int32) {
	o.Height = v
}

func (o MTOServiceItemDimension) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MTOServiceItemDimension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["length"] = o.Length
	toSerialize["width"] = o.Width
	toSerialize["height"] = o.Height
	return toSerialize, nil
}

func (o *MTOServiceItemDimension) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"length",
		"width",
		"height",
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

	varMTOServiceItemDimension := _MTOServiceItemDimension{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMTOServiceItemDimension)

	if err != nil {
		return err
	}

	*o = MTOServiceItemDimension(varMTOServiceItemDimension)

	return err
}

type NullableMTOServiceItemDimension struct {
	value *MTOServiceItemDimension
	isSet bool
}

func (v NullableMTOServiceItemDimension) Get() *MTOServiceItemDimension {
	return v.value
}

func (v *NullableMTOServiceItemDimension) Set(val *MTOServiceItemDimension) {
	v.value = val
	v.isSet = true
}

func (v NullableMTOServiceItemDimension) IsSet() bool {
	return v.isSet
}

func (v *NullableMTOServiceItemDimension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMTOServiceItemDimension(val *MTOServiceItemDimension) *NullableMTOServiceItemDimension {
	return &NullableMTOServiceItemDimension{value: val, isSet: true}
}

func (v NullableMTOServiceItemDimension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMTOServiceItemDimension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


