# UpdateMTOServiceItemShuttleV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActualWeight** | Pointer to **NullableInt32** | Provided by the movers, based on weight tickets. Relevant for shuttling (DDSHUT &amp; DOSHUT) service items. | [optional] 
**EstimatedWeight** | Pointer to **NullableInt32** | An estimate of how much weight from a shipment will be included in a shuttling (DDSHUT &amp; DOSHUT) service item. | [optional] 
**ReServiceCode** | Pointer to **string** | Service code allowed for this model type. | [optional] 

## Methods

### NewUpdateMTOServiceItemShuttleV3

`func NewUpdateMTOServiceItemShuttleV3() *UpdateMTOServiceItemShuttleV3`

NewUpdateMTOServiceItemShuttleV3 instantiates a new UpdateMTOServiceItemShuttleV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMTOServiceItemShuttleV3WithDefaults

`func NewUpdateMTOServiceItemShuttleV3WithDefaults() *UpdateMTOServiceItemShuttleV3`

NewUpdateMTOServiceItemShuttleV3WithDefaults instantiates a new UpdateMTOServiceItemShuttleV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActualWeight

`func (o *UpdateMTOServiceItemShuttleV3) GetActualWeight() int32`

GetActualWeight returns the ActualWeight field if non-nil, zero value otherwise.

### GetActualWeightOk

`func (o *UpdateMTOServiceItemShuttleV3) GetActualWeightOk() (*int32, bool)`

GetActualWeightOk returns a tuple with the ActualWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualWeight

`func (o *UpdateMTOServiceItemShuttleV3) SetActualWeight(v int32)`

SetActualWeight sets ActualWeight field to given value.

### HasActualWeight

`func (o *UpdateMTOServiceItemShuttleV3) HasActualWeight() bool`

HasActualWeight returns a boolean if a field has been set.

### SetActualWeightNil

`func (o *UpdateMTOServiceItemShuttleV3) SetActualWeightNil(b bool)`

 SetActualWeightNil sets the value for ActualWeight to be an explicit nil

### UnsetActualWeight
`func (o *UpdateMTOServiceItemShuttleV3) UnsetActualWeight()`

UnsetActualWeight ensures that no value is present for ActualWeight, not even an explicit nil
### GetEstimatedWeight

`func (o *UpdateMTOServiceItemShuttleV3) GetEstimatedWeight() int32`

GetEstimatedWeight returns the EstimatedWeight field if non-nil, zero value otherwise.

### GetEstimatedWeightOk

`func (o *UpdateMTOServiceItemShuttleV3) GetEstimatedWeightOk() (*int32, bool)`

GetEstimatedWeightOk returns a tuple with the EstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedWeight

`func (o *UpdateMTOServiceItemShuttleV3) SetEstimatedWeight(v int32)`

SetEstimatedWeight sets EstimatedWeight field to given value.

### HasEstimatedWeight

`func (o *UpdateMTOServiceItemShuttleV3) HasEstimatedWeight() bool`

HasEstimatedWeight returns a boolean if a field has been set.

### SetEstimatedWeightNil

`func (o *UpdateMTOServiceItemShuttleV3) SetEstimatedWeightNil(b bool)`

 SetEstimatedWeightNil sets the value for EstimatedWeight to be an explicit nil

### UnsetEstimatedWeight
`func (o *UpdateMTOServiceItemShuttleV3) UnsetEstimatedWeight()`

UnsetEstimatedWeight ensures that no value is present for EstimatedWeight, not even an explicit nil
### GetReServiceCode

`func (o *UpdateMTOServiceItemShuttleV3) GetReServiceCode() string`

GetReServiceCode returns the ReServiceCode field if non-nil, zero value otherwise.

### GetReServiceCodeOk

`func (o *UpdateMTOServiceItemShuttleV3) GetReServiceCodeOk() (*string, bool)`

GetReServiceCodeOk returns a tuple with the ReServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReServiceCode

`func (o *UpdateMTOServiceItemShuttleV3) SetReServiceCode(v string)`

SetReServiceCode sets ReServiceCode field to given value.

### HasReServiceCode

`func (o *UpdateMTOServiceItemShuttleV3) HasReServiceCode() bool`

HasReServiceCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


