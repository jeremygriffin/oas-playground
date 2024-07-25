# UpdateMTOServiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | ID of the service item. Must match path. | [optional] 
**ModelType** | [**UpdateMTOServiceItemModelType**](UpdateMTOServiceItemModelType.md) |  | 

## Methods

### NewUpdateMTOServiceItem

`func NewUpdateMTOServiceItem(modelType UpdateMTOServiceItemModelType, ) *UpdateMTOServiceItem`

NewUpdateMTOServiceItem instantiates a new UpdateMTOServiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateMTOServiceItemWithDefaults

`func NewUpdateMTOServiceItemWithDefaults() *UpdateMTOServiceItem`

NewUpdateMTOServiceItemWithDefaults instantiates a new UpdateMTOServiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateMTOServiceItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateMTOServiceItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateMTOServiceItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateMTOServiceItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetModelType

`func (o *UpdateMTOServiceItem) GetModelType() UpdateMTOServiceItemModelType`

GetModelType returns the ModelType field if non-nil, zero value otherwise.

### GetModelTypeOk

`func (o *UpdateMTOServiceItem) GetModelTypeOk() (*UpdateMTOServiceItemModelType, bool)`

GetModelTypeOk returns a tuple with the ModelType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModelType

`func (o *UpdateMTOServiceItem) SetModelType(v UpdateMTOServiceItemModelType)`

SetModelType sets ModelType field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


