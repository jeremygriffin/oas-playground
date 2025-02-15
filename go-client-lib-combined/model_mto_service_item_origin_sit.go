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

// checks if the MTOServiceItemOriginSIT type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MTOServiceItemOriginSIT{}

// MTOServiceItemOriginSIT Describes a domestic origin SIT service item. Subtype of a MTOServiceItem.
type MTOServiceItemOriginSIT struct {
	MTOServiceItem
	// Service code allowed for this model type.
	ReServiceCode string `json:"reServiceCode"`
	// Explanation of why Prime is picking up SIT item.
	Reason string `json:"reason"`
	SitPostalCode string `json:"sitPostalCode" validate:"regexp=^(\\\\d{5}([\\\\-]\\\\d{4})?)$"`
	// Entry date for the SIT
	SitEntryDate string `json:"sitEntryDate"`
	// Departure date for SIT. This is the end date of the SIT at either origin or destination. This is optional as it can be updated using the UpdateMTOServiceItemSIT modelType at a later date.
	SitDepartureDate NullableString `json:"sitDepartureDate,omitempty"`
	SitHHGActualOrigin *Address `json:"sitHHGActualOrigin,omitempty"`
	SitHHGOriginalOrigin *Address `json:"sitHHGOriginalOrigin,omitempty"`
	RequestApprovalsRequestedStatus *bool `json:"requestApprovalsRequestedStatus,omitempty"`
	// Date when the customer has requested delivery out of SIT.
	SitRequestedDelivery NullableString `json:"sitRequestedDelivery,omitempty"`
	// Date when the customer contacted the prime for a delivery out of SIT.
	SitCustomerContacted NullableString `json:"sitCustomerContacted,omitempty"`
}

type _MTOServiceItemOriginSIT MTOServiceItemOriginSIT

// NewMTOServiceItemOriginSIT instantiates a new MTOServiceItemOriginSIT object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMTOServiceItemOriginSIT(reServiceCode string, reason string, sitPostalCode string, sitEntryDate string, moveTaskOrderID string, modelType MTOServiceItemModelType) *MTOServiceItemOriginSIT {
	this := MTOServiceItemOriginSIT{}
	this.MoveTaskOrderID = moveTaskOrderID
	this.ModelType = modelType
	this.ReServiceCode = reServiceCode
	this.Reason = reason
	this.SitPostalCode = sitPostalCode
	this.SitEntryDate = sitEntryDate
	return &this
}

// NewMTOServiceItemOriginSITWithDefaults instantiates a new MTOServiceItemOriginSIT object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMTOServiceItemOriginSITWithDefaults() *MTOServiceItemOriginSIT {
	this := MTOServiceItemOriginSIT{}
	return &this
}

// GetReServiceCode returns the ReServiceCode field value
func (o *MTOServiceItemOriginSIT) GetReServiceCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReServiceCode
}

// GetReServiceCodeOk returns a tuple with the ReServiceCode field value
// and a boolean to check if the value has been set.
func (o *MTOServiceItemOriginSIT) GetReServiceCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReServiceCode, true
}

// SetReServiceCode sets field value
func (o *MTOServiceItemOriginSIT) SetReServiceCode(v string) {
	o.ReServiceCode = v
}

// GetReason returns the Reason field value
func (o *MTOServiceItemOriginSIT) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *MTOServiceItemOriginSIT) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *MTOServiceItemOriginSIT) SetReason(v string) {
	o.Reason = v
}

// GetSitPostalCode returns the SitPostalCode field value
func (o *MTOServiceItemOriginSIT) GetSitPostalCode() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SitPostalCode
}

// GetSitPostalCodeOk returns a tuple with the SitPostalCode field value
// and a boolean to check if the value has been set.
func (o *MTOServiceItemOriginSIT) GetSitPostalCodeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SitPostalCode, true
}

// SetSitPostalCode sets field value
func (o *MTOServiceItemOriginSIT) SetSitPostalCode(v string) {
	o.SitPostalCode = v
}

// GetSitEntryDate returns the SitEntryDate field value
func (o *MTOServiceItemOriginSIT) GetSitEntryDate() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SitEntryDate
}

// GetSitEntryDateOk returns a tuple with the SitEntryDate field value
// and a boolean to check if the value has been set.
func (o *MTOServiceItemOriginSIT) GetSitEntryDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SitEntryDate, true
}

// SetSitEntryDate sets field value
func (o *MTOServiceItemOriginSIT) SetSitEntryDate(v string) {
	o.SitEntryDate = v
}

// GetSitDepartureDate returns the SitDepartureDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MTOServiceItemOriginSIT) GetSitDepartureDate() string {
	if o == nil || IsNil(o.SitDepartureDate.Get()) {
		var ret string
		return ret
	}
	return *o.SitDepartureDate.Get()
}

// GetSitDepartureDateOk returns a tuple with the SitDepartureDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MTOServiceItemOriginSIT) GetSitDepartureDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SitDepartureDate.Get(), o.SitDepartureDate.IsSet()
}

// HasSitDepartureDate returns a boolean if a field has been set.
func (o *MTOServiceItemOriginSIT) HasSitDepartureDate() bool {
	if o != nil && o.SitDepartureDate.IsSet() {
		return true
	}

	return false
}

// SetSitDepartureDate gets a reference to the given NullableString and assigns it to the SitDepartureDate field.
func (o *MTOServiceItemOriginSIT) SetSitDepartureDate(v string) {
	o.SitDepartureDate.Set(&v)
}
// SetSitDepartureDateNil sets the value for SitDepartureDate to be an explicit nil
func (o *MTOServiceItemOriginSIT) SetSitDepartureDateNil() {
	o.SitDepartureDate.Set(nil)
}

// UnsetSitDepartureDate ensures that no value is present for SitDepartureDate, not even an explicit nil
func (o *MTOServiceItemOriginSIT) UnsetSitDepartureDate() {
	o.SitDepartureDate.Unset()
}

// GetSitHHGActualOrigin returns the SitHHGActualOrigin field value if set, zero value otherwise.
func (o *MTOServiceItemOriginSIT) GetSitHHGActualOrigin() Address {
	if o == nil || IsNil(o.SitHHGActualOrigin) {
		var ret Address
		return ret
	}
	return *o.SitHHGActualOrigin
}

// GetSitHHGActualOriginOk returns a tuple with the SitHHGActualOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOServiceItemOriginSIT) GetSitHHGActualOriginOk() (*Address, bool) {
	if o == nil || IsNil(o.SitHHGActualOrigin) {
		return nil, false
	}
	return o.SitHHGActualOrigin, true
}

// HasSitHHGActualOrigin returns a boolean if a field has been set.
func (o *MTOServiceItemOriginSIT) HasSitHHGActualOrigin() bool {
	if o != nil && !IsNil(o.SitHHGActualOrigin) {
		return true
	}

	return false
}

// SetSitHHGActualOrigin gets a reference to the given Address and assigns it to the SitHHGActualOrigin field.
func (o *MTOServiceItemOriginSIT) SetSitHHGActualOrigin(v Address) {
	o.SitHHGActualOrigin = &v
}

// GetSitHHGOriginalOrigin returns the SitHHGOriginalOrigin field value if set, zero value otherwise.
func (o *MTOServiceItemOriginSIT) GetSitHHGOriginalOrigin() Address {
	if o == nil || IsNil(o.SitHHGOriginalOrigin) {
		var ret Address
		return ret
	}
	return *o.SitHHGOriginalOrigin
}

// GetSitHHGOriginalOriginOk returns a tuple with the SitHHGOriginalOrigin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOServiceItemOriginSIT) GetSitHHGOriginalOriginOk() (*Address, bool) {
	if o == nil || IsNil(o.SitHHGOriginalOrigin) {
		return nil, false
	}
	return o.SitHHGOriginalOrigin, true
}

// HasSitHHGOriginalOrigin returns a boolean if a field has been set.
func (o *MTOServiceItemOriginSIT) HasSitHHGOriginalOrigin() bool {
	if o != nil && !IsNil(o.SitHHGOriginalOrigin) {
		return true
	}

	return false
}

// SetSitHHGOriginalOrigin gets a reference to the given Address and assigns it to the SitHHGOriginalOrigin field.
func (o *MTOServiceItemOriginSIT) SetSitHHGOriginalOrigin(v Address) {
	o.SitHHGOriginalOrigin = &v
}

// GetRequestApprovalsRequestedStatus returns the RequestApprovalsRequestedStatus field value if set, zero value otherwise.
func (o *MTOServiceItemOriginSIT) GetRequestApprovalsRequestedStatus() bool {
	if o == nil || IsNil(o.RequestApprovalsRequestedStatus) {
		var ret bool
		return ret
	}
	return *o.RequestApprovalsRequestedStatus
}

// GetRequestApprovalsRequestedStatusOk returns a tuple with the RequestApprovalsRequestedStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOServiceItemOriginSIT) GetRequestApprovalsRequestedStatusOk() (*bool, bool) {
	if o == nil || IsNil(o.RequestApprovalsRequestedStatus) {
		return nil, false
	}
	return o.RequestApprovalsRequestedStatus, true
}

// HasRequestApprovalsRequestedStatus returns a boolean if a field has been set.
func (o *MTOServiceItemOriginSIT) HasRequestApprovalsRequestedStatus() bool {
	if o != nil && !IsNil(o.RequestApprovalsRequestedStatus) {
		return true
	}

	return false
}

// SetRequestApprovalsRequestedStatus gets a reference to the given bool and assigns it to the RequestApprovalsRequestedStatus field.
func (o *MTOServiceItemOriginSIT) SetRequestApprovalsRequestedStatus(v bool) {
	o.RequestApprovalsRequestedStatus = &v
}

// GetSitRequestedDelivery returns the SitRequestedDelivery field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MTOServiceItemOriginSIT) GetSitRequestedDelivery() string {
	if o == nil || IsNil(o.SitRequestedDelivery.Get()) {
		var ret string
		return ret
	}
	return *o.SitRequestedDelivery.Get()
}

// GetSitRequestedDeliveryOk returns a tuple with the SitRequestedDelivery field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MTOServiceItemOriginSIT) GetSitRequestedDeliveryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SitRequestedDelivery.Get(), o.SitRequestedDelivery.IsSet()
}

// HasSitRequestedDelivery returns a boolean if a field has been set.
func (o *MTOServiceItemOriginSIT) HasSitRequestedDelivery() bool {
	if o != nil && o.SitRequestedDelivery.IsSet() {
		return true
	}

	return false
}

// SetSitRequestedDelivery gets a reference to the given NullableString and assigns it to the SitRequestedDelivery field.
func (o *MTOServiceItemOriginSIT) SetSitRequestedDelivery(v string) {
	o.SitRequestedDelivery.Set(&v)
}
// SetSitRequestedDeliveryNil sets the value for SitRequestedDelivery to be an explicit nil
func (o *MTOServiceItemOriginSIT) SetSitRequestedDeliveryNil() {
	o.SitRequestedDelivery.Set(nil)
}

// UnsetSitRequestedDelivery ensures that no value is present for SitRequestedDelivery, not even an explicit nil
func (o *MTOServiceItemOriginSIT) UnsetSitRequestedDelivery() {
	o.SitRequestedDelivery.Unset()
}

// GetSitCustomerContacted returns the SitCustomerContacted field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MTOServiceItemOriginSIT) GetSitCustomerContacted() string {
	if o == nil || IsNil(o.SitCustomerContacted.Get()) {
		var ret string
		return ret
	}
	return *o.SitCustomerContacted.Get()
}

// GetSitCustomerContactedOk returns a tuple with the SitCustomerContacted field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MTOServiceItemOriginSIT) GetSitCustomerContactedOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SitCustomerContacted.Get(), o.SitCustomerContacted.IsSet()
}

// HasSitCustomerContacted returns a boolean if a field has been set.
func (o *MTOServiceItemOriginSIT) HasSitCustomerContacted() bool {
	if o != nil && o.SitCustomerContacted.IsSet() {
		return true
	}

	return false
}

// SetSitCustomerContacted gets a reference to the given NullableString and assigns it to the SitCustomerContacted field.
func (o *MTOServiceItemOriginSIT) SetSitCustomerContacted(v string) {
	o.SitCustomerContacted.Set(&v)
}
// SetSitCustomerContactedNil sets the value for SitCustomerContacted to be an explicit nil
func (o *MTOServiceItemOriginSIT) SetSitCustomerContactedNil() {
	o.SitCustomerContacted.Set(nil)
}

// UnsetSitCustomerContacted ensures that no value is present for SitCustomerContacted, not even an explicit nil
func (o *MTOServiceItemOriginSIT) UnsetSitCustomerContacted() {
	o.SitCustomerContacted.Unset()
}

func (o MTOServiceItemOriginSIT) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MTOServiceItemOriginSIT) ToMap() (map[string]interface{}, error) {
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
	toSerialize["reason"] = o.Reason
	toSerialize["sitPostalCode"] = o.SitPostalCode
	toSerialize["sitEntryDate"] = o.SitEntryDate
	if o.SitDepartureDate.IsSet() {
		toSerialize["sitDepartureDate"] = o.SitDepartureDate.Get()
	}
	if !IsNil(o.SitHHGActualOrigin) {
		toSerialize["sitHHGActualOrigin"] = o.SitHHGActualOrigin
	}
	if !IsNil(o.SitHHGOriginalOrigin) {
		toSerialize["sitHHGOriginalOrigin"] = o.SitHHGOriginalOrigin
	}
	if !IsNil(o.RequestApprovalsRequestedStatus) {
		toSerialize["requestApprovalsRequestedStatus"] = o.RequestApprovalsRequestedStatus
	}
	if o.SitRequestedDelivery.IsSet() {
		toSerialize["sitRequestedDelivery"] = o.SitRequestedDelivery.Get()
	}
	if o.SitCustomerContacted.IsSet() {
		toSerialize["sitCustomerContacted"] = o.SitCustomerContacted.Get()
	}
	return toSerialize, nil
}

func (o *MTOServiceItemOriginSIT) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"reServiceCode",
		"reason",
		"sitPostalCode",
		"sitEntryDate",
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

	varMTOServiceItemOriginSIT := _MTOServiceItemOriginSIT{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMTOServiceItemOriginSIT)

	if err != nil {
		return err
	}

	*o = MTOServiceItemOriginSIT(varMTOServiceItemOriginSIT)

	return err
}

type NullableMTOServiceItemOriginSIT struct {
	value *MTOServiceItemOriginSIT
	isSet bool
}

func (v NullableMTOServiceItemOriginSIT) Get() *MTOServiceItemOriginSIT {
	return v.value
}

func (v *NullableMTOServiceItemOriginSIT) Set(val *MTOServiceItemOriginSIT) {
	v.value = val
	v.isSet = true
}

func (v NullableMTOServiceItemOriginSIT) IsSet() bool {
	return v.isSet
}

func (v *NullableMTOServiceItemOriginSIT) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMTOServiceItemOriginSIT(val *MTOServiceItemOriginSIT) *NullableMTOServiceItemOriginSIT {
	return &NullableMTOServiceItemOriginSIT{value: val, isSet: true}
}

func (v NullableMTOServiceItemOriginSIT) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMTOServiceItemOriginSIT) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


