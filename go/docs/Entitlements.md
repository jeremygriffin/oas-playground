# Entitlements

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**AuthorizedWeight** | Pointer to **NullableInt32** |  | [optional] 
**DependentsAuthorized** | Pointer to **NullableBool** |  | [optional] 
**GunSafe** | Pointer to **bool** |  | [optional] 
**NonTemporaryStorage** | Pointer to **NullableBool** |  | [optional] 
**PrivatelyOwnedVehicle** | Pointer to **NullableBool** |  | [optional] 
**ProGearWeight** | Pointer to **int32** |  | [optional] 
**ProGearWeightSpouse** | Pointer to **int32** |  | [optional] 
**RequiredMedicalEquipmentWeight** | Pointer to **int32** |  | [optional] 
**OrganizationalClothingAndIndividualEquipment** | Pointer to **bool** |  | [optional] 
**StorageInTransit** | Pointer to **int32** |  | [optional] 
**TotalWeight** | Pointer to **int32** |  | [optional] 
**TotalDependents** | Pointer to **int32** |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewEntitlements

`func NewEntitlements() *Entitlements`

NewEntitlements instantiates a new Entitlements object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementsWithDefaults

`func NewEntitlementsWithDefaults() *Entitlements`

NewEntitlementsWithDefaults instantiates a new Entitlements object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Entitlements) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Entitlements) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Entitlements) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Entitlements) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthorizedWeight

`func (o *Entitlements) GetAuthorizedWeight() int32`

GetAuthorizedWeight returns the AuthorizedWeight field if non-nil, zero value otherwise.

### GetAuthorizedWeightOk

`func (o *Entitlements) GetAuthorizedWeightOk() (*int32, bool)`

GetAuthorizedWeightOk returns a tuple with the AuthorizedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedWeight

`func (o *Entitlements) SetAuthorizedWeight(v int32)`

SetAuthorizedWeight sets AuthorizedWeight field to given value.

### HasAuthorizedWeight

`func (o *Entitlements) HasAuthorizedWeight() bool`

HasAuthorizedWeight returns a boolean if a field has been set.

### SetAuthorizedWeightNil

`func (o *Entitlements) SetAuthorizedWeightNil(b bool)`

 SetAuthorizedWeightNil sets the value for AuthorizedWeight to be an explicit nil

### UnsetAuthorizedWeight
`func (o *Entitlements) UnsetAuthorizedWeight()`

UnsetAuthorizedWeight ensures that no value is present for AuthorizedWeight, not even an explicit nil
### GetDependentsAuthorized

`func (o *Entitlements) GetDependentsAuthorized() bool`

GetDependentsAuthorized returns the DependentsAuthorized field if non-nil, zero value otherwise.

### GetDependentsAuthorizedOk

`func (o *Entitlements) GetDependentsAuthorizedOk() (*bool, bool)`

GetDependentsAuthorizedOk returns a tuple with the DependentsAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentsAuthorized

`func (o *Entitlements) SetDependentsAuthorized(v bool)`

SetDependentsAuthorized sets DependentsAuthorized field to given value.

### HasDependentsAuthorized

`func (o *Entitlements) HasDependentsAuthorized() bool`

HasDependentsAuthorized returns a boolean if a field has been set.

### SetDependentsAuthorizedNil

`func (o *Entitlements) SetDependentsAuthorizedNil(b bool)`

 SetDependentsAuthorizedNil sets the value for DependentsAuthorized to be an explicit nil

### UnsetDependentsAuthorized
`func (o *Entitlements) UnsetDependentsAuthorized()`

UnsetDependentsAuthorized ensures that no value is present for DependentsAuthorized, not even an explicit nil
### GetGunSafe

`func (o *Entitlements) GetGunSafe() bool`

GetGunSafe returns the GunSafe field if non-nil, zero value otherwise.

### GetGunSafeOk

`func (o *Entitlements) GetGunSafeOk() (*bool, bool)`

GetGunSafeOk returns a tuple with the GunSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGunSafe

`func (o *Entitlements) SetGunSafe(v bool)`

SetGunSafe sets GunSafe field to given value.

### HasGunSafe

`func (o *Entitlements) HasGunSafe() bool`

HasGunSafe returns a boolean if a field has been set.

### GetNonTemporaryStorage

`func (o *Entitlements) GetNonTemporaryStorage() bool`

GetNonTemporaryStorage returns the NonTemporaryStorage field if non-nil, zero value otherwise.

### GetNonTemporaryStorageOk

`func (o *Entitlements) GetNonTemporaryStorageOk() (*bool, bool)`

GetNonTemporaryStorageOk returns a tuple with the NonTemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonTemporaryStorage

`func (o *Entitlements) SetNonTemporaryStorage(v bool)`

SetNonTemporaryStorage sets NonTemporaryStorage field to given value.

### HasNonTemporaryStorage

`func (o *Entitlements) HasNonTemporaryStorage() bool`

HasNonTemporaryStorage returns a boolean if a field has been set.

### SetNonTemporaryStorageNil

`func (o *Entitlements) SetNonTemporaryStorageNil(b bool)`

 SetNonTemporaryStorageNil sets the value for NonTemporaryStorage to be an explicit nil

### UnsetNonTemporaryStorage
`func (o *Entitlements) UnsetNonTemporaryStorage()`

UnsetNonTemporaryStorage ensures that no value is present for NonTemporaryStorage, not even an explicit nil
### GetPrivatelyOwnedVehicle

`func (o *Entitlements) GetPrivatelyOwnedVehicle() bool`

GetPrivatelyOwnedVehicle returns the PrivatelyOwnedVehicle field if non-nil, zero value otherwise.

### GetPrivatelyOwnedVehicleOk

`func (o *Entitlements) GetPrivatelyOwnedVehicleOk() (*bool, bool)`

GetPrivatelyOwnedVehicleOk returns a tuple with the PrivatelyOwnedVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivatelyOwnedVehicle

`func (o *Entitlements) SetPrivatelyOwnedVehicle(v bool)`

SetPrivatelyOwnedVehicle sets PrivatelyOwnedVehicle field to given value.

### HasPrivatelyOwnedVehicle

`func (o *Entitlements) HasPrivatelyOwnedVehicle() bool`

HasPrivatelyOwnedVehicle returns a boolean if a field has been set.

### SetPrivatelyOwnedVehicleNil

`func (o *Entitlements) SetPrivatelyOwnedVehicleNil(b bool)`

 SetPrivatelyOwnedVehicleNil sets the value for PrivatelyOwnedVehicle to be an explicit nil

### UnsetPrivatelyOwnedVehicle
`func (o *Entitlements) UnsetPrivatelyOwnedVehicle()`

UnsetPrivatelyOwnedVehicle ensures that no value is present for PrivatelyOwnedVehicle, not even an explicit nil
### GetProGearWeight

`func (o *Entitlements) GetProGearWeight() int32`

GetProGearWeight returns the ProGearWeight field if non-nil, zero value otherwise.

### GetProGearWeightOk

`func (o *Entitlements) GetProGearWeightOk() (*int32, bool)`

GetProGearWeightOk returns a tuple with the ProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProGearWeight

`func (o *Entitlements) SetProGearWeight(v int32)`

SetProGearWeight sets ProGearWeight field to given value.

### HasProGearWeight

`func (o *Entitlements) HasProGearWeight() bool`

HasProGearWeight returns a boolean if a field has been set.

### GetProGearWeightSpouse

`func (o *Entitlements) GetProGearWeightSpouse() int32`

GetProGearWeightSpouse returns the ProGearWeightSpouse field if non-nil, zero value otherwise.

### GetProGearWeightSpouseOk

`func (o *Entitlements) GetProGearWeightSpouseOk() (*int32, bool)`

GetProGearWeightSpouseOk returns a tuple with the ProGearWeightSpouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProGearWeightSpouse

`func (o *Entitlements) SetProGearWeightSpouse(v int32)`

SetProGearWeightSpouse sets ProGearWeightSpouse field to given value.

### HasProGearWeightSpouse

`func (o *Entitlements) HasProGearWeightSpouse() bool`

HasProGearWeightSpouse returns a boolean if a field has been set.

### GetRequiredMedicalEquipmentWeight

`func (o *Entitlements) GetRequiredMedicalEquipmentWeight() int32`

GetRequiredMedicalEquipmentWeight returns the RequiredMedicalEquipmentWeight field if non-nil, zero value otherwise.

### GetRequiredMedicalEquipmentWeightOk

`func (o *Entitlements) GetRequiredMedicalEquipmentWeightOk() (*int32, bool)`

GetRequiredMedicalEquipmentWeightOk returns a tuple with the RequiredMedicalEquipmentWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredMedicalEquipmentWeight

`func (o *Entitlements) SetRequiredMedicalEquipmentWeight(v int32)`

SetRequiredMedicalEquipmentWeight sets RequiredMedicalEquipmentWeight field to given value.

### HasRequiredMedicalEquipmentWeight

`func (o *Entitlements) HasRequiredMedicalEquipmentWeight() bool`

HasRequiredMedicalEquipmentWeight returns a boolean if a field has been set.

### GetOrganizationalClothingAndIndividualEquipment

`func (o *Entitlements) GetOrganizationalClothingAndIndividualEquipment() bool`

GetOrganizationalClothingAndIndividualEquipment returns the OrganizationalClothingAndIndividualEquipment field if non-nil, zero value otherwise.

### GetOrganizationalClothingAndIndividualEquipmentOk

`func (o *Entitlements) GetOrganizationalClothingAndIndividualEquipmentOk() (*bool, bool)`

GetOrganizationalClothingAndIndividualEquipmentOk returns a tuple with the OrganizationalClothingAndIndividualEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalClothingAndIndividualEquipment

`func (o *Entitlements) SetOrganizationalClothingAndIndividualEquipment(v bool)`

SetOrganizationalClothingAndIndividualEquipment sets OrganizationalClothingAndIndividualEquipment field to given value.

### HasOrganizationalClothingAndIndividualEquipment

`func (o *Entitlements) HasOrganizationalClothingAndIndividualEquipment() bool`

HasOrganizationalClothingAndIndividualEquipment returns a boolean if a field has been set.

### GetStorageInTransit

`func (o *Entitlements) GetStorageInTransit() int32`

GetStorageInTransit returns the StorageInTransit field if non-nil, zero value otherwise.

### GetStorageInTransitOk

`func (o *Entitlements) GetStorageInTransitOk() (*int32, bool)`

GetStorageInTransitOk returns a tuple with the StorageInTransit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInTransit

`func (o *Entitlements) SetStorageInTransit(v int32)`

SetStorageInTransit sets StorageInTransit field to given value.

### HasStorageInTransit

`func (o *Entitlements) HasStorageInTransit() bool`

HasStorageInTransit returns a boolean if a field has been set.

### GetTotalWeight

`func (o *Entitlements) GetTotalWeight() int32`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *Entitlements) GetTotalWeightOk() (*int32, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *Entitlements) SetTotalWeight(v int32)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *Entitlements) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetTotalDependents

`func (o *Entitlements) GetTotalDependents() int32`

GetTotalDependents returns the TotalDependents field if non-nil, zero value otherwise.

### GetTotalDependentsOk

`func (o *Entitlements) GetTotalDependentsOk() (*int32, bool)`

GetTotalDependentsOk returns a tuple with the TotalDependents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDependents

`func (o *Entitlements) SetTotalDependents(v int32)`

SetTotalDependents sets TotalDependents field to given value.

### HasTotalDependents

`func (o *Entitlements) HasTotalDependents() bool`

HasTotalDependents returns a boolean if a field has been set.

### GetETag

`func (o *Entitlements) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *Entitlements) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *Entitlements) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *Entitlements) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


