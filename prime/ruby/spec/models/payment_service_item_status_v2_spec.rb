=begin
#MilMove Prime V2 API

#The Prime V2 API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v2/`. 

The version of the OpenAPI document: 0.0.1
Contact: milmove-developers@caci.com
Generated by: https://openapi-generator.tech
Generator version: 7.8.0-SNAPSHOT

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for OpenapiClient::PaymentServiceItemStatusV2
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe OpenapiClient::PaymentServiceItemStatusV2 do
  let(:instance) { OpenapiClient::PaymentServiceItemStatusV2.new }

  describe 'test an instance of PaymentServiceItemStatusV2' do
    it 'should create an instance of PaymentServiceItemStatusV2' do
      # uncomment below to test the instance creation
      #expect(instance).to be_instance_of(OpenapiClient::PaymentServiceItemStatusV2)
    end
  end

end
