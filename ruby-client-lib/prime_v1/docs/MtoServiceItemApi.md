# OpenapiClient::MtoServiceItemApi

All URIs are relative to */prime/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**create_mto_service_item**](MtoServiceItemApi.md#create_mto_service_item) | **POST** /mto-service-items | createMTOServiceItem |
| [**create_service_request_document_upload**](MtoServiceItemApi.md#create_service_request_document_upload) | **POST** /mto-service-items/{mtoServiceItemID}/uploads | createServiceRequestDocumentUpload |
| [**update_mto_service_item**](MtoServiceItemApi.md#update_mto_service_item) | **PATCH** /mto-service-items/{mtoServiceItemID} | updateMTOServiceItem |


## create_mto_service_item

> <Array<MTOServiceItem>> create_mto_service_item(opts)

createMTOServiceItem

Creates one or more MTOServiceItems. Not all service items may be created, please see details below.  This endpoint supports different body definitions. In the modelType field below, select the modelType corresponding  to the service item you wish to create and the documentation will update with the new definition.  Upon creation these items are associated with a Move Task Order and an MTO Shipment. The request must include UUIDs for the MTO and MTO Shipment connected to this service item. Some service item types require additional service items to be autogenerated when added - all created service items, autogenerated included, will be returned in the response.  To update a service item, please use [updateMTOServiceItem](#operation/updateMTOServiceItem) endpoint.  ---  **`MTOServiceItemOriginSIT`**  MTOServiceItemOriginSIT is a subtype of MTOServiceItem.  This model type describes a domestic origin SIT service item. Items can be created using this model type with the following codes:  **DOFSIT**  **1st day origin SIT service item**. When a DOFSIT is requested, the API will auto-create the following group of service items:   * DOFSIT - Domestic origin 1st day SIT   * DOASIT - Domestic origin Additional day SIT   * DOPSIT - Domestic origin SIT pickup   * DOSFSC - Domestic origin SIT fuel surcharge  **DOASIT**  **Addt'l days origin SIT service item**. This represents an additional day of storage for the same item. Additional DOASIT service items can be created and added to an existing shipment that **includes a DOFSIT service item**.  ---  **`MTOServiceItemDestSIT`**  MTOServiceItemDestSIT is a subtype of MTOServiceItem.  This model type describes a domestic destination SIT service item. Items can be created using this model type with the following codes:  **DDFSIT**  **1st day destination SIT service item**.  These additional fields are optional for creating a DDFSIT:   * `firstAvailableDeliveryDate1`     * string <date>     * First available date that Prime can deliver SIT service item.     * firstAvailableDeliveryDate1, dateOfContact1, and timeMilitary1 are required together   * `dateOfContact1`     * string <date>     * Date of attempted contact by the prime corresponding to `timeMilitary1`     * dateOfContact1, timeMilitary1, and firstAvailableDeliveryDate1 are required together   * `timeMilitary1`     * string\\d{4}Z     * Time of attempted contact corresponding to `dateOfContact1`, in military format.     * timeMilitary1, dateOfContact1, and firstAvailableDeliveryDate1 are required together   * `firstAvailableDeliveryDate2`     * string <date>     * Second available date that Prime can deliver SIT service item.     * firstAvailableDeliveryDate2, dateOfContact2, and timeMilitary2 are required together   * `dateOfContact2`     * string <date>     * Date of attempted contact delivery by the prime corresponding to `timeMilitary2`     * dateOfContact2, timeMilitary2, and firstAvailableDeliveryDate2 are required together   * `timeMilitary2`     * string\\d{4}Z     * Time of attempted contact corresponding to `dateOfContact2`, in military format.     * timeMilitary2, dateOfContact2, and firstAvailableDeliveryDate2 are required together  When a DDFSIT is requested, the API will auto-create the following group of service items:   * DDFSIT - Domestic destination 1st day SIT   * DDASIT - Domestic destination Additional day SIT   * DDDSIT - Domestic destination SIT delivery   * DDSFSC - Domestic destination SIT fuel surcharge  **DDASIT**  **Addt'l days destination SIT service item**. This represents an additional day of storage for the same item. Additional DDASIT service items can be created and added to an existing shipment that **includes a DDFSIT service item**. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MtoServiceItemApi.new
opts = {
  body: OpenapiClient::MTOServiceItem.new({move_task_order_id: '1f2270c7-7166-40ae-981e-b200ebdf3054', model_type: OpenapiClient::MTOServiceItemModelType::MTO_SERVICE_ITEM_BASIC}) # MTOServiceItem | 
}

begin
  # createMTOServiceItem
  result = api_instance.create_mto_service_item(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoServiceItemApi->create_mto_service_item: #{e}"
end
```

#### Using the create_mto_service_item_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Array<MTOServiceItem>>, Integer, Hash)> create_mto_service_item_with_http_info(opts)

```ruby
begin
  # createMTOServiceItem
  data, status_code, headers = api_instance.create_mto_service_item_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Array<MTOServiceItem>>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoServiceItemApi->create_mto_service_item_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **body** | [**MTOServiceItem**](MTOServiceItem.md) |  | [optional] |

### Return type

[**Array&lt;MTOServiceItem&gt;**](MTOServiceItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## create_service_request_document_upload

> <UploadWithOmissions> create_service_request_document_upload(mto_service_item_id, file)

createServiceRequestDocumentUpload

### Functionality  This endpoint **uploads** a Service Request document for a ServiceItem.  The ServiceItem should already exist.  ServiceItems are created with the [createMTOServiceItem](#operation/createMTOServiceItem) endpoint. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MtoServiceItemApi.new
mto_service_item_id = 'mto_service_item_id_example' # String | UUID of the service item to use.
file = File.new('/path/to/some/file') # File | The file to upload.

begin
  # createServiceRequestDocumentUpload
  result = api_instance.create_service_request_document_upload(mto_service_item_id, file)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoServiceItemApi->create_service_request_document_upload: #{e}"
end
```

#### Using the create_service_request_document_upload_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<UploadWithOmissions>, Integer, Hash)> create_service_request_document_upload_with_http_info(mto_service_item_id, file)

```ruby
begin
  # createServiceRequestDocumentUpload
  data, status_code, headers = api_instance.create_service_request_document_upload_with_http_info(mto_service_item_id, file)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <UploadWithOmissions>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoServiceItemApi->create_service_request_document_upload_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mto_service_item_id** | **String** | UUID of the service item to use. |  |
| **file** | **File** | The file to upload. |  |

### Return type

[**UploadWithOmissions**](UploadWithOmissions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json


## update_mto_service_item

> <MTOServiceItem> update_mto_service_item(mto_service_item_id, if_match, body)

updateMTOServiceItem

Updates MTOServiceItems after creation. Not all service items or fields may be updated, please see details below.  This endpoint supports different body definitions. In the modelType field below, select the modelType corresponding  to the service item you wish to update and the documentation will update with the new definition.  * Addresses: To update a destination service item's SIT destination final address, update the shipment destination address. For approved shipments, please use [updateShipmentDestinationAddress](#mtoShipment/updateShipmentDestinationAddress). For shipments not yet approved, please use [updateMTOShipmentAddress](#mtoShipment/updateMTOShipmentAddress).  * SIT Service Items: Take note that when updating `sitCustomerContacted`, `sitDepartureDate`, or `sitRequestedDelivery`, we want those to be updated on `DOASIT` (for origin SIT) and `DDASIT` (for destination SIT). If updating those values in other service items, the office users will not have as much attention to those values.  To create a service item, please use [createMTOServiceItem](#mtoServiceItem/createMTOServiceItem)) endpoint. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MtoServiceItemApi.new
mto_service_item_id = 'mto_service_item_id_example' # String | UUID of service item to update.
if_match = 'if_match_example' # String | Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 
body = OpenapiClient::UpdateMTOServiceItem.new({model_type: OpenapiClient::UpdateMTOServiceItemModelType::UPDATE_MTO_SERVICE_ITEM_SIT}) # UpdateMTOServiceItem | 

begin
  # updateMTOServiceItem
  result = api_instance.update_mto_service_item(mto_service_item_id, if_match, body)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoServiceItemApi->update_mto_service_item: #{e}"
end
```

#### Using the update_mto_service_item_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<MTOServiceItem>, Integer, Hash)> update_mto_service_item_with_http_info(mto_service_item_id, if_match, body)

```ruby
begin
  # updateMTOServiceItem
  data, status_code, headers = api_instance.update_mto_service_item_with_http_info(mto_service_item_id, if_match, body)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <MTOServiceItem>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoServiceItemApi->update_mto_service_item_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mto_service_item_id** | **String** | UUID of service item to update. |  |
| **if_match** | **String** | Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  |  |
| **body** | [**UpdateMTOServiceItem**](UpdateMTOServiceItem.md) |  |  |

### Return type

[**MTOServiceItem**](MTOServiceItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

