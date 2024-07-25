# ExcessWeightRecord

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
**MoveId** | **string** | The UUID of the move this excess weight record belongs to. | 
**MoveExcessWeightQualifiedAt** | Pointer to **NullableTime** | The date and time when the sum of all the move&#39;s shipments met the excess weight qualification threshold. The system monitors these weights and will update this field automatically.  | [optional] [readonly] 
**MoveExcessWeightAcknowledgedAt** | Pointer to **NullableTime** | The date and time when the TOO acknowledged the excess weight alert, either by dismissing the risk or updating the max billable weight. This will occur after the excess weight record has been uploaded.  | [optional] [readonly] 

## Methods

### NewExcessWeightRecord

`func NewExcessWeightRecord(filename string, contentType string, bytes int32, moveId string, ) *ExcessWeightRecord`

NewExcessWeightRecord instantiates a new ExcessWeightRecord object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExcessWeightRecordWithDefaults

`func NewExcessWeightRecordWithDefaults() *ExcessWeightRecord`

NewExcessWeightRecordWithDefaults instantiates a new ExcessWeightRecord object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExcessWeightRecord) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExcessWeightRecord) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExcessWeightRecord) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExcessWeightRecord) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUrl

`func (o *ExcessWeightRecord) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ExcessWeightRecord) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ExcessWeightRecord) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ExcessWeightRecord) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetFilename

`func (o *ExcessWeightRecord) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *ExcessWeightRecord) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *ExcessWeightRecord) SetFilename(v string)`

SetFilename sets Filename field to given value.


### GetContentType

`func (o *ExcessWeightRecord) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ExcessWeightRecord) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ExcessWeightRecord) SetContentType(v string)`

SetContentType sets ContentType field to given value.


### GetBytes

`func (o *ExcessWeightRecord) GetBytes() int32`

GetBytes returns the Bytes field if non-nil, zero value otherwise.

### GetBytesOk

`func (o *ExcessWeightRecord) GetBytesOk() (*int32, bool)`

GetBytesOk returns a tuple with the Bytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBytes

`func (o *ExcessWeightRecord) SetBytes(v int32)`

SetBytes sets Bytes field to given value.


### GetStatus

`func (o *ExcessWeightRecord) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ExcessWeightRecord) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ExcessWeightRecord) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ExcessWeightRecord) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ExcessWeightRecord) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ExcessWeightRecord) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ExcessWeightRecord) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ExcessWeightRecord) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *ExcessWeightRecord) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ExcessWeightRecord) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ExcessWeightRecord) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ExcessWeightRecord) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetMoveId

`func (o *ExcessWeightRecord) GetMoveId() string`

GetMoveId returns the MoveId field if non-nil, zero value otherwise.

### GetMoveIdOk

`func (o *ExcessWeightRecord) GetMoveIdOk() (*string, bool)`

GetMoveIdOk returns a tuple with the MoveId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveId

`func (o *ExcessWeightRecord) SetMoveId(v string)`

SetMoveId sets MoveId field to given value.


### GetMoveExcessWeightQualifiedAt

`func (o *ExcessWeightRecord) GetMoveExcessWeightQualifiedAt() time.Time`

GetMoveExcessWeightQualifiedAt returns the MoveExcessWeightQualifiedAt field if non-nil, zero value otherwise.

### GetMoveExcessWeightQualifiedAtOk

`func (o *ExcessWeightRecord) GetMoveExcessWeightQualifiedAtOk() (*time.Time, bool)`

GetMoveExcessWeightQualifiedAtOk returns a tuple with the MoveExcessWeightQualifiedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveExcessWeightQualifiedAt

`func (o *ExcessWeightRecord) SetMoveExcessWeightQualifiedAt(v time.Time)`

SetMoveExcessWeightQualifiedAt sets MoveExcessWeightQualifiedAt field to given value.

### HasMoveExcessWeightQualifiedAt

`func (o *ExcessWeightRecord) HasMoveExcessWeightQualifiedAt() bool`

HasMoveExcessWeightQualifiedAt returns a boolean if a field has been set.

### SetMoveExcessWeightQualifiedAtNil

`func (o *ExcessWeightRecord) SetMoveExcessWeightQualifiedAtNil(b bool)`

 SetMoveExcessWeightQualifiedAtNil sets the value for MoveExcessWeightQualifiedAt to be an explicit nil

### UnsetMoveExcessWeightQualifiedAt
`func (o *ExcessWeightRecord) UnsetMoveExcessWeightQualifiedAt()`

UnsetMoveExcessWeightQualifiedAt ensures that no value is present for MoveExcessWeightQualifiedAt, not even an explicit nil
### GetMoveExcessWeightAcknowledgedAt

`func (o *ExcessWeightRecord) GetMoveExcessWeightAcknowledgedAt() time.Time`

GetMoveExcessWeightAcknowledgedAt returns the MoveExcessWeightAcknowledgedAt field if non-nil, zero value otherwise.

### GetMoveExcessWeightAcknowledgedAtOk

`func (o *ExcessWeightRecord) GetMoveExcessWeightAcknowledgedAtOk() (*time.Time, bool)`

GetMoveExcessWeightAcknowledgedAtOk returns a tuple with the MoveExcessWeightAcknowledgedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveExcessWeightAcknowledgedAt

`func (o *ExcessWeightRecord) SetMoveExcessWeightAcknowledgedAt(v time.Time)`

SetMoveExcessWeightAcknowledgedAt sets MoveExcessWeightAcknowledgedAt field to given value.

### HasMoveExcessWeightAcknowledgedAt

`func (o *ExcessWeightRecord) HasMoveExcessWeightAcknowledgedAt() bool`

HasMoveExcessWeightAcknowledgedAt returns a boolean if a field has been set.

### SetMoveExcessWeightAcknowledgedAtNil

`func (o *ExcessWeightRecord) SetMoveExcessWeightAcknowledgedAtNil(b bool)`

 SetMoveExcessWeightAcknowledgedAtNil sets the value for MoveExcessWeightAcknowledgedAt to be an explicit nil

### UnsetMoveExcessWeightAcknowledgedAt
`func (o *ExcessWeightRecord) UnsetMoveExcessWeightAcknowledgedAt()`

UnsetMoveExcessWeightAcknowledgedAt ensures that no value is present for MoveExcessWeightAcknowledgedAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


