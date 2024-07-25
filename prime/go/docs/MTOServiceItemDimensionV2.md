# MTOServiceItemDimensionV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Length** | **int32** | Length in thousandth inches. 1000 thou &#x3D; 1 inch. | 
**Width** | **int32** | Width in thousandth inches. 1000 thou &#x3D; 1 inch. | 
**Height** | **int32** | Height in thousandth inches. 1000 thou &#x3D; 1 inch. | 

## Methods

### NewMTOServiceItemDimensionV2

`func NewMTOServiceItemDimensionV2(length int32, width int32, height int32, ) *MTOServiceItemDimensionV2`

NewMTOServiceItemDimensionV2 instantiates a new MTOServiceItemDimensionV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTOServiceItemDimensionV2WithDefaults

`func NewMTOServiceItemDimensionV2WithDefaults() *MTOServiceItemDimensionV2`

NewMTOServiceItemDimensionV2WithDefaults instantiates a new MTOServiceItemDimensionV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MTOServiceItemDimensionV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MTOServiceItemDimensionV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MTOServiceItemDimensionV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MTOServiceItemDimensionV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLength

`func (o *MTOServiceItemDimensionV2) GetLength() int32`

GetLength returns the Length field if non-nil, zero value otherwise.

### GetLengthOk

`func (o *MTOServiceItemDimensionV2) GetLengthOk() (*int32, bool)`

GetLengthOk returns a tuple with the Length field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLength

`func (o *MTOServiceItemDimensionV2) SetLength(v int32)`

SetLength sets Length field to given value.


### GetWidth

`func (o *MTOServiceItemDimensionV2) GetWidth() int32`

GetWidth returns the Width field if non-nil, zero value otherwise.

### GetWidthOk

`func (o *MTOServiceItemDimensionV2) GetWidthOk() (*int32, bool)`

GetWidthOk returns a tuple with the Width field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWidth

`func (o *MTOServiceItemDimensionV2) SetWidth(v int32)`

SetWidth sets Width field to given value.


### GetHeight

`func (o *MTOServiceItemDimensionV2) GetHeight() int32`

GetHeight returns the Height field if non-nil, zero value otherwise.

### GetHeightOk

`func (o *MTOServiceItemDimensionV2) GetHeightOk() (*int32, bool)`

GetHeightOk returns a tuple with the Height field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeight

`func (o *MTOServiceItemDimensionV2) SetHeight(v int32)`

SetHeight sets Height field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


