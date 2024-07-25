# CreateSITAddressUpdateRequestV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewAddress** | Pointer to [**AddressV2V2**](AddressV2.md) |  | [optional] 
**ContractorRemarks** | **string** |  | 
**MtoServiceItemID** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateSITAddressUpdateRequestV2

`func NewCreateSITAddressUpdateRequestV2(contractorRemarks string, ) *CreateSITAddressUpdateRequestV2`

NewCreateSITAddressUpdateRequestV2 instantiates a new CreateSITAddressUpdateRequestV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSITAddressUpdateRequestV2WithDefaults

`func NewCreateSITAddressUpdateRequestV2WithDefaults() *CreateSITAddressUpdateRequestV2`

NewCreateSITAddressUpdateRequestV2WithDefaults instantiates a new CreateSITAddressUpdateRequestV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewAddress

`func (o *CreateSITAddressUpdateRequestV2) GetNewAddress() AddressV2V2`

GetNewAddress returns the NewAddress field if non-nil, zero value otherwise.

### GetNewAddressOk

`func (o *CreateSITAddressUpdateRequestV2) GetNewAddressOk() (*AddressV2V2, bool)`

GetNewAddressOk returns a tuple with the NewAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAddress

`func (o *CreateSITAddressUpdateRequestV2) SetNewAddress(v AddressV2V2)`

SetNewAddress sets NewAddress field to given value.

### HasNewAddress

`func (o *CreateSITAddressUpdateRequestV2) HasNewAddress() bool`

HasNewAddress returns a boolean if a field has been set.

### GetContractorRemarks

`func (o *CreateSITAddressUpdateRequestV2) GetContractorRemarks() string`

GetContractorRemarks returns the ContractorRemarks field if non-nil, zero value otherwise.

### GetContractorRemarksOk

`func (o *CreateSITAddressUpdateRequestV2) GetContractorRemarksOk() (*string, bool)`

GetContractorRemarksOk returns a tuple with the ContractorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorRemarks

`func (o *CreateSITAddressUpdateRequestV2) SetContractorRemarks(v string)`

SetContractorRemarks sets ContractorRemarks field to given value.


### GetMtoServiceItemID

`func (o *CreateSITAddressUpdateRequestV2) GetMtoServiceItemID() string`

GetMtoServiceItemID returns the MtoServiceItemID field if non-nil, zero value otherwise.

### GetMtoServiceItemIDOk

`func (o *CreateSITAddressUpdateRequestV2) GetMtoServiceItemIDOk() (*string, bool)`

GetMtoServiceItemIDOk returns a tuple with the MtoServiceItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItemID

`func (o *CreateSITAddressUpdateRequestV2) SetMtoServiceItemID(v string)`

SetMtoServiceItemID sets MtoServiceItemID field to given value.

### HasMtoServiceItemID

`func (o *CreateSITAddressUpdateRequestV2) HasMtoServiceItemID() bool`

HasMtoServiceItemID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


