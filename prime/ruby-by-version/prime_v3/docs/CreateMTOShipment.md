# OpenapiClient::CreateMTOShipment

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **move_task_order_id** | **String** | The ID of the move this new shipment is for. |  |
| **requested_pickup_date** | **Date** | The customer&#39;s preferred pickup date. Other dates, such as required delivery date and (outside MilMove) the pack date, are derived from this date.  | [optional] |
| **prime_estimated_weight** | **Integer** | The estimated weight of this shipment, determined by the movers during the pre-move survey. This value **can only be updated once.** If there was an issue with estimating the weight and a mistake was made, the Prime contractor will need to contact the TOO to change it.  | [optional] |
| **customer_remarks** | **String** | The customer can use the customer remarks field to inform the services counselor and the movers about any special circumstances for this shipment. Typical examples:   * bulky or fragile items,   * weapons,   * access info for their address.  Customer enters this information during onboarding. Optional field.  | [optional] |
| **agents** | [**Array&lt;MTOAgent&gt;**](MTOAgent.md) | A list of the agents for a shipment. Agents are the people who the Prime contractor recognize as permitted to release (in the case of pickup) or receive (on delivery) a shipment.  | [optional] |
| **mto_service_items** | [**Array&lt;MTOServiceItem&gt;**](MTOServiceItem.md) | A list of service items connected to this shipment. | [optional] |
| **pickup_address** | [**Address**](Address.md) | The address where the movers should pick up this shipment. | [optional] |
| **destination_address** | [**Address**](Address.md) | Where the movers should deliver this shipment. | [optional] |
| **shipment_type** | [**MTOShipmentType**](MTOShipmentType.md) |  |  |
| **diversion** | **Boolean** | This value indicates whether or not this shipment is part of a diversion. If yes, the shipment can be either the starting or ending segment of the diversion. When this boolean is true, you must link it to a parent shipment with the divertedFromShipmentId parameter.  | [optional] |
| **diverted_from_shipment_id** | **String** | The ID of the shipment this is a diversion from. Aka the \&quot;Parent\&quot; shipment. The diversion boolean must be true if this parameter is supplied in the request. If provided, and if the diverted from shipment is also a diversion, the previous should must then also have a parent ID.  | [optional] |
| **point_of_contact** | **String** | Email or ID of the person who will be contacted in the event of questions or concerns about this update. May be the person performing the update, or someone else working with the Prime contractor.  | [optional] |
| **counselor_remarks** | **String** |  | [optional] |
| **ppm_shipment** | [**CreatePPMShipment**](CreatePPMShipment.md) |  | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::CreateMTOShipment.new(
  move_task_order_id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  requested_pickup_date: null,
  prime_estimated_weight: 4500,
  customer_remarks: handle with care,
  agents: null,
  mto_service_items: null,
  pickup_address: null,
  destination_address: null,
  shipment_type: null,
  diversion: null,
  diverted_from_shipment_id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  point_of_contact: null,
  counselor_remarks: counselor approved,
  ppm_shipment: null
)
```

