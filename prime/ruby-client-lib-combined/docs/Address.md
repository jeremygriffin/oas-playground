# OpenapiClient::Address

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **street_address1** | **String** |  |  |
| **street_address2** | **String** |  | [optional] |
| **street_address3** | **String** |  | [optional] |
| **city** | **String** |  |  |
| **e_tag** | **String** |  | [optional][readonly] |
| **state** | **String** |  |  |
| **postal_code** | **String** |  |  |
| **country** | **String** |  | [optional][default to &#39;USA&#39;] |
| **county** | **String** |  | [optional] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::Address.new(
  id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  street_address1: 123 Main Ave,
  street_address2: Apartment 9000,
  street_address3: Montmârtre,
  city: Anytown,
  e_tag: null,
  state: null,
  postal_code: 90210,
  country: USA,
  county: LOS ANGELES
)
```

