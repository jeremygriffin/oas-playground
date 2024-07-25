# UpdateReweigh

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Weight** | Pointer to **NullableInt32** | The total reweighed weight for the shipment in pounds. | [optional] 
**VerificationReason** | Pointer to **NullableString** | In lieu of a document being uploaded indicating why a reweigh did not occur. | [optional] 

## Methods

### NewUpdateReweigh

`func NewUpdateReweigh() *UpdateReweigh`

NewUpdateReweigh instantiates a new UpdateReweigh object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateReweighWithDefaults

`func NewUpdateReweighWithDefaults() *UpdateReweigh`

NewUpdateReweighWithDefaults instantiates a new UpdateReweigh object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWeight

`func (o *UpdateReweigh) GetWeight() int32`

GetWeight returns the Weight field if non-nil, zero value otherwise.

### GetWeightOk

`func (o *UpdateReweigh) GetWeightOk() (*int32, bool)`

GetWeightOk returns a tuple with the Weight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeight

`func (o *UpdateReweigh) SetWeight(v int32)`

SetWeight sets Weight field to given value.

### HasWeight

`func (o *UpdateReweigh) HasWeight() bool`

HasWeight returns a boolean if a field has been set.

### SetWeightNil

`func (o *UpdateReweigh) SetWeightNil(b bool)`

 SetWeightNil sets the value for Weight to be an explicit nil

### UnsetWeight
`func (o *UpdateReweigh) UnsetWeight()`

UnsetWeight ensures that no value is present for Weight, not even an explicit nil
### GetVerificationReason

`func (o *UpdateReweigh) GetVerificationReason() string`

GetVerificationReason returns the VerificationReason field if non-nil, zero value otherwise.

### GetVerificationReasonOk

`func (o *UpdateReweigh) GetVerificationReasonOk() (*string, bool)`

GetVerificationReasonOk returns a tuple with the VerificationReason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVerificationReason

`func (o *UpdateReweigh) SetVerificationReason(v string)`

SetVerificationReason sets VerificationReason field to given value.

### HasVerificationReason

`func (o *UpdateReweigh) HasVerificationReason() bool`

HasVerificationReason returns a boolean if a field has been set.

### SetVerificationReasonNil

`func (o *UpdateReweigh) SetVerificationReasonNil(b bool)`

 SetVerificationReasonNil sets the value for VerificationReason to be an explicit nil

### UnsetVerificationReason
`func (o *UpdateReweigh) UnsetVerificationReason()`

UnsetVerificationReason ensures that no value is present for VerificationReason, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


