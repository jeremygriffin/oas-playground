# OpenapiClient::UpdateMTOShipmentStorageFacility

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **facility_name** | **String** |  | [optional] |
| **address** | [**Address**](Address.md) |  | [optional] |
| **lot_number** | **String** |  | [optional] |
| **phone** | **String** |  | [optional] |
| **email** | **String** |  | [optional] |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::UpdateMTOShipmentStorageFacility.new(
  id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  facility_name: null,
  address: null,
  lot_number: null,
  phone: null,
  email: null,
  e_tag: null
)
```

