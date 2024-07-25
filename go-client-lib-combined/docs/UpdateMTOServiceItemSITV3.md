# UpdateMTOServiceItemSITV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReServiceCode** | Pointer to **string** | Service code allowed for this model type. | [optional] 
**SitDepartureDate** | Pointer to **string** | Departure date for SIT. This is the end date of the SIT at either origin or destination. | [optional] 
**SitDestinationFinalAddress** | Pointer to [**AddressV3V3**](AddressV3.md) |  | [optional] 
**DateOfContact1** | Pointer to **NullableString** | Date of attempted contact by the prime corresponding to &#39;timeMilitary1&#39;. | [optional] 
**TimeMilitary1** | Pointer to **NullableString** | Time of attempted contact by the prime corresponding to &#39;dateOfContact1&#39;, in military format. | [optional] 
**FirstAvailableDeliveryDate1** | Pointer to **NullableString** | First available date that Prime can deliver SIT service item. | [optional] 
**DateOfContact2** | Pointer to **NullableString** | Date of attempted contact by the prime corresponding to &#39;timeMilitary2&#39;. | [optional] 
**TimeMilitary2** | Pointer to **NullableString** | Time of attempted contact by the prime corresponding to &#39;dateOfContact2&#39;, in military format. | [optional] 
**FirstAvailableDeliveryDate2** | Pointer to **NullableString** | Second available date that Prime can deliver SIT service item. | [optional] 
**SitRequestedDelivery** | Pointer to **NullableString** | Date when the customer has requested delivery out of SIT. | [optional] 
**SitCustomerContacted** | Pointer to **NullableString** | Date when the customer contacted the prime for a delivery out of SIT. | [optional] 
**UpdateReason** | Pointer to **NullableString** | Reason for updating service item. | [optional] 
**SitPostalCode** | Pointer to **NullableString** |  | [optional] 
**SitEntryDate** | Pointer to **NullableString** | Entry date for the SIT. | [optional] 
**RequestApprovalsRequestedStatus** | Pointer to **NullableBool** | Indicates if \&quot;Approvals Requested\&quot; status is being requested. | [optional] 

## Methods

### NewUpdateMTOServiceItemSITV3

`func NewUpdateMTOServiceItemSITV3() *UpdateMTOServiceItemSITV3`

NewUpdateMTOServiceItemSITV3 instantiates a new UpdateMTOServiceItemSITV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMTOServiceItemSITV3WithDefaults

`func NewUpdateMTOServiceItemSITV3WithDefaults() *UpdateMTOServiceItemSITV3`

NewUpdateMTOServiceItemSITV3WithDefaults instantiates a new UpdateMTOServiceItemSITV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReServiceCode

`func (o *UpdateMTOServiceItemSITV3) GetReServiceCode() string`

GetReServiceCode returns the ReServiceCode field if non-nil, zero value otherwise.

### GetReServiceCodeOk

`func (o *UpdateMTOServiceItemSITV3) GetReServiceCodeOk() (*string, bool)`

GetReServiceCodeOk returns a tuple with the ReServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReServiceCode

`func (o *UpdateMTOServiceItemSITV3) SetReServiceCode(v string)`

SetReServiceCode sets ReServiceCode field to given value.

### HasReServiceCode

`func (o *UpdateMTOServiceItemSITV3) HasReServiceCode() bool`

HasReServiceCode returns a boolean if a field has been set.

### GetSitDepartureDate

`func (o *UpdateMTOServiceItemSITV3) GetSitDepartureDate() string`

GetSitDepartureDate returns the SitDepartureDate field if non-nil, zero value otherwise.

### GetSitDepartureDateOk

`func (o *UpdateMTOServiceItemSITV3) GetSitDepartureDateOk() (*string, bool)`

GetSitDepartureDateOk returns a tuple with the SitDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitDepartureDate

`func (o *UpdateMTOServiceItemSITV3) SetSitDepartureDate(v string)`

SetSitDepartureDate sets SitDepartureDate field to given value.

### HasSitDepartureDate

`func (o *UpdateMTOServiceItemSITV3) HasSitDepartureDate() bool`

HasSitDepartureDate returns a boolean if a field has been set.

### GetSitDestinationFinalAddress

`func (o *UpdateMTOServiceItemSITV3) GetSitDestinationFinalAddress() AddressV3V3`

GetSitDestinationFinalAddress returns the SitDestinationFinalAddress field if non-nil, zero value otherwise.

### GetSitDestinationFinalAddressOk

`func (o *UpdateMTOServiceItemSITV3) GetSitDestinationFinalAddressOk() (*AddressV3V3, bool)`

GetSitDestinationFinalAddressOk returns a tuple with the SitDestinationFinalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitDestinationFinalAddress

`func (o *UpdateMTOServiceItemSITV3) SetSitDestinationFinalAddress(v AddressV3V3)`

SetSitDestinationFinalAddress sets SitDestinationFinalAddress field to given value.

### HasSitDestinationFinalAddress

`func (o *UpdateMTOServiceItemSITV3) HasSitDestinationFinalAddress() bool`

HasSitDestinationFinalAddress returns a boolean if a field has been set.

### GetDateOfContact1

`func (o *UpdateMTOServiceItemSITV3) GetDateOfContact1() string`

GetDateOfContact1 returns the DateOfContact1 field if non-nil, zero value otherwise.

### GetDateOfContact1Ok

`func (o *UpdateMTOServiceItemSITV3) GetDateOfContact1Ok() (*string, bool)`

GetDateOfContact1Ok returns a tuple with the DateOfContact1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfContact1

`func (o *UpdateMTOServiceItemSITV3) SetDateOfContact1(v string)`

SetDateOfContact1 sets DateOfContact1 field to given value.

### HasDateOfContact1

`func (o *UpdateMTOServiceItemSITV3) HasDateOfContact1() bool`

HasDateOfContact1 returns a boolean if a field has been set.

### SetDateOfContact1Nil

`func (o *UpdateMTOServiceItemSITV3) SetDateOfContact1Nil(b bool)`

 SetDateOfContact1Nil sets the value for DateOfContact1 to be an explicit nil

### UnsetDateOfContact1
`func (o *UpdateMTOServiceItemSITV3) UnsetDateOfContact1()`

UnsetDateOfContact1 ensures that no value is present for DateOfContact1, not even an explicit nil
### GetTimeMilitary1

`func (o *UpdateMTOServiceItemSITV3) GetTimeMilitary1() string`

GetTimeMilitary1 returns the TimeMilitary1 field if non-nil, zero value otherwise.

### GetTimeMilitary1Ok

`func (o *UpdateMTOServiceItemSITV3) GetTimeMilitary1Ok() (*string, bool)`

GetTimeMilitary1Ok returns a tuple with the TimeMilitary1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMilitary1

`func (o *UpdateMTOServiceItemSITV3) SetTimeMilitary1(v string)`

SetTimeMilitary1 sets TimeMilitary1 field to given value.

### HasTimeMilitary1

`func (o *UpdateMTOServiceItemSITV3) HasTimeMilitary1() bool`

HasTimeMilitary1 returns a boolean if a field has been set.

### SetTimeMilitary1Nil

`func (o *UpdateMTOServiceItemSITV3) SetTimeMilitary1Nil(b bool)`

 SetTimeMilitary1Nil sets the value for TimeMilitary1 to be an explicit nil

### UnsetTimeMilitary1
`func (o *UpdateMTOServiceItemSITV3) UnsetTimeMilitary1()`

UnsetTimeMilitary1 ensures that no value is present for TimeMilitary1, not even an explicit nil
### GetFirstAvailableDeliveryDate1

`func (o *UpdateMTOServiceItemSITV3) GetFirstAvailableDeliveryDate1() string`

GetFirstAvailableDeliveryDate1 returns the FirstAvailableDeliveryDate1 field if non-nil, zero value otherwise.

### GetFirstAvailableDeliveryDate1Ok

`func (o *UpdateMTOServiceItemSITV3) GetFirstAvailableDeliveryDate1Ok() (*string, bool)`

GetFirstAvailableDeliveryDate1Ok returns a tuple with the FirstAvailableDeliveryDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAvailableDeliveryDate1

`func (o *UpdateMTOServiceItemSITV3) SetFirstAvailableDeliveryDate1(v string)`

SetFirstAvailableDeliveryDate1 sets FirstAvailableDeliveryDate1 field to given value.

### HasFirstAvailableDeliveryDate1

`func (o *UpdateMTOServiceItemSITV3) HasFirstAvailableDeliveryDate1() bool`

HasFirstAvailableDeliveryDate1 returns a boolean if a field has been set.

### SetFirstAvailableDeliveryDate1Nil

`func (o *UpdateMTOServiceItemSITV3) SetFirstAvailableDeliveryDate1Nil(b bool)`

 SetFirstAvailableDeliveryDate1Nil sets the value for FirstAvailableDeliveryDate1 to be an explicit nil

### UnsetFirstAvailableDeliveryDate1
`func (o *UpdateMTOServiceItemSITV3) UnsetFirstAvailableDeliveryDate1()`

UnsetFirstAvailableDeliveryDate1 ensures that no value is present for FirstAvailableDeliveryDate1, not even an explicit nil
### GetDateOfContact2

`func (o *UpdateMTOServiceItemSITV3) GetDateOfContact2() string`

GetDateOfContact2 returns the DateOfContact2 field if non-nil, zero value otherwise.

### GetDateOfContact2Ok

`func (o *UpdateMTOServiceItemSITV3) GetDateOfContact2Ok() (*string, bool)`

GetDateOfContact2Ok returns a tuple with the DateOfContact2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfContact2

`func (o *UpdateMTOServiceItemSITV3) SetDateOfContact2(v string)`

SetDateOfContact2 sets DateOfContact2 field to given value.

### HasDateOfContact2

`func (o *UpdateMTOServiceItemSITV3) HasDateOfContact2() bool`

HasDateOfContact2 returns a boolean if a field has been set.

### SetDateOfContact2Nil

`func (o *UpdateMTOServiceItemSITV3) SetDateOfContact2Nil(b bool)`

 SetDateOfContact2Nil sets the value for DateOfContact2 to be an explicit nil

### UnsetDateOfContact2
`func (o *UpdateMTOServiceItemSITV3) UnsetDateOfContact2()`

UnsetDateOfContact2 ensures that no value is present for DateOfContact2, not even an explicit nil
### GetTimeMilitary2

`func (o *UpdateMTOServiceItemSITV3) GetTimeMilitary2() string`

GetTimeMilitary2 returns the TimeMilitary2 field if non-nil, zero value otherwise.

### GetTimeMilitary2Ok

`func (o *UpdateMTOServiceItemSITV3) GetTimeMilitary2Ok() (*string, bool)`

GetTimeMilitary2Ok returns a tuple with the TimeMilitary2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMilitary2

`func (o *UpdateMTOServiceItemSITV3) SetTimeMilitary2(v string)`

SetTimeMilitary2 sets TimeMilitary2 field to given value.

### HasTimeMilitary2

`func (o *UpdateMTOServiceItemSITV3) HasTimeMilitary2() bool`

HasTimeMilitary2 returns a boolean if a field has been set.

### SetTimeMilitary2Nil

`func (o *UpdateMTOServiceItemSITV3) SetTimeMilitary2Nil(b bool)`

 SetTimeMilitary2Nil sets the value for TimeMilitary2 to be an explicit nil

### UnsetTimeMilitary2
`func (o *UpdateMTOServiceItemSITV3) UnsetTimeMilitary2()`

UnsetTimeMilitary2 ensures that no value is present for TimeMilitary2, not even an explicit nil
### GetFirstAvailableDeliveryDate2

`func (o *UpdateMTOServiceItemSITV3) GetFirstAvailableDeliveryDate2() string`

GetFirstAvailableDeliveryDate2 returns the FirstAvailableDeliveryDate2 field if non-nil, zero value otherwise.

### GetFirstAvailableDeliveryDate2Ok

`func (o *UpdateMTOServiceItemSITV3) GetFirstAvailableDeliveryDate2Ok() (*string, bool)`

GetFirstAvailableDeliveryDate2Ok returns a tuple with the FirstAvailableDeliveryDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAvailableDeliveryDate2

`func (o *UpdateMTOServiceItemSITV3) SetFirstAvailableDeliveryDate2(v string)`

SetFirstAvailableDeliveryDate2 sets FirstAvailableDeliveryDate2 field to given value.

### HasFirstAvailableDeliveryDate2

`func (o *UpdateMTOServiceItemSITV3) HasFirstAvailableDeliveryDate2() bool`

HasFirstAvailableDeliveryDate2 returns a boolean if a field has been set.

### SetFirstAvailableDeliveryDate2Nil

`func (o *UpdateMTOServiceItemSITV3) SetFirstAvailableDeliveryDate2Nil(b bool)`

 SetFirstAvailableDeliveryDate2Nil sets the value for FirstAvailableDeliveryDate2 to be an explicit nil

### UnsetFirstAvailableDeliveryDate2
`func (o *UpdateMTOServiceItemSITV3) UnsetFirstAvailableDeliveryDate2()`

UnsetFirstAvailableDeliveryDate2 ensures that no value is present for FirstAvailableDeliveryDate2, not even an explicit nil
### GetSitRequestedDelivery

`func (o *UpdateMTOServiceItemSITV3) GetSitRequestedDelivery() string`

GetSitRequestedDelivery returns the SitRequestedDelivery field if non-nil, zero value otherwise.

### GetSitRequestedDeliveryOk

`func (o *UpdateMTOServiceItemSITV3) GetSitRequestedDeliveryOk() (*string, bool)`

GetSitRequestedDeliveryOk returns a tuple with the SitRequestedDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitRequestedDelivery

`func (o *UpdateMTOServiceItemSITV3) SetSitRequestedDelivery(v string)`

SetSitRequestedDelivery sets SitRequestedDelivery field to given value.

### HasSitRequestedDelivery

`func (o *UpdateMTOServiceItemSITV3) HasSitRequestedDelivery() bool`

HasSitRequestedDelivery returns a boolean if a field has been set.

### SetSitRequestedDeliveryNil

`func (o *UpdateMTOServiceItemSITV3) SetSitRequestedDeliveryNil(b bool)`

 SetSitRequestedDeliveryNil sets the value for SitRequestedDelivery to be an explicit nil

### UnsetSitRequestedDelivery
`func (o *UpdateMTOServiceItemSITV3) UnsetSitRequestedDelivery()`

UnsetSitRequestedDelivery ensures that no value is present for SitRequestedDelivery, not even an explicit nil
### GetSitCustomerContacted

`func (o *UpdateMTOServiceItemSITV3) GetSitCustomerContacted() string`

GetSitCustomerContacted returns the SitCustomerContacted field if non-nil, zero value otherwise.

### GetSitCustomerContactedOk

`func (o *UpdateMTOServiceItemSITV3) GetSitCustomerContactedOk() (*string, bool)`

GetSitCustomerContactedOk returns a tuple with the SitCustomerContacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitCustomerContacted

`func (o *UpdateMTOServiceItemSITV3) SetSitCustomerContacted(v string)`

SetSitCustomerContacted sets SitCustomerContacted field to given value.

### HasSitCustomerContacted

`func (o *UpdateMTOServiceItemSITV3) HasSitCustomerContacted() bool`

HasSitCustomerContacted returns a boolean if a field has been set.

### SetSitCustomerContactedNil

`func (o *UpdateMTOServiceItemSITV3) SetSitCustomerContactedNil(b bool)`

 SetSitCustomerContactedNil sets the value for SitCustomerContacted to be an explicit nil

### UnsetSitCustomerContacted
`func (o *UpdateMTOServiceItemSITV3) UnsetSitCustomerContacted()`

UnsetSitCustomerContacted ensures that no value is present for SitCustomerContacted, not even an explicit nil
### GetUpdateReason

`func (o *UpdateMTOServiceItemSITV3) GetUpdateReason() string`

GetUpdateReason returns the UpdateReason field if non-nil, zero value otherwise.

### GetUpdateReasonOk

`func (o *UpdateMTOServiceItemSITV3) GetUpdateReasonOk() (*string, bool)`

GetUpdateReasonOk returns a tuple with the UpdateReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateReason

`func (o *UpdateMTOServiceItemSITV3) SetUpdateReason(v string)`

SetUpdateReason sets UpdateReason field to given value.

### HasUpdateReason

`func (o *UpdateMTOServiceItemSITV3) HasUpdateReason() bool`

HasUpdateReason returns a boolean if a field has been set.

### SetUpdateReasonNil

`func (o *UpdateMTOServiceItemSITV3) SetUpdateReasonNil(b bool)`

 SetUpdateReasonNil sets the value for UpdateReason to be an explicit nil

### UnsetUpdateReason
`func (o *UpdateMTOServiceItemSITV3) UnsetUpdateReason()`

UnsetUpdateReason ensures that no value is present for UpdateReason, not even an explicit nil
### GetSitPostalCode

`func (o *UpdateMTOServiceItemSITV3) GetSitPostalCode() string`

GetSitPostalCode returns the SitPostalCode field if non-nil, zero value otherwise.

### GetSitPostalCodeOk

`func (o *UpdateMTOServiceItemSITV3) GetSitPostalCodeOk() (*string, bool)`

GetSitPostalCodeOk returns a tuple with the SitPostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitPostalCode

`func (o *UpdateMTOServiceItemSITV3) SetSitPostalCode(v string)`

SetSitPostalCode sets SitPostalCode field to given value.

### HasSitPostalCode

`func (o *UpdateMTOServiceItemSITV3) HasSitPostalCode() bool`

HasSitPostalCode returns a boolean if a field has been set.

### SetSitPostalCodeNil

`func (o *UpdateMTOServiceItemSITV3) SetSitPostalCodeNil(b bool)`

 SetSitPostalCodeNil sets the value for SitPostalCode to be an explicit nil

### UnsetSitPostalCode
`func (o *UpdateMTOServiceItemSITV3) UnsetSitPostalCode()`

UnsetSitPostalCode ensures that no value is present for SitPostalCode, not even an explicit nil
### GetSitEntryDate

`func (o *UpdateMTOServiceItemSITV3) GetSitEntryDate() string`

GetSitEntryDate returns the SitEntryDate field if non-nil, zero value otherwise.

### GetSitEntryDateOk

`func (o *UpdateMTOServiceItemSITV3) GetSitEntryDateOk() (*string, bool)`

GetSitEntryDateOk returns a tuple with the SitEntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEntryDate

`func (o *UpdateMTOServiceItemSITV3) SetSitEntryDate(v string)`

SetSitEntryDate sets SitEntryDate field to given value.

### HasSitEntryDate

`func (o *UpdateMTOServiceItemSITV3) HasSitEntryDate() bool`

HasSitEntryDate returns a boolean if a field has been set.

### SetSitEntryDateNil

`func (o *UpdateMTOServiceItemSITV3) SetSitEntryDateNil(b bool)`

 SetSitEntryDateNil sets the value for SitEntryDate to be an explicit nil

### UnsetSitEntryDate
`func (o *UpdateMTOServiceItemSITV3) UnsetSitEntryDate()`

UnsetSitEntryDate ensures that no value is present for SitEntryDate, not even an explicit nil
### GetRequestApprovalsRequestedStatus

`func (o *UpdateMTOServiceItemSITV3) GetRequestApprovalsRequestedStatus() bool`

GetRequestApprovalsRequestedStatus returns the RequestApprovalsRequestedStatus field if non-nil, zero value otherwise.

### GetRequestApprovalsRequestedStatusOk

`func (o *UpdateMTOServiceItemSITV3) GetRequestApprovalsRequestedStatusOk() (*bool, bool)`

GetRequestApprovalsRequestedStatusOk returns a tuple with the RequestApprovalsRequestedStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestApprovalsRequestedStatus

`func (o *UpdateMTOServiceItemSITV3) SetRequestApprovalsRequestedStatus(v bool)`

SetRequestApprovalsRequestedStatus sets RequestApprovalsRequestedStatus field to given value.

### HasRequestApprovalsRequestedStatus

`func (o *UpdateMTOServiceItemSITV3) HasRequestApprovalsRequestedStatus() bool`

HasRequestApprovalsRequestedStatus returns a boolean if a field has been set.

### SetRequestApprovalsRequestedStatusNil

`func (o *UpdateMTOServiceItemSITV3) SetRequestApprovalsRequestedStatusNil(b bool)`

 SetRequestApprovalsRequestedStatusNil sets the value for RequestApprovalsRequestedStatus to be an explicit nil

### UnsetRequestApprovalsRequestedStatus
`func (o *UpdateMTOServiceItemSITV3) UnsetRequestApprovalsRequestedStatus()`

UnsetRequestApprovalsRequestedStatus ensures that no value is present for RequestApprovalsRequestedStatus, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


