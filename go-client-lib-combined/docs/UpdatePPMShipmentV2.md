# UpdatePPMShipmentV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpectedDepartureDate** | Pointer to **NullableString** | Date the customer expects to begin moving from their origin.  | [optional] 
**SitExpected** | Pointer to **NullableBool** | Captures whether some or all of the PPM shipment will require temporary storage at the origin or destination.  Must be set to &#x60;true&#x60; when providing &#x60;sitLocation&#x60;, &#x60;sitEstimatedWeight&#x60;, &#x60;sitEstimatedEntryDate&#x60;, and &#x60;sitEstimatedDepartureDate&#x60; values to calculate the &#x60;sitEstimatedCost&#x60;.  | [optional] 
**SitLocation** | Pointer to [**SITLocationTypeV2**](SITLocationType.md) |  | [optional] 
**SitEstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of the goods being put into storage. | [optional] 
**SitEstimatedEntryDate** | Pointer to **NullableString** | The date that goods will first enter the storage location. | [optional] 
**SitEstimatedDepartureDate** | Pointer to **NullableString** | The date that goods will exit the storage location. | [optional] 
**EstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of the PPM shipment goods being moved. | [optional] 
**HasProGear** | Pointer to **NullableBool** | Indicates whether PPM shipment has pro gear for themselves or their spouse.  | [optional] 
**ProGearWeight** | Pointer to **NullableInt32** | The estimated weight of the pro-gear being moved belonging to the service member. | [optional] 
**SpouseProGearWeight** | Pointer to **NullableInt32** | The estimated weight of the pro-gear being moved belonging to a spouse. | [optional] 

## Methods

### NewUpdatePPMShipmentV2

`func NewUpdatePPMShipmentV2() *UpdatePPMShipmentV2`

NewUpdatePPMShipmentV2 instantiates a new UpdatePPMShipmentV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePPMShipmentV2WithDefaults

`func NewUpdatePPMShipmentV2WithDefaults() *UpdatePPMShipmentV2`

NewUpdatePPMShipmentV2WithDefaults instantiates a new UpdatePPMShipmentV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpectedDepartureDate

`func (o *UpdatePPMShipmentV2) GetExpectedDepartureDate() string`

GetExpectedDepartureDate returns the ExpectedDepartureDate field if non-nil, zero value otherwise.

### GetExpectedDepartureDateOk

`func (o *UpdatePPMShipmentV2) GetExpectedDepartureDateOk() (*string, bool)`

GetExpectedDepartureDateOk returns a tuple with the ExpectedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDepartureDate

`func (o *UpdatePPMShipmentV2) SetExpectedDepartureDate(v string)`

SetExpectedDepartureDate sets ExpectedDepartureDate field to given value.

### HasExpectedDepartureDate

`func (o *UpdatePPMShipmentV2) HasExpectedDepartureDate() bool`

HasExpectedDepartureDate returns a boolean if a field has been set.

### SetExpectedDepartureDateNil

`func (o *UpdatePPMShipmentV2) SetExpectedDepartureDateNil(b bool)`

 SetExpectedDepartureDateNil sets the value for ExpectedDepartureDate to be an explicit nil

### UnsetExpectedDepartureDate
`func (o *UpdatePPMShipmentV2) UnsetExpectedDepartureDate()`

UnsetExpectedDepartureDate ensures that no value is present for ExpectedDepartureDate, not even an explicit nil
### GetSitExpected

`func (o *UpdatePPMShipmentV2) GetSitExpected() bool`

GetSitExpected returns the SitExpected field if non-nil, zero value otherwise.

### GetSitExpectedOk

`func (o *UpdatePPMShipmentV2) GetSitExpectedOk() (*bool, bool)`

GetSitExpectedOk returns a tuple with the SitExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitExpected

`func (o *UpdatePPMShipmentV2) SetSitExpected(v bool)`

SetSitExpected sets SitExpected field to given value.

### HasSitExpected

`func (o *UpdatePPMShipmentV2) HasSitExpected() bool`

HasSitExpected returns a boolean if a field has been set.

### SetSitExpectedNil

`func (o *UpdatePPMShipmentV2) SetSitExpectedNil(b bool)`

 SetSitExpectedNil sets the value for SitExpected to be an explicit nil

### UnsetSitExpected
`func (o *UpdatePPMShipmentV2) UnsetSitExpected()`

UnsetSitExpected ensures that no value is present for SitExpected, not even an explicit nil
### GetSitLocation

`func (o *UpdatePPMShipmentV2) GetSitLocation() SITLocationTypeV2`

GetSitLocation returns the SitLocation field if non-nil, zero value otherwise.

### GetSitLocationOk

`func (o *UpdatePPMShipmentV2) GetSitLocationOk() (*SITLocationTypeV2, bool)`

GetSitLocationOk returns a tuple with the SitLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitLocation

`func (o *UpdatePPMShipmentV2) SetSitLocation(v SITLocationTypeV2)`

SetSitLocation sets SitLocation field to given value.

### HasSitLocation

`func (o *UpdatePPMShipmentV2) HasSitLocation() bool`

HasSitLocation returns a boolean if a field has been set.

### GetSitEstimatedWeight

`func (o *UpdatePPMShipmentV2) GetSitEstimatedWeight() int32`

GetSitEstimatedWeight returns the SitEstimatedWeight field if non-nil, zero value otherwise.

### GetSitEstimatedWeightOk

`func (o *UpdatePPMShipmentV2) GetSitEstimatedWeightOk() (*int32, bool)`

GetSitEstimatedWeightOk returns a tuple with the SitEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedWeight

`func (o *UpdatePPMShipmentV2) SetSitEstimatedWeight(v int32)`

SetSitEstimatedWeight sets SitEstimatedWeight field to given value.

### HasSitEstimatedWeight

`func (o *UpdatePPMShipmentV2) HasSitEstimatedWeight() bool`

HasSitEstimatedWeight returns a boolean if a field has been set.

### SetSitEstimatedWeightNil

`func (o *UpdatePPMShipmentV2) SetSitEstimatedWeightNil(b bool)`

 SetSitEstimatedWeightNil sets the value for SitEstimatedWeight to be an explicit nil

### UnsetSitEstimatedWeight
`func (o *UpdatePPMShipmentV2) UnsetSitEstimatedWeight()`

UnsetSitEstimatedWeight ensures that no value is present for SitEstimatedWeight, not even an explicit nil
### GetSitEstimatedEntryDate

`func (o *UpdatePPMShipmentV2) GetSitEstimatedEntryDate() string`

GetSitEstimatedEntryDate returns the SitEstimatedEntryDate field if non-nil, zero value otherwise.

### GetSitEstimatedEntryDateOk

`func (o *UpdatePPMShipmentV2) GetSitEstimatedEntryDateOk() (*string, bool)`

GetSitEstimatedEntryDateOk returns a tuple with the SitEstimatedEntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedEntryDate

`func (o *UpdatePPMShipmentV2) SetSitEstimatedEntryDate(v string)`

SetSitEstimatedEntryDate sets SitEstimatedEntryDate field to given value.

### HasSitEstimatedEntryDate

`func (o *UpdatePPMShipmentV2) HasSitEstimatedEntryDate() bool`

HasSitEstimatedEntryDate returns a boolean if a field has been set.

### SetSitEstimatedEntryDateNil

`func (o *UpdatePPMShipmentV2) SetSitEstimatedEntryDateNil(b bool)`

 SetSitEstimatedEntryDateNil sets the value for SitEstimatedEntryDate to be an explicit nil

### UnsetSitEstimatedEntryDate
`func (o *UpdatePPMShipmentV2) UnsetSitEstimatedEntryDate()`

UnsetSitEstimatedEntryDate ensures that no value is present for SitEstimatedEntryDate, not even an explicit nil
### GetSitEstimatedDepartureDate

`func (o *UpdatePPMShipmentV2) GetSitEstimatedDepartureDate() string`

GetSitEstimatedDepartureDate returns the SitEstimatedDepartureDate field if non-nil, zero value otherwise.

### GetSitEstimatedDepartureDateOk

`func (o *UpdatePPMShipmentV2) GetSitEstimatedDepartureDateOk() (*string, bool)`

GetSitEstimatedDepartureDateOk returns a tuple with the SitEstimatedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedDepartureDate

`func (o *UpdatePPMShipmentV2) SetSitEstimatedDepartureDate(v string)`

SetSitEstimatedDepartureDate sets SitEstimatedDepartureDate field to given value.

### HasSitEstimatedDepartureDate

`func (o *UpdatePPMShipmentV2) HasSitEstimatedDepartureDate() bool`

HasSitEstimatedDepartureDate returns a boolean if a field has been set.

### SetSitEstimatedDepartureDateNil

`func (o *UpdatePPMShipmentV2) SetSitEstimatedDepartureDateNil(b bool)`

 SetSitEstimatedDepartureDateNil sets the value for SitEstimatedDepartureDate to be an explicit nil

### UnsetSitEstimatedDepartureDate
`func (o *UpdatePPMShipmentV2) UnsetSitEstimatedDepartureDate()`

UnsetSitEstimatedDepartureDate ensures that no value is present for SitEstimatedDepartureDate, not even an explicit nil
### GetEstimatedWeight

`func (o *UpdatePPMShipmentV2) GetEstimatedWeight() int32`

GetEstimatedWeight returns the EstimatedWeight field if non-nil, zero value otherwise.

### GetEstimatedWeightOk

`func (o *UpdatePPMShipmentV2) GetEstimatedWeightOk() (*int32, bool)`

GetEstimatedWeightOk returns a tuple with the EstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedWeight

`func (o *UpdatePPMShipmentV2) SetEstimatedWeight(v int32)`

SetEstimatedWeight sets EstimatedWeight field to given value.

### HasEstimatedWeight

`func (o *UpdatePPMShipmentV2) HasEstimatedWeight() bool`

HasEstimatedWeight returns a boolean if a field has been set.

### SetEstimatedWeightNil

`func (o *UpdatePPMShipmentV2) SetEstimatedWeightNil(b bool)`

 SetEstimatedWeightNil sets the value for EstimatedWeight to be an explicit nil

### UnsetEstimatedWeight
`func (o *UpdatePPMShipmentV2) UnsetEstimatedWeight()`

UnsetEstimatedWeight ensures that no value is present for EstimatedWeight, not even an explicit nil
### GetHasProGear

`func (o *UpdatePPMShipmentV2) GetHasProGear() bool`

GetHasProGear returns the HasProGear field if non-nil, zero value otherwise.

### GetHasProGearOk

`func (o *UpdatePPMShipmentV2) GetHasProGearOk() (*bool, bool)`

GetHasProGearOk returns a tuple with the HasProGear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProGear

`func (o *UpdatePPMShipmentV2) SetHasProGear(v bool)`

SetHasProGear sets HasProGear field to given value.

### HasHasProGear

`func (o *UpdatePPMShipmentV2) HasHasProGear() bool`

HasHasProGear returns a boolean if a field has been set.

### SetHasProGearNil

`func (o *UpdatePPMShipmentV2) SetHasProGearNil(b bool)`

 SetHasProGearNil sets the value for HasProGear to be an explicit nil

### UnsetHasProGear
`func (o *UpdatePPMShipmentV2) UnsetHasProGear()`

UnsetHasProGear ensures that no value is present for HasProGear, not even an explicit nil
### GetProGearWeight

`func (o *UpdatePPMShipmentV2) GetProGearWeight() int32`

GetProGearWeight returns the ProGearWeight field if non-nil, zero value otherwise.

### GetProGearWeightOk

`func (o *UpdatePPMShipmentV2) GetProGearWeightOk() (*int32, bool)`

GetProGearWeightOk returns a tuple with the ProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProGearWeight

`func (o *UpdatePPMShipmentV2) SetProGearWeight(v int32)`

SetProGearWeight sets ProGearWeight field to given value.

### HasProGearWeight

`func (o *UpdatePPMShipmentV2) HasProGearWeight() bool`

HasProGearWeight returns a boolean if a field has been set.

### SetProGearWeightNil

`func (o *UpdatePPMShipmentV2) SetProGearWeightNil(b bool)`

 SetProGearWeightNil sets the value for ProGearWeight to be an explicit nil

### UnsetProGearWeight
`func (o *UpdatePPMShipmentV2) UnsetProGearWeight()`

UnsetProGearWeight ensures that no value is present for ProGearWeight, not even an explicit nil
### GetSpouseProGearWeight

`func (o *UpdatePPMShipmentV2) GetSpouseProGearWeight() int32`

GetSpouseProGearWeight returns the SpouseProGearWeight field if non-nil, zero value otherwise.

### GetSpouseProGearWeightOk

`func (o *UpdatePPMShipmentV2) GetSpouseProGearWeightOk() (*int32, bool)`

GetSpouseProGearWeightOk returns a tuple with the SpouseProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpouseProGearWeight

`func (o *UpdatePPMShipmentV2) SetSpouseProGearWeight(v int32)`

SetSpouseProGearWeight sets SpouseProGearWeight field to given value.

### HasSpouseProGearWeight

`func (o *UpdatePPMShipmentV2) HasSpouseProGearWeight() bool`

HasSpouseProGearWeight returns a boolean if a field has been set.

### SetSpouseProGearWeightNil

`func (o *UpdatePPMShipmentV2) SetSpouseProGearWeightNil(b bool)`

 SetSpouseProGearWeightNil sets the value for SpouseProGearWeight to be an explicit nil

### UnsetSpouseProGearWeight
`func (o *UpdatePPMShipmentV2) UnsetSpouseProGearWeight()`

UnsetSpouseProGearWeight ensures that no value is present for SpouseProGearWeight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


