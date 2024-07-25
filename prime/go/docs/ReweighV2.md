# ReweighV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**RequestedBy** | Pointer to [**ReweighRequesterV2V2**](ReweighRequesterV2.md) |  | [optional] 
**VerificationProvidedAt** | Pointer to **NullableTime** |  | [optional] 
**VerificationReason** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **NullableInt32** |  | [optional] 
**ShipmentID** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewReweighV2

`func NewReweighV2() *ReweighV2`

NewReweighV2 instantiates a new ReweighV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReweighV2WithDefaults

`func NewReweighV2WithDefaults() *ReweighV2`

NewReweighV2WithDefaults instantiates a new ReweighV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReweighV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReweighV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReweighV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReweighV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequestedAt

`func (o *ReweighV2) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *ReweighV2) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *ReweighV2) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *ReweighV2) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetRequestedBy

`func (o *ReweighV2) GetRequestedBy() ReweighRequesterV2V2`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *ReweighV2) GetRequestedByOk() (*ReweighRequesterV2V2, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *ReweighV2) SetRequestedBy(v ReweighRequesterV2V2)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *ReweighV2) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetVerificationProvidedAt

`func (o *ReweighV2) GetVerificationProvidedAt() time.Time`

GetVerificationProvidedAt returns the VerificationProvidedAt field if non-nil, zero value otherwise.

### GetVerificationProvidedAtOk

`func (o *ReweighV2) GetVerificationProvidedAtOk() (*time.Time, bool)`

GetVerificationProvidedAtOk returns a tuple with the VerificationProvidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationProvidedAt

`func (o *ReweighV2) SetVerificationProvidedAt(v time.Time)`

SetVerificationProvidedAt sets VerificationProvidedAt field to given value.

### HasVerificationProvidedAt

`func (o *ReweighV2) HasVerificationProvidedAt() bool`

HasVerificationProvidedAt returns a boolean if a field has been set.

### SetVerificationProvidedAtNil

`func (o *ReweighV2) SetVerificationProvidedAtNil(b bool)`

 SetVerificationProvidedAtNil sets the value for VerificationProvidedAt to be an explicit nil

### UnsetVerificationProvidedAt
`func (o *ReweighV2) UnsetVerificationProvidedAt()`

UnsetVerificationProvidedAt ensures that no value is present for VerificationProvidedAt, not even an explicit nil
### GetVerificationReason

`func (o *ReweighV2) GetVerificationReason() string`

GetVerificationReason returns the VerificationReason field if non-nil, zero value otherwise.

### GetVerificationReasonOk

`func (o *ReweighV2) GetVerificationReasonOk() (*string, bool)`

GetVerificationReasonOk returns a tuple with the VerificationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationReason

`func (o *ReweighV2) SetVerificationReason(v string)`

SetVerificationReason sets VerificationReason field to given value.

### HasVerificationReason

`func (o *ReweighV2) HasVerificationReason() bool`

HasVerificationReason returns a boolean if a field has been set.

### SetVerificationReasonNil

`func (o *ReweighV2) SetVerificationReasonNil(b bool)`

 SetVerificationReasonNil sets the value for VerificationReason to be an explicit nil

### UnsetVerificationReason
`func (o *ReweighV2) UnsetVerificationReason()`

UnsetVerificationReason ensures that no value is present for VerificationReason, not even an explicit nil
### GetWeight

`func (o *ReweighV2) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ReweighV2) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ReweighV2) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ReweighV2) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *ReweighV2) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *ReweighV2) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetShipmentID

`func (o *ReweighV2) GetShipmentID() string`

GetShipmentID returns the ShipmentID field if non-nil, zero value otherwise.

### GetShipmentIDOk

`func (o *ReweighV2) GetShipmentIDOk() (*string, bool)`

GetShipmentIDOk returns a tuple with the ShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentID

`func (o *ReweighV2) SetShipmentID(v string)`

SetShipmentID sets ShipmentID field to given value.

### HasShipmentID

`func (o *ReweighV2) HasShipmentID() bool`

HasShipmentID returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReweighV2) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReweighV2) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReweighV2) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReweighV2) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ReweighV2) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReweighV2) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReweighV2) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ReweighV2) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetETag

`func (o *ReweighV2) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *ReweighV2) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *ReweighV2) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *ReweighV2) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


