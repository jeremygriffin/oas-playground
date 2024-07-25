# SITExtension

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

### NewSITExtension

`func NewSITExtension() *SITExtension`

NewSITExtension instantiates a new SITExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSITExtensionWithDefaults

`func NewSITExtensionWithDefaults() *SITExtension`

NewSITExtensionWithDefaults instantiates a new SITExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SITExtension) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SITExtension) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SITExtension) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SITExtension) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMtoShipmentID

`func (o *SITExtension) GetMtoShipmentID() string`

GetMtoShipmentID returns the MtoShipmentID field if non-nil, zero value otherwise.

### GetMtoShipmentIDOk

`func (o *SITExtension) GetMtoShipmentIDOk() (*string, bool)`

GetMtoShipmentIDOk returns a tuple with the MtoShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoShipmentID

`func (o *SITExtension) SetMtoShipmentID(v string)`

SetMtoShipmentID sets MtoShipmentID field to given value.

### HasMtoShipmentID

`func (o *SITExtension) HasMtoShipmentID() bool`

HasMtoShipmentID returns a boolean if a field has been set.

### GetRequestReason

`func (o *SITExtension) GetRequestReason() string`

GetRequestReason returns the RequestReason field if non-nil, zero value otherwise.

### GetRequestReasonOk

`func (o *SITExtension) GetRequestReasonOk() (*string, bool)`

GetRequestReasonOk returns a tuple with the RequestReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestReason

`func (o *SITExtension) SetRequestReason(v string)`

SetRequestReason sets RequestReason field to given value.

### HasRequestReason

`func (o *SITExtension) HasRequestReason() bool`

HasRequestReason returns a boolean if a field has been set.

### GetContractorRemarks

`func (o *SITExtension) GetContractorRemarks() string`

GetContractorRemarks returns the ContractorRemarks field if non-nil, zero value otherwise.

### GetContractorRemarksOk

`func (o *SITExtension) GetContractorRemarksOk() (*string, bool)`

GetContractorRemarksOk returns a tuple with the ContractorRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContractorRemarks

`func (o *SITExtension) SetContractorRemarks(v string)`

SetContractorRemarks sets ContractorRemarks field to given value.

### HasContractorRemarks

`func (o *SITExtension) HasContractorRemarks() bool`

HasContractorRemarks returns a boolean if a field has been set.

### SetContractorRemarksNil

`func (o *SITExtension) SetContractorRemarksNil(b bool)`

 SetContractorRemarksNil sets the value for ContractorRemarks to be an explicit nil

### UnsetContractorRemarks
`func (o *SITExtension) UnsetContractorRemarks()`

UnsetContractorRemarks ensures that no value is present for ContractorRemarks, not even an explicit nil
### GetRequestedDays

`func (o *SITExtension) GetRequestedDays() int32`

GetRequestedDays returns the RequestedDays field if non-nil, zero value otherwise.

### GetRequestedDaysOk

`func (o *SITExtension) GetRequestedDaysOk() (*int32, bool)`

GetRequestedDaysOk returns a tuple with the RequestedDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedDays

`func (o *SITExtension) SetRequestedDays(v int32)`

SetRequestedDays sets RequestedDays field to given value.

### HasRequestedDays

`func (o *SITExtension) HasRequestedDays() bool`

HasRequestedDays returns a boolean if a field has been set.

### GetStatus

`func (o *SITExtension) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SITExtension) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SITExtension) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SITExtension) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetApprovedDays

`func (o *SITExtension) GetApprovedDays() int32`

GetApprovedDays returns the ApprovedDays field if non-nil, zero value otherwise.

### GetApprovedDaysOk

`func (o *SITExtension) GetApprovedDaysOk() (*int32, bool)`

GetApprovedDaysOk returns a tuple with the ApprovedDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovedDays

`func (o *SITExtension) SetApprovedDays(v int32)`

SetApprovedDays sets ApprovedDays field to given value.

### HasApprovedDays

`func (o *SITExtension) HasApprovedDays() bool`

HasApprovedDays returns a boolean if a field has been set.

### SetApprovedDaysNil

`func (o *SITExtension) SetApprovedDaysNil(b bool)`

 SetApprovedDaysNil sets the value for ApprovedDays to be an explicit nil

### UnsetApprovedDays
`func (o *SITExtension) UnsetApprovedDays()`

UnsetApprovedDays ensures that no value is present for ApprovedDays, not even an explicit nil
### GetDecisionDate

`func (o *SITExtension) GetDecisionDate() time.Time`

GetDecisionDate returns the DecisionDate field if non-nil, zero value otherwise.

### GetDecisionDateOk

`func (o *SITExtension) GetDecisionDateOk() (*time.Time, bool)`

GetDecisionDateOk returns a tuple with the DecisionDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecisionDate

`func (o *SITExtension) SetDecisionDate(v time.Time)`

SetDecisionDate sets DecisionDate field to given value.

### HasDecisionDate

`func (o *SITExtension) HasDecisionDate() bool`

HasDecisionDate returns a boolean if a field has been set.

### SetDecisionDateNil

`func (o *SITExtension) SetDecisionDateNil(b bool)`

 SetDecisionDateNil sets the value for DecisionDate to be an explicit nil

### UnsetDecisionDate
`func (o *SITExtension) UnsetDecisionDate()`

UnsetDecisionDate ensures that no value is present for DecisionDate, not even an explicit nil
### GetOfficeRemarks

`func (o *SITExtension) GetOfficeRemarks() string`

GetOfficeRemarks returns the OfficeRemarks field if non-nil, zero value otherwise.

### GetOfficeRemarksOk

`func (o *SITExtension) GetOfficeRemarksOk() (*string, bool)`

GetOfficeRemarksOk returns a tuple with the OfficeRemarks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfficeRemarks

`func (o *SITExtension) SetOfficeRemarks(v string)`

SetOfficeRemarks sets OfficeRemarks field to given value.

### HasOfficeRemarks

`func (o *SITExtension) HasOfficeRemarks() bool`

HasOfficeRemarks returns a boolean if a field has been set.

### SetOfficeRemarksNil

`func (o *SITExtension) SetOfficeRemarksNil(b bool)`

 SetOfficeRemarksNil sets the value for OfficeRemarks to be an explicit nil

### UnsetOfficeRemarks
`func (o *SITExtension) UnsetOfficeRemarks()`

UnsetOfficeRemarks ensures that no value is present for OfficeRemarks, not even an explicit nil
### GetCreatedAt

`func (o *SITExtension) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *SITExtension) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *SITExtension) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *SITExtension) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *SITExtension) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *SITExtension) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *SITExtension) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *SITExtension) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetETag

`func (o *SITExtension) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *SITExtension) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *SITExtension) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *SITExtension) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


