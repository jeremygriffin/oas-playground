# SitAddressUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] [readonly] 
**MtoServiceItemId** | Pointer to **string** |  | [optional] [readonly] 
**NewAddressId** | Pointer to **string** |  | [optional] [readonly] 
**NewAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**OldAddressId** | Pointer to **string** |  | [optional] [readonly] 
**OldAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Status** | Pointer to [**SitAddressUpdateStatus**](SitAddressUpdateStatus.md) |  | [optional] 
**Distance** | Pointer to **int32** |  | [optional] [readonly] 
**ContractorRemarks** | Pointer to **NullableString** |  | [optional] 
**OfficeRemarks** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ETag** | Pointer to **string** | A hash unique to this shipment that should be used as the \&quot;If-Match\&quot; header for any updates. | [optional] [readonly] 

## Methods

### NewSitAddressUpdate

`func NewSitAddressUpdate() *SitAddressUpdate`

NewSitAddressUpdate instantiates a new SitAddressUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSitAddressUpdateWithDefaults

`func NewSitAddressUpdateWithDefaults() *SitAddressUpdate`

NewSitAddressUpdateWithDefaults instantiates a new SitAddressUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SitAddressUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SitAddressUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SitAddressUpdate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SitAddressUpdate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMtoServiceItemId

`func (o *SitAddressUpdate) GetMtoServiceItemId() string`

GetMtoServiceItemId returns the MtoServiceItemId field if non-nil, zero value otherwise.

### GetMtoServiceItemIdOk

`func (o *SitAddressUpdate) GetMtoServiceItemIdOk() (*string, bool)`

GetMtoServiceItemIdOk returns a tuple with the MtoServiceItemId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItemId

`func (o *SitAddressUpdate) SetMtoServiceItemId(v string)`

SetMtoServiceItemId sets MtoServiceItemId field to given value.

### HasMtoServiceItemId

`func (o *SitAddressUpdate) HasMtoServiceItemId() bool`

HasMtoServiceItemId returns a boolean if a field has been set.

### GetNewAddressId

`func (o *SitAddressUpdate) GetNewAddressId() string`

GetNewAddressId returns the NewAddressId field if non-nil, zero value otherwise.

### GetNewAddressIdOk

`func (o *SitAddressUpdate) GetNewAddressIdOk() (*string, bool)`

GetNewAddressIdOk returns a tuple with the NewAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAddressId

`func (o *SitAddressUpdate) SetNewAddressId(v string)`

SetNewAddressId sets NewAddressId field to given value.

### HasNewAddressId

`func (o *SitAddressUpdate) HasNewAddressId() bool`

HasNewAddressId returns a boolean if a field has been set.

### GetNewAddress

`func (o *SitAddressUpdate) GetNewAddress() Address`

GetNewAddress returns the NewAddress field if non-nil, zero value otherwise.

### GetNewAddressOk

`func (o *SitAddressUpdate) GetNewAddressOk() (*Address, bool)`

GetNewAddressOk returns a tuple with the NewAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAddress

`func (o *SitAddressUpdate) SetNewAddress(v Address)`

SetNewAddress sets NewAddress field to given value.

### HasNewAddress

`func (o *SitAddressUpdate) HasNewAddress() bool`

HasNewAddress returns a boolean if a field has been set.

### GetOldAddressId

`func (o *SitAddressUpdate) GetOldAddressId() string`

GetOldAddressId returns the OldAddressId field if non-nil, zero value otherwise.

### GetOldAddressIdOk

`func (o *SitAddressUpdate) GetOldAddressIdOk() (*string, bool)`

GetOldAddressIdOk returns a tuple with the OldAddressId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldAddressId

`func (o *SitAddressUpdate) SetOldAddressId(v string)`

SetOldAddressId sets OldAddressId field to given value.

### HasOldAddressId

`func (o *SitAddressUpdate) HasOldAddressId() bool`

HasOldAddressId returns a boolean if a field has been set.

### GetOldAddress

`func (o *SitAddressUpdate) GetOldAddress() Address`

GetOldAddress returns the OldAddress field if non-nil, zero value otherwise.

### GetOldAddressOk

`func (o *SitAddressUpdate) GetOldAddressOk() (*Address, bool)`

GetOldAddressOk returns a tuple with the OldAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldAddress

`func (o *SitAddressUpdate) SetOldAddress(v Address)`

SetOldAddress sets OldAddress field to given value.

### HasOldAddress

`func (o *SitAddressUpdate) HasOldAddress() bool`

HasOldAddress returns a boolean if a field has been set.

### GetStatus

`func (o *SitAddressUpdate) GetStatus() SitAddressUpdateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SitAddressUpdate) GetStatusOk() (*SitAddressUpdateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SitAddressUpdate) SetStatus(v SitAddressUpdateStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SitAddressUpdate) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDistance

`func (o *SitAddressUpdate) GetDistance() int32`

GetDistance returns the Distance field if non-nil, zero value otherwise.

### GetDistanceOk

`func (o *SitAddressUpdate) GetDistanceOk() (*int32, bool)`

GetDistanceOk returns a tuple with the Distance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDistance

`func (o *SitAddressUpdate) SetDistance(v int32)`

SetDistance sets Distance field to given value.

### HasDistance

`func (o *SitAddressUpdate) HasDistance() bool`

HasDistance returns a boolean if a field has been set.

### GetContractorRemarks

`func (o *SitAddressUpdate) GetContractorRemarks() string`

GetContractorRemarks returns the ContractorRemarks field if non-nil, zero value otherwise.

### GetContractorRemarksOk

`func (o *SitAddressUpdate) GetContractorRemarksOk() (*string, bool)`

GetContractorRemarksOk returns a tuple with the ContractorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorRemarks

`func (o *SitAddressUpdate) SetContractorRemarks(v string)`

SetContractorRemarks sets ContractorRemarks field to given value.

### HasContractorRemarks

`func (o *SitAddressUpdate) HasContractorRemarks() bool`

HasContractorRemarks returns a boolean if a field has been set.

### SetContractorRemarksNil

`func (o *SitAddressUpdate) SetContractorRemarksNil(b bool)`

 SetContractorRemarksNil sets the value for ContractorRemarks to be an explicit nil

### UnsetContractorRemarks
`func (o *SitAddressUpdate) UnsetContractorRemarks()`

UnsetContractorRemarks ensures that no value is present for ContractorRemarks, not even an explicit nil
### GetOfficeRemarks

`func (o *SitAddressUpdate) GetOfficeRemarks() string`

GetOfficeRemarks returns the OfficeRemarks field if non-nil, zero value otherwise.

### GetOfficeRemarksOk

`func (o *SitAddressUpdate) GetOfficeRemarksOk() (*string, bool)`

GetOfficeRemarksOk returns a tuple with the OfficeRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeRemarks

`func (o *SitAddressUpdate) SetOfficeRemarks(v string)`

SetOfficeRemarks sets OfficeRemarks field to given value.

### HasOfficeRemarks

`func (o *SitAddressUpdate) HasOfficeRemarks() bool`

HasOfficeRemarks returns a boolean if a field has been set.

### SetOfficeRemarksNil

`func (o *SitAddressUpdate) SetOfficeRemarksNil(b bool)`

 SetOfficeRemarksNil sets the value for OfficeRemarks to be an explicit nil

### UnsetOfficeRemarks
`func (o *SitAddressUpdate) UnsetOfficeRemarks()`

UnsetOfficeRemarks ensures that no value is present for OfficeRemarks, not even an explicit nil
### GetCreatedAt

`func (o *SitAddressUpdate) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SitAddressUpdate) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SitAddressUpdate) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SitAddressUpdate) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SitAddressUpdate) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SitAddressUpdate) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SitAddressUpdate) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SitAddressUpdate) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetETag

`func (o *SitAddressUpdate) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *SitAddressUpdate) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *SitAddressUpdate) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *SitAddressUpdate) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


