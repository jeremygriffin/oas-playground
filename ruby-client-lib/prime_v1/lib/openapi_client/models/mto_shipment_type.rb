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
  class MTOShipmentType
    BOAT_HAUL_AWAY = "BOAT_HAUL_AWAY".freeze
    BOAT_TOW_AWAY = "BOAT_TOW_AWAY".freeze
    HHG = "HHG".freeze
    HHG_INTO_NTS_DOMESTIC = "HHG_INTO_NTS_DOMESTIC".freeze
    HHG_OUTOF_NTS_DOMESTIC = "HHG_OUTOF_NTS_DOMESTIC".freeze
    INTERNATIONAL_HHG = "INTERNATIONAL_HHG".freeze
    INTERNATIONAL_UB = "INTERNATIONAL_UB".freeze
    MOTORHOME = "MOTORHOME".freeze
    PPM = "PPM".freeze

    def self.all_vars
      @all_vars ||= [BOAT_HAUL_AWAY, BOAT_TOW_AWAY, HHG, HHG_INTO_NTS_DOMESTIC, HHG_OUTOF_NTS_DOMESTIC, INTERNATIONAL_HHG, INTERNATIONAL_UB, MOTORHOME, PPM].freeze
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
      return value if MTOShipmentType.all_vars.include?(value)
      raise "Invalid ENUM value #{value} for class #MTOShipmentType"
    end
  end
end
