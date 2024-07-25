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
  class UpdateMTOShipmentV2
    # The date the Prime contractor scheduled to pick up this shipment after consultation with the customer.
    attr_accessor :scheduled_pickup_date

    # The date when the Prime contractor actually picked up the shipment. Updated after-the-fact.
    attr_accessor :actual_pickup_date

    # The date the Prime provides to the customer as the first possible delivery date so that they can plan their travel accordingly. 
    attr_accessor :first_available_delivery_date

    # The date the Prime contractor scheduled to deliver this shipment after consultation with the customer.
    attr_accessor :scheduled_delivery_date

    # The date when the Prime contractor actually delivered the shipment. Updated after-the-fact.
    attr_accessor :actual_delivery_date

    # The estimated weight of this shipment, determined by the movers during the pre-move survey. This value **can only be updated once.** If there was an issue with estimating the weight and a mistake was made, the Prime contracter will need to contact the TOO to change it. 
    attr_accessor :prime_estimated_weight

    # The actual weight of the shipment, provided after the Prime packs, picks up, and weighs a customer's shipment.
    attr_accessor :prime_actual_weight

    # The previously recorded weight for the NTS Shipment. Used for NTS Release to know what the previous primeActualWeight or billable weight was.
    attr_accessor :nts_recorded_weight

    # The address where the movers should pick up this shipment, entered by the customer during onboarding when they enter shipment details. 
    attr_accessor :pickup_address

    # Where the movers should deliver this shipment. Often provided by the customer when they enter shipment details during onboarding, if they know their new address already.  May be blank when entered by the customer, required when entered by the Prime. May not represent the true final destination due to the shipment being diverted or placed in SIT. 
    attr_accessor :destination_address

    attr_accessor :destination_type

    # A second pickup address for this shipment, if the customer entered one. An optional field.
    attr_accessor :secondary_pickup_address

    # A second delivery address for this shipment, if the customer entered one. An optional field.
    attr_accessor :secondary_delivery_address

    attr_accessor :storage_facility

    attr_accessor :shipment_type

    # This value indicates whether or not this shipment is part of a diversion. If yes, the shipment can be either the starting or ending segment of the diversion. 
    attr_accessor :diversion

    # Email or ID of the person who will be contacted in the event of questions or concerns about this update. May be the person performing the update, or someone else working with the Prime contractor. 
    attr_accessor :point_of_contact

    attr_accessor :counselor_remarks

    attr_accessor :actual_pro_gear_weight

    attr_accessor :actual_spouse_pro_gear_weight

    attr_accessor :ppm_shipment

    class EnumAttributeValidator
      attr_reader :datatype
      attr_reader :allowable_values

      def initialize(datatype, allowable_values)
        @allowable_values = allowable_values.map do |value|
          case datatype.to_s
          when /Integer/i
            value.to_i
          when /Float/i
            value.to_f
          else
            value
          end
        end
      end

      def valid?(value)
        !value || allowable_values.include?(value)
      end
    end

    # Attribute mapping from ruby-style variable name to JSON key.
    def self.attribute_map
      {
        :'scheduled_pickup_date' => :'scheduledPickupDate',
        :'actual_pickup_date' => :'actualPickupDate',
        :'first_available_delivery_date' => :'firstAvailableDeliveryDate',
        :'scheduled_delivery_date' => :'scheduledDeliveryDate',
        :'actual_delivery_date' => :'actualDeliveryDate',
        :'prime_estimated_weight' => :'primeEstimatedWeight',
        :'prime_actual_weight' => :'primeActualWeight',
        :'nts_recorded_weight' => :'ntsRecordedWeight',
        :'pickup_address' => :'pickupAddress',
        :'destination_address' => :'destinationAddress',
        :'destination_type' => :'destinationType',
        :'secondary_pickup_address' => :'secondaryPickupAddress',
        :'secondary_delivery_address' => :'secondaryDeliveryAddress',
        :'storage_facility' => :'storageFacility',
        :'shipment_type' => :'shipmentType',
        :'diversion' => :'diversion',
        :'point_of_contact' => :'pointOfContact',
        :'counselor_remarks' => :'counselorRemarks',
        :'actual_pro_gear_weight' => :'actualProGearWeight',
        :'actual_spouse_pro_gear_weight' => :'actualSpouseProGearWeight',
        :'ppm_shipment' => :'ppmShipment'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'scheduled_pickup_date' => :'Date',
        :'actual_pickup_date' => :'Date',
        :'first_available_delivery_date' => :'Date',
        :'scheduled_delivery_date' => :'Date',
        :'actual_delivery_date' => :'Date',
        :'prime_estimated_weight' => :'Integer',
        :'prime_actual_weight' => :'Integer',
        :'nts_recorded_weight' => :'Integer',
        :'pickup_address' => :'AddressV2',
        :'destination_address' => :'AddressV2',
        :'destination_type' => :'DestinationTypeV2',
        :'secondary_pickup_address' => :'AddressV2',
        :'secondary_delivery_address' => :'AddressV2',
        :'storage_facility' => :'UpdateMTOShipmentStorageFacilityV2',
        :'shipment_type' => :'MTOShipmentTypeV2',
        :'diversion' => :'Boolean',
        :'point_of_contact' => :'String',
        :'counselor_remarks' => :'String',
        :'actual_pro_gear_weight' => :'Integer',
        :'actual_spouse_pro_gear_weight' => :'Integer',
        :'ppm_shipment' => :'UpdatePPMShipmentV2'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'scheduled_pickup_date',
        :'actual_pickup_date',
        :'first_available_delivery_date',
        :'scheduled_delivery_date',
        :'actual_delivery_date',
        :'prime_estimated_weight',
        :'prime_actual_weight',
        :'nts_recorded_weight',
        :'destination_type',
        :'counselor_remarks',
        :'actual_pro_gear_weight',
        :'actual_spouse_pro_gear_weight',
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `OpenapiClient::UpdateMTOShipmentV2` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `OpenapiClient::UpdateMTOShipmentV2`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'scheduled_pickup_date')
        self.scheduled_pickup_date = attributes[:'scheduled_pickup_date']
      end

      if attributes.key?(:'actual_pickup_date')
        self.actual_pickup_date = attributes[:'actual_pickup_date']
      end

      if attributes.key?(:'first_available_delivery_date')
        self.first_available_delivery_date = attributes[:'first_available_delivery_date']
      end

      if attributes.key?(:'scheduled_delivery_date')
        self.scheduled_delivery_date = attributes[:'scheduled_delivery_date']
      end

      if attributes.key?(:'actual_delivery_date')
        self.actual_delivery_date = attributes[:'actual_delivery_date']
      end

      if attributes.key?(:'prime_estimated_weight')
        self.prime_estimated_weight = attributes[:'prime_estimated_weight']
      end

      if attributes.key?(:'prime_actual_weight')
        self.prime_actual_weight = attributes[:'prime_actual_weight']
      end

      if attributes.key?(:'nts_recorded_weight')
        self.nts_recorded_weight = attributes[:'nts_recorded_weight']
      end

      if attributes.key?(:'pickup_address')
        self.pickup_address = attributes[:'pickup_address']
      end

      if attributes.key?(:'destination_address')
        self.destination_address = attributes[:'destination_address']
      end

      if attributes.key?(:'destination_type')
        self.destination_type = attributes[:'destination_type']
      end

      if attributes.key?(:'secondary_pickup_address')
        self.secondary_pickup_address = attributes[:'secondary_pickup_address']
      end

      if attributes.key?(:'secondary_delivery_address')
        self.secondary_delivery_address = attributes[:'secondary_delivery_address']
      end

      if attributes.key?(:'storage_facility')
        self.storage_facility = attributes[:'storage_facility']
      end

      if attributes.key?(:'shipment_type')
        self.shipment_type = attributes[:'shipment_type']
      end

      if attributes.key?(:'diversion')
        self.diversion = attributes[:'diversion']
      end

      if attributes.key?(:'point_of_contact')
        self.point_of_contact = attributes[:'point_of_contact']
      end

      if attributes.key?(:'counselor_remarks')
        self.counselor_remarks = attributes[:'counselor_remarks']
      end

      if attributes.key?(:'actual_pro_gear_weight')
        self.actual_pro_gear_weight = attributes[:'actual_pro_gear_weight']
      end

      if attributes.key?(:'actual_spouse_pro_gear_weight')
        self.actual_spouse_pro_gear_weight = attributes[:'actual_spouse_pro_gear_weight']
      end

      if attributes.key?(:'ppm_shipment')
        self.ppm_shipment = attributes[:'ppm_shipment']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      warn '[DEPRECATED] the `list_invalid_properties` method is obsolete'
      invalid_properties = Array.new
      if !@prime_estimated_weight.nil? && @prime_estimated_weight < 1
        invalid_properties.push('invalid value for "prime_estimated_weight", must be greater than or equal to 1.')
      end

      if !@prime_actual_weight.nil? && @prime_actual_weight < 1
        invalid_properties.push('invalid value for "prime_actual_weight", must be greater than or equal to 1.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      warn '[DEPRECATED] the `valid?` method is obsolete'
      return false if !@prime_estimated_weight.nil? && @prime_estimated_weight < 1
      return false if !@prime_actual_weight.nil? && @prime_actual_weight < 1
      true
    end

    # Custom attribute writer method with validation
    # @param [Object] prime_estimated_weight Value to be assigned
    def prime_estimated_weight=(prime_estimated_weight)
      if !prime_estimated_weight.nil? && prime_estimated_weight < 1
        fail ArgumentError, 'invalid value for "prime_estimated_weight", must be greater than or equal to 1.'
      end

      @prime_estimated_weight = prime_estimated_weight
    end

    # Custom attribute writer method with validation
    # @param [Object] prime_actual_weight Value to be assigned
    def prime_actual_weight=(prime_actual_weight)
      if !prime_actual_weight.nil? && prime_actual_weight < 1
        fail ArgumentError, 'invalid value for "prime_actual_weight", must be greater than or equal to 1.'
      end

      @prime_actual_weight = prime_actual_weight
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          scheduled_pickup_date == o.scheduled_pickup_date &&
          actual_pickup_date == o.actual_pickup_date &&
          first_available_delivery_date == o.first_available_delivery_date &&
          scheduled_delivery_date == o.scheduled_delivery_date &&
          actual_delivery_date == o.actual_delivery_date &&
          prime_estimated_weight == o.prime_estimated_weight &&
          prime_actual_weight == o.prime_actual_weight &&
          nts_recorded_weight == o.nts_recorded_weight &&
          pickup_address == o.pickup_address &&
          destination_address == o.destination_address &&
          destination_type == o.destination_type &&
          secondary_pickup_address == o.secondary_pickup_address &&
          secondary_delivery_address == o.secondary_delivery_address &&
          storage_facility == o.storage_facility &&
          shipment_type == o.shipment_type &&
          diversion == o.diversion &&
          point_of_contact == o.point_of_contact &&
          counselor_remarks == o.counselor_remarks &&
          actual_pro_gear_weight == o.actual_pro_gear_weight &&
          actual_spouse_pro_gear_weight == o.actual_spouse_pro_gear_weight &&
          ppm_shipment == o.ppm_shipment
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [scheduled_pickup_date, actual_pickup_date, first_available_delivery_date, scheduled_delivery_date, actual_delivery_date, prime_estimated_weight, prime_actual_weight, nts_recorded_weight, pickup_address, destination_address, destination_type, secondary_pickup_address, secondary_delivery_address, storage_facility, shipment_type, diversion, point_of_contact, counselor_remarks, actual_pro_gear_weight, actual_spouse_pro_gear_weight, ppm_shipment].hash
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
