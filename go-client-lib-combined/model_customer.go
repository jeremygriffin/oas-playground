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

// checks if the Customer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Customer{}

// Customer struct for Customer
type Customer struct {
	Id *string `json:"id,omitempty"`
	DodID *string `json:"dodID,omitempty"`
	Emplid *string `json:"emplid,omitempty"`
	UserID *string `json:"userID,omitempty"`
	CurrentAddress *Address `json:"currentAddress,omitempty"`
	FirstName *string `json:"firstName,omitempty"`
	LastName *string `json:"lastName,omitempty"`
	Branch *string `json:"branch,omitempty"`
	Phone *string `json:"phone,omitempty"`
	Email *string `json:"email,omitempty" validate:"regexp=^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\\\.[a-zA-Z]{2,}$"`
	ETag *string `json:"eTag,omitempty"`
}

// NewCustomer instantiates a new Customer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomer() *Customer {
	this := Customer{}
	return &this
}

// NewCustomerWithDefaults instantiates a new Customer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerWithDefaults() *Customer {
	this := Customer{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Customer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Customer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Customer) SetId(v string) {
	o.Id = &v
}

// GetDodID returns the DodID field value if set, zero value otherwise.
func (o *Customer) GetDodID() string {
	if o == nil || IsNil(o.DodID) {
		var ret string
		return ret
	}
	return *o.DodID
}

// GetDodIDOk returns a tuple with the DodID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetDodIDOk() (*string, bool) {
	if o == nil || IsNil(o.DodID) {
		return nil, false
	}
	return o.DodID, true
}

// HasDodID returns a boolean if a field has been set.
func (o *Customer) HasDodID() bool {
	if o != nil && !IsNil(o.DodID) {
		return true
	}

	return false
}

// SetDodID gets a reference to the given string and assigns it to the DodID field.
func (o *Customer) SetDodID(v string) {
	o.DodID = &v
}

// GetEmplid returns the Emplid field value if set, zero value otherwise.
func (o *Customer) GetEmplid() string {
	if o == nil || IsNil(o.Emplid) {
		var ret string
		return ret
	}
	return *o.Emplid
}

// GetEmplidOk returns a tuple with the Emplid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetEmplidOk() (*string, bool) {
	if o == nil || IsNil(o.Emplid) {
		return nil, false
	}
	return o.Emplid, true
}

// HasEmplid returns a boolean if a field has been set.
func (o *Customer) HasEmplid() bool {
	if o != nil && !IsNil(o.Emplid) {
		return true
	}

	return false
}

// SetEmplid gets a reference to the given string and assigns it to the Emplid field.
func (o *Customer) SetEmplid(v string) {
	o.Emplid = &v
}

// GetUserID returns the UserID field value if set, zero value otherwise.
func (o *Customer) GetUserID() string {
	if o == nil || IsNil(o.UserID) {
		var ret string
		return ret
	}
	return *o.UserID
}

// GetUserIDOk returns a tuple with the UserID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetUserIDOk() (*string, bool) {
	if o == nil || IsNil(o.UserID) {
		return nil, false
	}
	return o.UserID, true
}

// HasUserID returns a boolean if a field has been set.
func (o *Customer) HasUserID() bool {
	if o != nil && !IsNil(o.UserID) {
		return true
	}

	return false
}

// SetUserID gets a reference to the given string and assigns it to the UserID field.
func (o *Customer) SetUserID(v string) {
	o.UserID = &v
}

// GetCurrentAddress returns the CurrentAddress field value if set, zero value otherwise.
func (o *Customer) GetCurrentAddress() Address {
	if o == nil || IsNil(o.CurrentAddress) {
		var ret Address
		return ret
	}
	return *o.CurrentAddress
}

// GetCurrentAddressOk returns a tuple with the CurrentAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetCurrentAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.CurrentAddress) {
		return nil, false
	}
	return o.CurrentAddress, true
}

// HasCurrentAddress returns a boolean if a field has been set.
func (o *Customer) HasCurrentAddress() bool {
	if o != nil && !IsNil(o.CurrentAddress) {
		return true
	}

	return false
}

// SetCurrentAddress gets a reference to the given Address and assigns it to the CurrentAddress field.
func (o *Customer) SetCurrentAddress(v Address) {
	o.CurrentAddress = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *Customer) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// HasFirstName returns a boolean if a field has been set.
func (o *Customer) HasFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *Customer) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *Customer) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// HasLastName returns a boolean if a field has been set.
func (o *Customer) HasLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *Customer) SetLastName(v string) {
	o.LastName = &v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *Customer) GetBranch() string {
	if o == nil || IsNil(o.Branch) {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetBranchOk() (*string, bool) {
	if o == nil || IsNil(o.Branch) {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *Customer) HasBranch() bool {
	if o != nil && !IsNil(o.Branch) {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *Customer) SetBranch(v string) {
	o.Branch = &v
}

// GetPhone returns the Phone field value if set, zero value otherwise.
func (o *Customer) GetPhone() string {
	if o == nil || IsNil(o.Phone) {
		var ret string
		return ret
	}
	return *o.Phone
}

// GetPhoneOk returns a tuple with the Phone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetPhoneOk() (*string, bool) {
	if o == nil || IsNil(o.Phone) {
		return nil, false
	}
	return o.Phone, true
}

// HasPhone returns a boolean if a field has been set.
func (o *Customer) HasPhone() bool {
	if o != nil && !IsNil(o.Phone) {
		return true
	}

	return false
}

// SetPhone gets a reference to the given string and assigns it to the Phone field.
func (o *Customer) SetPhone(v string) {
	o.Phone = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *Customer) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *Customer) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *Customer) SetEmail(v string) {
	o.Email = &v
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *Customer) GetETag() string {
	if o == nil || IsNil(o.ETag) {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Customer) GetETagOk() (*string, bool) {
	if o == nil || IsNil(o.ETag) {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *Customer) HasETag() bool {
	if o != nil && !IsNil(o.ETag) {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *Customer) SetETag(v string) {
	o.ETag = &v
}

func (o Customer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Customer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DodID) {
		toSerialize["dodID"] = o.DodID
	}
	if !IsNil(o.Emplid) {
		toSerialize["emplid"] = o.Emplid
	}
	if !IsNil(o.UserID) {
		toSerialize["userID"] = o.UserID
	}
	if !IsNil(o.CurrentAddress) {
		toSerialize["currentAddress"] = o.CurrentAddress
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Branch) {
		toSerialize["branch"] = o.Branch
	}
	if !IsNil(o.Phone) {
		toSerialize["phone"] = o.Phone
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.ETag) {
		toSerialize["eTag"] = o.ETag
	}
	return toSerialize, nil
}

type NullableCustomer struct {
	value *Customer
	isSet bool
}

func (v NullableCustomer) Get() *Customer {
	return v.value
}

func (v *NullableCustomer) Set(val *Customer) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomer(val *Customer) *NullableCustomer {
	return &NullableCustomer{value: val, isSet: true}
}

func (v NullableCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


