# UpdateMTOShipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualProGearWeight** | Pointer to **NullableInt32** | The actual weight of any pro gear shipped during a move. | [optional] 
**ActualSpouseProGearWeight** | Pointer to **NullableInt32** | The actual weight of any pro gear shipped during a move. | [optional] 
**ScheduledPickupDate** | Pointer to **NullableString** | The date the Prime contractor scheduled to pick up this shipment after consultation with the customer. | [optional] 
**ActualPickupDate** | Pointer to **NullableString** | The date when the Prime contractor actually picked up the shipment. Updated after-the-fact. | [optional] 
**FirstAvailableDeliveryDate** | Pointer to **NullableString** | The date the Prime provides to the customer as the first possible delivery date so that they can plan their travel accordingly.  | [optional] 
**ScheduledDeliveryDate** | Pointer to **NullableString** | The date the Prime contractor scheduled to deliver this shipment after consultation with the customer. | [optional] 
**ActualDeliveryDate** | Pointer to **NullableString** | The date when the Prime contractor actually delivered the shipment. Updated after-the-fact. | [optional] 
**PrimeEstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of this shipment, determined by the movers during the pre-move survey. This value **can only be updated once.** If there was an issue with estimating the weight and a mistake was made, the Prime contracter will need to contact the TOO to change it.  | [optional] 
**PrimeActualWeight** | Pointer to **NullableInt32** | The actual weight of the shipment, provided after the Prime packs, picks up, and weighs a customer&#39;s shipment. | [optional] 
**NtsRecordedWeight** | Pointer to **NullableInt32** | The previously recorded weight for the NTS Shipment. Used for NTS Release to know what the previous primeActualWeight or billable weight was. | [optional] 
**PickupAddress** | Pointer to [**Address**](Address.md) | The address where the movers should pick up this shipment, entered by the customer during onboarding when they enter shipment details.  | [optional] 
**DestinationAddress** | Pointer to [**Address**](Address.md) | Where the movers should deliver this shipment. Often provided by the customer when they enter shipment details during onboarding, if they know their new address already.  May be blank when entered by the customer, required when entered by the Prime. May not represent the true final destination due to the shipment being diverted or placed in SIT.  | [optional] 
**DestinationType** | Pointer to [**NullableDestinationType**](DestinationType.md) |  | [optional] 
**SecondaryPickupAddress** | Pointer to [**Address**](Address.md) | A second pickup address for this shipment, if the customer entered one. An optional field. | [optional] 
**SecondaryDeliveryAddress** | Pointer to [**Address**](Address.md) | A second delivery address for this shipment, if the customer entered one. An optional field. | [optional] 
**StorageFacility** | Pointer to [**UpdateMTOShipmentStorageFacility**](UpdateMTOShipmentStorageFacility.md) |  | [optional] 
**ShipmentType** | Pointer to [**MTOShipmentType**](MTOShipmentType.md) |  | [optional] 
**Diversion** | Pointer to **bool** | This value indicates whether or not this shipment is part of a diversion. If yes, the shipment can be either the starting or ending segment of the diversion.  | [optional] 
**PointOfContact** | Pointer to **string** | Email or ID of the person who will be contacted in the event of questions or concerns about this update. May be the person performing the update, or someone else working with the Prime contractor.  | [optional] 
**CounselorRemarks** | Pointer to **NullableString** |  | [optional] 
**PpmShipment** | Pointer to [**UpdatePPMShipment**](UpdatePPMShipment.md) |  | [optional] 

## Methods

### NewUpdateMTOShipment

`func NewUpdateMTOShipment() *UpdateMTOShipment`

NewUpdateMTOShipment instantiates a new UpdateMTOShipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMTOShipmentWithDefaults

`func NewUpdateMTOShipmentWithDefaults() *UpdateMTOShipment`

NewUpdateMTOShipmentWithDefaults instantiates a new UpdateMTOShipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualProGearWeight

`func (o *UpdateMTOShipment) GetActualProGearWeight() int32`

GetActualProGearWeight returns the ActualProGearWeight field if non-nil, zero value otherwise.

### GetActualProGearWeightOk

`func (o *UpdateMTOShipment) GetActualProGearWeightOk() (*int32, bool)`

GetActualProGearWeightOk returns a tuple with the ActualProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualProGearWeight

`func (o *UpdateMTOShipment) SetActualProGearWeight(v int32)`

SetActualProGearWeight sets ActualProGearWeight field to given value.

### HasActualProGearWeight

`func (o *UpdateMTOShipment) HasActualProGearWeight() bool`

HasActualProGearWeight returns a boolean if a field has been set.

### SetActualProGearWeightNil

`func (o *UpdateMTOShipment) SetActualProGearWeightNil(b bool)`

 SetActualProGearWeightNil sets the value for ActualProGearWeight to be an explicit nil

### UnsetActualProGearWeight
`func (o *UpdateMTOShipment) UnsetActualProGearWeight()`

UnsetActualProGearWeight ensures that no value is present for ActualProGearWeight, not even an explicit nil
### GetActualSpouseProGearWeight

`func (o *UpdateMTOShipment) GetActualSpouseProGearWeight() int32`

GetActualSpouseProGearWeight returns the ActualSpouseProGearWeight field if non-nil, zero value otherwise.

### GetActualSpouseProGearWeightOk

`func (o *UpdateMTOShipment) GetActualSpouseProGearWeightOk() (*int32, bool)`

GetActualSpouseProGearWeightOk returns a tuple with the ActualSpouseProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSpouseProGearWeight

`func (o *UpdateMTOShipment) SetActualSpouseProGearWeight(v int32)`

SetActualSpouseProGearWeight sets ActualSpouseProGearWeight field to given value.

### HasActualSpouseProGearWeight

`func (o *UpdateMTOShipment) HasActualSpouseProGearWeight() bool`

HasActualSpouseProGearWeight returns a boolean if a field has been set.

### SetActualSpouseProGearWeightNil

`func (o *UpdateMTOShipment) SetActualSpouseProGearWeightNil(b bool)`

 SetActualSpouseProGearWeightNil sets the value for ActualSpouseProGearWeight to be an explicit nil

### UnsetActualSpouseProGearWeight
`func (o *UpdateMTOShipment) UnsetActualSpouseProGearWeight()`

UnsetActualSpouseProGearWeight ensures that no value is present for ActualSpouseProGearWeight, not even an explicit nil
### GetScheduledPickupDate

`func (o *UpdateMTOShipment) GetScheduledPickupDate() string`

GetScheduledPickupDate returns the ScheduledPickupDate field if non-nil, zero value otherwise.

### GetScheduledPickupDateOk

`func (o *UpdateMTOShipment) GetScheduledPickupDateOk() (*string, bool)`

GetScheduledPickupDateOk returns a tuple with the ScheduledPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPickupDate

`func (o *UpdateMTOShipment) SetScheduledPickupDate(v string)`

SetScheduledPickupDate sets ScheduledPickupDate field to given value.

### HasScheduledPickupDate

`func (o *UpdateMTOShipment) HasScheduledPickupDate() bool`

HasScheduledPickupDate returns a boolean if a field has been set.

### SetScheduledPickupDateNil

`func (o *UpdateMTOShipment) SetScheduledPickupDateNil(b bool)`

 SetScheduledPickupDateNil sets the value for ScheduledPickupDate to be an explicit nil

### UnsetScheduledPickupDate
`func (o *UpdateMTOShipment) UnsetScheduledPickupDate()`

UnsetScheduledPickupDate ensures that no value is present for ScheduledPickupDate, not even an explicit nil
### GetActualPickupDate

`func (o *UpdateMTOShipment) GetActualPickupDate() string`

GetActualPickupDate returns the ActualPickupDate field if non-nil, zero value otherwise.

### GetActualPickupDateOk

`func (o *UpdateMTOShipment) GetActualPickupDateOk() (*string, bool)`

GetActualPickupDateOk returns a tuple with the ActualPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPickupDate

`func (o *UpdateMTOShipment) SetActualPickupDate(v string)`

SetActualPickupDate sets ActualPickupDate field to given value.

### HasActualPickupDate

`func (o *UpdateMTOShipment) HasActualPickupDate() bool`

HasActualPickupDate returns a boolean if a field has been set.

### SetActualPickupDateNil

`func (o *UpdateMTOShipment) SetActualPickupDateNil(b bool)`

 SetActualPickupDateNil sets the value for ActualPickupDate to be an explicit nil

### UnsetActualPickupDate
`func (o *UpdateMTOShipment) UnsetActualPickupDate()`

UnsetActualPickupDate ensures that no value is present for ActualPickupDate, not even an explicit nil
### GetFirstAvailableDeliveryDate

`func (o *UpdateMTOShipment) GetFirstAvailableDeliveryDate() string`

GetFirstAvailableDeliveryDate returns the FirstAvailableDeliveryDate field if non-nil, zero value otherwise.

### GetFirstAvailableDeliveryDateOk

`func (o *UpdateMTOShipment) GetFirstAvailableDeliveryDateOk() (*string, bool)`

GetFirstAvailableDeliveryDateOk returns a tuple with the FirstAvailableDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAvailableDeliveryDate

`func (o *UpdateMTOShipment) SetFirstAvailableDeliveryDate(v string)`

SetFirstAvailableDeliveryDate sets FirstAvailableDeliveryDate field to given value.

### HasFirstAvailableDeliveryDate

`func (o *UpdateMTOShipment) HasFirstAvailableDeliveryDate() bool`

HasFirstAvailableDeliveryDate returns a boolean if a field has been set.

### SetFirstAvailableDeliveryDateNil

`func (o *UpdateMTOShipment) SetFirstAvailableDeliveryDateNil(b bool)`

 SetFirstAvailableDeliveryDateNil sets the value for FirstAvailableDeliveryDate to be an explicit nil

### UnsetFirstAvailableDeliveryDate
`func (o *UpdateMTOShipment) UnsetFirstAvailableDeliveryDate()`

UnsetFirstAvailableDeliveryDate ensures that no value is present for FirstAvailableDeliveryDate, not even an explicit nil
### GetScheduledDeliveryDate

`func (o *UpdateMTOShipment) GetScheduledDeliveryDate() string`

GetScheduledDeliveryDate returns the ScheduledDeliveryDate field if non-nil, zero value otherwise.

### GetScheduledDeliveryDateOk

`func (o *UpdateMTOShipment) GetScheduledDeliveryDateOk() (*string, bool)`

GetScheduledDeliveryDateOk returns a tuple with the ScheduledDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeliveryDate

`func (o *UpdateMTOShipment) SetScheduledDeliveryDate(v string)`

SetScheduledDeliveryDate sets ScheduledDeliveryDate field to given value.

### HasScheduledDeliveryDate

`func (o *UpdateMTOShipment) HasScheduledDeliveryDate() bool`

HasScheduledDeliveryDate returns a boolean if a field has been set.

### SetScheduledDeliveryDateNil

`func (o *UpdateMTOShipment) SetScheduledDeliveryDateNil(b bool)`

 SetScheduledDeliveryDateNil sets the value for ScheduledDeliveryDate to be an explicit nil

### UnsetScheduledDeliveryDate
`func (o *UpdateMTOShipment) UnsetScheduledDeliveryDate()`

UnsetScheduledDeliveryDate ensures that no value is present for ScheduledDeliveryDate, not even an explicit nil
### GetActualDeliveryDate

`func (o *UpdateMTOShipment) GetActualDeliveryDate() string`

GetActualDeliveryDate returns the ActualDeliveryDate field if non-nil, zero value otherwise.

### GetActualDeliveryDateOk

`func (o *UpdateMTOShipment) GetActualDeliveryDateOk() (*string, bool)`

GetActualDeliveryDateOk returns a tuple with the ActualDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDeliveryDate

`func (o *UpdateMTOShipment) SetActualDeliveryDate(v string)`

SetActualDeliveryDate sets ActualDeliveryDate field to given value.

### HasActualDeliveryDate

`func (o *UpdateMTOShipment) HasActualDeliveryDate() bool`

HasActualDeliveryDate returns a boolean if a field has been set.

### SetActualDeliveryDateNil

`func (o *UpdateMTOShipment) SetActualDeliveryDateNil(b bool)`

 SetActualDeliveryDateNil sets the value for ActualDeliveryDate to be an explicit nil

### UnsetActualDeliveryDate
`func (o *UpdateMTOShipment) UnsetActualDeliveryDate()`

UnsetActualDeliveryDate ensures that no value is present for ActualDeliveryDate, not even an explicit nil
### GetPrimeEstimatedWeight

`func (o *UpdateMTOShipment) GetPrimeEstimatedWeight() int32`

GetPrimeEstimatedWeight returns the PrimeEstimatedWeight field if non-nil, zero value otherwise.

### GetPrimeEstimatedWeightOk

`func (o *UpdateMTOShipment) GetPrimeEstimatedWeightOk() (*int32, bool)`

GetPrimeEstimatedWeightOk returns a tuple with the PrimeEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeEstimatedWeight

`func (o *UpdateMTOShipment) SetPrimeEstimatedWeight(v int32)`

SetPrimeEstimatedWeight sets PrimeEstimatedWeight field to given value.

### HasPrimeEstimatedWeight

`func (o *UpdateMTOShipment) HasPrimeEstimatedWeight() bool`

HasPrimeEstimatedWeight returns a boolean if a field has been set.

### SetPrimeEstimatedWeightNil

`func (o *UpdateMTOShipment) SetPrimeEstimatedWeightNil(b bool)`

 SetPrimeEstimatedWeightNil sets the value for PrimeEstimatedWeight to be an explicit nil

### UnsetPrimeEstimatedWeight
`func (o *UpdateMTOShipment) UnsetPrimeEstimatedWeight()`

UnsetPrimeEstimatedWeight ensures that no value is present for PrimeEstimatedWeight, not even an explicit nil
### GetPrimeActualWeight

`func (o *UpdateMTOShipment) GetPrimeActualWeight() int32`

GetPrimeActualWeight returns the PrimeActualWeight field if non-nil, zero value otherwise.

### GetPrimeActualWeightOk

`func (o *UpdateMTOShipment) GetPrimeActualWeightOk() (*int32, bool)`

GetPrimeActualWeightOk returns a tuple with the PrimeActualWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeActualWeight

`func (o *UpdateMTOShipment) SetPrimeActualWeight(v int32)`

SetPrimeActualWeight sets PrimeActualWeight field to given value.

### HasPrimeActualWeight

`func (o *UpdateMTOShipment) HasPrimeActualWeight() bool`

HasPrimeActualWeight returns a boolean if a field has been set.

### SetPrimeActualWeightNil

`func (o *UpdateMTOShipment) SetPrimeActualWeightNil(b bool)`

 SetPrimeActualWeightNil sets the value for PrimeActualWeight to be an explicit nil

### UnsetPrimeActualWeight
`func (o *UpdateMTOShipment) UnsetPrimeActualWeight()`

UnsetPrimeActualWeight ensures that no value is present for PrimeActualWeight, not even an explicit nil
### GetNtsRecordedWeight

`func (o *UpdateMTOShipment) GetNtsRecordedWeight() int32`

GetNtsRecordedWeight returns the NtsRecordedWeight field if non-nil, zero value otherwise.

### GetNtsRecordedWeightOk

`func (o *UpdateMTOShipment) GetNtsRecordedWeightOk() (*int32, bool)`

GetNtsRecordedWeightOk returns a tuple with the NtsRecordedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtsRecordedWeight

`func (o *UpdateMTOShipment) SetNtsRecordedWeight(v int32)`

SetNtsRecordedWeight sets NtsRecordedWeight field to given value.

### HasNtsRecordedWeight

`func (o *UpdateMTOShipment) HasNtsRecordedWeight() bool`

HasNtsRecordedWeight returns a boolean if a field has been set.

### SetNtsRecordedWeightNil

`func (o *UpdateMTOShipment) SetNtsRecordedWeightNil(b bool)`

 SetNtsRecordedWeightNil sets the value for NtsRecordedWeight to be an explicit nil

### UnsetNtsRecordedWeight
`func (o *UpdateMTOShipment) UnsetNtsRecordedWeight()`

UnsetNtsRecordedWeight ensures that no value is present for NtsRecordedWeight, not even an explicit nil
### GetPickupAddress

`func (o *UpdateMTOShipment) GetPickupAddress() Address`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *UpdateMTOShipment) GetPickupAddressOk() (*Address, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *UpdateMTOShipment) SetPickupAddress(v Address)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *UpdateMTOShipment) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *UpdateMTOShipment) GetDestinationAddress() Address`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *UpdateMTOShipment) GetDestinationAddressOk() (*Address, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *UpdateMTOShipment) SetDestinationAddress(v Address)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *UpdateMTOShipment) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetDestinationType

`func (o *UpdateMTOShipment) GetDestinationType() DestinationType`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *UpdateMTOShipment) GetDestinationTypeOk() (*DestinationType, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *UpdateMTOShipment) SetDestinationType(v DestinationType)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *UpdateMTOShipment) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### SetDestinationTypeNil

`func (o *UpdateMTOShipment) SetDestinationTypeNil(b bool)`

 SetDestinationTypeNil sets the value for DestinationType to be an explicit nil

### UnsetDestinationType
`func (o *UpdateMTOShipment) UnsetDestinationType()`

UnsetDestinationType ensures that no value is present for DestinationType, not even an explicit nil
### GetSecondaryPickupAddress

`func (o *UpdateMTOShipment) GetSecondaryPickupAddress() Address`

GetSecondaryPickupAddress returns the SecondaryPickupAddress field if non-nil, zero value otherwise.

### GetSecondaryPickupAddressOk

`func (o *UpdateMTOShipment) GetSecondaryPickupAddressOk() (*Address, bool)`

GetSecondaryPickupAddressOk returns a tuple with the SecondaryPickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryPickupAddress

`func (o *UpdateMTOShipment) SetSecondaryPickupAddress(v Address)`

SetSecondaryPickupAddress sets SecondaryPickupAddress field to given value.

### HasSecondaryPickupAddress

`func (o *UpdateMTOShipment) HasSecondaryPickupAddress() bool`

HasSecondaryPickupAddress returns a boolean if a field has been set.

### GetSecondaryDeliveryAddress

`func (o *UpdateMTOShipment) GetSecondaryDeliveryAddress() Address`

GetSecondaryDeliveryAddress returns the SecondaryDeliveryAddress field if non-nil, zero value otherwise.

### GetSecondaryDeliveryAddressOk

`func (o *UpdateMTOShipment) GetSecondaryDeliveryAddressOk() (*Address, bool)`

GetSecondaryDeliveryAddressOk returns a tuple with the SecondaryDeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDeliveryAddress

`func (o *UpdateMTOShipment) SetSecondaryDeliveryAddress(v Address)`

SetSecondaryDeliveryAddress sets SecondaryDeliveryAddress field to given value.

### HasSecondaryDeliveryAddress

`func (o *UpdateMTOShipment) HasSecondaryDeliveryAddress() bool`

HasSecondaryDeliveryAddress returns a boolean if a field has been set.

### GetStorageFacility

`func (o *UpdateMTOShipment) GetStorageFacility() UpdateMTOShipmentStorageFacility`

GetStorageFacility returns the StorageFacility field if non-nil, zero value otherwise.

### GetStorageFacilityOk

`func (o *UpdateMTOShipment) GetStorageFacilityOk() (*UpdateMTOShipmentStorageFacility, bool)`

GetStorageFacilityOk returns a tuple with the StorageFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFacility

`func (o *UpdateMTOShipment) SetStorageFacility(v UpdateMTOShipmentStorageFacility)`

SetStorageFacility sets StorageFacility field to given value.

### HasStorageFacility

`func (o *UpdateMTOShipment) HasStorageFacility() bool`

HasStorageFacility returns a boolean if a field has been set.

### GetShipmentType

`func (o *UpdateMTOShipment) GetShipmentType() MTOShipmentType`

GetShipmentType returns the ShipmentType field if non-nil, zero value otherwise.

### GetShipmentTypeOk

`func (o *UpdateMTOShipment) GetShipmentTypeOk() (*MTOShipmentType, bool)`

GetShipmentTypeOk returns a tuple with the ShipmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentType

`func (o *UpdateMTOShipment) SetShipmentType(v MTOShipmentType)`

SetShipmentType sets ShipmentType field to given value.

### HasShipmentType

`func (o *UpdateMTOShipment) HasShipmentType() bool`

HasShipmentType returns a boolean if a field has been set.

### GetDiversion

`func (o *UpdateMTOShipment) GetDiversion() bool`

GetDiversion returns the Diversion field if non-nil, zero value otherwise.

### GetDiversionOk

`func (o *UpdateMTOShipment) GetDiversionOk() (*bool, bool)`

GetDiversionOk returns a tuple with the Diversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiversion

`func (o *UpdateMTOShipment) SetDiversion(v bool)`

SetDiversion sets Diversion field to given value.

### HasDiversion

`func (o *UpdateMTOShipment) HasDiversion() bool`

HasDiversion returns a boolean if a field has been set.

### GetPointOfContact

`func (o *UpdateMTOShipment) GetPointOfContact() string`

GetPointOfContact returns the PointOfContact field if non-nil, zero value otherwise.

### GetPointOfContactOk

`func (o *UpdateMTOShipment) GetPointOfContactOk() (*string, bool)`

GetPointOfContactOk returns a tuple with the PointOfContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointOfContact

`func (o *UpdateMTOShipment) SetPointOfContact(v string)`

SetPointOfContact sets PointOfContact field to given value.

### HasPointOfContact

`func (o *UpdateMTOShipment) HasPointOfContact() bool`

HasPointOfContact returns a boolean if a field has been set.

### GetCounselorRemarks

`func (o *UpdateMTOShipment) GetCounselorRemarks() string`

GetCounselorRemarks returns the CounselorRemarks field if non-nil, zero value otherwise.

### GetCounselorRemarksOk

`func (o *UpdateMTOShipment) GetCounselorRemarksOk() (*string, bool)`

GetCounselorRemarksOk returns a tuple with the CounselorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounselorRemarks

`func (o *UpdateMTOShipment) SetCounselorRemarks(v string)`

SetCounselorRemarks sets CounselorRemarks field to given value.

### HasCounselorRemarks

`func (o *UpdateMTOShipment) HasCounselorRemarks() bool`

HasCounselorRemarks returns a boolean if a field has been set.

### SetCounselorRemarksNil

`func (o *UpdateMTOShipment) SetCounselorRemarksNil(b bool)`

 SetCounselorRemarksNil sets the value for CounselorRemarks to be an explicit nil

### UnsetCounselorRemarks
`func (o *UpdateMTOShipment) UnsetCounselorRemarks()`

UnsetCounselorRemarks ensures that no value is present for CounselorRemarks, not even an explicit nil
### GetPpmShipment

`func (o *UpdateMTOShipment) GetPpmShipment() UpdatePPMShipment`

GetPpmShipment returns the PpmShipment field if non-nil, zero value otherwise.

### GetPpmShipmentOk

`func (o *UpdateMTOShipment) GetPpmShipmentOk() (*UpdatePPMShipment, bool)`

GetPpmShipmentOk returns a tuple with the PpmShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpmShipment

`func (o *UpdateMTOShipment) SetPpmShipment(v UpdatePPMShipment)`

SetPpmShipment sets PpmShipment field to given value.

### HasPpmShipment

`func (o *UpdateMTOShipment) HasPpmShipment() bool`

HasPpmShipment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


