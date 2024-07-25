# OrderV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Customer** | Pointer to [**CustomerV2V2**](CustomerV2.md) |  | [optional] 
**CustomerID** | Pointer to **string** |  | [optional] 
**Entitlement** | Pointer to [**EntitlementsV2V2**](EntitlementsV2.md) |  | [optional] 
**DestinationDutyLocation** | Pointer to [**DutyLocationV2V2**](DutyLocationV2.md) |  | [optional] 
**OriginDutyLocation** | Pointer to [**DutyLocationV2V2**](DutyLocationV2.md) |  | [optional] 
**OriginDutyLocationGBLOC** | Pointer to **string** |  | [optional] 
**Rank** | **string** |  | 
**ReportByDate** | Pointer to **string** |  | [optional] 
**OrdersType** | Pointer to [**OrdersTypeV2V2**](OrdersTypeV2.md) |  | [optional] 
**OrderNumber** | **string** |  | 
**LinesOfAccounting** | **string** |  | 
**SupplyAndServicesCostEstimate** | Pointer to **string** |  | [optional] [readonly] 
**PackingAndShippingInstructions** | Pointer to **string** |  | [optional] [readonly] 
**MethodOfPayment** | Pointer to **string** |  | [optional] [readonly] 
**Naics** | Pointer to **string** |  | [optional] [readonly] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewOrderV2

`func NewOrderV2(rank string, orderNumber string, linesOfAccounting string, ) *OrderV2`

NewOrderV2 instantiates a new OrderV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderV2WithDefaults

`func NewOrderV2WithDefaults() *OrderV2`

NewOrderV2WithDefaults instantiates a new OrderV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderV2) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderV2) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomer

`func (o *OrderV2) GetCustomer() CustomerV2V2`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderV2) GetCustomerOk() (*CustomerV2V2, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderV2) SetCustomer(v CustomerV2V2)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *OrderV2) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerID

`func (o *OrderV2) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *OrderV2) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *OrderV2) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *OrderV2) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetEntitlement

`func (o *OrderV2) GetEntitlement() EntitlementsV2V2`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *OrderV2) GetEntitlementOk() (*EntitlementsV2V2, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *OrderV2) SetEntitlement(v EntitlementsV2V2)`

SetEntitlement sets Entitlement field to given value.

### HasEntitlement

`func (o *OrderV2) HasEntitlement() bool`

HasEntitlement returns a boolean if a field has been set.

### GetDestinationDutyLocation

`func (o *OrderV2) GetDestinationDutyLocation() DutyLocationV2V2`

GetDestinationDutyLocation returns the DestinationDutyLocation field if non-nil, zero value otherwise.

### GetDestinationDutyLocationOk

`func (o *OrderV2) GetDestinationDutyLocationOk() (*DutyLocationV2V2, bool)`

GetDestinationDutyLocationOk returns a tuple with the DestinationDutyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDutyLocation

`func (o *OrderV2) SetDestinationDutyLocation(v DutyLocationV2V2)`

SetDestinationDutyLocation sets DestinationDutyLocation field to given value.

### HasDestinationDutyLocation

`func (o *OrderV2) HasDestinationDutyLocation() bool`

HasDestinationDutyLocation returns a boolean if a field has been set.

### GetOriginDutyLocation

`func (o *OrderV2) GetOriginDutyLocation() DutyLocationV2V2`

GetOriginDutyLocation returns the OriginDutyLocation field if non-nil, zero value otherwise.

### GetOriginDutyLocationOk

`func (o *OrderV2) GetOriginDutyLocationOk() (*DutyLocationV2V2, bool)`

GetOriginDutyLocationOk returns a tuple with the OriginDutyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDutyLocation

`func (o *OrderV2) SetOriginDutyLocation(v DutyLocationV2V2)`

SetOriginDutyLocation sets OriginDutyLocation field to given value.

### HasOriginDutyLocation

`func (o *OrderV2) HasOriginDutyLocation() bool`

HasOriginDutyLocation returns a boolean if a field has been set.

### GetOriginDutyLocationGBLOC

`func (o *OrderV2) GetOriginDutyLocationGBLOC() string`

GetOriginDutyLocationGBLOC returns the OriginDutyLocationGBLOC field if non-nil, zero value otherwise.

### GetOriginDutyLocationGBLOCOk

`func (o *OrderV2) GetOriginDutyLocationGBLOCOk() (*string, bool)`

GetOriginDutyLocationGBLOCOk returns a tuple with the OriginDutyLocationGBLOC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDutyLocationGBLOC

`func (o *OrderV2) SetOriginDutyLocationGBLOC(v string)`

SetOriginDutyLocationGBLOC sets OriginDutyLocationGBLOC field to given value.

### HasOriginDutyLocationGBLOC

`func (o *OrderV2) HasOriginDutyLocationGBLOC() bool`

HasOriginDutyLocationGBLOC returns a boolean if a field has been set.

### GetRank

`func (o *OrderV2) GetRank() string`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *OrderV2) GetRankOk() (*string, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *OrderV2) SetRank(v string)`

SetRank sets Rank field to given value.


### GetReportByDate

`func (o *OrderV2) GetReportByDate() string`

GetReportByDate returns the ReportByDate field if non-nil, zero value otherwise.

### GetReportByDateOk

`func (o *OrderV2) GetReportByDateOk() (*string, bool)`

GetReportByDateOk returns a tuple with the ReportByDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportByDate

`func (o *OrderV2) SetReportByDate(v string)`

SetReportByDate sets ReportByDate field to given value.

### HasReportByDate

`func (o *OrderV2) HasReportByDate() bool`

HasReportByDate returns a boolean if a field has been set.

### GetOrdersType

`func (o *OrderV2) GetOrdersType() OrdersTypeV2V2`

GetOrdersType returns the OrdersType field if non-nil, zero value otherwise.

### GetOrdersTypeOk

`func (o *OrderV2) GetOrdersTypeOk() (*OrdersTypeV2V2, bool)`

GetOrdersTypeOk returns a tuple with the OrdersType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersType

`func (o *OrderV2) SetOrdersType(v OrdersTypeV2V2)`

SetOrdersType sets OrdersType field to given value.

### HasOrdersType

`func (o *OrderV2) HasOrdersType() bool`

HasOrdersType returns a boolean if a field has been set.

### GetOrderNumber

`func (o *OrderV2) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *OrderV2) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *OrderV2) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.


### GetLinesOfAccounting

`func (o *OrderV2) GetLinesOfAccounting() string`

GetLinesOfAccounting returns the LinesOfAccounting field if non-nil, zero value otherwise.

### GetLinesOfAccountingOk

`func (o *OrderV2) GetLinesOfAccountingOk() (*string, bool)`

GetLinesOfAccountingOk returns a tuple with the LinesOfAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesOfAccounting

`func (o *OrderV2) SetLinesOfAccounting(v string)`

SetLinesOfAccounting sets LinesOfAccounting field to given value.


### GetSupplyAndServicesCostEstimate

`func (o *OrderV2) GetSupplyAndServicesCostEstimate() string`

GetSupplyAndServicesCostEstimate returns the SupplyAndServicesCostEstimate field if non-nil, zero value otherwise.

### GetSupplyAndServicesCostEstimateOk

`func (o *OrderV2) GetSupplyAndServicesCostEstimateOk() (*string, bool)`

GetSupplyAndServicesCostEstimateOk returns a tuple with the SupplyAndServicesCostEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplyAndServicesCostEstimate

`func (o *OrderV2) SetSupplyAndServicesCostEstimate(v string)`

SetSupplyAndServicesCostEstimate sets SupplyAndServicesCostEstimate field to given value.

### HasSupplyAndServicesCostEstimate

`func (o *OrderV2) HasSupplyAndServicesCostEstimate() bool`

HasSupplyAndServicesCostEstimate returns a boolean if a field has been set.

### GetPackingAndShippingInstructions

`func (o *OrderV2) GetPackingAndShippingInstructions() string`

GetPackingAndShippingInstructions returns the PackingAndShippingInstructions field if non-nil, zero value otherwise.

### GetPackingAndShippingInstructionsOk

`func (o *OrderV2) GetPackingAndShippingInstructionsOk() (*string, bool)`

GetPackingAndShippingInstructionsOk returns a tuple with the PackingAndShippingInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackingAndShippingInstructions

`func (o *OrderV2) SetPackingAndShippingInstructions(v string)`

SetPackingAndShippingInstructions sets PackingAndShippingInstructions field to given value.

### HasPackingAndShippingInstructions

`func (o *OrderV2) HasPackingAndShippingInstructions() bool`

HasPackingAndShippingInstructions returns a boolean if a field has been set.

### GetMethodOfPayment

`func (o *OrderV2) GetMethodOfPayment() string`

GetMethodOfPayment returns the MethodOfPayment field if non-nil, zero value otherwise.

### GetMethodOfPaymentOk

`func (o *OrderV2) GetMethodOfPaymentOk() (*string, bool)`

GetMethodOfPaymentOk returns a tuple with the MethodOfPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodOfPayment

`func (o *OrderV2) SetMethodOfPayment(v string)`

SetMethodOfPayment sets MethodOfPayment field to given value.

### HasMethodOfPayment

`func (o *OrderV2) HasMethodOfPayment() bool`

HasMethodOfPayment returns a boolean if a field has been set.

### GetNaics

`func (o *OrderV2) GetNaics() string`

GetNaics returns the Naics field if non-nil, zero value otherwise.

### GetNaicsOk

`func (o *OrderV2) GetNaicsOk() (*string, bool)`

GetNaicsOk returns a tuple with the Naics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaics

`func (o *OrderV2) SetNaics(v string)`

SetNaics sets Naics field to given value.

### HasNaics

`func (o *OrderV2) HasNaics() bool`

HasNaics returns a boolean if a field has been set.

### GetETag

`func (o *OrderV2) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *OrderV2) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *OrderV2) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *OrderV2) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


