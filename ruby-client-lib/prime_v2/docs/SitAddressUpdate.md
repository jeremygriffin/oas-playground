# OpenapiClient::SitAddressUpdate

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional][readonly] |
| **mto_service_item_id** | **String** |  | [optional][readonly] |
| **new_address_id** | **String** |  | [optional][readonly] |
| **new_address** | [**Address**](Address.md) |  | [optional] |
| **old_address_id** | **String** |  | [optional][readonly] |
| **old_address** | [**Address**](Address.md) |  | [optional] |
| **status** | [**SitAddressUpdateStatus**](SitAddressUpdateStatus.md) |  | [optional] |
| **distance** | **Integer** |  | [optional][readonly] |
| **contractor_remarks** | **String** |  | [optional] |
| **office_remarks** | **String** |  | [optional] |
| **created_at** | **Time** |  | [optional][readonly] |
| **updated_at** | **Time** |  | [optional][readonly] |
| **e_tag** | **String** | A hash unique to this shipment that should be used as the \&quot;If-Match\&quot; header for any updates. | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::SitAddressUpdate.new(
  id: ddd7bb48-4730-47c4-9781-6500384f4941,
  mto_service_item_id: 12d9e103-5a56-4636-906d-6e993b97ef51,
  new_address_id: 31a2ad3c-1682-4d5b-8423-ff40053a056b,
  new_address: null,
  old_address_id: 31a2ad3c-1682-4d5b-8423-ff40053a056b,
  old_address: null,
  status: null,
  distance: 25,
  contractor_remarks: Customer reached out to me this week &amp; let me know they want to move closer to family.,
  office_remarks: The customer has found a new house closer to base.,
  created_at: null,
  updated_at: null,
  e_tag: null
)
```

