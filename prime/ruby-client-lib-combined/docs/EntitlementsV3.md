# OpenapiClient::EntitlementsV3

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **authorized_weight** | **Integer** |  | [optional] |
| **dependents_authorized** | **Boolean** |  | [optional] |
| **gun_safe** | **Boolean** |  | [optional] |
| **non_temporary_storage** | **Boolean** |  | [optional] |
| **privately_owned_vehicle** | **Boolean** |  | [optional] |
| **pro_gear_weight** | **Integer** |  | [optional] |
| **pro_gear_weight_spouse** | **Integer** |  | [optional] |
| **required_medical_equipment_weight** | **Integer** |  | [optional] |
| **organizational_clothing_and_individual_equipment** | **Boolean** |  | [optional] |
| **storage_in_transit** | **Integer** |  | [optional] |
| **total_weight** | **Integer** |  | [optional] |
| **total_dependents** | **Integer** |  | [optional] |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::EntitlementsV3.new(
  id: 571008b1-b0de-454d-b843-d71be9f02c04,
  authorized_weight: 2000,
  dependents_authorized: true,
  gun_safe: false,
  non_temporary_storage: false,
  privately_owned_vehicle: false,
  pro_gear_weight: 2000,
  pro_gear_weight_spouse: 500,
  required_medical_equipment_weight: 500,
  organizational_clothing_and_individual_equipment: false,
  storage_in_transit: 90,
  total_weight: 500,
  total_dependents: 2,
  e_tag: null
)
```

