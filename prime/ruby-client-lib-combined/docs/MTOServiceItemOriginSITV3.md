# OpenapiClient::MTOServiceItemOriginSITV3

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **re_service_code** | **String** | Service code allowed for this model type. |  |
| **reason** | **String** | Explanation of why Prime is picking up SIT item. |  |
| **sit_postal_code** | **String** |  |  |
| **sit_entry_date** | **Date** | Entry date for the SIT |  |
| **sit_departure_date** | **Date** | Departure date for SIT. This is the end date of the SIT at either origin or destination. This is optional as it can be updated using the UpdateMTOServiceItemSIT modelType at a later date. | [optional] |
| **sit_hhg_actual_origin** | [**AddressV3**](AddressV3.md) |  | [optional] |
| **sit_hhg_original_origin** | [**AddressV3**](AddressV3.md) |  | [optional] |
| **request_approvals_requested_status** | **Boolean** |  | [optional] |
| **sit_requested_delivery** | **Date** | Date when the customer has requested delivery out of SIT. | [optional] |
| **sit_customer_contacted** | **Date** | Date when the customer contacted the prime for a delivery out of SIT. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::MTOServiceItemOriginSITV3.new(
  re_service_code: null,
  reason: Storage items need to be picked up,
  sit_postal_code: 90210,
  sit_entry_date: null,
  sit_departure_date: null,
  sit_hhg_actual_origin: null,
  sit_hhg_original_origin: null,
  request_approvals_requested_status: null,
  sit_requested_delivery: null,
  sit_customer_contacted: null
)
```

