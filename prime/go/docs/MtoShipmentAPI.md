# \MtoShipmentAPI

All URIs are relative to */prime/v3*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMTOShipment**](MtoShipmentAPI.md#CreateMTOShipment) | **Post** /mto-shipments | createMTOShipment
[**UpdateMTOShipment**](MtoShipmentAPI.md#UpdateMTOShipment) | **Patch** /mto-shipments/{mtoShipmentID} | updateMTOShipment



## CreateMTOShipment

> MTOShipmentV3V3 CreateMTOShipment(ctx).Body(body).Execute()

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
	body := *openapiclient.NewCreateMTOShipmentV3("1f2270c7-7166-40ae-981e-b200ebdf3054", "TODO") // CreateMTOShipmentV3 |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtoShipmentAPI.CreateMTOShipment(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoShipmentAPI.CreateMTOShipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMTOShipment`: MTOShipmentV3V3
	fmt.Fprintf(os.Stdout, "Response from `MtoShipmentAPI.CreateMTOShipment`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMTOShipmentRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateMTOShipmentV3**](CreateMTOShipmentV3.md) |  | 

### Return type

[**MTOShipmentV3V3**](MTOShipmentV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMTOShipment

> MTOShipmentV3V3 UpdateMTOShipment(ctx, mtoShipmentID).IfMatch(ifMatch).Body(body).Execute()

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
	body := *openapiclient.NewUpdateMTOShipmentV3() // UpdateMTOShipmentV3 | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtoShipmentAPI.UpdateMTOShipment(context.Background(), mtoShipmentID).IfMatch(ifMatch).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoShipmentAPI.UpdateMTOShipment``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMTOShipment`: MTOShipmentV3V3
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
 **body** | [**UpdateMTOShipmentV3**](UpdateMTOShipmentV3.md) |  | 

### Return type

[**MTOShipmentV3V3**](MTOShipmentV3.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

