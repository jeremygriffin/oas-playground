# ClientError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Detail** | **string** |  | 
**Instance** | **string** |  | 

## Methods

### NewClientError

`func NewClientError(title string, detail string, instance string, ) *ClientError`

NewClientError instantiates a new ClientError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientErrorWithDefaults

`func NewClientErrorWithDefaults() *ClientError`

NewClientErrorWithDefaults instantiates a new ClientError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ClientError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ClientError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ClientError) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *ClientError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ClientError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ClientError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetInstance

`func (o *ClientError) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ClientError) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ClientError) SetInstance(v string)`

SetInstance sets Instance field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


