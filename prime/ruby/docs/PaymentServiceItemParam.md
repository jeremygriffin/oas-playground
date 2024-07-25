# OpenapiClient::PaymentServiceItemParam

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional][readonly] |
| **payment_service_item_id** | **String** |  | [optional] |
| **key** | [**ServiceItemParamName**](ServiceItemParamName.md) |  | [optional] |
| **value** | **String** |  | [optional] |
| **type** | [**ServiceItemParamType**](ServiceItemParamType.md) |  | [optional] |
| **origin** | [**ServiceItemParamOrigin**](ServiceItemParamOrigin.md) |  | [optional] |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::PaymentServiceItemParam.new(
  id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  payment_service_item_id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  key: null,
  value: 3025,
  type: null,
  origin: null,
  e_tag: null
)
```

