# OpenapiClient::DutyLocation

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **name** | **String** |  | [optional] |
| **address_id** | **String** |  | [optional] |
| **address** | [**Address**](Address.md) |  | [optional] |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::DutyLocation.new(
  id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  name: Fort Bragg North Station,
  address_id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  address: null,
  e_tag: null
)
```

