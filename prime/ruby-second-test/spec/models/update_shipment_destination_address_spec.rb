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
require 'date'

# Unit tests for OpenapiClient::UpdateShipmentDestinationAddress
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe OpenapiClient::UpdateShipmentDestinationAddress do
  let(:instance) { OpenapiClient::UpdateShipmentDestinationAddress.new }

  describe 'test an instance of UpdateShipmentDestinationAddress' do
    it 'should create an instance of UpdateShipmentDestinationAddress' do
      # uncomment below to test the instance creation
      #expect(instance).to be_instance_of(OpenapiClient::UpdateShipmentDestinationAddress)
    end
  end

  describe 'test attribute "new_address"' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  describe 'test attribute "contractor_remarks"' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

end
