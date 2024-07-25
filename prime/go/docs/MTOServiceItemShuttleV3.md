# MTOServiceItemShuttleV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReServiceCode** | **string** | A unique code for the service item. Indicates if shuttling is requested for the shipment origin (&#x60;DOSHUT&#x60;) or destination (&#x60;DDSHUT&#x60;).  | 
**Reason** | **string** | The contractor&#39;s explanation for why a shuttle service is requested. Used by the TOO while deciding to approve or reject the service item.  | 
**EstimatedWeight** | Pointer to **NullableInt32** | An estimate of how much weight from a shipment will be included in the shuttling service. | [optional] 
**ActualWeight** | Pointer to **NullableInt32** | A record of the actual weight that was shuttled. Provided by the movers, based on weight tickets. | [optional] 

## Methods

### NewMTOServiceItemShuttleV3

`func NewMTOServiceItemShuttleV3(reServiceCode string, reason string, ) *MTOServiceItemShuttleV3`

NewMTOServiceItemShuttleV3 instantiates a new MTOServiceItemShuttleV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTOServiceItemShuttleV3WithDefaults

`func NewMTOServiceItemShuttleV3WithDefaults() *MTOServiceItemShuttleV3`

NewMTOServiceItemShuttleV3WithDefaults instantiates a new MTOServiceItemShuttleV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReServiceCode

`func (o *MTOServiceItemShuttleV3) GetReServiceCode() string`

GetReServiceCode returns the ReServiceCode field if non-nil, zero value otherwise.

### GetReServiceCodeOk

`func (o *MTOServiceItemShuttleV3) GetReServiceCodeOk() (*string, bool)`

GetReServiceCodeOk returns a tuple with the ReServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReServiceCode

`func (o *MTOServiceItemShuttleV3) SetReServiceCode(v string)`

SetReServiceCode sets ReServiceCode field to given value.


### GetReason

`func (o *MTOServiceItemShuttleV3) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MTOServiceItemShuttleV3) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MTOServiceItemShuttleV3) SetReason(v string)`

SetReason sets Reason field to given value.


### GetEstimatedWeight

`func (o *MTOServiceItemShuttleV3) GetEstimatedWeight() int32`

GetEstimatedWeight returns the EstimatedWeight field if non-nil, zero value otherwise.

### GetEstimatedWeightOk

`func (o *MTOServiceItemShuttleV3) GetEstimatedWeightOk() (*int32, bool)`

GetEstimatedWeightOk returns a tuple with the EstimatedWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedWeight

`func (o *MTOServiceItemShuttleV3) SetEstimatedWeight(v int32)`

SetEstimatedWeight sets EstimatedWeight field to given value.

### HasEstimatedWeight

`func (o *MTOServiceItemShuttleV3) HasEstimatedWeight() bool`

HasEstimatedWeight returns a boolean if a field has been set.

### SetEstimatedWeightNil

`func (o *MTOServiceItemShuttleV3) SetEstimatedWeightNil(b bool)`

 SetEstimatedWeightNil sets the value for EstimatedWeight to be an explicit nil

### UnsetEstimatedWeight
`func (o *MTOServiceItemShuttleV3) UnsetEstimatedWeight()`

UnsetEstimatedWeight ensures that no value is present for EstimatedWeight, not even an explicit nil
### GetActualWeight

`func (o *MTOServiceItemShuttleV3) GetActualWeight() int32`

GetActualWeight returns the ActualWeight field if non-nil, zero value otherwise.

### GetActualWeightOk

`func (o *MTOServiceItemShuttleV3) GetActualWeightOk() (*int32, bool)`

GetActualWeightOk returns a tuple with the ActualWeight field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActualWeight

`func (o *MTOServiceItemShuttleV3) SetActualWeight(v int32)`

SetActualWeight sets ActualWeight field to given value.

### HasActualWeight

`func (o *MTOServiceItemShuttleV3) HasActualWeight() bool`

HasActualWeight returns a boolean if a field has been set.

### SetActualWeightNil

`func (o *MTOServiceItemShuttleV3) SetActualWeightNil(b bool)`

 SetActualWeightNil sets the value for ActualWeight to be an explicit nil

### UnsetActualWeight
`func (o *MTOServiceItemShuttleV3) UnsetActualWeight()`

UnsetActualWeight ensures that no value is present for ActualWeight, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


