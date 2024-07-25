# DutyLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AddressID** | Pointer to **string** |  | [optional] 
**Address** | Pointer to [**Address**](Address.md) |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewDutyLocation

`func NewDutyLocation() *DutyLocation`

NewDutyLocation instantiates a new DutyLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDutyLocationWithDefaults

`func NewDutyLocationWithDefaults() *DutyLocation`

NewDutyLocationWithDefaults instantiates a new DutyLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DutyLocation) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DutyLocation) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DutyLocation) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DutyLocation) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *DutyLocation) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DutyLocation) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DutyLocation) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DutyLocation) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddressID

`func (o *DutyLocation) GetAddressID() string`

GetAddressID returns the AddressID field if non-nil, zero value otherwise.

### GetAddressIDOk

`func (o *DutyLocation) GetAddressIDOk() (*string, bool)`

GetAddressIDOk returns a tuple with the AddressID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddressID

`func (o *DutyLocation) SetAddressID(v string)`

SetAddressID sets AddressID field to given value.

### HasAddressID

`func (o *DutyLocation) HasAddressID() bool`

HasAddressID returns a boolean if a field has been set.

### GetAddress

`func (o *DutyLocation) GetAddress() Address`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *DutyLocation) GetAddressOk() (*Address, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *DutyLocation) SetAddress(v Address)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *DutyLocation) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetETag

`func (o *DutyLocation) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *DutyLocation) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *DutyLocation) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *DutyLocation) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


