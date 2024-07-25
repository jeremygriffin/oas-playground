# OpenapiClient::UpdatePPMShipmentV3

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **expected_departure_date** | **Date** | Date the customer expects to begin moving from their origin.  | [optional] |
| **pickup_address** | [**AddressV3**](AddressV3.md) | The address of the origin location where goods are being moved from.  | [optional] |
| **has_secondary_pickup_address** | **Boolean** |  | [optional] |
| **secondary_pickup_address** | [**AddressV3**](AddressV3.md) | An optional secondary pickup location near the origin where additional goods exist.  | [optional] |
| **destination_address** | [**AddressV3**](AddressV3.md) | The address of the destination location where goods are being delivered to.  | [optional] |
| **has_secondary_destination_address** | **Boolean** |  | [optional] |
| **secondary_destination_address** | [**AddressV3**](AddressV3.md) | An optional secondary address near the destination where goods will be dropped off.  | [optional] |
| **sit_expected** | **Boolean** | Captures whether some or all of the PPM shipment will require temporary storage at the origin or destination.  Must be set to &#x60;true&#x60; when providing &#x60;sitLocation&#x60;, &#x60;sitEstimatedWeight&#x60;, &#x60;sitEstimatedEntryDate&#x60;, and &#x60;sitEstimatedDepartureDate&#x60; values to calculate the &#x60;sitEstimatedCost&#x60;.  | [optional] |
| **sit_location** | [**SITLocationTypeV3**](SITLocationTypeV3.md) |  | [optional] |
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

instance = OpenapiClient::UpdatePPMShipmentV3.new(
  expected_departure_date: null,
  pickup_address: null,
  has_secondary_pickup_address: null,
  secondary_pickup_address: null,
  destination_address: null,
  has_secondary_destination_address: null,
  secondary_destination_address: null,
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

