# \MtoShipmentAPI

All URIs are relative to */prime/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMTOAgent**](MtoShipmentAPI.md#CreateMTOAgent) | **Post** /mto-shipments/{mtoShipmentID}/agents | createMTOAgent
[**CreateMTOShipment**](MtoShipmentAPI.md#CreateMTOShipment) | **Post** /mto-shipments | createMTOShipment
[**CreateSITExtension**](MtoShipmentAPI.md#CreateSITExtension) | **Post** /mto-shipments/{mtoShipmentID}/sit-extensions | createSITExtension
[**DeleteMTOShipment**](MtoShipmentAPI.md#DeleteMTOShipment) | **Delete** /mto-shipments/{mtoShipmentID} | deleteMTOShipment
[**UpdateMTOAgent**](MtoShipmentAPI.md#UpdateMTOAgent) | **Put** /mto-shipments/{mtoShipmentID}/agents/{agentID} | updateMTOAgent
[**UpdateMTOShipment**](MtoShipmentAPI.md#UpdateMTOShipment) | **Patch** /mto-shipments/{mtoShipmentID} | updateMTOShipment
[**UpdateMTOShipmentAddress**](MtoShipmentAPI.md#UpdateMTOShipmentAddress) | **Put** /mto-shipments/{mtoShipmentID}/addresses/{addressID} | updateMTOShipmentAddress
[**UpdateMTOShipmentStatus**](MtoShipmentAPI.md#UpdateMTOShipmentStatus) | **Patch** /mto-shipments/{mtoShipmentID}/status | updateMTOShipmentStatus
[**UpdateReweigh**](MtoShipmentAPI.md#UpdateReweigh) | **Patch** /mto-shipments/{mtoShipmentID}/reweighs/{reweighID} | updateReweigh
[**UpdateShipmentDestinationAddress**](MtoShipmentAPI.md#UpdateShipmentDestinationAddress) | **Post** /mto-shipments/{mtoShipmentID}/shipment-address-updates | updateShipmentDestinationAddress



## CreateMTOAgent

> MTOAgent CreateMTOAgent(ctx, mtoShipmentID).Body(body).Execute()

createMTOAgent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	mtoShipmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the shipment associated with the agent
	body := *openapiclient.NewMTOAgent() // MTOAgent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtoShipmentAPI.CreateMTOAgent(context.Background(), mtoShipmentID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoShipmentAPI.CreateMTOAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMTOAgent`: MTOAgent
	fmt.Fprintf(os.Stdout, "Response from `MtoShipmentAPI.CreateMTOAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mtoShipmentID** | **string** | UUID of the shipment associated with the agent | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateMTOAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**MTOAgent**](MTOAgent.md) |  | 

### Return type

[**MTOAgent**](MTOAgent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateMTOShipment

> MTOShipment CreateMTOShipment(ctx).Body(body).Execute()

createMTOShipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	body := *openapiclient.NewCreateMTOShipment("1f2270c7-7166-40ae-981e-b200ebdf3054", openapiclient.MTOShipmentType("BOAT_HAUL_AWAY")) // CreateMTOShipment |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtoShipmentAPI.CreateMTOShipment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoShipmentAPI.CreateMTOShipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMTOShipment`: MTOShipment
	fmt.Fprintf(os.Stdout, "Response from `MtoShipmentAPI.CreateMTOShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMTOShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateMTOShipment**](CreateMTOShipment.md) |  | 

### Return type

[**MTOShipment**](MTOShipment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateSITExtension

> SITExtension CreateSITExtension(ctx, mtoShipmentID).Body(body).Execute()

createSITExtension



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	mtoShipmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the shipment associated with the agent
	body := *openapiclient.NewCreateSITExtension("RequestReason_example", "We need SIT additional days. The customer has not found a house yet.", int32(30)) // CreateSITExtension | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtoShipmentAPI.CreateSITExtension(context.Background(), mtoShipmentID).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoShipmentAPI.CreateSITExtension``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateSITExtension`: SITExtension
	fmt.Fprintf(os.Stdout, "Response from `MtoShipmentAPI.CreateSITExtension`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mtoShipmentID** | **string** | UUID of the shipment associated with the agent | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateSITExtensionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CreateSITExtension**](CreateSITExtension.md) |  | 

### Return type

[**SITExtension**](SITExtension.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteMTOShipment

> DeleteMTOShipment(ctx, mtoShipmentID).Execute()

deleteMTOShipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	mtoShipmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the shipment to be deleted

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.MtoShipmentAPI.DeleteMTOShipment(context.Background(), mtoShipmentID).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoShipmentAPI.DeleteMTOShipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mtoShipmentID** | **string** | UUID of the shipment to be deleted | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteMTOShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMTOAgent

> MTOAgent UpdateMTOAgent(ctx, mtoShipmentID, agentID).IfMatch(ifMatch).Body(body).Execute()

updateMTOAgent



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	mtoShipmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the shipment associated with the agent
	agentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the agent being updated
	ifMatch := "ifMatch_example" // string | Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 
	body := *openapiclient.NewMTOAgent() // MTOAgent | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtoShipmentAPI.UpdateMTOAgent(context.Background(), mtoShipmentID, agentID).IfMatch(ifMatch).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoShipmentAPI.UpdateMTOAgent``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMTOAgent`: MTOAgent
	fmt.Fprintf(os.Stdout, "Response from `MtoShipmentAPI.UpdateMTOAgent`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mtoShipmentID** | **string** | UUID of the shipment associated with the agent | 
**agentID** | **string** | UUID of the agent being updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMTOAgentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  | 
 **body** | [**MTOAgent**](MTOAgent.md) |  | 

### Return type

[**MTOAgent**](MTOAgent.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMTOShipment

> MTOShipment UpdateMTOShipment(ctx, mtoShipmentID).IfMatch(ifMatch).Body(body).Execute()

updateMTOShipment



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	mtoShipmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the shipment being updated.
	ifMatch := "ifMatch_example" // string | Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 
	body := *openapiclient.NewUpdateMTOShipment() // UpdateMTOShipment | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtoShipmentAPI.UpdateMTOShipment(context.Background(), mtoShipmentID).IfMatch(ifMatch).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoShipmentAPI.UpdateMTOShipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMTOShipment`: MTOShipment
	fmt.Fprintf(os.Stdout, "Response from `MtoShipmentAPI.UpdateMTOShipment`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mtoShipmentID** | **string** | UUID of the shipment being updated. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMTOShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  | 
 **body** | [**UpdateMTOShipment**](UpdateMTOShipment.md) |  | 

### Return type

[**MTOShipment**](MTOShipment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMTOShipmentAddress

> Address UpdateMTOShipmentAddress(ctx, mtoShipmentID, addressID).IfMatch(ifMatch).Body(body).Execute()

updateMTOShipmentAddress



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	mtoShipmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the shipment associated with the address
	addressID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the address being updated
	ifMatch := "ifMatch_example" // string | Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 
	body := *openapiclient.NewAddress("123 Main Ave", "Anytown", "State_example", "90210") // Address | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtoShipmentAPI.UpdateMTOShipmentAddress(context.Background(), mtoShipmentID, addressID).IfMatch(ifMatch).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoShipmentAPI.UpdateMTOShipmentAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMTOShipmentAddress`: Address
	fmt.Fprintf(os.Stdout, "Response from `MtoShipmentAPI.UpdateMTOShipmentAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mtoShipmentID** | **string** | UUID of the shipment associated with the address | 
**addressID** | **string** | UUID of the address being updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMTOShipmentAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  | 
 **body** | [**Address**](Address.md) |  | 

### Return type

[**Address**](Address.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMTOShipmentStatus

> MTOShipment UpdateMTOShipmentStatus(ctx, mtoShipmentID).IfMatch(ifMatch).Body(body).Execute()

updateMTOShipmentStatus



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	mtoShipmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the shipment associated with the agent
	ifMatch := "ifMatch_example" // string | Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 
	body := *openapiclient.NewUpdateMTOShipmentStatus() // UpdateMTOShipmentStatus | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtoShipmentAPI.UpdateMTOShipmentStatus(context.Background(), mtoShipmentID).IfMatch(ifMatch).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoShipmentAPI.UpdateMTOShipmentStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMTOShipmentStatus`: MTOShipment
	fmt.Fprintf(os.Stdout, "Response from `MtoShipmentAPI.UpdateMTOShipmentStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mtoShipmentID** | **string** | UUID of the shipment associated with the agent | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMTOShipmentStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  | 
 **body** | [**UpdateMTOShipmentStatus**](UpdateMTOShipmentStatus.md) |  | 

### Return type

[**MTOShipment**](MTOShipment.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateReweigh

> Reweigh UpdateReweigh(ctx, mtoShipmentID, reweighID).IfMatch(ifMatch).Body(body).Execute()

updateReweigh



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	mtoShipmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the shipment associated with the reweigh
	reweighID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the reweigh being updated
	ifMatch := "ifMatch_example" // string | Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 
	body := *openapiclient.NewUpdateReweigh() // UpdateReweigh | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtoShipmentAPI.UpdateReweigh(context.Background(), mtoShipmentID, reweighID).IfMatch(ifMatch).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoShipmentAPI.UpdateReweigh``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateReweigh`: Reweigh
	fmt.Fprintf(os.Stdout, "Response from `MtoShipmentAPI.UpdateReweigh`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mtoShipmentID** | **string** | UUID of the shipment associated with the reweigh | 
**reweighID** | **string** | UUID of the reweigh being updated | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateReweighRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **ifMatch** | **string** | Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  | 
 **body** | [**UpdateReweigh**](UpdateReweigh.md) |  | 

### Return type

[**Reweigh**](Reweigh.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateShipmentDestinationAddress

> ShipmentAddressUpdate UpdateShipmentDestinationAddress(ctx, mtoShipmentID).IfMatch(ifMatch).Body(body).Execute()

updateShipmentDestinationAddress



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	mtoShipmentID := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | UUID of the shipment associated with the address
	ifMatch := "ifMatch_example" // string | Needs to be the eTag of the mtoShipment. Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 
	body := *openapiclient.NewUpdateShipmentDestinationAddress(*openapiclient.NewAddress("123 Main Ave", "Anytown", "State_example", "90210"), "Customer reached out to me this week and let me know they want to move somewhere else.") // UpdateShipmentDestinationAddress | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtoShipmentAPI.UpdateShipmentDestinationAddress(context.Background(), mtoShipmentID).IfMatch(ifMatch).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoShipmentAPI.UpdateShipmentDestinationAddress``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateShipmentDestinationAddress`: ShipmentAddressUpdate
	fmt.Fprintf(os.Stdout, "Response from `MtoShipmentAPI.UpdateShipmentDestinationAddress`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mtoShipmentID** | **string** | UUID of the shipment associated with the address | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateShipmentDestinationAddressRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Needs to be the eTag of the mtoShipment. Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  | 
 **body** | [**UpdateShipmentDestinationAddress**](UpdateShipmentDestinationAddress.md) |  | 

### Return type

[**ShipmentAddressUpdate**](ShipmentAddressUpdate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

