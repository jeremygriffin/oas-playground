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
  # Describes a domestic origin SIT service item. Subtype of a MTOServiceItem.
  class MTOServiceItemOriginSITV2 < MTOServiceItemV2
    # Service code allowed for this model type.
    attr_accessor :re_service_code

    # Explanation of why Prime is picking up SIT item.
    attr_accessor :reason

    attr_accessor :sit_postal_code

    # Entry date for the SIT
    attr_accessor :sit_entry_date

    # Departure date for SIT. This is the end date of the SIT at either origin or destination. This is optional as it can be updated using the UpdateMTOServiceItemSIT modelType at a later date.
    attr_accessor :sit_departure_date

    attr_accessor :sit_hhg_actual_origin

    attr_accessor :sit_hhg_original_origin

    attr_accessor :request_approvals_requested_status

    # Date when the customer has requested delivery out of SIT.
    attr_accessor :sit_requested_delivery

    # Date when the customer contacted the prime for a delivery out of SIT.
    attr_accessor :sit_customer_contacted

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
        :'re_service_code' => :'reServiceCode',
        :'reason' => :'reason',
        :'sit_postal_code' => :'sitPostalCode',
        :'sit_entry_date' => :'sitEntryDate',
        :'sit_departure_date' => :'sitDepartureDate',
        :'sit_hhg_actual_origin' => :'sitHHGActualOrigin',
        :'sit_hhg_original_origin' => :'sitHHGOriginalOrigin',
        :'request_approvals_requested_status' => :'requestApprovalsRequestedStatus',
        :'sit_requested_delivery' => :'sitRequestedDelivery',
        :'sit_customer_contacted' => :'sitCustomerContacted'
      }
    end

    # Returns all the JSON keys this model knows about, including the ones defined in its parent(s)
    def self.acceptable_attributes
      attribute_map.values.concat(superclass.acceptable_attributes)
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'re_service_code' => :'String',
        :'reason' => :'String',
        :'sit_postal_code' => :'String',
        :'sit_entry_date' => :'Date',
        :'sit_departure_date' => :'Date',
        :'sit_hhg_actual_origin' => :'AddressV2',
        :'sit_hhg_original_origin' => :'AddressV2',
        :'request_approvals_requested_status' => :'Boolean',
        :'sit_requested_delivery' => :'Date',
        :'sit_customer_contacted' => :'Date'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'sit_departure_date',
        :'sit_requested_delivery',
        :'sit_customer_contacted'
      ])
    end

    # List of class defined in allOf (OpenAPI v3)
    def self.openapi_all_of
      [
      :'MTOServiceItemV2'
      ]
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `OpenapiClient::MTOServiceItemOriginSITV2` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `OpenapiClient::MTOServiceItemOriginSITV2`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      # call parent's initialize
      super(attributes)

      if attributes.key?(:'re_service_code')
        self.re_service_code = attributes[:'re_service_code']
      else
        self.re_service_code = nil
      end

      if attributes.key?(:'reason')
        self.reason = attributes[:'reason']
      else
        self.reason = nil
      end

      if attributes.key?(:'sit_postal_code')
        self.sit_postal_code = attributes[:'sit_postal_code']
      else
        self.sit_postal_code = nil
      end

      if attributes.key?(:'sit_entry_date')
        self.sit_entry_date = attributes[:'sit_entry_date']
      else
        self.sit_entry_date = nil
      end

      if attributes.key?(:'sit_departure_date')
        self.sit_departure_date = attributes[:'sit_departure_date']
      end

      if attributes.key?(:'sit_hhg_actual_origin')
        self.sit_hhg_actual_origin = attributes[:'sit_hhg_actual_origin']
      end

      if attributes.key?(:'sit_hhg_original_origin')
        self.sit_hhg_original_origin = attributes[:'sit_hhg_original_origin']
      end

      if attributes.key?(:'request_approvals_requested_status')
        self.request_approvals_requested_status = attributes[:'request_approvals_requested_status']
      end

      if attributes.key?(:'sit_requested_delivery')
        self.sit_requested_delivery = attributes[:'sit_requested_delivery']
      end

      if attributes.key?(:'sit_customer_contacted')
        self.sit_customer_contacted = attributes[:'sit_customer_contacted']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      warn '[DEPRECATED] the `list_invalid_properties` method is obsolete'
      invalid_properties = super
      if @re_service_code.nil?
        invalid_properties.push('invalid value for "re_service_code", re_service_code cannot be nil.')
      end

      if @reason.nil?
        invalid_properties.push('invalid value for "reason", reason cannot be nil.')
      end

      if @sit_postal_code.nil?
        invalid_properties.push('invalid value for "sit_postal_code", sit_postal_code cannot be nil.')
      end

      pattern = Regexp.new(/^(\d{5}([\-]\d{4})?)$/)
      if @sit_postal_code !~ pattern
        invalid_properties.push("invalid value for \"sit_postal_code\", must conform to the pattern #{pattern}.")
      end

      if @sit_entry_date.nil?
        invalid_properties.push('invalid value for "sit_entry_date", sit_entry_date cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      warn '[DEPRECATED] the `valid?` method is obsolete'
      return false if @re_service_code.nil?
      re_service_code_validator = EnumAttributeValidator.new('String', ["DOFSIT", "DOASIT"])
      return false unless re_service_code_validator.valid?(@re_service_code)
      return false if @reason.nil?
      return false if @sit_postal_code.nil?
      return false if @sit_postal_code !~ Regexp.new(/^(\d{5}([\-]\d{4})?)$/)
      return false if @sit_entry_date.nil?
      true && super
    end

    # Custom attribute writer method checking allowed values (enum).
    # @param [Object] re_service_code Object to be assigned
    def re_service_code=(re_service_code)
      validator = EnumAttributeValidator.new('String', ["DOFSIT", "DOASIT"])
      unless validator.valid?(re_service_code)
        fail ArgumentError, "invalid value for \"re_service_code\", must be one of #{validator.allowable_values}."
      end
      @re_service_code = re_service_code
    end

    # Custom attribute writer method with validation
    # @param [Object] sit_postal_code Value to be assigned
    def sit_postal_code=(sit_postal_code)
      if sit_postal_code.nil?
        fail ArgumentError, 'sit_postal_code cannot be nil'
      end

      pattern = Regexp.new(/^(\d{5}([\-]\d{4})?)$/)
      if sit_postal_code !~ pattern
        fail ArgumentError, "invalid value for \"sit_postal_code\", must conform to the pattern #{pattern}."
      end

      @sit_postal_code = sit_postal_code
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          re_service_code == o.re_service_code &&
          reason == o.reason &&
          sit_postal_code == o.sit_postal_code &&
          sit_entry_date == o.sit_entry_date &&
          sit_departure_date == o.sit_departure_date &&
          sit_hhg_actual_origin == o.sit_hhg_actual_origin &&
          sit_hhg_original_origin == o.sit_hhg_original_origin &&
          request_approvals_requested_status == o.request_approvals_requested_status &&
          sit_requested_delivery == o.sit_requested_delivery &&
          sit_customer_contacted == o.sit_customer_contacted && super(o)
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [re_service_code, reason, sit_postal_code, sit_entry_date, sit_departure_date, sit_hhg_actual_origin, sit_hhg_original_origin, request_approvals_requested_status, sit_requested_delivery, sit_customer_contacted].hash
    end

    # Builds the object from hash
    # @param [Hash] attributes Model attributes in the form of hash
    # @return [Object] Returns the model itself
    def self.build_from_hash(attributes)
      return nil unless attributes.is_a?(Hash)
      super(attributes)
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
      hash = super
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
