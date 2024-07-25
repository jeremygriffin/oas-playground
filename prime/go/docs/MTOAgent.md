# MTOAgent

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
**AgentType** | Pointer to [**MTOAgentType**](MTOAgentType.md) |  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewMTOAgent

`func NewMTOAgent() *MTOAgent`

NewMTOAgent instantiates a new MTOAgent object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMTOAgentWithDefaults

`func NewMTOAgentWithDefaults() *MTOAgent`

NewMTOAgentWithDefaults instantiates a new MTOAgent object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *MTOAgent) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MTOAgent) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MTOAgent) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MTOAgent) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMtoShipmentID

`func (o *MTOAgent) GetMtoShipmentID() string`

GetMtoShipmentID returns the MtoShipmentID field if non-nil, zero value otherwise.

### GetMtoShipmentIDOk

`func (o *MTOAgent) GetMtoShipmentIDOk() (*string, bool)`

GetMtoShipmentIDOk returns a tuple with the MtoShipmentID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMtoShipmentID

`func (o *MTOAgent) SetMtoShipmentID(v string)`

SetMtoShipmentID sets MtoShipmentID field to given value.

### HasMtoShipmentID

`func (o *MTOAgent) HasMtoShipmentID() bool`

HasMtoShipmentID returns a boolean if a field has been set.

### GetCreatedAt

`func (o *MTOAgent) GetCreatedAt() time.Time`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *MTOAgent) GetCreatedAtOk() (*time.Time, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *MTOAgent) SetCreatedAt(v time.Time)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *MTOAgent) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetUpdatedAt

`func (o *MTOAgent) GetUpdatedAt() time.Time`

GetUpdatedAt returns the UpdatedAt field if non-nil, zero value otherwise.

### GetUpdatedAtOk

`func (o *MTOAgent) GetUpdatedAtOk() (*time.Time, bool)`

GetUpdatedAtOk returns a tuple with the UpdatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdatedAt

`func (o *MTOAgent) SetUpdatedAt(v time.Time)`

SetUpdatedAt sets UpdatedAt field to given value.

### HasUpdatedAt

`func (o *MTOAgent) HasUpdatedAt() bool`

HasUpdatedAt returns a boolean if a field has been set.

### GetFirstName

`func (o *MTOAgent) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *MTOAgent) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *MTOAgent) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *MTOAgent) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *MTOAgent) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *MTOAgent) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *MTOAgent) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *MTOAgent) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *MTOAgent) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *MTOAgent) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *MTOAgent) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *MTOAgent) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetEmail

`func (o *MTOAgent) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *MTOAgent) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *MTOAgent) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *MTOAgent) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### SetEmailNil

`func (o *MTOAgent) SetEmailNil(b bool)`

 SetEmailNil sets the value for Email to be an explicit nil

### UnsetEmail
`func (o *MTOAgent) UnsetEmail()`

UnsetEmail ensures that no value is present for Email, not even an explicit nil
### GetPhone

`func (o *MTOAgent) GetPhone() string`

GetPhone returns the Phone field if non-nil, zero value otherwise.

### GetPhoneOk

`func (o *MTOAgent) GetPhoneOk() (*string, bool)`

GetPhoneOk returns a tuple with the Phone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhone

`func (o *MTOAgent) SetPhone(v string)`

SetPhone sets Phone field to given value.

### HasPhone

`func (o *MTOAgent) HasPhone() bool`

HasPhone returns a boolean if a field has been set.

### SetPhoneNil

`func (o *MTOAgent) SetPhoneNil(b bool)`

 SetPhoneNil sets the value for Phone to be an explicit nil

### UnsetPhone
`func (o *MTOAgent) UnsetPhone()`

UnsetPhone ensures that no value is present for Phone, not even an explicit nil
### GetAgentType

`func (o *MTOAgent) GetAgentType() MTOAgentType`

GetAgentType returns the AgentType field if non-nil, zero value otherwise.

### GetAgentTypeOk

`func (o *MTOAgent) GetAgentTypeOk() (*MTOAgentType, bool)`

GetAgentTypeOk returns a tuple with the AgentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentType

`func (o *MTOAgent) SetAgentType(v MTOAgentType)`

SetAgentType sets AgentType field to given value.

### HasAgentType

`func (o *MTOAgent) HasAgentType() bool`

HasAgentType returns a boolean if a field has been set.

### GetETag

`func (o *MTOAgent) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *MTOAgent) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *MTOAgent) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *MTOAgent) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


