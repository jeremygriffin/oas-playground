# \MtoServiceItemAPI

All URIs are relative to */prime/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMTOServiceItem**](MtoServiceItemAPI.md#CreateMTOServiceItem) | **Post** /mto-service-items | createMTOServiceItem
[**CreateServiceRequestDocumentUpload**](MtoServiceItemAPI.md#CreateServiceRequestDocumentUpload) | **Post** /mto-service-items/{mtoServiceItemID}/uploads | createServiceRequestDocumentUpload
[**UpdateMTOServiceItem**](MtoServiceItemAPI.md#UpdateMTOServiceItem) | **Patch** /mto-service-items/{mtoServiceItemID} | updateMTOServiceItem



## CreateMTOServiceItem

> []MTOServiceItem CreateMTOServiceItem(ctx).Body(body).Execute()

createMTOServiceItem



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
	body := *openapiclient.NewMTOServiceItem("1f2270c7-7166-40ae-981e-b200ebdf3054", openapiclient.MTOServiceItemModelType("MTOServiceItemBasic")) // MTOServiceItem |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtoServiceItemAPI.CreateMTOServiceItem(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoServiceItemAPI.CreateMTOServiceItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMTOServiceItem`: []MTOServiceItem
	fmt.Fprintf(os.Stdout, "Response from `MtoServiceItemAPI.CreateMTOServiceItem`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMTOServiceItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**MTOServiceItem**](MTOServiceItem.md) |  | 

### Return type

[**[]MTOServiceItem**](MTOServiceItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateServiceRequestDocumentUpload

> UploadWithOmissions CreateServiceRequestDocumentUpload(ctx, mtoServiceItemID).File(file).Execute()

createServiceRequestDocumentUpload



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
	mtoServiceItemID := "mtoServiceItemID_example" // string | UUID of the service item to use.
	file := os.NewFile(1234, "some_file") // *os.File | The file to upload.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtoServiceItemAPI.CreateServiceRequestDocumentUpload(context.Background(), mtoServiceItemID).File(file).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoServiceItemAPI.CreateServiceRequestDocumentUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateServiceRequestDocumentUpload`: UploadWithOmissions
	fmt.Fprintf(os.Stdout, "Response from `MtoServiceItemAPI.CreateServiceRequestDocumentUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mtoServiceItemID** | **string** | UUID of the service item to use. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateServiceRequestDocumentUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The file to upload. | 

### Return type

[**UploadWithOmissions**](UploadWithOmissions.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateMTOServiceItem

> MTOServiceItem UpdateMTOServiceItem(ctx, mtoServiceItemID).IfMatch(ifMatch).Body(body).Execute()

updateMTOServiceItem



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
	mtoServiceItemID := "mtoServiceItemID_example" // string | UUID of service item to update.
	ifMatch := "ifMatch_example" // string | Optimistic locking is implemented via the `If-Match` header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a `412 Precondition Failed` error. 
	body := *openapiclient.NewUpdateMTOServiceItem(openapiclient.UpdateMTOServiceItemModelType("UpdateMTOServiceItemSIT")) // UpdateMTOServiceItem | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MtoServiceItemAPI.UpdateMTOServiceItem(context.Background(), mtoServiceItemID).IfMatch(ifMatch).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MtoServiceItemAPI.UpdateMTOServiceItem``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateMTOServiceItem`: MTOServiceItem
	fmt.Fprintf(os.Stdout, "Response from `MtoServiceItemAPI.UpdateMTOServiceItem`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**mtoServiceItemID** | **string** | UUID of service item to update. | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateMTOServiceItemRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **ifMatch** | **string** | Optimistic locking is implemented via the &#x60;If-Match&#x60; header. If the ETag header does not match the value of the resource on the server, the server rejects the change with a &#x60;412 Precondition Failed&#x60; error.  | 
 **body** | [**UpdateMTOServiceItem**](UpdateMTOServiceItem.md) |  | 

### Return type

[**MTOServiceItem**](MTOServiceItem.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

