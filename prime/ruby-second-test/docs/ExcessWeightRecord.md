# OpenapiClient::ExcessWeightRecord

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
| **move_id** | **String** | The UUID of the move this excess weight record belongs to. |  |
| **move_excess_weight_qualified_at** | **Time** | The date and time when the sum of all the move&#39;s shipments met the excess weight qualification threshold. The system monitors these weights and will update this field automatically.  | [optional][readonly] |
| **move_excess_weight_acknowledged_at** | **Time** | The date and time when the TOO acknowledged the excess weight alert, either by dismissing the risk or updating the max billable weight. This will occur after the excess weight record has been uploaded.  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::ExcessWeightRecord.new(
  id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  url: https://uploads.domain.test/dir/c56a4180-65aa-42ec-a945-5fd21dec0538,
  filename: filename.pdf,
  content_type: application/pdf,
  bytes: null,
  status: null,
  created_at: null,
  updated_at: null,
  move_id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  move_excess_weight_qualified_at: null,
  move_excess_weight_acknowledged_at: null
)
```

