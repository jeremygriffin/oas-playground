# UpdateMTOServiceItemV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the service item. Must match path. | [optional] 
**ModelType** | [**UpdateMTOServiceItemModelTypeV3V3**](UpdateMTOServiceItemModelTypeV3.md) |  | 

## Methods

### NewUpdateMTOServiceItemV3

`func NewUpdateMTOServiceItemV3(modelType UpdateMTOServiceItemModelTypeV3V3, ) *UpdateMTOServiceItemV3`

NewUpdateMTOServiceItemV3 instantiates a new UpdateMTOServiceItemV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMTOServiceItemV3WithDefaults

`func NewUpdateMTOServiceItemV3WithDefaults() *UpdateMTOServiceItemV3`

NewUpdateMTOServiceItemV3WithDefaults instantiates a new UpdateMTOServiceItemV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateMTOServiceItemV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateMTOServiceItemV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateMTOServiceItemV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateMTOServiceItemV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModelType

`func (o *UpdateMTOServiceItemV3) GetModelType() UpdateMTOServiceItemModelTypeV3V3`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *UpdateMTOServiceItemV3) GetModelTypeOk() (*UpdateMTOServiceItemModelTypeV3V3, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *UpdateMTOServiceItemV3) SetModelType(v UpdateMTOServiceItemModelTypeV3V3)`

SetModelType sets ModelType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


