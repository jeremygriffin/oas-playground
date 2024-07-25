# OpenapiClient::Order

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **customer** | [**Customer**](Customer.md) |  | [optional] |
| **customer_id** | **String** |  | [optional] |
| **entitlement** | [**Entitlements**](Entitlements.md) |  | [optional] |
| **destination_duty_location** | [**DutyLocation**](DutyLocation.md) |  | [optional] |
| **origin_duty_location** | [**DutyLocation**](DutyLocation.md) |  | [optional] |
| **origin_duty_location_gbloc** | **String** |  | [optional] |
| **rank** | **String** |  |  |
| **report_by_date** | **Date** |  | [optional] |
| **orders_type** | [**OrdersType**](OrdersType.md) |  | [optional] |
| **order_number** | **String** |  |  |
| **lines_of_accounting** | **String** |  |  |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::Order.new(
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
  e_tag: null
)
```

