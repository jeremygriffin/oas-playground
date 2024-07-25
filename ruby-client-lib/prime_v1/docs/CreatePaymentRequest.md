# OpenapiClient::CreatePaymentRequest

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **is_final** | **Boolean** |  | [optional][default to false] |
| **move_task_order_id** | **String** |  |  |
| **service_items** | [**Array&lt;ServiceItem&gt;**](ServiceItem.md) |  |  |
| **point_of_contact** | **String** | Email or id of a contact person for this update. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::CreatePaymentRequest.new(
  is_final: null,
  move_task_order_id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  service_items: null,
  point_of_contact: null
)
```

