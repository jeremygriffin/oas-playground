# OpenapiClient::OrderV3

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **customer** | [**CustomerV3**](CustomerV3.md) |  | [optional] |
| **customer_id** | **String** |  | [optional] |
| **entitlement** | [**EntitlementsV3**](EntitlementsV3.md) |  | [optional] |
| **destination_duty_location** | [**DutyLocationV3**](DutyLocationV3.md) |  | [optional] |
| **origin_duty_location** | [**DutyLocationV3**](DutyLocationV3.md) |  | [optional] |
| **origin_duty_location_gbloc** | **String** |  | [optional] |
| **rank** | **String** |  |  |
| **report_by_date** | **Date** |  | [optional] |
| **orders_type** | [**OrdersTypeV3**](OrdersTypeV3.md) |  | [optional] |
| **order_number** | **String** |  |  |
| **lines_of_accounting** | **String** |  |  |
| **supply_and_services_cost_estimate** | **String** |  | [optional][readonly] |
| **packing_and_shipping_instructions** | **String** |  | [optional][readonly] |
| **method_of_payment** | **String** |  | [optional][readonly] |
| **naics** | **String** |  | [optional][readonly] |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::OrderV3.new(
  id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  customer: null,
  customer_id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  entitlement: null,
  destination_duty_location: null,
  origin_duty_location: null,
  origin_duty_location_gbloc: KKFA,
  rank: E_5,
  report_by_date: null,
  orders_type: null,
  order_number: null,
  lines_of_accounting: null,
  supply_and_services_cost_estimate: null,
  packing_and_shipping_instructions: null,
  method_of_payment: null,
  naics: null,
  e_tag: null
)
```

