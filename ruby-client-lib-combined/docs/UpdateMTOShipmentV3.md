# OpenapiClient::UpdateMTOShipmentV3

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **actual_pro_gear_weight** | **Integer** | The actual weight of any pro gear shipped during a move. | [optional] |
| **actual_spouse_pro_gear_weight** | **Integer** | The actual weight of any pro gear shipped during a move. | [optional] |
| **scheduled_pickup_date** | **Date** | The date the Prime contractor scheduled to pick up this shipment after consultation with the customer. | [optional] |
| **actual_pickup_date** | **Date** | The date when the Prime contractor actually picked up the shipment. Updated after-the-fact. | [optional] |
| **first_available_delivery_date** | **Date** | The date the Prime provides to the customer as the first possible delivery date so that they can plan their travel accordingly.  | [optional] |
| **scheduled_delivery_date** | **Date** | The date the Prime contractor scheduled to deliver this shipment after consultation with the customer. | [optional] |
| **actual_delivery_date** | **Date** | The date when the Prime contractor actually delivered the shipment. Updated after-the-fact. | [optional] |
| **prime_estimated_weight** | **Integer** | The estimated weight of this shipment, determined by the movers during the pre-move survey. This value **can only be updated once.** If there was an issue with estimating the weight and a mistake was made, the Prime contracter will need to contact the TOO to change it.  | [optional] |
| **prime_actual_weight** | **Integer** | The actual weight of the shipment, provided after the Prime packs, picks up, and weighs a customer&#39;s shipment. | [optional] |
| **nts_recorded_weight** | **Integer** | The previously recorded weight for the NTS Shipment. Used for NTS Release to know what the previous primeActualWeight or billable weight was. | [optional] |
| **pickup_address** | [**AddressV3**](AddressV3.md) | The address where the movers should pick up this shipment, entered by the customer during onboarding when they enter shipment details.  | [optional] |
| **destination_address** | [**AddressV3**](AddressV3.md) | Where the movers should deliver this shipment. Often provided by the customer when they enter shipment details during onboarding, if they know their new address already.  May be blank when entered by the customer, required when entered by the Prime. May not represent the true final destination due to the shipment being diverted or placed in SIT.  | [optional] |
| **destination_type** | [**DestinationTypeV3**](DestinationTypeV3.md) |  | [optional] |
| **secondary_pickup_address** | [**AddressV3**](AddressV3.md) | A second pickup address for this shipment, if the customer entered one. An optional field. | [optional] |
| **secondary_delivery_address** | [**AddressV3**](AddressV3.md) | A second delivery address for this shipment, if the customer entered one. An optional field. | [optional] |
| **storage_facility** | [**UpdateMTOShipmentStorageFacilityV3**](UpdateMTOShipmentStorageFacilityV3.md) |  | [optional] |
| **shipment_type** | [**MTOShipmentTypeV3**](MTOShipmentTypeV3.md) |  | [optional] |
| **diversion** | **Boolean** | This value indicates whether or not this shipment is part of a diversion. If yes, the shipment can be either the starting or ending segment of the diversion.  | [optional] |
| **point_of_contact** | **String** | Email or ID of the person who will be contacted in the event of questions or concerns about this update. May be the person performing the update, or someone else working with the Prime contractor.  | [optional] |
| **counselor_remarks** | **String** |  | [optional] |
| **ppm_shipment** | [**UpdatePPMShipmentV3**](UpdatePPMShipmentV3.md) |  | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::UpdateMTOShipmentV3.new(
  actual_pro_gear_weight: 4500,
  actual_spouse_pro_gear_weight: 4500,
  scheduled_pickup_date: null,
  actual_pickup_date: null,
  first_available_delivery_date: null,
  scheduled_delivery_date: null,
  actual_delivery_date: null,
  prime_estimated_weight: 4500,
  prime_actual_weight: 4500,
  nts_recorded_weight: 4500,
  pickup_address: null,
  destination_address: null,
  destination_type: null,
  secondary_pickup_address: null,
  secondary_delivery_address: null,
  storage_facility: null,
  shipment_type: null,
  diversion: null,
  point_of_contact: null,
  counselor_remarks: counselor approved,
  ppm_shipment: null
)
```

