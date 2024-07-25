# OpenapiClient::UpdateShipmentDestinationAddressV2

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **new_address** | [**AddressV2**](AddressV2.md) |  |  |
| **contractor_remarks** | **String** | This is the remark the Prime has entered, which would be the reason there is an address change. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::UpdateShipmentDestinationAddressV2.new(
  new_address: null,
  contractor_remarks: Customer reached out to me this week and let me know they want to move somewhere else.
)
```

