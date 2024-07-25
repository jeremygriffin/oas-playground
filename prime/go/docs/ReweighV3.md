# ReweighV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**RequestedAt** | Pointer to **time.Time** |  | [optional] 
**RequestedBy** | Pointer to [**ReweighRequesterV3V3**](ReweighRequesterV3.md) |  | [optional] 
**VerificationProvidedAt** | Pointer to **NullableTime** |  | [optional] 
**VerificationReason** | Pointer to **NullableString** |  | [optional] 
**Weight** | Pointer to **NullableInt32** |  | [optional] 
**ShipmentID** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewReweighV3

`func NewReweighV3() *ReweighV3`

NewReweighV3 instantiates a new ReweighV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewReweighV3WithDefaults

`func NewReweighV3WithDefaults() *ReweighV3`

NewReweighV3WithDefaults instantiates a new ReweighV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ReweighV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ReweighV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ReweighV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ReweighV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRequestedAt

`func (o *ReweighV3) GetRequestedAt() time.Time`

GetRequestedAt returns the RequestedAt field if non-nil, zero value otherwise.

### GetRequestedAtOk

`func (o *ReweighV3) GetRequestedAtOk() (*time.Time, bool)`

GetRequestedAtOk returns a tuple with the RequestedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedAt

`func (o *ReweighV3) SetRequestedAt(v time.Time)`

SetRequestedAt sets RequestedAt field to given value.

### HasRequestedAt

`func (o *ReweighV3) HasRequestedAt() bool`

HasRequestedAt returns a boolean if a field has been set.

### GetRequestedBy

`func (o *ReweighV3) GetRequestedBy() ReweighRequesterV3V3`

GetRequestedBy returns the RequestedBy field if non-nil, zero value otherwise.

### GetRequestedByOk

`func (o *ReweighV3) GetRequestedByOk() (*ReweighRequesterV3V3, bool)`

GetRequestedByOk returns a tuple with the RequestedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedBy

`func (o *ReweighV3) SetRequestedBy(v ReweighRequesterV3V3)`

SetRequestedBy sets RequestedBy field to given value.

### HasRequestedBy

`func (o *ReweighV3) HasRequestedBy() bool`

HasRequestedBy returns a boolean if a field has been set.

### GetVerificationProvidedAt

`func (o *ReweighV3) GetVerificationProvidedAt() time.Time`

GetVerificationProvidedAt returns the VerificationProvidedAt field if non-nil, zero value otherwise.

### GetVerificationProvidedAtOk

`func (o *ReweighV3) GetVerificationProvidedAtOk() (*time.Time, bool)`

GetVerificationProvidedAtOk returns a tuple with the VerificationProvidedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationProvidedAt

`func (o *ReweighV3) SetVerificationProvidedAt(v time.Time)`

SetVerificationProvidedAt sets VerificationProvidedAt field to given value.

### HasVerificationProvidedAt

`func (o *ReweighV3) HasVerificationProvidedAt() bool`

HasVerificationProvidedAt returns a boolean if a field has been set.

### SetVerificationProvidedAtNil

`func (o *ReweighV3) SetVerificationProvidedAtNil(b bool)`

 SetVerificationProvidedAtNil sets the value for VerificationProvidedAt to be an explicit nil

### UnsetVerificationProvidedAt
`func (o *ReweighV3) UnsetVerificationProvidedAt()`

UnsetVerificationProvidedAt ensures that no value is present for VerificationProvidedAt, not even an explicit nil
### GetVerificationReason

`func (o *ReweighV3) GetVerificationReason() string`

GetVerificationReason returns the VerificationReason field if non-nil, zero value otherwise.

### GetVerificationReasonOk

`func (o *ReweighV3) GetVerificationReasonOk() (*string, bool)`

GetVerificationReasonOk returns a tuple with the VerificationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationReason

`func (o *ReweighV3) SetVerificationReason(v string)`

SetVerificationReason sets VerificationReason field to given value.

### HasVerificationReason

`func (o *ReweighV3) HasVerificationReason() bool`

HasVerificationReason returns a boolean if a field has been set.

### SetVerificationReasonNil

`func (o *ReweighV3) SetVerificationReasonNil(b bool)`

 SetVerificationReasonNil sets the value for VerificationReason to be an explicit nil

### UnsetVerificationReason
`func (o *ReweighV3) UnsetVerificationReason()`

UnsetVerificationReason ensures that no value is present for VerificationReason, not even an explicit nil
### GetWeight

`func (o *ReweighV3) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *ReweighV3) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *ReweighV3) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *ReweighV3) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *ReweighV3) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *ReweighV3) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetShipmentID

`func (o *ReweighV3) GetShipmentID() string`

GetShipmentID returns the ShipmentID field if non-nil, zero value otherwise.

### GetShipmentIDOk

`func (o *ReweighV3) GetShipmentIDOk() (*string, bool)`

GetShipmentIDOk returns a tuple with the ShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShipmentID

`func (o *ReweighV3) SetShipmentID(v string)`

SetShipmentID sets ShipmentID field to given value.

### HasShipmentID

`func (o *ReweighV3) HasShipmentID() bool`

HasShipmentID returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ReweighV3) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ReweighV3) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ReweighV3) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ReweighV3) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ReweighV3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ReweighV3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ReweighV3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ReweighV3) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetETag

`func (o *ReweighV3) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *ReweighV3) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *ReweighV3) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *ReweighV3) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


