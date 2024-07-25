# OpenapiClient::MTOServiceItemDomesticCratingV3

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **re_service_code** | **String** | A unique code for the service item. Indicates if the service is for crating (DCRT) or uncrating (DUCRT). |  |
| **item** | **Object** | The dimensions of the item being crated. |  |
| **crate** | **Object** | The dimensions for the crate the item will be shipped in. |  |
| **description** | **String** | A description of the item being crated. |  |
| **reason** | **String** | The contractor&#39;s explanation for why an item needed to be crated or uncrated. Used by the TOO while deciding to approve or reject the service item.  | [optional] |
| **standalone_crate** | **Boolean** |  | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::MTOServiceItemDomesticCratingV3.new(
  re_service_code: null,
  item: null,
  crate: null,
  description: Decorated horse head to be crated.,
  reason: Storage items need to be picked up,
  standalone_crate: null
)
```

