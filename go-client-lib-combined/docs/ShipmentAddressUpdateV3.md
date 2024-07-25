# ShipmentAddressUpdateV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ContractorRemarks** | **string** | The reason there is an address change. | [readonly] 
**OfficeRemarks** | Pointer to **NullableString** | The TOO comment on approval or rejection. | [optional] 
**Status** | [**ShipmentAddressUpdateStatusV3V3**](ShipmentAddressUpdateStatusV3.md) |  | 
**ShipmentID** | **string** |  | [readonly] 
**OriginalAddress** | [**AddressV3V3**](AddressV3.md) |  | 
**NewAddress** | [**AddressV3V3**](AddressV3.md) |  | 
**SitOriginalAddress** | Pointer to [**AddressV3V3**](AddressV3.md) |  | [optional] 
**OldSitDistanceBetween** | Pointer to **int32** | The distance between the original SIT address and the previous/old destination address of shipment | [optional] 
**NewSitDistanceBetween** | Pointer to **int32** | The distance between the original SIT address and requested new destination address of shipment | [optional] 

## Methods

### NewShipmentAddressUpdateV3

`func NewShipmentAddressUpdateV3(id string, contractorRemarks string, status ShipmentAddressUpdateStatusV3V3, shipmentID string, originalAddress AddressV3V3, newAddress AddressV3V3, ) *ShipmentAddressUpdateV3`

NewShipmentAddressUpdateV3 instantiates a new ShipmentAddressUpdateV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentAddressUpdateV3WithDefaults

`func NewShipmentAddressUpdateV3WithDefaults() *ShipmentAddressUpdateV3`

NewShipmentAddressUpdateV3WithDefaults instantiates a new ShipmentAddressUpdateV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShipmentAddressUpdateV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipmentAddressUpdateV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipmentAddressUpdateV3) SetId(v string)`

SetId sets Id field to given value.


### GetContractorRemarks

`func (o *ShipmentAddressUpdateV3) GetContractorRemarks() string`

GetContractorRemarks returns the ContractorRemarks field if non-nil, zero value otherwise.

### GetContractorRemarksOk

`func (o *ShipmentAddressUpdateV3) GetContractorRemarksOk() (*string, bool)`

GetContractorRemarksOk returns a tuple with the ContractorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorRemarks

`func (o *ShipmentAddressUpdateV3) SetContractorRemarks(v string)`

SetContractorRemarks sets ContractorRemarks field to given value.


### GetOfficeRemarks

`func (o *ShipmentAddressUpdateV3) GetOfficeRemarks() string`

GetOfficeRemarks returns the OfficeRemarks field if non-nil, zero value otherwise.

### GetOfficeRemarksOk

`func (o *ShipmentAddressUpdateV3) GetOfficeRemarksOk() (*string, bool)`

GetOfficeRemarksOk returns a tuple with the OfficeRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeRemarks

`func (o *ShipmentAddressUpdateV3) SetOfficeRemarks(v string)`

SetOfficeRemarks sets OfficeRemarks field to given value.

### HasOfficeRemarks

`func (o *ShipmentAddressUpdateV3) HasOfficeRemarks() bool`

HasOfficeRemarks returns a boolean if a field has been set.

### SetOfficeRemarksNil

`func (o *ShipmentAddressUpdateV3) SetOfficeRemarksNil(b bool)`

 SetOfficeRemarksNil sets the value for OfficeRemarks to be an explicit nil

### UnsetOfficeRemarks
`func (o *ShipmentAddressUpdateV3) UnsetOfficeRemarks()`

UnsetOfficeRemarks ensures that no value is present for OfficeRemarks, not even an explicit nil
### GetStatus

`func (o *ShipmentAddressUpdateV3) GetStatus() ShipmentAddressUpdateStatusV3V3`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShipmentAddressUpdateV3) GetStatusOk() (*ShipmentAddressUpdateStatusV3V3, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShipmentAddressUpdateV3) SetStatus(v ShipmentAddressUpdateStatusV3V3)`

SetStatus sets Status field to given value.


### GetShipmentID

`func (o *ShipmentAddressUpdateV3) GetShipmentID() string`

GetShipmentID returns the ShipmentID field if non-nil, zero value otherwise.

### GetShipmentIDOk

`func (o *ShipmentAddressUpdateV3) GetShipmentIDOk() (*string, bool)`

GetShipmentIDOk returns a tuple with the ShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentID

`func (o *ShipmentAddressUpdateV3) SetShipmentID(v string)`

SetShipmentID sets ShipmentID field to given value.


### GetOriginalAddress

`func (o *ShipmentAddressUpdateV3) GetOriginalAddress() AddressV3V3`

GetOriginalAddress returns the OriginalAddress field if non-nil, zero value otherwise.

### GetOriginalAddressOk

`func (o *ShipmentAddressUpdateV3) GetOriginalAddressOk() (*AddressV3V3, bool)`

GetOriginalAddressOk returns a tuple with the OriginalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAddress

`func (o *ShipmentAddressUpdateV3) SetOriginalAddress(v AddressV3V3)`

SetOriginalAddress sets OriginalAddress field to given value.


### GetNewAddress

`func (o *ShipmentAddressUpdateV3) GetNewAddress() AddressV3V3`

GetNewAddress returns the NewAddress field if non-nil, zero value otherwise.

### GetNewAddressOk

`func (o *ShipmentAddressUpdateV3) GetNewAddressOk() (*AddressV3V3, bool)`

GetNewAddressOk returns a tuple with the NewAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAddress

`func (o *ShipmentAddressUpdateV3) SetNewAddress(v AddressV3V3)`

SetNewAddress sets NewAddress field to given value.


### GetSitOriginalAddress

`func (o *ShipmentAddressUpdateV3) GetSitOriginalAddress() AddressV3V3`

GetSitOriginalAddress returns the SitOriginalAddress field if non-nil, zero value otherwise.

### GetSitOriginalAddressOk

`func (o *ShipmentAddressUpdateV3) GetSitOriginalAddressOk() (*AddressV3V3, bool)`

GetSitOriginalAddressOk returns a tuple with the SitOriginalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitOriginalAddress

`func (o *ShipmentAddressUpdateV3) SetSitOriginalAddress(v AddressV3V3)`

SetSitOriginalAddress sets SitOriginalAddress field to given value.

### HasSitOriginalAddress

`func (o *ShipmentAddressUpdateV3) HasSitOriginalAddress() bool`

HasSitOriginalAddress returns a boolean if a field has been set.

### GetOldSitDistanceBetween

`func (o *ShipmentAddressUpdateV3) GetOldSitDistanceBetween() int32`

GetOldSitDistanceBetween returns the OldSitDistanceBetween field if non-nil, zero value otherwise.

### GetOldSitDistanceBetweenOk

`func (o *ShipmentAddressUpdateV3) GetOldSitDistanceBetweenOk() (*int32, bool)`

GetOldSitDistanceBetweenOk returns a tuple with the OldSitDistanceBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSitDistanceBetween

`func (o *ShipmentAddressUpdateV3) SetOldSitDistanceBetween(v int32)`

SetOldSitDistanceBetween sets OldSitDistanceBetween field to given value.

### HasOldSitDistanceBetween

`func (o *ShipmentAddressUpdateV3) HasOldSitDistanceBetween() bool`

HasOldSitDistanceBetween returns a boolean if a field has been set.

### GetNewSitDistanceBetween

`func (o *ShipmentAddressUpdateV3) GetNewSitDistanceBetween() int32`

GetNewSitDistanceBetween returns the NewSitDistanceBetween field if non-nil, zero value otherwise.

### GetNewSitDistanceBetweenOk

`func (o *ShipmentAddressUpdateV3) GetNewSitDistanceBetweenOk() (*int32, bool)`

GetNewSitDistanceBetweenOk returns a tuple with the NewSitDistanceBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSitDistanceBetween

`func (o *ShipmentAddressUpdateV3) SetNewSitDistanceBetween(v int32)`

SetNewSitDistanceBetween sets NewSitDistanceBetween field to given value.

### HasNewSitDistanceBetween

`func (o *ShipmentAddressUpdateV3) HasNewSitDistanceBetween() bool`

HasNewSitDistanceBetween returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


