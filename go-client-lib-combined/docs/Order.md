# Order

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Customer** | Pointer to [**Customer**](Customer.md) |  | [optional] 
**CustomerID** | Pointer to **string** |  | [optional] 
**Entitlement** | Pointer to [**Entitlements**](Entitlements.md) |  | [optional] 
**DestinationDutyLocation** | Pointer to [**DutyLocation**](DutyLocation.md) |  | [optional] 
**OriginDutyLocation** | Pointer to [**DutyLocation**](DutyLocation.md) |  | [optional] 
**OriginDutyLocationGBLOC** | Pointer to **string** |  | [optional] 
**Rank** | **string** |  | 
**ReportByDate** | Pointer to **string** |  | [optional] 
**OrdersType** | Pointer to [**OrdersType**](OrdersType.md) |  | [optional] 
**OrderNumber** | **string** |  | 
**LinesOfAccounting** | **string** |  | 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewOrder

`func NewOrder(rank string, orderNumber string, linesOfAccounting string, ) *Order`

NewOrder instantiates a new Order object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderWithDefaults

`func NewOrderWithDefaults() *Order`

NewOrderWithDefaults instantiates a new Order object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Order) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Order) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Order) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Order) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomer

`func (o *Order) GetCustomer() Customer`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *Order) GetCustomerOk() (*Customer, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *Order) SetCustomer(v Customer)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *Order) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerID

`func (o *Order) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *Order) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *Order) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *Order) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetEntitlement

`func (o *Order) GetEntitlement() Entitlements`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *Order) GetEntitlementOk() (*Entitlements, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *Order) SetEntitlement(v Entitlements)`

SetEntitlement sets Entitlement field to given value.

### HasEntitlement

`func (o *Order) HasEntitlement() bool`

HasEntitlement returns a boolean if a field has been set.

### GetDestinationDutyLocation

`func (o *Order) GetDestinationDutyLocation() DutyLocation`

GetDestinationDutyLocation returns the DestinationDutyLocation field if non-nil, zero value otherwise.

### GetDestinationDutyLocationOk

`func (o *Order) GetDestinationDutyLocationOk() (*DutyLocation, bool)`

GetDestinationDutyLocationOk returns a tuple with the DestinationDutyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDutyLocation

`func (o *Order) SetDestinationDutyLocation(v DutyLocation)`

SetDestinationDutyLocation sets DestinationDutyLocation field to given value.

### HasDestinationDutyLocation

`func (o *Order) HasDestinationDutyLocation() bool`

HasDestinationDutyLocation returns a boolean if a field has been set.

### GetOriginDutyLocation

`func (o *Order) GetOriginDutyLocation() DutyLocation`

GetOriginDutyLocation returns the OriginDutyLocation field if non-nil, zero value otherwise.

### GetOriginDutyLocationOk

`func (o *Order) GetOriginDutyLocationOk() (*DutyLocation, bool)`

GetOriginDutyLocationOk returns a tuple with the OriginDutyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDutyLocation

`func (o *Order) SetOriginDutyLocation(v DutyLocation)`

SetOriginDutyLocation sets OriginDutyLocation field to given value.

### HasOriginDutyLocation

`func (o *Order) HasOriginDutyLocation() bool`

HasOriginDutyLocation returns a boolean if a field has been set.

### GetOriginDutyLocationGBLOC

`func (o *Order) GetOriginDutyLocationGBLOC() string`

GetOriginDutyLocationGBLOC returns the OriginDutyLocationGBLOC field if non-nil, zero value otherwise.

### GetOriginDutyLocationGBLOCOk

`func (o *Order) GetOriginDutyLocationGBLOCOk() (*string, bool)`

GetOriginDutyLocationGBLOCOk returns a tuple with the OriginDutyLocationGBLOC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDutyLocationGBLOC

`func (o *Order) SetOriginDutyLocationGBLOC(v string)`

SetOriginDutyLocationGBLOC sets OriginDutyLocationGBLOC field to given value.

### HasOriginDutyLocationGBLOC

`func (o *Order) HasOriginDutyLocationGBLOC() bool`

HasOriginDutyLocationGBLOC returns a boolean if a field has been set.

### GetRank

`func (o *Order) GetRank() string`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *Order) GetRankOk() (*string, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *Order) SetRank(v string)`

SetRank sets Rank field to given value.


### GetReportByDate

`func (o *Order) GetReportByDate() string`

GetReportByDate returns the ReportByDate field if non-nil, zero value otherwise.

### GetReportByDateOk

`func (o *Order) GetReportByDateOk() (*string, bool)`

GetReportByDateOk returns a tuple with the ReportByDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportByDate

`func (o *Order) SetReportByDate(v string)`

SetReportByDate sets ReportByDate field to given value.

### HasReportByDate

`func (o *Order) HasReportByDate() bool`

HasReportByDate returns a boolean if a field has been set.

### GetOrdersType

`func (o *Order) GetOrdersType() OrdersType`

GetOrdersType returns the OrdersType field if non-nil, zero value otherwise.

### GetOrdersTypeOk

`func (o *Order) GetOrdersTypeOk() (*OrdersType, bool)`

GetOrdersTypeOk returns a tuple with the OrdersType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersType

`func (o *Order) SetOrdersType(v OrdersType)`

SetOrdersType sets OrdersType field to given value.

### HasOrdersType

`func (o *Order) HasOrdersType() bool`

HasOrdersType returns a boolean if a field has been set.

### GetOrderNumber

`func (o *Order) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *Order) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *Order) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.


### GetLinesOfAccounting

`func (o *Order) GetLinesOfAccounting() string`

GetLinesOfAccounting returns the LinesOfAccounting field if non-nil, zero value otherwise.

### GetLinesOfAccountingOk

`func (o *Order) GetLinesOfAccountingOk() (*string, bool)`

GetLinesOfAccountingOk returns a tuple with the LinesOfAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesOfAccounting

`func (o *Order) SetLinesOfAccounting(v string)`

SetLinesOfAccounting sets LinesOfAccounting field to given value.


### GetETag

`func (o *Order) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *Order) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *Order) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *Order) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


