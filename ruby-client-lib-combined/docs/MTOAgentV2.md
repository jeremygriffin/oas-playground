# OpenapiClient::MTOAgentV2

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** | The ID of the agent. | [optional][readonly] |
| **mto_shipment_id** | **String** | The ID of the shipment this agent is permitted to release/receive. | [optional][readonly] |
| **created_at** | **Time** |  | [optional][readonly] |
| **updated_at** | **Time** |  | [optional][readonly] |
| **first_name** | **String** |  | [optional] |
| **last_name** | **String** |  | [optional] |
| **email** | **String** |  | [optional] |
| **phone** | **String** |  | [optional] |
| **agent_type** | [**MTOAgentTypeV2**](MTOAgentTypeV2.md) |  | [optional] |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::MTOAgentV2.new(
  id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  mto_shipment_id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  created_at: null,
  updated_at: null,
  first_name: null,
  last_name: null,
  email: null,
  phone: null,
  agent_type: null,
  e_tag: null
)
```

