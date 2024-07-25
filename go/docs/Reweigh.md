# Reweigh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**RequestedBy** | Pointer to [**ReweighRequester**](ReweighRequester.md) |  | [optional] 
**VerificationProvidedAt** | Pointer to **NullableTime** |  | [optional] 
**VerificationReason** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **NullableInt32** |  | [optional] 
**ShipmentID** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewReweigh

`func NewReweigh() *Reweigh`

NewReweigh instantiates a new Reweigh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReweighWithDefaults

`func NewReweighWithDefaults() *Reweigh`

NewReweighWithDefaults instantiates a new Reweigh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Reweigh) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Reweigh) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Reweigh) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Reweigh) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequestedAt

`func (o *Reweigh) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *Reweigh) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *Reweigh) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *Reweigh) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetRequestedBy

`func (o *Reweigh) GetRequestedBy() ReweighRequester`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *Reweigh) GetRequestedByOk() (*ReweighRequester, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *Reweigh) SetRequestedBy(v ReweighRequester)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *Reweigh) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetVerificationProvidedAt

`func (o *Reweigh) GetVerificationProvidedAt() time.Time`

GetVerificationProvidedAt returns the VerificationProvidedAt field if non-nil, zero value otherwise.

### GetVerificationProvidedAtOk

`func (o *Reweigh) GetVerificationProvidedAtOk() (*time.Time, bool)`

GetVerificationProvidedAtOk returns a tuple with the VerificationProvidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationProvidedAt

`func (o *Reweigh) SetVerificationProvidedAt(v time.Time)`

SetVerificationProvidedAt sets VerificationProvidedAt field to given value.

### HasVerificationProvidedAt

`func (o *Reweigh) HasVerificationProvidedAt() bool`

HasVerificationProvidedAt returns a boolean if a field has been set.

### SetVerificationProvidedAtNil

`func (o *Reweigh) SetVerificationProvidedAtNil(b bool)`

 SetVerificationProvidedAtNil sets the value for VerificationProvidedAt to be an explicit nil

### UnsetVerificationProvidedAt
`func (o *Reweigh) UnsetVerificationProvidedAt()`

UnsetVerificationProvidedAt ensures that no value is present for VerificationProvidedAt, not even an explicit nil
### GetVerificationReason

`func (o *Reweigh) GetVerificationReason() string`

GetVerificationReason returns the VerificationReason field if non-nil, zero value otherwise.

### GetVerificationReasonOk

`func (o *Reweigh) GetVerificationReasonOk() (*string, bool)`

GetVerificationReasonOk returns a tuple with the VerificationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationReason

`func (o *Reweigh) SetVerificationReason(v string)`

SetVerificationReason sets VerificationReason field to given value.

### HasVerificationReason

`func (o *Reweigh) HasVerificationReason() bool`

HasVerificationReason returns a boolean if a field has been set.

### SetVerificationReasonNil

`func (o *Reweigh) SetVerificationReasonNil(b bool)`

 SetVerificationReasonNil sets the value for VerificationReason to be an explicit nil

### UnsetVerificationReason
`func (o *Reweigh) UnsetVerificationReason()`

UnsetVerificationReason ensures that no value is present for VerificationReason, not even an explicit nil
### GetWeight

`func (o *Reweigh) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *Reweigh) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *Reweigh) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *Reweigh) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *Reweigh) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *Reweigh) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetShipmentID

`func (o *Reweigh) GetShipmentID() string`

GetShipmentID returns the ShipmentID field if non-nil, zero value otherwise.

### GetShipmentIDOk

`func (o *Reweigh) GetShipmentIDOk() (*string, bool)`

GetShipmentIDOk returns a tuple with the ShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentID

`func (o *Reweigh) SetShipmentID(v string)`

SetShipmentID sets ShipmentID field to given value.

### HasShipmentID

`func (o *Reweigh) HasShipmentID() bool`

HasShipmentID returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Reweigh) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Reweigh) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Reweigh) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Reweigh) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *Reweigh) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *Reweigh) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *Reweigh) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *Reweigh) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetETag

`func (o *Reweigh) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *Reweigh) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *Reweigh) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *Reweigh) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


