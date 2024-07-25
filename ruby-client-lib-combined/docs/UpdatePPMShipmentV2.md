# OpenapiClient::UpdatePPMShipmentV2

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **expected_departure_date** | **Date** | Date the customer expects to begin moving from their origin.  | [optional] |
| **sit_expected** | **Boolean** | Captures whether some or all of the PPM shipment will require temporary storage at the origin or destination.  Must be set to &#x60;true&#x60; when providing &#x60;sitLocation&#x60;, &#x60;sitEstimatedWeight&#x60;, &#x60;sitEstimatedEntryDate&#x60;, and &#x60;sitEstimatedDepartureDate&#x60; values to calculate the &#x60;sitEstimatedCost&#x60;.  | [optional] |
| **sit_location** | [**SITLocationTypeV2**](SITLocationTypeV2.md) |  | [optional] |
| **sit_estimated_weight** | **Integer** | The estimated weight of the goods being put into storage. | [optional] |
| **sit_estimated_entry_date** | **Date** | The date that goods will first enter the storage location. | [optional] |
| **sit_estimated_departure_date** | **Date** | The date that goods will exit the storage location. | [optional] |
| **estimated_weight** | **Integer** | The estimated weight of the PPM shipment goods being moved. | [optional] |
| **has_pro_gear** | **Boolean** | Indicates whether PPM shipment has pro gear for themselves or their spouse.  | [optional] |
| **pro_gear_weight** | **Integer** | The estimated weight of the pro-gear being moved belonging to the service member. | [optional] |
| **spouse_pro_gear_weight** | **Integer** | The estimated weight of the pro-gear being moved belonging to a spouse. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::UpdatePPMShipmentV2.new(
  expected_departure_date: null,
  sit_expected: null,
  sit_location: null,
  sit_estimated_weight: 2000,
  sit_estimated_entry_date: null,
  sit_estimated_departure_date: null,
  estimated_weight: 4200,
  has_pro_gear: null,
  pro_gear_weight: null,
  spouse_pro_gear_weight: null
)
```

