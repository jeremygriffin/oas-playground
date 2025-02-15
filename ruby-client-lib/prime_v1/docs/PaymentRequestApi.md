# OpenapiClient::PaymentRequestApi

All URIs are relative to */prime/v1*

| Method | HTTP request | Description |
| ------ | ------------ | ----------- |
| [**create_payment_request**](PaymentRequestApi.md#create_payment_request) | **POST** /payment-requests | createPaymentRequest |
| [**create_upload**](PaymentRequestApi.md#create_upload) | **POST** /payment-requests/{paymentRequestID}/uploads | createUpload |


## create_payment_request

> <PaymentRequest> create_payment_request(opts)

createPaymentRequest

Creates a new instance of a paymentRequest and is assigned the status `PENDING`. A move task order can have multiple payment requests, and a final payment request can be marked using boolean `isFinal`.  If a `PENDING` payment request is recalculated, a new payment request is created and the original request is marked with the status `DEPRECATED`.  **NOTE**: In order to create a payment request for most service items, the shipment *must* be updated with the `PrimeActualWeight` value via [updateMTOShipment](#operation/updateMTOShipment).  **FSC - Fuel Surcharge** service items require `ActualPickupDate` to be updated on the shipment.  A service item can be on several payment requests in the case of partial payment requests and payments.  In the request, if no params are necessary, then just the `serviceItem` `id` is required. For example: ```json {   \"isFinal\": false,   \"moveTaskOrderID\": \"uuid\",   \"serviceItems\": [     {       \"id\": \"uuid\",     },     {       \"id\": \"uuid\",       \"params\": [         {           \"key\": \"Service Item Parameter Name\",           \"value\": \"Service Item Parameter Value\"         }       ]     }   ],   \"pointOfContact\": \"string\" } ```  SIT Service Items & Accepted Payment Request Parameters: --- If `WeightBilled` is not provided then the full shipment weight (`PrimeActualWeight`) will be considered in the calculation.  **NOTE**: Diversions have a unique calcuation for payment requests without a `WeightBilled` parameter.  If you created a payment request for a diversion and `WeightBilled` is not provided, then the following will be used in the calculation: - The lowest shipment weight (`PrimeActualWeight`) found in the diverted shipment chain. - The lowest reweigh weight found in the diverted shipment chain.  The diverted shipment chain is created by referencing the `diversion` boolean, `divertedFromShipmentId` UUID, and matching destination to pickup addresses. If the chain cannot be established it will fall back to the `PrimeActualWeight` of the current shipment. This is utilized because diverted shipments are all one single shipment, but going to different locations. The lowest weight found is the true shipment weight, and thus we search the chain of shipments for the lowest weight found.  **DOFSIT - Domestic origin 1st day SIT** ```json   \"params\": [     {       \"key\": \"WeightBilled\",       \"value\": \"integer\"     }   ] ```  **DOASIT - Domestic origin add'l SIT** *(SITPaymentRequestStart & SITPaymentRequestEnd are **REQUIRED**)* *To create a paymentRequest for this service item, the `SITPaymentRequestStart` and `SITPaymentRequestEnd` dates must not overlap previously requested SIT dates.* ```json   \"params\": [     {       \"key\": \"WeightBilled\",       \"value\": \"integer\"     },     {       \"key\": \"SITPaymentRequestStart\",       \"value\": \"date\"     },     {       \"key\": \"SITPaymentRequestEnd\",       \"value\": \"date\"     }   ] ```  **DOPSIT - Domestic origin SIT pickup** ```json   \"params\": [     {       \"key\": \"WeightBilled\",       \"value\": \"integer\"     }   ] ```  **DOSHUT - Domestic origin shuttle service** ```json   \"params\": [     {       \"key\": \"WeightBilled\",       \"value\": \"integer\"     }   ] ```  **DDFSIT - Domestic destination 1st day SIT** ```json   \"params\": [     {       \"key\": \"WeightBilled\",       \"value\": \"integer\"     }   ] ```  **DDASIT - Domestic destination add'l SIT** *(SITPaymentRequestStart & SITPaymentRequestEnd are **REQUIRED**)* *To create a paymentRequest for this service item, the `SITPaymentRequestStart` and `SITPaymentRequestEnd` dates must not overlap previously requested SIT dates.* ```json   \"params\": [     {       \"key\": \"WeightBilled\",       \"value\": \"integer\"     },     {       \"key\": \"SITPaymentRequestStart\",       \"value\": \"date\"     },     {       \"key\": \"SITPaymentRequestEnd\",       \"value\": \"date\"     }   ] ```  **DDDSIT - Domestic destination SIT delivery** *To create a paymentRequest for this service item, it must first have a final address set via [updateMTOServiceItem](#operation/updateMTOServiceItem).* ```json   \"params\": [     {       \"key\": \"WeightBilled\",       \"value\": \"integer\"     }   ] ```  **DDSHUT - Domestic destination shuttle service** ```json   \"params\": [     {       \"key\": \"WeightBilled\",       \"value\": \"integer\"     }   ] ``` --- 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::PaymentRequestApi.new
opts = {
  body: OpenapiClient::CreatePaymentRequest.new({move_task_order_id: 'c56a4180-65aa-42ec-a945-5fd21dec0538', service_items: [OpenapiClient::ServiceItem.new]}) # CreatePaymentRequest | 
}

begin
  # createPaymentRequest
  result = api_instance.create_payment_request(opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PaymentRequestApi->create_payment_request: #{e}"
end
```

#### Using the create_payment_request_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<PaymentRequest>, Integer, Hash)> create_payment_request_with_http_info(opts)

```ruby
begin
  # createPaymentRequest
  data, status_code, headers = api_instance.create_payment_request_with_http_info(opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <PaymentRequest>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PaymentRequestApi->create_payment_request_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **body** | [**CreatePaymentRequest**](CreatePaymentRequest.md) |  | [optional] |

### Return type

[**PaymentRequest**](PaymentRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


## create_upload

> <UploadWithOmissions> create_upload(payment_request_id, file, opts)

createUpload

### Functionality This endpoint **uploads** a Proof of Service document for a PaymentRequest.  The PaymentRequest should already exist.  Optional key of **isWeightTicket** indicates if the document is a weight ticket or not. This will be used for partial and full deliveries and makes it easier for the Task Invoicing Officers to locate and review service item documents. If left empty, it will assume it is NOT a weight ticket.  The formdata in the body of the POST request that is sent should look like this if it IS a weight ticket being attached to an existing payment request:   ```json   {     \"file\": \"filePath\",     \"isWeightTicket\": true   }   ```   If the proof of service doc is NOT a weight ticket, it will look like this - or you can leave it empty:   ```json   {     \"file\": \"filePath\",     \"isWeightTicket\": false   }   ```   ```json   {     \"file\": \"filePath\",   }   ```  PaymentRequests are created with the [createPaymentRequest](#operation/createPaymentRequest) endpoint. 

### Examples

```ruby
require 'time'
require 'openapi_client'

api_instance = OpenapiClient::PaymentRequestApi.new
payment_request_id = 'payment_request_id_example' # String | UUID of payment request to use.
file = File.new('/path/to/some/file') # File | The file to upload.
opts = {
  is_weight_ticket: true # Boolean | Indicates whether the file is a weight ticket.
}

begin
  # createUpload
  result = api_instance.create_upload(payment_request_id, file, opts)
  p result
rescue OpenapiClient::ApiError => e
  puts "Error when calling PaymentRequestApi->create_upload: #{e}"
end
```

#### Using the create_upload_with_http_info variant

This returns an Array which contains the response data, status code and headers.

> <Array(<UploadWithOmissions>, Integer, Hash)> create_upload_with_http_info(payment_request_id, file, opts)

```ruby
begin
  # createUpload
  data, status_code, headers = api_instance.create_upload_with_http_info(payment_request_id, file, opts)
  p status_code # => 2xx
  p headers # => { ... }
  p data # => <UploadWithOmissions>
rescue OpenapiClient::ApiError => e
  puts "Error when calling PaymentRequestApi->create_upload_with_http_info: #{e}"
end
```

### Parameters

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **payment_request_id** | **String** | UUID of payment request to use. |  |
| **file** | **File** | The file to upload. |  |
| **is_weight_ticket** | **Boolean** | Indicates whether the file is a weight ticket. | [optional] |

### Return type

[**UploadWithOmissions**](UploadWithOmissions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

