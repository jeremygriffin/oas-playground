# MTOServiceItemDomesticCratingV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ReServiceCode** | **string** | A unique code for the service item. Indicates if the service is for crating (DCRT) or uncrating (DUCRT). | 
**Item** | **map[string]interface{}** | The dimensions of the item being crated. | 
**Crate** | **map[string]interface{}** | The dimensions for the crate the item will be shipped in. | 
**Description** | **string** | A description of the item being crated. | 
**Reason** | Pointer to **NullableString** | The contractor&#39;s explanation for why an item needed to be crated or uncrated. Used by the TOO while deciding to approve or reject the service item.  | [optional] 
**StandaloneCrate** | Pointer to **NullableBool** |  | [optional] 

## Methods

### NewMTOServiceItemDomesticCratingV3

`func NewMTOServiceItemDomesticCratingV3(reServiceCode string, item map[string]interface{}, crate map[string]interface{}, description string, ) *MTOServiceItemDomesticCratingV3`

NewMTOServiceItemDomesticCratingV3 instantiates a new MTOServiceItemDomesticCratingV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTOServiceItemDomesticCratingV3WithDefaults

`func NewMTOServiceItemDomesticCratingV3WithDefaults() *MTOServiceItemDomesticCratingV3`

NewMTOServiceItemDomesticCratingV3WithDefaults instantiates a new MTOServiceItemDomesticCratingV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReServiceCode

`func (o *MTOServiceItemDomesticCratingV3) GetReServiceCode() string`

GetReServiceCode returns the ReServiceCode field if non-nil, zero value otherwise.

### GetReServiceCodeOk

`func (o *MTOServiceItemDomesticCratingV3) GetReServiceCodeOk() (*string, bool)`

GetReServiceCodeOk returns a tuple with the ReServiceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReServiceCode

`func (o *MTOServiceItemDomesticCratingV3) SetReServiceCode(v string)`

SetReServiceCode sets ReServiceCode field to given value.


### GetItem

`func (o *MTOServiceItemDomesticCratingV3) GetItem() map[string]interface{}`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *MTOServiceItemDomesticCratingV3) GetItemOk() (*map[string]interface{}, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *MTOServiceItemDomesticCratingV3) SetItem(v map[string]interface{})`

SetItem sets Item field to given value.


### GetCrate

`func (o *MTOServiceItemDomesticCratingV3) GetCrate() map[string]interface{}`

GetCrate returns the Crate field if non-nil, zero value otherwise.

### GetCrateOk

`func (o *MTOServiceItemDomesticCratingV3) GetCrateOk() (*map[string]interface{}, bool)`

GetCrateOk returns a tuple with the Crate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrate

`func (o *MTOServiceItemDomesticCratingV3) SetCrate(v map[string]interface{})`

SetCrate sets Crate field to given value.


### GetDescription

`func (o *MTOServiceItemDomesticCratingV3) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MTOServiceItemDomesticCratingV3) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MTOServiceItemDomesticCratingV3) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetReason

`func (o *MTOServiceItemDomesticCratingV3) GetReason() string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *MTOServiceItemDomesticCratingV3) GetReasonOk() (*string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *MTOServiceItemDomesticCratingV3) SetReason(v string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *MTOServiceItemDomesticCratingV3) HasReason() bool`

HasReason returns a boolean if a field has been set.

### SetReasonNil

`func (o *MTOServiceItemDomesticCratingV3) SetReasonNil(b bool)`

 SetReasonNil sets the value for Reason to be an explicit nil

### UnsetReason
`func (o *MTOServiceItemDomesticCratingV3) UnsetReason()`

UnsetReason ensures that no value is present for Reason, not even an explicit nil
### GetStandaloneCrate

`func (o *MTOServiceItemDomesticCratingV3) GetStandaloneCrate() bool`

GetStandaloneCrate returns the StandaloneCrate field if non-nil, zero value otherwise.

### GetStandaloneCrateOk

`func (o *MTOServiceItemDomesticCratingV3) GetStandaloneCrateOk() (*bool, bool)`

GetStandaloneCrateOk returns a tuple with the StandaloneCrate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStandaloneCrate

`func (o *MTOServiceItemDomesticCratingV3) SetStandaloneCrate(v bool)`

SetStandaloneCrate sets StandaloneCrate field to given value.

### HasStandaloneCrate

`func (o *MTOServiceItemDomesticCratingV3) HasStandaloneCrate() bool`

HasStandaloneCrate returns a boolean if a field has been set.

### SetStandaloneCrateNil

`func (o *MTOServiceItemDomesticCratingV3) SetStandaloneCrateNil(b bool)`

 SetStandaloneCrateNil sets the value for StandaloneCrate to be an explicit nil

### UnsetStandaloneCrate
`func (o *MTOServiceItemDomesticCratingV3) UnsetStandaloneCrate()`

UnsetStandaloneCrate ensures that no value is present for StandaloneCrate, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


