# \PaymentRequestAPI

All URIs are relative to */prime/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaymentRequest**](PaymentRequestAPI.md#CreatePaymentRequest) | **Post** /payment-requests | createPaymentRequest
[**CreateUpload**](PaymentRequestAPI.md#CreateUpload) | **Post** /payment-requests/{paymentRequestID}/uploads | createUpload



## CreatePaymentRequest

> PaymentRequest CreatePaymentRequest(ctx).Body(body).Execute()

createPaymentRequest



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
	body := *openapiclient.NewCreatePaymentRequest("c56a4180-65aa-42ec-a945-5fd21dec0538", []openapiclient.ServiceItem{*openapiclient.NewServiceItem()}) // CreatePaymentRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRequestAPI.CreatePaymentRequest(context.Background()).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestAPI.CreatePaymentRequest``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreatePaymentRequest`: PaymentRequest
	fmt.Fprintf(os.Stdout, "Response from `PaymentRequestAPI.CreatePaymentRequest`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentRequestRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreatePaymentRequest**](CreatePaymentRequest.md) |  | 

### Return type

[**PaymentRequest**](PaymentRequest.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUpload

> UploadWithOmissions CreateUpload(ctx, paymentRequestID).File(file).IsWeightTicket(isWeightTicket).Execute()

createUpload



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
	paymentRequestID := "paymentRequestID_example" // string | UUID of payment request to use.
	file := os.NewFile(1234, "some_file") // *os.File | The file to upload.
	isWeightTicket := true // bool | Indicates whether the file is a weight ticket. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.PaymentRequestAPI.CreateUpload(context.Background(), paymentRequestID).File(file).IsWeightTicket(isWeightTicket).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `PaymentRequestAPI.CreateUpload``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateUpload`: UploadWithOmissions
	fmt.Fprintf(os.Stdout, "Response from `PaymentRequestAPI.CreateUpload`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paymentRequestID** | **string** | UUID of payment request to use. | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUploadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **file** | ***os.File** | The file to upload. | 
 **isWeightTicket** | **bool** | Indicates whether the file is a weight ticket. | 

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

