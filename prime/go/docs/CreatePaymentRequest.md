# CreatePaymentRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IsFinal** | Pointer to **bool** |  | [optional] [default to false]
**MoveTaskOrderID** | **string** |  | 
**ServiceItems** | [**[]ServiceItem**](ServiceItem.md) |  | 
**PointOfContact** | Pointer to **string** | Email or id of a contact person for this update. | [optional] 

## Methods

### NewCreatePaymentRequest

`func NewCreatePaymentRequest(moveTaskOrderID string, serviceItems []ServiceItem, ) *CreatePaymentRequest`

NewCreatePaymentRequest instantiates a new CreatePaymentRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreatePaymentRequestWithDefaults

`func NewCreatePaymentRequestWithDefaults() *CreatePaymentRequest`

NewCreatePaymentRequestWithDefaults instantiates a new CreatePaymentRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIsFinal

`func (o *CreatePaymentRequest) GetIsFinal() bool`

GetIsFinal returns the IsFinal field if non-nil, zero value otherwise.

### GetIsFinalOk

`func (o *CreatePaymentRequest) GetIsFinalOk() (*bool, bool)`

GetIsFinalOk returns a tuple with the IsFinal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsFinal

`func (o *CreatePaymentRequest) SetIsFinal(v bool)`

SetIsFinal sets IsFinal field to given value.

### HasIsFinal

`func (o *CreatePaymentRequest) HasIsFinal() bool`

HasIsFinal returns a boolean if a field has been set.

### GetMoveTaskOrderID

`func (o *CreatePaymentRequest) GetMoveTaskOrderID() string`

GetMoveTaskOrderID returns the MoveTaskOrderID field if non-nil, zero value otherwise.

### GetMoveTaskOrderIDOk

`func (o *CreatePaymentRequest) GetMoveTaskOrderIDOk() (*string, bool)`

GetMoveTaskOrderIDOk returns a tuple with the MoveTaskOrderID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveTaskOrderID

`func (o *CreatePaymentRequest) SetMoveTaskOrderID(v string)`

SetMoveTaskOrderID sets MoveTaskOrderID field to given value.


### GetServiceItems

`func (o *CreatePaymentRequest) GetServiceItems() []ServiceItem`

GetServiceItems returns the ServiceItems field if non-nil, zero value otherwise.

### GetServiceItemsOk

`func (o *CreatePaymentRequest) GetServiceItemsOk() (*[]ServiceItem, bool)`

GetServiceItemsOk returns a tuple with the ServiceItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceItems

`func (o *CreatePaymentRequest) SetServiceItems(v []ServiceItem)`

SetServiceItems sets ServiceItems field to given value.


### GetPointOfContact

`func (o *CreatePaymentRequest) GetPointOfContact() string`

GetPointOfContact returns the PointOfContact field if non-nil, zero value otherwise.

### GetPointOfContactOk

`func (o *CreatePaymentRequest) GetPointOfContactOk() (*string, bool)`

GetPointOfContactOk returns a tuple with the PointOfContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPointOfContact

`func (o *CreatePaymentRequest) SetPointOfContact(v string)`

SetPointOfContact sets PointOfContact field to given value.

### HasPointOfContact

`func (o *CreatePaymentRequest) HasPointOfContact() bool`

HasPointOfContact returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


