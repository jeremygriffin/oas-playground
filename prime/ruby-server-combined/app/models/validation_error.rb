=begin
MilMove Prime V3 API

The Prime V3 API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v3/`. 

The version of the OpenAPI document: 0.0.1
Contact: milmove-developers@caci.com
Generated by: https://github.com/openapitools/openapi-generator.git

=end


class ValidationError < ApplicationRecord
  validates_presence_of :title
  validates_presence_of :detail
  validates_presence_of :instance
  validates_presence_of :invalid_fields

  serialize :invalid_fields, Hash
end
