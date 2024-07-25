=begin
#MilMove Prime V3 API

#The Prime V3 API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v3/`. 

The version of the OpenAPI document: 0.0.1
Contact: milmove-developers@caci.com
Generated by: https://openapi-generator.tech
Generator version: 7.8.0-SNAPSHOT

=end

# Common files
require 'openapi_client/api_client'
require 'openapi_client/api_error'
require 'openapi_client/version'
require 'openapi_client/configuration'

# Models
require 'openapi_client/models/address_v3'
require 'openapi_client/models/client_error_v3'
require 'openapi_client/models/create_mto_shipment_v3'
require 'openapi_client/models/create_ppm_shipment_v3'
require 'openapi_client/models/create_sit_address_update_request_v3'
require 'openapi_client/models/create_sit_extension_v3'
require 'openapi_client/models/customer_v3'
require 'openapi_client/models/destination_type_v3'
require 'openapi_client/models/duty_location_v3'
require 'openapi_client/models/entitlements_v3'
require 'openapi_client/models/error_v3'
require 'openapi_client/models/mto_agent_type_v3'
require 'openapi_client/models/mto_agent_v3'
require 'openapi_client/models/mto_service_item_dimension_v3'
require 'openapi_client/models/mto_service_item_model_type_v3'
require 'openapi_client/models/mto_service_item_status_v3'
require 'openapi_client/models/mto_service_item_v3'
require 'openapi_client/models/mto_shipment_type_v3'
require 'openapi_client/models/mto_shipment_v3'
require 'openapi_client/models/mto_shipment_without_service_items_v3'
require 'openapi_client/models/move_task_order_v3'
require 'openapi_client/models/order_v3'
require 'openapi_client/models/orders_type_v3'
require 'openapi_client/models/ppm_shipment_status_v3'
require 'openapi_client/models/ppm_shipment_v3'
require 'openapi_client/models/payment_request_status_v3'
require 'openapi_client/models/payment_request_v3'
require 'openapi_client/models/payment_service_item_param_v3'
require 'openapi_client/models/payment_service_item_status_v3'
require 'openapi_client/models/payment_service_item_v3'
require 'openapi_client/models/proof_of_service_doc_v3'
require 'openapi_client/models/re_service_code_v3'
require 'openapi_client/models/reweigh_requester_v3'
require 'openapi_client/models/reweigh_v3'
require 'openapi_client/models/sit_extension_v3'
require 'openapi_client/models/sit_location_type_v3'
require 'openapi_client/models/service_item_param_name_v3'
require 'openapi_client/models/service_item_param_origin_v3'
require 'openapi_client/models/service_item_param_type_v3'
require 'openapi_client/models/service_request_document_v3'
require 'openapi_client/models/shipment_address_update_status_v3'
require 'openapi_client/models/shipment_address_update_v3'
require 'openapi_client/models/sit_address_update_status_v3'
require 'openapi_client/models/sit_address_update_v3'
require 'openapi_client/models/storage_facility_v3'
require 'openapi_client/models/update_mto_service_item_model_type_v3'
require 'openapi_client/models/update_mto_service_item_v3'
require 'openapi_client/models/update_mto_shipment_status_v3'
require 'openapi_client/models/update_mto_shipment_storage_facility_v3'
require 'openapi_client/models/update_mto_shipment_v3'
require 'openapi_client/models/update_ppm_shipment_v3'
require 'openapi_client/models/update_reweigh_v3'
require 'openapi_client/models/update_shipment_destination_address_v3'
require 'openapi_client/models/upload_with_omissions_v3'
require 'openapi_client/models/validation_error_v3'
require 'openapi_client/models/mto_service_item_basic_v3'
require 'openapi_client/models/mto_service_item_dest_sitv3'
require 'openapi_client/models/mto_service_item_domestic_crating_v3'
require 'openapi_client/models/mto_service_item_origin_sitv3'
require 'openapi_client/models/mto_service_item_shuttle_v3'
require 'openapi_client/models/update_mto_service_item_sitv3'
require 'openapi_client/models/update_mto_service_item_shuttle_v3'

# APIs
require 'openapi_client/api/move_task_order_api'
require 'openapi_client/api/mto_shipment_api'

module OpenapiClient
  class << self
    # Customize default settings for the SDK using block.
    #   OpenapiClient.configure do |config|
    #     config.username = "xxx"
    #     config.password = "xxx"
    #   end
    # If no block given, return the default Configuration object.
    def configure
      if block_given?
        yield(Configuration.default)
      else
        Configuration.default
      end
    end
  end
end
