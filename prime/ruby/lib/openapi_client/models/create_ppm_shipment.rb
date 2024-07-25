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
  # Creation object containing the `PPM` shipmentType specific data, not used for other shipment types.
  class CreatePPMShipment
    # Date the customer expects to begin moving from their origin. 
    attr_accessor :expected_departure_date

    # The address of the origin location where goods are being moved from.
    attr_accessor :pickup_address

    # The address of the destination location where goods are being delivered to.
    attr_accessor :destination_address

    # Captures whether some or all of the PPM shipment will require temporary storage at the origin or destination.  Must be set to `true` when providing `sitLocation`, `sitEstimatedWeight`, `sitEstimatedEntryDate`, and `sitEstimatedDepartureDate` values to calculate the `sitEstimatedCost`. 
    attr_accessor :sit_expected

    attr_accessor :sit_location

    # The estimated weight of the goods being put into storage in pounds.
    attr_accessor :sit_estimated_weight

    # The date that goods will first enter the storage location.
    attr_accessor :sit_estimated_entry_date

    # The date that goods will exit the storage location.
    attr_accessor :sit_estimated_departure_date

    # The estimated weight of the PPM shipment goods being moved in pounds.
    attr_accessor :estimated_weight

    # Indicates whether PPM shipment has pro gear for themselves or their spouse. 
    attr_accessor :has_pro_gear

    # The estimated weight of the pro-gear being moved belonging to the service member in pounds.
    attr_accessor :pro_gear_weight

    # The estimated weight of the pro-gear being moved belonging to a spouse in pounds.
    attr_accessor :spouse_pro_gear_weight

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'expected_departure_date' => :'expectedDepartureDate',
        :'pickup_address' => :'pickupAddress',
        :'destination_address' => :'destinationAddress',
        :'sit_expected' => :'sitExpected',
        :'sit_location' => :'sitLocation',
        :'sit_estimated_weight' => :'sitEstimatedWeight',
        :'sit_estimated_entry_date' => :'sitEstimatedEntryDate',
        :'sit_estimated_departure_date' => :'sitEstimatedDepartureDate',
        :'estimated_weight' => :'estimatedWeight',
        :'has_pro_gear' => :'hasProGear',
        :'pro_gear_weight' => :'proGearWeight',
        :'spouse_pro_gear_weight' => :'spouseProGearWeight'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'expected_departure_date' => :'Date',
        :'pickup_address' => :'Address',
        :'destination_address' => :'Address',
        :'sit_expected' => :'Boolean',
        :'sit_location' => :'SITLocationType',
        :'sit_estimated_weight' => :'Integer',
        :'sit_estimated_entry_date' => :'Date',
        :'sit_estimated_departure_date' => :'Date',
        :'estimated_weight' => :'Integer',
        :'has_pro_gear' => :'Boolean',
        :'pro_gear_weight' => :'Integer',
        :'spouse_pro_gear_weight' => :'Integer'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'sit_estimated_weight',
        :'sit_estimated_entry_date',
        :'sit_estimated_departure_date',
        :'pro_gear_weight',
        :'spouse_pro_gear_weight'
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `OpenapiClient::CreatePPMShipment` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `OpenapiClient::CreatePPMShipment`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'expected_departure_date')
        self.expected_departure_date = attributes[:'expected_departure_date']
      else
        self.expected_departure_date = nil
      end

      if attributes.key?(:'pickup_address')
        self.pickup_address = attributes[:'pickup_address']
      end

      if attributes.key?(:'destination_address')
        self.destination_address = attributes[:'destination_address']
      end

      if attributes.key?(:'sit_expected')
        self.sit_expected = attributes[:'sit_expected']
      else
        self.sit_expected = nil
      end

      if attributes.key?(:'sit_location')
        if (value = attributes[:'sit_location']).is_a?(Hash)
          self.sit_location = value
        end
        self.sit_location = attributes[:'sit_location']
      end

      if attributes.key?(:'sit_estimated_weight')
        self.sit_estimated_weight = attributes[:'sit_estimated_weight']
      end

      if attributes.key?(:'sit_estimated_entry_date')
        self.sit_estimated_entry_date = attributes[:'sit_estimated_entry_date']
      end

      if attributes.key?(:'sit_estimated_departure_date')
        self.sit_estimated_departure_date = attributes[:'sit_estimated_departure_date']
      end

      if attributes.key?(:'estimated_weight')
        self.estimated_weight = attributes[:'estimated_weight']
      else
        self.estimated_weight = nil
      end

      if attributes.key?(:'has_pro_gear')
        self.has_pro_gear = attributes[:'has_pro_gear']
      else
        self.has_pro_gear = nil
      end

      if attributes.key?(:'pro_gear_weight')
        self.pro_gear_weight = attributes[:'pro_gear_weight']
      end

      if attributes.key?(:'spouse_pro_gear_weight')
        self.spouse_pro_gear_weight = attributes[:'spouse_pro_gear_weight']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      warn '[DEPRECATED] the `list_invalid_properties` method is obsolete'
      invalid_properties = Array.new
      if @expected_departure_date.nil?
        invalid_properties.push('invalid value for "expected_departure_date", expected_departure_date cannot be nil.')
      end

      if @sit_expected.nil?
        invalid_properties.push('invalid value for "sit_expected", sit_expected cannot be nil.')
      end

      if @estimated_weight.nil?
        invalid_properties.push('invalid value for "estimated_weight", estimated_weight cannot be nil.')
      end

      if @has_pro_gear.nil?
        invalid_properties.push('invalid value for "has_pro_gear", has_pro_gear cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      warn '[DEPRECATED] the `valid?` method is obsolete'
      return false if @expected_departure_date.nil?
      return false if @sit_expected.nil?
      return false if @estimated_weight.nil?
      return false if @has_pro_gear.nil?
      true
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          expected_departure_date == o.expected_departure_date &&
          pickup_address == o.pickup_address &&
          destination_address == o.destination_address &&
          sit_expected == o.sit_expected &&
          sit_location == o.sit_location &&
          sit_estimated_weight == o.sit_estimated_weight &&
          sit_estimated_entry_date == o.sit_estimated_entry_date &&
          sit_estimated_departure_date == o.sit_estimated_departure_date &&
          estimated_weight == o.estimated_weight &&
          has_pro_gear == o.has_pro_gear &&
          pro_gear_weight == o.pro_gear_weight &&
          spouse_pro_gear_weight == o.spouse_pro_gear_weight
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [expected_departure_date, pickup_address, destination_address, sit_expected, sit_location, sit_estimated_weight, sit_estimated_entry_date, sit_estimated_departure_date, estimated_weight, has_pro_gear, pro_gear_weight, spouse_pro_gear_weight].hash
    end

    # Builds the object from hash
    # @param [Hash] attributes Model attributes in the form of hash
    # @return [Object] Returns the model itself
    def self.build_from_hash(attributes)
      return nil unless attributes.is_a?(Hash)
      attributes = attributes.transform_keys(&:to_sym)
      transformed_hash = {}
      openapi_types.each_pair do |key, type|
        if attributes.key?(attribute_map[key]) && attributes[attribute_map[key]].nil?
          transformed_hash["#{key}"] = nil
        elsif type =~ /\AArray<(.*)>/i
          # check to ensure the input is an array given that the attribute
          # is documented as an array but the input is not
          if attributes[attribute_map[key]].is_a?(Array)
            transformed_hash["#{key}"] = attributes[attribute_map[key]].map { |v| _deserialize($1, v) }
          end
        elsif !attributes[attribute_map[key]].nil?
          transformed_hash["#{key}"] = _deserialize(type, attributes[attribute_map[key]])
        end
      end
      new(transformed_hash)
    end

    # Deserializes the data based on type
    # @param string type Data type
    # @param string value Value to be deserialized
    # @return [Object] Deserialized data
    def self._deserialize(type, value)
      case type.to_sym
      when :Time
        Time.parse(value)
      when :Date
        Date.parse(value)
      when :String
        value.to_s
      when :Integer
        value.to_i
      when :Float
        value.to_f
      when :Boolean
        if value.to_s =~ /\A(true|t|yes|y|1)\z/i
          true
        else
          false
        end
      when :Object
        # generic object (usually a Hash), return directly
        value
      when /\AArray<(?<inner_type>.+)>\z/
        inner_type = Regexp.last_match[:inner_type]
        value.map { |v| _deserialize(inner_type, v) }
      when /\AHash<(?<k_type>.+?), (?<v_type>.+)>\z/
        k_type = Regexp.last_match[:k_type]
        v_type = Regexp.last_match[:v_type]
        {}.tap do |hash|
          value.each do |k, v|
            hash[_deserialize(k_type, k)] = _deserialize(v_type, v)
          end
        end
      else # model
        # models (e.g. Pet) or oneOf
        klass = OpenapiClient.const_get(type)
        klass.respond_to?(:openapi_any_of) || klass.respond_to?(:openapi_one_of) ? klass.build(value) : klass.build_from_hash(value)
      end
    end

    # Returns the string representation of the object
    # @return [String] String presentation of the object
    def to_s
      to_hash.to_s
    end

    # to_body is an alias to to_hash (backward compatibility)
    # @return [Hash] Returns the object in the form of hash
    def to_body
      to_hash
    end

    # Returns the object in the form of hash
    # @return [Hash] Returns the object in the form of hash
    def to_hash
      hash = {}
      self.class.attribute_map.each_pair do |attr, param|
        value = self.send(attr)
        if value.nil?
          is_nullable = self.class.openapi_nullable.include?(attr)
          next if !is_nullable || (is_nullable && !instance_variable_defined?(:"@#{attr}"))
        end

        hash[param] = _to_hash(value)
      end
      hash
    end

    # Outputs non-array value in the form of hash
    # For object, use to_hash. Otherwise, just return the value
    # @param [Object] value Any valid value
    # @return [Hash] Returns the value in the form of hash
    def _to_hash(value)
      if value.is_a?(Array)
        value.compact.map { |v| _to_hash(v) }
      elsif value.is_a?(Hash)
        {}.tap do |hash|
          value.each { |k, v| hash[k] = _to_hash(v) }
        end
      elsif value.respond_to? :to_hash
        value.to_hash
      else
        value
      end
    end

  end

end
