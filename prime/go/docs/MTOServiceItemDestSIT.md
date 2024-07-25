# MTOServiceItemDestSIT

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReServiceCode** | **string** | Service code allowed for this model type. | 
**DateOfContact1** | Pointer to **NullableString** | Date of attempted contact by the prime corresponding to &#x60;timeMilitary1&#x60;. | [optional] 
**DateOfContact2** | Pointer to **NullableString** | Date of attempted contact by the prime corresponding to &#x60;timeMilitary2&#x60;. | [optional] 
**TimeMilitary1** | Pointer to **NullableString** | Time of attempted contact corresponding to &#x60;dateOfContact1&#x60;, in military format. | [optional] 
**TimeMilitary2** | Pointer to **NullableString** | Time of attempted contact corresponding to &#x60;dateOfContact2&#x60;, in military format. | [optional] 
**FirstAvailableDeliveryDate1** | Pointer to **NullableString** | First available date that Prime can deliver SIT service item. | [optional] 
**FirstAvailableDeliveryDate2** | Pointer to **NullableString** | Second available date that Prime can deliver SIT service item. | [optional] 
**SitEntryDate** | **string** | Entry date for the SIT | 
**SitDepartureDate** | Pointer to **NullableString** | Departure date for SIT. This is the end date of the SIT at either origin or destination. This is optional as it can be updated using the UpdateMTOServiceItemSIT modelType at a later date. | [optional] 
**SitDestinationFinalAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**Reason** | **NullableString** | The reason item has been placed in SIT.  | 
**SitAddressUpdates** | Pointer to [**[]SitAddressUpdate**](SitAddressUpdate.md) | A list of updates to a SIT service item address. | [optional] 
**SitRequestedDelivery** | Pointer to **NullableString** | Date when the customer has requested delivery out of SIT. | [optional] 
**SitCustomerContacted** | Pointer to **NullableString** | Date when the customer contacted the prime for a delivery out of SIT. | [optional] 

## Methods

### NewMTOServiceItemDestSIT

`func NewMTOServiceItemDestSIT(reServiceCode string, sitEntryDate string, reason NullableString, ) *MTOServiceItemDestSIT`

NewMTOServiceItemDestSIT instantiates a new MTOServiceItemDestSIT object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTOServiceItemDestSITWithDefaults

`func NewMTOServiceItemDestSITWithDefaults() *MTOServiceItemDestSIT`

NewMTOServiceItemDestSITWithDefaults instantiates a new MTOServiceItemDestSIT object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReServiceCode

`func (o *MTOServiceItemDestSIT) GetReServiceCode() string`

GetReServiceCode returns the ReServiceCode field if non-nil, zero value otherwise.

### GetReServiceCodeOk

`func (o *MTOServiceItemDestSIT) GetReServiceCodeOk() (*string, bool)`

GetReServiceCodeOk returns a tuple with the ReServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReServiceCode

`func (o *MTOServiceItemDestSIT) SetReServiceCode(v string)`

SetReServiceCode sets ReServiceCode field to given value.


### GetDateOfContact1

`func (o *MTOServiceItemDestSIT) GetDateOfContact1() string`

GetDateOfContact1 returns the DateOfContact1 field if non-nil, zero value otherwise.

### GetDateOfContact1Ok

`func (o *MTOServiceItemDestSIT) GetDateOfContact1Ok() (*string, bool)`

GetDateOfContact1Ok returns a tuple with the DateOfContact1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfContact1

`func (o *MTOServiceItemDestSIT) SetDateOfContact1(v string)`

SetDateOfContact1 sets DateOfContact1 field to given value.

### HasDateOfContact1

`func (o *MTOServiceItemDestSIT) HasDateOfContact1() bool`

HasDateOfContact1 returns a boolean if a field has been set.

### SetDateOfContact1Nil

`func (o *MTOServiceItemDestSIT) SetDateOfContact1Nil(b bool)`

 SetDateOfContact1Nil sets the value for DateOfContact1 to be an explicit nil

### UnsetDateOfContact1
`func (o *MTOServiceItemDestSIT) UnsetDateOfContact1()`

UnsetDateOfContact1 ensures that no value is present for DateOfContact1, not even an explicit nil
### GetDateOfContact2

`func (o *MTOServiceItemDestSIT) GetDateOfContact2() string`

GetDateOfContact2 returns the DateOfContact2 field if non-nil, zero value otherwise.

### GetDateOfContact2Ok

`func (o *MTOServiceItemDestSIT) GetDateOfContact2Ok() (*string, bool)`

GetDateOfContact2Ok returns a tuple with the DateOfContact2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateOfContact2

`func (o *MTOServiceItemDestSIT) SetDateOfContact2(v string)`

SetDateOfContact2 sets DateOfContact2 field to given value.

### HasDateOfContact2

`func (o *MTOServiceItemDestSIT) HasDateOfContact2() bool`

HasDateOfContact2 returns a boolean if a field has been set.

### SetDateOfContact2Nil

`func (o *MTOServiceItemDestSIT) SetDateOfContact2Nil(b bool)`

 SetDateOfContact2Nil sets the value for DateOfContact2 to be an explicit nil

### UnsetDateOfContact2
`func (o *MTOServiceItemDestSIT) UnsetDateOfContact2()`

UnsetDateOfContact2 ensures that no value is present for DateOfContact2, not even an explicit nil
### GetTimeMilitary1

`func (o *MTOServiceItemDestSIT) GetTimeMilitary1() string`

GetTimeMilitary1 returns the TimeMilitary1 field if non-nil, zero value otherwise.

### GetTimeMilitary1Ok

`func (o *MTOServiceItemDestSIT) GetTimeMilitary1Ok() (*string, bool)`

GetTimeMilitary1Ok returns a tuple with the TimeMilitary1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMilitary1

`func (o *MTOServiceItemDestSIT) SetTimeMilitary1(v string)`

SetTimeMilitary1 sets TimeMilitary1 field to given value.

### HasTimeMilitary1

`func (o *MTOServiceItemDestSIT) HasTimeMilitary1() bool`

HasTimeMilitary1 returns a boolean if a field has been set.

### SetTimeMilitary1Nil

`func (o *MTOServiceItemDestSIT) SetTimeMilitary1Nil(b bool)`

 SetTimeMilitary1Nil sets the value for TimeMilitary1 to be an explicit nil

### UnsetTimeMilitary1
`func (o *MTOServiceItemDestSIT) UnsetTimeMilitary1()`

UnsetTimeMilitary1 ensures that no value is present for TimeMilitary1, not even an explicit nil
### GetTimeMilitary2

`func (o *MTOServiceItemDestSIT) GetTimeMilitary2() string`

GetTimeMilitary2 returns the TimeMilitary2 field if non-nil, zero value otherwise.

### GetTimeMilitary2Ok

`func (o *MTOServiceItemDestSIT) GetTimeMilitary2Ok() (*string, bool)`

GetTimeMilitary2Ok returns a tuple with the TimeMilitary2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeMilitary2

`func (o *MTOServiceItemDestSIT) SetTimeMilitary2(v string)`

SetTimeMilitary2 sets TimeMilitary2 field to given value.

### HasTimeMilitary2

`func (o *MTOServiceItemDestSIT) HasTimeMilitary2() bool`

HasTimeMilitary2 returns a boolean if a field has been set.

### SetTimeMilitary2Nil

`func (o *MTOServiceItemDestSIT) SetTimeMilitary2Nil(b bool)`

 SetTimeMilitary2Nil sets the value for TimeMilitary2 to be an explicit nil

### UnsetTimeMilitary2
`func (o *MTOServiceItemDestSIT) UnsetTimeMilitary2()`

UnsetTimeMilitary2 ensures that no value is present for TimeMilitary2, not even an explicit nil
### GetFirstAvailableDeliveryDate1

`func (o *MTOServiceItemDestSIT) GetFirstAvailableDeliveryDate1() string`

GetFirstAvailableDeliveryDate1 returns the FirstAvailableDeliveryDate1 field if non-nil, zero value otherwise.

### GetFirstAvailableDeliveryDate1Ok

`func (o *MTOServiceItemDestSIT) GetFirstAvailableDeliveryDate1Ok() (*string, bool)`

GetFirstAvailableDeliveryDate1Ok returns a tuple with the FirstAvailableDeliveryDate1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAvailableDeliveryDate1

`func (o *MTOServiceItemDestSIT) SetFirstAvailableDeliveryDate1(v string)`

SetFirstAvailableDeliveryDate1 sets FirstAvailableDeliveryDate1 field to given value.

### HasFirstAvailableDeliveryDate1

`func (o *MTOServiceItemDestSIT) HasFirstAvailableDeliveryDate1() bool`

HasFirstAvailableDeliveryDate1 returns a boolean if a field has been set.

### SetFirstAvailableDeliveryDate1Nil

`func (o *MTOServiceItemDestSIT) SetFirstAvailableDeliveryDate1Nil(b bool)`

 SetFirstAvailableDeliveryDate1Nil sets the value for FirstAvailableDeliveryDate1 to be an explicit nil

### UnsetFirstAvailableDeliveryDate1
`func (o *MTOServiceItemDestSIT) UnsetFirstAvailableDeliveryDate1()`

UnsetFirstAvailableDeliveryDate1 ensures that no value is present for FirstAvailableDeliveryDate1, not even an explicit nil
### GetFirstAvailableDeliveryDate2

`func (o *MTOServiceItemDestSIT) GetFirstAvailableDeliveryDate2() string`

GetFirstAvailableDeliveryDate2 returns the FirstAvailableDeliveryDate2 field if non-nil, zero value otherwise.

### GetFirstAvailableDeliveryDate2Ok

`func (o *MTOServiceItemDestSIT) GetFirstAvailableDeliveryDate2Ok() (*string, bool)`

GetFirstAvailableDeliveryDate2Ok returns a tuple with the FirstAvailableDeliveryDate2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstAvailableDeliveryDate2

`func (o *MTOServiceItemDestSIT) SetFirstAvailableDeliveryDate2(v string)`

SetFirstAvailableDeliveryDate2 sets FirstAvailableDeliveryDate2 field to given value.

### HasFirstAvailableDeliveryDate2

`func (o *MTOServiceItemDestSIT) HasFirstAvailableDeliveryDate2() bool`

HasFirstAvailableDeliveryDate2 returns a boolean if a field has been set.

### SetFirstAvailableDeliveryDate2Nil

`func (o *MTOServiceItemDestSIT) SetFirstAvailableDeliveryDate2Nil(b bool)`

 SetFirstAvailableDeliveryDate2Nil sets the value for FirstAvailableDeliveryDate2 to be an explicit nil

### UnsetFirstAvailableDeliveryDate2
`func (o *MTOServiceItemDestSIT) UnsetFirstAvailableDeliveryDate2()`

UnsetFirstAvailableDeliveryDate2 ensures that no value is present for FirstAvailableDeliveryDate2, not even an explicit nil
### GetSitEntryDate

`func (o *MTOServiceItemDestSIT) GetSitEntryDate() string`

GetSitEntryDate returns the SitEntryDate field if non-nil, zero value otherwise.

### GetSitEntryDateOk

`func (o *MTOServiceItemDestSIT) GetSitEntryDateOk() (*string, bool)`

GetSitEntryDateOk returns a tuple with the SitEntryDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitEntryDate

`func (o *MTOServiceItemDestSIT) SetSitEntryDate(v string)`

SetSitEntryDate sets SitEntryDate field to given value.


### GetSitDepartureDate

`func (o *MTOServiceItemDestSIT) GetSitDepartureDate() string`

GetSitDepartureDate returns the SitDepartureDate field if non-nil, zero value otherwise.

### GetSitDepartureDateOk

`func (o *MTOServiceItemDestSIT) GetSitDepartureDateOk() (*string, bool)`

GetSitDepartureDateOk returns a tuple with the SitDepartureDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitDepartureDate

`func (o *MTOServiceItemDestSIT) SetSitDepartureDate(v string)`

SetSitDepartureDate sets SitDepartureDate field to given value.

### HasSitDepartureDate

`func (o *MTOServiceItemDestSIT) HasSitDepartureDate() bool`

HasSitDepartureDate returns a boolean if a field has been set.

### SetSitDepartureDateNil

`func (o *MTOServiceItemDestSIT) SetSitDepartureDateNil(b bool)`

 SetSitDepartureDateNil sets the value for SitDepartureDate to be an explicit nil

### UnsetSitDepartureDate
`func (o *MTOServiceItemDestSIT) UnsetSitDepartureDate()`

UnsetSitDepartureDate ensures that no value is present for SitDepartureDate, not even an explicit nil
### GetSitDestinationFinalAddress

`func (o *MTOServiceItemDestSIT) GetSitDestinationFinalAddress() Address`

GetSitDestinationFinalAddress returns the SitDestinationFinalAddress field if non-nil, zero value otherwise.

### GetSitDestinationFinalAddressOk

`func (o *MTOServiceItemDestSIT) GetSitDestinationFinalAddressOk() (*Address, bool)`

GetSitDestinationFinalAddressOk returns a tuple with the SitDestinationFinalAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitDestinationFinalAddress

`func (o *MTOServiceItemDestSIT) SetSitDestinationFinalAddress(v Address)`

SetSitDestinationFinalAddress sets SitDestinationFinalAddress field to given value.

### HasSitDestinationFinalAddress

`func (o *MTOServiceItemDestSIT) HasSitDestinationFinalAddress() bool`

HasSitDestinationFinalAddress returns a boolean if a field has been set.

### GetReason

`func (o *MTOServiceItemDestSIT) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MTOServiceItemDestSIT) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MTOServiceItemDestSIT) SetReason(v string)`

SetReason sets Reason field to given value.


### SetReasonNil

`func (o *MTOServiceItemDestSIT) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *MTOServiceItemDestSIT) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetSitAddressUpdates

`func (o *MTOServiceItemDestSIT) GetSitAddressUpdates() []SitAddressUpdate`

GetSitAddressUpdates returns the SitAddressUpdates field if non-nil, zero value otherwise.

### GetSitAddressUpdatesOk

`func (o *MTOServiceItemDestSIT) GetSitAddressUpdatesOk() (*[]SitAddressUpdate, bool)`

GetSitAddressUpdatesOk returns a tuple with the SitAddressUpdates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitAddressUpdates

`func (o *MTOServiceItemDestSIT) SetSitAddressUpdates(v []SitAddressUpdate)`

SetSitAddressUpdates sets SitAddressUpdates field to given value.

### HasSitAddressUpdates

`func (o *MTOServiceItemDestSIT) HasSitAddressUpdates() bool`

HasSitAddressUpdates returns a boolean if a field has been set.

### GetSitRequestedDelivery

`func (o *MTOServiceItemDestSIT) GetSitRequestedDelivery() string`

GetSitRequestedDelivery returns the SitRequestedDelivery field if non-nil, zero value otherwise.

### GetSitRequestedDeliveryOk

`func (o *MTOServiceItemDestSIT) GetSitRequestedDeliveryOk() (*string, bool)`

GetSitRequestedDeliveryOk returns a tuple with the SitRequestedDelivery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitRequestedDelivery

`func (o *MTOServiceItemDestSIT) SetSitRequestedDelivery(v string)`

SetSitRequestedDelivery sets SitRequestedDelivery field to given value.

### HasSitRequestedDelivery

`func (o *MTOServiceItemDestSIT) HasSitRequestedDelivery() bool`

HasSitRequestedDelivery returns a boolean if a field has been set.

### SetSitRequestedDeliveryNil

`func (o *MTOServiceItemDestSIT) SetSitRequestedDeliveryNil(b bool)`

 SetSitRequestedDeliveryNil sets the value for SitRequestedDelivery to be an explicit nil

### UnsetSitRequestedDelivery
`func (o *MTOServiceItemDestSIT) UnsetSitRequestedDelivery()`

UnsetSitRequestedDelivery ensures that no value is present for SitRequestedDelivery, not even an explicit nil
### GetSitCustomerContacted

`func (o *MTOServiceItemDestSIT) GetSitCustomerContacted() string`

GetSitCustomerContacted returns the SitCustomerContacted field if non-nil, zero value otherwise.

### GetSitCustomerContactedOk

`func (o *MTOServiceItemDestSIT) GetSitCustomerContactedOk() (*string, bool)`

GetSitCustomerContactedOk returns a tuple with the SitCustomerContacted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSitCustomerContacted

`func (o *MTOServiceItemDestSIT) SetSitCustomerContacted(v string)`

SetSitCustomerContacted sets SitCustomerContacted field to given value.

### HasSitCustomerContacted

`func (o *MTOServiceItemDestSIT) HasSitCustomerContacted() bool`

HasSitCustomerContacted returns a boolean if a field has been set.

### SetSitCustomerContactedNil

`func (o *MTOServiceItemDestSIT) SetSitCustomerContactedNil(b bool)`

 SetSitCustomerContactedNil sets the value for SitCustomerContacted to be an explicit nil

### UnsetSitCustomerContacted
`func (o *MTOServiceItemDestSIT) UnsetSitCustomerContacted()`

UnsetSitCustomerContacted ensures that no value is present for SitCustomerContacted, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


