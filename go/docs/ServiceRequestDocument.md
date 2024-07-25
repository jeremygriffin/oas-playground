# ServiceRequestDocument

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uploads** | Pointer to [**[]UploadWithOmissions**](UploadWithOmissions.md) |  | [optional] 

## Methods

### NewServiceRequestDocument

`func NewServiceRequestDocument() *ServiceRequestDocument`

NewServiceRequestDocument instantiates a new ServiceRequestDocument object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceRequestDocumentWithDefaults

`func NewServiceRequestDocumentWithDefaults() *ServiceRequestDocument`

NewServiceRequestDocumentWithDefaults instantiates a new ServiceRequestDocument object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUploads

`func (o *ServiceRequestDocument) GetUploads() []UploadWithOmissions`

GetUploads returns the Uploads field if non-nil, zero value otherwise.

### GetUploadsOk

`func (o *ServiceRequestDocument) GetUploadsOk() (*[]UploadWithOmissions, bool)`

GetUploadsOk returns a tuple with the Uploads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUploads

`func (o *ServiceRequestDocument) SetUploads(v []UploadWithOmissions)`

SetUploads sets Uploads field to given value.

### HasUploads

`func (o *ServiceRequestDocument) HasUploads() bool`

HasUploads returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


