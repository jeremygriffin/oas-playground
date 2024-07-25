# UpdateShipmentDestinationAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewAddress** | [**Address**](Address.md) |  | 
**ContractorRemarks** | **string** | This is the remark the Prime has entered, which would be the reason there is an address change. | 

## Methods

### NewUpdateShipmentDestinationAddress

`func NewUpdateShipmentDestinationAddress(newAddress Address, contractorRemarks string, ) *UpdateShipmentDestinationAddress`

NewUpdateShipmentDestinationAddress instantiates a new UpdateShipmentDestinationAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateShipmentDestinationAddressWithDefaults

`func NewUpdateShipmentDestinationAddressWithDefaults() *UpdateShipmentDestinationAddress`

NewUpdateShipmentDestinationAddressWithDefaults instantiates a new UpdateShipmentDestinationAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewAddress

`func (o *UpdateShipmentDestinationAddress) GetNewAddress() Address`

GetNewAddress returns the NewAddress field if non-nil, zero value otherwise.

### GetNewAddressOk

`func (o *UpdateShipmentDestinationAddress) GetNewAddressOk() (*Address, bool)`

GetNewAddressOk returns a tuple with the NewAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAddress

`func (o *UpdateShipmentDestinationAddress) SetNewAddress(v Address)`

SetNewAddress sets NewAddress field to given value.


### GetContractorRemarks

`func (o *UpdateShipmentDestinationAddress) GetContractorRemarks() string`

GetContractorRemarks returns the ContractorRemarks field if non-nil, zero value otherwise.

### GetContractorRemarksOk

`func (o *UpdateShipmentDestinationAddress) GetContractorRemarksOk() (*string, bool)`

GetContractorRemarksOk returns a tuple with the ContractorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorRemarks

`func (o *UpdateShipmentDestinationAddress) SetContractorRemarks(v string)`

SetContractorRemarks sets ContractorRemarks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


