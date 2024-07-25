# PPMShipment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The primary unique identifier of this PPM shipment | [readonly] 
**ShipmentId** | **string** | The id of the parent MTOShipment record | [readonly] 
**CreatedAt** | **time.Time** | The timestamp of when the PPM shipment was created (UTC) | [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The timestamp of when a property of this object was last updated (UTC) | [optional] [readonly] 
**Status** | [**PPMShipmentStatus**](PPMShipmentStatus.md) |  | 
**ExpectedDepartureDate** | **string** | Date the customer expects to begin moving from their origin.  | 
**ActualMoveDate** | Pointer to **NullableString** | The actual start date of when the PPM shipment left the origin. | [optional] 
**SubmittedAt** | Pointer to **NullableTime** | The timestamp of when the customer submitted their PPM documentation to the counselor for review. | [optional] 
**ReviewedAt** | Pointer to **NullableTime** | The timestamp of when the Service Counselor has reviewed all of the closeout documents. | [optional] 
**ApprovedAt** | Pointer to **NullableTime** | The timestamp of when the shipment was approved and the service member can begin their move. | [optional] 
**ActualPickupPostalCode** | Pointer to **NullableString** | The actual postal code where the PPM shipment started. To be filled once the customer has moved the shipment.  | [optional] 
**ActualDestinationPostalCode** | Pointer to **NullableString** | The actual postal code where the PPM shipment ended. To be filled once the customer has moved the shipment.  | [optional] 
**SitExpected** | **bool** | Captures whether some or all of the PPM shipment will require temporary storage at the origin or destination.  Must be set to &#x60;true&#x60; when providing &#x60;sitLocation&#x60;, &#x60;sitEstimatedWeight&#x60;, &#x60;sitEstimatedEntryDate&#x60;, and &#x60;sitEstimatedDepartureDate&#x60; values to calculate the &#x60;sitEstimatedCost&#x60;.  | 
**EstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of the PPM shipment goods being moved in pounds. | [optional] 
**HasProGear** | Pointer to **NullableBool** | Indicates whether PPM shipment has pro gear for themselves or their spouse.  | [optional] 
**ProGearWeight** | Pointer to **NullableInt32** | The estimated weight of the pro-gear being moved belonging to the service member in pounds. | [optional] 
**SpouseProGearWeight** | Pointer to **NullableInt32** | The estimated weight of the pro-gear being moved belonging to a spouse in pounds. | [optional] 
**EstimatedIncentive** | Pointer to **NullableInt32** | The estimated amount the government will pay the service member to move their belongings based on the moving date, locations, and shipment weight. | [optional] 
**HasRequestedAdvance** | Pointer to **NullableBool** | Indicates whether an advance has been requested for the PPM shipment.  | [optional] 
**AdvanceAmountRequested** | Pointer to **NullableInt32** | The amount requested as an advance by the service member, up to a maximum percentage of the estimated incentive.  | [optional] 
**HasReceivedAdvance** | Pointer to **NullableBool** | Indicates whether an advance was received for the PPM shipment.  | [optional] 
**AdvanceAmountReceived** | Pointer to **NullableInt32** | The amount received for an advance, or null if no advance is received.  | [optional] 
**SitLocation** | Pointer to [**SITLocationType**](SITLocationType.md) |  | [optional] 
**SitEstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of the goods being put into storage in pounds. | [optional] 
**SitEstimatedEntryDate** | Pointer to **NullableString** | The date that goods will first enter the storage location. | [optional] 
**SitEstimatedDepartureDate** | Pointer to **NullableString** | The date that goods will exit the storage location. | [optional] 
**SitEstimatedCost** | Pointer to **NullableInt32** | The estimated amount that the government will pay the service member to put their goods into storage. This estimated storage cost is separate from the estimated incentive. | [optional] 
**ETag** | **string** | A hash unique to this shipment that should be used as the \&quot;If-Match\&quot; header for any updates. | [readonly] 

## Methods

### NewPPMShipment

`func NewPPMShipment(id string, shipmentId string, createdAt time.Time, status PPMShipmentStatus, expectedDepartureDate string, sitExpected bool, eTag string, ) *PPMShipment`

NewPPMShipment instantiates a new PPMShipment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPPMShipmentWithDefaults

`func NewPPMShipmentWithDefaults() *PPMShipment`

NewPPMShipmentWithDefaults instantiates a new PPMShipment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PPMShipment) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PPMShipment) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PPMShipment) SetId(v string)`

SetId sets Id field to given value.


### GetShipmentId

`func (o *PPMShipment) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *PPMShipment) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *PPMShipment) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.


### GetCreatedAt

`func (o *PPMShipment) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PPMShipment) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PPMShipment) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PPMShipment) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PPMShipment) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PPMShipment) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PPMShipment) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *PPMShipment) GetStatus() PPMShipmentStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PPMShipment) GetStatusOk() (*PPMShipmentStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PPMShipment) SetStatus(v PPMShipmentStatus)`

SetStatus sets Status field to given value.


### GetExpectedDepartureDate

`func (o *PPMShipment) GetExpectedDepartureDate() string`

GetExpectedDepartureDate returns the ExpectedDepartureDate field if non-nil, zero value otherwise.

### GetExpectedDepartureDateOk

`func (o *PPMShipment) GetExpectedDepartureDateOk() (*string, bool)`

GetExpectedDepartureDateOk returns a tuple with the ExpectedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDepartureDate

`func (o *PPMShipment) SetExpectedDepartureDate(v string)`

SetExpectedDepartureDate sets ExpectedDepartureDate field to given value.


### GetActualMoveDate

`func (o *PPMShipment) GetActualMoveDate() string`

GetActualMoveDate returns the ActualMoveDate field if non-nil, zero value otherwise.

### GetActualMoveDateOk

`func (o *PPMShipment) GetActualMoveDateOk() (*string, bool)`

GetActualMoveDateOk returns a tuple with the ActualMoveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualMoveDate

`func (o *PPMShipment) SetActualMoveDate(v string)`

SetActualMoveDate sets ActualMoveDate field to given value.

### HasActualMoveDate

`func (o *PPMShipment) HasActualMoveDate() bool`

HasActualMoveDate returns a boolean if a field has been set.

### SetActualMoveDateNil

`func (o *PPMShipment) SetActualMoveDateNil(b bool)`

 SetActualMoveDateNil sets the value for ActualMoveDate to be an explicit nil

### UnsetActualMoveDate
`func (o *PPMShipment) UnsetActualMoveDate()`

UnsetActualMoveDate ensures that no value is present for ActualMoveDate, not even an explicit nil
### GetSubmittedAt

`func (o *PPMShipment) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *PPMShipment) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *PPMShipment) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *PPMShipment) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### SetSubmittedAtNil

`func (o *PPMShipment) SetSubmittedAtNil(b bool)`

 SetSubmittedAtNil sets the value for SubmittedAt to be an explicit nil

### UnsetSubmittedAt
`func (o *PPMShipment) UnsetSubmittedAt()`

UnsetSubmittedAt ensures that no value is present for SubmittedAt, not even an explicit nil
### GetReviewedAt

`func (o *PPMShipment) GetReviewedAt() time.Time`

GetReviewedAt returns the ReviewedAt field if non-nil, zero value otherwise.

### GetReviewedAtOk

`func (o *PPMShipment) GetReviewedAtOk() (*time.Time, bool)`

GetReviewedAtOk returns a tuple with the ReviewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedAt

`func (o *PPMShipment) SetReviewedAt(v time.Time)`

SetReviewedAt sets ReviewedAt field to given value.

### HasReviewedAt

`func (o *PPMShipment) HasReviewedAt() bool`

HasReviewedAt returns a boolean if a field has been set.

### SetReviewedAtNil

`func (o *PPMShipment) SetReviewedAtNil(b bool)`

 SetReviewedAtNil sets the value for ReviewedAt to be an explicit nil

### UnsetReviewedAt
`func (o *PPMShipment) UnsetReviewedAt()`

UnsetReviewedAt ensures that no value is present for ReviewedAt, not even an explicit nil
### GetApprovedAt

`func (o *PPMShipment) GetApprovedAt() time.Time`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *PPMShipment) GetApprovedAtOk() (*time.Time, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *PPMShipment) SetApprovedAt(v time.Time)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *PPMShipment) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### SetApprovedAtNil

`func (o *PPMShipment) SetApprovedAtNil(b bool)`

 SetApprovedAtNil sets the value for ApprovedAt to be an explicit nil

### UnsetApprovedAt
`func (o *PPMShipment) UnsetApprovedAt()`

UnsetApprovedAt ensures that no value is present for ApprovedAt, not even an explicit nil
### GetActualPickupPostalCode

`func (o *PPMShipment) GetActualPickupPostalCode() string`

GetActualPickupPostalCode returns the ActualPickupPostalCode field if non-nil, zero value otherwise.

### GetActualPickupPostalCodeOk

`func (o *PPMShipment) GetActualPickupPostalCodeOk() (*string, bool)`

GetActualPickupPostalCodeOk returns a tuple with the ActualPickupPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPickupPostalCode

`func (o *PPMShipment) SetActualPickupPostalCode(v string)`

SetActualPickupPostalCode sets ActualPickupPostalCode field to given value.

### HasActualPickupPostalCode

`func (o *PPMShipment) HasActualPickupPostalCode() bool`

HasActualPickupPostalCode returns a boolean if a field has been set.

### SetActualPickupPostalCodeNil

`func (o *PPMShipment) SetActualPickupPostalCodeNil(b bool)`

 SetActualPickupPostalCodeNil sets the value for ActualPickupPostalCode to be an explicit nil

### UnsetActualPickupPostalCode
`func (o *PPMShipment) UnsetActualPickupPostalCode()`

UnsetActualPickupPostalCode ensures that no value is present for ActualPickupPostalCode, not even an explicit nil
### GetActualDestinationPostalCode

`func (o *PPMShipment) GetActualDestinationPostalCode() string`

GetActualDestinationPostalCode returns the ActualDestinationPostalCode field if non-nil, zero value otherwise.

### GetActualDestinationPostalCodeOk

`func (o *PPMShipment) GetActualDestinationPostalCodeOk() (*string, bool)`

GetActualDestinationPostalCodeOk returns a tuple with the ActualDestinationPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDestinationPostalCode

`func (o *PPMShipment) SetActualDestinationPostalCode(v string)`

SetActualDestinationPostalCode sets ActualDestinationPostalCode field to given value.

### HasActualDestinationPostalCode

`func (o *PPMShipment) HasActualDestinationPostalCode() bool`

HasActualDestinationPostalCode returns a boolean if a field has been set.

### SetActualDestinationPostalCodeNil

`func (o *PPMShipment) SetActualDestinationPostalCodeNil(b bool)`

 SetActualDestinationPostalCodeNil sets the value for ActualDestinationPostalCode to be an explicit nil

### UnsetActualDestinationPostalCode
`func (o *PPMShipment) UnsetActualDestinationPostalCode()`

UnsetActualDestinationPostalCode ensures that no value is present for ActualDestinationPostalCode, not even an explicit nil
### GetSitExpected

`func (o *PPMShipment) GetSitExpected() bool`

GetSitExpected returns the SitExpected field if non-nil, zero value otherwise.

### GetSitExpectedOk

`func (o *PPMShipment) GetSitExpectedOk() (*bool, bool)`

GetSitExpectedOk returns a tuple with the SitExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitExpected

`func (o *PPMShipment) SetSitExpected(v bool)`

SetSitExpected sets SitExpected field to given value.


### GetEstimatedWeight

`func (o *PPMShipment) GetEstimatedWeight() int32`

GetEstimatedWeight returns the EstimatedWeight field if non-nil, zero value otherwise.

### GetEstimatedWeightOk

`func (o *PPMShipment) GetEstimatedWeightOk() (*int32, bool)`

GetEstimatedWeightOk returns a tuple with the EstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedWeight

`func (o *PPMShipment) SetEstimatedWeight(v int32)`

SetEstimatedWeight sets EstimatedWeight field to given value.

### HasEstimatedWeight

`func (o *PPMShipment) HasEstimatedWeight() bool`

HasEstimatedWeight returns a boolean if a field has been set.

### SetEstimatedWeightNil

`func (o *PPMShipment) SetEstimatedWeightNil(b bool)`

 SetEstimatedWeightNil sets the value for EstimatedWeight to be an explicit nil

### UnsetEstimatedWeight
`func (o *PPMShipment) UnsetEstimatedWeight()`

UnsetEstimatedWeight ensures that no value is present for EstimatedWeight, not even an explicit nil
### GetHasProGear

`func (o *PPMShipment) GetHasProGear() bool`

GetHasProGear returns the HasProGear field if non-nil, zero value otherwise.

### GetHasProGearOk

`func (o *PPMShipment) GetHasProGearOk() (*bool, bool)`

GetHasProGearOk returns a tuple with the HasProGear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProGear

`func (o *PPMShipment) SetHasProGear(v bool)`

SetHasProGear sets HasProGear field to given value.

### HasHasProGear

`func (o *PPMShipment) HasHasProGear() bool`

HasHasProGear returns a boolean if a field has been set.

### SetHasProGearNil

`func (o *PPMShipment) SetHasProGearNil(b bool)`

 SetHasProGearNil sets the value for HasProGear to be an explicit nil

### UnsetHasProGear
`func (o *PPMShipment) UnsetHasProGear()`

UnsetHasProGear ensures that no value is present for HasProGear, not even an explicit nil
### GetProGearWeight

`func (o *PPMShipment) GetProGearWeight() int32`

GetProGearWeight returns the ProGearWeight field if non-nil, zero value otherwise.

### GetProGearWeightOk

`func (o *PPMShipment) GetProGearWeightOk() (*int32, bool)`

GetProGearWeightOk returns a tuple with the ProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProGearWeight

`func (o *PPMShipment) SetProGearWeight(v int32)`

SetProGearWeight sets ProGearWeight field to given value.

### HasProGearWeight

`func (o *PPMShipment) HasProGearWeight() bool`

HasProGearWeight returns a boolean if a field has been set.

### SetProGearWeightNil

`func (o *PPMShipment) SetProGearWeightNil(b bool)`

 SetProGearWeightNil sets the value for ProGearWeight to be an explicit nil

### UnsetProGearWeight
`func (o *PPMShipment) UnsetProGearWeight()`

UnsetProGearWeight ensures that no value is present for ProGearWeight, not even an explicit nil
### GetSpouseProGearWeight

`func (o *PPMShipment) GetSpouseProGearWeight() int32`

GetSpouseProGearWeight returns the SpouseProGearWeight field if non-nil, zero value otherwise.

### GetSpouseProGearWeightOk

`func (o *PPMShipment) GetSpouseProGearWeightOk() (*int32, bool)`

GetSpouseProGearWeightOk returns a tuple with the SpouseProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpouseProGearWeight

`func (o *PPMShipment) SetSpouseProGearWeight(v int32)`

SetSpouseProGearWeight sets SpouseProGearWeight field to given value.

### HasSpouseProGearWeight

`func (o *PPMShipment) HasSpouseProGearWeight() bool`

HasSpouseProGearWeight returns a boolean if a field has been set.

### SetSpouseProGearWeightNil

`func (o *PPMShipment) SetSpouseProGearWeightNil(b bool)`

 SetSpouseProGearWeightNil sets the value for SpouseProGearWeight to be an explicit nil

### UnsetSpouseProGearWeight
`func (o *PPMShipment) UnsetSpouseProGearWeight()`

UnsetSpouseProGearWeight ensures that no value is present for SpouseProGearWeight, not even an explicit nil
### GetEstimatedIncentive

`func (o *PPMShipment) GetEstimatedIncentive() int32`

GetEstimatedIncentive returns the EstimatedIncentive field if non-nil, zero value otherwise.

### GetEstimatedIncentiveOk

`func (o *PPMShipment) GetEstimatedIncentiveOk() (*int32, bool)`

GetEstimatedIncentiveOk returns a tuple with the EstimatedIncentive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedIncentive

`func (o *PPMShipment) SetEstimatedIncentive(v int32)`

SetEstimatedIncentive sets EstimatedIncentive field to given value.

### HasEstimatedIncentive

`func (o *PPMShipment) HasEstimatedIncentive() bool`

HasEstimatedIncentive returns a boolean if a field has been set.

### SetEstimatedIncentiveNil

`func (o *PPMShipment) SetEstimatedIncentiveNil(b bool)`

 SetEstimatedIncentiveNil sets the value for EstimatedIncentive to be an explicit nil

### UnsetEstimatedIncentive
`func (o *PPMShipment) UnsetEstimatedIncentive()`

UnsetEstimatedIncentive ensures that no value is present for EstimatedIncentive, not even an explicit nil
### GetHasRequestedAdvance

`func (o *PPMShipment) GetHasRequestedAdvance() bool`

GetHasRequestedAdvance returns the HasRequestedAdvance field if non-nil, zero value otherwise.

### GetHasRequestedAdvanceOk

`func (o *PPMShipment) GetHasRequestedAdvanceOk() (*bool, bool)`

GetHasRequestedAdvanceOk returns a tuple with the HasRequestedAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRequestedAdvance

`func (o *PPMShipment) SetHasRequestedAdvance(v bool)`

SetHasRequestedAdvance sets HasRequestedAdvance field to given value.

### HasHasRequestedAdvance

`func (o *PPMShipment) HasHasRequestedAdvance() bool`

HasHasRequestedAdvance returns a boolean if a field has been set.

### SetHasRequestedAdvanceNil

`func (o *PPMShipment) SetHasRequestedAdvanceNil(b bool)`

 SetHasRequestedAdvanceNil sets the value for HasRequestedAdvance to be an explicit nil

### UnsetHasRequestedAdvance
`func (o *PPMShipment) UnsetHasRequestedAdvance()`

UnsetHasRequestedAdvance ensures that no value is present for HasRequestedAdvance, not even an explicit nil
### GetAdvanceAmountRequested

`func (o *PPMShipment) GetAdvanceAmountRequested() int32`

GetAdvanceAmountRequested returns the AdvanceAmountRequested field if non-nil, zero value otherwise.

### GetAdvanceAmountRequestedOk

`func (o *PPMShipment) GetAdvanceAmountRequestedOk() (*int32, bool)`

GetAdvanceAmountRequestedOk returns a tuple with the AdvanceAmountRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanceAmountRequested

`func (o *PPMShipment) SetAdvanceAmountRequested(v int32)`

SetAdvanceAmountRequested sets AdvanceAmountRequested field to given value.

### HasAdvanceAmountRequested

`func (o *PPMShipment) HasAdvanceAmountRequested() bool`

HasAdvanceAmountRequested returns a boolean if a field has been set.

### SetAdvanceAmountRequestedNil

`func (o *PPMShipment) SetAdvanceAmountRequestedNil(b bool)`

 SetAdvanceAmountRequestedNil sets the value for AdvanceAmountRequested to be an explicit nil

### UnsetAdvanceAmountRequested
`func (o *PPMShipment) UnsetAdvanceAmountRequested()`

UnsetAdvanceAmountRequested ensures that no value is present for AdvanceAmountRequested, not even an explicit nil
### GetHasReceivedAdvance

`func (o *PPMShipment) GetHasReceivedAdvance() bool`

GetHasReceivedAdvance returns the HasReceivedAdvance field if non-nil, zero value otherwise.

### GetHasReceivedAdvanceOk

`func (o *PPMShipment) GetHasReceivedAdvanceOk() (*bool, bool)`

GetHasReceivedAdvanceOk returns a tuple with the HasReceivedAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReceivedAdvance

`func (o *PPMShipment) SetHasReceivedAdvance(v bool)`

SetHasReceivedAdvance sets HasReceivedAdvance field to given value.

### HasHasReceivedAdvance

`func (o *PPMShipment) HasHasReceivedAdvance() bool`

HasHasReceivedAdvance returns a boolean if a field has been set.

### SetHasReceivedAdvanceNil

`func (o *PPMShipment) SetHasReceivedAdvanceNil(b bool)`

 SetHasReceivedAdvanceNil sets the value for HasReceivedAdvance to be an explicit nil

### UnsetHasReceivedAdvance
`func (o *PPMShipment) UnsetHasReceivedAdvance()`

UnsetHasReceivedAdvance ensures that no value is present for HasReceivedAdvance, not even an explicit nil
### GetAdvanceAmountReceived

`func (o *PPMShipment) GetAdvanceAmountReceived() int32`

GetAdvanceAmountReceived returns the AdvanceAmountReceived field if non-nil, zero value otherwise.

### GetAdvanceAmountReceivedOk

`func (o *PPMShipment) GetAdvanceAmountReceivedOk() (*int32, bool)`

GetAdvanceAmountReceivedOk returns a tuple with the AdvanceAmountReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanceAmountReceived

`func (o *PPMShipment) SetAdvanceAmountReceived(v int32)`

SetAdvanceAmountReceived sets AdvanceAmountReceived field to given value.

### HasAdvanceAmountReceived

`func (o *PPMShipment) HasAdvanceAmountReceived() bool`

HasAdvanceAmountReceived returns a boolean if a field has been set.

### SetAdvanceAmountReceivedNil

`func (o *PPMShipment) SetAdvanceAmountReceivedNil(b bool)`

 SetAdvanceAmountReceivedNil sets the value for AdvanceAmountReceived to be an explicit nil

### UnsetAdvanceAmountReceived
`func (o *PPMShipment) UnsetAdvanceAmountReceived()`

UnsetAdvanceAmountReceived ensures that no value is present for AdvanceAmountReceived, not even an explicit nil
### GetSitLocation

`func (o *PPMShipment) GetSitLocation() SITLocationType`

GetSitLocation returns the SitLocation field if non-nil, zero value otherwise.

### GetSitLocationOk

`func (o *PPMShipment) GetSitLocationOk() (*SITLocationType, bool)`

GetSitLocationOk returns a tuple with the SitLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitLocation

`func (o *PPMShipment) SetSitLocation(v SITLocationType)`

SetSitLocation sets SitLocation field to given value.

### HasSitLocation

`func (o *PPMShipment) HasSitLocation() bool`

HasSitLocation returns a boolean if a field has been set.

### GetSitEstimatedWeight

`func (o *PPMShipment) GetSitEstimatedWeight() int32`

GetSitEstimatedWeight returns the SitEstimatedWeight field if non-nil, zero value otherwise.

### GetSitEstimatedWeightOk

`func (o *PPMShipment) GetSitEstimatedWeightOk() (*int32, bool)`

GetSitEstimatedWeightOk returns a tuple with the SitEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedWeight

`func (o *PPMShipment) SetSitEstimatedWeight(v int32)`

SetSitEstimatedWeight sets SitEstimatedWeight field to given value.

### HasSitEstimatedWeight

`func (o *PPMShipment) HasSitEstimatedWeight() bool`

HasSitEstimatedWeight returns a boolean if a field has been set.

### SetSitEstimatedWeightNil

`func (o *PPMShipment) SetSitEstimatedWeightNil(b bool)`

 SetSitEstimatedWeightNil sets the value for SitEstimatedWeight to be an explicit nil

### UnsetSitEstimatedWeight
`func (o *PPMShipment) UnsetSitEstimatedWeight()`

UnsetSitEstimatedWeight ensures that no value is present for SitEstimatedWeight, not even an explicit nil
### GetSitEstimatedEntryDate

`func (o *PPMShipment) GetSitEstimatedEntryDate() string`

GetSitEstimatedEntryDate returns the SitEstimatedEntryDate field if non-nil, zero value otherwise.

### GetSitEstimatedEntryDateOk

`func (o *PPMShipment) GetSitEstimatedEntryDateOk() (*string, bool)`

GetSitEstimatedEntryDateOk returns a tuple with the SitEstimatedEntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedEntryDate

`func (o *PPMShipment) SetSitEstimatedEntryDate(v string)`

SetSitEstimatedEntryDate sets SitEstimatedEntryDate field to given value.

### HasSitEstimatedEntryDate

`func (o *PPMShipment) HasSitEstimatedEntryDate() bool`

HasSitEstimatedEntryDate returns a boolean if a field has been set.

### SetSitEstimatedEntryDateNil

`func (o *PPMShipment) SetSitEstimatedEntryDateNil(b bool)`

 SetSitEstimatedEntryDateNil sets the value for SitEstimatedEntryDate to be an explicit nil

### UnsetSitEstimatedEntryDate
`func (o *PPMShipment) UnsetSitEstimatedEntryDate()`

UnsetSitEstimatedEntryDate ensures that no value is present for SitEstimatedEntryDate, not even an explicit nil
### GetSitEstimatedDepartureDate

`func (o *PPMShipment) GetSitEstimatedDepartureDate() string`

GetSitEstimatedDepartureDate returns the SitEstimatedDepartureDate field if non-nil, zero value otherwise.

### GetSitEstimatedDepartureDateOk

`func (o *PPMShipment) GetSitEstimatedDepartureDateOk() (*string, bool)`

GetSitEstimatedDepartureDateOk returns a tuple with the SitEstimatedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedDepartureDate

`func (o *PPMShipment) SetSitEstimatedDepartureDate(v string)`

SetSitEstimatedDepartureDate sets SitEstimatedDepartureDate field to given value.

### HasSitEstimatedDepartureDate

`func (o *PPMShipment) HasSitEstimatedDepartureDate() bool`

HasSitEstimatedDepartureDate returns a boolean if a field has been set.

### SetSitEstimatedDepartureDateNil

`func (o *PPMShipment) SetSitEstimatedDepartureDateNil(b bool)`

 SetSitEstimatedDepartureDateNil sets the value for SitEstimatedDepartureDate to be an explicit nil

### UnsetSitEstimatedDepartureDate
`func (o *PPMShipment) UnsetSitEstimatedDepartureDate()`

UnsetSitEstimatedDepartureDate ensures that no value is present for SitEstimatedDepartureDate, not even an explicit nil
### GetSitEstimatedCost

`func (o *PPMShipment) GetSitEstimatedCost() int32`

GetSitEstimatedCost returns the SitEstimatedCost field if non-nil, zero value otherwise.

### GetSitEstimatedCostOk

`func (o *PPMShipment) GetSitEstimatedCostOk() (*int32, bool)`

GetSitEstimatedCostOk returns a tuple with the SitEstimatedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedCost

`func (o *PPMShipment) SetSitEstimatedCost(v int32)`

SetSitEstimatedCost sets SitEstimatedCost field to given value.

### HasSitEstimatedCost

`func (o *PPMShipment) HasSitEstimatedCost() bool`

HasSitEstimatedCost returns a boolean if a field has been set.

### SetSitEstimatedCostNil

`func (o *PPMShipment) SetSitEstimatedCostNil(b bool)`

 SetSitEstimatedCostNil sets the value for SitEstimatedCost to be an explicit nil

### UnsetSitEstimatedCost
`func (o *PPMShipment) UnsetSitEstimatedCost()`

UnsetSitEstimatedCost ensures that no value is present for SitEstimatedCost, not even an explicit nil
### GetETag

`func (o *PPMShipment) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *PPMShipment) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *PPMShipment) SetETag(v string)`

SetETag sets ETag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


