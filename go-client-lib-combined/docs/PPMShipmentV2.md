# PPMShipmentV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | The primary unique identifier of this PPM shipment | [readonly] 
**ShipmentId** | **string** | The id of the parent MTOShipment record | [readonly] 
**CreatedAt** | **time.Time** | The timestamp of when the PPM shipment was created (UTC) | [readonly] 
**UpdatedAt** | Pointer to **time.Time** | The timestamp of when a property of this object was last updated (UTC) | [optional] [readonly] 
**Status** | [**PPMShipmentStatusV2V2**](PPMShipmentStatusV2.md) |  | 
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
**SitLocation** | Pointer to [**SITLocationTypeV2**](SITLocationType.md) |  | [optional] 
**SitEstimatedWeight** | Pointer to **NullableInt32** | The estimated weight of the goods being put into storage in pounds. | [optional] 
**SitEstimatedEntryDate** | Pointer to **NullableString** | The date that goods will first enter the storage location. | [optional] 
**SitEstimatedDepartureDate** | Pointer to **NullableString** | The date that goods will exit the storage location. | [optional] 
**SitEstimatedCost** | Pointer to **NullableInt32** | The estimated amount that the government will pay the service member to put their goods into storage. This estimated storage cost is separate from the estimated incentive. | [optional] 
**ETag** | **string** | A hash unique to this shipment that should be used as the \&quot;If-Match\&quot; header for any updates. | [readonly] 

## Methods

### NewPPMShipmentV2

`func NewPPMShipmentV2(id string, shipmentId string, createdAt time.Time, status PPMShipmentStatusV2V2, expectedDepartureDate string, sitExpected bool, eTag string, ) *PPMShipmentV2`

NewPPMShipmentV2 instantiates a new PPMShipmentV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPPMShipmentV2WithDefaults

`func NewPPMShipmentV2WithDefaults() *PPMShipmentV2`

NewPPMShipmentV2WithDefaults instantiates a new PPMShipmentV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PPMShipmentV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PPMShipmentV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PPMShipmentV2) SetId(v string)`

SetId sets Id field to given value.


### GetShipmentId

`func (o *PPMShipmentV2) GetShipmentId() string`

GetShipmentId returns the ShipmentId field if non-nil, zero value otherwise.

### GetShipmentIdOk

`func (o *PPMShipmentV2) GetShipmentIdOk() (*string, bool)`

GetShipmentIdOk returns a tuple with the ShipmentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentId

`func (o *PPMShipmentV2) SetShipmentId(v string)`

SetShipmentId sets ShipmentId field to given value.


### GetCreatedAt

`func (o *PPMShipmentV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *PPMShipmentV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *PPMShipmentV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.


### GetUpdatedAt

`func (o *PPMShipmentV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *PPMShipmentV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *PPMShipmentV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *PPMShipmentV2) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetStatus

`func (o *PPMShipmentV2) GetStatus() PPMShipmentStatusV2V2`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *PPMShipmentV2) GetStatusOk() (*PPMShipmentStatusV2V2, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *PPMShipmentV2) SetStatus(v PPMShipmentStatusV2V2)`

SetStatus sets Status field to given value.


### GetExpectedDepartureDate

`func (o *PPMShipmentV2) GetExpectedDepartureDate() string`

GetExpectedDepartureDate returns the ExpectedDepartureDate field if non-nil, zero value otherwise.

### GetExpectedDepartureDateOk

`func (o *PPMShipmentV2) GetExpectedDepartureDateOk() (*string, bool)`

GetExpectedDepartureDateOk returns a tuple with the ExpectedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpectedDepartureDate

`func (o *PPMShipmentV2) SetExpectedDepartureDate(v string)`

SetExpectedDepartureDate sets ExpectedDepartureDate field to given value.


### GetActualMoveDate

`func (o *PPMShipmentV2) GetActualMoveDate() string`

GetActualMoveDate returns the ActualMoveDate field if non-nil, zero value otherwise.

### GetActualMoveDateOk

`func (o *PPMShipmentV2) GetActualMoveDateOk() (*string, bool)`

GetActualMoveDateOk returns a tuple with the ActualMoveDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualMoveDate

`func (o *PPMShipmentV2) SetActualMoveDate(v string)`

SetActualMoveDate sets ActualMoveDate field to given value.

### HasActualMoveDate

`func (o *PPMShipmentV2) HasActualMoveDate() bool`

HasActualMoveDate returns a boolean if a field has been set.

### SetActualMoveDateNil

`func (o *PPMShipmentV2) SetActualMoveDateNil(b bool)`

 SetActualMoveDateNil sets the value for ActualMoveDate to be an explicit nil

### UnsetActualMoveDate
`func (o *PPMShipmentV2) UnsetActualMoveDate()`

UnsetActualMoveDate ensures that no value is present for ActualMoveDate, not even an explicit nil
### GetSubmittedAt

`func (o *PPMShipmentV2) GetSubmittedAt() time.Time`

GetSubmittedAt returns the SubmittedAt field if non-nil, zero value otherwise.

### GetSubmittedAtOk

`func (o *PPMShipmentV2) GetSubmittedAtOk() (*time.Time, bool)`

GetSubmittedAtOk returns a tuple with the SubmittedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubmittedAt

`func (o *PPMShipmentV2) SetSubmittedAt(v time.Time)`

SetSubmittedAt sets SubmittedAt field to given value.

### HasSubmittedAt

`func (o *PPMShipmentV2) HasSubmittedAt() bool`

HasSubmittedAt returns a boolean if a field has been set.

### SetSubmittedAtNil

`func (o *PPMShipmentV2) SetSubmittedAtNil(b bool)`

 SetSubmittedAtNil sets the value for SubmittedAt to be an explicit nil

### UnsetSubmittedAt
`func (o *PPMShipmentV2) UnsetSubmittedAt()`

UnsetSubmittedAt ensures that no value is present for SubmittedAt, not even an explicit nil
### GetReviewedAt

`func (o *PPMShipmentV2) GetReviewedAt() time.Time`

GetReviewedAt returns the ReviewedAt field if non-nil, zero value otherwise.

### GetReviewedAtOk

`func (o *PPMShipmentV2) GetReviewedAtOk() (*time.Time, bool)`

GetReviewedAtOk returns a tuple with the ReviewedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReviewedAt

`func (o *PPMShipmentV2) SetReviewedAt(v time.Time)`

SetReviewedAt sets ReviewedAt field to given value.

### HasReviewedAt

`func (o *PPMShipmentV2) HasReviewedAt() bool`

HasReviewedAt returns a boolean if a field has been set.

### SetReviewedAtNil

`func (o *PPMShipmentV2) SetReviewedAtNil(b bool)`

 SetReviewedAtNil sets the value for ReviewedAt to be an explicit nil

### UnsetReviewedAt
`func (o *PPMShipmentV2) UnsetReviewedAt()`

UnsetReviewedAt ensures that no value is present for ReviewedAt, not even an explicit nil
### GetApprovedAt

`func (o *PPMShipmentV2) GetApprovedAt() time.Time`

GetApprovedAt returns the ApprovedAt field if non-nil, zero value otherwise.

### GetApprovedAtOk

`func (o *PPMShipmentV2) GetApprovedAtOk() (*time.Time, bool)`

GetApprovedAtOk returns a tuple with the ApprovedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedAt

`func (o *PPMShipmentV2) SetApprovedAt(v time.Time)`

SetApprovedAt sets ApprovedAt field to given value.

### HasApprovedAt

`func (o *PPMShipmentV2) HasApprovedAt() bool`

HasApprovedAt returns a boolean if a field has been set.

### SetApprovedAtNil

`func (o *PPMShipmentV2) SetApprovedAtNil(b bool)`

 SetApprovedAtNil sets the value for ApprovedAt to be an explicit nil

### UnsetApprovedAt
`func (o *PPMShipmentV2) UnsetApprovedAt()`

UnsetApprovedAt ensures that no value is present for ApprovedAt, not even an explicit nil
### GetActualPickupPostalCode

`func (o *PPMShipmentV2) GetActualPickupPostalCode() string`

GetActualPickupPostalCode returns the ActualPickupPostalCode field if non-nil, zero value otherwise.

### GetActualPickupPostalCodeOk

`func (o *PPMShipmentV2) GetActualPickupPostalCodeOk() (*string, bool)`

GetActualPickupPostalCodeOk returns a tuple with the ActualPickupPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualPickupPostalCode

`func (o *PPMShipmentV2) SetActualPickupPostalCode(v string)`

SetActualPickupPostalCode sets ActualPickupPostalCode field to given value.

### HasActualPickupPostalCode

`func (o *PPMShipmentV2) HasActualPickupPostalCode() bool`

HasActualPickupPostalCode returns a boolean if a field has been set.

### SetActualPickupPostalCodeNil

`func (o *PPMShipmentV2) SetActualPickupPostalCodeNil(b bool)`

 SetActualPickupPostalCodeNil sets the value for ActualPickupPostalCode to be an explicit nil

### UnsetActualPickupPostalCode
`func (o *PPMShipmentV2) UnsetActualPickupPostalCode()`

UnsetActualPickupPostalCode ensures that no value is present for ActualPickupPostalCode, not even an explicit nil
### GetActualDestinationPostalCode

`func (o *PPMShipmentV2) GetActualDestinationPostalCode() string`

GetActualDestinationPostalCode returns the ActualDestinationPostalCode field if non-nil, zero value otherwise.

### GetActualDestinationPostalCodeOk

`func (o *PPMShipmentV2) GetActualDestinationPostalCodeOk() (*string, bool)`

GetActualDestinationPostalCodeOk returns a tuple with the ActualDestinationPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualDestinationPostalCode

`func (o *PPMShipmentV2) SetActualDestinationPostalCode(v string)`

SetActualDestinationPostalCode sets ActualDestinationPostalCode field to given value.

### HasActualDestinationPostalCode

`func (o *PPMShipmentV2) HasActualDestinationPostalCode() bool`

HasActualDestinationPostalCode returns a boolean if a field has been set.

### SetActualDestinationPostalCodeNil

`func (o *PPMShipmentV2) SetActualDestinationPostalCodeNil(b bool)`

 SetActualDestinationPostalCodeNil sets the value for ActualDestinationPostalCode to be an explicit nil

### UnsetActualDestinationPostalCode
`func (o *PPMShipmentV2) UnsetActualDestinationPostalCode()`

UnsetActualDestinationPostalCode ensures that no value is present for ActualDestinationPostalCode, not even an explicit nil
### GetSitExpected

`func (o *PPMShipmentV2) GetSitExpected() bool`

GetSitExpected returns the SitExpected field if non-nil, zero value otherwise.

### GetSitExpectedOk

`func (o *PPMShipmentV2) GetSitExpectedOk() (*bool, bool)`

GetSitExpectedOk returns a tuple with the SitExpected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitExpected

`func (o *PPMShipmentV2) SetSitExpected(v bool)`

SetSitExpected sets SitExpected field to given value.


### GetEstimatedWeight

`func (o *PPMShipmentV2) GetEstimatedWeight() int32`

GetEstimatedWeight returns the EstimatedWeight field if non-nil, zero value otherwise.

### GetEstimatedWeightOk

`func (o *PPMShipmentV2) GetEstimatedWeightOk() (*int32, bool)`

GetEstimatedWeightOk returns a tuple with the EstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedWeight

`func (o *PPMShipmentV2) SetEstimatedWeight(v int32)`

SetEstimatedWeight sets EstimatedWeight field to given value.

### HasEstimatedWeight

`func (o *PPMShipmentV2) HasEstimatedWeight() bool`

HasEstimatedWeight returns a boolean if a field has been set.

### SetEstimatedWeightNil

`func (o *PPMShipmentV2) SetEstimatedWeightNil(b bool)`

 SetEstimatedWeightNil sets the value for EstimatedWeight to be an explicit nil

### UnsetEstimatedWeight
`func (o *PPMShipmentV2) UnsetEstimatedWeight()`

UnsetEstimatedWeight ensures that no value is present for EstimatedWeight, not even an explicit nil
### GetHasProGear

`func (o *PPMShipmentV2) GetHasProGear() bool`

GetHasProGear returns the HasProGear field if non-nil, zero value otherwise.

### GetHasProGearOk

`func (o *PPMShipmentV2) GetHasProGearOk() (*bool, bool)`

GetHasProGearOk returns a tuple with the HasProGear field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasProGear

`func (o *PPMShipmentV2) SetHasProGear(v bool)`

SetHasProGear sets HasProGear field to given value.

### HasHasProGear

`func (o *PPMShipmentV2) HasHasProGear() bool`

HasHasProGear returns a boolean if a field has been set.

### SetHasProGearNil

`func (o *PPMShipmentV2) SetHasProGearNil(b bool)`

 SetHasProGearNil sets the value for HasProGear to be an explicit nil

### UnsetHasProGear
`func (o *PPMShipmentV2) UnsetHasProGear()`

UnsetHasProGear ensures that no value is present for HasProGear, not even an explicit nil
### GetProGearWeight

`func (o *PPMShipmentV2) GetProGearWeight() int32`

GetProGearWeight returns the ProGearWeight field if non-nil, zero value otherwise.

### GetProGearWeightOk

`func (o *PPMShipmentV2) GetProGearWeightOk() (*int32, bool)`

GetProGearWeightOk returns a tuple with the ProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProGearWeight

`func (o *PPMShipmentV2) SetProGearWeight(v int32)`

SetProGearWeight sets ProGearWeight field to given value.

### HasProGearWeight

`func (o *PPMShipmentV2) HasProGearWeight() bool`

HasProGearWeight returns a boolean if a field has been set.

### SetProGearWeightNil

`func (o *PPMShipmentV2) SetProGearWeightNil(b bool)`

 SetProGearWeightNil sets the value for ProGearWeight to be an explicit nil

### UnsetProGearWeight
`func (o *PPMShipmentV2) UnsetProGearWeight()`

UnsetProGearWeight ensures that no value is present for ProGearWeight, not even an explicit nil
### GetSpouseProGearWeight

`func (o *PPMShipmentV2) GetSpouseProGearWeight() int32`

GetSpouseProGearWeight returns the SpouseProGearWeight field if non-nil, zero value otherwise.

### GetSpouseProGearWeightOk

`func (o *PPMShipmentV2) GetSpouseProGearWeightOk() (*int32, bool)`

GetSpouseProGearWeightOk returns a tuple with the SpouseProGearWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpouseProGearWeight

`func (o *PPMShipmentV2) SetSpouseProGearWeight(v int32)`

SetSpouseProGearWeight sets SpouseProGearWeight field to given value.

### HasSpouseProGearWeight

`func (o *PPMShipmentV2) HasSpouseProGearWeight() bool`

HasSpouseProGearWeight returns a boolean if a field has been set.

### SetSpouseProGearWeightNil

`func (o *PPMShipmentV2) SetSpouseProGearWeightNil(b bool)`

 SetSpouseProGearWeightNil sets the value for SpouseProGearWeight to be an explicit nil

### UnsetSpouseProGearWeight
`func (o *PPMShipmentV2) UnsetSpouseProGearWeight()`

UnsetSpouseProGearWeight ensures that no value is present for SpouseProGearWeight, not even an explicit nil
### GetEstimatedIncentive

`func (o *PPMShipmentV2) GetEstimatedIncentive() int32`

GetEstimatedIncentive returns the EstimatedIncentive field if non-nil, zero value otherwise.

### GetEstimatedIncentiveOk

`func (o *PPMShipmentV2) GetEstimatedIncentiveOk() (*int32, bool)`

GetEstimatedIncentiveOk returns a tuple with the EstimatedIncentive field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedIncentive

`func (o *PPMShipmentV2) SetEstimatedIncentive(v int32)`

SetEstimatedIncentive sets EstimatedIncentive field to given value.

### HasEstimatedIncentive

`func (o *PPMShipmentV2) HasEstimatedIncentive() bool`

HasEstimatedIncentive returns a boolean if a field has been set.

### SetEstimatedIncentiveNil

`func (o *PPMShipmentV2) SetEstimatedIncentiveNil(b bool)`

 SetEstimatedIncentiveNil sets the value for EstimatedIncentive to be an explicit nil

### UnsetEstimatedIncentive
`func (o *PPMShipmentV2) UnsetEstimatedIncentive()`

UnsetEstimatedIncentive ensures that no value is present for EstimatedIncentive, not even an explicit nil
### GetHasRequestedAdvance

`func (o *PPMShipmentV2) GetHasRequestedAdvance() bool`

GetHasRequestedAdvance returns the HasRequestedAdvance field if non-nil, zero value otherwise.

### GetHasRequestedAdvanceOk

`func (o *PPMShipmentV2) GetHasRequestedAdvanceOk() (*bool, bool)`

GetHasRequestedAdvanceOk returns a tuple with the HasRequestedAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRequestedAdvance

`func (o *PPMShipmentV2) SetHasRequestedAdvance(v bool)`

SetHasRequestedAdvance sets HasRequestedAdvance field to given value.

### HasHasRequestedAdvance

`func (o *PPMShipmentV2) HasHasRequestedAdvance() bool`

HasHasRequestedAdvance returns a boolean if a field has been set.

### SetHasRequestedAdvanceNil

`func (o *PPMShipmentV2) SetHasRequestedAdvanceNil(b bool)`

 SetHasRequestedAdvanceNil sets the value for HasRequestedAdvance to be an explicit nil

### UnsetHasRequestedAdvance
`func (o *PPMShipmentV2) UnsetHasRequestedAdvance()`

UnsetHasRequestedAdvance ensures that no value is present for HasRequestedAdvance, not even an explicit nil
### GetAdvanceAmountRequested

`func (o *PPMShipmentV2) GetAdvanceAmountRequested() int32`

GetAdvanceAmountRequested returns the AdvanceAmountRequested field if non-nil, zero value otherwise.

### GetAdvanceAmountRequestedOk

`func (o *PPMShipmentV2) GetAdvanceAmountRequestedOk() (*int32, bool)`

GetAdvanceAmountRequestedOk returns a tuple with the AdvanceAmountRequested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanceAmountRequested

`func (o *PPMShipmentV2) SetAdvanceAmountRequested(v int32)`

SetAdvanceAmountRequested sets AdvanceAmountRequested field to given value.

### HasAdvanceAmountRequested

`func (o *PPMShipmentV2) HasAdvanceAmountRequested() bool`

HasAdvanceAmountRequested returns a boolean if a field has been set.

### SetAdvanceAmountRequestedNil

`func (o *PPMShipmentV2) SetAdvanceAmountRequestedNil(b bool)`

 SetAdvanceAmountRequestedNil sets the value for AdvanceAmountRequested to be an explicit nil

### UnsetAdvanceAmountRequested
`func (o *PPMShipmentV2) UnsetAdvanceAmountRequested()`

UnsetAdvanceAmountRequested ensures that no value is present for AdvanceAmountRequested, not even an explicit nil
### GetHasReceivedAdvance

`func (o *PPMShipmentV2) GetHasReceivedAdvance() bool`

GetHasReceivedAdvance returns the HasReceivedAdvance field if non-nil, zero value otherwise.

### GetHasReceivedAdvanceOk

`func (o *PPMShipmentV2) GetHasReceivedAdvanceOk() (*bool, bool)`

GetHasReceivedAdvanceOk returns a tuple with the HasReceivedAdvance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasReceivedAdvance

`func (o *PPMShipmentV2) SetHasReceivedAdvance(v bool)`

SetHasReceivedAdvance sets HasReceivedAdvance field to given value.

### HasHasReceivedAdvance

`func (o *PPMShipmentV2) HasHasReceivedAdvance() bool`

HasHasReceivedAdvance returns a boolean if a field has been set.

### SetHasReceivedAdvanceNil

`func (o *PPMShipmentV2) SetHasReceivedAdvanceNil(b bool)`

 SetHasReceivedAdvanceNil sets the value for HasReceivedAdvance to be an explicit nil

### UnsetHasReceivedAdvance
`func (o *PPMShipmentV2) UnsetHasReceivedAdvance()`

UnsetHasReceivedAdvance ensures that no value is present for HasReceivedAdvance, not even an explicit nil
### GetAdvanceAmountReceived

`func (o *PPMShipmentV2) GetAdvanceAmountReceived() int32`

GetAdvanceAmountReceived returns the AdvanceAmountReceived field if non-nil, zero value otherwise.

### GetAdvanceAmountReceivedOk

`func (o *PPMShipmentV2) GetAdvanceAmountReceivedOk() (*int32, bool)`

GetAdvanceAmountReceivedOk returns a tuple with the AdvanceAmountReceived field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdvanceAmountReceived

`func (o *PPMShipmentV2) SetAdvanceAmountReceived(v int32)`

SetAdvanceAmountReceived sets AdvanceAmountReceived field to given value.

### HasAdvanceAmountReceived

`func (o *PPMShipmentV2) HasAdvanceAmountReceived() bool`

HasAdvanceAmountReceived returns a boolean if a field has been set.

### SetAdvanceAmountReceivedNil

`func (o *PPMShipmentV2) SetAdvanceAmountReceivedNil(b bool)`

 SetAdvanceAmountReceivedNil sets the value for AdvanceAmountReceived to be an explicit nil

### UnsetAdvanceAmountReceived
`func (o *PPMShipmentV2) UnsetAdvanceAmountReceived()`

UnsetAdvanceAmountReceived ensures that no value is present for AdvanceAmountReceived, not even an explicit nil
### GetSitLocation

`func (o *PPMShipmentV2) GetSitLocation() SITLocationTypeV2`

GetSitLocation returns the SitLocation field if non-nil, zero value otherwise.

### GetSitLocationOk

`func (o *PPMShipmentV2) GetSitLocationOk() (*SITLocationTypeV2, bool)`

GetSitLocationOk returns a tuple with the SitLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitLocation

`func (o *PPMShipmentV2) SetSitLocation(v SITLocationTypeV2)`

SetSitLocation sets SitLocation field to given value.

### HasSitLocation

`func (o *PPMShipmentV2) HasSitLocation() bool`

HasSitLocation returns a boolean if a field has been set.

### GetSitEstimatedWeight

`func (o *PPMShipmentV2) GetSitEstimatedWeight() int32`

GetSitEstimatedWeight returns the SitEstimatedWeight field if non-nil, zero value otherwise.

### GetSitEstimatedWeightOk

`func (o *PPMShipmentV2) GetSitEstimatedWeightOk() (*int32, bool)`

GetSitEstimatedWeightOk returns a tuple with the SitEstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedWeight

`func (o *PPMShipmentV2) SetSitEstimatedWeight(v int32)`

SetSitEstimatedWeight sets SitEstimatedWeight field to given value.

### HasSitEstimatedWeight

`func (o *PPMShipmentV2) HasSitEstimatedWeight() bool`

HasSitEstimatedWeight returns a boolean if a field has been set.

### SetSitEstimatedWeightNil

`func (o *PPMShipmentV2) SetSitEstimatedWeightNil(b bool)`

 SetSitEstimatedWeightNil sets the value for SitEstimatedWeight to be an explicit nil

### UnsetSitEstimatedWeight
`func (o *PPMShipmentV2) UnsetSitEstimatedWeight()`

UnsetSitEstimatedWeight ensures that no value is present for SitEstimatedWeight, not even an explicit nil
### GetSitEstimatedEntryDate

`func (o *PPMShipmentV2) GetSitEstimatedEntryDate() string`

GetSitEstimatedEntryDate returns the SitEstimatedEntryDate field if non-nil, zero value otherwise.

### GetSitEstimatedEntryDateOk

`func (o *PPMShipmentV2) GetSitEstimatedEntryDateOk() (*string, bool)`

GetSitEstimatedEntryDateOk returns a tuple with the SitEstimatedEntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedEntryDate

`func (o *PPMShipmentV2) SetSitEstimatedEntryDate(v string)`

SetSitEstimatedEntryDate sets SitEstimatedEntryDate field to given value.

### HasSitEstimatedEntryDate

`func (o *PPMShipmentV2) HasSitEstimatedEntryDate() bool`

HasSitEstimatedEntryDate returns a boolean if a field has been set.

### SetSitEstimatedEntryDateNil

`func (o *PPMShipmentV2) SetSitEstimatedEntryDateNil(b bool)`

 SetSitEstimatedEntryDateNil sets the value for SitEstimatedEntryDate to be an explicit nil

### UnsetSitEstimatedEntryDate
`func (o *PPMShipmentV2) UnsetSitEstimatedEntryDate()`

UnsetSitEstimatedEntryDate ensures that no value is present for SitEstimatedEntryDate, not even an explicit nil
### GetSitEstimatedDepartureDate

`func (o *PPMShipmentV2) GetSitEstimatedDepartureDate() string`

GetSitEstimatedDepartureDate returns the SitEstimatedDepartureDate field if non-nil, zero value otherwise.

### GetSitEstimatedDepartureDateOk

`func (o *PPMShipmentV2) GetSitEstimatedDepartureDateOk() (*string, bool)`

GetSitEstimatedDepartureDateOk returns a tuple with the SitEstimatedDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedDepartureDate

`func (o *PPMShipmentV2) SetSitEstimatedDepartureDate(v string)`

SetSitEstimatedDepartureDate sets SitEstimatedDepartureDate field to given value.

### HasSitEstimatedDepartureDate

`func (o *PPMShipmentV2) HasSitEstimatedDepartureDate() bool`

HasSitEstimatedDepartureDate returns a boolean if a field has been set.

### SetSitEstimatedDepartureDateNil

`func (o *PPMShipmentV2) SetSitEstimatedDepartureDateNil(b bool)`

 SetSitEstimatedDepartureDateNil sets the value for SitEstimatedDepartureDate to be an explicit nil

### UnsetSitEstimatedDepartureDate
`func (o *PPMShipmentV2) UnsetSitEstimatedDepartureDate()`

UnsetSitEstimatedDepartureDate ensures that no value is present for SitEstimatedDepartureDate, not even an explicit nil
### GetSitEstimatedCost

`func (o *PPMShipmentV2) GetSitEstimatedCost() int32`

GetSitEstimatedCost returns the SitEstimatedCost field if non-nil, zero value otherwise.

### GetSitEstimatedCostOk

`func (o *PPMShipmentV2) GetSitEstimatedCostOk() (*int32, bool)`

GetSitEstimatedCostOk returns a tuple with the SitEstimatedCost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEstimatedCost

`func (o *PPMShipmentV2) SetSitEstimatedCost(v int32)`

SetSitEstimatedCost sets SitEstimatedCost field to given value.

### HasSitEstimatedCost

`func (o *PPMShipmentV2) HasSitEstimatedCost() bool`

HasSitEstimatedCost returns a boolean if a field has been set.

### SetSitEstimatedCostNil

`func (o *PPMShipmentV2) SetSitEstimatedCostNil(b bool)`

 SetSitEstimatedCostNil sets the value for SitEstimatedCost to be an explicit nil

### UnsetSitEstimatedCost
`func (o *PPMShipmentV2) UnsetSitEstimatedCost()`

UnsetSitEstimatedCost ensures that no value is present for SitEstimatedCost, not even an explicit nil
### GetETag

`func (o *PPMShipmentV2) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *PPMShipmentV2) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *PPMShipmentV2) SetETag(v string)`

SetETag sets ETag field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


