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

# Unit tests for OpenapiClient::CreateSITExtensionV3
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe OpenapiClient::CreateSITExtensionV3 do
  let(:instance) { OpenapiClient::CreateSITExtensionV3.new }

  describe 'test an instance of CreateSITExtensionV3' do
    it 'should create an instance of CreateSITExtensionV3' do
      # uncomment below to test the instance creation
      #expect(instance).to be_instance_of(OpenapiClient::CreateSITExtensionV3)
    end
  end

  describe 'test attribute "request_reason"' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
      # validator = Petstore::EnumTest::EnumAttributeValidator.new('String', ["SERIOUS_ILLNESS_MEMBER", "SERIOUS_ILLNESS_DEPENDENT", "IMPENDING_ASSIGNEMENT", "DIRECTED_TEMPORARY_DUTY", "NONAVAILABILITY_OF_CIVILIAN_HOUSING", "AWAITING_COMPLETION_OF_RESIDENCE", "OTHER"])
      # validator.allowable_values.each do |value|
      #   expect { instance.request_reason = value }.not_to raise_error
      # end
    end
  end

  describe 'test attribute "contractor_remarks"' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  describe 'test attribute "requested_days"' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

end
