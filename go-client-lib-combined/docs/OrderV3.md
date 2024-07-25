# OrderV3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** |  | [optional] 
**Customer** | Pointer to [**CustomerV3V3**](CustomerV3.md) |  | [optional] 
**CustomerID** | Pointer to **string** |  | [optional] 
**Entitlement** | Pointer to [**EntitlementsV3V3**](EntitlementsV3.md) |  | [optional] 
**DestinationDutyLocation** | Pointer to [**DutyLocationV3V3**](DutyLocationV3.md) |  | [optional] 
**OriginDutyLocation** | Pointer to [**DutyLocationV3V3**](DutyLocationV3.md) |  | [optional] 
**OriginDutyLocationGBLOC** | Pointer to **string** |  | [optional] 
**Rank** | **string** |  | 
**ReportByDate** | Pointer to **string** |  | [optional] 
**OrdersType** | Pointer to [**OrdersTypeV3V3**](OrdersTypeV3.md) |  | [optional] 
**OrderNumber** | **string** |  | 
**LinesOfAccounting** | **string** |  | 
**SupplyAndServicesCostEstimate** | Pointer to **string** |  | [optional] [readonly] 
**PackingAndShippingInstructions** | Pointer to **string** |  | [optional] [readonly] 
**MethodOfPayment** | Pointer to **string** |  | [optional] [readonly] 
**Naics** | Pointer to **string** |  | [optional] [readonly] 
**ETag** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewOrderV3

`func NewOrderV3(rank string, orderNumber string, linesOfAccounting string, ) *OrderV3`

NewOrderV3 instantiates a new OrderV3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOrderV3WithDefaults

`func NewOrderV3WithDefaults() *OrderV3`

NewOrderV3WithDefaults instantiates a new OrderV3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OrderV3) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OrderV3) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OrderV3) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *OrderV3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCustomer

`func (o *OrderV3) GetCustomer() CustomerV3V3`

GetCustomer returns the Customer field if non-nil, zero value otherwise.

### GetCustomerOk

`func (o *OrderV3) GetCustomerOk() (*CustomerV3V3, bool)`

GetCustomerOk returns a tuple with the Customer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomer

`func (o *OrderV3) SetCustomer(v CustomerV3V3)`

SetCustomer sets Customer field to given value.

### HasCustomer

`func (o *OrderV3) HasCustomer() bool`

HasCustomer returns a boolean if a field has been set.

### GetCustomerID

`func (o *OrderV3) GetCustomerID() string`

GetCustomerID returns the CustomerID field if non-nil, zero value otherwise.

### GetCustomerIDOk

`func (o *OrderV3) GetCustomerIDOk() (*string, bool)`

GetCustomerIDOk returns a tuple with the CustomerID field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomerID

`func (o *OrderV3) SetCustomerID(v string)`

SetCustomerID sets CustomerID field to given value.

### HasCustomerID

`func (o *OrderV3) HasCustomerID() bool`

HasCustomerID returns a boolean if a field has been set.

### GetEntitlement

`func (o *OrderV3) GetEntitlement() EntitlementsV3V3`

GetEntitlement returns the Entitlement field if non-nil, zero value otherwise.

### GetEntitlementOk

`func (o *OrderV3) GetEntitlementOk() (*EntitlementsV3V3, bool)`

GetEntitlementOk returns a tuple with the Entitlement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntitlement

`func (o *OrderV3) SetEntitlement(v EntitlementsV3V3)`

SetEntitlement sets Entitlement field to given value.

### HasEntitlement

`func (o *OrderV3) HasEntitlement() bool`

HasEntitlement returns a boolean if a field has been set.

### GetDestinationDutyLocation

`func (o *OrderV3) GetDestinationDutyLocation() DutyLocationV3V3`

GetDestinationDutyLocation returns the DestinationDutyLocation field if non-nil, zero value otherwise.

### GetDestinationDutyLocationOk

`func (o *OrderV3) GetDestinationDutyLocationOk() (*DutyLocationV3V3, bool)`

GetDestinationDutyLocationOk returns a tuple with the DestinationDutyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationDutyLocation

`func (o *OrderV3) SetDestinationDutyLocation(v DutyLocationV3V3)`

SetDestinationDutyLocation sets DestinationDutyLocation field to given value.

### HasDestinationDutyLocation

`func (o *OrderV3) HasDestinationDutyLocation() bool`

HasDestinationDutyLocation returns a boolean if a field has been set.

### GetOriginDutyLocation

`func (o *OrderV3) GetOriginDutyLocation() DutyLocationV3V3`

GetOriginDutyLocation returns the OriginDutyLocation field if non-nil, zero value otherwise.

### GetOriginDutyLocationOk

`func (o *OrderV3) GetOriginDutyLocationOk() (*DutyLocationV3V3, bool)`

GetOriginDutyLocationOk returns a tuple with the OriginDutyLocation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDutyLocation

`func (o *OrderV3) SetOriginDutyLocation(v DutyLocationV3V3)`

SetOriginDutyLocation sets OriginDutyLocation field to given value.

### HasOriginDutyLocation

`func (o *OrderV3) HasOriginDutyLocation() bool`

HasOriginDutyLocation returns a boolean if a field has been set.

### GetOriginDutyLocationGBLOC

`func (o *OrderV3) GetOriginDutyLocationGBLOC() string`

GetOriginDutyLocationGBLOC returns the OriginDutyLocationGBLOC field if non-nil, zero value otherwise.

### GetOriginDutyLocationGBLOCOk

`func (o *OrderV3) GetOriginDutyLocationGBLOCOk() (*string, bool)`

GetOriginDutyLocationGBLOCOk returns a tuple with the OriginDutyLocationGBLOC field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginDutyLocationGBLOC

`func (o *OrderV3) SetOriginDutyLocationGBLOC(v string)`

SetOriginDutyLocationGBLOC sets OriginDutyLocationGBLOC field to given value.

### HasOriginDutyLocationGBLOC

`func (o *OrderV3) HasOriginDutyLocationGBLOC() bool`

HasOriginDutyLocationGBLOC returns a boolean if a field has been set.

### GetRank

`func (o *OrderV3) GetRank() string`

GetRank returns the Rank field if non-nil, zero value otherwise.

### GetRankOk

`func (o *OrderV3) GetRankOk() (*string, bool)`

GetRankOk returns a tuple with the Rank field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRank

`func (o *OrderV3) SetRank(v string)`

SetRank sets Rank field to given value.


### GetReportByDate

`func (o *OrderV3) GetReportByDate() string`

GetReportByDate returns the ReportByDate field if non-nil, zero value otherwise.

### GetReportByDateOk

`func (o *OrderV3) GetReportByDateOk() (*string, bool)`

GetReportByDateOk returns a tuple with the ReportByDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReportByDate

`func (o *OrderV3) SetReportByDate(v string)`

SetReportByDate sets ReportByDate field to given value.

### HasReportByDate

`func (o *OrderV3) HasReportByDate() bool`

HasReportByDate returns a boolean if a field has been set.

### GetOrdersType

`func (o *OrderV3) GetOrdersType() OrdersTypeV3V3`

GetOrdersType returns the OrdersType field if non-nil, zero value otherwise.

### GetOrdersTypeOk

`func (o *OrderV3) GetOrdersTypeOk() (*OrdersTypeV3V3, bool)`

GetOrdersTypeOk returns a tuple with the OrdersType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrdersType

`func (o *OrderV3) SetOrdersType(v OrdersTypeV3V3)`

SetOrdersType sets OrdersType field to given value.

### HasOrdersType

`func (o *OrderV3) HasOrdersType() bool`

HasOrdersType returns a boolean if a field has been set.

### GetOrderNumber

`func (o *OrderV3) GetOrderNumber() string`

GetOrderNumber returns the OrderNumber field if non-nil, zero value otherwise.

### GetOrderNumberOk

`func (o *OrderV3) GetOrderNumberOk() (*string, bool)`

GetOrderNumberOk returns a tuple with the OrderNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrderNumber

`func (o *OrderV3) SetOrderNumber(v string)`

SetOrderNumber sets OrderNumber field to given value.


### GetLinesOfAccounting

`func (o *OrderV3) GetLinesOfAccounting() string`

GetLinesOfAccounting returns the LinesOfAccounting field if non-nil, zero value otherwise.

### GetLinesOfAccountingOk

`func (o *OrderV3) GetLinesOfAccountingOk() (*string, bool)`

GetLinesOfAccountingOk returns a tuple with the LinesOfAccounting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinesOfAccounting

`func (o *OrderV3) SetLinesOfAccounting(v string)`

SetLinesOfAccounting sets LinesOfAccounting field to given value.


### GetSupplyAndServicesCostEstimate

`func (o *OrderV3) GetSupplyAndServicesCostEstimate() string`

GetSupplyAndServicesCostEstimate returns the SupplyAndServicesCostEstimate field if non-nil, zero value otherwise.

### GetSupplyAndServicesCostEstimateOk

`func (o *OrderV3) GetSupplyAndServicesCostEstimateOk() (*string, bool)`

GetSupplyAndServicesCostEstimateOk returns a tuple with the SupplyAndServicesCostEstimate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupplyAndServicesCostEstimate

`func (o *OrderV3) SetSupplyAndServicesCostEstimate(v string)`

SetSupplyAndServicesCostEstimate sets SupplyAndServicesCostEstimate field to given value.

### HasSupplyAndServicesCostEstimate

`func (o *OrderV3) HasSupplyAndServicesCostEstimate() bool`

HasSupplyAndServicesCostEstimate returns a boolean if a field has been set.

### GetPackingAndShippingInstructions

`func (o *OrderV3) GetPackingAndShippingInstructions() string`

GetPackingAndShippingInstructions returns the PackingAndShippingInstructions field if non-nil, zero value otherwise.

### GetPackingAndShippingInstructionsOk

`func (o *OrderV3) GetPackingAndShippingInstructionsOk() (*string, bool)`

GetPackingAndShippingInstructionsOk returns a tuple with the PackingAndShippingInstructions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPackingAndShippingInstructions

`func (o *OrderV3) SetPackingAndShippingInstructions(v string)`

SetPackingAndShippingInstructions sets PackingAndShippingInstructions field to given value.

### HasPackingAndShippingInstructions

`func (o *OrderV3) HasPackingAndShippingInstructions() bool`

HasPackingAndShippingInstructions returns a boolean if a field has been set.

### GetMethodOfPayment

`func (o *OrderV3) GetMethodOfPayment() string`

GetMethodOfPayment returns the MethodOfPayment field if non-nil, zero value otherwise.

### GetMethodOfPaymentOk

`func (o *OrderV3) GetMethodOfPaymentOk() (*string, bool)`

GetMethodOfPaymentOk returns a tuple with the MethodOfPayment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMethodOfPayment

`func (o *OrderV3) SetMethodOfPayment(v string)`

SetMethodOfPayment sets MethodOfPayment field to given value.

### HasMethodOfPayment

`func (o *OrderV3) HasMethodOfPayment() bool`

HasMethodOfPayment returns a boolean if a field has been set.

### GetNaics

`func (o *OrderV3) GetNaics() string`

GetNaics returns the Naics field if non-nil, zero value otherwise.

### GetNaicsOk

`func (o *OrderV3) GetNaicsOk() (*string, bool)`

GetNaicsOk returns a tuple with the Naics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNaics

`func (o *OrderV3) SetNaics(v string)`

SetNaics sets Naics field to given value.

### HasNaics

`func (o *OrderV3) HasNaics() bool`

HasNaics returns a boolean if a field has been set.

### GetETag

`func (o *OrderV3) GetETag() string`

GetETag returns the ETag field if non-nil, zero value otherwise.

### GetETagOk

`func (o *OrderV3) GetETagOk() (*string, bool)`

GetETagOk returns a tuple with the ETag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetETag

`func (o *OrderV3) SetETag(v string)`

SetETag sets ETag field to given value.

### HasETag

`func (o *OrderV3) HasETag() bool`

HasETag returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


