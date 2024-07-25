# ShipmentAddressUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | [readonly] 
**ContractorRemarks** | **string** | The reason there is an address change. | [readonly] 
**OfficeRemarks** | Pointer to **NullableString** | The TOO comment on approval or rejection. | [optional] 
**Status** | [**ShipmentAddressUpdateStatus**](ShipmentAddressUpdateStatus.md) |  | 
**ShipmentID** | **string** |  | [readonly] 
**OriginalAddress** | [**Address**](Address.md) |  | 
**NewAddress** | [**Address**](Address.md) |  | 
**SitOriginalAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**OldSitDistanceBetween** | Pointer to **int32** | The distance between the original SIT address and the previous/old destination address of shipment | [optional] 
**NewSitDistanceBetween** | Pointer to **int32** | The distance between the original SIT address and requested new destination address of shipment | [optional] 

## Methods

### NewShipmentAddressUpdate

`func NewShipmentAddressUpdate(id string, contractorRemarks string, status ShipmentAddressUpdateStatus, shipmentID string, originalAddress Address, newAddress Address, ) *ShipmentAddressUpdate`

NewShipmentAddressUpdate instantiates a new ShipmentAddressUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewShipmentAddressUpdateWithDefaults

`func NewShipmentAddressUpdateWithDefaults() *ShipmentAddressUpdate`

NewShipmentAddressUpdateWithDefaults instantiates a new ShipmentAddressUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ShipmentAddressUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ShipmentAddressUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ShipmentAddressUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetContractorRemarks

`func (o *ShipmentAddressUpdate) GetContractorRemarks() string`

GetContractorRemarks returns the ContractorRemarks field if non-nil, zero value otherwise.

### GetContractorRemarksOk

`func (o *ShipmentAddressUpdate) GetContractorRemarksOk() (*string, bool)`

GetContractorRemarksOk returns a tuple with the ContractorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorRemarks

`func (o *ShipmentAddressUpdate) SetContractorRemarks(v string)`

SetContractorRemarks sets ContractorRemarks field to given value.


### GetOfficeRemarks

`func (o *ShipmentAddressUpdate) GetOfficeRemarks() string`

GetOfficeRemarks returns the OfficeRemarks field if non-nil, zero value otherwise.

### GetOfficeRemarksOk

`func (o *ShipmentAddressUpdate) GetOfficeRemarksOk() (*string, bool)`

GetOfficeRemarksOk returns a tuple with the OfficeRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeRemarks

`func (o *ShipmentAddressUpdate) SetOfficeRemarks(v string)`

SetOfficeRemarks sets OfficeRemarks field to given value.

### HasOfficeRemarks

`func (o *ShipmentAddressUpdate) HasOfficeRemarks() bool`

HasOfficeRemarks returns a boolean if a field has been set.

### SetOfficeRemarksNil

`func (o *ShipmentAddressUpdate) SetOfficeRemarksNil(b bool)`

 SetOfficeRemarksNil sets the value for OfficeRemarks to be an explicit nil

### UnsetOfficeRemarks
`func (o *ShipmentAddressUpdate) UnsetOfficeRemarks()`

UnsetOfficeRemarks ensures that no value is present for OfficeRemarks, not even an explicit nil
### GetStatus

`func (o *ShipmentAddressUpdate) GetStatus() ShipmentAddressUpdateStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ShipmentAddressUpdate) GetStatusOk() (*ShipmentAddressUpdateStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ShipmentAddressUpdate) SetStatus(v ShipmentAddressUpdateStatus)`

SetStatus sets Status field to given value.


### GetShipmentID

`func (o *ShipmentAddressUpdate) GetShipmentID() string`

GetShipmentID returns the ShipmentID field if non-nil, zero value otherwise.

### GetShipmentIDOk

`func (o *ShipmentAddressUpdate) GetShipmentIDOk() (*string, bool)`

GetShipmentIDOk returns a tuple with the ShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentID

`func (o *ShipmentAddressUpdate) SetShipmentID(v string)`

SetShipmentID sets ShipmentID field to given value.


### GetOriginalAddress

`func (o *ShipmentAddressUpdate) GetOriginalAddress() Address`

GetOriginalAddress returns the OriginalAddress field if non-nil, zero value otherwise.

### GetOriginalAddressOk

`func (o *ShipmentAddressUpdate) GetOriginalAddressOk() (*Address, bool)`

GetOriginalAddressOk returns a tuple with the OriginalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginalAddress

`func (o *ShipmentAddressUpdate) SetOriginalAddress(v Address)`

SetOriginalAddress sets OriginalAddress field to given value.


### GetNewAddress

`func (o *ShipmentAddressUpdate) GetNewAddress() Address`

GetNewAddress returns the NewAddress field if non-nil, zero value otherwise.

### GetNewAddressOk

`func (o *ShipmentAddressUpdate) GetNewAddressOk() (*Address, bool)`

GetNewAddressOk returns a tuple with the NewAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewAddress

`func (o *ShipmentAddressUpdate) SetNewAddress(v Address)`

SetNewAddress sets NewAddress field to given value.


### GetSitOriginalAddress

`func (o *ShipmentAddressUpdate) GetSitOriginalAddress() Address`

GetSitOriginalAddress returns the SitOriginalAddress field if non-nil, zero value otherwise.

### GetSitOriginalAddressOk

`func (o *ShipmentAddressUpdate) GetSitOriginalAddressOk() (*Address, bool)`

GetSitOriginalAddressOk returns a tuple with the SitOriginalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitOriginalAddress

`func (o *ShipmentAddressUpdate) SetSitOriginalAddress(v Address)`

SetSitOriginalAddress sets SitOriginalAddress field to given value.

### HasSitOriginalAddress

`func (o *ShipmentAddressUpdate) HasSitOriginalAddress() bool`

HasSitOriginalAddress returns a boolean if a field has been set.

### GetOldSitDistanceBetween

`func (o *ShipmentAddressUpdate) GetOldSitDistanceBetween() int32`

GetOldSitDistanceBetween returns the OldSitDistanceBetween field if non-nil, zero value otherwise.

### GetOldSitDistanceBetweenOk

`func (o *ShipmentAddressUpdate) GetOldSitDistanceBetweenOk() (*int32, bool)`

GetOldSitDistanceBetweenOk returns a tuple with the OldSitDistanceBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldSitDistanceBetween

`func (o *ShipmentAddressUpdate) SetOldSitDistanceBetween(v int32)`

SetOldSitDistanceBetween sets OldSitDistanceBetween field to given value.

### HasOldSitDistanceBetween

`func (o *ShipmentAddressUpdate) HasOldSitDistanceBetween() bool`

HasOldSitDistanceBetween returns a boolean if a field has been set.

### GetNewSitDistanceBetween

`func (o *ShipmentAddressUpdate) GetNewSitDistanceBetween() int32`

GetNewSitDistanceBetween returns the NewSitDistanceBetween field if non-nil, zero value otherwise.

### GetNewSitDistanceBetweenOk

`func (o *ShipmentAddressUpdate) GetNewSitDistanceBetweenOk() (*int32, bool)`

GetNewSitDistanceBetweenOk returns a tuple with the NewSitDistanceBetween field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewSitDistanceBetween

`func (o *ShipmentAddressUpdate) SetNewSitDistanceBetween(v int32)`

SetNewSitDistanceBetween sets NewSitDistanceBetween field to given value.

### HasNewSitDistanceBetween

`func (o *ShipmentAddressUpdate) HasNewSitDistanceBetween() bool`

HasNewSitDistanceBetween returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


