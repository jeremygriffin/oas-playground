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
  class MTOServiceItemModelTypeV2
    MTO_SERVICE_ITEM_BASIC = "MTOServiceItemBasic".freeze
    MTO_SERVICE_ITEM_ORIGIN_SIT = "MTOServiceItemOriginSIT".freeze
    MTO_SERVICE_ITEM_DEST_SIT = "MTOServiceItemDestSIT".freeze
    MTO_SERVICE_ITEM_SHUTTLE = "MTOServiceItemShuttle".freeze
    MTO_SERVICE_ITEM_DOMESTIC_CRATING = "MTOServiceItemDomesticCrating".freeze

    def self.all_vars
      @all_vars ||= [MTO_SERVICE_ITEM_BASIC, MTO_SERVICE_ITEM_ORIGIN_SIT, MTO_SERVICE_ITEM_DEST_SIT, MTO_SERVICE_ITEM_SHUTTLE, MTO_SERVICE_ITEM_DOMESTIC_CRATING].freeze
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
      return value if MTOServiceItemModelTypeV2.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #MTOServiceItemModelTypeV2"
    end
  end
end
