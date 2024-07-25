# ValidationErrorV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Title** | **string** |  | 
**Detail** | **string** |  | 
**Instance** | **string** |  | 
**InvalidFields** | **map[string][]string** |  | 

## Methods

### NewValidationErrorV2

`func NewValidationErrorV2(title string, detail string, instance string, invalidFields map[string][]string, ) *ValidationErrorV2`

NewValidationErrorV2 instantiates a new ValidationErrorV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewValidationErrorV2WithDefaults

`func NewValidationErrorV2WithDefaults() *ValidationErrorV2`

NewValidationErrorV2WithDefaults instantiates a new ValidationErrorV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTitle

`func (o *ValidationErrorV2) GetTitle() string`

GetTitle returns the Title field if non-nil, zero value otherwise.

### GetTitleOk

`func (o *ValidationErrorV2) GetTitleOk() (*string, bool)`

GetTitleOk returns a tuple with the Title field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTitle

`func (o *ValidationErrorV2) SetTitle(v string)`

SetTitle sets Title field to given value.


### GetDetail

`func (o *ValidationErrorV2) GetDetail() string`

GetDetail returns the Detail field if non-nil, zero value otherwise.

### GetDetailOk

`func (o *ValidationErrorV2) GetDetailOk() (*string, bool)`

GetDetailOk returns a tuple with the Detail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetail

`func (o *ValidationErrorV2) SetDetail(v string)`

SetDetail sets Detail field to given value.


### GetInstance

`func (o *ValidationErrorV2) GetInstance() string`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *ValidationErrorV2) GetInstanceOk() (*string, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *ValidationErrorV2) SetInstance(v string)`

SetInstance sets Instance field to given value.


### GetInvalidFields

`func (o *ValidationErrorV2) GetInvalidFields() map[string][]string`

GetInvalidFields returns the InvalidFields field if non-nil, zero value otherwise.

### GetInvalidFieldsOk

`func (o *ValidationErrorV2) GetInvalidFieldsOk() (*map[string][]string, bool)`

GetInvalidFieldsOk returns a tuple with the InvalidFields field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvalidFields

`func (o *ValidationErrorV2) SetInvalidFields(v map[string][]string)`

SetInvalidFields sets InvalidFields field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


