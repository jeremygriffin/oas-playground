# UpdatePPMShipmentV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedDepartureDate** | Pointer to **NullableString** | Date the customer expects to begin moving from their origin.  | [optional] 
**PickupAddress** | Pointer to [**AddressV3V3**](AddressV3.md) | The address of the origin location where goods are being moved from.  | [optional] 
**HasSecondaryPickupAddress** | Pointer to **NullableBool** |  | [optional] 
**SecondaryPickupAddress** | Pointer to [**AddressV3V3**](AddressV3.md) | An optional secondary pickup location near the origin where additional goods exist.  | [optional] 
**DestinationAddress** | Pointer to [**AddressV3V3**](AddressV3.md) | The address of the destination location where goods are being delivered to.  | [optional] 
**HasSecondaryDestinationAddress** | Pointer to **NullableBool** |  | [optional] 
**SecondaryDestinationAddress** | Pointer to [**AddressV3V3**](AddressV3.md) | An optional secondary address near the destination where goods will be dropped off.  | [optional] 
**SitExpected** | Pointer to **NullableBool** | Captures whether some or all of the PPM shipment will require temporary storage at the origin or destination.  Must be set to &#x60;true&#x60; when providing &#x60;sitLocation&#x60;, &#x60;sitEstimatedWeight&#x60;, &#x60;sitEstimatedEntryDate&#x60;, and &#x60;sitEstimatedDepartureDate&#x60; values to calculate the &#x60;sitEstimatedCost&#x60;.  | [optional] 
**SitLocation** | Pointer to [**SITLocationTypeV3**](SITLocationType.md) |  | [optional] 
**SitEstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of the goods being put into storage. | [optional] 
**SitEstimatedEntryDate** | Pointer to **NullableString** | The date that goods will first enter the storage location. | [optional] 
**SitEstimatedDepartureDate** | Pointer to **NullableString** | The date that goods will exit the storage location. | [optional] 
**EstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of the PPM shipment goods being moved. | [optional] 
**HasProGear** | Pointer to **NullableBool** | Indicates whether PPM shipment has pro gear for themselves or their spouse.  | [optional] 
**ProGearWeight** | Pointer to **NullableInt32** | The estimated weight of the pro-gear being moved belonging to the service member. | [optional] 
**SpouseProGearWeight** | Pointer to **NullableInt32** | The estimated weight of the pro-gear being moved belonging to a spouse. | [optional] 

## Methods

### NewUpdatePPMShipmentV3

`func NewUpdatePPMShipmentV3() *UpdatePPMShipmentV3`

NewUpdatePPMShipmentV3 instantiates a new UpdatePPMShipmentV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePPMShipmentV3WithDefaults

`func NewUpdatePPMShipmentV3WithDefaults() *UpdatePPMShipmentV3`

NewUpdatePPMShipmentV3WithDefaults instantiates a new UpdatePPMShipmentV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedDepartureDate

`func (o *UpdatePPMShipmentV3) GetExpectedDepartureDate() string`

GetExpectedDepartureDate returns the ExpectedDepartureDate field if non-nil, zero value otherwise.

### GetExpectedDepartureDateOk

`func (o *UpdatePPMShipmentV3) GetExpectedDepartureDateOk() (*string, bool)`

GetExpectedDepartureDateOk returns a tuple with the ExpectedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDepartureDate

`func (o *UpdatePPMShipmentV3) SetExpectedDepartureDate(v string)`

SetExpectedDepartureDate sets ExpectedDepartureDate field to given value.

### HasExpectedDepartureDate

`func (o *UpdatePPMShipmentV3) HasExpectedDepartureDate() bool`

HasExpectedDepartureDate returns a boolean if a field has been set.

### SetExpectedDepartureDateNil

`func (o *UpdatePPMShipmentV3) SetExpectedDepartureDateNil(b bool)`

 SetExpectedDepartureDateNil sets the value for ExpectedDepartureDate to be an explicit nil

### UnsetExpectedDepartureDate
`func (o *UpdatePPMShipmentV3) UnsetExpectedDepartureDate()`

UnsetExpectedDepartureDate ensures that no value is present for ExpectedDepartureDate, not even an explicit nil
### GetPickupAddress

`func (o *UpdatePPMShipmentV3) GetPickupAddress() AddressV3V3`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *UpdatePPMShipmentV3) GetPickupAddressOk() (*AddressV3V3, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *UpdatePPMShipmentV3) SetPickupAddress(v AddressV3V3)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *UpdatePPMShipmentV3) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetHasSecondaryPickupAddress

`func (o *UpdatePPMShipmentV3) GetHasSecondaryPickupAddress() bool`

GetHasSecondaryPickupAddress returns the HasSecondaryPickupAddress field if non-nil, zero value otherwise.

### GetHasSecondaryPickupAddressOk

`func (o *UpdatePPMShipmentV3) GetHasSecondaryPickupAddressOk() (*bool, bool)`

GetHasSecondaryPickupAddressOk returns a tuple with the HasSecondaryPickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecondaryPickupAddress

`func (o *UpdatePPMShipmentV3) SetHasSecondaryPickupAddress(v bool)`

SetHasSecondaryPickupAddress sets HasSecondaryPickupAddress field to given value.

### HasHasSecondaryPickupAddress

`func (o *UpdatePPMShipmentV3) HasHasSecondaryPickupAddress() bool`

HasHasSecondaryPickupAddress returns a boolean if a field has been set.

### SetHasSecondaryPickupAddressNil

`func (o *UpdatePPMShipmentV3) SetHasSecondaryPickupAddressNil(b bool)`

 SetHasSecondaryPickupAddressNil sets the value for HasSecondaryPickupAddress to be an explicit nil

### UnsetHasSecondaryPickupAddress
`func (o *UpdatePPMShipmentV3) UnsetHasSecondaryPickupAddress()`

UnsetHasSecondaryPickupAddress ensures that no value is present for HasSecondaryPickupAddress, not even an explicit nil
### GetSecondaryPickupAddress

`func (o *UpdatePPMShipmentV3) GetSecondaryPickupAddress() AddressV3V3`

GetSecondaryPickupAddress returns the SecondaryPickupAddress field if non-nil, zero value otherwise.

### GetSecondaryPickupAddressOk

`func (o *UpdatePPMShipmentV3) GetSecondaryPickupAddressOk() (*AddressV3V3, bool)`

GetSecondaryPickupAddressOk returns a tuple with the SecondaryPickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryPickupAddress

`func (o *UpdatePPMShipmentV3) SetSecondaryPickupAddress(v AddressV3V3)`

SetSecondaryPickupAddress sets SecondaryPickupAddress field to given value.

### HasSecondaryPickupAddress

`func (o *UpdatePPMShipmentV3) HasSecondaryPickupAddress() bool`

HasSecondaryPickupAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *UpdatePPMShipmentV3) GetDestinationAddress() AddressV3V3`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *UpdatePPMShipmentV3) GetDestinationAddressOk() (*AddressV3V3, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *UpdatePPMShipmentV3) SetDestinationAddress(v AddressV3V3)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *UpdatePPMShipmentV3) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetHasSecondaryDestinationAddress

`func (o *UpdatePPMShipmentV3) GetHasSecondaryDestinationAddress() bool`

GetHasSecondaryDestinationAddress returns the HasSecondaryDestinationAddress field if non-nil, zero value otherwise.

### GetHasSecondaryDestinationAddressOk

`func (o *UpdatePPMShipmentV3) GetHasSecondaryDestinationAddressOk() (*bool, bool)`

GetHasSecondaryDestinationAddressOk returns a tuple with the HasSecondaryDestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecondaryDestinationAddress

`func (o *UpdatePPMShipmentV3) SetHasSecondaryDestinationAddress(v bool)`

SetHasSecondaryDestinationAddress sets HasSecondaryDestinationAddress field to given value.

### HasHasSecondaryDestinationAddress

`func (o *UpdatePPMShipmentV3) HasHasSecondaryDestinationAddress() bool`

HasHasSecondaryDestinationAddress returns a boolean if a field has been set.

### SetHasSecondaryDestinationAddressNil

`func (o *UpdatePPMShipmentV3) SetHasSecondaryDestinationAddressNil(b bool)`

 SetHasSecondaryDestinationAddressNil sets the value for HasSecondaryDestinationAddress to be an explicit nil

### UnsetHasSecondaryDestinationAddress
`func (o *UpdatePPMShipmentV3) UnsetHasSecondaryDestinationAddress()`

UnsetHasSecondaryDestinationAddress ensures that no value is present for HasSecondaryDestinationAddress, not even an explicit nil
### GetSecondaryDestinationAddress

`func (o *UpdatePPMShipmentV3) GetSecondaryDestinationAddress() AddressV3V3`

GetSecondaryDestinationAddress returns the SecondaryDestinationAddress field if non-nil, zero value otherwise.

### GetSecondaryDestinationAddressOk

`func (o *UpdatePPMShipmentV3) GetSecondaryDestinationAddressOk() (*AddressV3V3, bool)`

GetSecondaryDestinationAddressOk returns a tuple with the SecondaryDestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDestinationAddress

`func (o *UpdatePPMShipmentV3) SetSecondaryDestinationAddress(v AddressV3V3)`

SetSecondaryDestinationAddress sets SecondaryDestinationAddress field to given value.

### HasSecondaryDestinationAddress

`func (o *UpdatePPMShipmentV3) HasSecondaryDestinationAddress() bool`

HasSecondaryDestinationAddress returns a boolean if a field has been set.

### GetSitExpected

`func (o *UpdatePPMShipmentV3) GetSitExpected() bool`

GetSitExpected returns the SitExpected field if non-nil, zero value otherwise.

### GetSitExpectedOk

`func (o *UpdatePPMShipmentV3) GetSitExpectedOk() (*bool, bool)`

GetSitExpectedOk returns a tuple with the SitExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitExpected

`func (o *UpdatePPMShipmentV3) SetSitExpected(v bool)`

SetSitExpected sets SitExpected field to given value.

### HasSitExpected

`func (o *UpdatePPMShipmentV3) HasSitExpected() bool`

HasSitExpected returns a boolean if a field has been set.

### SetSitExpectedNil

`func (o *UpdatePPMShipmentV3) SetSitExpectedNil(b bool)`

 SetSitExpectedNil sets the value for SitExpected to be an explicit nil

### UnsetSitExpected
`func (o *UpdatePPMShipmentV3) UnsetSitExpected()`

UnsetSitExpected ensures that no value is present for SitExpected, not even an explicit nil
### GetSitLocation

`func (o *UpdatePPMShipmentV3) GetSitLocation() SITLocationTypeV3`

GetSitLocation returns the SitLocation field if non-nil, zero value otherwise.

### GetSitLocationOk

`func (o *UpdatePPMShipmentV3) GetSitLocationOk() (*SITLocationTypeV3, bool)`

GetSitLocationOk returns a tuple with the SitLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitLocation

`func (o *UpdatePPMShipmentV3) SetSitLocation(v SITLocationTypeV3)`

SetSitLocation sets SitLocation field to given value.

### HasSitLocation

`func (o *UpdatePPMShipmentV3) HasSitLocation() bool`

HasSitLocation returns a boolean if a field has been set.

### GetSitEstimatedWeight

`func (o *UpdatePPMShipmentV3) GetSitEstimatedWeight() int32`

GetSitEstimatedWeight returns the SitEstimatedWeight field if non-nil, zero value otherwise.

### GetSitEstimatedWeightOk

`func (o *UpdatePPMShipmentV3) GetSitEstimatedWeightOk() (*int32, bool)`

GetSitEstimatedWeightOk returns a tuple with the SitEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedWeight

`func (o *UpdatePPMShipmentV3) SetSitEstimatedWeight(v int32)`

SetSitEstimatedWeight sets SitEstimatedWeight field to given value.

### HasSitEstimatedWeight

`func (o *UpdatePPMShipmentV3) HasSitEstimatedWeight() bool`

HasSitEstimatedWeight returns a boolean if a field has been set.

### SetSitEstimatedWeightNil

`func (o *UpdatePPMShipmentV3) SetSitEstimatedWeightNil(b bool)`

 SetSitEstimatedWeightNil sets the value for SitEstimatedWeight to be an explicit nil

### UnsetSitEstimatedWeight
`func (o *UpdatePPMShipmentV3) UnsetSitEstimatedWeight()`

UnsetSitEstimatedWeight ensures that no value is present for SitEstimatedWeight, not even an explicit nil
### GetSitEstimatedEntryDate

`func (o *UpdatePPMShipmentV3) GetSitEstimatedEntryDate() string`

GetSitEstimatedEntryDate returns the SitEstimatedEntryDate field if non-nil, zero value otherwise.

### GetSitEstimatedEntryDateOk

`func (o *UpdatePPMShipmentV3) GetSitEstimatedEntryDateOk() (*string, bool)`

GetSitEstimatedEntryDateOk returns a tuple with the SitEstimatedEntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedEntryDate

`func (o *UpdatePPMShipmentV3) SetSitEstimatedEntryDate(v string)`

SetSitEstimatedEntryDate sets SitEstimatedEntryDate field to given value.

### HasSitEstimatedEntryDate

`func (o *UpdatePPMShipmentV3) HasSitEstimatedEntryDate() bool`

HasSitEstimatedEntryDate returns a boolean if a field has been set.

### SetSitEstimatedEntryDateNil

`func (o *UpdatePPMShipmentV3) SetSitEstimatedEntryDateNil(b bool)`

 SetSitEstimatedEntryDateNil sets the value for SitEstimatedEntryDate to be an explicit nil

### UnsetSitEstimatedEntryDate
`func (o *UpdatePPMShipmentV3) UnsetSitEstimatedEntryDate()`

UnsetSitEstimatedEntryDate ensures that no value is present for SitEstimatedEntryDate, not even an explicit nil
### GetSitEstimatedDepartureDate

`func (o *UpdatePPMShipmentV3) GetSitEstimatedDepartureDate() string`

GetSitEstimatedDepartureDate returns the SitEstimatedDepartureDate field if non-nil, zero value otherwise.

### GetSitEstimatedDepartureDateOk

`func (o *UpdatePPMShipmentV3) GetSitEstimatedDepartureDateOk() (*string, bool)`

GetSitEstimatedDepartureDateOk returns a tuple with the SitEstimatedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedDepartureDate

`func (o *UpdatePPMShipmentV3) SetSitEstimatedDepartureDate(v string)`

SetSitEstimatedDepartureDate sets SitEstimatedDepartureDate field to given value.

### HasSitEstimatedDepartureDate

`func (o *UpdatePPMShipmentV3) HasSitEstimatedDepartureDate() bool`

HasSitEstimatedDepartureDate returns a boolean if a field has been set.

### SetSitEstimatedDepartureDateNil

`func (o *UpdatePPMShipmentV3) SetSitEstimatedDepartureDateNil(b bool)`

 SetSitEstimatedDepartureDateNil sets the value for SitEstimatedDepartureDate to be an explicit nil

### UnsetSitEstimatedDepartureDate
`func (o *UpdatePPMShipmentV3) UnsetSitEstimatedDepartureDate()`

UnsetSitEstimatedDepartureDate ensures that no value is present for SitEstimatedDepartureDate, not even an explicit nil
### GetEstimatedWeight

`func (o *UpdatePPMShipmentV3) GetEstimatedWeight() int32`

GetEstimatedWeight returns the EstimatedWeight field if non-nil, zero value otherwise.

### GetEstimatedWeightOk

`func (o *UpdatePPMShipmentV3) GetEstimatedWeightOk() (*int32, bool)`

GetEstimatedWeightOk returns a tuple with the EstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedWeight

`func (o *UpdatePPMShipmentV3) SetEstimatedWeight(v int32)`

SetEstimatedWeight sets EstimatedWeight field to given value.

### HasEstimatedWeight

`func (o *UpdatePPMShipmentV3) HasEstimatedWeight() bool`

HasEstimatedWeight returns a boolean if a field has been set.

### SetEstimatedWeightNil

`func (o *UpdatePPMShipmentV3) SetEstimatedWeightNil(b bool)`

 SetEstimatedWeightNil sets the value for EstimatedWeight to be an explicit nil

### UnsetEstimatedWeight
`func (o *UpdatePPMShipmentV3) UnsetEstimatedWeight()`

UnsetEstimatedWeight ensures that no value is present for EstimatedWeight, not even an explicit nil
### GetHasProGear

`func (o *UpdatePPMShipmentV3) GetHasProGear() bool`

GetHasProGear returns the HasProGear field if non-nil, zero value otherwise.

### GetHasProGearOk

`func (o *UpdatePPMShipmentV3) GetHasProGearOk() (*bool, bool)`

GetHasProGearOk returns a tuple with the HasProGear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProGear

`func (o *UpdatePPMShipmentV3) SetHasProGear(v bool)`

SetHasProGear sets HasProGear field to given value.

### HasHasProGear

`func (o *UpdatePPMShipmentV3) HasHasProGear() bool`

HasHasProGear returns a boolean if a field has been set.

### SetHasProGearNil

`func (o *UpdatePPMShipmentV3) SetHasProGearNil(b bool)`

 SetHasProGearNil sets the value for HasProGear to be an explicit nil

### UnsetHasProGear
`func (o *UpdatePPMShipmentV3) UnsetHasProGear()`

UnsetHasProGear ensures that no value is present for HasProGear, not even an explicit nil
### GetProGearWeight

`func (o *UpdatePPMShipmentV3) GetProGearWeight() int32`

GetProGearWeight returns the ProGearWeight field if non-nil, zero value otherwise.

### GetProGearWeightOk

`func (o *UpdatePPMShipmentV3) GetProGearWeightOk() (*int32, bool)`

GetProGearWeightOk returns a tuple with the ProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProGearWeight

`func (o *UpdatePPMShipmentV3) SetProGearWeight(v int32)`

SetProGearWeight sets ProGearWeight field to given value.

### HasProGearWeight

`func (o *UpdatePPMShipmentV3) HasProGearWeight() bool`

HasProGearWeight returns a boolean if a field has been set.

### SetProGearWeightNil

`func (o *UpdatePPMShipmentV3) SetProGearWeightNil(b bool)`

 SetProGearWeightNil sets the value for ProGearWeight to be an explicit nil

### UnsetProGearWeight
`func (o *UpdatePPMShipmentV3) UnsetProGearWeight()`

UnsetProGearWeight ensures that no value is present for ProGearWeight, not even an explicit nil
### GetSpouseProGearWeight

`func (o *UpdatePPMShipmentV3) GetSpouseProGearWeight() int32`

GetSpouseProGearWeight returns the SpouseProGearWeight field if non-nil, zero value otherwise.

### GetSpouseProGearWeightOk

`func (o *UpdatePPMShipmentV3) GetSpouseProGearWeightOk() (*int32, bool)`

GetSpouseProGearWeightOk returns a tuple with the SpouseProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpouseProGearWeight

`func (o *UpdatePPMShipmentV3) SetSpouseProGearWeight(v int32)`

SetSpouseProGearWeight sets SpouseProGearWeight field to given value.

### HasSpouseProGearWeight

`func (o *UpdatePPMShipmentV3) HasSpouseProGearWeight() bool`

HasSpouseProGearWeight returns a boolean if a field has been set.

### SetSpouseProGearWeightNil

`func (o *UpdatePPMShipmentV3) SetSpouseProGearWeightNil(b bool)`

 SetSpouseProGearWeightNil sets the value for SpouseProGearWeight to be an explicit nil

### UnsetSpouseProGearWeight
`func (o *UpdatePPMShipmentV3) UnsetSpouseProGearWeight()`

UnsetSpouseProGearWeight ensures that no value is present for SpouseProGearWeight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


