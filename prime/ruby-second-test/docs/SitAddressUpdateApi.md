# OpenapiClient::SitAddressUpdateApi

All URIs are relative to */prime/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**create_sit_address_update_request**](SitAddressUpdateApi.md#create_sit_address_update_request) | **POST** /sit-address-updates | createSITAddressUpdateRequest |


## create_sit_address_update_request

> <SitAddressUpdate> create_sit_address_update_request(opts)

createSITAddressUpdateRequest

**Functionality:** Creates an update request for a SIT service item's final delivery address. A newly created update request is assigned the status 'REQUESTED'  if the change in address is > 50 miles and automatically approved otherwise.  **Limitations:** The update can be requested for APPROVED SIT service items only. Only ONE request is allowed per approved SIT service item.  **DEPRECATION ON AUGUST 5TH, 2024** Following deprecation, when updating a service item's final delivery address, you will need to update the shipment's destination address. This will update the destination SIT service items' final delivery address upon approval. For `APPROVED` shipments, you can use [updateShipmentDestinationAddress](#mtoShipment/updateShipmentDestinationAddress) For shipments in any other status, you can use [updateMTOShipmentAddress](#mtoShipment/updateMTOShipmentAddress) 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::SitAddressUpdateApi.new
opts = {
  body: OpenapiClient::CreateSITAddressUpdateRequest.new({contractor_remarks: 'Customer reached out to me this week & let me know they want to move closer to family.'}) # CreateSITAddressUpdateRequest | 
}

begin
  # createSITAddressUpdateRequest
  result = api_instance.create_sit_address_update_request(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling SitAddressUpdateApi->create_sit_address_update_request: #{e}"
end
```

#### Using the create_sit_address_update_request_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SitAddressUpdate>, Integer, Hash)> create_sit_address_update_request_with_http_info(opts)

```ruby
begin
  # createSITAddressUpdateRequest
  data, status_code, headers = api_instance.create_sit_address_update_request_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SitAddressUpdate>
rescue OpenapiClient::ApiError => e
  puts "Error when calling SitAddressUpdateApi->create_sit_address_update_request_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **body** | [**CreateSITAddressUpdateRequest**](CreateSITAddressUpdateRequest.md) |  | [optional] |

### Return type

[**SitAddressUpdate**](SitAddressUpdate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

