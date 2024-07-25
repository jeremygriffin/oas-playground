# OpenapiClient::ServiceItem

## Properties

| Name | Type | Description | Notes |
| ---- | ---- | ----------- | ----- |
| **id** | **String** |  | [optional] |
| **params** | [**Array&lt;ServiceItemParamsInner&gt;**](ServiceItemParamsInner.md) | This should be populated for the following service items:   * DOASIT(Domestic origin Additional day SIT)   * DDASIT(Domestic destination Additional day SIT)  Both take in the following param keys:   * &#x60;SITPaymentRequestStart&#x60;   * &#x60;SITPaymentRequestEnd&#x60;  The value of each is a date string in the format \&quot;YYYY-MM-DD\&quot; (e.g. \&quot;2023-01-15\&quot;)  | [optional] |
| **e_tag** | **String** |  | [optional][readonly] |

## Example

```ruby
require 'openapi_client'

instance = OpenapiClient::ServiceItem.new(
  id: c56a4180-65aa-42ec-a945-5fd21dec0538,
  params: null,
  e_tag: null
)
```

