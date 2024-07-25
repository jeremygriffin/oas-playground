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

# Unit tests for OpenapiClient::UpdateMTOServiceItemShuttle
# Automatically generated by openapi-generator (https://openapi-generator.tech)
# Please update as you see appropriate
describe OpenapiClient::UpdateMTOServiceItemShuttle do
  let(:instance) { OpenapiClient::UpdateMTOServiceItemShuttle.new }

  describe 'test an instance of UpdateMTOServiceItemShuttle' do
    it 'should create an instance of UpdateMTOServiceItemShuttle' do
      # uncomment below to test the instance creation
      #expect(instance).to be_instance_of(OpenapiClient::UpdateMTOServiceItemShuttle)
    end
  end

  describe 'test attribute "actual_weight"' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  describe 'test attribute "estimated_weight"' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
    end
  end

  describe 'test attribute "re_service_code"' do
    it 'should work' do
      # assertion here. ref: https://rspec.info/features/3-12/rspec-expectations/built-in-matchers/
      # validator = Petstore::EnumTest::EnumAttributeValidator.new('String', ["DDSHUT", "DOSHUT"])
      # validator.allowable_values.each do |value|
      #   expect { instance.re_service_code = value }.not_to raise_error
      # end
    end
  end

end
