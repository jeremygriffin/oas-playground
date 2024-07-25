# CreatePPMShipmentV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedDepartureDate** | **string** | Date the customer expects to begin moving from their origin.  | 
**PickupAddress** | [**AddressV3V3**](AddressV3.md) | The address of the origin location where goods are being moved from. | 
**SecondaryPickupAddress** | Pointer to [**AddressV3V3**](AddressV3.md) | An optional secondary pickup location address near the origin where additional goods exist. | [optional] 
**DestinationAddress** | [**AddressV3V3**](AddressV3.md) | The address of the destination location where goods are being delivered to. | 
**SecondaryDestinationAddress** | Pointer to [**AddressV3V3**](AddressV3.md) | An optional secondary address near the destination where goods will be dropped off. | [optional] 
**SitExpected** | **bool** | Captures whether some or all of the PPM shipment will require temporary storage at the origin or destination.  Must be set to &#x60;true&#x60; when providing &#x60;sitLocation&#x60;, &#x60;sitEstimatedWeight&#x60;, &#x60;sitEstimatedEntryDate&#x60;, and &#x60;sitEstimatedDepartureDate&#x60; values to calculate the &#x60;sitEstimatedCost&#x60;.  | 
**SitLocation** | Pointer to [**SITLocationTypeV3**](SITLocationType.md) |  | [optional] 
**SitEstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of the goods being put into storage in pounds. | [optional] 
**SitEstimatedEntryDate** | Pointer to **NullableString** | The date that goods will first enter the storage location. | [optional] 
**SitEstimatedDepartureDate** | Pointer to **NullableString** | The date that goods will exit the storage location. | [optional] 
**EstimatedWeight** | **int32** | The estimated weight of the PPM shipment goods being moved in pounds. | 
**HasProGear** | **bool** | Indicates whether PPM shipment has pro gear for themselves or their spouse.  | 
**ProGearWeight** | Pointer to **NullableInt32** | The estimated weight of the pro-gear being moved belonging to the service member in pounds. | [optional] 
**SpouseProGearWeight** | Pointer to **NullableInt32** | The estimated weight of the pro-gear being moved belonging to a spouse in pounds. | [optional] 

## Methods

### NewCreatePPMShipmentV3

`func NewCreatePPMShipmentV3(expectedDepartureDate string, pickupAddress AddressV3V3, destinationAddress AddressV3V3, sitExpected bool, estimatedWeight int32, hasProGear bool, ) *CreatePPMShipmentV3`

NewCreatePPMShipmentV3 instantiates a new CreatePPMShipmentV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePPMShipmentV3WithDefaults

`func NewCreatePPMShipmentV3WithDefaults() *CreatePPMShipmentV3`

NewCreatePPMShipmentV3WithDefaults instantiates a new CreatePPMShipmentV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedDepartureDate

`func (o *CreatePPMShipmentV3) GetExpectedDepartureDate() string`

GetExpectedDepartureDate returns the ExpectedDepartureDate field if non-nil, zero value otherwise.

### GetExpectedDepartureDateOk

`func (o *CreatePPMShipmentV3) GetExpectedDepartureDateOk() (*string, bool)`

GetExpectedDepartureDateOk returns a tuple with the ExpectedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDepartureDate

`func (o *CreatePPMShipmentV3) SetExpectedDepartureDate(v string)`

SetExpectedDepartureDate sets ExpectedDepartureDate field to given value.


### GetPickupAddress

`func (o *CreatePPMShipmentV3) GetPickupAddress() AddressV3V3`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *CreatePPMShipmentV3) GetPickupAddressOk() (*AddressV3V3, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *CreatePPMShipmentV3) SetPickupAddress(v AddressV3V3)`

SetPickupAddress sets PickupAddress field to given value.


### GetSecondaryPickupAddress

`func (o *CreatePPMShipmentV3) GetSecondaryPickupAddress() AddressV3V3`

GetSecondaryPickupAddress returns the SecondaryPickupAddress field if non-nil, zero value otherwise.

### GetSecondaryPickupAddressOk

`func (o *CreatePPMShipmentV3) GetSecondaryPickupAddressOk() (*AddressV3V3, bool)`

GetSecondaryPickupAddressOk returns a tuple with the SecondaryPickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryPickupAddress

`func (o *CreatePPMShipmentV3) SetSecondaryPickupAddress(v AddressV3V3)`

SetSecondaryPickupAddress sets SecondaryPickupAddress field to given value.

### HasSecondaryPickupAddress

`func (o *CreatePPMShipmentV3) HasSecondaryPickupAddress() bool`

HasSecondaryPickupAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *CreatePPMShipmentV3) GetDestinationAddress() AddressV3V3`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *CreatePPMShipmentV3) GetDestinationAddressOk() (*AddressV3V3, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *CreatePPMShipmentV3) SetDestinationAddress(v AddressV3V3)`

SetDestinationAddress sets DestinationAddress field to given value.


### GetSecondaryDestinationAddress

`func (o *CreatePPMShipmentV3) GetSecondaryDestinationAddress() AddressV3V3`

GetSecondaryDestinationAddress returns the SecondaryDestinationAddress field if non-nil, zero value otherwise.

### GetSecondaryDestinationAddressOk

`func (o *CreatePPMShipmentV3) GetSecondaryDestinationAddressOk() (*AddressV3V3, bool)`

GetSecondaryDestinationAddressOk returns a tuple with the SecondaryDestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondaryDestinationAddress

`func (o *CreatePPMShipmentV3) SetSecondaryDestinationAddress(v AddressV3V3)`

SetSecondaryDestinationAddress sets SecondaryDestinationAddress field to given value.

### HasSecondaryDestinationAddress

`func (o *CreatePPMShipmentV3) HasSecondaryDestinationAddress() bool`

HasSecondaryDestinationAddress returns a boolean if a field has been set.

### GetSitExpected

`func (o *CreatePPMShipmentV3) GetSitExpected() bool`

GetSitExpected returns the SitExpected field if non-nil, zero value otherwise.

### GetSitExpectedOk

`func (o *CreatePPMShipmentV3) GetSitExpectedOk() (*bool, bool)`

GetSitExpectedOk returns a tuple with the SitExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitExpected

`func (o *CreatePPMShipmentV3) SetSitExpected(v bool)`

SetSitExpected sets SitExpected field to given value.


### GetSitLocation

`func (o *CreatePPMShipmentV3) GetSitLocation() SITLocationTypeV3`

GetSitLocation returns the SitLocation field if non-nil, zero value otherwise.

### GetSitLocationOk

`func (o *CreatePPMShipmentV3) GetSitLocationOk() (*SITLocationTypeV3, bool)`

GetSitLocationOk returns a tuple with the SitLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitLocation

`func (o *CreatePPMShipmentV3) SetSitLocation(v SITLocationTypeV3)`

SetSitLocation sets SitLocation field to given value.

### HasSitLocation

`func (o *CreatePPMShipmentV3) HasSitLocation() bool`

HasSitLocation returns a boolean if a field has been set.

### GetSitEstimatedWeight

`func (o *CreatePPMShipmentV3) GetSitEstimatedWeight() int32`

GetSitEstimatedWeight returns the SitEstimatedWeight field if non-nil, zero value otherwise.

### GetSitEstimatedWeightOk

`func (o *CreatePPMShipmentV3) GetSitEstimatedWeightOk() (*int32, bool)`

GetSitEstimatedWeightOk returns a tuple with the SitEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedWeight

`func (o *CreatePPMShipmentV3) SetSitEstimatedWeight(v int32)`

SetSitEstimatedWeight sets SitEstimatedWeight field to given value.

### HasSitEstimatedWeight

`func (o *CreatePPMShipmentV3) HasSitEstimatedWeight() bool`

HasSitEstimatedWeight returns a boolean if a field has been set.

### SetSitEstimatedWeightNil

`func (o *CreatePPMShipmentV3) SetSitEstimatedWeightNil(b bool)`

 SetSitEstimatedWeightNil sets the value for SitEstimatedWeight to be an explicit nil

### UnsetSitEstimatedWeight
`func (o *CreatePPMShipmentV3) UnsetSitEstimatedWeight()`

UnsetSitEstimatedWeight ensures that no value is present for SitEstimatedWeight, not even an explicit nil
### GetSitEstimatedEntryDate

`func (o *CreatePPMShipmentV3) GetSitEstimatedEntryDate() string`

GetSitEstimatedEntryDate returns the SitEstimatedEntryDate field if non-nil, zero value otherwise.

### GetSitEstimatedEntryDateOk

`func (o *CreatePPMShipmentV3) GetSitEstimatedEntryDateOk() (*string, bool)`

GetSitEstimatedEntryDateOk returns a tuple with the SitEstimatedEntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedEntryDate

`func (o *CreatePPMShipmentV3) SetSitEstimatedEntryDate(v string)`

SetSitEstimatedEntryDate sets SitEstimatedEntryDate field to given value.

### HasSitEstimatedEntryDate

`func (o *CreatePPMShipmentV3) HasSitEstimatedEntryDate() bool`

HasSitEstimatedEntryDate returns a boolean if a field has been set.

### SetSitEstimatedEntryDateNil

`func (o *CreatePPMShipmentV3) SetSitEstimatedEntryDateNil(b bool)`

 SetSitEstimatedEntryDateNil sets the value for SitEstimatedEntryDate to be an explicit nil

### UnsetSitEstimatedEntryDate
`func (o *CreatePPMShipmentV3) UnsetSitEstimatedEntryDate()`

UnsetSitEstimatedEntryDate ensures that no value is present for SitEstimatedEntryDate, not even an explicit nil
### GetSitEstimatedDepartureDate

`func (o *CreatePPMShipmentV3) GetSitEstimatedDepartureDate() string`

GetSitEstimatedDepartureDate returns the SitEstimatedDepartureDate field if non-nil, zero value otherwise.

### GetSitEstimatedDepartureDateOk

`func (o *CreatePPMShipmentV3) GetSitEstimatedDepartureDateOk() (*string, bool)`

GetSitEstimatedDepartureDateOk returns a tuple with the SitEstimatedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedDepartureDate

`func (o *CreatePPMShipmentV3) SetSitEstimatedDepartureDate(v string)`

SetSitEstimatedDepartureDate sets SitEstimatedDepartureDate field to given value.

### HasSitEstimatedDepartureDate

`func (o *CreatePPMShipmentV3) HasSitEstimatedDepartureDate() bool`

HasSitEstimatedDepartureDate returns a boolean if a field has been set.

### SetSitEstimatedDepartureDateNil

`func (o *CreatePPMShipmentV3) SetSitEstimatedDepartureDateNil(b bool)`

 SetSitEstimatedDepartureDateNil sets the value for SitEstimatedDepartureDate to be an explicit nil

### UnsetSitEstimatedDepartureDate
`func (o *CreatePPMShipmentV3) UnsetSitEstimatedDepartureDate()`

UnsetSitEstimatedDepartureDate ensures that no value is present for SitEstimatedDepartureDate, not even an explicit nil
### GetEstimatedWeight

`func (o *CreatePPMShipmentV3) GetEstimatedWeight() int32`

GetEstimatedWeight returns the EstimatedWeight field if non-nil, zero value otherwise.

### GetEstimatedWeightOk

`func (o *CreatePPMShipmentV3) GetEstimatedWeightOk() (*int32, bool)`

GetEstimatedWeightOk returns a tuple with the EstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedWeight

`func (o *CreatePPMShipmentV3) SetEstimatedWeight(v int32)`

SetEstimatedWeight sets EstimatedWeight field to given value.


### GetHasProGear

`func (o *CreatePPMShipmentV3) GetHasProGear() bool`

GetHasProGear returns the HasProGear field if non-nil, zero value otherwise.

### GetHasProGearOk

`func (o *CreatePPMShipmentV3) GetHasProGearOk() (*bool, bool)`

GetHasProGearOk returns a tuple with the HasProGear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProGear

`func (o *CreatePPMShipmentV3) SetHasProGear(v bool)`

SetHasProGear sets HasProGear field to given value.


### GetProGearWeight

`func (o *CreatePPMShipmentV3) GetProGearWeight() int32`

GetProGearWeight returns the ProGearWeight field if non-nil, zero value otherwise.

### GetProGearWeightOk

`func (o *CreatePPMShipmentV3) GetProGearWeightOk() (*int32, bool)`

GetProGearWeightOk returns a tuple with the ProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProGearWeight

`func (o *CreatePPMShipmentV3) SetProGearWeight(v int32)`

SetProGearWeight sets ProGearWeight field to given value.

### HasProGearWeight

`func (o *CreatePPMShipmentV3) HasProGearWeight() bool`

HasProGearWeight returns a boolean if a field has been set.

### SetProGearWeightNil

`func (o *CreatePPMShipmentV3) SetProGearWeightNil(b bool)`

 SetProGearWeightNil sets the value for ProGearWeight to be an explicit nil

### UnsetProGearWeight
`func (o *CreatePPMShipmentV3) UnsetProGearWeight()`

UnsetProGearWeight ensures that no value is present for ProGearWeight, not even an explicit nil
### GetSpouseProGearWeight

`func (o *CreatePPMShipmentV3) GetSpouseProGearWeight() int32`

GetSpouseProGearWeight returns the SpouseProGearWeight field if non-nil, zero value otherwise.

### GetSpouseProGearWeightOk

`func (o *CreatePPMShipmentV3) GetSpouseProGearWeightOk() (*int32, bool)`

GetSpouseProGearWeightOk returns a tuple with the SpouseProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpouseProGearWeight

`func (o *CreatePPMShipmentV3) SetSpouseProGearWeight(v int32)`

SetSpouseProGearWeight sets SpouseProGearWeight field to given value.

### HasSpouseProGearWeight

`func (o *CreatePPMShipmentV3) HasSpouseProGearWeight() bool`

HasSpouseProGearWeight returns a boolean if a field has been set.

### SetSpouseProGearWeightNil

`func (o *CreatePPMShipmentV3) SetSpouseProGearWeightNil(b bool)`

 SetSpouseProGearWeightNil sets the value for SpouseProGearWeight to be an explicit nil

### UnsetSpouseProGearWeight
`func (o *CreatePPMShipmentV3) UnsetSpouseProGearWeight()`

UnsetSpouseProGearWeight ensures that no value is present for SpouseProGearWeight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


