# OpenapiClient::SITExtensionV2

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **mto_shipment_id** | **String** |  | [optional] |
| **request_reason** | **String** |  | [optional] |
| **contractor_remarks** | **String** |  | [optional] |
| **requested_days** | **Integer** |  | [optional] |
| **status** | **String** |  | [optional] |
| **approved_days** | **Integer** |  | [optional] |
| **decision_date** | **Time** |  | [optional] |
| **office_remarks** | **String** |  | [optional] |
| **created_at** | **Time** |  | [optional][readonly] |
| **updated_at** | **Time** |  | [optional][readonly] |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::SITExtensionV2.new(
  id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  mto_shipment_id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  request_reason: null,
  contractor_remarks: We need SIT additional days. The customer has not found a house yet.,
  requested_days: 30,
  status: null,
  approved_days: 30,
  decision_date: null,
  office_remarks: null,
  created_at: null,
  updated_at: null,
  e_tag: null
)
```

