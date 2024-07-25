# UpdateMTOShipmentStorageFacilityV2

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

### NewUpdateMTOShipmentStorageFacilityV2

`func NewUpdateMTOShipmentStorageFacilityV2() *UpdateMTOShipmentStorageFacilityV2`

NewUpdateMTOShipmentStorageFacilityV2 instantiates a new UpdateMTOShipmentStorageFacilityV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMTOShipmentStorageFacilityV2WithDefaults

`func NewUpdateMTOShipmentStorageFacilityV2WithDefaults() *UpdateMTOShipmentStorageFacilityV2`

NewUpdateMTOShipmentStorageFacilityV2WithDefaults instantiates a new UpdateMTOShipmentStorageFacilityV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateMTOShipmentStorageFacilityV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateMTOShipmentStorageFacilityV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateMTOShipmentStorageFacilityV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateMTOShipmentStorageFacilityV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetFacilityName

`func (o *UpdateMTOShipmentStorageFacilityV2) GetFacilityName() string`

GetFacilityName returns the FacilityName field if non-nil, zero value otherwise.

### GetFacilityNameOk

`func (o *UpdateMTOShipmentStorageFacilityV2) GetFacilityNameOk() (*string, bool)`

GetFacilityNameOk returns a tuple with the FacilityName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFacilityName

`func (o *UpdateMTOShipmentStorageFacilityV2) SetFacilityName(v string)`

SetFacilityName sets FacilityName field to given value.

### HasFacilityName

`func (o *UpdateMTOShipmentStorageFacilityV2) HasFacilityName() bool`

HasFacilityName returns a boolean if a field has been set.

### GetAddress

`func (o *UpdateMTOShipmentStorageFacilityV2) GetAddress() AddressV2V2`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UpdateMTOShipmentStorageFacilityV2) GetAddressOk() (*AddressV2V2, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UpdateMTOShipmentStorageFacilityV2) SetAddress(v AddressV2V2)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UpdateMTOShipmentStorageFacilityV2) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetLotNumber

`func (o *UpdateMTOShipmentStorageFacilityV2) GetLotNumber() string`

GetLotNumber returns the LotNumber field if non-nil, zero value otherwise.

### GetLotNumberOk

`func (o *UpdateMTOShipmentStorageFacilityV2) GetLotNumberOk() (*string, bool)`

GetLotNumberOk returns a tuple with the LotNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLotNumber

`func (o *UpdateMTOShipmentStorageFacilityV2) SetLotNumber(v string)`

SetLotNumber sets LotNumber field to given value.

### HasLotNumber

`func (o *UpdateMTOShipmentStorageFacilityV2) HasLotNumber() bool`

HasLotNumber returns a boolean if a field has been set.

### SetLotNumberNil

`func (o *UpdateMTOShipmentStorageFacilityV2) SetLotNumberNil(b bool)`

 SetLotNumberNil sets the value for LotNumber to be an explicit nil

### UnsetLotNumber
`func (o *UpdateMTOShipmentStorageFacilityV2) UnsetLotNumber()`

UnsetLotNumber ensures that no value is present for LotNumber, not even an explicit nil
### GetPhone

`func (o *UpdateMTOShipmentStorageFacilityV2) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *UpdateMTOShipmentStorageFacilityV2) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *UpdateMTOShipmentStorageFacilityV2) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *UpdateMTOShipmentStorageFacilityV2) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *UpdateMTOShipmentStorageFacilityV2) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *UpdateMTOShipmentStorageFacilityV2) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetEmail

`func (o *UpdateMTOShipmentStorageFacilityV2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateMTOShipmentStorageFacilityV2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateMTOShipmentStorageFacilityV2) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateMTOShipmentStorageFacilityV2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *UpdateMTOShipmentStorageFacilityV2) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *UpdateMTOShipmentStorageFacilityV2) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetETag

`func (o *UpdateMTOShipmentStorageFacilityV2) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *UpdateMTOShipmentStorageFacilityV2) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *UpdateMTOShipmentStorageFacilityV2) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *UpdateMTOShipmentStorageFacilityV2) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


