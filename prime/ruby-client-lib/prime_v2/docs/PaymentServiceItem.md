# OpenapiClient::PaymentServiceItem

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional][readonly] |
| **payment_request_id** | **String** |  | [optional] |
| **mto_service_item_id** | **String** |  | [optional] |
| **status** | [**PaymentServiceItemStatus**](PaymentServiceItemStatus.md) |  | [optional] |
| **price_cents** | **Integer** |  | [optional] |
| **rejection_reason** | **String** |  | [optional] |
| **reference_id** | **Object** |  | [optional][readonly] |
| **payment_service_item_params** | [**Array&lt;PaymentServiceItemParam&gt;**](PaymentServiceItemParam.md) |  | [optional] |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::PaymentServiceItem.new(
  id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  payment_request_id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  mto_service_item_id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  status: null,
  price_cents: null,
  rejection_reason: documentation was incomplete,
  reference_id: 1234-5678-c56a4180,
  payment_service_item_params: null,
  e_tag: null
)
```

