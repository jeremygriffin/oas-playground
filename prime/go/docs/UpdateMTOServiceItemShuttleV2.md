# UpdateMTOServiceItemShuttleV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualWeight** | Pointer to **NullableInt32** | Provided by the movers, based on weight tickets. Relevant for shuttling (DDSHUT &amp; DOSHUT) service items. | [optional] 
**EstimatedWeight** | Pointer to **NullableInt32** | An estimate of how much weight from a shipment will be included in a shuttling (DDSHUT &amp; DOSHUT) service item. | [optional] 
**ReServiceCode** | Pointer to **string** | Service code allowed for this model type. | [optional] 

## Methods

### NewUpdateMTOServiceItemShuttleV2

`func NewUpdateMTOServiceItemShuttleV2() *UpdateMTOServiceItemShuttleV2`

NewUpdateMTOServiceItemShuttleV2 instantiates a new UpdateMTOServiceItemShuttleV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMTOServiceItemShuttleV2WithDefaults

`func NewUpdateMTOServiceItemShuttleV2WithDefaults() *UpdateMTOServiceItemShuttleV2`

NewUpdateMTOServiceItemShuttleV2WithDefaults instantiates a new UpdateMTOServiceItemShuttleV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualWeight

`func (o *UpdateMTOServiceItemShuttleV2) GetActualWeight() int32`

GetActualWeight returns the ActualWeight field if non-nil, zero value otherwise.

### GetActualWeightOk

`func (o *UpdateMTOServiceItemShuttleV2) GetActualWeightOk() (*int32, bool)`

GetActualWeightOk returns a tuple with the ActualWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualWeight

`func (o *UpdateMTOServiceItemShuttleV2) SetActualWeight(v int32)`

SetActualWeight sets ActualWeight field to given value.

### HasActualWeight

`func (o *UpdateMTOServiceItemShuttleV2) HasActualWeight() bool`

HasActualWeight returns a boolean if a field has been set.

### SetActualWeightNil

`func (o *UpdateMTOServiceItemShuttleV2) SetActualWeightNil(b bool)`

 SetActualWeightNil sets the value for ActualWeight to be an explicit nil

### UnsetActualWeight
`func (o *UpdateMTOServiceItemShuttleV2) UnsetActualWeight()`

UnsetActualWeight ensures that no value is present for ActualWeight, not even an explicit nil
### GetEstimatedWeight

`func (o *UpdateMTOServiceItemShuttleV2) GetEstimatedWeight() int32`

GetEstimatedWeight returns the EstimatedWeight field if non-nil, zero value otherwise.

### GetEstimatedWeightOk

`func (o *UpdateMTOServiceItemShuttleV2) GetEstimatedWeightOk() (*int32, bool)`

GetEstimatedWeightOk returns a tuple with the EstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedWeight

`func (o *UpdateMTOServiceItemShuttleV2) SetEstimatedWeight(v int32)`

SetEstimatedWeight sets EstimatedWeight field to given value.

### HasEstimatedWeight

`func (o *UpdateMTOServiceItemShuttleV2) HasEstimatedWeight() bool`

HasEstimatedWeight returns a boolean if a field has been set.

### SetEstimatedWeightNil

`func (o *UpdateMTOServiceItemShuttleV2) SetEstimatedWeightNil(b bool)`

 SetEstimatedWeightNil sets the value for EstimatedWeight to be an explicit nil

### UnsetEstimatedWeight
`func (o *UpdateMTOServiceItemShuttleV2) UnsetEstimatedWeight()`

UnsetEstimatedWeight ensures that no value is present for EstimatedWeight, not even an explicit nil
### GetReServiceCode

`func (o *UpdateMTOServiceItemShuttleV2) GetReServiceCode() string`

GetReServiceCode returns the ReServiceCode field if non-nil, zero value otherwise.

### GetReServiceCodeOk

`func (o *UpdateMTOServiceItemShuttleV2) GetReServiceCodeOk() (*string, bool)`

GetReServiceCodeOk returns a tuple with the ReServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReServiceCode

`func (o *UpdateMTOServiceItemShuttleV2) SetReServiceCode(v string)`

SetReServiceCode sets ReServiceCode field to given value.

### HasReServiceCode

`func (o *UpdateMTOServiceItemShuttleV2) HasReServiceCode() bool`

HasReServiceCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


