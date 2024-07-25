# CreateSITAddressUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**ContractorRemarks** | **string** |  | 
**MtoServiceItemID** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateSITAddressUpdateRequest

`func NewCreateSITAddressUpdateRequest(contractorRemarks string, ) *CreateSITAddressUpdateRequest`

NewCreateSITAddressUpdateRequest instantiates a new CreateSITAddressUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSITAddressUpdateRequestWithDefaults

`func NewCreateSITAddressUpdateRequestWithDefaults() *CreateSITAddressUpdateRequest`

NewCreateSITAddressUpdateRequestWithDefaults instantiates a new CreateSITAddressUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewAddress

`func (o *CreateSITAddressUpdateRequest) GetNewAddress() Address`

GetNewAddress returns the NewAddress field if non-nil, zero value otherwise.

### GetNewAddressOk

`func (o *CreateSITAddressUpdateRequest) GetNewAddressOk() (*Address, bool)`

GetNewAddressOk returns a tuple with the NewAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAddress

`func (o *CreateSITAddressUpdateRequest) SetNewAddress(v Address)`

SetNewAddress sets NewAddress field to given value.

### HasNewAddress

`func (o *CreateSITAddressUpdateRequest) HasNewAddress() bool`

HasNewAddress returns a boolean if a field has been set.

### GetContractorRemarks

`func (o *CreateSITAddressUpdateRequest) GetContractorRemarks() string`

GetContractorRemarks returns the ContractorRemarks field if non-nil, zero value otherwise.

### GetContractorRemarksOk

`func (o *CreateSITAddressUpdateRequest) GetContractorRemarksOk() (*string, bool)`

GetContractorRemarksOk returns a tuple with the ContractorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorRemarks

`func (o *CreateSITAddressUpdateRequest) SetContractorRemarks(v string)`

SetContractorRemarks sets ContractorRemarks field to given value.


### GetMtoServiceItemID

`func (o *CreateSITAddressUpdateRequest) GetMtoServiceItemID() string`

GetMtoServiceItemID returns the MtoServiceItemID field if non-nil, zero value otherwise.

### GetMtoServiceItemIDOk

`func (o *CreateSITAddressUpdateRequest) GetMtoServiceItemIDOk() (*string, bool)`

GetMtoServiceItemIDOk returns a tuple with the MtoServiceItemID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItemID

`func (o *CreateSITAddressUpdateRequest) SetMtoServiceItemID(v string)`

SetMtoServiceItemID sets MtoServiceItemID field to given value.

### HasMtoServiceItemID

`func (o *CreateSITAddressUpdateRequest) HasMtoServiceItemID() bool`

HasMtoServiceItemID returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


