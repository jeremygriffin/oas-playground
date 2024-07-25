# OpenapiClient::CreateSITAddressUpdateRequestV3

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **new_address** | [**AddressV3**](AddressV3.md) |  | [optional] |
| **contractor_remarks** | **String** |  |  |
| **mto_service_item_id** | **String** |  | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::CreateSITAddressUpdateRequestV3.new(
  new_address: null,
  contractor_remarks: Customer reached out to me this week &amp; let me know they want to move closer to family.,
  mto_service_item_id: c56a4180-65aa-42ec-a945-5fd21dec0538
)
```

