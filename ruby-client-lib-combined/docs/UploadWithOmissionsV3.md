# OpenapiClient::UploadWithOmissionsV3

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **url** | **String** |  | [optional] |
| **filename** | **String** |  |  |
| **content_type** | **String** |  |  |
| **bytes** | **Integer** |  |  |
| **status** | **String** |  | [optional] |
| **created_at** | **Time** |  | [optional][readonly] |
| **updated_at** | **Time** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::UploadWithOmissionsV3.new(
  id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  url: https://uploads.domain.test/dir/c56a4180-65aa-42ec-a945-5fd21dec0538,
  filename: filename.pdf,
  content_type: application/pdf,
  bytes: null,
  status: null,
  created_at: null,
  updated_at: null
)
```

