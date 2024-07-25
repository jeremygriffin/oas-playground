# OpenapiClient::Reweigh

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **requested_at** | **Time** |  | [optional] |
| **requested_by** | [**ReweighRequester**](ReweighRequester.md) |  | [optional] |
| **verification_provided_at** | **Time** |  | [optional] |
| **verification_reason** | **String** |  | [optional] |
| **weight** | **Integer** |  | [optional] |
| **shipment_id** | **String** |  | [optional] |
| **created_at** | **Time** |  | [optional][readonly] |
| **updated_at** | **Time** |  | [optional][readonly] |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::Reweigh.new(
  id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  requested_at: null,
  requested_by: null,
  verification_provided_at: null,
  verification_reason: The reweigh was not performed due to some justification provided by the Prime,
  weight: 2000,
  shipment_id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  created_at: null,
  updated_at: null,
  e_tag: null
)
```

