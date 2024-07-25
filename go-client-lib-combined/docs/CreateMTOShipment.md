# CreateMTOShipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveTaskOrderID** | **string** | The ID of the move this new shipment is for. | 
**RequestedPickupDate** | Pointer to **NullableString** | The customer&#39;s preferred pickup date. Other dates, such as required delivery date and (outside MilMove) the pack date, are derived from this date.  | [optional] 
**PrimeEstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of this shipment, determined by the movers during the pre-move survey. This value **can only be updated once.** If there was an issue with estimating the weight and a mistake was made, the Prime contractor will need to contact the TOO to change it.  | [optional] 
**CustomerRemarks** | Pointer to **NullableString** | The customer can use the customer remarks field to inform the services counselor and the movers about any special circumstances for this shipment. Typical examples:   * bulky or fragile items,   * weapons,   * access info for their address.  Customer enters this information during onboarding. Optional field.  | [optional] 
**Agents** | Pointer to [**[]MTOAgent**](MTOAgent.md) | A list of the agents for a shipment. Agents are the people who the Prime contractor recognize as permitted to release (in the case of pickup) or receive (on delivery) a shipment.  | [optional] 
**MtoServiceItems** | Pointer to [**[]MTOServiceItem**](MTOServiceItem.md) | A list of service items connected to this shipment. | [optional] 
**PickupAddress** | Pointer to [**Address**](Address.md) | The address where the movers should pick up this shipment. | [optional] 
**DestinationAddress** | Pointer to [**Address**](Address.md) | Where the movers should deliver this shipment. | [optional] 
**ShipmentType** | [**MTOShipmentType**](MTOShipmentType.md) |  | 
**Diversion** | Pointer to **bool** | This value indicates whether or not this shipment is part of a diversion. If yes, the shipment can be either the starting or ending segment of the diversion.  | [optional] 
**PointOfContact** | Pointer to **string** | Email or ID of the person who will be contacted in the event of questions or concerns about this update. May be the person performing the update, or someone else working with the Prime contractor.  | [optional] 
**CounselorRemarks** | Pointer to **NullableString** |  | [optional] 
**PpmShipment** | Pointer to [**CreatePPMShipment**](CreatePPMShipment.md) |  | [optional] 

## Methods

### NewCreateMTOShipment

`func NewCreateMTOShipment(moveTaskOrderID string, shipmentType MTOShipmentType, ) *CreateMTOShipment`

NewCreateMTOShipment instantiates a new CreateMTOShipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateMTOShipmentWithDefaults

`func NewCreateMTOShipmentWithDefaults() *CreateMTOShipment`

NewCreateMTOShipmentWithDefaults instantiates a new CreateMTOShipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoveTaskOrderID

`func (o *CreateMTOShipment) GetMoveTaskOrderID() string`

GetMoveTaskOrderID returns the MoveTaskOrderID field if non-nil, zero value otherwise.

### GetMoveTaskOrderIDOk

`func (o *CreateMTOShipment) GetMoveTaskOrderIDOk() (*string, bool)`

GetMoveTaskOrderIDOk returns a tuple with the MoveTaskOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTaskOrderID

`func (o *CreateMTOShipment) SetMoveTaskOrderID(v string)`

SetMoveTaskOrderID sets MoveTaskOrderID field to given value.


### GetRequestedPickupDate

`func (o *CreateMTOShipment) GetRequestedPickupDate() string`

GetRequestedPickupDate returns the RequestedPickupDate field if non-nil, zero value otherwise.

### GetRequestedPickupDateOk

`func (o *CreateMTOShipment) GetRequestedPickupDateOk() (*string, bool)`

GetRequestedPickupDateOk returns a tuple with the RequestedPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedPickupDate

`func (o *CreateMTOShipment) SetRequestedPickupDate(v string)`

SetRequestedPickupDate sets RequestedPickupDate field to given value.

### HasRequestedPickupDate

`func (o *CreateMTOShipment) HasRequestedPickupDate() bool`

HasRequestedPickupDate returns a boolean if a field has been set.

### SetRequestedPickupDateNil

`func (o *CreateMTOShipment) SetRequestedPickupDateNil(b bool)`

 SetRequestedPickupDateNil sets the value for RequestedPickupDate to be an explicit nil

### UnsetRequestedPickupDate
`func (o *CreateMTOShipment) UnsetRequestedPickupDate()`

UnsetRequestedPickupDate ensures that no value is present for RequestedPickupDate, not even an explicit nil
### GetPrimeEstimatedWeight

`func (o *CreateMTOShipment) GetPrimeEstimatedWeight() int32`

GetPrimeEstimatedWeight returns the PrimeEstimatedWeight field if non-nil, zero value otherwise.

### GetPrimeEstimatedWeightOk

`func (o *CreateMTOShipment) GetPrimeEstimatedWeightOk() (*int32, bool)`

GetPrimeEstimatedWeightOk returns a tuple with the PrimeEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeEstimatedWeight

`func (o *CreateMTOShipment) SetPrimeEstimatedWeight(v int32)`

SetPrimeEstimatedWeight sets PrimeEstimatedWeight field to given value.

### HasPrimeEstimatedWeight

`func (o *CreateMTOShipment) HasPrimeEstimatedWeight() bool`

HasPrimeEstimatedWeight returns a boolean if a field has been set.

### SetPrimeEstimatedWeightNil

`func (o *CreateMTOShipment) SetPrimeEstimatedWeightNil(b bool)`

 SetPrimeEstimatedWeightNil sets the value for PrimeEstimatedWeight to be an explicit nil

### UnsetPrimeEstimatedWeight
`func (o *CreateMTOShipment) UnsetPrimeEstimatedWeight()`

UnsetPrimeEstimatedWeight ensures that no value is present for PrimeEstimatedWeight, not even an explicit nil
### GetCustomerRemarks

`func (o *CreateMTOShipment) GetCustomerRemarks() string`

GetCustomerRemarks returns the CustomerRemarks field if non-nil, zero value otherwise.

### GetCustomerRemarksOk

`func (o *CreateMTOShipment) GetCustomerRemarksOk() (*string, bool)`

GetCustomerRemarksOk returns a tuple with the CustomerRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRemarks

`func (o *CreateMTOShipment) SetCustomerRemarks(v string)`

SetCustomerRemarks sets CustomerRemarks field to given value.

### HasCustomerRemarks

`func (o *CreateMTOShipment) HasCustomerRemarks() bool`

HasCustomerRemarks returns a boolean if a field has been set.

### SetCustomerRemarksNil

`func (o *CreateMTOShipment) SetCustomerRemarksNil(b bool)`

 SetCustomerRemarksNil sets the value for CustomerRemarks to be an explicit nil

### UnsetCustomerRemarks
`func (o *CreateMTOShipment) UnsetCustomerRemarks()`

UnsetCustomerRemarks ensures that no value is present for CustomerRemarks, not even an explicit nil
### GetAgents

`func (o *CreateMTOShipment) GetAgents() []MTOAgent`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *CreateMTOShipment) GetAgentsOk() (*[]MTOAgent, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *CreateMTOShipment) SetAgents(v []MTOAgent)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *CreateMTOShipment) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetMtoServiceItems

`func (o *CreateMTOShipment) GetMtoServiceItems() []MTOServiceItem`

GetMtoServiceItems returns the MtoServiceItems field if non-nil, zero value otherwise.

### GetMtoServiceItemsOk

`func (o *CreateMTOShipment) GetMtoServiceItemsOk() (*[]MTOServiceItem, bool)`

GetMtoServiceItemsOk returns a tuple with the MtoServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItems

`func (o *CreateMTOShipment) SetMtoServiceItems(v []MTOServiceItem)`

SetMtoServiceItems sets MtoServiceItems field to given value.

### HasMtoServiceItems

`func (o *CreateMTOShipment) HasMtoServiceItems() bool`

HasMtoServiceItems returns a boolean if a field has been set.

### GetPickupAddress

`func (o *CreateMTOShipment) GetPickupAddress() Address`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *CreateMTOShipment) GetPickupAddressOk() (*Address, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *CreateMTOShipment) SetPickupAddress(v Address)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *CreateMTOShipment) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *CreateMTOShipment) GetDestinationAddress() Address`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *CreateMTOShipment) GetDestinationAddressOk() (*Address, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *CreateMTOShipment) SetDestinationAddress(v Address)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *CreateMTOShipment) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetShipmentType

`func (o *CreateMTOShipment) GetShipmentType() MTOShipmentType`

GetShipmentType returns the ShipmentType field if non-nil, zero value otherwise.

### GetShipmentTypeOk

`func (o *CreateMTOShipment) GetShipmentTypeOk() (*MTOShipmentType, bool)`

GetShipmentTypeOk returns a tuple with the ShipmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentType

`func (o *CreateMTOShipment) SetShipmentType(v MTOShipmentType)`

SetShipmentType sets ShipmentType field to given value.


### GetDiversion

`func (o *CreateMTOShipment) GetDiversion() bool`

GetDiversion returns the Diversion field if non-nil, zero value otherwise.

### GetDiversionOk

`func (o *CreateMTOShipment) GetDiversionOk() (*bool, bool)`

GetDiversionOk returns a tuple with the Diversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiversion

`func (o *CreateMTOShipment) SetDiversion(v bool)`

SetDiversion sets Diversion field to given value.

### HasDiversion

`func (o *CreateMTOShipment) HasDiversion() bool`

HasDiversion returns a boolean if a field has been set.

### GetPointOfContact

`func (o *CreateMTOShipment) GetPointOfContact() string`

GetPointOfContact returns the PointOfContact field if non-nil, zero value otherwise.

### GetPointOfContactOk

`func (o *CreateMTOShipment) GetPointOfContactOk() (*string, bool)`

GetPointOfContactOk returns a tuple with the PointOfContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointOfContact

`func (o *CreateMTOShipment) SetPointOfContact(v string)`

SetPointOfContact sets PointOfContact field to given value.

### HasPointOfContact

`func (o *CreateMTOShipment) HasPointOfContact() bool`

HasPointOfContact returns a boolean if a field has been set.

### GetCounselorRemarks

`func (o *CreateMTOShipment) GetCounselorRemarks() string`

GetCounselorRemarks returns the CounselorRemarks field if non-nil, zero value otherwise.

### GetCounselorRemarksOk

`func (o *CreateMTOShipment) GetCounselorRemarksOk() (*string, bool)`

GetCounselorRemarksOk returns a tuple with the CounselorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounselorRemarks

`func (o *CreateMTOShipment) SetCounselorRemarks(v string)`

SetCounselorRemarks sets CounselorRemarks field to given value.

### HasCounselorRemarks

`func (o *CreateMTOShipment) HasCounselorRemarks() bool`

HasCounselorRemarks returns a boolean if a field has been set.

### SetCounselorRemarksNil

`func (o *CreateMTOShipment) SetCounselorRemarksNil(b bool)`

 SetCounselorRemarksNil sets the value for CounselorRemarks to be an explicit nil

### UnsetCounselorRemarks
`func (o *CreateMTOShipment) UnsetCounselorRemarks()`

UnsetCounselorRemarks ensures that no value is present for CounselorRemarks, not even an explicit nil
### GetPpmShipment

`func (o *CreateMTOShipment) GetPpmShipment() CreatePPMShipment`

GetPpmShipment returns the PpmShipment field if non-nil, zero value otherwise.

### GetPpmShipmentOk

`func (o *CreateMTOShipment) GetPpmShipmentOk() (*CreatePPMShipment, bool)`

GetPpmShipmentOk returns a tuple with the PpmShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpmShipment

`func (o *CreateMTOShipment) SetPpmShipment(v CreatePPMShipment)`

SetPpmShipment sets PpmShipment field to given value.

### HasPpmShipment

`func (o *CreateMTOShipment) HasPpmShipment() bool`

HasPpmShipment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


