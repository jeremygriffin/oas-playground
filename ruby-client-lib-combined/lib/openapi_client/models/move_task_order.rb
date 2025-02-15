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
  class MoveTaskOrder
    attr_accessor :id

    attr_accessor :move_code

    attr_accessor :created_at

    attr_accessor :order_id

    attr_accessor :order

    attr_accessor :reference_id

    attr_accessor :available_to_prime_at

    attr_accessor :updated_at

    attr_accessor :prime_counseling_completed_at

    attr_accessor :payment_requests

    attr_accessor :mto_service_items

    # A list of shipments without their associated service items.
    attr_accessor :mto_shipments

    attr_accessor :ppm_type

    attr_accessor :ppm_estimated_weight

    attr_accessor :excess_weight_qualified_at

    attr_accessor :excess_weight_acknowledged_at

    attr_accessor :excess_weight_upload_id

    attr_accessor :e_tag

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
        :'id' => :'id',
        :'move_code' => :'moveCode',
        :'created_at' => :'createdAt',
        :'order_id' => :'orderID',
        :'order' => :'order',
        :'reference_id' => :'referenceId',
        :'available_to_prime_at' => :'availableToPrimeAt',
        :'updated_at' => :'updatedAt',
        :'prime_counseling_completed_at' => :'primeCounselingCompletedAt',
        :'payment_requests' => :'paymentRequests',
        :'mto_service_items' => :'mtoServiceItems',
        :'mto_shipments' => :'mtoShipments',
        :'ppm_type' => :'ppmType',
        :'ppm_estimated_weight' => :'ppmEstimatedWeight',
        :'excess_weight_qualified_at' => :'excessWeightQualifiedAt',
        :'excess_weight_acknowledged_at' => :'excessWeightAcknowledgedAt',
        :'excess_weight_upload_id' => :'excessWeightUploadId',
        :'e_tag' => :'eTag'
      }
    end

    # Returns all the JSON keys this model knows about
    def self.acceptable_attributes
      attribute_map.values
    end

    # Attribute type mapping.
    def self.openapi_types
      {
        :'id' => :'String',
        :'move_code' => :'String',
        :'created_at' => :'Time',
        :'order_id' => :'String',
        :'order' => :'Order',
        :'reference_id' => :'String',
        :'available_to_prime_at' => :'Time',
        :'updated_at' => :'Time',
        :'prime_counseling_completed_at' => :'Time',
        :'payment_requests' => :'Array<PaymentRequest>',
        :'mto_service_items' => :'Array<MTOServiceItem>',
        :'mto_shipments' => :'Array<MTOShipmentWithoutServiceItems>',
        :'ppm_type' => :'String',
        :'ppm_estimated_weight' => :'Integer',
        :'excess_weight_qualified_at' => :'Time',
        :'excess_weight_acknowledged_at' => :'Time',
        :'excess_weight_upload_id' => :'String',
        :'e_tag' => :'String'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'available_to_prime_at',
        :'prime_counseling_completed_at',
        :'excess_weight_qualified_at',
        :'excess_weight_acknowledged_at',
        :'excess_weight_upload_id',
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `OpenapiClient::MoveTaskOrder` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `OpenapiClient::MoveTaskOrder`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'id')
        self.id = attributes[:'id']
      end

      if attributes.key?(:'move_code')
        self.move_code = attributes[:'move_code']
      end

      if attributes.key?(:'created_at')
        self.created_at = attributes[:'created_at']
      end

      if attributes.key?(:'order_id')
        self.order_id = attributes[:'order_id']
      end

      if attributes.key?(:'order')
        self.order = attributes[:'order']
      end

      if attributes.key?(:'reference_id')
        self.reference_id = attributes[:'reference_id']
      end

      if attributes.key?(:'available_to_prime_at')
        self.available_to_prime_at = attributes[:'available_to_prime_at']
      end

      if attributes.key?(:'updated_at')
        self.updated_at = attributes[:'updated_at']
      end

      if attributes.key?(:'prime_counseling_completed_at')
        self.prime_counseling_completed_at = attributes[:'prime_counseling_completed_at']
      end

      if attributes.key?(:'payment_requests')
        if (value = attributes[:'payment_requests']).is_a?(Array)
          self.payment_requests = value
        end
      else
        self.payment_requests = nil
      end

      if attributes.key?(:'mto_service_items')
        if (value = attributes[:'mto_service_items']).is_a?(Array)
          self.mto_service_items = value
        end
      else
        self.mto_service_items = nil
      end

      if attributes.key?(:'mto_shipments')
        if (value = attributes[:'mto_shipments']).is_a?(Array)
          self.mto_shipments = value
        end
      else
        self.mto_shipments = nil
      end

      if attributes.key?(:'ppm_type')
        self.ppm_type = attributes[:'ppm_type']
      end

      if attributes.key?(:'ppm_estimated_weight')
        self.ppm_estimated_weight = attributes[:'ppm_estimated_weight']
      end

      if attributes.key?(:'excess_weight_qualified_at')
        self.excess_weight_qualified_at = attributes[:'excess_weight_qualified_at']
      end

      if attributes.key?(:'excess_weight_acknowledged_at')
        self.excess_weight_acknowledged_at = attributes[:'excess_weight_acknowledged_at']
      end

      if attributes.key?(:'excess_weight_upload_id')
        self.excess_weight_upload_id = attributes[:'excess_weight_upload_id']
      end

      if attributes.key?(:'e_tag')
        self.e_tag = attributes[:'e_tag']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      warn '[DEPRECATED] the `list_invalid_properties` method is obsolete'
      invalid_properties = Array.new
      if @payment_requests.nil?
        invalid_properties.push('invalid value for "payment_requests", payment_requests cannot be nil.')
      end

      if @mto_service_items.nil?
        invalid_properties.push('invalid value for "mto_service_items", mto_service_items cannot be nil.')
      end

      if @mto_shipments.nil?
        invalid_properties.push('invalid value for "mto_shipments", mto_shipments cannot be nil.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      warn '[DEPRECATED] the `valid?` method is obsolete'
      return false if @payment_requests.nil?
      return false if @mto_service_items.nil?
      return false if @mto_shipments.nil?
      ppm_type_validator = EnumAttributeValidator.new('String', ["PARTIAL", "FULL"])
      return false unless ppm_type_validator.valid?(@ppm_type)
      true
    end

    # Custom attribute writer method checking allowed values (enum).
    # @param [Object] ppm_type Object to be assigned
    def ppm_type=(ppm_type)
      validator = EnumAttributeValidator.new('String', ["PARTIAL", "FULL"])
      unless validator.valid?(ppm_type)
        fail ArgumentError, "invalid value for \"ppm_type\", must be one of #{validator.allowable_values}."
      end
      @ppm_type = ppm_type
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          id == o.id &&
          move_code == o.move_code &&
          created_at == o.created_at &&
          order_id == o.order_id &&
          order == o.order &&
          reference_id == o.reference_id &&
          available_to_prime_at == o.available_to_prime_at &&
          updated_at == o.updated_at &&
          prime_counseling_completed_at == o.prime_counseling_completed_at &&
          payment_requests == o.payment_requests &&
          mto_service_items == o.mto_service_items &&
          mto_shipments == o.mto_shipments &&
          ppm_type == o.ppm_type &&
          ppm_estimated_weight == o.ppm_estimated_weight &&
          excess_weight_qualified_at == o.excess_weight_qualified_at &&
          excess_weight_acknowledged_at == o.excess_weight_acknowledged_at &&
          excess_weight_upload_id == o.excess_weight_upload_id &&
          e_tag == o.e_tag
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [id, move_code, created_at, order_id, order, reference_id, available_to_prime_at, updated_at, prime_counseling_completed_at, payment_requests, mto_service_items, mto_shipments, ppm_type, ppm_estimated_weight, excess_weight_qualified_at, excess_weight_acknowledged_at, excess_weight_upload_id, e_tag].hash
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
