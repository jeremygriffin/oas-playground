# CreateSITAddressUpdateRequestV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewAddress** | Pointer to [**AddressV3V3**](AddressV3.md) |  | [optional] 
**ContractorRemarks** | **string** |  | 
**MtoServiceItemID** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateSITAddressUpdateRequestV3

`func NewCreateSITAddressUpdateRequestV3(contractorRemarks string, ) *CreateSITAddressUpdateRequestV3`

NewCreateSITAddressUpdateRequestV3 instantiates a new CreateSITAddressUpdateRequestV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSITAddressUpdateRequestV3WithDefaults

`func NewCreateSITAddressUpdateRequestV3WithDefaults() *CreateSITAddressUpdateRequestV3`

NewCreateSITAddressUpdateRequestV3WithDefaults instantiates a new CreateSITAddressUpdateRequestV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewAddress

`func (o *CreateSITAddressUpdateRequestV3) GetNewAddress() AddressV3V3`

GetNewAddress returns the NewAddress field if non-nil, zero value otherwise.

### GetNewAddressOk

`func (o *CreateSITAddressUpdateRequestV3) GetNewAddressOk() (*AddressV3V3, bool)`

GetNewAddressOk returns a tuple with the NewAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAddress

`func (o *CreateSITAddressUpdateRequestV3) SetNewAddress(v AddressV3V3)`

SetNewAddress sets NewAddress field to given value.

### HasNewAddress

`func (o *CreateSITAddressUpdateRequestV3) HasNewAddress() bool`

HasNewAddress returns a boolean if a field has been set.

### GetContractorRemarks

`func (o *CreateSITAddressUpdateRequestV3) GetContractorRemarks() string`

GetContractorRemarks returns the ContractorRemarks field if non-nil, zero value otherwise.

### GetContractorRemarksOk

`func (o *CreateSITAddressUpdateRequestV3) GetContractorRemarksOk() (*string, bool)`

GetContractorRemarksOk returns a tuple with the ContractorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorRemarks

`func (o *CreateSITAddressUpdateRequestV3) SetContractorRemarks(v string)`

SetContractorRemarks sets ContractorRemarks field to given value.


### GetMtoServiceItemID

`func (o *CreateSITAddressUpdateRequestV3) GetMtoServiceItemID() string`

GetMtoServiceItemID returns the MtoServiceItemID field if non-nil, zero value otherwise.

### GetMtoServiceItemIDOk

`func (o *CreateSITAddressUpdateRequestV3) GetMtoServiceItemIDOk() (*string, bool)`

GetMtoServiceItemIDOk returns a tuple with the MtoServiceItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItemID

`func (o *CreateSITAddressUpdateRequestV3) SetMtoServiceItemID(v string)`

SetMtoServiceItemID sets MtoServiceItemID field to given value.

### HasMtoServiceItemID

`func (o *CreateSITAddressUpdateRequestV3) HasMtoServiceItemID() bool`

HasMtoServiceItemID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


