# OpenapiClient::MTOServiceItemDimension

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **length** | **Integer** | Length in thousandth inches. 1000 thou &#x3D; 1 inch. |  |
| **width** | **Integer** | Width in thousandth inches. 1000 thou &#x3D; 1 inch. |  |
| **height** | **Integer** | Height in thousandth inches. 1000 thou &#x3D; 1 inch. |  |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::MTOServiceItemDimension.new(
  id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  length: 1000,
  width: 1000,
  height: 1000
)
```

