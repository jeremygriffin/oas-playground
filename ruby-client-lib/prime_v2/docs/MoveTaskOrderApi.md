# OpenapiClient::MoveTaskOrderApi

All URIs are relative to */prime/v2*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**get_move_task_order**](MoveTaskOrderApi.md#get_move_task_order) | **GET** /move-task-orders/{moveID} | getMoveTaskOrder |


## get_move_task_order

> <MoveTaskOrder> get_move_task_order(move_id)

getMoveTaskOrder

### Functionality This endpoint gets an individual MoveTaskOrder by ID.  It will provide information about the Customer and any associated MTOShipments, MTOServiceItems and PaymentRequests.  **NOTE**: New version in v3. Version will return PPM addresses[pickupAddress, destinationAddress, secondaryPickupAddress secondaryDestinationAddress]. PPM postalCodes will be phased out[pickupPostalCode, secondaryPickupPostalCode, destinationPostalCode and secondaryDestinationPostalCode]. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MoveTaskOrderApi.new
move_id = 'move_id_example' # String | UUID or MoveCode of move task order to use.

begin
  # getMoveTaskOrder
  result = api_instance.get_move_task_order(move_id)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MoveTaskOrderApi->get_move_task_order: #{e}"
end
```

#### Using the get_move_task_order_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<MoveTaskOrder>, Integer, Hash)> get_move_task_order_with_http_info(move_id)

```ruby
begin
  # getMoveTaskOrder
  data, status_code, headers = api_instance.get_move_task_order_with_http_info(move_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <MoveTaskOrder>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MoveTaskOrderApi->get_move_task_order_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **move_id** | **String** | UUID or MoveCode of move task order to use. |  |

### Return type

[**MoveTaskOrder**](MoveTaskOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

