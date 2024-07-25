# OpenapiClient::MtoShipmentApi

All URIs are relative to */prime/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**create_mto_agent**](MtoShipmentApi.md#create_mto_agent) | **POST** /mto-shipments/{mtoShipmentID}/agents | createMTOAgent |
| [**create_mto_shipment**](MtoShipmentApi.md#create_mto_shipment) | **POST** /mto-shipments | createMTOShipment |
| [**create_sit_extension**](MtoShipmentApi.md#create_sit_extension) | **POST** /mto-shipments/{mtoShipmentID}/sit-extensions | createSITExtension |
| [**delete_mto_shipment**](MtoShipmentApi.md#delete_mto_shipment) | **DELETE** /mto-shipments/{mtoShipmentID} | deleteMTOShipment |
| [**update_mto_agent**](MtoShipmentApi.md#update_mto_agent) | **PUT** /mto-shipments/{mtoShipmentID}/agents/{agentID} | updateMTOAgent |
| [**update_mto_shipment**](MtoShipmentApi.md#update_mto_shipment) | **PATCH** /mto-shipments/{mtoShipmentID} | updateMTOShipment |
| [**update_mto_shipment_address**](MtoShipmentApi.md#update_mto_shipment_address) | **PUT** /mto-shipments/{mtoShipmentID}/addresses/{addressID} | updateMTOShipmentAddress |
| [**update_mto_shipment_status**](MtoShipmentApi.md#update_mto_shipment_status) | **PATCH** /mto-shipments/{mtoShipmentID}/status | updateMTOShipmentStatus |
| [**update_reweigh**](MtoShipmentApi.md#update_reweigh) | **PATCH** /mto-shipments/{mtoShipmentID}/reweighs/{reweighID} | updateReweigh |
| [**update_shipment_destination_address**](MtoShipmentApi.md#update_shipment_destination_address) | **POST** /mto-shipments/{mtoShipmentID}/shipment-address-updates | updateShipmentDestinationAddress |


## create_mto_agent

> <MTOAgent> create_mto_agent(mto_shipment_id, body)

createMTOAgent

### Functionality This endpoint is used to **create** and add agents for an existing MTO Shipment. Only the fields being modified need to be sent in the request body.  ### Errors The agent must always have a name and at least one method of contact (either `email` or `phone`).  The agent must be associated with the MTO shipment passed in the url.  The shipment should be associated with an MTO that is available to the Pime. If the caller requests a new agent, and the shipment is not on an available MTO, the caller will receive a **NotFound** response. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MtoShipmentApi.new
mto_shipment_id = '38400000-8cf0-11bd-b23e-10b96e4ef00d' # String | UUID of the shipment associated with the agent
body = OpenapiClient::MTOAgent.new # MTOAgent | 

begin
  # createMTOAgent
  result = api_instance.create_mto_agent(mto_shipment_id, body)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->create_mto_agent: #{e}"
end
```

#### Using the create_mto_agent_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<MTOAgent>, Integer, Hash)> create_mto_agent_with_http_info(mto_shipment_id, body)

```ruby
begin
  # createMTOAgent
  data, status_code, headers = api_instance.create_mto_agent_with_http_info(mto_shipment_id, body)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <MTOAgent>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->create_mto_agent_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mto_shipment_id** | **String** | UUID of the shipment associated with the agent |  |
| **body** | [**MTOAgent**](MTOAgent.md) |  |  |

### Return type

[**MTOAgent**](MTOAgent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## create_mto_shipment

> <MTOShipment> create_mto_shipment(opts)

createMTOShipment

_[Deprecated: sunset on 2024-04-08]_ This endpoint is deprecated and will be removed in a future version. Please use the new endpoint at `/prime/v2/createMTOShipment` instead.  Creates a new shipment within the specified move. This endpoint should be used whenever the movers identify a need for an additional shipment. The new shipment will be submitted to the TOO for review, and the TOO must approve it before the contractor can proceed with billing.  **WIP**: The Prime should be notified by a push notification whenever the TOO approves a shipment connected to one of their moves. Otherwise, the Prime can fetch the related move using the [getMoveTaskOrder](#operation/getMoveTaskOrder) endpoint and see if this shipment has the status `\"APPROVED\"`. 

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


## create_sit_extension

> <SITExtension> create_sit_extension(mto_shipment_id, body)

createSITExtension

### Functionality This endpoint creates a storage in transit (SIT) extension request for a shipment. A SIT extension request is a request an increase in the shipment day allowance for the number of days a shipment is allowed to be in SIT. The total SIT day allowance includes time spent in both origin and destination SIT. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MtoShipmentApi.new
mto_shipment_id = '38400000-8cf0-11bd-b23e-10b96e4ef00d' # String | UUID of the shipment associated with the agent
body = OpenapiClient::CreateSITExtension.new({request_reason: 'SERIOUS_ILLNESS_MEMBER', contractor_remarks: 'We need SIT additional days. The customer has not found a house yet.', requested_days: 30}) # CreateSITExtension | 

begin
  # createSITExtension
  result = api_instance.create_sit_extension(mto_shipment_id, body)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->create_sit_extension: #{e}"
end
```

#### Using the create_sit_extension_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<SITExtension>, Integer, Hash)> create_sit_extension_with_http_info(mto_shipment_id, body)

```ruby
begin
  # createSITExtension
  data, status_code, headers = api_instance.create_sit_extension_with_http_info(mto_shipment_id, body)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <SITExtension>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->create_sit_extension_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mto_shipment_id** | **String** | UUID of the shipment associated with the agent |  |
| **body** | [**CreateSITExtension**](CreateSITExtension.md) |  |  |

### Return type

[**SITExtension**](SITExtension.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## delete_mto_shipment

> delete_mto_shipment(mto_shipment_id)

deleteMTOShipment

### Functionality This endpoint deletes an individual shipment by ID.  ### Errors * The mtoShipment should be associated with an MTO that is available to prime. * The mtoShipment must be a PPM shipment. * Counseling should not have already been completed for the associated MTO. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MtoShipmentApi.new
mto_shipment_id = '38400000-8cf0-11bd-b23e-10b96e4ef00d' # String | UUID of the shipment to be deleted

begin
  # deleteMTOShipment
  api_instance.delete_mto_shipment(mto_shipment_id)
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->delete_mto_shipment: #{e}"
end
```

#### Using the delete_mto_shipment_with_http_info variant

This returns an Array which contains the response data (`nil` in this case), status code and headers.

> <Array(nil, Integer, Hash)> delete_mto_shipment_with_http_info(mto_shipment_id)

```ruby
begin
  # deleteMTOShipment
  data, status_code, headers = api_instance.delete_mto_shipment_with_http_info(mto_shipment_id)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => nil
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->delete_mto_shipment_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mto_shipment_id** | **String** | UUID of the shipment to be deleted |  |

### Return type

nil (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


## update_mto_agent

> <MTOAgent> update_mto_agent(mto_shipment_id, agent_id, if_match, body)

updateMTOAgent

### Functionality This endpoint is used to **update** the agents for an MTO Shipment. Only the fields being modified need to be sent in the request body.  ### Errors: The agent must always have a name and at least one method of contact (either `email` or `phone`).  The agent must be associated with the MTO shipment passed in the url.  The shipment should be associated with an MTO that is available to the Prime. If the caller requests an update to an agent, and the shipment is not on an available MTO, the caller will receive a **NotFound** response. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MtoShipmentApi.new
mto_shipment_id = '38400000-8cf0-11bd-b23e-10b96e4ef00d' # String | UUID of the shipment associated with the agent
agent_id = '38400000-8cf0-11bd-b23e-10b96e4ef00d' # String | UUID of the agent being updated
if_match = 'if_match_example' # String | Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 
body = OpenapiClient::MTOAgent.new # MTOAgent | 

begin
  # updateMTOAgent
  result = api_instance.update_mto_agent(mto_shipment_id, agent_id, if_match, body)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->update_mto_agent: #{e}"
end
```

#### Using the update_mto_agent_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<MTOAgent>, Integer, Hash)> update_mto_agent_with_http_info(mto_shipment_id, agent_id, if_match, body)

```ruby
begin
  # updateMTOAgent
  data, status_code, headers = api_instance.update_mto_agent_with_http_info(mto_shipment_id, agent_id, if_match, body)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <MTOAgent>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->update_mto_agent_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mto_shipment_id** | **String** | UUID of the shipment associated with the agent |  |
| **agent_id** | **String** | UUID of the agent being updated |  |
| **if_match** | **String** | Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  |  |
| **body** | [**MTOAgent**](MTOAgent.md) |  |  |

### Return type

[**MTOAgent**](MTOAgent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## update_mto_shipment

> <MTOShipment> update_mto_shipment(mto_shipment_id, if_match, body)

updateMTOShipment

_[Deprecated: sunset on August 5th, 2024]_ This endpoint is deprecated and will be removed in a future version. Please use the new endpoint at `/prime/v2/updateMTOShipment` instead.  **DEPRECATION ON AUGUST 5TH, 2024** Following deprecation, there is an edge case scenario where a PPM shipment with no addresses could be updated and it would also update the final destination SIT address for SIT service items. This edge case has been removed as you should not be able to update items using this endpoint. Third-party APIs have confirmed they will require deprecation for this change.  Updates an existing shipment for a move.  Note that there are some restrictions on nested objects:  * Service items: You cannot add or update service items using this endpoint. Please use [createMTOServiceItem](#operation/createMTOServiceItem) and [updateMTOServiceItem](#operation/updateMTOServiceItem) instead. * Agents: You cannot add or update agents using this endpoint. Please use [createMTOAgent](#operation/createMTOAgent) and [updateMTOAgent](#operation/updateMTOAgent) instead. * Addresses: You can add new addresses using this endpoint (and must use this endpoint to do so), but you cannot update existing ones. Please use [updateMTOShipmentAddress](#operation/updateMTOShipmentAddress) instead.  These restrictions are due to our [optimistic locking/concurrency control](https://transcom.github.io/mymove-docs/docs/dev/contributing/backend/use-optimistic-locking) mechanism.  Note that some fields cannot be manually changed but will still be updated automatically, such as `primeEstimatedWeightRecordedDate` and `requiredDeliveryDate`. 

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


## update_mto_shipment_address

> <Address> update_mto_shipment_address(mto_shipment_id, address_id, if_match, body)

updateMTOShipmentAddress

### Functionality This endpoint is used to **update** the pickup, secondary, and destination addresses on an MTO Shipment. mto-shipments/{mtoShipmentID}/shipment-address-updates is for updating a delivery address. The address details completely replace the original, except for the UUID. Therefore a complete address should be sent in the request. When a destination address on a shipment is updated, the destination SIT service items address ID will also be updated so that shipment and service item final destinations match.  This endpoint **cannot create** an address. To create an address on an MTO shipment, the caller must use [updateMTOShipment](#operation/updateMTOShipment) as the parent shipment has to be updated with the appropriate link to the address.  ### Errors The address must be associated with the mtoShipment passed in the url. In other words, it should be listed as pickupAddress, destinationAddress, secondaryPickupAddress or secondaryDeliveryAddress on the mtoShipment provided. If it is not, caller will receive a **Conflict** Error.  The mtoShipment should be associated with an MTO that is available to prime. If the caller requests an update to an address, and the shipment is not on an available MTO, the caller will receive a **NotFound** Error. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MtoShipmentApi.new
mto_shipment_id = '38400000-8cf0-11bd-b23e-10b96e4ef00d' # String | UUID of the shipment associated with the address
address_id = '38400000-8cf0-11bd-b23e-10b96e4ef00d' # String | UUID of the address being updated
if_match = 'if_match_example' # String | Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 
body = OpenapiClient::Address.new({street_address1: '123 Main Ave', city: 'Anytown', state: 'AL', postal_code: '90210'}) # Address | 

begin
  # updateMTOShipmentAddress
  result = api_instance.update_mto_shipment_address(mto_shipment_id, address_id, if_match, body)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->update_mto_shipment_address: #{e}"
end
```

#### Using the update_mto_shipment_address_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Address>, Integer, Hash)> update_mto_shipment_address_with_http_info(mto_shipment_id, address_id, if_match, body)

```ruby
begin
  # updateMTOShipmentAddress
  data, status_code, headers = api_instance.update_mto_shipment_address_with_http_info(mto_shipment_id, address_id, if_match, body)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Address>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->update_mto_shipment_address_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mto_shipment_id** | **String** | UUID of the shipment associated with the address |  |
| **address_id** | **String** | UUID of the address being updated |  |
| **if_match** | **String** | Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  |  |
| **body** | [**Address**](Address.md) |  |  |

### Return type

[**Address**](Address.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## update_mto_shipment_status

> <MTOShipment> update_mto_shipment_status(mto_shipment_id, if_match, body)

updateMTOShipmentStatus

### Functionality This endpoint should be used by the Prime to confirm the cancellation of a shipment. It allows the shipment status to be changed to \"CANCELED.\" Currently, the Prime cannot update the shipment to any other status. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MtoShipmentApi.new
mto_shipment_id = '38400000-8cf0-11bd-b23e-10b96e4ef00d' # String | UUID of the shipment associated with the agent
if_match = 'if_match_example' # String | Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 
body = OpenapiClient::UpdateMTOShipmentStatus.new # UpdateMTOShipmentStatus | 

begin
  # updateMTOShipmentStatus
  result = api_instance.update_mto_shipment_status(mto_shipment_id, if_match, body)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->update_mto_shipment_status: #{e}"
end
```

#### Using the update_mto_shipment_status_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<MTOShipment>, Integer, Hash)> update_mto_shipment_status_with_http_info(mto_shipment_id, if_match, body)

```ruby
begin
  # updateMTOShipmentStatus
  data, status_code, headers = api_instance.update_mto_shipment_status_with_http_info(mto_shipment_id, if_match, body)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <MTOShipment>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->update_mto_shipment_status_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mto_shipment_id** | **String** | UUID of the shipment associated with the agent |  |
| **if_match** | **String** | Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  |  |
| **body** | [**UpdateMTOShipmentStatus**](UpdateMTOShipmentStatus.md) |  |  |

### Return type

[**MTOShipment**](MTOShipment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## update_reweigh

> <Reweigh> update_reweigh(mto_shipment_id, reweigh_id, if_match, body)

updateReweigh

### Functionality This endpoint can be used to update a reweigh with a new weight or to provide the reason why a reweigh did not occur. Only one of weight or verificationReason should be sent in the request body.  A reweigh is the second recorded weight for a shipment, as validated by certified weight tickets. Applies to one shipment. A reweigh can be triggered automatically, or requested by the customer or transportation office. Not all shipments are reweighed, so not all shipments will have a reweigh weight. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MtoShipmentApi.new
mto_shipment_id = '38400000-8cf0-11bd-b23e-10b96e4ef00d' # String | UUID of the shipment associated with the reweigh
reweigh_id = '38400000-8cf0-11bd-b23e-10b96e4ef00d' # String | UUID of the reweigh being updated
if_match = 'if_match_example' # String | Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 
body = OpenapiClient::UpdateReweigh.new # UpdateReweigh | 

begin
  # updateReweigh
  result = api_instance.update_reweigh(mto_shipment_id, reweigh_id, if_match, body)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->update_reweigh: #{e}"
end
```

#### Using the update_reweigh_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<Reweigh>, Integer, Hash)> update_reweigh_with_http_info(mto_shipment_id, reweigh_id, if_match, body)

```ruby
begin
  # updateReweigh
  data, status_code, headers = api_instance.update_reweigh_with_http_info(mto_shipment_id, reweigh_id, if_match, body)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <Reweigh>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->update_reweigh_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mto_shipment_id** | **String** | UUID of the shipment associated with the reweigh |  |
| **reweigh_id** | **String** | UUID of the reweigh being updated |  |
| **if_match** | **String** | Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  |  |
| **body** | [**UpdateReweigh**](UpdateReweigh.md) |  |  |

### Return type

[**Reweigh**](Reweigh.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## update_shipment_destination_address

> <ShipmentAddressUpdate> update_shipment_destination_address(mto_shipment_id, if_match, body)

updateShipmentDestinationAddress

### Functionality This endpoint is used so the Prime can request an **update** for the destination address on an MTO Shipment, after the destination address has already been approved. If automatically approved or TOO approves, this will update the final destination address values for destination SIT service items to be the same as the changed destination address that was approved. Address updates will be automatically approved unless they change:   - The service area   - Mileage bracket for direct delivery   - the address and the distance between the old and new address is > 50   - Domestic Short Haul to Domestic Line Haul or vice versa       - Shipments that start and end in one ZIP3 use Short Haul pricing       - Shipments that start and end in different ZIP3s use Line Haul pricing  For those, changes will require TOO approval. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::MtoShipmentApi.new
mto_shipment_id = '38400000-8cf0-11bd-b23e-10b96e4ef00d' # String | UUID of the shipment associated with the address
if_match = 'if_match_example' # String | Needs to be the eTag of the mtoShipment. Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 
body = OpenapiClient::UpdateShipmentDestinationAddress.new({new_address: OpenapiClient::Address.new({street_address1: '123 Main Ave', city: 'Anytown', state: 'AL', postal_code: '90210'}), contractor_remarks: 'Customer reached out to me this week and let me know they want to move somewhere else.'}) # UpdateShipmentDestinationAddress | 

begin
  # updateShipmentDestinationAddress
  result = api_instance.update_shipment_destination_address(mto_shipment_id, if_match, body)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->update_shipment_destination_address: #{e}"
end
```

#### Using the update_shipment_destination_address_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<ShipmentAddressUpdate>, Integer, Hash)> update_shipment_destination_address_with_http_info(mto_shipment_id, if_match, body)

```ruby
begin
  # updateShipmentDestinationAddress
  data, status_code, headers = api_instance.update_shipment_destination_address_with_http_info(mto_shipment_id, if_match, body)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <ShipmentAddressUpdate>
rescue OpenapiClient::ApiError => e
  puts "Error when calling MtoShipmentApi->update_shipment_destination_address_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **mto_shipment_id** | **String** | UUID of the shipment associated with the address |  |
| **if_match** | **String** | Needs to be the eTag of the mtoShipment. Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  |  |
| **body** | [**UpdateShipmentDestinationAddress**](UpdateShipmentDestinationAddress.md) |  |  |

### Return type

[**ShipmentAddressUpdate**](ShipmentAddressUpdate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

