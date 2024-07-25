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

// checks if the MTOServiceItemV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MTOServiceItemV2{}

// MTOServiceItemV2 MTOServiceItem describes a base type of a service item. Polymorphic type.
type MTOServiceItemV2 struct {
	// The ID of the service item.
	Id *string `json:"id,omitempty"`
	// The ID of the move for this service item.
	MoveTaskOrderID string `json:"moveTaskOrderID"`
	// The ID of the shipment this service is for, if any. Optional.
	MtoShipmentID *string `json:"mtoShipmentID,omitempty"`
	// The full descriptive name of the service.
	ReServiceName *string `json:"reServiceName,omitempty"`
	Status *MTOServiceItemStatusV2V2 `json:"status,omitempty"`
	// The reason why this service item was rejected by the TOO.
	RejectionReason NullableString `json:"rejectionReason,omitempty"`
	ModelType MTOServiceItemModelTypeV2V2 `json:"modelType"`
	ServiceRequestDocuments []ServiceRequestDocumentV2V2 `json:"serviceRequestDocuments,omitempty"`
	// A hash unique to this service item that should be used as the \"If-Match\" header for any updates.
	ETag *string `json:"eTag,omitempty"`
}

type _MTOServiceItemV2 MTOServiceItemV2

// NewMTOServiceItemV2 instantiates a new MTOServiceItemV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMTOServiceItemV2(moveTaskOrderID string, modelType MTOServiceItemModelTypeV2V2) *MTOServiceItemV2 {
	this := MTOServiceItemV2{}
	this.MoveTaskOrderID = moveTaskOrderID
	this.ModelType = modelType
	return &this
}

// NewMTOServiceItemV2WithDefaults instantiates a new MTOServiceItemV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMTOServiceItemV2WithDefaults() *MTOServiceItemV2 {
	this := MTOServiceItemV2{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *MTOServiceItemV2) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOServiceItemV2) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *MTOServiceItemV2) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *MTOServiceItemV2) SetId(v string) {
	o.Id = &v
}

// GetMoveTaskOrderID returns the MoveTaskOrderID field value
func (o *MTOServiceItemV2) GetMoveTaskOrderID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MoveTaskOrderID
}

// GetMoveTaskOrderIDOk returns a tuple with the MoveTaskOrderID field value
// and a boolean to check if the value has been set.
func (o *MTOServiceItemV2) GetMoveTaskOrderIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MoveTaskOrderID, true
}

// SetMoveTaskOrderID sets field value
func (o *MTOServiceItemV2) SetMoveTaskOrderID(v string) {
	o.MoveTaskOrderID = v
}

// GetMtoShipmentID returns the MtoShipmentID field value if set, zero value otherwise.
func (o *MTOServiceItemV2) GetMtoShipmentID() string {
	if o == nil || IsNil(o.MtoShipmentID) {
		var ret string
		return ret
	}
	return *o.MtoShipmentID
}

// GetMtoShipmentIDOk returns a tuple with the MtoShipmentID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOServiceItemV2) GetMtoShipmentIDOk() (*string, bool) {
	if o == nil || IsNil(o.MtoShipmentID) {
		return nil, false
	}
	return o.MtoShipmentID, true
}

// HasMtoShipmentID returns a boolean if a field has been set.
func (o *MTOServiceItemV2) HasMtoShipmentID() bool {
	if o != nil && !IsNil(o.MtoShipmentID) {
		return true
	}

	return false
}

// SetMtoShipmentID gets a reference to the given string and assigns it to the MtoShipmentID field.
func (o *MTOServiceItemV2) SetMtoShipmentID(v string) {
	o.MtoShipmentID = &v
}

// GetReServiceName returns the ReServiceName field value if set, zero value otherwise.
func (o *MTOServiceItemV2) GetReServiceName() string {
	if o == nil || IsNil(o.ReServiceName) {
		var ret string
		return ret
	}
	return *o.ReServiceName
}

// GetReServiceNameOk returns a tuple with the ReServiceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOServiceItemV2) GetReServiceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReServiceName) {
		return nil, false
	}
	return o.ReServiceName, true
}

// HasReServiceName returns a boolean if a field has been set.
func (o *MTOServiceItemV2) HasReServiceName() bool {
	if o != nil && !IsNil(o.ReServiceName) {
		return true
	}

	return false
}

// SetReServiceName gets a reference to the given string and assigns it to the ReServiceName field.
func (o *MTOServiceItemV2) SetReServiceName(v string) {
	o.ReServiceName = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *MTOServiceItemV2) GetStatus() MTOServiceItemStatusV2V2 {
	if o == nil || IsNil(o.Status) {
		var ret MTOServiceItemStatusV2V2
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOServiceItemV2) GetStatusOk() (*MTOServiceItemStatusV2V2, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *MTOServiceItemV2) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given MTOServiceItemStatusV2V2 and assigns it to the Status field.
func (o *MTOServiceItemV2) SetStatus(v MTOServiceItemStatusV2V2) {
	o.Status = &v
}

// GetRejectionReason returns the RejectionReason field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *MTOServiceItemV2) GetRejectionReason() string {
	if o == nil || IsNil(o.RejectionReason.Get()) {
		var ret string
		return ret
	}
	return *o.RejectionReason.Get()
}

// GetRejectionReasonOk returns a tuple with the RejectionReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *MTOServiceItemV2) GetRejectionReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RejectionReason.Get(), o.RejectionReason.IsSet()
}

// HasRejectionReason returns a boolean if a field has been set.
func (o *MTOServiceItemV2) HasRejectionReason() bool {
	if o != nil && o.RejectionReason.IsSet() {
		return true
	}

	return false
}

// SetRejectionReason gets a reference to the given NullableString and assigns it to the RejectionReason field.
func (o *MTOServiceItemV2) SetRejectionReason(v string) {
	o.RejectionReason.Set(&v)
}
// SetRejectionReasonNil sets the value for RejectionReason to be an explicit nil
func (o *MTOServiceItemV2) SetRejectionReasonNil() {
	o.RejectionReason.Set(nil)
}

// UnsetRejectionReason ensures that no value is present for RejectionReason, not even an explicit nil
func (o *MTOServiceItemV2) UnsetRejectionReason() {
	o.RejectionReason.Unset()
}

// GetModelType returns the ModelType field value
func (o *MTOServiceItemV2) GetModelType() MTOServiceItemModelTypeV2V2 {
	if o == nil {
		var ret MTOServiceItemModelTypeV2V2
		return ret
	}

	return o.ModelType
}

// GetModelTypeOk returns a tuple with the ModelType field value
// and a boolean to check if the value has been set.
func (o *MTOServiceItemV2) GetModelTypeOk() (*MTOServiceItemModelTypeV2V2, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ModelType, true
}

// SetModelType sets field value
func (o *MTOServiceItemV2) SetModelType(v MTOServiceItemModelTypeV2V2) {
	o.ModelType = v
}

// GetServiceRequestDocuments returns the ServiceRequestDocuments field value if set, zero value otherwise.
func (o *MTOServiceItemV2) GetServiceRequestDocuments() []ServiceRequestDocumentV2V2 {
	if o == nil || IsNil(o.ServiceRequestDocuments) {
		var ret []ServiceRequestDocumentV2V2
		return ret
	}
	return o.ServiceRequestDocuments
}

// GetServiceRequestDocumentsOk returns a tuple with the ServiceRequestDocuments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOServiceItemV2) GetServiceRequestDocumentsOk() ([]ServiceRequestDocumentV2V2, bool) {
	if o == nil || IsNil(o.ServiceRequestDocuments) {
		return nil, false
	}
	return o.ServiceRequestDocuments, true
}

// HasServiceRequestDocuments returns a boolean if a field has been set.
func (o *MTOServiceItemV2) HasServiceRequestDocuments() bool {
	if o != nil && !IsNil(o.ServiceRequestDocuments) {
		return true
	}

	return false
}

// SetServiceRequestDocuments gets a reference to the given []ServiceRequestDocumentV2V2 and assigns it to the ServiceRequestDocuments field.
func (o *MTOServiceItemV2) SetServiceRequestDocuments(v []ServiceRequestDocumentV2V2) {
	o.ServiceRequestDocuments = v
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *MTOServiceItemV2) GetETag() string {
	if o == nil || IsNil(o.ETag) {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MTOServiceItemV2) GetETagOk() (*string, bool) {
	if o == nil || IsNil(o.ETag) {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *MTOServiceItemV2) HasETag() bool {
	if o != nil && !IsNil(o.ETag) {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *MTOServiceItemV2) SetETag(v string) {
	o.ETag = &v
}

func (o MTOServiceItemV2) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MTOServiceItemV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["moveTaskOrderID"] = o.MoveTaskOrderID
	if !IsNil(o.MtoShipmentID) {
		toSerialize["mtoShipmentID"] = o.MtoShipmentID
	}
	if !IsNil(o.ReServiceName) {
		toSerialize["reServiceName"] = o.ReServiceName
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.RejectionReason.IsSet() {
		toSerialize["rejectionReason"] = o.RejectionReason.Get()
	}
	toSerialize["modelType"] = o.ModelType
	if !IsNil(o.ServiceRequestDocuments) {
		toSerialize["serviceRequestDocuments"] = o.ServiceRequestDocuments
	}
	if !IsNil(o.ETag) {
		toSerialize["eTag"] = o.ETag
	}
	return toSerialize, nil
}

func (o *MTOServiceItemV2) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
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

	varMTOServiceItemV2 := _MTOServiceItemV2{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varMTOServiceItemV2)

	if err != nil {
		return err
	}

	*o = MTOServiceItemV2(varMTOServiceItemV2)

	return err
}

type NullableMTOServiceItemV2 struct {
	value *MTOServiceItemV2
	isSet bool
}

func (v NullableMTOServiceItemV2) Get() *MTOServiceItemV2 {
	return v.value
}

func (v *NullableMTOServiceItemV2) Set(val *MTOServiceItemV2) {
	v.value = val
	v.isSet = true
}

func (v NullableMTOServiceItemV2) IsSet() bool {
	return v.isSet
}

func (v *NullableMTOServiceItemV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMTOServiceItemV2(val *MTOServiceItemV2) *NullableMTOServiceItemV2 {
	return &NullableMTOServiceItemV2{value: val, isSet: true}
}

func (v NullableMTOServiceItemV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMTOServiceItemV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


