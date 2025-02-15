# Go API client for openapi

The Prime V3 API is a RESTful API that enables the Prime contractor to request
information about upcoming moves, update the details and status of those moves,
and make payment requests. It uses Mutual TLS for authentication procedures.

All endpoints are located at `/prime/v3/`.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.0.1
- Package version: 1.0.0
- Generator version: 7.8.0-SNAPSHOT
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```go
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `openapi.ContextServerIndex` of type `int`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `openapi.ContextServerVariables` of type `map[string]string`.

```go
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `openapi.ContextOperationServerIndices` and `openapi.ContextOperationServerVariables` context maps.

```go
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to */prime/v3*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*MoveTaskOrderAPI* | [**GetMoveTaskOrder**](docs/MoveTaskOrderAPI.md#getmovetaskorder) | **Get** /move-task-orders/{moveID} | getMoveTaskOrder
*MtoShipmentAPI* | [**CreateMTOShipment**](docs/MtoShipmentAPI.md#createmtoshipment) | **Post** /mto-shipments | createMTOShipment
*MtoShipmentAPI* | [**UpdateMTOShipment**](docs/MtoShipmentAPI.md#updatemtoshipment) | **Patch** /mto-shipments/{mtoShipmentID} | updateMTOShipment


## Documentation For Models

 - [AddressV3](docs/AddressV3.md)
 - [ClientErrorV3](docs/ClientErrorV3.md)
 - [CreateMTOShipmentV3](docs/CreateMTOShipmentV3.md)
 - [CreatePPMShipmentV3](docs/CreatePPMShipmentV3.md)
 - [CreateSITAddressUpdateRequestV3](docs/CreateSITAddressUpdateRequestV3.md)
 - [CreateSITExtensionV3](docs/CreateSITExtensionV3.md)
 - [CustomerV3](docs/CustomerV3.md)
 - [DestinationTypeV3](docs/DestinationTypeV3.md)
 - [DutyLocationV3](docs/DutyLocationV3.md)
 - [EntitlementsV3](docs/EntitlementsV3.md)
 - [ErrorV3](docs/ErrorV3.md)
 - [MTOAgentTypeV3](docs/MTOAgentTypeV3.md)
 - [MTOAgentV3](docs/MTOAgentV3.md)
 - [MTOServiceItemBasicV3](docs/MTOServiceItemBasicV3.md)
 - [MTOServiceItemDestSITV3](docs/MTOServiceItemDestSITV3.md)
 - [MTOServiceItemDimensionV3](docs/MTOServiceItemDimensionV3.md)
 - [MTOServiceItemDomesticCratingV3](docs/MTOServiceItemDomesticCratingV3.md)
 - [MTOServiceItemModelTypeV3](docs/MTOServiceItemModelTypeV3.md)
 - [MTOServiceItemOriginSITV3](docs/MTOServiceItemOriginSITV3.md)
 - [MTOServiceItemShuttleV3](docs/MTOServiceItemShuttleV3.md)
 - [MTOServiceItemStatusV3](docs/MTOServiceItemStatusV3.md)
 - [MTOServiceItemV3](docs/MTOServiceItemV3.md)
 - [MTOShipmentTypeV3](docs/MTOShipmentTypeV3.md)
 - [MTOShipmentV3](docs/MTOShipmentV3.md)
 - [MTOShipmentWithoutServiceItemsV3](docs/MTOShipmentWithoutServiceItemsV3.md)
 - [MoveTaskOrderV3](docs/MoveTaskOrderV3.md)
 - [OrderV3](docs/OrderV3.md)
 - [OrdersTypeV3](docs/OrdersTypeV3.md)
 - [PPMShipmentStatusV3](docs/PPMShipmentStatusV3.md)
 - [PPMShipmentV3](docs/PPMShipmentV3.md)
 - [PaymentRequestStatusV3](docs/PaymentRequestStatusV3.md)
 - [PaymentRequestV3](docs/PaymentRequestV3.md)
 - [PaymentServiceItemParamV3](docs/PaymentServiceItemParamV3.md)
 - [PaymentServiceItemStatusV3](docs/PaymentServiceItemStatusV3.md)
 - [PaymentServiceItemV3](docs/PaymentServiceItemV3.md)
 - [ProofOfServiceDocV3](docs/ProofOfServiceDocV3.md)
 - [ReServiceCodeV3](docs/ReServiceCodeV3.md)
 - [ReweighRequesterV3](docs/ReweighRequesterV3.md)
 - [ReweighV3](docs/ReweighV3.md)
 - [SITExtensionV3](docs/SITExtensionV3.md)
 - [SITLocationTypeV3](docs/SITLocationTypeV3.md)
 - [ServiceItemParamNameV3](docs/ServiceItemParamNameV3.md)
 - [ServiceItemParamOriginV3](docs/ServiceItemParamOriginV3.md)
 - [ServiceItemParamTypeV3](docs/ServiceItemParamTypeV3.md)
 - [ServiceRequestDocumentV3](docs/ServiceRequestDocumentV3.md)
 - [ShipmentAddressUpdateStatusV3](docs/ShipmentAddressUpdateStatusV3.md)
 - [ShipmentAddressUpdateV3](docs/ShipmentAddressUpdateV3.md)
 - [SitAddressUpdateStatusV3](docs/SitAddressUpdateStatusV3.md)
 - [SitAddressUpdateV3](docs/SitAddressUpdateV3.md)
 - [StorageFacilityV3](docs/StorageFacilityV3.md)
 - [UpdateMTOServiceItemModelTypeV3](docs/UpdateMTOServiceItemModelTypeV3.md)
 - [UpdateMTOServiceItemSITV3](docs/UpdateMTOServiceItemSITV3.md)
 - [UpdateMTOServiceItemShuttleV3](docs/UpdateMTOServiceItemShuttleV3.md)
 - [UpdateMTOServiceItemV3](docs/UpdateMTOServiceItemV3.md)
 - [UpdateMTOShipmentStatusV3](docs/UpdateMTOShipmentStatusV3.md)
 - [UpdateMTOShipmentStorageFacilityV3](docs/UpdateMTOShipmentStorageFacilityV3.md)
 - [UpdateMTOShipmentV3](docs/UpdateMTOShipmentV3.md)
 - [UpdatePPMShipmentV3](docs/UpdatePPMShipmentV3.md)
 - [UpdateReweighV3](docs/UpdateReweighV3.md)
 - [UpdateShipmentDestinationAddressV3](docs/UpdateShipmentDestinationAddressV3.md)
 - [UploadWithOmissionsV3](docs/UploadWithOmissionsV3.md)
 - [ValidationErrorV3](docs/ValidationErrorV3.md)


## Documentation For Authorization

Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author

milmove-developers@caci.com

