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

// checks if the CreateMTOShipment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateMTOShipment{}

// CreateMTOShipment struct for CreateMTOShipment
type CreateMTOShipment struct {
	// The ID of the move this new shipment is for.
	MoveTaskOrderID string `json:"moveTaskOrderID"`
	// The customer's preferred pickup date. Other dates, such as required delivery date and (outside MilMove) the pack date, are derived from this date. 
	RequestedPickupDate NullableString `json:"requestedPickupDate,omitempty"`
	// The estimated weight of this shipment, determined by the movers during the pre-move survey. This value **can only be updated once.** If there was an issue with estimating the weight and a mistake was made, the Prime contractor will need to contact the TOO to change it. 
	PrimeEstimatedWeight NullableInt32 `json:"primeEstimatedWeight,omitempty"`
	// The customer can use the customer remarks field to inform the services counselor and the movers about any special circumstances for this shipment. Typical examples:   * bulky or fragile items,   * weapons,   * access info for their address.  Customer enters this information during onboarding. Optional field. 
	CustomerRemarks NullableString `json:"customerRemarks,omitempty"`
	// A list of the agents for a shipment. Agents are the people who the Prime contractor recognize as permitted to release (in the case of pickup) or receive (on delivery) a shipment. 
	Agents []MTOAgent `json:"agents,omitempty"`
	// A list of service items connected to this shipment.
	MtoServiceItems []MTOServiceItem `json:"mtoServiceItems,omitempty"`
	// The address where the movers should pick up this shipment.
	PickupAddress *Address `json:"pickupAddress,omitempty"`
	// Where the movers should deliver this shipment.
	DestinationAddress *Address `json:"destinationAddress,omitempty"`
	ShipmentType MTOShipmentType `json:"shipmentType"`
	// This value indicates whether or not this shipment is part of a diversion. If yes, the shipment can be either the starting or ending segment of the diversion. 
	Diversion *bool `json:"diversion,omitempty"`
	// Email or ID of the person who will be contacted in the event of questions or concerns about this update. May be the person performing the update, or someone else working with the Prime contractor. 
	PointOfContact *string `json:"pointOfContact,omitempty"`
	CounselorRemarks NullableString `json:"counselorRemarks,omitempty"`
	PpmShipment *CreatePPMShipment `json:"ppmShipment,omitempty"`
}

type _CreateMTOShipment CreateMTOShipment

// NewCreateMTOShipment instantiates a new CreateMTOShipment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateMTOShipment(moveTaskOrderID string, shipmentType MTOShipmentType) *CreateMTOShipment {
	this := CreateMTOShipment{}
	this.MoveTaskOrderID = moveTaskOrderID
	this.ShipmentType = shipmentType
	return &this
}

// NewCreateMTOShipmentWithDefaults instantiates a new CreateMTOShipment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateMTOShipmentWithDefaults() *CreateMTOShipment {
	this := CreateMTOShipment{}
	return &this
}

// GetMoveTaskOrderID returns the MoveTaskOrderID field value
func (o *CreateMTOShipment) GetMoveTaskOrderID() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.MoveTaskOrderID
}

// GetMoveTaskOrderIDOk returns a tuple with the MoveTaskOrderID field value
// and a boolean to check if the value has been set.
func (o *CreateMTOShipment) GetMoveTaskOrderIDOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MoveTaskOrderID, true
}

// SetMoveTaskOrderID sets field value
func (o *CreateMTOShipment) SetMoveTaskOrderID(v string) {
	o.MoveTaskOrderID = v
}

// GetRequestedPickupDate returns the RequestedPickupDate field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateMTOShipment) GetRequestedPickupDate() string {
	if o == nil || IsNil(o.RequestedPickupDate.Get()) {
		var ret string
		return ret
	}
	return *o.RequestedPickupDate.Get()
}

// GetRequestedPickupDateOk returns a tuple with the RequestedPickupDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMTOShipment) GetRequestedPickupDateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequestedPickupDate.Get(), o.RequestedPickupDate.IsSet()
}

// HasRequestedPickupDate returns a boolean if a field has been set.
func (o *CreateMTOShipment) HasRequestedPickupDate() bool {
	if o != nil && o.RequestedPickupDate.IsSet() {
		return true
	}

	return false
}

// SetRequestedPickupDate gets a reference to the given NullableString and assigns it to the RequestedPickupDate field.
func (o *CreateMTOShipment) SetRequestedPickupDate(v string) {
	o.RequestedPickupDate.Set(&v)
}
// SetRequestedPickupDateNil sets the value for RequestedPickupDate to be an explicit nil
func (o *CreateMTOShipment) SetRequestedPickupDateNil() {
	o.RequestedPickupDate.Set(nil)
}

// UnsetRequestedPickupDate ensures that no value is present for RequestedPickupDate, not even an explicit nil
func (o *CreateMTOShipment) UnsetRequestedPickupDate() {
	o.RequestedPickupDate.Unset()
}

// GetPrimeEstimatedWeight returns the PrimeEstimatedWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateMTOShipment) GetPrimeEstimatedWeight() int32 {
	if o == nil || IsNil(o.PrimeEstimatedWeight.Get()) {
		var ret int32
		return ret
	}
	return *o.PrimeEstimatedWeight.Get()
}

// GetPrimeEstimatedWeightOk returns a tuple with the PrimeEstimatedWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMTOShipment) GetPrimeEstimatedWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrimeEstimatedWeight.Get(), o.PrimeEstimatedWeight.IsSet()
}

// HasPrimeEstimatedWeight returns a boolean if a field has been set.
func (o *CreateMTOShipment) HasPrimeEstimatedWeight() bool {
	if o != nil && o.PrimeEstimatedWeight.IsSet() {
		return true
	}

	return false
}

// SetPrimeEstimatedWeight gets a reference to the given NullableInt32 and assigns it to the PrimeEstimatedWeight field.
func (o *CreateMTOShipment) SetPrimeEstimatedWeight(v int32) {
	o.PrimeEstimatedWeight.Set(&v)
}
// SetPrimeEstimatedWeightNil sets the value for PrimeEstimatedWeight to be an explicit nil
func (o *CreateMTOShipment) SetPrimeEstimatedWeightNil() {
	o.PrimeEstimatedWeight.Set(nil)
}

// UnsetPrimeEstimatedWeight ensures that no value is present for PrimeEstimatedWeight, not even an explicit nil
func (o *CreateMTOShipment) UnsetPrimeEstimatedWeight() {
	o.PrimeEstimatedWeight.Unset()
}

// GetCustomerRemarks returns the CustomerRemarks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateMTOShipment) GetCustomerRemarks() string {
	if o == nil || IsNil(o.CustomerRemarks.Get()) {
		var ret string
		return ret
	}
	return *o.CustomerRemarks.Get()
}

// GetCustomerRemarksOk returns a tuple with the CustomerRemarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMTOShipment) GetCustomerRemarksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CustomerRemarks.Get(), o.CustomerRemarks.IsSet()
}

// HasCustomerRemarks returns a boolean if a field has been set.
func (o *CreateMTOShipment) HasCustomerRemarks() bool {
	if o != nil && o.CustomerRemarks.IsSet() {
		return true
	}

	return false
}

// SetCustomerRemarks gets a reference to the given NullableString and assigns it to the CustomerRemarks field.
func (o *CreateMTOShipment) SetCustomerRemarks(v string) {
	o.CustomerRemarks.Set(&v)
}
// SetCustomerRemarksNil sets the value for CustomerRemarks to be an explicit nil
func (o *CreateMTOShipment) SetCustomerRemarksNil() {
	o.CustomerRemarks.Set(nil)
}

// UnsetCustomerRemarks ensures that no value is present for CustomerRemarks, not even an explicit nil
func (o *CreateMTOShipment) UnsetCustomerRemarks() {
	o.CustomerRemarks.Unset()
}

// GetAgents returns the Agents field value if set, zero value otherwise.
func (o *CreateMTOShipment) GetAgents() []MTOAgent {
	if o == nil || IsNil(o.Agents) {
		var ret []MTOAgent
		return ret
	}
	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMTOShipment) GetAgentsOk() ([]MTOAgent, bool) {
	if o == nil || IsNil(o.Agents) {
		return nil, false
	}
	return o.Agents, true
}

// HasAgents returns a boolean if a field has been set.
func (o *CreateMTOShipment) HasAgents() bool {
	if o != nil && !IsNil(o.Agents) {
		return true
	}

	return false
}

// SetAgents gets a reference to the given []MTOAgent and assigns it to the Agents field.
func (o *CreateMTOShipment) SetAgents(v []MTOAgent) {
	o.Agents = v
}

// GetMtoServiceItems returns the MtoServiceItems field value if set, zero value otherwise.
func (o *CreateMTOShipment) GetMtoServiceItems() []MTOServiceItem {
	if o == nil || IsNil(o.MtoServiceItems) {
		var ret []MTOServiceItem
		return ret
	}
	return o.MtoServiceItems
}

// GetMtoServiceItemsOk returns a tuple with the MtoServiceItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMTOShipment) GetMtoServiceItemsOk() ([]MTOServiceItem, bool) {
	if o == nil || IsNil(o.MtoServiceItems) {
		return nil, false
	}
	return o.MtoServiceItems, true
}

// HasMtoServiceItems returns a boolean if a field has been set.
func (o *CreateMTOShipment) HasMtoServiceItems() bool {
	if o != nil && !IsNil(o.MtoServiceItems) {
		return true
	}

	return false
}

// SetMtoServiceItems gets a reference to the given []MTOServiceItem and assigns it to the MtoServiceItems field.
func (o *CreateMTOShipment) SetMtoServiceItems(v []MTOServiceItem) {
	o.MtoServiceItems = v
}

// GetPickupAddress returns the PickupAddress field value if set, zero value otherwise.
func (o *CreateMTOShipment) GetPickupAddress() Address {
	if o == nil || IsNil(o.PickupAddress) {
		var ret Address
		return ret
	}
	return *o.PickupAddress
}

// GetPickupAddressOk returns a tuple with the PickupAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMTOShipment) GetPickupAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.PickupAddress) {
		return nil, false
	}
	return o.PickupAddress, true
}

// HasPickupAddress returns a boolean if a field has been set.
func (o *CreateMTOShipment) HasPickupAddress() bool {
	if o != nil && !IsNil(o.PickupAddress) {
		return true
	}

	return false
}

// SetPickupAddress gets a reference to the given Address and assigns it to the PickupAddress field.
func (o *CreateMTOShipment) SetPickupAddress(v Address) {
	o.PickupAddress = &v
}

// GetDestinationAddress returns the DestinationAddress field value if set, zero value otherwise.
func (o *CreateMTOShipment) GetDestinationAddress() Address {
	if o == nil || IsNil(o.DestinationAddress) {
		var ret Address
		return ret
	}
	return *o.DestinationAddress
}

// GetDestinationAddressOk returns a tuple with the DestinationAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMTOShipment) GetDestinationAddressOk() (*Address, bool) {
	if o == nil || IsNil(o.DestinationAddress) {
		return nil, false
	}
	return o.DestinationAddress, true
}

// HasDestinationAddress returns a boolean if a field has been set.
func (o *CreateMTOShipment) HasDestinationAddress() bool {
	if o != nil && !IsNil(o.DestinationAddress) {
		return true
	}

	return false
}

// SetDestinationAddress gets a reference to the given Address and assigns it to the DestinationAddress field.
func (o *CreateMTOShipment) SetDestinationAddress(v Address) {
	o.DestinationAddress = &v
}

// GetShipmentType returns the ShipmentType field value
func (o *CreateMTOShipment) GetShipmentType() MTOShipmentType {
	if o == nil {
		var ret MTOShipmentType
		return ret
	}

	return o.ShipmentType
}

// GetShipmentTypeOk returns a tuple with the ShipmentType field value
// and a boolean to check if the value has been set.
func (o *CreateMTOShipment) GetShipmentTypeOk() (*MTOShipmentType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ShipmentType, true
}

// SetShipmentType sets field value
func (o *CreateMTOShipment) SetShipmentType(v MTOShipmentType) {
	o.ShipmentType = v
}

// GetDiversion returns the Diversion field value if set, zero value otherwise.
func (o *CreateMTOShipment) GetDiversion() bool {
	if o == nil || IsNil(o.Diversion) {
		var ret bool
		return ret
	}
	return *o.Diversion
}

// GetDiversionOk returns a tuple with the Diversion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMTOShipment) GetDiversionOk() (*bool, bool) {
	if o == nil || IsNil(o.Diversion) {
		return nil, false
	}
	return o.Diversion, true
}

// HasDiversion returns a boolean if a field has been set.
func (o *CreateMTOShipment) HasDiversion() bool {
	if o != nil && !IsNil(o.Diversion) {
		return true
	}

	return false
}

// SetDiversion gets a reference to the given bool and assigns it to the Diversion field.
func (o *CreateMTOShipment) SetDiversion(v bool) {
	o.Diversion = &v
}

// GetPointOfContact returns the PointOfContact field value if set, zero value otherwise.
func (o *CreateMTOShipment) GetPointOfContact() string {
	if o == nil || IsNil(o.PointOfContact) {
		var ret string
		return ret
	}
	return *o.PointOfContact
}

// GetPointOfContactOk returns a tuple with the PointOfContact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMTOShipment) GetPointOfContactOk() (*string, bool) {
	if o == nil || IsNil(o.PointOfContact) {
		return nil, false
	}
	return o.PointOfContact, true
}

// HasPointOfContact returns a boolean if a field has been set.
func (o *CreateMTOShipment) HasPointOfContact() bool {
	if o != nil && !IsNil(o.PointOfContact) {
		return true
	}

	return false
}

// SetPointOfContact gets a reference to the given string and assigns it to the PointOfContact field.
func (o *CreateMTOShipment) SetPointOfContact(v string) {
	o.PointOfContact = &v
}

// GetCounselorRemarks returns the CounselorRemarks field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CreateMTOShipment) GetCounselorRemarks() string {
	if o == nil || IsNil(o.CounselorRemarks.Get()) {
		var ret string
		return ret
	}
	return *o.CounselorRemarks.Get()
}

// GetCounselorRemarksOk returns a tuple with the CounselorRemarks field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CreateMTOShipment) GetCounselorRemarksOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CounselorRemarks.Get(), o.CounselorRemarks.IsSet()
}

// HasCounselorRemarks returns a boolean if a field has been set.
func (o *CreateMTOShipment) HasCounselorRemarks() bool {
	if o != nil && o.CounselorRemarks.IsSet() {
		return true
	}

	return false
}

// SetCounselorRemarks gets a reference to the given NullableString and assigns it to the CounselorRemarks field.
func (o *CreateMTOShipment) SetCounselorRemarks(v string) {
	o.CounselorRemarks.Set(&v)
}
// SetCounselorRemarksNil sets the value for CounselorRemarks to be an explicit nil
func (o *CreateMTOShipment) SetCounselorRemarksNil() {
	o.CounselorRemarks.Set(nil)
}

// UnsetCounselorRemarks ensures that no value is present for CounselorRemarks, not even an explicit nil
func (o *CreateMTOShipment) UnsetCounselorRemarks() {
	o.CounselorRemarks.Unset()
}

// GetPpmShipment returns the PpmShipment field value if set, zero value otherwise.
func (o *CreateMTOShipment) GetPpmShipment() CreatePPMShipment {
	if o == nil || IsNil(o.PpmShipment) {
		var ret CreatePPMShipment
		return ret
	}
	return *o.PpmShipment
}

// GetPpmShipmentOk returns a tuple with the PpmShipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateMTOShipment) GetPpmShipmentOk() (*CreatePPMShipment, bool) {
	if o == nil || IsNil(o.PpmShipment) {
		return nil, false
	}
	return o.PpmShipment, true
}

// HasPpmShipment returns a boolean if a field has been set.
func (o *CreateMTOShipment) HasPpmShipment() bool {
	if o != nil && !IsNil(o.PpmShipment) {
		return true
	}

	return false
}

// SetPpmShipment gets a reference to the given CreatePPMShipment and assigns it to the PpmShipment field.
func (o *CreateMTOShipment) SetPpmShipment(v CreatePPMShipment) {
	o.PpmShipment = &v
}

func (o CreateMTOShipment) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateMTOShipment) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["moveTaskOrderID"] = o.MoveTaskOrderID
	if o.RequestedPickupDate.IsSet() {
		toSerialize["requestedPickupDate"] = o.RequestedPickupDate.Get()
	}
	if o.PrimeEstimatedWeight.IsSet() {
		toSerialize["primeEstimatedWeight"] = o.PrimeEstimatedWeight.Get()
	}
	if o.CustomerRemarks.IsSet() {
		toSerialize["customerRemarks"] = o.CustomerRemarks.Get()
	}
	if !IsNil(o.Agents) {
		toSerialize["agents"] = o.Agents
	}
	if !IsNil(o.MtoServiceItems) {
		toSerialize["mtoServiceItems"] = o.MtoServiceItems
	}
	if !IsNil(o.PickupAddress) {
		toSerialize["pickupAddress"] = o.PickupAddress
	}
	if !IsNil(o.DestinationAddress) {
		toSerialize["destinationAddress"] = o.DestinationAddress
	}
	toSerialize["shipmentType"] = o.ShipmentType
	if !IsNil(o.Diversion) {
		toSerialize["diversion"] = o.Diversion
	}
	if !IsNil(o.PointOfContact) {
		toSerialize["pointOfContact"] = o.PointOfContact
	}
	if o.CounselorRemarks.IsSet() {
		toSerialize["counselorRemarks"] = o.CounselorRemarks.Get()
	}
	if !IsNil(o.PpmShipment) {
		toSerialize["ppmShipment"] = o.PpmShipment
	}
	return toSerialize, nil
}

func (o *CreateMTOShipment) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"moveTaskOrderID",
		"shipmentType",
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

	varCreateMTOShipment := _CreateMTOShipment{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateMTOShipment)

	if err != nil {
		return err
	}

	*o = CreateMTOShipment(varCreateMTOShipment)

	return err
}

type NullableCreateMTOShipment struct {
	value *CreateMTOShipment
	isSet bool
}

func (v NullableCreateMTOShipment) Get() *CreateMTOShipment {
	return v.value
}

func (v *NullableCreateMTOShipment) Set(val *CreateMTOShipment) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateMTOShipment) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateMTOShipment) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateMTOShipment(val *CreateMTOShipment) *NullableCreateMTOShipment {
	return &NullableCreateMTOShipment{value: val, isSet: true}
}

func (v NullableCreateMTOShipment) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateMTOShipment) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


