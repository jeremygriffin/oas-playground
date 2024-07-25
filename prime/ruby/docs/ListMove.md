# OpenapiClient::ListMove

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **move_code** | **String** |  | [optional][readonly] |
| **created_at** | **Time** |  | [optional][readonly] |
| **order_id** | **String** |  | [optional] |
| **reference_id** | **String** |  | [optional] |
| **available_to_prime_at** | **Time** |  | [optional][readonly] |
| **updated_at** | **Time** |  | [optional][readonly] |
| **ppm_type** | **String** |  | [optional] |
| **e_tag** | **String** |  | [optional][readonly] |
| **amendments** | [**Amendments**](Amendments.md) |  | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::ListMove.new(
  id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  move_code: HYXFJF,
  created_at: null,
  order_id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  reference_id: 1001-3456,
  available_to_prime_at: null,
  updated_at: null,
  ppm_type: null,
  e_tag: null,
  amendments: null
)
```

