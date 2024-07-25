# UpdateMTOShipmentV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ScheduledPickupDate** | Pointer to **NullableString** | The date the Prime contractor scheduled to pick up this shipment after consultation with the customer. | [optional] 
**ActualPickupDate** | Pointer to **NullableString** | The date when the Prime contractor actually picked up the shipment. Updated after-the-fact. | [optional] 
**FirstAvailableDeliveryDate** | Pointer to **NullableString** | The date the Prime provides to the customer as the first possible delivery date so that they can plan their travel accordingly.  | [optional] 
**ScheduledDeliveryDate** | Pointer to **NullableString** | The date the Prime contractor scheduled to deliver this shipment after consultation with the customer. | [optional] 
**ActualDeliveryDate** | Pointer to **NullableString** | The date when the Prime contractor actually delivered the shipment. Updated after-the-fact. | [optional] 
**PrimeEstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of this shipment, determined by the movers during the pre-move survey. This value **can only be updated once.** If there was an issue with estimating the weight and a mistake was made, the Prime contracter will need to contact the TOO to change it.  | [optional] 
**PrimeActualWeight** | Pointer to **NullableInt32** | The actual weight of the shipment, provided after the Prime packs, picks up, and weighs a customer&#39;s shipment. | [optional] 
**NtsRecordedWeight** | Pointer to **NullableInt32** | The previously recorded weight for the NTS Shipment. Used for NTS Release to know what the previous primeActualWeight or billable weight was. | [optional] 
**PickupAddress** | Pointer to [**AddressV2V2**](AddressV2.md) | The address where the movers should pick up this shipment, entered by the customer during onboarding when they enter shipment details.  | [optional] 
**DestinationAddress** | Pointer to [**AddressV2V2**](AddressV2.md) | Where the movers should deliver this shipment. Often provided by the customer when they enter shipment details during onboarding, if they know their new address already.  May be blank when entered by the customer, required when entered by the Prime. May not represent the true final destination due to the shipment being diverted or placed in SIT.  | [optional] 
**DestinationType** | Pointer to [**NullableDestinationTypeV2V2**](DestinationTypeV2.md) |  | [optional] 
**SecondaryPickupAddress** | Pointer to [**AddressV2V2**](AddressV2.md) | A second pickup address for this shipment, if the customer entered one. An optional field. | [optional] 
**SecondaryDeliveryAddress** | Pointer to [**AddressV2V2**](AddressV2.md) | A second delivery address for this shipment, if the customer entered one. An optional field. | [optional] 
**StorageFacility** | Pointer to [**UpdateMTOShipmentStorageFacilityV2V2**](UpdateMTOShipmentStorageFacilityV2.md) |  | [optional] 
**ShipmentType** | Pointer to [**MTOShipmentTypeV2V2**](MTOShipmentTypeV2.md) |  | [optional] 
**Diversion** | Pointer to **bool** | This value indicates whether or not this shipment is part of a diversion. If yes, the shipment can be either the starting or ending segment of the diversion.  | [optional] 
**PointOfContact** | Pointer to **string** | Email or ID of the person who will be contacted in the event of questions or concerns about this update. May be the person performing the update, or someone else working with the Prime contractor.  | [optional] 
**CounselorRemarks** | Pointer to **NullableString** |  | [optional] 
**ActualProGearWeight** | Pointer to **NullableInt32** |  | [optional] 
**ActualSpouseProGearWeight** | Pointer to **NullableInt32** |  | [optional] 
**PpmShipment** | Pointer to [**UpdatePPMShipmentV2V2**](UpdatePPMShipmentV2.md) |  | [optional] 

## Methods

### NewUpdateMTOShipmentV2

`func NewUpdateMTOShipmentV2() *UpdateMTOShipmentV2`

NewUpdateMTOShipmentV2 instantiates a new UpdateMTOShipmentV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMTOShipmentV2WithDefaults

`func NewUpdateMTOShipmentV2WithDefaults() *UpdateMTOShipmentV2`

NewUpdateMTOShipmentV2WithDefaults instantiates a new UpdateMTOShipmentV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScheduledPickupDate

`func (o *UpdateMTOShipmentV2) GetScheduledPickupDate() string`

GetScheduledPickupDate returns the ScheduledPickupDate field if non-nil, zero value otherwise.

### GetScheduledPickupDateOk

`func (o *UpdateMTOShipmentV2) GetScheduledPickupDateOk() (*string, bool)`

GetScheduledPickupDateOk returns a tuple with the ScheduledPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledPickupDate

`func (o *UpdateMTOShipmentV2) SetScheduledPickupDate(v string)`

SetScheduledPickupDate sets ScheduledPickupDate field to given value.

### HasScheduledPickupDate

`func (o *UpdateMTOShipmentV2) HasScheduledPickupDate() bool`

HasScheduledPickupDate returns a boolean if a field has been set.

### SetScheduledPickupDateNil

`func (o *UpdateMTOShipmentV2) SetScheduledPickupDateNil(b bool)`

 SetScheduledPickupDateNil sets the value for ScheduledPickupDate to be an explicit nil

### UnsetScheduledPickupDate
`func (o *UpdateMTOShipmentV2) UnsetScheduledPickupDate()`

UnsetScheduledPickupDate ensures that no value is present for ScheduledPickupDate, not even an explicit nil
### GetActualPickupDate

`func (o *UpdateMTOShipmentV2) GetActualPickupDate() string`

GetActualPickupDate returns the ActualPickupDate field if non-nil, zero value otherwise.

### GetActualPickupDateOk

`func (o *UpdateMTOShipmentV2) GetActualPickupDateOk() (*string, bool)`

GetActualPickupDateOk returns a tuple with the ActualPickupDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPickupDate

`func (o *UpdateMTOShipmentV2) SetActualPickupDate(v string)`

SetActualPickupDate sets ActualPickupDate field to given value.

### HasActualPickupDate

`func (o *UpdateMTOShipmentV2) HasActualPickupDate() bool`

HasActualPickupDate returns a boolean if a field has been set.

### SetActualPickupDateNil

`func (o *UpdateMTOShipmentV2) SetActualPickupDateNil(b bool)`

 SetActualPickupDateNil sets the value for ActualPickupDate to be an explicit nil

### UnsetActualPickupDate
`func (o *UpdateMTOShipmentV2) UnsetActualPickupDate()`

UnsetActualPickupDate ensures that no value is present for ActualPickupDate, not even an explicit nil
### GetFirstAvailableDeliveryDate

`func (o *UpdateMTOShipmentV2) GetFirstAvailableDeliveryDate() string`

GetFirstAvailableDeliveryDate returns the FirstAvailableDeliveryDate field if non-nil, zero value otherwise.

### GetFirstAvailableDeliveryDateOk

`func (o *UpdateMTOShipmentV2) GetFirstAvailableDeliveryDateOk() (*string, bool)`

GetFirstAvailableDeliveryDateOk returns a tuple with the FirstAvailableDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAvailableDeliveryDate

`func (o *UpdateMTOShipmentV2) SetFirstAvailableDeliveryDate(v string)`

SetFirstAvailableDeliveryDate sets FirstAvailableDeliveryDate field to given value.

### HasFirstAvailableDeliveryDate

`func (o *UpdateMTOShipmentV2) HasFirstAvailableDeliveryDate() bool`

HasFirstAvailableDeliveryDate returns a boolean if a field has been set.

### SetFirstAvailableDeliveryDateNil

`func (o *UpdateMTOShipmentV2) SetFirstAvailableDeliveryDateNil(b bool)`

 SetFirstAvailableDeliveryDateNil sets the value for FirstAvailableDeliveryDate to be an explicit nil

### UnsetFirstAvailableDeliveryDate
`func (o *UpdateMTOShipmentV2) UnsetFirstAvailableDeliveryDate()`

UnsetFirstAvailableDeliveryDate ensures that no value is present for FirstAvailableDeliveryDate, not even an explicit nil
### GetScheduledDeliveryDate

`func (o *UpdateMTOShipmentV2) GetScheduledDeliveryDate() string`

GetScheduledDeliveryDate returns the ScheduledDeliveryDate field if non-nil, zero value otherwise.

### GetScheduledDeliveryDateOk

`func (o *UpdateMTOShipmentV2) GetScheduledDeliveryDateOk() (*string, bool)`

GetScheduledDeliveryDateOk returns a tuple with the ScheduledDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScheduledDeliveryDate

`func (o *UpdateMTOShipmentV2) SetScheduledDeliveryDate(v string)`

SetScheduledDeliveryDate sets ScheduledDeliveryDate field to given value.

### HasScheduledDeliveryDate

`func (o *UpdateMTOShipmentV2) HasScheduledDeliveryDate() bool`

HasScheduledDeliveryDate returns a boolean if a field has been set.

### SetScheduledDeliveryDateNil

`func (o *UpdateMTOShipmentV2) SetScheduledDeliveryDateNil(b bool)`

 SetScheduledDeliveryDateNil sets the value for ScheduledDeliveryDate to be an explicit nil

### UnsetScheduledDeliveryDate
`func (o *UpdateMTOShipmentV2) UnsetScheduledDeliveryDate()`

UnsetScheduledDeliveryDate ensures that no value is present for ScheduledDeliveryDate, not even an explicit nil
### GetActualDeliveryDate

`func (o *UpdateMTOShipmentV2) GetActualDeliveryDate() string`

GetActualDeliveryDate returns the ActualDeliveryDate field if non-nil, zero value otherwise.

### GetActualDeliveryDateOk

`func (o *UpdateMTOShipmentV2) GetActualDeliveryDateOk() (*string, bool)`

GetActualDeliveryDateOk returns a tuple with the ActualDeliveryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDeliveryDate

`func (o *UpdateMTOShipmentV2) SetActualDeliveryDate(v string)`

SetActualDeliveryDate sets ActualDeliveryDate field to given value.

### HasActualDeliveryDate

`func (o *UpdateMTOShipmentV2) HasActualDeliveryDate() bool`

HasActualDeliveryDate returns a boolean if a field has been set.

### SetActualDeliveryDateNil

`func (o *UpdateMTOShipmentV2) SetActualDeliveryDateNil(b bool)`

 SetActualDeliveryDateNil sets the value for ActualDeliveryDate to be an explicit nil

### UnsetActualDeliveryDate
`func (o *UpdateMTOShipmentV2) UnsetActualDeliveryDate()`

UnsetActualDeliveryDate ensures that no value is present for ActualDeliveryDate, not even an explicit nil
### GetPrimeEstimatedWeight

`func (o *UpdateMTOShipmentV2) GetPrimeEstimatedWeight() int32`

GetPrimeEstimatedWeight returns the PrimeEstimatedWeight field if non-nil, zero value otherwise.

### GetPrimeEstimatedWeightOk

`func (o *UpdateMTOShipmentV2) GetPrimeEstimatedWeightOk() (*int32, bool)`

GetPrimeEstimatedWeightOk returns a tuple with the PrimeEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeEstimatedWeight

`func (o *UpdateMTOShipmentV2) SetPrimeEstimatedWeight(v int32)`

SetPrimeEstimatedWeight sets PrimeEstimatedWeight field to given value.

### HasPrimeEstimatedWeight

`func (o *UpdateMTOShipmentV2) HasPrimeEstimatedWeight() bool`

HasPrimeEstimatedWeight returns a boolean if a field has been set.

### SetPrimeEstimatedWeightNil

`func (o *UpdateMTOShipmentV2) SetPrimeEstimatedWeightNil(b bool)`

 SetPrimeEstimatedWeightNil sets the value for PrimeEstimatedWeight to be an explicit nil

### UnsetPrimeEstimatedWeight
`func (o *UpdateMTOShipmentV2) UnsetPrimeEstimatedWeight()`

UnsetPrimeEstimatedWeight ensures that no value is present for PrimeEstimatedWeight, not even an explicit nil
### GetPrimeActualWeight

`func (o *UpdateMTOShipmentV2) GetPrimeActualWeight() int32`

GetPrimeActualWeight returns the PrimeActualWeight field if non-nil, zero value otherwise.

### GetPrimeActualWeightOk

`func (o *UpdateMTOShipmentV2) GetPrimeActualWeightOk() (*int32, bool)`

GetPrimeActualWeightOk returns a tuple with the PrimeActualWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimeActualWeight

`func (o *UpdateMTOShipmentV2) SetPrimeActualWeight(v int32)`

SetPrimeActualWeight sets PrimeActualWeight field to given value.

### HasPrimeActualWeight

`func (o *UpdateMTOShipmentV2) HasPrimeActualWeight() bool`

HasPrimeActualWeight returns a boolean if a field has been set.

### SetPrimeActualWeightNil

`func (o *UpdateMTOShipmentV2) SetPrimeActualWeightNil(b bool)`

 SetPrimeActualWeightNil sets the value for PrimeActualWeight to be an explicit nil

### UnsetPrimeActualWeight
`func (o *UpdateMTOShipmentV2) UnsetPrimeActualWeight()`

UnsetPrimeActualWeight ensures that no value is present for PrimeActualWeight, not even an explicit nil
### GetNtsRecordedWeight

`func (o *UpdateMTOShipmentV2) GetNtsRecordedWeight() int32`

GetNtsRecordedWeight returns the NtsRecordedWeight field if non-nil, zero value otherwise.

### GetNtsRecordedWeightOk

`func (o *UpdateMTOShipmentV2) GetNtsRecordedWeightOk() (*int32, bool)`

GetNtsRecordedWeightOk returns a tuple with the NtsRecordedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNtsRecordedWeight

`func (o *UpdateMTOShipmentV2) SetNtsRecordedWeight(v int32)`

SetNtsRecordedWeight sets NtsRecordedWeight field to given value.

### HasNtsRecordedWeight

`func (o *UpdateMTOShipmentV2) HasNtsRecordedWeight() bool`

HasNtsRecordedWeight returns a boolean if a field has been set.

### SetNtsRecordedWeightNil

`func (o *UpdateMTOShipmentV2) SetNtsRecordedWeightNil(b bool)`

 SetNtsRecordedWeightNil sets the value for NtsRecordedWeight to be an explicit nil

### UnsetNtsRecordedWeight
`func (o *UpdateMTOShipmentV2) UnsetNtsRecordedWeight()`

UnsetNtsRecordedWeight ensures that no value is present for NtsRecordedWeight, not even an explicit nil
### GetPickupAddress

`func (o *UpdateMTOShipmentV2) GetPickupAddress() AddressV2V2`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *UpdateMTOShipmentV2) GetPickupAddressOk() (*AddressV2V2, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *UpdateMTOShipmentV2) SetPickupAddress(v AddressV2V2)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *UpdateMTOShipmentV2) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *UpdateMTOShipmentV2) GetDestinationAddress() AddressV2V2`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *UpdateMTOShipmentV2) GetDestinationAddressOk() (*AddressV2V2, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *UpdateMTOShipmentV2) SetDestinationAddress(v AddressV2V2)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *UpdateMTOShipmentV2) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetDestinationType

`func (o *UpdateMTOShipmentV2) GetDestinationType() DestinationTypeV2V2`

GetDestinationType returns the DestinationType field if non-nil, zero value otherwise.

### GetDestinationTypeOk

`func (o *UpdateMTOShipmentV2) GetDestinationTypeOk() (*DestinationTypeV2V2, bool)`

GetDestinationTypeOk returns a tuple with the DestinationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationType

`func (o *UpdateMTOShipmentV2) SetDestinationType(v DestinationTypeV2V2)`

SetDestinationType sets DestinationType field to given value.

### HasDestinationType

`func (o *UpdateMTOShipmentV2) HasDestinationType() bool`

HasDestinationType returns a boolean if a field has been set.

### SetDestinationTypeNil

`func (o *UpdateMTOShipmentV2) SetDestinationTypeNil(b bool)`

 SetDestinationTypeNil sets the value for DestinationType to be an explicit nil

### UnsetDestinationType
`func (o *UpdateMTOShipmentV2) UnsetDestinationType()`

UnsetDestinationType ensures that no value is present for DestinationType, not even an explicit nil
### GetSecondaryPickupAddress

`func (o *UpdateMTOShipmentV2) GetSecondaryPickupAddress() AddressV2V2`

GetSecondaryPickupAddress returns the SecondaryPickupAddress field if non-nil, zero value otherwise.

### GetSecondaryPickupAddressOk

`func (o *UpdateMTOShipmentV2) GetSecondaryPickupAddressOk() (*AddressV2V2, bool)`

GetSecondaryPickupAddressOk returns a tuple with the SecondaryPickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryPickupAddress

`func (o *UpdateMTOShipmentV2) SetSecondaryPickupAddress(v AddressV2V2)`

SetSecondaryPickupAddress sets SecondaryPickupAddress field to given value.

### HasSecondaryPickupAddress

`func (o *UpdateMTOShipmentV2) HasSecondaryPickupAddress() bool`

HasSecondaryPickupAddress returns a boolean if a field has been set.

### GetSecondaryDeliveryAddress

`func (o *UpdateMTOShipmentV2) GetSecondaryDeliveryAddress() AddressV2V2`

GetSecondaryDeliveryAddress returns the SecondaryDeliveryAddress field if non-nil, zero value otherwise.

### GetSecondaryDeliveryAddressOk

`func (o *UpdateMTOShipmentV2) GetSecondaryDeliveryAddressOk() (*AddressV2V2, bool)`

GetSecondaryDeliveryAddressOk returns a tuple with the SecondaryDeliveryAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDeliveryAddress

`func (o *UpdateMTOShipmentV2) SetSecondaryDeliveryAddress(v AddressV2V2)`

SetSecondaryDeliveryAddress sets SecondaryDeliveryAddress field to given value.

### HasSecondaryDeliveryAddress

`func (o *UpdateMTOShipmentV2) HasSecondaryDeliveryAddress() bool`

HasSecondaryDeliveryAddress returns a boolean if a field has been set.

### GetStorageFacility

`func (o *UpdateMTOShipmentV2) GetStorageFacility() UpdateMTOShipmentStorageFacilityV2V2`

GetStorageFacility returns the StorageFacility field if non-nil, zero value otherwise.

### GetStorageFacilityOk

`func (o *UpdateMTOShipmentV2) GetStorageFacilityOk() (*UpdateMTOShipmentStorageFacilityV2V2, bool)`

GetStorageFacilityOk returns a tuple with the StorageFacility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageFacility

`func (o *UpdateMTOShipmentV2) SetStorageFacility(v UpdateMTOShipmentStorageFacilityV2V2)`

SetStorageFacility sets StorageFacility field to given value.

### HasStorageFacility

`func (o *UpdateMTOShipmentV2) HasStorageFacility() bool`

HasStorageFacility returns a boolean if a field has been set.

### GetShipmentType

`func (o *UpdateMTOShipmentV2) GetShipmentType() MTOShipmentTypeV2V2`

GetShipmentType returns the ShipmentType field if non-nil, zero value otherwise.

### GetShipmentTypeOk

`func (o *UpdateMTOShipmentV2) GetShipmentTypeOk() (*MTOShipmentTypeV2V2, bool)`

GetShipmentTypeOk returns a tuple with the ShipmentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentType

`func (o *UpdateMTOShipmentV2) SetShipmentType(v MTOShipmentTypeV2V2)`

SetShipmentType sets ShipmentType field to given value.

### HasShipmentType

`func (o *UpdateMTOShipmentV2) HasShipmentType() bool`

HasShipmentType returns a boolean if a field has been set.

### GetDiversion

`func (o *UpdateMTOShipmentV2) GetDiversion() bool`

GetDiversion returns the Diversion field if non-nil, zero value otherwise.

### GetDiversionOk

`func (o *UpdateMTOShipmentV2) GetDiversionOk() (*bool, bool)`

GetDiversionOk returns a tuple with the Diversion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDiversion

`func (o *UpdateMTOShipmentV2) SetDiversion(v bool)`

SetDiversion sets Diversion field to given value.

### HasDiversion

`func (o *UpdateMTOShipmentV2) HasDiversion() bool`

HasDiversion returns a boolean if a field has been set.

### GetPointOfContact

`func (o *UpdateMTOShipmentV2) GetPointOfContact() string`

GetPointOfContact returns the PointOfContact field if non-nil, zero value otherwise.

### GetPointOfContactOk

`func (o *UpdateMTOShipmentV2) GetPointOfContactOk() (*string, bool)`

GetPointOfContactOk returns a tuple with the PointOfContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointOfContact

`func (o *UpdateMTOShipmentV2) SetPointOfContact(v string)`

SetPointOfContact sets PointOfContact field to given value.

### HasPointOfContact

`func (o *UpdateMTOShipmentV2) HasPointOfContact() bool`

HasPointOfContact returns a boolean if a field has been set.

### GetCounselorRemarks

`func (o *UpdateMTOShipmentV2) GetCounselorRemarks() string`

GetCounselorRemarks returns the CounselorRemarks field if non-nil, zero value otherwise.

### GetCounselorRemarksOk

`func (o *UpdateMTOShipmentV2) GetCounselorRemarksOk() (*string, bool)`

GetCounselorRemarksOk returns a tuple with the CounselorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounselorRemarks

`func (o *UpdateMTOShipmentV2) SetCounselorRemarks(v string)`

SetCounselorRemarks sets CounselorRemarks field to given value.

### HasCounselorRemarks

`func (o *UpdateMTOShipmentV2) HasCounselorRemarks() bool`

HasCounselorRemarks returns a boolean if a field has been set.

### SetCounselorRemarksNil

`func (o *UpdateMTOShipmentV2) SetCounselorRemarksNil(b bool)`

 SetCounselorRemarksNil sets the value for CounselorRemarks to be an explicit nil

### UnsetCounselorRemarks
`func (o *UpdateMTOShipmentV2) UnsetCounselorRemarks()`

UnsetCounselorRemarks ensures that no value is present for CounselorRemarks, not even an explicit nil
### GetActualProGearWeight

`func (o *UpdateMTOShipmentV2) GetActualProGearWeight() int32`

GetActualProGearWeight returns the ActualProGearWeight field if non-nil, zero value otherwise.

### GetActualProGearWeightOk

`func (o *UpdateMTOShipmentV2) GetActualProGearWeightOk() (*int32, bool)`

GetActualProGearWeightOk returns a tuple with the ActualProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualProGearWeight

`func (o *UpdateMTOShipmentV2) SetActualProGearWeight(v int32)`

SetActualProGearWeight sets ActualProGearWeight field to given value.

### HasActualProGearWeight

`func (o *UpdateMTOShipmentV2) HasActualProGearWeight() bool`

HasActualProGearWeight returns a boolean if a field has been set.

### SetActualProGearWeightNil

`func (o *UpdateMTOShipmentV2) SetActualProGearWeightNil(b bool)`

 SetActualProGearWeightNil sets the value for ActualProGearWeight to be an explicit nil

### UnsetActualProGearWeight
`func (o *UpdateMTOShipmentV2) UnsetActualProGearWeight()`

UnsetActualProGearWeight ensures that no value is present for ActualProGearWeight, not even an explicit nil
### GetActualSpouseProGearWeight

`func (o *UpdateMTOShipmentV2) GetActualSpouseProGearWeight() int32`

GetActualSpouseProGearWeight returns the ActualSpouseProGearWeight field if non-nil, zero value otherwise.

### GetActualSpouseProGearWeightOk

`func (o *UpdateMTOShipmentV2) GetActualSpouseProGearWeightOk() (*int32, bool)`

GetActualSpouseProGearWeightOk returns a tuple with the ActualSpouseProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualSpouseProGearWeight

`func (o *UpdateMTOShipmentV2) SetActualSpouseProGearWeight(v int32)`

SetActualSpouseProGearWeight sets ActualSpouseProGearWeight field to given value.

### HasActualSpouseProGearWeight

`func (o *UpdateMTOShipmentV2) HasActualSpouseProGearWeight() bool`

HasActualSpouseProGearWeight returns a boolean if a field has been set.

### SetActualSpouseProGearWeightNil

`func (o *UpdateMTOShipmentV2) SetActualSpouseProGearWeightNil(b bool)`

 SetActualSpouseProGearWeightNil sets the value for ActualSpouseProGearWeight to be an explicit nil

### UnsetActualSpouseProGearWeight
`func (o *UpdateMTOShipmentV2) UnsetActualSpouseProGearWeight()`

UnsetActualSpouseProGearWeight ensures that no value is present for ActualSpouseProGearWeight, not even an explicit nil
### GetPpmShipment

`func (o *UpdateMTOShipmentV2) GetPpmShipment() UpdatePPMShipmentV2V2`

GetPpmShipment returns the PpmShipment field if non-nil, zero value otherwise.

### GetPpmShipmentOk

`func (o *UpdateMTOShipmentV2) GetPpmShipmentOk() (*UpdatePPMShipmentV2V2, bool)`

GetPpmShipmentOk returns a tuple with the PpmShipment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpmShipment

`func (o *UpdateMTOShipmentV2) SetPpmShipment(v UpdatePPMShipmentV2V2)`

SetPpmShipment sets PpmShipment field to given value.

### HasPpmShipment

`func (o *UpdateMTOShipmentV2) HasPpmShipment() bool`

HasPpmShipment returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


