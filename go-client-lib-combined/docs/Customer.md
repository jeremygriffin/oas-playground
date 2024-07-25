# Customer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**DodID** | Pointer to **string** |  | [optional] 
**Emplid** | Pointer to **string** |  | [optional] 
**UserID** | Pointer to **string** |  | [optional] 
**CurrentAddress** | Pointer to [**Address**](Address.md) |  | [optional] 
**FirstName** | Pointer to **string** |  | [optional] 
**LastName** | Pointer to **string** |  | [optional] 
**Branch** | Pointer to **string** |  | [optional] 
**Phone** | Pointer to **string** |  | [optional] 
**Email** | Pointer to **string** |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewCustomer

`func NewCustomer() *Customer`

NewCustomer instantiates a new Customer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerWithDefaults

`func NewCustomerWithDefaults() *Customer`

NewCustomerWithDefaults instantiates a new Customer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Customer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Customer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Customer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Customer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDodID

`func (o *Customer) GetDodID() string`

GetDodID returns the DodID field if non-nil, zero value otherwise.

### GetDodIDOk

`func (o *Customer) GetDodIDOk() (*string, bool)`

GetDodIDOk returns a tuple with the DodID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDodID

`func (o *Customer) SetDodID(v string)`

SetDodID sets DodID field to given value.

### HasDodID

`func (o *Customer) HasDodID() bool`

HasDodID returns a boolean if a field has been set.

### GetEmplid

`func (o *Customer) GetEmplid() string`

GetEmplid returns the Emplid field if non-nil, zero value otherwise.

### GetEmplidOk

`func (o *Customer) GetEmplidOk() (*string, bool)`

GetEmplidOk returns a tuple with the Emplid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmplid

`func (o *Customer) SetEmplid(v string)`

SetEmplid sets Emplid field to given value.

### HasEmplid

`func (o *Customer) HasEmplid() bool`

HasEmplid returns a boolean if a field has been set.

### GetUserID

`func (o *Customer) GetUserID() string`

GetUserID returns the UserID field if non-nil, zero value otherwise.

### GetUserIDOk

`func (o *Customer) GetUserIDOk() (*string, bool)`

GetUserIDOk returns a tuple with the UserID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserID

`func (o *Customer) SetUserID(v string)`

SetUserID sets UserID field to given value.

### HasUserID

`func (o *Customer) HasUserID() bool`

HasUserID returns a boolean if a field has been set.

### GetCurrentAddress

`func (o *Customer) GetCurrentAddress() Address`

GetCurrentAddress returns the CurrentAddress field if non-nil, zero value otherwise.

### GetCurrentAddressOk

`func (o *Customer) GetCurrentAddressOk() (*Address, bool)`

GetCurrentAddressOk returns a tuple with the CurrentAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentAddress

`func (o *Customer) SetCurrentAddress(v Address)`

SetCurrentAddress sets CurrentAddress field to given value.

### HasCurrentAddress

`func (o *Customer) HasCurrentAddress() bool`

HasCurrentAddress returns a boolean if a field has been set.

### GetFirstName

`func (o *Customer) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Customer) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Customer) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *Customer) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *Customer) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Customer) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Customer) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *Customer) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetBranch

`func (o *Customer) GetBranch() string`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *Customer) GetBranchOk() (*string, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *Customer) SetBranch(v string)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *Customer) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetPhone

`func (o *Customer) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *Customer) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *Customer) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *Customer) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### GetEmail

`func (o *Customer) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Customer) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Customer) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *Customer) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetETag

`func (o *Customer) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *Customer) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *Customer) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *Customer) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


