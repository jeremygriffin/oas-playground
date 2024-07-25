# CustomerV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DodID** | Pointer to **string** |  | [optional] 
**Emplid** | Pointer to **string** |  | [optional] 
**UserID** | Pointer to **string** |  | [optional] 
**CurrentAddress** | Pointer to [**AddressV2V2**](AddressV2.md) |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewCustomerV2

`func NewCustomerV2() *CustomerV2`

NewCustomerV2 instantiates a new CustomerV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerV2WithDefaults

`func NewCustomerV2WithDefaults() *CustomerV2`

NewCustomerV2WithDefaults instantiates a new CustomerV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CustomerV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CustomerV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDodID

`func (o *CustomerV2) GetDodID() string`

GetDodID returns the DodID field if non-nil, zero value otherwise.

### GetDodIDOk

`func (o *CustomerV2) GetDodIDOk() (*string, bool)`

GetDodIDOk returns a tuple with the DodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDodID

`func (o *CustomerV2) SetDodID(v string)`

SetDodID sets DodID field to given value.

### HasDodID

`func (o *CustomerV2) HasDodID() bool`

HasDodID returns a boolean if a field has been set.

### GetEmplid

`func (o *CustomerV2) GetEmplid() string`

GetEmplid returns the Emplid field if non-nil, zero value otherwise.

### GetEmplidOk

`func (o *CustomerV2) GetEmplidOk() (*string, bool)`

GetEmplidOk returns a tuple with the Emplid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmplid

`func (o *CustomerV2) SetEmplid(v string)`

SetEmplid sets Emplid field to given value.

### HasEmplid

`func (o *CustomerV2) HasEmplid() bool`

HasEmplid returns a boolean if a field has been set.

### GetUserID

`func (o *CustomerV2) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *CustomerV2) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *CustomerV2) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *CustomerV2) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetCurrentAddress

`func (o *CustomerV2) GetCurrentAddress() AddressV2V2`

GetCurrentAddress returns the CurrentAddress field if non-nil, zero value otherwise.

### GetCurrentAddressOk

`func (o *CustomerV2) GetCurrentAddressOk() (*AddressV2V2, bool)`

GetCurrentAddressOk returns a tuple with the CurrentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAddress

`func (o *CustomerV2) SetCurrentAddress(v AddressV2V2)`

SetCurrentAddress sets CurrentAddress field to given value.

### HasCurrentAddress

`func (o *CustomerV2) HasCurrentAddress() bool`

HasCurrentAddress returns a boolean if a field has been set.

### GetFirstName

`func (o *CustomerV2) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *CustomerV2) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *CustomerV2) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *CustomerV2) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *CustomerV2) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *CustomerV2) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *CustomerV2) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *CustomerV2) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetBranch

`func (o *CustomerV2) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *CustomerV2) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *CustomerV2) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *CustomerV2) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetPhone

`func (o *CustomerV2) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *CustomerV2) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *CustomerV2) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *CustomerV2) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *CustomerV2) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *CustomerV2) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *CustomerV2) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *CustomerV2) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetETag

`func (o *CustomerV2) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *CustomerV2) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *CustomerV2) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *CustomerV2) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


