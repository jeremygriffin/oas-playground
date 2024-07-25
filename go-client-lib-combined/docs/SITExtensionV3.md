# SITExtensionV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**MtoShipmentID** | Pointer to **string** |  | [optional] 
**RequestReason** | Pointer to **string** |  | [optional] 
**ContractorRemarks** | Pointer to **NullableString** |  | [optional] 
**RequestedDays** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ApprovedDays** | Pointer to **NullableInt32** |  | [optional] 
**DecisionDate** | Pointer to **NullableTime** |  | [optional] 
**OfficeRemarks** | Pointer to **NullableString** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewSITExtensionV3

`func NewSITExtensionV3() *SITExtensionV3`

NewSITExtensionV3 instantiates a new SITExtensionV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSITExtensionV3WithDefaults

`func NewSITExtensionV3WithDefaults() *SITExtensionV3`

NewSITExtensionV3WithDefaults instantiates a new SITExtensionV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SITExtensionV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SITExtensionV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SITExtensionV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SITExtensionV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMtoShipmentID

`func (o *SITExtensionV3) GetMtoShipmentID() string`

GetMtoShipmentID returns the MtoShipmentID field if non-nil, zero value otherwise.

### GetMtoShipmentIDOk

`func (o *SITExtensionV3) GetMtoShipmentIDOk() (*string, bool)`

GetMtoShipmentIDOk returns a tuple with the MtoShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoShipmentID

`func (o *SITExtensionV3) SetMtoShipmentID(v string)`

SetMtoShipmentID sets MtoShipmentID field to given value.

### HasMtoShipmentID

`func (o *SITExtensionV3) HasMtoShipmentID() bool`

HasMtoShipmentID returns a boolean if a field has been set.

### GetRequestReason

`func (o *SITExtensionV3) GetRequestReason() string`

GetRequestReason returns the RequestReason field if non-nil, zero value otherwise.

### GetRequestReasonOk

`func (o *SITExtensionV3) GetRequestReasonOk() (*string, bool)`

GetRequestReasonOk returns a tuple with the RequestReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestReason

`func (o *SITExtensionV3) SetRequestReason(v string)`

SetRequestReason sets RequestReason field to given value.

### HasRequestReason

`func (o *SITExtensionV3) HasRequestReason() bool`

HasRequestReason returns a boolean if a field has been set.

### GetContractorRemarks

`func (o *SITExtensionV3) GetContractorRemarks() string`

GetContractorRemarks returns the ContractorRemarks field if non-nil, zero value otherwise.

### GetContractorRemarksOk

`func (o *SITExtensionV3) GetContractorRemarksOk() (*string, bool)`

GetContractorRemarksOk returns a tuple with the ContractorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorRemarks

`func (o *SITExtensionV3) SetContractorRemarks(v string)`

SetContractorRemarks sets ContractorRemarks field to given value.

### HasContractorRemarks

`func (o *SITExtensionV3) HasContractorRemarks() bool`

HasContractorRemarks returns a boolean if a field has been set.

### SetContractorRemarksNil

`func (o *SITExtensionV3) SetContractorRemarksNil(b bool)`

 SetContractorRemarksNil sets the value for ContractorRemarks to be an explicit nil

### UnsetContractorRemarks
`func (o *SITExtensionV3) UnsetContractorRemarks()`

UnsetContractorRemarks ensures that no value is present for ContractorRemarks, not even an explicit nil
### GetRequestedDays

`func (o *SITExtensionV3) GetRequestedDays() int32`

GetRequestedDays returns the RequestedDays field if non-nil, zero value otherwise.

### GetRequestedDaysOk

`func (o *SITExtensionV3) GetRequestedDaysOk() (*int32, bool)`

GetRequestedDaysOk returns a tuple with the RequestedDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDays

`func (o *SITExtensionV3) SetRequestedDays(v int32)`

SetRequestedDays sets RequestedDays field to given value.

### HasRequestedDays

`func (o *SITExtensionV3) HasRequestedDays() bool`

HasRequestedDays returns a boolean if a field has been set.

### GetStatus

`func (o *SITExtensionV3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SITExtensionV3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SITExtensionV3) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SITExtensionV3) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApprovedDays

`func (o *SITExtensionV3) GetApprovedDays() int32`

GetApprovedDays returns the ApprovedDays field if non-nil, zero value otherwise.

### GetApprovedDaysOk

`func (o *SITExtensionV3) GetApprovedDaysOk() (*int32, bool)`

GetApprovedDaysOk returns a tuple with the ApprovedDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDays

`func (o *SITExtensionV3) SetApprovedDays(v int32)`

SetApprovedDays sets ApprovedDays field to given value.

### HasApprovedDays

`func (o *SITExtensionV3) HasApprovedDays() bool`

HasApprovedDays returns a boolean if a field has been set.

### SetApprovedDaysNil

`func (o *SITExtensionV3) SetApprovedDaysNil(b bool)`

 SetApprovedDaysNil sets the value for ApprovedDays to be an explicit nil

### UnsetApprovedDays
`func (o *SITExtensionV3) UnsetApprovedDays()`

UnsetApprovedDays ensures that no value is present for ApprovedDays, not even an explicit nil
### GetDecisionDate

`func (o *SITExtensionV3) GetDecisionDate() time.Time`

GetDecisionDate returns the DecisionDate field if non-nil, zero value otherwise.

### GetDecisionDateOk

`func (o *SITExtensionV3) GetDecisionDateOk() (*time.Time, bool)`

GetDecisionDateOk returns a tuple with the DecisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDate

`func (o *SITExtensionV3) SetDecisionDate(v time.Time)`

SetDecisionDate sets DecisionDate field to given value.

### HasDecisionDate

`func (o *SITExtensionV3) HasDecisionDate() bool`

HasDecisionDate returns a boolean if a field has been set.

### SetDecisionDateNil

`func (o *SITExtensionV3) SetDecisionDateNil(b bool)`

 SetDecisionDateNil sets the value for DecisionDate to be an explicit nil

### UnsetDecisionDate
`func (o *SITExtensionV3) UnsetDecisionDate()`

UnsetDecisionDate ensures that no value is present for DecisionDate, not even an explicit nil
### GetOfficeRemarks

`func (o *SITExtensionV3) GetOfficeRemarks() string`

GetOfficeRemarks returns the OfficeRemarks field if non-nil, zero value otherwise.

### GetOfficeRemarksOk

`func (o *SITExtensionV3) GetOfficeRemarksOk() (*string, bool)`

GetOfficeRemarksOk returns a tuple with the OfficeRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeRemarks

`func (o *SITExtensionV3) SetOfficeRemarks(v string)`

SetOfficeRemarks sets OfficeRemarks field to given value.

### HasOfficeRemarks

`func (o *SITExtensionV3) HasOfficeRemarks() bool`

HasOfficeRemarks returns a boolean if a field has been set.

### SetOfficeRemarksNil

`func (o *SITExtensionV3) SetOfficeRemarksNil(b bool)`

 SetOfficeRemarksNil sets the value for OfficeRemarks to be an explicit nil

### UnsetOfficeRemarks
`func (o *SITExtensionV3) UnsetOfficeRemarks()`

UnsetOfficeRemarks ensures that no value is present for OfficeRemarks, not even an explicit nil
### GetCreatedAt

`func (o *SITExtensionV3) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SITExtensionV3) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SITExtensionV3) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SITExtensionV3) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SITExtensionV3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SITExtensionV3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SITExtensionV3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SITExtensionV3) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetETag

`func (o *SITExtensionV3) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *SITExtensionV3) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *SITExtensionV3) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *SITExtensionV3) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


