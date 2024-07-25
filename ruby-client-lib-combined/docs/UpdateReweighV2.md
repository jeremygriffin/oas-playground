# OpenapiClient::UpdateReweighV2

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **weight** | **Integer** | The total reweighed weight for the shipment in pounds. | [optional] |
| **verification_reason** | **String** | In lieu of a document being uploaded indicating why a reweigh did not occur. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::UpdateReweighV2.new(
  weight: 2000,
  verification_reason: The reweigh was not performed because the shipment was already delivered
)
```

