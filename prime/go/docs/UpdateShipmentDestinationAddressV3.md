# UpdateShipmentDestinationAddressV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewAddress** | [**AddressV3V3**](AddressV3.md) |  | 
**ContractorRemarks** | **string** | This is the remark the Prime has entered, which would be the reason there is an address change. | 

## Methods

### NewUpdateShipmentDestinationAddressV3

`func NewUpdateShipmentDestinationAddressV3(newAddress AddressV3V3, contractorRemarks string, ) *UpdateShipmentDestinationAddressV3`

NewUpdateShipmentDestinationAddressV3 instantiates a new UpdateShipmentDestinationAddressV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateShipmentDestinationAddressV3WithDefaults

`func NewUpdateShipmentDestinationAddressV3WithDefaults() *UpdateShipmentDestinationAddressV3`

NewUpdateShipmentDestinationAddressV3WithDefaults instantiates a new UpdateShipmentDestinationAddressV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewAddress

`func (o *UpdateShipmentDestinationAddressV3) GetNewAddress() AddressV3V3`

GetNewAddress returns the NewAddress field if non-nil, zero value otherwise.

### GetNewAddressOk

`func (o *UpdateShipmentDestinationAddressV3) GetNewAddressOk() (*AddressV3V3, bool)`

GetNewAddressOk returns a tuple with the NewAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAddress

`func (o *UpdateShipmentDestinationAddressV3) SetNewAddress(v AddressV3V3)`

SetNewAddress sets NewAddress field to given value.


### GetContractorRemarks

`func (o *UpdateShipmentDestinationAddressV3) GetContractorRemarks() string`

GetContractorRemarks returns the ContractorRemarks field if non-nil, zero value otherwise.

### GetContractorRemarksOk

`func (o *UpdateShipmentDestinationAddressV3) GetContractorRemarksOk() (*string, bool)`

GetContractorRemarksOk returns a tuple with the ContractorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorRemarks

`func (o *UpdateShipmentDestinationAddressV3) SetContractorRemarks(v string)`

SetContractorRemarks sets ContractorRemarks field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


