# OpenapiClient::MoveTaskOrderApi

All URIs are relative to */prime/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**create_excess_weight_record**](MoveTaskOrderApi.md#create_excess_weight_record) | **POST** /move-task-orders/{moveTaskOrderID}/excess-weight-record | createExcessWeightRecord |
| [**download_move_order**](MoveTaskOrderApi.md#download_move_order) | **GET** /moves/{locator}/documents | Downloads move order as a PDF |
| [**get_move_task_order**](MoveTaskOrderApi.md#get_move_task_order) | **GET** /move-task-orders/{moveID} | getMoveTaskOrder |
| [**list_moves**](MoveTaskOrderApi.md#list_moves) | **GET** /moves | listMoves |
| [**update_mto_post_counseling_information**](MoveTaskOrderApi.md#update_mto_post_counseling_information) | **PATCH** /move-task-orders/{moveTaskOrderID}/post-counseling-info | updateMTOPostCounselingInformation |


## create_excess_weight_record

> <ExcessWeightRecord> create_excess_weight_record(move_task_order_id, file)

createExcessWeightRecord

Uploads an excess weight record, which is a document that proves that the movers or contractors have counseled the customer about their excess weight. Excess weight counseling should occur after the sum of the shipments for the customer's move crosses the excess weight alert threshold. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MoveTaskOrderApi.new
move_task_order_id = '38400000-8cf0-11bd-b23e-10b96e4ef00d' # String | UUID of the move being updated.
file = File.new('/path/to/some/file') # File | The file to upload.

begin
  # createExcessWeightRecord
  result = api_instance.create_excess_weight_record(move_task_order_id, file)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MoveTaskOrderApi->create_excess_weight_record: #{e}"
end
```

#### Using the create_excess_weight_record_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ExcessWeightRecord>, Integer, Hash)> create_excess_weight_record_with_http_info(move_task_order_id, file)

```ruby
begin
  # createExcessWeightRecord
  data, status_code, headers = api_instance.create_excess_weight_record_with_http_info(move_task_order_id, file)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ExcessWeightRecord>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MoveTaskOrderApi->create_excess_weight_record_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **move_task_order_id** | **String** | UUID of the move being updated. |  |
| **file** | **File** | The file to upload. |  |

### Return type

[**ExcessWeightRecord**](ExcessWeightRecord.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json


## download_move_order

> File download_move_order(locator, opts)

Downloads move order as a PDF

### Functionality This endpoint downloads all uploaded move order documentations into one download file by locator.  ### Errors * The move must be in need counseling state. * The move client's origin duty location must not currently have gov counseling. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MoveTaskOrderApi.new
locator = 'locator_example' # String | the locator code for move order to be downloaded
opts = {
  type: 'ALL' # String | upload type
}

begin
  # Downloads move order as a PDF
  result = api_instance.download_move_order(locator, opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MoveTaskOrderApi->download_move_order: #{e}"
end
```

#### Using the download_move_order_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(File, Integer, Hash)> download_move_order_with_http_info(locator, opts)

```ruby
begin
  # Downloads move order as a PDF
  data, status_code, headers = api_instance.download_move_order_with_http_info(locator, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => File
rescue OpenapiClient::ApiError => e
  puts "Error when calling MoveTaskOrderApi->download_move_order_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **locator** | **String** | the locator code for move order to be downloaded |  |
| **type** | **String** | upload type | [optional][default to &#39;ALL&#39;] |

### Return type

**File**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/pdf


## get_move_task_order

> <MoveTaskOrder> get_move_task_order(move_id)

getMoveTaskOrder

### Functionality This endpoint gets an individual MoveTaskOrder by ID.  It will provide information about the Customer and any associated MTOShipments, MTOServiceItems and PaymentRequests. 

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


## list_moves

> <Array<ListMove>> list_moves(opts)

listMoves

Gets all moves that have been reviewed and approved by the TOO. The `since` parameter can be used to filter this list down to only the moves that have been updated since the provided timestamp. A move will be considered updated if the `updatedAt` timestamp on the move or on its orders, shipments, service items, or payment requests, is later than the provided date and time.  **WIP**: Include what causes moves to leave this list. Currently, once the `availableToPrimeAt` timestamp has been set, that move will always appear in this list. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MoveTaskOrderApi.new
opts = {
  since: Time.parse('2013-10-20T19:20:30+01:00') # Time | Only return moves updated since this time. Formatted like \"2021-07-23T18:30:47.116Z\"
}

begin
  # listMoves
  result = api_instance.list_moves(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MoveTaskOrderApi->list_moves: #{e}"
end
```

#### Using the list_moves_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<ListMove>>, Integer, Hash)> list_moves_with_http_info(opts)

```ruby
begin
  # listMoves
  data, status_code, headers = api_instance.list_moves_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<ListMove>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MoveTaskOrderApi->list_moves_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **since** | **Time** | Only return moves updated since this time. Formatted like \&quot;2021-07-23T18:30:47.116Z\&quot; | [optional] |

### Return type

[**Array&lt;ListMove&gt;**](ListMove.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## update_mto_post_counseling_information

> <MoveTaskOrder> update_mto_post_counseling_information(move_task_order_id, if_match)

updateMTOPostCounselingInformation

### Functionality This endpoint **updates** the MoveTaskOrder to indicate that the Prime has completed Counseling. This update uses the moveTaskOrderID provided in the path, updates the move status and marks child elements of the move to indicate the update. No body object is expected for this request.  **For Full/Partial PPMs**: This action is required so that the customer can start uploading their proof of service docs.  **For other move types**: This action is required for auditing reasons so that we have a record of when the Prime counseled the customer. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MoveTaskOrderApi.new
move_task_order_id = 'move_task_order_id_example' # String | ID of move task order to use.
if_match = 'if_match_example' # String | Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 

begin
  # updateMTOPostCounselingInformation
  result = api_instance.update_mto_post_counseling_information(move_task_order_id, if_match)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MoveTaskOrderApi->update_mto_post_counseling_information: #{e}"
end
```

#### Using the update_mto_post_counseling_information_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<MoveTaskOrder>, Integer, Hash)> update_mto_post_counseling_information_with_http_info(move_task_order_id, if_match)

```ruby
begin
  # updateMTOPostCounselingInformation
  data, status_code, headers = api_instance.update_mto_post_counseling_information_with_http_info(move_task_order_id, if_match)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <MoveTaskOrder>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MoveTaskOrderApi->update_mto_post_counseling_information_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **move_task_order_id** | **String** | ID of move task order to use. |  |
| **if_match** | **String** | Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  |  |

### Return type

[**MoveTaskOrder**](MoveTaskOrder.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

