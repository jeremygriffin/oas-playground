# MTOServiceItemOriginSIT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReServiceCode** | **string** | Service code allowed for this model type. | 
**Reason** | **string** | Explanation of why Prime is picking up SIT item. | 
**SitPostalCode** | **string** |  | 
**SitEntryDate** | **string** | Entry date for the SIT | 
**SitDepartureDate** | Pointer to **NullableString** | Departure date for SIT. This is the end date of the SIT at either origin or destination. This is optional as it can be updated using the UpdateMTOServiceItemSIT modelType at a later date. | [optional] 
**SitHHGActualOrigin** | Pointer to [**Address**](Address.md) |  | [optional] 
**SitHHGOriginalOrigin** | Pointer to [**Address**](Address.md) |  | [optional] 
**RequestApprovalsRequestedStatus** | Pointer to **bool** |  | [optional] 
**SitRequestedDelivery** | Pointer to **NullableString** | Date when the customer has requested delivery out of SIT. | [optional] 
**SitCustomerContacted** | Pointer to **NullableString** | Date when the customer contacted the prime for a delivery out of SIT. | [optional] 

## Methods

### NewMTOServiceItemOriginSIT

`func NewMTOServiceItemOriginSIT(reServiceCode string, reason string, sitPostalCode string, sitEntryDate string, ) *MTOServiceItemOriginSIT`

NewMTOServiceItemOriginSIT instantiates a new MTOServiceItemOriginSIT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTOServiceItemOriginSITWithDefaults

`func NewMTOServiceItemOriginSITWithDefaults() *MTOServiceItemOriginSIT`

NewMTOServiceItemOriginSITWithDefaults instantiates a new MTOServiceItemOriginSIT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReServiceCode

`func (o *MTOServiceItemOriginSIT) GetReServiceCode() string`

GetReServiceCode returns the ReServiceCode field if non-nil, zero value otherwise.

### GetReServiceCodeOk

`func (o *MTOServiceItemOriginSIT) GetReServiceCodeOk() (*string, bool)`

GetReServiceCodeOk returns a tuple with the ReServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReServiceCode

`func (o *MTOServiceItemOriginSIT) SetReServiceCode(v string)`

SetReServiceCode sets ReServiceCode field to given value.


### GetReason

`func (o *MTOServiceItemOriginSIT) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MTOServiceItemOriginSIT) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MTOServiceItemOriginSIT) SetReason(v string)`

SetReason sets Reason field to given value.


### GetSitPostalCode

`func (o *MTOServiceItemOriginSIT) GetSitPostalCode() string`

GetSitPostalCode returns the SitPostalCode field if non-nil, zero value otherwise.

### GetSitPostalCodeOk

`func (o *MTOServiceItemOriginSIT) GetSitPostalCodeOk() (*string, bool)`

GetSitPostalCodeOk returns a tuple with the SitPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitPostalCode

`func (o *MTOServiceItemOriginSIT) SetSitPostalCode(v string)`

SetSitPostalCode sets SitPostalCode field to given value.


### GetSitEntryDate

`func (o *MTOServiceItemOriginSIT) GetSitEntryDate() string`

GetSitEntryDate returns the SitEntryDate field if non-nil, zero value otherwise.

### GetSitEntryDateOk

`func (o *MTOServiceItemOriginSIT) GetSitEntryDateOk() (*string, bool)`

GetSitEntryDateOk returns a tuple with the SitEntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEntryDate

`func (o *MTOServiceItemOriginSIT) SetSitEntryDate(v string)`

SetSitEntryDate sets SitEntryDate field to given value.


### GetSitDepartureDate

`func (o *MTOServiceItemOriginSIT) GetSitDepartureDate() string`

GetSitDepartureDate returns the SitDepartureDate field if non-nil, zero value otherwise.

### GetSitDepartureDateOk

`func (o *MTOServiceItemOriginSIT) GetSitDepartureDateOk() (*string, bool)`

GetSitDepartureDateOk returns a tuple with the SitDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitDepartureDate

`func (o *MTOServiceItemOriginSIT) SetSitDepartureDate(v string)`

SetSitDepartureDate sets SitDepartureDate field to given value.

### HasSitDepartureDate

`func (o *MTOServiceItemOriginSIT) HasSitDepartureDate() bool`

HasSitDepartureDate returns a boolean if a field has been set.

### SetSitDepartureDateNil

`func (o *MTOServiceItemOriginSIT) SetSitDepartureDateNil(b bool)`

 SetSitDepartureDateNil sets the value for SitDepartureDate to be an explicit nil

### UnsetSitDepartureDate
`func (o *MTOServiceItemOriginSIT) UnsetSitDepartureDate()`

UnsetSitDepartureDate ensures that no value is present for SitDepartureDate, not even an explicit nil
### GetSitHHGActualOrigin

`func (o *MTOServiceItemOriginSIT) GetSitHHGActualOrigin() Address`

GetSitHHGActualOrigin returns the SitHHGActualOrigin field if non-nil, zero value otherwise.

### GetSitHHGActualOriginOk

`func (o *MTOServiceItemOriginSIT) GetSitHHGActualOriginOk() (*Address, bool)`

GetSitHHGActualOriginOk returns a tuple with the SitHHGActualOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitHHGActualOrigin

`func (o *MTOServiceItemOriginSIT) SetSitHHGActualOrigin(v Address)`

SetSitHHGActualOrigin sets SitHHGActualOrigin field to given value.

### HasSitHHGActualOrigin

`func (o *MTOServiceItemOriginSIT) HasSitHHGActualOrigin() bool`

HasSitHHGActualOrigin returns a boolean if a field has been set.

### GetSitHHGOriginalOrigin

`func (o *MTOServiceItemOriginSIT) GetSitHHGOriginalOrigin() Address`

GetSitHHGOriginalOrigin returns the SitHHGOriginalOrigin field if non-nil, zero value otherwise.

### GetSitHHGOriginalOriginOk

`func (o *MTOServiceItemOriginSIT) GetSitHHGOriginalOriginOk() (*Address, bool)`

GetSitHHGOriginalOriginOk returns a tuple with the SitHHGOriginalOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitHHGOriginalOrigin

`func (o *MTOServiceItemOriginSIT) SetSitHHGOriginalOrigin(v Address)`

SetSitHHGOriginalOrigin sets SitHHGOriginalOrigin field to given value.

### HasSitHHGOriginalOrigin

`func (o *MTOServiceItemOriginSIT) HasSitHHGOriginalOrigin() bool`

HasSitHHGOriginalOrigin returns a boolean if a field has been set.

### GetRequestApprovalsRequestedStatus

`func (o *MTOServiceItemOriginSIT) GetRequestApprovalsRequestedStatus() bool`

GetRequestApprovalsRequestedStatus returns the RequestApprovalsRequestedStatus field if non-nil, zero value otherwise.

### GetRequestApprovalsRequestedStatusOk

`func (o *MTOServiceItemOriginSIT) GetRequestApprovalsRequestedStatusOk() (*bool, bool)`

GetRequestApprovalsRequestedStatusOk returns a tuple with the RequestApprovalsRequestedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestApprovalsRequestedStatus

`func (o *MTOServiceItemOriginSIT) SetRequestApprovalsRequestedStatus(v bool)`

SetRequestApprovalsRequestedStatus sets RequestApprovalsRequestedStatus field to given value.

### HasRequestApprovalsRequestedStatus

`func (o *MTOServiceItemOriginSIT) HasRequestApprovalsRequestedStatus() bool`

HasRequestApprovalsRequestedStatus returns a boolean if a field has been set.

### GetSitRequestedDelivery

`func (o *MTOServiceItemOriginSIT) GetSitRequestedDelivery() string`

GetSitRequestedDelivery returns the SitRequestedDelivery field if non-nil, zero value otherwise.

### GetSitRequestedDeliveryOk

`func (o *MTOServiceItemOriginSIT) GetSitRequestedDeliveryOk() (*string, bool)`

GetSitRequestedDeliveryOk returns a tuple with the SitRequestedDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitRequestedDelivery

`func (o *MTOServiceItemOriginSIT) SetSitRequestedDelivery(v string)`

SetSitRequestedDelivery sets SitRequestedDelivery field to given value.

### HasSitRequestedDelivery

`func (o *MTOServiceItemOriginSIT) HasSitRequestedDelivery() bool`

HasSitRequestedDelivery returns a boolean if a field has been set.

### SetSitRequestedDeliveryNil

`func (o *MTOServiceItemOriginSIT) SetSitRequestedDeliveryNil(b bool)`

 SetSitRequestedDeliveryNil sets the value for SitRequestedDelivery to be an explicit nil

### UnsetSitRequestedDelivery
`func (o *MTOServiceItemOriginSIT) UnsetSitRequestedDelivery()`

UnsetSitRequestedDelivery ensures that no value is present for SitRequestedDelivery, not even an explicit nil
### GetSitCustomerContacted

`func (o *MTOServiceItemOriginSIT) GetSitCustomerContacted() string`

GetSitCustomerContacted returns the SitCustomerContacted field if non-nil, zero value otherwise.

### GetSitCustomerContactedOk

`func (o *MTOServiceItemOriginSIT) GetSitCustomerContactedOk() (*string, bool)`

GetSitCustomerContactedOk returns a tuple with the SitCustomerContacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitCustomerContacted

`func (o *MTOServiceItemOriginSIT) SetSitCustomerContacted(v string)`

SetSitCustomerContacted sets SitCustomerContacted field to given value.

### HasSitCustomerContacted

`func (o *MTOServiceItemOriginSIT) HasSitCustomerContacted() bool`

HasSitCustomerContacted returns a boolean if a field has been set.

### SetSitCustomerContactedNil

`func (o *MTOServiceItemOriginSIT) SetSitCustomerContactedNil(b bool)`

 SetSitCustomerContactedNil sets the value for SitCustomerContacted to be an explicit nil

### UnsetSitCustomerContacted
`func (o *MTOServiceItemOriginSIT) UnsetSitCustomerContacted()`

UnsetSitCustomerContacted ensures that no value is present for SitCustomerContacted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


