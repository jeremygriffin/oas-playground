=begin
#MilMove Prime V3 API

#The Prime V3 API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v3/`. 

The version of the OpenAPI document: 0.0.1
Contact: milmove-developers@caci.com
Generated by: https://openapi-generator.tech
Generator version: 7.8.0-SNAPSHOT

=end

require 'spec_helper'
require 'json'
require 'date'

# Unit tests for OpenapiClient::PaymentRequestStatusV3
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe OpenapiClient::PaymentRequestStatusV3 do
  let(:instance) { OpenapiClient::PaymentRequestStatusV3.new }

  describe 'test an instance of PaymentRequestStatusV3' do
    it 'should create an instance of PaymentRequestStatusV3' do
      # uncomment below to test the instance creation
      #expect(instance).to be_instance_of(OpenapiClient::PaymentRequestStatusV3)
    end
  end

end
