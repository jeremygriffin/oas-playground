# UpdateReweighV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weight** | Pointer to **NullableInt32** | The total reweighed weight for the shipment in pounds. | [optional] 
**VerificationReason** | Pointer to **NullableString** | In lieu of a document being uploaded indicating why a reweigh did not occur. | [optional] 

## Methods

### NewUpdateReweighV2

`func NewUpdateReweighV2() *UpdateReweighV2`

NewUpdateReweighV2 instantiates a new UpdateReweighV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateReweighV2WithDefaults

`func NewUpdateReweighV2WithDefaults() *UpdateReweighV2`

NewUpdateReweighV2WithDefaults instantiates a new UpdateReweighV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeight

`func (o *UpdateReweighV2) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *UpdateReweighV2) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *UpdateReweighV2) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *UpdateReweighV2) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *UpdateReweighV2) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *UpdateReweighV2) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetVerificationReason

`func (o *UpdateReweighV2) GetVerificationReason() string`

GetVerificationReason returns the VerificationReason field if non-nil, zero value otherwise.

### GetVerificationReasonOk

`func (o *UpdateReweighV2) GetVerificationReasonOk() (*string, bool)`

GetVerificationReasonOk returns a tuple with the VerificationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationReason

`func (o *UpdateReweighV2) SetVerificationReason(v string)`

SetVerificationReason sets VerificationReason field to given value.

### HasVerificationReason

`func (o *UpdateReweighV2) HasVerificationReason() bool`

HasVerificationReason returns a boolean if a field has been set.

### SetVerificationReasonNil

`func (o *UpdateReweighV2) SetVerificationReasonNil(b bool)`

 SetVerificationReasonNil sets the value for VerificationReason to be an explicit nil

### UnsetVerificationReason
`func (o *UpdateReweighV2) UnsetVerificationReason()`

UnsetVerificationReason ensures that no value is present for VerificationReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


