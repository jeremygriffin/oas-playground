=begin
#MilMove Prime V3 API

#The Prime V3 API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v3/`. 

The version of the OpenAPI document: 0.0.1
Contact: milmove-developers@caci.com
Generated by: https://openapi-generator.tech
Generator version: 7.8.0-SNAPSHOT

=end

require 'date'
require 'time'

module OpenapiClient
  class ServiceItemParamOriginV3
    PRIME = "PRIME".freeze
    SYSTEM = "SYSTEM".freeze
    PRICER = "PRICER".freeze
    PAYMENT_REQUEST = "PAYMENT_REQUEST".freeze

    def self.all_vars
      @all_vars ||= [PRIME, SYSTEM, PRICER, PAYMENT_REQUEST].freeze
    end

    # Builds the enum from string
    # @param [String] The enum value in the form of the string
    # @return [String] The enum value
    def self.build_from_hash(value)
      new.build_from_hash(value)
    end

    # Builds the enum from string
    # @param [String] The enum value in the form of the string
    # @return [String] The enum value
    def build_from_hash(value)
      return value if ServiceItemParamOriginV3.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #ServiceItemParamOriginV3"
    end
  end
end
