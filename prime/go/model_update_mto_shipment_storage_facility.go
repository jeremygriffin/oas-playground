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
)

// checks if the UpdateMTOShipmentStorageFacility type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateMTOShipmentStorageFacility{}

// UpdateMTOShipmentStorageFacility struct for UpdateMTOShipmentStorageFacility
type UpdateMTOShipmentStorageFacility struct {
	Id *string `json:"id,omitempty"`
	FacilityName *string `json:"facilityName,omitempty"`
	Address *Address `json:"address,omitempty"`
	LotNumber NullableString `json:"lotNumber,omitempty"`
	Phone NullableString `json:"phone,omitempty" validate:"regexp=^[2-9]\\\\d{2}-\\\\d{3}-\\\\d{4}$"`
	Email NullableString `json:"email,omitempty" validate:"regexp=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\\\.[a-zA-Z]{2,}$"`
	ETag *string `json:"eTag,omitempty"`
}

// NewUpdateMTOShipmentStorageFacility instantiates a new UpdateMTOShipmentStorageFacility object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateMTOShipmentStorageFacility() *UpdateMTOShipmentStorageFacility {
	this := UpdateMTOShipmentStorageFacility{}
	return &this
}

// NewUpdateMTOShipmentStorageFacilityWithDefaults instantiates a new UpdateMTOShipmentStorageFacility object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateMTOShipmentStorageFacilityWithDefaults() *UpdateMTOShipmentStorageFacility {
	this := UpdateMTOShipmentStorageFacility{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UpdateMTOShipmentStorageFacility) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMTOShipmentStorageFacility) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UpdateMTOShipmentStorageFacility) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *UpdateMTOShipmentStorageFacility) SetId(v string) {
	o.Id = &v
}

// GetFacilityName returns the FacilityName field value if set, zero value otherwise.
func (o *UpdateMTOShipmentStorageFacility) GetFacilityName() string {
	if o == nil || IsNil(o.FacilityName) {
		var ret string
		return ret
	}
	return *o.FacilityName
}

// GetFacilityNameOk returns a tuple with the FacilityName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMTOShipmentStorageFacility) GetFacilityNameOk() (*string, bool) {
	if o == nil || IsNil(o.FacilityName) {
		return nil, false
	}
	return o.FacilityName, true
}

// HasFacilityName returns a boolean if a field has been set.
func (o *UpdateMTOShipmentStorageFacility) HasFacilityName() bool {
	if o != nil && !IsNil(o.FacilityName) {
		return true
	}

	return false
}

// SetFacilityName gets a reference to the given string and assigns it to the FacilityName field.
func (o *UpdateMTOShipmentStorageFacility) SetFacilityName(v string) {
	o.FacilityName = &v
}

// GetAddress returns the Address field value if set, zero value otherwise.
func (o *UpdateMTOShipmentStorageFacility) GetAddress() Address {
	if o == nil || IsNil(o.Address) {
		var ret Address
		return ret
	}
	return *o.Address
}

// GetAddressOk returns a tuple with the Address field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMTOShipmentStorageFacility) GetAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.Address) {
		return nil, false
	}
	return o.Address, true
}

// HasAddress returns a boolean if a field has been set.
func (o *UpdateMTOShipmentStorageFacility) HasAddress() bool {
	if o != nil && !IsNil(o.Address) {
		return true
	}

	return false
}

// SetAddress gets a reference to the given Address and assigns it to the Address field.
func (o *UpdateMTOShipmentStorageFacility) SetAddress(v Address) {
	o.Address = &v
}

// GetLotNumber returns the LotNumber field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMTOShipmentStorageFacility) GetLotNumber() string {
	if o == nil || IsNil(o.LotNumber.Get()) {
		var ret string
		return ret
	}
	return *o.LotNumber.Get()
}

// GetLotNumberOk returns a tuple with the LotNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMTOShipmentStorageFacility) GetLotNumberOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LotNumber.Get(), o.LotNumber.IsSet()
}

// HasLotNumber returns a boolean if a field has been set.
func (o *UpdateMTOShipmentStorageFacility) HasLotNumber() bool {
	if o != nil && o.LotNumber.IsSet() {
		return true
	}

	return false
}

// SetLotNumber gets a reference to the given NullableString and assigns it to the LotNumber field.
func (o *UpdateMTOShipmentStorageFacility) SetLotNumber(v string) {
	o.LotNumber.Set(&v)
}
// SetLotNumberNil sets the value for LotNumber to be an explicit nil
func (o *UpdateMTOShipmentStorageFacility) SetLotNumberNil() {
	o.LotNumber.Set(nil)
}

// UnsetLotNumber ensures that no value is present for LotNumber, not even an explicit nil
func (o *UpdateMTOShipmentStorageFacility) UnsetLotNumber() {
	o.LotNumber.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMTOShipmentStorageFacility) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMTOShipmentStorageFacility) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *UpdateMTOShipmentStorageFacility) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *UpdateMTOShipmentStorageFacility) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *UpdateMTOShipmentStorageFacility) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *UpdateMTOShipmentStorageFacility) UnsetPhone() {
	o.Phone.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *UpdateMTOShipmentStorageFacility) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *UpdateMTOShipmentStorageFacility) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *UpdateMTOShipmentStorageFacility) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *UpdateMTOShipmentStorageFacility) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *UpdateMTOShipmentStorageFacility) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *UpdateMTOShipmentStorageFacility) UnsetEmail() {
	o.Email.Unset()
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *UpdateMTOShipmentStorageFacility) GetETag() string {
	if o == nil || IsNil(o.ETag) {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateMTOShipmentStorageFacility) GetETagOk() (*string, bool) {
	if o == nil || IsNil(o.ETag) {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *UpdateMTOShipmentStorageFacility) HasETag() bool {
	if o != nil && !IsNil(o.ETag) {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *UpdateMTOShipmentStorageFacility) SetETag(v string) {
	o.ETag = &v
}

func (o UpdateMTOShipmentStorageFacility) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateMTOShipmentStorageFacility) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.FacilityName) {
		toSerialize["facilityName"] = o.FacilityName
	}
	if !IsNil(o.Address) {
		toSerialize["address"] = o.Address
	}
	if o.LotNumber.IsSet() {
		toSerialize["lotNumber"] = o.LotNumber.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if !IsNil(o.ETag) {
		toSerialize["eTag"] = o.ETag
	}
	return toSerialize, nil
}

type NullableUpdateMTOShipmentStorageFacility struct {
	value *UpdateMTOShipmentStorageFacility
	isSet bool
}

func (v NullableUpdateMTOShipmentStorageFacility) Get() *UpdateMTOShipmentStorageFacility {
	return v.value
}

func (v *NullableUpdateMTOShipmentStorageFacility) Set(val *UpdateMTOShipmentStorageFacility) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateMTOShipmentStorageFacility) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateMTOShipmentStorageFacility) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateMTOShipmentStorageFacility(val *UpdateMTOShipmentStorageFacility) *NullableUpdateMTOShipmentStorageFacility {
	return &NullableUpdateMTOShipmentStorageFacility{value: val, isSet: true}
}

func (v NullableUpdateMTOShipmentStorageFacility) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateMTOShipmentStorageFacility) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


