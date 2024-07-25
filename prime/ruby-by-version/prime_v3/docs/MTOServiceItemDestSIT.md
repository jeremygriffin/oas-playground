# OpenapiClient::MTOServiceItemDestSIT

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **re_service_code** | **String** | Service code allowed for this model type. |  |
| **date_of_contact1** | **Date** | Date of attempted contact by the prime corresponding to &#x60;timeMilitary1&#x60;. | [optional] |
| **date_of_contact2** | **Date** | Date of attempted contact by the prime corresponding to &#x60;timeMilitary2&#x60;. | [optional] |
| **time_military1** | **String** | Time of attempted contact corresponding to &#x60;dateOfContact1&#x60;, in military format. | [optional] |
| **time_military2** | **String** | Time of attempted contact corresponding to &#x60;dateOfContact2&#x60;, in military format. | [optional] |
| **first_available_delivery_date1** | **Date** | First available date that Prime can deliver SIT service item. | [optional] |
| **first_available_delivery_date2** | **Date** | Second available date that Prime can deliver SIT service item. | [optional] |
| **sit_entry_date** | **Date** | Entry date for the SIT |  |
| **sit_departure_date** | **Date** | Departure date for SIT. This is the end date of the SIT at either origin or destination. This is optional as it can be updated using the UpdateMTOServiceItemSIT modelType at a later date. | [optional] |
| **sit_destination_final_address** | [**Address**](Address.md) |  | [optional] |
| **reason** | **String** | The reason item has been placed in SIT.  |  |
| **sit_address_updates** | [**Array&lt;SitAddressUpdate&gt;**](SitAddressUpdate.md) | A list of updates to a SIT service item address. | [optional] |
| **sit_requested_delivery** | **Date** | Date when the customer has requested delivery out of SIT. | [optional] |
| **sit_customer_contacted** | **Date** | Date when the customer contacted the prime for a delivery out of SIT. | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::MTOServiceItemDestSIT.new(
  re_service_code: null,
  date_of_contact1: null,
  date_of_contact2: null,
  time_military1: 1400Z,
  time_military2: 1400Z,
  first_available_delivery_date1: null,
  first_available_delivery_date2: null,
  sit_entry_date: null,
  sit_departure_date: null,
  sit_destination_final_address: null,
  reason: null,
  sit_address_updates: null,
  sit_requested_delivery: null,
  sit_customer_contacted: null
)
```

