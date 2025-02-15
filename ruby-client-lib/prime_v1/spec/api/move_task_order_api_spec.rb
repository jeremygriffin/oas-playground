=begin
#MilMove Prime API

#The Prime API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v1/`. 

The version of the OpenAPI document: 0.0.1
Contact: milmove-developers@caci.com
Generated by: https://openapi-generator.tech
Generator version: 7.8.0-SNAPSHOT

=end

require 'spec_helper'
require 'json'

# Unit tests for OpenapiClient::MoveTaskOrderApi
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe 'MoveTaskOrderApi' do
  before do
    # run before each test
    @api_instance = OpenapiClient::MoveTaskOrderApi.new
  end

  after do
    # run after each test
  end

  describe 'test an instance of MoveTaskOrderApi' do
    it 'should create an instance of MoveTaskOrderApi' do
      expect(@api_instance).to be_instance_of(OpenapiClient::MoveTaskOrderApi)
    end
  end

  # unit tests for create_excess_weight_record
  # createExcessWeightRecord
  # Uploads an excess weight record, which is a document that proves that the movers or contractors have counseled the customer about their excess weight. Excess weight counseling should occur after the sum of the shipments for the customer&#39;s move crosses the excess weight alert threshold. 
  # @param move_task_order_id UUID of the move being updated.
  # @param file The file to upload.
  # @param [Hash] opts the optional parameters
  # @return [ExcessWeightRecord]
  describe 'create_excess_weight_record test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for download_move_order
  # Downloads move order as a PDF
  # ### Functionality This endpoint downloads all uploaded move order documentations into one download file by locator.  ### Errors * The move must be in need counseling state. * The move client&#39;s origin duty location must not currently have gov counseling. 
  # @param locator the locator code for move order to be downloaded
  # @param [Hash] opts the optional parameters
  # @option opts [String] :type upload type
  # @return [File]
  describe 'download_move_order test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for get_move_task_order
  # getMoveTaskOrder
  # ### Functionality This endpoint gets an individual MoveTaskOrder by ID.  It will provide information about the Customer and any associated MTOShipments, MTOServiceItems and PaymentRequests. 
  # @param move_id UUID or MoveCode of move task order to use.
  # @param [Hash] opts the optional parameters
  # @return [MoveTaskOrder]
  describe 'get_move_task_order test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for list_moves
  # listMoves
  # Gets all moves that have been reviewed and approved by the TOO. The &#x60;since&#x60; parameter can be used to filter this list down to only the moves that have been updated since the provided timestamp. A move will be considered updated if the &#x60;updatedAt&#x60; timestamp on the move or on its orders, shipments, service items, or payment requests, is later than the provided date and time.  **WIP**: Include what causes moves to leave this list. Currently, once the &#x60;availableToPrimeAt&#x60; timestamp has been set, that move will always appear in this list. 
  # @param [Hash] opts the optional parameters
  # @option opts [Time] :since Only return moves updated since this time. Formatted like \&quot;2021-07-23T18:30:47.116Z\&quot;
  # @return [Array<ListMove>]
  describe 'list_moves test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  # unit tests for update_mto_post_counseling_information
  # updateMTOPostCounselingInformation
  # ### Functionality This endpoint **updates** the MoveTaskOrder to indicate that the Prime has completed Counseling. This update uses the moveTaskOrderID provided in the path, updates the move status and marks child elements of the move to indicate the update. No body object is expected for this request.  **For Full/Partial PPMs**: This action is required so that the customer can start uploading their proof of service docs.  **For other move types**: This action is required for auditing reasons so that we have a record of when the Prime counseled the customer. 
  # @param move_task_order_id ID of move task order to use.
  # @param if_match Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error. 
  # @param [Hash] opts the optional parameters
  # @return [MoveTaskOrder]
  describe 'update_mto_post_counseling_information test' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

end
