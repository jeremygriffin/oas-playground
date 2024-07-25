=begin
MilMove Prime V2 API

The Prime V2 API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v2/`. 

The version of the OpenAPI document: 0.0.1
Contact: milmove-developers@caci.com
Generated by: https://github.com/openapitools/openapi-generator.git

=end

class InitTables < ActiveRecord::Migration
  def change
    create_table "address".pluralize.to_sym, id: false do |t|
      t.String :id
      t.string :street_address1
      t.string :street_address2
      t.string :street_address3
      t.string :city
      t.string :e_tag
      t.string :state
      t.string :postal_code
      t.string :country
      t.string :county

      t.timestamps
    end

    create_table "client_error".pluralize.to_sym, id: false do |t|
      t.string :title
      t.string :detail
      t.String :instance

      t.timestamps
    end

    create_table "create_mto_shipment".pluralize.to_sym, id: false do |t|
      t.String :move_task_order_id
      t.Date :requested_pickup_date
      t.integer :prime_estimated_weight
      t.string :customer_remarks
      t.string :agents
      t.string :mto_service_items
      t.string :pickup_address
      t.string :destination_address
      t.string :shipment_type
      t.boolean :diversion
      t.String :diverted_from_shipment_id
      t.string :point_of_contact
      t.string :counselor_remarks
      t.string :ppm_shipment

      t.timestamps
    end

    create_table "create_ppm_shipment".pluralize.to_sym, id: false do |t|
      t.Date :expected_departure_date
      t.string :pickup_address
      t.string :destination_address
      t.boolean :sit_expected
      t.string :sit_location
      t.integer :sit_estimated_weight
      t.Date :sit_estimated_entry_date
      t.Date :sit_estimated_departure_date
      t.integer :estimated_weight
      t.boolean :has_pro_gear
      t.integer :pro_gear_weight
      t.integer :spouse_pro_gear_weight

      t.timestamps
    end

    create_table "create_sit_address_update_request".pluralize.to_sym, id: false do |t|
      t.string :new_address
      t.string :contractor_remarks
      t.String :mto_service_item_id

      t.timestamps
    end

    create_table "create_sit_extension".pluralize.to_sym, id: false do |t|
      t.string :request_reason
      t.string :contractor_remarks
      t.integer :requested_days

      t.timestamps
    end

    create_table "customer".pluralize.to_sym, id: false do |t|
      t.String :id
      t.string :dod_id
      t.string :emplid
      t.String :user_id
      t.string :current_address
      t.string :first_name
      t.string :last_name
      t.string :branch
      t.string :phone
      t.string :email
      t.string :e_tag

      t.timestamps
    end

    create_table "destination_type".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "duty_location".pluralize.to_sym, id: false do |t|
      t.String :id
      t.string :name
      t.String :address_id
      t.string :address
      t.string :e_tag

      t.timestamps
    end

    create_table "entitlements".pluralize.to_sym, id: false do |t|
      t.String :id
      t.integer :authorized_weight
      t.boolean :dependents_authorized
      t.boolean :gun_safe
      t.boolean :non_temporary_storage
      t.boolean :privately_owned_vehicle
      t.integer :pro_gear_weight
      t.integer :pro_gear_weight_spouse
      t.integer :required_medical_equipment_weight
      t.boolean :organizational_clothing_and_individual_equipment
      t.integer :storage_in_transit
      t.integer :total_weight
      t.integer :total_dependents
      t.string :e_tag

      t.timestamps
    end

    create_table "error".pluralize.to_sym, id: false do |t|
      t.string :title
      t.string :detail
      t.String :instance

      t.timestamps
    end

    create_table "mto_agent".pluralize.to_sym, id: false do |t|
      t.String :id
      t.String :mto_shipment_id
      t.datetime :created_at
      t.datetime :updated_at
      t.string :first_name
      t.string :last_name
      t.string :email
      t.string :phone
      t.string :agent_type
      t.string :e_tag

      t.timestamps
    end

    create_table "mto_agent_type".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "mto_service_item".pluralize.to_sym, id: false do |t|
      t.String :id
      t.String :move_task_order_id
      t.String :mto_shipment_id
      t.string :re_service_name
      t.string :status
      t.string :rejection_reason
      t.string :model_type
      t.string :service_request_documents
      t.string :e_tag

      t.timestamps
    end

    create_table "mto_service_item_basic".pluralize.to_sym, id: false do |t|
      t.String :id
      t.String :move_task_order_id
      t.String :mto_shipment_id
      t.string :re_service_name
      t.string :status
      t.string :rejection_reason
      t.string :model_type
      t.string :service_request_documents
      t.string :e_tag
      t.string :re_service_code

      t.timestamps
    end

    create_table "mto_service_item_dest_sit".pluralize.to_sym, id: false do |t|
      t.String :id
      t.String :move_task_order_id
      t.String :mto_shipment_id
      t.string :re_service_name
      t.string :status
      t.string :rejection_reason
      t.string :model_type
      t.string :service_request_documents
      t.string :e_tag
      t.string :re_service_code
      t.Date :date_of_contact1
      t.Date :date_of_contact2
      t.string :time_military1
      t.string :time_military2
      t.Date :first_available_delivery_date1
      t.Date :first_available_delivery_date2
      t.Date :sit_entry_date
      t.Date :sit_departure_date
      t.string :sit_destination_final_address
      t.string :reason
      t.string :sit_address_updates
      t.Date :sit_requested_delivery
      t.Date :sit_customer_contacted

      t.timestamps
    end

    create_table "mto_service_item_dimension".pluralize.to_sym, id: false do |t|
      t.String :id
      t.integer :length
      t.integer :width
      t.integer :height

      t.timestamps
    end

    create_table "mto_service_item_domestic_crating".pluralize.to_sym, id: false do |t|
      t.String :id
      t.String :move_task_order_id
      t.String :mto_shipment_id
      t.string :re_service_name
      t.string :status
      t.string :rejection_reason
      t.string :model_type
      t.string :service_request_documents
      t.string :e_tag
      t.string :re_service_code
      t.Object :item
      t.Object :crate
      t.string :description
      t.string :reason
      t.boolean :standalone_crate

      t.timestamps
    end

    create_table "mto_service_item_model_type".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "mto_service_item_origin_sit".pluralize.to_sym, id: false do |t|
      t.String :id
      t.String :move_task_order_id
      t.String :mto_shipment_id
      t.string :re_service_name
      t.string :status
      t.string :rejection_reason
      t.string :model_type
      t.string :service_request_documents
      t.string :e_tag
      t.string :re_service_code
      t.string :reason
      t.string :sit_postal_code
      t.Date :sit_entry_date
      t.Date :sit_departure_date
      t.string :sit_hhg_actual_origin
      t.string :sit_hhg_original_origin
      t.boolean :request_approvals_requested_status
      t.Date :sit_requested_delivery
      t.Date :sit_customer_contacted

      t.timestamps
    end

    create_table "mto_service_item_shuttle".pluralize.to_sym, id: false do |t|
      t.String :id
      t.String :move_task_order_id
      t.String :mto_shipment_id
      t.string :re_service_name
      t.string :status
      t.string :rejection_reason
      t.string :model_type
      t.string :service_request_documents
      t.string :e_tag
      t.string :re_service_code
      t.string :reason
      t.integer :estimated_weight
      t.integer :actual_weight

      t.timestamps
    end

    create_table "mto_service_item_status".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "mto_shipment".pluralize.to_sym, id: false do |t|
      t.string :mto_service_items
      t.String :id
      t.String :move_task_order_id
      t.Date :approved_date
      t.Date :requested_pickup_date
      t.Date :requested_delivery_date
      t.Date :scheduled_pickup_date
      t.Date :actual_pickup_date
      t.Date :first_available_delivery_date
      t.Date :required_delivery_date
      t.Date :scheduled_delivery_date
      t.Date :actual_delivery_date
      t.integer :prime_estimated_weight
      t.Date :prime_estimated_weight_recorded_date
      t.integer :prime_actual_weight
      t.integer :nts_recorded_weight
      t.string :customer_remarks
      t.string :counselor_remarks
      t.integer :actual_pro_gear_weight
      t.integer :actual_spouse_pro_gear_weight
      t.string :agents
      t.string :sit_extensions
      t.string :reweigh
      t.string :pickup_address
      t.string :destination_address
      t.string :destination_type
      t.string :secondary_pickup_address
      t.string :secondary_delivery_address
      t.string :storage_facility
      t.string :shipment_type
      t.boolean :diversion
      t.string :diversion_reason
      t.string :status
      t.string :ppm_shipment
      t.string :delivery_address_update
      t.string :e_tag
      t.datetime :created_at
      t.datetime :updated_at
      t.string :point_of_contact
      t.Date :origin_sit_auth_end_date
      t.Date :destination_sit_auth_end_date

      t.timestamps
    end

    create_table "mto_shipment_type".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "mto_shipment_without_service_items".pluralize.to_sym, id: false do |t|
      t.String :id
      t.String :move_task_order_id
      t.Date :approved_date
      t.Date :requested_pickup_date
      t.Date :requested_delivery_date
      t.Date :scheduled_pickup_date
      t.Date :actual_pickup_date
      t.Date :first_available_delivery_date
      t.Date :required_delivery_date
      t.Date :scheduled_delivery_date
      t.Date :actual_delivery_date
      t.integer :prime_estimated_weight
      t.Date :prime_estimated_weight_recorded_date
      t.integer :prime_actual_weight
      t.integer :nts_recorded_weight
      t.string :customer_remarks
      t.string :counselor_remarks
      t.integer :actual_pro_gear_weight
      t.integer :actual_spouse_pro_gear_weight
      t.string :agents
      t.string :sit_extensions
      t.string :reweigh
      t.string :pickup_address
      t.string :destination_address
      t.string :destination_type
      t.string :secondary_pickup_address
      t.string :secondary_delivery_address
      t.string :storage_facility
      t.string :shipment_type
      t.boolean :diversion
      t.string :diversion_reason
      t.string :status
      t.string :ppm_shipment
      t.string :delivery_address_update
      t.string :e_tag
      t.datetime :created_at
      t.datetime :updated_at
      t.string :point_of_contact
      t.Date :origin_sit_auth_end_date
      t.Date :destination_sit_auth_end_date

      t.timestamps
    end

    create_table "move_task_order".pluralize.to_sym, id: false do |t|
      t.String :id
      t.string :move_code
      t.datetime :created_at
      t.String :order_id
      t.string :order
      t.string :reference_id
      t.datetime :available_to_prime_at
      t.datetime :updated_at
      t.datetime :prime_counseling_completed_at
      t.string :payment_requests
      t.string :mto_service_items
      t.string :mto_shipments
      t.string :ppm_type
      t.integer :ppm_estimated_weight
      t.datetime :excess_weight_qualified_at
      t.datetime :excess_weight_acknowledged_at
      t.String :excess_weight_upload_id
      t.string :contract_number
      t.string :e_tag

      t.timestamps
    end

    create_table "order".pluralize.to_sym, id: false do |t|
      t.String :id
      t.string :customer
      t.String :customer_id
      t.string :entitlement
      t.string :destination_duty_location
      t.string :origin_duty_location
      t.string :origin_duty_location_gbloc
      t.string :rank
      t.Date :report_by_date
      t.string :orders_type
      t.string :order_number
      t.string :lines_of_accounting
      t.string :supply_and_services_cost_estimate
      t.string :packing_and_shipping_instructions
      t.string :method_of_payment
      t.string :naics
      t.string :e_tag

      t.timestamps
    end

    create_table "orders_type".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "ppm_shipment".pluralize.to_sym, id: false do |t|
      t.String :id
      t.String :shipment_id
      t.datetime :created_at
      t.datetime :updated_at
      t.string :status
      t.Date :expected_departure_date
      t.Date :actual_move_date
      t.datetime :submitted_at
      t.datetime :reviewed_at
      t.datetime :approved_at
      t.string :actual_pickup_postal_code
      t.string :actual_destination_postal_code
      t.boolean :sit_expected
      t.integer :estimated_weight
      t.boolean :has_pro_gear
      t.integer :pro_gear_weight
      t.integer :spouse_pro_gear_weight
      t.integer :estimated_incentive
      t.boolean :has_requested_advance
      t.integer :advance_amount_requested
      t.boolean :has_received_advance
      t.integer :advance_amount_received
      t.string :sit_location
      t.integer :sit_estimated_weight
      t.Date :sit_estimated_entry_date
      t.Date :sit_estimated_departure_date
      t.integer :sit_estimated_cost
      t.string :e_tag

      t.timestamps
    end

    create_table "ppm_shipment_status".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "payment_request".pluralize.to_sym, id: false do |t|
      t.String :id
      t.boolean :is_final
      t.String :move_task_order_id
      t.string :rejection_reason
      t.string :status
      t.string :payment_request_number
      t.String :recalculation_of_payment_request_id
      t.string :proof_of_service_docs
      t.string :payment_service_items
      t.string :e_tag

      t.timestamps
    end

    create_table "payment_request_status".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "payment_service_item".pluralize.to_sym, id: false do |t|
      t.String :id
      t.String :payment_request_id
      t.String :mto_service_item_id
      t.string :status
      t.integer :price_cents
      t.string :rejection_reason
      t.Object :reference_id
      t.string :payment_service_item_params
      t.string :e_tag

      t.timestamps
    end

    create_table "payment_service_item_param".pluralize.to_sym, id: false do |t|
      t.String :id
      t.String :payment_service_item_id
      t.string :key
      t.string :value
      t.string :type
      t.string :origin
      t.string :e_tag

      t.timestamps
    end

    create_table "payment_service_item_status".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "proof_of_service_doc".pluralize.to_sym, id: false do |t|
      t.string :uploads

      t.timestamps
    end

    create_table "re_service_code".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "reweigh".pluralize.to_sym, id: false do |t|
      t.String :id
      t.datetime :requested_at
      t.string :requested_by
      t.datetime :verification_provided_at
      t.string :verification_reason
      t.integer :weight
      t.String :shipment_id
      t.datetime :created_at
      t.datetime :updated_at
      t.string :e_tag

      t.timestamps
    end

    create_table "reweigh_requester".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "sit_extension".pluralize.to_sym, id: false do |t|
      t.String :id
      t.String :mto_shipment_id
      t.string :request_reason
      t.string :contractor_remarks
      t.integer :requested_days
      t.string :status
      t.integer :approved_days
      t.datetime :decision_date
      t.string :office_remarks
      t.datetime :created_at
      t.datetime :updated_at
      t.string :e_tag

      t.timestamps
    end

    create_table "sit_location_type".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "service_item_param_name".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "service_item_param_origin".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "service_item_param_type".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "service_request_document".pluralize.to_sym, id: false do |t|
      t.string :uploads

      t.timestamps
    end

    create_table "shipment_address_update".pluralize.to_sym, id: false do |t|
      t.String :id
      t.string :contractor_remarks
      t.string :office_remarks
      t.string :status
      t.String :shipment_id
      t.string :original_address
      t.string :new_address
      t.string :sit_original_address
      t.integer :old_sit_distance_between
      t.integer :new_sit_distance_between

      t.timestamps
    end

    create_table "shipment_address_update_status".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "sit_address_update".pluralize.to_sym, id: false do |t|
      t.String :id
      t.String :mto_service_item_id
      t.String :new_address_id
      t.string :new_address
      t.String :old_address_id
      t.string :old_address
      t.string :status
      t.integer :distance
      t.string :contractor_remarks
      t.string :office_remarks
      t.datetime :created_at
      t.datetime :updated_at
      t.string :e_tag

      t.timestamps
    end

    create_table "sit_address_update_status".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "storage_facility".pluralize.to_sym, id: false do |t|
      t.String :id
      t.string :facility_name
      t.string :address
      t.string :lot_number
      t.string :phone
      t.string :email
      t.string :e_tag

      t.timestamps
    end

    create_table "update_mto_service_item".pluralize.to_sym, id: false do |t|
      t.String :id
      t.string :model_type

      t.timestamps
    end

    create_table "update_mto_service_item_model_type".pluralize.to_sym, id: false do |t|

      t.timestamps
    end

    create_table "update_mto_service_item_sit".pluralize.to_sym, id: false do |t|
      t.String :id
      t.string :model_type
      t.string :re_service_code
      t.Date :sit_departure_date
      t.string :sit_destination_final_address
      t.Date :date_of_contact1
      t.string :time_military1
      t.Date :first_available_delivery_date1
      t.Date :date_of_contact2
      t.string :time_military2
      t.Date :first_available_delivery_date2
      t.Date :sit_requested_delivery
      t.Date :sit_customer_contacted
      t.string :update_reason
      t.string :sit_postal_code
      t.Date :sit_entry_date
      t.boolean :request_approvals_requested_status

      t.timestamps
    end

    create_table "update_mto_service_item_shuttle".pluralize.to_sym, id: false do |t|
      t.String :id
      t.string :model_type
      t.integer :actual_weight
      t.integer :estimated_weight
      t.string :re_service_code

      t.timestamps
    end

    create_table "update_mto_shipment".pluralize.to_sym, id: false do |t|
      t.Date :scheduled_pickup_date
      t.Date :actual_pickup_date
      t.Date :first_available_delivery_date
      t.Date :scheduled_delivery_date
      t.Date :actual_delivery_date
      t.integer :prime_estimated_weight
      t.integer :prime_actual_weight
      t.integer :nts_recorded_weight
      t.string :pickup_address
      t.string :destination_address
      t.string :destination_type
      t.string :secondary_pickup_address
      t.string :secondary_delivery_address
      t.string :storage_facility
      t.string :shipment_type
      t.boolean :diversion
      t.string :point_of_contact
      t.string :counselor_remarks
      t.integer :actual_pro_gear_weight
      t.integer :actual_spouse_pro_gear_weight
      t.string :ppm_shipment

      t.timestamps
    end

    create_table "update_mto_shipment_status".pluralize.to_sym, id: false do |t|
      t.string :status

      t.timestamps
    end

    create_table "update_mto_shipment_storage_facility".pluralize.to_sym, id: false do |t|
      t.String :id
      t.string :facility_name
      t.string :address
      t.string :lot_number
      t.string :phone
      t.string :email
      t.string :e_tag

      t.timestamps
    end

    create_table "update_ppm_shipment".pluralize.to_sym, id: false do |t|
      t.Date :expected_departure_date
      t.boolean :sit_expected
      t.string :sit_location
      t.integer :sit_estimated_weight
      t.Date :sit_estimated_entry_date
      t.Date :sit_estimated_departure_date
      t.integer :estimated_weight
      t.boolean :has_pro_gear
      t.integer :pro_gear_weight
      t.integer :spouse_pro_gear_weight

      t.timestamps
    end

    create_table "update_reweigh".pluralize.to_sym, id: false do |t|
      t.integer :weight
      t.string :verification_reason

      t.timestamps
    end

    create_table "update_shipment_destination_address".pluralize.to_sym, id: false do |t|
      t.string :new_address
      t.string :contractor_remarks

      t.timestamps
    end

    create_table "upload_with_omissions".pluralize.to_sym, id: false do |t|
      t.String :id
      t.String :url
      t.string :filename
      t.string :content_type
      t.integer :bytes
      t.string :status
      t.datetime :created_at
      t.datetime :updated_at

      t.timestamps
    end

    create_table "validation_error".pluralize.to_sym, id: false do |t|
      t.string :title
      t.string :detail
      t.String :instance
      t.string :invalid_fields

      t.timestamps
    end

  end
end
