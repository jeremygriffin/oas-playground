=begin
#MilMove Prime API

#The Prime API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v1/`. 

The version of the OpenAPI document: 0.0.1
Contact: milmove-developers@caci.com
Generated by: https://openapi-generator.tech
Generator version: 7.8.0-SNAPSHOT

=end

require 'date'
require 'time'

module OpenapiClient
  class PaymentServiceItemStatus
    REQUESTED = "REQUESTED".freeze
    APPROVED = "APPROVED".freeze
    DENIED = "DENIED".freeze
    SENT_TO_GEX = "SENT_TO_GEX".freeze
    PAID = "PAID".freeze
    EDI_ERROR = "EDI_ERROR".freeze

    def self.all_vars
      @all_vars ||= [REQUESTED, APPROVED, DENIED, SENT_TO_GEX, PAID, EDI_ERROR].freeze
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
      return value if PaymentServiceItemStatus.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #PaymentServiceItemStatus"
    end
  end
end
