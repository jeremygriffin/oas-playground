# OpenapiClient::PPMShipmentV2

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** | The primary unique identifier of this PPM shipment | [readonly] |
| **shipment_id** | **String** | The id of the parent MTOShipment record | [readonly] |
| **created_at** | **Time** | The timestamp of when the PPM shipment was created (UTC) | [readonly] |
| **updated_at** | **Time** | The timestamp of when a property of this object was last updated (UTC) | [optional][readonly] |
| **status** | [**PPMShipmentStatusV2**](PPMShipmentStatusV2.md) |  |  |
| **expected_departure_date** | **Date** | Date the customer expects to begin moving from their origin.  |  |
| **actual_move_date** | **Date** | The actual start date of when the PPM shipment left the origin. | [optional] |
| **submitted_at** | **Time** | The timestamp of when the customer submitted their PPM documentation to the counselor for review. | [optional] |
| **reviewed_at** | **Time** | The timestamp of when the Service Counselor has reviewed all of the closeout documents. | [optional] |
| **approved_at** | **Time** | The timestamp of when the shipment was approved and the service member can begin their move. | [optional] |
| **actual_pickup_postal_code** | **String** | The actual postal code where the PPM shipment started. To be filled once the customer has moved the shipment.  | [optional] |
| **actual_destination_postal_code** | **String** | The actual postal code where the PPM shipment ended. To be filled once the customer has moved the shipment.  | [optional] |
| **sit_expected** | **Boolean** | Captures whether some or all of the PPM shipment will require temporary storage at the origin or destination.  Must be set to &#x60;true&#x60; when providing &#x60;sitLocation&#x60;, &#x60;sitEstimatedWeight&#x60;, &#x60;sitEstimatedEntryDate&#x60;, and &#x60;sitEstimatedDepartureDate&#x60; values to calculate the &#x60;sitEstimatedCost&#x60;.  |  |
| **estimated_weight** | **Integer** | The estimated weight of the PPM shipment goods being moved in pounds. | [optional] |
| **has_pro_gear** | **Boolean** | Indicates whether PPM shipment has pro gear for themselves or their spouse.  | [optional] |
| **pro_gear_weight** | **Integer** | The estimated weight of the pro-gear being moved belonging to the service member in pounds. | [optional] |
| **spouse_pro_gear_weight** | **Integer** | The estimated weight of the pro-gear being moved belonging to a spouse in pounds. | [optional] |
| **estimated_incentive** | **Integer** | The estimated amount the government will pay the service member to move their belongings based on the moving date, locations, and shipment weight. | [optional] |
| **has_requested_advance** | **Boolean** | Indicates whether an advance has been requested for the PPM shipment.  | [optional] |
| **advance_amount_requested** | **Integer** | The amount requested as an advance by the service member, up to a maximum percentage of the estimated incentive.  | [optional] |
| **has_received_advance** | **Boolean** | Indicates whether an advance was received for the PPM shipment.  | [optional] |
| **advance_amount_received** | **Integer** | The amount received for an advance, or null if no advance is received.  | [optional] |
| **sit_location** | [**SITLocationTypeV2**](SITLocationTypeV2.md) |  | [optional] |
| **sit_estimated_weight** | **Integer** | The estimated weight of the goods being put into storage in pounds. | [optional] |
| **sit_estimated_entry_date** | **Date** | The date that goods will first enter the storage location. | [optional] |
| **sit_estimated_departure_date** | **Date** | The date that goods will exit the storage location. | [optional] |
| **sit_estimated_cost** | **Integer** | The estimated amount that the government will pay the service member to put their goods into storage. This estimated storage cost is separate from the estimated incentive. | [optional] |
| **e_tag** | **String** | A hash unique to this shipment that should be used as the \&quot;If-Match\&quot; header for any updates. | [readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::PPMShipmentV2.new(
  id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  shipment_id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  created_at: null,
  updated_at: null,
  status: null,
  expected_departure_date: null,
  actual_move_date: null,
  submitted_at: null,
  reviewed_at: null,
  approved_at: null,
  actual_pickup_postal_code: 90210,
  actual_destination_postal_code: 90210,
  sit_expected: null,
  estimated_weight: 4200,
  has_pro_gear: null,
  pro_gear_weight: null,
  spouse_pro_gear_weight: null,
  estimated_incentive: null,
  has_requested_advance: null,
  advance_amount_requested: null,
  has_received_advance: null,
  advance_amount_received: null,
  sit_location: null,
  sit_estimated_weight: 2000,
  sit_estimated_entry_date: null,
  sit_estimated_departure_date: null,
  sit_estimated_cost: null,
  e_tag: null
)
```

