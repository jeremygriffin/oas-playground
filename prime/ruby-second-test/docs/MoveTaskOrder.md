# OpenapiClient::MoveTaskOrder

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **move_code** | **String** |  | [optional][readonly] |
| **created_at** | **Time** |  | [optional][readonly] |
| **order_id** | **String** |  | [optional] |
| **order** | [**Order**](Order.md) |  | [optional] |
| **reference_id** | **String** |  | [optional] |
| **available_to_prime_at** | **Time** |  | [optional][readonly] |
| **updated_at** | **Time** |  | [optional][readonly] |
| **prime_counseling_completed_at** | **Time** |  | [optional][readonly] |
| **payment_requests** | [**Array&lt;PaymentRequest&gt;**](PaymentRequest.md) |  |  |
| **mto_service_items** | [**Array&lt;MTOServiceItem&gt;**](MTOServiceItem.md) |  |  |
| **mto_shipments** | [**Array&lt;MTOShipmentWithoutServiceItems&gt;**](MTOShipmentWithoutServiceItems.md) | A list of shipments without their associated service items. |  |
| **ppm_type** | **String** |  | [optional] |
| **ppm_estimated_weight** | **Integer** |  | [optional] |
| **excess_weight_qualified_at** | **Time** |  | [optional][readonly] |
| **excess_weight_acknowledged_at** | **Time** |  | [optional][readonly] |
| **excess_weight_upload_id** | **String** |  | [optional][readonly] |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::MoveTaskOrder.new(
  id: a502b4f1-b9c4-4faf-8bdd-68292501bf26,
  move_code: HYXFJF,
  created_at: null,
  order_id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  order: null,
  reference_id: 1001-3456,
  available_to_prime_at: null,
  updated_at: null,
  prime_counseling_completed_at: null,
  payment_requests: null,
  mto_service_items: null,
  mto_shipments: null,
  ppm_type: null,
  ppm_estimated_weight: null,
  excess_weight_qualified_at: null,
  excess_weight_acknowledged_at: null,
  excess_weight_upload_id: null,
  e_tag: null
)
```

