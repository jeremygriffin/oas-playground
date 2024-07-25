# ListMove

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**MoveCode** | Pointer to **string** |  | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**OrderID** | Pointer to **string** |  | [optional] 
**ReferenceId** | Pointer to **string** |  | [optional] 
**AvailableToPrimeAt** | Pointer to **NullableTime** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**PpmType** | Pointer to **string** |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 
**Amendments** | Pointer to [**Amendments**](Amendments.md) |  | [optional] 

## Methods

### NewListMove

`func NewListMove() *ListMove`

NewListMove instantiates a new ListMove object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListMoveWithDefaults

`func NewListMoveWithDefaults() *ListMove`

NewListMoveWithDefaults instantiates a new ListMove object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListMove) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListMove) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListMove) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ListMove) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMoveCode

`func (o *ListMove) GetMoveCode() string`

GetMoveCode returns the MoveCode field if non-nil, zero value otherwise.

### GetMoveCodeOk

`func (o *ListMove) GetMoveCodeOk() (*string, bool)`

GetMoveCodeOk returns a tuple with the MoveCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveCode

`func (o *ListMove) SetMoveCode(v string)`

SetMoveCode sets MoveCode field to given value.

### HasMoveCode

`func (o *ListMove) HasMoveCode() bool`

HasMoveCode returns a boolean if a field has been set.

### GetCreatedAt

`func (o *ListMove) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *ListMove) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *ListMove) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *ListMove) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetOrderID

`func (o *ListMove) GetOrderID() string`

GetOrderID returns the OrderID field if non-nil, zero value otherwise.

### GetOrderIDOk

`func (o *ListMove) GetOrderIDOk() (*string, bool)`

GetOrderIDOk returns a tuple with the OrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderID

`func (o *ListMove) SetOrderID(v string)`

SetOrderID sets OrderID field to given value.

### HasOrderID

`func (o *ListMove) HasOrderID() bool`

HasOrderID returns a boolean if a field has been set.

### GetReferenceId

`func (o *ListMove) GetReferenceId() string`

GetReferenceId returns the ReferenceId field if non-nil, zero value otherwise.

### GetReferenceIdOk

`func (o *ListMove) GetReferenceIdOk() (*string, bool)`

GetReferenceIdOk returns a tuple with the ReferenceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReferenceId

`func (o *ListMove) SetReferenceId(v string)`

SetReferenceId sets ReferenceId field to given value.

### HasReferenceId

`func (o *ListMove) HasReferenceId() bool`

HasReferenceId returns a boolean if a field has been set.

### GetAvailableToPrimeAt

`func (o *ListMove) GetAvailableToPrimeAt() time.Time`

GetAvailableToPrimeAt returns the AvailableToPrimeAt field if non-nil, zero value otherwise.

### GetAvailableToPrimeAtOk

`func (o *ListMove) GetAvailableToPrimeAtOk() (*time.Time, bool)`

GetAvailableToPrimeAtOk returns a tuple with the AvailableToPrimeAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailableToPrimeAt

`func (o *ListMove) SetAvailableToPrimeAt(v time.Time)`

SetAvailableToPrimeAt sets AvailableToPrimeAt field to given value.

### HasAvailableToPrimeAt

`func (o *ListMove) HasAvailableToPrimeAt() bool`

HasAvailableToPrimeAt returns a boolean if a field has been set.

### SetAvailableToPrimeAtNil

`func (o *ListMove) SetAvailableToPrimeAtNil(b bool)`

 SetAvailableToPrimeAtNil sets the value for AvailableToPrimeAt to be an explicit nil

### UnsetAvailableToPrimeAt
`func (o *ListMove) UnsetAvailableToPrimeAt()`

UnsetAvailableToPrimeAt ensures that no value is present for AvailableToPrimeAt, not even an explicit nil
### GetUpdatedAt

`func (o *ListMove) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *ListMove) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *ListMove) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *ListMove) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetPpmType

`func (o *ListMove) GetPpmType() string`

GetPpmType returns the PpmType field if non-nil, zero value otherwise.

### GetPpmTypeOk

`func (o *ListMove) GetPpmTypeOk() (*string, bool)`

GetPpmTypeOk returns a tuple with the PpmType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPpmType

`func (o *ListMove) SetPpmType(v string)`

SetPpmType sets PpmType field to given value.

### HasPpmType

`func (o *ListMove) HasPpmType() bool`

HasPpmType returns a boolean if a field has been set.

### GetETag

`func (o *ListMove) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *ListMove) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *ListMove) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *ListMove) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetAmendments

`func (o *ListMove) GetAmendments() Amendments`

GetAmendments returns the Amendments field if non-nil, zero value otherwise.

### GetAmendmentsOk

`func (o *ListMove) GetAmendmentsOk() (*Amendments, bool)`

GetAmendmentsOk returns a tuple with the Amendments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmendments

`func (o *ListMove) SetAmendments(v Amendments)`

SetAmendments sets Amendments field to given value.

### HasAmendments

`func (o *ListMove) HasAmendments() bool`

HasAmendments returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


