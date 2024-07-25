=begin
#MilMove Prime V3 API

#The Prime V3 API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v3/`. 

The version of the OpenAPI document: 0.0.1
Contact: milmove-developers@caci.com
Generated by: https://openapi-generator.tech
Generator version: 7.8.0-SNAPSHOT

=end

require 'date'
require 'time'

module OpenapiClient
  class MTOShipmentWithoutServiceItemsV3
    # The ID of the shipment.
    attr_accessor :id

    # The ID of the move for this shipment.
    attr_accessor :move_task_order_id

    # The date when the Task Ordering Officer first approved this shipment for the move.
    attr_accessor :approved_date

    # The date the customer selects during onboarding as their preferred pickup date. Other dates, such as required delivery date and (outside MilMove) the pack date, are derived from this date. 
    attr_accessor :requested_pickup_date

    # The customer's preferred delivery date.
    attr_accessor :requested_delivery_date

    # The date the Prime contractor scheduled to pick up this shipment after consultation with the customer.
    attr_accessor :scheduled_pickup_date

    # The date when the Prime contractor actually picked up the shipment. Updated after-the-fact.
    attr_accessor :actual_pickup_date

    # The date the Prime provides to the customer as the first possible delivery date so that they can plan their travel accordingly. 
    attr_accessor :first_available_delivery_date

    # The latest date by which the Prime can deliver a customer's shipment without violating the contract. This is calculated based on weight, distance, and the scheduled pickup date. It cannot be modified. 
    attr_accessor :required_delivery_date

    # The date the Prime contractor scheduled to deliver this shipment after consultation with the customer.
    attr_accessor :scheduled_delivery_date

    # The date when the Prime contractor actually delivered the shipment. Updated after-the-fact.
    attr_accessor :actual_delivery_date

    # The estimated weight of this shipment, determined by the movers during the pre-move survey. This value **can only be updated once.** If there was an issue with estimating the weight and a mistake was made, the Prime contracter will need to contact the TOO to change it. 
    attr_accessor :prime_estimated_weight

    # The date when the Prime contractor recorded the shipment's estimated weight.
    attr_accessor :prime_estimated_weight_recorded_date

    # The actual weight of the shipment, provided after the Prime packs, picks up, and weighs a customer's shipment.
    attr_accessor :prime_actual_weight

    # The previously recorded weight for the NTS Shipment. Used for NTS Release to know what the previous primeActualWeight or billable weight was.
    attr_accessor :nts_recorded_weight

    # The customer can use the customer remarks field to inform the services counselor and the movers about any special circumstances for this shipment. Typical examples:   * bulky or fragile items,   * weapons,   * access info for their address.  Customer enters this information during onboarding. Optional field. 
    attr_accessor :customer_remarks

    # The counselor can use the counselor remarks field to inform the movers about any special circumstances for this shipment. Typical examples:   * bulky or fragile items,   * weapons,   * access info for their address.  Counselors enters this information when creating or editing an MTO Shipment. Optional field. 
    attr_accessor :counselor_remarks

    # A list of the agents for a shipment. Agents are the people who the Prime contractor recognize as permitted to release (in the case of pickup) or receive (on delivery) a shipment. 
    attr_accessor :agents

    attr_accessor :sit_extensions

    attr_accessor :reweigh

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

    # The reason the TOO provided when requesting a diversion for this shipment. 
    attr_accessor :diversion_reason

    # The status of a shipment, indicating where it is in the TOO's approval process. Can only be updated by the contractor in special circumstances. 
    attr_accessor :status

    attr_accessor :ppm_shipment

    attr_accessor :delivery_address_update

    # A hash unique to this shipment that should be used as the \"If-Match\" header for any updates.
    attr_accessor :e_tag

    attr_accessor :created_at

    attr_accessor :updated_at

    # Email or ID of the person who will be contacted in the event of questions or concerns about this update. May be the person performing the update, or someone else working with the Prime contractor. 
    attr_accessor :point_of_contact

    # The SIT authorized end date for origin SIT.
    attr_accessor :origin_sit_auth_end_date

    # The SIT authorized end date for destination SIT.
    attr_accessor :destination_sit_auth_end_date

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
        :'move_task_order_id' => :'moveTaskOrderID',
        :'approved_date' => :'approvedDate',
        :'requested_pickup_date' => :'requestedPickupDate',
        :'requested_delivery_date' => :'requestedDeliveryDate',
        :'scheduled_pickup_date' => :'scheduledPickupDate',
        :'actual_pickup_date' => :'actualPickupDate',
        :'first_available_delivery_date' => :'firstAvailableDeliveryDate',
        :'required_delivery_date' => :'requiredDeliveryDate',
        :'scheduled_delivery_date' => :'scheduledDeliveryDate',
        :'actual_delivery_date' => :'actualDeliveryDate',
        :'prime_estimated_weight' => :'primeEstimatedWeight',
        :'prime_estimated_weight_recorded_date' => :'primeEstimatedWeightRecordedDate',
        :'prime_actual_weight' => :'primeActualWeight',
        :'nts_recorded_weight' => :'ntsRecordedWeight',
        :'customer_remarks' => :'customerRemarks',
        :'counselor_remarks' => :'counselorRemarks',
        :'agents' => :'agents',
        :'sit_extensions' => :'sitExtensions',
        :'reweigh' => :'reweigh',
        :'pickup_address' => :'pickupAddress',
        :'destination_address' => :'destinationAddress',
        :'destination_type' => :'destinationType',
        :'secondary_pickup_address' => :'secondaryPickupAddress',
        :'secondary_delivery_address' => :'secondaryDeliveryAddress',
        :'storage_facility' => :'storageFacility',
        :'shipment_type' => :'shipmentType',
        :'diversion' => :'diversion',
        :'diversion_reason' => :'diversionReason',
        :'status' => :'status',
        :'ppm_shipment' => :'ppmShipment',
        :'delivery_address_update' => :'deliveryAddressUpdate',
        :'e_tag' => :'eTag',
        :'created_at' => :'createdAt',
        :'updated_at' => :'updatedAt',
        :'point_of_contact' => :'pointOfContact',
        :'origin_sit_auth_end_date' => :'originSitAuthEndDate',
        :'destination_sit_auth_end_date' => :'destinationSitAuthEndDate'
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
        :'move_task_order_id' => :'String',
        :'approved_date' => :'Date',
        :'requested_pickup_date' => :'Date',
        :'requested_delivery_date' => :'Date',
        :'scheduled_pickup_date' => :'Date',
        :'actual_pickup_date' => :'Date',
        :'first_available_delivery_date' => :'Date',
        :'required_delivery_date' => :'Date',
        :'scheduled_delivery_date' => :'Date',
        :'actual_delivery_date' => :'Date',
        :'prime_estimated_weight' => :'Integer',
        :'prime_estimated_weight_recorded_date' => :'Date',
        :'prime_actual_weight' => :'Integer',
        :'nts_recorded_weight' => :'Integer',
        :'customer_remarks' => :'String',
        :'counselor_remarks' => :'String',
        :'agents' => :'Array<MTOAgentV3>',
        :'sit_extensions' => :'Array<SITExtensionV3>',
        :'reweigh' => :'ReweighV3',
        :'pickup_address' => :'AddressV3',
        :'destination_address' => :'AddressV3',
        :'destination_type' => :'DestinationTypeV3',
        :'secondary_pickup_address' => :'AddressV3',
        :'secondary_delivery_address' => :'AddressV3',
        :'storage_facility' => :'UpdateMTOShipmentStorageFacilityV3',
        :'shipment_type' => :'MTOShipmentTypeV3',
        :'diversion' => :'Boolean',
        :'diversion_reason' => :'String',
        :'status' => :'String',
        :'ppm_shipment' => :'PPMShipmentV3',
        :'delivery_address_update' => :'ShipmentAddressUpdateV3',
        :'e_tag' => :'String',
        :'created_at' => :'Time',
        :'updated_at' => :'Time',
        :'point_of_contact' => :'String',
        :'origin_sit_auth_end_date' => :'Date',
        :'destination_sit_auth_end_date' => :'Date'
      }
    end

    # List of attributes with nullable: true
    def self.openapi_nullable
      Set.new([
        :'approved_date',
        :'requested_pickup_date',
        :'requested_delivery_date',
        :'scheduled_pickup_date',
        :'actual_pickup_date',
        :'first_available_delivery_date',
        :'required_delivery_date',
        :'scheduled_delivery_date',
        :'actual_delivery_date',
        :'prime_estimated_weight',
        :'prime_estimated_weight_recorded_date',
        :'prime_actual_weight',
        :'nts_recorded_weight',
        :'customer_remarks',
        :'counselor_remarks',
        :'destination_type',
        :'diversion_reason',
        :'ppm_shipment',
        :'origin_sit_auth_end_date',
        :'destination_sit_auth_end_date'
      ])
    end

    # Initializes the object
    # @param [Hash] attributes Model attributes in the form of hash
    def initialize(attributes = {})
      if (!attributes.is_a?(Hash))
        fail ArgumentError, "The input argument (attributes) must be a hash in `OpenapiClient::MTOShipmentWithoutServiceItemsV3` initialize method"
      end

      # check to see if the attribute exists and convert string to symbol for hash key
      attributes = attributes.each_with_object({}) { |(k, v), h|
        if (!self.class.attribute_map.key?(k.to_sym))
          fail ArgumentError, "`#{k}` is not a valid attribute in `OpenapiClient::MTOShipmentWithoutServiceItemsV3`. Please check the name to make sure it's valid. List of attributes: " + self.class.attribute_map.keys.inspect
        end
        h[k.to_sym] = v
      }

      if attributes.key?(:'id')
        self.id = attributes[:'id']
      end

      if attributes.key?(:'move_task_order_id')
        self.move_task_order_id = attributes[:'move_task_order_id']
      end

      if attributes.key?(:'approved_date')
        self.approved_date = attributes[:'approved_date']
      end

      if attributes.key?(:'requested_pickup_date')
        self.requested_pickup_date = attributes[:'requested_pickup_date']
      end

      if attributes.key?(:'requested_delivery_date')
        self.requested_delivery_date = attributes[:'requested_delivery_date']
      end

      if attributes.key?(:'scheduled_pickup_date')
        self.scheduled_pickup_date = attributes[:'scheduled_pickup_date']
      end

      if attributes.key?(:'actual_pickup_date')
        self.actual_pickup_date = attributes[:'actual_pickup_date']
      end

      if attributes.key?(:'first_available_delivery_date')
        self.first_available_delivery_date = attributes[:'first_available_delivery_date']
      end

      if attributes.key?(:'required_delivery_date')
        self.required_delivery_date = attributes[:'required_delivery_date']
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

      if attributes.key?(:'prime_estimated_weight_recorded_date')
        self.prime_estimated_weight_recorded_date = attributes[:'prime_estimated_weight_recorded_date']
      end

      if attributes.key?(:'prime_actual_weight')
        self.prime_actual_weight = attributes[:'prime_actual_weight']
      end

      if attributes.key?(:'nts_recorded_weight')
        self.nts_recorded_weight = attributes[:'nts_recorded_weight']
      end

      if attributes.key?(:'customer_remarks')
        self.customer_remarks = attributes[:'customer_remarks']
      end

      if attributes.key?(:'counselor_remarks')
        self.counselor_remarks = attributes[:'counselor_remarks']
      end

      if attributes.key?(:'agents')
        if (value = attributes[:'agents']).is_a?(Array)
          self.agents = value
        end
      end

      if attributes.key?(:'sit_extensions')
        if (value = attributes[:'sit_extensions']).is_a?(Array)
          self.sit_extensions = value
        end
      end

      if attributes.key?(:'reweigh')
        self.reweigh = attributes[:'reweigh']
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

      if attributes.key?(:'diversion_reason')
        self.diversion_reason = attributes[:'diversion_reason']
      end

      if attributes.key?(:'status')
        self.status = attributes[:'status']
      end

      if attributes.key?(:'ppm_shipment')
        self.ppm_shipment = attributes[:'ppm_shipment']
      end

      if attributes.key?(:'delivery_address_update')
        self.delivery_address_update = attributes[:'delivery_address_update']
      end

      if attributes.key?(:'e_tag')
        self.e_tag = attributes[:'e_tag']
      end

      if attributes.key?(:'created_at')
        self.created_at = attributes[:'created_at']
      end

      if attributes.key?(:'updated_at')
        self.updated_at = attributes[:'updated_at']
      end

      if attributes.key?(:'point_of_contact')
        self.point_of_contact = attributes[:'point_of_contact']
      end

      if attributes.key?(:'origin_sit_auth_end_date')
        self.origin_sit_auth_end_date = attributes[:'origin_sit_auth_end_date']
      end

      if attributes.key?(:'destination_sit_auth_end_date')
        self.destination_sit_auth_end_date = attributes[:'destination_sit_auth_end_date']
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

      if !@agents.nil? && @agents.length > 2
        invalid_properties.push('invalid value for "agents", number of items must be less than or equal to 2.')
      end

      invalid_properties
    end

    # Check to see if the all the properties in the model are valid
    # @return true if the model is valid
    def valid?
      warn '[DEPRECATED] the `valid?` method is obsolete'
      return false if !@prime_estimated_weight.nil? && @prime_estimated_weight < 1
      return false if !@prime_actual_weight.nil? && @prime_actual_weight < 1
      return false if !@agents.nil? && @agents.length > 2
      status_validator = EnumAttributeValidator.new('String', ["SUBMITTED", "APPROVED", "REJECTED", "CANCELLATION_REQUESTED", "CANCELED", "DIVERSION_REQUESTED"])
      return false unless status_validator.valid?(@status)
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

    # Custom attribute writer method with validation
    # @param [Object] agents Value to be assigned
    def agents=(agents)
      if agents.nil?
        fail ArgumentError, 'agents cannot be nil'
      end

      if agents.length > 2
        fail ArgumentError, 'invalid value for "agents", number of items must be less than or equal to 2.'
      end

      @agents = agents
    end

    # Custom attribute writer method checking allowed values (enum).
    # @param [Object] status Object to be assigned
    def status=(status)
      validator = EnumAttributeValidator.new('String', ["SUBMITTED", "APPROVED", "REJECTED", "CANCELLATION_REQUESTED", "CANCELED", "DIVERSION_REQUESTED"])
      unless validator.valid?(status)
        fail ArgumentError, "invalid value for \"status\", must be one of #{validator.allowable_values}."
      end
      @status = status
    end

    # Checks equality by comparing each attribute.
    # @param [Object] Object to be compared
    def ==(o)
      return true if self.equal?(o)
      self.class == o.class &&
          id == o.id &&
          move_task_order_id == o.move_task_order_id &&
          approved_date == o.approved_date &&
          requested_pickup_date == o.requested_pickup_date &&
          requested_delivery_date == o.requested_delivery_date &&
          scheduled_pickup_date == o.scheduled_pickup_date &&
          actual_pickup_date == o.actual_pickup_date &&
          first_available_delivery_date == o.first_available_delivery_date &&
          required_delivery_date == o.required_delivery_date &&
          scheduled_delivery_date == o.scheduled_delivery_date &&
          actual_delivery_date == o.actual_delivery_date &&
          prime_estimated_weight == o.prime_estimated_weight &&
          prime_estimated_weight_recorded_date == o.prime_estimated_weight_recorded_date &&
          prime_actual_weight == o.prime_actual_weight &&
          nts_recorded_weight == o.nts_recorded_weight &&
          customer_remarks == o.customer_remarks &&
          counselor_remarks == o.counselor_remarks &&
          agents == o.agents &&
          sit_extensions == o.sit_extensions &&
          reweigh == o.reweigh &&
          pickup_address == o.pickup_address &&
          destination_address == o.destination_address &&
          destination_type == o.destination_type &&
          secondary_pickup_address == o.secondary_pickup_address &&
          secondary_delivery_address == o.secondary_delivery_address &&
          storage_facility == o.storage_facility &&
          shipment_type == o.shipment_type &&
          diversion == o.diversion &&
          diversion_reason == o.diversion_reason &&
          status == o.status &&
          ppm_shipment == o.ppm_shipment &&
          delivery_address_update == o.delivery_address_update &&
          e_tag == o.e_tag &&
          created_at == o.created_at &&
          updated_at == o.updated_at &&
          point_of_contact == o.point_of_contact &&
          origin_sit_auth_end_date == o.origin_sit_auth_end_date &&
          destination_sit_auth_end_date == o.destination_sit_auth_end_date
    end

    # @see the `==` method
    # @param [Object] Object to be compared
    def eql?(o)
      self == o
    end

    # Calculates hash code according to all attributes.
    # @return [Integer] Hash code
    def hash
      [id, move_task_order_id, approved_date, requested_pickup_date, requested_delivery_date, scheduled_pickup_date, actual_pickup_date, first_available_delivery_date, required_delivery_date, scheduled_delivery_date, actual_delivery_date, prime_estimated_weight, prime_estimated_weight_recorded_date, prime_actual_weight, nts_recorded_weight, customer_remarks, counselor_remarks, agents, sit_extensions, reweigh, pickup_address, destination_address, destination_type, secondary_pickup_address, secondary_delivery_address, storage_facility, shipment_type, diversion, diversion_reason, status, ppm_shipment, delivery_address_update, e_tag, created_at, updated_at, point_of_contact, origin_sit_auth_end_date, destination_sit_auth_end_date].hash
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
