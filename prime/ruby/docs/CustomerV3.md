# OpenapiClient::CustomerV3

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **dod_id** | **String** |  | [optional] |
| **emplid** | **String** |  | [optional] |
| **user_id** | **String** |  | [optional] |
| **current_address** | [**AddressV3**](AddressV3.md) |  | [optional] |
| **first_name** | **String** |  | [optional] |
| **last_name** | **String** |  | [optional] |
| **branch** | **String** |  | [optional] |
| **phone** | **String** |  | [optional] |
| **email** | **String** |  | [optional] |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::CustomerV3.new(
  id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  dod_id: null,
  emplid: null,
  user_id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  current_address: null,
  first_name: Vanya,
  last_name: Petrovna,
  branch: COAST_GUARD,
  phone: null,
  email: fake@example.com,
  e_tag: null
)
```

