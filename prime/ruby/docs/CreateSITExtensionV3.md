# OpenapiClient::CreateSITExtensionV3

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **request_reason** | **String** |  |  |
| **contractor_remarks** | **String** |  |  |
| **requested_days** | **Integer** |  |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::CreateSITExtensionV3.new(
  request_reason: null,
  contractor_remarks: We need SIT additional days. The customer has not found a house yet.,
  requested_days: 30
)
```

