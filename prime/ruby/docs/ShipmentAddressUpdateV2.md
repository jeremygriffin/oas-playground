# OpenapiClient::ShipmentAddressUpdateV2

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [readonly] |
| **contractor_remarks** | **String** | The reason there is an address change. | [readonly] |
| **office_remarks** | **String** | The TOO comment on approval or rejection. | [optional] |
| **status** | [**ShipmentAddressUpdateStatusV2**](ShipmentAddressUpdateStatusV2.md) |  |  |
| **shipment_id** | **String** |  | [readonly] |
| **original_address** | [**AddressV2**](AddressV2.md) |  |  |
| **new_address** | [**AddressV2**](AddressV2.md) |  |  |
| **sit_original_address** | [**AddressV2**](AddressV2.md) |  | [optional] |
| **old_sit_distance_between** | **Integer** | The distance between the original SIT address and the previous/old destination address of shipment | [optional] |
| **new_sit_distance_between** | **Integer** | The distance between the original SIT address and requested new destination address of shipment | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::ShipmentAddressUpdateV2.new(
  id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  contractor_remarks: This is a contractor remark,
  office_remarks: This is an office remark,
  status: null,
  shipment_id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  original_address: null,
  new_address: null,
  sit_original_address: null,
  old_sit_distance_between: 50,
  new_sit_distance_between: 88
)
```

