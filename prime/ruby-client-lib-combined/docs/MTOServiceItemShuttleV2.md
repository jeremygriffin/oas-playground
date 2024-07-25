# OpenapiClient::MTOServiceItemShuttleV2

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **re_service_code** | **String** | A unique code for the service item. Indicates if shuttling is requested for the shipment origin (&#x60;DOSHUT&#x60;) or destination (&#x60;DDSHUT&#x60;).  |  |
| **reason** | **String** | The contractor&#39;s explanation for why a shuttle service is requested. Used by the TOO while deciding to approve or reject the service item.  |  |
| **estimated_weight** | **Integer** | An estimate of how much weight from a shipment will be included in the shuttling service. | [optional] |
| **actual_weight** | **Integer** | A record of the actual weight that was shuttled. Provided by the movers, based on weight tickets. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::MTOServiceItemShuttleV2.new(
  re_service_code: null,
  reason: Storage items need to be picked up.,
  estimated_weight: 4200,
  actual_weight: 4000
)
```

