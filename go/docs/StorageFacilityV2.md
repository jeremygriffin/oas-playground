# StorageFacilityV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**FacilityName** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**AddressV2V2**](AddressV2.md) |  | [optional] 
**LotNumber** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewStorageFacilityV2

`func NewStorageFacilityV2() *StorageFacilityV2`

NewStorageFacilityV2 instantiates a new StorageFacilityV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStorageFacilityV2WithDefaults

`func NewStorageFacilityV2WithDefaults() *StorageFacilityV2`

NewStorageFacilityV2WithDefaults instantiates a new StorageFacilityV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *StorageFacilityV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StorageFacilityV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StorageFacilityV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StorageFacilityV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFacilityName

`func (o *StorageFacilityV2) GetFacilityName() string`

GetFacilityName returns the FacilityName field if non-nil, zero value otherwise.

### GetFacilityNameOk

`func (o *StorageFacilityV2) GetFacilityNameOk() (*string, bool)`

GetFacilityNameOk returns a tuple with the FacilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityName

`func (o *StorageFacilityV2) SetFacilityName(v string)`

SetFacilityName sets FacilityName field to given value.

### HasFacilityName

`func (o *StorageFacilityV2) HasFacilityName() bool`

HasFacilityName returns a boolean if a field has been set.

### GetAddress

`func (o *StorageFacilityV2) GetAddress() AddressV2V2`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *StorageFacilityV2) GetAddressOk() (*AddressV2V2, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *StorageFacilityV2) SetAddress(v AddressV2V2)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *StorageFacilityV2) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLotNumber

`func (o *StorageFacilityV2) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *StorageFacilityV2) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *StorageFacilityV2) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *StorageFacilityV2) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### SetLotNumberNil

`func (o *StorageFacilityV2) SetLotNumberNil(b bool)`

 SetLotNumberNil sets the value for LotNumber to be an explicit nil

### UnsetLotNumber
`func (o *StorageFacilityV2) UnsetLotNumber()`

UnsetLotNumber ensures that no value is present for LotNumber, not even an explicit nil
### GetPhone

`func (o *StorageFacilityV2) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *StorageFacilityV2) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *StorageFacilityV2) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *StorageFacilityV2) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *StorageFacilityV2) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *StorageFacilityV2) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *StorageFacilityV2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *StorageFacilityV2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *StorageFacilityV2) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *StorageFacilityV2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *StorageFacilityV2) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *StorageFacilityV2) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetETag

`func (o *StorageFacilityV2) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *StorageFacilityV2) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *StorageFacilityV2) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *StorageFacilityV2) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


