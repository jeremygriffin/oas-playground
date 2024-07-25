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

// checks if the CreateSITAddressUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSITAddressUpdateRequest{}

// CreateSITAddressUpdateRequest CreateSITAddressUpdateRequest contains the fields required for the prime to create a SIT address update request.
type CreateSITAddressUpdateRequest struct {
	NewAddress *Address `json:"newAddress,omitempty"`
	ContractorRemarks string `json:"contractorRemarks"`
	MtoServiceItemID *string `json:"mtoServiceItemID,omitempty"`
}

type _CreateSITAddressUpdateRequest CreateSITAddressUpdateRequest

// NewCreateSITAddressUpdateRequest instantiates a new CreateSITAddressUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSITAddressUpdateRequest(contractorRemarks string) *CreateSITAddressUpdateRequest {
	this := CreateSITAddressUpdateRequest{}
	this.ContractorRemarks = contractorRemarks
	return &this
}

// NewCreateSITAddressUpdateRequestWithDefaults instantiates a new CreateSITAddressUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSITAddressUpdateRequestWithDefaults() *CreateSITAddressUpdateRequest {
	this := CreateSITAddressUpdateRequest{}
	return &this
}

// GetNewAddress returns the NewAddress field value if set, zero value otherwise.
func (o *CreateSITAddressUpdateRequest) GetNewAddress() Address {
	if o == nil || IsNil(o.NewAddress) {
		var ret Address
		return ret
	}
	return *o.NewAddress
}

// GetNewAddressOk returns a tuple with the NewAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSITAddressUpdateRequest) GetNewAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.NewAddress) {
		return nil, false
	}
	return o.NewAddress, true
}

// HasNewAddress returns a boolean if a field has been set.
func (o *CreateSITAddressUpdateRequest) HasNewAddress() bool {
	if o != nil && !IsNil(o.NewAddress) {
		return true
	}

	return false
}

// SetNewAddress gets a reference to the given Address and assigns it to the NewAddress field.
func (o *CreateSITAddressUpdateRequest) SetNewAddress(v Address) {
	o.NewAddress = &v
}

// GetContractorRemarks returns the ContractorRemarks field value
func (o *CreateSITAddressUpdateRequest) GetContractorRemarks() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractorRemarks
}

// GetContractorRemarksOk returns a tuple with the ContractorRemarks field value
// and a boolean to check if the value has been set.
func (o *CreateSITAddressUpdateRequest) GetContractorRemarksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractorRemarks, true
}

// SetContractorRemarks sets field value
func (o *CreateSITAddressUpdateRequest) SetContractorRemarks(v string) {
	o.ContractorRemarks = v
}

// GetMtoServiceItemID returns the MtoServiceItemID field value if set, zero value otherwise.
func (o *CreateSITAddressUpdateRequest) GetMtoServiceItemID() string {
	if o == nil || IsNil(o.MtoServiceItemID) {
		var ret string
		return ret
	}
	return *o.MtoServiceItemID
}

// GetMtoServiceItemIDOk returns a tuple with the MtoServiceItemID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSITAddressUpdateRequest) GetMtoServiceItemIDOk() (*string, bool) {
	if o == nil || IsNil(o.MtoServiceItemID) {
		return nil, false
	}
	return o.MtoServiceItemID, true
}

// HasMtoServiceItemID returns a boolean if a field has been set.
func (o *CreateSITAddressUpdateRequest) HasMtoServiceItemID() bool {
	if o != nil && !IsNil(o.MtoServiceItemID) {
		return true
	}

	return false
}

// SetMtoServiceItemID gets a reference to the given string and assigns it to the MtoServiceItemID field.
func (o *CreateSITAddressUpdateRequest) SetMtoServiceItemID(v string) {
	o.MtoServiceItemID = &v
}

func (o CreateSITAddressUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSITAddressUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NewAddress) {
		toSerialize["newAddress"] = o.NewAddress
	}
	toSerialize["contractorRemarks"] = o.ContractorRemarks
	if !IsNil(o.MtoServiceItemID) {
		toSerialize["mtoServiceItemID"] = o.MtoServiceItemID
	}
	return toSerialize, nil
}

func (o *CreateSITAddressUpdateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varCreateSITAddressUpdateRequest := _CreateSITAddressUpdateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateSITAddressUpdateRequest)

	if err != nil {
		return err
	}

	*o = CreateSITAddressUpdateRequest(varCreateSITAddressUpdateRequest)

	return err
}

type NullableCreateSITAddressUpdateRequest struct {
	value *CreateSITAddressUpdateRequest
	isSet bool
}

func (v NullableCreateSITAddressUpdateRequest) Get() *CreateSITAddressUpdateRequest {
	return v.value
}

func (v *NullableCreateSITAddressUpdateRequest) Set(val *CreateSITAddressUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSITAddressUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSITAddressUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSITAddressUpdateRequest(val *CreateSITAddressUpdateRequest) *NullableCreateSITAddressUpdateRequest {
	return &NullableCreateSITAddressUpdateRequest{value: val, isSet: true}
}

func (v NullableCreateSITAddressUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSITAddressUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


