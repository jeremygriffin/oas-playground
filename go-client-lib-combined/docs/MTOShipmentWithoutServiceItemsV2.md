# MTOShipmentWithoutServiceItemsV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the shipment. | [optional] [readonly] 
**MoveTaskOrderID** | Pointer to **string** | The ID of the move for this shipment. | [optional] [readonly] 
**ApprovedDate** | Pointer to **NullableString** | The date when the Task Ordering Officer first approved this shipment for the move. | [optional] [readonly] 
**RequestedPickupDate** | Pointer to **NullableString** | The date the customer selects during onboarding as their preferred pickup date. Other dates, such as required delivery date and (outside MilMove) the pack date, are derived from this date.  | [optional] [readonly] 
**RequestedDeliveryDate** | Pointer to **NullableString** | The customer&#39;s preferred delivery date. | [optional] [readonly] 
**ScheduledPickupDate** | Pointer to **NullableString** | The date the Prime contractor scheduled to pick up this shipment after consultation with the customer. | [optional] 
**ActualPickupDate** | Pointer to **NullableString** | The date when the Prime contractor actually picked up the shipment. Updated after-the-fact. | [optional] 
**FirstAvailableDeliveryDate** | Pointer to **NullableString** | The date the Prime provides to the customer as the first possible delivery date so that they can plan their travel accordingly.  | [optional] 
**RequiredDeliveryDate** | Pointer to **NullableString** | The latest date by which the Prime can deliver a customer&#39;s shipment without violating the contract. This is calculated based on weight, distance, and the scheduled pickup date. It cannot be modified.  | [optional] [readonly] 
**ScheduledDeliveryDate** | Pointer to **NullableString** | The date the Prime contractor scheduled to deliver this shipment after consultation with the customer. | [optional] 
**ActualDeliveryDate** | Pointer to **NullableString** | The date when the Prime contractor actually delivered the shipment. Updated after-the-fact. | [optional] 
**PrimeEstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of this shipment, determined by the movers during the pre-move survey. This value **can only be updated once.** If there was an issue with estimating the weight and a mistake was made, the Prime contracter will need to contact the TOO to change it.  | [optional] 
**PrimeEstimatedWeightRecordedDate** | Pointer to **NullableString** | The date when the Prime contractor recorded the shipment&#39;s estimated weight. | [optional] [readonly] 
**PrimeActualWeight** | Pointer to **NullableInt32** | The actual weight of the shipment, provided after the Prime packs, picks up, and weighs a customer&#39;s shipment. | [optional] 
**NtsRecordedWeight** | Pointer to **NullableInt32** | The previously recorded weight for the NTS Shipment. Used for NTS Release to know what the previous primeActualWeight or billable weight was. | [optional] 
**CustomerRemarks** | Pointer to **NullableString** | The customer can use the customer remarks field to inform the services counselor and the movers about any special circumstances for this shipment. Typical examples:   * bulky or fragile items,   * weapons,   * access info for their address.  Customer enters this information during onboarding. Optional field.  | [optional] [readonly] 
**CounselorRemarks** | Pointer to **NullableString** | The counselor can use the counselor remarks field to inform the movers about any special circumstances for this shipment. Typical examples:   * bulky or fragile items,   * weapons,   * access info for their address.  Counselors enters this information when creating or editing an MTO Shipment. Optional field.  | [optional] [readonly] 
**ActualProGearWeight** | Pointer to **NullableInt32** | The actual weight of any pro gear being shipped.  | [optional] 
**ActualSpouseProGearWeight** | Pointer to **NullableInt32** | The actual weight of any spouse pro gear being shipped.  | [optional] 
**Agents** | Pointer to [**[]MTOAgentV2V2**](MTOAgentV2V2.md) | A list of the agents for a shipment. Agents are the people who the Prime contractor recognize as permitted to release (in the case of pickup) or receive (on delivery) a shipment.  | [optional] 
**SitExtensions** | Pointer to [**[]SITExtensionV2V2**](SITExtensionV2V2.md) |  | [optional] 
**Reweigh** | Pointer to [**ReweighV2V2**](ReweighV2.md) |  | [optional] 
**PickupAddress** | Pointer to [**AddressV2V2**](AddressV2.md) | The address where the movers should pick up this shipment, entered by the customer during onboarding when they enter shipment details.  | [optional] 
**DestinationAddress** | Pointer to [**AddressV2V2**](AddressV2.md) | Where the movers should deliver this shipment. Often provided by the customer when they enter shipment details during onboarding, if they know their new address already.  May be blank when entered by the customer, required when entered by the Prime. May not represent the true final destination due to the shipment being diverted or placed in SIT.  | [optional] 
**DestinationType** | Pointer to [**NullableDestinationTypeV2V2**](DestinationTypeV2.md) |  | [optional] 
**SecondaryPickupAddress** | Pointer to [**AddressV2V2**](AddressV2.md) | A second pickup address for this shipment, if the customer entered one. An optional field. | [optional] 
**SecondaryDeliveryAddress** | Pointer to [**AddressV2V2**](AddressV2.md) | A second delivery address for this shipment, if the customer entered one. An optional field. | [optional] 
**StorageFacility** | Pointer to [**UpdateMTOShipmentStorageFacilityV2V2**](UpdateMTOShipmentStorageFacilityV2.md) |  | [optional] 
**ShipmentType** | Pointer to [**MTOShipmentTypeV2V2**](MTOShipmentTypeV2.md) |  | [optional] 
**Diversion** | Pointer to **bool** | This value indicates whether or not this shipment is part of a diversion. If yes, the shipment can be either the starting or ending segment of the diversion.  | [optional] 
**DiversionReason** | Pointer to **NullableString** | The reason the TOO provided when requesting a diversion for this shipment.  | [optional] [readonly] 
**Status** | Pointer to **string** | The status of a shipment, indicating where it is in the TOO&#39;s approval process. Can only be updated by the contractor in special circumstances.  | [optional] [readonly] 
**PpmShipment** | Pointer to [**NullablePPMShipmentV2V2**](PPMShipmentV2.md) |  | [optional] 
**DeliveryAddressUpdate** | Pointer to [**ShipmentAddressUpdateV2V2**](ShipmentAddressUpdateV2.md) |  | [optional] 
**ETag** | Pointer to **string** | A hash unique to this shipment that should be used as the \&quot;If-Match\&quot; header for any updates. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**PointOfContact** | Pointer to **string** | Email or ID of the person who will be contacted in the event of questions or concerns about this update. May be the person performing the update, or someone else working with the Prime contractor.  | [optional] 
**OriginSitAuthEndDate** | Pointer to **NullableString** | The SIT authorized end date for origin SIT. | [optional] 
**DestinationSitAuthEndDate** | Pointer to **NullableString** | The SIT authorized end date for destination SIT. | [optional] 

## Methods

### NewMTOShipmentWithoutServiceItemsV2

`func NewMTOShipmentWithoutServiceItemsV2() *MTOShipmentWithoutServiceItemsV2`

NewMTOShipmentWithoutServiceItemsV2 instantiates a new MTOShipmentWithoutServiceItemsV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTOShipmentWithoutServiceItemsV2WithDefaults

`func NewMTOShipmentWithoutServiceItemsV2WithDefaults() *MTOShipmentWithoutServiceItemsV2`

NewMTOShipmentWithoutServiceItemsV2WithDefaults instantiates a new MTOShipmentWithoutServiceItemsV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MTOShipmentWithoutServiceItemsV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MTOShipmentWithoutServiceItemsV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MTOShipmentWithoutServiceItemsV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMoveTaskOrderID

`func (o *MTOShipmentWithoutServiceItemsV2) GetMoveTaskOrderID() string`

GetMoveTaskOrderID returns the MoveTaskOrderID field if non-nil, zero value otherwise.

### GetMoveTaskOrderIDOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetMoveTaskOrderIDOk() (*string, bool)`

GetMoveTaskOrderIDOk returns a tuple with the MoveTaskOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTaskOrderID

`func (o *MTOShipmentWithoutServiceItemsV2) SetMoveTaskOrderID(v string)`

SetMoveTaskOrderID sets MoveTaskOrderID field to given value.

### HasMoveTaskOrderID

`func (o *MTOShipmentWithoutServiceItemsV2) HasMoveTaskOrderID() bool`

HasMoveTaskOrderID returns a boolean if a field has been set.

### GetApprovedDate

`func (o *MTOShipmentWithoutServiceItemsV2) GetApprovedDate() string`

GetApprovedDate returns the ApprovedDate field if non-nil, zero value otherwise.

### GetApprovedDateOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetApprovedDateOk() (*string, bool)`

GetApprovedDateOk returns a tuple with the ApprovedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDate

`func (o *MTOShipmentWithoutServiceItemsV2) SetApprovedDate(v string)`

SetApprovedDate sets ApprovedDate field to given value.

### HasApprovedDate

`func (o *MTOShipmentWithoutServiceItemsV2) HasApprovedDate() bool`

HasApprovedDate returns a boolean if a field has been set.

### SetApprovedDateNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetApprovedDateNil(b bool)`

 SetApprovedDateNil sets the value for ApprovedDate to be an explicit nil

### UnsetApprovedDate
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetApprovedDate()`

UnsetApprovedDate ensures that no value is present for ApprovedDate, not even an explicit nil
### GetRequestedPickupDate

`func (o *MTOShipmentWithoutServiceItemsV2) GetRequestedPickupDate() string`

GetRequestedPickupDate returns the RequestedPickupDate field if non-nil, zero value otherwise.

### GetRequestedPickupDateOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetRequestedPickupDateOk() (*string, bool)`

GetRequestedPickupDateOk returns a tuple with the RequestedPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedPickupDate

`func (o *MTOShipmentWithoutServiceItemsV2) SetRequestedPickupDate(v string)`

SetRequestedPickupDate sets RequestedPickupDate field to given value.

### HasRequestedPickupDate

`func (o *MTOShipmentWithoutServiceItemsV2) HasRequestedPickupDate() bool`

HasRequestedPickupDate returns a boolean if a field has been set.

### SetRequestedPickupDateNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetRequestedPickupDateNil(b bool)`

 SetRequestedPickupDateNil sets the value for RequestedPickupDate to be an explicit nil

### UnsetRequestedPickupDate
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetRequestedPickupDate()`

UnsetRequestedPickupDate ensures that no value is present for RequestedPickupDate, not even an explicit nil
### GetRequestedDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) GetRequestedDeliveryDate() string`

GetRequestedDeliveryDate returns the RequestedDeliveryDate field if non-nil, zero value otherwise.

### GetRequestedDeliveryDateOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetRequestedDeliveryDateOk() (*string, bool)`

GetRequestedDeliveryDateOk returns a tuple with the RequestedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) SetRequestedDeliveryDate(v string)`

SetRequestedDeliveryDate sets RequestedDeliveryDate field to given value.

### HasRequestedDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) HasRequestedDeliveryDate() bool`

HasRequestedDeliveryDate returns a boolean if a field has been set.

### SetRequestedDeliveryDateNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetRequestedDeliveryDateNil(b bool)`

 SetRequestedDeliveryDateNil sets the value for RequestedDeliveryDate to be an explicit nil

### UnsetRequestedDeliveryDate
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetRequestedDeliveryDate()`

UnsetRequestedDeliveryDate ensures that no value is present for RequestedDeliveryDate, not even an explicit nil
### GetScheduledPickupDate

`func (o *MTOShipmentWithoutServiceItemsV2) GetScheduledPickupDate() string`

GetScheduledPickupDate returns the ScheduledPickupDate field if non-nil, zero value otherwise.

### GetScheduledPickupDateOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetScheduledPickupDateOk() (*string, bool)`

GetScheduledPickupDateOk returns a tuple with the ScheduledPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPickupDate

`func (o *MTOShipmentWithoutServiceItemsV2) SetScheduledPickupDate(v string)`

SetScheduledPickupDate sets ScheduledPickupDate field to given value.

### HasScheduledPickupDate

`func (o *MTOShipmentWithoutServiceItemsV2) HasScheduledPickupDate() bool`

HasScheduledPickupDate returns a boolean if a field has been set.

### SetScheduledPickupDateNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetScheduledPickupDateNil(b bool)`

 SetScheduledPickupDateNil sets the value for ScheduledPickupDate to be an explicit nil

### UnsetScheduledPickupDate
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetScheduledPickupDate()`

UnsetScheduledPickupDate ensures that no value is present for ScheduledPickupDate, not even an explicit nil
### GetActualPickupDate

`func (o *MTOShipmentWithoutServiceItemsV2) GetActualPickupDate() string`

GetActualPickupDate returns the ActualPickupDate field if non-nil, zero value otherwise.

### GetActualPickupDateOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetActualPickupDateOk() (*string, bool)`

GetActualPickupDateOk returns a tuple with the ActualPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPickupDate

`func (o *MTOShipmentWithoutServiceItemsV2) SetActualPickupDate(v string)`

SetActualPickupDate sets ActualPickupDate field to given value.

### HasActualPickupDate

`func (o *MTOShipmentWithoutServiceItemsV2) HasActualPickupDate() bool`

HasActualPickupDate returns a boolean if a field has been set.

### SetActualPickupDateNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetActualPickupDateNil(b bool)`

 SetActualPickupDateNil sets the value for ActualPickupDate to be an explicit nil

### UnsetActualPickupDate
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetActualPickupDate()`

UnsetActualPickupDate ensures that no value is present for ActualPickupDate, not even an explicit nil
### GetFirstAvailableDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) GetFirstAvailableDeliveryDate() string`

GetFirstAvailableDeliveryDate returns the FirstAvailableDeliveryDate field if non-nil, zero value otherwise.

### GetFirstAvailableDeliveryDateOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetFirstAvailableDeliveryDateOk() (*string, bool)`

GetFirstAvailableDeliveryDateOk returns a tuple with the FirstAvailableDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAvailableDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) SetFirstAvailableDeliveryDate(v string)`

SetFirstAvailableDeliveryDate sets FirstAvailableDeliveryDate field to given value.

### HasFirstAvailableDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) HasFirstAvailableDeliveryDate() bool`

HasFirstAvailableDeliveryDate returns a boolean if a field has been set.

### SetFirstAvailableDeliveryDateNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetFirstAvailableDeliveryDateNil(b bool)`

 SetFirstAvailableDeliveryDateNil sets the value for FirstAvailableDeliveryDate to be an explicit nil

### UnsetFirstAvailableDeliveryDate
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetFirstAvailableDeliveryDate()`

UnsetFirstAvailableDeliveryDate ensures that no value is present for FirstAvailableDeliveryDate, not even an explicit nil
### GetRequiredDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) GetRequiredDeliveryDate() string`

GetRequiredDeliveryDate returns the RequiredDeliveryDate field if non-nil, zero value otherwise.

### GetRequiredDeliveryDateOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetRequiredDeliveryDateOk() (*string, bool)`

GetRequiredDeliveryDateOk returns a tuple with the RequiredDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) SetRequiredDeliveryDate(v string)`

SetRequiredDeliveryDate sets RequiredDeliveryDate field to given value.

### HasRequiredDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) HasRequiredDeliveryDate() bool`

HasRequiredDeliveryDate returns a boolean if a field has been set.

### SetRequiredDeliveryDateNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetRequiredDeliveryDateNil(b bool)`

 SetRequiredDeliveryDateNil sets the value for RequiredDeliveryDate to be an explicit nil

### UnsetRequiredDeliveryDate
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetRequiredDeliveryDate()`

UnsetRequiredDeliveryDate ensures that no value is present for RequiredDeliveryDate, not even an explicit nil
### GetScheduledDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) GetScheduledDeliveryDate() string`

GetScheduledDeliveryDate returns the ScheduledDeliveryDate field if non-nil, zero value otherwise.

### GetScheduledDeliveryDateOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetScheduledDeliveryDateOk() (*string, bool)`

GetScheduledDeliveryDateOk returns a tuple with the ScheduledDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) SetScheduledDeliveryDate(v string)`

SetScheduledDeliveryDate sets ScheduledDeliveryDate field to given value.

### HasScheduledDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) HasScheduledDeliveryDate() bool`

HasScheduledDeliveryDate returns a boolean if a field has been set.

### SetScheduledDeliveryDateNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetScheduledDeliveryDateNil(b bool)`

 SetScheduledDeliveryDateNil sets the value for ScheduledDeliveryDate to be an explicit nil

### UnsetScheduledDeliveryDate
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetScheduledDeliveryDate()`

UnsetScheduledDeliveryDate ensures that no value is present for ScheduledDeliveryDate, not even an explicit nil
### GetActualDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) GetActualDeliveryDate() string`

GetActualDeliveryDate returns the ActualDeliveryDate field if non-nil, zero value otherwise.

### GetActualDeliveryDateOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetActualDeliveryDateOk() (*string, bool)`

GetActualDeliveryDateOk returns a tuple with the ActualDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) SetActualDeliveryDate(v string)`

SetActualDeliveryDate sets ActualDeliveryDate field to given value.

### HasActualDeliveryDate

`func (o *MTOShipmentWithoutServiceItemsV2) HasActualDeliveryDate() bool`

HasActualDeliveryDate returns a boolean if a field has been set.

### SetActualDeliveryDateNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetActualDeliveryDateNil(b bool)`

 SetActualDeliveryDateNil sets the value for ActualDeliveryDate to be an explicit nil

### UnsetActualDeliveryDate
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetActualDeliveryDate()`

UnsetActualDeliveryDate ensures that no value is present for ActualDeliveryDate, not even an explicit nil
### GetPrimeEstimatedWeight

`func (o *MTOShipmentWithoutServiceItemsV2) GetPrimeEstimatedWeight() int32`

GetPrimeEstimatedWeight returns the PrimeEstimatedWeight field if non-nil, zero value otherwise.

### GetPrimeEstimatedWeightOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetPrimeEstimatedWeightOk() (*int32, bool)`

GetPrimeEstimatedWeightOk returns a tuple with the PrimeEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeEstimatedWeight

`func (o *MTOShipmentWithoutServiceItemsV2) SetPrimeEstimatedWeight(v int32)`

SetPrimeEstimatedWeight sets PrimeEstimatedWeight field to given value.

### HasPrimeEstimatedWeight

`func (o *MTOShipmentWithoutServiceItemsV2) HasPrimeEstimatedWeight() bool`

HasPrimeEstimatedWeight returns a boolean if a field has been set.

### SetPrimeEstimatedWeightNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetPrimeEstimatedWeightNil(b bool)`

 SetPrimeEstimatedWeightNil sets the value for PrimeEstimatedWeight to be an explicit nil

### UnsetPrimeEstimatedWeight
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetPrimeEstimatedWeight()`

UnsetPrimeEstimatedWeight ensures that no value is present for PrimeEstimatedWeight, not even an explicit nil
### GetPrimeEstimatedWeightRecordedDate

`func (o *MTOShipmentWithoutServiceItemsV2) GetPrimeEstimatedWeightRecordedDate() string`

GetPrimeEstimatedWeightRecordedDate returns the PrimeEstimatedWeightRecordedDate field if non-nil, zero value otherwise.

### GetPrimeEstimatedWeightRecordedDateOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetPrimeEstimatedWeightRecordedDateOk() (*string, bool)`

GetPrimeEstimatedWeightRecordedDateOk returns a tuple with the PrimeEstimatedWeightRecordedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeEstimatedWeightRecordedDate

`func (o *MTOShipmentWithoutServiceItemsV2) SetPrimeEstimatedWeightRecordedDate(v string)`

SetPrimeEstimatedWeightRecordedDate sets PrimeEstimatedWeightRecordedDate field to given value.

### HasPrimeEstimatedWeightRecordedDate

`func (o *MTOShipmentWithoutServiceItemsV2) HasPrimeEstimatedWeightRecordedDate() bool`

HasPrimeEstimatedWeightRecordedDate returns a boolean if a field has been set.

### SetPrimeEstimatedWeightRecordedDateNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetPrimeEstimatedWeightRecordedDateNil(b bool)`

 SetPrimeEstimatedWeightRecordedDateNil sets the value for PrimeEstimatedWeightRecordedDate to be an explicit nil

### UnsetPrimeEstimatedWeightRecordedDate
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetPrimeEstimatedWeightRecordedDate()`

UnsetPrimeEstimatedWeightRecordedDate ensures that no value is present for PrimeEstimatedWeightRecordedDate, not even an explicit nil
### GetPrimeActualWeight

`func (o *MTOShipmentWithoutServiceItemsV2) GetPrimeActualWeight() int32`

GetPrimeActualWeight returns the PrimeActualWeight field if non-nil, zero value otherwise.

### GetPrimeActualWeightOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetPrimeActualWeightOk() (*int32, bool)`

GetPrimeActualWeightOk returns a tuple with the PrimeActualWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeActualWeight

`func (o *MTOShipmentWithoutServiceItemsV2) SetPrimeActualWeight(v int32)`

SetPrimeActualWeight sets PrimeActualWeight field to given value.

### HasPrimeActualWeight

`func (o *MTOShipmentWithoutServiceItemsV2) HasPrimeActualWeight() bool`

HasPrimeActualWeight returns a boolean if a field has been set.

### SetPrimeActualWeightNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetPrimeActualWeightNil(b bool)`

 SetPrimeActualWeightNil sets the value for PrimeActualWeight to be an explicit nil

### UnsetPrimeActualWeight
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetPrimeActualWeight()`

UnsetPrimeActualWeight ensures that no value is present for PrimeActualWeight, not even an explicit nil
### GetNtsRecordedWeight

`func (o *MTOShipmentWithoutServiceItemsV2) GetNtsRecordedWeight() int32`

GetNtsRecordedWeight returns the NtsRecordedWeight field if non-nil, zero value otherwise.

### GetNtsRecordedWeightOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetNtsRecordedWeightOk() (*int32, bool)`

GetNtsRecordedWeightOk returns a tuple with the NtsRecordedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtsRecordedWeight

`func (o *MTOShipmentWithoutServiceItemsV2) SetNtsRecordedWeight(v int32)`

SetNtsRecordedWeight sets NtsRecordedWeight field to given value.

### HasNtsRecordedWeight

`func (o *MTOShipmentWithoutServiceItemsV2) HasNtsRecordedWeight() bool`

HasNtsRecordedWeight returns a boolean if a field has been set.

### SetNtsRecordedWeightNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetNtsRecordedWeightNil(b bool)`

 SetNtsRecordedWeightNil sets the value for NtsRecordedWeight to be an explicit nil

### UnsetNtsRecordedWeight
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetNtsRecordedWeight()`

UnsetNtsRecordedWeight ensures that no value is present for NtsRecordedWeight, not even an explicit nil
### GetCustomerRemarks

`func (o *MTOShipmentWithoutServiceItemsV2) GetCustomerRemarks() string`

GetCustomerRemarks returns the CustomerRemarks field if non-nil, zero value otherwise.

### GetCustomerRemarksOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetCustomerRemarksOk() (*string, bool)`

GetCustomerRemarksOk returns a tuple with the CustomerRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRemarks

`func (o *MTOShipmentWithoutServiceItemsV2) SetCustomerRemarks(v string)`

SetCustomerRemarks sets CustomerRemarks field to given value.

### HasCustomerRemarks

`func (o *MTOShipmentWithoutServiceItemsV2) HasCustomerRemarks() bool`

HasCustomerRemarks returns a boolean if a field has been set.

### SetCustomerRemarksNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetCustomerRemarksNil(b bool)`

 SetCustomerRemarksNil sets the value for CustomerRemarks to be an explicit nil

### UnsetCustomerRemarks
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetCustomerRemarks()`

UnsetCustomerRemarks ensures that no value is present for CustomerRemarks, not even an explicit nil
### GetCounselorRemarks

`func (o *MTOShipmentWithoutServiceItemsV2) GetCounselorRemarks() string`

GetCounselorRemarks returns the CounselorRemarks field if non-nil, zero value otherwise.

### GetCounselorRemarksOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetCounselorRemarksOk() (*string, bool)`

GetCounselorRemarksOk returns a tuple with the CounselorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounselorRemarks

`func (o *MTOShipmentWithoutServiceItemsV2) SetCounselorRemarks(v string)`

SetCounselorRemarks sets CounselorRemarks field to given value.

### HasCounselorRemarks

`func (o *MTOShipmentWithoutServiceItemsV2) HasCounselorRemarks() bool`

HasCounselorRemarks returns a boolean if a field has been set.

### SetCounselorRemarksNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetCounselorRemarksNil(b bool)`

 SetCounselorRemarksNil sets the value for CounselorRemarks to be an explicit nil

### UnsetCounselorRemarks
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetCounselorRemarks()`

UnsetCounselorRemarks ensures that no value is present for CounselorRemarks, not even an explicit nil
### GetActualProGearWeight

`func (o *MTOShipmentWithoutServiceItemsV2) GetActualProGearWeight() int32`

GetActualProGearWeight returns the ActualProGearWeight field if non-nil, zero value otherwise.

### GetActualProGearWeightOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetActualProGearWeightOk() (*int32, bool)`

GetActualProGearWeightOk returns a tuple with the ActualProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualProGearWeight

`func (o *MTOShipmentWithoutServiceItemsV2) SetActualProGearWeight(v int32)`

SetActualProGearWeight sets ActualProGearWeight field to given value.

### HasActualProGearWeight

`func (o *MTOShipmentWithoutServiceItemsV2) HasActualProGearWeight() bool`

HasActualProGearWeight returns a boolean if a field has been set.

### SetActualProGearWeightNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetActualProGearWeightNil(b bool)`

 SetActualProGearWeightNil sets the value for ActualProGearWeight to be an explicit nil

### UnsetActualProGearWeight
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetActualProGearWeight()`

UnsetActualProGearWeight ensures that no value is present for ActualProGearWeight, not even an explicit nil
### GetActualSpouseProGearWeight

`func (o *MTOShipmentWithoutServiceItemsV2) GetActualSpouseProGearWeight() int32`

GetActualSpouseProGearWeight returns the ActualSpouseProGearWeight field if non-nil, zero value otherwise.

### GetActualSpouseProGearWeightOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetActualSpouseProGearWeightOk() (*int32, bool)`

GetActualSpouseProGearWeightOk returns a tuple with the ActualSpouseProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSpouseProGearWeight

`func (o *MTOShipmentWithoutServiceItemsV2) SetActualSpouseProGearWeight(v int32)`

SetActualSpouseProGearWeight sets ActualSpouseProGearWeight field to given value.

### HasActualSpouseProGearWeight

`func (o *MTOShipmentWithoutServiceItemsV2) HasActualSpouseProGearWeight() bool`

HasActualSpouseProGearWeight returns a boolean if a field has been set.

### SetActualSpouseProGearWeightNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetActualSpouseProGearWeightNil(b bool)`

 SetActualSpouseProGearWeightNil sets the value for ActualSpouseProGearWeight to be an explicit nil

### UnsetActualSpouseProGearWeight
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetActualSpouseProGearWeight()`

UnsetActualSpouseProGearWeight ensures that no value is present for ActualSpouseProGearWeight, not even an explicit nil
### GetAgents

`func (o *MTOShipmentWithoutServiceItemsV2) GetAgents() []MTOAgentV2V2`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetAgentsOk() (*[]MTOAgentV2V2, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *MTOShipmentWithoutServiceItemsV2) SetAgents(v []MTOAgentV2V2)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *MTOShipmentWithoutServiceItemsV2) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetSitExtensions

`func (o *MTOShipmentWithoutServiceItemsV2) GetSitExtensions() []SITExtensionV2V2`

GetSitExtensions returns the SitExtensions field if non-nil, zero value otherwise.

### GetSitExtensionsOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetSitExtensionsOk() (*[]SITExtensionV2V2, bool)`

GetSitExtensionsOk returns a tuple with the SitExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitExtensions

`func (o *MTOShipmentWithoutServiceItemsV2) SetSitExtensions(v []SITExtensionV2V2)`

SetSitExtensions sets SitExtensions field to given value.

### HasSitExtensions

`func (o *MTOShipmentWithoutServiceItemsV2) HasSitExtensions() bool`

HasSitExtensions returns a boolean if a field has been set.

### GetReweigh

`func (o *MTOShipmentWithoutServiceItemsV2) GetReweigh() ReweighV2V2`

GetReweigh returns the Reweigh field if non-nil, zero value otherwise.

### GetReweighOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetReweighOk() (*ReweighV2V2, bool)`

GetReweighOk returns a tuple with the Reweigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReweigh

`func (o *MTOShipmentWithoutServiceItemsV2) SetReweigh(v ReweighV2V2)`

SetReweigh sets Reweigh field to given value.

### HasReweigh

`func (o *MTOShipmentWithoutServiceItemsV2) HasReweigh() bool`

HasReweigh returns a boolean if a field has been set.

### GetPickupAddress

`func (o *MTOShipmentWithoutServiceItemsV2) GetPickupAddress() AddressV2V2`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetPickupAddressOk() (*AddressV2V2, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *MTOShipmentWithoutServiceItemsV2) SetPickupAddress(v AddressV2V2)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *MTOShipmentWithoutServiceItemsV2) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *MTOShipmentWithoutServiceItemsV2) GetDestinationAddress() AddressV2V2`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetDestinationAddressOk() (*AddressV2V2, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *MTOShipmentWithoutServiceItemsV2) SetDestinationAddress(v AddressV2V2)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *MTOShipmentWithoutServiceItemsV2) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetDestinationType

`func (o *MTOShipmentWithoutServiceItemsV2) GetDestinationType() DestinationTypeV2V2`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetDestinationTypeOk() (*DestinationTypeV2V2, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *MTOShipmentWithoutServiceItemsV2) SetDestinationType(v DestinationTypeV2V2)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *MTOShipmentWithoutServiceItemsV2) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### SetDestinationTypeNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetDestinationTypeNil(b bool)`

 SetDestinationTypeNil sets the value for DestinationType to be an explicit nil

### UnsetDestinationType
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetDestinationType()`

UnsetDestinationType ensures that no value is present for DestinationType, not even an explicit nil
### GetSecondaryPickupAddress

`func (o *MTOShipmentWithoutServiceItemsV2) GetSecondaryPickupAddress() AddressV2V2`

GetSecondaryPickupAddress returns the SecondaryPickupAddress field if non-nil, zero value otherwise.

### GetSecondaryPickupAddressOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetSecondaryPickupAddressOk() (*AddressV2V2, bool)`

GetSecondaryPickupAddressOk returns a tuple with the SecondaryPickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryPickupAddress

`func (o *MTOShipmentWithoutServiceItemsV2) SetSecondaryPickupAddress(v AddressV2V2)`

SetSecondaryPickupAddress sets SecondaryPickupAddress field to given value.

### HasSecondaryPickupAddress

`func (o *MTOShipmentWithoutServiceItemsV2) HasSecondaryPickupAddress() bool`

HasSecondaryPickupAddress returns a boolean if a field has been set.

### GetSecondaryDeliveryAddress

`func (o *MTOShipmentWithoutServiceItemsV2) GetSecondaryDeliveryAddress() AddressV2V2`

GetSecondaryDeliveryAddress returns the SecondaryDeliveryAddress field if non-nil, zero value otherwise.

### GetSecondaryDeliveryAddressOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetSecondaryDeliveryAddressOk() (*AddressV2V2, bool)`

GetSecondaryDeliveryAddressOk returns a tuple with the SecondaryDeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDeliveryAddress

`func (o *MTOShipmentWithoutServiceItemsV2) SetSecondaryDeliveryAddress(v AddressV2V2)`

SetSecondaryDeliveryAddress sets SecondaryDeliveryAddress field to given value.

### HasSecondaryDeliveryAddress

`func (o *MTOShipmentWithoutServiceItemsV2) HasSecondaryDeliveryAddress() bool`

HasSecondaryDeliveryAddress returns a boolean if a field has been set.

### GetStorageFacility

`func (o *MTOShipmentWithoutServiceItemsV2) GetStorageFacility() UpdateMTOShipmentStorageFacilityV2V2`

GetStorageFacility returns the StorageFacility field if non-nil, zero value otherwise.

### GetStorageFacilityOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetStorageFacilityOk() (*UpdateMTOShipmentStorageFacilityV2V2, bool)`

GetStorageFacilityOk returns a tuple with the StorageFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFacility

`func (o *MTOShipmentWithoutServiceItemsV2) SetStorageFacility(v UpdateMTOShipmentStorageFacilityV2V2)`

SetStorageFacility sets StorageFacility field to given value.

### HasStorageFacility

`func (o *MTOShipmentWithoutServiceItemsV2) HasStorageFacility() bool`

HasStorageFacility returns a boolean if a field has been set.

### GetShipmentType

`func (o *MTOShipmentWithoutServiceItemsV2) GetShipmentType() MTOShipmentTypeV2V2`

GetShipmentType returns the ShipmentType field if non-nil, zero value otherwise.

### GetShipmentTypeOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetShipmentTypeOk() (*MTOShipmentTypeV2V2, bool)`

GetShipmentTypeOk returns a tuple with the ShipmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentType

`func (o *MTOShipmentWithoutServiceItemsV2) SetShipmentType(v MTOShipmentTypeV2V2)`

SetShipmentType sets ShipmentType field to given value.

### HasShipmentType

`func (o *MTOShipmentWithoutServiceItemsV2) HasShipmentType() bool`

HasShipmentType returns a boolean if a field has been set.

### GetDiversion

`func (o *MTOShipmentWithoutServiceItemsV2) GetDiversion() bool`

GetDiversion returns the Diversion field if non-nil, zero value otherwise.

### GetDiversionOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetDiversionOk() (*bool, bool)`

GetDiversionOk returns a tuple with the Diversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiversion

`func (o *MTOShipmentWithoutServiceItemsV2) SetDiversion(v bool)`

SetDiversion sets Diversion field to given value.

### HasDiversion

`func (o *MTOShipmentWithoutServiceItemsV2) HasDiversion() bool`

HasDiversion returns a boolean if a field has been set.

### GetDiversionReason

`func (o *MTOShipmentWithoutServiceItemsV2) GetDiversionReason() string`

GetDiversionReason returns the DiversionReason field if non-nil, zero value otherwise.

### GetDiversionReasonOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetDiversionReasonOk() (*string, bool)`

GetDiversionReasonOk returns a tuple with the DiversionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiversionReason

`func (o *MTOShipmentWithoutServiceItemsV2) SetDiversionReason(v string)`

SetDiversionReason sets DiversionReason field to given value.

### HasDiversionReason

`func (o *MTOShipmentWithoutServiceItemsV2) HasDiversionReason() bool`

HasDiversionReason returns a boolean if a field has been set.

### SetDiversionReasonNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetDiversionReasonNil(b bool)`

 SetDiversionReasonNil sets the value for DiversionReason to be an explicit nil

### UnsetDiversionReason
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetDiversionReason()`

UnsetDiversionReason ensures that no value is present for DiversionReason, not even an explicit nil
### GetStatus

`func (o *MTOShipmentWithoutServiceItemsV2) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MTOShipmentWithoutServiceItemsV2) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MTOShipmentWithoutServiceItemsV2) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPpmShipment

`func (o *MTOShipmentWithoutServiceItemsV2) GetPpmShipment() PPMShipmentV2V2`

GetPpmShipment returns the PpmShipment field if non-nil, zero value otherwise.

### GetPpmShipmentOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetPpmShipmentOk() (*PPMShipmentV2V2, bool)`

GetPpmShipmentOk returns a tuple with the PpmShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpmShipment

`func (o *MTOShipmentWithoutServiceItemsV2) SetPpmShipment(v PPMShipmentV2V2)`

SetPpmShipment sets PpmShipment field to given value.

### HasPpmShipment

`func (o *MTOShipmentWithoutServiceItemsV2) HasPpmShipment() bool`

HasPpmShipment returns a boolean if a field has been set.

### SetPpmShipmentNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetPpmShipmentNil(b bool)`

 SetPpmShipmentNil sets the value for PpmShipment to be an explicit nil

### UnsetPpmShipment
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetPpmShipment()`

UnsetPpmShipment ensures that no value is present for PpmShipment, not even an explicit nil
### GetDeliveryAddressUpdate

`func (o *MTOShipmentWithoutServiceItemsV2) GetDeliveryAddressUpdate() ShipmentAddressUpdateV2V2`

GetDeliveryAddressUpdate returns the DeliveryAddressUpdate field if non-nil, zero value otherwise.

### GetDeliveryAddressUpdateOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetDeliveryAddressUpdateOk() (*ShipmentAddressUpdateV2V2, bool)`

GetDeliveryAddressUpdateOk returns a tuple with the DeliveryAddressUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddressUpdate

`func (o *MTOShipmentWithoutServiceItemsV2) SetDeliveryAddressUpdate(v ShipmentAddressUpdateV2V2)`

SetDeliveryAddressUpdate sets DeliveryAddressUpdate field to given value.

### HasDeliveryAddressUpdate

`func (o *MTOShipmentWithoutServiceItemsV2) HasDeliveryAddressUpdate() bool`

HasDeliveryAddressUpdate returns a boolean if a field has been set.

### GetETag

`func (o *MTOShipmentWithoutServiceItemsV2) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MTOShipmentWithoutServiceItemsV2) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MTOShipmentWithoutServiceItemsV2) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MTOShipmentWithoutServiceItemsV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MTOShipmentWithoutServiceItemsV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MTOShipmentWithoutServiceItemsV2) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MTOShipmentWithoutServiceItemsV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MTOShipmentWithoutServiceItemsV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MTOShipmentWithoutServiceItemsV2) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPointOfContact

`func (o *MTOShipmentWithoutServiceItemsV2) GetPointOfContact() string`

GetPointOfContact returns the PointOfContact field if non-nil, zero value otherwise.

### GetPointOfContactOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetPointOfContactOk() (*string, bool)`

GetPointOfContactOk returns a tuple with the PointOfContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointOfContact

`func (o *MTOShipmentWithoutServiceItemsV2) SetPointOfContact(v string)`

SetPointOfContact sets PointOfContact field to given value.

### HasPointOfContact

`func (o *MTOShipmentWithoutServiceItemsV2) HasPointOfContact() bool`

HasPointOfContact returns a boolean if a field has been set.

### GetOriginSitAuthEndDate

`func (o *MTOShipmentWithoutServiceItemsV2) GetOriginSitAuthEndDate() string`

GetOriginSitAuthEndDate returns the OriginSitAuthEndDate field if non-nil, zero value otherwise.

### GetOriginSitAuthEndDateOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetOriginSitAuthEndDateOk() (*string, bool)`

GetOriginSitAuthEndDateOk returns a tuple with the OriginSitAuthEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginSitAuthEndDate

`func (o *MTOShipmentWithoutServiceItemsV2) SetOriginSitAuthEndDate(v string)`

SetOriginSitAuthEndDate sets OriginSitAuthEndDate field to given value.

### HasOriginSitAuthEndDate

`func (o *MTOShipmentWithoutServiceItemsV2) HasOriginSitAuthEndDate() bool`

HasOriginSitAuthEndDate returns a boolean if a field has been set.

### SetOriginSitAuthEndDateNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetOriginSitAuthEndDateNil(b bool)`

 SetOriginSitAuthEndDateNil sets the value for OriginSitAuthEndDate to be an explicit nil

### UnsetOriginSitAuthEndDate
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetOriginSitAuthEndDate()`

UnsetOriginSitAuthEndDate ensures that no value is present for OriginSitAuthEndDate, not even an explicit nil
### GetDestinationSitAuthEndDate

`func (o *MTOShipmentWithoutServiceItemsV2) GetDestinationSitAuthEndDate() string`

GetDestinationSitAuthEndDate returns the DestinationSitAuthEndDate field if non-nil, zero value otherwise.

### GetDestinationSitAuthEndDateOk

`func (o *MTOShipmentWithoutServiceItemsV2) GetDestinationSitAuthEndDateOk() (*string, bool)`

GetDestinationSitAuthEndDateOk returns a tuple with the DestinationSitAuthEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationSitAuthEndDate

`func (o *MTOShipmentWithoutServiceItemsV2) SetDestinationSitAuthEndDate(v string)`

SetDestinationSitAuthEndDate sets DestinationSitAuthEndDate field to given value.

### HasDestinationSitAuthEndDate

`func (o *MTOShipmentWithoutServiceItemsV2) HasDestinationSitAuthEndDate() bool`

HasDestinationSitAuthEndDate returns a boolean if a field has been set.

### SetDestinationSitAuthEndDateNil

`func (o *MTOShipmentWithoutServiceItemsV2) SetDestinationSitAuthEndDateNil(b bool)`

 SetDestinationSitAuthEndDateNil sets the value for DestinationSitAuthEndDate to be an explicit nil

### UnsetDestinationSitAuthEndDate
`func (o *MTOShipmentWithoutServiceItemsV2) UnsetDestinationSitAuthEndDate()`

UnsetDestinationSitAuthEndDate ensures that no value is present for DestinationSitAuthEndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


