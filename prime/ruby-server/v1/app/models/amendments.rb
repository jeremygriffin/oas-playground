=begin
MilMove Prime API

The Prime API is a RESTful API that enables the Prime contractor to request information about upcoming moves, update the details and status of those moves, and make payment requests. It uses Mutual TLS for authentication procedures.  All endpoints are located at `/prime/v1/`. 

The version of the OpenAPI document: 0.0.1
Contact: milmove-developers@caci.com
Generated by: https://github.com/openapitools/openapi-generator.git

=end


class Amendments < ApplicationRecord
  validates_presence_of :total
  validates_presence_of :available_since

end
