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
  # This represents a destination address change request made by the Prime that is either auto-approved or requires review if the pricing criteria has changed. If criteria has changed, then it must be approved or rejected by a TOO. 
  class ShipmentAddressUpdateV2
    attr_accessor :id

    # The reason there is an address change.
    attr_accessor :contractor_remarks

    # The TOO comment on approval or rejection.
    attr_accessor :office_remarks

    attr_accessor :status

    attr_accessor :shipment_id

    attr_accessor :original_address

    attr_accessor :new_address

    attr_accessor :sit_original_address

    # The distance between the original SIT address and the previous/old destination address of shipment
    attr_accessor :old_sit_distance_between

    # The distance between the original SIT address and requested new destination address of shipment
    attr_accessor :new_sit_distance_between

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
        :'contractor_remarks' => :'contractorRemarks',
        :'office_remarks' => :'officeRemarks',
        :'status' => :'status',
        :'shipment_id' => :'shipmentID',
        :'original_address' => :'originalAddress',
        :'new_address' => :'newAddress',
        :'sit_original_address' => :'sitOriginalAddress',
        :'old_sit_distance_between' => :'oldSitDistanceBetween',
        :'new_sit_distance_between' => :'newSitDistanceBetween'
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
        :'contractor_remarks' => :'String',
        :'office_remarks' => :'String',
        :'status' => :'ShipmentAddressUpdateStatusV2',
        :'shipment_id' => :'String',
        :'original_address' => :'AddressV2',
        :'new_address' => :'AddressV2',
        :'sit_original_address' => :'AddressV2',
        :'old_sit_distance_between' => :'Integer',
        :'new_sit_distance_between' => :'Integer'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'office_remarks',
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `OpenapiClient::ShipmentAddressUpdateV2` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `OpenapiClient::ShipmentAddressUpdateV2`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'id')
        self.id = attributes[:'id']
      else
        self.id = nil
      end

      if attributes.key?(:'contractor_remarks')
        self.contractor_remarks = attributes[:'contractor_remarks']
      else
        self.contractor_remarks = nil
      end

      if attributes.key?(:'office_remarks')
        self.office_remarks = attributes[:'office_remarks']
      end

      if attributes.key?(:'status')
        self.status = attributes[:'status']
      else
        self.status = nil
      end

      if attributes.key?(:'shipment_id')
        self.shipment_id = attributes[:'shipment_id']
      else
        self.shipment_id = nil
      end

      if attributes.key?(:'original_address')
        self.original_address = attributes[:'original_address']
      else
        self.original_address = nil
      end

      if attributes.key?(:'new_address')
        self.new_address = attributes[:'new_address']
      else
        self.new_address = nil
      end

      if attributes.key?(:'sit_original_address')
        self.sit_original_address = attributes[:'sit_original_address']
      end

      if attributes.key?(:'old_sit_distance_between')
        self.old_sit_distance_between = attributes[:'old_sit_distance_between']
      end

      if attributes.key?(:'new_sit_distance_between')
        self.new_sit_distance_between = attributes[:'new_sit_distance_between']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      warn '[DEPRECATED] the `list_invalid_properties` method is obsolete'
      invalid_properties = Array.new
      if @id.nil?
        invalid_properties.push('invalid value for "id", id cannot be nil.')
      end

      if @contractor_remarks.nil?
        invalid_properties.push('invalid value for "contractor_remarks", contractor_remarks cannot be nil.')
      end

      if @status.nil?
        invalid_properties.push('invalid value for "status", status cannot be nil.')
      end

      if @shipment_id.nil?
        invalid_properties.push('invalid value for "shipment_id", shipment_id cannot be nil.')
      end

      if @original_address.nil?
        invalid_properties.push('invalid value for "original_address", original_address cannot be nil.')
      end

      if @new_address.nil?
        invalid_properties.push('invalid value for "new_address", new_address cannot be nil.')
      end

      if !@old_sit_distance_between.nil? && @old_sit_distance_between < 0
        invalid_properties.push('invalid value for "old_sit_distance_between", must be greater than or equal to 0.')
      end

      if !@new_sit_distance_between.nil? && @new_sit_distance_between < 0
        invalid_properties.push('invalid value for "new_sit_distance_between", must be greater than or equal to 0.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      warn '[DEPRECATED] the `valid?` method is obsolete'
      return false if @id.nil?
      return false if @contractor_remarks.nil?
      return false if @status.nil?
      return false if @shipment_id.nil?
      return false if @original_address.nil?
      return false if @new_address.nil?
      return false if !@old_sit_distance_between.nil? && @old_sit_distance_between < 0
      return false if !@new_sit_distance_between.nil? && @new_sit_distance_between < 0
      true
    end

    # Custom attribute writer method with validation
    # @param [Object] old_sit_distance_between Value to be assigned
    def old_sit_distance_between=(old_sit_distance_between)
      if old_sit_distance_between.nil?
        fail ArgumentError, 'old_sit_distance_between cannot be nil'
      end

      if old_sit_distance_between < 0
        fail ArgumentError, 'invalid value for "old_sit_distance_between", must be greater than or equal to 0.'
      end

      @old_sit_distance_between = old_sit_distance_between
    end

    # Custom attribute writer method with validation
    # @param [Object] new_sit_distance_between Value to be assigned
    def new_sit_distance_between=(new_sit_distance_between)
      if new_sit_distance_between.nil?
        fail ArgumentError, 'new_sit_distance_between cannot be nil'
      end

      if new_sit_distance_between < 0
        fail ArgumentError, 'invalid value for "new_sit_distance_between", must be greater than or equal to 0.'
      end

      @new_sit_distance_between = new_sit_distance_between
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          id == o.id &&
          contractor_remarks == o.contractor_remarks &&
          office_remarks == o.office_remarks &&
          status == o.status &&
          shipment_id == o.shipment_id &&
          original_address == o.original_address &&
          new_address == o.new_address &&
          sit_original_address == o.sit_original_address &&
          old_sit_distance_between == o.old_sit_distance_between &&
          new_sit_distance_between == o.new_sit_distance_between
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [id, contractor_remarks, office_remarks, status, shipment_id, original_address, new_address, sit_original_address, old_sit_distance_between, new_sit_distance_between].hash
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
