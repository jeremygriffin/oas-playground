=begin
#MilMove Prime API

#The Prime API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v1/`. 

The version of the OpenAPI document: 0.0.1
Contact: milmove-developers@caci.com
Generated by: https://openapi-generator.tech
Generator version: 7.8.0-SNAPSHOT

=end

require 'cgi'

module OpenapiClient
  class MtoServiceItemApi
    attr_accessor :api_client

    def initialize(api_client = ApiClient.default)
      @api_client = api_client
    end
    # createMTOServiceItem
    # Creates one or more MTOServiceItems. Not all service items may be created, please see details below.  This endpoint supports different body definitions. In the modelType field below, select the modelType corresponding  to the service item you wish to create and the documentation will update with the new definition.  Upon creation these items are associated with a Move Task Order and an MTO Shipment. The request must include UUIDs for the MTO and MTO Shipment connected to this service item. Some service item types require additional service items to be autogenerated when added - all created service items, autogenerated included, will be returned in the response.  To update a service item, please use [updateMTOServiceItem](#operation/updateMTOServiceItem) endpoint.  ---  **`MTOServiceItemOriginSIT`**  MTOServiceItemOriginSIT is a subtype of MTOServiceItem.  This model type describes a domestic origin SIT service item. Items can be created using this model type with the following codes:  **DOFSIT**  **1st day origin SIT service item**. When a DOFSIT is requested, the API will auto-create the following group of service items:   * DOFSIT - Domestic origin 1st day SIT   * DOASIT - Domestic origin Additional day SIT   * DOPSIT - Domestic origin SIT pickup   * DOSFSC - Domestic origin SIT fuel surcharge  **DOASIT**  **Addt'l days origin SIT service item**. This represents an additional day of storage for the same item. Additional DOASIT service items can be created and added to an existing shipment that **includes a DOFSIT service item**.  ---  **`MTOServiceItemDestSIT`**  MTOServiceItemDestSIT is a subtype of MTOServiceItem.  This model type describes a domestic destination SIT service item. Items can be created using this model type with the following codes:  **DDFSIT**  **1st day destination SIT service item**.  These additional fields are optional for creating a DDFSIT:   * `firstAvailableDeliveryDate1`     * string <date>     * First available date that Prime can deliver SIT service item.     * firstAvailableDeliveryDate1, dateOfContact1, and timeMilitary1 are required together   * `dateOfContact1`     * string <date>     * Date of attempted contact by the prime corresponding to `timeMilitary1`     * dateOfContact1, timeMilitary1, and firstAvailableDeliveryDate1 are required together   * `timeMilitary1`     * string\\d{4}Z     * Time of attempted contact corresponding to `dateOfContact1`, in military format.     * timeMilitary1, dateOfContact1, and firstAvailableDeliveryDate1 are required together   * `firstAvailableDeliveryDate2`     * string <date>     * Second available date that Prime can deliver SIT service item.     * firstAvailableDeliveryDate2, dateOfContact2, and timeMilitary2 are required together   * `dateOfContact2`     * string <date>     * Date of attempted contact delivery by the prime corresponding to `timeMilitary2`     * dateOfContact2, timeMilitary2, and firstAvailableDeliveryDate2 are required together   * `timeMilitary2`     * string\\d{4}Z     * Time of attempted contact corresponding to `dateOfContact2`, in military format.     * timeMilitary2, dateOfContact2, and firstAvailableDeliveryDate2 are required together  When a DDFSIT is requested, the API will auto-create the following group of service items:   * DDFSIT - Domestic destination 1st day SIT   * DDASIT - Domestic destination Additional day SIT   * DDDSIT - Domestic destination SIT delivery   * DDSFSC - Domestic destination SIT fuel surcharge  **DDASIT**  **Addt'l days destination SIT service item**. This represents an additional day of storage for the same item. Additional DDASIT service items can be created and added to an existing shipment that **includes a DDFSIT service item**. 
    # @param [Hash] opts the optional parameters
    # @option opts [MTOServiceItem] :body 
    # @return [Array<MTOServiceItem>]
    def create_mto_service_item(opts = {})
      data, _status_code, _headers = create_mto_service_item_with_http_info(opts)
      data
    end

    # createMTOServiceItem
    # Creates one or more MTOServiceItems. Not all service items may be created, please see details below.  This endpoint supports different body definitions. In the modelType field below, select the modelType corresponding  to the service item you wish to create and the documentation will update with the new definition.  Upon creation these items are associated with a Move Task Order and an MTO Shipment. The request must include UUIDs for the MTO and MTO Shipment connected to this service item. Some service item types require additional service items to be autogenerated when added - all created service items, autogenerated included, will be returned in the response.  To update a service item, please use [updateMTOServiceItem](#operation/updateMTOServiceItem) endpoint.  ---  **&#x60;MTOServiceItemOriginSIT&#x60;**  MTOServiceItemOriginSIT is a subtype of MTOServiceItem.  This model type describes a domestic origin SIT service item. Items can be created using this model type with the following codes:  **DOFSIT**  **1st day origin SIT service item**. When a DOFSIT is requested, the API will auto-create the following group of service items:   * DOFSIT - Domestic origin 1st day SIT   * DOASIT - Domestic origin Additional day SIT   * DOPSIT - Domestic origin SIT pickup   * DOSFSC - Domestic origin SIT fuel surcharge  **DOASIT**  **Addt&#39;l days origin SIT service item**. This represents an additional day of storage for the same item. Additional DOASIT service items can be created and added to an existing shipment that **includes a DOFSIT service item**.  ---  **&#x60;MTOServiceItemDestSIT&#x60;**  MTOServiceItemDestSIT is a subtype of MTOServiceItem.  This model type describes a domestic destination SIT service item. Items can be created using this model type with the following codes:  **DDFSIT**  **1st day destination SIT service item**.  These additional fields are optional for creating a DDFSIT:   * &#x60;firstAvailableDeliveryDate1&#x60;     * string &lt;date&gt;     * First available date that Prime can deliver SIT service item.     * firstAvailableDeliveryDate1, dateOfContact1, and timeMilitary1 are required together   * &#x60;dateOfContact1&#x60;     * string &lt;date&gt;     * Date of attempted contact by the prime corresponding to &#x60;timeMilitary1&#x60;     * dateOfContact1, timeMilitary1, and firstAvailableDeliveryDate1 are required together   * &#x60;timeMilitary1&#x60;     * string\\d{4}Z     * Time of attempted contact corresponding to &#x60;dateOfContact1&#x60;, in military format.     * timeMilitary1, dateOfContact1, and firstAvailableDeliveryDate1 are required together   * &#x60;firstAvailableDeliveryDate2&#x60;     * string &lt;date&gt;     * Second available date that Prime can deliver SIT service item.     * firstAvailableDeliveryDate2, dateOfContact2, and timeMilitary2 are required together   * &#x60;dateOfContact2&#x60;     * string &lt;date&gt;     * Date of attempted contact delivery by the prime corresponding to &#x60;timeMilitary2&#x60;     * dateOfContact2, timeMilitary2, and firstAvailableDeliveryDate2 are required together   * &#x60;timeMilitary2&#x60;     * string\\d{4}Z     * Time of attempted contact corresponding to &#x60;dateOfContact2&#x60;, in military format.     * timeMilitary2, dateOfContact2, and firstAvailableDeliveryDate2 are required together  When a DDFSIT is requested, the API will auto-create the following group of service items:   * DDFSIT - Domestic destination 1st day SIT   * DDASIT - Domestic destination Additional day SIT   * DDDSIT - Domestic destination SIT delivery   * DDSFSC - Domestic destination SIT fuel surcharge  **DDASIT**  **Addt&#39;l days destination SIT service item**. This represents an additional day of storage for the same item. Additional DDASIT service items can be created and added to an existing shipment that **includes a DDFSIT service item**. 
    # @param [Hash] opts the optional parameters
    # @option opts [MTOServiceItem] :body 
    # @return [Array<(Array<MTOServiceItem>, Integer, Hash)>] Array<MTOServiceItem> data, response status code and response headers
    def create_mto_service_item_with_http_info(opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: MtoServiceItemApi.create_mto_service_item ...'
      end
      # resource path
      local_var_path = '/mto-service-items'

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json']) unless header_params['Accept']
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(opts[:'body'])

      # return_type
      return_type = opts[:debug_return_type] || 'Array<MTOServiceItem>'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"MtoServiceItemApi.create_mto_service_item",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: MtoServiceItemApi#create_mto_service_item\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # createServiceRequestDocumentUpload
    # ### Functionality  This endpoint **uploads** a Service Request document for a ServiceItem.  The ServiceItem should already exist.  ServiceItems are created with the [createMTOServiceItem](#operation/createMTOServiceItem) endpoint. 
    # @param mto_service_item_id [String] UUID of the service item to use.
    # @param file [File] The file to upload.
    # @param [Hash] opts the optional parameters
    # @return [UploadWithOmissions]
    def create_service_request_document_upload(mto_service_item_id, file, opts = {})
      data, _status_code, _headers = create_service_request_document_upload_with_http_info(mto_service_item_id, file, opts)
      data
    end

    # createServiceRequestDocumentUpload
    # ### Functionality  This endpoint **uploads** a Service Request document for a ServiceItem.  The ServiceItem should already exist.  ServiceItems are created with the [createMTOServiceItem](#operation/createMTOServiceItem) endpoint. 
    # @param mto_service_item_id [String] UUID of the service item to use.
    # @param file [File] The file to upload.
    # @param [Hash] opts the optional parameters
    # @return [Array<(UploadWithOmissions, Integer, Hash)>] UploadWithOmissions data, response status code and response headers
    def create_service_request_document_upload_with_http_info(mto_service_item_id, file, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: MtoServiceItemApi.create_service_request_document_upload ...'
      end
      # verify the required parameter 'mto_service_item_id' is set
      if @api_client.config.client_side_validation && mto_service_item_id.nil?
        fail ArgumentError, "Missing the required parameter 'mto_service_item_id' when calling MtoServiceItemApi.create_service_request_document_upload"
      end
      # verify the required parameter 'file' is set
      if @api_client.config.client_side_validation && file.nil?
        fail ArgumentError, "Missing the required parameter 'file' when calling MtoServiceItemApi.create_service_request_document_upload"
      end
      # resource path
      local_var_path = '/mto-service-items/{mtoServiceItemID}/uploads'.sub('{' + 'mtoServiceItemID' + '}', CGI.escape(mto_service_item_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json']) unless header_params['Accept']
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['multipart/form-data'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end

      # form parameters
      form_params = opts[:form_params] || {}
      form_params['file'] = file

      # http body (model)
      post_body = opts[:debug_body]

      # return_type
      return_type = opts[:debug_return_type] || 'UploadWithOmissions'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"MtoServiceItemApi.create_service_request_document_upload",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:POST, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: MtoServiceItemApi#create_service_request_document_upload\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end

    # updateMTOServiceItem
    # Updates MTOServiceItems after creation. Not all service items or fields may be updated, please see details below.  This endpoint supports different body definitions. In the modelType field below, select the modelType corresponding  to the service item you wish to update and the documentation will update with the new definition.  * Addresses: To update a destination service item's SIT destination final address, update the shipment destination address. For approved shipments, please use [updateShipmentDestinationAddress](#mtoShipment/updateShipmentDestinationAddress). For shipments not yet approved, please use [updateMTOShipmentAddress](#mtoShipment/updateMTOShipmentAddress).  * SIT Service Items: Take note that when updating `sitCustomerContacted`, `sitDepartureDate`, or `sitRequestedDelivery`, we want those to be updated on `DOASIT` (for origin SIT) and `DDASIT` (for destination SIT). If updating those values in other service items, the office users will not have as much attention to those values.  To create a service item, please use [createMTOServiceItem](#mtoServiceItem/createMTOServiceItem)) endpoint. 
    # @param mto_service_item_id [String] UUID of service item to update.
    # @param if_match [String] Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error. 
    # @param body [UpdateMTOServiceItem] 
    # @param [Hash] opts the optional parameters
    # @return [MTOServiceItem]
    def update_mto_service_item(mto_service_item_id, if_match, body, opts = {})
      data, _status_code, _headers = update_mto_service_item_with_http_info(mto_service_item_id, if_match, body, opts)
      data
    end

    # updateMTOServiceItem
    # Updates MTOServiceItems after creation. Not all service items or fields may be updated, please see details below.  This endpoint supports different body definitions. In the modelType field below, select the modelType corresponding  to the service item you wish to update and the documentation will update with the new definition.  * Addresses: To update a destination service item&#39;s SIT destination final address, update the shipment destination address. For approved shipments, please use [updateShipmentDestinationAddress](#mtoShipment/updateShipmentDestinationAddress). For shipments not yet approved, please use [updateMTOShipmentAddress](#mtoShipment/updateMTOShipmentAddress).  * SIT Service Items: Take note that when updating &#x60;sitCustomerContacted&#x60;, &#x60;sitDepartureDate&#x60;, or &#x60;sitRequestedDelivery&#x60;, we want those to be updated on &#x60;DOASIT&#x60; (for origin SIT) and &#x60;DDASIT&#x60; (for destination SIT). If updating those values in other service items, the office users will not have as much attention to those values.  To create a service item, please use [createMTOServiceItem](#mtoServiceItem/createMTOServiceItem)) endpoint. 
    # @param mto_service_item_id [String] UUID of service item to update.
    # @param if_match [String] Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error. 
    # @param body [UpdateMTOServiceItem] 
    # @param [Hash] opts the optional parameters
    # @return [Array<(MTOServiceItem, Integer, Hash)>] MTOServiceItem data, response status code and response headers
    def update_mto_service_item_with_http_info(mto_service_item_id, if_match, body, opts = {})
      if @api_client.config.debugging
        @api_client.config.logger.debug 'Calling API: MtoServiceItemApi.update_mto_service_item ...'
      end
      # verify the required parameter 'mto_service_item_id' is set
      if @api_client.config.client_side_validation && mto_service_item_id.nil?
        fail ArgumentError, "Missing the required parameter 'mto_service_item_id' when calling MtoServiceItemApi.update_mto_service_item"
      end
      # verify the required parameter 'if_match' is set
      if @api_client.config.client_side_validation && if_match.nil?
        fail ArgumentError, "Missing the required parameter 'if_match' when calling MtoServiceItemApi.update_mto_service_item"
      end
      # verify the required parameter 'body' is set
      if @api_client.config.client_side_validation && body.nil?
        fail ArgumentError, "Missing the required parameter 'body' when calling MtoServiceItemApi.update_mto_service_item"
      end
      # resource path
      local_var_path = '/mto-service-items/{mtoServiceItemID}'.sub('{' + 'mtoServiceItemID' + '}', CGI.escape(mto_service_item_id.to_s))

      # query parameters
      query_params = opts[:query_params] || {}

      # header parameters
      header_params = opts[:header_params] || {}
      # HTTP header 'Accept' (if needed)
      header_params['Accept'] = @api_client.select_header_accept(['application/json']) unless header_params['Accept']
      # HTTP header 'Content-Type'
      content_type = @api_client.select_header_content_type(['application/json'])
      if !content_type.nil?
          header_params['Content-Type'] = content_type
      end
      header_params[:'If-Match'] = if_match

      # form parameters
      form_params = opts[:form_params] || {}

      # http body (model)
      post_body = opts[:debug_body] || @api_client.object_to_http_body(body)

      # return_type
      return_type = opts[:debug_return_type] || 'MTOServiceItem'

      # auth_names
      auth_names = opts[:debug_auth_names] || []

      new_options = opts.merge(
        :operation => :"MtoServiceItemApi.update_mto_service_item",
        :header_params => header_params,
        :query_params => query_params,
        :form_params => form_params,
        :body => post_body,
        :auth_names => auth_names,
        :return_type => return_type
      )

      data, status_code, headers = @api_client.call_api(:PATCH, local_var_path, new_options)
      if @api_client.config.debugging
        @api_client.config.logger.debug "API called: MtoServiceItemApi#update_mto_service_item\nData: #{data.inspect}\nStatus code: #{status_code}\nHeaders: #{headers}"
      end
      return data, status_code, headers
    end
  end
end
