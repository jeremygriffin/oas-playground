# OpenapiClient::MTOServiceItemV3

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** | The ID of the service item. | [optional][readonly] |
| **move_task_order_id** | **String** | The ID of the move for this service item. |  |
| **mto_shipment_id** | **String** | The ID of the shipment this service is for, if any. Optional. | [optional] |
| **re_service_name** | **String** | The full descriptive name of the service. | [optional][readonly] |
| **status** | [**MTOServiceItemStatusV3**](MTOServiceItemStatusV3.md) |  | [optional] |
| **rejection_reason** | **String** | The reason why this service item was rejected by the TOO. | [optional][readonly] |
| **model_type** | [**MTOServiceItemModelTypeV3**](MTOServiceItemModelTypeV3.md) |  |  |
| **service_request_documents** | [**Array&lt;ServiceRequestDocumentV3&gt;**](ServiceRequestDocumentV3.md) |  | [optional] |
| **e_tag** | **String** | A hash unique to this service item that should be used as the \&quot;If-Match\&quot; header for any updates. | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::MTOServiceItemV3.new(
  id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  move_task_order_id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  mto_shipment_id: 1f2270c7-7166-40ae-981e-b200ebdf3054,
  re_service_name: null,
  status: null,
  rejection_reason: item was too heavy,
  model_type: null,
  service_request_documents: null,
  e_tag: null
)
```

