# OpenapiClient::MTOShipmentV3

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mto_service_items** | [**Array&lt;MTOServiceItemV3&gt;**](MTOServiceItemV3.md) | A list of service items connected to this shipment. | [optional][readonly] |
| **id** | **String** | The ID of the shipment. | [optional][readonly] |
| **move_task_order_id** | **String** | The ID of the move for this shipment. | [optional][readonly] |
| **approved_date** | **Date** | The date when the Task Ordering Officer first approved this shipment for the move. | [optional][readonly] |
| **requested_pickup_date** | **Date** | The date the customer selects during onboarding as their preferred pickup date. Other dates, such as required delivery date and (outside MilMove) the pack date, are derived from this date.  | [optional][readonly] |
| **requested_delivery_date** | **Date** | The customer&#39;s preferred delivery date. | [optional][readonly] |
| **scheduled_pickup_date** | **Date** | The date the Prime contractor scheduled to pick up this shipment after consultation with the customer. | [optional] |
| **actual_pickup_date** | **Date** | The date when the Prime contractor actually picked up the shipment. Updated after-the-fact. | [optional] |
| **first_available_delivery_date** | **Date** | The date the Prime provides to the customer as the first possible delivery date so that they can plan their travel accordingly.  | [optional] |
| **required_delivery_date** | **Date** | The latest date by which the Prime can deliver a customer&#39;s shipment without violating the contract. This is calculated based on weight, distance, and the scheduled pickup date. It cannot be modified.  | [optional][readonly] |
| **scheduled_delivery_date** | **Date** | The date the Prime contractor scheduled to deliver this shipment after consultation with the customer. | [optional] |
| **actual_delivery_date** | **Date** | The date when the Prime contractor actually delivered the shipment. Updated after-the-fact. | [optional] |
| **prime_estimated_weight** | **Integer** | The estimated weight of this shipment, determined by the movers during the pre-move survey. This value **can only be updated once.** If there was an issue with estimating the weight and a mistake was made, the Prime contracter will need to contact the TOO to change it.  | [optional] |
| **prime_estimated_weight_recorded_date** | **Date** | The date when the Prime contractor recorded the shipment&#39;s estimated weight. | [optional][readonly] |
| **prime_actual_weight** | **Integer** | The actual weight of the shipment, provided after the Prime packs, picks up, and weighs a customer&#39;s shipment. | [optional] |
| **nts_recorded_weight** | **Integer** | The previously recorded weight for the NTS Shipment. Used for NTS Release to know what the previous primeActualWeight or billable weight was. | [optional] |
| **customer_remarks** | **String** | The customer can use the customer remarks field to inform the services counselor and the movers about any special circumstances for this shipment. Typical examples:   * bulky or fragile items,   * weapons,   * access info for their address.  Customer enters this information during onboarding. Optional field.  | [optional][readonly] |
| **counselor_remarks** | **String** | The counselor can use the counselor remarks field to inform the movers about any special circumstances for this shipment. Typical examples:   * bulky or fragile items,   * weapons,   * access info for their address.  Counselors enters this information when creating or editing an MTO Shipment. Optional field.  | [optional][readonly] |
| **agents** | [**Array&lt;MTOAgentV3&gt;**](MTOAgentV3.md) | A list of the agents for a shipment. Agents are the people who the Prime contractor recognize as permitted to release (in the case of pickup) or receive (on delivery) a shipment.  | [optional] |
| **sit_extensions** | [**Array&lt;SITExtensionV3&gt;**](SITExtensionV3.md) |  | [optional] |
| **reweigh** | [**ReweighV3**](ReweighV3.md) |  | [optional] |
| **pickup_address** | [**AddressV3**](AddressV3.md) | The address where the movers should pick up this shipment, entered by the customer during onboarding when they enter shipment details.  | [optional] |
| **destination_address** | [**AddressV3**](AddressV3.md) | Where the movers should deliver this shipment. Often provided by the customer when they enter shipment details during onboarding, if they know their new address already.  May be blank when entered by the customer, required when entered by the Prime. May not represent the true final destination due to the shipment being diverted or placed in SIT.  | [optional] |
| **destination_type** | [**DestinationTypeV3**](DestinationTypeV3.md) |  | [optional] |
| **secondary_pickup_address** | [**AddressV3**](AddressV3.md) | A second pickup address for this shipment, if the customer entered one. An optional field. | [optional] |
| **secondary_delivery_address** | [**AddressV3**](AddressV3.md) | A second delivery address for this shipment, if the customer entered one. An optional field. | [optional] |
| **storage_facility** | [**UpdateMTOShipmentStorageFacilityV3**](UpdateMTOShipmentStorageFacilityV3.md) |  | [optional] |
| **shipment_type** | [**MTOShipmentTypeV3**](MTOShipmentTypeV3.md) |  | [optional] |
| **diversion** | **Boolean** | This value indicates whether or not this shipment is part of a diversion. If yes, the shipment can be either the starting or ending segment of the diversion.  | [optional] |
| **diversion_reason** | **String** | The reason the TOO provided when requesting a diversion for this shipment.  | [optional][readonly] |
| **status** | **String** | The status of a shipment, indicating where it is in the TOO&#39;s approval process. Can only be updated by the contractor in special circumstances.  | [optional][readonly] |
| **ppm_shipment** | [**PPMShipmentV3**](PPMShipmentV3.md) |  | [optional] |
| **delivery_address_update** | [**ShipmentAddressUpdateV3**](ShipmentAddressUpdateV3.md) |  | [optional] |
| **e_tag** | **String** | A hash unique to this shipment that should be used as the \&quot;If-Match\&quot; header for any updates. | [optional][readonly] |
| **created_at** | **Time** |  | [optional][readonly] |
| **updated_at** | **Time** |  | [optional][readonly] |
| **point_of_contact** | **String** | Email or ID of the person who will be contacted in the event of questions or concerns about this update. May be the person performing the update, or someone else working with the Prime contractor.  | [optional] |
| **origin_sit_auth_end_date** | **Date** | The SIT authorized end date for origin SIT. | [optional] |
| **destination_sit_auth_end_date** | **Date** | The SIT authorized end date for destination SIT. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::MTOShipmentV3.new(
  mto_service_items: null,
  id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  move_task_order_id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  approved_date: null,
  requested_pickup_date: null,
  requested_delivery_date: null,
  scheduled_pickup_date: null,
  actual_pickup_date: null,
  first_available_delivery_date: null,
  required_delivery_date: null,
  scheduled_delivery_date: null,
  actual_delivery_date: null,
  prime_estimated_weight: 4500,
  prime_estimated_weight_recorded_date: null,
  prime_actual_weight: 4500,
  nts_recorded_weight: 4500,
  customer_remarks: handle with care,
  counselor_remarks: handle with care,
  agents: null,
  sit_extensions: null,
  reweigh: null,
  pickup_address: null,
  destination_address: null,
  destination_type: null,
  secondary_pickup_address: null,
  secondary_delivery_address: null,
  storage_facility: null,
  shipment_type: null,
  diversion: null,
  diversion_reason: null,
  status: null,
  ppm_shipment: null,
  delivery_address_update: null,
  e_tag: null,
  created_at: null,
  updated_at: null,
  point_of_contact: null,
  origin_sit_auth_end_date: null,
  destination_sit_auth_end_date: null
)
```

