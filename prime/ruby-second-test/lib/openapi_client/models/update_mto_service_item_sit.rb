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
  # Subtype used to provide the departure date for origin or destination SIT. This is not creating a new service item but rather updating and existing service item. 
  class UpdateMTOServiceItemSIT < UpdateMTOServiceItem
    # Service code allowed for this model type.
    attr_accessor :re_service_code

    # Departure date for SIT. This is the end date of the SIT at either origin or destination.
    attr_accessor :sit_departure_date

    attr_accessor :sit_destination_final_address

    # Date of attempted contact by the prime corresponding to 'timeMilitary1'.
    attr_accessor :date_of_contact1

    # Time of attempted contact by the prime corresponding to 'dateOfContact1', in military format.
    attr_accessor :time_military1

    # First available date that Prime can deliver SIT service item.
    attr_accessor :first_available_delivery_date1

    # Date of attempted contact by the prime corresponding to 'timeMilitary2'.
    attr_accessor :date_of_contact2

    # Time of attempted contact by the prime corresponding to 'dateOfContact2', in military format.
    attr_accessor :time_military2

    # Second available date that Prime can deliver SIT service item.
    attr_accessor :first_available_delivery_date2

    # Date when the customer has requested delivery out of SIT.
    attr_accessor :sit_requested_delivery

    # Date when the customer contacted the prime for a delivery out of SIT.
    attr_accessor :sit_customer_contacted

    # Reason for updating service item.
    attr_accessor :update_reason

    attr_accessor :sit_postal_code

    # Entry date for the SIT.
    attr_accessor :sit_entry_date

    # Indicates if \"Approvals Requested\" status is being requested.
    attr_accessor :request_approvals_requested_status

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
        :'sit_departure_date' => :'sitDepartureDate',
        :'sit_destination_final_address' => :'sitDestinationFinalAddress',
        :'date_of_contact1' => :'dateOfContact1',
        :'time_military1' => :'timeMilitary1',
        :'first_available_delivery_date1' => :'firstAvailableDeliveryDate1',
        :'date_of_contact2' => :'dateOfContact2',
        :'time_military2' => :'timeMilitary2',
        :'first_available_delivery_date2' => :'firstAvailableDeliveryDate2',
        :'sit_requested_delivery' => :'sitRequestedDelivery',
        :'sit_customer_contacted' => :'sitCustomerContacted',
        :'update_reason' => :'updateReason',
        :'sit_postal_code' => :'sitPostalCode',
        :'sit_entry_date' => :'sitEntryDate',
        :'request_approvals_requested_status' => :'requestApprovalsRequestedStatus'
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
        :'sit_departure_date' => :'Date',
        :'sit_destination_final_address' => :'Address',
        :'date_of_contact1' => :'Date',
        :'time_military1' => :'String',
        :'first_available_delivery_date1' => :'Date',
        :'date_of_contact2' => :'Date',
        :'time_military2' => :'String',
        :'first_available_delivery_date2' => :'Date',
        :'sit_requested_delivery' => :'Date',
        :'sit_customer_contacted' => :'Date',
        :'update_reason' => :'String',
        :'sit_postal_code' => :'String',
        :'sit_entry_date' => :'Date',
        :'request_approvals_requested_status' => :'Boolean'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'date_of_contact1',
        :'time_military1',
        :'first_available_delivery_date1',
        :'date_of_contact2',
        :'time_military2',
        :'first_available_delivery_date2',
        :'sit_requested_delivery',
        :'sit_customer_contacted',
        :'update_reason',
        :'sit_postal_code',
        :'sit_entry_date',
        :'request_approvals_requested_status'
      ])
    end

    # List of class defined in allOf (OpenAPI v3)
    def self.openapi_all_of
      [
      :'UpdateMTOServiceItem'
      ]
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `OpenapiClient::UpdateMTOServiceItemSIT` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `OpenapiClient::UpdateMTOServiceItemSIT`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      # call parent's initialize
      super(attributes)

      if attributes.key?(:'re_service_code')
        self.re_service_code = attributes[:'re_service_code']
      end

      if attributes.key?(:'sit_departure_date')
        self.sit_departure_date = attributes[:'sit_departure_date']
      end

      if attributes.key?(:'sit_destination_final_address')
        self.sit_destination_final_address = attributes[:'sit_destination_final_address']
      end

      if attributes.key?(:'date_of_contact1')
        self.date_of_contact1 = attributes[:'date_of_contact1']
      end

      if attributes.key?(:'time_military1')
        self.time_military1 = attributes[:'time_military1']
      end

      if attributes.key?(:'first_available_delivery_date1')
        self.first_available_delivery_date1 = attributes[:'first_available_delivery_date1']
      end

      if attributes.key?(:'date_of_contact2')
        self.date_of_contact2 = attributes[:'date_of_contact2']
      end

      if attributes.key?(:'time_military2')
        self.time_military2 = attributes[:'time_military2']
      end

      if attributes.key?(:'first_available_delivery_date2')
        self.first_available_delivery_date2 = attributes[:'first_available_delivery_date2']
      end

      if attributes.key?(:'sit_requested_delivery')
        self.sit_requested_delivery = attributes[:'sit_requested_delivery']
      end

      if attributes.key?(:'sit_customer_contacted')
        self.sit_customer_contacted = attributes[:'sit_customer_contacted']
      end

      if attributes.key?(:'update_reason')
        self.update_reason = attributes[:'update_reason']
      end

      if attributes.key?(:'sit_postal_code')
        self.sit_postal_code = attributes[:'sit_postal_code']
      end

      if attributes.key?(:'sit_entry_date')
        self.sit_entry_date = attributes[:'sit_entry_date']
      end

      if attributes.key?(:'request_approvals_requested_status')
        self.request_approvals_requested_status = attributes[:'request_approvals_requested_status']
      end
    end

    # Show invalid properties with the reasons. Usually used together with valid?
    # @return Array for valid properties with the reasons
    def list_invalid_properties
      warn '[DEPRECATED] the `list_invalid_properties` method is obsolete'
      invalid_properties = super
      pattern = Regexp.new(/\d{4}Z/)
      if !@time_military1.nil? && @time_military1 !~ pattern
        invalid_properties.push("invalid value for \"time_military1\", must conform to the pattern #{pattern}.")
      end

      pattern = Regexp.new(/\d{4}Z/)
      if !@time_military2.nil? && @time_military2 !~ pattern
        invalid_properties.push("invalid value for \"time_military2\", must conform to the pattern #{pattern}.")
      end

      pattern = Regexp.new(/^(\d{5}([\-]\d{4})?)$/)
      if !@sit_postal_code.nil? && @sit_postal_code !~ pattern
        invalid_properties.push("invalid value for \"sit_postal_code\", must conform to the pattern #{pattern}.")
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      warn '[DEPRECATED] the `valid?` method is obsolete'
      re_service_code_validator = EnumAttributeValidator.new('String', ["DDDSIT", "DDASIT", "DOPSIT", "DOASIT", "DOFSIT"])
      return false unless re_service_code_validator.valid?(@re_service_code)
      return false if !@time_military1.nil? && @time_military1 !~ Regexp.new(/\d{4}Z/)
      return false if !@time_military2.nil? && @time_military2 !~ Regexp.new(/\d{4}Z/)
      return false if !@sit_postal_code.nil? && @sit_postal_code !~ Regexp.new(/^(\d{5}([\-]\d{4})?)$/)
      true && super
    end

    # Custom attribute writer method checking allowed values (enum).
    # @param [Object] re_service_code Object to be assigned
    def re_service_code=(re_service_code)
      validator = EnumAttributeValidator.new('String', ["DDDSIT", "DDASIT", "DOPSIT", "DOASIT", "DOFSIT"])
      unless validator.valid?(re_service_code)
        fail ArgumentError, "invalid value for \"re_service_code\", must be one of #{validator.allowable_values}."
      end
      @re_service_code = re_service_code
    end

    # Custom attribute writer method with validation
    # @param [Object] time_military1 Value to be assigned
    def time_military1=(time_military1)
      pattern = Regexp.new(/\d{4}Z/)
      if !time_military1.nil? && time_military1 !~ pattern
        fail ArgumentError, "invalid value for \"time_military1\", must conform to the pattern #{pattern}."
      end

      @time_military1 = time_military1
    end

    # Custom attribute writer method with validation
    # @param [Object] time_military2 Value to be assigned
    def time_military2=(time_military2)
      pattern = Regexp.new(/\d{4}Z/)
      if !time_military2.nil? && time_military2 !~ pattern
        fail ArgumentError, "invalid value for \"time_military2\", must conform to the pattern #{pattern}."
      end

      @time_military2 = time_military2
    end

    # Custom attribute writer method with validation
    # @param [Object] sit_postal_code Value to be assigned
    def sit_postal_code=(sit_postal_code)
      pattern = Regexp.new(/^(\d{5}([\-]\d{4})?)$/)
      if !sit_postal_code.nil? && sit_postal_code !~ pattern
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
          sit_departure_date == o.sit_departure_date &&
          sit_destination_final_address == o.sit_destination_final_address &&
          date_of_contact1 == o.date_of_contact1 &&
          time_military1 == o.time_military1 &&
          first_available_delivery_date1 == o.first_available_delivery_date1 &&
          date_of_contact2 == o.date_of_contact2 &&
          time_military2 == o.time_military2 &&
          first_available_delivery_date2 == o.first_available_delivery_date2 &&
          sit_requested_delivery == o.sit_requested_delivery &&
          sit_customer_contacted == o.sit_customer_contacted &&
          update_reason == o.update_reason &&
          sit_postal_code == o.sit_postal_code &&
          sit_entry_date == o.sit_entry_date &&
          request_approvals_requested_status == o.request_approvals_requested_status && super(o)
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [re_service_code, sit_departure_date, sit_destination_final_address, date_of_contact1, time_military1, first_available_delivery_date1, date_of_contact2, time_military2, first_available_delivery_date2, sit_requested_delivery, sit_customer_contacted, update_reason, sit_postal_code, sit_entry_date, request_approvals_requested_status].hash
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
