# CustomerV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DodID** | Pointer to **string** |  | [optional] 
**Emplid** | Pointer to **string** |  | [optional] 
**UserID** | Pointer to **string** |  | [optional] 
**CurrentAddress** | Pointer to [**AddressV3V3**](AddressV3.md) |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewCustomerV3

`func NewCustomerV3() *CustomerV3`

NewCustomerV3 instantiates a new CustomerV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerV3WithDefaults

`func NewCustomerV3WithDefaults() *CustomerV3`

NewCustomerV3WithDefaults instantiates a new CustomerV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDodID

`func (o *CustomerV3) GetDodID() string`

GetDodID returns the DodID field if non-nil, zero value otherwise.

### GetDodIDOk

`func (o *CustomerV3) GetDodIDOk() (*string, bool)`

GetDodIDOk returns a tuple with the DodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDodID

`func (o *CustomerV3) SetDodID(v string)`

SetDodID sets DodID field to given value.

### HasDodID

`func (o *CustomerV3) HasDodID() bool`

HasDodID returns a boolean if a field has been set.

### GetEmplid

`func (o *CustomerV3) GetEmplid() string`

GetEmplid returns the Emplid field if non-nil, zero value otherwise.

### GetEmplidOk

`func (o *CustomerV3) GetEmplidOk() (*string, bool)`

GetEmplidOk returns a tuple with the Emplid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmplid

`func (o *CustomerV3) SetEmplid(v string)`

SetEmplid sets Emplid field to given value.

### HasEmplid

`func (o *CustomerV3) HasEmplid() bool`

HasEmplid returns a boolean if a field has been set.

### GetUserID

`func (o *CustomerV3) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *CustomerV3) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *CustomerV3) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *CustomerV3) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetCurrentAddress

`func (o *CustomerV3) GetCurrentAddress() AddressV3V3`

GetCurrentAddress returns the CurrentAddress field if non-nil, zero value otherwise.

### GetCurrentAddressOk

`func (o *CustomerV3) GetCurrentAddressOk() (*AddressV3V3, bool)`

GetCurrentAddressOk returns a tuple with the CurrentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAddress

`func (o *CustomerV3) SetCurrentAddress(v AddressV3V3)`

SetCurrentAddress sets CurrentAddress field to given value.

### HasCurrentAddress

`func (o *CustomerV3) HasCurrentAddress() bool`

HasCurrentAddress returns a boolean if a field has been set.

### GetFirstName

`func (o *CustomerV3) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CustomerV3) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CustomerV3) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CustomerV3) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CustomerV3) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CustomerV3) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CustomerV3) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CustomerV3) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetBranch

`func (o *CustomerV3) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CustomerV3) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CustomerV3) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *CustomerV3) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerV3) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerV3) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerV3) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerV3) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerV3) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerV3) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerV3) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerV3) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetETag

`func (o *CustomerV3) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *CustomerV3) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *CustomerV3) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *CustomerV3) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


