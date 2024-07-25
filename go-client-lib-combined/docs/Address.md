# Address

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**StreetAddress1** | **string** |  | 
**StreetAddress2** | Pointer to **NullableString** |  | [optional] 
**StreetAddress3** | Pointer to **NullableString** |  | [optional] 
**City** | **string** |  | 
**ETag** | Pointer to **string** |  | [optional] [readonly] 
**State** | **string** |  | 
**PostalCode** | **string** |  | 
**Country** | Pointer to **NullableString** |  | [optional] [default to "USA"]
**County** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAddress

`func NewAddress(streetAddress1 string, city string, state string, postalCode string, ) *Address`

NewAddress instantiates a new Address object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressWithDefaults

`func NewAddressWithDefaults() *Address`

NewAddressWithDefaults instantiates a new Address object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Address) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Address) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Address) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Address) HasId() bool`

HasId returns a boolean if a field has been set.

### GetStreetAddress1

`func (o *Address) GetStreetAddress1() string`

GetStreetAddress1 returns the StreetAddress1 field if non-nil, zero value otherwise.

### GetStreetAddress1Ok

`func (o *Address) GetStreetAddress1Ok() (*string, bool)`

GetStreetAddress1Ok returns a tuple with the StreetAddress1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress1

`func (o *Address) SetStreetAddress1(v string)`

SetStreetAddress1 sets StreetAddress1 field to given value.


### GetStreetAddress2

`func (o *Address) GetStreetAddress2() string`

GetStreetAddress2 returns the StreetAddress2 field if non-nil, zero value otherwise.

### GetStreetAddress2Ok

`func (o *Address) GetStreetAddress2Ok() (*string, bool)`

GetStreetAddress2Ok returns a tuple with the StreetAddress2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress2

`func (o *Address) SetStreetAddress2(v string)`

SetStreetAddress2 sets StreetAddress2 field to given value.

### HasStreetAddress2

`func (o *Address) HasStreetAddress2() bool`

HasStreetAddress2 returns a boolean if a field has been set.

### SetStreetAddress2Nil

`func (o *Address) SetStreetAddress2Nil(b bool)`

 SetStreetAddress2Nil sets the value for StreetAddress2 to be an explicit nil

### UnsetStreetAddress2
`func (o *Address) UnsetStreetAddress2()`

UnsetStreetAddress2 ensures that no value is present for StreetAddress2, not even an explicit nil
### GetStreetAddress3

`func (o *Address) GetStreetAddress3() string`

GetStreetAddress3 returns the StreetAddress3 field if non-nil, zero value otherwise.

### GetStreetAddress3Ok

`func (o *Address) GetStreetAddress3Ok() (*string, bool)`

GetStreetAddress3Ok returns a tuple with the StreetAddress3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreetAddress3

`func (o *Address) SetStreetAddress3(v string)`

SetStreetAddress3 sets StreetAddress3 field to given value.

### HasStreetAddress3

`func (o *Address) HasStreetAddress3() bool`

HasStreetAddress3 returns a boolean if a field has been set.

### SetStreetAddress3Nil

`func (o *Address) SetStreetAddress3Nil(b bool)`

 SetStreetAddress3Nil sets the value for StreetAddress3 to be an explicit nil

### UnsetStreetAddress3
`func (o *Address) UnsetStreetAddress3()`

UnsetStreetAddress3 ensures that no value is present for StreetAddress3, not even an explicit nil
### GetCity

`func (o *Address) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Address) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Address) SetCity(v string)`

SetCity sets City field to given value.


### GetETag

`func (o *Address) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *Address) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *Address) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *Address) HasETag() bool`

HasETag returns a boolean if a field has been set.

### GetState

`func (o *Address) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Address) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Address) SetState(v string)`

SetState sets State field to given value.


### GetPostalCode

`func (o *Address) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Address) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Address) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCountry

`func (o *Address) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Address) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Address) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Address) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *Address) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Address) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetCounty

`func (o *Address) GetCounty() string`

GetCounty returns the County field if non-nil, zero value otherwise.

### GetCountyOk

`func (o *Address) GetCountyOk() (*string, bool)`

GetCountyOk returns a tuple with the County field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCounty

`func (o *Address) SetCounty(v string)`

SetCounty sets County field to given value.

### HasCounty

`func (o *Address) HasCounty() bool`

HasCounty returns a boolean if a field has been set.

### SetCountyNil

`func (o *Address) SetCountyNil(b bool)`

 SetCountyNil sets the value for County to be an explicit nil

### UnsetCounty
`func (o *Address) UnsetCounty()`

UnsetCounty ensures that no value is present for County, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


