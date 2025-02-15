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

// checks if the MTOAgent type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MTOAgent{}

// MTOAgent struct for MTOAgent
type MTOAgent struct {
	// The ID of the agent.
	Id *string `json:"id,omitempty"`
	// The ID of the shipment this agent is permitted to release/receive.
	MtoShipmentID *string `json:"mtoShipmentID,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	FirstName NullableString `json:"firstName,omitempty"`
	LastName NullableString `json:"lastName,omitempty"`
	Email NullableString `json:"email,omitempty" validate:"regexp=^([a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\\\.[a-zA-Z]{2,})?$"`
	Phone NullableString `json:"phone,omitempty" validate:"regexp=^([2-9]\\\\d{2}-\\\\d{3}-\\\\d{4})?$"`
	AgentType *MTOAgentType `json:"agentType,omitempty"`
	ETag *string `json:"eTag,omitempty"`
}

// NewMTOAgent instantiates a new MTOAgent object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMTOAgent() *MTOAgent {
	this := MTOAgent{}
	return &this
}

// NewMTOAgentWithDefaults instantiates a new MTOAgent object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMTOAgentWithDefaults() *MTOAgent {
	this := MTOAgent{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MTOAgent) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOAgent) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MTOAgent) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MTOAgent) SetId(v string) {
	o.Id = &v
}

// GetMtoShipmentID returns the MtoShipmentID field value if set, zero value otherwise.
func (o *MTOAgent) GetMtoShipmentID() string {
	if o == nil || IsNil(o.MtoShipmentID) {
		var ret string
		return ret
	}
	return *o.MtoShipmentID
}

// GetMtoShipmentIDOk returns a tuple with the MtoShipmentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOAgent) GetMtoShipmentIDOk() (*string, bool) {
	if o == nil || IsNil(o.MtoShipmentID) {
		return nil, false
	}
	return o.MtoShipmentID, true
}

// HasMtoShipmentID returns a boolean if a field has been set.
func (o *MTOAgent) HasMtoShipmentID() bool {
	if o != nil && !IsNil(o.MtoShipmentID) {
		return true
	}

	return false
}

// SetMtoShipmentID gets a reference to the given string and assigns it to the MtoShipmentID field.
func (o *MTOAgent) SetMtoShipmentID(v string) {
	o.MtoShipmentID = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *MTOAgent) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOAgent) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *MTOAgent) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *MTOAgent) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *MTOAgent) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOAgent) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *MTOAgent) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *MTOAgent) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MTOAgent) GetFirstName() string {
	if o == nil || IsNil(o.FirstName.Get()) {
		var ret string
		return ret
	}
	return *o.FirstName.Get()
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MTOAgent) GetFirstNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FirstName.Get(), o.FirstName.IsSet()
}

// HasFirstName returns a boolean if a field has been set.
func (o *MTOAgent) HasFirstName() bool {
	if o != nil && o.FirstName.IsSet() {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given NullableString and assigns it to the FirstName field.
func (o *MTOAgent) SetFirstName(v string) {
	o.FirstName.Set(&v)
}
// SetFirstNameNil sets the value for FirstName to be an explicit nil
func (o *MTOAgent) SetFirstNameNil() {
	o.FirstName.Set(nil)
}

// UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
func (o *MTOAgent) UnsetFirstName() {
	o.FirstName.Unset()
}

// GetLastName returns the LastName field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MTOAgent) GetLastName() string {
	if o == nil || IsNil(o.LastName.Get()) {
		var ret string
		return ret
	}
	return *o.LastName.Get()
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MTOAgent) GetLastNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastName.Get(), o.LastName.IsSet()
}

// HasLastName returns a boolean if a field has been set.
func (o *MTOAgent) HasLastName() bool {
	if o != nil && o.LastName.IsSet() {
		return true
	}

	return false
}

// SetLastName gets a reference to the given NullableString and assigns it to the LastName field.
func (o *MTOAgent) SetLastName(v string) {
	o.LastName.Set(&v)
}
// SetLastNameNil sets the value for LastName to be an explicit nil
func (o *MTOAgent) SetLastNameNil() {
	o.LastName.Set(nil)
}

// UnsetLastName ensures that no value is present for LastName, not even an explicit nil
func (o *MTOAgent) UnsetLastName() {
	o.LastName.Unset()
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MTOAgent) GetEmail() string {
	if o == nil || IsNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MTOAgent) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *MTOAgent) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *MTOAgent) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *MTOAgent) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *MTOAgent) UnsetEmail() {
	o.Email.Unset()
}

// GetPhone returns the Phone field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MTOAgent) GetPhone() string {
	if o == nil || IsNil(o.Phone.Get()) {
		var ret string
		return ret
	}
	return *o.Phone.Get()
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MTOAgent) GetPhoneOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Phone.Get(), o.Phone.IsSet()
}

// HasPhone returns a boolean if a field has been set.
func (o *MTOAgent) HasPhone() bool {
	if o != nil && o.Phone.IsSet() {
		return true
	}

	return false
}

// SetPhone gets a reference to the given NullableString and assigns it to the Phone field.
func (o *MTOAgent) SetPhone(v string) {
	o.Phone.Set(&v)
}
// SetPhoneNil sets the value for Phone to be an explicit nil
func (o *MTOAgent) SetPhoneNil() {
	o.Phone.Set(nil)
}

// UnsetPhone ensures that no value is present for Phone, not even an explicit nil
func (o *MTOAgent) UnsetPhone() {
	o.Phone.Unset()
}

// GetAgentType returns the AgentType field value if set, zero value otherwise.
func (o *MTOAgent) GetAgentType() MTOAgentType {
	if o == nil || IsNil(o.AgentType) {
		var ret MTOAgentType
		return ret
	}
	return *o.AgentType
}

// GetAgentTypeOk returns a tuple with the AgentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOAgent) GetAgentTypeOk() (*MTOAgentType, bool) {
	if o == nil || IsNil(o.AgentType) {
		return nil, false
	}
	return o.AgentType, true
}

// HasAgentType returns a boolean if a field has been set.
func (o *MTOAgent) HasAgentType() bool {
	if o != nil && !IsNil(o.AgentType) {
		return true
	}

	return false
}

// SetAgentType gets a reference to the given MTOAgentType and assigns it to the AgentType field.
func (o *MTOAgent) SetAgentType(v MTOAgentType) {
	o.AgentType = &v
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *MTOAgent) GetETag() string {
	if o == nil || IsNil(o.ETag) {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOAgent) GetETagOk() (*string, bool) {
	if o == nil || IsNil(o.ETag) {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *MTOAgent) HasETag() bool {
	if o != nil && !IsNil(o.ETag) {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *MTOAgent) SetETag(v string) {
	o.ETag = &v
}

func (o MTOAgent) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MTOAgent) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.MtoShipmentID) {
		toSerialize["mtoShipmentID"] = o.MtoShipmentID
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if o.FirstName.IsSet() {
		toSerialize["firstName"] = o.FirstName.Get()
	}
	if o.LastName.IsSet() {
		toSerialize["lastName"] = o.LastName.Get()
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Phone.IsSet() {
		toSerialize["phone"] = o.Phone.Get()
	}
	if !IsNil(o.AgentType) {
		toSerialize["agentType"] = o.AgentType
	}
	if !IsNil(o.ETag) {
		toSerialize["eTag"] = o.ETag
	}
	return toSerialize, nil
}

type NullableMTOAgent struct {
	value *MTOAgent
	isSet bool
}

func (v NullableMTOAgent) Get() *MTOAgent {
	return v.value
}

func (v *NullableMTOAgent) Set(val *MTOAgent) {
	v.value = val
	v.isSet = true
}

func (v NullableMTOAgent) IsSet() bool {
	return v.isSet
}

func (v *NullableMTOAgent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMTOAgent(val *MTOAgent) *NullableMTOAgent {
	return &NullableMTOAgent{value: val, isSet: true}
}

func (v NullableMTOAgent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMTOAgent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


