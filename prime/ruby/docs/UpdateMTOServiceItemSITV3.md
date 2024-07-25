# OpenapiClient::UpdateMTOServiceItemSITV3

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **re_service_code** | **String** | Service code allowed for this model type. | [optional] |
| **sit_departure_date** | **Date** | Departure date for SIT. This is the end date of the SIT at either origin or destination. | [optional] |
| **sit_destination_final_address** | [**AddressV3**](AddressV3.md) |  | [optional] |
| **date_of_contact1** | **Date** | Date of attempted contact by the prime corresponding to &#39;timeMilitary1&#39;. | [optional] |
| **time_military1** | **String** | Time of attempted contact by the prime corresponding to &#39;dateOfContact1&#39;, in military format. | [optional] |
| **first_available_delivery_date1** | **Date** | First available date that Prime can deliver SIT service item. | [optional] |
| **date_of_contact2** | **Date** | Date of attempted contact by the prime corresponding to &#39;timeMilitary2&#39;. | [optional] |
| **time_military2** | **String** | Time of attempted contact by the prime corresponding to &#39;dateOfContact2&#39;, in military format. | [optional] |
| **first_available_delivery_date2** | **Date** | Second available date that Prime can deliver SIT service item. | [optional] |
| **sit_requested_delivery** | **Date** | Date when the customer has requested delivery out of SIT. | [optional] |
| **sit_customer_contacted** | **Date** | Date when the customer contacted the prime for a delivery out of SIT. | [optional] |
| **update_reason** | **String** | Reason for updating service item. | [optional] |
| **sit_postal_code** | **String** |  | [optional] |
| **sit_entry_date** | **Date** | Entry date for the SIT. | [optional] |
| **request_approvals_requested_status** | **Boolean** | Indicates if \&quot;Approvals Requested\&quot; status is being requested. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::UpdateMTOServiceItemSITV3.new(
  re_service_code: null,
  sit_departure_date: null,
  sit_destination_final_address: null,
  date_of_contact1: null,
  time_military1: 1400Z,
  first_available_delivery_date1: null,
  date_of_contact2: null,
  time_military2: 1400Z,
  first_available_delivery_date2: null,
  sit_requested_delivery: null,
  sit_customer_contacted: null,
  update_reason: null,
  sit_postal_code: 90210,
  sit_entry_date: null,
  request_approvals_requested_status: null
)
```

