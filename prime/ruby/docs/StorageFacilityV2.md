# OpenapiClient::StorageFacilityV2

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **facility_name** | **String** |  | [optional] |
| **address** | [**AddressV2**](AddressV2.md) |  | [optional] |
| **lot_number** | **String** |  | [optional] |
| **phone** | **String** |  | [optional] |
| **email** | **String** |  | [optional] |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::StorageFacilityV2.new(
  id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  facility_name: null,
  address: null,
  lot_number: null,
  phone: null,
  email: null,
  e_tag: null
)
```

