# CreateMTOShipmentV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveTaskOrderID** | **string** | The ID of the move this new shipment is for. | 
**RequestedPickupDate** | Pointer to **NullableString** | The customer&#39;s preferred pickup date. Other dates, such as required delivery date and (outside MilMove) the pack date, are derived from this date.  | [optional] 
**PrimeEstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of this shipment, determined by the movers during the pre-move survey. This value **can only be updated once.** If there was an issue with estimating the weight and a mistake was made, the Prime contractor will need to contact the TOO to change it.  | [optional] 
**CustomerRemarks** | Pointer to **NullableString** | The customer can use the customer remarks field to inform the services counselor and the movers about any special circumstances for this shipment. Typical examples:   * bulky or fragile items,   * weapons,   * access info for their address.  Customer enters this information during onboarding. Optional field.  | [optional] 
**Agents** | Pointer to [**[]MTOAgentV2V2**](MTOAgentV2V2.md) | A list of the agents for a shipment. Agents are the people who the Prime contractor recognize as permitted to release (in the case of pickup) or receive (on delivery) a shipment.  | [optional] 
**MtoServiceItems** | Pointer to [**[]MTOServiceItemV2V2**](MTOServiceItemV2V2.md) | A list of service items connected to this shipment. | [optional] 
**PickupAddress** | Pointer to [**AddressV2V2**](AddressV2.md) | The address where the movers should pick up this shipment. | [optional] 
**DestinationAddress** | Pointer to [**AddressV2V2**](AddressV2.md) | Where the movers should deliver this shipment. | [optional] 
**ShipmentType** | [**MTOShipmentTypeV2V2**](MTOShipmentTypeV2.md) |  | 
**Diversion** | Pointer to **bool** | This value indicates whether or not this shipment is part of a diversion. If yes, the shipment can be either the starting or ending segment of the diversion. When this boolean is true, you must link it to a parent shipment with the divertedFromShipmentId parameter.  | [optional] 
**DivertedFromShipmentId** | Pointer to **string** | The ID of the shipment this is a diversion from. Aka the \&quot;Parent\&quot; shipment. The diversion boolean must be true if this parameter is supplied in the request. If provided, and if the diverted from shipment is also a diversion, the previous should must then also have a parent ID.  | [optional] 
**PointOfContact** | Pointer to **string** | Email or ID of the person who will be contacted in the event of questions or concerns about this update. May be the person performing the update, or someone else working with the Prime contractor.  | [optional] 
**CounselorRemarks** | Pointer to **NullableString** |  | [optional] 
**PpmShipment** | Pointer to [**CreatePPMShipmentV2V2**](CreatePPMShipmentV2.md) |  | [optional] 

## Methods

### NewCreateMTOShipmentV2

`func NewCreateMTOShipmentV2(moveTaskOrderID string, shipmentType MTOShipmentTypeV2V2, ) *CreateMTOShipmentV2`

NewCreateMTOShipmentV2 instantiates a new CreateMTOShipmentV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMTOShipmentV2WithDefaults

`func NewCreateMTOShipmentV2WithDefaults() *CreateMTOShipmentV2`

NewCreateMTOShipmentV2WithDefaults instantiates a new CreateMTOShipmentV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoveTaskOrderID

`func (o *CreateMTOShipmentV2) GetMoveTaskOrderID() string`

GetMoveTaskOrderID returns the MoveTaskOrderID field if non-nil, zero value otherwise.

### GetMoveTaskOrderIDOk

`func (o *CreateMTOShipmentV2) GetMoveTaskOrderIDOk() (*string, bool)`

GetMoveTaskOrderIDOk returns a tuple with the MoveTaskOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTaskOrderID

`func (o *CreateMTOShipmentV2) SetMoveTaskOrderID(v string)`

SetMoveTaskOrderID sets MoveTaskOrderID field to given value.


### GetRequestedPickupDate

`func (o *CreateMTOShipmentV2) GetRequestedPickupDate() string`

GetRequestedPickupDate returns the RequestedPickupDate field if non-nil, zero value otherwise.

### GetRequestedPickupDateOk

`func (o *CreateMTOShipmentV2) GetRequestedPickupDateOk() (*string, bool)`

GetRequestedPickupDateOk returns a tuple with the RequestedPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedPickupDate

`func (o *CreateMTOShipmentV2) SetRequestedPickupDate(v string)`

SetRequestedPickupDate sets RequestedPickupDate field to given value.

### HasRequestedPickupDate

`func (o *CreateMTOShipmentV2) HasRequestedPickupDate() bool`

HasRequestedPickupDate returns a boolean if a field has been set.

### SetRequestedPickupDateNil

`func (o *CreateMTOShipmentV2) SetRequestedPickupDateNil(b bool)`

 SetRequestedPickupDateNil sets the value for RequestedPickupDate to be an explicit nil

### UnsetRequestedPickupDate
`func (o *CreateMTOShipmentV2) UnsetRequestedPickupDate()`

UnsetRequestedPickupDate ensures that no value is present for RequestedPickupDate, not even an explicit nil
### GetPrimeEstimatedWeight

`func (o *CreateMTOShipmentV2) GetPrimeEstimatedWeight() int32`

GetPrimeEstimatedWeight returns the PrimeEstimatedWeight field if non-nil, zero value otherwise.

### GetPrimeEstimatedWeightOk

`func (o *CreateMTOShipmentV2) GetPrimeEstimatedWeightOk() (*int32, bool)`

GetPrimeEstimatedWeightOk returns a tuple with the PrimeEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeEstimatedWeight

`func (o *CreateMTOShipmentV2) SetPrimeEstimatedWeight(v int32)`

SetPrimeEstimatedWeight sets PrimeEstimatedWeight field to given value.

### HasPrimeEstimatedWeight

`func (o *CreateMTOShipmentV2) HasPrimeEstimatedWeight() bool`

HasPrimeEstimatedWeight returns a boolean if a field has been set.

### SetPrimeEstimatedWeightNil

`func (o *CreateMTOShipmentV2) SetPrimeEstimatedWeightNil(b bool)`

 SetPrimeEstimatedWeightNil sets the value for PrimeEstimatedWeight to be an explicit nil

### UnsetPrimeEstimatedWeight
`func (o *CreateMTOShipmentV2) UnsetPrimeEstimatedWeight()`

UnsetPrimeEstimatedWeight ensures that no value is present for PrimeEstimatedWeight, not even an explicit nil
### GetCustomerRemarks

`func (o *CreateMTOShipmentV2) GetCustomerRemarks() string`

GetCustomerRemarks returns the CustomerRemarks field if non-nil, zero value otherwise.

### GetCustomerRemarksOk

`func (o *CreateMTOShipmentV2) GetCustomerRemarksOk() (*string, bool)`

GetCustomerRemarksOk returns a tuple with the CustomerRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRemarks

`func (o *CreateMTOShipmentV2) SetCustomerRemarks(v string)`

SetCustomerRemarks sets CustomerRemarks field to given value.

### HasCustomerRemarks

`func (o *CreateMTOShipmentV2) HasCustomerRemarks() bool`

HasCustomerRemarks returns a boolean if a field has been set.

### SetCustomerRemarksNil

`func (o *CreateMTOShipmentV2) SetCustomerRemarksNil(b bool)`

 SetCustomerRemarksNil sets the value for CustomerRemarks to be an explicit nil

### UnsetCustomerRemarks
`func (o *CreateMTOShipmentV2) UnsetCustomerRemarks()`

UnsetCustomerRemarks ensures that no value is present for CustomerRemarks, not even an explicit nil
### GetAgents

`func (o *CreateMTOShipmentV2) GetAgents() []MTOAgentV2V2`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *CreateMTOShipmentV2) GetAgentsOk() (*[]MTOAgentV2V2, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *CreateMTOShipmentV2) SetAgents(v []MTOAgentV2V2)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *CreateMTOShipmentV2) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetMtoServiceItems

`func (o *CreateMTOShipmentV2) GetMtoServiceItems() []MTOServiceItemV2V2`

GetMtoServiceItems returns the MtoServiceItems field if non-nil, zero value otherwise.

### GetMtoServiceItemsOk

`func (o *CreateMTOShipmentV2) GetMtoServiceItemsOk() (*[]MTOServiceItemV2V2, bool)`

GetMtoServiceItemsOk returns a tuple with the MtoServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItems

`func (o *CreateMTOShipmentV2) SetMtoServiceItems(v []MTOServiceItemV2V2)`

SetMtoServiceItems sets MtoServiceItems field to given value.

### HasMtoServiceItems

`func (o *CreateMTOShipmentV2) HasMtoServiceItems() bool`

HasMtoServiceItems returns a boolean if a field has been set.

### GetPickupAddress

`func (o *CreateMTOShipmentV2) GetPickupAddress() AddressV2V2`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *CreateMTOShipmentV2) GetPickupAddressOk() (*AddressV2V2, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *CreateMTOShipmentV2) SetPickupAddress(v AddressV2V2)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *CreateMTOShipmentV2) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *CreateMTOShipmentV2) GetDestinationAddress() AddressV2V2`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *CreateMTOShipmentV2) GetDestinationAddressOk() (*AddressV2V2, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *CreateMTOShipmentV2) SetDestinationAddress(v AddressV2V2)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *CreateMTOShipmentV2) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetShipmentType

`func (o *CreateMTOShipmentV2) GetShipmentType() MTOShipmentTypeV2V2`

GetShipmentType returns the ShipmentType field if non-nil, zero value otherwise.

### GetShipmentTypeOk

`func (o *CreateMTOShipmentV2) GetShipmentTypeOk() (*MTOShipmentTypeV2V2, bool)`

GetShipmentTypeOk returns a tuple with the ShipmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentType

`func (o *CreateMTOShipmentV2) SetShipmentType(v MTOShipmentTypeV2V2)`

SetShipmentType sets ShipmentType field to given value.


### GetDiversion

`func (o *CreateMTOShipmentV2) GetDiversion() bool`

GetDiversion returns the Diversion field if non-nil, zero value otherwise.

### GetDiversionOk

`func (o *CreateMTOShipmentV2) GetDiversionOk() (*bool, bool)`

GetDiversionOk returns a tuple with the Diversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiversion

`func (o *CreateMTOShipmentV2) SetDiversion(v bool)`

SetDiversion sets Diversion field to given value.

### HasDiversion

`func (o *CreateMTOShipmentV2) HasDiversion() bool`

HasDiversion returns a boolean if a field has been set.

### GetDivertedFromShipmentId

`func (o *CreateMTOShipmentV2) GetDivertedFromShipmentId() string`

GetDivertedFromShipmentId returns the DivertedFromShipmentId field if non-nil, zero value otherwise.

### GetDivertedFromShipmentIdOk

`func (o *CreateMTOShipmentV2) GetDivertedFromShipmentIdOk() (*string, bool)`

GetDivertedFromShipmentIdOk returns a tuple with the DivertedFromShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivertedFromShipmentId

`func (o *CreateMTOShipmentV2) SetDivertedFromShipmentId(v string)`

SetDivertedFromShipmentId sets DivertedFromShipmentId field to given value.

### HasDivertedFromShipmentId

`func (o *CreateMTOShipmentV2) HasDivertedFromShipmentId() bool`

HasDivertedFromShipmentId returns a boolean if a field has been set.

### GetPointOfContact

`func (o *CreateMTOShipmentV2) GetPointOfContact() string`

GetPointOfContact returns the PointOfContact field if non-nil, zero value otherwise.

### GetPointOfContactOk

`func (o *CreateMTOShipmentV2) GetPointOfContactOk() (*string, bool)`

GetPointOfContactOk returns a tuple with the PointOfContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointOfContact

`func (o *CreateMTOShipmentV2) SetPointOfContact(v string)`

SetPointOfContact sets PointOfContact field to given value.

### HasPointOfContact

`func (o *CreateMTOShipmentV2) HasPointOfContact() bool`

HasPointOfContact returns a boolean if a field has been set.

### GetCounselorRemarks

`func (o *CreateMTOShipmentV2) GetCounselorRemarks() string`

GetCounselorRemarks returns the CounselorRemarks field if non-nil, zero value otherwise.

### GetCounselorRemarksOk

`func (o *CreateMTOShipmentV2) GetCounselorRemarksOk() (*string, bool)`

GetCounselorRemarksOk returns a tuple with the CounselorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounselorRemarks

`func (o *CreateMTOShipmentV2) SetCounselorRemarks(v string)`

SetCounselorRemarks sets CounselorRemarks field to given value.

### HasCounselorRemarks

`func (o *CreateMTOShipmentV2) HasCounselorRemarks() bool`

HasCounselorRemarks returns a boolean if a field has been set.

### SetCounselorRemarksNil

`func (o *CreateMTOShipmentV2) SetCounselorRemarksNil(b bool)`

 SetCounselorRemarksNil sets the value for CounselorRemarks to be an explicit nil

### UnsetCounselorRemarks
`func (o *CreateMTOShipmentV2) UnsetCounselorRemarks()`

UnsetCounselorRemarks ensures that no value is present for CounselorRemarks, not even an explicit nil
### GetPpmShipment

`func (o *CreateMTOShipmentV2) GetPpmShipment() CreatePPMShipmentV2V2`

GetPpmShipment returns the PpmShipment field if non-nil, zero value otherwise.

### GetPpmShipmentOk

`func (o *CreateMTOShipmentV2) GetPpmShipmentOk() (*CreatePPMShipmentV2V2, bool)`

GetPpmShipmentOk returns a tuple with the PpmShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpmShipment

`func (o *CreateMTOShipmentV2) SetPpmShipment(v CreatePPMShipmentV2V2)`

SetPpmShipment sets PpmShipment field to given value.

### HasPpmShipment

`func (o *CreateMTOShipmentV2) HasPpmShipment() bool`

HasPpmShipment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


