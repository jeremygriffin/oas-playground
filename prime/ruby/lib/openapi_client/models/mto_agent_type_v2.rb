=begin
#MilMove Prime V2 API

#The Prime V2 API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v2/`. 

The version of the OpenAPI document: 0.0.1
Contact: milmove-developers@caci.com
Generated by: https://openapi-generator.tech
Generator version: 7.8.0-SNAPSHOT

=end

require 'date'
require 'time'

module OpenapiClient
  class MTOAgentTypeV2
    RELEASING_AGENT = "RELEASING_AGENT".freeze
    RECEIVING_AGENT = "RECEIVING_AGENT".freeze

    def self.all_vars
      @all_vars ||= [RELEASING_AGENT, RECEIVING_AGENT].freeze
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
      return value if MTOAgentTypeV2.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #MTOAgentTypeV2"
    end
  end
end
