# OpenapiClient::MtoShipmentApi

All URIs are relative to */prime/v2*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**create_mto_shipment**](MtoShipmentApi.md#create_mto_shipment) | **POST** /mto-shipments | createMTOShipment |
| [**update_mto_shipment**](MtoShipmentApi.md#update_mto_shipment) | **PATCH** /mto-shipments/{mtoShipmentID} | updateMTOShipment |


## create_mto_shipment

> <MTOShipment> create_mto_shipment(opts)

createMTOShipment

Creates a new shipment within the specified move. This endpoint should be used whenever the movers identify a need for an additional shipment. The new shipment will be submitted to the TOO for review, and the TOO must approve it before the contractor can proceed with billing.  **NOTE**: When creating a child shipment diversion, you can no longer specify the `primeActualWeight`. If you create a new diverted shipment with the `diversion` and `divertedFromShipmentId` parameter, it will automatically inherit the primeActualWeight of its `divertedFromShipmentId` parent. Payment requests created on a diverted shipment \"chain\" will utilize the lowest weight possible in the chain to prevent overcharging as they are still separate shipments.  **NOTE**: New version in v3. Version will accept PPM addresses[pickupAddress, destinationAddress, secondaryPickupAddress secondaryDestinationAddress]. PPM postalCodes will be phased out[pickupPostalCode, secondaryPickupPostalCode, destinationPostalCode and secondaryDestinationPostalCode].  **WIP**: The Prime should be notified by a push notification whenever the TOO approves a shipment connected to one of their moves. Otherwise, the Prime can fetch the related move using the [getMoveTaskOrder](#operation/getMoveTaskOrder) endpoint and see if this shipment has the status `\"APPROVED\"`. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MtoShipmentApi.new
opts = {
  body: OpenapiClient::CreateMTOShipment.new({move_task_order_id: '1f2270c7-7166-40ae-981e-b200ebdf3054', shipment_type: OpenapiClient::MTOShipmentType::BOAT_HAUL_AWAY}) # CreateMTOShipment | 
}

begin
  # createMTOShipment
  result = api_instance.create_mto_shipment(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->create_mto_shipment: #{e}"
end
```

#### Using the create_mto_shipment_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<MTOShipment>, Integer, Hash)> create_mto_shipment_with_http_info(opts)

```ruby
begin
  # createMTOShipment
  data, status_code, headers = api_instance.create_mto_shipment_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <MTOShipment>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->create_mto_shipment_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **body** | [**CreateMTOShipment**](CreateMTOShipment.md) |  | [optional] |

### Return type

[**MTOShipment**](MTOShipment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## update_mto_shipment

> <MTOShipment> update_mto_shipment(mto_shipment_id, if_match, body)

updateMTOShipment

Updates an existing shipment for a move.  Note that there are some restrictions on nested objects:  * Service items: You cannot add or update service items using this endpoint. Please use [createMTOServiceItem](#operation/createMTOServiceItem) and [updateMTOServiceItem](#operation/updateMTOServiceItem) instead. * Agents: You cannot add or update agents using this endpoint. Please use [createMTOAgent](#operation/createMTOAgent) and [updateMTOAgent](#operation/updateMTOAgent) instead. * Addresses: You can add new addresses using this endpoint (and must use this endpoint to do so), but you cannot update existing ones. Please use [updateMTOShipmentAddress](#operation/updateMTOShipmentAddress) instead.  These restrictions are due to our [optimistic locking/concurrency control](https://transcom.github.io/mymove-docs/docs/dev/contributing/backend/use-optimistic-locking) mechanism.  Note that some fields cannot be manually changed but will still be updated automatically, such as `primeEstimatedWeightRecordedDate` and `requiredDeliveryDate`.  **NOTE**: New version in v3. Version will accept PPM addresses[pickupAddress, destinationAddress, secondaryPickupAddress secondaryDestinationAddress]. PPM postalCodes will be phased out[pickupPostalCode, secondaryPickupPostalCode, destinationPostalCode and secondaryDestinationPostalCode]. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MtoShipmentApi.new
mto_shipment_id = '38400000-8cf0-11bd-b23e-10b96e4ef00d' # String | UUID of the shipment being updated.
if_match = 'if_match_example' # String | Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 
body = OpenapiClient::UpdateMTOShipment.new # UpdateMTOShipment | 

begin
  # updateMTOShipment
  result = api_instance.update_mto_shipment(mto_shipment_id, if_match, body)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->update_mto_shipment: #{e}"
end
```

#### Using the update_mto_shipment_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<MTOShipment>, Integer, Hash)> update_mto_shipment_with_http_info(mto_shipment_id, if_match, body)

```ruby
begin
  # updateMTOShipment
  data, status_code, headers = api_instance.update_mto_shipment_with_http_info(mto_shipment_id, if_match, body)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <MTOShipment>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->update_mto_shipment_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mto_shipment_id** | **String** | UUID of the shipment being updated. |  |
| **if_match** | **String** | Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  |  |
| **body** | [**UpdateMTOShipment**](UpdateMTOShipment.md) |  |  |

### Return type

[**MTOShipment**](MTOShipment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

