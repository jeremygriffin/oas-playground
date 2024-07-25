# OpenapiClient::PaymentRequestV3

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional][readonly] |
| **is_final** | **Boolean** |  | [optional][default to false] |
| **move_task_order_id** | **String** |  | [optional] |
| **rejection_reason** | **String** |  | [optional] |
| **status** | [**PaymentRequestStatusV3**](PaymentRequestStatusV3.md) |  | [optional] |
| **payment_request_number** | **String** |  | [optional][readonly] |
| **recalculation_of_payment_request_id** | **String** |  | [optional][readonly] |
| **proof_of_service_docs** | [**Array&lt;ProofOfServiceDocV3&gt;**](ProofOfServiceDocV3.md) |  | [optional] |
| **payment_service_items** | [**Array&lt;PaymentServiceItemV3&gt;**](PaymentServiceItemV3.md) |  | [optional] |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::PaymentRequestV3.new(
  id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  is_final: null,
  move_task_order_id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  rejection_reason: documentation was incomplete,
  status: null,
  payment_request_number: 1234-5678-1,
  recalculation_of_payment_request_id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  proof_of_service_docs: null,
  payment_service_items: null,
  e_tag: null
)
```

