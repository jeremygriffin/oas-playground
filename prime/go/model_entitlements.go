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

// checks if the Entitlements type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Entitlements{}

// Entitlements struct for Entitlements
type Entitlements struct {
	Id *string `json:"id,omitempty"`
	AuthorizedWeight NullableInt32 `json:"authorizedWeight,omitempty"`
	DependentsAuthorized NullableBool `json:"dependentsAuthorized,omitempty"`
	GunSafe *bool `json:"gunSafe,omitempty"`
	NonTemporaryStorage NullableBool `json:"nonTemporaryStorage,omitempty"`
	PrivatelyOwnedVehicle NullableBool `json:"privatelyOwnedVehicle,omitempty"`
	ProGearWeight *int32 `json:"proGearWeight,omitempty"`
	ProGearWeightSpouse *int32 `json:"proGearWeightSpouse,omitempty"`
	RequiredMedicalEquipmentWeight *int32 `json:"requiredMedicalEquipmentWeight,omitempty"`
	OrganizationalClothingAndIndividualEquipment *bool `json:"organizationalClothingAndIndividualEquipment,omitempty"`
	StorageInTransit *int32 `json:"storageInTransit,omitempty"`
	TotalWeight *int32 `json:"totalWeight,omitempty"`
	TotalDependents *int32 `json:"totalDependents,omitempty"`
	ETag *string `json:"eTag,omitempty"`
}

// NewEntitlements instantiates a new Entitlements object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEntitlements() *Entitlements {
	this := Entitlements{}
	return &this
}

// NewEntitlementsWithDefaults instantiates a new Entitlements object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEntitlementsWithDefaults() *Entitlements {
	this := Entitlements{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Entitlements) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlements) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Entitlements) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Entitlements) SetId(v string) {
	o.Id = &v
}

// GetAuthorizedWeight returns the AuthorizedWeight field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entitlements) GetAuthorizedWeight() int32 {
	if o == nil || IsNil(o.AuthorizedWeight.Get()) {
		var ret int32
		return ret
	}
	return *o.AuthorizedWeight.Get()
}

// GetAuthorizedWeightOk returns a tuple with the AuthorizedWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entitlements) GetAuthorizedWeightOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.AuthorizedWeight.Get(), o.AuthorizedWeight.IsSet()
}

// HasAuthorizedWeight returns a boolean if a field has been set.
func (o *Entitlements) HasAuthorizedWeight() bool {
	if o != nil && o.AuthorizedWeight.IsSet() {
		return true
	}

	return false
}

// SetAuthorizedWeight gets a reference to the given NullableInt32 and assigns it to the AuthorizedWeight field.
func (o *Entitlements) SetAuthorizedWeight(v int32) {
	o.AuthorizedWeight.Set(&v)
}
// SetAuthorizedWeightNil sets the value for AuthorizedWeight to be an explicit nil
func (o *Entitlements) SetAuthorizedWeightNil() {
	o.AuthorizedWeight.Set(nil)
}

// UnsetAuthorizedWeight ensures that no value is present for AuthorizedWeight, not even an explicit nil
func (o *Entitlements) UnsetAuthorizedWeight() {
	o.AuthorizedWeight.Unset()
}

// GetDependentsAuthorized returns the DependentsAuthorized field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entitlements) GetDependentsAuthorized() bool {
	if o == nil || IsNil(o.DependentsAuthorized.Get()) {
		var ret bool
		return ret
	}
	return *o.DependentsAuthorized.Get()
}

// GetDependentsAuthorizedOk returns a tuple with the DependentsAuthorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entitlements) GetDependentsAuthorizedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.DependentsAuthorized.Get(), o.DependentsAuthorized.IsSet()
}

// HasDependentsAuthorized returns a boolean if a field has been set.
func (o *Entitlements) HasDependentsAuthorized() bool {
	if o != nil && o.DependentsAuthorized.IsSet() {
		return true
	}

	return false
}

// SetDependentsAuthorized gets a reference to the given NullableBool and assigns it to the DependentsAuthorized field.
func (o *Entitlements) SetDependentsAuthorized(v bool) {
	o.DependentsAuthorized.Set(&v)
}
// SetDependentsAuthorizedNil sets the value for DependentsAuthorized to be an explicit nil
func (o *Entitlements) SetDependentsAuthorizedNil() {
	o.DependentsAuthorized.Set(nil)
}

// UnsetDependentsAuthorized ensures that no value is present for DependentsAuthorized, not even an explicit nil
func (o *Entitlements) UnsetDependentsAuthorized() {
	o.DependentsAuthorized.Unset()
}

// GetGunSafe returns the GunSafe field value if set, zero value otherwise.
func (o *Entitlements) GetGunSafe() bool {
	if o == nil || IsNil(o.GunSafe) {
		var ret bool
		return ret
	}
	return *o.GunSafe
}

// GetGunSafeOk returns a tuple with the GunSafe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlements) GetGunSafeOk() (*bool, bool) {
	if o == nil || IsNil(o.GunSafe) {
		return nil, false
	}
	return o.GunSafe, true
}

// HasGunSafe returns a boolean if a field has been set.
func (o *Entitlements) HasGunSafe() bool {
	if o != nil && !IsNil(o.GunSafe) {
		return true
	}

	return false
}

// SetGunSafe gets a reference to the given bool and assigns it to the GunSafe field.
func (o *Entitlements) SetGunSafe(v bool) {
	o.GunSafe = &v
}

// GetNonTemporaryStorage returns the NonTemporaryStorage field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entitlements) GetNonTemporaryStorage() bool {
	if o == nil || IsNil(o.NonTemporaryStorage.Get()) {
		var ret bool
		return ret
	}
	return *o.NonTemporaryStorage.Get()
}

// GetNonTemporaryStorageOk returns a tuple with the NonTemporaryStorage field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entitlements) GetNonTemporaryStorageOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.NonTemporaryStorage.Get(), o.NonTemporaryStorage.IsSet()
}

// HasNonTemporaryStorage returns a boolean if a field has been set.
func (o *Entitlements) HasNonTemporaryStorage() bool {
	if o != nil && o.NonTemporaryStorage.IsSet() {
		return true
	}

	return false
}

// SetNonTemporaryStorage gets a reference to the given NullableBool and assigns it to the NonTemporaryStorage field.
func (o *Entitlements) SetNonTemporaryStorage(v bool) {
	o.NonTemporaryStorage.Set(&v)
}
// SetNonTemporaryStorageNil sets the value for NonTemporaryStorage to be an explicit nil
func (o *Entitlements) SetNonTemporaryStorageNil() {
	o.NonTemporaryStorage.Set(nil)
}

// UnsetNonTemporaryStorage ensures that no value is present for NonTemporaryStorage, not even an explicit nil
func (o *Entitlements) UnsetNonTemporaryStorage() {
	o.NonTemporaryStorage.Unset()
}

// GetPrivatelyOwnedVehicle returns the PrivatelyOwnedVehicle field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *Entitlements) GetPrivatelyOwnedVehicle() bool {
	if o == nil || IsNil(o.PrivatelyOwnedVehicle.Get()) {
		var ret bool
		return ret
	}
	return *o.PrivatelyOwnedVehicle.Get()
}

// GetPrivatelyOwnedVehicleOk returns a tuple with the PrivatelyOwnedVehicle field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Entitlements) GetPrivatelyOwnedVehicleOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.PrivatelyOwnedVehicle.Get(), o.PrivatelyOwnedVehicle.IsSet()
}

// HasPrivatelyOwnedVehicle returns a boolean if a field has been set.
func (o *Entitlements) HasPrivatelyOwnedVehicle() bool {
	if o != nil && o.PrivatelyOwnedVehicle.IsSet() {
		return true
	}

	return false
}

// SetPrivatelyOwnedVehicle gets a reference to the given NullableBool and assigns it to the PrivatelyOwnedVehicle field.
func (o *Entitlements) SetPrivatelyOwnedVehicle(v bool) {
	o.PrivatelyOwnedVehicle.Set(&v)
}
// SetPrivatelyOwnedVehicleNil sets the value for PrivatelyOwnedVehicle to be an explicit nil
func (o *Entitlements) SetPrivatelyOwnedVehicleNil() {
	o.PrivatelyOwnedVehicle.Set(nil)
}

// UnsetPrivatelyOwnedVehicle ensures that no value is present for PrivatelyOwnedVehicle, not even an explicit nil
func (o *Entitlements) UnsetPrivatelyOwnedVehicle() {
	o.PrivatelyOwnedVehicle.Unset()
}

// GetProGearWeight returns the ProGearWeight field value if set, zero value otherwise.
func (o *Entitlements) GetProGearWeight() int32 {
	if o == nil || IsNil(o.ProGearWeight) {
		var ret int32
		return ret
	}
	return *o.ProGearWeight
}

// GetProGearWeightOk returns a tuple with the ProGearWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlements) GetProGearWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.ProGearWeight) {
		return nil, false
	}
	return o.ProGearWeight, true
}

// HasProGearWeight returns a boolean if a field has been set.
func (o *Entitlements) HasProGearWeight() bool {
	if o != nil && !IsNil(o.ProGearWeight) {
		return true
	}

	return false
}

// SetProGearWeight gets a reference to the given int32 and assigns it to the ProGearWeight field.
func (o *Entitlements) SetProGearWeight(v int32) {
	o.ProGearWeight = &v
}

// GetProGearWeightSpouse returns the ProGearWeightSpouse field value if set, zero value otherwise.
func (o *Entitlements) GetProGearWeightSpouse() int32 {
	if o == nil || IsNil(o.ProGearWeightSpouse) {
		var ret int32
		return ret
	}
	return *o.ProGearWeightSpouse
}

// GetProGearWeightSpouseOk returns a tuple with the ProGearWeightSpouse field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlements) GetProGearWeightSpouseOk() (*int32, bool) {
	if o == nil || IsNil(o.ProGearWeightSpouse) {
		return nil, false
	}
	return o.ProGearWeightSpouse, true
}

// HasProGearWeightSpouse returns a boolean if a field has been set.
func (o *Entitlements) HasProGearWeightSpouse() bool {
	if o != nil && !IsNil(o.ProGearWeightSpouse) {
		return true
	}

	return false
}

// SetProGearWeightSpouse gets a reference to the given int32 and assigns it to the ProGearWeightSpouse field.
func (o *Entitlements) SetProGearWeightSpouse(v int32) {
	o.ProGearWeightSpouse = &v
}

// GetRequiredMedicalEquipmentWeight returns the RequiredMedicalEquipmentWeight field value if set, zero value otherwise.
func (o *Entitlements) GetRequiredMedicalEquipmentWeight() int32 {
	if o == nil || IsNil(o.RequiredMedicalEquipmentWeight) {
		var ret int32
		return ret
	}
	return *o.RequiredMedicalEquipmentWeight
}

// GetRequiredMedicalEquipmentWeightOk returns a tuple with the RequiredMedicalEquipmentWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlements) GetRequiredMedicalEquipmentWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.RequiredMedicalEquipmentWeight) {
		return nil, false
	}
	return o.RequiredMedicalEquipmentWeight, true
}

// HasRequiredMedicalEquipmentWeight returns a boolean if a field has been set.
func (o *Entitlements) HasRequiredMedicalEquipmentWeight() bool {
	if o != nil && !IsNil(o.RequiredMedicalEquipmentWeight) {
		return true
	}

	return false
}

// SetRequiredMedicalEquipmentWeight gets a reference to the given int32 and assigns it to the RequiredMedicalEquipmentWeight field.
func (o *Entitlements) SetRequiredMedicalEquipmentWeight(v int32) {
	o.RequiredMedicalEquipmentWeight = &v
}

// GetOrganizationalClothingAndIndividualEquipment returns the OrganizationalClothingAndIndividualEquipment field value if set, zero value otherwise.
func (o *Entitlements) GetOrganizationalClothingAndIndividualEquipment() bool {
	if o == nil || IsNil(o.OrganizationalClothingAndIndividualEquipment) {
		var ret bool
		return ret
	}
	return *o.OrganizationalClothingAndIndividualEquipment
}

// GetOrganizationalClothingAndIndividualEquipmentOk returns a tuple with the OrganizationalClothingAndIndividualEquipment field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlements) GetOrganizationalClothingAndIndividualEquipmentOk() (*bool, bool) {
	if o == nil || IsNil(o.OrganizationalClothingAndIndividualEquipment) {
		return nil, false
	}
	return o.OrganizationalClothingAndIndividualEquipment, true
}

// HasOrganizationalClothingAndIndividualEquipment returns a boolean if a field has been set.
func (o *Entitlements) HasOrganizationalClothingAndIndividualEquipment() bool {
	if o != nil && !IsNil(o.OrganizationalClothingAndIndividualEquipment) {
		return true
	}

	return false
}

// SetOrganizationalClothingAndIndividualEquipment gets a reference to the given bool and assigns it to the OrganizationalClothingAndIndividualEquipment field.
func (o *Entitlements) SetOrganizationalClothingAndIndividualEquipment(v bool) {
	o.OrganizationalClothingAndIndividualEquipment = &v
}

// GetStorageInTransit returns the StorageInTransit field value if set, zero value otherwise.
func (o *Entitlements) GetStorageInTransit() int32 {
	if o == nil || IsNil(o.StorageInTransit) {
		var ret int32
		return ret
	}
	return *o.StorageInTransit
}

// GetStorageInTransitOk returns a tuple with the StorageInTransit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlements) GetStorageInTransitOk() (*int32, bool) {
	if o == nil || IsNil(o.StorageInTransit) {
		return nil, false
	}
	return o.StorageInTransit, true
}

// HasStorageInTransit returns a boolean if a field has been set.
func (o *Entitlements) HasStorageInTransit() bool {
	if o != nil && !IsNil(o.StorageInTransit) {
		return true
	}

	return false
}

// SetStorageInTransit gets a reference to the given int32 and assigns it to the StorageInTransit field.
func (o *Entitlements) SetStorageInTransit(v int32) {
	o.StorageInTransit = &v
}

// GetTotalWeight returns the TotalWeight field value if set, zero value otherwise.
func (o *Entitlements) GetTotalWeight() int32 {
	if o == nil || IsNil(o.TotalWeight) {
		var ret int32
		return ret
	}
	return *o.TotalWeight
}

// GetTotalWeightOk returns a tuple with the TotalWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlements) GetTotalWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalWeight) {
		return nil, false
	}
	return o.TotalWeight, true
}

// HasTotalWeight returns a boolean if a field has been set.
func (o *Entitlements) HasTotalWeight() bool {
	if o != nil && !IsNil(o.TotalWeight) {
		return true
	}

	return false
}

// SetTotalWeight gets a reference to the given int32 and assigns it to the TotalWeight field.
func (o *Entitlements) SetTotalWeight(v int32) {
	o.TotalWeight = &v
}

// GetTotalDependents returns the TotalDependents field value if set, zero value otherwise.
func (o *Entitlements) GetTotalDependents() int32 {
	if o == nil || IsNil(o.TotalDependents) {
		var ret int32
		return ret
	}
	return *o.TotalDependents
}

// GetTotalDependentsOk returns a tuple with the TotalDependents field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlements) GetTotalDependentsOk() (*int32, bool) {
	if o == nil || IsNil(o.TotalDependents) {
		return nil, false
	}
	return o.TotalDependents, true
}

// HasTotalDependents returns a boolean if a field has been set.
func (o *Entitlements) HasTotalDependents() bool {
	if o != nil && !IsNil(o.TotalDependents) {
		return true
	}

	return false
}

// SetTotalDependents gets a reference to the given int32 and assigns it to the TotalDependents field.
func (o *Entitlements) SetTotalDependents(v int32) {
	o.TotalDependents = &v
}

// GetETag returns the ETag field value if set, zero value otherwise.
func (o *Entitlements) GetETag() string {
	if o == nil || IsNil(o.ETag) {
		var ret string
		return ret
	}
	return *o.ETag
}

// GetETagOk returns a tuple with the ETag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Entitlements) GetETagOk() (*string, bool) {
	if o == nil || IsNil(o.ETag) {
		return nil, false
	}
	return o.ETag, true
}

// HasETag returns a boolean if a field has been set.
func (o *Entitlements) HasETag() bool {
	if o != nil && !IsNil(o.ETag) {
		return true
	}

	return false
}

// SetETag gets a reference to the given string and assigns it to the ETag field.
func (o *Entitlements) SetETag(v string) {
	o.ETag = &v
}

func (o Entitlements) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Entitlements) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if o.AuthorizedWeight.IsSet() {
		toSerialize["authorizedWeight"] = o.AuthorizedWeight.Get()
	}
	if o.DependentsAuthorized.IsSet() {
		toSerialize["dependentsAuthorized"] = o.DependentsAuthorized.Get()
	}
	if !IsNil(o.GunSafe) {
		toSerialize["gunSafe"] = o.GunSafe
	}
	if o.NonTemporaryStorage.IsSet() {
		toSerialize["nonTemporaryStorage"] = o.NonTemporaryStorage.Get()
	}
	if o.PrivatelyOwnedVehicle.IsSet() {
		toSerialize["privatelyOwnedVehicle"] = o.PrivatelyOwnedVehicle.Get()
	}
	if !IsNil(o.ProGearWeight) {
		toSerialize["proGearWeight"] = o.ProGearWeight
	}
	if !IsNil(o.ProGearWeightSpouse) {
		toSerialize["proGearWeightSpouse"] = o.ProGearWeightSpouse
	}
	if !IsNil(o.RequiredMedicalEquipmentWeight) {
		toSerialize["requiredMedicalEquipmentWeight"] = o.RequiredMedicalEquipmentWeight
	}
	if !IsNil(o.OrganizationalClothingAndIndividualEquipment) {
		toSerialize["organizationalClothingAndIndividualEquipment"] = o.OrganizationalClothingAndIndividualEquipment
	}
	if !IsNil(o.StorageInTransit) {
		toSerialize["storageInTransit"] = o.StorageInTransit
	}
	if !IsNil(o.TotalWeight) {
		toSerialize["totalWeight"] = o.TotalWeight
	}
	if !IsNil(o.TotalDependents) {
		toSerialize["totalDependents"] = o.TotalDependents
	}
	if !IsNil(o.ETag) {
		toSerialize["eTag"] = o.ETag
	}
	return toSerialize, nil
}

type NullableEntitlements struct {
	value *Entitlements
	isSet bool
}

func (v NullableEntitlements) Get() *Entitlements {
	return v.value
}

func (v *NullableEntitlements) Set(val *Entitlements) {
	v.value = val
	v.isSet = true
}

func (v NullableEntitlements) IsSet() bool {
	return v.isSet
}

func (v *NullableEntitlements) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEntitlements(val *Entitlements) *NullableEntitlements {
	return &NullableEntitlements{value: val, isSet: true}
}

func (v NullableEntitlements) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEntitlements) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


