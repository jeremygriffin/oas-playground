# UpdatePPMShipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedDepartureDate** | Pointer to **NullableString** | Date the customer expects to begin moving from their origin.  | [optional] 
**SitExpected** | Pointer to **NullableBool** | Captures whether some or all of the PPM shipment will require temporary storage at the origin or destination.  Must be set to &#x60;true&#x60; when providing &#x60;sitLocation&#x60;, &#x60;sitEstimatedWeight&#x60;, &#x60;sitEstimatedEntryDate&#x60;, and &#x60;sitEstimatedDepartureDate&#x60; values to calculate the &#x60;sitEstimatedCost&#x60;.  | [optional] 
**SitLocation** | Pointer to [**SITLocationType**](SITLocationType.md) |  | [optional] 
**SitEstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of the goods being put into storage. | [optional] 
**SitEstimatedEntryDate** | Pointer to **NullableString** | The date that goods will first enter the storage location. | [optional] 
**SitEstimatedDepartureDate** | Pointer to **NullableString** | The date that goods will exit the storage location. | [optional] 
**EstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of the PPM shipment goods being moved. | [optional] 
**HasProGear** | Pointer to **NullableBool** | Indicates whether PPM shipment has pro gear for themselves or their spouse.  | [optional] 
**ProGearWeight** | Pointer to **NullableInt32** | The estimated weight of the pro-gear being moved belonging to the service member. | [optional] 
**SpouseProGearWeight** | Pointer to **NullableInt32** | The estimated weight of the pro-gear being moved belonging to a spouse. | [optional] 

## Methods

### NewUpdatePPMShipment

`func NewUpdatePPMShipment() *UpdatePPMShipment`

NewUpdatePPMShipment instantiates a new UpdatePPMShipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePPMShipmentWithDefaults

`func NewUpdatePPMShipmentWithDefaults() *UpdatePPMShipment`

NewUpdatePPMShipmentWithDefaults instantiates a new UpdatePPMShipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedDepartureDate

`func (o *UpdatePPMShipment) GetExpectedDepartureDate() string`

GetExpectedDepartureDate returns the ExpectedDepartureDate field if non-nil, zero value otherwise.

### GetExpectedDepartureDateOk

`func (o *UpdatePPMShipment) GetExpectedDepartureDateOk() (*string, bool)`

GetExpectedDepartureDateOk returns a tuple with the ExpectedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDepartureDate

`func (o *UpdatePPMShipment) SetExpectedDepartureDate(v string)`

SetExpectedDepartureDate sets ExpectedDepartureDate field to given value.

### HasExpectedDepartureDate

`func (o *UpdatePPMShipment) HasExpectedDepartureDate() bool`

HasExpectedDepartureDate returns a boolean if a field has been set.

### SetExpectedDepartureDateNil

`func (o *UpdatePPMShipment) SetExpectedDepartureDateNil(b bool)`

 SetExpectedDepartureDateNil sets the value for ExpectedDepartureDate to be an explicit nil

### UnsetExpectedDepartureDate
`func (o *UpdatePPMShipment) UnsetExpectedDepartureDate()`

UnsetExpectedDepartureDate ensures that no value is present for ExpectedDepartureDate, not even an explicit nil
### GetSitExpected

`func (o *UpdatePPMShipment) GetSitExpected() bool`

GetSitExpected returns the SitExpected field if non-nil, zero value otherwise.

### GetSitExpectedOk

`func (o *UpdatePPMShipment) GetSitExpectedOk() (*bool, bool)`

GetSitExpectedOk returns a tuple with the SitExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitExpected

`func (o *UpdatePPMShipment) SetSitExpected(v bool)`

SetSitExpected sets SitExpected field to given value.

### HasSitExpected

`func (o *UpdatePPMShipment) HasSitExpected() bool`

HasSitExpected returns a boolean if a field has been set.

### SetSitExpectedNil

`func (o *UpdatePPMShipment) SetSitExpectedNil(b bool)`

 SetSitExpectedNil sets the value for SitExpected to be an explicit nil

### UnsetSitExpected
`func (o *UpdatePPMShipment) UnsetSitExpected()`

UnsetSitExpected ensures that no value is present for SitExpected, not even an explicit nil
### GetSitLocation

`func (o *UpdatePPMShipment) GetSitLocation() SITLocationType`

GetSitLocation returns the SitLocation field if non-nil, zero value otherwise.

### GetSitLocationOk

`func (o *UpdatePPMShipment) GetSitLocationOk() (*SITLocationType, bool)`

GetSitLocationOk returns a tuple with the SitLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitLocation

`func (o *UpdatePPMShipment) SetSitLocation(v SITLocationType)`

SetSitLocation sets SitLocation field to given value.

### HasSitLocation

`func (o *UpdatePPMShipment) HasSitLocation() bool`

HasSitLocation returns a boolean if a field has been set.

### GetSitEstimatedWeight

`func (o *UpdatePPMShipment) GetSitEstimatedWeight() int32`

GetSitEstimatedWeight returns the SitEstimatedWeight field if non-nil, zero value otherwise.

### GetSitEstimatedWeightOk

`func (o *UpdatePPMShipment) GetSitEstimatedWeightOk() (*int32, bool)`

GetSitEstimatedWeightOk returns a tuple with the SitEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedWeight

`func (o *UpdatePPMShipment) SetSitEstimatedWeight(v int32)`

SetSitEstimatedWeight sets SitEstimatedWeight field to given value.

### HasSitEstimatedWeight

`func (o *UpdatePPMShipment) HasSitEstimatedWeight() bool`

HasSitEstimatedWeight returns a boolean if a field has been set.

### SetSitEstimatedWeightNil

`func (o *UpdatePPMShipment) SetSitEstimatedWeightNil(b bool)`

 SetSitEstimatedWeightNil sets the value for SitEstimatedWeight to be an explicit nil

### UnsetSitEstimatedWeight
`func (o *UpdatePPMShipment) UnsetSitEstimatedWeight()`

UnsetSitEstimatedWeight ensures that no value is present for SitEstimatedWeight, not even an explicit nil
### GetSitEstimatedEntryDate

`func (o *UpdatePPMShipment) GetSitEstimatedEntryDate() string`

GetSitEstimatedEntryDate returns the SitEstimatedEntryDate field if non-nil, zero value otherwise.

### GetSitEstimatedEntryDateOk

`func (o *UpdatePPMShipment) GetSitEstimatedEntryDateOk() (*string, bool)`

GetSitEstimatedEntryDateOk returns a tuple with the SitEstimatedEntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedEntryDate

`func (o *UpdatePPMShipment) SetSitEstimatedEntryDate(v string)`

SetSitEstimatedEntryDate sets SitEstimatedEntryDate field to given value.

### HasSitEstimatedEntryDate

`func (o *UpdatePPMShipment) HasSitEstimatedEntryDate() bool`

HasSitEstimatedEntryDate returns a boolean if a field has been set.

### SetSitEstimatedEntryDateNil

`func (o *UpdatePPMShipment) SetSitEstimatedEntryDateNil(b bool)`

 SetSitEstimatedEntryDateNil sets the value for SitEstimatedEntryDate to be an explicit nil

### UnsetSitEstimatedEntryDate
`func (o *UpdatePPMShipment) UnsetSitEstimatedEntryDate()`

UnsetSitEstimatedEntryDate ensures that no value is present for SitEstimatedEntryDate, not even an explicit nil
### GetSitEstimatedDepartureDate

`func (o *UpdatePPMShipment) GetSitEstimatedDepartureDate() string`

GetSitEstimatedDepartureDate returns the SitEstimatedDepartureDate field if non-nil, zero value otherwise.

### GetSitEstimatedDepartureDateOk

`func (o *UpdatePPMShipment) GetSitEstimatedDepartureDateOk() (*string, bool)`

GetSitEstimatedDepartureDateOk returns a tuple with the SitEstimatedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedDepartureDate

`func (o *UpdatePPMShipment) SetSitEstimatedDepartureDate(v string)`

SetSitEstimatedDepartureDate sets SitEstimatedDepartureDate field to given value.

### HasSitEstimatedDepartureDate

`func (o *UpdatePPMShipment) HasSitEstimatedDepartureDate() bool`

HasSitEstimatedDepartureDate returns a boolean if a field has been set.

### SetSitEstimatedDepartureDateNil

`func (o *UpdatePPMShipment) SetSitEstimatedDepartureDateNil(b bool)`

 SetSitEstimatedDepartureDateNil sets the value for SitEstimatedDepartureDate to be an explicit nil

### UnsetSitEstimatedDepartureDate
`func (o *UpdatePPMShipment) UnsetSitEstimatedDepartureDate()`

UnsetSitEstimatedDepartureDate ensures that no value is present for SitEstimatedDepartureDate, not even an explicit nil
### GetEstimatedWeight

`func (o *UpdatePPMShipment) GetEstimatedWeight() int32`

GetEstimatedWeight returns the EstimatedWeight field if non-nil, zero value otherwise.

### GetEstimatedWeightOk

`func (o *UpdatePPMShipment) GetEstimatedWeightOk() (*int32, bool)`

GetEstimatedWeightOk returns a tuple with the EstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedWeight

`func (o *UpdatePPMShipment) SetEstimatedWeight(v int32)`

SetEstimatedWeight sets EstimatedWeight field to given value.

### HasEstimatedWeight

`func (o *UpdatePPMShipment) HasEstimatedWeight() bool`

HasEstimatedWeight returns a boolean if a field has been set.

### SetEstimatedWeightNil

`func (o *UpdatePPMShipment) SetEstimatedWeightNil(b bool)`

 SetEstimatedWeightNil sets the value for EstimatedWeight to be an explicit nil

### UnsetEstimatedWeight
`func (o *UpdatePPMShipment) UnsetEstimatedWeight()`

UnsetEstimatedWeight ensures that no value is present for EstimatedWeight, not even an explicit nil
### GetHasProGear

`func (o *UpdatePPMShipment) GetHasProGear() bool`

GetHasProGear returns the HasProGear field if non-nil, zero value otherwise.

### GetHasProGearOk

`func (o *UpdatePPMShipment) GetHasProGearOk() (*bool, bool)`

GetHasProGearOk returns a tuple with the HasProGear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProGear

`func (o *UpdatePPMShipment) SetHasProGear(v bool)`

SetHasProGear sets HasProGear field to given value.

### HasHasProGear

`func (o *UpdatePPMShipment) HasHasProGear() bool`

HasHasProGear returns a boolean if a field has been set.

### SetHasProGearNil

`func (o *UpdatePPMShipment) SetHasProGearNil(b bool)`

 SetHasProGearNil sets the value for HasProGear to be an explicit nil

### UnsetHasProGear
`func (o *UpdatePPMShipment) UnsetHasProGear()`

UnsetHasProGear ensures that no value is present for HasProGear, not even an explicit nil
### GetProGearWeight

`func (o *UpdatePPMShipment) GetProGearWeight() int32`

GetProGearWeight returns the ProGearWeight field if non-nil, zero value otherwise.

### GetProGearWeightOk

`func (o *UpdatePPMShipment) GetProGearWeightOk() (*int32, bool)`

GetProGearWeightOk returns a tuple with the ProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProGearWeight

`func (o *UpdatePPMShipment) SetProGearWeight(v int32)`

SetProGearWeight sets ProGearWeight field to given value.

### HasProGearWeight

`func (o *UpdatePPMShipment) HasProGearWeight() bool`

HasProGearWeight returns a boolean if a field has been set.

### SetProGearWeightNil

`func (o *UpdatePPMShipment) SetProGearWeightNil(b bool)`

 SetProGearWeightNil sets the value for ProGearWeight to be an explicit nil

### UnsetProGearWeight
`func (o *UpdatePPMShipment) UnsetProGearWeight()`

UnsetProGearWeight ensures that no value is present for ProGearWeight, not even an explicit nil
### GetSpouseProGearWeight

`func (o *UpdatePPMShipment) GetSpouseProGearWeight() int32`

GetSpouseProGearWeight returns the SpouseProGearWeight field if non-nil, zero value otherwise.

### GetSpouseProGearWeightOk

`func (o *UpdatePPMShipment) GetSpouseProGearWeightOk() (*int32, bool)`

GetSpouseProGearWeightOk returns a tuple with the SpouseProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpouseProGearWeight

`func (o *UpdatePPMShipment) SetSpouseProGearWeight(v int32)`

SetSpouseProGearWeight sets SpouseProGearWeight field to given value.

### HasSpouseProGearWeight

`func (o *UpdatePPMShipment) HasSpouseProGearWeight() bool`

HasSpouseProGearWeight returns a boolean if a field has been set.

### SetSpouseProGearWeightNil

`func (o *UpdatePPMShipment) SetSpouseProGearWeightNil(b bool)`

 SetSpouseProGearWeightNil sets the value for SpouseProGearWeight to be an explicit nil

### UnsetSpouseProGearWeight
`func (o *UpdatePPMShipment) UnsetSpouseProGearWeight()`

UnsetSpouseProGearWeight ensures that no value is present for SpouseProGearWeight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


