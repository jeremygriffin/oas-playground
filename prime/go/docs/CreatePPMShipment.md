# CreatePPMShipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedDepartureDate** | **string** | Date the customer expects to begin moving from their origin.  | 
**PickupAddress** | Pointer to [**Address**](Address.md) | The address of the origin location where goods are being moved from. | [optional] 
**DestinationAddress** | Pointer to [**Address**](Address.md) | The address of the destination location where goods are being delivered to. | [optional] 
**SitExpected** | **bool** | Captures whether some or all of the PPM shipment will require temporary storage at the origin or destination.  Must be set to &#x60;true&#x60; when providing &#x60;sitLocation&#x60;, &#x60;sitEstimatedWeight&#x60;, &#x60;sitEstimatedEntryDate&#x60;, and &#x60;sitEstimatedDepartureDate&#x60; values to calculate the &#x60;sitEstimatedCost&#x60;.  | 
**SitLocation** | Pointer to [**SITLocationType**](SITLocationType.md) |  | [optional] 
**SitEstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of the goods being put into storage in pounds. | [optional] 
**SitEstimatedEntryDate** | Pointer to **NullableString** | The date that goods will first enter the storage location. | [optional] 
**SitEstimatedDepartureDate** | Pointer to **NullableString** | The date that goods will exit the storage location. | [optional] 
**EstimatedWeight** | **int32** | The estimated weight of the PPM shipment goods being moved in pounds. | 
**HasProGear** | **bool** | Indicates whether PPM shipment has pro gear for themselves or their spouse.  | 
**ProGearWeight** | Pointer to **NullableInt32** | The estimated weight of the pro-gear being moved belonging to the service member in pounds. | [optional] 
**SpouseProGearWeight** | Pointer to **NullableInt32** | The estimated weight of the pro-gear being moved belonging to a spouse in pounds. | [optional] 

## Methods

### NewCreatePPMShipment

`func NewCreatePPMShipment(expectedDepartureDate string, sitExpected bool, estimatedWeight int32, hasProGear bool, ) *CreatePPMShipment`

NewCreatePPMShipment instantiates a new CreatePPMShipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePPMShipmentWithDefaults

`func NewCreatePPMShipmentWithDefaults() *CreatePPMShipment`

NewCreatePPMShipmentWithDefaults instantiates a new CreatePPMShipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedDepartureDate

`func (o *CreatePPMShipment) GetExpectedDepartureDate() string`

GetExpectedDepartureDate returns the ExpectedDepartureDate field if non-nil, zero value otherwise.

### GetExpectedDepartureDateOk

`func (o *CreatePPMShipment) GetExpectedDepartureDateOk() (*string, bool)`

GetExpectedDepartureDateOk returns a tuple with the ExpectedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDepartureDate

`func (o *CreatePPMShipment) SetExpectedDepartureDate(v string)`

SetExpectedDepartureDate sets ExpectedDepartureDate field to given value.


### GetPickupAddress

`func (o *CreatePPMShipment) GetPickupAddress() Address`

GetPickupAddress returns the PickupAddress field if non-nil, zero value otherwise.

### GetPickupAddressOk

`func (o *CreatePPMShipment) GetPickupAddressOk() (*Address, bool)`

GetPickupAddressOk returns a tuple with the PickupAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPickupAddress

`func (o *CreatePPMShipment) SetPickupAddress(v Address)`

SetPickupAddress sets PickupAddress field to given value.

### HasPickupAddress

`func (o *CreatePPMShipment) HasPickupAddress() bool`

HasPickupAddress returns a boolean if a field has been set.

### GetDestinationAddress

`func (o *CreatePPMShipment) GetDestinationAddress() Address`

GetDestinationAddress returns the DestinationAddress field if non-nil, zero value otherwise.

### GetDestinationAddressOk

`func (o *CreatePPMShipment) GetDestinationAddressOk() (*Address, bool)`

GetDestinationAddressOk returns a tuple with the DestinationAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationAddress

`func (o *CreatePPMShipment) SetDestinationAddress(v Address)`

SetDestinationAddress sets DestinationAddress field to given value.

### HasDestinationAddress

`func (o *CreatePPMShipment) HasDestinationAddress() bool`

HasDestinationAddress returns a boolean if a field has been set.

### GetSitExpected

`func (o *CreatePPMShipment) GetSitExpected() bool`

GetSitExpected returns the SitExpected field if non-nil, zero value otherwise.

### GetSitExpectedOk

`func (o *CreatePPMShipment) GetSitExpectedOk() (*bool, bool)`

GetSitExpectedOk returns a tuple with the SitExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitExpected

`func (o *CreatePPMShipment) SetSitExpected(v bool)`

SetSitExpected sets SitExpected field to given value.


### GetSitLocation

`func (o *CreatePPMShipment) GetSitLocation() SITLocationType`

GetSitLocation returns the SitLocation field if non-nil, zero value otherwise.

### GetSitLocationOk

`func (o *CreatePPMShipment) GetSitLocationOk() (*SITLocationType, bool)`

GetSitLocationOk returns a tuple with the SitLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitLocation

`func (o *CreatePPMShipment) SetSitLocation(v SITLocationType)`

SetSitLocation sets SitLocation field to given value.

### HasSitLocation

`func (o *CreatePPMShipment) HasSitLocation() bool`

HasSitLocation returns a boolean if a field has been set.

### GetSitEstimatedWeight

`func (o *CreatePPMShipment) GetSitEstimatedWeight() int32`

GetSitEstimatedWeight returns the SitEstimatedWeight field if non-nil, zero value otherwise.

### GetSitEstimatedWeightOk

`func (o *CreatePPMShipment) GetSitEstimatedWeightOk() (*int32, bool)`

GetSitEstimatedWeightOk returns a tuple with the SitEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedWeight

`func (o *CreatePPMShipment) SetSitEstimatedWeight(v int32)`

SetSitEstimatedWeight sets SitEstimatedWeight field to given value.

### HasSitEstimatedWeight

`func (o *CreatePPMShipment) HasSitEstimatedWeight() bool`

HasSitEstimatedWeight returns a boolean if a field has been set.

### SetSitEstimatedWeightNil

`func (o *CreatePPMShipment) SetSitEstimatedWeightNil(b bool)`

 SetSitEstimatedWeightNil sets the value for SitEstimatedWeight to be an explicit nil

### UnsetSitEstimatedWeight
`func (o *CreatePPMShipment) UnsetSitEstimatedWeight()`

UnsetSitEstimatedWeight ensures that no value is present for SitEstimatedWeight, not even an explicit nil
### GetSitEstimatedEntryDate

`func (o *CreatePPMShipment) GetSitEstimatedEntryDate() string`

GetSitEstimatedEntryDate returns the SitEstimatedEntryDate field if non-nil, zero value otherwise.

### GetSitEstimatedEntryDateOk

`func (o *CreatePPMShipment) GetSitEstimatedEntryDateOk() (*string, bool)`

GetSitEstimatedEntryDateOk returns a tuple with the SitEstimatedEntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedEntryDate

`func (o *CreatePPMShipment) SetSitEstimatedEntryDate(v string)`

SetSitEstimatedEntryDate sets SitEstimatedEntryDate field to given value.

### HasSitEstimatedEntryDate

`func (o *CreatePPMShipment) HasSitEstimatedEntryDate() bool`

HasSitEstimatedEntryDate returns a boolean if a field has been set.

### SetSitEstimatedEntryDateNil

`func (o *CreatePPMShipment) SetSitEstimatedEntryDateNil(b bool)`

 SetSitEstimatedEntryDateNil sets the value for SitEstimatedEntryDate to be an explicit nil

### UnsetSitEstimatedEntryDate
`func (o *CreatePPMShipment) UnsetSitEstimatedEntryDate()`

UnsetSitEstimatedEntryDate ensures that no value is present for SitEstimatedEntryDate, not even an explicit nil
### GetSitEstimatedDepartureDate

`func (o *CreatePPMShipment) GetSitEstimatedDepartureDate() string`

GetSitEstimatedDepartureDate returns the SitEstimatedDepartureDate field if non-nil, zero value otherwise.

### GetSitEstimatedDepartureDateOk

`func (o *CreatePPMShipment) GetSitEstimatedDepartureDateOk() (*string, bool)`

GetSitEstimatedDepartureDateOk returns a tuple with the SitEstimatedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedDepartureDate

`func (o *CreatePPMShipment) SetSitEstimatedDepartureDate(v string)`

SetSitEstimatedDepartureDate sets SitEstimatedDepartureDate field to given value.

### HasSitEstimatedDepartureDate

`func (o *CreatePPMShipment) HasSitEstimatedDepartureDate() bool`

HasSitEstimatedDepartureDate returns a boolean if a field has been set.

### SetSitEstimatedDepartureDateNil

`func (o *CreatePPMShipment) SetSitEstimatedDepartureDateNil(b bool)`

 SetSitEstimatedDepartureDateNil sets the value for SitEstimatedDepartureDate to be an explicit nil

### UnsetSitEstimatedDepartureDate
`func (o *CreatePPMShipment) UnsetSitEstimatedDepartureDate()`

UnsetSitEstimatedDepartureDate ensures that no value is present for SitEstimatedDepartureDate, not even an explicit nil
### GetEstimatedWeight

`func (o *CreatePPMShipment) GetEstimatedWeight() int32`

GetEstimatedWeight returns the EstimatedWeight field if non-nil, zero value otherwise.

### GetEstimatedWeightOk

`func (o *CreatePPMShipment) GetEstimatedWeightOk() (*int32, bool)`

GetEstimatedWeightOk returns a tuple with the EstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedWeight

`func (o *CreatePPMShipment) SetEstimatedWeight(v int32)`

SetEstimatedWeight sets EstimatedWeight field to given value.


### GetHasProGear

`func (o *CreatePPMShipment) GetHasProGear() bool`

GetHasProGear returns the HasProGear field if non-nil, zero value otherwise.

### GetHasProGearOk

`func (o *CreatePPMShipment) GetHasProGearOk() (*bool, bool)`

GetHasProGearOk returns a tuple with the HasProGear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProGear

`func (o *CreatePPMShipment) SetHasProGear(v bool)`

SetHasProGear sets HasProGear field to given value.


### GetProGearWeight

`func (o *CreatePPMShipment) GetProGearWeight() int32`

GetProGearWeight returns the ProGearWeight field if non-nil, zero value otherwise.

### GetProGearWeightOk

`func (o *CreatePPMShipment) GetProGearWeightOk() (*int32, bool)`

GetProGearWeightOk returns a tuple with the ProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProGearWeight

`func (o *CreatePPMShipment) SetProGearWeight(v int32)`

SetProGearWeight sets ProGearWeight field to given value.

### HasProGearWeight

`func (o *CreatePPMShipment) HasProGearWeight() bool`

HasProGearWeight returns a boolean if a field has been set.

### SetProGearWeightNil

`func (o *CreatePPMShipment) SetProGearWeightNil(b bool)`

 SetProGearWeightNil sets the value for ProGearWeight to be an explicit nil

### UnsetProGearWeight
`func (o *CreatePPMShipment) UnsetProGearWeight()`

UnsetProGearWeight ensures that no value is present for ProGearWeight, not even an explicit nil
### GetSpouseProGearWeight

`func (o *CreatePPMShipment) GetSpouseProGearWeight() int32`

GetSpouseProGearWeight returns the SpouseProGearWeight field if non-nil, zero value otherwise.

### GetSpouseProGearWeightOk

`func (o *CreatePPMShipment) GetSpouseProGearWeightOk() (*int32, bool)`

GetSpouseProGearWeightOk returns a tuple with the SpouseProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpouseProGearWeight

`func (o *CreatePPMShipment) SetSpouseProGearWeight(v int32)`

SetSpouseProGearWeight sets SpouseProGearWeight field to given value.

### HasSpouseProGearWeight

`func (o *CreatePPMShipment) HasSpouseProGearWeight() bool`

HasSpouseProGearWeight returns a boolean if a field has been set.

### SetSpouseProGearWeightNil

`func (o *CreatePPMShipment) SetSpouseProGearWeightNil(b bool)`

 SetSpouseProGearWeightNil sets the value for SpouseProGearWeight to be an explicit nil

### UnsetSpouseProGearWeight
`func (o *CreatePPMShipment) UnsetSpouseProGearWeight()`

UnsetSpouseProGearWeight ensures that no value is present for SpouseProGearWeight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


