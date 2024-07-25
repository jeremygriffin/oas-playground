# MTOServiceItemOriginSITV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReServiceCode** | **string** | Service code allowed for this model type. | 
**Reason** | **string** | Explanation of why Prime is picking up SIT item. | 
**SitPostalCode** | **string** |  | 
**SitEntryDate** | **string** | Entry date for the SIT | 
**SitDepartureDate** | Pointer to **NullableString** | Departure date for SIT. This is the end date of the SIT at either origin or destination. This is optional as it can be updated using the UpdateMTOServiceItemSIT modelType at a later date. | [optional] 
**SitHHGActualOrigin** | Pointer to [**AddressV3V3**](AddressV3.md) |  | [optional] 
**SitHHGOriginalOrigin** | Pointer to [**AddressV3V3**](AddressV3.md) |  | [optional] 
**RequestApprovalsRequestedStatus** | Pointer to **bool** |  | [optional] 
**SitRequestedDelivery** | Pointer to **NullableString** | Date when the customer has requested delivery out of SIT. | [optional] 
**SitCustomerContacted** | Pointer to **NullableString** | Date when the customer contacted the prime for a delivery out of SIT. | [optional] 

## Methods

### NewMTOServiceItemOriginSITV3

`func NewMTOServiceItemOriginSITV3(reServiceCode string, reason string, sitPostalCode string, sitEntryDate string, ) *MTOServiceItemOriginSITV3`

NewMTOServiceItemOriginSITV3 instantiates a new MTOServiceItemOriginSITV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTOServiceItemOriginSITV3WithDefaults

`func NewMTOServiceItemOriginSITV3WithDefaults() *MTOServiceItemOriginSITV3`

NewMTOServiceItemOriginSITV3WithDefaults instantiates a new MTOServiceItemOriginSITV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReServiceCode

`func (o *MTOServiceItemOriginSITV3) GetReServiceCode() string`

GetReServiceCode returns the ReServiceCode field if non-nil, zero value otherwise.

### GetReServiceCodeOk

`func (o *MTOServiceItemOriginSITV3) GetReServiceCodeOk() (*string, bool)`

GetReServiceCodeOk returns a tuple with the ReServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReServiceCode

`func (o *MTOServiceItemOriginSITV3) SetReServiceCode(v string)`

SetReServiceCode sets ReServiceCode field to given value.


### GetReason

`func (o *MTOServiceItemOriginSITV3) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MTOServiceItemOriginSITV3) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MTOServiceItemOriginSITV3) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSitPostalCode

`func (o *MTOServiceItemOriginSITV3) GetSitPostalCode() string`

GetSitPostalCode returns the SitPostalCode field if non-nil, zero value otherwise.

### GetSitPostalCodeOk

`func (o *MTOServiceItemOriginSITV3) GetSitPostalCodeOk() (*string, bool)`

GetSitPostalCodeOk returns a tuple with the SitPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitPostalCode

`func (o *MTOServiceItemOriginSITV3) SetSitPostalCode(v string)`

SetSitPostalCode sets SitPostalCode field to given value.


### GetSitEntryDate

`func (o *MTOServiceItemOriginSITV3) GetSitEntryDate() string`

GetSitEntryDate returns the SitEntryDate field if non-nil, zero value otherwise.

### GetSitEntryDateOk

`func (o *MTOServiceItemOriginSITV3) GetSitEntryDateOk() (*string, bool)`

GetSitEntryDateOk returns a tuple with the SitEntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEntryDate

`func (o *MTOServiceItemOriginSITV3) SetSitEntryDate(v string)`

SetSitEntryDate sets SitEntryDate field to given value.


### GetSitDepartureDate

`func (o *MTOServiceItemOriginSITV3) GetSitDepartureDate() string`

GetSitDepartureDate returns the SitDepartureDate field if non-nil, zero value otherwise.

### GetSitDepartureDateOk

`func (o *MTOServiceItemOriginSITV3) GetSitDepartureDateOk() (*string, bool)`

GetSitDepartureDateOk returns a tuple with the SitDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitDepartureDate

`func (o *MTOServiceItemOriginSITV3) SetSitDepartureDate(v string)`

SetSitDepartureDate sets SitDepartureDate field to given value.

### HasSitDepartureDate

`func (o *MTOServiceItemOriginSITV3) HasSitDepartureDate() bool`

HasSitDepartureDate returns a boolean if a field has been set.

### SetSitDepartureDateNil

`func (o *MTOServiceItemOriginSITV3) SetSitDepartureDateNil(b bool)`

 SetSitDepartureDateNil sets the value for SitDepartureDate to be an explicit nil

### UnsetSitDepartureDate
`func (o *MTOServiceItemOriginSITV3) UnsetSitDepartureDate()`

UnsetSitDepartureDate ensures that no value is present for SitDepartureDate, not even an explicit nil
### GetSitHHGActualOrigin

`func (o *MTOServiceItemOriginSITV3) GetSitHHGActualOrigin() AddressV3V3`

GetSitHHGActualOrigin returns the SitHHGActualOrigin field if non-nil, zero value otherwise.

### GetSitHHGActualOriginOk

`func (o *MTOServiceItemOriginSITV3) GetSitHHGActualOriginOk() (*AddressV3V3, bool)`

GetSitHHGActualOriginOk returns a tuple with the SitHHGActualOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitHHGActualOrigin

`func (o *MTOServiceItemOriginSITV3) SetSitHHGActualOrigin(v AddressV3V3)`

SetSitHHGActualOrigin sets SitHHGActualOrigin field to given value.

### HasSitHHGActualOrigin

`func (o *MTOServiceItemOriginSITV3) HasSitHHGActualOrigin() bool`

HasSitHHGActualOrigin returns a boolean if a field has been set.

### GetSitHHGOriginalOrigin

`func (o *MTOServiceItemOriginSITV3) GetSitHHGOriginalOrigin() AddressV3V3`

GetSitHHGOriginalOrigin returns the SitHHGOriginalOrigin field if non-nil, zero value otherwise.

### GetSitHHGOriginalOriginOk

`func (o *MTOServiceItemOriginSITV3) GetSitHHGOriginalOriginOk() (*AddressV3V3, bool)`

GetSitHHGOriginalOriginOk returns a tuple with the SitHHGOriginalOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitHHGOriginalOrigin

`func (o *MTOServiceItemOriginSITV3) SetSitHHGOriginalOrigin(v AddressV3V3)`

SetSitHHGOriginalOrigin sets SitHHGOriginalOrigin field to given value.

### HasSitHHGOriginalOrigin

`func (o *MTOServiceItemOriginSITV3) HasSitHHGOriginalOrigin() bool`

HasSitHHGOriginalOrigin returns a boolean if a field has been set.

### GetRequestApprovalsRequestedStatus

`func (o *MTOServiceItemOriginSITV3) GetRequestApprovalsRequestedStatus() bool`

GetRequestApprovalsRequestedStatus returns the RequestApprovalsRequestedStatus field if non-nil, zero value otherwise.

### GetRequestApprovalsRequestedStatusOk

`func (o *MTOServiceItemOriginSITV3) GetRequestApprovalsRequestedStatusOk() (*bool, bool)`

GetRequestApprovalsRequestedStatusOk returns a tuple with the RequestApprovalsRequestedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestApprovalsRequestedStatus

`func (o *MTOServiceItemOriginSITV3) SetRequestApprovalsRequestedStatus(v bool)`

SetRequestApprovalsRequestedStatus sets RequestApprovalsRequestedStatus field to given value.

### HasRequestApprovalsRequestedStatus

`func (o *MTOServiceItemOriginSITV3) HasRequestApprovalsRequestedStatus() bool`

HasRequestApprovalsRequestedStatus returns a boolean if a field has been set.

### GetSitRequestedDelivery

`func (o *MTOServiceItemOriginSITV3) GetSitRequestedDelivery() string`

GetSitRequestedDelivery returns the SitRequestedDelivery field if non-nil, zero value otherwise.

### GetSitRequestedDeliveryOk

`func (o *MTOServiceItemOriginSITV3) GetSitRequestedDeliveryOk() (*string, bool)`

GetSitRequestedDeliveryOk returns a tuple with the SitRequestedDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitRequestedDelivery

`func (o *MTOServiceItemOriginSITV3) SetSitRequestedDelivery(v string)`

SetSitRequestedDelivery sets SitRequestedDelivery field to given value.

### HasSitRequestedDelivery

`func (o *MTOServiceItemOriginSITV3) HasSitRequestedDelivery() bool`

HasSitRequestedDelivery returns a boolean if a field has been set.

### SetSitRequestedDeliveryNil

`func (o *MTOServiceItemOriginSITV3) SetSitRequestedDeliveryNil(b bool)`

 SetSitRequestedDeliveryNil sets the value for SitRequestedDelivery to be an explicit nil

### UnsetSitRequestedDelivery
`func (o *MTOServiceItemOriginSITV3) UnsetSitRequestedDelivery()`

UnsetSitRequestedDelivery ensures that no value is present for SitRequestedDelivery, not even an explicit nil
### GetSitCustomerContacted

`func (o *MTOServiceItemOriginSITV3) GetSitCustomerContacted() string`

GetSitCustomerContacted returns the SitCustomerContacted field if non-nil, zero value otherwise.

### GetSitCustomerContactedOk

`func (o *MTOServiceItemOriginSITV3) GetSitCustomerContactedOk() (*string, bool)`

GetSitCustomerContactedOk returns a tuple with the SitCustomerContacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitCustomerContacted

`func (o *MTOServiceItemOriginSITV3) SetSitCustomerContacted(v string)`

SetSitCustomerContacted sets SitCustomerContacted field to given value.

### HasSitCustomerContacted

`func (o *MTOServiceItemOriginSITV3) HasSitCustomerContacted() bool`

HasSitCustomerContacted returns a boolean if a field has been set.

### SetSitCustomerContactedNil

`func (o *MTOServiceItemOriginSITV3) SetSitCustomerContactedNil(b bool)`

 SetSitCustomerContactedNil sets the value for SitCustomerContacted to be an explicit nil

### UnsetSitCustomerContacted
`func (o *MTOServiceItemOriginSITV3) UnsetSitCustomerContacted()`

UnsetSitCustomerContacted ensures that no value is present for SitCustomerContacted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


