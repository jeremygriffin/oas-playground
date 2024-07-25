# UploadWithOmissionsV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Filename** | **string** |  | 
**ContentType** | **string** |  | 
**Bytes** | **int32** |  | 
**Status** | Pointer to **string** |  | [optional] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 

## Methods

### NewUploadWithOmissionsV3

`func NewUploadWithOmissionsV3(filename string, contentType string, bytes int32, ) *UploadWithOmissionsV3`

NewUploadWithOmissionsV3 instantiates a new UploadWithOmissionsV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadWithOmissionsV3WithDefaults

`func NewUploadWithOmissionsV3WithDefaults() *UploadWithOmissionsV3`

NewUploadWithOmissionsV3WithDefaults instantiates a new UploadWithOmissionsV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UploadWithOmissionsV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UploadWithOmissionsV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UploadWithOmissionsV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *UploadWithOmissionsV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *UploadWithOmissionsV3) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UploadWithOmissionsV3) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UploadWithOmissionsV3) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UploadWithOmissionsV3) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFilename

`func (o *UploadWithOmissionsV3) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *UploadWithOmissionsV3) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *UploadWithOmissionsV3) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetContentType

`func (o *UploadWithOmissionsV3) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *UploadWithOmissionsV3) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *UploadWithOmissionsV3) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetBytes

`func (o *UploadWithOmissionsV3) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *UploadWithOmissionsV3) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *UploadWithOmissionsV3) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### GetStatus

`func (o *UploadWithOmissionsV3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UploadWithOmissionsV3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UploadWithOmissionsV3) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UploadWithOmissionsV3) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *UploadWithOmissionsV3) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *UploadWithOmissionsV3) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *UploadWithOmissionsV3) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *UploadWithOmissionsV3) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *UploadWithOmissionsV3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *UploadWithOmissionsV3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *UploadWithOmissionsV3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *UploadWithOmissionsV3) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


