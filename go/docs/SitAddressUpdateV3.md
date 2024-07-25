# SitAddressUpdateV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**MtoServiceItemId** | Pointer to **string** |  | [optional] [readonly] 
**NewAddressId** | Pointer to **string** |  | [optional] [readonly] 
**NewAddress** | Pointer to [**AddressV3V3**](AddressV3.md) |  | [optional] 
**OldAddressId** | Pointer to **string** |  | [optional] [readonly] 
**OldAddress** | Pointer to [**AddressV3V3**](AddressV3.md) |  | [optional] 
**Status** | Pointer to [**SitAddressUpdateStatusV3V3**](SitAddressUpdateStatusV3.md) |  | [optional] 
**Distance** | Pointer to **int32** |  | [optional] [readonly] 
**ContractorRemarks** | Pointer to **NullableString** |  | [optional] 
**OfficeRemarks** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ETag** | Pointer to **string** | A hash unique to this shipment that should be used as the \&quot;If-Match\&quot; header for any updates. | [optional] [readonly] 

## Methods

### NewSitAddressUpdateV3

`func NewSitAddressUpdateV3() *SitAddressUpdateV3`

NewSitAddressUpdateV3 instantiates a new SitAddressUpdateV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSitAddressUpdateV3WithDefaults

`func NewSitAddressUpdateV3WithDefaults() *SitAddressUpdateV3`

NewSitAddressUpdateV3WithDefaults instantiates a new SitAddressUpdateV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SitAddressUpdateV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SitAddressUpdateV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SitAddressUpdateV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SitAddressUpdateV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMtoServiceItemId

`func (o *SitAddressUpdateV3) GetMtoServiceItemId() string`

GetMtoServiceItemId returns the MtoServiceItemId field if non-nil, zero value otherwise.

### GetMtoServiceItemIdOk

`func (o *SitAddressUpdateV3) GetMtoServiceItemIdOk() (*string, bool)`

GetMtoServiceItemIdOk returns a tuple with the MtoServiceItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItemId

`func (o *SitAddressUpdateV3) SetMtoServiceItemId(v string)`

SetMtoServiceItemId sets MtoServiceItemId field to given value.

### HasMtoServiceItemId

`func (o *SitAddressUpdateV3) HasMtoServiceItemId() bool`

HasMtoServiceItemId returns a boolean if a field has been set.

### GetNewAddressId

`func (o *SitAddressUpdateV3) GetNewAddressId() string`

GetNewAddressId returns the NewAddressId field if non-nil, zero value otherwise.

### GetNewAddressIdOk

`func (o *SitAddressUpdateV3) GetNewAddressIdOk() (*string, bool)`

GetNewAddressIdOk returns a tuple with the NewAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAddressId

`func (o *SitAddressUpdateV3) SetNewAddressId(v string)`

SetNewAddressId sets NewAddressId field to given value.

### HasNewAddressId

`func (o *SitAddressUpdateV3) HasNewAddressId() bool`

HasNewAddressId returns a boolean if a field has been set.

### GetNewAddress

`func (o *SitAddressUpdateV3) GetNewAddress() AddressV3V3`

GetNewAddress returns the NewAddress field if non-nil, zero value otherwise.

### GetNewAddressOk

`func (o *SitAddressUpdateV3) GetNewAddressOk() (*AddressV3V3, bool)`

GetNewAddressOk returns a tuple with the NewAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAddress

`func (o *SitAddressUpdateV3) SetNewAddress(v AddressV3V3)`

SetNewAddress sets NewAddress field to given value.

### HasNewAddress

`func (o *SitAddressUpdateV3) HasNewAddress() bool`

HasNewAddress returns a boolean if a field has been set.

### GetOldAddressId

`func (o *SitAddressUpdateV3) GetOldAddressId() string`

GetOldAddressId returns the OldAddressId field if non-nil, zero value otherwise.

### GetOldAddressIdOk

`func (o *SitAddressUpdateV3) GetOldAddressIdOk() (*string, bool)`

GetOldAddressIdOk returns a tuple with the OldAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldAddressId

`func (o *SitAddressUpdateV3) SetOldAddressId(v string)`

SetOldAddressId sets OldAddressId field to given value.

### HasOldAddressId

`func (o *SitAddressUpdateV3) HasOldAddressId() bool`

HasOldAddressId returns a boolean if a field has been set.

### GetOldAddress

`func (o *SitAddressUpdateV3) GetOldAddress() AddressV3V3`

GetOldAddress returns the OldAddress field if non-nil, zero value otherwise.

### GetOldAddressOk

`func (o *SitAddressUpdateV3) GetOldAddressOk() (*AddressV3V3, bool)`

GetOldAddressOk returns a tuple with the OldAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldAddress

`func (o *SitAddressUpdateV3) SetOldAddress(v AddressV3V3)`

SetOldAddress sets OldAddress field to given value.

### HasOldAddress

`func (o *SitAddressUpdateV3) HasOldAddress() bool`

HasOldAddress returns a boolean if a field has been set.

### GetStatus

`func (o *SitAddressUpdateV3) GetStatus() SitAddressUpdateStatusV3V3`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SitAddressUpdateV3) GetStatusOk() (*SitAddressUpdateStatusV3V3, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SitAddressUpdateV3) SetStatus(v SitAddressUpdateStatusV3V3)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SitAddressUpdateV3) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDistance

`func (o *SitAddressUpdateV3) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *SitAddressUpdateV3) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *SitAddressUpdateV3) SetDistance(v int32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *SitAddressUpdateV3) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetContractorRemarks

`func (o *SitAddressUpdateV3) GetContractorRemarks() string`

GetContractorRemarks returns the ContractorRemarks field if non-nil, zero value otherwise.

### GetContractorRemarksOk

`func (o *SitAddressUpdateV3) GetContractorRemarksOk() (*string, bool)`

GetContractorRemarksOk returns a tuple with the ContractorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorRemarks

`func (o *SitAddressUpdateV3) SetContractorRemarks(v string)`

SetContractorRemarks sets ContractorRemarks field to given value.

### HasContractorRemarks

`func (o *SitAddressUpdateV3) HasContractorRemarks() bool`

HasContractorRemarks returns a boolean if a field has been set.

### SetContractorRemarksNil

`func (o *SitAddressUpdateV3) SetContractorRemarksNil(b bool)`

 SetContractorRemarksNil sets the value for ContractorRemarks to be an explicit nil

### UnsetContractorRemarks
`func (o *SitAddressUpdateV3) UnsetContractorRemarks()`

UnsetContractorRemarks ensures that no value is present for ContractorRemarks, not even an explicit nil
### GetOfficeRemarks

`func (o *SitAddressUpdateV3) GetOfficeRemarks() string`

GetOfficeRemarks returns the OfficeRemarks field if non-nil, zero value otherwise.

### GetOfficeRemarksOk

`func (o *SitAddressUpdateV3) GetOfficeRemarksOk() (*string, bool)`

GetOfficeRemarksOk returns a tuple with the OfficeRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeRemarks

`func (o *SitAddressUpdateV3) SetOfficeRemarks(v string)`

SetOfficeRemarks sets OfficeRemarks field to given value.

### HasOfficeRemarks

`func (o *SitAddressUpdateV3) HasOfficeRemarks() bool`

HasOfficeRemarks returns a boolean if a field has been set.

### SetOfficeRemarksNil

`func (o *SitAddressUpdateV3) SetOfficeRemarksNil(b bool)`

 SetOfficeRemarksNil sets the value for OfficeRemarks to be an explicit nil

### UnsetOfficeRemarks
`func (o *SitAddressUpdateV3) UnsetOfficeRemarks()`

UnsetOfficeRemarks ensures that no value is present for OfficeRemarks, not even an explicit nil
### GetCreatedAt

`func (o *SitAddressUpdateV3) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SitAddressUpdateV3) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SitAddressUpdateV3) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SitAddressUpdateV3) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SitAddressUpdateV3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SitAddressUpdateV3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SitAddressUpdateV3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SitAddressUpdateV3) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetETag

`func (o *SitAddressUpdateV3) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *SitAddressUpdateV3) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *SitAddressUpdateV3) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *SitAddressUpdateV3) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


