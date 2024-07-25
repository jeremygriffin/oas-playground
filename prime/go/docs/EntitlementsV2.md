# EntitlementsV2

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

### NewEntitlementsV2

`func NewEntitlementsV2() *EntitlementsV2`

NewEntitlementsV2 instantiates a new EntitlementsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEntitlementsV2WithDefaults

`func NewEntitlementsV2WithDefaults() *EntitlementsV2`

NewEntitlementsV2WithDefaults instantiates a new EntitlementsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EntitlementsV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EntitlementsV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EntitlementsV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *EntitlementsV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthorizedWeight

`func (o *EntitlementsV2) GetAuthorizedWeight() int32`

GetAuthorizedWeight returns the AuthorizedWeight field if non-nil, zero value otherwise.

### GetAuthorizedWeightOk

`func (o *EntitlementsV2) GetAuthorizedWeightOk() (*int32, bool)`

GetAuthorizedWeightOk returns a tuple with the AuthorizedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthorizedWeight

`func (o *EntitlementsV2) SetAuthorizedWeight(v int32)`

SetAuthorizedWeight sets AuthorizedWeight field to given value.

### HasAuthorizedWeight

`func (o *EntitlementsV2) HasAuthorizedWeight() bool`

HasAuthorizedWeight returns a boolean if a field has been set.

### SetAuthorizedWeightNil

`func (o *EntitlementsV2) SetAuthorizedWeightNil(b bool)`

 SetAuthorizedWeightNil sets the value for AuthorizedWeight to be an explicit nil

### UnsetAuthorizedWeight
`func (o *EntitlementsV2) UnsetAuthorizedWeight()`

UnsetAuthorizedWeight ensures that no value is present for AuthorizedWeight, not even an explicit nil
### GetDependentsAuthorized

`func (o *EntitlementsV2) GetDependentsAuthorized() bool`

GetDependentsAuthorized returns the DependentsAuthorized field if non-nil, zero value otherwise.

### GetDependentsAuthorizedOk

`func (o *EntitlementsV2) GetDependentsAuthorizedOk() (*bool, bool)`

GetDependentsAuthorizedOk returns a tuple with the DependentsAuthorized field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependentsAuthorized

`func (o *EntitlementsV2) SetDependentsAuthorized(v bool)`

SetDependentsAuthorized sets DependentsAuthorized field to given value.

### HasDependentsAuthorized

`func (o *EntitlementsV2) HasDependentsAuthorized() bool`

HasDependentsAuthorized returns a boolean if a field has been set.

### SetDependentsAuthorizedNil

`func (o *EntitlementsV2) SetDependentsAuthorizedNil(b bool)`

 SetDependentsAuthorizedNil sets the value for DependentsAuthorized to be an explicit nil

### UnsetDependentsAuthorized
`func (o *EntitlementsV2) UnsetDependentsAuthorized()`

UnsetDependentsAuthorized ensures that no value is present for DependentsAuthorized, not even an explicit nil
### GetGunSafe

`func (o *EntitlementsV2) GetGunSafe() bool`

GetGunSafe returns the GunSafe field if non-nil, zero value otherwise.

### GetGunSafeOk

`func (o *EntitlementsV2) GetGunSafeOk() (*bool, bool)`

GetGunSafeOk returns a tuple with the GunSafe field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGunSafe

`func (o *EntitlementsV2) SetGunSafe(v bool)`

SetGunSafe sets GunSafe field to given value.

### HasGunSafe

`func (o *EntitlementsV2) HasGunSafe() bool`

HasGunSafe returns a boolean if a field has been set.

### GetNonTemporaryStorage

`func (o *EntitlementsV2) GetNonTemporaryStorage() bool`

GetNonTemporaryStorage returns the NonTemporaryStorage field if non-nil, zero value otherwise.

### GetNonTemporaryStorageOk

`func (o *EntitlementsV2) GetNonTemporaryStorageOk() (*bool, bool)`

GetNonTemporaryStorageOk returns a tuple with the NonTemporaryStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonTemporaryStorage

`func (o *EntitlementsV2) SetNonTemporaryStorage(v bool)`

SetNonTemporaryStorage sets NonTemporaryStorage field to given value.

### HasNonTemporaryStorage

`func (o *EntitlementsV2) HasNonTemporaryStorage() bool`

HasNonTemporaryStorage returns a boolean if a field has been set.

### SetNonTemporaryStorageNil

`func (o *EntitlementsV2) SetNonTemporaryStorageNil(b bool)`

 SetNonTemporaryStorageNil sets the value for NonTemporaryStorage to be an explicit nil

### UnsetNonTemporaryStorage
`func (o *EntitlementsV2) UnsetNonTemporaryStorage()`

UnsetNonTemporaryStorage ensures that no value is present for NonTemporaryStorage, not even an explicit nil
### GetPrivatelyOwnedVehicle

`func (o *EntitlementsV2) GetPrivatelyOwnedVehicle() bool`

GetPrivatelyOwnedVehicle returns the PrivatelyOwnedVehicle field if non-nil, zero value otherwise.

### GetPrivatelyOwnedVehicleOk

`func (o *EntitlementsV2) GetPrivatelyOwnedVehicleOk() (*bool, bool)`

GetPrivatelyOwnedVehicleOk returns a tuple with the PrivatelyOwnedVehicle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivatelyOwnedVehicle

`func (o *EntitlementsV2) SetPrivatelyOwnedVehicle(v bool)`

SetPrivatelyOwnedVehicle sets PrivatelyOwnedVehicle field to given value.

### HasPrivatelyOwnedVehicle

`func (o *EntitlementsV2) HasPrivatelyOwnedVehicle() bool`

HasPrivatelyOwnedVehicle returns a boolean if a field has been set.

### SetPrivatelyOwnedVehicleNil

`func (o *EntitlementsV2) SetPrivatelyOwnedVehicleNil(b bool)`

 SetPrivatelyOwnedVehicleNil sets the value for PrivatelyOwnedVehicle to be an explicit nil

### UnsetPrivatelyOwnedVehicle
`func (o *EntitlementsV2) UnsetPrivatelyOwnedVehicle()`

UnsetPrivatelyOwnedVehicle ensures that no value is present for PrivatelyOwnedVehicle, not even an explicit nil
### GetProGearWeight

`func (o *EntitlementsV2) GetProGearWeight() int32`

GetProGearWeight returns the ProGearWeight field if non-nil, zero value otherwise.

### GetProGearWeightOk

`func (o *EntitlementsV2) GetProGearWeightOk() (*int32, bool)`

GetProGearWeightOk returns a tuple with the ProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProGearWeight

`func (o *EntitlementsV2) SetProGearWeight(v int32)`

SetProGearWeight sets ProGearWeight field to given value.

### HasProGearWeight

`func (o *EntitlementsV2) HasProGearWeight() bool`

HasProGearWeight returns a boolean if a field has been set.

### GetProGearWeightSpouse

`func (o *EntitlementsV2) GetProGearWeightSpouse() int32`

GetProGearWeightSpouse returns the ProGearWeightSpouse field if non-nil, zero value otherwise.

### GetProGearWeightSpouseOk

`func (o *EntitlementsV2) GetProGearWeightSpouseOk() (*int32, bool)`

GetProGearWeightSpouseOk returns a tuple with the ProGearWeightSpouse field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProGearWeightSpouse

`func (o *EntitlementsV2) SetProGearWeightSpouse(v int32)`

SetProGearWeightSpouse sets ProGearWeightSpouse field to given value.

### HasProGearWeightSpouse

`func (o *EntitlementsV2) HasProGearWeightSpouse() bool`

HasProGearWeightSpouse returns a boolean if a field has been set.

### GetRequiredMedicalEquipmentWeight

`func (o *EntitlementsV2) GetRequiredMedicalEquipmentWeight() int32`

GetRequiredMedicalEquipmentWeight returns the RequiredMedicalEquipmentWeight field if non-nil, zero value otherwise.

### GetRequiredMedicalEquipmentWeightOk

`func (o *EntitlementsV2) GetRequiredMedicalEquipmentWeightOk() (*int32, bool)`

GetRequiredMedicalEquipmentWeightOk returns a tuple with the RequiredMedicalEquipmentWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredMedicalEquipmentWeight

`func (o *EntitlementsV2) SetRequiredMedicalEquipmentWeight(v int32)`

SetRequiredMedicalEquipmentWeight sets RequiredMedicalEquipmentWeight field to given value.

### HasRequiredMedicalEquipmentWeight

`func (o *EntitlementsV2) HasRequiredMedicalEquipmentWeight() bool`

HasRequiredMedicalEquipmentWeight returns a boolean if a field has been set.

### GetOrganizationalClothingAndIndividualEquipment

`func (o *EntitlementsV2) GetOrganizationalClothingAndIndividualEquipment() bool`

GetOrganizationalClothingAndIndividualEquipment returns the OrganizationalClothingAndIndividualEquipment field if non-nil, zero value otherwise.

### GetOrganizationalClothingAndIndividualEquipmentOk

`func (o *EntitlementsV2) GetOrganizationalClothingAndIndividualEquipmentOk() (*bool, bool)`

GetOrganizationalClothingAndIndividualEquipmentOk returns a tuple with the OrganizationalClothingAndIndividualEquipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationalClothingAndIndividualEquipment

`func (o *EntitlementsV2) SetOrganizationalClothingAndIndividualEquipment(v bool)`

SetOrganizationalClothingAndIndividualEquipment sets OrganizationalClothingAndIndividualEquipment field to given value.

### HasOrganizationalClothingAndIndividualEquipment

`func (o *EntitlementsV2) HasOrganizationalClothingAndIndividualEquipment() bool`

HasOrganizationalClothingAndIndividualEquipment returns a boolean if a field has been set.

### GetStorageInTransit

`func (o *EntitlementsV2) GetStorageInTransit() int32`

GetStorageInTransit returns the StorageInTransit field if non-nil, zero value otherwise.

### GetStorageInTransitOk

`func (o *EntitlementsV2) GetStorageInTransitOk() (*int32, bool)`

GetStorageInTransitOk returns a tuple with the StorageInTransit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageInTransit

`func (o *EntitlementsV2) SetStorageInTransit(v int32)`

SetStorageInTransit sets StorageInTransit field to given value.

### HasStorageInTransit

`func (o *EntitlementsV2) HasStorageInTransit() bool`

HasStorageInTransit returns a boolean if a field has been set.

### GetTotalWeight

`func (o *EntitlementsV2) GetTotalWeight() int32`

GetTotalWeight returns the TotalWeight field if non-nil, zero value otherwise.

### GetTotalWeightOk

`func (o *EntitlementsV2) GetTotalWeightOk() (*int32, bool)`

GetTotalWeightOk returns a tuple with the TotalWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalWeight

`func (o *EntitlementsV2) SetTotalWeight(v int32)`

SetTotalWeight sets TotalWeight field to given value.

### HasTotalWeight

`func (o *EntitlementsV2) HasTotalWeight() bool`

HasTotalWeight returns a boolean if a field has been set.

### GetTotalDependents

`func (o *EntitlementsV2) GetTotalDependents() int32`

GetTotalDependents returns the TotalDependents field if non-nil, zero value otherwise.

### GetTotalDependentsOk

`func (o *EntitlementsV2) GetTotalDependentsOk() (*int32, bool)`

GetTotalDependentsOk returns a tuple with the TotalDependents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalDependents

`func (o *EntitlementsV2) SetTotalDependents(v int32)`

SetTotalDependents sets TotalDependents field to given value.

### HasTotalDependents

`func (o *EntitlementsV2) HasTotalDependents() bool`

HasTotalDependents returns a boolean if a field has been set.

### GetETag

`func (o *EntitlementsV2) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *EntitlementsV2) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *EntitlementsV2) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *EntitlementsV2) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


