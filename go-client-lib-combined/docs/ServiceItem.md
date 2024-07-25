# ServiceItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Params** | Pointer to [**[]ServiceItemParamsInner**](ServiceItemParamsInner.md) | This should be populated for the following service items:   * DOASIT(Domestic origin Additional day SIT)   * DDASIT(Domestic destination Additional day SIT)  Both take in the following param keys:   * &#x60;SITPaymentRequestStart&#x60;   * &#x60;SITPaymentRequestEnd&#x60;  The value of each is a date string in the format \&quot;YYYY-MM-DD\&quot; (e.g. \&quot;2023-01-15\&quot;)  | [optional] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewServiceItem

`func NewServiceItem() *ServiceItem`

NewServiceItem instantiates a new ServiceItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServiceItemWithDefaults

`func NewServiceItemWithDefaults() *ServiceItem`

NewServiceItemWithDefaults instantiates a new ServiceItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServiceItem) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServiceItem) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServiceItem) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ServiceItem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParams

`func (o *ServiceItem) GetParams() []ServiceItemParamsInner`

GetParams returns the Params field if non-nil, zero value otherwise.

### GetParamsOk

`func (o *ServiceItem) GetParamsOk() (*[]ServiceItemParamsInner, bool)`

GetParamsOk returns a tuple with the Params field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParams

`func (o *ServiceItem) SetParams(v []ServiceItemParamsInner)`

SetParams sets Params field to given value.

### HasParams

`func (o *ServiceItem) HasParams() bool`

HasParams returns a boolean if a field has been set.

### GetETag

`func (o *ServiceItem) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *ServiceItem) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *ServiceItem) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *ServiceItem) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


