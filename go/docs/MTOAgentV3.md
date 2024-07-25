# MTOAgentV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | The ID of the agent. | [optional] [readonly] 
**MtoShipmentID** | Pointer to **string** | The ID of the shipment this agent is permitted to release/receive. | [optional] [readonly] 
**CreatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**UpdatedAt** | Pointer to **time.Time** |  | [optional] [readonly] 
**FirstName** | Pointer to **NullableString** |  | [optional] 
**LastName** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **NullableString** |  | [optional] 
**Phone** | Pointer to **NullableString** |  | [optional] 
**AgentType** | Pointer to [**MTOAgentTypeV3V3**](MTOAgentTypeV3.md) |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewMTOAgentV3

`func NewMTOAgentV3() *MTOAgentV3`

NewMTOAgentV3 instantiates a new MTOAgentV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTOAgentV3WithDefaults

`func NewMTOAgentV3WithDefaults() *MTOAgentV3`

NewMTOAgentV3WithDefaults instantiates a new MTOAgentV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MTOAgentV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MTOAgentV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MTOAgentV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MTOAgentV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMtoShipmentID

`func (o *MTOAgentV3) GetMtoShipmentID() string`

GetMtoShipmentID returns the MtoShipmentID field if non-nil, zero value otherwise.

### GetMtoShipmentIDOk

`func (o *MTOAgentV3) GetMtoShipmentIDOk() (*string, bool)`

GetMtoShipmentIDOk returns a tuple with the MtoShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoShipmentID

`func (o *MTOAgentV3) SetMtoShipmentID(v string)`

SetMtoShipmentID sets MtoShipmentID field to given value.

### HasMtoShipmentID

`func (o *MTOAgentV3) HasMtoShipmentID() bool`

HasMtoShipmentID returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MTOAgentV3) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MTOAgentV3) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MTOAgentV3) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MTOAgentV3) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MTOAgentV3) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MTOAgentV3) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MTOAgentV3) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MTOAgentV3) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFirstName

`func (o *MTOAgentV3) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MTOAgentV3) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MTOAgentV3) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MTOAgentV3) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *MTOAgentV3) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *MTOAgentV3) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *MTOAgentV3) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MTOAgentV3) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MTOAgentV3) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MTOAgentV3) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *MTOAgentV3) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *MTOAgentV3) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetEmail

`func (o *MTOAgentV3) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MTOAgentV3) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MTOAgentV3) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MTOAgentV3) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MTOAgentV3) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MTOAgentV3) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *MTOAgentV3) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MTOAgentV3) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MTOAgentV3) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MTOAgentV3) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *MTOAgentV3) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *MTOAgentV3) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetAgentType

`func (o *MTOAgentV3) GetAgentType() MTOAgentTypeV3V3`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *MTOAgentV3) GetAgentTypeOk() (*MTOAgentTypeV3V3, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *MTOAgentV3) SetAgentType(v MTOAgentTypeV3V3)`

SetAgentType sets AgentType field to given value.

### HasAgentType

`func (o *MTOAgentV3) HasAgentType() bool`

HasAgentType returns a boolean if a field has been set.

### GetETag

`func (o *MTOAgentV3) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MTOAgentV3) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MTOAgentV3) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MTOAgentV3) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


