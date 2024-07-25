# CreateMTOShipmentV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveTaskOrderID** | **string** | The ID of the move this new shipment is for. | 
**RequestedPickupDate** | Pointer to **NullableString** | The customer&#39;s preferred pickup date. Other dates, such as required delivery date and (outside MilMove) the pack date, are derived from this date.  | [optional] 
**PrimeEstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of this shipment, determined by the movers during the pre-move survey. This value **can only be updated once.** If there was an issue with estimating the weight and a mistake was made, the Prime contractor will need to contact the TOO to change it.  | [optional] 
**CustomerRemarks** | Pointer to **NullableString** | The customer can use the customer remarks field to inform the services counselor and the movers about any special circumstances for this shipment. Typical examples:   * bulky or fragile items,   * weapons,   * access info for their address.  Customer enters this information during onboarding. Optional field.  | [optional] 
**Agents** | Pointer to [**[]MTOAgentV3V3**](MTOAgentV3V3.md) | A list of the agents for a shipment. Agents are the people who the Prime contractor recognize as permitted to release (in the case of pickup) or receive (on delivery) a shipment.  | [optional] 
**MtoServiceItems** | Pointer to [**[]MTOServiceItemV3V3**](MTOServiceItemV3V3.md) | A list of service items connected to this shipment. | [optional] 
**PickupAddress** | Pointer to [**AddressV3V3**](AddressV3.md) | The address where the movers should pick up this shipment. | [optional] 
**DestinationAddress** | Pointer to [**AddressV3V3**](AddressV3.md) | Where the movers should deliver this shipment. | [optional] 
**ShipmentType** | [**MTOShipmentTypeV3V3**](MTOShipmentTypeV3.md) |  | 
**Diversion** | Pointer to **bool** | This value indicates whether or not this shipment is part of a diversion. If yes, the shipment can be either the starting or ending segment of the diversion. When this boolean is true, you must link it to a parent shipment with the divertedFromShipmentId parameter.  | [optional] 
**DivertedFromShipmentId** | Pointer to **string** | The ID of the shipment this is a diversion from. Aka the \&quot;Parent\&quot; shipment. The diversion boolean must be true if this parameter is supplied in the request. If provided, and if the diverted from shipment is also a diversion, the previous should must then also have a parent ID.  | [optional] 
**PointOfContact** | Pointer to **string** | Email or ID of the person who will be contacted in the event of questions or concerns about this update. May be the person performing the update, or someone else working with the Prime contractor.  | [optional] 
**CounselorRemarks** | Pointer to **NullableString** |  | [optional] 
**PpmShipment** | Pointer to [**CreatePPMShipmentV3V3**](CreatePPMShipmentV3.md) |  | [optional] 

## Methods

### NewCreateMTOShipmentV3

`func NewCreateMTOShipmentV3(moveTaskOrderID string, shipmentType MTOShipmentTypeV3V3, ) *CreateMTOShipmentV3`

NewCreateMTOShipmentV3 instantiates a new CreateMTOShipmentV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMTOShipmentV3WithDefaults

`func NewCreateMTOShipmentV3WithDefaults() *CreateMTOShipmentV3`

NewCreateMTOShipmentV3WithDefaults instantiates a new CreateMTOShipmentV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoveTaskOrderID

`func (o *CreateMTOShipmentV3) GetMoveTaskOrderID() string`

GetMoveTaskOrderID returns the MoveTaskOrderID field if non-nil, zero value otherwise.

### GetMoveTaskOrderIDOk

`func (o *CreateMTOShipmentV3) GetMoveTaskOrderIDOk() (*string, bool)`

GetMoveTaskOrderIDOk returns a tuple with the MoveTaskOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTaskOrderID

`func (o *CreateMTOShipmentV3) SetMoveTaskOrderID(v string)`

SetMoveTaskOrderID sets MoveTaskOrderID field to given value.


### GetRequestedPickupDate

`func (o *CreateMTOShipmentV3) GetRequestedPickupDate() string`

GetRequestedPickupDate returns the RequestedPickupDate field if non-nil, zero value otherwise.

### GetRequestedPickupDateOk

`func (o *CreateMTOShipmentV3) GetRequestedPickupDateOk() (*string, bool)`

GetRequestedPickupDateOk returns a tuple with the RequestedPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedPickupDate

`func (o *CreateMTOShipmentV3) SetRequestedPickupDate(v string)`

SetRequestedPickupDate sets RequestedPickupDate field to given value.

### HasRequestedPickupDate

`func (o *CreateMTOShipmentV3) HasRequestedPickupDate() bool`

HasRequestedPickupDate returns a boolean if a field has been set.

### SetRequestedPickupDateNil

`func (o *CreateMTOShipmentV3) SetRequestedPickupDateNil(b bool)`

 SetRequestedPickupDateNil sets the value for RequestedPickupDate to be an explicit nil

### UnsetRequestedPickupDate
`func (o *CreateMTOShipmentV3) UnsetRequestedPickupDate()`

UnsetRequestedPickupDate ensures that no value is present for RequestedPickupDate, not even an explicit nil
### GetPrimeEstimatedWeight

`func (o *CreateMTOShipmentV3) GetPrimeEstimatedWeight() int32`

GetPrimeEstimatedWeight returns the PrimeEstimatedWeight field if non-nil, zero value otherwise.

### GetPrimeEstimatedWeightOk

`func (o *CreateMTOShipmentV3) GetPrimeEstimatedWeightOk() (*int32, bool)`

GetPrimeEstimatedWeightOk returns a tuple with the PrimeEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeEstimatedWeight

`func (o *CreateMTOShipmentV3) SetPrimeEstimatedWeight(v int32)`

SetPrimeEstimatedWeight sets PrimeEstimatedWeight field to given value.

### HasPrimeEstimatedWeight

`func (o *CreateMTOShipmentV3) HasPrimeEstimatedWeight() bool`

HasPrimeEstimatedWeight returns a boolean if a field has been set.

### SetPrimeEstimatedWeightNil

`func (o *CreateMTOShipmentV3) SetPrimeEstimatedWeightNil(b bool)`

 SetPrimeEstimatedWeightNil sets the value for PrimeEstimatedWeight to be an explicit nil

### UnsetPrimeEstimatedWeight
`func (o *CreateMTOShipmentV3) UnsetPrimeEstimatedWeight()`

UnsetPrimeEstimatedWeight ensures that no value is present for PrimeEstimatedWeight, not even an explicit nil
### GetCustomerRemarks

`func (o *CreateMTOShipmentV3) GetCustomerRemarks() string`

GetCustomerRemarks returns the CustomerRemarks field if non-nil, zero value otherwise.

### GetCustomerRemarksOk

`func (o *CreateMTOShipmentV3) GetCustomerRemarksOk() (*string, bool)`

GetCustomerRemarksOk returns a tuple with the CustomerRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRemarks

`func (o *CreateMTOShipmentV3) SetCustomerRemarks(v string)`

SetCustomerRemarks sets CustomerRemarks field to given value.

### HasCustomerRemarks

`func (o *CreateMTOShipmentV3) HasCustomerRemarks() bool`

HasCustomerRemarks returns a boolean if a field has been set.

### SetCustomerRemarksNil

`func (o *CreateMTOShipmentV3) SetCustomerRemarksNil(b bool)`

 SetCustomerRemarksNil sets the value for CustomerRemarks to be an explicit nil

### UnsetCustomerRemarks
`func (o *CreateMTOShipmentV3) UnsetCustomerRemarks()`

UnsetCustomerRemarks ensures that no value is present for CustomerRemarks, not even an explicit nil
### GetAgents

`func (o *CreateMTOShipmentV3) GetAgents() []MTOAgentV3V3`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *CreateMTOShipmentV3) GetAgentsOk() (*[]MTOAgentV3V3, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *CreateMTOShipmentV3) SetAgents(v []MTOAgentV3V3)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *CreateMTOShipmentV3) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetMtoServiceItems

`func (o *CreateMTOShipmentV3) GetMtoServiceItems() []MTOServiceItemV3V3`

GetMtoServiceItems returns the MtoServiceItems field if non-nil, zero value otherwise.

### GetMtoServiceItemsOk

`func (o *CreateMTOShipmentV3) GetMtoServiceItemsOk() (*[]MTOServiceItemV3V3, bool)`

GetMtoServiceItemsOk returns a tuple with the MtoServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItems

`func (o *CreateMTOShipmentV3) SetMtoServiceItems(v []MTOServiceItemV3V3)`

SetMtoServiceItems sets MtoServiceItems field to given value.

### HasMtoServiceItems

`func (o *CreateMTOShipmentV3) HasMtoServiceItems() bool`

HasMtoServiceItems returns a boolean if a field has been set.

### GetPickupAddress

`func (o *CreateMTOShipmentV3) GetPickupAddress() AddressV3V3`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *CreateMTOShipmentV3) GetPickupAddressOk() (*AddressV3V3, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *CreateMTOShipmentV3) SetPickupAddress(v AddressV3V3)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *CreateMTOShipmentV3) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *CreateMTOShipmentV3) GetDestinationAddress() AddressV3V3`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *CreateMTOShipmentV3) GetDestinationAddressOk() (*AddressV3V3, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *CreateMTOShipmentV3) SetDestinationAddress(v AddressV3V3)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *CreateMTOShipmentV3) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetShipmentType

`func (o *CreateMTOShipmentV3) GetShipmentType() MTOShipmentTypeV3V3`

GetShipmentType returns the ShipmentType field if non-nil, zero value otherwise.

### GetShipmentTypeOk

`func (o *CreateMTOShipmentV3) GetShipmentTypeOk() (*MTOShipmentTypeV3V3, bool)`

GetShipmentTypeOk returns a tuple with the ShipmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentType

`func (o *CreateMTOShipmentV3) SetShipmentType(v MTOShipmentTypeV3V3)`

SetShipmentType sets ShipmentType field to given value.


### GetDiversion

`func (o *CreateMTOShipmentV3) GetDiversion() bool`

GetDiversion returns the Diversion field if non-nil, zero value otherwise.

### GetDiversionOk

`func (o *CreateMTOShipmentV3) GetDiversionOk() (*bool, bool)`

GetDiversionOk returns a tuple with the Diversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiversion

`func (o *CreateMTOShipmentV3) SetDiversion(v bool)`

SetDiversion sets Diversion field to given value.

### HasDiversion

`func (o *CreateMTOShipmentV3) HasDiversion() bool`

HasDiversion returns a boolean if a field has been set.

### GetDivertedFromShipmentId

`func (o *CreateMTOShipmentV3) GetDivertedFromShipmentId() string`

GetDivertedFromShipmentId returns the DivertedFromShipmentId field if non-nil, zero value otherwise.

### GetDivertedFromShipmentIdOk

`func (o *CreateMTOShipmentV3) GetDivertedFromShipmentIdOk() (*string, bool)`

GetDivertedFromShipmentIdOk returns a tuple with the DivertedFromShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDivertedFromShipmentId

`func (o *CreateMTOShipmentV3) SetDivertedFromShipmentId(v string)`

SetDivertedFromShipmentId sets DivertedFromShipmentId field to given value.

### HasDivertedFromShipmentId

`func (o *CreateMTOShipmentV3) HasDivertedFromShipmentId() bool`

HasDivertedFromShipmentId returns a boolean if a field has been set.

### GetPointOfContact

`func (o *CreateMTOShipmentV3) GetPointOfContact() string`

GetPointOfContact returns the PointOfContact field if non-nil, zero value otherwise.

### GetPointOfContactOk

`func (o *CreateMTOShipmentV3) GetPointOfContactOk() (*string, bool)`

GetPointOfContactOk returns a tuple with the PointOfContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointOfContact

`func (o *CreateMTOShipmentV3) SetPointOfContact(v string)`

SetPointOfContact sets PointOfContact field to given value.

### HasPointOfContact

`func (o *CreateMTOShipmentV3) HasPointOfContact() bool`

HasPointOfContact returns a boolean if a field has been set.

### GetCounselorRemarks

`func (o *CreateMTOShipmentV3) GetCounselorRemarks() string`

GetCounselorRemarks returns the CounselorRemarks field if non-nil, zero value otherwise.

### GetCounselorRemarksOk

`func (o *CreateMTOShipmentV3) GetCounselorRemarksOk() (*string, bool)`

GetCounselorRemarksOk returns a tuple with the CounselorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounselorRemarks

`func (o *CreateMTOShipmentV3) SetCounselorRemarks(v string)`

SetCounselorRemarks sets CounselorRemarks field to given value.

### HasCounselorRemarks

`func (o *CreateMTOShipmentV3) HasCounselorRemarks() bool`

HasCounselorRemarks returns a boolean if a field has been set.

### SetCounselorRemarksNil

`func (o *CreateMTOShipmentV3) SetCounselorRemarksNil(b bool)`

 SetCounselorRemarksNil sets the value for CounselorRemarks to be an explicit nil

### UnsetCounselorRemarks
`func (o *CreateMTOShipmentV3) UnsetCounselorRemarks()`

UnsetCounselorRemarks ensures that no value is present for CounselorRemarks, not even an explicit nil
### GetPpmShipment

`func (o *CreateMTOShipmentV3) GetPpmShipment() CreatePPMShipmentV3V3`

GetPpmShipment returns the PpmShipment field if non-nil, zero value otherwise.

### GetPpmShipmentOk

`func (o *CreateMTOShipmentV3) GetPpmShipmentOk() (*CreatePPMShipmentV3V3, bool)`

GetPpmShipmentOk returns a tuple with the PpmShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpmShipment

`func (o *CreateMTOShipmentV3) SetPpmShipment(v CreatePPMShipmentV3V3)`

SetPpmShipment sets PpmShipment field to given value.

### HasPpmShipment

`func (o *CreateMTOShipmentV3) HasPpmShipment() bool`

HasPpmShipment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


