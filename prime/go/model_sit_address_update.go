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
	"time"
)

// checks if the SitAddressUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SitAddressUpdate{}

// SitAddressUpdate struct for SitAddressUpdate
type SitAddressUpdate struct {
	Id *string `json:"id,omitempty"`
	MtoServiceItemId *string `json:"mtoServiceItemId,omitempty"`
	NewAddressId *string `json:"newAddressId,omitempty"`
	NewAddress *Address `json:"newAddress,omitempty"`
	OldAddressId *string `json:"oldAddressId,omitempty"`
	OldAddress *Address `json:"oldAddress,omitempty"`
	Status *SitAddressUpdateStatus `json:"status,omitempty"`
	Distance *int32 `json:"distance,omitempty"`
	ContractorRemarks NullableString `json:"contractorRemarks,omitempty"`
	OfficeRemarks NullableString `json:"officeRemarks,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// A hash unique to this shipment that should be used as the \"If-Match\" header for any updates.
	ETag *string `json:"eTag,omitempty"`
}

// NewSitAddressUpdate instantiates a new SitAddressUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSitAddressUpdate() *SitAddressUpdate {
	this := SitAddressUpdate{}
	return &this
}

// NewSitAddressUpdateWithDefaults instantiates a new SitAddressUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSitAddressUpdateWithDefaults() *SitAddressUpdate {
	this := SitAddressUpdate{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SitAddressUpdate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SitAddressUpdate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SitAddressUpdate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *SitAddressUpdate) SetId(v string) {
	o.Id = &v
}

// GetMtoServiceItemId returns the MtoServiceItemId field value if set, zero value otherwise.
func (o *SitAddressUpdate) GetMtoServiceItemId() string {
	if o == nil || IsNil(o.MtoServiceItemId) {
		var ret string
		return ret
	}
	return *o.MtoServiceItemId
}

// GetMtoServiceItemIdOk returns a tuple with the MtoServiceItemId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SitAddressUpdate) GetMtoServiceItemIdOk() (*string, bool) {
	if o == nil || IsNil(o.MtoServiceItemId) {
		return nil, false
	}
	return o.MtoServiceItemId, true
}

// HasMtoServiceItemId returns a boolean if a field has been set.
func (o *SitAddressUpdate) HasMtoServiceItemId() bool {
	if o != nil && !IsNil(o.MtoServiceItemId) {
		return true
	}

	return false
}

// SetMtoServiceItemId gets a reference to the given string and assigns it to the MtoServiceItemId field.
func (o *SitAddressUpdate) SetMtoServiceItemId(v string) {
	o.MtoServiceItemId = &v
}

// GetNewAddressId returns the NewAddressId field value if set, zero value otherwise.
func (o *SitAddressUpdate) GetNewAddressId() string {
	if o == nil || IsNil(o.NewAddressId) {
		var ret string
		return ret
	}
	return *o.NewAddressId
}

// GetNewAddressIdOk returns a tuple with the NewAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SitAddressUpdate) GetNewAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.NewAddressId) {
		return nil, false
	}
	return o.NewAddressId, true
}

// HasNewAddressId returns a boolean if a field has been set.
func (o *SitAddressUpdate) HasNewAddressId() bool {
	if o != nil && !IsNil(o.NewAddressId) {
		return true
	}

	return false
}

// SetNewAddressId gets a reference to the given string and assigns it to the NewAddressId field.
func (o *SitAddressUpdate) SetNewAddressId(v string) {
	o.NewAddressId = &v
}

// GetNewAddress returns the NewAddress field value if set, zero value otherwise.
func (o *SitAddressUpdate) GetNewAddress() Address {
	if o == nil || IsNil(o.NewAddress) {
		var ret Address
		return ret
	}
	return *o.NewAddress
}

// GetNewAddressOk returns a tuple with the NewAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SitAddressUpdate) GetNewAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.NewAddress) {
		return nil, false
	}
	return o.NewAddress, true
}

// HasNewAddress returns a boolean if a field has been set.
func (o *SitAddressUpdate) HasNewAddress() bool {
	if o != nil && !IsNil(o.NewAddress) {
		return true
	}

	return false
}

// SetNewAddress gets a reference to the given Address and assigns it to the NewAddress field.
func (o *SitAddressUpdate) SetNewAddress(v Address) {
	o.NewAddress = &v
}

// GetOldAddressId returns the OldAddressId field value if set, zero value otherwise.
func (o *SitAddressUpdate) GetOldAddressId() string {
	if o == nil || IsNil(o.OldAddressId) {
		var ret string
		return ret
	}
	return *o.OldAddressId
}

// GetOldAddressIdOk returns a tuple with the OldAddressId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SitAddressUpdate) GetOldAddressIdOk() (*string, bool) {
	if o == nil || IsNil(o.OldAddressId) {
		return nil, false
	}
	return o.OldAddressId, true
}

// HasOldAddressId returns a boolean if a field has been set.
func (o *SitAddressUpdate) HasOldAddressId() bool {
	if o != nil && !IsNil(o.OldAddressId) {
		return true
	}

	return false
}

// SetOldAddressId gets a reference to the given string and assigns it to the OldAddressId field.
func (o *SitAddressUpdate) SetOldAddressId(v string) {
	o.OldAddressId = &v
}

// GetOldAddress returns the OldAddress field value if set, zero value otherwise.
func (o *SitAddressUpdate) GetOldAddress() Address {
	if o == nil || IsNil(o.OldAddress) {
		var ret Address
		return ret
	}
	return *o.OldAddress
}

// GetOldAddressOk returns a tuple with the OldAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SitAddressUpdate) GetOldAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.OldAddress) {
		return nil, false
	}
	return o.OldAddress, true
}

// HasOldAddress returns a boolean if a field has been set.
func (o *SitAddressUpdate) HasOldAddress() bool {
	if o != nil && !IsNil(o.OldAddress) {
		return true
	}

	return false
}

// SetOldAddress gets a reference to the given Address and assigns it to the OldAddress field.
func (o *SitAddressUpdate) SetOldAddress(v Address) {
	o.OldAddress = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *SitAddressUpdate) GetStatus() SitAddressUpdateStatus {
	if o == nil || IsNil(o.Status) {
		var ret SitAddressUpdateStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SitAddressUpdate) GetStatusOk() (*SitAddressUpdateStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *SitAddressUpdate) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given SitAddressUpdateStatus and assigns it to the Status field.
func (o *SitAddressUpdate) SetStatus(v SitAddressUpdateStatus) {
	o.Status = &v
}

// GetDistance returns the Distance field value if set, zero value otherwise.
func (o *SitAddressUpdate) GetDistance() int32 {
	if o == nil || IsNil(o.Distance) {
		var ret int32
		return ret
	}
	return *o.Distance
}

// GetDistanceOk returns a tuple with the Distance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SitAddressUpdate) GetDistanceOk() (*int32, bool) {
	if o == nil || IsNil(o.Distance) {
		return nil, false
	}
	return o.Distance, true
}

// HasDistance returns a boolean if a field has been set.
func (o *SitAddressUpdate) HasDistance() bool {
	if o != nil && !IsNil(o.Distance) {
		return true
	}

	return false
}

// SetDistance gets a reference to the given int32 and assigns it to the Distance field.
func (o *SitAddressUpdate) SetDistance(v int32) {
	o.Distance = &v
}

// GetContractorRemarks returns the ContractorRemarks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SitAddressUpdate) GetContractorRemarks() string {
	if o == nil || IsNil(o.ContractorRemarks.Get()) {
		var ret string
		return ret
	}
	return *o.ContractorRemarks.Get()
}

// GetContractorRemarksOk returns a tuple with the ContractorRemarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SitAddressUpdate) GetContractorRemarksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ContractorRemarks.Get(), o.ContractorRemarks.IsSet()
}

// HasContractorRemarks returns a boolean if a field has been set.
func (o *SitAddressUpdate) HasContractorRemarks() bool {
	if o != nil && o.ContractorRemarks.IsSet() {
		return true
	}

	return false
}

// SetContractorRemarks gets a reference to the given NullableString and assigns it to the ContractorRemarks field.
func (o *SitAddressUpdate) SetContractorRemarks(v string) {
	o.ContractorRemarks.Set(&v)
}
// SetContractorRemarksNil sets the value for ContractorRemarks to be an explicit nil
func (o *SitAddressUpdate) SetContractorRemarksNil() {
	o.ContractorRemarks.Set(nil)
}

// UnsetContractorRemarks ensures that no value is present for ContractorRemarks, not even an explicit nil
func (o *SitAddressUpdate) UnsetContractorRemarks() {
	o.ContractorRemarks.Unset()
}

// GetOfficeRemarks returns the OfficeRemarks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *SitAddressUpdate) GetOfficeRemarks() string {
	if o == nil || IsNil(o.OfficeRemarks.Get()) {
		var ret string
		return ret
	}
	return *o.OfficeRemarks.Get()
}

// GetOfficeRemarksOk returns a tuple with the OfficeRemarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *SitAddressUpdate) GetOfficeRemarksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OfficeRemarks.Get(), o.OfficeRemarks.IsSet()
}

// HasOfficeRemarks returns a boolean if a field has been set.
func (o *SitAddressUpdate) HasOfficeRemarks() bool {
	if o != nil && o.OfficeRemarks.IsSet() {
		return true
	}

	return false
}

// SetOfficeRemarks gets a reference to the given NullableString and assigns it to the OfficeRemarks field.
func (o *SitAddressUpdate) SetOfficeRemarks(v string) {
	o.OfficeRemarks.Set(&v)
}
// SetOfficeRemarksNil sets the value for OfficeRemarks to be an explicit nil
func (o *SitAddressUpdate) SetOfficeRemarksNil() {
	o.OfficeRemarks.Set(nil)
}

// UnsetOfficeRemarks ensures that no value is present for OfficeRemarks, not even an explicit nil
func (o *SitAddressUpdate) UnsetOfficeRemarks() {
	o.OfficeRemarks.Unset()
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *SitAddressUpdate) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SitAddressUpdate) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *SitAddressUpdate) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *SitAddressUpdate) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *SitAddressUpdate) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SitAddressUpdate) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *SitAddressUpdate) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *SitAddressUpdate) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *SitAddressUpdate) GetETag() string {
	if o == nil || IsNil(o.ETag) {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SitAddressUpdate) GetETagOk() (*string, bool) {
	if o == nil || IsNil(o.ETag) {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *SitAddressUpdate) HasETag() bool {
	if o != nil && !IsNil(o.ETag) {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *SitAddressUpdate) SetETag(v string) {
	o.ETag = &v
}

func (o SitAddressUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SitAddressUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MtoServiceItemId) {
		toSerialize["mtoServiceItemId"] = o.MtoServiceItemId
	}
	if !IsNil(o.NewAddressId) {
		toSerialize["newAddressId"] = o.NewAddressId
	}
	if !IsNil(o.NewAddress) {
		toSerialize["newAddress"] = o.NewAddress
	}
	if !IsNil(o.OldAddressId) {
		toSerialize["oldAddressId"] = o.OldAddressId
	}
	if !IsNil(o.OldAddress) {
		toSerialize["oldAddress"] = o.OldAddress
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Distance) {
		toSerialize["distance"] = o.Distance
	}
	if o.ContractorRemarks.IsSet() {
		toSerialize["contractorRemarks"] = o.ContractorRemarks.Get()
	}
	if o.OfficeRemarks.IsSet() {
		toSerialize["officeRemarks"] = o.OfficeRemarks.Get()
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.ETag) {
		toSerialize["eTag"] = o.ETag
	}
	return toSerialize, nil
}

type NullableSitAddressUpdate struct {
	value *SitAddressUpdate
	isSet bool
}

func (v NullableSitAddressUpdate) Get() *SitAddressUpdate {
	return v.value
}

func (v *NullableSitAddressUpdate) Set(val *SitAddressUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableSitAddressUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableSitAddressUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSitAddressUpdate(val *SitAddressUpdate) *NullableSitAddressUpdate {
	return &NullableSitAddressUpdate{value: val, isSet: true}
}

func (v NullableSitAddressUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSitAddressUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


