# OpenapiClient::UpdateMTOServiceItemShuttle

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **actual_weight** | **Integer** | Provided by the movers, based on weight tickets. Relevant for shuttling (DDSHUT &amp; DOSHUT) service items. | [optional] |
| **estimated_weight** | **Integer** | An estimate of how much weight from a shipment will be included in a shuttling (DDSHUT &amp; DOSHUT) service item. | [optional] |
| **re_service_code** | **String** | Service code allowed for this model type. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::UpdateMTOServiceItemShuttle.new(
  actual_weight: 4000,
  estimated_weight: 4200,
  re_service_code: null
)
```

