# ValidationError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Detail** | **string** |  | 
**Instance** | **string** |  | 
**InvalidFields** | **map[string][]string** |  | 

## Methods

### NewValidationError

`func NewValidationError(title string, detail string, instance string, invalidFields map[string][]string, ) *ValidationError`

NewValidationError instantiates a new ValidationError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorWithDefaults

`func NewValidationErrorWithDefaults() *ValidationError`

NewValidationErrorWithDefaults instantiates a new ValidationError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ValidationError) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ValidationError) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ValidationError) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *ValidationError) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ValidationError) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ValidationError) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetInstance

`func (o *ValidationError) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ValidationError) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ValidationError) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetInvalidFields

`func (o *ValidationError) GetInvalidFields() map[string][]string`

GetInvalidFields returns the InvalidFields field if non-nil, zero value otherwise.

### GetInvalidFieldsOk

`func (o *ValidationError) GetInvalidFieldsOk() (*map[string][]string, bool)`

GetInvalidFieldsOk returns a tuple with the InvalidFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidFields

`func (o *ValidationError) SetInvalidFields(v map[string][]string)`

SetInvalidFields sets InvalidFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


