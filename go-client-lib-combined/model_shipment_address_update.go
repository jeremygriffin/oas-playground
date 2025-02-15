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

// checks if the ShipmentAddressUpdate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipmentAddressUpdate{}

// ShipmentAddressUpdate This represents a destination address change request made by the Prime that is either auto-approved or requires review if the pricing criteria has changed. If criteria has changed, then it must be approved or rejected by a TOO. 
type ShipmentAddressUpdate struct {
	Id string `json:"id"`
	// The reason there is an address change.
	ContractorRemarks string `json:"contractorRemarks"`
	// The TOO comment on approval or rejection.
	OfficeRemarks NullableString `json:"officeRemarks,omitempty"`
	Status ShipmentAddressUpdateStatus `json:"status"`
	ShipmentID string `json:"shipmentID"`
	OriginalAddress Address `json:"originalAddress"`
	NewAddress Address `json:"newAddress"`
	SitOriginalAddress *Address `json:"sitOriginalAddress,omitempty"`
	// The distance between the original SIT address and the previous/old destination address of shipment
	OldSitDistanceBetween *int32 `json:"oldSitDistanceBetween,omitempty"`
	// The distance between the original SIT address and requested new destination address of shipment
	NewSitDistanceBetween *int32 `json:"newSitDistanceBetween,omitempty"`
}

type _ShipmentAddressUpdate ShipmentAddressUpdate

// NewShipmentAddressUpdate instantiates a new ShipmentAddressUpdate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipmentAddressUpdate(id string, contractorRemarks string, status ShipmentAddressUpdateStatus, shipmentID string, originalAddress Address, newAddress Address) *ShipmentAddressUpdate {
	this := ShipmentAddressUpdate{}
	this.Id = id
	this.ContractorRemarks = contractorRemarks
	this.Status = status
	this.ShipmentID = shipmentID
	this.OriginalAddress = originalAddress
	this.NewAddress = newAddress
	return &this
}

// NewShipmentAddressUpdateWithDefaults instantiates a new ShipmentAddressUpdate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipmentAddressUpdateWithDefaults() *ShipmentAddressUpdate {
	this := ShipmentAddressUpdate{}
	return &this
}

// GetId returns the Id field value
func (o *ShipmentAddressUpdate) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ShipmentAddressUpdate) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ShipmentAddressUpdate) SetId(v string) {
	o.Id = v
}

// GetContractorRemarks returns the ContractorRemarks field value
func (o *ShipmentAddressUpdate) GetContractorRemarks() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ContractorRemarks
}

// GetContractorRemarksOk returns a tuple with the ContractorRemarks field value
// and a boolean to check if the value has been set.
func (o *ShipmentAddressUpdate) GetContractorRemarksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ContractorRemarks, true
}

// SetContractorRemarks sets field value
func (o *ShipmentAddressUpdate) SetContractorRemarks(v string) {
	o.ContractorRemarks = v
}

// GetOfficeRemarks returns the OfficeRemarks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ShipmentAddressUpdate) GetOfficeRemarks() string {
	if o == nil || IsNil(o.OfficeRemarks.Get()) {
		var ret string
		return ret
	}
	return *o.OfficeRemarks.Get()
}

// GetOfficeRemarksOk returns a tuple with the OfficeRemarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ShipmentAddressUpdate) GetOfficeRemarksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.OfficeRemarks.Get(), o.OfficeRemarks.IsSet()
}

// HasOfficeRemarks returns a boolean if a field has been set.
func (o *ShipmentAddressUpdate) HasOfficeRemarks() bool {
	if o != nil && o.OfficeRemarks.IsSet() {
		return true
	}

	return false
}

// SetOfficeRemarks gets a reference to the given NullableString and assigns it to the OfficeRemarks field.
func (o *ShipmentAddressUpdate) SetOfficeRemarks(v string) {
	o.OfficeRemarks.Set(&v)
}
// SetOfficeRemarksNil sets the value for OfficeRemarks to be an explicit nil
func (o *ShipmentAddressUpdate) SetOfficeRemarksNil() {
	o.OfficeRemarks.Set(nil)
}

// UnsetOfficeRemarks ensures that no value is present for OfficeRemarks, not even an explicit nil
func (o *ShipmentAddressUpdate) UnsetOfficeRemarks() {
	o.OfficeRemarks.Unset()
}

// GetStatus returns the Status field value
func (o *ShipmentAddressUpdate) GetStatus() ShipmentAddressUpdateStatus {
	if o == nil {
		var ret ShipmentAddressUpdateStatus
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *ShipmentAddressUpdate) GetStatusOk() (*ShipmentAddressUpdateStatus, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *ShipmentAddressUpdate) SetStatus(v ShipmentAddressUpdateStatus) {
	o.Status = v
}

// GetShipmentID returns the ShipmentID field value
func (o *ShipmentAddressUpdate) GetShipmentID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ShipmentID
}

// GetShipmentIDOk returns a tuple with the ShipmentID field value
// and a boolean to check if the value has been set.
func (o *ShipmentAddressUpdate) GetShipmentIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentID, true
}

// SetShipmentID sets field value
func (o *ShipmentAddressUpdate) SetShipmentID(v string) {
	o.ShipmentID = v
}

// GetOriginalAddress returns the OriginalAddress field value
func (o *ShipmentAddressUpdate) GetOriginalAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.OriginalAddress
}

// GetOriginalAddressOk returns a tuple with the OriginalAddress field value
// and a boolean to check if the value has been set.
func (o *ShipmentAddressUpdate) GetOriginalAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginalAddress, true
}

// SetOriginalAddress sets field value
func (o *ShipmentAddressUpdate) SetOriginalAddress(v Address) {
	o.OriginalAddress = v
}

// GetNewAddress returns the NewAddress field value
func (o *ShipmentAddressUpdate) GetNewAddress() Address {
	if o == nil {
		var ret Address
		return ret
	}

	return o.NewAddress
}

// GetNewAddressOk returns a tuple with the NewAddress field value
// and a boolean to check if the value has been set.
func (o *ShipmentAddressUpdate) GetNewAddressOk() (*Address, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NewAddress, true
}

// SetNewAddress sets field value
func (o *ShipmentAddressUpdate) SetNewAddress(v Address) {
	o.NewAddress = v
}

// GetSitOriginalAddress returns the SitOriginalAddress field value if set, zero value otherwise.
func (o *ShipmentAddressUpdate) GetSitOriginalAddress() Address {
	if o == nil || IsNil(o.SitOriginalAddress) {
		var ret Address
		return ret
	}
	return *o.SitOriginalAddress
}

// GetSitOriginalAddressOk returns a tuple with the SitOriginalAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentAddressUpdate) GetSitOriginalAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.SitOriginalAddress) {
		return nil, false
	}
	return o.SitOriginalAddress, true
}

// HasSitOriginalAddress returns a boolean if a field has been set.
func (o *ShipmentAddressUpdate) HasSitOriginalAddress() bool {
	if o != nil && !IsNil(o.SitOriginalAddress) {
		return true
	}

	return false
}

// SetSitOriginalAddress gets a reference to the given Address and assigns it to the SitOriginalAddress field.
func (o *ShipmentAddressUpdate) SetSitOriginalAddress(v Address) {
	o.SitOriginalAddress = &v
}

// GetOldSitDistanceBetween returns the OldSitDistanceBetween field value if set, zero value otherwise.
func (o *ShipmentAddressUpdate) GetOldSitDistanceBetween() int32 {
	if o == nil || IsNil(o.OldSitDistanceBetween) {
		var ret int32
		return ret
	}
	return *o.OldSitDistanceBetween
}

// GetOldSitDistanceBetweenOk returns a tuple with the OldSitDistanceBetween field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentAddressUpdate) GetOldSitDistanceBetweenOk() (*int32, bool) {
	if o == nil || IsNil(o.OldSitDistanceBetween) {
		return nil, false
	}
	return o.OldSitDistanceBetween, true
}

// HasOldSitDistanceBetween returns a boolean if a field has been set.
func (o *ShipmentAddressUpdate) HasOldSitDistanceBetween() bool {
	if o != nil && !IsNil(o.OldSitDistanceBetween) {
		return true
	}

	return false
}

// SetOldSitDistanceBetween gets a reference to the given int32 and assigns it to the OldSitDistanceBetween field.
func (o *ShipmentAddressUpdate) SetOldSitDistanceBetween(v int32) {
	o.OldSitDistanceBetween = &v
}

// GetNewSitDistanceBetween returns the NewSitDistanceBetween field value if set, zero value otherwise.
func (o *ShipmentAddressUpdate) GetNewSitDistanceBetween() int32 {
	if o == nil || IsNil(o.NewSitDistanceBetween) {
		var ret int32
		return ret
	}
	return *o.NewSitDistanceBetween
}

// GetNewSitDistanceBetweenOk returns a tuple with the NewSitDistanceBetween field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ShipmentAddressUpdate) GetNewSitDistanceBetweenOk() (*int32, bool) {
	if o == nil || IsNil(o.NewSitDistanceBetween) {
		return nil, false
	}
	return o.NewSitDistanceBetween, true
}

// HasNewSitDistanceBetween returns a boolean if a field has been set.
func (o *ShipmentAddressUpdate) HasNewSitDistanceBetween() bool {
	if o != nil && !IsNil(o.NewSitDistanceBetween) {
		return true
	}

	return false
}

// SetNewSitDistanceBetween gets a reference to the given int32 and assigns it to the NewSitDistanceBetween field.
func (o *ShipmentAddressUpdate) SetNewSitDistanceBetween(v int32) {
	o.NewSitDistanceBetween = &v
}

func (o ShipmentAddressUpdate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipmentAddressUpdate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["contractorRemarks"] = o.ContractorRemarks
	if o.OfficeRemarks.IsSet() {
		toSerialize["officeRemarks"] = o.OfficeRemarks.Get()
	}
	toSerialize["status"] = o.Status
	toSerialize["shipmentID"] = o.ShipmentID
	toSerialize["originalAddress"] = o.OriginalAddress
	toSerialize["newAddress"] = o.NewAddress
	if !IsNil(o.SitOriginalAddress) {
		toSerialize["sitOriginalAddress"] = o.SitOriginalAddress
	}
	if !IsNil(o.OldSitDistanceBetween) {
		toSerialize["oldSitDistanceBetween"] = o.OldSitDistanceBetween
	}
	if !IsNil(o.NewSitDistanceBetween) {
		toSerialize["newSitDistanceBetween"] = o.NewSitDistanceBetween
	}
	return toSerialize, nil
}

func (o *ShipmentAddressUpdate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"contractorRemarks",
		"status",
		"shipmentID",
		"originalAddress",
		"newAddress",
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

	varShipmentAddressUpdate := _ShipmentAddressUpdate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipmentAddressUpdate)

	if err != nil {
		return err
	}

	*o = ShipmentAddressUpdate(varShipmentAddressUpdate)

	return err
}

type NullableShipmentAddressUpdate struct {
	value *ShipmentAddressUpdate
	isSet bool
}

func (v NullableShipmentAddressUpdate) Get() *ShipmentAddressUpdate {
	return v.value
}

func (v *NullableShipmentAddressUpdate) Set(val *ShipmentAddressUpdate) {
	v.value = val
	v.isSet = true
}

func (v NullableShipmentAddressUpdate) IsSet() bool {
	return v.isSet
}

func (v *NullableShipmentAddressUpdate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipmentAddressUpdate(val *ShipmentAddressUpdate) *NullableShipmentAddressUpdate {
	return &NullableShipmentAddressUpdate{value: val, isSet: true}
}

func (v NullableShipmentAddressUpdate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipmentAddressUpdate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


