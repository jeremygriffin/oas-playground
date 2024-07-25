# MTOShipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MtoServiceItems** | Pointer to [**[]MTOServiceItem**](MTOServiceItem.md) | A list of service items connected to this shipment. | [optional] [readonly] 
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
**Agents** | Pointer to [**[]MTOAgent**](MTOAgent.md) | A list of the agents for a shipment. Agents are the people who the Prime contractor recognize as permitted to release (in the case of pickup) or receive (on delivery) a shipment.  | [optional] 
**SitExtensions** | Pointer to [**[]SITExtension**](SITExtension.md) |  | [optional] 
**Reweigh** | Pointer to [**Reweigh**](Reweigh.md) |  | [optional] 
**PickupAddress** | Pointer to [**Address**](Address.md) | The address where the movers should pick up this shipment, entered by the customer during onboarding when they enter shipment details.  | [optional] 
**DestinationAddress** | Pointer to [**Address**](Address.md) | Where the movers should deliver this shipment. Often provided by the customer when they enter shipment details during onboarding, if they know their new address already.  May be blank when entered by the customer, required when entered by the Prime. May not represent the true final destination due to the shipment being diverted or placed in SIT.  | [optional] 
**DestinationType** | Pointer to [**NullableDestinationType**](DestinationType.md) |  | [optional] 
**SecondaryPickupAddress** | Pointer to [**Address**](Address.md) | A second pickup address for this shipment, if the customer entered one. An optional field. | [optional] 
**SecondaryDeliveryAddress** | Pointer to [**Address**](Address.md) | A second delivery address for this shipment, if the customer entered one. An optional field. | [optional] 
**StorageFacility** | Pointer to [**UpdateMTOShipmentStorageFacility**](UpdateMTOShipmentStorageFacility.md) |  | [optional] 
**ShipmentType** | Pointer to [**MTOShipmentType**](MTOShipmentType.md) |  | [optional] 
**Diversion** | Pointer to **bool** | This value indicates whether or not this shipment is part of a diversion. If yes, the shipment can be either the starting or ending segment of the diversion.  | [optional] 
**DiversionReason** | Pointer to **NullableString** | The reason the TOO provided when requesting a diversion for this shipment.  | [optional] [readonly] 
**Status** | Pointer to **string** | The status of a shipment, indicating where it is in the TOO&#39;s approval process. Can only be updated by the contractor in special circumstances.  | [optional] [readonly] 
**PpmShipment** | Pointer to [**NullablePPMShipment**](PPMShipment.md) |  | [optional] 
**DeliveryAddressUpdate** | Pointer to [**ShipmentAddressUpdate**](ShipmentAddressUpdate.md) |  | [optional] 
**ETag** | Pointer to **string** | A hash unique to this shipment that should be used as the \&quot;If-Match\&quot; header for any updates. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**PointOfContact** | Pointer to **string** | Email or ID of the person who will be contacted in the event of questions or concerns about this update. May be the person performing the update, or someone else working with the Prime contractor.  | [optional] 
**OriginSitAuthEndDate** | Pointer to **NullableString** | The SIT authorized end date for origin SIT. | [optional] 
**DestinationSitAuthEndDate** | Pointer to **NullableString** | The SIT authorized end date for destination SIT. | [optional] 

## Methods

### NewMTOShipment

`func NewMTOShipment() *MTOShipment`

NewMTOShipment instantiates a new MTOShipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTOShipmentWithDefaults

`func NewMTOShipmentWithDefaults() *MTOShipment`

NewMTOShipmentWithDefaults instantiates a new MTOShipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMtoServiceItems

`func (o *MTOShipment) GetMtoServiceItems() []MTOServiceItem`

GetMtoServiceItems returns the MtoServiceItems field if non-nil, zero value otherwise.

### GetMtoServiceItemsOk

`func (o *MTOShipment) GetMtoServiceItemsOk() (*[]MTOServiceItem, bool)`

GetMtoServiceItemsOk returns a tuple with the MtoServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoServiceItems

`func (o *MTOShipment) SetMtoServiceItems(v []MTOServiceItem)`

SetMtoServiceItems sets MtoServiceItems field to given value.

### HasMtoServiceItems

`func (o *MTOShipment) HasMtoServiceItems() bool`

HasMtoServiceItems returns a boolean if a field has been set.

### GetId

`func (o *MTOShipment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MTOShipment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MTOShipment) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MTOShipment) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMoveTaskOrderID

`func (o *MTOShipment) GetMoveTaskOrderID() string`

GetMoveTaskOrderID returns the MoveTaskOrderID field if non-nil, zero value otherwise.

### GetMoveTaskOrderIDOk

`func (o *MTOShipment) GetMoveTaskOrderIDOk() (*string, bool)`

GetMoveTaskOrderIDOk returns a tuple with the MoveTaskOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTaskOrderID

`func (o *MTOShipment) SetMoveTaskOrderID(v string)`

SetMoveTaskOrderID sets MoveTaskOrderID field to given value.

### HasMoveTaskOrderID

`func (o *MTOShipment) HasMoveTaskOrderID() bool`

HasMoveTaskOrderID returns a boolean if a field has been set.

### GetApprovedDate

`func (o *MTOShipment) GetApprovedDate() string`

GetApprovedDate returns the ApprovedDate field if non-nil, zero value otherwise.

### GetApprovedDateOk

`func (o *MTOShipment) GetApprovedDateOk() (*string, bool)`

GetApprovedDateOk returns a tuple with the ApprovedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDate

`func (o *MTOShipment) SetApprovedDate(v string)`

SetApprovedDate sets ApprovedDate field to given value.

### HasApprovedDate

`func (o *MTOShipment) HasApprovedDate() bool`

HasApprovedDate returns a boolean if a field has been set.

### SetApprovedDateNil

`func (o *MTOShipment) SetApprovedDateNil(b bool)`

 SetApprovedDateNil sets the value for ApprovedDate to be an explicit nil

### UnsetApprovedDate
`func (o *MTOShipment) UnsetApprovedDate()`

UnsetApprovedDate ensures that no value is present for ApprovedDate, not even an explicit nil
### GetRequestedPickupDate

`func (o *MTOShipment) GetRequestedPickupDate() string`

GetRequestedPickupDate returns the RequestedPickupDate field if non-nil, zero value otherwise.

### GetRequestedPickupDateOk

`func (o *MTOShipment) GetRequestedPickupDateOk() (*string, bool)`

GetRequestedPickupDateOk returns a tuple with the RequestedPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedPickupDate

`func (o *MTOShipment) SetRequestedPickupDate(v string)`

SetRequestedPickupDate sets RequestedPickupDate field to given value.

### HasRequestedPickupDate

`func (o *MTOShipment) HasRequestedPickupDate() bool`

HasRequestedPickupDate returns a boolean if a field has been set.

### SetRequestedPickupDateNil

`func (o *MTOShipment) SetRequestedPickupDateNil(b bool)`

 SetRequestedPickupDateNil sets the value for RequestedPickupDate to be an explicit nil

### UnsetRequestedPickupDate
`func (o *MTOShipment) UnsetRequestedPickupDate()`

UnsetRequestedPickupDate ensures that no value is present for RequestedPickupDate, not even an explicit nil
### GetRequestedDeliveryDate

`func (o *MTOShipment) GetRequestedDeliveryDate() string`

GetRequestedDeliveryDate returns the RequestedDeliveryDate field if non-nil, zero value otherwise.

### GetRequestedDeliveryDateOk

`func (o *MTOShipment) GetRequestedDeliveryDateOk() (*string, bool)`

GetRequestedDeliveryDateOk returns a tuple with the RequestedDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDeliveryDate

`func (o *MTOShipment) SetRequestedDeliveryDate(v string)`

SetRequestedDeliveryDate sets RequestedDeliveryDate field to given value.

### HasRequestedDeliveryDate

`func (o *MTOShipment) HasRequestedDeliveryDate() bool`

HasRequestedDeliveryDate returns a boolean if a field has been set.

### SetRequestedDeliveryDateNil

`func (o *MTOShipment) SetRequestedDeliveryDateNil(b bool)`

 SetRequestedDeliveryDateNil sets the value for RequestedDeliveryDate to be an explicit nil

### UnsetRequestedDeliveryDate
`func (o *MTOShipment) UnsetRequestedDeliveryDate()`

UnsetRequestedDeliveryDate ensures that no value is present for RequestedDeliveryDate, not even an explicit nil
### GetScheduledPickupDate

`func (o *MTOShipment) GetScheduledPickupDate() string`

GetScheduledPickupDate returns the ScheduledPickupDate field if non-nil, zero value otherwise.

### GetScheduledPickupDateOk

`func (o *MTOShipment) GetScheduledPickupDateOk() (*string, bool)`

GetScheduledPickupDateOk returns a tuple with the ScheduledPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPickupDate

`func (o *MTOShipment) SetScheduledPickupDate(v string)`

SetScheduledPickupDate sets ScheduledPickupDate field to given value.

### HasScheduledPickupDate

`func (o *MTOShipment) HasScheduledPickupDate() bool`

HasScheduledPickupDate returns a boolean if a field has been set.

### SetScheduledPickupDateNil

`func (o *MTOShipment) SetScheduledPickupDateNil(b bool)`

 SetScheduledPickupDateNil sets the value for ScheduledPickupDate to be an explicit nil

### UnsetScheduledPickupDate
`func (o *MTOShipment) UnsetScheduledPickupDate()`

UnsetScheduledPickupDate ensures that no value is present for ScheduledPickupDate, not even an explicit nil
### GetActualPickupDate

`func (o *MTOShipment) GetActualPickupDate() string`

GetActualPickupDate returns the ActualPickupDate field if non-nil, zero value otherwise.

### GetActualPickupDateOk

`func (o *MTOShipment) GetActualPickupDateOk() (*string, bool)`

GetActualPickupDateOk returns a tuple with the ActualPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPickupDate

`func (o *MTOShipment) SetActualPickupDate(v string)`

SetActualPickupDate sets ActualPickupDate field to given value.

### HasActualPickupDate

`func (o *MTOShipment) HasActualPickupDate() bool`

HasActualPickupDate returns a boolean if a field has been set.

### SetActualPickupDateNil

`func (o *MTOShipment) SetActualPickupDateNil(b bool)`

 SetActualPickupDateNil sets the value for ActualPickupDate to be an explicit nil

### UnsetActualPickupDate
`func (o *MTOShipment) UnsetActualPickupDate()`

UnsetActualPickupDate ensures that no value is present for ActualPickupDate, not even an explicit nil
### GetFirstAvailableDeliveryDate

`func (o *MTOShipment) GetFirstAvailableDeliveryDate() string`

GetFirstAvailableDeliveryDate returns the FirstAvailableDeliveryDate field if non-nil, zero value otherwise.

### GetFirstAvailableDeliveryDateOk

`func (o *MTOShipment) GetFirstAvailableDeliveryDateOk() (*string, bool)`

GetFirstAvailableDeliveryDateOk returns a tuple with the FirstAvailableDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAvailableDeliveryDate

`func (o *MTOShipment) SetFirstAvailableDeliveryDate(v string)`

SetFirstAvailableDeliveryDate sets FirstAvailableDeliveryDate field to given value.

### HasFirstAvailableDeliveryDate

`func (o *MTOShipment) HasFirstAvailableDeliveryDate() bool`

HasFirstAvailableDeliveryDate returns a boolean if a field has been set.

### SetFirstAvailableDeliveryDateNil

`func (o *MTOShipment) SetFirstAvailableDeliveryDateNil(b bool)`

 SetFirstAvailableDeliveryDateNil sets the value for FirstAvailableDeliveryDate to be an explicit nil

### UnsetFirstAvailableDeliveryDate
`func (o *MTOShipment) UnsetFirstAvailableDeliveryDate()`

UnsetFirstAvailableDeliveryDate ensures that no value is present for FirstAvailableDeliveryDate, not even an explicit nil
### GetRequiredDeliveryDate

`func (o *MTOShipment) GetRequiredDeliveryDate() string`

GetRequiredDeliveryDate returns the RequiredDeliveryDate field if non-nil, zero value otherwise.

### GetRequiredDeliveryDateOk

`func (o *MTOShipment) GetRequiredDeliveryDateOk() (*string, bool)`

GetRequiredDeliveryDateOk returns a tuple with the RequiredDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredDeliveryDate

`func (o *MTOShipment) SetRequiredDeliveryDate(v string)`

SetRequiredDeliveryDate sets RequiredDeliveryDate field to given value.

### HasRequiredDeliveryDate

`func (o *MTOShipment) HasRequiredDeliveryDate() bool`

HasRequiredDeliveryDate returns a boolean if a field has been set.

### SetRequiredDeliveryDateNil

`func (o *MTOShipment) SetRequiredDeliveryDateNil(b bool)`

 SetRequiredDeliveryDateNil sets the value for RequiredDeliveryDate to be an explicit nil

### UnsetRequiredDeliveryDate
`func (o *MTOShipment) UnsetRequiredDeliveryDate()`

UnsetRequiredDeliveryDate ensures that no value is present for RequiredDeliveryDate, not even an explicit nil
### GetScheduledDeliveryDate

`func (o *MTOShipment) GetScheduledDeliveryDate() string`

GetScheduledDeliveryDate returns the ScheduledDeliveryDate field if non-nil, zero value otherwise.

### GetScheduledDeliveryDateOk

`func (o *MTOShipment) GetScheduledDeliveryDateOk() (*string, bool)`

GetScheduledDeliveryDateOk returns a tuple with the ScheduledDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeliveryDate

`func (o *MTOShipment) SetScheduledDeliveryDate(v string)`

SetScheduledDeliveryDate sets ScheduledDeliveryDate field to given value.

### HasScheduledDeliveryDate

`func (o *MTOShipment) HasScheduledDeliveryDate() bool`

HasScheduledDeliveryDate returns a boolean if a field has been set.

### SetScheduledDeliveryDateNil

`func (o *MTOShipment) SetScheduledDeliveryDateNil(b bool)`

 SetScheduledDeliveryDateNil sets the value for ScheduledDeliveryDate to be an explicit nil

### UnsetScheduledDeliveryDate
`func (o *MTOShipment) UnsetScheduledDeliveryDate()`

UnsetScheduledDeliveryDate ensures that no value is present for ScheduledDeliveryDate, not even an explicit nil
### GetActualDeliveryDate

`func (o *MTOShipment) GetActualDeliveryDate() string`

GetActualDeliveryDate returns the ActualDeliveryDate field if non-nil, zero value otherwise.

### GetActualDeliveryDateOk

`func (o *MTOShipment) GetActualDeliveryDateOk() (*string, bool)`

GetActualDeliveryDateOk returns a tuple with the ActualDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDeliveryDate

`func (o *MTOShipment) SetActualDeliveryDate(v string)`

SetActualDeliveryDate sets ActualDeliveryDate field to given value.

### HasActualDeliveryDate

`func (o *MTOShipment) HasActualDeliveryDate() bool`

HasActualDeliveryDate returns a boolean if a field has been set.

### SetActualDeliveryDateNil

`func (o *MTOShipment) SetActualDeliveryDateNil(b bool)`

 SetActualDeliveryDateNil sets the value for ActualDeliveryDate to be an explicit nil

### UnsetActualDeliveryDate
`func (o *MTOShipment) UnsetActualDeliveryDate()`

UnsetActualDeliveryDate ensures that no value is present for ActualDeliveryDate, not even an explicit nil
### GetPrimeEstimatedWeight

`func (o *MTOShipment) GetPrimeEstimatedWeight() int32`

GetPrimeEstimatedWeight returns the PrimeEstimatedWeight field if non-nil, zero value otherwise.

### GetPrimeEstimatedWeightOk

`func (o *MTOShipment) GetPrimeEstimatedWeightOk() (*int32, bool)`

GetPrimeEstimatedWeightOk returns a tuple with the PrimeEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeEstimatedWeight

`func (o *MTOShipment) SetPrimeEstimatedWeight(v int32)`

SetPrimeEstimatedWeight sets PrimeEstimatedWeight field to given value.

### HasPrimeEstimatedWeight

`func (o *MTOShipment) HasPrimeEstimatedWeight() bool`

HasPrimeEstimatedWeight returns a boolean if a field has been set.

### SetPrimeEstimatedWeightNil

`func (o *MTOShipment) SetPrimeEstimatedWeightNil(b bool)`

 SetPrimeEstimatedWeightNil sets the value for PrimeEstimatedWeight to be an explicit nil

### UnsetPrimeEstimatedWeight
`func (o *MTOShipment) UnsetPrimeEstimatedWeight()`

UnsetPrimeEstimatedWeight ensures that no value is present for PrimeEstimatedWeight, not even an explicit nil
### GetPrimeEstimatedWeightRecordedDate

`func (o *MTOShipment) GetPrimeEstimatedWeightRecordedDate() string`

GetPrimeEstimatedWeightRecordedDate returns the PrimeEstimatedWeightRecordedDate field if non-nil, zero value otherwise.

### GetPrimeEstimatedWeightRecordedDateOk

`func (o *MTOShipment) GetPrimeEstimatedWeightRecordedDateOk() (*string, bool)`

GetPrimeEstimatedWeightRecordedDateOk returns a tuple with the PrimeEstimatedWeightRecordedDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeEstimatedWeightRecordedDate

`func (o *MTOShipment) SetPrimeEstimatedWeightRecordedDate(v string)`

SetPrimeEstimatedWeightRecordedDate sets PrimeEstimatedWeightRecordedDate field to given value.

### HasPrimeEstimatedWeightRecordedDate

`func (o *MTOShipment) HasPrimeEstimatedWeightRecordedDate() bool`

HasPrimeEstimatedWeightRecordedDate returns a boolean if a field has been set.

### SetPrimeEstimatedWeightRecordedDateNil

`func (o *MTOShipment) SetPrimeEstimatedWeightRecordedDateNil(b bool)`

 SetPrimeEstimatedWeightRecordedDateNil sets the value for PrimeEstimatedWeightRecordedDate to be an explicit nil

### UnsetPrimeEstimatedWeightRecordedDate
`func (o *MTOShipment) UnsetPrimeEstimatedWeightRecordedDate()`

UnsetPrimeEstimatedWeightRecordedDate ensures that no value is present for PrimeEstimatedWeightRecordedDate, not even an explicit nil
### GetPrimeActualWeight

`func (o *MTOShipment) GetPrimeActualWeight() int32`

GetPrimeActualWeight returns the PrimeActualWeight field if non-nil, zero value otherwise.

### GetPrimeActualWeightOk

`func (o *MTOShipment) GetPrimeActualWeightOk() (*int32, bool)`

GetPrimeActualWeightOk returns a tuple with the PrimeActualWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeActualWeight

`func (o *MTOShipment) SetPrimeActualWeight(v int32)`

SetPrimeActualWeight sets PrimeActualWeight field to given value.

### HasPrimeActualWeight

`func (o *MTOShipment) HasPrimeActualWeight() bool`

HasPrimeActualWeight returns a boolean if a field has been set.

### SetPrimeActualWeightNil

`func (o *MTOShipment) SetPrimeActualWeightNil(b bool)`

 SetPrimeActualWeightNil sets the value for PrimeActualWeight to be an explicit nil

### UnsetPrimeActualWeight
`func (o *MTOShipment) UnsetPrimeActualWeight()`

UnsetPrimeActualWeight ensures that no value is present for PrimeActualWeight, not even an explicit nil
### GetNtsRecordedWeight

`func (o *MTOShipment) GetNtsRecordedWeight() int32`

GetNtsRecordedWeight returns the NtsRecordedWeight field if non-nil, zero value otherwise.

### GetNtsRecordedWeightOk

`func (o *MTOShipment) GetNtsRecordedWeightOk() (*int32, bool)`

GetNtsRecordedWeightOk returns a tuple with the NtsRecordedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtsRecordedWeight

`func (o *MTOShipment) SetNtsRecordedWeight(v int32)`

SetNtsRecordedWeight sets NtsRecordedWeight field to given value.

### HasNtsRecordedWeight

`func (o *MTOShipment) HasNtsRecordedWeight() bool`

HasNtsRecordedWeight returns a boolean if a field has been set.

### SetNtsRecordedWeightNil

`func (o *MTOShipment) SetNtsRecordedWeightNil(b bool)`

 SetNtsRecordedWeightNil sets the value for NtsRecordedWeight to be an explicit nil

### UnsetNtsRecordedWeight
`func (o *MTOShipment) UnsetNtsRecordedWeight()`

UnsetNtsRecordedWeight ensures that no value is present for NtsRecordedWeight, not even an explicit nil
### GetCustomerRemarks

`func (o *MTOShipment) GetCustomerRemarks() string`

GetCustomerRemarks returns the CustomerRemarks field if non-nil, zero value otherwise.

### GetCustomerRemarksOk

`func (o *MTOShipment) GetCustomerRemarksOk() (*string, bool)`

GetCustomerRemarksOk returns a tuple with the CustomerRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerRemarks

`func (o *MTOShipment) SetCustomerRemarks(v string)`

SetCustomerRemarks sets CustomerRemarks field to given value.

### HasCustomerRemarks

`func (o *MTOShipment) HasCustomerRemarks() bool`

HasCustomerRemarks returns a boolean if a field has been set.

### SetCustomerRemarksNil

`func (o *MTOShipment) SetCustomerRemarksNil(b bool)`

 SetCustomerRemarksNil sets the value for CustomerRemarks to be an explicit nil

### UnsetCustomerRemarks
`func (o *MTOShipment) UnsetCustomerRemarks()`

UnsetCustomerRemarks ensures that no value is present for CustomerRemarks, not even an explicit nil
### GetCounselorRemarks

`func (o *MTOShipment) GetCounselorRemarks() string`

GetCounselorRemarks returns the CounselorRemarks field if non-nil, zero value otherwise.

### GetCounselorRemarksOk

`func (o *MTOShipment) GetCounselorRemarksOk() (*string, bool)`

GetCounselorRemarksOk returns a tuple with the CounselorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounselorRemarks

`func (o *MTOShipment) SetCounselorRemarks(v string)`

SetCounselorRemarks sets CounselorRemarks field to given value.

### HasCounselorRemarks

`func (o *MTOShipment) HasCounselorRemarks() bool`

HasCounselorRemarks returns a boolean if a field has been set.

### SetCounselorRemarksNil

`func (o *MTOShipment) SetCounselorRemarksNil(b bool)`

 SetCounselorRemarksNil sets the value for CounselorRemarks to be an explicit nil

### UnsetCounselorRemarks
`func (o *MTOShipment) UnsetCounselorRemarks()`

UnsetCounselorRemarks ensures that no value is present for CounselorRemarks, not even an explicit nil
### GetActualProGearWeight

`func (o *MTOShipment) GetActualProGearWeight() int32`

GetActualProGearWeight returns the ActualProGearWeight field if non-nil, zero value otherwise.

### GetActualProGearWeightOk

`func (o *MTOShipment) GetActualProGearWeightOk() (*int32, bool)`

GetActualProGearWeightOk returns a tuple with the ActualProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualProGearWeight

`func (o *MTOShipment) SetActualProGearWeight(v int32)`

SetActualProGearWeight sets ActualProGearWeight field to given value.

### HasActualProGearWeight

`func (o *MTOShipment) HasActualProGearWeight() bool`

HasActualProGearWeight returns a boolean if a field has been set.

### SetActualProGearWeightNil

`func (o *MTOShipment) SetActualProGearWeightNil(b bool)`

 SetActualProGearWeightNil sets the value for ActualProGearWeight to be an explicit nil

### UnsetActualProGearWeight
`func (o *MTOShipment) UnsetActualProGearWeight()`

UnsetActualProGearWeight ensures that no value is present for ActualProGearWeight, not even an explicit nil
### GetActualSpouseProGearWeight

`func (o *MTOShipment) GetActualSpouseProGearWeight() int32`

GetActualSpouseProGearWeight returns the ActualSpouseProGearWeight field if non-nil, zero value otherwise.

### GetActualSpouseProGearWeightOk

`func (o *MTOShipment) GetActualSpouseProGearWeightOk() (*int32, bool)`

GetActualSpouseProGearWeightOk returns a tuple with the ActualSpouseProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSpouseProGearWeight

`func (o *MTOShipment) SetActualSpouseProGearWeight(v int32)`

SetActualSpouseProGearWeight sets ActualSpouseProGearWeight field to given value.

### HasActualSpouseProGearWeight

`func (o *MTOShipment) HasActualSpouseProGearWeight() bool`

HasActualSpouseProGearWeight returns a boolean if a field has been set.

### SetActualSpouseProGearWeightNil

`func (o *MTOShipment) SetActualSpouseProGearWeightNil(b bool)`

 SetActualSpouseProGearWeightNil sets the value for ActualSpouseProGearWeight to be an explicit nil

### UnsetActualSpouseProGearWeight
`func (o *MTOShipment) UnsetActualSpouseProGearWeight()`

UnsetActualSpouseProGearWeight ensures that no value is present for ActualSpouseProGearWeight, not even an explicit nil
### GetAgents

`func (o *MTOShipment) GetAgents() []MTOAgent`

GetAgents returns the Agents field if non-nil, zero value otherwise.

### GetAgentsOk

`func (o *MTOShipment) GetAgentsOk() (*[]MTOAgent, bool)`

GetAgentsOk returns a tuple with the Agents field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgents

`func (o *MTOShipment) SetAgents(v []MTOAgent)`

SetAgents sets Agents field to given value.

### HasAgents

`func (o *MTOShipment) HasAgents() bool`

HasAgents returns a boolean if a field has been set.

### GetSitExtensions

`func (o *MTOShipment) GetSitExtensions() []SITExtension`

GetSitExtensions returns the SitExtensions field if non-nil, zero value otherwise.

### GetSitExtensionsOk

`func (o *MTOShipment) GetSitExtensionsOk() (*[]SITExtension, bool)`

GetSitExtensionsOk returns a tuple with the SitExtensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitExtensions

`func (o *MTOShipment) SetSitExtensions(v []SITExtension)`

SetSitExtensions sets SitExtensions field to given value.

### HasSitExtensions

`func (o *MTOShipment) HasSitExtensions() bool`

HasSitExtensions returns a boolean if a field has been set.

### GetReweigh

`func (o *MTOShipment) GetReweigh() Reweigh`

GetReweigh returns the Reweigh field if non-nil, zero value otherwise.

### GetReweighOk

`func (o *MTOShipment) GetReweighOk() (*Reweigh, bool)`

GetReweighOk returns a tuple with the Reweigh field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReweigh

`func (o *MTOShipment) SetReweigh(v Reweigh)`

SetReweigh sets Reweigh field to given value.

### HasReweigh

`func (o *MTOShipment) HasReweigh() bool`

HasReweigh returns a boolean if a field has been set.

### GetPickupAddress

`func (o *MTOShipment) GetPickupAddress() Address`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *MTOShipment) GetPickupAddressOk() (*Address, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *MTOShipment) SetPickupAddress(v Address)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *MTOShipment) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *MTOShipment) GetDestinationAddress() Address`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *MTOShipment) GetDestinationAddressOk() (*Address, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *MTOShipment) SetDestinationAddress(v Address)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *MTOShipment) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetDestinationType

`func (o *MTOShipment) GetDestinationType() DestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *MTOShipment) GetDestinationTypeOk() (*DestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *MTOShipment) SetDestinationType(v DestinationType)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *MTOShipment) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### SetDestinationTypeNil

`func (o *MTOShipment) SetDestinationTypeNil(b bool)`

 SetDestinationTypeNil sets the value for DestinationType to be an explicit nil

### UnsetDestinationType
`func (o *MTOShipment) UnsetDestinationType()`

UnsetDestinationType ensures that no value is present for DestinationType, not even an explicit nil
### GetSecondaryPickupAddress

`func (o *MTOShipment) GetSecondaryPickupAddress() Address`

GetSecondaryPickupAddress returns the SecondaryPickupAddress field if non-nil, zero value otherwise.

### GetSecondaryPickupAddressOk

`func (o *MTOShipment) GetSecondaryPickupAddressOk() (*Address, bool)`

GetSecondaryPickupAddressOk returns a tuple with the SecondaryPickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryPickupAddress

`func (o *MTOShipment) SetSecondaryPickupAddress(v Address)`

SetSecondaryPickupAddress sets SecondaryPickupAddress field to given value.

### HasSecondaryPickupAddress

`func (o *MTOShipment) HasSecondaryPickupAddress() bool`

HasSecondaryPickupAddress returns a boolean if a field has been set.

### GetSecondaryDeliveryAddress

`func (o *MTOShipment) GetSecondaryDeliveryAddress() Address`

GetSecondaryDeliveryAddress returns the SecondaryDeliveryAddress field if non-nil, zero value otherwise.

### GetSecondaryDeliveryAddressOk

`func (o *MTOShipment) GetSecondaryDeliveryAddressOk() (*Address, bool)`

GetSecondaryDeliveryAddressOk returns a tuple with the SecondaryDeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDeliveryAddress

`func (o *MTOShipment) SetSecondaryDeliveryAddress(v Address)`

SetSecondaryDeliveryAddress sets SecondaryDeliveryAddress field to given value.

### HasSecondaryDeliveryAddress

`func (o *MTOShipment) HasSecondaryDeliveryAddress() bool`

HasSecondaryDeliveryAddress returns a boolean if a field has been set.

### GetStorageFacility

`func (o *MTOShipment) GetStorageFacility() UpdateMTOShipmentStorageFacility`

GetStorageFacility returns the StorageFacility field if non-nil, zero value otherwise.

### GetStorageFacilityOk

`func (o *MTOShipment) GetStorageFacilityOk() (*UpdateMTOShipmentStorageFacility, bool)`

GetStorageFacilityOk returns a tuple with the StorageFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFacility

`func (o *MTOShipment) SetStorageFacility(v UpdateMTOShipmentStorageFacility)`

SetStorageFacility sets StorageFacility field to given value.

### HasStorageFacility

`func (o *MTOShipment) HasStorageFacility() bool`

HasStorageFacility returns a boolean if a field has been set.

### GetShipmentType

`func (o *MTOShipment) GetShipmentType() MTOShipmentType`

GetShipmentType returns the ShipmentType field if non-nil, zero value otherwise.

### GetShipmentTypeOk

`func (o *MTOShipment) GetShipmentTypeOk() (*MTOShipmentType, bool)`

GetShipmentTypeOk returns a tuple with the ShipmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentType

`func (o *MTOShipment) SetShipmentType(v MTOShipmentType)`

SetShipmentType sets ShipmentType field to given value.

### HasShipmentType

`func (o *MTOShipment) HasShipmentType() bool`

HasShipmentType returns a boolean if a field has been set.

### GetDiversion

`func (o *MTOShipment) GetDiversion() bool`

GetDiversion returns the Diversion field if non-nil, zero value otherwise.

### GetDiversionOk

`func (o *MTOShipment) GetDiversionOk() (*bool, bool)`

GetDiversionOk returns a tuple with the Diversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiversion

`func (o *MTOShipment) SetDiversion(v bool)`

SetDiversion sets Diversion field to given value.

### HasDiversion

`func (o *MTOShipment) HasDiversion() bool`

HasDiversion returns a boolean if a field has been set.

### GetDiversionReason

`func (o *MTOShipment) GetDiversionReason() string`

GetDiversionReason returns the DiversionReason field if non-nil, zero value otherwise.

### GetDiversionReasonOk

`func (o *MTOShipment) GetDiversionReasonOk() (*string, bool)`

GetDiversionReasonOk returns a tuple with the DiversionReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiversionReason

`func (o *MTOShipment) SetDiversionReason(v string)`

SetDiversionReason sets DiversionReason field to given value.

### HasDiversionReason

`func (o *MTOShipment) HasDiversionReason() bool`

HasDiversionReason returns a boolean if a field has been set.

### SetDiversionReasonNil

`func (o *MTOShipment) SetDiversionReasonNil(b bool)`

 SetDiversionReasonNil sets the value for DiversionReason to be an explicit nil

### UnsetDiversionReason
`func (o *MTOShipment) UnsetDiversionReason()`

UnsetDiversionReason ensures that no value is present for DiversionReason, not even an explicit nil
### GetStatus

`func (o *MTOShipment) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MTOShipment) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MTOShipment) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MTOShipment) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetPpmShipment

`func (o *MTOShipment) GetPpmShipment() PPMShipment`

GetPpmShipment returns the PpmShipment field if non-nil, zero value otherwise.

### GetPpmShipmentOk

`func (o *MTOShipment) GetPpmShipmentOk() (*PPMShipment, bool)`

GetPpmShipmentOk returns a tuple with the PpmShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpmShipment

`func (o *MTOShipment) SetPpmShipment(v PPMShipment)`

SetPpmShipment sets PpmShipment field to given value.

### HasPpmShipment

`func (o *MTOShipment) HasPpmShipment() bool`

HasPpmShipment returns a boolean if a field has been set.

### SetPpmShipmentNil

`func (o *MTOShipment) SetPpmShipmentNil(b bool)`

 SetPpmShipmentNil sets the value for PpmShipment to be an explicit nil

### UnsetPpmShipment
`func (o *MTOShipment) UnsetPpmShipment()`

UnsetPpmShipment ensures that no value is present for PpmShipment, not even an explicit nil
### GetDeliveryAddressUpdate

`func (o *MTOShipment) GetDeliveryAddressUpdate() ShipmentAddressUpdate`

GetDeliveryAddressUpdate returns the DeliveryAddressUpdate field if non-nil, zero value otherwise.

### GetDeliveryAddressUpdateOk

`func (o *MTOShipment) GetDeliveryAddressUpdateOk() (*ShipmentAddressUpdate, bool)`

GetDeliveryAddressUpdateOk returns a tuple with the DeliveryAddressUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryAddressUpdate

`func (o *MTOShipment) SetDeliveryAddressUpdate(v ShipmentAddressUpdate)`

SetDeliveryAddressUpdate sets DeliveryAddressUpdate field to given value.

### HasDeliveryAddressUpdate

`func (o *MTOShipment) HasDeliveryAddressUpdate() bool`

HasDeliveryAddressUpdate returns a boolean if a field has been set.

### GetETag

`func (o *MTOShipment) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MTOShipment) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MTOShipment) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MTOShipment) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MTOShipment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MTOShipment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MTOShipment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MTOShipment) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MTOShipment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MTOShipment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MTOShipment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MTOShipment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPointOfContact

`func (o *MTOShipment) GetPointOfContact() string`

GetPointOfContact returns the PointOfContact field if non-nil, zero value otherwise.

### GetPointOfContactOk

`func (o *MTOShipment) GetPointOfContactOk() (*string, bool)`

GetPointOfContactOk returns a tuple with the PointOfContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointOfContact

`func (o *MTOShipment) SetPointOfContact(v string)`

SetPointOfContact sets PointOfContact field to given value.

### HasPointOfContact

`func (o *MTOShipment) HasPointOfContact() bool`

HasPointOfContact returns a boolean if a field has been set.

### GetOriginSitAuthEndDate

`func (o *MTOShipment) GetOriginSitAuthEndDate() string`

GetOriginSitAuthEndDate returns the OriginSitAuthEndDate field if non-nil, zero value otherwise.

### GetOriginSitAuthEndDateOk

`func (o *MTOShipment) GetOriginSitAuthEndDateOk() (*string, bool)`

GetOriginSitAuthEndDateOk returns a tuple with the OriginSitAuthEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginSitAuthEndDate

`func (o *MTOShipment) SetOriginSitAuthEndDate(v string)`

SetOriginSitAuthEndDate sets OriginSitAuthEndDate field to given value.

### HasOriginSitAuthEndDate

`func (o *MTOShipment) HasOriginSitAuthEndDate() bool`

HasOriginSitAuthEndDate returns a boolean if a field has been set.

### SetOriginSitAuthEndDateNil

`func (o *MTOShipment) SetOriginSitAuthEndDateNil(b bool)`

 SetOriginSitAuthEndDateNil sets the value for OriginSitAuthEndDate to be an explicit nil

### UnsetOriginSitAuthEndDate
`func (o *MTOShipment) UnsetOriginSitAuthEndDate()`

UnsetOriginSitAuthEndDate ensures that no value is present for OriginSitAuthEndDate, not even an explicit nil
### GetDestinationSitAuthEndDate

`func (o *MTOShipment) GetDestinationSitAuthEndDate() string`

GetDestinationSitAuthEndDate returns the DestinationSitAuthEndDate field if non-nil, zero value otherwise.

### GetDestinationSitAuthEndDateOk

`func (o *MTOShipment) GetDestinationSitAuthEndDateOk() (*string, bool)`

GetDestinationSitAuthEndDateOk returns a tuple with the DestinationSitAuthEndDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationSitAuthEndDate

`func (o *MTOShipment) SetDestinationSitAuthEndDate(v string)`

SetDestinationSitAuthEndDate sets DestinationSitAuthEndDate field to given value.

### HasDestinationSitAuthEndDate

`func (o *MTOShipment) HasDestinationSitAuthEndDate() bool`

HasDestinationSitAuthEndDate returns a boolean if a field has been set.

### SetDestinationSitAuthEndDateNil

`func (o *MTOShipment) SetDestinationSitAuthEndDateNil(b bool)`

 SetDestinationSitAuthEndDateNil sets the value for DestinationSitAuthEndDate to be an explicit nil

### UnsetDestinationSitAuthEndDate
`func (o *MTOShipment) UnsetDestinationSitAuthEndDate()`

UnsetDestinationSitAuthEndDate ensures that no value is present for DestinationSitAuthEndDate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


